// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {NativeTokenBridgeTest} from "./NativeTokenBridgeTest.t.sol";
import {ITokenSource} from "../ITokenSource.sol";
import {
    NativeTokenDestination,
    TeleporterMessageInput,
    TeleporterFeeInfo,
    ITeleporterMessenger
} from "../NativeTokenDestination.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";

contract NativeTokenDestinationTest is NativeTokenBridgeTest {
    NativeTokenDestination public nativeTokenDestination;

    event TransferToSource(
        address indexed sender,
        address indexed recipient,
        bytes32 indexed teleporterMessageID,
        uint256 amount
    );
    event CollateralAdded(uint256 amount, uint256 remaining);
    event NativeTokensMinted(address indexed recipient, uint256 amount);
    event ReportTotalBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 burnAddressBalance);

    function setUp() public virtual override {
        NativeTokenBridgeTest.setUp();
        vm.mockCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(INativeMinter.mintNativeCoin.selector),
            ""
        );
        nativeTokenDestination = new NativeTokenDestination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            _DEFAULT_INITIAL_RESERVE_IMBALANCE
        );
    }

    function collateralizeBridge() public {
        vm.expectEmit(true, true, true, true, address(nativeTokenDestination));
        emit CollateralAdded({amount: _DEFAULT_INITIAL_RESERVE_IMBALANCE, remaining: 0});

        // We shouldn't mint anything here.
        vm.expectCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(INativeMinter.mintNativeCoin.selector),
            0
        );

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenDestination.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_INITIAL_RESERVE_IMBALANCE)
        );
    }

    function testTransferToSource() public {
        collateralizeBridge();

        vm.expectEmit(true, true, true, true, address(nativeTokenDestination));
        emit TransferToSource({
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
            requiredGasLimit: nativeTokenDestination.TRANSFER_NATIVE_TOKENS_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: abi.encode(
                ITokenSource.SourceAction.Unlock,
                abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
                )
        });

        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        nativeTokenDestination.transferToSource{value: _DEFAULT_TRANSFER_AMOUNT}(
            _DEFAULT_RECIPIENT,
            TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: _DEFAULT_FEE_AMOUNT}),
            new address[](0)
        );
    }

    function testCollateralizeBridge() public {
        uint256 firstTransfer = _DEFAULT_INITIAL_RESERVE_IMBALANCE / 4;

        assertEq(_DEFAULT_INITIAL_RESERVE_IMBALANCE, nativeTokenDestination.totalSupply());

        vm.expectEmit(true, true, true, true, address(nativeTokenDestination));
        emit CollateralAdded({
            amount: firstTransfer,
            remaining: _DEFAULT_INITIAL_RESERVE_IMBALANCE - firstTransfer
        });

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenDestination.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(_DEFAULT_RECIPIENT, firstTransfer)
        );

        assertEq(
            _DEFAULT_INITIAL_RESERVE_IMBALANCE - firstTransfer,
            nativeTokenDestination.currentReserveImbalance()
        );
        assertEq(_DEFAULT_INITIAL_RESERVE_IMBALANCE, nativeTokenDestination.totalSupply());

        vm.expectEmit(true, true, true, true, address(nativeTokenDestination));
        emit CollateralAdded({
            amount: _DEFAULT_INITIAL_RESERVE_IMBALANCE - firstTransfer,
            remaining: 0
        });
        emit NativeTokensMinted(_DEFAULT_RECIPIENT, firstTransfer);

        vm.expectCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(INativeMinter.mintNativeCoin.selector),
            1
        );

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenDestination.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_INITIAL_RESERVE_IMBALANCE)
        );

        assertEq(0, nativeTokenDestination.currentReserveImbalance());
        assertEq(
            _DEFAULT_INITIAL_RESERVE_IMBALANCE + firstTransfer, nativeTokenDestination.totalSupply()
        );
    }

    function testReportBurnedTxFees() public {
        uint256 burnedFees = nativeTokenDestination.BURNED_TX_FEES_ADDRESS().balance;

        vm.expectEmit(true, true, true, true, address(nativeTokenDestination));
        emit ReportTotalBurnedTxFees({
            teleporterMessageID: _MOCK_MESSAGE_ID,
            burnAddressBalance: burnedFees
        });

        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            feeInfo: TeleporterFeeInfo({
                feeTokenAddress: address(mockERC20),
                amount: _DEFAULT_FEE_AMOUNT
            }),
            requiredGasLimit: nativeTokenDestination.REPORT_BURNED_TOKENS_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: abi.encode(ITokenSource.SourceAction.Burn, abi.encode(burnedFees))
        });

        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        nativeTokenDestination.reportTotalBurnedTxFees(
            TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: _DEFAULT_FEE_AMOUNT}),
            new address[](0)
        );
    }

    function testZeroTeleporterAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");

        new NativeTokenDestination(
            address(0x0),
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            _DEFAULT_INITIAL_RESERVE_IMBALANCE
        );
    }

    function testZeroSourceBlockchainID() public {
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("zero source blockchain ID"));

        new NativeTokenDestination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            bytes32(0),
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            _DEFAULT_INITIAL_RESERVE_IMBALANCE
        );
    }

    function testSameBlockchainID() public {
        vm.expectRevert(
            _formatNativeTokenDestinationErrorMessage("cannot bridge with same blockchain")
        );

        new NativeTokenDestination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _MOCK_BLOCKCHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            _DEFAULT_INITIAL_RESERVE_IMBALANCE
        );
    }

    function testZeroSourceContractAddress() public {
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("zero source contract address"));

        new NativeTokenDestination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            address(0x0),
            _DEFAULT_INITIAL_RESERVE_IMBALANCE
        );
    }

    function testZeroInitialReserveImbalance() public {
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("zero initial reserve imbalance"));

        new NativeTokenDestination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            0
        );
    }

    function testInvalidTeleporterAddress() public {
        vm.expectRevert("TeleporterUpgradeable: invalid Teleporter sender");

        vm.prank(address(0x123));
        nativeTokenDestination.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
        );
    }

    function testInvalidSourceBlockchain() public {
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("invalid source chain"));

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenDestination.receiveTeleporterMessage(
            _MOCK_BLOCKCHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
        );
    }

    function testInvalidSenderContract() public {
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("unauthorized sender"));

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenDestination.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            address(0x123),
            abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
        );
    }

    function testInvalidRecipientAddress() public {
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("zero recipient address"));

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenDestination.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(address(0x0), _DEFAULT_TRANSFER_AMOUNT)
        );
    }

    function testInvalidTransferAmount() public {
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("zero transfer value"));

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        nativeTokenDestination.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(_DEFAULT_RECIPIENT, 0)
        );
    }

    function testZeroRecipient() public {
        collateralizeBridge();
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("zero recipient address"));

        nativeTokenDestination.transferToSource{value: _DEFAULT_TRANSFER_AMOUNT}(
            address(0x0),
            TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: _DEFAULT_FEE_AMOUNT}),
            new address[](0)
        );
    }

    function testUncollateralizedBridge() public {
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("contract undercollateralized"));

        nativeTokenDestination.transferToSource{value: _DEFAULT_TRANSFER_AMOUNT}(
            _DEFAULT_RECIPIENT,
            TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: _DEFAULT_FEE_AMOUNT}),
            new address[](0)
        );
    }

    function _formatNativeTokenDestinationErrorMessage(string memory errorMessage)
        private
        pure
        returns (bytes memory)
    {
        return bytes(string.concat("NativeTokenDestination: ", errorMessage));
    }
}
