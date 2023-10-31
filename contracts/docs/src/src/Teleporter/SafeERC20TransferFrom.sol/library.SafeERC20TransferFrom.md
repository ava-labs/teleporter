# SafeERC20TransferFrom
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/Teleporter/SafeERC20TransferFrom.sol)

*Provides a wrapper used for calling an ERC20 transferFrom method
to receive tokens to a contract from msg.sender.
Checks the balance of the recipient before and after the call to transferFrom, and
returns balance increase. Designed for safely handling ERC20 "fee on transfer" and "burn on transfer" implementations.
Note: A reentrancy guard must always be used when calling token.safeTransferFrom in order to
prevent against possible "before-after" pattern vulnerabilities.*


## Functions
### safeTransferFrom


```solidity
function safeTransferFrom(IERC20 erc20, uint256 amount) internal returns (uint256);
```

