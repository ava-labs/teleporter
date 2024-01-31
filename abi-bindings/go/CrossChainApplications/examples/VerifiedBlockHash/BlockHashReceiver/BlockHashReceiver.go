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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"publisherBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"publisherContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"ReceiveBlockHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getLatestBlockInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourcePublisherContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b5060405162000f9438038062000f9483398101604081905261003191610223565b600160005582806001600160a01b0381166100b95760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa158015610103573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610127919061025f565b60025550610134336101b5565b506001600160a01b03811661019d5760405162461bcd60e51b815260206004820152602960248201527f426c6f636b4861736852656365697665723a207a65726f207075626c6973686560448201526872206164647265737360b81b60648201526084016100b0565b60a0919091526001600160a01b031660c05250610278565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80516001600160a01b038116811461021e57600080fd5b919050565b60008060006060848603121561023857600080fd5b61024184610207565b92506020840151915061025660408501610207565b90509250925092565b60006020828403121561027157600080fd5b5051919050565b60805160a05160c051610cd1620002c3600039600081816101c40152610a20015260008181610143015261099e01526000818160ff015281816104c401526106fb0152610cd16000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806379a0710c11610097578063c868efaa11610066578063c868efaa1461024e578063d2cc7a7014610261578063f2fde38b14610269578063f3f39ee51461027c57600080fd5b806379a0710c146101bf5780638da5cb5b146101e657806397314297146101f7578063b17810be1461023357600080fd5b80634511243e116100d35780634511243e146101885780635eb995141461019b5780636c4f6ba9146101ae578063715018a6146101b757600080fd5b80631a7f5bec146100fa57806329b7b3fd1461013e5780632b0d8f1814610173575b600080fd5b6101217f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6101657f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610135565b610186610181366004610b4e565b610285565b005b610186610196366004610b4e565b61038a565b6101866101a9366004610b70565b610487565b61016560055481565b61018661049b565b6101217f000000000000000000000000000000000000000000000000000000000000000081565b6003546001600160a01b0316610121565b610223610205366004610b4e565b6001600160a01b031660009081526001602052604090205460ff1690565b6040519015158152602001610135565b60045460055460408051928352602083019190915201610135565b61018661025c366004610b89565b6104af565b600254610165565b610186610277366004610b4e565b610679565b61016560045481565b61028d6106ef565b6001600160a01b0381166102bc5760405162461bcd60e51b81526004016102b390610c10565b60405180910390fd5b6001600160a01b03811660009081526001602052604090205460ff161561033b5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b60648201526084016102b3565b6001600160a01b0381166000818152600160208190526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b6103926106ef565b6001600160a01b0381166103b85760405162461bcd60e51b81526004016102b390610c10565b6001600160a01b03811660009081526001602052604090205460ff166104325760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b60648201526084016102b3565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600160205260409020805460ff19169055565b61048f6106ef565b610498816106f7565b50565b6104a3610897565b6104ad60006108f1565b565b6104b7610943565b6002546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa15801561052e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105529190610c5e565b10156105b95760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b60648201526084016102b3565b6105c233610205565b156106285760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b60648201526084016102b3565b610669848484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061099c92505050565b6106736001600055565b50505050565b610681610897565b6001600160a01b0381166106e65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102b3565b610498816108f1565b6104ad610897565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610757573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077b9190610c5e565b600254909150818311156107eb5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b60648201526084016102b3565b8083116108605760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e0060648201526084016102b3565b6002839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6003546001600160a01b031633146104ad5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b3565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6002600054036109955760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016102b3565b6002600055565b7f00000000000000000000000000000000000000000000000000000000000000008314610a1e5760405162461bcd60e51b815260206004820152602a60248201527f426c6f636b4861736852656365697665723a20696e76616c696420736f757263604482015269194818da185a5b88125160b21b60648201526084016102b3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b031614610ab95760405162461bcd60e51b815260206004820152603160248201527f426c6f636b4861736852656365697665723a20696e76616c696420736f757263604482015270329031b430b4b710383ab13634b9b432b960791b60648201526084016102b3565b60008082806020019051810190610ad09190610c77565b91509150600454821115610b2b576004829055600581905560405181815282906001600160a01b0386169087907f0bca78aa82d7575f42e4b4b2fe04765a4b2f3661786403788ce987e065ac590d9060200160405180910390a45b5050505050565b80356001600160a01b0381168114610b4957600080fd5b919050565b600060208284031215610b6057600080fd5b610b6982610b32565b9392505050565b600060208284031215610b8257600080fd5b5035919050565b60008060008060608587031215610b9f57600080fd5b84359350610baf60208601610b32565b9250604085013567ffffffffffffffff80821115610bcc57600080fd5b818701915087601f830112610be057600080fd5b813581811115610bef57600080fd5b886020828501011115610c0157600080fd5b95989497505060200194505050565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b600060208284031215610c7057600080fd5b5051919050565b60008060408385031215610c8a57600080fd5b50508051602090910151909290915056fea264697066735822122029c2c619026127977100acfed48415bb21c932d59c4bc5d029dd6022d946dfff64736f6c63430008120033",
}

// BlockHashReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockHashReceiverMetaData.ABI instead.
var BlockHashReceiverABI = BlockHashReceiverMetaData.ABI

// BlockHashReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlockHashReceiverMetaData.Bin instead.
var BlockHashReceiverBin = BlockHashReceiverMetaData.Bin

// DeployBlockHashReceiver deploys a new Ethereum contract, binding an instance of BlockHashReceiver to it.
func DeployBlockHashReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterRegistryAddress common.Address, publisherBlockchainID [32]byte, publisherContractAddress common.Address) (common.Address, *types.Transaction, *BlockHashReceiver, error) {
	parsed, err := BlockHashReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlockHashReceiverBin), backend, teleporterRegistryAddress, publisherBlockchainID, publisherContractAddress)
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
// Solidity: function getLatestBlockInfo() view returns(uint256, bytes32)
func (_BlockHashReceiver *BlockHashReceiverCaller) GetLatestBlockInfo(opts *bind.CallOpts) (*big.Int, [32]byte, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "getLatestBlockInfo")

	if err != nil {
		return *new(*big.Int), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetLatestBlockInfo is a free data retrieval call binding the contract method 0xb17810be.
//
// Solidity: function getLatestBlockInfo() view returns(uint256, bytes32)
func (_BlockHashReceiver *BlockHashReceiverSession) GetLatestBlockInfo() (*big.Int, [32]byte, error) {
	return _BlockHashReceiver.Contract.GetLatestBlockInfo(&_BlockHashReceiver.CallOpts)
}

// GetLatestBlockInfo is a free data retrieval call binding the contract method 0xb17810be.
//
// Solidity: function getLatestBlockInfo() view returns(uint256, bytes32)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) GetLatestBlockInfo() (*big.Int, [32]byte, error) {
	return _BlockHashReceiver.Contract.GetLatestBlockInfo(&_BlockHashReceiver.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _BlockHashReceiver.Contract.GetMinTeleporterVersion(&_BlockHashReceiver.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _BlockHashReceiver.Contract.GetMinTeleporterVersion(&_BlockHashReceiver.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_BlockHashReceiver *BlockHashReceiverCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_BlockHashReceiver *BlockHashReceiverSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _BlockHashReceiver.Contract.IsTeleporterAddressPaused(&_BlockHashReceiver.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _BlockHashReceiver.Contract.IsTeleporterAddressPaused(&_BlockHashReceiver.CallOpts, teleporterAddress)
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

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverCaller) SourceBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlockHashReceiver.contract.Call(opts, &out, "sourceBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverSession) SourceBlockchainID() ([32]byte, error) {
	return _BlockHashReceiver.Contract.SourceBlockchainID(&_BlockHashReceiver.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_BlockHashReceiver *BlockHashReceiverCallerSession) SourceBlockchainID() ([32]byte, error) {
	return _BlockHashReceiver.Contract.SourceBlockchainID(&_BlockHashReceiver.CallOpts)
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

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _BlockHashReceiver.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_BlockHashReceiver *BlockHashReceiverSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.PauseTeleporterAddress(&_BlockHashReceiver.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.PauseTeleporterAddress(&_BlockHashReceiver.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _BlockHashReceiver.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_BlockHashReceiver *BlockHashReceiverSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.ReceiveTeleporterMessage(&_BlockHashReceiver.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.ReceiveTeleporterMessage(&_BlockHashReceiver.TransactOpts, sourceBlockchainID, originSenderAddress, message)
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

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _BlockHashReceiver.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_BlockHashReceiver *BlockHashReceiverSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.UnpauseTeleporterAddress(&_BlockHashReceiver.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.UnpauseTeleporterAddress(&_BlockHashReceiver.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _BlockHashReceiver.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_BlockHashReceiver *BlockHashReceiverSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.UpdateMinTeleporterVersion(&_BlockHashReceiver.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_BlockHashReceiver *BlockHashReceiverTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _BlockHashReceiver.Contract.UpdateMinTeleporterVersion(&_BlockHashReceiver.TransactOpts, version)
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
	SourceBlockchainID  [32]byte
	OriginSenderAddress common.Address
	BlockHeight         *big.Int
	BlockHash           [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReceiveBlockHash is a free log retrieval operation binding the contract event 0x0bca78aa82d7575f42e4b4b2fe04765a4b2f3661786403788ce987e065ac590d.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_BlockHashReceiver *BlockHashReceiverFilterer) FilterReceiveBlockHash(opts *bind.FilterOpts, sourceBlockchainID [][32]byte, originSenderAddress []common.Address, blockHeight []*big.Int) (*BlockHashReceiverReceiveBlockHashIterator, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.FilterLogs(opts, "ReceiveBlockHash", sourceBlockchainIDRule, originSenderAddressRule, blockHeightRule)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverReceiveBlockHashIterator{contract: _BlockHashReceiver.contract, event: "ReceiveBlockHash", logs: logs, sub: sub}, nil
}

// WatchReceiveBlockHash is a free log subscription operation binding the contract event 0x0bca78aa82d7575f42e4b4b2fe04765a4b2f3661786403788ce987e065ac590d.
//
// Solidity: event ReceiveBlockHash(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_BlockHashReceiver *BlockHashReceiverFilterer) WatchReceiveBlockHash(opts *bind.WatchOpts, sink chan<- *BlockHashReceiverReceiveBlockHash, sourceBlockchainID [][32]byte, originSenderAddress []common.Address, blockHeight []*big.Int) (event.Subscription, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.WatchLogs(opts, "ReceiveBlockHash", sourceBlockchainIDRule, originSenderAddressRule, blockHeightRule)
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
// Solidity: event ReceiveBlockHash(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, uint256 indexed blockHeight, bytes32 blockHash)
func (_BlockHashReceiver *BlockHashReceiverFilterer) ParseReceiveBlockHash(log types.Log) (*BlockHashReceiverReceiveBlockHash, error) {
	event := new(BlockHashReceiverReceiveBlockHash)
	if err := _BlockHashReceiver.contract.UnpackLog(event, "ReceiveBlockHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockHashReceiverTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the BlockHashReceiver contract.
type BlockHashReceiverTeleporterAddressPausedIterator struct {
	Event *BlockHashReceiverTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *BlockHashReceiverTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockHashReceiverTeleporterAddressPaused)
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
		it.Event = new(BlockHashReceiverTeleporterAddressPaused)
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
func (it *BlockHashReceiverTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockHashReceiverTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockHashReceiverTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the BlockHashReceiver contract.
type BlockHashReceiverTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_BlockHashReceiver *BlockHashReceiverFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*BlockHashReceiverTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverTeleporterAddressPausedIterator{contract: _BlockHashReceiver.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_BlockHashReceiver *BlockHashReceiverFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *BlockHashReceiverTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockHashReceiverTeleporterAddressPaused)
				if err := _BlockHashReceiver.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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

// ParseTeleporterAddressPaused is a log parse operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_BlockHashReceiver *BlockHashReceiverFilterer) ParseTeleporterAddressPaused(log types.Log) (*BlockHashReceiverTeleporterAddressPaused, error) {
	event := new(BlockHashReceiverTeleporterAddressPaused)
	if err := _BlockHashReceiver.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockHashReceiverTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the BlockHashReceiver contract.
type BlockHashReceiverTeleporterAddressUnpausedIterator struct {
	Event *BlockHashReceiverTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *BlockHashReceiverTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockHashReceiverTeleporterAddressUnpaused)
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
		it.Event = new(BlockHashReceiverTeleporterAddressUnpaused)
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
func (it *BlockHashReceiverTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockHashReceiverTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockHashReceiverTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the BlockHashReceiver contract.
type BlockHashReceiverTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_BlockHashReceiver *BlockHashReceiverFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*BlockHashReceiverTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &BlockHashReceiverTeleporterAddressUnpausedIterator{contract: _BlockHashReceiver.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_BlockHashReceiver *BlockHashReceiverFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *BlockHashReceiverTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _BlockHashReceiver.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockHashReceiverTeleporterAddressUnpaused)
				if err := _BlockHashReceiver.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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

// ParseTeleporterAddressUnpaused is a log parse operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_BlockHashReceiver *BlockHashReceiverFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*BlockHashReceiverTeleporterAddressUnpaused, error) {
	event := new(BlockHashReceiverTeleporterAddressUnpaused)
	if err := _BlockHashReceiver.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
