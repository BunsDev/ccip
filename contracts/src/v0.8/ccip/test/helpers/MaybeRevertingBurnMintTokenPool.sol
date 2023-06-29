// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {BurnMintTokenPool} from "../../pools/BurnMintTokenPool.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {IBurnMintERC20} from "../../../shared/token/ERC20/IBurnMintERC20.sol";

contract MaybeRevertingBurnMintTokenPool is BurnMintTokenPool {
  bytes public s_revertReason = "";

  constructor(IBurnMintERC20 token, address[] memory allowlist) BurnMintTokenPool(token, allowlist) {}

  function setShouldRevert(bytes calldata revertReason) external {
    s_revertReason = revertReason;
  }

  /// @notice Reverts depending on the value of `s_revertReason`
  function releaseOrMint(
    bytes memory,
    address receiver,
    uint256 amount,
    uint64,
    bytes memory
  ) external virtual override onlyOffRamp {
    bytes memory revertReason = s_revertReason;
    if (revertReason.length != 0) {
      assembly {
        revert(add(32, revertReason), mload(revertReason))
      }
    }
    _consumeOnRampRateLimit(amount);
    IBurnMintERC20(address(i_token)).mint(receiver, amount);
    emit Minted(msg.sender, receiver, amount);
  }
}
