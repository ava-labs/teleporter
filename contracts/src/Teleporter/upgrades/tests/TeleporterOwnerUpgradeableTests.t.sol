// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterOwnerUpgradeable} from "../TeleporterOwnerUpgradeable.sol";
import {TeleporterUpgradeableTest, ExampleUpgradeableApp} from "./TeleporterUpgradeableTests.t.sol";

contract ExampleOwnerUpgradeableApp is TeleporterOwnerUpgradeable {
    constructor(
        address teleporterRegistryAddress
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress) {}

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
    address public constant MOCK_INVALID_OWNER_ADDRESS =
        0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;

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
        _updateMinTeleporterVersionSuccess(minTeleporterVersion + 1);

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
        _updateMinTeleporterVersionSuccess(latestVersion);

        // Check that after ownership renounce call reverts
        ownerApp.renounceOwnership();
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.updateMinTeleporterVersion(latestVersion);
    }

    function testPauseTeleporterAccess() public {
        // First pause the Teleporter address
        _pauseTeleporterAddressSuccess(teleporterAddress);

        // Try to unpause the Teleporter address from non-owner account
        vm.prank(MOCK_INVALID_OWNER_ADDRESS);
        vm.expectRevert("Ownable: caller is not the owner");
        ownerApp.unpauseTeleporterAddress(teleporterAddress);

        // Check that the Teleporter address is still paused
        vm.prank(teleporterAddress);
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage(
                "Teleporter address paused"
            )
        );
        ownerApp.receiveTeleporterMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_ORIGIN_ADDRESS,
            ""
        );
        assertTrue(ownerApp.isTeleporterAddressPaused(teleporterAddress));

        // Unpause the Teleporter address from owner account
        _unpauseTeleporterAddressSuccess(teleporterAddress);

        // Check that the Teleporter address can now deliver messages
        vm.prank(teleporterAddress);
        ownerApp.receiveTeleporterMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_ORIGIN_ADDRESS,
            ""
        );
        assertFalse(ownerApp.isTeleporterAddressPaused(teleporterAddress));
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

    function _updateMinTeleporterVersionSuccess(
        uint256 newMinTeleporterVersion
    ) internal virtual override {
        vm.expectEmit(true, true, true, true, address(ownerApp));
        emit MinTeleporterVersionUpdated(
            ownerApp.getMinTeleporterVersion(),
            newMinTeleporterVersion
        );
        ownerApp.updateMinTeleporterVersion(newMinTeleporterVersion);
    }

    function _pauseTeleporterAddressSuccess(
        address teleporterAddress_
    ) internal virtual override {
        vm.expectEmit(true, true, true, true, address(ownerApp));
        emit TeleporterAddressPaused(teleporterAddress_);
        ownerApp.pauseTeleporterAddress(teleporterAddress_);
    }

    function _unpauseTeleporterAddressSuccess(
        address teleporterAddress_
    ) internal virtual override {
        vm.expectEmit(true, true, true, true, address(ownerApp));
        emit TeleporterAddressUnpaused(teleporterAddress_);
        ownerApp.unpauseTeleporterAddress(teleporterAddress_);
    }
}
