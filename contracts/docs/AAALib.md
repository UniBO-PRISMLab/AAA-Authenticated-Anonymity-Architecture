# Solidity API

## AAALib

_Library that provides utility functions for the AAA protocol._

### selectNodes

```solidity
function selectNodes(bytes32 pid, address[] pool, uint256 wordsNeeded) internal pure returns (address[] selected)
```

_Selects a subset of nodes from a pool based on a given pid and the number of words needed._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's Public Identity Data. |
| pool | address[] | The array of available nodes to select from. |
| wordsNeeded | uint256 | The number of nodes to select. |

### selectNode

```solidity
function selectNode(bytes32 pid, address[] pool) internal pure returns (address selected)
```

_Selects a single node from a pool based on a given pid._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| pid | bytes32 | User's Public Identity Data. |
| pool | address[] | The array of available nodes to select from. |

### deriveSymK

```solidity
function deriveSymK(bytes[] words) internal view returns (bytes32 acc)
```

_Derives a symmetric key from the given words._

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| words | bytes[] | An array of encrypted words. |

