// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {Address} from "../../vendor/Address.sol";
import {HealthChecker, AFNInterface} from "../health/HealthChecker.sol";
import {TokenPoolRegistry} from "../pools/TokenPoolRegistry.sol";
import {AggregateRateLimiter} from "../rateLimiter/AggregateRateLimiter.sol";
import {BaseOffRampInterface, Any2EVMOffRampRouterInterface, BlobVerifierInterface} from "../interfaces/offRamp/BaseOffRampInterface.sol";
import {CCIP, IERC20, PoolInterface} from "../models/Models.sol";

/**
 * @notice A base OffRamp contract that every OffRamp should expand on
 */
contract BaseOffRamp is BaseOffRampInterface, HealthChecker, TokenPoolRegistry, AggregateRateLimiter {
  using Address for address;

  // Chain ID of the source chain
  uint256 internal immutable i_sourceChainId;
  // Chain ID of this chain
  uint256 internal immutable i_chainId;

  // The router through which all transactions will be executed
  Any2EVMOffRampRouterInterface internal s_router;

  // The blob verifier contract
  BlobVerifierInterface internal s_blobVerifier;

  // The on chain offRamp configuration values
  OffRampConfig internal s_config;

  uint256 internal constant EXTERNAL_CALL_OVERHEAD_GAS = 2600;
  uint256 internal constant RATE_LIMITER_OVERHEAD_GAS = (2_100 + 5_000); // COLD_SLOAD_COST for accessing token bucket // SSTORE_RESET_GAS for updating & decreasing token bucket
  uint256 internal constant EVM_ADDRESS_LENGTH_BYTES = 20;
  uint256 internal constant EVM_WORD_BYTES = 32;
  uint256 internal constant CALLDATA_GAS_PER_BYTE = 16;
  uint256 internal constant PER_TOKEN_OVERHEAD_GAS = (2_100 + // COLD_SLOAD_COST for first reading the pool
    2_100 + // COLD_SLOAD_COST for pool to ensure allowed offramp calls it
    2_100 + // COLD_SLOAD_COST for accessing pool balance slot
    5_000 + // SSTORE_RESET_GAS for decreasing pool balance from non-zero to non-zero
    2_100 + // COLD_SLOAD_COST for accessing receiver balance
    20_000 + // SSTORE_SET_GAS for increasing receiver balance from zero to non-zero
    2_100); // COLD_SLOAD_COST for obtanining price of token to use for aggregate token bucket

  // A mapping of sequence numbers to execution state.
  // This makes sure we never execute a message twice.
  mapping(uint64 => CCIP.MessageExecutionState) internal s_executedMessages;

  constructor(
    uint256 sourceChainId,
    uint256 chainId,
    OffRampConfig memory offRampConfig,
    BlobVerifierInterface blobVerifier,
    AFNInterface afn,
    IERC20[] memory sourceTokens,
    PoolInterface[] memory pools,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin
  )
    HealthChecker(afn)
    TokenPoolRegistry(sourceTokens, pools)
    AggregateRateLimiter(rateLimiterConfig, tokenLimitsAdmin)
  {
    if (offRampConfig.onRampAddress == address(0)) revert ZeroAddressNotAllowed();
    // TokenPoolRegistry does a check on tokens.length != pools.length
    i_sourceChainId = sourceChainId;
    i_chainId = chainId;
    s_config = offRampConfig;
    s_blobVerifier = blobVerifier;
  }

  /**
   * @notice Uses the pool to release or mint tokens and send them to
   *          the given `receiver` address.
   */
  function _releaseOrMintToken(
    PoolInterface pool,
    uint256 amount,
    address receiver
  ) internal {
    pool.releaseOrMint(receiver, amount);
  }

  /**
   * @notice Uses pools to release or mint a number of different tokens
   *           and send them to the given `receiver` address.
   */
  function _releaseOrMintTokens(
    PoolInterface[] memory pools,
    uint256[] memory amounts,
    address receiver
  ) internal {
    if (pools.length != amounts.length) revert TokenAndAmountMisMatch();
    for (uint256 i = 0; i < pools.length; ++i) {
      _releaseOrMintToken(pools[i], amounts[i], receiver);
    }
  }

  /**
   * @notice Verifies that the given hashed messages are valid leaves of
   *          a relayed merkle tree.
   */
  function _verifyMessages(
    bytes32[] memory hashedLeaves,
    bytes32[] memory innerProofs,
    uint256 innerProofFlagBits,
    bytes32[] memory outerProofs,
    uint256 outerProofFlagBits
  ) internal returns (uint256, uint256) {
    uint256 gasBegin = gasleft();
    uint256 timestampRelayed = s_blobVerifier.verify(
      hashedLeaves,
      innerProofs,
      innerProofFlagBits,
      outerProofs,
      outerProofFlagBits
    );
    if (timestampRelayed <= 0) revert RootNotRelayed();
    return (timestampRelayed, gasBegin - gasleft());
  }

  /**
   * @notice Try executing a message
   * @param message CCIP.Any2EVMMessageFromSender memory message
   * @return CCIP.ExecutionState
   */
  function _trialExecute(CCIP.Any2EVMMessageFromSender memory message) internal returns (CCIP.MessageExecutionState) {
    try this.executeSingleMessage(message) {} catch (bytes memory err) {
      if (BaseOffRampInterface.ReceiverError.selector == bytes4(err)) {
        return CCIP.MessageExecutionState.FAILURE;
      } else {
        revert ExecutionError(err);
      }
    }
    return CCIP.MessageExecutionState.SUCCESS;
  }

  /**
   * @notice Execute a single message
   * @param message The Any2EVMMessageFromSender message that will be executed
   * @dev this can only be called by the contract itself. It is part of
   * the Execute call, as we can only try/catch on external calls.
   */
  function executeSingleMessage(CCIP.Any2EVMMessageFromSender memory message) external {
    if (msg.sender != address(this)) revert CanOnlySelfCall();
    if (message.destTokens.length > 0) {
      _removeTokens(message.destTokens, message.amounts);
      _releaseOrMintTokens(message.destPools, message.amounts, message.receiver);
    }

    _callReceiver(message);
  }

  function _callReceiver(CCIP.Any2EVMMessageFromSender memory message) internal {
    if (!message.receiver.isContract()) return;
    if (!s_router.routeMessage(message)) revert ReceiverError();
  }

  /**
   * @notice Reverts as this contract should not access CCIP messages
   */
  function ccipReceive(CCIP.Any2EVMMessageFromSender calldata) external pure {
    // solhint-disable-next-line reason-string
    revert();
  }

  /// @inheritdoc BaseOffRampInterface
  function execute(CCIP.ExecutionReport memory, bool) external virtual override {
    // solhint-disable-next-line reason-string
    revert();
  }

  /// @inheritdoc BaseOffRampInterface
  function setRouter(Any2EVMOffRampRouterInterface router) external onlyOwner {
    s_router = router;
    emit OffRampRouterSet(address(router));
  }

  /// @inheritdoc BaseOffRampInterface
  function getRouter() external view override returns (Any2EVMOffRampRouterInterface) {
    return s_router;
  }

  /// @inheritdoc BaseOffRampInterface
  function getExecutionState(uint64 sequenceNumber) public view returns (CCIP.MessageExecutionState) {
    return s_executedMessages[sequenceNumber];
  }

  /// @inheritdoc BaseOffRampInterface
  function getBlobVerifier() external view returns (BlobVerifierInterface) {
    return s_blobVerifier;
  }

  /// @inheritdoc BaseOffRampInterface
  function setBlobVerifier(BlobVerifierInterface blobVerifier) external onlyOwner {
    s_blobVerifier = blobVerifier;
  }

  /// @inheritdoc BaseOffRampInterface
  function getConfig() external view returns (OffRampConfig memory) {
    return s_config;
  }

  /// @inheritdoc BaseOffRampInterface
  function setConfig(OffRampConfig memory config) external onlyOwner {
    if (config.onRampAddress == address(0)) revert ZeroAddressNotAllowed();
    s_config = config;

    emit OffRampConfigSet(config);
  }

  function getChainIDs() external view returns (uint256 sourceChainId, uint256 chainId) {
    sourceChainId = i_sourceChainId;
    chainId = i_chainId;
  }

  /**
   * @notice Returns the pool for a given source chain token.
   */
  function _getPool(IERC20 token) internal view returns (PoolInterface pool) {
    pool = getPool(token);
    if (address(pool) == address(0)) revert UnsupportedToken(token);
  }

  function _metadataHash(bytes32 prefix) internal view returns (bytes32) {
    return keccak256(abi.encode(prefix, i_sourceChainId, i_chainId, s_config.onRampAddress));
  }
}
