// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterRegistry} from "./TeleporterRegistry.sol";
import {ITeleporterReceiver} from "../ITeleporterReceiver.sol";

/**
 * @dev TeleporterUpgradeable provides upgrade utility for applications built on top
 * of the Teleporter protocol by integrating with the {TeleporterRegistry}.
 *
 * This contract is intended to be inherited by other contracts that wish to use the
 * upgrade mechanism. It provides a modifier that restricts access to only Teleporter
 * versions that are greater than or equal to `minTeleporterVersion`.
 */
abstract contract TeleporterUpgradeable is ITeleporterReceiver {
    TeleporterRegistry public immutable teleporterRegistry;

    /**
     * @dev The minimum required Teleporter version that the contract is allowed
     * to receive messages from. Should only be updated through the `updateMinTeleporterVersion`
     * implementation of child contracts inheriting `TeleporterUpgradeable`. The value is
     * public because inheriting contracts must be able to update it, and it should be
     * publicly viewable.
     */
    uint256 private minTeleporterVersion;

    event MinTeleporterVersionUpdated(
        uint256 indexed oldMinTeleporterVersion,
        uint256 indexed newMinTeleporterVersion
    );

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

    function receiveTeleporterMessage(
        bytes32 originBlockchainID,
        address originSenderAddress,
        bytes calldata message
    ) external {
        // Checks that `msg.sender` matches a Teleporter version greater than or equal to `minTeleporterVersion`.
        require(
            teleporterRegistry.getVersionFromAddress(msg.sender) >=
                minTeleporterVersion,
            "TeleporterUpgradeable: invalid teleporter sender"
        );

        _receiveTeleporterMessage(
            originBlockchainID,
            originSenderAddress,
            message
        );
    }

    /**
     * @dev This is a virtual function that should be overridden to update the `minTeleporterVersion`
     * allowed for modifier `onlyAllowedTeleporter`, and emit {MinTeleporterVersionUpdated} event after.
     *
     * Note: To prevent anyone from being able to call this function, which would disallow messages
     * from old Teleporter versions from being received, this function should be safeguarded with access
     * controls. For example, if the derived contract has an owner/admin, only they can call this function.
     */
    function updateMinTeleporterVersion(uint256 version) public virtual {
        _checkTeleporterUpgradeAccess();
        _setMinTeleporterVersion(version);
    }

    function getMinTeleporterVersion() public view returns (uint256) {
        return minTeleporterVersion;
    }

    function _receiveTeleporterMessage(
        bytes32 originBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal virtual;

    function _checkTeleporterUpgradeAccess() internal virtual;

    function _setMinTeleporterVersion(uint256 version) private {
        uint256 latestTeleporterVersion = teleporterRegistry.latestVersion();
        uint256 oldMinTeleporterVersion = minTeleporterVersion;

        require(
            version <= latestTeleporterVersion,
            "TeleporterUpgradeable: invalid version"
        );
        require(
            version > oldMinTeleporterVersion,
            "TeleporterUpgradeable: not greater than current version"
        );

        minTeleporterVersion = version;
        emit MinTeleporterVersionUpdated(oldMinTeleporterVersion, version);
    }
}
