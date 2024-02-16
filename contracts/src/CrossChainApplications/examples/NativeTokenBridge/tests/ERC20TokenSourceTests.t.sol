// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {NativeTokenBridgeTest} from "./NativeTokenBridgeTest.t.sol";
import {ITokenSource} from "../ITokenSource.sol";
import {ERC20TokenSource} from "../ERC20TokenSource.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";

contract ERC20TokenSourceTest is NativeTokenBridgeTest {
    ERC20TokenSource public erc20TokenSource;

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
        erc20TokenSource = new ERC20TokenSource(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _DEFAULT_OWNER_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(mockERC20)
        );
    }

    function testTransferToDestination() public {
        vm.expectEmit(true, true, true, true, address(erc20TokenSource));
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
            requiredGasLimit: erc20TokenSource.MINT_NATIVE_TOKENS_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
        });

        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        erc20TokenSource.transferToDestination(
            _DEFAULT_RECIPIENT,
            _DEFAULT_TRANSFER_AMOUNT + _DEFAULT_FEE_AMOUNT,
            _DEFAULT_FEE_AMOUNT,
            new address[](0)
        );
    }

    function testUnlock() public {
        // Give the contract some tokens to burn.
        erc20TokenSource.transferToDestination(
            _DEFAULT_RECIPIENT,
            _DEFAULT_TRANSFER_AMOUNT * 2 + _DEFAULT_FEE_AMOUNT,
            _DEFAULT_FEE_AMOUNT,
            new address[](0)
        );

        vm.expectEmit(true, true, true, true, address(erc20TokenSource));
        emit UnlockTokens(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(
                ITokenSource.SourceAction.Unlock,
                abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
            )
        );

        assertEq(_DEFAULT_TRANSFER_AMOUNT, mockERC20.balanceOf(_DEFAULT_RECIPIENT));
    }

    function testBurnedTxFees() public {
        // Give the contract some tokens to burn.
        erc20TokenSource.transferToDestination(
            _DEFAULT_RECIPIENT,
            _DEFAULT_TRANSFER_AMOUNT + _DEFAULT_FEE_AMOUNT,
            _DEFAULT_FEE_AMOUNT,
            new address[](0)
        );

        uint256 burnedTxFees = 100;
        uint256 additionalTxFees = 50;
        assertEq(0, erc20TokenSource.destinationBurnedTotal());

        vm.expectEmit(true, true, true, true, address(erc20TokenSource));
        emit BurnTokens(burnedTxFees);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(ITokenSource.SourceAction.Burn, abi.encode(burnedTxFees))
        );

        assertEq(burnedTxFees, erc20TokenSource.destinationBurnedTotal());
        assertEq(burnedTxFees, mockERC20.balanceOf(erc20TokenSource.BURN_ADDRESS()));

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(ITokenSource.SourceAction.Burn, abi.encode(burnedTxFees - 1))
        );

        assertEq(burnedTxFees, erc20TokenSource.destinationBurnedTotal());
        assertEq(burnedTxFees, mockERC20.balanceOf(erc20TokenSource.BURN_ADDRESS()));

        emit BurnTokens(additionalTxFees);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(ITokenSource.SourceAction.Burn, abi.encode(burnedTxFees + additionalTxFees))
        );

        assertEq(burnedTxFees + additionalTxFees, erc20TokenSource.destinationBurnedTotal());
        assertEq(
            burnedTxFees + additionalTxFees, mockERC20.balanceOf(erc20TokenSource.BURN_ADDRESS())
        );
    }

    function testZeroTeleporterAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");

        new ERC20TokenSource(
            address(0x0),
            _DEFAULT_OWNER_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(mockERC20)
        );
    }

    function testZeroDestinationBlockhainID() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("zero destination blockchain ID"));

        new ERC20TokenSource(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _DEFAULT_OWNER_ADDRESS,
            bytes32(0),
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(mockERC20)
        );
    }

    function testSameBlockchainID() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("cannot bridge with same blockchain"));

        new ERC20TokenSource(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _DEFAULT_OWNER_ADDRESS,
            _MOCK_BLOCKCHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(mockERC20)
        );
    }

    function testZeroDestinationContractAddress() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("zero destination contract address"));

        new ERC20TokenSource(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _DEFAULT_OWNER_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            address(0x0),
            address(mockERC20)
        );
    }

    function testInvalidTeleporterAddress() public {
        vm.expectRevert("TeleporterUpgradeable: invalid Teleporter sender");

        vm.prank(address(0x123));
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
        );
    }

    function testZeroERC20ContractAddress() public {
        vm.expectRevert(_formatERC20TokenSourceErrorMessage("zero ERC20 contract address"));

        new ERC20TokenSource(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _DEFAULT_OWNER_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(0x0)
        );
    }

    function testInvalidDestinationBlockchain() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("invalid destination chain"));

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
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
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            address(0x123),
            abi.encode(
                ITokenSource.SourceAction.Unlock,
                abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
            )
        );
    }

    function testInvalidRecipientAddress() public {
        vm.expectRevert(_formatERC20TokenSourceErrorMessage("zero recipient address"));

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(
                ITokenSource.SourceAction.Unlock, abi.encode(address(0x0), _DEFAULT_TRANSFER_AMOUNT)
            )
        );
    }

    function testZeroRecipient() public {
        vm.expectRevert(_formatERC20TokenSourceErrorMessage("zero recipient address"));

        erc20TokenSource.transferToDestination(
            address(0x0),
            _DEFAULT_TRANSFER_AMOUNT + _DEFAULT_FEE_AMOUNT,
            _DEFAULT_FEE_AMOUNT,
            new address[](0)
        );
    }

    function testTransferZeroAmount() public {
        vm.expectRevert(_formatERC20TokenSourceErrorMessage("zero transfer amount"));

        erc20TokenSource.transferToDestination(_DEFAULT_RECIPIENT, 0, 0, new address[](0));
    }

    function testInsufficientAdjustedAmount() public {
        vm.expectRevert(_formatERC20TokenSourceErrorMessage("insufficient adjusted amount"));

        erc20TokenSource.transferToDestination(
            _DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT, _DEFAULT_TRANSFER_AMOUNT, new address[](0)
        );
    }

    function _formatERC20TokenSourceErrorMessage(string memory errorMessage)
        private
        pure
        returns (bytes memory)
    {
        return bytes(string.concat("ERC20TokenSource: ", errorMessage));
    }

    function _formatTokenSourceErrorMessage(string memory errorMessage)
        private
        pure
        returns (bytes memory)
    {
        return bytes(string.concat("TokenSource: ", errorMessage));
    }
}
