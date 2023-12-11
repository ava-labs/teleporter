// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {ERC20TokenSource, IERC20, ITokenSource, TeleporterMessageInput, TeleporterFeeInfo, IWarpMessenger, ITeleporterMessenger} from "../ERC20TokenSource.sol";
import {UnitTestMockERC20} from "../../../Mocks/UnitTestMockERC20.sol";

contract ERC20TokenSourceTest is Test {
    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS =
        0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
    address public constant WARP_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000005);
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
    uint256 private constant _DEFAULT_TRANSFER_AMOUNT = 1e18;
    uint256 private constant _DEFAULT_FEE_AMOUNT = 123456;

    ERC20TokenSource public erc20TokenSource;
    UnitTestMockERC20 public mockERC20;

    event TransferToDestination(
        address indexed sender,
        address indexed recipient,
        uint256 indexed teleporterMessageID,
        uint256 amount
    );
    event UnlockTokens(address recipient, uint256 amount);
    event BurnTokens(uint256 amount);

    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(_MOCK_BLOCKCHAIN_ID)
        );
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeWithSelector(
                ITeleporterMessenger.sendCrossChainMessage.selector
            ),
            abi.encode(1)
        );

        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector)
        );

        mockERC20 = new UnitTestMockERC20();
        erc20TokenSource = new ERC20TokenSource(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(mockERC20)
        );

        vm.mockCall(
            address(mockERC20),
            abi.encodeWithSelector(IERC20.allowance.selector),
            abi.encode(1234)
        );
        vm.mockCall(
            address(mockERC20),
            abi.encodeWithSelector(IERC20.approve.selector),
            abi.encode(true)
        );
    }

    function testTransferToDestination() public {
        vm.expectEmit(true, true, true, true, address(erc20TokenSource));
        emit TransferToDestination({
            sender: address(this),
            recipient: _DEFAULT_RECIPIENT,
            amount: _DEFAULT_TRANSFER_AMOUNT,
            teleporterMessageID: 1
        });

        TeleporterMessageInput
            memory expectedMessageInput = TeleporterMessageInput({
                destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
                destinationAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: address(mockERC20),
                    amount: _DEFAULT_FEE_AMOUNT
                }),
                requiredGasLimit: erc20TokenSource
                    .MINT_NATIVE_TOKENS_REQUIRED_GAS(),
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
            });

        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(
                ITeleporterMessenger.sendCrossChainMessage,
                (expectedMessageInput)
            )
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
            _DEFAULT_TRANSFER_AMOUNT * 2 + _DEFAULT_FEE_AMOUNT ,
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
            abi.encode(
                ITokenSource.SourceAction.Burn,
                abi.encode(burnedTxFees)
            )
        );

        assertEq(burnedTxFees, erc20TokenSource.destinationBurnedTotal());
        assertEq(burnedTxFees, mockERC20.balanceOf(erc20TokenSource.BURNED_TX_FEES_ADDRESS()));

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(
                ITokenSource.SourceAction.Burn,
                abi.encode(burnedTxFees - 1)
            )
        );

        assertEq(burnedTxFees, erc20TokenSource.destinationBurnedTotal());
        assertEq(burnedTxFees, mockERC20.balanceOf(erc20TokenSource.BURNED_TX_FEES_ADDRESS()));

        emit BurnTokens(additionalTxFees);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(
                ITokenSource.SourceAction.Burn,
                abi.encode(burnedTxFees + additionalTxFees)
            )
        );

        assertEq(burnedTxFees + additionalTxFees, erc20TokenSource.destinationBurnedTotal());
        assertEq(burnedTxFees + additionalTxFees, mockERC20.balanceOf(erc20TokenSource.BURNED_TX_FEES_ADDRESS()));
    }

    function testZeroTeleporterAddress() public {
        vm.expectRevert(
            _formatERC20TokenSourceErrorMessage(
                "zero TeleporterMessenger address"
            )
        );

        new ERC20TokenSource(
            address(0x0),
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(mockERC20)
        );
    }

    function testZeroDestinationChainID() public {
        vm.expectRevert(
            _formatERC20TokenSourceErrorMessage(
                "zero destination blockchain ID"
            )
        );

        new ERC20TokenSource(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            bytes32(0),
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(mockERC20)
        );
    }

    function testSameBlockchainID() public {
        vm.expectRevert(
            _formatERC20TokenSourceErrorMessage(
                "cannot bridge with same blockchain"
            )
        );

        new ERC20TokenSource(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            _MOCK_BLOCKCHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(mockERC20)
        );
    }

    function testZeroDestinationContractAddress() public {
        vm.expectRevert(
            _formatERC20TokenSourceErrorMessage(
                "zero destination contract address"
            )
        );

        new ERC20TokenSource(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            address(0x0),
            address(mockERC20)
        );
    }

    function testZeroERC20ContractAddress() public {
        vm.expectRevert(
            _formatERC20TokenSourceErrorMessage(
                "zero ERC20 contract address"
            )
        );

        new ERC20TokenSource(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            address(0x0)
        );
    }

    function testInvalidTeleporterAddress() public {
        vm.expectRevert(
            _formatERC20TokenSourceErrorMessage(
                "unauthorized TeleporterMessenger contract"
            )
        );

        vm.prank(address(0x123));
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(
                ITokenSource.SourceAction.Unlock,
                abi.encode(_DEFAULT_RECIPIENT, _DEFAULT_TRANSFER_AMOUNT)
            )
        );
    }

    function testInvalidDestinationBlockchain() public {
        vm.expectRevert(
            _formatERC20TokenSourceErrorMessage("invalid destination chain")
        );

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
        vm.expectRevert(
            _formatERC20TokenSourceErrorMessage("unauthorized sender")
        );

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
        vm.expectRevert(
            _formatERC20TokenSourceErrorMessage("zero recipient address")
        );

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        erc20TokenSource.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID,
            _DEFAULT_OTHER_BRIDGE_ADDRESS,
            abi.encode(
                ITokenSource.SourceAction.Unlock,
                abi.encode(address(0x0), _DEFAULT_TRANSFER_AMOUNT)
            )
        );
    }

    function testZeroRecipient() public {
        vm.expectRevert(
            _formatERC20TokenSourceErrorMessage("zero recipient address")
        );

        erc20TokenSource.transferToDestination(
            address(0x0),
            _DEFAULT_TRANSFER_AMOUNT + _DEFAULT_FEE_AMOUNT,
            _DEFAULT_FEE_AMOUNT,
            new address[](0)
        );
    }

    function _formatERC20TokenSourceErrorMessage(
        string memory errorMessage
    ) private pure returns (bytes memory) {
        return bytes(string.concat("ERC20TokenSource: ", errorMessage));
    }
}
