// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "../../Teleporter/ITeleporterMessenger.sol";
import "../../Teleporter/ITeleporterReceiver.sol";
import "../../Teleporter/TeleporterRegistry.sol";

/**
 * Contract for receiving latest block hashes from another chain.
 */
contract BlockHashReceiver is ITeleporterReceiver {
    TeleporterRegistry public immutable teleporterRegistry;
    uint256 internal _minTeleporterVersion;

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

    // Errors
    error Unauthorized();
    error InvalidSourceChainID();
    error InvalidSourceChainPublisher();
    error InvalidTeleporterRegistryAddress();

    constructor(
        address teleporterRegistryAddress,
        bytes32 publisherChainID,
        address publisherContractAddress
    ) {
        if (teleporterRegistryAddress == address(0)) {
            revert InvalidTeleporterRegistryAddress();
        }

        teleporterRegistry = TeleporterRegistry(teleporterRegistryAddress);
        _minTeleporterVersion = teleporterRegistry.getLatestVersion();
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
        if (
            teleporterRegistry.getAddressToVersion(msg.sender) <
            _minTeleporterVersion
        ) {
            revert Unauthorized();
        }

        if (originChainID != sourceChainID) {
            revert InvalidSourceChainID();
        }

        if (originSenderAddress != sourcePublisherContractAddress) {
            revert InvalidSourceChainPublisher();
        }

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

    function updateMinTeleporterVersion() external {
        _minTeleporterVersion = teleporterRegistry.getLatestVersion();
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
