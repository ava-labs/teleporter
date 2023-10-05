// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokensource

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

// NativetokensourceMetaData contains all meta data concerning the Nativetokensource contract.
var NativetokensourceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationContractAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceNotIncreased\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"}],\"name\":\"TransferToDestination\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINT_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"senderBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"transferToDestination\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NativetokensourceABI is the input ABI used to generate the binding from.
// Deprecated: Use NativetokensourceMetaData.ABI instead.
var NativetokensourceABI = NativetokensourceMetaData.ABI

// Nativetokensource is an auto generated Go binding around an Ethereum contract.
type Nativetokensource struct {
	NativetokensourceCaller     // Read-only binding to the contract
	NativetokensourceTransactor // Write-only binding to the contract
	NativetokensourceFilterer   // Log filterer for contract events
}

// NativetokensourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativetokensourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativetokensourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativetokensourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativetokensourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativetokensourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativetokensourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativetokensourceSession struct {
	Contract     *Nativetokensource // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NativetokensourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativetokensourceCallerSession struct {
	Contract *NativetokensourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NativetokensourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativetokensourceTransactorSession struct {
	Contract     *NativetokensourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NativetokensourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativetokensourceRaw struct {
	Contract *Nativetokensource // Generic contract binding to access the raw methods on
}

// NativetokensourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativetokensourceCallerRaw struct {
	Contract *NativetokensourceCaller // Generic read-only contract binding to access the raw methods on
}

// NativetokensourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativetokensourceTransactorRaw struct {
	Contract *NativetokensourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativetokensource creates a new instance of Nativetokensource, bound to a specific deployed contract.
func NewNativetokensource(address common.Address, backend bind.ContractBackend) (*Nativetokensource, error) {
	contract, err := bindNativetokensource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nativetokensource{NativetokensourceCaller: NativetokensourceCaller{contract: contract}, NativetokensourceTransactor: NativetokensourceTransactor{contract: contract}, NativetokensourceFilterer: NativetokensourceFilterer{contract: contract}}, nil
}

// NewNativetokensourceCaller creates a new read-only instance of Nativetokensource, bound to a specific deployed contract.
func NewNativetokensourceCaller(address common.Address, caller bind.ContractCaller) (*NativetokensourceCaller, error) {
	contract, err := bindNativetokensource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativetokensourceCaller{contract: contract}, nil
}

// NewNativetokensourceTransactor creates a new write-only instance of Nativetokensource, bound to a specific deployed contract.
func NewNativetokensourceTransactor(address common.Address, transactor bind.ContractTransactor) (*NativetokensourceTransactor, error) {
	contract, err := bindNativetokensource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativetokensourceTransactor{contract: contract}, nil
}

// NewNativetokensourceFilterer creates a new log filterer instance of Nativetokensource, bound to a specific deployed contract.
func NewNativetokensourceFilterer(address common.Address, filterer bind.ContractFilterer) (*NativetokensourceFilterer, error) {
	contract, err := bindNativetokensource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativetokensourceFilterer{contract: contract}, nil
}

// bindNativetokensource binds a generic wrapper to an already deployed contract.
func bindNativetokensource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativetokensourceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nativetokensource *NativetokensourceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nativetokensource.Contract.NativetokensourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nativetokensource *NativetokensourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nativetokensource.Contract.NativetokensourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nativetokensource *NativetokensourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nativetokensource.Contract.NativetokensourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nativetokensource *NativetokensourceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nativetokensource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nativetokensource *NativetokensourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nativetokensource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nativetokensource *NativetokensourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nativetokensource.Contract.contract.Transact(opts, method, params...)
}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Nativetokensource *NativetokensourceCaller) MINTNATIVETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nativetokensource.contract.Call(opts, &out, "MINT_NATIVE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Nativetokensource *NativetokensourceSession) MINTNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _Nativetokensource.Contract.MINTNATIVETOKENSREQUIREDGAS(&_Nativetokensource.CallOpts)
}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Nativetokensource *NativetokensourceCallerSession) MINTNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _Nativetokensource.Contract.MINTNATIVETOKENSREQUIREDGAS(&_Nativetokensource.CallOpts)
}

// CurrentBlockchainID is a free data retrieval call binding the contract method 0x4950d2d0.
//
// Solidity: function currentBlockchainID() view returns(bytes32)
func (_Nativetokensource *NativetokensourceCaller) CurrentBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nativetokensource.contract.Call(opts, &out, "currentBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentBlockchainID is a free data retrieval call binding the contract method 0x4950d2d0.
//
// Solidity: function currentBlockchainID() view returns(bytes32)
func (_Nativetokensource *NativetokensourceSession) CurrentBlockchainID() ([32]byte, error) {
	return _Nativetokensource.Contract.CurrentBlockchainID(&_Nativetokensource.CallOpts)
}

// CurrentBlockchainID is a free data retrieval call binding the contract method 0x4950d2d0.
//
// Solidity: function currentBlockchainID() view returns(bytes32)
func (_Nativetokensource *NativetokensourceCallerSession) CurrentBlockchainID() ([32]byte, error) {
	return _Nativetokensource.Contract.CurrentBlockchainID(&_Nativetokensource.CallOpts)
}

// DestinationBlockchainID is a free data retrieval call binding the contract method 0x41d3014d.
//
// Solidity: function destinationBlockchainID() view returns(bytes32)
func (_Nativetokensource *NativetokensourceCaller) DestinationBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nativetokensource.contract.Call(opts, &out, "destinationBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DestinationBlockchainID is a free data retrieval call binding the contract method 0x41d3014d.
//
// Solidity: function destinationBlockchainID() view returns(bytes32)
func (_Nativetokensource *NativetokensourceSession) DestinationBlockchainID() ([32]byte, error) {
	return _Nativetokensource.Contract.DestinationBlockchainID(&_Nativetokensource.CallOpts)
}

// DestinationBlockchainID is a free data retrieval call binding the contract method 0x41d3014d.
//
// Solidity: function destinationBlockchainID() view returns(bytes32)
func (_Nativetokensource *NativetokensourceCallerSession) DestinationBlockchainID() ([32]byte, error) {
	return _Nativetokensource.Contract.DestinationBlockchainID(&_Nativetokensource.CallOpts)
}

// DestinationContractAddress is a free data retrieval call binding the contract method 0x04d6baf7.
//
// Solidity: function destinationContractAddress() view returns(address)
func (_Nativetokensource *NativetokensourceCaller) DestinationContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nativetokensource.contract.Call(opts, &out, "destinationContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DestinationContractAddress is a free data retrieval call binding the contract method 0x04d6baf7.
//
// Solidity: function destinationContractAddress() view returns(address)
func (_Nativetokensource *NativetokensourceSession) DestinationContractAddress() (common.Address, error) {
	return _Nativetokensource.Contract.DestinationContractAddress(&_Nativetokensource.CallOpts)
}

// DestinationContractAddress is a free data retrieval call binding the contract method 0x04d6baf7.
//
// Solidity: function destinationContractAddress() view returns(address)
func (_Nativetokensource *NativetokensourceCallerSession) DestinationContractAddress() (common.Address, error) {
	return _Nativetokensource.Contract.DestinationContractAddress(&_Nativetokensource.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Nativetokensource *NativetokensourceCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nativetokensource.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Nativetokensource *NativetokensourceSession) TeleporterMessenger() (common.Address, error) {
	return _Nativetokensource.Contract.TeleporterMessenger(&_Nativetokensource.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Nativetokensource *NativetokensourceCallerSession) TeleporterMessenger() (common.Address, error) {
	return _Nativetokensource.Contract.TeleporterMessenger(&_Nativetokensource.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_Nativetokensource *NativetokensourceTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Nativetokensource.contract.Transact(opts, "receiveTeleporterMessage", senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_Nativetokensource *NativetokensourceSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Nativetokensource.Contract.ReceiveTeleporterMessage(&_Nativetokensource.TransactOpts, senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_Nativetokensource *NativetokensourceTransactorSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Nativetokensource.Contract.ReceiveTeleporterMessage(&_Nativetokensource.TransactOpts, senderBlockchainID, senderAddress, message)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0x3c7bf510.
//
// Solidity: function transferToDestination(address recipient, address feeContractAddress, uint256 feeAmount) payable returns()
func (_Nativetokensource *NativetokensourceTransactor) TransferToDestination(opts *bind.TransactOpts, recipient common.Address, feeContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _Nativetokensource.contract.Transact(opts, "transferToDestination", recipient, feeContractAddress, feeAmount)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0x3c7bf510.
//
// Solidity: function transferToDestination(address recipient, address feeContractAddress, uint256 feeAmount) payable returns()
func (_Nativetokensource *NativetokensourceSession) TransferToDestination(recipient common.Address, feeContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _Nativetokensource.Contract.TransferToDestination(&_Nativetokensource.TransactOpts, recipient, feeContractAddress, feeAmount)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0x3c7bf510.
//
// Solidity: function transferToDestination(address recipient, address feeContractAddress, uint256 feeAmount) payable returns()
func (_Nativetokensource *NativetokensourceTransactorSession) TransferToDestination(recipient common.Address, feeContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _Nativetokensource.Contract.TransferToDestination(&_Nativetokensource.TransactOpts, recipient, feeContractAddress, feeAmount)
}

// NativetokensourceTransferToDestinationIterator is returned from FilterTransferToDestination and is used to iterate over the raw logs and unpacked data for TransferToDestination events raised by the Nativetokensource contract.
type NativetokensourceTransferToDestinationIterator struct {
	Event *NativetokensourceTransferToDestination // Event containing the contract specifics and raw log

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
func (it *NativetokensourceTransferToDestinationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativetokensourceTransferToDestination)
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
		it.Event = new(NativetokensourceTransferToDestination)
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
func (it *NativetokensourceTransferToDestinationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativetokensourceTransferToDestinationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativetokensourceTransferToDestination represents a TransferToDestination event raised by the Nativetokensource contract.
type NativetokensourceTransferToDestination struct {
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	FeeContractAddress  common.Address
	FeeAmount           *big.Int
	TeleporterMessageID *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTransferToDestination is a free log retrieval operation binding the contract event 0x32d7e0a933e2b2b5890fd273e9ba2e9293d619805a7ca90d11f41eddc56d2b97.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, uint256 amount, address feeContractAddress, uint256 feeAmount, uint256 teleporterMessageID)
func (_Nativetokensource *NativetokensourceFilterer) FilterTransferToDestination(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*NativetokensourceTransferToDestinationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Nativetokensource.contract.FilterLogs(opts, "TransferToDestination", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativetokensourceTransferToDestinationIterator{contract: _Nativetokensource.contract, event: "TransferToDestination", logs: logs, sub: sub}, nil
}

// WatchTransferToDestination is a free log subscription operation binding the contract event 0x32d7e0a933e2b2b5890fd273e9ba2e9293d619805a7ca90d11f41eddc56d2b97.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, uint256 amount, address feeContractAddress, uint256 feeAmount, uint256 teleporterMessageID)
func (_Nativetokensource *NativetokensourceFilterer) WatchTransferToDestination(opts *bind.WatchOpts, sink chan<- *NativetokensourceTransferToDestination, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Nativetokensource.contract.WatchLogs(opts, "TransferToDestination", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativetokensourceTransferToDestination)
				if err := _Nativetokensource.contract.UnpackLog(event, "TransferToDestination", log); err != nil {
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

// ParseTransferToDestination is a log parse operation binding the contract event 0x32d7e0a933e2b2b5890fd273e9ba2e9293d619805a7ca90d11f41eddc56d2b97.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, uint256 amount, address feeContractAddress, uint256 feeAmount, uint256 teleporterMessageID)
func (_Nativetokensource *NativetokensourceFilterer) ParseTransferToDestination(log types.Log) (*NativetokensourceTransferToDestination, error) {
	event := new(NativetokensourceTransferToDestination)
	if err := _Nativetokensource.contract.UnpackLog(event, "TransferToDestination", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativetokensourceUnlockTokensIterator is returned from FilterUnlockTokens and is used to iterate over the raw logs and unpacked data for UnlockTokens events raised by the Nativetokensource contract.
type NativetokensourceUnlockTokensIterator struct {
	Event *NativetokensourceUnlockTokens // Event containing the contract specifics and raw log

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
func (it *NativetokensourceUnlockTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativetokensourceUnlockTokens)
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
		it.Event = new(NativetokensourceUnlockTokens)
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
func (it *NativetokensourceUnlockTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativetokensourceUnlockTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativetokensourceUnlockTokens represents a UnlockTokens event raised by the Nativetokensource contract.
type NativetokensourceUnlockTokens struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlockTokens is a free log retrieval operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_Nativetokensource *NativetokensourceFilterer) FilterUnlockTokens(opts *bind.FilterOpts) (*NativetokensourceUnlockTokensIterator, error) {

	logs, sub, err := _Nativetokensource.contract.FilterLogs(opts, "UnlockTokens")
	if err != nil {
		return nil, err
	}
	return &NativetokensourceUnlockTokensIterator{contract: _Nativetokensource.contract, event: "UnlockTokens", logs: logs, sub: sub}, nil
}

// WatchUnlockTokens is a free log subscription operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_Nativetokensource *NativetokensourceFilterer) WatchUnlockTokens(opts *bind.WatchOpts, sink chan<- *NativetokensourceUnlockTokens) (event.Subscription, error) {

	logs, sub, err := _Nativetokensource.contract.WatchLogs(opts, "UnlockTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativetokensourceUnlockTokens)
				if err := _Nativetokensource.contract.UnpackLog(event, "UnlockTokens", log); err != nil {
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
func (_Nativetokensource *NativetokensourceFilterer) ParseUnlockTokens(log types.Log) (*NativetokensourceUnlockTokens, error) {
	event := new(NativetokensourceUnlockTokens)
	if err := _Nativetokensource.contract.UnpackLog(event, "UnlockTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
