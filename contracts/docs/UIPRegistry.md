# Solidity API

## UIPRegistry

_UIPRegistry manages the registry of authorized UIP nodes._

### owner

```solidity
address owner
```

_Contract owner's address._

### isNode

```solidity
mapping(address => bool) isNode
```

_Mapping to track authorized UIP nodes._

### nodeList

```solidity
address[] nodeList
```

_List of authorized UIP nodes._

### NodeAdded

```solidity
event NodeAdded(address node)
```

_Events for node addition and removal._

### NodeRemoved

```solidity
event NodeRemoved(address node)
```

_Event emitted when a UIP node is removed._

### onlyOwner

```solidity
modifier onlyOwner()
```

_Modifier to restrict function access to the contract owner._

### onlyUIPNode

```solidity
modifier onlyUIPNode()
```

_Modifier to restrict function access to authorized UIP nodes._

### constructor

```solidity
constructor(address[] initialNodes) internal
```

### addNode

```solidity
function addNode(address node) external
```

_Adds a new UIP node to the registry.

Requirements:
- Only the contract owner can add a node.

Emits {NodeAdded} event._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| node | address | The address of the UIP node to be added. |

### removeNode

```solidity
function removeNode(address node) external
```

_Removes a UIP node from the registry.

Requirements:
- Only the contract owner can remove a node.

Emits {NodeRemoved} event._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| node | address | The address of the UIP node to be removed. |

