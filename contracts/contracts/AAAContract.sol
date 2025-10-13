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
     * @dev Represents a seed phrase
     *
     * `selectedNodes`: List of selected nodes for the phrase of size equals to WORDS_NEEDED
     *
     * `originalEncryptedWords`: Encrypted words submitted by selected nodes in their order
     *
     * `redundantEncryptedWords`: Redundant encrypted words keyed by index: mapping(index => address[] submissions)
     *
     * `hasSubmittedOriginal`: Tracks if a node has submitted its original word
     *
     * `hasSubmittedRedundant`: Tracks if a node has submitted a redundant word for a given index
     *
     * `encWordsByPID`: Mapping from PID to EncryptedWord struct
     *
     * `uipToEncryptSID`: Node selected to encrypt the SID
     *
     * `encSID`: encrypted with user’s public key
     *
     * `symK`: Computed when the phrase is complete
     *
     * `encPIDSymK`: Opaque encrypted payload: ENC(PID, symK) encrypted with user's PK
     *
     * `finalized`: Indicates if the phrase has been finalized
     *
     * `pk`: Public Key of the user
     */
    struct Phrase {
        address[] selectedNodes;
        bytes[] originalEncryptedWords;
        mapping(uint => bytes[]) redundantEncryptedWords;
        mapping(address => bool) hasSubmittedOriginal;
        mapping(address => mapping(uint => bool)) hasSubmittedRedundant;
        mapping(bytes32 => EncryptedWord[]) encWordsByPID;
        address uipToEncryptSID;
        bytes encSID;
        bytes32 symK;
        bytes encPIDSymK;
        bool finalized;
        bytes pk;
    }

    /// @dev Represents an encrypted word submitted by a node
    struct EncryptedWord {
        bytes word;
        bytes nodePK;
        uint index;
    }

    /// @dev Mapping from PID to Phrase
    mapping(bytes32 => Phrase) private phrases;

    /// @dev Word requested to a UIP node
    event WordRequestedToUIPNode(
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

    /// @dev Phrase completed
    event PhraseComplete(bytes32 indexed pid, bytes encSID);

    /// @dev SymK encrypted with user's public key and stored on-chain
    event SymKEncryptedStored(bytes32 indexed pid, bytes encPIDSymK);

    /// @dev Seed phrase generation protocol initiated
    event SeedPhraseProtocolInitiated(bytes32 indexed pid);

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
            pid,
            nodeList,
            WORDS_NEEDED
        );
        for (uint i = 0; i < selected.length; i++) {
            p.selectedNodes.push(selected[i]);
            emit WordRequestedToUIPNode(pid, selected[i], pk);
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
        require(!p.hasSubmittedOriginal[msg.sender], "already submitted");

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

        p.originalEncryptedWords.push(encryptedWord);
        p.encWordsByPID[pid].push(
            EncryptedWord({word: encryptedWord, nodePK: nodePK, index: index})
        );
        p.hasSubmittedOriginal[msg.sender] = true;
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

        if (p.originalEncryptedWords.length == WORDS_NEEDED) {
            bytes32 acc;
            for (uint k = 0; k < p.originalEncryptedWords.length; k++) {
                bytes32 h = keccak256(p.originalEncryptedWords[k]);
                acc = keccak256(abi.encodePacked(acc, h));
            }

            // TODO: derive a proper symK from sid
            p.symK = acc;
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
        emit PhraseComplete(pid, p.encSID);
    }

    /**
     * @dev Returns the selected nodes for a given pid.
     * This will be probably removed since the determinism of node selection
     * is not something that we want.
     *
     * @param pid User's PID.
     * @return Array of selected node addresses.
     */
    function getSelectedNodes(
        bytes32 pid
    ) external view returns (address[] memory) {
        return phrases[pid].selectedNodes;
    }

    /**
     * @dev Returns the original encrypted words for a given pid.
     *
     * @param pid User's PID.
     * @return Array of original encrypted words.
     */
    function getOriginalEncryptedWords(
        bytes32 pid
    ) external view returns (bytes[] memory) {
        return phrases[pid].originalEncryptedWords;
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
     * @dev Returns the SID for a given pid.
     *
     * @param pid User's PID.
     * @return The SID associated with the pid.
     */
    function getSID(bytes32 pid) external view returns (bytes memory) {
        return phrases[pid].encSID;
    }

    /**
     * @dev Returns the symmetric key (symK) for a given pid.
     *
     * @param pid User's PID.
     * @return The symmetric key associated with the pid.
     */
    function getSymK(bytes32 pid) external view returns (bytes32) {
        return phrases[pid].symK;
    }

    /**
     * @dev Returns the encrypted PID and symmetric key (encPIDSymK) for a given pid.
     *
     * @param pid User's PID.
     * @return The encrypted PID and symmetric key associated with the pid.
     */
    function getEncryptedPIDSymK(
        bytes32 pid
    ) external view returns (bytes memory) {
        return phrases[pid].encPIDSymK;
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
}
