// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {WarpMessage, IWarpMessenger} from "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";

/**
 * @dev Registry entry that represents a mapping between protocolAddress and version.
 */
struct ProtocolRegistryEntry {
    uint256 version;
    address protocolAddress;
}

/**
 * @dev Implementation of an abstract `WarpProtocolRegistry` contract.
 *
 * This implementation is a contract that can be used as a base contract for protocols that are
 * built on top of Warp. It allows the protocol to be upgraded through a Warp out-of-band message.
 */
abstract contract WarpProtocolRegistry {
    // Address that the out-of-band Warp message sets as the "source" address.
    // The address is not owned by any EOA or smart contract account, so it
    // cannot possibly be the source address of any other Warp message emitted by the VM.
    address public constant VALIDATORS_SOURCE_ADDRESS = address(0);

    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    bytes32 internal immutable _blockchainID;

    // The latest protocol version. 0 means no protocol version has been added, and isn't a valid version.
    uint256 internal _latestVersion;

    // Mappings that keep track of the protocol version and corresponding contract address.
    mapping(uint256 version => address protocolAddress)
        internal _versionToAddress;
    mapping(address protocolAddress => uint256 version)
        internal _addressToVersion;

    /**
     * @dev Emitted when a new protocol version is added to the registry.
     */
    event AddProtocolVersion(
        uint256 indexed version,
        address indexed protocolAddress
    );

    /**
     * @dev Initializes the contract by setting `_blockchainID` and `_latestVersion`.
     * Also adds the initial protocol versions to the registry.
     */
    constructor(ProtocolRegistryEntry[] memory initialEntries) {
        _blockchainID = WARP_MESSENGER.getBlockchainID();
        _latestVersion = 0;

        for (uint256 i = 0; i < initialEntries.length; i++) {
            _addToRegistry(initialEntries[i]);
        }
    }

    /**
     * @dev Gets and verifies a warp out-of-band message, and adds the new protocol version
     * address to the registry.
     * If a version is greater than the current latest version, it will be set as the latest version.
     * If a version is less than the current latest version, it is added to the registry, but
     * doesn't change the latest version.
     *
     * Emits a {AddProtocolVersion} event when successful.
     * Requirements:
     *
     * - a valid Warp out-of-band message must be provided.
     * - source chain ID must be the same as the blockchain ID of the registry.
     * - origin sender address must be the same as the `VALIDATORS_SOURCE_ADDRESS`.
     * - destination chain ID must be the same as the blockchain ID of the registry.
     * - destination address must be the same as the address of the registry.
     * - version must not be zero.
     * - version must not already be registered.
     * - protocol address must not be zero address.
     */
    function addProtocolVersion(uint32 messageIndex) external virtual {
        // Get and validate for a warp out-of-band message.
        (WarpMessage memory message, bool success) = WARP_MESSENGER
            .getVerifiedWarpMessage(messageIndex);
        require(success, "WarpProtocolRegistry: invalid warp message");
        require(
            message.sourceChainID == _blockchainID,
            "WarpProtocolRegistry: invalid source chain ID"
        );
        // Check that the message is sent through a warp out of band message.
        require(
            message.originSenderAddress == VALIDATORS_SOURCE_ADDRESS,
            "WarpProtocolRegistry: invalid origin sender address"
        );

        (ProtocolRegistryEntry memory entry, address destinationAddress) = abi
            .decode(message.payload, (ProtocolRegistryEntry, address));

        // Check that the message is sent to the registry.
        require(
            destinationAddress == address(this),
            "WarpProtocolRegistry: invalid destination address"
        );

        _addToRegistry(entry);
    }

    /**
     * @dev Gets the address of a protocol version.
     * Requirements:
     *
     * - the version must be a valid version.
     */
    function getAddressFromVersion(
        uint256 version
    ) external view virtual returns (address) {
        return _getAddressFromVersion(version);
    }

    /**
     * @dev Gets the version of the given `protocolAddress`.
     * If `protocolAddress` is not a registered protocol address, returns 0, which is an invalid version.
     */
    function getVersionFromAddress(
        address protocolAddress
    ) external view virtual returns (uint256) {
        return _addressToVersion[protocolAddress];
    }

    /**
     * @dev Gets the latest protocol version.
     * If the registry has no versions, we return 0, which is an invalid version.
     */
    function getLatestVersion() external view virtual returns (uint256) {
        return _latestVersion;
    }

    /**
     * @dev Adds the new protocol version address to the registry.
     * Updates latest version if the version is greater than the current latest version.
     *
     * Emits a {AddProtocolVersion} event when successful.
     * Note: `protocolAddress` doesn't have to be a contract address, this is primarily
     * to support the case we want to register a new protocol address meant for a security patch
     * before the contract is deployed, to prevent the vulnerabilitiy from being exposed before registry update.
     * Requirements:
     *
     * - `version` is not zero
     * - `version` is not already registered
     * - `protocolAddress` is not zero address
     */
    function _addToRegistry(
        ProtocolRegistryEntry memory entry
    ) internal virtual {
        require(entry.version != 0, "WarpProtocolRegistry: zero version");
        // Check that the version has not previously been registered.
        require(
            _versionToAddress[entry.version] == address(0),
            "WarpProtocolRegistry: version already exists"
        );
        require(
            entry.protocolAddress != address(0),
            "WarpProtocolRegistry: zero protocol address"
        );

        _versionToAddress[entry.version] = entry.protocolAddress;
        _addressToVersion[entry.protocolAddress] = entry.version;
        emit AddProtocolVersion(entry.version, entry.protocolAddress);

        // Set latest version if the version is greater than the current latest version.
        if (entry.version > _latestVersion) {
            _latestVersion = entry.version;
        }
    }

    /**
     * @dev Gets the corresponding address of a protocol version.
     * Requirements:
     *
     * - `version` must be a valid version, i.e. greater than 0 and not greater than the latest version.
     */
    function _getAddressFromVersion(
        uint256 version
    ) internal view virtual returns (address) {
        require(version != 0, "WarpProtocolRegistry: zero version");
        require(
            version <= _latestVersion,
            "WarpProtocolRegistry: invalid version"
        );
        return _versionToAddress[version];
    }
}
