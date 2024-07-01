// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITokenTransferer, SendTokensInput, SendAndCallInput} from "./ITokenTransferer.sol";

/**
 * @notice Interface for an Avalanche interchain token transferer that sends native tokens to another chain.
 *
 * @custom:security-contact https://github.com/ava-labs/avalanche-interchain-token-transfer/blob/main/SECURITY.md
 */
interface INativeTokenTransferer is ITokenTransferer {
    /**
     * @notice Sends native tokens to the specified destination.
     * @param input Specifies information for delivery of the tokens
     */
    function send(SendTokensInput calldata input) external payable;

    /**
     * @notice Sends native tokens to the specified destination to be used in a smart contract interaction.
     * @param input Specifies information for delivery of the tokens to the remote contract and contract to be called
     * on the remote chain.
     */
    function sendAndCall(SendAndCallInput calldata input) external payable;
}
