// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TokenTransferrerTest} from "./TokenTransferrerTests.t.sol";
import {INativeTokenTransferrer} from "../src/interfaces/INativeTokenTransferrer.sol";
import {SendTokensInput, SendAndCallInput} from "../src/interfaces/ITokenTransferrer.sol";

abstract contract NativeTokenTransferrerTest is TokenTransferrerTest {
    INativeTokenTransferrer public nativeTokenTransferrer;

    event Deposit(address indexed sender, uint256 amount);
    event Withdrawal(address indexed sender, uint256 amount);

    function testZeroSendAmount() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        _setUpRegisteredRemote(
            input.destinationBlockchainID, input.destinationTokenTransferrerAddress, 0
        );
        _setUpExpectedZeroAmountRevert();
        _send(input, 0);
    }

    function _send(SendTokensInput memory input, uint256 amount) internal virtual override {
        nativeTokenTransferrer.send{value: amount}(input);
    }

    function _sendAndCall(
        SendAndCallInput memory input,
        uint256 amount
    ) internal virtual override {
        nativeTokenTransferrer.sendAndCall{value: amount}(input);
    }
}
