// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.23;

import {
    BaseTeleporterRegistryOwnableAppTest,
    ExampleRegistryOwnableApp,
    ExampleRegistryOwnableAppUpgradeable
} from "./BaseTeleporterRegistryOwnableAppTest.t.sol";
import {ExampleRegistryApp} from "./BaseTeleporterRegistryAppTests.t.sol";
import {Ownable} from "@openzeppelin/contracts@5.0.2/access/Ownable.sol";

contract TeleporterRegistryOwnableAppTest is BaseTeleporterRegistryOwnableAppTest {
    function setUp() public virtual override {
        BaseTeleporterRegistryOwnableAppTest.setUp();
        ownerApp = new ExampleRegistryOwnableApp(address(teleporterRegistry), DEFAULT_OWNER_ADDRESS);
        app = ExampleRegistryApp(address(ownerApp));
    }

    function testZeroInitialOwner() public virtual {
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableInvalidOwner.selector, address(0)));
        ownerApp = new ExampleRegistryOwnableApp(address(teleporterRegistry), address(0));
    }
}

contract TeleporterRegistryOwnableAppUpgradeableTest is BaseTeleporterRegistryOwnableAppTest {
    function setUp() public virtual override {
        BaseTeleporterRegistryOwnableAppTest.setUp();
        ExampleRegistryOwnableAppUpgradeable upgradeableApp =
            new ExampleRegistryOwnableAppUpgradeable();
        upgradeableApp.initialize(address(teleporterRegistry), DEFAULT_OWNER_ADDRESS);
        ownerApp = ExampleRegistryOwnableApp(address(upgradeableApp));
        app = ExampleRegistryApp(address(ownerApp));
    }

    function testZeroInitialOwner() public {
        ExampleRegistryOwnableAppUpgradeable upgradeableApp =
            new ExampleRegistryOwnableAppUpgradeable();
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableInvalidOwner.selector, address(0)));
        upgradeableApp.initialize(address(teleporterRegistry), address(0));
    }
}
