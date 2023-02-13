// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";
import {Router} from "../../router/Router.sol";
import "../fees/WETH9.sol";

contract RouterSetup is BaseTest {
  Router internal s_sourceRouter;

  function setUp() public virtual override {
    BaseTest.setUp();

    address[] memory offRamps = new address[](0);
    WETH9 weth = new WETH9();
    s_sourceRouter = new Router(offRamps, address(weth));
  }
}
