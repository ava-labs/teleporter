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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenDecimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162000d0d38038062000d0d833981016040819052620000349162000236565b6040518060400160405280600a81526020016926b7b1b5902a37b5b2b760b11b81525060405180604001604052806004815260200163045584d560e41b815250816003908162000085919062000307565b50600462000094828262000307565b505050620000b5336b204fce5e3e25026110000000620000c160201b60201c565b60ff16608052620003fb565b6001600160a01b038216620000f15760405163ec442f0560e01b8152600060048201526024015b60405180910390fd5b620000ff6000838362000103565b5050565b6001600160a01b03831662000132578060026000828254620001269190620003d3565b90915550620001a69050565b6001600160a01b03831660009081526020819052604090205481811015620001875760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401620000e8565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b038216620001c457600280548290039055620001e3565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200022991815260200190565b60405180910390a3505050565b6000602082840312156200024957600080fd5b815160ff811681146200025b57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200028d57607f821691505b602082108103620002ae57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200030257600081815260208120601f850160051c81016020861015620002dd5750805b601f850160051c820191505b81811015620002fe57828155600101620002e9565b5050505b505050565b81516001600160401b0381111562000323576200032362000262565b6200033b8162000334845462000278565b84620002b4565b602080601f8311600181146200037357600084156200035a5750858301515b600019600386901b1c1916600185901b178555620002fe565b600085815260208120601f198616915b82811015620003a45788860151825594840194600190910190840162000383565b5085821015620003c35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b80820180821115620003f557634e487b7160e01b600052601160045260246000fd5b92915050565b6080516108ef6200041e6000396000818161013c015261017301526108ef6000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806342966c681161008c57806395d89b411161006657806395d89b41146101e6578063a0712d68146101ee578063a9059cbb14610201578063dd62ed3e1461021457600080fd5b806342966c681461019557806370a08231146101aa57806379cc6790146101d357600080fd5b806306fdde03146100d4578063095ea7b3146100f257806318160ddd1461011557806323b872dd14610127578063313ce5671461013a5780633b97e8561461016e575b600080fd5b6100dc61024d565b6040516100e99190610720565b60405180910390f35b61010561010036600461078a565b6102df565b60405190151581526020016100e9565b6002545b6040519081526020016100e9565b6101056101353660046107b4565b6102f9565b7f00000000000000000000000000000000000000000000000000000000000000005b60405160ff90911681526020016100e9565b61015c7f000000000000000000000000000000000000000000000000000000000000000081565b6101a86101a33660046107f0565b61031d565b005b6101196101b8366004610809565b6001600160a01b031660009081526020819052604090205490565b6101a86101e136600461078a565b61032a565b6100dc610343565b6101a86101fc3660046107f0565b610352565b61010561020f36600461078a565b6103b8565b61011961022236600461082b565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b60606003805461025c9061085e565b80601f01602080910402602001604051908101604052809291908181526020018280546102889061085e565b80156102d55780601f106102aa576101008083540402835291602001916102d5565b820191906000526020600020905b8154815290600101906020018083116102b857829003601f168201915b5050505050905090565b6000336102ed8185856103c6565b60019150505b92915050565b6000336103078582856103d8565b610312858585610456565b506001949350505050565b61032733826104b5565b50565b6103358233836103d8565b61033f82826104b5565b5050565b60606004805461025c9061085e565b662386f26fc100008111156103ae5760405162461bcd60e51b815260206004820152601f60248201527f4578616d706c6545524332303a206d6178206d696e742065786365656465640060448201526064015b60405180910390fd5b61032733826104eb565b6000336102ed818585610456565b6103d38383836001610521565b505050565b6001600160a01b038381166000908152600160209081526040808320938616835292905220546000198114610450578181101561044157604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016103a5565b61045084848484036000610521565b50505050565b6001600160a01b03831661048057604051634b637e8f60e11b8152600060048201526024016103a5565b6001600160a01b0382166104aa5760405163ec442f0560e01b8152600060048201526024016103a5565b6103d38383836105f6565b6001600160a01b0382166104df57604051634b637e8f60e11b8152600060048201526024016103a5565b61033f826000836105f6565b6001600160a01b0382166105155760405163ec442f0560e01b8152600060048201526024016103a5565b61033f600083836105f6565b6001600160a01b03841661054b5760405163e602df0560e01b8152600060048201526024016103a5565b6001600160a01b03831661057557604051634a1406b160e11b8152600060048201526024016103a5565b6001600160a01b038085166000908152600160209081526040808320938716835292905220829055801561045057826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105e891815260200190565b60405180910390a350505050565b6001600160a01b0383166106215780600260008282546106169190610898565b909155506106939050565b6001600160a01b038316600090815260208190526040902054818110156106745760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016103a5565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b0382166106af576002805482900390556106ce565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161071391815260200190565b60405180910390a3505050565b600060208083528351808285015260005b8181101561074d57858101830151858201604001528201610731565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461078557600080fd5b919050565b6000806040838503121561079d57600080fd5b6107a68361076e565b946020939093013593505050565b6000806000606084860312156107c957600080fd5b6107d28461076e565b92506107e06020850161076e565b9150604084013590509250925092565b60006020828403121561080257600080fd5b5035919050565b60006020828403121561081b57600080fd5b6108248261076e565b9392505050565b6000806040838503121561083e57600080fd5b6108478361076e565b91506108556020840161076e565b90509250929050565b600181811c9082168061087257607f821691505b60208210810361089257634e487b7160e01b600052602260045260246000fd5b50919050565b808201808211156102f357634e487b7160e01b600052601160045260246000fdfea26469706673582212200720f299407e7bb8cc6db75f57f6f5ea73717e28e3ab1edc7a219fb614dfcb6864736f6c63430008140033",
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
