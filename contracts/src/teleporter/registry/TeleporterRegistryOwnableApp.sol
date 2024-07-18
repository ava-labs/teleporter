// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.20;

import {TeleporterRegistryApp} from "./TeleporterRegistryApp.sol";
import {Ownable} from "@openzeppelin/contracts@5.0.2/access/Ownable.sol";

/**
 * @dev Contract that inherits {TeleporterRegistryApp} and allows
 * only owners of the contract to update the minimum Teleporter version.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter/blob/main/SECURITY.md
 */
abstract contract TeleporterRegistryOwnableApp is TeleporterRegistryApp, Ownable {
    constructor(
        address teleporterRegistryAddress,
        address initialOwner
    ) TeleporterRegistryApp(teleporterRegistryAddress) {
        transferOwnership(initialOwner);
    }

    /**
     * @dev See {TeleporterRegistryApp-_checkTeleporterUpgradeAccess}
     *
     * Checks that the caller is the owner of the contract for upgrade access.
     */
    function _checkTeleporterUpgradeAccess() internal view virtual override {
        _checkOwner();
    }
}
