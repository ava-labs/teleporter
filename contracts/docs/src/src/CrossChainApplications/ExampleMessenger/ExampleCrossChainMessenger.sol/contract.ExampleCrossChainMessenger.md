# ExampleCrossChainMessenger
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/CrossChainApplications/ExampleMessenger/ExampleCrossChainMessenger.sol)

**Inherits:**
[ITeleporterReceiver](/src/Teleporter/ITeleporterReceiver.sol/interface.ITeleporterReceiver.md), ReentrancyGuard, [TeleporterUpgradeable](/src/Teleporter/upgrades/TeleporterUpgradeable.sol/abstract.TeleporterUpgradeable.md), Ownable

*ExampleCrossChainMessenger is an example contract that demonstrates how to send and receive
messages cross chain.*


## State Variables
### _messages

```solidity
mapping(bytes32 => Message) private _messages;
```


## Functions
### constructor


```solidity
constructor(address teleporterRegistryAddress) TeleporterUpgradeable(teleporterRegistryAddress);
```

### receiveTeleporterMessage

*See {ITeleporterReceiver-receiveTeleporterMessage}.
Receives a message from another chain.*


```solidity
function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes calldata message)
    external
    onlyAllowedTeleporter;
```

### sendMessage

*Sends a message to another chain.*


```solidity
function sendMessage(
    bytes32 destinationChainID,
    address destinationAddress,
    address feeContractAddress,
    uint256 feeAmount,
    uint256 requiredGasLimit,
    string calldata message
) external nonReentrant returns (uint256 messageID);
```

### updateMinTeleporterVersion

*See {TeleporterUpgradeable-updateMinTeleporterVersion}
Updates the minimum Teleporter version allowed for receiving on this contract
to the latest version registered in the {TeleporterRegistry}. Also restricts this function to
the owner of this contract.
Emits a {MinTeleporterVersionUpdated} event.*


```solidity
function updateMinTeleporterVersion() external override onlyOwner;
```

### getCurrentMessage

*Returns the current message from another chain.*


```solidity
function getCurrentMessage(bytes32 originChainID) external view returns (address sender, string memory message);
```

## Events
### SendMessage
*Emitted when a message is submited to be sent.*


```solidity
event SendMessage(
    bytes32 indexed destinationChainID,
    address indexed destinationAddress,
    address feeAsset,
    uint256 feeAmount,
    uint256 requiredGasLimit,
    string message
);
```

### ReceiveMessage
*Emitted when a new message is received from a given chain ID.*


```solidity
event ReceiveMessage(bytes32 indexed originChainID, address indexed originSenderAddress, string message);
```

## Structs
### Message

```solidity
struct Message {
    address sender;
    string message;
}
```

