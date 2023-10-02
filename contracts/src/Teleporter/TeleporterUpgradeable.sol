// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "./TeleporterRegistry.sol";

abstract contract TeleporterUpgradeable {
    TeleporterRegistry public immutable teleporterRegistry;
    uint256 internal _minTeleporterVersion;

    error InvalidTeleporterRegistryAddress();
    error InvalidTeleporterSender();

    constructor(address teleporterRegistryAddress) {
        if (teleporterRegistryAddress == address(0)) {
            revert InvalidTeleporterRegistryAddress();
        }

        teleporterRegistry = TeleporterRegistry(teleporterRegistryAddress);
        _minTeleporterVersion = teleporterRegistry.getLatestVersion();
    }

    modifier onlyAllowedTeleporter() {
        _checkTeleporterMinVersion();
        _;
    }

    function updateMinTeleporterversion() external {
        _minTeleporterVersion = teleporterRegistry.getLatestVersion();
    }

    function _checkTeleporterMinVersion() internal view virtual {
        if (
            teleporterRegistry.getAddressToVersion(msg.sender) <
            _minTeleporterVersion
        ) {
            revert InvalidTeleporterSender();
        }
    }
}
