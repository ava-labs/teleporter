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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenDecimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162000f1e38038062000f1e833981016040819052620000349162000188565b6040518060400160405280600a81526020016926b7b1b5902a37b5b2b760b11b81525060405180604001604052806004815260200163045584d560e41b815250816003908162000085919062000258565b50600462000094828262000258565b50505060ff8116608052620000b6336b204fce5e3e25026110000000620000bd565b506200034c565b6001600160a01b038216620001185760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b80600260008282546200012c919062000324565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b505050565b6000602082840312156200019b57600080fd5b815160ff81168114620001ad57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620001df57607f821691505b6020821081036200020057634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200018357600081815260208120601f850160051c810160208610156200022f5750805b601f850160051c820191505b8181101562000250578281556001016200023b565b505050505050565b81516001600160401b03811115620002745762000274620001b4565b6200028c81620002858454620001ca565b8462000206565b602080601f831160018114620002c45760008415620002ab5750858301515b600019600386901b1c1916600185901b17855562000250565b600085815260208120601f198616915b82811015620002f557888601518255948401946001909101908401620002d4565b5085821015620003145787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b808201808211156200034657634e487b7160e01b600052601160045260246000fd5b92915050565b608051610baf6200036f6000396000818161016201526101ac0152610baf6000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806342966c6811610097578063a0712d6811610066578063a0712d6814610227578063a457c2d71461023a578063a9059cbb1461024d578063dd62ed3e1461026057600080fd5b806342966c68146101ce57806370a08231146101e357806379cc67901461020c57806395d89b411461021f57600080fd5b806323b872dd116100d357806323b872dd1461014d578063313ce5671461016057806339509351146101945780633b97e856146101a757600080fd5b806306fdde03146100fa578063095ea7b31461011857806318160ddd1461013b575b600080fd5b610102610273565b60405161010f91906109e0565b60405180910390f35b61012b610126366004610a4a565b610305565b604051901515815260200161010f565b6002545b60405190815260200161010f565b61012b61015b366004610a74565b61031f565b7f00000000000000000000000000000000000000000000000000000000000000005b60405160ff909116815260200161010f565b61012b6101a2366004610a4a565b610343565b6101827f000000000000000000000000000000000000000000000000000000000000000081565b6101e16101dc366004610ab0565b610365565b005b61013f6101f1366004610ac9565b6001600160a01b031660009081526020819052604090205490565b6101e161021a366004610a4a565b610372565b61010261038b565b6101e1610235366004610ab0565b61039a565b61012b610248366004610a4a565b610400565b61012b61025b366004610a4a565b61047b565b61013f61026e366004610aeb565b610489565b60606003805461028290610b1e565b80601f01602080910402602001604051908101604052809291908181526020018280546102ae90610b1e565b80156102fb5780601f106102d0576101008083540402835291602001916102fb565b820191906000526020600020905b8154815290600101906020018083116102de57829003601f168201915b5050505050905090565b6000336103138185856104b4565b60019150505b92915050565b60003361032d8582856105d9565b610338858585610653565b506001949350505050565b6000336103138185856103568383610489565b6103609190610b58565b6104b4565b61036f33826107f7565b50565b61037d8233836105d9565b61038782826107f7565b5050565b60606004805461028290610b1e565b662386f26fc100008111156103f65760405162461bcd60e51b815260206004820152601f60248201527f4578616d706c6545524332303a206d6178206d696e742065786365656465640060448201526064015b60405180910390fd5b61036f3382610921565b6000338161040e8286610489565b90508381101561046e5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016103ed565b61033882868684036104b4565b600033610313818585610653565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166105165760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016103ed565b6001600160a01b0382166105775760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103ed565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b60006105e58484610489565b9050600019811461064d57818110156106405760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103ed565b61064d84848484036104b4565b50505050565b6001600160a01b0383166106b75760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103ed565b6001600160a01b0382166107195760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103ed565b6001600160a01b038316600090815260208190526040902054818110156107915760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016103ed565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a361064d565b6001600160a01b0382166108575760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b60648201526084016103ed565b6001600160a01b038216600090815260208190526040902054818110156108cb5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b60648201526084016103ed565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91016105cc565b6001600160a01b0382166109775760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016103ed565b80600260008282546109899190610b58565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b600060208083528351808285015260005b81811015610a0d578581018301518582016040015282016109f1565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610a4557600080fd5b919050565b60008060408385031215610a5d57600080fd5b610a6683610a2e565b946020939093013593505050565b600080600060608486031215610a8957600080fd5b610a9284610a2e565b9250610aa060208501610a2e565b9150604084013590509250925092565b600060208284031215610ac257600080fd5b5035919050565b600060208284031215610adb57600080fd5b610ae482610a2e565b9392505050565b60008060408385031215610afe57600080fd5b610b0783610a2e565b9150610b1560208401610a2e565b90509250929050565b600181811c90821680610b3257607f821691505b602082108103610b5257634e487b7160e01b600052602260045260246000fd5b50919050565b8082018082111561031957634e487b7160e01b600052601160045260246000fdfea264697066735822122078e349a0f04205118928beb19778406651cc1f615550248019e2796689cc6c9664736f6c63430008120033",
}

// ExampleERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleERC20MetaData.ABI instead.
var ExampleERC20ABI = ExampleERC20MetaData.ABI

// ExampleERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleERC20MetaData.Bin instead.
var ExampleERC20Bin = ExampleERC20MetaData.Bin

// DeployExampleERC20 deploys a new Ethereum contract, binding an instance of ExampleERC20 to it.
func DeployExampleERC20(auth *bind.TransactOpts, backend bind.ContractBackend, tokenDecimals_ uint8) (common.Address, *types.Transaction, *ExampleERC20, error) {
	parsed, err := ExampleERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleERC20Bin), backend, tokenDecimals_)
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

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_ExampleERC20 *ExampleERC20Caller) TokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "tokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_ExampleERC20 *ExampleERC20Session) TokenDecimals() (uint8, error) {
	return _ExampleERC20.Contract.TokenDecimals(&_ExampleERC20.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint8)
func (_ExampleERC20 *ExampleERC20CallerSession) TokenDecimals() (uint8, error) {
	return _ExampleERC20.Contract.TokenDecimals(&_ExampleERC20.CallOpts)
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
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Approve(&_ExampleERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Approve(&_ExampleERC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Burn(&_ExampleERC20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Burn(&_ExampleERC20.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20Transactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20Session) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.BurnFrom(&_ExampleERC20.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20TransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.BurnFrom(&_ExampleERC20.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ExampleERC20 *ExampleERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ExampleERC20 *ExampleERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.DecreaseAllowance(&_ExampleERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ExampleERC20 *ExampleERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.DecreaseAllowance(&_ExampleERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ExampleERC20 *ExampleERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ExampleERC20 *ExampleERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.IncreaseAllowance(&_ExampleERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ExampleERC20 *ExampleERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.IncreaseAllowance(&_ExampleERC20.TransactOpts, spender, addedValue)
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
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Transfer(&_ExampleERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Transfer(&_ExampleERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.TransferFrom(&_ExampleERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.TransferFrom(&_ExampleERC20.TransactOpts, from, to, amount)
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
