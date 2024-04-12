// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenBridgeTest} from "./TeleporterTokenBridgeTests.t.sol";
import {INativeTokenBridge} from "../src/interfaces/INativeTokenBridge.sol";
import {SendTokensInput, SendAndCallInput} from "../src/interfaces/ITeleporterTokenBridge.sol";

abstract contract NativeTokenBridgeTest is TeleporterTokenBridgeTest {
    INativeTokenBridge public nativeTokenBridge;

    event Deposit(address indexed sender, uint256 amount);
    event Withdrawal(address indexed sender, uint256 amount);

    function _send(SendTokensInput memory input, uint256 amount) internal virtual override {
        nativeTokenBridge.send{value: amount}(input);
    }

    function _sendAndCall(
        SendAndCallInput memory input,
        uint256 amount
    ) internal virtual override {
        nativeTokenBridge.sendAndCall{value: amount}(input);
    }
}
