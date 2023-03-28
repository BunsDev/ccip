// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IAFN} from "../../interfaces/IAFN.sol";
import {AFN} from "../../AFN.sol";

contract MockAFN is IAFN {
  bool private s_curse;

  function badSignalReceived() external view override returns (bool) {
    return s_curse;
  }

  function voteToCurse(bytes32) external {
    s_curse = true;
  }

  function ownerUnvoteToCurse(AFN.UnvoteToCurseRecord[] memory) external {
    s_curse = false;
  }

  function isBlessed(bytes32) external view override returns (bool) {
    return !s_curse;
  }
}
