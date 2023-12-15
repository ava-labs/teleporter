// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {
    NativeTokenDestination,
    IERC20,
    ITokenSource,
    TeleporterMessageInput,
    TeleporterFeeInfo,
    IWarpMessenger,
    ITeleporterMessenger
} from "../NativeTokenDestination.sol";
import {TeleporterRegistry} from "../../../Teleporter/upgrades/TeleporterRegistry.sol";
import {UnitTestMockERC20} from "../../../Mocks/UnitTestMockERC20.sol";
import {INativeMinter} from "@subnet-evm-contracts/interfaces/INativeMinter.sol";

contract NativeTokenDestinationTest is Test {
    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS =
        0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
    address public constant MOCK_TELEPORTER_REGISTRY_ADDRESS =
        0xf9FA4a0c696b659328DDaaBCB46Ae4eBFC9e68e4;
    address public constant WARP_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000005);
    address public constant NATIVE_MINTER_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000001);
    bytes32 private constant _MOCK_BLOCKCHAIN_ID = bytes32(uint256(123456));
    bytes32 private constant _DEFAULT_OTHER_CHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    address private constant _DEFAULT_OTHER_BRIDGE_ADDRESS =
        0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    uint256 private constant _DEFAULT_INITIAL_RESERVE_IMBALANCE = 1000000000;
    address private constant _DEFAULT_RECIPIENT = 0xa4CEE7d1aF6aDdDD33E3b1cC680AB84fdf1b6d1d;
    uint256 private constant _DEFAULT_TRANSFER_AMOUNT = 1e18;
    uint256 private constant _DEFAULT_FEE_AMOUNT = 123456;

    NativeTokenDestination public nativeTokenDestination;
    UnitTestMockERC20 public mockERC20;

    event TransferToSource(
        address indexed sender,
        address indexed recipient,
        uint256 indexed teleporterMessageID,
        uint256 amount
    );
    event CollateralAdded(uint256 amount, uint256 remaining);
    event NativeTokensMinted(address indexed recipient, uint256 amount);
    event ReportTotalBurnedTxFees(uint256 indexed teleporterMessageID, uint256 burnAddressBalance);

    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(_MOCK_BLOCKCHAIN_ID)
        );
        vm.mockCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(INativeMinter.mintNativeCoin.selector),
            ""
        );
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeWithSelector(ITeleporterMessenger.sendCrossChainMessage.selector),
            abi.encode(1)
        );

        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector)
        );

        _initMockTeleporterRegistry();

        nativeTokenDestination = new NativeTokenDestination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            _DEFAULT_INITIAL_RESERVE_IMBALANCE
        );
        mockERC20 = new UnitTestMockERC20();

        vm.mockCall(
            address(mockERC20), abi.encodeWithSelector(IERC20.allowance.selector), abi.encode(1234)
        );
        vm.mockCall(
            address(mockERC20), abi.encodeWithSelector(IERC20.approve.selector), abi.encode(true)
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
            teleporterMessageID: 1
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
        emit ReportTotalBurnedTxFees({burnAddressBalance: burnedFees, teleporterMessageID: 1});

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

    function testZeroSourceChainID() public {
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
        vm.expectRevert();

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
                TeleporterRegistry.getVersionFromAddress.selector,
                (MOCK_TELEPORTER_MESSENGER_ADDRESS)
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
    }

    function _formatNativeTokenDestinationErrorMessage(string memory errorMessage)
        private
        pure
        returns (bytes memory)
    {
        return bytes(string.concat("NativeTokenDestination: ", errorMessage));
    }
}
