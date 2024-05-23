// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {IWrappedNativeToken} from "./interfaces/IWrappedNativeToken.sol";
import {ERC20Upgradeable} from "@openzeppelin/contracts-upgradeable@4.9.6/token/ERC20/ERC20Upgradeable.sol";
import {Address} from "@openzeppelin/contracts@4.8.1/utils/Address.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */
contract WrappedNativeToken is IWrappedNativeToken, ERC20Upgradeable {
    using Address for address payable;

    function initialize(string memory symbol) public initializer
    {
        __ERC20_init(string.concat("Wrapped ", symbol), string.concat("W", symbol));
    }

    receive() external payable {
        deposit();
    }

    fallback() external payable {
        deposit();
    }

    function withdraw(uint256 amount) external {
        _burn(msg.sender, amount);
        emit Withdrawal(msg.sender, amount);
        payable(msg.sender).sendValue(amount);
    }

    function deposit() public payable {
        _mint(msg.sender, msg.value);
        emit Deposit(msg.sender, msg.value);
    }
}
