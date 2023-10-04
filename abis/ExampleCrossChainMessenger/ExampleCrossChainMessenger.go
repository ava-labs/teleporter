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

// ExamplecrosschainmessengerMetaData contains all meta data concerning the Examplecrosschainmessenger contract.
var ExamplecrosschainmessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceNotIncreased\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ReceiveMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"SendMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"}],\"name\":\"getCurrentMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ExamplecrosschainmessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use ExamplecrosschainmessengerMetaData.ABI instead.
var ExamplecrosschainmessengerABI = ExamplecrosschainmessengerMetaData.ABI

// Examplecrosschainmessenger is an auto generated Go binding around an Ethereum contract.
type Examplecrosschainmessenger struct {
	ExamplecrosschainmessengerCaller     // Read-only binding to the contract
	ExamplecrosschainmessengerTransactor // Write-only binding to the contract
	ExamplecrosschainmessengerFilterer   // Log filterer for contract events
}

// ExamplecrosschainmessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExamplecrosschainmessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExamplecrosschainmessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExamplecrosschainmessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExamplecrosschainmessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExamplecrosschainmessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExamplecrosschainmessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExamplecrosschainmessengerSession struct {
	Contract     *Examplecrosschainmessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExamplecrosschainmessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExamplecrosschainmessengerCallerSession struct {
	Contract *ExamplecrosschainmessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ExamplecrosschainmessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExamplecrosschainmessengerTransactorSession struct {
	Contract     *ExamplecrosschainmessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ExamplecrosschainmessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExamplecrosschainmessengerRaw struct {
	Contract *Examplecrosschainmessenger // Generic contract binding to access the raw methods on
}

// ExamplecrosschainmessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExamplecrosschainmessengerCallerRaw struct {
	Contract *ExamplecrosschainmessengerCaller // Generic read-only contract binding to access the raw methods on
}

// ExamplecrosschainmessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExamplecrosschainmessengerTransactorRaw struct {
	Contract *ExamplecrosschainmessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExamplecrosschainmessenger creates a new instance of Examplecrosschainmessenger, bound to a specific deployed contract.
func NewExamplecrosschainmessenger(address common.Address, backend bind.ContractBackend) (*Examplecrosschainmessenger, error) {
	contract, err := bindExamplecrosschainmessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Examplecrosschainmessenger{ExamplecrosschainmessengerCaller: ExamplecrosschainmessengerCaller{contract: contract}, ExamplecrosschainmessengerTransactor: ExamplecrosschainmessengerTransactor{contract: contract}, ExamplecrosschainmessengerFilterer: ExamplecrosschainmessengerFilterer{contract: contract}}, nil
}

// NewExamplecrosschainmessengerCaller creates a new read-only instance of Examplecrosschainmessenger, bound to a specific deployed contract.
func NewExamplecrosschainmessengerCaller(address common.Address, caller bind.ContractCaller) (*ExamplecrosschainmessengerCaller, error) {
	contract, err := bindExamplecrosschainmessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExamplecrosschainmessengerCaller{contract: contract}, nil
}

// NewExamplecrosschainmessengerTransactor creates a new write-only instance of Examplecrosschainmessenger, bound to a specific deployed contract.
func NewExamplecrosschainmessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*ExamplecrosschainmessengerTransactor, error) {
	contract, err := bindExamplecrosschainmessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExamplecrosschainmessengerTransactor{contract: contract}, nil
}

// NewExamplecrosschainmessengerFilterer creates a new log filterer instance of Examplecrosschainmessenger, bound to a specific deployed contract.
func NewExamplecrosschainmessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*ExamplecrosschainmessengerFilterer, error) {
	contract, err := bindExamplecrosschainmessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExamplecrosschainmessengerFilterer{contract: contract}, nil
}

// bindExamplecrosschainmessenger binds a generic wrapper to an already deployed contract.
func bindExamplecrosschainmessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExamplecrosschainmessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Examplecrosschainmessenger.Contract.ExamplecrosschainmessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Examplecrosschainmessenger.Contract.ExamplecrosschainmessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Examplecrosschainmessenger.Contract.ExamplecrosschainmessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Examplecrosschainmessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Examplecrosschainmessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Examplecrosschainmessenger.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 originChainID) view returns(address sender, string message)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerCaller) GetCurrentMessage(opts *bind.CallOpts, originChainID [32]byte) (struct {
	Sender  common.Address
	Message string
}, error) {
	var out []interface{}
	err := _Examplecrosschainmessenger.contract.Call(opts, &out, "getCurrentMessage", originChainID)

	outstruct := new(struct {
		Sender  common.Address
		Message string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Message = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 originChainID) view returns(address sender, string message)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerSession) GetCurrentMessage(originChainID [32]byte) (struct {
	Sender  common.Address
	Message string
}, error) {
	return _Examplecrosschainmessenger.Contract.GetCurrentMessage(&_Examplecrosschainmessenger.CallOpts, originChainID)
}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 originChainID) view returns(address sender, string message)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerCallerSession) GetCurrentMessage(originChainID [32]byte) (struct {
	Sender  common.Address
	Message string
}, error) {
	return _Examplecrosschainmessenger.Contract.GetCurrentMessage(&_Examplecrosschainmessenger.CallOpts, originChainID)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Examplecrosschainmessenger.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerSession) TeleporterMessenger() (common.Address, error) {
	return _Examplecrosschainmessenger.Contract.TeleporterMessenger(&_Examplecrosschainmessenger.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerCallerSession) TeleporterMessenger() (common.Address, error) {
	return _Examplecrosschainmessenger.Contract.TeleporterMessenger(&_Examplecrosschainmessenger.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Examplecrosschainmessenger.contract.Transact(opts, "receiveTeleporterMessage", originChainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerSession) ReceiveTeleporterMessage(originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Examplecrosschainmessenger.Contract.ReceiveTeleporterMessage(&_Examplecrosschainmessenger.TransactOpts, originChainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerTransactorSession) ReceiveTeleporterMessage(originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Examplecrosschainmessenger.Contract.ReceiveTeleporterMessage(&_Examplecrosschainmessenger.TransactOpts, originChainID, originSenderAddress, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationChainID, address destinationAddress, address feeContractAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(uint256 messageID)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerTransactor) SendMessage(opts *bind.TransactOpts, destinationChainID [32]byte, destinationAddress common.Address, feeContractAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _Examplecrosschainmessenger.contract.Transact(opts, "sendMessage", destinationChainID, destinationAddress, feeContractAddress, feeAmount, requiredGasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationChainID, address destinationAddress, address feeContractAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(uint256 messageID)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerSession) SendMessage(destinationChainID [32]byte, destinationAddress common.Address, feeContractAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _Examplecrosschainmessenger.Contract.SendMessage(&_Examplecrosschainmessenger.TransactOpts, destinationChainID, destinationAddress, feeContractAddress, feeAmount, requiredGasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationChainID, address destinationAddress, address feeContractAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(uint256 messageID)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerTransactorSession) SendMessage(destinationChainID [32]byte, destinationAddress common.Address, feeContractAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _Examplecrosschainmessenger.Contract.SendMessage(&_Examplecrosschainmessenger.TransactOpts, destinationChainID, destinationAddress, feeContractAddress, feeAmount, requiredGasLimit, message)
}

// ExamplecrosschainmessengerReceiveMessageIterator is returned from FilterReceiveMessage and is used to iterate over the raw logs and unpacked data for ReceiveMessage events raised by the Examplecrosschainmessenger contract.
type ExamplecrosschainmessengerReceiveMessageIterator struct {
	Event *ExamplecrosschainmessengerReceiveMessage // Event containing the contract specifics and raw log

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
func (it *ExamplecrosschainmessengerReceiveMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExamplecrosschainmessengerReceiveMessage)
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
		it.Event = new(ExamplecrosschainmessengerReceiveMessage)
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
func (it *ExamplecrosschainmessengerReceiveMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExamplecrosschainmessengerReceiveMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExamplecrosschainmessengerReceiveMessage represents a ReceiveMessage event raised by the Examplecrosschainmessenger contract.
type ExamplecrosschainmessengerReceiveMessage struct {
	OriginChainID       [32]byte
	OriginSenderAddress common.Address
	Message             string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReceiveMessage is a free log retrieval operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed originChainID, address indexed originSenderAddress, string message)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerFilterer) FilterReceiveMessage(opts *bind.FilterOpts, originChainID [][32]byte, originSenderAddress []common.Address) (*ExamplecrosschainmessengerReceiveMessageIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _Examplecrosschainmessenger.contract.FilterLogs(opts, "ReceiveMessage", originChainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return &ExamplecrosschainmessengerReceiveMessageIterator{contract: _Examplecrosschainmessenger.contract, event: "ReceiveMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveMessage is a free log subscription operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed originChainID, address indexed originSenderAddress, string message)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerFilterer) WatchReceiveMessage(opts *bind.WatchOpts, sink chan<- *ExamplecrosschainmessengerReceiveMessage, originChainID [][32]byte, originSenderAddress []common.Address) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _Examplecrosschainmessenger.contract.WatchLogs(opts, "ReceiveMessage", originChainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExamplecrosschainmessengerReceiveMessage)
				if err := _Examplecrosschainmessenger.contract.UnpackLog(event, "ReceiveMessage", log); err != nil {
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
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerFilterer) ParseReceiveMessage(log types.Log) (*ExamplecrosschainmessengerReceiveMessage, error) {
	event := new(ExamplecrosschainmessengerReceiveMessage)
	if err := _Examplecrosschainmessenger.contract.UnpackLog(event, "ReceiveMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExamplecrosschainmessengerSendMessageIterator is returned from FilterSendMessage and is used to iterate over the raw logs and unpacked data for SendMessage events raised by the Examplecrosschainmessenger contract.
type ExamplecrosschainmessengerSendMessageIterator struct {
	Event *ExamplecrosschainmessengerSendMessage // Event containing the contract specifics and raw log

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
func (it *ExamplecrosschainmessengerSendMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExamplecrosschainmessengerSendMessage)
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
		it.Event = new(ExamplecrosschainmessengerSendMessage)
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
func (it *ExamplecrosschainmessengerSendMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExamplecrosschainmessengerSendMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExamplecrosschainmessengerSendMessage represents a SendMessage event raised by the Examplecrosschainmessenger contract.
type ExamplecrosschainmessengerSendMessage struct {
	DestinationChainID [32]byte
	DestinationAddress common.Address
	FeeAsset           common.Address
	FeeAmount          *big.Int
	RequiredGasLimit   *big.Int
	Message            string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSendMessage is a free log retrieval operation binding the contract event 0xa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca.
//
// Solidity: event SendMessage(bytes32 indexed destinationChainID, address indexed destinationAddress, address feeAsset, uint256 feeAmount, uint256 requiredGasLimit, string message)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerFilterer) FilterSendMessage(opts *bind.FilterOpts, destinationChainID [][32]byte, destinationAddress []common.Address) (*ExamplecrosschainmessengerSendMessageIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _Examplecrosschainmessenger.contract.FilterLogs(opts, "SendMessage", destinationChainIDRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return &ExamplecrosschainmessengerSendMessageIterator{contract: _Examplecrosschainmessenger.contract, event: "SendMessage", logs: logs, sub: sub}, nil
}

// WatchSendMessage is a free log subscription operation binding the contract event 0xa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca.
//
// Solidity: event SendMessage(bytes32 indexed destinationChainID, address indexed destinationAddress, address feeAsset, uint256 feeAmount, uint256 requiredGasLimit, string message)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerFilterer) WatchSendMessage(opts *bind.WatchOpts, sink chan<- *ExamplecrosschainmessengerSendMessage, destinationChainID [][32]byte, destinationAddress []common.Address) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _Examplecrosschainmessenger.contract.WatchLogs(opts, "SendMessage", destinationChainIDRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExamplecrosschainmessengerSendMessage)
				if err := _Examplecrosschainmessenger.contract.UnpackLog(event, "SendMessage", log); err != nil {
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
// Solidity: event SendMessage(bytes32 indexed destinationChainID, address indexed destinationAddress, address feeAsset, uint256 feeAmount, uint256 requiredGasLimit, string message)
func (_Examplecrosschainmessenger *ExamplecrosschainmessengerFilterer) ParseSendMessage(log types.Log) (*ExamplecrosschainmessengerSendMessage, error) {
	event := new(ExamplecrosschainmessengerSendMessage)
	if err := _Examplecrosschainmessenger.contract.UnpackLog(event, "SendMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
