// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../commitStore/CommitStore.sol";

contract CommitStoreHelper is CommitStore {
  constructor(
    uint64 chainId,
    uint64 sourceChainId,
    IAFN afn,
    CommitStoreConfig memory config
  ) CommitStore(chainId, sourceChainId, afn, config) {}

  /**
   * @dev Expose _report for tests
   */
  function report(bytes memory commitReport) external {
    _report(commitReport);
  }
}
