// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: BSD-3-Clause

pragma solidity 0.8.18;

import "../WarpProtocolRegistry.sol";
import "./ITeleporterMessenger.sol";

contract TeleporterRegistry is WarpProtocolRegistry {
    mapping(address => uint256[] versions) private _allowedVersions;
    mapping(address => mapping(address => bool)) private _allowedSenders;
    mapping(address => bool) private _allowedLatestVersion;

    function getTeleporter(
        uint256 version
    ) external view returns (ITeleporterMessenger) {
        return ITeleporterMessenger(_getVersionAddress(version));
    }

    function getLatestVersion() external view returns (uint256) {
        return _latestVersion;
    }

    function set_getVersionAddresst256[] versions) external {
        // Remove all previous versions that were allowed
        for (uint256 i = 0; i < _allowedVersions[msg.sender].length; i++) {
            delete _allowedSenders[msg.sender][
                _getVersionAddress(_allowedVersions[msg.sender][i])
            ];
        }

        // Add all new versions that are allowed
        for (uint256 i = 0; i < versions.length; i++) {
            require(
                versions[i] < _latestVersion_getVersionAddress
                "TeleporterRegistry: invalid version"
            );
            require(
                _allowedSenders[msg.sender][_getVersionAddress(versions[i])] ==
                    false,
                _getVersionAddress: duplicate version"
            );

            _allowedSenders[msg.sender][
                _getVersionAddress(versions[i])
            ] = true;
        }
        _allowedVersions[msg.sender] = versions;
    }

    function allowLatestVersion(bool allow) external {
        _allowedLatestVersion[msg.sender] = allow;
    }

    function isAllowedTeleporterSender(
        address sender_getVersionAddress
    ) external view returns (bool) {
        if (
            _allowedLatestVersion[msg.sender] &&
            sender == _getVersionAddress(_latestVersion)
        ) {
            return true;
        }

        // If no allowed versions, then we always return false
        // require(
        //     _allowedVersions[msg.sender].length > 0,
        //     "TeleporterRegistry: no allowed versions"
        // );
        require(sender != address(0), "TeleporterRegistry: invalid sender");

        return _allowedSenders[msg.sender][sender];
    }

    function getLatestTeleporter()
        external_getVersionAddress
        view
        returns (ITeleporterMessenger)
    {
        return ITeleporterMessenger(_getVersionAddress(_latestVersion));
    }
}
