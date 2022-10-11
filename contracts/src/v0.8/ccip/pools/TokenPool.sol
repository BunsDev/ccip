// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {SafeERC20} from "../../vendor/SafeERC20.sol";
import {Pausable} from "../../vendor/Pausable.sol";
import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {BaseOnRamp, BaseOnRampInterface, PoolInterface, IERC20} from "../onRamp/BaseOnRamp.sol";
import {BaseOffRampInterface} from "../interfaces/offRamp/BaseOffRampInterface.sol";

/**
 * @notice Base abstract class with common functions for all token pools
 */
abstract contract TokenPool is PoolInterface, OwnerIsCreator, Pausable {
  IERC20 internal immutable i_token;
  mapping(BaseOnRampInterface => bool) internal s_onRamps;
  mapping(BaseOffRampInterface => bool) internal s_offRamps;

  constructor(IERC20 token) {
    i_token = token;
  }

  /**
   * @notice Pause the pool
   * @dev Only callable by the owner
   */
  function pause() external override onlyOwner {
    _pause();
  }

  /**
   * @notice Unpause the pool
   * @dev Only callable by the owner
   */
  function unpause() external override onlyOwner {
    _unpause();
  }

  /**
   * @notice Set an onRamp's permissions
   * @dev Only callable by the owner
   * @param onRamp The onRamp
   * @param permission Whether or not the onRamp has onRamp permissions on this contract
   */
  function setOnRamp(BaseOnRampInterface onRamp, bool permission) public onlyOwner {
    s_onRamps[onRamp] = permission;
  }

  /**
   * @notice Set an offRamp's permissions
   * @dev Only callable by the owner
   * @param offRamp The offRamp
   * @param permission Whether or not the offRamp has offRamp permissions on this contract
   */
  function setOffRamp(BaseOffRampInterface offRamp, bool permission) public onlyOwner {
    s_offRamps[offRamp] = permission;
  }

  /**
   * @notice Checks whether something is a permissioned onRamp on this contract
   * @return boolean
   */
  function isOnRamp(BaseOnRampInterface onRamp) public view returns (bool) {
    return s_onRamps[onRamp];
  }

  /**
   * @notice Checks whether something is a permissioned offRamp on this contract
   * @return boolean
   */
  function isOffRamp(BaseOffRampInterface offRamp) public view returns (bool) {
    return s_offRamps[offRamp];
  }

  /**
   * @notice Gets the underlying token
   * @return token
   */
  function getToken() public view override returns (IERC20 token) {
    return i_token;
  }

  /**
   * @notice Checks whether the msg.sender is either the owner, or a permissioned onRamp on this contract
   * @dev Reverts with a PermissionsError if check fails
   */
  function _validateOwnerOrOnRamp() internal view {
    if (msg.sender != owner() && !isOnRamp(BaseOnRampInterface(msg.sender))) revert PermissionsError();
  }

  /**
   * @notice Checks whether the msg.sender is either the owner, or a permissioned offRamp on this contract
   * @dev Reverts with a PermissionsError if check fails
   */
  function _validateOwnerOrOffRamp() internal view {
    if (msg.sender != owner() && !isOffRamp(BaseOffRampInterface(msg.sender))) revert PermissionsError();
  }

  /**
   * @notice Check permissions and limits of a lock or burn
   */
  modifier assertLockOrBurn() {
    _validateOwnerOrOnRamp();
    _;
  }

  /**
   * @notice Check permissions and limits of a lock or burn
   */
  modifier assertMintOrRelease() {
    _validateOwnerOrOffRamp();
    _;
  }
}
