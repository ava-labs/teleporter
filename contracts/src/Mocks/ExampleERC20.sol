// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: BSD-3-Clause

pragma solidity 0.8.18;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

contract ExampleERC20 is ERC20Burnable {
    string private constant _TOKEN_NAME = "Mock Token";
    string private constant _TOKEN_SYMBOL = "EXMP";

    constructor() ERC20(_TOKEN_NAME, _TOKEN_SYMBOL) {
        _mint(msg.sender, 10_000_000_000_000_000_000_000_000_000);
    }

    function mint(uint256 amount) public {
        require(
            amount <= 10_000_000_000_000_000_000,
            "Can only mint 10 at a time."
        );
        _mint(msg.sender, amount);
    }
}
