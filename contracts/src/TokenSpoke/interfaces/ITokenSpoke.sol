// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITokenBridge} from "../../interfaces/ITokenBridge.sol";
import {TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Settings for constructing a {ITokenSpoke} contract.
 * @param teleporterRegistryAddress The current blockchain ID's Teleporter registry
 * address. See here for details: https://github.com/ava-labs/teleporter/tree/main/contracts/src/Teleporter/upgrades.
 * @param teleporterManager Address that manages this contract's integration with the
 * Teleporter registry and Teleporter versions.
 * @param tokenHubBlockchainID The blockchain ID of the associated token hub instance.
 * @param tokenHubAddress The address of the associated token hub instance on the {tokenHubBlockchainID}.
 */
struct TokenSpokeSettings {
    address teleporterRegistryAddress;
    address teleporterManager;
    bytes32 tokenHubBlockchainID;
    address tokenHubAddress;
}

/**
 * @dev Interface for a spoke bridge contract that mints a representation token on its chain, and allows
 * for burning that token to redeem the backing asset on the hub chain, or bridging to other spokes.
 */
interface ITokenSpoke is ITokenBridge {
    /**
     * @notice Sends a Teleporter message to register the spoke instance with its configured hub.
     * @param feeInfo The fee asset and amount for the Teleporter message.
     */
    function registerWithHub(TeleporterFeeInfo calldata feeInfo) external;
}
