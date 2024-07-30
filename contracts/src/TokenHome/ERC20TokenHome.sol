// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ERC20TokenHomeUpgradeable} from "./ERC20TokenHomeUpgradeable.sol";
import {ICTTInitializable} from "../utils/ICTTInitializable.sol";

/**
 * @title ERC20TokenHome
 * @notice A non-upgradeable version of {ERC20TokenHomeUpgradeable} that calls the parent upgradeable contract's initialize function.
 * @custom:security-contact https://github.com/ava-labs/avalanche-interchain-token-transfer/blob/main/SECURITY.md
 */
contract ERC20TokenHome is ERC20TokenHomeUpgradeable {
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        address tokenAddress,
        uint8 tokenDecimals
    ) ERC20TokenHomeUpgradeable(ICTTInitializable.Allowed) {
        initialize(teleporterRegistryAddress, teleporterManager, tokenAddress, tokenDecimals);
    }
}
