// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "./Any2EVMSubscriptionOffRampRouter.sol";
import "../../ocr/OCR2Base.sol";
import "../BaseOffRamp.sol";

/**
 * @notice EVM2EVMSubscriptionOffRamp enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract EVM2EVMSubscriptionOffRamp is BaseOffRamp, TypeAndVersionInterface, OCR2Base {
  using CCIP for CCIP.EVM2EVMSubscriptionMessage;

  string public constant override typeAndVersion = "EVM2EVMSubscriptionOffRamp 1.0.0";

  mapping(address => uint64) internal s_receiverToNonce;

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
    OCR2Base(true)
    BaseOffRamp(
      sourceChainId,
      chainId,
      offRampConfig,
      blobVerifier,
      afn,
      sourceTokens,
      pools,
      rateLimiterConfig,
      tokenLimitsAdmin
    )
  {}

  function execute(CCIP.ExecutionReport memory report, bool manualExecution)
    external
    override
    whenNotPaused
    whenHealthy
  {
    address routerAddress = address(s_router);
    if (routerAddress == address(0)) revert RouterNotSet();
    uint256 numMsgs = report.encodedMessages.length;
    if (numMsgs == 0) revert NoMessagesToExecute();

    CCIP.EVM2EVMSubscriptionMessage[] memory decodedMessages = new CCIP.EVM2EVMSubscriptionMessage[](numMsgs);
    bytes32[] memory hashedLeaves = new bytes32[](numMsgs);
    // TODO optimise gas cost of hashing/caching hash
    bytes32 metadataHash = _metadataHash(CCIP.EVM_2_EVM_SUBSCRIPTION_MESSAGE_HASH);
    for (uint256 i = 0; i < numMsgs; ++i) {
      CCIP.EVM2EVMSubscriptionMessage memory decodedMessage = abi.decode(
        report.encodedMessages[i],
        (CCIP.EVM2EVMSubscriptionMessage)
      );
      // We do this hash here instead of in _verifyMessages to avoid two separate loops
      // over the same data, which increases gas cost
      // TODO: golf check
      hashedLeaves[i] = decodedMessage._hash(metadataHash);
      decodedMessages[i] = decodedMessage;
    }

    (uint256 timestampRelayed, uint256 gasUsedByMerkle) = _verifyMessages(
      hashedLeaves,
      report.innerProofs,
      report.innerProofFlagBits,
      report.outerProofs,
      report.outerProofFlagBits
    );
    uint256 merkleGasShare = gasUsedByMerkle / decodedMessages.length;

    // only allow manual execution if the report is old enough
    if (manualExecution && (block.timestamp - timestampRelayed) < s_config.permissionLessExecutionThresholdSeconds) {
      revert ManualExecutionNotYetEnabled();
    }

    // tokenPerFeeCoin[0] is used because all subscriptions use the same payment token
    uint256 tokenPerFeeCoin = report.tokenPerFeeCoin[0];
    for (uint256 i = 0; i < numMsgs; ++i) {
      uint256 gasBegin = gasleft();
      CCIP.EVM2EVMSubscriptionMessage memory message = decodedMessages[i];
      CCIP.MessageExecutionState state = getExecutionState(message.sequenceNumber);
      if (state == CCIP.MessageExecutionState.SUCCESS) revert AlreadyExecuted(message.sequenceNumber);
      SubscriptionInterface.OffRampSubscription memory subscription = Subscription(routerAddress).getSubscription(
        message.receiver
      );
      if (address(subscription.receiver) == address(0)) {
        revert SubscriptionInterface.SubscriptionNotFound(message.receiver);
      }

      // Reduce stack pressure
      {
        // Any message with a nonce that is n + 1 is allowed.
        // If strict sequencing is disabled then any failed message can be re-executed out-of-order.
        bool isNextInSequence = s_receiverToNonce[message.receiver] + 1 == message.nonce;
        if (!(isNextInSequence || (!subscription.strictSequencing && state == CCIP.MessageExecutionState.FAILURE))) {
          revert IncorrectNonce(message.nonce);
        }

        _isWellFormed(message);

        s_executedMessages[message.sequenceNumber] = CCIP.MessageExecutionState.IN_PROGRESS;
        CCIP.MessageExecutionState newState = _trialExecute(_toAny2EVMMessageFromSender(message));
        s_executedMessages[message.sequenceNumber] = newState;
        emit ExecutionStateChanged(message.sequenceNumber, newState);

        // Increment the nonce of the receiver if it's the next nonce in line and it was successfully
        // executed or if the subscription doesn't require strict sequencing.
        if (isNextInSequence && (newState == CCIP.MessageExecutionState.SUCCESS || !subscription.strictSequencing)) {
          s_receiverToNonce[message.receiver]++;
        }
      }

      if (!manualExecution) {
        Any2EVMSubscriptionOffRampRouter(routerAddress).chargeSubscription(
          message.receiver,
          message.sender,
          // Gas cost in wei: gasUsed * gasPrice
          // example: 100k gas, 20 gwei = 1e5 * 20e9  = 2e15
          // Gas cost in token: costInWei * 1e18 / tokenPerFeeCoin
          // example: costInWei 2e15, tokenPerFeeCoin 2e20 = 2e15 * 2e20 / 1e18 = 4e17 tokens
          ((gasBegin - gasleft() + merkleGasShare) * tx.gasprice * tokenPerFeeCoin) / 1 ether
        );
      }
    }
  }

  function _toAny2EVMMessageFromSender(CCIP.EVM2EVMSubscriptionMessage memory original)
    internal
    view
    returns (CCIP.Any2EVMMessageFromSender memory message)
  {
    uint256 numberOfTokens = original.tokens.length;
    IERC20[] memory destTokens = new IERC20[](numberOfTokens);
    PoolInterface[] memory destPools = new PoolInterface[](numberOfTokens);

    for (uint256 i = 0; i < numberOfTokens; ++i) {
      PoolInterface pool = _getPool(original.tokens[i]);
      destPools[i] = pool;
      destTokens[i] = pool.getToken();
    }

    message = CCIP.Any2EVMMessageFromSender({
      sourceChainId: original.sourceChainId,
      sender: abi.encode(original.sender),
      receiver: original.receiver,
      data: original.data,
      destTokens: destTokens,
      destPools: destPools,
      amounts: original.amounts,
      gasLimit: original.gasLimit
    });
  }

  function getNonce(address receiver) external view returns (uint64) {
    return s_receiverToNonce[receiver];
  }

  function _isWellFormed(CCIP.EVM2EVMSubscriptionMessage memory message) private view {
    if (message.sourceChainId != i_sourceChainId) revert InvalidSourceChain(message.sourceChainId);
    if (message.tokens.length > uint256(s_config.maxTokensLength) || message.tokens.length != message.amounts.length) {
      revert UnsupportedNumberOfTokens(message.sequenceNumber);
    }
    if (message.data.length > uint256(s_config.maxDataSize))
      revert MessageTooLarge(uint256(s_config.maxDataSize), message.data.length);
  }

  // ******* OCR BASE ***********
  /**
   * @notice Entry point for execution, called by the OCR network
   * @dev Expects an encoded ExecutionReport
   */
  function _report(
    bytes32, /*configDigest*/
    uint40, /*epochAndRound*/
    bytes memory report
  ) internal override {
    this.execute(abi.decode(report, (CCIP.ExecutionReport)), false);
  }

  function _beforeSetConfig(uint8 _threshold, bytes memory _onchainConfig) internal override {}

  function _afterSetConfig(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {}

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {}
}
