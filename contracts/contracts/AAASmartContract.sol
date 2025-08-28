// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.28;

contract AAASmartContract {

    uint256 public constant TOTAL_SEED_WORDS = 24;

    mapping(bytes32 => bool) public seedPhraseProtocolInitiated;
    mapping(bytes32 => string) public pidToUserPublicKey;
    mapping(bytes32 => bytes32) public pidToSid;
    mapping(bytes32 => mapping(uint256 => string)) public encryptedWordsForUser;
    mapping(bytes32 => mapping(uint256 => mapping(address => string))) public redundantEncryptedWords;
    mapping(bytes32 => uint256) public wordsReceivedCount;
    mapping(bytes32 => string) public encryptedSidForPid;
    mapping(bytes32 => string) public sidPidSymkAssociation;
    mapping(bytes32 => string) public sacToAnonymousAccountPk;

    mapping(address => bool) public isNIP;
    mapping(address => bool) public isUIP;

    address public owner;

    event SeedPhraseProtocolInitiated(bytes32 indexed pid, string userPk);
    event EncryptedWordFragmentSubmitted(bytes32 indexed pid, uint256 wordNumber, address indexed uipNode);
    event SidCreated(bytes32 indexed pid, bytes32 indexed sid);
    event SacRegistered(bytes32 indexed sac, string anonymousAccountPk);
    event PidSidAssociationRetrieved(bytes32 indexed sid, string encryptedPid);

    constructor() {
        owner = msg.sender;
        // Example: Add initial NIP/UIP registrations here or via admin functions.
        // isNIP[0xAbCdEf1234567890AbCdEf1234567890AbCdEf12] = true;
        // isUIP[0xFeDcBa9876543210FeDcBa9876543210FeDcBa98] = true;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "AAAC: Only owner can call this function.");
        _;
    }

    function registerNIP(address _nipAddress, bool _status) public onlyOwner {
        isNIP[_nipAddress] = _status;
    }

    function registerUIP(address _uipAddress, bool _status) public onlyOwner {
        isUIP[_uipAddress] = _status;
    }

    function requestSeedPhraseProtocol(bytes32 _pid, string memory _userPk) public {
        require(!seedPhraseProtocolInitiated[_pid], "AAAC: Seed phrase protocol already initiated for this PID.");
        require(bytes(_userPk).length > 0, "AAAC: User Public Key cannot be empty.");

        seedPhraseProtocolInitiated[_pid] = true;
        pidToUserPublicKey[_pid] = _userPk;
        // msg.sender might be a user's proxy wallet or an authorized frontend service.
        // pidToUserAddress[_pid] = msg.sender;

        emit SeedPhraseProtocolInitiated(_pid, _userPk);
    }

    function submitEncryptedWordFragment(bytes32 _pid, uint256 _wordNumber, string memory _encryptedWord) public {
        require(isUIP[msg.sender], "AAAC: Only registered UIP nodes can submit word fragments.");
        require(seedPhraseProtocolInitiated[_pid], "AAAC: Seed phrase protocol not initiated for this PID.");
        require(_wordNumber > 0 && _wordNumber <= TOTAL_SEED_WORDS, "AAAC: Invalid word number.");
        require(bytes(_encryptedWord).length > 0, "AAAC: Encrypted word cannot be empty.");

        // Store the encrypted word fragment.
        // Note: should check for duplicate submissions from the same UIP for the same word.
        encryptedWordsForUser[_pid][_wordNumber] = _encryptedWord;

        // Increment the count of words received for this PID.
        wordsReceivedCount[_pid]++;

        // If all expected words are received, trigger the finalization of SID creation.
        // This check might involve a more complex consensus mechanism
        // among UIPs to ensure data integrity and agreement before finalizing the SID.
        if (wordsReceivedCount[_pid] == TOTAL_SEED_WORDS) {
            _finalizeSeedPhraseAndSID(_pid);
        }

        emit EncryptedWordFragmentSubmitted(_pid, _wordNumber, msg.sender);
    }

    function submitRedundantEncryptedWord(bytes32 _pid, uint256 _originalWordNumber, string memory _encryptedWord, address _targetUipNodeAddress) public {
        require(isUIP[msg.sender], "AAAC: Only registered UIP nodes can submit redundant words.");
        require(isUIP[_targetUipNodeAddress], "AAAC: Target is not a registered UIP node.");
        require(_originalWordNumber > 0 && _originalWordNumber <= TOTAL_SEED_WORDS, "AAAC: Invalid word number.");
        require(bytes(_encryptedWord).length > 0, "AAAC: Encrypted word cannot be empty.");

        // Store the redundant encrypted word.
        redundantEncryptedWords[_pid][_originalWordNumber][_targetUipNodeAddress] = _encryptedWord;
    }

    function _finalizeSeedPhraseAndSID(bytes32 _pid) internal {
        bytes32 calculatedSid = keccak256(abi.encodePacked("SID_PLACEHOLDER_FOR_", _pid));

        string memory userPk = pidToUserPublicKey[_pid];
        string memory encryptedSid = string(abi.encodePacked("ENC_SID_WITH_", userPk));
        string memory pidWithSymk = string(abi.encodePacked("ENC_PID_WITH_SYMK_FOR_", _toHexString(calculatedSid)));

        pidToSid[_pid] = calculatedSid;
        encryptedSidForPid[_pid] = encryptedSid;
        sidPidSymkAssociation[calculatedSid] = pidWithSymk;

        emit SidCreated(_pid, calculatedSid);
    }


    function getEncryptedSeedPhraseAndSID(bytes32 _pid) public view returns (string[] memory encryptedWords, string memory encryptedSID, string memory symkEncryptedPidSidAssociation, string memory pkUserForSidEncryption) {
        require(seedPhraseProtocolInitiated[_pid], "AAAC: Seed phrase protocol not initiated for this PID.");
        require(pidToSid[_pid] != bytes32(0), "AAAC: SID not yet created for this PID. Please wait for protocol completion.");

        // Retrieve individual encrypted words
        string[] memory _encryptedWords = new string[](TOTAL_SEED_WORDS);
        for (uint256 i = 0; i < TOTAL_SEED_WORDS; i++) {
            _encryptedWords[i] = encryptedWordsForUser[_pid][i + 1]; // Word numbers are 1-indexed
        }

        return (
            _encryptedWords,
            encryptedSidForPid[_pid],
            sidPidSymkAssociation[pidToSid[_pid]], // Retrieve ENC(PID, SYMK) using the SID as key
            pidToUserPublicKey[_pid]
        );
    }

    function registerSACAssociation(bytes32 _sac, string memory _anonymousAccountPk) public {
        // In a real system, `msg.sender` must be a trusted NIP or an authorized entity.
        // require(isNIP[msg.sender], "AAAC: Only registered NIPs can register SAC associations.");
        require(bytes(sacToAnonymousAccountPk[_sac]).length == 0, "AAAC: SAC already registered with an anonymous account.");
        require(bytes(_anonymousAccountPk).length > 0, "AAAC: Anonymous account Public Key cannot be empty.");

        sacToAnonymousAccountPk[_sac] = _anonymousAccountPk;
        emit SacRegistered(_sac, _anonymousAccountPk);
    }

    function checkSACExistenceAndAssociation(bytes32 _sac, string memory _anonymousAccountPk) public view returns (bool) {
        // Compare the stored PK with the provided PK for the given SAC.
        // Using keccak256 for string comparison as direct string comparison can be gas inefficient.
        return keccak256(abi.encodePacked(sacToAnonymousAccountPk[_sac])) == keccak256(abi.encodePacked(_anonymousAccountPk)) &&
               bytes(sacToAnonymousAccountPk[_sac]).length > 0; // Ensure a mapping exists
    }

    function getPIDFromSID(bytes32 _sid) public view returns (string memory) {
        // This function requires very strict access control.
        // In a production system, `isAuthorizedForDeanonymization` would be a complex function
        // involving multi-signature checks, voting mechanisms, or specific role-based access.
        // For simplicity, it's currently commented out, but MUST be implemented.
        // require(isAuthorizedForDeanonymization(msg.sender), "AAAC: Caller not authorized for deanonymization.");
        require(bytes(sidPidSymkAssociation[_sid]).length > 0, "AAAC: PID-SYMK association not found for this SID.");

        return sidPidSymkAssociation[_sid];
    }

    function _toHexString(bytes32 data) internal pure returns (string memory) {
        bytes16 alphabet = 0x30313233343536373839616263646566; // "0123456789abcdef"
        bytes memory str = new bytes(2 + 64);
        str[0] = "0";
        str[1] = "x";
        for (uint256 i = 0; i < 32; i++) {
            uint8 b = uint8(data[i]);
            str[2 + 2 * i]     = bytes1(alphabet[b >> 4]);
            str[3 + 2 * i]     = bytes1(alphabet[b & 0x0f]);
        }
        return string(str);
    }
}