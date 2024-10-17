// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20tokenstakingmanager

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

// ERC20TokenStakingManagerMetaData contains all meta data concerning the ERC20TokenStakingManager contract.
var ERC20TokenStakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADDRESS_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BIPS_CONVERSION_FACTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERC20_STAKING_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POS_VALIDATOR_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimDelegationFees\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeEndDelegation\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeEndValidation\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitializeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structValidator\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"startingWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"endedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWeight\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structPoSValidatorManagerSettings\",\"components\":[{\"name\":\"baseSettings\",\"type\":\"tuple\",\"internalType\":\"structValidatorManagerSettings\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"churnPeriodSeconds\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"minimumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"weightToValueFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardCalculator\",\"type\":\"address\",\"internalType\":\"contractIRewardCalculator\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20Mintable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeDelegatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delegationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeValidatorRegistration\",\"inputs\":[{\"name\":\"registrationInput\",\"type\":\"tuple\",\"internalType\":\"structValidatorRegistrationInput\",\"components\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"disableOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeValidatorSet\",\"inputs\":[{\"name\":\"subnetConversionData\",\"type\":\"tuple\",\"internalType\":\"structSubnetConversionData\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialValidators\",\"type\":\"tuple[]\",\"internalType\":\"structInitialValidator[]\",\"components\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredValidators\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resendEndValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendRegisterValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendUpdateDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitUptimeProof\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"valueToWeight\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"weightToValue\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DelegationEnded\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewards\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorAdded\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"delegatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"delegatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRegistered\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRemovalInitialized\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitialValidatorCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UptimeUpdated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"uptime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodEnded\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumValidatorStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodRegistered\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRemovalInitialized\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorWeightUpdate\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DelegatorIneligibleForRewards\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSKeyLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationFee\",\"inputs\":[{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationID\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDelegatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitializationStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMaximumChurnPercentage\",\"inputs\":[{\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinStakeDuration\",\"inputs\":[{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeID\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPChainOwnerThreshold\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addressesLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRegistrationExpiry\",\"inputs\":[{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeAmount\",\"inputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeMultiplier\",\"inputs\":[{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidSubnetConversionID\",\"inputs\":[{\"name\":\"encodedSubnetConversionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedSubnetConversionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenAddress\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTotalWeight\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidValidationID\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorManagerAddress\",\"inputs\":[{\"name\":\"validatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorManagerBlockchainID\",\"inputs\":[{\"name\":\"blockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWarpOriginSenderAddress\",\"inputs\":[{\"name\":\"senderAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpSourceChainID\",\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MaxChurnRateExceeded\",\"inputs\":[{\"name\":\"churnAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MaxWeightExceeded\",\"inputs\":[{\"name\":\"newValidatorWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MinStakeDurationNotPassed\",\"inputs\":[{\"name\":\"endTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NodeAlreadyRegistered\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PChainOwnerAddressesNotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnexpectedRegistrationStatus\",\"inputs\":[{\"name\":\"validRegistration\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"ValidatorIneligibleForRewards\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotPoS\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroWeightToValueFactor\",\"inputs\":[]}]",
	Bin: "0x3630383036303430353233343830313536313030306635373566383066643562353036303430353136313561356533383033383036313561356538333339383130313630343038313930353236313030326539313631303130373536356236303031383136303031383131313135363130303432353736313030343236313031326335363562303336313030346635373631303034663631303035353536356235303631303134303536356237666630633537653136383430646630343066313530383864633266383166653339316333393233626563373365323361393636326566633963323239633661303038303534363830313030303030303030303030303030303039303034363066663136313536313030613535373630343035313633663932656538613936306530316238313532363030343031363034303531383039313033393066643562383035343630303136303031363034303162303339303831313631343631303130343537383035343630303136303031363034303162303331393136363030313630303136303430316230333930383131373832353536303430353139303831353237666337663530356232663337316165323137356565343931336634343939653166323633336137623539333633323165656431636461656236313135313831643239303630323030313630343035313830393130333930613135623530353635623566363032303832383430333132313536313031313735373566383066643562383135313630303238313130363130313235353735663830666435623933393235303530353035363562363334653438376237313630653031623566353236303231363030343532363032343566666435623631353931313830363130313464356633393566663366653630383036303430353233343830313536313030306635373566383066643562353036303034333631303631303163313537356633353630653031633830363338323830613235613131363130306636353738303633626133613462393731313631303039613537383036336261336134623937313436313033643935373830363362633566626665633134363130336563353738303633626565306130336631343631303431333537383036336339373464316236313436313034323635373830363364356632306666363134363130343265353738303633646639336438646531343631303434653537383036336534613633633430313436313034353835373830363366373463363037623134363130343766353738303633666437616335653731343631303439323537356638306664356238303633383238306132356131343631303334333537383036333933653234353938313436313033346235373830363339386633653262343134363130333565353738303633396531626334656631343631303337313537383036336133613635653438313436313033383435373830363361393737386137613134363130323730353738303633616662393830393631343631303339373537383036336237373162336263313436313033626535373566383066643562383036333361316366666636313136313031363835373830363333613163666666363134363130323863353738303633343637656630366631343631303239663537383036333462656530303430313436313032623235373830363335323937666165363134363130326433353738303633363033303564363231343631303265363537383036333632303635383536313436313033303335373830363336363433356162663134363130333136353738303633373332323134663831343631303332393537383036333736663738363231313436313033333035373566383066643562383036333031313861636334313436313031633535373830363330333232656439383134363130316461353738303633313531643330643131343631303165643537383036333165633434373234313436313032306335373830363332306439316237613134363130323166353738303633323565316337373631343631303233323537383036333265323139346438313436313032343535373830363333353435356465643134363130323730353735623566383066643562363130316438363130316433333636303034363134396432353635623631303461353536356230303562363130316438363130316538333636303034363134613064353635623631303464613536356236313031663536303061383135363562363034303531363066663930393131363831353236303230303135623630343035313830393130333930663335623631303164383631303231613336363030343631343964323536356236313037353635363562363130316438363130323264333636303034363134613234353635623631303736313536356236313031643836313032343033363630303436313461373235363562363130643362353635623631303235383631303235333336363030343631346130643536356236313064616635363562363034303531363030313630303136303430316230333930393131363831353236303230303136313032303335363562363130323739363132373130383135363562363034303531363166666666393039313136383135323630323030313631303230333536356236313031643836313032396133363630303436313439643235363562363130653033353635623631303164383631303261643336363030343631346139333536356236313065306535363562363130326335363130326330333636303034363134616431353635623631306562653536356236303430353139303831353236303230303136313032303335363562363130316438363130326531333636303034363134623338353635623631306565353536356236313032656536303134383135363562363034303531363366666666666666663930393131363831353236303230303136313032303335363562363130326335363130333131333636303034363134623630353635623631313161643536356236313032353836313033323433363630303436313461306435363562363131316364353635623631303263353566383135363562363130316438363130333365333636303034363134396432353635623631313165313536356236313031663536303330383135363562363130316438363130333539333636303034363134613064353635623631313230633536356236313031643836313033366333363630303436313462333835363562363131326339353635623631303263353631303337663336363030343631346237623536356236313134663535363562363130316438363130333932333636303034363134613933353635623631313531333536356236313032633537663433313737313366376563626464646434626339396539356439303361646564616138383362326537633235353136313062643133653263376534373364303038313536356236313033636336303035363030313630393931623031383135363562363034303531363130323033393139303631346239623536356236313031643836313033653733363630303436313461306435363562363131373035353635623631303263353766653932353436643639383935306464643338393130643265313565643164393233636430613762336464653965326136613366333830353635353539636230303831353635623631303164383631303432313336363030343631346130643536356236313139363535363562363130316635363031343831353635623631303434313631303433633336363030343631346130643536356236313161383235363562363034303531363130323033393139303631346332343536356236313032353836323032613330303831353635623631303263353766366535626466636365313565353363333430366561363762666365333764636432366635313532643534393238323465343366643565336338616335616230303831353635623631303164383631303438643336363030343631346362383536356236313162633435363562363130326335363130346130333636303034363134636636353635623631316361333536356236313034623038333833383336313163646235363562363130346435353736303430353136333130333663663931363065313162383135323630303438313031383439303532363032343031356236303430353138303931303339306664356235303530353035363562356636313034653336313230323835363562356638333831353236303037383230313630323035323630343038303832323038313531363065303831303139303932353238303534393339343530393139323930393139303832393036306666313636303035383131313135363130353163353736313035316336313462616635363562363030353831313131353631303532643537363130353264363134626166353635623831353236303230303136303031383230313830353436313035343139303631346436313536356238303630316630313630323038303931303430323630323030313630343035313930383130313630343035323830393239313930383138313532363032303031383238303534363130353664393036313464363135363562383031353631303562383537383036303166313036313035386635373631303130303830383335343034303238333532393136303230303139313631303562383536356238323031393139303566353236303230356632303930356238313534383135323930363030313031393036303230303138303833313136313035396235373832393030333630316631363832303139313562353035303530393138333532353035303630303238323031353436303031363030313630343031623033383038323136363032303834303135323630303136303430316238323034383131363630343038343031353236303031363038303162383230343831313636303630383430313532363030313630633031623930393130343831313636303830383330313532363030333932383330313534313636306130393039313031353239303931353038313531363030353831313131353631303632333537363130363233363134626166353635623134363130363536353735663833383135323630303738333031363032303532363034303930383139303230353439303531363331373063633933333630653231623831353236313034636339313630666631363930363030343031363134643939353635623630363038313031353136303430353136333432613265306235363065313162383135323630303438313031383539303532363030313630303136303430316230333930393131363630323438323031353235663630343438323031353236303035363030313630393931623031393036336565356234386562393037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f393036333835343563313661393036303634303135663630343035313830383330333831383635616634313538303135363130366364353733643566383033653364356666643562353035303530353036303430353133643566383233653630316633643930383130313630316631393136383230313630343035323631303666343931393038313031393036313465613235363562363034303531383236336666666666666666313636306530316238313532363030343031363130373130393139303631346564333536356236303230363034303531383038333033383135663837356166313135383031353631303732633537336435663830336533643566666435623530353035303530363034303531336436303166313936303166383230313136383230313830363034303532353038313031393036313037353039313930363134656535353635623530353035303530353635623631303735303833383338333631316364623536356235663631303736613631323032383536356236303039383130313534393039313530363066663136313536313037393335373630343035313633376661623831653536306530316238313532363030343031363034303531383039313033393066643562363030353630303136303939316230313630303136303031363061303162303331363633343231336366373836303430353138313633666666666666666631363630653031623831353236303034303136303230363034303531383038333033383138363561666131353830313536313037643635373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130376661393139303631346565353536356238333630323030313335313436313038323335373630343035313633373262306137653736306531316238313532363032303834303133353630303438323031353236303234303136313034636335363562333036313038333436303630383530313630343038363031363134656663353635623630303136303031363061303162303331363134363130383663353736313038353236303630383430313630343038353031363134656663353635623630343035313633326638383132306436306532316238313532363030343031363130346363393139303631346239623536356235663631303837613630363038353031383536313466313735363562393035303930353035663830356238323831363366666666666666663136313031353631306236313537356636313038396436303630383830313838363134663137353635623833363366666666666666663136383138313130363130386233353736313038623336313466356335363562393035303630323030323831303139303631303863353931393036313466373035363562363130386365393036313466646235363562383035313630343035313931393235303566393136303038383830313931363130386536393136313530353635363562393038313532363032303031363034303531383039313033393032303534313436313039313635373830353136303430353136336134316637373266363065303162383135323631303463633931393036303034303136313465643335363562356636303032383835663031333538343630343035313630323030313631303934353932393139303931383235323630653031623630303136303031363065303162303331393136363032303832303135323630323430313930353635623630343038303531363031663139383138343033303138313532393038323930353236313039356639313631353035363536356236303230363034303531383038333033383138353561666131353830313536313039376135373364356638303365336435666664356235303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631303939643931393036313465653535363562393035303830383636303038303138333566303135313630343035313631303962353931393036313530353635363562393038313532363034303830353136303230393238313930303338333031383132303933393039333535363065303833303138313532363030323833353238343531383238343031353238343831303138303531363030313630303136303430316230333930383131363835383430313532356636303630383630313831393035323931353138313136363038303836303135323432313636306130383530313532363063303834303138313930353238343831353236303037386130313930393235323930323038313531383135343832393036306666313931363630303138333630303538313131313536313061333735373631306133373631346261663536356230323137393035353530363032303832303135313630303138323031393036313061353039303832363135306162353635623530363034303832383130313531363030323833303138303534363036303836303135313630383038373031353136306130383830313531363030313630303136303430316230333935383631363630303136303031363038303162303331393930393431363933393039333137363030313630343031623932383631363932393039323032393139303931313736303031363030313630383031623033313636303031363038303162393138353136393139303931303236303031363030313630633031623033313631373630303136306330316239313834313639313930393130323137393035353630633039303933303135313630303339303932303138303534363030313630303136303430316230333139313639323834313639323930393231373930393135353833303135313631306166343931313638353631353137613536356238323531363034303531393139353530363130623035393136313530353635363562363034303830353139313832393030333832323039303834303135313630303136303031363034303162303331363832353239303832393037663964343766656639646130373736363135343665363436643631383330626663626461393035303663326535656564333831393565383263346562316362646639303630323030313630343035313830393130333930613335303530383036313062356139303631353138643536356239303530363130383831353635623530363030343833303138313930353536303031383330313534363036343930363130623835393036303031363034303162393030343630666631363833363135316166353635623130313536313062613735373630343035313633353934333331376636306530316238313532363030343831303138323930353236303234303136313034636335363562356637335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363331653664393738393631306263623837363132303463353635623630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363130626562393139303631346564333536356236303230363034303531383038333033383138363561663431353830313536313063303635373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130633261393139303631346565353536356239303530356637335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f3633383632626661363338383630343035313832363366666666666666663136363065303162383135323630303430313631306336343931393036313532663135363562356636303430353138303833303338313836356166343135383031353631306337653537336435663830336533643566666435623530353035303530363034303531336435663832336536303166336439303831303136303166313931363832303136303430353236313063613539313930383130313930363134656132353635623930353035663630303238323630343035313631306362383931393036313530353635363562363032303630343035313830383330333831383535616661313538303135363130636433353733643566383033653364356666643562353035303530363034303531336436303166313936303166383230313136383230313830363034303532353038313031393036313063663639313930363134656535353635623930353038323831313436313064323235373630343035313633313837326663386436306530316238313532363030343831303138323930353236303234383130313834393035323630343430313631303463633536356235303530353036303039393039323031383035343630666631393136363030313137393035353530353035303530353635623631306434343832363132313537353635623631306436343537363034303531363333306566613938623630653031623831353236303034383130313833393035323630323430313631303463633536356235663631306436653833363131613832353635623531393035303630303238313630303538313131313536313064383535373631306438353631346261663536356231343631306461353537383036303430353136333137306363393333363065323162383135323630303430313631303463633931393036313464393935363562363130373530383338333631323138303536356235663830363130646239363132343361353635623630303330313534363130646337393038343631353337633536356239303530383031353830363130646463353735303630303136303031363034303162303338313131356231353631306466643537363034303531363332323264313634333630653231623831353236303034383130313834393035323630323430313631303463633536356239323931353035303536356236313037353038333833383336313234356535363562363130653136363132363239353635623566363130653166363132343361353635623930353035663830363130653263383436313236373335363562393135303931353036313065333938323631323135373536356236313065343535373530353035303631306562333536356235663832383135323630303538343031363032303532363034303930323035343630303136303031363061303162303331363630303438323531363030353831313131353631306537333537363130653733363134626166353635623033363130653938353735663833383135323630303838353031363032303532363034303831323038303534393139303535363130653936383238323631326132323536356235303562363130656165383136313065613938343630343030313531363131316164353635623631326139323536356235303530353035303562363130656262363132616230353635623530353635623566363130656337363132363239353635623631306564333835383538353835363132616436353635623930353036313065646436313261623035363562393439333530353035303530353635623566363130656565363132343361353635623566383338313532363030363832303136303230353236303430383038323230383135313630653038313031393039323532383035343933393435303931393239303931393038323930363066663136363030333831313131353631306632373537363130663237363134626166353635623630303338313131313536313066333835373631306633383631346261663536356238313532383135343631303130303930303436303031363030313630613031623033313636303230383230313532363030313832303135343630343038303833303139313930393135323630303239303932303135343630303136303031363034303162303338303832313636303630383430313532363030313630343031623832303438313136363038303834303135323630303136303830316238323034383131363630613038343031353236303031363063303162393039313034313636306330393039313031353238313031353139303931353035663631306661653832363131613832353635623930353036303031383335313630303338313131313536313066633535373631306663353631346261663536356231343631306665363537383235313630343035313633336230643534306436306532316238313532363130346363393139303630303430313631353339623536356236303034383135313630303538313131313536313066666235373631306666623631346261663536356230333631313031313537363131303039383536313263386435363562353035303530353035303530353635623566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363339646532336434303631313033363861363132303463353635623630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363131303536393139303631346564333536356236303630363034303531383038333033383138363561663431353830313536313130373135373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363131303935393139303631353362353536356235303931353039313530383138343134363131306332353738343630343030313531363034303531363330383939333862333630653131623831353236303034303136313034636339313831353236303230303139303536356238303630303136303031363034303162303331363833363036303031353136303031363030313630343031623033313631303830363131306662353735303830363030313630303136303430316230333136383536306130303135313630303136303031363034303162303331363131356231353631313132343537363034303531363332653139626332643630653131623831353236303031363030313630343031623033383231363630303438323031353236303234303136313034636335363562356638373831353236303036383730313630323039303831353236303430393138323930323038303534363066663139313636303032393038313137383235353031383035343630303136303031363034303162303334323136363030313630343031623831303236376666666666666666666666666666666636303430316231393930393231363931393039313137393039313535393135313931383235323835393138393931376630343730353962343635303639623862373531383336623431663966316438336461666635383364323233386363376662623436313433376563323361346636393130313630343035313830393130333930613335303530353035303530353035303530353635623566363131316236363132343361353635623630303330313534363130646664393036303031363030313630343031623033383431363631353161663536356235663631313164373832363131613832353635623630383030313531393239313530353035363562363131316563383338333833363132343565353635623631303464353537363034303531363335626666363833663630653131623831353236303034383130313834393035323630323430313631303463633536356235663631313231353631323433613536356239303530356636313132323138333631316138323536356235313930353036303034383136303035383131313135363131323338353736313132333836313462616635363562313436313132353835373830363034303531363331373063633933333630653231623831353236303034303136313034636339313930363134643939353635623566383338313532363030353833303136303230353236303430393032303534363030313630303136306130316230333136333331343631313239333537333335623630343035313633366532636364373536306531316238313532363030343031363130346363393139303631346239623536356235663833383135323630303838333031363032303930383135323630343038303833323038303534393038343930353536303035383630313930393235323930393132303534363130373530393036303031363030313630613031623033313638323631326132323536356236313132643136313236323935363562356636313132646136313234336135363562356638333831353236303036383230313630323035323630343038303832323038313531363065303831303139303932353238303534393339343530393139323930393139303832393036306666313636303033383131313135363131333133353736313133313336313462616635363562363030333831313131353631313332343537363131333234363134626166353635623831353238313534363130313030393030343630303136303031363061303162303331363630323038323031353236303031383230313534363034303832303135323630303239303931303135343630303136303031363034303162303338303832313636303630383430313532363030313630343031623832303438313136363038303834303135323630303136303830316238323034383131363630613038343031353236303031363063303162393039313034313636306330393039313031353239303530363030333831353136303033383131313135363131333964353736313133396436313462616635363562313436313133626535373830353136303430353136333362306435343064363065323162383135323631303463633931393036303034303136313533396235363562363030343631313363643832363034303031353136313161383235363562353136303035383131313135363131336466353736313133646636313462616635363562313436313134646535373566363131336565383536313230346335363562393035303566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f3633396465323364343038343630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363131343264393139303631346564333536356236303630363034303531383038333033383138363561663431353830313536313134343835373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363131343663393139303631353362353536356235303931353039313530383138343630343030313531313436313134393835373630343035313633303839393338623336306531316238313532363030343831303138333930353236303234303136313034636335363562383036303031363030313630343031623033313638343630633030313531363030313630303136303430316230333136313131353631313464613537363034303531363332653139626332643630653131623831353236303031363030313630343031623033383231363630303438323031353236303234303136313034636335363562353035303530356236313134653738333631326338643536356235303530363131346631363132616230353635623530353035363562356636313134666536313236323935363562363131353039383333333834363132656164353635623930353036313064666436313261623035363562356636313135316336313230323835363562393035303566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363332653433636562353631313534333836363132303463353635623630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363131353633393139303631346564333536356236303430383035313830383330333831383635616634313538303135363131353764353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631313561313931393036313533663535363562393135303931353038303631313563373537363034303531363332643037313335333630653031623831353238313135313536303034383230313532363032343031363130346363353635623566383238313532363030363834303136303230353236303430393032303830353436313135653139303631346436313536356239303530356630333631313630353537363034303531363330383939333862333630653131623831353236303034383130313833393035323630323430313631303463633536356236303031356638333831353236303037383530313630323035323630343039303230353436306666313636303035383131313135363131363262353736313136326236313462616635363562313436313136356535373566383238313532363030373834303136303230353236303430393038313930323035343930353136333137306363393333363065323162383135323631303463633931363066663136393036303034303136313464393935363562356638323831353236303036383430313630323035323630343038313230363131363736393136313439323135363562356638323831353236303037383430313630323039303831353236303430393138323930323038303534363066663139313636303032393038313137383235353031383035343630303136303031363034303162303334323831383131363630303136306330316230323630303136303031363063303162303339303933313639323930393231373932383339303535383435313630303136303830316239303933303431363832353239313831303139313930393135323833393137666638666431633930666239636661326361323335386664663538303662303836616434333331356439326232323163393239656663376631303563653735363839313031363034303531383039313033393061323530353035303530353635623566363131373065363132343361353635623566383338313532363030363832303136303230353236303430383038323230383135313630653038313031393039323532383035343933393435303931393239303931393038323930363066663136363030333831313131353631313734373537363131373437363134626166353635623630303338313131313536313137353835373631313735383631346261663536356238313532383135343631303130303930303436303031363030313630613031623033313636303230383230313532363030313830383330313534363034303833303135323630303239303932303135343630303136303031363034303162303338303832313636303630383430313532363030313630343031623832303438313136363038303834303135323630303136303830316238323034383131363630613038343031353236303031363063303162393039313034313636306330393039313031353239303931353038313531363030333831313131353631313764313537363131376431363134626166353635623134313538303135363131376632353735303630303338313531363030333831313131353631313765663537363131376566363134626166353635623134313535623135363131383133353738303531363034303531363333623064353430643630653231623831353236313034636339313930363030343031363135333962353635623566363131383231383236303430303135313631316138323536356239303530383036303630303135313630303136303031363034303162303331363566303336313138353335373630343035313633333962383934663936306532316238313532363030343831303138353930353236303234303136313034636335363562363034303830383330313531363036303833303135313630383038343031353139323531363334326132653062353630653131623831353236303035363030313630393931623031393336336565356234386562393337335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f39333633383534356331366139333631313863313933393036303034303139323833353236303031363030313630343031623033393138323136363032303834303135323136363034303832303135323630363030313930353635623566363034303531383038333033383138363561663431353830313536313138646235373364356638303365336435666664356235303530353035303630343035313364356638323365363031663364393038313031363031663139313638323031363034303532363131393032393139303831303139303631346561323536356236303430353138323633666666666666666631363630653031623831353236303034303136313139316539313930363134656433353635623630323036303430353138303833303338313566383735616631313538303135363131393361353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631313935653931393036313465653535363562353035303530353035303536356235663631313936653631323032383536356235663833383135323630303638323031363032303532363034303930323038303534393139323530393036313139386339303631346436313536356239303530356630333631313962303537363034303531363330383939333862333630653131623831353236303034383130313833393035323630323430313631303463633536356236303031356638333831353236303037383330313630323035323630343039303230353436306666313636303035383131313135363131396436353736313139643636313462616635363562313436313161303935373566383238313532363030373832303136303230353236303430393038313930323035343930353136333137306363393333363065323162383135323631303463633931363066663136393036303034303136313464393935363562356638323831353236303036383230313630323035323630343039303831393032303930353136336565356234386562363065303162383135323630303536303031363039393162303139313633656535623438656239313631316134323931393036303034303136313534313835363562363032303630343035313830383330333831356638373561663131353830313536313161356535373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130346435393139303631346565353536356236313161386136313439353835363562356636313161393336313230323835363562356638343831353236303037383230313630323035323630343039303831393032303831353136306530383130313930393235323830353439323933353039303931383239303630666631363630303538313131313536313161636135373631316163613631346261663536356236303035383131313135363131616462353736313161646236313462616635363562383135323630323030313630303138323031383035343631316165663930363134643631353635623830363031663031363032303830393130343032363032303031363034303531393038313031363034303532383039323931393038313831353236303230303138323830353436313162316239303631346436313536356238303135363131623636353738303630316631303631316233643537363130313030383038333534303430323833353239313630323030313931363131623636353635623832303139313930356635323630323035663230393035623831353438313532393036303031303139303630323030313830383331313631316234393537383239303033363031663136383230313931356235303530353039313833353235303530363030323832303135343630303136303031363034303162303338303832313636303230383430313532363030313630343031623832303438313136363034303834303135323630303136303830316238323034383131363630363038343031353236303031363063303162393039313034383131363630383038333031353236303033393039323031353439303931313636306130393039313031353239333932353035303530353635623766663063353765313638343064663034306631353038386463326638316665333931633339323362656337336532336139363632656663396332323963366130303830353436303032393139303630303136303430316239303034363066663136383036313163306435373530383035343630303136303031363034303162303338303834313639313136313031353562313536313163326235373630343035313633663932656538613936306530316238313532363030343031363034303531383039313033393066643562383035343638666666666666666666666666666666666666313931363630303136303031363034303162303338333136313736303031363034303162313738313535363131633536383438343631333138323536356238303534363066663630343031623139313638313535363034303531363030313630303136303430316230333833313638313532376663376635303562326633373161653231373565653439313366343439396531663236333361376235393336333231656564316364616562363131353138316432393036303230303136303430353138303931303339306131353035303530353035363562356638303631316361643631323032383536356239303530383036303038303138343834363034303531363131636333393239313930363135346132353635623930383135323630323030313630343035313830393130333930323035343931353035303932393135303530353635623566383036313163653536313234336135363562356638363831353236303036383230313630323035323630343038303832323038313531363065303831303139303932353238303534393339343530393139323930393139303832393036306666313636303033383131313135363131643165353736313164316536313462616635363562363030333831313131353631316432663537363131643266363134626166353635623831353238313534363130313030393030343630303136303031363061303162303331363630323038323031353236303031383230313534363034303830383330313931393039313532363030323930393230313534363030313630303136303430316230333830383231363630363038343031353236303031363034303162383230343831313636303830383430313532363030313630383031623832303438313136363061303834303135323630303136306330316239303931303431363630633039303931303135323831303135313930393135303566363131646135383236313161383235363562393035303630303238333531363030333831313131353631316462633537363131646263363134626166353635623134363131646464353738323531363034303531363333623064353430643630653231623831353236313034636339313930363030343031363135333962353635623630323038333031353136303031363030313630613031623033313633333134363131653739353735663832383135323630303538353031363032303532363034303930323035343630303136303031363061303162303331363333313436313165313635373333363131323739353635623566383238313532363030353835303136303230353236303430393032303534363061303832303135313631316534353931363030313630623031623930303436303031363030313630343031623033313639303631353462313536356236303031363030313630343031623033313634323130313536313165373935373630343035313633666236636536336636306530316238313532363030313630303136303430316230333432313636303034383230313532363032343031363130346363353635623630303238313531363030353831313131353631316538653537363131653865363134626166353635623033363131666263353736303032383430313534363038303834303135313631316562303931363030313630303136303430316230333136393036313534623135363562363030313630303136303430316230333136343231303135363131656534353736303430353136336662366365363366363065303162383135323630303136303031363034303162303334323136363030343832303135323630323430313631303463633536356238363135363131656636353736313165663438323837363132313830353635623530356235663838383135323630303638353031363032303532363034303930323038303534363066663139313636303033313739303535363036303833303135313630383038323031353136313166326639313834393136313166326139313930363135346431353635623631333139633536356235303566383938313532363030363836303136303230353236303430383132303630303230313830353436303031363030313630343031623033393039333136363030313630633031623032363030313630303136306330316230333930393331363932393039323137393039313535363131663730383436313333363635363562356638613831353236303037383730313630323035323630343038303832323038333930353535313931393235303834393138623931376633363664333336633061623338306463373939663039356136663832613236333236353835633532393039636336393862303962613435343037303965643537393161333135313539343530363132303231393335303530353035303536356236303034383135313630303538313131313536313166643135373631316664313631346261663536356230333631323030353537363131666466383336313333363635363562356638393831353236303037383630313630323035323630343039303230353536313166663838383631326338643536356236303031393435303530353035303530363132303231353635623830353136303430353136333137306363393333363065323162383135323631303463633931393036303034303136313464393935363562393339323530353035303536356237666539323534366436393839353064646433383931306432653135656431643932336364306137623364646539653261366133663338303536353535396362303039303536356236303430383035313630363038303832303138333532356638303833353236303230383330313532393138313031393139303931353236303430353136333036663832353335363065343162383135323633666666666666666638333136363030343832303135323566393038313930363030353630303136303939316230313930363336663832353335303930363032343031356636303430353138303833303338313836356166613135383031353631323062303537336435663830336533643566666435623530353035303530363034303531336435663832336536303166336439303831303136303166313931363832303136303430353236313230643739313930383130313930363135346631353635623931353039313530383036313230663935373630343035313633366232663139653936306530316238313532363030343031363034303531383039313033393066643562383135313135363132313166353738313531363034303531363336626135383961353630653031623831353236303034383130313931393039313532363032343031363130346363353635623630323038323031353136303031363030313630613031623033313631353631323135303537383136303230303135313630343035313632346465373564363065333162383135323630303430313631303463633931393036313462396235363562353039323931353035303536356235663830363132313631363132343361353635623566393338343532363030353031363032303532353035303630343039303230353436303031363030313630613031623033313631353135393035363562363034303531363330366638323533353630653431623831353236336666666666666666383231363630303438323031353235663930383139303831393036303035363030313630393931623031393036333666383235333530393036303234303135663630343035313830383330333831383635616661313538303135363132316362353733643566383033653364356666643562353035303530353036303430353133643566383233653630316633643930383130313630316631393136383230313630343035323631323166323931393038313031393036313534663135363562393135303931353038303631323231343537363034303531363336623266313965393630653031623831353236303034303136303430353138303931303339306664356236303035363030313630393931623031363030313630303136306130316230333136363334323133636637383630343035313831363366666666666666663136363065303162383135323630303430313630323036303430353138303833303338313836356166613135383031353631323235373537336435663830336533643566666435623530353035303530363034303531336436303166313936303166383230313136383230313830363034303532353038313031393036313232376239313930363134656535353635623832353131343631323261313537383135313630343035313633366261353839613536306530316238313532363030343831303139313930393135323630323430313631303463633536356236303230383230313531363030313630303136306130316230333136313536313232643235373831363032303031353136303430353136323464653735643630653331623831353236303034303136313034636339313930363134623962353635623566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363330383863323436333835363034303031353136303430353138323633666666666666666631363630653031623831353236303034303136313233306639313930363134656433353635623630343038303531383038333033383138363561663431353830313536313233323935373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363132333464393139303631353538313536356239313530393135303831383731343631323337343537363034303531363330383939333862333630653131623831353236303034383130313838393035323630323430313631303463633536356235663631323337643631323433613536356235663839383135323630303538323031363032303532363034303930323036303031303135343930393135303630303136303031363034303162303339303831313639303833313631313135363132343130353735663838383135323630303538323031363032303930383135323630343039313832393032303630303130313830353436303031363030313630343031623033313931363630303136303031363034303162303338363136393038313137393039313535393135313931383235323839393137666563343431343865386666323731663264306261636566313134323135346162616362306162623361323965623365623530653263613937653836643034333539313031363034303531383039313033393061323631323432663536356235663838383135323630303538323031363032303532363034303930323036303031303135343630303136303031363034303162303331363931353035623530393639353530353035303530353035303536356237663433313737313366376563626464646434626339396539356439303361646564616138383362326537633235353136313062643133653263376534373364303039303536356235663830363132343638363132343361353635623930353035663631323437343836363133346539353635623930353036313234376638363631323135373536356236313234386535373630303139323530353035303631323032313536356235663836383135323630303538333031363032303532363034303930323035343630303136303031363061303162303331363333313436313234623335373333363131323739353635623566383638313532363030353833303136303230353236303430393032303534363061303832303135313631323465323931363030313630623031623930303436303031363030313630343031623033313639303631353462313536356236303031363030313630343031623033313638313630633030313531363030313630303136303430316230333136313031353631323532393537363063303831303135313630343035313633666236636536336636306530316238313532363030313630303136303430316230333930393131363630303438323031353236303234303136313034636335363562356638353135363132353431353736313235336138373836363132313830353635623930353036313235356635363562353035663836383135323630303538333031363032303532363034303930323036303031303135343630303136303031363034303162303331363562363030343833303135343630343038333031353135663931363030313630303136306130316230333136393036333466323234323966393036313235383439303631313161643536356236306130383630313531363063303837303135313630343035313630303136303031363065303162303331393630653038363930316231363831353236313235623439333932393138323931383939303630303430313631353561343536356236303230363034303531383038333033383138363561666131353830313536313235636635373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363132356633393139303631346565353536356239303530383038343630303830313566386138313532363032303031393038313532363032303031356632303566383238323534363132363137393139303631353137613536356239303931353535303530313531353937393635303530353035303530353035303536356237663962373739623137343232643064663932323233303138623332623464316661343665303731373233643638313765323438366430303362656363353566303038303534363030313139303136313236366435373630343035313633336565356165623536306530316238313532363030343031363034303531383039313033393066643562363030323930353535363562356636313236376336313439353835363562356636313236383536313230323835363562393035303566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363332653433636562353631323661633838363132303463353635623630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363132366363393139303631346564333536356236303430383035313830383330333831383635616634313538303135363132366536353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631323730613931393036313533663535363562393135303931353038303135363132373331353736303430353136333264303731333533363065303162383135323831313531353630303438323031353236303234303136313034636335363562356638323831353236303037383430313630323035323630343038303832323038313531363065303831303139303932353238303534383239303630666631363630303538313131313536313237363235373631323736323631346261663536356236303035383131313135363132373733353736313237373336313462616635363562383135323630323030313630303138323031383035343631323738373930363134643631353635623830363031663031363032303830393130343032363032303031363034303531393038313031363034303532383039323931393038313831353236303230303138323830353436313237623339303631346436313536356238303135363132376665353738303630316631303631323764353537363130313030383038333534303430323833353239313630323030313931363132376665353635623832303139313930356635323630323035663230393035623831353438313532393036303031303139303630323030313830383331313631323765313537383239303033363031663136383230313931356235303530353039313833353235303530363030323832303135343630303136303031363034303162303338303832313636303230383430313532363030313630343031623832303438313136363034303834303135323630303136303830316238323034383131363630363038343031353236303031363063303162393039313034383131363630383038333031353236303033393238333031353431363630613039303931303135323930393135303831353136303035383131313135363132383639353736313238363936313462616635363562313431353830313536313238386135373530363030313831353136303035383131313135363132383837353736313238383736313462616635363562313431353562313536313238616235373830353136303430353136333137306363393333363065323162383135323631303463633931393036303034303136313464393935363562363030333831353136303035383131313135363132386330353736313238633036313462616635363562303336313238636535373630303438313532363132386433353635623630303538313532356238333630303830313831363032303031353136303430353136313238653939313930363135303536353635623930383135323630343038303531363032303932383139303033383330313930323035663930383139303535383538313532363030373837303139303932353239303230383135313831353438333932393139303832393036306666313931363630303138333630303538313131313536313239326435373631323932643631346261663536356230323137393035353530363032303832303135313630303138323031393036313239343639303832363135306162353635623530363034303832303135313630303238323031383035343630363038353031353136303830383630313531363061303837303135313630303136303031363034303162303339353836313636303031363030313630383031623033313939303934313639333930393331373630303136303430316239323836313639323930393230323931393039313137363030313630303136303830316230333136363030313630383031623931383531363931393039313032363030313630303136306330316230333136313736303031363063303162393138343136393139303931303231373930353536306330393039323031353136303033393039313031383035343630303136303031363034303162303331393136393139303932313631373930353538303531363030353831313131353631323965623537363132396562363134626166353635623630343035313834393037663163303865353936353666316131386463326461373638323663646335323830356334336538393761313763353066616566623861623363313532366363313639303566393061333931393639313935353039303933353035303530353035363562356636313261326236313337633135363562383035343630343035313633343063313066313936306530316238313532363030313630303136306130316230333836383131363630303438333031353236303234383230313836393035323932393335303931313639303633343063313066313939303630343430313566363034303531383038333033383135663837383033623135383031353631326137373537356638306664356235303561663131353830313536313261383935373364356638303365336435666664356235303530353035303530353035303536356236313134663138323832363132613966363133376331353635623534363030313630303136306130316230333136393139303631333765353536356236303031376639623737396231373432326430646639323232333031386233326234643166613436653037313732336436383137653234383664303033626563633535663030353535363562356638303631326165303631323433613536356236303032383130313534393039313530363166666666363030313630343031623930393130343831313639303836313631303830363132623039353735303631323731303631666666663836313631313562313536313262326435373630343035313633356631326536633336306531316238313532363166666666383631363630303438323031353236303234303136313034636335363562363030323831303135343630303136303031363034303162303339303831313639303835313631303135363132623639353736303430353136323032613036643630653131623831353236303031363030313630343031623033383531363630303438323031353236303234303136313034636335363562383035343833313038303631326237623537353038303630303130313534383331313562313536313262396335373630343035313633323232643136343336306532316238313532363030343831303138343930353236303234303136313034636335363562356636313262613638343631333834343536356239303530356636313262623238323631306461663536356239303530356636313262626638393833363133383631353635623930353036303430353138303630383030313630343035323830363132626433333339303536356236303031363030313630613031623033393038313136383235323631666666663830386331363630323038303835303139313930393135323630303136303031363034303162303338303864313636303430383038373031393139303931353235663630363039363837303138313930353238383831353236303035393039623031383335323939386139303230383635313831353439333838303135313962383830313531383331363630303136306230316230323637666666666666666666666666666666663630623031623139396339303935313636303031363061303162303236303031363030313630623031623033313939303934313639353136393439303934313739313930393131373938393039383136313738313535393130313531363030313930393130313830353439313930393531363630303136303031363034303162303331393930393131363137393039333535353039303931353035303934393335303530353035303536356235663631326339363631323433613536356235663833383135323630303638323031363032303532363034303830383232303831353136306530383130313930393235323830353439333934353039313932393039313930383239303630666631363630303338313131313536313263636635373631326363663631346261663536356236303033383131313135363132636530353736313263653036313462616635363562383135323831353436313031303039303034363030313630303136306130316230333136363032303832303135323630303138323031353436303430383038333031393139303931353236303032393039323031353436303031363030313630343031623033383038323136363036303834303135323630303136303430316238323034383131363630383038343031353236303031363038303162383230343831313636306130383430313532363030313630633031623930393130343136363063303930393130313532383130313531393039313530363132643534363133653265353635623832363038303031353136313264363339313930363135346231353635623630303136303031363034303162303331363432313031353631326439373537363034303531363366623663653633663630653031623831353236303031363030313630343031623033343231363630303438323031353236303234303136313034636335363562356638343831353236303036383430313630323039303831353236303430383038333230383035343630303136303031363061383162303331393136383135353630303138313031383439303535363030323031383339303535363030373836303139303931353238313230383035343930383239303535393038303832313536313265353235373566383438313532363030353837303136303230353236303430393032303534363132373130393036313265303439303630303136306130316239303034363166666666313638353631353161663536356236313265306539313930363135333763353635623931353038313836363030383031356638363831353236303230303139303831353236303230303135663230356638323832353436313265333239313930363135313761353635623930393135353530363132653432393035303832383436313535643235363562393035303631326535323835363032303031353138323631326132323536356236313265363738353630323030313531363130656139383736303630303135313631313161643536356236303430383035313832383135323630323038313031383439303532383539313839393137663865636563663531303037306333323064396135353332336666616265333530653239346165353035666330633530396463353733366461366635636339393339313031363034303531383039313033393061333530353035303530353035303530353635623566383036313265623736313234336135363562393035303566363132656336363130323533383536313338343435363562393035303566363132656432383736313161383235363562393035303631326564643837363132313537353635623631326566643537363034303531363333306566613938623630653031623831353236303034383130313838393035323630323430313631303463633536356236303032383135313630303538313131313536313266313235373631326631323631346261663536356231343631326633333537383035313630343035313633313730636339333336306532316238313532363130346363393139303630303430313631346439393536356235663832383236303830303135313631326634343931393036313534623135363562393035303833363030323031363030613930353439303631303130303061393030343630303136303031363034303162303331363832363034303031353136313266366439313930363135356535353635623630303136303031363034303162303331363831363030313630303136303430316230333136313131353631326661613537363034303531363336643531666530353630653131623831353236303031363030313630343031623033383231363630303438323031353236303234303136313034636335363562356638303631326662363861383436313331396335363562393135303931353035663861383336303430353136303230303136313266653439323931393039313832353236306330316236303031363030313630633031623033313931363630323038323031353236303238303139303536356236303430383035313630316631393831383430333031383135323832383235323830353136303230393039313031323036306530383330313930393135323931353038303630303138313532363030313630303136306130316230333863313636303230383038333031393139303931353236303430383038333031386639303532363030313630303136303430316230333830386231363630363038353031353235663630383038353031383139303532393038383136363061303835303135323630633039303933303138333930353238343833353236303036386230313930393135323930323038313531383135343832393036306666313931363630303138333630303338313131313536313330373735373631333037373631346261663536356230323137393035353530363032303832383130313531383235343631303130303630303136306138316230333139313636313031303036303031363030313630613031623033393238333136303231373833353536303430383038353031353136303031383530313535363036303830383630313531363030323930393530313830353436303830383038393031353136306130386130313531363063303930396130313531363030313630303136303430316230333939386131363630303136303031363038303162303331393930393431363933393039333137363030313630343031623931386131363931393039313032313736303031363030313630383031623033313636303031363038303162393938393136393939303939303236303031363030313630633031623033313639383930393831373630303136306330316239313838313639313930393130323137393035353831353138393836313638313532386138363136393438313031393439303934353239333862313639303833303135323931383130313835393035323930386331363931386439313834393137666230303234623236336263336130623732386136656465613530613639656661383431313839663864333265653861663964316332623161316132323334323639313031363034303531383039313033393061343961393935303530353035303530353035303530353035303536356236313331386136313365343935363562363133313933383236313365393435363562363131346631383136313366303835363562356638303566363133316137363132303238353635623566383638313532363030373832303136303230353236303430393032303630303230313534393039313530363030313630383031623930303436303031363030313630343031623033313636313331643738353832363133663634353635623566363133316531383736313431383935363562356638383831353236303037383530313630323035323630343038303832323036303032303138303534363766666666666666666666666666666666363038303162313931363630303136303830316236303031363030313630343031623033386338313136393138323032393239303932313739303932353539313531363334326132653062353630653131623831353236303034383130313863393035323931383431363630323438333031353236303434383230313532393139323530393036303035363030313630393931623031393036336565356234386562393037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f39303633383534356331366139303630363430313566363034303531383038333033383138363561663431353830313536313332386135373364356638303365336435666664356235303530353035303630343035313364356638323365363031663364393038313031363031663139313638323031363034303532363133326231393139303831303139303631346561323536356236303430353138323633666666666666666631363630653031623831353236303034303136313332636439313930363134656433353635623630323036303430353138303833303338313566383735616631313538303135363133326539353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631333330643931393036313465653535363562363034303830353136303031363030313630343031623033386138313136383235323630323038323031383439303532383235313933393435303835313639323862393237663037646535666633356136373461383030356536363166333333336339303763613633333334363238303837363264313964633762336162623161386331646639323832393030333031393061333930393435303932353035303530356239323530393239303530353635623566383036313333373036313234336135363562393035303566363133333830383436303430303135313631316138323536356239303530356636303033383235313630303538313131313536313333393835373631333339383631346261663536356231343830363133336236353735303630303438323531363030353831313131353631333362343537363133336234363134626166353635623134356231353631333363363537353036306330383130313531363133343033353635623630303238323531363030353831313131353631333364623537363133336462363134626166353635623033363133336537353735303432363133343033353635623831353136303430353136333137306363393333363065323162383135323631303463633931393036303034303136313464393935363562383436303830303135313630303136303031363034303162303331363831363030313630303136303430316230333136313136313334326135373530356639343933353035303530353035363562363030343833303135343630363038363031353136303031363030313630613031623033393039313136393036333466323234323966393036313334346639303631313161643536356236306130383530313531363038303839303135313630343038303862303135313566393038313532363030353861303136303230353238313930323036303031303135343930353136303031363030313630653031623033313936306530383739303162313638313532363133346131393439333932393138383931363030313630303136303430316230333930393131363930363030343031363135356134353635623630323036303430353138303833303338313836356166613135383031353631333462633537336435663830336533643566666435623530353035303530363034303531336436303166313936303166383230313136383230313830363034303532353038313031393036313334653039313930363134656535353635623935393435303530353035303530353635623631333466313631343935383536356235663631333466613631323032383536356235663834383135323630303738323031363032303532363034303830383232303831353136306530383130313930393235323830353439333934353039313932393039313930383239303630666631363630303538313131313536313335333335373631333533333631346261663536356236303035383131313135363133353434353736313335343436313462616635363562383135323630323030313630303138323031383035343631333535383930363134643631353635623830363031663031363032303830393130343032363032303031363034303531393038313031363034303532383039323931393038313831353236303230303138323830353436313335383439303631346436313536356238303135363133356366353738303630316631303631333561363537363130313030383038333534303430323833353239313630323030313931363133356366353635623832303139313930356635323630323035663230393035623831353438313532393036303031303139303630323030313830383331313631333562323537383239303033363031663136383230313931356235303530353039313833353235303530363030323832383130313534363030313630303136303430316230333830383231363630323038353031353236303031363034303162383230343831313636303430383530313532363030313630383031623832303438313136363036303835303135323630303136306330316239303931303438313136363038303834303135323630303339303933303135343930393231363630613039303931303135323930393135303831353136303035383131313135363133363364353736313336336436313462616635363562313436313336373035373566383438313532363030373833303136303230353236303430393038313930323035343930353136333137306363393333363065323162383135323631303463633931363066663136393036303034303136313464393935363562363030333831353234323630303136303031363034303162303331363630633038323031353235663834383135323630303738333031363032303532363034303930323038313531383135343833393239313930383239303630666631393136363030313833363030353831313131353631333662343537363133366234363134626166353635623032313739303535353036303230383230313531363030313832303139303631333663643930383236313530616235363562353036303430383230313531363030323832303138303534363036303835303135313630383038363031353136306130383730313531363030313630303136303430316230333935383631363630303136303031363038303162303331393930393431363933393039333137363030313630343031623932383631363932393039323032393139303931313736303031363030313630383031623033313636303031363038303162393138353136393139303931303236303031363030313630633031623033313631373630303136306330316239313834313639313930393130323137393035353630633039303932303135313630303339303931303138303534363030313630303136303430316230333139313639313930393231363137393035353566363133373661383538323631333139633536356236303830383430313531363034303830353136303031363030313630343031623033393039323136383235323432363032303833303135323931393335303833393235303837393137663133643538333934636632363964343862636639323739353961323961356666656537633939323464616666663839323765636466336334386666613763363739313031363034303531383039313033393061333530393339323530353035303536356237663665356264666363653135653533633334303665613637626663653337646364323666353135326435343932383234653433666435653363386163356162303039303536356236303430353136303031363030313630613031623033383338313136363032343833303135323630343438323031383339303532363130346435393138353931383231363930363361393035396362623930363036343031356236303430353136303230383138333033303338313532393036303430353239313530363065303162363032303832303138303531363030313630303136306530316230333833383138333136313738333532353035303530353036313431663235363562356636313064666438323631333835313631333763313536356235343630303136303031363061303162303331363930363134323461353635623566363133383661363132303238353635623630303930313534363066663136363133383865353736303430353136333766616238316535363065303162383135323630303430313630343035313830393130333930666435623566363133383937363132303238353635623930353034323631333861613630363038363031363034303837303136313462363035363562363030313630303136303430316230333136313131353830363133386534353735303631333863383632303261333030343236313531376135363562363133386438363036303836303136303430383730313631346236303536356236303031363030313630343031623033313631303135356231353631333931653537363133386639363036303835303136303430383630313631346236303536356236303430353136333538373964613133363065313162383135323630303136303031363034303162303339303931313636303034383230313532363032343031363130346363353635623631333933333631333932653630363038363031383636313536313035363562363134336162353635623631333934333631333932653630383038363031383636313536313035363562363033303631333935323630323038363031383636313536323435363562393035303134363133393834353736313339363636303230383530313835363135363234353635623630343035313633323634373562326636306531316238313532363130346363393235303630303430313930383135323630323030313930353635623631333938653834383036313536323435363562393035303566303336313339626235373631333961303834383036313536323435363562363034303531363333653038613132353630653131623831353236303034303136313034636339323931393036313536363635363562356636303038383230313631333963613836383036313536323435363562363034303531363133396438393239313930363135346132353635623930383135323630323030313630343035313830393130333930323035343134363133613131353736313339663638343830363135363234353635623630343035313633613431663737326636306530316238313532363030343031363130346363393239313930363135363636353635623631336131623833356636313366363435363562363034303830353136306530383130313930393135323831353438313532356639303831393037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f3930363365303437623238333930363032303831303136313361353838613830363135363234353635623830383036303166303136303230383039313034303236303230303136303430353139303831303136303430353238303933393239313930383138313532363032303031383338333830383238343337356639323031393139303931353235303530353039303832353235303630323039303831303139303631336161303930386230313862363135363234353635623830383036303166303136303230383039313034303236303230303136303430353139303831303136303430353238303933393239313930383138313532363032303031383338333830383238343337356639323031393139303931353235303530353039303832353235303630323030313631336165393630363038623031363034303863303136313462363035363562363030313630303136303430316230333136383135323630323030313631336230343630363038623031386236313536313035363562363133623064393036313536373935363562383135323630323030313631336231663630383038623031386236313536313035363562363133623238393036313536373935363562383135323630323030313838363030313630303136303430316230333136383135323530363034303531383236336666666666666666313636306530316238313532363030343031363133623536393139303631353739623536356235663630343035313830383330333831383635616634313538303135363133623730353733643566383033653364356666643562353035303530353036303430353133643566383233653630316633643930383130313630316631393136383230313630343035323631336239373931393038313031393036313538353235363562356638323831353236303036383630313630323035323630343039303230393139333530393135303631336262353832383236313530616235363562353038313630303838343031363133626335383838303631353632343536356236303430353136313362643339323931393036313534613235363562393038313532363034303531393038313930303336303230303138313230393139303931353536336565356234386562363065303162383135323566393036303035363030313630393931623031393036336565356234386562393036313363306639303835393036303034303136313465643335363562363032303630343035313830383330333831356638373561663131353830313536313363326235373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363133633466393139303631346565353536356236303430383035313630653038313031393039313532393039313530383036303031383135323630323030313631336336663839383036313536323435363562383038303630316630313630323038303931303430323630323030313630343035313930383130313630343035323830393339323931393038313831353236303230303138333833383038323834333735663932303138323930353235303933383535323530353035303630303136303031363034303162303338393136363032303830383430313832393035323630343038303835303138343930353236303630383530313932393039323532363038303834303138333930353236306130393039333031383239303532383638323532363030373838303139303932353232303831353138313534383239303630666631393136363030313833363030353831313131353631336366653537363133636665363134626166353635623032313739303535353036303230383230313531363030313832303139303631336431373930383236313530616235363562353036303430383230313531363030323832303138303534363036303835303135313630383038363031353136306130383730313531363030313630303136303430316230333935383631363630303136303031363038303162303331393930393431363933393039333137363030313630343031623932383631363932393039323032393139303931313736303031363030313630383031623033313636303031363038303162393138353136393139303931303236303031363030313630633031623033313631373630303136306330316239313834313639313930393130323137393035353630633039303932303135313630303339303931303138303534363030313630303136303430316230333139313639313930393231363137393035353830363133646234383838303631353632343536356236303430353136313364633239323931393036313534613235363562363034303531383039313033393032303834376662373732393765336265666336393162666338363461383165323431663833653265663732326236653762656361613265636563323530633664353262343330383938623630343030313630323038313031393036313365303039313930363134623630353635623630343038303531363030313630303136303430316230333933383431363831353239323930393131363630323038333031353230313630343035313830393130333930613435303930393539343530353035303530353035363562356636313365333736313230323835363562363030313031353436303031363030313630343031623033313639313930353035363562376666306335376531363834306466303430663135303838646332663831666533393163333932336265633733653233613936363265666339633232396336613030353436303031363034303162393030343630666631363631336539323537363034303531363331616663643739663630653331623831353236303034303136303430353138303931303339306664356235363562363133653963363133653439353635623631336561353831363134353134353635623631336561643631343532643536356236313065626236303630383230313335363038303833303133353631336563613630633038353031363061303836303136313462363035363562363133656461363065303836303136306330383730313631353839353536356236313365656236313031303038373031363065303838303136313538616535363562363130313030383730313335363133663033363130313430383930313631303132303861303136313465666335363562363134353364353635623631336631303631336534393536356235663631336631393631333763313536356239303530363030313630303136306130316230333832313636313366343435373831363034303531363337333330363830333630653031623831353236303034303136313034636339313930363134623962353635623830353436303031363030313630613031623033313931363630303136303031363061303162303339323930393231363931393039313137393035353536356235663631336636643631323032383536356239303530356638323630303136303031363034303162303331363834363030313630303136303430316230333136313131353631336639623537363133663934383338353631353464313536356239303530363133666138353635623631336661353834383436313534643135363562393035303562363034303830353136303830383130313832353236303032383430313534383038323532363030333835303135343630323038333031353236303034383530313534393238323031393239303932353236303035383430313534363030313630303136303430316230333136363036303832303135323432393131353830363134303061353735303630303138343031353438313531363134303036393136303031363030313630343031623033313639303631353137613536356238323130313535623135363134303330353736303031363030313630343031623033383331363630363038323031353238313831353236303430383130313531363032303832303135323631343034663536356238323831363036303031383138313531363134303432393139303631353462313536356236303031363030313630343031623033313639303532353035623630363038313031353136313430356639303630363436313535653535363562363032303832303135313630303138363031353436303031363030313630343031623033393239303932313639313631343038613931393036303031363034303162393030343630666631363631353161663536356231303135363134306261353736303630383130313531363034303531363364666165383830313630653031623831353236303031363030313630343031623033393039313136363030343832303135323630323430313631303463633536356238353630303136303031363034303162303331363831363034303031383138313531363134306435393139303631353137613536356239303532353036303430383130313830353136303031363030313630343031623033383731363931393036313430663539303833393036313535643235363562393035323530363030313834303135343630343038323031353136303634393136313431316139313630303136303430316239303931303436306666313639303631353161663536356231303135363134313431353738303630343030313531363034303531363335393433333137663630653031623831353236303034303136313034636339313831353236303230303139303536356238303531363030323835303135353630323038313031353136303033383530313535363034303831303135313630303438353031353536303630303135313630303539303933303138303534363030313630303136303430316230333139313636303031363030313630343031623033393039343136393339303933313739303932353535303530353035303536356235663830363134313933363132303238353635623566383438313532363030373832303136303230353236303430393032303630303230313830353439313932353039303630303839303631343163373930363030313630343031623930303436303031363030313630343031623033313636313538636535363562393139303631303130303061383135343831363030313630303136303430316230333032313931363930383336303031363030313630343031623033313630323137393035353931353035303931393035303536356235663631343230363630303136303031363061303162303338343136383336313436643335363562393035303830353135663134313538303135363134323261353735303830383036303230303139303531383130313930363134323238393139303631353865393536356231353562313536313034643535373832363034303531363335323734616665373630653031623831353236303034303136313034636339313930363134623962353635623566383038333630303136303031363061303162303331363633373061303832333133303630343035313832363366666666666666663136363065303162383135323630303430313631343237383931393036313462396235363562363032303630343035313830383330333831383635616661313538303135363134323933353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631343262373931393036313465653535363562393035303631343263653630303136303031363061303162303338353136333333303836363134366530353635623630343035313633373061303832333136306530316238313532356639303630303136303031363061303162303338363136393036333730613038323331393036313432666339303330393036303034303136313462396235363562363032303630343035313830383330333831383635616661313538303135363134333137353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631343333623931393036313465653535363562393035303831383131313631343361313537363034303531363234363162636436306535316238313532363032303630303438323031353236303263363032343832303135323766353336313636363534353532343333323330353437323631366537333636363537323436373236663664336132303632363136633631366536333635323036653630343438323031353236623162646430383161356239386463393935383563643935393630613231623630363438323031353236303834303136313034636335363562363133346530383238323631353564323536356236313433623836303230383230313832363134613933353635623633666666666666666631363135383031353631343364383537353036313433643336303230383230313832363134663137353635623135313539303530356231353631343431663537363134336561363032303832303138323631346139333536356236313433663736303230383330313833363134663137353635623630343035313633633038613066316436306530316238313532363366666666666666663930393331363630303438343031353236303234383330313532353036303434303136313034636335363562363134343263363032303832303138323631346631373536356239303530363134343362363032303833303138333631346139333536356236336666666666666666313631313135363134343534353736313433656136303230383230313832363134613933353635623630303135623631343436343630323038333031383336313466313735363562393035303831313031353631313466313537363134343761363032303833303138333631346631373536356236313434383536303031383436313535643235363562383138313130363134343934353736313434393436313466356335363562393035303630323030323031363032303831303139303631343461393931393036313465666335363562363030313630303136306130316230333136363134346266363032303834303138343631346631373536356238333831383131303631343463663537363134346366363134663563353635623930353036303230303230313630323038313031393036313434653439313930363134656663353635623630303136303031363061303162303331363130313536313435306335373630343035313633306462633864356636306533316238313532363030343031363034303531383039313033393066643562363030313031363134343537353635623631343531633631336534393536356236313435323436313437313935363562363130656262383136313437323135363562363134353335363133653439353635623631336539323631343830343536356236313435343536313365343935363562356636313435346536313234336135363562393035303631666666663835313631353830363134353636353735303631323731303631666666663836313631313562313536313435386135373630343035313633356631326536633336306531316238313532363166666666383631363630303438323031353236303234303136313034636335363562383638383131313536313435616535373630343035313633323232643136343336306532316238313532363030343831303138393930353236303234303136313034636335363562363066663834313631353830363134356331353735303630306136306666383531363131356231353631343565343537363034303531363331373064623335393630653331623831353236306666383531363630303438323031353236303234303136313034636335363562363134356563363133653265353635623630303136303031363034303162303331363836363030313630303136303430316230333136313031353631343632383537363034303531363230326130366436306531316238313532363030313630303136303430316230333837313636303034383230313532363032343031363130346363353635623832356630333631343634383537363034303531363361373333303037313630653031623831353236303034303136303430353138303931303339306664356239363837353536303031383730313935393039353535363030323836303138303534363030313630303136303430316230333935393039353136363966666666666666666666666666666666666666663139393039353136393439303934313736303031363034303162363166666666393439303934313639333930393330323932393039323137363766666666666666666666666666666666363035303162313931363630666639313930393131363630303136303530316230323137393039313535363030333833303135353630303439303931303138303534363030313630303136306130316230333139313636303031363030313630613031623033393039323136393139303931313739303535353635623630363036313230323138333833356636313438306335363562363034303531363030313630303136306130316230333834383131363630323438333031353238333831313636303434383330313532363036343832303138333930353236313037353039313836393138323136393036333233623837326464393036303834303136313338313235363562363133653932363133653439353635623631343732393631336534393536356235663631343733323631323032383536356238323335383135353930353036303134363134373461363036303834303136303430383530313631353861653536356236306666313631313830363134373639353735303631343736343630363038333031363034303834303136313538616535363562363066663136313535623135363134373964353736313437376536303630383330313630343038343031363135386165353635623630343035313633346135396262666636306531316238313532363066663930393131363630303438323031353236303234303136313034636335363562363134376164363036303833303136303430383430313631353861653536356236303031383230313830353436306666393239303932313636303031363034303162303236306666363034303162313939303932313639313930393131373930353536313437646536303430383330313630323038343031363134623630353635623630303139313930393130313830353436303031363030313630343031623033313931363630303136303031363034303162303339303932313639313930393131373930353535303536356236313261623036313365343935363562363036303831343731303135363134383331353733303630343035313633636437383630353936306530316238313532363030343031363130346363393139303631346239623536356235663830383536303031363030313630613031623033313638343836363034303531363134383463393139303631353035363536356235663630343035313830383330333831383538373561663139323530353035303364383035663831313436313438383635373630343035313931353036303166313936303366336430313136383230313630343035323364383235323364356636303230383430313365363134383862353635623630363039313530356235303931353039313530363134383962383638333833363134386135353635623936393535303530353035303530353035363562363036303832363134386261353736313438623538323631343866383536356236313230323135363562383135313135383031353631343864313537353036303031363030313630613031623033383431363362313535623135363134386631353738333630343035313633393939366233313536306530316238313532363030343031363130346363393139303631346239623536356235303830363132303231353635623830353131353631343930383537383035313830383236303230303166643562363034303531363330613132663532313630653131623831353236303034303136303430353138303931303339306664356235303830353436313439326439303631346436313536356235663832353538303630316631303631343933633537353035303536356236303166303136303230393030343930356635323630323035663230393038313031393036313065626239313930363134393935353635623630343038303531363065303831303139303931353238303566383135323630363036303230383230313831393035323566363034303833303138313930353239303832303138313930353236303830383230313831393035323630613038323031383139303532363063303930393130313532393035363562356238303832313131353631343961393537356638313535363030313031363134393936353635623530393035363562383031353135383131343631306562623537356638306664356238303335363366666666666666663831313638313134363134396364353735663830666435623931393035303536356235663830356636303630383438363033313231353631343965343537356638306664356238333335393235303630323038343031333536313439663638313631343961643536356239313530363134613034363034303835303136313439626135363562393035303932353039323530393235363562356636303230383238343033313231353631346131643537356638306664356235303335393139303530353635623566383036303430383338353033313231353631346133353537356638306664356238323335363030313630303136303430316230333831313131353631346134613537356638306664356238333031363038303831383630333132313536313461356235373566383066643562393135303631346136393630323038343031363134396261353635623930353039323530393239303530353635623566383036303430383338353033313231353631346138333537356638306664356238323335393135303631346136393630323038343031363134396261353635623566363032303832383430333132313536313461613335373566383066643562363132303231383236313439626135363562383033353631666666663831313638313134363134396364353735663830666435623630303136303031363034303162303338313136383131343631306562623537356638306664356235663830356638303630383038353837303331323135363134616534353735663830666435623834333536303031363030313630343031623033383131313135363134616639353735663830666435623835303136306130383138383033313231353631346230613537356638306664356239333530363134623138363032303836303136313461616335363562393235303630343038353031333536313462323838313631346162643536356239333936393239353530393239333630363030313335393235303530353635623566383036303430383338353033313231353631346234393537356638306664356236313462353238333631343962613536356239343630323039333930393330313335393335303530353035363562356636303230383238343033313231353631346237303537356638306664356238313335363132303231383136313461626435363562356638303630343038333835303331323135363134623863353735663830666435623530353038303335393236303230393039313031333539313530353635623630303136303031363061303162303339313930393131363831353236303230303139303536356236333465343837623731363065303162356635323630323136303034353236303234356666643562363030363831313036313462643335373631346264333631346261663536356239303532353635623566356238333831313031353631346266313537383138313031353138333832303135323630323030313631346264393536356235303530356639313031353235363562356638313531383038343532363134633130383136303230383630313630323038363031363134626437353635623630316630313630316631393136393239303932303136303230303139323931353035303536356236303230383135323631346333363630323038323031383335313631346263333536356235663630323038333031353136306530363034303834303135323631346335313631303130303834303138323631346266393536356239303530363034303834303135313630303136303031363034303162303338303832313636303630383630313532383036303630383730313531313636303830383630313532383036303830383730313531313636306130383630313532383036306130383730313531313636306330383630313532383036306330383730313531313636306530383630313532353035303830393135303530393239313530353035363562363030313630303136306130316230333831313638313134363130656262353735663830666435623566383038323834303336313031363038313132313536313463636235373566383066643562363130313430383038323132313536313463646135373566383066643562383439333530383330313335393035303631346365623831363134636134353635623830393135303530393235303932393035303536356235663830363032303833383530333132313536313464303735373566383066643562383233353630303136303031363034303162303338303832313131353631346431643537356638306664356238313835303139313530383536303166383330313132363134643330353735663830666435623831333538313831313131353631346433653537356638306664356238363630323038323835303130313131313536313464346635373566383066643562363032303932393039323031393639313935353039303933353035303530353035363562363030313831383131633930383231363830363134643735353736303766383231363931353035623630323038323130383130333631346439333537363334653438376237313630653031623566353236303232363030343532363032343566666435623530393139303530353635623630323038313031363130646664383238343631346263333536356236333465343837623731363065303162356635323630343136303034353236303234356666643562363034303531363036303831303136303031363030313630343031623033383131313832383231303137313536313464646435373631346464643631346461373536356236303430353239303536356236303430383035313930383130313630303136303031363034303162303338313131383238323130313731353631346464643537363134646464363134646137353635623630343035313630316638323031363031663139313638313031363030313630303136303430316230333831313138323832313031373135363134653264353736313465326436313464613735363562363034303532393139303530353635623566363030313630303136303430316230333832313131353631346534643537363134653464363134646137353635623530363031663031363031663139313636303230303139303536356235663832363031663833303131323631346536613537356638306664356238313531363134653764363134653738383236313465333535363562363134653035353635623831383135323834363032303833383630313031313131353631346539313537356638306664356236313065646438323630323038333031363032303837303136313462643735363562356636303230383238343033313231353631346562323537356638306664356238313531363030313630303136303430316230333831313131353631346563373537356638306664356236313065646438343832383530313631346535623536356236303230383135323566363132303231363032303833303138343631346266393536356235663630323038323834303331323135363134656635353735663830666435623530353139313930353035363562356636303230383238343033313231353631346630633537356638306664356238313335363132303231383136313463613435363562356638303833333536303165313938343336303330313831313236313466326335373566383066643562383330313830333539313530363030313630303136303430316230333832313131353631346634353537356638306664356236303230303139313530363030353831393031623336303338323133313536313333356635373566383066643562363334653438376237313630653031623566353236303332363030343532363032343566666435623566383233353630356531393833333630333031383131323631346638343537356638306664356239313930393130313932393135303530353635623566383236303166383330313132363134663964353735663830666435623831333536313466616236313465373838323631346533353536356238313831353238343630323038333836303130313131313536313466626635373566383066643562383136303230383530313630323038333031333735663931383130313630323030313931393039313532393339323530353035303536356235663630363038323336303331323135363134666562353735663830666435623631346666333631346462623536356238323335363030313630303136303430316230333830383231313135363135303039353735663830666435623631353031353336383338373031363134663865353635623833353236303230383530313335393135303830383231313135363135303261353735663830666435623530363135303337333638323836303136313466386535363562363032303833303135323530363034303833303133353631353034623831363134616264353635623630343038323031353239323931353035303536356235663832353136313466383438313834363032303837303136313462643735363562363031663832313131353631303464353537383035663532363032303566323036303166383430313630303531633831303136303230383531303135363135303863353735303830356236303166383430313630303531633832303139313530356238313831313031353631313935653537356638313535363030313031363135303938353635623831353136303031363030313630343031623033383131313135363135306334353736313530633436313464613735363562363135306438383136313530643238343534363134643631353635623834363135303637353635623630323038303630316638333131363030313831313436313531306235373566383431353631353066343537353038353833303135313562356631393630303338363930316231633139313636303031383539303162313738353535363131303039353635623566383538313532363032303831323036303166313938363136393135623832383131303135363135313339353738383836303135313832353539343834303139343630303139303931303139303834303136313531316135363562353038353832313031353631353135363537383738353031353135663139363030333838393031623630663831363163313931363831353535623530353035303530353036303031393038313162303139303535353035363562363334653438376237313630653031623566353236303131363030343532363032343566666435623830383230313830383231313135363130646664353736313064666436313531363635363562356636336666666666666666383038333136383138313033363135316135353736313531613536313531363635363562363030313031393339323530353035303536356238303832303238313135383238323034383431343137363130646664353736313064666436313531363635363562356638303833333536303165313938343336303330313831313236313531646235373566383066643562383330313630323038313031393235303335393035303630303136303031363034303162303338313131313536313531663935373566383066643562383033363033383231333135363133333566353735663830666435623831383335323831383136303230383530313337353035663832383230313630323039303831303139313930393135323630316639303931303136303166313931363930393130313031393035363562356638333833383535323630323038303836303139353530383038353630303531623833303130313834356635623837383131303135363135326534353738343833303336303166313930313839353238313335333638383930303336303565313930313831313236313532366235373566383066643562383730313630363036313532373938323830363135316336353635623832383735323631353238393833383830313832383436313532303735363562393235303530353036313532393938363833303138333631353163363536356238363833303338383838303135323631353261623833383238343631353230373536356239323530353035303630343038303833303133353932353036313532633038333631346162643536356236303031363030313630343031623033393239303932313639343930393130313933393039333532393738333031393739303833303139303630303130313631353234363536356235303930393739363530353035303530353035303530353635623630323038313532383133353630323038323031353236303230383230313335363034303832303135323566363034303833303133353631353331353831363134636134353635623630303136303031363061303162303331363630363038333831303139313930393135323833303133353336383439303033363031653139303138313132363135333363353735663830666435623833303136303230383130313930333536303031363030313630343031623033383131313135363135333537353735663830666435623830363030353162333630333832313331353631353336383537356638306664356236303830383038353031353236313334653036306130383530313832383436313532326635363562356638323631353339363537363334653438376237313630653031623566353236303132363030343532363032343566666435623530303439303536356236303230383130313630303438333130363135336166353736313533616636313462616635363562393139303532393035363562356638303566363036303834383630333132313536313533633735373566383066643562383335313932353036303230383430313531363135336439383136313461626435363562363034303835303135313930393235303631353365613831363134616264353635623830393135303530393235303932353039323536356235663830363034303833383530333132313536313534303635373566383066643562383235313931353036303230383330313531363134636562383136313439616435363562356636303230383038333532356638343534363135343261383136313464363135363562383036303230383730313532363034303630303138303834313635663831313436313534346235373630303138313134363135343637353736313534393435363562363066663139383531363630343038613031353236303430383431353135363030353162386130313031393535303631353439343536356238393566353236303230356632303566356238353831313031353631353438623537383135343862383230313836303135323930383330313930383830313631353437303536356238613031363034303031393635303530356235303933393839373530353035303530353035303530353035363562383138333832333735663931303139303831353239313930353035363562363030313630303136303430316230333831383131363833383231363031393038303832313131353631323135303537363132313530363135313636353635623630303136303031363034303162303338323831313638323832313630333930383038323131313536313231353035373631323135303631353136363536356235663830363034303833383530333132313536313535303235373566383066643562383235313630303136303031363034303162303338303832313131353631353531383537356638306664356239303834303139303630363038323837303331323135363135353262353735663830666435623631353533333631346462623536356238323531383135323630323038333031353136313535343538313631346361343536356236303230383230313532363034303833303135313832383131313135363135353562353735663830666435623631353536373838383238363031363134653562353635623630343038333031353235303830393435303530353035303630323038333031353136313463656238313631343961643536356235663830363034303833383530333132313536313535393235373566383066643562383235313931353036303230383330313531363134636562383136313461626435363562393438353532363030313630303136303430316230333933383431363630323038363031353239313833313636303430383530313532383231363630363038343031353231363630383038323031353236306130303139303536356238313831303338313831313131353631306466643537363130646664363135313636353635623630303136303031363034303162303338313831313638333832313630323830383231363931393038323831313436313536303835373631353630383631353136363536356235303530393239313530353035363562356638323335363033653139383333363033303138313132363134663834353735663830666435623566383038333335363031653139383433363033303138313132363135363339353735663830666435623833303138303335393135303630303136303031363034303162303338323131313536313536353235373566383066643562363032303031393135303336383139303033383231333135363133333566353735663830666435623630323038313532356636313065646436303230383330313834383636313532303735363562356636303430383233363033313231353631353638393537356638306664356236313536393136313464653335363562363135363961383336313439626135363562383135323630323038303834303133353630303136303031363034303162303338303832313131353631353662363537356638306664356239303835303139303336363031663833303131323631353663383537356638306664356238313335383138313131313536313536646135373631353664613631346461373536356238303630303531623931353036313536656238343833303136313465303535363562383138313532393138333031383430313931383438313031393033363834313131353631353730343537356638306664356239333835303139333562383338353130313536313537326535373834333539323530363135373165383336313463613435363562383238323532393338353031393339303835303139303631353730393536356239343836303139343930393435323530393239353934353035303530353035303536356235663630343038333031363366666666666666663833353131363834353236303230383038343031353136303430363032303837303135323832383135313830383535323630363038383031393135303630323038333031393435303566393235303562383038333130313536313234326635373834353136303031363030313630613031623033313638323532393338333031393336303031393239303932303139313930383330313930363135373732353635623630323038313532383135313630323038323031353235663630323038333031353136306530363034303834303135323631353763313631303130303834303138323631346266393536356239303530363034303834303135313630316631393830383538343033303136303630383630313532363135376466383338333631346266393536356239323530363030313630303136303430316230333630363038373031353131363630383038363031353236303830383630313531393135303830383538343033303136306130383630313532363135383066383338333631353734303536356239323530363061303836303135313931353038303835383430333031363063303836303135323530363135383264383238323631353734303536356239313530353036306330383430313531363135383461363065303835303138323630303136303031363034303162303331363930353235363562353039333932353035303530353635623566383036303430383338353033313231353631353836333537356638306664356238323531393135303630323038333031353136303031363030313630343031623033383131313135363135383766353735663830666435623631353838623835383238363031363134653562353635623931353035303932353039323930353035363562356636303230383238343033313231353631353861353537356638306664356236313230323138323631346161633536356235663630323038323834303331323135363135386265353735663830666435623831333536306666383131363831313436313230323135373566383066643562356636303031363030313630343031623033383038333136383138313033363135316135353736313531613536313531363635363562356636303230383238343033313231353631353866393537356638306664356238313531363132303231383136313439616435366665613136343733366636633633343330303038313930303061",
}

// ERC20TokenStakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenStakingManagerMetaData.ABI instead.
var ERC20TokenStakingManagerABI = ERC20TokenStakingManagerMetaData.ABI

// ERC20TokenStakingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenStakingManagerMetaData.Bin instead.
var ERC20TokenStakingManagerBin = ERC20TokenStakingManagerMetaData.Bin

// DeployERC20TokenStakingManager deploys a new Ethereum contract, binding an instance of ERC20TokenStakingManager to it.
func DeployERC20TokenStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *ERC20TokenStakingManager, error) {
	parsed, err := ERC20TokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenStakingManagerBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20TokenStakingManager{ERC20TokenStakingManagerCaller: ERC20TokenStakingManagerCaller{contract: contract}, ERC20TokenStakingManagerTransactor: ERC20TokenStakingManagerTransactor{contract: contract}, ERC20TokenStakingManagerFilterer: ERC20TokenStakingManagerFilterer{contract: contract}}, nil
}

// ERC20TokenStakingManager is an auto generated Go binding around an Ethereum contract.
type ERC20TokenStakingManager struct {
	ERC20TokenStakingManagerCaller     // Read-only binding to the contract
	ERC20TokenStakingManagerTransactor // Write-only binding to the contract
	ERC20TokenStakingManagerFilterer   // Log filterer for contract events
}

// ERC20TokenStakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenStakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenStakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenStakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenStakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenStakingManagerSession struct {
	Contract     *ERC20TokenStakingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20TokenStakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenStakingManagerCallerSession struct {
	Contract *ERC20TokenStakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ERC20TokenStakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenStakingManagerTransactorSession struct {
	Contract     *ERC20TokenStakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ERC20TokenStakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenStakingManagerRaw struct {
	Contract *ERC20TokenStakingManager // Generic contract binding to access the raw methods on
}

// ERC20TokenStakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerCallerRaw struct {
	Contract *ERC20TokenStakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenStakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenStakingManagerTransactorRaw struct {
	Contract *ERC20TokenStakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenStakingManager creates a new instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManager(address common.Address, backend bind.ContractBackend) (*ERC20TokenStakingManager, error) {
	contract, err := bindERC20TokenStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManager{ERC20TokenStakingManagerCaller: ERC20TokenStakingManagerCaller{contract: contract}, ERC20TokenStakingManagerTransactor: ERC20TokenStakingManagerTransactor{contract: contract}, ERC20TokenStakingManagerFilterer: ERC20TokenStakingManagerFilterer{contract: contract}}, nil
}

// NewERC20TokenStakingManagerCaller creates a new read-only instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenStakingManagerCaller, error) {
	contract, err := bindERC20TokenStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerCaller{contract: contract}, nil
}

// NewERC20TokenStakingManagerTransactor creates a new write-only instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenStakingManagerTransactor, error) {
	contract, err := bindERC20TokenStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerTransactor{contract: contract}, nil
}

// NewERC20TokenStakingManagerFilterer creates a new log filterer instance of ERC20TokenStakingManager, bound to a specific deployed contract.
func NewERC20TokenStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenStakingManagerFilterer, error) {
	contract, err := bindERC20TokenStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerFilterer{contract: contract}, nil
}

// bindERC20TokenStakingManager binds a generic wrapper to an already deployed contract.
func bindERC20TokenStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TokenStakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenStakingManager.Contract.ERC20TokenStakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ERC20TokenStakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ERC20TokenStakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenStakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) ADDRESSLENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "ADDRESS_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ADDRESSLENGTH() (uint32, error) {
	return _ERC20TokenStakingManager.Contract.ADDRESSLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// ADDRESSLENGTH is a free data retrieval call binding the contract method 0x60305d62.
//
// Solidity: function ADDRESS_LENGTH() view returns(uint32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) ADDRESSLENGTH() (uint32, error) {
	return _ERC20TokenStakingManager.Contract.ADDRESSLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) BIPSCONVERSIONFACTOR(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "BIPS_CONVERSION_FACTOR")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_ERC20TokenStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.BIPSCONVERSIONFACTOR(&_ERC20TokenStakingManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) BLSPUBLICKEYLENGTH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "BLS_PUBLIC_KEY_LENGTH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.BLSPUBLICKEYLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// BLSPUBLICKEYLENGTH is a free data retrieval call binding the contract method 0x8280a25a.
//
// Solidity: function BLS_PUBLIC_KEY_LENGTH() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) BLSPUBLICKEYLENGTH() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.BLSPUBLICKEYLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// ERC20STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xe4a63c40.
//
// Solidity: function ERC20_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) ERC20STAKINGMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "ERC20_STAKING_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC20STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xe4a63c40.
//
// Solidity: function ERC20_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ERC20STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.ERC20STAKINGMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// ERC20STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xe4a63c40.
//
// Solidity: function ERC20_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) ERC20STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.ERC20STAKINGMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMCHURNPERCENTAGELIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_CHURN_PERCENTAGE_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMCHURNPERCENTAGELIMIT is a free data retrieval call binding the contract method 0xc974d1b6.
//
// Solidity: function MAXIMUM_CHURN_PERCENTAGE_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMCHURNPERCENTAGELIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMCHURNPERCENTAGELIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMDELEGATIONFEEBIPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_DELEGATION_FEE_BIPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMREGISTRATIONEXPIRYLENGTH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_REGISTRATION_EXPIRY_LENGTH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMREGISTRATIONEXPIRYLENGTH is a free data retrieval call binding the contract method 0xdf93d8de.
//
// Solidity: function MAXIMUM_REGISTRATION_EXPIRY_LENGTH() view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMREGISTRATIONEXPIRYLENGTH() (uint64, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMREGISTRATIONEXPIRYLENGTH(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) MAXIMUMSTAKEMULTIPLIERLIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "MAXIMUM_STAKE_MULTIPLIER_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_ERC20TokenStakingManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) POSVALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "POS_VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// POSVALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xafb98096.
//
// Solidity: function POS_VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) POSVALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.POSVALIDATORMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) PCHAINBLOCKCHAINID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "P_CHAIN_BLOCKCHAIN_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.PCHAINBLOCKCHAINID(&_ERC20TokenStakingManager.CallOpts)
}

// PCHAINBLOCKCHAINID is a free data retrieval call binding the contract method 0x732214f8.
//
// Solidity: function P_CHAIN_BLOCKCHAIN_ID() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) PCHAINBLOCKCHAINID() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.PCHAINBLOCKCHAINID(&_ERC20TokenStakingManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) VALIDATORMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "VALIDATOR_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// VALIDATORMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xbc5fbfec.
//
// Solidity: function VALIDATOR_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) VALIDATORMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.VALIDATORMANAGERSTORAGELOCATION(&_ERC20TokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) WARPMESSENGER() (common.Address, error) {
	return _ERC20TokenStakingManager.Contract.WARPMESSENGER(&_ERC20TokenStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _ERC20TokenStakingManager.Contract.WARPMESSENGER(&_ERC20TokenStakingManager.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) GetValidator(opts *bind.CallOpts, validationID [32]byte) (Validator, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "getValidator", validationID)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _ERC20TokenStakingManager.Contract.GetValidator(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _ERC20TokenStakingManager.Contract.GetValidator(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) GetWeight(opts *bind.CallOpts, validationID [32]byte) (uint64, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "getWeight", validationID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.GetWeight(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// GetWeight is a free data retrieval call binding the contract method 0x66435abf.
//
// Solidity: function getWeight(bytes32 validationID) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) GetWeight(validationID [32]byte) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.GetWeight(&_ERC20TokenStakingManager.CallOpts, validationID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) RegisteredValidators(opts *bind.CallOpts, nodeID []byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "registeredValidators", nodeID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.RegisteredValidators(&_ERC20TokenStakingManager.CallOpts, nodeID)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xfd7ac5e7.
//
// Solidity: function registeredValidators(bytes nodeID) view returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) RegisteredValidators(nodeID []byte) ([32]byte, error) {
	return _ERC20TokenStakingManager.Contract.RegisteredValidators(&_ERC20TokenStakingManager.CallOpts, nodeID)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) ValueToWeight(opts *bind.CallOpts, value *big.Int) (uint64, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "valueToWeight", value)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.ValueToWeight(&_ERC20TokenStakingManager.CallOpts, value)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.ValueToWeight(&_ERC20TokenStakingManager.CallOpts, value)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) WeightToValue(opts *bind.CallOpts, weight uint64) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "weightToValue", weight)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _ERC20TokenStakingManager.Contract.WeightToValue(&_ERC20TokenStakingManager.CallOpts, weight)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _ERC20TokenStakingManager.Contract.WeightToValue(&_ERC20TokenStakingManager.CallOpts, weight)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ClaimDelegationFees(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ClaimDelegationFees(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x5297fae6.
//
// Solidity: function completeDelegatorRegistration(uint32 messageIndex, bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, messageIndex uint32, delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeDelegatorRegistration", messageIndex, delegationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x5297fae6.
//
// Solidity: function completeDelegatorRegistration(uint32 messageIndex, bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteDelegatorRegistration(messageIndex uint32, delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, messageIndex, delegationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x5297fae6.
//
// Solidity: function completeDelegatorRegistration(uint32 messageIndex, bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteDelegatorRegistration(messageIndex uint32, delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, messageIndex, delegationID)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x98f3e2b4.
//
// Solidity: function completeEndDelegation(uint32 messageIndex, bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteEndDelegation(opts *bind.TransactOpts, messageIndex uint32, delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeEndDelegation", messageIndex, delegationID)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x98f3e2b4.
//
// Solidity: function completeEndDelegation(uint32 messageIndex, bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteEndDelegation(messageIndex uint32, delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndDelegation(&_ERC20TokenStakingManager.TransactOpts, messageIndex, delegationID)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x98f3e2b4.
//
// Solidity: function completeEndDelegation(uint32 messageIndex, bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteEndDelegation(messageIndex uint32, delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndDelegation(&_ERC20TokenStakingManager.TransactOpts, messageIndex, delegationID)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteEndValidation(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeEndValidation", messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndValidation(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndValidation(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ForceInitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "forceInitializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndValidation(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ForceInitializeEndValidation(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0xf74c607b.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) Initialize(opts *bind.TransactOpts, settings PoSValidatorManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initialize", settings, token)
}

// Initialize is a paid mutator transaction binding the contract method 0xf74c607b.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) Initialize(settings PoSValidatorManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.Initialize(&_ERC20TokenStakingManager.TransactOpts, settings, token)
}

// Initialize is a paid mutator transaction binding the contract method 0xf74c607b.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,uint256,address) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) Initialize(settings PoSValidatorManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.Initialize(&_ERC20TokenStakingManager.TransactOpts, settings, token)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0x9e1bc4ef.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID, uint256 delegationAmount) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte, delegationAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeDelegatorRegistration", validationID, delegationAmount)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0x9e1bc4ef.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID, uint256 delegationAmount) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, validationID, delegationAmount)
}

// InitializeDelegatorRegistration is a paid mutator transaction binding the contract method 0x9e1bc4ef.
//
// Solidity: function initializeDelegatorRegistration(bytes32 validationID, uint256 delegationAmount) returns(bytes32)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, validationID, delegationAmount)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndValidation(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeEndValidation(&_ERC20TokenStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x4bee0040.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount) returns(bytes32 validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeValidatorRegistration(opts *bind.TransactOpts, registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeValidatorRegistration", registrationInput, delegationFeeBips, minStakeDuration, stakeAmount)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x4bee0040.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount) returns(bytes32 validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration, stakeAmount)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x4bee0040.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) registrationInput, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount) returns(bytes32 validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeValidatorRegistration(registrationInput ValidatorRegistrationInput, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeValidatorRegistration(&_ERC20TokenStakingManager.TransactOpts, registrationInput, delegationFeeBips, minStakeDuration, stakeAmount)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) InitializeValidatorSet(opts *bind.TransactOpts, subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initializeValidatorSet", subnetConversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeValidatorSet(&_ERC20TokenStakingManager.TransactOpts, subnetConversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) InitializeValidatorSet(subnetConversionData SubnetConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.InitializeValidatorSet(&_ERC20TokenStakingManager.TransactOpts, subnetConversionData, messageIndex)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ResendEndValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "resendEndValidatorMessage", validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendEndValidatorMessage(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendEndValidatorMessage(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ResendRegisterValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "resendRegisterValidatorMessage", validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendRegisterValidatorMessage(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendRegisterValidatorMessage(&_ERC20TokenStakingManager.TransactOpts, validationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ResendUpdateDelegation(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "resendUpdateDelegation", delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendUpdateDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ResendUpdateDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.SubmitUptimeProof(&_ERC20TokenStakingManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.SubmitUptimeProof(&_ERC20TokenStakingManager.TransactOpts, validationID, messageIndex)
}

// ERC20TokenStakingManagerDelegationEndedIterator is returned from FilterDelegationEnded and is used to iterate over the raw logs and unpacked data for DelegationEnded events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegationEndedIterator struct {
	Event *ERC20TokenStakingManagerDelegationEnded // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerDelegationEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegationEnded)
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
		it.Event = new(ERC20TokenStakingManagerDelegationEnded)
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
func (it *ERC20TokenStakingManagerDelegationEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegationEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegationEnded represents a DelegationEnded event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegationEnded struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegationEnded is a free log retrieval operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegationEnded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerDelegationEndedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegationEndedIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegationEnded", logs: logs, sub: sub}, nil
}

// WatchDelegationEnded is a free log subscription operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegationEnded(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegationEnded, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegationEnded)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegationEnded(log types.Log) (*ERC20TokenStakingManagerDelegationEnded, error) {
	event := new(ERC20TokenStakingManagerDelegationEnded)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerDelegatorAddedIterator is returned from FilterDelegatorAdded and is used to iterate over the raw logs and unpacked data for DelegatorAdded events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorAddedIterator struct {
	Event *ERC20TokenStakingManagerDelegatorAdded // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerDelegatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegatorAdded)
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
		it.Event = new(ERC20TokenStakingManagerDelegatorAdded)
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
func (it *ERC20TokenStakingManagerDelegatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegatorAdded represents a DelegatorAdded event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorAdded struct {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegatorAdded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*ERC20TokenStakingManagerDelegatorAddedIterator, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegatorAddedIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegatorAdded", logs: logs, sub: sub}, nil
}

// WatchDelegatorAdded is a free log subscription operation binding the contract event 0xb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegatorAdded(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegatorAdded, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegatorAdded)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegatorAdded(log types.Log) (*ERC20TokenStakingManagerDelegatorAdded, error) {
	event := new(ERC20TokenStakingManagerDelegatorAdded)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerDelegatorRegisteredIterator is returned from FilterDelegatorRegistered and is used to iterate over the raw logs and unpacked data for DelegatorRegistered events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRegisteredIterator struct {
	Event *ERC20TokenStakingManagerDelegatorRegistered // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerDelegatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegatorRegistered)
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
		it.Event = new(ERC20TokenStakingManagerDelegatorRegistered)
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
func (it *ERC20TokenStakingManagerDelegatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegatorRegistered represents a DelegatorRegistered event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRegistered struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRegistered is a free log retrieval operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegatorRegistered(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerDelegatorRegisteredIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegatorRegisteredIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegatorRegistered", logs: logs, sub: sub}, nil
}

// WatchDelegatorRegistered is a free log subscription operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegatorRegistered(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegatorRegistered, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegatorRegistered)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegatorRegistered(log types.Log) (*ERC20TokenStakingManagerDelegatorRegistered, error) {
	event := new(ERC20TokenStakingManagerDelegatorRegistered)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerDelegatorRemovalInitializedIterator is returned from FilterDelegatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for DelegatorRemovalInitialized events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRemovalInitializedIterator struct {
	Event *ERC20TokenStakingManagerDelegatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerDelegatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerDelegatorRemovalInitialized)
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
		it.Event = new(ERC20TokenStakingManagerDelegatorRemovalInitialized)
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
func (it *ERC20TokenStakingManagerDelegatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerDelegatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerDelegatorRemovalInitialized represents a DelegatorRemovalInitialized event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerDelegatorRemovalInitialized struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRemovalInitialized is a free log retrieval operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterDelegatorRemovalInitialized(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20TokenStakingManagerDelegatorRemovalInitializedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerDelegatorRemovalInitializedIterator{contract: _ERC20TokenStakingManager.contract, event: "DelegatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchDelegatorRemovalInitialized is a free log subscription operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchDelegatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerDelegatorRemovalInitialized, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerDelegatorRemovalInitialized)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseDelegatorRemovalInitialized(log types.Log) (*ERC20TokenStakingManagerDelegatorRemovalInitialized, error) {
	event := new(ERC20TokenStakingManagerDelegatorRemovalInitialized)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitialValidatorCreatedIterator is returned from FilterInitialValidatorCreated and is used to iterate over the raw logs and unpacked data for InitialValidatorCreated events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitialValidatorCreatedIterator struct {
	Event *ERC20TokenStakingManagerInitialValidatorCreated // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerInitialValidatorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitialValidatorCreated)
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
		it.Event = new(ERC20TokenStakingManagerInitialValidatorCreated)
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
func (it *ERC20TokenStakingManagerInitialValidatorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitialValidatorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitialValidatorCreated represents a InitialValidatorCreated event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitialValidatorCreated struct {
	ValidationID [32]byte
	NodeID       common.Hash
	Weight       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialValidatorCreated is a free log retrieval operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitialValidatorCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte) (*ERC20TokenStakingManagerInitialValidatorCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitialValidatorCreatedIterator{contract: _ERC20TokenStakingManager.contract, event: "InitialValidatorCreated", logs: logs, sub: sub}, nil
}

// WatchInitialValidatorCreated is a free log subscription operation binding the contract event 0x9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf.
//
// Solidity: event InitialValidatorCreated(bytes32 indexed validationID, bytes indexed nodeID, uint256 weight)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitialValidatorCreated(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitialValidatorCreated, validationID [][32]byte, nodeID [][]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "InitialValidatorCreated", validationIDRule, nodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitialValidatorCreated)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitialValidatorCreated(log types.Log) (*ERC20TokenStakingManagerInitialValidatorCreated, error) {
	event := new(ERC20TokenStakingManagerInitialValidatorCreated)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "InitialValidatorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitializedIterator struct {
	Event *ERC20TokenStakingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerInitialized)
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
		it.Event = new(ERC20TokenStakingManagerInitialized)
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
func (it *ERC20TokenStakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerInitialized represents a Initialized event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20TokenStakingManagerInitializedIterator, error) {

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerInitializedIterator{contract: _ERC20TokenStakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerInitialized)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseInitialized(log types.Log) (*ERC20TokenStakingManagerInitialized, error) {
	event := new(ERC20TokenStakingManagerInitialized)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerUptimeUpdatedIterator struct {
	Event *ERC20TokenStakingManagerUptimeUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerUptimeUpdated)
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
		it.Event = new(ERC20TokenStakingManagerUptimeUpdated)
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
func (it *ERC20TokenStakingManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerUptimeUpdated represents a UptimeUpdated event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20TokenStakingManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerUptimeUpdatedIterator{contract: _ERC20TokenStakingManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerUptimeUpdated)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseUptimeUpdated(log types.Log) (*ERC20TokenStakingManagerUptimeUpdated, error) {
	event := new(ERC20TokenStakingManagerUptimeUpdated)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidationPeriodCreatedIterator is returned from FilterValidationPeriodCreated and is used to iterate over the raw logs and unpacked data for ValidationPeriodCreated events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodCreatedIterator struct {
	Event *ERC20TokenStakingManagerValidationPeriodCreated // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerValidationPeriodCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidationPeriodCreated)
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
		it.Event = new(ERC20TokenStakingManagerValidationPeriodCreated)
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
func (it *ERC20TokenStakingManagerValidationPeriodCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidationPeriodCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidationPeriodCreated represents a ValidationPeriodCreated event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodCreated struct {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidationPeriodCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (*ERC20TokenStakingManagerValidationPeriodCreatedIterator, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidationPeriodCreatedIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidationPeriodCreated", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodCreated is a free log subscription operation binding the contract event 0xb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidationPeriodCreated(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidationPeriodCreated, validationID [][32]byte, nodeID [][]byte, registerValidationMessageID [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidationPeriodCreated)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidationPeriodCreated(log types.Log) (*ERC20TokenStakingManagerValidationPeriodCreated, error) {
	event := new(ERC20TokenStakingManagerValidationPeriodCreated)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidationPeriodEndedIterator is returned from FilterValidationPeriodEnded and is used to iterate over the raw logs and unpacked data for ValidationPeriodEnded events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodEndedIterator struct {
	Event *ERC20TokenStakingManagerValidationPeriodEnded // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerValidationPeriodEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidationPeriodEnded)
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
		it.Event = new(ERC20TokenStakingManagerValidationPeriodEnded)
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
func (it *ERC20TokenStakingManagerValidationPeriodEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidationPeriodEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidationPeriodEnded represents a ValidationPeriodEnded event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodEnded struct {
	ValidationID [32]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodEnded is a free log retrieval operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidationPeriodEnded(opts *bind.FilterOpts, validationID [][32]byte, status []uint8) (*ERC20TokenStakingManagerValidationPeriodEndedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidationPeriodEndedIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidationPeriodEnded", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodEnded is a free log subscription operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidationPeriodEnded(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidationPeriodEnded, validationID [][32]byte, status []uint8) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidationPeriodEnded)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidationPeriodEnded(log types.Log) (*ERC20TokenStakingManagerValidationPeriodEnded, error) {
	event := new(ERC20TokenStakingManagerValidationPeriodEnded)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidationPeriodRegisteredIterator is returned from FilterValidationPeriodRegistered and is used to iterate over the raw logs and unpacked data for ValidationPeriodRegistered events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodRegisteredIterator struct {
	Event *ERC20TokenStakingManagerValidationPeriodRegistered // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerValidationPeriodRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidationPeriodRegistered)
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
		it.Event = new(ERC20TokenStakingManagerValidationPeriodRegistered)
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
func (it *ERC20TokenStakingManagerValidationPeriodRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidationPeriodRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidationPeriodRegistered represents a ValidationPeriodRegistered event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationPeriodRegistered struct {
	ValidationID [32]byte
	Weight       *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodRegistered is a free log retrieval operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidationPeriodRegistered(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20TokenStakingManagerValidationPeriodRegisteredIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidationPeriodRegisteredIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidationPeriodRegistered", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodRegistered is a free log subscription operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidationPeriodRegistered(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidationPeriodRegistered, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidationPeriodRegistered)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidationPeriodRegistered(log types.Log) (*ERC20TokenStakingManagerValidationPeriodRegistered, error) {
	event := new(ERC20TokenStakingManagerValidationPeriodRegistered)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidatorRemovalInitializedIterator is returned from FilterValidatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for ValidatorRemovalInitialized events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorRemovalInitializedIterator struct {
	Event *ERC20TokenStakingManagerValidatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerValidatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidatorRemovalInitialized)
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
		it.Event = new(ERC20TokenStakingManagerValidatorRemovalInitialized)
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
func (it *ERC20TokenStakingManagerValidatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidatorRemovalInitialized represents a ValidatorRemovalInitialized event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorRemovalInitialized struct {
	ValidationID       [32]byte
	SetWeightMessageID [32]byte
	Weight             *big.Int
	EndTime            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemovalInitialized is a free log retrieval operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidatorRemovalInitialized(opts *bind.FilterOpts, validationID [][32]byte, setWeightMessageID [][32]byte) (*ERC20TokenStakingManagerValidatorRemovalInitializedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidatorRemovalInitializedIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchValidatorRemovalInitialized is a free log subscription operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidatorRemovalInitialized, validationID [][32]byte, setWeightMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidatorRemovalInitialized)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidatorRemovalInitialized(log types.Log) (*ERC20TokenStakingManagerValidatorRemovalInitialized, error) {
	event := new(ERC20TokenStakingManagerValidatorRemovalInitialized)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenStakingManagerValidatorWeightUpdateIterator is returned from FilterValidatorWeightUpdate and is used to iterate over the raw logs and unpacked data for ValidatorWeightUpdate events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorWeightUpdateIterator struct {
	Event *ERC20TokenStakingManagerValidatorWeightUpdate // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerValidatorWeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidatorWeightUpdate)
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
		it.Event = new(ERC20TokenStakingManagerValidatorWeightUpdate)
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
func (it *ERC20TokenStakingManagerValidatorWeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidatorWeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidatorWeightUpdate represents a ValidatorWeightUpdate event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidatorWeightUpdate struct {
	ValidationID       [32]byte
	Nonce              uint64
	ValidatorWeight    uint64
	SetWeightMessageID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorWeightUpdate is a free log retrieval operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidatorWeightUpdate(opts *bind.FilterOpts, validationID [][32]byte, nonce []uint64) (*ERC20TokenStakingManagerValidatorWeightUpdateIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidatorWeightUpdateIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidatorWeightUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorWeightUpdate is a free log subscription operation binding the contract event 0x07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df.
//
// Solidity: event ValidatorWeightUpdate(bytes32 indexed validationID, uint64 indexed nonce, uint64 validatorWeight, bytes32 setWeightMessageID)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidatorWeightUpdate(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidatorWeightUpdate, validationID [][32]byte, nonce []uint64) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidatorWeightUpdate", validationIDRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidatorWeightUpdate)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidatorWeightUpdate(log types.Log) (*ERC20TokenStakingManagerValidatorWeightUpdate, error) {
	event := new(ERC20TokenStakingManagerValidatorWeightUpdate)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidatorWeightUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
