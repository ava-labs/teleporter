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

// NativeTokenMinterMetaData contains all meta data concerning the NativeTokenMinter contract.
var NativeTokenMinterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"partnerChainID_\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceNotIncreased\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotBridgeTokenWithinSameChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"adjustedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientAdjustedAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPartnerContractAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeleporterMessengerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"BridgeTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MintNativeTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TRANSFER_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_PRECOMPILE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"bridgeTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partnerChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nativeChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NativeTokenMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenMinterMetaData.ABI instead.
var NativeTokenMinterABI = NativeTokenMinterMetaData.ABI

// NativeTokenMinter is an auto generated Go binding around an Ethereum contract.
type NativeTokenMinter struct {
	NativeTokenMinterCaller     // Read-only binding to the contract
	NativeTokenMinterTransactor // Write-only binding to the contract
	NativeTokenMinterFilterer   // Log filterer for contract events
}

// NativeTokenMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenMinterSession struct {
	Contract     *NativeTokenMinter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NativeTokenMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenMinterCallerSession struct {
	Contract *NativeTokenMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NativeTokenMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenMinterTransactorSession struct {
	Contract     *NativeTokenMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NativeTokenMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenMinterRaw struct {
	Contract *NativeTokenMinter // Generic contract binding to access the raw methods on
}

// NativeTokenMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenMinterCallerRaw struct {
	Contract *NativeTokenMinterCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenMinterTransactorRaw struct {
	Contract *NativeTokenMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenMinter creates a new instance of NativeTokenMinter, bound to a specific deployed contract.
func NewNativeTokenMinter(address common.Address, backend bind.ContractBackend) (*NativeTokenMinter, error) {
	contract, err := bindNativeTokenMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenMinter{NativeTokenMinterCaller: NativeTokenMinterCaller{contract: contract}, NativeTokenMinterTransactor: NativeTokenMinterTransactor{contract: contract}, NativeTokenMinterFilterer: NativeTokenMinterFilterer{contract: contract}}, nil
}

// NewNativeTokenMinterCaller creates a new read-only instance of NativeTokenMinter, bound to a specific deployed contract.
func NewNativeTokenMinterCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenMinterCaller, error) {
	contract, err := bindNativeTokenMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenMinterCaller{contract: contract}, nil
}

// NewNativeTokenMinterTransactor creates a new write-only instance of NativeTokenMinter, bound to a specific deployed contract.
func NewNativeTokenMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenMinterTransactor, error) {
	contract, err := bindNativeTokenMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenMinterTransactor{contract: contract}, nil
}

// NewNativeTokenMinterFilterer creates a new log filterer instance of NativeTokenMinter, bound to a specific deployed contract.
func NewNativeTokenMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenMinterFilterer, error) {
	contract, err := bindNativeTokenMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenMinterFilterer{contract: contract}, nil
}

// bindNativeTokenMinter binds a generic wrapper to an already deployed contract.
func bindNativeTokenMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenMinterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenMinter *NativeTokenMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenMinter.Contract.NativeTokenMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenMinter *NativeTokenMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenMinter.Contract.NativeTokenMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenMinter *NativeTokenMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenMinter.Contract.NativeTokenMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenMinter *NativeTokenMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenMinter *NativeTokenMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenMinter *NativeTokenMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenMinter.Contract.contract.Transact(opts, method, params...)
}

// TRANSFERNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xab285230.
//
// Solidity: function TRANSFER_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenMinter *NativeTokenMinterCaller) TRANSFERNATIVETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenMinter.contract.Call(opts, &out, "TRANSFER_NATIVE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TRANSFERNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xab285230.
//
// Solidity: function TRANSFER_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenMinter *NativeTokenMinterSession) TRANSFERNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenMinter.Contract.TRANSFERNATIVETOKENSREQUIREDGAS(&_NativeTokenMinter.CallOpts)
}

// TRANSFERNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xab285230.
//
// Solidity: function TRANSFER_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenMinter *NativeTokenMinterCallerSession) TRANSFERNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenMinter.Contract.TRANSFERNATIVETOKENSREQUIREDGAS(&_NativeTokenMinter.CallOpts)
}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_NativeTokenMinter *NativeTokenMinterCaller) WARPPRECOMPILEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenMinter.contract.Call(opts, &out, "WARP_PRECOMPILE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_NativeTokenMinter *NativeTokenMinterSession) WARPPRECOMPILEADDRESS() (common.Address, error) {
	return _NativeTokenMinter.Contract.WARPPRECOMPILEADDRESS(&_NativeTokenMinter.CallOpts)
}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_NativeTokenMinter *NativeTokenMinterCallerSession) WARPPRECOMPILEADDRESS() (common.Address, error) {
	return _NativeTokenMinter.Contract.WARPPRECOMPILEADDRESS(&_NativeTokenMinter.CallOpts)
}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_NativeTokenMinter *NativeTokenMinterCaller) CurrentChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenMinter.contract.Call(opts, &out, "currentChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_NativeTokenMinter *NativeTokenMinterSession) CurrentChainID() ([32]byte, error) {
	return _NativeTokenMinter.Contract.CurrentChainID(&_NativeTokenMinter.CallOpts)
}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_NativeTokenMinter *NativeTokenMinterCallerSession) CurrentChainID() ([32]byte, error) {
	return _NativeTokenMinter.Contract.CurrentChainID(&_NativeTokenMinter.CallOpts)
}

// PartnerChainID is a free data retrieval call binding the contract method 0x9c224734.
//
// Solidity: function partnerChainID() view returns(bytes32)
func (_NativeTokenMinter *NativeTokenMinterCaller) PartnerChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenMinter.contract.Call(opts, &out, "partnerChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PartnerChainID is a free data retrieval call binding the contract method 0x9c224734.
//
// Solidity: function partnerChainID() view returns(bytes32)
func (_NativeTokenMinter *NativeTokenMinterSession) PartnerChainID() ([32]byte, error) {
	return _NativeTokenMinter.Contract.PartnerChainID(&_NativeTokenMinter.CallOpts)
}

// PartnerChainID is a free data retrieval call binding the contract method 0x9c224734.
//
// Solidity: function partnerChainID() view returns(bytes32)
func (_NativeTokenMinter *NativeTokenMinterCallerSession) PartnerChainID() ([32]byte, error) {
	return _NativeTokenMinter.Contract.PartnerChainID(&_NativeTokenMinter.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenMinter *NativeTokenMinterCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenMinter.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenMinter *NativeTokenMinterSession) TeleporterMessenger() (common.Address, error) {
	return _NativeTokenMinter.Contract.TeleporterMessenger(&_NativeTokenMinter.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenMinter *NativeTokenMinterCallerSession) TeleporterMessenger() (common.Address, error) {
	return _NativeTokenMinter.Contract.TeleporterMessenger(&_NativeTokenMinter.CallOpts)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0xf9fe683e.
//
// Solidity: function bridgeTokens(address recipient, address feeTokenContractAddress, uint256 feeAmount) payable returns()
func (_NativeTokenMinter *NativeTokenMinterTransactor) BridgeTokens(opts *bind.TransactOpts, recipient common.Address, feeTokenContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _NativeTokenMinter.contract.Transact(opts, "bridgeTokens", recipient, feeTokenContractAddress, feeAmount)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0xf9fe683e.
//
// Solidity: function bridgeTokens(address recipient, address feeTokenContractAddress, uint256 feeAmount) payable returns()
func (_NativeTokenMinter *NativeTokenMinterSession) BridgeTokens(recipient common.Address, feeTokenContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _NativeTokenMinter.Contract.BridgeTokens(&_NativeTokenMinter.TransactOpts, recipient, feeTokenContractAddress, feeAmount)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0xf9fe683e.
//
// Solidity: function bridgeTokens(address recipient, address feeTokenContractAddress, uint256 feeAmount) payable returns()
func (_NativeTokenMinter *NativeTokenMinterTransactorSession) BridgeTokens(recipient common.Address, feeTokenContractAddress common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _NativeTokenMinter.Contract.BridgeTokens(&_NativeTokenMinter.TransactOpts, recipient, feeTokenContractAddress, feeAmount)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_NativeTokenMinter *NativeTokenMinterTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenMinter.contract.Transact(opts, "receiveTeleporterMessage", nativeChainID, nativeBridgeAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_NativeTokenMinter *NativeTokenMinterSession) ReceiveTeleporterMessage(nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenMinter.Contract.ReceiveTeleporterMessage(&_NativeTokenMinter.TransactOpts, nativeChainID, nativeBridgeAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_NativeTokenMinter *NativeTokenMinterTransactorSession) ReceiveTeleporterMessage(nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenMinter.Contract.ReceiveTeleporterMessage(&_NativeTokenMinter.TransactOpts, nativeChainID, nativeBridgeAddress, message)
}

// NativeTokenMinterBridgeTokensIterator is returned from FilterBridgeTokens and is used to iterate over the raw logs and unpacked data for BridgeTokens events raised by the NativeTokenMinter contract.
type NativeTokenMinterBridgeTokensIterator struct {
	Event *NativeTokenMinterBridgeTokens // Event containing the contract specifics and raw log

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
func (it *NativeTokenMinterBridgeTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenMinterBridgeTokens)
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
		it.Event = new(NativeTokenMinterBridgeTokens)
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
func (it *NativeTokenMinterBridgeTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenMinterBridgeTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenMinterBridgeTokens represents a BridgeTokens event raised by the NativeTokenMinter contract.
type NativeTokenMinterBridgeTokens struct {
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
func (_NativeTokenMinter *NativeTokenMinterFilterer) FilterBridgeTokens(opts *bind.FilterOpts, tokenContractAddress []common.Address, teleporterMessageID []*big.Int) (*NativeTokenMinterBridgeTokensIterator, error) {

	var tokenContractAddressRule []interface{}
	for _, tokenContractAddressItem := range tokenContractAddress {
		tokenContractAddressRule = append(tokenContractAddressRule, tokenContractAddressItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenMinter.contract.FilterLogs(opts, "BridgeTokens", tokenContractAddressRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenMinterBridgeTokensIterator{contract: _NativeTokenMinter.contract, event: "BridgeTokens", logs: logs, sub: sub}, nil
}

// WatchBridgeTokens is a free log subscription operation binding the contract event 0xfd68c031379d96b674b8ff93f8bba98792a892b18f15c19e26dc84485f4f33e2.
//
// Solidity: event BridgeTokens(address indexed tokenContractAddress, uint256 indexed teleporterMessageID, bytes32 destinationChainID, address destinationBridgeAddress, address recipient, uint256 transferAmount, uint256 feeAmount)
func (_NativeTokenMinter *NativeTokenMinterFilterer) WatchBridgeTokens(opts *bind.WatchOpts, sink chan<- *NativeTokenMinterBridgeTokens, tokenContractAddress []common.Address, teleporterMessageID []*big.Int) (event.Subscription, error) {

	var tokenContractAddressRule []interface{}
	for _, tokenContractAddressItem := range tokenContractAddress {
		tokenContractAddressRule = append(tokenContractAddressRule, tokenContractAddressItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenMinter.contract.WatchLogs(opts, "BridgeTokens", tokenContractAddressRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenMinterBridgeTokens)
				if err := _NativeTokenMinter.contract.UnpackLog(event, "BridgeTokens", log); err != nil {
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
func (_NativeTokenMinter *NativeTokenMinterFilterer) ParseBridgeTokens(log types.Log) (*NativeTokenMinterBridgeTokens, error) {
	event := new(NativeTokenMinterBridgeTokens)
	if err := _NativeTokenMinter.contract.UnpackLog(event, "BridgeTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenMinterMintNativeTokensIterator is returned from FilterMintNativeTokens and is used to iterate over the raw logs and unpacked data for MintNativeTokens events raised by the NativeTokenMinter contract.
type NativeTokenMinterMintNativeTokensIterator struct {
	Event *NativeTokenMinterMintNativeTokens // Event containing the contract specifics and raw log

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
func (it *NativeTokenMinterMintNativeTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenMinterMintNativeTokens)
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
		it.Event = new(NativeTokenMinterMintNativeTokens)
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
func (it *NativeTokenMinterMintNativeTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenMinterMintNativeTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenMinterMintNativeTokens represents a MintNativeTokens event raised by the NativeTokenMinter contract.
type NativeTokenMinterMintNativeTokens struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintNativeTokens is a free log retrieval operation binding the contract event 0xe2899f7ef7618206fa13efbcde0c05cccd83f35ac3b18ef860021181d61fa680.
//
// Solidity: event MintNativeTokens(address recipient, uint256 amount)
func (_NativeTokenMinter *NativeTokenMinterFilterer) FilterMintNativeTokens(opts *bind.FilterOpts) (*NativeTokenMinterMintNativeTokensIterator, error) {

	logs, sub, err := _NativeTokenMinter.contract.FilterLogs(opts, "MintNativeTokens")
	if err != nil {
		return nil, err
	}
	return &NativeTokenMinterMintNativeTokensIterator{contract: _NativeTokenMinter.contract, event: "MintNativeTokens", logs: logs, sub: sub}, nil
}

// WatchMintNativeTokens is a free log subscription operation binding the contract event 0xe2899f7ef7618206fa13efbcde0c05cccd83f35ac3b18ef860021181d61fa680.
//
// Solidity: event MintNativeTokens(address recipient, uint256 amount)
func (_NativeTokenMinter *NativeTokenMinterFilterer) WatchMintNativeTokens(opts *bind.WatchOpts, sink chan<- *NativeTokenMinterMintNativeTokens) (event.Subscription, error) {

	logs, sub, err := _NativeTokenMinter.contract.WatchLogs(opts, "MintNativeTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenMinterMintNativeTokens)
				if err := _NativeTokenMinter.contract.UnpackLog(event, "MintNativeTokens", log); err != nil {
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
func (_NativeTokenMinter *NativeTokenMinterFilterer) ParseMintNativeTokens(log types.Log) (*NativeTokenMinterMintNativeTokens, error) {
	event := new(NativeTokenMinterMintNativeTokens)
	if err := _NativeTokenMinter.contract.UnpackLog(event, "MintNativeTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
