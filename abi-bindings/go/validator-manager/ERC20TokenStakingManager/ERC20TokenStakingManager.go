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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADDRESS_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERC20_STAKING_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POS_VALIDATOR_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPTIME_REWARDS_THRESHOLD_PERCENTAGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimDelegationFees\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimValidationRewards\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeEndValidation\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitializeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structValidator\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"startingWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"endedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWeight\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structPoSValidatorManagerSettings\",\"components\":[{\"name\":\"baseSettings\",\"type\":\"tuple\",\"internalType\":\"structValidatorManagerSettings\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"churnPeriodSeconds\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"minimumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rewardCalculator\",\"type\":\"address\",\"internalType\":\"contractIRewardCalculator\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20Mintable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeDelegatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delegationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeValidatorRegistration\",\"inputs\":[{\"name\":\"registrationInput\",\"type\":\"tuple\",\"internalType\":\"structValidatorRegistrationInput\",\"components\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"disableOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeValidatorSet\",\"inputs\":[{\"name\":\"subnetConversionData\",\"type\":\"tuple\",\"internalType\":\"structSubnetConversionData\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialValidators\",\"type\":\"tuple[]\",\"internalType\":\"structInitialValidator[]\",\"components\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredValidators\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resendEndValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendRegisterValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendUpdateDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitUptimeProof\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"valueToWeight\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"weightToValue\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"DelegationEnded\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewards\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorAdded\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"delegatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"delegatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRegistered\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRemovalInitialized\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitialValidatorCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UptimeUpdated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"uptime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodEnded\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumValidatorStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodRegistered\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationRewardsClaimed\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reward\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRemovalInitialized\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorWeightUpdate\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DelegatorIneligibleForRewards\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSKeyLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationFee\",\"inputs\":[{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationID\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDelegatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitializationStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMaximumChurnPercentage\",\"inputs\":[{\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinStakeDuration\",\"inputs\":[{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeID\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPChainOwnerThreshold\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addressesLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRegistrationExpiry\",\"inputs\":[{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeAmount\",\"inputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeMultiplier\",\"inputs\":[{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidSubnetConversionID\",\"inputs\":[{\"name\":\"encodedSubnetConversionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedSubnetConversionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenAddress\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidValidationID\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorManagerAddress\",\"inputs\":[{\"name\":\"validatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorManagerBlockchainID\",\"inputs\":[{\"name\":\"blockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWarpOriginSenderAddress\",\"inputs\":[{\"name\":\"senderAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpSourceChainID\",\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MaxChurnRateExceeded\",\"inputs\":[{\"name\":\"churnAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MaxWeightExceeded\",\"inputs\":[{\"name\":\"newValidatorWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MinStakeDurationNotPassed\",\"inputs\":[{\"name\":\"endTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NodeAlreadyRegistered\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PChainOwnerAddressesNotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnexpectedRegistrationStatus\",\"inputs\":[{\"name\":\"validRegistration\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"ValidatorIneligibleForRewards\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotPoS\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x36303830363034303532333438303135363130303066353735663830666435623530363034303531363135663737333830333830363135663737383333393831303136303430383139303532363130303265393136313031303735363562363030313831363030313831313131353631303034323537363130303432363130313263353635623033363130303466353736313030346636313030353535363562353036313031343035363562376666306335376531363834306466303430663135303838646332663831666533393163333932336265633733653233613936363265666339633232396336613030383035343638303130303030303030303030303030303030393030343630666631363135363130306135353736303430353136336639326565386139363065303162383135323630303430313630343035313830393130333930666435623830353436303031363030313630343031623033393038313136313436313031303435373830353436303031363030313630343031623033313931363630303136303031363034303162303339303831313738323535363034303531393038313532376663376635303562326633373161653231373565653439313366343439396531663236333361376235393336333231656564316364616562363131353138316432393036303230303136303430353138303931303339306131356235303536356235663630323038323834303331323135363130313137353735663830666435623831353136303032383131303631303132353537356638306664356239333932353035303530353635623633346534383762373136306530316235663532363032313630303435323630323435666664356236313565326138303631303134643566333935666633666536303830363034303532333438303135363130303066353735663830666435623530363030343336313036313032303835373566333536306530316338303633373666373836323131313631303131663537383036336261336134623937313136313030613935373830363364356632306666363131363130303739353738303633643566323066663631343631303437353537383036336466393364386465313436313034393535373830363365346136336334303134363130343966353738303633663039393639616531343631303462333537383036336664376163356537313436313034633635373566383066643562383036336261336134623937313436313034333335373830363362633566626665633134363130343436353738303633626565306130336631343631303435613537383036336339373464316236313436313034366435373566383066643562383036333965316263346566313136313030656635373830363339653162633465663134363130336362353738303633613361363565343831343631303364653537383036336166623938303936313436313033663135373830363361666261383738613134363130343035353738303633623737316233626331343631303430643537356638306664356238303633373666373836323131343631303338613537383036333830646436373266313436313033396435373830363338323830613235613134363130336230353738303633393365323435393831343631303362383537356638306664356238303633333534353564656431313631303161303537383036333630333035643632313136313031373035373830363336303330356436323134363130333264353738303633363061643737383431343631303334613537383036333632303635383536313436313033356435373830363336363433356162663134363130333730353738303633373332323134663831343631303338333537356638306664356238303633333534353564656431343631303263613537383036333361316366666636313436313032653635373830363334363765663036663134363130326639353738303633346265653030343031343631303330633537356638306664356238303633323064393162376131313631303164623537383036333230643931623761313436313032363635373830363332356531633737363134363130323739353738303633323839336430373731343631303238633537383036333265323139346438313436313032396635373566383066643562383036333031313861636334313436313032306335373830363330333232656439383134363130323231353738303633313531643330643131343631303233343537383036333165633434373234313436313032353335373562356638306664356236313032316636313032316133363630303436313465326135363562363130346439353635623030356236313032316636313032326633363630303436313465363535363562363130346562353635623631303233633630306138313536356236303430353136306666393039313136383135323630323030313562363034303531383039313033393066333562363130323166363130323631333636303034363134653261353635623631303737623536356236313032316636313032373433363630303436313465376335363562363130373837353635623631303231663631303238373336363030343631346563613536356236313064346635363562363130323166363130323961333636303034363134656361353635623631306463333536356236313032623236313032616433363630303436313465363535363562363131313133353635623630343035313630303136303031363034303162303339303931313638313532363032303031363130323461353635623631303264333631323731303831353635623630343035313631666666663930393131363831353236303230303136313032346135363562363130323166363130326634333636303034363134653261353635623631313132393536356236313032316636313033303733363630303436313465656235363562363131313335353635623631303331663631303331613336363030343631346632393536356236313132323335363562363034303531393038313532363032303031363130323461353635623631303333353630313438313536356236303430353136336666666666666666393039313136383135323630323030313631303234613536356236313032316636313033353833363630303436313465636135363562363131323538353635623631303331663631303336623336363030343631346639303536356236313135333035363562363130326232363130333765333636303034363134653635353635623631313534393536356236313033316635663831353635623631303231663631303339383336363030343631346532613536356236313135356435363562363130323166363130336162333636303034363134656361353635623631313536613536356236313032336336303330383135363562363130323166363130336336333636303034363134653635353635623631313761623536356236313033316636313033643933363630303436313466616235363562363131383535353635623631303231663631303365633336363030343631346565623536356236313138383135363562363130333166356638303531363032303631356439653833333938313531393135323831353635623631303233633630353038313536356236313034316236303035363030313630393931623031383135363562363034303531363030313630303136306130316230333930393131363831353236303230303136313032346135363562363130323166363130343431333636303034363134653635353635623631316137363536356236313033316635663830353136303230363135646265383333393831353139313532383135363562363130323166363130343638333636303034363134653635353635623631316365313536356236313032336336303134383135363562363130343838363130343833333636303034363134653635353635623631316531653536356236303430353136313032346139313930363135303430353635623631303262323632303261333030383135363562363130333166356638303531363032303631356437653833333938313531393135323831353635623631303231663631303463313336363030343631353064343536356236313166366435363562363130333166363130346434333636303034363135313132353635623631323034633536356236313034653638333833383336303031363132306137353635623530353035303536356235663831383135323566383035313630323036313564666538333339383135313931353236303230353236303430383038323230383135313630653038313031393039323532383035343566383035313630323036313564626538333339383135313931353239333932393139303832393036306666313636303035383131313135363130353338353736313035333836313466636235363562363030353831313131353631303534393537363130353439363134666362353635623831353236303230303136303031383230313830353436313035356439303631353137643536356238303630316630313630323038303931303430323630323030313630343035313930383130313630343035323830393239313930383138313532363032303031383238303534363130353839393036313531376435363562383031353631303564343537383036303166313036313035616235373631303130303830383335343034303238333532393136303230303139313631303564343536356238323031393139303566353236303230356632303930356238313534383135323930363030313031393036303230303138303833313136313035623735373832393030333630316631363832303139313562353035303530393138333532353035303630303238323031353436303031363030313630343031623033383038323136363032303834303135323630303136303430316238323034383131363630343038343031353236303031363038303162383230343831313636303630383430313532363030313630633031623930393130343831313636303830383330313532363030333932383330313534313636306130393039313031353239303931353038313531363030353831313131353631303633663537363130363366363134666362353635623134363130363762353735663833383135323630303738333031363032303532363034303930383139303230353439303531363331373063633933333630653231623831353236313036373239313630666631363930363030343031363135316235353635623630343035313830393130333930666435623630363038313031353136303430353136333432613265306235363065313162383135323630303438313031383539303532363030313630303136303430316230333930393131363630323438323031353235663630343438323031353236303035363030313630393931623031393036336565356234386562393037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f3930363338353435633136613930363036343031356636303430353138303833303338313836356166343135383031353631303666323537336435663830336533643566666435623530353035303530363034303531336435663832336536303166336439303831303136303166313931363832303136303430353236313037313939313930383130313930363135326265353635623630343035313832363366666666666666663136363065303162383135323630303430313631303733353931393036313532656635363562363032303630343035313830383330333831356638373561663131353830313536313037353135373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130373735393139303631353330313536356235303530353035303536356236313034653638333833383335663631323061373536356237666539323534366436393839353064646433383931306432653135656431643932336364306137623364646539653261366133663338303536353535396362303935343566383035313630323036313564626538333339383135313931353239303630666631363135363130376439353736303430353136333766616238316535363065303162383135323630303430313630343035313830393130333930666435623630303536303031363039393162303136303031363030313630613031623033313636333432313363663738363034303531383136336666666666666666313636306530316238313532363030343031363032303630343035313830383330333831383635616661313538303135363130383163353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631303834303931393036313533303135363562383336303230303133353134363130383639353736303430353136333732623061376537363065313162383135323630323038343031333536303034383230313532363032343031363130363732353635623330363130383761363036303835303136303430383630313631353331383536356236303031363030313630613031623033313631343631303862643537363130383938363036303834303136303430383530313631353331383536356236303430353136333266383831323064363065323162383135323630303136303031363061303162303339303931313636303034383230313532363032343031363130363732353635623566363130386362363036303835303138353631353333333536356239303530393035303566383035623832383136336666666666666666313631303135363130626233353735663631303865653630363038383031383836313533333335363562383336336666666666666666313638313831313036313039303435373631303930343631353337383536356239303530363032303032383130313930363130393136393139303631353338633536356236313039316639303631353366373536356238303531363034303531393139323530356639313630303838383031393136313039333739313631353437323536356239303831353236303230303136303430353138303931303339303230353431343631303936373537383035313630343035313633613431663737326636306530316238313532363130363732393139303630303430313631353265663536356235663630303238383566303133353834363034303531363032303031363130393936393239313930393138323532363065303162363030313630303136306530316230333139313636303230383230313532363032343031393035363562363034303830353136303166313938313834303330313831353239303832393035323631303962303931363135343732353635623630323036303430353138303833303338313835356166613135383031353631303963623537336435663830336533643566666435623530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130396565393139303631353330313536356239303530383038363630303830313833356630313531363034303531363130613036393139303631353437323536356239303831353236303430383035313630323039323831393030333833303138313230393339303933353536306530383330313831353236303032383335323834353138323834303135323834383130313830353136303031363030313630343031623033393038313136383538343031353235663630363038363031383139303532393135313831313636303830383630313532343231363630613038353031353236306330383430313831393035323834383135323630303738613031393039323532393032303831353138313534383239303630666631393136363030313833363030353831313131353631306138383537363130613838363134666362353635623032313739303535353036303230383230313531363030313832303139303631306161313930383236313534633735363562353036303430383238313031353136303032383330313830353436303630383630313531363038303837303135313630613038383031353136303031363030313630343031623033393538363136363030313630303136303830316230333139393039343136393339303933313736303031363034303162393238363136393239303932303239313930393131373630303136303031363038303162303331363630303136303830316239313835313639313930393130323630303136303031363063303162303331363137363030313630633031623931383431363931393039313032313739303535363063303930393330313531363030333930393230313830353436376666666666666666666666666666666631393136393238343136393239303932313739303931353538333031353136313062343639313136383536313535393635363562383235313630343035313931393535303631306235373931363135343732353635623630343038303531393138323930303338323230393038343031353136303031363030313630343031623033313638323532393038323930376639643437666566396461303737363631353436653634366436313833306266636264613930353036633265356565643338313935653832633465623163626466393036303230303136303430353138303931303339306133353035303830363130626163393036313535613935363562393035303631303864323536356235303630303438333031383139303535356637335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363331653664393738393631306264663837363132336433353635623630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363130626666393139303631353265663536356236303230363034303531383038333033383138363561663431353830313536313063316135373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130633365393139303631353330313536356239303530356637335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f36333836326266613633383836303430353138323633666666666666666631363630653031623831353236303034303136313063373839313930363135366636353635623566363034303531383038333033383138363561663431353830313536313063393235373364356638303365336435666664356235303530353035303630343035313364356638323365363031663364393038313031363031663139313638323031363034303532363130636239393139303831303139303631353262653536356239303530356636303032383236303430353136313063636339313930363135343732353635623630323036303430353138303833303338313835356166613135383031353631306365373537336435663830336533643566666435623530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363130643061393139303631353330313536356239303530383238313134363130643336353736303430353136333138373266633864363065303162383135323630303438313031383239303532363032343831303138343930353236303434303136313036373235363562353035303530363030393930393230313830353436306666313931363630303131373930353535303530353035303536356236313064353838323631323465393536356236313064373835373630343035313633333065666139386236306530316238313532363030343831303138333930353236303234303136313036373235363562356636313064383238333631316531653536356235313930353036303032383136303035383131313135363130643939353736313064393936313466636235363562313436313064623935373830363034303531363331373063633933333630653231623831353236303034303136313036373239313930363135316235353635623631303737353833383336313235323435363562363130646362363132383061353635623566383035313630323036313564396538333339383135313931353235663631306465333834363131653165353635623930353036303032383135313630303538313131313536313064666135373631306466613631346663623536356231343631306531623537383035313630343035313633313730636339333336306532316238313532363130363732393139303630303430313631353162353536356236313065323438343631323465393536356236313065343435373630343035313633333065666139386236306530316238313532363030343831303138353930353236303234303136313036373235363562356638343831353236303034383330313630323035323630343039303230353436303031363030313630613031623033313633333134363130653861353733333562363034303531363336653263636437353630653131623831353236303031363030313630613031623033393039313136363030343832303135323630323430313631303637323536356235663834383135323630303438333031363032303532363034303930323035343630613038323031353134323931363130656263393136303031363062303162393039313034363030313630303136303430316230333136393036313537383135363562363030313630303136303430316230333136383136303031363030313630343031623033313631303135363130656639353736303430353136336662366365363366363065303162383135323630303136303031363034303162303338323136363030343832303135323630323430313631303637323536356235663631306630343836383636313235323435363562356638373831353236303034383630313630323035323630343038313230363030313031353439313932353039303631306633353930363030313630343031623930303436303031363030313630343031623033313638333631353761313536356235663838383135323630303438373031363032303532363034303831323036303031303135343931393235303630303136303830316239303931303436303031363030313630343031623033313639303831393030333631306636623537353036306130383430313531356236313066373638323832383636313238343135363562363130663936353736303430353136333562666636383366363065313162383135323630303438313031383939303532363032343031363130363732353635623630303338363031353436303430383630313531356639313630303136303031363061303162303331363930363337373863303662353930363130666262393036313135333035363562383438353839383835663830363034303531383836336666666666666666313636306530316238313532363030343031363130666533393739363935393439333932393139303631353763313536356236303230363034303531383038333033383138363561666131353830313536313066666535373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363131303232393139303631353330313536356239303530363036343630353036306666313638373630613030313531383736313130336239313930363135376131353635623631313034353931393036313537666635363562363131303466393139303631353833653536356235663861383135323630303438393031363032303532363034303930323036303031383130313830353437376666666666666666666666666666666666666666666666666666666666666666303030303030303030303030303030303139313636303031363034303162363030313630303136303430316230333934383531363032363766666666666666666666666666666666363038303162313931363137363030313630383031623933383931363933393039333032393239303932313739303931353535343631313063363930363030313630303136306130316230333136383236313238386435363562363034303531383139303861393037663639656233646439363032393837376137353631643562353037363831306530623533623232336339616234366362613865666430323865626630386665396139303566393061333530353035303530353035303530363131313066363030313566383035313630323036313564646538333339383135313931353235353536356235303530353635623566363131313233363465386434613531303030383336313538363335363562393239313530353035363562363130346536383338333833356636313239313635363562363131313364363132383061353635623566383035313630323036313564396538333339383135313931353235663830363131313536383436313262373735363562393135303931353036313131363338323631323465393536356236313131366635373530353035303631313230613536356235663832383135323630303438303835303136303230353236303430393039313230353436303031363030313630613031623033313639303832353136303035383131313135363131313965353736313131396536313466636235363562303336313131656635373566383338313532363030373835303136303230353236303430383132303830353439313930353536313131633138323832363132383864353635623630343035313831393038353930376636396562336464393630323938373761373536316435623530373638313065306235336232323363396162343663626138656664303238656266303866653961393035663930613335303562363131323035383136313132303038343630343030313531363131353330353635623631326632613536356235303530353035303562363131323230363030313566383035313630323036313564646538333339383135313931353235353536356235303536356235663631313232633631323830613536356236313132333838353835383538353631326634643536356239303530363131323530363030313566383035313630323036313564646538333339383135313931353235353536356239343933353035303530353035363562356638323831353235663830353136303230363135643565383333393831353139313532363032303532363034303830383232303831353136306530383130313930393235323830353435663830353136303230363135643965383333393831353139313532393339323931393038323930363066663136363030333831313131353631313261353537363131326135363134666362353635623630303338313131313536313132623635373631313262363631346663623536356238313532383135343631303130303930303436303031363030313630613031623033313636303230383230313532363030313832303135343630343038303833303139313930393135323630303239303932303135343630303136303031363034303162303338303832313636303630383430313532363030313630343031623832303438313136363038303834303135323630303136303830316238323034383131363630613038343031353236303031363063303162393039313034313636306330393039313031353238313031353139303931353035663631313332633832363131653165353635623930353036303031383335313630303338313131313536313133343335373631313334333631346663623536356231343631313336343537383235313630343035313633336230643534306436306532316238313532363130363732393139303630303430313631353837363536356236303034383135313630303538313131313536313133373935373631313337393631346663623536356230333631313338663537363131333837383636313331366235363562353035303530353035303530353635623566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363339646532336434303631313362343839363132336433353635623630343030313531363034303531383236336666666666666666313636306530316238313532363030343031363131336434393139303631353265663536356236303630363034303531383038333033383138363561663431353830313536313133656635373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363131343133393139303631353839303536356235303931353039313530383138343134363131343430353738343630343030313531363034303531363330383939333862333630653131623831353236303034303136313036373239313831353236303230303139303536356238303630303136303031363034303162303331363833363036303031353136303031363030313630343031623033313631303830363131343739353735303830363030313630303136303430316230333136383536306130303135313630303136303031363034303162303331363131356231353631313461323537363034303531363332653139626332643630653131623831353236303031363030313630343031623033383231363630303438323031353236303234303136313036373235363562356638383831353236303035383730313630323039303831353236303430393138323930323038303534363066663139313636303032393038313137383235353031383035343630303136303031363034303162303334323136363030313630343031623831303236666666666666666666666666666666666630303030303030303030303030303030313939303932313639313930393131373930393135353931353139313832353238353931386139313766303437303539623436353036396238623735313833366234316639663164383364616666353833643232333863633766626234363134333765633233613466363931303136303430353138303931303339306133353035303530353035303530353035303536356235663631313132333630303136303031363034303162303338333136363465386434613531303030363135386430353635623566363131353533383236313165316535363562363038303031353139323931353035303536356236313034653638333833383336303031363132393136353635623631313537323631323830613536356235663832383135323566383035313630323036313564356538333339383135313931353236303230353236303430383038323230383135313630653038313031393039323532383035343566383035313630323036313564396538333339383135313931353239333932393139303832393036306666313636303033383131313135363131356266353736313135626636313466636235363562363030333831313131353631313564303537363131356430363134666362353635623831353238313534363130313030393030343630303136303031363061303162303331363630323038323031353236303031383230313534363034303832303135323630303239303931303135343630303136303031363034303162303338303832313636303630383430313532363030313630343031623832303438313136363038303834303135323630303136303830316238323034383131363630613038343031353236303031363063303162393039313034313636306330393039313031353239303530363030333831353136303033383131313135363131363439353736313136343936313466636235363562313436313136366135373830353136303430353136333362306435343064363065323162383135323631303637323931393036303034303136313538373635363562363030343631313637393832363034303031353136313165316535363562353136303035383131313135363131363862353736313136386236313466636235363562313436313137386135373566363131363961383436313233643335363562393035303566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f363339646532336434303834363034303031353136303430353138323633666666666666666631363630653031623831353236303034303136313136643939313930363135326566353635623630363036303430353138303833303338313836356166343135383031353631313666343537336435663830336533643566666435623530353035303530363034303531336436303166313936303166383230313136383230313830363034303532353038313031393036313137313839313930363135383930353635623530393135303931353038313834363034303031353131343631313734343537363034303531363330383939333862333630653131623831353236303034383130313833393035323630323430313631303637323536356238303630303136303031363034303162303331363834363063303031353136303031363030313630343031623033313631313135363131373836353736303430353136333265313962633264363065313162383135323630303136303031363034303162303338323136363030343832303135323630323430313631303637323536356235303530353035623631313739333834363133313662353635623530353036313131306636303031356638303531363032303631356464653833333938313531393135323535353635623566383035313630323036313564396538333339383135313931353235663631313763333833363131653165353635623531393035303630303438313630303538313131313536313137646135373631313764613631346663623536356231343631313766613537383036303430353136333137306363393333363065323162383135323630303430313631303637323931393036313531623535363562356638333831353236303034383330313630323035323630343039303230353436303031363030313630613031623033313633333134363131383166353733333631306536353536356235663833383135323630303738333031363032303930383135323630343038303833323038303534393038343930353536303034383630313930393235323930393132303534363130373735393036303031363030313630613031623033313638323631323838643536356235663631313835653631323830613536356236313138363938333333383436313333353235363562393035303631313132333630303135663830353136303230363135646465383333393831353139313532353535363562356638303531363032303631356462653833333938313531393135323566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f36333265343363656235363131386234383636313233643335363562363034303031353136303430353138323633666666666666666631363630653031623831353236303034303136313138643439313930363135326566353635623630343038303531383038333033383138363561663431353830313536313138656535373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363131393132393139303631353865373536356239313530393135303830363131393338353736303430353136333264303731333533363065303162383135323831313531353630303438323031353236303234303136313036373235363562356638323831353236303036383430313630323035323630343039303230383035343631313935323930363135313764353635623930353035663033363131393736353736303430353136333038393933386233363065313162383135323630303438313031383339303532363032343031363130363732353635623630303135663833383135323630303738353031363032303532363034303930323035343630666631363630303538313131313536313139396335373631313939633631346663623536356231343631313963663537356638323831353236303037383430313630323035323630343039303831393032303534393035313633313730636339333336306532316238313532363130363732393136306666313639303630303430313631353162353536356235663832383135323630303638343031363032303532363034303831323036313139653739313631346437393536356235663832383135323630303738343031363032303930383135323630343039313832393032303830353436306666313931363630303239303831313738323535303138303534363030313630303136303430316230333432383138313136363030313630633031623032363030313630303136306330316230333930393331363932393039323137393238333930353538343531363030313630383031623930393330343136383235323931383130313931393039313532383339313766663866643163393066623963666132636132333538666466353830366230383661643433333135643932623232316339323965666337663130356365373536383931303136303430353138303931303339306132353035303530353035363562356638313831353235663830353136303230363135643565383333393831353139313532363032303532363034303830383232303831353136306530383130313930393235323830353435663830353136303230363135643965383333393831353139313532393339323931393038323930363066663136363030333831313131353631316163333537363131616333363134666362353635623630303338313131313536313161643435373631316164343631346663623536356238313532383135343631303130303930303436303031363030313630613031623033313636303230383230313532363030313830383330313534363034303833303135323630303239303932303135343630303136303031363034303162303338303832313636303630383430313532363030313630343031623832303438313136363038303834303135323630303136303830316238323034383131363630613038343031353236303031363063303162393039313034313636306330393039313031353239303931353038313531363030333831313131353631316234643537363131623464363134666362353635623134313538303135363131623665353735303630303338313531363030333831313131353631316236623537363131623662363134666362353635623134313535623135363131623866353738303531363034303531363333623064353430643630653231623831353236313036373239313930363030343031363135383736353635623566363131623964383236303430303135313631316531653536356239303530383036303630303135313630303136303031363034303162303331363566303336313162636635373630343035313633333962383934663936306532316238313532363030343831303138353930353236303234303136313036373235363562363034303830383330313531363036303833303135313630383038343031353139323531363334326132653062353630653131623831353236303035363030313630393931623031393336336565356234386562393337335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f39333633383534356331366139333631316333643933393036303034303139323833353236303031363030313630343031623033393138323136363032303834303135323136363034303832303135323630363030313930353635623566363034303531383038333033383138363561663431353830313536313163353735373364356638303365336435666664356235303530353035303630343035313364356638323365363031663364393038313031363031663139313638323031363034303532363131633765393139303831303139303631353262653536356236303430353138323633666666666666666631363630653031623831353236303034303136313163396139313930363135326566353635623630323036303430353138303833303338313566383735616631313538303135363131636236353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631316364613931393036313533303135363562353035303530353035303536356235663831383135323766653932353436643639383935306464643338393130643265313565643164393233636430613762336464653965326136613366333830353635353539636230363630323035323630343039303230383035343566383035313630323036313564626538333339383135313931353239313930363131643238393036313531376435363562393035303566303336313164346335373630343035313633303839393338623336306531316238313532363030343831303138333930353236303234303136313036373235363562363030313566383338313532363030373833303136303230353236303430393032303534363066663136363030353831313131353631316437323537363131643732363134666362353635623134363131646135353735663832383135323630303738323031363032303532363034303930383139303230353439303531363331373063633933333630653231623831353236313036373239313630666631363930363030343031363135316235353635623566383238313532363030363832303136303230353236303430393038313930323039303531363365653562343865623630653031623831353236303035363030313630393931623031393136336565356234386562393136313164646539313930363030343031363135393061353635623630323036303430353138303833303338313566383735616631313538303135363131646661353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631303465363931393036313533303135363562363131653236363134646230353635623566383238313532356638303531363032303631356466653833333938313531393135323630323035323630343039303831393032303831353136306530383130313930393235323830353435663830353136303230363135646265383333393831353139313532393239313930383239303630666631363630303538313131313536313165373335373631316537333631346663623536356236303035383131313135363131653834353736313165383436313466636235363562383135323630323030313630303138323031383035343631316539383930363135313764353635623830363031663031363032303830393130343032363032303031363034303531393038313031363034303532383039323931393038313831353236303230303138323830353436313165633439303631353137643536356238303135363131663066353738303630316631303631316565363537363130313030383038333534303430323833353239313630323030313931363131663066353635623832303139313930356635323630323035663230393035623831353438313532393036303031303139303630323030313830383331313631316566323537383239303033363031663136383230313931356235303530353039313833353235303530363030323832303135343630303136303031363034303162303338303832313636303230383430313532363030313630343031623832303438313136363034303834303135323630303136303830316238323034383131363630363038343031353236303031363063303162393039313034383131363630383038333031353236303033393039323031353439303931313636306130393039313031353239333932353035303530353635623766663063353765313638343064663034306631353038386463326638316665333931633339323362656337336532336139363632656663396332323963366130303830353436303032393139303630303136303430316239303034363066663136383036313166623635373530383035343630303136303031363034303162303338303834313639313136313031353562313536313166643435373630343035313633663932656538613936306530316238313532363030343031363034303531383039313033393066643562383035343638666666666666666666666666666666666666313931363630303136303031363034303162303338333136313736303031363034303162313738313535363131666666383438343631333632613536356238303534363066663630343031623139313638313535363034303531363030313630303136303430316230333833313638313532376663376635303562326633373161653231373565653439313366343439396531663236333361376235393336333231656564316364616562363131353138316432393036303230303136303430353138303931303339306131353035303530353035363562363034303531356639303566383035313630323036313564626538333339383135313931353239303766653932353436643639383935306464643338393130643265313565643164393233636430613762336464653965326136613366333830353635353539636230383930363132303866393038363930383639303631353939343536356239303831353236303230303136303430353138303931303339303230353439313530353039323931353035303536356235663834383135323566383035313630323036313564356538333339383135313931353236303230353236303430383038323230383135313630653038313031393039323532383035343566383035313630323036313564396538333339383135313931353239333932393139303832393036306666313636303033383131313135363132306634353736313230663436313466636235363562363030333831313131353631323130353537363132313035363134666362353635623831353238313534363130313030393030343630303136303031363061303162303331363630323038323031353236303031383230313534363034303830383330313931393039313532363030323930393230313534363030313630303136303430316230333830383231363630363038343031353236303031363034303162383230343831313636303830383430313532363030313630383031623832303438313136363061303834303135323630303136306330316239303931303431363630633039303931303135323831303135313930393135303566363132313762383236313165316535363562393035303630303238333531363030333831313131353631323139323537363132313932363134666362353635623134363132316233353738323531363034303531363333623064353430643630653231623831353236313036373239313930363030343031363135383736353635623630323038333031353136303031363030313630613031623033313633333134363132323536353735663832383135323630303438353031363032303532363034303930323035343333363030313630303136306130316230333930393131363033363132323530353735663832383135323630303438353031363032303532363034303930323035343630613038323031353136313232313739313630303136306230316239303034363030313630303136303430316230333136393036313537383135363562363030313630303136303430316230333136343231303135363132323462353736303430353136336662366365363366363065303162383135323630303136303031363034303162303334323136363030343832303135323630323430313631303637323536356236313232353635363562333336313065363535363562363030323831353136303035383131313135363132323662353736313232366236313466636235363562303336313233366335373836313536313232383235373631323238303832383736313235323435363562353035623566383838313532363030353835303136303230353236303430393032303830353436306666313931363630303331373930353536303630383330313531363038303832303135313631323262623931383439313631323262363931393036313537613135363562363133363434353635623530356638393831353236303035383630313630323035323630343038313230363030323031383035343630303136303031363034303162303339303933313636303031363063303162303236303031363030313630633031623033393039333136393239303932313739303931353536313232666338343631333831623536356239303530383538303135363132333039353735303830313535623135363132333261353736303430353136333130333663663931363065313162383135323630303438313031386139303532363032343031363130363732353635623566383938313532363030363836303136303230353236303430383038323230383339303535353138343931386239313766333636643333366330616233383064633739396630393561366638326132363332363538356335323930396363363938623039626134353430373039656435373931393061333530363132336339353635623630303438313531363030353831313131353631323338313537363132333831363134666362353635623033363132336164353736313233386638333631333831623536356235663839383135323630303638363031363032303532363034303930323035353631323361383838363133313662353635623631323363393536356238303531363034303531363331373063633933333630653231623831353236313036373239313930363030343031363135316235353635623530353035303530353035303530353035363562363034303830353136303630383038323031383335323566383038333532363032303833303135323931383130313931393039313532363034303531363330366638323533353630653431623831353236336666666666666666383331363630303438323031353235663930383139303630303536303031363039393162303139303633366638323533353039303630323430313566363034303531383038333033383138363561666131353830313536313234333735373364356638303365336435666664356235303530353035303630343035313364356638323365363031663364393038313031363031663139313638323031363034303532363132343565393139303831303139303631353961333536356239313530393135303830363132343830353736303430353136333662326631396539363065303162383135323630303430313630343035313830393130333930666435623831353131353631323461363537383135313630343035313633366261353839613536306530316238313532363030343831303139313930393135323630323430313631303637323536356236303230383230313531363030313630303136306130316230333136313536313234653235373630323038323031353136303430353136323464653735643630653331623831353236303031363030313630613031623033393039313136363030343832303135323630323430313631303637323536356235303932393135303530353635623566393038313532376634333137373133663765636264646464346263393965393564393033616465646161383833623265376332353531363130626431336532633765343733643034363032303532363034303930323035343630303136303031363061303162303331363135313539303536356236303430353136333036663832353335363065343162383135323633666666666666666638323136363030343832303135323566393038313930383139303630303536303031363039393162303139303633366638323533353039303630323430313566363034303531383038333033383138363561666131353830313536313235366635373364356638303365336435666664356235303530353035303630343035313364356638323365363031663364393038313031363031663139313638323031363034303532363132353936393139303831303139303631353961333536356239313530393135303830363132356238353736303430353136333662326631396539363065303162383135323630303430313630343035313830393130333930666435623630303536303031363039393162303136303031363030313630613031623033313636333432313363663738363034303531383136336666666666666666313636306530316238313532363030343031363032303630343035313830383330333831383635616661313538303135363132356662353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631323631663931393036313533303135363562383235313134363132363435353738313531363034303531363336626135383961353630653031623831353236303034383130313931393039313532363032343031363130363732353635623630323038323031353136303031363030313630613031623033313631353631323638313537363032303832303135313630343035313632346465373564363065333162383135323630303136303031363061303162303339303931313636303034383230313532363032343031363130363732353635623566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f36333038386332343633383536303430303135313630343035313832363366666666666666663136363065303162383135323630303430313631323662653931393036313532656635363562363034303830353138303833303338313836356166343135383031353631323664383537336435663830336533643566666435623530353035303530363034303531336436303166313936303166383230313136383230313830363034303532353038313031393036313236666339313930363135613333353635623931353039313530383138373134363132373233353736303430353136333038393933386233363065313162383135323630303438313031383839303532363032343031363130363732353635623566383738313532376634333137373133663765636264646464346263393965393564393033616465646161383833623265376332353531363130626431336532633765343733643034363032303532363034303930323036303031303135343566383035313630323036313564396538333339383135313931353239303630303136303031363034303162303339303831313639303833313631313135363132376530353735663838383135323630303438323031363032303930383135323630343039313832393032303630303130313830353436376666666666666666666666666666666631393136363030313630303136303430316230333836313639303831313739303931353539313531393138323532383939313766656334343134386538666632373166326430626163656631313432313534616261636230616262336132396562336562353065326361393765383664303433353931303136303430353138303931303339306132363132376666353635623566383838313532363030343832303136303230353236303430393032303630303130313534363030313630303136303430316230333136393135303562353039363935353035303530353035303530353635623566383035313630323036313564646538333339383135313931353238303534363030313139303136313238336235373630343035313633336565356165623536306530316238313532363030343031363034303531383039313033393066643562363030323930353535363562356636303530363132383465383438343631353761313536356236313238353839313930363135376666353635623630303136303031363034303162303331363631323836633835363036343631353766663536356236303031363030313630343031623033313631303135363132383832353735303566363132383836353635623530363030313562393339323530353035303536356235663566383035313630323036313564376538333339383135313931353238303534363034303531363334306331306631393630653031623831353236303031363030313630613031623033383638313136363030343833303135323630323438323031383639303532393239333530393131363930363334306331306631393930363034343031356636303430353138303833303338313566383738303362313538303135363132386538353735663830666435623530356166313135383031353631323866613537336435663830336533643566666435623530353035303530353035303530353635623630303135663830353136303230363135646465383333393831353139313532353535363562356638303531363032303631356439653833333938313531393135323566363132393265383636313339653435363562393035303631323933393836363132346539353635623631323934343537353035303631303737353536356235663836383135323630303438333031363032303532363034303930323035343630303136303031363061303162303331363333313436313239363935373333363130653635353635623566383638313532363030343833303136303230353236303430393032303534363061303832303135313631323939383931363030313630623031623930303436303031363030313630343031623033313639303631353738313536356236303031363030313630343031623033313638313630633030313531363030313630303136303430316230333136313031353631323964663537363063303831303135313630343035313633666236636536336636306530316238313532363030313630303136303430316230333930393131363630303438323031353236303234303136313036373235363562356638353135363132396637353736313239663038373836363132353234353635623930353036313261313535363562353035663836383135323630303438333031363032303532363034303930323036303031303135343630303136303031363034303162303331363562356638373831353236303034383430313630323035323630343038313230363030313031353436313261343239303630303136303430316239303034363030313630303136303430316230333136383336313537613135363562356638393831353236303034383630313630323035323630343038313230363030313031353439313932353036303031363038303162393039313034363030313630303136303430316230333136393038313930303336313261373835373530363061303833303135313562383538303135363132613931353735303631326138663832383238363630633030313531363132383431353635623135356231353631326162323537363034303531363335626666363833663630653131623831353236303034383130313861393035323630323430313631303637323536356236303033383530313534363034303835303135313566393136303031363030313630613031623033313639303633373738633036623539303631326164373930363131353330353635623834383538393630633030313531383835663830363034303531383836336666666666666666313636306530316238313532363030343031363132623033393739363935393439333932393139303631353763313536356236303230363034303531383038333033383138363561666131353830313536313262316535373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363132623432393139303631353330313536356239303530383038363630303730313566386338313532363032303031393038313532363032303031356632303566383238323534363132623636393139303631353539363536356239303931353535303530353035303530353035303530353035303530353035363562356636313262383036313464623035363562356638303531363032303631356462653833333938313531393135323566383037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f36333265343363656235363132626233383836313233643335363562363034303031353136303430353138323633666666666666666631363630653031623831353236303034303136313262643339313930363135326566353635623630343038303531383038333033383138363561663431353830313536313262656435373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363132633131393139303631353865373536356239313530393135303830313536313263333835373630343035313633326430373133353336306530316238313532383131353135363030343832303135323630323430313631303637323536356235663832383135323630303738343031363032303532363034303830383232303831353136306530383130313930393235323830353438323930363066663136363030353831313131353631326336393537363132633639363134666362353635623630303538313131313536313263376135373631326337613631346663623536356238313532363032303031363030313832303138303534363132633865393036313531376435363562383036303166303136303230383039313034303236303230303136303430353139303831303136303430353238303932393139303831383135323630323030313832383035343631326362613930363135313764353635623830313536313264303535373830363031663130363132636463353736313031303038303833353430343032383335323931363032303031393136313264303535363562383230313931393035663532363032303566323039303562383135343831353239303630303130313930363032303031383038333131363132636538353738323930303336303166313638323031393135623530353035303931383335323530353036303032383230313534363030313630303136303430316230333830383231363630323038343031353236303031363034303162383230343831313636303430383430313532363030313630383031623832303438313136363036303834303135323630303136306330316239303931303438313136363038303833303135323630303339323833303135343136363061303930393130313532393039313530383135313630303538313131313536313264373035373631326437303631346663623536356231343135383031353631326439313537353036303031383135313630303538313131313536313264386535373631326438653631346663623536356231343135356231353631326462323537383035313630343035313633313730636339333336306532316238313532363130363732393139303630303430313631353162353536356236303033383135313630303538313131313536313264633735373631326463373631346663623536356230333631326464353537363030343831353236313264646135363562363030353831353235623833363030383031383136303230303135313630343035313631326466303931393036313534373235363562393038313532363034303830353136303230393238313930303338333031393032303566393038313930353538353831353236303037383730313930393235323930323038313531383135343833393239313930383239303630666631393136363030313833363030353831313131353631326533343537363132653334363134666362353635623032313739303535353036303230383230313531363030313832303139303631326534643930383236313534633735363562353036303430383230313531363030323832303138303534363036303835303135313630383038363031353136306130383730313531363030313630303136303430316230333935383631363630303136303031363038303162303331393930393431363933393039333137363030313630343031623932383631363932393039323032393139303931313736303031363030313630383031623033313636303031363038303162393138353136393139303931303236303031363030313630633031623033313631373630303136306330316239313834313639313930393130323137393035353630633039303932303135313630303339303931303138303534363766666666666666666666666666666666313931363931393039323136313739303535383035313630303538313131313536313265663335373631326566333631346663623536356236303430353138343930376631633038653539363536663161313864633264613736383236636463353238303563343365383937613137633530666165666238616233633135323663633136393035663930613339313936393139353530393039333530353035303530353635623566383035313630323036313564376538333339383135313931353235343631313130663930363030313630303136306130316230333136383338333631336363383536356237663433313737313366376563626464646434626339396539356439303361646564616138383362326537633235353136313062643133653263376534373364303235343566393035663830353136303230363135643965383333393831353139313532393036303031363034303162393030343631666666663930383131363930383631363130383036313266613135373530363132373130363166666666383631363131356231353631326663353537363034303531363335663132653663333630653131623831353236316666666638363136363030343832303135323630323430313631303637323536356236303032383130313534363030313630303136303430316230333930383131363930383531363130313536313330303135373630343035313632303261303664363065313162383135323630303136303031363034303162303338353136363030343832303135323630323430313631303637323536356238303534383331303830363133303133353735303830363030313031353438333131356231353631333033343537363034303531363332323264313634333630653231623831353236303034383130313834393035323630323430313631303637323536356235663631333033653834363133643237353635623930353035663631333034613832363131313133353635623930353035663631333035373839383336313364346135363562393035303630343035313830363063303031363034303532383036313330366233333930353635623630303136303031363061303162303339303831313638323532363166666666383038633136363032303830383530313931393039313532363030313630303136303430316230333830386431363630343038303837303139313930393135323566363036303830383830313832393035323630383038303839303138333930353236306130393838393031383339303532386138333532363030343930396430313835353239303832393032303838353138313534393538613031353139333861303135313835313636303031363062303162303236376666666666666666666666666666666636306230316231393934393039373136363030313630613031623032363030313630303136306230316230333139393039363136393731363936393039363137393339303933313731363932393039323137383335353834303135313630303139303932303138303534393838353031353139343930393330313531383131363630303136303830316230323637666666666666666666666666666666663630383031623139393438323136363030313630343031623032363030313630303136303830316230333139393039393136393239303931313639313930393131373936393039363137393139303931313639343930393431373930393335353530393039313530353039343933353035303530353035363562356638313831353235663830353136303230363135643565383333393831353139313532363032303532363034303830383232303831353136306530383130313930393235323830353435663830353136303230363135643965383333393831353139313532393339323931393038323930363066663136363030333831313131353631333162383537363133316238363134666362353635623630303338313131313536313331633935373631333163393631346663623536356238313532383135343631303130303930303436303031363030313630613031623033313636303230383038333031393139303931353236303031383038343031353436303430383038353031393139303931353236303032393438353031353436303031363030313630343031623033383038323136363036303837303135323630303136303430316238323034383131363630383038373031353236303031363038303162383230343831313636306130383730313532363030313630633031623930393130343136363063303930393430313933393039333532383438333031353135663839383135323630303538393031383435323834383132303830353436303031363030313630613831623033313931363831353539323833303138313930353539313930393430313831393035353630303638373031393039313532393038313230383035343930383239303535393239333530393039313930383038323135363133326637353735663834383135323630303438373031363032303532363034303930323035343631323731303930363133326139393036303031363061303162393030343631666666663136383536313538643035363562363133326233393139303631353836333536356239313530383138363630303730313566383638313532363032303031393038313532363032303031356632303566383238323534363133326437393139303631353539363536356239303931353535303631333265373930353038323834363135613536353635623930353036313332663738353630323030313531383236313238386435363562363133333063383536303230303135313631313230303837363036303031353136313135333035363562363034303830353138323831353236303230383130313834393035323835393138393931376638656365636635313030373063333230643961353533323366666162653335306532393461653530356663306335303964633537333664613666356363393933393130313630343035313830393130333930613335303530353035303530353035303536356235663566383035313630323036313564396538333339383135313931353238313631333336653631303261643835363133643237353635623930353035663631333337613837363131653165353635623930353036313333383538373631323465393536356236313333613535373630343035313633333065666139386236306530316238313532363030343831303138383930353236303234303136313036373235363562363030323831353136303035383131313135363133336261353736313333626136313466636235363562313436313333646235373830353136303430353136333137306363393333363065323162383135323631303637323931393036303034303136313531623535363562356638323832363038303031353136313333656339313930363135373831353635623930353038333630303230313630306139303534393036313031303030613930303436303031363030313630343031623033313638323630343030313531363133343135393139303631353766663536356236303031363030313630343031623033313638313630303136303031363034303162303331363131313536313334353235373630343035313633366435316665303536306531316238313532363030313630303136303430316230333832313636303034383230313532363032343031363130363732353635623566383036313334356538613834363133363434353635623931353039313530356638613833363034303531363032303031363133343863393239313930393138323532363063303162363030313630303136306330316230333139313636303230383230313532363032383031393035363562363034303830353136303166313938313834303330313831353238323832353238303531363032303930393130313230363065303833303139303931353239313530383036303031383135323630303136303031363061303162303338633136363032303830383330313931393039313532363034303830383330313866393035323630303136303031363034303162303338303862313636303630383530313532356636303830383530313831393035323930383831363630613038353031353236306330393039333031383339303532383438333532363030353862303139303931353239303230383135313831353438323930363066663139313636303031383336303033383131313135363133353166353736313335316636313466636235363562303231373930353535303630323038323831303135313832353436313031303036303031363061383162303331393136363130313030363030313630303136306130316230333932383331363032313738333535363034303830383530313531363030313835303135353630363038303836303135313630303239303935303138303534363038303830383930313531363061303861303135313630633039303961303135313630303136303031363034303162303339393861313636303031363030313630383031623033313939303934313639333930393331373630303136303430316239313861313639313930393130323137363030313630303136303830316230333136363030313630383031623939383931363939393039393032363030313630303136306330316230333136393839303938313736303031363063303162393138383136393139303931303231373930353538313531383938363136383135323861383631363934383130313934393039343532393338623136393038333031353239313831303138353930353239303863313639313864393138343931376662303032346232363362633361306237323861366564656135306136396566613834313138396638643332656538616639643163326231613161323233343236393130313630343035313830393130333930613439613939353035303530353035303530353035303530353035363562363133363332363134333332353635623631333633623832363134333764353635623631313130663831363134336562353635623566383238313532356638303531363032303631356466653833333938313531393135323630323035323630343038313230363030323031353438313930356638303531363032303631356462653833333938313531393135323930363030313630383031623930303436303031363030313630343031623033313636313336386338353832363134343533353635623566363133363936383736313436326435363562356638383831353236303037383530313630323035323630343038303832323036303032303138303534363766666666666666666666666666666666363038303162313931363630303136303830316236303031363030313630343031623033386338313136393138323032393239303932313739303932353539313531363334326132653062353630653131623831353236303034383130313863393035323931383431363630323438333031353236303434383230313532393139323530393036303035363030313630393931623031393036336565356234386562393037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f39303633383534356331366139303630363430313566363034303531383038333033383138363561663431353830313536313337336635373364356638303365336435666664356235303530353035303630343035313364356638323365363031663364393038313031363031663139313638323031363034303532363133373636393139303831303139303631353262653536356236303430353138323633666666666666666631363630653031623831353236303034303136313337383239313930363135326566353635623630323036303430353138303833303338313566383735616631313538303135363133373965353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631333763323931393036313533303135363562363034303830353136303031363030313630343031623033386138313136383235323630323038323031383439303532383235313933393435303835313639323862393237663037646535666633356136373461383030356536363166333333336339303763613633333334363238303837363264313964633762336162623161386331646639323832393030333031393061333930393435303932353035303530356239323530393239303530353635623566383035663830353136303230363135643965383333393831353139313532393035303566363133383362383436303430303135313631316531653536356239303530356636303033383235313630303538313131313536313338353335373631333835333631346663623536356231343830363133383731353735303630303438323531363030353831313131353631333836663537363133383666363134666362353635623134356231353631333838313537353036306330383130313531363133386265353635623630303238323531363030353831313131353631333839363537363133383936363134666362353635623033363133386132353735303432363133386265353635623831353136303430353136333137306363393333363065323162383135323631303637323931393036303034303136313531623535363562383436303830303135313630303136303031363034303162303331363831363030313630303136303430316230333136313136313338653535373530356639343933353035303530353035363562363034303830383630313531356639303831353236303034383530313630323035323230363030313031353436306130383330313531363133393134393136303031363030313630343031623033313639303833363132383431353635623631333932323537353035663934393335303530353035303536356236303033383330313534363036303836303135313630303136303031363061303162303339303931313639303633373738633036623539303631333934373930363131353330353635623630613038353031353136303830383930313531363034303830386230313531356639303831353236303034383038623031363032303532383238323230363030313031353439323531363030313630303136306530316230333139363065303839393031623136383135323631333939633936393539343933386139333630303136303031363034303162303339303931313639323930393138323931303136313537633135363562363032303630343035313830383330333831383635616661313538303135363133396237353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631333964623931393036313533303135363562393539343530353035303530353035363562363133396563363134646230353635623566383238313532356638303531363032303631356466653833333938313531393135323630323035323630343038303832323038313531363065303831303139303932353238303534356638303531363032303631356462653833333938313531393135323933393239313930383239303630666631363630303538313131313536313361333935373631336133393631346663623536356236303035383131313135363133613461353736313361346136313466636235363562383135323630323030313630303138323031383035343631336135653930363135313764353635623830363031663031363032303830393130343032363032303031363034303531393038313031363034303532383039323931393038313831353236303230303138323830353436313361386139303631353137643536356238303135363133616435353738303630316631303631336161633537363130313030383038333534303430323833353239313630323030313931363133616435353635623832303139313930356635323630323035663230393035623831353438313532393036303031303139303630323030313830383331313631336162383537383239303033363031663136383230313931356235303530353039313833353235303530363030323832383130313534363030313630303136303430316230333830383231363630323038353031353236303031363034303162383230343831313636303430383530313532363030313630383031623832303438313136363036303835303135323630303136306330316239303931303438313136363038303834303135323630303339303933303135343930393231363630613039303931303135323930393135303831353136303035383131313135363133623433353736313362343336313466636235363562313436313362373635373566383438313532363030373833303136303230353236303430393038313930323035343930353136333137306363393333363065323162383135323631303637323931363066663136393036303034303136313531623535363562363030333831353234323630303136303031363034303162303331363630633038323031353235663834383135323630303738333031363032303532363034303930323038313531383135343833393239313930383239303630666631393136363030313833363030353831313131353631336262613537363133626261363134666362353635623032313739303535353036303230383230313531363030313832303139303631336264333930383236313534633735363562353036303430383230313531363030323832303138303534363036303835303135313630383038363031353136306130383730313531363030313630303136303430316230333935383631363630303136303031363038303162303331393930393431363933393039333137363030313630343031623932383631363932393039323032393139303931313736303031363030313630383031623033313636303031363038303162393138353136393139303931303236303031363030313630633031623033313631373630303136306330316239313834313639313930393130323137393035353630633039303932303135313630303339303931303138303534363766666666666666666666666666666666313931363931393039323136313739303535356636313363373138353832363133363434353635623630383038343031353136303430383035313630303136303031363034303162303339303932313638323532343236303230383330313532393139333530383339323530383739313766313364353833393463663236396434386263663932373935396132396135666665653763393932346461666666383932376563646633633438666661376336373931303136303430353138303931303339306133353039333932353035303530353635623630343035313630303136303031363061303162303338333831313636303234383330313532363034343832303138333930353236313034653639313835393138323136393036336139303539636262393036303634303135623630343035313630323038313833303330333831353239303630343035323931353036306530316236303230383230313830353136303031363030313630653031623033383338313833313631373833353235303530353035303631343661323536356235663631313132333832356638303531363032303631356437653833333938313531393135323534363030313630303136306130316230333136393036313437303335363562376665393235343664363938393530646464333839313064326531356564316439323363643061376233646465396532613661336633383035363535353963623039353435663930363066663136363133643865353736303430353136333766616238316535363065303162383135323630303430313630343035313830393130333930666435623566383035313630323036313564626538333339383135313931353234323631336461643630363038363031363034303837303136313466393035363562363030313630303136303430316230333136313131353830363133646537353735303631336463623632303261333030343236313535393635363562363133646462363036303836303136303430383730313631346639303536356236303031363030313630343031623033313631303135356231353631336532313537363133646663363036303835303136303430383630313631346639303536356236303430353136333538373964613133363065313162383135323630303136303031363034303162303339303931313636303034383230313532363032343031363130363732353635623631336533363631336533313630363038363031383636313561363935363562363134383563353635623631336534363631336533313630383038363031383636313561363935363562363033303631336535353630323038363031383636313561376435363562393035303134363133653837353736313365363936303230383530313835363135613764353635623630343035313633323634373562326636306531316238313532363130363732393235303630303430313930383135323630323030313930353635623631336539313834383036313561376435363562393035303566303336313365626535373631336561333834383036313561376435363562363034303531363333653038613132353630653131623831353236303034303136313036373239323931393036313561626635363562356636303038383230313631336563643836383036313561376435363562363034303531363133656462393239313930363135393934353635623930383135323630323030313630343035313830393130333930323035343134363133663134353736313365663938343830363135613764353635623630343035313633613431663737326636306530316238313532363030343031363130363732393239313930363135616266353635623631336631653833356636313434353335363562363034303830353136306530383130313930393135323831353438313532356639303831393037335f5f2466643063313437623430333165656636303739623034393863626166613836356630245f5f39303633653034376232383339303630323038313031363133663562386138303631356137643536356238303830363031663031363032303830393130343032363032303031363034303531393038313031363034303532383039333932393139303831383135323630323030313833383338303832383433373566393230313931393039313532353035303530393038323532353036303230393038313031393036313366613339303862303138623631356137643536356238303830363031663031363032303830393130343032363032303031363034303531393038313031363034303532383039333932393139303831383135323630323030313833383338303832383433373566393230313931393039313532353035303530393038323532353036303230303136313366656336303630386230313630343038633031363134663930353635623630303136303031363034303162303331363831353236303230303136313430303736303630386230313862363135613639353635623631343031303930363135616432353635623831353236303230303136313430323236303830386230313862363135613639353635623631343032623930363135616432353635623831353236303230303138383630303136303031363034303162303331363831353235303630343035313832363366666666666666663136363065303162383135323630303430313631343035393931393036313562663435363562356636303430353138303833303338313836356166343135383031353631343037333537336435663830336533643566666435623530353035303530363034303531336435663832336536303166336439303831303136303166313931363832303136303430353236313430396139313930383130313930363135636162353635623566383238313532363030363836303136303230353236303430393032303931393335303931353036313430623838323832363135346337353635623530383136303038383430313631343063383838383036313561376435363562363034303531363134306436393239313930363135393934353635623930383135323630343035313930383139303033363032303031383132303931393039313535363365653562343865623630653031623831353235663930363030353630303136303939316230313930363365653562343865623930363134313132393038353930363030343031363135326566353635623630323036303430353138303833303338313566383735616631313538303135363134313265353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631343135323931393036313533303135363562363034303830353136306530383130313930393135323930393135303830363030313831353236303230303136313431373238393830363135613764353635623830383036303166303136303230383039313034303236303230303136303430353139303831303136303430353238303933393239313930383138313532363032303031383338333830383238343337356639323031383239303532353039333835353235303530353036303031363030313630343031623033383931363630323038303834303138323930353236303430383038353031383439303532363036303835303139323930393235323630383038343031383339303532363061303930393330313832393035323836383235323630303738383031393039323532323038313531383135343832393036306666313931363630303138333630303538313131313536313432303135373631343230313631346663623536356230323137393035353530363032303832303135313630303138323031393036313432316139303832363135346337353635623530363034303832303135313630303238323031383035343630363038353031353136303830383630313531363061303837303135313630303136303031363034303162303339353836313636303031363030313630383031623033313939303934313639333930393331373630303136303430316239323836313639323930393230323931393039313137363030313630303136303830316230333136363030313630383031623931383531363931393039313032363030313630303136306330316230333136313736303031363063303162393138343136393139303931303231373930353536306330393039323031353136303033393039313031383035343637666666666666666666666666666666663139313639313930393231363137393035353830363134326238383838303631356137643536356236303430353136313432633639323931393036313539393435363562363034303531383039313033393032303834376662373732393765336265666336393162666338363461383165323431663833653265663732326236653762656361613265636563323530633664353262343330383938623630343030313630323038313031393036313433303439313930363134663930353635623630343038303531363030313630303136303430316230333933383431363831353239323930393131363630323038333031353230313630343035313830393130333930613435303930393539343530353035303530353035363562376666306335376531363834306466303430663135303838646332663831666533393163333932336265633733653233613936363265666339633232396336613030353436303031363034303162393030343630666631363631343337623537363034303531363331616663643739663630653331623831353236303034303136303430353138303931303339306664356235363562363134333835363134333332353635623631343338653831363134396335353635623631343339363631343964653536356236313132323036303630383230313335363038303833303133353631343362333630633038353031363061303836303136313466393035363562363134336333363065303836303136306330383730313631356365653536356236313433643436313031303038373031363065303838303136313564303735363562363134336536363130313230383830313631303130303839303136313533313835363562363134396565353635623631343366333631343333323536356235663830353136303230363135643765383333393831353139313532363030313630303136306130316230333832313636313434333335373630343035313633373333303638303336306530316238313532363030313630303136306130316230333833313636303034383230313532363032343031363130363732353635623830353436303031363030313630613031623033313931363630303136303031363061303162303339323930393231363931393039313137393035353536356235663830353136303230363135646265383333393831353139313532356636303031363030313630343031623033383038343136393038353136313131353631343438373537363134343830383338353631353761313536356239303530363134343934353635623631343439313834383436313537613135363562393035303562363034303830353136303830383130313832353236303032383430313534383038323532363030333835303135343630323038333031353236303034383530313534393238323031393239303932353236303035383430313534363030313630303136303430316230333136363036303832303135323432393131353830363134346636353735303630303138343031353438313531363134346632393136303031363030313630343031623033313639303631353539363536356238323130313535623135363134353163353736303031363030313630343031623033383331363630363038323031353238313831353236303430383130313531363032303832303135323631343533623536356238323831363036303031383138313531363134353265393139303631353738313536356236303031363030313630343031623033313639303532353035623630363038313031353136313435346239303630363436313537666635363562363032303832303135313630303138363031353436303031363030313630343031623033393239303932313639313631343537363931393036303031363034303162393030343630666631363631353864303536356231303135363134356136353736303630383130313531363034303531363364666165383830313630653031623831353236303031363030313630343031623033393039313136363030343832303135323630323430313631303637323536356238353630303136303031363034303162303331363831363034303031383138313531363134356331393139303631353539363536356239303532353036303430383130313830353136303031363030313630343031623033383731363931393036313435653139303833393036313561353635363562393035323530383035313630303238353031353536303230383130313531363030333835303135353630343038313031353136303034383530313535363036303031353136303035393039333031383035343637666666666666666666666666666666663139313636303031363030313630343031623033393039343136393339303933313739303932353535303530353035303536356235663831383135323566383035313630323036313564666538333339383135313931353236303230353236303430383132303630303230313830353435663830353136303230363135646265383333393831353139313532393139303630303839303631343637373930363030313630343031623930303436303031363030313630343031623033313636313564323735363562393139303631303130303061383135343831363030313630303136303430316230333032313931363930383336303031363030313630343031623033313630323137393035353931353035303931393035303536356235663631343662363630303136303031363061303162303338343136383336313462316435363562393035303830353135663134313538303135363134366461353735303830383036303230303139303531383130313930363134366438393139303631356434323536356231353562313536313034653635373630343035313633353237346166653736306530316238313532363030313630303136306130316230333834313636303034383230313532363032343031363130363732353635623630343035313633373061303832333136306530316238313532333036303034383230313532356639303831393036303031363030313630613031623033383531363930363337306130383233313930363032343031363032303630343035313830383330333831383635616661313538303135363134373439353733643566383033653364356666643562353035303530353036303430353133643630316631393630316638323031313638323031383036303430353235303831303139303631343736643931393036313533303135363562393035303631343738343630303136303031363061303162303338353136333333303836363134623261353635623630343035313633373061303832333136306530316238313532333036303034383230313532356639303630303136303031363061303162303338363136393036333730613038323331393036303234303136303230363034303531383038333033383138363561666131353830313536313437633835373364356638303365336435666664356235303530353035303630343035313364363031663139363031663832303131363832303138303630343035323530383130313930363134376563393139303631353330313536356239303530383138313131363134383532353736303430353136323436316263643630653531623831353236303230363030343832303135323630326336303234383230313532376635333631363636353435353234333332333035343732363136653733363636353732343637323666366433613230363236313663363136653633363532303665363034343832303135323662316264643038316135623938646339393538356364393539363061323162363036343832303135323630383430313631303637323536356236313339646238323832363135613536353635623631343836393630323038323031383236313465656235363562363366666666666666663136313538303135363134383839353735303631343838343630323038323031383236313533333335363562313531353930353035623135363134386430353736313438396236303230383230313832363134656562353635623631343861383630323038333031383336313533333335363562363034303531363363303861306631643630653031623831353236336666666666666666393039333136363030343834303135323630323438333031353235303630343430313631303637323536356236313438646436303230383230313832363135333333353635623930353036313438656336303230383330313833363134656562353635623633666666666666666631363131313536313439303535373631343839623630323038323031383236313465656235363562363030313562363134393135363032303833303138333631353333333536356239303530383131303135363131313066353736313439326236303230383330313833363135333333353635623631343933363630303138343631356135363536356238313831313036313439343535373631343934353631353337383536356239303530363032303032303136303230383130313930363134393561393139303631353331383536356236303031363030313630613031623033313636313439373036303230383430313834363135333333353635623833383138313130363134393830353736313439383036313533373835363562393035303630323030323031363032303831303139303631343939353931393036313533313835363562363030313630303136306130316230333136313031353631343962643537363034303531363330646263386435663630653331623831353236303034303136303430353138303931303339306664356236303031303136313439303835363562363134396364363134333332353635623631343964353631346236333536356236313132323038313631346236623536356236313439653636313433333235363562363134333762363134633533353635623631343966363631343333323536356235663830353136303230363135643965383333393831353139313532363166666666383431363135383036313461316135373530363132373130363166666666383531363131356231353631346133653537363034303531363335663132653663333630653131623831353236316666666638353136363030343832303135323630323430313631303637323536356238353837313131353631346136323537363034303531363332323264313634333630653231623831353236303034383130313838393035323630323430313631303637323536356236306666383331363135383036313461373535373530363030613630666638343136313135623135363134613938353736303430353136333137306462333539363065333162383135323630666638343136363030343832303135323630323430313631303637323536356239353836353536303031383630313934393039343535363030323835303138303534363030313630303136303430316230333934393039343136363966666666666666666666666666666666666666663139393039343136393339303933313736303031363034303162363166666666393339303933313639323930393230323931393039313137363766666666666666666666666666666666363035303162313931363630666639313930393131363630303136303530316230323137393035353630303339303931303138303534363030313630303136306130316230333139313636303031363030313630613031623033393039323136393139303931313739303535353635623630363036313238383638333833356636313463356235363562363034303531363030313630303136306130316230333834383131363630323438333031353238333831313636303434383330313532363036343832303138333930353236313037373539313836393138323136393036333233623837326464393036303834303136313363663535363562363134333762363134333332353635623631346237333631343333323536356238303335356638303531363032303631356462653833333938313531393135323930383135353630313436313462393836303630383430313630343038353031363135643037353635623630666631363131383036313462623735373530363134626232363036303833303136303430383430313631356430373536356236306666313631353562313536313462656235373631346263633630363038333031363034303834303136313564303735363562363034303531363334613539626266663630653131623831353236306666393039313136363030343832303135323630323430313631303637323536356236313462666236303630383330313630343038343031363135643037353635623630303138323031383035343630666639323930393231363630303136303430316230323630666636303430316231393930393231363931393039313137393035353631346332633630343038333031363032303834303136313466393035363562363030313931393039313031383035343637666666666666666666666666666666663139313636303031363030313630343031623033393039323136393139303931313739303535353035363562363132393033363134333332353635623630363038313437313031353631346338303537363034303531363363643738363035393630653031623831353233303630303438323031353236303234303136313036373235363562356638303835363030313630303136306130316230333136383438363630343035313631346339623931393036313534373235363562356636303430353138303833303338313835383735616631393235303530353033643830356638313134363134636435353736303430353139313530363031663139363033663364303131363832303136303430353233643832353233643566363032303834303133653631346364613536356236303630393135303562353039313530393135303631346365613836383338333631346366343536356239363935353035303530353035303530353635623630363038323631346430393537363134643034383236313464353035363562363132383836353635623831353131353830313536313464323035373530363030313630303136306130316230333834313633623135356231353631346434393537363034303531363339393936623331353630653031623831353236303031363030313630613031623033383531363630303438323031353236303234303136313036373235363562353038303631323838363536356238303531313536313464363035373830353138303832363032303031666435623630343035313633306131326635323136306531316238313532363030343031363034303531383039313033393066643562353038303534363134643835393036313531376435363562356638323535383036303166313036313464393435373530353035363562363031663031363032303930303439303566353236303230356632303930383130313930363131323230393139303631346465643536356236303430383035313630653038313031393039313532383035663831353236303630363032303832303138313930353235663630343038333031383139303532393038323031383139303532363038303832303138313930353236306130383230313831393035323630633039303931303135323930353635623562383038323131313536313465303135373566383135353630303130313631346465653536356235303930353635623830313531353831313436313132323035373566383066643562383033353633666666666666666638313136383131343631346532353537356638306664356239313930353035363562356638303566363036303834383630333132313536313465336335373566383066643562383333353932353036303230383430313335363134653465383136313465303535363562393135303631346535633630343038353031363134653132353635623930353039323530393235303932353635623566363032303832383430333132313536313465373535373566383066643562353033353931393035303536356235663830363034303833383530333132313536313465386435373566383066643562383233353630303136303031363034303162303338313131313536313465613235373566383066643562383330313630383038313836303331323135363134656233353735663830666435623931353036313465633136303230383430313631346531323536356239303530393235303932393035303536356235663830363034303833383530333132313536313465646235373566383066643562383233353931353036313465633136303230383430313631346531323536356235663630323038323834303331323135363134656662353735663830666435623631323838363832363134653132353635623830333536316666666638313136383131343631346532353537356638306664356236303031363030313630343031623033383131363831313436313132323035373566383066643562356638303566383036303830383538373033313231353631346633633537356638306664356238343335363030313630303136303430316230333831313131353631346635313537356638306664356238353031363061303831383830333132313536313466363235373566383066643562393335303631346637303630323038363031363134663034353635623932353036303430383530313335363134663830383136313466313535363562393339363932393535303932393336303630303133353932353035303536356235663630323038323834303331323135363134666130353735663830666435623831333536313238383638313631346631353536356235663830363034303833383530333132313536313466626335373566383066643562353035303830333539323630323039303931303133353931353035363562363334653438376237313630653031623566353236303231363030343532363032343566666435623630303638313130363134666566353736313466656636313466636235363562393035323536356235663562383338313130313536313530306435373831383130313531383338323031353236303230303136313466663535363562353035303566393130313532353635623566383135313830383435323631353032633831363032303836303136303230383630313631346666333536356236303166303136303166313931363932393039323031363032303031393239313530353035363562363032303831353236313530353236303230383230313833353136313466646635363562356636303230383330313531363065303630343038343031353236313530366436313031303038343031383236313530313535363562393035303630343038343031353136303031363030313630343031623033383038323136363036303836303135323830363036303837303135313136363038303836303135323830363038303837303135313136363061303836303135323830363061303837303135313136363063303836303135323830363063303837303135313136363065303836303135323530353038303931353035303932393135303530353635623630303136303031363061303162303338313136383131343631313232303537356638306664356235663830383238343033363130313430383131323135363135306537353735663830666435623631303132303830383231323135363135306636353735663830666435623834393335303833303133353930353036313531303738313631353063303536356238303931353035303932353039323930353035363562356638303630323038333835303331323135363135313233353735663830666435623832333536303031363030313630343031623033383038323131313536313531333935373566383066643562383138353031393135303835363031663833303131323631353134633537356638306664356238313335383138313131313536313531356135373566383066643562383636303230383238353031303131313135363135313662353735663830666435623630323039323930393230313936393139353530393039333530353035303530353635623630303138313831316339303832313638303631353139313537363037663832313639313530356236303230383231303831303336313531616635373633346534383762373136306530316235663532363032323630303435323630323435666664356235303931393035303536356236303230383130313631313132333832383436313466646635363562363334653438376237313630653031623566353236303431363030343532363032343566666435623630343035313630363038313031363030313630303136303430316230333831313138323832313031373135363135316639353736313531663936313531633335363562363034303532393035363562363034303830353139303831303136303031363030313630343031623033383131313832383231303137313536313531663935373631353166393631353163333536356236303430353136303166383230313630316631393136383130313630303136303031363034303162303338313131383238323130313731353631353234393537363135323439363135316333353635623630343035323931393035303536356235663630303136303031363034303162303338323131313536313532363935373631353236393631353163333536356235303630316630313630316631393136363032303031393035363562356638323630316638333031313236313532383635373566383066643562383135313631353239393631353239343832363135323531353635623631353232313536356238313831353238343630323038333836303130313131313536313532616435373566383066643562363131323530383236303230383330313630323038373031363134666633353635623566363032303832383430333132313536313532636535373566383066643562383135313630303136303031363034303162303338313131313536313532653335373566383066643562363131323530383438323835303136313532373735363562363032303831353235663631323838363630323038333031383436313530313535363562356636303230383238343033313231353631353331313537356638306664356235303531393139303530353635623566363032303832383430333132313536313533323835373566383066643562383133353631323838363831363135306330353635623566383038333335363031653139383433363033303138313132363135333438353735663830666435623833303138303335393135303630303136303031363034303162303338323131313536313533363135373566383066643562363032303031393135303630303538313930316233363033383231333135363133383134353735663830666435623633346534383762373136306530316235663532363033323630303435323630323435666664356235663832333536303565313938333336303330313831313236313533613035373566383066643562393139303931303139323931353035303536356235663832363031663833303131323631353362393537356638306664356238313335363135336337363135323934383236313532353135363562383138313532383436303230383338363031303131313135363135336462353735663830666435623831363032303835303136303230383330313337356639313831303136303230303139313930393135323933393235303530353035363562356636303630383233363033313231353631353430373537356638306664356236313534306636313531643735363562383233353630303136303031363034303162303338303832313131353631353432353537356638306664356236313534333133363833383730313631353361613536356238333532363032303835303133353931353038303832313131353631353434363537356638306664356235303631353435333336383238363031363135336161353635623630323038333031353235303630343038333031333536313534363738313631346631353536356236303430383230313532393239313530353035363562356638323531363135336130383138343630323038373031363134666633353635623630316638323131313536313034653635373830356635323630323035663230363031663834303136303035316338313031363032303835313031353631353461383537353038303562363031663834303136303035316338323031393135303562383138313130313536313163646135373566383135353630303130313631353462343536356238313531363030313630303136303430316230333831313131353631353465303537363135346530363135316333353635623631353466343831363135346565383435343631353137643536356238343631353438333536356236303230383036303166383331313630303138313134363135353237353735663834313536313535313035373530383538333031353135623566313936303033383639303162316331393136363030313835393031623137383535353631313338373536356235663835383135323630323038313230363031663139383631363931356238323831313031353631353535353537383838363031353138323535393438343031393436303031393039313031393038343031363135353336353635623530383538323130313536313535373235373837383530313531356631393630303338383930316236306638313631633139313638313535356235303530353035303530363030313930383131623031393035353530353635623633346534383762373136306530316235663532363031313630303435323630323435666664356238303832303138303832313131353631313132333537363131313233363135353832353635623566363366666666666666663830383331363831383130333631353563313537363135356331363135353832353635623630303130313933393235303530353035363562356638303833333536303165313938343336303330313831313236313535653035373566383066643562383330313630323038313031393235303335393035303630303136303031363034303162303338313131313536313535666535373566383066643562383033363033383231333135363133383134353735663830666435623831383335323831383136303230383530313337353035663832383230313630323039303831303139313930393135323630316639303931303136303166313931363930393130313031393035363562356638333833383535323630323038303836303139353530383038353630303531623833303130313834356635623837383131303135363135366539353738343833303336303166313930313839353238313335333638383930303336303565313930313831313236313536373035373566383066643562383730313630363036313536376538323830363135356362353635623832383735323631353638653833383830313832383436313536306335363562393235303530353036313536396538363833303138333631353563623536356238363833303338383838303135323631353662303833383238343631353630633536356239323530353035303630343038303833303133353932353036313536633538333631346631353536356236303031363030313630343031623033393239303932313639343930393130313933393039333532393738333031393739303833303139303630303130313631353634623536356235303930393739363530353035303530353035303530353635623630323038313532383133353630323038323031353236303230383230313335363034303832303135323566363034303833303133353631353731613831363135306330353635623630303136303031363061303162303331363630363038333831303139313930393135323833303133353336383439303033363031653139303138313132363135373431353735663830666435623833303136303230383130313930333536303031363030313630343031623033383131313135363135373563353735663830666435623830363030353162333630333832313331353631353736643537356638306664356236303830383038353031353236313339646236306130383530313832383436313536333435363562363030313630303136303430316230333831383131363833383231363031393038303832313131353631323465323537363132346532363135353832353635623630303136303031363034303162303338323831313638323832313630333930383038323131313536313234653235373631323465323631353538323536356239363837353236303031363030313630343031623033393538363136363032303838303135323933383531363630343038373031353239313834313636303630383630313532393039323136363038303834303135323630613038333031393139303931353236306330383230313532363065303031393035363562363030313630303136303430316230333831383131363833383231363032383038323136393139303832383131343631353832323537363135383232363135353832353635623530353039323931353035303536356236333465343837623731363065303162356635323630313236303034353236303234356666643562356636303031363030313630343031623033383038343136383036313538353735373631353835373631353832613536356239323136393139303931303439323931353035303536356235663832363135383731353736313538373136313538326135363562353030343930353635623630323038313031363030343833313036313538386135373631353838613631346663623536356239313930353239303536356235663830356636303630383438363033313231353631353861323537356638306664356238333531393235303630323038343031353136313538623438313631346631353536356236303430383530313531393039323530363135386335383136313466313535363562383039313530353039323530393235303932353635623830383230323831313538323832303438343134313736313131323335373631313132333631353538323536356235663830363034303833383530333132313536313538663835373566383066643562383235313931353036303230383330313531363135313037383136313465303535363562356636303230383038333532356638343534363135393163383136313531376435363562383036303230383730313532363034303630303138303834313635663831313436313539336435373630303138313134363135393539353736313539383635363562363066663139383531363630343038613031353236303430383431353135363030353162386130313031393535303631353938363536356238393566353236303230356632303566356238353831313031353631353937643537383135343862383230313836303135323930383330313930383830313631353936323536356238613031363034303031393635303530356235303933393839373530353035303530353035303530353035363562383138333832333735663931303139303831353239313930353035363562356638303630343038333835303331323135363135396234353735663830666435623832353136303031363030313630343031623033383038323131313536313539636135373566383066643562393038343031393036303630383238373033313231353631353964643537356638306664356236313539653536313531643735363562383235313831353236303230383330313531363135396637383136313530633035363562363032303832303135323630343038333031353138323831313131353631356130643537356638306664356236313561313938383832383630313631353237373536356236303430383330313532353038303934353035303530353036303230383330313531363135313037383136313465303535363562356638303630343038333835303331323135363135613434353735663830666435623832353139313530363032303833303135313631353130373831363134663135353635623831383130333831383131313135363131313233353736313131323336313535383235363562356638323335363033653139383333363033303138313132363135336130353735663830666435623566383038333335363031653139383433363033303138313132363135613932353735663830666435623833303138303335393135303630303136303031363034303162303338323131313536313561616235373566383066643562363032303031393135303336383139303033383231333135363133383134353735663830666435623630323038313532356636313132353036303230383330313834383636313536306335363562356636303430383233363033313231353631356165323537356638306664356236313561656136313531666635363562363135616633383336313465313235363562383135323630323038303834303133353630303136303031363034303162303338303832313131353631356230663537356638306664356239303835303139303336363031663833303131323631356232313537356638306664356238313335383138313131313536313562333335373631356233333631353163333536356238303630303531623931353036313562343438343833303136313532323135363562383138313532393138333031383430313931383438313031393033363834313131353631356235643537356638306664356239333835303139333562383338353130313536313562383735373834333539323530363135623737383336313530633035363562383238323532393338353031393339303835303139303631356236323536356239343836303139343930393435323530393239353934353035303530353035303536356235663630343038333031363366666666666666663833353131363834353236303230383038343031353136303430363032303837303135323832383135313830383535323630363038383031393135303630323038333031393435303566393235303562383038333130313536313237666635373834353136303031363030313630613031623033313638323532393338333031393336303031393239303932303139313930383330313930363135626362353635623630323038313532383135313630323038323031353235663630323038333031353136306530363034303834303135323631356331613631303130303834303138323631353031353536356239303530363034303834303135313630316631393830383538343033303136303630383630313532363135633338383338333631353031353536356239323530363030313630303136303430316230333630363038373031353131363630383038363031353236303830383630313531393135303830383538343033303136306130383630313532363135633638383338333631356239393536356239323530363061303836303135313931353038303835383430333031363063303836303135323530363135633836383238323631356239393536356239313530353036306330383430313531363135636133363065303835303138323630303136303031363034303162303331363930353235363562353039333932353035303530353635623566383036303430383338353033313231353631356362633537356638306664356238323531393135303630323038333031353136303031363030313630343031623033383131313135363135636438353735663830666435623631356365343835383238363031363135323737353635623931353035303932353039323930353035363562356636303230383238343033313231353631356366653537356638306664356236313238383638323631346630343536356235663630323038323834303331323135363135643137353735663830666435623831333536306666383131363831313436313238383635373566383066643562356636303031363030313630343031623033383038333136383138313033363135356331353736313535633136313535383235363562356636303230383238343033313231353631356435323537356638306664356238313531363132383836383136313465303535366665343331373731336637656362646464643462633939653935643930336164656461613838336232653763323535313631306264313365326337653437336430353665356264666363653135653533633334303665613637626663653337646364323666353135326435343932383234653433666435653363386163356162303034333137373133663765636264646464346263393965393564393033616465646161383833623265376332353531363130626431336532633765343733643030653932353436643639383935306464643338393130643265313565643164393233636430613762336464653965326136613366333830353635353539636230303962373739623137343232643064663932323233303138623332623464316661343665303731373233643638313765323438366430303362656363353566303065393235343664363938393530646464333839313064326531356564316439323363643061376233646465396532613661336633383035363535353963623037613136343733366636633633343330303038313930303061",
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
// Solidity: function valueToWeight(uint256 value) pure returns(uint64)
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
// Solidity: function valueToWeight(uint256 value) pure returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.ValueToWeight(&_ERC20TokenStakingManager.CallOpts, value)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) pure returns(uint64)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerCallerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _ERC20TokenStakingManager.Contract.ValueToWeight(&_ERC20TokenStakingManager.CallOpts, value)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) pure returns(uint256)
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
// Solidity: function weightToValue(uint64 weight) pure returns(uint256)
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _ERC20TokenStakingManager.Contract.WeightToValue(&_ERC20TokenStakingManager.CallOpts, weight)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) pure returns(uint256)
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

// Initialize is a paid mutator transaction binding the contract method 0xf09969ae.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,address) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerTransactor) Initialize(opts *bind.TransactOpts, settings PoSValidatorManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.contract.Transact(opts, "initialize", settings, token)
}

// Initialize is a paid mutator transaction binding the contract method 0xf09969ae.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,address) settings, address token) returns()
func (_ERC20TokenStakingManager *ERC20TokenStakingManagerSession) Initialize(settings PoSValidatorManagerSettings, token common.Address) (*types.Transaction, error) {
	return _ERC20TokenStakingManager.Contract.Initialize(&_ERC20TokenStakingManager.TransactOpts, settings, token)
}

// Initialize is a paid mutator transaction binding the contract method 0xf09969ae.
//
// Solidity: function initialize(((bytes32,uint64,uint8),uint256,uint256,uint64,uint16,uint8,address) settings, address token) returns()
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
