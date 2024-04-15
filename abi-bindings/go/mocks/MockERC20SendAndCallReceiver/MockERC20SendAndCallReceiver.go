// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockerc20sendandcallreceiver

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

// MockERC20SendAndCallReceiverMetaData contains all meta data concerning the MockERC20SendAndCallReceiver contract.
var MockERC20SendAndCallReceiverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"TokensReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"receiveTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610713806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80638bc1e07614610030575b600080fd5b61004361003e36600461052a565b610045565b005b7f0f611859dc94a66d9bec0e66c7485098c0d5e4662896ff2f523c47f962a7abef8484848460405161007a94939291906105bf565b60405180910390a1806100e85760405162461bcd60e51b815260206004820152602b60248201527f4d6f636b455243323053656e64416e6443616c6c52656365697665723a20656d60448201526a1c1d1e481c185e5b1bd85960aa1b60648201526084015b60405180910390fd5b6100f284846100f9565b5050505050565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015610142573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101669190610607565b905061017d6001600160a01b038516333086610263565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa1580156101c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101e89190610607565b905081811161024e5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016100df565b6102588282610620565b925050505b92915050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526102bd9085906102c3565b50505050565b6000610318826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661039a9092919063ffffffff16565b80519091501561039557808060200190518101906103369190610641565b6103955760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016100df565b505050565b60606103a984846000856103b1565b949350505050565b6060824710156104125760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016100df565b600080866001600160a01b0316858760405161042e919061068e565b60006040518083038185875af1925050503d806000811461046b576040519150601f19603f3d011682016040523d82523d6000602084013e610470565b606091505b50915091506104818783838761048c565b979650505050505050565b606083156104fb5782516000036104f4576001600160a01b0385163b6104f45760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016100df565b50816103a9565b6103a983838151156105105781518083602001fd5b8060405162461bcd60e51b81526004016100df91906106aa565b6000806000806060858703121561054057600080fd5b84356001600160a01b038116811461055757600080fd5b935060208501359250604085013567ffffffffffffffff8082111561057b57600080fd5b818701915087601f83011261058f57600080fd5b81358181111561059e57600080fd5b8860208285010111156105b057600080fd5b95989497505060200194505050565b6001600160a01b0385168152602081018490526060604082018190528101829052818360808301376000818301608090810191909152601f909201601f191601019392505050565b60006020828403121561061957600080fd5b5051919050565b8181038181111561025d57634e487b7160e01b600052601160045260246000fd5b60006020828403121561065357600080fd5b8151801515811461066357600080fd5b9392505050565b60005b8381101561068557818101518382015260200161066d565b50506000910152565b600082516106a081846020870161066a565b9190910192915050565b60208152600082518060208401526106c981604085016020870161066a565b601f01601f1916919091016040019291505056fea26469706673582212205a0b338b523b794463c94473e21b518e9749c98f9ddcd2a2c53cc46b805fa12264736f6c63430008120033",
}

// MockERC20SendAndCallReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MockERC20SendAndCallReceiverMetaData.ABI instead.
var MockERC20SendAndCallReceiverABI = MockERC20SendAndCallReceiverMetaData.ABI

// MockERC20SendAndCallReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockERC20SendAndCallReceiverMetaData.Bin instead.
var MockERC20SendAndCallReceiverBin = MockERC20SendAndCallReceiverMetaData.Bin

// DeployMockERC20SendAndCallReceiver deploys a new Ethereum contract, binding an instance of MockERC20SendAndCallReceiver to it.
func DeployMockERC20SendAndCallReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockERC20SendAndCallReceiver, error) {
	parsed, err := MockERC20SendAndCallReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockERC20SendAndCallReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockERC20SendAndCallReceiver{MockERC20SendAndCallReceiverCaller: MockERC20SendAndCallReceiverCaller{contract: contract}, MockERC20SendAndCallReceiverTransactor: MockERC20SendAndCallReceiverTransactor{contract: contract}, MockERC20SendAndCallReceiverFilterer: MockERC20SendAndCallReceiverFilterer{contract: contract}}, nil
}

// MockERC20SendAndCallReceiver is an auto generated Go binding around an Ethereum contract.
type MockERC20SendAndCallReceiver struct {
	MockERC20SendAndCallReceiverCaller     // Read-only binding to the contract
	MockERC20SendAndCallReceiverTransactor // Write-only binding to the contract
	MockERC20SendAndCallReceiverFilterer   // Log filterer for contract events
}

// MockERC20SendAndCallReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockERC20SendAndCallReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC20SendAndCallReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockERC20SendAndCallReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC20SendAndCallReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockERC20SendAndCallReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC20SendAndCallReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockERC20SendAndCallReceiverSession struct {
	Contract     *MockERC20SendAndCallReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MockERC20SendAndCallReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockERC20SendAndCallReceiverCallerSession struct {
	Contract *MockERC20SendAndCallReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// MockERC20SendAndCallReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockERC20SendAndCallReceiverTransactorSession struct {
	Contract     *MockERC20SendAndCallReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// MockERC20SendAndCallReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockERC20SendAndCallReceiverRaw struct {
	Contract *MockERC20SendAndCallReceiver // Generic contract binding to access the raw methods on
}

// MockERC20SendAndCallReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockERC20SendAndCallReceiverCallerRaw struct {
	Contract *MockERC20SendAndCallReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// MockERC20SendAndCallReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockERC20SendAndCallReceiverTransactorRaw struct {
	Contract *MockERC20SendAndCallReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockERC20SendAndCallReceiver creates a new instance of MockERC20SendAndCallReceiver, bound to a specific deployed contract.
func NewMockERC20SendAndCallReceiver(address common.Address, backend bind.ContractBackend) (*MockERC20SendAndCallReceiver, error) {
	contract, err := bindMockERC20SendAndCallReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockERC20SendAndCallReceiver{MockERC20SendAndCallReceiverCaller: MockERC20SendAndCallReceiverCaller{contract: contract}, MockERC20SendAndCallReceiverTransactor: MockERC20SendAndCallReceiverTransactor{contract: contract}, MockERC20SendAndCallReceiverFilterer: MockERC20SendAndCallReceiverFilterer{contract: contract}}, nil
}

// NewMockERC20SendAndCallReceiverCaller creates a new read-only instance of MockERC20SendAndCallReceiver, bound to a specific deployed contract.
func NewMockERC20SendAndCallReceiverCaller(address common.Address, caller bind.ContractCaller) (*MockERC20SendAndCallReceiverCaller, error) {
	contract, err := bindMockERC20SendAndCallReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockERC20SendAndCallReceiverCaller{contract: contract}, nil
}

// NewMockERC20SendAndCallReceiverTransactor creates a new write-only instance of MockERC20SendAndCallReceiver, bound to a specific deployed contract.
func NewMockERC20SendAndCallReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MockERC20SendAndCallReceiverTransactor, error) {
	contract, err := bindMockERC20SendAndCallReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockERC20SendAndCallReceiverTransactor{contract: contract}, nil
}

// NewMockERC20SendAndCallReceiverFilterer creates a new log filterer instance of MockERC20SendAndCallReceiver, bound to a specific deployed contract.
func NewMockERC20SendAndCallReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MockERC20SendAndCallReceiverFilterer, error) {
	contract, err := bindMockERC20SendAndCallReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockERC20SendAndCallReceiverFilterer{contract: contract}, nil
}

// bindMockERC20SendAndCallReceiver binds a generic wrapper to an already deployed contract.
func bindMockERC20SendAndCallReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockERC20SendAndCallReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockERC20SendAndCallReceiver.Contract.MockERC20SendAndCallReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.MockERC20SendAndCallReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.MockERC20SendAndCallReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockERC20SendAndCallReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.contract.Transact(opts, method, params...)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x8bc1e076.
//
// Solidity: function receiveTokens(address token, uint256 amount, bytes payload) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactor) ReceiveTokens(opts *bind.TransactOpts, token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.contract.Transact(opts, "receiveTokens", token, amount, payload)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x8bc1e076.
//
// Solidity: function receiveTokens(address token, uint256 amount, bytes payload) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverSession) ReceiveTokens(token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.ReceiveTokens(&_MockERC20SendAndCallReceiver.TransactOpts, token, amount, payload)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x8bc1e076.
//
// Solidity: function receiveTokens(address token, uint256 amount, bytes payload) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactorSession) ReceiveTokens(token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.ReceiveTokens(&_MockERC20SendAndCallReceiver.TransactOpts, token, amount, payload)
}

// MockERC20SendAndCallReceiverTokensReceivedIterator is returned from FilterTokensReceived and is used to iterate over the raw logs and unpacked data for TokensReceived events raised by the MockERC20SendAndCallReceiver contract.
type MockERC20SendAndCallReceiverTokensReceivedIterator struct {
	Event *MockERC20SendAndCallReceiverTokensReceived // Event containing the contract specifics and raw log

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
func (it *MockERC20SendAndCallReceiverTokensReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockERC20SendAndCallReceiverTokensReceived)
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
		it.Event = new(MockERC20SendAndCallReceiverTokensReceived)
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
func (it *MockERC20SendAndCallReceiverTokensReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockERC20SendAndCallReceiverTokensReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockERC20SendAndCallReceiverTokensReceived represents a TokensReceived event raised by the MockERC20SendAndCallReceiver contract.
type MockERC20SendAndCallReceiverTokensReceived struct {
	Token   common.Address
	Amount  *big.Int
	Payload []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokensReceived is a free log retrieval operation binding the contract event 0x0f611859dc94a66d9bec0e66c7485098c0d5e4662896ff2f523c47f962a7abef.
//
// Solidity: event TokensReceived(address token, uint256 amount, bytes payload)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverFilterer) FilterTokensReceived(opts *bind.FilterOpts) (*MockERC20SendAndCallReceiverTokensReceivedIterator, error) {

	logs, sub, err := _MockERC20SendAndCallReceiver.contract.FilterLogs(opts, "TokensReceived")
	if err != nil {
		return nil, err
	}
	return &MockERC20SendAndCallReceiverTokensReceivedIterator{contract: _MockERC20SendAndCallReceiver.contract, event: "TokensReceived", logs: logs, sub: sub}, nil
}

// WatchTokensReceived is a free log subscription operation binding the contract event 0x0f611859dc94a66d9bec0e66c7485098c0d5e4662896ff2f523c47f962a7abef.
//
// Solidity: event TokensReceived(address token, uint256 amount, bytes payload)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverFilterer) WatchTokensReceived(opts *bind.WatchOpts, sink chan<- *MockERC20SendAndCallReceiverTokensReceived) (event.Subscription, error) {

	logs, sub, err := _MockERC20SendAndCallReceiver.contract.WatchLogs(opts, "TokensReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockERC20SendAndCallReceiverTokensReceived)
				if err := _MockERC20SendAndCallReceiver.contract.UnpackLog(event, "TokensReceived", log); err != nil {
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

// ParseTokensReceived is a log parse operation binding the contract event 0x0f611859dc94a66d9bec0e66c7485098c0d5e4662896ff2f523c47f962a7abef.
//
// Solidity: event TokensReceived(address token, uint256 amount, bytes payload)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverFilterer) ParseTokensReceived(log types.Log) (*MockERC20SendAndCallReceiverTokensReceived, error) {
	event := new(MockERC20SendAndCallReceiverTokensReceived)
	if err := _MockERC20SendAndCallReceiver.contract.UnpackLog(event, "TokensReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
