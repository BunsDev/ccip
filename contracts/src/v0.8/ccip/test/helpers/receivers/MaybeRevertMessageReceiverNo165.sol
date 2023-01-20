pragma solidity 0.8.15;

import "../../../interfaces/applications/IAny2EVMMessageReceiver.sol";

contract MaybeRevertMessageReceiverNo165 is IAny2EVMMessageReceiver {
  address private s_manager;
  bool public s_toRevert;
  event MessageReceived();

  constructor(bool toRevert) {
    s_manager = msg.sender;
    s_toRevert = toRevert;
  }

  function setRevert(bool toRevert) external {
    s_toRevert = toRevert;
  }

  function ccipReceive(Common.Any2EVMMessage calldata) external override {
    if (s_toRevert) {
      revert();
    }
    emit MessageReceived();
  }
}
