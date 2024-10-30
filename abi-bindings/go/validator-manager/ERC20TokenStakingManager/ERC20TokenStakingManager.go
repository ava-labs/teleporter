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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADDRESS_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BIPS_CONVERSION_FACTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERC20_STAKING_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POS_VALIDATOR_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPTIME_REWARDS_THRESHOLD_PERCENTAGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimDelegationFees\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimValidationRewards\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeEndValidation\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitializeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structValidator\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"startingWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"endedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWeight\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structPoSValidatorManagerSettings\",\"components\":[{\"name\":\"baseSettings\",\"type\":\"tuple\",\"internalType\":\"structValidatorManagerSettings\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"churnPeriodSeconds\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"minimumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"weightToValueFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardCalculator\",\"type\":\"address\",\"internalType\":\"contractIRewardCalculator\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20Mintable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeDelegatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delegationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeValidatorRegistration\",\"inputs\":[{\"name\":\"registrationInput\",\"type\":\"tuple\",\"internalType\":\"structValidatorRegistrationInput\",\"components\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"disableOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeValidatorSet\",\"inputs\":[{\"name\":\"subnetConversionData\",\"type\":\"tuple\",\"internalType\":\"structSubnetConversionData\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialValidators\",\"type\":\"tuple[]\",\"internalType\":\"structInitialValidator[]\",\"components\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredValidators\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resendEndValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendRegisterValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendUpdateDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitUptimeProof\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"valueToWeight\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"weightToValue\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DelegationEnded\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewards\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorAdded\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"delegatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"delegatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRegistered\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRemovalInitialized\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitialValidatorCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UptimeUpdated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"uptime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodEnded\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumValidatorStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodRegistered\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationRewardsClaimed\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reward\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRemovalInitialized\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorWeightUpdate\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DelegatorIneligibleForRewards\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSKeyLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidBLSPublicKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCodecID\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationFee\",\"inputs\":[{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationID\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDelegatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitializationStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMaximumChurnPercentage\",\"inputs\":[{\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMinStakeDuration\",\"inputs\":[{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeID\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPChainOwnerThreshold\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addressesLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRegistrationExpiry\",\"inputs\":[{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeAmount\",\"inputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeMultiplier\",\"inputs\":[{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidSubnetConversionID\",\"inputs\":[{\"name\":\"encodedSubnetConversionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedSubnetConversionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenAddress\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTotalWeight\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidValidationID\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorManagerAddress\",\"inputs\":[{\"name\":\"validatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorManagerBlockchainID\",\"inputs\":[{\"name\":\"blockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWarpOriginSenderAddress\",\"inputs\":[{\"name\":\"senderAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpSourceChainID\",\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MaxChurnRateExceeded\",\"inputs\":[{\"name\":\"churnAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MaxWeightExceeded\",\"inputs\":[{\"name\":\"newValidatorWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MinStakeDurationNotPassed\",\"inputs\":[{\"name\":\"endTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NodeAlreadyRegistered\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PChainOwnerAddressesNotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnexpectedRegistrationStatus\",\"inputs\":[{\"name\":\"validRegistration\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"ValidatorIneligibleForRewards\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotPoS\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroWeightToValueFactor\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516163e83803806163e883398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b61629b8061014d5f395ff3fe608060405234801561000f575f80fd5b50600436106101d7575f3560e01c806380dd672f11610101578063ba3a4b971161009a578063ba3a4b971461040a578063bc5fbfec1461041d578063bee0a03f14610444578063c974d1b614610457578063d5f20ff61461045f578063df93d8de1461047f578063e4a63c4014610489578063f74c607b146104b0578063fd7ac5e7146104c3575f80fd5b806380dd672f1461036c5780638280a25a1461037f57806393e24598146103875780639e1bc4ef1461039a578063a3a65e48146103ad578063a9778a7a14610299578063afb98096146103c0578063afba878a146103e7578063b771b3bc146103ef575f80fd5b80633a1cfff6116101735780633a1cfff6146102b5578063467ef06f146102c85780634bee0040146102db57806360305d62146102fc57806360ad778414610319578063620658561461032c57806366435abf1461033f578063732214f81461035257806376f7862114610359575f80fd5b80630118acc4146101db5780630322ed98146101f0578063151d30d1146102035780631ec447241461022257806320d91b7a1461023557806325e1c776146102485780632893d0771461025b5780632e2194d81461026e57806335455ded14610299575b5f80fd5b6101ee6101e9366004615537565b6104d6565b005b6101ee6101fe366004615572565b6104e8565b61020b600a81565b60405160ff90911681526020015b60405180910390f35b6101ee610230366004615537565b6106f4565b6101ee610243366004615589565b610700565b6101ee6102563660046155d7565b610bfe565b6101ee6102693660046155d7565b610c72565b61028161027c366004615572565b610f94565b6040516001600160401b039091168152602001610219565b6102a261271081565b60405161ffff9091168152602001610219565b6101ee6102c3366004615537565b610fe8565b6101ee6102d63660046155f8565b610ff4565b6102ee6102e9366004615638565b6110d0565b604051908152602001610219565b610304601481565b60405163ffffffff9091168152602001610219565b6101ee6103273660046155d7565b6110f7565b6102ee61033a36600461569d565b611352565b61028161034d366004615572565b611372565b6102ee5f81565b6101ee610367366004615537565b611386565b6101ee61037a3660046155d7565b611393565b61020b603081565b6101ee610395366004615572565b61154e565b6102ee6103a83660046156b6565b6115f5565b6101ee6103bb3660046155f8565b611613565b6102ee7f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0081565b61020b605081565b6103fd6005600160991b0181565b60405161021991906156d6565b6101ee610418366004615572565b611799565b6102ee7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0081565b6101ee610452366004615572565b611977565b61020b601481565b61047261046d366004615572565b611a94565b604051610219919061575f565b6102816202a30081565b6102ee7f6e5bdfcce15e53c3406ea67bfce37dcd26f5152d5492824e43fd5e3c8ac5ab0081565b6101ee6104be3660046157f3565b611bd6565b6102ee6104d1366004615831565b611cb5565b6104e38383836001611ced565b505050565b5f6104f1612058565b5f838152600782016020526040808220815160e0810190925280549394509192909190829060ff16600581111561052a5761052a6156ea565b600581111561053b5761053b6156ea565b815260200160018201805461054f9061589c565b80601f016020809104026020016040519081016040528092919081815260200182805461057b9061589c565b80156105c65780601f1061059d576101008083540402835291602001916105c6565b820191905f5260205f20905b8154815290600101906020018083116105a957829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b909104811660808301526003928301541660a09091015290915081516005811115610631576106316156ea565b1461066d575f8381526007830160205260409081902054905163170cc93360e21b81526106649160ff16906004016158d4565b60405180910390fd5b6005600160991b016001600160a01b031663ee5b48eb6106928584606001515f61207c565b6040518263ffffffff1660e01b81526004016106ae91906158e2565b6020604051808303815f875af11580156106ca573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106ee91906158f4565b50505050565b6104e38383835f611ced565b5f610709612058565b600981015490915060ff161561073257604051637fab81e560e01b815260040160405180910390fd5b6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015610775573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061079991906158f4565b8360200135146107c2576040516372b0a7e760e11b815260208401356004820152602401610664565b306107d3606085016040860161590b565b6001600160a01b03161461080b576107f1606084016040850161590b565b604051632f88120d60e21b815260040161066491906156d6565b5f6108196060850185615926565b905090505f805b828163ffffffff161015610b00575f61083c6060880188615926565b8363ffffffff168181106108525761085261596b565b9050602002810190610864919061597f565b61086d90615aa3565b80516040519192505f91600888019161088591615b1c565b908152602001604051809103902054146108b557805160405163a41f772f60e01b815261066491906004016158e2565b5f6002885f0135846040516020016108e492919091825260e01b6001600160e01b031916602082015260240190565b60408051601f19818403018152908290526108fe91615b1c565b602060405180830381855afa158015610919573d5f803e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061093c91906158f4565b90508086600801835f01516040516109549190615b1c565b90815260408051602092819003830181209390935560e0830181526002835284518284015284810180516001600160401b03908116858401525f60608601819052915181166080860152421660a085015260c0840181905284815260078a01909252902081518154829060ff191660018360058111156109d6576109d66156ea565b0217905550602082015160018201906109ef9082615b71565b506040828101516002830180546060860151608087015160a08801516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c090930151600390920180546001600160401b03191692841692909217909155830151610a93911685615c40565b8251604051919550610aa491615b1c565b60408051918290038220908401516001600160401b031682529082907f9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf9060200160405180910390a3505080610af990615c53565b9050610820565b50600483018190556001830154606490610b2490600160401b900460ff1683615c75565b1015610b4657604051635943317f60e01b815260048101829052602401610664565b5f610b5c610b53866120d3565b604001516121de565b90505f610b688761236b565b90505f600282604051610b7b9190615b1c565b602060405180830381855afa158015610b96573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610bb991906158f4565b9050828114610be557604051631872fc8d60e01b81526004810182905260248101849052604401610664565b5050506009909201805460ff1916600117905550505050565b610c0782612540565b610c27576040516330efa98b60e01b815260048101839052602401610664565b5f610c3183611a94565b5190506002816005811115610c4857610c486156ea565b14610c68578060405163170cc93360e21b815260040161066491906158d4565b6106ee8383612569565b610c7a6127b7565b5f610c83612801565b90505f610c8f84611a94565b9050600281516005811115610ca657610ca66156ea565b14610cc757805160405163170cc93360e21b815261066491906004016158d4565b610cd084612540565b610cf0576040516330efa98b60e01b815260048101859052602401610664565b5f8481526005830160205260409020546001600160a01b03163314610d2b57335b604051636e2ccd7560e11b815260040161066491906156d6565b5f84815260058301602052604090205460a08201514291610d5d91600160b01b9091046001600160401b031690615c8c565b6001600160401b0316816001600160401b03161015610d9a5760405163fb6ce63f60e01b81526001600160401b0382166004820152602401610664565b5f610da58686612569565b5f87815260058601602052604081206001015491925090610dd690600160401b90046001600160401b031683615cac565b5f888152600587016020526040812060010154919250600160801b9091046001600160401b031690819003610e0c575060a08401515b610e17828286612825565b610e3757604051635bff683f60e11b815260048101899052602401610664565b600486015460408601515f916001600160a01b031690634f22429f90610e5c90611352565b848589886040518663ffffffff1660e01b8152600401610e80959493929190615ccc565b602060405180830381865afa158015610e9b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ebf91906158f4565b90506064605060ff168760a0015187610ed89190615cac565b610ee29190615cfa565b610eec9190615d39565b5f8a81526005890160205260409020600181018054600160401b600160c01b031916600160401b6001600160401b039485160267ffffffffffffffff60801b191617600160801b938916939093029290921790915554610f55906001600160a01b031682612870565b60405181908a907f69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a905f90a350505050505050610f906128e0565b5050565b5f80610f9e612801565b60030154610fac9084615d5e565b9050801580610fc157506001600160401b0381115b15610fe25760405163222d164360e21b815260048101849052602401610664565b92915050565b6104e38383835f612906565b610ffc6127b7565b5f611005612801565b90505f8061101284612b60565b9150915061101f82612540565b61102b575050506110c5565b5f8281526005840160205260409020546001600160a01b0316600482516005811115611059576110596156ea565b036110aa575f8381526008850160205260408120805491905561107c8282612870565b604051819085907f69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a905f90a3505b6110c0816110bb8460400151611352565b612e9a565b505050505b6110cd6128e0565b50565b5f6110d96127b7565b6110e585858585612eb8565b90506110ef6128e0565b949350505050565b5f611100612801565b5f848152600682016020526040808220815160e0810190925280549394509192909190829060ff166003811115611139576111396156ea565b600381111561114a5761114a6156ea565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f6111c082611a94565b90506001835160038111156111d7576111d76156ea565b146111f8578251604051633b0d540d60e21b81526106649190600401615d71565b60048151600581111561120d5761120d6156ea565b036112235761121b866130b5565b505050505050565b5f8061123a611231886120d3565b604001516132d5565b509150915081841461126757846040015160405163089938b360e11b815260040161066491815260200190565b806001600160401b031683606001516001600160401b031610806112a05750806001600160401b03168560a001516001600160401b0316115b156112c957604051632e19bc2d60e11b81526001600160401b0382166004820152602401610664565b5f888152600687016020908152604091829020805460ff1916600290811782550180546001600160401b034216600160401b810267ffffffffffffffff60401b1990921691909117909155915191825285918a917f047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6910160405180910390a35050505050505050565b5f61135b612801565b60030154610fe2906001600160401b038416615c75565b5f61137c82611a94565b6080015192915050565b6104e38383836001612906565b61139b6127b7565b5f6113a4612801565b5f848152600682016020526040808220815160e0810190925280549394509192909190829060ff1660038111156113dd576113dd6156ea565b60038111156113ee576113ee6156ea565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101529050600381516003811115611467576114676156ea565b14611488578051604051633b0d540d60e21b81526106649190600401615d71565b60046114978260400151611a94565b5160058111156114a9576114a96156ea565b1461153b575f6114b8846120d3565b90505f806114c983604001516132d5565b5091509150818460400151146114f55760405163089938b360e11b815260048101839052602401610664565b806001600160401b03168460c001516001600160401b0316111561153757604051632e19bc2d60e11b81526001600160401b0382166004820152602401610664565b5050505b611544846130b5565b5050610f906128e0565b5f611557612801565b90505f61156383611a94565b519050600481600581111561157a5761157a6156ea565b1461159a578060405163170cc93360e21b815260040161066491906158d4565b5f8381526005830160205260409020546001600160a01b031633146115bf5733610d11565b5f838152600883016020908152604080832080549084905560058601909252909120546106ee906001600160a01b031682612870565b5f6115fe6127b7565b61160983338461352b565b9050610fe26128e0565b5f61161c612058565b90505f8061163561162c856120d3565b60400151613800565b915091508061165b57604051632d07135360e01b81528115156004820152602401610664565b5f828152600684016020526040902080546116759061589c565b90505f036116995760405163089938b360e11b815260048101839052602401610664565b60015f83815260078501602052604090205460ff1660058111156116bf576116bf6156ea565b146116f2575f8281526007840160205260409081902054905163170cc93360e21b81526106649160ff16906004016158d4565b5f828152600684016020526040812061170a91615486565b5f828152600784016020908152604091829020805460ff1916600290811782550180546001600160401b0342818116600160c01b026001600160c01b0390931692909217928390558451600160801b9093041682529181019190915283917ff8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568910160405180910390a250505050565b5f6117a2612801565b5f838152600682016020526040808220815160e0810190925280549394509192909190829060ff1660038111156117db576117db6156ea565b60038111156117ec576117ec6156ea565b8152815461010090046001600160a01b0316602082015260018083015460408301526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015290915081516003811115611865576118656156ea565b141580156118865750600381516003811115611883576118836156ea565b14155b156118a7578051604051633b0d540d60e21b81526106649190600401615d71565b5f6118b58260400151611a94565b905080606001516001600160401b03165f036118e7576040516339b894f960e21b815260048101859052602401610664565b6005600160991b016001600160a01b031663ee5b48eb61191484604001518460600151856080015161207c565b6040518263ffffffff1660e01b815260040161193091906158e2565b6020604051808303815f875af115801561194c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061197091906158f4565b5050505050565b5f611980612058565b5f838152600682016020526040902080549192509061199e9061589c565b90505f036119c25760405163089938b360e11b815260048101839052602401610664565b60015f83815260078301602052604090205460ff1660058111156119e8576119e86156ea565b14611a1b575f8281526007820160205260409081902054905163170cc93360e21b81526106649160ff16906004016158d4565b5f82815260068201602052604090819020905163ee5b48eb60e01b81526005600160991b019163ee5b48eb91611a549190600401615d8b565b6020604051808303815f875af1158015611a70573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104e391906158f4565b611a9c6154bd565b5f611aa5612058565b5f84815260078201602052604090819020815160e0810190925280549293509091829060ff166005811115611adc57611adc6156ea565b6005811115611aed57611aed6156ea565b8152602001600182018054611b019061589c565b80601f0160208091040260200160405190810160405280929190818152602001828054611b2d9061589c565b8015611b785780601f10611b4f57610100808354040283529160200191611b78565b820191905f5260205f20905b815481529060010190602001808311611b5b57829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b9091048116608083015260039092015490911660a0909101529392505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff1680611c1f575080546001600160401b03808416911610155b15611c3d5760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b178155611c6884846139bc565b805460ff60401b191681556040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050565b5f80611cbf612058565b9050806008018484604051611cd5929190615e15565b90815260200160405180910390205491505092915050565b5f611cf6612801565b5f868152600682016020526040808220815160e0810190925280549394509192909190829060ff166003811115611d2f57611d2f6156ea565b6003811115611d4057611d406156ea565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f611db682611a94565b9050600283516003811115611dcd57611dcd6156ea565b14611dee578251604051633b0d540d60e21b81526106649190600401615d71565b60208301516001600160a01b03163314611e8a575f8281526005850160205260409020546001600160a01b03163314611e275733610d11565b5f82815260058501602052604090205460a0820151611e5691600160b01b90046001600160401b031690615c8c565b6001600160401b0316421015611e8a5760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610664565b600281516005811115611e9f57611e9f6156ea565b03611ff15760028401546080840151611ec1916001600160401b031690615c8c565b6001600160401b0316421015611ef55760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610664565b8615611f0757611f058287612569565b505b5f8881526006850160205260409020805460ff1916600317905560608301516080820151611f40918491611f3b9190615cac565b6139d6565b505f898152600686016020526040812060020180546001600160401b03909316600160c01b026001600160c01b0390931692909217909155611f8184613b20565b9050858015611f8e575080155b15611faf57604051631036cf9160e11b8152600481018a9052602401610664565b5f8981526007860160205260408082208390555184918b917f366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed579190a35061204e565b600481516005811115612006576120066156ea565b036120325761201483613b20565b5f89815260078601602052604090205561202d886130b5565b61204e565b805160405163170cc93360e21b815261066491906004016158d4565b5050505050505050565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0090565b6040515f6020820152600360e01b6022820152602681018490526001600160c01b031960c084811b8216604684015283901b16604e82015260609060560160405160208183030381529060405290505b9392505050565b60408051606080820183525f8083526020830152918101919091526040516306f8253560e41b815263ffffffff831660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa158015612137573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261215e9190810190615e2f565b915091508061218057604051636b2f19e960e01b815260040160405180910390fd5b8151156121a6578151604051636ba589a560e01b81526004810191909152602401610664565b60208201516001600160a01b0316156121d7578160200151604051624de75d60e31b815260040161066491906156d6565b5092915050565b5f815160261461221357815160405163cc92daa160e01b815263ffffffff909116600482015260266024820152604401610664565b5f805b60028110156122625761222a816001615f01565b612235906008615c75565b61ffff1684828151811061224b5761224b61596b565b016020015160f81c901b9190911790600101612216565b5061ffff81161561228c5760405163407b587360e01b815261ffff82166004820152602401610664565b5f805b60048110156122e7576122a3816003615f01565b6122ae906008615c75565b63ffffffff16856122c0836002615c40565b815181106122d0576122d061596b565b016020015160f81c901b919091179060010161228f565b5063ffffffff81161561230d57604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156123625761232481601f615f01565b61232f906008615c75565b8661233b836006615c40565b8151811061234b5761234b61596b565b016020015160f81c901b9190911790600101612310565b50949350505050565b60605f808335602085013560146123878787016040890161590b565b6123946060890189615926565b60405160f09790971b6001600160f01b0319166020880152602287019590955250604285019290925260e090811b6001600160e01b0319908116606286015260609290921b6001600160601b03191660668501529190911b16607a820152607e0160405160208183030381529060405290505f5b6124156060850185615926565b90508110156121d7578161242c6060860186615926565b8381811061243c5761243c61596b565b905060200281019061244e919061597f565b6124589080615f14565b90506124676060870187615926565b848181106124775761247761596b565b9050602002810190612489919061597f565b6124939080615f14565b6124a06060890189615926565b868181106124b0576124b061596b565b90506020028101906124c2919061597f565b6124d0906020810190615f14565b6124dd60608b018b615926565b888181106124ed576124ed61596b565b90506020028101906124ff919061597f565b61251090606081019060400161569d565b6040516020016125269796959493929190615f56565b60408051601f198184030181529190529150600101612408565b5f8061254a612801565b5f938452600501602052505060409020546001600160a01b0316151590565b6040516306f8253560e41b815263ffffffff821660048201525f90819081906005600160991b0190636f825350906024015f60405180830381865afa1580156125b4573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526125db9190810190615e2f565b91509150806125fd57604051636b2f19e960e01b815260040160405180910390fd5b6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015612640573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061266491906158f4565b82511461268a578151604051636ba589a560e01b81526004810191909152602401610664565b60208201516001600160a01b0316156126bb578160200151604051624de75d60e31b815260040161066491906156d6565b5f806126ca8460400151613ce0565b915091508187146126f15760405163089938b360e11b815260048101889052602401610664565b5f6126fa612801565b5f8981526005820160205260409020600101549091506001600160401b03908116908316111561278d575f88815260058201602090815260409182902060010180546001600160401b0319166001600160401b038616908117909155915191825289917fec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435910160405180910390a26127ac565b5f8881526005820160205260409020600101546001600160401b031691505b509695505050505050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f008054600119016127fb57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b7f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0090565b5f60506128328484615cac565b61283c9190615cfa565b6001600160401b0316612850856064615cfa565b6001600160401b0316101561286657505f6120cc565b5060019392505050565b5f612879613ed1565b80546040516340c10f1960e01b81526001600160a01b038681166004830152602482018690529293509116906340c10f19906044015f604051808303815f87803b1580156128c5575f80fd5b505af11580156128d7573d5f803e3d5ffd5b50505050505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f61290f612801565b90505f61291b86613ef5565b905061292686612540565b6129315750506106ee565b5f8681526005830160205260409020546001600160a01b031633146129565733610d11565b5f86815260058301602052604090205460a082015161298591600160b01b90046001600160401b031690615c8c565b6001600160401b03168160c001516001600160401b031610156129cc5760c081015160405163fb6ce63f60e01b81526001600160401b039091166004820152602401610664565b5f85156129e4576129dd8786612569565b9050612a02565b505f8681526005830160205260409020600101546001600160401b03165b5f878152600584016020526040812060010154612a2f90600160401b90046001600160401b031683615cac565b5f898152600586016020526040812060010154919250600160801b9091046001600160401b031690819003612a65575060a08301515b858015612a7e5750612a7c82828660c00151612825565b155b15612a9f57604051635bff683f60e11b8152600481018a9052602401610664565b600485015460408501515f916001600160a01b031690634f22429f90612ac490611352565b84858960c00151886040518663ffffffff1660e01b8152600401612aec959493929190615ccc565b602060405180830381865afa158015612b07573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b2b91906158f4565b905080866008015f8c81526020019081526020015f205f828254612b4f9190615c40565b909155505050505050505050505050565b5f612b696154bd565b5f612b72612058565b90505f80612b8261162c876120d3565b915091508015612ba957604051632d07135360e01b81528115156004820152602401610664565b5f828152600784016020526040808220815160e081019092528054829060ff166005811115612bda57612bda6156ea565b6005811115612beb57612beb6156ea565b8152602001600182018054612bff9061589c565b80601f0160208091040260200160405190810160405280929190818152602001828054612c2b9061589c565b8015612c765780601f10612c4d57610100808354040283529160200191612c76565b820191905f5260205f20905b815481529060010190602001808311612c5957829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b909104811660808301526003928301541660a09091015290915081516005811115612ce157612ce16156ea565b14158015612d025750600181516005811115612cff57612cff6156ea565b14155b15612d2357805160405163170cc93360e21b815261066491906004016158d4565b600381516005811115612d3857612d386156ea565b03612d465760048152612d4b565b600581525b836008018160200151604051612d619190615b1c565b90815260408051602092819003830190205f908190558581526007870190925290208151815483929190829060ff19166001836005811115612da557612da56156ea565b021790555060208201516001820190612dbe9082615b71565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c090920151600390910180546001600160401b0319169190921617905580516005811115612e6357612e636156ea565b60405184907f1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16905f90a39196919550909350505050565b610f908282612ea7613ed1565b546001600160a01b031691906141cd565b5f80612ec2612801565b600281015490915061ffff600160401b90910481169086161080612eeb575061271061ffff8616115b15612f0f57604051635f12e6c360e11b815261ffff86166004820152602401610664565b60028101546001600160401b039081169085161015612f4b576040516202a06d60e11b81526001600160401b0385166004820152602401610664565b8054831080612f5d5750806001015483115b15612f7e5760405163222d164360e21b815260048101849052602401610664565b5f612f888461422c565b90505f612f9482610f94565b90505f612fa18983614249565b90506040518060c00160405280612fb53390565b6001600160a01b03908116825261ffff808c166020808501919091526001600160401b03808d166040808701919091525f6060808801829052608080890183905260a09889018390528a83526005909d0185529082902088518154958a0151938a01518516600160b01b0267ffffffffffffffff60b01b1994909716600160a01b026001600160b01b03199096169716969096179390931716929092178355840151600190920180549885015194909301518116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b031990991692909116919091179690961791909116949094179093555090915050949350505050565b5f6130be612801565b5f838152600682016020526040808220815160e0810190925280549394509192909190829060ff1660038111156130f7576130f76156ea565b6003811115613108576131086156ea565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015281015190915061317c6147a6565b826080015161318b9190615c8c565b6001600160401b03164210156131bf5760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610664565b5f848152600684016020908152604080832080546001600160a81b0319168155600181018490556002018390556007860190915281208054908290559080821561327a575f8481526005870160205260409020546127109061322c90600160a01b900461ffff1685615c75565b6132369190615d5e565b915081866008015f8681526020019081526020015f205f82825461325a9190615c40565b9091555061326a90508284615f01565b905061327a856020015182612870565b61328f85602001516110bb8760600151611352565b6040805182815260208101849052859189917f8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993910160405180910390a350505050505050565b5f805f835160361461330c57835160405163cc92daa160e01b815263ffffffff909116600482015260366024820152604401610664565b5f805b600281101561335b57613323816001615f01565b61332e906008615c75565b61ffff168682815181106133445761334461596b565b016020015160f81c901b919091179060010161330f565b5061ffff8116156133855760405163407b587360e01b815261ffff82166004820152602401610664565b5f805b60048110156133e05761339c816003615f01565b6133a7906008615c75565b63ffffffff16876133b9836002615c40565b815181106133c9576133c961596b565b016020015160f81c901b9190911790600101613388565b5063ffffffff811660031461340857604051635b60892f60e01b815260040160405180910390fd5b5f805b602081101561345d5761341f81601f615f01565b61342a906008615c75565b88613436836006615c40565b815181106134465761344661596b565b016020015160f81c901b919091179060010161340b565b505f805b60088110156134bc57613475816007615f01565b613480906008615c75565b6001600160401b031689613495836026615c40565b815181106134a5576134a561596b565b016020015160f81c901b9190911790600101613461565b505f805b600881101561351b576134d4816007615f01565b6134df906008615c75565b6001600160401b03168a6134f483602e615c40565b815181106135045761350461596b565b016020015160f81c901b91909117906001016134c0565b5091989097509095509350505050565b5f80613535612801565b90505f61354461027c8561422c565b90505f61355087611a94565b905061355b87612540565b61357b576040516330efa98b60e01b815260048101889052602401610664565b600281516005811115613590576135906156ea565b146135b157805160405163170cc93360e21b815261066491906004016158d4565b5f8282608001516135c29190615c8c565b905083600201600a9054906101000a90046001600160401b031682604001516135eb9190615cfa565b6001600160401b0316816001600160401b0316111561362857604051636d51fe0560e11b81526001600160401b0382166004820152602401610664565b5f806136348a846139d6565b915091505f8a8360405160200161366292919091825260c01b6001600160c01b031916602082015260280190565b60408051601f19818403018152828252805160209091012060e08301909152915080600181526001600160a01b038c1660208083019190915260408083018f90526001600160401b03808b1660608501525f6080850181905290881660a085015260c090930183905284835260068b01909152902081518154829060ff191660018360038111156136f5576136f56156ea565b02179055506020828101518254610100600160a81b0319166101006001600160a01b039283160217835560408085015160018501556060808601516002909501805460808089015160a08a015160c0909a01516001600160401b03998a166001600160801b031990941693909317600160401b918a1691909102176001600160801b0316600160801b998916999099026001600160c01b031698909817600160c01b91881691909102179055815189861681528a861694810194909452938b1690830152918101859052908c16918d9184917fb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426910160405180910390a49a9950505050505050505050565b5f80825160271461383657825160405163cc92daa160e01b815263ffffffff909116600482015260276024820152604401610664565b5f805b60028110156138855761384d816001615f01565b613858906008615c75565b61ffff1685828151811061386e5761386e61596b565b016020015160f81c901b9190911790600101613839565b5061ffff8116156138af5760405163407b587360e01b815261ffff82166004820152602401610664565b5f805b600481101561390a576138c6816003615f01565b6138d1906008615c75565b63ffffffff16866138e3836002615c40565b815181106138f3576138f361596b565b016020015160f81c901b91909117906001016138b2565b5063ffffffff811660021461393257604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156139875761394981601f615f01565b613954906008615c75565b87613960836006615c40565b815181106139705761397061596b565b016020015160f81c901b9190911790600101613935565b505f8660268151811061399c5761399c61596b565b016020015191976001600160f81b03199092161515965090945050505050565b6139c46147c1565b6139cd8261480c565b610f9081614880565b5f805f6139e1612058565b5f868152600782016020526040902060020154909150600160801b90046001600160401b0316613a1185826148dc565b5f613a1b87614b01565b5f8881526007850160205260408120600201805467ffffffffffffffff60801b1916600160801b6001600160401b038b16021790559091506005600160991b0163ee5b48eb613a6b8a858b61207c565b6040518263ffffffff1660e01b8152600401613a8791906158e2565b6020604051808303815f875af1158015613aa3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613ac791906158f4565b604080516001600160401b038a811682526020820184905282519394508516928b927f07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df928290030190a3909450925050505b9250929050565b5f80613b2a612801565b90505f613b3a8460400151611a94565b90505f600382516005811115613b5257613b526156ea565b1480613b705750600482516005811115613b6e57613b6e6156ea565b145b15613b80575060c0810151613bbd565b600282516005811115613b9557613b956156ea565b03613ba1575042613bbd565b815160405163170cc93360e21b815261066491906004016158d4565b84608001516001600160401b0316816001600160401b031611613be457505f949350505050565b6040808601515f90815260058501602052206001015460a0830151613c13916001600160401b03169083612825565b613c2157505f949350505050565b600483015460608601516001600160a01b0390911690634f22429f90613c4690611352565b60a085015160808901516040808b01515f90815260058a016020528190206001015490516001600160e01b031960e087901b168152613c989493929188916001600160401b0390911690600401615ccc565b602060405180830381865afa158015613cb3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613cd791906158f4565b95945050505050565b5f808251602e14613d1657825160405163cc92daa160e01b815263ffffffff9091166004820152602e6024820152604401610664565b5f805b6002811015613d6557613d2d816001615f01565b613d38906008615c75565b61ffff16858281518110613d4e57613d4e61596b565b016020015160f81c901b9190911790600101613d19565b5061ffff811615613d8f5760405163407b587360e01b815261ffff82166004820152602401610664565b5f805b6004811015613dea57613da6816003615f01565b613db1906008615c75565b63ffffffff1686613dc3836002615c40565b81518110613dd357613dd361596b565b016020015160f81c901b9190911790600101613d92565b5063ffffffff811615613e1057604051635b60892f60e01b815260040160405180910390fd5b5f805b6020811015613e6557613e2781601f615f01565b613e32906008615c75565b87613e3e836006615c40565b81518110613e4e57613e4e61596b565b016020015160f81c901b9190911790600101613e13565b505f805b6008811015613ec457613e7d816007615f01565b613e88906008615c75565b6001600160401b031688613e9d836026615c40565b81518110613ead57613ead61596b565b016020015160f81c901b9190911790600101613e69565b5090969095509350505050565b7f6e5bdfcce15e53c3406ea67bfce37dcd26f5152d5492824e43fd5e3c8ac5ab0090565b613efd6154bd565b5f613f06612058565b5f848152600782016020526040808220815160e0810190925280549394509192909190829060ff166005811115613f3f57613f3f6156ea565b6005811115613f5057613f506156ea565b8152602001600182018054613f649061589c565b80601f0160208091040260200160405190810160405280929190818152602001828054613f909061589c565b8015613fdb5780601f10613fb257610100808354040283529160200191613fdb565b820191905f5260205f20905b815481529060010190602001808311613fbe57829003601f168201915b50505091835250506002828101546001600160401b038082166020850152600160401b820481166040850152600160801b820481166060850152600160c01b9091048116608084015260039093015490921660a09091015290915081516005811115614049576140496156ea565b1461407c575f8481526007830160205260409081902054905163170cc93360e21b81526106649160ff16906004016158d4565b60038152426001600160401b031660c08201525f84815260078301602052604090208151815483929190829060ff191660018360058111156140c0576140c06156ea565b0217905550602082015160018201906140d99082615b71565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c090920151600390910180546001600160401b031916919092161790555f61417685826139d6565b6080840151604080516001600160401b03909216825242602083015291935083925087917f13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67910160405180910390a3509392505050565b6040516001600160a01b038381166024830152604482018390526104e391859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050614b6a565b5f610fe282614239613ed1565b546001600160a01b031690614bc2565b5f614252612058565b6009015460ff1661427657604051637fab81e560e01b815260040160405180910390fd5b5f61427f612058565b905042614292606086016040870161569d565b6001600160401b03161115806142cc57506142b06202a30042615c40565b6142c0606086016040870161569d565b6001600160401b031610155b15614306576142e1606085016040860161569d565b604051635879da1360e11b81526001600160401b039091166004820152602401610664565b61431b6143166060860186615fbf565b614d23565b61432b6143166080860186615fbf565b603061433a6020860186615f14565b90501461436c5761434e6020850185615f14565b6040516326475b2f60e11b8152610664925060040190815260200190565b6143768480615f14565b90505f036143a3576143888480615f14565b604051633e08a12560e11b8152600401610664929190615fd3565b5f600882016143b28680615f14565b6040516143c0929190615e15565b908152602001604051809103902054146143f9576143de8480615f14565b60405163a41f772f60e01b8152600401610664929190615fd3565b614403835f6148dc565b6040805160e08101909152815481525f90819061450f90602081016144288980615f14565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250505090825250602090810190614470908a018a615f14565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505050908252506020016144b960608a0160408b0161569d565b6001600160401b031681526020016144d460608a018a615fbf565b6144dd90616001565b81526020016144ef60808a018a615fbf565b6144f890616001565b8152602001876001600160401b0316815250614e8c565b5f8281526006860160205260409020919350915061452d8282615b71565b50816008840161453d8880615f14565b60405161454b929190615e15565b9081526040519081900360200181209190915563ee5b48eb60e01b81525f906005600160991b019063ee5b48eb906145879085906004016158e2565b6020604051808303815f875af11580156145a3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906145c791906158f4565b6040805160e0810190915290915080600181526020016145e78980615f14565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052509385525050506001600160401b0389166020808401829052604080850184905260608501929092526080840183905260a0909301829052868252600788019092522081518154829060ff19166001836005811115614676576146766156ea565b02179055506020820151600182019061468f9082615b71565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c090920151600390910180546001600160401b031916919092161790558061472c8880615f14565b60405161473a929190615e15565b6040518091039020847fb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430898b6040016020810190614778919061569d565b604080516001600160401b0393841681529290911660208301520160405180910390a4509095945050505050565b5f6147af612058565b600101546001600160401b0316919050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661480a57604051631afcd79f60e31b815260040160405180910390fd5b565b6148146147c1565b61481d81615079565b614825615092565b6110cd6060820135608083013561484260c0850160a0860161569d565b61485260e0860160c087016160c8565b614863610100870160e088016160e1565b61010087013561487b61014089016101208a0161590b565b6150a2565b6148886147c1565b5f614891613ed1565b90506001600160a01b0382166148bc5781604051637330680360e01b815260040161066491906156d6565b80546001600160a01b0319166001600160a01b0392909216919091179055565b5f6148e5612058565b90505f826001600160401b0316846001600160401b031611156149135761490c8385615cac565b9050614920565b61491d8484615cac565b90505b6040805160808101825260028401548082526003850154602083015260048501549282019290925260058401546001600160401b031660608201524291158061498257506001840154815161497e916001600160401b031690615c40565b8210155b156149a8576001600160401b0383166060820152818152604081015160208201526149c7565b82816060018181516149ba9190615c8c565b6001600160401b03169052505b60608101516149d7906064615cfa565b602082015160018601546001600160401b039290921691614a029190600160401b900460ff16615c75565b1015614a3257606081015160405163dfae880160e01b81526001600160401b039091166004820152602401610664565b856001600160401b031681604001818151614a4d9190615c40565b9052506040810180516001600160401b0387169190614a6d908390615f01565b90525060018401546040820151606491614a9291600160401b90910460ff1690615c75565b1015614ab9578060400151604051635943317f60e01b815260040161066491815260200190565b80516002850155602081015160038501556040810151600485015560600151600590930180546001600160401b0319166001600160401b039094169390931790925550505050565b5f80614b0b612058565b5f8481526007820160205260409020600201805491925090600890614b3f90600160401b90046001600160401b0316616101565b91906101000a8154816001600160401b0302191690836001600160401b031602179055915050919050565b5f614b7e6001600160a01b03841683615238565b905080515f14158015614ba2575080806020019051810190614ba0919061611c565b155b156104e35782604051635274afe760e01b815260040161066491906156d6565b5f80836001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401614bf091906156d6565b602060405180830381865afa158015614c0b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614c2f91906158f4565b9050614c466001600160a01b038516333086615245565b6040516370a0823160e01b81525f906001600160a01b038616906370a0823190614c749030906004016156d6565b602060405180830381865afa158015614c8f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614cb391906158f4565b9050818111614d195760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610664565b613cd78282615f01565b614d3060208201826155f8565b63ffffffff16158015614d505750614d4b6020820182615926565b151590505b15614d9757614d6260208201826155f8565b614d6f6020830183615926565b60405163c08a0f1d60e01b815263ffffffff9093166004840152602483015250604401610664565b614da46020820182615926565b9050614db360208301836155f8565b63ffffffff161115614dcc57614d6260208201826155f8565b60015b614ddc6020830183615926565b9050811015610f9057614df26020830183615926565b614dfd600184615f01565b818110614e0c57614e0c61596b565b9050602002016020810190614e21919061590b565b6001600160a01b0316614e376020840184615926565b83818110614e4757614e4761596b565b9050602002016020810190614e5c919061590b565b6001600160a01b03161015614e8457604051630dbc8d5f60e31b815260040160405180910390fd5b600101614dcf565b5f6060826040015151603014614eb55760405163180ffa0d60e01b815260040160405180910390fd5b82516020808501518051604080880151606089015160808a01518051908701515193515f98614ef6988a986001989297929690959094909390929101616137565b60405160208183030381529060405290505f5b84608001516020015151811015614f6857818560800151602001518281518110614f3557614f3561596b565b6020026020010151604051602001614f4e9291906161f1565b60408051601f198184030181529190529150600101614f09565b5060a0840151805160209182015151604051614f88938593929101616222565b60405160208183030381529060405290505f5b8460a001516020015151811015614ffa57818560a00151602001518281518110614fc757614fc761596b565b6020026020010151604051602001614fe09291906161f1565b60408051601f198184030181529190529150600101614f9b565b5060c084015160405161501191839160200161625d565b60405160208183030381529060405290506002816040516150329190615b1c565b602060405180830381855afa15801561504d573d5f803e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061507091906158f4565b94909350915050565b6150816147c1565b61508961527e565b6110cd81615286565b61509a6147c1565b61480a615369565b6150aa6147c1565b5f6150b3612801565b905061ffff851615806150cb575061271061ffff8616115b156150ef57604051635f12e6c360e11b815261ffff86166004820152602401610664565b868811156151135760405163222d164360e21b815260048101899052602401610664565b60ff841615806151265750600a60ff8516115b156151495760405163170db35960e31b815260ff85166004820152602401610664565b6151516147a6565b6001600160401b0316866001600160401b0316101561518d576040516202a06d60e11b81526001600160401b0387166004820152602401610664565b825f036151ad5760405163a733007160e01b815260040160405180910390fd5b96875560018701959095556002860180546001600160401b039590951669ffffffffffffffffffff1990951694909417600160401b61ffff94909416939093029290921767ffffffffffffffff60501b191660ff91909116600160501b02179091556003830155600490910180546001600160a01b0319166001600160a01b03909216919091179055565b60606120cc83835f615371565b6040516001600160a01b0384811660248301528381166044830152606482018390526106ee9186918216906323b872dd906084016141fa565b61480a6147c1565b61528e6147c1565b5f615297612058565b82358155905060146152af60608401604085016160e1565b60ff1611806152ce57506152c960608301604084016160e1565b60ff16155b15615302576152e360608301604084016160e1565b604051634a59bbff60e11b815260ff9091166004820152602401610664565b61531260608301604084016160e1565b60018201805460ff92909216600160401b0260ff60401b19909216919091179055615343604083016020840161569d565b60019190910180546001600160401b0319166001600160401b0390921691909117905550565b6128e06147c1565b606081471015615396573060405163cd78605960e01b815260040161066491906156d6565b5f80856001600160a01b031684866040516153b19190615b1c565b5f6040518083038185875af1925050503d805f81146153eb576040519150601f19603f3d011682016040523d82523d5f602084013e6153f0565b606091505b509150915061540086838361540a565b9695505050505050565b60608261541f5761541a8261545d565b6120cc565b815115801561543657506001600160a01b0384163b155b156154565783604051639996b31560e01b815260040161066491906156d6565b50806120cc565b80511561546d5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5080546154929061589c565b5f825580601f106154a1575050565b601f0160209004905f5260205f20908101906110cd91906154fa565b6040805160e08101909152805f81526060602082018190525f604083018190529082018190526080820181905260a0820181905260c09091015290565b5b8082111561550e575f81556001016154fb565b5090565b80151581146110cd575f80fd5b803563ffffffff81168114615532575f80fd5b919050565b5f805f60608486031215615549575f80fd5b83359250602084013561555b81615512565b91506155696040850161551f565b90509250925092565b5f60208284031215615582575f80fd5b5035919050565b5f806040838503121561559a575f80fd5b82356001600160401b038111156155af575f80fd5b8301608081860312156155c0575f80fd5b91506155ce6020840161551f565b90509250929050565b5f80604083850312156155e8575f80fd5b823591506155ce6020840161551f565b5f60208284031215615608575f80fd5b6120cc8261551f565b803561ffff81168114615532575f80fd5b80356001600160401b0381168114615532575f80fd5b5f805f806080858703121561564b575f80fd5b84356001600160401b03811115615660575f80fd5b850160a08188031215615671575f80fd5b935061567f60208601615611565b925061568d60408601615622565b9396929550929360600135925050565b5f602082840312156156ad575f80fd5b6120cc82615622565b5f80604083850312156156c7575f80fd5b50508035926020909101359150565b6001600160a01b0391909116815260200190565b634e487b7160e01b5f52602160045260245ffd5b6006811061570e5761570e6156ea565b9052565b5f5b8381101561572c578181015183820152602001615714565b50505f910152565b5f815180845261574b816020860160208601615712565b601f01601f19169290920160200192915050565b602081526157716020820183516156fe565b5f602083015160e0604084015261578c610100840182615734565b905060408401516001600160401b0380821660608601528060608701511660808601528060808701511660a08601528060a08701511660c08601528060c08701511660e086015250508091505092915050565b6001600160a01b03811681146110cd575f80fd5b5f80828403610160811215615806575f80fd5b61014080821215615815575f80fd5b8493508301359050615826816157df565b809150509250929050565b5f8060208385031215615842575f80fd5b82356001600160401b0380821115615858575f80fd5b818501915085601f83011261586b575f80fd5b813581811115615879575f80fd5b86602082850101111561588a575f80fd5b60209290920196919550909350505050565b600181811c908216806158b057607f821691505b6020821081036158ce57634e487b7160e01b5f52602260045260245ffd5b50919050565b60208101610fe282846156fe565b602081525f6120cc6020830184615734565b5f60208284031215615904575f80fd5b5051919050565b5f6020828403121561591b575f80fd5b81356120cc816157df565b5f808335601e1984360301811261593b575f80fd5b8301803591506001600160401b03821115615954575f80fd5b6020019150600581901b3603821315613b19575f80fd5b634e487b7160e01b5f52603260045260245ffd5b5f8235605e19833603018112615993575f80fd5b9190910192915050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b03811182821017156159d3576159d361599d565b60405290565b604080519081016001600160401b03811182821017156159d3576159d361599d565b604051601f8201601f191681016001600160401b0381118282101715615a2357615a2361599d565b604052919050565b5f6001600160401b03821115615a4357615a4361599d565b50601f01601f191660200190565b5f82601f830112615a60575f80fd5b8135615a73615a6e82615a2b565b6159fb565b818152846020838601011115615a87575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60608236031215615ab3575f80fd5b615abb6159b1565b82356001600160401b0380821115615ad1575f80fd5b615add36838701615a51565b83526020850135915080821115615af2575f80fd5b50615aff36828601615a51565b602083015250615b1160408401615622565b604082015292915050565b5f8251615993818460208701615712565b601f8211156104e357805f5260205f20601f840160051c81016020851015615b525750805b601f840160051c820191505b81811015611970575f8155600101615b5e565b81516001600160401b03811115615b8a57615b8a61599d565b615b9e81615b98845461589c565b84615b2d565b602080601f831160018114615bd1575f8415615bba5750858301515b5f19600386901b1c1916600185901b17855561121b565b5f85815260208120601f198616915b82811015615bff57888601518255948401946001909101908401615be0565b5085821015615c1c57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610fe257610fe2615c2c565b5f63ffffffff808316818103615c6b57615c6b615c2c565b6001019392505050565b8082028115828204841417610fe257610fe2615c2c565b6001600160401b038181168382160190808211156121d7576121d7615c2c565b6001600160401b038281168282160390808211156121d7576121d7615c2c565b9485526001600160401b03938416602086015291831660408501528216606084015216608082015260a00190565b6001600160401b03818116838216028082169190828114615d1d57615d1d615c2c565b505092915050565b634e487b7160e01b5f52601260045260245ffd5b5f6001600160401b0380841680615d5257615d52615d25565b92169190910492915050565b5f82615d6c57615d6c615d25565b500490565b6020810160048310615d8557615d856156ea565b91905290565b5f60208083525f8454615d9d8161589c565b806020870152604060018084165f8114615dbe5760018114615dda57615e07565b60ff19851660408a0152604084151560051b8a01019550615e07565b895f5260205f205f5b85811015615dfe5781548b8201860152908301908801615de3565b8a016040019650505b509398975050505050505050565b818382375f9101908152919050565b805161553281615512565b5f8060408385031215615e40575f80fd5b82516001600160401b0380821115615e56575f80fd5b9084019060608287031215615e69575f80fd5b615e716159b1565b82518152602080840151615e84816157df565b82820152604084015183811115615e99575f80fd5b80850194505087601f850112615ead575f80fd5b83519250615ebd615a6e84615a2b565b8381528882858701011115615ed0575f80fd5b615edf84838301848801615712565b80604084015250819550615ef4818801615e24565b9450505050509250929050565b81810381811115610fe257610fe2615c2c565b5f808335601e19843603018112615f29575f80fd5b8301803591506001600160401b03821115615f42575f80fd5b602001915036819003821315613b19575f80fd5b5f8851615f67818460208d01615712565b60e089901b6001600160e01b031916908301908152868860048301378681019050600481015f8152858782375060c09390931b6001600160c01b0319166004939094019283019390935250600c019695505050505050565b5f8235603e19833603018112615993575f80fd5b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b5f60408236031215616011575f80fd5b6160196159d9565b6160228361551f565b81526020808401356001600160401b038082111561603e575f80fd5b9085019036601f830112616050575f80fd5b8135818111156160625761606261599d565b8060051b91506160738483016159fb565b818152918301840191848101903684111561608c575f80fd5b938501935b838510156160b657843592506160a6836157df565b8282529385019390850190616091565b94860194909452509295945050505050565b5f602082840312156160d8575f80fd5b6120cc82615611565b5f602082840312156160f1575f80fd5b813560ff811681146120cc575f80fd5b5f6001600160401b03808316818103615c6b57615c6b615c2c565b5f6020828403121561612c575f80fd5b81516120cc81615512565b61ffff60f01b8a60f01b1681525f63ffffffff60e01b808b60e01b166002840152896006840152808960e01b16602684015250865161617d81602a850160208b01615712565b86519083019061619481602a840160208b01615712565b60c087901b6001600160c01b031916602a92909101918201526161c6603282018660e01b6001600160e01b0319169052565b6161df603682018560e01b6001600160e01b0319169052565b603a019b9a5050505050505050505050565b5f8351616202818460208801615712565b60609390931b6001600160601b0319169190920190815260140192915050565b5f8451616233818460208901615712565b6001600160e01b031960e095861b8116919093019081529290931b16600482015260080192915050565b5f835161626e818460208801615712565b60c09390931b6001600160c01b031916919092019081526008019291505056fea164736f6c6343000819000a",
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

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCaller) UPTIMEREWARDSTHRESHOLDPERCENTAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenStakingManager.contract.Call(opts, &out, "UPTIME_REWARDS_THRESHOLD_PERCENTAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) UPTIMEREWARDSTHRESHOLDPERCENTAGE() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.UPTIMEREWARDSTHRESHOLDPERCENTAGE(&_ERC20TokenStakingManager.CallOpts)
}

// UPTIMEREWARDSTHRESHOLDPERCENTAGE is a free data retrieval call binding the contract method 0xafba878a.
//
// Solidity: function UPTIME_REWARDS_THRESHOLD_PERCENTAGE() view returns(uint8)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) UPTIMEREWARDSTHRESHOLDPERCENTAGE() (uint8, error) {
	return _ERC20TokenStakingManager.Contract.UPTIMEREWARDSTHRESHOLDPERCENTAGE(&_ERC20TokenStakingManager.CallOpts)
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

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) ClaimValidationRewards(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "claimValidationRewards", validationID, messageIndex)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ClaimValidationRewards(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ClaimValidationRewards(&_ERC20TokenStakingManager.TransactOpts, validationID, messageIndex)
}

// ClaimValidationRewards is a paid mutator transaction binding the contract method 0x2893d077.
//
// Solidity: function claimValidationRewards(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) ClaimValidationRewards(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.ClaimValidationRewards(&_ERC20TokenStakingManager.TransactOpts, validationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteDelegatorRegistration(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) CompleteEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "completeEndDelegation", delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactorSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.CompleteEndDelegation(&_ERC20TokenStakingManager.TransactOpts, delegationID, messageIndex)
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

// ERC20TokenStakingManagerValidationRewardsClaimedIterator is returned from FilterValidationRewardsClaimed and is used to iterate over the raw logs and unpacked data for ValidationRewardsClaimed events raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationRewardsClaimedIterator struct {
	Event *ERC20TokenStakingManagerValidationRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *ERC20TokenStakingManagerValidationRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenStakingManagerValidationRewardsClaimed)
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
		it.Event = new(ERC20TokenStakingManagerValidationRewardsClaimed)
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
func (it *ERC20TokenStakingManagerValidationRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenStakingManagerValidationRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenStakingManagerValidationRewardsClaimed represents a ValidationRewardsClaimed event raised by the ERC20TokenStakingManager contract.
type ERC20TokenStakingManagerValidationRewardsClaimed struct {
	ValidationID [32]byte
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationRewardsClaimed is a free log retrieval operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) FilterValidationRewardsClaimed(opts *bind.FilterOpts, validationID [][32]byte, reward []*big.Int) (*ERC20TokenStakingManagerValidationRewardsClaimedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.FilterLogs(opts, "ValidationRewardsClaimed", validationIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenStakingManagerValidationRewardsClaimedIterator{contract: _ERC20TokenStakingManager.contract, event: "ValidationRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchValidationRewardsClaimed is a free log subscription operation binding the contract event 0x69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a.
//
// Solidity: event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) WatchValidationRewardsClaimed(opts *bind.WatchOpts, sink chan<- *ERC20TokenStakingManagerValidationRewardsClaimed, validationID [][32]byte, reward []*big.Int) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _ERC20TokenStakingManager.contract.WatchLogs(opts, "ValidationRewardsClaimed", validationIDRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenStakingManagerValidationRewardsClaimed)
				if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationRewardsClaimed", log); err != nil {
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
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerFilterer) ParseValidationRewardsClaimed(log types.Log) (*ERC20TokenStakingManagerValidationRewardsClaimed, error) {
	event := new(ERC20TokenStakingManagerValidationRewardsClaimed)
	if err := _ERC20TokenStakingManager.contract.UnpackLog(event, "ValidationRewardsClaimed", log); err != nil {
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
