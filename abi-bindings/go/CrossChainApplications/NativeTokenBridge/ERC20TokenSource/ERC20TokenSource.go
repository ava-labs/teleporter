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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeTokenDestinationAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20ContractAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurnTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transferAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"}],\"name\":\"TransferToDestination\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURNED_TX_FEES_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationBurnedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20ContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenDestinationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"senderBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"transferToDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620016f3380380620016f383398101604081905262000035916200031b565b60016000556001600160a01b038416620000b15760405162461bcd60e51b815260206004820152603260248201527f4552433230546f6b656e536f757263653a207a65726f2054656c65706f727465604482015271724d657373656e676572206164647265737360701b60648201526084015b60405180910390fd5b6001600160a01b03841660e05282620001105760405162461bcd60e51b815260206004820152602b6024820152600080516020620016d383398151915260448201526a1bdb8818da185a5b88125160aa1b6064820152608401620000a8565b7302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000163573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200018991906200036f565b8303620001ff5760405162461bcd60e51b815260206004820152603460248201527f4552433230546f6b656e536f757263653a2063616e6e6f74206272696467652060448201527f776974682073616d6520626c6f636b636861696e0000000000000000000000006064820152608401620000a8565b60808390526001600160a01b038216620002715760405162461bcd60e51b81526020600482015260336024820152600080516020620016d383398151915260448201527f6f6e20636f6e74726163742061646472657373000000000000000000000000006064820152608401620000a8565b6001600160a01b0380831660a0528116620002e85760405162461bcd60e51b815260206004820152603060248201527f4552433230546f6b656e536f757263653a20696e76616c69642045524332302060448201526f636f6e7472616374206164647265737360801b6064820152608401620000a8565b6001600160a01b031660c0525062000389915050565b80516001600160a01b03811681146200031657600080fd5b919050565b600080600080608085870312156200033257600080fd5b6200033d85620002fe565b9350602085015192506200035460408601620002fe565b91506200036460608601620002fe565b905092959194509250565b6000602082840312156200038257600080fd5b5051919050565b60805160a05160c05160e0516112be620004156000396000818160f50152818161029c015281816102d301526104ee015260008181610183015281816101e30152818161027b0152818161036d01528181610a590152610c3801526000818161013e015281816103330152610606015260008181609d0152818161030d015261058301526112be6000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063b6171f7311610066578063b6171f731461012f578063b8c9091a14610139578063c452165e14610160578063c868efaa1461016b578063e486df151461017e57600080fd5b806341d3014d1461009857806355db3e9e146100d257806387a2edba146100db5780639b3e5803146100f0575b600080fd5b6100bf7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b6100bf60015481565b6100ee6100e9366004610e3d565b6101a5565b005b6101177f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100c9565b6100bf620186a081565b6101177f000000000000000000000000000000000000000000000000000000000000000081565b610117600160981b81565b6100ee610179366004610ed6565b6104db565b6101177f000000000000000000000000000000000000000000000000000000000000000081565b6101ad610786565b6001600160a01b0385166101dc5760405162461bcd60e51b81526004016101d390610f5f565b60405180910390fd5b60006102087f0000000000000000000000000000000000000000000000000000000000000000866107df565b90508381116102705760405162461bcd60e51b815260206004820152602e60248201527f4552433230546f6b656e536f757263653a20696e73756666696369656e74206160448201526d191a9d5cdd195908185b5bdd5b9d60921b60648201526084016101d3565b83156102c1576102c17f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000086610949565b60006102cd8583610fbd565b905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663624488506040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020018b8152508152602001620186a08152602001888880806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250505090825250604080516001600160a01b038e166020808301919091529181018890529101906060016040516020818303038152906040528152506040518263ffffffff1660e01b81526004016104339190611064565b6020604051808303816000875af1158015610452573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061047691906110e2565b60408051848152602081018990529081018290529091506001600160a01b0389169033907f51f5c55ca5f8f5fde6736f19d1d035a8c67fe18c30d445cea44c4816c6a525a09060600160405180910390a35050506104d46001600055565b5050505050565b6104e3610786565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105815760405162461bcd60e51b815260206004820152603b60248201527f4552433230546f6b656e536f757263653a20756e617574686f72697a6564205460448201527f656c65706f727465724d657373656e67657220636f6e7472616374000000000060648201526084016101d3565b7f000000000000000000000000000000000000000000000000000000000000000084146106045760405162461bcd60e51b815260206004820152602b60248201527f4552433230546f6b656e536f757263653a20696e76616c69642064657374696e60448201526a30ba34b7b71031b430b4b760a91b60648201526084016101d3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316146106935760405162461bcd60e51b815260206004820152602560248201527f4552433230546f6b656e536f757263653a20756e617574686f72697a6564207360448201526432b73232b960d91b60648201526084016101d3565b6000806106a283850185611111565b909250905060008260018111156106bb576106bb6111d9565b036106ec57600080828060200190518101906106d791906111ef565b915091506106e58282610a2e565b5050610774565b6001826001811115610700576107006111d9565b0361072c5760008180602001905181019061071b91906110e2565b905061072681610ac5565b50610774565b60405162461bcd60e51b815260206004820181905260248201527f4552433230546f6b656e536f757263653a20696e76616c696420616374696f6e60448201526064016101d3565b50506107806001600055565b50505050565b6002600054036107d85760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016101d3565b6002600055565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015610828573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061084c91906110e2565b90506108636001600160a01b038516333086610af4565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa1580156108aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ce91906110e2565b90508181116109345760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016101d3565b61093e8282610fbd565b925050505b92915050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa15801561099a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109be91906110e2565b6109c8919061121d565b6040516001600160a01b03851660248201526044810182905290915061078090859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610b2c565b6001600160a01b038216610a545760405162461bcd60e51b81526004016101d390610f5f565b610a7f7f00000000000000000000000000000000000000000000000000000000000000008383610c03565b604080516001600160a01b0384168152602081018390527f55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407910160405180910390a15050565b600154811115610af157600060015482610adf9190610fbd565b9050610aea81610c33565b5060018190555b50565b6040516001600160a01b03808516602483015283166044820152606481018290526107809085906323b872dd60e01b906084016109f7565b6000610b81826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610c989092919063ffffffff16565b805190915015610bfe5780806020019051810190610b9f9190611230565b610bfe5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016101d3565b505050565b6040516001600160a01b038316602482015260448101829052610bfe90849063a9059cbb60e01b906064016109f7565b610c627f0000000000000000000000000000000000000000000000000000000000000000600160981b83610c03565b6040518181527f2cd3fd70cd5a5d6d805e90d22741aa1a84590ace7cf01b244719558d266143829060200160405180910390a150565b6060610ca78484600085610caf565b949350505050565b606082471015610d105760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016101d3565b600080866001600160a01b03168587604051610d2c9190611259565b60006040518083038185875af1925050503d8060008114610d69576040519150601f19603f3d011682016040523d82523d6000602084013e610d6e565b606091505b5091509150610d7f87838387610d8a565b979650505050505050565b60608315610df9578251600003610df2576001600160a01b0385163b610df25760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101d3565b5081610ca7565b610ca78383815115610e0e5781518083602001fd5b8060405162461bcd60e51b81526004016101d39190611275565b6001600160a01b0381168114610af157600080fd5b600080600080600060808688031215610e5557600080fd5b8535610e6081610e28565b94506020860135935060408601359250606086013567ffffffffffffffff80821115610e8b57600080fd5b818801915088601f830112610e9f57600080fd5b813581811115610eae57600080fd5b8960208260051b8501011115610ec357600080fd5b9699959850939650602001949392505050565b60008060008060608587031215610eec57600080fd5b843593506020850135610efe81610e28565b9250604085013567ffffffffffffffff80821115610f1b57600080fd5b818701915087601f830112610f2f57600080fd5b813581811115610f3e57600080fd5b886020828501011115610f5057600080fd5b95989497505060200194505050565b60208082526028908201527f4552433230546f6b656e536f757263653a207a65726f20726563697069656e74604082015267206164647265737360c01b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b8181038181111561094357610943610fa7565b600081518084526020808501945080840160005b838110156110095781516001600160a01b031687529582019590820190600101610fe4565b509495945050505050565b60005b8381101561102f578181015183820152602001611017565b50506000910152565b60008151808452611050816020860160208601611014565b601f01601f19169290920160200192915050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c08401526110c5610100840182610fd0565b905060a0840151601f198483030160e085015261093e8282611038565b6000602082840312156110f457600080fd5b5051919050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561112457600080fd5b82356002811061113357600080fd5b9150602083013567ffffffffffffffff8082111561115057600080fd5b818501915085601f83011261116457600080fd5b813581811115611176576111766110fb565b604051601f8201601f19908116603f0116810190838211818310171561119e5761119e6110fb565b816040528281528860208487010111156111b757600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b634e487b7160e01b600052602160045260246000fd5b6000806040838503121561120257600080fd5b825161120d81610e28565b6020939093015192949293505050565b8082018082111561094357610943610fa7565b60006020828403121561124257600080fd5b8151801515811461125257600080fd5b9392505050565b6000825161126b818460208701611014565b9190910192915050565b602081526000611252602083018461103856fea2646970667358221220f7e42fc360c2fbd286b857618d7456ef997666777dc8ca7c3ccc9a49b0b0088864736f6c634300081200334552433230546f6b656e536f757263653a207a65726f2064657374696e617469",
}

// ERC20TokenSourceABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenSourceMetaData.ABI instead.
var ERC20TokenSourceABI = ERC20TokenSourceMetaData.ABI

// ERC20TokenSourceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenSourceMetaData.Bin instead.
var ERC20TokenSourceBin = ERC20TokenSourceMetaData.Bin

// DeployERC20TokenSource deploys a new Ethereum contract, binding an instance of ERC20TokenSource to it.
func DeployERC20TokenSource(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterMessengerAddress common.Address, destinationBlockchainID_ [32]byte, nativeTokenDestinationAddress_ common.Address, erc20ContractAddress_ common.Address) (common.Address, *types.Transaction, *ERC20TokenSource, error) {
	parsed, err := ERC20TokenSourceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenSourceBin), backend, teleporterMessengerAddress, destinationBlockchainID_, nativeTokenDestinationAddress_, erc20ContractAddress_)
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

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCaller) BURNEDTXFEESADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "BURNED_TX_FEES_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _ERC20TokenSource.Contract.BURNEDTXFEESADDRESS(&_ERC20TokenSource.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _ERC20TokenSource.Contract.BURNEDTXFEESADDRESS(&_ERC20TokenSource.CallOpts)
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

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceSession) TeleporterMessenger() (common.Address, error) {
	return _ERC20TokenSource.Contract.TeleporterMessenger(&_ERC20TokenSource.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) TeleporterMessenger() (common.Address, error) {
	return _ERC20TokenSource.Contract.TeleporterMessenger(&_ERC20TokenSource.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenSource.contract.Transact(opts, "receiveTeleporterMessage", senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_ERC20TokenSource *ERC20TokenSourceSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.ReceiveTeleporterMessage(&_ERC20TokenSource.TransactOpts, senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactorSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.ReceiveTeleporterMessage(&_ERC20TokenSource.TransactOpts, senderBlockchainID, senderAddress, message)
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
	TransferAmount      *big.Int
	FeeAmount           *big.Int
	TeleporterMessageID *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTransferToDestination is a free log retrieval operation binding the contract event 0x51f5c55ca5f8f5fde6736f19d1d035a8c67fe18c30d445cea44c4816c6a525a0.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, uint256 transferAmount, uint256 feeAmount, uint256 teleporterMessageID)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) FilterTransferToDestination(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*ERC20TokenSourceTransferToDestinationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.FilterLogs(opts, "TransferToDestination", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceTransferToDestinationIterator{contract: _ERC20TokenSource.contract, event: "TransferToDestination", logs: logs, sub: sub}, nil
}

// WatchTransferToDestination is a free log subscription operation binding the contract event 0x51f5c55ca5f8f5fde6736f19d1d035a8c67fe18c30d445cea44c4816c6a525a0.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, uint256 transferAmount, uint256 feeAmount, uint256 teleporterMessageID)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) WatchTransferToDestination(opts *bind.WatchOpts, sink chan<- *ERC20TokenSourceTransferToDestination, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.WatchLogs(opts, "TransferToDestination", senderRule, recipientRule)
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

// ParseTransferToDestination is a log parse operation binding the contract event 0x51f5c55ca5f8f5fde6736f19d1d035a8c67fe18c30d445cea44c4816c6a525a0.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, uint256 transferAmount, uint256 feeAmount, uint256 teleporterMessageID)
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
