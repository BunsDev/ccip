// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/offRamp/BaseOffRampInterface.sol";
import "../../interfaces/applications/Any2EVMMessageReceiverInterface.sol";
import "../../interfaces/BlobVerifierInterface.sol";

contract MockOffRamp is BaseOffRampInterface {
  IERC20 public s_token;

  function deliverMessageTo(Any2EVMMessageReceiverInterface recipient, CCIP.Any2EVMMessage calldata message) external {
    recipient.ccipReceive(message);
  }

  function i_sourceChainId() external view returns (uint256) {}

  function i_chainId() external view returns (uint256) {}

  function setRouter(Any2EVMOffRampRouterInterface router) external {}

  function getRouter() external pure override returns (Any2EVMOffRampRouterInterface) {
    return Any2EVMOffRampRouterInterface(address(0));
  }

  /**
   * @notice ccipReceive implements the receive function to create a
   * collision if some other method happens to hash to the same signature/
   */
  function ccipReceive(CCIP.Any2EVMMessageFromSender calldata) external pure {
    revert();
  }

  function execute(CCIP.ExecutionReport memory report, bool needFee) external override {}

  function executeSingleMessage(CCIP.EVM2EVMTollMessage memory message) external {}

  function setToken(IERC20 token) external {
    s_token = token;
  }

  function TOKEN() external view returns (IERC20) {
    return s_token;
  }

  /// @inheritdoc BaseOffRampInterface
  function getExecutionState(uint64) public pure returns (CCIP.MessageExecutionState) {
    return CCIP.MessageExecutionState.SUCCESS;
  }

  /// @inheritdoc BaseOffRampInterface
  function getBlobVerifier() public pure returns (BlobVerifierInterface) {
    return BlobVerifierInterface(address(1));
  }

  /// @inheritdoc BaseOffRampInterface
  function setBlobVerifier(BlobVerifierInterface blobVerifier) public pure {}

  /// @inheritdoc BaseOffRampInterface
  function getConfig() public pure returns (OffRampConfig memory config) {}

  /// @inheritdoc BaseOffRampInterface
  function setConfig(OffRampConfig memory config) public {}
}
