// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @dev Library that provides utility functions for the AAA protocol.
 */
library AAALib {
    /**
     * @dev Selects a subset of nodes from a pool based on a given pid and the number of words needed.
     *
     * @param pid User's Public Identity Data.
     * @param pool The array of available nodes to select from.
     * @param wordsNeeded The number of nodes to select.
     */
    function selectNodes(
        bytes32 pid,
        address[] memory pool,
        uint wordsNeeded
    ) internal pure returns (address[] memory selected) {
        require(pool.length >= wordsNeeded, "pool too small");
        selected = new address[](wordsNeeded);
        address[] memory temp = new address[](pool.length);
        for (uint i = 0; i < pool.length; i++) temp[i] = pool[i];
        uint remaining = pool.length;

        // TODO: this is deterministic and we don't like it much (check for VRF)
        for (uint i = 0; i < wordsNeeded; i++) {
            uint idx = uint(keccak256(abi.encodePacked(pid, i))) % remaining;
            selected[i] = temp[idx];
            temp[idx] = temp[remaining - 1];
            remaining--;
        }
    }

    /**
     * @dev Derives a symmetric key from the given words.
     *
     * @param words An array of encrypted words.
     */
    function deriveSymK(
        bytes[] storage words
    ) internal view returns (bytes32 acc) {
        for (uint i = 0; i < words.length; i++) {
            bytes32 h = keccak256(words[i]);
            acc = keccak256(abi.encodePacked(acc, h));
        }
    }
}
