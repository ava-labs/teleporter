// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: BSD-3-Clause

pragma solidity 0.8.18;

// A mock contract for use in unit tests.
contract UnitTestMockERC20 {
    mapping(address => uint256) public mockBalances;
    mapping(address => uint256) public feeOnTransferSenders;

    // If an address on feeOnTransferSenders is the sender address in a transferFrom call,
    // the amount credited to the receiving account is reduced by the feeAmount set for the
    // sender. This is to mock "fee on token transfer" functionality of select ERC20 contracts.
    function setFeeOnTransferSender(address sender, uint256 feeAmount) public {
        feeOnTransferSenders[sender] = feeAmount;
    }

    // The mock allows anyone to call transferFrom to increment the balance of the
    // receipt address. Neither the call or sender need to have sufficient balances to send,
    // we just increment the balance the of the recipient.
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) public returns (bool) {
        uint256 feeAmount = feeOnTransferSenders[from];
        mockBalances[to] += (amount - feeAmount);
        return true;
    }

    function balanceOf(address account) public view returns (uint256) {
        return mockBalances[account];
    }
}
