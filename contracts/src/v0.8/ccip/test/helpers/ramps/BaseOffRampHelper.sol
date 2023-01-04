pragma solidity ^0.8.0;

import "../../../offRamp/BaseOffRamp.sol";

contract BaseOffRampHelper is BaseOffRamp {
  constructor(
    uint64 sourceChainId,
    uint64 chainId,
    address onRampAddress,
    ICommitStore commitStore,
    IAFN afn,
    IERC20[] memory sourceTokens,
    IPool[] memory pools,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin
  )
    BaseOffRamp(
      sourceChainId,
      chainId,
      onRampAddress,
      commitStore,
      afn,
      sourceTokens,
      pools,
      rateLimiterConfig,
      tokenLimitsAdmin
    )
  {}

  function setExecutionState(uint64 sequenceNumber, Internal.MessageExecutionState state) public {
    s_executedMessages[sequenceNumber] = state;
  }

  function releaseOrMintToken(
    IPool pool,
    uint256 amount,
    address receiver
  ) external {
    _releaseOrMintToken(pool, amount, receiver);
  }

  function releaseOrMintTokens(
    address[] memory pools,
    Common.EVMTokenAndAmount[] memory tokensAndAmounts,
    address receiver
  ) external {
    _releaseOrMintTokens(pools, tokensAndAmounts, receiver);
  }

  function verifyMessages(
    bytes32[] memory hashedLeaves,
    bytes32[] memory innerProofs,
    uint256 innerProofFlagBits,
    bytes32[] memory outerProofs,
    uint256 outerProofFlagBits
  ) external returns (uint256, uint256) {
    return _verifyMessages(hashedLeaves, innerProofs, innerProofFlagBits, outerProofs, outerProofFlagBits);
  }

  function getPool_helper(IERC20 token) external view returns (IPool pool) {
    return _getPool(token);
  }
}
