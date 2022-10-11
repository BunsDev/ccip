// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {CCIP} from "../../models/Models.sol";
import {EVM2EVMTollOnRampInterface} from "./EVM2EVMTollOnRampInterface.sol";
import {BaseOnRampRouterInterface} from "./BaseOnRampRouterInterface.sol";

interface EVM2AnyTollOnRampRouterInterface is BaseOnRampRouterInterface {
  error OnRampAlreadySet(uint256 chainId, EVM2EVMTollOnRampInterface onRamp);

  event OnRampSet(uint256 indexed chainId, EVM2EVMTollOnRampInterface indexed onRamp);

  /**
   * @notice Request a message to be sent to the destination chain
   * @param destinationChainId The destination chain ID
   * @param message The message payload
   * @return The sequence number assigned to message
   */
  function ccipSend(uint256 destinationChainId, CCIP.EVM2AnyTollMessage calldata message) external returns (uint64);

  /**
   * @notice Set chainId => onRamp mapping
   * @dev only callable by owner
   * @param chainId destination chain ID
   * @param onRamp OnRamp to use for that destination chain
   */
  function setOnRamp(uint256 chainId, EVM2EVMTollOnRampInterface onRamp) external;

  /**
   * @notice Gets the current OnRamp for the specified chain ID
   * @param chainId Chain ID to get ramp details for
   * @return onRamp
   */
  function getOnRamp(uint256 chainId) external view returns (EVM2EVMTollOnRampInterface);
}
