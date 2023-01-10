// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IPool} from "../pools/IPool.sol";

import {GEConsumer} from "../../models/GEConsumer.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

interface IEVM2AnyGEOnRamp {
  /**
   * @notice Get the fee for a given ccip message
   * @param message The message to calculate the cost for
   * @return fee The calculated fee
   */
  function getFee(GEConsumer.EVM2AnyGEMessage calldata message) external view returns (uint256 fee);

  /**
   * @notice Get the pool for a specific token
   * @param sourceToken The source chain token to get the pool for
   * @return pool IPool
   */
  function getPoolBySourceToken(IERC20 sourceToken) external view returns (IPool);

  /**
   * @notice Send a message to the remote chain
   * @dev approve() must have already been called on the token using the this ramp address as the spender.
   * @dev if the contract is paused, this function will revert.
   * @param message Message struct to send
   * @param originalSender The original initiator of the CCIP request
   */
  function forwardFromRouter(
    GEConsumer.EVM2AnyGEMessage memory message,
    uint256 feeTokenAmount,
    address originalSender
  ) external returns (bytes32);
}
