// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "forge-std/Test.sol";
import "../TeleporterOwnerUpgradeable.sol";
import "./TeleporterUpgradeableTests.t.sol";

contract ExampleOwnerUpgradeableApp is TeleporterOwnerUpgradeable {
    constructor(
        address teleporterRegistryAddress
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress) {}
}

contract TeleporterOwnerUpgradeableTest is TeleporterUpgradeableTest {
    ExampleOwnerUpgradeableApp app;

    function setUp() public virtual override {
        TeleporterUpgradeableTest.setUp();
        app = new ExampleOwnerUpgradeableApp(address(teleporterRegistry));
    }

    function testOwnerUpdateMinTeleporterVersion() public {
        uint256 minTeleporterVersion = app.minTeleporterVersion();

        // Check that call to update minimum Teleporter version reverts for non-owners
        vm.prank(teleporterAddress);
        vm.expectRevert("Ownable: caller is not the owner");
        app.updateMinTeleporterVersion();

        // Check that minimum Teleporter version was not updated
        assertEq(app.minTeleporterVersion(), minTeleporterVersion);

        _addProtocolVersion(teleporterRegistry, teleporterAddress);

        // Check that call to update minimum Teleporter version succeeds for owners
        vm.prank(address(this));
        app.updateMinTeleporterVersion();

        assertEq(app.minTeleporterVersion(), minTeleporterVersion + 1);
    }

    function testOwnerTransfer() public {
        // Check that call to transfer ownership reverts for non-owners
        vm.prank(teleporterAddress);
        vm.expectRevert("Ownable: caller is not the owner");
        app.transferOwnership(address(0));

        // Check that ownership was not transferred
        assertEq(app.owner(), address(this));

        // Check that call for non owners reverts
        vm.prank(teleporterAddress);
        vm.expectRevert("Ownable: caller is not the owner");
        app.updateMinTeleporterVersion();

        // Check that after ownership transfer call succeeds
        app.transferOwnership(teleporterAddress);
        vm.prank(teleporterAddress);
        app.updateMinTeleporterVersion();

        // Check that call with old owner reverts
        vm.prank(address(this));
        vm.expectRevert("Ownable: caller is not the owner");
        app.updateMinTeleporterVersion();
    }

    function testRenounceOwnership() public {
        // Check that call to renounce ownership reverts for non-owners
        vm.prank(teleporterAddress);
        vm.expectRevert("Ownable: caller is not the owner");
        app.renounceOwnership();

        // Check that ownership was not renounced
        assertEq(app.owner(), address(this));

        // Check that update Teleporter version call for owner succeeds
        app.updateMinTeleporterVersion();

        // Check that after ownership renounce call reverts
        app.renounceOwnership();
        vm.expectRevert("Ownable: caller is not the owner");
        app.updateMinTeleporterVersion();
    }
}
