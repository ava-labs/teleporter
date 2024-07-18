// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.20;

import {TeleporterRegistryAppUpgradeable} from "./TeleporterRegistryAppUpgradeable.sol";

abstract contract TeleporterRegistryApp is TeleporterRegistryAppUpgradeable {
    constructor(address teleporterRegistryAddress) {
        __TeleporterRegistryAppUpgradeable_init(teleporterRegistryAddress);
    }
}
