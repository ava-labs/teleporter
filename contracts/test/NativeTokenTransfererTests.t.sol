// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TokenTransfererTest} from "./TokenTransfererTests.t.sol";
import {INativeTokenTransferer} from "../src/interfaces/INativeTokenTransferer.sol";
import {SendTokensInput, SendAndCallInput} from "../src/interfaces/ITokenTransferer.sol";

abstract contract NativeTokenTransfererTest is TokenTransfererTest {
    INativeTokenTransferer public nativeTokenTransferer;

    event Deposit(address indexed sender, uint256 amount);
    event Withdrawal(address indexed sender, uint256 amount);

    function testZeroSendAmount() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        _setUpRegisteredRemote(
            input.destinationBlockchainID, input.destinationTokenTransfererAddress, 0
        );
        _setUpExpectedZeroAmountRevert();
        _send(input, 0);
    }

    function _send(SendTokensInput memory input, uint256 amount) internal virtual override {
        nativeTokenTransferer.send{value: amount}(input);
    }

    function _sendAndCall(
        SendAndCallInput memory input,
        uint256 amount
    ) internal virtual override {
        nativeTokenTransferer.sendAndCall{value: amount}(input);
    }
}
