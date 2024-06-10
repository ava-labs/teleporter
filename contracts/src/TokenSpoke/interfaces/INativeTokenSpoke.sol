// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {INativeTokenBridge} from "../../interfaces/INativeTokenBridge.sol";
import {ITokenSpoke} from "./ITokenSpoke.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Interface that describes functionalities for a contract that can mint native tokens when
 * paired with a {ITokenHub} contract that will lock tokens on another chain.
 */
interface INativeTokenSpoke is ITokenSpoke, INativeTokenBridge {
    /**
     * @notice Emitted when native tokens are burned for bridge transfers.
     */
    event BurnedForBridge(uint256 amount);

    /**
     * @notice Emitted when reporting burned tx fees to hub chain.
     */
    event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned);

    /**
     * @notice Reports the increase in total burned transaction fees on this chain to its corresponding hub instance.
     * @param requiredGasLimit The gas limit required to report the burned transaction fees to the hub contract.
     */
    function reportBurnedTxFees(uint256 requiredGasLimit) external;

    /**
     * @notice Returns a best-estimate (upper bound) of the supply of the native asset
     * in circulation on this chain. Does not account for other native asset burn mechanisms,
     * which can result in the value returned being greater than true circulating supply.
     */
    function totalNativeAssetSupply() external view returns (uint256);
}
