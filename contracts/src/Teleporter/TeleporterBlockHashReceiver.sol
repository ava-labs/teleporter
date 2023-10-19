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

    /**
     * @dev See {ITeleporterBlockHashReceiver-receiveBlockHash}
     *
     * Emits a {ReceiveBlockHash} event when a block hash is verified..
     */
    function receiveBlockHash(
        uint32 messageIndex
    ) external receiverNonReentrant {
        // Verify and parse the block hash payload included in the transaction access list
        // using the warp message precompile.
        (WarpBlockHash memory warpBlockHash, bool success) = WARP_MESSENGER
            .getVerifiedWarpBlockHash(messageIndex);

        require(success, "TeleporterBlockHashReceiver: invalid warp message");

        emit ReceiveBlockHash(
            warpBlockHash.sourceChainID,
            warpBlockHash.blockHash
        );
    }
}