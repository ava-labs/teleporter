# ITeleporterReceiver
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/Teleporter/ITeleporterReceiver.sol)

*Interface that cross-chain applications must implement to receive messages from Teleporter.*


## Functions
### receiveTeleporterMessage

*Called by TeleporterMessenger on the receiving chain.*


```solidity
function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes calldata message)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`originChainID`|`bytes32`|is provided by the TeleporterMessenger contract.|
|`originSenderAddress`|`address`|is provided by the TeleporterMessenger contract.|
|`message`|`bytes`|is the TeleporterMessage payload set by the sender.|


