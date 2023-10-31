# SampleMessageReceiver
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/Teleporter/tests/HandleInitialMessageExecutionTests.t.sol)

**Inherits:**
[ITeleporterReceiver](/src/Teleporter/ITeleporterReceiver.sol/interface.ITeleporterReceiver.md)


## State Variables
### teleporterContract

```solidity
address public immutable teleporterContract;
```


### latestMessage

```solidity
string public latestMessage;
```


### latestMessageSenderSubnetID

```solidity
bytes32 public latestMessageSenderSubnetID;
```


### latestMessageSenderAddress

```solidity
address public latestMessageSenderAddress;
```


## Functions
### constructor


```solidity
constructor(address teleporterContractAddress);
```

### receiveTeleporterMessage


```solidity
function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes calldata message)
    external;
```

### _receiveMessage


```solidity
function _receiveMessage(bytes32 originChainID, address originSenderAddress, string memory message, bool succeed)
    internal;
```

### _receiveMessageRecursive


```solidity
function _receiveMessageRecursive(bytes32 originChainID, address originSenderAddress, string memory message) internal;
```

