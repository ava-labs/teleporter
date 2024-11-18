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
	"sort"
	"time"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/proto/pb/platformvm"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/stakeable"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
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
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/mocks/ExampleERC20"
	erc20tokenstakingmanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/ERC20TokenStakingManager"
	examplerewardcalculator "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/ExampleRewardCalculator"
	nativetokenstakingmanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/NativeTokenStakingManager"
	poavalidatormanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/PoAValidatorManager"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/protobuf/proto"

	. "github.com/onsi/gomega"
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

//
// Deployment utils
//

func DeployNativeTokenStakingManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
) (common.Address, *nativetokenstakingmanager.NativeTokenStakingManager) {
	// Reset the global binary data for better test isolation
	nativetokenstakingmanager.NativeTokenStakingManagerBin =
		nativetokenstakingmanager.NativeTokenStakingManagerMetaData.Bin

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, stakingManager, err := nativetokenstakingmanager.DeployNativeTokenStakingManager(
		opts,
		l1.RPCClient,
		0, // ICMInitializable.Allowed
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, l1, tx.Hash())

	return address, stakingManager
}

func DeployAndInitializeNativeTokenStakingManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
) (common.Address, *nativetokenstakingmanager.NativeTokenStakingManager) {
	stakingManagerContractAddress, stakingManager := DeployNativeTokenStakingManager(
		ctx,
		senderKey,
		l1,
	)
	rewardCalculatorAddress, _ := DeployExampleRewardCalculator(
		ctx,
		senderKey,
		l1,
		uint64(10),
	)
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.Initialize(
		opts,
		nativetokenstakingmanager.PoSValidatorManagerSettings{
			BaseSettings: nativetokenstakingmanager.ValidatorManagerSettings{
				L1ID:                   l1.L1ID,
				ChurnPeriodSeconds:     DefaultChurnPeriodSeconds,
				MaximumChurnPercentage: DefaultMaxChurnPercentage,
			},
			MinimumStakeAmount:       big.NewInt(0).SetUint64(DefaultMinStakeAmount),
			MaximumStakeAmount:       big.NewInt(0).SetUint64(DefaultMaxStakeAmount),
			MinimumStakeDuration:     DefaultMinStakeDurationSeconds,
			MinimumDelegationFeeBips: DefaultMinDelegateFeeBips,
			MaximumStakeMultiplier:   DefaultMaxStakeMultiplier,
			WeightToValueFactor:      big.NewInt(0).SetUint64(DefaultWeightToValueFactor),
			RewardCalculator:         rewardCalculatorAddress,
		},
	)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(ctx, l1, tx.Hash())

	return stakingManagerContractAddress, stakingManager
}

func DeployERC20TokenStakingManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
) (common.Address, *erc20tokenstakingmanager.ERC20TokenStakingManager) {
	// Reset the global binary data for better test isolation
	erc20tokenstakingmanager.ERC20TokenStakingManagerBin =
		erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.Bin

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, stakingManager, err := erc20tokenstakingmanager.DeployERC20TokenStakingManager(
		opts,
		l1.RPCClient,
		0, // ICMInitializable.Allowed
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, l1, tx.Hash())

	return address, stakingManager
}

func DeployAndInitializeERC20TokenStakingManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
) (
	common.Address,
	*erc20tokenstakingmanager.ERC20TokenStakingManager,
	common.Address,
	*exampleerc20.ExampleERC20,
) {
	stakingManagerContractAddress, stakingManager := DeployERC20TokenStakingManager(
		ctx,
		senderKey,
		l1,
	)

	erc20Address, erc20 := DeployExampleERC20(ctx, senderKey, l1)
	rewardCalculatorAddress, _ := DeployExampleRewardCalculator(
		ctx,
		senderKey,
		l1,
		uint64(10),
	)
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.Initialize(
		opts,
		erc20tokenstakingmanager.PoSValidatorManagerSettings{
			BaseSettings: erc20tokenstakingmanager.ValidatorManagerSettings{
				L1ID:                   l1.L1ID,
				ChurnPeriodSeconds:     DefaultChurnPeriodSeconds,
				MaximumChurnPercentage: DefaultMaxChurnPercentage,
			},
			MinimumStakeAmount:       big.NewInt(0).SetUint64(DefaultMinStakeAmount),
			MaximumStakeAmount:       big.NewInt(0).SetUint64(DefaultMaxStakeAmount),
			MinimumStakeDuration:     DefaultMinStakeDurationSeconds,
			MinimumDelegationFeeBips: DefaultMinDelegateFeeBips,
			MaximumStakeMultiplier:   DefaultMaxStakeMultiplier,
			WeightToValueFactor:      big.NewInt(0).SetUint64(DefaultWeightToValueFactor),
			RewardCalculator:         rewardCalculatorAddress,
		},
		erc20Address,
	)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(ctx, l1, tx.Hash())

	return stakingManagerContractAddress, stakingManager, erc20Address, erc20
}

func DeployPoAValidatorManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
) (common.Address, *poavalidatormanager.PoAValidatorManager) {
	// Reset the global binary data for better test isolation
	poavalidatormanager.PoAValidatorManagerBin = poavalidatormanager.PoAValidatorManagerMetaData.Bin

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, validatorManager, err := poavalidatormanager.DeployPoAValidatorManager(
		opts,
		l1.RPCClient,
		0, // ICMInitializable.Allowed
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, l1, tx.Hash())

	return address, validatorManager
}

func DeployAndInitializePoAValidatorManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	ownerAddress common.Address,
) (common.Address, *poavalidatormanager.PoAValidatorManager) {
	validatorManagerAddress, validatorManager := DeployPoAValidatorManager(
		ctx,
		senderKey,
		l1,
	)
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := validatorManager.Initialize(
		opts,
		poavalidatormanager.ValidatorManagerSettings{
			L1ID:                   l1.L1ID,
			ChurnPeriodSeconds:     uint64(0),
			MaximumChurnPercentage: uint8(20),
		},
		ownerAddress,
	)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(ctx, l1, tx.Hash())

	return validatorManagerAddress, validatorManager
}

func DeployExampleRewardCalculator(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	rewardBasisPoints uint64,
) (common.Address, *examplerewardcalculator.ExampleRewardCalculator) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, calculator, err := examplerewardcalculator.DeployExampleRewardCalculator(
		opts,
		l1.RPCClient,
		rewardBasisPoints,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, l1, tx.Hash())

	return address, calculator
}

//
// Validator Set Initialization utils
//

func InitializeNativeTokenValidatorSet(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	validatorManager *nativetokenstakingmanager.NativeTokenStakingManager,
	validatorManagerAddress common.Address,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
	nodes []Node,
) []ids.ID {
	initialValidators := make([]warpMessage.SubnetConversionValidatorData, len(nodes))
	initialValidatorsABI := make([]nativetokenstakingmanager.InitialValidator, len(nodes))
	for i, node := range nodes {
		initialValidators[i] = warpMessage.SubnetConversionValidatorData{
			NodeID:       node.NodeID.Bytes(),
			BLSPublicKey: node.NodePoP.PublicKey,
			Weight:       nodes[i].Weight,
		}
		initialValidatorsABI[i] = nativetokenstakingmanager.InitialValidator{
			NodeID:       node.NodeID.Bytes(),
			BlsPublicKey: node.NodePoP.PublicKey[:],
			Weight:       nodes[i].Weight,
		}
	}

	l1ConversionData := warpMessage.SubnetConversionData{
		SubnetID:       l1Info.L1ID,
		ManagerChainID: l1Info.BlockchainID,
		ManagerAddress: validatorManagerAddress[:],
		Validators:     initialValidators,
	}
	l1ConversionDataABI := nativetokenstakingmanager.ConversionData{
		L1ID:                         l1Info.L1ID,
		ValidatorManagerBlockchainID: l1Info.BlockchainID,
		ValidatorManagerAddress:      validatorManagerAddress,
		InitialValidators:            initialValidatorsABI,
	}
	l1ConversionID, err := warpMessage.SubnetConversionID(l1ConversionData)
	Expect(err).Should(BeNil())
	l1ConversionSignedMessage := ConstructSubnetConversionMessage(
		l1ConversionID,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)
	// Deliver the Warp message to the l1
	receipt := DeliverNativeTokenSubnetConversion(
		ctx,
		senderKey,
		l1Info,
		validatorManagerAddress,
		l1ConversionSignedMessage,
		l1ConversionDataABI,
	)
	initialValidatorCreatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseInitialValidatorCreated,
	)
	Expect(err).Should(BeNil())
	var validationIDs []ids.ID
	for i := range nodes {
		validationIDs = append(validationIDs, l1Info.L1ID.Append(uint32(i)))
	}

	Expect(initialValidatorCreatedEvent.Weight).Should(Equal(new(big.Int).SetUint64(nodes[0].Weight)))

	emittedValidationID := ids.ID(initialValidatorCreatedEvent.ValidationID)
	Expect(emittedValidationID).Should(Equal(validationIDs[0]))

	return validationIDs
}

func InitializeERC20TokenValidatorSet(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	validatorManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	validatorManagerAddress common.Address,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
	nodes []Node,
) []ids.ID {
	initialValidators := make([]warpMessage.SubnetConversionValidatorData, len(nodes))
	initialValidatorsABI := make([]erc20tokenstakingmanager.InitialValidator, len(nodes))
	for i, node := range nodes {
		initialValidators[i] = warpMessage.SubnetConversionValidatorData{
			NodeID:       node.NodeID.Bytes(),
			BLSPublicKey: node.NodePoP.PublicKey,
			Weight:       nodes[i].Weight,
		}
		initialValidatorsABI[i] = erc20tokenstakingmanager.InitialValidator{
			NodeID:       node.NodeID.Bytes(),
			BlsPublicKey: node.NodePoP.PublicKey[:],
			Weight:       nodes[i].Weight,
		}
	}

	l1ConversionData := warpMessage.SubnetConversionData{
		SubnetID:       l1Info.L1ID,
		ManagerChainID: l1Info.BlockchainID,
		ManagerAddress: validatorManagerAddress[:],
		Validators:     initialValidators,
	}
	l1ConversionDataABI := erc20tokenstakingmanager.ConversionData{
		L1ID:                         l1Info.L1ID,
		ValidatorManagerBlockchainID: l1Info.BlockchainID,
		ValidatorManagerAddress:      validatorManagerAddress,
		InitialValidators:            initialValidatorsABI,
	}
	l1ConversionID, err := warpMessage.SubnetConversionID(l1ConversionData)
	Expect(err).Should(BeNil())
	l1ConversionSignedMessage := ConstructSubnetConversionMessage(
		l1ConversionID,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the l1
	receipt := DeliverERC20TokenSubnetConversion(
		ctx,
		senderKey,
		l1Info,
		validatorManagerAddress,
		l1ConversionSignedMessage,
		l1ConversionDataABI,
	)
	initialValidatorCreatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseInitialValidatorCreated,
	)
	Expect(err).Should(BeNil())
	var validationIDs []ids.ID
	for i := range nodes {
		validationIDs = append(validationIDs, l1Info.L1ID.Append(uint32(i)))
	}

	Expect(initialValidatorCreatedEvent.Weight).Should(Equal(new(big.Int).SetUint64(nodes[0].Weight)))

	emittedValidationID := ids.ID(initialValidatorCreatedEvent.ValidationID)
	Expect(emittedValidationID).Should(Equal(validationIDs[0]))

	return validationIDs
}

func InitializePoAValidatorSet(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	validatorManager *poavalidatormanager.PoAValidatorManager,
	validatorManagerAddress common.Address,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
	nodes []Node,
) []ids.ID {
	initialValidators := make([]warpMessage.SubnetConversionValidatorData, len(nodes))
	initialValidatorsABI := make([]poavalidatormanager.InitialValidator, len(nodes))
	for i, node := range nodes {
		initialValidators[i] = warpMessage.SubnetConversionValidatorData{
			NodeID:       node.NodeID.Bytes(),
			BLSPublicKey: node.NodePoP.PublicKey,
			Weight:       nodes[i].Weight,
		}
		initialValidatorsABI[i] = poavalidatormanager.InitialValidator{
			NodeID:       node.NodeID.Bytes(),
			BlsPublicKey: node.NodePoP.PublicKey[:],
			Weight:       nodes[i].Weight,
		}
	}

	l1ConversionData := warpMessage.SubnetConversionData{
		SubnetID:       l1Info.L1ID,
		ManagerChainID: l1Info.BlockchainID,
		ManagerAddress: validatorManagerAddress[:],
		Validators:     initialValidators,
	}
	l1ConversionDataABI := poavalidatormanager.ConversionData{
		L1ID:                         l1Info.L1ID,
		ValidatorManagerBlockchainID: l1Info.BlockchainID,
		ValidatorManagerAddress:      validatorManagerAddress,
		InitialValidators:            initialValidatorsABI,
	}
	l1ConversionID, err := warpMessage.SubnetConversionID(l1ConversionData)
	Expect(err).Should(BeNil())
	l1ConversionSignedMessage := ConstructSubnetConversionMessage(
		l1ConversionID,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)
	// Deliver the Warp message to the l1
	receipt := DeliverPoASubnetConversion(
		ctx,
		senderKey,
		l1Info,
		validatorManagerAddress,
		l1ConversionSignedMessage,
		l1ConversionDataABI,
	)
	initialValidatorCreatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseInitialValidatorCreated,
	)
	Expect(err).Should(BeNil())
	var validationIDs []ids.ID
	for i := range nodes {
		validationIDs = append(validationIDs, l1Info.L1ID.Append(uint32(i)))
	}

	Expect(initialValidatorCreatedEvent.Weight).Should(Equal(new(big.Int).SetUint64(nodes[0].Weight)))

	emittedValidationID := ids.ID(initialValidatorCreatedEvent.ValidationID)
	Expect(emittedValidationID).Should(Equal(validationIDs[0]))

	return validationIDs
}

func DeliverNativeTokenSubnetConversion(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	validatorManagerAddress common.Address,
	l1ConversionSignedMessage *avalancheWarp.Message,
	l1ConversionData nativetokenstakingmanager.ConversionData,
) *types.Receipt {
	abi, err := nativetokenstakingmanager.NativeTokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("initializeValidatorSet", l1ConversionData, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1,
		validatorManagerAddress,
		l1ConversionSignedMessage,
	)
}

func DeliverERC20TokenSubnetConversion(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	validatorManagerAddress common.Address,
	l1ConversionSignedMessage *avalancheWarp.Message,
	l1ConversionData erc20tokenstakingmanager.ConversionData,
) *types.Receipt {
	abi, err := erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("initializeValidatorSet", l1ConversionData, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1,
		validatorManagerAddress,
		l1ConversionSignedMessage,
	)
}

func DeliverPoASubnetConversion(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	validatorManagerAddress common.Address,
	l1ConversionSignedMessage *avalancheWarp.Message,
	l1ConversionData poavalidatormanager.ConversionData,
) *types.Receipt {
	abi, err := poavalidatormanager.PoAValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("initializeValidatorSet", l1ConversionData, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1,
		validatorManagerAddress,
		l1ConversionSignedMessage,
	)
}

//
// Function call utils
//

func InitializeNativeValidatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	stakeAmount *big.Int,
	node Node,
	expiry uint64,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = stakeAmount

	tx, err := stakingManager.InitializeValidatorRegistration(
		opts,
		nativetokenstakingmanager.ValidatorRegistrationInput{
			NodeID:             node.NodeID[:],
			RegistrationExpiry: expiry,
			BlsPublicKey:       node.NodePoP.PublicKey[:],
		},
		DefaultMinDelegateFeeBips,
		DefaultMinStakeDurationSeconds,
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(ctx, l1, tx.Hash())
	registrationInitiatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidationPeriodCreated,
	)
	Expect(err).Should(BeNil())
	return receipt, ids.ID(registrationInitiatedEvent.ValidationID)
}

func InitializeERC20ValidatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	stakeAmount *big.Int,
	token *exampleerc20.ExampleERC20,
	stakingManagerAddress common.Address,
	node Node,
	expiry uint64,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
) (*types.Receipt, ids.ID) {
	ERC20Approve(
		ctx,
		token,
		stakingManagerAddress,
		stakeAmount,
		l1,
		senderKey,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := stakingManager.InitializeValidatorRegistration(
		opts,
		erc20tokenstakingmanager.ValidatorRegistrationInput{
			NodeID:             node.NodeID[:],
			RegistrationExpiry: expiry,
			BlsPublicKey:       node.NodePoP.PublicKey[:],
		},
		DefaultMinDelegateFeeBips,
		DefaultMinStakeDurationSeconds,
		stakeAmount,
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(ctx, l1, tx.Hash())
	registrationInitiatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidationPeriodCreated,
	)
	Expect(err).Should(BeNil())
	return receipt, ids.ID(registrationInitiatedEvent.ValidationID)
}

func InitializePoAValidatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	node Node,
	expiry uint64,
	validatorManager *poavalidatormanager.PoAValidatorManager,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := validatorManager.InitializeValidatorRegistration(
		opts,
		poavalidatormanager.ValidatorRegistrationInput{
			NodeID:             node.NodeID[:],
			RegistrationExpiry: expiry,
			BlsPublicKey:       node.NodePoP.PublicKey[:],
		},
		node.Weight,
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(ctx, l1, tx.Hash())
	registrationInitiatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseValidationPeriodCreated,
	)
	Expect(err).Should(BeNil())
	return receipt, ids.ID(registrationInitiatedEvent.ValidationID)
}

func CompleteNativeValidatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	stakingManagerContractAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := nativetokenstakingmanager.NativeTokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeValidatorRegistration", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1,
		stakingManagerContractAddress,
		registrationSignedMessage,
	)
}

func CompleteERC20ValidatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	stakingManagerContractAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeValidatorRegistration", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1,
		stakingManagerContractAddress,
		registrationSignedMessage,
	)
}

func CompletePoAValidatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	validatorManagerAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := poavalidatormanager.PoAValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeValidatorRegistration", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1,
		validatorManagerAddress,
		registrationSignedMessage,
	)
}

// Calls a method that retreived a signed Warp message from the transaction's access list
func CallWarpReceiver(
	ctx context.Context,
	callData []byte,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	contract common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, l1, PrivateKeyToAddress(senderKey))
	registrationTx := predicateutils.NewPredicateTx(
		l1.EVMChainID,
		nonce,
		&contract,
		2_000_000,
		gasFeeCap,
		gasTipCap,
		big.NewInt(0),
		callData,
		types.AccessList{},
		warp.ContractAddress,
		signedMessage.Bytes(),
	)
	signedRegistrationTx := SignTransaction(registrationTx, senderKey, l1.EVMChainID)
	return SendTransactionAndWaitForSuccess(ctx, l1, signedRegistrationTx)
}

func InitializeAndCompleteNativeValidatorRegistration(
	ctx context.Context,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	stakingManagerContractAddress common.Address,
	expiry uint64,
	node Node,
	pchainWallet pwallet.Wallet,
	networkID uint32,
) ids.ID {
	stakeAmount, err := stakingManager.WeightToValue(
		&bind.CallOpts{},
		node.Weight,
	)
	Expect(err).Should(BeNil())
	// Initiate validator registration
	receipt, validationID := InitializeNativeValidatorRegistration(
		ctx,
		fundedKey,
		l1Info,
		stakeAmount,
		node,
		expiry,
		stakingManager,
	)

	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
	// (Sending to the P-Chain will be skipped for now)
	signedWarpMessage := ConstructSignedWarpMessage(ctx, receipt, l1Info, pChainInfo)

	_, err = pchainWallet.IssueRegisterSubnetValidatorTx(
		100*units.Avax,
		node.NodePoP.ProofOfPossession,
		signedWarpMessage.Bytes(),
	)
	Expect(err).Should(BeNil())
	PChainProposerVMWorkaround(pchainWallet)
	AdvanceProposerVM(ctx, l1Info, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing validator registration")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		expiry,
		node,
		true,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the l1
	receipt = CompleteNativeValidatorRegistration(
		ctx,
		fundedKey,
		l1Info,
		stakingManagerContractAddress,
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

func InitializeAndCompleteERC20ValidatorRegistration(
	ctx context.Context,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	stakingManagerAddress common.Address,
	erc20 *exampleerc20.ExampleERC20,
	expiry uint64,
	node Node,
	pchainWallet pwallet.Wallet,
	networkID uint32,
) ids.ID {
	stakeAmount, err := stakingManager.WeightToValue(
		&bind.CallOpts{},
		node.Weight,
	)
	Expect(err).Should(BeNil())
	// Initiate validator registration
	var receipt *types.Receipt
	log.Println("Initializing validator registration")
	receipt, validationID := InitializeERC20ValidatorRegistration(
		ctx,
		fundedKey,
		l1Info,
		stakeAmount,
		erc20,
		stakingManagerAddress,
		node,
		expiry,
		stakingManager,
	)

	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
	signedWarpMessage := ConstructSignedWarpMessage(ctx, receipt, l1Info, pChainInfo)

	_, err = pchainWallet.IssueRegisterSubnetValidatorTx(
		100*units.Avax,
		node.NodePoP.ProofOfPossession,
		signedWarpMessage.Bytes(),
	)
	Expect(err).Should(BeNil())
	PChainProposerVMWorkaround(pchainWallet)
	AdvanceProposerVM(ctx, l1Info, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing validator registration")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		expiry,
		node,
		true,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the l1
	receipt = CompleteERC20ValidatorRegistration(
		ctx,
		fundedKey,
		l1Info,
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

func InitializeAndCompletePoAValidatorRegistration(
	ctx context.Context,
	signatureAggregator *aggregator.SignatureAggregator,
	ownerKey *ecdsa.PrivateKey,
	fundedKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	validatorManager *poavalidatormanager.PoAValidatorManager,
	validatorManagerAddress common.Address,
	expiry uint64,
	node Node,
	pchainWallet pwallet.Wallet,
	networkID uint32,
) ids.ID {
	// Initiate validator registration
	receipt, validationID := InitializePoAValidatorRegistration(
		ctx,
		ownerKey,
		l1Info,
		node,
		expiry,
		validatorManager,
	)

	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
	signedWarpMessage := ConstructSignedWarpMessage(ctx, receipt, l1Info, pChainInfo)

	_, err := pchainWallet.IssueRegisterSubnetValidatorTx(
		100*units.Avax,
		node.NodePoP.ProofOfPossession,
		signedWarpMessage.Bytes(),
	)
	Expect(err).Should(BeNil())
	PChainProposerVMWorkaround(pchainWallet)
	AdvanceProposerVM(ctx, l1Info, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing validator registration")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		expiry,
		node,
		true,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the l1
	receipt = CompletePoAValidatorRegistration(
		ctx,
		fundedKey,
		l1Info,
		validatorManagerAddress,
		registrationSignedMessage,
	)
	// Check that the validator is registered in the staking contract
	registrationEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseValidationPeriodRegistered,
	)
	Expect(err).Should(BeNil())
	Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))

	return validationID
}

func InitializeEndNativeValidation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.InitializeEndValidation(
		opts,
		validationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, l1, tx.Hash())
}

func ForceInitializeEndNativeValidation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.ForceInitializeEndValidation(
		opts,
		validationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, l1, tx.Hash())
}

func InitializeEndERC20Validation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	validationID ids.ID,
	force bool,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.InitializeEndValidation(
		opts,
		validationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, l1, tx.Hash())
}

func ForceInitializeEndERC20Validation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.ForceInitializeEndValidation(
		opts,
		validationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, l1, tx.Hash())
}

func InitializeEndPoAValidation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	validatorManager *poavalidatormanager.PoAValidatorManager,
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := validatorManager.InitializeEndValidation(
		opts,
		validationID,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, l1, tx.Hash())
}

func CompleteEndNativeValidation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	stakingManagerContractAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := nativetokenstakingmanager.NativeTokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndValidation", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1,
		stakingManagerContractAddress,
		registrationSignedMessage,
	)
}

func CompleteEndERC20Validation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	stakingManagerContractAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndValidation", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1,
		stakingManagerContractAddress,
		registrationSignedMessage,
	)
}

func CompleteEndPoAValidation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	validatorManagerAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := poavalidatormanager.PoAValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndValidation", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1,
		validatorManagerAddress,
		registrationSignedMessage,
	)
}

func InitializeERC20DelegatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	validationID ids.ID,
	delegationAmount *big.Int,
	token *exampleerc20.ExampleERC20,
	stakingManagerAddress common.Address,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
) *types.Receipt {
	ERC20Approve(
		ctx,
		token,
		stakingManagerAddress,
		delegationAmount,
		l1,
		senderKey,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := stakingManager.InitializeDelegatorRegistration(
		opts,
		validationID,
		delegationAmount,
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(ctx, l1, tx.Hash())
	_, err = GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseDelegatorAdded,
	)
	Expect(err).Should(BeNil())
	return receipt
}

func CompleteERC20DelegatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	delegationID ids.ID,
	l1 interfaces.L1TestInfo,
	stakingManagerContractAddress common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeDelegatorRegistration", delegationID, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1,
		stakingManagerContractAddress,
		signedMessage,
	)
}

func InitializeEndERC20Delegation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	delegationID ids.ID,
) *types.Receipt {
	WaitMinStakeDuration(ctx, l1, senderKey)
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.ForceInitializeEndDelegation(
		opts,
		delegationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, l1, tx.Hash())
}

func CompleteEndERC20Delegation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	delegationID ids.ID,
	l1 interfaces.L1TestInfo,
	stakingManagerContractAddress common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndDelegation", delegationID, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1,
		stakingManagerContractAddress,
		signedMessage,
	)
}

func InitializeNativeDelegatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	validationID ids.ID,
	delegationAmount *big.Int,
	stakingManagerAddress common.Address,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = delegationAmount

	tx, err := stakingManager.InitializeDelegatorRegistration(
		opts,
		validationID,
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(ctx, l1, tx.Hash())
	_, err = GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseDelegatorAdded,
	)
	Expect(err).Should(BeNil())
	return receipt
}

func CompleteNativeDelegatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	delegationID ids.ID,
	l1 interfaces.L1TestInfo,
	stakingManagerContractAddress common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := nativetokenstakingmanager.NativeTokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeDelegatorRegistration", delegationID, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1,
		stakingManagerContractAddress,
		signedMessage,
	)
}

func InitializeEndNativeDelegation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	l1 interfaces.L1TestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	delegationID ids.ID,
) *types.Receipt {
	WaitMinStakeDuration(ctx, l1, senderKey)
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, l1.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.ForceInitializeEndDelegation(
		opts,
		delegationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, l1, tx.Hash())
}

func CompleteEndNativeDelegation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	delegationID ids.ID,
	l1 interfaces.L1TestInfo,
	stakingManagerContractAddress common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := nativetokenstakingmanager.NativeTokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndDelegation", delegationID, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		l1,
		stakingManagerContractAddress,
		signedMessage,
	)
}

func InitializeAndCompleteEndInitialNativeValidation(
	ctx context.Context,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	index uint32,
	weight uint64,
	pchainWallet pwallet.Wallet,
	networkID uint32,
) {
	log.Println("Initializing initial validator removal")
	WaitMinStakeDuration(ctx, l1Info, fundedKey)
	receipt := ForceInitializeEndNativeValidation(
		ctx,
		fundedKey,
		l1Info,
		stakingManager,
		validationID,
	)
	validatorRemovalEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidatorRemovalInitialized,
	)
	Expect(err).Should(BeNil())
	Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
	Expect(validatorRemovalEvent.Weight.Uint64()).Should(Equal(weight))

	// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
	// (Sending to the P-Chain will be skipped for now)
	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, l1Info)
	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedMessage,
		nil,
		l1Info.L1ID,
		67,
	)
	Expect(err).Should(BeNil())

	// Deliver the Warp message to the P-Chain
	pchainWallet.IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
	PChainProposerVMWorkaround(pchainWallet)
	AdvanceProposerVM(ctx, l1Info, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing initial validator removal")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessageForInitialValidator(
		validationID,
		index,
		false,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the l1
	receipt = CompleteEndNativeValidation(
		ctx,
		fundedKey,
		l1Info,
		stakingManagerAddress,
		registrationSignedMessage,
	)

	// Check that the validator is has been delisted from the staking contract
	validationEndedEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidationPeriodEnded,
	)
	Expect(err).Should(BeNil())
	Expect(validationEndedEvent.ValidationID[:]).Should(Equal(validationID[:]))
}

func InitializeAndCompleteEndNativeValidation(
	ctx context.Context,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	expiry uint64,
	node Node,
	nonce uint64,
	pchainWallet pwallet.Wallet,
	networkID uint32,
) {
	log.Println("Initializing validator removal")
	WaitMinStakeDuration(ctx, l1Info, fundedKey)
	receipt := ForceInitializeEndNativeValidation(
		ctx,
		fundedKey,
		l1Info,
		stakingManager,
		validationID,
	)
	validatorRemovalEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidatorRemovalInitialized,
	)
	Expect(err).Should(BeNil())
	Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
	Expect(validatorRemovalEvent.Weight.Uint64()).Should(Equal(node.Weight))

	// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, l1Info)
	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedMessage,
		nil,
		l1Info.L1ID,
		67,
	)
	Expect(err).Should(BeNil())

	// Deliver the Warp message to the P-Chain
	pchainWallet.IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
	PChainProposerVMWorkaround(pchainWallet)
	AdvanceProposerVM(ctx, l1Info, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing validator removal")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		expiry,
		node,
		false,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the l1
	receipt = CompleteEndNativeValidation(
		ctx,
		fundedKey,
		l1Info,
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

func InitializeAndCompleteEndInitialERC20Validation(
	ctx context.Context,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	index uint32,
	weight uint64,
	pchainWallet pwallet.Wallet,
	networkID uint32,
) {
	log.Println("Initializing initial validator removal")
	WaitMinStakeDuration(ctx, l1Info, fundedKey)
	receipt := ForceInitializeEndERC20Validation(
		ctx,
		fundedKey,
		l1Info,
		stakingManager,
		validationID,
	)
	validatorRemovalEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidatorRemovalInitialized,
	)
	Expect(err).Should(BeNil())
	Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
	Expect(validatorRemovalEvent.Weight.Uint64()).Should(Equal(weight))

	// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
	// (Sending to the P-Chain will be skipped for now)
	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, l1Info)
	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedMessage,
		nil,
		l1Info.L1ID,
		67,
	)
	Expect(err).Should(BeNil())

	// Deliver the Warp message to the P-Chain
	pchainWallet.IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
	PChainProposerVMWorkaround(pchainWallet)
	AdvanceProposerVM(ctx, l1Info, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing initial validator removal")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessageForInitialValidator(
		validationID,
		index,
		false,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the l1
	receipt = CompleteEndERC20Validation(
		ctx,
		fundedKey,
		l1Info,
		stakingManagerAddress,
		registrationSignedMessage,
	)

	// Check that the validator is has been delisted from the staking contract
	validationEndedEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidationPeriodEnded,
	)
	Expect(err).Should(BeNil())
	Expect(validationEndedEvent.ValidationID[:]).Should(Equal(validationID[:]))
}

func InitializeAndCompleteEndERC20Validation(
	ctx context.Context,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	expiry uint64,
	node Node,
	nonce uint64,
	pchainWallet pwallet.Wallet,
	networkID uint32,
) {
	log.Println("Initializing validator removal")
	WaitMinStakeDuration(ctx, l1Info, fundedKey)
	receipt := ForceInitializeEndERC20Validation(
		ctx,
		fundedKey,
		l1Info,
		stakingManager,
		validationID,
	)
	validatorRemovalEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidatorRemovalInitialized,
	)
	Expect(err).Should(BeNil())
	Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
	Expect(validatorRemovalEvent.Weight.Uint64()).Should(Equal(node.Weight))

	// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, l1Info)
	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedMessage,
		nil,
		l1Info.L1ID,
		67,
	)
	Expect(err).Should(BeNil())

	// Deliver the Warp message to the P-Chain
	pchainWallet.IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
	PChainProposerVMWorkaround(pchainWallet)
	AdvanceProposerVM(ctx, l1Info, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing validator removal")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		expiry,
		node,
		false,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the l1
	receipt = CompleteEndERC20Validation(
		ctx,
		fundedKey,
		l1Info,
		stakingManagerAddress,
		registrationSignedMessage,
	)

	// Check that the validator is has been delisted from the staking contract
	validationEndedEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidationPeriodEnded,
	)
	Expect(err).Should(BeNil())
	Expect(validationEndedEvent.ValidationID[:]).Should(Equal(validationID[:]))
}

func InitializeAndCompleteEndInitialPoAValidation(
	ctx context.Context,
	signatureAggregator *aggregator.SignatureAggregator,
	ownerKey *ecdsa.PrivateKey,
	fundedKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	stakingManager *poavalidatormanager.PoAValidatorManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	index uint32,
	weight uint64,
	pchainWallet pwallet.Wallet,
	networkID uint32,
) {
	log.Println("Initializing initial validator removal")
	WaitMinStakeDuration(ctx, l1Info, fundedKey)
	receipt := InitializeEndPoAValidation(
		ctx,
		ownerKey,
		l1Info,
		stakingManager,
		validationID,
	)
	validatorRemovalEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidatorRemovalInitialized,
	)
	Expect(err).Should(BeNil())
	Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
	Expect(validatorRemovalEvent.Weight.Uint64()).Should(Equal(weight))

	// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
	// (Sending to the P-Chain will be skipped for now)
	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, l1Info)
	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedMessage,
		nil,
		l1Info.L1ID,
		67,
	)
	Expect(err).Should(BeNil())

	// Deliver the Warp message to the P-Chain
	pchainWallet.IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
	PChainProposerVMWorkaround(pchainWallet)
	AdvanceProposerVM(ctx, l1Info, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing initial validator removal")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessageForInitialValidator(
		validationID,
		index,
		false,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the l1
	receipt = CompleteEndPoAValidation(
		ctx,
		fundedKey,
		l1Info,
		stakingManagerAddress,
		registrationSignedMessage,
	)

	// Check that the validator is has been delisted from the staking contract
	validationEndedEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidationPeriodEnded,
	)
	Expect(err).Should(BeNil())
	Expect(validationEndedEvent.ValidationID[:]).Should(Equal(validationID[:]))
}

func InitializeAndCompleteEndPoAValidation(
	ctx context.Context,
	signatureAggregator *aggregator.SignatureAggregator,
	ownerKey *ecdsa.PrivateKey,
	fundedKey *ecdsa.PrivateKey,
	l1Info interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	validatorManager *poavalidatormanager.PoAValidatorManager,
	validatorManagerAddress common.Address,
	validationID ids.ID,
	weight uint64,
	nonce uint64,
	networkID uint32,
) {
	receipt := InitializeEndPoAValidation(
		ctx,
		ownerKey,
		l1Info,
		validatorManager,
		validationID,
	)
	validatorRemovalEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseValidatorRemovalInitialized,
	)
	Expect(err).Should(BeNil())
	Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
	Expect(validatorRemovalEvent.Weight.Uint64()).Should(Equal(weight))

	// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
	// (Sending to the P-Chain will be skipped for now)
	signedWarpMessage := ConstructSignedWarpMessage(ctx, receipt, l1Info, pChainInfo)
	Expect(err).Should(BeNil())

	// Validate the Warp message, (this will be done on the P-Chain in the future)
	ValidateSubnetValidatorWeightMessage(signedWarpMessage, validationID, 0, nonce)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		0,
		Node{},
		false,
		l1Info,
		pChainInfo,
		networkID,
		signatureAggregator,
	)

	// Deliver the Warp message to the l1
	receipt = CompleteEndPoAValidation(
		ctx,
		ownerKey,
		l1Info,
		validatorManagerAddress,
		registrationSignedMessage,
	)

	// Check that the validator is has been delisted from the staking contract
	registrationEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseValidationPeriodEnded,
	)
	Expect(err).Should(BeNil())
	Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
}

//
// P-Chain utils
//

func ConstructSubnetValidatorRegistrationMessageForInitialValidator(
	validationID ids.ID,
	index uint32,
	valid bool,
	l1 interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
) *avalancheWarp.Message {
	justification := platformvm.SubnetValidatorRegistrationJustification{
		Preimage: &platformvm.SubnetValidatorRegistrationJustification_ConvertSubnetTxData{
			ConvertSubnetTxData: &platformvm.SubnetIDIndex{
				SubnetId: l1.L1ID[:],
				Index:    index,
			},
		},
	}
	justificationBytes, err := proto.Marshal(&justification)
	Expect(err).Should(BeNil())

	registrationPayload, err := warpMessage.NewSubnetValidatorRegistration(validationID, valid)
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
		l1.L1ID,
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
	l1 interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
) *avalancheWarp.Message {
	msg, err := warpMessage.NewRegisterSubnetValidator(
		l1.L1ID,
		node.NodeID,
		node.NodePoP.PublicKey,
		expiry,
		warpMessage.PChainOwner{},
		warpMessage.PChainOwner{},
		node.Weight,
	)
	Expect(err).Should(BeNil())
	justification := platformvm.SubnetValidatorRegistrationJustification{
		Preimage: &platformvm.SubnetValidatorRegistrationJustification_RegisterSubnetValidatorMessage{
			RegisterSubnetValidatorMessage: msg.Bytes(),
		},
	}
	justificationBytes, err := proto.Marshal(&justification)
	Expect(err).Should(BeNil())

	registrationPayload, err := warpMessage.NewSubnetValidatorRegistration(validationID, valid)
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
		l1.L1ID,
		67,
	)
	Expect(err).Should(BeNil())

	return registrationSignedMessage
}

func ConstructSubnetValidatorWeightUpdateMessage(
	validationID ids.ID,
	nonce uint64,
	weight uint64,
	l1 interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	signatureAggregator *aggregator.SignatureAggregator,
	networkID uint32,
) *avalancheWarp.Message {
	payload, err := warpMessage.NewSubnetValidatorWeight(validationID, nonce, weight)
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
		l1.L1ID,
		67,
	)
	Expect(err).Should(BeNil())
	return updateSignedMessage
}

func ConstructSubnetConversionMessage(
	l1ConversionID ids.ID,
	l1 interfaces.L1TestInfo,
	pChainInfo interfaces.L1TestInfo,
	networkID uint32,
	signatureAggregator *aggregator.SignatureAggregator,
) *avalancheWarp.Message {
	l1ConversionPayload, err := warpMessage.NewSubnetConversion(l1ConversionID)
	Expect(err).Should(BeNil())
	l1ConversionAddressedCall, err := warpPayload.NewAddressedCall(
		nil,
		l1ConversionPayload.Bytes(),
	)
	Expect(err).Should(BeNil())
	l1ConversionUnsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		pChainInfo.BlockchainID,
		l1ConversionAddressedCall.Bytes(),
	)
	Expect(err).Should(BeNil())

	l1ConversionSignedMessage, err := signatureAggregator.CreateSignedMessage(
		l1ConversionUnsignedMessage,
		l1.L1ID[:],
		l1.L1ID,
		67,
	)
	Expect(err).Should(BeNil())
	return l1ConversionSignedMessage
}

//
// Warp message validiation utils
// These will be replaced by the actual implementation on the P-Chain in the future
//

func ValidateRegisterSubnetValidatorMessage(
	signedWarpMessage *avalancheWarp.Message,
	nodeID ids.ID,
	weight uint64,
	l1ID ids.ID,
	blsPublicKey [bls.PublicKeyLen]byte,
) {
	// Validate the Warp message, (this will be done on the P-Chain in the future)
	msg, err := warpPayload.ParseAddressedCall(signedWarpMessage.UnsignedMessage.Payload)
	Expect(err).Should(BeNil())
	// Check that the addressed call payload is a registered Warp message type
	var payloadInterface warpMessage.Payload
	ver, err := warpMessage.Codec.Unmarshal(msg.Payload, &payloadInterface)
	Expect(err).Should(BeNil())
	payload, ok := payloadInterface.(*warpMessage.RegisterSubnetValidator)
	Expect(ok).Should(BeTrue())

	Expect(ver).Should(Equal(uint16(warpMessage.CodecVersion)))
	Expect(payload.NodeID).Should(Equal(nodeID))
	Expect(payload.Weight).Should(Equal(weight))
	Expect(payload.SubnetID).Should(Equal(l1ID))
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
	payload, ok := payloadInterface.(*warpMessage.SubnetValidatorWeight)
	Expect(ok).Should(BeTrue())

	Expect(ver).Should(Equal(uint16(warpMessage.CodecVersion)))
	Expect(payload.ValidationID).Should(Equal(validationID))
	Expect(payload.Weight).Should(Equal(weight))
	Expect(payload.Nonce).Should(Equal(nonce))
}

func WaitMinStakeDuration(
	ctx context.Context,
	l1 interfaces.L1TestInfo,
	fundedKey *ecdsa.PrivateKey,
) {
	// Make sure minimum stake duration has passed
	time.Sleep(time.Duration(DefaultMinStakeDurationSeconds) * time.Second)

	// Send a loopback transaction to self to force a block production
	// before delisting the validator.
	SendNativeTransfer(
		ctx,
		l1,
		fundedKey,
		common.Address{},
		big.NewInt(10),
	)
}

func CalculateSubnetConversionValidationId(l1ID ids.ID, validatorIdx uint32) ids.ID {
	preImage := make([]byte, 36)
	copy(preImage[0:32], l1ID[:])
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

	l1ID := v.FieldByName("SubnetID").Interface().([32]byte)
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
	copy(b[2:34], l1ID[:])
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
	l1 interfaces.L1TestInfo,
	fundedKey *ecdsa.PrivateKey,
	blocks int,
) {
	for i := 0; i < blocks; i++ {
		err := subnetEvmUtils.IssueTxsToActivateProposerVMFork(
			ctx, l1.EVMChainID, fundedKey, l1.WSClient,
		)
		Expect(err).Should(BeNil())
	}
}

func ConvertSubnet(
	ctx context.Context,
	l1Info interfaces.L1TestInfo,
	pchainWallet pwallet.Wallet,
	stakingManagerAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
) []Node {
	// Remove the current validators before converting the l1
	var nodes []Node
	for _, uri := range l1Info.NodeURIs {
		infoClient := info.NewClient(uri)
		nodeID, nodePoP, err := infoClient.GetNodeID(ctx)
		Expect(err).Should(BeNil())
		nodes = append(nodes, Node{
			NodeID:  nodeID,
			NodePoP: nodePoP,
		})

		_, err = pchainWallet.IssueRemoveSubnetValidatorTx(
			nodeID,
			l1Info.L1ID,
		)
		Expect(err).Should(BeNil())
	}

	// Sort the nodeIDs so that the l1 conversion ID matches the P-Chain
	sort.Slice(nodes, func(i, j int) bool {
		return string(nodes[i].NodeID.Bytes()) < string(nodes[j].NodeID.Bytes())
	})

	totalWeight := uint64(len(nodes)-1) * units.Schmeckle
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Weight = units.Schmeckle
		totalWeight += units.Schmeckle
	}
	// Set the last node's weight such that removing any other node will not violate the churn limit
	nodes[len(nodes)-1].Weight = 4 * totalWeight

	// Construct the convert l1 info
	destAddr, err := address.ParseToID(DefaultPChainAddress)
	Expect(err).Should(BeNil())
	vdrs := make([]*txs.ConvertSubnetValidator, len(nodes))
	for i, node := range nodes {
		vdrs[i] = &txs.ConvertSubnetValidator{
			NodeID:  node.NodeID.Bytes(),
			Weight:  nodes[i].Weight,
			Balance: units.Avax * 100,
			Signer:  *node.NodePoP,
			RemainingBalanceOwner: warpMessage.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
			DeactivationOwner: warpMessage.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
		}
	}

	log.Println("Issuing ConvertSubnetTx")
	_, err = pchainWallet.IssueConvertSubnetTx(
		l1Info.L1ID,
		l1Info.BlockchainID,
		stakingManagerAddress[:],
		vdrs,
	)
	Expect(err).Should(BeNil())

	PChainProposerVMWorkaround(pchainWallet)
	AdvanceProposerVM(ctx, l1Info, fundedKey, 5)

	return nodes
}
