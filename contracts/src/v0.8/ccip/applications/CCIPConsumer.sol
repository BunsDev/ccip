// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IAny2EVMMessageReceiver} from "../interfaces/applications/IAny2EVMMessageReceiver.sol";
import {IAny2EVMOffRampRouter} from "../interfaces/offRamp/IAny2EVMOffRampRouter.sol";
import {IGERouter} from "../interfaces/router/IGERouter.sol";
import {IERC165} from "../../vendor/IERC165.sol";

import {GEConsumer} from "../models/GEConsumer.sol";
import {Common} from "../models/Common.sol";

/// @title CCIPConsumer - Base contract for CCIP applications that can both send and receive messages.
abstract contract CCIPConsumer is IAny2EVMMessageReceiver, IERC165 {
  IGERouter private immutable i_router;
  address private s_feeToken;

  constructor(address router, address feeToken) {
    if (router == address(0)) revert InvalidRouter(address(0));
    i_router = IGERouter(router);

    _setFeeToken(feeToken);
  }

  /**
   * @notice IERC165 supports an interfaceId
   * @param interfaceId The interfaceId to check
   * @return true if the interfaceId is supported
   */
  function supportsInterface(bytes4 interfaceId) public pure override returns (bool) {
    return interfaceId == type(IAny2EVMMessageReceiver).interfaceId || interfaceId == type(IERC165).interfaceId;
  }

  /// @inheritdoc IAny2EVMMessageReceiver
  function ccipReceive(Common.Any2EVMMessage calldata message) external override onlyRouter {
    _ccipReceive(message);
  }

  /**
   * @notice Override this function in your implementation.
   * @param message Any2EVMMessage
   */
  function _ccipReceive(Common.Any2EVMMessage memory message) internal virtual;

  /**
   * @notice Request a message to be sent to the destination chain
   * @dev Internal - Accessible by inheriting contracts
   * @param destinationChainId The destination chain ID
   * @param message The message payload
   * @return messageId assigned to message
   */
  function _ccipSend(uint64 destinationChainId, GEConsumer.EVM2AnyGEMessage memory message)
    internal
    returns (bytes32 messageId)
  {
    return i_router.ccipSend(destinationChainId, message);
  }

  /////////////////////////////////////////////////////////////////////
  // Plumbing
  /////////////////////////////////////////////////////////////////////

  /**
   * @notice Return the current router
   * @return i_router address
   */
  function getRouter() public view returns (address) {
    return address(i_router);
  }

  event FeeTokenSet(address indexed feeToken);

  /**
   * @notice Set the feeToken
   * @dev Internal - Accessible by inheriting contracts
   */
  function _setFeeToken(address feeToken) internal {
    s_feeToken = feeToken;
    emit FeeTokenSet(feeToken);
  }

  /**
   * @notice Return the current feeToken address
   * @return feeToken address
   */
  function getFeeToken() public view returns (address) {
    return s_feeToken;
  }

  error InvalidRouter(address router);

  /**
   * @dev only calls from the set router are accepted.
   */
  modifier onlyRouter() {
    if (msg.sender != address(i_router)) revert InvalidRouter(msg.sender);
    _;
  }
}
