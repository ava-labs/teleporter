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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"originBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ReceiveMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"SendMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getCurrentMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600180553480156200001557600080fd5b5060405162001b3138038062001b3183398101604081905262000038916200019a565b600160005580806001600160a01b038116620000c05760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f727465722072656769737472792061646472657373000000000000000000606482015260840160405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa1580156200010b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001319190620001cc565b60035550620001403362000148565b5050620001e6565b600480546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600060208284031215620001ad57600080fd5b81516001600160a01b0381168114620001c557600080fd5b9392505050565b600060208284031215620001df57600080fd5b5051919050565b60805161191b620002166000396000818160be01528181610566015281816109650152610cc3015261191b6000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806397314297116100715780639731429714610151578063b33fead41461018d578063c868efaa146101ae578063d2cc7a70146101c1578063f2fde38b146101d3578063f63d09d7146101e657600080fd5b80631a7f5bec146100b95780632b0d8f18146100fd5780634511243e146101125780635eb9951414610125578063715018a6146101385780638da5cb5b14610140575b600080fd5b6100e07f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b61011061010b3660046112cb565b6101f9565b005b6101106101203660046112cb565b6102fb565b6101106101333660046112ef565b6103f8565b61011061040c565b6004546001600160a01b03166100e0565b61017d61015f3660046112cb565b6001600160a01b031660009081526002602052604090205460ff1690565b60405190151581526020016100f4565b6101a061019b3660046112ef565b610420565b6040516100f4929190611358565b6101106101bc3660046113c5565b6104f8565b6003545b6040519081526020016100f4565b6101106101e13660046112cb565b61071d565b6101c56101f4366004611421565b610793565b610201610959565b6001600160a01b0381166102305760405162461bcd60e51b8152600401610227906114a7565b60405180910390fd5b6001600160a01b03811660009081526002602052604090205460ff16156102af5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610227565b6001600160a01b038116600081815260026020526040808220805460ff19166001179055517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b610303610959565b6001600160a01b0381166103295760405162461bcd60e51b8152600401610227906114a7565b6001600160a01b03811660009081526002602052604090205460ff166103a35760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610227565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600260205260409020805460ff19169055565b610400610959565b61040981610961565b50565b610414610b01565b61041e6000610b5b565b565b6000818152600560209081526040808320815180830190925280546001600160a01b031682526001810180546060948694939290840191610460906114f5565b80601f016020809104026020016040519081016040528092919081815260200182805461048c906114f5565b80156104d95780601f106104ae576101008083540402835291602001916104d9565b820191906000526020600020905b8154815290600101906020018083116104bc57829003601f168201915b5050505050815250509050806000015181602001519250925050915091565b600180541461055c5760405162461bcd60e51b815260206004820152602a60248201527f54656c65706f727465725570677261646561626c653a207265636569766572206044820152697265656e7472616e637960b01b6064820152608401610227565b60026001556003547f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa1580156105d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105fc919061152f565b10156106635760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610227565b61066c3361015f565b156106d25760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610227565b610713848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610bad92505050565b5050600180555050565b610725610b01565b6001600160a01b03811661078a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610227565b61040981610b5b565b600061079d610c65565b60006107a7610cbe565b9050600086156107d1576107bb8888610dd2565b90506107d16001600160a01b0389168383610f3c565b886001600160a01b03168a7fa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca8a848a8a8a604051610813959493929190611571565b60405180910390a3816001600160a01b031663624488506040518060c001604052808d81526020018c6001600160a01b0316815260200160405180604001604052808d6001600160a01b03168152602001868152508152602001898152602001600067ffffffffffffffff81111561088d5761088d61159f565b6040519080825280602002602001820160405280156108b6578160200160208202803683370190505b50815260200188886040516020016108cf9291906115b5565b6040516020818303038152906040528152506040518263ffffffff1660e01b81526004016108fd919061160d565b6020604051808303816000875af115801561091c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610940919061152f565b9250505061094e6001600055565b979650505050505050565b61041e610b01565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e5919061152f565b60035490915081831115610a555760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610227565b808311610aca5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610227565b6003839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6004546001600160a01b0316331461041e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610227565b600480546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600081806020019051810190610bc3919061168b565b6040805180820182526001600160a01b038681168252602080830185815260008a81526005909252939020825181546001600160a01b03191692169190911781559151929350916001820190610c19908261177b565b50905050826001600160a01b0316847f1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa83604051610c57919061183b565b60405180910390a350505050565b600260005403610cb75760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610227565b6002600055565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d43919061184e565b9050610d67816001600160a01b031660009081526002602052604090205460ff1690565b15610dcd5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610227565b919050565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015610e1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3f919061152f565b9050610e566001600160a01b038516333086611027565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610e9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec1919061152f565b9050818111610f275760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610227565b610f318282611881565b925050505b92915050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015610f8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fb1919061152f565b610fbb9190611894565b6040516001600160a01b03851660248201526044810182905290915061102190859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261105f565b50505050565b6040516001600160a01b03808516602483015283166044820152606481018290526110219085906323b872dd60e01b90608401610fea565b60006110b4826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166111369092919063ffffffff16565b80519091501561113157808060200190518101906110d291906118a7565b6111315760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610227565b505050565b6060611145848460008561114d565b949350505050565b6060824710156111ae5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610227565b600080866001600160a01b031685876040516111ca91906118c9565b60006040518083038185875af1925050503d8060008114611207576040519150601f19603f3d011682016040523d82523d6000602084013e61120c565b606091505b509150915061094e8783838760608315611287578251600003611280576001600160a01b0385163b6112805760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610227565b5081611145565b611145838381511561129c5781518083602001fd5b8060405162461bcd60e51b8152600401610227919061183b565b6001600160a01b038116811461040957600080fd5b6000602082840312156112dd57600080fd5b81356112e8816112b6565b9392505050565b60006020828403121561130157600080fd5b5035919050565b60005b8381101561132357818101518382015260200161130b565b50506000910152565b60008151808452611344816020860160208601611308565b601f01601f19169290920160200192915050565b6001600160a01b03831681526040602082018190526000906111459083018461132c565b60008083601f84011261138e57600080fd5b50813567ffffffffffffffff8111156113a657600080fd5b6020830191508360208285010111156113be57600080fd5b9250929050565b600080600080606085870312156113db57600080fd5b8435935060208501356113ed816112b6565b9250604085013567ffffffffffffffff81111561140957600080fd5b6114158782880161137c565b95989497509550505050565b600080600080600080600060c0888a03121561143c57600080fd5b87359650602088013561144e816112b6565b9550604088013561145e816112b6565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561148857600080fd5b6114948a828b0161137c565b989b979a50959850939692959293505050565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b600181811c9082168061150957607f821691505b60208210810361152957634e487b7160e01b600052602260045260246000fd5b50919050565b60006020828403121561154157600080fd5b5051919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60018060a01b038616815284602082015283604082015260806060820152600061094e608083018486611548565b634e487b7160e01b600052604160045260246000fd5b602081526000611145602083018486611548565b600081518084526020808501945080840160005b838110156116025781516001600160a01b0316875295820195908201906001016115dd565b509495945050505050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c084015261166e6101008401826115c9565b905060a0840151601f198483030160e0850152610f31828261132c565b60006020828403121561169d57600080fd5b815167ffffffffffffffff808211156116b557600080fd5b818401915084601f8301126116c957600080fd5b8151818111156116db576116db61159f565b604051601f8201601f19908116603f011681019083821181831017156117035761170361159f565b8160405282815287602084870101111561171c57600080fd5b61094e836020830160208801611308565b601f82111561113157600081815260208120601f850160051c810160208610156117545750805b601f850160051c820191505b8181101561177357828155600101611760565b505050505050565b815167ffffffffffffffff8111156117955761179561159f565b6117a9816117a384546114f5565b8461172d565b602080601f8311600181146117de57600084156117c65750858301515b600019600386901b1c1916600185901b178555611773565b600085815260208120601f198616915b8281101561180d578886015182559484019460019091019084016117ee565b508582101561182b5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6020815260006112e8602083018461132c565b60006020828403121561186057600080fd5b81516112e8816112b6565b634e487b7160e01b600052601160045260246000fd5b81810381811115610f3657610f3661186b565b80820180821115610f3657610f3661186b565b6000602082840312156118b957600080fd5b815180151581146112e857600080fd5b600082516118db818460208701611308565b919091019291505056fea26469706673582212201e0af58281e9935afc1410e77d47962a5108c590f10d8379aea4ec2ae231bad764736f6c63430008120033",
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
// Solidity: function sendMessage(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(bytes32)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactor) SendMessage(opts *bind.TransactOpts, destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.contract.Transact(opts, "sendMessage", destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(bytes32)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) SendMessage(destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.SendMessage(&_ExampleCrossChainMessenger.TransactOpts, destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(bytes32)
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

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.UnpauseTeleporterAddress(&_ExampleCrossChainMessenger.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ExampleCrossChainMessenger.Contract.UnpauseTeleporterAddress(&_ExampleCrossChainMessenger.TransactOpts, teleporterAddress)
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

// ExampleCrossChainMessengerTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerTeleporterAddressUnpausedIterator struct {
	Event *ExampleCrossChainMessengerTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *ExampleCrossChainMessengerTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleCrossChainMessengerTeleporterAddressUnpaused)
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
		it.Event = new(ExampleCrossChainMessengerTeleporterAddressUnpaused)
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
func (it *ExampleCrossChainMessengerTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleCrossChainMessengerTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleCrossChainMessengerTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the ExampleCrossChainMessenger contract.
type ExampleCrossChainMessengerTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*ExampleCrossChainMessengerTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &ExampleCrossChainMessengerTeleporterAddressUnpausedIterator{contract: _ExampleCrossChainMessenger.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *ExampleCrossChainMessengerTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ExampleCrossChainMessenger.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleCrossChainMessengerTeleporterAddressUnpaused)
				if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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
func (_ExampleCrossChainMessenger *ExampleCrossChainMessengerFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*ExampleCrossChainMessengerTeleporterAddressUnpaused, error) {
	event := new(ExampleCrossChainMessengerTeleporterAddressUnpaused)
	if err := _ExampleCrossChainMessenger.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
