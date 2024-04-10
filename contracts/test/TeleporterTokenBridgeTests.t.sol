// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {
    ITeleporterTokenBridge,
    SendTokensInput,
    SendAndCallInput,
    BridgeMessageType,
    BridgeMessage,
    SingleHopSendMessage,
    SingleHopCallMessage,
    MultiHopSendMessage,
    MultiHopCallMessage
} from "../src/interfaces/ITeleporterTokenBridge.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

abstract contract TeleporterTokenBridgeTest is Test {
    using SafeERC20 for IERC20;

    bytes32 public constant DEFAULT_SOURCE_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    bytes32 public constant DEFAULT_DESTINATION_BLOCKCHAIN_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes32 public constant OTHER_BLOCKCHAIN_ID =
        bytes32(hex"9876987698769876987698769876987698769876987698769876987698769876");
    address public constant DEFAULT_DESTINATION_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address public constant TOKEN_SOURCE_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address public constant DEFAULT_RECIPIENT_ADDRESS = 0xABCDabcdABcDabcDaBCDAbcdABcdAbCdABcDABCd;
    address public constant DEFAULT_RECIPIENT_CONTRACT_ADDRESS =
        0xa83114A443dA1CecEFC50368531cACE9F37fCCcb;
    uint256 public constant DEFAULT_REQUIRED_GAS_LIMIT = 200_000;
    uint256 public constant DEFAULT_RECIPIENT_GAS_LIMIT = 100_000;
    address public constant DEFAULT_FALLBACK_RECIPIENT_ADDRESS =
        0xe69Ea1BAF997002602c0A3D451b2b5c9B7F8E6A1;
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
    address public constant NATIVE_MINTER_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000001);
    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS =
        0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
    address public constant MOCK_TELEPORTER_REGISTRY_ADDRESS =
        0xf9FA4a0c696b659328DDaaBCB46Ae4eBFC9e68e4;
    bytes32 internal constant _MOCK_MESSAGE_ID =
        bytes32(hex"1111111111111111111111111111111111111111111111111111111111111111");

    uint256 internal constant _DEFAULT_FEE_AMOUNT = 123456;
    uint256 internal constant _DEFAULT_TRANSFER_AMOUNT = 1e18;
    uint256 internal constant _DEFAULT_INITIAL_RESERVE_IMBALANCE = 1e18;
    uint8 internal constant _DEFAULT_DECIMALS_SHIFT = 1;
    uint256 internal constant _DEFAULT_TOKEN_MULTIPLIER = 10 ** _DEFAULT_DECIMALS_SHIFT;
    uint256 internal constant _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE = 1;

    ITeleporterTokenBridge public tokenBridge;
    IERC20 public feeToken;

    event TokensSent(
        bytes32 indexed teleporterMessageID,
        address indexed sender,
        SendTokensInput input,
        uint256 amount
    );
    event TokensAndCallSent(
        bytes32 indexed teleporterMessageID,
        address indexed sender,
        SendAndCallInput input,
        uint256 amount
    );
    event TokensWithdrawn(address indexed recipient, uint256 amount);
    event CallSucceeded(address indexed recipientContract, uint256 amount);
    event CallFailed(address indexed recipientContract, uint256 amount);
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

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

    function testSendZeroRecipient() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.recipient = address(0);
        feeToken.approve(address(tokenBridge), _DEFAULT_TRANSFER_AMOUNT);
        vm.expectRevert(_formatErrorMessage("zero recipient address"));
        _send(input, 0);
    }

    function testSendAndCallZeroRecipientContract() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.recipientContract = address(0);
        vm.expectRevert(_formatErrorMessage("zero recipient contract address"));
        _sendAndCall(input, 0);
    }

    function testSendZeroRequiredGasLimit() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.requiredGasLimit = 0;
        feeToken.approve(address(tokenBridge), _DEFAULT_TRANSFER_AMOUNT);
        vm.expectRevert(_formatErrorMessage("zero required gas limit"));
        _send(input, 0);
    }

    function testSendAndCallZeroRequiredGasLimit() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.requiredGasLimit = 0;
        vm.expectRevert(_formatErrorMessage("zero required gas limit"));
        _sendAndCall(input, 0);
    }

    function testSendAndCallZeroRecipientGasLimit() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.recipientGasLimit = 0;
        vm.expectRevert(_formatErrorMessage("zero recipient gas limit"));
        _sendAndCall(input, 0);
    }

    function testSendAndCallInvalidRecipientGasLimit() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.recipientGasLimit = input.requiredGasLimit + 1;
        vm.expectRevert(_formatErrorMessage("invalid recipient gas limit"));
        _sendAndCall(input, 0);
    }

    function testSendAndCallZeroFallbackRecipient() public {
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.fallbackRecipient = address(0);
        vm.expectRevert(_formatErrorMessage("zero fallback recipient address"));
        _sendAndCall(input, 0);
    }

    function testInsufficientAmountToCoverFees() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = 1;
        _setUpExpectedDeposit(input.primaryFee);
        vm.expectRevert(_formatErrorMessage("insufficient amount to cover fees"));
        _send(input, input.primaryFee);
    }

    function testSendWithFees() public {
        uint256 amount = 200_000;
        uint256 primaryFee = 100;
        _sendSingleHopSendSuccess(amount, primaryFee);
    }

    function testSendNoFees() public {
        uint256 amount = 200_000;
        uint256 primaryFee = 0;
        _sendSingleHopSendSuccess(amount, primaryFee);
    }

    function testSendAndCallWithFees() public {
        uint256 amount = 100_000;
        uint256 primaryFee = 10;
        _sendSingleHopCallSuccess(amount, primaryFee);
    }

    function _initMockTeleporterRegistry() internal {
        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(
                TeleporterRegistry(MOCK_TELEPORTER_REGISTRY_ADDRESS).latestVersion.selector
            ),
            abi.encode(1)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getVersionFromAddress.selector),
            abi.encode(1)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getAddressFromVersion.selector, (1)),
            abi.encode(MOCK_TELEPORTER_MESSENGER_ADDRESS)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getLatestTeleporter.selector),
            abi.encode(ITeleporterMessenger(MOCK_TELEPORTER_MESSENGER_ADDRESS))
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getVersionFromAddress.selector),
            abi.encode(1)
        );
    }

    function _send(SendTokensInput memory input, uint256 amount) internal virtual;

    function _sendAndCall(SendAndCallInput memory input, uint256 amount) internal virtual;

    function _sendSingleHopSendSuccess(uint256 amount, uint256 feeAmount) internal {
        uint256 bridgeAmount = amount - feeAmount;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = feeAmount;

        _setUpExpectedDeposit(amount);

        // Only tokens destinations scale tokens, so isReceive is always false here.
        uint256 scaledBridgedAmount = _scaleTokens(bridgeAmount, false);

        _checkExpectedTeleporterCallsForSend(
            _createSingleHopTeleporterMessageInput(input, scaledBridgedAmount), feeAmount
        );
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, scaledBridgedAmount);
        _send(input, amount);
    }

    function _sendSingleHopCallSuccess(uint256 amount, uint256 feeAmount) internal {
        uint256 bridgeAmount = amount - feeAmount;
        SendAndCallInput memory input = _createDefaultSendAndCallInput();
        input.primaryFee = feeAmount;

        // Only tokens destinations scale tokens, so isReceive is always false here.
        uint256 scaledBridgedAmount = _scaleTokens(bridgeAmount, false);

        _setUpExpectedDeposit(amount);
        _checkExpectedTeleporterCallsForSend(
            _createSingleHopCallTeleporterMessageInput(input, scaledBridgedAmount), feeAmount
        );
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensAndCallSent(_MOCK_MESSAGE_ID, address(this), input, scaledBridgedAmount);
        _sendAndCall(input, amount);
    }

    function _setUpExpectedDeposit(uint256 amount) internal virtual;

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal virtual;

    function _checkExpectedTeleporterCallsForSend(
        TeleporterMessageInput memory expectedMessageInput,
        uint256 expectedFeeAmount
    ) internal {
        if (expectedFeeAmount > 0) {
            vm.expectCall(
                address(feeToken),
                abi.encodeCall(
                    IERC20.allowance,
                    (address(tokenBridge), address(MOCK_TELEPORTER_MESSENGER_ADDRESS))
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

    // This function is overridden by TeleporterTokenDestinationTests
    function _scaleTokens(uint256 amount, bool) internal virtual returns (uint256) {
        return amount;
    }

    function _createSingleHopTeleporterMessageInput(
        SendTokensInput memory input,
        uint256 bridgeAmount
    ) internal view returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(feeToken), amount: input.primaryFee}),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeSingleHopSendMessage(bridgeAmount, input.recipient)
        });
    }

    function _createSingleHopCallTeleporterMessageInput(
        SendAndCallInput memory input,
        uint256 bridgeAmount
    ) internal view returns (TeleporterMessageInput memory) {
        return TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(feeToken), amount: input.primaryFee}),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeSingleHopCallMessage(
                bridgeAmount,
                input.recipientContract,
                input.recipientPayload,
                input.recipientGasLimit,
                input.fallbackRecipient
                )
        });
    }

    function _createDefaultSendTokensInput()
        internal
        pure
        virtual
        returns (SendTokensInput memory);

    function _createDefaultSendAndCallInput()
        internal
        pure
        virtual
        returns (SendAndCallInput memory);

    function _formatErrorMessage(string memory message)
        internal
        pure
        virtual
        returns (bytes memory);

    function _encodeSingleHopSendMessage(
        uint256 amount,
        address recipient
    ) internal pure returns (bytes memory) {
        return abi.encode(
            BridgeMessage({
                messageType: BridgeMessageType.SINGLE_HOP_SEND,
                amount: amount,
                payload: abi.encode(SingleHopSendMessage({recipient: recipient}))
            })
        );
    }

    function _encodeSingleHopCallMessage(
        uint256 amount,
        address recipientContract,
        bytes memory recipientPayload,
        uint256 recipientGasLimit,
        address fallbackRecipient
    ) internal pure returns (bytes memory) {
        return abi.encode(
            BridgeMessage({
                messageType: BridgeMessageType.SINGLE_HOP_CALL,
                amount: amount,
                payload: abi.encode(
                    SingleHopCallMessage({
                        recipientContract: recipientContract,
                        recipientPayload: recipientPayload,
                        recipientGasLimit: recipientGasLimit,
                        fallbackRecipient: fallbackRecipient
                    })
                    )
            })
        );
    }

    function _encodeMultiHopSendMessage(
        uint256 amount,
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        address recipient,
        uint256 secondaryFee,
        uint256 secondaryGasLimit
    ) internal pure returns (bytes memory) {
        return abi.encode(
            BridgeMessage({
                messageType: BridgeMessageType.MULTI_HOP_SEND,
                amount: amount,
                payload: abi.encode(
                    MultiHopSendMessage({
                        destinationBlockchainID: destinationBlockchainID,
                        destinationBridgeAddress: destinationBridgeAddress,
                        recipient: recipient,
                        secondaryFee: secondaryFee,
                        secondaryGasLimit: secondaryGasLimit
                    })
                    )
            })
        );
    }

    function _encodeMultiHopCallMessage(
        uint256 amount,
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        address recipientContract,
        bytes memory recipientPayload,
        uint256 requiredGasLimit,
        uint256 recipientGasLimit,
        address fallbackRecipient,
        uint256 secondaryFee
    ) internal pure returns (bytes memory) {
        return abi.encode(
            BridgeMessage({
                messageType: BridgeMessageType.MULTI_HOP_CALL,
                amount: amount,
                payload: abi.encode(
                    MultiHopCallMessage({
                        destinationBlockchainID: destinationBlockchainID,
                        destinationBridgeAddress: destinationBridgeAddress,
                        recipientContract: recipientContract,
                        recipientPayload: recipientPayload,
                        requiredGasLimit: requiredGasLimit,
                        recipientGasLimit: recipientGasLimit,
                        fallbackRecipient: fallbackRecipient,
                        secondaryFee: secondaryFee
                    })
                    )
            })
        );
    }
}
