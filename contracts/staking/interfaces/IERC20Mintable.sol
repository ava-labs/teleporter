// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";

interface IERC20Mintable is IERC20 {
    /**
     * @notice Mint tokens to the specified address.
     * @param account The address to mint tokens to.
     * @param amount How many tokens to mint.
     * @dev This function should have appropriate user controls to ensure that only the staking contract can mint.
     */
    function mint(address account, uint256 amount) external;
}
