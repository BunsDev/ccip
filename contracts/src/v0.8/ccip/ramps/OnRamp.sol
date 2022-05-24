// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../pools/PoolCollector.sol";
import "../interfaces/OnRampInterface.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../utils/CCIP.sol";
import "../health/HealthChecker.sol";
import "../pools/TokenPoolRegistry.sol";
import "./PriceFeedRegistry.sol";
import "../../vendor/SafeERC20.sol";

/**
 * @notice An implementation of an On Ramp, which enables just a single token to be
 * used in the protocol.
 */
contract OnRamp is
  OnRampInterface,
  TypeAndVersionInterface,
  HealthChecker,
  TokenPoolRegistry,
  PriceFeedRegistry,
  PoolCollector
{
  using SafeERC20 for IERC20;

  // Chain ID of the source chain (where this contract is deployed)
  uint256 public immutable CHAIN_ID;

  // Destination chain => sequence number
  mapping(uint256 => uint64) private s_sequenceNumberPerDestinationChain;
  // List of destination chains
  uint256[] private s_destinationChains;
  // OnRamp config
  OnRampConfig private s_config;

  // Whether the allowlist is enabled
  bool private s_allowlistEnabled;
  // Addresses that are allowed to send messages
  mapping(address => bool) private s_allowed;
  // List of allowed addresses
  address[] private s_allowList;

  constructor(
    uint256 chainId,
    uint256[] memory destinationChainIds,
    IERC20[] memory tokens,
    PoolInterface[] memory pools,
    AggregatorV2V3Interface[] memory feeds,
    address[] memory allowlist,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    OnRampConfig memory config
  ) HealthChecker(afn, maxTimeWithoutAFNSignal) TokenPoolRegistry(tokens, pools) PriceFeedRegistry(tokens, feeds) {
    CHAIN_ID = chainId;
    s_destinationChains = destinationChainIds;
    for (uint256 i = 0; i < destinationChainIds.length; i++) {
      s_sequenceNumberPerDestinationChain[destinationChainIds[i]] = 1;
    }
    if (allowlist.length > 0) {
      s_allowlistEnabled = true;
      s_allowList = allowlist;
    }
    for (uint256 i = 0; i < allowlist.length; i++) {
      s_allowed[allowlist[i]] = true;
    }
    s_config = config;
  }

  /**
   * @notice Send a message to the remote chain
   * @dev the first token in the payload is used as the fee token
   * @dev approve() must have already been called on the token using the this ramp address as the spender.
   * @dev if the contract is paused, this function will revert.
   * @param payload Message struct to send
   */
  function requestCrossChainSend(CCIP.MessagePayload memory payload, address originalSender)
    external
    override
    whenNotPaused
    whenHealthy
    returns (uint64)
  {
    address sender = msg.sender;
    if (sender != s_config.router) revert MustBeCalledByRouter();
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    uint64 sequenceNumber = s_sequenceNumberPerDestinationChain[payload.destinationChainId];
    // Check that the destination chain has been configured
    // Assumes that any configured destination chains sequence number are initialized with 1
    if (sequenceNumber == 0) revert UnsupportedDestinationChain(payload.destinationChainId);
    // Check that payload is formed correctly
    if (payload.data.length > uint256(s_config.maxDataSize))
      revert MessageTooLarge(uint256(s_config.maxDataSize), payload.data.length);
    if (payload.tokens.length > uint256(s_config.maxTokensLength) || payload.tokens.length != payload.amounts.length)
      revert UnsupportedNumberOfTokens();

    if (s_allowlistEnabled && !s_allowed[originalSender]) revert SenderNotAllowed(originalSender);

    for (uint256 i = 0; i < payload.tokens.length; i++) {
      IERC20 token = payload.tokens[i];
      PoolInterface pool = getPool(token);
      if (address(pool) == address(0)) revert UnsupportedToken(token);
      uint256 amount = payload.amounts[i];
      pool.lockOrBurn(amount);
    }

    // Emit message request
    CCIP.Message memory message = CCIP.Message({
      sequenceNumber: sequenceNumber,
      sourceChainId: CHAIN_ID,
      sender: originalSender,
      payload: payload
    });
    s_sequenceNumberPerDestinationChain[payload.destinationChainId] = sequenceNumber + 1;
    emit CrossChainSendRequested(message);
    return message.sequenceNumber;
  }

  /**
   * @notice Get the required fee for a specific fee token
   * @param feeToken token to get the fee for
   * @return fee uint256
   */
  function getRequiredFee(IERC20 feeToken) public view override returns (uint256) {
    AggregatorV2V3Interface feed = getFeed(feeToken);
    if (address(feed) == address(0)) revert UnsupportedFeeToken(feeToken);
    return s_config.relayingFeeJuels * uint256(feed.latestAnswer());
  }

  /**
   * @notice Get the pool for a specific token
   * @param token token to get the pool for
   * @return pool PoolInterface
   */
  function getTokenPool(IERC20 token) external view override returns (PoolInterface) {
    return getPool(token);
  }

  function setAllowlistEnabled(bool enabled) external onlyOwner {
    s_allowlistEnabled = enabled;
    emit AllowlistEnabledSet(enabled);
  }

  function getAllowlistEnabled() external view returns (bool) {
    return s_allowlistEnabled;
  }

  function setAllowlist(address[] calldata allowlist) external onlyOwner {
    // Remove existing allowlist
    address[] memory existingList = s_allowList;
    for (uint256 i = 0; i < existingList.length; i++) {
      s_allowed[existingList[i]] = false;
    }

    // Set the new allowlist
    s_allowList = allowlist;
    for (uint256 i = 0; i < allowlist.length; i++) {
      s_allowed[allowlist[i]] = true;
    }
    emit AllowlistSet(allowlist);
  }

  function getAllowlist() external view returns (address[] memory) {
    return s_allowList;
  }

  function setConfig(OnRampConfig calldata config) external onlyOwner {
    s_config = config;
    emit OnRampConfigSet(config);
  }

  function getConfig() external view returns (OnRampConfig memory config) {
    return s_config;
  }

  function getDestinationChains() external view returns (uint256[] memory) {
    return s_destinationChains;
  }

  function getSequenceNumberOfDestinationChain(uint256 destinationChainId) external view returns (uint64) {
    return s_sequenceNumberPerDestinationChain[destinationChainId];
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "OnRamp 0.0.1";
  }
}
