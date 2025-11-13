// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @dev Library that provides utility functions for the AAA protocol.
 */
library AAALib {
    /**
     * @dev Selects a subset of nodes from a pool based on a given random seed and the number of words needed.
     *
     * @param randomSeed A random seed for selection.
     * @param pool The array of available nodes to select from.
     * @param wordsNeeded The number of nodes to select.
     */
    function selectNodes(
        uint256 randomSeed,
        address[] memory pool,
        uint wordsNeeded
    ) internal pure returns (address[] memory selected) {
        require(pool.length >= wordsNeeded, "pool too small");

        selected = new address[](wordsNeeded);
        address[] memory temp = new address[](pool.length);
        for (uint256 i = 0; i < pool.length; i++) temp[i] = pool[i];
        uint256 remaining = pool.length;

        // Fisher–Yates shuffle using randomness derived from randomSeed
        for (uint256 i = 0; i < wordsNeeded; i++) {
            // mix in loop index to get fresh randomness per step
            randomSeed = uint256(keccak256(abi.encode(randomSeed, i)));
            uint256 idx = randomSeed % remaining;
            selected[i] = temp[idx];
            temp[idx] = temp[remaining - 1];
            remaining--;
        }
    }

    /**
     * @dev Selects a single node from a pool based on a given pid.
     *
     * @param pid User's Public Identity Data.
     * @param pool The array of available nodes to select from.
     */
    function selectNode(
        bytes32 pid,
        address[] memory pool
    ) internal pure returns (address selected) {
        require(pool.length > 0, "pool too small");

        uint idx = uint(keccak256(abi.encodePacked(pid))) % pool.length;
        selected = pool[idx];
    }

    /**
     * @dev Derives a symmetric key from the given words.
     *
     * @param words An array of encrypted words.
     */
    function deriveSymK(
        bytes[] memory words,
        bytes32 salt,
        bytes32 context
    ) internal pure returns (bytes32 key) {
        key = salt;
        for (uint i = 0; i < words.length; i++) {
            key = keccak256(
                abi.encodePacked(key, context, keccak256(words[i]))
            );
        }
        return key;
    }
}
