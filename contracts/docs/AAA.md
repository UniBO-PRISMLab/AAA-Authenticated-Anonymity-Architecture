# Solidity API

## AAA

Proof of concept implementation.

_AAA is a smart contract that manages the blockchain interactions for the AAA protocol._

### WORDS

```solidity
uint256 WORDS
```

_Number of words needed to complete the seed phrase protocol._

### REDUNDANCY_FACTOR

```solidity
uint256 REDUNDANCY_FACTOR
```

_Redundancy factor, each word will be duplicated REDUNDANCY_FACTOR - 1 times._

### WordRequested

```solidity
event WordRequested(bytes32 pid, address node, bytes userPK)
```

_Emitted to request word generation requested to a UIP node._

### WordSubmitted

```solidity
event WordSubmitted(bytes32 pid, address node, bytes32 wordHash, uint256 index)
```

_Emitted when a word is submitted by a UIP node._

### RedundantWordRequested

```solidity
event RedundantWordRequested(bytes32 pid, uint256 index, bytes32 hashedWord, address toNode)
```

_Emitted to request a redundant word from a UIP node._

### RedundantWordSubmitted

```solidity
event RedundantWordSubmitted(bytes32 pid, uint256 index, address node, bytes32 wordHash)
```

_Emitted when a redundant word is submitted by a UIP node._

### SIDEncryptionRequested

```solidity
event SIDEncryptionRequested(bytes32 pid, address node, bytes sid, bytes userPK)
```

_Emitted to request SID encryption from a UIP node._

### PIDEncryptionRequested

```solidity
event PIDEncryptionRequested(bytes32 pid, address node, bytes32 symK, bytes32 sid)
```

_Emitted to request PID encryption from a UIP node._

### SeedPhraseProtocolInitiated

```solidity
event SeedPhraseProtocolInitiated(bytes32 pid)
```

_Emitted when the seed phrase protocol starts._

### PhraseComplete

```solidity
event PhraseComplete(bytes32 pid, bytes encSID)
```

_Emitted when the seed phrase protocol is completed._

### Phrase

_Represents a seed phrase._

```solidity
struct Phrase {
  bool started;
  bytes pk;
  address encryptionResp;
  struct AAA.Word[] words;
  mapping(address => bool) hasSubmitted;
  mapping(uint256 => struct AAA.RedundantWord[]) redundantEncWords;
  mapping(address => mapping(uint256 => bool)) hasSubmittedRedundant;
}
```

### RedundantWord

_Represents a redundant word submitted by a node._

```solidity
struct RedundantWord {
  bytes word;
  bytes nodePK;
}
```

### Word

_Represents a word used in the seed phrase._

```solidity
struct Word {
  bytes word;
  uint256 index;
}
```

### SIDRecord

_Represents a SID record._

```solidity
struct SIDRecord {
  bytes encPID;
  bytes pk;
  bool exists;
}
```

### nonReentrant

```solidity
modifier nonReentrant()
```

### constructor

```solidity
constructor(address[] nodes, uint256 words, uint256 redundancyFactor) public
```

### seedPhraseGenerationProtocol

```solidity
function seedPhraseGenerationProtocol(bytes32 pid, bytes pk) external
```

Starts the seed phrase generation protocol.

_Requires that the protocol has not already been started for the given pid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |
| pk | bytes | User's Public Key as 2048 bit RSA PKCS#8 keys submitted as plain bytes. |

### submitEncryptedWord

```solidity
function submitEncryptedWord(bytes32 pid, bytes encryptedWord) external
```

_Submits an encrypted word for a given pid.
When the required number of words is reached, emits {SIDEncryptionRequested} event.

Requirements:
- The encrypted word must not be empty.
- The phrase must have been initiated.
- The sender must be one of the selected nodes.
- The sender must not have already submitted their word._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |
| encryptedWord | bytes | The word encrypted with the user's public key. |

### submitRedundantWord

```solidity
function submitRedundantWord(bytes32 pid, bytes encryptedWord, uint256 wordIndex, bytes nodePK) external
```

Submits a redundant encrypted word for a given pid.

_The sender must be a UIP node.
The phrase must have been initiated.
The sender must not have already submitted the redundant word for the given index._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |
| encryptedWord | bytes | The redundant word encrypted with the node's public key. |
| wordIndex | uint256 | The index of the word being submitted. |
| nodePK | bytes | The public key of the node submitting the redundant word. |

### submitEncryptedSID

```solidity
function submitEncryptedSID(bytes32 pid, bytes encSID) external
```

_Stores the encrypted SID for a given pid and marks the phrase as finalized.
Emits {PIDEncryptionRequested} event.

Requirements:
- The sender must be a UIP node.
- The protocol must have been started for the given pid.
- The encrypted SID must not have been already stored.
- The sender must be the node selected to encrypt the SID._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |
| encSID | bytes | The encrypted SID to be stored. |

### submitEncryptedPID

```solidity
function submitEncryptedPID(bytes32 pid, bytes32 sid, bytes encPID) external
```

_Stores the encrypted PID and symmetric key for a given pid.
Emits {PhraseComplete} event.

Requirements:
- The sender must be a UIP node.
- The {SIDRecord} must not have been already stored._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |
| sid | bytes32 | User's SID. |
| encPID | bytes | The encrypted PID and symmetric key to be stored. |

### submitSACRecord

```solidity
function submitSACRecord(bytes sac, bytes32 pkHash) external
```

Submits a SAC record linking a public key to a SAC code.

_Requirements:
- The public key hash must not be empty.
- The SAC code must be greater than zero.
- The public key hash must not have been already stored.
- The SAC code must exist._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| sac | bytes | The SAC code to be linked. |
| pkHash | bytes32 | The keccak256 hash of the user's public key. |

### submitSAC

```solidity
function submitSAC(bytes sac) external
```

_Submits a SAC code.

Requirements:
- The sender must be a UIP node._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| sac | bytes | The SAC code to be submitted. |

### getSID

```solidity
function getSID(bytes32 pid) external view returns (bytes)
```

_Returns the encrypted SID for a given pid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bytes | The encrypted SID associated with the pid. |

### getSIDRecord

```solidity
function getSIDRecord(bytes32 sid) external view returns (bytes, bytes)
```

_Returns the SID record for a given sid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| sid | bytes32 | User's SID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bytes | The encrypted PID. |
| [1] | bytes | The user's public key. |

### getSACRecord

```solidity
function getSACRecord(bytes32 pkHash) external view returns (bytes)
```

Returns the SAC code for a given public key hash.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pkHash | bytes32 | User's public key hash. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bytes | The SAC code associated with the public key hash. |

### sacExists

```solidity
function sacExists(bytes sac) external view returns (bool)
```

Checks if a SAC code exists.

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| sac | bytes | The SAC code to check. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bool | True if the SAC code exists, false otherwise. |

### getWords

```solidity
function getWords(bytes32 pid) external view returns (bytes[] words)
```

_Returns the encrypted words, node public keys, and indexes for a given pid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| words | bytes[] | The array of encrypted words. |

### getPhrase

```solidity
function getPhrase(bytes32 pid) external view returns (bool started, bytes pk, bytes[] encWords)
```

_Returns the phrase information for a given pid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| started | bool | Whether the phrase generation has started. |
| pk | bytes | The user's public key. |
| encWords | bytes[] | The array of encrypted words. |

### getRedundantWords

```solidity
function getRedundantWords(bytes32 pid, uint256 index) external view returns (bytes[] words, bytes[] nodePKs)
```

_Returns the redundant encrypted words and node public keys for a given pid and index._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |
| index | uint256 | The index of the word. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| words | bytes[] | The array of redundant encrypted words. |
| nodePKs | bytes[] | The array of node public keys associated with the redundant words. |

