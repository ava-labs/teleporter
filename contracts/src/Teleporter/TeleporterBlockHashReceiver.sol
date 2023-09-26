// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import "./ReentrancyGuards.sol";
import "./ITeleporterBlockHashReceiver.sol";

/**
 * @dev Implementation of the {ITeleporterBlockHashReceiver} interface.
 *
 * This implementation is used to receive block hashes from other chains using the WarpMessenger precompile
 */
contract TeleporterBlockHashReceiver is ITeleporterBlockHashReceiver, ReentrancyGuards {
    WarpMessenger public constant WARP_MESSENGER =
        WarpMessenger(0x0200000000000000000000000000000000000005);

    error InvalidWarpBlockHash();
    error UnexpectedSourceChainID();

    /**
     * @dev See {ITeleporterBlockHashReceiver-receiveCrossChainMessage}
     *
     * Emits a {ReceiveBlockHash} event when a block hash is verified..
     */
    function receiveCrossChainMessage(
        uint32 messageIndex,
        bytes32 sourceChainID
    ) external receiverNonReentrant {
        // Verify and parse the block hash payload included in the transaction access list
        // using the warp message precompile.
        (WarpBlockHash memory warpBlockHash, bool success) = WARP_MESSENGER
            .getVerifiedWarpBlockHash(messageIndex);

        if (!success) {
            revert InvalidWarpBlockHash();
        }
        if (warpBlockHash.sourceChainID != sourceChainID) {
            revert UnexpectedSourceChainID();
        }

        emit ReceiveBlockHash(
            warpBlockHash.sourceChainID,
            warpBlockHash.blockHash
        );
    }
}