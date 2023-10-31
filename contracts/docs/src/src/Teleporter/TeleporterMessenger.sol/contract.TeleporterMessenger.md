# TeleporterMessenger
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/Teleporter/TeleporterMessenger.sol)

**Inherits:**
[ITeleporterMessenger](/src/Teleporter/ITeleporterMessenger.sol/interface.ITeleporterMessenger.md), [ReentrancyGuards](/src/Teleporter/ReentrancyGuards.sol/abstract.ReentrancyGuards.md)

*Implementation of the {ITeleporterMessenger} interface.
This implementation is used to send messages cross-chain using the WarpMessenger precompile,
and to receive messages sent from other chains. Teleporter contracts should be deployed through Nick's method
of universal deployer, such that the same contract is deployed at the same address on all chains.*


## State Variables
### WARP_MESSENGER

```solidity
WarpMessenger public constant WARP_MESSENGER = WarpMessenger(0x0200000000000000000000000000000000000005);
```


### latestMessageIDs

```solidity
mapping(bytes32 => uint256) public latestMessageIDs;
```


### outstandingReceipts

```solidity
mapping(bytes32 => ReceiptQueue.TeleporterMessageReceiptQueue) public outstandingReceipts;
```


### sentMessageInfo

```solidity
mapping(bytes32 => mapping(uint256 => SentMessageInfo)) public sentMessageInfo;
```


### relayerRewardAddresses

```solidity
mapping(bytes32 => mapping(uint256 => address)) public relayerRewardAddresses;
```


### receivedFailedMessageHashes

```solidity
mapping(bytes32 => mapping(uint256 => bytes32)) public receivedFailedMessageHashes;
```


### relayerRewardAmounts

```solidity
mapping(address => mapping(address => uint256)) public relayerRewardAmounts;
```


### blockchainID

```solidity
bytes32 public blockchainID;
```


## Functions
### sendCrossChainMessage

*See {ITeleporterMessenger-sendCrossChainMessage}
When executed, a relayer may kick off an asynchronous event to have the validators of the
chain create an aggregate BLS signature of the message.
Emits a {SendCrossChainMessage} event when message successfully gets sent.*


```solidity
function sendCrossChainMessage(TeleporterMessageInput calldata messageInput)
    external
    senderNonReentrant
    returns (uint256 messageID);
```

### retrySendCrossChainMessage

*See {ITeleporterMessenger-retrySendCrossChainMessage}
Emits a {SendCrossChainMessage} event.
Requirements:
- `message` must have been previously sent to the given `destinationChainID`.
- `message` encoding mush match previously sent message.*


```solidity
function retrySendCrossChainMessage(bytes32 destinationChainID, TeleporterMessage calldata message)
    external
    senderNonReentrant;
```

### addFeeAmount

*See {ITeleporterMessenger-addFeeAmount}
Emits a {AddFeeAmount} event.
Requirements:
- `additionalFeeAmount` must be non-zero.
- `message` must exist and not have been delivered yet.
- `feeContractAddress` must match the fee asset contract address used in the original call to `sendCrossChainMessage`.*


```solidity
function addFeeAmount(
    bytes32 destinationChainID,
    uint256 messageID,
    address feeContractAddress,
    uint256 additionalFeeAmount
) external senderNonReentrant;
```

### receiveCrossChainMessage

*See {ITeleporterMessenger-receiveCrossChainMessage}
Emits a {ReceiveCrossChainMessage} event.
Re-entrancy is explicitly disallowed between receiving functions. One message is not able to receive another message.
Requirements:
- `relayerRewardAddress` must not be the zero address.
- `messageIndex` must specify a valid warp message in the transaction's storage slots.
- Valid warp message provided in storage slots, and sender address matches the address of this contract.
- Warp message `destinationChainID` must match the `blockchainID` of this contract.
- Warp message `destinationAddress` must match the address of this contract.
- Teleporter message was not previously delivered.
- Transaction was sent by an allowed relayer for corresponding teleporter message.*


```solidity
function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) external receiverNonReentrant;
```

### retryMessageExecution

*See {ITeleporterMessenger-retryMessageExecution}
Reverts if the message execution fails again on the specified message.
Emits a {MessageExecuted} event if the retry is successful.
Requirements:
- `message` must have previously failed to execute, and matches the hash of the failed message.*


```solidity
function retryMessageExecution(bytes32 originChainID, TeleporterMessage calldata message)
    external
    receiverNonReentrant;
```

### retryReceipts

*See {ITeleporterMessenger-retryReceipts}
There is no explicit limit to the number of receipts able to be retried in a single message because
this method is intended to be used by relayers themselves to ensure their receipts get returned.
There is no fee associated with the empty message, and the same relayer is expected to relay it
themselves in order to claim their rewards, so it is their responsibility to ensure that the necessary
gas is provided for however many receipts are being retried.
When receipts are retried, they are not removed from their corresponding receipt queue because there
is no efficient way to remove a specific receipt from an arbitrary position in the queue, and it is
harmless for receipts to be sent multiple times within the protocol.
Emits {SendCrossChainMessage} event.
Requirements:
- `messageIDs` must all be valid and have existing receipts.*


```solidity
function retryReceipts(
    bytes32 originChainID,
    uint256[] calldata messageIDs,
    TeleporterFeeInfo calldata feeInfo,
    address[] calldata allowedRelayerAddresses
) external senderNonReentrant returns (uint256 messageID);
```

### redeemRelayerRewards

*See {ITeleporterMessenger-redeemRelayerRewards}
Requirements:
- `rewardAmount` must be non-zero.*


```solidity
function redeemRelayerRewards(address feeAsset) external;
```

### getMessageHash

See {ITeleporterMessenger-getMessageHash}


```solidity
function getMessageHash(bytes32 destinationChainID, uint256 messageID) external view returns (bytes32 messageHash);
```

### messageReceived

*See {ITeleporterMessenger-messageReceived}*


```solidity
function messageReceived(bytes32 originChainID, uint256 messageID) external view returns (bool delivered);
```

### getRelayerRewardAddress

*See {ITeleporterMessenger-getRelayerRewardAddress}*


```solidity
function getRelayerRewardAddress(bytes32 originChainID, uint256 messageID)
    external
    view
    returns (address relayerRewardAddress);
```

### checkRelayerRewardAmount

*See {ITeleporterMessenger-checkRelayerRewardAmount}*


```solidity
function checkRelayerRewardAmount(address relayer, address feeAsset) external view returns (uint256);
```

### getFeeInfo

*See {ITeleporterMessenger-getFeeInfo}*


```solidity
function getFeeInfo(bytes32 destinationChainID, uint256 messageID)
    external
    view
    returns (address feeAsset, uint256 feeAmount);
```

### getNextMessageID

*Returns the next message ID to be used to send a message to the given chain ID.*


```solidity
function getNextMessageID(bytes32 chainID) external view returns (uint256 messageID);
```

### getReceiptQueueSize

*See {ITeleporterMessenger-getReceiptQueueSize}*


```solidity
function getReceiptQueueSize(bytes32 chainID) external view returns (uint256);
```

### getReceiptAtIndex

*See {ITeleporterMessenger-getReceiptAtIndex}*


```solidity
function getReceiptAtIndex(bytes32 chainID, uint256 index) external view returns (TeleporterMessageReceipt memory);
```

### checkIsAllowedRelayer

*Checks whether `delivererAddress` is allowed to deliver the message.*


```solidity
function checkIsAllowedRelayer(address delivererAddress, address[] calldata allowedRelayers)
    external
    pure
    returns (bool);
```

### _sendTeleporterMessage

*Helper function for sending a teleporter message cross chain.
Constructs the Teleporter message and sends it through the Warp Messenger precompile,
and performs fee transfer if necessary.
Emits a {SendCrossChainMessage} event.*


```solidity
function _sendTeleporterMessage(
    bytes32 destinationChainID,
    address destinationAddress,
    TeleporterFeeInfo calldata feeInfo,
    uint256 requiredGasLimit,
    address[] calldata allowedRelayerAddresses,
    bytes memory message,
    TeleporterMessageReceipt[] memory receipts
) private returns (uint256 messageID);
```

### _markReceipt

*Marks the receipt of a message from the given `destinationChainID` with the given `messageID`.
It is possible that the receipt was already received for this message, in which case we return early.
If existing message is found and not yet delivered, we delete it from state and increment the fee/reward*


```solidity
function _markReceipt(bytes32 destinationChainID, uint256 messageID, address relayerRewardAddress) private;
```

### _handleInitialMessageExecution

*Attempts to execute the newly delivered message.
Only revert in the event that the message deliverer (relayer) did not provide enough gas to handle the execution
(including possibly storing a failed message in state). All execution specific errors (i.e. invalid call data, etc)
that are not in the relayers control are caught and handled properly.
Emits a {MessageExecuted} event if the call on destination address is successful.
Emits a {MessageExecutionFailed} event if the call on destination address fails with formatted call data.
Requirements:
- There is enough gas left to cover `message.requiredGasLimit`.*


```solidity
function _handleInitialMessageExecution(bytes32 originChainID, TeleporterMessage memory message) private;
```

### _storeFailedMessageExecution

*Stores the hash of a message that has been successfully delivered but fails to execute properly
such that the message execution can be retried by anyone in the future.*


```solidity
function _storeFailedMessageExecution(bytes32 originChainID, TeleporterMessage memory message) private;
```

### _getNextMessageID

*Returns the next message ID to be used to send a message to the given `chainID`.*


```solidity
function _getNextMessageID(bytes32 chainID) private view returns (uint256 messageID);
```

### _checkIsAllowedRelayer

*Checks whether `delivererAddress` is allowed to deliver the message.*


```solidity
function _checkIsAllowedRelayer(address delivererAddress, address[] memory allowedRelayers)
    private
    pure
    returns (bool);
```

## Structs
### SentMessageInfo

```solidity
struct SentMessageInfo {
    bytes32 messageHash;
    TeleporterFeeInfo feeInfo;
}
```

