// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iacp99securitymodule

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

// IACP99SecurityModuleMetaData contains all meta data concerning the IACP99SecurityModule contract.
var IACP99SecurityModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"handleCompleteEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"handleCompleteValidatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"handleCompleteValidatorWeightChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"handleInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"handleInitializeValidatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"weight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"handleInitializeValidatorWeightChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IACP99SecurityModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use IACP99SecurityModuleMetaData.ABI instead.
var IACP99SecurityModuleABI = IACP99SecurityModuleMetaData.ABI

// IACP99SecurityModule is an auto generated Go binding around an Ethereum contract.
type IACP99SecurityModule struct {
	IACP99SecurityModuleCaller     // Read-only binding to the contract
	IACP99SecurityModuleTransactor // Write-only binding to the contract
	IACP99SecurityModuleFilterer   // Log filterer for contract events
}

// IACP99SecurityModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IACP99SecurityModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IACP99SecurityModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IACP99SecurityModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IACP99SecurityModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IACP99SecurityModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IACP99SecurityModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IACP99SecurityModuleSession struct {
	Contract     *IACP99SecurityModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IACP99SecurityModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IACP99SecurityModuleCallerSession struct {
	Contract *IACP99SecurityModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IACP99SecurityModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IACP99SecurityModuleTransactorSession struct {
	Contract     *IACP99SecurityModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IACP99SecurityModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IACP99SecurityModuleRaw struct {
	Contract *IACP99SecurityModule // Generic contract binding to access the raw methods on
}

// IACP99SecurityModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IACP99SecurityModuleCallerRaw struct {
	Contract *IACP99SecurityModuleCaller // Generic read-only contract binding to access the raw methods on
}

// IACP99SecurityModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IACP99SecurityModuleTransactorRaw struct {
	Contract *IACP99SecurityModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIACP99SecurityModule creates a new instance of IACP99SecurityModule, bound to a specific deployed contract.
func NewIACP99SecurityModule(address common.Address, backend bind.ContractBackend) (*IACP99SecurityModule, error) {
	contract, err := bindIACP99SecurityModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IACP99SecurityModule{IACP99SecurityModuleCaller: IACP99SecurityModuleCaller{contract: contract}, IACP99SecurityModuleTransactor: IACP99SecurityModuleTransactor{contract: contract}, IACP99SecurityModuleFilterer: IACP99SecurityModuleFilterer{contract: contract}}, nil
}

// NewIACP99SecurityModuleCaller creates a new read-only instance of IACP99SecurityModule, bound to a specific deployed contract.
func NewIACP99SecurityModuleCaller(address common.Address, caller bind.ContractCaller) (*IACP99SecurityModuleCaller, error) {
	contract, err := bindIACP99SecurityModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IACP99SecurityModuleCaller{contract: contract}, nil
}

// NewIACP99SecurityModuleTransactor creates a new write-only instance of IACP99SecurityModule, bound to a specific deployed contract.
func NewIACP99SecurityModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*IACP99SecurityModuleTransactor, error) {
	contract, err := bindIACP99SecurityModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IACP99SecurityModuleTransactor{contract: contract}, nil
}

// NewIACP99SecurityModuleFilterer creates a new log filterer instance of IACP99SecurityModule, bound to a specific deployed contract.
func NewIACP99SecurityModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*IACP99SecurityModuleFilterer, error) {
	contract, err := bindIACP99SecurityModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IACP99SecurityModuleFilterer{contract: contract}, nil
}

// bindIACP99SecurityModule binds a generic wrapper to an already deployed contract.
func bindIACP99SecurityModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IACP99SecurityModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IACP99SecurityModule *IACP99SecurityModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IACP99SecurityModule.Contract.IACP99SecurityModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IACP99SecurityModule *IACP99SecurityModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.IACP99SecurityModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IACP99SecurityModule *IACP99SecurityModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.IACP99SecurityModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IACP99SecurityModule *IACP99SecurityModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IACP99SecurityModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IACP99SecurityModule *IACP99SecurityModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IACP99SecurityModule *IACP99SecurityModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.contract.Transact(opts, method, params...)
}

// HandleCompleteEndValidation is a paid mutator transaction binding the contract method 0xf7795a8c.
//
// Solidity: function handleCompleteEndValidation(bytes32 validationID) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleTransactor) HandleCompleteEndValidation(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.contract.Transact(opts, "handleCompleteEndValidation", validationID)
}

// HandleCompleteEndValidation is a paid mutator transaction binding the contract method 0xf7795a8c.
//
// Solidity: function handleCompleteEndValidation(bytes32 validationID) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleSession) HandleCompleteEndValidation(validationID [32]byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.HandleCompleteEndValidation(&_IACP99SecurityModule.TransactOpts, validationID)
}

// HandleCompleteEndValidation is a paid mutator transaction binding the contract method 0xf7795a8c.
//
// Solidity: function handleCompleteEndValidation(bytes32 validationID) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleTransactorSession) HandleCompleteEndValidation(validationID [32]byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.HandleCompleteEndValidation(&_IACP99SecurityModule.TransactOpts, validationID)
}

// HandleCompleteValidatorRegistration is a paid mutator transaction binding the contract method 0x56ab447d.
//
// Solidity: function handleCompleteValidatorRegistration(bytes32 validationID) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleTransactor) HandleCompleteValidatorRegistration(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.contract.Transact(opts, "handleCompleteValidatorRegistration", validationID)
}

// HandleCompleteValidatorRegistration is a paid mutator transaction binding the contract method 0x56ab447d.
//
// Solidity: function handleCompleteValidatorRegistration(bytes32 validationID) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleSession) HandleCompleteValidatorRegistration(validationID [32]byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.HandleCompleteValidatorRegistration(&_IACP99SecurityModule.TransactOpts, validationID)
}

// HandleCompleteValidatorRegistration is a paid mutator transaction binding the contract method 0x56ab447d.
//
// Solidity: function handleCompleteValidatorRegistration(bytes32 validationID) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleTransactorSession) HandleCompleteValidatorRegistration(validationID [32]byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.HandleCompleteValidatorRegistration(&_IACP99SecurityModule.TransactOpts, validationID)
}

// HandleCompleteValidatorWeightChange is a paid mutator transaction binding the contract method 0xeffb55b4.
//
// Solidity: function handleCompleteValidatorWeightChange(bytes32 validationID, bytes args) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleTransactor) HandleCompleteValidatorWeightChange(opts *bind.TransactOpts, validationID [32]byte, args []byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.contract.Transact(opts, "handleCompleteValidatorWeightChange", validationID, args)
}

// HandleCompleteValidatorWeightChange is a paid mutator transaction binding the contract method 0xeffb55b4.
//
// Solidity: function handleCompleteValidatorWeightChange(bytes32 validationID, bytes args) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleSession) HandleCompleteValidatorWeightChange(validationID [32]byte, args []byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.HandleCompleteValidatorWeightChange(&_IACP99SecurityModule.TransactOpts, validationID, args)
}

// HandleCompleteValidatorWeightChange is a paid mutator transaction binding the contract method 0xeffb55b4.
//
// Solidity: function handleCompleteValidatorWeightChange(bytes32 validationID, bytes args) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleTransactorSession) HandleCompleteValidatorWeightChange(validationID [32]byte, args []byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.HandleCompleteValidatorWeightChange(&_IACP99SecurityModule.TransactOpts, validationID, args)
}

// HandleInitializeEndValidation is a paid mutator transaction binding the contract method 0xbddb5a14.
//
// Solidity: function handleInitializeEndValidation(bytes32 validationID, bytes args) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleTransactor) HandleInitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, args []byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.contract.Transact(opts, "handleInitializeEndValidation", validationID, args)
}

// HandleInitializeEndValidation is a paid mutator transaction binding the contract method 0xbddb5a14.
//
// Solidity: function handleInitializeEndValidation(bytes32 validationID, bytes args) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleSession) HandleInitializeEndValidation(validationID [32]byte, args []byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.HandleInitializeEndValidation(&_IACP99SecurityModule.TransactOpts, validationID, args)
}

// HandleInitializeEndValidation is a paid mutator transaction binding the contract method 0xbddb5a14.
//
// Solidity: function handleInitializeEndValidation(bytes32 validationID, bytes args) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleTransactorSession) HandleInitializeEndValidation(validationID [32]byte, args []byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.HandleInitializeEndValidation(&_IACP99SecurityModule.TransactOpts, validationID, args)
}

// HandleInitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x004050c8.
//
// Solidity: function handleInitializeValidatorRegistration(bytes32 validationID, uint64 weight, bytes args) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleTransactor) HandleInitializeValidatorRegistration(opts *bind.TransactOpts, validationID [32]byte, weight uint64, args []byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.contract.Transact(opts, "handleInitializeValidatorRegistration", validationID, weight, args)
}

// HandleInitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x004050c8.
//
// Solidity: function handleInitializeValidatorRegistration(bytes32 validationID, uint64 weight, bytes args) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleSession) HandleInitializeValidatorRegistration(validationID [32]byte, weight uint64, args []byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.HandleInitializeValidatorRegistration(&_IACP99SecurityModule.TransactOpts, validationID, weight, args)
}

// HandleInitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x004050c8.
//
// Solidity: function handleInitializeValidatorRegistration(bytes32 validationID, uint64 weight, bytes args) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleTransactorSession) HandleInitializeValidatorRegistration(validationID [32]byte, weight uint64, args []byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.HandleInitializeValidatorRegistration(&_IACP99SecurityModule.TransactOpts, validationID, weight, args)
}

// HandleInitializeValidatorWeightChange is a paid mutator transaction binding the contract method 0x98d9fa7c.
//
// Solidity: function handleInitializeValidatorWeightChange(bytes32 validationID, uint64 weight, uint64 nonce, bytes args) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleTransactor) HandleInitializeValidatorWeightChange(opts *bind.TransactOpts, validationID [32]byte, weight uint64, nonce uint64, args []byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.contract.Transact(opts, "handleInitializeValidatorWeightChange", validationID, weight, nonce, args)
}

// HandleInitializeValidatorWeightChange is a paid mutator transaction binding the contract method 0x98d9fa7c.
//
// Solidity: function handleInitializeValidatorWeightChange(bytes32 validationID, uint64 weight, uint64 nonce, bytes args) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleSession) HandleInitializeValidatorWeightChange(validationID [32]byte, weight uint64, nonce uint64, args []byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.HandleInitializeValidatorWeightChange(&_IACP99SecurityModule.TransactOpts, validationID, weight, nonce, args)
}

// HandleInitializeValidatorWeightChange is a paid mutator transaction binding the contract method 0x98d9fa7c.
//
// Solidity: function handleInitializeValidatorWeightChange(bytes32 validationID, uint64 weight, uint64 nonce, bytes args) returns()
func (_IACP99SecurityModule *IACP99SecurityModuleTransactorSession) HandleInitializeValidatorWeightChange(validationID [32]byte, weight uint64, nonce uint64, args []byte) (*types.Transaction, error) {
	return _IACP99SecurityModule.Contract.HandleInitializeValidatorWeightChange(&_IACP99SecurityModule.TransactOpts, validationID, weight, nonce, args)
}
