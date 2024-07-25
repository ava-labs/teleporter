// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappednativetoken

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

// WrappedNativeTokenMetaData contains all meta data concerning the WrappedNativeToken contract.
var WrappedNativeTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000cb738038062000cb78339810160408190526200003491620000dc565b8060405160200162000047919062000194565b60405160208183030381529060405281604051602001620000699190620001c6565b60408051601f19818403018152919052600362000087838262000282565b50600462000096828262000282565b505050506200034e565b634e487b7160e01b600052604160045260246000fd5b60005b83811015620000d3578181015183820152602001620000b9565b50506000910152565b600060208284031215620000ef57600080fd5b81516001600160401b03808211156200010757600080fd5b818401915084601f8301126200011c57600080fd5b815181811115620001315762000131620000a0565b604051601f8201601f19908116603f011681019083821181831017156200015c576200015c620000a0565b816040528281528760208487010111156200017657600080fd5b62000189836020830160208801620000b6565b979650505050505050565b6702bb930b83832b2160c51b815260008251620001b9816008850160208701620000b6565b9190910160080192915050565b605760f81b815260008251620001e4816001850160208701620000b6565b9190910160010192915050565b600181811c908216806200020657607f821691505b6020821081036200022757634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200027d576000816000526020600020601f850160051c81016020861015620002585750805b601f850160051c820191505b81811015620002795782815560010162000264565b5050505b505050565b81516001600160401b038111156200029e576200029e620000a0565b620002b681620002af8454620001f1565b846200022d565b602080601f831160018114620002ee5760008415620002d55750858301515b600019600386901b1c1916600185901b17855562000279565b600085815260208120601f198616915b828110156200031f57888601518255948401946001909101908401620002fe565b50858210156200033e5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610959806200035e6000396000f3fe6080604052600436106100a05760003560e01c8063313ce56711610064578063313ce5671461017157806370a082311461018d57806395d89b41146101c3578063a9059cbb146101d8578063d0e30db0146100af578063dd62ed3e146101f8576100af565b806306fdde03146100b7578063095ea7b3146100e257806318160ddd1461011257806323b872dd146101315780632e1a7d4d14610151576100af565b366100af576100ad61023e565b005b6100ad61023e565b3480156100c357600080fd5b506100cc61027f565b6040516100d991906107b2565b60405180910390f35b3480156100ee57600080fd5b506101026100fd36600461081d565b610311565b60405190151581526020016100d9565b34801561011e57600080fd5b506002545b6040519081526020016100d9565b34801561013d57600080fd5b5061010261014c366004610847565b61032b565b34801561015d57600080fd5b506100ad61016c366004610883565b61034f565b34801561017d57600080fd5b50604051601281526020016100d9565b34801561019957600080fd5b506101236101a836600461089c565b6001600160a01b031660009081526020819052604090205490565b3480156101cf57600080fd5b506100cc61039b565b3480156101e457600080fd5b506101026101f336600461081d565b6103aa565b34801561020457600080fd5b506101236102133660046108be565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b61024833346103b8565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b60606003805461028e906108f1565b80601f01602080910402602001604051908101604052809291908181526020018280546102ba906108f1565b80156103075780601f106102dc57610100808354040283529160200191610307565b820191906000526020600020905b8154815290600101906020018083116102ea57829003601f168201915b5050505050905090565b60003361031f8185856103f7565b60019150505b92915050565b600033610339858285610409565b610344858585610487565b506001949350505050565b61035933826104e6565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a2610398338261051c565b50565b60606004805461028e906108f1565b60003361031f818585610487565b6001600160a01b0382166103e75760405163ec442f0560e01b8152600060048201526024015b60405180910390fd5b6103f3600083836105b3565b5050565b61040483838360016106dd565b505050565b6001600160a01b038381166000908152600160209081526040808320938616835292905220546000198114610481578181101561047257604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016103de565b610481848484840360006106dd565b50505050565b6001600160a01b0383166104b157604051634b637e8f60e11b8152600060048201526024016103de565b6001600160a01b0382166104db5760405163ec442f0560e01b8152600060048201526024016103de565b6104048383836105b3565b6001600160a01b03821661051057604051634b637e8f60e11b8152600060048201526024016103de565b6103f3826000836105b3565b8047101561053f5760405163cd78605960e01b81523060048201526024016103de565b6000826001600160a01b03168260405160006040518083038185875af1925050503d806000811461058c576040519150601f19603f3d011682016040523d82523d6000602084013e610591565b606091505b505090508061040457604051630a12f52160e11b815260040160405180910390fd5b6001600160a01b0383166105de5780600260008282546105d3919061092b565b909155506106509050565b6001600160a01b038316600090815260208190526040902054818110156106315760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016103de565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b03821661066c5760028054829003905561068b565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516106d091815260200190565b60405180910390a3505050565b6001600160a01b0384166107075760405163e602df0560e01b8152600060048201526024016103de565b6001600160a01b03831661073157604051634a1406b160e11b8152600060048201526024016103de565b6001600160a01b038085166000908152600160209081526040808320938716835292905220829055801561048157826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516107a491815260200190565b60405180910390a350505050565b60006020808352835180602085015260005b818110156107e0578581018301518582016040015282016107c4565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461081857600080fd5b919050565b6000806040838503121561083057600080fd5b61083983610801565b946020939093013593505050565b60008060006060848603121561085c57600080fd5b61086584610801565b925061087360208501610801565b9150604084013590509250925092565b60006020828403121561089557600080fd5b5035919050565b6000602082840312156108ae57600080fd5b6108b782610801565b9392505050565b600080604083850312156108d157600080fd5b6108da83610801565b91506108e860208401610801565b90509250929050565b600181811c9082168061090557607f821691505b60208210810361092557634e487b7160e01b600052602260045260246000fd5b50919050565b8082018082111561032557634e487b7160e01b600052601160045260246000fdfea164736f6c6343000817000a",
}

// WrappedNativeTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use WrappedNativeTokenMetaData.ABI instead.
var WrappedNativeTokenABI = WrappedNativeTokenMetaData.ABI

// WrappedNativeTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WrappedNativeTokenMetaData.Bin instead.
var WrappedNativeTokenBin = WrappedNativeTokenMetaData.Bin

// DeployWrappedNativeToken deploys a new Ethereum contract, binding an instance of WrappedNativeToken to it.
func DeployWrappedNativeToken(auth *bind.TransactOpts, backend bind.ContractBackend, symbol string) (common.Address, *types.Transaction, *WrappedNativeToken, error) {
	parsed, err := WrappedNativeTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WrappedNativeTokenBin), backend, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WrappedNativeToken{WrappedNativeTokenCaller: WrappedNativeTokenCaller{contract: contract}, WrappedNativeTokenTransactor: WrappedNativeTokenTransactor{contract: contract}, WrappedNativeTokenFilterer: WrappedNativeTokenFilterer{contract: contract}}, nil
}

// WrappedNativeToken is an auto generated Go binding around an Ethereum contract.
type WrappedNativeToken struct {
	WrappedNativeTokenCaller     // Read-only binding to the contract
	WrappedNativeTokenTransactor // Write-only binding to the contract
	WrappedNativeTokenFilterer   // Log filterer for contract events
}

// WrappedNativeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrappedNativeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedNativeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrappedNativeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedNativeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrappedNativeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedNativeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrappedNativeTokenSession struct {
	Contract     *WrappedNativeToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WrappedNativeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrappedNativeTokenCallerSession struct {
	Contract *WrappedNativeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// WrappedNativeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrappedNativeTokenTransactorSession struct {
	Contract     *WrappedNativeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// WrappedNativeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrappedNativeTokenRaw struct {
	Contract *WrappedNativeToken // Generic contract binding to access the raw methods on
}

// WrappedNativeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrappedNativeTokenCallerRaw struct {
	Contract *WrappedNativeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// WrappedNativeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrappedNativeTokenTransactorRaw struct {
	Contract *WrappedNativeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrappedNativeToken creates a new instance of WrappedNativeToken, bound to a specific deployed contract.
func NewWrappedNativeToken(address common.Address, backend bind.ContractBackend) (*WrappedNativeToken, error) {
	contract, err := bindWrappedNativeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeToken{WrappedNativeTokenCaller: WrappedNativeTokenCaller{contract: contract}, WrappedNativeTokenTransactor: WrappedNativeTokenTransactor{contract: contract}, WrappedNativeTokenFilterer: WrappedNativeTokenFilterer{contract: contract}}, nil
}

// NewWrappedNativeTokenCaller creates a new read-only instance of WrappedNativeToken, bound to a specific deployed contract.
func NewWrappedNativeTokenCaller(address common.Address, caller bind.ContractCaller) (*WrappedNativeTokenCaller, error) {
	contract, err := bindWrappedNativeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenCaller{contract: contract}, nil
}

// NewWrappedNativeTokenTransactor creates a new write-only instance of WrappedNativeToken, bound to a specific deployed contract.
func NewWrappedNativeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*WrappedNativeTokenTransactor, error) {
	contract, err := bindWrappedNativeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenTransactor{contract: contract}, nil
}

// NewWrappedNativeTokenFilterer creates a new log filterer instance of WrappedNativeToken, bound to a specific deployed contract.
func NewWrappedNativeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*WrappedNativeTokenFilterer, error) {
	contract, err := bindWrappedNativeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenFilterer{contract: contract}, nil
}

// bindWrappedNativeToken binds a generic wrapper to an already deployed contract.
func bindWrappedNativeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WrappedNativeTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedNativeToken *WrappedNativeTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedNativeToken.Contract.WrappedNativeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedNativeToken *WrappedNativeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.WrappedNativeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedNativeToken *WrappedNativeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.WrappedNativeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedNativeToken *WrappedNativeTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedNativeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedNativeToken *WrappedNativeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedNativeToken *WrappedNativeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedNativeToken.Contract.Allowance(&_WrappedNativeToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedNativeToken.Contract.Allowance(&_WrappedNativeToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedNativeToken.Contract.BalanceOf(&_WrappedNativeToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedNativeToken.Contract.BalanceOf(&_WrappedNativeToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedNativeToken *WrappedNativeTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedNativeToken *WrappedNativeTokenSession) Decimals() (uint8, error) {
	return _WrappedNativeToken.Contract.Decimals(&_WrappedNativeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) Decimals() (uint8, error) {
	return _WrappedNativeToken.Contract.Decimals(&_WrappedNativeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenSession) Name() (string, error) {
	return _WrappedNativeToken.Contract.Name(&_WrappedNativeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) Name() (string, error) {
	return _WrappedNativeToken.Contract.Name(&_WrappedNativeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenSession) Symbol() (string, error) {
	return _WrappedNativeToken.Contract.Symbol(&_WrappedNativeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) Symbol() (string, error) {
	return _WrappedNativeToken.Contract.Symbol(&_WrappedNativeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrappedNativeToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenSession) TotalSupply() (*big.Int, error) {
	return _WrappedNativeToken.Contract.TotalSupply(&_WrappedNativeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedNativeToken *WrappedNativeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _WrappedNativeToken.Contract.TotalSupply(&_WrappedNativeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Approve(&_WrappedNativeToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Approve(&_WrappedNativeToken.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenSession) Deposit() (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Deposit(&_WrappedNativeToken.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Deposit() (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Deposit(&_WrappedNativeToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Transfer(&_WrappedNativeToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Transfer(&_WrappedNativeToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.TransferFrom(&_WrappedNativeToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.TransferFrom(&_WrappedNativeToken.TransactOpts, from, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WrappedNativeToken *WrappedNativeTokenSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Withdraw(&_WrappedNativeToken.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Withdraw(&_WrappedNativeToken.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Fallback(&_WrappedNativeToken.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Fallback(&_WrappedNativeToken.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenSession) Receive() (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Receive(&_WrappedNativeToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Receive() (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Receive(&_WrappedNativeToken.TransactOpts)
}

// WrappedNativeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WrappedNativeToken contract.
type WrappedNativeTokenApprovalIterator struct {
	Event *WrappedNativeTokenApproval // Event containing the contract specifics and raw log

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
func (it *WrappedNativeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeTokenApproval)
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
		it.Event = new(WrappedNativeTokenApproval)
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
func (it *WrappedNativeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeTokenApproval represents a Approval event raised by the WrappedNativeToken contract.
type WrappedNativeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WrappedNativeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenApprovalIterator{contract: _WrappedNativeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WrappedNativeTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeTokenApproval)
				if err := _WrappedNativeToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WrappedNativeToken *WrappedNativeTokenFilterer) ParseApproval(log types.Log) (*WrappedNativeTokenApproval, error) {
	event := new(WrappedNativeTokenApproval)
	if err := _WrappedNativeToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedNativeTokenDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WrappedNativeToken contract.
type WrappedNativeTokenDepositIterator struct {
	Event *WrappedNativeTokenDeposit // Event containing the contract specifics and raw log

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
func (it *WrappedNativeTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeTokenDeposit)
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
		it.Event = new(WrappedNativeTokenDeposit)
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
func (it *WrappedNativeTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeTokenDeposit represents a Deposit event raised by the WrappedNativeToken contract.
type WrappedNativeTokenDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*WrappedNativeTokenDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenDepositIterator{contract: _WrappedNativeToken.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WrappedNativeTokenDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeTokenDeposit)
				if err := _WrappedNativeToken.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_WrappedNativeToken *WrappedNativeTokenFilterer) ParseDeposit(log types.Log) (*WrappedNativeTokenDeposit, error) {
	event := new(WrappedNativeTokenDeposit)
	if err := _WrappedNativeToken.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedNativeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WrappedNativeToken contract.
type WrappedNativeTokenTransferIterator struct {
	Event *WrappedNativeTokenTransfer // Event containing the contract specifics and raw log

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
func (it *WrappedNativeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeTokenTransfer)
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
		it.Event = new(WrappedNativeTokenTransfer)
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
func (it *WrappedNativeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeTokenTransfer represents a Transfer event raised by the WrappedNativeToken contract.
type WrappedNativeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedNativeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenTransferIterator{contract: _WrappedNativeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WrappedNativeTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeTokenTransfer)
				if err := _WrappedNativeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WrappedNativeToken *WrappedNativeTokenFilterer) ParseTransfer(log types.Log) (*WrappedNativeTokenTransfer, error) {
	event := new(WrappedNativeTokenTransfer)
	if err := _WrappedNativeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedNativeTokenWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WrappedNativeToken contract.
type WrappedNativeTokenWithdrawalIterator struct {
	Event *WrappedNativeTokenWithdrawal // Event containing the contract specifics and raw log

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
func (it *WrappedNativeTokenWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedNativeTokenWithdrawal)
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
		it.Event = new(WrappedNativeTokenWithdrawal)
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
func (it *WrappedNativeTokenWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedNativeTokenWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedNativeTokenWithdrawal represents a Withdrawal event raised by the WrappedNativeToken contract.
type WrappedNativeTokenWithdrawal struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address) (*WrappedNativeTokenWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.FilterLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedNativeTokenWithdrawalIterator{contract: _WrappedNativeToken.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed sender, uint256 amount)
func (_WrappedNativeToken *WrappedNativeTokenFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WrappedNativeTokenWithdrawal, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WrappedNativeToken.contract.WatchLogs(opts, "Withdrawal", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedNativeTokenWithdrawal)
				if err := _WrappedNativeToken.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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
func (_WrappedNativeToken *WrappedNativeTokenFilterer) ParseWithdrawal(log types.Log) (*WrappedNativeTokenWithdrawal, error) {
	event := new(WrappedNativeTokenWithdrawal)
	if err := _WrappedNativeToken.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
