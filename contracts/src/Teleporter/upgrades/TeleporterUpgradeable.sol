// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "./TeleporterRegistry.sol";

/**
 * @dev TeleporterUpgradeable provides upgrade utility for applications built on top
 * of the Teleporter protocol by integrating with the {TeleporterRegistry}.
 */
abstract contract TeleporterUpgradeable {
    TeleporterRegistry public immutable teleporterRegistry;
    uint256 internal _minTeleporterVersion;

    error InvalidTeleporterRegistryAddress();
    error InvalidTeleporterSender();

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
     * @dev Throws if called by a `msg.sender` that is not an allowed Telepoter version.
     * Checks that `msg.sender` matches a Teleporter version greater than or equal to `_minTeleporterVersion`.
     */
    modifier onlyAllowedTeleporter() {
        if (
            teleporterRegistry.getAddressToVersion(msg.sender) <
            _minTeleporterVersion
        ) {
            revert InvalidTeleporterSender();
        }
        _;
    }

    /**
     * @dev Updates `_minTeleporterVersion` to the latest version.
     */
    function updateMinTeleporterversion() external {
        _minTeleporterVersion = teleporterRegistry.getLatestVersion();
    }

    /**
     * @dev Gets the minimum Teleporter version that is allowed to call this contract.
     */
    function getMinTeleporterVersion() external view returns (uint256) {
        return _minTeleporterVersion;
    }
}
