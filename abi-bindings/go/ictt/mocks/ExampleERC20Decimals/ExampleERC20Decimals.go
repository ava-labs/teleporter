// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exampleerc20decimals

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

// ExampleERC20DecimalsMetaData contains all meta data concerning the ExampleERC20Decimals contract.
var ExampleERC20DecimalsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenDecimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051610cba380380610cba83398101604081905261002e91610217565b6040518060400160405280600a81526020016926b7b1b5902a37b5b2b760b11b81525060405180604001604052806004815260200163045584d560e41b815250816003908161007d91906102d6565b50600461008a82826102d6565b5050506100a9336b204fce5e3e250261100000006100b460201b60201c565b60ff166080526103ba565b6001600160a01b0382166100e25760405163ec442f0560e01b81525f60048201526024015b60405180910390fd5b6100ed5f83836100f1565b5050565b6001600160a01b03831661011b578060025f8282546101109190610395565b9091555061018b9050565b6001600160a01b0383165f908152602081905260409020548181101561016d5760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016100d9565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b0382166101a7576002805482900390556101c5565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161020a91815260200190565b60405180910390a3505050565b5f60208284031215610227575f80fd5b815160ff81168114610237575f80fd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061026657607f821691505b60208210810361028457634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156102d157805f5260205f20601f840160051c810160208510156102af5750805b601f840160051c820191505b818110156102ce575f81556001016102bb565b50505b505050565b81516001600160401b038111156102ef576102ef61023e565b610303816102fd8454610252565b8461028a565b602080601f831160018114610336575f841561031f5750858301515b5f19600386901b1c1916600185901b17855561038d565b5f85815260208120601f198616915b8281101561036457888601518255948401946001909101908401610345565b508582101561038157878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b808201808211156103b457634e487b7160e01b5f52601160045260245ffd5b92915050565b6080516108e16103d95f395f8181610133015261016a01526108e15ff3fe608060405234801561000f575f80fd5b50600436106100c7575f3560e01c806342966c681161007957806342966c68146101a157806370a08231146101b457806379cc6790146101dc57806395d89b41146101ef578063a0712d68146101f7578063a9059cbb1461020a578063dd62ed3e1461021d575f80fd5b806306fdde03146100cb578063095ea7b3146100e957806318160ddd1461010c57806323b872dd1461011e578063313ce567146101315780633b97e8561461016557806340c10f191461018c575b5f80fd5b6100d3610230565b6040516100e091906106e1565b60405180910390f35b6100fc6100f7366004610748565b6102c0565b60405190151581526020016100e0565b6002545b6040519081526020016100e0565b6100fc61012c366004610770565b6102d9565b7f00000000000000000000000000000000000000000000000000000000000000005b60405160ff90911681526020016100e0565b6101537f000000000000000000000000000000000000000000000000000000000000000081565b61019f61019a366004610748565b6102fc565b005b61019f6101af3660046107a9565b61033b565b6101106101c23660046107c0565b6001600160a01b03165f9081526020819052604090205490565b61019f6101ea366004610748565b610348565b6100d361035d565b61019f6102053660046107a9565b61036c565b6100fc610218366004610748565b61039e565b61011061022b3660046107e0565b6103ab565b60606003805461023f90610811565b80601f016020809104026020016040519081016040528092919081815260200182805461026b90610811565b80156102b65780601f1061028d576101008083540402835291602001916102b6565b820191905f5260205f20905b81548152906001019060200180831161029957829003601f168201915b5050505050905090565b5f336102cd8185856103d5565b60019150505b92915050565b5f336102e68582856103e7565b6102f1858585610437565b506001949350505050565b678ac7230489e8000081111561032d5760405162461bcd60e51b815260040161032490610849565b60405180910390fd5b6103378282610494565b5050565b61034533826104c8565b50565b6103538233836103e7565b61033782826104c8565b60606004805461023f90610811565b678ac7230489e800008111156103945760405162461bcd60e51b815260040161032490610849565b6103453382610494565b5f336102cd818585610437565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6103e283838360016104fc565b505050565b5f6103f284846103ab565b90505f198114610431578181101561042357828183604051637dc7a0d960e11b815260040161032493929190610880565b61043184848484035f6104fc565b50505050565b6001600160a01b038316610460575f604051634b637e8f60e11b815260040161032491906108a1565b6001600160a01b038216610489575f60405163ec442f0560e01b815260040161032491906108a1565b6103e28383836105ce565b6001600160a01b0382166104bd575f60405163ec442f0560e01b815260040161032491906108a1565b6103375f83836105ce565b6001600160a01b0382166104f1575f604051634b637e8f60e11b815260040161032491906108a1565b610337825f836105ce565b6001600160a01b038416610525575f60405163e602df0560e01b815260040161032491906108a1565b6001600160a01b03831661054e575f604051634a1406b160e11b815260040161032491906108a1565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561043157826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105c091815260200190565b60405180910390a350505050565b6001600160a01b0383166105f8578060025f8282546105ed91906108b5565b909155506106559050565b6001600160a01b0383165f90815260208190526040902054818110156106375783818360405163391434e360e21b815260040161032493929190610880565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b0382166106715760028054829003905561068f565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516106d491815260200190565b60405180910390a3505050565b5f602080835283518060208501525f5b8181101561070d578581018301518582016040015282016106f1565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610743575f80fd5b919050565b5f8060408385031215610759575f80fd5b6107628361072d565b946020939093013593505050565b5f805f60608486031215610782575f80fd5b61078b8461072d565b92506107996020850161072d565b9150604084013590509250925092565b5f602082840312156107b9575f80fd5b5035919050565b5f602082840312156107d0575f80fd5b6107d98261072d565b9392505050565b5f80604083850312156107f1575f80fd5b6107fa8361072d565b91506108086020840161072d565b90509250929050565b600181811c9082168061082557607f821691505b60208210810361084357634e487b7160e01b5f52602260045260245ffd5b50919050565b6020808252601f908201527f4578616d706c6545524332303a206d6178206d696e7420657863656564656400604082015260600190565b6001600160a01b039390931683526020830191909152604082015260600190565b6001600160a01b0391909116815260200190565b808201808211156102d357634e487b7160e01b5f52601160045260245ffdfea164736f6c6343000819000a",
}

// ExampleERC20DecimalsABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleERC20DecimalsMetaData.ABI instead.
var ExampleERC20DecimalsABI = ExampleERC20DecimalsMetaData.ABI

// ExampleERC20DecimalsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleERC20DecimalsMetaData.Bin instead.
var ExampleERC20DecimalsBin = ExampleERC20DecimalsMetaData.Bin

// DeployExampleERC20Decimals deploys a new Ethereum contract, binding an instance of ExampleERC20Decimals to it.
func DeployExampleERC20Decimals(auth *bind.TransactOpts, backend bind.ContractBackend, tokenDecimals_ uint8) (common.Address, *types.Transaction, *ExampleERC20Decimals, error) {
	parsed, err := ExampleERC20DecimalsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleERC20DecimalsBin), backend, tokenDecimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExampleERC20Decimals{ExampleERC20DecimalsCaller: ExampleERC20DecimalsCaller{contract: contract}, ExampleERC20DecimalsTransactor: ExampleERC20DecimalsTransactor{contract: contract}, ExampleERC20DecimalsFilterer: ExampleERC20DecimalsFilterer{contract: contract}}, nil
}

// ExampleERC20Decimals is an auto generated Go binding around an Ethereum contract.
type ExampleERC20Decimals struct {
	ExampleERC20DecimalsCaller     // Read-only binding to the contract
	ExampleERC20DecimalsTransactor // Write-only binding to the contract
	ExampleERC20DecimalsFilterer   // Log filterer for contract events
}

// ExampleERC20DecimalsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleERC20DecimalsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleERC20DecimalsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleERC20DecimalsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleERC20DecimalsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleERC20DecimalsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleERC20DecimalsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleERC20DecimalsSession struct {
	Contract     *ExampleERC20Decimals // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ExampleERC20DecimalsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleERC20DecimalsCallerSession struct {
	Contract *ExampleERC20DecimalsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ExampleERC20DecimalsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleERC20DecimalsTransactorSession struct {
	Contract     *ExampleERC20DecimalsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ExampleERC20DecimalsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleERC20DecimalsRaw struct {
	Contract *ExampleERC20Decimals // Generic contract binding to access the raw methods on
}

// ExampleERC20DecimalsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleERC20DecimalsCallerRaw struct {
	Contract *ExampleERC20DecimalsCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleERC20DecimalsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleERC20DecimalsTransactorRaw struct {
	Contract *ExampleERC20DecimalsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleERC20Decimals creates a new instance of ExampleERC20Decimals, bound to a specific deployed contract.
func NewExampleERC20Decimals(address common.Address, backend bind.ContractBackend) (*ExampleERC20Decimals, error) {
	contract, err := bindExampleERC20Decimals(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20Decimals{ExampleERC20DecimalsCaller: ExampleERC20DecimalsCaller{contract: contract}, ExampleERC20DecimalsTransactor: ExampleERC20DecimalsTransactor{contract: contract}, ExampleERC20DecimalsFilterer: ExampleERC20DecimalsFilterer{contract: contract}}, nil
}

// NewExampleERC20DecimalsCaller creates a new read-only instance of ExampleERC20Decimals, bound to a specific deployed contract.
func NewExampleERC20DecimalsCaller(address common.Address, caller bind.ContractCaller) (*ExampleERC20DecimalsCaller, error) {
	contract, err := bindExampleERC20Decimals(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20DecimalsCaller{contract: contract}, nil
}

// NewExampleERC20DecimalsTransactor creates a new write-only instance of ExampleERC20Decimals, bound to a specific deployed contract.
func NewExampleERC20DecimalsTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleERC20DecimalsTransactor, error) {
	contract, err := bindExampleERC20Decimals(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20DecimalsTransactor{contract: contract}, nil
}

// NewExampleERC20DecimalsFilterer creates a new log filterer instance of ExampleERC20Decimals, bound to a specific deployed contract.
func NewExampleERC20DecimalsFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleERC20DecimalsFilterer, error) {
	contract, err := bindExampleERC20Decimals(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20DecimalsFilterer{contract: contract}, nil
}

// bindExampleERC20Decimals binds a generic wrapper to an already deployed contract.
func bindExampleERC20Decimals(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExampleERC20DecimalsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleERC20Decimals *ExampleERC20DecimalsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleERC20Decimals.Contract.ExampleERC20DecimalsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleERC20Decimals *ExampleERC20DecimalsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.ExampleERC20DecimalsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleERC20Decimals *ExampleERC20DecimalsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.ExampleERC20DecimalsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleERC20Decimals *ExampleERC20DecimalsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleERC20Decimals.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExampleERC20Decimals.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ExampleERC20Decimals.Contract.Allowance(&_ExampleERC20Decimals.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ExampleERC20Decimals.Contract.Allowance(&_ExampleERC20Decimals.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExampleERC20Decimals.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExampleERC20Decimals.Contract.BalanceOf(&_ExampleERC20Decimals.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExampleERC20Decimals.Contract.BalanceOf(&_ExampleERC20Decimals.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ExampleERC20Decimals.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) Decimals() (uint8, error) {
	return _ExampleERC20Decimals.Contract.Decimals(&_ExampleERC20Decimals.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCallerSession) Decimals() (uint8, error) {
	return _ExampleERC20Decimals.Contract.Decimals(&_ExampleERC20Decimals.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExampleERC20Decimals.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) Name() (string, error) {
	return _ExampleERC20Decimals.Contract.Name(&_ExampleERC20Decimals.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCallerSession) Name() (string, error) {
	return _ExampleERC20Decimals.Contract.Name(&_ExampleERC20Decimals.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExampleERC20Decimals.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) Symbol() (string, error) {
	return _ExampleERC20Decimals.Contract.Symbol(&_ExampleERC20Decimals.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCallerSession) Symbol() (string, error) {
	return _ExampleERC20Decimals.Contract.Symbol(&_ExampleERC20Decimals.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCaller) TokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ExampleERC20Decimals.contract.Call(opts, &out, "tokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) TokenDecimals() (uint8, error) {
	return _ExampleERC20Decimals.Contract.TokenDecimals(&_ExampleERC20Decimals.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCallerSession) TokenDecimals() (uint8, error) {
	return _ExampleERC20Decimals.Contract.TokenDecimals(&_ExampleERC20Decimals.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExampleERC20Decimals.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) TotalSupply() (*big.Int, error) {
	return _ExampleERC20Decimals.Contract.TotalSupply(&_ExampleERC20Decimals.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExampleERC20Decimals *ExampleERC20DecimalsCallerSession) TotalSupply() (*big.Int, error) {
	return _ExampleERC20Decimals.Contract.TotalSupply(&_ExampleERC20Decimals.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Approve(&_ExampleERC20Decimals.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Approve(&_ExampleERC20Decimals.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Burn(&_ExampleERC20Decimals.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Burn(&_ExampleERC20Decimals.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.BurnFrom(&_ExampleERC20Decimals.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.BurnFrom(&_ExampleERC20Decimals.TransactOpts, account, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Mint(&_ExampleERC20Decimals.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Mint(&_ExampleERC20Decimals.TransactOpts, account, amount)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) Mint0(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "mint0", amount)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) Mint0(amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Mint0(&_ExampleERC20Decimals.TransactOpts, amount)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) Mint0(amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Mint0(&_ExampleERC20Decimals.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Transfer(&_ExampleERC20Decimals.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Transfer(&_ExampleERC20Decimals.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.TransferFrom(&_ExampleERC20Decimals.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.TransferFrom(&_ExampleERC20Decimals.TransactOpts, from, to, value)
}

// ExampleERC20DecimalsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ExampleERC20Decimals contract.
type ExampleERC20DecimalsApprovalIterator struct {
	Event *ExampleERC20DecimalsApproval // Event containing the contract specifics and raw log

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
func (it *ExampleERC20DecimalsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleERC20DecimalsApproval)
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
		it.Event = new(ExampleERC20DecimalsApproval)
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
func (it *ExampleERC20DecimalsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleERC20DecimalsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleERC20DecimalsApproval represents a Approval event raised by the ExampleERC20Decimals contract.
type ExampleERC20DecimalsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ExampleERC20Decimals *ExampleERC20DecimalsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ExampleERC20DecimalsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ExampleERC20Decimals.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20DecimalsApprovalIterator{contract: _ExampleERC20Decimals.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ExampleERC20Decimals *ExampleERC20DecimalsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ExampleERC20DecimalsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ExampleERC20Decimals.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleERC20DecimalsApproval)
				if err := _ExampleERC20Decimals.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ExampleERC20Decimals *ExampleERC20DecimalsFilterer) ParseApproval(log types.Log) (*ExampleERC20DecimalsApproval, error) {
	event := new(ExampleERC20DecimalsApproval)
	if err := _ExampleERC20Decimals.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleERC20DecimalsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ExampleERC20Decimals contract.
type ExampleERC20DecimalsTransferIterator struct {
	Event *ExampleERC20DecimalsTransfer // Event containing the contract specifics and raw log

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
func (it *ExampleERC20DecimalsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleERC20DecimalsTransfer)
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
		it.Event = new(ExampleERC20DecimalsTransfer)
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
func (it *ExampleERC20DecimalsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleERC20DecimalsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleERC20DecimalsTransfer represents a Transfer event raised by the ExampleERC20Decimals contract.
type ExampleERC20DecimalsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ExampleERC20Decimals *ExampleERC20DecimalsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExampleERC20DecimalsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExampleERC20Decimals.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20DecimalsTransferIterator{contract: _ExampleERC20Decimals.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ExampleERC20Decimals *ExampleERC20DecimalsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ExampleERC20DecimalsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExampleERC20Decimals.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleERC20DecimalsTransfer)
				if err := _ExampleERC20Decimals.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ExampleERC20Decimals *ExampleERC20DecimalsFilterer) ParseTransfer(log types.Log) (*ExampleERC20DecimalsTransfer, error) {
	event := new(ExampleERC20DecimalsTransfer)
	if err := _ExampleERC20Decimals.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
