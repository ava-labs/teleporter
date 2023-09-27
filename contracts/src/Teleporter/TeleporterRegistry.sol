// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "../WarpProtocolRegistry.sol";
import "./ITeleporterMessenger.sol";

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
    function getTeleporter(
        uint256 version
    ) external view returns (ITeleporterMessenger) {
        return ITeleporterMessenger(_getVersionToAddress(version));
    }

    /**
     * @dev Gets the latest {ITeleporterMessenger} contract.
     */
    function getLatestTeleporter()
        external
        view
        returns (ITeleporterMessenger)
    {
        return ITeleporterMessenger(_getVersionToAddress(_latestVersion));
    }

    /**
     * @dev Gets the version of the given `protocolAddress`.
     * If `protocolAddress` is not a valid protocol address, returns 0.
     */
    function getAddressToVersion(
        address protocolAddress
    ) external view returns (uint256) {
        return _addressToVersion[protocolAddress];
    }

    /**
     * @dev See {WarpProtocolRegistry-addProtocolVersion}
     *
     * Adds the new protocol version address to the `_addressToVersion` mapping
     * so that the registry can be queried for the version of a given protocol address.
     */
    function _addProtocolVersion(uint32 messageIndex) internal override {
        super._addProtocolVersion(messageIndex);
        _addressToVersion[
            _getVersionToAddress(_latestVersion)
        ] = _latestVersion;
    }
}
