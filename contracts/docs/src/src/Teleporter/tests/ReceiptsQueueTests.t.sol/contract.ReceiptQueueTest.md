# ReceiptQueueTest
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/Teleporter/tests/ReceiptsQueueTests.t.sol)

**Inherits:**
Test


## State Variables
### _queue

```solidity
ReceiptQueue.TeleporterMessageReceiptQueue private _queue;
```


### _receipt1

```solidity
TeleporterMessageReceipt private _receipt1 =
    TeleporterMessageReceipt({receivedMessageID: 543, relayerRewardAddress: 0x10eB43ef5982628728E3E4bb9F78834f67Fbb40b});
```


### _receipt2

```solidity
TeleporterMessageReceipt private _receipt2 = TeleporterMessageReceipt({
    receivedMessageID: 684384,
    relayerRewardAddress: 0x10eB43ef5982628728E3E4bb9F78834f67Fbb40b
});
```


### _receipt3

```solidity
TeleporterMessageReceipt private _receipt3 = TeleporterMessageReceipt({
    receivedMessageID: 654351,
    relayerRewardAddress: 0xcC8E718045817AebA89592C72Ae1C9917f5D0894
});
```


## Functions
### testEnqueueDequeueSuccess


```solidity
function testEnqueueDequeueSuccess() public;
```

### testDequeueRevertIfEmptyQueue


```solidity
function testDequeueRevertIfEmptyQueue() public;
```

### testGetReceiptAtIndex


```solidity
function testGetReceiptAtIndex() public;
```

### testGetReceiptAtIndexWithEmptyQueue


```solidity
function testGetReceiptAtIndexWithEmptyQueue() public;
```

### _formatReceiptQueueErrorMessage


```solidity
function _formatReceiptQueueErrorMessage(string memory errorMessage) private pure returns (bytes memory);
```

