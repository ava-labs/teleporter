// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/ERC20.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

interface IWrappedNativeToken is IERC20 {
    function deposit() external payable;

    function withdraw(uint wad) external;
}
