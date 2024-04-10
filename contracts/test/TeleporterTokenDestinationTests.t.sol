// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterTokenBridgeTest} from "./TeleporterTokenBridgeTests.t.sol";
import {TeleporterTokenDestination, IWarpMessenger} from "../src/TeleporterTokenDestination.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";
import {SendTokensInput, SendAndCallInput} from "../src/interfaces/ITeleporterTokenBridge.sol";
import {ITeleporterMessenger} from "@teleporter/ITeleporterMessenger.sol";

abstract contract TeleporterTokenDestinationTest is TeleporterTokenBridgeTest {
    TeleporterTokenDestination public tokenDestination;

    function setUp() public virtual {
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeWithSelector(ITeleporterMessenger.sendCrossChainMessage.selector),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.getBlockchainID, ()),
            abi.encode(DEFAULT_DESTINATION_BLOCKCHAIN_ID)
        );
        vm.expectCall(WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getBlockchainID, ()));

        _initMockTeleporterRegistry();

        vm.expectCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeCall(TeleporterRegistry(MOCK_TELEPORTER_REGISTRY_ADDRESS).latestVersion, ())
        );
    }

    function testInvalidSendingBackToSourceBlockchain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(this);
        _setUpExpectedDeposit(_DEFAULT_TRANSFER_AMOUNT);
        vm.expectRevert(_formatErrorMessage("invalid destination bridge address"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testInvalidSendandCallingBackToSourceBlockchain() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationBridgeAddress = address(this);
        _setUpExpectedDeposit(_DEFAULT_TRANSFER_AMOUNT);
        vm.expectRevert(_formatErrorMessage("invalid destination bridge address"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testNonZeroSecondaryFeeToSourceBlockchain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.secondaryFee = 1;
        _setUpExpectedDeposit(_DEFAULT_TRANSFER_AMOUNT);
        vm.expectRevert(_formatErrorMessage("non-zero secondary fee"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendingToSameInstance() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = tokenDestination.blockchainID();
        input.destinationBridgeAddress = address(tokenDestination);
        _setUpExpectedDeposit(_DEFAULT_TRANSFER_AMOUNT);
        vm.expectRevert(_formatErrorMessage("invalid destination bridge address"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendAndCallingToSameInstance() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationBlockchainID = tokenDestination.blockchainID();
        input.destinationBridgeAddress = address(tokenDestination);
        _setUpExpectedDeposit(_DEFAULT_TRANSFER_AMOUNT);
        vm.expectRevert(_formatErrorMessage("invalid destination bridge address"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendZeroAmountAfterScaling() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.primaryFee = 0;
        input.secondaryFee = 0;
        uint256 amount = 1;
        _setUpExpectedDeposit(amount);
        if (_scaleTokens(amount, false) == 0) {
            vm.expectRevert(_formatErrorMessage("insufficient tokens to transfer"));
        }
        _sendAndCall(input, amount);
    }

    function testSendToSameBlockchainDifferentDestination() public {
        uint256 amount = 200;
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
        uint256 amount = 200;
        uint256 scaledAmount = _scaleTokens(amount, true);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        vm.expectEmit(true, true, true, true, address(tokenDestination));
        emit TokensWithdrawn(DEFAULT_RECIPIENT_ADDRESS, scaledAmount);
        _checkExpectedWithdrawal(DEFAULT_RECIPIENT_ADDRESS, amount);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeSingleHopSendMessage(amount, DEFAULT_RECIPIENT_ADDRESS)
        );
    }

    function testReceiveSendAndCallSuccess() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        _setUpExpectedSendAndCall({
            recipient: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            amount: amount,
            payload: payload,
            gasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            targetHasCode: true,
            expectSuccess: true
        });

        bytes memory message = _encodeSingleHopCallMessage(
            amount,
            DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            payload,
            DEFAULT_RECIPIENT_GAS_LIMIT,
            DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        );
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, message
        );
    }

    function testReceiveSendAndCallFailure() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        _setUpExpectedSendAndCall({
            recipient: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            amount: amount,
            payload: payload,
            gasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            targetHasCode: true,
            expectSuccess: false
        });

        bytes memory message = _encodeSingleHopCallMessage(
            amount,
            DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            payload,
            DEFAULT_RECIPIENT_GAS_LIMIT,
            DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        );
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, message
        );
    }

    function testReceiveSendAndCallFailureNoCode() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        _setUpExpectedSendAndCall({
            recipient: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            amount: amount,
            payload: payload,
            gasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            targetHasCode: false,
            expectSuccess: false
        });

        bytes memory message = _encodeSingleHopCallMessage(
            amount,
            DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            payload,
            DEFAULT_RECIPIENT_GAS_LIMIT,
            DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        );
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, message
        );
    }

    function testReceiveSendAndCallFailureInsufficientGas() public {
        uint256 amount = 200;
        bytes memory payload = hex"DEADBEEF";
        uint256 gasLimit = 5_000_000;
        bytes memory message = _encodeSingleHopCallMessage({
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: gasLimit,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        _setUpMockMint(address(tokenDestination), amount);
        vm.expectRevert("GasUtils: insufficient gas");
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage{gas: gasLimit - 1}(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, message
        );
    }

    function testReceiveInvalidMessageType() public {
        vm.expectRevert(_formatErrorMessage("invalid message type"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeMultiHopSendMessage({
                amount: 1,
                destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
                recipient: DEFAULT_RECIPIENT_ADDRESS,
                secondaryFee: 0,
                secondaryGasLimit: 1_000
            })
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

        // Only tokens destinations scale tokens, so isReceive is always false here.
        uint256 scaledBridgedAmount = _scaleTokens(bridgeAmount, false);

        _checkExpectedTeleporterCallsForSend(
            _createMultiHopSendTeleporterMessageInput(input, scaledBridgedAmount), primaryFee
        );
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, scaledBridgedAmount);
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

        // Only tokens destinations scale tokens, so isReceive is always false here.
        uint256 scaledBridgedAmount = _scaleTokens(bridgeAmount, false);

        _checkExpectedTeleporterCallsForSend(
            _createMultiHopCallTeleporterMessageInput(input, scaledBridgedAmount), primaryFee
        );
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensAndCallSent(_MOCK_MESSAGE_ID, address(this), input, scaledBridgedAmount);
        _sendAndCall(input, amount);
    }

    function _setUpMockMint(address recipient, uint256 amount) internal virtual;

    function _setUpExpectedSendAndCall(
        address recipient,
        uint256 amount,
        bytes memory payload,
        uint256 gasLimit,
        bool targetHasCode,
        bool expectSuccess
    ) internal virtual;

    function _createMultiHopSendTeleporterMessageInput(
        SendTokensInput memory input,
        uint256 bridgeAmount
    ) internal view returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: tokenDestination.sourceBlockchainID(),
            destinationAddress: tokenDestination.tokenSourceAddress(),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(feeToken), amount: input.primaryFee}),
            requiredGasLimit: tokenDestination.MULTIHOP_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: _encodeMultiHopSendMessage({
                amount: bridgeAmount,
                destinationBlockchainID: input.destinationBlockchainID,
                destinationBridgeAddress: input.destinationBridgeAddress,
                recipient: input.recipient,
                secondaryFee: input.secondaryFee,
                secondaryGasLimit: input.requiredGasLimit
            })
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
            requiredGasLimit: tokenDestination.MULTIHOP_REQUIRED_GAS()
                + (input.recipientPayload.length * tokenDestination.MULTIHOP_CALL_GAS_PER_BYTE()),
            allowedRelayerAddresses: new address[](0),
            message: _encodeMultiHopCallMessage({
                amount: bridgeAmount,
                destinationBlockchainID: input.destinationBlockchainID,
                destinationBridgeAddress: input.destinationBridgeAddress,
                recipientContract: input.recipientContract,
                recipientPayload: input.recipientPayload,
                requiredGasLimit: input.requiredGasLimit,
                recipientGasLimit: input.recipientGasLimit,
                fallbackRecipient: input.fallbackRecipient,
                secondaryFee: input.secondaryFee
            })
        });
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
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT
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
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS,
            primaryFee: 0,
            secondaryFee: 0
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
}
