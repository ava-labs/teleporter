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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originBridgeAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"TokensReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"blockSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"blockedSenders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"blocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"receiveTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506108ae806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063487bd69e146100465780637f450d8d1461008557806394395edd146100c3575b600080fd5b610071610054366004610670565b600060208181529281526040808220909352908152205460ff1681565b604051901515815260200160405180910390f35b6100c1610093366004610670565b6000918252602082815260408084206001600160a01b0390931684529190529020805460ff19166001179055565b005b6100c16100d136600461069c565b6000878152602081815260408083206001600160a01b038916845290915290205460ff161561015c5760405162461bcd60e51b815260206004820152602c60248201527f4d6f636b455243323053656e64416e6443616c6c52656365697665723a20736560448201526b1b99195c88189b1bd8dad95960a21b60648201526084015b60405180910390fd5b846001600160a01b0316866001600160a01b0316887f7149eb81ada224b86bfda05a4cf439b25596d3c0d8015aebd1cdc11ab17fe190878787876040516101a69493929190610754565b60405180910390a48061020f5760405162461bcd60e51b815260206004820152602b60248201527f4d6f636b455243323053656e64416e6443616c6c52656365697665723a20656d60448201526a1c1d1e481c185e5b1bd85960aa1b6064820152608401610153565b61021a843385610224565b5050505050505050565b6040516370a0823160e01b815230600482015260009081906001600160a01b038616906370a0823190602401602060405180830381865afa15801561026d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610291919061079c565b90506102a86001600160a01b03861685308661038d565b6040516370a0823160e01b81523060048201526000906001600160a01b038716906370a0823190602401602060405180830381865afa1580156102ef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610313919061079c565b90508181116103795760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610153565b61038382826107b5565b9695505050505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526103e79085906103ed565b50505050565b6000610442826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166104c49092919063ffffffff16565b8051909150156104bf578080602001905181019061046091906107dc565b6104bf5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610153565b505050565b60606104d384846000856104db565b949350505050565b60608247101561053c5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610153565b600080866001600160a01b031685876040516105589190610829565b60006040518083038185875af1925050503d8060008114610595576040519150601f19603f3d011682016040523d82523d6000602084013e61059a565b606091505b50915091506105ab878383876105b6565b979650505050505050565b6060831561062557825160000361061e576001600160a01b0385163b61061e5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610153565b50816104d3565b6104d3838381511561063a5781518083602001fd5b8060405162461bcd60e51b81526004016101539190610845565b80356001600160a01b038116811461066b57600080fd5b919050565b6000806040838503121561068357600080fd5b8235915061069360208401610654565b90509250929050565b600080600080600080600060c0888a0312156106b757600080fd5b873596506106c760208901610654565b95506106d560408901610654565b94506106e360608901610654565b93506080880135925060a088013567ffffffffffffffff8082111561070757600080fd5b818a0191508a601f83011261071b57600080fd5b81358181111561072a57600080fd5b8b602082850101111561073c57600080fd5b60208301945080935050505092959891949750929550565b6001600160a01b0385168152602081018490526060604082018190528101829052818360808301376000818301608090810191909152601f909201601f191601019392505050565b6000602082840312156107ae57600080fd5b5051919050565b818103818111156107d657634e487b7160e01b600052601160045260246000fd5b92915050565b6000602082840312156107ee57600080fd5b815180151581146107fe57600080fd5b9392505050565b60005b83811015610820578181015183820152602001610808565b50506000910152565b6000825161083b818460208701610805565b9190910192915050565b6020815260008251806020840152610864816040850160208701610805565b601f01601f1916919091016040019291505056fea26469706673582212209da9fca30e63475e44fe4ee7dacc68bd9ec453d8e470cda24f5bacf3787a2ffc64736f6c63430008120033",
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

// ReceiveTokens is a paid mutator transaction binding the contract method 0x94395edd.
//
// Solidity: function receiveTokens(bytes32 sourceBlockchainID, address originBridgeAddress, address originSenderAddress, address token, uint256 amount, bytes payload) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactor) ReceiveTokens(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originBridgeAddress common.Address, originSenderAddress common.Address, token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.contract.Transact(opts, "receiveTokens", sourceBlockchainID, originBridgeAddress, originSenderAddress, token, amount, payload)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x94395edd.
//
// Solidity: function receiveTokens(bytes32 sourceBlockchainID, address originBridgeAddress, address originSenderAddress, address token, uint256 amount, bytes payload) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverSession) ReceiveTokens(sourceBlockchainID [32]byte, originBridgeAddress common.Address, originSenderAddress common.Address, token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.ReceiveTokens(&_MockERC20SendAndCallReceiver.TransactOpts, sourceBlockchainID, originBridgeAddress, originSenderAddress, token, amount, payload)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x94395edd.
//
// Solidity: function receiveTokens(bytes32 sourceBlockchainID, address originBridgeAddress, address originSenderAddress, address token, uint256 amount, bytes payload) returns()
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverTransactorSession) ReceiveTokens(sourceBlockchainID [32]byte, originBridgeAddress common.Address, originSenderAddress common.Address, token common.Address, amount *big.Int, payload []byte) (*types.Transaction, error) {
	return _MockERC20SendAndCallReceiver.Contract.ReceiveTokens(&_MockERC20SendAndCallReceiver.TransactOpts, sourceBlockchainID, originBridgeAddress, originSenderAddress, token, amount, payload)
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
	OriginBridgeAddress common.Address
	OriginSenderAddress common.Address
	Token               common.Address
	Amount              *big.Int
	Payload             []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensReceived is a free log retrieval operation binding the contract event 0x7149eb81ada224b86bfda05a4cf439b25596d3c0d8015aebd1cdc11ab17fe190.
//
// Solidity: event TokensReceived(bytes32 indexed sourceBlockchainID, address indexed originBridgeAddress, address indexed originSenderAddress, address token, uint256 amount, bytes payload)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverFilterer) FilterTokensReceived(opts *bind.FilterOpts, sourceBlockchainID [][32]byte, originBridgeAddress []common.Address, originSenderAddress []common.Address) (*MockERC20SendAndCallReceiverTokensReceivedIterator, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originBridgeAddressRule []interface{}
	for _, originBridgeAddressItem := range originBridgeAddress {
		originBridgeAddressRule = append(originBridgeAddressRule, originBridgeAddressItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _MockERC20SendAndCallReceiver.contract.FilterLogs(opts, "TokensReceived", sourceBlockchainIDRule, originBridgeAddressRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return &MockERC20SendAndCallReceiverTokensReceivedIterator{contract: _MockERC20SendAndCallReceiver.contract, event: "TokensReceived", logs: logs, sub: sub}, nil
}

// WatchTokensReceived is a free log subscription operation binding the contract event 0x7149eb81ada224b86bfda05a4cf439b25596d3c0d8015aebd1cdc11ab17fe190.
//
// Solidity: event TokensReceived(bytes32 indexed sourceBlockchainID, address indexed originBridgeAddress, address indexed originSenderAddress, address token, uint256 amount, bytes payload)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverFilterer) WatchTokensReceived(opts *bind.WatchOpts, sink chan<- *MockERC20SendAndCallReceiverTokensReceived, sourceBlockchainID [][32]byte, originBridgeAddress []common.Address, originSenderAddress []common.Address) (event.Subscription, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originBridgeAddressRule []interface{}
	for _, originBridgeAddressItem := range originBridgeAddress {
		originBridgeAddressRule = append(originBridgeAddressRule, originBridgeAddressItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _MockERC20SendAndCallReceiver.contract.WatchLogs(opts, "TokensReceived", sourceBlockchainIDRule, originBridgeAddressRule, originSenderAddressRule)
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

// ParseTokensReceived is a log parse operation binding the contract event 0x7149eb81ada224b86bfda05a4cf439b25596d3c0d8015aebd1cdc11ab17fe190.
//
// Solidity: event TokensReceived(bytes32 indexed sourceBlockchainID, address indexed originBridgeAddress, address indexed originSenderAddress, address token, uint256 amount, bytes payload)
func (_MockERC20SendAndCallReceiver *MockERC20SendAndCallReceiverFilterer) ParseTokensReceived(log types.Log) (*MockERC20SendAndCallReceiverTokensReceived, error) {
	event := new(MockERC20SendAndCallReceiverTokensReceived)
	if err := _MockERC20SendAndCallReceiver.contract.UnpackLog(event, "TokensReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
