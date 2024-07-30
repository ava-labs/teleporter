// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "forge-std/Test.sol";
import {StakingManager} from "../StakingManager.sol";
import {IStakingManager} from "../interfaces/IStakingManager.sol";

abstract contract StakingManagerTest is Test {
    IStakingManager public stakingManager;

    event ValidationPeriodCreated(
        bytes32 indexed validationID,
        bytes32 indexed nodeID,
        bytes32 indexed registerValidationMessageID,
        uint256 weight,
        uint64 registrationExpiry
    );

    event ValidationPeriodRegistered(
        bytes32 indexed validationID, uint256 stakeAmount, uint256 timestamp
    );

    event ValidatorRemovalInitialized(
        bytes32 indexed validationID,
        bytes32 indexed setWeightMessageID,
        uint256 stakeAmount,
        uint256 endTime,
        uint64 uptime
    );

    event ValidationPeriodEnded(bytes32 indexed validationID);

    function setup() public virtual {}

    function testProvideInitialStake() public {}

    function testResendRegisterValidatorMessage() public {}

    function testCompleteValidatorRegistration() public {}

    function testInitializeEndValidation() public {}

    function testResendEndValidation() public {}

    function testCompleteEndValidation() public {}

    function testValueToWeight() public {}

    function testWeightToValue() public {}
}
