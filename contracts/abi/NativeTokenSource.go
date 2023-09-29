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

// NativeTokenSourceMetaData contains all meta data concerning the NativeTokenSource contract.
var NativeTokenSourceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destinationContractAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationChainID_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceNotIncreased\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotBridgeTokenWithinSameChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"adjustedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientAdjustedAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationContractAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPartnerContractAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeleporterMessengerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"TransferToDestination\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINT_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_PRECOMPILE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nativeChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"transferToDestination\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NativeTokenSourceABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenSourceMetaData.ABI instead.
var NativeTokenSourceABI = NativeTokenSourceMetaData.ABI

// NativeTokenSource is an auto generated Go binding around an Ethereum contract.
type NativeTokenSource struct {
	NativeTokenSourceCaller     // Read-only binding to the contract
	NativeTokenSourceTransactor // Write-only binding to the contract
	NativeTokenSourceFilterer   // Log filterer for contract events
}

// NativeTokenSourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenSourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenSourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenSourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenSourceSession struct {
	Contract     *NativeTokenSource // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NativeTokenSourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenSourceCallerSession struct {
	Contract *NativeTokenSourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NativeTokenSourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenSourceTransactorSession struct {
	Contract     *NativeTokenSourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NativeTokenSourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenSourceRaw struct {
	Contract *NativeTokenSource // Generic contract binding to access the raw methods on
}

// NativeTokenSourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenSourceCallerRaw struct {
	Contract *NativeTokenSourceCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenSourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenSourceTransactorRaw struct {
	Contract *NativeTokenSourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenSource creates a new instance of NativeTokenSource, bound to a specific deployed contract.
func NewNativeTokenSource(address common.Address, backend bind.ContractBackend) (*NativeTokenSource, error) {
	contract, err := bindNativeTokenSource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSource{NativeTokenSourceCaller: NativeTokenSourceCaller{contract: contract}, NativeTokenSourceTransactor: NativeTokenSourceTransactor{contract: contract}, NativeTokenSourceFilterer: NativeTokenSourceFilterer{contract: contract}}, nil
}

// NewNativeTokenSourceCaller creates a new read-only instance of NativeTokenSource, bound to a specific deployed contract.
func NewNativeTokenSourceCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenSourceCaller, error) {
	contract, err := bindNativeTokenSource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceCaller{contract: contract}, nil
}

// NewNativeTokenSourceTransactor creates a new write-only instance of NativeTokenSource, bound to a specific deployed contract.
func NewNativeTokenSourceTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenSourceTransactor, error) {
	contract, err := bindNativeTokenSource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceTransactor{contract: contract}, nil
}

// NewNativeTokenSourceFilterer creates a new log filterer instance of NativeTokenSource, bound to a specific deployed contract.
func NewNativeTokenSourceFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenSourceFilterer, error) {
	contract, err := bindNativeTokenSource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceFilterer{contract: contract}, nil
}

// bindNativeTokenSource binds a generic wrapper to an already deployed contract.
func bindNativeTokenSource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenSourceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenSource *NativeTokenSourceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenSource.Contract.NativeTokenSourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenSource *NativeTokenSourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.NativeTokenSourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenSource *NativeTokenSourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.NativeTokenSourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenSource *NativeTokenSourceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenSource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenSource *NativeTokenSourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenSource *NativeTokenSourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.contract.Transact(opts, method, params...)
}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceCaller) MINTNATIVETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "MINT_NATIVE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceSession) MINTNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSource.Contract.MINTNATIVETOKENSREQUIREDGAS(&_NativeTokenSource.CallOpts)
}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceCallerSession) MINTNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSource.Contract.MINTNATIVETOKENSREQUIREDGAS(&_NativeTokenSource.CallOpts)
}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCaller) WARPPRECOMPILEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "WARP_PRECOMPILE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_NativeTokenSource *NativeTokenSourceSession) WARPPRECOMPILEADDRESS() (common.Address, error) {
	return _NativeTokenSource.Contract.WARPPRECOMPILEADDRESS(&_NativeTokenSource.CallOpts)
}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCallerSession) WARPPRECOMPILEADDRESS() (common.Address, error) {
	return _NativeTokenSource.Contract.WARPPRECOMPILEADDRESS(&_NativeTokenSource.CallOpts)
}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_NativeTokenSource *NativeTokenSourceCaller) CurrentChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "currentChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_NativeTokenSource *NativeTokenSourceSession) CurrentChainID() ([32]byte, error) {
	return _NativeTokenSource.Contract.CurrentChainID(&_NativeTokenSource.CallOpts)
}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_NativeTokenSource *NativeTokenSourceCallerSession) CurrentChainID() ([32]byte, error) {
	return _NativeTokenSource.Contract.CurrentChainID(&_NativeTokenSource.CallOpts)
}

// DestinationChainID is a free data retrieval call binding the contract method 0x7a9cffc8.
//
// Solidity: function destinationChainID() view returns(bytes32)
func (_NativeTokenSource *NativeTokenSourceCaller) DestinationChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "destinationChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DestinationChainID is a free data retrieval call binding the contract method 0x7a9cffc8.
//
// Solidity: function destinationChainID() view returns(bytes32)
func (_NativeTokenSource *NativeTokenSourceSession) DestinationChainID() ([32]byte, error) {
	return _NativeTokenSource.Contract.DestinationChainID(&_NativeTokenSource.CallOpts)
}

// DestinationChainID is a free data retrieval call binding the contract method 0x7a9cffc8.
//
// Solidity: function destinationChainID() view returns(bytes32)
func (_NativeTokenSource *NativeTokenSourceCallerSession) DestinationChainID() ([32]byte, error) {
	return _NativeTokenSource.Contract.DestinationChainID(&_NativeTokenSource.CallOpts)
}

// DestinationContractAddress is a free data retrieval call binding the contract method 0x04d6baf7.
//
// Solidity: function destinationContractAddress() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCaller) DestinationContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "destinationContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DestinationContractAddress is a free data retrieval call binding the contract method 0x04d6baf7.
//
// Solidity: function destinationContractAddress() view returns(address)
func (_NativeTokenSource *NativeTokenSourceSession) DestinationContractAddress() (common.Address, error) {
	return _NativeTokenSource.Contract.DestinationContractAddress(&_NativeTokenSource.CallOpts)
}

// DestinationContractAddress is a free data retrieval call binding the contract method 0x04d6baf7.
//
// Solidity: function destinationContractAddress() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCallerSession) DestinationContractAddress() (common.Address, error) {
	return _NativeTokenSource.Contract.DestinationContractAddress(&_NativeTokenSource.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenSource *NativeTokenSourceSession) TeleporterMessenger() (common.Address, error) {
	return _NativeTokenSource.Contract.TeleporterMessenger(&_NativeTokenSource.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCallerSession) TeleporterMessenger() (common.Address, error) {
	return _NativeTokenSource.Contract.TeleporterMessenger(&_NativeTokenSource.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_NativeTokenSource *NativeTokenSourceTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSource.contract.Transact(opts, "receiveTeleporterMessage", nativeChainID, nativeBridgeAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_NativeTokenSource *NativeTokenSourceSession) ReceiveTeleporterMessage(nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.ReceiveTeleporterMessage(&_NativeTokenSource.TransactOpts, nativeChainID, nativeBridgeAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_NativeTokenSource *NativeTokenSourceTransactorSession) ReceiveTeleporterMessage(nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.ReceiveTeleporterMessage(&_NativeTokenSource.TransactOpts, nativeChainID, nativeBridgeAddress, message)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0x3c7bf510.
//
// Solidity: function transferToDestination(address recipient, address feeTokenContractAddress, uint256 feeAmount) payable returns()
func (_NativeTokenSource *NativeTokenSourceTransactor) TransferToDestination(opts *bind.TransactOpts, recipient common.Address, feeTokenContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSource.contract.Transact(opts, "transferToDestination", recipient, feeTokenContractAddress, feeAmount)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0x3c7bf510.
//
// Solidity: function transferToDestination(address recipient, address feeTokenContractAddress, uint256 feeAmount) payable returns()
func (_NativeTokenSource *NativeTokenSourceSession) TransferToDestination(recipient common.Address, feeTokenContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.TransferToDestination(&_NativeTokenSource.TransactOpts, recipient, feeTokenContractAddress, feeAmount)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0x3c7bf510.
//
// Solidity: function transferToDestination(address recipient, address feeTokenContractAddress, uint256 feeAmount) payable returns()
func (_NativeTokenSource *NativeTokenSourceTransactorSession) TransferToDestination(recipient common.Address, feeTokenContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.TransferToDestination(&_NativeTokenSource.TransactOpts, recipient, feeTokenContractAddress, feeAmount)
}

// NativeTokenSourceTransferToDestinationIterator is returned from FilterTransferToDestination and is used to iterate over the raw logs and unpacked data for TransferToDestination events raised by the NativeTokenSource contract.
type NativeTokenSourceTransferToDestinationIterator struct {
	Event *NativeTokenSourceTransferToDestination // Event containing the contract specifics and raw log

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
func (it *NativeTokenSourceTransferToDestinationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSourceTransferToDestination)
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
		it.Event = new(NativeTokenSourceTransferToDestination)
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
func (it *NativeTokenSourceTransferToDestinationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSourceTransferToDestinationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSourceTransferToDestination represents a TransferToDestination event raised by the NativeTokenSource contract.
type NativeTokenSourceTransferToDestination struct {
	TokenContractAddress     common.Address
	TeleporterMessageID      *big.Int
	DestinationChainID       [32]byte
	DestinationBridgeAddress common.Address
	Recipient                common.Address
	TransferAmount           *big.Int
	FeeAmount                *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterTransferToDestination is a free log retrieval operation binding the contract event 0x0bd44c3ad8913eab944c838cd210d05b9c7ec0254d9c1fdb0571af8b573c5b94.
//
// Solidity: event TransferToDestination(address indexed tokenContractAddress, uint256 indexed teleporterMessageID, bytes32 destinationChainID, address destinationBridgeAddress, address recipient, uint256 transferAmount, uint256 feeAmount)
func (_NativeTokenSource *NativeTokenSourceFilterer) FilterTransferToDestination(opts *bind.FilterOpts, tokenContractAddress []common.Address, teleporterMessageID []*big.Int) (*NativeTokenSourceTransferToDestinationIterator, error) {

	var tokenContractAddressRule []interface{}
	for _, tokenContractAddressItem := range tokenContractAddress {
		tokenContractAddressRule = append(tokenContractAddressRule, tokenContractAddressItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenSource.contract.FilterLogs(opts, "TransferToDestination", tokenContractAddressRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceTransferToDestinationIterator{contract: _NativeTokenSource.contract, event: "TransferToDestination", logs: logs, sub: sub}, nil
}

// WatchTransferToDestination is a free log subscription operation binding the contract event 0x0bd44c3ad8913eab944c838cd210d05b9c7ec0254d9c1fdb0571af8b573c5b94.
//
// Solidity: event TransferToDestination(address indexed tokenContractAddress, uint256 indexed teleporterMessageID, bytes32 destinationChainID, address destinationBridgeAddress, address recipient, uint256 transferAmount, uint256 feeAmount)
func (_NativeTokenSource *NativeTokenSourceFilterer) WatchTransferToDestination(opts *bind.WatchOpts, sink chan<- *NativeTokenSourceTransferToDestination, tokenContractAddress []common.Address, teleporterMessageID []*big.Int) (event.Subscription, error) {

	var tokenContractAddressRule []interface{}
	for _, tokenContractAddressItem := range tokenContractAddress {
		tokenContractAddressRule = append(tokenContractAddressRule, tokenContractAddressItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenSource.contract.WatchLogs(opts, "TransferToDestination", tokenContractAddressRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSourceTransferToDestination)
				if err := _NativeTokenSource.contract.UnpackLog(event, "TransferToDestination", log); err != nil {
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

// ParseTransferToDestination is a log parse operation binding the contract event 0x0bd44c3ad8913eab944c838cd210d05b9c7ec0254d9c1fdb0571af8b573c5b94.
//
// Solidity: event TransferToDestination(address indexed tokenContractAddress, uint256 indexed teleporterMessageID, bytes32 destinationChainID, address destinationBridgeAddress, address recipient, uint256 transferAmount, uint256 feeAmount)
func (_NativeTokenSource *NativeTokenSourceFilterer) ParseTransferToDestination(log types.Log) (*NativeTokenSourceTransferToDestination, error) {
	event := new(NativeTokenSourceTransferToDestination)
	if err := _NativeTokenSource.contract.UnpackLog(event, "TransferToDestination", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSourceUnlockTokensIterator is returned from FilterUnlockTokens and is used to iterate over the raw logs and unpacked data for UnlockTokens events raised by the NativeTokenSource contract.
type NativeTokenSourceUnlockTokensIterator struct {
	Event *NativeTokenSourceUnlockTokens // Event containing the contract specifics and raw log

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
func (it *NativeTokenSourceUnlockTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSourceUnlockTokens)
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
		it.Event = new(NativeTokenSourceUnlockTokens)
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
func (it *NativeTokenSourceUnlockTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSourceUnlockTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSourceUnlockTokens represents a UnlockTokens event raised by the NativeTokenSource contract.
type NativeTokenSourceUnlockTokens struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlockTokens is a free log retrieval operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_NativeTokenSource *NativeTokenSourceFilterer) FilterUnlockTokens(opts *bind.FilterOpts) (*NativeTokenSourceUnlockTokensIterator, error) {

	logs, sub, err := _NativeTokenSource.contract.FilterLogs(opts, "UnlockTokens")
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceUnlockTokensIterator{contract: _NativeTokenSource.contract, event: "UnlockTokens", logs: logs, sub: sub}, nil
}

// WatchUnlockTokens is a free log subscription operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_NativeTokenSource *NativeTokenSourceFilterer) WatchUnlockTokens(opts *bind.WatchOpts, sink chan<- *NativeTokenSourceUnlockTokens) (event.Subscription, error) {

	logs, sub, err := _NativeTokenSource.contract.WatchLogs(opts, "UnlockTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSourceUnlockTokens)
				if err := _NativeTokenSource.contract.UnpackLog(event, "UnlockTokens", log); err != nil {
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
func (_NativeTokenSource *NativeTokenSourceFilterer) ParseUnlockTokens(log types.Log) (*NativeTokenSourceUnlockTokens, error) {
	event := new(NativeTokenSourceUnlockTokens)
	if err := _NativeTokenSource.contract.UnpackLog(event, "UnlockTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
