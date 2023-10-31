# RetryReceiptTest
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/Teleporter/tests/RetryReceiptTests.t.sol)

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

