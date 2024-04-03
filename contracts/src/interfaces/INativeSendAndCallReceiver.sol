// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Interface for a contracts that are called to receive native tokens.
 */
interface INativeSendAndCallReceiver {
    /**
     * @notice Called to receive the amount of the native token
     * @param payload arbitrary data provided by the caller
     */
    function receiveTokens(bytes calldata payload) external payable;
}
