// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {console2} from "forge-std/console2.sol";

contract ComputeSlotTest is Test {
    function testComputeStorageSlot() public {
        bytes32 slot = keccak256(
            abi.encode(uint256(keccak256("avalanche-ictt.storage.SendReentrancyGuard")) - 1)
        ) & ~bytes32(uint256(0xff));
        console2.logBytes32(slot);
    }
}
