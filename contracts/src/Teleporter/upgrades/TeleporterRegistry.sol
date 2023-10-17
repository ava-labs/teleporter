// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "../../WarpProtocolRegistry.sol";
import "../ITeleporterMessenger.sol";

/**
 * @dev TeleporterRegistry contract is a {WarpProtocolRegistry} and provides an upgrade
 * mechanism for {ITeleporterMessenger} contracts.
 */
contract TeleporterRegistry is WarpProtocolRegistry {
    constructor(
        ProtocolRegistryEntry[] memory initialEntries
    ) WarpProtocolRegistry(initialEntries) {}

    /**
     * @dev Gets the {ITeleporterMessenger} contract of the given `version`.
     * Requirements:
     *
     * - `version` must be a valid version, i.e. greater than 0 and not greater than the latest version.
     */
    function getTeleporterFromVersion(
        uint256 version
    ) external view returns (ITeleporterMessenger) {
        return ITeleporterMessenger(_getAddressFromVersion(version));
    }

    /**
     * @dev Gets the latest {ITeleporterMessenger} contract.
     */
    function getLatestTeleporter()
        external
        view
        returns (ITeleporterMessenger)
    {
        return ITeleporterMessenger(_getAddressFromVersion(_latestVersion));
    }
}
