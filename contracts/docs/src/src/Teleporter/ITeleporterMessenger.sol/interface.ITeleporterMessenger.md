# ITeleporterMessenger
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/Teleporter/ITeleporterMessenger.sol)

*Interface that describes functionalities for a cross-chain messenger implementing the Teleporter protcol.*


## Functions
### sendCrossChainMessage

*Called by transactions to initiate the sending of a cross-chain message.*


```solidity
function sendCrossChainMessage(TeleporterMessageInput calldata messageInput) external returns (uint256 messageID);
```

### retrySendCrossChainMessage

*Called by transactions to retry the sending of a cross-chain message.
Retriggers the sending of a message previously emitted by sendCrossChainMessage that has not yet been acknowledged
with a receipt from the destination chain. This may be necessary in the unlikely event that less than the required
threshold of stake weight successfully inserted the message in their messages DB at the time of the first submission.
The message is checked to have already been previously submitted by comparing its message hash against those kept in
state until a receipt is received for the message.*


```solidity
function retrySendCrossChainMessage(bytes32 destinationChainID, TeleporterMessage calldata message) external;
```

### addFeeAmount

*Adds the additional fee amount to the amount to be paid to the relayer that delivers
the given message ID to the destination chain.
The fee contract address must be the same asset type as the fee asset specified in the original
call to sendCrossChainMessage. Returns a failure if the message doesn't exist or there is already
receipt of delivery of the message.*


```solidity
function addFeeAmount(
    bytes32 destinationChainID,
    uint256 messageID,
    address feeContractAddress,
    uint256 additionalFeeAmount
) external;
```

### receiveCrossChainMessage

*Receives a cross-chain message, and marks the `relayerRewardAddress` for fee reward for a successful delivery.
The message specified by `messageIndex` must be provided at that index in the access list storage slots of the transaction,
and is verified in the precompile predicate.*


```solidity
function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) external;
```

### retryMessageExecution

*Retries the execution of a previously delivered message by verifying the payload matches
the hash of the payload originally delivered, and calling the destination address again.
Intended to be used if message excution failed on initial delivery of the Teleporter message.
For example, this may occur if the original required gas limit was not sufficient for the message
execution, or if the destination address did not contain a contract, but a compatible contract
was later deployed to that address. Messages are ensured to be successfully executed at most once.*


```solidity
function retryMessageExecution(bytes32 originChainID, TeleporterMessage calldata message) external;
```

### retryReceipts

*Retries the sending of receipts for the given `messageIDs`.
Sends the specified message receipts in a new message (with an empty payload) back to the origin chain.
This is intended to be used if the message receipts were originally included in messages that were dropped
or otherwise not delivered in a timely manner.*


```solidity
function retryReceipts(
    bytes32 originChainID,
    uint256[] calldata messageIDs,
    TeleporterFeeInfo calldata feeInfo,
    address[] calldata allowedRelayerAddresses
) external returns (uint256 messageID);
```

### redeemRelayerRewards

*Sends any fee amount rewards for the given fee asset out to the caller.*


```solidity
function redeemRelayerRewards(address feeAsset) external;
```

### getMessageHash

*Gets the hash of a given message stored in the EVM state, if the message exists.*


```solidity
function getMessageHash(bytes32 destinationChainID, uint256 messageID) external view returns (bytes32 messageHash);
```

### messageReceived

*Checks whether or not the given message has been received by this chain.*


```solidity
function messageReceived(bytes32 originChainID, uint256 messageID) external view returns (bool delivered);
```

### getRelayerRewardAddress

*Returns the address the relayer reward should be sent to on the origin chain
for a given message, assuming that the message has already been delivered.*


```solidity
function getRelayerRewardAddress(bytes32 originChainID, uint256 messageID)
    external
    view
    returns (address relayerRewardAddress);
```

### checkRelayerRewardAmount

Gets the current reward amount of a given fee asset that is redeemable by the given relayer.


```solidity
function checkRelayerRewardAmount(address relayer, address feeAsset) external view returns (uint256);
```

### getFeeInfo

*Gets the fee asset and amount for a given message.*


```solidity
function getFeeInfo(bytes32 destinationChainID, uint256 messageID)
    external
    view
    returns (address feeAsset, uint256 feeAmount);
```

### getReceiptQueueSize

*Gets the number of receipts that have been sent to the given destination chain ID.*


```solidity
function getReceiptQueueSize(bytes32 chainID) external view returns (uint256 size);
```

### getReceiptAtIndex

*Gets the receipt at the given index in the queue for the given chain ID.*


```solidity
function getReceiptAtIndex(bytes32 chainID, uint256 index)
    external
    view
    returns (TeleporterMessageReceipt memory receipt);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainID`|`bytes32`|The chain ID to get the receipt queue for.|
|`index`|`uint256`|The index of the receipt to get, starting from 0.|


## Events
### SendCrossChainMessage
*Emitted when sending a Teleporter message cross-chain.*


```solidity
event SendCrossChainMessage(
    bytes32 indexed destinationChainID, uint256 indexed messageID, TeleporterMessage message, TeleporterFeeInfo feeInfo
);
```

### AddFeeAmount
*Emitted when an additional fee amount is added to a Teleporter message that had previously
been sent, but not yet delivered to the destination chain.*


```solidity
event AddFeeAmount(bytes32 indexed destinationChainID, uint256 indexed messageID, TeleporterFeeInfo updatedFeeInfo);
```

### MessageExecutionFailed
*Emitted when a Teleporter message is being delivered on the destination chain to an address,
but message execution fails. Failed messages can then be retried with `retryMessageExecution`*


```solidity
event MessageExecutionFailed(bytes32 indexed originChainID, uint256 indexed messageID, TeleporterMessage message);
```

### MessageExecuted
*Emitted when a Teleporter message is successfully executed with the
specified destination address and message call data. This can occur either when
the message is initially received, or on a retry attempt.
Each message received can be executed successfully at most once.*


```solidity
event MessageExecuted(bytes32 indexed originChainID, uint256 indexed messageID);
```

### ReceiveCrossChainMessage
*Emitted when a TeleporterMessage is successfully received.*


```solidity
event ReceiveCrossChainMessage(
    bytes32 indexed originChainID,
    uint256 indexed messageID,
    address indexed deliverer,
    address rewardRedeemer,
    TeleporterMessage message
);
```

### RelayerRewardsRedeemed
*Emitted when an account redeems accumulated relayer rewards.*


```solidity
event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount);
```

