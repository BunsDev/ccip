// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {OwnerIsCreator} from "../../../shared/access/OwnerIsCreator.sol";

import {IBridge} from "./IBridge.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice Interface for a liquidity container, this can be a CCIP token pool.
interface ILiquidityContainer {
  function provideLiquidity(uint256 amount) external;

  function withdrawLiquidity(uint256 amount) external;
}

contract LiquidityManager is OwnerIsCreator {
  using SafeERC20 for IERC20;

  error CannotBeZero();
  error InvalidDestinationChain(uint64 chainSelector);
  error Unauthorized(address caller);
  error InsufficientLiquidity(uint256 requested, uint256 available);

  event LiquidityTransferred(
    uint64 indexed fromChainSelector,
    uint64 indexed toChainSelector,
    address indexed to,
    uint256 amount
  );
  event LiquidityAdded(address indexed provider, uint256 indexed amount);
  event LiquidityRemoved(address indexed remover, uint256 indexed amount);

  struct CrossChainLiquidityContainer {
    address liquidityContainer;
    IBridge bridge;
    // Potentially some fields related to the bridge
  }

  /// @notice The token that this pool manages liquidity for.
  IERC20 public immutable i_localToken;

  /// @notice The chain selector belonging to the chain this pool is deployed on.
  uint64 internal immutable i_localChainSelector;

  /// @notice Mapping of chain selector to liquidity container on other chains
  mapping(uint64 chainSelector => CrossChainLiquidityContainer) private s_crossChainLiquidityContainers;

  /// @notice The liquidity container on the local chain
  /// @dev In the case of CCIP, this would be the token pool.
  ILiquidityContainer private s_localLiquidityContainer;

  /// @notice The address that is allowed to call rebalanceLiquidity
  /// @dev This address cannot withdraw liquidity from the system, only rebalance it
  /// between pre-approved locations. Only the owner can remove liquidity.
  address private s_liquidityManager;

  constructor(IERC20 token, uint64 localChainSelector, address liquidityManager) {
    if (address(token) == address(0) || localChainSelector == 0) {
      revert CannotBeZero();
    }
    i_localToken = token;
    i_localChainSelector = localChainSelector;
    s_liquidityManager = liquidityManager;
  }

  // ================================================================
  // │                    Liquidity management                      │
  // ================================================================

  function rebalanceLiquidity(uint64 chainSelector, uint256 amount) external {
    // Ensure only the owner and the liquidity manager can transfer liquidity
    if (msg.sender != s_liquidityManager && msg.sender != owner()) {
      revert Unauthorized(msg.sender);
    }

    uint256 currentBalance = i_localToken.balanceOf(address(s_localLiquidityContainer));
    if (currentBalance < amount) {
      revert InsufficientLiquidity(amount, currentBalance);
    }

    address destChainAddress = s_crossChainLiquidityContainers[chainSelector].liquidityContainer;
    if (destChainAddress == address(0)) {
      revert InvalidDestinationChain(chainSelector);
    }

    s_localLiquidityContainer.withdrawLiquidity(amount);

    // TODO send funds to bridge

    emit LiquidityTransferred(i_localChainSelector, chainSelector, destChainAddress, amount);
  }

  /// @notice Adds liquidity to the multi-chain system.
  /// @dev Anyone can call this function, but anyone other than the owner should regard
  /// adding liquidity as a donation to the system, as there is no way to get it out.
  function addLiquidity(uint256 amount) external {
    i_localToken.safeTransferFrom(msg.sender, address(this), amount);

    // Make sure this is tether compatible, as they have strange approval requirements
    // Should be good since all approvals are always immediately used.
    i_localToken.approve(address(s_localLiquidityContainer), amount);
    s_localLiquidityContainer.provideLiquidity(amount);

    emit LiquidityAdded(msg.sender, amount);
  }

  function removeLiquidity(uint256 amount) external onlyOwner {
    uint256 currentBalance = i_localToken.balanceOf(address(s_localLiquidityContainer));
    if (currentBalance < amount) {
      revert InsufficientLiquidity(amount, currentBalance);
    }

    s_localLiquidityContainer.withdrawLiquidity(amount);

    // Do we want to send to the msg.sender/owner or do we want a 2-phase withdraw?
    // 2 phase would be safer if the multisig doesn't want to have tokens and we
    // need to send to a different address.
    i_localToken.safeTransfer(msg.sender, amount);

    emit LiquidityRemoved(msg.sender, amount);
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  // TODO
}
