// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleportermessenger

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

// TeleporterMessage is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessage struct {
	MessageNonce            *big.Int
	OriginSenderAddress     common.Address
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Receipts                []TeleporterMessageReceipt
	Message                 []byte
}

// TeleporterMessageInput is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageInput struct {
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	FeeInfo                 TeleporterFeeInfo
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Message                 []byte
}

// TeleporterMessageReceipt is an auto generated low-level Go binding around an user-defined struct.
type TeleporterMessageReceipt struct {
	ReceivedMessageNonce *big.Int
	RelayerRewardAddress common.Address
}

// WarpBlockHash is an auto generated low-level Go binding around an user-defined struct.
type WarpBlockHash struct {
	SourceChainID [32]byte
	BlockHash     [32]byte
}

// WarpMessage is an auto generated low-level Go binding around an user-defined struct.
type WarpMessage struct {
	SourceChainID       [32]byte
	OriginSenderAddress common.Address
	Payload             []byte
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"}]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212209d2dabc5051a0f1ba488def6f101c5ffcdd34e71b16d3b6b944eda1ed980237464736f6c63430008190033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20PermitMetaData contains all meta data concerning the IERC20Permit contract.
var IERC20PermitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20PermitABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20PermitMetaData.ABI instead.
var IERC20PermitABI = IERC20PermitMetaData.ABI

// IERC20Permit is an auto generated Go binding around an Ethereum contract.
type IERC20Permit struct {
	IERC20PermitCaller     // Read-only binding to the contract
	IERC20PermitTransactor // Write-only binding to the contract
	IERC20PermitFilterer   // Log filterer for contract events
}

// IERC20PermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20PermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20PermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20PermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20PermitSession struct {
	Contract     *IERC20Permit     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20PermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20PermitCallerSession struct {
	Contract *IERC20PermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IERC20PermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20PermitTransactorSession struct {
	Contract     *IERC20PermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20PermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20PermitRaw struct {
	Contract *IERC20Permit // Generic contract binding to access the raw methods on
}

// IERC20PermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20PermitCallerRaw struct {
	Contract *IERC20PermitCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20PermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20PermitTransactorRaw struct {
	Contract *IERC20PermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Permit creates a new instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20Permit(address common.Address, backend bind.ContractBackend) (*IERC20Permit, error) {
	contract, err := bindIERC20Permit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Permit{IERC20PermitCaller: IERC20PermitCaller{contract: contract}, IERC20PermitTransactor: IERC20PermitTransactor{contract: contract}, IERC20PermitFilterer: IERC20PermitFilterer{contract: contract}}, nil
}

// NewIERC20PermitCaller creates a new read-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitCaller(address common.Address, caller bind.ContractCaller) (*IERC20PermitCaller, error) {
	contract, err := bindIERC20Permit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitCaller{contract: contract}, nil
}

// NewIERC20PermitTransactor creates a new write-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20PermitTransactor, error) {
	contract, err := bindIERC20Permit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitTransactor{contract: contract}, nil
}

// NewIERC20PermitFilterer creates a new log filterer instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20PermitFilterer, error) {
	contract, err := bindIERC20Permit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitFilterer{contract: contract}, nil
}

// bindIERC20Permit binds a generic wrapper to an already deployed contract.
func bindIERC20Permit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20PermitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.IERC20PermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC20Permit.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20Permit.Contract.DOMAINSEPARATOR(&_IERC20Permit.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20Permit.Contract.DOMAINSEPARATOR(&_IERC20Permit.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Permit.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20Permit.Contract.Nonces(&_IERC20Permit.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20Permit.Contract.Nonces(&_IERC20Permit.CallOpts, owner)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// ITeleporterMessengerMetaData contains all meta data concerning the ITeleporterMessenger contract.
var ITeleporterMessengerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"updatedFeeInfo\",\"type\":\"tuple\"}],\"name\":\"AddFeeAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"name\":\"BlockchainIDInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"MessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"ReceiptReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardRedeemer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ReceiveCrossChainMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RelayerRewardsRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"SendCrossChainMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalFeeAmount\",\"type\":\"uint256\"}],\"name\":\"addFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"}],\"name\":\"checkRelayerRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getNextMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getReceiptAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getReceiptQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getRelayerRewardAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"messageReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"name\":\"receiveCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"}],\"name\":\"redeemRelayerRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retryMessageExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retrySendCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageInput\",\"name\":\"messageInput\",\"type\":\"tuple\"}],\"name\":\"sendCrossChainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"messageIDs\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"sendSpecifiedReceipts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITeleporterMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use ITeleporterMessengerMetaData.ABI instead.
var ITeleporterMessengerABI = ITeleporterMessengerMetaData.ABI

// ITeleporterMessenger is an auto generated Go binding around an Ethereum contract.
type ITeleporterMessenger struct {
	ITeleporterMessengerCaller     // Read-only binding to the contract
	ITeleporterMessengerTransactor // Write-only binding to the contract
	ITeleporterMessengerFilterer   // Log filterer for contract events
}

// ITeleporterMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITeleporterMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITeleporterMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITeleporterMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITeleporterMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITeleporterMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITeleporterMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITeleporterMessengerSession struct {
	Contract     *ITeleporterMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ITeleporterMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITeleporterMessengerCallerSession struct {
	Contract *ITeleporterMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ITeleporterMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITeleporterMessengerTransactorSession struct {
	Contract     *ITeleporterMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ITeleporterMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITeleporterMessengerRaw struct {
	Contract *ITeleporterMessenger // Generic contract binding to access the raw methods on
}

// ITeleporterMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITeleporterMessengerCallerRaw struct {
	Contract *ITeleporterMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// ITeleporterMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITeleporterMessengerTransactorRaw struct {
	Contract *ITeleporterMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITeleporterMessenger creates a new instance of ITeleporterMessenger, bound to a specific deployed contract.
func NewITeleporterMessenger(address common.Address, backend bind.ContractBackend) (*ITeleporterMessenger, error) {
	contract, err := bindITeleporterMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessenger{ITeleporterMessengerCaller: ITeleporterMessengerCaller{contract: contract}, ITeleporterMessengerTransactor: ITeleporterMessengerTransactor{contract: contract}, ITeleporterMessengerFilterer: ITeleporterMessengerFilterer{contract: contract}}, nil
}

// NewITeleporterMessengerCaller creates a new read-only instance of ITeleporterMessenger, bound to a specific deployed contract.
func NewITeleporterMessengerCaller(address common.Address, caller bind.ContractCaller) (*ITeleporterMessengerCaller, error) {
	contract, err := bindITeleporterMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerCaller{contract: contract}, nil
}

// NewITeleporterMessengerTransactor creates a new write-only instance of ITeleporterMessenger, bound to a specific deployed contract.
func NewITeleporterMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*ITeleporterMessengerTransactor, error) {
	contract, err := bindITeleporterMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerTransactor{contract: contract}, nil
}

// NewITeleporterMessengerFilterer creates a new log filterer instance of ITeleporterMessenger, bound to a specific deployed contract.
func NewITeleporterMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*ITeleporterMessengerFilterer, error) {
	contract, err := bindITeleporterMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerFilterer{contract: contract}, nil
}

// bindITeleporterMessenger binds a generic wrapper to an already deployed contract.
func bindITeleporterMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITeleporterMessenger *ITeleporterMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITeleporterMessenger.Contract.ITeleporterMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITeleporterMessenger *ITeleporterMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.ITeleporterMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITeleporterMessenger *ITeleporterMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.ITeleporterMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITeleporterMessenger *ITeleporterMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITeleporterMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITeleporterMessenger *ITeleporterMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITeleporterMessenger *ITeleporterMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.contract.Transact(opts, method, params...)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeTokenAddress) view returns(uint256)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) CheckRelayerRewardAmount(opts *bind.CallOpts, relayer common.Address, feeTokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "checkRelayerRewardAmount", relayer, feeTokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeTokenAddress) view returns(uint256)
func (_ITeleporterMessenger *ITeleporterMessengerSession) CheckRelayerRewardAmount(relayer common.Address, feeTokenAddress common.Address) (*big.Int, error) {
	return _ITeleporterMessenger.Contract.CheckRelayerRewardAmount(&_ITeleporterMessenger.CallOpts, relayer, feeTokenAddress)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeTokenAddress) view returns(uint256)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) CheckRelayerRewardAmount(relayer common.Address, feeTokenAddress common.Address) (*big.Int, error) {
	return _ITeleporterMessenger.Contract.CheckRelayerRewardAmount(&_ITeleporterMessenger.CallOpts, relayer, feeTokenAddress)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) GetFeeInfo(opts *bind.CallOpts, messageID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "getFeeInfo", messageID)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_ITeleporterMessenger *ITeleporterMessengerSession) GetFeeInfo(messageID [32]byte) (common.Address, *big.Int, error) {
	return _ITeleporterMessenger.Contract.GetFeeInfo(&_ITeleporterMessenger.CallOpts, messageID)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) GetFeeInfo(messageID [32]byte) (common.Address, *big.Int, error) {
	return _ITeleporterMessenger.Contract.GetFeeInfo(&_ITeleporterMessenger.CallOpts, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) GetMessageHash(opts *bind.CallOpts, messageID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "getMessageHash", messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerSession) GetMessageHash(messageID [32]byte) ([32]byte, error) {
	return _ITeleporterMessenger.Contract.GetMessageHash(&_ITeleporterMessenger.CallOpts, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) GetMessageHash(messageID [32]byte) ([32]byte, error) {
	return _ITeleporterMessenger.Contract.GetMessageHash(&_ITeleporterMessenger.CallOpts, messageID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) GetNextMessageID(opts *bind.CallOpts, destinationBlockchainID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "getNextMessageID", destinationBlockchainID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerSession) GetNextMessageID(destinationBlockchainID [32]byte) ([32]byte, error) {
	return _ITeleporterMessenger.Contract.GetNextMessageID(&_ITeleporterMessenger.CallOpts, destinationBlockchainID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) GetNextMessageID(destinationBlockchainID [32]byte) ([32]byte, error) {
	return _ITeleporterMessenger.Contract.GetNextMessageID(&_ITeleporterMessenger.CallOpts, destinationBlockchainID)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_ITeleporterMessenger *ITeleporterMessengerCaller) GetReceiptAtIndex(opts *bind.CallOpts, sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "getReceiptAtIndex", sourceBlockchainID, index)

	if err != nil {
		return *new(TeleporterMessageReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(TeleporterMessageReceipt)).(*TeleporterMessageReceipt)

	return out0, err

}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_ITeleporterMessenger *ITeleporterMessengerSession) GetReceiptAtIndex(sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _ITeleporterMessenger.Contract.GetReceiptAtIndex(&_ITeleporterMessenger.CallOpts, sourceBlockchainID, index)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) GetReceiptAtIndex(sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _ITeleporterMessenger.Contract.GetReceiptAtIndex(&_ITeleporterMessenger.CallOpts, sourceBlockchainID, index)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) GetReceiptQueueSize(opts *bind.CallOpts, sourceBlockchainID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "getReceiptQueueSize", sourceBlockchainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_ITeleporterMessenger *ITeleporterMessengerSession) GetReceiptQueueSize(sourceBlockchainID [32]byte) (*big.Int, error) {
	return _ITeleporterMessenger.Contract.GetReceiptQueueSize(&_ITeleporterMessenger.CallOpts, sourceBlockchainID)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) GetReceiptQueueSize(sourceBlockchainID [32]byte) (*big.Int, error) {
	return _ITeleporterMessenger.Contract.GetReceiptQueueSize(&_ITeleporterMessenger.CallOpts, sourceBlockchainID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) GetRelayerRewardAddress(opts *bind.CallOpts, messageID [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "getRelayerRewardAddress", messageID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_ITeleporterMessenger *ITeleporterMessengerSession) GetRelayerRewardAddress(messageID [32]byte) (common.Address, error) {
	return _ITeleporterMessenger.Contract.GetRelayerRewardAddress(&_ITeleporterMessenger.CallOpts, messageID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) GetRelayerRewardAddress(messageID [32]byte) (common.Address, error) {
	return _ITeleporterMessenger.Contract.GetRelayerRewardAddress(&_ITeleporterMessenger.CallOpts, messageID)
}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_ITeleporterMessenger *ITeleporterMessengerCaller) MessageReceived(opts *bind.CallOpts, messageID [32]byte) (bool, error) {
	var out []interface{}
	err := _ITeleporterMessenger.contract.Call(opts, &out, "messageReceived", messageID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_ITeleporterMessenger *ITeleporterMessengerSession) MessageReceived(messageID [32]byte) (bool, error) {
	return _ITeleporterMessenger.Contract.MessageReceived(&_ITeleporterMessenger.CallOpts, messageID)
}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_ITeleporterMessenger *ITeleporterMessengerCallerSession) MessageReceived(messageID [32]byte) (bool, error) {
	return _ITeleporterMessenger.Contract.MessageReceived(&_ITeleporterMessenger.CallOpts, messageID)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) AddFeeAmount(opts *bind.TransactOpts, messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "addFeeAmount", messageID, feeTokenAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_ITeleporterMessenger *ITeleporterMessengerSession) AddFeeAmount(messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.AddFeeAmount(&_ITeleporterMessenger.TransactOpts, messageID, feeTokenAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) AddFeeAmount(messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.AddFeeAmount(&_ITeleporterMessenger.TransactOpts, messageID, feeTokenAddress, additionalFeeAmount)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) ReceiveCrossChainMessage(opts *bind.TransactOpts, messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "receiveCrossChainMessage", messageIndex, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_ITeleporterMessenger *ITeleporterMessengerSession) ReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.ReceiveCrossChainMessage(&_ITeleporterMessenger.TransactOpts, messageIndex, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) ReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.ReceiveCrossChainMessage(&_ITeleporterMessenger.TransactOpts, messageIndex, relayerRewardAddress)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeTokenAddress) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) RedeemRelayerRewards(opts *bind.TransactOpts, feeTokenAddress common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "redeemRelayerRewards", feeTokenAddress)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeTokenAddress) returns()
func (_ITeleporterMessenger *ITeleporterMessengerSession) RedeemRelayerRewards(feeTokenAddress common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.RedeemRelayerRewards(&_ITeleporterMessenger.TransactOpts, feeTokenAddress)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeTokenAddress) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) RedeemRelayerRewards(feeTokenAddress common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.RedeemRelayerRewards(&_ITeleporterMessenger.TransactOpts, feeTokenAddress)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) RetryMessageExecution(opts *bind.TransactOpts, sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "retryMessageExecution", sourceBlockchainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ITeleporterMessenger *ITeleporterMessengerSession) RetryMessageExecution(sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.RetryMessageExecution(&_ITeleporterMessenger.TransactOpts, sourceBlockchainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) RetryMessageExecution(sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.RetryMessageExecution(&_ITeleporterMessenger.TransactOpts, sourceBlockchainID, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x8245a1b0.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) RetrySendCrossChainMessage(opts *bind.TransactOpts, message TeleporterMessage) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "retrySendCrossChainMessage", message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x8245a1b0.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ITeleporterMessenger *ITeleporterMessengerSession) RetrySendCrossChainMessage(message TeleporterMessage) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.RetrySendCrossChainMessage(&_ITeleporterMessenger.TransactOpts, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x8245a1b0.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) RetrySendCrossChainMessage(message TeleporterMessage) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.RetrySendCrossChainMessage(&_ITeleporterMessenger.TransactOpts, message)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) SendCrossChainMessage(opts *bind.TransactOpts, messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "sendCrossChainMessage", messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.SendCrossChainMessage(&_ITeleporterMessenger.TransactOpts, messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.SendCrossChainMessage(&_ITeleporterMessenger.TransactOpts, messageInput)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerTransactor) SendSpecifiedReceipts(opts *bind.TransactOpts, sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.contract.Transact(opts, "sendSpecifiedReceipts", sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerSession) SendSpecifiedReceipts(sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.SendSpecifiedReceipts(&_ITeleporterMessenger.TransactOpts, sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_ITeleporterMessenger *ITeleporterMessengerTransactorSession) SendSpecifiedReceipts(sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _ITeleporterMessenger.Contract.SendSpecifiedReceipts(&_ITeleporterMessenger.TransactOpts, sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// ITeleporterMessengerAddFeeAmountIterator is returned from FilterAddFeeAmount and is used to iterate over the raw logs and unpacked data for AddFeeAmount events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerAddFeeAmountIterator struct {
	Event *ITeleporterMessengerAddFeeAmount // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerAddFeeAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerAddFeeAmount)
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
		it.Event = new(ITeleporterMessengerAddFeeAmount)
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
func (it *ITeleporterMessengerAddFeeAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerAddFeeAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerAddFeeAmount represents a AddFeeAmount event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerAddFeeAmount struct {
	MessageID      [32]byte
	UpdatedFeeInfo TeleporterFeeInfo
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddFeeAmount is a free log retrieval operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterAddFeeAmount(opts *bind.FilterOpts, messageID [][32]byte) (*ITeleporterMessengerAddFeeAmountIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "AddFeeAmount", messageIDRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerAddFeeAmountIterator{contract: _ITeleporterMessenger.contract, event: "AddFeeAmount", logs: logs, sub: sub}, nil
}

// WatchAddFeeAmount is a free log subscription operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchAddFeeAmount(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerAddFeeAmount, messageID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "AddFeeAmount", messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerAddFeeAmount)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
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

// ParseAddFeeAmount is a log parse operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseAddFeeAmount(log types.Log) (*ITeleporterMessengerAddFeeAmount, error) {
	event := new(ITeleporterMessengerAddFeeAmount)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerBlockchainIDInitializedIterator is returned from FilterBlockchainIDInitialized and is used to iterate over the raw logs and unpacked data for BlockchainIDInitialized events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerBlockchainIDInitializedIterator struct {
	Event *ITeleporterMessengerBlockchainIDInitialized // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerBlockchainIDInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerBlockchainIDInitialized)
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
		it.Event = new(ITeleporterMessengerBlockchainIDInitialized)
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
func (it *ITeleporterMessengerBlockchainIDInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerBlockchainIDInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerBlockchainIDInitialized represents a BlockchainIDInitialized event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerBlockchainIDInitialized struct {
	BlockchainID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBlockchainIDInitialized is a free log retrieval operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterBlockchainIDInitialized(opts *bind.FilterOpts, blockchainID [][32]byte) (*ITeleporterMessengerBlockchainIDInitializedIterator, error) {

	var blockchainIDRule []interface{}
	for _, blockchainIDItem := range blockchainID {
		blockchainIDRule = append(blockchainIDRule, blockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "BlockchainIDInitialized", blockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerBlockchainIDInitializedIterator{contract: _ITeleporterMessenger.contract, event: "BlockchainIDInitialized", logs: logs, sub: sub}, nil
}

// WatchBlockchainIDInitialized is a free log subscription operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchBlockchainIDInitialized(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerBlockchainIDInitialized, blockchainID [][32]byte) (event.Subscription, error) {

	var blockchainIDRule []interface{}
	for _, blockchainIDItem := range blockchainID {
		blockchainIDRule = append(blockchainIDRule, blockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "BlockchainIDInitialized", blockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerBlockchainIDInitialized)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "BlockchainIDInitialized", log); err != nil {
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

// ParseBlockchainIDInitialized is a log parse operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseBlockchainIDInitialized(log types.Log) (*ITeleporterMessengerBlockchainIDInitialized, error) {
	event := new(ITeleporterMessengerBlockchainIDInitialized)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "BlockchainIDInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerMessageExecutedIterator is returned from FilterMessageExecuted and is used to iterate over the raw logs and unpacked data for MessageExecuted events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerMessageExecutedIterator struct {
	Event *ITeleporterMessengerMessageExecuted // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerMessageExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerMessageExecuted)
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
		it.Event = new(ITeleporterMessengerMessageExecuted)
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
func (it *ITeleporterMessengerMessageExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerMessageExecuted represents a MessageExecuted event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerMessageExecuted struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageExecuted is a free log retrieval operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterMessageExecuted(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte) (*ITeleporterMessengerMessageExecutedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "MessageExecuted", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerMessageExecutedIterator{contract: _ITeleporterMessenger.contract, event: "MessageExecuted", logs: logs, sub: sub}, nil
}

// WatchMessageExecuted is a free log subscription operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerMessageExecuted, messageID [][32]byte, sourceBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "MessageExecuted", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerMessageExecuted)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
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

// ParseMessageExecuted is a log parse operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseMessageExecuted(log types.Log) (*ITeleporterMessengerMessageExecuted, error) {
	event := new(ITeleporterMessengerMessageExecuted)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerMessageExecutionFailedIterator is returned from FilterMessageExecutionFailed and is used to iterate over the raw logs and unpacked data for MessageExecutionFailed events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerMessageExecutionFailedIterator struct {
	Event *ITeleporterMessengerMessageExecutionFailed // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerMessageExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerMessageExecutionFailed)
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
		it.Event = new(ITeleporterMessengerMessageExecutionFailed)
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
func (it *ITeleporterMessengerMessageExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerMessageExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerMessageExecutionFailed represents a MessageExecutionFailed event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerMessageExecutionFailed struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageExecutionFailed is a free log retrieval operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterMessageExecutionFailed(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte) (*ITeleporterMessengerMessageExecutionFailedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "MessageExecutionFailed", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerMessageExecutionFailedIterator{contract: _ITeleporterMessenger.contract, event: "MessageExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchMessageExecutionFailed is a free log subscription operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchMessageExecutionFailed(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerMessageExecutionFailed, messageID [][32]byte, sourceBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "MessageExecutionFailed", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerMessageExecutionFailed)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
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

// ParseMessageExecutionFailed is a log parse operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseMessageExecutionFailed(log types.Log) (*ITeleporterMessengerMessageExecutionFailed, error) {
	event := new(ITeleporterMessengerMessageExecutionFailed)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerReceiptReceivedIterator is returned from FilterReceiptReceived and is used to iterate over the raw logs and unpacked data for ReceiptReceived events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerReceiptReceivedIterator struct {
	Event *ITeleporterMessengerReceiptReceived // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerReceiptReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerReceiptReceived)
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
		it.Event = new(ITeleporterMessengerReceiptReceived)
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
func (it *ITeleporterMessengerReceiptReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerReceiptReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerReceiptReceived represents a ReceiptReceived event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerReceiptReceived struct {
	MessageID               [32]byte
	DestinationBlockchainID [32]byte
	RelayerRewardAddress    common.Address
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterReceiptReceived is a free log retrieval operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterReceiptReceived(opts *bind.FilterOpts, messageID [][32]byte, destinationBlockchainID [][32]byte, relayerRewardAddress []common.Address) (*ITeleporterMessengerReceiptReceivedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var relayerRewardAddressRule []interface{}
	for _, relayerRewardAddressItem := range relayerRewardAddress {
		relayerRewardAddressRule = append(relayerRewardAddressRule, relayerRewardAddressItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "ReceiptReceived", messageIDRule, destinationBlockchainIDRule, relayerRewardAddressRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerReceiptReceivedIterator{contract: _ITeleporterMessenger.contract, event: "ReceiptReceived", logs: logs, sub: sub}, nil
}

// WatchReceiptReceived is a free log subscription operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchReceiptReceived(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerReceiptReceived, messageID [][32]byte, destinationBlockchainID [][32]byte, relayerRewardAddress []common.Address) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var relayerRewardAddressRule []interface{}
	for _, relayerRewardAddressItem := range relayerRewardAddress {
		relayerRewardAddressRule = append(relayerRewardAddressRule, relayerRewardAddressItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "ReceiptReceived", messageIDRule, destinationBlockchainIDRule, relayerRewardAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerReceiptReceived)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "ReceiptReceived", log); err != nil {
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

// ParseReceiptReceived is a log parse operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseReceiptReceived(log types.Log) (*ITeleporterMessengerReceiptReceived, error) {
	event := new(ITeleporterMessengerReceiptReceived)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "ReceiptReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerReceiveCrossChainMessageIterator is returned from FilterReceiveCrossChainMessage and is used to iterate over the raw logs and unpacked data for ReceiveCrossChainMessage events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerReceiveCrossChainMessageIterator struct {
	Event *ITeleporterMessengerReceiveCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerReceiveCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerReceiveCrossChainMessage)
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
		it.Event = new(ITeleporterMessengerReceiveCrossChainMessage)
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
func (it *ITeleporterMessengerReceiveCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerReceiveCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerReceiveCrossChainMessage represents a ReceiveCrossChainMessage event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerReceiveCrossChainMessage struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Deliverer          common.Address
	RewardRedeemer     common.Address
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReceiveCrossChainMessage is a free log retrieval operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterReceiveCrossChainMessage(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte, deliverer []common.Address) (*ITeleporterMessengerReceiveCrossChainMessageIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var delivererRule []interface{}
	for _, delivererItem := range deliverer {
		delivererRule = append(delivererRule, delivererItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "ReceiveCrossChainMessage", messageIDRule, sourceBlockchainIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerReceiveCrossChainMessageIterator{contract: _ITeleporterMessenger.contract, event: "ReceiveCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveCrossChainMessage is a free log subscription operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchReceiveCrossChainMessage(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerReceiveCrossChainMessage, messageID [][32]byte, sourceBlockchainID [][32]byte, deliverer []common.Address) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var delivererRule []interface{}
	for _, delivererItem := range deliverer {
		delivererRule = append(delivererRule, delivererItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "ReceiveCrossChainMessage", messageIDRule, sourceBlockchainIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerReceiveCrossChainMessage)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
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

// ParseReceiveCrossChainMessage is a log parse operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseReceiveCrossChainMessage(log types.Log) (*ITeleporterMessengerReceiveCrossChainMessage, error) {
	event := new(ITeleporterMessengerReceiveCrossChainMessage)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerRelayerRewardsRedeemedIterator is returned from FilterRelayerRewardsRedeemed and is used to iterate over the raw logs and unpacked data for RelayerRewardsRedeemed events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerRelayerRewardsRedeemedIterator struct {
	Event *ITeleporterMessengerRelayerRewardsRedeemed // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerRelayerRewardsRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerRelayerRewardsRedeemed)
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
		it.Event = new(ITeleporterMessengerRelayerRewardsRedeemed)
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
func (it *ITeleporterMessengerRelayerRewardsRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerRelayerRewardsRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerRelayerRewardsRedeemed represents a RelayerRewardsRedeemed event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerRelayerRewardsRedeemed struct {
	Redeemer common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRelayerRewardsRedeemed is a free log retrieval operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterRelayerRewardsRedeemed(opts *bind.FilterOpts, redeemer []common.Address, asset []common.Address) (*ITeleporterMessengerRelayerRewardsRedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerRelayerRewardsRedeemedIterator{contract: _ITeleporterMessenger.contract, event: "RelayerRewardsRedeemed", logs: logs, sub: sub}, nil
}

// WatchRelayerRewardsRedeemed is a free log subscription operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchRelayerRewardsRedeemed(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerRelayerRewardsRedeemed, redeemer []common.Address, asset []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerRelayerRewardsRedeemed)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
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

// ParseRelayerRewardsRedeemed is a log parse operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseRelayerRewardsRedeemed(log types.Log) (*ITeleporterMessengerRelayerRewardsRedeemed, error) {
	event := new(ITeleporterMessengerRelayerRewardsRedeemed)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterMessengerSendCrossChainMessageIterator is returned from FilterSendCrossChainMessage and is used to iterate over the raw logs and unpacked data for SendCrossChainMessage events raised by the ITeleporterMessenger contract.
type ITeleporterMessengerSendCrossChainMessageIterator struct {
	Event *ITeleporterMessengerSendCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *ITeleporterMessengerSendCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITeleporterMessengerSendCrossChainMessage)
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
		it.Event = new(ITeleporterMessengerSendCrossChainMessage)
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
func (it *ITeleporterMessengerSendCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITeleporterMessengerSendCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITeleporterMessengerSendCrossChainMessage represents a SendCrossChainMessage event raised by the ITeleporterMessenger contract.
type ITeleporterMessengerSendCrossChainMessage struct {
	MessageID               [32]byte
	DestinationBlockchainID [32]byte
	Message                 TeleporterMessage
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSendCrossChainMessage is a free log retrieval operation binding the contract event 0x2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) FilterSendCrossChainMessage(opts *bind.FilterOpts, messageID [][32]byte, destinationBlockchainID [][32]byte) (*ITeleporterMessengerSendCrossChainMessageIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.FilterLogs(opts, "SendCrossChainMessage", messageIDRule, destinationBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &ITeleporterMessengerSendCrossChainMessageIterator{contract: _ITeleporterMessenger.contract, event: "SendCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchSendCrossChainMessage is a free log subscription operation binding the contract event 0x2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) WatchSendCrossChainMessage(opts *bind.WatchOpts, sink chan<- *ITeleporterMessengerSendCrossChainMessage, messageID [][32]byte, destinationBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}

	logs, sub, err := _ITeleporterMessenger.contract.WatchLogs(opts, "SendCrossChainMessage", messageIDRule, destinationBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITeleporterMessengerSendCrossChainMessage)
				if err := _ITeleporterMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
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

// ParseSendCrossChainMessage is a log parse operation binding the contract event 0x2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_ITeleporterMessenger *ITeleporterMessengerFilterer) ParseSendCrossChainMessage(log types.Log) (*ITeleporterMessengerSendCrossChainMessage, error) {
	event := new(ITeleporterMessengerSendCrossChainMessage)
	if err := _ITeleporterMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITeleporterReceiverMetaData contains all meta data concerning the ITeleporterReceiver contract.
var ITeleporterReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITeleporterReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use ITeleporterReceiverMetaData.ABI instead.
var ITeleporterReceiverABI = ITeleporterReceiverMetaData.ABI

// ITeleporterReceiver is an auto generated Go binding around an Ethereum contract.
type ITeleporterReceiver struct {
	ITeleporterReceiverCaller     // Read-only binding to the contract
	ITeleporterReceiverTransactor // Write-only binding to the contract
	ITeleporterReceiverFilterer   // Log filterer for contract events
}

// ITeleporterReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITeleporterReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITeleporterReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITeleporterReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITeleporterReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITeleporterReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITeleporterReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITeleporterReceiverSession struct {
	Contract     *ITeleporterReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ITeleporterReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITeleporterReceiverCallerSession struct {
	Contract *ITeleporterReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ITeleporterReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITeleporterReceiverTransactorSession struct {
	Contract     *ITeleporterReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ITeleporterReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITeleporterReceiverRaw struct {
	Contract *ITeleporterReceiver // Generic contract binding to access the raw methods on
}

// ITeleporterReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITeleporterReceiverCallerRaw struct {
	Contract *ITeleporterReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ITeleporterReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITeleporterReceiverTransactorRaw struct {
	Contract *ITeleporterReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITeleporterReceiver creates a new instance of ITeleporterReceiver, bound to a specific deployed contract.
func NewITeleporterReceiver(address common.Address, backend bind.ContractBackend) (*ITeleporterReceiver, error) {
	contract, err := bindITeleporterReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITeleporterReceiver{ITeleporterReceiverCaller: ITeleporterReceiverCaller{contract: contract}, ITeleporterReceiverTransactor: ITeleporterReceiverTransactor{contract: contract}, ITeleporterReceiverFilterer: ITeleporterReceiverFilterer{contract: contract}}, nil
}

// NewITeleporterReceiverCaller creates a new read-only instance of ITeleporterReceiver, bound to a specific deployed contract.
func NewITeleporterReceiverCaller(address common.Address, caller bind.ContractCaller) (*ITeleporterReceiverCaller, error) {
	contract, err := bindITeleporterReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITeleporterReceiverCaller{contract: contract}, nil
}

// NewITeleporterReceiverTransactor creates a new write-only instance of ITeleporterReceiver, bound to a specific deployed contract.
func NewITeleporterReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ITeleporterReceiverTransactor, error) {
	contract, err := bindITeleporterReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITeleporterReceiverTransactor{contract: contract}, nil
}

// NewITeleporterReceiverFilterer creates a new log filterer instance of ITeleporterReceiver, bound to a specific deployed contract.
func NewITeleporterReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ITeleporterReceiverFilterer, error) {
	contract, err := bindITeleporterReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITeleporterReceiverFilterer{contract: contract}, nil
}

// bindITeleporterReceiver binds a generic wrapper to an already deployed contract.
func bindITeleporterReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITeleporterReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITeleporterReceiver *ITeleporterReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITeleporterReceiver.Contract.ITeleporterReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITeleporterReceiver *ITeleporterReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITeleporterReceiver.Contract.ITeleporterReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITeleporterReceiver *ITeleporterReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITeleporterReceiver.Contract.ITeleporterReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITeleporterReceiver *ITeleporterReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITeleporterReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITeleporterReceiver *ITeleporterReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITeleporterReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITeleporterReceiver *ITeleporterReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITeleporterReceiver.Contract.contract.Transact(opts, method, params...)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ITeleporterReceiver *ITeleporterReceiverTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ITeleporterReceiver.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ITeleporterReceiver *ITeleporterReceiverSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ITeleporterReceiver.Contract.ReceiveTeleporterMessage(&_ITeleporterReceiver.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ITeleporterReceiver *ITeleporterReceiverTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ITeleporterReceiver.Contract.ReceiveTeleporterMessage(&_ITeleporterReceiver.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// IWarpMessengerMetaData contains all meta data concerning the IWarpMessenger contract.
var IWarpMessengerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SendWarpMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"getVerifiedWarpBlockHash\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structWarpBlockHash\",\"name\":\"warpBlockHash\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"getVerifiedWarpMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structWarpMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"sendWarpMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IWarpMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use IWarpMessengerMetaData.ABI instead.
var IWarpMessengerABI = IWarpMessengerMetaData.ABI

// IWarpMessenger is an auto generated Go binding around an Ethereum contract.
type IWarpMessenger struct {
	IWarpMessengerCaller     // Read-only binding to the contract
	IWarpMessengerTransactor // Write-only binding to the contract
	IWarpMessengerFilterer   // Log filterer for contract events
}

// IWarpMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWarpMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWarpMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWarpMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWarpMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWarpMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWarpMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWarpMessengerSession struct {
	Contract     *IWarpMessenger   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWarpMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWarpMessengerCallerSession struct {
	Contract *IWarpMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IWarpMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWarpMessengerTransactorSession struct {
	Contract     *IWarpMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IWarpMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWarpMessengerRaw struct {
	Contract *IWarpMessenger // Generic contract binding to access the raw methods on
}

// IWarpMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWarpMessengerCallerRaw struct {
	Contract *IWarpMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// IWarpMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWarpMessengerTransactorRaw struct {
	Contract *IWarpMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWarpMessenger creates a new instance of IWarpMessenger, bound to a specific deployed contract.
func NewIWarpMessenger(address common.Address, backend bind.ContractBackend) (*IWarpMessenger, error) {
	contract, err := bindIWarpMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWarpMessenger{IWarpMessengerCaller: IWarpMessengerCaller{contract: contract}, IWarpMessengerTransactor: IWarpMessengerTransactor{contract: contract}, IWarpMessengerFilterer: IWarpMessengerFilterer{contract: contract}}, nil
}

// NewIWarpMessengerCaller creates a new read-only instance of IWarpMessenger, bound to a specific deployed contract.
func NewIWarpMessengerCaller(address common.Address, caller bind.ContractCaller) (*IWarpMessengerCaller, error) {
	contract, err := bindIWarpMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWarpMessengerCaller{contract: contract}, nil
}

// NewIWarpMessengerTransactor creates a new write-only instance of IWarpMessenger, bound to a specific deployed contract.
func NewIWarpMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*IWarpMessengerTransactor, error) {
	contract, err := bindIWarpMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWarpMessengerTransactor{contract: contract}, nil
}

// NewIWarpMessengerFilterer creates a new log filterer instance of IWarpMessenger, bound to a specific deployed contract.
func NewIWarpMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*IWarpMessengerFilterer, error) {
	contract, err := bindIWarpMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWarpMessengerFilterer{contract: contract}, nil
}

// bindIWarpMessenger binds a generic wrapper to an already deployed contract.
func bindIWarpMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWarpMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWarpMessenger *IWarpMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWarpMessenger.Contract.IWarpMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWarpMessenger *IWarpMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWarpMessenger.Contract.IWarpMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWarpMessenger *IWarpMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWarpMessenger.Contract.IWarpMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWarpMessenger *IWarpMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWarpMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWarpMessenger *IWarpMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWarpMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWarpMessenger *IWarpMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWarpMessenger.Contract.contract.Transact(opts, method, params...)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32 blockchainID)
func (_IWarpMessenger *IWarpMessengerCaller) GetBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IWarpMessenger.contract.Call(opts, &out, "getBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32 blockchainID)
func (_IWarpMessenger *IWarpMessengerSession) GetBlockchainID() ([32]byte, error) {
	return _IWarpMessenger.Contract.GetBlockchainID(&_IWarpMessenger.CallOpts)
}

// GetBlockchainID is a free data retrieval call binding the contract method 0x4213cf78.
//
// Solidity: function getBlockchainID() view returns(bytes32 blockchainID)
func (_IWarpMessenger *IWarpMessengerCallerSession) GetBlockchainID() ([32]byte, error) {
	return _IWarpMessenger.Contract.GetBlockchainID(&_IWarpMessenger.CallOpts)
}

// GetVerifiedWarpBlockHash is a free data retrieval call binding the contract method 0xce7f5929.
//
// Solidity: function getVerifiedWarpBlockHash(uint32 index) view returns((bytes32,bytes32) warpBlockHash, bool valid)
func (_IWarpMessenger *IWarpMessengerCaller) GetVerifiedWarpBlockHash(opts *bind.CallOpts, index uint32) (struct {
	WarpBlockHash WarpBlockHash
	Valid         bool
}, error) {
	var out []interface{}
	err := _IWarpMessenger.contract.Call(opts, &out, "getVerifiedWarpBlockHash", index)

	outstruct := new(struct {
		WarpBlockHash WarpBlockHash
		Valid         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WarpBlockHash = *abi.ConvertType(out[0], new(WarpBlockHash)).(*WarpBlockHash)
	outstruct.Valid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetVerifiedWarpBlockHash is a free data retrieval call binding the contract method 0xce7f5929.
//
// Solidity: function getVerifiedWarpBlockHash(uint32 index) view returns((bytes32,bytes32) warpBlockHash, bool valid)
func (_IWarpMessenger *IWarpMessengerSession) GetVerifiedWarpBlockHash(index uint32) (struct {
	WarpBlockHash WarpBlockHash
	Valid         bool
}, error) {
	return _IWarpMessenger.Contract.GetVerifiedWarpBlockHash(&_IWarpMessenger.CallOpts, index)
}

// GetVerifiedWarpBlockHash is a free data retrieval call binding the contract method 0xce7f5929.
//
// Solidity: function getVerifiedWarpBlockHash(uint32 index) view returns((bytes32,bytes32) warpBlockHash, bool valid)
func (_IWarpMessenger *IWarpMessengerCallerSession) GetVerifiedWarpBlockHash(index uint32) (struct {
	WarpBlockHash WarpBlockHash
	Valid         bool
}, error) {
	return _IWarpMessenger.Contract.GetVerifiedWarpBlockHash(&_IWarpMessenger.CallOpts, index)
}

// GetVerifiedWarpMessage is a free data retrieval call binding the contract method 0x6f825350.
//
// Solidity: function getVerifiedWarpMessage(uint32 index) view returns((bytes32,address,bytes) message, bool valid)
func (_IWarpMessenger *IWarpMessengerCaller) GetVerifiedWarpMessage(opts *bind.CallOpts, index uint32) (struct {
	Message WarpMessage
	Valid   bool
}, error) {
	var out []interface{}
	err := _IWarpMessenger.contract.Call(opts, &out, "getVerifiedWarpMessage", index)

	outstruct := new(struct {
		Message WarpMessage
		Valid   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Message = *abi.ConvertType(out[0], new(WarpMessage)).(*WarpMessage)
	outstruct.Valid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetVerifiedWarpMessage is a free data retrieval call binding the contract method 0x6f825350.
//
// Solidity: function getVerifiedWarpMessage(uint32 index) view returns((bytes32,address,bytes) message, bool valid)
func (_IWarpMessenger *IWarpMessengerSession) GetVerifiedWarpMessage(index uint32) (struct {
	Message WarpMessage
	Valid   bool
}, error) {
	return _IWarpMessenger.Contract.GetVerifiedWarpMessage(&_IWarpMessenger.CallOpts, index)
}

// GetVerifiedWarpMessage is a free data retrieval call binding the contract method 0x6f825350.
//
// Solidity: function getVerifiedWarpMessage(uint32 index) view returns((bytes32,address,bytes) message, bool valid)
func (_IWarpMessenger *IWarpMessengerCallerSession) GetVerifiedWarpMessage(index uint32) (struct {
	Message WarpMessage
	Valid   bool
}, error) {
	return _IWarpMessenger.Contract.GetVerifiedWarpMessage(&_IWarpMessenger.CallOpts, index)
}

// SendWarpMessage is a paid mutator transaction binding the contract method 0xee5b48eb.
//
// Solidity: function sendWarpMessage(bytes payload) returns(bytes32 messageID)
func (_IWarpMessenger *IWarpMessengerTransactor) SendWarpMessage(opts *bind.TransactOpts, payload []byte) (*types.Transaction, error) {
	return _IWarpMessenger.contract.Transact(opts, "sendWarpMessage", payload)
}

// SendWarpMessage is a paid mutator transaction binding the contract method 0xee5b48eb.
//
// Solidity: function sendWarpMessage(bytes payload) returns(bytes32 messageID)
func (_IWarpMessenger *IWarpMessengerSession) SendWarpMessage(payload []byte) (*types.Transaction, error) {
	return _IWarpMessenger.Contract.SendWarpMessage(&_IWarpMessenger.TransactOpts, payload)
}

// SendWarpMessage is a paid mutator transaction binding the contract method 0xee5b48eb.
//
// Solidity: function sendWarpMessage(bytes payload) returns(bytes32 messageID)
func (_IWarpMessenger *IWarpMessengerTransactorSession) SendWarpMessage(payload []byte) (*types.Transaction, error) {
	return _IWarpMessenger.Contract.SendWarpMessage(&_IWarpMessenger.TransactOpts, payload)
}

// IWarpMessengerSendWarpMessageIterator is returned from FilterSendWarpMessage and is used to iterate over the raw logs and unpacked data for SendWarpMessage events raised by the IWarpMessenger contract.
type IWarpMessengerSendWarpMessageIterator struct {
	Event *IWarpMessengerSendWarpMessage // Event containing the contract specifics and raw log

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
func (it *IWarpMessengerSendWarpMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWarpMessengerSendWarpMessage)
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
		it.Event = new(IWarpMessengerSendWarpMessage)
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
func (it *IWarpMessengerSendWarpMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWarpMessengerSendWarpMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWarpMessengerSendWarpMessage represents a SendWarpMessage event raised by the IWarpMessenger contract.
type IWarpMessengerSendWarpMessage struct {
	Sender    common.Address
	MessageID [32]byte
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSendWarpMessage is a free log retrieval operation binding the contract event 0x56600c567728a800c0aa927500f831cb451df66a7af570eb4df4dfbf4674887d.
//
// Solidity: event SendWarpMessage(address indexed sender, bytes32 indexed messageID, bytes message)
func (_IWarpMessenger *IWarpMessengerFilterer) FilterSendWarpMessage(opts *bind.FilterOpts, sender []common.Address, messageID [][32]byte) (*IWarpMessengerSendWarpMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _IWarpMessenger.contract.FilterLogs(opts, "SendWarpMessage", senderRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return &IWarpMessengerSendWarpMessageIterator{contract: _IWarpMessenger.contract, event: "SendWarpMessage", logs: logs, sub: sub}, nil
}

// WatchSendWarpMessage is a free log subscription operation binding the contract event 0x56600c567728a800c0aa927500f831cb451df66a7af570eb4df4dfbf4674887d.
//
// Solidity: event SendWarpMessage(address indexed sender, bytes32 indexed messageID, bytes message)
func (_IWarpMessenger *IWarpMessengerFilterer) WatchSendWarpMessage(opts *bind.WatchOpts, sink chan<- *IWarpMessengerSendWarpMessage, sender []common.Address, messageID [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _IWarpMessenger.contract.WatchLogs(opts, "SendWarpMessage", senderRule, messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWarpMessengerSendWarpMessage)
				if err := _IWarpMessenger.contract.UnpackLog(event, "SendWarpMessage", log); err != nil {
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

// ParseSendWarpMessage is a log parse operation binding the contract event 0x56600c567728a800c0aa927500f831cb451df66a7af570eb4df4dfbf4674887d.
//
// Solidity: event SendWarpMessage(address indexed sender, bytes32 indexed messageID, bytes message)
func (_IWarpMessenger *IWarpMessengerFilterer) ParseSendWarpMessage(log types.Log) (*IWarpMessengerSendWarpMessage, error) {
	event := new(IWarpMessengerSendWarpMessage)
	if err := _IWarpMessenger.contract.UnpackLog(event, "SendWarpMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"}]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212203d0969b61d0d77012e6ad59a2d8b7436adebee3c994fb3ef324be3735eee46fd64736f6c63430008190033",
}

// MathABI is the input ABI used to generate the binding from.
// Deprecated: Use MathMetaData.ABI instead.
var MathABI = MathMetaData.ABI

// MathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathMetaData.Bin instead.
var MathBin = MathMetaData.Bin

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// ReceiptQueueMetaData contains all meta data concerning the ReceiptQueue contract.
var ReceiptQueueMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea264697066735822122038a742d6e870fa1fd5ffc8e616284b2f5a03e535930dff90cb0e0ee6e67eec8e64736f6c63430008190033",
}

// ReceiptQueueABI is the input ABI used to generate the binding from.
// Deprecated: Use ReceiptQueueMetaData.ABI instead.
var ReceiptQueueABI = ReceiptQueueMetaData.ABI

// ReceiptQueueBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReceiptQueueMetaData.Bin instead.
var ReceiptQueueBin = ReceiptQueueMetaData.Bin

// DeployReceiptQueue deploys a new Ethereum contract, binding an instance of ReceiptQueue to it.
func DeployReceiptQueue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReceiptQueue, error) {
	parsed, err := ReceiptQueueMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiptQueueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReceiptQueue{ReceiptQueueCaller: ReceiptQueueCaller{contract: contract}, ReceiptQueueTransactor: ReceiptQueueTransactor{contract: contract}, ReceiptQueueFilterer: ReceiptQueueFilterer{contract: contract}}, nil
}

// ReceiptQueue is an auto generated Go binding around an Ethereum contract.
type ReceiptQueue struct {
	ReceiptQueueCaller     // Read-only binding to the contract
	ReceiptQueueTransactor // Write-only binding to the contract
	ReceiptQueueFilterer   // Log filterer for contract events
}

// ReceiptQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiptQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiptQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiptQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiptQueueSession struct {
	Contract     *ReceiptQueue     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiptQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiptQueueCallerSession struct {
	Contract *ReceiptQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ReceiptQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiptQueueTransactorSession struct {
	Contract     *ReceiptQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ReceiptQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiptQueueRaw struct {
	Contract *ReceiptQueue // Generic contract binding to access the raw methods on
}

// ReceiptQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiptQueueCallerRaw struct {
	Contract *ReceiptQueueCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiptQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiptQueueTransactorRaw struct {
	Contract *ReceiptQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiptQueue creates a new instance of ReceiptQueue, bound to a specific deployed contract.
func NewReceiptQueue(address common.Address, backend bind.ContractBackend) (*ReceiptQueue, error) {
	contract, err := bindReceiptQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiptQueue{ReceiptQueueCaller: ReceiptQueueCaller{contract: contract}, ReceiptQueueTransactor: ReceiptQueueTransactor{contract: contract}, ReceiptQueueFilterer: ReceiptQueueFilterer{contract: contract}}, nil
}

// NewReceiptQueueCaller creates a new read-only instance of ReceiptQueue, bound to a specific deployed contract.
func NewReceiptQueueCaller(address common.Address, caller bind.ContractCaller) (*ReceiptQueueCaller, error) {
	contract, err := bindReceiptQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptQueueCaller{contract: contract}, nil
}

// NewReceiptQueueTransactor creates a new write-only instance of ReceiptQueue, bound to a specific deployed contract.
func NewReceiptQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiptQueueTransactor, error) {
	contract, err := bindReceiptQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptQueueTransactor{contract: contract}, nil
}

// NewReceiptQueueFilterer creates a new log filterer instance of ReceiptQueue, bound to a specific deployed contract.
func NewReceiptQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiptQueueFilterer, error) {
	contract, err := bindReceiptQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiptQueueFilterer{contract: contract}, nil
}

// bindReceiptQueue binds a generic wrapper to an already deployed contract.
func bindReceiptQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReceiptQueueMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptQueue *ReceiptQueueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptQueue.Contract.ReceiptQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptQueue *ReceiptQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptQueue.Contract.ReceiptQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptQueue *ReceiptQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptQueue.Contract.ReceiptQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptQueue *ReceiptQueueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptQueue *ReceiptQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptQueue *ReceiptQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptQueue.Contract.contract.Transact(opts, method, params...)
}

// ReentrancyGuardsMetaData contains all meta data concerning the ReentrancyGuards contract.
var ReentrancyGuardsMetaData = &bind.MetaData{
	ABI: "[]",
}

// ReentrancyGuardsABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyGuardsMetaData.ABI instead.
var ReentrancyGuardsABI = ReentrancyGuardsMetaData.ABI

// ReentrancyGuards is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuards struct {
	ReentrancyGuardsCaller     // Read-only binding to the contract
	ReentrancyGuardsTransactor // Write-only binding to the contract
	ReentrancyGuardsFilterer   // Log filterer for contract events
}

// ReentrancyGuardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardsSession struct {
	Contract     *ReentrancyGuards // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardsCallerSession struct {
	Contract *ReentrancyGuardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ReentrancyGuardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardsTransactorSession struct {
	Contract     *ReentrancyGuardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ReentrancyGuardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardsRaw struct {
	Contract *ReentrancyGuards // Generic contract binding to access the raw methods on
}

// ReentrancyGuardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardsCallerRaw struct {
	Contract *ReentrancyGuardsCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardsTransactorRaw struct {
	Contract *ReentrancyGuardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuards creates a new instance of ReentrancyGuards, bound to a specific deployed contract.
func NewReentrancyGuards(address common.Address, backend bind.ContractBackend) (*ReentrancyGuards, error) {
	contract, err := bindReentrancyGuards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuards{ReentrancyGuardsCaller: ReentrancyGuardsCaller{contract: contract}, ReentrancyGuardsTransactor: ReentrancyGuardsTransactor{contract: contract}, ReentrancyGuardsFilterer: ReentrancyGuardsFilterer{contract: contract}}, nil
}

// NewReentrancyGuardsCaller creates a new read-only instance of ReentrancyGuards, bound to a specific deployed contract.
func NewReentrancyGuardsCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardsCaller, error) {
	contract, err := bindReentrancyGuards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardsCaller{contract: contract}, nil
}

// NewReentrancyGuardsTransactor creates a new write-only instance of ReentrancyGuards, bound to a specific deployed contract.
func NewReentrancyGuardsTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardsTransactor, error) {
	contract, err := bindReentrancyGuards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardsTransactor{contract: contract}, nil
}

// NewReentrancyGuardsFilterer creates a new log filterer instance of ReentrancyGuards, bound to a specific deployed contract.
func NewReentrancyGuardsFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardsFilterer, error) {
	contract, err := bindReentrancyGuards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardsFilterer{contract: contract}, nil
}

// bindReentrancyGuards binds a generic wrapper to an already deployed contract.
func bindReentrancyGuards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReentrancyGuardsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuards *ReentrancyGuardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuards.Contract.ReentrancyGuardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuards *ReentrancyGuardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuards.Contract.ReentrancyGuardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuards *ReentrancyGuardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuards.Contract.ReentrancyGuardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuards *ReentrancyGuardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuards *ReentrancyGuardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuards *ReentrancyGuardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuards.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedDecrease\",\"type\":\"uint256\"}],\"name\":\"SafeERC20FailedDecreaseAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea264697066735822122044774e8ae9cb3b510a51dacb03d3b3d6bee0ee00ebb240bae28efc962ab9298a64736f6c63430008190033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20TransferFromMetaData contains all meta data concerning the SafeERC20TransferFrom contract.
var SafeERC20TransferFromMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212206a67e765459c7c24f187543272046fd3214ad556da06c492ce0a4aee92d9549264736f6c63430008190033",
}

// SafeERC20TransferFromABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20TransferFromMetaData.ABI instead.
var SafeERC20TransferFromABI = SafeERC20TransferFromMetaData.ABI

// SafeERC20TransferFromBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20TransferFromMetaData.Bin instead.
var SafeERC20TransferFromBin = SafeERC20TransferFromMetaData.Bin

// DeploySafeERC20TransferFrom deploys a new Ethereum contract, binding an instance of SafeERC20TransferFrom to it.
func DeploySafeERC20TransferFrom(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20TransferFrom, error) {
	parsed, err := SafeERC20TransferFromMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20TransferFromBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20TransferFrom{SafeERC20TransferFromCaller: SafeERC20TransferFromCaller{contract: contract}, SafeERC20TransferFromTransactor: SafeERC20TransferFromTransactor{contract: contract}, SafeERC20TransferFromFilterer: SafeERC20TransferFromFilterer{contract: contract}}, nil
}

// SafeERC20TransferFrom is an auto generated Go binding around an Ethereum contract.
type SafeERC20TransferFrom struct {
	SafeERC20TransferFromCaller     // Read-only binding to the contract
	SafeERC20TransferFromTransactor // Write-only binding to the contract
	SafeERC20TransferFromFilterer   // Log filterer for contract events
}

// SafeERC20TransferFromCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20TransferFromCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20TransferFromTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20TransferFromTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20TransferFromFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20TransferFromFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20TransferFromSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20TransferFromSession struct {
	Contract     *SafeERC20TransferFrom // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SafeERC20TransferFromCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20TransferFromCallerSession struct {
	Contract *SafeERC20TransferFromCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// SafeERC20TransferFromTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransferFromTransactorSession struct {
	Contract     *SafeERC20TransferFromTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// SafeERC20TransferFromRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20TransferFromRaw struct {
	Contract *SafeERC20TransferFrom // Generic contract binding to access the raw methods on
}

// SafeERC20TransferFromCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20TransferFromCallerRaw struct {
	Contract *SafeERC20TransferFromCaller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransferFromTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransferFromTransactorRaw struct {
	Contract *SafeERC20TransferFromTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20TransferFrom creates a new instance of SafeERC20TransferFrom, bound to a specific deployed contract.
func NewSafeERC20TransferFrom(address common.Address, backend bind.ContractBackend) (*SafeERC20TransferFrom, error) {
	contract, err := bindSafeERC20TransferFrom(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20TransferFrom{SafeERC20TransferFromCaller: SafeERC20TransferFromCaller{contract: contract}, SafeERC20TransferFromTransactor: SafeERC20TransferFromTransactor{contract: contract}, SafeERC20TransferFromFilterer: SafeERC20TransferFromFilterer{contract: contract}}, nil
}

// NewSafeERC20TransferFromCaller creates a new read-only instance of SafeERC20TransferFrom, bound to a specific deployed contract.
func NewSafeERC20TransferFromCaller(address common.Address, caller bind.ContractCaller) (*SafeERC20TransferFromCaller, error) {
	contract, err := bindSafeERC20TransferFrom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20TransferFromCaller{contract: contract}, nil
}

// NewSafeERC20TransferFromTransactor creates a new write-only instance of SafeERC20TransferFrom, bound to a specific deployed contract.
func NewSafeERC20TransferFromTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20TransferFromTransactor, error) {
	contract, err := bindSafeERC20TransferFrom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20TransferFromTransactor{contract: contract}, nil
}

// NewSafeERC20TransferFromFilterer creates a new log filterer instance of SafeERC20TransferFrom, bound to a specific deployed contract.
func NewSafeERC20TransferFromFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20TransferFromFilterer, error) {
	contract, err := bindSafeERC20TransferFrom(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20TransferFromFilterer{contract: contract}, nil
}

// bindSafeERC20TransferFrom binds a generic wrapper to an already deployed contract.
func bindSafeERC20TransferFrom(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeERC20TransferFromMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20TransferFrom *SafeERC20TransferFromRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20TransferFrom.Contract.SafeERC20TransferFromCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20TransferFrom *SafeERC20TransferFromRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20TransferFrom.Contract.SafeERC20TransferFromTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20TransferFrom *SafeERC20TransferFromRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20TransferFrom.Contract.SafeERC20TransferFromTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20TransferFrom *SafeERC20TransferFromCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20TransferFrom.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20TransferFrom *SafeERC20TransferFromTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20TransferFrom.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20TransferFrom *SafeERC20TransferFromTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20TransferFrom.Contract.contract.Transact(opts, method, params...)
}

// TeleporterMessengerMetaData contains all meta data concerning the TeleporterMessenger contract.
var TeleporterMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"updatedFeeInfo\",\"type\":\"tuple\"}],\"name\":\"AddFeeAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockchainID\",\"type\":\"bytes32\"}],\"name\":\"BlockchainIDInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"MessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"MessageExecutionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"ReceiptReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deliverer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardRedeemer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"ReceiveCrossChainMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RelayerRewardsRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"name\":\"SendCrossChainMessage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"additionalFeeAmount\",\"type\":\"uint256\"}],\"name\":\"addFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"calculateMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"checkRelayerRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getNextMessageID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getReceiptAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getReceiptQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"getRelayerRewardAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"messageReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"receiptQueues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"first\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"name\":\"receiveCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"receivedFailedMessageHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"redeemRelayerRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retryMessageExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receivedMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"relayerRewardAddress\",\"type\":\"address\"}],\"internalType\":\"structTeleporterMessageReceipt[]\",\"name\":\"receipts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"retrySendCrossChainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structTeleporterMessageInput\",\"name\":\"messageInput\",\"type\":\"tuple\"}],\"name\":\"sendCrossChainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"messageIDs\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"sendSpecifiedReceipts\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageID\",\"type\":\"bytes32\"}],\"name\":\"sentMessageInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b5060015f8190558055613112806100245f395ff3fe608060405234801561000f575f80fd5b5060043610610148575f3560e01c8063a8898181116100bf578063df20e8bc11610079578063df20e8bc14610331578063e69d606a14610344578063e6e67bd5146103ab578063ebc3b1ba146103e6578063ecc7042814610409578063fc2d619714610412575f80fd5b8063a8898181146102a9578063a9a85614146102bc578063b771b3bc146102cf578063c473eef8146102dd578063ccb5f80914610315578063d127dc9b14610328575f80fd5b8063399b77da11610110578063399b77da1461021257806362448850146102315780638245a1b014610244578063860a3b0614610257578063892bf412146102765780638ac0fd0414610296575f80fd5b80630af5b4ff1461014c57806322296c3a146101675780632bc8b0bf1461017c5780632ca40f551461018f5780632e27c223146101e7575b5f80fd5b610154610425565b6040519081526020015b60405180910390f35b61017a610175366004612111565b6104f3565b005b61015461018a36600461212c565b6105e6565b6101d961019d36600461212c565b600560209081525f9182526040918290208054835180850190945260018201546001600160a01b03168452600290910154918301919091529082565b60405161015e929190612143565b6101fa6101f536600461212c565b610602565b6040516001600160a01b03909116815260200161015e565b61015461022036600461212c565b5f9081526005602052604090205490565b61015461023f36600461216a565b610689565b61017a6102523660046121b7565b6106e2565b61015461026536600461212c565b60066020525f908152604090205481565b6102896102843660046121e8565b610885565b60405161015e9190612208565b61017a6102a4366004612228565b6108b6565b6101546102b736600461225d565b610aed565b6101546102ca3660046122cd565b610b2f565b6101fa6005600160991b0181565b6101546102eb36600461235e565b6001600160a01b039182165f90815260096020908152604080832093909416825291909152205490565b61017a610323366004612395565b610dc1565b61015460025481565b61015461033f36600461212c565b6111e3565b61038c61035236600461212c565b5f90815260056020908152604091829020825180840190935260018101546001600160a01b03168084526002909101549290910182905291565b604080516001600160a01b03909316835260208301919091520161015e565b6103d16103b936600461212c565b60046020525f90815260409020805460019091015482565b6040805192835260208301919091520161015e565b6103f96103f436600461212c565b61122a565b604051901515815260200161015e565b61015460035481565b61017a6104203660046123b9565b61123f565b6002545f90806104ee576005600160991b016001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015610472573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061049691906123fc565b9050806104be5760405162461bcd60e51b81526004016104b590612413565b60405180910390fd5b600281905560405181907f1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104905f90a25b919050565b335f9081526009602090815260408083206001600160a01b0385168452909152902054806105745760405162461bcd60e51b815260206004820152602860248201527f54656c65706f727465724d657373656e6765723a206e6f2072657761726420746044820152676f2072656465656d60c01b60648201526084016104b5565b335f8181526009602090815260408083206001600160a01b03871680855290835281842093909355518481529192917f3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88910160405180910390a36105e26001600160a01b0383163383611494565b5050565b5f8181526004602052604081206105fc906114f8565b92915050565b5f8181526007602052604081205461066e5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f74604482015268081c9958d95a5d995960ba1b60648201526084016104b5565b505f908152600860205260409020546001600160a01b031690565b5f60015f54146106ab5760405162461bcd60e51b81526004016104b59061245a565b60025f556106d86106bb83612691565b83355f9081526004602052604090206106d39061150a565b611604565b60015f5592915050565b60015f54146107035760405162461bcd60e51b81526004016104b59061245a565b60025f818155905461071b9060408401358435610aed565b5f818152600560209081526040918290208251808401845281548152835180850190945260018201546001600160a01b0316845260029091015483830152908101919091528051919250906107825760405162461bcd60e51b81526004016104b590612730565b5f8360405160200161079491906129af565b60408051601f19818403018152919052825181516020830120919250146107cd5760405162461bcd60e51b81526004016104b5906129c1565b8360400135837f2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8868560200151604051610808929190612a0a565b60405180910390a360405163ee5b48eb60e01b81526005600160991b019063ee5b48eb9061083a908490600401612a8b565b6020604051808303815f875af1158015610856573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061087a91906123fc565b505060015f55505050565b604080518082019091525f80825260208201525f8381526004602052604090206108af9083611837565b9392505050565b60015f54146108d75760405162461bcd60e51b81526004016104b59061245a565b60025f5560018054146108fc5760405162461bcd60e51b81526004016104b590612a9d565b6002600155806109665760405162461bcd60e51b815260206004820152602f60248201527f54656c65706f727465724d657373656e6765723a207a65726f2061646469746960448201526e1bdb985b0819995948185b5bdd5b9d608a1b60648201526084016104b5565b6001600160a01b03821661098c5760405162461bcd60e51b81526004016104b590612ae2565b5f838152600560205260409020546109b65760405162461bcd60e51b81526004016104b590612730565b5f838152600560205260409020600101546001600160a01b03838116911614610a475760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465724d657373656e6765723a20696e76616c69642066656560448201527f20617373657420636f6e7472616374206164647265737300000000000000000060648201526084016104b5565b5f610a5283836118f8565b5f85815260056020526040812060020180549293508392909190610a77908490612b4a565b90915550505f8481526005602052604090819020905185917fc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e191610ad8916001019081546001600160a01b0316815260019190910154602082015260400190565b60405180910390a2505060018080555f555050565b6040805130602082015290810184905260608101839052608081018290525f9060a0016040516020818303038152906040528051906020012090509392505050565b5f60015f5414610b515760405162461bcd60e51b81526004016104b59061245a565b60025f818155905490866001600160401b03811115610b7257610b7261249d565b604051908082528060200260200182016040528015610bb657816020015b604080518082019091525f8082526020820152815260200190600190039081610b905790505b509050865f5b81811015610d2e575f8a8a83818110610bd757610bd7612b5d565b9050602002013590505f60075f8381526020019081526020015f20549050805f03610c535760405162461bcd60e51b815260206004820152602660248201527f54656c65706f727465724d657373656e6765723a2072656365697074206e6f7460448201526508199bdd5b9960d21b60648201526084016104b5565b610c5e8d8783610aed565b8214610cd25760405162461bcd60e51b815260206004820152603a60248201527f54656c65706f727465724d657373656e6765723a206d6573736167652049442060448201527f6e6f742066726f6d20736f7572636520626c6f636b636861696e00000000000060648201526084016104b5565b5f828152600860209081526040918290205482518084019093528383526001600160a01b03169082018190528651909190879086908110610d1557610d15612b5d565b6020026020010181905250505050806001019050610bbc565b506040805160c0810182528b81525f6020820152610daf918101610d57368b90038b018b612b71565b81526020015f81526020018888808060200260200160405190810160405280939291908181526020018383602002808284375f9201829052509385525050604080519283526020808401909152909201525083611604565b60015f559a9950505050505050505050565b6001805414610de25760405162461bcd60e51b81526004016104b590612a9d565b60026001556040516306f8253560e41b815263ffffffff831660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa158015610e30573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610e579190810190612be7565b9150915080610eba5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a20696e76616c69642077617260448201526870206d65737361676560b81b60648201526084016104b5565b60208201516001600160a01b03163014610f315760405162461bcd60e51b815260206004820152603260248201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206f726960448201527167696e2073656e646572206164647265737360701b60648201526084016104b5565b5f8260400151806020019051810190610f4a9190612d7b565b90505f610f55610425565b905080826040015114610fc45760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206465736044820152701d1a5b985d1a5bdb8818da185a5b881251607a1b60648201526084016104b5565b835182515f91610fd5918490610aed565b5f81815260076020526040902054909150156110495760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465724d657373656e6765723a206d65737361676520616c7260448201526c1958591e481c9958d95a5d9959609a1b60648201526084016104b5565b611057338460a00151611a5a565b6110b55760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465724d657373656e6765723a20756e617574686f72697a6560448201526832103932b630bcb2b960b91b60648201526084016104b5565b6110c281845f0151611ac6565b6001600160a01b038616156110f8575f81815260086020526040902080546001600160a01b0319166001600160a01b0388161790555b60c0830151515f5b8181101561113b5761113384885f01518760c00151848151811061112657611126612b5d565b6020026020010151611b36565b600101611100565b50604080518082018252855181526001600160a01b03891660208083019190915288515f90815260049091529190912061117491611c5a565b336001600160a01b0316865f0151837f292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e348a886040516111b4929190612f81565b60405180910390a460e084015151156111d5576111d582875f015186611cb4565b505060018055505050505050565b6002545f90806112055760405162461bcd60e51b81526004016104b590612413565b5f60035460016112159190612b4a565b9050611222828583610aed565b949350505050565b5f8181526007602052604081205415156105fc565b60018054146112605760405162461bcd60e51b81526004016104b590612a9d565b60026001819055545f906112779084908435610aed565b5f81815260066020526040902054909150806112a55760405162461bcd60e51b81526004016104b590612730565b80836040516020016112b791906129af565b60405160208183030381529060405280519060200120146112ea5760405162461bcd60e51b81526004016104b5906129c1565b5f6112fb6080850160608601612111565b6001600160a01b03163b1161136f5760405162461bcd60e51b815260206004820152603460248201527f54656c65706f727465724d657373656e6765723a2064657374696e6174696f6e604482015273206164647265737320686173206e6f20636f646560601b60648201526084016104b5565b604051849083907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c905f90a35f82815260066020908152604080832083905586916113be918701908701612111565b6113cb60e0870187612fa4565b6040516024016113de9493929190612fe6565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b17905290505f61142561141e6080870160608801612111565b5a84611de3565b9050806114885760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f727465724d657373656e6765723a20726574727920657865637560448201526a1d1a5bdb8819985a5b195960aa1b60648201526084016104b5565b50506001805550505050565b6040516001600160a01b038381166024830152604482018390526114f391859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050611dfa565b505050565b805460018201545f916105fc91613010565b60605f611520600561151b856114f8565b611e5b565b9050805f0361156c57604080515f8082526020820190925290611564565b604080518082019091525f808252602082015281526020019060019003908161153e5790505b509392505050565b5f816001600160401b038111156115855761158561249d565b6040519080825280602002602001820160405280156115c957816020015b604080518082019091525f80825260208201528152602001906001900390816115a35790505b5090505f5b82811015611564576115df85611e70565b8282815181106115f1576115f1612b5d565b60209081029190910101526001016115ce565b5f8061160e610425565b90505f60035f815461161f90613023565b91905081905590505f61163683875f015184610aed565b90505f604051806101000160405280848152602001336001600160a01b03168152602001885f0151815260200188602001516001600160a01b0316815260200188606001518152602001886080015181526020018781526020018860a0015181525090505f816040516020016116ac919061303b565b60405160208183030381529060405290505f808960400151602001511115611713576040890151516001600160a01b03166116f95760405162461bcd60e51b81526004016104b590612ae2565b6040890151805160209091015161171091906118f8565b90505b6040805180820182528a820151516001600160a01b03908116825260208083018590528351808501855286518783012081528082018481525f8a815260058452869020915182555180516001830180546001600160a01b03191691909516179093559101516002909101558a51915190919086907f2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8906117b6908890869061304d565b60405180910390a360405163ee5b48eb60e01b81526005600160991b019063ee5b48eb906117e8908690600401612a8b565b6020604051808303815f875af1158015611804573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061182891906123fc565b50939998505050505050505050565b604080518082019091525f8082526020820152611853836114f8565b82106118ab5760405162461bcd60e51b815260206004820152602160248201527f5265636569707451756575653a20696e646578206f7574206f6620626f756e646044820152607360f81b60648201526084016104b5565b826002015f83855f01546118bf9190612b4a565b815260208082019290925260409081015f20815180830190925280548252600101546001600160a01b0316918101919091529392505050565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038516906370a0823190602401602060405180830381865afa15801561193e573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061196291906123fc565b90506119796001600160a01b038516333086611f3a565b6040516370a0823160e01b81523060048201525f906001600160a01b038616906370a0823190602401602060405180830381865afa1580156119bd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119e191906123fc565b9050818111611a475760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016104b5565b611a518282613010565b95945050505050565b5f81515f03611a6b575060016105fc565b81515f5b81811015611abc57846001600160a01b0316848281518110611a9357611a93612b5d565b60200260200101516001600160a01b031603611ab4576001925050506105fc565b600101611a6f565b505f949350505050565b805f03611b255760405162461bcd60e51b815260206004820152602760248201527f54656c65706f727465724d657373656e6765723a207a65726f206d657373616760448201526665206e6f6e636560c81b60648201526084016104b5565b5f9182526007602052604090912055565b5f611b458484845f0151610aed565b5f818152600560209081526040918290208251808401845281548152835180850190945260018201546001600160a01b031684526002909101548383015290810191909152805191925090611b9b575050505050565b5f8281526005602090815260408083208381556001810180546001600160a01b03191690556002018390558382018051830151878401516001600160a01b0390811686526009855283862092515116855292528220805491929091611c01908490612b4a565b9250508190555082602001516001600160a01b031684837fd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf578460200151604051611c4b919061305f565b60405180910390a45050505050565b600182018054829160028501915f9182611c7383613023565b9091555081526020808201929092526040015f2082518155910151600190910180546001600160a01b0319166001600160a01b039092169190911790555050565b80608001515a1015611d165760405162461bcd60e51b815260206004820152602560248201527f54656c65706f727465724d657373656e6765723a20696e73756666696369656e604482015264742067617360d81b60648201526084016104b5565b80606001516001600160a01b03163b5f03611d36576114f3838383611f79565b602081015160e08201516040515f92611d5392869260240161307f565b60408051601f198184030181529190526020810180516001600160e01b031663643477d560e11b179052606083015160808401519192505f91611d97919084611de3565b905080611db057611da9858585611f79565b5050505050565b604051849086907f34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c905f90a35050505050565b5f805f808451602086015f8989f195945050505050565b5f611e0e6001600160a01b03841683611fed565b905080515f14158015611e32575080806020019051810190611e3091906130a8565b155b156114f357604051635274afe760e01b81526001600160a01b03841660048201526024016104b5565b5f818310611e6957816108af565b5090919050565b604080518082019091525f808252602082015281546001830154819003611ed95760405162461bcd60e51b815260206004820152601960248201527f5265636569707451756575653a20656d7074792071756575650000000000000060448201526064016104b5565b5f8181526002840160208181526040808420815180830190925280548252600180820180546001600160a01b03811685870152888852959094529490556001600160a01b0319909216905590611f30908390612b4a565b9093555090919050565b6040516001600160a01b038481166024830152838116604483015260648201839052611f739186918216906323b872dd906084016114c1565b50505050565b80604051602001611f8a919061303b565b60408051601f1981840301815282825280516020918201205f878152600690925291902055829084907f4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c98590611fe090859061303b565b60405180910390a3505050565b60606108af83835f845f80856001600160a01b0316848660405161201191906130c1565b5f6040518083038185875af1925050503d805f811461204b576040519150601f19603f3d011682016040523d82523d5f602084013e612050565b606091505b509150915061206086838361206a565b9695505050505050565b60608261207f5761207a826120c6565b6108af565b815115801561209657506001600160a01b0384163b155b156120bf57604051639996b31560e01b81526001600160a01b03851660048201526024016104b5565b50806108af565b8051156120d65780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b50565b6001600160a01b03811681146120ef575f80fd5b80356104ee816120f2565b5f60208284031215612121575f80fd5b81356108af816120f2565b5f6020828403121561213c575f80fd5b5035919050565b828152606081016108af602083018480516001600160a01b03168252602090810151910152565b5f6020828403121561217a575f80fd5b81356001600160401b0381111561218f575f80fd5b820160e081850312156108af575f80fd5b5f61010082840312156121b1575f80fd5b50919050565b5f602082840312156121c7575f80fd5b81356001600160401b038111156121dc575f80fd5b611222848285016121a0565b5f80604083850312156121f9575f80fd5b50508035926020909101359150565b815181526020808301516001600160a01b031690820152604081016105fc565b5f805f6060848603121561223a575f80fd5b83359250602084013561224c816120f2565b929592945050506040919091013590565b5f805f6060848603121561226f575f80fd5b505081359360208301359350604090920135919050565b5f8083601f840112612296575f80fd5b5081356001600160401b038111156122ac575f80fd5b6020830191508360208260051b85010111156122c6575f80fd5b9250929050565b5f805f805f8086880360a08112156122e3575f80fd5b8735965060208801356001600160401b0380821115612300575f80fd5b61230c8b838c01612286565b90985096508691506040603f1984011215612325575f80fd5b60408a01955060808a013592508083111561233e575f80fd5b505061234c89828a01612286565b979a9699509497509295939492505050565b5f806040838503121561236f575f80fd5b823561237a816120f2565b9150602083013561238a816120f2565b809150509250929050565b5f80604083850312156123a6575f80fd5b823563ffffffff8116811461237a575f80fd5b5f80604083850312156123ca575f80fd5b8235915060208301356001600160401b038111156123e6575f80fd5b6123f2858286016121a0565b9150509250929050565b5f6020828403121561240c575f80fd5b5051919050565b60208082526027908201527f54656c65706f727465724d657373656e6765723a207a65726f20626c6f636b636040820152661a185a5b88125160ca1b606082015260800190565b60208082526023908201527f5265656e7472616e63794775617264733a2073656e646572207265656e7472616040820152626e637960e81b606082015260800190565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156124d3576124d361249d565b60405290565b60405160c081016001600160401b03811182821017156124d3576124d361249d565b60405161010081016001600160401b03811182821017156124d3576124d361249d565b604051601f8201601f191681016001600160401b03811182821017156125465761254661249d565b604052919050565b5f6040828403121561255e575f80fd5b6125666124b1565b90508135612573816120f2565b808252506020820135602082015292915050565b5f6001600160401b0382111561259f5761259f61249d565b5060051b60200190565b5f82601f8301126125b8575f80fd5b813560206125cd6125c883612587565b61251e565b8083825260208201915060208460051b8701019350868411156125ee575f80fd5b602086015b84811015612613578035612606816120f2565b83529183019183016125f3565b509695505050505050565b5f6001600160401b038211156126365761263661249d565b50601f01601f191660200190565b5f82601f830112612653575f80fd5b81356126616125c88261261e565b818152846020838601011115612675575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60e082360312156126a1575f80fd5b6126a96124d9565b823581526126b960208401612106565b60208201526126cb366040850161254e565b60408201526080830135606082015260a08301356001600160401b03808211156126f3575f80fd5b6126ff368387016125a9565b608084015260c0850135915080821115612717575f80fd5b5061272436828601612644565b60a08301525092915050565b60208082526026908201527f54656c65706f727465724d657373656e6765723a206d657373616765206e6f7460408201526508199bdd5b9960d21b606082015260800190565b5f808335601e1984360301811261278b575f80fd5b83016020810192503590506001600160401b038111156127a9575f80fd5b8060051b36038213156122c6575f80fd5b8183525f60208085019450825f5b858110156127f65781356127db816120f2565b6001600160a01b0316875295820195908201906001016127c8565b509495945050505050565b5f808335601e19843603018112612816575f80fd5b83016020810192503590506001600160401b03811115612834575f80fd5b8060061b36038213156122c6575f80fd5b8183525f60208085019450825f5b858110156127f657813587528282013561286c816120f2565b6001600160a01b0316878401526040968701969190910190600101612853565b5f808335601e198436030181126128a1575f80fd5b83016020810192503590506001600160401b038111156128bf575f80fd5b8036038213156122c6575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f61010082358452602083013561290b816120f2565b6001600160a01b031660208501526040838101359085015261292f60608401612106565b6001600160a01b031660608501526080838101359085015261295460a0840184612776565b8260a087015261296783870182846127ba565b9250505061297860c0840184612801565b85830360c087015261298b838284612845565b9250505061299c60e084018461288c565b85830360e08701526120608382846128cd565b602081525f6108af60208301846128f5565b60208082526029908201527f54656c65706f727465724d657373656e6765723a20696e76616c6964206d65736040820152680e6c2ceca40d0c2e6d60bb1b606082015260800190565b606081525f612a1c60608301856128f5565b90506108af602083018480516001600160a01b03168252602090810151910152565b5f5b83811015612a58578181015183820152602001612a40565b50505f910152565b5f8151808452612a77816020860160208601612a3e565b601f01601f19169290920160200192915050565b602081525f6108af6020830184612a60565b60208082526025908201527f5265656e7472616e63794775617264733a207265636569766572207265656e7460408201526472616e637960d81b606082015260800190565b60208082526034908201527f54656c65706f727465724d657373656e6765723a207a65726f2066656520617360408201527373657420636f6e7472616374206164647265737360601b606082015260800190565b634e487b7160e01b5f52601160045260245ffd5b808201808211156105fc576105fc612b36565b634e487b7160e01b5f52603260045260245ffd5b5f60408284031215612b81575f80fd5b6108af838361254e565b80516104ee816120f2565b5f82601f830112612ba5575f80fd5b8151612bb36125c88261261e565b818152846020838601011115612bc7575f80fd5b611222826020830160208701612a3e565b805180151581146104ee575f80fd5b5f8060408385031215612bf8575f80fd5b82516001600160401b0380821115612c0e575f80fd5b9084019060608287031215612c21575f80fd5b604051606081018181108382111715612c3c57612c3c61249d565b604052825181526020830151612c51816120f2565b6020820152604083015182811115612c67575f80fd5b612c7388828601612b96565b6040830152509350612c8a91505060208401612bd8565b90509250929050565b5f82601f830112612ca2575f80fd5b81516020612cb26125c883612587565b8083825260208201915060208460051b870101935086841115612cd3575f80fd5b602086015b84811015612613578051612ceb816120f2565b8352918301918301612cd8565b5f82601f830112612d07575f80fd5b81516020612d176125c883612587565b82815260069290921b84018101918181019086841115612d35575f80fd5b8286015b848110156126135760408189031215612d50575f80fd5b612d586124b1565b8151815284820151612d69816120f2565b81860152835291830191604001612d39565b5f60208284031215612d8b575f80fd5b81516001600160401b0380821115612da1575f80fd5b908301906101008286031215612db5575f80fd5b612dbd6124fb565b82518152612dcd60208401612b8b565b602082015260408301516040820152612de860608401612b8b565b60608201526080830151608082015260a083015182811115612e08575f80fd5b612e1487828601612c93565b60a08301525060c083015182811115612e2b575f80fd5b612e3787828601612cf8565b60c08301525060e083015182811115612e4e575f80fd5b612e5a87828601612b96565b60e08301525095945050505050565b5f815180845260208085019450602084015f5b838110156127f65781516001600160a01b031687529582019590820190600101612e7c565b5f815180845260208085019450602084015f5b838110156127f657612eda878351805182526020908101516001600160a01b0316910152565b6040969096019590820190600101612eb4565b5f6101008251845260018060a01b036020840151166020850152604083015160408501526060830151612f2b60608601826001600160a01b03169052565b506080830151608085015260a08301518160a0860152612f4d82860182612e69565b91505060c083015184820360c0860152612f678282612ea1565b91505060e083015184820360e0860152611a518282612a60565b6001600160a01b03831681526040602082018190525f9061122290830184612eed565b5f808335601e19843603018112612fb9575f80fd5b8301803591506001600160401b03821115612fd2575f80fd5b6020019150368190038213156122c6575f80fd5b8481526001600160a01b03841660208201526060604082018190525f9061206090830184866128cd565b818103818111156105fc576105fc612b36565b5f6001820161303457613034612b36565b5060010190565b602081525f6108af6020830184612eed565b606081525f612a1c6060830185612eed565b81516001600160a01b0316815260208083015190820152604081016105fc565b8381526001600160a01b03831660208201526060604082018190525f90611a5190830184612a60565b5f602082840312156130b8575f80fd5b6108af82612bd8565b5f82516130d2818460208701612a3e565b919091019291505056fea2646970667358221220fd03902031337abdc8de400f99ee880fd9e0f4a12687daaae4c9c3451e31d2f164736f6c63430008190033",
}

// TeleporterMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterMessengerMetaData.ABI instead.
var TeleporterMessengerABI = TeleporterMessengerMetaData.ABI

// TeleporterMessengerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TeleporterMessengerMetaData.Bin instead.
var TeleporterMessengerBin = TeleporterMessengerMetaData.Bin

// DeployTeleporterMessenger deploys a new Ethereum contract, binding an instance of TeleporterMessenger to it.
func DeployTeleporterMessenger(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TeleporterMessenger, error) {
	parsed, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TeleporterMessengerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TeleporterMessenger{TeleporterMessengerCaller: TeleporterMessengerCaller{contract: contract}, TeleporterMessengerTransactor: TeleporterMessengerTransactor{contract: contract}, TeleporterMessengerFilterer: TeleporterMessengerFilterer{contract: contract}}, nil
}

// TeleporterMessenger is an auto generated Go binding around an Ethereum contract.
type TeleporterMessenger struct {
	TeleporterMessengerCaller     // Read-only binding to the contract
	TeleporterMessengerTransactor // Write-only binding to the contract
	TeleporterMessengerFilterer   // Log filterer for contract events
}

// TeleporterMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleporterMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleporterMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleporterMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleporterMessengerSession struct {
	Contract     *TeleporterMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TeleporterMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleporterMessengerCallerSession struct {
	Contract *TeleporterMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TeleporterMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleporterMessengerTransactorSession struct {
	Contract     *TeleporterMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TeleporterMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleporterMessengerRaw struct {
	Contract *TeleporterMessenger // Generic contract binding to access the raw methods on
}

// TeleporterMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleporterMessengerCallerRaw struct {
	Contract *TeleporterMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// TeleporterMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleporterMessengerTransactorRaw struct {
	Contract *TeleporterMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleporterMessenger creates a new instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessenger(address common.Address, backend bind.ContractBackend) (*TeleporterMessenger, error) {
	contract, err := bindTeleporterMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessenger{TeleporterMessengerCaller: TeleporterMessengerCaller{contract: contract}, TeleporterMessengerTransactor: TeleporterMessengerTransactor{contract: contract}, TeleporterMessengerFilterer: TeleporterMessengerFilterer{contract: contract}}, nil
}

// NewTeleporterMessengerCaller creates a new read-only instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessengerCaller(address common.Address, caller bind.ContractCaller) (*TeleporterMessengerCaller, error) {
	contract, err := bindTeleporterMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerCaller{contract: contract}, nil
}

// NewTeleporterMessengerTransactor creates a new write-only instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleporterMessengerTransactor, error) {
	contract, err := bindTeleporterMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerTransactor{contract: contract}, nil
}

// NewTeleporterMessengerFilterer creates a new log filterer instance of TeleporterMessenger, bound to a specific deployed contract.
func NewTeleporterMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleporterMessengerFilterer, error) {
	contract, err := bindTeleporterMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerFilterer{contract: contract}, nil
}

// bindTeleporterMessenger binds a generic wrapper to an already deployed contract.
func bindTeleporterMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterMessenger *TeleporterMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterMessenger.Contract.TeleporterMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterMessenger *TeleporterMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.TeleporterMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterMessenger *TeleporterMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.TeleporterMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterMessenger *TeleporterMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterMessenger *TeleporterMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterMessenger *TeleporterMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.contract.Transact(opts, method, params...)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterMessenger.Contract.WARPMESSENGER(&_TeleporterMessenger.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterMessenger.Contract.WARPMESSENGER(&_TeleporterMessenger.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) BlockchainID() ([32]byte, error) {
	return _TeleporterMessenger.Contract.BlockchainID(&_TeleporterMessenger.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) BlockchainID() ([32]byte, error) {
	return _TeleporterMessenger.Contract.BlockchainID(&_TeleporterMessenger.CallOpts)
}

// CalculateMessageID is a free data retrieval call binding the contract method 0xa8898181.
//
// Solidity: function calculateMessageID(bytes32 sourceBlockchainID, bytes32 destinationBlockchainID, uint256 nonce) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) CalculateMessageID(opts *bind.CallOpts, sourceBlockchainID [32]byte, destinationBlockchainID [32]byte, nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "calculateMessageID", sourceBlockchainID, destinationBlockchainID, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMessageID is a free data retrieval call binding the contract method 0xa8898181.
//
// Solidity: function calculateMessageID(bytes32 sourceBlockchainID, bytes32 destinationBlockchainID, uint256 nonce) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) CalculateMessageID(sourceBlockchainID [32]byte, destinationBlockchainID [32]byte, nonce *big.Int) ([32]byte, error) {
	return _TeleporterMessenger.Contract.CalculateMessageID(&_TeleporterMessenger.CallOpts, sourceBlockchainID, destinationBlockchainID, nonce)
}

// CalculateMessageID is a free data retrieval call binding the contract method 0xa8898181.
//
// Solidity: function calculateMessageID(bytes32 sourceBlockchainID, bytes32 destinationBlockchainID, uint256 nonce) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) CalculateMessageID(sourceBlockchainID [32]byte, destinationBlockchainID [32]byte, nonce *big.Int) ([32]byte, error) {
	return _TeleporterMessenger.Contract.CalculateMessageID(&_TeleporterMessenger.CallOpts, sourceBlockchainID, destinationBlockchainID, nonce)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) CheckRelayerRewardAmount(opts *bind.CallOpts, relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "checkRelayerRewardAmount", relayer, feeAsset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _TeleporterMessenger.Contract.CheckRelayerRewardAmount(&_TeleporterMessenger.CallOpts, relayer, feeAsset)
}

// CheckRelayerRewardAmount is a free data retrieval call binding the contract method 0xc473eef8.
//
// Solidity: function checkRelayerRewardAmount(address relayer, address feeAsset) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) CheckRelayerRewardAmount(relayer common.Address, feeAsset common.Address) (*big.Int, error) {
	return _TeleporterMessenger.Contract.CheckRelayerRewardAmount(&_TeleporterMessenger.CallOpts, relayer, feeAsset)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetFeeInfo(opts *bind.CallOpts, messageID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getFeeInfo", messageID)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) GetFeeInfo(messageID [32]byte) (common.Address, *big.Int, error) {
	return _TeleporterMessenger.Contract.GetFeeInfo(&_TeleporterMessenger.CallOpts, messageID)
}

// GetFeeInfo is a free data retrieval call binding the contract method 0xe69d606a.
//
// Solidity: function getFeeInfo(bytes32 messageID) view returns(address, uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetFeeInfo(messageID [32]byte) (common.Address, *big.Int, error) {
	return _TeleporterMessenger.Contract.GetFeeInfo(&_TeleporterMessenger.CallOpts, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetMessageHash(opts *bind.CallOpts, messageID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getMessageHash", messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) GetMessageHash(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetMessageHash(&_TeleporterMessenger.CallOpts, messageID)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x399b77da.
//
// Solidity: function getMessageHash(bytes32 messageID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetMessageHash(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetMessageHash(&_TeleporterMessenger.CallOpts, messageID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetNextMessageID(opts *bind.CallOpts, destinationBlockchainID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getNextMessageID", destinationBlockchainID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) GetNextMessageID(destinationBlockchainID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetNextMessageID(&_TeleporterMessenger.CallOpts, destinationBlockchainID)
}

// GetNextMessageID is a free data retrieval call binding the contract method 0xdf20e8bc.
//
// Solidity: function getNextMessageID(bytes32 destinationBlockchainID) view returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetNextMessageID(destinationBlockchainID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.GetNextMessageID(&_TeleporterMessenger.CallOpts, destinationBlockchainID)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessenger *TeleporterMessengerCaller) GetReceiptAtIndex(opts *bind.CallOpts, sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getReceiptAtIndex", sourceBlockchainID, index)

	if err != nil {
		return *new(TeleporterMessageReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(TeleporterMessageReceipt)).(*TeleporterMessageReceipt)

	return out0, err

}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessenger *TeleporterMessengerSession) GetReceiptAtIndex(sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _TeleporterMessenger.Contract.GetReceiptAtIndex(&_TeleporterMessenger.CallOpts, sourceBlockchainID, index)
}

// GetReceiptAtIndex is a free data retrieval call binding the contract method 0x892bf412.
//
// Solidity: function getReceiptAtIndex(bytes32 sourceBlockchainID, uint256 index) view returns((uint256,address))
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetReceiptAtIndex(sourceBlockchainID [32]byte, index *big.Int) (TeleporterMessageReceipt, error) {
	return _TeleporterMessenger.Contract.GetReceiptAtIndex(&_TeleporterMessenger.CallOpts, sourceBlockchainID, index)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetReceiptQueueSize(opts *bind.CallOpts, sourceBlockchainID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getReceiptQueueSize", sourceBlockchainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) GetReceiptQueueSize(sourceBlockchainID [32]byte) (*big.Int, error) {
	return _TeleporterMessenger.Contract.GetReceiptQueueSize(&_TeleporterMessenger.CallOpts, sourceBlockchainID)
}

// GetReceiptQueueSize is a free data retrieval call binding the contract method 0x2bc8b0bf.
//
// Solidity: function getReceiptQueueSize(bytes32 sourceBlockchainID) view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetReceiptQueueSize(sourceBlockchainID [32]byte) (*big.Int, error) {
	return _TeleporterMessenger.Contract.GetReceiptQueueSize(&_TeleporterMessenger.CallOpts, sourceBlockchainID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCaller) GetRelayerRewardAddress(opts *bind.CallOpts, messageID [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "getRelayerRewardAddress", messageID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_TeleporterMessenger *TeleporterMessengerSession) GetRelayerRewardAddress(messageID [32]byte) (common.Address, error) {
	return _TeleporterMessenger.Contract.GetRelayerRewardAddress(&_TeleporterMessenger.CallOpts, messageID)
}

// GetRelayerRewardAddress is a free data retrieval call binding the contract method 0x2e27c223.
//
// Solidity: function getRelayerRewardAddress(bytes32 messageID) view returns(address)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) GetRelayerRewardAddress(messageID [32]byte) (common.Address, error) {
	return _TeleporterMessenger.Contract.GetRelayerRewardAddress(&_TeleporterMessenger.CallOpts, messageID)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerSession) MessageNonce() (*big.Int, error) {
	return _TeleporterMessenger.Contract.MessageNonce(&_TeleporterMessenger.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) MessageNonce() (*big.Int, error) {
	return _TeleporterMessenger.Contract.MessageNonce(&_TeleporterMessenger.CallOpts)
}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_TeleporterMessenger *TeleporterMessengerCaller) MessageReceived(opts *bind.CallOpts, messageID [32]byte) (bool, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "messageReceived", messageID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_TeleporterMessenger *TeleporterMessengerSession) MessageReceived(messageID [32]byte) (bool, error) {
	return _TeleporterMessenger.Contract.MessageReceived(&_TeleporterMessenger.CallOpts, messageID)
}

// MessageReceived is a free data retrieval call binding the contract method 0xebc3b1ba.
//
// Solidity: function messageReceived(bytes32 messageID) view returns(bool)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) MessageReceived(messageID [32]byte) (bool, error) {
	return _TeleporterMessenger.Contract.MessageReceived(&_TeleporterMessenger.CallOpts, messageID)
}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceBlockchainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessenger *TeleporterMessengerCaller) ReceiptQueues(opts *bind.CallOpts, sourceBlockchainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "receiptQueues", sourceBlockchainID)

	outstruct := new(struct {
		First *big.Int
		Last  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.First = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Last = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceBlockchainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessenger *TeleporterMessengerSession) ReceiptQueues(sourceBlockchainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _TeleporterMessenger.Contract.ReceiptQueues(&_TeleporterMessenger.CallOpts, sourceBlockchainID)
}

// ReceiptQueues is a free data retrieval call binding the contract method 0xe6e67bd5.
//
// Solidity: function receiptQueues(bytes32 sourceBlockchainID) view returns(uint256 first, uint256 last)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) ReceiptQueues(sourceBlockchainID [32]byte) (struct {
	First *big.Int
	Last  *big.Int
}, error) {
	return _TeleporterMessenger.Contract.ReceiptQueues(&_TeleporterMessenger.CallOpts, sourceBlockchainID)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0x860a3b06.
//
// Solidity: function receivedFailedMessageHashes(bytes32 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessenger *TeleporterMessengerCaller) ReceivedFailedMessageHashes(opts *bind.CallOpts, messageID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "receivedFailedMessageHashes", messageID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0x860a3b06.
//
// Solidity: function receivedFailedMessageHashes(bytes32 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessenger *TeleporterMessengerSession) ReceivedFailedMessageHashes(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.ReceivedFailedMessageHashes(&_TeleporterMessenger.CallOpts, messageID)
}

// ReceivedFailedMessageHashes is a free data retrieval call binding the contract method 0x860a3b06.
//
// Solidity: function receivedFailedMessageHashes(bytes32 messageID) view returns(bytes32 messageHash)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) ReceivedFailedMessageHashes(messageID [32]byte) ([32]byte, error) {
	return _TeleporterMessenger.Contract.ReceivedFailedMessageHashes(&_TeleporterMessenger.CallOpts, messageID)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x2ca40f55.
//
// Solidity: function sentMessageInfo(bytes32 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerCaller) SentMessageInfo(opts *bind.CallOpts, messageID [32]byte) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	var out []interface{}
	err := _TeleporterMessenger.contract.Call(opts, &out, "sentMessageInfo", messageID)

	outstruct := new(struct {
		MessageHash [32]byte
		FeeInfo     TeleporterFeeInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MessageHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.FeeInfo = *abi.ConvertType(out[1], new(TeleporterFeeInfo)).(*TeleporterFeeInfo)

	return *outstruct, err

}

// SentMessageInfo is a free data retrieval call binding the contract method 0x2ca40f55.
//
// Solidity: function sentMessageInfo(bytes32 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerSession) SentMessageInfo(messageID [32]byte) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _TeleporterMessenger.Contract.SentMessageInfo(&_TeleporterMessenger.CallOpts, messageID)
}

// SentMessageInfo is a free data retrieval call binding the contract method 0x2ca40f55.
//
// Solidity: function sentMessageInfo(bytes32 messageID) view returns(bytes32 messageHash, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerCallerSession) SentMessageInfo(messageID [32]byte) (struct {
	MessageHash [32]byte
	FeeInfo     TeleporterFeeInfo
}, error) {
	return _TeleporterMessenger.Contract.SentMessageInfo(&_TeleporterMessenger.CallOpts, messageID)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) AddFeeAmount(opts *bind.TransactOpts, messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "addFeeAmount", messageID, feeTokenAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) AddFeeAmount(messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.AddFeeAmount(&_TeleporterMessenger.TransactOpts, messageID, feeTokenAddress, additionalFeeAmount)
}

// AddFeeAmount is a paid mutator transaction binding the contract method 0x8ac0fd04.
//
// Solidity: function addFeeAmount(bytes32 messageID, address feeTokenAddress, uint256 additionalFeeAmount) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) AddFeeAmount(messageID [32]byte, feeTokenAddress common.Address, additionalFeeAmount *big.Int) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.AddFeeAmount(&_TeleporterMessenger.TransactOpts, messageID, feeTokenAddress, additionalFeeAmount)
}

// InitializeBlockchainID is a paid mutator transaction binding the contract method 0x0af5b4ff.
//
// Solidity: function initializeBlockchainID() returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactor) InitializeBlockchainID(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "initializeBlockchainID")
}

// InitializeBlockchainID is a paid mutator transaction binding the contract method 0x0af5b4ff.
//
// Solidity: function initializeBlockchainID() returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) InitializeBlockchainID() (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.InitializeBlockchainID(&_TeleporterMessenger.TransactOpts)
}

// InitializeBlockchainID is a paid mutator transaction binding the contract method 0x0af5b4ff.
//
// Solidity: function initializeBlockchainID() returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) InitializeBlockchainID() (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.InitializeBlockchainID(&_TeleporterMessenger.TransactOpts)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) ReceiveCrossChainMessage(opts *bind.TransactOpts, messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "receiveCrossChainMessage", messageIndex, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) ReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.ReceiveCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageIndex, relayerRewardAddress)
}

// ReceiveCrossChainMessage is a paid mutator transaction binding the contract method 0xccb5f809.
//
// Solidity: function receiveCrossChainMessage(uint32 messageIndex, address relayerRewardAddress) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) ReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.ReceiveCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageIndex, relayerRewardAddress)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) RedeemRelayerRewards(opts *bind.TransactOpts, feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "redeemRelayerRewards", feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RedeemRelayerRewards(&_TeleporterMessenger.TransactOpts, feeAsset)
}

// RedeemRelayerRewards is a paid mutator transaction binding the contract method 0x22296c3a.
//
// Solidity: function redeemRelayerRewards(address feeAsset) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) RedeemRelayerRewards(feeAsset common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RedeemRelayerRewards(&_TeleporterMessenger.TransactOpts, feeAsset)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) RetryMessageExecution(opts *bind.TransactOpts, sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "retryMessageExecution", sourceBlockchainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) RetryMessageExecution(sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetryMessageExecution(&_TeleporterMessenger.TransactOpts, sourceBlockchainID, message)
}

// RetryMessageExecution is a paid mutator transaction binding the contract method 0xfc2d6197.
//
// Solidity: function retryMessageExecution(bytes32 sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) RetryMessageExecution(sourceBlockchainID [32]byte, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetryMessageExecution(&_TeleporterMessenger.TransactOpts, sourceBlockchainID, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x8245a1b0.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactor) RetrySendCrossChainMessage(opts *bind.TransactOpts, message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "retrySendCrossChainMessage", message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x8245a1b0.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerSession) RetrySendCrossChainMessage(message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetrySendCrossChainMessage(&_TeleporterMessenger.TransactOpts, message)
}

// RetrySendCrossChainMessage is a paid mutator transaction binding the contract method 0x8245a1b0.
//
// Solidity: function retrySendCrossChainMessage((uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message) returns()
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) RetrySendCrossChainMessage(message TeleporterMessage) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.RetrySendCrossChainMessage(&_TeleporterMessenger.TransactOpts, message)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactor) SendCrossChainMessage(opts *bind.TransactOpts, messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "sendCrossChainMessage", messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageInput)
}

// SendCrossChainMessage is a paid mutator transaction binding the contract method 0x62448850.
//
// Solidity: function sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes) messageInput) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) SendCrossChainMessage(messageInput TeleporterMessageInput) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendCrossChainMessage(&_TeleporterMessenger.TransactOpts, messageInput)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactor) SendSpecifiedReceipts(opts *bind.TransactOpts, sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.contract.Transact(opts, "sendSpecifiedReceipts", sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerSession) SendSpecifiedReceipts(sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendSpecifiedReceipts(&_TeleporterMessenger.TransactOpts, sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// SendSpecifiedReceipts is a paid mutator transaction binding the contract method 0xa9a85614.
//
// Solidity: function sendSpecifiedReceipts(bytes32 sourceBlockchainID, bytes32[] messageIDs, (address,uint256) feeInfo, address[] allowedRelayerAddresses) returns(bytes32)
func (_TeleporterMessenger *TeleporterMessengerTransactorSession) SendSpecifiedReceipts(sourceBlockchainID [32]byte, messageIDs [][32]byte, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _TeleporterMessenger.Contract.SendSpecifiedReceipts(&_TeleporterMessenger.TransactOpts, sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
}

// TeleporterMessengerAddFeeAmountIterator is returned from FilterAddFeeAmount and is used to iterate over the raw logs and unpacked data for AddFeeAmount events raised by the TeleporterMessenger contract.
type TeleporterMessengerAddFeeAmountIterator struct {
	Event *TeleporterMessengerAddFeeAmount // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerAddFeeAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerAddFeeAmount)
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
		it.Event = new(TeleporterMessengerAddFeeAmount)
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
func (it *TeleporterMessengerAddFeeAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerAddFeeAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerAddFeeAmount represents a AddFeeAmount event raised by the TeleporterMessenger contract.
type TeleporterMessengerAddFeeAmount struct {
	MessageID      [32]byte
	UpdatedFeeInfo TeleporterFeeInfo
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddFeeAmount is a free log retrieval operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterAddFeeAmount(opts *bind.FilterOpts, messageID [][32]byte) (*TeleporterMessengerAddFeeAmountIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "AddFeeAmount", messageIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerAddFeeAmountIterator{contract: _TeleporterMessenger.contract, event: "AddFeeAmount", logs: logs, sub: sub}, nil
}

// WatchAddFeeAmount is a free log subscription operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchAddFeeAmount(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerAddFeeAmount, messageID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "AddFeeAmount", messageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerAddFeeAmount)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
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

// ParseAddFeeAmount is a log parse operation binding the contract event 0xc1bfd1f1208927dfbd414041dcb5256e6c9ad90dd61aec3249facbd34ff7b3e1.
//
// Solidity: event AddFeeAmount(bytes32 indexed messageID, (address,uint256) updatedFeeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseAddFeeAmount(log types.Log) (*TeleporterMessengerAddFeeAmount, error) {
	event := new(TeleporterMessengerAddFeeAmount)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "AddFeeAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerBlockchainIDInitializedIterator is returned from FilterBlockchainIDInitialized and is used to iterate over the raw logs and unpacked data for BlockchainIDInitialized events raised by the TeleporterMessenger contract.
type TeleporterMessengerBlockchainIDInitializedIterator struct {
	Event *TeleporterMessengerBlockchainIDInitialized // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerBlockchainIDInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerBlockchainIDInitialized)
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
		it.Event = new(TeleporterMessengerBlockchainIDInitialized)
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
func (it *TeleporterMessengerBlockchainIDInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerBlockchainIDInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerBlockchainIDInitialized represents a BlockchainIDInitialized event raised by the TeleporterMessenger contract.
type TeleporterMessengerBlockchainIDInitialized struct {
	BlockchainID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBlockchainIDInitialized is a free log retrieval operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterBlockchainIDInitialized(opts *bind.FilterOpts, blockchainID [][32]byte) (*TeleporterMessengerBlockchainIDInitializedIterator, error) {

	var blockchainIDRule []interface{}
	for _, blockchainIDItem := range blockchainID {
		blockchainIDRule = append(blockchainIDRule, blockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "BlockchainIDInitialized", blockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerBlockchainIDInitializedIterator{contract: _TeleporterMessenger.contract, event: "BlockchainIDInitialized", logs: logs, sub: sub}, nil
}

// WatchBlockchainIDInitialized is a free log subscription operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchBlockchainIDInitialized(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerBlockchainIDInitialized, blockchainID [][32]byte) (event.Subscription, error) {

	var blockchainIDRule []interface{}
	for _, blockchainIDItem := range blockchainID {
		blockchainIDRule = append(blockchainIDRule, blockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "BlockchainIDInitialized", blockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerBlockchainIDInitialized)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "BlockchainIDInitialized", log); err != nil {
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

// ParseBlockchainIDInitialized is a log parse operation binding the contract event 0x1eac640109dc937d2a9f42735a05f794b39a5e3759d681951d671aabbce4b104.
//
// Solidity: event BlockchainIDInitialized(bytes32 indexed blockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseBlockchainIDInitialized(log types.Log) (*TeleporterMessengerBlockchainIDInitialized, error) {
	event := new(TeleporterMessengerBlockchainIDInitialized)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "BlockchainIDInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerMessageExecutedIterator is returned from FilterMessageExecuted and is used to iterate over the raw logs and unpacked data for MessageExecuted events raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecutedIterator struct {
	Event *TeleporterMessengerMessageExecuted // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerMessageExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerMessageExecuted)
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
		it.Event = new(TeleporterMessengerMessageExecuted)
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
func (it *TeleporterMessengerMessageExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerMessageExecuted represents a MessageExecuted event raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecuted struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageExecuted is a free log retrieval operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterMessageExecuted(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte) (*TeleporterMessengerMessageExecutedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "MessageExecuted", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerMessageExecutedIterator{contract: _TeleporterMessenger.contract, event: "MessageExecuted", logs: logs, sub: sub}, nil
}

// WatchMessageExecuted is a free log subscription operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerMessageExecuted, messageID [][32]byte, sourceBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "MessageExecuted", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerMessageExecuted)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
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

// ParseMessageExecuted is a log parse operation binding the contract event 0x34795cc6b122b9a0ae684946319f1e14a577b4e8f9b3dda9ac94c21a54d3188c.
//
// Solidity: event MessageExecuted(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseMessageExecuted(log types.Log) (*TeleporterMessengerMessageExecuted, error) {
	event := new(TeleporterMessengerMessageExecuted)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerMessageExecutionFailedIterator is returned from FilterMessageExecutionFailed and is used to iterate over the raw logs and unpacked data for MessageExecutionFailed events raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecutionFailedIterator struct {
	Event *TeleporterMessengerMessageExecutionFailed // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerMessageExecutionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerMessageExecutionFailed)
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
		it.Event = new(TeleporterMessengerMessageExecutionFailed)
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
func (it *TeleporterMessengerMessageExecutionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerMessageExecutionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerMessageExecutionFailed represents a MessageExecutionFailed event raised by the TeleporterMessenger contract.
type TeleporterMessengerMessageExecutionFailed struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMessageExecutionFailed is a free log retrieval operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterMessageExecutionFailed(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte) (*TeleporterMessengerMessageExecutionFailedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "MessageExecutionFailed", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerMessageExecutionFailedIterator{contract: _TeleporterMessenger.contract, event: "MessageExecutionFailed", logs: logs, sub: sub}, nil
}

// WatchMessageExecutionFailed is a free log subscription operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchMessageExecutionFailed(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerMessageExecutionFailed, messageID [][32]byte, sourceBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "MessageExecutionFailed", messageIDRule, sourceBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerMessageExecutionFailed)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
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

// ParseMessageExecutionFailed is a log parse operation binding the contract event 0x4619adc1017b82e02eaefac01a43d50d6d8de4460774bc370c3ff0210d40c985.
//
// Solidity: event MessageExecutionFailed(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseMessageExecutionFailed(log types.Log) (*TeleporterMessengerMessageExecutionFailed, error) {
	event := new(TeleporterMessengerMessageExecutionFailed)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "MessageExecutionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerReceiptReceivedIterator is returned from FilterReceiptReceived and is used to iterate over the raw logs and unpacked data for ReceiptReceived events raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiptReceivedIterator struct {
	Event *TeleporterMessengerReceiptReceived // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerReceiptReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerReceiptReceived)
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
		it.Event = new(TeleporterMessengerReceiptReceived)
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
func (it *TeleporterMessengerReceiptReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerReceiptReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerReceiptReceived represents a ReceiptReceived event raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiptReceived struct {
	MessageID               [32]byte
	DestinationBlockchainID [32]byte
	RelayerRewardAddress    common.Address
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterReceiptReceived is a free log retrieval operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterReceiptReceived(opts *bind.FilterOpts, messageID [][32]byte, destinationBlockchainID [][32]byte, relayerRewardAddress []common.Address) (*TeleporterMessengerReceiptReceivedIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var relayerRewardAddressRule []interface{}
	for _, relayerRewardAddressItem := range relayerRewardAddress {
		relayerRewardAddressRule = append(relayerRewardAddressRule, relayerRewardAddressItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "ReceiptReceived", messageIDRule, destinationBlockchainIDRule, relayerRewardAddressRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerReceiptReceivedIterator{contract: _TeleporterMessenger.contract, event: "ReceiptReceived", logs: logs, sub: sub}, nil
}

// WatchReceiptReceived is a free log subscription operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchReceiptReceived(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerReceiptReceived, messageID [][32]byte, destinationBlockchainID [][32]byte, relayerRewardAddress []common.Address) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var relayerRewardAddressRule []interface{}
	for _, relayerRewardAddressItem := range relayerRewardAddress {
		relayerRewardAddressRule = append(relayerRewardAddressRule, relayerRewardAddressItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "ReceiptReceived", messageIDRule, destinationBlockchainIDRule, relayerRewardAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerReceiptReceived)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiptReceived", log); err != nil {
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

// ParseReceiptReceived is a log parse operation binding the contract event 0xd13a7935f29af029349bed0a2097455b91fd06190a30478c575db3f31e00bf57.
//
// Solidity: event ReceiptReceived(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, address indexed relayerRewardAddress, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseReceiptReceived(log types.Log) (*TeleporterMessengerReceiptReceived, error) {
	event := new(TeleporterMessengerReceiptReceived)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiptReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerReceiveCrossChainMessageIterator is returned from FilterReceiveCrossChainMessage and is used to iterate over the raw logs and unpacked data for ReceiveCrossChainMessage events raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiveCrossChainMessageIterator struct {
	Event *TeleporterMessengerReceiveCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerReceiveCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerReceiveCrossChainMessage)
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
		it.Event = new(TeleporterMessengerReceiveCrossChainMessage)
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
func (it *TeleporterMessengerReceiveCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerReceiveCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerReceiveCrossChainMessage represents a ReceiveCrossChainMessage event raised by the TeleporterMessenger contract.
type TeleporterMessengerReceiveCrossChainMessage struct {
	MessageID          [32]byte
	SourceBlockchainID [32]byte
	Deliverer          common.Address
	RewardRedeemer     common.Address
	Message            TeleporterMessage
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReceiveCrossChainMessage is a free log retrieval operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterReceiveCrossChainMessage(opts *bind.FilterOpts, messageID [][32]byte, sourceBlockchainID [][32]byte, deliverer []common.Address) (*TeleporterMessengerReceiveCrossChainMessageIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var delivererRule []interface{}
	for _, delivererItem := range deliverer {
		delivererRule = append(delivererRule, delivererItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "ReceiveCrossChainMessage", messageIDRule, sourceBlockchainIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerReceiveCrossChainMessageIterator{contract: _TeleporterMessenger.contract, event: "ReceiveCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveCrossChainMessage is a free log subscription operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchReceiveCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerReceiveCrossChainMessage, messageID [][32]byte, sourceBlockchainID [][32]byte, deliverer []common.Address) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var delivererRule []interface{}
	for _, delivererItem := range deliverer {
		delivererRule = append(delivererRule, delivererItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "ReceiveCrossChainMessage", messageIDRule, sourceBlockchainIDRule, delivererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerReceiveCrossChainMessage)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
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

// ParseReceiveCrossChainMessage is a log parse operation binding the contract event 0x292ee90bbaf70b5d4936025e09d56ba08f3e421156b6a568cf3c2840d9343e34.
//
// Solidity: event ReceiveCrossChainMessage(bytes32 indexed messageID, bytes32 indexed sourceBlockchainID, address indexed deliverer, address rewardRedeemer, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseReceiveCrossChainMessage(log types.Log) (*TeleporterMessengerReceiveCrossChainMessage, error) {
	event := new(TeleporterMessengerReceiveCrossChainMessage)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "ReceiveCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerRelayerRewardsRedeemedIterator is returned from FilterRelayerRewardsRedeemed and is used to iterate over the raw logs and unpacked data for RelayerRewardsRedeemed events raised by the TeleporterMessenger contract.
type TeleporterMessengerRelayerRewardsRedeemedIterator struct {
	Event *TeleporterMessengerRelayerRewardsRedeemed // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerRelayerRewardsRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerRelayerRewardsRedeemed)
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
		it.Event = new(TeleporterMessengerRelayerRewardsRedeemed)
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
func (it *TeleporterMessengerRelayerRewardsRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerRelayerRewardsRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerRelayerRewardsRedeemed represents a RelayerRewardsRedeemed event raised by the TeleporterMessenger contract.
type TeleporterMessengerRelayerRewardsRedeemed struct {
	Redeemer common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRelayerRewardsRedeemed is a free log retrieval operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterRelayerRewardsRedeemed(opts *bind.FilterOpts, redeemer []common.Address, asset []common.Address) (*TeleporterMessengerRelayerRewardsRedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerRelayerRewardsRedeemedIterator{contract: _TeleporterMessenger.contract, event: "RelayerRewardsRedeemed", logs: logs, sub: sub}, nil
}

// WatchRelayerRewardsRedeemed is a free log subscription operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchRelayerRewardsRedeemed(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerRelayerRewardsRedeemed, redeemer []common.Address, asset []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "RelayerRewardsRedeemed", redeemerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerRelayerRewardsRedeemed)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
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

// ParseRelayerRewardsRedeemed is a log parse operation binding the contract event 0x3294c84e5b0f29d9803655319087207bc94f4db29f7927846944822773780b88.
//
// Solidity: event RelayerRewardsRedeemed(address indexed redeemer, address indexed asset, uint256 amount)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseRelayerRewardsRedeemed(log types.Log) (*TeleporterMessengerRelayerRewardsRedeemed, error) {
	event := new(TeleporterMessengerRelayerRewardsRedeemed)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "RelayerRewardsRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterMessengerSendCrossChainMessageIterator is returned from FilterSendCrossChainMessage and is used to iterate over the raw logs and unpacked data for SendCrossChainMessage events raised by the TeleporterMessenger contract.
type TeleporterMessengerSendCrossChainMessageIterator struct {
	Event *TeleporterMessengerSendCrossChainMessage // Event containing the contract specifics and raw log

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
func (it *TeleporterMessengerSendCrossChainMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterMessengerSendCrossChainMessage)
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
		it.Event = new(TeleporterMessengerSendCrossChainMessage)
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
func (it *TeleporterMessengerSendCrossChainMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterMessengerSendCrossChainMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterMessengerSendCrossChainMessage represents a SendCrossChainMessage event raised by the TeleporterMessenger contract.
type TeleporterMessengerSendCrossChainMessage struct {
	MessageID               [32]byte
	DestinationBlockchainID [32]byte
	Message                 TeleporterMessage
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSendCrossChainMessage is a free log retrieval operation binding the contract event 0x2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) FilterSendCrossChainMessage(opts *bind.FilterOpts, messageID [][32]byte, destinationBlockchainID [][32]byte) (*TeleporterMessengerSendCrossChainMessageIterator, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.FilterLogs(opts, "SendCrossChainMessage", messageIDRule, destinationBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterMessengerSendCrossChainMessageIterator{contract: _TeleporterMessenger.contract, event: "SendCrossChainMessage", logs: logs, sub: sub}, nil
}

// WatchSendCrossChainMessage is a free log subscription operation binding the contract event 0x2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) WatchSendCrossChainMessage(opts *bind.WatchOpts, sink chan<- *TeleporterMessengerSendCrossChainMessage, messageID [][32]byte, destinationBlockchainID [][32]byte) (event.Subscription, error) {

	var messageIDRule []interface{}
	for _, messageIDItem := range messageID {
		messageIDRule = append(messageIDRule, messageIDItem)
	}
	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}

	logs, sub, err := _TeleporterMessenger.contract.WatchLogs(opts, "SendCrossChainMessage", messageIDRule, destinationBlockchainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterMessengerSendCrossChainMessage)
				if err := _TeleporterMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
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

// ParseSendCrossChainMessage is a log parse operation binding the contract event 0x2a211ad4a59ab9d003852404f9c57c690704ee755f3c79d2c2812ad32da99df8.
//
// Solidity: event SendCrossChainMessage(bytes32 indexed messageID, bytes32 indexed destinationBlockchainID, (uint256,address,bytes32,address,uint256,address[],(uint256,address)[],bytes) message, (address,uint256) feeInfo)
func (_TeleporterMessenger *TeleporterMessengerFilterer) ParseSendCrossChainMessage(log types.Log) (*TeleporterMessengerSendCrossChainMessage, error) {
	event := new(TeleporterMessengerSendCrossChainMessage)
	if err := _TeleporterMessenger.contract.UnpackLog(event, "SendCrossChainMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
