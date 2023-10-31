# RedeemRelayerRewardsTest
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/Teleporter/tests/RedeemRelayerRewardsTests.t.sol)

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

