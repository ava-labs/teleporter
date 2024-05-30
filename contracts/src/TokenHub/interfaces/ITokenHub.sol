// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITokenBridge, SendTokensInput, SendAndCallInput} from "../../interfaces/ITokenBridge.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Interface for a "hub" bridge contract that locks a specific token
 * on its chain to be bridged to supported "spoke" bridge contracts on other chains.
 */
interface ITokenHub is ITokenBridge {
    /**
     * @dev Emitted when tokens are added as collateral for a given spoke instance.
     * The event emits a {remaining} value of 0 when the spoke instance is fully collateralized.
     */
    event CollateralAdded(
        bytes32 indexed spokeBlockchainID,
        address indexed spokeBridgeAddress,
        uint256 amount,
        uint256 remaining
    );

    /**
     * @notice Emitted when a new spoke instance is registered with the token bridge.
     */
    event SpokeRegistered(
        bytes32 indexed spokeBlockchainID,
        address indexed spokeBridgeAddress,
        uint256 initialCollateralNeeded,
        uint8 tokenDecimals
    );

    /**
     * @notice Emitted when tokens are routed from a multi-hop send message to another chain.
     */
    event TokensRouted(bytes32 indexed teleporterMessageID, SendTokensInput input, uint256 amount);

    /**
     * @notice Emitted when tokens are routed from a mulit-hop send message,
     * with calldata for a contract recipient, to another chain.
     */
    event TokensAndCallRouted(
        bytes32 indexed teleporterMessageID, SendAndCallInput input, uint256 amount
    );
}
