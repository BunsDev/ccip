// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {IBaseOffRamp} from "../interfaces/offRamp/IBaseOffRamp.sol";
import {ICommitStore} from "../interfaces/ICommitStore.sol";
import {IAFN} from "../interfaces/health/IAFN.sol";
import {IPool} from "../interfaces/pools/IPool.sol";
import {IEVM2EVMOffRamp} from "../interfaces/offRamp/IEVM2EVMOffRamp.sol";
import {IAny2EVMMessageReceiver} from "../interfaces/applications/IAny2EVMMessageReceiver.sol";

import {Internal} from "../models/Internal.sol";
import {Common} from "../models/Common.sol";
import {Consumer} from "../models/Consumer.sol";
import {Internal} from "../models/Internal.sol";
import {OCR2Base} from "../ocr/OCR2Base.sol";
import {Any2EVMBaseOffRamp} from "./Any2EVMBaseOffRamp.sol";

import {IERC20} from "../../vendor/IERC20.sol";
import {Address} from "../../vendor/Address.sol";
import {ERC165Checker} from "../../vendor/ERC165Checker.sol";

/**
 * @notice EVM2EVMOffRamp enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract EVM2EVMOffRamp is IEVM2EVMOffRamp, Any2EVMBaseOffRamp, TypeAndVersionInterface, OCR2Base {
  using Address for address;
  using ERC165Checker for address;

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMOffRamp 1.0.0";

  bytes32 internal immutable i_metadataHash;

  mapping(address => uint64) internal s_senderNonce;

  OffRampConfig internal s_config;

  constructor(
    uint64 sourceChainId,
    uint64 chainId,
    OffRampConfig memory offRampConfig,
    address onRampAddress,
    ICommitStore commitStore,
    IAFN afn,
    IERC20[] memory sourceTokens,
    IPool[] memory pools,
    RateLimiterConfig memory rateLimiterConfig
  )
    OCR2Base()
    Any2EVMBaseOffRamp(sourceChainId, chainId, onRampAddress, commitStore, afn, sourceTokens, pools, rateLimiterConfig)
  {
    s_config = offRampConfig;
    i_metadataHash = _metadataHash(Internal.EVM_2_EVM_MESSAGE_HASH);
  }

  /// @inheritdoc IEVM2EVMOffRamp
  function manuallyExecute(Internal.ExecutionReport memory report) external override {
    _execute(report, true);
  }

  /// @inheritdoc IEVM2EVMOffRamp
  function getSenderNonce(address sender) public view override returns (uint64 nonce) {
    return s_senderNonce[sender];
  }

  /**
   * @notice Try executing a message
   * @param message Common.Any2EVMMessage memory message
   * @param manualExecution bool to indicate manual instead of DON execution
   * @return Internal.ExecutionState
   */
  function _trialExecute(Internal.EVM2EVMMessage memory message, bool manualExecution)
    internal
    returns (Internal.MessageExecutionState)
  {
    try this.executeSingleMessage(message, manualExecution) {} catch (bytes memory err) {
      if (IBaseOffRamp.ReceiverError.selector == bytes4(err)) {
        return Internal.MessageExecutionState.FAILURE;
      } else {
        revert ExecutionError(err);
      }
    }
    return Internal.MessageExecutionState.SUCCESS;
  }

  /**
   * @notice Execute a single message
   * @param message The Any2EVMMessageFromSender message that will be executed
   * @param manualExecution bool to indicate manual instead of DON execution
   * @dev this can only be called by the contract itself. It is part of
   * the Execute call, as we can only try/catch on external calls.
   */
  function executeSingleMessage(Internal.EVM2EVMMessage memory message, bool manualExecution) external {
    if (msg.sender != address(this)) revert CanOnlySelfCall();
    Common.EVMTokenAndAmount[] memory destTokensAndAmounts = new Common.EVMTokenAndAmount[](0);
    if (message.tokensAndAmounts.length > 0) {
      destTokensAndAmounts = _releaseOrMintTokens(message.tokensAndAmounts, message.receiver);
    }
    if (
      !message.receiver.isContract() || !message.receiver.supportsInterface(type(IAny2EVMMessageReceiver).interfaceId)
    ) return;
    if (
      !s_router.routeMessage(
        Internal._toAny2EVMMessage(message, destTokensAndAmounts),
        manualExecution,
        message.gasLimit,
        message.receiver
      )
    ) revert ReceiverError();
  }

  function _executeMessages(Internal.ExecutionReport memory report, bool manualExecution) internal {
    // Report may have only price updates, so we only process messages if there are some.
    uint256 numMsgs = report.encodedMessages.length;
    if (numMsgs == 0) {
      return;
    }

    bytes32[] memory hashedLeaves = new bytes32[](numMsgs);
    Internal.EVM2EVMMessage[] memory decodedMessages = new Internal.EVM2EVMMessage[](numMsgs);

    for (uint256 i = 0; i < numMsgs; ++i) {
      Internal.EVM2EVMMessage memory decodedMessage = abi.decode(report.encodedMessages[i], (Internal.EVM2EVMMessage));
      // We do this hash here instead of in _verifyMessages to avoid two separate loops
      // over the same data, which increases gas cost
      hashedLeaves[i] = Internal._hash(decodedMessage, i_metadataHash);
      decodedMessages[i] = decodedMessage;
    }

    (uint256 timestampCommitted, ) = _verifyMessages(
      hashedLeaves,
      report.innerProofs,
      report.innerProofFlagBits,
      report.outerProofs,
      report.outerProofFlagBits
    );
    bool isOldCommitReport = (block.timestamp - timestampCommitted) > s_config.permissionLessExecutionThresholdSeconds;

    // Execute messages
    for (uint256 i = 0; i < numMsgs; ++i) {
      Internal.EVM2EVMMessage memory message = decodedMessages[i];
      Internal.MessageExecutionState originalState = getExecutionState(message.sequenceNumber);
      // Two valid cases here, we either have never touched this message before, or we tried to execute
      // and failed. This check protects against reentry and re-execution because the other states are
      // IN_PROGRESS and SUCCESS, both should not be allowed to execute.
      if (
        !(originalState == Internal.MessageExecutionState.UNTOUCHED ||
          originalState == Internal.MessageExecutionState.FAILURE)
      ) revert AlreadyExecuted(message.sequenceNumber);

      if (manualExecution) {
        // Manually execution is fine if we previously failed or if the commit report is just too old
        // Acceptable state transitions: FAILURE->SUCCESS, UNTOUCHED->SUCCESS, FAILURE->FAILURE
        if (!(isOldCommitReport || originalState == Internal.MessageExecutionState.FAILURE))
          revert ManualExecutionNotYetEnabled();
      } else {
        // DON can only execute a message once
        // Acceptable state transitions: UNTOUCHED->SUCCESS, UNTOUCHED->FAILURE
        if (originalState != Internal.MessageExecutionState.UNTOUCHED) revert AlreadyAttempted(message.sequenceNumber);
      }

      // If this is the first time executing this message we take the fee
      if (originalState == Internal.MessageExecutionState.UNTOUCHED) {
        // UNTOUCHED messages MUST be executed in order always.
        if (s_senderNonce[message.sender] + 1 != message.nonce) {
          // We skip the message if the nonce is incorrect
          emit SkippedIncorrectNonce(message.nonce, message.sender);
          continue;
        }
      }

      _isWellFormed(message);

      s_executedMessages[message.sequenceNumber] = Internal.MessageExecutionState.IN_PROGRESS;
      Internal.MessageExecutionState newState = _trialExecute(message, manualExecution);
      s_executedMessages[message.sequenceNumber] = newState;

      if (manualExecution) {
        // Nonce changes per state transition:
        // FAILURE->SUCCESS: no nonce bump unless strict
        // UNTOUCHED->SUCCESS: nonce bump
        // FAILURE->FAILURE: no nonce bump
        if (
          (message.strict &&
            originalState == Internal.MessageExecutionState.FAILURE &&
            newState == Internal.MessageExecutionState.SUCCESS) ||
          (originalState == Internal.MessageExecutionState.UNTOUCHED &&
            newState == Internal.MessageExecutionState.SUCCESS)
        ) {
          s_senderNonce[message.sender]++;
        }
      } else {
        // Nonce changes per state transition:
        // UNTOUCHED->SUCCESS: nonce bump
        // UNTOUCHED->FAILURE: nonce bump unless strict
        if (!(message.strict && newState == Internal.MessageExecutionState.FAILURE)) {
          s_senderNonce[message.sender]++;
        }
      }

      emit ExecutionStateChanged(message.sequenceNumber, message.messageId, newState);
    }
  }

  /**
   * @notice Execute a series of one or more messages using a merkle proof and update one or more
   * feeManager prices.
   * @param report ExecutionReport
   * @param manualExecution Whether the DON auto executes or it is manually initiated
   */
  function _execute(Internal.ExecutionReport memory report, bool manualExecution) internal whenNotPaused whenHealthy {
    if (address(s_router) == address(0)) revert RouterNotSet();

    // Fee updates
    if (report.feeUpdates.length != 0) {
      if (manualExecution) revert UnauthorizedGasPriceUpdate();
      s_config.feeManager.updateFees(report.feeUpdates);
    }

    // Message execution
    _executeMessages(report, manualExecution);
  }

  function _isWellFormed(Internal.EVM2EVMMessage memory message) private view {
    if (message.sourceChainId != i_sourceChainId) revert InvalidSourceChain(message.sourceChainId);
    if (message.tokensAndAmounts.length > uint256(s_config.maxTokensLength))
      revert UnsupportedNumberOfTokens(message.sequenceNumber);
    if (message.data.length > uint256(s_config.maxDataSize))
      revert MessageTooLarge(uint256(s_config.maxDataSize), message.data.length);
  }

  /// @inheritdoc IEVM2EVMOffRamp
  function getOffRampConfig() external view override returns (OffRampConfig memory) {
    return s_config;
  }

  /// @inheritdoc IEVM2EVMOffRamp
  function setOffRampConfig(OffRampConfig memory config) external override onlyOwner {
    s_config = config;

    emit OffRampConfigChanged(config);
  }

  // ******* OCR BASE ***********
  /**
   * @notice Entry point for execution, called by the OCR network
   * @dev Expects an encoded ExecutionReport
   */
  function _report(bytes memory report) internal override {
    _execute(abi.decode(report, (Internal.ExecutionReport)), false);
  }

  function _beforeSetOCR2Config(uint8 f, bytes memory onchainConfig) internal override {}
}
