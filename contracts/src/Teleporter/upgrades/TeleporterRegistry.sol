// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "../../WarpProtocolRegistry.sol";
import "../ITeleporterMessenger.sol";

/**
 * @dev TeleporterRegistry contract is a {WarpProtocolRegistry} and provides an upgrade
 * mechanism for teleporter messenger contracts.
 */
contract TeleporterRegistry is WarpProtocolRegistry {
    mapping(address => uint256) internal _addressToVersion;

    constructor(
        uint256[] memory initialVersions,
        address[] memory initialProtocolAddresses
    ) WarpProtocolRegistry(initialVersions, initialProtocolAddresses) {}

    /**
     * @dev Gets the {ITeleporterMessenger} contract of the given `version`.
     */
    function getTeleporterFromVersion(
        uint256 version
    ) external view returns (ITeleporterMessenger) {
        return ITeleporterMessenger(_getAddressToVersion(version));
    }

    /**
     * @dev Gets the latest {ITeleporterMessenger} contract.
     */
    function getLatestTeleporter()
        external
        view
        returns (ITeleporterMessenger)
    {
        return ITeleporterMessenger(_getAddressToVersion(_latestVersion));
    }

    /**
     * @dev Gets the version of the given `protocolAddress`.
     * If `protocolAddress` is not a valid protocol address, returns 0.
     */
    function getVersionFromAddress(
        address protocolAddress
    ) external view returns (uint256) {
        return _addressToVersion[protocolAddress];
    }

    /**
     * @dev See {WarpProtocolRegistry-_addToRegistry}
     *
     * Adds the new protocol version address to the `_addressToVersion` mapping
     * so that the registry can be queried for the version of a given protocol address.
     * Requirements:
     *
     * - `protocolAddress` must not have been previously registered.
     */
    function _addToRegistry(
        uint256 version,
        address protocolAddress
    ) internal override {
        WarpProtocolRegistry._addToRegistry(version, protocolAddress);
        _addressToVersion[protocolAddress] = version;
    }
}
