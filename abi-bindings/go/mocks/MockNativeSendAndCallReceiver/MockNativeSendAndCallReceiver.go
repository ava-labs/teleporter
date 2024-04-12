// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocknativesendandcallreceiver

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

// MockNativeSendAndCallReceiverMetaData contains all meta data concerning the MockNativeSendAndCallReceiver contract.
var MockNativeSendAndCallReceiverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"TokensReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"receiveTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506101be806100206000396000f3fe60806040526004361061001d5760003560e01c80628a403e14610022575b600080fd5b6100356100303660046100e0565b610037565b005b7f7aaf2e01aab6b767f76075e0fefc4618cf6b4a3b7597d1e81b914eb66175655b34838360405161006a93929190610152565b60405180910390a160008190036100dc5760405162461bcd60e51b815260206004820152602c60248201527f4d6f636b4e617469766553656e64416e6443616c6c52656365697665723a206560448201526b1b5c1d1e481c185e5b1bd85960a21b606482015260840160405180910390fd5b5050565b600080602083850312156100f357600080fd5b823567ffffffffffffffff8082111561010b57600080fd5b818501915085601f83011261011f57600080fd5b81358181111561012e57600080fd5b86602082850101111561014057600080fd5b60209290920196919550909350505050565b83815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f191601019291505056fea26469706673582212203b9a7303a3c2148c218936e8bd84202c5b69e3bf6350af291764eb900955097064736f6c63430008120033",
}

// MockNativeSendAndCallReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MockNativeSendAndCallReceiverMetaData.ABI instead.
var MockNativeSendAndCallReceiverABI = MockNativeSendAndCallReceiverMetaData.ABI

// MockNativeSendAndCallReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockNativeSendAndCallReceiverMetaData.Bin instead.
var MockNativeSendAndCallReceiverBin = MockNativeSendAndCallReceiverMetaData.Bin

// DeployMockNativeSendAndCallReceiver deploys a new Ethereum contract, binding an instance of MockNativeSendAndCallReceiver to it.
func DeployMockNativeSendAndCallReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockNativeSendAndCallReceiver, error) {
	parsed, err := MockNativeSendAndCallReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockNativeSendAndCallReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockNativeSendAndCallReceiver{MockNativeSendAndCallReceiverCaller: MockNativeSendAndCallReceiverCaller{contract: contract}, MockNativeSendAndCallReceiverTransactor: MockNativeSendAndCallReceiverTransactor{contract: contract}, MockNativeSendAndCallReceiverFilterer: MockNativeSendAndCallReceiverFilterer{contract: contract}}, nil
}

// MockNativeSendAndCallReceiver is an auto generated Go binding around an Ethereum contract.
type MockNativeSendAndCallReceiver struct {
	MockNativeSendAndCallReceiverCaller     // Read-only binding to the contract
	MockNativeSendAndCallReceiverTransactor // Write-only binding to the contract
	MockNativeSendAndCallReceiverFilterer   // Log filterer for contract events
}

// MockNativeSendAndCallReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockNativeSendAndCallReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockNativeSendAndCallReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockNativeSendAndCallReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockNativeSendAndCallReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockNativeSendAndCallReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockNativeSendAndCallReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockNativeSendAndCallReceiverSession struct {
	Contract     *MockNativeSendAndCallReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MockNativeSendAndCallReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockNativeSendAndCallReceiverCallerSession struct {
	Contract *MockNativeSendAndCallReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// MockNativeSendAndCallReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockNativeSendAndCallReceiverTransactorSession struct {
	Contract     *MockNativeSendAndCallReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// MockNativeSendAndCallReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockNativeSendAndCallReceiverRaw struct {
	Contract *MockNativeSendAndCallReceiver // Generic contract binding to access the raw methods on
}

// MockNativeSendAndCallReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockNativeSendAndCallReceiverCallerRaw struct {
	Contract *MockNativeSendAndCallReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// MockNativeSendAndCallReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockNativeSendAndCallReceiverTransactorRaw struct {
	Contract *MockNativeSendAndCallReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockNativeSendAndCallReceiver creates a new instance of MockNativeSendAndCallReceiver, bound to a specific deployed contract.
func NewMockNativeSendAndCallReceiver(address common.Address, backend bind.ContractBackend) (*MockNativeSendAndCallReceiver, error) {
	contract, err := bindMockNativeSendAndCallReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockNativeSendAndCallReceiver{MockNativeSendAndCallReceiverCaller: MockNativeSendAndCallReceiverCaller{contract: contract}, MockNativeSendAndCallReceiverTransactor: MockNativeSendAndCallReceiverTransactor{contract: contract}, MockNativeSendAndCallReceiverFilterer: MockNativeSendAndCallReceiverFilterer{contract: contract}}, nil
}

// NewMockNativeSendAndCallReceiverCaller creates a new read-only instance of MockNativeSendAndCallReceiver, bound to a specific deployed contract.
func NewMockNativeSendAndCallReceiverCaller(address common.Address, caller bind.ContractCaller) (*MockNativeSendAndCallReceiverCaller, error) {
	contract, err := bindMockNativeSendAndCallReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockNativeSendAndCallReceiverCaller{contract: contract}, nil
}

// NewMockNativeSendAndCallReceiverTransactor creates a new write-only instance of MockNativeSendAndCallReceiver, bound to a specific deployed contract.
func NewMockNativeSendAndCallReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MockNativeSendAndCallReceiverTransactor, error) {
	contract, err := bindMockNativeSendAndCallReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockNativeSendAndCallReceiverTransactor{contract: contract}, nil
}

// NewMockNativeSendAndCallReceiverFilterer creates a new log filterer instance of MockNativeSendAndCallReceiver, bound to a specific deployed contract.
func NewMockNativeSendAndCallReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MockNativeSendAndCallReceiverFilterer, error) {
	contract, err := bindMockNativeSendAndCallReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockNativeSendAndCallReceiverFilterer{contract: contract}, nil
}

// bindMockNativeSendAndCallReceiver binds a generic wrapper to an already deployed contract.
func bindMockNativeSendAndCallReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockNativeSendAndCallReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockNativeSendAndCallReceiver.Contract.MockNativeSendAndCallReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.MockNativeSendAndCallReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.MockNativeSendAndCallReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockNativeSendAndCallReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.contract.Transact(opts, method, params...)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x008a403e.
//
// Solidity: function receiveTokens(bytes payload) payable returns()
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverTransactor) ReceiveTokens(opts *bind.TransactOpts, payload []byte) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.contract.Transact(opts, "receiveTokens", payload)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x008a403e.
//
// Solidity: function receiveTokens(bytes payload) payable returns()
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverSession) ReceiveTokens(payload []byte) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.ReceiveTokens(&_MockNativeSendAndCallReceiver.TransactOpts, payload)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x008a403e.
//
// Solidity: function receiveTokens(bytes payload) payable returns()
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverTransactorSession) ReceiveTokens(payload []byte) (*types.Transaction, error) {
	return _MockNativeSendAndCallReceiver.Contract.ReceiveTokens(&_MockNativeSendAndCallReceiver.TransactOpts, payload)
}

// MockNativeSendAndCallReceiverTokensReceivedIterator is returned from FilterTokensReceived and is used to iterate over the raw logs and unpacked data for TokensReceived events raised by the MockNativeSendAndCallReceiver contract.
type MockNativeSendAndCallReceiverTokensReceivedIterator struct {
	Event *MockNativeSendAndCallReceiverTokensReceived // Event containing the contract specifics and raw log

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
func (it *MockNativeSendAndCallReceiverTokensReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockNativeSendAndCallReceiverTokensReceived)
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
		it.Event = new(MockNativeSendAndCallReceiverTokensReceived)
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
func (it *MockNativeSendAndCallReceiverTokensReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockNativeSendAndCallReceiverTokensReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockNativeSendAndCallReceiverTokensReceived represents a TokensReceived event raised by the MockNativeSendAndCallReceiver contract.
type MockNativeSendAndCallReceiverTokensReceived struct {
	Amount  *big.Int
	Payload []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokensReceived is a free log retrieval operation binding the contract event 0x7aaf2e01aab6b767f76075e0fefc4618cf6b4a3b7597d1e81b914eb66175655b.
//
// Solidity: event TokensReceived(uint256 amount, bytes payload)
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverFilterer) FilterTokensReceived(opts *bind.FilterOpts) (*MockNativeSendAndCallReceiverTokensReceivedIterator, error) {

	logs, sub, err := _MockNativeSendAndCallReceiver.contract.FilterLogs(opts, "TokensReceived")
	if err != nil {
		return nil, err
	}
	return &MockNativeSendAndCallReceiverTokensReceivedIterator{contract: _MockNativeSendAndCallReceiver.contract, event: "TokensReceived", logs: logs, sub: sub}, nil
}

// WatchTokensReceived is a free log subscription operation binding the contract event 0x7aaf2e01aab6b767f76075e0fefc4618cf6b4a3b7597d1e81b914eb66175655b.
//
// Solidity: event TokensReceived(uint256 amount, bytes payload)
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverFilterer) WatchTokensReceived(opts *bind.WatchOpts, sink chan<- *MockNativeSendAndCallReceiverTokensReceived) (event.Subscription, error) {

	logs, sub, err := _MockNativeSendAndCallReceiver.contract.WatchLogs(opts, "TokensReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockNativeSendAndCallReceiverTokensReceived)
				if err := _MockNativeSendAndCallReceiver.contract.UnpackLog(event, "TokensReceived", log); err != nil {
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

// ParseTokensReceived is a log parse operation binding the contract event 0x7aaf2e01aab6b767f76075e0fefc4618cf6b4a3b7597d1e81b914eb66175655b.
//
// Solidity: event TokensReceived(uint256 amount, bytes payload)
func (_MockNativeSendAndCallReceiver *MockNativeSendAndCallReceiverFilterer) ParseTokensReceived(log types.Log) (*MockNativeSendAndCallReceiverTokensReceived, error) {
	event := new(MockNativeSendAndCallReceiverTokensReceived)
	if err := _MockNativeSendAndCallReceiver.contract.UnpackLog(event, "TokensReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
