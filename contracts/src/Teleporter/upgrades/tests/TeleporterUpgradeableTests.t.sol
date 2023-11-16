// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterUpgradeable} from "../TeleporterUpgradeable.sol";
import {TeleporterRegistryTest, TeleporterMessenger} from "./TeleporterRegistryTests.t.sol";

contract ExampleUpgradeableApp is TeleporterUpgradeable {
    constructor(
        address teleporterRegistryAddress
    ) TeleporterUpgradeable(teleporterRegistryAddress) {}

    function updateMinTeleporterVersion() external override {
        uint256 oldMinTeleporterVersion = minTeleporterVersion;
        minTeleporterVersion = teleporterRegistry.latestVersion();
        emit MinTeleporterVersionUpdated(
            oldMinTeleporterVersion,
            minTeleporterVersion
        );
    }

    // solhint-disable-next-line no-empty-blocks
    function teleporterCall() public onlyAllowedTeleporter {}
}

contract TeleporterUpgradeableTest is TeleporterRegistryTest {
    event MinTeleporterVersionUpdated(
        uint256 indexed oldMinTeleporterVersion,
        uint256 indexed newMinTeleporterVersion
    );

    function setUp() public virtual override {
        TeleporterRegistryTest.setUp();
        _addProtocolVersion(teleporterRegistry, teleporterAddress);
    }

    function testInvalidRegistryAddress() public {
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage(
                "zero teleporter registry address"
            )
        );
        new ExampleUpgradeableApp(address(0));
    }

    function testOnlyAllowedTeleporter() public {
        ExampleUpgradeableApp app = new ExampleUpgradeableApp(
            address(teleporterRegistry)
        );

        assertEq(app.minTeleporterVersion(), 1);

        vm.expectRevert(
            _formatRegistryErrorMessage("protocol address not found")
        );
        app.teleporterCall();

        vm.prank(teleporterAddress);
        app.teleporterCall();
    }

    function testUpdateMinTeleporterVersion() public {
        ExampleUpgradeableApp app = new ExampleUpgradeableApp(
            address(teleporterRegistry)
        );

        // First check that calling with initial teleporter address works
        assertEq(app.minTeleporterVersion(), 1);
        vm.prank(teleporterAddress);
        app.teleporterCall();

        // Now add new protocol version to registry and update the app's min version
        address newTeleporterAddress = address(new TeleporterMessenger());
        _addProtocolVersion(teleporterRegistry, newTeleporterAddress);

        vm.expectEmit(true, true, true, true, address(app));
        emit MinTeleporterVersionUpdated(1, 2);

        app.updateMinTeleporterVersion();
        assertEq(app.minTeleporterVersion(), 2);

        // Check that calling with the old teleporter address fails
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage(
                "invalid teleporter sender"
            )
        );
        vm.prank(teleporterAddress);
        app.teleporterCall();

        // Check that calling with the new teleporter address works
        vm.prank(newTeleporterAddress);
        app.teleporterCall();
    }

    function _formatTeleporterUpgradeableErrorMessage(
        string memory errorMessage
    ) internal pure returns (bytes memory) {
        return bytes(string.concat("TeleporterUpgradeable: ", errorMessage));
    }
}
