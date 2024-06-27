// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITokenBridge, SendTokensInput, SendAndCallInput} from "./ITokenBridge.sol";

/**
 * @notice Interface for a Teleporter token bridge that sends ERC20 tokens to another chain.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
interface IERC20TokenBridge is ITokenBridge {
    /**
     * @notice Sends ERC20 tokens to the specified destination.
     * @param input Specifies information for delivery of the tokens
     * @param amount Amount of tokens to send
     */
    function send(SendTokensInput calldata input, uint256 amount) external;

    /**
     * @notice Sends ERC20 tokens to the specified destination to be used in a smart contract interaction.
     * @param input Specifies information for delivery of the tokens
     * @param amount Amount of tokens to send
     */
    function sendAndCall(SendAndCallInput calldata input, uint256 amount) external;
}
