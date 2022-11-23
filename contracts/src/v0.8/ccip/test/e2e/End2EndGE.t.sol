// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../commitStore/CommitStore.t.sol";
import "../onRamp/ge/EVM2EVMGEOnRampSetup.t.sol";
import "../offRamp/ge/EVM2EVMGEOffRampSetup.t.sol";

contract E2E_GE is EVM2EVMGEOnRampSetup, CommitStoreSetup, EVM2EVMGEOffRampSetup {
  using CCIP for CCIP.EVM2EVMGEMessage;

  Any2EVMOffRampRouterInterface public s_router;

  MerkleHelper public s_merkleHelper;

  function setUp() public virtual override(EVM2EVMGEOnRampSetup, CommitStoreSetup, EVM2EVMGEOffRampSetup) {
    EVM2EVMGEOnRampSetup.setUp();
    CommitStoreSetup.setUp();
    EVM2EVMGEOffRampSetup.setUp();

    deployOffRamp(s_commitStore, s_gasFeeCache);

    s_merkleHelper = new MerkleHelper();

    BaseOffRampInterface[] memory offRamps = new BaseOffRampInterface[](1);
    offRamps[0] = s_offRamp;
    s_router = new GERouter(offRamps);
    s_offRamp.setRouter(s_router);
  }

  function testSuccess() public {
    IERC20 token0 = IERC20(s_sourceTokens[0]);
    IERC20 token1 = IERC20(s_sourceTokens[1]);
    uint256 balance0Pre = token0.balanceOf(OWNER);
    uint256 balance1Pre = token1.balanceOf(OWNER);

    CCIP.EVM2EVMGEMessage[] memory messages = new CCIP.EVM2EVMGEMessage[](3);
    messages[0] = sendRequest(1);
    messages[1] = sendRequest(2);
    messages[2] = sendRequest(3);

    uint256 expectedFee = s_onRampRouter.getFee(DEST_CHAIN_ID, _generateTokenMessage());
    // Asserts that the tokens have been sent and the fee has been paid.
    assertEq(balance0Pre - messages.length * (i_tokenAmount0 + expectedFee), token0.balanceOf(OWNER));
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
    onRamps[0] = commitStoreConfig().onRamps[0];

    CCIP.CommitReport memory report = CCIP.CommitReport({
      onRamps: onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: merkleRoots[0]
    });

    s_commitStore.report(abi.encode(report));
    bytes32[] memory proofs = new bytes32[](0);
    uint256 timestamp = s_commitStore.verify(merkleRoots, proofs, 2**2 - 1, proofs, 2**2 - 1);
    assertEq(BLOCK_TIME, timestamp);

    // We change the block time so when execute would e.g. use the current
    // block time instead of the committed block time the value would be
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

  function sendRequest(uint64 expectedSeqNum) public returns (CCIP.EVM2EVMGEMessage memory) {
    CCIP.EVM2AnyGEMessage memory message = _generateTokenMessage();
    uint256 expectedFee = s_onRampRouter.getFee(DEST_CHAIN_ID, message);

    IERC20(s_sourceTokens[0]).approve(address(s_onRampRouter), i_tokenAmount0 + expectedFee);
    IERC20(s_sourceTokens[1]).approve(address(s_onRampRouter), i_tokenAmount1);

    message.receiver = abi.encode(address(s_receiver));
    CCIP.EVM2EVMGEMessage memory geEvent = _messageToEvent(message, expectedSeqNum, expectedSeqNum, expectedFee);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(geEvent);

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);

    return geEvent;
  }
}
