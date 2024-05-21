// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {
    ITeleporterTokenBridge, SendTokensInput, SendAndCallInput
} from "./ITeleporterTokenBridge.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Interface for a Teleporter token bridge that sends native tokens to another chain.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
interface INativeTokenBridge is ITeleporterTokenBridge {
    /**
     * @notice Sends native tokens transferred to this contract to the spoke contract instance.
     * @param input Specifies information for delivery of the tokens
     */
    function send(SendTokensInput calldata input) external payable;

    /**
     * @notice Sends native tokens transferred to this contract to the spoke contract instance.
     * @param input Specifies information for delivery of the tokens to the spoke contract and contract to be called
     * on the spoke chain.
     */
    function sendAndCall(SendAndCallInput calldata input) external payable;
}
