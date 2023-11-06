// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITeleporterReceiver} from "../../Teleporter/ITeleporterReceiver.sol";
import {TeleporterUpgradeable} from "../../Teleporter/upgrades/TeleporterUpgradeable.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/**
 * Contract for receiving latest block hashes from another chain.
 */
contract BlockHashReceiver is
    ITeleporterReceiver,
    TeleporterUpgradeable,
    Ownable
{
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
        address teleporterRegistryAddress,
        bytes32 publisherChainID,
        address publisherContractAddress
    ) TeleporterUpgradeable(teleporterRegistryAddress) {
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
    ) external onlyAllowedTeleporter {
        require(
            originChainID == sourceChainID,
            "BlockHashReceiver: invalid source chain ID"
        );
        require(
            originSenderAddress == sourcePublisherContractAddress,
            "BlockHashReceiver: invalid source chain publisher"
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
     * @dev See {TeleporterUpgradeable-updateMinTeleporterVersion}
     *
     * Updates the minimum Teleporter version allowed for receiving on this contract
     * to the latest version registered in the {TeleporterRegistry}.
     * Restricted to only owners of the contract.
     * Emits a {MinTeleporterVersionUpdated} event.
     */
    function updateMinTeleporterVersion() external override onlyOwner {
        uint256 oldMinTeleporterVersion = minTeleporterVersion;
        minTeleporterVersion = teleporterRegistry.getLatestVersion();
        emit MinTeleporterVersionUpdated(
            oldMinTeleporterVersion,
            minTeleporterVersion
        );
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
