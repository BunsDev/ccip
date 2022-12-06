// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../applications/PingPongDemo.sol";
import "../../onRamp/ge/EVM2EVMGEOnRampSetup.t.sol";
import "../../../models/Common.sol";

// setup
contract PingPongDappSetup is EVM2EVMGEOnRampSetup {
  event Ping(uint256 pingPongs);
  event Pong(uint256 pingPongs);

  PingPongDemo s_pingPong;
  IERC20 s_feeToken;

  address immutable i_pongContract = address(10);

  function setUp() public virtual override {
    EVM2EVMGEOnRampSetup.setUp();

    s_feeToken = IERC20(s_sourceTokens[0]);
    s_pingPong = new PingPongDemo(GERouterInterface(address(s_sourceRouter)), s_sourceFeeToken);
    s_pingPong.setCounterpart(DEST_CHAIN_ID, i_pongContract);

    uint256 fundingAmount = 1e18;

    // Fund the contract with LINK tokens
    IERC20(s_sourceFeeToken).transfer(address(s_pingPong), fundingAmount);

    // Approve the link tokens from the dapp
    changePrank(address(s_pingPong));
    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), fundingAmount);

    // Change back to te deployer of the contracts
    changePrank(OWNER);
  }
}

/// @notice #startPingPong
contract PingPong_startPingPong is PingPongDappSetup {
  event ConfigPropagated(uint256 chainId, address contractAddress);

  // Success
  function testSuccess() public {
    uint256 pingPongNumber = 1;
    bytes memory data = abi.encode(pingPongNumber);

    GEConsumer.EVM2AnyGEMessage memory sentMessage = GEConsumer.EVM2AnyGEMessage({
      receiver: abi.encode(i_pongContract),
      data: data,
      tokensAndAmounts: new Common.EVMTokenAndAmount[](0),
      feeToken: s_sourceFeeToken,
      extraArgs: GEConsumer._argsToBytes(GEConsumer.EVMExtraArgsV1({gasLimit: 2e5, strict: false}))
    });

    uint256 expectedFee = s_sourceRouter.getFee(DEST_CHAIN_ID, sentMessage);

    GE.EVM2EVMGEMessage memory message = GE.EVM2EVMGEMessage({
      sequenceNumber: 1,
      feeTokenAmount: expectedFee,
      sourceChainId: SOURCE_CHAIN_ID,
      sender: address(s_pingPong),
      receiver: i_pongContract,
      nonce: 1,
      data: data,
      tokensAndAmounts: sentMessage.tokensAndAmounts,
      gasLimit: 2e5,
      feeToken: sentMessage.feeToken,
      strict: false,
      messageId: ""
    });
    message.messageId = GE._hash(message, s_metadataHash);

    vm.expectEmit(false, false, false, true);
    emit Ping(pingPongNumber);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(message);

    s_pingPong.startPingPong();
  }
}

/// @notice #ccipReceive
contract PingPong_ccipReceive is PingPongDappSetup {
  // Success

  function testSuccess() public {
    Common.EVMTokenAndAmount[] memory tokensAndAmounts = new Common.EVMTokenAndAmount[](0);

    uint256 pingPongNumber = 5;

    Common.Any2EVMMessage memory message = Common.Any2EVMMessage({
      sourceChainId: DEST_CHAIN_ID,
      sender: abi.encode(i_pongContract),
      data: abi.encode(pingPongNumber),
      destTokensAndAmounts: tokensAndAmounts
    });

    changePrank(address(s_sourceRouter));

    vm.expectEmit(false, false, false, true);
    emit Pong(pingPongNumber + 1);

    s_pingPong.ccipReceive(message);
  }
  // Revert
}
