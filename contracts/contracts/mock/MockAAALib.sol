// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../AAALib.sol";

contract MockAAALib {
    using AAALib for *;

    function testSelectNodes(
        bytes32 pid,
        address[] memory pool,
        uint wordsNeeded
    ) external pure returns (address[] memory) {
        return AAALib.selectNodes(pid, pool, wordsNeeded);
    }

    function testSelectNode(
        bytes32 pid,
        address[] memory pool
    ) external pure returns (address) {
        return AAALib.selectNode(pid, pool);
    }

    function testDeriveSymK(
        bytes[] memory words
    ) external pure returns (bytes32) {
        bytes32 acc;
        for (uint i = 0; i < words.length; i++) {
            bytes32 h = keccak256(words[i]);
            acc = keccak256(abi.encodePacked(acc, h));
        }
        return acc;
    }
}
