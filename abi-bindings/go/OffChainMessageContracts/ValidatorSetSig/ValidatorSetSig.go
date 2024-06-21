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
	ValidatorSetSigAddress common.Address
	TargetContractAddress  common.Address
	TargetBlockChainID     [32]byte
	Nonce                  *big.Int
	TxPayload              []byte
}

// ValidatorSetSigMetaData contains all meta data concerning the ValidatorSetSig contract.
var ValidatorSetSigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validatorBlockChainID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Delivered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VALIDATORS_SOURCE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WARP_MESSENGER\",\"outputs\":[{\"internalType\":\"contractIWarpMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"messageIndex\",\"type\":\"uint32\"}],\"name\":\"executeCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContractAddress\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validatorSetSigAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"targetBlockChainID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"txPayload\",\"type\":\"bytes\"}],\"internalType\":\"structValidatorSetSigMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"validateMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b50604051610c2c380380610c2c83398101604081905261002f916100b6565b600160005560808190526040805163084279ef60e31b8152905173020000000000000000000000000000000000000591634213cf789160048083019260209291908290030181865afa158015610089573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100ad91906100b6565b60a052506100cf565b6000602082840312156100c857600080fd5b5051919050565b60805160a051610b2b61010160003960008181610124015261058801526000818160ef015261024e0152610b2b6000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638d6e579d1161005b5780638d6e579d146100ea578063b771b3bc14610111578063d127dc9b1461011f578063e5a4c2fc1461014657600080fd5b80630731775d146100825780635433da42146100a75780637ecebe00146100bc575b600080fd5b61008a600081565b6040516001600160a01b0390911681526020015b60405180910390f35b6100ba6100b53660046106e8565b610169565b005b6100dc6100ca36600461072a565b60016020526000908152604090205481565b60405190815260200161009e565b6100dc7f000000000000000000000000000000000000000000000000000000000000000081565b61008a6005600160991b0181565b6100dc7f000000000000000000000000000000000000000000000000000000000000000081565b610159610154366004610858565b610512565b604051901515815260200161009e565b61017161068f565b6040516306f8253560e41b815263ffffffff8216600482015260009081906005600160991b0190636f82535090602401600060405180830381865afa1580156101be573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526101e69190810190610972565b915091508061024a5760405162461bcd60e51b815260206004820152602560248201527f56616c696461746f725365745369673a20696e76616c69642077617270206d65604482015264737361676560d81b60648201526084015b60405180910390fd5b81517f0000000000000000000000000000000000000000000000000000000000000000146102c95760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f725365745369673a20696e76616c696420736f75726365436044820152651a185a5b925160d21b6064820152608401610241565b60208201516001600160a01b03161561033a5760405162461bcd60e51b815260206004820152602d60248201527f56616c696461746f725365745369673a206e6f6e2d7a65726f206f726967696e60448201526c53656e6465724164647265737360981b6064820152608401610241565b6040805160a0810182526000808252602082018190529181018290526060808201929092526080810191909152826040015180602001905181019061037f9190610a18565b9050600061038c82610512565b9050806103db5760405162461bcd60e51b815260206004820181905260248201527f56616c696461746f725365745369673a20696e76616c6964206d6573736167656044820152606401610241565b81606001516001600084602001516001600160a01b03166001600160a01b031681526020019081526020016000208190555081602001516001600160a01b0316826080015160405161042d9190610ab2565b6000604051808303816000865af19150503d806000811461046a576040519150601f19603f3d011682016040523d82523d6000602084013e61046f565b606091505b505080915050806104c25760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f725365745369673a2063616c6c206661696c6564000000006044820152606401610241565b816060015182602001516001600160a01b03167f5942a9a3968c7d49fc51c027041544ea295f5c1e395d6d8aa35c4369959f8ed960405160405180910390a35050505061050f6001600055565b50565b80516000906001600160a01b031630146105865760405162461bcd60e51b815260206004820152602f60248201527f56616c696461746f725365745369673a20696e76616c69642076616c6964617460448201526e6f725365745369674164647265737360881b6064820152608401610241565b7f000000000000000000000000000000000000000000000000000000000000000082604001511461060d5760405162461bcd60e51b815260206004820152602b60248201527f56616c696461746f725365745369673a20696e76616c6964207461726765744260448201526a1b1bd8dad0da185a5b925160aa1b6064820152608401610241565b60608201516020808401516001600160a01b03166000908152600191829052604090205461063a91610ace565b146106875760405162461bcd60e51b815260206004820152601e60248201527f56616c696461746f725365745369673a20696e76616c6964206e6f6e636500006044820152606401610241565b506001919050565b6002600054036106e15760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610241565b6002600055565b6000602082840312156106fa57600080fd5b813563ffffffff8116811461070e57600080fd5b9392505050565b6001600160a01b038116811461050f57600080fd5b60006020828403121561073c57600080fd5b813561070e81610715565b634e487b7160e01b600052604160045260246000fd5b60405160a0810167ffffffffffffffff8111828210171561078057610780610747565b60405290565b6040516060810167ffffffffffffffff8111828210171561078057610780610747565b604051601f8201601f1916810167ffffffffffffffff811182821017156107d2576107d2610747565b604052919050565b600067ffffffffffffffff8211156107f4576107f4610747565b50601f01601f191660200190565b600082601f83011261081357600080fd5b8135610826610821826107da565b6107a9565b81815284602083860101111561083b57600080fd5b816020850160208301376000918101602001919091529392505050565b60006020828403121561086a57600080fd5b813567ffffffffffffffff8082111561088257600080fd5b9083019060a0828603121561089657600080fd5b61089e61075d565b82356108a981610715565b815260208301356108b981610715565b8060208301525060408301356040820152606083013560608201526080830135828111156108e657600080fd5b6108f287828601610802565b60808301525095945050505050565b60005b8381101561091c578181015183820152602001610904565b50506000910152565b600082601f83011261093657600080fd5b8151610944610821826107da565b81815284602083860101111561095957600080fd5b61096a826020830160208701610901565b949350505050565b6000806040838503121561098557600080fd5b825167ffffffffffffffff8082111561099d57600080fd5b90840190606082870312156109b157600080fd5b6109b9610786565b8251815260208301516109cb81610715565b60208201526040830151828111156109e257600080fd5b6109ee88828601610925565b60408301525080945050505060208301518015158114610a0d57600080fd5b809150509250929050565b600060208284031215610a2a57600080fd5b815167ffffffffffffffff80821115610a4257600080fd5b9083019060a08286031215610a5657600080fd5b610a5e61075d565b8251610a6981610715565b81526020830151610a7981610715565b806020830152506040830151604082015260608301516060820152608083015182811115610aa657600080fd5b6108f287828601610925565b60008251610ac4818460208701610901565b9190910192915050565b80820180821115610aef57634e487b7160e01b600052601160045260246000fd5b9291505056fea2646970667358221220c97679778682b65afd450575c44d7dc688c586fe7674d8a7ab2f25b445517c4664736f6c63430008120033",
}

// ValidatorSetSigABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorSetSigMetaData.ABI instead.
var ValidatorSetSigABI = ValidatorSetSigMetaData.ABI

// ValidatorSetSigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorSetSigMetaData.Bin instead.
var ValidatorSetSigBin = ValidatorSetSigMetaData.Bin

// DeployValidatorSetSig deploys a new Ethereum contract, binding an instance of ValidatorSetSig to it.
func DeployValidatorSetSig(auth *bind.TransactOpts, backend bind.ContractBackend, validatorBlockChainID [32]byte) (common.Address, *types.Transaction, *ValidatorSetSig, error) {
	parsed, err := ValidatorSetSigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorSetSigBin), backend, validatorBlockChainID)
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

// ValidateMessage is a free data retrieval call binding the contract method 0xe5a4c2fc.
//
// Solidity: function validateMessage((address,address,bytes32,uint256,bytes) message) view returns(bool)
func (_ValidatorSetSig *ValidatorSetSigCaller) ValidateMessage(opts *bind.CallOpts, message ValidatorSetSigMessage) (bool, error) {
	var out []interface{}
	err := _ValidatorSetSig.contract.Call(opts, &out, "validateMessage", message)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateMessage is a free data retrieval call binding the contract method 0xe5a4c2fc.
//
// Solidity: function validateMessage((address,address,bytes32,uint256,bytes) message) view returns(bool)
func (_ValidatorSetSig *ValidatorSetSigSession) ValidateMessage(message ValidatorSetSigMessage) (bool, error) {
	return _ValidatorSetSig.Contract.ValidateMessage(&_ValidatorSetSig.CallOpts, message)
}

// ValidateMessage is a free data retrieval call binding the contract method 0xe5a4c2fc.
//
// Solidity: function validateMessage((address,address,bytes32,uint256,bytes) message) view returns(bool)
func (_ValidatorSetSig *ValidatorSetSigCallerSession) ValidateMessage(message ValidatorSetSigMessage) (bool, error) {
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
