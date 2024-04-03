// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20source

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

// SendTokensInput is an auto generated low-level Go binding around an user-defined struct.
type SendTokensInput struct {
	DestinationBlockchainID  [32]byte
	DestinationBridgeAddress common.Address
	Recipient                common.Address
	PrimaryFee               *big.Int
	SecondaryFee             *big.Int
	RequiredGasLimit         *big.Int
}

// ERC20SourceMetaData contains all meta data concerning the ERC20Source contract.
var ERC20SourceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SendTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"}],\"name\":\"bridgedBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162002197380380620021978339810160408190526200003591620003a5565b60016000558282828282816001600160a01b038116620000c25760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001339190620003ef565b60025550620001423362000259565b6200014d81620002ab565b50507302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001a2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001c89190620003ef565b60a0526001600160a01b038116620002395760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401620000b9565b6001600160a01b0390811660c0529290921660e052506200040992505050565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b620002b56200032a565b6001600160a01b0381166200031c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401620000b9565b620003278162000259565b50565b6003546001600160a01b03163314620003865760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401620000b9565b565b80516001600160a01b0381168114620003a057600080fd5b919050565b600080600060608486031215620003bb57600080fd5b620003c68462000388565b9250620003d66020850162000388565b9150620003e66040850162000388565b90509250925092565b6000602082840312156200040257600080fd5b5051919050565b60805160a05160c05160e051611d1c6200047b6000396000818161029b0152818161109c01526111fa01526000818161021f0152610b17015260008181610259015281816108280152610f9501526000818161013d0152818161052c01528181610c72015261138e0152611d1c6000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80638da5cb5b11610097578063d127dc9b11610066578063d127dc9b14610254578063d2cc7a701461027b578063f2fde38b14610283578063fc0c546a1461029657600080fd5b80638da5cb5b146101cd57806397314297146101de578063b8df0dea1461021a578063c868efaa1461024157600080fd5b806332709bc1116100d357806332709bc11461018c5780634511243e1461019f5780635eb99514146101b2578063715018a6146101c557600080fd5b806302ee3e9c146100fa5780631a7f5bec146101385780632b0d8f1814610177575b600080fd5b610125610108366004611866565b600460209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b61015f7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161012f565b61018a610185366004611896565b6102bd565b005b61018a61019a3660046118b3565b6103c2565b61018a6101ad366004611896565b6103f2565b61018a6101c03660046118e4565b6104ef565b61018a610503565b6003546001600160a01b031661015f565b61020a6101ec366004611896565b6001600160a01b031660009081526001602052604090205460ff1690565b604051901515815260200161012f565b61015f7f000000000000000000000000000000000000000000000000000000000000000081565b61018a61024f3660046118fd565b610517565b6101257f000000000000000000000000000000000000000000000000000000000000000081565b600254610125565b61018a610291366004611896565b6106e1565b61015f7f000000000000000000000000000000000000000000000000000000000000000081565b6102c5610757565b6001600160a01b0381166102f45760405162461bcd60e51b81526004016102eb90611986565b60405180910390fd5b6001600160a01b03811660009081526001602052604090205460ff16156103735760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b60648201526084016102eb565b6001600160a01b0381166000818152600160208190526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b6103ca61075f565b6103e46103dc36849003840184611a21565b8260006107b8565b6103ee6001600055565b5050565b6103fa610757565b6001600160a01b0381166104205760405162461bcd60e51b81526004016102eb90611986565b6001600160a01b03811660009081526001602052604090205460ff1661049a5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b60648201526084016102eb565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600160205260409020805460ff19169055565b6104f7610757565b61050081610c6e565b50565b61050b610e0e565b6105156000610e68565b565b61051f61075f565b6002546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015610596573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ba9190611a8f565b10156106215760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b60648201526084016102eb565b61062a336101ec565b156106905760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b60648201526084016102eb565b6106d1848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610eba92505050565b6106db6001600055565b50505050565b6106e9610e0e565b6001600160a01b03811661074e5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102eb565b61050081610e68565b610515610e0e565b6002600054036107b15760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016102eb565b6002600055565b82516108245760405162461bcd60e51b815260206004820152603560248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f20646573746044820152741a5b985d1a5bdb88189b1bd8dad8da185a5b881251605a1b60648201526084016102eb565b82517f000000000000000000000000000000000000000000000000000000000000000090036108b05760405162461bcd60e51b815260206004820152603260248201527f54656c65706f72746572546f6b656e536f757263653a2063616e6e6f7420627260448201527134b233b2903a379039b0b6b29031b430b4b760711b60648201526084016102eb565b60208301516001600160a01b03166109295760405162461bcd60e51b815260206004820152603660248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f2064657374604482015275696e6174696f6e20627269646765206164647265737360501b60648201526084016102eb565b60408301516001600160a01b03166109995760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f207265636960448201526c7069656e74206164647265737360981b60648201526084016102eb565b608083015115610a015760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572546f6b656e536f757263653a206e6f6e2d7a65726f2060448201526c7365636f6e646172792066656560981b60648201526084016102eb565b80610a1257610a0f82611095565b91505b82606001518211610a8b5760405162461bcd60e51b815260206004820152603860248201527f54656c65706f72746572546f6b656e536f757263653a20696e7375666669636960448201527f656e7420616d6f756e7420746f20636f7665722066656573000000000000000060648201526084016102eb565b6060830151610a9a9083611abe565b83516000908152600460209081526040808320828801516001600160a01b03168452909152812080549294508492909190610ad6908490611ad1565b925050819055506000610be66040518060c001604052808660000151815260200186602001516001600160a01b0316815260200160405180604001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602001886060015181525081526020018660a001518152602001600067ffffffffffffffff811115610b7557610b756119d4565b604051908082528060200260200182016040528015610b9e578160200160208202803683370190505b508152602001866040015186604051602001610bcf9291906001600160a01b03929092168252602082015260400190565b6040516020818303038152906040528152506110c7565b60408051865181526020808801516001600160a01b0390811691830191909152878301511681830152606080880151908201526080808801519082015260a0808801519082015260c081018690529051919250339183917fdc72eade375f64388998c651c5137da74fbe64aec18f766dea3da18f154ab0a0919081900360e00190a350505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf29190611a8f565b60025490915081831115610d625760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b60648201526084016102eb565b808311610dd75760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e0060648201526084016102eb565b6002839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6003546001600160a01b031633146105155760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102eb565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60008082806020019051810190610ed19190611ae4565b60008781526004602090815260408083206001600160a01b038a168452909152902054919350915081811015610f645760405162461bcd60e51b815260206004820152603260248201527f54656c65706f72746572546f6b656e536f757263653a20696e73756666696369604482015271656e74206272696467652062616c616e636560701b60648201526084016102eb565b610f6e8282611abe565b60008781526004602090815260408083206001600160a01b038a16845290915290205582517f000000000000000000000000000000000000000000000000000000000000000090036110895760208301516001600160a01b0316301461102c5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572546f6b656e536f757263653a20696e76616c6964206260448201526c7269646765206164647265737360981b60648201526084016102eb565b82604001516001600160a01b03167f680f2e4f4032ebf1774e8cdbaddcb1b617a5a606411c8ca96257ada338d3833c8360405161106b91815260200190565b60405180910390a26110818360400151836111ed565b505050505050565b611081838360016107b8565b60006110c17f000000000000000000000000000000000000000000000000000000000000000083611221565b92915050565b6000806110d2611389565b60408401516020015190915015611177576040830151516001600160a01b03166111545760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b60648201526084016102eb565b604083015160208101519051611177916001600160a01b0390911690839061149d565b604051630624488560e41b81526001600160a01b038216906362448850906111a3908690600401611bfa565b6020604051808303816000875af11580156111c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111e69190611a8f565b9392505050565b6103ee6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168383611582565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa15801561126a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061128e9190611a8f565b90506112a56001600160a01b0385163330866115b7565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa1580156112ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113109190611a8f565b90508181116113765760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016102eb565b6113808282611abe565b95945050505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061140e9190611c78565b9050611432816001600160a01b031660009081526001602052604090205460ff1690565b156114985760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b60648201526084016102eb565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa1580156114ee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115129190611a8f565b61151c9190611ad1565b6040516001600160a01b0385166024820152604481018290529091506106db90859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526115ef565b6040516001600160a01b0383166024820152604481018290526115b290849063a9059cbb60e01b9060640161154b565b505050565b6040516001600160a01b03808516602483015283166044820152606481018290526106db9085906323b872dd60e01b9060840161154b565b6000611644826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166116c19092919063ffffffff16565b8051909150156115b257808060200190518101906116629190611c95565b6115b25760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016102eb565b60606116d084846000856116d8565b949350505050565b6060824710156117395760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016102eb565b600080866001600160a01b031685876040516117559190611cb7565b60006040518083038185875af1925050503d8060008114611792576040519150601f19603f3d011682016040523d82523d6000602084013e611797565b606091505b50915091506117a8878383876117b3565b979650505050505050565b6060831561182257825160000361181b576001600160a01b0385163b61181b5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102eb565b50816116d0565b6116d083838151156118375781518083602001fd5b8060405162461bcd60e51b81526004016102eb9190611cd3565b6001600160a01b038116811461050057600080fd5b6000806040838503121561187957600080fd5b82359150602083013561188b81611851565b809150509250929050565b6000602082840312156118a857600080fd5b81356111e681611851565b60008082840360e08112156118c757600080fd5b60c08112156118d557600080fd5b50919360c08501359350915050565b6000602082840312156118f657600080fd5b5035919050565b6000806000806060858703121561191357600080fd5b84359350602085013561192581611851565b9250604085013567ffffffffffffffff8082111561194257600080fd5b818701915087601f83011261195657600080fd5b81358181111561196557600080fd5b88602082850101111561197757600080fd5b95989497505060200194505050565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b60405160c0810167ffffffffffffffff81118282101715611a1b57634e487b7160e01b600052604160045260246000fd5b60405290565b600060c08284031215611a3357600080fd5b611a3b6119ea565b823581526020830135611a4d81611851565b60208201526040830135611a6081611851565b80604083015250606083013560608201526080830135608082015260a083013560a08201528091505092915050565b600060208284031215611aa157600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b818103818111156110c1576110c1611aa8565b808201808211156110c1576110c1611aa8565b60008082840360e0811215611af857600080fd5b60c0811215611b0657600080fd5b50611b0f6119ea565b835181526020840151611b2181611851565b60208201526040840151611b3481611851565b6040820152606084810151908201526080808501519082015260a0808501519082015260c09093015192949293505050565b600081518084526020808501945080840160005b83811015611b9f5781516001600160a01b031687529582019590820190600101611b7a565b509495945050505050565b60005b83811015611bc5578181015183820152602001611bad565b50506000910152565b60008151808452611be6816020860160208601611baa565b601f01601f19169290920160200192915050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c0840152611c5b610100840182611b66565b905060a0840151601f198483030160e08501526113808282611bce565b600060208284031215611c8a57600080fd5b81516111e681611851565b600060208284031215611ca757600080fd5b815180151581146111e657600080fd5b60008251611cc9818460208701611baa565b9190910192915050565b6020815260006111e66020830184611bce56fea2646970667358221220dc16a6759a0a1a4fe3d7d0b11109a294b9fa8664196e767645d5eee1d809e30564736f6c63430008120033",
}

// ERC20SourceABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20SourceMetaData.ABI instead.
var ERC20SourceABI = ERC20SourceMetaData.ABI

// ERC20SourceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20SourceMetaData.Bin instead.
var ERC20SourceBin = ERC20SourceMetaData.Bin

// DeployERC20Source deploys a new Ethereum contract, binding an instance of ERC20Source to it.
func DeployERC20Source(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterRegistryAddress common.Address, teleporterManager common.Address, tokenAddress common.Address) (common.Address, *types.Transaction, *ERC20Source, error) {
	parsed, err := ERC20SourceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20SourceBin), backend, teleporterRegistryAddress, teleporterManager, tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Source{ERC20SourceCaller: ERC20SourceCaller{contract: contract}, ERC20SourceTransactor: ERC20SourceTransactor{contract: contract}, ERC20SourceFilterer: ERC20SourceFilterer{contract: contract}}, nil
}

// ERC20Source is an auto generated Go binding around an Ethereum contract.
type ERC20Source struct {
	ERC20SourceCaller     // Read-only binding to the contract
	ERC20SourceTransactor // Write-only binding to the contract
	ERC20SourceFilterer   // Log filterer for contract events
}

// ERC20SourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20SourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20SourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20SourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20SourceSession struct {
	Contract     *ERC20Source      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20SourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20SourceCallerSession struct {
	Contract *ERC20SourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20SourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20SourceTransactorSession struct {
	Contract     *ERC20SourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20SourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20SourceRaw struct {
	Contract *ERC20Source // Generic contract binding to access the raw methods on
}

// ERC20SourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20SourceCallerRaw struct {
	Contract *ERC20SourceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20SourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20SourceTransactorRaw struct {
	Contract *ERC20SourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Source creates a new instance of ERC20Source, bound to a specific deployed contract.
func NewERC20Source(address common.Address, backend bind.ContractBackend) (*ERC20Source, error) {
	contract, err := bindERC20Source(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Source{ERC20SourceCaller: ERC20SourceCaller{contract: contract}, ERC20SourceTransactor: ERC20SourceTransactor{contract: contract}, ERC20SourceFilterer: ERC20SourceFilterer{contract: contract}}, nil
}

// NewERC20SourceCaller creates a new read-only instance of ERC20Source, bound to a specific deployed contract.
func NewERC20SourceCaller(address common.Address, caller bind.ContractCaller) (*ERC20SourceCaller, error) {
	contract, err := bindERC20Source(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20SourceCaller{contract: contract}, nil
}

// NewERC20SourceTransactor creates a new write-only instance of ERC20Source, bound to a specific deployed contract.
func NewERC20SourceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20SourceTransactor, error) {
	contract, err := bindERC20Source(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20SourceTransactor{contract: contract}, nil
}

// NewERC20SourceFilterer creates a new log filterer instance of ERC20Source, bound to a specific deployed contract.
func NewERC20SourceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20SourceFilterer, error) {
	contract, err := bindERC20Source(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20SourceFilterer{contract: contract}, nil
}

// bindERC20Source binds a generic wrapper to an already deployed contract.
func bindERC20Source(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20SourceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Source *ERC20SourceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Source.Contract.ERC20SourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Source *ERC20SourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Source.Contract.ERC20SourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Source *ERC20SourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Source.Contract.ERC20SourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Source *ERC20SourceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Source.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Source *ERC20SourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Source.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Source *ERC20SourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Source.Contract.contract.Transact(opts, method, params...)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_ERC20Source *ERC20SourceCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Source.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_ERC20Source *ERC20SourceSession) BlockchainID() ([32]byte, error) {
	return _ERC20Source.Contract.BlockchainID(&_ERC20Source.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_ERC20Source *ERC20SourceCallerSession) BlockchainID() ([32]byte, error) {
	return _ERC20Source.Contract.BlockchainID(&_ERC20Source.CallOpts)
}

// BridgedBalances is a free data retrieval call binding the contract method 0x02ee3e9c.
//
// Solidity: function bridgedBalances(bytes32 destinationBlockchainID, address destinationBridgeAddress) view returns(uint256 balance)
func (_ERC20Source *ERC20SourceCaller) BridgedBalances(opts *bind.CallOpts, destinationBlockchainID [32]byte, destinationBridgeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Source.contract.Call(opts, &out, "bridgedBalances", destinationBlockchainID, destinationBridgeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgedBalances is a free data retrieval call binding the contract method 0x02ee3e9c.
//
// Solidity: function bridgedBalances(bytes32 destinationBlockchainID, address destinationBridgeAddress) view returns(uint256 balance)
func (_ERC20Source *ERC20SourceSession) BridgedBalances(destinationBlockchainID [32]byte, destinationBridgeAddress common.Address) (*big.Int, error) {
	return _ERC20Source.Contract.BridgedBalances(&_ERC20Source.CallOpts, destinationBlockchainID, destinationBridgeAddress)
}

// BridgedBalances is a free data retrieval call binding the contract method 0x02ee3e9c.
//
// Solidity: function bridgedBalances(bytes32 destinationBlockchainID, address destinationBridgeAddress) view returns(uint256 balance)
func (_ERC20Source *ERC20SourceCallerSession) BridgedBalances(destinationBlockchainID [32]byte, destinationBridgeAddress common.Address) (*big.Int, error) {
	return _ERC20Source.Contract.BridgedBalances(&_ERC20Source.CallOpts, destinationBlockchainID, destinationBridgeAddress)
}

// FeeTokenAddress is a free data retrieval call binding the contract method 0xb8df0dea.
//
// Solidity: function feeTokenAddress() view returns(address)
func (_ERC20Source *ERC20SourceCaller) FeeTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Source.contract.Call(opts, &out, "feeTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenAddress is a free data retrieval call binding the contract method 0xb8df0dea.
//
// Solidity: function feeTokenAddress() view returns(address)
func (_ERC20Source *ERC20SourceSession) FeeTokenAddress() (common.Address, error) {
	return _ERC20Source.Contract.FeeTokenAddress(&_ERC20Source.CallOpts)
}

// FeeTokenAddress is a free data retrieval call binding the contract method 0xb8df0dea.
//
// Solidity: function feeTokenAddress() view returns(address)
func (_ERC20Source *ERC20SourceCallerSession) FeeTokenAddress() (common.Address, error) {
	return _ERC20Source.Contract.FeeTokenAddress(&_ERC20Source.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20Source *ERC20SourceCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Source.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20Source *ERC20SourceSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ERC20Source.Contract.GetMinTeleporterVersion(&_ERC20Source.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20Source *ERC20SourceCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ERC20Source.Contract.GetMinTeleporterVersion(&_ERC20Source.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20Source *ERC20SourceCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Source.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20Source *ERC20SourceSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _ERC20Source.Contract.IsTeleporterAddressPaused(&_ERC20Source.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20Source *ERC20SourceCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _ERC20Source.Contract.IsTeleporterAddressPaused(&_ERC20Source.CallOpts, teleporterAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Source *ERC20SourceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Source.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Source *ERC20SourceSession) Owner() (common.Address, error) {
	return _ERC20Source.Contract.Owner(&_ERC20Source.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Source *ERC20SourceCallerSession) Owner() (common.Address, error) {
	return _ERC20Source.Contract.Owner(&_ERC20Source.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ERC20Source *ERC20SourceCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Source.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ERC20Source *ERC20SourceSession) TeleporterRegistry() (common.Address, error) {
	return _ERC20Source.Contract.TeleporterRegistry(&_ERC20Source.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ERC20Source *ERC20SourceCallerSession) TeleporterRegistry() (common.Address, error) {
	return _ERC20Source.Contract.TeleporterRegistry(&_ERC20Source.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ERC20Source *ERC20SourceCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Source.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ERC20Source *ERC20SourceSession) Token() (common.Address, error) {
	return _ERC20Source.Contract.Token(&_ERC20Source.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ERC20Source *ERC20SourceCallerSession) Token() (common.Address, error) {
	return _ERC20Source.Contract.Token(&_ERC20Source.CallOpts)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20Source *ERC20SourceTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20Source.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20Source *ERC20SourceSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20Source.Contract.PauseTeleporterAddress(&_ERC20Source.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20Source *ERC20SourceTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20Source.Contract.PauseTeleporterAddress(&_ERC20Source.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20Source *ERC20SourceTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20Source.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20Source *ERC20SourceSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20Source.Contract.ReceiveTeleporterMessage(&_ERC20Source.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20Source *ERC20SourceTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20Source.Contract.ReceiveTeleporterMessage(&_ERC20Source.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Source *ERC20SourceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Source.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Source *ERC20SourceSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20Source.Contract.RenounceOwnership(&_ERC20Source.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Source *ERC20SourceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20Source.Contract.RenounceOwnership(&_ERC20Source.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x32709bc1.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,uint256) input, uint256 amount) returns()
func (_ERC20Source *ERC20SourceTransactor) Send(opts *bind.TransactOpts, input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Source.contract.Transact(opts, "send", input, amount)
}

// Send is a paid mutator transaction binding the contract method 0x32709bc1.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,uint256) input, uint256 amount) returns()
func (_ERC20Source *ERC20SourceSession) Send(input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Source.Contract.Send(&_ERC20Source.TransactOpts, input, amount)
}

// Send is a paid mutator transaction binding the contract method 0x32709bc1.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,uint256) input, uint256 amount) returns()
func (_ERC20Source *ERC20SourceTransactorSession) Send(input SendTokensInput, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Source.Contract.Send(&_ERC20Source.TransactOpts, input, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Source *ERC20SourceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Source.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Source *ERC20SourceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Source.Contract.TransferOwnership(&_ERC20Source.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Source *ERC20SourceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Source.Contract.TransferOwnership(&_ERC20Source.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20Source *ERC20SourceTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20Source.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20Source *ERC20SourceSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20Source.Contract.UnpauseTeleporterAddress(&_ERC20Source.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20Source *ERC20SourceTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20Source.Contract.UnpauseTeleporterAddress(&_ERC20Source.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20Source *ERC20SourceTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _ERC20Source.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20Source *ERC20SourceSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _ERC20Source.Contract.UpdateMinTeleporterVersion(&_ERC20Source.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20Source *ERC20SourceTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _ERC20Source.Contract.UpdateMinTeleporterVersion(&_ERC20Source.TransactOpts, version)
}

// ERC20SourceMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the ERC20Source contract.
type ERC20SourceMinTeleporterVersionUpdatedIterator struct {
	Event *ERC20SourceMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20SourceMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SourceMinTeleporterVersionUpdated)
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
		it.Event = new(ERC20SourceMinTeleporterVersionUpdated)
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
func (it *ERC20SourceMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SourceMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SourceMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the ERC20Source contract.
type ERC20SourceMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ERC20Source *ERC20SourceFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*ERC20SourceMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ERC20Source.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SourceMinTeleporterVersionUpdatedIterator{contract: _ERC20Source.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ERC20Source *ERC20SourceFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *ERC20SourceMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ERC20Source.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SourceMinTeleporterVersionUpdated)
				if err := _ERC20Source.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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
func (_ERC20Source *ERC20SourceFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*ERC20SourceMinTeleporterVersionUpdated, error) {
	event := new(ERC20SourceMinTeleporterVersionUpdated)
	if err := _ERC20Source.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SourceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20Source contract.
type ERC20SourceOwnershipTransferredIterator struct {
	Event *ERC20SourceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20SourceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SourceOwnershipTransferred)
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
		it.Event = new(ERC20SourceOwnershipTransferred)
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
func (it *ERC20SourceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SourceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SourceOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20Source contract.
type ERC20SourceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20Source *ERC20SourceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20SourceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20Source.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SourceOwnershipTransferredIterator{contract: _ERC20Source.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20Source *ERC20SourceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20SourceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20Source.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SourceOwnershipTransferred)
				if err := _ERC20Source.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC20Source *ERC20SourceFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20SourceOwnershipTransferred, error) {
	event := new(ERC20SourceOwnershipTransferred)
	if err := _ERC20Source.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SourceSendTokensIterator is returned from FilterSendTokens and is used to iterate over the raw logs and unpacked data for SendTokens events raised by the ERC20Source contract.
type ERC20SourceSendTokensIterator struct {
	Event *ERC20SourceSendTokens // Event containing the contract specifics and raw log

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
func (it *ERC20SourceSendTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SourceSendTokens)
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
		it.Event = new(ERC20SourceSendTokens)
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
func (it *ERC20SourceSendTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SourceSendTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SourceSendTokens represents a SendTokens event raised by the ERC20Source contract.
type ERC20SourceSendTokens struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSendTokens is a free log retrieval operation binding the contract event 0xdc72eade375f64388998c651c5137da74fbe64aec18f766dea3da18f154ab0a0.
//
// Solidity: event SendTokens(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,uint256,uint256,uint256) input, uint256 amount)
func (_ERC20Source *ERC20SourceFilterer) FilterSendTokens(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*ERC20SourceSendTokensIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20Source.contract.FilterLogs(opts, "SendTokens", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SourceSendTokensIterator{contract: _ERC20Source.contract, event: "SendTokens", logs: logs, sub: sub}, nil
}

// WatchSendTokens is a free log subscription operation binding the contract event 0xdc72eade375f64388998c651c5137da74fbe64aec18f766dea3da18f154ab0a0.
//
// Solidity: event SendTokens(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,uint256,uint256,uint256) input, uint256 amount)
func (_ERC20Source *ERC20SourceFilterer) WatchSendTokens(opts *bind.WatchOpts, sink chan<- *ERC20SourceSendTokens, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20Source.contract.WatchLogs(opts, "SendTokens", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SourceSendTokens)
				if err := _ERC20Source.contract.UnpackLog(event, "SendTokens", log); err != nil {
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

// ParseSendTokens is a log parse operation binding the contract event 0xdc72eade375f64388998c651c5137da74fbe64aec18f766dea3da18f154ab0a0.
//
// Solidity: event SendTokens(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,uint256,uint256,uint256) input, uint256 amount)
func (_ERC20Source *ERC20SourceFilterer) ParseSendTokens(log types.Log) (*ERC20SourceSendTokens, error) {
	event := new(ERC20SourceSendTokens)
	if err := _ERC20Source.contract.UnpackLog(event, "SendTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SourceTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the ERC20Source contract.
type ERC20SourceTeleporterAddressPausedIterator struct {
	Event *ERC20SourceTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *ERC20SourceTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SourceTeleporterAddressPaused)
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
		it.Event = new(ERC20SourceTeleporterAddressPaused)
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
func (it *ERC20SourceTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SourceTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SourceTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the ERC20Source contract.
type ERC20SourceTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ERC20Source *ERC20SourceFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*ERC20SourceTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20Source.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SourceTeleporterAddressPausedIterator{contract: _ERC20Source.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ERC20Source *ERC20SourceFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *ERC20SourceTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20Source.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SourceTeleporterAddressPaused)
				if err := _ERC20Source.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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
func (_ERC20Source *ERC20SourceFilterer) ParseTeleporterAddressPaused(log types.Log) (*ERC20SourceTeleporterAddressPaused, error) {
	event := new(ERC20SourceTeleporterAddressPaused)
	if err := _ERC20Source.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SourceTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the ERC20Source contract.
type ERC20SourceTeleporterAddressUnpausedIterator struct {
	Event *ERC20SourceTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20SourceTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SourceTeleporterAddressUnpaused)
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
		it.Event = new(ERC20SourceTeleporterAddressUnpaused)
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
func (it *ERC20SourceTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SourceTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SourceTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the ERC20Source contract.
type ERC20SourceTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ERC20Source *ERC20SourceFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*ERC20SourceTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20Source.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SourceTeleporterAddressUnpausedIterator{contract: _ERC20Source.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ERC20Source *ERC20SourceFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20SourceTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20Source.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SourceTeleporterAddressUnpaused)
				if err := _ERC20Source.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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
func (_ERC20Source *ERC20SourceFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*ERC20SourceTeleporterAddressUnpaused, error) {
	event := new(ERC20SourceTeleporterAddressUnpaused)
	if err := _ERC20Source.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SourceWithdrawTokensIterator is returned from FilterWithdrawTokens and is used to iterate over the raw logs and unpacked data for WithdrawTokens events raised by the ERC20Source contract.
type ERC20SourceWithdrawTokensIterator struct {
	Event *ERC20SourceWithdrawTokens // Event containing the contract specifics and raw log

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
func (it *ERC20SourceWithdrawTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SourceWithdrawTokens)
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
		it.Event = new(ERC20SourceWithdrawTokens)
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
func (it *ERC20SourceWithdrawTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SourceWithdrawTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SourceWithdrawTokens represents a WithdrawTokens event raised by the ERC20Source contract.
type ERC20SourceWithdrawTokens struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTokens is a free log retrieval operation binding the contract event 0x680f2e4f4032ebf1774e8cdbaddcb1b617a5a606411c8ca96257ada338d3833c.
//
// Solidity: event WithdrawTokens(address indexed recipient, uint256 amount)
func (_ERC20Source *ERC20SourceFilterer) FilterWithdrawTokens(opts *bind.FilterOpts, recipient []common.Address) (*ERC20SourceWithdrawTokensIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20Source.contract.FilterLogs(opts, "WithdrawTokens", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SourceWithdrawTokensIterator{contract: _ERC20Source.contract, event: "WithdrawTokens", logs: logs, sub: sub}, nil
}

// WatchWithdrawTokens is a free log subscription operation binding the contract event 0x680f2e4f4032ebf1774e8cdbaddcb1b617a5a606411c8ca96257ada338d3833c.
//
// Solidity: event WithdrawTokens(address indexed recipient, uint256 amount)
func (_ERC20Source *ERC20SourceFilterer) WatchWithdrawTokens(opts *bind.WatchOpts, sink chan<- *ERC20SourceWithdrawTokens, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20Source.contract.WatchLogs(opts, "WithdrawTokens", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SourceWithdrawTokens)
				if err := _ERC20Source.contract.UnpackLog(event, "WithdrawTokens", log); err != nil {
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

// ParseWithdrawTokens is a log parse operation binding the contract event 0x680f2e4f4032ebf1774e8cdbaddcb1b617a5a606411c8ca96257ada338d3833c.
//
// Solidity: event WithdrawTokens(address indexed recipient, uint256 amount)
func (_ERC20Source *ERC20SourceFilterer) ParseWithdrawTokens(log types.Log) (*ERC20SourceWithdrawTokens, error) {
	event := new(ERC20SourceWithdrawTokens)
	if err := _ERC20Source.contract.UnpackLog(event, "WithdrawTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
