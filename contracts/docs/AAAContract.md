# Solidity API

## AAAContract

_AAAContract is a smart contract that manages the blockchain interactions for the AAA protocol._

### encryptedWords

```solidity
mapping(bytes32 => bytes[]) encryptedWords
```

_Stores the encrypted words submitted by nodes for each pid_

### phraseHashes

```solidity
mapping(bytes32 => bytes32) phraseHashes
```

_Stores the hash of the complete seed phrase for verification_

### WordRequestedToUIPNode

```solidity
event WordRequestedToUIPNode(bytes32 pid, address node, bytes publicKey)
```

_A word is requested from a UIP node_

### PhraseComplete

```solidity
event PhraseComplete(bytes32 pid)
```

_The seed phrase is completed_

### SeedPhraseProtocoloInitiated

```solidity
event SeedPhraseProtocoloInitiated(bytes32 pid)
```

_The seed phrase generation protocol is initiated_

### constructor

```solidity
constructor(address[] _nodes) public
```

### seedPhraseGenerationProtocol

```solidity
function seedPhraseGenerationProtocol(bytes32 pid, bytes publicKey) external
```

_Initiates the seed phrase generation protocol
by selecting nodes and emitting {WordRequestedToUIPNode} events._

### submitEncryptedWord

```solidity
function submitEncryptedWord(bytes32 pid, bytes encryptedWord) external
```

_Submits an encrypted word for a given pid.
When the required number of words is reached, emits {PhraseComplete} event._

## selectNodes

```solidity
function selectNodes(address[] nodePool, bytes32 pid, uint256 wordsNeeded) internal view returns (address[] selectedNodes)
```

