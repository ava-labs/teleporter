// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleporter_block_hash

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

// TeleporterBlockHashMetaData contains all meta data concerning the TeleporterBlockHash contract.
var TeleporterBlockHashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidWarpBlockHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedSourceChainID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"ReceiveBlockHash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"}],\"name\":\"receiveBlockHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeleporterBlockHashABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterBlockHashMetaData.ABI instead.
var TeleporterBlockHashABI = TeleporterBlockHashMetaData.ABI

// TeleporterBlockHash is an auto generated Go binding around an Ethereum contract.
type TeleporterBlockHash struct {
	TeleporterBlockHashCaller     // Read-only binding to the contract
	TeleporterBlockHashTransactor // Write-only binding to the contract
	TeleporterBlockHashFilterer   // Log filterer for contract events
}

// TeleporterBlockHashCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleporterBlockHashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterBlockHashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleporterBlockHashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterBlockHashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleporterBlockHashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterBlockHashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleporterBlockHashSession struct {
	Contract     *TeleporterBlockHash // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TeleporterBlockHashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleporterBlockHashCallerSession struct {
	Contract *TeleporterBlockHashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TeleporterBlockHashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleporterBlockHashTransactorSession struct {
	Contract     *TeleporterBlockHashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TeleporterBlockHashRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleporterBlockHashRaw struct {
	Contract *TeleporterBlockHash // Generic contract binding to access the raw methods on
}

// TeleporterBlockHashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleporterBlockHashCallerRaw struct {
	Contract *TeleporterBlockHashCaller // Generic read-only contract binding to access the raw methods on
}

// TeleporterBlockHashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleporterBlockHashTransactorRaw struct {
	Contract *TeleporterBlockHashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleporterBlockHash creates a new instance of TeleporterBlockHash, bound to a specific deployed contract.
func NewTeleporterBlockHash(address common.Address, backend bind.ContractBackend) (*TeleporterBlockHash, error) {
	contract, err := bindTeleporterBlockHash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeleporterBlockHash{TeleporterBlockHashCaller: TeleporterBlockHashCaller{contract: contract}, TeleporterBlockHashTransactor: TeleporterBlockHashTransactor{contract: contract}, TeleporterBlockHashFilterer: TeleporterBlockHashFilterer{contract: contract}}, nil
}

// NewTeleporterBlockHashCaller creates a new read-only instance of TeleporterBlockHash, bound to a specific deployed contract.
func NewTeleporterBlockHashCaller(address common.Address, caller bind.ContractCaller) (*TeleporterBlockHashCaller, error) {
	contract, err := bindTeleporterBlockHash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterBlockHashCaller{contract: contract}, nil
}

// NewTeleporterBlockHashTransactor creates a new write-only instance of TeleporterBlockHash, bound to a specific deployed contract.
func NewTeleporterBlockHashTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleporterBlockHashTransactor, error) {
	contract, err := bindTeleporterBlockHash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterBlockHashTransactor{contract: contract}, nil
}

// NewTeleporterBlockHashFilterer creates a new log filterer instance of TeleporterBlockHash, bound to a specific deployed contract.
func NewTeleporterBlockHashFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleporterBlockHashFilterer, error) {
	contract, err := bindTeleporterBlockHash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleporterBlockHashFilterer{contract: contract}, nil
}

// bindTeleporterBlockHash binds a generic wrapper to an already deployed contract.
func bindTeleporterBlockHash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleporterBlockHashMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterBlockHash *TeleporterBlockHashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterBlockHash.Contract.TeleporterBlockHashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterBlockHash *TeleporterBlockHashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterBlockHash.Contract.TeleporterBlockHashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterBlockHash *TeleporterBlockHashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterBlockHash.Contract.TeleporterBlockHashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterBlockHash *TeleporterBlockHashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterBlockHash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterBlockHash *TeleporterBlockHashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterBlockHash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterBlockHash *TeleporterBlockHashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterBlockHash.Contract.contract.Transact(opts, method, params...)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterBlockHash *TeleporterBlockHashCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterBlockHash.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterBlockHash *TeleporterBlockHashSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterBlockHash.Contract.WARPMESSENGER(&_TeleporterBlockHash.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterBlockHash *TeleporterBlockHashCallerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterBlockHash.Contract.WARPMESSENGER(&_TeleporterBlockHash.CallOpts)
}

// ReceiveBlockHash is a paid mutator transaction binding the contract method 0x30b52559.
//
// Solidity: function receiveBlockHash(uint32 messageIndex, bytes32 sourceChainID) returns()
func (_TeleporterBlockHash *TeleporterBlockHashTransactor) ReceiveBlockHash(opts *bind.TransactOpts, messageIndex uint32, sourceChainID [32]byte) (*types.Transaction, error) {
	return _TeleporterBlockHash.contract.Transact(opts, "receiveBlockHash", messageIndex, sourceChainID)
}

// ReceiveBlockHash is a paid mutator transaction binding the contract method 0x30b52559.
//
// Solidity: function receiveBlockHash(uint32 messageIndex, bytes32 sourceChainID) returns()
func (_TeleporterBlockHash *TeleporterBlockHashSession) ReceiveBlockHash(messageIndex uint32, sourceChainID [32]byte) (*types.Transaction, error) {
	return _TeleporterBlockHash.Contract.ReceiveBlockHash(&_TeleporterBlockHash.TransactOpts, messageIndex, sourceChainID)
}

// ReceiveBlockHash is a paid mutator transaction binding the contract method 0x30b52559.
//
// Solidity: function receiveBlockHash(uint32 messageIndex, bytes32 sourceChainID) returns()
func (_TeleporterBlockHash *TeleporterBlockHashTransactorSession) ReceiveBlockHash(messageIndex uint32, sourceChainID [32]byte) (*types.Transaction, error) {
	return _TeleporterBlockHash.Contract.ReceiveBlockHash(&_TeleporterBlockHash.TransactOpts, messageIndex, sourceChainID)
}

// TeleporterBlockHashReceiveBlockHashIterator is returned from FilterReceiveBlockHash and is used to iterate over the raw logs and unpacked data for ReceiveBlockHash events raised by the TeleporterBlockHash contract.
type TeleporterBlockHashReceiveBlockHashIterator struct {
	Event *TeleporterBlockHashReceiveBlockHash // Event containing the contract specifics and raw log

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
func (it *TeleporterBlockHashReceiveBlockHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterBlockHashReceiveBlockHash)
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
		it.Event = new(TeleporterBlockHashReceiveBlockHash)
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
func (it *TeleporterBlockHashReceiveBlockHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterBlockHashReceiveBlockHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterBlockHashReceiveBlockHash represents a ReceiveBlockHash event raised by the TeleporterBlockHash contract.
type TeleporterBlockHashReceiveBlockHash struct {
	OriginChainID [32]byte
	BlockHash     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceiveBlockHash is a free log retrieval operation binding the contract event 0x7770e5f72465e9b05c8076c3f2eac70898abe6a84f0259307d127c13e2a1a4e4.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, bytes32 indexed blockHash)
func (_TeleporterBlockHash *TeleporterBlockHashFilterer) FilterReceiveBlockHash(opts *bind.FilterOpts, originChainID [][32]byte, blockHash [][32]byte) (*TeleporterBlockHashReceiveBlockHashIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}

	logs, sub, err := _TeleporterBlockHash.contract.FilterLogs(opts, "ReceiveBlockHash", originChainIDRule, blockHashRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterBlockHashReceiveBlockHashIterator{contract: _TeleporterBlockHash.contract, event: "ReceiveBlockHash", logs: logs, sub: sub}, nil
}

// WatchReceiveBlockHash is a free log subscription operation binding the contract event 0x7770e5f72465e9b05c8076c3f2eac70898abe6a84f0259307d127c13e2a1a4e4.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, bytes32 indexed blockHash)
func (_TeleporterBlockHash *TeleporterBlockHashFilterer) WatchReceiveBlockHash(opts *bind.WatchOpts, sink chan<- *TeleporterBlockHashReceiveBlockHash, originChainID [][32]byte, blockHash [][32]byte) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}

	logs, sub, err := _TeleporterBlockHash.contract.WatchLogs(opts, "ReceiveBlockHash", originChainIDRule, blockHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterBlockHashReceiveBlockHash)
				if err := _TeleporterBlockHash.contract.UnpackLog(event, "ReceiveBlockHash", log); err != nil {
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
func (_TeleporterBlockHash *TeleporterBlockHashFilterer) ParseReceiveBlockHash(log types.Log) (*TeleporterBlockHashReceiveBlockHash, error) {
	event := new(TeleporterBlockHashReceiveBlockHash)
	if err := _TeleporterBlockHash.contract.UnpackLog(event, "ReceiveBlockHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
