// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {INativeTokenBridge} from "../../interfaces/INativeTokenBridge.sol";
import {ITokenHub} from "./ITokenHub.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Interface for a native token "hub" contract that locks the native token
 * on its chain to be bridged to supported spoke bridge contracts on other chains.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
interface INativeTokenHub is INativeTokenBridge, ITokenHub {
    /**
     * @notice Adds collateral to the hub bridge contract for the specified spoke instance. If more value is provided
     * than the amount of collateral needed, the excess amount is returned to the caller.
     * @param spokeBlockchainID The blockchain ID of the spoke bridge contract to add collateral for.
     * @param spokeBridgeAddress The address of the spoke bridge contract to add collateral for on
     * the {spokeBlockchainID}.
     */
    function addCollateral(
        bytes32 spokeBlockchainID,
        address spokeBridgeAddress
    ) external payable;
}
