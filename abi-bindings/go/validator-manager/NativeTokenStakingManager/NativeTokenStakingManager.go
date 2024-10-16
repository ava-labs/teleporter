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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADDRESS_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BLS_PUBLIC_KEY_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_CHURN_PERCENTAGE_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_REGISTRATION_EXPIRY_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NATIVE_MINTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINativeMinter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POS_VALIDATOR_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"P_CHAIN_BLOCKCHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPTIME_REWARDS_THRESHOLD_PERCENTAGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimDelegationFees\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimValidationRewards\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeEndValidation\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitializeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structValidator\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"startingWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"endedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWeight\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structPoSValidatorManagerSettings\",\"components\":[{\"name\":\"baseSettings\",\"type\":\"tuple\",\"internalType\":\"structValidatorManagerSettings\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"churnPeriodSeconds\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"minimumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rewardCalculator\",\"type\":\"address\",\"internalType\":\"contractIRewardCalculator\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeDelegatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initializeEndDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeValidatorRegistration\",\"inputs\":[{\"name\":\"registrationInput\",\"type\":\"tuple\",\"internalType\":\"structValidatorRegistrationInput\",\"components\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"disableOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initializeValidatorSet\",\"inputs\":[{\"name\":\"subnetConversionData\",\"type\":\"tuple\",\"internalType\":\"structSubnetConversionData\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialValidators\",\"type\":\"tuple[]\",\"internalType\":\"structInitialValidator[]\",\"components\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredValidators\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resendEndValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendRegisterValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendUpdateDelegation\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitUptimeProof\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"valueToWeight\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"weightToValue\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"DelegationEnded\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewards\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorAdded\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"delegatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"delegatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRegistered\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRemovalInitialized\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitialValidatorCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UptimeUpdated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"uptime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodEnded\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumValidatorStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodRegistered\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationRewardsClaimed\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reward\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRemovalInitialized\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorWeightUpdate\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DelegatorIneligibleForRewards\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSKeyLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidBLSPublicKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCodecID\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationFee\",\"inputs\":[{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationID\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDelegatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitializationStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMaximumChurnPercentage\",\"inputs\":[{\"name\":\"maximumChurnPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMinStakeDuration\",\"inputs\":[{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNodeID\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPChainOwnerThreshold\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addressesLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRegistrationExpiry\",\"inputs\":[{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeAmount\",\"inputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeMultiplier\",\"inputs\":[{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidSubnetConversionID\",\"inputs\":[{\"name\":\"encodedSubnetConversionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedSubnetConversionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidationID\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorManagerAddress\",\"inputs\":[{\"name\":\"validatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorManagerBlockchainID\",\"inputs\":[{\"name\":\"blockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWarpOriginSenderAddress\",\"inputs\":[{\"name\":\"senderAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpSourceChainID\",\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MaxChurnRateExceeded\",\"inputs\":[{\"name\":\"churnAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MaxWeightExceeded\",\"inputs\":[{\"name\":\"newValidatorWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MinStakeDurationNotPassed\",\"inputs\":[{\"name\":\"endTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NodeAlreadyRegistered\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PChainOwnerAddressesNotSorted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnexpectedRegistrationStatus\",\"inputs\":[{\"name\":\"validRegistration\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"ValidatorIneligibleForRewards\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotPoS\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516160eb3803806160eb83398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b615f9e8061014d5f395ff3fe6080604052600436106101fc575f3560e01c8063732214f811610113578063b771b3bc1161009d578063c599e24f1161006d578063c599e24f146105cb578063c974d1b6146105de578063d5f20ff6146105f2578063df93d8de1461061e578063fd7ac5e714610634575f80fd5b8063b771b3bc14610553578063ba3a4b971461056d578063bc5fbfec1461058c578063bee0a03f146105ac575f80fd5b806393e24598116100e357806393e24598146104ce578063a3a65e48146104ed578063af2f5feb1461050c578063afb980961461051f578063afba878a1461053f575f80fd5b8063732214f81461046957806376f786211461047c57806380dd672f1461049b5780638280a25a146104ba575f80fd5b80632e2194d811610194578063467ef06f11610164578063467ef06f146103b657806360305d62146103d557806360ad7784146103fe578063620658561461041d57806366435abf1461044a575f80fd5b80632e2194d814610306578063329c3e121461033d57806335455ded1461036f5780633a1cfff614610397575f80fd5b806320d91b7a116101cf57806320d91b7a1461028a57806320e55565146102a957806325e1c776146102c85780632893d077146102e7575f80fd5b80630118acc4146102005780630322ed9814610221578063151d30d1146102405780631ec447241461026b575b5f80fd5b34801561020b575f80fd5b5061021f61021a366004615212565b610653565b005b34801561022c575f80fd5b5061021f61023b36600461524d565b610665565b34801561024b575f80fd5b50610254600a81565b60405160ff90911681526020015b60405180910390f35b348015610276575f80fd5b5061021f610285366004615212565b61087c565b348015610295575f80fd5b5061021f6102a4366004615264565b610888565b3480156102b4575f80fd5b5061021f6102c33660046152b2565b610d74565b3480156102d3575f80fd5b5061021f6102e23660046152c9565b610e51565b3480156102f2575f80fd5b5061021f6103013660046152c9565b610ec5565b348015610311575f80fd5b5061032561032036600461524d565b61120c565b6040516001600160401b039091168152602001610262565b348015610348575f80fd5b506103576001600160991b0181565b6040516001600160a01b039091168152602001610262565b34801561037a575f80fd5b5061038461271081565b60405161ffff9091168152602001610262565b3480156103a2575f80fd5b5061021f6103b1366004615212565b611222565b3480156103c1575f80fd5b5061021f6103d03660046152ea565b61122e565b3480156103e0575f80fd5b506103e9601481565b60405163ffffffff9091168152602001610262565b348015610409575f80fd5b5061021f6104183660046152c9565b61131c565b348015610428575f80fd5b5061043c610437366004615319565b611587565b604051908152602001610262565b348015610455575f80fd5b5061032561046436600461524d565b6115a0565b348015610474575f80fd5b5061043c5f81565b348015610487575f80fd5b5061021f610496366004615212565b6115b4565b3480156104a6575f80fd5b5061021f6104b53660046152c9565b6115c1565b3480156104c5575f80fd5b50610254603081565b3480156104d9575f80fd5b5061021f6104e836600461524d565b611795565b3480156104f8575f80fd5b5061021f6105073660046152ea565b61183f565b61043c61051a366004615343565b6119c8565b34801561052a575f80fd5b5061043c5f80516020615f1283398151915281565b34801561054a575f80fd5b50610254605081565b34801561055e575f80fd5b506103576005600160991b0181565b348015610578575f80fd5b5061021f61058736600461524d565b6119fc565b348015610597575f80fd5b5061043c5f80516020615f3283398151915281565b3480156105b7575f80fd5b5061021f6105c636600461524d565b611be5565b61043c6105d936600461524d565b611d22565b3480156105e9575f80fd5b50610254601481565b3480156105fd575f80fd5b5061061161060c36600461524d565b611d53565b604051610262919061540c565b348015610629575f80fd5b506103256202a30081565b34801561063f575f80fd5b5061043c61064e36600461548c565b611ea2565b6106608383836001611efd565b505050565b5f8181525f80516020615f728339815191526020526040808220815160e0810190925280545f80516020615f3283398151915293929190829060ff1660058111156106b2576106b2615397565b60058111156106c3576106c3615397565b81526020016001820180546106d7906154f7565b80601f0160208091040260200160405190810160405280929190818152602001828054610703906154f7565b801561074e5780601f106107255761010080835404028352916020019161074e565b820191905f5260205f20905b81548152906001019060200180831161073157829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b909104811660808301526003928301541660a090910152909150815160058111156107b9576107b9615397565b146107f5575f8381526007830160205260409081902054905163170cc93360e21b81526107ec9160ff1690600401615529565b60405180910390fd5b6005600160991b016001600160a01b031663ee5b48eb61081a8584606001515f6121a0565b6040518263ffffffff1660e01b81526004016108369190615537565b6020604051808303815f875af1158015610852573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108769190615549565b50505050565b6106608383835f611efd565b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb09545f80516020615f328339815191529060ff16156108da57604051637fab81e560e01b815260040160405180910390fd5b6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa15801561091d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109419190615549565b83602001351461096a576040516372b0a7e760e11b8152602084013560048201526024016107ec565b3061097b6060850160408601615574565b6001600160a01b0316146109be576109996060840160408501615574565b604051632f88120d60e21b81526001600160a01b0390911660048201526024016107ec565b5f6109cc606085018561558f565b905090505f805b828163ffffffff161015610cb4575f6109ef606088018861558f565b8363ffffffff16818110610a0557610a056155d4565b9050602002810190610a1791906155e8565b610a209061570c565b80516040519192505f916008880191610a3891615785565b90815260200160405180910390205414610a6857805160405163a41f772f60e01b81526107ec9190600401615537565b5f6002885f013584604051602001610a9792919091825260e01b6001600160e01b031916602082015260240190565b60408051601f1981840301815290829052610ab191615785565b602060405180830381855afa158015610acc573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610aef9190615549565b90508086600801835f0151604051610b079190615785565b90815260408051602092819003830181209390935560e0830181526002835284518284015284810180516001600160401b03908116858401525f60608601819052915181166080860152421660a085015260c0840181905284815260078a01909252902081518154829060ff19166001836005811115610b8957610b89615397565b021790555060208201516001820190610ba290826157da565b506040828101516002830180546060860151608087015160a08801516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909301516003909201805467ffffffffffffffff191692841692909217909155830151610c479116856158a9565b8251604051919550610c5891615785565b60408051918290038220908401516001600160401b031682529082907f9d47fef9da077661546e646d61830bfcbda90506c2e5eed38195e82c4eb1cbdf9060200160405180910390a3505080610cad906158bc565b90506109d3565b50600483018190555f610cd2610cc9866121ef565b60400151612305565b90505f610cde87612492565b90505f600282604051610cf19190615785565b602060405180830381855afa158015610d0c573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610d2f9190615549565b9050828114610d5b57604051631872fc8d60e01b815260048101829052602481018490526044016107ec565b5050506009909201805460ff1916600117905550505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff1680610dbd575080546001600160401b03808416911610155b15610ddb5760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b178155610e058361266c565b805460ff60401b191681556040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050565b610e5a8261267d565b610e7a576040516330efa98b60e01b8152600481018390526024016107ec565b5f610e8483611d53565b5190506002816005811115610e9b57610e9b615397565b14610ebb578060405163170cc93360e21b81526004016107ec9190615529565b61087683836126b8565b610ecd612932565b5f80516020615f128339815191525f610ee584611d53565b9050600281516005811115610efc57610efc615397565b14610f1d57805160405163170cc93360e21b81526107ec9190600401615529565b610f268461267d565b610f46576040516330efa98b60e01b8152600481018590526024016107ec565b5f8481526004830160205260409020546001600160a01b03163314610f8c57335b604051636e2ccd7560e11b81526001600160a01b0390911660048201526024016107ec565b5f84815260048301602052604090205460a08201514291610fbe91600160b01b9091046001600160401b0316906158de565b6001600160401b0316816001600160401b03161015610ffb5760405163fb6ce63f60e01b81526001600160401b03821660048201526024016107ec565b5f61100686866126b8565b5f8781526004860160205260408120600101549192509061103790600160401b90046001600160401b0316836158fe565b5f888152600487016020526040812060010154919250600160801b9091046001600160401b03169081900361106d575060a08401515b611078828286612969565b61109857604051635bff683f60e11b8152600481018990526024016107ec565b600386015460408601515f916001600160a01b03169063778c06b5906110bd90611587565b848589885f806040518863ffffffff1660e01b81526004016110e5979695949392919061591e565b602060405180830381865afa158015611100573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111249190615549565b90506064605061113484886158fe565b61113e919061595c565b611148919061599b565b5f8a8152600489016020526040902060018101805477ffffffffffffffffffffffffffffffff00000000000000001916600160401b6001600160401b039485160267ffffffffffffffff60801b191617600160801b9389169390930292909217909155546111bf906001600160a01b0316826129b4565b60405181908a907f69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a905f90a35050505050505061120860015f80516020615f5283398151915255565b5050565b5f61121c64e8d4a51000836159c0565b92915050565b6106608383835f612a25565b611236612932565b5f80516020615f128339815191525f8061124f84612c86565b9150915061125c8261267d565b61126857505050611303565b5f82815260048085016020526040909120546001600160a01b0316908251600581111561129757611297615397565b036112e8575f838152600785016020526040812080549190556112ba82826129b4565b604051819085907f69eb3dd96029877a7561d5b5076810e0b53b223c9ab46cba8efd028ebf08fe9a905f90a3505b6112fe816112f98460400151611587565b612fc4565b505050505b61131960015f80516020615f5283398151915255565b50565b5f8281525f80516020615ef28339815191526020526040808220815160e0810190925280545f80516020615f1283398151915293929190829060ff16600381111561136957611369615397565b600381111561137a5761137a615397565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f6113f082611d53565b905060018351600381111561140757611407615397565b14611428578251604051633b0d540d60e21b81526107ec91906004016159d3565b60048151600581111561143d5761143d615397565b036114535761144b86612fd7565b505050505050565b5f8061146a611461886121ef565b604001516131be565b509150915081841461149757846040015160405163089938b360e11b81526004016107ec91815260200190565b806001600160401b031683606001516001600160401b031610806114d05750806001600160401b03168560a001516001600160401b0316115b156114f957604051632e19bc2d60e11b81526001600160401b03821660048201526024016107ec565b5f888152600587016020908152604091829020805460ff1916600290811782550180546001600160401b034216600160401b81026fffffffffffffffff00000000000000001990921691909117909155915191825285918a917f047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6910160405180910390a35050505050505050565b5f61121c6001600160401b03831664e8d4a510006159ed565b5f6115aa82611d53565b6080015192915050565b6106608383836001612a25565b6115c9612932565b5f8281525f80516020615ef28339815191526020526040808220815160e0810190925280545f80516020615f1283398151915293929190829060ff16600381111561161657611616615397565b600381111561162757611627615397565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015290506003815160038111156116a0576116a0615397565b146116c1578051604051633b0d540d60e21b81526107ec91906004016159d3565b60046116d08260400151611d53565b5160058111156116e2576116e2615397565b14611774575f6116f1846121ef565b90505f8061170283604001516131be565b50915091508184604001511461172e5760405163089938b360e11b8152600481018390526024016107ec565b806001600160401b03168460c001516001600160401b0316111561177057604051632e19bc2d60e11b81526001600160401b03821660048201526024016107ec565b5050505b61177d84612fd7565b505061120860015f80516020615f5283398151915255565b5f80516020615f128339815191525f6117ad83611d53565b51905060048160058111156117c4576117c4615397565b146117e4578060405163170cc93360e21b81526004016107ec9190615529565b5f8381526004830160205260409020546001600160a01b031633146118095733610f67565b5f83815260078301602090815260408083208054908490556004860190925290912054610876906001600160a01b0316826129b4565b5f80516020615f328339815191525f8061186461185b856121ef565b60400151613414565b915091508061188a57604051632d07135360e01b815281151560048201526024016107ec565b5f828152600684016020526040902080546118a4906154f7565b90505f036118c85760405163089938b360e11b8152600481018390526024016107ec565b60015f83815260078501602052604090205460ff1660058111156118ee576118ee615397565b14611921575f8281526007840160205260409081902054905163170cc93360e21b81526107ec9160ff1690600401615529565b5f828152600684016020526040812061193991615166565b5f828152600784016020908152604091829020805460ff1916600290811782550180546001600160401b0342818116600160c01b026001600160c01b0390931692909217928390558451600160801b9093041682529181019190915283917ff8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568910160405180910390a250505050565b5f6119d1612932565b6119dd848484346135d0565b90506119f560015f80516020615f5283398151915255565b9392505050565b5f8181525f80516020615ef28339815191526020526040808220815160e0810190925280545f80516020615f1283398151915293929190829060ff166003811115611a4957611a49615397565b6003811115611a5a57611a5a615397565b8152815461010090046001600160a01b0316602082015260018083015460408301526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015290915081516003811115611ad357611ad3615397565b14158015611af45750600381516003811115611af157611af1615397565b14155b15611b15578051604051633b0d540d60e21b81526107ec91906004016159d3565b5f611b238260400151611d53565b905080606001516001600160401b03165f03611b55576040516339b894f960e21b8152600481018590526024016107ec565b6005600160991b016001600160a01b031663ee5b48eb611b828460400151846060015185608001516121a0565b6040518263ffffffff1660e01b8152600401611b9e9190615537565b6020604051808303815f875af1158015611bba573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611bde9190615549565b5050505050565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb066020526040902080545f80516020615f328339815191529190611c2c906154f7565b90505f03611c505760405163089938b360e11b8152600481018390526024016107ec565b60015f83815260078301602052604090205460ff166005811115611c7657611c76615397565b14611ca9575f8281526007820160205260409081902054905163170cc93360e21b81526107ec9160ff1690600401615529565b5f82815260068201602052604090819020905163ee5b48eb60e01b81526005600160991b019163ee5b48eb91611ce29190600401615a04565b6020604051808303815f875af1158015611cfe573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106609190615549565b5f611d2b612932565b611d368233346137e3565b9050611d4e60015f80516020615f5283398151915255565b919050565b611d5b61519d565b5f8281525f80516020615f72833981519152602052604090819020815160e0810190925280545f80516020615f32833981519152929190829060ff166005811115611da857611da8615397565b6005811115611db957611db9615397565b8152602001600182018054611dcd906154f7565b80601f0160208091040260200160405190810160405280929190818152602001828054611df9906154f7565b8015611e445780601f10611e1b57610100808354040283529160200191611e44565b820191905f5260205f20905b815481529060010190602001808311611e2757829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b9091048116608083015260039092015490911660a0909101529392505050565b6040515f905f80516020615f32833981519152907fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0890611ee59086908690615a8e565b90815260200160405180910390205491505092915050565b5f8481525f80516020615ef28339815191526020526040808220815160e0810190925280545f80516020615f1283398151915293929190829060ff166003811115611f4a57611f4a615397565b6003811115611f5b57611f5b615397565b8152815461010090046001600160a01b0316602082015260018201546040808301919091526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101528101519091505f611fd182611d53565b9050600283516003811115611fe857611fe8615397565b14612009578251604051633b0d540d60e21b81526107ec91906004016159d3565b60208301516001600160a01b031633146120235733610f67565b60028151600581111561203857612038615397565b0361213957861561204f5761204d82876126b8565b505b5f8881526005850160205260409020805460ff191660031790556060830151608082015161208891849161208391906158fe565b613ab8565b505f898152600586016020526040812060020180546001600160401b03909316600160c01b026001600160c01b03909316929092179091556120c984613c0f565b90508580156120d6575080155b156120f757604051631036cf9160e11b8152600481018a90526024016107ec565b5f8981526006860160205260408082208390555184918b917f366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed579190a350612196565b60048151600581111561214e5761214e615397565b0361217a5761215c83613c0f565b5f89815260068601602052604090205561217588612fd7565b612196565b805160405163170cc93360e21b81526107ec9190600401615529565b5050505050505050565b604080515f6020820152600360e01b602282015260268101949094526001600160c01b031960c093841b811660468601529190921b16604e830152805180830360360181526056909201905290565b60408051606080820183525f8083526020830152918101919091526040516306f8253560e41b815263ffffffff831660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa158015612253573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261227a9190810190615aa8565b915091508061229c57604051636b2f19e960e01b815260040160405180910390fd5b8151156122c2578151604051636ba589a560e01b815260048101919091526024016107ec565b60208201516001600160a01b0316156122fe576020820151604051624de75d60e31b81526001600160a01b0390911660048201526024016107ec565b5092915050565b5f815160261461233a57815160405163cc92daa160e01b815263ffffffff9091166004820152602660248201526044016107ec565b5f805b600281101561238957612351816001615b7a565b61235c9060086159ed565b61ffff16848281518110612372576123726155d4565b016020015160f81c901b919091179060010161233d565b5061ffff8116156123b35760405163407b587360e01b815261ffff821660048201526024016107ec565b5f805b600481101561240e576123ca816003615b7a565b6123d59060086159ed565b63ffffffff16856123e78360026158a9565b815181106123f7576123f76155d4565b016020015160f81c901b91909117906001016123b6565b5063ffffffff81161561243457604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156124895761244b81601f615b7a565b6124569060086159ed565b866124628360066158a9565b81518110612472576124726155d4565b016020015160f81c901b9190911790600101612437565b50949350505050565b60605f808335602085013560146124ae87870160408901615574565b6124bb606089018961558f565b60405160f09790971b6001600160f01b0319166020880152602287019590955250604285019290925260e090811b6001600160e01b0319908116606286015260609290921b6bffffffffffffffffffffffff191660668501529190911b16607a820152607e0160405160208183030381529060405290505f5b612541606085018561558f565b90508110156122fe5781612558606086018661558f565b83818110612568576125686155d4565b905060200281019061257a91906155e8565b6125849080615b8d565b9050612593606087018761558f565b848181106125a3576125a36155d4565b90506020028101906125b591906155e8565b6125bf9080615b8d565b6125cc606089018961558f565b868181106125dc576125dc6155d4565b90506020028101906125ee91906155e8565b6125fc906020810190615b8d565b61260960608b018b61558f565b88818110612619576126196155d4565b905060200281019061262b91906155e8565b61263c906060810190604001615319565b6040516020016126529796959493929190615bcf565b60408051601f198184030181529190529150600101612534565b612674613dd8565b61131981613e23565b5f9081527f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0460205260409020546001600160a01b0316151590565b6040516306f8253560e41b815263ffffffff821660048201525f90819081906005600160991b0190636f825350906024015f60405180830381865afa158015612703573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261272a9190810190615aa8565b915091508061274c57604051636b2f19e960e01b815260040160405180910390fd5b6005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa15801561278f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127b39190615549565b8251146127d9578151604051636ba589a560e01b815260048101919091526024016107ec565b60208201516001600160a01b031615612815576020820151604051624de75d60e31b81526001600160a01b0390911660048201526024016107ec565b5f806128248460400151613e91565b9150915081871461284b5760405163089938b360e11b8152600481018890526024016107ec565b5f8781527f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d0460205260409020600101545f80516020615f12833981519152906001600160401b039081169083161115612908575f888152600482016020908152604091829020600101805467ffffffffffffffff19166001600160401b038616908117909155915191825289917fec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435910160405180910390a2612927565b5f8881526004820160205260409020600101546001600160401b031691505b509695505050505050565b5f80516020615f5283398151915280546001190161296357604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f605061297684846158fe565b612980919061595c565b6001600160401b031661299485606461595c565b6001600160401b031610156129aa57505f6119f5565b5060019392505050565b6040516327ad555d60e11b81526001600160a01b0383166004820152602481018290526001600160991b0190634f5aaaba906044015f604051808303815f87803b158015612a00575f80fd5b505af115801561144b573d5f803e3d5ffd5b60015f80516020615f5283398151915255565b5f80516020615f128339815191525f612a3d86614082565b9050612a488661267d565b612a53575050610876565b5f8681526004830160205260409020546001600160a01b03163314612a785733610f67565b5f86815260048301602052604090205460a0820151612aa791600160b01b90046001600160401b0316906158de565b6001600160401b03168160c001516001600160401b03161015612aee5760c081015160405163fb6ce63f60e01b81526001600160401b0390911660048201526024016107ec565b5f8515612b0657612aff87866126b8565b9050612b24565b505f8681526004830160205260409020600101546001600160401b03165b5f878152600484016020526040812060010154612b5190600160401b90046001600160401b0316836158fe565b5f898152600486016020526040812060010154919250600160801b9091046001600160401b031690819003612b87575060a08301515b858015612ba05750612b9e82828660c00151612969565b155b15612bc157604051635bff683f60e11b8152600481018a90526024016107ec565b600385015460408501515f916001600160a01b03169063778c06b590612be690611587565b84858960c00151885f806040518863ffffffff1660e01b8152600401612c12979695949392919061591e565b602060405180830381865afa158015612c2d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c519190615549565b905080866007015f8c81526020019081526020015f205f828254612c7591906158a9565b909155505050505050505050505050565b5f612c8f61519d565b5f80516020615f328339815191525f80612cab61185b876121ef565b915091508015612cd257604051632d07135360e01b815281151560048201526024016107ec565b5f828152600784016020526040808220815160e081019092528054829060ff166005811115612d0357612d03615397565b6005811115612d1457612d14615397565b8152602001600182018054612d28906154f7565b80601f0160208091040260200160405190810160405280929190818152602001828054612d54906154f7565b8015612d9f5780601f10612d7657610100808354040283529160200191612d9f565b820191905f5260205f20905b815481529060010190602001808311612d8257829003601f168201915b505050918352505060028201546001600160401b038082166020840152600160401b820481166040840152600160801b820481166060840152600160c01b909104811660808301526003928301541660a09091015290915081516005811115612e0a57612e0a615397565b14158015612e2b5750600181516005811115612e2857612e28615397565b14155b15612e4c57805160405163170cc93360e21b81526107ec9190600401615529565b600381516005811115612e6157612e61615397565b03612e6f5760048152612e74565b600581525b836008018160200151604051612e8a9190615785565b90815260408051602092819003830190205f908190558581526007870190925290208151815483929190829060ff19166001836005811115612ece57612ece615397565b021790555060208201516001820190612ee790826157da565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909201516003909101805467ffffffffffffffff19169190921617905580516005811115612f8d57612f8d615397565b60405184907f1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16905f90a39196919550909350505050565b6112086001600160a01b03831682614366565b5f8181525f80516020615ef28339815191526020526040808220815160e0810190925280545f80516020615f1283398151915293929190829060ff16600381111561302457613024615397565b600381111561303557613035615397565b8152815461010090046001600160a01b03166020808301919091526001808401546040808501919091526002948501546001600160401b038082166060870152600160401b820481166080870152600160801b8204811660a0870152600160c01b9091041660c090940193909352848301515f89815260058901845284812080546001600160a81b03191681559283018190559190940181905560068701909152908120805490829055929350909190808215613163575f8481526004870160205260409020546127109061311590600160a01b900461ffff16856159ed565b61311f91906159c0565b915081866007015f8681526020019081526020015f205f82825461314391906158a9565b9091555061315390508284615b7a565b90506131638560200151826129b4565b61317885602001516112f98760600151611587565b6040805182815260208101849052859189917f8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993910160405180910390a350505050505050565b5f805f83516036146131f557835160405163cc92daa160e01b815263ffffffff9091166004820152603660248201526044016107ec565b5f805b60028110156132445761320c816001615b7a565b6132179060086159ed565b61ffff1686828151811061322d5761322d6155d4565b016020015160f81c901b91909117906001016131f8565b5061ffff81161561326e5760405163407b587360e01b815261ffff821660048201526024016107ec565b5f805b60048110156132c957613285816003615b7a565b6132909060086159ed565b63ffffffff16876132a28360026158a9565b815181106132b2576132b26155d4565b016020015160f81c901b9190911790600101613271565b5063ffffffff81166003146132f157604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156133465761330881601f615b7a565b6133139060086159ed565b8861331f8360066158a9565b8151811061332f5761332f6155d4565b016020015160f81c901b91909117906001016132f4565b505f805b60088110156133a55761335e816007615b7a565b6133699060086159ed565b6001600160401b03168961337e8360266158a9565b8151811061338e5761338e6155d4565b016020015160f81c901b919091179060010161334a565b505f805b6008811015613404576133bd816007615b7a565b6133c89060086159ed565b6001600160401b03168a6133dd83602e6158a9565b815181106133ed576133ed6155d4565b016020015160f81c901b91909117906001016133a9565b5091989097509095509350505050565b5f80825160271461344a57825160405163cc92daa160e01b815263ffffffff9091166004820152602760248201526044016107ec565b5f805b600281101561349957613461816001615b7a565b61346c9060086159ed565b61ffff16858281518110613482576134826155d4565b016020015160f81c901b919091179060010161344d565b5061ffff8116156134c35760405163407b587360e01b815261ffff821660048201526024016107ec565b5f805b600481101561351e576134da816003615b7a565b6134e59060086159ed565b63ffffffff16866134f78360026158a9565b81518110613507576135076155d4565b016020015160f81c901b91909117906001016134c6565b5063ffffffff811660021461354657604051635b60892f60e01b815260040160405180910390fd5b5f805b602081101561359b5761355d81601f615b7a565b6135689060086159ed565b876135748360066158a9565b81518110613584576135846155d4565b016020015160f81c901b9190911790600101613549565b505f866026815181106135b0576135b06155d4565b016020015191976001600160f81b03199092161515965090945050505050565b7f4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d02545f905f80516020615f1283398151915290600160401b900461ffff9081169086161080613624575061271061ffff8616115b1561364857604051635f12e6c360e11b815261ffff861660048201526024016107ec565b60028101546001600160401b039081169085161015613684576040516202a06d60e11b81526001600160401b03851660048201526024016107ec565b80548310806136965750806001015483115b156136b75760405163222d164360e21b8152600481018490526024016107ec565b825f6136c28261120c565b90505f6136cf89836143f9565b90506040518060c001604052806136e33390565b6001600160a01b03908116825261ffff808c166020808501919091526001600160401b03808d166040808701919091525f6060808801829052608080890183905260a09889018390528a83526004909d0185529082902088518154958a0151938a01518516600160b01b0267ffffffffffffffff60b01b1994909716600160a01b026001600160b01b03199096169716969096179390931716929092178355840151600190920180549885015194909301518116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b031990991692909116919091179690961791909116949094179093555090915050949350505050565b5f5f80516020615f12833981519152816137fc8461120c565b90505f61380887611d53565b90506138138761267d565b613833576040516330efa98b60e01b8152600481018890526024016107ec565b60028151600581111561384857613848615397565b1461386957805160405163170cc93360e21b81526107ec9190600401615529565b5f82826080015161387a91906158de565b905083600201600a9054906101000a90046001600160401b031682604001516138a3919061595c565b6001600160401b0316816001600160401b031611156138e057604051636d51fe0560e11b81526001600160401b03821660048201526024016107ec565b5f806138ec8a84613ab8565b915091505f8a8360405160200161391a92919091825260c01b6001600160c01b031916602082015260280190565b60408051601f19818403018152828252805160209091012060e08301909152915080600181526001600160a01b038c1660208083019190915260408083018f90526001600160401b03808b1660608501525f6080850181905290881660a085015260c090930183905284835260058b01909152902081518154829060ff191660018360038111156139ad576139ad615397565b02179055506020828101518254610100600160a81b0319166101006001600160a01b039283160217835560408085015160018501556060808601516002909501805460808089015160a08a015160c0909a01516001600160401b03998a166001600160801b031990941693909317600160401b918a1691909102176001600160801b0316600160801b998916999099026001600160c01b031698909817600160c01b91881691909102179055815189861681528a861694810194909452938b1690830152918101859052908c16918d9184917fb0024b263bc3a0b728a6edea50a69efa841189f8d32ee8af9d1c2b1a1a223426910160405180910390a49a9950505050505050505050565b5f8281525f80516020615f72833981519152602052604081206002015481905f80516020615f3283398151915290600160801b90046001600160401b0316613b008582614971565b5f613b0a87614b4b565b5f8881526007850160205260408120600201805467ffffffffffffffff60801b1916600160801b6001600160401b038b16021790559091506005600160991b0163ee5b48eb613b5a8a858b6121a0565b6040518263ffffffff1660e01b8152600401613b769190615537565b6020604051808303815f875af1158015613b92573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613bb69190615549565b604080516001600160401b038a811682526020820184905282519394508516928b927f07de5ff35a674a8005e661f3333c907ca6333462808762d19dc7b3abb1a8c1df928290030190a3909450925050505b9250929050565b5f805f80516020615f1283398151915290505f613c2f8460400151611d53565b90505f600382516005811115613c4757613c47615397565b1480613c655750600482516005811115613c6357613c63615397565b145b15613c75575060c0810151613cb2565b600282516005811115613c8a57613c8a615397565b03613c96575042613cb2565b815160405163170cc93360e21b81526107ec9190600401615529565b84608001516001600160401b0316816001600160401b031611613cd957505f949350505050565b6040808601515f90815260048501602052206001015460a0830151613d08916001600160401b03169083612969565b613d1657505f949350505050565b600383015460608601516001600160a01b039091169063778c06b590613d3b90611587565b60a085015160808901516040808b01515f9081526004808b016020528282206001015492516001600160e01b031960e089901b168152613d90969594938a936001600160401b0390911692909182910161591e565b602060405180830381865afa158015613dab573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613dcf9190615549565b95945050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16613e2157604051631afcd79f60e31b815260040160405180910390fd5b565b613e2b613dd8565b613e3481614bc0565b613e3c614bd9565b61131960608201356080830135613e5960c0850160a08601615319565b613e6960e0860160c08701615c38565b613e7a610100870160e08801615c51565b613e8c61012088016101008901615574565b614be9565b5f808251602e14613ec757825160405163cc92daa160e01b815263ffffffff9091166004820152602e60248201526044016107ec565b5f805b6002811015613f1657613ede816001615b7a565b613ee99060086159ed565b61ffff16858281518110613eff57613eff6155d4565b016020015160f81c901b9190911790600101613eca565b5061ffff811615613f405760405163407b587360e01b815261ffff821660048201526024016107ec565b5f805b6004811015613f9b57613f57816003615b7a565b613f629060086159ed565b63ffffffff1686613f748360026158a9565b81518110613f8457613f846155d4565b016020015160f81c901b9190911790600101613f43565b5063ffffffff811615613fc157604051635b60892f60e01b815260040160405180910390fd5b5f805b602081101561401657613fd881601f615b7a565b613fe39060086159ed565b87613fef8360066158a9565b81518110613fff57613fff6155d4565b016020015160f81c901b9190911790600101613fc4565b505f805b60088110156140755761402e816007615b7a565b6140399060086159ed565b6001600160401b03168861404e8360266158a9565b8151811061405e5761405e6155d4565b016020015160f81c901b919091179060010161401a565b5090969095509350505050565b61408a61519d565b5f8281525f80516020615f728339815191526020526040808220815160e0810190925280545f80516020615f3283398151915293929190829060ff1660058111156140d7576140d7615397565b60058111156140e8576140e8615397565b81526020016001820180546140fc906154f7565b80601f0160208091040260200160405190810160405280929190818152602001828054614128906154f7565b80156141735780601f1061414a57610100808354040283529160200191614173565b820191905f5260205f20905b81548152906001019060200180831161415657829003601f168201915b50505091835250506002828101546001600160401b038082166020850152600160401b820481166040850152600160801b820481166060850152600160c01b9091048116608084015260039093015490921660a090910152909150815160058111156141e1576141e1615397565b14614214575f8481526007830160205260409081902054905163170cc93360e21b81526107ec9160ff1690600401615529565b60038152426001600160401b031660c08201525f84815260078301602052604090208151815483929190829060ff1916600183600581111561425857614258615397565b02179055506020820151600182019061427190826157da565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909201516003909101805467ffffffffffffffff1916919092161790555f61430f8582613ab8565b6080840151604080516001600160401b03909216825242602083015291935083925087917f13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67910160405180910390a3509392505050565b804710156143895760405163cd78605960e01b81523060048201526024016107ec565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f81146143d2576040519150601f19603f3d011682016040523d82523d5f602084013e6143d7565b606091505b505090508061066057604051630a12f52160e11b815260040160405180910390fd5b7fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb09545f9060ff1661443d57604051637fab81e560e01b815260040160405180910390fd5b5f80516020615f328339815191524261445c6060860160408701615319565b6001600160401b0316111580614496575061447a6202a300426158a9565b61448a6060860160408701615319565b6001600160401b031610155b156144d0576144ab6060850160408601615319565b604051635879da1360e11b81526001600160401b0390911660048201526024016107ec565b6144e56144e06060860186615c71565b614d18565b6144f56144e06080860186615c71565b60306145046020860186615b8d565b905014614536576145186020850185615b8d565b6040516326475b2f60e11b81526107ec925060040190815260200190565b6145408480615b8d565b90505f0361456d576145528480615b8d565b604051633e08a12560e11b81526004016107ec929190615c85565b5f6008820161457c8680615b8d565b60405161458a929190615a8e565b908152602001604051809103902054146145c3576145a88480615b8d565b60405163a41f772f60e01b81526004016107ec929190615c85565b6145cd835f614971565b6040805160e08101909152815481525f9081906146d990602081016145f28980615b8d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050509082525060209081019061463a908a018a615b8d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050509082525060200161468360608a0160408b01615319565b6001600160401b0316815260200161469e60608a018a615c71565b6146a790615cb3565b81526020016146b960808a018a615c71565b6146c290615cb3565b8152602001876001600160401b0316815250614e81565b5f828152600686016020526040902091935091506146f782826157da565b5081600884016147078880615b8d565b604051614715929190615a8e565b9081526040519081900360200181209190915563ee5b48eb60e01b81525f906005600160991b019063ee5b48eb90614751908590600401615537565b6020604051808303815f875af115801561476d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906147919190615549565b6040805160e0810190915290915080600181526020016147b18980615b8d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052509385525050506001600160401b0389166020808401829052604080850184905260608501929092526080840183905260a0909301829052868252600788019092522081518154829060ff1916600183600581111561484057614840615397565b02179055506020820151600182019061485990826157da565b5060408201516002820180546060850151608086015160a08701516001600160401b039586166001600160801b031990941693909317600160401b92861692909202919091176001600160801b0316600160801b918516919091026001600160c01b031617600160c01b9184169190910217905560c0909201516003909101805467ffffffffffffffff191691909216179055806148f78880615b8d565b604051614905929190615a8e565b6040518091039020847fb77297e3befc691bfc864a81e241f83e2ef722b6e7becaa2ecec250c6d52b430898b60400160208101906149439190615319565b604080516001600160401b0393841681529290911660208301520160405180910390a4509095945050505050565b5f80516020615f328339815191525f6001600160401b0380841690851611156149a55761499e83856158fe565b90506149b2565b6149af84846158fe565b90505b6040805160808101825260028401548082526003850154602083015260048501549282019290925260058401546001600160401b0316606082015242911580614a14575060018401548151614a10916001600160401b0316906158a9565b8210155b15614a3a576001600160401b038316606082015281815260408101516020820152614a59565b8281606001818151614a4c91906158de565b6001600160401b03169052505b6060810151614a6990606461595c565b602082015160018601546001600160401b039290921691614a949190600160401b900460ff166159ed565b1015614ac457606081015160405163dfae880160e01b81526001600160401b0390911660048201526024016107ec565b856001600160401b031681604001818151614adf91906158a9565b9052506040810180516001600160401b0387169190614aff908390615b7a565b905250805160028501556020810151600385015560408101516004850155606001516005909301805467ffffffffffffffff19166001600160401b039094169390931790925550505050565b5f8181525f80516020615f728339815191526020526040812060020180545f80516020615f328339815191529190600890614b9590600160401b90046001600160401b0316615d7a565b91906101000a8154816001600160401b0302191690836001600160401b031602179055915050919050565b614bc8613dd8565b614bd061506e565b61131981615076565b614be1613dd8565b613e2161515e565b614bf1613dd8565b5f80516020615f1283398151915261ffff84161580614c15575061271061ffff8516115b15614c3957604051635f12e6c360e11b815261ffff851660048201526024016107ec565b85871115614c5d5760405163222d164360e21b8152600481018890526024016107ec565b60ff83161580614c705750600a60ff8416115b15614c935760405163170db35960e31b815260ff841660048201526024016107ec565b95865560018601949094556002850180546001600160401b039490941669ffffffffffffffffffff1990941693909317600160401b61ffff93909316929092029190911767ffffffffffffffff60501b191660ff91909116600160501b02179055600390910180546001600160a01b0319166001600160a01b03909216919091179055565b614d2560208201826152ea565b63ffffffff16158015614d455750614d40602082018261558f565b151590505b15614d8c57614d5760208201826152ea565b614d64602083018361558f565b60405163c08a0f1d60e01b815263ffffffff90931660048401526024830152506044016107ec565b614d99602082018261558f565b9050614da860208301836152ea565b63ffffffff161115614dc157614d5760208201826152ea565b60015b614dd1602083018361558f565b905081101561120857614de7602083018361558f565b614df2600184615b7a565b818110614e0157614e016155d4565b9050602002016020810190614e169190615574565b6001600160a01b0316614e2c602084018461558f565b83818110614e3c57614e3c6155d4565b9050602002016020810190614e519190615574565b6001600160a01b03161015614e7957604051630dbc8d5f60e31b815260040160405180910390fd5b600101614dc4565b5f6060826040015151603014614eaa5760405163180ffa0d60e01b815260040160405180910390fd5b82516020808501518051604080880151606089015160808a01518051908701515193515f98614eeb988a986001989297929690959094909390929101615d95565b60405160208183030381529060405290505f5b84608001516020015151811015614f5d57818560800151602001518281518110614f2a57614f2a6155d4565b6020026020010151604051602001614f43929190615e4f565b60408051601f198184030181529190529150600101614efe565b5060a0840151805160209182015151604051614f7d938593929101615e85565b60405160208183030381529060405290505f5b8460a001516020015151811015614fef57818560a00151602001518281518110614fbc57614fbc6155d4565b6020026020010151604051602001614fd5929190615e4f565b60408051601f198184030181529190529150600101614f90565b5060c0840151604051615006918391602001615ec0565b60405160208183030381529060405290506002816040516150279190615785565b602060405180830381855afa158015615042573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906150659190615549565b94909350915050565b613e21613dd8565b61507e613dd8565b80355f80516020615f3283398151915290815560146150a36060840160408501615c51565b60ff1611806150c257506150bd6060830160408401615c51565b60ff16155b156150f6576150d76060830160408401615c51565b604051634a59bbff60e11b815260ff90911660048201526024016107ec565b6151066060830160408401615c51565b60018201805460ff92909216600160401b0260ff60401b199092169190911790556151376040830160208401615319565b600191909101805467ffffffffffffffff19166001600160401b0390921691909117905550565b612a12613dd8565b508054615172906154f7565b5f825580601f10615181575050565b601f0160209004905f5260205f209081019061131991906151da565b6040805160e08101909152805f81526060602082018190525f604083018190529082018190526080820181905260a0820181905260c09091015290565b5b808211156151ee575f81556001016151db565b5090565b8015158114611319575f80fd5b803563ffffffff81168114611d4e575f80fd5b5f805f60608486031215615224575f80fd5b833592506020840135615236816151f2565b9150615244604085016151ff565b90509250925092565b5f6020828403121561525d575f80fd5b5035919050565b5f8060408385031215615275575f80fd5b82356001600160401b0381111561528a575f80fd5b83016080818603121561529b575f80fd5b91506152a9602084016151ff565b90509250929050565b5f61012082840312156152c3575f80fd5b50919050565b5f80604083850312156152da575f80fd5b823591506152a9602084016151ff565b5f602082840312156152fa575f80fd5b6119f5826151ff565b80356001600160401b0381168114611d4e575f80fd5b5f60208284031215615329575f80fd5b6119f582615303565b803561ffff81168114611d4e575f80fd5b5f805f60608486031215615355575f80fd5b83356001600160401b0381111561536a575f80fd5b840160a0818703121561537b575f80fd5b925061538960208501615332565b915061524460408501615303565b634e487b7160e01b5f52602160045260245ffd5b600681106153bb576153bb615397565b9052565b5f5b838110156153d95781810151838201526020016153c1565b50505f910152565b5f81518084526153f88160208601602086016153bf565b601f01601f19169290920160200192915050565b6020815261541e6020820183516153ab565b5f602083015160e060408401526154396101008401826153e1565b905060408401516001600160401b0380821660608601528060608701511660808601528060808701511660a08601528060a08701511660c08601528060c08701511660e086015250508091505092915050565b5f806020838503121561549d575f80fd5b82356001600160401b03808211156154b3575f80fd5b818501915085601f8301126154c6575f80fd5b8135818111156154d4575f80fd5b8660208285010111156154e5575f80fd5b60209290920196919550909350505050565b600181811c9082168061550b57607f821691505b6020821081036152c357634e487b7160e01b5f52602260045260245ffd5b6020810161121c82846153ab565b602081525f6119f560208301846153e1565b5f60208284031215615559575f80fd5b5051919050565b6001600160a01b0381168114611319575f80fd5b5f60208284031215615584575f80fd5b81356119f581615560565b5f808335601e198436030181126155a4575f80fd5b8301803591506001600160401b038211156155bd575f80fd5b6020019150600581901b3603821315613c08575f80fd5b634e487b7160e01b5f52603260045260245ffd5b5f8235605e198336030181126155fc575f80fd5b9190910192915050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b038111828210171561563c5761563c615606565b60405290565b604080519081016001600160401b038111828210171561563c5761563c615606565b604051601f8201601f191681016001600160401b038111828210171561568c5761568c615606565b604052919050565b5f6001600160401b038211156156ac576156ac615606565b50601f01601f191660200190565b5f82601f8301126156c9575f80fd5b81356156dc6156d782615694565b615664565b8181528460208386010111156156f0575f80fd5b816020850160208301375f918101602001919091529392505050565b5f6060823603121561571c575f80fd5b61572461561a565b82356001600160401b038082111561573a575f80fd5b615746368387016156ba565b8352602085013591508082111561575b575f80fd5b50615768368286016156ba565b60208301525061577a60408401615303565b604082015292915050565b5f82516155fc8184602087016153bf565b601f82111561066057805f5260205f20601f840160051c810160208510156157bb5750805b601f840160051c820191505b81811015611bde575f81556001016157c7565b81516001600160401b038111156157f3576157f3615606565b6158078161580184546154f7565b84615796565b602080601f83116001811461583a575f84156158235750858301515b5f19600386901b1c1916600185901b17855561144b565b5f85815260208120601f198616915b8281101561586857888601518255948401946001909101908401615849565b508582101561588557878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561121c5761121c615895565b5f63ffffffff8083168181036158d4576158d4615895565b6001019392505050565b6001600160401b038181168382160190808211156122fe576122fe615895565b6001600160401b038281168282160390808211156122fe576122fe615895565b9687526001600160401b03958616602088015293851660408701529184166060860152909216608084015260a083019190915260c082015260e00190565b6001600160401b0381811683821602808216919082811461597f5761597f615895565b505092915050565b634e487b7160e01b5f52601260045260245ffd5b5f6001600160401b03808416806159b4576159b4615987565b92169190910492915050565b5f826159ce576159ce615987565b500490565b60208101600483106159e7576159e7615397565b91905290565b808202811582820484141761121c5761121c615895565b5f60208083525f8454615a16816154f7565b806020870152604060018084165f8114615a375760018114615a5357615a80565b60ff19851660408a0152604084151560051b8a01019550615a80565b895f5260205f205f5b85811015615a775781548b8201860152908301908801615a5c565b8a016040019650505b509398975050505050505050565b818382375f9101908152919050565b8051611d4e816151f2565b5f8060408385031215615ab9575f80fd5b82516001600160401b0380821115615acf575f80fd5b9084019060608287031215615ae2575f80fd5b615aea61561a565b82518152602080840151615afd81615560565b82820152604084015183811115615b12575f80fd5b80850194505087601f850112615b26575f80fd5b83519250615b366156d784615694565b8381528882858701011115615b49575f80fd5b615b58848383018488016153bf565b80604084015250819550615b6d818801615a9d565b9450505050509250929050565b8181038181111561121c5761121c615895565b5f808335601e19843603018112615ba2575f80fd5b8301803591506001600160401b03821115615bbb575f80fd5b602001915036819003821315613c08575f80fd5b5f8851615be0818460208d016153bf565b60e089901b6001600160e01b031916908301908152868860048301378681019050600481015f8152858782375060c09390931b6001600160c01b0319166004939094019283019390935250600c019695505050505050565b5f60208284031215615c48575f80fd5b6119f582615332565b5f60208284031215615c61575f80fd5b813560ff811681146119f5575f80fd5b5f8235603e198336030181126155fc575f80fd5b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b5f60408236031215615cc3575f80fd5b615ccb615642565b615cd4836151ff565b81526020808401356001600160401b0380821115615cf0575f80fd5b9085019036601f830112615d02575f80fd5b813581811115615d1457615d14615606565b8060051b9150615d25848301615664565b8181529183018401918481019036841115615d3e575f80fd5b938501935b83851015615d685784359250615d5883615560565b8282529385019390850190615d43565b94860194909452509295945050505050565b5f6001600160401b038083168181036158d4576158d4615895565b61ffff60f01b8a60f01b1681525f63ffffffff60e01b808b60e01b166002840152896006840152808960e01b166026840152508651615ddb81602a850160208b016153bf565b865190830190615df281602a840160208b016153bf565b60c087901b6001600160c01b031916602a9290910191820152615e24603282018660e01b6001600160e01b0319169052565b615e3d603682018560e01b6001600160e01b0319169052565b603a019b9a5050505050505050505050565b5f8351615e608184602088016153bf565b60609390931b6bffffffffffffffffffffffff19169190920190815260140192915050565b5f8451615e968184602089016153bf565b6001600160e01b031960e095861b8116919093019081529290931b16600482015260080192915050565b5f8351615ed18184602088016153bf565b60c09390931b6001600160c01b031916919092019081526008019291505056fe4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d054317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d00e92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00e92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb07a164736f6c6343000819000a",
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
