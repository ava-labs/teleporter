// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {INativeTokenBridge} from "./INativeTokenBridge.sol";
import {ITeleporterTokenSource} from "./ITeleporterTokenSource.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Interface for a "home" or "source" native token bridge contract that locks
 * tokens on its chain to be bridged to supported destination bridge contracts on other chains.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
interface INativeTokenSource is INativeTokenBridge, ITeleporterTokenSource {
    /**
     * @notice Adds collateral to the bridge contract.
     * @param destinationBlockchainID the destination blockchain ID of the bridge to add collateral for.
     * @param destinationBridgeAddress the address on the bridge to add collateral for on the remote chain.
     */
    function addCollateral(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress
    ) external payable;
}
