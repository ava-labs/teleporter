// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {UnitTestMockERC20} from "@mocks/UnitTestMockERC20.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ITeleporterMessenger} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";

contract NativeTokenBridgeTest is Test {
    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS =
        0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
    address public constant MOCK_TELEPORTER_REGISTRY_ADDRESS =
        0xf9FA4a0c696b659328DDaaBCB46Ae4eBFC9e68e4;
    address public constant WARP_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000005);
    address public constant NATIVE_MINTER_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000001);
    bytes32 internal constant _MOCK_BLOCKCHAIN_ID = bytes32(uint256(123456));
    bytes32 internal constant _DEFAULT_OTHER_CHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    address internal constant _DEFAULT_OTHER_BRIDGE_ADDRESS =
        0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    uint256 internal constant _DEFAULT_INITIAL_RESERVE_IMBALANCE = 1000000000;
    address internal constant _DEFAULT_RECIPIENT = 0xa4CEE7d1aF6aDdDD33E3b1cC680AB84fdf1b6d1d;
    uint256 internal constant _DEFAULT_TRANSFER_AMOUNT = 1e18;
    uint256 internal constant _DEFAULT_FEE_AMOUNT = 123456;
    bytes32 internal constant _MOCK_MESSAGE_ID =
        bytes32(hex"1111111111111111111111111111111111111111111111111111111111111111");

    UnitTestMockERC20 public mockERC20;

    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(_MOCK_BLOCKCHAIN_ID)
        );
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeWithSelector(ITeleporterMessenger.sendCrossChainMessage.selector),
            abi.encode(_MOCK_MESSAGE_ID)
        );

        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector)
        );

        _initMockTeleporterRegistry();

        mockERC20 = new UnitTestMockERC20();
        vm.mockCall(
            address(mockERC20), abi.encodeWithSelector(IERC20.allowance.selector), abi.encode(1234)
        );
        vm.mockCall(
            address(mockERC20), abi.encodeWithSelector(IERC20.approve.selector), abi.encode(true)
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
            abi.encodeWithSelector(TeleporterRegistry.getVersionFromAddress.selector),
            abi.encode(0)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getLatestTeleporter.selector),
            abi.encode(ITeleporterMessenger(MOCK_TELEPORTER_MESSENGER_ADDRESS))
        );
    }
}
