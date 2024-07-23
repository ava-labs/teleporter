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

contract TeleporterRegistryOwnableAppTest is BaseTeleporterRegistryOwnableAppTest {
    function setUp() public virtual override {
        BaseTeleporterRegistryOwnableAppTest.setUp();
        ownerApp = new ExampleRegistryOwnableApp(address(teleporterRegistry), DEFAULT_OWNER_ADDRESS);
        app = ExampleRegistryApp(address(ownerApp));
    }

    function _formatErrorMessage(string memory errorMessage)
        internal
        pure
        virtual
        override
        returns (bytes memory)
    {
        return bytes(string.concat("TeleporterRegistryApp: ", errorMessage));
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

    function _formatErrorMessage(string memory errorMessage)
        internal
        pure
        virtual
        override
        returns (bytes memory)
    {
        return bytes(string.concat("TeleporterRegistryAppUpgradeable: ", errorMessage));
    }
}
