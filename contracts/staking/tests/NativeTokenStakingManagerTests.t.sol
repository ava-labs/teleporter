// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {PoSValidatorManagerTest} from "./PoSValidatorManagerTests.t.sol";
import {NativeTokenStakingManager} from "../NativeTokenStakingManager.sol";
import {ValidatorManagerSettings} from "../interfaces/IValidatorManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";

// TODO: Remove this once all unit tests implemented
// solhint-disable no-empty-blocks
contract NativeTokenStakingManagerTest is PoSValidatorManagerTest {
    NativeTokenStakingManager public app;

    function setUp() public virtual {
        // Construct the object under test
        app = new NativeTokenStakingManager(ICMInitializable.Allowed);
        app.initialize(
            ValidatorManagerSettings({
                pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                subnetID: DEFAULT_SUBNET_ID,
                maximumHourlyChurn: DEFAULT_MAXIMUM_HOURLY_CHURN
            }),
            DEFAULT_MINIMUM_STAKE,
            DEFAULT_MAXIMUM_STAKE,
            DEFAULT_MINIMUM_STAKE_DURATION,
            IRewardCalculator(address(0))
        );
        validatorManager = app;
        posValidatorManager = app;
    }

    // Helpers
    function _initializeValidatorRegistration(
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory signature,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        return app.initializeValidatorRegistration{value: app.weightToValue(weight)}(
            nodeID, registrationExpiry, signature
        );
    }

    function _beforeSend(uint64 weight) internal override {
        // Native tokens no need pre approve
    }
}
// TODO: Remove this once all unit tests implemented
// solhint-enable no-empty-blocks
