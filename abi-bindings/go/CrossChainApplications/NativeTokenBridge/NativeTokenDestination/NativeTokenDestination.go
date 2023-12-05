// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokendestination

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

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// NativeTokenDestinationMetaData contains all meta data concerning the NativeTokenDestination contract.
var NativeTokenDestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeTokenSourceAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialReserveImbalance_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"CollateralAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NativeTokensMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAddressBalance\",\"type\":\"uint256\"}],\"name\":\"ReportTotalBurnedTxFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferToSource\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURNED_TX_FEES_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_FOR_TRANSFER_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORT_BURNED_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialReserveImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenSourceAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"senderBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"reportTotalBurnedTxFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"transferToSource\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x61012060405273020000000000000000000000000000000000000160805260006002553480156200002f57600080fd5b5060405162001bae38038062001bae833981016040819052620000529162000355565b60016000556001600160a01b038416620000d95760405162461bcd60e51b815260206004820152603860248201527f4e6174697665546f6b656e44657374696e6174696f6e3a207a65726f2054656c60448201527f65706f727465724d657373656e6765722061646472657373000000000000000060648201526084015b60405180910390fd5b6001600160a01b03841661010052826200013f5760405162461bcd60e51b8152602060048201526031602482015260008051602062001b8e8339815191526044820152701c98d948189b1bd8dad8da185a5b881251607a1b6064820152608401620000d0565b7302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000192573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b891906200039e565b83036200022e5760405162461bcd60e51b815260206004820152603a60248201527f4e6174697665546f6b656e44657374696e6174696f6e3a2063616e6e6f74206260448201527f726964676520776974682073616d6520626c6f636b636861696e0000000000006064820152608401620000d0565b60a08390526001600160a01b038216620002a05760405162461bcd60e51b8152602060048201526034602482015260008051602062001b8e83398151915260448201527f72636520636f6e747261637420616464726573730000000000000000000000006064820152608401620000d0565b6001600160a01b03821660c0526000819003620003265760405162461bcd60e51b815260206004820152603660248201527f4e6174697665546f6b656e44657374696e6174696f6e3a207a65726f20696e6960448201527f7469616c207265736572766520696d62616c616e6365000000000000000000006064820152608401620000d0565b60e081905260015550620003b8915050565b80516001600160a01b03811681146200035057600080fd5b919050565b600080600080608085870312156200036c57600080fd5b620003778562000338565b9350602085015192506200038e6040860162000338565b6060959095015193969295505050565b600060208284031215620003b157600080fd5b5051919050565b60805160a05160c05160e0516101005161173d62000451600039600081816102440152818161040c015281816106a50152818161070201526109120152600081816102100152818161030401526103b90152600081816101b10152818161046c015281816107620152610a3501526000818161013c015281816104460152818161073c01526109b101526000610c51015261173d6000f3fe6080604052600436106100e75760003560e01c80638ac7dd201161008a578063ab28523011610059578063ab28523014610297578063c452165e146102ae578063c868efaa146102c6578063d30951261461029757600080fd5b80638ac7dd20146101fe5780639b3e580314610232578063a2309ff814610266578063a2a950171461027c57600080fd5b80633a94fe51116100c65780633a94fe511461015e57806349e3284e146101805780635d93f9af1461019f57806375846562146101eb57600080fd5b8062d872ae146100ec57806318160ddd1461011557806329b7b3fd1461012a575b600080fd5b3480156100f857600080fd5b5061010260015481565b6040519081526020015b60405180910390f35b34801561012157600080fd5b506101026102e6565b34801561013657600080fd5b506101027f000000000000000000000000000000000000000000000000000000000000000081565b34801561016a57600080fd5b5061017e6101793660046112bb565b6103f5565b005b34801561018c57600080fd5b506001546040519015815260200161010c565b3480156101ab57600080fd5b506101d37f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161010c565b61017e6101f9366004611327565b6105d0565b34801561020a57600080fd5b506101027f000000000000000000000000000000000000000000000000000000000000000081565b34801561023e57600080fd5b506101d37f000000000000000000000000000000000000000000000000000000000000000081565b34801561027257600080fd5b5061010260025481565b34801561028857600080fd5b506101d36001600160981b0181565b3480156102a357600080fd5b50610102620186a081565b3480156102ba57600080fd5b506101d3600160981b81565b3480156102d257600080fd5b5061017e6102e136600461138b565b6108ff565b6000806103006001600160981b0131600160981b3161142a565b90507f0000000000000000000000000000000000000000000000000000000000000000600254610330919061142a565b8111156103b65760405162461bcd60e51b815260206004820152604360248201527f4e6174697665546f6b656e44657374696e6174696f6e3a20464154414c202d2060448201527f636f6e74726163742068617320746f6b656e7320756e6163636f756e746564206064820152623337b960e91b608482015260a4015b60405180910390fd5b807f00000000000000000000000000000000000000000000000000000000000000006002546103e5919061142a565b6103ef919061143d565b91505090565b6000600160981b6001600160a01b031631905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663624488506040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602001888036038101906104aa9190611450565b8152602001620186a081526020018787808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152505050908252506040805160208082018990528251808303820181528284019093529092019161051c91600191606001611506565b6040516020818303038152906040528152506040518263ffffffff1660e01b815260040161054a9190611581565b6020604051808303816000875af1158015610569573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061058d91906115ff565b9050807f2550fa6041684d40e635e29e93dde9017d70c25b46aa88393317b5182ed6ae7c836040516105c191815260200190565b60405180910390a25050505050565b6105d8610d10565b6001600160a01b0384166105fe5760405162461bcd60e51b81526004016103ad90611618565b6001541561066b5760405162461bcd60e51b815260206004820152603460248201527f4e6174697665546f6b656e44657374696e6174696f6e3a20636f6e7472616374604482015273081d5b99195c98dbdb1b185d195c985b1a5e995960621b60648201526084016103ad565b60006020840135156106ca576106916106876020860186611666565b8560200135610d69565b90506106ca6106a36020860186611666565b7f000000000000000000000000000000000000000000000000000000000000000083610ed3565b6040516001600160981b01903480156108fc02916000818181858888f193505050501580156106fd573d6000803e3d6000fd5b5060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663624488506040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602001888036038101906107a09190611450565b8152602001620186a081526020018787808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250938552505060405160209384019361080e92508d913491016001600160a01b03929092168252602082015260400190565b60408051601f198184030181529082905261082c9291602001611506565b6040516020818303038152906040528152506040518263ffffffff1660e01b815260040161085a9190611581565b6020604051808303816000875af1158015610879573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089d91906115ff565b905080866001600160a01b0316336001600160a01b03167f0322cbb1d3c23f6dbf1deddb3b4ef3ce0f93ae6eec7b44e4f395804104466d14346040516108e591815260200190565b60405180910390a450506108f96001600055565b50505050565b610907610d10565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109af5760405162461bcd60e51b815260206004820152604160248201527f4e6174697665546f6b656e44657374696e6174696f6e3a20756e617574686f7260448201527f697a65642054656c65706f727465724d657373656e67657220636f6e747261636064820152601d60fa1b608482015260a4016103ad565b7f00000000000000000000000000000000000000000000000000000000000000008414610a335760405162461bcd60e51b815260206004820152602c60248201527f4e6174697665546f6b656e44657374696e6174696f6e3a20696e76616c69642060448201526b39b7bab931b29031b430b4b760a11b60648201526084016103ad565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b031614610ac85760405162461bcd60e51b815260206004820152602b60248201527f4e6174697665546f6b656e44657374696e6174696f6e3a20756e617574686f7260448201526a34bd32b21039b2b73232b960a91b60648201526084016103ad565b600080610ad78385018561168a565b90925090506001600160a01b038216610b025760405162461bcd60e51b81526004016103ad90611618565b80600003610b665760405162461bcd60e51b815260206004820152602b60248201527f4e6174697665546f6b656e44657374696e6174696f6e3a207a65726f2074726160448201526a6e736665722076616c756560a81b60648201526084016103ad565b600154819015610c2b57600154821115610bcf5760015460408051918252600060208301527f244160b15e69cc411f041d94ae7fab6f6bba85dade8403216c05ff4b920d5449910160405180910390a1600154610bc3908361143d565b60006001559050610c2b565b8160016000828254610be1919061143d565b90915550506001546040805184815260208101929092527f244160b15e69cc411f041d94ae7fab6f6bba85dade8403216c05ff4b920d5449910160405180910390a1505050610d06565b6040516327ad555d60e11b81526001600160a01b038481166004830152602482018390527f00000000000000000000000000000000000000000000000000000000000000001690634f5aaaba90604401600060405180830381600087803b158015610c9557600080fd5b505af1158015610ca9573d6000803e3d6000fd5b505050508060026000828254610cbf919061142a565b90915550506040518181526001600160a01b038416907fd949ea0e9d5db53492d77f28fd5467fb2f6c4f5b88e3350e3c36729b76e99cf29060200160405180910390a25050505b6108f96001600055565b600260005403610d625760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016103ad565b6002600055565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015610db2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dd691906115ff565b9050610ded6001600160a01b038516333086610fb8565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610e34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5891906115ff565b9050818111610ebe5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016103ad565b610ec8828261143d565b925050505b92915050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015610f24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4891906115ff565b610f52919061142a565b6040516001600160a01b0385166024820152604481018290529091506108f990859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610ff0565b6040516001600160a01b03808516602483015283166044820152606481018290526108f99085906323b872dd60e01b90608401610f81565b6000611045826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166110c79092919063ffffffff16565b8051909150156110c2578080602001905181019061106391906116b6565b6110c25760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016103ad565b505050565b60606110d684846000856110de565b949350505050565b60608247101561113f5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016103ad565b600080866001600160a01b0316858760405161115b91906116d8565b60006040518083038185875af1925050503d8060008114611198576040519150601f19603f3d011682016040523d82523d6000602084013e61119d565b606091505b50915091506111ae878383876111b9565b979650505050505050565b60608315611228578251600003611221576001600160a01b0385163b6112215760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103ad565b50816110d6565b6110d6838381511561123d5781518083602001fd5b8060405162461bcd60e51b81526004016103ad91906116f4565b60006040828403121561126957600080fd5b50919050565b60008083601f84011261128157600080fd5b50813567ffffffffffffffff81111561129957600080fd5b6020830191508360208260051b85010111156112b457600080fd5b9250929050565b6000806000606084860312156112d057600080fd5b6112da8585611257565b9250604084013567ffffffffffffffff8111156112f657600080fd5b6113028682870161126f565b9497909650939450505050565b6001600160a01b038116811461132457600080fd5b50565b6000806000806080858703121561133d57600080fd5b84356113488161130f565b93506113578660208701611257565b9250606085013567ffffffffffffffff81111561137357600080fd5b61137f8782880161126f565b95989497509550505050565b600080600080606085870312156113a157600080fd5b8435935060208501356113b38161130f565b9250604085013567ffffffffffffffff808211156113d057600080fd5b818701915087601f8301126113e457600080fd5b8135818111156113f357600080fd5b88602082850101111561140557600080fd5b95989497505060200194505050565b634e487b7160e01b600052601160045260246000fd5b80820180821115610ecd57610ecd611414565b81810381811115610ecd57610ecd611414565b60006040828403121561146257600080fd5b6040516040810181811067ffffffffffffffff8211171561149357634e487b7160e01b600052604160045260246000fd5b60405282356114a18161130f565b81526020928301359281019290925250919050565b60005b838110156114d15781810151838201526020016114b9565b50506000910152565b600081518084526114f28160208601602086016114b6565b601f01601f19169290920160200192915050565b60006002841061152657634e487b7160e01b600052602160045260246000fd5b838252604060208301526110d660408301846114da565b600081518084526020808501945080840160005b838110156115765781516001600160a01b031687529582019590820190600101611551565b509495945050505050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c08401526115e261010084018261153d565b905060a0840151601f198483030160e0850152610ec882826114da565b60006020828403121561161157600080fd5b5051919050565b6020808252602e908201527f4e6174697665546f6b656e44657374696e6174696f6e3a207a65726f2072656360408201526d697069656e74206164647265737360901b606082015260800190565b60006020828403121561167857600080fd5b81356116838161130f565b9392505050565b6000806040838503121561169d57600080fd5b82356116a88161130f565b946020939093013593505050565b6000602082840312156116c857600080fd5b8151801515811461168357600080fd5b600082516116ea8184602087016114b6565b9190910192915050565b60208152600061168360208301846114da56fea26469706673582212203e788042f3e702711fff43e16b2fe5b4ba57bcaa7ce99474ec58f977be39e37b64736f6c634300081200334e6174697665546f6b656e44657374696e6174696f6e3a207a65726f20736f75",
}

// NativeTokenDestinationABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenDestinationMetaData.ABI instead.
var NativeTokenDestinationABI = NativeTokenDestinationMetaData.ABI

// NativeTokenDestinationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenDestinationMetaData.Bin instead.
var NativeTokenDestinationBin = NativeTokenDestinationMetaData.Bin

// DeployNativeTokenDestination deploys a new Ethereum contract, binding an instance of NativeTokenDestination to it.
func DeployNativeTokenDestination(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterMessengerAddress common.Address, sourceBlockchainID_ [32]byte, nativeTokenSourceAddress_ common.Address, initialReserveImbalance_ *big.Int) (common.Address, *types.Transaction, *NativeTokenDestination, error) {
	parsed, err := NativeTokenDestinationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenDestinationBin), backend, teleporterMessengerAddress, sourceBlockchainID_, nativeTokenSourceAddress_, initialReserveImbalance_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenDestination{NativeTokenDestinationCaller: NativeTokenDestinationCaller{contract: contract}, NativeTokenDestinationTransactor: NativeTokenDestinationTransactor{contract: contract}, NativeTokenDestinationFilterer: NativeTokenDestinationFilterer{contract: contract}}, nil
}

// NativeTokenDestination is an auto generated Go binding around an Ethereum contract.
type NativeTokenDestination struct {
	NativeTokenDestinationCaller     // Read-only binding to the contract
	NativeTokenDestinationTransactor // Write-only binding to the contract
	NativeTokenDestinationFilterer   // Log filterer for contract events
}

// NativeTokenDestinationCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenDestinationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenDestinationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenDestinationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenDestinationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenDestinationSession struct {
	Contract     *NativeTokenDestination // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NativeTokenDestinationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenDestinationCallerSession struct {
	Contract *NativeTokenDestinationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// NativeTokenDestinationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenDestinationTransactorSession struct {
	Contract     *NativeTokenDestinationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// NativeTokenDestinationRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenDestinationRaw struct {
	Contract *NativeTokenDestination // Generic contract binding to access the raw methods on
}

// NativeTokenDestinationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenDestinationCallerRaw struct {
	Contract *NativeTokenDestinationCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenDestinationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenDestinationTransactorRaw struct {
	Contract *NativeTokenDestinationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenDestination creates a new instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestination(address common.Address, backend bind.ContractBackend) (*NativeTokenDestination, error) {
	contract, err := bindNativeTokenDestination(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestination{NativeTokenDestinationCaller: NativeTokenDestinationCaller{contract: contract}, NativeTokenDestinationTransactor: NativeTokenDestinationTransactor{contract: contract}, NativeTokenDestinationFilterer: NativeTokenDestinationFilterer{contract: contract}}, nil
}

// NewNativeTokenDestinationCaller creates a new read-only instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenDestinationCaller, error) {
	contract, err := bindNativeTokenDestination(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCaller{contract: contract}, nil
}

// NewNativeTokenDestinationTransactor creates a new write-only instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenDestinationTransactor, error) {
	contract, err := bindNativeTokenDestination(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTransactor{contract: contract}, nil
}

// NewNativeTokenDestinationFilterer creates a new log filterer instance of NativeTokenDestination, bound to a specific deployed contract.
func NewNativeTokenDestinationFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenDestinationFilterer, error) {
	contract, err := bindNativeTokenDestination(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationFilterer{contract: contract}, nil
}

// bindNativeTokenDestination binds a generic wrapper to an already deployed contract.
func bindNativeTokenDestination(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenDestinationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenDestination.Contract.NativeTokenDestinationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.NativeTokenDestinationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenDestination *NativeTokenDestinationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.NativeTokenDestinationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenDestination *NativeTokenDestinationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenDestination.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenDestination *NativeTokenDestinationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenDestination *NativeTokenDestinationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.contract.Transact(opts, method, params...)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BURNEDTXFEESADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "BURNED_TX_FEES_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDTXFEESADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNEDTXFEESADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0xa2a95017.
//
// Solidity: function BURN_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) BURNFORTRANSFERADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "BURN_FOR_TRANSFER_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0xa2a95017.
//
// Solidity: function BURN_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) BURNFORTRANSFERADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNFORTRANSFERADDRESS(&_NativeTokenDestination.CallOpts)
}

// BURNFORTRANSFERADDRESS is a free data retrieval call binding the contract method 0xa2a95017.
//
// Solidity: function BURN_FOR_TRANSFER_ADDRESS() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) BURNFORTRANSFERADDRESS() (common.Address, error) {
	return _NativeTokenDestination.Contract.BURNFORTRANSFERADDRESS(&_NativeTokenDestination.CallOpts)
}

// REPORTBURNEDTOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xd3095126.
//
// Solidity: function REPORT_BURNED_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) REPORTBURNEDTOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "REPORT_BURNED_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REPORTBURNEDTOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xd3095126.
//
// Solidity: function REPORT_BURNED_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) REPORTBURNEDTOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.REPORTBURNEDTOKENSREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// REPORTBURNEDTOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xd3095126.
//
// Solidity: function REPORT_BURNED_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) REPORTBURNEDTOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.REPORTBURNEDTOKENSREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// TRANSFERNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xab285230.
//
// Solidity: function TRANSFER_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TRANSFERNATIVETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "TRANSFER_NATIVE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TRANSFERNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xab285230.
//
// Solidity: function TRANSFER_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TRANSFERNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TRANSFERNATIVETOKENSREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// TRANSFERNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xab285230.
//
// Solidity: function TRANSFER_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TRANSFERNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TRANSFERNATIVETOKENSREQUIREDGAS(&_NativeTokenDestination.CallOpts)
}

// CurrentReserveImbalance is a free data retrieval call binding the contract method 0x00d872ae.
//
// Solidity: function currentReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) CurrentReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "currentReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentReserveImbalance is a free data retrieval call binding the contract method 0x00d872ae.
//
// Solidity: function currentReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) CurrentReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.CurrentReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// CurrentReserveImbalance is a free data retrieval call binding the contract method 0x00d872ae.
//
// Solidity: function currentReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) CurrentReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.CurrentReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) InitialReserveImbalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "initialReserveImbalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.InitialReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// InitialReserveImbalance is a free data retrieval call binding the contract method 0x8ac7dd20.
//
// Solidity: function initialReserveImbalance() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) InitialReserveImbalance() (*big.Int, error) {
	return _NativeTokenDestination.Contract.InitialReserveImbalance(&_NativeTokenDestination.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCaller) IsCollateralized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "isCollateralized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationSession) IsCollateralized() (bool, error) {
	return _NativeTokenDestination.Contract.IsCollateralized(&_NativeTokenDestination.CallOpts)
}

// IsCollateralized is a free data retrieval call binding the contract method 0x49e3284e.
//
// Solidity: function isCollateralized() view returns(bool)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) IsCollateralized() (bool, error) {
	return _NativeTokenDestination.Contract.IsCollateralized(&_NativeTokenDestination.CallOpts)
}

// NativeTokenSourceAddress is a free data retrieval call binding the contract method 0x5d93f9af.
//
// Solidity: function nativeTokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) NativeTokenSourceAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "nativeTokenSourceAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeTokenSourceAddress is a free data retrieval call binding the contract method 0x5d93f9af.
//
// Solidity: function nativeTokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) NativeTokenSourceAddress() (common.Address, error) {
	return _NativeTokenDestination.Contract.NativeTokenSourceAddress(&_NativeTokenDestination.CallOpts)
}

// NativeTokenSourceAddress is a free data retrieval call binding the contract method 0x5d93f9af.
//
// Solidity: function nativeTokenSourceAddress() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) NativeTokenSourceAddress() (common.Address, error) {
	return _NativeTokenDestination.Contract.NativeTokenSourceAddress(&_NativeTokenDestination.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCaller) SourceBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "sourceBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationSession) SourceBlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.SourceBlockchainID(&_NativeTokenDestination.CallOpts)
}

// SourceBlockchainID is a free data retrieval call binding the contract method 0x29b7b3fd.
//
// Solidity: function sourceBlockchainID() view returns(bytes32)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) SourceBlockchainID() ([32]byte, error) {
	return _NativeTokenDestination.Contract.SourceBlockchainID(&_NativeTokenDestination.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationSession) TeleporterMessenger() (common.Address, error) {
	return _NativeTokenDestination.Contract.TeleporterMessenger(&_NativeTokenDestination.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TeleporterMessenger() (common.Address, error) {
	return _NativeTokenDestination.Contract.TeleporterMessenger(&_NativeTokenDestination.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalMinted(&_NativeTokenDestination.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TotalMinted() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalMinted(&_NativeTokenDestination.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenDestination.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalSupply(&_NativeTokenDestination.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NativeTokenDestination *NativeTokenDestinationCallerSession) TotalSupply() (*big.Int, error) {
	return _NativeTokenDestination.Contract.TotalSupply(&_NativeTokenDestination.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "receiveTeleporterMessage", senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReceiveTeleporterMessage(&_NativeTokenDestination.TransactOpts, senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReceiveTeleporterMessage(&_NativeTokenDestination.TransactOpts, senderBlockchainID, senderAddress, message)
}

// ReportTotalBurnedTxFees is a paid mutator transaction binding the contract method 0x3a94fe51.
//
// Solidity: function reportTotalBurnedTxFees((address,uint256) feeInfo, address[] allowedRelayerAddresses) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) ReportTotalBurnedTxFees(opts *bind.TransactOpts, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "reportTotalBurnedTxFees", feeInfo, allowedRelayerAddresses)
}

// ReportTotalBurnedTxFees is a paid mutator transaction binding the contract method 0x3a94fe51.
//
// Solidity: function reportTotalBurnedTxFees((address,uint256) feeInfo, address[] allowedRelayerAddresses) returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) ReportTotalBurnedTxFees(feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReportTotalBurnedTxFees(&_NativeTokenDestination.TransactOpts, feeInfo, allowedRelayerAddresses)
}

// ReportTotalBurnedTxFees is a paid mutator transaction binding the contract method 0x3a94fe51.
//
// Solidity: function reportTotalBurnedTxFees((address,uint256) feeInfo, address[] allowedRelayerAddresses) returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) ReportTotalBurnedTxFees(feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.ReportTotalBurnedTxFees(&_NativeTokenDestination.TransactOpts, feeInfo, allowedRelayerAddresses)
}

// TransferToSource is a paid mutator transaction binding the contract method 0x75846562.
//
// Solidity: function transferToSource(address recipient, (address,uint256) feeInfo, address[] allowedRelayerAddresses) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactor) TransferToSource(opts *bind.TransactOpts, recipient common.Address, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.contract.Transact(opts, "transferToSource", recipient, feeInfo, allowedRelayerAddresses)
}

// TransferToSource is a paid mutator transaction binding the contract method 0x75846562.
//
// Solidity: function transferToSource(address recipient, (address,uint256) feeInfo, address[] allowedRelayerAddresses) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationSession) TransferToSource(recipient common.Address, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferToSource(&_NativeTokenDestination.TransactOpts, recipient, feeInfo, allowedRelayerAddresses)
}

// TransferToSource is a paid mutator transaction binding the contract method 0x75846562.
//
// Solidity: function transferToSource(address recipient, (address,uint256) feeInfo, address[] allowedRelayerAddresses) payable returns()
func (_NativeTokenDestination *NativeTokenDestinationTransactorSession) TransferToSource(recipient common.Address, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenDestination.Contract.TransferToSource(&_NativeTokenDestination.TransactOpts, recipient, feeInfo, allowedRelayerAddresses)
}

// NativeTokenDestinationCollateralAddedIterator is returned from FilterCollateralAdded and is used to iterate over the raw logs and unpacked data for CollateralAdded events raised by the NativeTokenDestination contract.
type NativeTokenDestinationCollateralAddedIterator struct {
	Event *NativeTokenDestinationCollateralAdded // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationCollateralAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationCollateralAdded)
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
		it.Event = new(NativeTokenDestinationCollateralAdded)
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
func (it *NativeTokenDestinationCollateralAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationCollateralAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationCollateralAdded represents a CollateralAdded event raised by the NativeTokenDestination contract.
type NativeTokenDestinationCollateralAdded struct {
	Amount    *big.Int
	Remaining *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdded is a free log retrieval operation binding the contract event 0x244160b15e69cc411f041d94ae7fab6f6bba85dade8403216c05ff4b920d5449.
//
// Solidity: event CollateralAdded(uint256 amount, uint256 remaining)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterCollateralAdded(opts *bind.FilterOpts) (*NativeTokenDestinationCollateralAddedIterator, error) {

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "CollateralAdded")
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationCollateralAddedIterator{contract: _NativeTokenDestination.contract, event: "CollateralAdded", logs: logs, sub: sub}, nil
}

// WatchCollateralAdded is a free log subscription operation binding the contract event 0x244160b15e69cc411f041d94ae7fab6f6bba85dade8403216c05ff4b920d5449.
//
// Solidity: event CollateralAdded(uint256 amount, uint256 remaining)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchCollateralAdded(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationCollateralAdded) (event.Subscription, error) {

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "CollateralAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationCollateralAdded)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
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

// ParseCollateralAdded is a log parse operation binding the contract event 0x244160b15e69cc411f041d94ae7fab6f6bba85dade8403216c05ff4b920d5449.
//
// Solidity: event CollateralAdded(uint256 amount, uint256 remaining)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseCollateralAdded(log types.Log) (*NativeTokenDestinationCollateralAdded, error) {
	event := new(NativeTokenDestinationCollateralAdded)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "CollateralAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationNativeTokensMintedIterator is returned from FilterNativeTokensMinted and is used to iterate over the raw logs and unpacked data for NativeTokensMinted events raised by the NativeTokenDestination contract.
type NativeTokenDestinationNativeTokensMintedIterator struct {
	Event *NativeTokenDestinationNativeTokensMinted // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationNativeTokensMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationNativeTokensMinted)
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
		it.Event = new(NativeTokenDestinationNativeTokensMinted)
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
func (it *NativeTokenDestinationNativeTokensMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationNativeTokensMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationNativeTokensMinted represents a NativeTokensMinted event raised by the NativeTokenDestination contract.
type NativeTokenDestinationNativeTokensMinted struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNativeTokensMinted is a free log retrieval operation binding the contract event 0xd949ea0e9d5db53492d77f28fd5467fb2f6c4f5b88e3350e3c36729b76e99cf2.
//
// Solidity: event NativeTokensMinted(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterNativeTokensMinted(opts *bind.FilterOpts, recipient []common.Address) (*NativeTokenDestinationNativeTokensMintedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "NativeTokensMinted", recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationNativeTokensMintedIterator{contract: _NativeTokenDestination.contract, event: "NativeTokensMinted", logs: logs, sub: sub}, nil
}

// WatchNativeTokensMinted is a free log subscription operation binding the contract event 0xd949ea0e9d5db53492d77f28fd5467fb2f6c4f5b88e3350e3c36729b76e99cf2.
//
// Solidity: event NativeTokensMinted(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchNativeTokensMinted(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationNativeTokensMinted, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "NativeTokensMinted", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationNativeTokensMinted)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "NativeTokensMinted", log); err != nil {
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

// ParseNativeTokensMinted is a log parse operation binding the contract event 0xd949ea0e9d5db53492d77f28fd5467fb2f6c4f5b88e3350e3c36729b76e99cf2.
//
// Solidity: event NativeTokensMinted(address indexed recipient, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseNativeTokensMinted(log types.Log) (*NativeTokenDestinationNativeTokensMinted, error) {
	event := new(NativeTokenDestinationNativeTokensMinted)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "NativeTokensMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationReportTotalBurnedTxFeesIterator is returned from FilterReportTotalBurnedTxFees and is used to iterate over the raw logs and unpacked data for ReportTotalBurnedTxFees events raised by the NativeTokenDestination contract.
type NativeTokenDestinationReportTotalBurnedTxFeesIterator struct {
	Event *NativeTokenDestinationReportTotalBurnedTxFees // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationReportTotalBurnedTxFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationReportTotalBurnedTxFees)
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
		it.Event = new(NativeTokenDestinationReportTotalBurnedTxFees)
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
func (it *NativeTokenDestinationReportTotalBurnedTxFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationReportTotalBurnedTxFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationReportTotalBurnedTxFees represents a ReportTotalBurnedTxFees event raised by the NativeTokenDestination contract.
type NativeTokenDestinationReportTotalBurnedTxFees struct {
	TeleporterMessageID *big.Int
	BurnAddressBalance  *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReportTotalBurnedTxFees is a free log retrieval operation binding the contract event 0x2550fa6041684d40e635e29e93dde9017d70c25b46aa88393317b5182ed6ae7c.
//
// Solidity: event ReportTotalBurnedTxFees(uint256 indexed teleporterMessageID, uint256 burnAddressBalance)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterReportTotalBurnedTxFees(opts *bind.FilterOpts, teleporterMessageID []*big.Int) (*NativeTokenDestinationReportTotalBurnedTxFeesIterator, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "ReportTotalBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationReportTotalBurnedTxFeesIterator{contract: _NativeTokenDestination.contract, event: "ReportTotalBurnedTxFees", logs: logs, sub: sub}, nil
}

// WatchReportTotalBurnedTxFees is a free log subscription operation binding the contract event 0x2550fa6041684d40e635e29e93dde9017d70c25b46aa88393317b5182ed6ae7c.
//
// Solidity: event ReportTotalBurnedTxFees(uint256 indexed teleporterMessageID, uint256 burnAddressBalance)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchReportTotalBurnedTxFees(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationReportTotalBurnedTxFees, teleporterMessageID []*big.Int) (event.Subscription, error) {

	var teleporterMessageIDRule []interface{}
	for _, teleporterMessageIDItem := range teleporterMessageID {
		teleporterMessageIDRule = append(teleporterMessageIDRule, teleporterMessageIDItem)
	}

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "ReportTotalBurnedTxFees", teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationReportTotalBurnedTxFees)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "ReportTotalBurnedTxFees", log); err != nil {
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

// ParseReportTotalBurnedTxFees is a log parse operation binding the contract event 0x2550fa6041684d40e635e29e93dde9017d70c25b46aa88393317b5182ed6ae7c.
//
// Solidity: event ReportTotalBurnedTxFees(uint256 indexed teleporterMessageID, uint256 burnAddressBalance)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseReportTotalBurnedTxFees(log types.Log) (*NativeTokenDestinationReportTotalBurnedTxFees, error) {
	event := new(NativeTokenDestinationReportTotalBurnedTxFees)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "ReportTotalBurnedTxFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenDestinationTransferToSourceIterator is returned from FilterTransferToSource and is used to iterate over the raw logs and unpacked data for TransferToSource events raised by the NativeTokenDestination contract.
type NativeTokenDestinationTransferToSourceIterator struct {
	Event *NativeTokenDestinationTransferToSource // Event containing the contract specifics and raw log

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
func (it *NativeTokenDestinationTransferToSourceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenDestinationTransferToSource)
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
		it.Event = new(NativeTokenDestinationTransferToSource)
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
func (it *NativeTokenDestinationTransferToSourceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenDestinationTransferToSourceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenDestinationTransferToSource represents a TransferToSource event raised by the NativeTokenDestination contract.
type NativeTokenDestinationTransferToSource struct {
	Sender              common.Address
	Recipient           common.Address
	TeleporterMessageID *big.Int
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTransferToSource is a free log retrieval operation binding the contract event 0x0322cbb1d3c23f6dbf1deddb3b4ef3ce0f93ae6eec7b44e4f395804104466d14.
//
// Solidity: event TransferToSource(address indexed sender, address indexed recipient, uint256 indexed teleporterMessageID, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) FilterTransferToSource(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, teleporterMessageID []*big.Int) (*NativeTokenDestinationTransferToSourceIterator, error) {

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

	logs, sub, err := _NativeTokenDestination.contract.FilterLogs(opts, "TransferToSource", senderRule, recipientRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenDestinationTransferToSourceIterator{contract: _NativeTokenDestination.contract, event: "TransferToSource", logs: logs, sub: sub}, nil
}

// WatchTransferToSource is a free log subscription operation binding the contract event 0x0322cbb1d3c23f6dbf1deddb3b4ef3ce0f93ae6eec7b44e4f395804104466d14.
//
// Solidity: event TransferToSource(address indexed sender, address indexed recipient, uint256 indexed teleporterMessageID, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) WatchTransferToSource(opts *bind.WatchOpts, sink chan<- *NativeTokenDestinationTransferToSource, sender []common.Address, recipient []common.Address, teleporterMessageID []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _NativeTokenDestination.contract.WatchLogs(opts, "TransferToSource", senderRule, recipientRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenDestinationTransferToSource)
				if err := _NativeTokenDestination.contract.UnpackLog(event, "TransferToSource", log); err != nil {
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

// ParseTransferToSource is a log parse operation binding the contract event 0x0322cbb1d3c23f6dbf1deddb3b4ef3ce0f93ae6eec7b44e4f395804104466d14.
//
// Solidity: event TransferToSource(address indexed sender, address indexed recipient, uint256 indexed teleporterMessageID, uint256 amount)
func (_NativeTokenDestination *NativeTokenDestinationFilterer) ParseTransferToSource(log types.Log) (*NativeTokenDestinationTransferToSource, error) {
	event := new(NativeTokenDestinationTransferToSource)
	if err := _NativeTokenDestination.contract.UnpackLog(event, "TransferToSource", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
