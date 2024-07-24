// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validatorsetsig

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

// ValidatorSetSigMessage is an auto generated low-level Go binding around an user-defined struct.
type ValidatorSetSigMessage struct {
	TargetBlockchainID     [32]byte
	ValidatorSetSigAddress common.Address
	TargetContractAddress  common.Address
	Nonce                  *big.Int
	Payload                []byte
}

// ValidatorSetSigMetaData contains all meta data concerning the ValidatorSetSig contract.
var ValidatorSetSigMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"validatorBlockchainID_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"VALIDATORS_SOURCE_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blockchainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeCall\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"targetContractAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structValidatorSetSigMessage\",\"components\":[{\"name\":\"targetBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorSetSigAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"targetContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorBlockchainID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Delivered\",\"inputs\":[{\"name\":\"targetContractAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f80fd5b50604051610b05380380610b0583398101604081905261002e916100b2565b60015f5560808190526040805163084279ef60e31b8152905173020000000000000000000000000000000000000591634213cf789160048083019260209291908290030181865afa158015610085573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100a991906100b2565b60a052506100c9565b5f602082840312156100c2575f80fd5b5051919050565b60805160a051610a0e6100f75f395f8181610131015261046c01525f818160fc01526102330152610a0e5ff3fe608060405234801561000f575f80fd5b506004361061007a575f3560e01c80637ecebe00116100585780637ecebe00146100ca5780638d6e579d146100f7578063b771b3bc1461011e578063d127dc9b1461012c575f80fd5b80630731775d1461007e5780635433da42146100a25780635f659d8d146100b7575b5f80fd5b6100855f81565b6040516001600160a01b0390911681526020015b60405180910390f35b6100b56100b03660046105f4565b610153565b005b6100b56100c536600461073c565b610468565b6100e96100d83660046107de565b60016020525f908152604090205481565b604051908152602001610099565b6100e97f000000000000000000000000000000000000000000000000000000000000000081565b6100856005600160991b0181565b6100e97f000000000000000000000000000000000000000000000000000000000000000081565b61015b6105cc565b6040516306f8253560e41b815263ffffffff821660048201525f9081906005600160991b0190636f825350906024015f60405180830381865afa1580156101a4573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526101cb9190810190610865565b915091508061022f5760405162461bcd60e51b815260206004820152602560248201527f56616c696461746f725365745369673a20696e76616c69642077617270206d65604482015264737361676560d81b60648201526084015b60405180910390fd5b81517f0000000000000000000000000000000000000000000000000000000000000000146102ae5760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f725365745369673a20696e76616c696420736f75726365436044820152651a185a5b925160d21b6064820152608401610226565b60208201516001600160a01b03161561031f5760405162461bcd60e51b815260206004820152602d60248201527f56616c696461746f725365745369673a206e6f6e2d7a65726f206f726967696e60448201526c53656e6465724164647265737360981b6064820152608401610226565b5f82604001518060200190518101906103389190610905565b905061034381610468565b6060810151610353906001610998565b604080830180516001600160a01b039081165f90815260016020528381209490945590516080850151925191169161038a916109bd565b5f604051808303815f865af19150503d805f81146103c3576040519150601f19603f3d011682016040523d82523d5f602084013e6103c8565b606091505b50509050806104195760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f725365745369673a2063616c6c206661696c6564000000006044820152606401610226565b816060015182604001516001600160a01b03167f5942a9a3968c7d49fc51c027041544ea295f5c1e395d6d8aa35c4369959f8ed960405160405180910390a35050505061046560015f55565b50565b80517f0000000000000000000000000000000000000000000000000000000000000000146104ec5760405162461bcd60e51b815260206004820152602b60248201527f56616c696461746f725365745369673a20696e76616c6964207461726765744260448201526a1b1bd8dad8da185a5b925160aa1b6064820152608401610226565b60208101516001600160a01b031630146105605760405162461bcd60e51b815260206004820152602f60248201527f56616c696461746f725365745369673a20696e76616c69642076616c6964617460448201526e6f725365745369674164647265737360881b6064820152608401610226565b60608101516040808301516001600160a01b03165f90815260016020522054146104655760405162461bcd60e51b815260206004820152601e60248201527f56616c696461746f725365745369673a20696e76616c6964206e6f6e636500006044820152606401610226565b60025f54036105ee57604051633ee5aeb560e01b815260040160405180910390fd5b60025f55565b5f60208284031215610604575f80fd5b813563ffffffff81168114610617575f80fd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b60405160a0810167ffffffffffffffff811182821017156106555761065561061e565b60405290565b6040516060810167ffffffffffffffff811182821017156106555761065561061e565b604051601f8201601f1916810167ffffffffffffffff811182821017156106a7576106a761061e565b604052919050565b6001600160a01b0381168114610465575f80fd5b5f67ffffffffffffffff8211156106dc576106dc61061e565b50601f01601f191660200190565b5f82601f8301126106f9575f80fd5b813561070c610707826106c3565b61067e565b818152846020838601011115610720575f80fd5b816020850160208301375f918101602001919091529392505050565b5f6020828403121561074c575f80fd5b813567ffffffffffffffff80821115610763575f80fd5b9083019060a08286031215610776575f80fd5b61077e610632565b823581526020830135610790816106af565b602082015260408301356107a3816106af565b6040820152606083810135908201526080830135828111156107c3575f80fd5b6107cf878286016106ea565b60808301525095945050505050565b5f602082840312156107ee575f80fd5b8135610617816106af565b5f5b838110156108135781810151838201526020016107fb565b50505f910152565b5f82601f83011261082a575f80fd5b8151610838610707826106c3565b81815284602083860101111561084c575f80fd5b61085d8260208301602087016107f9565b949350505050565b5f8060408385031215610876575f80fd5b825167ffffffffffffffff8082111561088d575f80fd5b90840190606082870312156108a0575f80fd5b6108a861065b565b8251815260208301516108ba816106af565b60208201526040830151828111156108d0575f80fd5b6108dc8882860161081b565b604083015250809450505050602083015180151581146108fa575f80fd5b809150509250929050565b5f60208284031215610915575f80fd5b815167ffffffffffffffff8082111561092c575f80fd5b9083019060a0828603121561093f575f80fd5b610947610632565b825181526020830151610959816106af565b6020820152604083015161096c816106af565b60408201526060838101519082015260808301518281111561098c575f80fd5b6107cf8782860161081b565b808201808211156109b757634e487b7160e01b5f52601160045260245ffd5b92915050565b5f82516109ce8184602087016107f9565b919091019291505056fea26469706673582212207c2f37bfaac3d19c6939f706062920b1511cb3cdfc0f9c0f54cc943a5f8b005f64736f6c63430008190033",
}

// ValidatorSetSigABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorSetSigMetaData.ABI instead.
var ValidatorSetSigABI = ValidatorSetSigMetaData.ABI

// ValidatorSetSigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorSetSigMetaData.Bin instead.
var ValidatorSetSigBin = ValidatorSetSigMetaData.Bin

// DeployValidatorSetSig deploys a new Ethereum contract, binding an instance of ValidatorSetSig to it.
func DeployValidatorSetSig(auth *bind.TransactOpts, backend bind.ContractBackend, validatorBlockchainID_ [32]byte) (common.Address, *types.Transaction, *ValidatorSetSig, error) {
	parsed, err := ValidatorSetSigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorSetSigBin), backend, validatorBlockchainID_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorSetSig{ValidatorSetSigCaller: ValidatorSetSigCaller{contract: contract}, ValidatorSetSigTransactor: ValidatorSetSigTransactor{contract: contract}, ValidatorSetSigFilterer: ValidatorSetSigFilterer{contract: contract}}, nil
}

// ValidatorSetSig is an auto generated Go binding around an Ethereum contract.
type ValidatorSetSig struct {
	ValidatorSetSigCaller     // Read-only binding to the contract
	ValidatorSetSigTransactor // Write-only binding to the contract
	ValidatorSetSigFilterer   // Log filterer for contract events
}

// ValidatorSetSigCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorSetSigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorSetSigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorSetSigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSetSigSession struct {
	Contract     *ValidatorSetSig  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorSetSigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorSetSigCallerSession struct {
	Contract *ValidatorSetSigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ValidatorSetSigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorSetSigTransactorSession struct {
	Contract     *ValidatorSetSigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ValidatorSetSigRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorSetSigRaw struct {
	Contract *ValidatorSetSig // Generic contract binding to access the raw methods on
}

// ValidatorSetSigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorSetSigCallerRaw struct {
	Contract *ValidatorSetSigCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorSetSigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorSetSigTransactorRaw struct {
	Contract *ValidatorSetSigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorSetSig creates a new instance of ValidatorSetSig, bound to a specific deployed contract.
func NewValidatorSetSig(address common.Address, backend bind.ContractBackend) (*ValidatorSetSig, error) {
	contract, err := bindValidatorSetSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSig{ValidatorSetSigCaller: ValidatorSetSigCaller{contract: contract}, ValidatorSetSigTransactor: ValidatorSetSigTransactor{contract: contract}, ValidatorSetSigFilterer: ValidatorSetSigFilterer{contract: contract}}, nil
}

// NewValidatorSetSigCaller creates a new read-only instance of ValidatorSetSig, bound to a specific deployed contract.
func NewValidatorSetSigCaller(address common.Address, caller bind.ContractCaller) (*ValidatorSetSigCaller, error) {
	contract, err := bindValidatorSetSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSigCaller{contract: contract}, nil
}

// NewValidatorSetSigTransactor creates a new write-only instance of ValidatorSetSig, bound to a specific deployed contract.
func NewValidatorSetSigTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorSetSigTransactor, error) {
	contract, err := bindValidatorSetSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSigTransactor{contract: contract}, nil
}

// NewValidatorSetSigFilterer creates a new log filterer instance of ValidatorSetSig, bound to a specific deployed contract.
func NewValidatorSetSigFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorSetSigFilterer, error) {
	contract, err := bindValidatorSetSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSigFilterer{contract: contract}, nil
}

// bindValidatorSetSig binds a generic wrapper to an already deployed contract.
func bindValidatorSetSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorSetSigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSetSig *ValidatorSetSigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorSetSig.Contract.ValidatorSetSigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSetSig *ValidatorSetSigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.ValidatorSetSigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSetSig *ValidatorSetSigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.ValidatorSetSigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSetSig *ValidatorSetSigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorSetSig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSetSig *ValidatorSetSigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSetSig *ValidatorSetSigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.contract.Transact(opts, method, params...)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigCaller) VALIDATORSSOURCEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "VALIDATORS_SOURCE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigSession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _ValidatorSetSig.Contract.VALIDATORSSOURCEADDRESS(&_ValidatorSetSig.CallOpts)
}

// VALIDATORSSOURCEADDRESS is a free data retrieval call binding the contract method 0x0731775d.
//
// Solidity: function VALIDATORS_SOURCE_ADDRESS() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) VALIDATORSSOURCEADDRESS() (common.Address, error) {
	return _ValidatorSetSig.Contract.VALIDATORSSOURCEADDRESS(&_ValidatorSetSig.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigSession) WARPMESSENGER() (common.Address, error) {
	return _ValidatorSetSig.Contract.WARPMESSENGER(&_ValidatorSetSig.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) WARPMESSENGER() (common.Address, error) {
	return _ValidatorSetSig.Contract.WARPMESSENGER(&_ValidatorSetSig.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigSession) BlockchainID() ([32]byte, error) {
	return _ValidatorSetSig.Contract.BlockchainID(&_ValidatorSetSig.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) BlockchainID() ([32]byte, error) {
	return _ValidatorSetSig.Contract.BlockchainID(&_ValidatorSetSig.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address targetContractAddress) view returns(uint256 nonce)
func (_ValidatorSetSig *ValidatorSetSigCaller) Nonces(opts *bind.CallOpts, targetContractAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "nonces", targetContractAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address targetContractAddress) view returns(uint256 nonce)
func (_ValidatorSetSig *ValidatorSetSigSession) Nonces(targetContractAddress common.Address) (*big.Int, error) {
	return _ValidatorSetSig.Contract.Nonces(&_ValidatorSetSig.CallOpts, targetContractAddress)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address targetContractAddress) view returns(uint256 nonce)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) Nonces(targetContractAddress common.Address) (*big.Int, error) {
	return _ValidatorSetSig.Contract.Nonces(&_ValidatorSetSig.CallOpts, targetContractAddress)
}

// ValidateMessage is a free data retrieval call binding the contract method 0x5f659d8d.
//
// Solidity: function validateMessage((bytes32,address,address,uint256,bytes) message) view returns()
func (_ValidatorSetSig *ValidatorSetSigCaller) ValidateMessage(opts *bind.CallOpts, message ValidatorSetSigMessage) error {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "validateMessage", message)

	if err != nil {
		return err
	}

	return err

}

// ValidateMessage is a free data retrieval call binding the contract method 0x5f659d8d.
//
// Solidity: function validateMessage((bytes32,address,address,uint256,bytes) message) view returns()
func (_ValidatorSetSig *ValidatorSetSigSession) ValidateMessage(message ValidatorSetSigMessage) error {
	return _ValidatorSetSig.Contract.ValidateMessage(&_ValidatorSetSig.CallOpts, message)
}

// ValidateMessage is a free data retrieval call binding the contract method 0x5f659d8d.
//
// Solidity: function validateMessage((bytes32,address,address,uint256,bytes) message) view returns()
func (_ValidatorSetSig *ValidatorSetSigCallerSession) ValidateMessage(message ValidatorSetSigMessage) error {
	return _ValidatorSetSig.Contract.ValidateMessage(&_ValidatorSetSig.CallOpts, message)
}

// ValidatorBlockchainID is a free data retrieval call binding the contract method 0x8d6e579d.
//
// Solidity: function validatorBlockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigCaller) ValidatorBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "validatorBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ValidatorBlockchainID is a free data retrieval call binding the contract method 0x8d6e579d.
//
// Solidity: function validatorBlockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigSession) ValidatorBlockchainID() ([32]byte, error) {
	return _ValidatorSetSig.Contract.ValidatorBlockchainID(&_ValidatorSetSig.CallOpts)
}

// ValidatorBlockchainID is a free data retrieval call binding the contract method 0x8d6e579d.
//
// Solidity: function validatorBlockchainID() view returns(bytes32)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) ValidatorBlockchainID() ([32]byte, error) {
	return _ValidatorSetSig.Contract.ValidatorBlockchainID(&_ValidatorSetSig.CallOpts)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x5433da42.
//
// Solidity: function executeCall(uint32 messageIndex) returns()
func (_ValidatorSetSig *ValidatorSetSigTransactor) ExecuteCall(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorSetSig.contract.Transact(opts, "executeCall", messageIndex)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x5433da42.
//
// Solidity: function executeCall(uint32 messageIndex) returns()
func (_ValidatorSetSig *ValidatorSetSigSession) ExecuteCall(messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.ExecuteCall(&_ValidatorSetSig.TransactOpts, messageIndex)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x5433da42.
//
// Solidity: function executeCall(uint32 messageIndex) returns()
func (_ValidatorSetSig *ValidatorSetSigTransactorSession) ExecuteCall(messageIndex uint32) (*types.Transaction, error) {
	return _ValidatorSetSig.Contract.ExecuteCall(&_ValidatorSetSig.TransactOpts, messageIndex)
}

// ValidatorSetSigDeliveredIterator is returned from FilterDelivered and is used to iterate over the raw logs and unpacked data for Delivered events raised by the ValidatorSetSig contract.
type ValidatorSetSigDeliveredIterator struct {
	Event *ValidatorSetSigDelivered // Event containing the contract specifics and raw log

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
func (it *ValidatorSetSigDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSetSigDelivered)
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
		it.Event = new(ValidatorSetSigDelivered)
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
func (it *ValidatorSetSigDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSetSigDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSetSigDelivered represents a Delivered event raised by the ValidatorSetSig contract.
type ValidatorSetSigDelivered struct {
	TargetContractAddress common.Address
	Nonce                 *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDelivered is a free log retrieval operation binding the contract event 0x5942a9a3968c7d49fc51c027041544ea295f5c1e395d6d8aa35c4369959f8ed9.
//
// Solidity: event Delivered(address indexed targetContractAddress, uint256 indexed nonce)
func (_ValidatorSetSig *ValidatorSetSigFilterer) FilterDelivered(opts *bind.FilterOpts, targetContractAddress []common.Address, nonce []*big.Int) (*ValidatorSetSigDeliveredIterator, error) {

	var targetContractAddressRule []interface{}
	for _, targetContractAddressItem := range targetContractAddress {
		targetContractAddressRule = append(targetContractAddressRule, targetContractAddressItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ValidatorSetSig.contract.FilterLogs(opts, "Delivered", targetContractAddressRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetSigDeliveredIterator{contract: _ValidatorSetSig.contract, event: "Delivered", logs: logs, sub: sub}, nil
}

// WatchDelivered is a free log subscription operation binding the contract event 0x5942a9a3968c7d49fc51c027041544ea295f5c1e395d6d8aa35c4369959f8ed9.
//
// Solidity: event Delivered(address indexed targetContractAddress, uint256 indexed nonce)
func (_ValidatorSetSig *ValidatorSetSigFilterer) WatchDelivered(opts *bind.WatchOpts, sink chan<- *ValidatorSetSigDelivered, targetContractAddress []common.Address, nonce []*big.Int) (event.Subscription, error) {

	var targetContractAddressRule []interface{}
	for _, targetContractAddressItem := range targetContractAddress {
		targetContractAddressRule = append(targetContractAddressRule, targetContractAddressItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ValidatorSetSig.contract.WatchLogs(opts, "Delivered", targetContractAddressRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSetSigDelivered)
				if err := _ValidatorSetSig.contract.UnpackLog(event, "Delivered", log); err != nil {
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

// ParseDelivered is a log parse operation binding the contract event 0x5942a9a3968c7d49fc51c027041544ea295f5c1e395d6d8aa35c4369959f8ed9.
//
// Solidity: event Delivered(address indexed targetContractAddress, uint256 indexed nonce)
func (_ValidatorSetSig *ValidatorSetSigFilterer) ParseDelivered(log types.Log) (*ValidatorSetSigDelivered, error) {
	event := new(ValidatorSetSigDelivered)
	if err := _ValidatorSetSig.contract.UnpackLog(event, "Delivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
