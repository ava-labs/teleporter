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
	subnet interfaces.SubnetTestInfo,
) (common.Address, *nativetokenstakingmanager.NativeTokenStakingManager) {
	// Reset the global binary data for better test isolation
	nativetokenstakingmanager.NativeTokenStakingManagerBin =
		nativetokenstakingmanager.NativeTokenStakingManagerMetaData.Bin

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, stakingManager, err := nativetokenstakingmanager.DeployNativeTokenStakingManager(
		opts,
		subnet.RPCClient,
		0, // ICMInitializable.Allowed
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, stakingManager
}

func DeployAndInitializeNativeTokenStakingManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
) (common.Address, *nativetokenstakingmanager.NativeTokenStakingManager) {
	stakingManagerContractAddress, stakingManager := DeployNativeTokenStakingManager(
		ctx,
		senderKey,
		subnet,
	)
	rewardCalculatorAddress, _ := DeployExampleRewardCalculator(
		ctx,
		senderKey,
		subnet,
		uint64(10),
	)
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.Initialize(
		opts,
		nativetokenstakingmanager.PoSValidatorManagerSettings{
			BaseSettings: nativetokenstakingmanager.ValidatorManagerSettings{
				SubnetID:               subnet.SubnetID,
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
	WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return stakingManagerContractAddress, stakingManager
}

func DeployERC20TokenStakingManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
) (common.Address, *erc20tokenstakingmanager.ERC20TokenStakingManager) {
	// Reset the global binary data for better test isolation
	erc20tokenstakingmanager.ERC20TokenStakingManagerBin =
		erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.Bin

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, stakingManager, err := erc20tokenstakingmanager.DeployERC20TokenStakingManager(
		opts,
		subnet.RPCClient,
		0, // ICMInitializable.Allowed
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, stakingManager
}

func DeployAndInitializeERC20TokenStakingManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
) (
	common.Address,
	*erc20tokenstakingmanager.ERC20TokenStakingManager,
	common.Address,
	*exampleerc20.ExampleERC20,
) {
	stakingManagerContractAddress, stakingManager := DeployERC20TokenStakingManager(
		ctx,
		senderKey,
		subnet,
	)

	erc20Address, erc20 := DeployExampleERC20(ctx, senderKey, subnet)
	rewardCalculatorAddress, _ := DeployExampleRewardCalculator(
		ctx,
		senderKey,
		subnet,
		uint64(10),
	)
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.Initialize(
		opts,
		erc20tokenstakingmanager.PoSValidatorManagerSettings{
			BaseSettings: erc20tokenstakingmanager.ValidatorManagerSettings{
				SubnetID:               subnet.SubnetID,
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
	WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return stakingManagerContractAddress, stakingManager, erc20Address, erc20
}

func DeployPoAValidatorManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
) (common.Address, *poavalidatormanager.PoAValidatorManager) {
	// Reset the global binary data for better test isolation
	poavalidatormanager.PoAValidatorManagerBin = poavalidatormanager.PoAValidatorManagerMetaData.Bin

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, validatorManager, err := poavalidatormanager.DeployPoAValidatorManager(
		opts,
		subnet.RPCClient,
		0, // ICMInitializable.Allowed
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, validatorManager
}

func DeployAndInitializePoAValidatorManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	ownerAddress common.Address,
) (common.Address, *poavalidatormanager.PoAValidatorManager) {
	validatorManagerAddress, validatorManager := DeployPoAValidatorManager(
		ctx,
		senderKey,
		subnet,
	)
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := validatorManager.Initialize(
		opts,
		poavalidatormanager.ValidatorManagerSettings{
			SubnetID:               subnet.SubnetID,
			ChurnPeriodSeconds:     uint64(0),
			MaximumChurnPercentage: uint8(20),
		},
		ownerAddress,
	)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return validatorManagerAddress, validatorManager
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

func InitializeNativeTokenValidatorSet(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	validatorManager *nativetokenstakingmanager.NativeTokenStakingManager,
	validatorManagerAddress common.Address,
	network interfaces.LocalNetwork,
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

	subnetConversionData := warpMessage.SubnetConversionData{
		SubnetID:       subnetInfo.SubnetID,
		ManagerChainID: subnetInfo.BlockchainID,
		ManagerAddress: validatorManagerAddress[:],
		Validators:     initialValidators,
	}
	subnetConversionDataABI := nativetokenstakingmanager.ConversionData{
		SubnetID:                     subnetInfo.SubnetID,
		ValidatorManagerBlockchainID: subnetInfo.BlockchainID,
		ValidatorManagerAddress:      validatorManagerAddress,
		InitialValidators:            initialValidatorsABI,
	}
	subnetConversionID, err := warpMessage.SubnetConversionID(subnetConversionData)
	Expect(err).Should(BeNil())
	subnetConversionSignedMessage := ConstructSubnetConversionMessage(
		subnetConversionID,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)
	// Deliver the Warp message to the subnet
	receipt := DeliverNativeTokenSubnetConversion(
		ctx,
		senderKey,
		subnetInfo,
		validatorManagerAddress,
		subnetConversionSignedMessage,
		subnetConversionDataABI,
	)
	initialValidatorCreatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseInitialValidatorCreated,
	)
	Expect(err).Should(BeNil())
	var validationIDs []ids.ID
	for i := range nodes {
		validationIDs = append(validationIDs, subnetInfo.SubnetID.Append(uint32(i)))
	}

	Expect(initialValidatorCreatedEvent.Weight).Should(Equal(new(big.Int).SetUint64(nodes[0].Weight)))

	emittedValidationID := ids.ID(initialValidatorCreatedEvent.ValidationID)
	Expect(emittedValidationID).Should(Equal(validationIDs[0]))

	return validationIDs
}

func InitializeERC20TokenValidatorSet(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	validatorManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	validatorManagerAddress common.Address,
	network interfaces.LocalNetwork,
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

	subnetConversionData := warpMessage.SubnetConversionData{
		SubnetID:       subnetInfo.SubnetID,
		ManagerChainID: subnetInfo.BlockchainID,
		ManagerAddress: validatorManagerAddress[:],
		Validators:     initialValidators,
	}
	subnetConversionDataABI := erc20tokenstakingmanager.ConversionData{
		SubnetID:                     subnetInfo.SubnetID,
		ValidatorManagerBlockchainID: subnetInfo.BlockchainID,
		ValidatorManagerAddress:      validatorManagerAddress,
		InitialValidators:            initialValidatorsABI,
	}
	subnetConversionID, err := warpMessage.SubnetConversionID(subnetConversionData)
	Expect(err).Should(BeNil())
	subnetConversionSignedMessage := ConstructSubnetConversionMessage(
		subnetConversionID,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt := DeliverERC20TokenSubnetConversion(
		ctx,
		senderKey,
		subnetInfo,
		validatorManagerAddress,
		subnetConversionSignedMessage,
		subnetConversionDataABI,
	)
	initialValidatorCreatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseInitialValidatorCreated,
	)
	Expect(err).Should(BeNil())
	var validationIDs []ids.ID
	for i := range nodes {
		validationIDs = append(validationIDs, subnetInfo.SubnetID.Append(uint32(i)))
	}

	Expect(initialValidatorCreatedEvent.Weight).Should(Equal(new(big.Int).SetUint64(nodes[0].Weight)))

	emittedValidationID := ids.ID(initialValidatorCreatedEvent.ValidationID)
	Expect(emittedValidationID).Should(Equal(validationIDs[0]))

	return validationIDs
}

func InitializePoAValidatorSet(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	validatorManager *poavalidatormanager.PoAValidatorManager,
	validatorManagerAddress common.Address,
	network interfaces.LocalNetwork,
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

	subnetConversionData := warpMessage.SubnetConversionData{
		SubnetID:       subnetInfo.SubnetID,
		ManagerChainID: subnetInfo.BlockchainID,
		ManagerAddress: validatorManagerAddress[:],
		Validators:     initialValidators,
	}
	subnetConversionDataABI := poavalidatormanager.ConversionData{
		SubnetID:                     subnetInfo.SubnetID,
		ValidatorManagerBlockchainID: subnetInfo.BlockchainID,
		ValidatorManagerAddress:      validatorManagerAddress,
		InitialValidators:            initialValidatorsABI,
	}
	subnetConversionID, err := warpMessage.SubnetConversionID(subnetConversionData)
	Expect(err).Should(BeNil())
	subnetConversionSignedMessage := ConstructSubnetConversionMessage(
		subnetConversionID,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)
	// Deliver the Warp message to the subnet
	receipt := DeliverPoASubnetConversion(
		ctx,
		senderKey,
		subnetInfo,
		validatorManagerAddress,
		subnetConversionSignedMessage,
		subnetConversionDataABI,
	)
	initialValidatorCreatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseInitialValidatorCreated,
	)
	Expect(err).Should(BeNil())
	var validationIDs []ids.ID
	for i := range nodes {
		validationIDs = append(validationIDs, subnetInfo.SubnetID.Append(uint32(i)))
	}

	Expect(initialValidatorCreatedEvent.Weight).Should(Equal(new(big.Int).SetUint64(nodes[0].Weight)))

	emittedValidationID := ids.ID(initialValidatorCreatedEvent.ValidationID)
	Expect(emittedValidationID).Should(Equal(validationIDs[0]))

	return validationIDs
}

func DeliverNativeTokenSubnetConversion(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	validatorManagerAddress common.Address,
	subnetConversionSignedMessage *avalancheWarp.Message,
	subnetConversionData nativetokenstakingmanager.ConversionData,
) *types.Receipt {
	abi, err := nativetokenstakingmanager.NativeTokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("initializeValidatorSet", subnetConversionData, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		subnet,
		validatorManagerAddress,
		subnetConversionSignedMessage,
	)
}

func DeliverERC20TokenSubnetConversion(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	validatorManagerAddress common.Address,
	subnetConversionSignedMessage *avalancheWarp.Message,
	subnetConversionData erc20tokenstakingmanager.ConversionData,
) *types.Receipt {
	abi, err := erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("initializeValidatorSet", subnetConversionData, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		subnet,
		validatorManagerAddress,
		subnetConversionSignedMessage,
	)
}

func DeliverPoASubnetConversion(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	validatorManagerAddress common.Address,
	subnetConversionSignedMessage *avalancheWarp.Message,
	subnetConversionData poavalidatormanager.ConversionData,
) *types.Receipt {
	abi, err := poavalidatormanager.PoAValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("initializeValidatorSet", subnetConversionData, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		ctx,
		callData,
		senderKey,
		subnet,
		validatorManagerAddress,
		subnetConversionSignedMessage,
	)
}

//
// Function call utils
//

func InitializeNativeValidatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakeAmount *big.Int,
	node Node,
	expiry uint64,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
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
	receipt := WaitForTransactionSuccess(ctx, subnet, tx.Hash())
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
	subnet interfaces.SubnetTestInfo,
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
		subnet,
		senderKey,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
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
	receipt := WaitForTransactionSuccess(ctx, subnet, tx.Hash())
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
	subnet interfaces.SubnetTestInfo,
	node Node,
	expiry uint64,
	validatorManager *poavalidatormanager.PoAValidatorManager,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
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
	receipt := WaitForTransactionSuccess(ctx, subnet, tx.Hash())
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
	subnet interfaces.SubnetTestInfo,
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
		subnet,
		stakingManagerContractAddress,
		registrationSignedMessage,
	)
}

func CompleteERC20ValidatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
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
		subnet,
		stakingManagerContractAddress,
		registrationSignedMessage,
	)
}

func CompletePoAValidatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
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
		subnet,
		validatorManagerAddress,
		registrationSignedMessage,
	)
}

// Calls a method that retreived a signed Warp message from the transaction's access list
func CallWarpReceiver(
	ctx context.Context,
	callData []byte,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	contract common.Address,
	signedMessage *avalancheWarp.Message,
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
		signedMessage.Bytes(),
	)
	signedRegistrationTx := SignTransaction(registrationTx, senderKey, subnet.EVMChainID)
	return SendTransactionAndWaitForSuccess(ctx, subnet, signedRegistrationTx)
}

func InitializeAndCompleteNativeValidatorRegistration(
	ctx context.Context,
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	stakingManagerContractAddress common.Address,
	expiry uint64,
	node Node,
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
		subnetInfo,
		stakeAmount,
		node,
		expiry,
		stakingManager,
	)

	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
	// (Sending to the P-Chain will be skipped for now)
	signedWarpMessage := ConstructSignedWarpMessage(ctx, receipt, subnetInfo, pChainInfo)

	_, err = network.GetPChainWallet().IssueRegisterSubnetValidatorTx(
		100*units.Avax,
		node.NodePoP.ProofOfPossession,
		signedWarpMessage.Bytes(),
	)
	Expect(err).Should(BeNil())
	PChainProposerVMWorkaround(network)
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
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteNativeValidatorRegistration(
		ctx,
		fundedKey,
		subnetInfo,
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
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	stakingManagerAddress common.Address,
	erc20 *exampleerc20.ExampleERC20,
	expiry uint64,
	node Node,
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
		subnetInfo,
		stakeAmount,
		erc20,
		stakingManagerAddress,
		node,
		expiry,
		stakingManager,
	)

	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
	signedWarpMessage := ConstructSignedWarpMessage(ctx, receipt, subnetInfo, pChainInfo)

	_, err = network.GetPChainWallet().IssueRegisterSubnetValidatorTx(
		100*units.Avax,
		node.NodePoP.ProofOfPossession,
		signedWarpMessage.Bytes(),
	)
	Expect(err).Should(BeNil())
	PChainProposerVMWorkaround(network)
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
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteERC20ValidatorRegistration(
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

func InitializeAndCompletePoAValidatorRegistration(
	ctx context.Context,
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	ownerKey *ecdsa.PrivateKey,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	validatorManager *poavalidatormanager.PoAValidatorManager,
	validatorManagerAddress common.Address,
	expiry uint64,
	node Node,
) ids.ID {
	// Initiate validator registration
	receipt, validationID := InitializePoAValidatorRegistration(
		ctx,
		ownerKey,
		subnetInfo,
		node,
		expiry,
		validatorManager,
	)

	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
	signedWarpMessage := ConstructSignedWarpMessage(ctx, receipt, subnetInfo, pChainInfo)

	_, err := network.GetPChainWallet().IssueRegisterSubnetValidatorTx(
		100*units.Avax,
		node.NodePoP.ProofOfPossession,
		signedWarpMessage.Bytes(),
	)
	Expect(err).Should(BeNil())
	PChainProposerVMWorkaround(network)
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
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompletePoAValidatorRegistration(
		ctx,
		fundedKey,
		subnetInfo,
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
	subnet interfaces.SubnetTestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.InitializeEndValidation(
		opts,
		validationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, subnet, tx.Hash())
}

func ForceInitializeEndNativeValidation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.ForceInitializeEndValidation(
		opts,
		validationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, subnet, tx.Hash())
}

func InitializeEndERC20Validation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	validationID ids.ID,
	force bool,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.InitializeEndValidation(
		opts,
		validationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, subnet, tx.Hash())
}

func ForceInitializeEndERC20Validation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.ForceInitializeEndValidation(
		opts,
		validationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, subnet, tx.Hash())
}

func InitializeEndPoAValidation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	validatorManager *poavalidatormanager.PoAValidatorManager,
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := validatorManager.InitializeEndValidation(
		opts,
		validationID,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, subnet, tx.Hash())
}

func CompleteEndNativeValidation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
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
		subnet,
		stakingManagerContractAddress,
		registrationSignedMessage,
	)
}

func CompleteEndERC20Validation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
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
		subnet,
		stakingManagerContractAddress,
		registrationSignedMessage,
	)
}

func CompleteEndPoAValidation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
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
		subnet,
		validatorManagerAddress,
		registrationSignedMessage,
	)
}

func InitializeERC20DelegatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
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
		subnet,
		senderKey,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := stakingManager.InitializeDelegatorRegistration(
		opts,
		validationID,
		delegationAmount,
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(ctx, subnet, tx.Hash())
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
	subnet interfaces.SubnetTestInfo,
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
		subnet,
		stakingManagerContractAddress,
		signedMessage,
	)
}

func InitializeEndERC20Delegation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	delegationID ids.ID,
) *types.Receipt {
	WaitMinStakeDuration(ctx, subnet, senderKey)
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.ForceInitializeEndDelegation(
		opts,
		delegationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, subnet, tx.Hash())
}

func CompleteEndERC20Delegation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	delegationID ids.ID,
	subnet interfaces.SubnetTestInfo,
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
		subnet,
		stakingManagerContractAddress,
		signedMessage,
	)
}

func InitializeNativeDelegatorRegistration(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	validationID ids.ID,
	delegationAmount *big.Int,
	stakingManagerAddress common.Address,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = delegationAmount

	tx, err := stakingManager.InitializeDelegatorRegistration(
		opts,
		validationID,
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(ctx, subnet, tx.Hash())
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
	subnet interfaces.SubnetTestInfo,
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
		subnet,
		stakingManagerContractAddress,
		signedMessage,
	)
}

func InitializeEndNativeDelegation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	delegationID ids.ID,
) *types.Receipt {
	WaitMinStakeDuration(ctx, subnet, senderKey)
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.ForceInitializeEndDelegation(
		opts,
		delegationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(ctx, subnet, tx.Hash())
}

func CompleteEndNativeDelegation(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	delegationID ids.ID,
	subnet interfaces.SubnetTestInfo,
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
		subnet,
		stakingManagerContractAddress,
		signedMessage,
	)
}

func InitializeAndCompleteEndInitialNativeValidation(
	ctx context.Context,
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	index uint32,
	weight uint64,
) {
	log.Println("Initializing initial validator removal")
	WaitMinStakeDuration(ctx, subnetInfo, fundedKey)
	receipt := ForceInitializeEndNativeValidation(
		ctx,
		fundedKey,
		subnetInfo,
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
	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, subnetInfo)
	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedMessage,
		nil,
		subnetInfo.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())

	// Deliver the Warp message to the P-Chain
	network.GetPChainWallet().IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
	PChainProposerVMWorkaround(network)
	AdvanceProposerVM(ctx, subnetInfo, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing initial validator removal")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessageForInitialValidator(
		validationID,
		index,
		false,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteEndNativeValidation(
		ctx,
		fundedKey,
		subnetInfo,
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
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	expiry uint64,
	node Node,
	nonce uint64,
) {
	log.Println("Initializing validator removal")
	WaitMinStakeDuration(ctx, subnetInfo, fundedKey)
	receipt := ForceInitializeEndNativeValidation(
		ctx,
		fundedKey,
		subnetInfo,
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
	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, subnetInfo)
	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedMessage,
		nil,
		subnetInfo.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())

	// Deliver the Warp message to the P-Chain
	network.GetPChainWallet().IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
	PChainProposerVMWorkaround(network)
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
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteEndNativeValidation(
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

func InitializeAndCompleteEndInitialERC20Validation(
	ctx context.Context,
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	index uint32,
	weight uint64,
) {
	log.Println("Initializing initial validator removal")
	WaitMinStakeDuration(ctx, subnetInfo, fundedKey)
	receipt := ForceInitializeEndERC20Validation(
		ctx,
		fundedKey,
		subnetInfo,
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
	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, subnetInfo)
	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedMessage,
		nil,
		subnetInfo.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())

	// Deliver the Warp message to the P-Chain
	network.GetPChainWallet().IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
	PChainProposerVMWorkaround(network)
	AdvanceProposerVM(ctx, subnetInfo, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing initial validator removal")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessageForInitialValidator(
		validationID,
		index,
		false,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteEndERC20Validation(
		ctx,
		fundedKey,
		subnetInfo,
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
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	expiry uint64,
	node Node,
	nonce uint64,
) {
	log.Println("Initializing validator removal")
	WaitMinStakeDuration(ctx, subnetInfo, fundedKey)
	receipt := ForceInitializeEndERC20Validation(
		ctx,
		fundedKey,
		subnetInfo,
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
	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, subnetInfo)
	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedMessage,
		nil,
		subnetInfo.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())

	// Deliver the Warp message to the P-Chain
	network.GetPChainWallet().IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
	PChainProposerVMWorkaround(network)
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
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteEndERC20Validation(
		ctx,
		fundedKey,
		subnetInfo,
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
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	ownerKey *ecdsa.PrivateKey,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *poavalidatormanager.PoAValidatorManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	index uint32,
	weight uint64,
) {
	log.Println("Initializing initial validator removal")
	WaitMinStakeDuration(ctx, subnetInfo, fundedKey)
	receipt := InitializeEndPoAValidation(
		ctx,
		ownerKey,
		subnetInfo,
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
	unsignedMessage := ExtractWarpMessageFromLog(ctx, receipt, subnetInfo)
	signedWarpMessage, err := signatureAggregator.CreateSignedMessage(
		unsignedMessage,
		nil,
		subnetInfo.SubnetID,
		67,
	)
	Expect(err).Should(BeNil())

	// Deliver the Warp message to the P-Chain
	network.GetPChainWallet().IssueSetSubnetValidatorWeightTx(signedWarpMessage.Bytes())
	PChainProposerVMWorkaround(network)
	AdvanceProposerVM(ctx, subnetInfo, fundedKey, 5)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	log.Println("Completing initial validator removal")
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessageForInitialValidator(
		validationID,
		index,
		false,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteEndPoAValidation(
		ctx,
		fundedKey,
		subnetInfo,
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
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	ownerKey *ecdsa.PrivateKey,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	validatorManager *poavalidatormanager.PoAValidatorManager,
	validatorManagerAddress common.Address,
	validationID ids.ID,
	weight uint64,
	nonce uint64,
) {
	receipt := InitializeEndPoAValidation(
		ctx,
		ownerKey,
		subnetInfo,
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
	signedWarpMessage := ConstructSignedWarpMessage(ctx, receipt, subnetInfo, pChainInfo)
	Expect(err).Should(BeNil())

	// Validate the Warp message, (this will be done on the P-Chain in the future)
	ValidateSubnetValidatorWeightMessage(signedWarpMessage, validationID, 0, nonce)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		0,
		Node{},
		false,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteEndPoAValidation(
		ctx,
		ownerKey,
		subnetInfo,
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
	subnet interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
) *avalancheWarp.Message {
	justification := platformvm.SubnetValidatorRegistrationJustification{
		Preimage: &platformvm.SubnetValidatorRegistrationJustification_ConvertSubnetTxData{
			ConvertSubnetTxData: &platformvm.SubnetIDIndex{
				SubnetId: subnet.SubnetID[:],
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
		network.GetNetworkID(),
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
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
) *avalancheWarp.Message {
	msg, err := warpMessage.NewRegisterSubnetValidator(
		subnet.SubnetID,
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
		network.GetNetworkID(),
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
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
) *avalancheWarp.Message {
	payload, err := warpMessage.NewSubnetValidatorWeight(validationID, nonce, weight)
	Expect(err).Should(BeNil())
	updateAddressedCall, err := warpPayload.NewAddressedCall(nil, payload.Bytes())
	Expect(err).Should(BeNil())
	updateUnsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		network.GetNetworkID(),
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
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
) *avalancheWarp.Message {
	subnetConversionPayload, err := warpMessage.NewSubnetConversion(subnetConversionID)
	Expect(err).Should(BeNil())
	subnetConversionAddressedCall, err := warpPayload.NewAddressedCall(
		nil,
		subnetConversionPayload.Bytes(),
	)
	Expect(err).Should(BeNil())
	subnetConversionUnsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		network.GetNetworkID(),
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
	payload, ok := payloadInterface.(*warpMessage.RegisterSubnetValidator)
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
	payload, ok := payloadInterface.(*warpMessage.SubnetValidatorWeight)
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
	network interfaces.LocalNetwork,
) {
	// Workaround current block map rules
	destAddr, err := address.ParseToID(DefaultPChainAddress)
	Expect(err).Should(BeNil())
	log.Println("Waiting for P-Chain...")
	time.Sleep(30 * time.Second)

	pBuilder := network.GetPChainWallet().Builder()
	pContext := pBuilder.Context()
	avaxAssetID := pContext.AVAXAssetID
	locktime := uint64(time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC).Unix())
	amount := 500 * units.MilliAvax
	_, err = network.GetPChainWallet().IssueBaseTx([]*avax.TransferableOutput{
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
	_, err = network.GetPChainWallet().IssueBaseTx([]*avax.TransferableOutput{
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
	for i := 0; i < blocks; i++ {
		err := subnetEvmUtils.IssueTxsToActivateProposerVMFork(
			ctx, subnet.EVMChainID, fundedKey, subnet.WSClient,
		)
		Expect(err).Should(BeNil())
	}
}

func ConvertSubnet(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	network interfaces.LocalNetwork,
	stakingManagerAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
) []Node {
	// Remove the current validators before converting the subnet
	var nodes []Node
	for _, uri := range subnetInfo.NodeURIs {
		infoClient := info.NewClient(uri)
		nodeID, nodePoP, err := infoClient.GetNodeID(ctx)
		Expect(err).Should(BeNil())
		nodes = append(nodes, Node{
			NodeID:  nodeID,
			NodePoP: nodePoP,
		})

		_, err = network.GetPChainWallet().IssueRemoveSubnetValidatorTx(
			nodeID,
			subnetInfo.SubnetID,
		)
		Expect(err).Should(BeNil())
	}

	// Sort the nodeIDs so that the subnet conversion ID matches the P-Chain
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

	// Construct the convert subnet info
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
	_, err = network.GetPChainWallet().IssueConvertSubnetTx(
		subnetInfo.SubnetID,
		subnetInfo.BlockchainID,
		stakingManagerAddress[:],
		vdrs,
	)
	Expect(err).Should(BeNil())

	PChainProposerVMWorkaround(network)
	AdvanceProposerVM(ctx, subnetInfo, fundedKey, 5)

	return nodes
}
