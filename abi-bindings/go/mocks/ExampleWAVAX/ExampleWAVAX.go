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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600c81526020016b0aee4c2e0e0cac84082ac82b60a31b815250604051806040016040528060058152602001640ae82ac82b60db1b815250816003908162000066919062000123565b50600462000075828262000123565b505050620001ef565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620000a957607f821691505b602082108103620000ca57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200011e57600081815260208120601f850160051c81016020861015620000f95750805b601f850160051c820191505b818110156200011a5782815560010162000105565b5050505b505050565b81516001600160401b038111156200013f576200013f6200007e565b620001578162000150845462000094565b84620000d0565b602080601f8311600181146200018f5760008415620001765750858301515b600019600386901b1c1916600185901b1785556200011a565b600085815260208120601f198616915b82811015620001c0578886015182559484019460019091019084016200019f565b5085821015620001df5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610cde80620001ff6000396000f3fe6080604052600436106100c65760003560e01c8063395093511161007f578063a457c2d711610059578063a457c2d71461021e578063a9059cbb1461023e578063d0e30db0146100d5578063dd62ed3e1461025e576100d5565b806339509351146101b357806370a08231146101d357806395d89b4114610209576100d5565b806306fdde03146100dd578063095ea7b31461010857806318160ddd1461013857806323b872dd146101575780632e1a7d4d14610177578063313ce56714610197576100d5565b366100d5576100d361027e565b005b6100d361027e565b3480156100e957600080fd5b506100f26102bf565b6040516100ff9190610b0f565b60405180910390f35b34801561011457600080fd5b50610128610123366004610b79565b610351565b60405190151581526020016100ff565b34801561014457600080fd5b506002545b6040519081526020016100ff565b34801561016357600080fd5b50610128610172366004610ba3565b61036b565b34801561018357600080fd5b506100d3610192366004610bdf565b61038f565b3480156101a357600080fd5b50604051601281526020016100ff565b3480156101bf57600080fd5b506101286101ce366004610b79565b6103db565b3480156101df57600080fd5b506101496101ee366004610bf8565b6001600160a01b031660009081526020819052604090205490565b34801561021557600080fd5b506100f26103fd565b34801561022a57600080fd5b50610128610239366004610b79565b61040c565b34801561024a57600080fd5b50610128610259366004610b79565b61048c565b34801561026a57600080fd5b50610149610279366004610c1a565b61049a565b61028833346104c5565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b6060600380546102ce90610c4d565b80601f01602080910402602001604051908101604052809291908181526020018280546102fa90610c4d565b80156103475780601f1061031c57610100808354040283529160200191610347565b820191906000526020600020905b81548152906001019060200180831161032a57829003601f168201915b5050505050905090565b60003361035f818585610584565b60019150505b92915050565b6000336103798582856106a9565b610384858585610723565b506001949350505050565b61039933826108c7565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a26103d833826109f6565b50565b60003361035f8185856103ee838361049a565b6103f89190610c87565b610584565b6060600480546102ce90610c4d565b6000338161041a828661049a565b90508381101561047f5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6103848286868403610584565b60003361035f818585610723565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b03821661051b5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610476565b806002600082825461052d9190610c87565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b0383166105e65760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610476565b6001600160a01b0382166106475760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610476565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b60006106b5848461049a565b9050600019811461071d57818110156107105760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610476565b61071d8484848403610584565b50505050565b6001600160a01b0383166107875760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610476565b6001600160a01b0382166107e95760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610476565b6001600160a01b038316600090815260208190526040902054818110156108615760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610476565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a361071d565b6001600160a01b0382166109275760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610476565b6001600160a01b0382166000908152602081905260409020548181101561099b5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610476565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910161069c565b505050565b80471015610a465760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610476565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114610a93576040519150601f19603f3d011682016040523d82523d6000602084013e610a98565b606091505b50509050806109f15760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610476565b600060208083528351808285015260005b81811015610b3c57858101830151858201604001528201610b20565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610b7457600080fd5b919050565b60008060408385031215610b8c57600080fd5b610b9583610b5d565b946020939093013593505050565b600080600060608486031215610bb857600080fd5b610bc184610b5d565b9250610bcf60208501610b5d565b9150604084013590509250925092565b600060208284031215610bf157600080fd5b5035919050565b600060208284031215610c0a57600080fd5b610c1382610b5d565b9392505050565b60008060408385031215610c2d57600080fd5b610c3683610b5d565b9150610c4460208401610b5d565b90509250929050565b600181811c90821680610c6157607f821691505b602082108103610c8157634e487b7160e01b600052602260045260246000fd5b50919050565b8082018082111561036557634e487b7160e01b600052601160045260246000fdfea26469706673582212205cb6d838da656ef49f3c86d1511900e2cbc52ba1af30f2c50deec404f349e44f64736f6c63430008120033",
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
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExampleWAVAX.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ExampleWAVAX.Contract.Allowance(&_ExampleWAVAX.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ExampleWAVAX.Contract.Allowance(&_ExampleWAVAX.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExampleWAVAX.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExampleWAVAX.Contract.BalanceOf(&_ExampleWAVAX.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExampleWAVAX *ExampleWAVAXCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExampleWAVAX.Contract.BalanceOf(&_ExampleWAVAX.CallOpts, account)
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

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.DecreaseAllowance(&_ExampleWAVAX.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.DecreaseAllowance(&_ExampleWAVAX.TransactOpts, spender, subtractedValue)
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

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.IncreaseAllowance(&_ExampleWAVAX.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ExampleWAVAX *ExampleWAVAXTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ExampleWAVAX.Contract.IncreaseAllowance(&_ExampleWAVAX.TransactOpts, spender, addedValue)
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
