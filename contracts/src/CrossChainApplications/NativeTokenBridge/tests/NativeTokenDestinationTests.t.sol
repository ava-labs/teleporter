// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {NativeTokenDestination, IERC20, ITokenSource, SafeERC20, SafeERC20TransferFrom, TeleporterMessageInput, TeleporterFeeInfo, IWarpMessenger, ITeleporterMessenger} from "../NativeTokenDestination.sol";
import {INativeTokenDestination} from "../INativeTokenDestination.sol";
import {ITeleporterReceiver} from "../../../Teleporter/ITeleporterReceiver.sol";
import {UnitTestMockERC20} from "../../../Mocks/UnitTestMockERC20.sol";
import {INativeMinter} from "@subnet-evm-contracts/interfaces/INativeMinter.sol";

contract NativeTokenDestinationTest is Test {
    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS =
        0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
    address public constant WARP_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000005);
    address public constant NATIVE_MINTER_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000001);
    bytes32 private constant _MOCK_BLOCKCHAIN_ID = bytes32(uint256(123456));
    bytes32 private constant _DEFAULT_OTHER_CHAIN_ID =
        bytes32(
            hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd"
        );
    address private constant _DEFAULT_OTHER_BRIDGE_ADDRESS =
        0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    uint256 private constant _DEFAULT_INITIAL_RESERVE_IMBALANCE = 1000000000;
    address private constant _DEFAULT_RECIPIENT =
        0xa4CEE7d1aF6aDdDD33E3b1cC680AB84fdf1b6d1d;

    event TransferToSource(
        address indexed sender,
        address indexed recipient,
        uint256 indexed teleporterMessageID,
        uint256 amount
    );
    event CollateralAdded(uint256 amount, uint256 remaining, address recipient);
    event NativeTokensMinted(address indexed recipient, uint256 amount);
    event ReportTotalBurnedTxFees(
        uint256 indexed teleporterMessageID,
        uint256 burnAddressBalance
    );

    NativeTokenDestination public nativeTokenDestination;
    UnitTestMockERC20 public mockERC20;

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
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector)
        );

        nativeTokenDestination = new NativeTokenDestination(MOCK_TELEPORTER_MESSENGER_ADDRESS, _DEFAULT_OTHER_CHAIN_ID, _DEFAULT_OTHER_BRIDGE_ADDRESS, _DEFAULT_INITIAL_RESERVE_IMBALANCE);
        mockERC20 = new UnitTestMockERC20();

        vm.mockCall(
            address(mockERC20),
            abi.encodeWithSelector(IERC20.allowance.selector),
            abi.encode(1234)
        );
        vm.mockCall(
            address(mockERC20),
            abi.encodeWithSelector(IERC20.transfer.selector),
            abi.encode(true)
        );
        vm.mockCall(
            address(mockERC20),
            abi.encodeWithSelector(IERC20.approve.selector),
            abi.encode(true)
        );
    }

    function collateralizeBridge() public {
        uint256 amount = _DEFAULT_INITIAL_RESERVE_IMBALANCE;
                
        vm.expectEmit(true, true, true, true, address(nativeTokenDestination));
        emit CollateralAdded({
            amount: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            remaining: 0,
            recipient: _DEFAULT_RECIPIENT
        });

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        bytes memory message = abi.encode(_DEFAULT_RECIPIENT, amount);
        nativeTokenDestination.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            message
        );
        vm.stopPrank();
    }

    function testTransferToSource() public {
        collateralizeBridge();
        uint256 feeAmount = 12345;
        uint256 amount = 1e18;

        vm.expectEmit(true, true, true, true, address(nativeTokenDestination));
        emit TransferToSource({
            sender: address(this),
            recipient: _DEFAULT_RECIPIENT,
            amount: amount,
            teleporterMessageID: 1
        });

        TeleporterMessageInput
            memory expectedMessageInput = TeleporterMessageInput({
                destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
                destinationAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: address(mockERC20),
                    amount: feeAmount
                }),
                requiredGasLimit: nativeTokenDestination.TRANSFER_NATIVE_TOKENS_REQUIRED_GAS(),
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(
                    ITokenSource.SourceAction.Unlock,
                    abi.encode(_DEFAULT_RECIPIENT, amount)
                )
            });

        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(
                ITeleporterMessenger.sendCrossChainMessage,
                (expectedMessageInput)
            )
        );
        
        nativeTokenDestination.transferToSource{value: amount}(
            _DEFAULT_RECIPIENT, 
            TeleporterFeeInfo({
                feeTokenAddress: address(mockERC20),
                amount: feeAmount
            }), 
            new address[](0)
        );
    }

    function _formatNativeTokenDestinationErrorMessage(
        string memory errorMessage
    ) private pure returns (bytes memory) {
        return bytes(string.concat("NativeTokenDestination: ", errorMessage));
    }
}
