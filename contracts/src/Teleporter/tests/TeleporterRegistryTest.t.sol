// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "forge-std/Test.sol";
import "../TeleporterRegistry.sol";

contract TeleporterRegistryTest is Test {
    TeleporterRegistry teleporterRegistry;

    function setUp() public override {
        teleporterRegistry = new TeleporterRegistry(
            new uint256[](0),
            new address[](0)
        );
    }

    function test_getLatestVersion() public {
        Assert.equal(teleporterRegistry.getLatestVersion(), 0, "should be 0");
    }

    function test_addProtocolVersion() public {
        teleporterRegistry.addProtocolVersion(address(0x1), 1);
        Assert.equal(teleporterRegistry.getLatestVersion(), 1, "should be 1");
        Assert.equal(
            teleporterRegistry.getAddressToVersion(address(0x1)),
            1,
            "should be 1"
        );
    }

    function test_getVersionToAddress() public {
        teleporterRegistry.addProtocolVersion(address(0x1), 1);
        Assert.equal(
            teleporterRegistry.getVersionToAddress(1),
            address(0x1),
            "should be 0x1"
        );
    }

    function test_getAddressToVersion() public {
        teleporterRegistry.addProtocolVersion(address(0x1), 1);
        Assert.equal(
            teleporterRegistry.getAddressToVersion(address(0x1)),
            1,
            "should be 1"
        );
    }

    function test_getTeleporter() public {
        teleporterRegistry.addProtocolVersion(address(0x1), 1);
        Assert.equal(
            teleporterRegistry.getTeleporter(1),
            address(0x1),
            "should be 0x1"
        );
    }

    function test_getLatestTeleporter() public {
        teleporterRegistry.addProtocolVersion(address(0x1), 1);
        Assert.equal(
            teleporterRegistry.getLatestTeleporter(),
            address(0x1),
            "should be 0x1"
        );
    }
}
