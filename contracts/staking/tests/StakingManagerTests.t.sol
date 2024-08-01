// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {StakingManager} from "../StakingManager.sol";

abstract contract StakingManagerTest is Test {
    bytes32 public constant P_CHAIN_BLOCKCHAIN_ID =
        bytes32(hex"0000000000000000000000000000000000000000000000000000000000000000");
    bytes32 public constant DEFAULT_SUBNET_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes32 public constant DEFAULT_NODE_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    uint64 public constant DEFAULT_EXPIRY = 1000;
    bytes public constant DEFAULT_ED25519_SIGNATURE = bytes(
        hex"12345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678"
    );
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;

    StakingManager public stakingManager;

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

    function setUp() public virtual {}

    function testValueToWeight() public {
        uint64 w1 = stakingManager.valueToWeight(1e12);
        uint64 w2 = stakingManager.valueToWeight(1e18);
        uint64 w3 = stakingManager.valueToWeight(1e27);

        assertEq(w1, 1);
        assertEq(w2, 1e6);
        assertEq(w3, 1e15);
    }

    function testWeightToValue() public {
        uint256 v1 = stakingManager.weightToValue(1);
        uint256 v2 = stakingManager.weightToValue(1e6);
        uint256 v3 = stakingManager.weightToValue(1e15);

        assertEq(v1, 1e12);
        assertEq(v2, 1e18);
        assertEq(v3, 1e27);
    }
}
