# BlockHashPublisher
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/CrossChainApplications/VerifiedBlockHash/BlockHashPublisher.sol)

Contract that publishes the latest block hash of current chain to another chain.


## State Variables
### RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT

```solidity
uint256 public constant RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT = 1.5e5;
```


### teleporterRegistry

```solidity
TeleporterRegistry public immutable teleporterRegistry;
```


## Functions
### constructor


```solidity
constructor(address teleporterRegistryAddress);
```

### publishLatestBlockHash

*Publishes the latest block hash to another chain.*


```solidity
function publishLatestBlockHash(bytes32 destinationChainID, address destinationAddress)
    external
    returns (uint256 messageID);
```

## Events
### PublishBlockHash
*Emitted when a block hash is submitted to be published to another chain.*


```solidity
event PublishBlockHash(
    bytes32 indexed destinationChainID,
    address indexed destinationAddress,
    uint256 indexed blockHeight,
    bytes32 blockHash
);
```

