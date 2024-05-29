// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iteleporterreceiver

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

// ITeleporterReceiverMetaData contains all meta data concerning the ITeleporterReceiver contract.
var ITeleporterReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITeleporterReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use ITeleporterReceiverMetaData.ABI instead.
var ITeleporterReceiverABI = ITeleporterReceiverMetaData.ABI

// ITeleporterReceiver is an auto generated Go binding around an Ethereum contract.
type ITeleporterReceiver struct {
	ITeleporterReceiverCaller     // Read-only binding to the contract
	ITeleporterReceiverTransactor // Write-only binding to the contract
	ITeleporterReceiverFilterer   // Log filterer for contract events
}

// ITeleporterReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITeleporterReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITeleporterReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITeleporterReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITeleporterReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITeleporterReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITeleporterReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITeleporterReceiverSession struct {
	Contract     *ITeleporterReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ITeleporterReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITeleporterReceiverCallerSession struct {
	Contract *ITeleporterReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ITeleporterReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITeleporterReceiverTransactorSession struct {
	Contract     *ITeleporterReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ITeleporterReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITeleporterReceiverRaw struct {
	Contract *ITeleporterReceiver // Generic contract binding to access the raw methods on
}

// ITeleporterReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITeleporterReceiverCallerRaw struct {
	Contract *ITeleporterReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ITeleporterReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITeleporterReceiverTransactorRaw struct {
	Contract *ITeleporterReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITeleporterReceiver creates a new instance of ITeleporterReceiver, bound to a specific deployed contract.
func NewITeleporterReceiver(address common.Address, backend bind.ContractBackend) (*ITeleporterReceiver, error) {
	contract, err := bindITeleporterReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITeleporterReceiver{ITeleporterReceiverCaller: ITeleporterReceiverCaller{contract: contract}, ITeleporterReceiverTransactor: ITeleporterReceiverTransactor{contract: contract}, ITeleporterReceiverFilterer: ITeleporterReceiverFilterer{contract: contract}}, nil
}

// NewITeleporterReceiverCaller creates a new read-only instance of ITeleporterReceiver, bound to a specific deployed contract.
func NewITeleporterReceiverCaller(address common.Address, caller bind.ContractCaller) (*ITeleporterReceiverCaller, error) {
	contract, err := bindITeleporterReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITeleporterReceiverCaller{contract: contract}, nil
}

// NewITeleporterReceiverTransactor creates a new write-only instance of ITeleporterReceiver, bound to a specific deployed contract.
func NewITeleporterReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ITeleporterReceiverTransactor, error) {
	contract, err := bindITeleporterReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITeleporterReceiverTransactor{contract: contract}, nil
}

// NewITeleporterReceiverFilterer creates a new log filterer instance of ITeleporterReceiver, bound to a specific deployed contract.
func NewITeleporterReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ITeleporterReceiverFilterer, error) {
	contract, err := bindITeleporterReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITeleporterReceiverFilterer{contract: contract}, nil
}

// bindITeleporterReceiver binds a generic wrapper to an already deployed contract.
func bindITeleporterReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITeleporterReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITeleporterReceiver *ITeleporterReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITeleporterReceiver.Contract.ITeleporterReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITeleporterReceiver *ITeleporterReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITeleporterReceiver.Contract.ITeleporterReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITeleporterReceiver *ITeleporterReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITeleporterReceiver.Contract.ITeleporterReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITeleporterReceiver *ITeleporterReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITeleporterReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITeleporterReceiver *ITeleporterReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITeleporterReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITeleporterReceiver *ITeleporterReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITeleporterReceiver.Contract.contract.Transact(opts, method, params...)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ITeleporterReceiver *ITeleporterReceiverTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ITeleporterReceiver.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ITeleporterReceiver *ITeleporterReceiverSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ITeleporterReceiver.Contract.ReceiveTeleporterMessage(&_ITeleporterReceiver.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ITeleporterReceiver *ITeleporterReceiverTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ITeleporterReceiver.Contract.ReceiveTeleporterMessage(&_ITeleporterReceiver.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}
