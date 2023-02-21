// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../interfaces/ICommitStore.sol";
import "../interfaces/offRamp/IEVM2EVMOffRamp.sol";
import "../interfaces/rateLimiter/IAggregateRateLimiter.sol";
import "../interfaces/router/IRouter.sol";
import "../interfaces/onRamp/IEVM2EVMOnRamp.sol";

contract StructFactory {
  // addresses
  address internal constant OWNER = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;
  address internal constant STRANGER = address(999999);
  address internal constant DUMMY_CONTRACT_ADDRESS = 0x1111111111111111111111111111111111111112;
  address internal constant ON_RAMP_ADDRESS = 0x11118e64e1FB0c487f25dD6D3601FF6aF8d32E4e;
  address internal constant ZERO_ADDRESS = address(0);
  address internal constant USER_1 = address(1);
  address internal constant USER_2 = address(2);
  address internal constant USER_3 = address(3);
  address internal constant USER_4 = address(4);

  // afn

  function afnConstructorArgs()
    internal
    pure
    returns (
      address[] memory,
      uint256[] memory,
      uint256,
      uint256
    )
  {
    address[] memory participants = new address[](4);
    participants[0] = USER_1;
    participants[1] = USER_2;
    participants[2] = USER_3;
    participants[3] = USER_4;
    uint256[] memory weights = new uint256[](4);
    weights[0] = WEIGHT_1;
    weights[1] = WEIGHT_10;
    weights[2] = WEIGHT_20;
    weights[3] = WEIGHT_40;
    uint256 blessingThreshold = WEIGHT_10 + WEIGHT_20 + WEIGHT_40;
    uint256 badSignalThreshold = WEIGHT_1 + WEIGHT_10 + WEIGHT_20 + WEIGHT_40;
    return (participants, weights, blessingThreshold, badSignalThreshold);
  }

  uint256 internal constant ZERO = 0;
  uint256 internal constant WEIGHT_1 = 1;
  uint256 internal constant WEIGHT_10 = 10;
  uint256 internal constant WEIGHT_20 = 20;
  uint256 internal constant WEIGHT_40 = 40;

  // message info
  uint64 internal constant SOURCE_CHAIN_ID = 1;
  uint64 internal constant DEST_CHAIN_ID = 2;
  uint256 internal constant GAS_LIMIT = 100_000;

  // timing
  uint256 internal constant BLOCK_TIME = 1234567890;
  uint32 internal constant TWELVE_HOURS = 60 * 60 * 12;

  // offRamp
  uint256 internal constant POOL_BALANCE = 1e25;
  uint32 internal constant EXECUTION_DELAY_SECONDS = 0;
  uint32 internal constant MAX_DATA_SIZE = 500;
  uint16 internal constant MAX_TOKENS_LENGTH = 5;
  uint32 internal constant PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS = 500;
  uint64 internal constant MAX_GAS_LIMIT = 4_000_000;

  function offRampConfig(ICommitStore commitStore, IRouter router)
    internal
    pure
    returns (IEVM2EVMOffRamp.OffRampConfig memory)
  {
    return
      IEVM2EVMOffRamp.OffRampConfig({
        router: address(router),
        commitStore: address(commitStore),
        executionDelaySeconds: EXECUTION_DELAY_SECONDS,
        maxDataSize: MAX_DATA_SIZE,
        maxTokensLength: MAX_TOKENS_LENGTH,
        permissionLessExecutionThresholdSeconds: PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS
      });
  }

  function onRampConfig() internal pure returns (IEVM2EVMOnRamp.OnRampConfig memory) {
    return
      IEVM2EVMOnRamp.OnRampConfig({
        maxDataSize: MAX_DATA_SIZE,
        maxTokensLength: MAX_TOKENS_LENGTH,
        maxGasLimit: MAX_GAS_LIMIT
      });
  }

  // Rate limiter
  address constant TOKEN_LIMIT_ADMIN = 0x11118e64e1FB0c487f25dD6D3601FF6aF8d32E4e;

  function rateLimiterConfig() internal pure returns (IAggregateRateLimiter.RateLimiterConfig memory) {
    return IAggregateRateLimiter.RateLimiterConfig({capacity: 100e28, rate: 1e15, admin: TOKEN_LIMIT_ADMIN});
  }

  function getTokenPrices() internal pure returns (uint256[] memory prices) {
    prices = new uint256[](2);
    prices[0] = 1;
    prices[1] = 8;
    return prices;
  }
}
