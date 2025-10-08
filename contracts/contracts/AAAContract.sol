// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @author @micheledinelli
 * @dev AAAContract is a smart contract that manages the blockchain interactions for the AAA protocol.
 */
contract AAAContract {

    /// @dev Number of words required to complete the seed phrase
    uint private constant WORDS_NEEDED = 6;

    /// @dev The pool of available nodes to select from
    address[] private nodePool;

    /// @dev Stores the encrypted words submitted by nodes for each pid
    mapping(bytes32 => bytes[]) public encryptedWords;

    /// @dev Stores the hash of the complete seed phrase for verification
    mapping(bytes32 => bytes32) public phraseHashes;

    /// @dev Tracks if a node has already submitted a word for a given pid
    mapping(bytes32 => mapping(address => bool)) public hasSubmittedWord;

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
     * @dev Modifier to restrict access to only authorized UIP nodes.
     */
    modifier onlyUIPNode() {
        bool authorized = false;
        for (uint i = 0; i < nodePool.length; i++) {
            if (nodePool[i] == msg.sender) {
                authorized = true;
                break;
            }
        }
        require(authorized, "Not an authorized UIP node");
        _;
    }

    /**
     * @dev Initiates the seed phrase generation protocol
     * by selecting nodes and emitting {WordRequestedToUIPNode} events.
     */
    function seedPhraseGenerationProtocol(bytes32 pid, bytes calldata publicKey) external {
        require(phraseHashes[pid] == 0, "Phrase already completed");
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
    function submitEncryptedWord(bytes32 pid, bytes calldata encryptedWord) external onlyUIPNode {
        require(encryptedWord.length > 0, "Encrypted word required");
        require(phraseHashes[pid] == 0, "Phrase already completed");
        require(!hasSubmittedWord[pid][msg.sender], "Node already submitted a word");

        encryptedWords[pid].push(encryptedWord);
        hasSubmittedWord[pid][msg.sender] = true;

        if (encryptedWords[pid].length == WORDS_NEEDED) {
            bytes memory full;
            for (uint i = 0; i < encryptedWords[pid].length; i++) {
                full = bytes.concat(full, encryptedWords[pid][i]);
            }
            phraseHashes[pid] = keccak256(full);
            emit PhraseComplete(pid);
        }
    }

    /**
     * @dev Returns the encrypted words for a given PID.
     *      The user can decrypt these off-chain using their private key.
     */
    function getEncryptedWords(bytes32 pid) external view returns (bytes[] memory) {
        require(encryptedWords[pid].length > 0, "No words found for this PID");
        return encryptedWords[pid];
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