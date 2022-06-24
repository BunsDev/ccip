// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./TokenPool.sol";
import "../../vendor/ERC20.sol";

/**
 * @notice This pool mints and burns its own tokens, representing a wrapped form of the native token
 * on a source chain - similar to WBTC.
 */
contract WrappedTokenPool is TokenPool, ERC20 {
  using TokenLimits for TokenLimits.TokenBucket;

  constructor(
    string memory name,
    string memory symbol,
    BucketConfig memory burnConfig,
    BucketConfig memory mintConfig
  ) TokenPool(IERC20(address(this)), burnConfig, mintConfig) ERC20(name, symbol) {}

  /**
   * @notice Burn the token in the pool
   * @param amount Amount to burn
   */
  function lockOrBurn(uint256 amount) external override whenNotPaused assertLockOrBurn(amount) {
    _burn(address(this), amount);
    emit Burned(msg.sender, amount);
  }

  /**
   * @notice Mint tokens fromm the pool to the recipient
   * @param recipient Recipient address
   * @param amount Amount to mint
   */
  function releaseOrMint(address recipient, uint256 amount)
    external
    override
    whenNotPaused
    assertMintOrRelease(amount)
  {
    _mint(recipient, amount);
    emit Minted(msg.sender, recipient, amount);
  }
}
