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
    ITeleporterTokenBridge, SendTokensInput
} from "../src/interfaces/ITeleporterTokenBridge.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

abstract contract TeleporterTokenBridgeTest is Test {
    using SafeERC20 for IERC20;

    bytes32 public constant DEFAULT_SOURCE_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    bytes32 public constant DEFAULT_DESTINATION_BLOCKCHAIN_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    address public constant DEFAULT_DESTINATION_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address public constant TOKEN_SOURCE_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address public constant DEFAULT_RECIPIENT_ADDRESS = 0xABCDabcdABcDabcDaBCDAbcdABcdAbCdABcDABCd;
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
    address public constant NATIVE_MINTER_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000001);
    uint256 public constant DEFAULT_REQUIRED_GAS_LIMIT = 100_000;

    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS =
        0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
    address public constant MOCK_TELEPORTER_REGISTRY_ADDRESS =
        0xf9FA4a0c696b659328DDaaBCB46Ae4eBFC9e68e4;
    bytes32 internal constant _MOCK_MESSAGE_ID =
        bytes32(hex"1111111111111111111111111111111111111111111111111111111111111111");

    uint256 internal constant _DEFAULT_FEE_AMOUNT = 123456;
    uint256 internal constant _DEFAULT_TRANSFER_AMOUNT = 1e18;
    uint256 internal constant _DEFAULT_INITIAL_RESERVE_IMBALANCE = 1e18;
    uint256 internal constant _DEFAULT_DECIMALS_SHIFT = 1;
    uint256 internal constant _DEFAULT_TOKEN_MULTIPLIER = 10 ** _DEFAULT_DECIMALS_SHIFT;
    uint256 internal constant _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE = 1;

    ITeleporterTokenBridge public tokenBridge;
    IERC20 public feeToken;

    event SendTokens(
        bytes32 indexed teleporterMessageID,
        address indexed sender,
        SendTokensInput input,
        uint256 amount
    );

    event WithdrawTokens(address indexed recipient, uint256 amount);

    event Transfer(address indexed from, address indexed to, uint256 value);

    function testZeroDestinationBlockchainID() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        feeToken.approve(address(tokenBridge), _DEFAULT_TRANSFER_AMOUNT);
        input.destinationBlockchainID = bytes32(0);
        vm.expectRevert(_formatErrorMessage("zero destination blockchain ID"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testZeroDestinationBridge() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        feeToken.approve(address(tokenBridge), _DEFAULT_TRANSFER_AMOUNT);
        input.destinationBridgeAddress = address(0);
        vm.expectRevert(_formatErrorMessage("zero destination bridge address"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testZeroRecipient() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.recipient = address(0);
        feeToken.approve(address(tokenBridge), _DEFAULT_TRANSFER_AMOUNT);
        vm.expectRevert(_formatErrorMessage("zero recipient address"));
        _send(input, _DEFAULT_TRANSFER_AMOUNT);
    }

    function testSendWithFees() public {
        uint256 amount = 200;
        uint256 primaryFee = 100;
        _sendSuccess(amount, primaryFee);
    }

    function testSendNoFees() public {
        uint256 amount = 200;
        uint256 primaryFee = 0;
        _sendSuccess(amount, primaryFee);
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
            abi.encodeWithSelector(
                TeleporterRegistry.getVersionFromAddress.selector
            ),
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

    function _sendSuccess(uint256 amount, uint256 feeAmount) internal {
        uint256 bridgedAmount = amount - feeAmount;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = feeAmount;

        _setUpExpectedDeposit(amount);

        // Only tokens destinations scale tokens, so isReceive is always false here.
        uint256 scaledBridgedAmount = _scaleTokens(bridgedAmount, false);

        _checkExpectedTeleporterCalls(input, bridgedAmount);
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit SendTokens(_MOCK_MESSAGE_ID, address(this), input, scaledBridgedAmount);
        _send(input, amount);
    }

    function _setUpExpectedDeposit(uint256 amount) internal virtual;

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal virtual;

    function _checkExpectedTeleporterCalls(
        SendTokensInput memory input,
        uint256 bridgeAmount
    ) internal {
        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(feeToken), amount: input.primaryFee}),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeMessage(input, bridgeAmount)
        });

        if (input.primaryFee > 0) {
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

    // This function is overridden by NativeTokenDestinationTests
    function _scaleTokens(
        uint256 amount,
        bool
    ) internal virtual returns (uint256) {
        return amount;
    }

    function _encodeMessage(
        SendTokensInput memory input,
        uint256 amount
    ) internal virtual returns (bytes memory);

    function _createDefaultSendTokensInput()
        internal
        pure
        virtual
        returns (SendTokensInput memory);

    function _formatErrorMessage(string memory message)
        internal
        pure
        virtual
        returns (bytes memory);
}
