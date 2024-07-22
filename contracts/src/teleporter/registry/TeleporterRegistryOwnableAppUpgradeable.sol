// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.23;

import {TeleporterRegistryAppUpgradeable} from "./TeleporterRegistryAppUpgradeable.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";

/**
 * @dev Contract that inherits {TeleporterRegistryAppUpgradeable} and allows
 * only owners of the contract to update the minimum Teleporter version.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter/blob/main/SECURITY.md
 */
abstract contract TeleporterRegistryOwnableAppUpgradeable is
    TeleporterRegistryAppUpgradeable,
    OwnableUpgradeable
{
    // solhint-disable-next-line func-name-mixedcase
    function __TeleporterRegistryOwnableAppUpgradeable_init(
        address teleporterRegistryAddress,
        address initialOwner
    ) internal onlyInitializing {
        __TeleporterRegistryAppUpgradeable_init(teleporterRegistryAddress);
        __Ownable_init(initialOwner);
    }

    // solhint-disable-next-line func-name-mixedcase, no-empty-blocks, func-name-mixedcase
    function _TeleporterRegistryOwnableAppUpgradeable_init_unchained() internal onlyInitializing {}

    /**
     * @dev See {TeleporterRegistryAppUpgradeable-_checkTeleporterUpgradeAccess}
     *
     * Checks that the caller is the owner of the contract for upgrade access.
     */
    function _checkTeleporterUpgradeAccess() internal view virtual override {
        _checkOwner();
    }
}
