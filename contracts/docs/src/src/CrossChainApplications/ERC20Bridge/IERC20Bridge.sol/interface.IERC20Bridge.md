# IERC20Bridge
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/CrossChainApplications/ERC20Bridge/IERC20Bridge.sol)

*Interface that describes functionalities for a cross-chain ERC20 bridge.*


## Functions
### bridgeTokens

*Transfers ERC20 tokens to another chain.
This can be wrapping, unwrapping, and transferring a wrapped token between two non-native chains.*


```solidity
function bridgeTokens(
    bytes32 destinationChainID,
    address destinationBridgeAddress,
    address tokenContractAddress,
    address recipient,
    uint256 totalAmount,
    uint256 primaryFeeAmount,
    uint256 secondaryFeeAmount
) external;
```

### submitCreateBridgeToken

*Creates a new bridge token on another chain.*


```solidity
function submitCreateBridgeToken(
    bytes32 destinationChainID,
    address destinationBridgeAddress,
    ERC20 nativeToken,
    address messageFeeAsset,
    uint256 messageFeeAmount
) external;
```

## Events
### BridgeTokens
*Emitted when tokens are locked in this bridge contract to be bridged to another chain.*


```solidity
event BridgeTokens(
    address indexed tokenContractAddress,
    bytes32 indexed destinationChainID,
    uint256 indexed teleporterMessageID,
    address destinationBridgeAddress,
    address recipient,
    uint256 amount
);
```

### SubmitCreateBridgeToken
*Emitted when submitting a request to create a new bridge token on another chain.*


```solidity
event SubmitCreateBridgeToken(
    bytes32 indexed destinationChainID,
    address indexed destinationBridgeAddress,
    address indexed nativeContractAddress,
    uint256 teleporterMessageID
);
```

### CreateBridgeToken
*Emitted when creating a new bridge token.*


```solidity
event CreateBridgeToken(
    bytes32 indexed nativeChainID,
    address indexed nativeBridgeAddress,
    address indexed nativeContractAddress,
    address bridgeTokenAddress
);
```

### MintBridgeTokens
*Emitted when minting bridge tokens.*


```solidity
event MintBridgeTokens(address indexed contractAddress, address recipient, uint256 amount);
```

## Enums
### BridgeAction
*Enum representing the action to take on receiving a Teleporter message.*


```solidity
enum BridgeAction {
    Create,
    Mint,
    Transfer
}
```

