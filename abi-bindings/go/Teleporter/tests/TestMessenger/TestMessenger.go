// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testmessenger

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

// TestMessengerMetaData contains all meta data concerning the TestMessenger contract.
var TestMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"ReceiveMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"SendMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"}],\"name\":\"getCurrentMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001ea438038062001ea4833981016040819052620000349162000522565b6200003e62000052565b6200004a8282620000be565b505062000574565b600054610100900460ff16620000b25760405162461bcd60e51b815260206004820152602b602482015260008051602062001e8483398151915260448201526a6e697469616c697a696e6760a81b60648201526084015b60405180910390fd5b620000bc6200013e565b565b600054610100900460ff166200011a5760405162461bcd60e51b815260206004820152602b602482015260008051602062001e8483398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000a9565b6200012582620001a1565b6200012f6200030b565b6200013a8162000371565b5050565b600054610100900460ff166200019a5760405162461bcd60e51b815260206004820152602b602482015260008051602062001e8483398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000a9565b6001603355565b600054610100900460ff16620001fd5760405162461bcd60e51b815260206004820152602b602482015260008051602062001e8483398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000a9565b6200020762000052565b6001600160a01b038116620002855760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f7274657220726567697374727920616464726573730000000000000000006064820152608401620000a9565b606580546001600160a01b0319166001600160a01b0383169081179091556040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa158015620002df573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200030591906200055a565b60675550565b600054610100900460ff16620003675760405162461bcd60e51b815260206004820152602b602482015260008051602062001e8483398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000a9565b620000bc620003f0565b6200037b62000457565b6001600160a01b038116620003e25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401620000a9565b620003ed81620004b3565b50565b600054610100900460ff166200044c5760405162461bcd60e51b815260206004820152602b602482015260008051602062001e8483398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000a9565b620000bc33620004b3565b6068546001600160a01b03163314620000bc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401620000a9565b606880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80516001600160a01b03811681146200051d57600080fd5b919050565b600080604083850312156200053657600080fd5b620005418362000505565b9150620005516020840162000505565b90509250929050565b6000602082840312156200056d57600080fd5b5051919050565b61190080620005846000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80639731429711610071578063973142971461013d578063b33fead414610179578063c868efaa1461019a578063d2cc7a70146101ad578063f2fde38b146101bf578063f63d09d7146101d257600080fd5b80631a7f5bec146100b95780632b0d8f18146100e95780634511243e146100fe5780635eb9951414610111578063715018a6146101245780638da5cb5b1461012c575b600080fd5b6065546100cc906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100fc6100f73660046112b7565b6101e5565b005b6100fc61010c3660046112b7565b6102e7565b6100fc61011f3660046112d4565b6103e4565b6100fc6103f8565b6068546001600160a01b03166100cc565b61016961014b3660046112b7565b6001600160a01b031660009081526066602052604090205460ff1690565b60405190151581526020016100e0565b61018c6101873660046112d4565b61040c565b6040516100e092919061133d565b6100fc6101a83660046113aa565b6104e4565b6067545b6040519081526020016100e0565b6100fc6101cd3660046112b7565b610690565b6101b16101e0366004611406565b610706565b6101ed610846565b6001600160a01b03811661021c5760405162461bcd60e51b81526004016102139061148c565b60405180910390fd5b6001600160a01b03811660009081526066602052604090205460ff161561029b5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b6064820152608401610213565b6001600160a01b038116600081815260666020526040808220805460ff19166001179055517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b6102ef610846565b6001600160a01b0381166103155760405162461bcd60e51b81526004016102139061148c565b6001600160a01b03811660009081526066602052604090205460ff1661038f5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b6064820152608401610213565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152606660205260409020805460ff19169055565b6103ec610846565b6103f58161084e565b50565b6104006109d8565b61040a6000610a32565b565b6000818152609a60209081526040808320815180830190925280546001600160a01b03168252600181018054606094869493929084019161044c906114da565b80601f0160208091040260200160405190810160405280929190818152602001828054610478906114da565b80156104c55780601f1061049a576101008083540402835291602001916104c5565b820191906000526020600020905b8154815290600101906020018083116104a857829003601f168201915b5050505050815250509050806000015181602001519250925050915091565b6104ec610a84565b6067546065546001600160a01b0316634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015610545573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105699190611514565b10156105d05760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b6064820152608401610213565b6105d93361014b565b1561063f5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b6064820152608401610213565b610680848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610add92505050565b61068a6001603355565b50505050565b6106986109d8565b6001600160a01b0381166106fd5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610213565b6103f581610a32565b6000610710610a84565b60008515610725576107228787610b95565b90505b876001600160a01b0316897fa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca8984898989604051610767959493929190611556565b60405180910390a361082e6040518060c001604052808b81526020018a6001600160a01b0316815260200160405180604001604052808b6001600160a01b03168152602001858152508152602001878152602001600067ffffffffffffffff8111156107d5576107d5611584565b6040519080825280602002602001820160405280156107fe578160200160208202803683370190505b508152602001868660405160200161081792919061159a565b604051602081830303815290604052815250610cff565b91505061083b6001603355565b979650505050505050565b61040a6109d8565b6065546040805163301fd1f560e21b815290516000926001600160a01b03169163c07f47d49160048083019260209291908290030181865afa158015610898573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bc9190611514565b6067549091508183111561092c5760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b6064820152608401610213565b8083116109a15760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e006064820152608401610213565b6067839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6068546001600160a01b0316331461040a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610213565b606880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600260335403610ad65760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610213565b6002603355565b600081806020019051810190610af391906115ae565b6040805180820182526001600160a01b038681168252602080830185815260008a8152609a909252939020825181546001600160a01b03191692169190911781559151929350916001820190610b49908261169e565b50905050826001600160a01b0316847f1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa83604051610b87919061175e565b60405180910390a350505050565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015610bde573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c029190611514565b9050610c196001600160a01b038516333086610e25565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610c60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c849190611514565b9050818111610cea5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b6064820152608401610213565b610cf48282611787565b925050505b92915050565b600080610d0a610e90565b60408401516020015190915015610daf576040830151516001600160a01b0316610d8c5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b6064820152608401610213565b604083015160208101519051610daf916001600160a01b03909116908390610f99565b604051630624488560e41b81526001600160a01b03821690636244885090610ddb9086906004016117de565b6020604051808303816000875af1158015610dfa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1e9190611514565b9392505050565b6040516001600160a01b038085166024830152831660448201526064810182905261068a9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261104b565b600080606560009054906101000a90046001600160a01b03166001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ee6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0a919061185c565b9050610f2e816001600160a01b031660009081526066602052604090205460ff1690565b15610f945760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b6064820152608401610213565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015610fea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061100e9190611514565b6110189190611879565b6040516001600160a01b03851660248201526044810182905290915061068a90859063095ea7b360e01b90606401610e59565b60006110a0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166111229092919063ffffffff16565b80519091501561111d57808060200190518101906110be919061188c565b61111d5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610213565b505050565b60606111318484600085611139565b949350505050565b60608247101561119a5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610213565b600080866001600160a01b031685876040516111b691906118ae565b60006040518083038185875af1925050503d80600081146111f3576040519150601f19603f3d011682016040523d82523d6000602084013e6111f8565b606091505b509150915061083b878383876060831561127357825160000361126c576001600160a01b0385163b61126c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610213565b5081611131565b61113183838151156112885781518083602001fd5b8060405162461bcd60e51b8152600401610213919061175e565b6001600160a01b03811681146103f557600080fd5b6000602082840312156112c957600080fd5b8135610e1e816112a2565b6000602082840312156112e657600080fd5b5035919050565b60005b838110156113085781810151838201526020016112f0565b50506000910152565b600081518084526113298160208601602086016112ed565b601f01601f19169290920160200192915050565b6001600160a01b038316815260406020820181905260009061113190830184611311565b60008083601f84011261137357600080fd5b50813567ffffffffffffffff81111561138b57600080fd5b6020830191508360208285010111156113a357600080fd5b9250929050565b600080600080606085870312156113c057600080fd5b8435935060208501356113d2816112a2565b9250604085013567ffffffffffffffff8111156113ee57600080fd5b6113fa87828801611361565b95989497509550505050565b600080600080600080600060c0888a03121561142157600080fd5b873596506020880135611433816112a2565b95506040880135611443816112a2565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561146d57600080fd5b6114798a828b01611361565b989b979a50959850939692959293505050565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b600181811c908216806114ee57607f821691505b60208210810361150e57634e487b7160e01b600052602260045260246000fd5b50919050565b60006020828403121561152657600080fd5b5051919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60018060a01b038616815284602082015283604082015260806060820152600061083b60808301848661152d565b634e487b7160e01b600052604160045260246000fd5b60208152600061113160208301848661152d565b6000602082840312156115c057600080fd5b815167ffffffffffffffff808211156115d857600080fd5b818401915084601f8301126115ec57600080fd5b8151818111156115fe576115fe611584565b604051601f8201601f19908116603f0116810190838211818310171561162657611626611584565b8160405282815287602084870101111561163f57600080fd5b61083b8360208301602088016112ed565b601f82111561111d57600081815260208120601f850160051c810160208610156116775750805b601f850160051c820191505b8181101561169657828155600101611683565b505050505050565b815167ffffffffffffffff8111156116b8576116b8611584565b6116cc816116c684546114da565b84611650565b602080601f83116001811461170157600084156116e95750858301515b600019600386901b1c1916600185901b178555611696565b600085815260208120601f198616915b8281101561173057888601518255948401946001909101908401611711565b508582101561174e5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b602081526000610e1e6020830184611311565b634e487b7160e01b600052601160045260246000fd5b81810381811115610cf957610cf9611771565b600081518084526020808501945080840160005b838110156117d35781516001600160a01b0316875295820195908201906001016117ae565b509495945050505050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c084015261183f61010084018261179a565b905060a0840151601f198483030160e0850152610cf48282611311565b60006020828403121561186e57600080fd5b8151610e1e816112a2565b80820180821115610cf957610cf9611771565b60006020828403121561189e57600080fd5b81518015158114610e1e57600080fd5b600082516118c08184602087016112ed565b919091019291505056fea2646970667358221220f55c41c4ace58713cd09459c78638b0dfda9e2030caad918884a5f9b4e77a8a064736f6c63430008120033496e697469616c697a61626c653a20636f6e7472616374206973206e6f742069",
}

// TestMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use TestMessengerMetaData.ABI instead.
var TestMessengerABI = TestMessengerMetaData.ABI

// TestMessengerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestMessengerMetaData.Bin instead.
var TestMessengerBin = TestMessengerMetaData.Bin

// DeployTestMessenger deploys a new Ethereum contract, binding an instance of TestMessenger to it.
func DeployTestMessenger(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterRegistryAddress common.Address, teleporterManager common.Address) (common.Address, *types.Transaction, *TestMessenger, error) {
	parsed, err := TestMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestMessengerBin), backend, teleporterRegistryAddress, teleporterManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestMessenger{TestMessengerCaller: TestMessengerCaller{contract: contract}, TestMessengerTransactor: TestMessengerTransactor{contract: contract}, TestMessengerFilterer: TestMessengerFilterer{contract: contract}}, nil
}

// TestMessenger is an auto generated Go binding around an Ethereum contract.
type TestMessenger struct {
	TestMessengerCaller     // Read-only binding to the contract
	TestMessengerTransactor // Write-only binding to the contract
	TestMessengerFilterer   // Log filterer for contract events
}

// TestMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestMessengerSession struct {
	Contract     *TestMessenger    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestMessengerCallerSession struct {
	Contract *TestMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TestMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestMessengerTransactorSession struct {
	Contract     *TestMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestMessengerRaw struct {
	Contract *TestMessenger // Generic contract binding to access the raw methods on
}

// TestMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestMessengerCallerRaw struct {
	Contract *TestMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// TestMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestMessengerTransactorRaw struct {
	Contract *TestMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestMessenger creates a new instance of TestMessenger, bound to a specific deployed contract.
func NewTestMessenger(address common.Address, backend bind.ContractBackend) (*TestMessenger, error) {
	contract, err := bindTestMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestMessenger{TestMessengerCaller: TestMessengerCaller{contract: contract}, TestMessengerTransactor: TestMessengerTransactor{contract: contract}, TestMessengerFilterer: TestMessengerFilterer{contract: contract}}, nil
}

// NewTestMessengerCaller creates a new read-only instance of TestMessenger, bound to a specific deployed contract.
func NewTestMessengerCaller(address common.Address, caller bind.ContractCaller) (*TestMessengerCaller, error) {
	contract, err := bindTestMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestMessengerCaller{contract: contract}, nil
}

// NewTestMessengerTransactor creates a new write-only instance of TestMessenger, bound to a specific deployed contract.
func NewTestMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*TestMessengerTransactor, error) {
	contract, err := bindTestMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestMessengerTransactor{contract: contract}, nil
}

// NewTestMessengerFilterer creates a new log filterer instance of TestMessenger, bound to a specific deployed contract.
func NewTestMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*TestMessengerFilterer, error) {
	contract, err := bindTestMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestMessengerFilterer{contract: contract}, nil
}

// bindTestMessenger binds a generic wrapper to an already deployed contract.
func bindTestMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestMessenger *TestMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestMessenger.Contract.TestMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestMessenger *TestMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestMessenger.Contract.TestMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestMessenger *TestMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestMessenger.Contract.TestMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestMessenger *TestMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestMessenger *TestMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestMessenger *TestMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestMessenger.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 sourceBlockchainID) view returns(address, string)
func (_TestMessenger *TestMessengerCaller) GetCurrentMessage(opts *bind.CallOpts, sourceBlockchainID [32]byte) (common.Address, string, error) {
	var out []interface{}
	err := _TestMessenger.contract.Call(opts, &out, "getCurrentMessage", sourceBlockchainID)

	if err != nil {
		return *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 sourceBlockchainID) view returns(address, string)
func (_TestMessenger *TestMessengerSession) GetCurrentMessage(sourceBlockchainID [32]byte) (common.Address, string, error) {
	return _TestMessenger.Contract.GetCurrentMessage(&_TestMessenger.CallOpts, sourceBlockchainID)
}

// GetCurrentMessage is a free data retrieval call binding the contract method 0xb33fead4.
//
// Solidity: function getCurrentMessage(bytes32 sourceBlockchainID) view returns(address, string)
func (_TestMessenger *TestMessengerCallerSession) GetCurrentMessage(sourceBlockchainID [32]byte) (common.Address, string, error) {
	return _TestMessenger.Contract.GetCurrentMessage(&_TestMessenger.CallOpts, sourceBlockchainID)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_TestMessenger *TestMessengerCaller) GetMinTeleporterVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestMessenger.contract.Call(opts, &out, "getMinTeleporterVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_TestMessenger *TestMessengerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _TestMessenger.Contract.GetMinTeleporterVersion(&_TestMessenger.CallOpts)
}

// GetMinTeleporterVersion is a free data retrieval call binding the contract method 0xd2cc7a70.
//
// Solidity: function getMinTeleporterVersion() view returns(uint256)
func (_TestMessenger *TestMessengerCallerSession) GetMinTeleporterVersion() (*big.Int, error) {
	return _TestMessenger.Contract.GetMinTeleporterVersion(&_TestMessenger.CallOpts)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_TestMessenger *TestMessengerCaller) IsTeleporterAddressPaused(opts *bind.CallOpts, teleporterAddress common.Address) (bool, error) {
	var out []interface{}
	err := _TestMessenger.contract.Call(opts, &out, "isTeleporterAddressPaused", teleporterAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_TestMessenger *TestMessengerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _TestMessenger.Contract.IsTeleporterAddressPaused(&_TestMessenger.CallOpts, teleporterAddress)
}

// IsTeleporterAddressPaused is a free data retrieval call binding the contract method 0x97314297.
//
// Solidity: function isTeleporterAddressPaused(address teleporterAddress) view returns(bool)
func (_TestMessenger *TestMessengerCallerSession) IsTeleporterAddressPaused(teleporterAddress common.Address) (bool, error) {
	return _TestMessenger.Contract.IsTeleporterAddressPaused(&_TestMessenger.CallOpts, teleporterAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestMessenger *TestMessengerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestMessenger.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestMessenger *TestMessengerSession) Owner() (common.Address, error) {
	return _TestMessenger.Contract.Owner(&_TestMessenger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestMessenger *TestMessengerCallerSession) Owner() (common.Address, error) {
	return _TestMessenger.Contract.Owner(&_TestMessenger.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_TestMessenger *TestMessengerCaller) TeleporterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestMessenger.contract.Call(opts, &out, "teleporterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_TestMessenger *TestMessengerSession) TeleporterRegistry() (common.Address, error) {
	return _TestMessenger.Contract.TeleporterRegistry(&_TestMessenger.CallOpts)
}

// TeleporterRegistry is a free data retrieval call binding the contract method 0x1a7f5bec.
//
// Solidity: function teleporterRegistry() view returns(address)
func (_TestMessenger *TestMessengerCallerSession) TeleporterRegistry() (common.Address, error) {
	return _TestMessenger.Contract.TeleporterRegistry(&_TestMessenger.CallOpts)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_TestMessenger *TestMessengerTransactor) PauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "pauseTeleporterAddress", teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_TestMessenger *TestMessengerSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _TestMessenger.Contract.PauseTeleporterAddress(&_TestMessenger.TransactOpts, teleporterAddress)
}

// PauseTeleporterAddress is a paid mutator transaction binding the contract method 0x2b0d8f18.
//
// Solidity: function pauseTeleporterAddress(address teleporterAddress) returns()
func (_TestMessenger *TestMessengerTransactorSession) PauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _TestMessenger.Contract.PauseTeleporterAddress(&_TestMessenger.TransactOpts, teleporterAddress)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_TestMessenger *TestMessengerTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_TestMessenger *TestMessengerSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _TestMessenger.Contract.ReceiveTeleporterMessage(&_TestMessenger.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_TestMessenger *TestMessengerTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _TestMessenger.Contract.ReceiveTeleporterMessage(&_TestMessenger.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestMessenger *TestMessengerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestMessenger *TestMessengerSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestMessenger.Contract.RenounceOwnership(&_TestMessenger.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestMessenger *TestMessengerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestMessenger.Contract.RenounceOwnership(&_TestMessenger.TransactOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(bytes32)
func (_TestMessenger *TestMessengerTransactor) SendMessage(opts *bind.TransactOpts, destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "sendMessage", destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(bytes32)
func (_TestMessenger *TestMessengerSession) SendMessage(destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _TestMessenger.Contract.SendMessage(&_TestMessenger.TransactOpts, destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0xf63d09d7.
//
// Solidity: function sendMessage(bytes32 destinationBlockchainID, address destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message) returns(bytes32)
func (_TestMessenger *TestMessengerTransactorSession) SendMessage(destinationBlockchainID [32]byte, destinationAddress common.Address, feeTokenAddress common.Address, feeAmount *big.Int, requiredGasLimit *big.Int, message string) (*types.Transaction, error) {
	return _TestMessenger.Contract.SendMessage(&_TestMessenger.TransactOpts, destinationBlockchainID, destinationAddress, feeTokenAddress, feeAmount, requiredGasLimit, message)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestMessenger *TestMessengerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestMessenger *TestMessengerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestMessenger.Contract.TransferOwnership(&_TestMessenger.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestMessenger *TestMessengerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestMessenger.Contract.TransferOwnership(&_TestMessenger.TransactOpts, newOwner)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_TestMessenger *TestMessengerTransactor) UnpauseTeleporterAddress(opts *bind.TransactOpts, teleporterAddress common.Address) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "unpauseTeleporterAddress", teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_TestMessenger *TestMessengerSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _TestMessenger.Contract.UnpauseTeleporterAddress(&_TestMessenger.TransactOpts, teleporterAddress)
}

// UnpauseTeleporterAddress is a paid mutator transaction binding the contract method 0x4511243e.
//
// Solidity: function unpauseTeleporterAddress(address teleporterAddress) returns()
func (_TestMessenger *TestMessengerTransactorSession) UnpauseTeleporterAddress(teleporterAddress common.Address) (*types.Transaction, error) {
	return _TestMessenger.Contract.UnpauseTeleporterAddress(&_TestMessenger.TransactOpts, teleporterAddress)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_TestMessenger *TestMessengerTransactor) UpdateMinTeleporterVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _TestMessenger.contract.Transact(opts, "updateMinTeleporterVersion", version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_TestMessenger *TestMessengerSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _TestMessenger.Contract.UpdateMinTeleporterVersion(&_TestMessenger.TransactOpts, version)
}

// UpdateMinTeleporterVersion is a paid mutator transaction binding the contract method 0x5eb99514.
//
// Solidity: function updateMinTeleporterVersion(uint256 version) returns()
func (_TestMessenger *TestMessengerTransactorSession) UpdateMinTeleporterVersion(version *big.Int) (*types.Transaction, error) {
	return _TestMessenger.Contract.UpdateMinTeleporterVersion(&_TestMessenger.TransactOpts, version)
}

// TestMessengerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TestMessenger contract.
type TestMessengerInitializedIterator struct {
	Event *TestMessengerInitialized // Event containing the contract specifics and raw log

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
func (it *TestMessengerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerInitialized)
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
		it.Event = new(TestMessengerInitialized)
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
func (it *TestMessengerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerInitialized represents a Initialized event raised by the TestMessenger contract.
type TestMessengerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TestMessenger *TestMessengerFilterer) FilterInitialized(opts *bind.FilterOpts) (*TestMessengerInitializedIterator, error) {

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TestMessengerInitializedIterator{contract: _TestMessenger.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TestMessenger *TestMessengerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TestMessengerInitialized) (event.Subscription, error) {

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerInitialized)
				if err := _TestMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TestMessenger *TestMessengerFilterer) ParseInitialized(log types.Log) (*TestMessengerInitialized, error) {
	event := new(TestMessengerInitialized)
	if err := _TestMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessengerMinTeleporterVersionUpdatedIterator is returned from FilterMinTeleporterVersionUpdated and is used to iterate over the raw logs and unpacked data for MinTeleporterVersionUpdated events raised by the TestMessenger contract.
type TestMessengerMinTeleporterVersionUpdatedIterator struct {
	Event *TestMessengerMinTeleporterVersionUpdated // Event containing the contract specifics and raw log

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
func (it *TestMessengerMinTeleporterVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerMinTeleporterVersionUpdated)
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
		it.Event = new(TestMessengerMinTeleporterVersionUpdated)
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
func (it *TestMessengerMinTeleporterVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerMinTeleporterVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerMinTeleporterVersionUpdated represents a MinTeleporterVersionUpdated event raised by the TestMessenger contract.
type TestMessengerMinTeleporterVersionUpdated struct {
	OldMinTeleporterVersion *big.Int
	NewMinTeleporterVersion *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMinTeleporterVersionUpdated is a free log retrieval operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_TestMessenger *TestMessengerFilterer) FilterMinTeleporterVersionUpdated(opts *bind.FilterOpts, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (*TestMessengerMinTeleporterVersionUpdatedIterator, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return &TestMessengerMinTeleporterVersionUpdatedIterator{contract: _TestMessenger.contract, event: "MinTeleporterVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchMinTeleporterVersionUpdated is a free log subscription operation binding the contract event 0xa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d.
//
// Solidity: event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion)
func (_TestMessenger *TestMessengerFilterer) WatchMinTeleporterVersionUpdated(opts *bind.WatchOpts, sink chan<- *TestMessengerMinTeleporterVersionUpdated, oldMinTeleporterVersion []*big.Int, newMinTeleporterVersion []*big.Int) (event.Subscription, error) {

	var oldMinTeleporterVersionRule []interface{}
	for _, oldMinTeleporterVersionItem := range oldMinTeleporterVersion {
		oldMinTeleporterVersionRule = append(oldMinTeleporterVersionRule, oldMinTeleporterVersionItem)
	}
	var newMinTeleporterVersionRule []interface{}
	for _, newMinTeleporterVersionItem := range newMinTeleporterVersion {
		newMinTeleporterVersionRule = append(newMinTeleporterVersionRule, newMinTeleporterVersionItem)
	}

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "MinTeleporterVersionUpdated", oldMinTeleporterVersionRule, newMinTeleporterVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerMinTeleporterVersionUpdated)
				if err := _TestMessenger.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
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
func (_TestMessenger *TestMessengerFilterer) ParseMinTeleporterVersionUpdated(log types.Log) (*TestMessengerMinTeleporterVersionUpdated, error) {
	event := new(TestMessengerMinTeleporterVersionUpdated)
	if err := _TestMessenger.contract.UnpackLog(event, "MinTeleporterVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessengerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TestMessenger contract.
type TestMessengerOwnershipTransferredIterator struct {
	Event *TestMessengerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TestMessengerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerOwnershipTransferred)
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
		it.Event = new(TestMessengerOwnershipTransferred)
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
func (it *TestMessengerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerOwnershipTransferred represents a OwnershipTransferred event raised by the TestMessenger contract.
type TestMessengerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestMessenger *TestMessengerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestMessengerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestMessengerOwnershipTransferredIterator{contract: _TestMessenger.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestMessenger *TestMessengerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestMessengerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerOwnershipTransferred)
				if err := _TestMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TestMessenger *TestMessengerFilterer) ParseOwnershipTransferred(log types.Log) (*TestMessengerOwnershipTransferred, error) {
	event := new(TestMessengerOwnershipTransferred)
	if err := _TestMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessengerReceiveMessageIterator is returned from FilterReceiveMessage and is used to iterate over the raw logs and unpacked data for ReceiveMessage events raised by the TestMessenger contract.
type TestMessengerReceiveMessageIterator struct {
	Event *TestMessengerReceiveMessage // Event containing the contract specifics and raw log

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
func (it *TestMessengerReceiveMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerReceiveMessage)
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
		it.Event = new(TestMessengerReceiveMessage)
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
func (it *TestMessengerReceiveMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerReceiveMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerReceiveMessage represents a ReceiveMessage event raised by the TestMessenger contract.
type TestMessengerReceiveMessage struct {
	SourceBlockchainID  [32]byte
	OriginSenderAddress common.Address
	Message             string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReceiveMessage is a free log retrieval operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, string message)
func (_TestMessenger *TestMessengerFilterer) FilterReceiveMessage(opts *bind.FilterOpts, sourceBlockchainID [][32]byte, originSenderAddress []common.Address) (*TestMessengerReceiveMessageIterator, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "ReceiveMessage", sourceBlockchainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return &TestMessengerReceiveMessageIterator{contract: _TestMessenger.contract, event: "ReceiveMessage", logs: logs, sub: sub}, nil
}

// WatchReceiveMessage is a free log subscription operation binding the contract event 0x1f5c800b5f2b573929a7948f82a199c2a212851b53a6c5bd703ece23999d24aa.
//
// Solidity: event ReceiveMessage(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, string message)
func (_TestMessenger *TestMessengerFilterer) WatchReceiveMessage(opts *bind.WatchOpts, sink chan<- *TestMessengerReceiveMessage, sourceBlockchainID [][32]byte, originSenderAddress []common.Address) (event.Subscription, error) {

	var sourceBlockchainIDRule []interface{}
	for _, sourceBlockchainIDItem := range sourceBlockchainID {
		sourceBlockchainIDRule = append(sourceBlockchainIDRule, sourceBlockchainIDItem)
	}
	var originSenderAddressRule []interface{}
	for _, originSenderAddressItem := range originSenderAddress {
		originSenderAddressRule = append(originSenderAddressRule, originSenderAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "ReceiveMessage", sourceBlockchainIDRule, originSenderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerReceiveMessage)
				if err := _TestMessenger.contract.UnpackLog(event, "ReceiveMessage", log); err != nil {
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
// Solidity: event ReceiveMessage(bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, string message)
func (_TestMessenger *TestMessengerFilterer) ParseReceiveMessage(log types.Log) (*TestMessengerReceiveMessage, error) {
	event := new(TestMessengerReceiveMessage)
	if err := _TestMessenger.contract.UnpackLog(event, "ReceiveMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessengerSendMessageIterator is returned from FilterSendMessage and is used to iterate over the raw logs and unpacked data for SendMessage events raised by the TestMessenger contract.
type TestMessengerSendMessageIterator struct {
	Event *TestMessengerSendMessage // Event containing the contract specifics and raw log

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
func (it *TestMessengerSendMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerSendMessage)
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
		it.Event = new(TestMessengerSendMessage)
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
func (it *TestMessengerSendMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerSendMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerSendMessage represents a SendMessage event raised by the TestMessenger contract.
type TestMessengerSendMessage struct {
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
func (_TestMessenger *TestMessengerFilterer) FilterSendMessage(opts *bind.FilterOpts, destinationBlockchainID [][32]byte, destinationAddress []common.Address) (*TestMessengerSendMessageIterator, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "SendMessage", destinationBlockchainIDRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return &TestMessengerSendMessageIterator{contract: _TestMessenger.contract, event: "SendMessage", logs: logs, sub: sub}, nil
}

// WatchSendMessage is a free log subscription operation binding the contract event 0xa06eff1edd0c66b8dc96d086dda7ba263edf88d7417e6cb15073b5e7bff8a8ca.
//
// Solidity: event SendMessage(bytes32 indexed destinationBlockchainID, address indexed destinationAddress, address feeTokenAddress, uint256 feeAmount, uint256 requiredGasLimit, string message)
func (_TestMessenger *TestMessengerFilterer) WatchSendMessage(opts *bind.WatchOpts, sink chan<- *TestMessengerSendMessage, destinationBlockchainID [][32]byte, destinationAddress []common.Address) (event.Subscription, error) {

	var destinationBlockchainIDRule []interface{}
	for _, destinationBlockchainIDItem := range destinationBlockchainID {
		destinationBlockchainIDRule = append(destinationBlockchainIDRule, destinationBlockchainIDItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "SendMessage", destinationBlockchainIDRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerSendMessage)
				if err := _TestMessenger.contract.UnpackLog(event, "SendMessage", log); err != nil {
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
func (_TestMessenger *TestMessengerFilterer) ParseSendMessage(log types.Log) (*TestMessengerSendMessage, error) {
	event := new(TestMessengerSendMessage)
	if err := _TestMessenger.contract.UnpackLog(event, "SendMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessengerTeleporterAddressPausedIterator is returned from FilterTeleporterAddressPaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressPaused events raised by the TestMessenger contract.
type TestMessengerTeleporterAddressPausedIterator struct {
	Event *TestMessengerTeleporterAddressPaused // Event containing the contract specifics and raw log

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
func (it *TestMessengerTeleporterAddressPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerTeleporterAddressPaused)
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
		it.Event = new(TestMessengerTeleporterAddressPaused)
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
func (it *TestMessengerTeleporterAddressPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerTeleporterAddressPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerTeleporterAddressPaused represents a TeleporterAddressPaused event raised by the TestMessenger contract.
type TestMessengerTeleporterAddressPaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressPaused is a free log retrieval operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_TestMessenger *TestMessengerFilterer) FilterTeleporterAddressPaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*TestMessengerTeleporterAddressPausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &TestMessengerTeleporterAddressPausedIterator{contract: _TestMessenger.contract, event: "TeleporterAddressPaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressPaused is a free log subscription operation binding the contract event 0x933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c.
//
// Solidity: event TeleporterAddressPaused(address indexed teleporterAddress)
func (_TestMessenger *TestMessengerFilterer) WatchTeleporterAddressPaused(opts *bind.WatchOpts, sink chan<- *TestMessengerTeleporterAddressPaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "TeleporterAddressPaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerTeleporterAddressPaused)
				if err := _TestMessenger.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
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
func (_TestMessenger *TestMessengerFilterer) ParseTeleporterAddressPaused(log types.Log) (*TestMessengerTeleporterAddressPaused, error) {
	event := new(TestMessengerTeleporterAddressPaused)
	if err := _TestMessenger.contract.UnpackLog(event, "TeleporterAddressPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestMessengerTeleporterAddressUnpausedIterator is returned from FilterTeleporterAddressUnpaused and is used to iterate over the raw logs and unpacked data for TeleporterAddressUnpaused events raised by the TestMessenger contract.
type TestMessengerTeleporterAddressUnpausedIterator struct {
	Event *TestMessengerTeleporterAddressUnpaused // Event containing the contract specifics and raw log

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
func (it *TestMessengerTeleporterAddressUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestMessengerTeleporterAddressUnpaused)
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
		it.Event = new(TestMessengerTeleporterAddressUnpaused)
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
func (it *TestMessengerTeleporterAddressUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestMessengerTeleporterAddressUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestMessengerTeleporterAddressUnpaused represents a TeleporterAddressUnpaused event raised by the TestMessenger contract.
type TestMessengerTeleporterAddressUnpaused struct {
	TeleporterAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTeleporterAddressUnpaused is a free log retrieval operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_TestMessenger *TestMessengerFilterer) FilterTeleporterAddressUnpaused(opts *bind.FilterOpts, teleporterAddress []common.Address) (*TestMessengerTeleporterAddressUnpausedIterator, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.FilterLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return &TestMessengerTeleporterAddressUnpausedIterator{contract: _TestMessenger.contract, event: "TeleporterAddressUnpaused", logs: logs, sub: sub}, nil
}

// WatchTeleporterAddressUnpaused is a free log subscription operation binding the contract event 0x844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c3.
//
// Solidity: event TeleporterAddressUnpaused(address indexed teleporterAddress)
func (_TestMessenger *TestMessengerFilterer) WatchTeleporterAddressUnpaused(opts *bind.WatchOpts, sink chan<- *TestMessengerTeleporterAddressUnpaused, teleporterAddress []common.Address) (event.Subscription, error) {

	var teleporterAddressRule []interface{}
	for _, teleporterAddressItem := range teleporterAddress {
		teleporterAddressRule = append(teleporterAddressRule, teleporterAddressItem)
	}

	logs, sub, err := _TestMessenger.contract.WatchLogs(opts, "TeleporterAddressUnpaused", teleporterAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestMessengerTeleporterAddressUnpaused)
				if err := _TestMessenger.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
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
func (_TestMessenger *TestMessengerFilterer) ParseTeleporterAddressUnpaused(log types.Log) (*TestMessengerTeleporterAddressUnpaused, error) {
	event := new(TestMessengerTeleporterAddressUnpaused)
	if err := _TestMessenger.contract.UnpackLog(event, "TeleporterAddressUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
