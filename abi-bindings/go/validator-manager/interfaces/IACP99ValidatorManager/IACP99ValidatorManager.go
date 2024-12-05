// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iacp99validatormanager

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

// ConversionData is an auto generated low-level Go binding around an user-defined struct.
type ConversionData struct {
	SubnetID                     [32]byte
	ValidatorManagerBlockchainID [32]byte
	ValidatorManagerAddress      common.Address
	InitialValidators            []InitialValidator
}

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

// ValidatorRegistrationInput is an auto generated low-level Go binding around an user-defined struct.
type ValidatorRegistrationInput struct {
	NodeID                []byte
	BlsPublicKey          []byte
	RegistrationExpiry    uint64
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
}

// IACP99ValidatorManagerMetaData contains all meta data concerning the IACP99ValidatorManager contract.
var IACP99ValidatorManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndValidation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"completeValidatorWeightChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChurnPeriodSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSecurityModule\",\"outputs\":[{\"internalType\":\"contractIACP99SecurityModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"startingWeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"messageNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endedAt\",\"type\":\"uint64\"}],\"internalType\":\"structValidator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"registrationExpiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"internalType\":\"structPChainOwner\",\"name\":\"disableOwner\",\"type\":\"tuple\"}],\"internalType\":\"structValidatorRegistrationInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"initializeValidatorRegistration\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnetID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validatorManagerAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nodeID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"internalType\":\"structInitialValidator[]\",\"name\":\"initialValidators\",\"type\":\"tuple[]\"}],\"internalType\":\"structConversionData\",\"name\":\"conversionData\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"}],\"name\":\"initializeValidatorWeightChange\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IACP99ValidatorManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IACP99ValidatorManagerMetaData.ABI instead.
var IACP99ValidatorManagerABI = IACP99ValidatorManagerMetaData.ABI

// IACP99ValidatorManager is an auto generated Go binding around an Ethereum contract.
type IACP99ValidatorManager struct {
	IACP99ValidatorManagerCaller     // Read-only binding to the contract
	IACP99ValidatorManagerTransactor // Write-only binding to the contract
	IACP99ValidatorManagerFilterer   // Log filterer for contract events
}

// IACP99ValidatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IACP99ValidatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IACP99ValidatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IACP99ValidatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IACP99ValidatorManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IACP99ValidatorManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IACP99ValidatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IACP99ValidatorManagerSession struct {
	Contract     *IACP99ValidatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IACP99ValidatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IACP99ValidatorManagerCallerSession struct {
	Contract *IACP99ValidatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IACP99ValidatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IACP99ValidatorManagerTransactorSession struct {
	Contract     *IACP99ValidatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IACP99ValidatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IACP99ValidatorManagerRaw struct {
	Contract *IACP99ValidatorManager // Generic contract binding to access the raw methods on
}

// IACP99ValidatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IACP99ValidatorManagerCallerRaw struct {
	Contract *IACP99ValidatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IACP99ValidatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IACP99ValidatorManagerTransactorRaw struct {
	Contract *IACP99ValidatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIACP99ValidatorManager creates a new instance of IACP99ValidatorManager, bound to a specific deployed contract.
func NewIACP99ValidatorManager(address common.Address, backend bind.ContractBackend) (*IACP99ValidatorManager, error) {
	contract, err := bindIACP99ValidatorManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IACP99ValidatorManager{IACP99ValidatorManagerCaller: IACP99ValidatorManagerCaller{contract: contract}, IACP99ValidatorManagerTransactor: IACP99ValidatorManagerTransactor{contract: contract}, IACP99ValidatorManagerFilterer: IACP99ValidatorManagerFilterer{contract: contract}}, nil
}

// NewIACP99ValidatorManagerCaller creates a new read-only instance of IACP99ValidatorManager, bound to a specific deployed contract.
func NewIACP99ValidatorManagerCaller(address common.Address, caller bind.ContractCaller) (*IACP99ValidatorManagerCaller, error) {
	contract, err := bindIACP99ValidatorManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IACP99ValidatorManagerCaller{contract: contract}, nil
}

// NewIACP99ValidatorManagerTransactor creates a new write-only instance of IACP99ValidatorManager, bound to a specific deployed contract.
func NewIACP99ValidatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IACP99ValidatorManagerTransactor, error) {
	contract, err := bindIACP99ValidatorManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IACP99ValidatorManagerTransactor{contract: contract}, nil
}

// NewIACP99ValidatorManagerFilterer creates a new log filterer instance of IACP99ValidatorManager, bound to a specific deployed contract.
func NewIACP99ValidatorManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IACP99ValidatorManagerFilterer, error) {
	contract, err := bindIACP99ValidatorManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IACP99ValidatorManagerFilterer{contract: contract}, nil
}

// bindIACP99ValidatorManager binds a generic wrapper to an already deployed contract.
func bindIACP99ValidatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IACP99ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IACP99ValidatorManager *IACP99ValidatorManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IACP99ValidatorManager.Contract.IACP99ValidatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IACP99ValidatorManager *IACP99ValidatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.IACP99ValidatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IACP99ValidatorManager *IACP99ValidatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.IACP99ValidatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IACP99ValidatorManager *IACP99ValidatorManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IACP99ValidatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.contract.Transact(opts, method, params...)
}

// GetChurnPeriodSeconds is a free data retrieval call binding the contract method 0x09c1df66.
//
// Solidity: function getChurnPeriodSeconds() view returns(uint64)
func (_IACP99ValidatorManager *IACP99ValidatorManagerCaller) GetChurnPeriodSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IACP99ValidatorManager.contract.Call(opts, &out, "getChurnPeriodSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetChurnPeriodSeconds is a free data retrieval call binding the contract method 0x09c1df66.
//
// Solidity: function getChurnPeriodSeconds() view returns(uint64)
func (_IACP99ValidatorManager *IACP99ValidatorManagerSession) GetChurnPeriodSeconds() (uint64, error) {
	return _IACP99ValidatorManager.Contract.GetChurnPeriodSeconds(&_IACP99ValidatorManager.CallOpts)
}

// GetChurnPeriodSeconds is a free data retrieval call binding the contract method 0x09c1df66.
//
// Solidity: function getChurnPeriodSeconds() view returns(uint64)
func (_IACP99ValidatorManager *IACP99ValidatorManagerCallerSession) GetChurnPeriodSeconds() (uint64, error) {
	return _IACP99ValidatorManager.Contract.GetChurnPeriodSeconds(&_IACP99ValidatorManager.CallOpts)
}

// GetSecurityModule is a free data retrieval call binding the contract method 0xcdeea3c9.
//
// Solidity: function getSecurityModule() view returns(address)
func (_IACP99ValidatorManager *IACP99ValidatorManagerCaller) GetSecurityModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IACP99ValidatorManager.contract.Call(opts, &out, "getSecurityModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSecurityModule is a free data retrieval call binding the contract method 0xcdeea3c9.
//
// Solidity: function getSecurityModule() view returns(address)
func (_IACP99ValidatorManager *IACP99ValidatorManagerSession) GetSecurityModule() (common.Address, error) {
	return _IACP99ValidatorManager.Contract.GetSecurityModule(&_IACP99ValidatorManager.CallOpts)
}

// GetSecurityModule is a free data retrieval call binding the contract method 0xcdeea3c9.
//
// Solidity: function getSecurityModule() view returns(address)
func (_IACP99ValidatorManager *IACP99ValidatorManagerCallerSession) GetSecurityModule() (common.Address, error) {
	return _IACP99ValidatorManager.Contract.GetSecurityModule(&_IACP99ValidatorManager.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_IACP99ValidatorManager *IACP99ValidatorManagerCaller) GetValidator(opts *bind.CallOpts, validationID [32]byte) (Validator, error) {
	var out []interface{}
	err := _IACP99ValidatorManager.contract.Call(opts, &out, "getValidator", validationID)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_IACP99ValidatorManager *IACP99ValidatorManagerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _IACP99ValidatorManager.Contract.GetValidator(&_IACP99ValidatorManager.CallOpts, validationID)
}

// GetValidator is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 validationID) view returns((uint8,bytes,uint64,uint64,uint64,uint64,uint64))
func (_IACP99ValidatorManager *IACP99ValidatorManagerCallerSession) GetValidator(validationID [32]byte) (Validator, error) {
	return _IACP99ValidatorManager.Contract.GetValidator(&_IACP99ValidatorManager.CallOpts, validationID)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns(bytes32)
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactor) CompleteEndValidation(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _IACP99ValidatorManager.contract.Transact(opts, "completeEndValidation", messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns(bytes32)
func (_IACP99ValidatorManager *IACP99ValidatorManagerSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.CompleteEndValidation(&_IACP99ValidatorManager.TransactOpts, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns(bytes32)
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactorSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.CompleteEndValidation(&_IACP99ValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _IACP99ValidatorManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_IACP99ValidatorManager *IACP99ValidatorManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.CompleteValidatorRegistration(&_IACP99ValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.CompleteValidatorRegistration(&_IACP99ValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorWeightChange is a paid mutator transaction binding the contract method 0x5fb0236c.
//
// Solidity: function completeValidatorWeightChange(bytes32 validationID) returns()
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactor) CompleteValidatorWeightChange(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _IACP99ValidatorManager.contract.Transact(opts, "completeValidatorWeightChange", validationID)
}

// CompleteValidatorWeightChange is a paid mutator transaction binding the contract method 0x5fb0236c.
//
// Solidity: function completeValidatorWeightChange(bytes32 validationID) returns()
func (_IACP99ValidatorManager *IACP99ValidatorManagerSession) CompleteValidatorWeightChange(validationID [32]byte) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.CompleteValidatorWeightChange(&_IACP99ValidatorManager.TransactOpts, validationID)
}

// CompleteValidatorWeightChange is a paid mutator transaction binding the contract method 0x5fb0236c.
//
// Solidity: function completeValidatorWeightChange(bytes32 validationID) returns()
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactorSession) CompleteValidatorWeightChange(validationID [32]byte) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.CompleteValidatorWeightChange(&_IACP99ValidatorManager.TransactOpts, validationID)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x97fb70d4.
//
// Solidity: function initializeEndValidation(bytes32 validationID) returns()
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactor) InitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _IACP99ValidatorManager.contract.Transact(opts, "initializeEndValidation", validationID)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x97fb70d4.
//
// Solidity: function initializeEndValidation(bytes32 validationID) returns()
func (_IACP99ValidatorManager *IACP99ValidatorManagerSession) InitializeEndValidation(validationID [32]byte) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.InitializeEndValidation(&_IACP99ValidatorManager.TransactOpts, validationID)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x97fb70d4.
//
// Solidity: function initializeEndValidation(bytes32 validationID) returns()
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactorSession) InitializeEndValidation(validationID [32]byte) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.InitializeEndValidation(&_IACP99ValidatorManager.TransactOpts, validationID)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x9ba96b86.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) input, uint64 weight) returns(bytes32)
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactor) InitializeValidatorRegistration(opts *bind.TransactOpts, input ValidatorRegistrationInput, weight uint64) (*types.Transaction, error) {
	return _IACP99ValidatorManager.contract.Transact(opts, "initializeValidatorRegistration", input, weight)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x9ba96b86.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) input, uint64 weight) returns(bytes32)
func (_IACP99ValidatorManager *IACP99ValidatorManagerSession) InitializeValidatorRegistration(input ValidatorRegistrationInput, weight uint64) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.InitializeValidatorRegistration(&_IACP99ValidatorManager.TransactOpts, input, weight)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x9ba96b86.
//
// Solidity: function initializeValidatorRegistration((bytes,bytes,uint64,(uint32,address[]),(uint32,address[])) input, uint64 weight) returns(bytes32)
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactorSession) InitializeValidatorRegistration(input ValidatorRegistrationInput, weight uint64) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.InitializeValidatorRegistration(&_IACP99ValidatorManager.TransactOpts, input, weight)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactor) InitializeValidatorSet(opts *bind.TransactOpts, conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _IACP99ValidatorManager.contract.Transact(opts, "initializeValidatorSet", conversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_IACP99ValidatorManager *IACP99ValidatorManagerSession) InitializeValidatorSet(conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.InitializeValidatorSet(&_IACP99ValidatorManager.TransactOpts, conversionData, messageIndex)
}

// InitializeValidatorSet is a paid mutator transaction binding the contract method 0x20d91b7a.
//
// Solidity: function initializeValidatorSet((bytes32,bytes32,address,(bytes,bytes,uint64)[]) conversionData, uint32 messageIndex) returns()
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactorSession) InitializeValidatorSet(conversionData ConversionData, messageIndex uint32) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.InitializeValidatorSet(&_IACP99ValidatorManager.TransactOpts, conversionData, messageIndex)
}

// InitializeValidatorWeightChange is a paid mutator transaction binding the contract method 0xe3bb1234.
//
// Solidity: function initializeValidatorWeightChange(bytes32 validationID, uint64 weight) returns(uint64)
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactor) InitializeValidatorWeightChange(opts *bind.TransactOpts, validationID [32]byte, weight uint64) (*types.Transaction, error) {
	return _IACP99ValidatorManager.contract.Transact(opts, "initializeValidatorWeightChange", validationID, weight)
}

// InitializeValidatorWeightChange is a paid mutator transaction binding the contract method 0xe3bb1234.
//
// Solidity: function initializeValidatorWeightChange(bytes32 validationID, uint64 weight) returns(uint64)
func (_IACP99ValidatorManager *IACP99ValidatorManagerSession) InitializeValidatorWeightChange(validationID [32]byte, weight uint64) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.InitializeValidatorWeightChange(&_IACP99ValidatorManager.TransactOpts, validationID, weight)
}

// InitializeValidatorWeightChange is a paid mutator transaction binding the contract method 0xe3bb1234.
//
// Solidity: function initializeValidatorWeightChange(bytes32 validationID, uint64 weight) returns(uint64)
func (_IACP99ValidatorManager *IACP99ValidatorManagerTransactorSession) InitializeValidatorWeightChange(validationID [32]byte, weight uint64) (*types.Transaction, error) {
	return _IACP99ValidatorManager.Contract.InitializeValidatorWeightChange(&_IACP99ValidatorManager.TransactOpts, validationID, weight)
}
