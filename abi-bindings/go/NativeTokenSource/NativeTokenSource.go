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
	RequiredGasLimit         *big.Int
}

// NativeTokenSourceMetaData contains all meta data concerning the NativeTokenSource contract.
var NativeTokenSourceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SendTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"blockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"}],\"name\":\"bridgedBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"primaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structSendTokensInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIWrappedNativeToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620022fc380380620022fc8339810160408190526200003591620003a5565b60016000558282828282816001600160a01b038116620000c25760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001339190620003ef565b60025550620001423362000259565b6200014d81620002ab565b50507302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001a2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001c89190620003ef565b60a0526001600160a01b038116620002395760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401620000b9565b6001600160a01b0390811660c0529290921660e052506200040992505050565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b620002b56200032a565b6001600160a01b0381166200031c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401620000b9565b620003278162000259565b50565b6003546001600160a01b03163314620003865760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401620000b9565b565b80516001600160a01b0381168114620003a057600080fd5b919050565b600080600060608486031215620003bb57600080fd5b620003c68462000388565b9250620003d66020850162000388565b9150620003e66040850162000388565b90509250925092565b6000602082840312156200040257600080fd5b5051919050565b60805160a05160c05160e051611e7b62000481600039600081816103e4015281816111e0015261134801526000818160fc015281816103270152610c4b01526000818161037b0152818161095c01526110d10152600081816101ff0152818161066801528181610dae015261159b0152611e7b6000f3fe6080604052600436106100ec5760003560e01c80638da5cb5b1161008a578063d127dc9b11610059578063d127dc9b14610369578063d2cc7a701461039d578063f2fde38b146103b2578063fc0c546a146103d257600080fd5b80638da5cb5b146102ae57806397314297146102cc578063b8df0dea14610315578063c868efaa1461034957600080fd5b80632b0d8f18116100c65780632b0d8f18146102395780634511243e146102595780635eb9951414610279578063715018a61461029957600080fd5b806302ee3e9c1461018f578063146dfd3c146101da5780631a7f5bec146101ed57600080fd5b3661018a57336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101885760405162461bcd60e51b815260206004820152603160248201527f4e6174697665546f6b656e536f757263653a20696e76616c69642072656365696044820152703b32903830bcb0b136329039b2b73232b960791b60648201526084015b60405180910390fd5b005b600080fd5b34801561019b57600080fd5b506101c76101aa3660046119de565b600460209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b6101886101e8366004611a0e565b610406565b3480156101f957600080fd5b506102217f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101d1565b34801561024557600080fd5b50610188610254366004611a26565b610435565b34801561026557600080fd5b50610188610274366004611a26565b610531565b34801561028557600080fd5b50610188610294366004611a43565b61062e565b3480156102a557600080fd5b5061018861063f565b3480156102ba57600080fd5b506003546001600160a01b0316610221565b3480156102d857600080fd5b506103056102e7366004611a26565b6001600160a01b031660009081526001602052604090205460ff1690565b60405190151581526020016101d1565b34801561032157600080fd5b506102217f000000000000000000000000000000000000000000000000000000000000000081565b34801561035557600080fd5b50610188610364366004611a5c565b610653565b34801561037557600080fd5b506101c77f000000000000000000000000000000000000000000000000000000000000000081565b3480156103a957600080fd5b506002546101c7565b3480156103be57600080fd5b506101886103cd366004611a26565b61081d565b3480156103de57600080fd5b506102217f000000000000000000000000000000000000000000000000000000000000000081565b61040e610893565b61042861042036839003830183611b32565b3460006108ec565b6104326001600055565b50565b61043d610da2565b6001600160a01b0381166104635760405162461bcd60e51b815260040161017f90611ba0565b6001600160a01b03811660009081526001602052604090205460ff16156104e25760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b606482015260840161017f565b6001600160a01b0381166000818152600160208190526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b610539610da2565b6001600160a01b03811661055f5760405162461bcd60e51b815260040161017f90611ba0565b6001600160a01b03811660009081526001602052604090205460ff166105d95760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b606482015260840161017f565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600160205260409020805460ff19169055565b610636610da2565b61043281610daa565b610647610f4a565b6106516000610fa4565b565b61065b610893565b6002546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa1580156106d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f69190611bee565b101561075d5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b606482015260840161017f565b610766336102e7565b156107cc5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b606482015260840161017f565b61080d848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610ff692505050565b6108176001600055565b50505050565b610825610f4a565b6001600160a01b03811661088a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161017f565b61043281610fa4565b6002600054036108e55760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161017f565b6002600055565b82516109585760405162461bcd60e51b815260206004820152603560248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f20646573746044820152741a5b985d1a5bdb88189b1bd8dad8da185a5b881251605a1b606482015260840161017f565b82517f000000000000000000000000000000000000000000000000000000000000000090036109e45760405162461bcd60e51b815260206004820152603260248201527f54656c65706f72746572546f6b656e536f757263653a2063616e6e6f7420627260448201527134b233b2903a379039b0b6b29031b430b4b760711b606482015260840161017f565b60208301516001600160a01b0316610a5d5760405162461bcd60e51b815260206004820152603660248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f2064657374604482015275696e6174696f6e20627269646765206164647265737360501b606482015260840161017f565b60408301516001600160a01b0316610acd5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572546f6b656e536f757263653a207a65726f207265636960448201526c7069656e74206164647265737360981b606482015260840161017f565b608083015115610b355760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572546f6b656e536f757263653a206e6f6e2d7a65726f2060448201526c7365636f6e646172792066656560981b606482015260840161017f565b80610b4657610b43826111d1565b91505b82606001518211610bbf5760405162461bcd60e51b815260206004820152603860248201527f54656c65706f72746572546f6b656e536f757263653a20696e7375666669636960448201527f656e7420616d6f756e7420746f20636f76657220666565730000000000000000606482015260840161017f565b6060830151610bce9083611c1d565b83516000908152600460209081526040808320828801516001600160a01b03168452909152812080549294508492909190610c0a908490611c30565b925050819055506000610d1a6040518060c001604052808660000151815260200186602001516001600160a01b0316815260200160405180604001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602001886060015181525081526020018660a001518152602001600067ffffffffffffffff811115610ca957610ca9611ae5565b604051908082528060200260200182016040528015610cd2578160200160208202803683370190505b508152602001866040015186604051602001610d039291906001600160a01b03929092168252602082015260400190565b60405160208183030381529060405281525061120c565b60408051865181526020808801516001600160a01b0390811691830191909152878301511681830152606080880151908201526080808801519082015260a0808801519082015260c081018690529051919250339183917fdc72eade375f64388998c651c5137da74fbe64aec18f766dea3da18f154ab0a0919081900360e00190a350505050565b610651610f4a565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e0a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e2e9190611bee565b60025490915081831115610e9e5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b606482015260840161017f565b808311610f135760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e00606482015260840161017f565b6002839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6003546001600160a01b031633146106515760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161017f565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000808280602001905181019061100d9190611c43565b60008781526004602090815260408083206001600160a01b038a1684529091529020549193509150818110156110a05760405162461bcd60e51b815260206004820152603260248201527f54656c65706f72746572546f6b656e536f757263653a20696e73756666696369604482015271656e74206272696467652062616c616e636560701b606482015260840161017f565b6110aa8282611c1d565b60008781526004602090815260408083206001600160a01b038a16845290915290205582517f000000000000000000000000000000000000000000000000000000000000000090036111c55760208301516001600160a01b031630146111685760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f72746572546f6b656e536f757263653a20696e76616c6964206260448201526c7269646765206164647265737360981b606482015260840161017f565b82604001516001600160a01b03167f680f2e4f4032ebf1774e8cdbaddcb1b617a5a606411c8ca96257ada338d3833c836040516111a791815260200190565b60405180910390a26111bd836040015183611332565b505050505050565b6111bd838360016108ec565b60006112066001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016836113e7565b92915050565b600080611217611596565b604084015160200151909150156112bc576040830151516001600160a01b03166112995760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b606482015260840161017f565b6040830151602081015190516112bc916001600160a01b039091169083906116aa565b604051630624488560e41b81526001600160a01b038216906362448850906112e8908690600401611d59565b6020604051808303816000875af1158015611307573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061132b9190611bee565b9392505050565b604051632e1a7d4d60e01b8152600481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d90602401600060405180830381600087803b15801561139457600080fd5b505af11580156113a8573d6000803e3d6000fd5b50506040516001600160a01b038516925083156108fc02915083906000818181858888f193505050501580156113e2573d6000803e3d6000fd5b505050565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015611430573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114549190611bee565b9050836001600160a01b031663d0e30db0846040518263ffffffff1660e01b81526004016000604051808303818588803b15801561149157600080fd5b505af11580156114a5573d6000803e3d6000fd5b50506040516370a0823160e01b8152306004820152600093506001600160a01b03881692506370a082319150602401602060405180830381865afa1580156114f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115159190611bee565b90508181116115835760405162461bcd60e51b815260206004820152603460248201527f53616665577261707065644e6174697665546f6b656e4465706f7369743a2062604482015273185b185b98d9481b9bdd081a5b98dc99585cd95960621b606482015260840161017f565b61158d8282611c1d565b95945050505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061161b9190611dd7565b905061163f816001600160a01b031660009081526001602052604090205460ff1690565b156116a55760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b606482015260840161017f565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa1580156116fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061171f9190611bee565b6117299190611c30565b604080516001600160a01b03868116602483015260448083018590528351808403909101815260649092018352602080830180516001600160e01b031663095ea7b360e01b17905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564908401529293506108179287926000916117bc91908516908490611839565b8051909150156113e257808060200190518101906117da9190611df4565b6113e25760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161017f565b60606118488484600085611850565b949350505050565b6060824710156118b15760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161017f565b600080866001600160a01b031685876040516118cd9190611e16565b60006040518083038185875af1925050503d806000811461190a576040519150601f19603f3d011682016040523d82523d6000602084013e61190f565b606091505b50915091506119208783838761192b565b979650505050505050565b6060831561199a578251600003611993576001600160a01b0385163b6119935760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161017f565b5081611848565b61184883838151156119af5781518083602001fd5b8060405162461bcd60e51b815260040161017f9190611e32565b6001600160a01b038116811461043257600080fd5b600080604083850312156119f157600080fd5b823591506020830135611a03816119c9565b809150509250929050565b600060c08284031215611a2057600080fd5b50919050565b600060208284031215611a3857600080fd5b813561132b816119c9565b600060208284031215611a5557600080fd5b5035919050565b60008060008060608587031215611a7257600080fd5b843593506020850135611a84816119c9565b9250604085013567ffffffffffffffff80821115611aa157600080fd5b818701915087601f830112611ab557600080fd5b813581811115611ac457600080fd5b886020828501011115611ad657600080fd5b95989497505060200194505050565b634e487b7160e01b600052604160045260246000fd5b60405160c0810167ffffffffffffffff81118282101715611b2c57634e487b7160e01b600052604160045260246000fd5b60405290565b600060c08284031215611b4457600080fd5b611b4c611afb565b823581526020830135611b5e816119c9565b60208201526040830135611b71816119c9565b80604083015250606083013560608201526080830135608082015260a083013560a08201528091505092915050565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b600060208284031215611c0057600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b8181038181111561120657611206611c07565b8082018082111561120657611206611c07565b60008082840360e0811215611c5757600080fd5b60c0811215611c6557600080fd5b50611c6e611afb565b835181526020840151611c80816119c9565b60208201526040840151611c93816119c9565b6040820152606084810151908201526080808501519082015260a0808501519082015260c09093015192949293505050565b600081518084526020808501945080840160005b83811015611cfe5781516001600160a01b031687529582019590820190600101611cd9565b509495945050505050565b60005b83811015611d24578181015183820152602001611d0c565b50506000910152565b60008151808452611d45816020860160208601611d09565b601f01601f19169290920160200192915050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c0840152611dba610100840182611cc5565b905060a0840151601f198483030160e085015261158d8282611d2d565b600060208284031215611de957600080fd5b815161132b816119c9565b600060208284031215611e0657600080fd5b8151801515811461132b57600080fd5b60008251611e28818460208701611d09565b9190910192915050565b60208152600061132b6020830184611d2d56fea26469706673582212202212b2105d688a2c822cf40904e9be1e41ce1c6a5c8064f11ca5ff854a76b16c64736f6c63430008120033",
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

// Send is a paid mutator transaction binding the contract method 0x146dfd3c.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,uint256) input) payable returns()
func (_NativeTokenSource *NativeTokenSourceTransactor) Send(opts *bind.TransactOpts, input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenSource.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0x146dfd3c.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,uint256) input) payable returns()
func (_NativeTokenSource *NativeTokenSourceSession) Send(input SendTokensInput) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.Send(&_NativeTokenSource.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0x146dfd3c.
//
// Solidity: function send((bytes32,address,address,uint256,uint256,uint256) input) payable returns()
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
	Input               SendTokensInput
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSendTokens is a free log retrieval operation binding the contract event 0xdc72eade375f64388998c651c5137da74fbe64aec18f766dea3da18f154ab0a0.
//
// Solidity: event SendTokens(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,uint256,uint256,uint256) input, uint256 amount)
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

// WatchSendTokens is a free log subscription operation binding the contract event 0xdc72eade375f64388998c651c5137da74fbe64aec18f766dea3da18f154ab0a0.
//
// Solidity: event SendTokens(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,uint256,uint256,uint256) input, uint256 amount)
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

// ParseSendTokens is a log parse operation binding the contract event 0xdc72eade375f64388998c651c5137da74fbe64aec18f766dea3da18f154ab0a0.
//
// Solidity: event SendTokens(bytes32 indexed teleporterMessageID, address indexed sender, (bytes32,address,address,uint256,uint256,uint256) input, uint256 amount)
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

// NativeTokenSourceWithdrawTokensIterator is returned from FilterWithdrawTokens and is used to iterate over the raw logs and unpacked data for WithdrawTokens events raised by the NativeTokenSource contract.
type NativeTokenSourceWithdrawTokensIterator struct {
	Event *NativeTokenSourceWithdrawTokens // Event containing the contract specifics and raw log

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
func (it *NativeTokenSourceWithdrawTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSourceWithdrawTokens)
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
		it.Event = new(NativeTokenSourceWithdrawTokens)
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
func (it *NativeTokenSourceWithdrawTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSourceWithdrawTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSourceWithdrawTokens represents a WithdrawTokens event raised by the NativeTokenSource contract.
type NativeTokenSourceWithdrawTokens struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTokens is a free log retrieval operation binding the contract event 0x680f2e4f4032ebf1774e8cdbaddcb1b617a5a606411c8ca96257ada338d3833c.
//
// Solidity: event WithdrawTokens(address indexed recipient, uint256 amount)
func (_NativeTokenSource *NativeTokenSourceFilterer) FilterWithdrawTokens(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenSourceWithdrawTokensIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenSource.contract.FilterLogs(opts, "WithdrawTokens", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceWithdrawTokensIterator{contract: _NativeTokenSource.contract, event: "WithdrawTokens", logs: logs, sub: sub}, nil
}

// WatchWithdrawTokens is a free log subscription operation binding the contract event 0x680f2e4f4032ebf1774e8cdbaddcb1b617a5a606411c8ca96257ada338d3833c.
//
// Solidity: event WithdrawTokens(address indexed recipient, uint256 amount)
func (_NativeTokenSource *NativeTokenSourceFilterer) WatchWithdrawTokens(opts *bind.WatchOpts, sink chan<- *NativeTokenSourceWithdrawTokens, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenSource.contract.WatchLogs(opts, "WithdrawTokens", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSourceWithdrawTokens)
				if err := _NativeTokenSource.contract.UnpackLog(event, "WithdrawTokens", log); err != nil {
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
func (_NativeTokenSource *NativeTokenSourceFilterer) ParseWithdrawTokens(log types.Log) (*NativeTokenSourceWithdrawTokens, error) {
	event := new(NativeTokenSourceWithdrawTokens)
	if err := _NativeTokenSource.contract.UnpackLog(event, "WithdrawTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
