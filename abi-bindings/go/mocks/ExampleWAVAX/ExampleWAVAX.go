// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package examplewavax

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

// ExampleWAVAXMetaData contains all meta data concerning the ExampleWAVAX contract.
var ExampleWAVAXMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c0604052600c60809081526b0aee4c2e0e0cac84082ac82b60a31b60a05260009061002b9082610114565b506040805180820190915260058152640ae82ac82b60db1b60208201526001906100559082610114565b506002805460ff1916601217905534801561006f57600080fd5b506101d3565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061009f57607f821691505b6020821081036100bf57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561010f57600081815260208120601f850160051c810160208610156100ec5750805b601f850160051c820191505b8181101561010b578281556001016100f8565b5050505b505050565b81516001600160401b0381111561012d5761012d610075565b6101418161013b845461008b565b846100c5565b602080601f831160018114610176576000841561015e5750858301515b600019600386901b1c1916600185901b17855561010b565b600085815260208120601f198616915b828110156101a557888601518255948401946001909101908401610186565b50858210156101c35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610846806101e26000396000f3fe6080604052600436106100a05760003560e01c8063313ce56711610064578063313ce5671461016f57806370a082311461019b57806395d89b41146101c8578063a9059cbb146101dd578063d0e30db0146100af578063dd62ed3e146101fd576100af565b806306fdde03146100b7578063095ea7b3146100e257806318160ddd1461011257806323b872dd1461012f5780632e1a7d4d1461014f576100af565b366100af576100ad610235565b005b6100ad610235565b3480156100c357600080fd5b506100cc610290565b6040516100d99190610621565b60405180910390f35b3480156100ee57600080fd5b506101026100fd36600461068b565b61031e565b60405190151581526020016100d9565b34801561011e57600080fd5b50475b6040519081526020016100d9565b34801561013b57600080fd5b5061010261014a3660046106b5565b61038b565b34801561015b57600080fd5b506100ad61016a3660046106f1565b610547565b34801561017b57600080fd5b506002546101899060ff1681565b60405160ff90911681526020016100d9565b3480156101a757600080fd5b506101216101b636600461070a565b60036020526000908152604090205481565b3480156101d457600080fd5b506100cc610600565b3480156101e957600080fd5b506101026101f836600461068b565b61060d565b34801561020957600080fd5b50610121610218366004610725565b600460209081526000928352604080842090915290825290205481565b336000908152600360205260408120805434929061025490849061076e565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b6000805461029d90610781565b80601f01602080910402602001604051908101604052809291908181526020018280546102c990610781565b80156103165780601f106102eb57610100808354040283529160200191610316565b820191906000526020600020905b8154815290600101906020018083116102f957829003601f168201915b505050505081565b3360008181526004602090815260408083206001600160a01b038716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103799086815260200190565b60405180910390a35060015b92915050565b6001600160a01b0383166000908152600360205260408120548211156103cc5760405162461bcd60e51b81526004016103c3906107bb565b60405180910390fd5b6001600160a01b0384163314610494576001600160a01b038416600090815260046020908152604080832033845290915290205482111561045b5760405162461bcd60e51b8152602060048201526024808201527f4578616d706c6557415641583a20696e73756666696369656e7420616c6c6f77604482015263616e636560e01b60648201526084016103c3565b6001600160a01b03841660009081526004602090815260408083203384529091528120805484929061048e9084906107fd565b90915550505b6001600160a01b038416600090815260036020526040812080548492906104bc9084906107fd565b90915550506001600160a01b038316600090815260036020526040812080548492906104e990849061076e565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161053591815260200190565b60405180910390a35060019392505050565b336000908152600360205260409020548111156105765760405162461bcd60e51b81526004016103c3906107bb565b33600090815260036020526040812080548392906105959084906107fd565b9091555050604051339082156108fc029083906000818181858888f193505050501580156105c7573d6000803e3d6000fd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b6001805461029d90610781565b600061061a33848461038b565b9392505050565b600060208083528351808285015260005b8181101561064e57858101830151858201604001528201610632565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461068657600080fd5b919050565b6000806040838503121561069e57600080fd5b6106a78361066f565b946020939093013593505050565b6000806000606084860312156106ca57600080fd5b6106d38461066f565b92506106e16020850161066f565b9150604084013590509250925092565b60006020828403121561070357600080fd5b5035919050565b60006020828403121561071c57600080fd5b61061a8261066f565b6000806040838503121561073857600080fd5b6107418361066f565b915061074f6020840161066f565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561038557610385610758565b600181811c9082168061079557607f821691505b6020821081036107b557634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526022908201527f4578616d706c6557415641583a20696e73756666696369656e742062616c616e604082015261636560f01b606082015260800190565b818103818111156103855761038561075856fea26469706673582212206da8c514c3d694fcfb8b70946a320bb0c9142e349f8d3dac4693097d5f1cfd2664736f6c63430008120033",
}

// ExampleWAVAXABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleWAVAXMetaData.ABI instead.
var ExampleWAVAXABI = ExampleWAVAXMetaData.ABI

// ExampleWAVAXBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleWAVAXMetaData.Bin instead.
var ExampleWAVAXBin = ExampleWAVAXMetaData.Bin

// DeployExampleWAVAX deploys a new Ethereum contract, binding an instance of ExampleWAVAX to it.
func DeployExampleWAVAX(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExampleWAVAX, error) {
	parsed, err := ExampleWAVAXMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleWAVAXBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExampleWAVAX{ExampleWAVAXCaller: ExampleWAVAXCaller{contract: contract}, ExampleWAVAXTransactor: ExampleWAVAXTransactor{contract: contract}, ExampleWAVAXFilterer: ExampleWAVAXFilterer{contract: contract}}, nil
}

// ExampleWAVAX is an auto generated Go binding around an Ethereum contract.
type ExampleWAVAX struct {
	ExampleWAVAXCaller     // Read-only binding to the contract
	ExampleWAVAXTransactor // Write-only binding to the contract
	ExampleWAVAXFilterer   // Log filterer for contract events
}

// ExampleWAVAXCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleWAVAXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleWAVAXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleWAVAXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleWAVAXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleWAVAXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleWAVAXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleWAVAXSession struct {
	Contract     *ExampleWAVAX     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExampleWAVAXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleWAVAXCallerSession struct {
	Contract *ExampleWAVAXCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ExampleWAVAXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleWAVAXTransactorSession struct {
	Contract     *ExampleWAVAXTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExampleWAVAXRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleWAVAXRaw struct {
	Contract *ExampleWAVAX // Generic contract binding to access the raw methods on
}

// ExampleWAVAXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleWAVAXCallerRaw struct {
	Contract *ExampleWAVAXCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleWAVAXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleWAVAXTransactorRaw struct {
	Contract *ExampleWAVAXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleWAVAX creates a new instance of ExampleWAVAX, bound to a specific deployed contract.
func NewExampleWAVAX(address common.Address, backend bind.ContractBackend) (*ExampleWAVAX, error) {
	contract, err := bindExampleWAVAX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleWAVAX{ExampleWAVAXCaller: ExampleWAVAXCaller{contract: contract}, ExampleWAVAXTransactor: ExampleWAVAXTransactor{contract: contract}, ExampleWAVAXFilterer: ExampleWAVAXFilterer{contract: contract}}, nil
}

// NewExampleWAVAXCaller creates a new read-only instance of ExampleWAVAX, bound to a specific deployed contract.
func NewExampleWAVAXCaller(address common.Address, caller bind.ContractCaller) (*ExampleWAVAXCaller, error) {
	contract, err := bindExampleWAVAX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleWAVAXCaller{contract: contract}, nil
}

// NewExampleWAVAXTransactor creates a new write-only instance of ExampleWAVAX, bound to a specific deployed contract.
func NewExampleWAVAXTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleWAVAXTransactor, error) {
	contract, err := bindExampleWAVAX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleWAVAXTransactor{contract: contract}, nil
}

// NewExampleWAVAXFilterer creates a new log filterer instance of ExampleWAVAX, bound to a specific deployed contract.
func NewExampleWAVAXFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleWAVAXFilterer, error) {
	contract, err := bindExampleWAVAX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleWAVAXFilterer{contract: contract}, nil
}

// bindExampleWAVAX binds a generic wrapper to an already deployed contract.
func bindExampleWAVAX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExampleWAVAXMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleWAVAX *ExampleWAVAXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleWAVAX.Contract.ExampleWAVAXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleWAVAX *ExampleWAVAXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.ExampleWAVAXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleWAVAX *ExampleWAVAXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.ExampleWAVAXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleWAVAX *ExampleWAVAXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleWAVAX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleWAVAX *ExampleWAVAXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleWAVAX *ExampleWAVAXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExampleWAVAX.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ExampleWAVAX.Contract.Allowance(&_ExampleWAVAX.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ExampleWAVAX.Contract.Allowance(&_ExampleWAVAX.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExampleWAVAX.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ExampleWAVAX.Contract.BalanceOf(&_ExampleWAVAX.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ExampleWAVAX.Contract.BalanceOf(&_ExampleWAVAX.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExampleWAVAX *ExampleWAVAXCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ExampleWAVAX.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExampleWAVAX *ExampleWAVAXSession) Decimals() (uint8, error) {
	return _ExampleWAVAX.Contract.Decimals(&_ExampleWAVAX.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExampleWAVAX *ExampleWAVAXCallerSession) Decimals() (uint8, error) {
	return _ExampleWAVAX.Contract.Decimals(&_ExampleWAVAX.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleWAVAX *ExampleWAVAXCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExampleWAVAX.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleWAVAX *ExampleWAVAXSession) Name() (string, error) {
	return _ExampleWAVAX.Contract.Name(&_ExampleWAVAX.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleWAVAX *ExampleWAVAXCallerSession) Name() (string, error) {
	return _ExampleWAVAX.Contract.Name(&_ExampleWAVAX.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleWAVAX *ExampleWAVAXCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExampleWAVAX.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleWAVAX *ExampleWAVAXSession) Symbol() (string, error) {
	return _ExampleWAVAX.Contract.Symbol(&_ExampleWAVAX.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleWAVAX *ExampleWAVAXCallerSession) Symbol() (string, error) {
	return _ExampleWAVAX.Contract.Symbol(&_ExampleWAVAX.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExampleWAVAX.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXSession) TotalSupply() (*big.Int, error) {
	return _ExampleWAVAX.Contract.TotalSupply(&_ExampleWAVAX.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXCallerSession) TotalSupply() (*big.Int, error) {
	return _ExampleWAVAX.Contract.TotalSupply(&_ExampleWAVAX.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.Approve(&_ExampleWAVAX.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.Approve(&_ExampleWAVAX.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ExampleWAVAX *ExampleWAVAXTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleWAVAX.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ExampleWAVAX *ExampleWAVAXSession) Deposit() (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.Deposit(&_ExampleWAVAX.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ExampleWAVAX *ExampleWAVAXTransactorSession) Deposit() (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.Deposit(&_ExampleWAVAX.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.Transfer(&_ExampleWAVAX.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.Transfer(&_ExampleWAVAX.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.TransferFrom(&_ExampleWAVAX.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.TransferFrom(&_ExampleWAVAX.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ExampleWAVAX *ExampleWAVAXTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ExampleWAVAX *ExampleWAVAXSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.Withdraw(&_ExampleWAVAX.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ExampleWAVAX *ExampleWAVAXTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.Withdraw(&_ExampleWAVAX.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ExampleWAVAX *ExampleWAVAXTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ExampleWAVAX.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ExampleWAVAX *ExampleWAVAXSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.Fallback(&_ExampleWAVAX.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ExampleWAVAX *ExampleWAVAXTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.Fallback(&_ExampleWAVAX.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ExampleWAVAX *ExampleWAVAXTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleWAVAX.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ExampleWAVAX *ExampleWAVAXSession) Receive() (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.Receive(&_ExampleWAVAX.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ExampleWAVAX *ExampleWAVAXTransactorSession) Receive() (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.Receive(&_ExampleWAVAX.TransactOpts)
}

// ExampleWAVAXApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ExampleWAVAX contract.
type ExampleWAVAXApprovalIterator struct {
	Event *ExampleWAVAXApproval // Event containing the contract specifics and raw log

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
func (it *ExampleWAVAXApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleWAVAXApproval)
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
		it.Event = new(ExampleWAVAXApproval)
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
func (it *ExampleWAVAXApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleWAVAXApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleWAVAXApproval represents a Approval event raised by the ExampleWAVAX contract.
type ExampleWAVAXApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ExampleWAVAX *ExampleWAVAXFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ExampleWAVAXApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ExampleWAVAX.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ExampleWAVAXApprovalIterator{contract: _ExampleWAVAX.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ExampleWAVAX *ExampleWAVAXFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ExampleWAVAXApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ExampleWAVAX.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleWAVAXApproval)
				if err := _ExampleWAVAX.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ExampleWAVAX *ExampleWAVAXFilterer) ParseApproval(log types.Log) (*ExampleWAVAXApproval, error) {
	event := new(ExampleWAVAXApproval)
	if err := _ExampleWAVAX.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleWAVAXDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ExampleWAVAX contract.
type ExampleWAVAXDepositIterator struct {
	Event *ExampleWAVAXDeposit // Event containing the contract specifics and raw log

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
func (it *ExampleWAVAXDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleWAVAXDeposit)
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
		it.Event = new(ExampleWAVAXDeposit)
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
func (it *ExampleWAVAXDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleWAVAXDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleWAVAXDeposit represents a Deposit event raised by the ExampleWAVAX contract.
type ExampleWAVAXDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_ExampleWAVAX *ExampleWAVAXFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*ExampleWAVAXDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExampleWAVAX.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &ExampleWAVAXDepositIterator{contract: _ExampleWAVAX.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_ExampleWAVAX *ExampleWAVAXFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ExampleWAVAXDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExampleWAVAX.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleWAVAXDeposit)
				if err := _ExampleWAVAX.contract.UnpackLog(event, "Deposit", log); err != nil {
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
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_ExampleWAVAX *ExampleWAVAXFilterer) ParseDeposit(log types.Log) (*ExampleWAVAXDeposit, error) {
	event := new(ExampleWAVAXDeposit)
	if err := _ExampleWAVAX.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleWAVAXTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ExampleWAVAX contract.
type ExampleWAVAXTransferIterator struct {
	Event *ExampleWAVAXTransfer // Event containing the contract specifics and raw log

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
func (it *ExampleWAVAXTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleWAVAXTransfer)
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
		it.Event = new(ExampleWAVAXTransfer)
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
func (it *ExampleWAVAXTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleWAVAXTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleWAVAXTransfer represents a Transfer event raised by the ExampleWAVAX contract.
type ExampleWAVAXTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ExampleWAVAX *ExampleWAVAXFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExampleWAVAXTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExampleWAVAX.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExampleWAVAXTransferIterator{contract: _ExampleWAVAX.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ExampleWAVAX *ExampleWAVAXFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ExampleWAVAXTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExampleWAVAX.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleWAVAXTransfer)
				if err := _ExampleWAVAX.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ExampleWAVAX *ExampleWAVAXFilterer) ParseTransfer(log types.Log) (*ExampleWAVAXTransfer, error) {
	event := new(ExampleWAVAXTransfer)
	if err := _ExampleWAVAX.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleWAVAXWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ExampleWAVAX contract.
type ExampleWAVAXWithdrawalIterator struct {
	Event *ExampleWAVAXWithdrawal // Event containing the contract specifics and raw log

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
func (it *ExampleWAVAXWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleWAVAXWithdrawal)
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
		it.Event = new(ExampleWAVAXWithdrawal)
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
func (it *ExampleWAVAXWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleWAVAXWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleWAVAXWithdrawal represents a Withdrawal event raised by the ExampleWAVAX contract.
type ExampleWAVAXWithdrawal struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_ExampleWAVAX *ExampleWAVAXFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address) (*ExampleWAVAXWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExampleWAVAX.contract.FilterLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return &ExampleWAVAXWithdrawalIterator{contract: _ExampleWAVAX.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_ExampleWAVAX *ExampleWAVAXFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ExampleWAVAXWithdrawal, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExampleWAVAX.contract.WatchLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleWAVAXWithdrawal)
				if err := _ExampleWAVAX.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_ExampleWAVAX *ExampleWAVAXFilterer) ParseWithdrawal(log types.Log) (*ExampleWAVAXWithdrawal, error) {
	event := new(ExampleWAVAXWithdrawal)
	if err := _ExampleWAVAX.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
