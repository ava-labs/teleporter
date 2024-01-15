// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Contract that publishes the latest block hash of current chain to another chain.
 */
contract BlockHashPublisher {
    // The gas limit required to receive a block hash on the destination chain.
    uint256 public constant RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT = 1.5e5;

    // The Teleporter registry contract manages different Teleporter contract versions.
    TeleporterRegistry public immutable teleporterRegistry;

    /**
     * @dev Emitted when a block hash is submitted to be published to another chain.
     */
    event PublishBlockHash(
        bytes32 indexed destinationBlockchainID,
        address indexed destinationAddress,
        uint256 indexed blockHeight,
        bytes32 blockHash
    );

    constructor(address teleporterRegistryAddress) {
        require(
            teleporterRegistryAddress != address(0),
            "BlockHashPublisher: zero teleporter registry address"
        );

        teleporterRegistry = TeleporterRegistry(teleporterRegistryAddress);
    }

    /**
     * @dev Publishes the latest block hash to another chain.
     * @return The message of the of the message sent to publish the hash.
     */
    function publishLatestBlockHash(
        bytes32 destinationBlockchainID,
        address destinationAddress
    ) external returns (bytes32) {
        // Get the latest block info. Note it must the previous block
        // because the current block hash is not available during execution.
        uint256 blockHeight = block.number - 1;
        bytes32 blockHash = blockhash(blockHeight);

        // ABI encode the function arguments to be called on the destination.
        // The sourceBlockchainID and originSenderAddress arguments of the target function are provided by Warp/Teleporter.
        bytes memory messageData = abi.encode(blockHeight, blockHash);

        emit PublishBlockHash(destinationBlockchainID, destinationAddress, blockHeight, blockHash);
        return teleporterRegistry.getLatestTeleporter().sendCrossChainMessage(
            TeleporterMessageInput({
                destinationBlockchainID: destinationBlockchainID,
                destinationAddress: destinationAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
                requiredGasLimit: RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT,
                allowedRelayerAddresses: new address[](0),
                message: messageData
            })
        );
    }
}
