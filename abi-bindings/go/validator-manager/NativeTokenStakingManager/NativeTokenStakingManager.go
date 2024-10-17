// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenstakingmanager

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// InitialValidator is an auto generated low-level Go binding around an user-defined struct.
type InitialValidator struct {
	NodeID       []byte
	BlsPublicKey []byte
	Weight       uint64
}

// PChainOwner is an auto generated low-level Go binding around an user-defined struct.
type PChainOwner struct {
	Threshold uint32
	Addresses []common.Address
}

// PoSValidatorManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type PoSValidatorManagerSettings struct {
	BaseSettings             ValidatorManagerSettings
	MinimumStakeAmount       *big.Int
	MaximumStakeAmount       *big.Int
	MinimumStakeDuration     uint64
	MinimumDelegationFeeBips uint16
	MaximumStakeMultiplier   uint8
	WeightToValueFactor      *big.Int
	RewardCalculator         common.Address
}

// SubnetConversionData is an auto generated low-level Go binding around an user-defined struct.
type SubnetConversionData struct {
	SubnetID                     [32]byte
	ValidatorManagerBlockchainID [32]byte
	ValidatorManagerAddress      common.Address
	InitialValidators            []InitialValidator
}

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	Status         uint8
	NodeID         []byte
	StartingWeight uint64
	MessageNonce   uint64
	Weight         uint64
	StartedAt      uint64
	EndedAt        uint64
}

// ValidatorManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type ValidatorManagerSettings struct {
	SubnetID               [32]byte
	ChurnPeriodSeconds     uint64
	MaximumChurnPercentage uint8
}

// ValidatorRegistrationInput is an auto generated low-level Go binding around an user-defined struct.
type ValidatorRegistrationInput struct {
	NodeID                []byte
	BlsPublicKey          []byte
	RegistrationExpiry    uint64
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}

// NativeTokenStakingManagerMetaData contains all meta data concerning the NativeTokenStakingManager contract.
var NativeTokenStakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADDRESS_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BIPS_CONVERSION_FACTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NATIVE_MINTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINativeMinter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POS_VALIDATOR_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimDelegationFees\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeEndDelegation\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeEndValidation\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitializeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structValidator\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"startingWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"endedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWeight\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structPoSValidatorManagerSettings\",\"components\":[{\"name\":\"baseSettings\",\"type\":\"tuple\",\"internalType\":\"structValidatorManagerSettings\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"churnPeriodSeconds\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"minimumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"weightToValueFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardCalculator\",\"type\":\"address\",\"internalType\":\"contractIRewardCalculator\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeDelegatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initializeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeValidatorRegistration\",\"inputs\":[{\"name\":\"registrationInput\",\"type\":\"tuple\",\"internalType\":\"structValidatorRegistrationInput\",\"components\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"disableOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initializeValidatorSet\",\"inputs\":[{\"name\":\"subnetConversionData\",\"type\":\"tuple\",\"internalType\":\"structSubnetConversionData\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialValidators\",\"type\":\"tuple[]\",\"internalType\":\"structInitialValidator[]\",\"components\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredValidators\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resendEndValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendRegisterValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendUpdateDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitUptimeProof\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"valueToWeight\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"weightToValue\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DelegationEnded\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewards\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorAdded\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"delegatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"delegatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRegistered\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRemovalInitialized\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitialValidatorCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UptimeUpdated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"uptime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodEnded\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumValidatorStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodRegistered\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRemovalInitialized\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorWeightUpdate\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DelegatorIneligibleForRewards\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSKeyLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationFee\",\"inputs\":[{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationID\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDelegatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitializationStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMaximumChurnPercentage\",\"inputs\":[{\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinStakeDuration\",\"inputs\":[{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeID\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPChainOwnerThreshold\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addressesLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRegistrationExpiry\",\"inputs\":[{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeAmount\",\"inputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeMultiplier\",\"inputs\":[{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidSubnetConversionID\",\"inputs\":[{\"name\":\"encodedSubnetConversionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedSubnetConversionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidTotalWeight\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidValidationID\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorManagerAddress\",\"inputs\":[{\"name\":\"validatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorManagerBlockchainID\",\"inputs\":[{\"name\":\"blockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWarpOriginSenderAddress\",\"inputs\":[{\"name\":\"senderAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpSourceChainID\",\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MaxChurnRateExceeded\",\"inputs\":[{\"name\":\"churnAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MaxWeightExceeded\",\"inputs\":[{\"name\":\"newValidatorWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MinStakeDurationNotPassed\",\"inputs\":[{\"name\":\"endTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NodeAlreadyRegistered\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PChainOwnerAddressesNotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnexpectedRegistrationStatus\",\"inputs\":[{\"name\":\"validRegistration\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"ValidatorIneligibleForRewards\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotPoS\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroWeightToValueFactor\",\"inputs\":[]}]",
	Bin: "0x363038303630343035323334383031353631303030663537356638306664356235303630343035313631353738633338303338303631353738633833333938313031363034303831393035323631303032653931363130313037353635623630303138313630303138313131313536313030343235373631303034323631303132633536356230333631303034663537363130303466363130303535353635623530363130313430353635623766663063353765313638343064663034306631353038386463326638316665333931633339323362656337336532336139363632656663396332323963366130303830353436383031303030303030303030303030303030303930303436306666313631353631303061353537363034303531363366393265653861393630653031623831353236303034303136303430353138303931303339306664356238303534363030313630303136303430316230333930383131363134363130313034353738303534363030313630303136303430316230333139313636303031363030313630343031623033393038313137383235353630343035313930383135323766633766353035623266333731616532313735656534393133663434393965316632363333613762353933363332316565643163646165623631313531383164323930363032303031363034303531383039313033393061313562353035363562356636303230383238343033313231353631303131373537356638306664356238313531363030323831313036313031323535373566383066643562393339323530353035303536356236333465343837623731363065303162356635323630323136303034353236303234356666643562363135363366383036313031346435663339356666336665363038303630343035323630303433363130363130316235353735663335363065303163383036333832383061323561313136313030656135373830363362613361346239373131363130303865353738303633626133613462393731343631303464633537383036336263356662666563313436313034666235373830363362656530613033663134363130353265353738303633633235376130663531343631303534643537383036336335393965323466313436313035366335373830363363393734643162363134363130353766353738303633643566323066663631343631303539333537383036336466393364386465313436313035626635373830363366643761633565373134363130356435353735663830666435623830363338323830613235613134363130343062353738303633393365323435393831343631303431663537383036333938663365326234313436313034336535373830363361336136356534383134363130343564353738303633613937373861376131343631303264663537383036336166326635666562313436313034376335373830363361666239383039363134363130343866353738303633623737316233626331343631303463323537356638306664356238303633333534353564656431313631303135633537383036333335343535646564313436313032646635373830363333613163666666363134363130333037353738303633343637656630366631343631303332363537383036333532393766616536313436313033343535373830363336303330356436323134363130333634353738303633363230363538353631343631303338643537383036333636343335616266313436313033626135373830363337333232313466383134363130336439353738303633373666373836323131343631303365633537356638306664356238303633303131386163633431343631303162393537383036333033323265643938313436313031646135373830363331353164333064313134363130316639353738303633316563343437323431343631303232343537383036333230643931623761313436313032343335373830363332356531633737363134363130323632353738303633326532313934643831343631303238313537383036333332396333653132313436313032623835373562356638306664356233343830313536313031633435373566383066643562353036313031643836313031643333363630303436313437363635363562363130356634353635623030356233343830313536313031653535373566383066643562353036313031643836313031663433363630303436313437613135363562363130363239353635623334383031353631303230343537356638306664356235303631303230643630306138313536356236303430353136306666393039313136383135323630323030313562363034303531383039313033393066333562333438303135363130323266353735663830666435623530363130316438363130323365333636303034363134373636353635623631303861353536356233343830313536313032346535373566383066643562353036313031643836313032356433363630303436313437623835363562363130386230353635623334383031353631303236643537356638306664356235303631303164383631303237633336363030343631343830363536356236313065386135363562333438303135363130323863353735663830666435623530363130326130363130323962333636303034363134376131353635623631306566653536356236303430353136303031363030313630343031623033393039313136383135323630323030313631303231623536356233343830313536313032633335373566383066643562353036313032643236303031363030313630393931623031383135363562363034303531363130323162393139303631343832373536356233343830313536313032656135373566383066643562353036313032663436313237313038313536356236303430353136316666666639303931313638313532363032303031363130323162353635623334383031353631303331323537356638306664356235303631303164383631303332313336363030343631343736363536356236313066353235363562333438303135363130333331353735663830666435623530363130316438363130333430333636303034363134383362353635623631306635643536356233343830313536313033353035373566383066643562353036313031643836313033356633363630303436313438353435363562363131303064353635623334383031353631303336663537356638306664356235303631303337383630313438313536356236303430353136336666666666666666393039313136383135323630323030313631303231623536356233343830313536313033393835373566383066643562353036313033616336313033613733363630303436313438393035363562363131326435353635623630343035313930383135323630323030313631303231623536356233343830313536313033633535373566383066643562353036313032613036313033643433363630303436313437613135363562363131326635353635623334383031353631303365343537356638306664356235303631303361633566383135363562333438303135363130336637353735663830666435623530363130316438363130343036333636303034363134373636353635623631313330393536356233343830313536313034313635373566383066643562353036313032306436303330383135363562333438303135363130343261353735663830666435623530363130316438363130343339333636303034363134376131353635623631313333343536356233343830313536313034343935373566383066643562353036313031643836313034353833363630303436313438353435363562363131336631353635623334383031353631303436383537356638306664356235303631303164383631303437373336363030343631343833623536356236313136316435363562363130336163363130343861333636303034363134386263353635623631313830663536356233343830313536313034396135373566383066643562353036313033616337663433313737313366376563626464646434626339396539356439303361646564616138383362326537633235353136313062643133653263376534373364303038313536356233343830313536313034636435373566383066643562353036313032643236303035363030313630393931623031383135363562333438303135363130346537353735663830666435623530363130316438363130346636333636303034363134376131353635623631313833353536356233343830313536313035303635373566383066643562353036313033616337666539323534366436393839353064646433383931306432653135656431643932336364306137623364646539653261366133663338303536353535396362303038313536356233343830313536313035333935373566383066643562353036313031643836313035343833363630303436313437613135363562363131613935353635623334383031353631303535383537356638306664356235303631303164383631303536373336363030343631343931643536356236313162623235363562363130336163363130353761333636303034363134376131353635623631316338663536356233343830313536313035386135373566383066643562353036313032306436303134383135363562333438303135363130353965353735663830666435623530363130356232363130356164333636303034363134376131353635623631316362323536356236303430353136313032316239313930363134396139353635623334383031353631303563613537356638306664356235303631303261303632303261333030383135363562333438303135363130356530353735663830666435623530363130336163363130356566333636303034363134613239353635623631316466343536356236313035666638333833383336313165326335363562363130363234353736303430353136333130333663663931363065313162383135323630303438313031383439303532363032343031356236303430353138303931303339306664356235303530353035363562356636313036333236313231373235363562356638333831353236303037383230313630323035323630343038303832323038313531363065303831303139303932353238303534393339343530393139323930393139303832393036306666313636303035383131313135363130363662353736313036366236313439333435363562363030353831313131353631303637633537363130363763363134393334353635623831353236303230303136303031383230313830353436313036393039303631346139343536356238303630316630313630323038303931303430323630323030313630343035313930383130313630343035323830393239313930383138313532363032303031383238303534363130366263393036313461393435363562383031353631303730373537383036303166313036313036646535373631303130303830383335343034303238333532393136303230303139313631303730373536356238323031393139303566353236303230356632303930356238313534383135323930363030313031393036303230303138303833313136313036656135373832393030333630316631363832303139313562353035303530393138333532353035303630303238323031353436303031363030313630343031623033383038323136363032303834303135323630303136303430316238323034383131363630343038343031353236303031363038303162383230343831313636303630383430313532363030313630633031623930393130343831313636303830383330313532363030333932383330313534313636306130393039313031353239303931353038313531363030353831313131353631303737323537363130373732363134393334353635623134363130376135353735663833383135323630303738333031363032303532363034303930383139303230353439303531363331373063633933333630653231623831353236313036316239313630666631363930363030343031363134616336353635623630363038313031353136303430353136333432613265306235363065313162383135323630303438313031383539303532363030313630303136303430316230333930393131363630323438323031353235663630343438323031353236303035363030313630393931623031393036336565356234386562393037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f393036333835343563313661393036303634303135663630343035313830383330333831383635616634313538303135363130383163353733643566383033653364356666643562353035303530353036303430353133643566383233653630316633643930383130313630316631393136383230313630343035323631303834333931393038313031393036313462643735363562363034303531383236336666666666666666313636306530316238313532363030343031363130383566393139303631346330383536356236303230363034303531383038333033383135663837356166313135383031353631303837623537336435663830336533643566666435623530353035303530363034303531336436303166313936303166383230313136383230313830363034303532353038313031393036313038396639313930363134633161353635623530353035303530353635623631303839663833383338333631316532633536356235663631303862393631323137323536356236303039383130313534393039313530363066663136313536313038653235373630343035313633376661623831653536306530316238313532363030343031363034303531383039313033393066643562363030353630303136303939316230313630303136303031363061303162303331363633343231336366373836303430353138313633666666666666666631363630653031623831353236303034303136303230363034303531383038333033383138363561666131353830313536313039323535373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130393439393139303631346331613536356238333630323030313335313436313039373235373630343035313633373262306137653736306531316238313532363032303834303133353630303438323031353236303234303136313036316235363562333036313039383336303630383530313630343038363031363134633435353635623630303136303031363061303162303331363134363130396262353736313039613136303630383430313630343038353031363134633435353635623630343035313633326638383132306436306532316238313532363030343031363130363162393139303631343832373536356235663631303963393630363038353031383536313463363035363562393035303930353035663830356238323831363366666666666666663136313031353631306362303537356636313039656336303630383830313838363134633630353635623833363366666666666666663136383138313130363130613032353736313061303236313463613535363562393035303630323030323831303139303631306131343931393036313463623935363562363130613164393036313464323435363562383035313630343035313931393235303566393136303038383830313931363130613335393136313464396635363562393038313532363032303031363034303531383039313033393032303534313436313061363535373830353136303430353136336134316637373266363065303162383135323631303631623931393036303034303136313463303835363562356636303032383835663031333538343630343035313630323030313631306139343932393139303931383235323630653031623630303136303031363065303162303331393136363032303832303135323630323430313930353635623630343038303531363031663139383138343033303138313532393038323930353236313061616539313631346439663536356236303230363034303531383038333033383138353561666131353830313536313061633935373364356638303365336435666664356235303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631306165633931393036313463316135363562393035303830383636303038303138333566303135313630343035313631306230343931393036313464396635363562393038313532363034303830353136303230393238313930303338333031383132303933393039333535363065303833303138313532363030323833353238343531383238343031353238343831303138303531363030313630303136303430316230333930383131363835383430313532356636303630383630313831393035323931353138313136363038303836303135323432313636306130383530313532363063303834303138313930353238343831353236303037386130313930393235323930323038313531383135343832393036306666313931363630303138333630303538313131313536313062383635373631306238363631343933343536356230323137393035353530363032303832303135313630303138323031393036313062396639303832363134646634353635623530363034303832383130313531363030323833303138303534363036303836303135313630383038373031353136306130383830313531363030313630303136303430316230333935383631363630303136303031363038303162303331393930393431363933393039333137363030313630343031623932383631363932393039323032393139303931313736303031363030313630383031623033313636303031363038303162393138353136393139303931303236303031363030313630633031623033313631373630303136306330316239313834313639313930393130323137393035353630633039303933303135313630303339303932303138303534363030313630303136303430316230333139313639323834313639323930393231373930393135353833303135313631306334333931313638353631346563333536356238323531363034303531393139353530363130633534393136313464396635363562363034303830353139313832393030333832323039303834303135313630303136303031363034303162303331363832353239303832393037663964343766656639646130373736363135343665363436643631383330626663626461393035303663326535656564333831393565383263346562316362646639303630323030313630343035313830393130333930613335303530383036313063613939303631346564363536356239303530363130396430353635623530363030343833303138313930353536303031383330313534363036343930363130636434393036303031363034303162393030343630666631363833363134656638353635623130313536313063663635373630343035313633353934333331376636306530316238313532363030343831303138323930353236303234303136313036316235363562356637335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363331653664393738393631306431613837363132313936353635623630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363130643361393139303631346330383536356236303230363034303531383038333033383138363561663431353830313536313064353535373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130643739393139303631346331613536356239303530356637335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f3633383632626661363338383630343035313832363366666666666666663136363065303162383135323630303430313631306462333931393036313530336135363562356636303430353138303833303338313836356166343135383031353631306463643537336435663830336533643566666435623530353035303530363034303531336435663832336536303166336439303831303136303166313931363832303136303430353236313064663439313930383130313930363134626437353635623930353035663630303238323630343035313631306530373931393036313464396635363562363032303630343035313830383330333831383535616661313538303135363130653232353733643566383033653364356666643562353035303530363034303531336436303166313936303166383230313136383230313830363034303532353038313031393036313065343539313930363134633161353635623930353038323831313436313065373135373630343035313633313837326663386436306530316238313532363030343831303138323930353236303234383130313834393035323630343430313631303631623536356235303530353036303039393039323031383035343630666631393136363030313137393035353530353035303530353635623631306539333832363132326131353635623631306562333537363034303531363333306566613938623630653031623831353236303034383130313833393035323630323430313631303631623536356235663631306562643833363131636232353635623531393035303630303238313630303538313131313536313065643435373631306564343631343933343536356231343631306566343537383036303430353136333137306363393333363065323162383135323630303430313631303631623931393036313461633635363562363130383966383338333631323263613536356235663830363130663038363132353834353635623630303330313534363130663136393038343631353063353536356239303530383031353830363130663262353735303630303136303031363034303162303338313131356231353631306634633537363034303531363332323264313634333630653231623831353236303034383130313834393035323630323430313631303631623536356239323931353035303536356236313038396638333833383336313235613835363562363130663635363132373733353635623566363130663665363132353834353635623930353035663830363130663762383436313237626435363562393135303931353036313066383838323631323261313536356236313066393435373530353035303631313030323536356235663832383135323630303538343031363032303532363034303930323035343630303136303031363061303162303331363630303438323531363030353831313131353631306663323537363130666332363134393334353635623033363130666537353735663833383135323630303838353031363032303532363034303831323038303534393139303535363130666535383238323631326236633536356235303562363130666664383136313066663838343630343030313531363131326435353635623631326263613536356235303530353035303562363131303061363132626464353635623530353635623566363131303136363132353834353635623566383338313532363030363832303136303230353236303430383038323230383135313630653038313031393039323532383035343933393435303931393239303931393038323930363066663136363030333831313131353631313034663537363131303466363134393334353635623630303338313131313536313130363035373631313036303631343933343536356238313532383135343631303130303930303436303031363030313630613031623033313636303230383230313532363030313832303135343630343038303833303139313930393135323630303239303932303135343630303136303031363034303162303338303832313636303630383430313532363030313630343031623832303438313136363038303834303135323630303136303830316238323034383131363630613038343031353236303031363063303162393039313034313636306330393039313031353238313031353139303931353035663631313064363832363131636232353635623930353036303031383335313630303338313131313536313130656435373631313065643631343933343536356231343631313130653537383235313630343035313633336230643534306436306532316238313532363130363162393139303630303430313631353065343536356236303034383135313630303538313131313536313131323335373631313132333631343933343536356230333631313133393537363131313331383536313263303335363562353035303530353035303530353635623566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363339646532336434303631313135653861363132313936353635623630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363131313765393139303631346330383536356236303630363034303531383038333033383138363561663431353830313536313131393935373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363131316264393139303631353066653536356235303931353039313530383138343134363131316561353738343630343030313531363034303531363330383939333862333630653131623831353236303034303136313036316239313831353236303230303139303536356238303630303136303031363034303162303331363833363036303031353136303031363030313630343031623033313631303830363131323233353735303830363030313630303136303430316230333136383536306130303135313630303136303031363034303162303331363131356231353631313234633537363034303531363332653139626332643630653131623831353236303031363030313630343031623033383231363630303438323031353236303234303136313036316235363562356638373831353236303036383730313630323039303831353236303430393138323930323038303534363066663139313636303032393038313137383235353031383035343630303136303031363034303162303334323136363030313630343031623831303236376666666666666666666666666666666636303430316231393930393231363931393039313137393039313535393135313931383235323835393138393931376630343730353962343635303639623862373531383336623431663966316438336461666635383364323233386363376662623436313433376563323361346636393130313630343035313830393130333930613335303530353035303530353035303530353635623566363131326465363132353834353635623630303330313534363130663463393036303031363030313630343031623033383431363631346566383536356235663631313266663832363131636232353635623630383030313531393239313530353035363562363131333134383338333833363132356138353635623631303632343537363034303531363335626666363833663630653131623831353236303034383130313834393035323630323430313631303631623536356235663631313333643631323538343536356239303530356636313133343938333631316362323536356235313930353036303034383136303035383131313135363131333630353736313133363036313439333435363562313436313133383035373830363034303531363331373063633933333630653231623831353236303034303136313036316239313930363134616336353635623566383338313532363030353833303136303230353236303430393032303534363030313630303136306130316230333136333331343631313362623537333335623630343035313633366532636364373536306531316238313532363030343031363130363162393139303631343832373536356235663833383135323630303838333031363032303930383135323630343038303833323038303534393038343930353536303035383630313930393235323930393132303534363130383966393036303031363030313630613031623033313638323631326236633536356236313133663936313237373335363562356636313134303236313235383435363562356638333831353236303036383230313630323035323630343038303832323038313531363065303831303139303932353238303534393339343530393139323930393139303832393036306666313636303033383131313135363131343362353736313134336236313439333435363562363030333831313131353631313434633537363131343463363134393334353635623831353238313534363130313030393030343630303136303031363061303162303331363630323038323031353236303031383230313534363034303832303135323630303239303931303135343630303136303031363034303162303338303832313636303630383430313532363030313630343031623832303438313136363038303834303135323630303136303830316238323034383131363630613038343031353236303031363063303162393039313034313636306330393039313031353239303530363030333831353136303033383131313135363131346335353736313134633536313439333435363562313436313134653635373830353136303430353136333362306435343064363065323162383135323631303631623931393036303034303136313530653435363562363030343631313466353832363034303031353136313163623235363562353136303035383131313135363131353037353736313135303736313439333435363562313436313136303635373566363131353136383536313231393635363562393035303566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f3633396465323364343038343630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363131353535393139303631346330383536356236303630363034303531383038333033383138363561663431353830313536313135373035373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363131353934393139303631353066653536356235303931353039313530383138343630343030313531313436313135633035373630343035313633303839393338623336306531316238313532363030343831303138333930353236303234303136313036316235363562383036303031363030313630343031623033313638343630633030313531363030313630303136303430316230333136313131353631313630323537363034303531363332653139626332643630653131623831353236303031363030313630343031623033383231363630303438323031353236303234303136313036316235363562353035303530356236313136306638333631326330333536356235303530363131363139363132626464353635623530353035363562356636313136323636313231373235363562393035303566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f36333265343363656235363131363464383636313231393635363562363034303031353136303430353138323633666666666666666631363630653031623831353236303034303136313136366439313930363134633038353635623630343038303531383038333033383138363561663431353830313536313136383735373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363131366162393139303631353133333536356239313530393135303830363131366431353736303430353136333264303731333533363065303162383135323831313531353630303438323031353236303234303136313036316235363562356638323831353236303036383430313630323035323630343039303230383035343631313665623930363134613934353635623930353035663033363131373066353736303430353136333038393933386233363065313162383135323630303438313031383339303532363032343031363130363162353635623630303135663833383135323630303738353031363032303532363034303930323035343630666631363630303538313131313536313137333535373631313733353631343933343536356231343631313736383537356638323831353236303037383430313630323035323630343039303831393032303534393035313633313730636339333336306532316238313532363130363162393136306666313639303630303430313631346163363536356235663832383135323630303638343031363032303532363034303831323036313137383039313631343662613536356235663832383135323630303738343031363032303930383135323630343039313832393032303830353436306666313931363630303239303831313738323535303138303534363030313630303136303430316230333432383138313136363030313630633031623032363030313630303136306330316230333930393331363932393039323137393238333930353538343531363030313630383031623930393330343136383235323931383130313931393039313532383339313766663866643163393066623963666132636132333538666466353830366230383661643433333135643932623232316339323965666337663130356365373536383931303136303430353138303931303339306132353035303530353035363562356636313138313836313237373335363562363131383234383438343834333436313265323335363562393035303631313832653631326264643536356239333932353035303530353635623566363131383365363132353834353635623566383338313532363030363832303136303230353236303430383038323230383135313630653038313031393039323532383035343933393435303931393239303931393038323930363066663136363030333831313131353631313837373537363131383737363134393334353635623630303338313131313536313138383835373631313838383631343933343536356238313532383135343631303130303930303436303031363030313630613031623033313636303230383230313532363030313830383330313534363034303833303135323630303239303932303135343630303136303031363034303162303338303832313636303630383430313532363030313630343031623832303438313136363038303834303135323630303136303830316238323034383131363630613038343031353236303031363063303162393039313034313636306330393039313031353239303931353038313531363030333831313131353631313930313537363131393031363134393334353635623134313538303135363131393232353735303630303338313531363030333831313131353631313931663537363131393166363134393334353635623134313535623135363131393433353738303531363034303531363333623064353430643630653231623831353236313036316239313930363030343031363135306534353635623566363131393531383236303430303135313631316362323536356239303530383036303630303135313630303136303031363034303162303331363566303336313139383335373630343035313633333962383934663936306532316238313532363030343831303138353930353236303234303136313036316235363562363034303830383330313531363036303833303135313630383038343031353139323531363334326132653062353630653131623831353236303035363030313630393931623031393336336565356234386562393337335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f3933363338353435633136613933363131396631393339303630303430313932383335323630303136303031363034303162303339313832313636303230383430313532313636303430383230313532363036303031393035363562356636303430353138303833303338313836356166343135383031353631316130623537336435663830336533643566666435623530353035303530363034303531336435663832336536303166336439303831303136303166313931363832303136303430353236313161333239313930383130313930363134626437353635623630343035313832363366666666666666663136363065303162383135323630303430313631316134653931393036313463303835363562363032303630343035313830383330333831356638373561663131353830313536313161366135373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363131613865393139303631346331613536356235303530353035303530353635623566363131613965363132313732353635623566383338313532363030363832303136303230353236303430393032303830353439313932353039303631316162633930363134613934353635623930353035663033363131616530353736303430353136333038393933386233363065313162383135323630303438313031383339303532363032343031363130363162353635623630303135663833383135323630303738333031363032303532363034303930323035343630666631363630303538313131313536313162303635373631316230363631343933343536356231343631316233393537356638323831353236303037383230313630323035323630343039303831393032303534393035313633313730636339333336306532316238313532363130363162393136306666313639303630303430313631346163363536356235663832383135323630303638323031363032303532363034303930383139303230393035313633656535623438656236306530316238313532363030353630303136303939316230313931363365653562343865623931363131623732393139303630303430313631353136313536356236303230363034303531383038333033383135663837356166313135383031353631316238653537336435663830336533643566666435623530353035303530363034303531336436303166313936303166383230313136383230313830363034303532353038313031393036313036323439313930363134633161353635623766663063353765313638343064663034306631353038386463326638316665333931633339323362656337336532336139363632656663396332323963366130303830353436303032393139303630303136303430316239303034363066663136383036313162666235373530383035343630303136303031363034303162303338303834313639313136313031353562313536313163313935373630343035313633663932656538613936306530316238313532363030343031363034303531383039313033393066643562383035343638666666666666666666666666666666666666313931363630303136303031363034303162303338333136313736303031363034303162313738313535363131633433383336313266636635363562383035343630666636303430316231393136383135353630343035313630303136303031363034303162303338333136383135323766633766353035623266333731616532313735656534393133663434393965316632363333613762353933363332316565643163646165623631313531383164323930363032303031363034303531383039313033393061313530353035303536356235663631316339383631323737333536356236313163613338323333333436313266653035363562393035303631316361643631326264643536356239313930353035363562363131636261363134366631353635623566363131636333363132313732353635623566383438313532363030373832303136303230353236303430393038313930323038313531363065303831303139303932353238303534393239333530393039313832393036306666313636303035383131313135363131636661353736313163666136313439333435363562363030353831313131353631316430623537363131643062363134393334353635623831353236303230303136303031383230313830353436313164316639303631346139343536356238303630316630313630323038303931303430323630323030313630343035313930383130313630343035323830393239313930383138313532363032303031383238303534363131643462393036313461393435363562383031353631316439363537383036303166313036313164366435373631303130303830383335343034303238333532393136303230303139313631316439363536356238323031393139303566353236303230356632303930356238313534383135323930363030313031393036303230303138303833313136313164373935373832393030333630316631363832303139313562353035303530393138333532353035303630303238323031353436303031363030313630343031623033383038323136363032303834303135323630303136303430316238323034383131363630343038343031353236303031363038303162383230343831313636303630383430313532363030313630633031623930393130343831313636303830383330313532363030333930393230313534393039313136363061303930393130313532393339323530353035303536356235663830363131646665363132313732353635623930353038303630303830313834383436303430353136313165313439323931393036313531656235363562393038313532363032303031363034303531383039313033393032303534393135303530393239313530353035363562356638303631316533363631323538343536356235663836383135323630303638323031363032303532363034303830383232303831353136306530383130313930393235323830353439333934353039313932393039313930383239303630666631363630303338313131313536313165366635373631316536663631343933343536356236303033383131313135363131653830353736313165383036313439333435363562383135323831353436313031303039303034363030313630303136306130316230333136363032303832303135323630303138323031353436303430383038333031393139303931353236303032393039323031353436303031363030313630343031623033383038323136363036303834303135323630303136303430316238323034383131363630383038343031353236303031363038303162383230343831313636306130383430313532363030313630633031623930393130343136363063303930393130313532383130313531393039313530356636313165663638323631316362323536356239303530363030323833353136303033383131313135363131663064353736313166306436313439333435363562313436313166326535373832353136303430353136333362306435343064363065323162383135323631303631623931393036303034303136313530653435363562363032303833303135313630303136303031363061303162303331363333313436313166636135373566383238313532363030353835303136303230353236303430393032303534363030313630303136306130316230333136333331343631316636373537333336313133613135363562356638323831353236303035383530313630323035323630343039303230353436306130383230313531363131663936393136303031363062303162393030343630303136303031363034303162303331363930363135316661353635623630303136303031363034303162303331363432313031353631316663613537363034303531363366623663653633663630653031623831353236303031363030313630343031623033343231363630303438323031353236303234303136313036316235363562363030323831353136303035383131313135363131666466353736313166646636313439333435363562303336313231306435373630303238343031353436303830383430313531363132303031393136303031363030313630343031623033313639303631353166613536356236303031363030313630343031623033313634323130313536313230333535373630343035313633666236636536336636306530316238313532363030313630303136303430316230333432313636303034383230313532363032343031363130363162353635623836313536313230343735373631323034353832383736313232636135363562353035623566383838313532363030363835303136303230353236303430393032303830353436306666313931363630303331373930353536303630383330313531363038303832303135313631323038303931383439313631323037623931393036313532316135363562363133326232353635623530356638393831353236303036383630313630323035323630343038313230363030323031383035343630303136303031363034303162303339303933313636303031363063303162303236303031363030313630633031623033393039333136393239303932313739303931353536313230633138343631333437633536356235663861383135323630303738373031363032303532363034303830383232303833393035353531393139323530383439313862393137663336366433333663306162333830646337393966303935613666383261323633323635383563353239303963633639386230396261343534303730396564353739316133313531353934353036313138326539333530353035303530353635623630303438313531363030353831313131353631323132323537363132313232363134393334353635623033363132313536353736313231333038333631333437633536356235663839383135323630303738363031363032303532363034303930323035353631323134393838363132633033353635623630303139343530353035303530353036313138326535363562383035313630343035313633313730636339333336306532316238313532363130363162393139303630303430313631346163363536356237666539323534366436393839353064646433383931306432653135656431643932336364306137623364646539653261366133663338303536353535396362303039303536356236303430383035313630363038303832303138333532356638303833353236303230383330313532393138313031393139303931353236303430353136333036663832353335363065343162383135323633666666666666666638333136363030343832303135323566393038313930363030353630303136303939316230313930363336663832353335303930363032343031356636303430353138303833303338313836356166613135383031353631323166613537336435663830336533643566666435623530353035303530363034303531336435663832336536303166336439303831303136303166313931363832303136303430353236313232323139313930383130313930363135323361353635623931353039313530383036313232343335373630343035313633366232663139653936306530316238313532363030343031363034303531383039313033393066643562383135313135363132323639353738313531363034303531363336626135383961353630653031623831353236303034383130313931393039313532363032343031363130363162353635623630323038323031353136303031363030313630613031623033313631353631323239613537383136303230303135313630343035313632346465373564363065333162383135323630303430313631303631623931393036313438323735363562353039323931353035303536356235663830363132326162363132353834353635623566393338343532363030353031363032303532353035303630343039303230353436303031363030313630613031623033313631353135393035363562363034303531363330366638323533353630653431623831353236336666666666666666383231363630303438323031353235663930383139303831393036303035363030313630393931623031393036333666383235333530393036303234303135663630343035313830383330333831383635616661313538303135363132333135353733643566383033653364356666643562353035303530353036303430353133643566383233653630316633643930383130313630316631393136383230313630343035323631323333633931393038313031393036313532336135363562393135303931353038303631323335653537363034303531363336623266313965393630653031623831353236303034303136303430353138303931303339306664356236303035363030313630393931623031363030313630303136306130316230333136363334323133636637383630343035313831363366666666666666663136363065303162383135323630303430313630323036303430353138303833303338313836356166613135383031353631323361313537336435663830336533643566666435623530353035303530363034303531336436303166313936303166383230313136383230313830363034303532353038313031393036313233633539313930363134633161353635623832353131343631323365623537383135313630343035313633366261353839613536306530316238313532363030343831303139313930393135323630323430313631303631623536356236303230383230313531363030313630303136306130316230333136313536313234316335373831363032303031353136303430353136323464653735643630653331623831353236303034303136313036316239313930363134383237353635623566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363330383863323436333835363034303031353136303430353138323633666666666666666631363630653031623831353236303034303136313234353939313930363134633038353635623630343038303531383038333033383138363561663431353830313536313234373335373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363132343937393139303631353263613536356239313530393135303831383731343631323462653537363034303531363330383939333862333630653131623831353236303034383130313838393035323630323430313631303631623536356235663631323463373631323538343536356235663839383135323630303538323031363032303532363034303930323036303031303135343930393135303630303136303031363034303162303339303831313639303833313631313135363132353561353735663838383135323630303538323031363032303930383135323630343039313832393032303630303130313830353436303031363030313630343031623033313931363630303136303031363034303162303338363136393038313137393039313535393135313931383235323839393137666563343431343865386666323731663264306261636566313134323135346162616362306162623361323965623365623530653263613937653836643034333539313031363034303531383039313033393061323631323537393536356235663838383135323630303538323031363032303532363034303930323036303031303135343630303136303031363034303162303331363931353035623530393639353530353035303530353035303536356237663433313737313366376563626464646434626339396539356439303361646564616138383362326537633235353136313062643133653263376534373364303039303536356235663830363132356232363132353834353635623930353035663631323562653836363133356666353635623930353036313235633938363631323261313536356236313235643835373630303139323530353035303631313832653536356235663836383135323630303538333031363032303532363034303930323035343630303136303031363061303162303331363333313436313235666435373333363131336131353635623566383638313532363030353833303136303230353236303430393032303534363061303832303135313631323632633931363030313630623031623930303436303031363030313630343031623033313639303631353166613536356236303031363030313630343031623033313638313630633030313531363030313630303136303430316230333136313031353631323637333537363063303831303135313630343035313633666236636536336636306530316238313532363030313630303136303430316230333930393131363630303438323031353236303234303136313036316235363562356638353135363132363862353736313236383438373836363132326361353635623930353036313236613935363562353035663836383135323630303538333031363032303532363034303930323036303031303135343630303136303031363034303162303331363562363030343833303135343630343038333031353135663931363030313630303136306130316230333136393036333466323234323966393036313236636539303631313264353536356236306130383630313531363063303837303135313630343035313630303136303031363065303162303331393630653038363930316231363831353236313236666539333932393138323931383939303630303430313631353265643536356236303230363034303531383038333033383138363561666131353830313536313237313935373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363132373364393139303631346331613536356239303530383038343630303830313566386138313532363032303031393038313532363032303031356632303566383238323534363132373631393139303631346563333536356239303931353535303530313531353937393635303530353035303530353035303536356237663962373739623137343232643064663932323233303138623332623464316661343665303731373233643638313765323438366430303362656363353566303038303534363030313139303136313237623735373630343035313633336565356165623536306530316238313532363030343031363034303531383039313033393066643562363030323930353535363562356636313237633636313436663135363562356636313237636636313231373235363562393035303566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f36333265343363656235363132376636383836313231393635363562363034303031353136303430353138323633666666666666666631363630653031623831353236303034303136313238313639313930363134633038353635623630343038303531383038333033383138363561663431353830313536313238333035373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363132383534393139303631353133333536356239313530393135303830313536313238376235373630343035313633326430373133353336306530316238313532383131353135363030343832303135323630323430313631303631623536356235663832383135323630303738343031363032303532363034303830383232303831353136306530383130313930393235323830353438323930363066663136363030353831313131353631323861633537363132386163363134393334353635623630303538313131313536313238626435373631323862643631343933343536356238313532363032303031363030313832303138303534363132386431393036313461393435363562383036303166303136303230383039313034303236303230303136303430353139303831303136303430353238303932393139303831383135323630323030313832383035343631323866643930363134613934353635623830313536313239343835373830363031663130363132393166353736313031303038303833353430343032383335323931363032303031393136313239343835363562383230313931393035663532363032303566323039303562383135343831353239303630303130313930363032303031383038333131363132393262353738323930303336303166313638323031393135623530353035303931383335323530353036303032383230313534363030313630303136303430316230333830383231363630323038343031353236303031363034303162383230343831313636303430383430313532363030313630383031623832303438313136363036303834303135323630303136306330316239303931303438313136363038303833303135323630303339323833303135343136363061303930393130313532393039313530383135313630303538313131313536313239623335373631323962333631343933343536356231343135383031353631323964343537353036303031383135313630303538313131313536313239643135373631323964313631343933343536356231343135356231353631323966353537383035313630343035313633313730636339333336306532316238313532363130363162393139303630303430313631346163363536356236303033383135313630303538313131313536313261306135373631326130613631343933343536356230333631326131383537363030343831353236313261316435363562363030353831353235623833363030383031383136303230303135313630343035313631326133333931393036313464396635363562393038313532363034303830353136303230393238313930303338333031393032303566393038313930353538353831353236303037383730313930393235323930323038313531383135343833393239313930383239303630666631393136363030313833363030353831313131353631326137373537363132613737363134393334353635623032313739303535353036303230383230313531363030313832303139303631326139303930383236313464663435363562353036303430383230313531363030323832303138303534363036303835303135313630383038363031353136306130383730313531363030313630303136303430316230333935383631363630303136303031363038303162303331393930393431363933393039333137363030313630343031623932383631363932393039323032393139303931313736303031363030313630383031623033313636303031363038303162393138353136393139303931303236303031363030313630633031623033313631373630303136306330316239313834313639313930393130323137393035353630633039303932303135313630303339303931303138303534363030313630303136303430316230333139313639313930393231363137393035353830353136303035383131313135363132623335353736313262333536313439333435363562363034303531383439303766316330386535393635366631613138646332646137363832366364633532383035633433653839376131376335306661656662386162336331353236636331363930356639306133393139363931393535303930393335303530353035303536356236303430353136333237616435353564363065313162383135323630303136303031363061303162303338333136363030343832303135323630323438313031383239303532363030313630303136303939316230313930363334663561616162613930363034343031356636303430353138303833303338313566383738303362313538303135363132626238353735663830666435623530356166313135383031353631313133313537336435663830336533643566666435623631313631393630303136303031363061303162303338333136383236313338643735363562363030313766396237373962313734323264306466393232323330313862333262346431666134366530373137323364363831376532343836643030336265636335356630303535353635623566363132633063363132353834353635623566383338313532363030363832303136303230353236303430383038323230383135313630653038313031393039323532383035343933393435303931393239303931393038323930363066663136363030333831313131353631326334353537363132633435363134393334353635623630303338313131313536313263353635373631326335363631343933343536356238313532383135343631303130303930303436303031363030313630613031623033313636303230383230313532363030313832303135343630343038303833303139313930393135323630303239303932303135343630303136303031363034303162303338303832313636303630383430313532363030313630343031623832303438313136363038303834303135323630303136303830316238323034383131363630613038343031353236303031363063303162393039313034313636306330393039313031353238313031353139303931353036313263636136313339366135363562383236303830303135313631326364393931393036313531666135363562363030313630303136303430316230333136343231303135363132643064353736303430353136336662366365363366363065303162383135323630303136303031363034303162303334323136363030343832303135323630323430313631303631623536356235663834383135323630303638343031363032303930383135323630343038303833323038303534363030313630303136306138316230333139313638313535363030313831303138343930353536303032303138333930353536303037383630313930393135323831323038303534393038323930353539303830383231353631326463383537356638343831353236303035383730313630323035323630343039303230353436313237313039303631326437613930363030313630613031623930303436316666666631363835363134656638353635623631326438343931393036313530633535363562393135303831383636303038303135663836383135323630323030313930383135323630323030313566323035663832383235343631326461383931393036313465633335363562393039313535353036313264623839303530383238343631353331623536356239303530363132646338383536303230303135313832363132623663353635623631326464643835363032303031353136313066663838373630363030313531363131326435353635623630343038303531383238313532363032303831303138343930353238353931383939313766386563656366353130303730633332306439613535333233666661626533353065323934616535303566633063353039646335373336646136663563633939333931303136303430353138303931303339306133353035303530353035303530353035363562356638303631326532643631323538343536356236303032383130313534393039313530363166666666363030313630343031623930393130343831313639303836313631303830363132653536353735303631323731303631666666663836313631313562313536313265376135373630343035313633356631326536633336306531316238313532363166666666383631363630303438323031353236303234303136313036316235363562363030323831303135343630303136303031363034303162303339303831313639303835313631303135363132656236353736303430353136323032613036643630653131623831353236303031363030313630343031623033383531363630303438323031353236303234303136313036316235363562383035343833313038303631326563383537353038303630303130313534383331313562313536313265653935373630343035313633323232643136343336306532316238313532363030343831303138343930353236303234303136313036316235363562383235663631326566343832363130656665353635623930353035663631326630313839383336313339383535363562393035303630343035313830363038303031363034303532383036313266313533333930353635623630303136303031363061303162303339303831313638323532363166666666383038633136363032303830383530313931393039313532363030313630303136303430316230333830386431363630343038303837303139313930393135323566363036303936383730313831393035323838383135323630303539303962303138333532393938613930323038363531383135343933383830313531396238383031353138333136363030313630623031623032363766666666666666666666666666666666363062303162313939633930393531363630303136306130316230323630303136303031363062303162303331393930393431363935313639343930393431373931393039313137393839303938313631373831353539313031353136303031393039313031383035343931393039353136363030313630303136303430316230333139393039313136313739303933353535303930393135303530393439333530353035303530353635623631326664373631336635323536356236313130306138313631336639643536356235663830363132666561363132353834353635623930353035663631326666363834363130656665353635623930353035663631333030323837363131636232353635623930353036313330306438373631323261313536356236313330326435373630343035313633333065666139386236306530316238313532363030343831303138383930353236303234303136313036316235363562363030323831353136303035383131313135363133303432353736313330343236313439333435363562313436313330363335373830353136303430353136333137306363393333363065323162383135323631303631623931393036303034303136313461633635363562356638323832363038303031353136313330373439313930363135316661353635623930353038333630303230313630306139303534393036313031303030613930303436303031363030313630343031623033313638323630343030313531363133303964393139303631353332653536356236303031363030313630343031623033313638313630303136303031363034303162303331363131313536313330646135373630343035313633366435316665303536306531316238313532363030313630303136303430316230333832313636303034383230313532363032343031363130363162353635623566383036313330653638613834363133326232353635623931353039313530356638613833363034303531363032303031363133313134393239313930393138323532363063303162363030313630303136306330316230333139313636303230383230313532363032383031393035363562363034303830353136303166313938313834303330313831353238323832353238303531363032303930393130313230363065303833303139303931353239313530383036303031383135323630303136303031363061303162303338633136363032303830383330313931393039313532363034303830383330313866393035323630303136303031363034303162303338303862313636303630383530313532356636303830383530313831393035323930383831363630613038353031353236306330393039333031383339303532383438333532363030363862303139303931353239303230383135313831353438323930363066663139313636303031383336303033383131313135363133316137353736313331613736313439333435363562303231373930353535303630323038323831303135313832353436313031303036303031363061383162303331393136363130313030363030313630303136306130316230333932383331363032313738333535363034303830383530313531363030313835303135353630363038303836303135313630303239303935303138303534363038303830383930313531363061303861303135313630633039303961303135313630303136303031363034303162303339393861313636303031363030313630383031623033313939303934313639333930393331373630303136303430316239313861313639313930393130323137363030313630303136303830316230333136363030313630383031623939383931363939393039393032363030313630303136306330316230333136393839303938313736303031363063303162393138383136393139303931303231373930353538313531383938363136383135323861383631363934383130313934393039343532393338623136393038333031353239313831303138353930353239303863313639313864393138343931376662303032346232363362633361306237323861366564656135306136396566613834313138396638643332656538616639643163326231613161323233343236393130313630343035313830393130333930613439613939353035303530353035303530353035303530353035363562356638303566363133326264363132313732353635623566383638313532363030373832303136303230353236303430393032303630303230313534393039313530363030313630383031623930303436303031363030313630343031623033313636313332656438353832363134303131353635623566363133326637383736313432333635363562356638383831353236303037383530313630323035323630343038303832323036303032303138303534363766666666666666666666666666666666363038303162313931363630303136303830316236303031363030313630343031623033386338313136393138323032393239303932313739303932353539313531363334326132653062353630653131623831353236303034383130313863393035323931383431363630323438333031353236303434383230313532393139323530393036303035363030313630393931623031393036336565356234386562393037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f3930363338353435633136613930363036343031356636303430353138303833303338313836356166343135383031353631333361303537336435663830336533643566666435623530353035303530363034303531336435663832336536303166336439303831303136303166313931363832303136303430353236313333633739313930383130313930363134626437353635623630343035313832363366666666666666663136363065303162383135323630303430313631333365333931393036313463303835363562363032303630343035313830383330333831356638373561663131353830313536313333666635373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363133343233393139303631346331613536356236303430383035313630303136303031363034303162303338613831313638323532363032303832303138343930353238323531393339343530383531363932386239323766303764653566663335613637346138303035653636316633333333633930376361363333333436323830383736326431396463376233616262316138633164663932383239303033303139306133393039343530393235303530353035623932353039323930353035363562356638303631333438363631323538343536356239303530356636313334393638343630343030313531363131636232353635623930353035663630303338323531363030353831313131353631333461653537363133346165363134393334353635623134383036313334636335373530363030343832353136303035383131313135363133346361353736313334636136313439333435363562313435623135363133346463353735303630633038313031353136313335313935363562363030323832353136303035383131313135363133346631353736313334663136313439333435363562303336313334666435373530343236313335313935363562383135313630343035313633313730636339333336306532316238313532363130363162393139303630303430313631346163363536356238343630383030313531363030313630303136303430316230333136383136303031363030313630343031623033313631313631333534303537353035663934393335303530353035303536356236303034383330313534363036303836303135313630303136303031363061303162303339303931313639303633346632323432396639303631333536353930363131326435353635623630613038353031353136303830383930313531363034303830386230313531356639303831353236303035386130313630323035323831393032303630303130313534393035313630303136303031363065303162303331393630653038373930316231363831353236313335623739343933393239313838393136303031363030313630343031623033393039313136393036303034303136313532656435363562363032303630343035313830383330333831383635616661313538303135363133356432353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631333566363931393036313463316135363562393539343530353035303530353035363562363133363037363134366631353635623566363133363130363132313732353635623566383438313532363030373832303136303230353236303430383038323230383135313630653038313031393039323532383035343933393435303931393239303931393038323930363066663136363030353831313131353631333634393537363133363439363134393334353635623630303538313131313536313336356135373631333635613631343933343536356238313532363032303031363030313832303138303534363133363665393036313461393435363562383036303166303136303230383039313034303236303230303136303430353139303831303136303430353238303932393139303831383135323630323030313832383035343631333639613930363134613934353635623830313536313336653535373830363031663130363133366263353736313031303038303833353430343032383335323931363032303031393136313336653535363562383230313931393035663532363032303566323039303562383135343831353239303630303130313930363032303031383038333131363133366338353738323930303336303166313638323031393135623530353035303931383335323530353036303032383238313031353436303031363030313630343031623033383038323136363032303835303135323630303136303430316238323034383131363630343038353031353236303031363038303162383230343831313636303630383530313532363030313630633031623930393130343831313636303830383430313532363030333930393330313534393039323136363061303930393130313532393039313530383135313630303538313131313536313337353335373631333735333631343933343536356231343631333738363537356638343831353236303037383330313630323035323630343039303831393032303534393035313633313730636339333336306532316238313532363130363162393136306666313639303630303430313631346163363536356236303033383135323432363030313630303136303430316230333136363063303832303135323566383438313532363030373833303136303230353236303430393032303831353138313534383339323931393038323930363066663139313636303031383336303035383131313135363133376361353736313337636136313439333435363562303231373930353535303630323038323031353136303031383230313930363133376533393038323631346466343536356235303630343038323031353136303032383230313830353436303630383530313531363038303836303135313630613038373031353136303031363030313630343031623033393538363136363030313630303136303830316230333139393039343136393339303933313736303031363034303162393238363136393239303932303239313930393131373630303136303031363038303162303331363630303136303830316239313835313639313930393130323630303136303031363063303162303331363137363030313630633031623931383431363931393039313032313739303535363063303930393230313531363030333930393130313830353436303031363030313630343031623033313931363931393039323136313739303535356636313338383038353832363133326232353635623630383038343031353136303430383035313630303136303031363034303162303339303932313638323532343236303230383330313532393139333530383339323530383739313766313364353833393463663236396434386263663932373935396132396135666665653763393932346461666666383932376563646633633438666661376336373931303136303430353138303931303339306133353039333932353035303530353635623830343731303135363133386661353733303630343035313633636437383630353936306530316238313532363030343031363130363162393139303631343832373536356235663832363030313630303136306130316230333136383236303430353135663630343035313830383330333831383538373561663139323530353035303364383035663831313436313339343335373630343035313931353036303166313936303366336430313136383230313630343035323364383235323364356636303230383430313365363133393438353635623630363039313530356235303530393035303830363130363234353736303430353136333061313266353231363065313162383135323630303430313630343035313830393130333930666435623566363133393733363132313732353635623630303130313534363030313630303136303430316230333136393139303530353635623566363133393865363132313732353635623630303930313534363066663136363133396232353736303430353136333766616238316535363065303162383135323630303430313630343035313830393130333930666435623566363133396262363132313732353635623930353034323631333963653630363038363031363034303837303136313438393035363562363030313630303136303430316230333136313131353830363133613038353735303631333965633632303261333030343236313465633335363562363133396663363036303836303136303430383730313631343839303536356236303031363030313630343031623033313631303135356231353631336134323537363133613164363036303835303136303430383630313631343839303536356236303430353136333538373964613133363065313162383135323630303136303031363034303162303339303931313636303034383230313532363032343031363130363162353635623631336135373631336135323630363038363031383636313533353935363562363134323966353635623631336136373631336135323630383038363031383636313533353935363562363033303631336137363630323038363031383636313533366435363562393035303134363133616138353736313361386136303230383530313835363135333664353635623630343035313633323634373562326636306531316238313532363130363162393235303630303430313930383135323630323030313930353635623631336162323834383036313533366435363562393035303566303336313361646635373631336163343834383036313533366435363562363034303531363333653038613132353630653131623831353236303034303136313036316239323931393036313533616635363562356636303038383230313631336165653836383036313533366435363562363034303531363133616663393239313930363135316562353635623930383135323630323030313630343035313830393130333930323035343134363133623335353736313362316138343830363135333664353635623630343035313633613431663737326636306530316238313532363030343031363130363162393239313930363135336166353635623631336233663833356636313430313135363562363034303830353136306530383130313930393135323831353438313532356639303831393037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f3930363365303437623238333930363032303831303136313362376338613830363135333664353635623830383036303166303136303230383039313034303236303230303136303430353139303831303136303430353238303933393239313930383138313532363032303031383338333830383238343337356639323031393139303931353235303530353039303832353235303630323039303831303139303631336263343930386230313862363135333664353635623830383036303166303136303230383039313034303236303230303136303430353139303831303136303430353238303933393239313930383138313532363032303031383338333830383238343337356639323031393139303931353235303530353039303832353235303630323030313631336330643630363038623031363034303863303136313438393035363562363030313630303136303430316230333136383135323630323030313631336332383630363038623031386236313533353935363562363133633331393036313533633235363562383135323630323030313631336334333630383038623031386236313533353935363562363133633463393036313533633235363562383135323630323030313838363030313630303136303430316230333136383135323530363034303531383236336666666666666666313636306530316238313532363030343031363133633761393139303631353465343536356235663630343035313830383330333831383635616634313538303135363133633934353733643566383033653364356666643562353035303530353036303430353133643566383233653630316633643930383130313630316631393136383230313630343035323631336362623931393038313031393036313535396235363562356638323831353236303036383630313630323035323630343039303230393139333530393135303631336364393832383236313464663435363562353038313630303838343031363133636539383838303631353336643536356236303430353136313363663739323931393036313531656235363562393038313532363034303531393038313930303336303230303138313230393139303931353536336565356234386562363065303162383135323566393036303035363030313630393931623031393036336565356234386562393036313364333339303835393036303034303136313463303835363562363032303630343035313830383330333831356638373561663131353830313536313364346635373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363133643733393139303631346331613536356236303430383035313630653038313031393039313532393039313530383036303031383135323630323030313631336439333839383036313533366435363562383038303630316630313630323038303931303430323630323030313630343035313930383130313630343035323830393339323931393038313831353236303230303138333833383038323834333735663932303138323930353235303933383535323530353035303630303136303031363034303162303338393136363032303830383430313832393035323630343038303835303138343930353236303630383530313932393039323532363038303834303138333930353236306130393039333031383239303532383638323532363030373838303139303932353232303831353138313534383239303630666631393136363030313833363030353831313131353631336532323537363133653232363134393334353635623032313739303535353036303230383230313531363030313832303139303631336533623930383236313464663435363562353036303430383230313531363030323832303138303534363036303835303135313630383038363031353136306130383730313531363030313630303136303430316230333935383631363630303136303031363038303162303331393930393431363933393039333137363030313630343031623932383631363932393039323032393139303931313736303031363030313630383031623033313636303031363038303162393138353136393139303931303236303031363030313630633031623033313631373630303136306330316239313834313639313930393130323137393035353630633039303932303135313630303339303931303138303534363030313630303136303430316230333139313639313930393231363137393035353830363133656438383838303631353336643536356236303430353136313365653639323931393036313531656235363562363034303531383039313033393032303834376662373732393765336265666336393162666338363461383165323431663833653265663732326236653762656361613265636563323530633664353262343330383938623630343030313630323038313031393036313366323439313930363134383930353635623630343038303531363030313630303136303430316230333933383431363831353239323930393131363630323038333031353230313630343035313830393130333930613435303930393539343530353035303530353035363562376666306335376531363834306466303430663135303838646332663831666533393163333932336265633733653233613936363265666339633232396336613030353436303031363034303162393030343630666631363631336639623537363034303531363331616663643739663630653331623831353236303034303136303430353138303931303339306664356235363562363133666135363133663532353635623631336661653831363134343038353635623631336662363631343432313536356236313130306136303630383230313335363038303833303133353631336664333630633038353031363061303836303136313438393035363562363133666533363065303836303136306330383730313631353564653536356236313366663436313031303038373031363065303838303136313535663735363562363130313030383730313335363134303063363130313430383930313631303132303861303136313463343535363562363134343331353635623566363134303161363132313732353635623930353035663832363030313630303136303430316230333136383436303031363030313630343031623033313631313135363134303438353736313430343138333835363135323161353635623930353036313430353535363562363134303532383438343631353231613536356239303530356236303430383035313630383038313031383235323630303238343031353438303832353236303033383530313534363032303833303135323630303438353031353439323832303139323930393235323630303538343031353436303031363030313630343031623033313636303630383230313532343239313135383036313430623735373530363030313834303135343831353136313430623339313630303136303031363034303162303331363930363134656333353635623832313031353562313536313430646435373630303136303031363034303162303338333136363036303832303135323831383135323630343038313031353136303230383230313532363134306663353635623832383136303630303138313831353136313430656639313930363135316661353635623630303136303031363034303162303331363930353235303562363036303831303135313631343130633930363036343631353332653536356236303230383230313531363030313836303135343630303136303031363034303162303339323930393231363931363134313337393139303630303136303430316239303034363066663136363134656638353635623130313536313431363735373630363038313031353136303430353136336466616538383031363065303162383135323630303136303031363034303162303339303931313636303034383230313532363032343031363130363162353635623835363030313630303136303430316230333136383136303430303138313831353136313431383239313930363134656333353635623930353235303630343038313031383035313630303136303031363034303162303338373136393139303631343161323930383339303631353331623536356239303532353036303031383430313534363034303832303135313630363439313631343163373931363030313630343031623930393130343630666631363930363134656638353635623130313536313431656535373830363034303031353136303430353136333539343333313766363065303162383135323630303430313631303631623931383135323630323030313930353635623830353136303032383530313535363032303831303135313630303338353031353536303430383130313531363030343835303135353630363030313531363030353930393330313830353436303031363030313630343031623033313931363630303136303031363034303162303339303934313639333930393331373930393235353530353035303530353635623566383036313432343036313231373235363562356638343831353236303037383230313630323035323630343039303230363030323031383035343931393235303930363030383930363134323734393036303031363034303162393030343630303136303031363034303162303331363631353631373536356239313930363130313030306138313534383136303031363030313630343031623033303231393136393038333630303136303031363034303162303331363032313739303535393135303530393139303530353635623631343261633630323038323031383236313438336235363562363366666666666666663136313538303135363134326363353735303631343263373630323038323031383236313463363035363562313531353930353035623135363134333133353736313432646536303230383230313832363134383362353635623631343265623630323038333031383336313463363035363562363034303531363363303861306631643630653031623831353236336666666666666666393039333136363030343834303135323630323438333031353235303630343430313631303631623536356236313433323036303230383230313832363134633630353635623930353036313433326636303230383330313833363134383362353635623633666666666666666631363131313536313433343835373631343264653630323038323031383236313438336235363562363030313562363134333538363032303833303138333631346336303536356239303530383131303135363131363139353736313433366536303230383330313833363134633630353635623631343337393630303138343631353331623536356238313831313036313433383835373631343338383631346361353536356239303530363032303032303136303230383130313930363134333964393139303631346334353536356236303031363030313630613031623033313636313433623336303230383430313834363134633630353635623833383138313130363134336333353736313433633336313463613535363562393035303630323030323031363032303831303139303631343364383931393036313463343535363562363030313630303136306130316230333136313031353631343430303537363034303531363330646263386435663630653331623831353236303034303136303430353138303931303339306664356236303031303136313433346235363562363134343130363133663532353635623631343431383631343563373536356236313130306138313631343563663536356236313434323936313366353235363562363133663962363134366232353635623631343433393631336635323536356235663631343434323631323538343536356239303530363166666666383531363135383036313434356135373530363132373130363166666666383631363131356231353631343437653537363034303531363335663132653663333630653131623831353236316666666638363136363030343832303135323630323430313631303631623536356238363838313131353631343461323537363034303531363332323264313634333630653231623831353236303034383130313839393035323630323430313631303631623536356236306666383431363135383036313434623535373530363030613630666638353136313135623135363134346438353736303430353136333137306462333539363065333162383135323630666638353136363030343832303135323630323430313631303631623536356236313434653036313339366135363562363030313630303136303430316230333136383636303031363030313630343031623033313631303135363134353163353736303430353136323032613036643630653131623831353236303031363030313630343031623033383731363630303438323031353236303234303136313036316235363562383235663033363134353363353736303430353136336137333330303731363065303162383135323630303430313630343035313830393130333930666435623936383735353630303138373031393539303935353536303032383630313830353436303031363030313630343031623033393539303935313636396666666666666666666666666666666666666666313939303935313639343930393431373630303136303430316236316666666639343930393431363933393039333032393239303932313736376666666666666666666666666666666636303530316231393136363066663931393039313136363030313630353031623032313739303931353536303033383330313535363030343930393130313830353436303031363030313630613031623033313931363630303136303031363061303162303339303932313639313930393131373930353535363562363133663962363133663532353635623631343564373631336635323536356235663631343565303631323137323536356238323335383135353930353036303134363134356638363036303834303136303430383530313631353566373536356236306666313631313830363134363137353735303631343631323630363038333031363034303834303136313535663735363562363066663136313535623135363134363462353736313436326336303630383330313630343038343031363135356637353635623630343035313633346135396262666636306531316238313532363066663930393131363630303438323031353236303234303136313036316235363562363134363562363036303833303136303430383430313631353566373536356236303031383230313830353436306666393239303932313636303031363034303162303236306666363034303162313939303932313639313930393131373930353536313436386336303430383330313630323038343031363134383930353635623630303139313930393130313830353436303031363030313630343031623033313931363630303136303031363034303162303339303932313639313930393131373930353535303536356236313262646436313366353235363562353038303534363134366336393036313461393435363562356638323535383036303166313036313436643535373530353035363562363031663031363032303930303439303566353236303230356632303930383130313930363131303061393139303631343732653536356236303430383035313630653038313031393039313532383035663831353236303630363032303832303138313930353235663630343038333031383139303532393038323031383139303532363038303832303138313930353236306130383230313831393035323630633039303931303135323930353635623562383038323131313536313437343235373566383135353630303130313631343732663536356235303930353635623830313531353831313436313130306135373566383066643562383033353633666666666666666638313136383131343631316361643537356638306664356235663830356636303630383438363033313231353631343737383537356638306664356238333335393235303630323038343031333536313437386138313631343734363536356239313530363134373938363034303835303136313437353335363562393035303932353039323530393235363562356636303230383238343033313231353631343762313537356638306664356235303335393139303530353635623566383036303430383338353033313231353631343763393537356638306664356238323335363030313630303136303430316230333831313131353631343764653537356638306664356238333031363038303831383630333132313536313437656635373566383066643562393135303631343766643630323038343031363134373533353635623930353039323530393239303530353635623566383036303430383338353033313231353631343831373537356638306664356238323335393135303631343766643630323038343031363134373533353635623630303136303031363061303162303339313930393131363831353236303230303139303536356235663630323038323834303331323135363134383462353735663830666435623631313832653832363134373533353635623566383036303430383338353033313231353631343836353537356638306664356236313438366538333631343735333536356239343630323039333930393330313335393335303530353035363562363030313630303136303430316230333831313638313134363131303061353735663830666435623566363032303832383430333132313536313438613035373566383066643562383133353631313832653831363134383763353635623830333536316666666638313136383131343631316361643537356638306664356235663830356636303630383438363033313231353631343863653537356638306664356238333335363030313630303136303430316230333831313131353631343865333537356638306664356238343031363061303831383730333132313536313438663435373566383066643562393235303631343930323630323038353031363134386162353635623931353036303430383430313335363134393132383136313438376335363562383039313530353039323530393235303932353635623566363130313430383238343033313231353631343932653537356638306664356235303931393035303536356236333465343837623731363065303162356635323630323136303034353236303234356666643562363030363831313036313439353835373631343935383631343933343536356239303532353635623566356238333831313031353631343937363537383138313031353138333832303135323630323030313631343935653536356235303530356639313031353235363562356638313531383038343532363134393935383136303230383630313630323038363031363134393563353635623630316630313630316631393136393239303932303136303230303139323931353035303536356236303230383135323631343962623630323038323031383335313631343934383536356235663630323038333031353136306530363034303834303135323631343964363631303130303834303138323631343937653536356239303530363034303834303135313630303136303031363034303162303338303832313636303630383630313532383036303630383730313531313636303830383630313532383036303830383730313531313636306130383630313532383036306130383730313531313636306330383630313532383036306330383730313531313636306530383630313532353035303830393135303530393239313530353035363562356638303630323038333835303331323135363134613361353735663830666435623832333536303031363030313630343031623033383038323131313536313461353035373566383066643562383138353031393135303835363031663833303131323631346136333537356638306664356238313335383138313131313536313461373135373566383066643562383636303230383238353031303131313135363134613832353735663830666435623630323039323930393230313936393139353530393039333530353035303530353635623630303138313831316339303832313638303631346161383537363037663832313639313530356236303230383231303831303336313439326535373633346534383762373136306530316235663532363032323630303435323630323435666664356236303230383130313631306634633832383436313439343835363562363334653438376237313630653031623566353236303431363030343532363032343566666435623630343035313630363038313031363030313630303136303430316230333831313138323832313031373135363134623061353736313462306136313461643435363562363034303532393035363562363034303830353139303831303136303031363030313630343031623033383131313832383231303137313536313462306135373631346230613631346164343536356236303430353136303166383230313630316631393136383130313630303136303031363034303162303338313131383238323130313731353631346235613537363134623561363134616434353635623630343035323931393035303536356235663630303136303031363034303162303338323131313536313462376135373631346237613631346164343536356235303630316630313630316631393136363032303031393035363562356638323630316638333031313236313462393735373566383066643562383135313631346261613631346261353832363134623632353635623631346233323536356238313831353238343630323038333836303130313131313536313462626535373566383066643562363134626366383236303230383330313630323038373031363134393563353635623934393335303530353035303536356235663630323038323834303331323135363134626537353735663830666435623831353136303031363030313630343031623033383131313135363134626663353735663830666435623631346263663834383238353031363134623838353635623630323038313532356636313138326536303230383330313834363134393765353635623566363032303832383430333132313536313463326135373566383066643562353035313931393035303536356236303031363030313630613031623033383131363831313436313130306135373566383066643562356636303230383238343033313231353631346335353537356638306664356238313335363131383265383136313463333135363562356638303833333536303165313938343336303330313831313236313463373535373566383066643562383330313830333539313530363030313630303136303430316230333832313131353631346338653537356638306664356236303230303139313530363030353831393031623336303338323133313536313334373535373566383066643562363334653438376237313630653031623566353236303332363030343532363032343566666435623566383233353630356531393833333630333031383131323631346363643537356638306664356239313930393130313932393135303530353635623566383236303166383330313132363134636536353735663830666435623831333536313463663436313462613538323631346236323536356238313831353238343630323038333836303130313131313536313464303835373566383066643562383136303230383530313630323038333031333735663931383130313630323030313931393039313532393339323530353035303536356235663630363038323336303331323135363134643334353735663830666435623631346433633631346165383536356238323335363030313630303136303430316230333830383231313135363134643532353735663830666435623631346435653336383338373031363134636437353635623833353236303230383530313335393135303830383231313135363134643733353735663830666435623530363134643830333638323836303136313463643735363562363032303833303135323530363034303833303133353631346439343831363134383763353635623630343038323031353239323931353035303536356235663832353136313463636438313834363032303837303136313439356335363562363031663832313131353631303632343537383035663532363032303566323036303166383430313630303531633831303136303230383531303135363134646435353735303830356236303166383430313630303531633832303139313530356238313831313031353631316138653537356638313535363030313031363134646531353635623831353136303031363030313630343031623033383131313135363134653064353736313465306436313461643435363562363134653231383136313465316238343534363134613934353635623834363134646230353635623630323038303630316638333131363030313831313436313465353435373566383431353631346533643537353038353833303135313562356631393630303338363930316231633139313636303031383539303162313738353535363131313331353635623566383538313532363032303831323036303166313938363136393135623832383131303135363134653832353738383836303135313832353539343834303139343630303139303931303139303834303136313465363335363562353038353832313031353631346539663537383738353031353135663139363030333838393031623630663831363163313931363831353535623530353035303530353036303031393038313162303139303535353035363562363334653438376237313630653031623566353236303131363030343532363032343566666435623830383230313830383231313135363130663463353736313066346336313465616635363562356636336666666666666666383038333136383138313033363134656565353736313465656536313465616635363562363030313031393339323530353035303536356238303832303238313135383238323034383431343137363130663463353736313066346336313465616635363562356638303833333536303165313938343336303330313831313236313466323435373566383066643562383330313630323038313031393235303335393035303630303136303031363034303162303338313131313536313466343235373566383066643562383033363033383231333135363133343735353735663830666435623831383335323831383136303230383530313337353035663832383230313630323039303831303139313930393135323630316639303931303136303166313931363930393130313031393035363562356638333833383535323630323038303836303139353530383038353630303531623833303130313834356635623837383131303135363135303264353738343833303336303166313930313839353238313335333638383930303336303565313930313831313236313466623435373566383066643562383730313630363036313466633238323830363134663066353635623832383735323631346664323833383830313832383436313466353035363562393235303530353036313466653238363833303138333631346630663536356238363833303338383838303135323631346666343833383238343631346635303536356239323530353035303630343038303833303133353932353036313530303938333631343837633536356236303031363030313630343031623033393239303932313639343930393130313933393039333532393738333031393739303833303139303630303130313631346638663536356235303930393739363530353035303530353035303530353635623630323038313532383133353630323038323031353236303230383230313335363034303832303135323566363034303833303133353631353035653831363134633331353635623630303136303031363061303162303331363630363038333831303139313930393135323833303133353336383439303033363031653139303138313132363135303835353735663830666435623833303136303230383130313930333536303031363030313630343031623033383131313135363135306130353735663830666435623830363030353162333630333832313331353631353062313537356638306664356236303830383038353031353236313335663636306130383530313832383436313466373835363562356638323631353064663537363334653438376237313630653031623566353236303132363030343532363032343566666435623530303439303536356236303230383130313630303438333130363135306638353736313530663836313439333435363562393139303532393035363562356638303566363036303834383630333132313536313531313035373566383066643562383335313932353036303230383430313531363135313232383136313438376335363562363034303835303135313930393235303631343931323831363134383763353635623566383036303430383338353033313231353631353134343537356638306664356238323531393135303630323038333031353136313531353638313631343734363536356238303931353035303932353039323930353035363562356636303230383038333532356638343534363135313733383136313461393435363562383036303230383730313532363034303630303138303834313635663831313436313531393435373630303138313134363135316230353736313531646435363562363066663139383531363630343038613031353236303430383431353135363030353162386130313031393535303631353164643536356238393566353236303230356632303566356238353831313031353631353164343537383135343862383230313836303135323930383330313930383830313631353162393536356238613031363034303031393635303530356235303933393839373530353035303530353035303530353035363562383138333832333735663931303139303831353239313930353035363562363030313630303136303430316230333831383131363833383231363031393038303832313131353631323239613537363132323961363134656166353635623630303136303031363034303162303338323831313638323832313630333930383038323131313536313232396135373631323239613631346561663536356235663830363034303833383530333132313536313532346235373566383066643562383235313630303136303031363034303162303338303832313131353631353236313537356638306664356239303834303139303630363038323837303331323135363135323734353735663830666435623631353237633631346165383536356238323531383135323630323038333031353136313532386538313631346333313536356236303230383230313532363034303833303135313832383131313135363135326134353735663830666435623631353262303838383238363031363134623838353635623630343038333031353235303830393435303530353035303630323038333031353136313531353638313631343734363536356235663830363034303833383530333132313536313532646235373566383066643562383235313931353036303230383330313531363135313536383136313438376335363562393438353532363030313630303136303430316230333933383431363630323038363031353239313833313636303430383530313532383231363630363038343031353231363630383038323031353236306130303139303536356238313831303338313831313131353631306634633537363130663463363134656166353635623630303136303031363034303162303338313831313638333832313630323830383231363931393038323831313436313533353135373631353335313631346561663536356235303530393239313530353035363562356638323335363033653139383333363033303138313132363134636364353735663830666435623566383038333335363031653139383433363033303138313132363135333832353735663830666435623833303138303335393135303630303136303031363034303162303338323131313536313533396235373566383066643562363032303031393135303336383139303033383231333135363133343735353735663830666435623630323038313532356636313462636636303230383330313834383636313466353035363562356636303430383233363033313231353631353364323537356638306664356236313533646136313462313035363562363135336533383336313437353335363562383135323630323038303834303133353630303136303031363034303162303338303832313131353631353366663537356638306664356239303835303139303336363031663833303131323631353431313537356638306664356238313335383138313131313536313534323335373631353432333631346164343536356238303630303531623931353036313534333438343833303136313462333235363562383138313532393138333031383430313931383438313031393033363834313131353631353434643537356638306664356239333835303139333562383338353130313536313534373735373834333539323530363135343637383336313463333135363562383238323532393338353031393339303835303139303631353435323536356239343836303139343930393435323530393239353934353035303530353035303536356235663630343038333031363366666666666666663833353131363834353236303230383038343031353136303430363032303837303135323832383135313830383535323630363038383031393135303630323038333031393435303566393235303562383038333130313536313235373935373834353136303031363030313630613031623033313638323532393338333031393336303031393239303932303139313930383330313930363135346262353635623630323038313532383135313630323038323031353235663630323038333031353136306530363034303834303135323631353530613631303130303834303138323631343937653536356239303530363034303834303135313630316631393830383538343033303136303630383630313532363135353238383338333631343937653536356239323530363030313630303136303430316230333630363038373031353131363630383038363031353236303830383630313531393135303830383538343033303136306130383630313532363135353538383338333631353438393536356239323530363061303836303135313931353038303835383430333031363063303836303135323530363135353736383238323631353438393536356239313530353036306330383430313531363135353933363065303835303138323630303136303031363034303162303331363930353235363562353039333932353035303530353635623566383036303430383338353033313231353631353561633537356638306664356238323531393135303630323038333031353136303031363030313630343031623033383131313135363135356338353735663830666435623631353564343835383238363031363134623838353635623931353035303932353039323930353035363562356636303230383238343033313231353631353565653537356638306664356236313138326538323631343861623536356235663630323038323834303331323135363135363037353735663830666435623831333536306666383131363831313436313138326535373566383066643562356636303031363030313630343031623033383038333136383138313033363134656565353736313465656536313465616635366665613136343733366636633633343330303038313930303061",
}

// NativeTokenStakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenStakingManagerMetaData.ABI instead.
var NativeTokenStakingManagerABI = NativeTokenStakingManagerMetaData.ABI

// NativeTokenStakingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenStakingManagerMetaData.Bin instead.
var NativeTokenStakingManagerBin = NativeTokenStakingManagerMetaData.Bin

// DeployNativeTokenStakingManager deploys a new Ethereum contract, binding an instance of NativeTokenStakingManager to it.
func DeployNativeTokenStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *NativeTokenStakingManager, error) {
	parsed, err := NativeTokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenStakingManagerBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenStakingManager{NativeTokenStakingManagerCaller: NativeTokenStakingManagerCaller{contract: contract}, NativeTokenStakingManagerTransactor: NativeTokenStakingManagerTransactor{contract: contract}, NativeTokenStakingManagerFilterer: NativeTokenStakingManagerFilterer{contract: contract}}, nil
}

// NativeTokenStakingManager is an auto generated Go binding around an Ethereum contract.
type NativeTokenStakingManager struct {
	NativeTokenStakingManagerCaller     // Read-only binding to the contract
	NativeTokenStakingManagerTransactor // Write-only binding to the contract
	NativeTokenStakingManagerFilterer   // Log filterer for contract events
}

// NativeTokenStakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenStakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenStakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenStakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenStakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenStakingManagerSession struct {
	Contract     *NativeTokenStakingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NativeTokenStakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenStakingManagerCallerSession struct {
	Contract *NativeTokenStakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// NativeTokenStakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenStakingManagerTransactorSession struct {
	Contract     *NativeTokenStakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// NativeTokenStakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenStakingManagerRaw struct {
	Contract *NativeTokenStakingManager // Generic contract binding to access the raw methods on
}

// NativeTokenStakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerCallerRaw struct {
	Contract *NativeTokenStakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenStakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenStakingManagerTransactorRaw struct {
	Contract *NativeTokenStakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenStakingManager creates a new instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManager(address common.Address, backend bind.ContractBackend) (*NativeTokenStakingManager, error) {
	contract, err := bindNativeTokenStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManager{NativeTokenStakingManagerCaller: NativeTokenStakingManagerCaller{contract: contract}, NativeTokenStakingManagerTransactor: NativeTokenStakingManagerTransactor{contract: contract}, NativeTokenStakingManagerFilterer: NativeTokenStakingManagerFilterer{contract: contract}}, nil
}

// NewNativeTokenStakingManagerCaller creates a new read-only instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenStakingManagerCaller, error) {
	contract, err := bindNativeTokenStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerCaller{contract: contract}, nil
}

// NewNativeTokenStakingManagerTransactor creates a new write-only instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenStakingManagerTransactor, error) {
	contract, err := bindNativeTokenStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerTransactor{contract: contract}, nil
}

// NewNativeTokenStakingManagerFilterer creates a new log filterer instance of NativeTokenStakingManager, bound to a specific deployed contract.
func NewNativeTokenStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenStakingManagerFilterer, error) {
	contract, err := bindNativeTokenStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerFilterer{contract: contract}, nil
}

// bindNativeTokenStakingManager binds a generic wrapper to an already deployed contract.
func bindNativeTokenStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenStakingManager *NativeTokenStakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenStakingManager.Contract.NativeTokenStakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenStakingManager *NativeTokenStakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.NativeTokenStakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenStakingManager *NativeTokenStakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.NativeTokenStakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenStakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) ADDRESSLENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "ADDRESS_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ADDRESSLENGTH() (uint32, error) {
	return _NativeTokenStakingManager.Contract.ADDRESSLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) ADDRESSLENGTH() (uint32, error) {
	return _NativeTokenStakingManager.Contract.ADDRESSLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) BIPSCONVERSIONFACTOR(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "BIPS_CONVERSION_FACTOR")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _NativeTokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_NativeTokenStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _NativeTokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_NativeTokenStakingManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) BLSPUBLICKEYLENGTH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "BLS_PUBLIC_KEY_LENGTH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _NativeTokenStakingManager.Contract.BLSPUBLICKEYLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _NativeTokenStakingManager.Contract.BLSPUBLICKEYLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMCHURNPERCENTAGELIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_CHURN_PERCENTAGE_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMDELEGATIONFEEBIPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_DELEGATION_FEE_BIPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMREGISTRATIONEXPIRYLENGTH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_REGISTRATION_EXPIRY_LENGTH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) MAXIMUMSTAKEMULTIPLIERLIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "MAXIMUM_STAKE_MULTIPLIER_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_NativeTokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _NativeTokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_NativeTokenStakingManager.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) NATIVEMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "NATIVE_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.NATIVEMINTER(&_NativeTokenStakingManager.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.NATIVEMINTER(&_NativeTokenStakingManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) POSVALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "POS_VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) PCHAINBLOCKCHAINID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "P_CHAIN_BLOCKCHAIN_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.PCHAINBLOCKCHAINID(&_NativeTokenStakingManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.PCHAINBLOCKCHAINID(&_NativeTokenStakingManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) VALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_NativeTokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) WARPMESSENGER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.WARPMESSENGER(&_NativeTokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _NativeTokenStakingManager.Contract.WARPMESSENGER(&_NativeTokenStakingManager.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) GetValidator(opts *bind.CallOpts, validationID [32]byte) (Validator, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "getValidator", validationID)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _NativeTokenStakingManager.Contract.GetValidator(&_NativeTokenStakingManager.CallOpts, validationID)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _NativeTokenStakingManager.Contract.GetValidator(&_NativeTokenStakingManager.CallOpts, validationID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) GetWeight(opts *bind.CallOpts, validationID [32]byte) (uint64, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "getWeight", validationID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _NativeTokenStakingManager.Contract.GetWeight(&_NativeTokenStakingManager.CallOpts, validationID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _NativeTokenStakingManager.Contract.GetWeight(&_NativeTokenStakingManager.CallOpts, validationID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) RegisteredValidators(opts *bind.CallOpts, nodeID []byte) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "registeredValidators", nodeID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.RegisteredValidators(&_NativeTokenStakingManager.CallOpts, nodeID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _NativeTokenStakingManager.Contract.RegisteredValidators(&_NativeTokenStakingManager.CallOpts, nodeID)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) ValueToWeight(opts *bind.CallOpts, value *big.Int) (uint64, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "valueToWeight", value)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _NativeTokenStakingManager.Contract.ValueToWeight(&_NativeTokenStakingManager.CallOpts, value)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _NativeTokenStakingManager.Contract.ValueToWeight(&_NativeTokenStakingManager.CallOpts, value)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) WeightToValue(opts *bind.CallOpts, weight uint64) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "weightToValue", weight)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _NativeTokenStakingManager.Contract.WeightToValue(&_NativeTokenStakingManager.CallOpts, weight)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _NativeTokenStakingManager.Contract.WeightToValue(&_NativeTokenStakingManager.CallOpts, weight)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ClaimDelegationFees(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ClaimDelegationFees(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x5297fae6.
//
// Solidity: function completeDelegatorRegistration(uint32 messageIndex, bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, messageIndex uint32, delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeDelegatorRegistration", messageIndex, delegationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x5297fae6.
//
// Solidity: function completeDelegatorRegistration(uint32 messageIndex, bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteDelegatorRegistration(messageIndex uint32, delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, messageIndex, delegationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x5297fae6.
//
// Solidity: function completeDelegatorRegistration(uint32 messageIndex, bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteDelegatorRegistration(messageIndex uint32, delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, messageIndex, delegationID)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x98f3e2b4.
//
// Solidity: function completeEndDelegation(uint32 messageIndex, bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteEndDelegation(opts *bind.TransactOpts, messageIndex uint32, delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeEndDelegation", messageIndex, delegationID)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x98f3e2b4.
//
// Solidity: function completeEndDelegation(uint32 messageIndex, bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteEndDelegation(messageIndex uint32, delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteEndDelegation(&_NativeTokenStakingManager.TransactOpts, messageIndex, delegationID)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x98f3e2b4.
//
// Solidity: function completeEndDelegation(uint32 messageIndex, bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteEndDelegation(messageIndex uint32, delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteEndDelegation(&_NativeTokenStakingManager.TransactOpts, messageIndex, delegationID)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteEndValidation(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeEndValidation", messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteEndValidation(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteEndValidation(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ForceInitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "forceInitializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ForceInitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "forceInitializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndValidation(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ForceInitializeEndValidation(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0xc257a0f5.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) Initialize(opts *bind.TransactOpts, settings PoSValidatorManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initialize", settings)
}

// Initialize is a paid mutator transaction binding the contract method 0xc257a0f5.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) Initialize(settings PoSValidatorManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.Initialize(&_NativeTokenStakingManager.TransactOpts, settings)
}

// Initialize is a paid mutator transaction binding the contract method 0xc257a0f5.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) Initialize(settings PoSValidatorManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.Initialize(&_NativeTokenStakingManager.TransactOpts, settings)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0xc599e24f.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeDelegatorRegistration", validationID)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0xc599e24f.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeDelegatorRegistration(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0xc599e24f.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeDelegatorRegistration(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndValidation(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeEndValidation(&_NativeTokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xaf2f5feb.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeValidatorRegistration(opts *bind.TransactOpts, registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeValidatorRegistration", registrationInput, delegationFeeBips, minStakeDuration)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xaf2f5feb.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0xaf2f5feb.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration) payable returns(bytes32)
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeValidatorRegistration(&_NativeTokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) InitializeValidatorSet(opts *bind.TransactOpts, subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initializeValidatorSet", subnetConversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeValidatorSet(&_NativeTokenStakingManager.TransactOpts, subnetConversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.InitializeValidatorSet(&_NativeTokenStakingManager.TransactOpts, subnetConversionData, messageIndex)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ResendEndValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "resendEndValidatorMessage", validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendEndValidatorMessage(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendEndValidatorMessage(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ResendRegisterValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "resendRegisterValidatorMessage", validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendRegisterValidatorMessage(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendRegisterValidatorMessage(&_NativeTokenStakingManager.TransactOpts, validationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ResendUpdateDelegation(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "resendUpdateDelegation", delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendUpdateDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ResendUpdateDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.SubmitUptimeProof(&_NativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.SubmitUptimeProof(&_NativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// NativeTokenStakingManagerDelegationEndedIterator is returned from FilterDelegationEnded and is used to iterate over the raw logs and unpacked data for DelegationEnded events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegationEndedIterator struct {
	Event *NativeTokenStakingManagerDelegationEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerDelegationEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegationEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerDelegationEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerDelegationEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegationEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegationEnded represents a DelegationEnded event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegationEnded struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegationEnded is a free log retrieval operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegationEnded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenStakingManagerDelegationEndedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegationEndedIterator{contract: _NativeTokenStakingManager.contract, event: "DelegationEnded", logs: logs, sub: sub}, nil
}

// WatchDelegationEnded is a free log subscription operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegationEnded(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegationEnded, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegationEnded)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegationEnded is a log parse operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegationEnded(log types.Log) (*NativeTokenStakingManagerDelegationEnded, error) {
	event := new(NativeTokenStakingManagerDelegationEnded)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerDelegatorAddedIterator is returned from FilterDelegatorAdded and is used to iterate over the raw logs and unpacked data for DelegatorAdded events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorAddedIterator struct {
	Event *NativeTokenStakingManagerDelegatorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerDelegatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegatorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerDelegatorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerDelegatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegatorAdded represents a DelegatorAdded event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorAdded struct {
	DelegationID       [32]byte
	ValidationID       [32]byte
	DelegatorAddress   common.Address
	Nonce              uint64
	ValidatorWeight    uint64
	DelegatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDelegatorAdded is a free log retrieval operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegatorAdded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*NativeTokenStakingManagerDelegatorAddedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegatorAddedIterator{contract: _NativeTokenStakingManager.contract, event: "DelegatorAdded", logs: logs, sub: sub}, nil
}

// WatchDelegatorAdded is a free log subscription operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegatorAdded(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegatorAdded, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegatorAdded)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorAdded is a log parse operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegatorAdded(log types.Log) (*NativeTokenStakingManagerDelegatorAdded, error) {
	event := new(NativeTokenStakingManagerDelegatorAdded)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerDelegatorRegisteredIterator is returned from FilterDelegatorRegistered and is used to iterate over the raw logs and unpacked data for DelegatorRegistered events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRegisteredIterator struct {
	Event *NativeTokenStakingManagerDelegatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerDelegatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerDelegatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerDelegatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegatorRegistered represents a DelegatorRegistered event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRegistered struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRegistered is a free log retrieval operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegatorRegistered(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenStakingManagerDelegatorRegisteredIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegatorRegisteredIterator{contract: _NativeTokenStakingManager.contract, event: "DelegatorRegistered", logs: logs, sub: sub}, nil
}

// WatchDelegatorRegistered is a free log subscription operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegatorRegistered(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegatorRegistered, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegatorRegistered)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorRegistered is a log parse operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegatorRegistered(log types.Log) (*NativeTokenStakingManagerDelegatorRegistered, error) {
	event := new(NativeTokenStakingManagerDelegatorRegistered)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerDelegatorRemovalInitializedIterator is returned from FilterDelegatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for DelegatorRemovalInitialized events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRemovalInitializedIterator struct {
	Event *NativeTokenStakingManagerDelegatorRemovalInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerDelegatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerDelegatorRemovalInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerDelegatorRemovalInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerDelegatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerDelegatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerDelegatorRemovalInitialized represents a DelegatorRemovalInitialized event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerDelegatorRemovalInitialized struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRemovalInitialized is a free log retrieval operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterDelegatorRemovalInitialized(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenStakingManagerDelegatorRemovalInitializedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerDelegatorRemovalInitializedIterator{contract: _NativeTokenStakingManager.contract, event: "DelegatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchDelegatorRemovalInitialized is a free log subscription operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchDelegatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerDelegatorRemovalInitialized, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerDelegatorRemovalInitialized)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegatorRemovalInitialized is a log parse operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseDelegatorRemovalInitialized(log types.Log) (*NativeTokenStakingManagerDelegatorRemovalInitialized, error) {
	event := new(NativeTokenStakingManagerDelegatorRemovalInitialized)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerInitialValidatorCreatedIterator is returned from FilterInitialValidatorCreated and is used to iterate over the raw logs and unpacked data for InitialValidatorCreated events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitialValidatorCreatedIterator struct {
	Event *NativeTokenStakingManagerInitialValidatorCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerInitialValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerInitialValidatorCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerInitialValidatorCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerInitialValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerInitialValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerInitialValidatorCreated represents a InitialValidatorCreated event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitialValidatorCreated struct {
	ValidationID [32]byte
	NodeID       common.Hash
	Weight       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialValidatorCreated is a free log retrieval operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterInitialValidatorCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte) (*NativeTokenStakingManagerInitialValidatorCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerInitialValidatorCreatedIterator{contract: _NativeTokenStakingManager.contract, event: "InitialValidatorCreated", logs: logs, sub: sub}, nil
}

// WatchInitialValidatorCreated is a free log subscription operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchInitialValidatorCreated(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerInitialValidatorCreated, validationID [][32]byte, nodeID [][]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerInitialValidatorCreated)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialValidatorCreated is a log parse operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseInitialValidatorCreated(log types.Log) (*NativeTokenStakingManagerInitialValidatorCreated, error) {
	event := new(NativeTokenStakingManagerInitialValidatorCreated)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitializedIterator struct {
	Event *NativeTokenStakingManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerInitialized represents a Initialized event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeTokenStakingManagerInitializedIterator, error) {

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerInitializedIterator{contract: _NativeTokenStakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerInitialized)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseInitialized(log types.Log) (*NativeTokenStakingManagerInitialized, error) {
	event := new(NativeTokenStakingManagerInitialized)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerUptimeUpdatedIterator struct {
	Event *NativeTokenStakingManagerUptimeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerUptimeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerUptimeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerUptimeUpdated represents a UptimeUpdated event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*NativeTokenStakingManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerUptimeUpdatedIterator{contract: _NativeTokenStakingManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerUptimeUpdated)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUptimeUpdated is a log parse operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseUptimeUpdated(log types.Log) (*NativeTokenStakingManagerUptimeUpdated, error) {
	event := new(NativeTokenStakingManagerUptimeUpdated)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerValidationPeriodCreatedIterator is returned from FilterValidationPeriodCreated and is used to iterate over the raw logs and unpacked data for ValidationPeriodCreated events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationPeriodCreatedIterator struct {
	Event *NativeTokenStakingManagerValidationPeriodCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerValidationPeriodCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidationPeriodCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerValidationPeriodCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerValidationPeriodCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidationPeriodCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidationPeriodCreated represents a ValidationPeriodCreated event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationPeriodCreated struct {
	ValidationID                [32]byte
	NodeID                      common.Hash
	RegisterValidationMessageID [32]byte
	Weight                      *big.Int
	RegistrationExpiry          uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodCreated is a free log retrieval operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidationPeriodCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (*NativeTokenStakingManagerValidationPeriodCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidationPeriodCreatedIterator{contract: _NativeTokenStakingManager.contract, event: "ValidationPeriodCreated", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodCreated is a free log subscription operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidationPeriodCreated(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidationPeriodCreated, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidationPeriodCreated)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationPeriodCreated is a log parse operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidationPeriodCreated(log types.Log) (*NativeTokenStakingManagerValidationPeriodCreated, error) {
	event := new(NativeTokenStakingManagerValidationPeriodCreated)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerValidationPeriodEndedIterator is returned from FilterValidationPeriodEnded and is used to iterate over the raw logs and unpacked data for ValidationPeriodEnded events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationPeriodEndedIterator struct {
	Event *NativeTokenStakingManagerValidationPeriodEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerValidationPeriodEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidationPeriodEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerValidationPeriodEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerValidationPeriodEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidationPeriodEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidationPeriodEnded represents a ValidationPeriodEnded event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationPeriodEnded struct {
	ValidationID [32]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodEnded is a free log retrieval operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidationPeriodEnded(opts *bind.FilterOpts, validationID [][32]byte, status []uint8) (*NativeTokenStakingManagerValidationPeriodEndedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidationPeriodEndedIterator{contract: _NativeTokenStakingManager.contract, event: "ValidationPeriodEnded", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodEnded is a free log subscription operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidationPeriodEnded(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidationPeriodEnded, validationID [][32]byte, status []uint8) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidationPeriodEnded)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationPeriodEnded is a log parse operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidationPeriodEnded(log types.Log) (*NativeTokenStakingManagerValidationPeriodEnded, error) {
	event := new(NativeTokenStakingManagerValidationPeriodEnded)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerValidationPeriodRegisteredIterator is returned from FilterValidationPeriodRegistered and is used to iterate over the raw logs and unpacked data for ValidationPeriodRegistered events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationPeriodRegisteredIterator struct {
	Event *NativeTokenStakingManagerValidationPeriodRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerValidationPeriodRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidationPeriodRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerValidationPeriodRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerValidationPeriodRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidationPeriodRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidationPeriodRegistered represents a ValidationPeriodRegistered event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationPeriodRegistered struct {
	ValidationID [32]byte
	Weight       *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodRegistered is a free log retrieval operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidationPeriodRegistered(opts *bind.FilterOpts, validationID [][32]byte) (*NativeTokenStakingManagerValidationPeriodRegisteredIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidationPeriodRegisteredIterator{contract: _NativeTokenStakingManager.contract, event: "ValidationPeriodRegistered", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodRegistered is a free log subscription operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidationPeriodRegistered(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidationPeriodRegistered, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidationPeriodRegistered)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationPeriodRegistered is a log parse operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidationPeriodRegistered(log types.Log) (*NativeTokenStakingManagerValidationPeriodRegistered, error) {
	event := new(NativeTokenStakingManagerValidationPeriodRegistered)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerValidatorRemovalInitializedIterator is returned from FilterValidatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for ValidatorRemovalInitialized events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidatorRemovalInitializedIterator struct {
	Event *NativeTokenStakingManagerValidatorRemovalInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerValidatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidatorRemovalInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerValidatorRemovalInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerValidatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidatorRemovalInitialized represents a ValidatorRemovalInitialized event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidatorRemovalInitialized struct {
	ValidationID       [32]byte
	SetWeightMessageID [32]byte
	Weight             *big.Int
	EndTime            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemovalInitialized is a free log retrieval operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidatorRemovalInitialized(opts *bind.FilterOpts, validationID [][32]byte, setWeightMessageID [][32]byte) (*NativeTokenStakingManagerValidatorRemovalInitializedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidatorRemovalInitializedIterator{contract: _NativeTokenStakingManager.contract, event: "ValidatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchValidatorRemovalInitialized is a free log subscription operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidatorRemovalInitialized, validationID [][32]byte, setWeightMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidatorRemovalInitialized)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRemovalInitialized is a log parse operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidatorRemovalInitialized(log types.Log) (*NativeTokenStakingManagerValidatorRemovalInitialized, error) {
	event := new(NativeTokenStakingManagerValidatorRemovalInitialized)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenStakingManagerValidatorWeightUpdateIterator is returned from FilterValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for ValidatorWeightUpdate events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidatorWeightUpdateIterator struct {
	Event *NativeTokenStakingManagerValidatorWeightUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenStakingManagerValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidatorWeightUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeTokenStakingManagerValidatorWeightUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeTokenStakingManagerValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidatorWeightUpdate represents a ValidatorWeightUpdate event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidatorWeightUpdate struct {
	ValidationID       [32]byte
	Nonce              uint64
	ValidatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorWeightUpdate is a free log retrieval operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte, nonce []uint64) (*NativeTokenStakingManagerValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidatorWeightUpdateIterator{contract: _NativeTokenStakingManager.contract, event: "ValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorWeightUpdate is a free log subscription operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidatorWeightUpdate, validationID [][32]byte, nonce []uint64) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidatorWeightUpdate)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorWeightUpdate is a log parse operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidatorWeightUpdate(log types.Log) (*NativeTokenStakingManagerValidatorWeightUpdate, error) {
	event := new(NativeTokenStakingManagerValidatorWeightUpdate)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
