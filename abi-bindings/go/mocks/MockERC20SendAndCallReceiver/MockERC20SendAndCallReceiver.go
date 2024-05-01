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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"TokensReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"blockSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"blockedSenders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"blocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"receiveTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061096b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063329d940d14610046578063487bd69e1461005b5780637f450d8d1461009a575b600080fd5b610059610054366004610745565b6100ad565b005b6100866100693660046107eb565b600060208181529281526040808220909352908152205460ff1681565b604051901515815260200160405180910390f35b6100596100a83660046107eb565b6101f4565b6000868152602081815260408083206001600160a01b038916845290915290205460ff16156101385760405162461bcd60e51b815260206004820152602c60248201527f4d6f636b455243323053656e64416e6443616c6c52656365697665723a20736560448201526b1b99195c88189b1bd8dad95960a21b60648201526084015b60405180910390fd5b846001600160a01b0316867f98f64f0ad4e0e2a42535fa15b05dc6e800e16e439c98143fefabb72b43bad53e868686866040516101789493929190610817565b60405180910390a3806101e15760405162461bcd60e51b815260206004820152602b60248201527f4d6f636b455243323053656e64416e6443616c6c52656365697665723a20656d60448201526a1c1d1e481c185e5b1bd85960aa1b606482015260840161012f565b6101eb84846102f8565b50505050505050565b8161025a5760405162461bcd60e51b815260206004820152603060248201527f4d6f636b455243323053656e64416e6443616c6c52656365697665723a207a6560448201526f1c9bc8189b1bd8dad8da185a5b88125160821b606482015260840161012f565b6001600160a01b0381166102ca5760405162461bcd60e51b815260206004820152603160248201527f4d6f636b455243323053656e64416e6443616c6c52656365697665723a207a65604482015270726f2073656e646572206164647265737360781b606482015260840161012f565b6000918252602082815260408084206001600160a01b0390931684529190529020805460ff19166001179055565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015610341573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610365919061085f565b905061037c6001600160a01b038516333086610462565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa1580156103c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e7919061085f565b905081811161044d5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b606482015260840161012f565b6104578282610878565b925050505b92915050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526104bc9085906104c2565b50505050565b6000610517826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166105999092919063ffffffff16565b80519091501561059457808060200190518101906105359190610899565b6105945760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161012f565b505050565b60606105a884846000856105b0565b949350505050565b6060824710156106115760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161012f565b600080866001600160a01b0316858760405161062d91906108e6565b60006040518083038185875af1925050503d806000811461066a576040519150601f19603f3d011682016040523d82523d6000602084013e61066f565b606091505b50915091506106808783838761068b565b979650505050505050565b606083156106fa5782516000036106f3576001600160a01b0385163b6106f35760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161012f565b50816105a8565b6105a8838381511561070f5781518083602001fd5b8060405162461bcd60e51b815260040161012f9190610902565b80356001600160a01b038116811461074057600080fd5b919050565b60008060008060008060a0878903121561075e57600080fd5b8635955061076e60208801610729565b945061077c60408801610729565b935060608701359250608087013567ffffffffffffffff808211156107a057600080fd5b818901915089601f8301126107b457600080fd5b8135818111156107c357600080fd5b8a60208285010111156107d557600080fd5b6020830194508093505050509295509295509295565b600080604083850312156107fe57600080fd5b8235915061080e60208401610729565b90509250929050565b6001600160a01b0385168152602081018490526060604082018190528101829052818360808301376000818301608090810191909152601f909201601f191601019392505050565b60006020828403121561087157600080fd5b5051919050565b8181038181111561045c57634e487b7160e01b600052601160045260246000fd5b6000602082840312156108ab57600080fd5b815180151581146108bb57600080fd5b9392505050565b60005b838110156108dd5781810151838201526020016108c5565b50506000910152565b600082516108f88184602087016108c2565b9190910192915050565b60208152600082518060208401526109218160408501602087016108c2565b601f01601f1916919091016040019291505056fea2646970667358221220580a7972aac3081a2d1cb640e85a6c81bbaedecb2306736119a10d4d14f50bd364736f6c63430008120033",
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

// BlockedSenders is a free data retrieval call binding the contract method 0x487bd69e.
//
// Solidity: function blockedSenders(bytes32 blockchainID, address senderAddress) view returns(bool blocked)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverCaller) BlockedSenders(opts *bind.CallOpts, blockchainID [32]byte, senderAddress common.Address) (bool, error) {
	var out []interface{}
	err := _MockERC20SendAndCallReceiver.contract.Call(opts, &out, "blockedSenders", blockchainID, senderAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlockedSenders is a free data retrieval call binding the contract method 0x487bd69e.
//
// Solidity: function blockedSenders(bytes32 blockchainID, address senderAddress) view returns(bool blocked)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverSession) BlockedSenders(blockchainID [32]byte, senderAddress common.Address) (bool, error) {
	return _MockERC20SendAndCallReceiver.Contract.BlockedSenders(&_MockERC20SendAndCallReceiver.CallOpts, blockchainID, senderAddress)
}

// BlockedSenders is a free data retrieval call binding the contract method 0x487bd69e.
//
// Solidity: function blockedSenders(bytes32 blockchainID, address senderAddress) view returns(bool blocked)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverCallerSession) BlockedSenders(blockchainID [32]byte, senderAddress common.Address) (bool, error) {
	return _MockERC20SendAndCallReceiver.Contract.BlockedSenders(&_MockERC20SendAndCallReceiver.CallOpts, blockchainID, senderAddress)
}

// BlockSender is a paid mutator transaction binding the contract method 0x7f450d8d.
//
// Solidity: function blockSender(bytes32 blockchainID, address senderAddress) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactor) BlockSender(opts *bind.TransactOpts, blockchainID [32]byte, senderAddress common.Address) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.contract.Transact(opts, "blockSender", blockchainID, senderAddress)
}

// BlockSender is a paid mutator transaction binding the contract method 0x7f450d8d.
//
// Solidity: function blockSender(bytes32 blockchainID, address senderAddress) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverSession) BlockSender(blockchainID [32]byte, senderAddress common.Address) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.BlockSender(&_MockERC20SendAndCallReceiver.TransactOpts, blockchainID, senderAddress)
}

// BlockSender is a paid mutator transaction binding the contract method 0x7f450d8d.
//
// Solidity: function blockSender(bytes32 blockchainID, address senderAddress) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactorSession) BlockSender(blockchainID [32]byte, senderAddress common.Address) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.BlockSender(&_MockERC20SendAndCallReceiver.TransactOpts, blockchainID, senderAddress)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x329d940d.
//
// Solidity: function receiveTokens(bytes32 sourceBlockchainID, address originSenderAddress, address token, uint256 amount, bytes payload) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactor) ReceiveTokens(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.contract.Transact(opts, "receiveTokens", sourceBlockchainID, originSenderAddress, token, amount, payload)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x329d940d.
//
// Solidity: function receiveTokens(bytes32 sourceBlockchainID, address originSenderAddress, address token, uint256 amount, bytes payload) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverSession) ReceiveTokens(sourceBlockchainID [32]byte, originSenderAddress common.Address, token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.ReceiveTokens(&_MockERC20SendAndCallReceiver.TransactOpts, sourceBlockchainID, originSenderAddress, token, amount, payload)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x329d940d.
//
// Solidity: function receiveTokens(bytes32 sourceBlockchainID, address originSenderAddress, address token, uint256 amount, bytes payload) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactorSession) ReceiveTokens(sourceBlockchainID [32]byte, originSenderAddress common.Address, token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.ReceiveTokens(&_MockERC20SendAndCallReceiver.TransactOpts, sourceBlockchainID, originSenderAddress, token, amount, payload)
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
	SourceBlockchainID  [32]byte
	OriginSenderAddress common.Address
	Token               common.Address
	Amount              *big.Int
	Payload             []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensReceived is a free log retrieval operation binding the contract event 0x98f64f0ad4e0e2a42535fa15b05dc6e800e16e439c98143fefabb72b43bad53e.
//
// Solidity: event TokensReceived(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, address token, uint256 amount, bytes payload)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverFilterer) FilterTokensReceived(opts *bind.FilterOpts, sourceBlockchainID [][32]byte, originSenderAddress []common.Address) (*MockERC20SendAndCallReceiverTokensReceivedIterator, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _MockERC20SendAndCallReceiver.contract.FilterLogs(opts, "TokensReceived", sourceBlockchainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return &MockERC20SendAndCallReceiverTokensReceivedIterator{contract: _MockERC20SendAndCallReceiver.contract, event: "TokensReceived", logs: logs, sub: sub}, nil
}

// WatchTokensReceived is a free log subscription operation binding the contract event 0x98f64f0ad4e0e2a42535fa15b05dc6e800e16e439c98143fefabb72b43bad53e.
//
// Solidity: event TokensReceived(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, address token, uint256 amount, bytes payload)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverFilterer) WatchTokensReceived(opts *bind.WatchOpts, sink chan<- *MockERC20SendAndCallReceiverTokensReceived, sourceBlockchainID [][32]byte, originSenderAddress []common.Address) (event.Subscription, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _MockERC20SendAndCallReceiver.contract.WatchLogs(opts, "TokensReceived", sourceBlockchainIDRule, originSenderAddressRule)
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

// ParseTokensReceived is a log parse operation binding the contract event 0x98f64f0ad4e0e2a42535fa15b05dc6e800e16e439c98143fefabb72b43bad53e.
//
// Solidity: event TokensReceived(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, address token, uint256 amount, bytes payload)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverFilterer) ParseTokensReceived(log types.Log) (*MockERC20SendAndCallReceiverTokensReceived, error) {
	event := new(MockERC20SendAndCallReceiverTokensReceived)
	if err := _MockERC20SendAndCallReceiver.contract.UnpackLog(event, "TokensReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
