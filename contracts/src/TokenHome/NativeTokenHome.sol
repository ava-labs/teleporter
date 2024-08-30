// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {NativeTokenHomeUpgradeable} from "./NativeTokenHomeUpgradeable.sol";
import {ICTTInitializable} from "../utils/ICTTInitializable.sol";

/**
 * @title NativeTokenHome
 * @notice A non-upgradeable version of {NativeTokenHomeUpgradeable} that calls the parent upgradeable contract's initialize function.
 * @custom:security-contact https://github.com/ava-labs/avalanche-interchain-token-transfer/blob/main/SECURITY.md
 */
contract NativeTokenHome is NativeTokenHomeUpgradeable {
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        uint256 minTeleporterVersion,
        address wrappedTokenAddress
    ) NativeTokenHomeUpgradeable(ICTTInitializable.Allowed) {
        initialize(
            teleporterRegistryAddress, teleporterManager, minTeleporterVersion, wrappedTokenAddress
        );
    }
}
