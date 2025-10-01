// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @author @micheledinelli
 * @dev AAAContract is a smart contract that manages the blockchain interactions for the AAA protocol.
 */
contract AAAContract {

    /// @dev Number of words required to complete the seed phrase.
    uint private constant WORDS_NEEDED = 6;

    /// @dev The pool of available nodes to select from.
    address[] private nodePool;

    /// @dev Stores the encrypted words submitted by nodes for each pid
    mapping(bytes32 => bytes[]) public encryptedWords;

    /// @dev Stores the hash of the complete seed phrase for verification
    mapping(bytes32 => bytes32) public phraseHashes;

    /// @dev A word is requested from a UIP node
    event WordRequestedToUIPNode(bytes32 indexed pid, address indexed node, bytes publicKey);

    /// @dev The seed phrase is completed
    event PhraseComplete(bytes32 indexed pid);

    /// @dev The seed phrase generation protocol is initiated
    event SeedPhraseProtocoloInitiated(bytes32 indexed pid);
    
    constructor(address[] memory _nodes) {
        require(_nodes.length >= WORDS_NEEDED, "Not enough nodes in the pool");
        nodePool = _nodes;
    }

    /**
     * @dev Initiates the seed phrase generation protocol
     * by selecting nodes and emitting {WordRequestedToUIPNode} events.
     */
    function seedPhraseGenerationProtocol(bytes32 pid, bytes calldata publicKey) external {
        address[] memory selectedNodes = selectNodes(nodePool, pid, WORDS_NEEDED);
        for (uint i = 0; i < selectedNodes.length; i++) {
            emit WordRequestedToUIPNode(pid, selectedNodes[i], publicKey);
        }

        emit SeedPhraseProtocoloInitiated(pid);
    }
    
    /**
     * @dev Submits an encrypted word for a given pid.
     * When the required number of words is reached, emits {PhraseComplete} event.
     */
    function submitEncryptedWord(bytes32 pid, bytes calldata encryptedWord) external {
        require(encryptedWord.length > 0, "Encrypted word required");
        encryptedWords[pid].push(encryptedWord);

        if (encryptedWords[pid].length == WORDS_NEEDED) {
            string memory completePhrase = "";
            for (uint i = 0; i < encryptedWords[pid].length; i++) {
                completePhrase = string(abi.encodePacked(completePhrase, encryptedWords[pid][i]));
            }                
            emit PhraseComplete(pid);
        }
    }
}

function selectNodes(address[] memory nodePool, bytes32 pid, uint wordsNeeded) view returns (address[] memory selectedNodes) {
    selectedNodes = new address[](wordsNeeded);
    address[] memory tempPool = new address[](nodePool.length);

    for (uint i = 0; i < nodePool.length; i++) {
        tempPool[i] = nodePool[i];
    }
    
    uint remainingNodes = nodePool.length;
    
    for (uint i = 0; i < wordsNeeded; i++) {
        uint randomIndex = uint(keccak256(abi.encodePacked(
            pid,
            block.timestamp,
            block.prevrandao,
            block.number,
            msg.sender,
            i
        ))) % remainingNodes;
        
        selectedNodes[i] = tempPool[randomIndex];
        
        tempPool[randomIndex] = tempPool[remainingNodes - 1];
        remainingNodes--;
    }
    return selectedNodes;
}