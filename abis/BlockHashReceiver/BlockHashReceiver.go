// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockhashreceiver

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

// BlockhashreceiverMetaData contains all meta data concerning the Blockhashreceiver contract.
var BlockhashreceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"publisherChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"publisherContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidSourceChainID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSourceChainPublisher\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"ReceiveBlockHash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getLatestBlockInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourcePublisherContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BlockhashreceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockhashreceiverMetaData.ABI instead.
var BlockhashreceiverABI = BlockhashreceiverMetaData.ABI

// Blockhashreceiver is an auto generated Go binding around an Ethereum contract.
type Blockhashreceiver struct {
	BlockhashreceiverCaller     // Read-only binding to the contract
	BlockhashreceiverTransactor // Write-only binding to the contract
	BlockhashreceiverFilterer   // Log filterer for contract events
}

// BlockhashreceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockhashreceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockhashreceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockhashreceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockhashreceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockhashreceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockhashreceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockhashreceiverSession struct {
	Contract     *Blockhashreceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BlockhashreceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockhashreceiverCallerSession struct {
	Contract *BlockhashreceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BlockhashreceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockhashreceiverTransactorSession struct {
	Contract     *BlockhashreceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BlockhashreceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockhashreceiverRaw struct {
	Contract *Blockhashreceiver // Generic contract binding to access the raw methods on
}

// BlockhashreceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockhashreceiverCallerRaw struct {
	Contract *BlockhashreceiverCaller // Generic read-only contract binding to access the raw methods on
}

// BlockhashreceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockhashreceiverTransactorRaw struct {
	Contract *BlockhashreceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockhashreceiver creates a new instance of Blockhashreceiver, bound to a specific deployed contract.
func NewBlockhashreceiver(address common.Address, backend bind.ContractBackend) (*Blockhashreceiver, error) {
	contract, err := bindBlockhashreceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockhashreceiver{BlockhashreceiverCaller: BlockhashreceiverCaller{contract: contract}, BlockhashreceiverTransactor: BlockhashreceiverTransactor{contract: contract}, BlockhashreceiverFilterer: BlockhashreceiverFilterer{contract: contract}}, nil
}

// NewBlockhashreceiverCaller creates a new read-only instance of Blockhashreceiver, bound to a specific deployed contract.
func NewBlockhashreceiverCaller(address common.Address, caller bind.ContractCaller) (*BlockhashreceiverCaller, error) {
	contract, err := bindBlockhashreceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockhashreceiverCaller{contract: contract}, nil
}

// NewBlockhashreceiverTransactor creates a new write-only instance of Blockhashreceiver, bound to a specific deployed contract.
func NewBlockhashreceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockhashreceiverTransactor, error) {
	contract, err := bindBlockhashreceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockhashreceiverTransactor{contract: contract}, nil
}

// NewBlockhashreceiverFilterer creates a new log filterer instance of Blockhashreceiver, bound to a specific deployed contract.
func NewBlockhashreceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockhashreceiverFilterer, error) {
	contract, err := bindBlockhashreceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockhashreceiverFilterer{contract: contract}, nil
}

// bindBlockhashreceiver binds a generic wrapper to an already deployed contract.
func bindBlockhashreceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockhashreceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockhashreceiver *BlockhashreceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockhashreceiver.Contract.BlockhashreceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockhashreceiver *BlockhashreceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockhashreceiver.Contract.BlockhashreceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockhashreceiver *BlockhashreceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockhashreceiver.Contract.BlockhashreceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockhashreceiver *BlockhashreceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockhashreceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockhashreceiver *BlockhashreceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockhashreceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockhashreceiver *BlockhashreceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockhashreceiver.Contract.contract.Transact(opts, method, params...)
}

// GetLatestBlockInfo is a free data retrieval call binding the contract method 0xb17810be.
//
// Solidity: function getLatestBlockInfo() view returns(uint256 height, bytes32 hash)
func (_Blockhashreceiver *BlockhashreceiverCaller) GetLatestBlockInfo(opts *bind.CallOpts) (struct {
	Height *big.Int
	Hash   [32]byte
}, error) {
	var out []interface{}
	err := _Blockhashreceiver.contract.Call(opts, &out, "getLatestBlockInfo")

	outstruct := new(struct {
		Height *big.Int
		Hash   [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Height = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Hash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetLatestBlockInfo is a free data retrieval call binding the contract method 0xb17810be.
//
// Solidity: function getLatestBlockInfo() view returns(uint256 height, bytes32 hash)
func (_Blockhashreceiver *BlockhashreceiverSession) GetLatestBlockInfo() (struct {
	Height *big.Int
	Hash   [32]byte
}, error) {
	return _Blockhashreceiver.Contract.GetLatestBlockInfo(&_Blockhashreceiver.CallOpts)
}

// GetLatestBlockInfo is a free data retrieval call binding the contract method 0xb17810be.
//
// Solidity: function getLatestBlockInfo() view returns(uint256 height, bytes32 hash)
func (_Blockhashreceiver *BlockhashreceiverCallerSession) GetLatestBlockInfo() (struct {
	Height *big.Int
	Hash   [32]byte
}, error) {
	return _Blockhashreceiver.Contract.GetLatestBlockInfo(&_Blockhashreceiver.CallOpts)
}

// LatestBlockHash is a free data retrieval call binding the contract method 0x6c4f6ba9.
//
// Solidity: function latestBlockHash() view returns(bytes32)
func (_Blockhashreceiver *BlockhashreceiverCaller) LatestBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blockhashreceiver.contract.Call(opts, &out, "latestBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestBlockHash is a free data retrieval call binding the contract method 0x6c4f6ba9.
//
// Solidity: function latestBlockHash() view returns(bytes32)
func (_Blockhashreceiver *BlockhashreceiverSession) LatestBlockHash() ([32]byte, error) {
	return _Blockhashreceiver.Contract.LatestBlockHash(&_Blockhashreceiver.CallOpts)
}

// LatestBlockHash is a free data retrieval call binding the contract method 0x6c4f6ba9.
//
// Solidity: function latestBlockHash() view returns(bytes32)
func (_Blockhashreceiver *BlockhashreceiverCallerSession) LatestBlockHash() ([32]byte, error) {
	return _Blockhashreceiver.Contract.LatestBlockHash(&_Blockhashreceiver.CallOpts)
}

// LatestBlockHeight is a free data retrieval call binding the contract method 0xf3f39ee5.
//
// Solidity: function latestBlockHeight() view returns(uint256)
func (_Blockhashreceiver *BlockhashreceiverCaller) LatestBlockHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blockhashreceiver.contract.Call(opts, &out, "latestBlockHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockHeight is a free data retrieval call binding the contract method 0xf3f39ee5.
//
// Solidity: function latestBlockHeight() view returns(uint256)
func (_Blockhashreceiver *BlockhashreceiverSession) LatestBlockHeight() (*big.Int, error) {
	return _Blockhashreceiver.Contract.LatestBlockHeight(&_Blockhashreceiver.CallOpts)
}

// LatestBlockHeight is a free data retrieval call binding the contract method 0xf3f39ee5.
//
// Solidity: function latestBlockHeight() view returns(uint256)
func (_Blockhashreceiver *BlockhashreceiverCallerSession) LatestBlockHeight() (*big.Int, error) {
	return _Blockhashreceiver.Contract.LatestBlockHeight(&_Blockhashreceiver.CallOpts)
}

// SourceChainID is a free data retrieval call binding the contract method 0x4c335368.
//
// Solidity: function sourceChainID() view returns(bytes32)
func (_Blockhashreceiver *BlockhashreceiverCaller) SourceChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blockhashreceiver.contract.Call(opts, &out, "sourceChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SourceChainID is a free data retrieval call binding the contract method 0x4c335368.
//
// Solidity: function sourceChainID() view returns(bytes32)
func (_Blockhashreceiver *BlockhashreceiverSession) SourceChainID() ([32]byte, error) {
	return _Blockhashreceiver.Contract.SourceChainID(&_Blockhashreceiver.CallOpts)
}

// SourceChainID is a free data retrieval call binding the contract method 0x4c335368.
//
// Solidity: function sourceChainID() view returns(bytes32)
func (_Blockhashreceiver *BlockhashreceiverCallerSession) SourceChainID() ([32]byte, error) {
	return _Blockhashreceiver.Contract.SourceChainID(&_Blockhashreceiver.CallOpts)
}

// SourcePublisherContractAddress is a free data retrieval call binding the contract method 0x79a0710c.
//
// Solidity: function sourcePublisherContractAddress() view returns(address)
func (_Blockhashreceiver *BlockhashreceiverCaller) SourcePublisherContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockhashreceiver.contract.Call(opts, &out, "sourcePublisherContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SourcePublisherContractAddress is a free data retrieval call binding the contract method 0x79a0710c.
//
// Solidity: function sourcePublisherContractAddress() view returns(address)
func (_Blockhashreceiver *BlockhashreceiverSession) SourcePublisherContractAddress() (common.Address, error) {
	return _Blockhashreceiver.Contract.SourcePublisherContractAddress(&_Blockhashreceiver.CallOpts)
}

// SourcePublisherContractAddress is a free data retrieval call binding the contract method 0x79a0710c.
//
// Solidity: function sourcePublisherContractAddress() view returns(address)
func (_Blockhashreceiver *BlockhashreceiverCallerSession) SourcePublisherContractAddress() (common.Address, error) {
	return _Blockhashreceiver.Contract.SourcePublisherContractAddress(&_Blockhashreceiver.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Blockhashreceiver *BlockhashreceiverCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockhashreceiver.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Blockhashreceiver *BlockhashreceiverSession) TeleporterMessenger() (common.Address, error) {
	return _Blockhashreceiver.Contract.TeleporterMessenger(&_Blockhashreceiver.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Blockhashreceiver *BlockhashreceiverCallerSession) TeleporterMessenger() (common.Address, error) {
	return _Blockhashreceiver.Contract.TeleporterMessenger(&_Blockhashreceiver.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_Blockhashreceiver *BlockhashreceiverTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Blockhashreceiver.contract.Transact(opts, "receiveTeleporterMessage", originChainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_Blockhashreceiver *BlockhashreceiverSession) ReceiveTeleporterMessage(originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Blockhashreceiver.Contract.ReceiveTeleporterMessage(&_Blockhashreceiver.TransactOpts, originChainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_Blockhashreceiver *BlockhashreceiverTransactorSession) ReceiveTeleporterMessage(originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Blockhashreceiver.Contract.ReceiveTeleporterMessage(&_Blockhashreceiver.TransactOpts, originChainID, originSenderAddress, message)
}

// BlockhashreceiverReceiveBlockHashIterator is returned from FilterReceiveBlockHash and is used to iterate over the raw logs and unpacked data for ReceiveBlockHash events raised by the Blockhashreceiver contract.
type BlockhashreceiverReceiveBlockHashIterator struct {
	Event *BlockhashreceiverReceiveBlockHash // Event containing the contract specifics and raw log

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
func (it *BlockhashreceiverReceiveBlockHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockhashreceiverReceiveBlockHash)
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
		it.Event = new(BlockhashreceiverReceiveBlockHash)
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
func (it *BlockhashreceiverReceiveBlockHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockhashreceiverReceiveBlockHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockhashreceiverReceiveBlockHash represents a ReceiveBlockHash event raised by the Blockhashreceiver contract.
type BlockhashreceiverReceiveBlockHash struct {
	OriginChainID       [32]byte
	OriginSenderAddress common.Address
	BlockHeight         *big.Int
	BlockHash           [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReceiveBlockHash is a free log retrieval operation binding the contract event 0x0bca78aa82d7575f42e4b4b2fe04765a4b2f3661786403788ce987e065ac590d.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, address indexed originSenderAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_Blockhashreceiver *BlockhashreceiverFilterer) FilterReceiveBlockHash(opts *bind.FilterOpts, originChainID [][32]byte, originSenderAddress []common.Address, blockHeight []*big.Int) (*BlockhashreceiverReceiveBlockHashIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _Blockhashreceiver.contract.FilterLogs(opts, "ReceiveBlockHash", originChainIDRule, originSenderAddressRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return &BlockhashreceiverReceiveBlockHashIterator{contract: _Blockhashreceiver.contract, event: "ReceiveBlockHash", logs: logs, sub: sub}, nil
}

// WatchReceiveBlockHash is a free log subscription operation binding the contract event 0x0bca78aa82d7575f42e4b4b2fe04765a4b2f3661786403788ce987e065ac590d.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, address indexed originSenderAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_Blockhashreceiver *BlockhashreceiverFilterer) WatchReceiveBlockHash(opts *bind.WatchOpts, sink chan<- *BlockhashreceiverReceiveBlockHash, originChainID [][32]byte, originSenderAddress []common.Address, blockHeight []*big.Int) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _Blockhashreceiver.contract.WatchLogs(opts, "ReceiveBlockHash", originChainIDRule, originSenderAddressRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockhashreceiverReceiveBlockHash)
				if err := _Blockhashreceiver.contract.UnpackLog(event, "ReceiveBlockHash", log); err != nil {
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

// ParseReceiveBlockHash is a log parse operation binding the contract event 0x0bca78aa82d7575f42e4b4b2fe04765a4b2f3661786403788ce987e065ac590d.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, address indexed originSenderAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_Blockhashreceiver *BlockhashreceiverFilterer) ParseReceiveBlockHash(log types.Log) (*BlockhashreceiverReceiveBlockHash, error) {
	event := new(BlockhashreceiverReceiveBlockHash)
	if err := _Blockhashreceiver.contract.UnpackLog(event, "ReceiveBlockHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
