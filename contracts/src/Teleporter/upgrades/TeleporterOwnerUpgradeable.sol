// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterUpgradeable} from "./TeleporterUpgradeable.sol";
import {Ownable} from "@openzeppelin/contracts@4.8.1/access/Ownable.sol";

/**
 * @dev Contract that inherits {TeleporterUpgradeable} and allows
 * only owners of the contract to update the minimum Teleporter version.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter/blob/main/SECURITY.md
 */
abstract contract TeleporterOwnerUpgradeable is TeleporterUpgradeable, Ownable {
    constructor(address teleporterRegistryAddress)
        TeleporterUpgradeable(teleporterRegistryAddress)
    {}

    /**
     * @dev See {TeleporterUpgradeable-_checkTeleporterUpgradeAccess}
     *
     * Checks that the caller is the owner of the contract for upgrade access.
     */
    function _checkTeleporterUpgradeAccess() internal view virtual override {
        _checkOwner();
    }
}
