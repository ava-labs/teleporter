// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.23;

import {GetTeleporterMessengerTest} from "./GetTeleporterMessengerTests.t.sol";
import {PauseTeleporterAddressTest} from "./PauseTeleporterAddressTests.t.sol";
import {UpdateMinTeleporterVersionTest} from "./UpdateMinTeleporterVersionTests.t.sol";
import {SendTeleporterMessageTest} from "./SendTeleporterMessageTests.t.sol";
import {UnpauseTeleporterAddressTest} from "./UnpauseTeleporterAddressTests.t.sol";
import {BaseTeleporterRegistryAppTest} from "./BaseTeleporterRegistryAppTests.t.sol";
import {ExampleRegistryApp} from "./BaseTeleporterRegistryAppTests.t.sol";

contract TeleporterRegistryAppTest is
    GetTeleporterMessengerTest,
    PauseTeleporterAddressTest,
    UpdateMinTeleporterVersionTest,
    SendTeleporterMessageTest,
    UnpauseTeleporterAddressTest
{
    function setUp() public virtual override {
        BaseTeleporterRegistryAppTest.setUp();
        app = new ExampleRegistryApp(address(teleporterRegistry));
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
