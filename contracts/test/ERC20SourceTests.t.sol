// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenSourceTest} from "./TeleporterTokenSourceTests.t.sol";
import {ERC20Source} from "../src/ERC20Source.sol";
import {
    ITeleporterTokenBridge, SendTokensInput
} from "../src/interfaces/ITeleporterTokenBridge.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

contract ERC20SourceTest is TeleporterTokenSourceTest {
    using SafeERC20 for IERC20;

    ERC20Source public app;

    function setUp() public virtual override {
        TeleporterTokenSourceTest.setUp();
        app = new ERC20Source(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            address(mockERC20)
        );
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");
        new ERC20Source(address(0), address(this), address(mockERC20));
    }

    function testZeroTeleporterManagerAddress() public {
        vm.expectRevert("Ownable: new owner is the zero address");
        new ERC20Source(MOCK_TELEPORTER_REGISTRY_ADDRESS, address(0), address(mockERC20));
    }

    function testZeroFeeTokenAddress() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("zero fee token address"));
        new ERC20Source(MOCK_TELEPORTER_REGISTRY_ADDRESS, address(this), address(0));
    }

    /**
     * Send tokens unit tests
     */

    function testSendToSameChain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = DEFAULT_SOURCE_BLOCKCHAIN_ID;
        vm.expectRevert(_formatTokenSourceErrorMessage("cannot bridge to same chain"));
        app.send(input, 0);
    }

    function testZeroDestinationBridge() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(0);
        vm.expectRevert(_formatTokenSourceErrorMessage("zero destination bridge address"));
        app.send(input, 0);
    }

    function testZeroRecipient() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.recipient = address(0);
        vm.expectRevert(_formatTokenSourceErrorMessage("zero recipient address"));
        app.send(input, 0);
    }

    function testZeroSendAmount() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("zero send amount"));
        app.send(_createDefaultSendTokensInput(), 0);
    }

    function testInsufficientAmountToCoverFees() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = 1;
        vm.expectRevert(_formatTokenSourceErrorMessage("insufficient amount to cover fees"));
        mockERC20.safeIncreaseAllowance(address(app), input.primaryFee);
        app.send(input, input.primaryFee);
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

    function testReceiveInvalidMessage() public {
        vm.expectRevert();
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, bytes("test")
        );
    }

    function testReceiveInsufficientBridgeBalance() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("insufficient bridge balance"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            abi.encode(
                SendTokensInput({
                    destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
                    destinationBridgeAddress: address(this),
                    recipient: DEFAULT_RECIPIENT_ADDRESS,
                    primaryFee: 0,
                    secondaryFee: 0,
                    allowedRelayerAddresses: new address[](0)
                }),
                1
            )
        );
    }

    function testReceiveInvalidDestinationBridgeAddress() public {
        // First send to destination blockchain to increase the bridge balance
        _sendSuccess(2, 0);
        vm.expectRevert(_formatTokenSourceErrorMessage("invalid bridge address"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            abi.encode(
                SendTokensInput({
                    destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
                    destinationBridgeAddress: address(0),
                    recipient: DEFAULT_RECIPIENT_ADDRESS,
                    primaryFee: 0,
                    secondaryFee: 0,
                    allowedRelayerAddresses: new address[](0)
                }),
                1
            )
        );
    }

    function testReceiveBridgeBackSuccess() public {
        uint256 amount = 2;
        _sendSuccess(amount, 0);

        uint256 feeAmount = 1;
        uint256 bridgedAmount = amount - feeAmount;
        SendTokensInput memory input = SendTokensInput({
            destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            destinationBridgeAddress: address(app),
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: feeAmount,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            abi.encode(input, bridgedAmount)
        );

        // Make sure the bridge balance is increased
        assertEq(
            app.bridgedBalances(DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS),
            bridgedAmount
        );
    }

    function testMultiHopTransfer() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 2;
        _sendSuccess(amount, 0);

        uint256 feeAmount = 1;
        uint256 bridgedAmount = amount - feeAmount;
        SendTokensInput memory input = SendTokensInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: feeAmount,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });

        _checkExpectedTeleporterCalls(input, bridgedAmount);

        vm.expectEmit(true, true, true, true, address(app));
        emit SendTokens(_MOCK_MESSAGE_ID, address(MOCK_TELEPORTER_MESSENGER_ADDRESS), bridgedAmount);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            abi.encode(input, amount)
        );
    }

    function testMultiHopTransferFails() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 2;
        _sendSuccess(amount, 0);
        uint256 balanceBefore =
            app.bridgedBalances(DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS);
        assertEq(balanceBefore, amount);

        // Fail due to insufficient amount to cover fees
        uint256 feeAmount = amount;
        SendTokensInput memory input = SendTokensInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: feeAmount,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });

        vm.expectRevert(_formatTokenSourceErrorMessage("insufficient amount to cover fees"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            abi.encode(input, amount)
        );

        // Make sure the bridge balance is still the same
        assertEq(
            app.bridgedBalances(DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS),
            balanceBefore
        );
    }

    function _sendSuccess(uint256 amount, uint256 feeAmount) internal {
        uint256 bridgedAmount = amount - feeAmount;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = feeAmount;

        // TODO: Check that deposit was called.

        _checkExpectedTeleporterCalls(input, bridgedAmount);

        vm.expectEmit(true, true, true, true, address(app));
        emit SendTokens(_MOCK_MESSAGE_ID, address(this), bridgedAmount);
        app.send(input, amount);
    }

    function _checkExpectedTeleporterCalls(
        SendTokensInput memory input,
        uint256 bridgeAmount
    ) internal {
        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: input.primaryFee}),
            requiredGasLimit: app.SEND_TOKENS_REQUIRED_GAS(),
            allowedRelayerAddresses: input.allowedRelayerAddresses,
            message: abi.encode(DEFAULT_RECIPIENT_ADDRESS, bridgeAmount)
        });

        if (input.primaryFee > 0) {
            vm.expectCall(
                address(mockERC20),
                abi.encodeCall(
                    IERC20.allowance, (address(app), address(MOCK_TELEPORTER_MESSENGER_ADDRESS))
                )
            );
        }
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput)),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );
    }
}
