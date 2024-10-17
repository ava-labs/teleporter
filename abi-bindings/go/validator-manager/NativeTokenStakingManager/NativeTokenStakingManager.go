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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADDRESS_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NATIVE_MINTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINativeMinter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POS_VALIDATOR_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPTIME_REWARDS_THRESHOLD_PERCENTAGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimDelegationFees\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimValidationRewards\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeEndValidation\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitializeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structValidator\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"startingWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"endedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWeight\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structPoSValidatorManagerSettings\",\"components\":[{\"name\":\"baseSettings\",\"type\":\"tuple\",\"internalType\":\"structValidatorManagerSettings\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"churnPeriodSeconds\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"minimumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rewardCalculator\",\"type\":\"address\",\"internalType\":\"contractIRewardCalculator\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeDelegatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initializeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeValidatorRegistration\",\"inputs\":[{\"name\":\"registrationInput\",\"type\":\"tuple\",\"internalType\":\"structValidatorRegistrationInput\",\"components\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"disableOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initializeValidatorSet\",\"inputs\":[{\"name\":\"subnetConversionData\",\"type\":\"tuple\",\"internalType\":\"structSubnetConversionData\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialValidators\",\"type\":\"tuple[]\",\"internalType\":\"structInitialValidator[]\",\"components\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredValidators\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resendEndValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendRegisterValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendUpdateDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitUptimeProof\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"valueToWeight\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"weightToValue\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"DelegationEnded\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewards\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorAdded\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"delegatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"delegatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRegistered\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRemovalInitialized\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitialValidatorCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UptimeUpdated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"uptime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodEnded\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumValidatorStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodRegistered\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationRewardsClaimed\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reward\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRemovalInitialized\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorWeightUpdate\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DelegatorIneligibleForRewards\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSKeyLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationFee\",\"inputs\":[{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationID\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDelegatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitializationStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMaximumChurnPercentage\",\"inputs\":[{\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinStakeDuration\",\"inputs\":[{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeID\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPChainOwnerThreshold\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addressesLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRegistrationExpiry\",\"inputs\":[{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeAmount\",\"inputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeMultiplier\",\"inputs\":[{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidSubnetConversionID\",\"inputs\":[{\"name\":\"encodedSubnetConversionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedSubnetConversionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidationID\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorManagerAddress\",\"inputs\":[{\"name\":\"validatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorManagerBlockchainID\",\"inputs\":[{\"name\":\"blockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWarpOriginSenderAddress\",\"inputs\":[{\"name\":\"senderAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpSourceChainID\",\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MaxChurnRateExceeded\",\"inputs\":[{\"name\":\"churnAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MaxWeightExceeded\",\"inputs\":[{\"name\":\"newValidatorWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MinStakeDurationNotPassed\",\"inputs\":[{\"name\":\"endTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NodeAlreadyRegistered\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PChainOwnerAddressesNotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnexpectedRegistrationStatus\",\"inputs\":[{\"name\":\"validRegistration\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"ValidatorIneligibleForRewards\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotPoS\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x36303830363034303532333438303135363130303066353735663830666435623530363034303531363135636233333830333830363135636233383333393831303136303430383139303532363130303265393136313031303735363562363030313831363030313831313131353631303034323537363130303432363130313263353635623033363130303466353736313030346636313030353535363562353036313031343035363562376666306335376531363834306466303430663135303838646332663831666533393163333932336265633733653233613936363265666339633232396336613030383035343638303130303030303030303030303030303030393030343630666631363135363130306135353736303430353136336639326565386139363065303162383135323630303430313630343035313830393130333930666435623830353436303031363030313630343031623033393038313136313436313031303435373830353436303031363030313630343031623033313931363630303136303031363034303162303339303831313738323535363034303531393038313532376663376635303562326633373161653231373565653439313366343439396531663236333361376235393336333231656564316364616562363131353138316432393036303230303136303430353138303931303339306131356235303536356235663630323038323834303331323135363130313137353735663830666435623831353136303032383131303631303132353537356638306664356239333932353035303530353635623633346534383762373136306530316235663532363032313630303435323630323435666664356236313562363638303631303134643566333935666633666536303830363034303532363030343336313036313031666335373566333536306530316338303633373332323134663831313631303131333537383036336237373162336263313136313030396435373830363363353939653234663131363130303664353738303633633539396532346631343631303563623537383036336339373464316236313436313035646535373830363364356632306666363134363130356632353738303633646639336438646531343631303631653537383036336664376163356537313436313036333435373566383066643562383036336237373162336263313436313035353335373830363362613361346239373134363130353664353738303633626335666266656331343631303538633537383036336265653061303366313436313035616335373566383066643562383036333933653234353938313136313030653335373830363339336532343539383134363130346365353738303633613361363565343831343631303465643537383036336166326635666562313436313035306335373830363361666239383039363134363130353166353738303633616662613837386131343631303533663537356638306664356238303633373332323134663831343631303436393537383036333736663738363231313436313034376335373830363338306464363732663134363130343962353738303633383238306132356131343631303462613537356638306664356238303633326532313934643831313631303139343537383036333436376566303666313136313031363435373830363334363765663036663134363130336236353738303633363033303564363231343631303364353537383036333630616437373834313436313033666535373830363336323036353835363134363130343164353738303633363634333561626631343631303434613537356638306664356238303633326532313934643831343631303330363537383036333332396333653132313436313033336435373830363333353435356465643134363130333666353738303633336131636666663631343631303339373537356638306664356238303633323064393162376131313631303163663537383036333230643931623761313436313032386135373830363332306535353536353134363130326139353738303633323565316337373631343631303263383537383036333238393364303737313436313032653735373566383066643562383036333031313861636334313436313032303035373830363330333232656439383134363130323231353738303633313531643330643131343631303234303537383036333165633434373234313436313032366235373562356638306664356233343830313536313032306235373566383066643562353036313032316636313032316133363630303436313462656335363562363130363533353635623030356233343830313536313032326335373566383066643562353036313032316636313032336233363630303436313463323735363562363130363635353635623334383031353631303234623537356638306664356235303631303235343630306138313536356236303430353136306666393039313136383135323630323030313562363034303531383039313033393066333562333438303135363130323736353735663830666435623530363130323166363130323835333636303034363134626563353635623631303866353536356233343830313536313032393535373566383066643562353036313032316636313032613433363630303436313463336535363562363130393031353635623334383031353631303262343537356638306664356235303631303231663631303263333336363030343631346338633536356236313065633935363562333438303135363130326433353735663830666435623530363130323166363130326532333636303034363134636133353635623631306661363536356233343830313536313032663235373566383066643562353036313032316636313033303133363630303436313463613335363562363131303161353635623334383031353631303331313537356638306664356235303631303332353631303332303336363030343631346332373536356236313133366135363562363034303531363030313630303136303430316230333930393131363831353236303230303136313032363235363562333438303135363130333438353735663830666435623530363130333537363030313630303136303939316230313831353635623630343035313630303136303031363061303162303339303931313638313532363032303031363130323632353635623334383031353631303337613537356638306664356235303631303338343631323731303831353635623630343035313631666666663930393131363831353236303230303136313032363235363562333438303135363130336132353735663830666435623530363130323166363130336231333636303034363134626563353635623631313338303536356233343830313536313033633135373566383066643562353036313032316636313033643033363630303436313463633435363562363131333863353635623334383031353631303365303537356638306664356235303631303365393630313438313536356236303430353136336666666666666666393039313136383135323630323030313631303236323536356233343830313536313034303935373566383066643562353036313032316636313034313833363630303436313463613335363562363131343761353635623334383031353631303432383537356638306664356235303631303433633631303433373336363030343631346366313536356236313137353235363562363034303531393038313532363032303031363130323632353635623334383031353631303435353537356638306664356235303631303332353631303436343336363030343631346332373536356236313137366235363562333438303135363130343734353735663830666435623530363130343363356638313536356233343830313536313034383735373566383066643562353036313032316636313034393633363630303436313462656335363562363131373766353635623334383031353631303461363537356638306664356235303631303231663631303462353336363030343631346361333536356236313137386335363562333438303135363130346335353735663830666435623530363130323534363033303831353635623334383031353631303464393537356638306664356235303631303231663631303465383336363030343631346332373536356236313139636435363562333438303135363130346638353735663830666435623530363130323166363130353037333636303034363134636334353635623631316137373536356236313034336336313035316133363630303436313464316435363562363131633663353635623334383031353631303532613537356638306664356235303631303433633566383035313630323036313561646138333339383135313931353238313536356233343830313536313035346135373566383066643562353036313032353436303530383135363562333438303135363130353565353735663830666435623530363130333537363030353630303136303939316230313831353635623334383031353631303537383537356638306664356235303631303231663631303538373336363030343631346332373536356236313163613035363562333438303135363130353937353735663830666435623530363130343363356638303531363032303631356166613833333938313531393135323831353635623334383031353631303562373537356638306664356235303631303231663631303563363336363030343631346332373536356236313166306235363562363130343363363130356439333636303034363134633237353635623631323034383536356233343830313536313035653935373566383066643562353036313032353436303134383135363562333438303135363130356664353735663830666435623530363130363131363130363063333636303034363134633237353635623631323037393536356236303430353136313032363239313930363134646633353635623334383031353631303632393537356638306664356235303631303332353632303261333030383135363562333438303135363130363366353735663830666435623530363130343363363130363465333636303034363134653733353635623631323163383536356236313036363038333833383336303031363132323233353635623530353035303536356235663831383135323566383035313630323036313562336138333339383135313931353236303230353236303430383038323230383135313630653038313031393039323532383035343566383035313630323036313561666138333339383135313931353239333932393139303832393036306666313636303035383131313135363130366232353736313036623236313464376535363562363030353831313131353631303663333537363130366333363134643765353635623831353236303230303136303031383230313830353436313036643739303631346564653536356238303630316630313630323038303931303430323630323030313630343035313930383130313630343035323830393239313930383138313532363032303031383238303534363130373033393036313465646535363562383031353631303734653537383036303166313036313037323535373631303130303830383335343034303238333532393136303230303139313631303734653536356238323031393139303566353236303230356632303930356238313534383135323930363030313031393036303230303138303833313136313037333135373832393030333630316631363832303139313562353035303530393138333532353035303630303238323031353436303031363030313630343031623033383038323136363032303834303135323630303136303430316238323034383131363630343038343031353236303031363038303162383230343831313636303630383430313532363030313630633031623930393130343831313636303830383330313532363030333932383330313534313636306130393039313031353239303931353038313531363030353831313131353631303762393537363130376239363134643765353635623134363130376635353735663833383135323630303738333031363032303532363034303930383139303230353439303531363331373063633933333630653231623831353236313037656339313630666631363930363030343031363134663130353635623630343035313830393130333930666435623630363038313031353136303430353136333432613265306235363065313162383135323630303438313031383539303532363030313630303136303430316230333930393131363630323438323031353235663630343438323031353236303035363030313630393931623031393036336565356234386562393037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f3930363338353435633136613930363036343031356636303430353138303833303338313836356166343135383031353631303836633537336435663830336533643566666435623530353035303530363034303531336435663832336536303166336439303831303136303166313931363832303136303430353236313038393339313930383130313930363135303231353635623630343035313832363366666666666666663136363065303162383135323630303430313631303861663931393036313530353235363562363032303630343035313830383330333831356638373561663131353830313536313038636235373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130386566393139303631353036343536356235303530353035303536356236313036363038333833383335663631323232333536356237666539323534366436393839353064646433383931306432653135656431643932336364306137623364646539653261366133663338303536353535396362303935343566383035313630323036313561666138333339383135313931353239303630666631363135363130393533353736303430353136333766616238316535363065303162383135323630303430313630343035313830393130333930666435623630303536303031363039393162303136303031363030313630613031623033313636333432313363663738363034303531383136336666666666666666313636306530316238313532363030343031363032303630343035313830383330333831383635616661313538303135363130393936353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631303962613931393036313530363435363562383336303230303133353134363130396533353736303430353136333732623061376537363065313162383135323630323038343031333536303034383230313532363032343031363130376563353635623330363130396634363036303835303136303430383630313631353038663536356236303031363030313630613031623033313631343631306133373537363130613132363036303834303136303430383530313631353038663536356236303430353136333266383831323064363065323162383135323630303136303031363061303162303339303931313636303034383230313532363032343031363130376563353635623566363130613435363036303835303138353631353061613536356239303530393035303566383035623832383136336666666666666666313631303135363130643264353735663631306136383630363038383031383836313530616135363562383336336666666666666666313638313831313036313061376535373631306137653631353065663536356239303530363032303032383130313930363130613930393139303631353130333536356236313061393939303631353136653536356238303531363034303531393139323530356639313630303838383031393136313061623139313631353165393536356239303831353236303230303136303430353138303931303339303230353431343631306165313537383035313630343035313633613431663737326636306530316238313532363130376563393139303630303430313631353035323536356235663630303238383566303133353834363034303531363032303031363130623130393239313930393138323532363065303162363030313630303136306530316230333139313636303230383230313532363032343031393035363562363034303830353136303166313938313834303330313831353239303832393035323631306232613931363135316539353635623630323036303430353138303833303338313835356166613135383031353631306234353537336435663830336533643566666435623530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130623638393139303631353036343536356239303530383038363630303830313833356630313531363034303531363130623830393139303631353165393536356239303831353236303430383035313630323039323831393030333833303138313230393339303933353536306530383330313831353236303032383335323834353138323834303135323834383130313830353136303031363030313630343031623033393038313136383538343031353235663630363038363031383139303532393135313831313636303830383630313532343231363630613038353031353236306330383430313831393035323834383135323630303738613031393039323532393032303831353138313534383239303630666631393136363030313833363030353831313131353631306330323537363130633032363134643765353635623032313739303535353036303230383230313531363030313832303139303631306331623930383236313532336535363562353036303430383238313031353136303032383330313830353436303630383630313531363038303837303135313630613038383031353136303031363030313630343031623033393538363136363030313630303136303830316230333139393039343136393339303933313736303031363034303162393238363136393239303932303239313930393131373630303136303031363038303162303331363630303136303830316239313835313639313930393130323630303136303031363063303162303331363137363030313630633031623931383431363931393039313032313739303535363063303930393330313531363030333930393230313830353436376666666666666666666666666666666631393136393238343136393239303932313739303931353538333031353136313063633039313136383536313533306435363562383235313630343035313931393535303631306364313931363135316539353635623630343038303531393138323930303338323230393038343031353136303031363030313630343031623033313638323532393038323930376639643437666566396461303737363631353436653634366436313833306266636264613930353036633265356565643338313935653832633465623163626466393036303230303136303430353138303931303339306133353035303830363130643236393036313533323035363562393035303631306134633536356235303630303438333031383139303535356637335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363331653664393738393631306435393837363132353466353635623630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363130643739393139303631353035323536356236303230363034303531383038333033383138363561663431353830313536313064393435373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130646238393139303631353036343536356239303530356637335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f36333836326266613633383836303430353138323633666666666666666631363630653031623831353236303034303136313064663239313930363135343664353635623566363034303531383038333033383138363561663431353830313536313065306335373364356638303365336435666664356235303530353035303630343035313364356638323365363031663364393038313031363031663139313638323031363034303532363130653333393139303831303139303631353032313536356239303530356636303032383236303430353136313065343639313930363135316539353635623630323036303430353138303833303338313835356166613135383031353631306536313537336435663830336533643566666435623530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130653834393139303631353036343536356239303530383238313134363130656230353736303430353136333138373266633864363065303162383135323630303438313031383239303532363032343831303138343930353236303434303136313037656335363562353035303530363030393930393230313830353436306666313931363630303131373930353535303530353035303536356237666630633537653136383430646630343066313530383864633266383166653339316333393233626563373365323361393636326566633963323239633661303038303534363030323931393036303031363034303162393030343630666631363830363130663132353735303830353436303031363030313630343031623033383038343136393131363130313535623135363130663330353736303430353136336639326565386139363065303162383135323630303430313630343035313830393130333930666435623830353436386666666666666666666666666666666666663139313636303031363030313630343031623033383331363137363030313630343031623137383135353631306635613833363132363635353635623830353436306666363034303162313931363831353536303430353136303031363030313630343031623033383331363831353237666337663530356232663337316165323137356565343931336634343939653166323633336137623539333633323165656431636461656236313135313831643239303630323030313630343035313830393130333930613135303530353035363562363130666166383236313236373635363562363130666366353736303430353136333330656661393862363065303162383135323630303438313031383339303532363032343031363130376563353635623566363130666439383336313230373935363562353139303530363030323831363030353831313131353631306666303537363130666630363134643765353635623134363131303130353738303630343035313633313730636339333336306532316238313532363030343031363130376563393139303631346631303536356236313038656638333833363132366231353635623631313032323631323939373536356235663830353136303230363135616461383333393831353139313532356636313130336138343631323037393536356239303530363030323831353136303035383131313135363131303531353736313130353136313464376535363562313436313130373235373830353136303430353136333137306363393333363065323162383135323631303765633931393036303034303136313466313035363562363131303762383436313236373635363562363131303962353736303430353136333330656661393862363065303162383135323630303438313031383539303532363032343031363130376563353635623566383438313532363030343833303136303230353236303430393032303534363030313630303136306130316230333136333331343631313065313537333335623630343035313633366532636364373536306531316238313532363030313630303136306130316230333930393131363630303438323031353236303234303136313037656335363562356638343831353236303034383330313630323035323630343039303230353436306130383230313531343239313631313131333931363030313630623031623930393130343630303136303031363034303162303331363930363135346638353635623630303136303031363034303162303331363831363030313630303136303430316230333136313031353631313135303537363034303531363366623663653633663630653031623831353236303031363030313630343031623033383231363630303438323031353236303234303136313037656335363562356636313131356238363836363132366231353635623566383738313532363030343836303136303230353236303430383132303630303130313534393139323530393036313131386339303630303136303430316239303034363030313630303136303430316230333136383336313535313835363562356638383831353236303034383730313630323035323630343038313230363030313031353439313932353036303031363038303162393039313034363030313630303136303430316230333136393038313930303336313131633235373530363061303834303135313562363131316364383238323836363132396365353635623631313165643537363034303531363335626666363833663630653131623831353236303034383130313839393035323630323430313631303765633536356236303033383630313534363034303836303135313566393136303031363030313630613031623033313639303633373738633036623539303631313231323930363131373532353635623834383538393838356638303630343035313838363366666666666666663136363065303162383135323630303430313631313233613937393639353934393339323931393036313535333835363562363032303630343035313830383330333831383635616661313538303135363131323535353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631313237393931393036313530363435363562393035303630363436303530363066663136383736306130303135313837363131323932393139303631353531383536356236313132396339313930363135353736353635623631313261363931393036313535623535363562356638613831353236303034383930313630323035323630343039303230363030313831303138303534373766666666666666666666666666666666666666666666666666666666666666663030303030303030303030303030303031393136363030313630343031623630303136303031363034303162303339343835313630323637666666666666666666666666666666663630383031623139313631373630303136303830316239333839313639333930393330323932393039323137393039313535353436313133316439303630303136303031363061303162303331363832363132613139353635623630343035313831393038613930376636396562336464393630323938373761373536316435623530373638313065306235336232323363396162343663626138656664303238656266303866653961393035663930613335303530353035303530353035303631313336363630303135663830353136303230363135623161383333393831353139313532353535363562353035303536356235663631313337613634653864346135313030303833363135356461353635623932393135303530353635623631303636303833383338333566363132613861353635623631313339343631323939373536356235663830353136303230363135616461383333393831353139313532356638303631313361643834363132636562353635623931353039313530363131336261383236313236373635363562363131336336353735303530353036313134363135363562356638323831353236303034383038353031363032303532363034303930393132303534363030313630303136306130316230333136393038323531363030353831313131353631313366353537363131336635363134643765353635623033363131343436353735663833383135323630303738353031363032303532363034303831323038303534393139303535363131343138383238323631326131393536356236303430353138313930383539303766363965623364643936303239383737613735363164356235303736383130653062353362323233633961623436636261386566643032386562663038666539613930356639306133353035623631313435633831363131343537383436303430303135313631313735323536356236313330396535363562353035303530353035623631313437373630303135663830353136303230363135623161383333393831353139313532353535363562353035363562356638323831353235663830353136303230363135616261383333393831353139313532363032303532363034303830383232303831353136306530383130313930393235323830353435663830353136303230363135616461383333393831353139313532393339323931393038323930363066663136363030333831313131353631313463373537363131346337363134643765353635623630303338313131313536313134643835373631313464383631346437653536356238313532383135343631303130303930303436303031363030313630613031623033313636303230383230313532363030313832303135343630343038303833303139313930393135323630303239303932303135343630303136303031363034303162303338303832313636303630383430313532363030313630343031623832303438313136363038303834303135323630303136303830316238323034383131363630613038343031353236303031363063303162393039313034313636306330393039313031353238313031353139303931353035663631313534653832363132303739353635623930353036303031383335313630303338313131313536313135363535373631313536353631346437653536356231343631313538363537383235313630343035313633336230643534306436306532316238313532363130376563393139303630303430313631353565643536356236303034383135313630303538313131313536313135396235373631313539623631346437653536356230333631313562313537363131356139383636313330623135363562353035303530353035303530353635623566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363339646532336434303631313564363839363132353466353635623630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363131356636393139303631353035323536356236303630363034303531383038333033383138363561663431353830313536313136313135373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363131363335393139303631353630373536356235303931353039313530383138343134363131363632353738343630343030313531363034303531363330383939333862333630653131623831353236303034303136313037656339313831353236303230303139303536356238303630303136303031363034303162303331363833363036303031353136303031363030313630343031623033313631303830363131363962353735303830363030313630303136303430316230333136383536306130303135313630303136303031363034303162303331363131356231353631313663343537363034303531363332653139626332643630653131623831353236303031363030313630343031623033383231363630303438323031353236303234303136313037656335363562356638383831353236303035383730313630323039303831353236303430393138323930323038303534363066663139313636303032393038313137383235353031383035343630303136303031363034303162303334323136363030313630343031623831303236666666666666666666666666666666666630303030303030303030303030303030313939303932313639313930393131373930393135353931353139313832353238353931386139313766303437303539623436353036396238623735313833366234316639663164383364616666353833643232333863633766626234363134333765633233613466363931303136303430353138303931303339306133353035303530353035303530353035303536356235663631313337613630303136303031363034303162303338333136363465386434613531303030363135363363353635623566363131373735383236313230373935363562363038303031353139323931353035303536356236313036363038333833383336303031363132613861353635623631313739343631323939373536356235663832383135323566383035313630323036313561626138333339383135313931353236303230353236303430383038323230383135313630653038313031393039323532383035343566383035313630323036313561646138333339383135313931353239333932393139303832393036306666313636303033383131313135363131376531353736313137653136313464376535363562363030333831313131353631313766323537363131376632363134643765353635623831353238313534363130313030393030343630303136303031363061303162303331363630323038323031353236303031383230313534363034303832303135323630303239303931303135343630303136303031363034303162303338303832313636303630383430313532363030313630343031623832303438313136363038303834303135323630303136303830316238323034383131363630613038343031353236303031363063303162393039313034313636306330393039313031353239303530363030333831353136303033383131313135363131383662353736313138366236313464376535363562313436313138386335373830353136303430353136333362306435343064363065323162383135323631303765633931393036303034303136313535656435363562363030343631313839623832363034303031353136313230373935363562353136303035383131313135363131386164353736313138616436313464376535363562313436313139616335373566363131386263383436313235346635363562393035303566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f3633396465323364343038343630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363131386662393139303631353035323536356236303630363034303531383038333033383138363561663431353830313536313139313635373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363131393361393139303631353630373536356235303931353039313530383138343630343030313531313436313139363635373630343035313633303839393338623336306531316238313532363030343831303138333930353236303234303136313037656335363562383036303031363030313630343031623033313638343630633030313531363030313630303136303430316230333136313131353631313961383537363034303531363332653139626332643630653131623831353236303031363030313630343031623033383231363630303438323031353236303234303136313037656335363562353035303530356236313139623538343631333062313536356235303530363131333636363030313566383035313630323036313562316138333339383135313931353235353536356235663830353136303230363135616461383333393831353139313532356636313139653538333631323037393536356235313930353036303034383136303035383131313135363131396663353736313139666336313464376535363562313436313161316335373830363034303531363331373063633933333630653231623831353236303034303136313037656339313930363134663130353635623566383338313532363030343833303136303230353236303430393032303534363030313630303136306130316230333136333331343631316134313537333336313130626335363562356638333831353236303037383330313630323039303831353236303430383038333230383035343930383439303535363030343836303139303932353239303931323035343631303865663930363030313630303136306130316230333136383236313261313935363562356638303531363032303631356166613833333938313531393135323566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363332653433636562353631316161613836363132353466353635623630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363131616361393139303631353035323536356236303430383035313830383330333831383635616634313538303135363131616534353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631316230383931393036313536353335363562393135303931353038303631316232653537363034303531363332643037313335333630653031623831353238313135313536303034383230313532363032343031363130376563353635623566383238313532363030363834303136303230353236303430393032303830353436313162343839303631346564653536356239303530356630333631316236633537363034303531363330383939333862333630653131623831353236303034383130313833393035323630323430313631303765633536356236303031356638333831353236303037383530313630323035323630343039303230353436306666313636303035383131313135363131623932353736313162393236313464376535363562313436313162633535373566383238313532363030373834303136303230353236303430393038313930323035343930353136333137306363393333363065323162383135323631303765633931363066663136393036303034303136313466313035363562356638323831353236303036383430313630323035323630343038313230363131626464393136313462343035363562356638323831353236303037383430313630323039303831353236303430393138323930323038303534363066663139313636303032393038313137383235353031383035343630303136303031363034303162303334323831383131363630303136306330316230323630303136303031363063303162303339303933313639323930393231373932383339303535383435313630303136303830316239303933303431363832353239313831303139313930393135323833393137666638666431633930666239636661326361323335386664663538303662303836616434333331356439326232323163393239656663376631303563653735363839313031363034303531383039313033393061323530353035303530353635623566363131633735363132393937353635623631316338313834383438343334363133323938353635623930353036313163393936303031356638303531363032303631356231613833333938313531393135323535353635623933393235303530353035363562356638313831353235663830353136303230363135616261383333393831353139313532363032303532363034303830383232303831353136306530383130313930393235323830353435663830353136303230363135616461383333393831353139313532393339323931393038323930363066663136363030333831313131353631316365643537363131636564363134643765353635623630303338313131313536313163666535373631316366653631346437653536356238313532383135343631303130303930303436303031363030313630613031623033313636303230383230313532363030313830383330313534363034303833303135323630303239303932303135343630303136303031363034303162303338303832313636303630383430313532363030313630343031623832303438313136363038303834303135323630303136303830316238323034383131363630613038343031353236303031363063303162393039313034313636306330393039313031353239303931353038313531363030333831313131353631316437373537363131643737363134643765353635623134313538303135363131643938353735303630303338313531363030333831313131353631316439353537363131643935363134643765353635623134313535623135363131646239353738303531363034303531363333623064353430643630653231623831353236313037656339313930363030343031363135356564353635623566363131646337383236303430303135313631323037393536356239303530383036303630303135313630303136303031363034303162303331363566303336313164663935373630343035313633333962383934663936306532316238313532363030343831303138353930353236303234303136313037656335363562363034303830383330313531363036303833303135313630383038343031353139323531363334326132653062353630653131623831353236303035363030313630393931623031393336336565356234386562393337335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f3933363338353435633136613933363131653637393339303630303430313932383335323630303136303031363034303162303339313832313636303230383430313532313636303430383230313532363036303031393035363562356636303430353138303833303338313836356166343135383031353631316538313537336435663830336533643566666435623530353035303530363034303531336435663832336536303166336439303831303136303166313931363832303136303430353236313165613839313930383130313930363135303231353635623630343035313832363366666666666666663136363065303162383135323630303430313631316563343931393036313530353235363562363032303630343035313830383330333831356638373561663131353830313536313165653035373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363131663034393139303631353036343536356235303530353035303530353635623566383138313532376665393235343664363938393530646464333839313064326531356564316439323363643061376233646465396532613661336633383035363535353963623036363032303532363034303930323038303534356638303531363032303631356166613833333938313531393135323931393036313166353239303631346564653536356239303530356630333631316637363537363034303531363330383939333862333630653131623831353236303034383130313833393035323630323430313631303765633536356236303031356638333831353236303037383330313630323035323630343039303230353436306666313636303035383131313135363131663963353736313166396336313464376535363562313436313166636635373566383238313532363030373832303136303230353236303430393038313930323035343930353136333137306363393333363065323162383135323631303765633931363066663136393036303034303136313466313035363562356638323831353236303036383230313630323035323630343039303831393032303930353136336565356234386562363065303162383135323630303536303031363039393162303139313633656535623438656239313631323030383931393036303034303136313536383135363562363032303630343035313830383330333831356638373561663131353830313536313230323435373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130363630393139303631353036343536356235663631323035313631323939373536356236313230356338323333333436313334616235363562393035303631323037343630303135663830353136303230363135623161383333393831353139313532353535363562393139303530353635623631323038313631346237373536356235663832383135323566383035313630323036313562336138333339383135313931353236303230353236303430393038313930323038313531363065303831303139303932353238303534356638303531363032303631356166613833333938313531393135323932393139303832393036306666313636303035383131313135363132306365353736313230636536313464376535363562363030353831313131353631323064663537363132306466363134643765353635623831353236303230303136303031383230313830353436313230663339303631346564653536356238303630316630313630323038303931303430323630323030313630343035313930383130313630343035323830393239313930383138313532363032303031383238303534363132313166393036313465646535363562383031353631323136613537383036303166313036313231343135373631303130303830383335343034303238333532393136303230303139313631323136613536356238323031393139303566353236303230356632303930356238313534383135323930363030313031393036303230303138303833313136313231346435373832393030333630316631363832303139313562353035303530393138333532353035303630303238323031353436303031363030313630343031623033383038323136363032303834303135323630303136303430316238323034383131363630343038343031353236303031363038303162383230343831313636303630383430313532363030313630633031623930393130343831313636303830383330313532363030333930393230313534393039313136363061303930393130313532393339323530353035303536356236303430353135663930356638303531363032303631356166613833333938313531393135323930376665393235343664363938393530646464333839313064326531356564316439323363643061376233646465396532613661336633383035363535353963623038393036313232306239303836393038363930363135373062353635623930383135323630323030313630343035313830393130333930323035343931353035303932393135303530353635623566383438313532356638303531363032303631356162613833333938313531393135323630323035323630343038303832323038313531363065303831303139303932353238303534356638303531363032303631356164613833333938313531393135323933393239313930383239303630666631363630303338313131313536313232373035373631323237303631346437653536356236303033383131313135363132323831353736313232383136313464376535363562383135323831353436313031303039303034363030313630303136306130316230333136363032303832303135323630303138323031353436303430383038333031393139303931353236303032393039323031353436303031363030313630343031623033383038323136363036303834303135323630303136303430316238323034383131363630383038343031353236303031363038303162383230343831313636306130383430313532363030313630633031623930393130343136363063303930393130313532383130313531393039313530356636313232663738323631323037393536356239303530363030323833353136303033383131313135363132333065353736313233306536313464376535363562313436313233326635373832353136303430353136333362306435343064363065323162383135323631303765633931393036303034303136313535656435363562363032303833303135313630303136303031363061303162303331363333313436313233643235373566383238313532363030343835303136303230353236303430393032303534333336303031363030313630613031623033393039313136303336313233636335373566383238313532363030343835303136303230353236303430393032303534363061303832303135313631323339333931363030313630623031623930303436303031363030313630343031623033313639303631353466383536356236303031363030313630343031623033313634323130313536313233633735373630343035313633666236636536336636306530316238313532363030313630303136303430316230333432313636303034383230313532363032343031363130376563353635623631323364323536356233333631313062633536356236303032383135313630303538313131313536313233653735373631323365373631346437653536356230333631323465383537383631353631323366653537363132336663383238373631323662313536356235303562356638383831353236303035383530313630323035323630343039303230383035343630666631393136363030333137393035353630363038333031353136303830383230313531363132343337393138343931363132343332393139303631353531383536356236313337383035363562353035663839383135323630303538363031363032303532363034303831323036303032303138303534363030313630303136303430316230333930393331363630303136306330316230323630303136303031363063303162303339303933313639323930393231373930393135353631323437383834363133393537353635623930353038353830313536313234383535373530383031353562313536313234613635373630343035313633313033366366393136306531316238313532363030343831303138613930353236303234303136313037656335363562356638393831353236303036383630313630323035323630343038303832323038333930353535313834393138623931376633363664333336633061623338306463373939663039356136663832613236333236353835633532393039636336393862303962613435343037303965643537393139306133353036313235343535363562363030343831353136303035383131313135363132346664353736313234666436313464376535363562303336313235323935373631323530623833363133393537353635623566383938313532363030363836303136303230353236303430393032303535363132353234383836313330623135363562363132353435353635623830353136303430353136333137306363393333363065323162383135323631303765633931393036303034303136313466313035363562353035303530353035303530353035303536356236303430383035313630363038303832303138333532356638303833353236303230383330313532393138313031393139303931353236303430353136333036663832353335363065343162383135323633666666666666666638333136363030343832303135323566393038313930363030353630303136303939316230313930363336663832353335303930363032343031356636303430353138303833303338313836356166613135383031353631323562333537336435663830336533643566666435623530353035303530363034303531336435663832336536303166336439303831303136303166313931363832303136303430353236313235646139313930383130313930363135373161353635623931353039313530383036313235666335373630343035313633366232663139653936306530316238313532363030343031363034303531383039313033393066643562383135313135363132363232353738313531363034303531363336626135383961353630653031623831353236303034383130313931393039313532363032343031363130376563353635623630323038323031353136303031363030313630613031623033313631353631323635653537363032303832303135313630343035313632346465373564363065333162383135323630303136303031363061303162303339303931313636303034383230313532363032343031363130376563353635623530393239313530353035363562363132363664363133623230353635623631313437373831363133623662353635623566393038313532376634333137373133663765636264646464346263393965393564393033616465646161383833623265376332353531363130626431336532633765343733643034363032303532363034303930323035343630303136303031363061303162303331363135313539303536356236303430353136333036663832353335363065343162383135323633666666666666666638323136363030343832303135323566393038313930383139303630303536303031363039393162303139303633366638323533353039303630323430313566363034303531383038333033383138363561666131353830313536313236666335373364356638303365336435666664356235303530353035303630343035313364356638323365363031663364393038313031363031663139313638323031363034303532363132373233393139303831303139303631353731613536356239313530393135303830363132373435353736303430353136333662326631396539363065303162383135323630303430313630343035313830393130333930666435623630303536303031363039393162303136303031363030313630613031623033313636333432313363663738363034303531383136336666666666666666313636306530316238313532363030343031363032303630343035313830383330333831383635616661313538303135363132373838353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631323761633931393036313530363435363562383235313134363132376432353738313531363034303531363336626135383961353630653031623831353236303034383130313931393039313532363032343031363130376563353635623630323038323031353136303031363030313630613031623033313631353631323830653537363032303832303135313630343035313632346465373564363065333162383135323630303136303031363061303162303339303931313636303034383230313532363032343031363130376563353635623566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f3633303838633234363338353630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363132383462393139303631353035323536356236303430383035313830383330333831383635616634313538303135363132383635353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631323838393931393036313537616135363562393135303931353038313837313436313238623035373630343035313633303839393338623336306531316238313532363030343831303138383930353236303234303136313037656335363562356638373831353237663433313737313366376563626464646434626339396539356439303361646564616138383362326537633235353136313062643133653263376534373364303436303230353236303430393032303630303130313534356638303531363032303631356164613833333938313531393135323930363030313630303136303430316230333930383131363930383331363131313536313239366435373566383838313532363030343832303136303230393038313532363034303931383239303230363030313031383035343637666666666666666666666666666666663139313636303031363030313630343031623033383631363930383131373930393135353931353139313832353238393931376665633434313438653866663237316632643062616365663131343231353461626163623061626233613239656233656235306532636139376538366430343335393130313630343035313830393130333930613236313239386335363562356638383831353236303034383230313630323035323630343039303230363030313031353436303031363030313630343031623033313639313530356235303936393535303530353035303530353035363562356638303531363032303631356231613833333938313531393135323830353436303031313930313631323963383537363034303531363333656535616562353630653031623831353236303034303136303430353138303931303339306664356236303032393035353536356235663630353036313239646238343834363135353138353635623631323965353931393036313535373635363562363030313630303136303430316230333136363132396639383536303634363135353736353635623630303136303031363034303162303331363130313536313261306635373530356636313163393935363562353036303031393339323530353035303536356236303430353136333237616435353564363065313162383135323630303136303031363061303162303338333136363030343832303135323630323438313031383239303532363030313630303136303939316230313930363334663561616162613930363034343031356636303430353138303833303338313566383738303362313538303135363132613635353735663830666435623530356166313135383031353631313561393537336435663830336533643566666435623630303135663830353136303230363135623161383333393831353139313532353535363562356638303531363032303631356164613833333938313531393135323566363132616132383636313362643935363562393035303631326161643836363132363736353635623631326162383537353035303631303865663536356235663836383135323630303438333031363032303532363034303930323035343630303136303031363061303162303331363333313436313261646435373333363131306263353635623566383638313532363030343833303136303230353236303430393032303534363061303832303135313631326230633931363030313630623031623930303436303031363030313630343031623033313639303631353466383536356236303031363030313630343031623033313638313630633030313531363030313630303136303430316230333136313031353631326235333537363063303831303135313630343035313633666236636536336636306530316238313532363030313630303136303430316230333930393131363630303438323031353236303234303136313037656335363562356638353135363132623662353736313262363438373836363132366231353635623930353036313262383935363562353035663836383135323630303438333031363032303532363034303930323036303031303135343630303136303031363034303162303331363562356638373831353236303034383430313630323035323630343038313230363030313031353436313262623639303630303136303430316239303034363030313630303136303430316230333136383336313535313835363562356638393831353236303034383630313630323035323630343038313230363030313031353439313932353036303031363038303162393039313034363030313630303136303430316230333136393038313930303336313262656335373530363061303833303135313562383538303135363132633035353735303631326330333832383238363630633030313531363132396365353635623135356231353631326332363537363034303531363335626666363833663630653131623831353236303034383130313861393035323630323430313631303765633536356236303033383530313534363034303835303135313566393136303031363030313630613031623033313639303633373738633036623539303631326334623930363131373532353635623834383538393630633030313531383835663830363034303531383836336666666666666666313636306530316238313532363030343031363132633737393739363935393439333932393139303631353533383536356236303230363034303531383038333033383138363561666131353830313536313263393235373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363132636236393139303631353036343536356239303530383038363630303730313566386338313532363032303031393038313532363032303031356632303566383238323534363132636461393139303631353330643536356239303931353535303530353035303530353035303530353035303530353035363562356636313263663436313462373735363562356638303531363032303631356166613833333938313531393135323566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363332653433636562353631326432373838363132353466353635623630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363132643437393139303631353035323536356236303430383035313830383330333831383635616634313538303135363132643631353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631326438353931393036313536353335363562393135303931353038303135363132646163353736303430353136333264303731333533363065303162383135323831313531353630303438323031353236303234303136313037656335363562356638323831353236303037383430313630323035323630343038303832323038313531363065303831303139303932353238303534383239303630666631363630303538313131313536313264646435373631326464643631346437653536356236303035383131313135363132646565353736313264656536313464376535363562383135323630323030313630303138323031383035343631326530323930363134656465353635623830363031663031363032303830393130343032363032303031363034303531393038313031363034303532383039323931393038313831353236303230303138323830353436313265326539303631346564653536356238303135363132653739353738303630316631303631326535303537363130313030383038333534303430323833353239313630323030313931363132653739353635623832303139313930356635323630323035663230393035623831353438313532393036303031303139303630323030313830383331313631326535633537383239303033363031663136383230313931356235303530353039313833353235303530363030323832303135343630303136303031363034303162303338303832313636303230383430313532363030313630343031623832303438313136363034303834303135323630303136303830316238323034383131363630363038343031353236303031363063303162393039313034383131363630383038333031353236303033393238333031353431363630613039303931303135323930393135303831353136303035383131313135363132656534353736313265653436313464376535363562313431353830313536313266303535373530363030313831353136303035383131313135363132663032353736313266303236313464376535363562313431353562313536313266323635373830353136303430353136333137306363393333363065323162383135323631303765633931393036303034303136313466313035363562363030333831353136303035383131313135363132663362353736313266336236313464376535363562303336313266343935373630303438313532363132663465353635623630303538313532356238333630303830313831363032303031353136303430353136313266363439313930363135316539353635623930383135323630343038303531363032303932383139303033383330313930323035663930383139303535383538313532363030373837303139303932353239303230383135313831353438333932393139303832393036306666313931363630303138333630303538313131313536313266613835373631326661383631346437653536356230323137393035353530363032303832303135313630303138323031393036313266633139303832363135323365353635623530363034303832303135313630303238323031383035343630363038353031353136303830383630313531363061303837303135313630303136303031363034303162303339353836313636303031363030313630383031623033313939303934313639333930393331373630303136303430316239323836313639323930393230323931393039313137363030313630303136303830316230333136363030313630383031623931383531363931393039313032363030313630303136306330316230333136313736303031363063303162393138343136393139303931303231373930353536306330393039323031353136303033393039313031383035343637666666666666666666666666666666663139313639313930393231363137393035353830353136303035383131313135363133303637353736313330363736313464376535363562363034303531383439303766316330386535393635366631613138646332646137363832366364633532383035633433653839376131376335306661656662386162336331353236636331363930356639306133393139363931393535303930393335303530353035303536356236313133363636303031363030313630613031623033383331363832363133656264353635623566383138313532356638303531363032303631356162613833333938313531393135323630323035323630343038303832323038313531363065303831303139303932353238303534356638303531363032303631356164613833333938313531393135323933393239313930383239303630666631363630303338313131313536313330666535373631333066653631346437653536356236303033383131313135363133313066353736313331306636313464376535363562383135323831353436313031303039303034363030313630303136306130316230333136363032303830383330313931393039313532363030313830383430313534363034303830383530313931393039313532363030323934383530313534363030313630303136303430316230333830383231363630363038373031353236303031363034303162383230343831313636303830383730313532363030313630383031623832303438313136363061303837303135323630303136306330316239303931303431363630633039303934303139333930393335323834383330313531356638393831353236303035383930313834353238343831323038303534363030313630303136306138316230333139313638313535393238333031383139303535393139303934303138313930353536303036383730313930393135323930383132303830353439303832393035353932393335303930393139303830383231353631333233643537356638343831353236303034383730313630323035323630343039303230353436313237313039303631333165663930363030313630613031623930303436316666666631363835363135363363353635623631333166393931393036313535646135363562393135303831383636303037303135663836383135323630323030313930383135323630323030313566323035663832383235343631333231643931393036313533306435363562393039313535353036313332326439303530383238343631353763643536356239303530363133323364383536303230303135313832363132613139353635623631333235323835363032303031353136313134353738373630363030313531363131373532353635623630343038303531383238313532363032303831303138343930353238353931383939313766386563656366353130303730633332306439613535333233666661626533353065323934616535303566633063353039646335373336646136663563633939333931303136303430353138303931303339306133353035303530353035303530353035363562376634333137373133663765636264646464346263393965393564393033616465646161383833623265376332353531363130626431336532633765343733643032353435663930356638303531363032303631356164613833333938313531393135323930363030313630343031623930303436316666666639303831313639303836313631303830363133326563353735303631323731303631666666663836313631313562313536313333313035373630343035313633356631326536633336306531316238313532363166666666383631363630303438323031353236303234303136313037656335363562363030323831303135343630303136303031363034303162303339303831313639303835313631303135363133333463353736303430353136323032613036643630653131623831353236303031363030313630343031623033383531363630303438323031353236303234303136313037656335363562383035343833313038303631333335653537353038303630303130313534383331313562313536313333376635373630343035313633323232643136343336306532316238313532363030343831303138343930353236303234303136313037656335363562383235663631333338613832363131333661353635623930353035663631333339373839383336313366353035363562393035303630343035313830363063303031363034303532383036313333616233333930353635623630303136303031363061303162303339303831313638323532363166666666383038633136363032303830383530313931393039313532363030313630303136303430316230333830386431363630343038303837303139313930393135323566363036303830383830313832393035323630383038303839303138333930353236306130393838393031383339303532386138333532363030343930396430313835353239303832393032303838353138313534393538613031353139333861303135313835313636303031363062303162303236376666666666666666666666666666666636306230316231393934393039373136363030313630613031623032363030313630303136306230316230333139393039363136393731363936393039363137393339303933313731363932393039323137383335353834303135313630303139303932303138303534393838353031353139343930393330313531383131363630303136303830316230323637666666666666666666666666666666663630383031623139393438323136363030313630343031623032363030313630303136303830316230333139393039393136393239303931313639313930393131373936393039363137393139303931313639343930393431373930393335353530393039313530353039343933353035303530353035363562356635663830353136303230363135616461383333393831353139313532383136313334633438343631313336613536356239303530356636313334643038373631323037393536356239303530363133346462383736313236373635363562363133346662353736303430353136333330656661393862363065303162383135323630303438313031383839303532363032343031363130376563353635623630303238313531363030353831313131353631333531303537363133353130363134643765353635623134363133353331353738303531363034303531363331373063633933333630653231623831353236313037656339313930363030343031363134663130353635623566383238323630383030313531363133353432393139303631353466383536356239303530383336303032303136303061393035343930363130313030306139303034363030313630303136303430316230333136383236303430303135313631333536623931393036313535373635363562363030313630303136303430316230333136383136303031363030313630343031623033313631313135363133356138353736303430353136333664353166653035363065313162383135323630303136303031363034303162303338323136363030343832303135323630323430313631303765633536356235663830363133356234386138343631333738303536356239313530393135303566386138333630343035313630323030313631333565323932393139303931383235323630633031623630303136303031363063303162303331393136363032303832303135323630323830313930353635623630343038303531363031663139383138343033303138313532383238323532383035313630323039303931303132303630653038333031393039313532393135303830363030313831353236303031363030313630613031623033386331363630323038303833303139313930393135323630343038303833303138663930353236303031363030313630343031623033383038623136363036303835303135323566363038303835303138313930353239303838313636306130383530313532363063303930393330313833393035323834383335323630303538623031393039313532393032303831353138313534383239303630666631393136363030313833363030333831313131353631333637353537363133363735363134643765353635623032313739303535353036303230383238313031353138323534363130313030363030313630613831623033313931363631303130303630303136303031363061303162303339323833313630323137383335353630343038303835303135313630303138353031353536303630383038363031353136303032393039353031383035343630383038303839303135313630613038613031353136306330393039613031353136303031363030313630343031623033393938613136363030313630303136303830316230333139393039343136393339303933313736303031363034303162393138613136393139303931303231373630303136303031363038303162303331363630303136303830316239393839313639393930393930323630303136303031363063303162303331363938393039383137363030313630633031623931383831363931393039313032313739303535383135313839383631363831353238613836313639343831303139343930393435323933386231363930383330313532393138313031383539303532393038633136393138643931383439313766623030323462323633626333613062373238613665646561353061363965666138343131383966386433326565386166396431633262316131613232333432363931303136303430353138303931303339306134396139393530353035303530353035303530353035303530353635623566383238313532356638303531363032303631356233613833333938313531393135323630323035323630343038313230363030323031353438313930356638303531363032303631356166613833333938313531393135323930363030313630383031623930303436303031363030313630343031623033313636313337633838353832363134353338353635623566363133376432383736313437313235363562356638383831353236303037383530313630323035323630343038303832323036303032303138303534363766666666666666666666666666666666363038303162313931363630303136303830316236303031363030313630343031623033386338313136393138323032393239303932313739303932353539313531363334326132653062353630653131623831353236303034383130313863393035323931383431363630323438333031353236303434383230313532393139323530393036303035363030313630393931623031393036336565356234386562393037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f393036333835343563313661393036303634303135663630343035313830383330333831383635616634313538303135363133383762353733643566383033653364356666643562353035303530353036303430353133643566383233653630316633643930383130313630316631393136383230313630343035323631333861323931393038313031393036313530323135363562363034303531383236336666666666666666313636306530316238313532363030343031363133386265393139303631353035323536356236303230363034303531383038333033383135663837356166313135383031353631333864613537336435663830336533643566666435623530353035303530363034303531336436303166313936303166383230313136383230313830363034303532353038313031393036313338666539313930363135303634353635623630343038303531363030313630303136303430316230333861383131363832353236303230383230313834393035323832353139333934353038353136393238623932376630376465356666333561363734613830303565363631663333333363393037636136333333343632383038373632643139646337623361626231613863316466393238323930303330313930613339303934353039323530353035303562393235303932393035303536356235663830356638303531363032303631356164613833333938313531393135323930353035663631333937373834363034303031353136313230373935363562393035303566363030333832353136303035383131313135363133393866353736313339386636313464376535363562313438303631333961643537353036303034383235313630303538313131313536313339616235373631333961623631346437653536356231343562313536313339626435373530363063303831303135313631333966613536356236303032383235313630303538313131313536313339643235373631333964323631346437653536356230333631333964653537353034323631333966613536356238313531363034303531363331373063633933333630653231623831353236313037656339313930363030343031363134663130353635623834363038303031353136303031363030313630343031623033313638313630303136303031363034303162303331363131363133613231353735303566393439333530353035303530353635623630343038303836303135313566393038313532363030343835303136303230353232303630303130313534363061303833303135313631336135303931363030313630303136303430316230333136393038333631323963653536356236313361356535373530356639343933353035303530353035363562363030333833303135343630363038363031353136303031363030313630613031623033393039313136393036333737386330366235393036313361383339303631313735323536356236306130383530313531363038303839303135313630343038303862303135313566393038313532363030343830386230313630323035323832383232303630303130313534393235313630303136303031363065303162303331393630653038393930316231363831353236313361643839363935393439333861393336303031363030313630343031623033393039313136393239303931383239313031363135353338353635623630323036303430353138303833303338313836356166613135383031353631336166333537336435663830336533643566666435623530353035303530363034303531336436303166313936303166383230313136383230313830363034303532353038313031393036313362313739313930363135303634353635623935393435303530353035303530353635623766663063353765313638343064663034306631353038386463326638316665333931633339323362656337336532336139363632656663396332323963366130303534363030313630343031623930303436306666313636313362363935373630343035313633316166636437396636306533316238313532363030343031363034303531383039313033393066643562353635623631336237333631336232303536356236313362376338313631343738373536356236313362383436313437613035363562363131343737363036303832303133353630383038333031333536313362613136306330383530313630613038363031363134636631353635623631336262313630653038363031363063303837303136313537653035363562363133626332363130313030383730313630653038383031363135376639353635623631336264343631303132303838303136313031303038393031363135303866353635623631343762303536356236313362653136313462373735363562356638323831353235663830353136303230363135623361383333393831353139313532363032303532363034303830383232303831353136306530383130313930393235323830353435663830353136303230363135616661383333393831353139313532393339323931393038323930363066663136363030353831313131353631336332653537363133633265363134643765353635623630303538313131313536313363336635373631336333663631346437653536356238313532363032303031363030313832303138303534363133633533393036313465646535363562383036303166303136303230383039313034303236303230303136303430353139303831303136303430353238303932393139303831383135323630323030313832383035343631336337663930363134656465353635623830313536313363636135373830363031663130363133636131353736313031303038303833353430343032383335323931363032303031393136313363636135363562383230313931393035663532363032303566323039303562383135343831353239303630303130313930363032303031383038333131363133636164353738323930303336303166313638323031393135623530353035303931383335323530353036303032383238313031353436303031363030313630343031623033383038323136363032303835303135323630303136303430316238323034383131363630343038353031353236303031363038303162383230343831313636303630383530313532363030313630633031623930393130343831313636303830383430313532363030333930393330313534393039323136363061303930393130313532393039313530383135313630303538313131313536313364333835373631336433383631346437653536356231343631336436623537356638343831353236303037383330313630323035323630343039303831393032303534393035313633313730636339333336306532316238313532363130376563393136306666313639303630303430313631346631303536356236303033383135323432363030313630303136303430316230333136363063303832303135323566383438313532363030373833303136303230353236303430393032303831353138313534383339323931393038323930363066663139313636303031383336303035383131313135363133646166353736313364616636313464376535363562303231373930353535303630323038323031353136303031383230313930363133646338393038323631353233653536356235303630343038323031353136303032383230313830353436303630383530313531363038303836303135313630613038373031353136303031363030313630343031623033393538363136363030313630303136303830316230333139393039343136393339303933313736303031363034303162393238363136393239303932303239313930393131373630303136303031363038303162303331363630303136303830316239313835313639313930393130323630303136303031363063303162303331363137363030313630633031623931383431363931393039313032313739303535363063303930393230313531363030333930393130313830353436376666666666666666666666666666666631393136393139303932313631373930353535663631336536363835383236313337383035363562363038303834303135313630343038303531363030313630303136303430316230333930393231363832353234323630323038333031353239313933353038333932353038373931376631336435383339346366323639643438626366393237393539613239613566666565376339393234646166666638393237656364663363343866666137633637393130313630343035313830393130333930613335303933393235303530353035363562383034373130313536313365653035373630343035313633636437383630353936306530316238313532333036303034383230313532363032343031363130376563353635623566383236303031363030313630613031623033313638323630343035313566363034303531383038333033383138353837356166313932353035303530336438303566383131343631336632393537363034303531393135303630316631393630336633643031313638323031363034303532336438323532336435663630323038343031336536313366326535363562363036303931353035623530353039303530383036313036363035373630343035313633306131326635323136306531316238313532363030343031363034303531383039313033393066643562376665393235343664363938393530646464333839313064326531356564316439323363643061376233646465396532613661336633383035363535353963623039353435663930363066663136363133663934353736303430353136333766616238316535363065303162383135323630303430313630343035313830393130333930666435623566383035313630323036313561666138333339383135313931353234323631336662333630363038363031363034303837303136313463663135363562363030313630303136303430316230333136313131353830363133666564353735303631336664313632303261333030343236313533306435363562363133666531363036303836303136303430383730313631346366313536356236303031363030313630343031623033313631303135356231353631343032373537363134303032363036303835303136303430383630313631346366313536356236303430353136333538373964613133363065313162383135323630303136303031363034303162303339303931313636303034383230313532363032343031363130376563353635623631343033633631343033373630363038363031383636313538313935363562363134386466353635623631343034633631343033373630383038363031383636313538313935363562363033303631343035623630323038363031383636313538326435363562393035303134363134303864353736313430366636303230383530313835363135383264353635623630343035313633323634373562326636306531316238313532363130376563393235303630303430313930383135323630323030313930353635623631343039373834383036313538326435363562393035303566303336313430633435373631343061393834383036313538326435363562363034303531363333653038613132353630653131623831353236303034303136313037656339323931393036313538366635363562356636303038383230313631343064333836383036313538326435363562363034303531363134306531393239313930363135373062353635623930383135323630323030313630343035313830393130333930323035343134363134313161353736313430666638343830363135383264353635623630343035313633613431663737326636306530316238313532363030343031363130376563393239313930363135383666353635623631343132343833356636313435333835363562363034303830353136306530383130313930393135323831353438313532356639303831393037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f393036336530343762323833393036303230383130313631343136313861383036313538326435363562383038303630316630313630323038303931303430323630323030313630343035313930383130313630343035323830393339323931393038313831353236303230303138333833383038323834333735663932303139313930393135323530353035303930383235323530363032303930383130313930363134316139393038623031386236313538326435363562383038303630316630313630323038303931303430323630323030313630343035313930383130313630343035323830393339323931393038313831353236303230303138333833383038323834333735663932303139313930393135323530353035303930383235323530363032303031363134316632363036303862303136303430386330313631346366313536356236303031363030313630343031623033313638313532363032303031363134323064363036303862303138623631353831393536356236313432313639303631353838323536356238313532363032303031363134323238363038303862303138623631353831393536356236313432333139303631353838323536356238313532363032303031383836303031363030313630343031623033313638313532353036303430353138323633666666666666666631363630653031623831353236303034303136313432356639313930363135396134353635623566363034303531383038333033383138363561663431353830313536313432373935373364356638303365336435666664356235303530353035303630343035313364356638323365363031663364393038313031363031663139313638323031363034303532363134326130393139303831303139303631356135623536356235663832383135323630303638363031363032303532363034303930323039313933353039313530363134326265383238323631353233653536356235303831363030383834303136313432636538383830363135383264353635623630343035313631343264633932393139303631353730623536356239303831353236303430353139303831393030333630323030313831323039313930393135353633656535623438656236306530316238313532356639303630303536303031363039393162303139303633656535623438656239303631343331383930383539303630303430313631353035323536356236303230363034303531383038333033383135663837356166313135383031353631343333343537336435663830336533643566666435623530353035303530363034303531336436303166313936303166383230313136383230313830363034303532353038313031393036313433353839313930363135303634353635623630343038303531363065303831303139303931353239303931353038303630303138313532363032303031363134333738383938303631353832643536356238303830363031663031363032303830393130343032363032303031363034303531393038313031363034303532383039333932393139303831383135323630323030313833383338303832383433373566393230313832393035323530393338353532353035303530363030313630303136303430316230333839313636303230383038343031383239303532363034303830383530313834393035323630363038353031393239303932353236303830383430313833393035323630613039303933303138323930353238363832353236303037383830313930393235323230383135313831353438323930363066663139313636303031383336303035383131313135363134343037353736313434303736313464376535363562303231373930353535303630323038323031353136303031383230313930363134343230393038323631353233653536356235303630343038323031353136303032383230313830353436303630383530313531363038303836303135313630613038373031353136303031363030313630343031623033393538363136363030313630303136303830316230333139393039343136393339303933313736303031363034303162393238363136393239303932303239313930393131373630303136303031363038303162303331363630303136303830316239313835313639313930393130323630303136303031363063303162303331363137363030313630633031623931383431363931393039313032313739303535363063303930393230313531363030333930393130313830353436376666666666666666666666666666666631393136393139303932313631373930353538303631343462653838383036313538326435363562363034303531363134346363393239313930363135373062353635623630343035313830393130333930323038343766623737323937653362656663363931626663383634613831653234316638336532656637323262366537626563616132656365633235306336643532623433303839386236303430303136303230383130313930363134353061393139303631346366313536356236303430383035313630303136303031363034303162303339333834313638313532393239303931313636303230383330313532303136303430353138303931303339306134353039303935393435303530353035303530353635623566383035313630323036313561666138333339383135313931353235663630303136303031363034303162303338303834313639303835313631313135363134353663353736313435363538333835363135353138353635623930353036313435373935363562363134353736383438343631353531383536356239303530356236303430383035313630383038313031383235323630303238343031353438303832353236303033383530313534363032303833303135323630303438353031353439323832303139323930393235323630303538343031353436303031363030313630343031623033313636303630383230313532343239313135383036313435646235373530363030313834303135343831353136313435643739313630303136303031363034303162303331363930363135333064353635623832313031353562313536313436303135373630303136303031363034303162303338333136363036303832303135323831383135323630343038313031353136303230383230313532363134363230353635623832383136303630303138313831353136313436313339313930363135346638353635623630303136303031363034303162303331363930353235303562363036303831303135313631343633303930363036343631353537363536356236303230383230313531363030313836303135343630303136303031363034303162303339323930393231363931363134363562393139303630303136303430316239303034363066663136363135363363353635623130313536313436386235373630363038313031353136303430353136336466616538383031363065303162383135323630303136303031363034303162303339303931313636303034383230313532363032343031363130376563353635623835363030313630303136303430316230333136383136303430303138313831353136313436613639313930363135333064353635623930353235303630343038313031383035313630303136303031363034303162303338373136393139303631343663363930383339303631353763643536356239303532353038303531363030323835303135353630323038313031353136303033383530313535363034303831303135313630303438353031353536303630303135313630303539303933303138303534363766666666666666666666666666666666313931363630303136303031363034303162303339303934313639333930393331373930393235353530353035303530353635623566383138313532356638303531363032303631356233613833333938313531393135323630323035323630343038313230363030323031383035343566383035313630323036313561666138333339383135313931353239313930363030383930363134373563393036303031363034303162393030343630303136303031363034303162303331363631356139653536356239313930363130313030306138313534383136303031363030313630343031623033303231393136393038333630303136303031363034303162303331363032313739303535393135303530393139303530353635623631343738663631336232303536356236313437393736313461343835363562363131343737383136313461353035363562363134376138363133623230353635623631336236393631346233383536356236313437623836313362323035363562356638303531363032303631356164613833333938313531393135323631666666663834313631353830363134376463353735303631323731303631666666663835313631313562313536313438303035373630343035313633356631326536633336306531316238313532363166666666383531363630303438323031353236303234303136313037656335363562383538373131313536313438323435373630343035313633323232643136343336306532316238313532363030343831303138383930353236303234303136313037656335363562363066663833313631353830363134383337353735303630306136306666383431363131356231353631343835613537363034303531363331373064623335393630653331623831353236306666383431363630303438323031353236303234303136313037656335363562393538363535363030313836303139343930393435353630303238353031383035343630303136303031363034303162303339343930393431363639666666666666666666666666666666666666666631393930393431363933393039333137363030313630343031623631666666663933393039333136393239303932303239313930393131373637666666666666666666666666666666663630353031623139313636306666393139303931313636303031363035303162303231373930353536303033393039313031383035343630303136303031363061303162303331393136363030313630303136306130316230333930393231363931393039313137393035353536356236313438656336303230383230313832363134636334353635623633666666666666666631363135383031353631343930633537353036313439303736303230383230313832363135306161353635623135313539303530356231353631343935333537363134393165363032303832303138323631346363343536356236313439326236303230383330313833363135306161353635623630343035313633633038613066316436306530316238313532363366666666666666663930393331363630303438343031353236303234383330313532353036303434303136313037656335363562363134393630363032303832303138323631353061613536356239303530363134393666363032303833303138333631346363343536356236336666666666666666313631313135363134393838353736313439316536303230383230313832363134636334353635623630303135623631343939383630323038333031383336313530616135363562393035303831313031353631313336363537363134396165363032303833303138333631353061613536356236313439623936303031383436313537636435363562383138313130363134396338353736313439633836313530656635363562393035303630323030323031363032303831303139303631343964643931393036313530386635363562363030313630303136306130316230333136363134396633363032303834303138343631353061613536356238333831383131303631346130333537363134613033363135306566353635623930353036303230303230313630323038313031393036313461313839313930363135303866353635623630303136303031363061303162303331363130313536313461343035373630343035313633306462633864356636306533316238313532363030343031363034303531383039313033393066643562363030313031363134393862353635623631336236393631336232303536356236313461353836313362323035363562383033353566383035313630323036313561666138333339383135313931353239303831353536303134363134613764363036303834303136303430383530313631353766393536356236306666313631313830363134613963353735303631346139373630363038333031363034303834303136313537663935363562363066663136313535623135363134616430353736313461623136303630383330313630343038343031363135376639353635623630343035313633346135396262666636306531316238313532363066663930393131363630303438323031353236303234303136313037656335363562363134616530363036303833303136303430383430313631353766393536356236303031383230313830353436306666393239303932313636303031363034303162303236306666363034303162313939303932313639313930393131373930353536313462313136303430383330313630323038343031363134636631353635623630303139313930393130313830353436376666666666666666666666666666666631393136363030313630303136303430316230333930393231363931393039313137393035353530353635623631326137373631336232303536356235303830353436313462346339303631346564653536356235663832353538303630316631303631346235623537353035303536356236303166303136303230393030343930356635323630323035663230393038313031393036313134373739313930363134626234353635623630343038303531363065303831303139303931353238303566383135323630363036303230383230313831393035323566363034303833303138313930353239303832303138313930353236303830383230313831393035323630613038323031383139303532363063303930393130313532393035363562356238303832313131353631346263383537356638313535363030313031363134626235353635623530393035363562383031353135383131343631313437373537356638306664356238303335363366666666666666663831313638313134363132303734353735663830666435623566383035663630363038343836303331323135363134626665353735663830666435623833333539323530363032303834303133353631346331303831363134626363353635623931353036313463316536303430383530313631346264393536356239303530393235303932353039323536356235663630323038323834303331323135363134633337353735663830666435623530333539313930353035363562356638303630343038333835303331323135363134633466353735663830666435623832333536303031363030313630343031623033383131313135363134633634353735663830666435623833303136303830383138363033313231353631346337353537356638306664356239313530363134633833363032303834303136313462643935363562393035303932353039323930353035363562356636313031323038323834303331323135363134633964353735663830666435623530393139303530353635623566383036303430383338353033313231353631346362343537356638306664356238323335393135303631346338333630323038343031363134626439353635623566363032303832383430333132313536313463643435373566383066643562363131633939383236313462643935363562363030313630303136303430316230333831313638313134363131343737353735663830666435623566363032303832383430333132313536313464303135373566383066643562383133353631316339393831363134636464353635623830333536316666666638313136383131343631323037343537356638306664356235663830356636303630383438363033313231353631346432663537356638306664356238333335363030313630303136303430316230333831313131353631346434343537356638306664356238343031363061303831383730333132313536313464353535373566383066643562393235303631346436333630323038353031363134643063353635623931353036303430383430313335363134643733383136313463646435363562383039313530353039323530393235303932353635623633346534383762373136306530316235663532363032313630303435323630323435666664356236303036383131303631346461323537363134646132363134643765353635623930353235363562356635623833383131303135363134646330353738313831303135313833383230313532363032303031363134646138353635623530353035663931303135323536356235663831353138303834353236313464646638313630323038363031363032303836303136313464613635363562363031663031363031663139313639323930393230313630323030313932393135303530353635623630323038313532363134653035363032303832303138333531363134643932353635623566363032303833303135313630653036303430383430313532363134653230363130313030383430313832363134646338353635623930353036303430383430313531363030313630303136303430316230333830383231363630363038363031353238303630363038373031353131363630383038363031353238303630383038373031353131363630613038363031353238303630613038373031353131363630633038363031353238303630633038373031353131363630653038363031353235303530383039313530353039323931353035303536356235663830363032303833383530333132313536313465383435373566383066643562383233353630303136303031363034303162303338303832313131353631346539613537356638306664356238313835303139313530383536303166383330313132363134656164353735663830666435623831333538313831313131353631346562623537356638306664356238363630323038323835303130313131313536313465636335373566383066643562363032303932393039323031393639313935353039303933353035303530353035363562363030313831383131633930383231363830363134656632353736303766383231363931353035623630323038323130383130333631346339643537363334653438376237313630653031623566353236303232363030343532363032343566666435623630323038313031363131333761383238343631346439323536356236333465343837623731363065303162356635323630343136303034353236303234356666643562363034303531363036303831303136303031363030313630343031623033383131313832383231303137313536313466353435373631346635343631346631653536356236303430353239303536356236303430383035313930383130313630303136303031363034303162303338313131383238323130313731353631346635343537363134663534363134663165353635623630343035313630316638323031363031663139313638313031363030313630303136303430316230333831313138323832313031373135363134666134353736313466613436313466316535363562363034303532393139303530353635623566363030313630303136303430316230333832313131353631346663343537363134666334363134663165353635623530363031663031363031663139313636303230303139303536356235663832363031663833303131323631346665313537356638306664356238313531363134666634363134666566383236313466616335363562363134663763353635623831383135323834363032303833383630313031313131353631353030383537356638306664356236313530313938323630323038333031363032303837303136313464613635363562393439333530353035303530353635623566363032303832383430333132313536313530333135373566383066643562383135313630303136303031363034303162303338313131313536313530343635373566383066643562363135303139383438323835303136313466643235363562363032303831353235663631316339393630323038333031383436313464633835363562356636303230383238343033313231353631353037343537356638306664356235303531393139303530353635623630303136303031363061303162303338313136383131343631313437373537356638306664356235663630323038323834303331323135363135303966353735663830666435623831333536313163393938313631353037623536356235663830383333353630316531393834333630333031383131323631353062663537356638306664356238333031383033353931353036303031363030313630343031623033383231313135363135306438353735663830666435623630323030313931353036303035383139303162333630333832313331353631333935303537356638306664356236333465343837623731363065303162356635323630333236303034353236303234356666643562356638323335363035653139383333363033303138313132363135313137353735663830666435623931393039313031393239313530353035363562356638323630316638333031313236313531333035373566383066643562383133353631353133653631346665663832363134666163353635623831383135323834363032303833383630313031313131353631353135323537356638306664356238313630323038353031363032303833303133373566393138313031363032303031393139303931353239333932353035303530353635623566363036303832333630333132313536313531376535373566383066643562363135313836363134663332353635623832333536303031363030313630343031623033383038323131313536313531396335373566383066643562363135316138333638333837303136313531323135363562383335323630323038353031333539313530383038323131313536313531626435373566383066643562353036313531636133363832383630313631353132313536356236303230383330313532353036303430383330313335363135316465383136313463646435363562363034303832303135323932393135303530353635623566383235313631353131373831383436303230383730313631346461363536356236303166383231313135363130363630353738303566353236303230356632303630316638343031363030353163383130313630323038353130313536313532316635373530383035623630316638343031363030353163383230313931353035623831383131303135363131663034353735663831353536303031303136313532326235363562383135313630303136303031363034303162303338313131313536313532353735373631353235373631346631653536356236313532366238313631353236353834353436313465646535363562383436313531666135363562363032303830363031663833313136303031383131343631353239653537356638343135363135323837353735303835383330313531356235663139363030333836393031623163313931363630303138353930316231373835353536313135613935363562356638353831353236303230383132303630316631393836313639313562383238313130313536313532636335373838383630313531383235353934383430313934363030313930393130313930383430313631353261643536356235303835383231303135363135326539353738373835303135313566313936303033383839303162363066383136316331393136383135353562353035303530353035303630303139303831316230313930353535303536356236333465343837623731363065303162356635323630313136303034353236303234356666643562383038323031383038323131313536313133376135373631313337613631353266393536356235663633666666666666666638303833313638313831303336313533333835373631353333383631353266393536356236303031303139333932353035303530353635623566383038333335363031653139383433363033303138313132363135333537353735663830666435623833303136303230383130313932353033353930353036303031363030313630343031623033383131313135363135333735353735663830666435623830333630333832313331353631333935303537356638306664356238313833353238313831363032303835303133373530356638323832303136303230393038313031393139303931353236303166393039313031363031663139313639303931303130313930353635623566383338333835353236303230383038363031393535303830383536303035316238333031303138343566356238373831313031353631353436303537383438333033363031663139303138393532383133353336383839303033363035653139303138313132363135336537353735663830666435623837303136303630363135336635383238303631353334323536356238323837353236313534303538333838303138323834363135333833353635623932353035303530363135343135383638333031383336313533343235363562383638333033383838383031353236313534323738333832383436313533383335363562393235303530353036303430383038333031333539323530363135343363383336313463646435363562363030313630303136303430316230333932393039323136393439303931303139333930393335323937383330313937393038333031393036303031303136313533633235363562353039303937393635303530353035303530353035303536356236303230383135323831333536303230383230313532363032303832303133353630343038323031353235663630343038333031333536313534393138313631353037623536356236303031363030313630613031623033313636303630383338313031393139303931353238333031333533363834393030333630316531393031383131323631353462383537356638306664356238333031363032303831303139303335363030313630303136303430316230333831313131353631353464333537356638306664356238303630303531623336303338323133313536313534653435373566383066643562363038303830383530313532363133623137363061303835303138323834363135336162353635623630303136303031363034303162303338313831313638333832313630313930383038323131313536313236356535373631323635653631353266393536356236303031363030313630343031623033383238313136383238323136303339303830383231313135363132363565353736313236356536313532663935363562393638373532363030313630303136303430316230333935383631363630323038383031353239333835313636303430383730313532393138343136363036303836303135323930393231363630383038343031353236306130383330313931393039313532363063303832303135323630653030313930353635623630303136303031363034303162303338313831313638333832313630323830383231363931393038323831313436313535393935373631353539393631353266393536356235303530393239313530353035363562363334653438376237313630653031623566353236303132363030343532363032343566666435623566363030313630303136303430316230333830383431363830363135356365353736313535636536313535613135363562393231363931393039313034393239313530353035363562356638323631353565383537363135356538363135356131353635623530303439303536356236303230383130313630303438333130363135363031353736313536303136313464376535363562393139303532393035363562356638303566363036303834383630333132313536313536313935373566383066643562383335313932353036303230383430313531363135363262383136313463646435363562363034303835303135313930393235303631346437333831363134636464353635623830383230323831313538323832303438343134313736313133376135373631313337613631353266393536356235663830363034303833383530333132313536313536363435373566383066643562383235313931353036303230383330313531363135363736383136313462636335363562383039313530353039323530393239303530353635623566363032303830383335323566383435343631353639333831363134656465353635623830363032303837303135323630343036303031383038343136356638313134363135366234353736303031383131343631353664303537363135366664353635623630666631393835313636303430386130313532363034303834313531353630303531623861303130313935353036313536666435363562383935663532363032303566323035663562383538313130313536313536663435373831353438623832303138363031353239303833303139303838303136313536643935363562386130313630343030313936353035303562353039333938393735303530353035303530353035303530353635623831383338323337356639313031393038313532393139303530353635623566383036303430383338353033313231353631353732623537356638306664356238323531363030313630303136303430316230333830383231313135363135373431353735663830666435623930383430313930363036303832383730333132313536313537353435373566383066643562363135373563363134663332353635623832353138313532363032303833303135313631353736653831363135303762353635623630323038323031353236303430383330313531383238313131313536313537383435373566383066643562363135373930383838323836303136313466643235363562363034303833303135323530383039343530353035303530363032303833303135313631353637363831363134626363353635623566383036303430383338353033313231353631353762623537356638306664356238323531393135303630323038333031353136313536373638313631346364643536356238313831303338313831313131353631313337613537363131333761363135326639353635623566363032303832383430333132313536313537663035373566383066643562363131633939383236313464306335363562356636303230383238343033313231353631353830393537356638306664356238313335363066663831313638313134363131633939353735663830666435623566383233353630336531393833333630333031383131323631353131373537356638306664356235663830383333353630316531393834333630333031383131323631353834323537356638306664356238333031383033353931353036303031363030313630343031623033383231313135363135383562353735663830666435623630323030313931353033363831393030333832313331353631333935303537356638306664356236303230383135323566363135303139363032303833303138343836363135333833353635623566363034303832333630333132313536313538393235373566383066643562363135383961363134663561353635623631353861333833363134626439353635623831353236303230383038343031333536303031363030313630343031623033383038323131313536313538626635373566383066643562393038353031393033363630316638333031313236313538643135373566383066643562383133353831383131313135363135386533353736313538653336313466316535363562383036303035316239313530363135386634383438333031363134663763353635623831383135323931383330313834303139313834383130313930333638343131313536313539306435373566383066643562393338353031393335623833383531303135363135393337353738343335393235303631353932373833363135303762353635623832383235323933383530313933393038353031393036313539313235363562393438363031393439303934353235303932393539343530353035303530353035363562356636303430383330313633666666666666666638333531313638343532363032303830383430313531363034303630323038373031353238323831353138303835353236303630383830313931353036303230383330313934353035663932353035623830383331303135363132393863353738343531363030313630303136306130316230333136383235323933383330313933363030313932393039323031393139303833303139303631353937623536356236303230383135323831353136303230383230313532356636303230383330313531363065303630343038343031353236313539636136313031303038343031383236313464633835363562393035303630343038343031353136303166313938303835383430333031363036303836303135323631353965383833383336313464633835363562393235303630303136303031363034303162303336303630383730313531313636303830383630313532363038303836303135313931353038303835383430333031363061303836303135323631356131383833383336313539343935363562393235303630613038363031353139313530383038353834303330313630633038363031353235303631356133363832383236313539343935363562393135303530363063303834303135313631356135333630653038353031383236303031363030313630343031623033313639303532353635623530393339323530353035303536356235663830363034303833383530333132313536313561366335373566383066643562383235313931353036303230383330313531363030313630303136303430316230333831313131353631356138383537356638306664356236313561393438353832383630313631346664323536356239313530353039323530393239303530353635623566363030313630303136303430316230333830383331363831383130333631353333383537363135333338363135326639353666653433313737313366376563626464646434626339396539356439303361646564616138383362326537633235353136313062643133653263376534373364303534333137373133663765636264646464346263393965393564393033616465646161383833623265376332353531363130626431336532633765343733643030653932353436643639383935306464643338393130643265313565643164393233636430613762336464653965326136613366333830353635353539636230303962373739623137343232643064663932323233303138623332623464316661343665303731373233643638313765323438366430303362656363353566303065393235343664363938393530646464333839313064326531356564316439323363643061376233646465396532613661336633383035363535353963623037613136343733366636633633343330303038313930303061",
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

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCaller) UPTIMEREWARDSTHRESHOLDPERCENTAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenStakingManager.contract.Call(opts, &out, "UPTIME_REWARDS_THRESHOLD_PERCENTAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) UPTIMEREWARDSTHRESHOLDPERCENTAGE() (uint8, error) {
	return _NativeTokenStakingManager.Contract.UPTIMEREWARDSTHRESHOLDPERCENTAGE(&_NativeTokenStakingManager.CallOpts)
}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) UPTIMEREWARDSTHRESHOLDPERCENTAGE() (uint8, error) {
	return _NativeTokenStakingManager.Contract.UPTIMEREWARDSTHRESHOLDPERCENTAGE(&_NativeTokenStakingManager.CallOpts)
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
// Solidity: function valueToWeight(uint256 value) pure returns(uint64)
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
// Solidity: function valueToWeight(uint256 value) pure returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _NativeTokenStakingManager.Contract.ValueToWeight(&_NativeTokenStakingManager.CallOpts, value)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) pure returns(uint64)
func (_NativeTokenStakingManager *NativeTokenStakingManagerCallerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _NativeTokenStakingManager.Contract.ValueToWeight(&_NativeTokenStakingManager.CallOpts, value)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) pure returns(uint256)
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
// Solidity: function weightToValue(uint64 weight) pure returns(uint256)
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _NativeTokenStakingManager.Contract.WeightToValue(&_NativeTokenStakingManager.CallOpts, weight)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) pure returns(uint256)
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

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) ClaimValidationRewards(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "claimValidationRewards", validationID, messageIndex)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) ClaimValidationRewards(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ClaimValidationRewards(&_NativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) ClaimValidationRewards(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.ClaimValidationRewards(&_NativeTokenStakingManager.TransactOpts, validationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteDelegatorRegistration(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) CompleteEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "completeEndDelegation", delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactorSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.CompleteEndDelegation(&_NativeTokenStakingManager.TransactOpts, delegationID, messageIndex)
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

// Initialize is a paid mutator transaction binding the contract method 0x20e55565.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,address) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerTransactor) Initialize(opts *bind.TransactOpts, settings PoSValidatorManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.contract.Transact(opts, "initialize", settings)
}

// Initialize is a paid mutator transaction binding the contract method 0x20e55565.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,address) settings) returns()
func (_NativeTokenStakingManager *NativeTokenStakingManagerSession) Initialize(settings PoSValidatorManagerSettings) (*types.Transaction, error) {
	return _NativeTokenStakingManager.Contract.Initialize(&_NativeTokenStakingManager.TransactOpts, settings)
}

// Initialize is a paid mutator transaction binding the contract method 0x20e55565.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,address) settings) returns()
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

// NativeTokenStakingManagerValidationRewardsClaimedIterator is returned from FilterValidationRewardsClaimed and is used to iterate over the raw logs and unpacked data for ValidationRewardsClaimed events raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationRewardsClaimedIterator struct {
	Event *NativeTokenStakingManagerValidationRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *NativeTokenStakingManagerValidationRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenStakingManagerValidationRewardsClaimed)
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
		it.Event = new(NativeTokenStakingManagerValidationRewardsClaimed)
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
func (it *NativeTokenStakingManagerValidationRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenStakingManagerValidationRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenStakingManagerValidationRewardsClaimed represents a ValidationRewardsClaimed event raised by the NativeTokenStakingManager contract.
type NativeTokenStakingManagerValidationRewardsClaimed struct {
	ValidationID [32]byte
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationRewardsClaimed is a free log retrieval operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) FilterValidationRewardsClaimed(opts *bind.FilterOpts, validationID [][32]byte, reward []*big.Int) (*NativeTokenStakingManagerValidationRewardsClaimedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.FilterLogs(opts, "ValidationRewardsClaimed", validationIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenStakingManagerValidationRewardsClaimedIterator{contract: _NativeTokenStakingManager.contract, event: "ValidationRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchValidationRewardsClaimed is a free log subscription operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) WatchValidationRewardsClaimed(opts *bind.WatchOpts, sink chan<- *NativeTokenStakingManagerValidationRewardsClaimed, validationID [][32]byte, reward []*big.Int) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _NativeTokenStakingManager.contract.WatchLogs(opts, "ValidationRewardsClaimed", validationIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenStakingManagerValidationRewardsClaimed)
				if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationRewardsClaimed", log); err != nil {
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

// ParseValidationRewardsClaimed is a log parse operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_NativeTokenStakingManager *NativeTokenStakingManagerFilterer) ParseValidationRewardsClaimed(log types.Log) (*NativeTokenStakingManagerValidationRewardsClaimed, error) {
	event := new(NativeTokenStakingManagerValidationRewardsClaimed)
	if err := _NativeTokenStakingManager.contract.UnpackLog(event, "ValidationRewardsClaimed", log); err != nil {
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
