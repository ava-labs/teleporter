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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationBlockchainID_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nativeTokenDestinationAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc20ContractAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenMultiplier_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"multiplyOnSend_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurnTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinTeleporterVersion\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinTeleporterVersion\",\"type\":\"uint256\"}],\"name\":\"MinTeleporterVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"TeleporterAddressUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"teleporterMessageID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferToDestination\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_NATIVE_TOKENS_REQUIRED_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationBlockchainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destinationBurnedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20ContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinTeleporterVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"isTeleporterAddressPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplyOnSend\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenDestinationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"pauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceBlockchainID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"originSenderAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveTeleporterMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teleporterRegistry\",\"outputs\":[{\"internalType\":\"contractTeleporterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"allowedRelayerAddresses\",\"type\":\"address[]\"}],\"name\":\"transferToDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teleporterAddress\",\"type\":\"address\"}],\"name\":\"unpauseTeleporterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateMinTeleporterVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b506040516200235f3803806200235f833981016040819052620000359162000468565b600160005585858582806001600160a01b038116620000c15760405162461bcd60e51b815260206004820152603760248201527f54656c65706f727465725570677261646561626c653a207a65726f2074656c6560448201527f706f72746572207265676973747279206164647265737300000000000000000060648201526084015b60405180910390fd5b6001600160a01b03811660808190526040805163301fd1f560e21b8152905163c07f47d4916004808201926020929091908290030181865afa1580156200010c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001329190620004df565b600255506200014133620003f9565b5081620001a55760405162461bcd60e51b815260206004820152602b60248201527f546f6b656e536f757263653a207a65726f2064657374696e6174696f6e20626c60448201526a1bd8dad8da185a5b88125160aa1b6064820152608401620000b8565b7302000000000000000000000000000000000000056001600160a01b0316634213cf786040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001f8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200021e9190620004df565b8203620002865760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e536f757263653a2063616e6e6f742062726964676520776974682060448201526e39b0b6b290313637b1b5b1b430b4b760891b6064820152608401620000b8565b60a08290526001600160a01b038116620002fa5760405162461bcd60e51b815260206004820152602e60248201527f546f6b656e536f757263653a207a65726f2064657374696e6174696f6e20636f60448201526d6e7472616374206164647265737360901b6064820152608401620000b8565b6001600160a01b0390811660c052851691506200037290505760405162461bcd60e51b815260206004820152602d60248201527f4552433230546f6b656e536f757263653a207a65726f20455243323020636f6e60448201526c7472616374206164647265737360981b6064820152608401620000b8565b6001600160a01b03831660e0526000829003620003e15760405162461bcd60e51b815260206004820152602660248201527f4552433230546f6b656e536f757263653a207a65726f20746f6b656e4d756c7460448201526534b83634b2b960d11b6064820152608401620000b8565b6101009190915215156101205250620004f992505050565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80516001600160a01b03811681146200046357600080fd5b919050565b60008060008060008060c087890312156200048257600080fd5b6200048d876200044b565b955060208701519450620004a4604088016200044b565b9350620004b4606088016200044b565b92506080870151915060a08701518015158114620004d157600080fd5b809150509295509295509295565b600060208284031215620004f257600080fd5b5051919050565b60805160a05160c05160e0516101005161012051611d83620005dc600039600081816101840152818161066801528181611447015261153b0152600081816102a801528181610691015281816106c201528181611470015281816114a10152818161156401526115950152600081816102ea0152818161059a015281816106320152818161076801528181611511015261166c0152600081816102810152818161072e01526111a60152600081816101bb01528181610708015261112801526000818161012b015281816108eb01528181610b220152610dc80152611d836000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80638da5cb5b116100ad578063c868efaa11610071578063c868efaa146102ca578063d2cc7a70146102dd578063e486df15146102e5578063f2fde38b1461030c578063fccc28131461031f57600080fd5b80638da5cb5b146102355780639731429714610246578063b6171f7314610272578063b8c9091a1461027c578063ba3f5a12146102a357600080fd5b80634511243e116100f45780634511243e146101eb57806355db3e9e146101fe5780635eb9951414610207578063715018a61461021a57806387a2edba1461022257600080fd5b80631a7f5bec146101265780632b0d8f181461016a578063355fceb71461017f57806341d3014d146101b6575b600080fd5b61014d7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b61017d610178366004611829565b61032f565b005b6101a67f000000000000000000000000000000000000000000000000000000000000000081565b6040519015158152602001610161565b6101dd7f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610161565b61017d6101f9366004611829565b610434565b6101dd60045481565b61017d61021536600461184d565b610531565b61017d610545565b61017d610230366004611866565b610559565b6003546001600160a01b031661014d565b6101a6610254366004611829565b6001600160a01b031660009081526001602052604090205460ff1690565b6101dd620186a081565b61014d7f000000000000000000000000000000000000000000000000000000000000000081565b6101dd7f000000000000000000000000000000000000000000000000000000000000000081565b61017d6102d83660046118ff565b6108d6565b6002546101dd565b61014d7f000000000000000000000000000000000000000000000000000000000000000081565b61017d61031a366004611829565b610aa0565b61014d62010203600160981b0181565b610337610b16565b6001600160a01b0381166103665760405162461bcd60e51b815260040161035d90611988565b60405180910390fd5b6001600160a01b03811660009081526001602052604090205460ff16156103e55760405162461bcd60e51b815260206004820152602d60248201527f54656c65706f727465725570677261646561626c653a2061646472657373206160448201526c1b1c9958591e481c185d5cd959609a1b606482015260840161035d565b6001600160a01b0381166000818152600160208190526040808320805460ff1916909217909155517f933f93e57a222e6330362af8b376d0a8725b6901e9a2fb86d00f169702b28a4c9190a250565b61043c610b16565b6001600160a01b0381166104625760405162461bcd60e51b815260040161035d90611988565b6001600160a01b03811660009081526001602052604090205460ff166104dc5760405162461bcd60e51b815260206004820152602960248201527f54656c65706f727465725570677261646561626c653a2061646472657373206e6044820152681bdd081c185d5cd95960ba1b606482015260840161035d565b6040516001600160a01b038216907f844e2f3154214672229235858fd029d1dfd543901c6d05931f0bc2480a2d72c390600090a26001600160a01b03166000908152600160205260409020805460ff19169055565b610539610b16565b61054281610b1e565b50565b61054d610cbe565b6105576000610d18565b565b610561610d6a565b600061056b610dc3565b90506001600160a01b0386166105935760405162461bcd60e51b815260040161035d906119d6565b60006105bf7f000000000000000000000000000000000000000000000000000000000000000087610ed7565b90508481116106275760405162461bcd60e51b815260206004820152602e60248201527f4552433230546f6b656e536f757263653a20696e73756666696369656e74206160448201526d191a9d5cdd195908185b5bdd5b9d60921b606482015260840161035d565b8415610658576106587f00000000000000000000000000000000000000000000000000000000000000008387611041565b60006106648683611a34565b90507f0000000000000000000000000000000000000000000000000000000000000000156106bd576106b67f000000000000000000000000000000000000000000000000000000000000000082611a47565b90506106ea565b6106e77f000000000000000000000000000000000000000000000000000000000000000082611a5e565b90505b6000836001600160a01b031663624488506040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200160405180604001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020018c8152508152602001620186a08152602001898980806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250505090825250604080516001600160a01b038f166020808301919091529181018890529101906060016040516020818303038152906040528152506040518263ffffffff1660e01b815260040161082e9190611b14565b6020604051808303816000875af115801561084d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108719190611b92565b905080896001600160a01b0316336001600160a01b03167f6cf14fdf618c440df3de7de7dcacf59541a464e55f360cbe73724c12e0c4cf99856040516108b991815260200190565b60405180910390a4505050506108cf6001600055565b5050505050565b6108de610d6a565b6002546001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016634c1f08ce336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015610955573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109799190611b92565b10156109e05760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201526f32b632b837b93a32b91039b2b73232b960811b606482015260840161035d565b6109e933610254565b15610a4f5760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881859191c995cdcc81c185d5cd95960821b606482015260840161035d565b610a90848484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061112692505050565b610a9a6001600055565b50505050565b610aa8610cbe565b6001600160a01b038116610b0d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161035d565b61054281610d18565b610557610cbe565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c07f47d46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba29190611b92565b60025490915081831115610c125760405162461bcd60e51b815260206004820152603160248201527f54656c65706f727465725570677261646561626c653a20696e76616c6964205460448201527032b632b837b93a32b9103b32b939b4b7b760791b606482015260840161035d565b808311610c875760405162461bcd60e51b815260206004820152603f60248201527f54656c65706f727465725570677261646561626c653a206e6f7420677265617460448201527f6572207468616e2063757272656e74206d696e696d756d2076657273696f6e00606482015260840161035d565b6002839055604051839082907fa9a7ef57e41f05b4c15480842f5f0c27edfcbb553fed281f7c4068452cc1c02d90600090a3505050565b6003546001600160a01b031633146105575760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161035d565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600260005403610dbc5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161035d565b6002600055565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d820e64f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e489190611bab565b9050610e6c816001600160a01b031660009081526001602052604090205460ff1690565b15610ed25760405162461bcd60e51b815260206004820152603060248201527f54656c65706f727465725570677261646561626c653a2054656c65706f72746560448201526f1c881cd95b991a5b99c81c185d5cd95960821b606482015260840161035d565b919050565b6040516370a0823160e01b815230600482015260009081906001600160a01b038516906370a0823190602401602060405180830381865afa158015610f20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f449190611b92565b9050610f5b6001600160a01b03851633308661130e565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015610fa2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fc69190611b92565b905081811161102c5760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b606482015260840161035d565b6110368282611a34565b925050505b92915050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e90604401602060405180830381865afa158015611092573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b69190611b92565b6110c09190611bc8565b6040516001600160a01b038516602482015260448101829052909150610a9a90859063095ea7b360e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152611346565b7f000000000000000000000000000000000000000000000000000000000000000083146111a45760405162461bcd60e51b815260206004820152602660248201527f546f6b656e536f757263653a20696e76616c69642064657374696e6174696f6e6044820152651031b430b4b760d11b606482015260840161035d565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316146112255760405162461bcd60e51b815260206004820181905260248201527f546f6b656e536f757263653a20756e617574686f72697a65642073656e646572604482015260640161035d565b6000808280602001905181019061123c9190611bf1565b9092509050600082600181111561125557611255611cb8565b0361128657600080828060200190518101906112719190611cce565b9150915061127f828261141d565b50506108cf565b600182600181111561129a5761129a611cb8565b036112c6576000818060200190518101906112b59190611b92565b90506112c081611537565b506108cf565b60405162461bcd60e51b815260206004820152601b60248201527f546f6b656e536f757263653a20696e76616c696420616374696f6e0000000000604482015260640161035d565b6040516001600160a01b0380851660248301528316604482015260648101829052610a9a9085906323b872dd60e01b906084016110ef565b600061139b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166115ed9092919063ffffffff16565b80519091501561141857808060200190518101906113b99190611cfc565b6114185760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161035d565b505050565b6001600160a01b0382166114435760405162461bcd60e51b815260040161035d906119d6565b60007f00000000000000000000000000000000000000000000000000000000000000001561149c576114957f000000000000000000000000000000000000000000000000000000000000000083611a5e565b90506114c9565b6114c67f000000000000000000000000000000000000000000000000000000000000000083611a47565b90505b826001600160a01b03167f55aaef8fd8c07238c3618a93c8a1627194187d3b0952908e58f2ab0f944fb4078260405161150491815260200190565b60405180910390a26114187f00000000000000000000000000000000000000000000000000000000000000008483611604565b60007f000000000000000000000000000000000000000000000000000000000000000015611590576115897f000000000000000000000000000000000000000000000000000000000000000083611a5e565b90506115bd565b6115ba7f000000000000000000000000000000000000000000000000000000000000000083611a47565b90505b6004548111156115e9576000600454826115d79190611a34565b90506115e281611634565b5060048190555b5050565b60606115fc848460008561169b565b949350505050565b6040516001600160a01b03831660248201526044810182905261141890849063a9059cbb60e01b906064016110ef565b6040518181527f2cd3fd70cd5a5d6d805e90d22741aa1a84590ace7cf01b244719558d266143829060200160405180910390a16105427f000000000000000000000000000000000000000000000000000000000000000062010203600160981b0183611604565b6060824710156116fc5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161035d565b600080866001600160a01b031685876040516117189190611d1e565b60006040518083038185875af1925050503d8060008114611755576040519150601f19603f3d011682016040523d82523d6000602084013e61175a565b606091505b509150915061176b87838387611776565b979650505050505050565b606083156117e55782516000036117de576001600160a01b0385163b6117de5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161035d565b50816115fc565b6115fc83838151156117fa5781518083602001fd5b8060405162461bcd60e51b815260040161035d9190611d3a565b6001600160a01b038116811461054257600080fd5b60006020828403121561183b57600080fd5b813561184681611814565b9392505050565b60006020828403121561185f57600080fd5b5035919050565b60008060008060006080868803121561187e57600080fd5b853561188981611814565b94506020860135935060408601359250606086013567ffffffffffffffff808211156118b457600080fd5b818801915088601f8301126118c857600080fd5b8135818111156118d757600080fd5b8960208260051b85010111156118ec57600080fd5b9699959850939650602001949392505050565b6000806000806060858703121561191557600080fd5b84359350602085013561192781611814565b9250604085013567ffffffffffffffff8082111561194457600080fd5b818701915087601f83011261195857600080fd5b81358181111561196757600080fd5b88602082850101111561197957600080fd5b95989497505060200194505050565b6020808252602e908201527f54656c65706f727465725570677261646561626c653a207a65726f2054656c6560408201526d706f72746572206164647265737360901b606082015260800190565b60208082526028908201527f4552433230546f6b656e536f757263653a207a65726f20726563697069656e74604082015267206164647265737360c01b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b8181038181111561103b5761103b611a1e565b808202811582820484141761103b5761103b611a1e565b600082611a7b57634e487b7160e01b600052601260045260246000fd5b500490565b600081518084526020808501945080840160005b83811015611ab95781516001600160a01b031687529582019590820190600101611a94565b509495945050505050565b60005b83811015611adf578181015183820152602001611ac7565b50506000910152565b60008151808452611b00816020860160208601611ac4565b601f01601f19169290920160200192915050565b60208152815160208201526000602083015160018060a01b03808216604085015260408501519150808251166060850152506020810151608084015250606083015160a0830152608083015160e060c0840152611b75610100840182611a80565b905060a0840151601f198483030160e08501526110368282611ae8565b600060208284031215611ba457600080fd5b5051919050565b600060208284031215611bbd57600080fd5b815161184681611814565b8082018082111561103b5761103b611a1e565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215611c0457600080fd5b825160028110611c1357600080fd5b602084015190925067ffffffffffffffff80821115611c3157600080fd5b818501915085601f830112611c4557600080fd5b815181811115611c5757611c57611bdb565b604051601f8201601f19908116603f01168101908382118183101715611c7f57611c7f611bdb565b81604052828152886020848701011115611c9857600080fd5b611ca9836020830160208801611ac4565b80955050505050509250929050565b634e487b7160e01b600052602160045260246000fd5b60008060408385031215611ce157600080fd5b8251611cec81611814565b6020939093015192949293505050565b600060208284031215611d0e57600080fd5b8151801515811461184657600080fd5b60008251611d30818460208701611ac4565b9190910192915050565b6020815260006118466020830184611ae856fea2646970667358221220d28e6615a73bf690a4a4ed679e679c66c417b84652ca36b5ef92aa729fea692d64736f6c63430008120033",
}

// ERC20TokenSourceABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TokenSourceMetaData.ABI instead.
var ERC20TokenSourceABI = ERC20TokenSourceMetaData.ABI

// ERC20TokenSourceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20TokenSourceMetaData.Bin instead.
var ERC20TokenSourceBin = ERC20TokenSourceMetaData.Bin

// DeployERC20TokenSource deploys a new Ethereum contract, binding an instance of ERC20TokenSource to it.
func DeployERC20TokenSource(auth *bind.TransactOpts, backend bind.ContractBackend, teleporterRegistryAddress common.Address, destinationBlockchainID_ [32]byte, nativeTokenDestinationAddress_ common.Address, erc20ContractAddress_ common.Address, tokenMultiplier_ *big.Int, multiplyOnSend_ bool) (common.Address, *types.Transaction, *ERC20TokenSource, error) {
	parsed, err := ERC20TokenSourceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20TokenSourceBin), backend, teleporterRegistryAddress, destinationBlockchainID_, nativeTokenDestinationAddress_, erc20ContractAddress_, tokenMultiplier_, multiplyOnSend_)
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

// MultiplyOnSend is a free data retrieval call binding the contract method 0x355fceb7.
//
// Solidity: function multiplyOnSend() view returns(bool)
func (_ERC20TokenSource *ERC20TokenSourceCaller) MultiplyOnSend(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "multiplyOnSend")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MultiplyOnSend is a free data retrieval call binding the contract method 0x355fceb7.
//
// Solidity: function multiplyOnSend() view returns(bool)
func (_ERC20TokenSource *ERC20TokenSourceSession) MultiplyOnSend() (bool, error) {
	return _ERC20TokenSource.Contract.MultiplyOnSend(&_ERC20TokenSource.CallOpts)
}

// MultiplyOnSend is a free data retrieval call binding the contract method 0x355fceb7.
//
// Solidity: function multiplyOnSend() view returns(bool)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) MultiplyOnSend() (bool, error) {
	return _ERC20TokenSource.Contract.MultiplyOnSend(&_ERC20TokenSource.CallOpts)
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

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceCaller) TokenMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenSource.contract.Call(opts, &out, "tokenMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceSession) TokenMultiplier() (*big.Int, error) {
	return _ERC20TokenSource.Contract.TokenMultiplier(&_ERC20TokenSource.CallOpts)
}

// TokenMultiplier is a free data retrieval call binding the contract method 0xba3f5a12.
//
// Solidity: function tokenMultiplier() view returns(uint256)
func (_ERC20TokenSource *ERC20TokenSourceCallerSession) TokenMultiplier() (*big.Int, error) {
	return _ERC20TokenSource.Contract.TokenMultiplier(&_ERC20TokenSource.CallOpts)
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
