# BlockHashReceiver
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/CrossChainApplications/VerifiedBlockHash/BlockHashReceiver.sol)

**Inherits:**
[ITeleporterReceiver](/src/Teleporter/ITeleporterReceiver.sol/interface.ITeleporterReceiver.md), [TeleporterUpgradeable](/src/Teleporter/upgrades/TeleporterUpgradeable.sol/abstract.TeleporterUpgradeable.md), Ownable

Contract for receiving latest block hashes from another chain.


## State Variables
### sourceChainID

```solidity
bytes32 public immutable sourceChainID;
```


### sourcePublisherContractAddress

```solidity
address public immutable sourcePublisherContractAddress;
```


### latestBlockHeight

```solidity
uint256 public latestBlockHeight;
```


### latestBlockHash

```solidity
bytes32 public latestBlockHash;
```


## Functions
### constructor


```solidity
constructor(address teleporterRegistryAddress, bytes32 publisherChainID, address publisherContractAddress)
    TeleporterUpgradeable(teleporterRegistryAddress);
```

### receiveTeleporterMessage

*See {ITeleporterReceiver-receiveTeleporterMessage}.
Receives the latest block hash from another chain
Requirements:
- Sender must be the Teleporter contract.
- Origin sender address must be the source publisher contract address that initiated the BlockHashReceiver.*


```solidity
function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes calldata message)
    external
    onlyAllowedTeleporter;
```

### updateMinTeleporterVersion

*See {TeleporterUpgradeable-updateMinTeleporterVersion}
Updates the minimum Teleporter version allowed for receiving on this contract
to the latest version registered in the {TeleporterRegistry}.
Restricted to only owners of the contract.
Emits a {MinTeleporterVersionUpdated} event.*


```solidity
function updateMinTeleporterVersion() external override onlyOwner;
```

### getLatestBlockInfo

*Returns the latest block information.*


```solidity
function getLatestBlockInfo() public view returns (uint256 height, bytes32 hash);
```

## Events
### ReceiveBlockHash
*Emitted when a new block hash is received from a given origin chain ID.*


```solidity
event ReceiveBlockHash(
    bytes32 indexed originChainID, address indexed originSenderAddress, uint256 indexed blockHeight, bytes32 blockHash
);
```

