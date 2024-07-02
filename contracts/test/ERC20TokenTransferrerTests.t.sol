// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TokenTransferrerTest} from "./TokenTransferrerTests.t.sol";
import {IERC20TokenTransferrer} from "../src/interfaces/IERC20TokenTransferrer.sol";
import {SendTokensInput, SendAndCallInput} from "../src/interfaces/ITokenTransferrer.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

abstract contract ERC20TokenTransferrerTest is TokenTransferrerTest {
    using SafeERC20 for IERC20;

    IERC20TokenTransferrer public erc20TokenTransferrer;

    function testZeroSendAmount() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        _setUpRegisteredRemote(
            input.destinationBlockchainID, input.destinationTokenTransferrerAddress, 0
        );
        _setUpExpectedZeroAmountRevert();
        _send(input, 0);
    }

    function testSendNoAllowance() public {
        uint256 amount = 2e15;

        SendTokensInput memory input = _createDefaultSendTokensInput();
        _setUpRegisteredRemote(
            input.destinationBlockchainID, input.destinationTokenTransferrerAddress, 0
        );
        vm.expectRevert("ERC20: insufficient allowance");
        _send(input, amount);
    }

    function testSendAndCallNoAllowance() public {
        uint256 amount = 2e15;

        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        _setUpRegisteredRemote(
            input.destinationBlockchainID, input.destinationTokenTransferrerAddress, 0
        );
        vm.expectRevert("ERC20: insufficient allowance");
        _sendAndCall(input, amount);
    }

    function _send(SendTokensInput memory input, uint256 amount) internal virtual override {
        erc20TokenTransferrer.send(input, amount);
    }

    function _sendAndCall(
        SendAndCallInput memory input,
        uint256 amount
    ) internal virtual override {
        erc20TokenTransferrer.sendAndCall(input, amount);
    }
}
