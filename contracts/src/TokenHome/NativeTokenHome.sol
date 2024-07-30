// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {NativeTokenHomeUpgradeable} from "./NativeTokenHomeUpgradeable.sol";
import {ICTTInitializable} from "../utils/ICTTInitializable.sol";

contract NativeTokenHome is NativeTokenHomeUpgradeable {
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        address wrappedTokenAddress
    ) NativeTokenHomeUpgradeable(ICTTInitializable.Allowed) {
        initialize(teleporterRegistryAddress, teleporterManager, wrappedTokenAddress);
    }
}
