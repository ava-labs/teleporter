# TeleporterMessage
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/Teleporter/ITeleporterMessenger.sol)


```solidity
struct TeleporterMessage {
    uint256 messageID;
    address senderAddress;
    address destinationAddress;
    uint256 requiredGasLimit;
    address[] allowedRelayerAddresses;
    TeleporterMessageReceipt[] receipts;
    bytes message;
}
```

