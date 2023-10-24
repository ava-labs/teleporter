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

// BlockHashReceiverMetaData contains all meta data concerning the BlockHashReceiver contract.
var BlockHashReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"publisherChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"publisherContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"ReceiveBlockHash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getLatestBlockInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourcePublisherContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BlockHashReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockHashReceiverMetaData.ABI instead.
var BlockHashReceiverABI = BlockHashReceiverMetaData.ABI

// BlockHashReceiver is an auto generated Go binding around an Ethereum contract.
type BlockHashReceiver struct {
	BlockHashReceiverCaller     // Read-only binding to the contract
	BlockHashReceiverTransactor // Write-only binding to the contract
	BlockHashReceiverFilterer   // Log filterer for contract events
}

// BlockHashReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockHashReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockHashReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockHashReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockHashReceiverSession struct {
	Contract     *BlockHashReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BlockHashReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockHashReceiverCallerSession struct {
	Contract *BlockHashReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BlockHashReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockHashReceiverTransactorSession struct {
	Contract     *BlockHashReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BlockHashReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockHashReceiverRaw struct {
	Contract *BlockHashReceiver // Generic contract binding to access the raw methods on
}

// BlockHashReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockHashReceiverCallerRaw struct {
	Contract *BlockHashReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// BlockHashReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockHashReceiverTransactorRaw struct {
	Contract *BlockHashReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockHashReceiver creates a new instance of BlockHashReceiver, bound to a specific deployed contract.
func NewBlockHashReceiver(address common.Address, backend bind.ContractBackend) (*BlockHashReceiver, error) {
	contract, err := bindBlockHashReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiver{BlockHashReceiverCaller: BlockHashReceiverCaller{contract: contract}, BlockHashReceiverTransactor: BlockHashReceiverTransactor{contract: contract}, BlockHashReceiverFilterer: BlockHashReceiverFilterer{contract: contract}}, nil
}

// NewBlockHashReceiverCaller creates a new read-only instance of BlockHashReceiver, bound to a specific deployed contract.
func NewBlockHashReceiverCaller(address common.Address, caller bind.ContractCaller) (*BlockHashReceiverCaller, error) {
	contract, err := bindBlockHashReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverCaller{contract: contract}, nil
}

// NewBlockHashReceiverTransactor creates a new write-only instance of BlockHashReceiver, bound to a specific deployed contract.
func NewBlockHashReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockHashReceiverTransactor, error) {
	contract, err := bindBlockHashReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverTransactor{contract: contract}, nil
}

// NewBlockHashReceiverFilterer creates a new log filterer instance of BlockHashReceiver, bound to a specific deployed contract.
func NewBlockHashReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockHashReceiverFilterer, error) {
	contract, err := bindBlockHashReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverFilterer{contract: contract}, nil
}

// bindBlockHashReceiver binds a generic wrapper to an already deployed contract.
func bindBlockHashReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockHashReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockHashReceiver *BlockHashReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockHashReceiver.Contract.BlockHashReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockHashReceiver *BlockHashReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.BlockHashReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockHashReceiver *BlockHashReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.BlockHashReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockHashReceiver *BlockHashReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockHashReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockHashReceiver *BlockHashReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockHashReceiver *BlockHashReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.contract.Transact(opts, method, params...)
}

// GetLatestBlockInfo is a free data retrieval call binding the contract method 0xb17810be.
//
// Solidity: function getLatestBlockInfo() view returns(uint256 height, bytes32 hash)
func (_BlockHashReceiver *BlockHashReceiverCaller) GetLatestBlockInfo(opts *bind.CallOpts) (struct {
	Height *big.Int
	Hash   [32]byte
}, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "getLatestBlockInfo")

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
func (_BlockHashReceiver *BlockHashReceiverSession) GetLatestBlockInfo() (struct {
	Height *big.Int
	Hash   [32]byte
}, error) {
	return _BlockHashReceiver.Contract.GetLatestBlockInfo(&_BlockHashReceiver.CallOpts)
}

// GetLatestBlockInfo is a free data retrieval call binding the contract method 0xb17810be.
//
// Solidity: function getLatestBlockInfo() view returns(uint256 height, bytes32 hash)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) GetLatestBlockInfo() (struct {
	Height *big.Int
	Hash   [32]byte
}, error) {
	return _BlockHashReceiver.Contract.GetLatestBlockInfo(&_BlockHashReceiver.CallOpts)
}

// LatestBlockHash is a free data retrieval call binding the contract method 0x6c4f6ba9.
//
// Solidity: function latestBlockHash() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverCaller) LatestBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "latestBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestBlockHash is a free data retrieval call binding the contract method 0x6c4f6ba9.
//
// Solidity: function latestBlockHash() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverSession) LatestBlockHash() ([32]byte, error) {
	return _BlockHashReceiver.Contract.LatestBlockHash(&_BlockHashReceiver.CallOpts)
}

// LatestBlockHash is a free data retrieval call binding the contract method 0x6c4f6ba9.
//
// Solidity: function latestBlockHash() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) LatestBlockHash() ([32]byte, error) {
	return _BlockHashReceiver.Contract.LatestBlockHash(&_BlockHashReceiver.CallOpts)
}

// LatestBlockHeight is a free data retrieval call binding the contract method 0xf3f39ee5.
//
// Solidity: function latestBlockHeight() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverCaller) LatestBlockHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "latestBlockHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockHeight is a free data retrieval call binding the contract method 0xf3f39ee5.
//
// Solidity: function latestBlockHeight() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverSession) LatestBlockHeight() (*big.Int, error) {
	return _BlockHashReceiver.Contract.LatestBlockHeight(&_BlockHashReceiver.CallOpts)
}

// LatestBlockHeight is a free data retrieval call binding the contract method 0xf3f39ee5.
//
// Solidity: function latestBlockHeight() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) LatestBlockHeight() (*big.Int, error) {
	return _BlockHashReceiver.Contract.LatestBlockHeight(&_BlockHashReceiver.CallOpts)
}

// MinTeleporterVersion is a free data retrieval call binding the contract method 0xe49cc553.
//
// Solidity: function minTeleporterVersion() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverCaller) MinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "minTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTeleporterVersion is a free data retrieval call binding the contract method 0xe49cc553.
//
// Solidity: function minTeleporterVersion() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverSession) MinTeleporterVersion() (*big.Int, error) {
	return _BlockHashReceiver.Contract.MinTeleporterVersion(&_BlockHashReceiver.CallOpts)
}

// MinTeleporterVersion is a free data retrieval call binding the contract method 0xe49cc553.
//
// Solidity: function minTeleporterVersion() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) MinTeleporterVersion() (*big.Int, error) {
	return _BlockHashReceiver.Contract.MinTeleporterVersion(&_BlockHashReceiver.CallOpts)
}

// SourceChainID is a free data retrieval call binding the contract method 0x4c335368.
//
// Solidity: function sourceChainID() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverCaller) SourceChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "sourceChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SourceChainID is a free data retrieval call binding the contract method 0x4c335368.
//
// Solidity: function sourceChainID() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverSession) SourceChainID() ([32]byte, error) {
	return _BlockHashReceiver.Contract.SourceChainID(&_BlockHashReceiver.CallOpts)
}

// SourceChainID is a free data retrieval call binding the contract method 0x4c335368.
//
// Solidity: function sourceChainID() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) SourceChainID() ([32]byte, error) {
	return _BlockHashReceiver.Contract.SourceChainID(&_BlockHashReceiver.CallOpts)
}

// SourcePublisherContractAddress is a free data retrieval call binding the contract method 0x79a0710c.
//
// Solidity: function sourcePublisherContractAddress() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverCaller) SourcePublisherContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "sourcePublisherContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SourcePublisherContractAddress is a free data retrieval call binding the contract method 0x79a0710c.
//
// Solidity: function sourcePublisherContractAddress() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverSession) SourcePublisherContractAddress() (common.Address, error) {
	return _BlockHashReceiver.Contract.SourcePublisherContractAddress(&_BlockHashReceiver.CallOpts)
}

// SourcePublisherContractAddress is a free data retrieval call binding the contract method 0x79a0710c.
//
// Solidity: function sourcePublisherContractAddress() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) SourcePublisherContractAddress() (common.Address, error) {
	return _BlockHashReceiver.Contract.SourcePublisherContractAddress(&_BlockHashReceiver.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverSession) TeleporterRegistry() (common.Address, error) {
	return _BlockHashReceiver.Contract.TeleporterRegistry(&_BlockHashReceiver.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) TeleporterRegistry() (common.Address, error) {
	return _BlockHashReceiver.Contract.TeleporterRegistry(&_BlockHashReceiver.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _BlockHashReceiver.contract.Transact(opts, "receiveTeleporterMessage", originChainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_BlockHashReceiver *BlockHashReceiverSession) ReceiveTeleporterMessage(originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.ReceiveTeleporterMessage(&_BlockHashReceiver.TransactOpts, originChainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactorSession) ReceiveTeleporterMessage(originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.ReceiveTeleporterMessage(&_BlockHashReceiver.TransactOpts, originChainID, originSenderAddress, message)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0xb6109d9d.
//
// Solidity: function updateMinTeleporterVersion() returns()
func (_BlockHashReceiver *BlockHashReceiverTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashReceiver.contract.Transact(opts, "updateMinTeleporterVersion")
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0xb6109d9d.
//
// Solidity: function updateMinTeleporterVersion() returns()
func (_BlockHashReceiver *BlockHashReceiverSession) UpdateMinTeleporterVersion() (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.UpdateMinTeleporterVersion(&_BlockHashReceiver.TransactOpts)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0xb6109d9d.
//
// Solidity: function updateMinTeleporterVersion() returns()
func (_BlockHashReceiver *BlockHashReceiverTransactorSession) UpdateMinTeleporterVersion() (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.UpdateMinTeleporterVersion(&_BlockHashReceiver.TransactOpts)
}

// BlockHashReceiverMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the BlockHashReceiver contract.
type BlockHashReceiverMinTeleporterVersionUpdatedIterator struct {
	Event *BlockHashReceiverMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *BlockHashReceiverMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockHashReceiverMinTeleporterVersionUpdated)
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
		it.Event = new(BlockHashReceiverMinTeleporterVersionUpdated)
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
func (it *BlockHashReceiverMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockHashReceiverMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockHashReceiverMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the BlockHashReceiver contract.
type BlockHashReceiverMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_BlockHashReceiver *BlockHashReceiverFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*BlockHashReceiverMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverMinTeleporterVersionUpdatedIterator{contract: _BlockHashReceiver.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_BlockHashReceiver *BlockHashReceiverFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *BlockHashReceiverMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockHashReceiverMinTeleporterVersionUpdated)
				if err := _BlockHashReceiver.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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

// ParseMinTeleporterVersionUpdated is a log parse operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_BlockHashReceiver *BlockHashReceiverFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*BlockHashReceiverMinTeleporterVersionUpdated, error) {
	event := new(BlockHashReceiverMinTeleporterVersionUpdated)
	if err := _BlockHashReceiver.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockHashReceiverReceiveBlockHashIterator is returned from FilterReceiveBlockHash and is used to iterate over the raw logs and unpacked data for ReceiveBlockHash events raised by the BlockHashReceiver contract.
type BlockHashReceiverReceiveBlockHashIterator struct {
	Event *BlockHashReceiverReceiveBlockHash // Event containing the contract specifics and raw log

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
func (it *BlockHashReceiverReceiveBlockHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockHashReceiverReceiveBlockHash)
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
		it.Event = new(BlockHashReceiverReceiveBlockHash)
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
func (it *BlockHashReceiverReceiveBlockHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockHashReceiverReceiveBlockHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockHashReceiverReceiveBlockHash represents a ReceiveBlockHash event raised by the BlockHashReceiver contract.
type BlockHashReceiverReceiveBlockHash struct {
	OriginChainID       [32]byte
	OriginSenderAddress common.Address
	BlockHeight         *big.Int
	BlockHash           [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReceiveBlockHash is a free log retrieval operation binding the contract event 0x0bca78aa82d7575f42e4b4b2fe04765a4b2f3661786403788ce987e065ac590d.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, address indexed originSenderAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_BlockHashReceiver *BlockHashReceiverFilterer) FilterReceiveBlockHash(opts *bind.FilterOpts, originChainID [][32]byte, originSenderAddress []common.Address, blockHeight []*big.Int) (*BlockHashReceiverReceiveBlockHashIterator, error) {

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

	logs, sub, err := _BlockHashReceiver.contract.FilterLogs(opts, "ReceiveBlockHash", originChainIDRule, originSenderAddressRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverReceiveBlockHashIterator{contract: _BlockHashReceiver.contract, event: "ReceiveBlockHash", logs: logs, sub: sub}, nil
}

// WatchReceiveBlockHash is a free log subscription operation binding the contract event 0x0bca78aa82d7575f42e4b4b2fe04765a4b2f3661786403788ce987e065ac590d.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, address indexed originSenderAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_BlockHashReceiver *BlockHashReceiverFilterer) WatchReceiveBlockHash(opts *bind.WatchOpts, sink chan<- *BlockHashReceiverReceiveBlockHash, originChainID [][32]byte, originSenderAddress []common.Address, blockHeight []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _BlockHashReceiver.contract.WatchLogs(opts, "ReceiveBlockHash", originChainIDRule, originSenderAddressRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockHashReceiverReceiveBlockHash)
				if err := _BlockHashReceiver.contract.UnpackLog(event, "ReceiveBlockHash", log); err != nil {
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
func (_BlockHashReceiver *BlockHashReceiverFilterer) ParseReceiveBlockHash(log types.Log) (*BlockHashReceiverReceiveBlockHash, error) {
	event := new(BlockHashReceiverReceiveBlockHash)
	if err := _BlockHashReceiver.contract.UnpackLog(event, "ReceiveBlockHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
