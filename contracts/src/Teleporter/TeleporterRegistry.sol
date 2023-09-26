// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: BSD-3-Clause

pragma solidity 0.8.18;

import "../WarpProtocolRegistry.sol";
import "./ITeleporterMessenger.sol";

contract TeleporterRegistry is WarpProtocolRegistry {
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
}
