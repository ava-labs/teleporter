# MarkReceiptTest
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/Teleporter/tests/MarkReceiptTests.t.sol)

**Inherits:**
[TeleporterMessengerTest](/src/Teleporter/tests/TeleporterMessengerTest.t.sol/contract.TeleporterMessengerTest.md)


## Functions
### setUp


```solidity
function setUp() public virtual override;
```

### testCheckRelayersUponReceipt


```solidity
function testCheckRelayersUponReceipt() public;
```

### testReceiptForNoFeeMessage


```solidity
function testReceiptForNoFeeMessage() public;
```

### testDuplicateReceiptAllowed


```solidity
function testDuplicateReceiptAllowed() public;
```

## Structs
### FeeRewardInfo

```solidity
struct FeeRewardInfo {
    uint256 feeAmount;
    address relayerRewardAddress;
}
```

