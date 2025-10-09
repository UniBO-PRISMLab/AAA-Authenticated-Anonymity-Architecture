// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../UIPRegistry.sol";

contract MockUIPRegistry is UIPRegistry {
    constructor(address[] memory initialNodes) UIPRegistry(initialNodes) {}

    function onlyOwnerFn() external onlyOwner {}
    function onlyNodeFn() external onlyUIPNode {}
}
