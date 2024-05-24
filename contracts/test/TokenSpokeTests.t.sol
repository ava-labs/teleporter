// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TokenBridgeTest} from "./TokenBridgeTests.t.sol";
import {TokenSpoke, IWarpMessenger} from "../src/TokenSpoke/TokenSpoke.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";
import {SendTokensInput, SendAndCallInput} from "../src/interfaces/ITokenBridge.sol";
import {ITeleporterMessenger} from "@teleporter/ITeleporterMessenger.sol";
import {TokenScalingUtils} from "../src/utils/TokenScalingUtils.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    BridgeMessageType,
    BridgeMessage,
    RegisterSpokeMessage
} from "../src/interfaces/ITokenBridge.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {ExampleERC20} from "../lib/teleporter/contracts/src/Mocks/ExampleERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

abstract contract TokenSpokeTest is TokenBridgeTest {
    using SafeERC20 for IERC20;

    TokenSpoke public tokenSpoke;
    uint8 public tokenDecimals = 18;

    function setUp() public virtual {
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeWithSelector(ITeleporterMessenger.sendCrossChainMessage.selector),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeCall(IWarpMessenger.getBlockchainID, ()),
            abi.encode(DEFAULT_TOKEN_SPOKE_BLOCKCHAIN_ID)
        );
        vm.expectCall(WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getBlockchainID, ()));

        _initMockTeleporterRegistry();

        vm.expectCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeCall(TeleporterRegistry(MOCK_TELEPORTER_REGISTRY_ADDRESS).latestVersion, ())
        );
    }

    function testZeroDestinationBlockchain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = bytes32(0);
        vm.expectRevert(_formatErrorMessage("zero destination blockchain ID"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testZeroDestinationBlockchainWithSendAndCall() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationBlockchainID = bytes32(0);
        vm.expectRevert(_formatErrorMessage("zero destination blockchain ID"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testZeroDestinationBridgeAddress() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(0);
        vm.expectRevert(_formatErrorMessage("zero destination bridge address"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testZeroDestinationBridgeAddressWithSendAndCall() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationBridgeAddress = address(0);
        vm.expectRevert(_formatErrorMessage("zero destination bridge address"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testInvalidSendingBackToSourceBlockchain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(this);
        vm.expectRevert(_formatErrorMessage("invalid destination bridge address"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testInvalidSendAndCallingBackToHubBlockchain() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationBridgeAddress = address(this);
        vm.expectRevert(_formatErrorMessage("invalid destination bridge address"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testNonZeroSecondaryFeeToHubBlockchain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.secondaryFee = 1;
        vm.expectRevert(_formatErrorMessage("non-zero secondary fee"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testNonZeroSecondaryFeeToHubBlockchainCall() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.secondaryFee = 1;
        vm.expectRevert(_formatErrorMessage("non-zero secondary fee"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testNonZeroMultiHopFallbackForSingleHop() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;
        vm.expectRevert(_formatErrorMessage("non-zero multi-hop fallback"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testNonZeroMultiHopFallbackForSingleHopCall() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;
        vm.expectRevert(_formatErrorMessage("non-zero multi-hop fallback"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendingToSameInstance() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = tokenSpoke.blockchainID();
        input.destinationBridgeAddress = address(tokenSpoke);
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;
        vm.expectRevert(_formatErrorMessage("invalid destination bridge address"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendingToOtherInstanceOnSameChain() public {
        _sendMultiHopSendSuccess(tokenSpoke.blockchainID(), 1e18, 999, 555);
    }

    function testSendAndCallingToSameInstance() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationBlockchainID = tokenSpoke.blockchainID();
        input.destinationBridgeAddress = address(tokenSpoke);
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;
        vm.expectRevert(_formatErrorMessage("invalid destination bridge address"));
        _sendAndCall(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendAndCallingToOtherInstanceOnSameChain() public {
        _sendMultiHopCallSuccess(tokenSpoke.blockchainID(), 1e18, 999, 555);
    }

    function testSendZeroAmountAfterRemoveScaling() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.primaryFee = 0;
        input.secondaryFee = 0;
        uint256 amount = 1;

        if (
            TokenScalingUtils.removeTokenScale(
                tokenSpoke.tokenMultiplier(), tokenSpoke.multiplyOnSpoke(), amount
            ) != 0
        ) {
            return;
        }

        _setUpExpectedDeposit(amount, input.primaryFee);
        vm.expectRevert(_formatErrorMessage("insufficient tokens to transfer"));
        _sendAndCall(input, amount);
    }

    function testSendToSameBlockchainDifferentDestination() public {
        uint256 amount = 2e15;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = tokenSpoke.blockchainID();
        // Set the desintation bridge address to an address different than the token spoke contract.
        input.destinationBridgeAddress = address(0x55);

        _sendSingleHopSendSuccess(amount, input.primaryFee);
    }

    function testSendMultiHopInsufficientAmountToCoverFees() public {
        SendTokensInput memory input = _createDefaultSendMultiHopInput();
        uint256 amount = 1;
        input.secondaryFee = 2;
        _setUpExpectedDeposit(amount, input.primaryFee);
        vm.expectRevert(_formatErrorMessage("insufficient tokens to transfer"));
        _send(input, amount);
    }

    function testSendMultiHopZeroMultiHopFallback() public {
        uint256 amount = 200_000;
        SendTokensInput memory input = _createDefaultSendMultiHopInput();
        input.multiHopFallback = address(0);
        vm.expectRevert(_formatErrorMessage("zero multi-hop fallback"));
        _send(input, amount);
    }

    function testSendAndCallMultiHopZeroMultiHopFallback() public {
        uint256 amount = 200_000;
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationBlockchainID = OTHER_BLOCKCHAIN_ID;
        input.multiHopFallback = address(0);
        vm.expectRevert(_formatErrorMessage("zero multi-hop fallback"));
        _sendAndCall(input, amount);
    }

    function testSendMultiHopSendSuccess() public {
        uint256 amount = 4e15;
        uint256 primaryFee = 5;
        uint256 secondaryFee = 2;
        _sendMultiHopSendSuccess(OTHER_BLOCKCHAIN_ID, amount, primaryFee, secondaryFee);
    }

    function testSendMultiHopCallSuccess() public {
        uint256 amount = 4e18;
        uint256 primaryFee = 5;
        uint256 secondaryFee = 1;
        _sendMultiHopCallSuccess(OTHER_BLOCKCHAIN_ID, amount, primaryFee, secondaryFee);
    }

    function testReceiveInvalidSourceBlockchainID() public {
        vm.expectRevert(_formatErrorMessage("invalid source blockchain ID"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSpoke.receiveTeleporterMessage(
            DEFAULT_TOKEN_SPOKE_BLOCKCHAIN_ID, DEFAULT_TOKEN_HUB_ADDRESS, new bytes(0)
        );
    }

    function testReceiveInvalidOriginSenderAddress() public {
        vm.expectRevert(_formatErrorMessage("invalid origin sender address"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSpoke.receiveTeleporterMessage(
            DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID, address(0), new bytes(0)
        );
    }

    function testReceiveInvalidMessage() public {
        vm.expectRevert();
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSpoke.receiveTeleporterMessage(
            DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID, DEFAULT_TOKEN_HUB_ADDRESS, bytes("test")
        );
    }

    function testReceiveWithdrawSuccess() public {
        uint256 amount = 200;
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        _checkExpectedWithdrawal(DEFAULT_RECIPIENT_ADDRESS, amount);
        uint256 initialSupply = _getTotalSupply();
        tokenSpoke.receiveTeleporterMessage(
            DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_HUB_ADDRESS,
            _encodeSingleHopSendMessage(amount, DEFAULT_RECIPIENT_ADDRESS)
        );
        assertEq(_getTotalSupply(), initialSupply + amount);
    }

    function testReceiveSendAndCallSuccess() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        bytes32 sourceBlockchainID = DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID;
        OriginSenderInfo memory originInfo;
        originInfo.senderAddress = address(this);
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

        uint256 initialSupply = _getTotalSupply();

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSpoke.receiveTeleporterMessage(
            DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID, DEFAULT_TOKEN_HUB_ADDRESS, message
        );

        assertEq(_getTotalSupply(), initialSupply + amount);
    }

    function testReceiveSendAndCallFailure() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        OriginSenderInfo memory originInfo;
        originInfo.bridgeAddress = address(tokenBridge);
        originInfo.senderAddress = address(this);
        _setUpExpectedSendAndCall({
            sourceBlockchainID: DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID,
            originInfo: originInfo,
            recipient: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            amount: amount,
            payload: payload,
            gasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            targetHasCode: true,
            expectSuccess: false
        });

        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID,
            originInfo: originInfo,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSpoke.receiveTeleporterMessage(
            DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID, DEFAULT_TOKEN_HUB_ADDRESS, message
        );
    }

    function testReceiveSendAndCallFailureNoCode() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        bytes32 sourceBlockchainID = DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID;
        OriginSenderInfo memory originInfo;
        originInfo.bridgeAddress = address(tokenBridge);
        originInfo.senderAddress = address(this);
        _setUpExpectedSendAndCall({
            sourceBlockchainID: sourceBlockchainID,
            originInfo: originInfo,
            recipient: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            amount: amount,
            payload: payload,
            gasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            targetHasCode: false,
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
        tokenSpoke.receiveTeleporterMessage(
            DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID, DEFAULT_TOKEN_HUB_ADDRESS, message
        );
    }

    function testReceiveSendAndCallFailureInsufficientGas() public {
        uint256 amount = 200;
        bytes memory payload = hex"DEADBEEF";
        uint256 gasLimit = 5_000_000;
        OriginSenderInfo memory originInfo;
        originInfo.bridgeAddress = address(tokenBridge);
        originInfo.senderAddress = address(this);

        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID,
            originInfo: originInfo,
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: gasLimit,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        _setUpMockMint(address(tokenSpoke), amount);
        vm.expectRevert("CallUtils: insufficient gas");
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSpoke.receiveTeleporterMessage{gas: gasLimit - 1}(
            DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID, DEFAULT_TOKEN_HUB_ADDRESS, message
        );
    }

    function testReceiveInvalidMessageType() public {
        vm.expectRevert(_formatErrorMessage("invalid message type"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSpoke.receiveTeleporterMessage(
            DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_HUB_ADDRESS,
            _encodeMultiHopSendMessage({
                amount: 1,
                destinationBlockchainID: DEFAULT_TOKEN_SPOKE_BLOCKCHAIN_ID,
                destinationBridgeAddress: DEFAULT_TOKEN_SPOKE_ADDRESS,
                recipient: DEFAULT_RECIPIENT_ADDRESS,
                secondaryFee: 0,
                secondaryGasLimit: 1_000,
                multiHopFallback: DEFAULT_MULTIHOP_FALLBACK_ADDRESS
            })
        );
    }

    function testCalculateNumWords() public {
        assertEq(tokenSpoke.calculateNumWords(0), 0);
        assertEq(tokenSpoke.calculateNumWords(1), 1);
        assertEq(tokenSpoke.calculateNumWords(32), 1);
        assertEq(tokenSpoke.calculateNumWords(33), 2);
        assertEq(tokenSpoke.calculateNumWords(64), 2);
        assertEq(tokenSpoke.calculateNumWords(65), 3);
    }

    function testRegisterWithHubSuccess() public {
        // Create a new instance that has not yet received any messages.
        tokenSpoke = _createNewSpokeInstance();
        tokenBridge = tokenSpoke;
        bridgedToken = IERC20(address(tokenSpoke));

        // Deploy a separate fee asset for registering with the token hub.
        ExampleERC20 separateFeeAsset = new ExampleERC20();
        uint256 feeAmount = 13;
        TeleporterFeeInfo memory feeInfo =
            TeleporterFeeInfo({feeTokenAddress: address(separateFeeAsset), amount: feeAmount});

        IERC20(separateFeeAsset).safeIncreaseAllowance(address(tokenBridge), feeAmount);
        vm.expectCall(
            address(separateFeeAsset),
            abi.encodeCall(IERC20.transferFrom, (address(this), address(tokenBridge), feeAmount))
        );

        BridgeMessage memory expectedBridgeMessage = BridgeMessage({
            messageType: BridgeMessageType.REGISTER_SPOKE,
            payload: abi.encode(
                RegisterSpokeMessage({
                    initialReserveImbalance: tokenSpoke.initialReserveImbalance(),
                    spokeTokenDecimals: tokenDecimals,
                    hubTokenDecimals: tokenHubDecimals
                })
                )
        });
        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: tokenSpoke.tokenHubBlockchainID(),
            destinationAddress: tokenSpoke.tokenHubAddress(),
            feeInfo: feeInfo,
            requiredGasLimit: tokenSpoke.REGISTER_SPOKE_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: abi.encode(expectedBridgeMessage)
        });
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput)),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        tokenSpoke.registerWithHub(feeInfo);
    }

    function testRegisterWithHubAlreadyRegistered() public {
        // Mock receiving a message from the hub so that the spoke knows
        // that it is registered already.
        uint256 amount = 1;
        _setUpMockMint(DEFAULT_RECIPIENT_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSpoke.receiveTeleporterMessage(
            tokenSpoke.tokenHubBlockchainID(),
            tokenSpoke.tokenHubAddress(),
            _encodeSingleHopSendMessage(amount, DEFAULT_RECIPIENT_ADDRESS)
        );

        vm.expectRevert(_formatErrorMessage("already registered"));
        tokenSpoke.registerWithHub(TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}));
    }

    function _sendMultiHopSendSuccess(
        bytes32 destinationBlockchainID,
        uint256 amount,
        uint256 primaryFee,
        uint256 secondaryFee
    ) internal {
        uint256 bridgeAmount = amount;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = destinationBlockchainID;
        input.primaryFee = primaryFee;
        input.secondaryFee = secondaryFee;
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;

        _setUpExpectedDeposit(amount, input.primaryFee);
        _checkExpectedTeleporterCallsForSend(
            _createMultiHopSendTeleporterMessageInput(input, bridgeAmount)
        );
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, bridgeAmount);
        _send(input, amount);
    }

    function _sendMultiHopCallSuccess(
        bytes32 destinationBlockchainID,
        uint256 amount,
        uint256 primaryFee,
        uint256 secondaryFee
    ) internal {
        uint256 bridgeAmount = amount;
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.destinationBlockchainID = destinationBlockchainID;
        input.primaryFee = primaryFee;
        input.secondaryFee = secondaryFee;
        input.multiHopFallback = DEFAULT_MULTIHOP_FALLBACK_ADDRESS;

        _setUpExpectedDeposit(amount, input.primaryFee);

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
        OriginSenderInfo memory originInfo,
        address recipient,
        uint256 amount,
        bytes memory payload,
        uint256 gasLimit,
        bool targetHasCode,
        bool expectSuccess
    ) internal virtual;

    // Spokes don't need to register supported spokes because they
    // only send messages to their configured token hub.
    function _setUpRegisteredSpoke(bytes32, address, uint256) internal virtual override {
        return;
    }

    function _createNewSpokeInstance() internal virtual returns (TokenSpoke);

    function _getTotalSupply() internal view virtual returns (uint256);

    function _createMultiHopSendTeleporterMessageInput(
        SendTokensInput memory input,
        uint256 bridgeAmount
    ) internal view returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: tokenSpoke.tokenHubBlockchainID(),
            destinationAddress: tokenSpoke.tokenHubAddress(),
            feeInfo: TeleporterFeeInfo({
                feeTokenAddress: address(bridgedToken),
                amount: input.primaryFee
            }),
            requiredGasLimit: tokenSpoke.MULTI_HOP_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: _encodeMultiHopSendMessage({
                amount: bridgeAmount,
                destinationBlockchainID: input.destinationBlockchainID,
                destinationBridgeAddress: input.destinationBridgeAddress,
                recipient: input.recipient,
                secondaryFee: input.secondaryFee,
                secondaryGasLimit: input.requiredGasLimit,
                multiHopFallback: input.multiHopFallback
            })
        });
    }

    function _createMultiHopCallTeleporterMessageInput(
        address originSenderAddress,
        SendAndCallInput memory input,
        uint256 bridgeAmount
    ) internal view returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: tokenSpoke.tokenHubBlockchainID(),
            destinationAddress: tokenSpoke.tokenHubAddress(),
            feeInfo: TeleporterFeeInfo({
                feeTokenAddress: address(bridgedToken),
                amount: input.primaryFee
            }),
            requiredGasLimit: tokenSpoke.MULTI_HOP_REQUIRED_GAS()
                + (
                    tokenSpoke.calculateNumWords(input.recipientPayload.length)
                        * tokenSpoke.MULTI_HOP_CALL_GAS_PER_WORD()
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
                multiHopFallback: input.multiHopFallback,
                secondaryRequiredGasLimit: input.requiredGasLimit,
                secondaryFee: input.secondaryFee
            })
        });
    }

    function _createDefaultSendTokensInput()
        internal
        view
        override
        returns (SendTokensInput memory)
    {
        return SendTokensInput({
            destinationBlockchainID: DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_TOKEN_HUB_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFeeTokenAddress: address(bridgedToken),
            primaryFee: 0,
            secondaryFee: 0,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            multiHopFallback: address(0)
        });
    }

    function _createDefaultSendMultiHopInput() internal view returns (SendTokensInput memory) {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = OTHER_BLOCKCHAIN_ID;
        input.multiHopFallback = DEFAULT_FALLBACK_RECIPIENT_ADDRESS;
        return input;
    }

    function _createDefaultSendAndCallInput()
        internal
        view
        override
        returns (SendAndCallInput memory)
    {
        return SendAndCallInput({
            destinationBlockchainID: DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_TOKEN_HUB_ADDRESS,
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

    function _getDefaultMessageSourceBlockchainID() internal pure override returns (bytes32) {
        return DEFAULT_TOKEN_SPOKE_BLOCKCHAIN_ID;
    }

    function _formatErrorMessage(string memory message)
        internal
        pure
        override
        returns (bytes memory)
    {
        return bytes(string.concat("TokenSpoke: ", message));
    }
}
