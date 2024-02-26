// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";

contract WAVAX is IERC20 {
    string public constant NAME = "Wrapped AVAX";
    string public constant SYMBOL = "WAVAX";
    uint8 public constant DECIMALS = 18;

    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;

    event Deposit(address indexed destination, uint256 amount);
    event Withdrawal(address indexed source, uint256 amount);

    receive() external payable {
        deposit();
    }

    fallback() external payable {
        deposit();
    }

    function deposit() public payable {
        balanceOf[msg.sender] += msg.value;
        emit Deposit(msg.sender, msg.value);
    }

    function withdraw(uint256 amount) public {
        require(balanceOf[msg.sender] >= amount, "WAVAX: insufficient balance");
        balanceOf[msg.sender] -= amount;
        payable(msg.sender).transfer(amount);
        emit Withdrawal(msg.sender, amount);
    }

    function approve(address addr, uint256 amount) public returns (bool) {
        allowance[msg.sender][addr] = amount;
        emit Approval(msg.sender, addr, amount);
        return true;
    }

    function transfer(address destination, uint256 amount) public returns (bool) {
        return transferFrom(msg.sender, destination, amount);
    }

    function transferFrom(
        address source,
        address destination,
        uint256 amount
    ) public returns (bool) {
        require(balanceOf[source] >= amount, "WAVAX: insufficient balance");

        if (source != msg.sender) {
            require(allowance[source][msg.sender] >= amount, "WAVAX: insufficient allowance");
            allowance[source][msg.sender] -= amount;
        }

        balanceOf[source] -= amount;
        balanceOf[destination] += amount;

        emit Transfer(source, destination, amount);

        return true;
    }

    function totalSupply() public view returns (uint256) {
        return address(this).balance;
    }
}
