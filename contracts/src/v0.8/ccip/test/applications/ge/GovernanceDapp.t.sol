// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../applications/GovernanceDapp.sol";
import "../../onRamp/ge/EVM2EVMGEOnRampSetup.t.sol";

// setup
contract GovernanceDappSetup is EVM2EVMGEOnRampSetup {
  GovernanceDapp s_governanceDapp;
  IERC20 s_feeToken;

  GovernanceDapp.FeeConfig s_feeConfig;
  GovernanceDapp.CrossChainClone s_crossChainClone;

  function setUp() public virtual override {
    EVM2EVMGEOnRampSetup.setUp();

    s_crossChainClone = GovernanceDapp.CrossChainClone({chainId: DEST_CHAIN_ID, contractAddress: address(1)});

    s_feeToken = IERC20(s_sourceTokens[0]);
    s_governanceDapp = new GovernanceDapp(s_sourceRouter, s_feeConfig, s_sourceFeeToken);
    s_governanceDapp.addClone(s_crossChainClone);

    uint256 fundingAmount = 1e18;

    // Fund the contract with LINK tokens
    IERC20(s_sourceFeeToken).transfer(address(s_governanceDapp), fundingAmount);

    // Approve the link tokens from the dapp
    changePrank(address(s_governanceDapp));
    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), fundingAmount);

    // Change back to te deployer of the contracts
    changePrank(OWNER);
  }
}

/// @notice #constructor
contract GovernanceDapp_constructor is GovernanceDappSetup {
  // Success
  function testSuccess() public {
    // typeAndVersion
    assertEq("GovernanceDapp 1.0.0", s_governanceDapp.typeAndVersion());
  }
}

/// @notice #voteForNewFeeConfig
contract GovernanceDapp_voteForNewFeeConfig is GovernanceDappSetup {
  using CCIP for CCIP.EVM2EVMGEMessage;

  event ConfigPropagated(uint256 chainId, address contractAddress);

  // Success
  function testSuccess() public {
    GovernanceDapp.FeeConfig memory feeConfig = GovernanceDapp.FeeConfig({feeAmount: 10000, changedAtBlock: 100});
    bytes memory data = abi.encode(feeConfig);
    CCIP.EVM2EVMGEMessage memory message = CCIP.EVM2EVMGEMessage({
      sequenceNumber: 1,
      sourceChainId: SOURCE_CHAIN_ID,
      sender: address(s_governanceDapp),
      receiver: s_crossChainClone.contractAddress,
      nonce: 1,
      data: data,
      tokensAndAmounts: new CCIP.EVMTokenAndAmount[](0),
      gasLimit: 3e5,
      strict: false,
      feeToken: s_sourceFeeToken,
      feeTokenAmount: 32400109, // todo
      messageId: ""
    });
    message.messageId = message._hash(s_metadataHash);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(message);

    vm.expectEmit(false, false, false, true);
    emit ConfigPropagated(s_crossChainClone.chainId, s_crossChainClone.contractAddress);

    s_governanceDapp.voteForNewFeeConfig(feeConfig);
  }
}

/// @notice #ccipReceive
contract GovernanceDapp_ccipReceive is GovernanceDappSetup {
  // Success

  function testSuccess() public {
    GovernanceDapp.FeeConfig memory feeConfig = GovernanceDapp.FeeConfig({feeAmount: 10000, changedAtBlock: 100});

    CCIP.Any2EVMMessage memory message = CCIP.Any2EVMMessage({
      sourceChainId: SOURCE_CHAIN_ID,
      sender: abi.encode(OWNER),
      data: abi.encode(feeConfig),
      destTokensAndAmounts: new CCIP.EVMTokenAndAmount[](0)
    });

    changePrank(address(s_sourceRouter));

    s_governanceDapp.ccipReceive(message);

    GovernanceDapp.FeeConfig memory newConfig = s_governanceDapp.getFeeConfig();
    assertEq(feeConfig.changedAtBlock, newConfig.changedAtBlock);
    assertEq(feeConfig.feeAmount, newConfig.feeAmount);
  }
  // Revert
}
