// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./EVM2EVMSubscriptionOnRampSetup.t.sol";

/// @notice #constructor
contract EVM2EVMSubscriptionOnRamp_constructor is EVM2EVMSubscriptionOnRampSetup {
  function testSuccess() public {
    // typeAndVersion
    assertEq("EVM2EVMSubscriptionOnRamp 1.0.0", s_onRamp.typeAndVersion());

    // owner
    assertEq(OWNER, s_onRamp.owner());

    // baseOnRamp
    assertEq(RELAYING_FEE_JUELS, s_onRamp.getConfig().relayingFeeJuels);
    assertEq(MAX_DATA_SIZE, s_onRamp.getConfig().maxDataSize);
    assertEq(MAX_TOKENS_LENGTH, s_onRamp.getConfig().maxTokensLength);

    assertEq(SOURCE_CHAIN_ID, s_onRamp.i_chainId());
    assertEq(DEST_CHAIN_ID, s_onRamp.i_destinationChainId());

    assertEq(address(s_onRampRouter), s_onRamp.getRouter());
    assertEq(1, s_onRamp.getExpectedNextSequenceNumber());

    // HealthChecker
    assertEq(address(s_afn), address(s_onRamp.getAFN()));
  }
}

/// @notice #forwardFromRouter
contract EVM2EVMSubscriptionOnRamp_forwardFromRouter is EVM2EVMSubscriptionOnRampSetup {
  function setUp() public virtual override {
    EVM2EVMSubscriptionOnRampSetup.setUp();

    // Since we'll mostly be testing for valid calls from the router we'll
    // mock all calls to be originating from the router and re-mock in
    // tests that require failure.
    changePrank(address(s_onRampRouter));
  }

  // Success

  // Asserts that forwardFromRouter succeeds when called from the
  // router.
  function testSuccess() public {
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
  }

  // Asserts that multiple forwardFromRouter calls should result in
  // incrementing sequence number values.
  function testShouldIncrementSeqNumSuccess() public {
    uint64 seqNum = s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
    assertEq(seqNum, 1);
    seqNum = s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
    assertEq(seqNum, 2);
    seqNum = s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
    assertEq(seqNum, 3);
  }

  // Asserts that forwardFromRouter emits the correct event when sending
  // properly approved tokens.
  function testExactApproveSuccess() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    message.amounts = new uint256[](1);
    message.amounts[0] = 2**64;
    message.tokens = new IERC20[](1);
    message.tokens[0] = s_sourceTokens[0];

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(_messageToEvent(message, 1, 1));

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Assert that sending messages increments the nonce and the sequence numbers
  // on the onramp. Sending to a different receiver should start at 1 again.
  function testShouldIncrementReceiverNonceSuccess() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    CCIP.EVM2EVMSubscriptionMessage memory tollEvent = _messageToEvent(message, 1, 1);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(tollEvent);
    s_onRamp.forwardFromRouter(message, OWNER);

    message = _generateEmptyMessage();
    tollEvent = _messageToEvent(message, 2, 2);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(tollEvent);
    s_onRamp.forwardFromRouter(message, OWNER);

    message = _generateEmptyMessage();
    message.receiver = abi.encode(address(s_onRampRouter));
    tollEvent = _messageToEvent(message, 3, 1);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(tollEvent);
    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Reverts

  function testPausedReverts() public {
    changePrank(OWNER);
    s_onRamp.pause();
    vm.expectRevert("Pausable: paused");
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
  }

  function testUnhealthyReverts() public {
    s_afn.voteBad();
    vm.expectRevert(HealthChecker.BadAFNSignal.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
  }

  function testSenderNotAllowedReverts() public {
    changePrank(OWNER);
    s_onRamp.setAllowlistEnabled(true);

    vm.expectRevert(abi.encodeWithSelector(AllowListInterface.SenderNotAllowed.selector, STRANGER));
    changePrank(address(s_onRampRouter));
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), STRANGER);
  }

  function testUnsupportedTokenReverts() public {
    IERC20 wrongToken = IERC20(address(1));

    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    message.tokens = new IERC20[](1);
    message.tokens[0] = wrongToken;
    message.amounts = new uint256[](1);
    message.amounts[0] = 1;

    // We need to set the price of this new token to be able to reach
    // the proper revert point. This must be called by the owner.
    changePrank(OWNER);
    s_onRamp.setPrices(message.tokens, message.amounts);

    // Change back to the router
    changePrank(address(s_onRampRouter));

    vm.expectRevert(abi.encodeWithSelector(BaseOnRampInterface.UnsupportedToken.selector, wrongToken));
    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Asserts that forwardFromRouter reverts when it's not called by
  // the router
  function testMustBeCalledByRouterReverts() public {
    vm.stopPrank();
    vm.expectRevert(BaseOnRampInterface.MustBeCalledByRouter.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), OWNER);
  }

  // Asserts that forwardFromRouter reverts when the original sender
  // is not set by the router.
  function testRouterMustSetOriginalSenderReverts() public {
    vm.expectRevert(BaseOnRampInterface.RouterMustSetOriginalSender.selector);
    s_onRamp.forwardFromRouter(_generateEmptyMessage(), address(0));
  }

  // Asserts that forwardFromRouter reverts when the number of supplied tokens
  // is larger than the maxTokenLength.
  function testUnsupportedNumberOfTokensReverts() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    message.tokens = new IERC20[](MAX_TOKENS_LENGTH + 1);
    vm.expectRevert(BaseOnRampInterface.UnsupportedNumberOfTokens.selector);
    s_onRamp.forwardFromRouter(message, OWNER);
  }

  // Asserts that forwardFromRouter reverts when the data length is too long.
  function testMessageTooLargeReverts() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    message.data = new bytes(MAX_DATA_SIZE + 1);
    vm.expectRevert(
      abi.encodeWithSelector(
        BaseOnRampInterface.MessageTooLarge.selector,
        onRampConfig().maxDataSize,
        message.data.length
      )
    );

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  function testValueExceedsAllowedThresholdReverts() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    message.amounts = new uint256[](1);
    message.amounts[0] = 2**128;
    message.tokens = new IERC20[](1);
    message.tokens[0] = s_sourceTokens[0];

    s_sourceTokens[0].approve(address(s_onRamp), 2**128);

    vm.expectRevert(AggregateRateLimiterInterface.ValueExceedsAllowedThreshold.selector);

    s_onRamp.forwardFromRouter(message, OWNER);
  }

  function testPriceNotFoundForTokenReverts() public {
    CCIP.EVM2AnySubscriptionMessage memory message = _generateEmptyMessage();
    address fakeToken = address(1);
    message.amounts = new uint256[](1);
    message.tokens = new IERC20[](1);
    message.tokens[0] = IERC20(fakeToken);

    vm.expectRevert(abi.encodeWithSelector(AggregateRateLimiterInterface.PriceNotFoundForToken.selector, fakeToken));

    s_onRamp.forwardFromRouter(message, OWNER);
  }
}
