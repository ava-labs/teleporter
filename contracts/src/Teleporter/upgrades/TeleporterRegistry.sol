// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITeleporterMessenger} from "../ITeleporterMessenger.sol";
import {IWarpMessenger, WarpMessage} from "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";

/**
 * @dev Registry entry that represents a mapping between protocolAddress and version.
 */
struct ProtocolRegistryEntry {
    uint256 version;
    address protocolAddress;
}

/**
 * @dev TeleporterRegistry contract provides an upgrade mechanism for {ITeleporterMessenger} contracts
 * through Warp off-chain messages
 */
contract TeleporterRegistry {
    // Address that the off-chain Warp message sets as the "source" address.
    // The address is not owned by any EOA or smart contract account, so it
    // cannot possibly be the source address of any other Warp message emitted by the VM.
    address public constant VALIDATORS_SOURCE_ADDRESS = address(0);

    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    bytes32 public immutable blockchainID;

    // The latest protocol version. 0 means no protocol version has been added, and isn't a valid version.
    uint256 public latestVersion;

    // Mappings that keep track of the protocol version and corresponding contract address.
    mapping(uint256 version => address protocolAddress)
        private _versionToAddress;
    mapping(address protocolAddress => uint256 version)
        private _addressToVersion;

    /**
     * @dev Emitted when a new protocol version is added to the registry.
     */
    event AddProtocolVersion(
        uint256 indexed version,
        address indexed protocolAddress
    );

    /**
     * @dev Emitted when the latest version is updated.
     */
    event LatestVersionUpdated(
        uint256 indexed oldVersion,
        uint256 indexed newVersion
    );

    /**
     * @dev Initializes the contract by setting `blockchainID` and `latestVersion`.
     * Also adds the initial protocol versions to the registry.
     */
    constructor(ProtocolRegistryEntry[] memory initialEntries) {
        blockchainID = WARP_MESSENGER.getBlockchainID();
        latestVersion = 0;

        for (uint256 i = 0; i < initialEntries.length; i++) {
            _addToRegistry(initialEntries[i]);
        }
    }

    /**
     * @dev Gets and verifies a Warp off-chain message, and adds the new version to procotol address mapping,
     * specified in the Warp off-chain message, to the registry.
     * If a version is greater than the current latest version, it will be set as the latest version.
     * If a version is less than the current latest version, it is added to the registry, but
     * doesn't change the latest version.
     *
     * Emits a {AddProtocolVersion} event when successful.
     * Emits a {LatestVersionUpdated} event when a new protocol version great than the current latest version is added.
     * Requirements:
     *
     * - a valid Warp off-chain message must be provided.
     * - source chain ID must be the same as the blockchain ID of the registry.
     * - origin sender address must be the same as the `VALIDATORS_SOURCE_ADDRESS`.
     * - destination address must be the same as the address of this registry.
     */
    function addProtocolVersion(uint32 messageIndex) external {
        // Get and validate for a Warp off-chain message.
        (WarpMessage memory message, bool success) = WARP_MESSENGER
            .getVerifiedWarpMessage(messageIndex);
        require(success, "TeleporterRegistry: invalid warp message");
        require(
            message.sourceChainID == blockchainID,
            "TeleporterRegistry: invalid source chain ID"
        );
        // Check that the message is sent through a Warp off-chain message.
        require(
            message.originSenderAddress == VALIDATORS_SOURCE_ADDRESS,
            "TeleporterRegistry: invalid origin sender address"
        );

        (ProtocolRegistryEntry memory entry, address destinationAddress) = abi
            .decode(message.payload, (ProtocolRegistryEntry, address));

        // Check that the message is sent to the registry.
        require(
            destinationAddress == address(this),
            "TeleporterRegistry: invalid destination address"
        );

        _addToRegistry(entry);
    }

    /**
     * @dev Gets the latest {ITeleporterMessenger} contract.
     */
    function getLatestTeleporter()
        external
        view
        returns (ITeleporterMessenger)
    {
        return ITeleporterMessenger(getAddressFromVersion(latestVersion));
    }

    /**
     * @dev Gets the {ITeleporterMessenger} contract of the given `version`.
     * Requirements:
     *
     * - `version` must be a valid version registered.
     */
    function getTeleporterFromVersion(
        uint256 version
    ) external view returns (ITeleporterMessenger) {
        return ITeleporterMessenger(getAddressFromVersion(version));
    }

    /**
     * @dev Gets the address of a protocol version.
     * Requirements:
     *
     * - `version` must be a valid version registered.
     */
    function getAddressFromVersion(
        uint256 version
    ) public view returns (address) {
        require(version != 0, "TeleporterRegistry: zero version");
        address protocolAddress = _versionToAddress[version];
        require(
            protocolAddress != address(0),
            "TeleporterRegistry: version not found"
        );
        return protocolAddress;
    }

    /**
     * @dev Gets the version of the given `protocolAddress`.
     * Rrequirements:
     *
     * - `protocolAddress` must be a valid protocol address registered.
     */
    function getVersionFromAddress(
        address protocolAddress
    ) public view returns (uint256) {
        require(
            protocolAddress != address(0),
            "TeleporterRegistry: zero protocol address"
        );
        uint256 version = _addressToVersion[protocolAddress];
        require(version != 0, "TeleporterRegistry: protocol address not found");
        return version;
    }

    /**
     * @dev Adds the new protocol version address to the registry.
     * Updates latest version if the version is greater than the current latest version.
     *
     * Emits a {AddProtocolVersion} event when successful.
     * Emits a {LatestVersionUpdated} event when a new protocol version great than the current latest version is added.
     * Note: `protocolAddress` doesn't have to be a contract address, this is primarily
     * to support the case of registering a new protocol address meant for a security patch
     * before the contract is deployed, and prevent the vulnerabilitiy from being exposed before the registry update.
     * Requirements:
     *
     * - `version` is not zero
     * - `version` is not already registered
     * - `protocolAddress` is not zero address
     */
    function _addToRegistry(ProtocolRegistryEntry memory entry) private {
        require(entry.version != 0, "TeleporterRegistry: zero version");
        // Check that the version has not previously been registered.
        require(
            _versionToAddress[entry.version] == address(0),
            "TeleporterRegistry: version already exists"
        );
        require(
            entry.protocolAddress != address(0),
            "TeleporterRegistry: zero protocol address"
        );

        _versionToAddress[entry.version] = entry.protocolAddress;
        _addressToVersion[entry.protocolAddress] = entry.version;
        emit AddProtocolVersion(entry.version, entry.protocolAddress);

        // Set latest version if the version is greater than the current latest version.
        if (entry.version > latestVersion) {
            uint256 oldLatestVersion = latestVersion;
            latestVersion = entry.version;
            emit LatestVersionUpdated(oldLatestVersion, latestVersion);
        }
    }
}
