// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * @dev Interface that describes functionalities for receiving block hashes from other chains.
 */
interface ITeleporterBlockHashReceiver {
    /**
     * @dev Emitted when a block hash is successfully received.
     */
    event ReceiveBlockHash(
        bytes32 indexed originChainID,
        bytes32 indexed blockHash
    );

    /**
     * @dev Receives a block hash from another chain.
     *
     * The message specified by `messageIndex` must be provided at that index in the access list storage slots of the transaction,
     * and is verified in the precompile predicate.
     */
    function receiveBlockHash(
        uint32 messageIndex
    ) external;
}