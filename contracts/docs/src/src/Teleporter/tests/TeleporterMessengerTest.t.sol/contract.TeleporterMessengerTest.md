# TeleporterMessengerTest
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/Teleporter/tests/TeleporterMessengerTest.t.sol)

**Inherits:**
Test


## State Variables
### teleporterMessenger

```solidity
TeleporterMessenger public teleporterMessenger;
```


### MOCK_BLOCK_CHAIN_ID

```solidity
bytes32 public constant MOCK_BLOCK_CHAIN_ID = bytes32(uint256(123456));
```


### DEFAULT_ORIGIN_CHAIN_ID

```solidity
bytes32 public constant DEFAULT_ORIGIN_CHAIN_ID =
    bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
```


### DEFAULT_DESTINATION_CHAIN_ID

```solidity
bytes32 public constant DEFAULT_DESTINATION_CHAIN_ID =
    bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
```


### DEFAULT_DESTINATION_ADDRESS

```solidity
address public constant DEFAULT_DESTINATION_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
```


### DEFAULT_REQUIRED_GAS_LIMIT

```solidity
uint256 public constant DEFAULT_REQUIRED_GAS_LIMIT = 1e6;
```


### WARP_PRECOMPILE_ADDRESS

```solidity
address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
```


### DEFAULT_RELAYER_REWARD_ADDRESS

```solidity
address public constant DEFAULT_RELAYER_REWARD_ADDRESS = 0xa4CEE7d1aF6aDdDD33E3b1cC680AB84fdf1b6d1d;
```


### _mockFeeAsset

```solidity
UnitTestMockERC20 internal _mockFeeAsset;
```


## Functions
### setUp


```solidity
function setUp() public virtual;
```

### testEmptyReceiptQueue


```solidity
function testEmptyReceiptQueue() public;
```

### _sendTestMessageWithFee


```solidity
function _sendTestMessageWithFee(bytes32 chainID, uint256 feeAmount) internal returns (uint256);
```

### _sendTestMessageWithNoFee


```solidity
function _sendTestMessageWithNoFee(bytes32 chainID) internal returns (uint256);
```

### _setUpSuccessGetVerifiedWarpMessageMock


```solidity
function _setUpSuccessGetVerifiedWarpMessageMock(uint32 index, WarpMessage memory warpMessage) internal;
```

### _receiveTestMessage


```solidity
function _receiveTestMessage(
    bytes32 originChainID,
    uint256 messageID,
    address relayerRewardAddress,
    TeleporterMessageReceipt[] memory receipts
) internal;
```

### _receiveMessageSentToEOA


```solidity
function _receiveMessageSentToEOA() internal returns (bytes32, address, TeleporterMessage memory);
```

### _createMockTeleporterMessage


```solidity
function _createMockTeleporterMessage(uint256 messageID, bytes memory message)
    internal
    view
    returns (TeleporterMessage memory);
```

### _createDefaultWarpMessage


```solidity
function _createDefaultWarpMessage(bytes32 originChainID, bytes memory payload)
    internal
    view
    returns (WarpMessage memory);
```

### _formatTeleporterErrorMessage


```solidity
function _formatTeleporterErrorMessage(string memory errorMessage) internal pure returns (bytes memory);
```

## Events
### SendCrossChainMessage

```solidity
event SendCrossChainMessage(
    bytes32 indexed destinationChainID, uint256 indexed messageID, TeleporterMessage message, TeleporterFeeInfo feeInfo
);
```

### AddFeeAmount

```solidity
event AddFeeAmount(bytes32 indexed destinationChainID, uint256 indexed messageID, TeleporterFeeInfo updatedFeeInfo);
```

### ReceiveCrossChainMessage

```solidity
event ReceiveCrossChainMessage(
    bytes32 indexed originChainID,
    uint256 indexed messageID,
    address indexed deliverer,
    address rewardRedeemer,
    TeleporterMessage message
);
```

### MessageExecutionFailed

```solidity
event MessageExecutionFailed(bytes32 indexed originChainID, uint256 indexed messageID, TeleporterMessage message);
```

### MessageExecuted

```solidity
event MessageExecuted(bytes32 indexed originChainID, uint256 indexed messageID);
```

### FailedFeePayment

```solidity
event FailedFeePayment(
    bytes32 indexed destinationChainID,
    uint256 indexed messageID,
    address indexed feeAsset,
    uint256 feeAmount,
    address relayerRewardAddress
);
```

### RelayerRewardsRedeemed

```solidity
event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount);
```

