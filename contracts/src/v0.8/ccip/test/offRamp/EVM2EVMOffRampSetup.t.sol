// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {ICommitStore} from "../../interfaces/ICommitStore.sol";
import {IAny2EVMMessageReceiver} from "../../interfaces/router/IAny2EVMMessageReceiver.sol";
import {IEVM2EVMOffRamp} from "../../interfaces/offRamp/IEVM2EVMOffRamp.sol";
import {IFeeManager} from "../../interfaces/fees/IFeeManager.sol";
import {IRouter} from "../../interfaces/router/IRouter.sol";

import {Internal} from "../../models/Internal.sol";
import {Common} from "../../models/Common.sol";
import {FeeManagerSetup} from "../fees/FeeManager.t.sol";
import {MockCommitStore} from "../mocks/MockCommitStore.sol";
import {SimpleMessageReceiver} from "../helpers/receivers/SimpleMessageReceiver.sol";
import {EVM2EVMOffRampHelper} from "../helpers/ramps/EVM2EVMOffRampHelper.sol";
import "../TokenSetup.t.sol";
import "../router/RouterSetup.t.sol";

contract EVM2EVMOffRampSetup is TokenSetup, FeeManagerSetup {
  ICommitStore internal s_mockCommitStore;
  IAny2EVMMessageReceiver internal s_receiver;
  IAny2EVMMessageReceiver internal s_secondary_receiver;

  EVM2EVMOffRampHelper internal s_offRamp;

  uint256 internal constant EXECUTION_FEE_AMOUNT = 1e18;

  event ExecutionStateChanged(
    uint64 indexed sequenceNumber,
    bytes32 indexed messageId,
    Internal.MessageExecutionState state
  );
  event SkippedIncorrectNonce(uint64 indexed nonce, address indexed sender);

  function setUp() public virtual override(TokenSetup, FeeManagerSetup) {
    TokenSetup.setUp();
    FeeManagerSetup.setUp();

    s_mockCommitStore = new MockCommitStore();
    s_receiver = new SimpleMessageReceiver();
    s_secondary_receiver = new SimpleMessageReceiver();

    deployOffRamp(s_mockCommitStore, s_destFeeManager, s_destRouter);
  }

  function deployOffRamp(
    ICommitStore commitStore,
    IFeeManager feeManager,
    IRouter router
  ) internal {
    s_offRamp = new EVM2EVMOffRampHelper(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      ON_RAMP_ADDRESS,
      offRampConfig(feeManager, commitStore, router),
      s_afn,
      getCastedSourceTokens(),
      getCastedDestinationPools(),
      rateLimiterConfig()
    );

    s_offRamp.setPrices(getCastedDestinationTokens(), getTokenPrices());
    s_destFeeManager.setFeeUpdater(address(s_offRamp));
    address[] memory s_offRamps = new address[](1);
    s_offRamps[0] = address(s_offRamp);
    IRouter.OnRampUpdate[] memory onRampUpdates = new IRouter.OnRampUpdate[](0);
    IRouter.OffRampUpdate[] memory offRampUpdates = new IRouter.OffRampUpdate[](1);
    offRampUpdates[0] = IRouter.OffRampUpdate({sourceChainId: SOURCE_CHAIN_ID, offRamps: s_offRamps});
    s_destRouter.applyRampUpdates(onRampUpdates, offRampUpdates);

    LockReleaseTokenPool(address(s_destPools[0])).setOffRamp(address(s_offRamp), true);
    LockReleaseTokenPool(address(s_destPools[1])).setOffRamp(address(s_offRamp), true);
  }

  function _convertToGeneralMessage(Internal.EVM2EVMMessage memory original)
    internal
    view
    returns (Common.Any2EVMMessage memory message)
  {
    uint256 numberOfTokens = original.tokensAndAmounts.length;
    Common.EVMTokenAndAmount[] memory destTokensAndAmounts = new Common.EVMTokenAndAmount[](numberOfTokens);

    for (uint256 i = 0; i < numberOfTokens; ++i) {
      IPool pool = s_offRamp.getPoolBySourceToken(IERC20(original.tokensAndAmounts[i].token));
      destTokensAndAmounts[i].token = address(pool.getToken());
      destTokensAndAmounts[i].amount = original.tokensAndAmounts[i].amount;
    }

    return
      Common.Any2EVMMessage({
        sourceChainId: original.sourceChainId,
        sender: abi.encode(original.sender),
        data: original.data,
        destTokensAndAmounts: destTokensAndAmounts
      });
  }

  function _generateAny2EVMMessageNoTokens(uint64 sequenceNumber)
    internal
    view
    returns (Internal.EVM2EVMMessage memory)
  {
    return _generateAny2EVMMessage(sequenceNumber, getCastedSourceEVMTokenAndAmountsWithZeroAmounts());
  }

  function _generateAny2EVMMessageWithTokens(uint64 sequenceNumber, uint256[] memory amounts)
    internal
    view
    returns (Internal.EVM2EVMMessage memory)
  {
    Common.EVMTokenAndAmount[] memory tokensAndAmounts = getCastedSourceEVMTokenAndAmountsWithZeroAmounts();
    for (uint256 i = 0; i < tokensAndAmounts.length; ++i) {
      tokensAndAmounts[i].amount = amounts[i];
    }
    return _generateAny2EVMMessage(sequenceNumber, tokensAndAmounts);
  }

  function _generateAny2EVMMessage(uint64 sequenceNumber, Common.EVMTokenAndAmount[] memory tokensAndAmounts)
    internal
    view
    returns (Internal.EVM2EVMMessage memory)
  {
    bytes memory data = abi.encode(0);
    Internal.EVM2EVMMessage memory message = Internal.EVM2EVMMessage({
      sequenceNumber: sequenceNumber,
      sender: OWNER,
      nonce: sequenceNumber,
      gasLimit: GAS_LIMIT,
      strict: false,
      sourceChainId: SOURCE_CHAIN_ID,
      receiver: address(s_receiver),
      data: data,
      tokensAndAmounts: tokensAndAmounts,
      feeToken: tokensAndAmounts[0].token,
      feeTokenAmount: uint256(0),
      messageId: ""
    });
    message.messageId = Internal._hash(
      message,
      keccak256(abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_ID, DEST_CHAIN_ID, ON_RAMP_ADDRESS))
    );

    return message;
  }

  function _generateBasicMessages() internal view returns (Internal.EVM2EVMMessage[] memory) {
    Internal.EVM2EVMMessage[] memory messages = new Internal.EVM2EVMMessage[](1);
    messages[0] = _generateAny2EVMMessageNoTokens(1);
    return messages;
  }

  function _generateMessagesWithTokens() internal view returns (Internal.EVM2EVMMessage[] memory) {
    Internal.EVM2EVMMessage[] memory messages = new Internal.EVM2EVMMessage[](2);
    Common.EVMTokenAndAmount[] memory tokensAndAmounts = getCastedSourceEVMTokenAndAmountsWithZeroAmounts();
    tokensAndAmounts[0].amount = 1e18;
    tokensAndAmounts[1].amount = 5e18;
    messages[0] = _generateAny2EVMMessage(1, tokensAndAmounts);
    messages[1] = _generateAny2EVMMessage(2, tokensAndAmounts);
    return messages;
  }

  function _generateReportFromMessages(Internal.EVM2EVMMessage[] memory messages)
    internal
    pure
    returns (Internal.ExecutionReport memory)
  {
    bytes[] memory encodedMessages = new bytes[](messages.length);
    uint64[] memory sequenceNumbers = new uint64[](messages.length);
    for (uint256 i = 0; i < messages.length; ++i) {
      encodedMessages[i] = abi.encode(messages[i]);
      sequenceNumbers[i] = messages[i].sequenceNumber;
    }

    bytes32[] memory proofs = new bytes32[](0);
    Internal.FeeUpdate[] memory feeUpdates = new Internal.FeeUpdate[](0);

    return
      Internal.ExecutionReport({
        sequenceNumbers: sequenceNumbers,
        proofs: proofs,
        proofFlagBits: 2**256 - 1,
        encodedMessages: encodedMessages,
        feeUpdates: feeUpdates
      });
  }
}
