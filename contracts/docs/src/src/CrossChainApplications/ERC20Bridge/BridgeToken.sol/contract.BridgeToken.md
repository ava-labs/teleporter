# BridgeToken
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/CrossChainApplications/ERC20Bridge/BridgeToken.sol)

**Inherits:**
ERC20Burnable

*BridgeToken is an ERC20Burnable token contract that is associated with a specific native chain bridge and asset, and is only mintable by the bridge contract on this chain.*


## State Variables
### bridgeContract

```solidity
address public immutable bridgeContract;
```


### nativeChainID

```solidity
bytes32 public immutable nativeChainID;
```


### nativeBridge

```solidity
address public immutable nativeBridge;
```


### nativeAsset

```solidity
address public immutable nativeAsset;
```


### _decimals

```solidity
uint8 private immutable _decimals;
```


## Functions
### constructor

*Initializes a BridgeToken instance.*


```solidity
constructor(
    bytes32 sourceChainID,
    address sourceBridge,
    address sourceAsset,
    string memory tokenName,
    string memory tokenSymbol,
    uint8 tokenDecimals
) ERC20(tokenName, tokenSymbol);
```

### mint

*Mints tokens to `account` if called by original `bridgeContract`.*


```solidity
function mint(address account, uint256 amount) public;
```

### decimals


```solidity
function decimals() public view virtual override returns (uint8);
```

