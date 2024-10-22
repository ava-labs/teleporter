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
	ABI: "[{\"inputs\":[{\"internalType\":\"enumICMInitializable\",\"name\":\"init\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"InvalidBLSKeyLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"}],\"name\":\"InvalidDelegationFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidDelegationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumDelegatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidDelegatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"name\":\"InvalidMaximumChurnPercentage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"}],\"name\":\"InvalidMinStakeDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"InvalidNodeID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addressesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidPChainOwnerThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"InvalidRegistrationExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"InvalidStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"}],\"name\":\"InvalidStakeMultiplier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"encodedSubnetConversionID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedSubnetConversionID\",\"type\":\"bytes32\"}],\"name\":\"InvalidSubnetConversionID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"InvalidTokenAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidationID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"}],\"name\":\"InvalidValidatorManagerAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidatorManagerBlockchainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidValidatorStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWarpMessage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"InvalidWarpOriginSenderAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"}],\"name\":\"InvalidWarpSourceChainID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"churnAmount\",\"type\":\"uint64\"}],\"name\":\"MaxChurnRateExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newValidatorWeight\",\"type\":\"uint64\"}],\"name\":\"MaxWeightExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"}],\"name\":\"MinStakeDurationNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"NodeAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PChainOwnerAddressesNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UnauthorizedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"validRegistration\",\"type\":\"bool\"}],\"name\":\"UnexpectedRegistrationStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorIneligibleForRewards\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorNotPoS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"DelegationEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"delegatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"DelegatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"InitialValidatorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"UptimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"}],\"name\":\"ValidationPeriodCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidationPeriodEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidationPeriodRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"ValidationRewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"ValidatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"setWeightMessageID\",\"type\":\"bytes32\"}],\"name\":\"ValidatorWeightUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_LENGTH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERC20_STAKING_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POS_VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPTIME_REWARDS_THRESHOLD_PERCENTAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"claimDelegationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"claimValidationRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"startingWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"messageNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endedAt\",\"type\":\"uint64\"}],\"internalType\":\"structValidator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"churnPeriodSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\"}],\"internalType\":\"structValidatorManagerSettings\",\"name\":\"baseSettings\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minimumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"minimumStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"contractIRewardCalculator\",\"name\":\"rewardCalculator\",\"type\":\"address\"}],\"internalType\":\"structPoSValidatorManagerSettings\",\"name\":\"settings\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20Mintable\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"delegationAmount\",\"type\":\"uint256\"}],\"name\":\"initializeDelegatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"internalType\":\"structValidatorRegistrationInput\",\"name\":\"registrationInput\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"delegationFeeBips\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minStakeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"initializeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structSubnetConversionData\",\"name\":\"subnetConversionData\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"}],\"name\":\"registeredValidators\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendEndValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"resendRegisterValidatorMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"resendUpdateDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"submitUptimeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"valueToWeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"weightToValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051615fa0380380615fa083398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b615e538061014d5f395ff3fe608060405234801561000f575f80fd5b5060043610610208575f3560e01c806376f786211161011f578063ba3a4b97116100a9578063d5f20ff611610079578063d5f20ff614610475578063df93d8de14610495578063e4a63c401461049f578063f09969ae146104b3578063fd7ac5e7146104c6575f80fd5b8063ba3a4b9714610433578063bc5fbfec14610446578063bee0a03f1461045a578063c974d1b61461046d575f80fd5b80639e1bc4ef116100ef5780639e1bc4ef146103cb578063a3a65e48146103de578063afb98096146103f1578063afba878a14610405578063b771b3bc1461040d575f80fd5b806376f786211461038a57806380dd672f1461039d5780638280a25a146103b057806393e24598146103b8575f80fd5b806335455ded116101a057806360305d621161017057806360305d621461032d57806360ad77841461034a578063620658561461035d57806366435abf14610370578063732214f814610383575f80fd5b806335455ded146102ca5780633a1cfff6146102e6578063467ef06f146102f95780634bee00401461030c575f80fd5b806320d91b7a116101db57806320d91b7a1461026657806325e1c776146102795780632893d0771461028c5780632e2194d81461029f575f80fd5b80630118acc41461020c5780630322ed9814610221578063151d30d1146102345780631ec4472414610253575b5f80fd5b61021f61021a366004614e2a565b6104d9565b005b61021f61022f366004614e65565b6104eb565b61023c600a81565b60405160ff90911681526020015b60405180910390f35b61021f610261366004614e2a565b61077b565b61021f610274366004614e7c565b610787565b61021f610287366004614eca565b610d4f565b61021f61029a366004614eca565b610dc3565b6102b26102ad366004614e65565b611113565b6040516001600160401b03909116815260200161024a565b6102d361271081565b60405161ffff909116815260200161024a565b61021f6102f4366004614e2a565b611129565b61021f610307366004614eeb565b611135565b61031f61031a366004614f29565b611223565b60405190815260200161024a565b610335601481565b60405163ffffffff909116815260200161024a565b61021f610358366004614eca565b611258565b61031f61036b366004614f90565b611530565b6102b261037e366004614e65565b611549565b61031f5f81565b61021f610398366004614e2a565b61155d565b61021f6103ab366004614eca565b61156a565b61023c603081565b61021f6103c6366004614e65565b6117ab565b61031f6103d9366004614fab565b611855565b61021f6103ec366004614eeb565b611881565b61031f5f80516020615d9e83398151915281565b61023c605081565b61041b6005600160991b0181565b6040516001600160a01b03909116815260200161024a565b61021f610441366004614e65565b611a76565b61031f5f80516020615dbe83398151915281565b61021f610468366004614e65565b611ce1565b61023c601481565b610488610483366004614e65565b611e1e565b60405161024a9190615040565b6102b26202a30081565b61031f5f80516020615d7e83398151915281565b61021f6104c13660046150d4565b611f6d565b61031f6104d4366004615112565b61204c565b6104e683838360016120a7565b505050565b5f8181525f80516020615dfe8339815191526020526040808220815160e0810190925280545f80516020615dbe83398151915293929190829060ff16600581111561053857610538614fcb565b600581111561054957610549614fcb565b815260200160018201805461055d9061517d565b80601f01602080910402602001604051908101604052809291908181526020018280546105899061517d565b80156105d45780601f106105ab576101008083540402835291602001916105d4565b820191905f5260205f20905b8154815290600101906020018083116105b757829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b909104811660808301526003928301541660a0909101529091508151600581111561063f5761063f614fcb565b1461067b575f8381526007830160205260409081902054905163170cc93360e21b81526106729160ff16906004016151b5565b60405180910390fd5b60608101516040516342a2e0b560e11b8152600481018590526001600160401b0390911660248201525f60448201526005600160991b019063ee5b48eb9073__$fd0c147b4031eef6079b0498cbafa865f0$__90638545c16a906064015f60405180830381865af41580156106f2573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261071991908101906152be565b6040518263ffffffff1660e01b815260040161073591906152ef565b6020604051808303815f875af1158015610751573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107759190615301565b50505050565b6104e68383835f6120a7565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb09545f80516020615dbe8339815191529060ff16156107d957604051637fab81e560e01b815260040160405180910390fd5b6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa15801561081c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108409190615301565b836020013514610869576040516372b0a7e760e11b815260208401356004820152602401610672565b3061087a6060850160408601615318565b6001600160a01b0316146108bd576108986060840160408501615318565b604051632f88120d60e21b81526001600160a01b039091166004820152602401610672565b5f6108cb6060850185615333565b905090505f805b828163ffffffff161015610bb3575f6108ee6060880188615333565b8363ffffffff1681811061090457610904615378565b9050602002810190610916919061538c565b61091f906153f7565b80516040519192505f91600888019161093791615472565b9081526020016040518091039020541461096757805160405163a41f772f60e01b815261067291906004016152ef565b5f6002885f01358460405160200161099692919091825260e01b6001600160e01b031916602082015260240190565b60408051601f19818403018152908290526109b091615472565b602060405180830381855afa1580156109cb573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906109ee9190615301565b90508086600801835f0151604051610a069190615472565b90815260408051602092819003830181209390935560e0830181526002835284518284015284810180516001600160401b03908116858401525f60608601819052915181166080860152421660a085015260c0840181905284815260078a01909252902081518154829060ff19166001836005811115610a8857610a88614fcb565b021790555060208201516001820190610aa190826154c7565b506040828101516002830180546060860151608087015160a08801516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909301516003909201805467ffffffffffffffff191692841692909217909155830151610b46911685615596565b8251604051919550610b5791615472565b60408051918290038220908401516001600160401b031682529082907f9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf9060200160405180910390a3505080610bac906155a9565b90506108d2565b50600483018190555f73__$fd0c147b4031eef6079b0498cbafa865f0$__631e6d9789610bdf876123d3565b604001516040518263ffffffff1660e01b8152600401610bff91906152ef565b602060405180830381865af4158015610c1a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c3e9190615301565b90505f73__$fd0c147b4031eef6079b0498cbafa865f0$__63862bfa63886040518263ffffffff1660e01b8152600401610c7891906156f6565b5f60405180830381865af4158015610c92573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610cb991908101906152be565b90505f600282604051610ccc9190615472565b602060405180830381855afa158015610ce7573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610d0a9190615301565b9050828114610d3657604051631872fc8d60e01b81526004810182905260248101849052604401610672565b5050506009909201805460ff1916600117905550505050565b610d58826124e9565b610d78576040516330efa98b60e01b815260048101839052602401610672565b5f610d8283611e1e565b5190506002816005811115610d9957610d99614fcb565b14610db9578060405163170cc93360e21b815260040161067291906151b5565b6107758383612524565b610dcb61280a565b5f80516020615d9e8339815191525f610de384611e1e565b9050600281516005811115610dfa57610dfa614fcb565b14610e1b57805160405163170cc93360e21b815261067291906004016151b5565b610e24846124e9565b610e44576040516330efa98b60e01b815260048101859052602401610672565b5f8481526004830160205260409020546001600160a01b03163314610e8a57335b604051636e2ccd7560e11b81526001600160a01b039091166004820152602401610672565b5f84815260048301602052604090205460a08201514291610ebc91600160b01b9091046001600160401b031690615781565b6001600160401b0316816001600160401b03161015610ef95760405163fb6ce63f60e01b81526001600160401b0382166004820152602401610672565b5f610f048686612524565b5f87815260048601602052604081206001015491925090610f3590600160401b90046001600160401b0316836157a1565b5f888152600487016020526040812060010154919250600160801b9091046001600160401b031690819003610f6b575060a08401515b610f76828286612841565b610f9657604051635bff683f60e11b815260048101899052602401610672565b600386015460408601515f916001600160a01b03169063778c06b590610fbb90611530565b848589885f806040518863ffffffff1660e01b8152600401610fe397969594939291906157c1565b602060405180830381865afa158015610ffe573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110229190615301565b90506064605060ff168760a001518761103b91906157a1565b61104591906157ff565b61104f919061583e565b5f8a8152600489016020526040902060018101805477ffffffffffffffffffffffffffffffff00000000000000001916600160401b6001600160401b039485160267ffffffffffffffff60801b191617600160801b9389169390930292909217909155546110c6906001600160a01b03168261288d565b60405181908a907f69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a905f90a35050505050505061110f60015f80516020615dde83398151915255565b5050565b5f61112364e8d4a5100083615863565b92915050565b6104e68383835f612916565b61113d61280a565b5f80516020615d9e8339815191525f8061115684612b77565b91509150611163826124e9565b61116f5750505061120a565b5f82815260048085016020526040909120546001600160a01b0316908251600581111561119e5761119e614fcb565b036111ef575f838152600785016020526040812080549190556111c1828261288d565b604051819085907f69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a905f90a3505b611205816112008460400151611530565b612f2a565b505050505b61122060015f80516020615dde83398151915255565b50565b5f61122c61280a565b61123885858585612f4d565b905061125060015f80516020615dde83398151915255565b949350505050565b5f8281525f80516020615d5e8339815191526020526040808220815160e0810190925280545f80516020615d9e83398151915293929190829060ff1660038111156112a5576112a5614fcb565b60038111156112b6576112b6614fcb565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f61132c82611e1e565b905060018351600381111561134357611343614fcb565b14611364578251604051633b0d540d60e21b81526106729190600401615876565b60048151600581111561137957611379614fcb565b0361138f576113878661316b565b505050505050565b5f8073__$fd0c147b4031eef6079b0498cbafa865f0$__639de23d406113b4896123d3565b604001516040518263ffffffff1660e01b81526004016113d491906152ef565b606060405180830381865af41580156113ef573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114139190615890565b509150915081841461144057846040015160405163089938b360e11b815260040161067291815260200190565b806001600160401b031683606001516001600160401b031610806114795750806001600160401b03168560a001516001600160401b0316115b156114a257604051632e19bc2d60e11b81526001600160401b0382166004820152602401610672565b5f888152600587016020908152604091829020805460ff1916600290811782550180546001600160401b034216600160401b81026fffffffffffffffff00000000000000001990921691909117909155915191825285918a917f047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6910160405180910390a35050505050505050565b5f6111236001600160401b03831664e8d4a510006158d0565b5f61155382611e1e565b6080015192915050565b6104e68383836001612916565b61157261280a565b5f8281525f80516020615d5e8339815191526020526040808220815160e0810190925280545f80516020615d9e83398151915293929190829060ff1660038111156115bf576115bf614fcb565b60038111156115d0576115d0614fcb565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c090910152905060038151600381111561164957611649614fcb565b1461166a578051604051633b0d540d60e21b81526106729190600401615876565b60046116798260400151611e1e565b51600581111561168b5761168b614fcb565b1461178a575f61169a846123d3565b90505f8073__$fd0c147b4031eef6079b0498cbafa865f0$__639de23d4084604001516040518263ffffffff1660e01b81526004016116d991906152ef565b606060405180830381865af41580156116f4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117189190615890565b5091509150818460400151146117445760405163089938b360e11b815260048101839052602401610672565b806001600160401b03168460c001516001600160401b0316111561178657604051632e19bc2d60e11b81526001600160401b0382166004820152602401610672565b5050505b6117938461316b565b505061110f60015f80516020615dde83398151915255565b5f80516020615d9e8339815191525f6117c383611e1e565b51905060048160058111156117da576117da614fcb565b146117fa578060405163170cc93360e21b815260040161067291906151b5565b5f8381526004830160205260409020546001600160a01b0316331461181f5733610e65565b5f83815260078301602090815260408083208054908490556004860190925290912054610775906001600160a01b03168261288d565b5f61185e61280a565b611869833384613352565b905061112360015f80516020615dde83398151915255565b5f80516020615dbe8339815191525f8073__$fd0c147b4031eef6079b0498cbafa865f0$__632e43ceb56118b4866123d3565b604001516040518263ffffffff1660e01b81526004016118d491906152ef565b6040805180830381865af41580156118ee573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061191291906158e7565b915091508061193857604051632d07135360e01b81528115156004820152602401610672565b5f828152600684016020526040902080546119529061517d565b90505f036119765760405163089938b360e11b815260048101839052602401610672565b60015f83815260078501602052604090205460ff16600581111561199c5761199c614fcb565b146119cf575f8281526007840160205260409081902054905163170cc93360e21b81526106729160ff16906004016151b5565b5f82815260068401602052604081206119e791614d79565b5f828152600784016020908152604091829020805460ff1916600290811782550180546001600160401b0342818116600160c01b026001600160c01b0390931692909217928390558451600160801b9093041682529181019190915283917ff8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568910160405180910390a250505050565b5f8181525f80516020615d5e8339815191526020526040808220815160e0810190925280545f80516020615d9e83398151915293929190829060ff166003811115611ac357611ac3614fcb565b6003811115611ad457611ad4614fcb565b8152815461010090046001600160a01b0316602082015260018083015460408301526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015290915081516003811115611b4d57611b4d614fcb565b14158015611b6e5750600381516003811115611b6b57611b6b614fcb565b14155b15611b8f578051604051633b0d540d60e21b81526106729190600401615876565b5f611b9d8260400151611e1e565b905080606001516001600160401b03165f03611bcf576040516339b894f960e21b815260048101859052602401610672565b6040808301516060830151608084015192516342a2e0b560e11b81526005600160991b019363ee5b48eb9373__$fd0c147b4031eef6079b0498cbafa865f0$__93638545c16a93611c3d93906004019283526001600160401b03918216602084015216604082015260600190565b5f60405180830381865af4158015611c57573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611c7e91908101906152be565b6040518263ffffffff1660e01b8152600401611c9a91906152ef565b6020604051808303815f875af1158015611cb6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611cda9190615301565b5050505050565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb066020526040902080545f80516020615dbe8339815191529190611d289061517d565b90505f03611d4c5760405163089938b360e11b815260048101839052602401610672565b60015f83815260078301602052604090205460ff166005811115611d7257611d72614fcb565b14611da5575f8281526007820160205260409081902054905163170cc93360e21b81526106729160ff16906004016151b5565b5f82815260068201602052604090819020905163ee5b48eb60e01b81526005600160991b019163ee5b48eb91611dde919060040161590a565b6020604051808303815f875af1158015611dfa573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104e69190615301565b611e26614db0565b5f8281525f80516020615dfe833981519152602052604090819020815160e0810190925280545f80516020615dbe833981519152929190829060ff166005811115611e7357611e73614fcb565b6005811115611e8457611e84614fcb565b8152602001600182018054611e989061517d565b80601f0160208091040260200160405190810160405280929190818152602001828054611ec49061517d565b8015611f0f5780601f10611ee657610100808354040283529160200191611f0f565b820191905f5260205f20905b815481529060010190602001808311611ef257829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b9091048116608083015260039092015490911660a0909101529392505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff1680611fb6575080546001600160401b03808416911610155b15611fd45760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b178155611fff848461362a565b805460ff60401b191681556040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050565b6040515f905f80516020615dbe833981519152907fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb089061208f9086908690615994565b90815260200160405180910390205491505092915050565b5f8481525f80516020615d5e8339815191526020526040808220815160e0810190925280545f80516020615d9e83398151915293929190829060ff1660038111156120f4576120f4614fcb565b600381111561210557612105614fcb565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f61217b82611e1e565b905060028351600381111561219257612192614fcb565b146121b3578251604051633b0d540d60e21b81526106729190600401615876565b60208301516001600160a01b03163314612256575f828152600485016020526040902054336001600160a01b0390911603612250575f82815260048501602052604090205460a082015161221791600160b01b90046001600160401b031690615781565b6001600160401b031642101561224b5760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610672565b612256565b33610e65565b60028151600581111561226b5761226b614fcb565b0361236c578615612282576122808287612524565b505b5f8881526005850160205260409020805460ff19166003179055606083015160808201516122bb9184916122b691906157a1565b613644565b505f898152600586016020526040812060020180546001600160401b03909316600160c01b026001600160c01b03909316929092179091556122fc8461381b565b9050858015612309575080155b1561232a57604051631036cf9160e11b8152600481018a9052602401610672565b5f8981526006860160205260408082208390555184918b917f366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed579190a3506123c9565b60048151600581111561238157612381614fcb565b036123ad5761238f8361381b565b5f8981526006860160205260409020556123a88861316b565b6123c9565b805160405163170cc93360e21b815261067291906004016151b5565b5050505050505050565b60408051606080820183525f8083526020830152918101919091526040516306f8253560e41b815263ffffffff831660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa158015612437573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261245e91908101906159a3565b915091508061248057604051636b2f19e960e01b815260040160405180910390fd5b8151156124a6578151604051636ba589a560e01b81526004810191909152602401610672565b60208201516001600160a01b0316156124e2576020820151604051624de75d60e31b81526001600160a01b039091166004820152602401610672565b5092915050565b5f9081527f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0460205260409020546001600160a01b0316151590565b6040516306f8253560e41b815263ffffffff821660048201525f90819081906005600160991b0190636f825350906024015f60405180830381865afa15801561256f573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261259691908101906159a3565b91509150806125b857604051636b2f19e960e01b815260040160405180910390fd5b6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa1580156125fb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061261f9190615301565b825114612645578151604051636ba589a560e01b81526004810191909152602401610672565b60208201516001600160a01b031615612681576020820151604051624de75d60e31b81526001600160a01b039091166004820152602401610672565b5f8073__$fd0c147b4031eef6079b0498cbafa865f0$__63088c246385604001516040518263ffffffff1660e01b81526004016126be91906152ef565b6040805180830381865af41580156126d8573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126fc9190615a33565b915091508187146127235760405163089938b360e11b815260048101889052602401610672565b5f8781527f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0460205260409020600101545f80516020615d9e833981519152906001600160401b0390811690831611156127e0575f888152600482016020908152604091829020600101805467ffffffffffffffff19166001600160401b038616908117909155915191825289917fec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435910160405180910390a26127ff565b5f8881526004820160205260409020600101546001600160401b031691505b509695505050505050565b5f80516020615dde83398151915280546001190161283b57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f605061284e84846157a1565b61285891906157ff565b6001600160401b031661286c8560646157ff565b6001600160401b0316101561288257505f612886565b5060015b9392505050565b5f5f80516020615d7e83398151915280546040516340c10f1960e01b81526001600160a01b038681166004830152602482018690529293509116906340c10f19906044015f604051808303815f87803b1580156128e8575f80fd5b505af11580156128fa573d5f803e3d5ffd5b50505050505050565b60015f80516020615dde83398151915255565b5f80516020615d9e8339815191525f61292e866139e4565b9050612939866124e9565b612944575050610775565b5f8681526004830160205260409020546001600160a01b031633146129695733610e65565b5f86815260048301602052604090205460a082015161299891600160b01b90046001600160401b031690615781565b6001600160401b03168160c001516001600160401b031610156129df5760c081015160405163fb6ce63f60e01b81526001600160401b039091166004820152602401610672565b5f85156129f7576129f08786612524565b9050612a15565b505f8681526004830160205260409020600101546001600160401b03165b5f878152600484016020526040812060010154612a4290600160401b90046001600160401b0316836157a1565b5f898152600486016020526040812060010154919250600160801b9091046001600160401b031690819003612a78575060a08301515b858015612a915750612a8f82828660c00151612841565b155b15612ab257604051635bff683f60e11b8152600481018a9052602401610672565b600385015460408501515f916001600160a01b03169063778c06b590612ad790611530565b84858960c00151885f806040518863ffffffff1660e01b8152600401612b0397969594939291906157c1565b602060405180830381865afa158015612b1e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b429190615301565b905080866007015f8c81526020019081526020015f205f828254612b669190615596565b909155505050505050505050505050565b5f612b80614db0565b5f80516020615dbe8339815191525f8073__$fd0c147b4031eef6079b0498cbafa865f0$__632e43ceb5612bb3886123d3565b604001516040518263ffffffff1660e01b8152600401612bd391906152ef565b6040805180830381865af4158015612bed573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c1191906158e7565b915091508015612c3857604051632d07135360e01b81528115156004820152602401610672565b5f828152600784016020526040808220815160e081019092528054829060ff166005811115612c6957612c69614fcb565b6005811115612c7a57612c7a614fcb565b8152602001600182018054612c8e9061517d565b80601f0160208091040260200160405190810160405280929190818152602001828054612cba9061517d565b8015612d055780601f10612cdc57610100808354040283529160200191612d05565b820191905f5260205f20905b815481529060010190602001808311612ce857829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b909104811660808301526003928301541660a09091015290915081516005811115612d7057612d70614fcb565b14158015612d915750600181516005811115612d8e57612d8e614fcb565b14155b15612db257805160405163170cc93360e21b815261067291906004016151b5565b600381516005811115612dc757612dc7614fcb565b03612dd55760048152612dda565b600581525b836008018160200151604051612df09190615472565b90815260408051602092819003830190205f908190558581526007870190925290208151815483929190829060ff19166001836005811115612e3457612e34614fcb565b021790555060208201516001820190612e4d90826154c7565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909201516003909101805467ffffffffffffffff19169190921617905580516005811115612ef357612ef3614fcb565b60405184907f1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16905f90a39196919550909350505050565b5f80516020615d7e8339815191525461110f906001600160a01b03168383613cc8565b7f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d02545f905f80516020615d9e83398151915290600160401b900461ffff9081169086161080612fa1575061271061ffff8616115b15612fc557604051635f12e6c360e11b815261ffff86166004820152602401610672565b60028101546001600160401b039081169085161015613001576040516202a06d60e11b81526001600160401b0385166004820152602401610672565b80548310806130135750806001015483115b156130345760405163222d164360e21b815260048101849052602401610672565b5f61303e84613d27565b90505f61304a82611113565b90505f6130578983613d4a565b90506040518060c0016040528061306b3390565b6001600160a01b03908116825261ffff808c166020808501919091526001600160401b03808d166040808701919091525f6060808801829052608080890183905260a09889018390528a83526004909d0185529082902088518154958a0151938a01518516600160b01b0267ffffffffffffffff60b01b1994909716600160a01b026001600160b01b03199096169716969096179390931716929092178355840151600190920180549885015194909301518116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b031990991692909116919091179690961791909116949094179093555090915050949350505050565b5f8181525f80516020615d5e8339815191526020526040808220815160e0810190925280545f80516020615d9e83398151915293929190829060ff1660038111156131b8576131b8614fcb565b60038111156131c9576131c9614fcb565b8152815461010090046001600160a01b03166020808301919091526001808401546040808501919091526002948501546001600160401b038082166060870152600160401b820481166080870152600160801b8204811660a0870152600160c01b9091041660c090940193909352848301515f89815260058901845284812080546001600160a81b031916815592830181905591909401819055600687019091529081208054908290559293509091908082156132f7575f848152600487016020526040902054612710906132a990600160a01b900461ffff16856158d0565b6132b39190615863565b915081866007015f8681526020019081526020015f205f8282546132d79190615596565b909155506132e790508284615a56565b90506132f785602001518261288d565b61330c85602001516112008760600151611530565b6040805182815260208101849052859189917f8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993910160405180910390a350505050505050565b5f5f80516020615d9e8339815191528161336e6102ad85613d27565b90505f61337a87611e1e565b9050613385876124e9565b6133a5576040516330efa98b60e01b815260048101889052602401610672565b6002815160058111156133ba576133ba614fcb565b146133db57805160405163170cc93360e21b815261067291906004016151b5565b5f8282608001516133ec9190615781565b905083600201600a9054906101000a90046001600160401b0316826040015161341591906157ff565b6001600160401b0316816001600160401b0316111561345257604051636d51fe0560e11b81526001600160401b0382166004820152602401610672565b5f8061345e8a84613644565b915091505f8a8360405160200161348c92919091825260c01b6001600160c01b031916602082015260280190565b60408051601f19818403018152828252805160209091012060e08301909152915080600181526001600160a01b038c1660208083019190915260408083018f90526001600160401b03808b1660608501525f6080850181905290881660a085015260c090930183905284835260058b01909152902081518154829060ff1916600183600381111561351f5761351f614fcb565b02179055506020828101518254610100600160a81b0319166101006001600160a01b039283160217835560408085015160018501556060808601516002909501805460808089015160a08a015160c0909a01516001600160401b03998a166001600160801b031990941693909317600160401b918a1691909102176001600160801b0316600160801b998916999099026001600160c01b031698909817600160c01b91881691909102179055815189861681528a861694810194909452938b1690830152918101859052908c16918d9184917fb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426910160405180910390a49a9950505050505050505050565b613632614332565b61363b8261437d565b61110f816143eb565b5f8281525f80516020615dfe833981519152602052604081206002015481905f80516020615dbe83398151915290600160801b90046001600160401b031661368c8582614453565b5f6136968761462d565b5f888152600785016020526040808220600201805467ffffffffffffffff60801b1916600160801b6001600160401b038c81169182029290921790925591516342a2e0b560e11b8152600481018c905291841660248301526044820152919250906005600160991b019063ee5b48eb9073__$fd0c147b4031eef6079b0498cbafa865f0$__90638545c16a906064015f60405180830381865af415801561373f573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261376691908101906152be565b6040518263ffffffff1660e01b815260040161378291906152ef565b6020604051808303815f875af115801561379e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137c29190615301565b604080516001600160401b038a811682526020820184905282519394508516928b927f07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df928290030190a3909450925050505b9250929050565b5f805f80516020615d9e83398151915290505f61383b8460400151611e1e565b90505f60038251600581111561385357613853614fcb565b1480613871575060048251600581111561386f5761386f614fcb565b145b15613881575060c08101516138be565b60028251600581111561389657613896614fcb565b036138a25750426138be565b815160405163170cc93360e21b815261067291906004016151b5565b84608001516001600160401b0316816001600160401b0316116138e557505f949350505050565b6040808601515f90815260048501602052206001015460a0830151613914916001600160401b03169083612841565b61392257505f949350505050565b600383015460608601516001600160a01b039091169063778c06b59061394790611530565b60a085015160808901516040808b01515f9081526004808b016020528282206001015492516001600160e01b031960e089901b16815261399c969594938a936001600160401b039091169290918291016157c1565b602060405180830381865afa1580156139b7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906139db9190615301565b95945050505050565b6139ec614db0565b5f8281525f80516020615dfe8339815191526020526040808220815160e0810190925280545f80516020615dbe83398151915293929190829060ff166005811115613a3957613a39614fcb565b6005811115613a4a57613a4a614fcb565b8152602001600182018054613a5e9061517d565b80601f0160208091040260200160405190810160405280929190818152602001828054613a8a9061517d565b8015613ad55780601f10613aac57610100808354040283529160200191613ad5565b820191905f5260205f20905b815481529060010190602001808311613ab857829003601f168201915b50505091835250506002828101546001600160401b038082166020850152600160401b820481166040850152600160801b820481166060850152600160c01b9091048116608084015260039093015490921660a09091015290915081516005811115613b4357613b43614fcb565b14613b76575f8481526007830160205260409081902054905163170cc93360e21b81526106729160ff16906004016151b5565b60038152426001600160401b031660c08201525f84815260078301602052604090208151815483929190829060ff19166001836005811115613bba57613bba614fcb565b021790555060208201516001820190613bd390826154c7565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909201516003909101805467ffffffffffffffff1916919092161790555f613c718582613644565b6080840151604080516001600160401b03909216825242602083015291935083925087917f13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67910160405180910390a3509392505050565b6040516001600160a01b038381166024830152604482018390526104e691859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506146a2565b5f611123825f80516020615d7e833981519152546001600160a01b031690614703565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb09545f9060ff16613d8e57604051637fab81e560e01b815260040160405180910390fd5b5f80516020615dbe83398151915242613dad6060860160408701614f90565b6001600160401b0316111580613de75750613dcb6202a30042615596565b613ddb6060860160408701614f90565b6001600160401b031610155b15613e2157613dfc6060850160408601614f90565b604051635879da1360e11b81526001600160401b039091166004820152602401610672565b613e36613e316060860186615a69565b61485c565b613e46613e316080860186615a69565b6030613e556020860186615a7d565b905014613e8757613e696020850185615a7d565b6040516326475b2f60e11b8152610672925060040190815260200190565b613e918480615a7d565b90505f03613ebe57613ea38480615a7d565b604051633e08a12560e11b8152600401610672929190615abf565b5f60088201613ecd8680615a7d565b604051613edb929190615994565b90815260200160405180910390205414613f1457613ef98480615a7d565b60405163a41f772f60e01b8152600401610672929190615abf565b613f1e835f614453565b6040805160e08101909152815481525f90819073__$fd0c147b4031eef6079b0498cbafa865f0$__9063e047b2839060208101613f5b8a80615a7d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250505090825250602090810190613fa3908b018b615a7d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250505090825250602001613fec60608b0160408c01614f90565b6001600160401b0316815260200161400760608b018b615a69565b61401090615ad2565b815260200161402260808b018b615a69565b61402b90615ad2565b8152602001886001600160401b03168152506040518263ffffffff1660e01b81526004016140599190615bf4565b5f60405180830381865af4158015614073573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261409a9190810190615cab565b5f828152600686016020526040902091935091506140b882826154c7565b5081600884016140c88880615a7d565b6040516140d6929190615994565b9081526040519081900360200181209190915563ee5b48eb60e01b81525f906005600160991b019063ee5b48eb906141129085906004016152ef565b6020604051808303815f875af115801561412e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906141529190615301565b6040805160e0810190915290915080600181526020016141728980615a7d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052509385525050506001600160401b0389166020808401829052604080850184905260608501929092526080840183905260a0909301829052868252600788019092522081518154829060ff1916600183600581111561420157614201614fcb565b02179055506020820151600182019061421a90826154c7565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909201516003909101805467ffffffffffffffff191691909216179055806142b88880615a7d565b6040516142c6929190615994565b6040518091039020847fb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430898b60400160208101906143049190614f90565b604080516001600160401b0393841681529290911660208301520160405180910390a4509095945050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661437b57604051631afcd79f60e31b815260040160405180910390fd5b565b614385614332565b61438e816149c5565b6143966149de565b611220606082013560808301356143b360c0850160a08601614f90565b6143c360e0860160c08701615cee565b6143d4610100870160e08801615d07565b6143e661012088016101008901615318565b6149ee565b6143f3614332565b5f80516020615d7e8339815191526001600160a01b03821661443357604051637330680360e01b81526001600160a01b0383166004820152602401610672565b80546001600160a01b0319166001600160a01b0392909216919091179055565b5f80516020615dbe8339815191525f6001600160401b0380841690851611156144875761448083856157a1565b9050614494565b61449184846157a1565b90505b6040805160808101825260028401548082526003850154602083015260048501549282019290925260058401546001600160401b03166060820152429115806144f65750600184015481516144f2916001600160401b031690615596565b8210155b1561451c576001600160401b03831660608201528181526040810151602082015261453b565b828160600181815161452e9190615781565b6001600160401b03169052505b606081015161454b9060646157ff565b602082015160018601546001600160401b0392909216916145769190600160401b900460ff166158d0565b10156145a657606081015160405163dfae880160e01b81526001600160401b039091166004820152602401610672565b856001600160401b0316816040018181516145c19190615596565b9052506040810180516001600160401b03871691906145e1908390615a56565b905250805160028501556020810151600385015560408101516004850155606001516005909301805467ffffffffffffffff19166001600160401b039094169390931790925550505050565b5f8181525f80516020615dfe8339815191526020526040812060020180545f80516020615dbe833981519152919060089061467790600160401b90046001600160401b0316615d27565b91906101000a8154816001600160401b0302191690836001600160401b031602179055915050919050565b5f6146b66001600160a01b03841683614b1d565b905080515f141580156146da5750808060200190518101906146d89190615d42565b155b156104e657604051635274afe760e01b81526001600160a01b0384166004820152602401610672565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038516906370a0823190602401602060405180830381865afa158015614749573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061476d9190615301565b90506147846001600160a01b038516333086614b2a565b6040516370a0823160e01b81523060048201525f906001600160a01b038616906370a0823190602401602060405180830381865afa1580156147c8573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906147ec9190615301565b90508181116148525760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610672565b6139db8282615a56565b6148696020820182614eeb565b63ffffffff1615801561488957506148846020820182615333565b151590505b156148d05761489b6020820182614eeb565b6148a86020830183615333565b60405163c08a0f1d60e01b815263ffffffff9093166004840152602483015250604401610672565b6148dd6020820182615333565b90506148ec6020830183614eeb565b63ffffffff1611156149055761489b6020820182614eeb565b60015b6149156020830183615333565b905081101561110f5761492b6020830183615333565b614936600184615a56565b81811061494557614945615378565b905060200201602081019061495a9190615318565b6001600160a01b03166149706020840184615333565b8381811061498057614980615378565b90506020020160208101906149959190615318565b6001600160a01b031610156149bd57604051630dbc8d5f60e31b815260040160405180910390fd5b600101614908565b6149cd614332565b6149d5614b63565b61122081614b6b565b6149e6614332565b61437b614c53565b6149f6614332565b5f80516020615d9e83398151915261ffff84161580614a1a575061271061ffff8516115b15614a3e57604051635f12e6c360e11b815261ffff85166004820152602401610672565b85871115614a625760405163222d164360e21b815260048101889052602401610672565b60ff83161580614a755750600a60ff8416115b15614a985760405163170db35960e31b815260ff84166004820152602401610672565b95865560018601949094556002850180546001600160401b039490941669ffffffffffffffffffff1990941693909317600160401b61ffff93909316929092029190911767ffffffffffffffff60501b191660ff91909116600160501b02179055600390910180546001600160a01b0319166001600160a01b03909216919091179055565b606061288683835f614c5b565b6040516001600160a01b0384811660248301528381166044830152606482018390526107759186918216906323b872dd90608401613cf5565b61437b614332565b614b73614332565b80355f80516020615dbe8339815191529081556014614b986060840160408501615d07565b60ff161180614bb75750614bb26060830160408401615d07565b60ff16155b15614beb57614bcc6060830160408401615d07565b604051634a59bbff60e11b815260ff9091166004820152602401610672565b614bfb6060830160408401615d07565b60018201805460ff92909216600160401b0260ff60401b19909216919091179055614c2c6040830160208401614f90565b600191909101805467ffffffffffffffff19166001600160401b0390921691909117905550565b612903614332565b606081471015614c805760405163cd78605960e01b8152306004820152602401610672565b5f80856001600160a01b03168486604051614c9b9190615472565b5f6040518083038185875af1925050503d805f8114614cd5576040519150601f19603f3d011682016040523d82523d5f602084013e614cda565b606091505b5091509150614cea868383614cf4565b9695505050505050565b606082614d0957614d0482614d50565b612886565b8151158015614d2057506001600160a01b0384163b155b15614d4957604051639996b31560e01b81526001600160a01b0385166004820152602401610672565b5080612886565b805115614d605780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b508054614d859061517d565b5f825580601f10614d94575050565b601f0160209004905f5260205f20908101906112209190614ded565b6040805160e08101909152805f81526060602082018190525f604083018190529082018190526080820181905260a0820181905260c09091015290565b5b80821115614e01575f8155600101614dee565b5090565b8015158114611220575f80fd5b803563ffffffff81168114614e25575f80fd5b919050565b5f805f60608486031215614e3c575f80fd5b833592506020840135614e4e81614e05565b9150614e5c60408501614e12565b90509250925092565b5f60208284031215614e75575f80fd5b5035919050565b5f8060408385031215614e8d575f80fd5b82356001600160401b03811115614ea2575f80fd5b830160808186031215614eb3575f80fd5b9150614ec160208401614e12565b90509250929050565b5f8060408385031215614edb575f80fd5b82359150614ec160208401614e12565b5f60208284031215614efb575f80fd5b61288682614e12565b803561ffff81168114614e25575f80fd5b6001600160401b0381168114611220575f80fd5b5f805f8060808587031215614f3c575f80fd5b84356001600160401b03811115614f51575f80fd5b850160a08188031215614f62575f80fd5b9350614f7060208601614f04565b92506040850135614f8081614f15565b9396929550929360600135925050565b5f60208284031215614fa0575f80fd5b813561288681614f15565b5f8060408385031215614fbc575f80fd5b50508035926020909101359150565b634e487b7160e01b5f52602160045260245ffd5b60068110614fef57614fef614fcb565b9052565b5f5b8381101561500d578181015183820152602001614ff5565b50505f910152565b5f815180845261502c816020860160208601614ff3565b601f01601f19169290920160200192915050565b60208152615052602082018351614fdf565b5f602083015160e0604084015261506d610100840182615015565b905060408401516001600160401b0380821660608601528060608701511660808601528060808701511660a08601528060a08701511660c08601528060c08701511660e086015250508091505092915050565b6001600160a01b0381168114611220575f80fd5b5f808284036101408112156150e7575f80fd5b610120808212156150f6575f80fd5b8493508301359050615107816150c0565b809150509250929050565b5f8060208385031215615123575f80fd5b82356001600160401b0380821115615139575f80fd5b818501915085601f83011261514c575f80fd5b81358181111561515a575f80fd5b86602082850101111561516b575f80fd5b60209290920196919550909350505050565b600181811c9082168061519157607f821691505b6020821081036151af57634e487b7160e01b5f52602260045260245ffd5b50919050565b602081016111238284614fdf565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b03811182821017156151f9576151f96151c3565b60405290565b604080519081016001600160401b03811182821017156151f9576151f96151c3565b604051601f8201601f191681016001600160401b0381118282101715615249576152496151c3565b604052919050565b5f6001600160401b03821115615269576152696151c3565b50601f01601f191660200190565b5f82601f830112615286575f80fd5b815161529961529482615251565b615221565b8181528460208386010111156152ad575f80fd5b611250826020830160208701614ff3565b5f602082840312156152ce575f80fd5b81516001600160401b038111156152e3575f80fd5b61125084828501615277565b602081525f6128866020830184615015565b5f60208284031215615311575f80fd5b5051919050565b5f60208284031215615328575f80fd5b8135612886816150c0565b5f808335601e19843603018112615348575f80fd5b8301803591506001600160401b03821115615361575f80fd5b6020019150600581901b3603821315613814575f80fd5b634e487b7160e01b5f52603260045260245ffd5b5f8235605e198336030181126153a0575f80fd5b9190910192915050565b5f82601f8301126153b9575f80fd5b81356153c761529482615251565b8181528460208386010111156153db575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60608236031215615407575f80fd5b61540f6151d7565b82356001600160401b0380821115615425575f80fd5b615431368387016153aa565b83526020850135915080821115615446575f80fd5b50615453368286016153aa565b602083015250604083013561546781614f15565b604082015292915050565b5f82516153a0818460208701614ff3565b601f8211156104e657805f5260205f20601f840160051c810160208510156154a85750805b601f840160051c820191505b81811015611cda575f81556001016154b4565b81516001600160401b038111156154e0576154e06151c3565b6154f4816154ee845461517d565b84615483565b602080601f831160018114615527575f84156155105750858301515b5f19600386901b1c1916600185901b178555611387565b5f85815260208120601f198616915b8281101561555557888601518255948401946001909101908401615536565b508582101561557257878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561112357611123615582565b5f63ffffffff8083168181036155c1576155c1615582565b6001019392505050565b5f808335601e198436030181126155e0575f80fd5b83016020810192503590506001600160401b038111156155fe575f80fd5b803603821315613814575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f8383855260208086019550808560051b830101845f5b878110156156e957848303601f19018952813536889003605e19018112615670575f80fd5b8701606061567e82806155cb565b82875261568e838801828461560c565b9250505061569e868301836155cb565b868303888801526156b083828461560c565b9250505060408083013592506156c583614f15565b6001600160401b03929092169490910193909352978301979083019060010161564b565b5090979650505050505050565b6020815281356020820152602082013560408201525f604083013561571a816150c0565b6001600160a01b031660608381019190915283013536849003601e19018112615741575f80fd5b83016020810190356001600160401b0381111561575c575f80fd5b8060051b360382131561576d575f80fd5b6080808501526139db60a085018284615634565b6001600160401b038181168382160190808211156124e2576124e2615582565b6001600160401b038281168282160390808211156124e2576124e2615582565b9687526001600160401b03958616602088015293851660408701529184166060860152909216608084015260a083019190915260c082015260e00190565b6001600160401b0381811683821602808216919082811461582257615822615582565b505092915050565b634e487b7160e01b5f52601260045260245ffd5b5f6001600160401b03808416806158575761585761582a565b92169190910492915050565b5f826158715761587161582a565b500490565b602081016004831061588a5761588a614fcb565b91905290565b5f805f606084860312156158a2575f80fd5b8351925060208401516158b481614f15565b60408501519092506158c581614f15565b809150509250925092565b808202811582820484141761112357611123615582565b5f80604083850312156158f8575f80fd5b82519150602083015161510781614e05565b5f60208083525f845461591c8161517d565b806020870152604060018084165f811461593d576001811461595957615986565b60ff19851660408a0152604084151560051b8a01019550615986565b895f5260205f205f5b8581101561597d5781548b8201860152908301908801615962565b8a016040019650505b509398975050505050505050565b818382375f9101908152919050565b5f80604083850312156159b4575f80fd5b82516001600160401b03808211156159ca575f80fd5b90840190606082870312156159dd575f80fd5b6159e56151d7565b8251815260208301516159f7816150c0565b6020820152604083015182811115615a0d575f80fd5b615a1988828601615277565b604083015250809450505050602083015161510781614e05565b5f8060408385031215615a44575f80fd5b82519150602083015161510781614f15565b8181038181111561112357611123615582565b5f8235603e198336030181126153a0575f80fd5b5f808335601e19843603018112615a92575f80fd5b8301803591506001600160401b03821115615aab575f80fd5b602001915036819003821315613814575f80fd5b602081525f61125060208301848661560c565b5f60408236031215615ae2575f80fd5b615aea6151ff565b615af383614e12565b81526020808401356001600160401b0380821115615b0f575f80fd5b9085019036601f830112615b21575f80fd5b813581811115615b3357615b336151c3565b8060051b9150615b44848301615221565b8181529183018401918481019036841115615b5d575f80fd5b938501935b83851015615b875784359250615b77836150c0565b8282529385019390850190615b62565b94860194909452509295945050505050565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b808310156127ff5784516001600160a01b03168252938301936001929092019190830190615bcb565b60208152815160208201525f602083015160e06040840152615c1a610100840182615015565b90506040840151601f1980858403016060860152615c388383615015565b92506001600160401b03606087015116608086015260808601519150808584030160a0860152615c688383615b99565b925060a08601519150808584030160c086015250615c868282615b99565b91505060c0840151615ca360e08501826001600160401b03169052565b509392505050565b5f8060408385031215615cbc575f80fd5b8251915060208301516001600160401b03811115615cd8575f80fd5b615ce485828601615277565b9150509250929050565b5f60208284031215615cfe575f80fd5b61288682614f04565b5f60208284031215615d17575f80fd5b813560ff81168114612886575f80fd5b5f6001600160401b038083168181036155c1576155c1615582565b5f60208284031215615d52575f80fd5b815161288681614e0556fe4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d056e5bdfcce15e53c3406ea67bfce37dcd26f5152d5492824e43fd5e3c8ac5ab004317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d00e92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00e92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb07a2646970667358221220c29e76ccadf773954d1a11079de0204fa05783e932cfc5408ce23c30f8f60c5e64736f6c63430008190033",
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
