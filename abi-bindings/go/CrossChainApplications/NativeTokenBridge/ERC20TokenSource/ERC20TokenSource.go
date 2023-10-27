// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20tokensource

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

// ERC20TokenSourceMetaData contains all meta data concerning the ERC20TokenSource contract.
var ERC20TokenSourceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeTokenDestinationAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ERC20ContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"}],\"name\":\"TransferToDestination\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINT_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenDestinationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"senderBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20ContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"transferToDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20TokenSourceABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenSourceMetaData.ABI instead.
var ERC20TokenSourceABI = ERC20TokenSourceMetaData.ABI

// ERC20TokenSource is an auto generated Go binding around an Ethereum contract.
type ERC20TokenSource struct {
	ERC20TokenSourceCaller     // Read-only binding to the contract
	ERC20TokenSourceTransactor // Write-only binding to the contract
	ERC20TokenSourceFilterer   // Log filterer for contract events
}

// ERC20TokenSourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenSourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenSourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenSourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenSourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenSourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenSourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenSourceSession struct {
	Contract     *ERC20TokenSource // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20TokenSourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenSourceCallerSession struct {
	Contract *ERC20TokenSourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC20TokenSourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenSourceTransactorSession struct {
	Contract     *ERC20TokenSourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC20TokenSourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenSourceRaw struct {
	Contract *ERC20TokenSource // Generic contract binding to access the raw methods on
}

// ERC20TokenSourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenSourceCallerRaw struct {
	Contract *ERC20TokenSourceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenSourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenSourceTransactorRaw struct {
	Contract *ERC20TokenSourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenSource creates a new instance of ERC20TokenSource, bound to a specific deployed contract.
func NewERC20TokenSource(address common.Address, backend bind.ContractBackend) (*ERC20TokenSource, error) {
	contract, err := bindERC20TokenSource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSource{ERC20TokenSourceCaller: ERC20TokenSourceCaller{contract: contract}, ERC20TokenSourceTransactor: ERC20TokenSourceTransactor{contract: contract}, ERC20TokenSourceFilterer: ERC20TokenSourceFilterer{contract: contract}}, nil
}

// NewERC20TokenSourceCaller creates a new read-only instance of ERC20TokenSource, bound to a specific deployed contract.
func NewERC20TokenSourceCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenSourceCaller, error) {
	contract, err := bindERC20TokenSource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceCaller{contract: contract}, nil
}

// NewERC20TokenSourceTransactor creates a new write-only instance of ERC20TokenSource, bound to a specific deployed contract.
func NewERC20TokenSourceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenSourceTransactor, error) {
	contract, err := bindERC20TokenSource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceTransactor{contract: contract}, nil
}

// NewERC20TokenSourceFilterer creates a new log filterer instance of ERC20TokenSource, bound to a specific deployed contract.
func NewERC20TokenSourceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenSourceFilterer, error) {
	contract, err := bindERC20TokenSource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceFilterer{contract: contract}, nil
}

// bindERC20TokenSource binds a generic wrapper to an already deployed contract.
func bindERC20TokenSource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TokenSourceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenSource *ERC20TokenSourceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenSource.Contract.ERC20TokenSourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenSource *ERC20TokenSourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.ERC20TokenSourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenSource *ERC20TokenSourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.ERC20TokenSourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenSource *ERC20TokenSourceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenSource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenSource *ERC20TokenSourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenSource *ERC20TokenSourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.contract.Transact(opts, method, params...)
}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceCaller) MINTNATIVETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "MINT_NATIVE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceSession) MINTNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenSource.Contract.MINTNATIVETOKENSREQUIREDGAS(&_ERC20TokenSource.CallOpts)
}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) MINTNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenSource.Contract.MINTNATIVETOKENSREQUIREDGAS(&_ERC20TokenSource.CallOpts)
}

// CurrentBlockchainID is a free data retrieval call binding the contract method 0x4950d2d0.
//
// Solidity: function currentBlockchainID() view returns(bytes32)
func (_ERC20TokenSource *ERC20TokenSourceCaller) CurrentBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "currentBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentBlockchainID is a free data retrieval call binding the contract method 0x4950d2d0.
//
// Solidity: function currentBlockchainID() view returns(bytes32)
func (_ERC20TokenSource *ERC20TokenSourceSession) CurrentBlockchainID() ([32]byte, error) {
	return _ERC20TokenSource.Contract.CurrentBlockchainID(&_ERC20TokenSource.CallOpts)
}

// CurrentBlockchainID is a free data retrieval call binding the contract method 0x4950d2d0.
//
// Solidity: function currentBlockchainID() view returns(bytes32)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) CurrentBlockchainID() ([32]byte, error) {
	return _ERC20TokenSource.Contract.CurrentBlockchainID(&_ERC20TokenSource.CallOpts)
}

// DestinationBlockchainID is a free data retrieval call binding the contract method 0x41d3014d.
//
// Solidity: function destinationBlockchainID() view returns(bytes32)
func (_ERC20TokenSource *ERC20TokenSourceCaller) DestinationBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "destinationBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DestinationBlockchainID is a free data retrieval call binding the contract method 0x41d3014d.
//
// Solidity: function destinationBlockchainID() view returns(bytes32)
func (_ERC20TokenSource *ERC20TokenSourceSession) DestinationBlockchainID() ([32]byte, error) {
	return _ERC20TokenSource.Contract.DestinationBlockchainID(&_ERC20TokenSource.CallOpts)
}

// DestinationBlockchainID is a free data retrieval call binding the contract method 0x41d3014d.
//
// Solidity: function destinationBlockchainID() view returns(bytes32)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) DestinationBlockchainID() ([32]byte, error) {
	return _ERC20TokenSource.Contract.DestinationBlockchainID(&_ERC20TokenSource.CallOpts)
}

// NativeTokenDestinationAddress is a free data retrieval call binding the contract method 0xb8c9091a.
//
// Solidity: function nativeTokenDestinationAddress() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCaller) NativeTokenDestinationAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "nativeTokenDestinationAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeTokenDestinationAddress is a free data retrieval call binding the contract method 0xb8c9091a.
//
// Solidity: function nativeTokenDestinationAddress() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceSession) NativeTokenDestinationAddress() (common.Address, error) {
	return _ERC20TokenSource.Contract.NativeTokenDestinationAddress(&_ERC20TokenSource.CallOpts)
}

// NativeTokenDestinationAddress is a free data retrieval call binding the contract method 0xb8c9091a.
//
// Solidity: function nativeTokenDestinationAddress() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) NativeTokenDestinationAddress() (common.Address, error) {
	return _ERC20TokenSource.Contract.NativeTokenDestinationAddress(&_ERC20TokenSource.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceSession) TeleporterMessenger() (common.Address, error) {
	return _ERC20TokenSource.Contract.TeleporterMessenger(&_ERC20TokenSource.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) TeleporterMessenger() (common.Address, error) {
	return _ERC20TokenSource.Contract.TeleporterMessenger(&_ERC20TokenSource.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenSource.contract.Transact(opts, "receiveTeleporterMessage", senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_ERC20TokenSource *ERC20TokenSourceSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.ReceiveTeleporterMessage(&_ERC20TokenSource.TransactOpts, senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactorSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.ReceiveTeleporterMessage(&_ERC20TokenSource.TransactOpts, senderBlockchainID, senderAddress, message)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0x10a61f2b.
//
// Solidity: function transferToDestination(address recipient, address ERC20ContractAddress, uint256 totalAmount, uint256 feeAmount, address[] allowedRelayerAddresses) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactor) TransferToDestination(opts *bind.TransactOpts, recipient common.Address, ERC20ContractAddress common.Address, totalAmount *big.Int, feeAmount *big.Int, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.contract.Transact(opts, "transferToDestination", recipient, ERC20ContractAddress, totalAmount, feeAmount, allowedRelayerAddresses)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0x10a61f2b.
//
// Solidity: function transferToDestination(address recipient, address ERC20ContractAddress, uint256 totalAmount, uint256 feeAmount, address[] allowedRelayerAddresses) returns()
func (_ERC20TokenSource *ERC20TokenSourceSession) TransferToDestination(recipient common.Address, ERC20ContractAddress common.Address, totalAmount *big.Int, feeAmount *big.Int, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.TransferToDestination(&_ERC20TokenSource.TransactOpts, recipient, ERC20ContractAddress, totalAmount, feeAmount, allowedRelayerAddresses)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0x10a61f2b.
//
// Solidity: function transferToDestination(address recipient, address ERC20ContractAddress, uint256 totalAmount, uint256 feeAmount, address[] allowedRelayerAddresses) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactorSession) TransferToDestination(recipient common.Address, ERC20ContractAddress common.Address, totalAmount *big.Int, feeAmount *big.Int, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.TransferToDestination(&_ERC20TokenSource.TransactOpts, recipient, ERC20ContractAddress, totalAmount, feeAmount, allowedRelayerAddresses)
}

// ERC20TokenSourceTransferToDestinationIterator is returned from FilterTransferToDestination and is used to iterate over the raw logs and unpacked data for TransferToDestination events raised by the ERC20TokenSource contract.
type ERC20TokenSourceTransferToDestinationIterator struct {
	Event *ERC20TokenSourceTransferToDestination // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSourceTransferToDestinationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSourceTransferToDestination)
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
		it.Event = new(ERC20TokenSourceTransferToDestination)
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
func (it *ERC20TokenSourceTransferToDestinationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSourceTransferToDestinationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSourceTransferToDestination represents a TransferToDestination event raised by the ERC20TokenSource contract.
type ERC20TokenSourceTransferToDestination struct {
	Sender               common.Address
	Recipient            common.Address
	ERC20ContractAddress common.Address
	TransferAmount       *big.Int
	FeeAmount            *big.Int
	TeleporterMessageID  *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTransferToDestination is a free log retrieval operation binding the contract event 0x13a202fe5d8659b297514aa4630f23f3a30583706bf3aa067257552aa20e9ec3.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, address ERC20ContractAddress, uint256 transferAmount, uint256 feeAmount, uint256 teleporterMessageID)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) FilterTransferToDestination(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*ERC20TokenSourceTransferToDestinationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.FilterLogs(opts, "TransferToDestination", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceTransferToDestinationIterator{contract: _ERC20TokenSource.contract, event: "TransferToDestination", logs: logs, sub: sub}, nil
}

// WatchTransferToDestination is a free log subscription operation binding the contract event 0x13a202fe5d8659b297514aa4630f23f3a30583706bf3aa067257552aa20e9ec3.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, address ERC20ContractAddress, uint256 transferAmount, uint256 feeAmount, uint256 teleporterMessageID)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) WatchTransferToDestination(opts *bind.WatchOpts, sink chan<- *ERC20TokenSourceTransferToDestination, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.WatchLogs(opts, "TransferToDestination", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSourceTransferToDestination)
				if err := _ERC20TokenSource.contract.UnpackLog(event, "TransferToDestination", log); err != nil {
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

// ParseTransferToDestination is a log parse operation binding the contract event 0x13a202fe5d8659b297514aa4630f23f3a30583706bf3aa067257552aa20e9ec3.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, address ERC20ContractAddress, uint256 transferAmount, uint256 feeAmount, uint256 teleporterMessageID)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) ParseTransferToDestination(log types.Log) (*ERC20TokenSourceTransferToDestination, error) {
	event := new(ERC20TokenSourceTransferToDestination)
	if err := _ERC20TokenSource.contract.UnpackLog(event, "TransferToDestination", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenSourceUnlockTokensIterator is returned from FilterUnlockTokens and is used to iterate over the raw logs and unpacked data for UnlockTokens events raised by the ERC20TokenSource contract.
type ERC20TokenSourceUnlockTokensIterator struct {
	Event *ERC20TokenSourceUnlockTokens // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSourceUnlockTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSourceUnlockTokens)
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
		it.Event = new(ERC20TokenSourceUnlockTokens)
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
func (it *ERC20TokenSourceUnlockTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSourceUnlockTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSourceUnlockTokens represents a UnlockTokens event raised by the ERC20TokenSource contract.
type ERC20TokenSourceUnlockTokens struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlockTokens is a free log retrieval operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) FilterUnlockTokens(opts *bind.FilterOpts) (*ERC20TokenSourceUnlockTokensIterator, error) {

	logs, sub, err := _ERC20TokenSource.contract.FilterLogs(opts, "UnlockTokens")
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceUnlockTokensIterator{contract: _ERC20TokenSource.contract, event: "UnlockTokens", logs: logs, sub: sub}, nil
}

// WatchUnlockTokens is a free log subscription operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) WatchUnlockTokens(opts *bind.WatchOpts, sink chan<- *ERC20TokenSourceUnlockTokens) (event.Subscription, error) {

	logs, sub, err := _ERC20TokenSource.contract.WatchLogs(opts, "UnlockTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSourceUnlockTokens)
				if err := _ERC20TokenSource.contract.UnpackLog(event, "UnlockTokens", log); err != nil {
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

// ParseUnlockTokens is a log parse operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) ParseUnlockTokens(log types.Log) (*ERC20TokenSourceUnlockTokens, error) {
	event := new(ERC20TokenSourceUnlockTokens)
	if err := _ERC20TokenSource.contract.UnpackLog(event, "UnlockTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
