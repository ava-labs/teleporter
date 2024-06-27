// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITokenBridge, SendTokensInput, SendAndCallInput} from "../../interfaces/ITokenBridge.sol";

/**
 * @dev Interface for a "home" bridge contract that locks a specific token
 * on its chain to be bridged to supported "remote" bridge contracts on other chains.
 */
interface ITokenHome is ITokenBridge {
    /**
     * @dev Emitted when tokens are added as collateral for a given TokenRemote instance.
     * The event emits a {remaining} value of 0 when the TokenRemote instance is fully collateralized.
     */
    event CollateralAdded(
        bytes32 indexed remoteBlockchainID,
        address indexed remoteBridgeAddress,
        uint256 amount,
        uint256 remaining
    );

    /**
     * @notice Emitted when a new TokenRemote instance is registered with the token bridge.
     */
    event RemoteRegistered(
        bytes32 indexed remoteBlockchainID,
        address indexed remoteBridgeAddress,
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
