// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterOwnerUpgradeable} from "../TeleporterOwnerUpgradeable.sol";
import {TeleporterUpgradeableTest} from "./TeleporterUpgradeableTests.t.sol";

contract ExampleOwnerUpgradeableApp is TeleporterOwnerUpgradeable {
    constructor(
        address teleporterRegistryAddress
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress) {}
}

contract TeleporterOwnerUpgradeableTest is TeleporterUpgradeableTest {
    ExampleOwnerUpgradeableApp public app;
    address public constant MOCK_INVALID_OWNER_ADDRESS =
        0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;

    function setUp() public virtual override {
        TeleporterUpgradeableTest.setUp();
        app = new ExampleOwnerUpgradeableApp(address(teleporterRegistry));
    }

    function testOwnerUpdateMinTeleporterVersion() public {
        uint256 minTeleporterVersion = app.minTeleporterVersion();

        // Check that call to update minimum Teleporter version reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        app.updateMinTeleporterVersion();

        // Check that minimum Teleporter version was not updated
        assertEq(app.minTeleporterVersion(), minTeleporterVersion);

        _addProtocolVersion(teleporterRegistry, teleporterAddress);

        // Check that call to update minimum Teleporter version succeeds for owners
        vm.prank(address(this));
        vm.expectEmit(true, true, true, true, address(app));
        emit MinTeleporterVersionUpdated(
            minTeleporterVersion,
            minTeleporterVersion + 1
        );
        app.updateMinTeleporterVersion();

        assertEq(app.minTeleporterVersion(), minTeleporterVersion + 1);
    }

    function testOwnerTransfer() public {
        uint256 minTeleporterVersion = app.minTeleporterVersion();
        _addProtocolVersion(teleporterRegistry, teleporterAddress);

        // Check that call to transfer ownership reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        app.transferOwnership(address(0));

        // Check that ownership was not transferred
        assertEq(app.owner(), address(this));

        // Check that call for non owners reverts
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        app.updateMinTeleporterVersion();

        // Check that after ownership transfer call succeeds
        address newOwner = address(0x123);
        app.transferOwnership(newOwner);
        vm.prank(newOwner);
        vm.expectEmit(true, true, true, true, address(app));
        emit MinTeleporterVersionUpdated(
            minTeleporterVersion,
            minTeleporterVersion + 1
        );
        app.updateMinTeleporterVersion();

        // Check that call with old owner reverts
        vm.expectRevert("Ownable: caller is not the owner");
        app.updateMinTeleporterVersion();
    }

    function testRenounceOwnership() public {
        // Check that call to renounce ownership reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
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
