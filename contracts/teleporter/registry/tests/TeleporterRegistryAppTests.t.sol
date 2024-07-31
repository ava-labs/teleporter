// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {GetTeleporterMessengerTest} from "./GetTeleporterMessengerTests.t.sol";
import {PauseTeleporterAddressTest} from "./PauseTeleporterAddressTests.t.sol";
import {UpdateMinTeleporterVersionTest} from "./UpdateMinTeleporterVersionTests.t.sol";
import {SendTeleporterMessageTest} from "./SendTeleporterMessageTests.t.sol";
import {UnpauseTeleporterAddressTest} from "./UnpauseTeleporterAddressTests.t.sol";
import {BaseTeleporterRegistryAppTest} from "./BaseTeleporterRegistryAppTests.t.sol";
import {
    ExampleRegistryApp,
    ExampleRegistryAppUpgradeable
} from "./BaseTeleporterRegistryAppTests.t.sol";
import {NonReentrantTest} from "./NonReentrantTests.t.sol";

contract TeleporterRegistryAppTest is
    GetTeleporterMessengerTest,
    PauseTeleporterAddressTest,
    UpdateMinTeleporterVersionTest,
    SendTeleporterMessageTest,
    UnpauseTeleporterAddressTest,
    NonReentrantTest
{
    function setUp() public virtual override (BaseTeleporterRegistryAppTest, NonReentrantTest) {
        BaseTeleporterRegistryAppTest.setUp();
        app = new ExampleRegistryApp(address(teleporterRegistry));
        NonReentrantTest.setUp();
    }

    function testInvalidRegistryAddress() public virtual {
        vm.expectRevert(_formatErrorMessage("zero teleporter registry address"));
        app = new ExampleRegistryApp(address(0));
    }
}

contract TeleporterRegistryAppUpgradeableTest is TeleporterRegistryAppTest {
    function setUp() public virtual override {
        TeleporterRegistryAppTest.setUp();
        ExampleRegistryAppUpgradeable upgradeableApp = new ExampleRegistryAppUpgradeable();
        upgradeableApp.initialize(address(teleporterRegistry));
        app = ExampleRegistryApp(address(upgradeableApp));
    }

    function testInvalidRegistryAddress() public override {
        ExampleRegistryAppUpgradeable upgradeableApp = new ExampleRegistryAppUpgradeable();
        vm.expectRevert(_formatErrorMessage("zero teleporter registry address"));
        upgradeableApp.initialize(address(0));
    }

    function testStorageSlot() public {
        assertEq(
            _erc7201StorageSlot("TeleporterRegistryApp"),
            new ExampleRegistryAppUpgradeable().TELEPORTER_REGISTRY_APP_STORAGE_LOCATION()
        );
    }

    function _erc7201StorageSlot(bytes memory storageName) private pure returns (bytes32) {
        return keccak256(
            abi.encode(uint256(keccak256(abi.encodePacked("teleporter.storage.", storageName))) - 1)
        ) & ~bytes32(uint256(0xff));
    }
}
