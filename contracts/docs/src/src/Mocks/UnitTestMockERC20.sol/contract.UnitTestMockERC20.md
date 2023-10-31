# UnitTestMockERC20
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/Mocks/UnitTestMockERC20.sol)


## State Variables
### mockBalances

```solidity
mapping(address => uint256) public mockBalances;
```


### feeOnTransferSenders

```solidity
mapping(address => uint256) public feeOnTransferSenders;
```


## Functions
### setFeeOnTransferSender


```solidity
function setFeeOnTransferSender(address sender, uint256 feeAmount) public;
```

### transferFrom


```solidity
function transferFrom(address from, address to, uint256 amount) public returns (bool);
```

### balanceOf


```solidity
function balanceOf(address account) public view returns (uint256);
```

