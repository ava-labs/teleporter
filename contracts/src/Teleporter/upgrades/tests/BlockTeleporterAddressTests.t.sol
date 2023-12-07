// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterUpgradeable} from "../TeleporterUpgradeable.sol";
import {TeleporterUpgradeableTest, ExampleUpgradeableApp} from "./TeleporterUpgradeableTests.t.sol";

contract BlockTeleporterAddressTest is TeleporterUpgradeableTest {
    function setUp() public virtual override {
        TeleporterUpgradeableTest.setUp();
    }

    function testBlockTeleporterAddressBasic() public {
        ExampleUpgradeableApp app = new ExampleUpgradeableApp(
            address(teleporterRegistry)
        );
        // Check that the teleporterAddress is not blocked initially
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_ORIGIN_ADDRESS,
            ""
        );
        // Check that teleporterAddress can not deliver messages once blocked
        app.blockTeleporterAddress(teleporterAddress);
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage(
                "Teleporter address blocked"
            )
        );
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_ORIGIN_ADDRESS,
            ""
        );
    }

    function testAlreadyBlockedTeleporterAddress() public {
        ExampleUpgradeableApp app = new ExampleUpgradeableApp(
            address(teleporterRegistry)
        );

        // Check that teleporterAddress can not deliver messages once blocked
        app.blockTeleporterAddress(teleporterAddress);
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage(
                "Teleporter address blocked"
            )
        );
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_ORIGIN_ADDRESS,
            ""
        );
        // Check that teleporterAddress can not be blocked again
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage("address already blocked")
        );
        app.blockTeleporterAddress(teleporterAddress);
    }

    function testBlockMultipleVersions() public {
        // Since a Teleporter address can be registered with multiple versions,
        // when blocking a Teleporter address, we should block all versions of that address.

        ExampleUpgradeableApp app = new ExampleUpgradeableApp(
            address(teleporterRegistry)
        );
        // Add the same Teleporter address to the registry with a different version
        _addProtocolVersion(teleporterRegistry, teleporterAddress);

        // Check that the teleporterAddress is not blocked initially
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_ORIGIN_ADDRESS,
            ""
        );

        // Check that teleporterAddress can not deliver messages once blocked
        app.blockTeleporterAddress(teleporterAddress);
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage(
                "Teleporter address blocked"
            )
        );
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_ORIGIN_ADDRESS,
            ""
        );

        // Check that after updating mininum Teleporter version, the address is still blocked
        app.updateMinTeleporterVersion(teleporterRegistry.latestVersion());
        vm.prank(teleporterAddress);
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage(
                "Teleporter address blocked"
            )
        );
        app.receiveTeleporterMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_ORIGIN_ADDRESS,
            ""
        );
    }

    function testBlockNonRegisteredTeleporterAddress() public {
        ExampleUpgradeableApp app = new ExampleUpgradeableApp(
            address(teleporterRegistry)
        );

        vm.expectRevert(_formatTele("version not found"));
        app.blockTeleporterAddress(address(this));
    }
}
