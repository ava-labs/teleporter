// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: BSD-3-Clause

pragma solidity 0.8.18;

import "../WarpProtocolRegistry.sol";
import "./ITeleporterMessenger.sol";

contract TeleporterRegistry is WarpProtocolRegistry {
    mapping(address => uint256) internal _addressToVersion;

    function _addProtocolVersion() internal override {
        super._addProtocolVersion();
        _addressToVersion[
            _getVersionToAddress(_latestVersion)
        ] = _latestVersion;
    }

    function getTeleporter(
        uint256 version
    ) external view returns (ITeleporterMessenger) {
        return ITeleporterMessenger(_getVersionToAddress(version));
    }

    function getLatestTeleporter()
        external
        view
        returns (ITeleporterMessenger)
    {
        return ITeleporterMessenger(_getVersionToAddress(_latestVersion));
    }

    function getAddressToVersion(
        address protocolAddress
    ) external view returns (uint256) {
        return _addressToVersion[protocolAddress];
    }
}
