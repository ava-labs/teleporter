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
    BridgeMessageType,
    BridgeMessage,
    MultiHopSendMessage,
    RegisterDestinationMessage
} from "../src/interfaces/ITeleporterTokenBridge.sol";
import {ITeleporterMessenger} from "@teleporter/ITeleporterMessenger.sol";

abstract contract TeleporterTokenSourceTest is TeleporterTokenBridgeTest {
    TeleporterTokenSource public tokenSource;

    function setUp() public virtual {
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeWithSelector(ITeleporterMessenger.sendCrossChainMessage.selector),
            abi.encode(_MOCK_MESSAGE_ID)
        );
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

    function testSendToUnregisteredDestination() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = DEFAULT_SOURCE_BLOCKCHAIN_ID;
        vm.expectRevert(_formatErrorMessage("destination bridge not registered"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testNonZeroSecondaryFee() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.secondaryFee = 1;
        vm.expectRevert(_formatErrorMessage("non-zero secondary fee"));
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
        uint256 amount = 200;
        _sendSingleHopSendSuccess(amount, 0);

        uint256 feeAmount = 10;
        uint256 bridgeAmount = amount - feeAmount;
        SendTokensInput memory input = _createDefaultReceiveTokensInput();
        input.primaryFee = feeAmount;

        _checkExpectedWithdrawal(DEFAULT_RECIPIENT_ADDRESS, bridgeAmount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeSingleHopSendMessage(bridgeAmount, DEFAULT_RECIPIENT_ADDRESS)
        );

        // Make sure the bridge balance is correct. Only the fee amount remains locked in the source
        // contract. The rest is withdrawn.
        assertEq(
            tokenSource.bridgedBalances(
                DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS
            ),
            feeAmount
        );
    }

    function testReceiveSendAndCallSuccess() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        bytes32 sourceBlockchainID = DEFAULT_DESTINATION_BLOCKCHAIN_ID;
        address originSenderAddress = address(this);
        bytes memory payload = hex"DEADBEEF";
        _setUpExpectedSendAndCall({
            sourceBlockchainID: sourceBlockchainID,
            originSenderAddress: originSenderAddress,
            recipient: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            amount: amount,
            payload: payload,
            gasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            targetHasCode: true,
            expectSuccess: true
        });

        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: sourceBlockchainID,
            originSenderAddress: originSenderAddress,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, message
        );
    }

    function testReceiveSendAndCallFailure() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 2;
        _sendSingleHopSendSuccess(amount, 0);

        bytes32 sourceBlockchainID = DEFAULT_DESTINATION_BLOCKCHAIN_ID;
        address originSenderAddress = address(this);
        bytes memory payload = hex"DEADBEEF";
        _setUpExpectedSendAndCall({
            sourceBlockchainID: sourceBlockchainID,
            originSenderAddress: originSenderAddress,
            recipient: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            amount: amount,
            payload: payload,
            gasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            targetHasCode: true,
            expectSuccess: false
        });

        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: sourceBlockchainID,
            originSenderAddress: originSenderAddress,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, message
        );
    }

    function testReceiveMismatchedSourceBlockchainID() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        bytes32 sourceBlockchainID = DEFAULT_DESTINATION_BLOCKCHAIN_ID;
        address originSenderAddress = address(this);
        bytes memory payload = hex"DEADBEEF";

        bytes32 wrongSourceBlockchainID = DEFAULT_SOURCE_BLOCKCHAIN_ID;
        assertNotEq(sourceBlockchainID, wrongSourceBlockchainID);
        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: wrongSourceBlockchainID,
            originSenderAddress: originSenderAddress,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        vm.expectRevert(_formatErrorMessage("mismatched source blockchain ID"));
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, message
        );
    }

    function testMultiHopSendSuccess() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        uint256 feeAmount = 1;
        uint256 bridgeAmount = amount - feeAmount;
        SendTokensInput memory input = SendTokensInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: feeAmount,
            secondaryFee: 0,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });
        _checkExpectedTeleporterCallsForSend(
            _createSingleHopTeleporterMessageInput(input, bridgeAmount)
        );

        vm.expectEmit(true, true, true, true, address(tokenSource));
        emit TokensRouted(_MOCK_MESSAGE_ID, input, bridgeAmount);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeMultiHopSendMessage({
                amount: amount,
                destinationBlockchainID: input.destinationBlockchainID,
                destinationBridgeAddress: input.destinationBridgeAddress,
                recipient: input.recipient,
                secondaryFee: input.primaryFee,
                secondaryGasLimit: input.requiredGasLimit,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
            })
        );
    }

    function testMultiHopSendInsufficientFees() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 2;
        _sendSingleHopSendSuccess(amount, 0);
        uint256 balanceBefore = tokenSource.bridgedBalances(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS
        );
        assertEq(balanceBefore, amount);

        // Fail due to insufficient amount to cover fees
        MultiHopSendMessage memory message = MultiHopSendMessage({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            amount: amount,
            secondaryFee: amount,
            secondaryGasLimit: 50_000,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        vm.expectRevert(_formatErrorMessage("insufficient amount to cover fees"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeMultiHopSendMessage({
                amount: amount,
                destinationBlockchainID: message.destinationBlockchainID,
                destinationBridgeAddress: message.destinationBridgeAddress,
                recipient: message.recipient,
                secondaryFee: message.secondaryFee,
                secondaryGasLimit: message.secondaryGasLimit,
                fallbackRecipient: message.fallbackRecipient
            })
        );

        // Make sure the bridge balance is still the same
        assertEq(
            tokenSource.bridgedBalances(
                DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS
            ),
            balanceBefore
        );
    }

    function testMultiHopCallSuccess() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        uint256 feeAmount = 1;
        uint256 bridgeAmount = amount - feeAmount;
        address originSenderAddress = address(this);
        SendAndCallInput memory input = SendAndCallInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: hex"65465465",
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS,
            primaryFee: feeAmount,
            secondaryFee: 0
        });
        _checkExpectedTeleporterCallsForSend(
            _createSingleHopCallTeleporterMessageInput(
                DEFAULT_DESTINATION_BLOCKCHAIN_ID, originSenderAddress, input, bridgeAmount
            )
        );

        vm.expectEmit(true, true, true, true, address(tokenSource));
        emit TokensAndCallRouted(_MOCK_MESSAGE_ID, input, bridgeAmount);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeMultiHopCallMessage({
                originSenderAddress: originSenderAddress,
                amount: amount,
                destinationBlockchainID: input.destinationBlockchainID,
                destinationBridgeAddress: input.destinationBridgeAddress,
                recipientContract: input.recipientContract,
                recipientPayload: input.recipientPayload,
                recipientGasLimit: input.recipientGasLimit,
                fallbackRecipient: input.fallbackRecipient,
                secondaryRequiredGasLimit: input.requiredGasLimit,
                secondaryFee: input.primaryFee
            })
        );
    }

    function _setUpExpectedSendAndCall(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        address recipient,
        uint256 amount,
        bytes memory payload,
        uint256 gasLimit,
        bool targetHasCode,
        bool expectSuccess
    ) internal virtual;

    function _setUpRegisteredDestination(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress
    ) internal virtual override {
        RegisterDestinationMessage memory payload = RegisterDestinationMessage({
            initialReserveImbalance: 0,
            tokenMultiplier: 1,
            multiplyOnReceive: false
        });
        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.REGISTER_DESTINATION,
            payload: abi.encode(payload)
        });
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            destinationBlockchainID, destinationBridgeAddress, abi.encode(message)
        );
    }

    function _createDefaultReceiveTokensInput() internal view returns (SendTokensInput memory) {
        return SendTokensInput({
            destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            destinationBridgeAddress: address(tokenSource),
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: 0,
            secondaryFee: 0,
            requiredGasLimit: 0,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });
    }

    function _getDefaultSourceBlockchainID() internal pure override returns (bytes32) {
        return DEFAULT_SOURCE_BLOCKCHAIN_ID;
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
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
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
        return bytes(string.concat("TeleporterTokenSource: ", message));
    }
}
