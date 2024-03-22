// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITeleporterTokenBridgeTest} from "./ITeleporterTokenBridgeTests.t.sol";
import {IERC20Bridge} from "../src/interfaces/IERC20Bridge.sol";
import {SendTokensInput} from "../src/interfaces/ITeleporterTokenBridge.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

abstract contract IERC20BridgeTest is ITeleporterTokenBridgeTest {
    using SafeERC20 for IERC20;

    IERC20Bridge public erc20Bridge;

    // constructor(IERC20Bridge _app, IERC20 _feeToken) {
    //     erc20Bridge = _app;
    //     feeToken = _feeToken;
    // }

    // function setUp() public virtual override {
    //     super.setUp();
    // }

    function testZeroDestinationBridge() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(0);
        vm.expectRevert(_formatErrorMessage("zero destination bridge address"));
        erc20Bridge.send(input, 0);
    }

    function testZeroRecipient() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.recipient = address(0);
        vm.expectRevert(_formatErrorMessage("zero recipient address"));
        erc20Bridge.send(input, 0);
    }

    function testZeroSendAmount() public {
        vm.expectRevert("SafeERC20TransferFrom: balance not increased");
        erc20Bridge.send(_createDefaultSendTokensInput(), 0);
    }

    function testInsufficientAmountToCoverFees() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = 1;
        feeToken.safeIncreaseAllowance(address(erc20Bridge), input.primaryFee);
        vm.expectRevert(_formatErrorMessage("insufficient amount to cover fees"));
        erc20Bridge.send(input, input.primaryFee);
    }

    function testSendWithFees() public {
        uint256 amount = 2;
        uint256 primaryFee = 1;
        _sendSuccess(amount, primaryFee);
    }

    function testSendNoFees() public {
        uint256 amount = 2;
        uint256 primaryFee = 0;
        _sendSuccess(amount, primaryFee);
    }

    function _send(SendTokensInput memory input, uint256 amount) internal virtual override {
        erc20Bridge.send(input, amount);
    }

    function _checkDeposit(uint256 amount) internal virtual override {
        // Check that transferFrom is called to deposit the funds sent from the user to the bridge
        feeToken.safeIncreaseAllowance(address(tokenBridge), amount);
        vm.expectCall(
            address(feeToken),
            abi.encodeCall(IERC20.transferFrom, (address(this), address(tokenBridge), amount))
        );
    }
}
