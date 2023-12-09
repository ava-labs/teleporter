// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package examplecrosschainmessenger

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

// ExampleCrossChainMessengerMetaData contains all meta data concerning the ExampleCrossChainMessenger contract.
var ExampleCrossChainMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ReceiveMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"SendMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getCurrentMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200199138038062001991833981016040819052620000349162000196565b600160005580806001600160a01b038116620000bc5760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f727465722072656769737472792061646472657373000000000000000000606482015260840160405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa15801562000107573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200012d9190620001c8565b600255506200013c3362000144565b5050620001e2565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600060208284031215620001a957600080fd5b81516001600160a01b0381168114620001c157600080fd5b9392505050565b600060208284031215620001db57600080fd5b5051919050565b60805161177f620002126000396000818160b30152818161043b0152818161081d0152610b7b015261177f6000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806397314297116100715780639731429714610133578063b33fead41461016f578063c868efaa14610190578063d2cc7a70146101a3578063f2fde38b146101b5578063f63d09d7146101c857600080fd5b80631a7f5bec146100ae5780632b0d8f18146100f25780635eb9951414610107578063715018a61461011a5780638da5cb5b14610122575b600080fd5b6100d57f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b61010561010036600461117d565b6101db565b005b6101056101153660046111a1565b610323565b610105610337565b6003546001600160a01b03166100d5565b61015f61014136600461117d565b6001600160a01b031660009081526001602052604090205460ff1690565b60405190151581526020016100e9565b61018261017d3660046111a1565b61034b565b6040516100e992919061120a565b61010561019e366004611277565b610423565b6002545b6040519081526020016100e9565b6101056101c336600461117d565b6105d5565b6101a76101d63660046112d3565b61064b565b6101e3610811565b6001600160a01b0381166102555760405162461bcd60e51b815260206004820152602e60248201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560448201526d706f72746572206164647265737360901b60648201526084015b60405180910390fd5b6001600160a01b03811660009081526001602052604090205460ff16156102d45760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b606482015260840161024c565b6001600160a01b0381166000818152600160208190526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b61032b610811565b61033481610819565b50565b61033f6109b9565b6103496000610a13565b565b6000818152600460209081526040808320815180830190925280546001600160a01b03168252600181018054606094869493929084019161038b90611359565b80601f01602080910402602001604051908101604052809291908181526020018280546103b790611359565b80156104045780601f106103d957610100808354040283529160200191610404565b820191906000526020600020905b8154815290600101906020018083116103e757829003601f168201915b5050505050815250509050806000015181602001519250925050915091565b60025460405163260f846760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690634c1f08ce90602401602060405180830381865afa15801561048a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ae9190611393565b10156105155760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b606482015260840161024c565b3360009081526001602052604090205460ff161561058e5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b606482015260840161024c565b6105cf848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610a6592505050565b50505050565b6105dd6109b9565b6001600160a01b0381166106425760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161024c565b61033481610a13565b6000610655610b1d565b600061065f610b76565b905060008615610689576106738888610c8a565b90506106896001600160a01b0389168383610df4565b886001600160a01b03168a7fa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca8a848a8a8a6040516106cb9594939291906113d5565b60405180910390a3816001600160a01b031663624488506040518060c001604052808d81526020018c6001600160a01b0316815260200160405180604001604052808d6001600160a01b03168152602001868152508152602001898152602001600067ffffffffffffffff81111561074557610745611403565b60405190808252806020026020018201604052801561076e578160200160208202803683370190505b5081526020018888604051602001610787929190611419565b6040516020818303038152906040528152506040518263ffffffff1660e01b81526004016107b59190611471565b6020604051808303816000875af11580156107d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107f89190611393565b925050506108066001600055565b979650505050505050565b6103496109b9565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610879573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089d9190611393565b6002549091508183111561090d5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b606482015260840161024c565b8083116109825760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e00606482015260840161024c565b6002839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6003546001600160a01b031633146103495760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161024c565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600081806020019051810190610a7b91906114ef565b6040805180820182526001600160a01b038681168252602080830185815260008a81526004909252939020825181546001600160a01b03191692169190911781559151929350916001820190610ad190826115df565b50905050826001600160a01b0316847f1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa83604051610b0f919061169f565b60405180910390a350505050565b600260005403610b6f5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161024c565b6002600055565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bd7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bfb91906116b2565b9050610c1f816001600160a01b031660009081526001602052604090205460ff1690565b15610c855760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b606482015260840161024c565b919050565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015610cd3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf79190611393565b9050610d0e6001600160a01b038516333086610ed9565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610d55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d799190611393565b9050818111610ddf5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b606482015260840161024c565b610de982826116e5565b925050505b92915050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015610e45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e699190611393565b610e7391906116f8565b6040516001600160a01b0385166024820152604481018290529091506105cf90859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610f11565b6040516001600160a01b03808516602483015283166044820152606481018290526105cf9085906323b872dd60e01b90608401610ea2565b6000610f66826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610fe89092919063ffffffff16565b805190915015610fe35780806020019051810190610f84919061170b565b610fe35760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161024c565b505050565b6060610ff78484600085610fff565b949350505050565b6060824710156110605760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161024c565b600080866001600160a01b0316858760405161107c919061172d565b60006040518083038185875af1925050503d80600081146110b9576040519150601f19603f3d011682016040523d82523d6000602084013e6110be565b606091505b50915091506108068783838760608315611139578251600003611132576001600160a01b0385163b6111325760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161024c565b5081610ff7565b610ff7838381511561114e5781518083602001fd5b8060405162461bcd60e51b815260040161024c919061169f565b6001600160a01b038116811461033457600080fd5b60006020828403121561118f57600080fd5b813561119a81611168565b9392505050565b6000602082840312156111b357600080fd5b5035919050565b60005b838110156111d55781810151838201526020016111bd565b50506000910152565b600081518084526111f68160208601602086016111ba565b601f01601f19169290920160200192915050565b6001600160a01b0383168152604060208201819052600090610ff7908301846111de565b60008083601f84011261124057600080fd5b50813567ffffffffffffffff81111561125857600080fd5b60208301915083602082850101111561127057600080fd5b9250929050565b6000806000806060858703121561128d57600080fd5b84359350602085013561129f81611168565b9250604085013567ffffffffffffffff8111156112bb57600080fd5b6112c78782880161122e565b95989497509550505050565b600080600080600080600060c0888a0312156112ee57600080fd5b87359650602088013561130081611168565b9550604088013561131081611168565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561133a57600080fd5b6113468a828b0161122e565b989b979a50959850939692959293505050565b600181811c9082168061136d57607f821691505b60208210810361138d57634e487b7160e01b600052602260045260246000fd5b50919050565b6000602082840312156113a557600080fd5b5051919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60018060a01b03861681528460208201528360408201526080606082015260006108066080830184866113ac565b634e487b7160e01b600052604160045260246000fd5b602081526000610ff76020830184866113ac565b600081518084526020808501945080840160005b838110156114665781516001600160a01b031687529582019590820190600101611441565b509495945050505050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c08401526114d261010084018261142d565b905060a0840151601f198483030160e0850152610de982826111de565b60006020828403121561150157600080fd5b815167ffffffffffffffff8082111561151957600080fd5b818401915084601f83011261152d57600080fd5b81518181111561153f5761153f611403565b604051601f8201601f19908116603f0116810190838211818310171561156757611567611403565b8160405282815287602084870101111561158057600080fd5b6108068360208301602088016111ba565b601f821115610fe357600081815260208120601f850160051c810160208610156115b85750805b601f850160051c820191505b818110156115d7578281556001016115c4565b505050505050565b815167ffffffffffffffff8111156115f9576115f9611403565b61160d816116078454611359565b84611591565b602080601f831160018114611642576000841561162a5750858301515b600019600386901b1c1916600185901b1785556115d7565b600085815260208120601f198616915b8281101561167157888601518255948401946001909101908401611652565b508582101561168f5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60208152600061119a60208301846111de565b6000602082840312156116c457600080fd5b815161119a81611168565b634e487b7160e01b600052601160045260246000fd5b81810381811115610dee57610dee6116cf565b80820180821115610dee57610dee6116cf565b60006020828403121561171d57600080fd5b8151801515811461119a57600080fd5b6000825161173f8184602087016111ba565b919091019291505056fea26469706673582212206e426c632d35ceb528a4e22953725e0b1075340d6336aa6d29dadbc92a976ff664736f6c63430008120033",
}

// ExampleCrossChainMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleCrossChainMessengerMetaData.ABI instead.
var ExampleCrossChainMessengerABI = ExampleCrossChainMessengerMetaData.ABI

// ExampleCrossChainMessengerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleCrossChainMessengerMetaData.Bin instead.
var ExampleCrossChainMessengerBin = ExampleCrossChainMessengerMetaData.Bin

// DeployExampleCrossChainMessenger deploys a new Ethereum contract, binding an instance of ExampleCrossChainMessenger to it.
func DeployExampleCrossChainMessenger(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterRegistryAddress common.Address) (common.Address, *types.Transaction, *ExampleCrossChainMessenger, error) {
	parsed, err := ExampleCrossChainMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleCrossChainMessengerBin), backend, teleporterRegistryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExampleCrossChainMessenger{ExampleCrossChainMessengerCaller: ExampleCrossChainMessengerCaller{contract: contract}, ExampleCrossChainMessengerTransactor: ExampleCrossChainMessengerTransactor{contract: contract}, ExampleCrossChainMessengerFilterer: ExampleCrossChainMessengerFilterer{contract: contract}}, nil
}

// ExampleCrossChainMessenger is an auto generated Go binding around an Ethereum contract.
type ExampleCrossChainMessenger struct {
	ExampleCrossChainMessengerCaller     // Read-only binding to the contract
	ExampleCrossChainMessengerTransactor // Write-only binding to the contract
	ExampleCrossChainMessengerFilterer   // Log filterer for contract events
}

// ExampleCrossChainMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleCrossChainMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleCrossChainMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleCrossChainMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleCrossChainMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleCrossChainMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleCrossChainMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleCrossChainMessengerSession struct {
	Contract     *ExampleCrossChainMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExampleCrossChainMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleCrossChainMessengerCallerSession struct {
	Contract *ExampleCrossChainMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ExampleCrossChainMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleCrossChainMessengerTransactorSession struct {
	Contract     *ExampleCrossChainMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ExampleCrossChainMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleCrossChainMessengerRaw struct {
	Contract *ExampleCrossChainMessenger // Generic contract binding to access the raw methods on
}

// ExampleCrossChainMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleCrossChainMessengerCallerRaw struct {
	Contract *ExampleCrossChainMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleCrossChainMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleCrossChainMessengerTransactorRaw struct {
	Contract *ExampleCrossChainMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleCrossChainMessenger creates a new instance of ExampleCrossChainMessenger, bound to a specific deployed contract.
func NewExampleCrossChainMessenger(address common.Address, backend bind.ContractBackend) (*ExampleCrossChainMessenger, error) {
	contract, err := bindExampleCrossChainMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessenger{ExampleCrossChainMessengerCaller: ExampleCrossChainMessengerCaller{contract: contract}, ExampleCrossChainMessengerTransactor: ExampleCrossChainMessengerTransactor{contract: contract}, ExampleCrossChainMessengerFilterer: ExampleCrossChainMessengerFilterer{contract: contract}}, nil
}

// NewExampleCrossChainMessengerCaller creates a new read-only instance of ExampleCrossChainMessenger, bound to a specific deployed contract.
func NewExampleCrossChainMessengerCaller(address common.Address, caller bind.ContractCaller) (*ExampleCrossChainMessengerCaller, error) {
	contract, err := bindExampleCrossChainMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerCaller{contract: contract}, nil
}

// NewExampleCrossChainMessengerTransactor creates a new write-only instance of ExampleCrossChainMessenger, bound to a specific deployed contract.
func NewExampleCrossChainMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleCrossChainMessengerTransactor, error) {
	contract, err := bindExampleCrossChainMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerTransactor{contract: contract}, nil
}

// NewExampleCrossChainMessengerFilterer creates a new log filterer instance of ExampleCrossChainMessenger, bound to a specific deployed contract.
func NewExampleCrossChainMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleCrossChainMessengerFilterer, error) {
	contract, err := bindExampleCrossChainMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerFilterer{contract: contract}, nil
}

// bindExampleCrossChainMessenger binds a generic wrapper to an already deployed contract.
func bindExampleCrossChainMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExampleCrossChainMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleCrossChainMessenger.Contract.ExampleCrossChainMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.ExampleCrossChainMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.ExampleCrossChainMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleCrossChainMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 originBlockchainID) view returns(address, string)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCaller) GetCurrentMessage(opts *bind.CallOpts, originBlockchainID [32]byte) (common.Address, string, error) {
	var out []interface{}
	err := _ExampleCrossChainMessenger.contract.Call(opts, &out, "getCurrentMessage", originBlockchainID)

	if err != nil {
		return *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 originBlockchainID) view returns(address, string)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) GetCurrentMessage(originBlockchainID [32]byte) (common.Address, string, error) {
	return _ExampleCrossChainMessenger.Contract.GetCurrentMessage(&_ExampleCrossChainMessenger.CallOpts, originBlockchainID)
}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 originBlockchainID) view returns(address, string)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCallerSession) GetCurrentMessage(originBlockchainID [32]byte) (common.Address, string, error) {
	return _ExampleCrossChainMessenger.Contract.GetCurrentMessage(&_ExampleCrossChainMessenger.CallOpts, originBlockchainID)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExampleCrossChainMessenger.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ExampleCrossChainMessenger.Contract.GetMinTeleporterVersion(&_ExampleCrossChainMessenger.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ExampleCrossChainMessenger.Contract.GetMinTeleporterVersion(&_ExampleCrossChainMessenger.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _ExampleCrossChainMessenger.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _ExampleCrossChainMessenger.Contract.IsTeleporterAddressPaused(&_ExampleCrossChainMessenger.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _ExampleCrossChainMessenger.Contract.IsTeleporterAddressPaused(&_ExampleCrossChainMessenger.CallOpts, teleporterAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExampleCrossChainMessenger.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) Owner() (common.Address, error) {
	return _ExampleCrossChainMessenger.Contract.Owner(&_ExampleCrossChainMessenger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCallerSession) Owner() (common.Address, error) {
	return _ExampleCrossChainMessenger.Contract.Owner(&_ExampleCrossChainMessenger.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExampleCrossChainMessenger.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) TeleporterRegistry() (common.Address, error) {
	return _ExampleCrossChainMessenger.Contract.TeleporterRegistry(&_ExampleCrossChainMessenger.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerCallerSession) TeleporterRegistry() (common.Address, error) {
	return _ExampleCrossChainMessenger.Contract.TeleporterRegistry(&_ExampleCrossChainMessenger.CallOpts)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.PauseTeleporterAddress(&_ExampleCrossChainMessenger.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.PauseTeleporterAddress(&_ExampleCrossChainMessenger.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originBlockchainID, address originSenderAddress, bytes message) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, originBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.contract.Transact(opts, "receiveTeleporterMessage", originBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originBlockchainID, address originSenderAddress, bytes message) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) ReceiveTeleporterMessage(originBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.ReceiveTeleporterMessage(&_ExampleCrossChainMessenger.TransactOpts, originBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originBlockchainID, address originSenderAddress, bytes message) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorSession) ReceiveTeleporterMessage(originBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.ReceiveTeleporterMessage(&_ExampleCrossChainMessenger.TransactOpts, originBlockchainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.RenounceOwnership(&_ExampleCrossChainMessenger.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.RenounceOwnership(&_ExampleCrossChainMessenger.TransactOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(uint256)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactor) SendMessage(opts *bind.TransactOpts, destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.contract.Transact(opts, "sendMessage", destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(uint256)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) SendMessage(destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.SendMessage(&_ExampleCrossChainMessenger.TransactOpts, destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(uint256)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorSession) SendMessage(destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.SendMessage(&_ExampleCrossChainMessenger.TransactOpts, destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.TransferOwnership(&_ExampleCrossChainMessenger.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.TransferOwnership(&_ExampleCrossChainMessenger.TransactOpts, newOwner)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.UpdateMinTeleporterVersion(&_ExampleCrossChainMessenger.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.UpdateMinTeleporterVersion(&_ExampleCrossChainMessenger.TransactOpts, version)
}

// ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator struct {
	Event *ExampleCrossChainMessengerMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleCrossChainMessengerMinTeleporterVersionUpdated)
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
		it.Event = new(ExampleCrossChainMessengerMinTeleporterVersionUpdated)
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
func (it *ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleCrossChainMessengerMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerMinTeleporterVersionUpdatedIterator{contract: _ExampleCrossChainMessenger.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *ExampleCrossChainMessengerMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleCrossChainMessengerMinTeleporterVersionUpdated)
				if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*ExampleCrossChainMessengerMinTeleporterVersionUpdated, error) {
	event := new(ExampleCrossChainMessengerMinTeleporterVersionUpdated)
	if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleCrossChainMessengerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerOwnershipTransferredIterator struct {
	Event *ExampleCrossChainMessengerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExampleCrossChainMessengerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleCrossChainMessengerOwnershipTransferred)
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
		it.Event = new(ExampleCrossChainMessengerOwnershipTransferred)
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
func (it *ExampleCrossChainMessengerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleCrossChainMessengerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleCrossChainMessengerOwnershipTransferred represents a OwnershipTransferred event raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExampleCrossChainMessengerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerOwnershipTransferredIterator{contract: _ExampleCrossChainMessenger.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExampleCrossChainMessengerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleCrossChainMessengerOwnershipTransferred)
				if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) ParseOwnershipTransferred(log types.Log) (*ExampleCrossChainMessengerOwnershipTransferred, error) {
	event := new(ExampleCrossChainMessengerOwnershipTransferred)
	if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleCrossChainMessengerReceiveMessageIterator is returned from FilterReceiveMessage and is used to iterate over the raw logs and unpacked data for ReceiveMessage events raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerReceiveMessageIterator struct {
	Event *ExampleCrossChainMessengerReceiveMessage // Event containing the contract specifics and raw log

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
func (it *ExampleCrossChainMessengerReceiveMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleCrossChainMessengerReceiveMessage)
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
		it.Event = new(ExampleCrossChainMessengerReceiveMessage)
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
func (it *ExampleCrossChainMessengerReceiveMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleCrossChainMessengerReceiveMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleCrossChainMessengerReceiveMessage represents a ReceiveMessage event raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerReceiveMessage struct {
	OriginBlockchainID  [32]byte
	OriginSenderAddress common.Address
	Message             string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReceiveMessage is a free log retrieval operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed originBlockchainID, address indexed originSenderAddress, string message)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) FilterReceiveMessage(opts *bind.FilterOpts, originBlockchainID [][32]byte, originSenderAddress []common.Address) (*ExampleCrossChainMessengerReceiveMessageIterator, error) {

	var originBlockchainIDRule []interface{}
	for _, originBlockchainIDItem := range originBlockchainID {
		originBlockchainIDRule = append(originBlockchainIDRule, originBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.FilterLogs(opts, "ReceiveMessage", originBlockchainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerReceiveMessageIterator{contract: _ExampleCrossChainMessenger.contract, event: "ReceiveMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveMessage is a free log subscription operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed originBlockchainID, address indexed originSenderAddress, string message)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) WatchReceiveMessage(opts *bind.WatchOpts, sink chan<- *ExampleCrossChainMessengerReceiveMessage, originBlockchainID [][32]byte, originSenderAddress []common.Address) (event.Subscription, error) {

	var originBlockchainIDRule []interface{}
	for _, originBlockchainIDItem := range originBlockchainID {
		originBlockchainIDRule = append(originBlockchainIDRule, originBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.WatchLogs(opts, "ReceiveMessage", originBlockchainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleCrossChainMessengerReceiveMessage)
				if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "ReceiveMessage", log); err != nil {
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

// ParseReceiveMessage is a log parse operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed originBlockchainID, address indexed originSenderAddress, string message)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) ParseReceiveMessage(log types.Log) (*ExampleCrossChainMessengerReceiveMessage, error) {
	event := new(ExampleCrossChainMessengerReceiveMessage)
	if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "ReceiveMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleCrossChainMessengerSendMessageIterator is returned from FilterSendMessage and is used to iterate over the raw logs and unpacked data for SendMessage events raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerSendMessageIterator struct {
	Event *ExampleCrossChainMessengerSendMessage // Event containing the contract specifics and raw log

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
func (it *ExampleCrossChainMessengerSendMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleCrossChainMessengerSendMessage)
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
		it.Event = new(ExampleCrossChainMessengerSendMessage)
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
func (it *ExampleCrossChainMessengerSendMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleCrossChainMessengerSendMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleCrossChainMessengerSendMessage represents a SendMessage event raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerSendMessage struct {
	DestinationBlockchainID [32]byte
	DestinationAddress      common.Address
	FeeTokenAddress         common.Address
	FeeAmount               *big.Int
	RequiredGasLimit        *big.Int
	Message                 string
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSendMessage is a free log retrieval operation binding the contract event 0xa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca.
//
// Solidity: event SendMessage(bytes32 indexed destinationBlockchainID, address indexed destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) FilterSendMessage(opts *bind.FilterOpts, destinationBlockchainID [][32]byte, destinationAddress []common.Address) (*ExampleCrossChainMessengerSendMessageIterator, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.FilterLogs(opts, "SendMessage", destinationBlockchainIDRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerSendMessageIterator{contract: _ExampleCrossChainMessenger.contract, event: "SendMessage", logs: logs, sub: sub}, nil
}

// WatchSendMessage is a free log subscription operation binding the contract event 0xa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca.
//
// Solidity: event SendMessage(bytes32 indexed destinationBlockchainID, address indexed destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) WatchSendMessage(opts *bind.WatchOpts, sink chan<- *ExampleCrossChainMessengerSendMessage, destinationBlockchainID [][32]byte, destinationAddress []common.Address) (event.Subscription, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.WatchLogs(opts, "SendMessage", destinationBlockchainIDRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleCrossChainMessengerSendMessage)
				if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "SendMessage", log); err != nil {
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

// ParseSendMessage is a log parse operation binding the contract event 0xa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca.
//
// Solidity: event SendMessage(bytes32 indexed destinationBlockchainID, address indexed destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) ParseSendMessage(log types.Log) (*ExampleCrossChainMessengerSendMessage, error) {
	event := new(ExampleCrossChainMessengerSendMessage)
	if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "SendMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleCrossChainMessengerTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerTeleporterAddressPausedIterator struct {
	Event *ExampleCrossChainMessengerTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *ExampleCrossChainMessengerTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleCrossChainMessengerTeleporterAddressPaused)
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
		it.Event = new(ExampleCrossChainMessengerTeleporterAddressPaused)
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
func (it *ExampleCrossChainMessengerTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleCrossChainMessengerTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleCrossChainMessengerTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*ExampleCrossChainMessengerTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerTeleporterAddressPausedIterator{contract: _ExampleCrossChainMessenger.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *ExampleCrossChainMessengerTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleCrossChainMessengerTeleporterAddressPaused)
				if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) ParseTeleporterAddressPaused(log types.Log) (*ExampleCrossChainMessengerTeleporterAddressPaused, error) {
	event := new(ExampleCrossChainMessengerTeleporterAddressPaused)
	if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
