// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package foo

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

// FooMetaData contains all meta data concerning the Foo contract.
var FooMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BAR_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FOO_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFoo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setBar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setFoo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b5061037b8061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061007a575f3560e01c80634321d9d2116100585780634321d9d2146100d457806367a23d13146100fb5780638129fc1c14610103578063dc80035d1461010b575f80fd5b8063243dc8da1461007e578063352d3fba1461009857806337c7d9fb146100ad575b5f80fd5b61008661011e565b60405190815260200160405180910390f35b6100ab6100a636600461032e565b610147565b005b6100867f92c6022ab2e4d8395662f3b814d5b9a489d615215f1269fc366f81bf3a17f70081565b6100867fab087b855a66dd5f7ac611eac86e092a2f3fc08a745c96afb2bc0553ff79ea0081565b610086610171565b6100ab610198565b6100ab61011936600461032e565b6102a4565b5f7f92c6022ab2e4d8395662f3b814d5b9a489d615215f1269fc366f81bf3a17f7005b54919050565b5f7fab087b855a66dd5f7ac611eac86e092a2f3fc08a745c96afb2bc0553ff79ea005b9190915550565b5f7fab087b855a66dd5f7ac611eac86e092a2f3fc08a745c96afb2bc0553ff79ea00610141565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156101dd5750825b90505f8267ffffffffffffffff1660011480156101f95750303b155b905081158015610207575080155b156102255760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561024f57845460ff60401b1916600160401b1785555b6102576102cb565b831561029d57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050565b5f7f92c6022ab2e4d8395662f3b814d5b9a489d615215f1269fc366f81bf3a17f70061016a565b6102d36102dd565b6102db610326565b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166102db57604051631afcd79f60e31b815260040160405180910390fd5b6102db6102dd565b5f6020828403121561033e575f80fd5b503591905056fea26469706673582212205f2b29b10a01165f2d0b46155421e8c14d58e86f686d6d8b51d5666918fe283864736f6c63430008190033",
}

// FooABI is the input ABI used to generate the binding from.
// Deprecated: Use FooMetaData.ABI instead.
var FooABI = FooMetaData.ABI

// FooBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FooMetaData.Bin instead.
var FooBin = FooMetaData.Bin

// DeployFoo deploys a new Ethereum contract, binding an instance of Foo to it.
func DeployFoo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Foo, error) {
	parsed, err := FooMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FooBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Foo{FooCaller: FooCaller{contract: contract}, FooTransactor: FooTransactor{contract: contract}, FooFilterer: FooFilterer{contract: contract}}, nil
}

// Foo is an auto generated Go binding around an Ethereum contract.
type Foo struct {
	FooCaller     // Read-only binding to the contract
	FooTransactor // Write-only binding to the contract
	FooFilterer   // Log filterer for contract events
}

// FooCaller is an auto generated read-only Go binding around an Ethereum contract.
type FooCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FooTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FooTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FooFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FooFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FooSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FooSession struct {
	Contract     *Foo              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FooCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FooCallerSession struct {
	Contract *FooCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FooTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FooTransactorSession struct {
	Contract     *FooTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FooRaw is an auto generated low-level Go binding around an Ethereum contract.
type FooRaw struct {
	Contract *Foo // Generic contract binding to access the raw methods on
}

// FooCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FooCallerRaw struct {
	Contract *FooCaller // Generic read-only contract binding to access the raw methods on
}

// FooTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FooTransactorRaw struct {
	Contract *FooTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFoo creates a new instance of Foo, bound to a specific deployed contract.
func NewFoo(address common.Address, backend bind.ContractBackend) (*Foo, error) {
	contract, err := bindFoo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Foo{FooCaller: FooCaller{contract: contract}, FooTransactor: FooTransactor{contract: contract}, FooFilterer: FooFilterer{contract: contract}}, nil
}

// NewFooCaller creates a new read-only instance of Foo, bound to a specific deployed contract.
func NewFooCaller(address common.Address, caller bind.ContractCaller) (*FooCaller, error) {
	contract, err := bindFoo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FooCaller{contract: contract}, nil
}

// NewFooTransactor creates a new write-only instance of Foo, bound to a specific deployed contract.
func NewFooTransactor(address common.Address, transactor bind.ContractTransactor) (*FooTransactor, error) {
	contract, err := bindFoo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FooTransactor{contract: contract}, nil
}

// NewFooFilterer creates a new log filterer instance of Foo, bound to a specific deployed contract.
func NewFooFilterer(address common.Address, filterer bind.ContractFilterer) (*FooFilterer, error) {
	contract, err := bindFoo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FooFilterer{contract: contract}, nil
}

// bindFoo binds a generic wrapper to an already deployed contract.
func bindFoo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FooMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foo *FooRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Foo.Contract.FooCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foo *FooRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foo.Contract.FooTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foo *FooRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foo.Contract.FooTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foo *FooCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Foo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foo *FooTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foo *FooTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foo.Contract.contract.Transact(opts, method, params...)
}

// BARSTORAGELOCATION is a free data retrieval call binding the contract method 0x4321d9d2.
//
// Solidity: function BAR_STORAGE_LOCATION() view returns(bytes32)
func (_Foo *FooCaller) BARSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Foo.contract.Call(opts, &out, "BAR_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BARSTORAGELOCATION is a free data retrieval call binding the contract method 0x4321d9d2.
//
// Solidity: function BAR_STORAGE_LOCATION() view returns(bytes32)
func (_Foo *FooSession) BARSTORAGELOCATION() ([32]byte, error) {
	return _Foo.Contract.BARSTORAGELOCATION(&_Foo.CallOpts)
}

// BARSTORAGELOCATION is a free data retrieval call binding the contract method 0x4321d9d2.
//
// Solidity: function BAR_STORAGE_LOCATION() view returns(bytes32)
func (_Foo *FooCallerSession) BARSTORAGELOCATION() ([32]byte, error) {
	return _Foo.Contract.BARSTORAGELOCATION(&_Foo.CallOpts)
}

// FOOSTORAGELOCATION is a free data retrieval call binding the contract method 0x37c7d9fb.
//
// Solidity: function FOO_STORAGE_LOCATION() view returns(bytes32)
func (_Foo *FooCaller) FOOSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Foo.contract.Call(opts, &out, "FOO_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FOOSTORAGELOCATION is a free data retrieval call binding the contract method 0x37c7d9fb.
//
// Solidity: function FOO_STORAGE_LOCATION() view returns(bytes32)
func (_Foo *FooSession) FOOSTORAGELOCATION() ([32]byte, error) {
	return _Foo.Contract.FOOSTORAGELOCATION(&_Foo.CallOpts)
}

// FOOSTORAGELOCATION is a free data retrieval call binding the contract method 0x37c7d9fb.
//
// Solidity: function FOO_STORAGE_LOCATION() view returns(bytes32)
func (_Foo *FooCallerSession) FOOSTORAGELOCATION() ([32]byte, error) {
	return _Foo.Contract.FOOSTORAGELOCATION(&_Foo.CallOpts)
}

// GetBar is a free data retrieval call binding the contract method 0x67a23d13.
//
// Solidity: function getBar() view returns(uint256)
func (_Foo *FooCaller) GetBar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Foo.contract.Call(opts, &out, "getBar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBar is a free data retrieval call binding the contract method 0x67a23d13.
//
// Solidity: function getBar() view returns(uint256)
func (_Foo *FooSession) GetBar() (*big.Int, error) {
	return _Foo.Contract.GetBar(&_Foo.CallOpts)
}

// GetBar is a free data retrieval call binding the contract method 0x67a23d13.
//
// Solidity: function getBar() view returns(uint256)
func (_Foo *FooCallerSession) GetBar() (*big.Int, error) {
	return _Foo.Contract.GetBar(&_Foo.CallOpts)
}

// GetFoo is a free data retrieval call binding the contract method 0x243dc8da.
//
// Solidity: function getFoo() view returns(uint256)
func (_Foo *FooCaller) GetFoo(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Foo.contract.Call(opts, &out, "getFoo")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFoo is a free data retrieval call binding the contract method 0x243dc8da.
//
// Solidity: function getFoo() view returns(uint256)
func (_Foo *FooSession) GetFoo() (*big.Int, error) {
	return _Foo.Contract.GetFoo(&_Foo.CallOpts)
}

// GetFoo is a free data retrieval call binding the contract method 0x243dc8da.
//
// Solidity: function getFoo() view returns(uint256)
func (_Foo *FooCallerSession) GetFoo() (*big.Int, error) {
	return _Foo.Contract.GetFoo(&_Foo.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Foo *FooTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foo.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Foo *FooSession) Initialize() (*types.Transaction, error) {
	return _Foo.Contract.Initialize(&_Foo.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Foo *FooTransactorSession) Initialize() (*types.Transaction, error) {
	return _Foo.Contract.Initialize(&_Foo.TransactOpts)
}

// SetBar is a paid mutator transaction binding the contract method 0x352d3fba.
//
// Solidity: function setBar(uint256 value) returns()
func (_Foo *FooTransactor) SetBar(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Foo.contract.Transact(opts, "setBar", value)
}

// SetBar is a paid mutator transaction binding the contract method 0x352d3fba.
//
// Solidity: function setBar(uint256 value) returns()
func (_Foo *FooSession) SetBar(value *big.Int) (*types.Transaction, error) {
	return _Foo.Contract.SetBar(&_Foo.TransactOpts, value)
}

// SetBar is a paid mutator transaction binding the contract method 0x352d3fba.
//
// Solidity: function setBar(uint256 value) returns()
func (_Foo *FooTransactorSession) SetBar(value *big.Int) (*types.Transaction, error) {
	return _Foo.Contract.SetBar(&_Foo.TransactOpts, value)
}

// SetFoo is a paid mutator transaction binding the contract method 0xdc80035d.
//
// Solidity: function setFoo(uint256 value) returns()
func (_Foo *FooTransactor) SetFoo(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Foo.contract.Transact(opts, "setFoo", value)
}

// SetFoo is a paid mutator transaction binding the contract method 0xdc80035d.
//
// Solidity: function setFoo(uint256 value) returns()
func (_Foo *FooSession) SetFoo(value *big.Int) (*types.Transaction, error) {
	return _Foo.Contract.SetFoo(&_Foo.TransactOpts, value)
}

// SetFoo is a paid mutator transaction binding the contract method 0xdc80035d.
//
// Solidity: function setFoo(uint256 value) returns()
func (_Foo *FooTransactorSession) SetFoo(value *big.Int) (*types.Transaction, error) {
	return _Foo.Contract.SetFoo(&_Foo.TransactOpts, value)
}

// FooInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Foo contract.
type FooInitializedIterator struct {
	Event *FooInitialized // Event containing the contract specifics and raw log

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
func (it *FooInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FooInitialized)
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
		it.Event = new(FooInitialized)
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
func (it *FooInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FooInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FooInitialized represents a Initialized event raised by the Foo contract.
type FooInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Foo *FooFilterer) FilterInitialized(opts *bind.FilterOpts) (*FooInitializedIterator, error) {

	logs, sub, err := _Foo.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FooInitializedIterator{contract: _Foo.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Foo *FooFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FooInitialized) (event.Subscription, error) {

	logs, sub, err := _Foo.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FooInitialized)
				if err := _Foo.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Foo *FooFilterer) ParseInitialized(log types.Log) (*FooInitialized, error) {
	event := new(FooInitialized)
	if err := _Foo.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
