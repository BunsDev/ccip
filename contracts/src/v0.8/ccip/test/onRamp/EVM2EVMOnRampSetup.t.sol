// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IFeeManager} from "../../interfaces/fees/IFeeManager.sol";
import {IEVM2EVMOnRamp} from "../../interfaces/onRamp/IEVM2EVMOnRamp.sol";

import {EVM2EVMOnRamp} from "../../onRamp/EVM2EVMOnRamp.sol";
import {FeeManager} from "../../fees/FeeManager.sol";
import {Router} from "../../router/Router.sol";
import {RouterSetup} from "../router/RouterSetup.t.sol";
import {Internal} from "../../models/Internal.sol";
import {Consumer} from "../../models/Consumer.sol";
import "../../offRamp/EVM2EVMOffRamp.sol";
import "../TokenSetup.t.sol";

contract EVM2EVMOnRampSetup is TokenSetup, RouterSetup {
  // Duplicate event of the CCIPSendRequested in the IOnRamp
  event CCIPSendRequested(Internal.EVM2EVMMessage message);

  uint256 internal immutable i_tokenAmount0 = 9;
  uint256 internal immutable i_tokenAmount1 = 7;

  bytes32 internal s_metadataHash;

  address[] internal s_allowList;

  EVM2EVMOnRamp internal s_onRamp;
  address[] s_offRamps;
  // Naming chosen to not collide with s_feeManager in the offRampSetup since both
  // are imported into the e2e test.
  IFeeManager internal s_IFeeManager;

  function setUp() public virtual override(TokenSetup, RouterSetup) {
    TokenSetup.setUp();
    RouterSetup.setUp();

    Internal.FeeUpdate[] memory feeUpdates = new Internal.FeeUpdate[](2);
    feeUpdates[0] = Internal.FeeUpdate({
      sourceFeeToken: s_sourceTokens[0],
      destChainId: DEST_CHAIN_ID,
      feeTokenBaseUnitsPerUnitGas: 100
    });
    feeUpdates[1] = Internal.FeeUpdate({
      sourceFeeToken: s_sourceRouter.getWrappedNative(),
      destChainId: DEST_CHAIN_ID,
      feeTokenBaseUnitsPerUnitGas: 101
    });

    IEVM2EVMOnRamp.FeeTokenConfigArgs[] memory feeTokenConfigArgs = new IEVM2EVMOnRamp.FeeTokenConfigArgs[](2);
    feeTokenConfigArgs[0] = IEVM2EVMOnRamp.FeeTokenConfigArgs({
      token: s_sourceTokens[0],
      feeAmount: 1,
      multiplier: 108e16,
      destGasOverhead: 1
    });
    feeTokenConfigArgs[1] = IEVM2EVMOnRamp.FeeTokenConfigArgs({
      token: s_sourceRouter.getWrappedNative(),
      feeAmount: 2,
      multiplier: 108e16,
      destGasOverhead: 2
    });
    address[] memory feeUpdaters = new address[](0);
    s_IFeeManager = new FeeManager(feeUpdates, feeUpdaters, TWELVE_HOURS);

    s_onRamp = new EVM2EVMOnRamp(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      s_sourceTokens,
      getCastedSourcePools(),
      s_allowList,
      s_afn,
      onRampConfig(),
      rateLimiterConfig(),
      address(s_sourceRouter),
      address(s_IFeeManager),
      feeTokenConfigArgs
    );

    s_metadataHash = keccak256(
      abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_ID, DEST_CHAIN_ID, address(s_onRamp))
    );

    s_onRamp.setPrices(getCastedSourceTokens(), getTokenPrices());

    LockReleaseTokenPool(address(s_sourcePools[0])).setOnRamp(address(s_onRamp), true);
    LockReleaseTokenPool(address(s_sourcePools[1])).setOnRamp(address(s_onRamp), true);

    s_sourceRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);

    s_offRamps = new address[](2);
    s_offRamps[0] = address(10);
    s_offRamps[1] = address(11);
    s_sourceRouter.addOffRamp(s_offRamps[0]);
    s_sourceRouter.addOffRamp(s_offRamps[1]);

    // Pre approve the first token so the gas estimates of the tests
    // only cover actual gas usage from the ramps
    IERC20(s_sourceTokens[0]).approve(address(s_sourceRouter), 2**128);
  }

  function assertSameConfig(IEVM2EVMOnRamp.OnRampConfig memory a, IEVM2EVMOnRamp.OnRampConfig memory b) public {
    assertEq(a.maxDataSize, b.maxDataSize);
    assertEq(a.maxTokensLength, b.maxTokensLength);
    assertEq(a.maxGasLimit, b.maxGasLimit);
  }

  function _generateTokenMessage() public view returns (Consumer.EVM2AnyMessage memory) {
    Common.EVMTokenAndAmount[] memory tokensAndAmounts = getCastedSourceEVMTokenAndAmountsWithZeroAmounts();
    tokensAndAmounts[0].amount = i_tokenAmount0;
    tokensAndAmounts[1].amount = i_tokenAmount1;
    return
      Consumer.EVM2AnyMessage({
        receiver: abi.encode(OWNER),
        data: "",
        tokensAndAmounts: tokensAndAmounts,
        feeToken: s_sourceFeeToken,
        extraArgs: Consumer._argsToBytes(Consumer.EVMExtraArgsV1({gasLimit: GAS_LIMIT, strict: false}))
      });
  }

  function _generateEmptyMessage() public view returns (Consumer.EVM2AnyMessage memory) {
    return
      Consumer.EVM2AnyMessage({
        receiver: abi.encode(OWNER),
        data: "",
        tokensAndAmounts: new Common.EVMTokenAndAmount[](0),
        feeToken: s_sourceFeeToken,
        extraArgs: Consumer._argsToBytes(Consumer.EVMExtraArgsV1({gasLimit: GAS_LIMIT, strict: false}))
      });
  }

  function _messageToEvent(
    Consumer.EVM2AnyMessage memory message,
    uint64 seqNum,
    uint64 nonce,
    uint256 feeTokenAmount
  ) public view returns (Internal.EVM2EVMMessage memory) {
    Internal.EVM2EVMMessage memory messageEvent = Internal.EVM2EVMMessage({
      sequenceNumber: seqNum,
      feeTokenAmount: feeTokenAmount,
      sender: OWNER,
      nonce: nonce,
      gasLimit: GAS_LIMIT,
      strict: false,
      sourceChainId: SOURCE_CHAIN_ID,
      receiver: abi.decode(message.receiver, (address)),
      data: message.data,
      tokensAndAmounts: message.tokensAndAmounts,
      feeToken: message.feeToken,
      messageId: ""
    });

    messageEvent.messageId = Internal._hash(messageEvent, s_metadataHash);
    return messageEvent;
  }
}
