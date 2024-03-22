// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenSourceTest} from "./TeleporterTokenSourceTests.t.sol";
import {ITeleporterTokenBridgeTest} from "./ITeleporterTokenBridgeTests.t.sol";
import {NativeTokenSource} from "../src/NativeTokenSource.sol";
import {
    ITeleporterTokenBridge, SendTokensInput
} from "../src/interfaces/ITeleporterTokenBridge.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {IWrappedNativeToken} from "../src/interfaces/IWrappedNativeToken.sol";
import {ExampleWAVAX} from "../src/mocks/ExampleWAVAX.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/ERC20.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";

contract NativeTokenSourceTest is TeleporterTokenSourceTest {
    NativeTokenSource public app;
    IWrappedNativeToken public mockWrappedToken;

    function setUp() public override {
        TeleporterTokenSourceTest.setUp();
        // IERC20BridgeTest.setUp();

        mockWrappedToken = new ExampleWAVAX();
        app = new NativeTokenSource(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            address(mockWrappedToken)
        );
        tokenSource = app;
        tokenBridge = app;
        feeToken = mockWrappedToken;
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");
        new NativeTokenSource(address(0), address(this), address(mockWrappedToken));
    }

    function testZeroTeleporterManagerAddress() public {
        vm.expectRevert("Ownable: new owner is the zero address");
        new NativeTokenSource(MOCK_TELEPORTER_REGISTRY_ADDRESS, address(0), address(mockWrappedToken));
    }

    function testZeroFeeTokenAddress() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("zero fee token address"));
        new NativeTokenSource(MOCK_TELEPORTER_REGISTRY_ADDRESS, address(this), address(0));
    }

    /**
     * Send tokens unit tests
     */

    function testZeroDestinationBridge() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(0);
        vm.expectRevert(_formatTokenSourceErrorMessage("zero destination bridge address"));
        app.send{value: 0}(input);
    }

    function testZeroRecipient() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.recipient = address(0);
        vm.expectRevert(_formatTokenSourceErrorMessage("zero recipient address"));
        app.send{value: 0}(input);
    }

    function testZeroSendAmount() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("insufficient amount to cover fees"));
        app.send{value: 0}(_createDefaultSendTokensInput());
    }

    function testInsufficientAmountToCoverFees() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = 1;
        // mockWrappedToken.safeIncreaseAllowance(address(app), input.primaryFee);
        vm.expectRevert(_formatTokenSourceErrorMessage("insufficient amount to cover fees"));
        app.send{value: input.primaryFee}(input);
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

    /**
     * Receive tokens unit tests
     */

    function _send(SendTokensInput memory input, uint256 amount) internal virtual override {
        app.send{value: amount}(input);
    }

    function _checkDeposit(uint256 amount) internal virtual override {}
}
