// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../mocks/MockBlobVerifier.sol";
import "../helpers/BaseOffRampHelper.sol";
import "../TokenSetup.t.sol";

contract BaseOffRampSetup is TokenSetup {
  event OffRampConfigSet(BaseOffRampInterface.OffRampConfig config);

  BaseOffRampHelper s_offRamp;
  MockBlobVerifier s_mockBlobVerifier;

  function setUp() public virtual override {
    TokenSetup.setUp();

    s_mockBlobVerifier = new MockBlobVerifier();

    s_offRamp = new BaseOffRampHelper(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      offRampConfig(),
      s_mockBlobVerifier,
      ON_RAMP_ADDRESS,
      s_afn,
      s_sourceTokens,
      s_destPools,
      rateLimiterConfig(),
      TOKEN_LIMIT_ADMIN
    );

    s_offRamp.setPrices(s_destTokens, getTokenPrices());

    NativeTokenPool(address(s_destPools[0])).setOffRamp(s_offRamp, true);
    NativeTokenPool(address(s_destPools[1])).setOffRamp(s_offRamp, true);
  }

  function assertSameConfig(BaseOffRampInterface.OffRampConfig memory a, BaseOffRampInterface.OffRampConfig memory b)
    public
  {
    assertEq(a.executionDelaySeconds, b.executionDelaySeconds);
    assertEq(a.maxDataSize, b.maxDataSize);
    assertEq(a.maxTokensLength, b.maxTokensLength);
    assertEq(a.permissionLessExecutionThresholdSeconds, b.permissionLessExecutionThresholdSeconds);
  }
}

/// @notice #constructor
contract BaseOffRamp_constructor is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    // owner
    assertEq(OWNER, s_offRamp.owner());

    assertEq(address(s_mockBlobVerifier), address(s_offRamp.getBlobVerifier()));

    assertEq(SOURCE_CHAIN_ID, s_offRamp.i_sourceChainId());
    assertEq(DEST_CHAIN_ID, s_offRamp.i_chainId());

    assertSameConfig(offRampConfig(), s_offRamp.getConfig());
  }

  // Revert
  function testTokenConfigMismatchReverts() public {
    vm.expectRevert(TokenPoolRegistry.InvalidTokenPoolConfig.selector);

    PoolInterface[] memory pools = new PoolInterface[](1);

    IERC20[] memory wrongTokens = new IERC20[](5);
    s_offRamp = new BaseOffRampHelper(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      offRampConfig(),
      s_mockBlobVerifier,
      ON_RAMP_ADDRESS,
      s_afn,
      wrongTokens,
      pools,
      rateLimiterConfig(),
      TOKEN_LIMIT_ADMIN
    );
  }
}

/// @notice #getExecutionState
contract BaseOffRamp_getExecutionState is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    // setting the execution state is done with a helper function. This
    // is normally not exposed.
    s_offRamp.setExecutionState(1, CCIP.MessageExecutionState.FAILURE);
    s_offRamp.setExecutionState(10, CCIP.MessageExecutionState.IN_PROGRESS);
    s_offRamp.setExecutionState(33, CCIP.MessageExecutionState.UNTOUCHED);
    s_offRamp.setExecutionState(50, CCIP.MessageExecutionState.SUCCESS);

    assertEq(uint256(CCIP.MessageExecutionState.FAILURE), uint256(s_offRamp.getExecutionState(1)));
    assertEq(uint256(CCIP.MessageExecutionState.IN_PROGRESS), uint256(s_offRamp.getExecutionState(10)));
    assertEq(uint256(CCIP.MessageExecutionState.UNTOUCHED), uint256(s_offRamp.getExecutionState(33)));
    assertEq(uint256(CCIP.MessageExecutionState.SUCCESS), uint256(s_offRamp.getExecutionState(50)));
  }
}

/// @notice #getBlobVerifier
contract BaseOffRamp_getBlobVerifier is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    assertEq(address(s_mockBlobVerifier), address(s_offRamp.getBlobVerifier()));
  }
}

/// @notice #setBlobVerifier
contract BaseOffRamp_setBlobVerifier is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    assertEq(address(s_mockBlobVerifier), address(s_offRamp.getBlobVerifier()));

    MockBlobVerifier blobVerifier = new MockBlobVerifier();
    s_offRamp.setBlobVerifier(blobVerifier);

    assertEq(address(blobVerifier), address(s_offRamp.getBlobVerifier()));
  }
}

/// @notice #getConfig
contract BaseOffRamp_getConfig is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    assertSameConfig(offRampConfig(), s_offRamp.getConfig());
  }
}

/// @notice #setConfig
contract BaseOffRamp_setConfig is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    BaseOffRampInterface.OffRampConfig memory newConfig = generateNewConfig();

    vm.expectEmit(false, false, false, true);
    emit OffRampConfigSet(newConfig);

    s_offRamp.setConfig(newConfig);

    assertSameConfig(newConfig, s_offRamp.getConfig());
  }

  // Revert
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_offRamp.setConfig(offRampConfig());
  }

  function generateNewConfig() internal pure returns (BaseOffRampInterface.OffRampConfig memory) {
    return
      BaseOffRampInterface.OffRampConfig({
        executionDelaySeconds: 20,
        maxDataSize: 1,
        maxTokensLength: 15,
        permissionLessExecutionThresholdSeconds: 200
      });
  }
}

/// @notice #_releaseOrMintToken internal function
contract BaseOffRamp__releaseOrMintToken is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    uint256 startingBalance = s_destTokens[1].balanceOf(OWNER);
    uint256 amount = POOL_BALANCE / 2;
    s_offRamp.releaseOrMintToken(s_destPools[1], amount, OWNER);
    assertEq(startingBalance + amount, s_destTokens[1].balanceOf(OWNER));
  }

  // Revert
  function testExceedsPoolReverts() public {
    vm.expectRevert("ERC20: transfer amount exceeds balance");
    s_offRamp.releaseOrMintToken(s_destPools[1], POOL_BALANCE * 2, OWNER);
  }
}

/// @notice #_releaseOrMintTokens
contract BaseOffRamp__releaseOrMintTokens is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    uint256 startingBalance = s_destTokens[1].balanceOf(OWNER);

    PoolInterface[] memory pools = new PoolInterface[](2);
    pools[0] = s_destPools[1];
    pools[1] = s_destPools[1];

    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 100;
    amounts[1] = 50;

    s_offRamp.releaseOrMintTokens(pools, amounts, OWNER);
    assertEq(startingBalance + amounts[0] + amounts[1], s_destTokens[1].balanceOf(OWNER));
  }

  // Revert

  function testTokenAndAmountMisMatchReverts() public {
    uint256[] memory amounts = new uint256[](1);

    vm.expectRevert(BaseOffRampInterface.TokenAndAmountMisMatch.selector);
    s_offRamp.releaseOrMintTokens(s_destPools, amounts, OWNER);
  }
}

/// @notice #_verifyMessages
contract BaseOffRamp__verifyMessages is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    bytes32[] memory mockBytes = new bytes32[](5);
    // Since we use a mock blob verifier it should always return 1
    (uint256 timestamp, ) = s_offRamp.verifyMessages(mockBytes, mockBytes, 1, mockBytes, 1);
    assertEq(1, timestamp);
  }
}

/// @notice #_getPool
contract BaseOffRamp__getPool is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    address expectedPoolAddress = address(s_destPools[0]);
    address actualPoolAddress = address(s_offRamp.getPool(s_sourceTokens[0]));
    assertEq(expectedPoolAddress, actualPoolAddress);
  }

  // Revert
  function testUnsupportedTokenReverts() public {
    IERC20 wrongToken = IERC20(address(1));

    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.UnsupportedToken.selector, wrongToken));
    s_offRamp.getPool_helper(wrongToken);
  }
}

/// @notice #execute
contract BaseOffRamp_execute is BaseOffRampSetup {
  // Reverts
  function testReverts() public {
    vm.expectRevert();
    CCIP.ExecutionReport memory report;
    s_offRamp.execute(report, false);
  }
}
