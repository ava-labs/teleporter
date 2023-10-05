// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleporterblockhashreceiver

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

// TeleporterblockhashreceiverMetaData contains all meta data concerning the Teleporterblockhashreceiver contract.
var TeleporterblockhashreceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidWarpBlockHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedSourceChainID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"ReceiveBlockHash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"}],\"name\":\"receiveBlockHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeleporterblockhashreceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterblockhashreceiverMetaData.ABI instead.
var TeleporterblockhashreceiverABI = TeleporterblockhashreceiverMetaData.ABI

// Teleporterblockhashreceiver is an auto generated Go binding around an Ethereum contract.
type Teleporterblockhashreceiver struct {
	TeleporterblockhashreceiverCaller     // Read-only binding to the contract
	TeleporterblockhashreceiverTransactor // Write-only binding to the contract
	TeleporterblockhashreceiverFilterer   // Log filterer for contract events
}

// TeleporterblockhashreceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleporterblockhashreceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterblockhashreceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleporterblockhashreceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterblockhashreceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleporterblockhashreceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterblockhashreceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleporterblockhashreceiverSession struct {
	Contract     *Teleporterblockhashreceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TeleporterblockhashreceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleporterblockhashreceiverCallerSession struct {
	Contract *TeleporterblockhashreceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// TeleporterblockhashreceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleporterblockhashreceiverTransactorSession struct {
	Contract     *TeleporterblockhashreceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// TeleporterblockhashreceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleporterblockhashreceiverRaw struct {
	Contract *Teleporterblockhashreceiver // Generic contract binding to access the raw methods on
}

// TeleporterblockhashreceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleporterblockhashreceiverCallerRaw struct {
	Contract *TeleporterblockhashreceiverCaller // Generic read-only contract binding to access the raw methods on
}

// TeleporterblockhashreceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleporterblockhashreceiverTransactorRaw struct {
	Contract *TeleporterblockhashreceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleporterblockhashreceiver creates a new instance of Teleporterblockhashreceiver, bound to a specific deployed contract.
func NewTeleporterblockhashreceiver(address common.Address, backend bind.ContractBackend) (*Teleporterblockhashreceiver, error) {
	contract, err := bindTeleporterblockhashreceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Teleporterblockhashreceiver{TeleporterblockhashreceiverCaller: TeleporterblockhashreceiverCaller{contract: contract}, TeleporterblockhashreceiverTransactor: TeleporterblockhashreceiverTransactor{contract: contract}, TeleporterblockhashreceiverFilterer: TeleporterblockhashreceiverFilterer{contract: contract}}, nil
}

// NewTeleporterblockhashreceiverCaller creates a new read-only instance of Teleporterblockhashreceiver, bound to a specific deployed contract.
func NewTeleporterblockhashreceiverCaller(address common.Address, caller bind.ContractCaller) (*TeleporterblockhashreceiverCaller, error) {
	contract, err := bindTeleporterblockhashreceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterblockhashreceiverCaller{contract: contract}, nil
}

// NewTeleporterblockhashreceiverTransactor creates a new write-only instance of Teleporterblockhashreceiver, bound to a specific deployed contract.
func NewTeleporterblockhashreceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleporterblockhashreceiverTransactor, error) {
	contract, err := bindTeleporterblockhashreceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterblockhashreceiverTransactor{contract: contract}, nil
}

// NewTeleporterblockhashreceiverFilterer creates a new log filterer instance of Teleporterblockhashreceiver, bound to a specific deployed contract.
func NewTeleporterblockhashreceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleporterblockhashreceiverFilterer, error) {
	contract, err := bindTeleporterblockhashreceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleporterblockhashreceiverFilterer{contract: contract}, nil
}

// bindTeleporterblockhashreceiver binds a generic wrapper to an already deployed contract.
func bindTeleporterblockhashreceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleporterblockhashreceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Teleporterblockhashreceiver.Contract.TeleporterblockhashreceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Teleporterblockhashreceiver.Contract.TeleporterblockhashreceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Teleporterblockhashreceiver.Contract.TeleporterblockhashreceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Teleporterblockhashreceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Teleporterblockhashreceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Teleporterblockhashreceiver.Contract.contract.Transact(opts, method, params...)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Teleporterblockhashreceiver.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverSession) WARPMESSENGER() (common.Address, error) {
	return _Teleporterblockhashreceiver.Contract.WARPMESSENGER(&_Teleporterblockhashreceiver.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverCallerSession) WARPMESSENGER() (common.Address, error) {
	return _Teleporterblockhashreceiver.Contract.WARPMESSENGER(&_Teleporterblockhashreceiver.CallOpts)
}

// ReceiveBlockHash is a paid mutator transaction binding the contract method 0x30b52559.
//
// Solidity: function receiveBlockHash(uint32 messageIndex, bytes32 sourceChainID) returns()
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverTransactor) ReceiveBlockHash(opts *bind.TransactOpts, messageIndex uint32, sourceChainID [32]byte) (*types.Transaction, error) {
	return _Teleporterblockhashreceiver.contract.Transact(opts, "receiveBlockHash", messageIndex, sourceChainID)
}

// ReceiveBlockHash is a paid mutator transaction binding the contract method 0x30b52559.
//
// Solidity: function receiveBlockHash(uint32 messageIndex, bytes32 sourceChainID) returns()
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverSession) ReceiveBlockHash(messageIndex uint32, sourceChainID [32]byte) (*types.Transaction, error) {
	return _Teleporterblockhashreceiver.Contract.ReceiveBlockHash(&_Teleporterblockhashreceiver.TransactOpts, messageIndex, sourceChainID)
}

// ReceiveBlockHash is a paid mutator transaction binding the contract method 0x30b52559.
//
// Solidity: function receiveBlockHash(uint32 messageIndex, bytes32 sourceChainID) returns()
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverTransactorSession) ReceiveBlockHash(messageIndex uint32, sourceChainID [32]byte) (*types.Transaction, error) {
	return _Teleporterblockhashreceiver.Contract.ReceiveBlockHash(&_Teleporterblockhashreceiver.TransactOpts, messageIndex, sourceChainID)
}

// TeleporterblockhashreceiverReceiveBlockHashIterator is returned from FilterReceiveBlockHash and is used to iterate over the raw logs and unpacked data for ReceiveBlockHash events raised by the Teleporterblockhashreceiver contract.
type TeleporterblockhashreceiverReceiveBlockHashIterator struct {
	Event *TeleporterblockhashreceiverReceiveBlockHash // Event containing the contract specifics and raw log

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
func (it *TeleporterblockhashreceiverReceiveBlockHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterblockhashreceiverReceiveBlockHash)
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
		it.Event = new(TeleporterblockhashreceiverReceiveBlockHash)
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
func (it *TeleporterblockhashreceiverReceiveBlockHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterblockhashreceiverReceiveBlockHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterblockhashreceiverReceiveBlockHash represents a ReceiveBlockHash event raised by the Teleporterblockhashreceiver contract.
type TeleporterblockhashreceiverReceiveBlockHash struct {
	OriginChainID [32]byte
	BlockHash     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceiveBlockHash is a free log retrieval operation binding the contract event 0x7770e5f72465e9b05c8076c3f2eac70898abe6a84f0259307d127c13e2a1a4e4.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, bytes32 indexed blockHash)
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverFilterer) FilterReceiveBlockHash(opts *bind.FilterOpts, originChainID [][32]byte, blockHash [][32]byte) (*TeleporterblockhashreceiverReceiveBlockHashIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}

	logs, sub, err := _Teleporterblockhashreceiver.contract.FilterLogs(opts, "ReceiveBlockHash", originChainIDRule, blockHashRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterblockhashreceiverReceiveBlockHashIterator{contract: _Teleporterblockhashreceiver.contract, event: "ReceiveBlockHash", logs: logs, sub: sub}, nil
}

// WatchReceiveBlockHash is a free log subscription operation binding the contract event 0x7770e5f72465e9b05c8076c3f2eac70898abe6a84f0259307d127c13e2a1a4e4.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, bytes32 indexed blockHash)
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverFilterer) WatchReceiveBlockHash(opts *bind.WatchOpts, sink chan<- *TeleporterblockhashreceiverReceiveBlockHash, originChainID [][32]byte, blockHash [][32]byte) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}

	logs, sub, err := _Teleporterblockhashreceiver.contract.WatchLogs(opts, "ReceiveBlockHash", originChainIDRule, blockHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterblockhashreceiverReceiveBlockHash)
				if err := _Teleporterblockhashreceiver.contract.UnpackLog(event, "ReceiveBlockHash", log); err != nil {
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

// ParseReceiveBlockHash is a log parse operation binding the contract event 0x7770e5f72465e9b05c8076c3f2eac70898abe6a84f0259307d127c13e2a1a4e4.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, bytes32 indexed blockHash)
func (_Teleporterblockhashreceiver *TeleporterblockhashreceiverFilterer) ParseReceiveBlockHash(log types.Log) (*TeleporterblockhashreceiverReceiveBlockHash, error) {
	event := new(TeleporterblockhashreceiverReceiveBlockHash)
	if err := _Teleporterblockhashreceiver.contract.UnpackLog(event, "ReceiveBlockHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
