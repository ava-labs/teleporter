// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {NativeTokenBridgeTest} from "./NativeTokenBridgeTest.t.sol";
import {ITokenSource} from "../ITokenSource.sol";
import {NativeTokenSource} from "../NativeTokenSource.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";

contract NativeTokenSourceTest is NativeTokenBridgeTest {
    NativeTokenSource public nativeTokenSource;

    event TransferToDestination(
        address indexed sender,
        address indexed recipient,
        bytes32 indexed teleporterMessageID,
        uint256 amount
    );
    event UnlockTokens(address indexed recipient, uint256 amount);
    event BurnTokens(uint256 amount);

    function setUp() public virtual override {
        NativeTokenBridgeTest.setUp();
        nativeTokenSource = new NativeTokenSource(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _DEFAULT_OWNER_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS
        );
    }

    function testTransferToDestination() public {
        vm.expectEmit(true, true, true, true, address(nativeTokenSource));
        emit TransferToDestination({
            sender: address(this),
            recipient: _DEFAULT_RECIPIENT,
            amount: _DEFAULT_TRANSFER_AMOUNT,
            teleporterMessageID: _MOCK_MESSAGE_ID
        });

        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            feeInfo: TeleporterFeeInfo({
                feeTokenAddress: address(mockERC20),
                amount: _DEFAULT_FEE_AMOUNT
            }),
            requiredGasLimit: nativeTokenSource.MINT_NATIVE_TOKENS_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
        });

        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        nativeTokenSource.transferToDestination{value: _DEFAULT_TRANSFER_AMOUNT}(
            _DEFAULT_RECIPIENT,
            TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: _DEFAULT_FEE_AMOUNT}),
            new address[](0)
        );

        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        nativeTokenSource.transferToDestination{value: _DEFAULT_TRANSFER_AMOUNT}(
            _DEFAULT_RECIPIENT,
            TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: _DEFAULT_FEE_AMOUNT}),
            new address[](0)
        );
    }

    function testUnlock() public {
        // Give the contract some tokens to burn.
        nativeTokenSource.transferToDestination{value: _DEFAULT_TRANSFER_AMOUNT * 2}(
            _DEFAULT_RECIPIENT,
            TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: _DEFAULT_FEE_AMOUNT}),
            new address[](0)
        );

        vm.expectEmit(true, true, true, true, address(nativeTokenSource));
        emit UnlockTokens(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(
                ITokenSource.SourceAction.Unlock,
                abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
            )
        );

        assertEq(_DEFAULT_TRANSFER_AMOUNT, _DEFAULT_RECIPIENT.balance);
    }

    function testBurnedTxFees() public {
        // Give the contract some tokens to burn.
        nativeTokenSource.transferToDestination{value: _DEFAULT_TRANSFER_AMOUNT}(
            _DEFAULT_RECIPIENT,
            TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: _DEFAULT_FEE_AMOUNT}),
            new address[](0)
        );

        uint256 burnedTxFees = 100;
        uint256 additionalTxFees = 50;
        assertEq(0, nativeTokenSource.destinationBurnedTotal());

        vm.expectEmit(true, true, true, true, address(nativeTokenSource));
        emit BurnTokens(burnedTxFees);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(ITokenSource.SourceAction.Burn, abi.encode(burnedTxFees))
        );

        assertEq(burnedTxFees, nativeTokenSource.destinationBurnedTotal());
        assertEq(burnedTxFees, nativeTokenSource.BURN_ADDRESS().balance);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(ITokenSource.SourceAction.Burn, abi.encode(burnedTxFees - 1))
        );

        assertEq(burnedTxFees, nativeTokenSource.destinationBurnedTotal());
        assertEq(burnedTxFees, nativeTokenSource.BURN_ADDRESS().balance);

        emit BurnTokens(additionalTxFees);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(ITokenSource.SourceAction.Burn, abi.encode(burnedTxFees + additionalTxFees))
        );

        assertEq(burnedTxFees + additionalTxFees, nativeTokenSource.destinationBurnedTotal());
        assertEq(burnedTxFees + additionalTxFees, nativeTokenSource.BURN_ADDRESS().balance);
    }

    function testZeroTeleporterAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");

        new NativeTokenSource(
            address(0x0),
            _DEFAULT_OWNER_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS
        );
    }

    function testZeroDestinationBlockchainID() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("zero destination blockchain ID"));

        new NativeTokenSource(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _DEFAULT_OWNER_ADDRESS,
            bytes32(0),
            _DEFAULT_OTHER_BRIDGE_ADDRESS
        );
    }

    function testSameBlockchainID() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("cannot bridge with same blockchain"));

        new NativeTokenSource(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _DEFAULT_OWNER_ADDRESS,
            _MOCK_BLOCKCHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS
        );
    }

    function testZeroDestinationContractAddress() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("zero destination contract address"));

        new NativeTokenSource(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _DEFAULT_OWNER_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            address(0x0)
        );
    }

    function testInvalidTeleporterAddress() public {
        vm.expectRevert("TeleporterUpgradeable: invalid Teleporter sender");

        vm.prank(address(0x123));
        nativeTokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(
                ITokenSource.SourceAction.Unlock,
                abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
            )
        );
    }

    function testInvalidDestinationBlockchain() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("invalid destination chain"));

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenSource.receiveTeleporterMessage(
            _MOCK_BLOCKCHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(
                ITokenSource.SourceAction.Unlock,
                abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
            )
        );
    }

    function testInvalidSenderContract() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("unauthorized sender"));

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            address(0x123),
            abi.encode(
                ITokenSource.SourceAction.Unlock,
                abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
            )
        );
    }

    function testInvalidRecipientAddress() public {
        vm.expectRevert(_formatNativeTokenSourceErrorMessage("zero recipient address"));

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(
                ITokenSource.SourceAction.Unlock, abi.encode(address(0x0), _DEFAULT_TRANSFER_AMOUNT)
            )
        );
    }

    function testZeroRecipient() public {
        vm.expectRevert(_formatNativeTokenSourceErrorMessage("zero recipient address"));

        nativeTokenSource.transferToDestination{value: _DEFAULT_TRANSFER_AMOUNT}(
            address(0x0),
            TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: _DEFAULT_FEE_AMOUNT}),
            new address[](0)
        );
    }

    function testInsufficientCollateral() public {
        vm.expectRevert(_formatNativeTokenSourceErrorMessage("insufficient collateral"));

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(
                ITokenSource.SourceAction.Unlock,
                abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
            )
        );
    }

    function testTransferZeroAmount() public {
        vm.expectRevert(_formatNativeTokenSourceErrorMessage("zero transfer value"));

        nativeTokenSource.transferToDestination{value: 0}(
            _DEFAULT_RECIPIENT,
            TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: _DEFAULT_FEE_AMOUNT}),
            new address[](0)
        );
    }

    function _formatNativeTokenSourceErrorMessage(string memory errorMessage)
        private
        pure
        returns (bytes memory)
    {
        return bytes(string.concat("NativeTokenSource: ", errorMessage));
    }

    function _formatTokenSourceErrorMessage(string memory errorMessage)
        private
        pure
        returns (bytes memory)
    {
        return bytes(string.concat("TokenSource: ", errorMessage));
    }
}
