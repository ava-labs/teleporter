// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterUpgradeable} from "../TeleporterUpgradeable.sol";
import {TeleporterRegistryTest, TeleporterMessenger} from "./TeleporterRegistryTests.t.sol";

contract ExampleUpgradeableApp is TeleporterUpgradeable {
    constructor(
        address teleporterRegistryAddress
    ) TeleporterUpgradeable(teleporterRegistryAddress) {}

    function setMinTeleporterVersion(uint256 version) public {
        _setMinTeleporterVersion(version);
    }

    function _receiveTeleporterMessage(
        bytes32 originBlockchainID,
        address originSenderAddress,
        bytes memory message // solhint-disable-next-line no-empty-blocks
    ) internal override {}

    // solhint-disable-next-line no-empty-blocks
    function _checkTeleporterUpgradeAccess() internal override {}
}

contract TeleporterUpgradeableTest is TeleporterRegistryTest {
    bytes32 public constant DEFAULT_ORIGIN_CHAIN_ID =
        bytes32(
            hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd"
        );

    address public constant DEFAULT_ORIGIN_ADDRESS =
        0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;

    event MinTeleporterVersionUpdated(
        uint256 indexed oldMinTeleporterVersion,
        uint256 indexed newMinTeleporterVersion
    );

    function setUp() public virtual override {
        TeleporterRegistryTest.setUp();
        _addProtocolVersion(teleporterRegistry, teleporterAddress);
    }

    function testInvalidRegistryAddress() public {
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage(
                "zero teleporter registry address"
            )
        );
        new ExampleUpgradeableApp(address(0));
    }

    function testOnlyAllowedTeleporter() public {
        ExampleUpgradeableApp app = new ExampleUpgradeableApp(
            address(teleporterRegistry)
        );

        assertEq(app.getMinTeleporterVersion(), 1);

        vm.expectRevert(
            _formatRegistryErrorMessage("protocol address not found")
        );
        app.receiveTeleporterMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_ORIGIN_ADDRESS,
            ""
        );

        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_ORIGIN_ADDRESS,
            ""
        );
    }

    function testUpdateMinTeleporterVersion() public {
        ExampleUpgradeableApp app = new ExampleUpgradeableApp(
            address(teleporterRegistry)
        );

        // First check that calling with initial teleporter address works
        assertEq(app.getMinTeleporterVersion(), 1);
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_ORIGIN_ADDRESS,
            ""
        );

        // Now add new protocol version to registry and update the app's min version
        address newTeleporterAddress = address(new TeleporterMessenger());
        _addProtocolVersion(teleporterRegistry, newTeleporterAddress);

        vm.expectEmit(true, true, true, true, address(app));
        emit MinTeleporterVersionUpdated(1, 2);

        app.updateMinTeleporterVersion(teleporterRegistry.latestVersion());
        assertEq(app.getMinTeleporterVersion(), 2);

        // Check that calling with the old teleporter address fails
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage(
                "invalid Teleporter sender"
            )
        );
        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_ORIGIN_ADDRESS,
            ""
        );

        // Check that calling with the new teleporter address works
        vm.prank(newTeleporterAddress);
        app.receiveTeleporterMessage(
            DEFAULT_ORIGIN_CHAIN_ID,
            DEFAULT_ORIGIN_ADDRESS,
            ""
        );
    }

    function testSetMinTeleporterVersion() public {
        ExampleUpgradeableApp app = new ExampleUpgradeableApp(
            address(teleporterRegistry)
        );

        uint256 latestVersion = teleporterRegistry.latestVersion();

        // Check setting for a version > latest version fails
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage(
                "invalid Teleporter version"
            )
        );
        app.setMinTeleporterVersion(latestVersion + 1);

        // Check setting for a version <= min version fails
        uint256 minVersion = app.getMinTeleporterVersion();
        assertEq(minVersion, teleporterRegistry.latestVersion());
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage(
                "not greater than current minimum version"
            )
        );
        app.setMinTeleporterVersion(minVersion);

        // Add a new protocol version to the registry
        _addProtocolVersion(teleporterRegistry, teleporterAddress);

        // Check setting a new valid minimum version
        vm.expectEmit(true, true, true, true, address(app));
        emit MinTeleporterVersionUpdated(latestVersion, latestVersion + 1);
        app.setMinTeleporterVersion(teleporterRegistry.latestVersion());
        assertEq(app.getMinTeleporterVersion(), latestVersion + 1);
    }

    function _formatTeleporterUpgradeableErrorMessage(
        string memory errorMessage
    ) internal pure returns (bytes memory) {
        return bytes(string.concat("TeleporterUpgradeable: ", errorMessage));
    }
}
