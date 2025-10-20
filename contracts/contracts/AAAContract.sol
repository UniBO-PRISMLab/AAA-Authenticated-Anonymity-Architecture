// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./AAALib.sol";
import "./UIPRegistry.sol";

/**
 * @dev AAAContract is a smart contract that manages the blockchain interactions for the AAA protocol.
 */

contract AAAContract is UIPRegistry {
    using AAALib for *;

    /// @dev Number of words needed to complete the seed phrase.
    uint public immutable WORDS_NEEDED;

    /// @dev Redundancy factor.
    uint public immutable REDUNDANCY_M;

    /// @dev If the contract is locked or not.
    bool private _isLocked;

    modifier nonReentrant() {
        require(!_isLocked, "Reentrancy");
        _isLocked = true;
        _;
        _isLocked = false;
    }

    /**
     * @dev Represents a seed phrase.
     *
     * `selectedNodes`: List of selected nodes for the phrase.
     *
     * `redundantEncryptedWords`: Redundant encrypted words keyed by index: mapping(index => address[] submissions).
     *
     * `hasSubmitted`: Tracks if a node has submitted its original word.
     *
     * `hasSubmittedRedundant`: Tracks if a node has submitted a redundant word for a given index.
     *`
     * `encWordsByPID`: Mapping from PID to EncryptedWord struct.
     *
     * `uipToEncryptSID`: address of the node selected to encrypt the SID with the user's public key.
     *
     * `encSID`: SID encrypted with user’s public key.
     *
     * `finalized`: Indicates if the phrase has been finalized.
     */
    struct Phrase {
        address[] selectedNodes;
        mapping(uint => bytes[]) redundantEncryptedWords;
        mapping(address => bool) hasSubmitted;
        mapping(address => mapping(uint => bool)) hasSubmittedRedundant;
        mapping(bytes32 => EncryptedWord[]) encWordsByPID;
        address uipToEncryptSID;
        bytes encSID;
        bool finalized;
        bytes pk;
    }

    /// @dev Represents an encrypted word submitted by a node
    struct EncryptedWord {
        bytes word;
        bytes nodePK;
        uint index;
    }

    /// @dev Represents a SID record
    struct SIDRecord {
        bytes encPID;
        bytes pk;
        bool exists;
    }

    /// @dev Mapping from PID to Phrase
    mapping(bytes32 => Phrase) private phrases;

    /// @dev Mapping from SID to PID encrypted with symK and PK
    mapping(bytes32 => SIDRecord) private sidRecords;

    /// @dev Represents a SID record

    /// @dev Word requested to a UIP node
    event WordRequested(
        bytes32 indexed pid,
        address indexed node,
        bytes userPK
    );

    /// @dev Word submitted by a UIP node
    event WordSubmitted(
        bytes32 indexed pid,
        address indexed node,
        uint indexed index,
        bytes32 wordHash
    );

    /// @dev Redundancy requested from a UIP node
    event RedundancyRequested(
        bytes32 indexed pid,
        uint indexed index,
        address indexed fromNode,
        address toNode
    );

    /// @dev Redundant word submitted by a UIP node
    event RedundantWordSubmitted(
        bytes32 indexed pid,
        uint indexed index,
        address indexed node,
        bytes32 wordHash
    );

    /// @dev Emitted to request SID encryption from a UIP node
    event SIDEncryptionRequested(
        bytes32 indexed pid,
        address indexed node,
        bytes sid,
        bytes userPK
    );

    /// @dev Emitted to request PID encryption from a UIP node
    event PIDEncryptionRequested(
        bytes32 indexed pid,
        address indexed node,
        bytes32 symK
    );

    /// @dev Seed phrase generation protocol initiated
    event SeedPhraseProtocolInitiated(bytes32 indexed pid);

    /// @dev Phrase completed
    event PhraseComplete(bytes32 indexed pid, bytes encSID);

    constructor(
        address[] memory nodes,
        uint wordsNeeded,
        uint redundancyM
    ) UIPRegistry(nodes) {
        require(nodes.length >= wordsNeeded, "too few nodes");
        WORDS_NEEDED = wordsNeeded;
        REDUNDANCY_M = redundancyM;
    }

    /**
     * @dev Initiates the seed phrase generation protocol.
     *
     * Requirements:
     * - The phrase must not be finalized.
     * - The phrase must not have been started.
     *
     * @param pid User's PID.
     * @param pk User's Public Key.
     */
    function seedPhraseGenerationProtocol(
        bytes32 pid,
        bytes calldata pk
    ) external nonReentrant {
        Phrase storage p = phrases[pid];
        require(!p.finalized, "finalized");
        require(p.selectedNodes.length == 0, "already started");

        address[] memory selected = AAALib.selectNodes(
            132456,
            nodeList,
            WORDS_NEEDED
        );
        for (uint i = 0; i < selected.length; i++) {
            p.selectedNodes.push(selected[i]);
            emit WordRequested(pid, selected[i], pk);
        }
        p.pk = pk;
        emit SeedPhraseProtocolInitiated(pid);
    }

    /**
     * @dev Submits an encrypted word for a given pid.
     * When the required number of words is reached, emits {PhraseComplete} event.
     *
     * Requirements:
     * - The encrypted word must not be empty.
     * - The phrase must have been initiated.
     * - The phrase must not be finalized.
     * - The sender must be one of the selected nodes.
     * - The sender must not have already submitted their original word.
     *
     * @param pid User's PID.
     * @param encryptedWord The encrypted word submitted by the node.
     * @param nodePK The public key of the node submitting the word.
     */
    function submitEncryptedWord(
        bytes32 pid,
        bytes calldata encryptedWord,
        bytes calldata nodePK
    ) external nonReentrant onlyUIPNode {
        require(encryptedWord.length > 0, "empty");

        Phrase storage p = phrases[pid];
        require(p.selectedNodes.length == WORDS_NEEDED, "not initiated");
        require(!p.finalized, "done");
        require(!p.hasSubmitted[msg.sender], "already submitted");

        bool isSelected;
        uint index;
        for (uint i = 0; i < p.selectedNodes.length; i++) {
            if (p.selectedNodes[i] == msg.sender) {
                isSelected = true;
                index = i;
                break;
            }
        }
        require(isSelected, "not selected");

        p.encWordsByPID[pid].push(
            EncryptedWord({word: encryptedWord, nodePK: nodePK, index: index})
        );
        p.hasSubmitted[msg.sender] = true;

        emit WordSubmitted(pid, msg.sender, index, keccak256(encryptedWord));

        if (REDUNDANCY_M > 1) {
            address[] memory pool = nodeList;
            uint remaining = pool.length;
            address[] memory temp = new address[](pool.length);
            for (uint i = 0; i < pool.length; i++) temp[i] = pool[i];
            for (uint i = 0; i < remaining; i++) {
                if (temp[i] == msg.sender) {
                    temp[i] = temp[remaining - 1];
                    remaining--;
                    break;
                }
            }
            for (uint j = 0; j < REDUNDANCY_M - 1 && remaining > 0; j++) {
                uint targetIdx = uint(
                    keccak256(abi.encodePacked(pid, index, j))
                ) % remaining;
                address target = temp[targetIdx];
                emit RedundancyRequested(pid, index, msg.sender, target);
                temp[targetIdx] = temp[remaining - 1];
                remaining--;
            }
        }

        if (p.encWordsByPID[pid].length == WORDS_NEEDED) {
            bytes32 acc;
            for (uint k = 0; k < p.encWordsByPID[pid].length; k++) {
                bytes32 h = keccak256(p.encWordsByPID[pid][k].word);
                acc = keccak256(abi.encodePacked(acc, h));
            }

            bytes[] memory wordBytes = new bytes[](WORDS_NEEDED);
            for (uint i = 0; i < WORDS_NEEDED; i++) {
                wordBytes[i] = p.encWordsByPID[pid][i].word;
            }

            p.finalized = true;

            address nodeAddress = AAALib.selectNode(pid, nodeList);
            p.uipToEncryptSID = nodeAddress;

            emit SIDEncryptionRequested(
                pid,
                nodeAddress,
                abi.encodePacked(acc),
                p.pk
            );
        }
    }

    /**
     * @dev Submits a redundant encrypted word for a given pid and index.
     *
     * Requirements:
     * - The phrase must be finalized.
     * - The sender must be a UIP node.
     * - The sender must not have already submitted a redundant word for the given index.
     *
     * @param pid User's PID.
     * @param index Index of the word for which redundancy is being submitted.
     * @param encryptedWordForTarget The redundant encrypted word submitted by the node.
     */
    function submitRedundantEncryptedWord(
        bytes32 pid,
        uint index,
        bytes calldata encryptedWordForTarget
    ) external nonReentrant onlyUIPNode {
        Phrase storage p = phrases[pid];
        require(
            !p.hasSubmittedRedundant[msg.sender][index],
            "already submitted"
        );
        p.redundantEncryptedWords[index].push(encryptedWordForTarget);
        p.hasSubmittedRedundant[msg.sender][index] = true;
        emit RedundantWordSubmitted(
            pid,
            index,
            msg.sender,
            keccak256(encryptedWordForTarget)
        );
    }

    /**
     * @dev Stores the encrypted SID for a given pid and marks the phrase as finalized.
     * Emits {PhraseComplete} event.
     *
     * Requirements:
     * - The phrase must be finalized.
     * - The sender must be a UIP node.
     * - The encrypted SID must not have been already stored.
     * - The sender must be the node selected to encrypt the SID.
     *
     * @param pid User's PID.
     * @param encSID The encrypted SID to be stored.
     */
    function storeEncryptedSID(
        bytes32 pid,
        bytes calldata encSID
    ) external nonReentrant onlyUIPNode {
        Phrase storage p = phrases[pid];
        require(p.encSID.length == 0, "already stored");
        require(p.finalized, "not finalized");
        require(p.uipToEncryptSID == msg.sender, "not selected");
        p.encSID = encSID;

        bytes[] memory wordBytes = new bytes[](WORDS_NEEDED);
        for (uint i = 0; i < WORDS_NEEDED; i++) {
            wordBytes[i] = p.encWordsByPID[pid][i].word;
        }

        bytes32 symK = AAALib.deriveSymK(wordBytes);

        emit PIDEncryptionRequested(pid, msg.sender, symK);
    }

    /**
     * @dev Stores the encrypted PID and symmetric key for a given pid.
     *
     * Requirements:
     * - The sender must be a UIP node.
     * - The encrypted PID and symmetric key must not have been already stored.
     *
     * @param pid User's PID.
     * @param encPID The encrypted PID and symmetric key to be stored.
     */
    function storeEncryptedPID(
        bytes32 pid,
        bytes calldata encPID
    ) external nonReentrant onlyUIPNode {
        SIDRecord storage record = sidRecords[pid];
        require(!record.exists, "already stored");
        record.encPID = encPID;
        record.pk = phrases[pid].pk;
        emit PhraseComplete(pid, phrases[pid].encSID);
    }

    /**
     * @dev Returns the redundant encrypted words for a given pid and index.
     *
     * @param pid User's PID.
     * @param index Index of the word.
     * @return Array of redundant encrypted words for the specified index.
     */
    function getRedundantEncryptedWords(
        bytes32 pid,
        uint index
    ) external view returns (bytes[] memory) {
        return phrases[pid].redundantEncryptedWords[index];
    }

    /**
     * @dev Returns the encrypted SID for a given pid.
     *
     * @param pid User's PID.
     * @return The encrypted SID associated with the pid.
     */
    function getSID(bytes32 pid) external view returns (bytes memory) {
        return phrases[pid].encSID;
    }

    /**
     * @dev Returns the selected nodes for a given pid.
     *
     * @param pid User's PID.
     * @return Array of selected node addresses for the specified pid.
     */
    function getSelectedNodes(
        bytes32 pid
    ) external view returns (address[] memory) {
        return phrases[pid].selectedNodes;
    }

    /**
     * @dev Returns the public key (pk) for a given pid.
     *
     * @param pid User's PID.
     * @return The public key associated with the pid.
     */
    function getUserPK(bytes32 pid) external view returns (bytes memory) {
        return phrases[pid].pk;
    }

    /**
     * @dev Returns the encrypted words for a given pid.
     *
     * @param pid User's PID.
     * @return Array of encrypted words associated with the pid.
     */
    function getWords(bytes32 pid) external view returns (bytes[] memory) {
        EncryptedWord[] storage encWords = phrases[pid].encWordsByPID[pid];
        bytes[] memory words = new bytes[](encWords.length);
        for (uint i = 0; i < encWords.length; i++) {
            words[i] = encWords[i].word;
        }
        return words;
    }
}
