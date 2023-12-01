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
 * upgrade mechanism. It provides an interface that restricts access to only Teleporter
 * versions that are greater than or equal to `minTeleporterVersion`.
 */
abstract contract TeleporterUpgradeable is ITeleporterReceiver {
    TeleporterRegistry public immutable teleporterRegistry;

    /**
     * @dev The minimum required Teleporter version that the contract is allowed
     * to receive messages from. Should only be updated by `_setMinTeleporterVersion`
     */
    uint256 private minTeleporterVersion;

    /**
     * @dev Emitted when `minTeleporterVersion` is updated.
     */
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

    /**
     * @dev See {ITeleporterReceiver-receiveTeleporterMessage}
     *
     * Checks that the caller is a Teleporter version greater than or equal to `minTeleporterVersion`.
     */
    function receiveTeleporterMessage(
        bytes32 originBlockchainID,
        address originSenderAddress,
        bytes calldata message
    ) external {
        // Checks that `msg.sender` matches a Teleporter version greater than or equal to `minTeleporterVersion`.
        require(
            teleporterRegistry.getVersionFromAddress(msg.sender) >=
                minTeleporterVersion,
            "TeleporterUpgradeable: invalid Teleporter sender"
        );

        _receiveTeleporterMessage(
            originBlockchainID,
            originSenderAddress,
            message
        );
    }

    /**
     * @dev Updates the minimum Teleporter version allowed for delivering Teleporer messages
     * to this contract.
     *
     * To prevent anyone from being able to call this function, which would disallow messages
     * from old Teleporter versions from being received, this function should be safeguarded with access
     * controls. This is done by overriding the implementation of {_checkTeleporterUpgradeAccess}.
     */
    function updateMinTeleporterVersion(uint256 version) public virtual {
        _checkTeleporterUpgradeAccess();
        _setMinTeleporterVersion(version);
    }

    /**
     * @dev Public getter for `minTeleporterVersion`.
     */
    function getMinTeleporterVersion() public view returns (uint256) {
        return minTeleporterVersion;
    }

    /**
     * @dev Receives Teleporter messages and handles accordingly.
     * This function should be overridden by contracts that inherit from this contract.
     */
    function _receiveTeleporterMessage(
        bytes32 originBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal virtual;

    /**
     * @dev Checks that the caller has access to update the minimum Teleporter version
     * allowed for delivering Teleporter messages to this contract.
     *
     * This function should be overridden by contracts that inherit from this contract.
     */
    function _checkTeleporterUpgradeAccess() internal virtual;

    /**
     * @dev Sets the minimum Teleporter version allowed for delivering Teleporter messages.
     * Emits a {MinTeleporterVersionUpdated} event if the minimum Teleporter version was updated.
     * Requirements:
     *
     * - `version` must be less than or equal to the latest Teleporter version.
     * - `version` must be greater than the current minimum Teleporter version.
     *
     */
    function _setMinTeleporterVersion(uint256 version) internal virtual {
        uint256 latestTeleporterVersion = teleporterRegistry.latestVersion();
        uint256 oldMinTeleporterVersion = minTeleporterVersion;

        require(
            version <= latestTeleporterVersion,
            "TeleporterUpgradeable: invalid Teleporter version"
        );
        require(
            version > oldMinTeleporterVersion,
            "TeleporterUpgradeable: not greater than current minimum version"
        );

        minTeleporterVersion = version;
        emit MinTeleporterVersionUpdated(oldMinTeleporterVersion, version);
    }
}
