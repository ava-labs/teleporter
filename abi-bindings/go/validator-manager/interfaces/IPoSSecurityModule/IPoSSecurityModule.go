// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipossecuritymodule

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

// IPoSSecurityModuleMetaData contains all meta data concerning the IPoSSecurityModule contract.
var IPoSSecurityModuleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"}],\"name\":\"DelegationEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validatorWeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"delegatorWeight\",\"type\":\"uint64\"}],\"name\":\"DelegatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"DelegatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"DelegatorRemovalInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"uptime\",\"type\":\"uint64\"}],\"name\":\"UptimeUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"changeDelegatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"changeValidatorRewardRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"}],\"name\":\"claimDelegationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeDelegatorRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"completeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"forceInitializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"forceInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"forceInitializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"initializeEndDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"includeUptimeProof\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"initializeEndValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"delegationID\",\"type\":\"bytes32\"}],\"name\":\"resendUpdateDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validationID\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"submitUptimeProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPoSSecurityModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use IPoSSecurityModuleMetaData.ABI instead.
var IPoSSecurityModuleABI = IPoSSecurityModuleMetaData.ABI

// IPoSSecurityModule is an auto generated Go binding around an Ethereum contract.
type IPoSSecurityModule struct {
	IPoSSecurityModuleCaller     // Read-only binding to the contract
	IPoSSecurityModuleTransactor // Write-only binding to the contract
	IPoSSecurityModuleFilterer   // Log filterer for contract events
}

// IPoSSecurityModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoSSecurityModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoSSecurityModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoSSecurityModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoSSecurityModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoSSecurityModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoSSecurityModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoSSecurityModuleSession struct {
	Contract     *IPoSSecurityModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IPoSSecurityModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoSSecurityModuleCallerSession struct {
	Contract *IPoSSecurityModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IPoSSecurityModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoSSecurityModuleTransactorSession struct {
	Contract     *IPoSSecurityModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IPoSSecurityModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoSSecurityModuleRaw struct {
	Contract *IPoSSecurityModule // Generic contract binding to access the raw methods on
}

// IPoSSecurityModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoSSecurityModuleCallerRaw struct {
	Contract *IPoSSecurityModuleCaller // Generic read-only contract binding to access the raw methods on
}

// IPoSSecurityModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoSSecurityModuleTransactorRaw struct {
	Contract *IPoSSecurityModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPoSSecurityModule creates a new instance of IPoSSecurityModule, bound to a specific deployed contract.
func NewIPoSSecurityModule(address common.Address, backend bind.ContractBackend) (*IPoSSecurityModule, error) {
	contract, err := bindIPoSSecurityModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPoSSecurityModule{IPoSSecurityModuleCaller: IPoSSecurityModuleCaller{contract: contract}, IPoSSecurityModuleTransactor: IPoSSecurityModuleTransactor{contract: contract}, IPoSSecurityModuleFilterer: IPoSSecurityModuleFilterer{contract: contract}}, nil
}

// NewIPoSSecurityModuleCaller creates a new read-only instance of IPoSSecurityModule, bound to a specific deployed contract.
func NewIPoSSecurityModuleCaller(address common.Address, caller bind.ContractCaller) (*IPoSSecurityModuleCaller, error) {
	contract, err := bindIPoSSecurityModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoSSecurityModuleCaller{contract: contract}, nil
}

// NewIPoSSecurityModuleTransactor creates a new write-only instance of IPoSSecurityModule, bound to a specific deployed contract.
func NewIPoSSecurityModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoSSecurityModuleTransactor, error) {
	contract, err := bindIPoSSecurityModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoSSecurityModuleTransactor{contract: contract}, nil
}

// NewIPoSSecurityModuleFilterer creates a new log filterer instance of IPoSSecurityModule, bound to a specific deployed contract.
func NewIPoSSecurityModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoSSecurityModuleFilterer, error) {
	contract, err := bindIPoSSecurityModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoSSecurityModuleFilterer{contract: contract}, nil
}

// bindIPoSSecurityModule binds a generic wrapper to an already deployed contract.
func bindIPoSSecurityModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPoSSecurityModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoSSecurityModule *IPoSSecurityModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoSSecurityModule.Contract.IPoSSecurityModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoSSecurityModule *IPoSSecurityModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.IPoSSecurityModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoSSecurityModule *IPoSSecurityModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.IPoSSecurityModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoSSecurityModule *IPoSSecurityModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoSSecurityModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.contract.Transact(opts, method, params...)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address recipient) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) ChangeDelegatorRewardRecipient(opts *bind.TransactOpts, delegationID [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "changeDelegatorRewardRecipient", delegationID, recipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address recipient) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ChangeDelegatorRewardRecipient(&_IPoSSecurityModule.TransactOpts, delegationID, recipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address recipient) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ChangeDelegatorRewardRecipient(&_IPoSSecurityModule.TransactOpts, delegationID, recipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address recipient) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) ChangeValidatorRewardRecipient(opts *bind.TransactOpts, validationID [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "changeValidatorRewardRecipient", validationID, recipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address recipient) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) ChangeValidatorRewardRecipient(validationID [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ChangeValidatorRewardRecipient(&_IPoSSecurityModule.TransactOpts, validationID, recipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address recipient) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) ChangeValidatorRewardRecipient(validationID [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ChangeValidatorRewardRecipient(&_IPoSSecurityModule.TransactOpts, validationID, recipient)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ClaimDelegationFees(&_IPoSSecurityModule.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ClaimDelegationFees(&_IPoSSecurityModule.TransactOpts, validationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.CompleteDelegatorRegistration(&_IPoSSecurityModule.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.CompleteDelegatorRegistration(&_IPoSSecurityModule.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) CompleteEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "completeEndDelegation", delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.CompleteEndDelegation(&_IPoSSecurityModule.TransactOpts, delegationID, messageIndex)
}

// CompleteEndDelegation is a paid mutator transaction binding the contract method 0x80dd672f.
//
// Solidity: function completeEndDelegation(bytes32 delegationID, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) CompleteEndDelegation(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.CompleteEndDelegation(&_IPoSSecurityModule.TransactOpts, delegationID, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) ForceInitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "forceInitializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ForceInitializeEndDelegation(&_IPoSSecurityModule.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation is a paid mutator transaction binding the contract method 0x1ec44724.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) ForceInitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ForceInitializeEndDelegation(&_IPoSSecurityModule.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x37b9be8f.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address recipientAddress) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) ForceInitializeEndDelegation0(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, recipientAddress common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "forceInitializeEndDelegation0", delegationID, includeUptimeProof, messageIndex, recipientAddress)
}

// ForceInitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x37b9be8f.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address recipientAddress) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) ForceInitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, recipientAddress common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ForceInitializeEndDelegation0(&_IPoSSecurityModule.TransactOpts, delegationID, includeUptimeProof, messageIndex, recipientAddress)
}

// ForceInitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x37b9be8f.
//
// Solidity: function forceInitializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address recipientAddress) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) ForceInitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, recipientAddress common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ForceInitializeEndDelegation0(&_IPoSSecurityModule.TransactOpts, delegationID, includeUptimeProof, messageIndex, recipientAddress)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) ForceInitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "forceInitializeEndValidation", validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ForceInitializeEndValidation(&_IPoSSecurityModule.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation is a paid mutator transaction binding the contract method 0x3a1cfff6.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) ForceInitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ForceInitializeEndValidation(&_IPoSSecurityModule.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitializeEndValidation0 is a paid mutator transaction binding the contract method 0x7d8d2f77.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address recipientAddress) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) ForceInitializeEndValidation0(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32, recipientAddress common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "forceInitializeEndValidation0", validationID, includeUptimeProof, messageIndex, recipientAddress)
}

// ForceInitializeEndValidation0 is a paid mutator transaction binding the contract method 0x7d8d2f77.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address recipientAddress) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) ForceInitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, recipientAddress common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ForceInitializeEndValidation0(&_IPoSSecurityModule.TransactOpts, validationID, includeUptimeProof, messageIndex, recipientAddress)
}

// ForceInitializeEndValidation0 is a paid mutator transaction binding the contract method 0x7d8d2f77.
//
// Solidity: function forceInitializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address recipientAddress) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) ForceInitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, recipientAddress common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ForceInitializeEndValidation0(&_IPoSSecurityModule.TransactOpts, validationID, includeUptimeProof, messageIndex, recipientAddress)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) InitializeEndDelegation(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "initializeEndDelegation", delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.InitializeEndDelegation(&_IPoSSecurityModule.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation is a paid mutator transaction binding the contract method 0x0118acc4.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) InitializeEndDelegation(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.InitializeEndDelegation(&_IPoSSecurityModule.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x9ae06447.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address recipientAddress) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) InitializeEndDelegation0(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, recipientAddress common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "initializeEndDelegation0", delegationID, includeUptimeProof, messageIndex, recipientAddress)
}

// InitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x9ae06447.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address recipientAddress) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) InitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, recipientAddress common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.InitializeEndDelegation0(&_IPoSSecurityModule.TransactOpts, delegationID, includeUptimeProof, messageIndex, recipientAddress)
}

// InitializeEndDelegation0 is a paid mutator transaction binding the contract method 0x9ae06447.
//
// Solidity: function initializeEndDelegation(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex, address recipientAddress) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) InitializeEndDelegation0(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32, recipientAddress common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.InitializeEndDelegation0(&_IPoSSecurityModule.TransactOpts, delegationID, includeUptimeProof, messageIndex, recipientAddress)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x5dd6a6cb.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address recipientAddress) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) InitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32, recipientAddress common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "initializeEndValidation", validationID, includeUptimeProof, messageIndex, recipientAddress)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x5dd6a6cb.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address recipientAddress) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, recipientAddress common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.InitializeEndValidation(&_IPoSSecurityModule.TransactOpts, validationID, includeUptimeProof, messageIndex, recipientAddress)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x5dd6a6cb.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex, address recipientAddress) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) InitializeEndValidation(validationID [32]byte, includeUptimeProof bool, messageIndex uint32, recipientAddress common.Address) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.InitializeEndValidation(&_IPoSSecurityModule.TransactOpts, validationID, includeUptimeProof, messageIndex, recipientAddress)
}

// InitializeEndValidation0 is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) InitializeEndValidation0(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "initializeEndValidation0", validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation0 is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) InitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.InitializeEndValidation0(&_IPoSSecurityModule.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitializeEndValidation0 is a paid mutator transaction binding the contract method 0x76f78621.
//
// Solidity: function initializeEndValidation(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) InitializeEndValidation0(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.InitializeEndValidation0(&_IPoSSecurityModule.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) ResendUpdateDelegation(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "resendUpdateDelegation", delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ResendUpdateDelegation(&_IPoSSecurityModule.TransactOpts, delegationID)
}

// ResendUpdateDelegation is a paid mutator transaction binding the contract method 0xba3a4b97.
//
// Solidity: function resendUpdateDelegation(bytes32 delegationID) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) ResendUpdateDelegation(delegationID [32]byte) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.ResendUpdateDelegation(&_IPoSSecurityModule.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.SubmitUptimeProof(&_IPoSSecurityModule.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_IPoSSecurityModule *IPoSSecurityModuleTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _IPoSSecurityModule.Contract.SubmitUptimeProof(&_IPoSSecurityModule.TransactOpts, validationID, messageIndex)
}

// IPoSSecurityModuleDelegationEndedIterator is returned from FilterDelegationEnded and is used to iterate over the raw logs and unpacked data for DelegationEnded events raised by the IPoSSecurityModule contract.
type IPoSSecurityModuleDelegationEndedIterator struct {
	Event *IPoSSecurityModuleDelegationEnded // Event containing the contract specifics and raw log

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
func (it *IPoSSecurityModuleDelegationEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSSecurityModuleDelegationEnded)
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
		it.Event = new(IPoSSecurityModuleDelegationEnded)
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
func (it *IPoSSecurityModuleDelegationEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSSecurityModuleDelegationEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSSecurityModuleDelegationEnded represents a DelegationEnded event raised by the IPoSSecurityModule contract.
type IPoSSecurityModuleDelegationEnded struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegationEnded is a free log retrieval operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) FilterDelegationEnded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*IPoSSecurityModuleDelegationEndedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSSecurityModule.contract.FilterLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &IPoSSecurityModuleDelegationEndedIterator{contract: _IPoSSecurityModule.contract, event: "DelegationEnded", logs: logs, sub: sub}, nil
}

// WatchDelegationEnded is a free log subscription operation binding the contract event 0x8ececf510070c320d9a55323ffabe350e294ae505fc0c509dc5736da6f5cc993.
//
// Solidity: event DelegationEnded(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) WatchDelegationEnded(opts *bind.WatchOpts, sink chan<- *IPoSSecurityModuleDelegationEnded, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSSecurityModule.contract.WatchLogs(opts, "DelegationEnded", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSSecurityModuleDelegationEnded)
				if err := _IPoSSecurityModule.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
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
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) ParseDelegationEnded(log types.Log) (*IPoSSecurityModuleDelegationEnded, error) {
	event := new(IPoSSecurityModuleDelegationEnded)
	if err := _IPoSSecurityModule.contract.UnpackLog(event, "DelegationEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSSecurityModuleDelegatorAddedIterator is returned from FilterDelegatorAdded and is used to iterate over the raw logs and unpacked data for DelegatorAdded events raised by the IPoSSecurityModule contract.
type IPoSSecurityModuleDelegatorAddedIterator struct {
	Event *IPoSSecurityModuleDelegatorAdded // Event containing the contract specifics and raw log

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
func (it *IPoSSecurityModuleDelegatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSSecurityModuleDelegatorAdded)
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
		it.Event = new(IPoSSecurityModuleDelegatorAdded)
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
func (it *IPoSSecurityModuleDelegatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSSecurityModuleDelegatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSSecurityModuleDelegatorAdded represents a DelegatorAdded event raised by the IPoSSecurityModule contract.
type IPoSSecurityModuleDelegatorAdded struct {
	DelegationID     [32]byte
	ValidationID     [32]byte
	DelegatorAddress common.Address
	Nonce            uint64
	ValidatorWeight  uint64
	DelegatorWeight  uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDelegatorAdded is a free log retrieval operation binding the contract event 0xeea0c6dd7e326ed588ce418df1563a24af963b7a54714ba659fab8152435795b.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight)
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) FilterDelegatorAdded(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*IPoSSecurityModuleDelegatorAddedIterator, error) {

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

	logs, sub, err := _IPoSSecurityModule.contract.FilterLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPoSSecurityModuleDelegatorAddedIterator{contract: _IPoSSecurityModule.contract, event: "DelegatorAdded", logs: logs, sub: sub}, nil
}

// WatchDelegatorAdded is a free log subscription operation binding the contract event 0xeea0c6dd7e326ed588ce418df1563a24af963b7a54714ba659fab8152435795b.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight)
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) WatchDelegatorAdded(opts *bind.WatchOpts, sink chan<- *IPoSSecurityModuleDelegatorAdded, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IPoSSecurityModule.contract.WatchLogs(opts, "DelegatorAdded", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSSecurityModuleDelegatorAdded)
				if err := _IPoSSecurityModule.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
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

// ParseDelegatorAdded is a log parse operation binding the contract event 0xeea0c6dd7e326ed588ce418df1563a24af963b7a54714ba659fab8152435795b.
//
// Solidity: event DelegatorAdded(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight)
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) ParseDelegatorAdded(log types.Log) (*IPoSSecurityModuleDelegatorAdded, error) {
	event := new(IPoSSecurityModuleDelegatorAdded)
	if err := _IPoSSecurityModule.contract.UnpackLog(event, "DelegatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSSecurityModuleDelegatorRegisteredIterator is returned from FilterDelegatorRegistered and is used to iterate over the raw logs and unpacked data for DelegatorRegistered events raised by the IPoSSecurityModule contract.
type IPoSSecurityModuleDelegatorRegisteredIterator struct {
	Event *IPoSSecurityModuleDelegatorRegistered // Event containing the contract specifics and raw log

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
func (it *IPoSSecurityModuleDelegatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSSecurityModuleDelegatorRegistered)
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
		it.Event = new(IPoSSecurityModuleDelegatorRegistered)
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
func (it *IPoSSecurityModuleDelegatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSSecurityModuleDelegatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSSecurityModuleDelegatorRegistered represents a DelegatorRegistered event raised by the IPoSSecurityModule contract.
type IPoSSecurityModuleDelegatorRegistered struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRegistered is a free log retrieval operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) FilterDelegatorRegistered(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*IPoSSecurityModuleDelegatorRegisteredIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSSecurityModule.contract.FilterLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &IPoSSecurityModuleDelegatorRegisteredIterator{contract: _IPoSSecurityModule.contract, event: "DelegatorRegistered", logs: logs, sub: sub}, nil
}

// WatchDelegatorRegistered is a free log subscription operation binding the contract event 0x047059b465069b8b751836b41f9f1d83daff583d2238cc7fbb461437ec23a4f6.
//
// Solidity: event DelegatorRegistered(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) WatchDelegatorRegistered(opts *bind.WatchOpts, sink chan<- *IPoSSecurityModuleDelegatorRegistered, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSSecurityModule.contract.WatchLogs(opts, "DelegatorRegistered", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSSecurityModuleDelegatorRegistered)
				if err := _IPoSSecurityModule.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
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
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) ParseDelegatorRegistered(log types.Log) (*IPoSSecurityModuleDelegatorRegistered, error) {
	event := new(IPoSSecurityModuleDelegatorRegistered)
	if err := _IPoSSecurityModule.contract.UnpackLog(event, "DelegatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSSecurityModuleDelegatorRemovalInitializedIterator is returned from FilterDelegatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for DelegatorRemovalInitialized events raised by the IPoSSecurityModule contract.
type IPoSSecurityModuleDelegatorRemovalInitializedIterator struct {
	Event *IPoSSecurityModuleDelegatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *IPoSSecurityModuleDelegatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSSecurityModuleDelegatorRemovalInitialized)
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
		it.Event = new(IPoSSecurityModuleDelegatorRemovalInitialized)
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
func (it *IPoSSecurityModuleDelegatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSSecurityModuleDelegatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSSecurityModuleDelegatorRemovalInitialized represents a DelegatorRemovalInitialized event raised by the IPoSSecurityModule contract.
type IPoSSecurityModuleDelegatorRemovalInitialized struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRemovalInitialized is a free log retrieval operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) FilterDelegatorRemovalInitialized(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*IPoSSecurityModuleDelegatorRemovalInitializedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSSecurityModule.contract.FilterLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &IPoSSecurityModuleDelegatorRemovalInitializedIterator{contract: _IPoSSecurityModule.contract, event: "DelegatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchDelegatorRemovalInitialized is a free log subscription operation binding the contract event 0x366d336c0ab380dc799f095a6f82a26326585c52909cc698b09ba4540709ed57.
//
// Solidity: event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) WatchDelegatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *IPoSSecurityModuleDelegatorRemovalInitialized, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSSecurityModule.contract.WatchLogs(opts, "DelegatorRemovalInitialized", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSSecurityModuleDelegatorRemovalInitialized)
				if err := _IPoSSecurityModule.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
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
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) ParseDelegatorRemovalInitialized(log types.Log) (*IPoSSecurityModuleDelegatorRemovalInitialized, error) {
	event := new(IPoSSecurityModuleDelegatorRemovalInitialized)
	if err := _IPoSSecurityModule.contract.UnpackLog(event, "DelegatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoSSecurityModuleUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the IPoSSecurityModule contract.
type IPoSSecurityModuleUptimeUpdatedIterator struct {
	Event *IPoSSecurityModuleUptimeUpdated // Event containing the contract specifics and raw log

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
func (it *IPoSSecurityModuleUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoSSecurityModuleUptimeUpdated)
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
		it.Event = new(IPoSSecurityModuleUptimeUpdated)
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
func (it *IPoSSecurityModuleUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoSSecurityModuleUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoSSecurityModuleUptimeUpdated represents a UptimeUpdated event raised by the IPoSSecurityModule contract.
type IPoSSecurityModuleUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*IPoSSecurityModuleUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSSecurityModule.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &IPoSSecurityModuleUptimeUpdatedIterator{contract: _IPoSSecurityModule.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *IPoSSecurityModuleUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _IPoSSecurityModule.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoSSecurityModuleUptimeUpdated)
				if err := _IPoSSecurityModule.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
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
func (_IPoSSecurityModule *IPoSSecurityModuleFilterer) ParseUptimeUpdated(log types.Log) (*IPoSSecurityModuleUptimeUpdated, error) {
	event := new(IPoSSecurityModuleUptimeUpdated)
	if err := _IPoSSecurityModule.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
