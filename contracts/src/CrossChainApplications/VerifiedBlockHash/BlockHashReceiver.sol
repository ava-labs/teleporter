// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: BSD-3-Clause

pragma solidity 0.8.18;

import "../../Teleporter/ITeleporterMessenger.sol";
import "../../Teleporter/ITeleporterReceiver.sol";

/**
 * Contract for receiving latest block hashes from another chain.
 */
contract BlockHashReceiver is ITeleporterReceiver {
    ITeleporterMessenger public immutable teleporterMessenger;

    // Source chain information
    bytes32 public immutable sourceChainID;
    address public immutable sourcePublisherContractAddress;

    // Latest received block information
    uint256 public latestBlockHeight;
    bytes32 public latestBlockHash;

    /**
     * @dev Emitted when a new block hash is received from a given origin chain ID.
     */
    event ReceiveBlockHash(
        bytes32 indexed originChainID,
        address indexed originSenderAddress,
        uint256 indexed blockHeight,
        bytes32 blockHash
    );

    constructor(
        address teleporterMessengerAddress,
        bytes32 publisherChainID,
        address publisherContractAddress
    ) {
        teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);
        sourceChainID = publisherChainID;
        sourcePublisherContractAddress = publisherContractAddress;
    }

    /**
     * @dev See {ITeleporterReceiver-receiveTeleporterMessage}.
     *
     * Receives the latest block hash from another chain
     *
     * Requirements:
     *
     * - Sender must be the Teleporter contract.
     * - Origin sender address must be the source publisher contract address that initiated the BlockHashReceiver.
     */

    function receiveTeleporterMessage(
        bytes32 originChainID,
        address originSenderAddress,
        bytes calldata message
    ) external {
        require(
            msg.sender == address(teleporterMessenger),
            "Unauthorized caller."
        );
        require(originChainID == sourceChainID, "Invalid source chain ID.");
        require(
            originSenderAddress == sourcePublisherContractAddress,
            "Unauthorized source chain publisher"
        );

        (uint256 blockHeight, bytes32 blockHash) = abi.decode(
            message,
            (uint256, bytes32)
        );

        if (blockHeight > latestBlockHeight) {
            latestBlockHeight = blockHeight;
            latestBlockHash = blockHash;
            emit ReceiveBlockHash(
                originChainID,
                originSenderAddress,
                blockHeight,
                blockHash
            );
        }
    }

    /**
     * @dev Returns the latest block information.
     */
    function getLatestBlockInfo()
        public
        view
        returns (uint256 height, bytes32 hash)
    {
        return (latestBlockHeight, latestBlockHash);
    }
}
