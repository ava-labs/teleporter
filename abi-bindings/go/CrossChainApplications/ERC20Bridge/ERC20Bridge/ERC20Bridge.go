// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20bridge

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

// ERC20BridgeMetaData contains all meta data concerning the ERC20Bridge contract.
var ERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nativeChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nativeBridgeAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nativeContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeTokenAddress\",\"type\":\"address\"}],\"name\":\"CreateBridgeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MintBridgeTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nativeContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"}],\"name\":\"SubmitCreateBridgeToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CREATE_BRIDGE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_BRIDGE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_PRECOMPILE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"primaryFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFeeAmount\",\"type\":\"uint256\"}],\"name\":\"bridgeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bridgedBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nativeContractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nativeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nativeSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"nativeDecimals\",\"type\":\"uint8\"}],\"name\":\"encodeCreateBridgeTokenData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nativeContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bridgeAmount\",\"type\":\"uint256\"}],\"name\":\"encodeMintBridgeTokensData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nativeContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"encodeTransferBridgeTokensData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nativeToWrappedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nativeChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"nativeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageFeeAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"messageFeeAmount\",\"type\":\"uint256\"}],\"name\":\"submitCreateBridgeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"submittedBridgeTokenCreations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenContracts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC20BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20BridgeMetaData.ABI instead.
var ERC20BridgeABI = ERC20BridgeMetaData.ABI

// ERC20Bridge is an auto generated Go binding around an Ethereum contract.
type ERC20Bridge struct {
	ERC20BridgeCaller     // Read-only binding to the contract
	ERC20BridgeTransactor // Write-only binding to the contract
	ERC20BridgeFilterer   // Log filterer for contract events
}

// ERC20BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BridgeSession struct {
	Contract     *ERC20Bridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BridgeCallerSession struct {
	Contract *ERC20BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BridgeTransactorSession struct {
	Contract     *ERC20BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BridgeRaw struct {
	Contract *ERC20Bridge // Generic contract binding to access the raw methods on
}

// ERC20BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BridgeCallerRaw struct {
	Contract *ERC20BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BridgeTransactorRaw struct {
	Contract *ERC20BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Bridge creates a new instance of ERC20Bridge, bound to a specific deployed contract.
func NewERC20Bridge(address common.Address, backend bind.ContractBackend) (*ERC20Bridge, error) {
	contract, err := bindERC20Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Bridge{ERC20BridgeCaller: ERC20BridgeCaller{contract: contract}, ERC20BridgeTransactor: ERC20BridgeTransactor{contract: contract}, ERC20BridgeFilterer: ERC20BridgeFilterer{contract: contract}}, nil
}

// NewERC20BridgeCaller creates a new read-only instance of ERC20Bridge, bound to a specific deployed contract.
func NewERC20BridgeCaller(address common.Address, caller bind.ContractCaller) (*ERC20BridgeCaller, error) {
	contract, err := bindERC20Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeCaller{contract: contract}, nil
}

// NewERC20BridgeTransactor creates a new write-only instance of ERC20Bridge, bound to a specific deployed contract.
func NewERC20BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BridgeTransactor, error) {
	contract, err := bindERC20Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeTransactor{contract: contract}, nil
}

// NewERC20BridgeFilterer creates a new log filterer instance of ERC20Bridge, bound to a specific deployed contract.
func NewERC20BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BridgeFilterer, error) {
	contract, err := bindERC20Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeFilterer{contract: contract}, nil
}

// bindERC20Bridge binds a generic wrapper to an already deployed contract.
func bindERC20Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Bridge *ERC20BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Bridge.Contract.ERC20BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Bridge *ERC20BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.ERC20BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Bridge *ERC20BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.ERC20BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Bridge *ERC20BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Bridge *ERC20BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Bridge *ERC20BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.contract.Transact(opts, method, params...)
}

// CREATEBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x5f217bcc.
//
// Solidity: function CREATE_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCaller) CREATEBRIDGETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "CREATE_BRIDGE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CREATEBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x5f217bcc.
//
// Solidity: function CREATE_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeSession) CREATEBRIDGETOKENSREQUIREDGAS() (*big.Int, error) {
	return _ERC20Bridge.Contract.CREATEBRIDGETOKENSREQUIREDGAS(&_ERC20Bridge.CallOpts)
}

// CREATEBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x5f217bcc.
//
// Solidity: function CREATE_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCallerSession) CREATEBRIDGETOKENSREQUIREDGAS() (*big.Int, error) {
	return _ERC20Bridge.Contract.CREATEBRIDGETOKENSREQUIREDGAS(&_ERC20Bridge.CallOpts)
}

// MINTBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x7a465fd9.
//
// Solidity: function MINT_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCaller) MINTBRIDGETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "MINT_BRIDGE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x7a465fd9.
//
// Solidity: function MINT_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeSession) MINTBRIDGETOKENSREQUIREDGAS() (*big.Int, error) {
	return _ERC20Bridge.Contract.MINTBRIDGETOKENSREQUIREDGAS(&_ERC20Bridge.CallOpts)
}

// MINTBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x7a465fd9.
//
// Solidity: function MINT_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCallerSession) MINTBRIDGETOKENSREQUIREDGAS() (*big.Int, error) {
	return _ERC20Bridge.Contract.MINTBRIDGETOKENSREQUIREDGAS(&_ERC20Bridge.CallOpts)
}

// TRANSFERBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x6b47cd9a.
//
// Solidity: function TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCaller) TRANSFERBRIDGETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TRANSFERBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x6b47cd9a.
//
// Solidity: function TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeSession) TRANSFERBRIDGETOKENSREQUIREDGAS() (*big.Int, error) {
	return _ERC20Bridge.Contract.TRANSFERBRIDGETOKENSREQUIREDGAS(&_ERC20Bridge.CallOpts)
}

// TRANSFERBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x6b47cd9a.
//
// Solidity: function TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCallerSession) TRANSFERBRIDGETOKENSREQUIREDGAS() (*big.Int, error) {
	return _ERC20Bridge.Contract.TRANSFERBRIDGETOKENSREQUIREDGAS(&_ERC20Bridge.CallOpts)
}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_ERC20Bridge *ERC20BridgeCaller) WARPPRECOMPILEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "WARP_PRECOMPILE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_ERC20Bridge *ERC20BridgeSession) WARPPRECOMPILEADDRESS() (common.Address, error) {
	return _ERC20Bridge.Contract.WARPPRECOMPILEADDRESS(&_ERC20Bridge.CallOpts)
}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_ERC20Bridge *ERC20BridgeCallerSession) WARPPRECOMPILEADDRESS() (common.Address, error) {
	return _ERC20Bridge.Contract.WARPPRECOMPILEADDRESS(&_ERC20Bridge.CallOpts)
}

// BridgedBalances is a free data retrieval call binding the contract method 0xb9e55da1.
//
// Solidity: function bridgedBalances(bytes32 , address , address ) view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCaller) BridgedBalances(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "bridgedBalances", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgedBalances is a free data retrieval call binding the contract method 0xb9e55da1.
//
// Solidity: function bridgedBalances(bytes32 , address , address ) view returns(uint256)
func (_ERC20Bridge *ERC20BridgeSession) BridgedBalances(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _ERC20Bridge.Contract.BridgedBalances(&_ERC20Bridge.CallOpts, arg0, arg1, arg2)
}

// BridgedBalances is a free data retrieval call binding the contract method 0xb9e55da1.
//
// Solidity: function bridgedBalances(bytes32 , address , address ) view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCallerSession) BridgedBalances(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _ERC20Bridge.Contract.BridgedBalances(&_ERC20Bridge.CallOpts, arg0, arg1, arg2)
}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_ERC20Bridge *ERC20BridgeCaller) CurrentChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "currentChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_ERC20Bridge *ERC20BridgeSession) CurrentChainID() ([32]byte, error) {
	return _ERC20Bridge.Contract.CurrentChainID(&_ERC20Bridge.CallOpts)
}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_ERC20Bridge *ERC20BridgeCallerSession) CurrentChainID() ([32]byte, error) {
	return _ERC20Bridge.Contract.CurrentChainID(&_ERC20Bridge.CallOpts)
}

// EncodeCreateBridgeTokenData is a free data retrieval call binding the contract method 0x367e9584.
//
// Solidity: function encodeCreateBridgeTokenData(address nativeContractAddress, string nativeName, string nativeSymbol, uint8 nativeDecimals) pure returns(bytes)
func (_ERC20Bridge *ERC20BridgeCaller) EncodeCreateBridgeTokenData(opts *bind.CallOpts, nativeContractAddress common.Address, nativeName string, nativeSymbol string, nativeDecimals uint8) ([]byte, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "encodeCreateBridgeTokenData", nativeContractAddress, nativeName, nativeSymbol, nativeDecimals)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeCreateBridgeTokenData is a free data retrieval call binding the contract method 0x367e9584.
//
// Solidity: function encodeCreateBridgeTokenData(address nativeContractAddress, string nativeName, string nativeSymbol, uint8 nativeDecimals) pure returns(bytes)
func (_ERC20Bridge *ERC20BridgeSession) EncodeCreateBridgeTokenData(nativeContractAddress common.Address, nativeName string, nativeSymbol string, nativeDecimals uint8) ([]byte, error) {
	return _ERC20Bridge.Contract.EncodeCreateBridgeTokenData(&_ERC20Bridge.CallOpts, nativeContractAddress, nativeName, nativeSymbol, nativeDecimals)
}

// EncodeCreateBridgeTokenData is a free data retrieval call binding the contract method 0x367e9584.
//
// Solidity: function encodeCreateBridgeTokenData(address nativeContractAddress, string nativeName, string nativeSymbol, uint8 nativeDecimals) pure returns(bytes)
func (_ERC20Bridge *ERC20BridgeCallerSession) EncodeCreateBridgeTokenData(nativeContractAddress common.Address, nativeName string, nativeSymbol string, nativeDecimals uint8) ([]byte, error) {
	return _ERC20Bridge.Contract.EncodeCreateBridgeTokenData(&_ERC20Bridge.CallOpts, nativeContractAddress, nativeName, nativeSymbol, nativeDecimals)
}

// EncodeMintBridgeTokensData is a free data retrieval call binding the contract method 0x8c56fcf0.
//
// Solidity: function encodeMintBridgeTokensData(address nativeContractAddress, address recipient, uint256 bridgeAmount) pure returns(bytes)
func (_ERC20Bridge *ERC20BridgeCaller) EncodeMintBridgeTokensData(opts *bind.CallOpts, nativeContractAddress common.Address, recipient common.Address, bridgeAmount *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "encodeMintBridgeTokensData", nativeContractAddress, recipient, bridgeAmount)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeMintBridgeTokensData is a free data retrieval call binding the contract method 0x8c56fcf0.
//
// Solidity: function encodeMintBridgeTokensData(address nativeContractAddress, address recipient, uint256 bridgeAmount) pure returns(bytes)
func (_ERC20Bridge *ERC20BridgeSession) EncodeMintBridgeTokensData(nativeContractAddress common.Address, recipient common.Address, bridgeAmount *big.Int) ([]byte, error) {
	return _ERC20Bridge.Contract.EncodeMintBridgeTokensData(&_ERC20Bridge.CallOpts, nativeContractAddress, recipient, bridgeAmount)
}

// EncodeMintBridgeTokensData is a free data retrieval call binding the contract method 0x8c56fcf0.
//
// Solidity: function encodeMintBridgeTokensData(address nativeContractAddress, address recipient, uint256 bridgeAmount) pure returns(bytes)
func (_ERC20Bridge *ERC20BridgeCallerSession) EncodeMintBridgeTokensData(nativeContractAddress common.Address, recipient common.Address, bridgeAmount *big.Int) ([]byte, error) {
	return _ERC20Bridge.Contract.EncodeMintBridgeTokensData(&_ERC20Bridge.CallOpts, nativeContractAddress, recipient, bridgeAmount)
}

// EncodeTransferBridgeTokensData is a free data retrieval call binding the contract method 0xc60da612.
//
// Solidity: function encodeTransferBridgeTokensData(bytes32 destinationChainID, address destinationBridgeAddress, address nativeContractAddress, address recipient, uint256 amount, uint256 feeAmount) pure returns(bytes)
func (_ERC20Bridge *ERC20BridgeCaller) EncodeTransferBridgeTokensData(opts *bind.CallOpts, destinationChainID [32]byte, destinationBridgeAddress common.Address, nativeContractAddress common.Address, recipient common.Address, amount *big.Int, feeAmount *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "encodeTransferBridgeTokensData", destinationChainID, destinationBridgeAddress, nativeContractAddress, recipient, amount, feeAmount)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTransferBridgeTokensData is a free data retrieval call binding the contract method 0xc60da612.
//
// Solidity: function encodeTransferBridgeTokensData(bytes32 destinationChainID, address destinationBridgeAddress, address nativeContractAddress, address recipient, uint256 amount, uint256 feeAmount) pure returns(bytes)
func (_ERC20Bridge *ERC20BridgeSession) EncodeTransferBridgeTokensData(destinationChainID [32]byte, destinationBridgeAddress common.Address, nativeContractAddress common.Address, recipient common.Address, amount *big.Int, feeAmount *big.Int) ([]byte, error) {
	return _ERC20Bridge.Contract.EncodeTransferBridgeTokensData(&_ERC20Bridge.CallOpts, destinationChainID, destinationBridgeAddress, nativeContractAddress, recipient, amount, feeAmount)
}

// EncodeTransferBridgeTokensData is a free data retrieval call binding the contract method 0xc60da612.
//
// Solidity: function encodeTransferBridgeTokensData(bytes32 destinationChainID, address destinationBridgeAddress, address nativeContractAddress, address recipient, uint256 amount, uint256 feeAmount) pure returns(bytes)
func (_ERC20Bridge *ERC20BridgeCallerSession) EncodeTransferBridgeTokensData(destinationChainID [32]byte, destinationBridgeAddress common.Address, nativeContractAddress common.Address, recipient common.Address, amount *big.Int, feeAmount *big.Int) ([]byte, error) {
	return _ERC20Bridge.Contract.EncodeTransferBridgeTokensData(&_ERC20Bridge.CallOpts, destinationChainID, destinationBridgeAddress, nativeContractAddress, recipient, amount, feeAmount)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ERC20Bridge.Contract.GetMinTeleporterVersion(&_ERC20Bridge.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20Bridge *ERC20BridgeCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ERC20Bridge.Contract.GetMinTeleporterVersion(&_ERC20Bridge.CallOpts)
}

// NativeToWrappedTokens is a free data retrieval call binding the contract method 0x65435568.
//
// Solidity: function nativeToWrappedTokens(bytes32 , address , address ) view returns(address)
func (_ERC20Bridge *ERC20BridgeCaller) NativeToWrappedTokens(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address, arg2 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "nativeToWrappedTokens", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToWrappedTokens is a free data retrieval call binding the contract method 0x65435568.
//
// Solidity: function nativeToWrappedTokens(bytes32 , address , address ) view returns(address)
func (_ERC20Bridge *ERC20BridgeSession) NativeToWrappedTokens(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (common.Address, error) {
	return _ERC20Bridge.Contract.NativeToWrappedTokens(&_ERC20Bridge.CallOpts, arg0, arg1, arg2)
}

// NativeToWrappedTokens is a free data retrieval call binding the contract method 0x65435568.
//
// Solidity: function nativeToWrappedTokens(bytes32 , address , address ) view returns(address)
func (_ERC20Bridge *ERC20BridgeCallerSession) NativeToWrappedTokens(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (common.Address, error) {
	return _ERC20Bridge.Contract.NativeToWrappedTokens(&_ERC20Bridge.CallOpts, arg0, arg1, arg2)
}

// SubmittedBridgeTokenCreations is a free data retrieval call binding the contract method 0x8343f661.
//
// Solidity: function submittedBridgeTokenCreations(bytes32 , address , address ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCaller) SubmittedBridgeTokenCreations(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "submittedBridgeTokenCreations", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SubmittedBridgeTokenCreations is a free data retrieval call binding the contract method 0x8343f661.
//
// Solidity: function submittedBridgeTokenCreations(bytes32 , address , address ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeSession) SubmittedBridgeTokenCreations(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (bool, error) {
	return _ERC20Bridge.Contract.SubmittedBridgeTokenCreations(&_ERC20Bridge.CallOpts, arg0, arg1, arg2)
}

// SubmittedBridgeTokenCreations is a free data retrieval call binding the contract method 0x8343f661.
//
// Solidity: function submittedBridgeTokenCreations(bytes32 , address , address ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCallerSession) SubmittedBridgeTokenCreations(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (bool, error) {
	return _ERC20Bridge.Contract.SubmittedBridgeTokenCreations(&_ERC20Bridge.CallOpts, arg0, arg1, arg2)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ERC20Bridge *ERC20BridgeCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ERC20Bridge *ERC20BridgeSession) TeleporterRegistry() (common.Address, error) {
	return _ERC20Bridge.Contract.TeleporterRegistry(&_ERC20Bridge.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ERC20Bridge *ERC20BridgeCallerSession) TeleporterRegistry() (common.Address, error) {
	return _ERC20Bridge.Contract.TeleporterRegistry(&_ERC20Bridge.CallOpts)
}

// WrappedTokenContracts is a free data retrieval call binding the contract method 0x9bd9abc0.
//
// Solidity: function wrappedTokenContracts(address ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCaller) WrappedTokenContracts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Bridge.contract.Call(opts, &out, "wrappedTokenContracts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WrappedTokenContracts is a free data retrieval call binding the contract method 0x9bd9abc0.
//
// Solidity: function wrappedTokenContracts(address ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeSession) WrappedTokenContracts(arg0 common.Address) (bool, error) {
	return _ERC20Bridge.Contract.WrappedTokenContracts(&_ERC20Bridge.CallOpts, arg0)
}

// WrappedTokenContracts is a free data retrieval call binding the contract method 0x9bd9abc0.
//
// Solidity: function wrappedTokenContracts(address ) view returns(bool)
func (_ERC20Bridge *ERC20BridgeCallerSession) WrappedTokenContracts(arg0 common.Address) (bool, error) {
	return _ERC20Bridge.Contract.WrappedTokenContracts(&_ERC20Bridge.CallOpts, arg0)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0xc63d2207.
//
// Solidity: function bridgeTokens(bytes32 destinationChainID, address destinationBridgeAddress, address tokenContractAddress, address recipient, uint256 totalAmount, uint256 primaryFeeAmount, uint256 secondaryFeeAmount) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) BridgeTokens(opts *bind.TransactOpts, destinationChainID [32]byte, destinationBridgeAddress common.Address, tokenContractAddress common.Address, recipient common.Address, totalAmount *big.Int, primaryFeeAmount *big.Int, secondaryFeeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "bridgeTokens", destinationChainID, destinationBridgeAddress, tokenContractAddress, recipient, totalAmount, primaryFeeAmount, secondaryFeeAmount)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0xc63d2207.
//
// Solidity: function bridgeTokens(bytes32 destinationChainID, address destinationBridgeAddress, address tokenContractAddress, address recipient, uint256 totalAmount, uint256 primaryFeeAmount, uint256 secondaryFeeAmount) returns()
func (_ERC20Bridge *ERC20BridgeSession) BridgeTokens(destinationChainID [32]byte, destinationBridgeAddress common.Address, tokenContractAddress common.Address, recipient common.Address, totalAmount *big.Int, primaryFeeAmount *big.Int, secondaryFeeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.BridgeTokens(&_ERC20Bridge.TransactOpts, destinationChainID, destinationBridgeAddress, tokenContractAddress, recipient, totalAmount, primaryFeeAmount, secondaryFeeAmount)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0xc63d2207.
//
// Solidity: function bridgeTokens(bytes32 destinationChainID, address destinationBridgeAddress, address tokenContractAddress, address recipient, uint256 totalAmount, uint256 primaryFeeAmount, uint256 secondaryFeeAmount) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) BridgeTokens(destinationChainID [32]byte, destinationBridgeAddress common.Address, tokenContractAddress common.Address, recipient common.Address, totalAmount *big.Int, primaryFeeAmount *big.Int, secondaryFeeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.BridgeTokens(&_ERC20Bridge.TransactOpts, destinationChainID, destinationBridgeAddress, tokenContractAddress, recipient, totalAmount, primaryFeeAmount, secondaryFeeAmount)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "receiveTeleporterMessage", nativeChainID, nativeBridgeAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_ERC20Bridge *ERC20BridgeSession) ReceiveTeleporterMessage(nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.ReceiveTeleporterMessage(&_ERC20Bridge.TransactOpts, nativeChainID, nativeBridgeAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) ReceiveTeleporterMessage(nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.ReceiveTeleporterMessage(&_ERC20Bridge.TransactOpts, nativeChainID, nativeBridgeAddress, message)
}

// SubmitCreateBridgeToken is a paid mutator transaction binding the contract method 0x6c7e40d1.
//
// Solidity: function submitCreateBridgeToken(bytes32 destinationChainID, address destinationBridgeAddress, address nativeToken, address messageFeeAsset, uint256 messageFeeAmount) returns()
func (_ERC20Bridge *ERC20BridgeTransactor) SubmitCreateBridgeToken(opts *bind.TransactOpts, destinationChainID [32]byte, destinationBridgeAddress common.Address, nativeToken common.Address, messageFeeAsset common.Address, messageFeeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "submitCreateBridgeToken", destinationChainID, destinationBridgeAddress, nativeToken, messageFeeAsset, messageFeeAmount)
}

// SubmitCreateBridgeToken is a paid mutator transaction binding the contract method 0x6c7e40d1.
//
// Solidity: function submitCreateBridgeToken(bytes32 destinationChainID, address destinationBridgeAddress, address nativeToken, address messageFeeAsset, uint256 messageFeeAmount) returns()
func (_ERC20Bridge *ERC20BridgeSession) SubmitCreateBridgeToken(destinationChainID [32]byte, destinationBridgeAddress common.Address, nativeToken common.Address, messageFeeAsset common.Address, messageFeeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.SubmitCreateBridgeToken(&_ERC20Bridge.TransactOpts, destinationChainID, destinationBridgeAddress, nativeToken, messageFeeAsset, messageFeeAmount)
}

// SubmitCreateBridgeToken is a paid mutator transaction binding the contract method 0x6c7e40d1.
//
// Solidity: function submitCreateBridgeToken(bytes32 destinationChainID, address destinationBridgeAddress, address nativeToken, address messageFeeAsset, uint256 messageFeeAmount) returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) SubmitCreateBridgeToken(destinationChainID [32]byte, destinationBridgeAddress common.Address, nativeToken common.Address, messageFeeAsset common.Address, messageFeeAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Bridge.Contract.SubmitCreateBridgeToken(&_ERC20Bridge.TransactOpts, destinationChainID, destinationBridgeAddress, nativeToken, messageFeeAsset, messageFeeAmount)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0xb6109d9d.
//
// Solidity: function updateMinTeleporterVersion() returns()
func (_ERC20Bridge *ERC20BridgeTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Bridge.contract.Transact(opts, "updateMinTeleporterVersion")
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0xb6109d9d.
//
// Solidity: function updateMinTeleporterVersion() returns()
func (_ERC20Bridge *ERC20BridgeSession) UpdateMinTeleporterVersion() (*types.Transaction, error) {
	return _ERC20Bridge.Contract.UpdateMinTeleporterVersion(&_ERC20Bridge.TransactOpts)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0xb6109d9d.
//
// Solidity: function updateMinTeleporterVersion() returns()
func (_ERC20Bridge *ERC20BridgeTransactorSession) UpdateMinTeleporterVersion() (*types.Transaction, error) {
	return _ERC20Bridge.Contract.UpdateMinTeleporterVersion(&_ERC20Bridge.TransactOpts)
}

// ERC20BridgeBridgeTokensIterator is returned from FilterBridgeTokens and is used to iterate over the raw logs and unpacked data for BridgeTokens events raised by the ERC20Bridge contract.
type ERC20BridgeBridgeTokensIterator struct {
	Event *ERC20BridgeBridgeTokens // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeBridgeTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeBridgeTokens)
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
		it.Event = new(ERC20BridgeBridgeTokens)
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
func (it *ERC20BridgeBridgeTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeBridgeTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeBridgeTokens represents a BridgeTokens event raised by the ERC20Bridge contract.
type ERC20BridgeBridgeTokens struct {
	TokenContractAddress     common.Address
	DestinationChainID       [32]byte
	TeleporterMessageID      *big.Int
	DestinationBridgeAddress common.Address
	Recipient                common.Address
	Amount                   *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterBridgeTokens is a free log retrieval operation binding the contract event 0x97935c4470efae40c8440c3abfe968a5512232dd375cc974e712f487c2b99c31.
//
// Solidity: event BridgeTokens(address indexed tokenContractAddress, bytes32 indexed destinationChainID, uint256 indexed teleporterMessageID, address destinationBridgeAddress, address recipient, uint256 amount)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterBridgeTokens(opts *bind.FilterOpts, tokenContractAddress []common.Address, destinationChainID [][32]byte, teleporterMessageID []*big.Int) (*ERC20BridgeBridgeTokensIterator, error) {

	var tokenContractAddressRule []interface{}
	for _, tokenContractAddressItem := range tokenContractAddress {
		tokenContractAddressRule = append(tokenContractAddressRule, tokenContractAddressItem)
	}
	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "BridgeTokens", tokenContractAddressRule, destinationChainIDRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeBridgeTokensIterator{contract: _ERC20Bridge.contract, event: "BridgeTokens", logs: logs, sub: sub}, nil
}

// WatchBridgeTokens is a free log subscription operation binding the contract event 0x97935c4470efae40c8440c3abfe968a5512232dd375cc974e712f487c2b99c31.
//
// Solidity: event BridgeTokens(address indexed tokenContractAddress, bytes32 indexed destinationChainID, uint256 indexed teleporterMessageID, address destinationBridgeAddress, address recipient, uint256 amount)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchBridgeTokens(opts *bind.WatchOpts, sink chan<- *ERC20BridgeBridgeTokens, tokenContractAddress []common.Address, destinationChainID [][32]byte, teleporterMessageID []*big.Int) (event.Subscription, error) {

	var tokenContractAddressRule []interface{}
	for _, tokenContractAddressItem := range tokenContractAddress {
		tokenContractAddressRule = append(tokenContractAddressRule, tokenContractAddressItem)
	}
	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "BridgeTokens", tokenContractAddressRule, destinationChainIDRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeBridgeTokens)
				if err := _ERC20Bridge.contract.UnpackLog(event, "BridgeTokens", log); err != nil {
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

// ParseBridgeTokens is a log parse operation binding the contract event 0x97935c4470efae40c8440c3abfe968a5512232dd375cc974e712f487c2b99c31.
//
// Solidity: event BridgeTokens(address indexed tokenContractAddress, bytes32 indexed destinationChainID, uint256 indexed teleporterMessageID, address destinationBridgeAddress, address recipient, uint256 amount)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseBridgeTokens(log types.Log) (*ERC20BridgeBridgeTokens, error) {
	event := new(ERC20BridgeBridgeTokens)
	if err := _ERC20Bridge.contract.UnpackLog(event, "BridgeTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeCreateBridgeTokenIterator is returned from FilterCreateBridgeToken and is used to iterate over the raw logs and unpacked data for CreateBridgeToken events raised by the ERC20Bridge contract.
type ERC20BridgeCreateBridgeTokenIterator struct {
	Event *ERC20BridgeCreateBridgeToken // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeCreateBridgeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeCreateBridgeToken)
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
		it.Event = new(ERC20BridgeCreateBridgeToken)
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
func (it *ERC20BridgeCreateBridgeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeCreateBridgeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeCreateBridgeToken represents a CreateBridgeToken event raised by the ERC20Bridge contract.
type ERC20BridgeCreateBridgeToken struct {
	NativeChainID         [32]byte
	NativeBridgeAddress   common.Address
	NativeContractAddress common.Address
	BridgeTokenAddress    common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCreateBridgeToken is a free log retrieval operation binding the contract event 0xe1c61a845f79534e11924517ddbedc668d0c20e467eafb4d3bd2858e2815f3b5.
//
// Solidity: event CreateBridgeToken(bytes32 indexed nativeChainID, address indexed nativeBridgeAddress, address indexed nativeContractAddress, address bridgeTokenAddress)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterCreateBridgeToken(opts *bind.FilterOpts, nativeChainID [][32]byte, nativeBridgeAddress []common.Address, nativeContractAddress []common.Address) (*ERC20BridgeCreateBridgeTokenIterator, error) {

	var nativeChainIDRule []interface{}
	for _, nativeChainIDItem := range nativeChainID {
		nativeChainIDRule = append(nativeChainIDRule, nativeChainIDItem)
	}
	var nativeBridgeAddressRule []interface{}
	for _, nativeBridgeAddressItem := range nativeBridgeAddress {
		nativeBridgeAddressRule = append(nativeBridgeAddressRule, nativeBridgeAddressItem)
	}
	var nativeContractAddressRule []interface{}
	for _, nativeContractAddressItem := range nativeContractAddress {
		nativeContractAddressRule = append(nativeContractAddressRule, nativeContractAddressItem)
	}

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "CreateBridgeToken", nativeChainIDRule, nativeBridgeAddressRule, nativeContractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeCreateBridgeTokenIterator{contract: _ERC20Bridge.contract, event: "CreateBridgeToken", logs: logs, sub: sub}, nil
}

// WatchCreateBridgeToken is a free log subscription operation binding the contract event 0xe1c61a845f79534e11924517ddbedc668d0c20e467eafb4d3bd2858e2815f3b5.
//
// Solidity: event CreateBridgeToken(bytes32 indexed nativeChainID, address indexed nativeBridgeAddress, address indexed nativeContractAddress, address bridgeTokenAddress)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchCreateBridgeToken(opts *bind.WatchOpts, sink chan<- *ERC20BridgeCreateBridgeToken, nativeChainID [][32]byte, nativeBridgeAddress []common.Address, nativeContractAddress []common.Address) (event.Subscription, error) {

	var nativeChainIDRule []interface{}
	for _, nativeChainIDItem := range nativeChainID {
		nativeChainIDRule = append(nativeChainIDRule, nativeChainIDItem)
	}
	var nativeBridgeAddressRule []interface{}
	for _, nativeBridgeAddressItem := range nativeBridgeAddress {
		nativeBridgeAddressRule = append(nativeBridgeAddressRule, nativeBridgeAddressItem)
	}
	var nativeContractAddressRule []interface{}
	for _, nativeContractAddressItem := range nativeContractAddress {
		nativeContractAddressRule = append(nativeContractAddressRule, nativeContractAddressItem)
	}

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "CreateBridgeToken", nativeChainIDRule, nativeBridgeAddressRule, nativeContractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeCreateBridgeToken)
				if err := _ERC20Bridge.contract.UnpackLog(event, "CreateBridgeToken", log); err != nil {
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

// ParseCreateBridgeToken is a log parse operation binding the contract event 0xe1c61a845f79534e11924517ddbedc668d0c20e467eafb4d3bd2858e2815f3b5.
//
// Solidity: event CreateBridgeToken(bytes32 indexed nativeChainID, address indexed nativeBridgeAddress, address indexed nativeContractAddress, address bridgeTokenAddress)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseCreateBridgeToken(log types.Log) (*ERC20BridgeCreateBridgeToken, error) {
	event := new(ERC20BridgeCreateBridgeToken)
	if err := _ERC20Bridge.contract.UnpackLog(event, "CreateBridgeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeMintBridgeTokensIterator is returned from FilterMintBridgeTokens and is used to iterate over the raw logs and unpacked data for MintBridgeTokens events raised by the ERC20Bridge contract.
type ERC20BridgeMintBridgeTokensIterator struct {
	Event *ERC20BridgeMintBridgeTokens // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeMintBridgeTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeMintBridgeTokens)
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
		it.Event = new(ERC20BridgeMintBridgeTokens)
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
func (it *ERC20BridgeMintBridgeTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeMintBridgeTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeMintBridgeTokens represents a MintBridgeTokens event raised by the ERC20Bridge contract.
type ERC20BridgeMintBridgeTokens struct {
	ContractAddress common.Address
	Recipient       common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintBridgeTokens is a free log retrieval operation binding the contract event 0xc0767f158f0d5394b598489a51ed607cd55a8be2dcef113ba1626efcf4c63954.
//
// Solidity: event MintBridgeTokens(address indexed contractAddress, address recipient, uint256 amount)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterMintBridgeTokens(opts *bind.FilterOpts, contractAddress []common.Address) (*ERC20BridgeMintBridgeTokensIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "MintBridgeTokens", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeMintBridgeTokensIterator{contract: _ERC20Bridge.contract, event: "MintBridgeTokens", logs: logs, sub: sub}, nil
}

// WatchMintBridgeTokens is a free log subscription operation binding the contract event 0xc0767f158f0d5394b598489a51ed607cd55a8be2dcef113ba1626efcf4c63954.
//
// Solidity: event MintBridgeTokens(address indexed contractAddress, address recipient, uint256 amount)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchMintBridgeTokens(opts *bind.WatchOpts, sink chan<- *ERC20BridgeMintBridgeTokens, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "MintBridgeTokens", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeMintBridgeTokens)
				if err := _ERC20Bridge.contract.UnpackLog(event, "MintBridgeTokens", log); err != nil {
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

// ParseMintBridgeTokens is a log parse operation binding the contract event 0xc0767f158f0d5394b598489a51ed607cd55a8be2dcef113ba1626efcf4c63954.
//
// Solidity: event MintBridgeTokens(address indexed contractAddress, address recipient, uint256 amount)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseMintBridgeTokens(log types.Log) (*ERC20BridgeMintBridgeTokens, error) {
	event := new(ERC20BridgeMintBridgeTokens)
	if err := _ERC20Bridge.contract.UnpackLog(event, "MintBridgeTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BridgeSubmitCreateBridgeTokenIterator is returned from FilterSubmitCreateBridgeToken and is used to iterate over the raw logs and unpacked data for SubmitCreateBridgeToken events raised by the ERC20Bridge contract.
type ERC20BridgeSubmitCreateBridgeTokenIterator struct {
	Event *ERC20BridgeSubmitCreateBridgeToken // Event containing the contract specifics and raw log

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
func (it *ERC20BridgeSubmitCreateBridgeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BridgeSubmitCreateBridgeToken)
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
		it.Event = new(ERC20BridgeSubmitCreateBridgeToken)
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
func (it *ERC20BridgeSubmitCreateBridgeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BridgeSubmitCreateBridgeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BridgeSubmitCreateBridgeToken represents a SubmitCreateBridgeToken event raised by the ERC20Bridge contract.
type ERC20BridgeSubmitCreateBridgeToken struct {
	DestinationChainID       [32]byte
	DestinationBridgeAddress common.Address
	NativeContractAddress    common.Address
	TeleporterMessageID      *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSubmitCreateBridgeToken is a free log retrieval operation binding the contract event 0x110b902745a3d7d6b66732479f01de654a3bc6e501be7c8ba2c3a6f9868cb539.
//
// Solidity: event SubmitCreateBridgeToken(bytes32 indexed destinationChainID, address indexed destinationBridgeAddress, address indexed nativeContractAddress, uint256 teleporterMessageID)
func (_ERC20Bridge *ERC20BridgeFilterer) FilterSubmitCreateBridgeToken(opts *bind.FilterOpts, destinationChainID [][32]byte, destinationBridgeAddress []common.Address, nativeContractAddress []common.Address) (*ERC20BridgeSubmitCreateBridgeTokenIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var destinationBridgeAddressRule []interface{}
	for _, destinationBridgeAddressItem := range destinationBridgeAddress {
		destinationBridgeAddressRule = append(destinationBridgeAddressRule, destinationBridgeAddressItem)
	}
	var nativeContractAddressRule []interface{}
	for _, nativeContractAddressItem := range nativeContractAddress {
		nativeContractAddressRule = append(nativeContractAddressRule, nativeContractAddressItem)
	}

	logs, sub, err := _ERC20Bridge.contract.FilterLogs(opts, "SubmitCreateBridgeToken", destinationChainIDRule, destinationBridgeAddressRule, nativeContractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BridgeSubmitCreateBridgeTokenIterator{contract: _ERC20Bridge.contract, event: "SubmitCreateBridgeToken", logs: logs, sub: sub}, nil
}

// WatchSubmitCreateBridgeToken is a free log subscription operation binding the contract event 0x110b902745a3d7d6b66732479f01de654a3bc6e501be7c8ba2c3a6f9868cb539.
//
// Solidity: event SubmitCreateBridgeToken(bytes32 indexed destinationChainID, address indexed destinationBridgeAddress, address indexed nativeContractAddress, uint256 teleporterMessageID)
func (_ERC20Bridge *ERC20BridgeFilterer) WatchSubmitCreateBridgeToken(opts *bind.WatchOpts, sink chan<- *ERC20BridgeSubmitCreateBridgeToken, destinationChainID [][32]byte, destinationBridgeAddress []common.Address, nativeContractAddress []common.Address) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var destinationBridgeAddressRule []interface{}
	for _, destinationBridgeAddressItem := range destinationBridgeAddress {
		destinationBridgeAddressRule = append(destinationBridgeAddressRule, destinationBridgeAddressItem)
	}
	var nativeContractAddressRule []interface{}
	for _, nativeContractAddressItem := range nativeContractAddress {
		nativeContractAddressRule = append(nativeContractAddressRule, nativeContractAddressItem)
	}

	logs, sub, err := _ERC20Bridge.contract.WatchLogs(opts, "SubmitCreateBridgeToken", destinationChainIDRule, destinationBridgeAddressRule, nativeContractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BridgeSubmitCreateBridgeToken)
				if err := _ERC20Bridge.contract.UnpackLog(event, "SubmitCreateBridgeToken", log); err != nil {
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

// ParseSubmitCreateBridgeToken is a log parse operation binding the contract event 0x110b902745a3d7d6b66732479f01de654a3bc6e501be7c8ba2c3a6f9868cb539.
//
// Solidity: event SubmitCreateBridgeToken(bytes32 indexed destinationChainID, address indexed destinationBridgeAddress, address indexed nativeContractAddress, uint256 teleporterMessageID)
func (_ERC20Bridge *ERC20BridgeFilterer) ParseSubmitCreateBridgeToken(log types.Log) (*ERC20BridgeSubmitCreateBridgeToken, error) {
	event := new(ERC20BridgeSubmitCreateBridgeToken)
	if err := _ERC20Bridge.contract.UnpackLog(event, "SubmitCreateBridgeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
