# ITeleporterReceiver
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/Teleporter/ITeleporterReceiver.sol)

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


