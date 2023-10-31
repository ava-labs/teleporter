# RedeemRelayerRewardsTest
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/Teleporter/tests/RedeemRelayerRewardsTests.t.sol)

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

