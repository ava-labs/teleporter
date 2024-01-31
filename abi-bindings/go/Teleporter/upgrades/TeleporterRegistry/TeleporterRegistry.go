// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teleporterregistry

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

// ProtocolRegistryEntry is an auto generated low-level Go binding around an user-defined struct.
type ProtocolRegistryEntry struct {
	Version         *big.Int
	ProtocolAddress common.Address
}

// TeleporterRegistryMetaData contains all meta data concerning the TeleporterRegistry contract.
var TeleporterRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"internalType\":\"structProtocolRegistryEntry[]\",\"name\":\"initialEntries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"name\":\"AddProtocolVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newVersion\",\"type\":\"uint256\"}],\"name\":\"LatestVersionUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_VERSION_INCREMENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORS_SOURCE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"addProtocolVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getAddressFromVersion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestTeleporter\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getTeleporterFromVersion\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"name\":\"getVersionFromAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620010963803806200109683398101604081905262000034916200041a565b7302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000087573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000ad91906200050c565b608052805160005b81811015620000fe57620000eb838281518110620000d757620000d762000526565b60200260200101516200010760201b60201c565b620000f68162000552565b9050620000b5565b5050506200058a565b80516000036200015e5760405162461bcd60e51b815260206004820181905260248201527f54656c65706f7274657252656769737472793a207a65726f2076657273696f6e60448201526064015b60405180910390fd5b80516000908152600160205260409020546001600160a01b031615620001da5760405162461bcd60e51b815260206004820152602a60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20616c72656044820152696164792065786973747360b01b606482015260840162000155565b60208101516001600160a01b0316620002485760405162461bcd60e51b815260206004820152602960248201527f54656c65706f7274657252656769737472793a207a65726f2070726f746f636f6044820152686c206164647265737360b81b606482015260840162000155565b600054620002596101f4826200056e565b82511115620002c25760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20696e637260448201526d0cadacadce840e8dede40d0d2ced60931b606482015260840162000155565b602082810180518451600090815260018452604080822080546001600160a01b0319166001600160a01b03948516179055925190911681526002909252902054825111156200032c5781516020808401516001600160a01b03166000908152600290915260409020555b602082015182516040516001600160a01b03909216917fa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a90600090a38151811015620003a2578151600081815560405183917f30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b091a35b5050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715620003e157620003e1620003a6565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620004125762000412620003a6565b604052919050565b600060208083850312156200042e57600080fd5b82516001600160401b03808211156200044657600080fd5b818501915085601f8301126200045b57600080fd5b815181811115620004705762000470620003a6565b62000480848260051b01620003e7565b818152848101925060069190911b830184019087821115620004a157600080fd5b928401925b81841015620005015760408489031215620004c15760008081fd5b620004cb620003bc565b84518152858501516001600160a01b0381168114620004ea5760008081fd5b8187015283526040939093019291840191620004a6565b979650505050505050565b6000602082840312156200051f57600080fd5b5051919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016200056757620005676200053c565b5060010190565b808201808211156200058457620005846200053c565b92915050565b608051610ae9620005ad6000396000818161014901526102640152610ae96000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063ac473ac311610066578063ac473ac314610124578063b771b3bc1461012d578063c07f47d41461013b578063d127dc9b14610144578063d820e64f1461016b57600080fd5b80630731775d146100a3578063215abce9146100c857806341f34ed9146100db57806346f9ef49146100f05780634c1f08ce14610103575b600080fd5b6100ab600081565b6040516001600160a01b0390911681526020015b60405180910390f35b6100ab6100d63660046107c5565b610173565b6100ee6100e93660046107de565b610184565b005b6100ab6100fe3660046107c5565b6103f9565b610116610111366004610823565b6104be565b6040519081526020016100bf565b6101166101f481565b6100ab6005600160991b0181565b61011660005481565b6101167f000000000000000000000000000000000000000000000000000000000000000081565b6100ab610566565b600061017e826103f9565b92915050565b6040516306f8253560e41b815263ffffffff8216600482015260009081906005600160991b0190636f82535090602401600060405180830381865afa1580156101d1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526101f991908101906108c5565b91509150806102605760405162461bcd60e51b815260206004820152602860248201527f54656c65706f7274657252656769737472793a20696e76616c69642077617270604482015267206d65737361676560c01b60648201526084015b60405180910390fd5b81517f0000000000000000000000000000000000000000000000000000000000000000146102e45760405162461bcd60e51b815260206004820152602b60248201527f54656c65706f7274657252656769737472793a20696e76616c696420736f757260448201526a18d94818da185a5b88125160aa1b6064820152608401610257565b60208201516001600160a01b0316156103595760405162461bcd60e51b815260206004820152603160248201527f54656c65706f7274657252656769737472793a20696e76616c6964206f726967604482015270696e2073656e646572206164647265737360781b6064820152608401610257565b600080836040015180602001905181019061037491906109cd565b90925090506001600160a01b03811630146103e95760405162461bcd60e51b815260206004820152602f60248201527f54656c65706f7274657252656769737472793a20696e76616c6964206465737460448201526e696e6174696f6e206164647265737360881b6064820152608401610257565b6103f282610578565b5050505050565b60008160000361044b5760405162461bcd60e51b815260206004820181905260248201527f54656c65706f7274657252656769737472793a207a65726f2076657273696f6e6044820152606401610257565b6000828152600160205260409020546001600160a01b03168061017e5760405162461bcd60e51b815260206004820152602560248201527f54656c65706f7274657252656769737472793a2076657273696f6e206e6f7420604482015264199bdd5b9960da1b6064820152608401610257565b60006001600160a01b0382166104e65760405162461bcd60e51b815260040161025790610a49565b6001600160a01b0382166000908152600260205260408120549081900361017e5760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f7274657252656769737472793a2070726f746f636f6c2061646460448201526d1c995cdcc81b9bdd08199bdd5b9960921b6064820152608401610257565b60006105736000546103f9565b905090565b80516000036105c95760405162461bcd60e51b815260206004820181905260248201527f54656c65706f7274657252656769737472793a207a65726f2076657273696f6e6044820152606401610257565b80516000908152600160205260409020546001600160a01b0316156106435760405162461bcd60e51b815260206004820152602a60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20616c72656044820152696164792065786973747360b01b6064820152608401610257565b60208101516001600160a01b031661066d5760405162461bcd60e51b815260040161025790610a49565b60005461067c6101f482610a92565b825111156106e35760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f7274657252656769737472793a2076657273696f6e20696e637260448201526d0cadacadce840e8dede40d0d2ced60931b6064820152608401610257565b602082810180518451600090815260018452604080822080546001600160a01b0319166001600160a01b039485161790559251909116815260029092529020548251111561074c5781516020808401516001600160a01b03166000908152600290915260409020555b602082015182516040516001600160a01b03909216917fa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a90600090a381518110156107c1578151600081815560405183917f30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b091a35b5050565b6000602082840312156107d757600080fd5b5035919050565b6000602082840312156107f057600080fd5b813563ffffffff8116811461080457600080fd5b9392505050565b6001600160a01b038116811461082057600080fd5b50565b60006020828403121561083557600080fd5b81356108048161080b565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561087957610879610840565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156108a8576108a8610840565b604052919050565b805180151581146108c057600080fd5b919050565b600080604083850312156108d857600080fd5b825167ffffffffffffffff808211156108f057600080fd5b908401906060828703121561090457600080fd5b61090c610856565b8251815260208084015161091f8161080b565b8282015260408401518381111561093557600080fd5b80850194505087601f85011261094a57600080fd5b83518381111561095c5761095c610840565b61096e601f8201601f1916830161087f565b9350808452888282870101111561098457600080fd5b60005b818110156109a2578581018301518582018401528201610987565b506000828286010152508260408301528195506109c08188016108b0565b9450505050509250929050565b60008082840360608112156109e157600080fd5b60408112156109ef57600080fd5b506040516040810181811067ffffffffffffffff82111715610a1357610a13610840565b604052835181526020840151610a288161080b565b60208201526040840151909250610a3e8161080b565b809150509250929050565b60208082526029908201527f54656c65706f7274657252656769737472793a207a65726f2070726f746f636f6040820152686c206164647265737360b81b606082015260800190565b8082018082111561017e57634e487b7160e01b600052601160045260246000fdfea264697066735822122087260e058e11432415f580187bb4f42226bbc96d98b5c3bf8267153d86c4a40864736f6c63430008120033",
}

// TeleporterRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeleporterRegistryMetaData.ABI instead.
var TeleporterRegistryABI = TeleporterRegistryMetaData.ABI

// TeleporterRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TeleporterRegistryMetaData.Bin instead.
var TeleporterRegistryBin = TeleporterRegistryMetaData.Bin

// DeployTeleporterRegistry deploys a new Ethereum contract, binding an instance of TeleporterRegistry to it.
func DeployTeleporterRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, initialEntries []ProtocolRegistryEntry) (common.Address, *types.Transaction, *TeleporterRegistry, error) {
	parsed, err := TeleporterRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TeleporterRegistryBin), backend, initialEntries)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TeleporterRegistry{TeleporterRegistryCaller: TeleporterRegistryCaller{contract: contract}, TeleporterRegistryTransactor: TeleporterRegistryTransactor{contract: contract}, TeleporterRegistryFilterer: TeleporterRegistryFilterer{contract: contract}}, nil
}

// TeleporterRegistry is an auto generated Go binding around an Ethereum contract.
type TeleporterRegistry struct {
	TeleporterRegistryCaller     // Read-only binding to the contract
	TeleporterRegistryTransactor // Write-only binding to the contract
	TeleporterRegistryFilterer   // Log filterer for contract events
}

// TeleporterRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeleporterRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeleporterRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeleporterRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeleporterRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeleporterRegistrySession struct {
	Contract     *TeleporterRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TeleporterRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeleporterRegistryCallerSession struct {
	Contract *TeleporterRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TeleporterRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeleporterRegistryTransactorSession struct {
	Contract     *TeleporterRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TeleporterRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeleporterRegistryRaw struct {
	Contract *TeleporterRegistry // Generic contract binding to access the raw methods on
}

// TeleporterRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeleporterRegistryCallerRaw struct {
	Contract *TeleporterRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TeleporterRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeleporterRegistryTransactorRaw struct {
	Contract *TeleporterRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeleporterRegistry creates a new instance of TeleporterRegistry, bound to a specific deployed contract.
func NewTeleporterRegistry(address common.Address, backend bind.ContractBackend) (*TeleporterRegistry, error) {
	contract, err := bindTeleporterRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistry{TeleporterRegistryCaller: TeleporterRegistryCaller{contract: contract}, TeleporterRegistryTransactor: TeleporterRegistryTransactor{contract: contract}, TeleporterRegistryFilterer: TeleporterRegistryFilterer{contract: contract}}, nil
}

// NewTeleporterRegistryCaller creates a new read-only instance of TeleporterRegistry, bound to a specific deployed contract.
func NewTeleporterRegistryCaller(address common.Address, caller bind.ContractCaller) (*TeleporterRegistryCaller, error) {
	contract, err := bindTeleporterRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryCaller{contract: contract}, nil
}

// NewTeleporterRegistryTransactor creates a new write-only instance of TeleporterRegistry, bound to a specific deployed contract.
func NewTeleporterRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TeleporterRegistryTransactor, error) {
	contract, err := bindTeleporterRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryTransactor{contract: contract}, nil
}

// NewTeleporterRegistryFilterer creates a new log filterer instance of TeleporterRegistry, bound to a specific deployed contract.
func NewTeleporterRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TeleporterRegistryFilterer, error) {
	contract, err := bindTeleporterRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryFilterer{contract: contract}, nil
}

// bindTeleporterRegistry binds a generic wrapper to an already deployed contract.
func bindTeleporterRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeleporterRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterRegistry *TeleporterRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterRegistry.Contract.TeleporterRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterRegistry *TeleporterRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.TeleporterRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterRegistry *TeleporterRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.TeleporterRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeleporterRegistry *TeleporterRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeleporterRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeleporterRegistry *TeleporterRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeleporterRegistry *TeleporterRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.contract.Transact(opts, method, params...)
}

// MAXVERSIONINCREMENT is a free data retrieval call binding the contract method 0xac473ac3.
//
// Solidity: function MAX_VERSION_INCREMENT() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCaller) MAXVERSIONINCREMENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "MAX_VERSION_INCREMENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVERSIONINCREMENT is a free data retrieval call binding the contract method 0xac473ac3.
//
// Solidity: function MAX_VERSION_INCREMENT() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistrySession) MAXVERSIONINCREMENT() (*big.Int, error) {
	return _TeleporterRegistry.Contract.MAXVERSIONINCREMENT(&_TeleporterRegistry.CallOpts)
}

// MAXVERSIONINCREMENT is a free data retrieval call binding the contract method 0xac473ac3.
//
// Solidity: function MAX_VERSION_INCREMENT() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) MAXVERSIONINCREMENT() (*big.Int, error) {
	return _TeleporterRegistry.Contract.MAXVERSIONINCREMENT(&_TeleporterRegistry.CallOpts)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) VALIDATORSSOURCEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "VALIDATORS_SOURCE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _TeleporterRegistry.Contract.VALIDATORSSOURCEADDRESS(&_TeleporterRegistry.CallOpts)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _TeleporterRegistry.Contract.VALIDATORSSOURCEADDRESS(&_TeleporterRegistry.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterRegistry.Contract.WARPMESSENGER(&_TeleporterRegistry.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) WARPMESSENGER() (common.Address, error) {
	return _TeleporterRegistry.Contract.WARPMESSENGER(&_TeleporterRegistry.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterRegistry *TeleporterRegistryCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterRegistry *TeleporterRegistrySession) BlockchainID() ([32]byte, error) {
	return _TeleporterRegistry.Contract.BlockchainID(&_TeleporterRegistry.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) BlockchainID() ([32]byte, error) {
	return _TeleporterRegistry.Contract.BlockchainID(&_TeleporterRegistry.CallOpts)
}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetAddressFromVersion(opts *bind.CallOpts, version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getAddressFromVersion", version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) GetAddressFromVersion(version *big.Int) (common.Address, error) {
	return _TeleporterRegistry.Contract.GetAddressFromVersion(&_TeleporterRegistry.CallOpts, version)
}

// GetAddressFromVersion is a free data retrieval call binding the contract method 0x46f9ef49.
//
// Solidity: function getAddressFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetAddressFromVersion(version *big.Int) (common.Address, error) {
	return _TeleporterRegistry.Contract.GetAddressFromVersion(&_TeleporterRegistry.CallOpts, version)
}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetLatestTeleporter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getLatestTeleporter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) GetLatestTeleporter() (common.Address, error) {
	return _TeleporterRegistry.Contract.GetLatestTeleporter(&_TeleporterRegistry.CallOpts)
}

// GetLatestTeleporter is a free data retrieval call binding the contract method 0xd820e64f.
//
// Solidity: function getLatestTeleporter() view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetLatestTeleporter() (common.Address, error) {
	return _TeleporterRegistry.Contract.GetLatestTeleporter(&_TeleporterRegistry.CallOpts)
}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetTeleporterFromVersion(opts *bind.CallOpts, version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getTeleporterFromVersion", version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistrySession) GetTeleporterFromVersion(version *big.Int) (common.Address, error) {
	return _TeleporterRegistry.Contract.GetTeleporterFromVersion(&_TeleporterRegistry.CallOpts, version)
}

// GetTeleporterFromVersion is a free data retrieval call binding the contract method 0x215abce9.
//
// Solidity: function getTeleporterFromVersion(uint256 version) view returns(address)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetTeleporterFromVersion(version *big.Int) (common.Address, error) {
	return _TeleporterRegistry.Contract.GetTeleporterFromVersion(&_TeleporterRegistry.CallOpts, version)
}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetVersionFromAddress(opts *bind.CallOpts, protocolAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getVersionFromAddress", protocolAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistrySession) GetVersionFromAddress(protocolAddress common.Address) (*big.Int, error) {
	return _TeleporterRegistry.Contract.GetVersionFromAddress(&_TeleporterRegistry.CallOpts, protocolAddress)
}

// GetVersionFromAddress is a free data retrieval call binding the contract method 0x4c1f08ce.
//
// Solidity: function getVersionFromAddress(address protocolAddress) view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetVersionFromAddress(protocolAddress common.Address) (*big.Int, error) {
	return _TeleporterRegistry.Contract.GetVersionFromAddress(&_TeleporterRegistry.CallOpts, protocolAddress)
}

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCaller) LatestVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "latestVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistrySession) LatestVersion() (*big.Int, error) {
	return _TeleporterRegistry.Contract.LatestVersion(&_TeleporterRegistry.CallOpts)
}

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) LatestVersion() (*big.Int, error) {
	return _TeleporterRegistry.Contract.LatestVersion(&_TeleporterRegistry.CallOpts)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_TeleporterRegistry *TeleporterRegistryTransactor) AddProtocolVersion(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _TeleporterRegistry.contract.Transact(opts, "addProtocolVersion", messageIndex)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_TeleporterRegistry *TeleporterRegistrySession) AddProtocolVersion(messageIndex uint32) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.AddProtocolVersion(&_TeleporterRegistry.TransactOpts, messageIndex)
}

// AddProtocolVersion is a paid mutator transaction binding the contract method 0x41f34ed9.
//
// Solidity: function addProtocolVersion(uint32 messageIndex) returns()
func (_TeleporterRegistry *TeleporterRegistryTransactorSession) AddProtocolVersion(messageIndex uint32) (*types.Transaction, error) {
	return _TeleporterRegistry.Contract.AddProtocolVersion(&_TeleporterRegistry.TransactOpts, messageIndex)
}

// TeleporterRegistryAddProtocolVersionIterator is returned from FilterAddProtocolVersion and is used to iterate over the raw logs and unpacked data for AddProtocolVersion events raised by the TeleporterRegistry contract.
type TeleporterRegistryAddProtocolVersionIterator struct {
	Event *TeleporterRegistryAddProtocolVersion // Event containing the contract specifics and raw log

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
func (it *TeleporterRegistryAddProtocolVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterRegistryAddProtocolVersion)
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
		it.Event = new(TeleporterRegistryAddProtocolVersion)
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
func (it *TeleporterRegistryAddProtocolVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterRegistryAddProtocolVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterRegistryAddProtocolVersion represents a AddProtocolVersion event raised by the TeleporterRegistry contract.
type TeleporterRegistryAddProtocolVersion struct {
	Version         *big.Int
	ProtocolAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddProtocolVersion is a free log retrieval operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_TeleporterRegistry *TeleporterRegistryFilterer) FilterAddProtocolVersion(opts *bind.FilterOpts, version []*big.Int, protocolAddress []common.Address) (*TeleporterRegistryAddProtocolVersionIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _TeleporterRegistry.contract.FilterLogs(opts, "AddProtocolVersion", versionRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryAddProtocolVersionIterator{contract: _TeleporterRegistry.contract, event: "AddProtocolVersion", logs: logs, sub: sub}, nil
}

// WatchAddProtocolVersion is a free log subscription operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_TeleporterRegistry *TeleporterRegistryFilterer) WatchAddProtocolVersion(opts *bind.WatchOpts, sink chan<- *TeleporterRegistryAddProtocolVersion, version []*big.Int, protocolAddress []common.Address) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var protocolAddressRule []interface{}
	for _, protocolAddressItem := range protocolAddress {
		protocolAddressRule = append(protocolAddressRule, protocolAddressItem)
	}

	logs, sub, err := _TeleporterRegistry.contract.WatchLogs(opts, "AddProtocolVersion", versionRule, protocolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterRegistryAddProtocolVersion)
				if err := _TeleporterRegistry.contract.UnpackLog(event, "AddProtocolVersion", log); err != nil {
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

// ParseAddProtocolVersion is a log parse operation binding the contract event 0xa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a.
//
// Solidity: event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress)
func (_TeleporterRegistry *TeleporterRegistryFilterer) ParseAddProtocolVersion(log types.Log) (*TeleporterRegistryAddProtocolVersion, error) {
	event := new(TeleporterRegistryAddProtocolVersion)
	if err := _TeleporterRegistry.contract.UnpackLog(event, "AddProtocolVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeleporterRegistryLatestVersionUpdatedIterator is returned from FilterLatestVersionUpdated and is used to iterate over the raw logs and unpacked data for LatestVersionUpdated events raised by the TeleporterRegistry contract.
type TeleporterRegistryLatestVersionUpdatedIterator struct {
	Event *TeleporterRegistryLatestVersionUpdated // Event containing the contract specifics and raw log

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
func (it *TeleporterRegistryLatestVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeleporterRegistryLatestVersionUpdated)
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
		it.Event = new(TeleporterRegistryLatestVersionUpdated)
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
func (it *TeleporterRegistryLatestVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeleporterRegistryLatestVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeleporterRegistryLatestVersionUpdated represents a LatestVersionUpdated event raised by the TeleporterRegistry contract.
type TeleporterRegistryLatestVersionUpdated struct {
	OldVersion *big.Int
	NewVersion *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLatestVersionUpdated is a free log retrieval operation binding the contract event 0x30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b0.
//
// Solidity: event LatestVersionUpdated(uint256 indexed oldVersion, uint256 indexed newVersion)
func (_TeleporterRegistry *TeleporterRegistryFilterer) FilterLatestVersionUpdated(opts *bind.FilterOpts, oldVersion []*big.Int, newVersion []*big.Int) (*TeleporterRegistryLatestVersionUpdatedIterator, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _TeleporterRegistry.contract.FilterLogs(opts, "LatestVersionUpdated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return &TeleporterRegistryLatestVersionUpdatedIterator{contract: _TeleporterRegistry.contract, event: "LatestVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchLatestVersionUpdated is a free log subscription operation binding the contract event 0x30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b0.
//
// Solidity: event LatestVersionUpdated(uint256 indexed oldVersion, uint256 indexed newVersion)
func (_TeleporterRegistry *TeleporterRegistryFilterer) WatchLatestVersionUpdated(opts *bind.WatchOpts, sink chan<- *TeleporterRegistryLatestVersionUpdated, oldVersion []*big.Int, newVersion []*big.Int) (event.Subscription, error) {

	var oldVersionRule []interface{}
	for _, oldVersionItem := range oldVersion {
		oldVersionRule = append(oldVersionRule, oldVersionItem)
	}
	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _TeleporterRegistry.contract.WatchLogs(opts, "LatestVersionUpdated", oldVersionRule, newVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeleporterRegistryLatestVersionUpdated)
				if err := _TeleporterRegistry.contract.UnpackLog(event, "LatestVersionUpdated", log); err != nil {
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

// ParseLatestVersionUpdated is a log parse operation binding the contract event 0x30623e953733f6474dabdfbef1103ce15ab73cdc77c6dfad0f9874d167e8a9b0.
//
// Solidity: event LatestVersionUpdated(uint256 indexed oldVersion, uint256 indexed newVersion)
func (_TeleporterRegistry *TeleporterRegistryFilterer) ParseLatestVersionUpdated(log types.Log) (*TeleporterRegistryLatestVersionUpdated, error) {
	event := new(TeleporterRegistryLatestVersionUpdated)
	if err := _TeleporterRegistry.contract.UnpackLog(event, "LatestVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
