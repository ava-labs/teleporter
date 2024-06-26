// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITokenBridge} from "../../interfaces/ITokenBridge.sol";
import {TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";

/**
 * @notice Settings for constructing a {ITokenRemote} contract.
 * @param teleporterRegistryAddress The current blockchain ID's Teleporter registry
 * address. See here for details: https://github.com/ava-labs/teleporter/tree/main/contracts/src/Teleporter/upgrades.
 * @param teleporterManager Address that manages this contract's integration with the
 * Teleporter registry and Teleporter versions.
 * @param tokenHomeBlockchainID The blockchain ID of the associated TokenHome instance.
 * @param tokenHomeAddress The address of the associated token home contract.
 * @param tokenHomeDecimals The number of decimal places used by the token home's token.
 */
struct TokenRemoteSettings {
    address teleporterRegistryAddress;
    address teleporterManager;
    bytes32 tokenHomeBlockchainID;
    address tokenHomeAddress;
    uint8 tokenHomeDecimals;
}

/**
 * @dev Interface for a remote bridge contract that mints a representation token on its chain, and allows
 * for burning that token to redeem the backing asset on the home chain, or bridging to other remotes.
 */
interface ITokenRemote is ITokenBridge {
    /**
     * @notice Sends a Teleporter message to register the TokenRemote instance with its configured home.
     * @param feeInfo The optional fee asset and amount for the Teleporter message, for relayer incentivization.
     */
    function registerWithHome(TeleporterFeeInfo calldata feeInfo) external;
}
