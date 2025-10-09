// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @dev UIPRegistry manages the registry of authorized UIP nodes.
 */
abstract contract UIPRegistry {
    /// @dev Contract owner's address.
    address public owner;

    /// @dev Mapping to track authorized UIP nodes.
    mapping(address => bool) public isNode;

    /// @dev List of authorized UIP nodes.
    address[] public nodeList;

    /// @dev Events for node addition and removal.
    event NodeAdded(address node);

    /// @dev Event emitted when a UIP node is removed.
    event NodeRemoved(address node);

    /// @dev Modifier to restrict function access to the contract owner.
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner");
        _;
    }

    /// @dev Modifier to restrict function access to authorized UIP nodes.
    modifier onlyUIPNode() {
        require(isNode[msg.sender], "Not UIP node");
        _;
    }

    constructor(address[] memory initialNodes) {
        owner = msg.sender;
        for (uint i = 0; i < initialNodes.length; i++) {
            address n = initialNodes[i];
            if (!isNode[n]) {
                isNode[n] = true;
                nodeList.push(n);
                emit NodeAdded(n);
            }
        }
    }

    /**
     * @dev Adds a new UIP node to the registry.
     *
     * Requirements:
     * - Only the contract owner can add a node.
     *
     * @param node The address of the UIP node to be added.
     */
    function addNode(address node) external onlyOwner {
        require(!isNode[node], "exists");
        isNode[node] = true;
        nodeList.push(node);
        emit NodeAdded(node);
    }

    /**
     * @dev Removes a UIP node from the registry.
     *
     * Requirements:
     * - Only the contract owner can remove a node.
     *
     * @param node The address of the UIP node to be removed.
     */
    function removeNode(address node) external onlyOwner {
        require(isNode[node], "absent");
        isNode[node] = false;
        for (uint i = 0; i < nodeList.length; i++) {
            if (nodeList[i] == node) {
                nodeList[i] = nodeList[nodeList.length - 1];
                nodeList.pop();
                break;
            }
        }
        emit NodeRemoved(node);
    }
}
