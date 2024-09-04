// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "../../teleporter/ITeleporterMessenger.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {TeleporterRegistry} from "@teleporter/registry/TeleporterRegistry.sol";
import {SequentialDeliverAppExample, SequentialMessage} from "../SequentialDeliverAppExample.sol";

contract SequentialDeliverAppExampleTest is Test {
    SequentialDeliverAppExample public app;

    address public constant MOCK_TELEPORTER_REGISTRY_ADDRESS = address(0x1);
    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS = address(0x2);
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
    bytes32 public constant DEFAULT_CURRENT_BLOCKCHAIN_ID =
        bytes32(hex"1111111111111111111111111111111111111111111111111111111111111111");
    uint256 public constant DEFAULT_MIN_TELEPORTER_VERSION = 1;
    bytes32 public constant DEFAULT_PARTNER_APP_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    address public constant DEFAULT_PARTNER_APP_ADDRESS = address(0x4);
    bytes32 internal constant _MOCK_MESSAGE_ID =
        bytes32(hex"2222222222222222222222222222222222222222222222222222222222222222");

    event SequentialMessageSent(uint256 nonce, bytes payload);
    event SequentialMessageReceived(uint256 nonce, bytes payload);

    function setUp() public virtual {
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeWithSelector(ITeleporterMessenger.sendCrossChainMessage.selector),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(DEFAULT_CURRENT_BLOCKCHAIN_ID)
        );

        _initMockTeleporterRegistry();
        app = new SequentialDeliverAppExample(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            DEFAULT_MIN_TELEPORTER_VERSION,
            DEFAULT_PARTNER_APP_BLOCKCHAIN_ID,
            DEFAULT_PARTNER_APP_ADDRESS
        );
    }

    function testSendSequentialMessage() public {
        bytes memory payload = "0x1234";
        vm.expectEmit(true, true, true, true, address(app));
        emit SequentialMessageSent(1, payload);
        app.sendMessage(payload);
        vm.expectEmit(true, true, true, true, address(app));
        emit SequentialMessageSent(2, payload);
        app.sendMessage(payload);
        vm.expectEmit(true, true, true, true, address(app));
        emit SequentialMessageSent(3, payload);
        app.sendMessage(payload);
    }

    function testReceiveSequentialMessage() public {
        // Correct nonce.
        SequentialMessage memory message = SequentialMessage(1, "0x1234");
        vm.expectEmit(true, true, true, true, address(app));
        emit SequentialMessageReceived(message.nonce, message.payload);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_PARTNER_APP_BLOCKCHAIN_ID, DEFAULT_PARTNER_APP_ADDRESS, abi.encode(message)
        );

        // Incorrect nonce.
        message = SequentialMessage(3, "0x1234");
        vm.expectRevert("SequentialDeliverAppExample: Invalid nonce");
        app.receiveTeleporterMessage(
            DEFAULT_PARTNER_APP_BLOCKCHAIN_ID, DEFAULT_PARTNER_APP_ADDRESS, abi.encode(message)
        );

        // Correct nonce.
        message = SequentialMessage(2, "0x1234");
        vm.expectEmit(true, true, true, true, address(app));
        emit SequentialMessageReceived(message.nonce, message.payload);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_PARTNER_APP_BLOCKCHAIN_ID, DEFAULT_PARTNER_APP_ADDRESS, abi.encode(message)
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
            abi.encodeWithSelector(TeleporterRegistry.getVersionFromAddress.selector),
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

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getVersionFromAddress.selector),
            abi.encode(1)
        );
    }
}
