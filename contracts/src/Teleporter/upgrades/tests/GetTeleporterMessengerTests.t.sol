// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterUpgradeableTest, ExampleUpgradeableApp} from "./TeleporterUpgradeableTests.t.sol";
import {ITeleporterMessenger} from "../../ITeleporterMessenger.sol";
import {TeleporterMessenger} from "../../TeleporterMessenger.sol";

contract GetTeleporterMessengerTest is TeleporterUpgradeableTest {
    function setUp() public virtual override {
        TeleporterUpgradeableTest.setUp();
    }

    function testGetTeleporterMessengerBasic() public {
        ITeleporterMessenger messenger = app.getTeleporterMessenger();
        assertEq(
            teleporterRegistry.getVersionFromAddress(address(messenger)),
            teleporterRegistry.latestVersion()
        );
    }

    function testGetPausedTeleporterMessenger() public {
        app.pauseTeleporterAddress(teleporterAddress);
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage(
                "Teleporter sending paused"
            )
        );
        app.getTeleporterMessenger();
    }

    function testGetNewTeleporterMessenger() public {
        // Pause the current latest version of Teleporter
        app.pauseTeleporterAddress(teleporterAddress);
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage(
                "Teleporter sending paused"
            )
        );
        app.getTeleporterMessenger();

        // Add a new version of Teleporter, and make sure we can get
        // the new Teleporter successfully.
        address newTeleporterAddress = address(new TeleporterMessenger());
        _addProtocolVersion(teleporterRegistry, newTeleporterAddress);
        ITeleporterMessenger messenger = app.getTeleporterMessenger();
        assertEq(address(messenger), newTeleporterAddress);
        assertEq(
            teleporterRegistry.getVersionFromAddress(address(messenger)),
            teleporterRegistry.latestVersion()
        );
    }

    function testPauseNonLatestTeleporter() public {
        address newTeleporterAddress = address(new TeleporterMessenger());
        _addProtocolVersion(teleporterRegistry, newTeleporterAddress);

        // Pause a non-latest version of Teleporter
        app.pauseTeleporterAddress(teleporterAddress);

        // Make sure we can still get the latest version of Teleporter for sending
        ITeleporterMessenger messenger = app.getTeleporterMessenger();
        assertEq(address(messenger), newTeleporterAddress);
        assertEq(
            teleporterRegistry.getVersionFromAddress(address(messenger)),
            teleporterRegistry.latestVersion()
        );
    }
}
