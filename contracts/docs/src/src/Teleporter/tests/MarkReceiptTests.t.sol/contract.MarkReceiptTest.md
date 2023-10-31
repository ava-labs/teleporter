# MarkReceiptTest
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/Teleporter/tests/MarkReceiptTests.t.sol)

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

