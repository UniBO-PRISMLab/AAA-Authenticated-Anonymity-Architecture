// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./AAALib.sol";
import "./UIPRegistry.sol";

/// @title A simulator for trees
/// @dev AAA is a smart contract that manages the blockchain interactions for the AAA protocol.
/// @author Michele Dinelli
/// @notice Proof of concept implementation.
contract AAA is UIPRegistry {
    using AAALib for *;

    /// @dev Number of words needed to complete the seed phrase protocol.
    uint public immutable WORDS;

    /// @dev Redundancy factor, each word will be duplicated REDUNDANCY_FACTOR - 1 times.
    uint public immutable REDUNDANCY_FACTOR;

    /// @dev Tells if the contract is locked or not to avoid reentrancy.
    bool private _isLocked;

    /// @dev Mapping from PID to {Phrase} struct.
    mapping(bytes32 => Phrase) private phrases;

    /// @dev Mapping from SID to {SIDRecord} struct.
    mapping(bytes32 => SIDRecord) private sidRecords;

    /// @dev Mapping from PID to SID encrypted with user's public key.
    mapping(bytes32 => bytes) private pidToEncSid;

    /// @dev Mapping from SID to SAC codes.
    mapping(bytes32 => uint256[]) private sidToSac;

    /// @dev Determines the existence of SAC codes.
    mapping(bytes => bool) private sacCodes;

    /// @dev Mapping from a public key hash to its associated SAC code.
    mapping(bytes32 => bytes) private pkToSac;

    /// @dev Emitted to request word generation requested to a UIP node.
    event WordRequested(
        bytes32 indexed pid,
        address indexed node,
        bytes userPK
    );

    /// @dev Emitted when a word is submitted by a UIP node.
    event WordSubmitted(
        bytes32 indexed pid,
        address indexed node,
        bytes32 wordHash,
        uint index
    );

    /// @dev Emitted to request a redundant word from a UIP node.
    event RedundantWordRequested(
        bytes32 indexed pid,
        uint indexed index,
        bytes32 hashedWord,
        address indexed toNode
    );

    /// @dev Emitted when a redundant word is submitted by a UIP node.
    event RedundantWordSubmitted(
        bytes32 indexed pid,
        uint indexed index,
        address indexed node,
        bytes32 wordHash
    );

    /// @dev Emitted to request SID encryption from a UIP node.
    event SIDEncryptionRequested(
        bytes32 indexed pid,
        address indexed node,
        bytes sid,
        bytes userPK
    );

    /// @dev Emitted to request PID encryption from a UIP node.
    event PIDEncryptionRequested(
        bytes32 indexed pid,
        address indexed node,
        bytes32 symK,
        bytes32 sid
    );

    /// @dev Emitted when the seed phrase protocol starts.
    event SeedPhraseProtocolInitiated(bytes32 indexed pid);

    /// @dev Emitted when the seed phrase protocol is completed.
    event PhraseComplete(bytes32 indexed pid, bytes encSID);

    /// @dev Represents a seed phrase.
    struct Phrase {
        bool started;
        bytes pk;
        address encryptionResp;
        Word[] words;
        mapping(address => bool) hasSubmitted;
        mapping(uint => RedundantWord[]) redundantEncWords;
        mapping(address => mapping(uint => bool)) hasSubmittedRedundant;
    }

    /// @dev Represents a redundant word submitted by a node.
    struct RedundantWord {
        bytes word;
        bytes nodePK;
    }

    /// @dev Represents a word used in the seed phrase.
    struct Word {
        bytes word;
        uint index;
    }

    /// @dev Represents a SID record.
    struct SIDRecord {
        bytes encPID;
        bytes pk;
        bool exists;
    }

    modifier nonReentrant() {
        require(!_isLocked, "Reentrancy");
        _isLocked = true;
        _;
        _isLocked = false;
    }

    constructor(
        address[] memory nodes,
        uint words,
        uint redundancyFactor
    ) UIPRegistry(nodes) {
        require(nodes.length >= words, "too few nodes");
        require(redundancyFactor > 1, "invalid redundancy");
        WORDS = words;
        REDUNDANCY_FACTOR = redundancyFactor;
    }

    /**
     * @notice Starts the seed phrase generation protocol.
     * @dev Requires that the protocol has not already been started for the given pid.
     * @param pid User's PID.
     * @param pk User's Public Key as 2048 bit RSA PKCS#8 keys submitted as plain bytes.
     */
    function seedPhraseGenerationProtocol(
        bytes32 pid,
        bytes calldata pk
    ) external nonReentrant {
        Phrase storage p = phrases[pid];
        require(!p.started, "already started");

        p.started = true;
        p.pk = pk;

        address[] memory selected = AAALib.selectNodes(
            uint256(pid),
            nodeList,
            WORDS
        );

        for (uint i = 0; i < selected.length; i++) {
            selectedNodesByPID[pid].push(selected[i]);
            emit WordRequested(pid, selected[i], pk);
        }

        emit SeedPhraseProtocolInitiated(pid);
    }

    /**
     * @dev Submits an encrypted word for a given pid.
     * When the required number of words is reached, emits {SIDEncryptionRequested} event.
     *
     * Requirements:
     * - The encrypted word must not be empty.
     * - The phrase must have been initiated.
     * - The sender must be one of the selected nodes.
     * - The sender must not have already submitted their word.
     *
     * @param pid User's PID.
     * @param encryptedWord The word encrypted with the user's public key.
     */
    function submitEncryptedWord(
        bytes32 pid,
        bytes calldata encryptedWord
    ) external nonReentrant onlyUIPNode {
        Phrase storage p = phrases[pid];
        require(p.started, "not started");
        require(!p.hasSubmitted[msg.sender], "already submitted");
        require(encryptedWord.length > 0, "empty");

        bool isSelected;
        uint wordIndex;
        address[] memory selectedNodes = selectedNodesByPID[pid];
        for (uint i = 0; i < selectedNodes.length; i++) {
            if (selectedNodes[i] == msg.sender) {
                isSelected = true;
                wordIndex = i;
                break;
            }
        }
        require(isSelected, "not selected");

        p.words.push(Word({word: encryptedWord, index: wordIndex}));
        p.hasSubmitted[msg.sender] = true;
        emit WordSubmitted(
            pid,
            msg.sender,
            keccak256(encryptedWord),
            wordIndex
        );

        address[] memory redundantNodes = AAALib.selectNodes(
            uint256(keccak256(abi.encodePacked(pid, msg.sender))),
            nodeList,
            REDUNDANCY_FACTOR - 1
        );

        for (uint i = 0; i < redundantNodes.length; i++) {
            redundantNodesByPID[pid][wordIndex].push(redundantNodes[i]);
            emit RedundantWordRequested(
                pid,
                wordIndex,
                keccak256(encryptedWord),
                redundantNodes[i]
            );
        }

        if (p.words.length == WORDS) {
            bytes memory acc;
            bytes[] memory wordBytes = new bytes[](WORDS);
            for (uint i = 0; i < p.words.length; i++) {
                bytes32 h = keccak256(p.words[i].word);
                wordBytes[i] = p.words[i].word;
                acc = abi.encodePacked(acc, h);
            }

            address nodeAddress = AAALib.selectNode(pid, nodeList);
            p.encryptionResp = nodeAddress;

            emit SIDEncryptionRequested(
                pid,
                nodeAddress,
                abi.encodePacked(keccak256(acc)),
                p.pk
            );
        }
    }

    /**
     * @notice Submits a redundant encrypted word for a given pid.
     * @dev The sender must be a UIP node.
     * The phrase must have been initiated.
     * The sender must not have already submitted the redundant word for the given index.
     *
     * @param pid User's PID.
     * @param encryptedWord The redundant word encrypted with the node's public key.
     * @param wordIndex The index of the word being submitted.
     * @param nodePK The public key of the node submitting the redundant word.
     */
    function submitRedundantWord(
        bytes32 pid,
        bytes calldata encryptedWord,
        uint256 wordIndex,
        bytes calldata nodePK
    ) external nonReentrant onlyUIPNode {
        Phrase storage p = phrases[pid];
        require(p.started, "not started");
        require(
            !p.hasSubmittedRedundant[msg.sender][wordIndex],
            "already submitted"
        );
        require(encryptedWord.length > 0, "empty");

        bool wasRequested = false;
        address[] memory redundantNodes = redundantNodesByPID[pid][wordIndex];
        for (uint i = 0; i < redundantNodes.length; i++) {
            if (redundantNodes[i] == msg.sender) {
                wasRequested = true;
                break;
            }
        }
        require(wasRequested, "not selected");

        p.redundantEncWords[wordIndex].push(
            RedundantWord({word: encryptedWord, nodePK: nodePK})
        );
        p.hasSubmittedRedundant[msg.sender][wordIndex] = true;

        emit RedundantWordSubmitted(
            pid,
            wordIndex,
            msg.sender,
            keccak256(encryptedWord)
        );
    }

    /**
     * @dev Stores the encrypted SID for a given pid and marks the phrase as finalized.
     * Emits {PIDEncryptionRequested} event.
     *
     * Requirements:
     * - The sender must be a UIP node.
     * - The protocol must have been started for the given pid.
     * - The encrypted SID must not have been already stored.
     * - The sender must be the node selected to encrypt the SID.
     *
     * @param pid User's PID.
     * @param encSID The encrypted SID to be stored.
     */
    function submitEncryptedSID(
        bytes32 pid,
        bytes calldata encSID
    ) external nonReentrant onlyUIPNode {
        Phrase storage p = phrases[pid];
        require(p.started, "not started");
        require(pidToEncSid[pid].length == 0, "already stored");
        require(p.encryptionResp == msg.sender, "not selected");
        require(encSID.length > 0, "empty");
        pidToEncSid[pid] = encSID;

        bytes[] memory wordBytes = new bytes[](WORDS);
        bytes memory acc;
        for (uint i = 0; i < WORDS; i++) {
            wordBytes[i] = p.words[i].word;
            bytes32 h = keccak256(p.words[i].word);
            acc = abi.encodePacked(acc, h);
        }

        bytes32 symK = AAALib.deriveSymK(
            wordBytes,
            keccak256(abi.encodePacked(block.timestamp, msg.sender)),
            pid
        );

        emit PIDEncryptionRequested(
            pid,
            p.encryptionResp,
            symK,
            keccak256(acc)
        );
    }

    /**
     * @dev Stores the encrypted PID and symmetric key for a given pid.
     * Emits {PhraseComplete} event.
     *
     * Requirements:
     * - The sender must be a UIP node.
     * - The {SIDRecord} must not have been already stored.
     *
     * @param pid User's PID.
     * @param sid User's SID.
     * @param encPID The encrypted PID and symmetric key to be stored.
     */
    function submitEncryptedPID(
        bytes32 pid,
        bytes32 sid,
        bytes calldata encPID
    ) external nonReentrant onlyUIPNode {
        SIDRecord storage record = sidRecords[sid];
        require(!record.exists, "already stored");
        record.encPID = encPID;
        record.pk = phrases[pid].pk;
        record.exists = true;
        emit PhraseComplete(pid, pidToEncSid[pid]);
    }

    /**
     * @notice Submits a SAC record linking a public key to a SAC code.
     * @dev Requires:
     * - The public key hash must not be empty.
     * - The SAC code must be greater than zero.
     * - The public key hash must not have been already stored.
     * - The SAC code must exist.
     *
     * @param sac The SAC code to be linked.
     * @param pkHash The keccak256 hash of the user's public key.
     */
    function submitSACRecord(bytes calldata sac, bytes32 pkHash) external {
        require(sac.length > 0, "invalid sac");
        require(sacCodes[sac], "sac not found");
        pkToSac[pkHash] = sac;
    }

    /**
     * @dev Submits a SAC code.
     *
     * Requirements:
     * - The sender must be a UIP node.
     *
     * @param sac The SAC code to be submitted.
     */
    function submitSAC(bytes calldata sac) external nonReentrant onlyUIPNode {
        require(sac.length > 0, "invalid sac");
        require(!sacCodes[sac], "already stored");
        sacCodes[sac] = true;
    }

    /**
     * @dev Returns the encrypted SID for a given pid.
     *
     * @param pid User's PID.
     * @return The encrypted SID associated with the pid.
     */
    function getSID(bytes32 pid) external view returns (bytes memory) {
        return pidToEncSid[pid];
    }

    /**
     * @dev Returns the SID record for a given sid.
     * @param sid User's SID.
     * @return The encrypted PID.
     * @return The user's public key.
     */
    function getSIDRecord(
        bytes32 sid
    ) external view returns (bytes memory, bytes memory) {
        SIDRecord storage record = sidRecords[sid];
        require(record.exists, "not found");
        return (record.encPID, record.pk);
    }

    /**
     * @notice Returns the SAC code for a given public key hash.
     * @param pkHash User's public key hash.
     * @return The SAC code associated with the public key hash.
     */
    function getSACRecord(bytes32 pkHash) external view returns (bytes memory) {
        bytes memory sac = pkToSac[pkHash];
        require(sac.length > 0, "not found");
        return sac;
    }

    /**
     * @notice Checks if a SAC code exists.
     * @param sac The SAC code to check.
     * @return True if the SAC code exists, false otherwise.
     */
    function sacExists(bytes calldata sac) external view returns (bool) {
        return sacCodes[sac];
    }

    /**
     * @dev Returns the encrypted words, node public keys, and indexes for a given pid.
     * @param pid User's PID.
     * @return words The array of encrypted words.
     */
    function getWords(
        bytes32 pid
    ) external view returns (bytes[] memory words) {
        Word[] storage encWords = phrases[pid].words;
        uint len = encWords.length;
        words = new bytes[](len);
        for (uint i = 0; i < len; i++) {
            words[i] = encWords[i].word;
        }
    }

    /**
     * @dev Returns the phrase information for a given pid.
     * @param pid User's PID.
     * @return started Whether the phrase generation has started.
     * @return pk The user's public key.
     * @return encWords The array of encrypted words.
     */
    function getPhrase(
        bytes32 pid
    )
        external
        view
        returns (bool started, bytes memory pk, bytes[] memory encWords)
    {
        Phrase storage p = phrases[pid];
        bytes[] memory words = new bytes[](p.words.length);
        for (uint i = 0; i < p.words.length; i++) {
            words[i] = p.words[i].word;
        }
        return (p.started, p.pk, words);
    }

    /**
     * @dev Returns the redundant encrypted words and node public keys for a given pid and index.
     * @param pid User's PID.
     * @param index The index of the word.
     * @return words The array of redundant encrypted words.
     * @return nodePKs The array of node public keys associated with the redundant words.
     */
    function getRedundantWords(
        bytes32 pid,
        uint index
    ) external view returns (bytes[] memory words, bytes[] memory nodePKs) {
        RedundantWord[] storage redWords = phrases[pid].redundantEncWords[
            index
        ];
        uint len = redWords.length;
        words = new bytes[](len);
        nodePKs = new bytes[](len);
        for (uint i = 0; i < len; i++) {
            words[i] = redWords[i].word;
            nodePKs[i] = redWords[i].nodePK;
        }
    }
}
