// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.20;

import {TeleporterRegistryOwnableAppUpgradeable} from
    "./TeleporterRegistryOwnableAppUpgradeable.sol";

abstract contract TeleporterRegistryOwnableApp is TeleporterRegistryOwnableAppUpgradeable {
    constructor(address teleporterRegistryAddress, address initialOwner) {
        __TeleporterRegistryOwnableAppUpgradeable_init(teleporterRegistryAddress, initialOwner);
    }
}
