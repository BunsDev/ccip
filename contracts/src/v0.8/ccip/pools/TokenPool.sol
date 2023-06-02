// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {IPool} from "../interfaces/pools/IPool.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";

import {Pausable} from "../../vendor/Pausable.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.0/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.0/token/ERC20/IERC20.sol";
import {IERC165} from "../../vendor/openzeppelin-solidity/v4.8.0/utils/introspection/IERC165.sol";
import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableSet.sol";

/// @notice Base abstract class with common functions for all token pools.
abstract contract TokenPool is IPool, OwnerIsCreator, Pausable, IERC165 {
  using EnumerableSet for EnumerableSet.AddressSet;
  using RateLimiter for RateLimiter.TokenBucket;

  error PermissionsError();
  error NullAddressNotAllowed();
  error SenderNotAllowed(address sender);
  error AllowListNotEnabled();

  event Locked(address indexed sender, uint256 amount);
  event Burned(address indexed sender, uint256 amount);
  event Released(address indexed sender, address indexed recipient, uint256 amount);
  event Minted(address indexed sender, address indexed recipient, uint256 amount);
  event OnRampAllowanceSet(address onRamp, bool allowed);
  event OffRampAllowanceSet(address onRamp, bool allowed);
  event AllowListAdd(address sender);
  event AllowListRemove(address sender);

  struct RampUpdate {
    address ramp;
    bool allowed;
  }

  // The immutable token that belongs to this pool.
  IERC20 internal immutable i_token;
  // The immutable flag that indicates if the pool is access-controled.
  bool internal immutable i_allowlistEnabled;
  // A set of allowed onRamps.
  EnumerableSet.AddressSet internal s_onRamps;
  // A set of allowed offRamps.
  EnumerableSet.AddressSet internal s_offRamps;
  // A set of addresses allowed to trigger lockOrBurn as original senders.
  EnumerableSet.AddressSet internal s_allowList;
  // The token bucket object that contains the bucket state.
  RateLimiter.TokenBucket private s_rateLimiter;

  constructor(IERC20 token, address[] memory allowlist, RateLimiter.Config memory rateLimiterConfig) {
    if (address(token) == address(0)) revert NullAddressNotAllowed();

    s_rateLimiter = RateLimiter.TokenBucket({
      rate: rateLimiterConfig.rate,
      capacity: rateLimiterConfig.capacity,
      tokens: rateLimiterConfig.capacity,
      lastUpdated: uint32(block.timestamp),
      isEnabled: rateLimiterConfig.isEnabled
    });

    i_token = token;

    // pool can be set as permissioned or permissionless at deployment time
    i_allowlistEnabled = allowlist.length > 0;
    if (i_allowlistEnabled) {
      _applyAllowListUpdates(new address[](0), allowlist);
    }
  }

  /// @inheritdoc IPool
  function getToken() public view override returns (IERC20 token) {
    return i_token;
  }

  /// @inheritdoc IERC165
  function supportsInterface(bytes4 interfaceId) public pure virtual override returns (bool) {
    return interfaceId == type(IPool).interfaceId || interfaceId == type(IERC165).interfaceId;
  }

  // ================================================================
  // |                      Ramp permissions                        |
  // ================================================================

  /// @notice Checks whether something is a permissioned onRamp on this contract.
  /// @return true if the given address is a permissioned onRamp.
  function isOnRamp(address onRamp) public view returns (bool) {
    return s_onRamps.contains(onRamp);
  }

  /// @notice Checks whether something is a permissioned offRamp on this contract.
  /// @return true is the given address is a permissioned offRamp.
  function isOffRamp(address offRamp) public view returns (bool) {
    return s_offRamps.contains(offRamp);
  }

  /// @notice Sets permissions for all on and offRamps.
  /// @dev Only callable by the owner
  /// @param onRamps A list of onRamps and their new permission status
  /// @param offRamps A list of offRamps and their new permission status
  function applyRampUpdates(RampUpdate[] memory onRamps, RampUpdate[] memory offRamps) public virtual onlyOwner {
    for (uint256 i = 0; i < onRamps.length; ++i) {
      RampUpdate memory update = onRamps[i];

      if (update.allowed ? s_onRamps.add(update.ramp) : s_onRamps.remove(update.ramp)) {
        emit OnRampAllowanceSet(onRamps[i].ramp, onRamps[i].allowed);
      }
    }

    for (uint256 i = 0; i < offRamps.length; ++i) {
      RampUpdate memory update = offRamps[i];

      if (update.allowed ? s_offRamps.add(update.ramp) : s_offRamps.remove(update.ramp)) {
        emit OffRampAllowanceSet(offRamps[i].ramp, offRamps[i].allowed);
      }
    }
  }

  // ================================================================
  // |                        Rate limiting                         |
  // ================================================================

  /// @notice Consumes rate limiting capacity in this pool
  function _consumeRateLimit(uint256 amount) internal {
    s_rateLimiter._consume(amount);
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @return The token bucket.
  function currentRateLimiterState() public view returns (RateLimiter.TokenBucket memory) {
    return s_rateLimiter._currentTokenBucketState();
  }

  /// @notice Sets the rate limited config.
  /// @param config The new rate limiter config.
  /// @dev should only be callable by the owner or token limit admin.
  function setRateLimiterConfig(RateLimiter.Config memory config) public onlyOwner {
    s_rateLimiter._setTokenBucketConfig(config);
  }

  // ================================================================
  // |                           Access                             |
  // ================================================================

  /// @notice Checks whether the msg.sender is a permissioned onRamp on this contract
  /// @dev Reverts with a PermissionsError if check fails
  modifier onlyOnRamp() {
    if (!isOnRamp(msg.sender)) revert PermissionsError();
    _;
  }

  /// @notice Checks whether the msg.sender is a permissioned offRamp on this contract
  /// @dev Reverts with a PermissionsError if check fails
  modifier onlyOffRamp() {
    if (!isOffRamp(msg.sender)) revert PermissionsError();
    _;
  }

  /// @notice Pauses the token pool.
  function pause() external onlyOwner {
    _pause();
  }

  /// @notice Unpauses the token pool.
  function unpause() external onlyOwner {
    _unpause();
  }

  // ================================================================
  // |                          Allowlist                           |
  // ================================================================

  modifier checkAllowList(address sender) {
    if (i_allowlistEnabled && !s_allowList.contains(sender)) revert SenderNotAllowed(sender);
    _;
  }

  /// @notice Gets whether the allowList functionality is enabled.
  /// @return true is enabled, false if not.
  function getAllowListEnabled() external view returns (bool) {
    return i_allowlistEnabled;
  }

  /// @notice Gets the allowed addresses.
  /// @return The allowed addresses.
  function getAllowList() external view returns (address[] memory) {
    address[] memory allowList = new address[](s_allowList.length());
    for (uint256 i = 0; i < s_allowList.length(); ++i) {
      allowList[i] = s_allowList.at(i);
    }
    return allowList;
  }

  /// @notice Apply updates to the allow list.
  /// @param removes The addresses to be removed.
  /// @param adds The addresses to be added.
  /// @dev allowListing will be removed before public launch
  function applyAllowListUpdates(address[] calldata removes, address[] calldata adds) external onlyOwner {
    _applyAllowListUpdates(removes, adds);
  }

  /// @notice Internal version of applyAllowListUpdates to allow for reuse in the constructor.
  function _applyAllowListUpdates(address[] memory removes, address[] memory adds) internal {
    if (!i_allowlistEnabled) revert AllowListNotEnabled();

    for (uint256 i = 0; i < removes.length; ++i) {
      address toRemove = removes[i];
      if (s_allowList.remove(toRemove)) {
        emit AllowListRemove(toRemove);
      }
    }
    for (uint256 i = 0; i < adds.length; ++i) {
      address toAdd = adds[i];
      if (toAdd == address(0)) {
        continue;
      }
      if (s_allowList.add(toAdd)) {
        emit AllowListAdd(toAdd);
      }
    }
  }
}
