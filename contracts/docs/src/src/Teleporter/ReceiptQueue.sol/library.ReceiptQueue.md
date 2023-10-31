# ReceiptQueue
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/Teleporter/ReceiptQueue.sol)

*ReceiptQueue is a convenience library that creates a queue-like interface of
TeleporterMessageReceipt structs. It provides FIFO properties.
Note: All functions in this library are internal so that the library is not deployed as a contract.*


## State Variables
### _MAXIMUM_RECEIPT_COUNT

```solidity
uint256 private constant _MAXIMUM_RECEIPT_COUNT = 5;
```


## Functions
### enqueue

*Adds a receipt to the queue.*


```solidity
function enqueue(TeleporterMessageReceiptQueue storage queue, TeleporterMessageReceipt memory receipt) internal;
```

### dequeue

*Removes the oldest outstanding receipt from the queue.
Requirements:
- The queue must be non-empty.*


```solidity
function dequeue(TeleporterMessageReceiptQueue storage queue)
    internal
    returns (TeleporterMessageReceipt memory result);
```

### getOutstandingReceiptsToSend

*Returns the outstanding receipts for the given chain ID that should be included in the next message sent.*


```solidity
function getOutstandingReceiptsToSend(TeleporterMessageReceiptQueue storage queue)
    internal
    returns (TeleporterMessageReceipt[] memory result);
```

### size

*Returns the number of outstanding receipts in the queue.*


```solidity
function size(TeleporterMessageReceiptQueue storage queue) internal view returns (uint256);
```

### getReceiptAtIndex

*Returns the receipt at the given index in the queue.*


```solidity
function getReceiptAtIndex(TeleporterMessageReceiptQueue storage queue, uint256 index)
    internal
    view
    returns (TeleporterMessageReceipt memory);
```

## Structs
### TeleporterMessageReceiptQueue

```solidity
struct TeleporterMessageReceiptQueue {
    uint256 first;
    uint256 last;
    mapping(uint256 index => TeleporterMessageReceipt) data;
}
```

