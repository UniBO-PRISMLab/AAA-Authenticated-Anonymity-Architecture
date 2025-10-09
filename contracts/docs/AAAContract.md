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

_Represents a seed phrase

`selectedNodes`: List of selected nodes for the phrase of size equals to WORDS_NEEDED

`originalEncryptedWords`: Encrypted words submitted by selected nodes in their order

`redundantEncryptedWords`: Redundant encrypted words keyed by index: mapping(index => address[] submissions)

`hasSubmittedOriginal`: Tracks if a node has submitted its original word

`hasSubmittedRedundant`: Tracks if a node has submitted a redundant word for a given index

`sid`: Computed when the phrase is complete

`symK`: Computed when the phrase is complete

`encPIDSymK`: Opaque encrypted payload: ENC(PID, symK) encrypted with user's PK

`finalized`: Indicates if the phrase has been finalized

`pk`: Public Key of the user_

```solidity
struct Phrase {
  address[] selectedNodes;
  bytes[] originalEncryptedWords;
  mapping(uint256 => bytes[]) redundantEncryptedWords;
  mapping(address => bool) hasSubmittedOriginal;
  mapping(address => mapping(uint256 => bool)) hasSubmittedRedundant;
  bytes32 sid;
  bytes32 symK;
  bytes encPIDSymK;
  bool finalized;
  bytes pk;
}
```

### WordRequestedToUIPNode

```solidity
event WordRequestedToUIPNode(bytes32 pid, address node, bytes userPK)
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

### PhraseComplete

```solidity
event PhraseComplete(bytes32 pid, bytes32 sid, bytes32 symK)
```

_Phrase completed_

### SymKEncryptedStored

```solidity
event SymKEncryptedStored(bytes32 pid, bytes encPIDSymK)
```

_SymK encrypted with user's public key and stored on-chain_

### SeedPhraseProtocolInitiated

```solidity
event SeedPhraseProtocolInitiated(bytes32 pid)
```

_Seed phrase generation protocol initiated_

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
function submitEncryptedWord(bytes32 pid, bytes encryptedWord) external
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

### getSelectedNodes

```solidity
function getSelectedNodes(bytes32 pid) external view returns (address[])
```

_Returns the selected nodes for a given pid.
This will be probably removed since the determinism of node selection
is not something that we want._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | address[] | Array of selected node addresses. |

### getOriginalEncryptedWords

```solidity
function getOriginalEncryptedWords(bytes32 pid) external view returns (bytes[])
```

_Returns the original encrypted words for a given pid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bytes[] | Array of original encrypted words. |

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
function getSID(bytes32 pid) external view returns (bytes32)
```

_Returns the SID for a given pid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bytes32 | The SID associated with the pid. |

### getSymK

```solidity
function getSymK(bytes32 pid) external view returns (bytes32)
```

_Returns the symmetric key (symK) for a given pid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bytes32 | The symmetric key associated with the pid. |

### getEncryptedPIDSymK

```solidity
function getEncryptedPIDSymK(bytes32 pid) external view returns (bytes)
```

_Returns the encrypted PID and symmetric key (encPIDSymK) for a given pid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's PID. |

#### Return Values

| Name | Type | Description |
| ---- | ---- | ----------- |
| [0] | bytes | The encrypted PID and symmetric key associated with the pid. |

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

