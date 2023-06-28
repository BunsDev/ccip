// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IARM} from "../../interfaces/IARM.sol";
import {ARM} from "../../ARM.sol";
import {OwnerIsCreator} from "./../../../shared/access/OwnerIsCreator.sol";

contract MockARM is IARM, OwnerIsCreator {
  bool private s_curse;
  ARM.VersionedConfig private s_versionedConfig;

  function isCursed() external view override returns (bool) {
    return s_curse;
  }

  function voteToCurse(bytes32) external {
    s_curse = true;
  }

  function ownerUnvoteToCurse(ARM.UnvoteToCurseRecord[] memory) external {
    s_curse = false;
  }

  function isBlessed(IARM.TaggedRoot calldata) external view override returns (bool) {
    return !s_curse;
  }

  function getConfigDetails() external view returns (uint32 version, uint32 blockNumber, ARM.Config memory config) {
    version = s_versionedConfig.configVersion;
    blockNumber = s_versionedConfig.blockNumber;
    config = s_versionedConfig.config;
  }
}
