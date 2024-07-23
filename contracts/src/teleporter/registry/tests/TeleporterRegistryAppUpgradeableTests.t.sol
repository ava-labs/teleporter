// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.23;

import {TeleporterRegistryAppTest} from "./TeleporterRegistryAppTests.t.sol";
import {
    ExampleRegistryAppUpgradeable,
    ExampleRegistryApp
} from "./BaseTeleporterRegistryAppTests.t.sol";

contract TeleporterRegistryAppUpgradeableTest is TeleporterRegistryAppTest {
    function setUp() public virtual override {
        TeleporterRegistryAppTest.setUp();
        ExampleRegistryAppUpgradeable upgradeableApp = new ExampleRegistryAppUpgradeable();
        upgradeableApp.initialize(address(teleporterRegistry));
        app = ExampleRegistryApp(address(upgradeableApp));
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
