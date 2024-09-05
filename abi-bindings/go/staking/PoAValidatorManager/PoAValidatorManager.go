// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poavalidatormanager

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

// ValidatorManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type ValidatorManagerSettings struct {
	PChainBlockchainID [32]byte
	SubnetID           [32]byte
}

// PoAValidatorManagerMetaData contains all meta data concerning the PoAValidatorManager contract.
var PoAValidatorManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeValidators\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"completeEndValidation\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structValidatorManagerSettings\",\"components\":[{\"name\":\"pChainBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeValidatorRegistration\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nodeID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendEndValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendRegisterValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodEnded\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumValidatorStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodRegistered\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRemovalInitialized\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516127ad3803806127ad83398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b6126608061014d5f395ff3fe608060405234801561000f575f80fd5b50600436106100b1575f3560e01c806397fb70d41161006e57806397fb70d41461018d578063a3a65e48146101a0578063b771b3bc146101b3578063bee0a03f146101c1578063f2fde38b146101d4578063f3e148eb146101e7575f80fd5b80630322ed98146100b55780630cdd0985146100ca5780633aaa9f251461011b578063467ef06f1461012e578063715018a6146101415780638da5cb5b14610149575b5f80fd5b6100c86100c3366004611f67565b6101fa565b005b6101086100d8366004611f67565b5f9081527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb04602052604090205490565b6040519081526020015b60405180910390f35b61010861012936600461202b565b6103d0565b6100c861013c3660046120ce565b6103f0565b6100c8610808565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03165b6040516001600160a01b039091168152602001610112565b6100c861019b366004611f67565b61081b565b6100c86101ae3660046120ce565b610827565b6101756005600160991b0181565b6100c86101cf366004611f67565b610ab5565b6100c86101e236600461210c565b610bc9565b6100c86101f5366004612127565b610c03565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb036020526040808220815160e0810190925280545f8051602061263483398151915293929190829060ff16600581111561025a5761025a612163565b600581111561026b5761026b612163565b81526001820154602082015260028201546001600160a01b03811660408301526001600160401b03600160a01b909104811660608301526003928301548082166080840152600160401b8104821660a0840152600160801b90041660c090910152909150815160058111156102e2576102e2612163565b1461034c5760405162461bcd60e51b815260206004820152602f60248201527f56616c696461746f724d616e616765723a2076616c696461746f72206e6f742060448201526e1c195b991a5b99c81c995b5bdd985b608a1b60648201526084015b60405180910390fd5b5f61035c8483606001515f610d11565b60405163ee5b48eb60e01b81529091506005600160991b019063ee5b48eb90610389908490600401612199565b6020604051808303815f875af11580156103a5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103c991906121cb565b5050505050565b5f6103d9610f2e565b6103e584868585610f89565b90505b949350505050565b5f5f805160206126348339815191526040516306f8253560e41b815263ffffffff841660048201529091505f9081906005600160991b0190636f825350906024015f60405180830381865afa15801561044b573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261047291908101906121f1565b91509150806104935760405162461bcd60e51b8152600401610343906122c3565b82548251146104b45760405162461bcd60e51b815260040161034390612309565b60208201516001600160a01b0316156104df5760405162461bcd60e51b815260040161034390612352565b5f806104ee84604001516113b7565b9150915080156105535760405162461bcd60e51b815260206004820152602a60248201527f56616c696461746f724d616e616765723a20726567697374726174696f6e20736044820152691d1a5b1b081d985b1a5960b21b6064820152608401610343565b5f828152600386016020526040808220815160e081019092528054829060ff16600581111561058457610584612163565b600581111561059557610595612163565b81526001820154602082015260028201546001600160a01b03811660408301526001600160401b03600160a01b909104811660608301526003928301548082166080840152600160401b8104821660a0840152600160801b90041660c0909101529091505f908251600581111561060e5761060e612163565b148061062c575060018251600581111561062a5761062a612163565b145b61068b5760405162461bcd60e51b815260206004820152602a60248201527f56616c696461746f724d616e616765723a20696e76616c69642076616c696461604482015269746f722073746174757360b01b6064820152608401610343565b6003825160058111156106a0576106a0612163565b036106ad575060046106b1565b5060055b6020808301515f908152600489019091526040812055818160058111156106da576106da612163565b908160058111156106ed576106ed612163565b9052505f84815260038801602052604090208251815484929190829060ff1916600183600581111561072157610721612163565b021790555060208201516001820155604082015160028201805460608501516001600160a01b039093166001600160e01b031990911617600160a01b6001600160401b039384160217905560808301516003909201805460a085015160c0909501519383166001600160801b031990911617600160401b948316949094029390931767ffffffffffffffff60801b1916600160801b9290911691909102179055815160058111156107d4576107d4612163565b60405185907f1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16905f90a35050505050505050565b610810610f2e565b6108195f611616565b565b61082481611686565b50565b5f5f805160206126348339815191526040516306f8253560e41b815263ffffffff841660048201529091505f9081906005600160991b0190636f825350906024015f60405180830381865afa158015610882573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108a991908101906121f1565b91509150806108ca5760405162461bcd60e51b8152600401610343906122c3565b82548251146108eb5760405162461bcd60e51b815260040161034390612309565b60208201516001600160a01b0316156109165760405162461bcd60e51b815260040161034390612352565b5f8061092584604001516113b7565b91509150806109875760405162461bcd60e51b815260206004820152602860248201527f56616c696461746f724d616e616765723a20726567697374726174696f6e206e6044820152671bdd081d985b1a5960c21b6064820152608401610343565b5f828152600286016020526040812080546109a1906123a1565b90501180156109d3575060015f83815260038701602052604090205460ff1660058111156109d1576109d1612163565b145b6109ef5760405162461bcd60e51b8152600401610343906123d9565b5f8281526002860160205260408120610a0791611f1d565b5f828152600386810160208181526040808520805460ff1916600217815593840180546001600160401b0342818116600160401b026fffffffffffffffff000000000000000019909316929092178355600190960154875260048c01845282872089905595889052928252915482519316835282019290925283917ff8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568910160405180910390a2505050505050565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb026020526040812080545f80516020612634833981519152929190610afd906123a1565b9050118015610b2f575060015f83815260038301602052604090205460ff166005811115610b2d57610b2d612163565b145b610b4b5760405162461bcd60e51b8152600401610343906123d9565b5f82815260028201602052604090819020905163ee5b48eb60e01b81526005600160991b019163ee5b48eb91610b849190600401612420565b6020604051808303815f875af1158015610ba0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bc491906121cb565b505050565b610bd1610f2e565b6001600160a01b038116610bfa57604051631e4fbdf760e01b81525f6004820152602401610343565b61082481611616565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015610c475750825b90505f826001600160401b03166001148015610c625750303b155b905081158015610c70575080155b15610c8e5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610cb857845460ff60401b1916600160401b1785555b610cc28787611a06565b8315610d0857845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b60408051603680825260608281019093525f91906020820181803683370190505090505f5b6002811015610d8a57610d4a8160016124be565b610d559060086124d7565b5081515f90839083908110610d6c57610d6c6124ee565b60200101906001600160f81b03191690815f1a905350600101610d36565b505f5b6004811015610ded57610da18160036124be565b610dac9060086124d7565b6001901c60f81b82610dbf836002612502565b81518110610dcf57610dcf6124ee565b60200101906001600160f81b03191690815f1a905350600101610d8d565b505f5b6020811015610e4f57610e0481601f6124be565b610e0f9060086124d7565b86901c60f81b82610e21836006612502565b81518110610e3157610e316124ee565b60200101906001600160f81b03191690815f1a905350600101610df0565b505f5b6008811015610eba57610e668160076124be565b610e719060086124d7565b6001600160401b038616901c60f81b82610e8c836026612502565b81518110610e9c57610e9c6124ee565b60200101906001600160f81b03191690815f1a905350600101610e52565b505f5b6008811015610f2557610ed18160076124be565b610edc9060086124d7565b6001600160401b038516901c60f81b82610ef783602e612502565b81518110610f0757610f076124ee565b60200101906001600160f81b03191690815f1a905350600101610ebd565b50949350505050565b33610f607f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146108195760405163118cdaa760e01b8152336004820152602401610343565b5f610f92611a24565b5f80516020612634833981519152426001600160401b038516118015610fcc57506001600160401b038416610fca426202a300612502565b115b61102e5760405162461bcd60e51b815260206004820152602d60248201527f56616c696461746f724d616e616765723a20696e76616c69642072656769737460448201526c726174696f6e2065787069727960981b6064820152608401610343565b856110855760405162461bcd60e51b815260206004820152602160248201527f56616c696461746f724d616e616765723a20696e76616c6964206e6f646520496044820152601160fa1b6064820152608401610343565b5f868152600482016020526040902054156110f35760405162461bcd60e51b815260206004820152602860248201527f56616c696461746f724d616e616765723a206e6f646520494420616c72656164604482015267792061637469766560c01b6064820152608401610343565b825160301461115a5760405162461bcd60e51b815260206004820152602d60248201527f56616c696461746f724d616e616765723a20696e76616c696420626c7350756260448201526c0d8d2c696caf240d8cadccee8d609b1b6064820152608401610343565b5f806111a16040518060a00160405280856001015481526020018a8152602001896001600160401b03168152602001886001600160401b0316815260200187815250611a6e565b5f828152600286016020526040902091935091506111bf8282612559565b5060405163ee5b48eb60e01b81525f906005600160991b019063ee5b48eb906111ec908590600401612199565b6020604051808303815f875af1158015611208573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061122c91906121cb565b6040805160e08101825260018152602081018c90529192508101336001600160a01b031681525f60208083018290526001600160401b038c166040808501919091526060840183905260809093018290528682526003880190522081518154829060ff191660018360058111156112a5576112a5612163565b0217905550602082810151600183015560408084015160028401805460608701516001600160401b03908116600160a01b026001600160e01b03199092166001600160a01b039094169390931717905560808501516003909401805460a087015160c0909701518316600160801b0267ffffffffffffffff60801b19978416600160401b026001600160801b03199092169684169690961717959095169390931790935582518b83168152918a169082015282918b9186917f79b81620b81daf2c08cd5bb3dbb79e75d2d7a87f52171fde5aadc8c47823026e910160405180910390a45090925050506103e860017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f80825160271461141c5760405162461bcd60e51b815260206004820152602960248201527f56616c696461746f724d657373616765733a20696e76616c6964206d657373616044820152680ceca40d8cadccee8d60bb1b6064820152608401610343565b5f805b600281101561146b576114338160016124be565b61143e9060086124d7565b61ffff16858281518110611454576114546124ee565b016020015160f81c901b919091179060010161141f565b5061ffff8116156114ca5760405162461bcd60e51b815260206004820152602360248201527f56616c696461746f724d657373616765733a20696e76616c696420636f64656360448201526208125160ea1b6064820152608401610343565b5f805b6004811015611525576114e18160036124be565b6114ec9060086124d7565b63ffffffff16866114fe836002612502565b8151811061150e5761150e6124ee565b016020015160f81c901b91909117906001016114cd565b5063ffffffff811660021461158c5760405162461bcd60e51b815260206004820152602760248201527f56616c696461746f724d657373616765733a20696e76616c6964206d657373616044820152666765207479706560c81b6064820152608401610343565b5f805b60208110156115e1576115a381601f6124be565b6115ae9060086124d7565b876115ba836006612502565b815181106115ca576115ca6124ee565b016020015160f81c901b919091179060010161158f565b505f866026815181106115f6576115f66124ee565b016020015191976001600160f81b03199092161515965090945050505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb036020526040808220815160e0810190925280545f8051602061263483398151915293929190829060ff1660058111156116e6576116e6612163565b60058111156116f7576116f7612163565b8152600182015460208201526002808301546001600160a01b03811660408401526001600160401b03600160a01b909104811660608401526003909301548084166080840152600160401b8104841660a0840152600160801b900490921660c0909101529091508151600581111561177157611771612163565b146117cd5760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f724d616e616765723a2076616c696461746f72206e6f742060448201526561637469766560d01b6064820152608401610343565b60408101516001600160a01b0316336001600160a01b0316146118475760405162461bcd60e51b815260206004820152602c60248201527f56616c696461746f724d616e616765723a2073656e646572206e6f742076616c60448201526b34b230ba37b91037bbb732b960a11b6064820152608401610343565b6003808252426001600160401b031660c08301525f848152908301602052604090208151815483929190829060ff1916600183600581111561188b5761188b612163565b02179055506020820151600182015560408201516002820180546060808601516001600160a01b039094166001600160e01b031990921691909117600160a01b6001600160401b03948516021790915560808401516003909301805460a086015160c0909601519484166001600160801b031990911617600160401b958416959095029490941767ffffffffffffffff60801b1916600160801b9390921692909202179091558101515f9061194290859083610d11565b60405163ee5b48eb60e01b81529091505f906005600160991b019063ee5b48eb90611971908590600401612199565b6020604051808303815f875af115801561198d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119b191906121cb565b6080840151604080516001600160401b039092168252426020830152919250829187917f13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67910160405180910390a35050505050565b611a0e611e39565b611a1782611e82565b611a2081611ea3565b5050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00805460011901611a6857604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f6060826080015151603014611ad85760405162461bcd60e51b815260206004820152602960248201527f5374616b696e674d657373616765733a20696e76616c6964207369676e6174756044820152680e4ca40d8cadccee8d60bb1b6064820152608401610343565b60408051608680825260c082019092525f916020820181803683370190505090505f5b6002811015611b4f57611b0f8160016124be565b611b1a9060086124d7565b5081515f90839083908110611b3157611b316124ee565b60200101906001600160f81b03191690815f1a905350600101611afb565b505f5b6004811015611bad57611b668160036124be565b611b719060086124d7565b505f82611b7f836002612502565b81518110611b8f57611b8f6124ee565b60200101906001600160f81b03191690815f1a905350600101611b52565b505f5b6020811015611c0a5784518160208110611bcc57611bcc6124ee565b1a60f81b82611bdc836006612502565b81518110611bec57611bec6124ee565b60200101906001600160f81b03191690815f1a905350600101611bb0565b505f5b6020811015611c6a5784602001518160208110611c2c57611c2c6124ee565b1a60f81b82611c3c836026612502565b81518110611c4c57611c4c6124ee565b60200101906001600160f81b03191690815f1a905350600101611c0d565b505f5b6008811015611cde57611c818160076124be565b611c8c9060086124d7565b60ff1685604001516001600160401b0316901c60f81b82826046611cb09190612502565b81518110611cc057611cc06124ee565b60200101906001600160f81b03191690815f1a905350600101611c6d565b505f5b6030811015611d495784608001518181518110611d0057611d006124ee565b01602001516001600160f81b03191682611d1b83604e612502565b81518110611d2b57611d2b6124ee565b60200101906001600160f81b03191690815f1a905350600101611ce1565b505f5b6008811015611dbb57611d608160076124be565b611d6b9060086124d7565b60608601516001600160401b0390811691161c60f81b82611d8d83607e612502565b81518110611d9d57611d9d6124ee565b60200101906001600160f81b03191690815f1a905350600101611d4c565b50600281604051611dcc9190612618565b602060405180830381855afa158015611de7573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190611e0a91906121cb565b94909350915050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661081957604051631afcd79f60e31b815260040160405180910390fd5b611e8a611e39565b611e92611eb4565b611e9a611ec4565b61082481611ecc565b611eab611e39565b61082481611f0d565b611ebc611e39565b610819611f15565b610819611e39565b611ed4611e39565b80355f8051602061263483398151915255602001357fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0155565b610bd1611e39565b611e13611e39565b508054611f29906123a1565b5f825580601f10611f38575050565b601f0160209004905f5260205f209081019061082491905b80821115611f63575f8155600101611f50565b5090565b5f60208284031215611f77575f80fd5b5035919050565b80356001600160401b0381168114611f94575f80fd5b919050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715611fcf57611fcf611f99565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611ffd57611ffd611f99565b604052919050565b5f6001600160401b0382111561201d5761201d611f99565b50601f01601f191660200190565b5f805f806080858703121561203e575f80fd5b61204785611f7e565b93506020850135925061205c60408601611f7e565b915060608501356001600160401b03811115612076575f80fd5b8501601f81018713612086575f80fd5b803561209961209482612005565b611fd5565b8181528860208385010111156120ad575f80fd5b816020840160208301375f6020838301015280935050505092959194509250565b5f602082840312156120de575f80fd5b813563ffffffff811681146120f1575f80fd5b9392505050565b6001600160a01b0381168114610824575f80fd5b5f6020828403121561211c575f80fd5b81356120f1816120f8565b5f808284036060811215612139575f80fd5b6040811215612146575f80fd5b508291506040830135612158816120f8565b809150509250929050565b634e487b7160e01b5f52602160045260245ffd5b5f5b83811015612191578181015183820152602001612179565b50505f910152565b602081525f82518060208401526121b7816040850160208701612177565b601f01601f19169190910160400192915050565b5f602082840312156121db575f80fd5b5051919050565b80518015158114611f94575f80fd5b5f8060408385031215612202575f80fd5b82516001600160401b0380821115612218575f80fd5b908401906060828703121561222b575f80fd5b612233611fad565b82518152602080840151612246816120f8565b8282015260408401518381111561225b575f80fd5b80850194505087601f85011261226f575f80fd5b8351925061227f61209484612005565b8381528882858701011115612292575f80fd5b6122a184838301848801612177565b806040840152508195506122b68188016121e2565b9450505050509250929050565b60208082526026908201527f56616c696461746f724d616e616765723a20696e76616c69642077617270206d60408201526565737361676560d01b606082015260800190565b60208082526029908201527f56616c696461746f724d616e616765723a20696e76616c696420736f757263656040820152680818da185a5b88125160ba1b606082015260800190565b6020808252602f908201527f56616c696461746f724d616e616765723a20696e76616c6964206f726967696e60408201526e2073656e646572206164647265737360881b606082015260800190565b600181811c908216806123b557607f821691505b6020821081036123d357634e487b7160e01b5f52602260045260245ffd5b50919050565b60208082526027908201527f56616c696461746f724d616e616765723a20696e76616c69642076616c6964616040820152661d1a5bdb88125160ca1b606082015260800190565b5f60208083525f8454612432816123a1565b806020870152604060018084165f8114612453576001811461246f5761249c565b60ff19851660408a0152604084151560051b8a0101955061249c565b895f5260205f205f5b858110156124935781548b8201860152908301908801612478565b8a016040019650505b509398975050505050505050565b634e487b7160e01b5f52601160045260245ffd5b818103818111156124d1576124d16124aa565b92915050565b80820281158282048414176124d1576124d16124aa565b634e487b7160e01b5f52603260045260245ffd5b808201808211156124d1576124d16124aa565b601f821115610bc457805f5260205f20601f840160051c8101602085101561253a5750805b601f840160051c820191505b818110156103c9575f8155600101612546565b81516001600160401b0381111561257257612572611f99565b6125868161258084546123a1565b84612515565b602080601f8311600181146125b9575f84156125a25750858301515b5f19600386901b1c1916600185901b178555612610565b5f85815260208120601f198616915b828110156125e7578886015182559484019460019091019084016125c8565b508582101561260457878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f8251612629818460208701612177565b919091019291505056fee92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb00a164736f6c6343000819000a",
}

// PoAValidatorManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PoAValidatorManagerMetaData.ABI instead.
var PoAValidatorManagerABI = PoAValidatorManagerMetaData.ABI

// PoAValidatorManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoAValidatorManagerMetaData.Bin instead.
var PoAValidatorManagerBin = PoAValidatorManagerMetaData.Bin

// DeployPoAValidatorManager deploys a new Ethereum contract, binding an instance of PoAValidatorManager to it.
func DeployPoAValidatorManager(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *PoAValidatorManager, error) {
	parsed, err := PoAValidatorManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoAValidatorManagerBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoAValidatorManager{PoAValidatorManagerCaller: PoAValidatorManagerCaller{contract: contract}, PoAValidatorManagerTransactor: PoAValidatorManagerTransactor{contract: contract}, PoAValidatorManagerFilterer: PoAValidatorManagerFilterer{contract: contract}}, nil
}

// PoAValidatorManager is an auto generated Go binding around an Ethereum contract.
type PoAValidatorManager struct {
	PoAValidatorManagerCaller     // Read-only binding to the contract
	PoAValidatorManagerTransactor // Write-only binding to the contract
	PoAValidatorManagerFilterer   // Log filterer for contract events
}

// PoAValidatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoAValidatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoAValidatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoAValidatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoAValidatorManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoAValidatorManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoAValidatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoAValidatorManagerSession struct {
	Contract     *PoAValidatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PoAValidatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoAValidatorManagerCallerSession struct {
	Contract *PoAValidatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PoAValidatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoAValidatorManagerTransactorSession struct {
	Contract     *PoAValidatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PoAValidatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoAValidatorManagerRaw struct {
	Contract *PoAValidatorManager // Generic contract binding to access the raw methods on
}

// PoAValidatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoAValidatorManagerCallerRaw struct {
	Contract *PoAValidatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PoAValidatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoAValidatorManagerTransactorRaw struct {
	Contract *PoAValidatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoAValidatorManager creates a new instance of PoAValidatorManager, bound to a specific deployed contract.
func NewPoAValidatorManager(address common.Address, backend bind.ContractBackend) (*PoAValidatorManager, error) {
	contract, err := bindPoAValidatorManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManager{PoAValidatorManagerCaller: PoAValidatorManagerCaller{contract: contract}, PoAValidatorManagerTransactor: PoAValidatorManagerTransactor{contract: contract}, PoAValidatorManagerFilterer: PoAValidatorManagerFilterer{contract: contract}}, nil
}

// NewPoAValidatorManagerCaller creates a new read-only instance of PoAValidatorManager, bound to a specific deployed contract.
func NewPoAValidatorManagerCaller(address common.Address, caller bind.ContractCaller) (*PoAValidatorManagerCaller, error) {
	contract, err := bindPoAValidatorManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerCaller{contract: contract}, nil
}

// NewPoAValidatorManagerTransactor creates a new write-only instance of PoAValidatorManager, bound to a specific deployed contract.
func NewPoAValidatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PoAValidatorManagerTransactor, error) {
	contract, err := bindPoAValidatorManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerTransactor{contract: contract}, nil
}

// NewPoAValidatorManagerFilterer creates a new log filterer instance of PoAValidatorManager, bound to a specific deployed contract.
func NewPoAValidatorManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PoAValidatorManagerFilterer, error) {
	contract, err := bindPoAValidatorManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerFilterer{contract: contract}, nil
}

// bindPoAValidatorManager binds a generic wrapper to an already deployed contract.
func bindPoAValidatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoAValidatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoAValidatorManager *PoAValidatorManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoAValidatorManager.Contract.PoAValidatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoAValidatorManager *PoAValidatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.PoAValidatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoAValidatorManager *PoAValidatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.PoAValidatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoAValidatorManager *PoAValidatorManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoAValidatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoAValidatorManager *PoAValidatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoAValidatorManager *PoAValidatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.contract.Transact(opts, method, params...)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoAValidatorManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerSession) WARPMESSENGER() (common.Address, error) {
	return _PoAValidatorManager.Contract.WARPMESSENGER(&_PoAValidatorManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _PoAValidatorManager.Contract.WARPMESSENGER(&_PoAValidatorManager.CallOpts)
}

// ActiveValidators is a free data retrieval call binding the contract method 0x0cdd0985.
//
// Solidity: function activeValidators(bytes32 nodeID) view returns(bytes32)
func (_PoAValidatorManager *PoAValidatorManagerCaller) ActiveValidators(opts *bind.CallOpts, nodeID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PoAValidatorManager.contract.Call(opts, &out, "activeValidators", nodeID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActiveValidators is a free data retrieval call binding the contract method 0x0cdd0985.
//
// Solidity: function activeValidators(bytes32 nodeID) view returns(bytes32)
func (_PoAValidatorManager *PoAValidatorManagerSession) ActiveValidators(nodeID [32]byte) ([32]byte, error) {
	return _PoAValidatorManager.Contract.ActiveValidators(&_PoAValidatorManager.CallOpts, nodeID)
}

// ActiveValidators is a free data retrieval call binding the contract method 0x0cdd0985.
//
// Solidity: function activeValidators(bytes32 nodeID) view returns(bytes32)
func (_PoAValidatorManager *PoAValidatorManagerCallerSession) ActiveValidators(nodeID [32]byte) ([32]byte, error) {
	return _PoAValidatorManager.Contract.ActiveValidators(&_PoAValidatorManager.CallOpts, nodeID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoAValidatorManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerSession) Owner() (common.Address, error) {
	return _PoAValidatorManager.Contract.Owner(&_PoAValidatorManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoAValidatorManager *PoAValidatorManagerCallerSession) Owner() (common.Address, error) {
	return _PoAValidatorManager.Contract.Owner(&_PoAValidatorManager.CallOpts)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) CompleteEndValidation(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "completeEndValidation", messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.CompleteEndValidation(&_PoAValidatorManager.TransactOpts, messageIndex)
}

// CompleteEndValidation is a paid mutator transaction binding the contract method 0x467ef06f.
//
// Solidity: function completeEndValidation(uint32 messageIndex) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) CompleteEndValidation(messageIndex uint32) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.CompleteEndValidation(&_PoAValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.CompleteValidatorRegistration(&_PoAValidatorManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.CompleteValidatorRegistration(&_PoAValidatorManager.TransactOpts, messageIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0xf3e148eb.
//
// Solidity: function initialize((bytes32,bytes32) settings, address initialOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) Initialize(opts *bind.TransactOpts, settings ValidatorManagerSettings, initialOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "initialize", settings, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf3e148eb.
//
// Solidity: function initialize((bytes32,bytes32) settings, address initialOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) Initialize(settings ValidatorManagerSettings, initialOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.Initialize(&_PoAValidatorManager.TransactOpts, settings, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf3e148eb.
//
// Solidity: function initialize((bytes32,bytes32) settings, address initialOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) Initialize(settings ValidatorManagerSettings, initialOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.Initialize(&_PoAValidatorManager.TransactOpts, settings, initialOwner)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x97fb70d4.
//
// Solidity: function initializeEndValidation(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) InitializeEndValidation(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "initializeEndValidation", validationID)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x97fb70d4.
//
// Solidity: function initializeEndValidation(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) InitializeEndValidation(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.InitializeEndValidation(&_PoAValidatorManager.TransactOpts, validationID)
}

// InitializeEndValidation is a paid mutator transaction binding the contract method 0x97fb70d4.
//
// Solidity: function initializeEndValidation(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) InitializeEndValidation(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.InitializeEndValidation(&_PoAValidatorManager.TransactOpts, validationID)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x3aaa9f25.
//
// Solidity: function initializeValidatorRegistration(uint64 weight, bytes32 nodeID, uint64 registrationExpiry, bytes signature) returns(bytes32 validationID)
func (_PoAValidatorManager *PoAValidatorManagerTransactor) InitializeValidatorRegistration(opts *bind.TransactOpts, weight uint64, nodeID [32]byte, registrationExpiry uint64, signature []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "initializeValidatorRegistration", weight, nodeID, registrationExpiry, signature)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x3aaa9f25.
//
// Solidity: function initializeValidatorRegistration(uint64 weight, bytes32 nodeID, uint64 registrationExpiry, bytes signature) returns(bytes32 validationID)
func (_PoAValidatorManager *PoAValidatorManagerSession) InitializeValidatorRegistration(weight uint64, nodeID [32]byte, registrationExpiry uint64, signature []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.InitializeValidatorRegistration(&_PoAValidatorManager.TransactOpts, weight, nodeID, registrationExpiry, signature)
}

// InitializeValidatorRegistration is a paid mutator transaction binding the contract method 0x3aaa9f25.
//
// Solidity: function initializeValidatorRegistration(uint64 weight, bytes32 nodeID, uint64 registrationExpiry, bytes signature) returns(bytes32 validationID)
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) InitializeValidatorRegistration(weight uint64, nodeID [32]byte, registrationExpiry uint64, signature []byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.InitializeValidatorRegistration(&_PoAValidatorManager.TransactOpts, weight, nodeID, registrationExpiry, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.RenounceOwnership(&_PoAValidatorManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.RenounceOwnership(&_PoAValidatorManager.TransactOpts)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) ResendEndValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "resendEndValidatorMessage", validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.ResendEndValidatorMessage(&_PoAValidatorManager.TransactOpts, validationID)
}

// ResendEndValidatorMessage is a paid mutator transaction binding the contract method 0x0322ed98.
//
// Solidity: function resendEndValidatorMessage(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) ResendEndValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.ResendEndValidatorMessage(&_PoAValidatorManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) ResendRegisterValidatorMessage(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "resendRegisterValidatorMessage", validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.ResendRegisterValidatorMessage(&_PoAValidatorManager.TransactOpts, validationID)
}

// ResendRegisterValidatorMessage is a paid mutator transaction binding the contract method 0xbee0a03f.
//
// Solidity: function resendRegisterValidatorMessage(bytes32 validationID) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) ResendRegisterValidatorMessage(validationID [32]byte) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.ResendRegisterValidatorMessage(&_PoAValidatorManager.TransactOpts, validationID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.TransferOwnership(&_PoAValidatorManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoAValidatorManager *PoAValidatorManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoAValidatorManager.Contract.TransferOwnership(&_PoAValidatorManager.TransactOpts, newOwner)
}

// PoAValidatorManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PoAValidatorManager contract.
type PoAValidatorManagerInitializedIterator struct {
	Event *PoAValidatorManagerInitialized // Event containing the contract specifics and raw log

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
func (it *PoAValidatorManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerInitialized)
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
		it.Event = new(PoAValidatorManagerInitialized)
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
func (it *PoAValidatorManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerInitialized represents a Initialized event raised by the PoAValidatorManager contract.
type PoAValidatorManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoAValidatorManagerInitializedIterator, error) {

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerInitializedIterator{contract: _PoAValidatorManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerInitialized)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseInitialized(log types.Log) (*PoAValidatorManagerInitialized, error) {
	event := new(PoAValidatorManagerInitialized)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoAValidatorManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PoAValidatorManager contract.
type PoAValidatorManagerOwnershipTransferredIterator struct {
	Event *PoAValidatorManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PoAValidatorManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerOwnershipTransferred)
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
		it.Event = new(PoAValidatorManagerOwnershipTransferred)
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
func (it *PoAValidatorManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerOwnershipTransferred represents a OwnershipTransferred event raised by the PoAValidatorManager contract.
type PoAValidatorManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PoAValidatorManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerOwnershipTransferredIterator{contract: _PoAValidatorManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerOwnershipTransferred)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseOwnershipTransferred(log types.Log) (*PoAValidatorManagerOwnershipTransferred, error) {
	event := new(PoAValidatorManagerOwnershipTransferred)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoAValidatorManagerValidationPeriodCreatedIterator is returned from FilterValidationPeriodCreated and is used to iterate over the raw logs and unpacked data for ValidationPeriodCreated events raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidationPeriodCreatedIterator struct {
	Event *PoAValidatorManagerValidationPeriodCreated // Event containing the contract specifics and raw log

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
func (it *PoAValidatorManagerValidationPeriodCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerValidationPeriodCreated)
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
		it.Event = new(PoAValidatorManagerValidationPeriodCreated)
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
func (it *PoAValidatorManagerValidationPeriodCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerValidationPeriodCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerValidationPeriodCreated represents a ValidationPeriodCreated event raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidationPeriodCreated struct {
	ValidationID                [32]byte
	NodeID                      [32]byte
	RegisterValidationMessageID [32]byte
	Weight                      *big.Int
	RegistrationExpiry          uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodCreated is a free log retrieval operation binding the contract event 0x79b81620b81daf2c08cd5bb3dbb79e75d2d7a87f52171fde5aadc8c47823026e.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes32 indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterValidationPeriodCreated(opts *bind.FilterOpts, validationID [][32]byte, nodeID [][32]byte, registerValidationMessageID [][32]byte) (*PoAValidatorManagerValidationPeriodCreatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerValidationPeriodCreatedIterator{contract: _PoAValidatorManager.contract, event: "ValidationPeriodCreated", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodCreated is a free log subscription operation binding the contract event 0x79b81620b81daf2c08cd5bb3dbb79e75d2d7a87f52171fde5aadc8c47823026e.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes32 indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchValidationPeriodCreated(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerValidationPeriodCreated, validationID [][32]byte, nodeID [][32]byte, registerValidationMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var nodeIDRule []interface{}
	for _, nodeIDItem := range nodeID {
		nodeIDRule = append(nodeIDRule, nodeIDItem)
	}
	var registerValidationMessageIDRule []interface{}
	for _, registerValidationMessageIDItem := range registerValidationMessageID {
		registerValidationMessageIDRule = append(registerValidationMessageIDRule, registerValidationMessageIDItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "ValidationPeriodCreated", validationIDRule, nodeIDRule, registerValidationMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerValidationPeriodCreated)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
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

// ParseValidationPeriodCreated is a log parse operation binding the contract event 0x79b81620b81daf2c08cd5bb3dbb79e75d2d7a87f52171fde5aadc8c47823026e.
//
// Solidity: event ValidationPeriodCreated(bytes32 indexed validationID, bytes32 indexed nodeID, bytes32 indexed registerValidationMessageID, uint256 weight, uint64 registrationExpiry)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseValidationPeriodCreated(log types.Log) (*PoAValidatorManagerValidationPeriodCreated, error) {
	event := new(PoAValidatorManagerValidationPeriodCreated)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidationPeriodCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoAValidatorManagerValidationPeriodEndedIterator is returned from FilterValidationPeriodEnded and is used to iterate over the raw logs and unpacked data for ValidationPeriodEnded events raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidationPeriodEndedIterator struct {
	Event *PoAValidatorManagerValidationPeriodEnded // Event containing the contract specifics and raw log

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
func (it *PoAValidatorManagerValidationPeriodEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerValidationPeriodEnded)
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
		it.Event = new(PoAValidatorManagerValidationPeriodEnded)
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
func (it *PoAValidatorManagerValidationPeriodEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerValidationPeriodEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerValidationPeriodEnded represents a ValidationPeriodEnded event raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidationPeriodEnded struct {
	ValidationID [32]byte
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodEnded is a free log retrieval operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterValidationPeriodEnded(opts *bind.FilterOpts, validationID [][32]byte, status []uint8) (*PoAValidatorManagerValidationPeriodEndedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerValidationPeriodEndedIterator{contract: _PoAValidatorManager.contract, event: "ValidationPeriodEnded", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodEnded is a free log subscription operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchValidationPeriodEnded(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerValidationPeriodEnded, validationID [][32]byte, status []uint8) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "ValidationPeriodEnded", validationIDRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerValidationPeriodEnded)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
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

// ParseValidationPeriodEnded is a log parse operation binding the contract event 0x1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16.
//
// Solidity: event ValidationPeriodEnded(bytes32 indexed validationID, uint8 indexed status)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseValidationPeriodEnded(log types.Log) (*PoAValidatorManagerValidationPeriodEnded, error) {
	event := new(PoAValidatorManagerValidationPeriodEnded)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidationPeriodEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoAValidatorManagerValidationPeriodRegisteredIterator is returned from FilterValidationPeriodRegistered and is used to iterate over the raw logs and unpacked data for ValidationPeriodRegistered events raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidationPeriodRegisteredIterator struct {
	Event *PoAValidatorManagerValidationPeriodRegistered // Event containing the contract specifics and raw log

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
func (it *PoAValidatorManagerValidationPeriodRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerValidationPeriodRegistered)
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
		it.Event = new(PoAValidatorManagerValidationPeriodRegistered)
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
func (it *PoAValidatorManagerValidationPeriodRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerValidationPeriodRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerValidationPeriodRegistered represents a ValidationPeriodRegistered event raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidationPeriodRegistered struct {
	ValidationID [32]byte
	Weight       *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidationPeriodRegistered is a free log retrieval operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterValidationPeriodRegistered(opts *bind.FilterOpts, validationID [][32]byte) (*PoAValidatorManagerValidationPeriodRegisteredIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerValidationPeriodRegisteredIterator{contract: _PoAValidatorManager.contract, event: "ValidationPeriodRegistered", logs: logs, sub: sub}, nil
}

// WatchValidationPeriodRegistered is a free log subscription operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchValidationPeriodRegistered(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerValidationPeriodRegistered, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "ValidationPeriodRegistered", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerValidationPeriodRegistered)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
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

// ParseValidationPeriodRegistered is a log parse operation binding the contract event 0xf8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568.
//
// Solidity: event ValidationPeriodRegistered(bytes32 indexed validationID, uint256 weight, uint256 timestamp)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseValidationPeriodRegistered(log types.Log) (*PoAValidatorManagerValidationPeriodRegistered, error) {
	event := new(PoAValidatorManagerValidationPeriodRegistered)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidationPeriodRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoAValidatorManagerValidatorRemovalInitializedIterator is returned from FilterValidatorRemovalInitialized and is used to iterate over the raw logs and unpacked data for ValidatorRemovalInitialized events raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidatorRemovalInitializedIterator struct {
	Event *PoAValidatorManagerValidatorRemovalInitialized // Event containing the contract specifics and raw log

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
func (it *PoAValidatorManagerValidatorRemovalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoAValidatorManagerValidatorRemovalInitialized)
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
		it.Event = new(PoAValidatorManagerValidatorRemovalInitialized)
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
func (it *PoAValidatorManagerValidatorRemovalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoAValidatorManagerValidatorRemovalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoAValidatorManagerValidatorRemovalInitialized represents a ValidatorRemovalInitialized event raised by the PoAValidatorManager contract.
type PoAValidatorManagerValidatorRemovalInitialized struct {
	ValidationID       [32]byte
	SetWeightMessageID [32]byte
	Weight             *big.Int
	EndTime            *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemovalInitialized is a free log retrieval operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) FilterValidatorRemovalInitialized(opts *bind.FilterOpts, validationID [][32]byte, setWeightMessageID [][32]byte) (*PoAValidatorManagerValidatorRemovalInitializedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.FilterLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return &PoAValidatorManagerValidatorRemovalInitializedIterator{contract: _PoAValidatorManager.contract, event: "ValidatorRemovalInitialized", logs: logs, sub: sub}, nil
}

// WatchValidatorRemovalInitialized is a free log subscription operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) WatchValidatorRemovalInitialized(opts *bind.WatchOpts, sink chan<- *PoAValidatorManagerValidatorRemovalInitialized, validationID [][32]byte, setWeightMessageID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var setWeightMessageIDRule []interface{}
	for _, setWeightMessageIDItem := range setWeightMessageID {
		setWeightMessageIDRule = append(setWeightMessageIDRule, setWeightMessageIDItem)
	}

	logs, sub, err := _PoAValidatorManager.contract.WatchLogs(opts, "ValidatorRemovalInitialized", validationIDRule, setWeightMessageIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoAValidatorManagerValidatorRemovalInitialized)
				if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
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

// ParseValidatorRemovalInitialized is a log parse operation binding the contract event 0x13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67.
//
// Solidity: event ValidatorRemovalInitialized(bytes32 indexed validationID, bytes32 indexed setWeightMessageID, uint256 weight, uint256 endTime)
func (_PoAValidatorManager *PoAValidatorManagerFilterer) ParseValidatorRemovalInitialized(log types.Log) (*PoAValidatorManagerValidatorRemovalInitialized, error) {
	event := new(PoAValidatorManagerValidatorRemovalInitialized)
	if err := _PoAValidatorManager.contract.UnpackLog(event, "ValidatorRemovalInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
