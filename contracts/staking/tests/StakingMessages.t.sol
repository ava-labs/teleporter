// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {StakingMessages} from "../StakingMessages.sol";
import {StakingMessages as Original} from "./ManualStakingMessages.sol";

contract StakingMessagesTest is Test {
    function testSubnetValidatorWeightPackRoundTrip(
        bytes32 validationID,
        uint64 nonce,
        uint64 weight
    ) public view {
        uint256 startGas;
        uint256 endGas;

        startGas = gasleft();
        bytes memory original =
            Original.packSetSubnetValidatorWeightMessage(validationID, nonce, weight);
        endGas = gasleft();
        uint256 originalGas = startGas - endGas;

        startGas = gasleft();
        bytes memory packed =
            StakingMessages.packSetSubnetValidatorWeightMessage(validationID, nonce, weight);
        endGas = gasleft();
        // Doing this after the original will actually penalise it because
        // memory expansion is quadratic, so the gas saving is higher than the
        // test suggests.
        uint256 gas = startGas - endGas;

        assertEq(original, packed, "same data");
        assertGe(originalGas - gas, 19_500, "gas saving");
        assertLe(gas, 250, "gas cost");

        (bytes32 gotValidationID, uint64 gotNonce, uint64 gotWeight) =
            StakingMessages.unpackSetSubnetValidatorWeightMessage(packed);

        assertEq(validationID, gotValidationID, "validation ID");
        assertEq(nonce, gotNonce, "nonce");
        assertEq(weight, gotWeight, "weight");
    }

    function testUnpackSubnetValidatorWeightGas(
        bytes32 validationID,
        uint64 nonce,
        uint64 weight
    ) public view {
        bytes memory buffer =
            StakingMessages.packSetSubnetValidatorWeightMessage(validationID, nonce, weight);

        uint256 startGas;
        uint256 endGas;

        startGas = gasleft();
        StakingMessages.unpackSetSubnetValidatorWeightMessage(buffer);
        endGas = gasleft();
        uint256 gas = startGas - endGas;

        startGas = gasleft();
        Original.unpackSetSubnetValidatorWeightMessage(buffer);
        endGas = gasleft();
        uint256 originalGas = startGas - endGas;

        assertGe(originalGas - gas, 18_000, "gas saving");
        assertLe(gas, 350, "gas cost");
    }
}
