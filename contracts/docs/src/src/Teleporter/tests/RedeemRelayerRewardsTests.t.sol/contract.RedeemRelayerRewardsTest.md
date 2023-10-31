# RedeemRelayerRewardsTest
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/Teleporter/tests/RedeemRelayerRewardsTests.t.sol)

**Inherits:**
[TeleporterMessengerTest](/src/Teleporter/tests/TeleporterMessengerTest.t.sol/contract.TeleporterMessengerTest.md)


## Functions
### setUp


```solidity
function setUp() public virtual override;
```

### testZeroRewardBalance


```solidity
function testZeroRewardBalance() public;
```

### testRedemptionFails


```solidity
function testRedemptionFails() public;
```

### testRedemptionSucceeds


```solidity
function testRedemptionSucceeds() public;
```

### _setUpRelayerRewards


```solidity
function _setUpRelayerRewards(FeeRewardInfo memory feeRewardInfo) private;
```

## Structs
### FeeRewardInfo

```solidity
struct FeeRewardInfo {
    uint256 feeAmount;
    address relayerRewardAddress;
}
```

