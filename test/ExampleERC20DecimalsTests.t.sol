// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {ExampleERC20Decimals} from "../src/mocks/ExampleERC20Decimals.sol";

contract ExampleERC20DecimalsTest is Test {
    uint8 public constant MOCK_DECIMALS = 11;
    ExampleERC20Decimals public exampleERC20;

    function setUp() public virtual {
        exampleERC20 = new ExampleERC20Decimals(MOCK_DECIMALS);
    }

    function testDecimals() public {
        assertEq(exampleERC20.decimals(), MOCK_DECIMALS);
    }
}
