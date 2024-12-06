// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bar

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

// BarMetaData contains all meta data concerning the Bar contract.
var BarMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BAR_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setBar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506101198061001c5f395ff3fe6080604052348015600e575f80fd5b5060043610603a575f3560e01c8063352d3fba14603e5780634321d9d214606f57806367a23d131460a7575b5f80fd5b606d604936600460cd565b7fab087b855a66dd5f7ac611eac86e092a2f3fc08a745c96afb2bc0553ff79ea0055565b005b60957fab087b855a66dd5f7ac611eac86e092a2f3fc08a745c96afb2bc0553ff79ea0081565b60405190815260200160405180910390f35b7fab087b855a66dd5f7ac611eac86e092a2f3fc08a745c96afb2bc0553ff79ea00546095565b5f6020828403121560dc575f80fd5b503591905056fea26469706673582212200f2ca6707e4c64e7279bab2759d11b01bd05a1ec8e5e2fc760ee8b07b0ff293d64736f6c63430008190033",
}

// BarABI is the input ABI used to generate the binding from.
// Deprecated: Use BarMetaData.ABI instead.
var BarABI = BarMetaData.ABI

// BarBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BarMetaData.Bin instead.
var BarBin = BarMetaData.Bin

// DeployBar deploys a new Ethereum contract, binding an instance of Bar to it.
func DeployBar(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bar, error) {
	parsed, err := BarMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BarBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bar{BarCaller: BarCaller{contract: contract}, BarTransactor: BarTransactor{contract: contract}, BarFilterer: BarFilterer{contract: contract}}, nil
}

// Bar is an auto generated Go binding around an Ethereum contract.
type Bar struct {
	BarCaller     // Read-only binding to the contract
	BarTransactor // Write-only binding to the contract
	BarFilterer   // Log filterer for contract events
}

// BarCaller is an auto generated read-only Go binding around an Ethereum contract.
type BarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BarSession struct {
	Contract     *Bar              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BarCallerSession struct {
	Contract *BarCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BarTransactorSession struct {
	Contract     *BarTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BarRaw is an auto generated low-level Go binding around an Ethereum contract.
type BarRaw struct {
	Contract *Bar // Generic contract binding to access the raw methods on
}

// BarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BarCallerRaw struct {
	Contract *BarCaller // Generic read-only contract binding to access the raw methods on
}

// BarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BarTransactorRaw struct {
	Contract *BarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBar creates a new instance of Bar, bound to a specific deployed contract.
func NewBar(address common.Address, backend bind.ContractBackend) (*Bar, error) {
	contract, err := bindBar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bar{BarCaller: BarCaller{contract: contract}, BarTransactor: BarTransactor{contract: contract}, BarFilterer: BarFilterer{contract: contract}}, nil
}

// NewBarCaller creates a new read-only instance of Bar, bound to a specific deployed contract.
func NewBarCaller(address common.Address, caller bind.ContractCaller) (*BarCaller, error) {
	contract, err := bindBar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BarCaller{contract: contract}, nil
}

// NewBarTransactor creates a new write-only instance of Bar, bound to a specific deployed contract.
func NewBarTransactor(address common.Address, transactor bind.ContractTransactor) (*BarTransactor, error) {
	contract, err := bindBar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BarTransactor{contract: contract}, nil
}

// NewBarFilterer creates a new log filterer instance of Bar, bound to a specific deployed contract.
func NewBarFilterer(address common.Address, filterer bind.ContractFilterer) (*BarFilterer, error) {
	contract, err := bindBar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BarFilterer{contract: contract}, nil
}

// bindBar binds a generic wrapper to an already deployed contract.
func bindBar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bar *BarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bar.Contract.BarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bar *BarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bar.Contract.BarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bar *BarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bar.Contract.BarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bar *BarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bar *BarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bar *BarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bar.Contract.contract.Transact(opts, method, params...)
}

// BARSTORAGELOCATION is a free data retrieval call binding the contract method 0x4321d9d2.
//
// Solidity: function BAR_STORAGE_LOCATION() view returns(bytes32)
func (_Bar *BarCaller) BARSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bar.contract.Call(opts, &out, "BAR_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BARSTORAGELOCATION is a free data retrieval call binding the contract method 0x4321d9d2.
//
// Solidity: function BAR_STORAGE_LOCATION() view returns(bytes32)
func (_Bar *BarSession) BARSTORAGELOCATION() ([32]byte, error) {
	return _Bar.Contract.BARSTORAGELOCATION(&_Bar.CallOpts)
}

// BARSTORAGELOCATION is a free data retrieval call binding the contract method 0x4321d9d2.
//
// Solidity: function BAR_STORAGE_LOCATION() view returns(bytes32)
func (_Bar *BarCallerSession) BARSTORAGELOCATION() ([32]byte, error) {
	return _Bar.Contract.BARSTORAGELOCATION(&_Bar.CallOpts)
}

// GetBar is a free data retrieval call binding the contract method 0x67a23d13.
//
// Solidity: function getBar() view returns(uint256)
func (_Bar *BarCaller) GetBar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bar.contract.Call(opts, &out, "getBar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBar is a free data retrieval call binding the contract method 0x67a23d13.
//
// Solidity: function getBar() view returns(uint256)
func (_Bar *BarSession) GetBar() (*big.Int, error) {
	return _Bar.Contract.GetBar(&_Bar.CallOpts)
}

// GetBar is a free data retrieval call binding the contract method 0x67a23d13.
//
// Solidity: function getBar() view returns(uint256)
func (_Bar *BarCallerSession) GetBar() (*big.Int, error) {
	return _Bar.Contract.GetBar(&_Bar.CallOpts)
}

// SetBar is a paid mutator transaction binding the contract method 0x352d3fba.
//
// Solidity: function setBar(uint256 value) returns()
func (_Bar *BarTransactor) SetBar(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Bar.contract.Transact(opts, "setBar", value)
}

// SetBar is a paid mutator transaction binding the contract method 0x352d3fba.
//
// Solidity: function setBar(uint256 value) returns()
func (_Bar *BarSession) SetBar(value *big.Int) (*types.Transaction, error) {
	return _Bar.Contract.SetBar(&_Bar.TransactOpts, value)
}

// SetBar is a paid mutator transaction binding the contract method 0x352d3fba.
//
// Solidity: function setBar(uint256 value) returns()
func (_Bar *BarTransactorSession) SetBar(value *big.Int) (*types.Transaction, error) {
	return _Bar.Contract.SetBar(&_Bar.TransactOpts, value)
}

// BarInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bar contract.
type BarInitializedIterator struct {
	Event *BarInitialized // Event containing the contract specifics and raw log

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
func (it *BarInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BarInitialized)
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
		it.Event = new(BarInitialized)
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
func (it *BarInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BarInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BarInitialized represents a Initialized event raised by the Bar contract.
type BarInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bar *BarFilterer) FilterInitialized(opts *bind.FilterOpts) (*BarInitializedIterator, error) {

	logs, sub, err := _Bar.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BarInitializedIterator{contract: _Bar.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bar *BarFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BarInitialized) (event.Subscription, error) {

	logs, sub, err := _Bar.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BarInitialized)
				if err := _Bar.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bar *BarFilterer) ParseInitialized(log types.Log) (*BarInitialized, error) {
	event := new(BarInitialized)
	if err := _Bar.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
