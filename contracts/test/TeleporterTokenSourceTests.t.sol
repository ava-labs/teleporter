// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenBridgeTest} from "./TeleporterTokenBridgeTests.t.sol";
import {TeleporterTokenSource, IWarpMessenger} from "../src/TeleporterTokenSource.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";
import {SendTokensInput} from "../src/interfaces/ITeleporterTokenBridge.sol";

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
        SendTokensInput memory input = _createDefaultReceiveTokensInput();
        vm.expectRevert(_formatErrorMessage("insufficient bridge balance"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, abi.encode(input, 1)
        );
    }

    function testReceiveInvalidDestinationBridgeAddress() public {
        // First send to destination blockchain to increase the bridge balance
        _sendSuccess(2, 0);

        SendTokensInput memory input = _createDefaultReceiveTokensInput();
        input.destinationBridgeAddress = address(0);
        vm.expectRevert(_formatErrorMessage("invalid bridge address"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, abi.encode(input, 1)
        );
    }

    function testReceiveWithdrawSuccess() public {
        uint256 amount = 2;
        _sendSuccess(amount, 0);

        uint256 feeAmount = 1;
        uint256 bridgedAmount = amount - feeAmount;
        SendTokensInput memory input = _createDefaultReceiveTokensInput();
        input.primaryFee = feeAmount;

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        vm.expectEmit(true, true, true, true, address(tokenSource));
        emit WithdrawTokens(DEFAULT_RECIPIENT_ADDRESS, bridgedAmount);
        _checkExpectedWithdrawal(DEFAULT_RECIPIENT_ADDRESS, bridgedAmount);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            abi.encode(input, bridgedAmount)
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
        uint256 bridgedAmount = amount - feeAmount;
        SendTokensInput memory input = SendTokensInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: feeAmount,
            secondaryFee: 0,
            requiredGasLimit: _expectedRequiredGasLimit()
        });

        _checkExpectedTeleporterCalls(input, bridgedAmount);

        vm.expectEmit(true, true, true, true, address(tokenSource));
        emit SendTokens(
            _MOCK_MESSAGE_ID, address(MOCK_TELEPORTER_MESSENGER_ADDRESS), input, bridgedAmount
        );

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            abi.encode(input, amount)
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
        SendTokensInput memory input = SendTokensInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: feeAmount,
            secondaryFee: 0,
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT
        });

        vm.expectRevert(_formatErrorMessage("insufficient amount to cover fees"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            abi.encode(input, amount)
        );

        // Make sure the bridge balance is still the same
        assertEq(
            tokenSource.bridgedBalances(
                DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS
            ),
            balanceBefore
        );
    }

    function _expectedRequiredGasLimit() internal view virtual override returns (uint256) {
        return DEFAULT_REQUIRED_GAS_LIMIT;
    }

    function _createDefaultSendTokensInput()
        internal
        view
        override
        returns (SendTokensInput memory)
    {
        return SendTokensInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: 0,
            secondaryFee: 0,
            requiredGasLimit: _expectedRequiredGasLimit()
        });
    }

    function _createDefaultReceiveTokensInput() internal view returns (SendTokensInput memory) {
        return SendTokensInput({
            destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            destinationBridgeAddress: address(tokenSource),
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: 0,
            secondaryFee: 0,
            requiredGasLimit: 0
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

    function _encodeMessage(
        SendTokensInput memory input,
        uint256 amount
    ) internal pure virtual override returns (bytes memory) {
        return abi.encode(input.recipient, amount);
    }
}
