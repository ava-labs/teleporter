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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenDecimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162000f2238038062000f2283398101604081905262000034916200018c565b6040518060400160405280600a81526020016926b7b1b5902a37b5b2b760b11b81525060405180604001604052806004815260200163045584d560e41b81525081600390816200008591906200025c565b5060046200009482826200025c565b505050620000b5336b204fce5e3e25026110000000620000c160201b60201c565b60ff1660805262000350565b6001600160a01b0382166200011c5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b806002600082825462000130919062000328565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b505050565b6000602082840312156200019f57600080fd5b815160ff81168114620001b157600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620001e357607f821691505b6020821081036200020457634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200018757600081815260208120601f850160051c81016020861015620002335750805b601f850160051c820191505b8181101562000254578281556001016200023f565b505050505050565b81516001600160401b03811115620002785762000278620001b8565b6200029081620002898454620001ce565b846200020a565b602080601f831160018114620002c85760008415620002af5750858301515b600019600386901b1c1916600185901b17855562000254565b600085815260208120601f198616915b82811015620002f957888601518255948401946001909101908401620002d8565b5085821015620003185787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b808201808211156200034a57634e487b7160e01b600052601160045260246000fd5b92915050565b608051610baf620003736000396000818161016201526101ac0152610baf6000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806342966c6811610097578063a0712d6811610066578063a0712d6814610227578063a457c2d71461023a578063a9059cbb1461024d578063dd62ed3e1461026057600080fd5b806342966c68146101ce57806370a08231146101e357806379cc67901461020c57806395d89b411461021f57600080fd5b806323b872dd116100d357806323b872dd1461014d578063313ce5671461016057806339509351146101945780633b97e856146101a757600080fd5b806306fdde03146100fa578063095ea7b31461011857806318160ddd1461013b575b600080fd5b610102610273565b60405161010f91906109e0565b60405180910390f35b61012b610126366004610a4a565b610305565b604051901515815260200161010f565b6002545b60405190815260200161010f565b61012b61015b366004610a74565b61031f565b7f00000000000000000000000000000000000000000000000000000000000000005b60405160ff909116815260200161010f565b61012b6101a2366004610a4a565b610343565b6101827f000000000000000000000000000000000000000000000000000000000000000081565b6101e16101dc366004610ab0565b610365565b005b61013f6101f1366004610ac9565b6001600160a01b031660009081526020819052604090205490565b6101e161021a366004610a4a565b610372565b61010261038b565b6101e1610235366004610ab0565b61039a565b61012b610248366004610a4a565b610400565b61012b61025b366004610a4a565b61047b565b61013f61026e366004610aeb565b610489565b60606003805461028290610b1e565b80601f01602080910402602001604051908101604052809291908181526020018280546102ae90610b1e565b80156102fb5780601f106102d0576101008083540402835291602001916102fb565b820191906000526020600020905b8154815290600101906020018083116102de57829003601f168201915b5050505050905090565b6000336103138185856104b4565b60019150505b92915050565b60003361032d8582856105d9565b610338858585610653565b506001949350505050565b6000336103138185856103568383610489565b6103609190610b58565b6104b4565b61036f33826107f7565b50565b61037d8233836105d9565b61038782826107f7565b5050565b60606004805461028290610b1e565b662386f26fc100008111156103f65760405162461bcd60e51b815260206004820152601f60248201527f4578616d706c6545524332303a206d6178206d696e742065786365656465640060448201526064015b60405180910390fd5b61036f3382610921565b6000338161040e8286610489565b90508381101561046e5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016103ed565b61033882868684036104b4565b600033610313818585610653565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166105165760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016103ed565b6001600160a01b0382166105775760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103ed565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b60006105e58484610489565b9050600019811461064d57818110156106405760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103ed565b61064d84848484036104b4565b50505050565b6001600160a01b0383166106b75760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103ed565b6001600160a01b0382166107195760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103ed565b6001600160a01b038316600090815260208190526040902054818110156107915760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016103ed565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a361064d565b6001600160a01b0382166108575760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b60648201526084016103ed565b6001600160a01b038216600090815260208190526040902054818110156108cb5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b60648201526084016103ed565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91016105cc565b6001600160a01b0382166109775760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016103ed565b80600260008282546109899190610b58565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b600060208083528351808285015260005b81811015610a0d578581018301518582016040015282016109f1565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610a4557600080fd5b919050565b60008060408385031215610a5d57600080fd5b610a6683610a2e565b946020939093013593505050565b600080600060608486031215610a8957600080fd5b610a9284610a2e565b9250610aa060208501610a2e565b9150604084013590509250925092565b600060208284031215610ac257600080fd5b5035919050565b600060208284031215610adb57600080fd5b610ae482610a2e565b9392505050565b60008060408385031215610afe57600080fd5b610b0783610a2e565b9150610b1560208401610a2e565b90509250929050565b600181811c90821680610b3257607f821691505b602082108103610b5257634e487b7160e01b600052602260045260246000fd5b50919050565b8082018082111561031957634e487b7160e01b600052601160045260246000fdfea2646970667358221220a8d229bbe7e5635897f569497697fc7d36f8af64ab416f4df4727bb2479c5cf664736f6c63430008120033",
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
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Approve(&_ExampleERC20Decimals.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Approve(&_ExampleERC20Decimals.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Burn(&_ExampleERC20Decimals.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Burn(&_ExampleERC20Decimals.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.BurnFrom(&_ExampleERC20Decimals.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.BurnFrom(&_ExampleERC20Decimals.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.DecreaseAllowance(&_ExampleERC20Decimals.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.DecreaseAllowance(&_ExampleERC20Decimals.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.IncreaseAllowance(&_ExampleERC20Decimals.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.IncreaseAllowance(&_ExampleERC20Decimals.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Mint(&_ExampleERC20Decimals.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Mint(&_ExampleERC20Decimals.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Transfer(&_ExampleERC20Decimals.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.Transfer(&_ExampleERC20Decimals.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.TransferFrom(&_ExampleERC20Decimals.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ExampleERC20Decimals *ExampleERC20DecimalsTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20Decimals.Contract.TransferFrom(&_ExampleERC20Decimals.TransactOpts, from, to, amount)
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
