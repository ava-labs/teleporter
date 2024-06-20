// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TokenBridgeTest} from "./TokenBridgeTests.t.sol";
import {TokenHome, IWarpMessenger} from "../src/TokenHome/TokenHome.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    BridgeMessageType,
    BridgeMessage,
    MultiHopSendMessage,
    RegisterRemoteMessage
} from "../src/interfaces/ITokenBridge.sol";
import {ITeleporterMessenger, TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {Math} from "@openzeppelin/contracts@4.8.1/utils/math/Math.sol";

abstract contract TokenHomeTest is TokenBridgeTest {
    TokenHome public tokenHome;

    event CollateralAdded(
        bytes32 indexed remoteBlockchainID, address indexed remoteBridgeAddress, uint256 amount, uint256 remaining
    );

    function setUp() public virtual {
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeWithSelector(ITeleporterMessenger.sendCrossChainMessage.selector),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID)
        );
        vm.expectCall(WARP_PRECOMPILE_ADDRESS, abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector));

        _initMockTeleporterRegistry();

        vm.expectCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry(MOCK_TELEPORTER_REGISTRY_ADDRESS).latestVersion.selector)
        );
    }

    function testAddCollateralRemoteNotRegistered() public {
        vm.expectRevert(_formatErrorMessage("remote not registered"));
        _addCollateral(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 100);
    }

    function testAddCollateralAlreadyCollateralized() public {
        _setUpRegisteredRemote(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 0);
        vm.expectRevert(_formatErrorMessage("zero collateral needed"));
        _addCollateral(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 100);
    }

    function testAddCollateralPartialAmount() public {
        uint256 initialReserveImbalance = 100;
        uint256 collateralAmount = 50;
        _setUpRegisteredRemote(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, initialReserveImbalance);
        _setUpExpectedDeposit(collateralAmount, 0);
        vm.expectEmit(true, true, true, true, address(tokenHome));
        emit CollateralAdded(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            collateralAmount,
            initialReserveImbalance - collateralAmount
        );
        _addCollateral(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, collateralAmount);
        (, uint256 updateReserveImbalance,,) =
            tokenHome.registeredRemotes(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS);
        assertEq(updateReserveImbalance, initialReserveImbalance - collateralAmount);
    }

    function testAddCollateralMoreThanFullAmount() public {
        uint256 initialReserveImbalance = 100;
        uint256 collateralAmount = 150;
        _setUpRegisteredRemote(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, initialReserveImbalance);
        _setUpExpectedDeposit(collateralAmount, 0);
        vm.expectEmit(true, true, true, true, address(tokenHome));
        emit CollateralAdded(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, initialReserveImbalance, 0);

        uint256 excessAmount = collateralAmount - initialReserveImbalance;
        _checkExpectedWithdrawal(address(this), excessAmount);

        _addCollateral(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, collateralAmount);
        (, uint256 updateReserveImbalance,,) =
            tokenHome.registeredRemotes(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS);
        assertEq(updateReserveImbalance, 0);
        assertTrue(address(this).balance > 0);
    }

    function testSendToUnregisteredRemote() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        vm.expectRevert(_formatErrorMessage("remote not registered"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendDestinationNotCollateralized() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        _setUpRegisteredRemote(input.destinationBlockchainID, input.destinationBridgeAddress, 1);
        vm.expectRevert(_formatErrorMessage("collateral needed for remote"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendZeroScaledAmount() public {
        _setUpRegisteredRemote(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 0, 100_000, false);
        SendTokensInput memory input = _createDefaultSendTokensInput();
        uint256 amount = 10_000; // Amount is less than the token multiplier, so will be rounded down to zero.
        _setUpDeposit(amount);
        vm.expectRevert(_formatErrorMessage("zero scaled amount"));
        _send(input, amount);
    }

    function testNonZeroFallbackRecipientForSingleHop() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;
        vm.expectRevert(_formatErrorMessage("non-zero multi-hop fallback"));
        _send(input, 100_000);
    }

    function testNonZeroFallbackRecipientForSingleHopCall() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;
        vm.expectRevert(_formatErrorMessage("non-zero multi-hop fallback"));
        _sendAndCall(input, 100_000);
    }

    function testNonZeroSecondaryFee() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.secondaryFee = 1;
        vm.expectRevert(_formatErrorMessage("non-zero secondary fee"));
        _send(input, 0);
    }

    function testNonZeroSecondaryFeeCall() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.secondaryFee = 1;
        vm.expectRevert(_formatErrorMessage("non-zero secondary fee"));
        _sendAndCall(input, 0);
    }

    function testReceiveInvalidMessage() public {
        vm.expectRevert();
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, bytes("test")
        );
    }

    function testReceiveFromNonRegisteredRemote() public {
        vm.expectRevert(_formatErrorMessage("remote not registered"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeSingleHopSendMessage(1, DEFAULT_RECIPIENT_ADDRESS)
        );
    }

    function testReceiveFromNonCollateralizedRemote() public {
        _setUpRegisteredRemote(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 100);
        vm.expectRevert(_formatErrorMessage("remote not collateralized"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeSingleHopSendMessage(1, DEFAULT_RECIPIENT_ADDRESS)
        );
    }

    function testReceiveInsufficientBridgeBalance() public {
        uint256 collateralAmount = 100;
        _setUpRegisteredRemote(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, collateralAmount);
        _setUpExpectedDeposit(collateralAmount, 0);
        _addCollateral(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, collateralAmount);
        vm.expectRevert(_formatErrorMessage("insufficient bridge balance"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeSingleHopSendMessage(1, DEFAULT_RECIPIENT_ADDRESS)
        );
    }

    function testReceiveWithdrawSuccess() public {
        uint256 amount = 200;
        uint256 feeAmount = 10;
        _sendSingleHopSendSuccess(amount, feeAmount);

        // Withdraw an amount less than bridged amount
        uint256 withdrawAmount = amount / 2;

        _checkExpectedWithdrawal(DEFAULT_RECIPIENT_ADDRESS, withdrawAmount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeSingleHopSendMessage(withdrawAmount, DEFAULT_RECIPIENT_ADDRESS)
        );

        // Make sure the bridge balance is correct. Only the remaining amount remains locked in the home
        // contract. The rest is withdrawn.
        assertEq(
            tokenHome.bridgedBalances(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS), withdrawAmount
        );
    }

    function testReceiveSendAndCallSuccess() public {
        // First send to ç to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        bytes32 sourceBlockchainID = DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID;
        OriginSenderInfo memory originInfo;
        originInfo.bridgeAddress = DEFAULT_TOKEN_REMOTE_ADDRESS;
        originInfo.senderAddress = address(this);
        bytes memory payload = hex"DEADBEEF";
        _setUpExpectedSendAndCall({
            sourceBlockchainID: sourceBlockchainID,
            originInfo: originInfo,
            recipient: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            amount: amount,
            payload: payload,
            gasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            targetHasCode: true,
            expectSuccess: true
        });

        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: sourceBlockchainID,
            originInfo: originInfo,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, originInfo.bridgeAddress, message);
    }

    function testReceiveSendAndCallFailure() public {
        // First send to remote instance to increase the bridge balance
        uint256 amount = 2;
        _sendSingleHopSendSuccess(amount, 0);

        bytes32 sourceBlockchainID = DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID;
        OriginSenderInfo memory originInfo;
        originInfo.bridgeAddress = DEFAULT_TOKEN_REMOTE_ADDRESS;
        originInfo.senderAddress = address(this);
        bytes memory payload = hex"DEADBEEF";
        _setUpExpectedSendAndCall({
            sourceBlockchainID: sourceBlockchainID,
            originInfo: originInfo,
            recipient: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            amount: amount,
            payload: payload,
            gasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            targetHasCode: true,
            expectSuccess: false
        });

        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: sourceBlockchainID,
            originInfo: originInfo,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, originInfo.bridgeAddress, message);
    }

    function testReceiveMismatchedSourceBlockchainID() public {
        // First send to remote instance to increase the bridge balance
        uint256 amount = 200_000;
        bytes memory payload = hex"DEADBEEF";
        _sendSingleHopSendSuccess(amount, 0);

        OriginSenderInfo memory originInfo;
        originInfo.bridgeAddress = DEFAULT_TOKEN_REMOTE_ADDRESS;
        originInfo.senderAddress = address(this);

        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: OTHER_BLOCKCHAIN_ID,
            originInfo: originInfo,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        vm.expectRevert(_formatErrorMessage("mismatched source blockchain ID"));
        tokenHome.receiveTeleporterMessage(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, message);
    }

    function testReceiveWrongOriginBridgeAddress() public {
        // First send to remote instance to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        address originBridgeAddress = DEFAULT_TOKEN_REMOTE_ADDRESS;
        bytes memory payload = hex"DEADBEEF";

        address wrongAddress = address(0x1);
        assertNotEq(originBridgeAddress, wrongAddress);

        OriginSenderInfo memory originInfo;
        originInfo.bridgeAddress = wrongAddress;
        originInfo.senderAddress = address(this);

        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            originInfo: originInfo,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        vm.expectRevert(_formatErrorMessage("mismatched origin sender address"));
        tokenHome.receiveTeleporterMessage(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, originBridgeAddress, message);
    }

    function testMultiHopInvalidDestinationSendsToFallback() public {
        // First send to remote instance to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        // The multi-hop will not be routed to the OTHER_BLOCKCHAIN_ID remote since it
        // is not registered. Instead, the tokens are sent to the multi-hop fallback.
        _checkExpectedWithdrawal(DEFAULT_MULTIHOP_FALLBACK_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeMultiHopSendMessage({
                amount: amount,
                destinationBlockchainID: OTHER_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_TOKEN_REMOTE_ADDRESS,
                recipient: DEFAULT_RECIPIENT_ADDRESS,
                secondaryFee: 0,
                secondaryGasLimit: 500_000,
                multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS
            })
        );
    }

    function testMultiHopDestinationNotCollateralized() public {
        // First send to remote instance to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        _setUpRegisteredRemote(OTHER_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 100);

        // The multi-hop will not be routed to the OTHER_BLOCKCHAIN_ID remote since it is not yet
        // fully collateralized. Instead, the tokens are sent to the multi-hop fallback.
        _checkExpectedWithdrawal(DEFAULT_MULTIHOP_FALLBACK_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeMultiHopSendMessage({
                amount: amount,
                destinationBlockchainID: OTHER_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_TOKEN_REMOTE_ADDRESS,
                recipient: DEFAULT_RECIPIENT_ADDRESS,
                secondaryFee: 0,
                secondaryGasLimit: 500_000,
                multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS
            })
        );
    }

    function testMultiHopDestinationAmountScaledToZero() public {
        // First send to remote instance to increase the bridge balance
        uint256 amount = 1_000;
        _sendSingleHopSendSuccess(amount, 0);

        _setUpRegisteredRemote(OTHER_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 0, 100_000, false);

        // The multi-hop will not be routed to the OTHER_BLOCKCHAIN_ID remote since the token
        // amount would be scaled to zero. Instead, the tokens are sent to the multi-hop fallback.
        _checkExpectedWithdrawal(DEFAULT_MULTIHOP_FALLBACK_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeMultiHopSendMessage({
                amount: amount,
                destinationBlockchainID: OTHER_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_TOKEN_REMOTE_ADDRESS,
                recipient: DEFAULT_RECIPIENT_ADDRESS,
                secondaryFee: 0,
                secondaryGasLimit: 500_000,
                multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS
            })
        );
    }

    function testMultiHopSendSuccess() public {
        // First send to remote instance to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        uint256 feeAmount = 1;
        uint256 bridgeAmount = amount - feeAmount;
        SendTokensInput memory input = SendTokensInput({
            destinationBlockchainID: DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_TOKEN_REMOTE_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFeeTokenAddress: address(bridgedToken),
            primaryFee: feeAmount,
            secondaryFee: 0,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS
        });
        _checkExpectedTeleporterCallsForSend(_createSingleHopTeleporterMessageInput(input, bridgeAmount));

        vm.expectEmit(true, true, true, true, address(tokenHome));
        emit TokensRouted(_MOCK_MESSAGE_ID, input, bridgeAmount);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeMultiHopSendMessage({
                amount: amount,
                destinationBlockchainID: input.destinationBlockchainID,
                destinationBridgeAddress: input.destinationBridgeAddress,
                recipient: input.recipient,
                secondaryFee: input.primaryFee,
                secondaryGasLimit: input.requiredGasLimit,
                multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS
            })
        );
    }

    function testMultiHopSendInsufficientFees() public {
        // First send to remote instance to increase the bridge balance
        uint256 amount = 2;
        _sendSingleHopSendSuccess(amount, 0);
        uint256 balanceBefore =
            tokenHome.bridgedBalances(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS);
        assertEq(balanceBefore, amount);

        // Fail due to insufficient amount to cover fees
        MultiHopSendMessage memory message = MultiHopSendMessage({
            destinationBlockchainID: DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_TOKEN_REMOTE_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            amount: amount,
            secondaryFee: amount,
            secondaryGasLimit: 50_000,
            multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS
        });

        vm.expectRevert(_formatErrorMessage("insufficient amount to cover fees"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeMultiHopSendMessage({
                amount: amount,
                destinationBlockchainID: message.destinationBlockchainID,
                destinationBridgeAddress: message.destinationBridgeAddress,
                recipient: message.recipient,
                secondaryFee: message.secondaryFee,
                secondaryGasLimit: message.secondaryGasLimit,
                multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS
            })
        );

        // Make sure the bridge balance is still the same
        assertEq(
            tokenHome.bridgedBalances(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS), balanceBefore
        );
    }

    function testMultiHopCallInvalidDestination() public {
        // First send to remote instance to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        // The multi-hop will not be routed to the OTHER_BLOCKCHAIN_ID remote since it is not registered.
        // Instead, the tokens are sent to the multi-hop fallback.
        _checkExpectedWithdrawal(DEFAULT_MULTIHOP_FALLBACK_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        address originSenderAddress = address(this);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeMultiHopCallMessage({
                originSenderAddress: originSenderAddress,
                amount: amount,
                destinationBlockchainID: OTHER_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_TOKEN_REMOTE_ADDRESS,
                recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
                recipientPayload: new bytes(16),
                recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS,
                multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS,
                secondaryRequiredGasLimit: 500_000,
                secondaryFee: 0
            })
        );
    }

    function testMultiHopCallDestinationNotCollateralized() public {
        // First send to remote instance to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        _setUpRegisteredRemote(OTHER_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 100);

        // The multi-hop will not be routed to the OTHER_BLOCKCHAIN_ID remote since it is not yet
        // fully collateralized. Instead, the tokens are sent to the multi-hop fallback.
        _checkExpectedWithdrawal(DEFAULT_MULTIHOP_FALLBACK_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        address originSenderAddress = address(this);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeMultiHopCallMessage({
                originSenderAddress: originSenderAddress,
                amount: amount,
                destinationBlockchainID: OTHER_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_TOKEN_REMOTE_ADDRESS,
                recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
                recipientPayload: new bytes(16),
                recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS,
                multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS,
                secondaryRequiredGasLimit: 500_000,
                secondaryFee: 0
            })
        );
    }

    function testMultiHopCallAmountScaledToZero() public {
        // First send to remote instance to increase the bridge balance
        uint256 amount = 1_000;
        _sendSingleHopSendSuccess(amount, 0);

        _setUpRegisteredRemote(OTHER_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 0, 100_000, false);

        // The multi-hop will not be routed to the OTHER_BLOCKCHAIN_ID remote since the token
        // amount would be scaled to zero. Instead, the tokens are sent to the multi-hop fallback.
        _checkExpectedWithdrawal(DEFAULT_MULTIHOP_FALLBACK_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        address originSenderAddress = address(this);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeMultiHopCallMessage({
                originSenderAddress: originSenderAddress,
                amount: amount,
                destinationBlockchainID: OTHER_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_TOKEN_REMOTE_ADDRESS,
                recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
                recipientPayload: new bytes(16),
                recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS,
                multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS,
                secondaryRequiredGasLimit: 500_000,
                secondaryFee: 0
            })
        );
    }

    function testMultiHopCallSuccess() public {
        // First send to remote instance to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        uint256 feeAmount = 1;
        uint256 bridgeAmount = amount - feeAmount;
        OriginSenderInfo memory originInfo;
        originInfo.bridgeAddress = DEFAULT_TOKEN_REMOTE_ADDRESS;
        originInfo.senderAddress = address(this);
        SendAndCallInput memory input = SendAndCallInput({
            destinationBlockchainID: DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_TOKEN_REMOTE_ADDRESS,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: hex"65465465",
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS,
            primaryFeeTokenAddress: address(bridgedToken),
            multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS,
            primaryFee: feeAmount,
            secondaryFee: 0
        });
        _checkExpectedTeleporterCallsForSend(
            _createSingleHopCallTeleporterMessageInput(
                DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, originInfo, input, bridgeAmount
            )
        );

        vm.expectEmit(true, true, true, true, address(tokenHome));
        emit TokensAndCallRouted(_MOCK_MESSAGE_ID, input, bridgeAmount);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_REMOTE_ADDRESS,
            _encodeMultiHopCallMessage({
                originSenderAddress: originInfo.senderAddress,
                amount: amount,
                destinationBlockchainID: input.destinationBlockchainID,
                destinationBridgeAddress: input.destinationBridgeAddress,
                recipientContract: input.recipientContract,
                recipientPayload: input.recipientPayload,
                recipientGasLimit: input.recipientGasLimit,
                fallbackRecipient: input.fallbackRecipient,
                multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS,
                secondaryRequiredGasLimit: input.requiredGasLimit,
                secondaryFee: input.primaryFee
            })
        );
    }

    function testRegisterRemoteZeroBlockchainID() public {
        vm.expectRevert(_formatErrorMessage("zero remote blockchain ID"));
        _setUpRegisteredRemote(bytes32(0), DEFAULT_TOKEN_REMOTE_ADDRESS, 0);
    }

    function testRegisterRemoteSameChain() public {
        bytes32 localBlockchainID = tokenHome.blockchainID();
        vm.expectRevert(_formatErrorMessage("cannot register remote on same chain"));
        _setUpRegisteredRemote(localBlockchainID, DEFAULT_TOKEN_REMOTE_ADDRESS, 0);
    }

    function testRegisterRemoteZeroAddress() public {
        vm.expectRevert(_formatErrorMessage("zero remote bridge address"));
        _setUpRegisteredRemote(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, address(0), 0);
    }

    function testRegisterRemoteAlreadyReigstered() public {
        _setUpRegisteredRemote(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 0);
        vm.expectRevert(_formatErrorMessage("remote already registered"));
        _setUpRegisteredRemote(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 0);
    }

    function testRegisterRemoteTokenDecimalsTooHigh() public {
        vm.expectRevert(_formatErrorMessage("remote token decimals too high"));
        _setUpRegisteredRemote(DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, 0, 1e19, true);
    }

    function testRegisterInvalidHomeTokenDecimals() public {
        vm.expectRevert(_formatErrorMessage("invalid home token decimals"));
        uint8 remoteTokenDecimals = uint8(18);
        RegisterRemoteMessage memory payload = RegisterRemoteMessage({
            initialReserveImbalance: 0,
            homeTokenDecimals: tokenHomeDecimals + 1,
            remoteTokenDecimals: uint8(remoteTokenDecimals)
        });
        BridgeMessage memory message =
            BridgeMessage({messageType: BridgeMessageType.REGISTER_REMOTE, payload: abi.encode(payload)});
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(
            DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID, DEFAULT_TOKEN_REMOTE_ADDRESS, abi.encode(message)
        );
    }

    function testSendScaledDownAmount() public {
        uint256 amount = 100;
        uint256 feeAmount = 1;

        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = feeAmount;

        uint256 tokenMultiplier = 1e2;

        // For another destination, the raw amount sent over the wire should be divided by 1e2.
        amount = 42 * tokenMultiplier;
        feeAmount = 0;
        input.destinationBlockchainID = OTHER_BLOCKCHAIN_ID;
        input.primaryFee = feeAmount;
        _setUpRegisteredRemote(input.destinationBlockchainID, input.destinationBridgeAddress, 0, tokenMultiplier, false);
        _setUpExpectedDeposit(amount, input.primaryFee);
        TeleporterMessageInput memory expectedMessage = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(bridgedToken), amount: input.primaryFee}),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeSingleHopSendMessage(amount / tokenMultiplier, DEFAULT_RECIPIENT_ADDRESS)
        });
        _checkExpectedTeleporterCallsForSend(expectedMessage);
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, amount / tokenMultiplier);
        _send(input, amount);
    }

    function _setUpExpectedSendAndCall(
        bytes32 sourceBlockchainID,
        OriginSenderInfo memory originInfo,
        address recipient,
        uint256 amount,
        bytes memory payload,
        uint256 gasLimit,
        bool targetHasCode,
        bool expectSuccess
    ) internal virtual;

    function _addCollateral(bytes32 remoteBlockchainID, address remoteBridgeAddress, uint256 amount) internal virtual;

    /**
     * @notice Set up necessary calls to deposit funds and call `send` or `sendAndCall`
     * on the token home contract.
     * Different from `_setUpExpectedDeposit` since the send that follows isn't
     * expected to succeed. So `setUpDeposit` does not check for expected emit events.
     */
    function _setUpDeposit(uint256 amount) internal virtual {}

    function _setUpRegisteredRemote(
        bytes32 remoteBlockchainID,
        address remoteBridgeAddress,
        uint256 initialReserveImbalance
    ) internal virtual override {
        _setUpRegisteredRemote(remoteBlockchainID, remoteBridgeAddress, initialReserveImbalance, 1, true);
    }

    function _setUpRegisteredRemote(
        bytes32 remoteBlockchainID,
        address remoteBridgeAddress,
        uint256 initialReserveImbalance,
        uint256 tokenMultiplier,
        bool multiplyOnRemote
    ) internal virtual {
        uint8 remoteTokenDecimals = uint8(
            multiplyOnRemote
                ? tokenHomeDecimals + Math.log10(tokenMultiplier)
                : tokenHomeDecimals - Math.log10(tokenMultiplier)
        );
        RegisterRemoteMessage memory payload = RegisterRemoteMessage({
            initialReserveImbalance: initialReserveImbalance,
            homeTokenDecimals: tokenHomeDecimals,
            remoteTokenDecimals: remoteTokenDecimals
        });
        BridgeMessage memory message =
            BridgeMessage({messageType: BridgeMessageType.REGISTER_REMOTE, payload: abi.encode(payload)});
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHome.receiveTeleporterMessage(remoteBlockchainID, remoteBridgeAddress, abi.encode(message));
    }

    function _createDefaultSendTokensInput() internal view override returns (SendTokensInput memory) {
        return SendTokensInput({
            destinationBlockchainID: DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_TOKEN_REMOTE_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFeeTokenAddress: address(bridgedToken),
            primaryFee: 0,
            secondaryFee: 0,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            multiHopFallback: address(0)
        });
    }

    function _createDefaultSendAndCallInput() internal view override returns (SendAndCallInput memory) {
        return SendAndCallInput({
            destinationBlockchainID: DEFAULT_TOKEN_REMOTE_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_TOKEN_REMOTE_ADDRESS,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: new bytes(16),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS,
            multiHopFallback: address(0),
            primaryFeeTokenAddress: address(bridgedToken),
            primaryFee: 0,
            secondaryFee: 0
        });
    }

    function _createDefaultReceiveTokensInput() internal view returns (SendTokensInput memory) {
        return SendTokensInput({
            destinationBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            destinationBridgeAddress: address(tokenHome),
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFeeTokenAddress: address(bridgedToken),
            primaryFee: 0,
            secondaryFee: 0,
            requiredGasLimit: 0,
            multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS
        });
    }

    function _getDefaultMessageSourceBlockchainID() internal pure override returns (bytes32) {
        return DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID;
    }

    function _formatErrorMessage(string memory message) internal pure override returns (bytes memory) {
        return bytes(string.concat("TokenHome: ", message));
    }
}
