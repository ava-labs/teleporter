// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokendestination

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

// NativetokendestinationMetaData contains all meta data concerning the Nativetokendestination contract.
var NativetokendestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sourceContractAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenReserve_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceNotIncreased\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MintNativeTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"}],\"name\":\"TransferToSource\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TRANSFER_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"senderBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"transferToSource\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NativetokendestinationABI is the input ABI used to generate the binding from.
// Deprecated: Use NativetokendestinationMetaData.ABI instead.
var NativetokendestinationABI = NativetokendestinationMetaData.ABI

// Nativetokendestination is an auto generated Go binding around an Ethereum contract.
type Nativetokendestination struct {
	NativetokendestinationCaller     // Read-only binding to the contract
	NativetokendestinationTransactor // Write-only binding to the contract
	NativetokendestinationFilterer   // Log filterer for contract events
}

// NativetokendestinationCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativetokendestinationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativetokendestinationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativetokendestinationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativetokendestinationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativetokendestinationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativetokendestinationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativetokendestinationSession struct {
	Contract     *Nativetokendestination // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NativetokendestinationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativetokendestinationCallerSession struct {
	Contract *NativetokendestinationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// NativetokendestinationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativetokendestinationTransactorSession struct {
	Contract     *NativetokendestinationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// NativetokendestinationRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativetokendestinationRaw struct {
	Contract *Nativetokendestination // Generic contract binding to access the raw methods on
}

// NativetokendestinationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativetokendestinationCallerRaw struct {
	Contract *NativetokendestinationCaller // Generic read-only contract binding to access the raw methods on
}

// NativetokendestinationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativetokendestinationTransactorRaw struct {
	Contract *NativetokendestinationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativetokendestination creates a new instance of Nativetokendestination, bound to a specific deployed contract.
func NewNativetokendestination(address common.Address, backend bind.ContractBackend) (*Nativetokendestination, error) {
	contract, err := bindNativetokendestination(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nativetokendestination{NativetokendestinationCaller: NativetokendestinationCaller{contract: contract}, NativetokendestinationTransactor: NativetokendestinationTransactor{contract: contract}, NativetokendestinationFilterer: NativetokendestinationFilterer{contract: contract}}, nil
}

// NewNativetokendestinationCaller creates a new read-only instance of Nativetokendestination, bound to a specific deployed contract.
func NewNativetokendestinationCaller(address common.Address, caller bind.ContractCaller) (*NativetokendestinationCaller, error) {
	contract, err := bindNativetokendestination(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativetokendestinationCaller{contract: contract}, nil
}

// NewNativetokendestinationTransactor creates a new write-only instance of Nativetokendestination, bound to a specific deployed contract.
func NewNativetokendestinationTransactor(address common.Address, transactor bind.ContractTransactor) (*NativetokendestinationTransactor, error) {
	contract, err := bindNativetokendestination(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativetokendestinationTransactor{contract: contract}, nil
}

// NewNativetokendestinationFilterer creates a new log filterer instance of Nativetokendestination, bound to a specific deployed contract.
func NewNativetokendestinationFilterer(address common.Address, filterer bind.ContractFilterer) (*NativetokendestinationFilterer, error) {
	contract, err := bindNativetokendestination(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativetokendestinationFilterer{contract: contract}, nil
}

// bindNativetokendestination binds a generic wrapper to an already deployed contract.
func bindNativetokendestination(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativetokendestinationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nativetokendestination *NativetokendestinationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nativetokendestination.Contract.NativetokendestinationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nativetokendestination *NativetokendestinationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nativetokendestination.Contract.NativetokendestinationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nativetokendestination *NativetokendestinationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nativetokendestination.Contract.NativetokendestinationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nativetokendestination *NativetokendestinationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nativetokendestination.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nativetokendestination *NativetokendestinationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nativetokendestination.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nativetokendestination *NativetokendestinationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nativetokendestination.Contract.contract.Transact(opts, method, params...)
}

// TRANSFERNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xab285230.
//
// Solidity: function TRANSFER_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Nativetokendestination *NativetokendestinationCaller) TRANSFERNATIVETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nativetokendestination.contract.Call(opts, &out, "TRANSFER_NATIVE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TRANSFERNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xab285230.
//
// Solidity: function TRANSFER_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Nativetokendestination *NativetokendestinationSession) TRANSFERNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _Nativetokendestination.Contract.TRANSFERNATIVETOKENSREQUIREDGAS(&_Nativetokendestination.CallOpts)
}

// TRANSFERNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xab285230.
//
// Solidity: function TRANSFER_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Nativetokendestination *NativetokendestinationCallerSession) TRANSFERNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _Nativetokendestination.Contract.TRANSFERNATIVETOKENSREQUIREDGAS(&_Nativetokendestination.CallOpts)
}

// CurrentBlockchainID is a free data retrieval call binding the contract method 0x4950d2d0.
//
// Solidity: function currentBlockchainID() view returns(bytes32)
func (_Nativetokendestination *NativetokendestinationCaller) CurrentBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nativetokendestination.contract.Call(opts, &out, "currentBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentBlockchainID is a free data retrieval call binding the contract method 0x4950d2d0.
//
// Solidity: function currentBlockchainID() view returns(bytes32)
func (_Nativetokendestination *NativetokendestinationSession) CurrentBlockchainID() ([32]byte, error) {
	return _Nativetokendestination.Contract.CurrentBlockchainID(&_Nativetokendestination.CallOpts)
}

// CurrentBlockchainID is a free data retrieval call binding the contract method 0x4950d2d0.
//
// Solidity: function currentBlockchainID() view returns(bytes32)
func (_Nativetokendestination *NativetokendestinationCallerSession) CurrentBlockchainID() ([32]byte, error) {
	return _Nativetokendestination.Contract.CurrentBlockchainID(&_Nativetokendestination.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_Nativetokendestination *NativetokendestinationCaller) SourceBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nativetokendestination.contract.Call(opts, &out, "sourceBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_Nativetokendestination *NativetokendestinationSession) SourceBlockchainID() ([32]byte, error) {
	return _Nativetokendestination.Contract.SourceBlockchainID(&_Nativetokendestination.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_Nativetokendestination *NativetokendestinationCallerSession) SourceBlockchainID() ([32]byte, error) {
	return _Nativetokendestination.Contract.SourceBlockchainID(&_Nativetokendestination.CallOpts)
}

// SourceContractAddress is a free data retrieval call binding the contract method 0x3a85f16a.
//
// Solidity: function sourceContractAddress() view returns(address)
func (_Nativetokendestination *NativetokendestinationCaller) SourceContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nativetokendestination.contract.Call(opts, &out, "sourceContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SourceContractAddress is a free data retrieval call binding the contract method 0x3a85f16a.
//
// Solidity: function sourceContractAddress() view returns(address)
func (_Nativetokendestination *NativetokendestinationSession) SourceContractAddress() (common.Address, error) {
	return _Nativetokendestination.Contract.SourceContractAddress(&_Nativetokendestination.CallOpts)
}

// SourceContractAddress is a free data retrieval call binding the contract method 0x3a85f16a.
//
// Solidity: function sourceContractAddress() view returns(address)
func (_Nativetokendestination *NativetokendestinationCallerSession) SourceContractAddress() (common.Address, error) {
	return _Nativetokendestination.Contract.SourceContractAddress(&_Nativetokendestination.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Nativetokendestination *NativetokendestinationCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nativetokendestination.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Nativetokendestination *NativetokendestinationSession) TeleporterMessenger() (common.Address, error) {
	return _Nativetokendestination.Contract.TeleporterMessenger(&_Nativetokendestination.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Nativetokendestination *NativetokendestinationCallerSession) TeleporterMessenger() (common.Address, error) {
	return _Nativetokendestination.Contract.TeleporterMessenger(&_Nativetokendestination.CallOpts)
}

// TokenReserve is a free data retrieval call binding the contract method 0xcbcb3171.
//
// Solidity: function tokenReserve() view returns(uint256)
func (_Nativetokendestination *NativetokendestinationCaller) TokenReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nativetokendestination.contract.Call(opts, &out, "tokenReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenReserve is a free data retrieval call binding the contract method 0xcbcb3171.
//
// Solidity: function tokenReserve() view returns(uint256)
func (_Nativetokendestination *NativetokendestinationSession) TokenReserve() (*big.Int, error) {
	return _Nativetokendestination.Contract.TokenReserve(&_Nativetokendestination.CallOpts)
}

// TokenReserve is a free data retrieval call binding the contract method 0xcbcb3171.
//
// Solidity: function tokenReserve() view returns(uint256)
func (_Nativetokendestination *NativetokendestinationCallerSession) TokenReserve() (*big.Int, error) {
	return _Nativetokendestination.Contract.TokenReserve(&_Nativetokendestination.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_Nativetokendestination *NativetokendestinationTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Nativetokendestination.contract.Transact(opts, "receiveTeleporterMessage", senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_Nativetokendestination *NativetokendestinationSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Nativetokendestination.Contract.ReceiveTeleporterMessage(&_Nativetokendestination.TransactOpts, senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_Nativetokendestination *NativetokendestinationTransactorSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Nativetokendestination.Contract.ReceiveTeleporterMessage(&_Nativetokendestination.TransactOpts, senderBlockchainID, senderAddress, message)
}

// TransferToSource is a paid mutator transaction binding the contract method 0x281c91de.
//
// Solidity: function transferToSource(address recipient, address feeContractAddress, uint256 feeAmount) payable returns()
func (_Nativetokendestination *NativetokendestinationTransactor) TransferToSource(opts *bind.TransactOpts, recipient common.Address, feeContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _Nativetokendestination.contract.Transact(opts, "transferToSource", recipient, feeContractAddress, feeAmount)
}

// TransferToSource is a paid mutator transaction binding the contract method 0x281c91de.
//
// Solidity: function transferToSource(address recipient, address feeContractAddress, uint256 feeAmount) payable returns()
func (_Nativetokendestination *NativetokendestinationSession) TransferToSource(recipient common.Address, feeContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _Nativetokendestination.Contract.TransferToSource(&_Nativetokendestination.TransactOpts, recipient, feeContractAddress, feeAmount)
}

// TransferToSource is a paid mutator transaction binding the contract method 0x281c91de.
//
// Solidity: function transferToSource(address recipient, address feeContractAddress, uint256 feeAmount) payable returns()
func (_Nativetokendestination *NativetokendestinationTransactorSession) TransferToSource(recipient common.Address, feeContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _Nativetokendestination.Contract.TransferToSource(&_Nativetokendestination.TransactOpts, recipient, feeContractAddress, feeAmount)
}

// NativetokendestinationCollateralAddedIterator is returned from FilterCollateralAdded and is used to iterate over the raw logs and unpacked data for CollateralAdded events raised by the Nativetokendestination contract.
type NativetokendestinationCollateralAddedIterator struct {
	Event *NativetokendestinationCollateralAdded // Event containing the contract specifics and raw log

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
func (it *NativetokendestinationCollateralAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativetokendestinationCollateralAdded)
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
		it.Event = new(NativetokendestinationCollateralAdded)
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
func (it *NativetokendestinationCollateralAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativetokendestinationCollateralAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativetokendestinationCollateralAdded represents a CollateralAdded event raised by the Nativetokendestination contract.
type NativetokendestinationCollateralAdded struct {
	Amount    *big.Int
	Remaining *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdded is a free log retrieval operation binding the contract event 0x244160b15e69cc411f041d94ae7fab6f6bba85dade8403216c05ff4b920d5449.
//
// Solidity: event CollateralAdded(uint256 amount, uint256 remaining)
func (_Nativetokendestination *NativetokendestinationFilterer) FilterCollateralAdded(opts *bind.FilterOpts) (*NativetokendestinationCollateralAddedIterator, error) {

	logs, sub, err := _Nativetokendestination.contract.FilterLogs(opts, "CollateralAdded")
	if err != nil {
		return nil, err
	}
	return &NativetokendestinationCollateralAddedIterator{contract: _Nativetokendestination.contract, event: "CollateralAdded", logs: logs, sub: sub}, nil
}

// WatchCollateralAdded is a free log subscription operation binding the contract event 0x244160b15e69cc411f041d94ae7fab6f6bba85dade8403216c05ff4b920d5449.
//
// Solidity: event CollateralAdded(uint256 amount, uint256 remaining)
func (_Nativetokendestination *NativetokendestinationFilterer) WatchCollateralAdded(opts *bind.WatchOpts, sink chan<- *NativetokendestinationCollateralAdded) (event.Subscription, error) {

	logs, sub, err := _Nativetokendestination.contract.WatchLogs(opts, "CollateralAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativetokendestinationCollateralAdded)
				if err := _Nativetokendestination.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
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

// ParseCollateralAdded is a log parse operation binding the contract event 0x244160b15e69cc411f041d94ae7fab6f6bba85dade8403216c05ff4b920d5449.
//
// Solidity: event CollateralAdded(uint256 amount, uint256 remaining)
func (_Nativetokendestination *NativetokendestinationFilterer) ParseCollateralAdded(log types.Log) (*NativetokendestinationCollateralAdded, error) {
	event := new(NativetokendestinationCollateralAdded)
	if err := _Nativetokendestination.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativetokendestinationMintNativeTokensIterator is returned from FilterMintNativeTokens and is used to iterate over the raw logs and unpacked data for MintNativeTokens events raised by the Nativetokendestination contract.
type NativetokendestinationMintNativeTokensIterator struct {
	Event *NativetokendestinationMintNativeTokens // Event containing the contract specifics and raw log

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
func (it *NativetokendestinationMintNativeTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativetokendestinationMintNativeTokens)
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
		it.Event = new(NativetokendestinationMintNativeTokens)
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
func (it *NativetokendestinationMintNativeTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativetokendestinationMintNativeTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativetokendestinationMintNativeTokens represents a MintNativeTokens event raised by the Nativetokendestination contract.
type NativetokendestinationMintNativeTokens struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintNativeTokens is a free log retrieval operation binding the contract event 0xe2899f7ef7618206fa13efbcde0c05cccd83f35ac3b18ef860021181d61fa680.
//
// Solidity: event MintNativeTokens(address recipient, uint256 amount)
func (_Nativetokendestination *NativetokendestinationFilterer) FilterMintNativeTokens(opts *bind.FilterOpts) (*NativetokendestinationMintNativeTokensIterator, error) {

	logs, sub, err := _Nativetokendestination.contract.FilterLogs(opts, "MintNativeTokens")
	if err != nil {
		return nil, err
	}
	return &NativetokendestinationMintNativeTokensIterator{contract: _Nativetokendestination.contract, event: "MintNativeTokens", logs: logs, sub: sub}, nil
}

// WatchMintNativeTokens is a free log subscription operation binding the contract event 0xe2899f7ef7618206fa13efbcde0c05cccd83f35ac3b18ef860021181d61fa680.
//
// Solidity: event MintNativeTokens(address recipient, uint256 amount)
func (_Nativetokendestination *NativetokendestinationFilterer) WatchMintNativeTokens(opts *bind.WatchOpts, sink chan<- *NativetokendestinationMintNativeTokens) (event.Subscription, error) {

	logs, sub, err := _Nativetokendestination.contract.WatchLogs(opts, "MintNativeTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativetokendestinationMintNativeTokens)
				if err := _Nativetokendestination.contract.UnpackLog(event, "MintNativeTokens", log); err != nil {
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

// ParseMintNativeTokens is a log parse operation binding the contract event 0xe2899f7ef7618206fa13efbcde0c05cccd83f35ac3b18ef860021181d61fa680.
//
// Solidity: event MintNativeTokens(address recipient, uint256 amount)
func (_Nativetokendestination *NativetokendestinationFilterer) ParseMintNativeTokens(log types.Log) (*NativetokendestinationMintNativeTokens, error) {
	event := new(NativetokendestinationMintNativeTokens)
	if err := _Nativetokendestination.contract.UnpackLog(event, "MintNativeTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativetokendestinationTransferToSourceIterator is returned from FilterTransferToSource and is used to iterate over the raw logs and unpacked data for TransferToSource events raised by the Nativetokendestination contract.
type NativetokendestinationTransferToSourceIterator struct {
	Event *NativetokendestinationTransferToSource // Event containing the contract specifics and raw log

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
func (it *NativetokendestinationTransferToSourceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativetokendestinationTransferToSource)
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
		it.Event = new(NativetokendestinationTransferToSource)
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
func (it *NativetokendestinationTransferToSourceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativetokendestinationTransferToSourceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativetokendestinationTransferToSource represents a TransferToSource event raised by the Nativetokendestination contract.
type NativetokendestinationTransferToSource struct {
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	FeeContractAddress  common.Address
	FeeAmount           *big.Int
	TeleporterMessageID *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTransferToSource is a free log retrieval operation binding the contract event 0x343e0fe771af9fa258c07a602fad6dfd80c51eda78604397e553106360933bfc.
//
// Solidity: event TransferToSource(address indexed sender, address indexed recipient, uint256 amount, address feeContractAddress, uint256 feeAmount, uint256 teleporterMessageID)
func (_Nativetokendestination *NativetokendestinationFilterer) FilterTransferToSource(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*NativetokendestinationTransferToSourceIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Nativetokendestination.contract.FilterLogs(opts, "TransferToSource", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativetokendestinationTransferToSourceIterator{contract: _Nativetokendestination.contract, event: "TransferToSource", logs: logs, sub: sub}, nil
}

// WatchTransferToSource is a free log subscription operation binding the contract event 0x343e0fe771af9fa258c07a602fad6dfd80c51eda78604397e553106360933bfc.
//
// Solidity: event TransferToSource(address indexed sender, address indexed recipient, uint256 amount, address feeContractAddress, uint256 feeAmount, uint256 teleporterMessageID)
func (_Nativetokendestination *NativetokendestinationFilterer) WatchTransferToSource(opts *bind.WatchOpts, sink chan<- *NativetokendestinationTransferToSource, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Nativetokendestination.contract.WatchLogs(opts, "TransferToSource", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativetokendestinationTransferToSource)
				if err := _Nativetokendestination.contract.UnpackLog(event, "TransferToSource", log); err != nil {
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

// ParseTransferToSource is a log parse operation binding the contract event 0x343e0fe771af9fa258c07a602fad6dfd80c51eda78604397e553106360933bfc.
//
// Solidity: event TransferToSource(address indexed sender, address indexed recipient, uint256 amount, address feeContractAddress, uint256 feeAmount, uint256 teleporterMessageID)
func (_Nativetokendestination *NativetokendestinationFilterer) ParseTransferToSource(log types.Log) (*NativetokendestinationTransferToSource, error) {
	event := new(NativetokendestinationTransferToSource)
	if err := _Nativetokendestination.contract.UnpackLog(event, "TransferToSource", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
