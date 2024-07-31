// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "forge-std/Test.sol";
import {StakingManagerTest} from "./StakingManagerTests.t.sol";
import {NativeTokenStakingManager} from "../NativeTokenStakingManager.sol";
import {StakingManagerSettings, InitialStakerInfo} from "../interfaces/IStakingManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";
import {IWarpMessenger} from "../StakingManager.sol";
import {StakingMessages} from "../StakingMessages.sol";

contract NativeTokenStakingManagerTest is StakingManagerTest {
    NativeTokenStakingManager public app;

    function setUp() public override {
        // Construct the object under test
        app = new NativeTokenStakingManager(
            StakingManagerSettings({
                pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                subnetID: DEFAULT_SUBNET_ID,
                minimumStakeAmount: 20,
                maximumStakeAmount: 1e10,
                minimumStakeDuration: 24 hours,
                maximumHourlyChurn: 0,
                initialStakers: new InitialStakerInfo[](0),
                rewardCalculator: IRewardCalculator(address(0))
            })
        );
        stakingManager = app;

        // Setup the test
        StakingManagerTest.setUp();
    }

    function testInitializeValidatorRegistrationSuccess() public {
        (bytes32 validationID, ) = StakingMessages.packValidationInfo(
            StakingMessages.ValidationInfo({
                nodeID: DEFAULT_NODE_ID,
                subnetID: DEFAULT_SUBNET_ID,
                weight: 1e6,
                registrationExpiry: DEFAULT_EXPIRY,
                signature: DEFAULT_ED25519_SIGNATURE
            })
        );

        vm.warp(DEFAULT_EXPIRY-1);
        
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(bytes32(0))
        );
        // TODO: construct the expected message in the same way as the contract
        // vm.expectCall(
        //     WARP_PRECOMPILE_ADDRESS,
        //     abi.encodeCall(IWarpMessenger.sendWarpMessage, (abi.encode(expectedMessage)))
        // );
        vm.expectEmit(true, true, true, true, address(app));
        emit ValidationPeriodCreated(validationID, DEFAULT_NODE_ID, bytes32(0), 1e6, DEFAULT_EXPIRY);

        app.initializeValidatorRegistration{value: 1e18}(
            DEFAULT_NODE_ID,
            DEFAULT_EXPIRY,
            DEFAULT_ED25519_SIGNATURE
        );
    }

    function testInitializeValidatorRegistrationExcessiveChurn() public {}
    function testInitializeValidatorRegistrationInsufficientStake() public {}
    function testInitializeValidatorRegistrationExcessiveStake() public {}
    function testInitializeValidatorRegistrationInsufficientDuration() public {}
}
