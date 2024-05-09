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
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";

abstract contract TeleporterTokenSourceTest is TeleporterTokenBridgeTest {
    TeleporterTokenSource public tokenSource;

    event CollateralAdded(
        bytes32 indexed destinationBlockchainID,
        address indexed destinationBridgeAddress,
        uint256 amount,
        uint256 remaining
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

    function testAddCollateralDestinationNotRegistered() public {
        vm.expectRevert(_formatErrorMessage("destination bridge not registered"));
        _addCollateral(DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, 100);
    }

    function testAddCollateralAlreadyCollateralized() public {
        _setUpRegisteredDestination(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, 0
        );
        vm.expectRevert(_formatErrorMessage("zero collateral needed"));
        _addCollateral(DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, 100);
    }

    function testAddCollateralPartialAmount() public {
        uint256 initialReserveImbalance = 100;
        uint256 collateralAmount = 50;
        _setUpRegisteredDestination(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, initialReserveImbalance
        );
        _setUpExpectedDeposit(collateralAmount);
        vm.expectEmit(true, true, true, true, address(tokenSource));
        emit CollateralAdded(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            collateralAmount,
            initialReserveImbalance - collateralAmount
        );
        _addCollateral(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, collateralAmount
        );
        (, uint256 updateReserveImbalance,,) = tokenSource.registeredDestinations(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS
        );
        assertEq(updateReserveImbalance, initialReserveImbalance - collateralAmount);
    }

    function testAddCollateralMoreThanFullAmount() public {
        uint256 initialReserveImbalance = 100;
        uint256 collateralAmount = 150;
        _setUpRegisteredDestination(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, initialReserveImbalance
        );
        _setUpExpectedDeposit(collateralAmount);
        vm.expectEmit(true, true, true, true, address(tokenSource));
        emit CollateralAdded(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            initialReserveImbalance,
            0
        );

        uint256 excessAmount = collateralAmount - initialReserveImbalance;
        _checkExpectedWithdrawal(address(this), excessAmount);

        _addCollateral(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, collateralAmount
        );
        (, uint256 updateReserveImbalance,,) = tokenSource.registeredDestinations(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS
        );
        assertEq(updateReserveImbalance, 0);
        assertTrue(address(this).balance > 0);
    }

    function testSendToUnregisteredDestination() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        vm.expectRevert(_formatErrorMessage("destination not registered"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendDestinationNotCollateralized() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        _setUpRegisteredDestination(
            input.destinationBlockchainID, input.destinationBridgeAddress, 1
        );
        vm.expectRevert(_formatErrorMessage("non-zero collateral needed for destination"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendZeroScaledAmount() public {
        _setUpRegisteredDestination(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, 0, 100_000, false
        );
        SendTokensInput memory input = _createDefaultSendTokensInput();
        uint256 amount = 10_000; // Amount is less than the token multiplier, so will be rounded down to zero.
        _setUpDeposit(amount);
        vm.expectRevert(_formatErrorMessage("zero scaled amount"));
        _send(input, amount);
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

    function testMultiHopInvalidDestinationSendsToFallback() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        // The multi-hop will not be routed to the OTHER_BLOCKCHAIN_ID destination since it
        // is not registered. Instead, the tokens are sent to the fallback recipient.
        _checkExpectedWithdrawal(DEFAULT_FALLBACK_RECIPIENT_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeMultiHopSendMessage({
                amount: amount,
                destinationBlockchainID: OTHER_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
                recipient: DEFAULT_RECIPIENT_ADDRESS,
                secondaryFee: 0,
                secondaryGasLimit: 500_000,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
            })
        );
    }

    function testMultiHopDestinationNotCollateralized() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        _setUpRegisteredDestination(OTHER_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, 100);

        // The multi-hop will not be routed to the OTHER_BLOCKCHAIN_ID destination since it is not yet
        // fully collateralized. Instead, the tokens are sent to the fallback recipient.
        _checkExpectedWithdrawal(DEFAULT_FALLBACK_RECIPIENT_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeMultiHopSendMessage({
                amount: amount,
                destinationBlockchainID: OTHER_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
                recipient: DEFAULT_RECIPIENT_ADDRESS,
                secondaryFee: 0,
                secondaryGasLimit: 500_000,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
            })
        );
    }

    function testMultiHopDestinationAmountScaledToZero() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 1_000;
        _sendSingleHopSendSuccess(amount, 0);

        _setUpRegisteredDestination(
            OTHER_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, 0, 100_000, false
        );

        // The multi-hop will not be routed to the OTHER_BLOCKCHAIN_ID destination since the token
        // amount would be scaled to zero. Instead, the tokens are sent to the fallback recipient.
        _checkExpectedWithdrawal(DEFAULT_FALLBACK_RECIPIENT_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeMultiHopSendMessage({
                amount: amount,
                destinationBlockchainID: OTHER_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
                recipient: DEFAULT_RECIPIENT_ADDRESS,
                secondaryFee: 0,
                secondaryGasLimit: 500_000,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
            })
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

    function testMultiHopCallInvalidDestination() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        // The multi-hop will not be routed to the OTHER_BLOCKCHAIN_ID destination since it is not registered.
        // Instead, the tokens are sent to the fallback recipient.
        _checkExpectedWithdrawal(DEFAULT_FALLBACK_RECIPIENT_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeMultiHopCallMessage({
                originSenderAddress: address(this),
                amount: amount,
                destinationBlockchainID: OTHER_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
                recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
                recipientPayload: new bytes(16),
                recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS,
                secondaryRequiredGasLimit: 500_000,
                secondaryFee: 0
            })
        );
    }

    function testMultiHopCallDestinationNotCollateralized() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        _setUpRegisteredDestination(OTHER_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, 100);

        // The multi-hop will not be routed to the OTHER_BLOCKCHAIN_ID destination since it is not yet
        // fully collateralized. Instead, the tokens are sent to the fallback recipient.
        _checkExpectedWithdrawal(DEFAULT_FALLBACK_RECIPIENT_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeMultiHopCallMessage({
                originSenderAddress: address(this),
                amount: amount,
                destinationBlockchainID: OTHER_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
                recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
                recipientPayload: new bytes(16),
                recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS,
                secondaryRequiredGasLimit: 500_000,
                secondaryFee: 0
            })
        );
    }

    function testMultiHopCallAmountScaledToZero() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 1_000;
        _sendSingleHopSendSuccess(amount, 0);

        _setUpRegisteredDestination(
            OTHER_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, 0, 100_000, false
        );

        // The multi-hop will not be routed to the OTHER_BLOCKCHAIN_ID destination since the token
        // amount would be scaled to zero. Instead, the tokens are sent to the fallback recipient.
        _checkExpectedWithdrawal(DEFAULT_FALLBACK_RECIPIENT_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            _encodeMultiHopCallMessage({
                originSenderAddress: address(this),
                amount: amount,
                destinationBlockchainID: OTHER_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
                recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
                recipientPayload: new bytes(16),
                recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS,
                secondaryRequiredGasLimit: 500_000,
                secondaryFee: 0
            })
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

    function testRegisterDestinationZeroBlockchainID() public {
        vm.expectRevert(_formatErrorMessage("zero destination blockchain ID"));
        _setUpRegisteredDestination(bytes32(0), DEFAULT_DESTINATION_ADDRESS, 0);
    }

    function testRegisterDestinationSameChain() public {
        bytes32 localBlockchainID = tokenSource.blockchainID();
        vm.expectRevert(_formatErrorMessage("cannot register bridge on same chain"));
        _setUpRegisteredDestination(localBlockchainID, DEFAULT_DESTINATION_ADDRESS, 0);
    }

    function testRegisterDestinationZeroAddress() public {
        vm.expectRevert(_formatErrorMessage("zero destination bridge address"));
        _setUpRegisteredDestination(DEFAULT_DESTINATION_BLOCKCHAIN_ID, address(0), 0);
    }

    function testRegisterDestinationZeroTokenMultiplier() public {
        vm.expectRevert(_formatErrorMessage("invalid token multiplier"));
        _setUpRegisteredDestination(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, 0, 0, false
        );
    }

    function testRegisterDestinationTokenMultiplierToLarge() public {
        vm.expectRevert(_formatErrorMessage("invalid token multiplier"));
        _setUpRegisteredDestination(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, 0, 1e18 + 1, false
        );
    }

    function testRegisterDestinationRoundUpCollateralNeeded() public {
        _setUpRegisteredDestination(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, 11, 10, true
        );
        (, uint256 collateralNeeded,,) = tokenSource.registeredDestinations(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS
        );
        assertEq(collateralNeeded, 2);
    }

    function testRegisterDestinationAlreadyReigstered() public {
        _setUpRegisteredDestination(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, 0
        );
        vm.expectRevert(_formatErrorMessage("destination already registered"));
        _setUpRegisteredDestination(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, 0
        );
    }

    function testSendScaledAmount() public {
        uint256 amount = 100;
        uint256 feeAmount = 1;
        uint256 bridgeAmount = amount - feeAmount;

        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = feeAmount;

        // Raw amount sent over wire should be multipled by 1e12.
        uint256 tokenMultiplier = 1e12;
        _setUpRegisteredDestination(
            input.destinationBlockchainID, input.destinationBridgeAddress, 0, tokenMultiplier, true
        );
        _setUpExpectedDeposit(amount);
        TeleporterMessageInput memory expectedMessage = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(feeToken), amount: input.primaryFee}),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeSingleHopSendMessage(
                bridgeAmount * tokenMultiplier, DEFAULT_RECIPIENT_ADDRESS
                )
        });
        _checkExpectedTeleporterCallsForSend(expectedMessage);
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, bridgeAmount * tokenMultiplier);
        _send(input, amount);

        // For another destination, the raw amount sent over the wire should be divided by 1e12.
        amount = 42 * tokenMultiplier;
        feeAmount = 0;
        input.destinationBlockchainID = OTHER_BLOCKCHAIN_ID;
        input.primaryFee = feeAmount;
        _setUpRegisteredDestination(
            input.destinationBlockchainID, input.destinationBridgeAddress, 0, tokenMultiplier, false
        );
        _setUpExpectedDeposit(amount);
        expectedMessage = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(feeToken), amount: input.primaryFee}),
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
        address originSenderAddress,
        address recipient,
        uint256 amount,
        bytes memory payload,
        uint256 gasLimit,
        bool targetHasCode,
        bool expectSuccess
    ) internal virtual;

    function _addCollateral(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        uint256 amount
    ) internal virtual;

    function _setUpDeposit(uint256 amount) internal virtual {}

    function _setUpRegisteredDestination(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        uint256 initialReserveImbalance
    ) internal virtual override {
        _setUpRegisteredDestination(
            destinationBlockchainID, destinationBridgeAddress, initialReserveImbalance, 1, true
        );
    }

    function _setUpRegisteredDestination(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        uint256 initialReserveImbalance,
        uint256 tokenMultiplier,
        bool multiplyOnSend
    ) internal virtual {
        RegisterDestinationMessage memory payload = RegisterDestinationMessage({
            initialReserveImbalance: initialReserveImbalance,
            tokenMultiplier: tokenMultiplier,
            multiplyOnSend: multiplyOnSend
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
