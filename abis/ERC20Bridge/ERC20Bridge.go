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

// Erc20bridgeMetaData contains all meta data concerning the Erc20bridge contract.
var Erc20bridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceNotIncreased\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeTokenAddress\",\"type\":\"address\"}],\"name\":\"BridgeTokenAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotBridgeTokenWithinSameChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nativeTokenAddress\",\"type\":\"address\"}],\"name\":\"CannotBridgeWrappedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"adjustedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientAdjustedAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientTotalAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientWrappedTokenBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBridgeTokenAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationBridgeAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeleporterMessengerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nativeChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nativeBridgeAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nativeContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeTokenAddress\",\"type\":\"address\"}],\"name\":\"CreateBridgeToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MintBridgeTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nativeContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"}],\"name\":\"SubmitCreateBridgeToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CREATE_BRIDGE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_BRIDGE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_PRECOMPILE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"primaryFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFeeAmount\",\"type\":\"uint256\"}],\"name\":\"bridgeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bridgedBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nativeContractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nativeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nativeSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"nativeDecimals\",\"type\":\"uint8\"}],\"name\":\"encodeCreateBridgeTokenData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nativeContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bridgeAmount\",\"type\":\"uint256\"}],\"name\":\"encodeMintBridgeTokensData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nativeContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"encodeTransferBridgeTokensData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nativeToWrappedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nativeChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"nativeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageFeeAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"messageFeeAmount\",\"type\":\"uint256\"}],\"name\":\"submitCreateBridgeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"submittedBridgeTokenCreations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenContracts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Erc20bridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20bridgeMetaData.ABI instead.
var Erc20bridgeABI = Erc20bridgeMetaData.ABI

// Erc20bridge is an auto generated Go binding around an Ethereum contract.
type Erc20bridge struct {
	Erc20bridgeCaller     // Read-only binding to the contract
	Erc20bridgeTransactor // Write-only binding to the contract
	Erc20bridgeFilterer   // Log filterer for contract events
}

// Erc20bridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20bridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20bridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20bridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20bridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20bridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20bridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20bridgeSession struct {
	Contract     *Erc20bridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20bridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20bridgeCallerSession struct {
	Contract *Erc20bridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Erc20bridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20bridgeTransactorSession struct {
	Contract     *Erc20bridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Erc20bridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20bridgeRaw struct {
	Contract *Erc20bridge // Generic contract binding to access the raw methods on
}

// Erc20bridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20bridgeCallerRaw struct {
	Contract *Erc20bridgeCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20bridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20bridgeTransactorRaw struct {
	Contract *Erc20bridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20bridge creates a new instance of Erc20bridge, bound to a specific deployed contract.
func NewErc20bridge(address common.Address, backend bind.ContractBackend) (*Erc20bridge, error) {
	contract, err := bindErc20bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20bridge{Erc20bridgeCaller: Erc20bridgeCaller{contract: contract}, Erc20bridgeTransactor: Erc20bridgeTransactor{contract: contract}, Erc20bridgeFilterer: Erc20bridgeFilterer{contract: contract}}, nil
}

// NewErc20bridgeCaller creates a new read-only instance of Erc20bridge, bound to a specific deployed contract.
func NewErc20bridgeCaller(address common.Address, caller bind.ContractCaller) (*Erc20bridgeCaller, error) {
	contract, err := bindErc20bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20bridgeCaller{contract: contract}, nil
}

// NewErc20bridgeTransactor creates a new write-only instance of Erc20bridge, bound to a specific deployed contract.
func NewErc20bridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20bridgeTransactor, error) {
	contract, err := bindErc20bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20bridgeTransactor{contract: contract}, nil
}

// NewErc20bridgeFilterer creates a new log filterer instance of Erc20bridge, bound to a specific deployed contract.
func NewErc20bridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20bridgeFilterer, error) {
	contract, err := bindErc20bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20bridgeFilterer{contract: contract}, nil
}

// bindErc20bridge binds a generic wrapper to an already deployed contract.
func bindErc20bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20bridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20bridge *Erc20bridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20bridge.Contract.Erc20bridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20bridge *Erc20bridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20bridge.Contract.Erc20bridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20bridge *Erc20bridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20bridge.Contract.Erc20bridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20bridge *Erc20bridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20bridge *Erc20bridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20bridge *Erc20bridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20bridge.Contract.contract.Transact(opts, method, params...)
}

// CREATEBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x5f217bcc.
//
// Solidity: function CREATE_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Erc20bridge *Erc20bridgeCaller) CREATEBRIDGETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20bridge.contract.Call(opts, &out, "CREATE_BRIDGE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CREATEBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x5f217bcc.
//
// Solidity: function CREATE_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Erc20bridge *Erc20bridgeSession) CREATEBRIDGETOKENSREQUIREDGAS() (*big.Int, error) {
	return _Erc20bridge.Contract.CREATEBRIDGETOKENSREQUIREDGAS(&_Erc20bridge.CallOpts)
}

// CREATEBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x5f217bcc.
//
// Solidity: function CREATE_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Erc20bridge *Erc20bridgeCallerSession) CREATEBRIDGETOKENSREQUIREDGAS() (*big.Int, error) {
	return _Erc20bridge.Contract.CREATEBRIDGETOKENSREQUIREDGAS(&_Erc20bridge.CallOpts)
}

// MINTBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x7a465fd9.
//
// Solidity: function MINT_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Erc20bridge *Erc20bridgeCaller) MINTBRIDGETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20bridge.contract.Call(opts, &out, "MINT_BRIDGE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x7a465fd9.
//
// Solidity: function MINT_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Erc20bridge *Erc20bridgeSession) MINTBRIDGETOKENSREQUIREDGAS() (*big.Int, error) {
	return _Erc20bridge.Contract.MINTBRIDGETOKENSREQUIREDGAS(&_Erc20bridge.CallOpts)
}

// MINTBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x7a465fd9.
//
// Solidity: function MINT_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Erc20bridge *Erc20bridgeCallerSession) MINTBRIDGETOKENSREQUIREDGAS() (*big.Int, error) {
	return _Erc20bridge.Contract.MINTBRIDGETOKENSREQUIREDGAS(&_Erc20bridge.CallOpts)
}

// TRANSFERBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x6b47cd9a.
//
// Solidity: function TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Erc20bridge *Erc20bridgeCaller) TRANSFERBRIDGETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20bridge.contract.Call(opts, &out, "TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TRANSFERBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x6b47cd9a.
//
// Solidity: function TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Erc20bridge *Erc20bridgeSession) TRANSFERBRIDGETOKENSREQUIREDGAS() (*big.Int, error) {
	return _Erc20bridge.Contract.TRANSFERBRIDGETOKENSREQUIREDGAS(&_Erc20bridge.CallOpts)
}

// TRANSFERBRIDGETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x6b47cd9a.
//
// Solidity: function TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_Erc20bridge *Erc20bridgeCallerSession) TRANSFERBRIDGETOKENSREQUIREDGAS() (*big.Int, error) {
	return _Erc20bridge.Contract.TRANSFERBRIDGETOKENSREQUIREDGAS(&_Erc20bridge.CallOpts)
}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_Erc20bridge *Erc20bridgeCaller) WARPPRECOMPILEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20bridge.contract.Call(opts, &out, "WARP_PRECOMPILE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_Erc20bridge *Erc20bridgeSession) WARPPRECOMPILEADDRESS() (common.Address, error) {
	return _Erc20bridge.Contract.WARPPRECOMPILEADDRESS(&_Erc20bridge.CallOpts)
}

// WARPPRECOMPILEADDRESS is a free data retrieval call binding the contract method 0x74971856.
//
// Solidity: function WARP_PRECOMPILE_ADDRESS() view returns(address)
func (_Erc20bridge *Erc20bridgeCallerSession) WARPPRECOMPILEADDRESS() (common.Address, error) {
	return _Erc20bridge.Contract.WARPPRECOMPILEADDRESS(&_Erc20bridge.CallOpts)
}

// BridgedBalances is a free data retrieval call binding the contract method 0xb9e55da1.
//
// Solidity: function bridgedBalances(bytes32 , address , address ) view returns(uint256)
func (_Erc20bridge *Erc20bridgeCaller) BridgedBalances(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20bridge.contract.Call(opts, &out, "bridgedBalances", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgedBalances is a free data retrieval call binding the contract method 0xb9e55da1.
//
// Solidity: function bridgedBalances(bytes32 , address , address ) view returns(uint256)
func (_Erc20bridge *Erc20bridgeSession) BridgedBalances(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _Erc20bridge.Contract.BridgedBalances(&_Erc20bridge.CallOpts, arg0, arg1, arg2)
}

// BridgedBalances is a free data retrieval call binding the contract method 0xb9e55da1.
//
// Solidity: function bridgedBalances(bytes32 , address , address ) view returns(uint256)
func (_Erc20bridge *Erc20bridgeCallerSession) BridgedBalances(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _Erc20bridge.Contract.BridgedBalances(&_Erc20bridge.CallOpts, arg0, arg1, arg2)
}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_Erc20bridge *Erc20bridgeCaller) CurrentChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20bridge.contract.Call(opts, &out, "currentChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_Erc20bridge *Erc20bridgeSession) CurrentChainID() ([32]byte, error) {
	return _Erc20bridge.Contract.CurrentChainID(&_Erc20bridge.CallOpts)
}

// CurrentChainID is a free data retrieval call binding the contract method 0xb179e1e7.
//
// Solidity: function currentChainID() view returns(bytes32)
func (_Erc20bridge *Erc20bridgeCallerSession) CurrentChainID() ([32]byte, error) {
	return _Erc20bridge.Contract.CurrentChainID(&_Erc20bridge.CallOpts)
}

// EncodeCreateBridgeTokenData is a free data retrieval call binding the contract method 0x367e9584.
//
// Solidity: function encodeCreateBridgeTokenData(address nativeContractAddress, string nativeName, string nativeSymbol, uint8 nativeDecimals) pure returns(bytes)
func (_Erc20bridge *Erc20bridgeCaller) EncodeCreateBridgeTokenData(opts *bind.CallOpts, nativeContractAddress common.Address, nativeName string, nativeSymbol string, nativeDecimals uint8) ([]byte, error) {
	var out []interface{}
	err := _Erc20bridge.contract.Call(opts, &out, "encodeCreateBridgeTokenData", nativeContractAddress, nativeName, nativeSymbol, nativeDecimals)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeCreateBridgeTokenData is a free data retrieval call binding the contract method 0x367e9584.
//
// Solidity: function encodeCreateBridgeTokenData(address nativeContractAddress, string nativeName, string nativeSymbol, uint8 nativeDecimals) pure returns(bytes)
func (_Erc20bridge *Erc20bridgeSession) EncodeCreateBridgeTokenData(nativeContractAddress common.Address, nativeName string, nativeSymbol string, nativeDecimals uint8) ([]byte, error) {
	return _Erc20bridge.Contract.EncodeCreateBridgeTokenData(&_Erc20bridge.CallOpts, nativeContractAddress, nativeName, nativeSymbol, nativeDecimals)
}

// EncodeCreateBridgeTokenData is a free data retrieval call binding the contract method 0x367e9584.
//
// Solidity: function encodeCreateBridgeTokenData(address nativeContractAddress, string nativeName, string nativeSymbol, uint8 nativeDecimals) pure returns(bytes)
func (_Erc20bridge *Erc20bridgeCallerSession) EncodeCreateBridgeTokenData(nativeContractAddress common.Address, nativeName string, nativeSymbol string, nativeDecimals uint8) ([]byte, error) {
	return _Erc20bridge.Contract.EncodeCreateBridgeTokenData(&_Erc20bridge.CallOpts, nativeContractAddress, nativeName, nativeSymbol, nativeDecimals)
}

// EncodeMintBridgeTokensData is a free data retrieval call binding the contract method 0x8c56fcf0.
//
// Solidity: function encodeMintBridgeTokensData(address nativeContractAddress, address recipient, uint256 bridgeAmount) pure returns(bytes)
func (_Erc20bridge *Erc20bridgeCaller) EncodeMintBridgeTokensData(opts *bind.CallOpts, nativeContractAddress common.Address, recipient common.Address, bridgeAmount *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Erc20bridge.contract.Call(opts, &out, "encodeMintBridgeTokensData", nativeContractAddress, recipient, bridgeAmount)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeMintBridgeTokensData is a free data retrieval call binding the contract method 0x8c56fcf0.
//
// Solidity: function encodeMintBridgeTokensData(address nativeContractAddress, address recipient, uint256 bridgeAmount) pure returns(bytes)
func (_Erc20bridge *Erc20bridgeSession) EncodeMintBridgeTokensData(nativeContractAddress common.Address, recipient common.Address, bridgeAmount *big.Int) ([]byte, error) {
	return _Erc20bridge.Contract.EncodeMintBridgeTokensData(&_Erc20bridge.CallOpts, nativeContractAddress, recipient, bridgeAmount)
}

// EncodeMintBridgeTokensData is a free data retrieval call binding the contract method 0x8c56fcf0.
//
// Solidity: function encodeMintBridgeTokensData(address nativeContractAddress, address recipient, uint256 bridgeAmount) pure returns(bytes)
func (_Erc20bridge *Erc20bridgeCallerSession) EncodeMintBridgeTokensData(nativeContractAddress common.Address, recipient common.Address, bridgeAmount *big.Int) ([]byte, error) {
	return _Erc20bridge.Contract.EncodeMintBridgeTokensData(&_Erc20bridge.CallOpts, nativeContractAddress, recipient, bridgeAmount)
}

// EncodeTransferBridgeTokensData is a free data retrieval call binding the contract method 0xc60da612.
//
// Solidity: function encodeTransferBridgeTokensData(bytes32 destinationChainID, address destinationBridgeAddress, address nativeContractAddress, address recipient, uint256 amount, uint256 feeAmount) pure returns(bytes)
func (_Erc20bridge *Erc20bridgeCaller) EncodeTransferBridgeTokensData(opts *bind.CallOpts, destinationChainID [32]byte, destinationBridgeAddress common.Address, nativeContractAddress common.Address, recipient common.Address, amount *big.Int, feeAmount *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Erc20bridge.contract.Call(opts, &out, "encodeTransferBridgeTokensData", destinationChainID, destinationBridgeAddress, nativeContractAddress, recipient, amount, feeAmount)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTransferBridgeTokensData is a free data retrieval call binding the contract method 0xc60da612.
//
// Solidity: function encodeTransferBridgeTokensData(bytes32 destinationChainID, address destinationBridgeAddress, address nativeContractAddress, address recipient, uint256 amount, uint256 feeAmount) pure returns(bytes)
func (_Erc20bridge *Erc20bridgeSession) EncodeTransferBridgeTokensData(destinationChainID [32]byte, destinationBridgeAddress common.Address, nativeContractAddress common.Address, recipient common.Address, amount *big.Int, feeAmount *big.Int) ([]byte, error) {
	return _Erc20bridge.Contract.EncodeTransferBridgeTokensData(&_Erc20bridge.CallOpts, destinationChainID, destinationBridgeAddress, nativeContractAddress, recipient, amount, feeAmount)
}

// EncodeTransferBridgeTokensData is a free data retrieval call binding the contract method 0xc60da612.
//
// Solidity: function encodeTransferBridgeTokensData(bytes32 destinationChainID, address destinationBridgeAddress, address nativeContractAddress, address recipient, uint256 amount, uint256 feeAmount) pure returns(bytes)
func (_Erc20bridge *Erc20bridgeCallerSession) EncodeTransferBridgeTokensData(destinationChainID [32]byte, destinationBridgeAddress common.Address, nativeContractAddress common.Address, recipient common.Address, amount *big.Int, feeAmount *big.Int) ([]byte, error) {
	return _Erc20bridge.Contract.EncodeTransferBridgeTokensData(&_Erc20bridge.CallOpts, destinationChainID, destinationBridgeAddress, nativeContractAddress, recipient, amount, feeAmount)
}

// NativeToWrappedTokens is a free data retrieval call binding the contract method 0x65435568.
//
// Solidity: function nativeToWrappedTokens(bytes32 , address , address ) view returns(address)
func (_Erc20bridge *Erc20bridgeCaller) NativeToWrappedTokens(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address, arg2 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Erc20bridge.contract.Call(opts, &out, "nativeToWrappedTokens", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToWrappedTokens is a free data retrieval call binding the contract method 0x65435568.
//
// Solidity: function nativeToWrappedTokens(bytes32 , address , address ) view returns(address)
func (_Erc20bridge *Erc20bridgeSession) NativeToWrappedTokens(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (common.Address, error) {
	return _Erc20bridge.Contract.NativeToWrappedTokens(&_Erc20bridge.CallOpts, arg0, arg1, arg2)
}

// NativeToWrappedTokens is a free data retrieval call binding the contract method 0x65435568.
//
// Solidity: function nativeToWrappedTokens(bytes32 , address , address ) view returns(address)
func (_Erc20bridge *Erc20bridgeCallerSession) NativeToWrappedTokens(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (common.Address, error) {
	return _Erc20bridge.Contract.NativeToWrappedTokens(&_Erc20bridge.CallOpts, arg0, arg1, arg2)
}

// SubmittedBridgeTokenCreations is a free data retrieval call binding the contract method 0x8343f661.
//
// Solidity: function submittedBridgeTokenCreations(bytes32 , address , address ) view returns(bool)
func (_Erc20bridge *Erc20bridgeCaller) SubmittedBridgeTokenCreations(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _Erc20bridge.contract.Call(opts, &out, "submittedBridgeTokenCreations", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SubmittedBridgeTokenCreations is a free data retrieval call binding the contract method 0x8343f661.
//
// Solidity: function submittedBridgeTokenCreations(bytes32 , address , address ) view returns(bool)
func (_Erc20bridge *Erc20bridgeSession) SubmittedBridgeTokenCreations(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (bool, error) {
	return _Erc20bridge.Contract.SubmittedBridgeTokenCreations(&_Erc20bridge.CallOpts, arg0, arg1, arg2)
}

// SubmittedBridgeTokenCreations is a free data retrieval call binding the contract method 0x8343f661.
//
// Solidity: function submittedBridgeTokenCreations(bytes32 , address , address ) view returns(bool)
func (_Erc20bridge *Erc20bridgeCallerSession) SubmittedBridgeTokenCreations(arg0 [32]byte, arg1 common.Address, arg2 common.Address) (bool, error) {
	return _Erc20bridge.Contract.SubmittedBridgeTokenCreations(&_Erc20bridge.CallOpts, arg0, arg1, arg2)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Erc20bridge *Erc20bridgeCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20bridge.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Erc20bridge *Erc20bridgeSession) TeleporterMessenger() (common.Address, error) {
	return _Erc20bridge.Contract.TeleporterMessenger(&_Erc20bridge.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_Erc20bridge *Erc20bridgeCallerSession) TeleporterMessenger() (common.Address, error) {
	return _Erc20bridge.Contract.TeleporterMessenger(&_Erc20bridge.CallOpts)
}

// WrappedTokenContracts is a free data retrieval call binding the contract method 0x9bd9abc0.
//
// Solidity: function wrappedTokenContracts(address ) view returns(bool)
func (_Erc20bridge *Erc20bridgeCaller) WrappedTokenContracts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Erc20bridge.contract.Call(opts, &out, "wrappedTokenContracts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WrappedTokenContracts is a free data retrieval call binding the contract method 0x9bd9abc0.
//
// Solidity: function wrappedTokenContracts(address ) view returns(bool)
func (_Erc20bridge *Erc20bridgeSession) WrappedTokenContracts(arg0 common.Address) (bool, error) {
	return _Erc20bridge.Contract.WrappedTokenContracts(&_Erc20bridge.CallOpts, arg0)
}

// WrappedTokenContracts is a free data retrieval call binding the contract method 0x9bd9abc0.
//
// Solidity: function wrappedTokenContracts(address ) view returns(bool)
func (_Erc20bridge *Erc20bridgeCallerSession) WrappedTokenContracts(arg0 common.Address) (bool, error) {
	return _Erc20bridge.Contract.WrappedTokenContracts(&_Erc20bridge.CallOpts, arg0)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0xc63d2207.
//
// Solidity: function bridgeTokens(bytes32 destinationChainID, address destinationBridgeAddress, address tokenContractAddress, address recipient, uint256 totalAmount, uint256 primaryFeeAmount, uint256 secondaryFeeAmount) returns()
func (_Erc20bridge *Erc20bridgeTransactor) BridgeTokens(opts *bind.TransactOpts, destinationChainID [32]byte, destinationBridgeAddress common.Address, tokenContractAddress common.Address, recipient common.Address, totalAmount *big.Int, primaryFeeAmount *big.Int, secondaryFeeAmount *big.Int) (*types.Transaction, error) {
	return _Erc20bridge.contract.Transact(opts, "bridgeTokens", destinationChainID, destinationBridgeAddress, tokenContractAddress, recipient, totalAmount, primaryFeeAmount, secondaryFeeAmount)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0xc63d2207.
//
// Solidity: function bridgeTokens(bytes32 destinationChainID, address destinationBridgeAddress, address tokenContractAddress, address recipient, uint256 totalAmount, uint256 primaryFeeAmount, uint256 secondaryFeeAmount) returns()
func (_Erc20bridge *Erc20bridgeSession) BridgeTokens(destinationChainID [32]byte, destinationBridgeAddress common.Address, tokenContractAddress common.Address, recipient common.Address, totalAmount *big.Int, primaryFeeAmount *big.Int, secondaryFeeAmount *big.Int) (*types.Transaction, error) {
	return _Erc20bridge.Contract.BridgeTokens(&_Erc20bridge.TransactOpts, destinationChainID, destinationBridgeAddress, tokenContractAddress, recipient, totalAmount, primaryFeeAmount, secondaryFeeAmount)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0xc63d2207.
//
// Solidity: function bridgeTokens(bytes32 destinationChainID, address destinationBridgeAddress, address tokenContractAddress, address recipient, uint256 totalAmount, uint256 primaryFeeAmount, uint256 secondaryFeeAmount) returns()
func (_Erc20bridge *Erc20bridgeTransactorSession) BridgeTokens(destinationChainID [32]byte, destinationBridgeAddress common.Address, tokenContractAddress common.Address, recipient common.Address, totalAmount *big.Int, primaryFeeAmount *big.Int, secondaryFeeAmount *big.Int) (*types.Transaction, error) {
	return _Erc20bridge.Contract.BridgeTokens(&_Erc20bridge.TransactOpts, destinationChainID, destinationBridgeAddress, tokenContractAddress, recipient, totalAmount, primaryFeeAmount, secondaryFeeAmount)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_Erc20bridge *Erc20bridgeTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Erc20bridge.contract.Transact(opts, "receiveTeleporterMessage", nativeChainID, nativeBridgeAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_Erc20bridge *Erc20bridgeSession) ReceiveTeleporterMessage(nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Erc20bridge.Contract.ReceiveTeleporterMessage(&_Erc20bridge.TransactOpts, nativeChainID, nativeBridgeAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes message) returns()
func (_Erc20bridge *Erc20bridgeTransactorSession) ReceiveTeleporterMessage(nativeChainID [32]byte, nativeBridgeAddress common.Address, message []byte) (*types.Transaction, error) {
	return _Erc20bridge.Contract.ReceiveTeleporterMessage(&_Erc20bridge.TransactOpts, nativeChainID, nativeBridgeAddress, message)
}

// SubmitCreateBridgeToken is a paid mutator transaction binding the contract method 0x6c7e40d1.
//
// Solidity: function submitCreateBridgeToken(bytes32 destinationChainID, address destinationBridgeAddress, address nativeToken, address messageFeeAsset, uint256 messageFeeAmount) returns()
func (_Erc20bridge *Erc20bridgeTransactor) SubmitCreateBridgeToken(opts *bind.TransactOpts, destinationChainID [32]byte, destinationBridgeAddress common.Address, nativeToken common.Address, messageFeeAsset common.Address, messageFeeAmount *big.Int) (*types.Transaction, error) {
	return _Erc20bridge.contract.Transact(opts, "submitCreateBridgeToken", destinationChainID, destinationBridgeAddress, nativeToken, messageFeeAsset, messageFeeAmount)
}

// SubmitCreateBridgeToken is a paid mutator transaction binding the contract method 0x6c7e40d1.
//
// Solidity: function submitCreateBridgeToken(bytes32 destinationChainID, address destinationBridgeAddress, address nativeToken, address messageFeeAsset, uint256 messageFeeAmount) returns()
func (_Erc20bridge *Erc20bridgeSession) SubmitCreateBridgeToken(destinationChainID [32]byte, destinationBridgeAddress common.Address, nativeToken common.Address, messageFeeAsset common.Address, messageFeeAmount *big.Int) (*types.Transaction, error) {
	return _Erc20bridge.Contract.SubmitCreateBridgeToken(&_Erc20bridge.TransactOpts, destinationChainID, destinationBridgeAddress, nativeToken, messageFeeAsset, messageFeeAmount)
}

// SubmitCreateBridgeToken is a paid mutator transaction binding the contract method 0x6c7e40d1.
//
// Solidity: function submitCreateBridgeToken(bytes32 destinationChainID, address destinationBridgeAddress, address nativeToken, address messageFeeAsset, uint256 messageFeeAmount) returns()
func (_Erc20bridge *Erc20bridgeTransactorSession) SubmitCreateBridgeToken(destinationChainID [32]byte, destinationBridgeAddress common.Address, nativeToken common.Address, messageFeeAsset common.Address, messageFeeAmount *big.Int) (*types.Transaction, error) {
	return _Erc20bridge.Contract.SubmitCreateBridgeToken(&_Erc20bridge.TransactOpts, destinationChainID, destinationBridgeAddress, nativeToken, messageFeeAsset, messageFeeAmount)
}

// Erc20bridgeBridgeTokensIterator is returned from FilterBridgeTokens and is used to iterate over the raw logs and unpacked data for BridgeTokens events raised by the Erc20bridge contract.
type Erc20bridgeBridgeTokensIterator struct {
	Event *Erc20bridgeBridgeTokens // Event containing the contract specifics and raw log

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
func (it *Erc20bridgeBridgeTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20bridgeBridgeTokens)
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
		it.Event = new(Erc20bridgeBridgeTokens)
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
func (it *Erc20bridgeBridgeTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20bridgeBridgeTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20bridgeBridgeTokens represents a BridgeTokens event raised by the Erc20bridge contract.
type Erc20bridgeBridgeTokens struct {
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
func (_Erc20bridge *Erc20bridgeFilterer) FilterBridgeTokens(opts *bind.FilterOpts, tokenContractAddress []common.Address, destinationChainID [][32]byte, teleporterMessageID []*big.Int) (*Erc20bridgeBridgeTokensIterator, error) {

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

	logs, sub, err := _Erc20bridge.contract.FilterLogs(opts, "BridgeTokens", tokenContractAddressRule, destinationChainIDRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &Erc20bridgeBridgeTokensIterator{contract: _Erc20bridge.contract, event: "BridgeTokens", logs: logs, sub: sub}, nil
}

// WatchBridgeTokens is a free log subscription operation binding the contract event 0x97935c4470efae40c8440c3abfe968a5512232dd375cc974e712f487c2b99c31.
//
// Solidity: event BridgeTokens(address indexed tokenContractAddress, bytes32 indexed destinationChainID, uint256 indexed teleporterMessageID, address destinationBridgeAddress, address recipient, uint256 amount)
func (_Erc20bridge *Erc20bridgeFilterer) WatchBridgeTokens(opts *bind.WatchOpts, sink chan<- *Erc20bridgeBridgeTokens, tokenContractAddress []common.Address, destinationChainID [][32]byte, teleporterMessageID []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Erc20bridge.contract.WatchLogs(opts, "BridgeTokens", tokenContractAddressRule, destinationChainIDRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20bridgeBridgeTokens)
				if err := _Erc20bridge.contract.UnpackLog(event, "BridgeTokens", log); err != nil {
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
func (_Erc20bridge *Erc20bridgeFilterer) ParseBridgeTokens(log types.Log) (*Erc20bridgeBridgeTokens, error) {
	event := new(Erc20bridgeBridgeTokens)
	if err := _Erc20bridge.contract.UnpackLog(event, "BridgeTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20bridgeCreateBridgeTokenIterator is returned from FilterCreateBridgeToken and is used to iterate over the raw logs and unpacked data for CreateBridgeToken events raised by the Erc20bridge contract.
type Erc20bridgeCreateBridgeTokenIterator struct {
	Event *Erc20bridgeCreateBridgeToken // Event containing the contract specifics and raw log

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
func (it *Erc20bridgeCreateBridgeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20bridgeCreateBridgeToken)
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
		it.Event = new(Erc20bridgeCreateBridgeToken)
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
func (it *Erc20bridgeCreateBridgeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20bridgeCreateBridgeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20bridgeCreateBridgeToken represents a CreateBridgeToken event raised by the Erc20bridge contract.
type Erc20bridgeCreateBridgeToken struct {
	NativeChainID         [32]byte
	NativeBridgeAddress   common.Address
	NativeContractAddress common.Address
	BridgeTokenAddress    common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCreateBridgeToken is a free log retrieval operation binding the contract event 0xe1c61a845f79534e11924517ddbedc668d0c20e467eafb4d3bd2858e2815f3b5.
//
// Solidity: event CreateBridgeToken(bytes32 indexed nativeChainID, address indexed nativeBridgeAddress, address indexed nativeContractAddress, address bridgeTokenAddress)
func (_Erc20bridge *Erc20bridgeFilterer) FilterCreateBridgeToken(opts *bind.FilterOpts, nativeChainID [][32]byte, nativeBridgeAddress []common.Address, nativeContractAddress []common.Address) (*Erc20bridgeCreateBridgeTokenIterator, error) {

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

	logs, sub, err := _Erc20bridge.contract.FilterLogs(opts, "CreateBridgeToken", nativeChainIDRule, nativeBridgeAddressRule, nativeContractAddressRule)
	if err != nil {
		return nil, err
	}
	return &Erc20bridgeCreateBridgeTokenIterator{contract: _Erc20bridge.contract, event: "CreateBridgeToken", logs: logs, sub: sub}, nil
}

// WatchCreateBridgeToken is a free log subscription operation binding the contract event 0xe1c61a845f79534e11924517ddbedc668d0c20e467eafb4d3bd2858e2815f3b5.
//
// Solidity: event CreateBridgeToken(bytes32 indexed nativeChainID, address indexed nativeBridgeAddress, address indexed nativeContractAddress, address bridgeTokenAddress)
func (_Erc20bridge *Erc20bridgeFilterer) WatchCreateBridgeToken(opts *bind.WatchOpts, sink chan<- *Erc20bridgeCreateBridgeToken, nativeChainID [][32]byte, nativeBridgeAddress []common.Address, nativeContractAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Erc20bridge.contract.WatchLogs(opts, "CreateBridgeToken", nativeChainIDRule, nativeBridgeAddressRule, nativeContractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20bridgeCreateBridgeToken)
				if err := _Erc20bridge.contract.UnpackLog(event, "CreateBridgeToken", log); err != nil {
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
func (_Erc20bridge *Erc20bridgeFilterer) ParseCreateBridgeToken(log types.Log) (*Erc20bridgeCreateBridgeToken, error) {
	event := new(Erc20bridgeCreateBridgeToken)
	if err := _Erc20bridge.contract.UnpackLog(event, "CreateBridgeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20bridgeMintBridgeTokensIterator is returned from FilterMintBridgeTokens and is used to iterate over the raw logs and unpacked data for MintBridgeTokens events raised by the Erc20bridge contract.
type Erc20bridgeMintBridgeTokensIterator struct {
	Event *Erc20bridgeMintBridgeTokens // Event containing the contract specifics and raw log

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
func (it *Erc20bridgeMintBridgeTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20bridgeMintBridgeTokens)
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
		it.Event = new(Erc20bridgeMintBridgeTokens)
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
func (it *Erc20bridgeMintBridgeTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20bridgeMintBridgeTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20bridgeMintBridgeTokens represents a MintBridgeTokens event raised by the Erc20bridge contract.
type Erc20bridgeMintBridgeTokens struct {
	ContractAddress common.Address
	Recipient       common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMintBridgeTokens is a free log retrieval operation binding the contract event 0xc0767f158f0d5394b598489a51ed607cd55a8be2dcef113ba1626efcf4c63954.
//
// Solidity: event MintBridgeTokens(address indexed contractAddress, address recipient, uint256 amount)
func (_Erc20bridge *Erc20bridgeFilterer) FilterMintBridgeTokens(opts *bind.FilterOpts, contractAddress []common.Address) (*Erc20bridgeMintBridgeTokensIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Erc20bridge.contract.FilterLogs(opts, "MintBridgeTokens", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &Erc20bridgeMintBridgeTokensIterator{contract: _Erc20bridge.contract, event: "MintBridgeTokens", logs: logs, sub: sub}, nil
}

// WatchMintBridgeTokens is a free log subscription operation binding the contract event 0xc0767f158f0d5394b598489a51ed607cd55a8be2dcef113ba1626efcf4c63954.
//
// Solidity: event MintBridgeTokens(address indexed contractAddress, address recipient, uint256 amount)
func (_Erc20bridge *Erc20bridgeFilterer) WatchMintBridgeTokens(opts *bind.WatchOpts, sink chan<- *Erc20bridgeMintBridgeTokens, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _Erc20bridge.contract.WatchLogs(opts, "MintBridgeTokens", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20bridgeMintBridgeTokens)
				if err := _Erc20bridge.contract.UnpackLog(event, "MintBridgeTokens", log); err != nil {
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
func (_Erc20bridge *Erc20bridgeFilterer) ParseMintBridgeTokens(log types.Log) (*Erc20bridgeMintBridgeTokens, error) {
	event := new(Erc20bridgeMintBridgeTokens)
	if err := _Erc20bridge.contract.UnpackLog(event, "MintBridgeTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20bridgeSubmitCreateBridgeTokenIterator is returned from FilterSubmitCreateBridgeToken and is used to iterate over the raw logs and unpacked data for SubmitCreateBridgeToken events raised by the Erc20bridge contract.
type Erc20bridgeSubmitCreateBridgeTokenIterator struct {
	Event *Erc20bridgeSubmitCreateBridgeToken // Event containing the contract specifics and raw log

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
func (it *Erc20bridgeSubmitCreateBridgeTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20bridgeSubmitCreateBridgeToken)
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
		it.Event = new(Erc20bridgeSubmitCreateBridgeToken)
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
func (it *Erc20bridgeSubmitCreateBridgeTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20bridgeSubmitCreateBridgeTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20bridgeSubmitCreateBridgeToken represents a SubmitCreateBridgeToken event raised by the Erc20bridge contract.
type Erc20bridgeSubmitCreateBridgeToken struct {
	DestinationChainID       [32]byte
	DestinationBridgeAddress common.Address
	NativeContractAddress    common.Address
	TeleporterMessageID      *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSubmitCreateBridgeToken is a free log retrieval operation binding the contract event 0x110b902745a3d7d6b66732479f01de654a3bc6e501be7c8ba2c3a6f9868cb539.
//
// Solidity: event SubmitCreateBridgeToken(bytes32 indexed destinationChainID, address indexed destinationBridgeAddress, address indexed nativeContractAddress, uint256 teleporterMessageID)
func (_Erc20bridge *Erc20bridgeFilterer) FilterSubmitCreateBridgeToken(opts *bind.FilterOpts, destinationChainID [][32]byte, destinationBridgeAddress []common.Address, nativeContractAddress []common.Address) (*Erc20bridgeSubmitCreateBridgeTokenIterator, error) {

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

	logs, sub, err := _Erc20bridge.contract.FilterLogs(opts, "SubmitCreateBridgeToken", destinationChainIDRule, destinationBridgeAddressRule, nativeContractAddressRule)
	if err != nil {
		return nil, err
	}
	return &Erc20bridgeSubmitCreateBridgeTokenIterator{contract: _Erc20bridge.contract, event: "SubmitCreateBridgeToken", logs: logs, sub: sub}, nil
}

// WatchSubmitCreateBridgeToken is a free log subscription operation binding the contract event 0x110b902745a3d7d6b66732479f01de654a3bc6e501be7c8ba2c3a6f9868cb539.
//
// Solidity: event SubmitCreateBridgeToken(bytes32 indexed destinationChainID, address indexed destinationBridgeAddress, address indexed nativeContractAddress, uint256 teleporterMessageID)
func (_Erc20bridge *Erc20bridgeFilterer) WatchSubmitCreateBridgeToken(opts *bind.WatchOpts, sink chan<- *Erc20bridgeSubmitCreateBridgeToken, destinationChainID [][32]byte, destinationBridgeAddress []common.Address, nativeContractAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Erc20bridge.contract.WatchLogs(opts, "SubmitCreateBridgeToken", destinationChainIDRule, destinationBridgeAddressRule, nativeContractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20bridgeSubmitCreateBridgeToken)
				if err := _Erc20bridge.contract.UnpackLog(event, "SubmitCreateBridgeToken", log); err != nil {
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
func (_Erc20bridge *Erc20bridgeFilterer) ParseSubmitCreateBridgeToken(log types.Log) (*Erc20bridgeSubmitCreateBridgeToken, error) {
	event := new(Erc20bridgeSubmitCreateBridgeToken)
	if err := _Erc20bridge.contract.UnpackLog(event, "SubmitCreateBridgeToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
