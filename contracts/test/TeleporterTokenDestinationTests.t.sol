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
import {TokenScalingUtils} from "../src/utils/TokenScalingUtils.sol";

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

    function testInvalidSendAndCallingBackToSourceBlockchain() public {
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

        if (
            TokenScalingUtils.scaleTokens(
                tokenDestination.tokenMultiplier(),
                tokenDestination.multiplyOnReceive(),
                amount,
                false
            ) != 0
        ) {
            return;
        }

        _setUpExpectedDeposit(amount);
        vm.expectRevert(_formatErrorMessage("insufficient tokens to transfer"));
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

    function testSendZeroDestinationBlockchainID() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        feeToken.approve(address(tokenBridge), _DEFAULT_TRANSFER_AMOUNT);
        input.destinationBlockchainID = bytes32(0);
        vm.expectRevert(_formatErrorMessage("zero destination blockchain ID"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendZeroDestinationBridge() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        feeToken.approve(address(tokenBridge), _DEFAULT_TRANSFER_AMOUNT);
        input.destinationBridgeAddress = address(0);
        vm.expectRevert(_formatErrorMessage("zero destination bridge address"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
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
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        _checkExpectedWithdrawal(DEFAULT_RECIPIENT_ADDRESS, amount);
        uint256 initialSupply = _getTotalSupply();
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeSingleHopSendMessage(amount, DEFAULT_RECIPIENT_ADDRESS)
        );
        assertEq(_getTotalSupply(), initialSupply + amount);
    }

    function testReceiveSendAndCallSuccess() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        bytes32 sourceBlockchainID = DEFAULT_SOURCE_BLOCKCHAIN_ID;
        address originSenderAddress = address(this);
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

        uint256 initialSupply = _getTotalSupply();

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, message
        );

        assertEq(_getTotalSupply(), initialSupply + amount);
    }

    function testReceiveSendAndCallFailure() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        bytes32 sourceBlockchainID = DEFAULT_SOURCE_BLOCKCHAIN_ID;
        address originSenderAddress = address(this);
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
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, message
        );
    }

    function testReceiveSendAndCallFailureNoCode() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        bytes32 sourceBlockchainID = DEFAULT_SOURCE_BLOCKCHAIN_ID;
        address originSenderAddress = address(this);
        _setUpExpectedSendAndCall({
            sourceBlockchainID: sourceBlockchainID,
            originSenderAddress: originSenderAddress,
            recipient: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            amount: amount,
            payload: payload,
            gasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            targetHasCode: false,
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
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, message
        );
    }

    function testReceiveSendAndCallFailureInsufficientGas() public {
        uint256 amount = 200;
        bytes memory payload = hex"DEADBEEF";
        uint256 gasLimit = 5_000_000;
        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            originSenderAddress: address(this),
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: gasLimit,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        _setUpMockMint(address(tokenDestination), amount);
        vm.expectRevert("CallUtils: insufficient gas");
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
                secondaryGasLimit: 1_000,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
            })
        );
    }

    function testCalculateNumWords() public {
        assertEq(tokenDestination.calculateNumWords(0), 0);
        assertEq(tokenDestination.calculateNumWords(1), 1);
        assertEq(tokenDestination.calculateNumWords(32), 1);
        assertEq(tokenDestination.calculateNumWords(33), 2);
        assertEq(tokenDestination.calculateNumWords(64), 2);
        assertEq(tokenDestination.calculateNumWords(65), 3);
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
            _createMultiHopSendTeleporterMessageInput(input, bridgeAmount)
        );
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, bridgeAmount);
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
        address originSenderAddress = address(this);

        _checkExpectedTeleporterCallsForSend(
            _createMultiHopCallTeleporterMessageInput(originSenderAddress, input, bridgeAmount)
        );
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensAndCallSent(_MOCK_MESSAGE_ID, originSenderAddress, input, bridgeAmount);
        _sendAndCall(input, amount);
    }

    function _setUpMockMint(address recipient, uint256 amount) internal virtual;

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

    // Destinations don't need to registered supported destinations because they
    // only send messages to their configured token source.
    function _setUpRegisteredDestination(bytes32, address) internal virtual override {
        return;
    }

    function _getTotalSupply() internal view virtual returns (uint256);

    function _createMultiHopSendTeleporterMessageInput(
        SendTokensInput memory input,
        uint256 bridgeAmount
    ) internal view returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: tokenDestination.sourceBlockchainID(),
            destinationAddress: tokenDestination.tokenSourceAddress(),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(feeToken), amount: input.primaryFee}),
            requiredGasLimit: tokenDestination.MULTI_HOP_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: _encodeMultiHopSendMessage({
                amount: bridgeAmount,
                destinationBlockchainID: input.destinationBlockchainID,
                destinationBridgeAddress: input.destinationBridgeAddress,
                recipient: input.recipient,
                secondaryFee: input.secondaryFee,
                secondaryGasLimit: input.requiredGasLimit,
                fallbackRecipient: input.fallbackRecipient
            })
        });
    }

    function _createMultiHopCallTeleporterMessageInput(
        address originSenderAddress,
        SendAndCallInput memory input,
        uint256 bridgeAmount
    ) internal view returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: tokenDestination.sourceBlockchainID(),
            destinationAddress: tokenDestination.tokenSourceAddress(),
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(feeToken), amount: input.primaryFee}),
            requiredGasLimit: tokenDestination.MULTI_HOP_REQUIRED_GAS()
                + (
                    tokenDestination.calculateNumWords(input.recipientPayload.length)
                        * tokenDestination.MULTI_HOP_CALL_GAS_PER_WORD()
                ),
            allowedRelayerAddresses: new address[](0),
            message: _encodeMultiHopCallMessage({
                originSenderAddress: originSenderAddress,
                amount: bridgeAmount,
                destinationBlockchainID: input.destinationBlockchainID,
                destinationBridgeAddress: input.destinationBridgeAddress,
                recipientContract: input.recipientContract,
                recipientPayload: input.recipientPayload,
                recipientGasLimit: input.recipientGasLimit,
                fallbackRecipient: input.fallbackRecipient,
                secondaryRequiredGasLimit: input.requiredGasLimit,
                secondaryFee: input.secondaryFee
            })
        });
    }

    function _getDefaultSourceBlockchainID() internal pure override returns (bytes32) {
        return DEFAULT_DESTINATION_BLOCKCHAIN_ID;
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
