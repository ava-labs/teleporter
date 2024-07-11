// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.20;

import {TeleporterUpgradeable} from "./TeleporterUpgradeable.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@4.9.6/access/OwnableUpgradeable.sol";

/**
 * @dev Contract that inherits {TeleporterUpgradeable} and allows
 * only owners of the contract to update the minimum Teleporter version.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter/blob/main/SECURITY.md
 */
abstract contract TeleporterOwnerUpgradeable is TeleporterUpgradeable, OwnableUpgradeable {
    // solhint-disable-next-line func-name-mixedcase
    function __TeleporterOwnerUpgradeable_init(
        address teleporterRegistryAddress,
        address initialOwner
    ) internal onlyInitializing {
        __TeleporterUpgradeable_init(teleporterRegistryAddress);
        __Ownable_init_unchained();
        _TeleporterOwnerUpgradeable_init_unchained(initialOwner);
    }

    function _TeleporterOwnerUpgradeable_init_unchained(address initialOwner)
        internal
        onlyInitializing
    {
        transferOwnership(initialOwner);
    }

    /**
     * @dev See {TeleporterUpgradeable-_checkTeleporterUpgradeAccess}
     *
     * Checks that the caller is the owner of the contract for upgrade access.
     */
    function _checkTeleporterUpgradeAccess() internal view virtual override {
        _checkOwner();
    }
}
