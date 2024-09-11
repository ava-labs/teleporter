// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {PoSValidatorManagerTest} from "./PoSValidatorManagerTests.t.sol";
import {NativeTokenStakingManager} from "../NativeTokenStakingManager.sol";
import {ValidatorManagerSettings} from "../interfaces/IValidatorManager.sol";
import {PoSValidatorManagerSettings} from "../interfaces/IPoSValidatorManager.sol";
import {ExampleRewardCalculator} from "../ExampleRewardCalculator.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";

// TODO: Remove this once all unit tests implemented
// solhint-disable no-empty-blocks
contract NativeTokenStakingManagerTest is PoSValidatorManagerTest {
    NativeTokenStakingManager public app;

    function setUp() public virtual {
        // Construct the object under test
        app = new NativeTokenStakingManager(ICMInitializable.Allowed);
        rewardCalculator = new ExampleRewardCalculator(DEFAULT_REWARD_RATE);
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
                    pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                    subnetID: DEFAULT_SUBNET_ID,
                    maximumHourlyChurn: DEFAULT_MAXIMUM_HOURLY_CHURN
                }),
                minimumStakeAmount: DEFAULT_MINIMUM_STAKE,
                maximumStakeAmount: DEFAULT_MAXIMUM_STAKE,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                rewardCalculator: rewardCalculator
            })
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

    function _initializeDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        uint256 value = app.weightToValue(weight);
        vm.prank(delegatorAddress);
        vm.deal(delegatorAddress, value);
        return app.initializeDelegatorRegistration{value: value}(validationID);
    }

    function _beforeSend(uint64 weight, address spender) internal override {
        // Native tokens no need pre approve
    }

    function _expectStakeUnlock(address account, uint256 amount) internal override {
        // empty calldata implies the receive function will be called
        vm.expectCall(account, amount, "");
    }

    function _expectRewardIssuance(address account, uint256 amount) internal override {
        vm.mockCall(
            address(app.NATIVE_MINTER()),
            abi.encodeCall(INativeMinter.mintNativeCoin, (account, amount)),
            ""
        );
        // empty calldata implies the receive function will be called:
        vm.mockCall({
            callee: account,
            msgValue: amount,
            data: "", // implies receive()
            returnData: ""
        });
        vm.deal(account, account.balance + amount);
    }

    function _getStakeAssetBalance(address account) internal view override returns (uint256) {
        return account.balance;
    }
}
// TODO: Remove this once all unit tests implemented
// solhint-enable no-empty-blocks
