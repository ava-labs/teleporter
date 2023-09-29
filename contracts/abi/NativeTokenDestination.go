// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// NativeTokenDestinationMetaData contains all meta data concerning the NativeTokenDestination contract.
var NativeTokenDestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceContractAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceChainID_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceNotIncreased\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotBridgeTokenWithinSameChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"adjustedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientAdjustedAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPartnerContractAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSourceContractAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeleporterMessengerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MintNativeTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"TransferToSource\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TRANSFER_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_PRECOMPILE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nativeChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"transferToSource\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NativeTokenDestinationABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenDestinationMetaData.ABI instead.
var NativeTokenDestinationABI = NativeTokenDestinationMetaData.ABI

// NativeTokenDestination is an auto generated Go binding around an Ethereum contract.
type NativeTokenDestination struct {
	NativeTokenDestinationCaller     // Read-only binding to the contract
	NativeTokenDestinationTransactor // Write-only binding to the contract
	NativeTokenDestinationFilterer   // Log filterer for contract events
}

// NativeTokenDestinationCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenDestinationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenDestinationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenDestinationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenDestinationSession struct {
	Contract     *NativeTokenDestination // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NativeTokenDestinationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenDestinationCallerSession struct {
	Contract *NativeTokenDestinationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// NativeTokenDestinationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenDestinationTransactorSession struct {
	Contract     *NativeTokenDestinationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// NativeTokenDestinationRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenDestinationRaw struct {
	Contract *NativeTokenDestination // Generic contract binding to access the raw methods on
}

// NativeTokenDestinationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenDestinationCallerRaw struct {
	Contract *NativeTokenDestinationCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenDestinationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenDestinationTransactorRaw struct {
	Contract *NativeTokenDestinationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenDestination creates a new instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestination(address common.Address, backend bind.ContractBackend) (*NativeTokenDestination, error) {
	contract, err := bindNativeTokenDestination(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestination{NativeTokenDestinationCaller: NativeTokenDestinationCaller{contract: contract}, NativeTokenDestinationTransactor: NativeTokenDestinationTransactor{contract: contract}, NativeTokenDestinationFilterer: NativeTokenDestinationFilterer{contract: contract}}, nil
}

// NewNativeTokenDestinationCaller creates a new read-only instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenDestinationCaller, error) {
	contract, err := bindNativeTokenDestination(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCaller{contract: contract}, nil
}

// NewNativeTokenDestinationTransactor creates a new write-only instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenDestinationTransactor, error) {
	contract, err := bindNativeTokenDestination(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTransactor{contract: contract}, nil
}

// NewNativeTokenDestinationFilterer creates a new log filterer instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenDestinationFilterer, error) {
	contract, err := bindNativeTokenDestination(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationFilterer{contract: contract}, nil
}

// bindNativeTokenDestination binds a generic wrapper to an already deployed contract.
func bindNativeTokenDestination(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenDestinationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenDestination.Contract.NativeTokenDestinationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.NativeTokenDestinationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.NativeTokenDestinationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenDestination *NativeTokenDestinationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenDestination.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenDestination *NativeTokenDestinationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenDestination *NativeTokenDestinationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.contract.Transact(opts, method, params...)
}

// TRANSFERNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xab285230.
//
// Solidity: function TRANSFER_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TRANSFERNATIVETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "TRANSFER_NATIVE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TRANSFERNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xab285230.
//
// Solidity: function TRANSFER_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TRANSFERNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TRANSFERNATIVETOKENSREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// TRANSFERNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xab285230.
//
// Solidity: function TRANSFER_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TRANSFERNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TRANSFERNATIVETOKENSREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) WARPPRECOMPILEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "WARP_PRECOMPILE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) WARPPRECOMPILEADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.WARPPRECOMPILEADDRESS(&_NativeTokenDestination.CallOpts)
}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) WARPPRECOMPILEADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.WARPPRECOMPILEADDRESS(&_NativeTokenDestination.CallOpts)
}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCaller) CurrentChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "currentChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationSession) CurrentChainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.CurrentChainID(&_NativeTokenDestination.CallOpts)
}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) CurrentChainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.CurrentChainID(&_NativeTokenDestination.CallOpts)
}

// SourceChainID is a free data retrieval call binding the contract method 0x4c335368.
//
// Solidity: function sourceChainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCaller) SourceChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "sourceChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SourceChainID is a free data retrieval call binding the contract method 0x4c335368.
//
// Solidity: function sourceChainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationSession) SourceChainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.SourceChainID(&_NativeTokenDestination.CallOpts)
}

// SourceChainID is a free data retrieval call binding the contract method 0x4c335368.
//
// Solidity: function sourceChainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) SourceChainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.SourceChainID(&_NativeTokenDestination.CallOpts)
}

// SourceContractAddress is a free data retrieval call binding the contract method 0x3a85f16a.
//
// Solidity: function sourceContractAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) SourceContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "sourceContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SourceContractAddress is a free data retrieval call binding the contract method 0x3a85f16a.
//
// Solidity: function sourceContractAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) SourceContractAddress() (common.Address, error) {
	return _NativeTokenDestination.Contract.SourceContractAddress(&_NativeTokenDestination.CallOpts)
}

// SourceContractAddress is a free data retrieval call binding the contract method 0x3a85f16a.
//
// Solidity: function sourceContractAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) SourceContractAddress() (common.Address, error) {
	return _NativeTokenDestination.Contract.SourceContractAddress(&_NativeTokenDestination.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) TeleporterMessenger() (common.Address, error) {
	return _NativeTokenDestination.Contract.TeleporterMessenger(&_NativeTokenDestination.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TeleporterMessenger() (common.Address, error) {
	return _NativeTokenDestination.Contract.TeleporterMessenger(&_NativeTokenDestination.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "receiveTeleporterMessage", nativeChainID, nativeBridgeAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) ReceiveTeleporterMessage(nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReceiveTeleporterMessage(&_NativeTokenDestination.TransactOpts, nativeChainID, nativeBridgeAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) ReceiveTeleporterMessage(nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReceiveTeleporterMessage(&_NativeTokenDestination.TransactOpts, nativeChainID, nativeBridgeAddress, message)
}

// TransferToSource is a paid mutator transaction binding the contract method 0x281c91de.
//
// Solidity: function transferToSource(address recipient, address feeTokenContractAddress, uint256 feeAmount) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) TransferToSource(opts *bind.TransactOpts, recipient common.Address, feeTokenContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "transferToSource", recipient, feeTokenContractAddress, feeAmount)
}

// TransferToSource is a paid mutator transaction binding the contract method 0x281c91de.
//
// Solidity: function transferToSource(address recipient, address feeTokenContractAddress, uint256 feeAmount) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) TransferToSource(recipient common.Address, feeTokenContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferToSource(&_NativeTokenDestination.TransactOpts, recipient, feeTokenContractAddress, feeAmount)
}

// TransferToSource is a paid mutator transaction binding the contract method 0x281c91de.
//
// Solidity: function transferToSource(address recipient, address feeTokenContractAddress, uint256 feeAmount) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) TransferToSource(recipient common.Address, feeTokenContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferToSource(&_NativeTokenDestination.TransactOpts, recipient, feeTokenContractAddress, feeAmount)
}

// NativeTokenDestinationMintNativeTokensIterator is returned from FilterMintNativeTokens and is used to iterate over the raw logs and unpacked data for MintNativeTokens events raised by the NativeTokenDestination contract.
type NativeTokenDestinationMintNativeTokensIterator struct {
	Event *NativeTokenDestinationMintNativeTokens // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationMintNativeTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationMintNativeTokens)
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
		it.Event = new(NativeTokenDestinationMintNativeTokens)
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
func (it *NativeTokenDestinationMintNativeTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationMintNativeTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationMintNativeTokens represents a MintNativeTokens event raised by the NativeTokenDestination contract.
type NativeTokenDestinationMintNativeTokens struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintNativeTokens is a free log retrieval operation binding the contract event 0xe2899f7ef7618206fa13efbcde0c05cccd83f35ac3b18ef860021181d61fa680.
//
// Solidity: event MintNativeTokens(address recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterMintNativeTokens(opts *bind.FilterOpts) (*NativeTokenDestinationMintNativeTokensIterator, error) {

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "MintNativeTokens")
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationMintNativeTokensIterator{contract: _NativeTokenDestination.contract, event: "MintNativeTokens", logs: logs, sub: sub}, nil
}

// WatchMintNativeTokens is a free log subscription operation binding the contract event 0xe2899f7ef7618206fa13efbcde0c05cccd83f35ac3b18ef860021181d61fa680.
//
// Solidity: event MintNativeTokens(address recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchMintNativeTokens(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationMintNativeTokens) (event.Subscription, error) {

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "MintNativeTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationMintNativeTokens)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "MintNativeTokens", log); err != nil {
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
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseMintNativeTokens(log types.Log) (*NativeTokenDestinationMintNativeTokens, error) {
	event := new(NativeTokenDestinationMintNativeTokens)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "MintNativeTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTransferToSourceIterator is returned from FilterTransferToSource and is used to iterate over the raw logs and unpacked data for TransferToSource events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTransferToSourceIterator struct {
	Event *NativeTokenDestinationTransferToSource // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationTransferToSourceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTransferToSource)
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
		it.Event = new(NativeTokenDestinationTransferToSource)
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
func (it *NativeTokenDestinationTransferToSourceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTransferToSourceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTransferToSource represents a TransferToSource event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTransferToSource struct {
	TokenContractAddress     common.Address
	TeleporterMessageID      *big.Int
	DestinationChainID       [32]byte
	DestinationBridgeAddress common.Address
	Recipient                common.Address
	TransferAmount           *big.Int
	FeeAmount                *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterTransferToSource is a free log retrieval operation binding the contract event 0xd68353fcc47e706f53e30e6761a3cf2576d0e791c663e962b868bf38010db9a3.
//
// Solidity: event TransferToSource(address indexed tokenContractAddress, uint256 indexed teleporterMessageID, bytes32 destinationChainID, address destinationBridgeAddress, address recipient, uint256 transferAmount, uint256 feeAmount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTransferToSource(opts *bind.FilterOpts, tokenContractAddress []common.Address, teleporterMessageID []*big.Int) (*NativeTokenDestinationTransferToSourceIterator, error) {

	var tokenContractAddressRule []interface{}
	for _, tokenContractAddressItem := range tokenContractAddress {
		tokenContractAddressRule = append(tokenContractAddressRule, tokenContractAddressItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TransferToSource", tokenContractAddressRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTransferToSourceIterator{contract: _NativeTokenDestination.contract, event: "TransferToSource", logs: logs, sub: sub}, nil
}

// WatchTransferToSource is a free log subscription operation binding the contract event 0xd68353fcc47e706f53e30e6761a3cf2576d0e791c663e962b868bf38010db9a3.
//
// Solidity: event TransferToSource(address indexed tokenContractAddress, uint256 indexed teleporterMessageID, bytes32 destinationChainID, address destinationBridgeAddress, address recipient, uint256 transferAmount, uint256 feeAmount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTransferToSource(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTransferToSource, tokenContractAddress []common.Address, teleporterMessageID []*big.Int) (event.Subscription, error) {

	var tokenContractAddressRule []interface{}
	for _, tokenContractAddressItem := range tokenContractAddress {
		tokenContractAddressRule = append(tokenContractAddressRule, tokenContractAddressItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TransferToSource", tokenContractAddressRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTransferToSource)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TransferToSource", log); err != nil {
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

// ParseTransferToSource is a log parse operation binding the contract event 0xd68353fcc47e706f53e30e6761a3cf2576d0e791c663e962b868bf38010db9a3.
//
// Solidity: event TransferToSource(address indexed tokenContractAddress, uint256 indexed teleporterMessageID, bytes32 destinationChainID, address destinationBridgeAddress, address recipient, uint256 transferAmount, uint256 feeAmount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTransferToSource(log types.Log) (*NativeTokenDestinationTransferToSource, error) {
	event := new(NativeTokenDestinationTransferToSource)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TransferToSource", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
