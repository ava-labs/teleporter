// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package codec

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

// InitialValidator is an auto generated low-level Go binding around an user-defined struct.
type InitialValidator struct {
	NodeID       []byte
	BlsPublicKey []byte
	Weight       uint64
}

// PChainOwner is an auto generated low-level Go binding around an user-defined struct.
type PChainOwner struct {
	Threshold uint32
	Addresses []common.Address
}

// SubnetConversionData is an auto generated low-level Go binding around an user-defined struct.
type SubnetConversionData struct {
	SubnetID                     [32]byte
	ValidatorManagerBlockchainID [32]byte
	ValidatorManagerAddress      common.Address
	InitialValidators            []InitialValidator
}

// ValidatorMessagesValidationPeriod is an auto generated low-level Go binding around an user-defined struct.
type ValidatorMessagesValidationPeriod struct {
	SubnetID              [32]byte
	NodeID                []byte
	BlsPublicKey          []byte
	RegistrationExpiry    uint64
	RemainingBalanceOwner PChainOwner
	DisableOwner          PChainOwner
	Weight                uint64
}

// CodecMetaData contains all meta data concerning the Codec contract.
var CodecMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"packRegisterSubnetValidatorMessage\",\"inputs\":[{\"name\":\"validationPeriod\",\"type\":\"tuple\",\"internalType\":\"structValidatorMessages.ValidationPeriod\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"disableOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"packSubnetConversionData\",\"inputs\":[{\"name\":\"subnetConversionData\",\"type\":\"tuple\",\"internalType\":\"structSubnetConversionData\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialValidators\",\"type\":\"tuple[]\",\"internalType\":\"structInitialValidator[]\",\"components\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"packSubnetConversionMessage\",\"inputs\":[{\"name\":\"subnetConversionID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"packSubnetValidatorRegistrationMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"packSubnetValidatorWeightMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"packValidationUptimeMessage\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"uptime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"unpackRegisterSubnetValidatorMessage\",\"inputs\":[{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structValidatorMessages.ValidationPeriod\",\"components\":[{\"name\":\"subnetID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"registrationExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"disableOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"unpackSubnetConversionMessage\",\"inputs\":[{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"unpackSubnetValidatorRegistrationMessage\",\"inputs\":[{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"unpackSubnetValidatorWeightMessage\",\"inputs\":[{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"unpackValidationUptimeMessage\",\"inputs\":[{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"InvalidBLSPublicKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCodecID\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageType\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f80fd5b506121c78061001c5f395ff3fe608060405234801561000f575f80fd5b50600436106100a6575f3560e01c80638545c16a1161006e5780638545c16a1461016d5780639de23d4014610180578063a5233770146101b8578063e1d68f30146101d8578063f65e4b33146101eb578063fa1f8dfb146101fe575f80fd5b806301bbec74146100aa578063088c2463146100d45780631e6d9789146101045780631fd979c7146101255780632e43ceb514610145575b5f80fd5b6100bd6100b8366004611aa2565b610211565b6040516100cb929190611beb565b60405180910390f35b6100e76100e2366004611c03565b610226565b604080519283526001600160401b039091166020830152016100cb565b610117610112366004611c03565b610231565b6040519081526020016100cb565b610138610133366004611c34565b610241565b6040516100cb9190611c4b565b610158610153366004611c03565b610273565b604080519283529015156020830152016100cb565b61013861017b366004611c5d565b61027e565b61019361018e366004611c03565b6102d5565b604080519384526001600160401b0392831660208501529116908201526060016100cb565b6101cb6101c6366004611c03565b6102ee565b6040516100cb9190611cfc565b6101386101e6366004611db3565b6102ff565b6101386101f9366004611ddd565b610312565b61013861020c366004611e13565b61031d565b5f606061021d83610329565b91509150915091565b5f8061021d83610516565b5f61023b8261070c565b92915050565b606061023b82604080515f60208201819052602282015260268082019390935281518082039093018352604601905290565b5f8061021d83610899565b604080515f6020820152600360e01b6022820152602681018590526001600160c01b031960c085811b8216604684015284901b16604e8201528151808203603601815260569091019091526060905b949350505050565b5f805f6102e184610a55565b9250925092509193909250565b6102f6611850565b61023b82610cab565b606061030b83836115f6565b9392505050565b606061023b8261163f565b606061030b8383611820565b5f60608260400151516030146103525760405163180ffa0d60e01b815260040160405180910390fd5b82516020808501518051604080880151606089015160808a01518051908701515193515f98610393988a986001989297929690959094909390929101611e45565b60405160208183030381529060405290505f5b84608001516020015151811015610405578185608001516020015182815181106103d2576103d2611eff565b60200260200101516040516020016103eb929190611f13565b60408051601f1981840301815291905291506001016103a6565b5060a0840151805160209182015151604051610425938593929101611f49565b60405160208183030381529060405290505f5b8460a00151602001515181101561049757818560a0015160200151828151811061046457610464611eff565b602002602001015160405160200161047d929190611f13565b60408051601f198184030181529190529150600101610438565b5060c08401516040516104ae918391602001611f84565b60405160208183030381529060405290506002816040516104cf9190611fb5565b602060405180830381855afa1580156104ea573d5f803e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061050d9190611fd0565b94909350915050565b5f808251602e1461055157825160405163cc92daa160e01b815263ffffffff9091166004820152602e60248201526044015b60405180910390fd5b5f805b60028110156105a057610568816001611ffb565b61057390600861200e565b61ffff1685828151811061058957610589611eff565b016020015160f81c901b9190911790600101610554565b5061ffff8116156105ca5760405163407b587360e01b815261ffff82166004820152602401610548565b5f805b6004811015610625576105e1816003611ffb565b6105ec90600861200e565b63ffffffff16866105fe836002612025565b8151811061060e5761060e611eff565b016020015160f81c901b91909117906001016105cd565b5063ffffffff81161561064b57604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156106a05761066281601f611ffb565b61066d90600861200e565b87610679836006612025565b8151811061068957610689611eff565b016020015160f81c901b919091179060010161064e565b505f805b60088110156106ff576106b8816007611ffb565b6106c390600861200e565b6001600160401b0316886106d8836026612025565b815181106106e8576106e8611eff565b016020015160f81c901b91909117906001016106a4565b5090969095509350505050565b5f815160261461074157815160405163cc92daa160e01b815263ffffffff909116600482015260266024820152604401610548565b5f805b600281101561079057610758816001611ffb565b61076390600861200e565b61ffff1684828151811061077957610779611eff565b016020015160f81c901b9190911790600101610744565b5061ffff8116156107ba5760405163407b587360e01b815261ffff82166004820152602401610548565b5f805b6004811015610815576107d1816003611ffb565b6107dc90600861200e565b63ffffffff16856107ee836002612025565b815181106107fe576107fe611eff565b016020015160f81c901b91909117906001016107bd565b5063ffffffff81161561083b57604051635b60892f60e01b815260040160405180910390fd5b5f805b60208110156108905761085281601f611ffb565b61085d90600861200e565b86610869836006612025565b8151811061087957610879611eff565b016020015160f81c901b919091179060010161083e565b50949350505050565b5f8082516027146108cf57825160405163cc92daa160e01b815263ffffffff909116600482015260276024820152604401610548565b5f805b600281101561091e576108e6816001611ffb565b6108f190600861200e565b61ffff1685828151811061090757610907611eff565b016020015160f81c901b91909117906001016108d2565b5061ffff8116156109485760405163407b587360e01b815261ffff82166004820152602401610548565b5f805b60048110156109a35761095f816003611ffb565b61096a90600861200e565b63ffffffff168661097c836002612025565b8151811061098c5761098c611eff565b016020015160f81c901b919091179060010161094b565b5063ffffffff81166002146109cb57604051635b60892f60e01b815260040160405180910390fd5b5f805b6020811015610a20576109e281601f611ffb565b6109ed90600861200e565b876109f9836006612025565b81518110610a0957610a09611eff565b016020015160f81c901b91909117906001016109ce565b505f86602681518110610a3557610a35611eff565b016020015191976001600160f81b03199092161515965090945050505050565b5f805f8351603614610a8c57835160405163cc92daa160e01b815263ffffffff909116600482015260366024820152604401610548565b5f805b6002811015610adb57610aa3816001611ffb565b610aae90600861200e565b61ffff16868281518110610ac457610ac4611eff565b016020015160f81c901b9190911790600101610a8f565b5061ffff811615610b055760405163407b587360e01b815261ffff82166004820152602401610548565b5f805b6004811015610b6057610b1c816003611ffb565b610b2790600861200e565b63ffffffff1687610b39836002612025565b81518110610b4957610b49611eff565b016020015160f81c901b9190911790600101610b08565b5063ffffffff8116600314610b8857604051635b60892f60e01b815260040160405180910390fd5b5f805b6020811015610bdd57610b9f81601f611ffb565b610baa90600861200e565b88610bb6836006612025565b81518110610bc657610bc6611eff565b016020015160f81c901b9190911790600101610b8b565b505f805b6008811015610c3c57610bf5816007611ffb565b610c0090600861200e565b6001600160401b031689610c15836026612025565b81518110610c2557610c25611eff565b016020015160f81c901b9190911790600101610be1565b505f805b6008811015610c9b57610c54816007611ffb565b610c5f90600861200e565b6001600160401b03168a610c7483602e612025565b81518110610c8457610c84611eff565b016020015160f81c901b9190911790600101610c40565b5091989097509095509350505050565b610cb3611850565b5f610cbc611850565b5f805b6002811015610d1a57610cd3816001611ffb565b610cde90600861200e565b61ffff1686610cf363ffffffff871684612025565b81518110610d0357610d03611eff565b016020015160f81c901b9190911790600101610cbf565b5061ffff811615610d445760405163407b587360e01b815261ffff82166004820152602401610548565b610d4f600284612038565b9250505f805b6004811015610db457610d69816003611ffb565b610d7490600861200e565b63ffffffff16868563ffffffff1683610d8d9190612025565b81518110610d9d57610d9d611eff565b016020015160f81c901b9190911790600101610d55565b5063ffffffff8116600114610ddc57604051635b60892f60e01b815260040160405180910390fd5b610de7600484612038565b9250505f805b6020811015610e4457610e0181601f611ffb565b610e0c90600861200e565b86610e1d63ffffffff871684612025565b81518110610e2d57610e2d611eff565b016020015160f81c901b9190911790600101610ded565b50808252610e53602084612038565b9250505f805b6004811015610eb857610e6d816003611ffb565b610e7890600861200e565b63ffffffff16868563ffffffff1683610e919190612025565b81518110610ea157610ea1611eff565b016020015160f81c901b9190911790600101610e59565b50610ec4600484612038565b92505f8163ffffffff166001600160401b03811115610ee557610ee56118aa565b6040519080825280601f01601f191660200182016040528015610f0f576020820181803683370190505b5090505f5b8263ffffffff16811015610f7e5786610f3363ffffffff871683612025565b81518110610f4357610f43611eff565b602001015160f81c60f81b828281518110610f6057610f60611eff565b60200101906001600160f81b03191690815f1a905350600101610f14565b5060208301819052610f908285612038565b604080516030808252606082019092529195505f92506020820181803683370190505090505f5b603081101561101c5786610fd163ffffffff871683612025565b81518110610fe157610fe1611eff565b602001015160f81c60f81b828281518110610ffe57610ffe611eff565b60200101906001600160f81b03191690815f1a905350600101610fb7565b506040830181905261102f603085612038565b9350505f805b600881101561109557611049816007611ffb565b61105490600861200e565b6001600160401b03168761106e63ffffffff881684612025565b8151811061107e5761107e611eff565b016020015160f81c901b9190911790600101611035565b506001600160401b03811660608401526110b0600885612038565b9350505f805f5b6004811015611116576110cb816003611ffb565b6110d690600861200e565b63ffffffff16888763ffffffff16836110ef9190612025565b815181106110ff576110ff611eff565b016020015160f81c901b91909117906001016110b7565b50611122600486612038565b94505f5b60048110156111855761113a816003611ffb565b61114590600861200e565b63ffffffff16888763ffffffff168361115e9190612025565b8151811061116e5761116e611eff565b016020015160f81c901b9290921791600101611126565b50611191600486612038565b94505f8263ffffffff166001600160401b038111156111b2576111b26118aa565b6040519080825280602002602001820160405280156111db578160200160208202803683370190505b5090505f5b8363ffffffff168110156112c3576040805160148082528183019092525f916020820181803683370190505090505f5b6014811015611275578a61122a63ffffffff8b1683612025565b8151811061123a5761123a611eff565b602001015160f81c60f81b82828151811061125757611257611eff565b60200101906001600160f81b03191690815f1a905350600101611210565b505f601482015190508084848151811061129157611291611eff565b6001600160a01b03909216602092830291909101909101526112b460148a612038565b985050508060010190506111e0565b506040805180820190915263ffffffff9092168252602082015260808401525f80805b6004811015611345576112fa816003611ffb565b61130590600861200e565b63ffffffff16898863ffffffff168361131e9190612025565b8151811061132e5761132e611eff565b016020015160f81c901b91909117906001016112e6565b50611351600487612038565b95505f5b60048110156113b457611369816003611ffb565b61137490600861200e565b63ffffffff16898863ffffffff168361138d9190612025565b8151811061139d5761139d611eff565b016020015160f81c901b9290921791600101611355565b506113c0600487612038565b95505f8263ffffffff166001600160401b038111156113e1576113e16118aa565b60405190808252806020026020018201604052801561140a578160200160208202803683370190505b5090505f5b8363ffffffff168110156114f2576040805160148082528183019092525f916020820181803683370190505090505f5b60148110156114a4578b61145963ffffffff8c1683612025565b8151811061146957611469611eff565b602001015160f81c60f81b82828151811061148657611486611eff565b60200101906001600160f81b03191690815f1a90535060010161143f565b505f60148201519050808484815181106114c0576114c0611eff565b6001600160a01b03909216602092830291909101909101526114e360148b612038565b9950505080600101905061140f565b506040805180820190915263ffffffff9092168252602082015260a08501525f61151c8284612038565b611527906014612055565b61153285607a612038565b61153c9190612038565b90508063ffffffff1688511461157857875160405163cc92daa160e01b815263ffffffff91821660048201529082166024820152604401610548565b5f805b60088110156115db5761158f816007611ffb565b61159a90600861200e565b6001600160401b03168a6115b463ffffffff8b1684612025565b815181106115c4576115c4611eff565b016020015160f81c901b919091179060010161157b565b506001600160401b031660c086015250929695505050505050565b6040515f602082018190526022820152602681018390526001600160c01b031960c083901b166046820152606090604e015b604051602081830303815290604052905092915050565b60605f8083356020850135601461165b8787016040890161207d565b6116686060890189612096565b60405160f09790971b6001600160f01b0319166020880152602287019590955250604285019290925260e090811b6001600160e01b0319908116606286015260609290921b6bffffffffffffffffffffffff191660668501529190911b16607a820152607e0160405160208183030381529060405290505f5b6116ee6060850185612096565b905081101561181957816117056060860186612096565b8381811061171557611715611eff565b905060200281019061172791906120e2565b61173190806120f6565b90506117406060870187612096565b8481811061175057611750611eff565b905060200281019061176291906120e2565b61176c90806120f6565b6117796060890189612096565b8681811061178957611789611eff565b905060200281019061179b91906120e2565b6117a99060208101906120f6565b6117b660608b018b612096565b888181106117c6576117c6611eff565b90506020028101906117d891906120e2565b6117e9906060810190604001612138565b6040516020016117ff9796959493929190612151565b60408051601f1981840301815291905291506001016116e1565b5092915050565b6040515f6020820152600160e11b60228201526026810183905281151560f81b6046820152606090604701611628565b6040805160e0810182525f808252606060208084018290528385018290528184018390528451808601865283815280820183905260808501528451808601909552918452908301529060a082019081525f60209091015290565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156118e0576118e06118aa565b60405290565b60405160e081016001600160401b03811182821017156118e0576118e06118aa565b604051601f8201601f191681016001600160401b0381118282101715611930576119306118aa565b604052919050565b5f82601f830112611947575f80fd5b81356001600160401b03811115611960576119606118aa565b611973601f8201601f1916602001611908565b818152846020838601011115611987575f80fd5b816020850160208301375f918101602001919091529392505050565b80356001600160401b03811681146119b9575f80fd5b919050565b80356001600160a01b03811681146119b9575f80fd5b5f604082840312156119e4575f80fd5b6119ec6118be565b9050813563ffffffff81168114611a01575f80fd5b81526020828101356001600160401b0380821115611a1d575f80fd5b818501915085601f830112611a30575f80fd5b813581811115611a4257611a426118aa565b8060051b9150611a53848301611908565b8181529183018401918481019088841115611a6c575f80fd5b938501935b83851015611a9157611a82856119be565b82529385019390850190611a71565b808688015250505050505092915050565b5f60208284031215611ab2575f80fd5b81356001600160401b0380821115611ac8575f80fd5b9083019060e08286031215611adb575f80fd5b611ae36118e6565b82358152602083013582811115611af8575f80fd5b611b0487828601611938565b602083015250604083013582811115611b1b575f80fd5b611b2787828601611938565b604083015250611b39606084016119a3565b6060820152608083013582811115611b4f575f80fd5b611b5b878286016119d4565b60808301525060a083013582811115611b72575f80fd5b611b7e878286016119d4565b60a083015250611b9060c084016119a3565b60c082015295945050505050565b5f5b83811015611bb8578181015183820152602001611ba0565b50505f910152565b5f8151808452611bd7816020860160208601611b9e565b601f01601f19169290920160200192915050565b828152604060208201525f6102cd6040830184611bc0565b5f60208284031215611c13575f80fd5b81356001600160401b03811115611c28575f80fd5b6102cd84828501611938565b5f60208284031215611c44575f80fd5b5035919050565b602081525f61030b6020830184611bc0565b5f805f60608486031215611c6f575f80fd5b83359250611c7f602085016119a3565b9150611c8d604085016119a3565b90509250925092565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015611cf15784516001600160a01b03168252938301936001929092019190830190611cc8565b509695505050505050565b60208152815160208201525f602083015160e06040840152611d22610100840182611bc0565b90506040840151601f1980858403016060860152611d408383611bc0565b92506001600160401b03606087015116608086015260808601519150808584030160a0860152611d708383611c96565b925060a08601519150808584030160c086015250611d8e8282611c96565b91505060c0840151611dab60e08501826001600160401b03169052565b509392505050565b5f8060408385031215611dc4575f80fd5b82359150611dd4602084016119a3565b90509250929050565b5f60208284031215611ded575f80fd5b81356001600160401b03811115611e02575f80fd5b82016080818503121561030b575f80fd5b5f8060408385031215611e24575f80fd5b8235915060208301358015158114611e3a575f80fd5b809150509250929050565b61ffff60f01b8a60f01b1681525f63ffffffff60e01b808b60e01b166002840152896006840152808960e01b166026840152508651611e8b81602a850160208b01611b9e565b865190830190611ea281602a840160208b01611b9e565b60c087901b6001600160c01b031916602a9290910191820152611ed4603282018660e01b6001600160e01b0319169052565b611eed603682018560e01b6001600160e01b0319169052565b603a019b9a5050505050505050505050565b634e487b7160e01b5f52603260045260245ffd5b5f8351611f24818460208801611b9e565b60609390931b6bffffffffffffffffffffffff19169190920190815260140192915050565b5f8451611f5a818460208901611b9e565b6001600160e01b031960e095861b8116919093019081529290931b16600482015260080192915050565b5f8351611f95818460208801611b9e565b60c09390931b6001600160c01b0319169190920190815260080192915050565b5f8251611fc6818460208701611b9e565b9190910192915050565b5f60208284031215611fe0575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561023b5761023b611fe7565b808202811582820484141761023b5761023b611fe7565b8082018082111561023b5761023b611fe7565b63ffffffff81811683821601908082111561181957611819611fe7565b63ffffffff81811683821602808216919082811461207557612075611fe7565b505092915050565b5f6020828403121561208d575f80fd5b61030b826119be565b5f808335601e198436030181126120ab575f80fd5b8301803591506001600160401b038211156120c4575f80fd5b6020019150600581901b36038213156120db575f80fd5b9250929050565b5f8235605e19833603018112611fc6575f80fd5b5f808335601e1984360301811261210b575f80fd5b8301803591506001600160401b03821115612124575f80fd5b6020019150368190038213156120db575f80fd5b5f60208284031215612148575f80fd5b61030b826119a3565b5f8851612162818460208d01611b9e565b60e089901b6001600160e01b031916908301908152868860048301378681019050600481015f8152858782375060c09390931b6001600160c01b0319166004939094019283019390935250600c01969550505050505056fea164736f6c6343000819000a",
}

// CodecABI is the input ABI used to generate the binding from.
// Deprecated: Use CodecMetaData.ABI instead.
var CodecABI = CodecMetaData.ABI

// CodecBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CodecMetaData.Bin instead.
var CodecBin = CodecMetaData.Bin

// DeployCodec deploys a new Ethereum contract, binding an instance of Codec to it.
func DeployCodec(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Codec, error) {
	parsed, err := CodecMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CodecBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Codec{CodecCaller: CodecCaller{contract: contract}, CodecTransactor: CodecTransactor{contract: contract}, CodecFilterer: CodecFilterer{contract: contract}}, nil
}

// Codec is an auto generated Go binding around an Ethereum contract.
type Codec struct {
	CodecCaller     // Read-only binding to the contract
	CodecTransactor // Write-only binding to the contract
	CodecFilterer   // Log filterer for contract events
}

// CodecCaller is an auto generated read-only Go binding around an Ethereum contract.
type CodecCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CodecTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CodecTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CodecFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CodecFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CodecSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CodecSession struct {
	Contract     *Codec            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CodecCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CodecCallerSession struct {
	Contract *CodecCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CodecTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CodecTransactorSession struct {
	Contract     *CodecTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CodecRaw is an auto generated low-level Go binding around an Ethereum contract.
type CodecRaw struct {
	Contract *Codec // Generic contract binding to access the raw methods on
}

// CodecCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CodecCallerRaw struct {
	Contract *CodecCaller // Generic read-only contract binding to access the raw methods on
}

// CodecTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CodecTransactorRaw struct {
	Contract *CodecTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCodec creates a new instance of Codec, bound to a specific deployed contract.
func NewCodec(address common.Address, backend bind.ContractBackend) (*Codec, error) {
	contract, err := bindCodec(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Codec{CodecCaller: CodecCaller{contract: contract}, CodecTransactor: CodecTransactor{contract: contract}, CodecFilterer: CodecFilterer{contract: contract}}, nil
}

// NewCodecCaller creates a new read-only instance of Codec, bound to a specific deployed contract.
func NewCodecCaller(address common.Address, caller bind.ContractCaller) (*CodecCaller, error) {
	contract, err := bindCodec(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CodecCaller{contract: contract}, nil
}

// NewCodecTransactor creates a new write-only instance of Codec, bound to a specific deployed contract.
func NewCodecTransactor(address common.Address, transactor bind.ContractTransactor) (*CodecTransactor, error) {
	contract, err := bindCodec(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CodecTransactor{contract: contract}, nil
}

// NewCodecFilterer creates a new log filterer instance of Codec, bound to a specific deployed contract.
func NewCodecFilterer(address common.Address, filterer bind.ContractFilterer) (*CodecFilterer, error) {
	contract, err := bindCodec(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CodecFilterer{contract: contract}, nil
}

// bindCodec binds a generic wrapper to an already deployed contract.
func bindCodec(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CodecMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Codec *CodecRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Codec.Contract.CodecCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Codec *CodecRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Codec.Contract.CodecTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Codec *CodecRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Codec.Contract.CodecTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Codec *CodecCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Codec.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Codec *CodecTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Codec.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Codec *CodecTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Codec.Contract.contract.Transact(opts, method, params...)
}

// PackRegisterSubnetValidatorMessage is a free data retrieval call binding the contract method 0x01bbec74.
//
// Solidity: function packRegisterSubnetValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_Codec *CodecCaller) PackRegisterSubnetValidatorMessage(opts *bind.CallOpts, validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	var out []interface{}
	err := _Codec.contract.Call(opts, &out, "packRegisterSubnetValidatorMessage", validationPeriod)

	if err != nil {
		return *new([32]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// PackRegisterSubnetValidatorMessage is a free data retrieval call binding the contract method 0x01bbec74.
//
// Solidity: function packRegisterSubnetValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_Codec *CodecSession) PackRegisterSubnetValidatorMessage(validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	return _Codec.Contract.PackRegisterSubnetValidatorMessage(&_Codec.CallOpts, validationPeriod)
}

// PackRegisterSubnetValidatorMessage is a free data retrieval call binding the contract method 0x01bbec74.
//
// Solidity: function packRegisterSubnetValidatorMessage((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64) validationPeriod) pure returns(bytes32, bytes)
func (_Codec *CodecCallerSession) PackRegisterSubnetValidatorMessage(validationPeriod ValidatorMessagesValidationPeriod) ([32]byte, []byte, error) {
	return _Codec.Contract.PackRegisterSubnetValidatorMessage(&_Codec.CallOpts, validationPeriod)
}

// PackSubnetConversionData is a free data retrieval call binding the contract method 0xf65e4b33.
//
// Solidity: function packSubnetConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData) pure returns(bytes)
func (_Codec *CodecCaller) PackSubnetConversionData(opts *bind.CallOpts, subnetConversionData SubnetConversionData) ([]byte, error) {
	var out []interface{}
	err := _Codec.contract.Call(opts, &out, "packSubnetConversionData", subnetConversionData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackSubnetConversionData is a free data retrieval call binding the contract method 0xf65e4b33.
//
// Solidity: function packSubnetConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData) pure returns(bytes)
func (_Codec *CodecSession) PackSubnetConversionData(subnetConversionData SubnetConversionData) ([]byte, error) {
	return _Codec.Contract.PackSubnetConversionData(&_Codec.CallOpts, subnetConversionData)
}

// PackSubnetConversionData is a free data retrieval call binding the contract method 0xf65e4b33.
//
// Solidity: function packSubnetConversionData((bytes32,bytes32,address,(bytes,bytes,uint64)[]) subnetConversionData) pure returns(bytes)
func (_Codec *CodecCallerSession) PackSubnetConversionData(subnetConversionData SubnetConversionData) ([]byte, error) {
	return _Codec.Contract.PackSubnetConversionData(&_Codec.CallOpts, subnetConversionData)
}

// PackSubnetConversionMessage is a free data retrieval call binding the contract method 0x1fd979c7.
//
// Solidity: function packSubnetConversionMessage(bytes32 subnetConversionID) pure returns(bytes)
func (_Codec *CodecCaller) PackSubnetConversionMessage(opts *bind.CallOpts, subnetConversionID [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Codec.contract.Call(opts, &out, "packSubnetConversionMessage", subnetConversionID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackSubnetConversionMessage is a free data retrieval call binding the contract method 0x1fd979c7.
//
// Solidity: function packSubnetConversionMessage(bytes32 subnetConversionID) pure returns(bytes)
func (_Codec *CodecSession) PackSubnetConversionMessage(subnetConversionID [32]byte) ([]byte, error) {
	return _Codec.Contract.PackSubnetConversionMessage(&_Codec.CallOpts, subnetConversionID)
}

// PackSubnetConversionMessage is a free data retrieval call binding the contract method 0x1fd979c7.
//
// Solidity: function packSubnetConversionMessage(bytes32 subnetConversionID) pure returns(bytes)
func (_Codec *CodecCallerSession) PackSubnetConversionMessage(subnetConversionID [32]byte) ([]byte, error) {
	return _Codec.Contract.PackSubnetConversionMessage(&_Codec.CallOpts, subnetConversionID)
}

// PackSubnetValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xfa1f8dfb.
//
// Solidity: function packSubnetValidatorRegistrationMessage(bytes32 validationID, bool valid) pure returns(bytes)
func (_Codec *CodecCaller) PackSubnetValidatorRegistrationMessage(opts *bind.CallOpts, validationID [32]byte, valid bool) ([]byte, error) {
	var out []interface{}
	err := _Codec.contract.Call(opts, &out, "packSubnetValidatorRegistrationMessage", validationID, valid)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackSubnetValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xfa1f8dfb.
//
// Solidity: function packSubnetValidatorRegistrationMessage(bytes32 validationID, bool valid) pure returns(bytes)
func (_Codec *CodecSession) PackSubnetValidatorRegistrationMessage(validationID [32]byte, valid bool) ([]byte, error) {
	return _Codec.Contract.PackSubnetValidatorRegistrationMessage(&_Codec.CallOpts, validationID, valid)
}

// PackSubnetValidatorRegistrationMessage is a free data retrieval call binding the contract method 0xfa1f8dfb.
//
// Solidity: function packSubnetValidatorRegistrationMessage(bytes32 validationID, bool valid) pure returns(bytes)
func (_Codec *CodecCallerSession) PackSubnetValidatorRegistrationMessage(validationID [32]byte, valid bool) ([]byte, error) {
	return _Codec.Contract.PackSubnetValidatorRegistrationMessage(&_Codec.CallOpts, validationID, valid)
}

// PackSubnetValidatorWeightMessage is a free data retrieval call binding the contract method 0x8545c16a.
//
// Solidity: function packSubnetValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_Codec *CodecCaller) PackSubnetValidatorWeightMessage(opts *bind.CallOpts, validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	var out []interface{}
	err := _Codec.contract.Call(opts, &out, "packSubnetValidatorWeightMessage", validationID, nonce, weight)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackSubnetValidatorWeightMessage is a free data retrieval call binding the contract method 0x8545c16a.
//
// Solidity: function packSubnetValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_Codec *CodecSession) PackSubnetValidatorWeightMessage(validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	return _Codec.Contract.PackSubnetValidatorWeightMessage(&_Codec.CallOpts, validationID, nonce, weight)
}

// PackSubnetValidatorWeightMessage is a free data retrieval call binding the contract method 0x8545c16a.
//
// Solidity: function packSubnetValidatorWeightMessage(bytes32 validationID, uint64 nonce, uint64 weight) pure returns(bytes)
func (_Codec *CodecCallerSession) PackSubnetValidatorWeightMessage(validationID [32]byte, nonce uint64, weight uint64) ([]byte, error) {
	return _Codec.Contract.PackSubnetValidatorWeightMessage(&_Codec.CallOpts, validationID, nonce, weight)
}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_Codec *CodecCaller) PackValidationUptimeMessage(opts *bind.CallOpts, validationID [32]byte, uptime uint64) ([]byte, error) {
	var out []interface{}
	err := _Codec.contract.Call(opts, &out, "packValidationUptimeMessage", validationID, uptime)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_Codec *CodecSession) PackValidationUptimeMessage(validationID [32]byte, uptime uint64) ([]byte, error) {
	return _Codec.Contract.PackValidationUptimeMessage(&_Codec.CallOpts, validationID, uptime)
}

// PackValidationUptimeMessage is a free data retrieval call binding the contract method 0xe1d68f30.
//
// Solidity: function packValidationUptimeMessage(bytes32 validationID, uint64 uptime) pure returns(bytes)
func (_Codec *CodecCallerSession) PackValidationUptimeMessage(validationID [32]byte, uptime uint64) ([]byte, error) {
	return _Codec.Contract.PackValidationUptimeMessage(&_Codec.CallOpts, validationID, uptime)
}

// UnpackRegisterSubnetValidatorMessage is a free data retrieval call binding the contract method 0xa5233770.
//
// Solidity: function unpackRegisterSubnetValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_Codec *CodecCaller) UnpackRegisterSubnetValidatorMessage(opts *bind.CallOpts, input []byte) (ValidatorMessagesValidationPeriod, error) {
	var out []interface{}
	err := _Codec.contract.Call(opts, &out, "unpackRegisterSubnetValidatorMessage", input)

	if err != nil {
		return *new(ValidatorMessagesValidationPeriod), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorMessagesValidationPeriod)).(*ValidatorMessagesValidationPeriod)

	return out0, err

}

// UnpackRegisterSubnetValidatorMessage is a free data retrieval call binding the contract method 0xa5233770.
//
// Solidity: function unpackRegisterSubnetValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_Codec *CodecSession) UnpackRegisterSubnetValidatorMessage(input []byte) (ValidatorMessagesValidationPeriod, error) {
	return _Codec.Contract.UnpackRegisterSubnetValidatorMessage(&_Codec.CallOpts, input)
}

// UnpackRegisterSubnetValidatorMessage is a free data retrieval call binding the contract method 0xa5233770.
//
// Solidity: function unpackRegisterSubnetValidatorMessage(bytes input) pure returns((bytes32,bytes,bytes,uint64,(uint32,address[]),(uint32,address[]),uint64))
func (_Codec *CodecCallerSession) UnpackRegisterSubnetValidatorMessage(input []byte) (ValidatorMessagesValidationPeriod, error) {
	return _Codec.Contract.UnpackRegisterSubnetValidatorMessage(&_Codec.CallOpts, input)
}

// UnpackSubnetConversionMessage is a free data retrieval call binding the contract method 0x1e6d9789.
//
// Solidity: function unpackSubnetConversionMessage(bytes input) pure returns(bytes32)
func (_Codec *CodecCaller) UnpackSubnetConversionMessage(opts *bind.CallOpts, input []byte) ([32]byte, error) {
	var out []interface{}
	err := _Codec.contract.Call(opts, &out, "unpackSubnetConversionMessage", input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UnpackSubnetConversionMessage is a free data retrieval call binding the contract method 0x1e6d9789.
//
// Solidity: function unpackSubnetConversionMessage(bytes input) pure returns(bytes32)
func (_Codec *CodecSession) UnpackSubnetConversionMessage(input []byte) ([32]byte, error) {
	return _Codec.Contract.UnpackSubnetConversionMessage(&_Codec.CallOpts, input)
}

// UnpackSubnetConversionMessage is a free data retrieval call binding the contract method 0x1e6d9789.
//
// Solidity: function unpackSubnetConversionMessage(bytes input) pure returns(bytes32)
func (_Codec *CodecCallerSession) UnpackSubnetConversionMessage(input []byte) ([32]byte, error) {
	return _Codec.Contract.UnpackSubnetConversionMessage(&_Codec.CallOpts, input)
}

// UnpackSubnetValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x2e43ceb5.
//
// Solidity: function unpackSubnetValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_Codec *CodecCaller) UnpackSubnetValidatorRegistrationMessage(opts *bind.CallOpts, input []byte) ([32]byte, bool, error) {
	var out []interface{}
	err := _Codec.contract.Call(opts, &out, "unpackSubnetValidatorRegistrationMessage", input)

	if err != nil {
		return *new([32]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// UnpackSubnetValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x2e43ceb5.
//
// Solidity: function unpackSubnetValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_Codec *CodecSession) UnpackSubnetValidatorRegistrationMessage(input []byte) ([32]byte, bool, error) {
	return _Codec.Contract.UnpackSubnetValidatorRegistrationMessage(&_Codec.CallOpts, input)
}

// UnpackSubnetValidatorRegistrationMessage is a free data retrieval call binding the contract method 0x2e43ceb5.
//
// Solidity: function unpackSubnetValidatorRegistrationMessage(bytes input) pure returns(bytes32, bool)
func (_Codec *CodecCallerSession) UnpackSubnetValidatorRegistrationMessage(input []byte) ([32]byte, bool, error) {
	return _Codec.Contract.UnpackSubnetValidatorRegistrationMessage(&_Codec.CallOpts, input)
}

// UnpackSubnetValidatorWeightMessage is a free data retrieval call binding the contract method 0x9de23d40.
//
// Solidity: function unpackSubnetValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_Codec *CodecCaller) UnpackSubnetValidatorWeightMessage(opts *bind.CallOpts, input []byte) ([32]byte, uint64, uint64, error) {
	var out []interface{}
	err := _Codec.contract.Call(opts, &out, "unpackSubnetValidatorWeightMessage", input)

	if err != nil {
		return *new([32]byte), *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// UnpackSubnetValidatorWeightMessage is a free data retrieval call binding the contract method 0x9de23d40.
//
// Solidity: function unpackSubnetValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_Codec *CodecSession) UnpackSubnetValidatorWeightMessage(input []byte) ([32]byte, uint64, uint64, error) {
	return _Codec.Contract.UnpackSubnetValidatorWeightMessage(&_Codec.CallOpts, input)
}

// UnpackSubnetValidatorWeightMessage is a free data retrieval call binding the contract method 0x9de23d40.
//
// Solidity: function unpackSubnetValidatorWeightMessage(bytes input) pure returns(bytes32, uint64, uint64)
func (_Codec *CodecCallerSession) UnpackSubnetValidatorWeightMessage(input []byte) ([32]byte, uint64, uint64, error) {
	return _Codec.Contract.UnpackSubnetValidatorWeightMessage(&_Codec.CallOpts, input)
}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_Codec *CodecCaller) UnpackValidationUptimeMessage(opts *bind.CallOpts, input []byte) ([32]byte, uint64, error) {
	var out []interface{}
	err := _Codec.contract.Call(opts, &out, "unpackValidationUptimeMessage", input)

	if err != nil {
		return *new([32]byte), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_Codec *CodecSession) UnpackValidationUptimeMessage(input []byte) ([32]byte, uint64, error) {
	return _Codec.Contract.UnpackValidationUptimeMessage(&_Codec.CallOpts, input)
}

// UnpackValidationUptimeMessage is a free data retrieval call binding the contract method 0x088c2463.
//
// Solidity: function unpackValidationUptimeMessage(bytes input) pure returns(bytes32, uint64)
func (_Codec *CodecCallerSession) UnpackValidationUptimeMessage(input []byte) ([32]byte, uint64, error) {
	return _Codec.Contract.UnpackValidationUptimeMessage(&_Codec.CallOpts, input)
}
