// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "forge-std/Test.sol";
import "../TeleporterRegistry.sol";
import "../TeleporterUpgradeable.sol";

contract ExampleUpgradeableApp is TeleporterUpgradeable {
    constructor(
        address teleporterRegistryAddress
    ) TeleporterUpgradeable(teleporterRegistryAddress) {}

    // solhint-disable-next-line no-empty-blocks
    function teleporterCall() public onlyAllowedTeleporter {}
}

contract TeleporterUpgradeableTest is Test {
    address public constant MOCK_TELEPORTER_REGISTRY_ADDRESS =
        0xf9FA4a0c696b659328DDaaBCB46Ae4eBFC9e68e4;
    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS =
        0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;

    function setUp() public {
        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeCall(WarpProtocolRegistry.getLatestVersion, ()),
            abi.encode(1)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeCall(
                WarpProtocolRegistry.getVersionFromAddress,
                (MOCK_TELEPORTER_MESSENGER_ADDRESS)
            ),
            abi.encode(1)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeCall(
                WarpProtocolRegistry.getVersionFromAddress,
                (address(this))
            ),
            abi.encode(0)
        );
    }

    function testInvalidRegistryAddress() public {
        vm.expectRevert(
            TeleporterUpgradeable.InvalidTeleporterRegistryAddress.selector
        );
        new ExampleUpgradeableApp(address(0));
    }

    function testOnlyAllowedTeleporter() public {
        ExampleUpgradeableApp app = new ExampleUpgradeableApp(
            MOCK_TELEPORTER_REGISTRY_ADDRESS
        );

        assertEq(app.getMinTeleporterVersion(), 1);

        vm.expectRevert(TeleporterUpgradeable.InvalidTeleporterSender.selector);
        app.teleporterCall();

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.teleporterCall();
    }

    function testUpdateMinTeleporterVersion() public {
        ExampleUpgradeableApp app = new ExampleUpgradeableApp(
            MOCK_TELEPORTER_REGISTRY_ADDRESS
        );

        // First check that calling with initial teleporter address works
        assertEq(app.getMinTeleporterVersion(), 1);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.teleporterCall();

        // Now mock the registry getting a new version and updating the app's min version
        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeCall(WarpProtocolRegistry.getLatestVersion, ()),
            abi.encode(2)
        );

        app.updateMinTeleporterVersion();
        assertEq(app.getMinTeleporterVersion(), 2);

        // Check that calling with the old teleporter address fails
        vm.expectRevert(TeleporterUpgradeable.InvalidTeleporterSender.selector);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.teleporterCall();
    }
}
