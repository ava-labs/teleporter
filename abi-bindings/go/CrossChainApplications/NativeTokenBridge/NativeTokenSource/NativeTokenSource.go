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

// TeleporterFeeInfo is an auto generated low-level Go binding around an user-defined struct.
type TeleporterFeeInfo struct {
	FeeTokenAddress common.Address
	Amount          *big.Int
}

// NativeTokenSourceMetaData contains all meta data concerning the NativeTokenSource contract.
var NativeTokenSourceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterMessengerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeTokenDestinationAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurnTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teleporterMessageID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferToDestination\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURNED_TX_FEES_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationBurnedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenDestinationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"senderBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterMessenger\",\"outputs\":[{\"internalType\":\"contractITeleporterMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTeleporterFeeInfo\",\"name\":\"feeInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"transferToDestination\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162001643380380620016438339810160408190526200003491620002ae565b60016000556001600160a01b038316620000bb5760405162461bcd60e51b815260206004820152603360248201527f4e6174697665546f6b656e536f757263653a207a65726f2054656c65706f727460448201527f65724d657373656e67657220616464726573730000000000000000000000000060648201526084015b60405180910390fd5b6001600160a01b03831660c052816200011b5760405162461bcd60e51b815260206004820152602c60248201526000805160206200162383398151915260448201526b1a5bdb8818da185a5b88125160a21b6064820152608401620000b2565b7302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200016e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001949190620002ef565b82036200020a5760405162461bcd60e51b815260206004820152603560248201527f4e6174697665546f6b656e536f757263653a2063616e6e6f742062726964676560448201527f20776974682073616d6520626c6f636b636861696e00000000000000000000006064820152608401620000b2565b60808290526001600160a01b0381166200027c5760405162461bcd60e51b815260206004820152603460248201526000805160206200162383398151915260448201527f696f6e20636f6e747261637420616464726573730000000000000000000000006064820152608401620000b2565b6001600160a01b031660a05250620003099050565b80516001600160a01b0381168114620002a957600080fd5b919050565b600080600060608486031215620002c457600080fd5b620002cf8462000291565b925060208401519150620002e66040850162000291565b90509250925092565b6000602082840312156200030257600080fd5b5051919050565b60805160a05160c0516112bb620003686000396000818160ef0152818161027001528181610299015261048c015260008181610167015281816102f901526105a50152600081816092015281816102d3015261052101526112bb6000f3fe60806040526004361061007b5760003560e01c8063b6171f731161004e578063b6171f731461013e578063b8c9091a14610155578063c452165e14610189578063c868efaa146101a157600080fd5b806341d3014d1461008057806355db3e9e146100c75780639b3e5803146100dd578063ad0aee2514610129575b600080fd5b34801561008c57600080fd5b506100b47f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b3480156100d357600080fd5b506100b460015481565b3480156100e957600080fd5b506101117f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100be565b61013c610137366004610e01565b6101c1565b005b34801561014a57600080fd5b506100b4620186a081565b34801561016157600080fd5b506101117f000000000000000000000000000000000000000000000000000000000000000081565b34801561019557600080fd5b50610111600160981b81565b3480156101ad57600080fd5b5061013c6101bc366004610ea0565b610479565b6101c961072a565b6001600160a01b0384166102365760405162461bcd60e51b815260206004820152602960248201527f4e6174697665546f6b656e536f757263653a207a65726f20657263697069656e60448201526874206164647265737360b81b60648201526084015b60405180910390fd5b60006020840135156102955761025c6102526020860186610f1a565b8560200135610783565b905061029561026e6020860186610f1a565b7f0000000000000000000000000000000000000000000000000000000000000000836108ed565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663624488506040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602001888036038101906103379190610f85565b8152602001620186a081526020018787808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152505050908252506040516020918201916103a6918c913491016001600160a01b03929092168252602082015260400190565b6040516020818303038152906040528152506040518263ffffffff1660e01b81526004016103d49190611071565b6020604051808303816000875af11580156103f3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061041791906110ef565b905080866001600160a01b0316336001600160a01b03167f2b4e8f08417773e367064a6aea9ca2df303a60876676f70b6c3c5e66b314ca5a3460405161045f91815260200190565b60405180910390a450506104736001600055565b50505050565b61048161072a565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461051f5760405162461bcd60e51b815260206004820152603c60248201527f4e6174697665546f6b656e536f757263653a20756e617574686f72697a65642060448201527f74656c65706f727465724d657373656e67657220636f6e747261637400000000606482015260840161022d565b7f000000000000000000000000000000000000000000000000000000000000000084146105a35760405162461bcd60e51b815260206004820152602c60248201527f4e6174697665546f6b656e536f757263653a20696e76616c696420646573746960448201526b3730ba34b7b71031b430b4b760a11b606482015260840161022d565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316146106335760405162461bcd60e51b815260206004820152602660248201527f4e6174697665546f6b656e536f757263653a20756e617574686f72697a65642060448201526539b2b73232b960d11b606482015260840161022d565b60008061064283850185611108565b9092509050600082600181111561065b5761065b6111b4565b0361068c576000808280602001905181019061067791906111ca565b9150915061068582826109d2565b505061071e565b60018260018111156106a0576106a06111b4565b036106cc576000818060200190518101906106bb91906110ef565b90506106c681610ab6565b5061071e565b60405162461bcd60e51b815260206004820152602160248201527f4e6174697665546f6b656e536f757263653a20696e76616c696420616374696f6044820152603760f91b606482015260840161022d565b50506104736001600055565b60026000540361077c5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161022d565b6002600055565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa1580156107cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107f091906110ef565b90506108076001600160a01b038516333086610ae5565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa15801561084e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087291906110ef565b90508181116108d85760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b606482015260840161022d565b6108e2828261120e565b925050505b92915050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa15801561093e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096291906110ef565b61096c9190611221565b6040516001600160a01b03851660248201526044810182905290915061047390859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610b1d565b6001600160a01b038216610a395760405162461bcd60e51b815260206004820152602860248201527f4552433230546f6b656e536f757263653a207a65726f20726563697069656e74604482015267206164647265737360c01b606482015260840161022d565b6040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015610a6f573d6000803e3d6000fd5b50604080516001600160a01b0384168152602081018390527f55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407910160405180910390a15050565b600154811115610ae257600060015482610ad0919061120e565b9050610adb81610bf4565b5060018190555b50565b6040516001600160a01b03808516602483015283166044820152606481018290526104739085906323b872dd60e01b9060840161099b565b6000610b72826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610c5c9092919063ffffffff16565b805190915015610bef5780806020019051810190610b909190611234565b610bef5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161022d565b505050565b604051600160981b9082156108fc029083906000818181858888f19350505050158015610c25573d6000803e3d6000fd5b506040518181527f2cd3fd70cd5a5d6d805e90d22741aa1a84590ace7cf01b244719558d266143829060200160405180910390a150565b6060610c6b8484600085610c73565b949350505050565b606082471015610cd45760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161022d565b600080866001600160a01b03168587604051610cf09190611256565b60006040518083038185875af1925050503d8060008114610d2d576040519150601f19603f3d011682016040523d82523d6000602084013e610d32565b606091505b5091509150610d4387838387610d4e565b979650505050505050565b60608315610dbd578251600003610db6576001600160a01b0385163b610db65760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161022d565b5081610c6b565b610c6b8383815115610dd25781518083602001fd5b8060405162461bcd60e51b815260040161022d9190611272565b6001600160a01b0381168114610ae257600080fd5b6000806000808486036080811215610e1857600080fd5b8535610e2381610dec565b94506040601f1982011215610e3757600080fd5b50602085019250606085013567ffffffffffffffff80821115610e5957600080fd5b818701915087601f830112610e6d57600080fd5b813581811115610e7c57600080fd5b8860208260051b8501011115610e9157600080fd5b95989497505060200194505050565b60008060008060608587031215610eb657600080fd5b843593506020850135610ec881610dec565b9250604085013567ffffffffffffffff80821115610ee557600080fd5b818701915087601f830112610ef957600080fd5b813581811115610f0857600080fd5b886020828501011115610e9157600080fd5b600060208284031215610f2c57600080fd5b8135610f3781610dec565b9392505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610f7d57610f7d610f3e565b604052919050565b600060408284031215610f9757600080fd5b6040516040810181811067ffffffffffffffff82111715610fba57610fba610f3e565b6040528235610fc881610dec565b81526020928301359281019290925250919050565b600081518084526020808501945080840160005b838110156110165781516001600160a01b031687529582019590820190600101610ff1565b509495945050505050565b60005b8381101561103c578181015183820152602001611024565b50506000910152565b6000815180845261105d816020860160208601611021565b601f01601f19169290920160200192915050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c08401526110d2610100840182610fdd565b905060a0840151601f198483030160e08501526108e28282611045565b60006020828403121561110157600080fd5b5051919050565b6000806040838503121561111b57600080fd5b82356002811061112a57600080fd5b915060208381013567ffffffffffffffff8082111561114857600080fd5b818601915086601f83011261115c57600080fd5b81358181111561116e5761116e610f3e565b611180601f8201601f19168501610f54565b9150808252878482850101111561119657600080fd5b80848401858401376000848284010152508093505050509250929050565b634e487b7160e01b600052602160045260246000fd5b600080604083850312156111dd57600080fd5b82516111e881610dec565b6020939093015192949293505050565b634e487b7160e01b600052601160045260246000fd5b818103818111156108e7576108e76111f8565b808201808211156108e7576108e76111f8565b60006020828403121561124657600080fd5b81518015158114610f3757600080fd5b60008251611268818460208701611021565b9190910192915050565b602081526000610f37602083018461104556fea2646970667358221220e8cbdf2ffcdd9a41fcf7f19146c49a2513c9847816ef7fceeb854242be953f7c64736f6c634300081200334e6174697665546f6b656e536f757263653a207a65726f2064657374696e6174",
}

// NativeTokenSourceABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenSourceMetaData.ABI instead.
var NativeTokenSourceABI = NativeTokenSourceMetaData.ABI

// NativeTokenSourceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenSourceMetaData.Bin instead.
var NativeTokenSourceBin = NativeTokenSourceMetaData.Bin

// DeployNativeTokenSource deploys a new Ethereum contract, binding an instance of NativeTokenSource to it.
func DeployNativeTokenSource(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterMessengerAddress common.Address, destinationBlockchainID_ [32]byte, nativeTokenDestinationAddress_ common.Address) (common.Address, *types.Transaction, *NativeTokenSource, error) {
	parsed, err := NativeTokenSourceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenSourceBin), backend, teleporterMessengerAddress, destinationBlockchainID_, nativeTokenDestinationAddress_)
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

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCaller) BURNEDTXFEESADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "BURNED_TX_FEES_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenSource *NativeTokenSourceSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenSource.Contract.BURNEDTXFEESADDRESS(&_NativeTokenSource.CallOpts)
}

// BURNEDTXFEESADDRESS is a free data retrieval call binding the contract method 0xc452165e.
//
// Solidity: function BURNED_TX_FEES_ADDRESS() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCallerSession) BURNEDTXFEESADDRESS() (common.Address, error) {
	return _NativeTokenSource.Contract.BURNEDTXFEESADDRESS(&_NativeTokenSource.CallOpts)
}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceCaller) MINTNATIVETOKENSREQUIREDGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "MINT_NATIVE_TOKENS_REQUIRED_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceSession) MINTNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSource.Contract.MINTNATIVETOKENSREQUIREDGAS(&_NativeTokenSource.CallOpts)
}

// MINTNATIVETOKENSREQUIREDGAS is a free data retrieval call binding the contract method 0xb6171f73.
//
// Solidity: function MINT_NATIVE_TOKENS_REQUIRED_GAS() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceCallerSession) MINTNATIVETOKENSREQUIREDGAS() (*big.Int, error) {
	return _NativeTokenSource.Contract.MINTNATIVETOKENSREQUIREDGAS(&_NativeTokenSource.CallOpts)
}

// DestinationBlockchainID is a free data retrieval call binding the contract method 0x41d3014d.
//
// Solidity: function destinationBlockchainID() view returns(bytes32)
func (_NativeTokenSource *NativeTokenSourceCaller) DestinationBlockchainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "destinationBlockchainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DestinationBlockchainID is a free data retrieval call binding the contract method 0x41d3014d.
//
// Solidity: function destinationBlockchainID() view returns(bytes32)
func (_NativeTokenSource *NativeTokenSourceSession) DestinationBlockchainID() ([32]byte, error) {
	return _NativeTokenSource.Contract.DestinationBlockchainID(&_NativeTokenSource.CallOpts)
}

// DestinationBlockchainID is a free data retrieval call binding the contract method 0x41d3014d.
//
// Solidity: function destinationBlockchainID() view returns(bytes32)
func (_NativeTokenSource *NativeTokenSourceCallerSession) DestinationBlockchainID() ([32]byte, error) {
	return _NativeTokenSource.Contract.DestinationBlockchainID(&_NativeTokenSource.CallOpts)
}

// DestinationBurnedTotal is a free data retrieval call binding the contract method 0x55db3e9e.
//
// Solidity: function destinationBurnedTotal() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceCaller) DestinationBurnedTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "destinationBurnedTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DestinationBurnedTotal is a free data retrieval call binding the contract method 0x55db3e9e.
//
// Solidity: function destinationBurnedTotal() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceSession) DestinationBurnedTotal() (*big.Int, error) {
	return _NativeTokenSource.Contract.DestinationBurnedTotal(&_NativeTokenSource.CallOpts)
}

// DestinationBurnedTotal is a free data retrieval call binding the contract method 0x55db3e9e.
//
// Solidity: function destinationBurnedTotal() view returns(uint256)
func (_NativeTokenSource *NativeTokenSourceCallerSession) DestinationBurnedTotal() (*big.Int, error) {
	return _NativeTokenSource.Contract.DestinationBurnedTotal(&_NativeTokenSource.CallOpts)
}

// NativeTokenDestinationAddress is a free data retrieval call binding the contract method 0xb8c9091a.
//
// Solidity: function nativeTokenDestinationAddress() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCaller) NativeTokenDestinationAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "nativeTokenDestinationAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeTokenDestinationAddress is a free data retrieval call binding the contract method 0xb8c9091a.
//
// Solidity: function nativeTokenDestinationAddress() view returns(address)
func (_NativeTokenSource *NativeTokenSourceSession) NativeTokenDestinationAddress() (common.Address, error) {
	return _NativeTokenSource.Contract.NativeTokenDestinationAddress(&_NativeTokenSource.CallOpts)
}

// NativeTokenDestinationAddress is a free data retrieval call binding the contract method 0xb8c9091a.
//
// Solidity: function nativeTokenDestinationAddress() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCallerSession) NativeTokenDestinationAddress() (common.Address, error) {
	return _NativeTokenSource.Contract.NativeTokenDestinationAddress(&_NativeTokenSource.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCaller) TeleporterMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenSource.contract.Call(opts, &out, "teleporterMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenSource *NativeTokenSourceSession) TeleporterMessenger() (common.Address, error) {
	return _NativeTokenSource.Contract.TeleporterMessenger(&_NativeTokenSource.CallOpts)
}

// TeleporterMessenger is a free data retrieval call binding the contract method 0x9b3e5803.
//
// Solidity: function teleporterMessenger() view returns(address)
func (_NativeTokenSource *NativeTokenSourceCallerSession) TeleporterMessenger() (common.Address, error) {
	return _NativeTokenSource.Contract.TeleporterMessenger(&_NativeTokenSource.CallOpts)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_NativeTokenSource *NativeTokenSourceTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSource.contract.Transact(opts, "receiveTeleporterMessage", senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_NativeTokenSource *NativeTokenSourceSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.ReceiveTeleporterMessage(&_NativeTokenSource.TransactOpts, senderBlockchainID, senderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 senderBlockchainID, address senderAddress, bytes message) returns()
func (_NativeTokenSource *NativeTokenSourceTransactorSession) ReceiveTeleporterMessage(senderBlockchainID [32]byte, senderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.ReceiveTeleporterMessage(&_NativeTokenSource.TransactOpts, senderBlockchainID, senderAddress, message)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0xad0aee25.
//
// Solidity: function transferToDestination(address recipient, (address,uint256) feeInfo, address[] allowedRelayerAddresses) payable returns()
func (_NativeTokenSource *NativeTokenSourceTransactor) TransferToDestination(opts *bind.TransactOpts, recipient common.Address, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenSource.contract.Transact(opts, "transferToDestination", recipient, feeInfo, allowedRelayerAddresses)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0xad0aee25.
//
// Solidity: function transferToDestination(address recipient, (address,uint256) feeInfo, address[] allowedRelayerAddresses) payable returns()
func (_NativeTokenSource *NativeTokenSourceSession) TransferToDestination(recipient common.Address, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.TransferToDestination(&_NativeTokenSource.TransactOpts, recipient, feeInfo, allowedRelayerAddresses)
}

// TransferToDestination is a paid mutator transaction binding the contract method 0xad0aee25.
//
// Solidity: function transferToDestination(address recipient, (address,uint256) feeInfo, address[] allowedRelayerAddresses) payable returns()
func (_NativeTokenSource *NativeTokenSourceTransactorSession) TransferToDestination(recipient common.Address, feeInfo TeleporterFeeInfo, allowedRelayerAddresses []common.Address) (*types.Transaction, error) {
	return _NativeTokenSource.Contract.TransferToDestination(&_NativeTokenSource.TransactOpts, recipient, feeInfo, allowedRelayerAddresses)
}

// NativeTokenSourceBurnTokensIterator is returned from FilterBurnTokens and is used to iterate over the raw logs and unpacked data for BurnTokens events raised by the NativeTokenSource contract.
type NativeTokenSourceBurnTokensIterator struct {
	Event *NativeTokenSourceBurnTokens // Event containing the contract specifics and raw log

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
func (it *NativeTokenSourceBurnTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSourceBurnTokens)
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
		it.Event = new(NativeTokenSourceBurnTokens)
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
func (it *NativeTokenSourceBurnTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSourceBurnTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSourceBurnTokens represents a BurnTokens event raised by the NativeTokenSource contract.
type NativeTokenSourceBurnTokens struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnTokens is a free log retrieval operation binding the contract event 0x2cd3fd70cd5a5d6d805e90d22741aa1a84590ace7cf01b244719558d26614382.
//
// Solidity: event BurnTokens(uint256 amount)
func (_NativeTokenSource *NativeTokenSourceFilterer) FilterBurnTokens(opts *bind.FilterOpts) (*NativeTokenSourceBurnTokensIterator, error) {

	logs, sub, err := _NativeTokenSource.contract.FilterLogs(opts, "BurnTokens")
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceBurnTokensIterator{contract: _NativeTokenSource.contract, event: "BurnTokens", logs: logs, sub: sub}, nil
}

// WatchBurnTokens is a free log subscription operation binding the contract event 0x2cd3fd70cd5a5d6d805e90d22741aa1a84590ace7cf01b244719558d26614382.
//
// Solidity: event BurnTokens(uint256 amount)
func (_NativeTokenSource *NativeTokenSourceFilterer) WatchBurnTokens(opts *bind.WatchOpts, sink chan<- *NativeTokenSourceBurnTokens) (event.Subscription, error) {

	logs, sub, err := _NativeTokenSource.contract.WatchLogs(opts, "BurnTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSourceBurnTokens)
				if err := _NativeTokenSource.contract.UnpackLog(event, "BurnTokens", log); err != nil {
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
func (_NativeTokenSource *NativeTokenSourceFilterer) ParseBurnTokens(log types.Log) (*NativeTokenSourceBurnTokens, error) {
	event := new(NativeTokenSourceBurnTokens)
	if err := _NativeTokenSource.contract.UnpackLog(event, "BurnTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSourceTransferToDestinationIterator is returned from FilterTransferToDestination and is used to iterate over the raw logs and unpacked data for TransferToDestination events raised by the NativeTokenSource contract.
type NativeTokenSourceTransferToDestinationIterator struct {
	Event *NativeTokenSourceTransferToDestination // Event containing the contract specifics and raw log

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
func (it *NativeTokenSourceTransferToDestinationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSourceTransferToDestination)
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
		it.Event = new(NativeTokenSourceTransferToDestination)
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
func (it *NativeTokenSourceTransferToDestinationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSourceTransferToDestinationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSourceTransferToDestination represents a TransferToDestination event raised by the NativeTokenSource contract.
type NativeTokenSourceTransferToDestination struct {
	Sender              common.Address
	Recipient           common.Address
	TeleporterMessageID *big.Int
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTransferToDestination is a free log retrieval operation binding the contract event 0x2b4e8f08417773e367064a6aea9ca2df303a60876676f70b6c3c5e66b314ca5a.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, uint256 indexed teleporterMessageID, uint256 amount)
func (_NativeTokenSource *NativeTokenSourceFilterer) FilterTransferToDestination(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, teleporterMessageID []*big.Int) (*NativeTokenSourceTransferToDestinationIterator, error) {

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

	logs, sub, err := _NativeTokenSource.contract.FilterLogs(opts, "TransferToDestination", senderRule, recipientRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceTransferToDestinationIterator{contract: _NativeTokenSource.contract, event: "TransferToDestination", logs: logs, sub: sub}, nil
}

// WatchTransferToDestination is a free log subscription operation binding the contract event 0x2b4e8f08417773e367064a6aea9ca2df303a60876676f70b6c3c5e66b314ca5a.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, uint256 indexed teleporterMessageID, uint256 amount)
func (_NativeTokenSource *NativeTokenSourceFilterer) WatchTransferToDestination(opts *bind.WatchOpts, sink chan<- *NativeTokenSourceTransferToDestination, sender []common.Address, recipient []common.Address, teleporterMessageID []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _NativeTokenSource.contract.WatchLogs(opts, "TransferToDestination", senderRule, recipientRule, teleporterMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSourceTransferToDestination)
				if err := _NativeTokenSource.contract.UnpackLog(event, "TransferToDestination", log); err != nil {
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
func (_NativeTokenSource *NativeTokenSourceFilterer) ParseTransferToDestination(log types.Log) (*NativeTokenSourceTransferToDestination, error) {
	event := new(NativeTokenSourceTransferToDestination)
	if err := _NativeTokenSource.contract.UnpackLog(event, "TransferToDestination", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenSourceUnlockTokensIterator is returned from FilterUnlockTokens and is used to iterate over the raw logs and unpacked data for UnlockTokens events raised by the NativeTokenSource contract.
type NativeTokenSourceUnlockTokensIterator struct {
	Event *NativeTokenSourceUnlockTokens // Event containing the contract specifics and raw log

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
func (it *NativeTokenSourceUnlockTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenSourceUnlockTokens)
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
		it.Event = new(NativeTokenSourceUnlockTokens)
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
func (it *NativeTokenSourceUnlockTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenSourceUnlockTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenSourceUnlockTokens represents a UnlockTokens event raised by the NativeTokenSource contract.
type NativeTokenSourceUnlockTokens struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlockTokens is a free log retrieval operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_NativeTokenSource *NativeTokenSourceFilterer) FilterUnlockTokens(opts *bind.FilterOpts) (*NativeTokenSourceUnlockTokensIterator, error) {

	logs, sub, err := _NativeTokenSource.contract.FilterLogs(opts, "UnlockTokens")
	if err != nil {
		return nil, err
	}
	return &NativeTokenSourceUnlockTokensIterator{contract: _NativeTokenSource.contract, event: "UnlockTokens", logs: logs, sub: sub}, nil
}

// WatchUnlockTokens is a free log subscription operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address recipient, uint256 amount)
func (_NativeTokenSource *NativeTokenSourceFilterer) WatchUnlockTokens(opts *bind.WatchOpts, sink chan<- *NativeTokenSourceUnlockTokens) (event.Subscription, error) {

	logs, sub, err := _NativeTokenSource.contract.WatchLogs(opts, "UnlockTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenSourceUnlockTokens)
				if err := _NativeTokenSource.contract.UnpackLog(event, "UnlockTokens", log); err != nil {
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
func (_NativeTokenSource *NativeTokenSourceFilterer) ParseUnlockTokens(log types.Log) (*NativeTokenSourceUnlockTokens, error) {
	event := new(NativeTokenSourceUnlockTokens)
	if err := _NativeTokenSource.contract.UnpackLog(event, "UnlockTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
