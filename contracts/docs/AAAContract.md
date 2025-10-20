# Solidity API

## AAAContract

_AAAContract is a smart contract that manages the blockchain interactions for the AAA protocol._

### WORDS_NEEDED

```solidity
uint256 WORDS_NEEDED
```

_Number of words needed to complete the seed phrase._

### REDUNDANCY_M

```solidity
uint256 REDUNDANCY_M
```

_Redundancy factor._

### nonReentrant

```solidity
modifier nonReentrant()
```

### Phrase

_Represents a seed phrase.

`selectedNodes`: List of selected nodes for the phrase.

`redundantEncryptedWords`: Redundant encrypted words keyed by index: mapping(index => address[] submissions).

`hasSubmitted`: Tracks if a node has submitted its original word.

`hasSubmittedRedundant`: Tracks if a node has submitted a redundant word for a given index.
`
`encWordsByPID`: Mapping from PID to EncryptedWord struct.

`uipToEncryptSID`: address of the node selected to encrypt the SID with the user's public key.

`encSID`: SID encrypted with user’s public key.

`finalized`: Indicates if the phrase has been finalized._

```solidity
struct Phrase {
  address[] selectedNodes;
  mapping(uint256 => bytes[]) redundantEncryptedWords;
  mapping(address => bool) hasSubmitted;
  mapping(address => mapping(uint256 => bool)) hasSubmittedRedundant;
  mapping(bytes32 => struct AAAContract.EncryptedWord[]) encWordsByPID;
  address uipToEncryptSID;
  bytes encSID;
  bool finalized;
  bytes pk;
}
```

### EncryptedWord

_Represents an encrypted word submitted by a node_

```solidity
struct EncryptedWord {
  bytes word;
  bytes nodePK;
  uint256 index;
}
```

### SIDRecord

_Represents a SID record_

```solidity
struct SIDRecord {
  bytes encPID;
  bytes pk;
  bool exists;
}
```

### WordRequested

```solidity
event WordRequested(bytes32 pid, address node, bytes userPK)
```

_Word requested to a UIP node_

### WordSubmitted

```solidity
event WordSubmitted(bytes32 pid, address node, uint256 index, bytes32 wordHash)
```

_Word submitted by a UIP node_

### RedundancyRequested

```solidity
event RedundancyRequested(bytes32 pid, uint256 index, address fromNode, address toNode)
```

_Redundancy requested from a UIP node_

### RedundantWordSubmitted

```solidity
event RedundantWordSubmitted(bytes32 pid, uint256 index, address node, bytes32 wordHash)
```

_Redundant word submitted by a UIP node_

### SIDEncryptionRequested

```solidity
event SIDEncryptionRequested(bytes32 pid, address node, bytes sid, bytes userPK)
```

_Emitted to request SID encryption from a UIP node_

### PIDEncryptionRequested

```solidity
event PIDEncryptionRequested(bytes32 pid, address node, bytes32 symK)
```

_Emitted to request PID encryption from a UIP node_

### SeedPhraseProtocolInitiated

```solidity
event SeedPhraseProtocolInitiated(bytes32 pid)
```

_Seed phrase generation protocol initiated_

### PhraseComplete

```solidity
event PhraseComplete(bytes32 pid, bytes encSID)
```

_Phrase completed_

### constructor

```solidity
constructor(address[] nodes, uint256 wordsNeeded, uint256 redundancyM) public
```

### seedPhraseGenerationProtocol

```solidity
function seedPhraseGenerationProtocol(bytes32 pid, bytes pk) external
```

_Initiates the seed phrase generation protocol.

Requirements:
- The phrase must not be finalized.
- The phrase must not have been started._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |
| pk | bytes | User's Public Key. |

### submitEncryptedWord

```solidity
function submitEncryptedWord(bytes32 pid, bytes encryptedWord, bytes nodePK) external
```

_Submits an encrypted word for a given pid.
When the required number of words is reached, emits {PhraseComplete} event.

Requirements:
- The encrypted word must not be empty.
- The phrase must have been initiated.
- The phrase must not be finalized.
- The sender must be one of the selected nodes.
- The sender must not have already submitted their original word._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |
| encryptedWord | bytes | The encrypted word submitted by the node. |
| nodePK | bytes | The public key of the node submitting the word. |

### submitRedundantEncryptedWord

```solidity
function submitRedundantEncryptedWord(bytes32 pid, uint256 index, bytes encryptedWordForTarget) external
```

_Submits a redundant encrypted word for a given pid and index.

Requirements:
- The phrase must be finalized.
- The sender must be a UIP node.
- The sender must not have already submitted a redundant word for the given index._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |
| index | uint256 | Index of the word for which redundancy is being submitted. |
| encryptedWordForTarget | bytes | The redundant encrypted word submitted by the node. |

### storeEncryptedSID

```solidity
function storeEncryptedSID(bytes32 pid, bytes encSID) external
```

_Stores the encrypted SID for a given pid and marks the phrase as finalized.
Emits {PhraseComplete} event.

Requirements:
- The phrase must be finalized.
- The sender must be a UIP node.
- The encrypted SID must not have been already stored.
- The sender must be the node selected to encrypt the SID._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |
| encSID | bytes | The encrypted SID to be stored. |

### storeEncryptedPID

```solidity
function storeEncryptedPID(bytes32 pid, bytes encPID) external
```

_Stores the encrypted PID and symmetric key for a given pid.

Requirements:
- The sender must be a UIP node.
- The encrypted PID and symmetric key must not have been already stored._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |
| encPID | bytes | The encrypted PID and symmetric key to be stored. |

### getRedundantEncryptedWords

```solidity
function getRedundantEncryptedWords(bytes32 pid, uint256 index) external view returns (bytes[])
```

_Returns the redundant encrypted words for a given pid and index._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |
| index | uint256 | Index of the word. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bytes[] | Array of redundant encrypted words for the specified index. |

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

### getSelectedNodes

```solidity
function getSelectedNodes(bytes32 pid) external view returns (address[])
```

_Returns the selected nodes for a given pid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | address[] | Array of selected node addresses for the specified pid. |

### getUserPK

```solidity
function getUserPK(bytes32 pid) external view returns (bytes)
```

_Returns the public key (pk) for a given pid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bytes | The public key associated with the pid. |

### getWords

```solidity
function getWords(bytes32 pid) external view returns (bytes[])
```

_Returns the encrypted words for a given pid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bytes[] | Array of encrypted words associated with the pid. |

