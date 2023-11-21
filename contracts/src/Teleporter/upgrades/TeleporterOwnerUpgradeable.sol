// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterUpgradeable} from "./TeleporterUpgradeable.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @dev Contract that implements the {TeleporterUpgradeable} interface and allows
 * only owners of the contract to update the minimum Teleporter version.
 */
abstract contract TeleporterOwnerUpgradeable is TeleporterUpgradeable, Ownable {
    constructor(
        address teleporterRegistryAddress
    ) TeleporterUpgradeable(teleporterRegistryAddress) {}

    /**
     * @dev See {TeleporterUpgradeable-updateMinTeleporterVersion}
     *
     * Updates the minimum Teleporter version allowed for receiving on this contract
     * to the latest version registered in the {TeleporterRegistry}.
     * Restricted to only owners of the contract.
     * Emits a {MinTeleporterVersionUpdated} event if the minimum Teleporter version
     * was updated.
     */
    function updateMinTeleporterVersion() external override onlyOwner {
        uint256 oldMinTeleporterVersion = minTeleporterVersion;
        minTeleporterVersion = teleporterRegistry.latestVersion();
        if (minTeleporterVersion > oldMinTeleporterVersion) {
            emit MinTeleporterVersionUpdated(
                oldMinTeleporterVersion,
                minTeleporterVersion
            );
        }
    }
}
