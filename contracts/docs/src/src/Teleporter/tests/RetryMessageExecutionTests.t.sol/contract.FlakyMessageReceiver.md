# FlakyMessageReceiver
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/Teleporter/tests/RetryMessageExecutionTests.t.sol)

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
function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes calldata messageBytes)
    external;
```

### _receiveMessage


```solidity
function _receiveMessage(bytes32 originChainID, address originSenderAddress, string memory message) internal;
```

### _retryReceive


```solidity
function _retryReceive(bytes32 originChainID, address originSenderAddress, string memory message) internal;
```

