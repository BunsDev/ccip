// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";
import {GERouter} from "../../router/GERouter.sol";

contract GESRouterSetup is BaseTest {
  GERouter internal s_onRampRouter;

  function setUp() public virtual override {
    BaseTest.setUp();

    BaseOffRampInterface[] memory offRamps = new BaseOffRampInterface[](0);
    s_onRampRouter = new GERouter(offRamps);
  }
}
