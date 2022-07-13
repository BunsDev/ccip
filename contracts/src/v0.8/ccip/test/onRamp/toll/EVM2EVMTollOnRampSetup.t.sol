// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../TokenSetup.t.sol";
import "../../../utils/CCIP.sol";
import "../../../onRamp/toll/EVM2EVMTollOnRamp.sol";
import "../../../onRamp/toll/EVM2AnyTollOnRampRouter.sol";

contract EVM2EVMTollOnRampSetup is TokenSetup {
  // Duplicate event of the CCIPSendRequested in the TollOnRampInterface
  event CCIPSendRequested(CCIP.EVM2EVMTollEvent message);

  uint256 immutable FEE_AMOUNT = 1;
  uint256 immutable TOKEN_AMOUNT_0 = 9;
  uint256 immutable TOKEN_AMOUNT_1 = 7;

  address[] internal s_allowList;

  EVM2AnyTollOnRampRouter internal s_onRampRouter;
  EVM2EVMTollOnRamp internal s_onRamp;
  BaseOnRampInterface.OnRampConfig internal s_onRampConfig;

  function setUp() public virtual override {
    TokenSetup.setUp();

    s_onRampRouter = new EVM2AnyTollOnRampRouter();

    s_onRampConfig = BaseOnRampInterface.OnRampConfig({
      relayingFeeJuels: uint64(FEE_AMOUNT),
      maxDataSize: 50,
      maxTokensLength: 3
    });

    s_onRamp = new EVM2EVMTollOnRamp(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      s_sourceTokens,
      s_sourcePools,
      s_sourceFeeds,
      s_allowList,
      s_afn,
      HEARTBEAT,
      s_onRampConfig,
      s_onRampRouter
    );

    NativeTokenPool(address(s_sourcePools[0])).setOnRamp(s_onRamp, true);
    NativeTokenPool(address(s_sourcePools[1])).setOnRamp(s_onRamp, true);

    s_onRampRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);

    // Pre approve the first token so the gas estimates of the tests
    // only cover actual gas usage from the ramps
    s_sourceTokens[0].approve(address(s_onRampRouter), 2**128);
  }

  function getTokenMessage() public view returns (CCIP.EVM2AnyTollMessage memory) {
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = TOKEN_AMOUNT_0;
    amounts[1] = TOKEN_AMOUNT_1;
    IERC20[] memory tokens = s_sourceTokens;
    return
      CCIP.EVM2AnyTollMessage({
        receiver: OWNER,
        data: "",
        tokens: tokens,
        amounts: amounts,
        feeToken: s_sourceTokens[0],
        feeTokenAmount: 1,
        gasLimit: 0
      });
  }

  function getEmptyMessage() public view returns (CCIP.EVM2AnyTollMessage memory) {
    uint256[] memory amounts = new uint256[](0);
    IERC20[] memory tokens = new IERC20[](0);
    return
      CCIP.EVM2AnyTollMessage({
        receiver: OWNER,
        data: "",
        tokens: tokens,
        amounts: amounts,
        feeToken: s_sourceTokens[0],
        feeTokenAmount: 1,
        gasLimit: 0
      });
  }

  function messageToEvent(CCIP.EVM2AnyTollMessage memory message, uint64 seqNum)
    public
    pure
    returns (CCIP.EVM2EVMTollEvent memory)
  {
    return
      CCIP.EVM2EVMTollEvent({
        sequenceNumber: seqNum,
        sourceChainId: SOURCE_CHAIN_ID,
        sender: OWNER,
        receiver: message.receiver,
        data: message.data,
        tokens: message.tokens,
        amounts: message.amounts,
        feeToken: message.feeToken,
        feeTokenAmount: message.feeTokenAmount - FEE_AMOUNT,
        gasLimit: message.gasLimit
      });
  }
}
