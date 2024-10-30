// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {ValidatorManager, SubnetConversionData, InitialValidator} from "../ValidatorManager.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";
import {
    ValidatorStatus,
    ValidatorRegistrationInput,
    PChainOwner,
    IValidatorManager
} from "../interfaces/IValidatorManager.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";

// TODO: Remove this once all unit tests implemented
// solhint-disable no-empty-blocks
abstract contract ValidatorManagerTest is Test {
    bytes32 public constant DEFAULT_SUBNET_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes public constant DEFAULT_NODE_ID =
        bytes(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes public constant DEFAULT_INITIAL_VALIDATOR_NODE_ID_1 =
        bytes(hex"2345678123456781234567812345678123456781234567812345678123456781");
    bytes public constant DEFAULT_INITIAL_VALIDATOR_NODE_ID_2 =
        bytes(hex"1345678123456781234567812345678123456781234567812345678123456781");
    bytes public constant DEFAULT_BLS_PUBLIC_KEY = bytes(
        hex"123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678"
    );
    bytes32 public constant DEFAULT_SOURCE_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    bytes32 public constant DEFAULT_SUBNET_CONVERSION_ID =
        bytes32(hex"76a386628f079b7b00452f8cab0925740363fcd52b721a8cf91773e857327b36");
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;

    uint64 public constant DEFAULT_WEIGHT = 1e6;
    // Set the default weight to 1e10 to avoid churn issues
    uint64 public constant DEFAULT_INITIAL_VALIDATOR_WEIGHT = DEFAULT_WEIGHT * 1e4;
    uint256 public constant DEFAULT_MINIMUM_STAKE_AMOUNT = 20e12;
    uint256 public constant DEFAULT_MAXIMUM_STAKE_AMOUNT = 1e22;
    uint64 public constant DEFAULT_CHURN_PERIOD = 1 hours;
    uint8 public constant DEFAULT_MAXIMUM_CHURN_PERCENTAGE = 20;
    uint64 public constant DEFAULT_EXPIRY = 1000;
    uint8 public constant DEFAULT_MAXIMUM_HOURLY_CHURN = 0;
    uint64 public constant DEFAULT_REGISTRATION_TIMESTAMP = 1000;
    uint256 public constant DEFAULT_STARTING_TOTAL_WEIGHT = 1e10 + DEFAULT_WEIGHT;
    uint64 public constant DEFAULT_MINIMUM_VALIDATION_DURATION = 24 hours;
    uint64 public constant DEFAULT_COMPLETION_TIMESTAMP =
        DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_VALIDATION_DURATION;
    // solhint-disable-next-line var-name-mixedcase
    PChainOwner public DEFAULT_P_CHAIN_OWNER;

    ValidatorManager public validatorManager;

    // Used to create unique validator IDs in {_newNodeID}
    uint64 public nodeIDCounter = 0;

    event ValidationPeriodCreated(
        bytes32 indexed validationID,
        bytes indexed nodeID,
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
        uint256 endTime
    );

    event ValidationPeriodEnded(bytes32 indexed validationID, ValidatorStatus indexed status);

    receive() external payable {}
    fallback() external payable {}

    function setUp() public virtual {
        address[] memory addresses = new address[](1);
        addresses[0] = 0x1234567812345678123456781234567812345678;
        DEFAULT_P_CHAIN_OWNER = PChainOwner({threshold: 1, addresses: addresses});
    }

    function testInitializeValidatorRegistrationSuccess() public {
        _setUpInitializeValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_SUBNET_ID,
            DEFAULT_WEIGHT,
            DEFAULT_EXPIRY,
            DEFAULT_BLS_PUBLIC_KEY
        );
    }

    function testInitializeValidatorRegistrationExcessiveChurn() public {
        // TODO: implement
    }

    function testInitializeValidatorRegistrationInsufficientStake() public {
        // TODO: implement
    }

    function testInitializeValidatorRegistrationExcessiveStake() public {
        // TODO: implement
    }

    function testInitializeValidatorRegistrationInsufficientDuration() public {
        // TODO: implement
    }

    function testInitializeValidatorRegistrationPChainOwnerThresholdTooLarge() public {
        // Threshold too large
        address[] memory addresses = new address[](1);
        addresses[0] = 0x1234567812345678123456781234567812345678;
        PChainOwner memory invalidPChainOwner1 = PChainOwner({threshold: 2, addresses: addresses});
        _beforeSend(_weightToValue(DEFAULT_WEIGHT), address(this));
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorManager.InvalidPChainOwnerThreshold.selector, 2, 1)
        );
        _initializeValidatorRegistration(
            ValidatorRegistrationInput({
                nodeID: DEFAULT_NODE_ID,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
                remainingBalanceOwner: invalidPChainOwner1,
                disableOwner: DEFAULT_P_CHAIN_OWNER,
                registrationExpiry: DEFAULT_EXPIRY
            }),
            DEFAULT_WEIGHT
        );
    }

    function testInitializeValidatorRegistrationZeroPChainOwnerThreshold() public {
        // Zero threshold for non-zero address
        address[] memory addresses = new address[](1);
        addresses[0] = 0x1234567812345678123456781234567812345678;
        PChainOwner memory invalidPChainOwner1 = PChainOwner({threshold: 0, addresses: addresses});
        _beforeSend(_weightToValue(DEFAULT_WEIGHT), address(this));
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorManager.InvalidPChainOwnerThreshold.selector, 0, 1)
        );
        _initializeValidatorRegistration(
            ValidatorRegistrationInput({
                nodeID: DEFAULT_NODE_ID,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
                remainingBalanceOwner: invalidPChainOwner1,
                disableOwner: DEFAULT_P_CHAIN_OWNER,
                registrationExpiry: DEFAULT_EXPIRY
            }),
            DEFAULT_WEIGHT
        );
    }

    function testInitializeValidatorRegistrationPChainOwnerAddressesUnsorted() public {
        // Addresses not sorted
        address[] memory addresses = new address[](2);
        addresses[0] = 0x1234567812345678123456781234567812345678;
        addresses[1] = 0x0123456781234567812345678123456781234567;
        PChainOwner memory invalidPChainOwner1 = PChainOwner({threshold: 1, addresses: addresses});

        _beforeSend(_weightToValue(DEFAULT_WEIGHT), address(this));
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorManager.PChainOwnerAddressesNotSorted.selector)
        );
        _initializeValidatorRegistration(
            ValidatorRegistrationInput({
                nodeID: DEFAULT_NODE_ID,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
                remainingBalanceOwner: invalidPChainOwner1,
                disableOwner: DEFAULT_P_CHAIN_OWNER,
                registrationExpiry: DEFAULT_EXPIRY
            }),
            DEFAULT_WEIGHT
        );
    }

    // The following tests call functions that are  implemented in ValidatorManager, but access state that's
    // only set in NativeTokenValidatorManager. Therefore we call them via the concrete type, rather than a
    // reference to the abstract type.
    function testResendRegisterValidatorMessage() public {
        bytes32 validationID = _setUpInitializeValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_SUBNET_ID,
            DEFAULT_WEIGHT,
            DEFAULT_EXPIRY,
            DEFAULT_BLS_PUBLIC_KEY
        );
        (, bytes memory registerSubnetValidatorMessage) = ValidatorMessages
            .packRegisterSubnetValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                subnetID: DEFAULT_SUBNET_ID,
                nodeID: DEFAULT_NODE_ID,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
                registrationExpiry: DEFAULT_EXPIRY,
                remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
                disableOwner: DEFAULT_P_CHAIN_OWNER,
                weight: DEFAULT_WEIGHT
            })
        );
        _mockSendWarpMessage(registerSubnetValidatorMessage, bytes32(0));
        validatorManager.resendRegisterValidatorMessage(validationID);
    }

    function testCompleteValidatorRegistration() public {
        _registerDefaultValidator();
    }

    function testInitializeEndValidation() public virtual {
        bytes32 validationID = _registerDefaultValidator();
        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            expectedNonce: 1,
            includeUptime: false,
            force: false
        });
    }

    function testResendEndValidation() public virtual {
        bytes32 validationID = _registerDefaultValidator();
        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            expectedNonce: 1,
            includeUptime: false,
            force: false
        });

        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSubnetValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        validatorManager.resendEndValidatorMessage(validationID);
    }

    function testCompleteEndValidation() public virtual {
        bytes32 validationID = _registerDefaultValidator();
        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            expectedNonce: 1,
            includeUptime: false,
            force: false
        });

        bytes memory subnetValidatorRegistrationMessage =
            ValidatorMessages.packSubnetValidatorRegistrationMessage(validationID, false);

        _mockGetPChainWarpMessage(subnetValidatorRegistrationMessage, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidationPeriodEnded(validationID, ValidatorStatus.Completed);

        validatorManager.completeEndValidation(0);
    }

    function testCompleteInvalidatedValidation() public {
        bytes32 validationID = _setUpInitializeValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_SUBNET_ID,
            DEFAULT_WEIGHT,
            DEFAULT_EXPIRY,
            DEFAULT_BLS_PUBLIC_KEY
        );
        bytes memory subnetValidatorRegistrationMessage =
            ValidatorMessages.packSubnetValidatorRegistrationMessage(validationID, false);

        _mockGetPChainWarpMessage(subnetValidatorRegistrationMessage, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidationPeriodEnded(validationID, ValidatorStatus.Invalidated);

        validatorManager.completeEndValidation(0);
    }

    function testInitialWeightsTooLow() public {
        vm.prank(address(123));
        IValidatorManager manager = _setUp();

        _mockGetBlockchainID();
        vm.expectRevert(abi.encodeWithSelector(ValidatorManager.InvalidTotalWeight.selector, 4));
        manager.initializeValidatorSet(_defaultSubnetConversionDataWeightsTooLow(), 0);
    }

    function testRemoveValidatorTotalWeight5() public {
        // Use prank here, because otherwise each test will end up with a different contract address, leading to a different subnet conversion hash.
        vm.prank(address(123));
        IValidatorManager manager = _setUp();

        _mockGetBlockchainID();
        _mockGetPChainWarpMessage(
            ValidatorMessages.packSubnetConversionMessage(
                bytes32(hex"1d72565851401e05d6351ebf5443d9bdc04953f3233da1345af126e7e4be7464")
            ),
            true
        );
        manager.initializeValidatorSet(_defaultSubnetConversionDataTotalWeight5(), 0);

        bytes32 validationID = sha256(abi.encodePacked(DEFAULT_SUBNET_ID, uint32(0)));
        vm.expectRevert(abi.encodeWithSelector(ValidatorManager.InvalidTotalWeight.selector, 4));
        _forceInitializeEndValidation(validationID, false);
    }

    function testCumulativeChurnRegistration() public {
        uint64 churnThreshold =
            uint64(DEFAULT_STARTING_TOTAL_WEIGHT) * DEFAULT_MAXIMUM_CHURN_PERCENTAGE / 100;
        _beforeSend(_weightToValue(churnThreshold), address(this));

        // First registration should succeed
        _registerValidator({
            nodeID: _newNodeID(),
            subnetID: DEFAULT_SUBNET_ID,
            weight: churnThreshold,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        _beforeSend(DEFAULT_MINIMUM_STAKE_AMOUNT, address(this));

        // Second call should fail
        vm.expectRevert(
            abi.encodeWithSelector(
                ValidatorManager.MaxChurnRateExceeded.selector,
                churnThreshold + _valueToWeight(DEFAULT_MINIMUM_STAKE_AMOUNT)
            )
        );
        _initializeValidatorRegistration(
            ValidatorRegistrationInput({
                nodeID: DEFAULT_NODE_ID,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
                remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
                disableOwner: DEFAULT_P_CHAIN_OWNER,
                registrationExpiry: DEFAULT_REGISTRATION_TIMESTAMP + 1
            }),
            _valueToWeight(DEFAULT_MINIMUM_STAKE_AMOUNT)
        );
    }

    function testCumulativeChurnRegistrationAndEndValidation() public {
        // Registration should succeed
        bytes32 validationID = _registerValidator({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: _valueToWeight(DEFAULT_MINIMUM_STAKE_AMOUNT),
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint64 churnThreshold =
            uint64(DEFAULT_STARTING_TOTAL_WEIGHT) * DEFAULT_MAXIMUM_CHURN_PERCENTAGE / 100;
        _beforeSend(_weightToValue(churnThreshold), address(this));

        // Registration should succeed
        _registerValidator({
            nodeID: _newNodeID(),
            subnetID: DEFAULT_SUBNET_ID,
            weight: churnThreshold,
            registrationExpiry: DEFAULT_EXPIRY + 25 hours,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP + 25 hours
        });

        // Second call should fail
        // The first registration churn amount is not part of the new churn amount since
        // a new churn period has started.
        vm.expectRevert(
            abi.encodeWithSelector(
                ValidatorManager.MaxChurnRateExceeded.selector,
                _valueToWeight(DEFAULT_MINIMUM_STAKE_AMOUNT) + churnThreshold
            )
        );
        _initializeEndValidation(validationID, false);
    }

    function testValidatorManagerStorageSlot() public view {
        assertEq(
            _erc7201StorageSlot("ValidatorManager"),
            validatorManager.VALIDATOR_MANAGER_STORAGE_LOCATION()
        );
    }

    function _newNodeID() internal returns (bytes memory) {
        nodeIDCounter++;
        return abi.encodePacked(sha256(new bytes(nodeIDCounter)));
    }

    function _setUpInitializeValidatorRegistration(
        bytes memory nodeID,
        bytes32 subnetID,
        uint64 weight,
        uint64 registrationExpiry,
        bytes memory blsPublicKey
    ) internal returns (bytes32 validationID) {
        (validationID,) = ValidatorMessages.packRegisterSubnetValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                nodeID: nodeID,
                subnetID: subnetID,
                blsPublicKey: blsPublicKey,
                registrationExpiry: registrationExpiry,
                remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
                disableOwner: DEFAULT_P_CHAIN_OWNER,
                weight: weight
            })
        );
        (, bytes memory registerSubnetValidatorMessage) = ValidatorMessages
            .packRegisterSubnetValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                subnetID: subnetID,
                nodeID: nodeID,
                blsPublicKey: blsPublicKey,
                registrationExpiry: registrationExpiry,
                remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
                disableOwner: DEFAULT_P_CHAIN_OWNER,
                weight: weight
            })
        );
        vm.warp(registrationExpiry - 1);
        _mockSendWarpMessage(registerSubnetValidatorMessage, bytes32(0));

        _beforeSend(_weightToValue(weight), address(this));
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidationPeriodCreated(validationID, nodeID, bytes32(0), weight, registrationExpiry);

        _initializeValidatorRegistration(
            ValidatorRegistrationInput({
                nodeID: nodeID,
                blsPublicKey: blsPublicKey,
                remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
                disableOwner: DEFAULT_P_CHAIN_OWNER,
                registrationExpiry: registrationExpiry
            }),
            weight
        );
    }

    function _registerValidator(
        bytes memory nodeID,
        bytes32 subnetID,
        uint64 weight,
        uint64 registrationExpiry,
        bytes memory blsPublicKey,
        uint64 registrationTimestamp
    ) internal returns (bytes32 validationID) {
        validationID = _setUpInitializeValidatorRegistration(
            nodeID, subnetID, weight, registrationExpiry, blsPublicKey
        );
        bytes memory subnetValidatorRegistrationMessage =
            ValidatorMessages.packSubnetValidatorRegistrationMessage(validationID, true);

        _mockGetPChainWarpMessage(subnetValidatorRegistrationMessage, true);

        vm.warp(registrationTimestamp);
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidationPeriodRegistered(validationID, weight, registrationTimestamp);

        validatorManager.completeValidatorRegistration(0);
    }

    function _initializeEndValidation(
        bytes32 validationID,
        uint64 registrationTimestamp,
        uint64 completionTimestamp,
        uint64 expectedNonce,
        bool includeUptime,
        bool force
    ) internal {
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSubnetValidatorWeightMessage(validationID, expectedNonce, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        if (includeUptime) {
            bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
                validationID, completionTimestamp - registrationTimestamp
            );
            _mockGetUptimeWarpMessage(uptimeMsg, true);
            _mockGetBlockchainID();
        }

        vm.warp(completionTimestamp);
        if (force) {
            _forceInitializeEndValidation(validationID, includeUptime);
        } else {
            _initializeEndValidation(validationID, includeUptime);
        }
    }

    function _registerDefaultValidator() internal returns (bytes32 validationID) {
        return _registerValidator({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
    }

    function _mockSendWarpMessage(bytes memory payload, bytes32 expectedMessageID) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(expectedMessageID)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.sendWarpMessage, payload)
        );
    }

    function _mockGetPChainWarpMessage(bytes memory expectedPayload, bool valid) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(
                WarpMessage({
                    sourceChainID: validatorManager.P_CHAIN_BLOCKCHAIN_ID(),
                    originSenderAddress: address(0),
                    payload: expectedPayload
                }),
                valid
            )
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, 0)
        );
    }

    function _mockGetUptimeWarpMessage(bytes memory expectedPayload, bool valid) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(
                WarpMessage({
                    sourceChainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
                    originSenderAddress: address(0),
                    payload: expectedPayload
                }),
                valid
            )
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, 0)
        );
    }

    function _mockGetBlockchainID() internal {
        _mockGetBlockchainID(DEFAULT_SOURCE_BLOCKCHAIN_ID);
    }

    function _mockGetBlockchainID(bytes32 blockchainID) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(blockchainID)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector)
        );
    }

    function _mockInitializeValidatorSet() internal {
        _mockGetPChainWarpMessage(
            ValidatorMessages.packSubnetConversionMessage(DEFAULT_SUBNET_CONVERSION_ID), true
        );
    }

    function _initializeValidatorRegistration(
        ValidatorRegistrationInput memory input,
        uint64 weight
    ) internal virtual returns (bytes32);

    function _initializeEndValidation(bytes32 validationID, bool includeUptime) internal virtual;

    function _forceInitializeEndValidation(
        bytes32 validationID,
        bool includeUptime
    ) internal virtual;

    function _setUp() internal virtual returns (IValidatorManager);

    function _beforeSend(uint256 amount, address spender) internal virtual;

    function _defaultSubnetConversionData() internal view returns (SubnetConversionData memory) {
        InitialValidator[] memory initialValidators = new InitialValidator[](2);
        // The first initial validator has a high weight relative to the default PoS validator weight
        // to avoid churn issues
        initialValidators[0] = InitialValidator({
            nodeID: DEFAULT_INITIAL_VALIDATOR_NODE_ID_1,
            weight: DEFAULT_INITIAL_VALIDATOR_WEIGHT,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY
        });
        // The second initial validator has a low weight so that it can be safely removed in tests
        initialValidators[1] = InitialValidator({
            nodeID: DEFAULT_INITIAL_VALIDATOR_NODE_ID_2,
            weight: DEFAULT_WEIGHT,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY
        });
        return SubnetConversionData({
            subnetID: DEFAULT_SUBNET_ID,
            validatorManagerBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            validatorManagerAddress: address(validatorManager),
            initialValidators: initialValidators
        });
    }

    function _defaultSubnetConversionDataWeightsTooLow()
        internal
        view
        returns (SubnetConversionData memory)
    {
        InitialValidator[] memory initialValidators = new InitialValidator[](2);

        initialValidators[0] = InitialValidator({
            nodeID: DEFAULT_INITIAL_VALIDATOR_NODE_ID_1,
            weight: 1,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY
        });
        initialValidators[1] = InitialValidator({
            nodeID: DEFAULT_INITIAL_VALIDATOR_NODE_ID_2,
            weight: 3,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY
        });

        return SubnetConversionData({
            subnetID: DEFAULT_SUBNET_ID,
            validatorManagerBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            validatorManagerAddress: address(validatorManager),
            initialValidators: initialValidators
        });
    }

    function _defaultSubnetConversionDataTotalWeight5()
        internal
        view
        returns (SubnetConversionData memory)
    {
        InitialValidator[] memory initialValidators = new InitialValidator[](2);

        initialValidators[0] = InitialValidator({
            nodeID: DEFAULT_INITIAL_VALIDATOR_NODE_ID_1,
            weight: 1,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY
        });
        initialValidators[1] = InitialValidator({
            nodeID: DEFAULT_INITIAL_VALIDATOR_NODE_ID_2,
            weight: 4,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY
        });

        return SubnetConversionData({
            subnetID: DEFAULT_SUBNET_ID,
            validatorManagerBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            validatorManagerAddress: address(validatorManager),
            initialValidators: initialValidators
        });
    }

    // This needs to be kept in line with the contract conversions, but we can't make external calls
    // to the contract and use vm.expectRevert at the same time.
    // These are okay to use for PoA as well, because they're just used for conversions inside the tests.
    function _valueToWeight(uint256 value) internal pure returns (uint64) {
        return uint64(value / 1e12);
    }

    // This needs to be kept in line with the contract conversions, but we can't make external calls
    // to the contract and use vm.expectRevert at the same time.
    // These are okay to use for PoA as well, because they're just used for conversions inside the tests.
    function _weightToValue(uint64 weight) internal pure returns (uint256) {
        return uint256(weight) * 1e12;
    }

    function _erc7201StorageSlot(bytes memory storageName) internal pure returns (bytes32) {
        return keccak256(
            abi.encode(
                uint256(keccak256(abi.encodePacked("avalanche-icm.storage.", storageName))) - 1
            )
        ) & ~bytes32(uint256(0xff));
    }
}
// solhint-enable no-empty-blocks
