// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {
    TeleporterUpgradeableTest, ExampleUpgradeableApp
} from "./TeleporterUpgradeableTests.t.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "../../ITeleporterMessenger.sol";
import {TeleporterRegistry, ProtocolRegistryEntry} from "../TeleporterRegistry.sol";
import {IERC20} from "../../TeleporterMessenger.sol";
import {UnitTestMockERC20} from "@mocks/UnitTestMockERC20.sol";

contract SendTeleporterMessageTest is TeleporterUpgradeableTest {
    UnitTestMockERC20 internal _mockFeeAsset;

    function setUp() public virtual override {
        TeleporterUpgradeableTest.setUp();

        _mockFeeAsset = new UnitTestMockERC20();
    }

    function testSendTeleporterPaused() public {
        // Pause the Teleporter address returned from _getTeleporterMessenger()
        address teleporter = address(app.getTeleporterMessenger());
        _pauseTeleporterAddressSuccess(app, teleporter);

        // Check that the app reverts when trying to send a message with paused Teleporter
        vm.expectRevert(_formatTeleporterUpgradeableErrorMessage("Teleporter sending paused"));
        app.sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
                destinationAddress: DEFAULT_DESTINATION_ADDRESS,
                feeInfo: TeleporterFeeInfo(address(0), 0),
                requiredGasLimit: 0,
                allowedRelayerAddresses: new address[](0),
                message: hex"deadbeef"
            })
        );
    }

    function testNoRegisteredTeleporterSend() public {
        // Create a new Teleporter registry with no registered Teleporters
        TeleporterRegistry teleporterRegistry = new TeleporterRegistry(
            new ProtocolRegistryEntry[](0)
        );

        // Create a new app with the new Teleporter registry
        ExampleUpgradeableApp app = new ExampleUpgradeableApp(address(teleporterRegistry));

        // Check that the app reverts when trying to send a message with no registered Teleporter
        vm.expectRevert(_formatRegistryErrorMessage("zero version"));
        app.sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
                destinationAddress: DEFAULT_DESTINATION_ADDRESS,
                feeInfo: TeleporterFeeInfo(address(0), 0),
                requiredGasLimit: 0,
                allowedRelayerAddresses: new address[](0),
                message: hex"deadbeef"
            })
        );
    }

    function testSendNoFee() public {
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(address(0), 0),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            message: hex"deadbeef"
        });
        vm.mockCall(
            address(app.getTeleporterMessenger()),
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (messageInput)),
            abi.encode(bytes32(0))
        );
        vm.expectCall(
            address(app.getTeleporterMessenger()),
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (messageInput))
        );

        app.sendTeleporterMessage(messageInput);
    }

    function testSendWithFee() public {
        uint256 feeAmount = 1;
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(address(_mockFeeAsset), feeAmount),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        vm.mockCall(
            address(app.getTeleporterMessenger()),
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (messageInput)),
            abi.encode(bytes32(0))
        );
        vm.expectCall(
            address(_mockFeeAsset),
            abi.encodeCall(IERC20.transferFrom, (address(this), address(app), feeAmount))
        );
        vm.expectCall(
            address(app.getTeleporterMessenger()),
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (messageInput))
        );

        app.sendTeleporterMessage(messageInput);
    }

    function testInvalidFeeAsset() public {
        address invalidFeeAsset = 0xb8be9140D8717f4a8fd7e8ae23C5668bc3A4B39c;
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(invalidFeeAsset, 1),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        // We expect a call to the fee asset address, but since we did not
        // mock it there is no code there to execute.
        vm.expectCall(invalidFeeAsset, abi.encodeCall(IERC20.balanceOf, (address(app))));
        vm.expectRevert();
        app.sendTeleporterMessage(messageInput);
    }

    function testFeeTransferFailure() public {
        uint256 feeAmount = 1;
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(address(_mockFeeAsset), feeAmount),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        vm.mockCall(
            address(_mockFeeAsset),
            abi.encodeCall(IERC20.transferFrom, (address(this), address(app), feeAmount)),
            abi.encode(false)
        );
        vm.expectCall(
            address(_mockFeeAsset),
            abi.encodeCall(IERC20.transferFrom, (address(this), address(app), feeAmount))
        );
        vm.expectRevert("SafeERC20: ERC20 operation did not succeed");
        app.sendTeleporterMessage(messageInput);
    }

    function testZeroFeeAsset() public {
        TeleporterMessageInput memory messageInput = TeleporterMessageInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationAddress: DEFAULT_DESTINATION_ADDRESS,
            feeInfo: TeleporterFeeInfo(address(0), 1),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            message: new bytes(0)
        });

        vm.expectRevert(_formatTeleporterUpgradeableErrorMessage("zero fee token address"));
        app.sendTeleporterMessage(messageInput);
    }
}
