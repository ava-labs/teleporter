// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenBridgeTest} from "./TeleporterTokenBridgeTests.t.sol";
import {TeleporterTokenSource, IWarpMessenger} from "../src/TeleporterTokenSource.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    MultiHopSendMessage
} from "../src/interfaces/ITeleporterTokenBridge.sol";

abstract contract TeleporterTokenSourceTest is TeleporterTokenBridgeTest {
    TeleporterTokenSource public tokenSource;

    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(DEFAULT_SOURCE_BLOCKCHAIN_ID)
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

    function testSendToSameChain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = DEFAULT_SOURCE_BLOCKCHAIN_ID;
        vm.expectRevert(_formatErrorMessage("cannot bridge to same chain"));
        _send(input, 0);
    }

    function testReceiveInvalidMessage() public {
        vm.expectRevert();
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, bytes("test")
        );
    }

    function testReceiveInsufficientBridgeBalance() public {
        vm.expectRevert(_formatErrorMessage("insufficient bridge balance"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeSingleHopSendMessage(1, DEFAULT_RECIPIENT_ADDRESS)
        );
    }

    function testReceiveWithdrawSuccess() public {
        uint256 amount = 2;
        _sendSuccess(amount, 0);

        uint256 feeAmount = 1;
        uint256 bridgedAmount = amount - feeAmount;
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        _checkExpectedWithdrawal(DEFAULT_RECIPIENT_ADDRESS, bridgedAmount);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeSingleHopSendMessage(bridgedAmount, DEFAULT_RECIPIENT_ADDRESS)
        );

        // Make sure the bridge balance is increased
        assertEq(
            tokenSource.bridgedBalances(
                DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS
            ),
            bridgedAmount
        );
    }

    function testMultiHopTransfer() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 2;
        _sendSuccess(amount, 0);

        uint256 feeAmount = 1;
        uint256 bridgeAmount = amount - feeAmount;
        SendTokensInput memory input = SendTokensInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: feeAmount,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });
        _checkExpectedTeleporterCallsForSend(
            _createExpectedTeleporterMessageInput(input, bridgeAmount), feeAmount
        );

        vm.expectEmit(true, true, true, true, address(tokenSource));
        emit SendTokens(_MOCK_MESSAGE_ID, address(MOCK_TELEPORTER_MESSENGER_ADDRESS), bridgeAmount);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeMultiHopSendMessage(
                amount,
                input.destinationBlockchainID,
                input.destinationBridgeAddress,
                input.recipient,
                input.primaryFee
            )
        );
    }

    function testMultiHopTransferFails() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 2;
        _sendSuccess(amount, 0);
        uint256 balanceBefore = tokenSource.bridgedBalances(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS
        );
        assertEq(balanceBefore, amount);

        // Fail due to insufficient amount to cover fees
        uint256 feeAmount = amount;
        MultiHopSendMessage memory message = MultiHopSendMessage({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            secondaryFee: feeAmount
        });

        vm.expectRevert(_formatErrorMessage("insufficient amount to cover fees"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeMultiHopSendMessage(
                amount,
                message.destinationBlockchainID,
                message.destinationBridgeAddress,
                message.recipient,
                message.secondaryFee
            )
        );

        // Make sure the bridge balance is still the same
        assertEq(
            tokenSource.bridgedBalances(
                DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS
            ),
            balanceBefore
        );
    }

    function _requiredGasLimit() internal view virtual override returns (uint256) {
        return tokenSource.SEND_TOKENS_REQUIRED_GAS();
    }

    function _createDefaultSendTokensInput()
        internal
        pure
        override
        returns (SendTokensInput memory)
    {
        return SendTokensInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
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
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
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
        return bytes(string.concat("TeleporterTokenSource: ", message));
    }
}
