# ExampleERC20
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/Mocks/ExampleERC20.sol)

**Inherits:**
ERC20Burnable


## State Variables
### _TOKEN_NAME

```solidity
string private constant _TOKEN_NAME = "Mock Token";
```


### _TOKEN_SYMBOL

```solidity
string private constant _TOKEN_SYMBOL = "EXMP";
```


### _MAX_MINT

```solidity
uint256 private constant _MAX_MINT = 10_000_000_000_000_000;
```


## Functions
### constructor


```solidity
constructor() ERC20(_TOKEN_NAME, _TOKEN_SYMBOL);
```

### mint


```solidity
function mint(uint256 amount) public;
```

