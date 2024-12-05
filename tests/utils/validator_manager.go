package utils

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/proto/pb/platformvm"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/stakeable"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpMessage "github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	pwallet "github.com/ava-labs/avalanchego/wallet/chain/p/wallet"
	"github.com/ava-labs/awm-relayer/signature-aggregator/aggregator"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	subnetEvmUtils "github.com/ava-labs/subnet-evm/tests/utils"
	"github.com/ava-labs/subnet-evm/warp/messages"
	proxyadmin "github.com/ava-labs/teleporter/abi-bindings/go/ProxyAdmin"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/mocks/ExampleERC20"
	acp99validatormanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/ACP99ValidatorManager"
	arguments "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/ArgumentsExternal"
	erc20tokenstakingmanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/ERC20TokenStakingManager"
	examplerewardcalculator "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/ExampleRewardCalculator"
	nativetokenstakingmanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/NativeTokenStakingManager"
	poavalidatormanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/PoAValidatorManager"
	iacp99validatormanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/interfaces/IACP99ValidatorManager"
	iposvalidatormanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/interfaces/IPoSValidatorManager"
	ivalidatormanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/interfaces/IValidatorManager"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/proto"
)

const (
	DefaultMinDelegateFeeBips      uint16 = 1
	DefaultMinStakeDurationSeconds uint64 = 1
	DefaultMinStakeAmount          uint64 = 1e16
	DefaultMaxStakeAmount          uint64 = 10e18
	DefaultMaxStakeMultiplier      uint8  = 4
	DefaultMaxChurnPercentage      uint8  = 20
	DefaultChurnPeriodSeconds      uint64 = 1
	DefaultWeightToValueFactor     uint64 = 1e12
	DefaultPChainAddress           string = "P-local18jma8ppw3nhx5r4ap8clazz0dps7rv5u00z96u"
)

type ValidatorManagerConcreteType int

const (
	PoAValidatorManager ValidatorManagerConcreteType = iota
	ERC20TokenStakingManager
	NativeTokenStakingManager
)

//
// Deployment utils
//

func DeployValidatorManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	securityModuleAddress common.Address,
) (common.Address, *acp99validatormanager.ACP99ValidatorManager) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	acp99validatormanager.ACP99ValidatorManagerBin = acp99validatormanager.ACP99ValidatorManagerMetaData.Bin
	address, tx, _, err := acp99validatormanager.DeployACP99ValidatorManager(opts, subnet.RPCClient, 0)
	Expect(err).Should(BeNil())

	validatorManager, err := acp99validatormanager.NewACP99ValidatorManager(address, subnet.RPCClient)
	Expect(err).Should(BeNil())
	validatorManager.Initialize(
		opts,
		acp99validatormanager.ValidatorManagerSettings{
			SubnetID:               subnet.SubnetID,
			ChurnPeriodSeconds:     DefaultChurnPeriodSeconds,
			MaximumChurnPercentage: DefaultMaxChurnPercentage,
		},
		securityModuleAddress,
	)

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, validatorManager
}

// The senderKey is used as the owner of proxy and PoAValidatorManager contracts
func DeployAndInitializeValidatorManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	managerType ValidatorManagerConcreteType,
	proxy bool,
) (common.Address, *proxyadmin.ProxyAdmin) {
	// Deploy the security module
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	var (
		tx                      *types.Transaction
		securityModuleAddress   common.Address
		validatorManagerAddress common.Address
		validatorManager        *acp99validatormanager.ACP99ValidatorManager
	)
	switch managerType {
	case PoAValidatorManager:
		poavalidatormanager.PoAValidatorManagerBin = poavalidatormanager.PoAValidatorManagerMetaData.Bin
		var poaValidatorManager *poavalidatormanager.PoAValidatorManager
		securityModuleAddress, tx, poaValidatorManager, err = poavalidatormanager.DeployPoAValidatorManager(opts, subnet.RPCClient, 0)
		Expect(err).Should(BeNil())
		WaitForTransactionSuccess(ctx, subnet, tx.Hash())

		// Deploy the validator manager
		validatorManagerAddress, validatorManager = DeployValidatorManager(
			ctx,
			senderKey,
			subnet,
			securityModuleAddress,
		)

		tx, err = poaValidatorManager.Initialize(
			opts,
			PrivateKeyToAddress(senderKey),
		)
		Expect(err).Should(BeNil())
	case ERC20TokenStakingManager:
		erc20Address, _ := DeployExampleERC20(ctx, senderKey, subnet)
		rewardCalculatorAddress, _ := DeployExampleRewardCalculator(
			ctx,
			senderKey,
			subnet,
			uint64(10),
		)
		erc20tokenstakingmanager.ERC20TokenStakingManagerBin = erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.Bin
		var erc20StakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager
		securityModuleAddress, tx, erc20StakingManager, err = erc20tokenstakingmanager.DeployERC20TokenStakingManager(
			opts,
			subnet.RPCClient,
			0,
		)
		Expect(err).Should(BeNil())
		WaitForTransactionSuccess(ctx, subnet, tx.Hash())

		// Deploy the validator manager
		validatorManagerAddress, validatorManager = DeployValidatorManager(
			ctx,
			senderKey,
			subnet,
			securityModuleAddress,
		)

		tx, err = erc20StakingManager.Initialize(
			opts,
			erc20tokenstakingmanager.PoSValidatorManagerSettings{
				ValidatorManager:         validatorManagerAddress,
				MinimumStakeAmount:       big.NewInt(0).SetUint64(DefaultMinStakeAmount),
				MaximumStakeAmount:       big.NewInt(0).SetUint64(DefaultMaxStakeAmount),
				MinimumStakeDuration:     DefaultMinStakeDurationSeconds,
				MinimumDelegationFeeBips: DefaultMinDelegateFeeBips,
				MaximumStakeMultiplier:   DefaultMaxStakeMultiplier,
				WeightToValueFactor:      big.NewInt(0).SetUint64(DefaultWeightToValueFactor),
				RewardCalculator:         rewardCalculatorAddress,
				UptimeBlockchainID:       subnet.BlockchainID,
			},
			erc20Address,
		)
		Expect(err).Should(BeNil())
	case NativeTokenStakingManager:
		rewardCalculatorAddress, _ := DeployExampleRewardCalculator(
			ctx,
			senderKey,
			subnet,
			uint64(10),
		)
		nativetokenstakingmanager.NativeTokenStakingManagerBin = nativetokenstakingmanager.NativeTokenStakingManagerMetaData.Bin
		var nativeStakingManager *nativetokenstakingmanager.NativeTokenStakingManager
		securityModuleAddress, tx, nativeStakingManager, err = nativetokenstakingmanager.DeployNativeTokenStakingManager(
			opts,
			subnet.RPCClient,
			0,
		)
		Expect(err).Should(BeNil())
		WaitForTransactionSuccess(ctx, subnet, tx.Hash())

		// Deploy the validator manager
		validatorManagerAddress, validatorManager = DeployValidatorManager(
			ctx,
			senderKey,
			subnet,
			securityModuleAddress,
		)

		tx, err = nativeStakingManager.Initialize(
			opts,
			nativetokenstakingmanager.PoSValidatorManagerSettings{
				ValidatorManager:         validatorManagerAddress,
				MinimumStakeAmount:       big.NewInt(0).SetUint64(DefaultMinStakeAmount),
				MaximumStakeAmount:       big.NewInt(0).SetUint64(DefaultMaxStakeAmount),
				MinimumStakeDuration:     DefaultMinStakeDurationSeconds,
				MinimumDelegationFeeBips: DefaultMinDelegateFeeBips,
				MaximumStakeMultiplier:   DefaultMaxStakeMultiplier,
				WeightToValueFactor:      big.NewInt(0).SetUint64(DefaultWeightToValueFactor),
				RewardCalculator:         rewardCalculatorAddress,
				UptimeBlockchainID:       subnet.BlockchainID,
			},
		)
		Expect(err).Should(BeNil())
	}
	WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	// Initialize the validator manager
	tx, err = validatorManager.Initialize(
		opts,
		acp99validatormanager.ValidatorManagerSettings{
			SubnetID:               subnet.SubnetID,
			ChurnPeriodSeconds:     DefaultChurnPeriodSeconds,
			MaximumChurnPercentage: DefaultMaxChurnPercentage,
		},
		securityModuleAddress,
	)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	var proxyAdmin *proxyadmin.ProxyAdmin

	if proxy {
		// Overwrite the manager address with the proxy address
		validatorManagerAddress, proxyAdmin = DeployTransparentUpgradeableProxy(
			ctx,
			subnet,
			senderKey,
			validatorManagerAddress,
		)
	}

	return validatorManagerAddress, proxyAdmin
}

func DeployExampleRewardCalculator(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	rewardBasisPoints uint64,
) (common.Address, *examplerewardcalculator.ExampleRewardCalculator) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, calculator, err := examplerewardcalculator.DeployExampleRewardCalculator(
		opts,
		subnet.RPCClient,
		rewardBasisPoints,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, calculator
}

//
// Validator Set Initialization utils
//

func InitializeValidatorSet(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	validatorManagerAddress common.Address,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
	nodes []Node,
) []ids.ID {
	log.Println("Initializing validator set")
	initialValidators := make([]warpMessage.SubnetToL1ConverstionValidatorData, len(nodes))
	initialValidatorsABI := make([]iacp99validatormanager.InitialValidator, len(nodes))
	for i, node := range nodes {
		initialValidators[i] = warpMessage.SubnetToL1ConverstionValidatorData{
			NodeID:       node.NodeID.Bytes(),
			BLSPublicKey: node.NodePoP.PublicKey,
			Weight:       nodes[i].Weight,
		}
		initialValidatorsABI[i] = iacp99validatormanager.InitialValidator{
			NodeID:       node.NodeID.Bytes(),
			BlsPublicKey: node.NodePoP.PublicKey[:],
			Weight:       nodes[i].Weight,
		}
	}

	subnetConversionData := warpMessage.SubnetToL1ConversionData{
		SubnetID:       subnetInfo.SubnetID,
		ManagerChainID: subnetInfo.BlockchainID,
		ManagerAddress: validatorManagerAddress[:],
		Validators:     initialValidators,
	}
	subnetConversionDataABI := iacp99validatormanager.ConversionData{
		SubnetID:                     subnetInfo.SubnetID,
		ValidatorManagerBlockchainID: subnetInfo.BlockchainID,
		ValidatorManagerAddress:      validatorManagerAddress,
		InitialValidators:            initialValidatorsABI,
	}
	subnetConversionID, err := warpMessage.SubnetToL1ConversionID(subnetConversionData)
	Expect(err).Should(BeNil())
	subnetConversionSignedMessage := ConstructSubnetConversionMessage(
		subnetConversionID,
		subnetInfo,
		pChainInfo,
		networkID,
		signatureAggregator,
	)
	// Deliver the Warp message to the subnet
	receipt := DeliverSubnetConversion(
		ctx,
		senderKey,
		subnetInfo,
		validatorManagerAddress,
		subnetConversionSignedMessage,
		subnetConversionDataABI,
	)
	manager, err := ivalidatormanager.NewIValidatorManager(validatorManagerAddress, subnetInfo.RPCClient)
	Expect(err).Should(BeNil())
	initialValidatorCreatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		manager.ParseInitialValidatorCreated,
	)
	Expect(err).Should(BeNil())
	var validationIDs []ids.ID
	for i := range nodes {
		validationIDs = append(validationIDs, subnetInfo.SubnetID.Append(uint32(i)))
	}

	Expect(initialValidatorCreatedEvent.Weight).Should(Equal(nodes[0].Weight))

	emittedValidationID := ids.ID(initialValidatorCreatedEvent.ValidationID)
	Expect(emittedValidationID).Should(Equal(validationIDs[0]))

	return validationIDs
}

func DeliverSubnetConversion(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	validatorManagerAddress common.Address,
	subnetConversionSignedMessage *avalancheWarp.Message,
	subnetConversionData iacp99validatormanager.ConversionData,
) *types.Receipt {
	abi, err := iacp99validatormanager.IACP99ValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("initializeValidatorSet", subnetConversionData, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		subnet,
		validatorManagerAddress,
		subnetConversionSignedMessage.Bytes(),
	)
}

//
// Function call utils
//

// func InitializeNativeValidatorRegistration(
// 	ctx context.Context,
// 	senderKey *ecdsa.PrivateKey,
// 	subnet interfaces.SubnetTestInfo,
// 	stakeAmount *big.Int,
// 	node Node,
// 	expiry uint64,
// 	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
// ) (*types.Receipt, ids.ID) {
// 	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
// 	Expect(err).Should(BeNil())
// 	opts.Value = stakeAmount

// 	tx, err := stakingManager.InitializeValidatorRegistration(
// 		opts,
// 		nativetokenstakingmanager.ValidatorRegistrationInput{
// 			NodeID:             node.NodeID[:],
// 			RegistrationExpiry: expiry,
// 			BlsPublicKey:       node.NodePoP.PublicKey[:],
// 		},
// 		DefaultMinDelegateFeeBips,
// 		DefaultMinStakeDurationSeconds,
// 	)
// 	Expect(err).Should(BeNil())
// 	receipt := WaitForTransactionSuccess(ctx, subnet, tx.Hash())
// 	registrationInitiatedEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		stakingManager.ParseValidationPeriodCreated,
// 	)
// 	Expect(err).Should(BeNil())
// 	return receipt, ids.ID(registrationInitiatedEvent.ValidationID)
// }

func InitializeACP99ERC20ValidatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	weight uint64,
	token *exampleerc20.ExampleERC20,
	stakingManagerAddress common.Address,
	node Node,
	expiry uint64,
	stakingManager *acp99validatormanager.ACP99ValidatorManager,
) (*types.Receipt, ids.ID) {
	securityModule, err := stakingManager.GetSecurityModule(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	erc20Module, err := erc20tokenstakingmanager.NewERC20TokenStakingManager(securityModule, subnet.RPCClient)
	Expect(err).Should(BeNil())

	stakeAmount, err := erc20Module.WeightToValue(
		&bind.CallOpts{},
		node.Weight,
	)
	Expect(err).Should(BeNil())

	// Approve ERC20 transfer to the validator manager, which delegatecall's the security module
	ERC20Approve(
		ctx,
		token,
		securityModule,
		stakeAmount,
		subnet,
		senderKey,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())

	args := arguments.InitializeValidatorRegistrationArgs{
		DelegationFeeBips: DefaultMinDelegateFeeBips,
		MinStakeDuration:  DefaultMinStakeDurationSeconds,
	}
	argsBytes, err := args.Pack()
	Expect(err).Should(BeNil())

	tx, err := stakingManager.InitializeValidatorRegistration(
		opts,
		acp99validatormanager.ValidatorRegistrationInput{
			NodeID:             node.NodeID[:],
			RegistrationExpiry: expiry,
			BlsPublicKey:       node.NodePoP.PublicKey[:],
		},
		weight,
		argsBytes,
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	registrationInitiatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidationPeriodCreated,
	)
	Expect(err).Should(BeNil())
	return receipt, ids.ID(registrationInitiatedEvent.ValidationID)
}

// func InitializeERC20ValidatorRegistration(
// 	ctx context.Context,
// 	senderKey *ecdsa.PrivateKey,
// 	subnet interfaces.SubnetTestInfo,
// 	stakeAmount *big.Int,
// 	token *exampleerc20.ExampleERC20,
// 	stakingManagerAddress common.Address,
// 	node Node,
// 	expiry uint64,
// 	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
// ) (*types.Receipt, ids.ID) {
// 	ERC20Approve(
// 		ctx,
// 		token,
// 		stakingManagerAddress,
// 		stakeAmount,
// 		subnet,
// 		senderKey,
// 	)

// 	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
// 	Expect(err).Should(BeNil())

// 	tx, err := stakingManager.InitializeValidatorRegistration(
// 		opts,
// 		erc20tokenstakingmanager.ValidatorRegistrationInput{
// 			NodeID:             node.NodeID[:],
// 			RegistrationExpiry: expiry,
// 			BlsPublicKey:       node.NodePoP.PublicKey[:],
// 		},
// 		DefaultMinDelegateFeeBips,
// 		DefaultMinStakeDurationSeconds,
// 		stakeAmount,
// 	)
// 	Expect(err).Should(BeNil())
// 	receipt := WaitForTransactionSuccess(ctx, subnet, tx.Hash())
// 	registrationInitiatedEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		stakingManager.ParseValidationPeriodCreated,
// 	)
// 	Expect(err).Should(BeNil())
// 	return receipt, ids.ID(registrationInitiatedEvent.ValidationID)
// }

// func InitializePoAValidatorRegistration(
// 	ctx context.Context,
// 	senderKey *ecdsa.PrivateKey,
// 	subnet interfaces.SubnetTestInfo,
// 	node Node,
// 	expiry uint64,
// 	validatorManager *poavalidatormanager.PoAValidatorManager,
// ) (*types.Receipt, ids.ID) {
// 	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
// 	Expect(err).Should(BeNil())

// 	tx, err := validatorManager.InitializeValidatorRegistration(
// 		opts,
// 		poavalidatormanager.ValidatorRegistrationInput{
// 			NodeID:             node.NodeID[:],
// 			RegistrationExpiry: expiry,
// 			BlsPublicKey:       node.NodePoP.PublicKey[:],
// 		},
// 		node.Weight,
// 	)
// 	Expect(err).Should(BeNil())
// 	receipt := WaitForTransactionSuccess(ctx, subnet, tx.Hash())
// 	registrationInitiatedEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		validatorManager.ParseValidationPeriodCreated,
// 	)
// 	Expect(err).Should(BeNil())
// 	return receipt, ids.ID(registrationInitiatedEvent.ValidationID)
// }

func CompleteValidatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := iacp99validatormanager.IACP99ValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeValidatorRegistration", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		subnet,
		stakingManagerContractAddress,
		registrationSignedMessage.Bytes(),
	)
}

// Calls a method that retreived a signed Warp message from the transaction's access list
func CallWarpReceiver(
	ctx context.Context,
	callData []byte,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	contract common.Address,
	signedMessageBytes []byte,
) *types.Receipt {
	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, subnet, PrivateKeyToAddress(senderKey))
	registrationTx := predicateutils.NewPredicateTx(
		subnet.EVMChainID,
		nonce,
		&contract,
		2_000_000,
		gasFeeCap,
		gasTipCap,
		big.NewInt(0),
		callData,
		types.AccessList{},
		warp.ContractAddress,
		signedMessageBytes,
	)
	signedRegistrationTx := SignTransaction(registrationTx, senderKey, subnet.EVMChainID)
	return SendTransactionAndWaitForSuccess(ctx, subnet, signedRegistrationTx)
}

// func InitializeAndCompleteNativeValidatorRegistration(
// 	ctx context.Context,
// 	signatureAggregator *aggregator.SignatureAggregator,
// 	fundedKey *ecdsa.PrivateKey,
// 	subnetInfo interfaces.SubnetTestInfo,
// 	pChainInfo interfaces.SubnetTestInfo,
// 	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
// 	stakingManagerContractAddress common.Address,
// 	expiry uint64,
// 	node Node,
// 	pchainWallet pwallet.Wallet,
// 	networkID uint32,
// ) ids.ID {
// 	stakeAmount, err := stakingManager.WeightToValue(
// 		&bind.CallOpts{},
// 		node.Weight,
// 	)
// 	Expect(err).Should(BeNil())
// 	// Initiate validator registration
// 	receipt, validationID := InitializeNativeValidatorRegistration(
// 		ctx,
// 		fundedKey,
// 		subnetInfo,
// 		stakeAmount,
// 		node,
// 		expiry,
// 		stakingManager,
// 	)

// 	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
// 	// (Sending to the P-Chain will be skipped for now)
// 	signedWarpMessage := ConstructSignedWarpMessage(ctx, receipt, subnetInfo, pChainInfo, nil, signatureAggregator)

// 	_, err = pchainWallet.IssueRegisterL1ValidatorTx(
// 		100*units.Avax,
// 		node.NodePoP.ProofOfPossession,
// 		signedWarpMessage.Bytes(),
// 	)
// 	Expect(err).Should(BeNil())
// 	PChainProposerVMWorkaround(pchainWallet)
// 	AdvanceProposerVM(ctx, subnetInfo, fundedKey, 5)

// 	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
// 	log.Println("Completing validator registration")
// 	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
// 		validationID,
// 		expiry,
// 		node,
// 		true,
// 		subnetInfo,
// 		pChainInfo,
// 		networkID,
// 		signatureAggregator,
// 	)

// 	// Deliver the Warp message to the subnet
// 	receipt = CompleteValidatorRegistration(
// 		ctx,
// 		fundedKey,
// 		subnetInfo,
// 		stakingManagerContractAddress,
// 		registrationSignedMessage,
// 	)
// 	// Check that the validator is registered in the staking contract
// 	registrationEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		stakingManager.ParseValidationPeriodRegistered,
// 	)
// 	Expect(err).Should(BeNil())
// 	Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))

// 	return validationID
// }

func InitializeAndCompleteACP99ERC20ValidatorRegistration(
	ctx context.Context,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *acp99validatormanager.ACP99ValidatorManager,
	stakingManagerAddress common.Address,
	erc20 *exampleerc20.ExampleERC20,
	expiry uint64,
	node Node,
	pchainWallet pwallet.Wallet,
	networkID uint32,
) ids.ID {
	// Initiate validator registration
	var receipt *types.Receipt
	log.Println("Initializing validator  v registration")
	receipt, validationID := InitializeACP99ERC20ValidatorRegistration(
		ctx,
		fundedKey,
		subnetInfo,
		node.Weight,
		erc20,
		stakingManagerAddress,
		node,
		expiry,
		stakingManager,
	)

	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
	signedWarpMessage := ConstructSignedWarpMessage(ctx, receipt, subnetInfo, pChainInfo, nil, signatureAggregator)

	_, err := pchainWallet.IssueRegisterL1ValidatorTx(
		100*units.Avax,
		node.NodePoP.ProofOfPossession,
		signedWarpMessage.Bytes(),
	)
	Expect(err).Should(BeNil())
	PChainProposerVMWorkaround(pchainWallet)
	AdvanceProposerVM(ctx, subnetInfo, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing validator registration")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		expiry,
		node,
		true,
		subnetInfo,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteValidatorRegistration(
		ctx,
		fundedKey,
		subnetInfo,
		stakingManagerAddress,
		registrationSignedMessage,
	)
	// Check that the validator is registered in the staking contract
	registrationEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidationPeriodRegistered,
	)
	Expect(err).Should(BeNil())
	Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))

	return validationID
}

// func InitializeAndCompleteERC20ValidatorRegistration(
// 	ctx context.Context,
// 	signatureAggregator *aggregator.SignatureAggregator,
// 	fundedKey *ecdsa.PrivateKey,
// 	subnetInfo interfaces.SubnetTestInfo,
// 	pChainInfo interfaces.SubnetTestInfo,
// 	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
// 	stakingManagerAddress common.Address,
// 	erc20 *exampleerc20.ExampleERC20,
// 	expiry uint64,
// 	node Node,
// 	pchainWallet pwallet.Wallet,
// 	networkID uint32,
// ) ids.ID {
// 	stakeAmount, err := stakingManager.WeightToValue(
// 		&bind.CallOpts{},
// 		node.Weight,
// 	)
// 	Expect(err).Should(BeNil())
// 	// Initiate validator registration
// 	var receipt *types.Receipt
// 	log.Println("Initializing validator registration")
// 	receipt, validationID := InitializeERC20ValidatorRegistration(
// 		ctx,
// 		fundedKey,
// 		subnetInfo,
// 		stakeAmount,
// 		erc20,
// 		stakingManagerAddress,
// 		node,
// 		expiry,
// 		stakingManager,
// 	)

// 	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
// 	signedWarpMessage := ConstructSignedWarpMessage(ctx, receipt, subnetInfo, pChainInfo, nil, signatureAggregator)

// 	_, err = pchainWallet.IssueRegisterL1ValidatorTx(
// 		100*units.Avax,
// 		node.NodePoP.ProofOfPossession,
// 		signedWarpMessage.Bytes(),
// 	)
// 	Expect(err).Should(BeNil())
// 	PChainProposerVMWorkaround(pchainWallet)
// 	AdvanceProposerVM(ctx, subnetInfo, fundedKey, 5)

// 	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
// 	log.Println("Completing validator registration")
// 	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
// 		validationID,
// 		expiry,
// 		node,
// 		true,
// 		subnetInfo,
// 		pChainInfo,
// 		networkID,
// 		signatureAggregator,
// 	)

// 	// Deliver the Warp message to the subnet
// 	receipt = CompleteValidatorRegistration(
// 		ctx,
// 		fundedKey,
// 		subnetInfo,
// 		stakingManagerAddress,
// 		registrationSignedMessage,
// 	)
// 	// Check that the validator is registered in the staking contract
// 	registrationEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		stakingManager.ParseValidationPeriodRegistered,
// 	)
// 	Expect(err).Should(BeNil())
// 	Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))

// 	return validationID
// }

// func InitializeAndCompletePoAValidatorRegistration(
// 	ctx context.Context,
// 	signatureAggregator *aggregator.SignatureAggregator,
// 	ownerKey *ecdsa.PrivateKey,
// 	fundedKey *ecdsa.PrivateKey,
// 	subnetInfo interfaces.SubnetTestInfo,
// 	pChainInfo interfaces.SubnetTestInfo,
// 	validatorManager *poavalidatormanager.PoAValidatorManager,
// 	validatorManagerAddress common.Address,
// 	expiry uint64,
// 	node Node,
// 	pchainWallet pwallet.Wallet,
// 	networkID uint32,
// ) ids.ID {
// 	// Initiate validator registration
// 	receipt, validationID := InitializePoAValidatorRegistration(
// 		ctx,
// 		ownerKey,
// 		subnetInfo,
// 		node,
// 		expiry,
// 		validatorManager,
// 	)

// 	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
// 	signedWarpMessage := ConstructSignedWarpMessage(ctx, receipt, subnetInfo, pChainInfo, nil, signatureAggregator)

// 	_, err := pchainWallet.IssueRegisterL1ValidatorTx(
// 		100*units.Avax,
// 		node.NodePoP.ProofOfPossession,
// 		signedWarpMessage.Bytes(),
// 	)
// 	Expect(err).Should(BeNil())
// 	PChainProposerVMWorkaround(pchainWallet)
// 	AdvanceProposerVM(ctx, subnetInfo, fundedKey, 5)

// 	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
// 	log.Println("Completing validator registration")
// 	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
// 		validationID,
// 		expiry,
// 		node,
// 		true,
// 		subnetInfo,
// 		pChainInfo,
// 		networkID,
// 		signatureAggregator,
// 	)

// 	// Deliver the Warp message to the subnet
// 	receipt = CompleteValidatorRegistration(
// 		ctx,
// 		fundedKey,
// 		subnetInfo,
// 		validatorManagerAddress,
// 		registrationSignedMessage,
// 	)
// 	// Check that the validator is registered in the staking contract
// 	registrationEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		validatorManager.ParseValidationPeriodRegistered,
// 	)
// 	Expect(err).Should(BeNil())
// 	Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))

// 	return validationID
// }

// func InitializeEndPoSValidation(
// 	ctx context.Context,
// 	senderKey *ecdsa.PrivateKey,
// 	subnet interfaces.SubnetTestInfo,
// 	stakingManager *iposvalidatormanager.IPoSValidatorManager,
// 	validationID ids.ID,
// ) *types.Receipt {
// 	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
// 	Expect(err).Should(BeNil())
// 	tx, err := stakingManager.InitializeEndValidation0(
// 		opts,
// 		validationID,
// 		false,
// 		0,
// 	)
// 	Expect(err).Should(BeNil())
// 	return WaitForTransactionSuccess(ctx, subnet, tx.Hash())
// }

func InitializeEndACP99PoSValidation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManager *acp99validatormanager.ACP99ValidatorManager,
	validationID ids.ID,
	force bool,
	includeUptime bool,
	messageIndex uint32,
	rewardRecipient common.Address,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	args := arguments.InitializeEndValidationArgs{
		Force:              force,
		IncludeUptimeProof: includeUptime,
		MessageIndex:       messageIndex,
		RewardRecipient:    rewardRecipient,
	}
	argsPacked, err := args.Pack()
	Expect(err).Should(BeNil())

	tx, err := stakingManager.InitializeEndValidation(
		opts,
		validationID,
		argsPacked,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, subnet, tx.Hash())
}

// func ForceInitializeEndPoSValidation(
// 	ctx context.Context,
// 	senderKey *ecdsa.PrivateKey,
// 	subnet interfaces.SubnetTestInfo,
// 	stakingManager *iposvalidatormanager.IPoSValidatorManager,
// 	validationID ids.ID,
// ) *types.Receipt {
// 	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
// 	Expect(err).Should(BeNil())
// 	tx, err := stakingManager.ForceInitializeEndValidation(
// 		opts,
// 		validationID,
// 		false,
// 		0,
// 	)
// 	Expect(err).Should(BeNil())
// 	return WaitForTransactionSuccess(ctx, subnet, tx.Hash())
// }

func ConstructUptimeProofMessage(
	validationID ids.ID,
	uptime uint64,
	subnet interfaces.SubnetTestInfo,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
) *avalancheWarp.Message {
	uptimePayload, err := messages.NewValidatorUptime(validationID, uptime)
	Expect(err).Should(BeNil())
	addressedCall, err := warpPayload.NewAddressedCall(nil, uptimePayload.Bytes())
	Expect(err).Should(BeNil())
	uptimeProofUnsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		subnet.BlockchainID,
		addressedCall.Bytes(),
	)
	Expect(err).Should(BeNil())

	uptimeProofSignedMessage, err := signatureAggregator.CreateSignedMessage(
		uptimeProofUnsignedMessage,
		nil,
		subnet.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())
	return uptimeProofSignedMessage
}

func InitializeEndACP99PoSValidationWithUptime(
	ctx context.Context,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	uptime uint64,
) *types.Receipt {
	uptimeMsg := ConstructUptimeProofMessage(
		validationID,
		uptime,
		subnet,
		networkID,
		signatureAggregator,
	)
	args := arguments.InitializeEndValidationArgs{
		Force:              true,
		IncludeUptimeProof: true,
		MessageIndex:       0,
		RewardRecipient:    common.Address{},
	}
	argsPacked, err := args.Pack()
	Expect(err).Should(BeNil())

	abi, err := acp99validatormanager.ACP99ValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("initializeEndValidation", validationID, argsPacked)
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		subnet,
		stakingManagerAddress,
		uptimeMsg.Bytes(),
	)
}

func ForceInitializeEndPoSValidationWithUptime(
	ctx context.Context,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	uptime uint64,
) *types.Receipt {
	uptimeMsg := ConstructUptimeProofMessage(
		validationID,
		uptime,
		subnet,
		networkID,
		signatureAggregator,
	)

	abi, err := iposvalidatormanager.IPoSValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("forceInitializeEndValidation", validationID, true, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		subnet,
		stakingManagerAddress,
		uptimeMsg.Bytes(),
	)
}

func InitializeEndPoSValidationWithUptime(
	ctx context.Context,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	uptime uint64,
) *types.Receipt {
	uptimeMsg := ConstructUptimeProofMessage(
		validationID,
		uptime,
		subnet,
		networkID,
		signatureAggregator,
	)

	abi, err := iposvalidatormanager.IPoSValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("initializeEndValidation", validationID, true, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		subnet,
		stakingManagerAddress,
		uptimeMsg.Bytes(),
	)
}

// func InitializeEndPoAValidation(
// 	ctx context.Context,
// 	senderKey *ecdsa.PrivateKey,
// 	subnet interfaces.SubnetTestInfo,
// 	validatorManager *poavalidatormanager.PoAValidatorManager,
// 	validationID ids.ID,
// ) *types.Receipt {
// 	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
// 	Expect(err).Should(BeNil())
// 	tx, err := validatorManager.InitializeEndValidation(
// 		opts,
// 		validationID,
// 	)
// 	Expect(err).Should(BeNil())
// 	return WaitForTransactionSuccess(ctx, subnet, tx.Hash())
// }

func CompleteEndValidation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := iacp99validatormanager.IACP99ValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndValidation", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		subnet,
		stakingManagerContractAddress,
		registrationSignedMessage.Bytes(),
	)
}

// func InitializeERC20DelegatorRegistration(
// 	ctx context.Context,
// 	senderKey *ecdsa.PrivateKey,
// 	subnet interfaces.SubnetTestInfo,
// 	validationID ids.ID,
// 	delegationAmount *big.Int,
// 	token *exampleerc20.ExampleERC20,
// 	stakingManagerAddress common.Address,
// 	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
// ) *types.Receipt {
// 	ERC20Approve(
// 		ctx,
// 		token,
// 		stakingManagerAddress,
// 		delegationAmount,
// 		subnet,
// 		senderKey,
// 	)

// 	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
// 	Expect(err).Should(BeNil())

// 	tx, err := stakingManager.InitializeDelegatorRegistration(
// 		opts,
// 		validationID,
// 		delegationAmount,
// 	)
// 	Expect(err).Should(BeNil())
// 	receipt := WaitForTransactionSuccess(ctx, subnet, tx.Hash())
// 	_, err = GetEventFromLogs(
// 		receipt.Logs,
// 		stakingManager.ParseDelegatorAdded,
// 	)
// 	Expect(err).Should(BeNil())
// 	return receipt
// }

// func InitializeNativeDelegatorRegistration(
// 	ctx context.Context,
// 	senderKey *ecdsa.PrivateKey,
// 	subnet interfaces.SubnetTestInfo,
// 	validationID ids.ID,
// 	delegationAmount *big.Int,
// 	stakingManagerAddress common.Address,
// 	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
// ) *types.Receipt {
// 	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
// 	Expect(err).Should(BeNil())
// 	opts.Value = delegationAmount

// 	tx, err := stakingManager.InitializeDelegatorRegistration(
// 		opts,
// 		validationID,
// 	)
// 	Expect(err).Should(BeNil())
// 	receipt := WaitForTransactionSuccess(ctx, subnet, tx.Hash())
// 	_, err = GetEventFromLogs(
// 		receipt.Logs,
// 		stakingManager.ParseDelegatorAdded,
// 	)
// 	Expect(err).Should(BeNil())
// 	return receipt
// }

func CompleteDelegatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	delegationID ids.ID,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := iposvalidatormanager.IPoSValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeDelegatorRegistration", delegationID, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		subnet,
		stakingManagerContractAddress,
		signedMessage.Bytes(),
	)
}

// func InitializeEndDelegation(
// 	ctx context.Context,
// 	senderKey *ecdsa.PrivateKey,
// 	subnet interfaces.SubnetTestInfo,
// 	stakingManagerAddress common.Address,
// 	delegationID ids.ID,
// ) *types.Receipt {
// 	stakingManager, err := iposvalidatormanager.NewIPoSValidatorManager(stakingManagerAddress, subnet.RPCClient)
// 	Expect(err).Should(BeNil())
// 	WaitMinStakeDuration(ctx, subnet, senderKey)
// 	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
// 	Expect(err).Should(BeNil())
// 	tx, err := stakingManager.ForceInitializeEndDelegation(
// 		opts,
// 		delegationID,
// 		false,
// 		0,
// 	)
// 	Expect(err).Should(BeNil())
// 	return WaitForTransactionSuccess(ctx, subnet, tx.Hash())
// }

func CompleteEndDelegation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	delegationID ids.ID,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := iposvalidatormanager.IPoSValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndDelegation", delegationID, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		subnet,
		stakingManagerContractAddress,
		signedMessage.Bytes(),
	)
}

func InitializeAndCompleteEndInitialACP99PoSValidation(
	ctx context.Context,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *acp99validatormanager.ACP99ValidatorManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	index uint32,
	weight uint64,
	pchainWallet pwallet.Wallet,
	networkID uint32,
) {
	log.Println("Initializing initial validator removal")
	WaitMinStakeDuration(ctx, subnetInfo, fundedKey)
	receipt := InitializeEndACP99PoSValidation(
		ctx,
		fundedKey,
		subnetInfo,
		stakingManager,
		validationID,
		true,
		false,
		0,
		common.Address{},
	)
	ivdrmgr, err := ivalidatormanager.NewIValidatorManager(stakingManagerAddress, subnetInfo.RPCClient)
	Expect(err).Should(BeNil())
	validatorRemovalEvent, err := GetEventFromLogs(
		receipt.Logs,
		ivdrmgr.ParseValidatorRemovalInitialized,
	)
	Expect(err).Should(BeNil())
	Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
	Expect(validatorRemovalEvent.Weight).Should(Equal(weight))

	// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
	// (Sending to the P-Chain will be skipped for now)
	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, subnetInfo)
	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedMessage,
		nil,
		subnetInfo.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())

	// Deliver the Warp message to the P-Chain
	pchainWallet.IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
	PChainProposerVMWorkaround(pchainWallet)
	AdvanceProposerVM(ctx, subnetInfo, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing initial validator removal")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessageForInitialValidator(
		validationID,
		index,
		false,
		subnetInfo,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteEndValidation(
		ctx,
		fundedKey,
		subnetInfo,
		stakingManagerAddress,
		registrationSignedMessage,
	)

	// Check that the validator is has been delisted from the staking contract
	validationEndedEvent, err := GetEventFromLogs(
		receipt.Logs,
		ivdrmgr.ParseValidationPeriodEnded,
	)
	Expect(err).Should(BeNil())
	Expect(validationEndedEvent.ValidationID[:]).Should(Equal(validationID[:]))
}

// func InitializeAndCompleteEndInitialPoSValidation(
// 	ctx context.Context,
// 	signatureAggregator *aggregator.SignatureAggregator,
// 	fundedKey *ecdsa.PrivateKey,
// 	subnetInfo interfaces.SubnetTestInfo,
// 	pChainInfo interfaces.SubnetTestInfo,
// 	stakingManager *iposvalidatormanager.IPoSValidatorManager,
// 	stakingManagerAddress common.Address,
// 	validationID ids.ID,
// 	index uint32,
// 	weight uint64,
// 	pchainWallet pwallet.Wallet,
// 	networkID uint32,
// ) {
// 	log.Println("Initializing initial validator removal")
// 	WaitMinStakeDuration(ctx, subnetInfo, fundedKey)
// 	receipt := ForceInitializeEndPoSValidation(
// 		ctx,
// 		fundedKey,
// 		subnetInfo,
// 		stakingManager,
// 		validationID,
// 	)
// 	validatorRemovalEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		stakingManager.ParseValidatorRemovalInitialized,
// 	)
// 	Expect(err).Should(BeNil())
// 	Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
// 	Expect(validatorRemovalEvent.Weight).Should(Equal(weight))

// 	// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
// 	// (Sending to the P-Chain will be skipped for now)
// 	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, subnetInfo)
// 	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
// 		unsignedMessage,
// 		nil,
// 		subnetInfo.SubnetID,
// 		67,
// 	)
// 	Expect(err).Should(BeNil())

// 	// Deliver the Warp message to the P-Chain
// 	pchainWallet.IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
// 	PChainProposerVMWorkaround(pchainWallet)
// 	AdvanceProposerVM(ctx, subnetInfo, fundedKey, 5)

// 	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
// 	log.Println("Completing initial validator removal")
// 	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessageForInitialValidator(
// 		validationID,
// 		index,
// 		false,
// 		subnetInfo,
// 		pChainInfo,
// 		networkID,
// 		signatureAggregator,
// 	)

// 	// Deliver the Warp message to the subnet
// 	receipt = CompleteEndValidation(
// 		ctx,
// 		fundedKey,
// 		subnetInfo,
// 		stakingManagerAddress,
// 		registrationSignedMessage,
// 	)

// 	// Check that the validator is has been delisted from the staking contract
// 	validationEndedEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		stakingManager.ParseValidationPeriodEnded,
// 	)
// 	Expect(err).Should(BeNil())
// 	Expect(validationEndedEvent.ValidationID[:]).Should(Equal(validationID[:]))
// }

func InitializeAndCompleteEndACP99PoSValidation(
	ctx context.Context,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *acp99validatormanager.ACP99ValidatorManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	expiry uint64,
	node Node,
	nonce uint64,
	includeUptime bool,
	validatorStartTime time.Time,
	pchainWallet pwallet.Wallet,
	networkID uint32,
) {
	log.Println("Initializing validator removal")
	WaitMinStakeDuration(ctx, subnetInfo, fundedKey)

	var receipt *types.Receipt
	if includeUptime {
		uptime := uint64(time.Since(validatorStartTime).Seconds())
		receipt = InitializeEndACP99PoSValidationWithUptime(
			ctx,
			networkID,
			signatureAggregator,
			fundedKey,
			subnetInfo,
			stakingManagerAddress,
			validationID,
			uptime,
		)
	} else {
		receipt = InitializeEndACP99PoSValidation(
			ctx,
			fundedKey,
			subnetInfo,
			stakingManager,
			validationID,
			true,
			false,
			0,
			common.Address{},
		)
	}

	validatorRemovalEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidatorRemovalInitialized,
	)
	Expect(err).Should(BeNil())
	Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
	Expect(validatorRemovalEvent.Weight).Should(Equal(node.Weight))

	// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, subnetInfo)
	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedMessage,
		nil,
		subnetInfo.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())

	// Deliver the Warp message to the P-Chain
	pchainWallet.IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
	PChainProposerVMWorkaround(pchainWallet)
	AdvanceProposerVM(ctx, subnetInfo, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing validator removal")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		expiry,
		node,
		false,
		subnetInfo,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteEndValidation(
		ctx,
		fundedKey,
		subnetInfo,
		stakingManagerAddress,
		registrationSignedMessage,
	)

	// Check that the validator is has been delisted from the staking contract
	registrationEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidationPeriodEnded,
	)
	Expect(err).Should(BeNil())
	Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
}

// func InitializeAndCompleteEndPoSValidation(
// 	ctx context.Context,
// 	signatureAggregator *aggregator.SignatureAggregator,
// 	fundedKey *ecdsa.PrivateKey,
// 	subnetInfo interfaces.SubnetTestInfo,
// 	pChainInfo interfaces.SubnetTestInfo,
// 	stakingManager *iposvalidatormanager.IPoSValidatorManager,
// 	stakingManagerAddress common.Address,
// 	validationID ids.ID,
// 	expiry uint64,
// 	node Node,
// 	nonce uint64,
// 	includeUptime bool,
// 	validatorStartTime time.Time,
// 	pchainWallet pwallet.Wallet,
// 	networkID uint32,
// ) {
// 	log.Println("Initializing validator removal")
// 	WaitMinStakeDuration(ctx, subnetInfo, fundedKey)

// 	var receipt *types.Receipt
// 	if includeUptime {
// 		uptime := uint64(time.Since(validatorStartTime).Seconds())
// 		receipt = ForceInitializeEndPoSValidationWithUptime(
// 			ctx,
// 			networkID,
// 			signatureAggregator,
// 			fundedKey,
// 			subnetInfo,
// 			stakingManagerAddress,
// 			validationID,
// 			uptime,
// 		)
// 	} else {
// 		receipt = ForceInitializeEndPoSValidation(
// 			ctx,
// 			fundedKey,
// 			subnetInfo,
// 			stakingManager,
// 			validationID,
// 		)
// 	}

// 	validatorRemovalEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		stakingManager.ParseValidatorRemovalInitialized,
// 	)
// 	Expect(err).Should(BeNil())
// 	Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
// 	Expect(validatorRemovalEvent.Weight).Should(Equal(node.Weight))

// 	// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
// 	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, subnetInfo)
// 	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
// 		unsignedMessage,
// 		nil,
// 		subnetInfo.SubnetID,
// 		67,
// 	)
// 	Expect(err).Should(BeNil())

// 	// Deliver the Warp message to the P-Chain
// 	pchainWallet.IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
// 	PChainProposerVMWorkaround(pchainWallet)
// 	AdvanceProposerVM(ctx, subnetInfo, fundedKey, 5)

// 	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
// 	log.Println("Completing validator removal")
// 	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
// 		validationID,
// 		expiry,
// 		node,
// 		false,
// 		subnetInfo,
// 		pChainInfo,
// 		networkID,
// 		signatureAggregator,
// 	)

// 	// Deliver the Warp message to the subnet
// 	receipt = CompleteEndValidation(
// 		ctx,
// 		fundedKey,
// 		subnetInfo,
// 		stakingManagerAddress,
// 		registrationSignedMessage,
// 	)

// 	// Check that the validator is has been delisted from the staking contract
// 	registrationEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		stakingManager.ParseValidationPeriodEnded,
// 	)
// 	Expect(err).Should(BeNil())
// 	Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
// }

// func InitializeAndCompleteEndInitialPoAValidation(
// 	ctx context.Context,
// 	signatureAggregator *aggregator.SignatureAggregator,
// 	ownerKey *ecdsa.PrivateKey,
// 	fundedKey *ecdsa.PrivateKey,
// 	subnetInfo interfaces.SubnetTestInfo,
// 	pChainInfo interfaces.SubnetTestInfo,
// 	stakingManager *poavalidatormanager.PoAValidatorManager,
// 	stakingManagerAddress common.Address,
// 	validationID ids.ID,
// 	index uint32,
// 	weight uint64,
// 	pchainWallet pwallet.Wallet,
// 	networkID uint32,
// ) {
// 	log.Println("Initializing initial validator removal")
// 	WaitMinStakeDuration(ctx, subnetInfo, fundedKey)
// 	receipt := InitializeEndPoAValidation(
// 		ctx,
// 		ownerKey,
// 		subnetInfo,
// 		stakingManager,
// 		validationID,
// 	)
// 	validatorRemovalEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		stakingManager.ParseValidatorRemovalInitialized,
// 	)
// 	Expect(err).Should(BeNil())
// 	Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
// 	Expect(validatorRemovalEvent.Weight).Should(Equal(weight))

// 	// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
// 	// (Sending to the P-Chain will be skipped for now)
// 	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, subnetInfo)
// 	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
// 		unsignedMessage,
// 		nil,
// 		subnetInfo.SubnetID,
// 		67,
// 	)
// 	Expect(err).Should(BeNil())

// 	// Deliver the Warp message to the P-Chain
// 	pchainWallet.IssueSetL1ValidatorWeightTx(signedWarpMessage.Bytes())
// 	PChainProposerVMWorkaround(pchainWallet)
// 	AdvanceProposerVM(ctx, subnetInfo, fundedKey, 5)

// 	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
// 	log.Println("Completing initial validator removal")
// 	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessageForInitialValidator(
// 		validationID,
// 		index,
// 		false,
// 		subnetInfo,
// 		pChainInfo,
// 		networkID,
// 		signatureAggregator,
// 	)

// 	// Deliver the Warp message to the subnet
// 	receipt = CompleteEndValidation(
// 		ctx,
// 		fundedKey,
// 		subnetInfo,
// 		stakingManagerAddress,
// 		registrationSignedMessage,
// 	)

// 	// Check that the validator is has been delisted from the staking contract
// 	validationEndedEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		stakingManager.ParseValidationPeriodEnded,
// 	)
// 	Expect(err).Should(BeNil())
// 	Expect(validationEndedEvent.ValidationID[:]).Should(Equal(validationID[:]))
// }

// func InitializeAndCompleteEndPoAValidation(
// 	ctx context.Context,
// 	signatureAggregator *aggregator.SignatureAggregator,
// 	ownerKey *ecdsa.PrivateKey,
// 	fundedKey *ecdsa.PrivateKey,
// 	subnetInfo interfaces.SubnetTestInfo,
// 	pChainInfo interfaces.SubnetTestInfo,
// 	validatorManager *poavalidatormanager.PoAValidatorManager,
// 	validatorManagerAddress common.Address,
// 	validationID ids.ID,
// 	weight uint64,
// 	nonce uint64,
// 	networkID uint32,
// ) {
// 	receipt := InitializeEndPoAValidation(
// 		ctx,
// 		ownerKey,
// 		subnetInfo,
// 		validatorManager,
// 		validationID,
// 	)
// 	validatorRemovalEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		validatorManager.ParseValidatorRemovalInitialized,
// 	)
// 	Expect(err).Should(BeNil())
// 	Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
// 	Expect(validatorRemovalEvent.Weight).Should(Equal(weight))

// 	// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
// 	// (Sending to the P-Chain will be skipped for now)
// 	signedWarpMessage := ConstructSignedWarpMessage(ctx, receipt, subnetInfo, pChainInfo, nil, signatureAggregator)
// 	Expect(err).Should(BeNil())

// 	// Validate the Warp message, (this will be done on the P-Chain in the future)
// 	ValidateSubnetValidatorWeightMessage(signedWarpMessage, validationID, 0, nonce)

// 	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
// 	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
// 		validationID,
// 		0,
// 		Node{},
// 		false,
// 		subnetInfo,
// 		pChainInfo,
// 		networkID,
// 		signatureAggregator,
// 	)

// 	// Deliver the Warp message to the subnet
// 	receipt = CompleteEndValidation(
// 		ctx,
// 		ownerKey,
// 		subnetInfo,
// 		validatorManagerAddress,
// 		registrationSignedMessage,
// 	)

// 	// Check that the validator is has been delisted from the staking contract
// 	registrationEvent, err := GetEventFromLogs(
// 		receipt.Logs,
// 		validatorManager.ParseValidationPeriodEnded,
// 	)
// 	Expect(err).Should(BeNil())
// 	Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
// }

//
// P-Chain utils
//

func ConstructSubnetValidatorRegistrationMessageForInitialValidator(
	validationID ids.ID,
	index uint32,
	valid bool,
	subnet interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
) *avalancheWarp.Message {
	justification := platformvm.L1ValidatorRegistrationJustification{
		Preimage: &platformvm.L1ValidatorRegistrationJustification_ConvertSubnetToL1TxData{
			ConvertSubnetToL1TxData: &platformvm.SubnetIDIndex{
				SubnetId: subnet.SubnetID[:],
				Index:    index,
			},
		},
	}
	justificationBytes, err := proto.Marshal(&justification)
	Expect(err).Should(BeNil())

	registrationPayload, err := warpMessage.NewL1ValidatorRegistration(validationID, valid)
	Expect(err).Should(BeNil())
	registrationAddressedCall, err := warpPayload.NewAddressedCall(nil, registrationPayload.Bytes())
	Expect(err).Should(BeNil())
	registrationUnsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		pChainInfo.BlockchainID,
		registrationAddressedCall.Bytes(),
	)
	Expect(err).Should(BeNil())

	registrationSignedMessage, err := signatureAggregator.CreateSignedMessage(
		registrationUnsignedMessage,
		justificationBytes,
		subnet.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())

	return registrationSignedMessage
}

func ConstructSubnetValidatorRegistrationMessage(
	validationID ids.ID,
	expiry uint64,
	node Node,
	valid bool,
	subnet interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
) *avalancheWarp.Message {
	msg, err := warpMessage.NewRegisterL1Validator(
		subnet.SubnetID,
		node.NodeID,
		node.NodePoP.PublicKey,
		expiry,
		warpMessage.PChainOwner{},
		warpMessage.PChainOwner{},
		node.Weight,
	)
	Expect(err).Should(BeNil())
	justification := platformvm.L1ValidatorRegistrationJustification{
		Preimage: &platformvm.L1ValidatorRegistrationJustification_RegisterL1ValidatorMessage{
			RegisterL1ValidatorMessage: msg.Bytes(),
		},
	}
	justificationBytes, err := proto.Marshal(&justification)
	Expect(err).Should(BeNil())

	registrationPayload, err := warpMessage.NewL1ValidatorRegistration(validationID, valid)
	Expect(err).Should(BeNil())
	registrationAddressedCall, err := warpPayload.NewAddressedCall(nil, registrationPayload.Bytes())
	Expect(err).Should(BeNil())
	registrationUnsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		pChainInfo.BlockchainID,
		registrationAddressedCall.Bytes(),
	)
	Expect(err).Should(BeNil())

	registrationSignedMessage, err := signatureAggregator.CreateSignedMessage(
		registrationUnsignedMessage,
		justificationBytes,
		subnet.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())

	return registrationSignedMessage
}

func ConstructSubnetValidatorWeightUpdateMessage(
	validationID ids.ID,
	nonce uint64,
	weight uint64,
	subnet interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	signatureAggregator *aggregator.SignatureAggregator,
	networkID uint32,
) *avalancheWarp.Message {
	payload, err := warpMessage.NewL1ValidatorWeight(validationID, nonce, weight)
	Expect(err).Should(BeNil())
	updateAddressedCall, err := warpPayload.NewAddressedCall(nil, payload.Bytes())
	Expect(err).Should(BeNil())
	updateUnsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		pChainInfo.BlockchainID,
		updateAddressedCall.Bytes(),
	)
	Expect(err).Should(BeNil())

	updateSignedMessage, err := signatureAggregator.CreateSignedMessage(
		updateUnsignedMessage,
		nil,
		subnet.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())
	return updateSignedMessage
}

func ConstructSubnetConversionMessage(
	subnetConversionID ids.ID,
	subnet interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
) *avalancheWarp.Message {
	subnetConversionPayload, err := warpMessage.NewSubnetToL1Conversion(subnetConversionID)
	Expect(err).Should(BeNil())
	subnetConversionAddressedCall, err := warpPayload.NewAddressedCall(
		nil,
		subnetConversionPayload.Bytes(),
	)
	Expect(err).Should(BeNil())
	subnetConversionUnsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		pChainInfo.BlockchainID,
		subnetConversionAddressedCall.Bytes(),
	)
	Expect(err).Should(BeNil())

	subnetConversionSignedMessage, err := signatureAggregator.CreateSignedMessage(
		subnetConversionUnsignedMessage,
		subnet.SubnetID[:],
		subnet.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())
	return subnetConversionSignedMessage
}

//
// Warp message validiation utils
// These will be replaced by the actual implementation on the P-Chain in the future
//

func ValidateRegisterSubnetValidatorMessage(
	signedWarpMessage *avalancheWarp.Message,
	nodeID ids.ID,
	weight uint64,
	subnetID ids.ID,
	blsPublicKey [bls.PublicKeyLen]byte,
) {
	// Validate the Warp message, (this will be done on the P-Chain in the future)
	msg, err := warpPayload.ParseAddressedCall(signedWarpMessage.UnsignedMessage.Payload)
	Expect(err).Should(BeNil())
	// Check that the addressed call payload is a registered Warp message type
	var payloadInterface warpMessage.Payload
	ver, err := warpMessage.Codec.Unmarshal(msg.Payload, &payloadInterface)
	Expect(err).Should(BeNil())
	payload, ok := payloadInterface.(*warpMessage.RegisterL1Validator)
	Expect(ok).Should(BeTrue())

	Expect(ver).Should(Equal(uint16(warpMessage.CodecVersion)))
	Expect(payload.NodeID).Should(Equal(nodeID))
	Expect(payload.Weight).Should(Equal(weight))
	Expect(payload.SubnetID).Should(Equal(subnetID))
	Expect(payload.BLSPublicKey[:]).Should(Equal(blsPublicKey[:]))
}

func ValidateSubnetValidatorWeightMessage(
	signedWarpMessage *avalancheWarp.Message,
	validationID ids.ID,
	weight uint64,
	nonce uint64,
) {
	msg, err := warpPayload.ParseAddressedCall(signedWarpMessage.UnsignedMessage.Payload)
	Expect(err).Should(BeNil())
	// Check that the addressed call payload is a registered Warp message type
	var payloadInterface warpMessage.Payload
	ver, err := warpMessage.Codec.Unmarshal(msg.Payload, &payloadInterface)
	Expect(err).Should(BeNil())
	payload, ok := payloadInterface.(*warpMessage.L1ValidatorWeight)
	Expect(ok).Should(BeTrue())

	Expect(ver).Should(Equal(uint16(warpMessage.CodecVersion)))
	Expect(payload.ValidationID).Should(Equal(validationID))
	Expect(payload.Weight).Should(Equal(weight))
	Expect(payload.Nonce).Should(Equal(nonce))
}

func WaitMinStakeDuration(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	fundedKey *ecdsa.PrivateKey,
) {
	// Make sure minimum stake duration has passed
	time.Sleep(time.Duration(DefaultMinStakeDurationSeconds) * time.Second)

	// Send a loopback transaction to self to force a block production
	// before delisting the validator.
	SendNativeTransfer(
		ctx,
		subnet,
		fundedKey,
		common.Address{},
		big.NewInt(10),
	)
}

func CalculateSubnetConversionValidationId(subnetID ids.ID, validatorIdx uint32) ids.ID {
	preImage := make([]byte, 36)
	copy(preImage[0:32], subnetID[:])
	binary.BigEndian.PutUint32(preImage[32:36], validatorIdx)
	return sha256.Sum256(preImage)
}

// PackSubnetConversionData defines a packing function that works
// over any struct instance of SubnetConversionData since the abi-bindings
// process generates one for each of the different contracts.
func PackSubnetConversionData(data interface{}) ([]byte, error) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %s", v.Kind())
	}
	// Define required fields and their expected types
	requiredFields := map[string]reflect.Type{
		"SubnetID":                     reflect.TypeOf([32]byte{}),
		"ValidatorManagerBlockchainID": reflect.TypeOf([32]byte{}),
		"ValidatorManagerAddress":      reflect.TypeOf(common.Address{}),
		// InitialValidators is a slice of structs and handled separately
		"InitialValidators": reflect.TypeOf([]struct{}{}),
	}
	// Check for required fields and types
	for fieldName, expectedType := range requiredFields {
		field := v.FieldByName(fieldName)

		if !field.IsValid() {
			return nil, fmt.Errorf("field %s is missing", fieldName)
		}
		// Allow flexible types for InitialValidators by checking it contains structs
		if fieldName == "InitialValidators" {
			if field.Kind() != reflect.Slice || field.Type().Elem().Kind() != reflect.Struct {
				return nil, fmt.Errorf("field %s has incorrect type: expected a slice of structs", fieldName)
			}
		} else {
			if field.Type() != expectedType {
				return nil, fmt.Errorf("field %s has incorrect type: expected %s, got %s", fieldName, expectedType, field.Type())
			}
		}
	}

	subnetID := v.FieldByName("SubnetID").Interface().([32]byte)
	validatorManagerBlockchainID := v.FieldByName("ValidatorManagerBlockchainID").Interface().([32]byte)
	validatorManagerAddress := v.FieldByName("ValidatorManagerAddress").Interface().(common.Address)
	initialValidators := v.FieldByName("InitialValidators")

	// Pack each InitialValidator struct
	packedInitialValidators := make([][]byte, initialValidators.Len())
	var packedInitialValidatorsLen uint32
	for i := 0; i < initialValidators.Len(); i++ {
		iv := initialValidators.Index(i).Interface()
		ivPacked, err := PackInitialValidator(iv)
		if err != nil {
			return nil, fmt.Errorf("failed to pack InitialValidator: %w", err)
		}

		packedInitialValidators[i] = ivPacked
		packedInitialValidatorsLen += uint32(len(ivPacked))
	}

	b := make([]byte, 94+packedInitialValidatorsLen)
	binary.BigEndian.PutUint16(b[0:2], uint16(warpMessage.CodecVersion))
	copy(b[2:34], subnetID[:])
	copy(b[34:66], validatorManagerBlockchainID[:])
	// These are evm addresses and have lengths of 20 so hardcoding here
	binary.BigEndian.PutUint32(b[66:70], uint32(20))
	copy(b[70:90], validatorManagerAddress.Bytes())
	binary.BigEndian.PutUint32(b[90:94], uint32(initialValidators.Len()))
	offset := 94
	for _, ivPacked := range packedInitialValidators {
		copy(b[offset:offset+len(ivPacked)], ivPacked)
		offset += len(ivPacked)
	}

	return b, nil
}

// PackInitialValidator defines a packing function that works
// over any struct instance of InitialValidator since the abi-bindings
// process generates one for each of the different contracts.
func PackInitialValidator(iv interface{}) ([]byte, error) {
	v := reflect.ValueOf(iv)

	// Ensure the passed interface is a struct
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got %s", v.Kind())
	}

	// Define required fields and their expected types
	requiredFields := map[string]reflect.Type{
		"NodeID":       reflect.TypeOf([]byte{}),
		"Weight":       reflect.TypeOf(uint64(0)),
		"BlsPublicKey": reflect.TypeOf([]byte{}),
	}

	// Check for required fields and types
	for fieldName, expectedType := range requiredFields {
		field := v.FieldByName(fieldName)

		if !field.IsValid() {
			return nil, fmt.Errorf("field %s is missing", fieldName)
		}

		if field.Type() != expectedType {
			return nil, fmt.Errorf("field %s has incorrect type: expected %s, got %s", fieldName, expectedType, field.Type())
		}
	}

	// At this point, we know the struct has all required fields with correct types
	// Use reflection to retrieve field values and perform canonical packing
	nodeID := v.FieldByName("NodeID").Interface().([]byte)
	weight := v.FieldByName("Weight").Interface().(uint64)
	blsPublicKey := v.FieldByName("BlsPublicKey").Interface().([]byte)

	b := make([]byte, 60+len(nodeID))
	binary.BigEndian.PutUint32(b[0:4], uint32(len(nodeID)))
	copy(b[4:4+len(nodeID)], nodeID[:])
	binary.BigEndian.PutUint64(b[4+len(nodeID):4+len(nodeID)+8], weight)
	copy(b[4+len(nodeID)+8:4+len(nodeID)+8+48], blsPublicKey)
	return b, nil
}

func PChainProposerVMWorkaround(
	pchainWallet pwallet.Wallet,
) {
	// Workaround current block map rules
	destAddr, err := address.ParseToID(DefaultPChainAddress)
	Expect(err).Should(BeNil())
	log.Println("Waiting for P-Chain...")
	time.Sleep(30 * time.Second)

	pBuilder := pchainWallet.Builder()
	pContext := pBuilder.Context()
	avaxAssetID := pContext.AVAXAssetID
	locktime := uint64(time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC).Unix())
	amount := 500 * units.MilliAvax
	_, err = pchainWallet.IssueBaseTx([]*avax.TransferableOutput{
		{
			Asset: avax.Asset{
				ID: avaxAssetID,
			},
			Out: &stakeable.LockOut{
				Locktime: locktime,
				TransferableOut: &secp256k1fx.TransferOutput{
					Amt: amount,
					OutputOwners: secp256k1fx.OutputOwners{
						Threshold: 1,
						Addrs: []ids.ShortID{
							destAddr,
						},
					},
				},
			},
		},
	})
	Expect(err).Should(BeNil())
	_, err = pchainWallet.IssueBaseTx([]*avax.TransferableOutput{
		{
			Asset: avax.Asset{
				ID: avaxAssetID,
			},
			Out: &stakeable.LockOut{
				Locktime: locktime,
				TransferableOut: &secp256k1fx.TransferOutput{
					Amt: amount,
					OutputOwners: secp256k1fx.OutputOwners{
						Threshold: 1,
						Addrs: []ids.ShortID{
							destAddr,
						},
					},
				},
			},
		},
	})
	Expect(err).Should(BeNil())
	// End workaround
}

func AdvanceProposerVM(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	fundedKey *ecdsa.PrivateKey,
	blocks int,
) {
	log.Println("Advancing proposer VM")
	for i := 0; i < blocks; i++ {
		err := subnetEvmUtils.IssueTxsToActivateProposerVMFork(
			ctx, subnet.EVMChainID, fundedKey, subnet.WSClient,
		)
		Expect(err).Should(BeNil())
	}
}
