// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "./TeleporterRegistry.sol";

/**
 * @dev TeleporterUpgradeable provides upgrade utility for applications built on top
 * of the Teleporter protocol by integrating with the {TeleporterRegistry}.
 *
 * This contract is intended to be inherited by other contracts that wish to use the
 * upgrade mechanism. It provides a modifier that restricts access to only Teleporter
 * versions that are greater than or equal to `_minTeleporterVersion`.
 */
abstract contract TeleporterUpgradeable {
    TeleporterRegistry public immutable teleporterRegistry;
    uint256 internal _minTeleporterVersion;

    error InvalidTeleporterRegistryAddress();
    error InvalidTeleporterSender();

    /**
     * @dev Throws if called by a `msg.sender` that is not an allowed Teleporter version.
     * Checks that `msg.sender` matches a Teleporter version greater than or equal to `_minTeleporterVersion`.
     */
    modifier onlyAllowedTeleporter() {
        if (
            teleporterRegistry.getVersionFromAddress(msg.sender) <
            _minTeleporterVersion
        ) {
            revert InvalidTeleporterSender();
        }
        _;
    }

    /**
     * @dev Initializes the {TeleporterUpgradeable} contract by getting `teleporterRegistry`
     * instance and setting `_minTeleporterVersion`.
     */
    constructor(address teleporterRegistryAddress) {
        if (teleporterRegistryAddress == address(0)) {
            revert InvalidTeleporterRegistryAddress();
        }

        teleporterRegistry = TeleporterRegistry(teleporterRegistryAddress);
        _minTeleporterVersion = teleporterRegistry.getLatestVersion();
    }

    /**
     * @dev Updates `_minTeleporterVersion` to the latest version.
     */
    function updateMinTeleporterVersion() external {
        _minTeleporterVersion = teleporterRegistry.getLatestVersion();
    }

    /**
     * @dev Gets the minimum Teleporter version that is allowed to call this contract.
     */
    function getMinTeleporterVersion() external view returns (uint256) {
        return _minTeleporterVersion;
    }
}
