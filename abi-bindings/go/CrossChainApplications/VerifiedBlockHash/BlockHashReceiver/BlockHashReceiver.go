// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockhashreceiver

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

// BlockHashReceiverMetaData contains all meta data concerning the BlockHashReceiver contract.
var BlockHashReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"publisherChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"publisherContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"ReceiveBlockHash\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getLatestBlockInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originChainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceChainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourcePublisherContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b506040516109ec3803806109ec83398101604081905261002f916101b3565b82806001600160a01b0381166100b15760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f727465722072656769737472792061646472657373000000000000000000606482015260840160405180910390fd5b6001600160a01b038116608081905260408051630e6d1de960e01b81529051630e6d1de9916004808201926020929091908290030181865afa1580156100fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061011f91906101ef565b6000555061012c33610145565b5060a0919091526001600160a01b031660c05250610208565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80516001600160a01b03811681146101ae57600080fd5b919050565b6000806000606084860312156101c857600080fd5b6101d184610197565b9250602084015191506101e660408501610197565b90509250925092565b60006020828403121561020157600080fd5b5051919050565b60805160a05160c05161079a6102526000396000818161014a015261043701526000818161010201526103b501526000818160be015281816101fc01526102d4015261079a6000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063b17810be11610071578063b17810be1461017d578063b6109d9d14610198578063c868efaa146101a0578063e49cc553146101b3578063f2fde38b146101bc578063f3f39ee5146101cf57600080fd5b80631a7f5bec146100b95780634c335368146100fd5780636c4f6ba914610132578063715018a61461013b57806379a0710c146101455780638da5cb5b1461016c575b600080fd5b6100e07f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6101247f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100f4565b61012460035481565b6101436101d8565b005b6100e07f000000000000000000000000000000000000000000000000000000000000000081565b6001546001600160a01b03166100e0565b600254600354604080519283526020830191909152016100f4565b6101436101ec565b6101436101ae366004610680565b6102bc565b61012460005481565b6101436101ca366004610707565b610542565b61012460025481565b6101e06105b8565b6101ea6000610612565b565b6101f46105b8565b6000805490507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630e6d1de96040518163ffffffff1660e01b8152600401602060405180830381865afa158015610258573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061027c9190610729565b60008190558110156102b957600054817fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d60405160405180910390a35b50565b60005460405163260f846760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690634c1f08ce90602401602060405180830381865afa158015610323573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103479190610729565b10156103b35760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964207460448201526f32b632b837b93a32b91039b2b73232b960811b60648201526084015b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000084146104355760405162461bcd60e51b815260206004820152602a60248201527f426c6f636b4861736852656365697665723a20696e76616c696420736f757263604482015269194818da185a5b88125160b21b60648201526084016103aa565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316146104d05760405162461bcd60e51b815260206004820152603160248201527f426c6f636b4861736852656365697665723a20696e76616c696420736f757263604482015270329031b430b4b710383ab13634b9b432b960791b60648201526084016103aa565b6000806104df83850185610742565b9150915060025482111561053a576002829055600381905560405181815282906001600160a01b0387169088907f0bca78aa82d7575f42e4b4b2fe04765a4b2f3661786403788ce987e065ac590d9060200160405180910390a45b505050505050565b61054a6105b8565b6001600160a01b0381166105af5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103aa565b6102b981610612565b6001546001600160a01b031633146101ea5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103aa565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80356001600160a01b038116811461067b57600080fd5b919050565b6000806000806060858703121561069657600080fd5b843593506106a660208601610664565b9250604085013567ffffffffffffffff808211156106c357600080fd5b818701915087601f8301126106d757600080fd5b8135818111156106e657600080fd5b8860208285010111156106f857600080fd5b95989497505060200194505050565b60006020828403121561071957600080fd5b61072282610664565b9392505050565b60006020828403121561073b57600080fd5b5051919050565b6000806040838503121561075557600080fd5b5050803592602090910135915056fea26469706673582212207a6bb017fd95d20104c259d457e7ac7b06b46f18e1e79360a5623a6e74e7815364736f6c63430008120033",
}

// BlockHashReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockHashReceiverMetaData.ABI instead.
var BlockHashReceiverABI = BlockHashReceiverMetaData.ABI

// BlockHashReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlockHashReceiverMetaData.Bin instead.
var BlockHashReceiverBin = BlockHashReceiverMetaData.Bin

// DeployBlockHashReceiver deploys a new Ethereum contract, binding an instance of BlockHashReceiver to it.
func DeployBlockHashReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterRegistryAddress common.Address, publisherChainID [32]byte, publisherContractAddress common.Address) (common.Address, *types.Transaction, *BlockHashReceiver, error) {
	parsed, err := BlockHashReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlockHashReceiverBin), backend, teleporterRegistryAddress, publisherChainID, publisherContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlockHashReceiver{BlockHashReceiverCaller: BlockHashReceiverCaller{contract: contract}, BlockHashReceiverTransactor: BlockHashReceiverTransactor{contract: contract}, BlockHashReceiverFilterer: BlockHashReceiverFilterer{contract: contract}}, nil
}

// BlockHashReceiver is an auto generated Go binding around an Ethereum contract.
type BlockHashReceiver struct {
	BlockHashReceiverCaller     // Read-only binding to the contract
	BlockHashReceiverTransactor // Write-only binding to the contract
	BlockHashReceiverFilterer   // Log filterer for contract events
}

// BlockHashReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockHashReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockHashReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockHashReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockHashReceiverSession struct {
	Contract     *BlockHashReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BlockHashReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockHashReceiverCallerSession struct {
	Contract *BlockHashReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BlockHashReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockHashReceiverTransactorSession struct {
	Contract     *BlockHashReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BlockHashReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockHashReceiverRaw struct {
	Contract *BlockHashReceiver // Generic contract binding to access the raw methods on
}

// BlockHashReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockHashReceiverCallerRaw struct {
	Contract *BlockHashReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// BlockHashReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockHashReceiverTransactorRaw struct {
	Contract *BlockHashReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockHashReceiver creates a new instance of BlockHashReceiver, bound to a specific deployed contract.
func NewBlockHashReceiver(address common.Address, backend bind.ContractBackend) (*BlockHashReceiver, error) {
	contract, err := bindBlockHashReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiver{BlockHashReceiverCaller: BlockHashReceiverCaller{contract: contract}, BlockHashReceiverTransactor: BlockHashReceiverTransactor{contract: contract}, BlockHashReceiverFilterer: BlockHashReceiverFilterer{contract: contract}}, nil
}

// NewBlockHashReceiverCaller creates a new read-only instance of BlockHashReceiver, bound to a specific deployed contract.
func NewBlockHashReceiverCaller(address common.Address, caller bind.ContractCaller) (*BlockHashReceiverCaller, error) {
	contract, err := bindBlockHashReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverCaller{contract: contract}, nil
}

// NewBlockHashReceiverTransactor creates a new write-only instance of BlockHashReceiver, bound to a specific deployed contract.
func NewBlockHashReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockHashReceiverTransactor, error) {
	contract, err := bindBlockHashReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverTransactor{contract: contract}, nil
}

// NewBlockHashReceiverFilterer creates a new log filterer instance of BlockHashReceiver, bound to a specific deployed contract.
func NewBlockHashReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockHashReceiverFilterer, error) {
	contract, err := bindBlockHashReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverFilterer{contract: contract}, nil
}

// bindBlockHashReceiver binds a generic wrapper to an already deployed contract.
func bindBlockHashReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockHashReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockHashReceiver *BlockHashReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockHashReceiver.Contract.BlockHashReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockHashReceiver *BlockHashReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.BlockHashReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockHashReceiver *BlockHashReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.BlockHashReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockHashReceiver *BlockHashReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockHashReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockHashReceiver *BlockHashReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockHashReceiver *BlockHashReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.contract.Transact(opts, method, params...)
}

// GetLatestBlockInfo is a free data retrieval call binding the contract method 0xb17810be.
//
// Solidity: function getLatestBlockInfo() view returns(uint256 height, bytes32 hash)
func (_BlockHashReceiver *BlockHashReceiverCaller) GetLatestBlockInfo(opts *bind.CallOpts) (struct {
	Height *big.Int
	Hash   [32]byte
}, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "getLatestBlockInfo")

	outstruct := new(struct {
		Height *big.Int
		Hash   [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Height = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Hash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetLatestBlockInfo is a free data retrieval call binding the contract method 0xb17810be.
//
// Solidity: function getLatestBlockInfo() view returns(uint256 height, bytes32 hash)
func (_BlockHashReceiver *BlockHashReceiverSession) GetLatestBlockInfo() (struct {
	Height *big.Int
	Hash   [32]byte
}, error) {
	return _BlockHashReceiver.Contract.GetLatestBlockInfo(&_BlockHashReceiver.CallOpts)
}

// GetLatestBlockInfo is a free data retrieval call binding the contract method 0xb17810be.
//
// Solidity: function getLatestBlockInfo() view returns(uint256 height, bytes32 hash)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) GetLatestBlockInfo() (struct {
	Height *big.Int
	Hash   [32]byte
}, error) {
	return _BlockHashReceiver.Contract.GetLatestBlockInfo(&_BlockHashReceiver.CallOpts)
}

// LatestBlockHash is a free data retrieval call binding the contract method 0x6c4f6ba9.
//
// Solidity: function latestBlockHash() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverCaller) LatestBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "latestBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestBlockHash is a free data retrieval call binding the contract method 0x6c4f6ba9.
//
// Solidity: function latestBlockHash() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverSession) LatestBlockHash() ([32]byte, error) {
	return _BlockHashReceiver.Contract.LatestBlockHash(&_BlockHashReceiver.CallOpts)
}

// LatestBlockHash is a free data retrieval call binding the contract method 0x6c4f6ba9.
//
// Solidity: function latestBlockHash() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) LatestBlockHash() ([32]byte, error) {
	return _BlockHashReceiver.Contract.LatestBlockHash(&_BlockHashReceiver.CallOpts)
}

// LatestBlockHeight is a free data retrieval call binding the contract method 0xf3f39ee5.
//
// Solidity: function latestBlockHeight() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverCaller) LatestBlockHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "latestBlockHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockHeight is a free data retrieval call binding the contract method 0xf3f39ee5.
//
// Solidity: function latestBlockHeight() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverSession) LatestBlockHeight() (*big.Int, error) {
	return _BlockHashReceiver.Contract.LatestBlockHeight(&_BlockHashReceiver.CallOpts)
}

// LatestBlockHeight is a free data retrieval call binding the contract method 0xf3f39ee5.
//
// Solidity: function latestBlockHeight() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) LatestBlockHeight() (*big.Int, error) {
	return _BlockHashReceiver.Contract.LatestBlockHeight(&_BlockHashReceiver.CallOpts)
}

// MinTeleporterVersion is a free data retrieval call binding the contract method 0xe49cc553.
//
// Solidity: function minTeleporterVersion() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverCaller) MinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "minTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTeleporterVersion is a free data retrieval call binding the contract method 0xe49cc553.
//
// Solidity: function minTeleporterVersion() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverSession) MinTeleporterVersion() (*big.Int, error) {
	return _BlockHashReceiver.Contract.MinTeleporterVersion(&_BlockHashReceiver.CallOpts)
}

// MinTeleporterVersion is a free data retrieval call binding the contract method 0xe49cc553.
//
// Solidity: function minTeleporterVersion() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) MinTeleporterVersion() (*big.Int, error) {
	return _BlockHashReceiver.Contract.MinTeleporterVersion(&_BlockHashReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverSession) Owner() (common.Address, error) {
	return _BlockHashReceiver.Contract.Owner(&_BlockHashReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) Owner() (common.Address, error) {
	return _BlockHashReceiver.Contract.Owner(&_BlockHashReceiver.CallOpts)
}

// SourceChainID is a free data retrieval call binding the contract method 0x4c335368.
//
// Solidity: function sourceChainID() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverCaller) SourceChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "sourceChainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SourceChainID is a free data retrieval call binding the contract method 0x4c335368.
//
// Solidity: function sourceChainID() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverSession) SourceChainID() ([32]byte, error) {
	return _BlockHashReceiver.Contract.SourceChainID(&_BlockHashReceiver.CallOpts)
}

// SourceChainID is a free data retrieval call binding the contract method 0x4c335368.
//
// Solidity: function sourceChainID() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) SourceChainID() ([32]byte, error) {
	return _BlockHashReceiver.Contract.SourceChainID(&_BlockHashReceiver.CallOpts)
}

// SourcePublisherContractAddress is a free data retrieval call binding the contract method 0x79a0710c.
//
// Solidity: function sourcePublisherContractAddress() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverCaller) SourcePublisherContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "sourcePublisherContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SourcePublisherContractAddress is a free data retrieval call binding the contract method 0x79a0710c.
//
// Solidity: function sourcePublisherContractAddress() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverSession) SourcePublisherContractAddress() (common.Address, error) {
	return _BlockHashReceiver.Contract.SourcePublisherContractAddress(&_BlockHashReceiver.CallOpts)
}

// SourcePublisherContractAddress is a free data retrieval call binding the contract method 0x79a0710c.
//
// Solidity: function sourcePublisherContractAddress() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) SourcePublisherContractAddress() (common.Address, error) {
	return _BlockHashReceiver.Contract.SourcePublisherContractAddress(&_BlockHashReceiver.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverSession) TeleporterRegistry() (common.Address, error) {
	return _BlockHashReceiver.Contract.TeleporterRegistry(&_BlockHashReceiver.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) TeleporterRegistry() (common.Address, error) {
	return _BlockHashReceiver.Contract.TeleporterRegistry(&_BlockHashReceiver.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _BlockHashReceiver.contract.Transact(opts, "receiveTeleporterMessage", originChainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_BlockHashReceiver *BlockHashReceiverSession) ReceiveTeleporterMessage(originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.ReceiveTeleporterMessage(&_BlockHashReceiver.TransactOpts, originChainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originChainID, address originSenderAddress, bytes message) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactorSession) ReceiveTeleporterMessage(originChainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.ReceiveTeleporterMessage(&_BlockHashReceiver.TransactOpts, originChainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockHashReceiver *BlockHashReceiverTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashReceiver.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockHashReceiver *BlockHashReceiverSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.RenounceOwnership(&_BlockHashReceiver.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockHashReceiver *BlockHashReceiverTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.RenounceOwnership(&_BlockHashReceiver.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BlockHashReceiver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockHashReceiver *BlockHashReceiverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.TransferOwnership(&_BlockHashReceiver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.TransferOwnership(&_BlockHashReceiver.TransactOpts, newOwner)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0xb6109d9d.
//
// Solidity: function updateMinTeleporterVersion() returns()
func (_BlockHashReceiver *BlockHashReceiverTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashReceiver.contract.Transact(opts, "updateMinTeleporterVersion")
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0xb6109d9d.
//
// Solidity: function updateMinTeleporterVersion() returns()
func (_BlockHashReceiver *BlockHashReceiverSession) UpdateMinTeleporterVersion() (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.UpdateMinTeleporterVersion(&_BlockHashReceiver.TransactOpts)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0xb6109d9d.
//
// Solidity: function updateMinTeleporterVersion() returns()
func (_BlockHashReceiver *BlockHashReceiverTransactorSession) UpdateMinTeleporterVersion() (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.UpdateMinTeleporterVersion(&_BlockHashReceiver.TransactOpts)
}

// BlockHashReceiverMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the BlockHashReceiver contract.
type BlockHashReceiverMinTeleporterVersionUpdatedIterator struct {
	Event *BlockHashReceiverMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *BlockHashReceiverMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockHashReceiverMinTeleporterVersionUpdated)
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
		it.Event = new(BlockHashReceiverMinTeleporterVersionUpdated)
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
func (it *BlockHashReceiverMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockHashReceiverMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockHashReceiverMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the BlockHashReceiver contract.
type BlockHashReceiverMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_BlockHashReceiver *BlockHashReceiverFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*BlockHashReceiverMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverMinTeleporterVersionUpdatedIterator{contract: _BlockHashReceiver.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_BlockHashReceiver *BlockHashReceiverFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *BlockHashReceiverMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockHashReceiverMinTeleporterVersionUpdated)
				if err := _BlockHashReceiver.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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

// ParseMinTeleporterVersionUpdated is a log parse operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_BlockHashReceiver *BlockHashReceiverFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*BlockHashReceiverMinTeleporterVersionUpdated, error) {
	event := new(BlockHashReceiverMinTeleporterVersionUpdated)
	if err := _BlockHashReceiver.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockHashReceiverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BlockHashReceiver contract.
type BlockHashReceiverOwnershipTransferredIterator struct {
	Event *BlockHashReceiverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlockHashReceiverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockHashReceiverOwnershipTransferred)
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
		it.Event = new(BlockHashReceiverOwnershipTransferred)
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
func (it *BlockHashReceiverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockHashReceiverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockHashReceiverOwnershipTransferred represents a OwnershipTransferred event raised by the BlockHashReceiver contract.
type BlockHashReceiverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockHashReceiver *BlockHashReceiverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlockHashReceiverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverOwnershipTransferredIterator{contract: _BlockHashReceiver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockHashReceiver *BlockHashReceiverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockHashReceiverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockHashReceiverOwnershipTransferred)
				if err := _BlockHashReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockHashReceiver *BlockHashReceiverFilterer) ParseOwnershipTransferred(log types.Log) (*BlockHashReceiverOwnershipTransferred, error) {
	event := new(BlockHashReceiverOwnershipTransferred)
	if err := _BlockHashReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockHashReceiverReceiveBlockHashIterator is returned from FilterReceiveBlockHash and is used to iterate over the raw logs and unpacked data for ReceiveBlockHash events raised by the BlockHashReceiver contract.
type BlockHashReceiverReceiveBlockHashIterator struct {
	Event *BlockHashReceiverReceiveBlockHash // Event containing the contract specifics and raw log

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
func (it *BlockHashReceiverReceiveBlockHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockHashReceiverReceiveBlockHash)
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
		it.Event = new(BlockHashReceiverReceiveBlockHash)
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
func (it *BlockHashReceiverReceiveBlockHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockHashReceiverReceiveBlockHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockHashReceiverReceiveBlockHash represents a ReceiveBlockHash event raised by the BlockHashReceiver contract.
type BlockHashReceiverReceiveBlockHash struct {
	OriginChainID       [32]byte
	OriginSenderAddress common.Address
	BlockHeight         *big.Int
	BlockHash           [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReceiveBlockHash is a free log retrieval operation binding the contract event 0x0bca78aa82d7575f42e4b4b2fe04765a4b2f3661786403788ce987e065ac590d.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, address indexed originSenderAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_BlockHashReceiver *BlockHashReceiverFilterer) FilterReceiveBlockHash(opts *bind.FilterOpts, originChainID [][32]byte, originSenderAddress []common.Address, blockHeight []*big.Int) (*BlockHashReceiverReceiveBlockHashIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.FilterLogs(opts, "ReceiveBlockHash", originChainIDRule, originSenderAddressRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverReceiveBlockHashIterator{contract: _BlockHashReceiver.contract, event: "ReceiveBlockHash", logs: logs, sub: sub}, nil
}

// WatchReceiveBlockHash is a free log subscription operation binding the contract event 0x0bca78aa82d7575f42e4b4b2fe04765a4b2f3661786403788ce987e065ac590d.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, address indexed originSenderAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_BlockHashReceiver *BlockHashReceiverFilterer) WatchReceiveBlockHash(opts *bind.WatchOpts, sink chan<- *BlockHashReceiverReceiveBlockHash, originChainID [][32]byte, originSenderAddress []common.Address, blockHeight []*big.Int) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.WatchLogs(opts, "ReceiveBlockHash", originChainIDRule, originSenderAddressRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockHashReceiverReceiveBlockHash)
				if err := _BlockHashReceiver.contract.UnpackLog(event, "ReceiveBlockHash", log); err != nil {
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

// ParseReceiveBlockHash is a log parse operation binding the contract event 0x0bca78aa82d7575f42e4b4b2fe04765a4b2f3661786403788ce987e065ac590d.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed originChainID, address indexed originSenderAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_BlockHashReceiver *BlockHashReceiverFilterer) ParseReceiveBlockHash(log types.Log) (*BlockHashReceiverReceiveBlockHash, error) {
	event := new(BlockHashReceiverReceiveBlockHash)
	if err := _BlockHashReceiver.contract.UnpackLog(event, "ReceiveBlockHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
