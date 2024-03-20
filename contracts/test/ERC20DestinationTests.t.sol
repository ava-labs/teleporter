// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenDestinationTest} from "./TeleporterTokenDestinationTests.t.sol";
import {ERC20Destination} from "../src/ERC20Destination.sol";
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

contract ERC20DestinationTest is TeleporterTokenDestinationTest {
    using SafeERC20 for IERC20;

    event Transfer(address indexed from, address indexed to, uint256 value);

    ERC20Destination public app;

    string public constant MOCK_TOKEN_NAME = "Test Token";
    string public constant MOCK_TOKEN_SYMBOL = "TST";
    uint8 public constant MOCK_TOKEN_DECIMALS = 18;

    function setUp() public override {
        super.setUp();
        app = new ERC20Destination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            MOCK_TOKEN_DECIMALS);

        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(this), 10e18);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, abi.encode(address(this), 10e18)
        );
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");
        new ERC20Destination(
            address(0),
            address(this),
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            MOCK_TOKEN_DECIMALS);
    }

    function testZeroTeleporterManagerAddress() public {
        vm.expectRevert("Ownable: new owner is the zero address");
        new ERC20Destination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(0),
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            MOCK_TOKEN_DECIMALS);
    }

    function testZeroBlockchainID() public {
        vm.expectRevert(_formatTokenDestinationErrorMessage("zero source blockchain ID"));
        new ERC20Destination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            bytes32(0),
            TOKEN_SOURCE_ADDRESS,
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            MOCK_TOKEN_DECIMALS);
    }

    function testDeployToSameBlockchain() public {
        vm.expectRevert(
            _formatTokenDestinationErrorMessage("cannot deploy to same blockchain as source")
        );
        new ERC20Destination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            MOCK_TOKEN_DECIMALS);
    }

    function testZeroTokenSourceAddress() public {
        vm.expectRevert(_formatTokenDestinationErrorMessage("zero token source address"));
        new ERC20Destination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            address(0),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            MOCK_TOKEN_DECIMALS);
    }

    // function testZeroFeeTokenAddress() public {
    //     vm.expectRevert(_formatTokenDestinationErrorMessage("zero fee token address"));
    //     new ERC20Destination(
    //         MOCK_TELEPORTER_REGISTRY_ADDRESS,
    //         address(this),
    //         DEFAULT_SOURCE_BLOCKCHAIN_ID,
    //         address(0),
    //         MOCK_TOKEN_NAME,
    //         MOCK_TOKEN_SYMBOL,
    //         MOCK_TOKEN_DECIMALS);
    // }

    /**
     * Send tokens unit tests
     */
    function testZeroDestinationBridge() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(0);
        vm.expectRevert(_formatTokenDestinationErrorMessage("zero destination bridge address"));
        app.send(input, 0);
    }

    function testZeroRecipient() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.recipient = address(0);
        vm.expectRevert(_formatTokenDestinationErrorMessage("zero recipient address"));
        app.send(input, 0);
    }

    function testInvalidSendingBackToSourceBlockchain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(this);
        vm.expectRevert(_formatTokenDestinationErrorMessage("invalid destination bridge address"));
        app.send(input, 0);
    }

    function testSendingToSameInstance() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = app.blockchainID();
        input.destinationBridgeAddress = address(app);
        vm.expectRevert(_formatTokenDestinationErrorMessage("invalid destination bridge address"));
        app.send(input, 0);
    }

    function testZeroSendAmount() public {
        vm.expectRevert("SafeERC20TransferFrom: balance not increased");
        app.send(_createDefaultSendTokensInput(), 0);
    }

    function testInsufficientAmountToCoverFees() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = 1;
        uint256 amount = input.primaryFee;

        IERC20(app).safeIncreaseAllowance(address(app), amount);
        vm.expectRevert(_formatTokenDestinationErrorMessage("insufficient amount to cover fees"));
        app.send(input, amount);
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

    function testSendToSameBlockchainDifferentDestination() public {
        // Send a transfer to the same app itself
        uint256 amount = 2;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = app.blockchainID();
        input.destinationBridgeAddress = address(this);

        _sendSuccess(amount, input.primaryFee);
    }

    /**
     * Receive tokens unit tests
     */

    function testInvalidSourceChain() public {
        vm.expectRevert(_formatTokenDestinationErrorMessage("invalid source chain"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, new bytes(0)
        );
    }

    function testInvalidTokenSourceAddress() public {
        vm.expectRevert(_formatTokenDestinationErrorMessage("invalid token source address"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, address(0), new bytes(0));
    }

    function testReceivedInvalidMessage() public {
        vm.expectRevert();
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, bytes("test")
        );
    }

    function _sendSuccess(uint256 amount, uint256 feeAmount) internal {
        IERC20(app).safeIncreaseAllowance(address(app), amount);
        uint256 bridgedAmount = amount - feeAmount;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = feeAmount;

        // Check that transferFrom is called to deposit the funds sent from the user to the bridge
        vm.expectCall(
            address(app), abi.encodeCall(IERC20.transferFrom, (address(this), address(app), amount))
        );

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
            destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            destinationAddress: TOKEN_SOURCE_ADDRESS,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(app), amount: input.primaryFee}),
            requiredGasLimit: app.SEND_TOKENS_REQUIRED_GAS(),
            allowedRelayerAddresses: input.allowedRelayerAddresses,
            message: abi.encode(
                SendTokensInput({
                    destinationBlockchainID: input.destinationBlockchainID,
                    destinationBridgeAddress: input.destinationBridgeAddress,
                    recipient: input.recipient,
                    primaryFee: input.secondaryFee,
                    secondaryFee: 0,
                    allowedRelayerAddresses: input.allowedRelayerAddresses
                }),
                bridgeAmount
                )
        });

        if (input.primaryFee > 0) {
            vm.expectCall(
                address(app),
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
