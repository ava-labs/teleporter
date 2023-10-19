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

// TeleporterBlockHashReceiverMetaData contains all meta data concerning the TeleporterBlockHashReceiver contract.
var TeleporterBlockHashReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ReceiverReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderReentrancy\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"ReceiveBlockHash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"receiveBlockHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeleporterBlockHashReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterBlockHashReceiverMetaData.ABI instead.
var TeleporterBlockHashReceiverABI = TeleporterBlockHashReceiverMetaData.ABI

// TeleporterBlockHashReceiver is an auto generated Go binding around an Ethereum contract.
type TeleporterBlockHashReceiver struct {
	TeleporterBlockHashReceiverCaller     // Read-only binding to the contract
	TeleporterBlockHashReceiverTransactor // Write-only binding to the contract
	TeleporterBlockHashReceiverFilterer   // Log filterer for contract events
}

// TeleporterBlockHashReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleporterBlockHashReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterBlockHashReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleporterBlockHashReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterBlockHashReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleporterBlockHashReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterBlockHashReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleporterBlockHashReceiverSession struct {
	Contract     *TeleporterBlockHashReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TeleporterBlockHashReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleporterBlockHashReceiverCallerSession struct {
	Contract *TeleporterBlockHashReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// TeleporterBlockHashReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleporterBlockHashReceiverTransactorSession struct {
	Contract     *TeleporterBlockHashReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// TeleporterBlockHashReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleporterBlockHashReceiverRaw struct {
	Contract *TeleporterBlockHashReceiver // Generic contract binding to access the raw methods on
}

// TeleporterBlockHashReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleporterBlockHashReceiverCallerRaw struct {
	Contract *TeleporterBlockHashReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// TeleporterBlockHashReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleporterBlockHashReceiverTransactorRaw struct {
	Contract *TeleporterBlockHashReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleporterBlockHashReceiver creates a new instance of TeleporterBlockHashReceiver, bound to a specific deployed contract.
func NewTeleporterBlockHashReceiver(address common.Address, backend bind.ContractBackend) (*TeleporterBlockHashReceiver, error) {
	contract, err := bindTeleporterBlockHashReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeleporterBlockHashReceiver{TeleporterBlockHashReceiverCaller: TeleporterBlockHashReceiverCaller{contract: contract}, TeleporterBlockHashReceiverTransactor: TeleporterBlockHashReceiverTransactor{contract: contract}, TeleporterBlockHashReceiverFilterer: TeleporterBlockHashReceiverFilterer{contract: contract}}, nil
}

// NewTeleporterBlockHashReceiverCaller creates a new read-only instance of TeleporterBlockHashReceiver, bound to a specific deployed contract.
func NewTeleporterBlockHashReceiverCaller(address common.Address, caller bind.ContractCaller) (*TeleporterBlockHashReceiverCaller, error) {
	contract, err := bindTeleporterBlockHashReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterBlockHashReceiverCaller{contract: contract}, nil
}

// NewTeleporterBlockHashReceiverTransactor creates a new write-only instance of TeleporterBlockHashReceiver, bound to a specific deployed contract.
func NewTeleporterBlockHashReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleporterBlockHashReceiverTransactor, error) {
	contract, err := bindTeleporterBlockHashReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterBlockHashReceiverTransactor{contract: contract}, nil
}

// NewTeleporterBlockHashReceiverFilterer creates a new log filterer instance of TeleporterBlockHashReceiver, bound to a specific deployed contract.
func NewTeleporterBlockHashReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleporterBlockHashReceiverFilterer, error) {
	contract, err := bindTeleporterBlockHashReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleporterBlockHashReceiverFilterer{contract: contract}, nil
}

// bindTeleporterBlockHashReceiver binds a generic wrapper to an already deployed contract.
func bindTeleporterBlockHashReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleporterBlockHashReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterBlockHashReceiver.Contract.TeleporterBlockHashReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterBlockHashReceiver.Contract.TeleporterBlockHashReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterBlockHashReceiver.Contract.TeleporterBlockHashReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterBlockHashReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterBlockHashReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterBlockHashReceiver.Contract.contract.Transact(opts, method, params...)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterBlockHashReceiver.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterBlockHashReceiver.Contract.WARPMESSENGER(&_TeleporterBlockHashReceiver.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverCallerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterBlockHashReceiver.Contract.WARPMESSENGER(&_TeleporterBlockHashReceiver.CallOpts)
}

// ReceiveBlockHash is a paid mutator transaction binding the contract method 0x303bb425.
//
// Solidity: function receiveBlockHash(uint32 messageIndex) returns()
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverTransactor) ReceiveBlockHash(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _TeleporterBlockHashReceiver.contract.Transact(opts, "receiveBlockHash", messageIndex)
}

// ReceiveBlockHash is a paid mutator transaction binding the contract method 0x303bb425.
//
// Solidity: function receiveBlockHash(uint32 messageIndex) returns()
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverSession) ReceiveBlockHash(messageIndex uint32) (*types.Transaction, error) {
	return _TeleporterBlockHashReceiver.Contract.ReceiveBlockHash(&_TeleporterBlockHashReceiver.TransactOpts, messageIndex)
}

// ReceiveBlockHash is a paid mutator transaction binding the contract method 0x303bb425.
//
// Solidity: function receiveBlockHash(uint32 messageIndex) returns()
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverTransactorSession) ReceiveBlockHash(messageIndex uint32) (*types.Transaction, error) {
	return _TeleporterBlockHashReceiver.Contract.ReceiveBlockHash(&_TeleporterBlockHashReceiver.TransactOpts, messageIndex)
}

// TeleporterBlockHashReceiverReceiveBlockHashIterator is returned from FilterReceiveBlockHash and is used to iterate over the raw logs and unpacked data for ReceiveBlockHash events raised by the TeleporterBlockHashReceiver contract.
type TeleporterBlockHashReceiverReceiveBlockHashIterator struct {
	Event *TeleporterBlockHashReceiverReceiveBlockHash // Event containing the contract specifics and raw log

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
func (it *TeleporterBlockHashReceiverReceiveBlockHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterBlockHashReceiverReceiveBlockHash)
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
		it.Event = new(TeleporterBlockHashReceiverReceiveBlockHash)
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
func (it *TeleporterBlockHashReceiverReceiveBlockHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterBlockHashReceiverReceiveBlockHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterBlockHashReceiverReceiveBlockHash represents a ReceiveBlockHash event raised by the TeleporterBlockHashReceiver contract.
type TeleporterBlockHashReceiverReceiveBlockHash struct {
	OriginChainID [32]byte
	BlockHash     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceiveBlockHash is a free log retrieval operation binding the contract event 0x7770e5f72465e9b05c8076c3f2eac70898abe6a84f0259307d127c13e2a1a4e4.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, bytes32 indexed blockHash)
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverFilterer) FilterReceiveBlockHash(opts *bind.FilterOpts, originChainID [][32]byte, blockHash [][32]byte) (*TeleporterBlockHashReceiverReceiveBlockHashIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}

	logs, sub, err := _TeleporterBlockHashReceiver.contract.FilterLogs(opts, "ReceiveBlockHash", originChainIDRule, blockHashRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterBlockHashReceiverReceiveBlockHashIterator{contract: _TeleporterBlockHashReceiver.contract, event: "ReceiveBlockHash", logs: logs, sub: sub}, nil
}

// WatchReceiveBlockHash is a free log subscription operation binding the contract event 0x7770e5f72465e9b05c8076c3f2eac70898abe6a84f0259307d127c13e2a1a4e4.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, bytes32 indexed blockHash)
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverFilterer) WatchReceiveBlockHash(opts *bind.WatchOpts, sink chan<- *TeleporterBlockHashReceiverReceiveBlockHash, originChainID [][32]byte, blockHash [][32]byte) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}

	logs, sub, err := _TeleporterBlockHashReceiver.contract.WatchLogs(opts, "ReceiveBlockHash", originChainIDRule, blockHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterBlockHashReceiverReceiveBlockHash)
				if err := _TeleporterBlockHashReceiver.contract.UnpackLog(event, "ReceiveBlockHash", log); err != nil {
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
func (_TeleporterBlockHashReceiver *TeleporterBlockHashReceiverFilterer) ParseReceiveBlockHash(log types.Log) (*TeleporterBlockHashReceiverReceiveBlockHash, error) {
	event := new(TeleporterBlockHashReceiverReceiveBlockHash)
	if err := _TeleporterBlockHashReceiver.contract.UnpackLog(event, "ReceiveBlockHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
