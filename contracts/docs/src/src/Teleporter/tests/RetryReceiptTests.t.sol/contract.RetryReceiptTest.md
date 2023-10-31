# RetryReceiptTest
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/Teleporter/tests/RetryReceiptTests.t.sol)

**Inherits:**
[TeleporterMessengerTest](/src/Teleporter/tests/TeleporterMessengerTest.t.sol/contract.TeleporterMessengerTest.md)


## Functions
### setUp


```solidity
function setUp() public virtual override;
```

### testSuccess


```solidity
function testSuccess() public;
```

### testDuplicateAllowed


```solidity
function testDuplicateAllowed() public;
```

### testMissingMessage


```solidity
function testMissingMessage() public;
```

### _retryTestReceiptsWithFee


```solidity
function _retryTestReceiptsWithFee(bytes32 chainID, uint256[] memory messageIDs, address feeAddress, uint256 feeAmount)
    private
    returns (uint256);
```

### _retryTestReceiptsWithNoFee


```solidity
function _retryTestReceiptsWithNoFee(bytes32 chainID, uint256[] memory messageIDs) private returns (uint256);
```

