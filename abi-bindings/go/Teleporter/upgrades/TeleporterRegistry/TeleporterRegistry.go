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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"internalType\":\"structProtocolRegistryEntry[]\",\"name\":\"initialEntries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"name\":\"AddProtocolVersion\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VALIDATORS_SOURCE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"addProtocolVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getAddressFromVersion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestTeleporter\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"getTeleporterFromVersion\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"protocolAddress\",\"type\":\"address\"}],\"name\":\"getVersionFromAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162000e3b38038062000e3b833981016040819052620000349162000357565b807302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000088573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000ae919062000449565b60805260008080555b81518110156200010357620000ee828281518110620000da57620000da62000463565b60200260200101516200010c60201b60201c565b80620000fa8162000479565b915050620000b7565b505050620004a1565b80516000036200016e5760405162461bcd60e51b815260206004820152602260248201527f5761727050726f746f636f6c52656769737472793a207a65726f20766572736960448201526137b760f11b60648201526084015b60405180910390fd5b80516000908152600160205260409020546001600160a01b031615620001ec5760405162461bcd60e51b815260206004820152602c60248201527f5761727050726f746f636f6c52656769737472793a2076657273696f6e20616c60448201526b72656164792065786973747360a01b606482015260840162000165565b60208101516001600160a01b03166200025c5760405162461bcd60e51b815260206004820152602b60248201527f5761727050726f746f636f6c52656769737472793a207a65726f2070726f746f60448201526a636f6c206164647265737360a81b606482015260840162000165565b602081810180518351600090815260018452604080822080546001600160a01b0319166001600160a01b039485161790558551845184168352600290955280822094909455915184519351911692917fa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a91a360005481511115620002e05780516000555b50565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156200031e576200031e620002e3565b60405290565b604051601f8201601f191681016001600160401b03811182821017156200034f576200034f620002e3565b604052919050565b600060208083850312156200036b57600080fd5b82516001600160401b03808211156200038357600080fd5b818501915085601f8301126200039857600080fd5b815181811115620003ad57620003ad620002e3565b620003bd848260051b0162000324565b818152848101925060069190911b830184019087821115620003de57600080fd5b928401925b818410156200043e5760408489031215620003fe5760008081fd5b62000408620002f9565b84518152858501516001600160a01b0381168114620004275760008081fd5b8187015283526040939093019291840191620003e3565b979650505050505050565b6000602082840312156200045c57600080fd5b5051919050565b634e487b7160e01b600052603260045260246000fd5b6000600182016200049a57634e487b7160e01b600052601160045260246000fd5b5060010190565b608051610977620004c46000396000818161021e015261031901526109776000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806346f9ef491161005b57806346f9ef49146100c45780634c1f08ce146100ec578063b771b3bc14610115578063d820e64f1461012357600080fd5b80630731775d1461008d5780630e6d1de9146100b2578063215abce9146100c457806341f34ed9146100d7575b600080fd5b610095600081565b6040516001600160a01b0390911681526020015b60405180910390f35b6000545b6040519081526020016100a9565b6100956100d2366004610687565b61012b565b6100ea6100e53660046106a0565b61013c565b005b6100b66100fa3660046106e2565b6001600160a01b031660009081526002602052604090205490565b6100956005600160991b0181565b610095610446565b600061013682610458565b92915050565b6040516306f8253560e41b815263ffffffff8216600482015260009081906005600160991b0190636f82535090602401600060405180830381865afa158015610189573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526101b19190810190610784565b915091508061021a5760405162461bcd60e51b815260206004820152602a60248201527f5761727050726f746f636f6c52656769737472793a20696e76616c69642077616044820152697270206d65737361676560b01b60648201526084015b60405180910390fd5b81517f0000000000000000000000000000000000000000000000000000000000000000146102a05760405162461bcd60e51b815260206004820152602d60248201527f5761727050726f746f636f6c52656769737472793a20696e76616c696420736f60448201526c1d5c98d94818da185a5b881251609a1b6064820152608401610211565b60208201516001600160a01b0316156103175760405162461bcd60e51b815260206004820152603360248201527f5761727050726f746f636f6c52656769737472793a20696e76616c6964206f726044820152726967696e2073656e646572206164647265737360681b6064820152608401610211565b7f00000000000000000000000000000000000000000000000000000000000000008260400151146103a55760405162461bcd60e51b815260206004820152603260248201527f5761727050726f746f636f6c52656769737472793a20696e76616c69642064656044820152711cdd1a5b985d1a5bdb8818da185a5b88125160721b6064820152608401610211565b60608201516001600160a01b0316301461041b5760405162461bcd60e51b815260206004820152603160248201527f5761727050726f746f636f6c52656769737472793a20696e76616c69642064656044820152707374696e6174696f6e206164647265737360781b6064820152608401610211565b6000826080015180602001905181019061043591906108a9565b9050610440816104f6565b50505050565b6000610453600054610458565b905090565b60008160000361047a5760405162461bcd60e51b8152600401610211906108ff565b6000548211156104da5760405162461bcd60e51b815260206004820152602560248201527f5761727050726f746f636f6c52656769737472793a20696e76616c69642076656044820152643939b4b7b760d91b6064820152608401610211565b506000908152600160205260409020546001600160a01b031690565b80516000036105175760405162461bcd60e51b8152600401610211906108ff565b80516000908152600160205260409020546001600160a01b0316156105935760405162461bcd60e51b815260206004820152602c60248201527f5761727050726f746f636f6c52656769737472793a2076657273696f6e20616c60448201526b72656164792065786973747360a01b6064820152608401610211565b60208101516001600160a01b03166106015760405162461bcd60e51b815260206004820152602b60248201527f5761727050726f746f636f6c52656769737472793a207a65726f2070726f746f60448201526a636f6c206164647265737360a81b6064820152608401610211565b602081810180518351600090815260018452604080822080546001600160a01b0319166001600160a01b039485161790558551845184168352600290955280822094909455915184519351911692917fa5eed93d951a9603d5f7c0a57de79a299dd3dbd5e51429be209d8053a42ab43a91a3600054815111156106845780516000555b50565b60006020828403121561069957600080fd5b5035919050565b6000602082840312156106b257600080fd5b813563ffffffff811681146106c657600080fd5b9392505050565b6001600160a01b038116811461068457600080fd5b6000602082840312156106f457600080fd5b81356106c6816106cd565b634e487b7160e01b600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715610738576107386106ff565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610767576107676106ff565b604052919050565b8051801515811461077f57600080fd5b919050565b6000806040838503121561079757600080fd5b825167ffffffffffffffff808211156107af57600080fd5b9084019060a082870312156107c357600080fd5b6107cb610715565b825181526020808401516107de816106cd565b828201526040848101519083015260608401516107fa816106cd565b606083015260808401518381111561081157600080fd5b80850194505087601f85011261082657600080fd5b835183811115610838576108386106ff565b61084a601f8201601f1916830161073e565b9350808452888282870101111561086057600080fd5b60005b8181101561087e578581018301518582018401528201610863565b5060008282860101525082608083015281955061089c81880161076f565b9450505050509250929050565b6000604082840312156108bb57600080fd5b6040516040810181811067ffffffffffffffff821117156108de576108de6106ff565b6040528251815260208301516108f3816106cd565b60208201529392505050565b60208082526022908201527f5761727050726f746f636f6c52656769737472793a207a65726f20766572736960408201526137b760f11b60608201526080019056fea26469706673582212208d425d0954c81b922dbffdcb4592fc6dda68b1b6b7296ee1d73e880aa76e06fe64736f6c63430008120033",
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

// GetLatestVersion is a free data retrieval call binding the contract method 0x0e6d1de9.
//
// Solidity: function getLatestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCaller) GetLatestVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeleporterRegistry.contract.Call(opts, &out, "getLatestVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestVersion is a free data retrieval call binding the contract method 0x0e6d1de9.
//
// Solidity: function getLatestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistrySession) GetLatestVersion() (*big.Int, error) {
	return _TeleporterRegistry.Contract.GetLatestVersion(&_TeleporterRegistry.CallOpts)
}

// GetLatestVersion is a free data retrieval call binding the contract method 0x0e6d1de9.
//
// Solidity: function getLatestVersion() view returns(uint256)
func (_TeleporterRegistry *TeleporterRegistryCallerSession) GetLatestVersion() (*big.Int, error) {
	return _TeleporterRegistry.Contract.GetLatestVersion(&_TeleporterRegistry.CallOpts)
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
