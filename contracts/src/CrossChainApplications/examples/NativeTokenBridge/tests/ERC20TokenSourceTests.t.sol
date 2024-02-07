// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {NativeTokenBridgeTest} from "./NativeTokenBridgeTest.t.sol";
import {ITokenSource} from "../ITokenSource.sol";
import {
    ERC20TokenSource,
    TeleporterMessageInput,
    TeleporterFeeInfo,
    ITeleporterMessenger
} from "../ERC20TokenSource.sol";

contract ERC20TokenSourceTest is NativeTokenBridgeTest {
    ERC20TokenSource public erc20TokenSource;
    uint256 public constant DEFAULT_TOKEN_MULTIPLIER = 2;

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
        erc20TokenSource = new ERC20TokenSource({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            destinationBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenDestinationAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            erc20ContractAddress_: address(mockERC20),
            tokenMultiplier_: DEFAULT_TOKEN_MULTIPLIER,
            multiplyOnSend_: true
        });
    }

    function testTransferToDestination() public {
        vm.expectEmit(true, true, true, true, address(erc20TokenSource));
        emit TransferToDestination({
            sender: address(this),
            recipient: _DEFAULT_RECIPIENT,
            amount: _DEFAULT_TRANSFER_AMOUNT * DEFAULT_TOKEN_MULTIPLIER,
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
            message: abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT * DEFAULT_TOKEN_MULTIPLIER)
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

    function testTransferToDestinationWithMultiplier() public {
        erc20TokenSource = new ERC20TokenSource({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            destinationBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenDestinationAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            erc20ContractAddress_: address(mockERC20),
            tokenMultiplier_: DEFAULT_TOKEN_MULTIPLIER,
            multiplyOnSend_: false
        });

        vm.expectEmit(true, true, true, true, address(erc20TokenSource));
        emit TransferToDestination({
            sender: address(this),
            recipient: _DEFAULT_RECIPIENT,
            amount: _DEFAULT_TRANSFER_AMOUNT / DEFAULT_TOKEN_MULTIPLIER,
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
            message: abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT / DEFAULT_TOKEN_MULTIPLIER)
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
            _DEFAULT_TRANSFER_AMOUNT * DEFAULT_TOKEN_MULTIPLIER + _DEFAULT_FEE_AMOUNT,
            _DEFAULT_FEE_AMOUNT,
            new address[](0)
        );

        vm.expectEmit(true, true, true, true, address(erc20TokenSource));
        emit UnlockTokens(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT / DEFAULT_TOKEN_MULTIPLIER);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(
                ITokenSource.SourceAction.Unlock,
                abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
            )
        );

        assertEq(
            _DEFAULT_TRANSFER_AMOUNT / DEFAULT_TOKEN_MULTIPLIER,
            mockERC20.balanceOf(_DEFAULT_RECIPIENT)
        );
    }

    function testUnlockWithMultiplier() public {
        erc20TokenSource = new ERC20TokenSource({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            destinationBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenDestinationAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            erc20ContractAddress_: address(mockERC20),
            tokenMultiplier_: DEFAULT_TOKEN_MULTIPLIER,
            multiplyOnSend_: false
        });

        // Give the contract some tokens to burn.
        erc20TokenSource.transferToDestination(
            _DEFAULT_RECIPIENT,
            _DEFAULT_TRANSFER_AMOUNT * DEFAULT_TOKEN_MULTIPLIER + _DEFAULT_FEE_AMOUNT,
            _DEFAULT_FEE_AMOUNT,
            new address[](0)
        );

        vm.expectEmit(true, true, true, true, address(erc20TokenSource));
        emit UnlockTokens(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT * DEFAULT_TOKEN_MULTIPLIER);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(
                ITokenSource.SourceAction.Unlock,
                abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
            )
        );

        assertEq(
            _DEFAULT_TRANSFER_AMOUNT * DEFAULT_TOKEN_MULTIPLIER,
            mockERC20.balanceOf(_DEFAULT_RECIPIENT)
        );
    }

    function testBurnedTxFees() public {
        // Give the contract some tokens to burn.
        erc20TokenSource.transferToDestination(
            _DEFAULT_RECIPIENT,
            _DEFAULT_TRANSFER_AMOUNT + _DEFAULT_FEE_AMOUNT,
            _DEFAULT_FEE_AMOUNT,
            new address[](0)
        );

        uint256 burnedTxFees = 99;
        uint256 additionalTxFees = 50;
        assertEq(0, erc20TokenSource.destinationBurnedTotal());

        vm.expectEmit(true, true, true, true, address(erc20TokenSource));
        emit BurnTokens(burnedTxFees / DEFAULT_TOKEN_MULTIPLIER);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(ITokenSource.SourceAction.Burn, abi.encode(burnedTxFees))
        );

        assertEq(burnedTxFees / DEFAULT_TOKEN_MULTIPLIER, erc20TokenSource.destinationBurnedTotal());
        assertEq(
            burnedTxFees / DEFAULT_TOKEN_MULTIPLIER,
            mockERC20.balanceOf(erc20TokenSource.BURN_ADDRESS())
        );

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(ITokenSource.SourceAction.Burn, abi.encode(burnedTxFees - 1))
        );

        assertEq(burnedTxFees / DEFAULT_TOKEN_MULTIPLIER, erc20TokenSource.destinationBurnedTotal());
        assertEq(
            burnedTxFees / DEFAULT_TOKEN_MULTIPLIER,
            mockERC20.balanceOf(erc20TokenSource.BURN_ADDRESS())
        );

        emit BurnTokens(additionalTxFees / DEFAULT_TOKEN_MULTIPLIER);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(ITokenSource.SourceAction.Burn, abi.encode(burnedTxFees + additionalTxFees))
        );

        assertEq(
            burnedTxFees / DEFAULT_TOKEN_MULTIPLIER + additionalTxFees / DEFAULT_TOKEN_MULTIPLIER,
            erc20TokenSource.destinationBurnedTotal()
        );
        assertEq(
            burnedTxFees / DEFAULT_TOKEN_MULTIPLIER + additionalTxFees / DEFAULT_TOKEN_MULTIPLIER,
            mockERC20.balanceOf(erc20TokenSource.BURN_ADDRESS())
        );
    }

    function testBurnedTxFeesWithMultiplier() public {
        erc20TokenSource = new ERC20TokenSource({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            destinationBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenDestinationAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            erc20ContractAddress_: address(mockERC20),
            tokenMultiplier_: DEFAULT_TOKEN_MULTIPLIER,
            multiplyOnSend_: false
        });

        // Give the contract some tokens to burn.
        erc20TokenSource.transferToDestination(
            _DEFAULT_RECIPIENT,
            _DEFAULT_TRANSFER_AMOUNT + _DEFAULT_FEE_AMOUNT,
            _DEFAULT_FEE_AMOUNT,
            new address[](0)
        );

        uint256 burnedTxFees = 99;
        uint256 additionalTxFees = 50;
        assertEq(0, erc20TokenSource.destinationBurnedTotal());

        vm.expectEmit(true, true, true, true, address(erc20TokenSource));
        emit BurnTokens(burnedTxFees * DEFAULT_TOKEN_MULTIPLIER);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(ITokenSource.SourceAction.Burn, abi.encode(burnedTxFees))
        );

        assertEq(burnedTxFees * DEFAULT_TOKEN_MULTIPLIER, erc20TokenSource.destinationBurnedTotal());
        assertEq(
            burnedTxFees * DEFAULT_TOKEN_MULTIPLIER,
            mockERC20.balanceOf(erc20TokenSource.BURN_ADDRESS())
        );

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(ITokenSource.SourceAction.Burn, abi.encode(burnedTxFees - 1))
        );

        assertEq(burnedTxFees * DEFAULT_TOKEN_MULTIPLIER, erc20TokenSource.destinationBurnedTotal());
        assertEq(
            burnedTxFees * DEFAULT_TOKEN_MULTIPLIER,
            mockERC20.balanceOf(erc20TokenSource.BURN_ADDRESS())
        );

        emit BurnTokens(additionalTxFees * DEFAULT_TOKEN_MULTIPLIER);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(ITokenSource.SourceAction.Burn, abi.encode(burnedTxFees + additionalTxFees))
        );

        assertEq(
            burnedTxFees * DEFAULT_TOKEN_MULTIPLIER + additionalTxFees * DEFAULT_TOKEN_MULTIPLIER,
            erc20TokenSource.destinationBurnedTotal()
        );
        assertEq(
            burnedTxFees * DEFAULT_TOKEN_MULTIPLIER + additionalTxFees * DEFAULT_TOKEN_MULTIPLIER,
            mockERC20.balanceOf(erc20TokenSource.BURN_ADDRESS())
        );
    }

    function testZeroTeleporterAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");

        erc20TokenSource = new ERC20TokenSource({
            teleporterRegistryAddress: address(0x0),
            destinationBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenDestinationAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            erc20ContractAddress_: address(mockERC20),
            tokenMultiplier_: DEFAULT_TOKEN_MULTIPLIER,
            multiplyOnSend_: true
        });
    }

    function testZeroDestinationBlockhainID() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("zero destination blockchain ID"));

        erc20TokenSource = new ERC20TokenSource({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            destinationBlockchainID_: bytes32(0),
            nativeTokenDestinationAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            erc20ContractAddress_: address(mockERC20),
            tokenMultiplier_: DEFAULT_TOKEN_MULTIPLIER,
            multiplyOnSend_: true
        });
    }

    function testSameBlockchainID() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("cannot bridge with same blockchain"));

        erc20TokenSource = new ERC20TokenSource({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            destinationBlockchainID_: _MOCK_BLOCKCHAIN_ID,
            nativeTokenDestinationAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            erc20ContractAddress_: address(mockERC20),
            tokenMultiplier_: DEFAULT_TOKEN_MULTIPLIER,
            multiplyOnSend_: true
        });
    }

    function testZeroDestinationContractAddress() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("zero destination contract address"));

        erc20TokenSource = new ERC20TokenSource({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            destinationBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenDestinationAddress_: address(0x0),
            erc20ContractAddress_: address(mockERC20),
            tokenMultiplier_: DEFAULT_TOKEN_MULTIPLIER,
            multiplyOnSend_: true
        });
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

        erc20TokenSource = new ERC20TokenSource({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            destinationBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenDestinationAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            erc20ContractAddress_: address(0x0),
            tokenMultiplier_: DEFAULT_TOKEN_MULTIPLIER,
            multiplyOnSend_: true
        });
    }

    function testZeroTokenMultiplier() public {
        vm.expectRevert(_formatERC20TokenSourceErrorMessage("zero tokenMultiplier"));

        erc20TokenSource = new ERC20TokenSource({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            destinationBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenDestinationAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            erc20ContractAddress_: address(mockERC20),
            tokenMultiplier_: 0,
            multiplyOnSend_: true
        });
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
