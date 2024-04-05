// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterTokenBridgeTest} from "./TeleporterTokenBridgeTests.t.sol";
import {TeleporterTokenDestination, IWarpMessenger} from "../src/TeleporterTokenDestination.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";
import {SendTokensInput, SendAndCallInput} from "../src/interfaces/ITeleporterTokenBridge.sol";

abstract contract TeleporterTokenDestinationTest is TeleporterTokenBridgeTest {
    TeleporterTokenDestination public tokenDestination;

    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(DEFAULT_DESTINATION_BLOCKCHAIN_ID)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector)
        );

        _initMockTeleporterRegistry();

        vm.expectCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(
                TeleporterRegistry(MOCK_TELEPORTER_REGISTRY_ADDRESS).latestVersion.selector
            )
        );
    }

    function testInvalidSendingBackToSourceBlockchain() public {
        uint256 amount = 3;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(this);
        _setUpExpectedDeposit(amount);
        vm.expectRevert(_formatErrorMessage("invalid destination bridge address"));
        _send(input, amount);
    }

    function testSendingToSameInstance() public {
        uint256 amount = 3;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = tokenDestination.blockchainID();
        input.destinationBridgeAddress = address(tokenDestination);
        _setUpExpectedDeposit(amount);
        vm.expectRevert(_formatErrorMessage("invalid destination bridge address"));
        _send(input, amount);
    }

    function testSendToSameBlockchainDifferentDestination() public {
        uint256 amount = 2;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = tokenDestination.blockchainID();
        // Set the desintation bridge address to an address different than the token destination contract.
        input.destinationBridgeAddress = address(0x55);

        _sendSingleHopSendSuccess(amount, input.primaryFee);
    }

    function testSendMultiHopSendSuccess() public {
        uint256 amount = 40;
        uint256 primaryFee = 5;
        uint256 secondaryFee = 2;
        _sendMultiHopSendSuccess(amount, primaryFee, secondaryFee);
    }

    function testSendMultiHopCallSuccess() public {
        uint256 amount = 40;
        uint256 primaryFee = 5;
        uint256 secondaryFee = 1;
        _sendMultiHopCallSuccess(amount, primaryFee, secondaryFee);
    }

    function testReceiveInvalidSourceChain() public {
        vm.expectRevert(_formatErrorMessage("invalid source chain"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, new bytes(0)
        );
    }

    function testReceiveInvalidTokenSourceAddress() public {
        vm.expectRevert(_formatErrorMessage("invalid token source address"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, address(0), new bytes(0)
        );
    }

    function testReceiveInvalidMessage() public {
        vm.expectRevert();
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, bytes("test")
        );
    }

    function testReceiveWithdrawSuccess() public {
        uint256 amount = 2;
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        _checkExpectedWithdrawal(DEFAULT_RECIPIENT_ADDRESS, amount);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeSingleHopSendMessage(amount, DEFAULT_RECIPIENT_ADDRESS)
        );
    }

    function testReceiveInvalidMessageType() public {
        vm.expectRevert(_formatErrorMessage("invalid message type"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeMultiHopSendMessage(
                1,
                DEFAULT_DESTINATION_BLOCKCHAIN_ID,
                DEFAULT_DESTINATION_ADDRESS,
                DEFAULT_RECIPIENT_ADDRESS,
                0
            )
        );
    }

    function _sendMultiHopSendSuccess(
        uint256 amount,
        uint256 primaryFee,
        uint256 secondaryFee
    ) internal {
        uint256 bridgeAmount = amount - primaryFee;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = OTHER_BLOCKCHAIN_ID;
        input.primaryFee = primaryFee;
        input.secondaryFee = secondaryFee;

        _setUpExpectedDeposit(amount);
        _checkExpectedTeleporterCallsForSend(
            _createMultiHopSendTeleporterMessageInput(input, bridgeAmount), primaryFee
        );
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit SendTokens(_MOCK_MESSAGE_ID, address(this), bridgeAmount);
        _send(input, amount);
    }

    function _sendMultiHopCallSuccess(
        uint256 amount,
        uint256 primaryFee,
        uint256 secondaryFee
    ) internal {
        uint256 bridgeAmount = amount - primaryFee;
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationBlockchainID = OTHER_BLOCKCHAIN_ID;
        input.primaryFee = primaryFee;
        input.secondaryFee = secondaryFee;

        _setUpExpectedDeposit(amount);
        _checkExpectedTeleporterCallsForSend(
            _createMultiHopCallTeleporterMessageInput(input, bridgeAmount), primaryFee
        );
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit SendTokens(_MOCK_MESSAGE_ID, address(this), bridgeAmount);
        _sendAndCall(input, amount);
    }

    function _requiredGasLimit() internal view virtual override returns (uint256) {
        return tokenDestination.SEND_TOKENS_REQUIRED_GAS();
    }

    function _createDefaultSendTokensInput()
        internal
        pure
        override
        returns (SendTokensInput memory)
    {
        return SendTokensInput({
            destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            destinationBridgeAddress: TOKEN_SOURCE_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: 0,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });
    }

    function _createDefaultSendAndCallInput()
        internal
        pure
        override
        returns (SendAndCallInput memory)
    {
        return SendAndCallInput({
            destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            destinationBridgeAddress: TOKEN_SOURCE_ADDRESS,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: new bytes(16),
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS,
            primaryFee: 0,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });
    }

    function _formatErrorMessage(string memory message)
        internal
        pure
        override
        returns (bytes memory)
    {
        return bytes(string.concat("TeleporterTokenDestination: ", message));
    }

    function _createMultiHopSendTeleporterMessageInput(
        SendTokensInput memory input,
        uint256 bridgeAmount
    ) internal view returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: tokenDestination.sourceBlockchainID(),
            destinationAddress: tokenDestination.tokenSourceAddress(),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(feeToken), amount: input.primaryFee}),
            requiredGasLimit: _requiredGasLimit(),
            allowedRelayerAddresses: input.allowedRelayerAddresses,
            message: _encodeMultiHopSendMessage(
                bridgeAmount,
                input.destinationBlockchainID,
                input.destinationBridgeAddress,
                input.recipient,
                input.secondaryFee
                )
        });
    }

    function _createMultiHopCallTeleporterMessageInput(
        SendAndCallInput memory input,
        uint256 bridgeAmount
    ) internal view returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: tokenDestination.sourceBlockchainID(),
            destinationAddress: tokenDestination.tokenSourceAddress(),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(feeToken), amount: input.primaryFee}),
            requiredGasLimit: _requiredGasLimit(),
            allowedRelayerAddresses: input.allowedRelayerAddresses,
            message: _encodeMultiHopCallMessage({
                amount: bridgeAmount,
                destinationBlockchainID: input.destinationBlockchainID,
                destinationBridgeAddress: input.destinationBridgeAddress,
                recipientContract: input.recipientContract,
                recipientPayload: input.recipientPayload,
                recipientGasLimit: input.recipientGasLimit,
                fallbackRecipient: input.fallbackRecipient,
                secondaryFee: input.secondaryFee
            })
        });
    }
}
