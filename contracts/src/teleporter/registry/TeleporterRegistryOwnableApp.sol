// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.20;

import {TeleporterRegistryApp} from "./TeleporterRegistryApp.sol";
import {Ownable} from "@openzeppelin/contracts@5.0.2/access/Ownable.sol";
import {Context} from "@openzeppelin/contracts@5.0.2/utils/Context.sol";
import {ContextUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ContextUpgradeable.sol";

abstract contract TeleporterRegistryOwnableApp is TeleporterRegistryApp, Ownable {
    constructor(
        address teleporterRegistryAddress,
        address initialOwner
    ) TeleporterRegistryApp(teleporterRegistryAddress) Ownable(initialOwner) {
        // solhint-disable-previous-line no-empty-blocks
    }

    function _contextSuffixLength()
        internal
        view
        override (ContextUpgradeable, Context)
        returns (uint256)
    {
        return Context._contextSuffixLength();
    }

    function _msgData()
        internal
        view
        override (ContextUpgradeable, Context)
        returns (bytes calldata)
    {
        return Context._msgData();
    }

    function _msgSender() internal view override (ContextUpgradeable, Context) returns (address) {
        return Context._msgSender();
    }
}
