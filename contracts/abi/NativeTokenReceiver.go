// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// NativeTokenReceiverMetaData contains all meta data concerning the NativeTokenReceiver contract.
var NativeTokenReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"partnerChainID_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceNotIncreased\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotBridgeTokenWithinSameChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"adjustedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientAdjustedAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPartnerContractAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeleporterMessengerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"BridgeTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINT_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_PRECOMPILE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"bridgeTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partnerChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nativeChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NativeTokenReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenReceiverMetaData.ABI instead.
var NativeTokenReceiverABI = NativeTokenReceiverMetaData.ABI

// NativeTokenReceiver is an auto generated Go binding around an Ethereum contract.
type NativeTokenReceiver struct {
	NativeTokenReceiverCaller     // Read-only binding to the contract
	NativeTokenReceiverTransactor // Write-only binding to the contract
	NativeTokenReceiverFilterer   // Log filterer for contract events
}

// NativeTokenReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenReceiverSession struct {
	Contract     *NativeTokenReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NativeTokenReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenReceiverCallerSession struct {
	Contract *NativeTokenReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// NativeTokenReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenReceiverTransactorSession struct {
	Contract     *NativeTokenReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// NativeTokenReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenReceiverRaw struct {
	Contract *NativeTokenReceiver // Generic contract binding to access the raw methods on
}

// NativeTokenReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenReceiverCallerRaw struct {
	Contract *NativeTokenReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenReceiverTransactorRaw struct {
	Contract *NativeTokenReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenReceiver creates a new instance of NativeTokenReceiver, bound to a specific deployed contract.
func NewNativeTokenReceiver(address common.Address, backend bind.ContractBackend) (*NativeTokenReceiver, error) {
	contract, err := bindNativeTokenReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenReceiver{NativeTokenReceiverCaller: NativeTokenReceiverCaller{contract: contract}, NativeTokenReceiverTransactor: NativeTokenReceiverTransactor{contract: contract}, NativeTokenReceiverFilterer: NativeTokenReceiverFilterer{contract: contract}}, nil
}

// NewNativeTokenReceiverCaller creates a new read-only instance of NativeTokenReceiver, bound to a specific deployed contract.
func NewNativeTokenReceiverCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenReceiverCaller, error) {
	contract, err := bindNativeTokenReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenReceiverCaller{contract: contract}, nil
}

// NewNativeTokenReceiverTransactor creates a new write-only instance of NativeTokenReceiver, bound to a specific deployed contract.
func NewNativeTokenReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenReceiverTransactor, error) {
	contract, err := bindNativeTokenReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenReceiverTransactor{contract: contract}, nil
}

// NewNativeTokenReceiverFilterer creates a new log filterer instance of NativeTokenReceiver, bound to a specific deployed contract.
func NewNativeTokenReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenReceiverFilterer, error) {
	contract, err := bindNativeTokenReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenReceiverFilterer{contract: contract}, nil
}

// bindNativeTokenReceiver binds a generic wrapper to an already deployed contract.
func bindNativeTokenReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenReceiver *NativeTokenReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenReceiver.Contract.NativeTokenReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenReceiver *NativeTokenReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenReceiver.Contract.NativeTokenReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenReceiver *NativeTokenReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenReceiver.Contract.NativeTokenReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenReceiver *NativeTokenReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenReceiver *NativeTokenReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenReceiver *NativeTokenReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenReceiver.Contract.contract.Transact(opts, method, params...)
}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenReceiver *NativeTokenReceiverCaller) MINTNATIVETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenReceiver.contract.Call(opts, &out, "MINT_NATIVE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenReceiver *NativeTokenReceiverSession) MINTNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenReceiver.Contract.MINTNATIVETOKENSREQUIREDGAS(&_NativeTokenReceiver.CallOpts)
}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenReceiver *NativeTokenReceiverCallerSession) MINTNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenReceiver.Contract.MINTNATIVETOKENSREQUIREDGAS(&_NativeTokenReceiver.CallOpts)
}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_NativeTokenReceiver *NativeTokenReceiverCaller) WARPPRECOMPILEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenReceiver.contract.Call(opts, &out, "WARP_PRECOMPILE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_NativeTokenReceiver *NativeTokenReceiverSession) WARPPRECOMPILEADDRESS() (common.Address, error) {
	return _NativeTokenReceiver.Contract.WARPPRECOMPILEADDRESS(&_NativeTokenReceiver.CallOpts)
}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_NativeTokenReceiver *NativeTokenReceiverCallerSession) WARPPRECOMPILEADDRESS() (common.Address, error) {
	return _NativeTokenReceiver.Contract.WARPPRECOMPILEADDRESS(&_NativeTokenReceiver.CallOpts)
}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_NativeTokenReceiver *NativeTokenReceiverCaller) CurrentChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenReceiver.contract.Call(opts, &out, "currentChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_NativeTokenReceiver *NativeTokenReceiverSession) CurrentChainID() ([32]byte, error) {
	return _NativeTokenReceiver.Contract.CurrentChainID(&_NativeTokenReceiver.CallOpts)
}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_NativeTokenReceiver *NativeTokenReceiverCallerSession) CurrentChainID() ([32]byte, error) {
	return _NativeTokenReceiver.Contract.CurrentChainID(&_NativeTokenReceiver.CallOpts)
}

// PartnerChainID is a free data retrieval call binding the contract method 0x9c224734.
//
// Solidity: function partnerChainID() view returns(bytes32)
func (_NativeTokenReceiver *NativeTokenReceiverCaller) PartnerChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenReceiver.contract.Call(opts, &out, "partnerChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PartnerChainID is a free data retrieval call binding the contract method 0x9c224734.
//
// Solidity: function partnerChainID() view returns(bytes32)
func (_NativeTokenReceiver *NativeTokenReceiverSession) PartnerChainID() ([32]byte, error) {
	return _NativeTokenReceiver.Contract.PartnerChainID(&_NativeTokenReceiver.CallOpts)
}

// PartnerChainID is a free data retrieval call binding the contract method 0x9c224734.
//
// Solidity: function partnerChainID() view returns(bytes32)
func (_NativeTokenReceiver *NativeTokenReceiverCallerSession) PartnerChainID() ([32]byte, error) {
	return _NativeTokenReceiver.Contract.PartnerChainID(&_NativeTokenReceiver.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenReceiver *NativeTokenReceiverCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenReceiver.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenReceiver *NativeTokenReceiverSession) TeleporterMessenger() (common.Address, error) {
	return _NativeTokenReceiver.Contract.TeleporterMessenger(&_NativeTokenReceiver.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenReceiver *NativeTokenReceiverCallerSession) TeleporterMessenger() (common.Address, error) {
	return _NativeTokenReceiver.Contract.TeleporterMessenger(&_NativeTokenReceiver.CallOpts)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0xf9fe683e.
//
// Solidity: function bridgeTokens(address recipient, address feeTokenContractAddress, uint256 feeAmount) payable returns()
func (_NativeTokenReceiver *NativeTokenReceiverTransactor) BridgeTokens(opts *bind.TransactOpts, recipient common.Address, feeTokenContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _NativeTokenReceiver.contract.Transact(opts, "bridgeTokens", recipient, feeTokenContractAddress, feeAmount)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0xf9fe683e.
//
// Solidity: function bridgeTokens(address recipient, address feeTokenContractAddress, uint256 feeAmount) payable returns()
func (_NativeTokenReceiver *NativeTokenReceiverSession) BridgeTokens(recipient common.Address, feeTokenContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _NativeTokenReceiver.Contract.BridgeTokens(&_NativeTokenReceiver.TransactOpts, recipient, feeTokenContractAddress, feeAmount)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0xf9fe683e.
//
// Solidity: function bridgeTokens(address recipient, address feeTokenContractAddress, uint256 feeAmount) payable returns()
func (_NativeTokenReceiver *NativeTokenReceiverTransactorSession) BridgeTokens(recipient common.Address, feeTokenContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _NativeTokenReceiver.Contract.BridgeTokens(&_NativeTokenReceiver.TransactOpts, recipient, feeTokenContractAddress, feeAmount)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_NativeTokenReceiver *NativeTokenReceiverTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenReceiver.contract.Transact(opts, "receiveTeleporterMessage", nativeChainID, nativeBridgeAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_NativeTokenReceiver *NativeTokenReceiverSession) ReceiveTeleporterMessage(nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenReceiver.Contract.ReceiveTeleporterMessage(&_NativeTokenReceiver.TransactOpts, nativeChainID, nativeBridgeAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_NativeTokenReceiver *NativeTokenReceiverTransactorSession) ReceiveTeleporterMessage(nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenReceiver.Contract.ReceiveTeleporterMessage(&_NativeTokenReceiver.TransactOpts, nativeChainID, nativeBridgeAddress, message)
}

// NativeTokenReceiverBridgeTokensIterator is returned from FilterBridgeTokens and is used to iterate over the raw logs and unpacked data for BridgeTokens events raised by the NativeTokenReceiver contract.
type NativeTokenReceiverBridgeTokensIterator struct {
	Event *NativeTokenReceiverBridgeTokens // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenReceiverBridgeTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenReceiverBridgeTokens)
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
		it.Event = new(NativeTokenReceiverBridgeTokens)
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
func (it *NativeTokenReceiverBridgeTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenReceiverBridgeTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenReceiverBridgeTokens represents a BridgeTokens event raised by the NativeTokenReceiver contract.
type NativeTokenReceiverBridgeTokens struct {
	TokenContractAddress     common.Address
	TeleporterMessageID      *big.Int
	DestinationChainID       [32]byte
	DestinationBridgeAddress common.Address
	Recipient                common.Address
	TransferAmount           *big.Int
	FeeAmount                *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterBridgeTokens is a free log retrieval operation binding the contract event 0xfd68c031379d96b674b8ff93f8bba98792a892b18f15c19e26dc84485f4f33e2.
//
// Solidity: event BridgeTokens(address indexed tokenContractAddress, uint256 indexed teleporterMessageID, bytes32 destinationChainID, address destinationBridgeAddress, address recipient, uint256 transferAmount, uint256 feeAmount)
func (_NativeTokenReceiver *NativeTokenReceiverFilterer) FilterBridgeTokens(opts *bind.FilterOpts, tokenContractAddress []common.Address, teleporterMessageID []*big.Int) (*NativeTokenReceiverBridgeTokensIterator, error) {

	var tokenContractAddressRule []interface{}
	for _, tokenContractAddressItem := range tokenContractAddress {
		tokenContractAddressRule = append(tokenContractAddressRule, tokenContractAddressItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenReceiver.contract.FilterLogs(opts, "BridgeTokens", tokenContractAddressRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenReceiverBridgeTokensIterator{contract: _NativeTokenReceiver.contract, event: "BridgeTokens", logs: logs, sub: sub}, nil
}

// WatchBridgeTokens is a free log subscription operation binding the contract event 0xfd68c031379d96b674b8ff93f8bba98792a892b18f15c19e26dc84485f4f33e2.
//
// Solidity: event BridgeTokens(address indexed tokenContractAddress, uint256 indexed teleporterMessageID, bytes32 destinationChainID, address destinationBridgeAddress, address recipient, uint256 transferAmount, uint256 feeAmount)
func (_NativeTokenReceiver *NativeTokenReceiverFilterer) WatchBridgeTokens(opts *bind.WatchOpts, sink chan<- *NativeTokenReceiverBridgeTokens, tokenContractAddress []common.Address, teleporterMessageID []*big.Int) (event.Subscription, error) {

	var tokenContractAddressRule []interface{}
	for _, tokenContractAddressItem := range tokenContractAddress {
		tokenContractAddressRule = append(tokenContractAddressRule, tokenContractAddressItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenReceiver.contract.WatchLogs(opts, "BridgeTokens", tokenContractAddressRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenReceiverBridgeTokens)
				if err := _NativeTokenReceiver.contract.UnpackLog(event, "BridgeTokens", log); err != nil {
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

// ParseBridgeTokens is a log parse operation binding the contract event 0xfd68c031379d96b674b8ff93f8bba98792a892b18f15c19e26dc84485f4f33e2.
//
// Solidity: event BridgeTokens(address indexed tokenContractAddress, uint256 indexed teleporterMessageID, bytes32 destinationChainID, address destinationBridgeAddress, address recipient, uint256 transferAmount, uint256 feeAmount)
func (_NativeTokenReceiver *NativeTokenReceiverFilterer) ParseBridgeTokens(log types.Log) (*NativeTokenReceiverBridgeTokens, error) {
	event := new(NativeTokenReceiverBridgeTokens)
	if err := _NativeTokenReceiver.contract.UnpackLog(event, "BridgeTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenReceiverUnlockTokensIterator is returned from FilterUnlockTokens and is used to iterate over the raw logs and unpacked data for UnlockTokens events raised by the NativeTokenReceiver contract.
type NativeTokenReceiverUnlockTokensIterator struct {
	Event *NativeTokenReceiverUnlockTokens // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeTokenReceiverUnlockTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenReceiverUnlockTokens)
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
		it.Event = new(NativeTokenReceiverUnlockTokens)
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
func (it *NativeTokenReceiverUnlockTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenReceiverUnlockTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenReceiverUnlockTokens represents a UnlockTokens event raised by the NativeTokenReceiver contract.
type NativeTokenReceiverUnlockTokens struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlockTokens is a free log retrieval operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_NativeTokenReceiver *NativeTokenReceiverFilterer) FilterUnlockTokens(opts *bind.FilterOpts) (*NativeTokenReceiverUnlockTokensIterator, error) {

	logs, sub, err := _NativeTokenReceiver.contract.FilterLogs(opts, "UnlockTokens")
	if err != nil {
		return nil, err
	}
	return &NativeTokenReceiverUnlockTokensIterator{contract: _NativeTokenReceiver.contract, event: "UnlockTokens", logs: logs, sub: sub}, nil
}

// WatchUnlockTokens is a free log subscription operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_NativeTokenReceiver *NativeTokenReceiverFilterer) WatchUnlockTokens(opts *bind.WatchOpts, sink chan<- *NativeTokenReceiverUnlockTokens) (event.Subscription, error) {

	logs, sub, err := _NativeTokenReceiver.contract.WatchLogs(opts, "UnlockTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenReceiverUnlockTokens)
				if err := _NativeTokenReceiver.contract.UnpackLog(event, "UnlockTokens", log); err != nil {
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
func (_NativeTokenReceiver *NativeTokenReceiverFilterer) ParseUnlockTokens(log types.Log) (*NativeTokenReceiverUnlockTokens, error) {
	event := new(NativeTokenReceiverUnlockTokens)
	if err := _NativeTokenReceiver.contract.UnpackLog(event, "UnlockTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
