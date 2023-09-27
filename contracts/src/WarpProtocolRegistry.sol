// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import "@openzeppelin/contracts/utils/Address.sol";

/**
 * @dev Implementation of an abstract `WarpProtocolRegistry` contract.
 *
 * This implementation is a contract that can be used as a base contract for protocols that are
 * build on top of Warp. It allows the protocol to be upgraded through a warp out of band message.
 */
abstract contract WarpProtocolRegistry {
    /**
     * @dev Emitted when a new protocol version is added to the registry.
     */
    event AddProtocolVersion(
        uint256 indexed version,
        address indexed protocolAddress
    );

    // Address that the out-of-band warp message sets as the "source" address.
    // The address is obviously not owned by any EOA or smart contract account, so it
    // can not possibly be the source address of any other warp message emitted by the VM.
    address public constant VALIDATORS_SOURCE_ADDRESS = address(0);

    WarpMessenger public constant WARP_MESSENGER =
        WarpMessenger(0x0200000000000000000000000000000000000005);

    bytes32 internal immutable _chainID;

    uint256 internal _latestVersion;

    mapping(uint256 => address) internal _versionToAddress;

    // Errors
    error InvalidWarpMessage();
    error InvalidOriginChainID();
    error InvalidOriginSenderAddress();
    error InvalidDestinationChainID();
    error InvalidDestinationAddress();
    error InvalidProtocolAddress();
    error InvalidProtocolVersion();

    /**
     * @dev Initializes the contract by setting a `chainID` and `latestVersion`.
     */
    constructor() {
        _latestVersion = 0;
        _chainID = WARP_MESSENGER.getBlockchainID();
    }

    /**
     * @dev Gets and verifies a warp out of band message, and adds the new protocol version
     * addres to the registry.
     *
     * Emits a {AddProtocolVersion} event when successful.
     * Requirements:
     *
     * - a valid warp out of band message must be provided.
     * - the version must be the increment of the latest version.
     * - the protocol address must be a contract address.
     */
    function addProtocolVersion() external {
        _addProtocolVersion();
    }

    /**
     * @dev Gets the address of a protocol version.
     * Requirements:
     *
     * - the version must be a valid version.
     */
    function getVersionToAddress(
        uint256 version
    ) external view returns (address) {
        return _getVersionToAddress(version);
    }

    /**
     * @dev Gets the latest protocol version.
     */
    function getLatestVersion() external view returns (uint256) {
        return _latestVersion;
    }

    function _addProtocolVersion() internal virtual {
        (WarpMessage memory message, bool valid) = WARP_MESSENGER
            .getVerifiedWarpMessage();
        if (!valid) {
            revert InvalidWarpMessage();
        }
        if (message.originChainID != _chainID) {
            revert InvalidOriginChainID();
        }
        if (message.originSenderAddress != VALIDATORS_SOURCE_ADDRESS) {
            revert InvalidOriginSenderAddress();
        }
        if (message.destinationChainID != _chainID) {
            revert InvalidDestinationChainID();
        }
        if (message.destinationAddress != address(this)) {
            revert InvalidDestinationAddress();
        }

        (uint256 version, address protocolAddress) = abi.decode(
            message.payload,
            (uint256, address)
        );

        if (version != _latestVersion) {
            revert InvalidProtocolVersion();
        }
        if (!Address.isContract(protocolAddress)) {
            revert InvalidProtocolAddress();
        }

        _latestVersion++;
        _versionToAddress[_latestVersion] = protocolAddress;
        emit AddProtocolVersion(_latestVersion, protocolAddress);
    }

    function _getVersionToAddress(
        uint256 version
    ) internal view virtual returns (address) {
        if (!(0 < version && version < _latestVersion)) {
            revert InvalidProtocolVersion();
        }
        return _versionToAddress[version];
    }
}
