# Solidity API

## AAA

_AAA is a smart contract that manages the blockchain interactions for the AAA protocol._

### WORDS

```solidity
uint256 WORDS
```

_Number of words needed to complete the seed phrase._

### REDUNDANCY_FACTOR

```solidity
uint256 REDUNDANCY_FACTOR
```

_Redundancy factor._

### WordRequested

```solidity
event WordRequested(bytes32 pid, address node, bytes userPK)
```

_Emitted to request word generation requested to a UIP node._

### WordSubmitted

```solidity
event WordSubmitted(bytes32 pid, address node, bytes32 wordHash)
```

_Emitted when a word is submitted by a UIP node._

### RedundantWordRequested

```solidity
event RedundantWordRequested(bytes32 pid, uint256 index, address fromNode, address toNode)
```

_Redundancy requested from a UIP node._

### RedundantWordSubmitted

```solidity
event RedundantWordSubmitted(bytes32 pid, uint256 index, address node, bytes32 wordHash)
```

_Redundant word submitted by a UIP node_

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

_Seed phrase generation protocol initiated._

### PhraseComplete

```solidity
event PhraseComplete(bytes32 pid, bytes encSID)
```

_Seed phrase protocol is completed._

### Phrase

_Represents a seed phrase._

```solidity
struct Phrase {
  bool started;
  bytes pk;
  address encryptionResp;
  mapping(address => bool) hasSubmitted;
  mapping(bytes32 => struct AAA.Word[]) words;
  mapping(address => bool) hasSubmittedRedundant;
  mapping(uint256 => struct AAA.RedundantWord[]) redundantEncWords;
}
```

### RedundantWord

_Represents a redundant word submitted by a node._

```solidity
struct RedundantWord {
  bytes word;
  bytes nodePK;
  uint256 index;
}
```

### Word

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

\_Initiates the seed phrase generation protocol.

Requirements:

- The phrase must not be started.\_

#### Parameters

| Name | Type    | Description        |
| ---- | ------- | ------------------ |
| pid  | bytes32 | User's PID.        |
| pk   | bytes   | User's Public Key. |

### submitEncryptedWord

```solidity
function submitEncryptedWord(bytes32 pid, bytes encryptedWord) external
```

\_Submits an encrypted word for a given pid.
When the required number of words is reached, emits {SIDEncryptionRequested} event.

Requirements:

- The encrypted word must not be empty.
- The phrase must have been initiated.
- The sender must be one of the selected nodes.
- The sender must not have already submitted their word.\_

#### Parameters

| Name          | Type    | Description                                    |
| ------------- | ------- | ---------------------------------------------- |
| pid           | bytes32 | User's PID.                                    |
| encryptedWord | bytes   | The word encrypted with the user's public key. |

### submitRedundantWord

```solidity
function submitRedundantWord(bytes32 pid, bytes encryptedWord, bytes nodePK) external
```

\_Submits a redundant encrypted word for a given pid.

Requirements:

- The sender must be a UIP node.\_

#### Parameters

| Name          | Type    | Description                                               |
| ------------- | ------- | --------------------------------------------------------- |
| pid           | bytes32 | User's PID.                                               |
| encryptedWord | bytes   | The redundant word encrypted with the node's public key.  |
| nodePK        | bytes   | The public key of the node submitting the redundant word. |

### submitEncryptedSID

```solidity
function submitEncryptedSID(bytes32 pid, bytes encSID) external
```

\_Stores the encrypted SID for a given pid and marks the phrase as finalized.
Emits {PIDEncryptionRequested} event.

Requirements:

- The sender must be a UIP node.
- The protocol must have been started for the given pid.
- The encrypted SID must not have been already stored.
- The sender must be the node selected to encrypt the SID.\_

#### Parameters

| Name   | Type    | Description                     |
| ------ | ------- | ------------------------------- |
| pid    | bytes32 | User's PID.                     |
| encSID | bytes   | The encrypted SID to be stored. |

### submitEncryptedPID

```solidity
function submitEncryptedPID(bytes32 pid, bytes32 sid, bytes encPID) external
```

\_Stores the encrypted PID and symmetric key for a given pid.
Emits {PhraseComplete} event.

Requirements:

- The sender must be a UIP node.
- The {SIDRecord} must not have been already stored.\_

#### Parameters

| Name   | Type    | Description                                       |
| ------ | ------- | ------------------------------------------------- |
| pid    | bytes32 | User's PID.                                       |
| sid    | bytes32 | User's SID.                                       |
| encPID | bytes   | The encrypted PID and symmetric key to be stored. |

### submitSACRecord

```solidity
function submitSACRecord(uint256 sac, bytes pk) external
```

\_Submits a SAC record linking a public key to a SAC code.

Requirements:

- The public key must not be empty.
- The SAC code must be greater than zero.
- The public key must not have been already stored.
- The SAC code must exist.\_

#### Parameters

| Name | Type    | Description                  |
| ---- | ------- | ---------------------------- |
| sac  | uint256 | The SAC code to be linked.   |
| pk   | bytes   | The public key to be linked. |

### submitSAC

```solidity
function submitSAC(uint256 sac) external
```

\_Submits a SAC code.

Requirements:

- The sender must be a UIP node.\_

#### Parameters

| Name | Type    | Description                   |
| ---- | ------- | ----------------------------- |
| sac  | uint256 | The SAC code to be submitted. |

### getSID

```solidity
function getSID(bytes32 pid) external view returns (bytes)
```

_Returns the encrypted SID for a given pid._

#### Parameters

| Name | Type    | Description |
| ---- | ------- | ----------- |
| pid  | bytes32 | User's PID. |

#### Return Values

| Name | Type  | Description                                |
| ---- | ----- | ------------------------------------------ |
| [0]  | bytes | The encrypted SID associated with the pid. |

### getSIDRecord

```solidity
function getSIDRecord(bytes32 sid) external view returns (bytes, bytes)
```

_Returns the SID record for a given sid._

#### Parameters

| Name | Type    | Description |
| ---- | ------- | ----------- |
| sid  | bytes32 | User's SID. |

#### Return Values

| Name | Type  | Description            |
| ---- | ----- | ---------------------- |
| [0]  | bytes | The encrypted PID.     |
| [1]  | bytes | The user's public key. |

### getSACRecord

```solidity
function getSACRecord(bytes pk) external view returns (uint256)
```

_Returns the SAC code for a given public key._

#### Parameters

| Name | Type  | Description        |
| ---- | ----- | ------------------ |
| pk   | bytes | User's public key. |

#### Return Values

| Name | Type    | Description                                  |
| ---- | ------- | -------------------------------------------- |
| [0]  | uint256 | The SAC code associated with the public key. |

### getWords

```solidity
function getWords(bytes32 pid) external view returns (bytes[] words)
```

_Returns the encrypted words, node public keys, and indexes for a given pid._

#### Parameters

| Name | Type    | Description |
| ---- | ------- | ----------- |
| pid  | bytes32 | User's PID. |

#### Return Values

| Name  | Type    | Description                   |
| ----- | ------- | ----------------------------- |
| words | bytes[] | The array of encrypted words. |
