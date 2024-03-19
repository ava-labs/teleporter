// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokensource

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
	AllowedRelayerAddresses  []common.Address
}

// NativeTokenSourceMetaData contains all meta data concerning the NativeTokenSource contract.
var NativeTokenSourceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SendTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SEND_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"}],\"name\":\"bridgedBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIWrappedNativeToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620021793803806200217983398101604081905262000035916200033a565b60016000558282828282816001600160a01b038116620000c25760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000133919062000384565b600255506200014233620001ee565b6200014d8162000240565b50506001600160a01b03811660c0526040805163084279ef60e31b8152905173020000000000000000000000000000000000000591634213cf789160048083019260209291908290030181865afa158015620001ad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001d3919062000384565b60a0525050506001600160a01b031660e052506200039e9050565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6200024a620002bf565b6001600160a01b038116620002b15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401620000b9565b620002bc81620001ee565b50565b6003546001600160a01b031633146200031b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401620000b9565b565b80516001600160a01b03811681146200033557600080fd5b919050565b6000806000606084860312156200035057600080fd5b6200035b846200031d565b92506200036b602085016200031d565b91506200037b604085016200031d565b90509250925092565b6000602082840312156200039757600080fd5b5051919050565b60805160a05160c05160e051611d626200041760003960008181610406015281816110b60152611159015260008181610107015281816103360152610fc501526000818161038a01528181610c250152610cde01526000818161020e0152818161065e015281816108a901526112fc0152611d626000f3fe6080604052600436106100f75760003560e01c8063973142971161008a578063d2cc7a7011610059578063d2cc7a70146103ac578063e30fd71d146103c1578063f2fde38b146103d4578063fc0c546a146103f457600080fd5b806397314297146102db578063b8df0dea14610324578063c868efaa14610358578063d127dc9b1461037857600080fd5b80634511243e116100c65780634511243e146102685780635eb9951414610288578063715018a6146102a85780638da5cb5b146102bd57600080fd5b806302ee3e9c1461019a57806310397b01146101e55780631a7f5bec146101fc5780632b0d8f181461024857600080fd5b3661019557336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101935760405162461bcd60e51b815260206004820152603160248201527f4e6174697665546f6b656e536f757263653a20696e76616c69642072656365696044820152703b32903830bcb0b136329039b2b73232b960791b60648201526084015b60405180910390fd5b005b600080fd5b3480156101a657600080fd5b506101d26101b536600461173f565b600460209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b3480156101f157600080fd5b506101d2620493e081565b34801561020857600080fd5b506102307f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101dc565b34801561025457600080fd5b5061019361026336600461176f565b610428565b34801561027457600080fd5b5061019361028336600461176f565b610524565b34801561029457600080fd5b506101936102a336600461178c565b610621565b3480156102b457600080fd5b50610193610635565b3480156102c957600080fd5b506003546001600160a01b0316610230565b3480156102e757600080fd5b506103146102f636600461176f565b6001600160a01b031660009081526001602052604090205460ff1690565b60405190151581526020016101dc565b34801561033057600080fd5b506102307f000000000000000000000000000000000000000000000000000000000000000081565b34801561036457600080fd5b506101936103733660046117a5565b610649565b34801561038457600080fd5b506101d27f000000000000000000000000000000000000000000000000000000000000000081565b3480156103b857600080fd5b506002546101d2565b6101936103cf36600461182e565b610813565b3480156103e057600080fd5b506101936103ef36600461176f565b610827565b34801561040057600080fd5b506102307f000000000000000000000000000000000000000000000000000000000000000081565b61043061089d565b6001600160a01b0381166104565760405162461bcd60e51b815260040161018a90611869565b6001600160a01b03811660009081526001602052604090205460ff16156104d55760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b606482015260840161018a565b6001600160a01b0381166000818152600160208190526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b61052c61089d565b6001600160a01b0381166105525760405162461bcd60e51b815260040161018a90611869565b6001600160a01b03811660009081526001602052604090205460ff166105cc5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b606482015260840161018a565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600160205260409020805460ff19169055565b61062961089d565b610632816108a5565b50565b61063d610a45565b6106476000610a9f565b565b610651610af1565b6002546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa1580156106c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ec91906118b7565b10156107535760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b606482015260840161018a565b61075c336102f6565b156107c25760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b606482015260840161018a565b610803848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610b4a92505050565b61080d6001600055565b50505050565b61063261081f82611964565b346000610cda565b61082f610a45565b6001600160a01b0381166108945760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161018a565b61063281610a9f565b610647610a45565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610905573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092991906118b7565b600254909150818311156109995760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b606482015260840161018a565b808311610a0e5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e00606482015260840161018a565b6002839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6003546001600160a01b031633146106475760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161018a565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600260005403610b435760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161018a565b6002600055565b60008082806020019051810190610b619190611a52565b60008781526004602090815260408083206001600160a01b038a168452909152902054919350915081811015610bf45760405162461bcd60e51b815260206004820152603260248201527f54656c65706f72746572546f6b656e536f757263653a20696e73756666696369604482015271656e74206272696467652062616c616e636560701b606482015260840161018a565b610bfe8282611b82565b60008781526004602090815260408083206001600160a01b038a16845290915290205582517f00000000000000000000000000000000000000000000000000000000000000009003610cd25760208301516001600160a01b03163014610cbc5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572546f6b656e536f757263653a20696e76616c6964206260448201526c7269646765206164647265737360981b606482015260840161018a565b610cca8360400151836110a0565b505050505050565b610cca838360015b82517f00000000000000000000000000000000000000000000000000000000000000009003610d665760405162461bcd60e51b815260206004820152603260248201527f54656c65706f72746572546f6b656e536f757263653a2063616e6e6f7420627260448201527134b233b2903a379039b0b6b29031b430b4b760711b606482015260840161018a565b60208301516001600160a01b0316610ddf5760405162461bcd60e51b815260206004820152603660248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f2064657374604482015275696e6174696f6e20627269646765206164647265737360501b606482015260840161018a565b60408301516001600160a01b0316610e4f5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f207265636960448201526c7069656e74206164647265737360981b606482015260840161018a565b80610e6057610e5d82611155565b91505b60008211610ec05760405162461bcd60e51b815260206004820152602760248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f2073656e6460448201526608185b5bdd5b9d60ca1b606482015260840161018a565b82606001518211610f395760405162461bcd60e51b815260206004820152603860248201527f54656c65706f72746572546f6b656e536f757263653a20696e7375666669636960448201527f656e7420616d6f756e7420746f20636f76657220666565730000000000000000606482015260840161018a565b6060830151610f489083611b82565b83516000908152600460209081526040808320828801516001600160a01b03168452909152812080549294508492909190610f84908490611b9b565b9250508190555060006110546040518060c001604052808660000151815260200186602001516001600160a01b0316815260200160405180604001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200188606001518152508152602001620493e081526020018660a00151815260200186604001518660405160200161103d9291906001600160a01b03929092168252602082015260400190565b6040516020818303038152906040528152506111d1565b9050336001600160a01b0316817f35971507d5275c40cb225115b67b12c9e28b4bb02e60366a5863d8320bd638658560405161109291815260200190565b60405180910390a350505050565b604051632e1a7d4d60e01b8152600481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d90602401600060405180830381600087803b15801561110257600080fd5b505af1158015611116573d6000803e3d6000fd5b50506040516001600160a01b038516925083156108fc02915083906000818181858888f19350505050158015611150573d6000803e3d6000fd5b505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0836040518263ffffffff1660e01b81526004016000604051808303818588803b1580156111b257600080fd5b505af11580156111c6573d6000803e3d6000fd5b509495945050505050565b6000806111dc6112f7565b60408401516020015190915015611281576040830151516001600160a01b031661125e5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b606482015260840161018a565b604083015160208101519051611281916001600160a01b0390911690839061140b565b604051630624488560e41b81526001600160a01b038216906362448850906112ad908690600401611c37565b6020604051808303816000875af11580156112cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112f091906118b7565b9392505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611358573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061137c9190611cbe565b90506113a0816001600160a01b031660009081526001602052604090205460ff1690565b156114065760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b606482015260840161018a565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa15801561145c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061148091906118b7565b61148a9190611b9b565b604080516001600160a01b03868116602483015260448083018590528351808403909101815260649092018352602080830180516001600160e01b031663095ea7b360e01b17905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65649084015292935061080d92879260009161151d9190851690849061159a565b805190915015611150578080602001905181019061153b9190611cdb565b6111505760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161018a565b60606115a984846000856115b1565b949350505050565b6060824710156116125760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161018a565b600080866001600160a01b0316858760405161162e9190611cfd565b60006040518083038185875af1925050503d806000811461166b576040519150601f19603f3d011682016040523d82523d6000602084013e611670565b606091505b50915091506116818783838761168c565b979650505050505050565b606083156116fb5782516000036116f4576001600160a01b0385163b6116f45760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161018a565b50816115a9565b6115a983838151156117105781518083602001fd5b8060405162461bcd60e51b815260040161018a9190611d19565b6001600160a01b038116811461063257600080fd5b6000806040838503121561175257600080fd5b8235915060208301356117648161172a565b809150509250929050565b60006020828403121561178157600080fd5b81356112f08161172a565b60006020828403121561179e57600080fd5b5035919050565b600080600080606085870312156117bb57600080fd5b8435935060208501356117cd8161172a565b9250604085013567ffffffffffffffff808211156117ea57600080fd5b818701915087601f8301126117fe57600080fd5b81358181111561180d57600080fd5b88602082850101111561181f57600080fd5b95989497505060200194505050565b60006020828403121561184057600080fd5b813567ffffffffffffffff81111561185757600080fd5b820160c081850312156112f057600080fd5b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b6000602082840312156118c957600080fd5b5051919050565b634e487b7160e01b600052604160045260246000fd5b60405160c0810167ffffffffffffffff81118282101715611909576119096118d0565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715611938576119386118d0565b604052919050565b600067ffffffffffffffff82111561195a5761195a6118d0565b5060051b60200190565b600060c0823603121561197657600080fd5b61197e6118e6565b823581526020808401356119918161172a565b8282015260408401356119a38161172a565b80604084015250606084013560608301526080840135608083015260a084013567ffffffffffffffff8111156119d857600080fd5b840136601f8201126119e957600080fd5b80356119fc6119f782611940565b61190f565b81815260059190911b82018301908381019036831115611a1b57600080fd5b928401925b82841015611a42578335611a338161172a565b82529284019290840190611a20565b60a0860152509295945050505050565b60008060408385031215611a6557600080fd5b825167ffffffffffffffff80821115611a7d57600080fd5b9084019060c08287031215611a9157600080fd5b611a996118e6565b82518152602080840151611aac8161172a565b828201526040840151611abe8161172a565b80604084015250606084015160608301526080840151608083015260a084015183811115611aeb57600080fd5b80850194505087601f850112611b0057600080fd5b83519250611b106119f784611940565b83815260059390931b84018101928181019089851115611b2f57600080fd5b948201945b84861015611b56578551611b478161172a565b82529482019490820190611b34565b60a0840152509590950151949694955050505050565b634e487b7160e01b600052601160045260246000fd5b81810381811115611b9557611b95611b6c565b92915050565b80820180821115611b9557611b95611b6c565b600081518084526020808501945080840160005b838110156111c65781516001600160a01b031687529582019590820190600101611bc2565b60005b83811015611c02578181015183820152602001611bea565b50506000910152565b60008151808452611c23816020860160208601611be7565b601f01601f19169290920160200192915050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c0840152611c98610100840182611bae565b905060a0840151601f198483030160e0850152611cb58282611c0b565b95945050505050565b600060208284031215611cd057600080fd5b81516112f08161172a565b600060208284031215611ced57600080fd5b815180151581146112f057600080fd5b60008251611d0f818460208701611be7565b9190910192915050565b6020815260006112f06020830184611c0b56fea26469706673582212208ddf6a04b7421f0133dfda1673ecf4c8d2121d874fc01e20aae918fed223dd9b64736f6c63430008120033",
}

// NativeTokenSourceABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenSourceMetaData.ABI instead.
var NativeTokenSourceABI = NativeTokenSourceMetaData.ABI

// NativeTokenSourceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenSourceMetaData.Bin instead.
var NativeTokenSourceBin = NativeTokenSourceMetaData.Bin

// DeployNativeTokenSource deploys a new Ethereum contract, binding an instance of NativeTokenSource to it.
func DeployNativeTokenSource(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterRegistryAddress common.Address, teleporterManager common.Address, feeTokenAddress common.Address) (common.Address, *types.Transaction, *NativeTokenSource, error) {
	parsed, err := NativeTokenSourceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenSourceBin), backend, teleporterRegistryAddress, teleporterManager, feeTokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenSource{NativeTokenSourceCaller: NativeTokenSourceCaller{contract: contract}, NativeTokenSourceTransactor: NativeTokenSourceTransactor{contract: contract}, NativeTokenSourceFilterer: NativeTokenSourceFilterer{contract: contract}}, nil
}

// NativeTokenSource is an auto generated Go binding around an Ethereum contract.
type NativeTokenSource struct {
	NativeTokenSourceCaller     // Read-only binding to the contract
	NativeTokenSourceTransactor // Write-only binding to the contract
	NativeTokenSourceFilterer   // Log filterer for contract events
}

// NativeTokenSourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenSourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenSourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenSourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenSourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenSourceSession struct {
	Contract     *NativeTokenSource // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NativeTokenSourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenSourceCallerSession struct {
	Contract *NativeTokenSourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NativeTokenSourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenSourceTransactorSession struct {
	Contract     *NativeTokenSourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NativeTokenSourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenSourceRaw struct {
	Contract *NativeTokenSource // Generic contract binding to access the raw methods on
}

// NativeTokenSourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenSourceCallerRaw struct {
	Contract *NativeTokenSourceCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenSourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenSourceTransactorRaw struct {
	Contract *NativeTokenSourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenSource creates a new instance of NativeTokenSource, bound to a specific deployed contract.
func NewNativeTokenSource(address common.Address, backend bind.ContractBackend) (*NativeTokenSource, error) {
	contract, err := bindNativeTokenSource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSource{NativeTokenSourceCaller: NativeTokenSourceCaller{contract: contract}, NativeTokenSourceTransactor: NativeTokenSourceTransactor{contract: contract}, NativeTokenSourceFilterer: NativeTokenSourceFilterer{contract: contract}}, nil
}

// NewNativeTokenSourceCaller creates a new read-only instance of NativeTokenSource, bound to a specific deployed contract.
func NewNativeTokenSourceCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenSourceCaller, error) {
	contract, err := bindNativeTokenSource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceCaller{contract: contract}, nil
}

// NewNativeTokenSourceTransactor creates a new write-only instance of NativeTokenSource, bound to a specific deployed contract.
func NewNativeTokenSourceTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenSourceTransactor, error) {
	contract, err := bindNativeTokenSource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceTransactor{contract: contract}, nil
}

// NewNativeTokenSourceFilterer creates a new log filterer instance of NativeTokenSource, bound to a specific deployed contract.
func NewNativeTokenSourceFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenSourceFilterer, error) {
	contract, err := bindNativeTokenSource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceFilterer{contract: contract}, nil
}

// bindNativeTokenSource binds a generic wrapper to an already deployed contract.
func bindNativeTokenSource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenSourceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenSource *NativeTokenSourceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenSource.Contract.NativeTokenSourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenSource *NativeTokenSourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.NativeTokenSourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenSource *NativeTokenSourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.NativeTokenSourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenSource *NativeTokenSourceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenSource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenSource *NativeTokenSourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenSource *NativeTokenSourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.contract.Transact(opts, method, params...)
}

// SENDTOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x10397b01.
//
// Solidity: function SEND_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceCaller) SENDTOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "SEND_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SENDTOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x10397b01.
//
// Solidity: function SEND_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceSession) SENDTOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSource.Contract.SENDTOKENSREQUIREDGAS(&_NativeTokenSource.CallOpts)
}

// SENDTOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0x10397b01.
//
// Solidity: function SEND_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceCallerSession) SENDTOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSource.Contract.SENDTOKENSREQUIREDGAS(&_NativeTokenSource.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenSource *NativeTokenSourceCaller) BlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "blockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenSource *NativeTokenSourceSession) BlockchainID() ([32]byte, error) {
	return _NativeTokenSource.Contract.BlockchainID(&_NativeTokenSource.CallOpts)
}

// BlockchainID is a free data retrieval call binding the contract method 0xd127dc9b.
//
// Solidity: function blockchainID() view returns(bytes32)
func (_NativeTokenSource *NativeTokenSourceCallerSession) BlockchainID() ([32]byte, error) {
	return _NativeTokenSource.Contract.BlockchainID(&_NativeTokenSource.CallOpts)
}

// BridgedBalances is a free data retrieval call binding the contract method 0x02ee3e9c.
//
// Solidity: function bridgedBalances(bytes32 destinationBlockchainID, address destinationBridgeAddress) view returns(uint256 balance)
func (_NativeTokenSource *NativeTokenSourceCaller) BridgedBalances(opts *bind.CallOpts, destinationBlockchainID [32]byte, destinationBridgeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "bridgedBalances", destinationBlockchainID, destinationBridgeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgedBalances is a free data retrieval call binding the contract method 0x02ee3e9c.
//
// Solidity: function bridgedBalances(bytes32 destinationBlockchainID, address destinationBridgeAddress) view returns(uint256 balance)
func (_NativeTokenSource *NativeTokenSourceSession) BridgedBalances(destinationBlockchainID [32]byte, destinationBridgeAddress common.Address) (*big.Int, error) {
	return _NativeTokenSource.Contract.BridgedBalances(&_NativeTokenSource.CallOpts, destinationBlockchainID, destinationBridgeAddress)
}

// BridgedBalances is a free data retrieval call binding the contract method 0x02ee3e9c.
//
// Solidity: function bridgedBalances(bytes32 destinationBlockchainID, address destinationBridgeAddress) view returns(uint256 balance)
func (_NativeTokenSource *NativeTokenSourceCallerSession) BridgedBalances(destinationBlockchainID [32]byte, destinationBridgeAddress common.Address) (*big.Int, error) {
	return _NativeTokenSource.Contract.BridgedBalances(&_NativeTokenSource.CallOpts, destinationBlockchainID, destinationBridgeAddress)
}

// FeeTokenAddress is a free data retrieval call binding the contract method 0xb8df0dea.
//
// Solidity: function feeTokenAddress() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCaller) FeeTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "feeTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenAddress is a free data retrieval call binding the contract method 0xb8df0dea.
//
// Solidity: function feeTokenAddress() view returns(address)
func (_NativeTokenSource *NativeTokenSourceSession) FeeTokenAddress() (common.Address, error) {
	return _NativeTokenSource.Contract.FeeTokenAddress(&_NativeTokenSource.CallOpts)
}

// FeeTokenAddress is a free data retrieval call binding the contract method 0xb8df0dea.
//
// Solidity: function feeTokenAddress() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCallerSession) FeeTokenAddress() (common.Address, error) {
	return _NativeTokenSource.Contract.FeeTokenAddress(&_NativeTokenSource.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenSource.Contract.GetMinTeleporterVersion(&_NativeTokenSource.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _NativeTokenSource.Contract.GetMinTeleporterVersion(&_NativeTokenSource.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenSource *NativeTokenSourceCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenSource *NativeTokenSourceSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenSource.Contract.IsTeleporterAddressPaused(&_NativeTokenSource.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_NativeTokenSource *NativeTokenSourceCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _NativeTokenSource.Contract.IsTeleporterAddressPaused(&_NativeTokenSource.CallOpts, teleporterAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenSource *NativeTokenSourceSession) Owner() (common.Address, error) {
	return _NativeTokenSource.Contract.Owner(&_NativeTokenSource.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCallerSession) Owner() (common.Address, error) {
	return _NativeTokenSource.Contract.Owner(&_NativeTokenSource.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenSource *NativeTokenSourceSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenSource.Contract.TeleporterRegistry(&_NativeTokenSource.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCallerSession) TeleporterRegistry() (common.Address, error) {
	return _NativeTokenSource.Contract.TeleporterRegistry(&_NativeTokenSource.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_NativeTokenSource *NativeTokenSourceSession) Token() (common.Address, error) {
	return _NativeTokenSource.Contract.Token(&_NativeTokenSource.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCallerSession) Token() (common.Address, error) {
	return _NativeTokenSource.Contract.Token(&_NativeTokenSource.CallOpts)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSource *NativeTokenSourceTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSource.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSource *NativeTokenSourceSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.PauseTeleporterAddress(&_NativeTokenSource.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSource *NativeTokenSourceTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.PauseTeleporterAddress(&_NativeTokenSource.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenSource *NativeTokenSourceTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSource.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenSource *NativeTokenSourceSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.ReceiveTeleporterMessage(&_NativeTokenSource.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_NativeTokenSource *NativeTokenSourceTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.ReceiveTeleporterMessage(&_NativeTokenSource.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenSource *NativeTokenSourceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSource.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenSource *NativeTokenSourceSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenSource.Contract.RenounceOwnership(&_NativeTokenSource.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NativeTokenSource *NativeTokenSourceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NativeTokenSource.Contract.RenounceOwnership(&_NativeTokenSource.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xe30fd71d.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,address[]) input) payable returns()
func (_NativeTokenSource *NativeTokenSourceTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenSource.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xe30fd71d.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,address[]) input) payable returns()
func (_NativeTokenSource *NativeTokenSourceSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.Send(&_NativeTokenSource.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xe30fd71d.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,address[]) input) payable returns()
func (_NativeTokenSource *NativeTokenSourceTransactorSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.Send(&_NativeTokenSource.TransactOpts, input)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenSource *NativeTokenSourceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenSource.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenSource *NativeTokenSourceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.TransferOwnership(&_NativeTokenSource.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NativeTokenSource *NativeTokenSourceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.TransferOwnership(&_NativeTokenSource.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSource *NativeTokenSourceTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSource.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSource *NativeTokenSourceSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.UnpauseTeleporterAddress(&_NativeTokenSource.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_NativeTokenSource *NativeTokenSourceTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.UnpauseTeleporterAddress(&_NativeTokenSource.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenSource *NativeTokenSourceTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _NativeTokenSource.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenSource *NativeTokenSourceSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.UpdateMinTeleporterVersion(&_NativeTokenSource.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_NativeTokenSource *NativeTokenSourceTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.UpdateMinTeleporterVersion(&_NativeTokenSource.TransactOpts, version)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenSource *NativeTokenSourceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenSource.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenSource *NativeTokenSourceSession) Receive() (*types.Transaction, error) {
	return _NativeTokenSource.Contract.Receive(&_NativeTokenSource.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NativeTokenSource *NativeTokenSourceTransactorSession) Receive() (*types.Transaction, error) {
	return _NativeTokenSource.Contract.Receive(&_NativeTokenSource.TransactOpts)
}

// NativeTokenSourceMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the NativeTokenSource contract.
type NativeTokenSourceMinTeleporterVersionUpdatedIterator struct {
	Event *NativeTokenSourceMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *NativeTokenSourceMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSourceMinTeleporterVersionUpdated)
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
		it.Event = new(NativeTokenSourceMinTeleporterVersionUpdated)
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
func (it *NativeTokenSourceMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSourceMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSourceMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the NativeTokenSource contract.
type NativeTokenSourceMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenSource *NativeTokenSourceFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*NativeTokenSourceMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenSource.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceMinTeleporterVersionUpdatedIterator{contract: _NativeTokenSource.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_NativeTokenSource *NativeTokenSourceFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenSourceMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _NativeTokenSource.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSourceMinTeleporterVersionUpdated)
				if err := _NativeTokenSource.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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
func (_NativeTokenSource *NativeTokenSourceFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*NativeTokenSourceMinTeleporterVersionUpdated, error) {
	event := new(NativeTokenSourceMinTeleporterVersionUpdated)
	if err := _NativeTokenSource.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSourceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NativeTokenSource contract.
type NativeTokenSourceOwnershipTransferredIterator struct {
	Event *NativeTokenSourceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NativeTokenSourceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSourceOwnershipTransferred)
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
		it.Event = new(NativeTokenSourceOwnershipTransferred)
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
func (it *NativeTokenSourceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSourceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSourceOwnershipTransferred represents a OwnershipTransferred event raised by the NativeTokenSource contract.
type NativeTokenSourceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenSource *NativeTokenSourceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeTokenSourceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenSource.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceOwnershipTransferredIterator{contract: _NativeTokenSource.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NativeTokenSource *NativeTokenSourceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeTokenSourceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NativeTokenSource.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSourceOwnershipTransferred)
				if err := _NativeTokenSource.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NativeTokenSource *NativeTokenSourceFilterer) ParseOwnershipTransferred(log types.Log) (*NativeTokenSourceOwnershipTransferred, error) {
	event := new(NativeTokenSourceOwnershipTransferred)
	if err := _NativeTokenSource.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSourceSendTokensIterator is returned from FilterSendTokens and is used to iterate over the raw logs and unpacked data for SendTokens events raised by the NativeTokenSource contract.
type NativeTokenSourceSendTokensIterator struct {
	Event *NativeTokenSourceSendTokens // Event containing the contract specifics and raw log

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
func (it *NativeTokenSourceSendTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSourceSendTokens)
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
		it.Event = new(NativeTokenSourceSendTokens)
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
func (it *NativeTokenSourceSendTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSourceSendTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSourceSendTokens represents a SendTokens event raised by the NativeTokenSource contract.
type NativeTokenSourceSendTokens struct {
	TeleporterMessageID [32]byte
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSendTokens is a free log retrieval operation binding the contract event 0x35971507d5275c40cb225115b67b12c9e28b4bb02e60366a5863d8320bd63865.
//
// Solidity: event SendTokens(bytes32 indexed teleporterMessageID, address indexed sender, uint256 amount)
func (_NativeTokenSource *NativeTokenSourceFilterer) FilterSendTokens(opts *bind.FilterOpts, teleporterMessageID [][32]byte, sender []common.Address) (*NativeTokenSourceSendTokensIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSource.contract.FilterLogs(opts, "SendTokens", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceSendTokensIterator{contract: _NativeTokenSource.contract, event: "SendTokens", logs: logs, sub: sub}, nil
}

// WatchSendTokens is a free log subscription operation binding the contract event 0x35971507d5275c40cb225115b67b12c9e28b4bb02e60366a5863d8320bd63865.
//
// Solidity: event SendTokens(bytes32 indexed teleporterMessageID, address indexed sender, uint256 amount)
func (_NativeTokenSource *NativeTokenSourceFilterer) WatchSendTokens(opts *bind.WatchOpts, sink chan<- *NativeTokenSourceSendTokens, teleporterMessageID [][32]byte, sender []common.Address) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NativeTokenSource.contract.WatchLogs(opts, "SendTokens", teleporterMessageIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSourceSendTokens)
				if err := _NativeTokenSource.contract.UnpackLog(event, "SendTokens", log); err != nil {
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

// ParseSendTokens is a log parse operation binding the contract event 0x35971507d5275c40cb225115b67b12c9e28b4bb02e60366a5863d8320bd63865.
//
// Solidity: event SendTokens(bytes32 indexed teleporterMessageID, address indexed sender, uint256 amount)
func (_NativeTokenSource *NativeTokenSourceFilterer) ParseSendTokens(log types.Log) (*NativeTokenSourceSendTokens, error) {
	event := new(NativeTokenSourceSendTokens)
	if err := _NativeTokenSource.contract.UnpackLog(event, "SendTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSourceTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the NativeTokenSource contract.
type NativeTokenSourceTeleporterAddressPausedIterator struct {
	Event *NativeTokenSourceTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenSourceTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSourceTeleporterAddressPaused)
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
		it.Event = new(NativeTokenSourceTeleporterAddressPaused)
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
func (it *NativeTokenSourceTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSourceTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSourceTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the NativeTokenSource contract.
type NativeTokenSourceTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenSource *NativeTokenSourceFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenSourceTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenSource.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceTeleporterAddressPausedIterator{contract: _NativeTokenSource.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_NativeTokenSource *NativeTokenSourceFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *NativeTokenSourceTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenSource.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSourceTeleporterAddressPaused)
				if err := _NativeTokenSource.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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
func (_NativeTokenSource *NativeTokenSourceFilterer) ParseTeleporterAddressPaused(log types.Log) (*NativeTokenSourceTeleporterAddressPaused, error) {
	event := new(NativeTokenSourceTeleporterAddressPaused)
	if err := _NativeTokenSource.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSourceTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the NativeTokenSource contract.
type NativeTokenSourceTeleporterAddressUnpausedIterator struct {
	Event *NativeTokenSourceTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *NativeTokenSourceTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSourceTeleporterAddressUnpaused)
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
		it.Event = new(NativeTokenSourceTeleporterAddressUnpaused)
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
func (it *NativeTokenSourceTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSourceTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSourceTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the NativeTokenSource contract.
type NativeTokenSourceTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenSource *NativeTokenSourceFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*NativeTokenSourceTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenSource.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceTeleporterAddressUnpausedIterator{contract: _NativeTokenSource.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_NativeTokenSource *NativeTokenSourceFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *NativeTokenSourceTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _NativeTokenSource.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSourceTeleporterAddressUnpaused)
				if err := _NativeTokenSource.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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
func (_NativeTokenSource *NativeTokenSourceFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*NativeTokenSourceTeleporterAddressUnpaused, error) {
	event := new(NativeTokenSourceTeleporterAddressUnpaused)
	if err := _NativeTokenSource.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
