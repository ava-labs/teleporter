// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {NativeTokenHomeUpgradeable} from "./NativeTokenHomeUpgradeable.sol";
import {Initializable} from "../utils/Initializable.sol";

contract NativeTokenHome is NativeTokenHomeUpgradeable {
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        address wrappedTokenAddress
    ) NativeTokenHomeUpgradeable(Initializable.Allowed) {
        initialize(teleporterRegistryAddress, teleporterManager, wrappedTokenAddress);
    }
}
