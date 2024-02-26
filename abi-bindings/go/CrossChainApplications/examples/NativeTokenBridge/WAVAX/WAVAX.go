// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wavax

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

// WAVAXMetaData contains all meta data concerning the WAVAX contract.
var WAVAXMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYMBOL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b506107b5806100206000396000f3fe6080604052600436106100a05760003560e01c806370a082311161006457806370a0823114610170578063a3f4df7e1461019d578063a9059cbb146101e2578063d0e30db0146100af578063dd62ed3e14610202578063f76f8d781461023a576100af565b8063095ea7b3146100b757806318160ddd146100ec57806323b872dd146101095780632e0f2625146101295780632e1a7d4d14610150576100af565b366100af576100ad61026b565b005b6100ad61026b565b3480156100c357600080fd5b506100d76100d2366004610628565b6102c6565b60405190151581526020015b60405180910390f35b3480156100f857600080fd5b50475b6040519081526020016100e3565b34801561011557600080fd5b506100d7610124366004610652565b610333565b34801561013557600080fd5b5061013e601281565b60405160ff90911681526020016100e3565b34801561015c57600080fd5b506100ad61016b36600461068e565b61050f565b34801561017c57600080fd5b506100fb61018b3660046106a7565b60006020819052908152604090205481565b3480156101a957600080fd5b506101d56040518060400160405280600c81526020016b0aee4c2e0e0cac84082ac82b60a31b81525081565b6040516100e391906106c2565b3480156101ee57600080fd5b506100d76101fd366004610628565b6105f8565b34801561020e57600080fd5b506100fb61021d366004610710565b600160209081526000928352604080842090915290825290205481565b34801561024657600080fd5b506101d5604051806040016040528060058152602001640ae82ac82b60db1b81525081565b336000908152602081905260408120805434929061028a908490610759565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b3360008181526001602090815260408083206001600160a01b038716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103219086815260200190565b60405180910390a35060015b92915050565b6001600160a01b0383166000908152602081905260408120548211156103a05760405162461bcd60e51b815260206004820152601b60248201527f57415641583a20696e73756666696369656e742062616c616e6365000000000060448201526064015b60405180910390fd5b6001600160a01b038416331461045c576001600160a01b03841660009081526001602090815260408083203384529091529020548211156104235760405162461bcd60e51b815260206004820152601d60248201527f57415641583a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610397565b6001600160a01b03841660009081526001602090815260408083203384529091528120805484929061045690849061076c565b90915550505b6001600160a01b0384166000908152602081905260408120805484929061048490849061076c565b90915550506001600160a01b038316600090815260208190526040812080548492906104b1908490610759565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516104fd91815260200190565b60405180910390a35060019392505050565b3360009081526020819052604090205481111561056e5760405162461bcd60e51b815260206004820152601b60248201527f57415641583a20696e73756666696369656e742062616c616e636500000000006044820152606401610397565b336000908152602081905260408120805483929061058d90849061076c565b9091555050604051339082156108fc029083906000818181858888f193505050501580156105bf573d6000803e3d6000fd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b6000610605338484610333565b9392505050565b80356001600160a01b038116811461062357600080fd5b919050565b6000806040838503121561063b57600080fd5b6106448361060c565b946020939093013593505050565b60008060006060848603121561066757600080fd5b6106708461060c565b925061067e6020850161060c565b9150604084013590509250925092565b6000602082840312156106a057600080fd5b5035919050565b6000602082840312156106b957600080fd5b6106058261060c565b600060208083528351808285015260005b818110156106ef578581018301518582016040015282016106d3565b506000604082860101526040601f19601f8301168501019250505092915050565b6000806040838503121561072357600080fd5b61072c8361060c565b915061073a6020840161060c565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561032d5761032d610743565b8181038181111561032d5761032d61074356fea2646970667358221220c14b365384a908763eced2192d27045b4234f464a9b5db322339f033c704d78964736f6c63430008120033",
}

// WAVAXABI is the input ABI used to generate the binding from.
// Deprecated: Use WAVAXMetaData.ABI instead.
var WAVAXABI = WAVAXMetaData.ABI

// WAVAXBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WAVAXMetaData.Bin instead.
var WAVAXBin = WAVAXMetaData.Bin

// DeployWAVAX deploys a new Ethereum contract, binding an instance of WAVAX to it.
func DeployWAVAX(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WAVAX, error) {
	parsed, err := WAVAXMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WAVAXBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WAVAX{WAVAXCaller: WAVAXCaller{contract: contract}, WAVAXTransactor: WAVAXTransactor{contract: contract}, WAVAXFilterer: WAVAXFilterer{contract: contract}}, nil
}

// WAVAX is an auto generated Go binding around an Ethereum contract.
type WAVAX struct {
	WAVAXCaller     // Read-only binding to the contract
	WAVAXTransactor // Write-only binding to the contract
	WAVAXFilterer   // Log filterer for contract events
}

// WAVAXCaller is an auto generated read-only Go binding around an Ethereum contract.
type WAVAXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WAVAXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WAVAXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WAVAXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WAVAXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WAVAXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WAVAXSession struct {
	Contract     *WAVAX            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WAVAXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WAVAXCallerSession struct {
	Contract *WAVAXCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WAVAXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WAVAXTransactorSession struct {
	Contract     *WAVAXTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WAVAXRaw is an auto generated low-level Go binding around an Ethereum contract.
type WAVAXRaw struct {
	Contract *WAVAX // Generic contract binding to access the raw methods on
}

// WAVAXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WAVAXCallerRaw struct {
	Contract *WAVAXCaller // Generic read-only contract binding to access the raw methods on
}

// WAVAXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WAVAXTransactorRaw struct {
	Contract *WAVAXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWAVAX creates a new instance of WAVAX, bound to a specific deployed contract.
func NewWAVAX(address common.Address, backend bind.ContractBackend) (*WAVAX, error) {
	contract, err := bindWAVAX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WAVAX{WAVAXCaller: WAVAXCaller{contract: contract}, WAVAXTransactor: WAVAXTransactor{contract: contract}, WAVAXFilterer: WAVAXFilterer{contract: contract}}, nil
}

// NewWAVAXCaller creates a new read-only instance of WAVAX, bound to a specific deployed contract.
func NewWAVAXCaller(address common.Address, caller bind.ContractCaller) (*WAVAXCaller, error) {
	contract, err := bindWAVAX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WAVAXCaller{contract: contract}, nil
}

// NewWAVAXTransactor creates a new write-only instance of WAVAX, bound to a specific deployed contract.
func NewWAVAXTransactor(address common.Address, transactor bind.ContractTransactor) (*WAVAXTransactor, error) {
	contract, err := bindWAVAX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WAVAXTransactor{contract: contract}, nil
}

// NewWAVAXFilterer creates a new log filterer instance of WAVAX, bound to a specific deployed contract.
func NewWAVAXFilterer(address common.Address, filterer bind.ContractFilterer) (*WAVAXFilterer, error) {
	contract, err := bindWAVAX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WAVAXFilterer{contract: contract}, nil
}

// bindWAVAX binds a generic wrapper to an already deployed contract.
func bindWAVAX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WAVAXMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WAVAX *WAVAXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WAVAX.Contract.WAVAXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WAVAX *WAVAXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WAVAX.Contract.WAVAXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WAVAX *WAVAXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WAVAX.Contract.WAVAXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WAVAX *WAVAXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WAVAX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WAVAX *WAVAXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WAVAX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WAVAX *WAVAXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WAVAX.Contract.contract.Transact(opts, method, params...)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_WAVAX *WAVAXCaller) DECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WAVAX.contract.Call(opts, &out, "DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_WAVAX *WAVAXSession) DECIMALS() (uint8, error) {
	return _WAVAX.Contract.DECIMALS(&_WAVAX.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_WAVAX *WAVAXCallerSession) DECIMALS() (uint8, error) {
	return _WAVAX.Contract.DECIMALS(&_WAVAX.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_WAVAX *WAVAXCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WAVAX.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_WAVAX *WAVAXSession) NAME() (string, error) {
	return _WAVAX.Contract.NAME(&_WAVAX.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_WAVAX *WAVAXCallerSession) NAME() (string, error) {
	return _WAVAX.Contract.NAME(&_WAVAX.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() view returns(string)
func (_WAVAX *WAVAXCaller) SYMBOL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WAVAX.contract.Call(opts, &out, "SYMBOL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() view returns(string)
func (_WAVAX *WAVAXSession) SYMBOL() (string, error) {
	return _WAVAX.Contract.SYMBOL(&_WAVAX.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() view returns(string)
func (_WAVAX *WAVAXCallerSession) SYMBOL() (string, error) {
	return _WAVAX.Contract.SYMBOL(&_WAVAX.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WAVAX *WAVAXCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WAVAX.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WAVAX *WAVAXSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WAVAX.Contract.Allowance(&_WAVAX.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WAVAX *WAVAXCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WAVAX.Contract.Allowance(&_WAVAX.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WAVAX *WAVAXCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WAVAX.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WAVAX *WAVAXSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WAVAX.Contract.BalanceOf(&_WAVAX.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WAVAX *WAVAXCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WAVAX.Contract.BalanceOf(&_WAVAX.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WAVAX *WAVAXCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WAVAX.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WAVAX *WAVAXSession) TotalSupply() (*big.Int, error) {
	return _WAVAX.Contract.TotalSupply(&_WAVAX.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WAVAX *WAVAXCallerSession) TotalSupply() (*big.Int, error) {
	return _WAVAX.Contract.TotalSupply(&_WAVAX.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address addr, uint256 amount) returns(bool)
func (_WAVAX *WAVAXTransactor) Approve(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WAVAX.contract.Transact(opts, "approve", addr, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address addr, uint256 amount) returns(bool)
func (_WAVAX *WAVAXSession) Approve(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WAVAX.Contract.Approve(&_WAVAX.TransactOpts, addr, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address addr, uint256 amount) returns(bool)
func (_WAVAX *WAVAXTransactorSession) Approve(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WAVAX.Contract.Approve(&_WAVAX.TransactOpts, addr, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WAVAX *WAVAXTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WAVAX.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WAVAX *WAVAXSession) Deposit() (*types.Transaction, error) {
	return _WAVAX.Contract.Deposit(&_WAVAX.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WAVAX *WAVAXTransactorSession) Deposit() (*types.Transaction, error) {
	return _WAVAX.Contract.Deposit(&_WAVAX.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address destination, uint256 amount) returns(bool)
func (_WAVAX *WAVAXTransactor) Transfer(opts *bind.TransactOpts, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WAVAX.contract.Transact(opts, "transfer", destination, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address destination, uint256 amount) returns(bool)
func (_WAVAX *WAVAXSession) Transfer(destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WAVAX.Contract.Transfer(&_WAVAX.TransactOpts, destination, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address destination, uint256 amount) returns(bool)
func (_WAVAX *WAVAXTransactorSession) Transfer(destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WAVAX.Contract.Transfer(&_WAVAX.TransactOpts, destination, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address source, address destination, uint256 amount) returns(bool)
func (_WAVAX *WAVAXTransactor) TransferFrom(opts *bind.TransactOpts, source common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WAVAX.contract.Transact(opts, "transferFrom", source, destination, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address source, address destination, uint256 amount) returns(bool)
func (_WAVAX *WAVAXSession) TransferFrom(source common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WAVAX.Contract.TransferFrom(&_WAVAX.TransactOpts, source, destination, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address source, address destination, uint256 amount) returns(bool)
func (_WAVAX *WAVAXTransactorSession) TransferFrom(source common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WAVAX.Contract.TransferFrom(&_WAVAX.TransactOpts, source, destination, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WAVAX *WAVAXTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WAVAX.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WAVAX *WAVAXSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WAVAX.Contract.Withdraw(&_WAVAX.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WAVAX *WAVAXTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WAVAX.Contract.Withdraw(&_WAVAX.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WAVAX *WAVAXTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WAVAX.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WAVAX *WAVAXSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WAVAX.Contract.Fallback(&_WAVAX.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WAVAX *WAVAXTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WAVAX.Contract.Fallback(&_WAVAX.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WAVAX *WAVAXTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WAVAX.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WAVAX *WAVAXSession) Receive() (*types.Transaction, error) {
	return _WAVAX.Contract.Receive(&_WAVAX.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WAVAX *WAVAXTransactorSession) Receive() (*types.Transaction, error) {
	return _WAVAX.Contract.Receive(&_WAVAX.TransactOpts)
}

// WAVAXApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WAVAX contract.
type WAVAXApprovalIterator struct {
	Event *WAVAXApproval // Event containing the contract specifics and raw log

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
func (it *WAVAXApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WAVAXApproval)
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
		it.Event = new(WAVAXApproval)
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
func (it *WAVAXApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WAVAXApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WAVAXApproval represents a Approval event raised by the WAVAX contract.
type WAVAXApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WAVAX *WAVAXFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WAVAXApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WAVAX.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WAVAXApprovalIterator{contract: _WAVAX.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WAVAX *WAVAXFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WAVAXApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WAVAX.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WAVAXApproval)
				if err := _WAVAX.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WAVAX *WAVAXFilterer) ParseApproval(log types.Log) (*WAVAXApproval, error) {
	event := new(WAVAXApproval)
	if err := _WAVAX.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WAVAXDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WAVAX contract.
type WAVAXDepositIterator struct {
	Event *WAVAXDeposit // Event containing the contract specifics and raw log

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
func (it *WAVAXDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WAVAXDeposit)
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
		it.Event = new(WAVAXDeposit)
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
func (it *WAVAXDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WAVAXDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WAVAXDeposit represents a Deposit event raised by the WAVAX contract.
type WAVAXDeposit struct {
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed destination, uint256 amount)
func (_WAVAX *WAVAXFilterer) FilterDeposit(opts *bind.FilterOpts, destination []common.Address) (*WAVAXDepositIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _WAVAX.contract.FilterLogs(opts, "Deposit", destinationRule)
	if err != nil {
		return nil, err
	}
	return &WAVAXDepositIterator{contract: _WAVAX.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed destination, uint256 amount)
func (_WAVAX *WAVAXFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WAVAXDeposit, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _WAVAX.contract.WatchLogs(opts, "Deposit", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WAVAXDeposit)
				if err := _WAVAX.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed destination, uint256 amount)
func (_WAVAX *WAVAXFilterer) ParseDeposit(log types.Log) (*WAVAXDeposit, error) {
	event := new(WAVAXDeposit)
	if err := _WAVAX.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WAVAXTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WAVAX contract.
type WAVAXTransferIterator struct {
	Event *WAVAXTransfer // Event containing the contract specifics and raw log

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
func (it *WAVAXTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WAVAXTransfer)
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
		it.Event = new(WAVAXTransfer)
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
func (it *WAVAXTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WAVAXTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WAVAXTransfer represents a Transfer event raised by the WAVAX contract.
type WAVAXTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WAVAX *WAVAXFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WAVAXTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WAVAX.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WAVAXTransferIterator{contract: _WAVAX.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WAVAX *WAVAXFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WAVAXTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WAVAX.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WAVAXTransfer)
				if err := _WAVAX.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WAVAX *WAVAXFilterer) ParseTransfer(log types.Log) (*WAVAXTransfer, error) {
	event := new(WAVAXTransfer)
	if err := _WAVAX.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WAVAXWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WAVAX contract.
type WAVAXWithdrawalIterator struct {
	Event *WAVAXWithdrawal // Event containing the contract specifics and raw log

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
func (it *WAVAXWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WAVAXWithdrawal)
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
		it.Event = new(WAVAXWithdrawal)
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
func (it *WAVAXWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WAVAXWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WAVAXWithdrawal represents a Withdrawal event raised by the WAVAX contract.
type WAVAXWithdrawal struct {
	Source common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed source, uint256 amount)
func (_WAVAX *WAVAXFilterer) FilterWithdrawal(opts *bind.FilterOpts, source []common.Address) (*WAVAXWithdrawalIterator, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _WAVAX.contract.FilterLogs(opts, "Withdrawal", sourceRule)
	if err != nil {
		return nil, err
	}
	return &WAVAXWithdrawalIterator{contract: _WAVAX.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed source, uint256 amount)
func (_WAVAX *WAVAXFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WAVAXWithdrawal, source []common.Address) (event.Subscription, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _WAVAX.contract.WatchLogs(opts, "Withdrawal", sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WAVAXWithdrawal)
				if err := _WAVAX.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed source, uint256 amount)
func (_WAVAX *WAVAXFilterer) ParseWithdrawal(log types.Log) (*WAVAXWithdrawal, error) {
	event := new(WAVAXWithdrawal)
	if err := _WAVAX.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
