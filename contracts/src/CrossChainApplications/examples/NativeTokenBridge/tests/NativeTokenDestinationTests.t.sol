// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {NativeTokenBridgeTest} from "./NativeTokenBridgeTest.t.sol";
import {ITokenSource} from "../ITokenSource.sol";
import {NativeTokenDestination} from "../NativeTokenDestination.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";

contract NativeTokenDestinationTest is NativeTokenBridgeTest {
    NativeTokenDestination public nativeTokenDestination;
    uint256 internal constant _DEFAULT_DECIMALS_SHIFT = 1;
    uint256 internal constant _DEFAULT_TOKEN_MULTIPLIER = 10 ** _DEFAULT_DECIMALS_SHIFT;

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
        nativeTokenDestination = new NativeTokenDestination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: _DEFAULT_OWNER_ADDRESS,
            sourceBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenSourceAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            initialReserveImbalance_: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: _DEFAULT_DECIMALS_SHIFT,
            multiplyOnReceive_: true
        });
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
            abi.encode(
                _DEFAULT_RECIPIENT, _DEFAULT_INITIAL_RESERVE_IMBALANCE / _DEFAULT_TOKEN_MULTIPLIER
            )
        );
    }

    function testTransferToSource() public {
        collateralizeBridge();

        vm.expectEmit(true, true, true, true, address(nativeTokenDestination));
        emit TransferToSource({
            sender: address(this),
            recipient: _DEFAULT_RECIPIENT,
            amount: _DEFAULT_TRANSFER_AMOUNT / _DEFAULT_TOKEN_MULTIPLIER,
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
                abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT / _DEFAULT_TOKEN_MULTIPLIER)
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

    function testTransferToSourceDivideOnSend() public {
        nativeTokenDestination = new NativeTokenDestination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: _DEFAULT_OWNER_ADDRESS,
            sourceBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenSourceAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            initialReserveImbalance_: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: _DEFAULT_DECIMALS_SHIFT,
            multiplyOnReceive_: false
        });

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
            abi.encode(
                _DEFAULT_RECIPIENT, _DEFAULT_INITIAL_RESERVE_IMBALANCE * _DEFAULT_TOKEN_MULTIPLIER
            )
        );

        vm.expectEmit(true, true, true, true, address(nativeTokenDestination));
        emit TransferToSource({
            sender: address(this),
            recipient: _DEFAULT_RECIPIENT,
            amount: _DEFAULT_TRANSFER_AMOUNT * _DEFAULT_TOKEN_MULTIPLIER,
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
                abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT * _DEFAULT_TOKEN_MULTIPLIER)
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
            abi.encode(_DEFAULT_RECIPIENT, firstTransfer / _DEFAULT_TOKEN_MULTIPLIER)
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
            abi.encode(
                _DEFAULT_RECIPIENT, _DEFAULT_INITIAL_RESERVE_IMBALANCE / _DEFAULT_TOKEN_MULTIPLIER
            )
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
            burnAddressBalance: burnedFees / _DEFAULT_TOKEN_MULTIPLIER
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
            message: abi.encode(
                ITokenSource.SourceAction.Burn, abi.encode(burnedFees / _DEFAULT_TOKEN_MULTIPLIER)
                )
        });

        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        vm.expectCall(
            address(mockERC20),
            abi.encodeCall(
                IERC20.transferFrom,
                (address(this), address(nativeTokenDestination), _DEFAULT_FEE_AMOUNT)
            )
        );
        vm.expectCall(
            address(mockERC20),
            abi.encodeCall(
                IERC20.allowance,
                (address(nativeTokenDestination), MOCK_TELEPORTER_MESSENGER_ADDRESS)
            )
        );

        nativeTokenDestination.reportTotalBurnedTxFees(
            TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: _DEFAULT_FEE_AMOUNT}),
            new address[](0)
        );
    }

    function testZeroTeleporterAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");

        nativeTokenDestination = new NativeTokenDestination({
            teleporterRegistryAddress: address(0x0),
            teleporterManager: _DEFAULT_OWNER_ADDRESS,
            sourceBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenSourceAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            initialReserveImbalance_: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: _DEFAULT_DECIMALS_SHIFT,
            multiplyOnReceive_: true
        });
    }

    function testZeroSourceBlockchainID() public {
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("zero source blockchain ID"));

        nativeTokenDestination = new NativeTokenDestination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: _DEFAULT_OWNER_ADDRESS,
            sourceBlockchainID_: bytes32(0),
            nativeTokenSourceAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            initialReserveImbalance_: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: _DEFAULT_DECIMALS_SHIFT,
            multiplyOnReceive_: true
        });
    }

    function testSameBlockchainID() public {
        vm.expectRevert(
            _formatNativeTokenDestinationErrorMessage("cannot bridge with same blockchain")
        );
        nativeTokenDestination = new NativeTokenDestination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: _DEFAULT_OWNER_ADDRESS,
            sourceBlockchainID_: _MOCK_BLOCKCHAIN_ID,
            nativeTokenSourceAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            initialReserveImbalance_: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: _DEFAULT_DECIMALS_SHIFT,
            multiplyOnReceive_: true
        });
    }

    function invalidDecimalsShift() public {
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("invalid decimalsShift"));

        nativeTokenDestination = new NativeTokenDestination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: _DEFAULT_OWNER_ADDRESS,
            sourceBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenSourceAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            initialReserveImbalance_: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: 19,
            multiplyOnReceive_: true
        });
    }

    function testZeroSourceContractAddress() public {
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("zero source contract address"));

        nativeTokenDestination = new NativeTokenDestination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: _DEFAULT_OWNER_ADDRESS,
            sourceBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenSourceAddress_: address(0),
            initialReserveImbalance_: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: _DEFAULT_DECIMALS_SHIFT,
            multiplyOnReceive_: true
        });
    }

    function testZeroInitialReserveImbalance() public {
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("zero initial reserve imbalance"));

        nativeTokenDestination = new NativeTokenDestination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: _DEFAULT_OWNER_ADDRESS,
            sourceBlockchainID_: _DEFAULT_OTHER_CHAIN_ID,
            nativeTokenSourceAddress_: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            initialReserveImbalance_: 0,
            decimalsShift: _DEFAULT_DECIMALS_SHIFT,
            multiplyOnReceive_: true
        });
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

    function testTransferZeroAmount() public {
        collateralizeBridge();
        vm.expectRevert(_formatNativeTokenDestinationErrorMessage("zero transfer value"));

        nativeTokenDestination.transferToSource{value: 0}(
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
