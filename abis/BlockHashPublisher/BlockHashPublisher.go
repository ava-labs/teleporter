// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockhashpublisher

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

// BlockhashpublisherMetaData contains all meta data concerning the Blockhashpublisher contract.
var BlockhashpublisherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"PublishBlockHash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"}],\"name\":\"publishLatestBlockHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BlockhashpublisherABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockhashpublisherMetaData.ABI instead.
var BlockhashpublisherABI = BlockhashpublisherMetaData.ABI

// Blockhashpublisher is an auto generated Go binding around an Ethereum contract.
type Blockhashpublisher struct {
	BlockhashpublisherCaller     // Read-only binding to the contract
	BlockhashpublisherTransactor // Write-only binding to the contract
	BlockhashpublisherFilterer   // Log filterer for contract events
}

// BlockhashpublisherCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockhashpublisherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockhashpublisherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockhashpublisherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockhashpublisherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockhashpublisherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockhashpublisherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockhashpublisherSession struct {
	Contract     *Blockhashpublisher // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BlockhashpublisherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockhashpublisherCallerSession struct {
	Contract *BlockhashpublisherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BlockhashpublisherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockhashpublisherTransactorSession struct {
	Contract     *BlockhashpublisherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BlockhashpublisherRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockhashpublisherRaw struct {
	Contract *Blockhashpublisher // Generic contract binding to access the raw methods on
}

// BlockhashpublisherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockhashpublisherCallerRaw struct {
	Contract *BlockhashpublisherCaller // Generic read-only contract binding to access the raw methods on
}

// BlockhashpublisherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockhashpublisherTransactorRaw struct {
	Contract *BlockhashpublisherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockhashpublisher creates a new instance of Blockhashpublisher, bound to a specific deployed contract.
func NewBlockhashpublisher(address common.Address, backend bind.ContractBackend) (*Blockhashpublisher, error) {
	contract, err := bindBlockhashpublisher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockhashpublisher{BlockhashpublisherCaller: BlockhashpublisherCaller{contract: contract}, BlockhashpublisherTransactor: BlockhashpublisherTransactor{contract: contract}, BlockhashpublisherFilterer: BlockhashpublisherFilterer{contract: contract}}, nil
}

// NewBlockhashpublisherCaller creates a new read-only instance of Blockhashpublisher, bound to a specific deployed contract.
func NewBlockhashpublisherCaller(address common.Address, caller bind.ContractCaller) (*BlockhashpublisherCaller, error) {
	contract, err := bindBlockhashpublisher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockhashpublisherCaller{contract: contract}, nil
}

// NewBlockhashpublisherTransactor creates a new write-only instance of Blockhashpublisher, bound to a specific deployed contract.
func NewBlockhashpublisherTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockhashpublisherTransactor, error) {
	contract, err := bindBlockhashpublisher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockhashpublisherTransactor{contract: contract}, nil
}

// NewBlockhashpublisherFilterer creates a new log filterer instance of Blockhashpublisher, bound to a specific deployed contract.
func NewBlockhashpublisherFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockhashpublisherFilterer, error) {
	contract, err := bindBlockhashpublisher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockhashpublisherFilterer{contract: contract}, nil
}

// bindBlockhashpublisher binds a generic wrapper to an already deployed contract.
func bindBlockhashpublisher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockhashpublisherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockhashpublisher *BlockhashpublisherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockhashpublisher.Contract.BlockhashpublisherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockhashpublisher *BlockhashpublisherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockhashpublisher.Contract.BlockhashpublisherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockhashpublisher *BlockhashpublisherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockhashpublisher.Contract.BlockhashpublisherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockhashpublisher *BlockhashpublisherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockhashpublisher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockhashpublisher *BlockhashpublisherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockhashpublisher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockhashpublisher *BlockhashpublisherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockhashpublisher.Contract.contract.Transact(opts, method, params...)
}

// RECEIVEBLOCKHASHREQUIREDGASLIMIT is a free data retrieval call binding the contract method 0x18aef19b.
//
// Solidity: function RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT() view returns(uint256)
func (_Blockhashpublisher *BlockhashpublisherCaller) RECEIVEBLOCKHASHREQUIREDGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blockhashpublisher.contract.Call(opts, &out, "RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RECEIVEBLOCKHASHREQUIREDGASLIMIT is a free data retrieval call binding the contract method 0x18aef19b.
//
// Solidity: function RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT() view returns(uint256)
func (_Blockhashpublisher *BlockhashpublisherSession) RECEIVEBLOCKHASHREQUIREDGASLIMIT() (*big.Int, error) {
	return _Blockhashpublisher.Contract.RECEIVEBLOCKHASHREQUIREDGASLIMIT(&_Blockhashpublisher.CallOpts)
}

// RECEIVEBLOCKHASHREQUIREDGASLIMIT is a free data retrieval call binding the contract method 0x18aef19b.
//
// Solidity: function RECEIVE_BLOCK_HASH_REQUIRED_GAS_LIMIT() view returns(uint256)
func (_Blockhashpublisher *BlockhashpublisherCallerSession) RECEIVEBLOCKHASHREQUIREDGASLIMIT() (*big.Int, error) {
	return _Blockhashpublisher.Contract.RECEIVEBLOCKHASHREQUIREDGASLIMIT(&_Blockhashpublisher.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Blockhashpublisher *BlockhashpublisherCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockhashpublisher.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Blockhashpublisher *BlockhashpublisherSession) TeleporterMessenger() (common.Address, error) {
	return _Blockhashpublisher.Contract.TeleporterMessenger(&_Blockhashpublisher.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Blockhashpublisher *BlockhashpublisherCallerSession) TeleporterMessenger() (common.Address, error) {
	return _Blockhashpublisher.Contract.TeleporterMessenger(&_Blockhashpublisher.CallOpts)
}

// PublishLatestBlockHash is a paid mutator transaction binding the contract method 0x82ab2b86.
//
// Solidity: function publishLatestBlockHash(bytes32 destinationChainID, address destinationAddress) returns(uint256 messageID)
func (_Blockhashpublisher *BlockhashpublisherTransactor) PublishLatestBlockHash(opts *bind.TransactOpts, destinationChainID [32]byte, destinationAddress common.Address) (*types.Transaction, error) {
	return _Blockhashpublisher.contract.Transact(opts, "publishLatestBlockHash", destinationChainID, destinationAddress)
}

// PublishLatestBlockHash is a paid mutator transaction binding the contract method 0x82ab2b86.
//
// Solidity: function publishLatestBlockHash(bytes32 destinationChainID, address destinationAddress) returns(uint256 messageID)
func (_Blockhashpublisher *BlockhashpublisherSession) PublishLatestBlockHash(destinationChainID [32]byte, destinationAddress common.Address) (*types.Transaction, error) {
	return _Blockhashpublisher.Contract.PublishLatestBlockHash(&_Blockhashpublisher.TransactOpts, destinationChainID, destinationAddress)
}

// PublishLatestBlockHash is a paid mutator transaction binding the contract method 0x82ab2b86.
//
// Solidity: function publishLatestBlockHash(bytes32 destinationChainID, address destinationAddress) returns(uint256 messageID)
func (_Blockhashpublisher *BlockhashpublisherTransactorSession) PublishLatestBlockHash(destinationChainID [32]byte, destinationAddress common.Address) (*types.Transaction, error) {
	return _Blockhashpublisher.Contract.PublishLatestBlockHash(&_Blockhashpublisher.TransactOpts, destinationChainID, destinationAddress)
}

// BlockhashpublisherPublishBlockHashIterator is returned from FilterPublishBlockHash and is used to iterate over the raw logs and unpacked data for PublishBlockHash events raised by the Blockhashpublisher contract.
type BlockhashpublisherPublishBlockHashIterator struct {
	Event *BlockhashpublisherPublishBlockHash // Event containing the contract specifics and raw log

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
func (it *BlockhashpublisherPublishBlockHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockhashpublisherPublishBlockHash)
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
		it.Event = new(BlockhashpublisherPublishBlockHash)
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
func (it *BlockhashpublisherPublishBlockHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockhashpublisherPublishBlockHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockhashpublisherPublishBlockHash represents a PublishBlockHash event raised by the Blockhashpublisher contract.
type BlockhashpublisherPublishBlockHash struct {
	DestinationChainID [32]byte
	DestinationAddress common.Address
	BlockHeight        *big.Int
	BlockHash          [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPublishBlockHash is a free log retrieval operation binding the contract event 0xe13623d33d18131ce960c33b1282ceac1fe7b5ccfcf7f8c0f6dad32dd61e3bdd.
//
// Solidity: event PublishBlockHash(bytes32 indexed destinationChainID, address indexed destinationAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_Blockhashpublisher *BlockhashpublisherFilterer) FilterPublishBlockHash(opts *bind.FilterOpts, destinationChainID [][32]byte, destinationAddress []common.Address, blockHeight []*big.Int) (*BlockhashpublisherPublishBlockHashIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _Blockhashpublisher.contract.FilterLogs(opts, "PublishBlockHash", destinationChainIDRule, destinationAddressRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return &BlockhashpublisherPublishBlockHashIterator{contract: _Blockhashpublisher.contract, event: "PublishBlockHash", logs: logs, sub: sub}, nil
}

// WatchPublishBlockHash is a free log subscription operation binding the contract event 0xe13623d33d18131ce960c33b1282ceac1fe7b5ccfcf7f8c0f6dad32dd61e3bdd.
//
// Solidity: event PublishBlockHash(bytes32 indexed destinationChainID, address indexed destinationAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_Blockhashpublisher *BlockhashpublisherFilterer) WatchPublishBlockHash(opts *bind.WatchOpts, sink chan<- *BlockhashpublisherPublishBlockHash, destinationChainID [][32]byte, destinationAddress []common.Address, blockHeight []*big.Int) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _Blockhashpublisher.contract.WatchLogs(opts, "PublishBlockHash", destinationChainIDRule, destinationAddressRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockhashpublisherPublishBlockHash)
				if err := _Blockhashpublisher.contract.UnpackLog(event, "PublishBlockHash", log); err != nil {
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

// ParsePublishBlockHash is a log parse operation binding the contract event 0xe13623d33d18131ce960c33b1282ceac1fe7b5ccfcf7f8c0f6dad32dd61e3bdd.
//
// Solidity: event PublishBlockHash(bytes32 indexed destinationChainID, address indexed destinationAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_Blockhashpublisher *BlockhashpublisherFilterer) ParsePublishBlockHash(log types.Log) (*BlockhashpublisherPublishBlockHash, error) {
	event := new(BlockhashpublisherPublishBlockHash)
	if err := _Blockhashpublisher.contract.UnpackLog(event, "PublishBlockHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
