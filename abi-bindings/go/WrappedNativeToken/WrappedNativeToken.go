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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200103a3803806200103a8339810160408190526200003491620000dc565b8060405160200162000047919062000194565b60405160208183030381529060405281604051602001620000699190620001c6565b60408051601f19818403018152919052600362000087838262000280565b50600462000096828262000280565b505050506200034c565b634e487b7160e01b600052604160045260246000fd5b60005b83811015620000d3578181015183820152602001620000b9565b50506000910152565b600060208284031215620000ef57600080fd5b81516001600160401b03808211156200010757600080fd5b818401915084601f8301126200011c57600080fd5b815181811115620001315762000131620000a0565b604051601f8201601f19908116603f011681019083821181831017156200015c576200015c620000a0565b816040528281528760208487010111156200017657600080fd5b62000189836020830160208801620000b6565b979650505050505050565b6702bb930b83832b2160c51b815260008251620001b9816008850160208701620000b6565b9190910160080192915050565b605760f81b815260008251620001e4816001850160208701620000b6565b9190910160010192915050565b600181811c908216806200020657607f821691505b6020821081036200022757634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200027b57600081815260208120601f850160051c81016020861015620002565750805b601f850160051c820191505b81811015620002775782815560010162000262565b5050505b505050565b81516001600160401b038111156200029c576200029c620000a0565b620002b481620002ad8454620001f1565b846200022d565b602080601f831160018114620002ec5760008415620002d35750858301515b600019600386901b1c1916600185901b17855562000277565b600085815260208120601f198616915b828110156200031d57888601518255948401946001909101908401620002fc565b50858210156200033c5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610cde806200035c6000396000f3fe6080604052600436106100c65760003560e01c8063395093511161007f578063a457c2d711610059578063a457c2d71461021e578063a9059cbb1461023e578063d0e30db0146100d5578063dd62ed3e1461025e576100d5565b806339509351146101b357806370a08231146101d357806395d89b4114610209576100d5565b806306fdde03146100dd578063095ea7b31461010857806318160ddd1461013857806323b872dd146101575780632e1a7d4d14610177578063313ce56714610197576100d5565b366100d5576100d361027e565b005b6100d361027e565b3480156100e957600080fd5b506100f26102bf565b6040516100ff9190610b0f565b60405180910390f35b34801561011457600080fd5b50610128610123366004610b79565b610351565b60405190151581526020016100ff565b34801561014457600080fd5b506002545b6040519081526020016100ff565b34801561016357600080fd5b50610128610172366004610ba3565b61036b565b34801561018357600080fd5b506100d3610192366004610bdf565b61038f565b3480156101a357600080fd5b50604051601281526020016100ff565b3480156101bf57600080fd5b506101286101ce366004610b79565b6103db565b3480156101df57600080fd5b506101496101ee366004610bf8565b6001600160a01b031660009081526020819052604090205490565b34801561021557600080fd5b506100f26103fd565b34801561022a57600080fd5b50610128610239366004610b79565b61040c565b34801561024a57600080fd5b50610128610259366004610b79565b61048c565b34801561026a57600080fd5b50610149610279366004610c1a565b61049a565b61028833346104c5565b60405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b6060600380546102ce90610c4d565b80601f01602080910402602001604051908101604052809291908181526020018280546102fa90610c4d565b80156103475780601f1061031c57610100808354040283529160200191610347565b820191906000526020600020905b81548152906001019060200180831161032a57829003601f168201915b5050505050905090565b60003361035f818585610584565b60019150505b92915050565b6000336103798582856106a9565b610384858585610723565b506001949350505050565b61039933826108c7565b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a26103d833826109f6565b50565b60003361035f8185856103ee838361049a565b6103f89190610c87565b610584565b6060600480546102ce90610c4d565b6000338161041a828661049a565b90508381101561047f5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6103848286868403610584565b60003361035f818585610723565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b03821661051b5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610476565b806002600082825461052d9190610c87565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b0383166105e65760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610476565b6001600160a01b0382166106475760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610476565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b60006106b5848461049a565b9050600019811461071d57818110156107105760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610476565b61071d8484848403610584565b50505050565b6001600160a01b0383166107875760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610476565b6001600160a01b0382166107e95760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610476565b6001600160a01b038316600090815260208190526040902054818110156108615760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610476565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a361071d565b6001600160a01b0382166109275760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610476565b6001600160a01b0382166000908152602081905260409020548181101561099b5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610476565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910161069c565b505050565b80471015610a465760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610476565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114610a93576040519150601f19603f3d011682016040523d82523d6000602084013e610a98565b606091505b50509050806109f15760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610476565b600060208083528351808285015260005b81811015610b3c57858101830151858201604001528201610b20565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610b7457600080fd5b919050565b60008060408385031215610b8c57600080fd5b610b9583610b5d565b946020939093013593505050565b600080600060608486031215610bb857600080fd5b610bc184610b5d565b9250610bcf60208501610b5d565b9150604084013590509250925092565b600060208284031215610bf157600080fd5b5035919050565b600060208284031215610c0a57600080fd5b610c1382610b5d565b9392505050565b60008060408385031215610c2d57600080fd5b610c3683610b5d565b9150610c4460208401610b5d565b90509250929050565b600181811c90821680610c6157607f821691505b602082108103610c8157634e487b7160e01b600052602260045260246000fd5b50919050565b8082018082111561036557634e487b7160e01b600052601160045260246000fdfea26469706673582212209ad8eb6ac9ed4b9d750722d80d92cc9dfaec6d89ebc58789cf6d8da07d833a2664736f6c63430008120033",
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
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Approve(&_WrappedNativeToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Approve(&_WrappedNativeToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.DecreaseAllowance(&_WrappedNativeToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.DecreaseAllowance(&_WrappedNativeToken.TransactOpts, spender, subtractedValue)
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

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.IncreaseAllowance(&_WrappedNativeToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.IncreaseAllowance(&_WrappedNativeToken.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Transfer(&_WrappedNativeToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.Transfer(&_WrappedNativeToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.TransferFrom(&_WrappedNativeToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WrappedNativeToken *WrappedNativeTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedNativeToken.Contract.TransferFrom(&_WrappedNativeToken.TransactOpts, from, to, amount)
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
