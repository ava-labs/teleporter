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

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// NativeTokenDestinationMetaData contains all meta data concerning the NativeTokenDestination contract.
var NativeTokenDestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeTokenSourceAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialReserveImbalance_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NativeTokensMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAddressBalance\",\"type\":\"uint256\"}],\"name\":\"ReportTotalBurnedTxFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferToSource\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURNED_TX_FEES_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_FOR_TRANSFER_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORT_BURNED_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenSourceAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"senderBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"reportTotalBurnedTxFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"transferToSource\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BURNEDTXFEESADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "BURNED_TX_FEES_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDTXFEESADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDTXFEESADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0xa2a95017.
//
// Solidity: function BURN_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BURNFORTRANSFERADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "BURN_FOR_TRANSFER_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0xa2a95017.
//
// Solidity: function BURN_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) BURNFORTRANSFERADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNFORTRANSFERADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0xa2a95017.
//
// Solidity: function BURN_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BURNFORTRANSFERADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNFORTRANSFERADDRESS(&_NativeTokenDestination.CallOpts)
}

// REPORTBURNEDTOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xd3095126.
//
// Solidity: function REPORT_BURNED_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) REPORTBURNEDTOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "REPORT_BURNED_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REPORTBURNEDTOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xd3095126.
//
// Solidity: function REPORT_BURNED_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) REPORTBURNEDTOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.REPORTBURNEDTOKENSREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// REPORTBURNEDTOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xd3095126.
//
// Solidity: function REPORT_BURNED_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) REPORTBURNEDTOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.REPORTBURNEDTOKENSREQUIREDGAS(&_NativeTokenDestination.CallOpts)
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

// CurrentReserveImbalance is a free data retrieval call binding the contract method 0x00d872ae.
//
// Solidity: function currentReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) CurrentReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "currentReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentReserveImbalance is a free data retrieval call binding the contract method 0x00d872ae.
//
// Solidity: function currentReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) CurrentReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.CurrentReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// CurrentReserveImbalance is a free data retrieval call binding the contract method 0x00d872ae.
//
// Solidity: function currentReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) CurrentReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.CurrentReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) InitialReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "initialReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.InitialReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.InitialReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) IsCollateralized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "isCollateralized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) IsCollateralized() (bool, error) {
	return _NativeTokenDestination.Contract.IsCollateralized(&_NativeTokenDestination.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) IsCollateralized() (bool, error) {
	return _NativeTokenDestination.Contract.IsCollateralized(&_NativeTokenDestination.CallOpts)
}

// NativeTokenSourceAddress is a free data retrieval call binding the contract method 0x5d93f9af.
//
// Solidity: function nativeTokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) NativeTokenSourceAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "nativeTokenSourceAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeTokenSourceAddress is a free data retrieval call binding the contract method 0x5d93f9af.
//
// Solidity: function nativeTokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) NativeTokenSourceAddress() (common.Address, error) {
	return _NativeTokenDestination.Contract.NativeTokenSourceAddress(&_NativeTokenDestination.CallOpts)
}

// NativeTokenSourceAddress is a free data retrieval call binding the contract method 0x5d93f9af.
//
// Solidity: function nativeTokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) NativeTokenSourceAddress() (common.Address, error) {
	return _NativeTokenDestination.Contract.NativeTokenSourceAddress(&_NativeTokenDestination.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCaller) SourceBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "sourceBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationSession) SourceBlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.SourceBlockchainID(&_NativeTokenDestination.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) SourceBlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.SourceBlockchainID(&_NativeTokenDestination.CallOpts)
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

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalMinted(&_NativeTokenDestination.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalMinted(&_NativeTokenDestination.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalSupply(&_NativeTokenDestination.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalSupply(&_NativeTokenDestination.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "receiveTeleporterMessage", senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReceiveTeleporterMessage(&_NativeTokenDestination.TransactOpts, senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReceiveTeleporterMessage(&_NativeTokenDestination.TransactOpts, senderBlockchainID, senderAddress, message)
}

// ReportTotalBurnedTxFees is a paid mutator transaction binding the contract method 0x3a94fe51.
//
// Solidity: function reportTotalBurnedTxFees((address,uint256) feeInfo, address[] allowedRelayerAddresses) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) ReportTotalBurnedTxFees(opts *bind.TransactOpts, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "reportTotalBurnedTxFees", feeInfo, allowedRelayerAddresses)
}

// ReportTotalBurnedTxFees is a paid mutator transaction binding the contract method 0x3a94fe51.
//
// Solidity: function reportTotalBurnedTxFees((address,uint256) feeInfo, address[] allowedRelayerAddresses) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) ReportTotalBurnedTxFees(feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReportTotalBurnedTxFees(&_NativeTokenDestination.TransactOpts, feeInfo, allowedRelayerAddresses)
}

// ReportTotalBurnedTxFees is a paid mutator transaction binding the contract method 0x3a94fe51.
//
// Solidity: function reportTotalBurnedTxFees((address,uint256) feeInfo, address[] allowedRelayerAddresses) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) ReportTotalBurnedTxFees(feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReportTotalBurnedTxFees(&_NativeTokenDestination.TransactOpts, feeInfo, allowedRelayerAddresses)
}

// TransferToSource is a paid mutator transaction binding the contract method 0x75846562.
//
// Solidity: function transferToSource(address recipient, (address,uint256) feeInfo, address[] allowedRelayerAddresses) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) TransferToSource(opts *bind.TransactOpts, recipient common.Address, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "transferToSource", recipient, feeInfo, allowedRelayerAddresses)
}

// TransferToSource is a paid mutator transaction binding the contract method 0x75846562.
//
// Solidity: function transferToSource(address recipient, (address,uint256) feeInfo, address[] allowedRelayerAddresses) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) TransferToSource(recipient common.Address, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferToSource(&_NativeTokenDestination.TransactOpts, recipient, feeInfo, allowedRelayerAddresses)
}

// TransferToSource is a paid mutator transaction binding the contract method 0x75846562.
//
// Solidity: function transferToSource(address recipient, (address,uint256) feeInfo, address[] allowedRelayerAddresses) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) TransferToSource(recipient common.Address, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferToSource(&_NativeTokenDestination.TransactOpts, recipient, feeInfo, allowedRelayerAddresses)
}

// NativeTokenDestinationCollateralAddedIterator is returned from FilterCollateralAdded and is used to iterate over the raw logs and unpacked data for CollateralAdded events raised by the NativeTokenDestination contract.
type NativeTokenDestinationCollateralAddedIterator struct {
	Event *NativeTokenDestinationCollateralAdded // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationCollateralAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationCollateralAdded)
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
		it.Event = new(NativeTokenDestinationCollateralAdded)
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
func (it *NativeTokenDestinationCollateralAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationCollateralAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationCollateralAdded represents a CollateralAdded event raised by the NativeTokenDestination contract.
type NativeTokenDestinationCollateralAdded struct {
	Amount    *big.Int
	Remaining *big.Int
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdded is a free log retrieval operation binding the contract event 0x66c12a5fefbc1f35dd64a1d9b069fc5bda187e6f68812c58a330b77b2a695ba9.
//
// Solidity: event CollateralAdded(uint256 amount, uint256 remaining, address sender)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterCollateralAdded(opts *bind.FilterOpts) (*NativeTokenDestinationCollateralAddedIterator, error) {

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "CollateralAdded")
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCollateralAddedIterator{contract: _NativeTokenDestination.contract, event: "CollateralAdded", logs: logs, sub: sub}, nil
}

// WatchCollateralAdded is a free log subscription operation binding the contract event 0x66c12a5fefbc1f35dd64a1d9b069fc5bda187e6f68812c58a330b77b2a695ba9.
//
// Solidity: event CollateralAdded(uint256 amount, uint256 remaining, address sender)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchCollateralAdded(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationCollateralAdded) (event.Subscription, error) {

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "CollateralAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationCollateralAdded)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
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

// ParseCollateralAdded is a log parse operation binding the contract event 0x66c12a5fefbc1f35dd64a1d9b069fc5bda187e6f68812c58a330b77b2a695ba9.
//
// Solidity: event CollateralAdded(uint256 amount, uint256 remaining, address sender)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseCollateralAdded(log types.Log) (*NativeTokenDestinationCollateralAdded, error) {
	event := new(NativeTokenDestinationCollateralAdded)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationNativeTokensMintedIterator is returned from FilterNativeTokensMinted and is used to iterate over the raw logs and unpacked data for NativeTokensMinted events raised by the NativeTokenDestination contract.
type NativeTokenDestinationNativeTokensMintedIterator struct {
	Event *NativeTokenDestinationNativeTokensMinted // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationNativeTokensMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationNativeTokensMinted)
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
		it.Event = new(NativeTokenDestinationNativeTokensMinted)
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
func (it *NativeTokenDestinationNativeTokensMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationNativeTokensMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationNativeTokensMinted represents a NativeTokensMinted event raised by the NativeTokenDestination contract.
type NativeTokenDestinationNativeTokensMinted struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNativeTokensMinted is a free log retrieval operation binding the contract event 0xd949ea0e9d5db53492d77f28fd5467fb2f6c4f5b88e3350e3c36729b76e99cf2.
//
// Solidity: event NativeTokensMinted(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterNativeTokensMinted(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenDestinationNativeTokensMintedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "NativeTokensMinted", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationNativeTokensMintedIterator{contract: _NativeTokenDestination.contract, event: "NativeTokensMinted", logs: logs, sub: sub}, nil
}

// WatchNativeTokensMinted is a free log subscription operation binding the contract event 0xd949ea0e9d5db53492d77f28fd5467fb2f6c4f5b88e3350e3c36729b76e99cf2.
//
// Solidity: event NativeTokensMinted(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchNativeTokensMinted(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationNativeTokensMinted, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "NativeTokensMinted", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationNativeTokensMinted)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "NativeTokensMinted", log); err != nil {
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

// ParseNativeTokensMinted is a log parse operation binding the contract event 0xd949ea0e9d5db53492d77f28fd5467fb2f6c4f5b88e3350e3c36729b76e99cf2.
//
// Solidity: event NativeTokensMinted(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseNativeTokensMinted(log types.Log) (*NativeTokenDestinationNativeTokensMinted, error) {
	event := new(NativeTokenDestinationNativeTokensMinted)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "NativeTokensMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationReportTotalBurnedTxFeesIterator is returned from FilterReportTotalBurnedTxFees and is used to iterate over the raw logs and unpacked data for ReportTotalBurnedTxFees events raised by the NativeTokenDestination contract.
type NativeTokenDestinationReportTotalBurnedTxFeesIterator struct {
	Event *NativeTokenDestinationReportTotalBurnedTxFees // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationReportTotalBurnedTxFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationReportTotalBurnedTxFees)
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
		it.Event = new(NativeTokenDestinationReportTotalBurnedTxFees)
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
func (it *NativeTokenDestinationReportTotalBurnedTxFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationReportTotalBurnedTxFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationReportTotalBurnedTxFees represents a ReportTotalBurnedTxFees event raised by the NativeTokenDestination contract.
type NativeTokenDestinationReportTotalBurnedTxFees struct {
	TeleporterMessageID *big.Int
	BurnAddressBalance  *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReportTotalBurnedTxFees is a free log retrieval operation binding the contract event 0x2550fa6041684d40e635e29e93dde9017d70c25b46aa88393317b5182ed6ae7c.
//
// Solidity: event ReportTotalBurnedTxFees(uint256 indexed teleporterMessageID, uint256 burnAddressBalance)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterReportTotalBurnedTxFees(opts *bind.FilterOpts, teleporterMessageID []*big.Int) (*NativeTokenDestinationReportTotalBurnedTxFeesIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "ReportTotalBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationReportTotalBurnedTxFeesIterator{contract: _NativeTokenDestination.contract, event: "ReportTotalBurnedTxFees", logs: logs, sub: sub}, nil
}

// WatchReportTotalBurnedTxFees is a free log subscription operation binding the contract event 0x2550fa6041684d40e635e29e93dde9017d70c25b46aa88393317b5182ed6ae7c.
//
// Solidity: event ReportTotalBurnedTxFees(uint256 indexed teleporterMessageID, uint256 burnAddressBalance)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchReportTotalBurnedTxFees(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationReportTotalBurnedTxFees, teleporterMessageID []*big.Int) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "ReportTotalBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationReportTotalBurnedTxFees)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "ReportTotalBurnedTxFees", log); err != nil {
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

// ParseReportTotalBurnedTxFees is a log parse operation binding the contract event 0x2550fa6041684d40e635e29e93dde9017d70c25b46aa88393317b5182ed6ae7c.
//
// Solidity: event ReportTotalBurnedTxFees(uint256 indexed teleporterMessageID, uint256 burnAddressBalance)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseReportTotalBurnedTxFees(log types.Log) (*NativeTokenDestinationReportTotalBurnedTxFees, error) {
	event := new(NativeTokenDestinationReportTotalBurnedTxFees)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "ReportTotalBurnedTxFees", log); err != nil {
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
	Sender              common.Address
	Recipient           common.Address
	TeleporterMessageID *big.Int
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTransferToSource is a free log retrieval operation binding the contract event 0x0322cbb1d3c23f6dbf1deddb3b4ef3ce0f93ae6eec7b44e4f395804104466d14.
//
// Solidity: event TransferToSource(address indexed sender, address indexed recipient, uint256 indexed teleporterMessageID, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTransferToSource(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, teleporterMessageID []*big.Int) (*NativeTokenDestinationTransferToSourceIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TransferToSource", senderRule, recipientRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTransferToSourceIterator{contract: _NativeTokenDestination.contract, event: "TransferToSource", logs: logs, sub: sub}, nil
}

// WatchTransferToSource is a free log subscription operation binding the contract event 0x0322cbb1d3c23f6dbf1deddb3b4ef3ce0f93ae6eec7b44e4f395804104466d14.
//
// Solidity: event TransferToSource(address indexed sender, address indexed recipient, uint256 indexed teleporterMessageID, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTransferToSource(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTransferToSource, sender []common.Address, recipient []common.Address, teleporterMessageID []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TransferToSource", senderRule, recipientRule, teleporterMessageIDRule)
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

// ParseTransferToSource is a log parse operation binding the contract event 0x0322cbb1d3c23f6dbf1deddb3b4ef3ce0f93ae6eec7b44e4f395804104466d14.
//
// Solidity: event TransferToSource(address indexed sender, address indexed recipient, uint256 indexed teleporterMessageID, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTransferToSource(log types.Log) (*NativeTokenDestinationTransferToSource, error) {
	event := new(NativeTokenDestinationTransferToSource)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TransferToSource", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
