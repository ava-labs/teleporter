// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20tokensource

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

// ERC20TokenSourceMetaData contains all meta data concerning the ERC20TokenSource contract.
var ERC20TokenSourceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeTokenDestinationAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20ContractAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurnTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferToDestination\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationBurnedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20ContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenDestinationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"originBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"transferToDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162002014380380620020148339810160408190526200003591620003f0565b83806001600160a01b038116620000b95760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa15801562000104573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200012a919062000444565b60015550620001393362000381565b50600160035582620001965760405162461bcd60e51b8152602060048201526030602482015260008051602062001ff483398151915260448201526f1bdb88189b1bd8dad8da185a5b88125160821b6064820152608401620000b0565b7302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001e9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200020f919062000444565b8303620002855760405162461bcd60e51b815260206004820152603460248201527f4552433230546f6b656e536f757263653a2063616e6e6f74206272696467652060448201527f776974682073616d6520626c6f636b636861696e0000000000000000000000006064820152608401620000b0565b60a08390526001600160a01b038216620002f75760405162461bcd60e51b8152602060048201526033602482015260008051602062001ff483398151915260448201527f6f6e20636f6e74726163742061646472657373000000000000000000000000006064820152608401620000b0565b6001600160a01b0380831660c05281166200036b5760405162461bcd60e51b815260206004820152602d60248201527f4552433230546f6b656e536f757263653a207a65726f20455243323020636f6e60448201526c7472616374206164647265737360981b6064820152608401620000b0565b6001600160a01b031660e052506200045e915050565b600280546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80516001600160a01b0381168114620003eb57600080fd5b919050565b600080600080608085870312156200040757600080fd5b6200041285620003d3565b9350602085015192506200042960408601620003d3565b91506200043960608601620003d3565b905092959194509250565b6000602082840312156200045757600080fd5b5051919050565b60805160a05160c05160e051611b08620004ec6000396000818161026d015281816104f40152818161058c0152818161063e01528181611356015261142a01526000818161022b01528181610604015261106401526000818161016e015281816105de0152610fe1015260008181610115015281816107e2015281816109f40152610c9a0152611b086000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80638da5cb5b116100a2578063c868efaa11610071578063c868efaa1461024d578063d2cc7a7014610260578063e486df1514610268578063f2fde38b1461028f578063fccc2813146102a257600080fd5b80638da5cb5b146101e857806397314297146101f9578063b6171f731461021c578063b8c9091a1461022657600080fd5b806355db3e9e116100de57806355db3e9e146101b15780635eb99514146101ba578063715018a6146101cd57806387a2edba146101d557600080fd5b80631a7f5bec146101105780632b0d8f181461015457806341d3014d146101695780634511243e1461019e575b600080fd5b6101377f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6101676101623660046115e7565b6102b2565b005b6101907f000000000000000000000000000000000000000000000000000000000000000081565b60405190815260200161014b565b6101676101ac3660046115e7565b6103a1565b61019060045481565b6101676101c836600461160b565b61048b565b61016761049f565b6101676101e3366004611624565b6104b3565b6002546001600160a01b0316610137565b61020c6102073660046115e7565b6107ac565b604051901515815260200161014b565b610190620186a081565b6101377f000000000000000000000000000000000000000000000000000000000000000081565b61016761025b3660046116bd565b6107ca565b600154610190565b6101377f000000000000000000000000000000000000000000000000000000000000000081565b61016761029d3660046115e7565b610972565b61013762010203600160981b0181565b6102ba6109e8565b6001600160a01b0381166102e95760405162461bcd60e51b81526004016102e090611746565b60405180910390fd5b6102f2816107ac565b156103555760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b60648201526084016102e0565b6001600160a01b038116600081815260208190526040808220805460ff19166001179055517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b6103a96109e8565b6001600160a01b0381166103cf5760405162461bcd60e51b81526004016102e090611746565b6103d8816107ac565b6104365760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b60648201526084016102e0565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152602081905260409020805460ff19169055565b6104936109e8565b61049c816109f0565b50565b6104a7610b90565b6104b16000610bea565b565b6104bb610c3c565b60006104c5610c95565b90506001600160a01b0386166104ed5760405162461bcd60e51b81526004016102e090611794565b60006105197f000000000000000000000000000000000000000000000000000000000000000087610d90565b90508481116105815760405162461bcd60e51b815260206004820152602e60248201527f4552433230546f6b656e536f757263653a20696e73756666696369656e74206160448201526d191a9d5cdd195908185b5bdd5b9d60921b60648201526084016102e0565b84156105b2576105b27f00000000000000000000000000000000000000000000000000000000000000008387610efa565b60006105be86836117f2565b90506000836001600160a01b031663624488506040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020018c8152508152602001620186a08152602001898980806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250505090825250604080516001600160a01b038f166020808301919091529181018890529101906060016040516020818303038152906040528152506040518263ffffffff1660e01b81526004016107049190611899565b6020604051808303816000875af1158015610723573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107479190611917565b905080896001600160a01b0316336001600160a01b03167f2b4e8f08417773e367064a6aea9ca2df303a60876676f70b6c3c5e66b314ca5a8560405161078f91815260200190565b60405180910390a4505050506107a56001600355565b5050505050565b6001600160a01b031660009081526020819052604090205460ff1690565b60015460405163260f846760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690634c1f08ce90602401602060405180830381865afa158015610831573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108559190611917565b10156108bc5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b60648201526084016102e0565b6108c5336107ac565b1561092b5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b60648201526084016102e0565b61096c848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610fdf92505050565b50505050565b61097a610b90565b6001600160a01b0381166109df5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102e0565b61049c81610bea565b6104b1610b90565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a749190611917565b60015490915081831115610ae45760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b60648201526084016102e0565b808311610b595760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e0060648201526084016102e0565b6001839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6002546001600160a01b031633146104b15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102e0565b600280546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600260035403610c8e5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016102e0565b6002600355565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cf6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1a9190611930565b9050610d25816107ac565b15610d8b5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b60648201526084016102e0565b919050565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015610dd9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dfd9190611917565b9050610e146001600160a01b0385163330866111da565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610e5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e7f9190611917565b9050818111610ee55760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016102e0565b610eef82826117f2565b925050505b92915050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015610f4b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f6f9190611917565b610f79919061194d565b6040516001600160a01b03851660248201526044810182905290915061096c90859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152611212565b7f000000000000000000000000000000000000000000000000000000000000000083146110625760405162461bcd60e51b815260206004820152602b60248201527f4552433230546f6b656e536f757263653a20696e76616c69642064657374696e60448201526a30ba34b7b71031b430b4b760a91b60648201526084016102e0565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316146110f15760405162461bcd60e51b815260206004820152602560248201527f4552433230546f6b656e536f757263653a20756e617574686f72697a6564207360448201526432b73232b960d91b60648201526084016102e0565b600080828060200190518101906111089190611976565b9092509050600082600181111561112157611121611a3d565b03611152576000808280602001905181019061113d9190611a53565b9150915061114b82826112e9565b50506107a5565b600182600181111561116657611166611a3d565b03611192576000818060200190518101906111819190611917565b905061118c81611380565b506107a5565b60405162461bcd60e51b815260206004820181905260248201527f4552433230546f6b656e536f757263653a20696e76616c696420616374696f6e60448201526064016102e0565b6040516001600160a01b038085166024830152831660448201526064810182905261096c9085906323b872dd60e01b90608401610fa8565b6000611267826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166113ab9092919063ffffffff16565b8051909150156112e457808060200190518101906112859190611a81565b6112e45760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016102e0565b505050565b6001600160a01b03821661130f5760405162461bcd60e51b81526004016102e090611794565b604080516001600160a01b0384168152602081018390527f55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407910160405180910390a161137c7f000000000000000000000000000000000000000000000000000000000000000083836113c2565b5050565b60045481111561049c5760006004548261139a91906117f2565b90506113a5816113f2565b50600455565b60606113ba8484600085611459565b949350505050565b6040516001600160a01b0383166024820152604481018290526112e490849063a9059cbb60e01b90606401610fa8565b6040518181527f2cd3fd70cd5a5d6d805e90d22741aa1a84590ace7cf01b244719558d266143829060200160405180910390a161049c7f000000000000000000000000000000000000000000000000000000000000000062010203600160981b01836113c2565b6060824710156114ba5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016102e0565b600080866001600160a01b031685876040516114d69190611aa3565b60006040518083038185875af1925050503d8060008114611513576040519150601f19603f3d011682016040523d82523d6000602084013e611518565b606091505b509150915061152987838387611534565b979650505050505050565b606083156115a357825160000361159c576001600160a01b0385163b61159c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102e0565b50816113ba565b6113ba83838151156115b85781518083602001fd5b8060405162461bcd60e51b81526004016102e09190611abf565b6001600160a01b038116811461049c57600080fd5b6000602082840312156115f957600080fd5b8135611604816115d2565b9392505050565b60006020828403121561161d57600080fd5b5035919050565b60008060008060006080868803121561163c57600080fd5b8535611647816115d2565b94506020860135935060408601359250606086013567ffffffffffffffff8082111561167257600080fd5b818801915088601f83011261168657600080fd5b81358181111561169557600080fd5b8960208260051b85010111156116aa57600080fd5b9699959850939650602001949392505050565b600080600080606085870312156116d357600080fd5b8435935060208501356116e5816115d2565b9250604085013567ffffffffffffffff8082111561170257600080fd5b818701915087601f83011261171657600080fd5b81358181111561172557600080fd5b88602082850101111561173757600080fd5b95989497505060200194505050565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b60208082526028908201527f4552433230546f6b656e536f757263653a207a65726f20726563697069656e74604082015267206164647265737360c01b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b81810381811115610ef457610ef46117dc565b600081518084526020808501945080840160005b8381101561183e5781516001600160a01b031687529582019590820190600101611819565b509495945050505050565b60005b8381101561186457818101518382015260200161184c565b50506000910152565b60008151808452611885816020860160208601611849565b601f01601f19169290920160200192915050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c08401526118fa610100840182611805565b905060a0840151601f198483030160e0850152610eef828261186d565b60006020828403121561192957600080fd5b5051919050565b60006020828403121561194257600080fd5b8151611604816115d2565b80820180821115610ef457610ef46117dc565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561198957600080fd5b82516002811061199857600080fd5b602084015190925067ffffffffffffffff808211156119b657600080fd5b818501915085601f8301126119ca57600080fd5b8151818111156119dc576119dc611960565b604051601f8201601f19908116603f01168101908382118183101715611a0457611a04611960565b81604052828152886020848701011115611a1d57600080fd5b611a2e836020830160208801611849565b80955050505050509250929050565b634e487b7160e01b600052602160045260246000fd5b60008060408385031215611a6657600080fd5b8251611a71816115d2565b6020939093015192949293505050565b600060208284031215611a9357600080fd5b8151801515811461160457600080fd5b60008251611ab5818460208701611849565b9190910192915050565b602081526000611604602083018461186d56fea26469706673582212208691c17188a871117336b9ec1b50b30b559af3fdd6c62d52672cf498f59acbb764736f6c634300081200334552433230546f6b656e536f757263653a207a65726f2064657374696e617469",
}

// ERC20TokenSourceABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenSourceMetaData.ABI instead.
var ERC20TokenSourceABI = ERC20TokenSourceMetaData.ABI

// ERC20TokenSourceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenSourceMetaData.Bin instead.
var ERC20TokenSourceBin = ERC20TokenSourceMetaData.Bin

// DeployERC20TokenSource deploys a new Ethereum contract, binding an instance of ERC20TokenSource to it.
func DeployERC20TokenSource(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterRegistryAddress common.Address, destinationBlockchainID_ [32]byte, nativeTokenDestinationAddress_ common.Address, erc20ContractAddress_ common.Address) (common.Address, *types.Transaction, *ERC20TokenSource, error) {
	parsed, err := ERC20TokenSourceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenSourceBin), backend, teleporterRegistryAddress, destinationBlockchainID_, nativeTokenDestinationAddress_, erc20ContractAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20TokenSource{ERC20TokenSourceCaller: ERC20TokenSourceCaller{contract: contract}, ERC20TokenSourceTransactor: ERC20TokenSourceTransactor{contract: contract}, ERC20TokenSourceFilterer: ERC20TokenSourceFilterer{contract: contract}}, nil
}

// ERC20TokenSource is an auto generated Go binding around an Ethereum contract.
type ERC20TokenSource struct {
	ERC20TokenSourceCaller     // Read-only binding to the contract
	ERC20TokenSourceTransactor // Write-only binding to the contract
	ERC20TokenSourceFilterer   // Log filterer for contract events
}

// ERC20TokenSourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenSourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenSourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenSourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenSourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenSourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenSourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenSourceSession struct {
	Contract     *ERC20TokenSource // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20TokenSourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenSourceCallerSession struct {
	Contract *ERC20TokenSourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC20TokenSourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenSourceTransactorSession struct {
	Contract     *ERC20TokenSourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC20TokenSourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenSourceRaw struct {
	Contract *ERC20TokenSource // Generic contract binding to access the raw methods on
}

// ERC20TokenSourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenSourceCallerRaw struct {
	Contract *ERC20TokenSourceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenSourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenSourceTransactorRaw struct {
	Contract *ERC20TokenSourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenSource creates a new instance of ERC20TokenSource, bound to a specific deployed contract.
func NewERC20TokenSource(address common.Address, backend bind.ContractBackend) (*ERC20TokenSource, error) {
	contract, err := bindERC20TokenSource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSource{ERC20TokenSourceCaller: ERC20TokenSourceCaller{contract: contract}, ERC20TokenSourceTransactor: ERC20TokenSourceTransactor{contract: contract}, ERC20TokenSourceFilterer: ERC20TokenSourceFilterer{contract: contract}}, nil
}

// NewERC20TokenSourceCaller creates a new read-only instance of ERC20TokenSource, bound to a specific deployed contract.
func NewERC20TokenSourceCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenSourceCaller, error) {
	contract, err := bindERC20TokenSource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceCaller{contract: contract}, nil
}

// NewERC20TokenSourceTransactor creates a new write-only instance of ERC20TokenSource, bound to a specific deployed contract.
func NewERC20TokenSourceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenSourceTransactor, error) {
	contract, err := bindERC20TokenSource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceTransactor{contract: contract}, nil
}

// NewERC20TokenSourceFilterer creates a new log filterer instance of ERC20TokenSource, bound to a specific deployed contract.
func NewERC20TokenSourceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenSourceFilterer, error) {
	contract, err := bindERC20TokenSource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceFilterer{contract: contract}, nil
}

// bindERC20TokenSource binds a generic wrapper to an already deployed contract.
func bindERC20TokenSource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TokenSourceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenSource *ERC20TokenSourceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenSource.Contract.ERC20TokenSourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenSource *ERC20TokenSourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.ERC20TokenSourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenSource *ERC20TokenSourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.ERC20TokenSourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenSource *ERC20TokenSourceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenSource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenSource *ERC20TokenSourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenSource *ERC20TokenSourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.contract.Transact(opts, method, params...)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCaller) BURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceSession) BURNADDRESS() (common.Address, error) {
	return _ERC20TokenSource.Contract.BURNADDRESS(&_ERC20TokenSource.CallOpts)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) BURNADDRESS() (common.Address, error) {
	return _ERC20TokenSource.Contract.BURNADDRESS(&_ERC20TokenSource.CallOpts)
}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceCaller) MINTNATIVETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "MINT_NATIVE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceSession) MINTNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenSource.Contract.MINTNATIVETOKENSREQUIREDGAS(&_ERC20TokenSource.CallOpts)
}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) MINTNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _ERC20TokenSource.Contract.MINTNATIVETOKENSREQUIREDGAS(&_ERC20TokenSource.CallOpts)
}

// DestinationBlockchainID is a free data retrieval call binding the contract method 0x41d3014d.
//
// Solidity: function destinationBlockchainID() view returns(bytes32)
func (_ERC20TokenSource *ERC20TokenSourceCaller) DestinationBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "destinationBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DestinationBlockchainID is a free data retrieval call binding the contract method 0x41d3014d.
//
// Solidity: function destinationBlockchainID() view returns(bytes32)
func (_ERC20TokenSource *ERC20TokenSourceSession) DestinationBlockchainID() ([32]byte, error) {
	return _ERC20TokenSource.Contract.DestinationBlockchainID(&_ERC20TokenSource.CallOpts)
}

// DestinationBlockchainID is a free data retrieval call binding the contract method 0x41d3014d.
//
// Solidity: function destinationBlockchainID() view returns(bytes32)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) DestinationBlockchainID() ([32]byte, error) {
	return _ERC20TokenSource.Contract.DestinationBlockchainID(&_ERC20TokenSource.CallOpts)
}

// DestinationBurnedTotal is a free data retrieval call binding the contract method 0x55db3e9e.
//
// Solidity: function destinationBurnedTotal() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceCaller) DestinationBurnedTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "destinationBurnedTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DestinationBurnedTotal is a free data retrieval call binding the contract method 0x55db3e9e.
//
// Solidity: function destinationBurnedTotal() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceSession) DestinationBurnedTotal() (*big.Int, error) {
	return _ERC20TokenSource.Contract.DestinationBurnedTotal(&_ERC20TokenSource.CallOpts)
}

// DestinationBurnedTotal is a free data retrieval call binding the contract method 0x55db3e9e.
//
// Solidity: function destinationBurnedTotal() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) DestinationBurnedTotal() (*big.Int, error) {
	return _ERC20TokenSource.Contract.DestinationBurnedTotal(&_ERC20TokenSource.CallOpts)
}

// Erc20ContractAddress is a free data retrieval call binding the contract method 0xe486df15.
//
// Solidity: function erc20ContractAddress() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCaller) Erc20ContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "erc20ContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20ContractAddress is a free data retrieval call binding the contract method 0xe486df15.
//
// Solidity: function erc20ContractAddress() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceSession) Erc20ContractAddress() (common.Address, error) {
	return _ERC20TokenSource.Contract.Erc20ContractAddress(&_ERC20TokenSource.CallOpts)
}

// Erc20ContractAddress is a free data retrieval call binding the contract method 0xe486df15.
//
// Solidity: function erc20ContractAddress() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) Erc20ContractAddress() (common.Address, error) {
	return _ERC20TokenSource.Contract.Erc20ContractAddress(&_ERC20TokenSource.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ERC20TokenSource.Contract.GetMinTeleporterVersion(&_ERC20TokenSource.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _ERC20TokenSource.Contract.GetMinTeleporterVersion(&_ERC20TokenSource.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20TokenSource *ERC20TokenSourceCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20TokenSource *ERC20TokenSourceSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _ERC20TokenSource.Contract.IsTeleporterAddressPaused(&_ERC20TokenSource.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _ERC20TokenSource.Contract.IsTeleporterAddressPaused(&_ERC20TokenSource.CallOpts, teleporterAddress)
}

// NativeTokenDestinationAddress is a free data retrieval call binding the contract method 0xb8c9091a.
//
// Solidity: function nativeTokenDestinationAddress() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCaller) NativeTokenDestinationAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "nativeTokenDestinationAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeTokenDestinationAddress is a free data retrieval call binding the contract method 0xb8c9091a.
//
// Solidity: function nativeTokenDestinationAddress() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceSession) NativeTokenDestinationAddress() (common.Address, error) {
	return _ERC20TokenSource.Contract.NativeTokenDestinationAddress(&_ERC20TokenSource.CallOpts)
}

// NativeTokenDestinationAddress is a free data retrieval call binding the contract method 0xb8c9091a.
//
// Solidity: function nativeTokenDestinationAddress() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) NativeTokenDestinationAddress() (common.Address, error) {
	return _ERC20TokenSource.Contract.NativeTokenDestinationAddress(&_ERC20TokenSource.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceSession) Owner() (common.Address, error) {
	return _ERC20TokenSource.Contract.Owner(&_ERC20TokenSource.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) Owner() (common.Address, error) {
	return _ERC20TokenSource.Contract.Owner(&_ERC20TokenSource.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceSession) TeleporterRegistry() (common.Address, error) {
	return _ERC20TokenSource.Contract.TeleporterRegistry(&_ERC20TokenSource.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) TeleporterRegistry() (common.Address, error) {
	return _ERC20TokenSource.Contract.TeleporterRegistry(&_ERC20TokenSource.CallOpts)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenSource *ERC20TokenSourceSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.PauseTeleporterAddress(&_ERC20TokenSource.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.PauseTeleporterAddress(&_ERC20TokenSource.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, originBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenSource.contract.Transact(opts, "receiveTeleporterMessage", originBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenSource *ERC20TokenSourceSession) ReceiveTeleporterMessage(originBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.ReceiveTeleporterMessage(&_ERC20TokenSource.TransactOpts, originBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 originBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactorSession) ReceiveTeleporterMessage(originBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.ReceiveTeleporterMessage(&_ERC20TokenSource.TransactOpts, originBlockchainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenSource.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenSource *ERC20TokenSourceSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.RenounceOwnership(&_ERC20TokenSource.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.RenounceOwnership(&_ERC20TokenSource.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenSource *ERC20TokenSourceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.TransferOwnership(&_ERC20TokenSource.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.TransferOwnership(&_ERC20TokenSource.TransactOpts, newOwner)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0x87a2edba.
//
// Solidity: function transferToDestination(address recipient, uint256 totalAmount, uint256 feeAmount, address[] allowedRelayerAddresses) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactor) TransferToDestination(opts *bind.TransactOpts, recipient common.Address, totalAmount *big.Int, feeAmount *big.Int, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.contract.Transact(opts, "transferToDestination", recipient, totalAmount, feeAmount, allowedRelayerAddresses)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0x87a2edba.
//
// Solidity: function transferToDestination(address recipient, uint256 totalAmount, uint256 feeAmount, address[] allowedRelayerAddresses) returns()
func (_ERC20TokenSource *ERC20TokenSourceSession) TransferToDestination(recipient common.Address, totalAmount *big.Int, feeAmount *big.Int, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.TransferToDestination(&_ERC20TokenSource.TransactOpts, recipient, totalAmount, feeAmount, allowedRelayerAddresses)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0x87a2edba.
//
// Solidity: function transferToDestination(address recipient, uint256 totalAmount, uint256 feeAmount, address[] allowedRelayerAddresses) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactorSession) TransferToDestination(recipient common.Address, totalAmount *big.Int, feeAmount *big.Int, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.TransferToDestination(&_ERC20TokenSource.TransactOpts, recipient, totalAmount, feeAmount, allowedRelayerAddresses)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenSource *ERC20TokenSourceSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.UnpauseTeleporterAddress(&_ERC20TokenSource.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.UnpauseTeleporterAddress(&_ERC20TokenSource.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _ERC20TokenSource.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20TokenSource *ERC20TokenSourceSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.UpdateMinTeleporterVersion(&_ERC20TokenSource.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.UpdateMinTeleporterVersion(&_ERC20TokenSource.TransactOpts, version)
}

// ERC20TokenSourceBurnTokensIterator is returned from FilterBurnTokens and is used to iterate over the raw logs and unpacked data for BurnTokens events raised by the ERC20TokenSource contract.
type ERC20TokenSourceBurnTokensIterator struct {
	Event *ERC20TokenSourceBurnTokens // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSourceBurnTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSourceBurnTokens)
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
		it.Event = new(ERC20TokenSourceBurnTokens)
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
func (it *ERC20TokenSourceBurnTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSourceBurnTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSourceBurnTokens represents a BurnTokens event raised by the ERC20TokenSource contract.
type ERC20TokenSourceBurnTokens struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnTokens is a free log retrieval operation binding the contract event 0x2cd3fd70cd5a5d6d805e90d22741aa1a84590ace7cf01b244719558d26614382.
//
// Solidity: event BurnTokens(uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) FilterBurnTokens(opts *bind.FilterOpts) (*ERC20TokenSourceBurnTokensIterator, error) {

	logs, sub, err := _ERC20TokenSource.contract.FilterLogs(opts, "BurnTokens")
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceBurnTokensIterator{contract: _ERC20TokenSource.contract, event: "BurnTokens", logs: logs, sub: sub}, nil
}

// WatchBurnTokens is a free log subscription operation binding the contract event 0x2cd3fd70cd5a5d6d805e90d22741aa1a84590ace7cf01b244719558d26614382.
//
// Solidity: event BurnTokens(uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) WatchBurnTokens(opts *bind.WatchOpts, sink chan<- *ERC20TokenSourceBurnTokens) (event.Subscription, error) {

	logs, sub, err := _ERC20TokenSource.contract.WatchLogs(opts, "BurnTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSourceBurnTokens)
				if err := _ERC20TokenSource.contract.UnpackLog(event, "BurnTokens", log); err != nil {
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

// ParseBurnTokens is a log parse operation binding the contract event 0x2cd3fd70cd5a5d6d805e90d22741aa1a84590ace7cf01b244719558d26614382.
//
// Solidity: event BurnTokens(uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) ParseBurnTokens(log types.Log) (*ERC20TokenSourceBurnTokens, error) {
	event := new(ERC20TokenSourceBurnTokens)
	if err := _ERC20TokenSource.contract.UnpackLog(event, "BurnTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenSourceMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the ERC20TokenSource contract.
type ERC20TokenSourceMinTeleporterVersionUpdatedIterator struct {
	Event *ERC20TokenSourceMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSourceMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSourceMinTeleporterVersionUpdated)
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
		it.Event = new(ERC20TokenSourceMinTeleporterVersionUpdated)
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
func (it *ERC20TokenSourceMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSourceMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSourceMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the ERC20TokenSource contract.
type ERC20TokenSourceMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*ERC20TokenSourceMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceMinTeleporterVersionUpdatedIterator{contract: _ERC20TokenSource.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *ERC20TokenSourceMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSourceMinTeleporterVersionUpdated)
				if err := _ERC20TokenSource.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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
func (_ERC20TokenSource *ERC20TokenSourceFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*ERC20TokenSourceMinTeleporterVersionUpdated, error) {
	event := new(ERC20TokenSourceMinTeleporterVersionUpdated)
	if err := _ERC20TokenSource.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenSourceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20TokenSource contract.
type ERC20TokenSourceOwnershipTransferredIterator struct {
	Event *ERC20TokenSourceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSourceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSourceOwnershipTransferred)
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
		it.Event = new(ERC20TokenSourceOwnershipTransferred)
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
func (it *ERC20TokenSourceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSourceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSourceOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20TokenSource contract.
type ERC20TokenSourceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20TokenSourceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceOwnershipTransferredIterator{contract: _ERC20TokenSource.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20TokenSourceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSourceOwnershipTransferred)
				if err := _ERC20TokenSource.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC20TokenSource *ERC20TokenSourceFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20TokenSourceOwnershipTransferred, error) {
	event := new(ERC20TokenSourceOwnershipTransferred)
	if err := _ERC20TokenSource.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenSourceTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the ERC20TokenSource contract.
type ERC20TokenSourceTeleporterAddressPausedIterator struct {
	Event *ERC20TokenSourceTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSourceTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSourceTeleporterAddressPaused)
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
		it.Event = new(ERC20TokenSourceTeleporterAddressPaused)
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
func (it *ERC20TokenSourceTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSourceTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSourceTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the ERC20TokenSource contract.
type ERC20TokenSourceTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*ERC20TokenSourceTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceTeleporterAddressPausedIterator{contract: _ERC20TokenSource.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *ERC20TokenSourceTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSourceTeleporterAddressPaused)
				if err := _ERC20TokenSource.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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
func (_ERC20TokenSource *ERC20TokenSourceFilterer) ParseTeleporterAddressPaused(log types.Log) (*ERC20TokenSourceTeleporterAddressPaused, error) {
	event := new(ERC20TokenSourceTeleporterAddressPaused)
	if err := _ERC20TokenSource.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenSourceTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the ERC20TokenSource contract.
type ERC20TokenSourceTeleporterAddressUnpausedIterator struct {
	Event *ERC20TokenSourceTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSourceTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSourceTeleporterAddressUnpaused)
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
		it.Event = new(ERC20TokenSourceTeleporterAddressUnpaused)
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
func (it *ERC20TokenSourceTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSourceTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSourceTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the ERC20TokenSource contract.
type ERC20TokenSourceTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*ERC20TokenSourceTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceTeleporterAddressUnpausedIterator{contract: _ERC20TokenSource.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20TokenSourceTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSourceTeleporterAddressUnpaused)
				if err := _ERC20TokenSource.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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
func (_ERC20TokenSource *ERC20TokenSourceFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*ERC20TokenSourceTeleporterAddressUnpaused, error) {
	event := new(ERC20TokenSourceTeleporterAddressUnpaused)
	if err := _ERC20TokenSource.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenSourceTransferToDestinationIterator is returned from FilterTransferToDestination and is used to iterate over the raw logs and unpacked data for TransferToDestination events raised by the ERC20TokenSource contract.
type ERC20TokenSourceTransferToDestinationIterator struct {
	Event *ERC20TokenSourceTransferToDestination // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSourceTransferToDestinationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSourceTransferToDestination)
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
		it.Event = new(ERC20TokenSourceTransferToDestination)
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
func (it *ERC20TokenSourceTransferToDestinationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSourceTransferToDestinationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSourceTransferToDestination represents a TransferToDestination event raised by the ERC20TokenSource contract.
type ERC20TokenSourceTransferToDestination struct {
	Sender              common.Address
	Recipient           common.Address
	TeleporterMessageID *big.Int
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTransferToDestination is a free log retrieval operation binding the contract event 0x2b4e8f08417773e367064a6aea9ca2df303a60876676f70b6c3c5e66b314ca5a.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, uint256 indexed teleporterMessageID, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) FilterTransferToDestination(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, teleporterMessageID []*big.Int) (*ERC20TokenSourceTransferToDestinationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.FilterLogs(opts, "TransferToDestination", senderRule, recipientRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceTransferToDestinationIterator{contract: _ERC20TokenSource.contract, event: "TransferToDestination", logs: logs, sub: sub}, nil
}

// WatchTransferToDestination is a free log subscription operation binding the contract event 0x2b4e8f08417773e367064a6aea9ca2df303a60876676f70b6c3c5e66b314ca5a.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, uint256 indexed teleporterMessageID, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) WatchTransferToDestination(opts *bind.WatchOpts, sink chan<- *ERC20TokenSourceTransferToDestination, sender []common.Address, recipient []common.Address, teleporterMessageID []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.WatchLogs(opts, "TransferToDestination", senderRule, recipientRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSourceTransferToDestination)
				if err := _ERC20TokenSource.contract.UnpackLog(event, "TransferToDestination", log); err != nil {
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

// ParseTransferToDestination is a log parse operation binding the contract event 0x2b4e8f08417773e367064a6aea9ca2df303a60876676f70b6c3c5e66b314ca5a.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, uint256 indexed teleporterMessageID, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) ParseTransferToDestination(log types.Log) (*ERC20TokenSourceTransferToDestination, error) {
	event := new(ERC20TokenSourceTransferToDestination)
	if err := _ERC20TokenSource.contract.UnpackLog(event, "TransferToDestination", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenSourceUnlockTokensIterator is returned from FilterUnlockTokens and is used to iterate over the raw logs and unpacked data for UnlockTokens events raised by the ERC20TokenSource contract.
type ERC20TokenSourceUnlockTokensIterator struct {
	Event *ERC20TokenSourceUnlockTokens // Event containing the contract specifics and raw log

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
func (it *ERC20TokenSourceUnlockTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenSourceUnlockTokens)
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
		it.Event = new(ERC20TokenSourceUnlockTokens)
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
func (it *ERC20TokenSourceUnlockTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenSourceUnlockTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenSourceUnlockTokens represents a UnlockTokens event raised by the ERC20TokenSource contract.
type ERC20TokenSourceUnlockTokens struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlockTokens is a free log retrieval operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) FilterUnlockTokens(opts *bind.FilterOpts) (*ERC20TokenSourceUnlockTokensIterator, error) {

	logs, sub, err := _ERC20TokenSource.contract.FilterLogs(opts, "UnlockTokens")
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceUnlockTokensIterator{contract: _ERC20TokenSource.contract, event: "UnlockTokens", logs: logs, sub: sub}, nil
}

// WatchUnlockTokens is a free log subscription operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) WatchUnlockTokens(opts *bind.WatchOpts, sink chan<- *ERC20TokenSourceUnlockTokens) (event.Subscription, error) {

	logs, sub, err := _ERC20TokenSource.contract.WatchLogs(opts, "UnlockTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenSourceUnlockTokens)
				if err := _ERC20TokenSource.contract.UnpackLog(event, "UnlockTokens", log); err != nil {
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

// ParseUnlockTokens is a log parse operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) ParseUnlockTokens(log types.Log) (*ERC20TokenSourceUnlockTokens, error) {
	event := new(ERC20TokenSourceUnlockTokens)
	if err := _ERC20TokenSource.contract.UnpackLog(event, "UnlockTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
