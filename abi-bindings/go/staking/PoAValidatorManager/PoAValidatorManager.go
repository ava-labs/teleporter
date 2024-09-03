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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"completeEndValidation\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structValidatorManagerSettings\",\"components\":[{\"name\":\"pChainBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeEndValidation\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeValidatorRegistration\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nodeID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendEndValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resendRegisterValidatorMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodCreated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"registerValidationMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodEnded\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumValidatorStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationPeriodRegistered\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRemovalInitialized\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161276438038061276483398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b6126178061014d5f395ff3fe608060405234801561000f575f80fd5b50600436106100a6575f3560e01c806397fb70d41161006e57806397fb70d414610144578063a3a65e4814610157578063b771b3bc1461016a578063bee0a03f14610178578063f2fde38b1461018b578063f3e148eb1461019e575f80fd5b80630322ed98146100aa5780633aaa9f25146100bf578063467ef06f146100e5578063715018a6146100f85780638da5cb5b14610100575b5f80fd5b6100bd6100b8366004611f1e565b6101b1565b005b6100d26100cd366004611fe2565b610387565b6040519081526020015b60405180910390f35b6100bd6100f3366004612085565b6103a7565b6100bd6107bf565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03165b6040516001600160a01b0390911681526020016100dc565b6100bd610152366004611f1e565b6107d2565b6100bd610165366004612085565b6107de565b61012c6005600160991b0181565b6100bd610186366004611f1e565b610a6c565b6100bd6101993660046120c3565b610b80565b6100bd6101ac3660046120de565b610bba565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb036020526040808220815160e0810190925280545f805160206125eb83398151915293929190829060ff1660058111156102115761021161211a565b60058111156102225761022261211a565b81526001820154602082015260028201546001600160a01b03811660408301526001600160401b03600160a01b909104811660608301526003928301548082166080840152600160401b8104821660a0840152600160801b90041660c090910152909150815160058111156102995761029961211a565b146103035760405162461bcd60e51b815260206004820152602f60248201527f56616c696461746f724d616e616765723a2076616c696461746f72206e6f742060448201526e1c195b991a5b99c81c995b5bdd985b608a1b60648201526084015b60405180910390fd5b5f6103138483606001515f610cc8565b60405163ee5b48eb60e01b81529091506005600160991b019063ee5b48eb90610340908490600401612150565b6020604051808303815f875af115801561035c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103809190612182565b5050505050565b5f610390610ee5565b61039c84868585610f40565b90505b949350505050565b5f5f805160206125eb8339815191526040516306f8253560e41b815263ffffffff841660048201529091505f9081906005600160991b0190636f825350906024015f60405180830381865afa158015610402573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261042991908101906121a8565b915091508061044a5760405162461bcd60e51b81526004016102fa9061227a565b825482511461046b5760405162461bcd60e51b81526004016102fa906122c0565b60208201516001600160a01b0316156104965760405162461bcd60e51b81526004016102fa90612309565b5f806104a5846040015161136e565b91509150801561050a5760405162461bcd60e51b815260206004820152602a60248201527f56616c696461746f724d616e616765723a20726567697374726174696f6e20736044820152691d1a5b1b081d985b1a5960b21b60648201526084016102fa565b5f828152600386016020526040808220815160e081019092528054829060ff16600581111561053b5761053b61211a565b600581111561054c5761054c61211a565b81526001820154602082015260028201546001600160a01b03811660408301526001600160401b03600160a01b909104811660608301526003928301548082166080840152600160401b8104821660a0840152600160801b90041660c0909101529091505f90825160058111156105c5576105c561211a565b14806105e357506001825160058111156105e1576105e161211a565b145b6106425760405162461bcd60e51b815260206004820152602a60248201527f56616c696461746f724d616e616765723a20696e76616c69642076616c696461604482015269746f722073746174757360b01b60648201526084016102fa565b6003825160058111156106575761065761211a565b0361066457506004610668565b5060055b6020808301515f908152600489019091526040812055818160058111156106915761069161211a565b908160058111156106a4576106a461211a565b9052505f84815260038801602052604090208251815484929190829060ff191660018360058111156106d8576106d861211a565b021790555060208201516001820155604082015160028201805460608501516001600160a01b039093166001600160e01b031990911617600160a01b6001600160401b039384160217905560808301516003909201805460a085015160c0909501519383166001600160801b031990911617600160401b948316949094029390931767ffffffffffffffff60801b1916600160801b92909116919091021790558151600581111561078b5761078b61211a565b60405185907f1c08e59656f1a18dc2da76826cdc52805c43e897a17c50faefb8ab3c1526cc16905f90a35050505050505050565b6107c7610ee5565b6107d05f6115cd565b565b6107db8161163d565b50565b5f5f805160206125eb8339815191526040516306f8253560e41b815263ffffffff841660048201529091505f9081906005600160991b0190636f825350906024015f60405180830381865afa158015610839573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261086091908101906121a8565b91509150806108815760405162461bcd60e51b81526004016102fa9061227a565b82548251146108a25760405162461bcd60e51b81526004016102fa906122c0565b60208201516001600160a01b0316156108cd5760405162461bcd60e51b81526004016102fa90612309565b5f806108dc846040015161136e565b915091508061093e5760405162461bcd60e51b815260206004820152602860248201527f56616c696461746f724d616e616765723a20726567697374726174696f6e206e6044820152671bdd081d985b1a5960c21b60648201526084016102fa565b5f8281526002860160205260408120805461095890612358565b905011801561098a575060015f83815260038701602052604090205460ff1660058111156109885761098861211a565b145b6109a65760405162461bcd60e51b81526004016102fa90612390565b5f82815260028601602052604081206109be91611ed4565b5f828152600386810160208181526040808520805460ff1916600217815593840180546001600160401b0342818116600160401b026fffffffffffffffff000000000000000019909316929092178355600190960154875260048c01845282872089905595889052928252915482519316835282019290925283917ff8fd1c90fb9cfa2ca2358fdf5806b086ad43315d92b221c929efc7f105ce7568910160405180910390a2505050505050565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb026020526040812080545f805160206125eb833981519152929190610ab490612358565b9050118015610ae6575060015f83815260038301602052604090205460ff166005811115610ae457610ae461211a565b145b610b025760405162461bcd60e51b81526004016102fa90612390565b5f82815260028201602052604090819020905163ee5b48eb60e01b81526005600160991b019163ee5b48eb91610b3b91906004016123d7565b6020604051808303815f875af1158015610b57573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b7b9190612182565b505050565b610b88610ee5565b6001600160a01b038116610bb157604051631e4fbdf760e01b81525f60048201526024016102fa565b6107db816115cd565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015610bfe5750825b90505f826001600160401b03166001148015610c195750303b155b905081158015610c27575080155b15610c455760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610c6f57845460ff60401b1916600160401b1785555b610c7987876119bd565b8315610cbf57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b60408051603680825260608281019093525f91906020820181803683370190505090505f5b6002811015610d4157610d01816001612475565b610d0c90600861248e565b5081515f90839083908110610d2357610d236124a5565b60200101906001600160f81b03191690815f1a905350600101610ced565b505f5b6004811015610da457610d58816003612475565b610d6390600861248e565b6001901c60f81b82610d768360026124b9565b81518110610d8657610d866124a5565b60200101906001600160f81b03191690815f1a905350600101610d44565b505f5b6020811015610e0657610dbb81601f612475565b610dc690600861248e565b86901c60f81b82610dd88360066124b9565b81518110610de857610de86124a5565b60200101906001600160f81b03191690815f1a905350600101610da7565b505f5b6008811015610e7157610e1d816007612475565b610e2890600861248e565b6001600160401b038616901c60f81b82610e438360266124b9565b81518110610e5357610e536124a5565b60200101906001600160f81b03191690815f1a905350600101610e09565b505f5b6008811015610edc57610e88816007612475565b610e9390600861248e565b6001600160401b038516901c60f81b82610eae83602e6124b9565b81518110610ebe57610ebe6124a5565b60200101906001600160f81b03191690815f1a905350600101610e74565b50949350505050565b33610f177f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146107d05760405163118cdaa760e01b81523360048201526024016102fa565b5f610f496119db565b5f805160206125eb833981519152426001600160401b038516118015610f8357506001600160401b038416610f81426202a3006124b9565b115b610fe55760405162461bcd60e51b815260206004820152602d60248201527f56616c696461746f724d616e616765723a20696e76616c69642072656769737460448201526c726174696f6e2065787069727960981b60648201526084016102fa565b8561103c5760405162461bcd60e51b815260206004820152602160248201527f56616c696461746f724d616e616765723a20696e76616c6964206e6f646520496044820152601160fa1b60648201526084016102fa565b5f868152600482016020526040902054156110aa5760405162461bcd60e51b815260206004820152602860248201527f56616c696461746f724d616e616765723a206e6f646520494420616c72656164604482015267792061637469766560c01b60648201526084016102fa565b82516030146111115760405162461bcd60e51b815260206004820152602d60248201527f56616c696461746f724d616e616765723a20696e76616c696420626c7350756260448201526c0d8d2c696caf240d8cadccee8d609b1b60648201526084016102fa565b5f806111586040518060a00160405280856001015481526020018a8152602001896001600160401b03168152602001886001600160401b0316815260200187815250611a25565b5f828152600286016020526040902091935091506111768282612510565b5060405163ee5b48eb60e01b81525f906005600160991b019063ee5b48eb906111a3908590600401612150565b6020604051808303815f875af11580156111bf573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111e39190612182565b6040805160e08101825260018152602081018c90529192508101336001600160a01b031681525f60208083018290526001600160401b038c166040808501919091526060840183905260809093018290528682526003880190522081518154829060ff1916600183600581111561125c5761125c61211a565b0217905550602082810151600183015560408084015160028401805460608701516001600160401b03908116600160a01b026001600160e01b03199092166001600160a01b039094169390931717905560808501516003909401805460a087015160c0909701518316600160801b0267ffffffffffffffff60801b19978416600160401b026001600160801b03199092169684169690961717959095169390931790935582518b83168152918a169082015282918b9186917f79b81620b81daf2c08cd5bb3dbb79e75d2d7a87f52171fde5aadc8c47823026e910160405180910390a450909250505061039f60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f8082516027146113d35760405162461bcd60e51b815260206004820152602960248201527f56616c696461746f724d657373616765733a20696e76616c6964206d657373616044820152680ceca40d8cadccee8d60bb1b60648201526084016102fa565b5f805b6002811015611422576113ea816001612475565b6113f590600861248e565b61ffff1685828151811061140b5761140b6124a5565b016020015160f81c901b91909117906001016113d6565b5061ffff8116156114815760405162461bcd60e51b815260206004820152602360248201527f56616c696461746f724d657373616765733a20696e76616c696420636f64656360448201526208125160ea1b60648201526084016102fa565b5f805b60048110156114dc57611498816003612475565b6114a390600861248e565b63ffffffff16866114b58360026124b9565b815181106114c5576114c56124a5565b016020015160f81c901b9190911790600101611484565b5063ffffffff81166002146115435760405162461bcd60e51b815260206004820152602760248201527f56616c696461746f724d657373616765733a20696e76616c6964206d657373616044820152666765207479706560c81b60648201526084016102fa565b5f805b60208110156115985761155a81601f612475565b61156590600861248e565b876115718360066124b9565b81518110611581576115816124a5565b016020015160f81c901b9190911790600101611546565b505f866026815181106115ad576115ad6124a5565b016020015191976001600160f81b03199092161515965090945050505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f8181527fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb036020526040808220815160e0810190925280545f805160206125eb83398151915293929190829060ff16600581111561169d5761169d61211a565b60058111156116ae576116ae61211a565b8152600182015460208201526002808301546001600160a01b03811660408401526001600160401b03600160a01b909104811660608401526003909301548084166080840152600160401b8104841660a0840152600160801b900490921660c090910152909150815160058111156117285761172861211a565b146117845760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f724d616e616765723a2076616c696461746f72206e6f742060448201526561637469766560d01b60648201526084016102fa565b60408101516001600160a01b0316336001600160a01b0316146117fe5760405162461bcd60e51b815260206004820152602c60248201527f56616c696461746f724d616e616765723a2073656e646572206e6f742076616c60448201526b34b230ba37b91037bbb732b960a11b60648201526084016102fa565b6003808252426001600160401b031660c08301525f848152908301602052604090208151815483929190829060ff191660018360058111156118425761184261211a565b02179055506020820151600182015560408201516002820180546060808601516001600160a01b039094166001600160e01b031990921691909117600160a01b6001600160401b03948516021790915560808401516003909301805460a086015160c0909601519484166001600160801b031990911617600160401b958416959095029490941767ffffffffffffffff60801b1916600160801b9390921692909202179091558101515f906118f990859083610cc8565b60405163ee5b48eb60e01b81529091505f906005600160991b019063ee5b48eb90611928908590600401612150565b6020604051808303815f875af1158015611944573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119689190612182565b6080840151604080516001600160401b039092168252426020830152919250829187917f13d58394cf269d48bcf927959a29a5ffee7c9924dafff8927ecdf3c48ffa7c67910160405180910390a35050505050565b6119c5611df0565b6119ce82611e39565b6119d781611e5a565b5050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00805460011901611a1f57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f6060826080015151603014611a8f5760405162461bcd60e51b815260206004820152602960248201527f5374616b696e674d657373616765733a20696e76616c6964207369676e6174756044820152680e4ca40d8cadccee8d60bb1b60648201526084016102fa565b60408051608680825260c082019092525f916020820181803683370190505090505f5b6002811015611b0657611ac6816001612475565b611ad190600861248e565b5081515f90839083908110611ae857611ae86124a5565b60200101906001600160f81b03191690815f1a905350600101611ab2565b505f5b6004811015611b6457611b1d816003612475565b611b2890600861248e565b505f82611b368360026124b9565b81518110611b4657611b466124a5565b60200101906001600160f81b03191690815f1a905350600101611b09565b505f5b6020811015611bc15784518160208110611b8357611b836124a5565b1a60f81b82611b938360066124b9565b81518110611ba357611ba36124a5565b60200101906001600160f81b03191690815f1a905350600101611b67565b505f5b6020811015611c215784602001518160208110611be357611be36124a5565b1a60f81b82611bf38360266124b9565b81518110611c0357611c036124a5565b60200101906001600160f81b03191690815f1a905350600101611bc4565b505f5b6008811015611c9557611c38816007612475565b611c4390600861248e565b60ff1685604001516001600160401b0316901c60f81b82826046611c6791906124b9565b81518110611c7757611c776124a5565b60200101906001600160f81b03191690815f1a905350600101611c24565b505f5b6030811015611d005784608001518181518110611cb757611cb76124a5565b01602001516001600160f81b03191682611cd283604e6124b9565b81518110611ce257611ce26124a5565b60200101906001600160f81b03191690815f1a905350600101611c98565b505f5b6008811015611d7257611d17816007612475565b611d2290600861248e565b60608601516001600160401b0390811691161c60f81b82611d4483607e6124b9565b81518110611d5457611d546124a5565b60200101906001600160f81b03191690815f1a905350600101611d03565b50600281604051611d8391906125cf565b602060405180830381855afa158015611d9e573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190611dc19190612182565b94909350915050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166107d057604051631afcd79f60e31b815260040160405180910390fd5b611e41611df0565b611e49611e6b565b611e51611e7b565b6107db81611e83565b611e62611df0565b6107db81611ec4565b611e73611df0565b6107d0611ecc565b6107d0611df0565b611e8b611df0565b80355f805160206125eb83398151915255602001357fe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb0155565b610b88611df0565b611dca611df0565b508054611ee090612358565b5f825580601f10611eef575050565b601f0160209004905f5260205f20908101906107db91905b80821115611f1a575f8155600101611f07565b5090565b5f60208284031215611f2e575f80fd5b5035919050565b80356001600160401b0381168114611f4b575f80fd5b919050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715611f8657611f86611f50565b60405290565b604051601f8201601f191681016001600160401b0381118282101715611fb457611fb4611f50565b604052919050565b5f6001600160401b03821115611fd457611fd4611f50565b50601f01601f191660200190565b5f805f8060808587031215611ff5575f80fd5b611ffe85611f35565b93506020850135925061201360408601611f35565b915060608501356001600160401b0381111561202d575f80fd5b8501601f8101871361203d575f80fd5b803561205061204b82611fbc565b611f8c565b818152886020838501011115612064575f80fd5b816020840160208301375f6020838301015280935050505092959194509250565b5f60208284031215612095575f80fd5b813563ffffffff811681146120a8575f80fd5b9392505050565b6001600160a01b03811681146107db575f80fd5b5f602082840312156120d3575f80fd5b81356120a8816120af565b5f8082840360608112156120f0575f80fd5b60408112156120fd575f80fd5b50829150604083013561210f816120af565b809150509250929050565b634e487b7160e01b5f52602160045260245ffd5b5f5b83811015612148578181015183820152602001612130565b50505f910152565b602081525f825180602084015261216e81604085016020870161212e565b601f01601f19169190910160400192915050565b5f60208284031215612192575f80fd5b5051919050565b80518015158114611f4b575f80fd5b5f80604083850312156121b9575f80fd5b82516001600160401b03808211156121cf575f80fd5b90840190606082870312156121e2575f80fd5b6121ea611f64565b825181526020808401516121fd816120af565b82820152604084015183811115612212575f80fd5b80850194505087601f850112612226575f80fd5b8351925061223661204b84611fbc565b8381528882858701011115612249575f80fd5b6122588483830184880161212e565b8060408401525081955061226d818801612199565b9450505050509250929050565b60208082526026908201527f56616c696461746f724d616e616765723a20696e76616c69642077617270206d60408201526565737361676560d01b606082015260800190565b60208082526029908201527f56616c696461746f724d616e616765723a20696e76616c696420736f757263656040820152680818da185a5b88125160ba1b606082015260800190565b6020808252602f908201527f56616c696461746f724d616e616765723a20696e76616c6964206f726967696e60408201526e2073656e646572206164647265737360881b606082015260800190565b600181811c9082168061236c57607f821691505b60208210810361238a57634e487b7160e01b5f52602260045260245ffd5b50919050565b60208082526027908201527f56616c696461746f724d616e616765723a20696e76616c69642076616c6964616040820152661d1a5bdb88125160ca1b606082015260800190565b5f60208083525f84546123e981612358565b806020870152604060018084165f811461240a576001811461242657612453565b60ff19851660408a0152604084151560051b8a01019550612453565b895f5260205f205f5b8581101561244a5781548b820186015290830190880161242f565b8a016040019650505b509398975050505050505050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561248857612488612461565b92915050565b808202811582820484141761248857612488612461565b634e487b7160e01b5f52603260045260245ffd5b8082018082111561248857612488612461565b601f821115610b7b57805f5260205f20601f840160051c810160208510156124f15750805b601f840160051c820191505b81811015610380575f81556001016124fd565b81516001600160401b0381111561252957612529611f50565b61253d816125378454612358565b846124cc565b602080601f831160018114612570575f84156125595750858301515b5f19600386901b1c1916600185901b1785556125c7565b5f85815260208120601f198616915b8281101561259e5788860151825594840194600190910190840161257f565b50858210156125bb57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f82516125e081846020870161212e565b919091019291505056fee92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb00a164736f6c6343000819000a",
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
