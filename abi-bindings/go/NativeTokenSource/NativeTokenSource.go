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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SendTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SEND_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"}],\"name\":\"bridgedBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162001f7738038062001f7783398101604081905262000034916200032e565b60016000558282828282816001600160a01b038116620000c15760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa1580156200010c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000132919062000378565b600255506200014133620001e2565b6200014c8162000234565b50506001600160a01b03811660c0526040805163084279ef60e31b8152905173020000000000000000000000000000000000000591634213cf789160048083019260209291908290030181865afa158015620001ac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001d2919062000378565b60a0525062000392945050505050565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6200023e620002b3565b6001600160a01b038116620002a55760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401620000b8565b620002b081620001e2565b50565b6003546001600160a01b031633146200030f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401620000b8565b565b80516001600160a01b03811681146200032957600080fd5b919050565b6000806000606084860312156200034457600080fd5b6200034f8462000311565b92506200035f6020850162000311565b91506200036f6040850162000311565b90509250925092565b6000602082840312156200038b57600080fd5b5051919050565b60805160a05160c051611b8b620003ec6000396000818161028b0152610eda0152600081816102df01528181610b4d0152610c0401526000818161016101528181610588015281816107d1015261111a0152611b8b6000f3fe6080604052600436106100e85760003560e01c80638da5cb5b1161008a578063d127dc9b11610059578063d127dc9b146102cd578063d2cc7a7014610301578063e30fd71d14610316578063f2fde38b1461032957600080fd5b80638da5cb5b146102125780639731429714610230578063b8df0dea14610279578063c868efaa146102ad57600080fd5b80632b0d8f18116100c65780632b0d8f181461019b5780634511243e146101bd5780635eb99514146101dd578063715018a6146101fd57600080fd5b806302ee3e9c146100ed57806310397b01146101385780631a7f5bec1461014f575b600080fd5b3480156100f957600080fd5b5061012561010836600461155d565b600460209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b34801561014457600080fd5b50610125620493e081565b34801561015b57600080fd5b506101837f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161012f565b3480156101a757600080fd5b506101bb6101b636600461158d565b610349565b005b3480156101c957600080fd5b506101bb6101d836600461158d565b61044e565b3480156101e957600080fd5b506101bb6101f83660046115aa565b61054b565b34801561020957600080fd5b506101bb61055f565b34801561021e57600080fd5b506003546001600160a01b0316610183565b34801561023c57600080fd5b5061026961024b36600461158d565b6001600160a01b031660009081526001602052604090205460ff1690565b604051901515815260200161012f565b34801561028557600080fd5b506101837f000000000000000000000000000000000000000000000000000000000000000081565b3480156102b957600080fd5b506101bb6102c83660046115c3565b610573565b3480156102d957600080fd5b506101257f000000000000000000000000000000000000000000000000000000000000000081565b34801561030d57600080fd5b50600254610125565b6101bb61032436600461164c565b61073d565b34801561033557600080fd5b506101bb61034436600461158d565b61074f565b6103516107c5565b6001600160a01b0381166103805760405162461bcd60e51b815260040161037790611687565b60405180910390fd5b6001600160a01b03811660009081526001602052604090205460ff16156103ff5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610377565b6001600160a01b0381166000818152600160208190526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b6104566107c5565b6001600160a01b03811661047c5760405162461bcd60e51b815260040161037790611687565b6001600160a01b03811660009081526001602052604090205460ff166104f65760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610377565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600160205260409020805460ff19169055565b6105536107c5565b61055c816107cd565b50565b61056761096d565b61057160006109c7565b565b61057b610a19565b6002546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa1580156105f2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061691906116d5565b101561067d5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610377565b6106863361024b565b156106ec5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610377565b61072d848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610a7292505050565b6107376001600055565b50505050565b61055c61074982611782565b34610c00565b61075761096d565b6001600160a01b0381166107bc5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610377565b61055c816109c7565b61057161096d565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa15801561082d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085191906116d5565b600254909150818311156108c15760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610377565b8083116109365760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610377565b6002839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6003546001600160a01b031633146105715760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610377565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600260005403610a6b5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610377565b6002600055565b60008082806020019051810190610a899190611870565b60008781526004602090815260408083206001600160a01b038a168452909152902054919350915081811015610b1c5760405162461bcd60e51b815260206004820152603260248201527f54656c65706f72746572546f6b656e536f757263653a20696e73756666696369604482015271656e74206272696467652062616c616e636560701b6064820152608401610377565b610b2682826119a0565b60008781526004602090815260408083206001600160a01b038a16845290915290205582517f00000000000000000000000000000000000000000000000000000000000000009003610bfa5760208301516001600160a01b03163014610be45760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572546f6b656e536f757263653a20696e76616c6964206260448201526c7269646765206164647265737360981b6064820152608401610377565b610bf2836040015183610fb4565b505050505050565b610bf283835b81517f00000000000000000000000000000000000000000000000000000000000000009003610c8c5760405162461bcd60e51b815260206004820152603260248201527f54656c65706f72746572546f6b656e536f757263653a2063616e6e6f7420627260448201527134b233b2903a379039b0b6b29031b430b4b760711b6064820152608401610377565b60208201516001600160a01b0316610d055760405162461bcd60e51b815260206004820152603660248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f2064657374604482015275696e6174696f6e20627269646765206164647265737360501b6064820152608401610377565b60408201516001600160a01b0316610d755760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f207265636960448201526c7069656e74206164647265737360981b6064820152608401610377565b60008111610dd55760405162461bcd60e51b815260206004820152602760248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f2073656e6460448201526608185b5bdd5b9d60ca1b6064820152608401610377565b81606001518111610e4e5760405162461bcd60e51b815260206004820152603860248201527f54656c65706f72746572546f6b656e536f757263653a20696e7375666669636960448201527f656e7420616d6f756e7420746f20636f766572206665657300000000000000006064820152608401610377565b6060820151610e5d90826119a0565b82516000908152600460209081526040808320828701516001600160a01b03168452909152812080549293508392909190610e999084906119b9565b925050819055506000610f696040518060c001604052808560000151815260200185602001516001600160a01b0316815260200160405180604001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200187606001518152508152602001620493e081526020018560a001518152602001856040015185604051602001610f529291906001600160a01b03929092168252602082015260400190565b604051602081830303815290604052815250610fef565b9050336001600160a01b0316817f35971507d5275c40cb225115b67b12c9e28b4bb02e60366a5863d8320bd6386584604051610fa791815260200190565b60405180910390a3505050565b6040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015610fea573d6000803e3d6000fd5b505050565b600080610ffa611115565b6040840151602001519091501561109f576040830151516001600160a01b031661107c5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610377565b60408301516020810151905161109f916001600160a01b03909116908390611229565b604051630624488560e41b81526001600160a01b038216906362448850906110cb908690600401611a60565b6020604051808303816000875af11580156110ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061110e91906116d5565b9392505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611176573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061119a9190611ae7565b90506111be816001600160a01b031660009081526001602052604090205460ff1690565b156112245760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610377565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa15801561127a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061129e91906116d5565b6112a891906119b9565b604080516001600160a01b03868116602483015260448083018590528351808403909101815260649092018352602080830180516001600160e01b031663095ea7b360e01b17905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65649084015292935061073792879260009161133b919085169084906113b8565b805190915015610fea57808060200190518101906113599190611b04565b610fea5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610377565b60606113c784846000856113cf565b949350505050565b6060824710156114305760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610377565b600080866001600160a01b0316858760405161144c9190611b26565b60006040518083038185875af1925050503d8060008114611489576040519150601f19603f3d011682016040523d82523d6000602084013e61148e565b606091505b509150915061149f878383876114aa565b979650505050505050565b60608315611519578251600003611512576001600160a01b0385163b6115125760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610377565b50816113c7565b6113c7838381511561152e5781518083602001fd5b8060405162461bcd60e51b81526004016103779190611b42565b6001600160a01b038116811461055c57600080fd5b6000806040838503121561157057600080fd5b82359150602083013561158281611548565b809150509250929050565b60006020828403121561159f57600080fd5b813561110e81611548565b6000602082840312156115bc57600080fd5b5035919050565b600080600080606085870312156115d957600080fd5b8435935060208501356115eb81611548565b9250604085013567ffffffffffffffff8082111561160857600080fd5b818701915087601f83011261161c57600080fd5b81358181111561162b57600080fd5b88602082850101111561163d57600080fd5b95989497505060200194505050565b60006020828403121561165e57600080fd5b813567ffffffffffffffff81111561167557600080fd5b820160c0818503121561110e57600080fd5b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b6000602082840312156116e757600080fd5b5051919050565b634e487b7160e01b600052604160045260246000fd5b60405160c0810167ffffffffffffffff81118282101715611727576117276116ee565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715611756576117566116ee565b604052919050565b600067ffffffffffffffff821115611778576117786116ee565b5060051b60200190565b600060c0823603121561179457600080fd5b61179c611704565b823581526020808401356117af81611548565b8282015260408401356117c181611548565b80604084015250606084013560608301526080840135608083015260a084013567ffffffffffffffff8111156117f657600080fd5b840136601f82011261180757600080fd5b803561181a6118158261175e565b61172d565b81815260059190911b8201830190838101903683111561183957600080fd5b928401925b8284101561186057833561185181611548565b8252928401929084019061183e565b60a0860152509295945050505050565b6000806040838503121561188357600080fd5b825167ffffffffffffffff8082111561189b57600080fd5b9084019060c082870312156118af57600080fd5b6118b7611704565b825181526020808401516118ca81611548565b8282015260408401516118dc81611548565b80604084015250606084015160608301526080840151608083015260a08401518381111561190957600080fd5b80850194505087601f85011261191e57600080fd5b8351925061192e6118158461175e565b83815260059390931b8401810192818101908985111561194d57600080fd5b948201945b8486101561197457855161196581611548565b82529482019490820190611952565b60a0840152509590950151949694955050505050565b634e487b7160e01b600052601160045260246000fd5b818103818111156119b3576119b361198a565b92915050565b808201808211156119b3576119b361198a565b600081518084526020808501945080840160005b83811015611a055781516001600160a01b0316875295820195908201906001016119e0565b509495945050505050565b60005b83811015611a2b578181015183820152602001611a13565b50506000910152565b60008151808452611a4c816020860160208601611a10565b601f01601f19169290920160200192915050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c0840152611ac16101008401826119cc565b905060a0840151601f198483030160e0850152611ade8282611a34565b95945050505050565b600060208284031215611af957600080fd5b815161110e81611548565b600060208284031215611b1657600080fd5b8151801515811461110e57600080fd5b60008251611b38818460208701611a10565b9190910192915050565b60208152600061110e6020830184611a3456fea264697066735822122075ccf8349ad1fc713ae89e83f1921bdd69e3a291eb28b276044c94036e9d0d3e64736f6c63430008120033",
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
