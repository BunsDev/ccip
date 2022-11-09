// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../offRamp/toll/Any2EVMTollOffRampRouter.sol";
import "../offRamp/toll/EVM2EVMTollOffRampSetup.t.sol";
import "../onRamp/toll/EVM2EVMTollOnRampSetup.t.sol";
import "../blobVerifier/BlobVerifier.t.sol";

contract E2E_toll is EVM2EVMTollOnRampSetup, BlobVerifierSetup, EVM2EVMTollOffRampSetup {
  using CCIP for CCIP.EVM2EVMTollMessage;

  Any2EVMOffRampRouterInterface public s_router;

  MerkleHelper public s_merkleHelper;

  function setUp() public virtual override(EVM2EVMTollOnRampSetup, BlobVerifierSetup, EVM2EVMTollOffRampSetup) {
    EVM2EVMTollOnRampSetup.setUp();
    BlobVerifierSetup.setUp();
    EVM2EVMTollOffRampSetup.setUp();

    deployOffRamp(s_blobVerifier);

    s_merkleHelper = new MerkleHelper();

    BaseOffRampInterface[] memory offRamps = new BaseOffRampInterface[](1);
    offRamps[0] = s_offRamp;
    s_router = new Any2EVMTollOffRampRouter(offRamps);
    s_offRamp.setRouter(s_router);
  }

  function testSuccess() public {
    IERC20 token0 = IERC20(s_sourceTokens[0]);
    IERC20 token1 = IERC20(s_sourceTokens[1]);
    uint256 balance0Pre = token0.balanceOf(OWNER);
    uint256 balance1Pre = token1.balanceOf(OWNER);

    CCIP.EVM2EVMTollMessage[] memory messages = new CCIP.EVM2EVMTollMessage[](3);
    messages[0] = parseEventToDestChainMessage(sendRequest(1));
    messages[1] = parseEventToDestChainMessage(sendRequest(2));
    messages[2] = parseEventToDestChainMessage(sendRequest(3));

    // Asserts that the tokens have been sent and the fee has been paid.
    assertEq(
      balance0Pre - messages.length * (i_tokenAmount0 + RELAYING_FEE_JUELS + EXECUTION_FEE_AMOUNT),
      token0.balanceOf(OWNER)
    );
    assertEq(balance1Pre - messages.length * i_tokenAmount1, token1.balanceOf(OWNER));

    bytes32 metaDataHash = s_offRamp.metadataHash();

    bytes32[] memory hashedMessages = new bytes32[](3);
    hashedMessages[0] = messages[0]._hash(metaDataHash);
    hashedMessages[1] = messages[1]._hash(metaDataHash);
    hashedMessages[2] = messages[2]._hash(metaDataHash);

    CCIP.Interval[] memory intervals = new CCIP.Interval[](1);
    intervals[0] = CCIP.Interval(messages[0].sequenceNumber, messages[2].sequenceNumber);

    bytes32[] memory merkleRoots = new bytes32[](1);
    merkleRoots[0] = s_merkleHelper.getMerkleRoot(hashedMessages);

    address[] memory onRamps = new address[](1);
    onRamps[0] = blobVerifierConfig().onRamps[0];

    CCIP.RelayReport memory report = CCIP.RelayReport({
      onRamps: onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: merkleRoots[0]
    });

    s_blobVerifier.report(abi.encode(report));
    bytes32[] memory proofs = new bytes32[](0);
    uint256 timestamp = s_blobVerifier.verify(merkleRoots, proofs, 2**2 - 1, proofs, 2**2 - 1);
    assertEq(BLOCK_TIME, timestamp);

    // We change the block time so when execute would e.g. use the current
    // block time instead of the relayed block time the value would be
    // incorrect in the checks below.
    vm.warp(BLOCK_TIME + 2000);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[0].sequenceNumber, CCIP.MessageExecutionState.SUCCESS);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[1].sequenceNumber, CCIP.MessageExecutionState.SUCCESS);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[2].sequenceNumber, CCIP.MessageExecutionState.SUCCESS);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function sendRequest(uint64 expectedSeqNum) public returns (CCIP.EVM2EVMTollMessage memory) {
    CCIP.EVM2AnyTollMessage memory message = _generateTokenMessage();
    message.feeTokenAndAmount.amount = RELAYING_FEE_JUELS + EXECUTION_FEE_AMOUNT;

    IERC20(s_sourceTokens[0]).approve(
      address(s_onRampRouter),
      i_tokenAmount0 + RELAYING_FEE_JUELS + EXECUTION_FEE_AMOUNT
    );
    IERC20(s_sourceTokens[1]).approve(address(s_onRampRouter), i_tokenAmount1);

    message.receiver = abi.encode(address(s_receiver));
    CCIP.EVM2EVMTollMessage memory tollEvent = _messageToEvent(message, expectedSeqNum);
    vm.expectEmit(false, false, false, true);

    emit CCIPSendRequested(tollEvent);

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);

    return tollEvent;
  }

  function parseEventToDestChainMessage(CCIP.EVM2EVMTollMessage memory sendEvent)
    public
    pure
    returns (CCIP.EVM2EVMTollMessage memory)
  {
    return
      CCIP.EVM2EVMTollMessage({
        sourceChainId: sendEvent.sourceChainId,
        sequenceNumber: sendEvent.sequenceNumber,
        sender: sendEvent.sender,
        receiver: sendEvent.receiver,
        data: sendEvent.data,
        tokensAndAmounts: sendEvent.tokensAndAmounts,
        feeTokenAndAmount: sendEvent.feeTokenAndAmount,
        gasLimit: sendEvent.gasLimit
      });
  }
}
