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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teleporterManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeTokenDestinationAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20ContractAddress_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferToDestination\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINT_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationBurnedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20ContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenDestinationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"transferToDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162001fad38038062001fad8339810160408190526200003591620004e5565b6001600055848484848383816001600160a01b038116620000c35760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa1580156200010e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013491906200054c565b60025550620001433362000399565b6200014e81620003eb565b50829050620001b45760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e536f757263653a207a65726f2064657374696e6174696f6e20626c60448201526a1bd8dad8da185a5b88125160aa1b6064820152608401620000ba565b7302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000207573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200022d91906200054c565b8203620002955760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e536f757263653a2063616e6e6f742062726964676520776974682060448201526e39b0b6b290313637b1b5b1b430b4b760891b6064820152608401620000ba565b60a08290526001600160a01b038116620003095760405162461bcd60e51b815260206004820152602e60248201527f546f6b656e536f757263653a207a65726f2064657374696e6174696f6e20636f60448201526d6e7472616374206164647265737360901b6064820152608401620000ba565b6001600160a01b0390811660c05284169250620003829150505760405162461bcd60e51b815260206004820152602d60248201527f4552433230546f6b656e536f757263653a207a65726f20455243323020636f6e60448201526c7472616374206164647265737360981b6064820152608401620000ba565b6001600160a01b031660e052506200056692505050565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b620003f56200046a565b6001600160a01b0381166200045c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401620000ba565b620004678162000399565b50565b6003546001600160a01b03163314620004c65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401620000ba565b565b80516001600160a01b0381168114620004e057600080fd5b919050565b600080600080600060a08688031215620004fe57600080fd5b6200050986620004c8565b94506200051960208701620004c8565b9350604086015192506200053060608701620004c8565b91506200054060808701620004c8565b90509295509295909350565b6000602082840312156200055f57600080fd5b5051919050565b60805160a05160c05160e0516119c7620005e66000396000818161027b0152818161056e0152818161067b01526112ca015260008181610239015281816106410152610f870152600081816101630152818161061b0152610f0901526000818161010a0152818161079f015281816109d6015261109b01526119c76000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638da5cb5b11610097578063c868efaa11610066578063c868efaa1461025b578063d2cc7a701461026e578063e486df1514610276578063f2fde38b1461029d57600080fd5b80638da5cb5b146101dd57806397314297146101ee578063b6171f731461022a578063b8c9091a1461023457600080fd5b806355db3e9e116100d357806355db3e9e146101a65780635eb99514146101af578063715018a6146101c257806387a2edba146101ca57600080fd5b80631a7f5bec146101055780632b0d8f181461014957806341d3014d1461015e5780634511243e14610193575b600080fd5b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b61015c6101573660046115a0565b6102b0565b005b6101857f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610140565b61015c6101a13660046115a0565b6103b5565b61018560045481565b61015c6101bd3660046115bd565b6104b2565b61015c6104c6565b61015c6101d83660046115d6565b6104da565b6003546001600160a01b031661012c565b61021a6101fc3660046115a0565b6001600160a01b031660009081526001602052604090205460ff1690565b6040519015158152602001610140565b610185620186a081565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b61015c61026936600461166f565b61078a565b600254610185565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b61015c6102ab3660046115a0565b610954565b6102b86109ca565b6001600160a01b0381166102e75760405162461bcd60e51b81526004016102de906116f8565b60405180910390fd5b6001600160a01b03811660009081526001602052604090205460ff16156103665760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b60648201526084016102de565b6001600160a01b0381166000818152600160208190526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b6103bd6109ca565b6001600160a01b0381166103e35760405162461bcd60e51b81526004016102de906116f8565b6001600160a01b03811660009081526001602052604090205460ff1661045d5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b60648201526084016102de565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600160205260409020805460ff19169055565b6104ba6109ca565b6104c3816109d2565b50565b6104ce610b72565b6104d86000610bcc565b565b6104e2610c1e565b6001600160a01b0385166105085760405162461bcd60e51b81526004016102de90611746565b600084116105675760405162461bcd60e51b815260206004820152602660248201527f4552433230546f6b656e536f757263653a207a65726f207472616e7366657220604482015265185b5bdd5b9d60d21b60648201526084016102de565b60006105937f000000000000000000000000000000000000000000000000000000000000000086610c77565b90508381116105fb5760405162461bcd60e51b815260206004820152602e60248201527f4552433230546f6b656e536f757263653a20696e73756666696369656e74206160448201526d191a9d5cdd195908185b5bdd5b9d60921b60648201526084016102de565b600061060785836117a4565b905060006107266040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020018a8152508152602001620186a08152602001878780806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250505090825250604080516001600160a01b038d166020808301919091528183018890528251808303840181526060909201909252910152610de1565b905080886001600160a01b0316336001600160a01b03167f6cf14fdf618c440df3de7de7dcacf59541a464e55f360cbe73724c12e0c4cf998560405161076e91815260200190565b60405180910390a45050506107836001600055565b5050505050565b610792610c1e565b6002546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015610809573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061082d91906117b7565b10156108945760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b60648201526084016102de565b61089d336101fc565b156109035760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b60648201526084016102de565b610944848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610f0792505050565b61094e6001600055565b50505050565b61095c610b72565b6001600160a01b0381166109c15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102de565b6104c381610bcc565b6104d8610b72565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5691906117b7565b60025490915081831115610ac65760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b60648201526084016102de565b808311610b3b5760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e0060648201526084016102de565b6002839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6003546001600160a01b031633146104d85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102de565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600260005403610c705760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016102de565b6002600055565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015610cc0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce491906117b7565b9050610cfb6001600160a01b03851633308661102b565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610d42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d6691906117b7565b9050818111610dcc5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016102de565b610dd682826117a4565b925050505b92915050565b600080610dec611096565b60408401516020015190915015610e91576040830151516001600160a01b0316610e6e5760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a207a65726f206665652060448201526c746f6b656e206164647265737360981b60648201526084016102de565b604083015160208101519051610e91916001600160a01b039091169083906111aa565b604051630624488560e41b81526001600160a01b03821690636244885090610ebd908690600401611864565b6020604051808303816000875af1158015610edc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0091906117b7565b9392505050565b7f00000000000000000000000000000000000000000000000000000000000000008314610f855760405162461bcd60e51b815260206004820152602660248201527f546f6b656e536f757263653a20696e76616c69642064657374696e6174696f6e6044820152651031b430b4b760d11b60648201526084016102de565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316146110065760405162461bcd60e51b815260206004820181905260248201527f546f6b656e536f757263653a20756e617574686f72697a65642073656e64657260448201526064016102de565b6000808280602001905181019061101d91906118e2565b91509150610783828261125c565b6040516001600160a01b038085166024830152831660448201526064810182905261094e9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526112f4565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061111b9190611910565b905061113f816001600160a01b031660009081526001602052604090205460ff1690565b156111a55760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b60648201526084016102de565b919050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa1580156111fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061121f91906117b7565b611229919061192d565b6040516001600160a01b03851660248201526044810182905290915061094e90859063095ea7b360e01b9060640161105f565b6001600160a01b0382166112825760405162461bcd60e51b81526004016102de90611746565b816001600160a01b03167f55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407826040516112bd91815260200190565b60405180910390a26112f07f000000000000000000000000000000000000000000000000000000000000000083836113cb565b5050565b6000611349826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166113fb9092919063ffffffff16565b8051909150156113c657808060200190518101906113679190611940565b6113c65760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016102de565b505050565b6040516001600160a01b0383166024820152604481018290526113c690849063a9059cbb60e01b9060640161105f565b606061140a8484600085611412565b949350505050565b6060824710156114735760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016102de565b600080866001600160a01b0316858760405161148f9190611962565b60006040518083038185875af1925050503d80600081146114cc576040519150601f19603f3d011682016040523d82523d6000602084013e6114d1565b606091505b50915091506114e2878383876114ed565b979650505050505050565b6060831561155c578251600003611555576001600160a01b0385163b6115555760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102de565b508161140a565b61140a83838151156115715781518083602001fd5b8060405162461bcd60e51b81526004016102de919061197e565b6001600160a01b03811681146104c357600080fd5b6000602082840312156115b257600080fd5b8135610f008161158b565b6000602082840312156115cf57600080fd5b5035919050565b6000806000806000608086880312156115ee57600080fd5b85356115f98161158b565b94506020860135935060408601359250606086013567ffffffffffffffff8082111561162457600080fd5b818801915088601f83011261163857600080fd5b81358181111561164757600080fd5b8960208260051b850101111561165c57600080fd5b9699959850939650602001949392505050565b6000806000806060858703121561168557600080fd5b8435935060208501356116978161158b565b9250604085013567ffffffffffffffff808211156116b457600080fd5b818701915087601f8301126116c857600080fd5b8135818111156116d757600080fd5b8860208285010111156116e957600080fd5b95989497505060200194505050565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b60208082526028908201527f4552433230546f6b656e536f757263653a207a65726f20726563697069656e74604082015267206164647265737360c01b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b81810381811115610ddb57610ddb61178e565b6000602082840312156117c957600080fd5b5051919050565b600081518084526020808501945080840160005b838110156118095781516001600160a01b0316875295820195908201906001016117e4565b509495945050505050565b60005b8381101561182f578181015183820152602001611817565b50506000910152565b60008151808452611850816020860160208601611814565b601f01601f19169290920160200192915050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c08401526118c56101008401826117d0565b905060a0840151601f198483030160e0850152610dd68282611838565b600080604083850312156118f557600080fd5b82516119008161158b565b6020939093015192949293505050565b60006020828403121561192257600080fd5b8151610f008161158b565b80820180821115610ddb57610ddb61178e565b60006020828403121561195257600080fd5b81518015158114610f0057600080fd5b60008251611974818460208701611814565b9190910192915050565b602081526000610f00602083018461183856fea26469706673582212200b049d4dd835014e7abaeccd99c9095398ba50b5e7a08e79b209f26caff7554c64736f6c63430008120033",
}

// ERC20TokenSourceABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenSourceMetaData.ABI instead.
var ERC20TokenSourceABI = ERC20TokenSourceMetaData.ABI

// ERC20TokenSourceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenSourceMetaData.Bin instead.
var ERC20TokenSourceBin = ERC20TokenSourceMetaData.Bin

// DeployERC20TokenSource deploys a new Ethereum contract, binding an instance of ERC20TokenSource to it.
func DeployERC20TokenSource(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterRegistryAddress common.Address, teleporterManager common.Address, destinationBlockchainID_ [32]byte, nativeTokenDestinationAddress_ common.Address, erc20ContractAddress_ common.Address) (common.Address, *types.Transaction, *ERC20TokenSource, error) {
	parsed, err := ERC20TokenSourceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenSourceBin), backend, teleporterRegistryAddress, teleporterManager, destinationBlockchainID_, nativeTokenDestinationAddress_, erc20ContractAddress_)
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
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactor) ReceiveTeleporterMessage(opts *bind.TransactOpts, sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenSource.contract.Transact(opts, "receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenSource *ERC20TokenSourceSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.ReceiveTeleporterMessage(&_ERC20TokenSource.TransactOpts, sourceBlockchainID, originSenderAddress, message)
}

// ReceiveTeleporterMessage is a paid mutator transaction binding the contract method 0xc868efaa.
//
// Solidity: function receiveTeleporterMessage(bytes32 sourceBlockchainID, address originSenderAddress, bytes message) returns()
func (_ERC20TokenSource *ERC20TokenSourceTransactorSession) ReceiveTeleporterMessage(sourceBlockchainID [32]byte, originSenderAddress common.Address, message []byte) (*types.Transaction, error) {
	return _ERC20TokenSource.Contract.ReceiveTeleporterMessage(&_ERC20TokenSource.TransactOpts, sourceBlockchainID, originSenderAddress, message)
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
	TeleporterMessageID [32]byte
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTransferToDestination is a free log retrieval operation binding the contract event 0x6cf14fdf618c440df3de7de7dcacf59541a464e55f360cbe73724c12e0c4cf99.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, bytes32 indexed teleporterMessageID, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) FilterTransferToDestination(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, teleporterMessageID [][32]byte) (*ERC20TokenSourceTransferToDestinationIterator, error) {

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

// WatchTransferToDestination is a free log subscription operation binding the contract event 0x6cf14fdf618c440df3de7de7dcacf59541a464e55f360cbe73724c12e0c4cf99.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, bytes32 indexed teleporterMessageID, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) WatchTransferToDestination(opts *bind.WatchOpts, sink chan<- *ERC20TokenSourceTransferToDestination, sender []common.Address, recipient []common.Address, teleporterMessageID [][32]byte) (event.Subscription, error) {

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

// ParseTransferToDestination is a log parse operation binding the contract event 0x6cf14fdf618c440df3de7de7dcacf59541a464e55f360cbe73724c12e0c4cf99.
//
// Solidity: event TransferToDestination(address indexed sender, address indexed recipient, bytes32 indexed teleporterMessageID, uint256 amount)
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
// Solidity: event UnlockTokens(address indexed recipient, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) FilterUnlockTokens(opts *bind.FilterOpts, recipient []common.Address) (*ERC20TokenSourceUnlockTokensIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.FilterLogs(opts, "UnlockTokens", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenSourceUnlockTokensIterator{contract: _ERC20TokenSource.contract, event: "UnlockTokens", logs: logs, sub: sub}, nil
}

// WatchUnlockTokens is a free log subscription operation binding the contract event 0x55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb407.
//
// Solidity: event UnlockTokens(address indexed recipient, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) WatchUnlockTokens(opts *bind.WatchOpts, sink chan<- *ERC20TokenSourceUnlockTokens, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20TokenSource.contract.WatchLogs(opts, "UnlockTokens", recipientRule)
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
// Solidity: event UnlockTokens(address indexed recipient, uint256 amount)
func (_ERC20TokenSource *ERC20TokenSourceFilterer) ParseUnlockTokens(log types.Log) (*ERC20TokenSourceUnlockTokens, error) {
	event := new(ERC20TokenSourceUnlockTokens)
	if err := _ERC20TokenSource.contract.UnpackLog(event, "UnlockTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
