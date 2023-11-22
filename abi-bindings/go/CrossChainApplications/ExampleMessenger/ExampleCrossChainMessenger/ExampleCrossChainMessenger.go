// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package examplecrosschainmessenger

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

// ExampleCrossChainMessengerMetaData contains all meta data concerning the ExampleCrossChainMessenger contract.
var ExampleCrossChainMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ReceiveMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"SendMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"}],\"name\":\"getCurrentMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ExampleCrossChainMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleCrossChainMessengerMetaData.ABI instead.
var ExampleCrossChainMessengerABI = ExampleCrossChainMessengerMetaData.ABI

// ExampleCrossChainMessenger is an auto generated Go binding around an Ethereum contract.
type ExampleCrossChainMessenger struct {
	ExampleCrossChainMessengerCaller     // Read-only binding to the contract
	ExampleCrossChainMessengerTransactor // Write-only binding to the contract
	ExampleCrossChainMessengerFilterer   // Log filterer for contract events
}

// ExampleCrossChainMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleCrossChainMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleCrossChainMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleCrossChainMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleCrossChainMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleCrossChainMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleCrossChainMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleCrossChainMessengerSession struct {
	Contract     *ExampleCrossChainMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExampleCrossChainMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleCrossChainMessengerCallerSession struct {
	Contract *ExampleCrossChainMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ExampleCrossChainMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleCrossChainMessengerTransactorSession struct {
	Contract     *ExampleCrossChainMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ExampleCrossChainMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleCrossChainMessengerRaw struct {
	Contract *ExampleCrossChainMessenger // Generic contract binding to access the raw methods on
}

// ExampleCrossChainMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleCrossChainMessengerCallerRaw struct {
	Contract *ExampleCrossChainMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleCrossChainMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleCrossChainMessengerTransactorRaw struct {
	Contract *ExampleCrossChainMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleCrossChainMessenger creates a new instance of ExampleCrossChainMessenger, bound to a specific deployed contract.
func NewExampleCrossChainMessenger(address common.Address, backend bind.ContractBackend) (*ExampleCrossChainMessenger, error) {
	contract, err := bindExampleCrossChainMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessenger{ExampleCrossChainMessengerCaller: ExampleCrossChainMessengerCaller{contract: contract}, ExampleCrossChainMessengerTransactor: ExampleCrossChainMessengerTransactor{contract: contract}, ExampleCrossChainMessengerFilterer: ExampleCrossChainMessengerFilterer{contract: contract}}, nil
}

// NewExampleCrossChainMessengerCaller creates a new read-only instance of ExampleCrossChainMessenger, bound to a specific deployed contract.
func NewExampleCrossChainMessengerCaller(address common.Address, caller bind.ContractCaller) (*ExampleCrossChainMessengerCaller, error) {
	contract, err := bindExampleCrossChainMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerCaller{contract: contract}, nil
}

// NewExampleCrossChainMessengerTransactor creates a new write-only instance of ExampleCrossChainMessenger, bound to a specific deployed contract.
func NewExampleCrossChainMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleCrossChainMessengerTransactor, error) {
	contract, err := bindExampleCrossChainMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerTransactor{contract: contract}, nil
}

// NewExampleCrossChainMessengerFilterer creates a new log filterer instance of ExampleCrossChainMessenger, bound to a specific deployed contract.
func NewExampleCrossChainMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleCrossChainMessengerFilterer, error) {
	contract, err := bindExampleCrossChainMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerFilterer{contract: contract}, nil
}

// bindExampleCrossChainMessenger binds a generic wrapper to an already deployed contract.
func bindExampleCrossChainMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExampleCrossChainMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleCrossChainMessenger.Contract.ExampleCrossChainMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.ExampleCrossChainMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.ExampleCrossChainMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleCrossChainMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 originChainID) view returns(address, string)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCaller) GetCurrentMessage(opts *bind.CallOpts, originChainID [32]byte) (common.Address, string, error) {
	var out []interface{}
	err := _ExampleCrossChainMessenger.contract.Call(opts, &out, "getCurrentMessage", originChainID)

	if err != nil {
		return *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 originChainID) view returns(address, string)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) GetCurrentMessage(originChainID [32]byte) (common.Address, string, error) {
	return _ExampleCrossChainMessenger.Contract.GetCurrentMessage(&_ExampleCrossChainMessenger.CallOpts, originChainID)
}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 originChainID) view returns(address, string)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCallerSession) GetCurrentMessage(originChainID [32]byte) (common.Address, string, error) {
	return _ExampleCrossChainMessenger.Contract.GetCurrentMessage(&_ExampleCrossChainMessenger.CallOpts, originChainID)
}

// MinTeleporterVersion is a free data retrieval call binding the contract method 0xe49cc553.
//
// Solidity: function minTeleporterVersion() view returns(uint256)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCaller) MinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExampleCrossChainMessenger.contract.Call(opts, &out, "minTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTeleporterVersion is a free data retrieval call binding the contract method 0xe49cc553.
//
// Solidity: function minTeleporterVersion() view returns(uint256)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) MinTeleporterVersion() (*big.Int, error) {
	return _ExampleCrossChainMessenger.Contract.MinTeleporterVersion(&_ExampleCrossChainMessenger.CallOpts)
}

// MinTeleporterVersion is a free data retrieval call binding the contract method 0xe49cc553.
//
// Solidity: function minTeleporterVersion() view returns(uint256)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCallerSession) MinTeleporterVersion() (*big.Int, error) {
	return _ExampleCrossChainMessenger.Contract.MinTeleporterVersion(&_ExampleCrossChainMessenger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExampleCrossChainMessenger.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) Owner() (common.Address, error) {
	return _ExampleCrossChainMessenger.Contract.Owner(&_ExampleCrossChainMessenger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCallerSession) Owner() (common.Address, error) {
	return _ExampleCrossChainMessenger.Contract.Owner(&_ExampleCrossChainMessenger.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExampleCrossChainMessenger.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) TeleporterRegistry() (common.Address, error) {
	return _ExampleCrossChainMessenger.Contract.TeleporterRegistry(&_ExampleCrossChainMessenger.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCallerSession) TeleporterRegistry() (common.Address, error) {
	return _ExampleCrossChainMessenger.Contract.TeleporterRegistry(&_ExampleCrossChainMessenger.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.contract.Transact(opts, "receiveTeleporterMessage", originChainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) ReceiveTeleporterMessage(originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.ReceiveTeleporterMessage(&_ExampleCrossChainMessenger.TransactOpts, originChainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorSession) ReceiveTeleporterMessage(originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.ReceiveTeleporterMessage(&_ExampleCrossChainMessenger.TransactOpts, originChainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.RenounceOwnership(&_ExampleCrossChainMessenger.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.RenounceOwnership(&_ExampleCrossChainMessenger.TransactOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationChainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(uint256)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactor) SendMessage(opts *bind.TransactOpts, destinationChainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.contract.Transact(opts, "sendMessage", destinationChainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationChainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(uint256)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) SendMessage(destinationChainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.SendMessage(&_ExampleCrossChainMessenger.TransactOpts, destinationChainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationChainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(uint256)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorSession) SendMessage(destinationChainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.SendMessage(&_ExampleCrossChainMessenger.TransactOpts, destinationChainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.TransferOwnership(&_ExampleCrossChainMessenger.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.TransferOwnership(&_ExampleCrossChainMessenger.TransactOpts, newOwner)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0xb6109d9d.
//
// Solidity: function updateMinTeleporterVersion() returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.contract.Transact(opts, "updateMinTeleporterVersion")
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0xb6109d9d.
//
// Solidity: function updateMinTeleporterVersion() returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) UpdateMinTeleporterVersion() (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.UpdateMinTeleporterVersion(&_ExampleCrossChainMessenger.TransactOpts)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0xb6109d9d.
//
// Solidity: function updateMinTeleporterVersion() returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorSession) UpdateMinTeleporterVersion() (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.UpdateMinTeleporterVersion(&_ExampleCrossChainMessenger.TransactOpts)
}

// ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator struct {
	Event *ExampleCrossChainMessengerMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleCrossChainMessengerMinTeleporterVersionUpdated)
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
		it.Event = new(ExampleCrossChainMessengerMinTeleporterVersionUpdated)
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
func (it *ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleCrossChainMessengerMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator{contract: _ExampleCrossChainMessenger.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *ExampleCrossChainMessengerMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleCrossChainMessengerMinTeleporterVersionUpdated)
				if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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

// ParseMinTeleporterVersionUpdated is a log parse operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*ExampleCrossChainMessengerMinTeleporterVersionUpdated, error) {
	event := new(ExampleCrossChainMessengerMinTeleporterVersionUpdated)
	if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleCrossChainMessengerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerOwnershipTransferredIterator struct {
	Event *ExampleCrossChainMessengerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExampleCrossChainMessengerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleCrossChainMessengerOwnershipTransferred)
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
		it.Event = new(ExampleCrossChainMessengerOwnershipTransferred)
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
func (it *ExampleCrossChainMessengerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleCrossChainMessengerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleCrossChainMessengerOwnershipTransferred represents a OwnershipTransferred event raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExampleCrossChainMessengerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerOwnershipTransferredIterator{contract: _ExampleCrossChainMessenger.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExampleCrossChainMessengerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleCrossChainMessengerOwnershipTransferred)
				if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) ParseOwnershipTransferred(log types.Log) (*ExampleCrossChainMessengerOwnershipTransferred, error) {
	event := new(ExampleCrossChainMessengerOwnershipTransferred)
	if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleCrossChainMessengerReceiveMessageIterator is returned from FilterReceiveMessage and is used to iterate over the raw logs and unpacked data for ReceiveMessage events raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerReceiveMessageIterator struct {
	Event *ExampleCrossChainMessengerReceiveMessage // Event containing the contract specifics and raw log

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
func (it *ExampleCrossChainMessengerReceiveMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleCrossChainMessengerReceiveMessage)
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
		it.Event = new(ExampleCrossChainMessengerReceiveMessage)
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
func (it *ExampleCrossChainMessengerReceiveMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleCrossChainMessengerReceiveMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleCrossChainMessengerReceiveMessage represents a ReceiveMessage event raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerReceiveMessage struct {
	OriginChainID       [32]byte
	OriginSenderAddress common.Address
	Message             string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReceiveMessage is a free log retrieval operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed originChainID, address indexed originSenderAddress, string message)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) FilterReceiveMessage(opts *bind.FilterOpts, originChainID [][32]byte, originSenderAddress []common.Address) (*ExampleCrossChainMessengerReceiveMessageIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.FilterLogs(opts, "ReceiveMessage", originChainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerReceiveMessageIterator{contract: _ExampleCrossChainMessenger.contract, event: "ReceiveMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveMessage is a free log subscription operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed originChainID, address indexed originSenderAddress, string message)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) WatchReceiveMessage(opts *bind.WatchOpts, sink chan<- *ExampleCrossChainMessengerReceiveMessage, originChainID [][32]byte, originSenderAddress []common.Address) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.WatchLogs(opts, "ReceiveMessage", originChainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleCrossChainMessengerReceiveMessage)
				if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "ReceiveMessage", log); err != nil {
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

// ParseReceiveMessage is a log parse operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed originChainID, address indexed originSenderAddress, string message)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) ParseReceiveMessage(log types.Log) (*ExampleCrossChainMessengerReceiveMessage, error) {
	event := new(ExampleCrossChainMessengerReceiveMessage)
	if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "ReceiveMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleCrossChainMessengerSendMessageIterator is returned from FilterSendMessage and is used to iterate over the raw logs and unpacked data for SendMessage events raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerSendMessageIterator struct {
	Event *ExampleCrossChainMessengerSendMessage // Event containing the contract specifics and raw log

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
func (it *ExampleCrossChainMessengerSendMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleCrossChainMessengerSendMessage)
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
		it.Event = new(ExampleCrossChainMessengerSendMessage)
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
func (it *ExampleCrossChainMessengerSendMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleCrossChainMessengerSendMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleCrossChainMessengerSendMessage represents a SendMessage event raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerSendMessage struct {
	DestinationChainID [32]byte
	DestinationAddress common.Address
	FeeTokenAddress    common.Address
	FeeAmount          *big.Int
	RequiredGasLimit   *big.Int
	Message            string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSendMessage is a free log retrieval operation binding the contract event 0xa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca.
//
// Solidity: event SendMessage(bytes32 indexed destinationChainID, address indexed destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) FilterSendMessage(opts *bind.FilterOpts, destinationChainID [][32]byte, destinationAddress []common.Address) (*ExampleCrossChainMessengerSendMessageIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.FilterLogs(opts, "SendMessage", destinationChainIDRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerSendMessageIterator{contract: _ExampleCrossChainMessenger.contract, event: "SendMessage", logs: logs, sub: sub}, nil
}

// WatchSendMessage is a free log subscription operation binding the contract event 0xa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca.
//
// Solidity: event SendMessage(bytes32 indexed destinationChainID, address indexed destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) WatchSendMessage(opts *bind.WatchOpts, sink chan<- *ExampleCrossChainMessengerSendMessage, destinationChainID [][32]byte, destinationAddress []common.Address) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.WatchLogs(opts, "SendMessage", destinationChainIDRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleCrossChainMessengerSendMessage)
				if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "SendMessage", log); err != nil {
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

// ParseSendMessage is a log parse operation binding the contract event 0xa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca.
//
// Solidity: event SendMessage(bytes32 indexed destinationChainID, address indexed destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) ParseSendMessage(log types.Log) (*ExampleCrossChainMessengerSendMessage, error) {
	event := new(ExampleCrossChainMessengerSendMessage)
	if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "SendMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
