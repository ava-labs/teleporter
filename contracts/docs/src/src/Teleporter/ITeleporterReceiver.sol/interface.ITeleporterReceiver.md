# ITeleporterReceiver
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/Teleporter/ITeleporterReceiver.sol)

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


