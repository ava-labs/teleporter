// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterOwnerUpgradeable} from "../TeleporterOwnerUpgradeable.sol";
import {TeleporterUpgradeableTest} from "./TeleporterUpgradeableTests.t.sol";

contract ExampleOwnerUpgradeableApp is TeleporterOwnerUpgradeable {
    constructor(address teleporterRegistryAddress) TeleporterOwnerUpgradeable(teleporterRegistryAddress) {}

    function checkTeleporterUpgradeAccess() external view {
        _checkTeleporterUpgradeAccess();
    }

    function _receiveTeleporterMessage(
        bytes32 originBlockchainID,
        address originSenderAddress,
        bytes memory message // solhint-disable-next-line no-empty-blocks
    ) internal override {}
}

contract TeleporterOwnerUpgradeableTest is TeleporterUpgradeableTest {
    ExampleOwnerUpgradeableApp public ownerApp;
    address public constant MOCK_INVALID_OWNER_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;

    function setUp() public virtual override {
        TeleporterUpgradeableTest.setUp();
        ownerApp = new ExampleOwnerUpgradeableApp(address(teleporterRegistry));
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
        vm.prank(address(this));
        vm.expectEmit(true, true, true, true, address(ownerApp));
        emit MinTeleporterVersionUpdated(minTeleporterVersion, minTeleporterVersion + 1);
        ownerApp.updateMinTeleporterVersion(minTeleporterVersion + 1);

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
        assertEq(ownerApp.owner(), address(this));

        // Check that call for non owners reverts
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.updateMinTeleporterVersion(minTeleporterVersion + 1);

        // Check that after ownership transfer call succeeds
        address newOwner = address(0x123);
        ownerApp.transferOwnership(newOwner);
        vm.prank(newOwner);
        vm.expectEmit(true, true, true, true, address(ownerApp));
        emit MinTeleporterVersionUpdated(minTeleporterVersion, minTeleporterVersion + 1);
        ownerApp.updateMinTeleporterVersion(minTeleporterVersion + 1);

        // Check that call with old owner reverts
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.updateMinTeleporterVersion(minTeleporterVersion + 1);
    }

    function testRenounceOwnership() public {
        // Check that call to renounce ownership reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.renounceOwnership();

        // Check that ownership was not renounced
        assertEq(ownerApp.owner(), address(this));
        _addProtocolVersion(teleporterRegistry, teleporterAddress);
        uint256 latestVersion = teleporterRegistry.latestVersion();
        // Check that update Teleporter version call for owner succeeds
        ownerApp.updateMinTeleporterVersion(latestVersion);

        // Check that after ownership renounce call reverts
        ownerApp.renounceOwnership();
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.updateMinTeleporterVersion(latestVersion);
    }

    function testPauseTeleporterAccess() public {
        // Check that call to pause Teleporter address reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.pauseTeleporterAddress(teleporterAddress);

        // Check that Teleporter address was not paused
        vm.prank(teleporterAddress);
        ownerApp.receiveTeleporterMessage(DEFAULT_ORIGIN_CHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");

        // Check that call to pause Teleporter address succeeds for owners
        vm.prank(address(this));
        ownerApp.pauseTeleporterAddress(teleporterAddress);

        // Check that Teleporter address was paused
        vm.prank(teleporterAddress);
        vm.expectRevert(_formatTeleporterUpgradeableErrorMessage("Teleporter address paused"));
        ownerApp.receiveTeleporterMessage(DEFAULT_ORIGIN_CHAIN_ID, DEFAULT_ORIGIN_ADDRESS, "");
    }

    function testOwnerUpgradeAccess() public {
        // Check that call to check upgrade access reverts for non-owners
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.checkTeleporterUpgradeAccess();

        // Check that call to check upgrade access succeeds for owners
        vm.prank(address(this));
        ownerApp.checkTeleporterUpgradeAccess();
    }
}
