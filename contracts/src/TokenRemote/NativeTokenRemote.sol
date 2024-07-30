// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {NativeTokenRemoteUpgradeable} from "./NativeTokenRemoteUpgradeable.sol";
import {TokenRemoteSettings} from "./interfaces/ITokenRemote.sol";
import {ICTTInitializable} from "../utils/ICTTInitializable.sol";

contract NativeTokenRemote is NativeTokenRemoteUpgradeable {
    constructor(
        TokenRemoteSettings memory settings,
        string memory nativeAssetSymbol,
        uint256 initialReserveImbalance,
        uint256 burnedFeesReportingRewardPercentage
    ) NativeTokenRemoteUpgradeable(ICTTInitializable.Allowed) {
        initialize(
            settings,
            nativeAssetSymbol,
            initialReserveImbalance,
            burnedFeesReportingRewardPercentage
        );
    }
}
