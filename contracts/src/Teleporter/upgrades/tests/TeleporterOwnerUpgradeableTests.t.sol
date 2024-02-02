// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterOwnerUpgradeable} from "../TeleporterOwnerUpgradeable.sol";
import {TeleporterUpgradeable} from "../TeleporterUpgradeable.sol";
import {TeleporterUpgradeableTest} from "./TeleporterUpgradeableTests.t.sol";

contract ExampleOwnerUpgradeableApp is TeleporterOwnerUpgradeable {
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress, teleporterManager) {}

    function checkTeleporterUpgradeAccess() external view {
        _checkTeleporterUpgradeAccess();
    }

    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message // solhint-disable-next-line no-empty-blocks
    ) internal override {}
}

contract TeleporterOwnerUpgradeableTest is TeleporterUpgradeableTest {
    ExampleOwnerUpgradeableApp public ownerApp;
    address public constant MOCK_INVALID_OWNER_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address public constant DEFAULT_OWNER_ADDRESS = 0x1234512345123451234512345123451234512345;

    function setUp() public virtual override {
        TeleporterUpgradeableTest.setUp();
        ownerApp =
            new ExampleOwnerUpgradeableApp(address(teleporterRegistry), DEFAULT_OWNER_ADDRESS);
    }

    function testOwnerUpdateMinTeleporterVersion() public {
        uint256 minTeleporterVersion = ownerApp.getMinTeleporterVersion();
        _addProtocolVersion(teleporterRegistry, teleporterAddress);

        // Check that call to update minimum Teleporter version reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.updateMinTeleporterVersion(minTeleporterVersion + 1);

        // Check that minimum Teleporter version was not updated
        assertEq(ownerApp.getMinTeleporterVersion(), minTeleporterVersion);

        // Check that call to update minimum Teleporter version succeeds for owners
        _updateMinTeleporterVersionSuccess(ownerApp, minTeleporterVersion + 1);

        assertEq(ownerApp.getMinTeleporterVersion(), minTeleporterVersion + 1);
    }

    function testOwnerTransfer() public {
        uint256 minTeleporterVersion = ownerApp.getMinTeleporterVersion();
        _addProtocolVersion(teleporterRegistry, teleporterAddress);

        // Check that call to transfer ownership reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.transferOwnership(address(0));

        // Check that ownership was not transferred
        assertEq(ownerApp.owner(), DEFAULT_OWNER_ADDRESS);

        // Check that call for non owners reverts
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.updateMinTeleporterVersion(minTeleporterVersion + 1);

        // Check that after ownership transfer call succeeds
        address newOwner = address(0x123);
        vm.prank(DEFAULT_OWNER_ADDRESS);
        ownerApp.transferOwnership(newOwner);
        vm.prank(newOwner);
        ownerApp.updateMinTeleporterVersion(minTeleporterVersion + 1);

        // Check that call with old owner reverts
        vm.prank(DEFAULT_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.updateMinTeleporterVersion(minTeleporterVersion + 1);
    }

    function testRenounceOwnership() public {
        // Check that call to renounce ownership reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.renounceOwnership();

        // Check that ownership was not renounced
        assertEq(ownerApp.owner(), DEFAULT_OWNER_ADDRESS);
        vm.prank(DEFAULT_OWNER_ADDRESS);
        _addProtocolVersion(teleporterRegistry, teleporterAddress);
        uint256 latestVersion = teleporterRegistry.latestVersion();
        // Check that update Teleporter version call for owner succeeds
        _updateMinTeleporterVersionSuccess(ownerApp, latestVersion);

        // Check that after ownership renounce call reverts
        vm.prank(DEFAULT_OWNER_ADDRESS);
        ownerApp.renounceOwnership();
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.updateMinTeleporterVersion(latestVersion);
        vm.stopPrank();
    }

    function testPauseTeleporterAccess() public {
        // First pause the Teleporter address
        vm.prank(DEFAULT_OWNER_ADDRESS);
        _pauseTeleporterAddressSuccess(ownerApp, teleporterAddress);

        // Try to unpause the Teleporter address from non-owner account
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.unpauseTeleporterAddress(teleporterAddress);

        // Check that the Teleporter address is still paused
        vm.prank(teleporterAddress);
        vm.expectRevert("TeleporterUpgradeable: Teleporter address paused");
        ownerApp.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");

        // Unpause the Teleporter address from owner account
        vm.prank(DEFAULT_OWNER_ADDRESS);
        _unpauseTeleporterAddressSuccess(ownerApp, teleporterAddress);

        // Check that the Teleporter address can now deliver messages
        vm.prank(teleporterAddress);
        ownerApp.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testOwnerUpgradeAccess() public {
        // Check that call to check upgrade access reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.checkTeleporterUpgradeAccess();

        // Check that call to check upgrade access succeeds for owners
        vm.prank(DEFAULT_OWNER_ADDRESS);
        ownerApp.checkTeleporterUpgradeAccess();
    }

    function testInitalOwner() public {
        // Create a new Ownable app with a passed in teleporterManager
        ExampleOwnerUpgradeableApp newOwnerApp =
            new ExampleOwnerUpgradeableApp(address(teleporterRegistry), DEFAULT_OWNER_ADDRESS);

        // Check that the teleporterManager is set correctly
        assertEq(newOwnerApp.owner(), DEFAULT_OWNER_ADDRESS);
        vm.prank(DEFAULT_OWNER_ADDRESS);
        newOwnerApp.checkTeleporterUpgradeAccess();

        // Check that address(this) as the caller is not by default owner
        assertFalse(newOwnerApp.owner() == address(this));
        vm.expectRevert("Ownable: caller is not the owner");
        newOwnerApp.checkTeleporterUpgradeAccess();
    }

    function _updateMinTeleporterVersionSuccess(
        TeleporterUpgradeable app_,
        uint256 newMinTeleporterVersion
    ) internal virtual override {
        vm.expectEmit(true, true, true, true, address(app_));
        emit MinTeleporterVersionUpdated(app_.getMinTeleporterVersion(), newMinTeleporterVersion);
        vm.prank(DEFAULT_OWNER_ADDRESS);
        app_.updateMinTeleporterVersion(newMinTeleporterVersion);
    }
}
