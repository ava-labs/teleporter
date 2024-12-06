// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package foov2

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

// BarLibraryMetaData contains all meta data concerning the BarLibrary contract.
var BarLibraryMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60e8610033600b8282823980515f1a607314602757634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe7300000000000000000000000000000000000000003014608060405260043610603c575f3560e01c80635a55a96b146040578063a050083114605c575b5f80fd5b818015604a575f80fd5b50605a6056366004607d565b9055565b005b606b6067366004609c565b5490565b60405190815260200160405180910390f35b5f8060408385031215608d575f80fd5b50508035926020909101359150565b5f6020828403121560ab575f80fd5b503591905056fea2646970667358221220513d35ce61983ed22bc200ff7d449f3c6101eac58091f4875ad462305dec605e64736f6c63430008190033",
}

// BarLibraryABI is the input ABI used to generate the binding from.
// Deprecated: Use BarLibraryMetaData.ABI instead.
var BarLibraryABI = BarLibraryMetaData.ABI

// BarLibraryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BarLibraryMetaData.Bin instead.
var BarLibraryBin = BarLibraryMetaData.Bin

// DeployBarLibrary deploys a new Ethereum contract, binding an instance of BarLibrary to it.
func DeployBarLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BarLibrary, error) {
	parsed, err := BarLibraryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BarLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BarLibrary{BarLibraryCaller: BarLibraryCaller{contract: contract}, BarLibraryTransactor: BarLibraryTransactor{contract: contract}, BarLibraryFilterer: BarLibraryFilterer{contract: contract}}, nil
}

// BarLibrary is an auto generated Go binding around an Ethereum contract.
type BarLibrary struct {
	BarLibraryCaller     // Read-only binding to the contract
	BarLibraryTransactor // Write-only binding to the contract
	BarLibraryFilterer   // Log filterer for contract events
}

// BarLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BarLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BarLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BarLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BarLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BarLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BarLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BarLibrarySession struct {
	Contract     *BarLibrary       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BarLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BarLibraryCallerSession struct {
	Contract *BarLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BarLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BarLibraryTransactorSession struct {
	Contract     *BarLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BarLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BarLibraryRaw struct {
	Contract *BarLibrary // Generic contract binding to access the raw methods on
}

// BarLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BarLibraryCallerRaw struct {
	Contract *BarLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// BarLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BarLibraryTransactorRaw struct {
	Contract *BarLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBarLibrary creates a new instance of BarLibrary, bound to a specific deployed contract.
func NewBarLibrary(address common.Address, backend bind.ContractBackend) (*BarLibrary, error) {
	contract, err := bindBarLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BarLibrary{BarLibraryCaller: BarLibraryCaller{contract: contract}, BarLibraryTransactor: BarLibraryTransactor{contract: contract}, BarLibraryFilterer: BarLibraryFilterer{contract: contract}}, nil
}

// NewBarLibraryCaller creates a new read-only instance of BarLibrary, bound to a specific deployed contract.
func NewBarLibraryCaller(address common.Address, caller bind.ContractCaller) (*BarLibraryCaller, error) {
	contract, err := bindBarLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BarLibraryCaller{contract: contract}, nil
}

// NewBarLibraryTransactor creates a new write-only instance of BarLibrary, bound to a specific deployed contract.
func NewBarLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*BarLibraryTransactor, error) {
	contract, err := bindBarLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BarLibraryTransactor{contract: contract}, nil
}

// NewBarLibraryFilterer creates a new log filterer instance of BarLibrary, bound to a specific deployed contract.
func NewBarLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*BarLibraryFilterer, error) {
	contract, err := bindBarLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BarLibraryFilterer{contract: contract}, nil
}

// bindBarLibrary binds a generic wrapper to an already deployed contract.
func bindBarLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BarLibraryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BarLibrary *BarLibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BarLibrary.Contract.BarLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BarLibrary *BarLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BarLibrary.Contract.BarLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BarLibrary *BarLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BarLibrary.Contract.BarLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BarLibrary *BarLibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BarLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BarLibrary *BarLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BarLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BarLibrary *BarLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BarLibrary.Contract.contract.Transact(opts, method, params...)
}

// FooV2MetaData contains all meta data concerning the FooV2 contract.
var FooV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BAR_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FOO_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFoo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setBar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setFoo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506103d18061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061007a575f3560e01c80634321d9d2116100585780634321d9d2146100c957806367a23d13146100f05780638129fc1c146100f8578063dc80035d14610100575f80fd5b8063243dc8da1461007e578063352d3fba146100a057806337c7d9fb146100b5575b5f80fd5b5f8051602061037c833981519152545b60405190815260200160405180910390f35b6100b36100ae36600461034d565b61011f565b005b61008e5f8051602061037c83398151915281565b61008e7fab087b855a66dd5f7ac611eac86e092a2f3fc08a745c96afb2bc0553ff79ea0081565b61008e61019d565b6100b3610226565b6100b361010e36600461034d565b5f8051602061037c83398151915255565b5f5f8051602061037c833981519152604051635a55a96b60e01b81526001820160048201526024810184905290915073__$2d409e9966c58daad6cac335ac62f53e24$__90635a55a96b906044015f6040518083038186803b158015610183575f80fd5b505af4158015610195573d5f803e3d5ffd5b505050505050565b5f5f8051602061037c83398151915260405163a050083160e01b81526001909101600482015273__$2d409e9966c58daad6cac335ac62f53e24$__9063a050083190602401602060405180830381865af41580156101fd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102219190610364565b905090565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff16806102705750805467ffffffffffffffff808416911610155b1561028e5760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff191667ffffffffffffffff831617600160401b1781555f6102d87fab087b855a66dd5f7ac611eac86e092a2f3fc08a745c96afb2bc0553ff79ea0090565b547f92c6022ab2e4d8395662f3b814d5b9a489d615215f1269fc366f81bf3a17f7015550805468ff0000000000000000191681556040805167ffffffffffffffff8416815290517fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29181900360200190a15050565b5f6020828403121561035d575f80fd5b5035919050565b5f60208284031215610374575f80fd5b505191905056fe92c6022ab2e4d8395662f3b814d5b9a489d615215f1269fc366f81bf3a17f700a26469706673582212205886f725e0c9d1b63f41e6be295eaa7481a7f255e512bbdee62ef87b34d6f88264736f6c63430008190033",
}

// FooV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use FooV2MetaData.ABI instead.
var FooV2ABI = FooV2MetaData.ABI

// FooV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FooV2MetaData.Bin instead.
var FooV2Bin = FooV2MetaData.Bin

// DeployFooV2 deploys a new Ethereum contract, binding an instance of FooV2 to it.
func DeployFooV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FooV2, error) {
	parsed, err := FooV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	barLibraryAddr, _, _, _ := DeployBarLibrary(auth, backend)
	FooV2Bin = strings.ReplaceAll(FooV2Bin, "__$2d409e9966c58daad6cac335ac62f53e24$__", barLibraryAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FooV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FooV2{FooV2Caller: FooV2Caller{contract: contract}, FooV2Transactor: FooV2Transactor{contract: contract}, FooV2Filterer: FooV2Filterer{contract: contract}}, nil
}

// FooV2 is an auto generated Go binding around an Ethereum contract.
type FooV2 struct {
	FooV2Caller     // Read-only binding to the contract
	FooV2Transactor // Write-only binding to the contract
	FooV2Filterer   // Log filterer for contract events
}

// FooV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type FooV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FooV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FooV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FooV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FooV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FooV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FooV2Session struct {
	Contract     *FooV2            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FooV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FooV2CallerSession struct {
	Contract *FooV2Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FooV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FooV2TransactorSession struct {
	Contract     *FooV2Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FooV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type FooV2Raw struct {
	Contract *FooV2 // Generic contract binding to access the raw methods on
}

// FooV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FooV2CallerRaw struct {
	Contract *FooV2Caller // Generic read-only contract binding to access the raw methods on
}

// FooV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FooV2TransactorRaw struct {
	Contract *FooV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFooV2 creates a new instance of FooV2, bound to a specific deployed contract.
func NewFooV2(address common.Address, backend bind.ContractBackend) (*FooV2, error) {
	contract, err := bindFooV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FooV2{FooV2Caller: FooV2Caller{contract: contract}, FooV2Transactor: FooV2Transactor{contract: contract}, FooV2Filterer: FooV2Filterer{contract: contract}}, nil
}

// NewFooV2Caller creates a new read-only instance of FooV2, bound to a specific deployed contract.
func NewFooV2Caller(address common.Address, caller bind.ContractCaller) (*FooV2Caller, error) {
	contract, err := bindFooV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FooV2Caller{contract: contract}, nil
}

// NewFooV2Transactor creates a new write-only instance of FooV2, bound to a specific deployed contract.
func NewFooV2Transactor(address common.Address, transactor bind.ContractTransactor) (*FooV2Transactor, error) {
	contract, err := bindFooV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FooV2Transactor{contract: contract}, nil
}

// NewFooV2Filterer creates a new log filterer instance of FooV2, bound to a specific deployed contract.
func NewFooV2Filterer(address common.Address, filterer bind.ContractFilterer) (*FooV2Filterer, error) {
	contract, err := bindFooV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FooV2Filterer{contract: contract}, nil
}

// bindFooV2 binds a generic wrapper to an already deployed contract.
func bindFooV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FooV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FooV2 *FooV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FooV2.Contract.FooV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FooV2 *FooV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FooV2.Contract.FooV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FooV2 *FooV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FooV2.Contract.FooV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FooV2 *FooV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FooV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FooV2 *FooV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FooV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FooV2 *FooV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FooV2.Contract.contract.Transact(opts, method, params...)
}

// BARSTORAGELOCATION is a free data retrieval call binding the contract method 0x4321d9d2.
//
// Solidity: function BAR_STORAGE_LOCATION() view returns(bytes32)
func (_FooV2 *FooV2Caller) BARSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FooV2.contract.Call(opts, &out, "BAR_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BARSTORAGELOCATION is a free data retrieval call binding the contract method 0x4321d9d2.
//
// Solidity: function BAR_STORAGE_LOCATION() view returns(bytes32)
func (_FooV2 *FooV2Session) BARSTORAGELOCATION() ([32]byte, error) {
	return _FooV2.Contract.BARSTORAGELOCATION(&_FooV2.CallOpts)
}

// BARSTORAGELOCATION is a free data retrieval call binding the contract method 0x4321d9d2.
//
// Solidity: function BAR_STORAGE_LOCATION() view returns(bytes32)
func (_FooV2 *FooV2CallerSession) BARSTORAGELOCATION() ([32]byte, error) {
	return _FooV2.Contract.BARSTORAGELOCATION(&_FooV2.CallOpts)
}

// FOOSTORAGELOCATION is a free data retrieval call binding the contract method 0x37c7d9fb.
//
// Solidity: function FOO_STORAGE_LOCATION() view returns(bytes32)
func (_FooV2 *FooV2Caller) FOOSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FooV2.contract.Call(opts, &out, "FOO_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FOOSTORAGELOCATION is a free data retrieval call binding the contract method 0x37c7d9fb.
//
// Solidity: function FOO_STORAGE_LOCATION() view returns(bytes32)
func (_FooV2 *FooV2Session) FOOSTORAGELOCATION() ([32]byte, error) {
	return _FooV2.Contract.FOOSTORAGELOCATION(&_FooV2.CallOpts)
}

// FOOSTORAGELOCATION is a free data retrieval call binding the contract method 0x37c7d9fb.
//
// Solidity: function FOO_STORAGE_LOCATION() view returns(bytes32)
func (_FooV2 *FooV2CallerSession) FOOSTORAGELOCATION() ([32]byte, error) {
	return _FooV2.Contract.FOOSTORAGELOCATION(&_FooV2.CallOpts)
}

// GetBar is a free data retrieval call binding the contract method 0x67a23d13.
//
// Solidity: function getBar() view returns(uint256)
func (_FooV2 *FooV2Caller) GetBar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FooV2.contract.Call(opts, &out, "getBar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBar is a free data retrieval call binding the contract method 0x67a23d13.
//
// Solidity: function getBar() view returns(uint256)
func (_FooV2 *FooV2Session) GetBar() (*big.Int, error) {
	return _FooV2.Contract.GetBar(&_FooV2.CallOpts)
}

// GetBar is a free data retrieval call binding the contract method 0x67a23d13.
//
// Solidity: function getBar() view returns(uint256)
func (_FooV2 *FooV2CallerSession) GetBar() (*big.Int, error) {
	return _FooV2.Contract.GetBar(&_FooV2.CallOpts)
}

// GetFoo is a free data retrieval call binding the contract method 0x243dc8da.
//
// Solidity: function getFoo() view returns(uint256)
func (_FooV2 *FooV2Caller) GetFoo(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FooV2.contract.Call(opts, &out, "getFoo")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFoo is a free data retrieval call binding the contract method 0x243dc8da.
//
// Solidity: function getFoo() view returns(uint256)
func (_FooV2 *FooV2Session) GetFoo() (*big.Int, error) {
	return _FooV2.Contract.GetFoo(&_FooV2.CallOpts)
}

// GetFoo is a free data retrieval call binding the contract method 0x243dc8da.
//
// Solidity: function getFoo() view returns(uint256)
func (_FooV2 *FooV2CallerSession) GetFoo() (*big.Int, error) {
	return _FooV2.Contract.GetFoo(&_FooV2.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FooV2 *FooV2Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FooV2.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FooV2 *FooV2Session) Initialize() (*types.Transaction, error) {
	return _FooV2.Contract.Initialize(&_FooV2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FooV2 *FooV2TransactorSession) Initialize() (*types.Transaction, error) {
	return _FooV2.Contract.Initialize(&_FooV2.TransactOpts)
}

// SetBar is a paid mutator transaction binding the contract method 0x352d3fba.
//
// Solidity: function setBar(uint256 value) returns()
func (_FooV2 *FooV2Transactor) SetBar(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _FooV2.contract.Transact(opts, "setBar", value)
}

// SetBar is a paid mutator transaction binding the contract method 0x352d3fba.
//
// Solidity: function setBar(uint256 value) returns()
func (_FooV2 *FooV2Session) SetBar(value *big.Int) (*types.Transaction, error) {
	return _FooV2.Contract.SetBar(&_FooV2.TransactOpts, value)
}

// SetBar is a paid mutator transaction binding the contract method 0x352d3fba.
//
// Solidity: function setBar(uint256 value) returns()
func (_FooV2 *FooV2TransactorSession) SetBar(value *big.Int) (*types.Transaction, error) {
	return _FooV2.Contract.SetBar(&_FooV2.TransactOpts, value)
}

// SetFoo is a paid mutator transaction binding the contract method 0xdc80035d.
//
// Solidity: function setFoo(uint256 value) returns()
func (_FooV2 *FooV2Transactor) SetFoo(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _FooV2.contract.Transact(opts, "setFoo", value)
}

// SetFoo is a paid mutator transaction binding the contract method 0xdc80035d.
//
// Solidity: function setFoo(uint256 value) returns()
func (_FooV2 *FooV2Session) SetFoo(value *big.Int) (*types.Transaction, error) {
	return _FooV2.Contract.SetFoo(&_FooV2.TransactOpts, value)
}

// SetFoo is a paid mutator transaction binding the contract method 0xdc80035d.
//
// Solidity: function setFoo(uint256 value) returns()
func (_FooV2 *FooV2TransactorSession) SetFoo(value *big.Int) (*types.Transaction, error) {
	return _FooV2.Contract.SetFoo(&_FooV2.TransactOpts, value)
}

// FooV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FooV2 contract.
type FooV2InitializedIterator struct {
	Event *FooV2Initialized // Event containing the contract specifics and raw log

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
func (it *FooV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FooV2Initialized)
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
		it.Event = new(FooV2Initialized)
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
func (it *FooV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FooV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FooV2Initialized represents a Initialized event raised by the FooV2 contract.
type FooV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FooV2 *FooV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*FooV2InitializedIterator, error) {

	logs, sub, err := _FooV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FooV2InitializedIterator{contract: _FooV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FooV2 *FooV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FooV2Initialized) (event.Subscription, error) {

	logs, sub, err := _FooV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FooV2Initialized)
				if err := _FooV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FooV2 *FooV2Filterer) ParseInitialized(log types.Log) (*FooV2Initialized, error) {
	event := new(FooV2Initialized)
	if err := _FooV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
