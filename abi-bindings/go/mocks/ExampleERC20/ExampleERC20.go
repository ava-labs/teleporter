// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exampleerc20

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

// ExampleERC20MetaData contains all meta data concerning the ExampleERC20 contract.
var ExampleERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b506040518060400160405280600a81526020016926b7b1b5902a37b5b2b760b11b81525060405180604001604052806004815260200163045584d560e41b815250816003908161005f919061028b565b50600461006c828261028b565b50505061008b336b204fce5e3e2502611000000061009060201b60201c565b61036f565b6001600160a01b0382166100be5760405163ec442f0560e01b81525f60048201526024015b60405180910390fd5b6100c95f83836100cd565b5050565b6001600160a01b0383166100f7578060025f8282546100ec919061034a565b909155506101679050565b6001600160a01b0383165f90815260208190526040902054818110156101495760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016100b5565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b038216610183576002805482900390556101a1565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516101e691815260200190565b60405180910390a3505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061021b57607f821691505b60208210810361023957634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561028657805f5260205f20601f840160051c810160208510156102645750805b601f840160051c820191505b81811015610283575f8155600101610270565b50505b505050565b81516001600160401b038111156102a4576102a46101f3565b6102b8816102b28454610207565b8461023f565b602080601f8311600181146102eb575f84156102d45750858301515b5f19600386901b1c1916600185901b178555610342565b5f85815260208120601f198616915b82811015610319578886015182559484019460019091019084016102fa565b508582101561033657878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b8082018082111561036957634e487b7160e01b5f52601160045260245ffd5b92915050565b6108328061037c5f395ff3fe608060405234801561000f575f80fd5b50600436106100b1575f3560e01c806370a082311161006e57806370a082311461013f57806379cc67901461016757806395d89b411461017a578063a0712d6814610182578063a9059cbb14610195578063dd62ed3e146101a8575f80fd5b806306fdde03146100b5578063095ea7b3146100d357806318160ddd146100f657806323b872dd14610108578063313ce5671461011b57806342966c681461012a575b5f80fd5b6100bd6101e0565b6040516100ca919061069e565b60405180910390f35b6100e66100e1366004610705565b610270565b60405190151581526020016100ca565b6002545b6040519081526020016100ca565b6100e661011636600461072d565b610289565b604051601281526020016100ca565b61013d610138366004610766565b6102ac565b005b6100fa61014d36600461077d565b6001600160a01b03165f9081526020819052604090205490565b61013d610175366004610705565b6102b9565b6100bd6102d2565b61013d610190366004610766565b6102e1565b6100e66101a3366004610705565b610347565b6100fa6101b636600461079d565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6060600380546101ef906107ce565b80601f016020809104026020016040519081016040528092919081815260200182805461021b906107ce565b80156102665780601f1061023d57610100808354040283529160200191610266565b820191905f5260205f20905b81548152906001019060200180831161024957829003601f168201915b5050505050905090565b5f3361027d818585610354565b60019150505b92915050565b5f33610296858285610366565b6102a18585856103e1565b506001949350505050565b6102b6338261043e565b50565b6102c4823383610366565b6102ce828261043e565b5050565b6060600480546101ef906107ce565b662386f26fc1000081111561033d5760405162461bcd60e51b815260206004820152601f60248201527f4578616d706c6545524332303a206d6178206d696e742065786365656465640060448201526064015b60405180910390fd5b6102b63382610472565b5f3361027d8185856103e1565b61036183838360016104a6565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f1981146103db57818110156103cd57604051637dc7a0d960e11b81526001600160a01b03841660048201526024810182905260448101839052606401610334565b6103db84848484035f6104a6565b50505050565b6001600160a01b03831661040a57604051634b637e8f60e11b81525f6004820152602401610334565b6001600160a01b0382166104335760405163ec442f0560e01b81525f6004820152602401610334565b610361838383610578565b6001600160a01b03821661046757604051634b637e8f60e11b81525f6004820152602401610334565b6102ce825f83610578565b6001600160a01b03821661049b5760405163ec442f0560e01b81525f6004820152602401610334565b6102ce5f8383610578565b6001600160a01b0384166104cf5760405163e602df0560e01b81525f6004820152602401610334565b6001600160a01b0383166104f857604051634a1406b160e11b81525f6004820152602401610334565b6001600160a01b038085165f90815260016020908152604080832093871683529290522082905580156103db57826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161056a91815260200190565b60405180910390a350505050565b6001600160a01b0383166105a2578060025f8282546105979190610806565b909155506106129050565b6001600160a01b0383165f90815260208190526040902054818110156105f45760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401610334565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b03821661062e5760028054829003905561064c565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161069191815260200190565b60405180910390a3505050565b5f602080835283518060208501525f5b818110156106ca578581018301518582016040015282016106ae565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610700575f80fd5b919050565b5f8060408385031215610716575f80fd5b61071f836106ea565b946020939093013593505050565b5f805f6060848603121561073f575f80fd5b610748846106ea565b9250610756602085016106ea565b9150604084013590509250925092565b5f60208284031215610776575f80fd5b5035919050565b5f6020828403121561078d575f80fd5b610796826106ea565b9392505050565b5f80604083850312156107ae575f80fd5b6107b7836106ea565b91506107c5602084016106ea565b90509250929050565b600181811c908216806107e257607f821691505b60208210810361080057634e487b7160e01b5f52602260045260245ffd5b50919050565b8082018082111561028357634e487b7160e01b5f52601160045260245ffdfea164736f6c6343000819000a",
}

// ExampleERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleERC20MetaData.ABI instead.
var ExampleERC20ABI = ExampleERC20MetaData.ABI

// ExampleERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleERC20MetaData.Bin instead.
var ExampleERC20Bin = ExampleERC20MetaData.Bin

// DeployExampleERC20 deploys a new Ethereum contract, binding an instance of ExampleERC20 to it.
func DeployExampleERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExampleERC20, error) {
	parsed, err := ExampleERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExampleERC20{ExampleERC20Caller: ExampleERC20Caller{contract: contract}, ExampleERC20Transactor: ExampleERC20Transactor{contract: contract}, ExampleERC20Filterer: ExampleERC20Filterer{contract: contract}}, nil
}

// ExampleERC20 is an auto generated Go binding around an Ethereum contract.
type ExampleERC20 struct {
	ExampleERC20Caller     // Read-only binding to the contract
	ExampleERC20Transactor // Write-only binding to the contract
	ExampleERC20Filterer   // Log filterer for contract events
}

// ExampleERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleERC20Session struct {
	Contract     *ExampleERC20     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExampleERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleERC20CallerSession struct {
	Contract *ExampleERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ExampleERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleERC20TransactorSession struct {
	Contract     *ExampleERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExampleERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleERC20Raw struct {
	Contract *ExampleERC20 // Generic contract binding to access the raw methods on
}

// ExampleERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleERC20CallerRaw struct {
	Contract *ExampleERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ExampleERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleERC20TransactorRaw struct {
	Contract *ExampleERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleERC20 creates a new instance of ExampleERC20, bound to a specific deployed contract.
func NewExampleERC20(address common.Address, backend bind.ContractBackend) (*ExampleERC20, error) {
	contract, err := bindExampleERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20{ExampleERC20Caller: ExampleERC20Caller{contract: contract}, ExampleERC20Transactor: ExampleERC20Transactor{contract: contract}, ExampleERC20Filterer: ExampleERC20Filterer{contract: contract}}, nil
}

// NewExampleERC20Caller creates a new read-only instance of ExampleERC20, bound to a specific deployed contract.
func NewExampleERC20Caller(address common.Address, caller bind.ContractCaller) (*ExampleERC20Caller, error) {
	contract, err := bindExampleERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20Caller{contract: contract}, nil
}

// NewExampleERC20Transactor creates a new write-only instance of ExampleERC20, bound to a specific deployed contract.
func NewExampleERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ExampleERC20Transactor, error) {
	contract, err := bindExampleERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20Transactor{contract: contract}, nil
}

// NewExampleERC20Filterer creates a new log filterer instance of ExampleERC20, bound to a specific deployed contract.
func NewExampleERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ExampleERC20Filterer, error) {
	contract, err := bindExampleERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20Filterer{contract: contract}, nil
}

// bindExampleERC20 binds a generic wrapper to an already deployed contract.
func bindExampleERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExampleERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleERC20 *ExampleERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleERC20.Contract.ExampleERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleERC20 *ExampleERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleERC20.Contract.ExampleERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleERC20 *ExampleERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleERC20.Contract.ExampleERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleERC20 *ExampleERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleERC20 *ExampleERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleERC20 *ExampleERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExampleERC20 *ExampleERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExampleERC20 *ExampleERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ExampleERC20.Contract.Allowance(&_ExampleERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExampleERC20 *ExampleERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ExampleERC20.Contract.Allowance(&_ExampleERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExampleERC20 *ExampleERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExampleERC20 *ExampleERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExampleERC20.Contract.BalanceOf(&_ExampleERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExampleERC20 *ExampleERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExampleERC20.Contract.BalanceOf(&_ExampleERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExampleERC20 *ExampleERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExampleERC20 *ExampleERC20Session) Decimals() (uint8, error) {
	return _ExampleERC20.Contract.Decimals(&_ExampleERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExampleERC20 *ExampleERC20CallerSession) Decimals() (uint8, error) {
	return _ExampleERC20.Contract.Decimals(&_ExampleERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleERC20 *ExampleERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleERC20 *ExampleERC20Session) Name() (string, error) {
	return _ExampleERC20.Contract.Name(&_ExampleERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleERC20 *ExampleERC20CallerSession) Name() (string, error) {
	return _ExampleERC20.Contract.Name(&_ExampleERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleERC20 *ExampleERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleERC20 *ExampleERC20Session) Symbol() (string, error) {
	return _ExampleERC20.Contract.Symbol(&_ExampleERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleERC20 *ExampleERC20CallerSession) Symbol() (string, error) {
	return _ExampleERC20.Contract.Symbol(&_ExampleERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExampleERC20 *ExampleERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExampleERC20 *ExampleERC20Session) TotalSupply() (*big.Int, error) {
	return _ExampleERC20.Contract.TotalSupply(&_ExampleERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExampleERC20 *ExampleERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ExampleERC20.Contract.TotalSupply(&_ExampleERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ExampleERC20 *ExampleERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ExampleERC20 *ExampleERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Approve(&_ExampleERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ExampleERC20 *ExampleERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Approve(&_ExampleERC20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ExampleERC20 *ExampleERC20Transactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ExampleERC20 *ExampleERC20Session) Burn(value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Burn(&_ExampleERC20.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ExampleERC20 *ExampleERC20TransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Burn(&_ExampleERC20.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ExampleERC20 *ExampleERC20Transactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ExampleERC20 *ExampleERC20Session) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.BurnFrom(&_ExampleERC20.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ExampleERC20 *ExampleERC20TransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.BurnFrom(&_ExampleERC20.TransactOpts, account, value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20Transactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20Session) Mint(amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Mint(&_ExampleERC20.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20TransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Mint(&_ExampleERC20.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ExampleERC20 *ExampleERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ExampleERC20 *ExampleERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Transfer(&_ExampleERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ExampleERC20 *ExampleERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Transfer(&_ExampleERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ExampleERC20 *ExampleERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ExampleERC20 *ExampleERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.TransferFrom(&_ExampleERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ExampleERC20 *ExampleERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.TransferFrom(&_ExampleERC20.TransactOpts, from, to, value)
}

// ExampleERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ExampleERC20 contract.
type ExampleERC20ApprovalIterator struct {
	Event *ExampleERC20Approval // Event containing the contract specifics and raw log

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
func (it *ExampleERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleERC20Approval)
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
		it.Event = new(ExampleERC20Approval)
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
func (it *ExampleERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleERC20Approval represents a Approval event raised by the ExampleERC20 contract.
type ExampleERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ExampleERC20 *ExampleERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ExampleERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ExampleERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20ApprovalIterator{contract: _ExampleERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ExampleERC20 *ExampleERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ExampleERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ExampleERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleERC20Approval)
				if err := _ExampleERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ExampleERC20 *ExampleERC20Filterer) ParseApproval(log types.Log) (*ExampleERC20Approval, error) {
	event := new(ExampleERC20Approval)
	if err := _ExampleERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ExampleERC20 contract.
type ExampleERC20TransferIterator struct {
	Event *ExampleERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ExampleERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleERC20Transfer)
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
		it.Event = new(ExampleERC20Transfer)
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
func (it *ExampleERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleERC20Transfer represents a Transfer event raised by the ExampleERC20 contract.
type ExampleERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ExampleERC20 *ExampleERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExampleERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExampleERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20TransferIterator{contract: _ExampleERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ExampleERC20 *ExampleERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ExampleERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExampleERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleERC20Transfer)
				if err := _ExampleERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ExampleERC20 *ExampleERC20Filterer) ParseTransfer(log types.Log) (*ExampleERC20Transfer, error) {
	event := new(ExampleERC20Transfer)
	if err := _ExampleERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
