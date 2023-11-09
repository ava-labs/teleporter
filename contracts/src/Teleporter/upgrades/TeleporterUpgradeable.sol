// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterRegistry} from "./TeleporterRegistry.sol";

/**
 * @dev TeleporterUpgradeable provides upgrade utility for applications built on top
 * of the Teleporter protocol by integrating with the {TeleporterRegistry}.
 *
 * This contract is intended to be inherited by other contracts that wish to use the
 * upgrade mechanism. It provides a modifier that restricts access to only Teleporter
 * versions that are greater than or equal to `minTeleporterVersion`.
 */
abstract contract TeleporterUpgradeable {
    TeleporterRegistry public immutable teleporterRegistry;

    /**
     * @dev The minimum required Teleporter version that the contract is allowed
     * to receive messages from. Should only be updated through the `updateMinTeleporterVersion`
     * implementation of child contracts inheriting `TeleporterUpgradeable`. The value is
     * public because inheriting contracts must be able to update it, and it should be
     * publicly viewable.
     */
    uint256 public minTeleporterVersion;

    event MinTeleporterVersionUpdated(
        uint256 indexed oldMinTeleporterVersion,
        uint256 indexed newMinTeleporterVersion
    );

    /**
     * @dev Throws if called by a `msg.sender` that is not an allowed Teleporter version.
     * Checks that `msg.sender` matches a Teleporter version greater than or equal to `minTeleporterVersion`.
     */
    modifier onlyAllowedTeleporter() {
        require(
            teleporterRegistry.getVersionFromAddress(msg.sender) >=
                minTeleporterVersion,
            "TeleporterUpgradeable: invalid teleporter sender"
        );
        _;
    }

    /**
     * @dev Initializes the {TeleporterUpgradeable} contract by getting `teleporterRegistry`
     * instance and setting `_minTeleporterVersion`.
     */
    constructor(address teleporterRegistryAddress) {
        require(
            teleporterRegistryAddress != address(0),
            "TeleporterUpgradeable: zero teleporter registry address"
        );

        teleporterRegistry = TeleporterRegistry(teleporterRegistryAddress);
        minTeleporterVersion = teleporterRegistry.latestVersion();
    }

    /**
     * @dev This is a virtual function that should be overridden to update the `minTeleporterVersion`
     * allowed for modifier `onlyAllowedTeleporter`, and emit {MinTeleporterVersionUpdated} event after.
     *
     * Note: To prevent anyone from being able to call this function, which would disallow messages
     * from old Teleporter versions from being received, this function should be safeguarded with access
     * controls. For example, if the derived contract has an owner/admin, only they can call this function.
     */
    function updateMinTeleporterVersion() external virtual;
}
