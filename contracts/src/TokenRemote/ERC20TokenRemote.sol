// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ERC20TokenRemoteUpgradeable} from "./ERC20TokenRemoteUpgradeable.sol";
import {TokenRemoteSettings} from "./interfaces/ITokenRemote.sol";
import {ICTTInitializable} from "../utils/ICTTInitializable.sol";

contract ERC20TokenRemote is ERC20TokenRemoteUpgradeable {
    constructor(
        TokenRemoteSettings memory settings,
        string memory tokenName,
        string memory tokenSymbol,
        uint8 tokenDecimals
    ) ERC20TokenRemoteUpgradeable(ICTTInitializable.Allowed) {
        initialize(settings, tokenName, tokenSymbol, tokenDecimals);
    }
}
