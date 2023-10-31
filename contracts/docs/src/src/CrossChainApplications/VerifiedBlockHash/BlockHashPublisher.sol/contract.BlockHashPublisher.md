# BlockHashPublisher
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/CrossChainApplications/VerifiedBlockHash/BlockHashPublisher.sol)

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

