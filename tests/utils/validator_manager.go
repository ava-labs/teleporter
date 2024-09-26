package utils

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"
	"reflect"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpMessages "github.com/ava-labs/avalanchego/vms/platformvm/warp/messages"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/awm-relayer/signature-aggregator/aggregator"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/mocks/ExampleERC20"
	erc20tokenstakingmanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/ERC20TokenStakingManager"
	examplerewardcalculator "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/ExampleRewardCalculator"
	nativetokenstakingmanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/NativeTokenStakingManager"
	poavalidatormanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/PoAValidatorManager"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/onsi/gomega"
)

const (
	DefaultMinDelegateFeeBips      uint16 = 1
	DefaultMinStakeDurationSeconds uint64 = 1
	DefaultMinStakeAmount          uint64 = 1e18
	DefaultMaxStakeAmount          uint64 = 10e18
	DefaultMaxStakeMultiplier      uint8  = 4
	DefaultMaxChurnPercentage      uint8  = 20
	DefaultChurnPeriodSeconds      uint64 = 1

	initialValidatorPackedLen = 88
)

//
// Deployment utils
//

func DeployNativeTokenStakingManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
) (common.Address, *nativetokenstakingmanager.NativeTokenStakingManager) {
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
			RewardCalculator:         rewardCalculatorAddress,
		},
	)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())

	return stakingManagerContractAddress, stakingManager
}

func DeployERC20TokenStakingManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
) (common.Address, *erc20tokenstakingmanager.ERC20TokenStakingManager) {
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
			RewardCalculator:         rewardCalculatorAddress,
		},
		erc20Address,
	)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())

	return stakingManagerContractAddress, stakingManager, erc20Address, erc20
}

func DeployPoAValidatorManager(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
) (common.Address, *poavalidatormanager.PoAValidatorManager) {
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
	WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())

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
	sendingKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	validatorManager *nativetokenstakingmanager.NativeTokenStakingManager,
	validatorManagerAddress common.Address,
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	initialValidatorWeight uint64,
) ids.ID {
	convertSubnetTxId := ids.GenerateTestID()
	blsPublicKey := [bls.PublicKeyLen]byte{}
	subnetConversionData := nativetokenstakingmanager.SubnetConversionData{
		ConvertSubnetTxID:            convertSubnetTxId,
		ValidatorManagerBlockchainID: subnetInfo.BlockchainID,
		ValidatorManagerAddress:      validatorManagerAddress,
		InitialValidators: []nativetokenstakingmanager.InitialValidator{
			{
				NodeID:       ids.GenerateTestID(),
				Weight:       initialValidatorWeight,
				BlsPublicKey: blsPublicKey[:],
			},
		},
	}
	subnetConversionDataBytes, err := PackSubnetConversionData(subnetConversionData)
	Expect(err).Should(BeNil())
	subnetConversionID := sha256.Sum256(subnetConversionDataBytes)
	subnetConversionSignedMessage := ConstructSubnetConversionMessage(
		subnetConversionID,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)
	// Deliver the Warp message to the subnet
	receipt := DeliverNativeTokenSubnetConversion(
		sendingKey,
		subnetInfo,
		validatorManagerAddress,
		subnetConversionSignedMessage,
		subnetConversionData,
	)
	initialValidatorCreatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseInitialValidatorCreated,
	)
	Expect(err).Should(BeNil())
	Expect(initialValidatorCreatedEvent.NodeID).Should(Equal(subnetConversionData.InitialValidators[0].NodeID))
	Expect(initialValidatorCreatedEvent.Weight).Should(Equal(new(big.Int).SetUint64(initialValidatorWeight)))

	expectedValidationID := CalculateSubnetConversionValidationId(convertSubnetTxId, 0)
	emittedValidationID := ids.ID(initialValidatorCreatedEvent.ValidationID)
	Expect(emittedValidationID).Should(Equal(expectedValidationID))

	return emittedValidationID
}

func InitializeERC20TokenValidatorSet(
	ctx context.Context,
	sendingKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	validatorManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	validatorManagerAddress common.Address,
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	initialValidatorWeight uint64,
) ids.ID {
	convertSubnetTxId := ids.GenerateTestID()
	blsPublicKey := [bls.PublicKeyLen]byte{}
	subnetConversionData := erc20tokenstakingmanager.SubnetConversionData{
		ConvertSubnetTxID:            convertSubnetTxId,
		ValidatorManagerBlockchainID: subnetInfo.BlockchainID,
		ValidatorManagerAddress:      validatorManagerAddress,
		InitialValidators: []erc20tokenstakingmanager.InitialValidator{
			{
				NodeID:       ids.GenerateTestID(),
				Weight:       initialValidatorWeight,
				BlsPublicKey: blsPublicKey[:],
			},
		},
	}
	subnetConversionDataBytes, err := PackSubnetConversionData(subnetConversionData)
	Expect(err).Should(BeNil())
	subnetConversionID := sha256.Sum256(subnetConversionDataBytes)
	subnetConversionSignedMessage := ConstructSubnetConversionMessage(
		subnetConversionID,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)
	// Deliver the Warp message to the subnet
	receipt := DeliverERC20TokenSubnetConversion(
		sendingKey,
		subnetInfo,
		validatorManagerAddress,
		subnetConversionSignedMessage,
		subnetConversionData,
	)
	initialValidatorCreatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseInitialValidatorCreated,
	)
	Expect(err).Should(BeNil())
	Expect(initialValidatorCreatedEvent.NodeID).Should(Equal(subnetConversionData.InitialValidators[0].NodeID))
	Expect(initialValidatorCreatedEvent.Weight).Should(Equal(new(big.Int).SetUint64(initialValidatorWeight)))

	expectedValidationID := CalculateSubnetConversionValidationId(convertSubnetTxId, 0)
	emittedValidationID := ids.ID(initialValidatorCreatedEvent.ValidationID)
	Expect(emittedValidationID).Should(Equal(expectedValidationID))

	return emittedValidationID
}

func InitializePoAValidatorSet(
	ctx context.Context,
	sendingKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	validatorManager *poavalidatormanager.PoAValidatorManager,
	validatorManagerAddress common.Address,
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	initialValidatorWeight uint64,
) ids.ID {
	convertSubnetTxId := ids.GenerateTestID()
	blsPublicKey := [bls.PublicKeyLen]byte{}
	subnetConversionData := poavalidatormanager.SubnetConversionData{
		ConvertSubnetTxID:            convertSubnetTxId,
		ValidatorManagerBlockchainID: subnetInfo.BlockchainID,
		ValidatorManagerAddress:      validatorManagerAddress,
		InitialValidators: []poavalidatormanager.InitialValidator{
			{
				NodeID:       ids.GenerateTestID(),
				Weight:       initialValidatorWeight,
				BlsPublicKey: blsPublicKey[:],
			},
		},
	}

	subnetConversionDataBytes, err := PackSubnetConversionData(subnetConversionData)
	Expect(err).Should(BeNil())
	subnetConversionID := sha256.Sum256(subnetConversionDataBytes)
	subnetConversionSignedMessage := ConstructSubnetConversionMessage(
		subnetConversionID,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)
	// Deliver the Warp message to the subnet
	receipt := DeliverPoASubnetConversion(
		sendingKey,
		subnetInfo,
		validatorManagerAddress,
		subnetConversionSignedMessage,
		subnetConversionData,
	)
	initialValidatorCreatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseInitialValidatorCreated,
	)
	Expect(err).Should(BeNil())
	Expect(initialValidatorCreatedEvent.NodeID).Should(Equal(subnetConversionData.InitialValidators[0].NodeID))
	Expect(initialValidatorCreatedEvent.Weight).Should(Equal(new(big.Int).SetUint64(initialValidatorWeight)))

	expectedValidationID := CalculateSubnetConversionValidationId(convertSubnetTxId, 0)
	emittedValidationID := ids.ID(initialValidatorCreatedEvent.ValidationID)
	Expect(emittedValidationID).Should(Equal(expectedValidationID))

	return emittedValidationID
}

func DeliverNativeTokenSubnetConversion(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	validatorManagerAddress common.Address,
	subnetConversionSignedMessage *avalancheWarp.Message,
	subnetConversionData nativetokenstakingmanager.SubnetConversionData,
) *types.Receipt {
	abi, err := nativetokenstakingmanager.NativeTokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("initializeValidatorSet", subnetConversionData, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		validatorManagerAddress,
		subnetConversionSignedMessage,
	)
}

func DeliverERC20TokenSubnetConversion(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	validatorManagerAddress common.Address,
	subnetConversionSignedMessage *avalancheWarp.Message,
	subnetConversionData erc20tokenstakingmanager.SubnetConversionData,
) *types.Receipt {
	abi, err := erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("initializeValidatorSet", subnetConversionData, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		validatorManagerAddress,
		subnetConversionSignedMessage,
	)
}

func DeliverPoASubnetConversion(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	validatorManagerAddress common.Address,
	subnetConversionSignedMessage *avalancheWarp.Message,
	subnetConversionData poavalidatormanager.SubnetConversionData,
) *types.Receipt {
	abi, err := poavalidatormanager.PoAValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("initializeValidatorSet", subnetConversionData, uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		validatorManagerAddress,
		subnetConversionSignedMessage,
	)
}

//
// Function call utils
//

func InitializeNativeValidatorRegistration(
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakeAmount *big.Int,
	nodeID ids.ID,
	blsPublicKey [bls.PublicKeyLen]byte,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = stakeAmount

	tx, err := stakingManager.InitializeValidatorRegistration(
		opts,
		nativetokenstakingmanager.ValidatorRegistrationInput{
			NodeID:             nodeID,
			RegistrationExpiry: uint64(time.Now().Add(24 * time.Hour).Unix()),
			BlsPublicKey:       blsPublicKey[:],
		},
		DefaultMinDelegateFeeBips,
		DefaultMinStakeDurationSeconds,
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())
	registrationInitiatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidationPeriodCreated,
	)
	Expect(err).Should(BeNil())
	return receipt, ids.ID(registrationInitiatedEvent.ValidationID)
}

func InitializeERC20ValidatorRegistration(
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakeAmount *big.Int,
	token *exampleerc20.ExampleERC20,
	stakingManagerAddress common.Address,
	nodeID ids.ID,
	blsPublicKey [bls.PublicKeyLen]byte,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
) (*types.Receipt, ids.ID) {
	ERC20Approve(
		context.Background(),
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
			NodeID:             nodeID,
			RegistrationExpiry: uint64(time.Now().Add(24 * time.Hour).Unix()),
			BlsPublicKey:       blsPublicKey[:],
		},
		DefaultMinDelegateFeeBips,
		DefaultMinStakeDurationSeconds,
		stakeAmount,
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())
	registrationInitiatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseValidationPeriodCreated,
	)
	Expect(err).Should(BeNil())
	return receipt, ids.ID(registrationInitiatedEvent.ValidationID)
}

func InitializePoAValidatorRegistration(
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	weight uint64,
	nodeID ids.ID,
	blsPublicKey [bls.PublicKeyLen]byte,
	validatorManager *poavalidatormanager.PoAValidatorManager,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := validatorManager.InitializeValidatorRegistration(
		opts,
		poavalidatormanager.ValidatorRegistrationInput{
			NodeID:             nodeID,
			RegistrationExpiry: uint64(time.Now().Add(24 * time.Hour).Unix()),
			BlsPublicKey:       blsPublicKey[:],
		},
		weight,
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())
	registrationInitiatedEvent, err := GetEventFromLogs(
		receipt.Logs,
		validatorManager.ParseValidationPeriodCreated,
	)
	Expect(err).Should(BeNil())
	return receipt, ids.ID(registrationInitiatedEvent.ValidationID)
}

func CompleteNativeValidatorRegistration(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := nativetokenstakingmanager.NativeTokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeValidatorRegistration", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		stakingManagerContractAddress,
		registrationSignedMessage,
	)
}

func CompleteERC20ValidatorRegistration(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeValidatorRegistration", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		stakingManagerContractAddress,
		registrationSignedMessage,
	)
}

func CompletePoAValidatorRegistration(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	validatorManagerAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := poavalidatormanager.PoAValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeValidatorRegistration", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		validatorManagerAddress,
		registrationSignedMessage,
	)
}

// Calls a method that retreived a signed Warp message from the transaction's access list
func CallWarpReceiver(
	callData []byte,
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	contract common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	gasFeeCap, gasTipCap, nonce := CalculateTxParams(context.Background(), subnet, PrivateKeyToAddress(sendingKey))
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
	signedRegistrationTx := SignTransaction(registrationTx, sendingKey, subnet.EVMChainID)
	return SendTransactionAndWaitForSuccess(context.Background(), subnet, signedRegistrationTx)
}

func InitializeAndCompleteNativeValidatorRegistration(
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	stakingManagerContractAddress common.Address,
	stakeAmount *big.Int,
) ids.ID {
	// Initiate validator registration
	nodeID := ids.GenerateTestID()
	blsPublicKey := [bls.PublicKeyLen]byte{}
	receipt, validationID := InitializeNativeValidatorRegistration(
		fundedKey,
		subnetInfo,
		stakeAmount,
		nodeID,
		blsPublicKey,
		stakingManager,
	)

	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
	// (Sending to the P-Chain will be skipped for now)
	signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetInfo, pChainInfo)

	weight, err := stakingManager.ValueToWeight(
		&bind.CallOpts{},
		stakeAmount,
	)
	Expect(err).Should(BeNil())
	// Validate the Warp message, (this will be done on the P-Chain in the future)
	ValidateRegisterSubnetValidatorMessage(
		signedWarpMessage,
		nodeID,
		weight,
		subnetInfo.SubnetID,
		blsPublicKey,
	)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		true,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteNativeValidatorRegistration(
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
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	stakingManagerAddress common.Address,
	erc20 *exampleerc20.ExampleERC20,
	stakeAmount *big.Int,
) ids.ID {
	// Initiate validator registration
	nodeID := ids.GenerateTestID()
	blsPublicKey := [bls.PublicKeyLen]byte{}
	var receipt *types.Receipt
	receipt, validationID := InitializeERC20ValidatorRegistration(
		fundedKey,
		subnetInfo,
		stakeAmount,
		erc20,
		stakingManagerAddress,
		nodeID,
		blsPublicKey,
		stakingManager,
	)

	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
	// (Sending to the P-Chain will be skipped for now)
	signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetInfo, pChainInfo)

	weight, err := stakingManager.ValueToWeight(
		&bind.CallOpts{},
		stakeAmount,
	)
	Expect(err).Should(BeNil())
	// Validate the Warp message, (this will be done on the P-Chain in the future)
	ValidateRegisterSubnetValidatorMessage(
		signedWarpMessage,
		nodeID,
		weight,
		subnetInfo.SubnetID,
		blsPublicKey,
	)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		true,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteERC20ValidatorRegistration(
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
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	ownerKey *ecdsa.PrivateKey,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	validatorManager *poavalidatormanager.PoAValidatorManager,
	validatorManagerAddress common.Address,
	weight uint64,
) ids.ID {
	// Initiate validator registration
	nodeID := ids.GenerateTestID()
	blsPublicKey := [bls.PublicKeyLen]byte{}
	receipt, validationID := InitializePoAValidatorRegistration(
		ownerKey,
		subnetInfo,
		weight,
		nodeID,
		blsPublicKey,
		validatorManager,
	)

	// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
	// (Sending to the P-Chain will be skipped for now)
	signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetInfo, pChainInfo)

	// Validate the Warp message, (this will be done on the P-Chain in the future)
	ValidateRegisterSubnetValidatorMessage(
		signedWarpMessage,
		nodeID,
		weight,
		subnetInfo.SubnetID,
		blsPublicKey,
	)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		true,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompletePoAValidatorRegistration(
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
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(sendingKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.InitializeEndValidation(
		opts,
		validationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())
}

func ForceInitializeEndNativeValidation(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(sendingKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.ForceInitializeEndValidation(
		opts,
		validationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())
}

func InitializeEndERC20Validation(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	validationID ids.ID,
	force bool,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(sendingKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.InitializeEndValidation(
		opts,
		validationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())
}

func ForceInitializeEndERC20Validation(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(sendingKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.ForceInitializeEndValidation(
		opts,
		validationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())
}

func InitializeEndPoAValidation(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	validatorManager *poavalidatormanager.PoAValidatorManager,
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(sendingKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := validatorManager.InitializeEndValidation(
		opts,
		validationID,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())
}

func CompleteEndNativeValidation(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := nativetokenstakingmanager.NativeTokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndValidation", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		stakingManagerContractAddress,
		registrationSignedMessage,
	)
}

func CompleteEndERC20Validation(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndValidation", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		stakingManagerContractAddress,
		registrationSignedMessage,
	)
}

func CompleteEndPoAValidation(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	validatorManagerAddress common.Address,
	registrationSignedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := poavalidatormanager.PoAValidatorManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndValidation", uint32(0))
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		validatorManagerAddress,
		registrationSignedMessage,
	)
}

func InitializeERC20DelegatorRegistration(
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	validationID ids.ID,
	delegationAmount *big.Int,
	token *exampleerc20.ExampleERC20,
	stakingManagerAddress common.Address,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
) *types.Receipt {
	ERC20Approve(
		context.Background(),
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
	receipt := WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())
	_, err = GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseDelegatorAdded,
	)
	Expect(err).Should(BeNil())
	return receipt
}

func CompleteERC20DelegatorRegistration(
	sendingKey *ecdsa.PrivateKey,
	delegationID ids.ID,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeDelegatorRegistration", uint32(0), delegationID)
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		stakingManagerContractAddress,
		signedMessage,
	)
}

func InitializeEndERC20Delegation(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	delegationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(sendingKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.InitializeEndDelegation(
		opts,
		delegationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())
}

func CompleteEndERC20Delegation(
	sendingKey *ecdsa.PrivateKey,
	delegationID ids.ID,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndDelegation", uint32(0), delegationID)
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		stakingManagerContractAddress,
		signedMessage,
	)
}

func InitializeNativeDelegatorRegistration(senderKey *ecdsa.PrivateKey,
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
	receipt := WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())
	_, err = GetEventFromLogs(
		receipt.Logs,
		stakingManager.ParseDelegatorAdded,
	)
	Expect(err).Should(BeNil())
	return receipt
}

func CompleteNativeDelegatorRegistration(
	sendingKey *ecdsa.PrivateKey,
	delegationID ids.ID,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := nativetokenstakingmanager.NativeTokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeDelegatorRegistration", uint32(0), delegationID)
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		stakingManagerContractAddress,
		signedMessage,
	)
}

func InitializeEndNativeDelegation(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	delegationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(sendingKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.InitializeEndDelegation(
		opts,
		delegationID,
		false,
		0,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())
}

func CompleteEndNativeDelegation(
	sendingKey *ecdsa.PrivateKey,
	delegationID ids.ID,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := nativetokenstakingmanager.NativeTokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndDelegation", uint32(0), delegationID)
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		stakingManagerContractAddress,
		signedMessage,
	)
}

func InitializeAndCompleteEndNativeValidation(
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *nativetokenstakingmanager.NativeTokenStakingManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	weight uint64,
	nonce uint64,
) {
	WaitMinStakeDuration(subnetInfo, fundedKey)
	receipt := ForceInitializeEndNativeValidation(
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
	signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetInfo, pChainInfo)
	Expect(err).Should(BeNil())

	// Validate the Warp message, (this will be done on the P-Chain in the future)
	ValidateSetSubnetValidatorWeightMessage(signedWarpMessage, validationID, 0, nonce)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		false,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteEndNativeValidation(
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

func InitializeAndCompleteEndERC20Validation(
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
	stakingManagerAddress common.Address,
	validationID ids.ID,
	weight uint64,
	nonce uint64,
) {
	WaitMinStakeDuration(subnetInfo, fundedKey)
	receipt := ForceInitializeEndERC20Validation(
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
	signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetInfo, pChainInfo)
	Expect(err).Should(BeNil())

	// Validate the Warp message, (this will be done on the P-Chain in the future)
	ValidateSetSubnetValidatorWeightMessage(signedWarpMessage, validationID, 0, nonce)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		false,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteEndERC20Validation(
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

func InitializeAndCompleteEndPoAValidation(
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
	signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetInfo, pChainInfo)
	Expect(err).Should(BeNil())

	// Validate the Warp message, (this will be done on the P-Chain in the future)
	ValidateSetSubnetValidatorWeightMessage(signedWarpMessage, validationID, 0, nonce)

	// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
	registrationSignedMessage := ConstructSubnetValidatorRegistrationMessage(
		validationID,
		false,
		subnetInfo,
		pChainInfo,
		network,
		signatureAggregator,
	)

	// Deliver the Warp message to the subnet
	receipt = CompleteEndPoAValidation(
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

func ConstructSubnetValidatorRegistrationMessage(
	validationID ids.ID,
	valid bool,
	subnet interfaces.SubnetTestInfo,
	pChainInfo interfaces.SubnetTestInfo,
	network interfaces.LocalNetwork,
	signatureAggregator *aggregator.SignatureAggregator,
) *avalancheWarp.Message {
	registrationPayload, err := warpMessages.NewSubnetValidatorRegistration(validationID, valid)
	Expect(err).Should(BeNil())
	registrationAddressedCall, err := warpPayload.NewAddressedCall(common.Address{}.Bytes(), registrationPayload.Bytes())
	Expect(err).Should(BeNil())
	registrationUnsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		network.GetNetworkID(),
		pChainInfo.BlockchainID,
		registrationAddressedCall.Bytes(),
	)
	Expect(err).Should(BeNil())

	registrationSignedMessage, err := signatureAggregator.CreateSignedMessage(
		registrationUnsignedMessage,
		nil,
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
	updatePayload, err := warpMessages.NewSubnetValidatorWeightUpdate(validationID, nonce, weight)
	Expect(err).Should(BeNil())
	updateAddressedCall, err := warpPayload.NewAddressedCall(common.Address{}.Bytes(), updatePayload.Bytes())
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
	subnetConversionPayload, err := warpMessages.NewSubnetConversion(subnetConversionID)
	Expect(err).Should(BeNil())
	subnetConversionAddressedCall, err := warpPayload.NewAddressedCall(
		common.Address{}.Bytes(),
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
		nil,
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
	var payloadInterface warpMessages.Payload
	ver, err := warpMessages.Codec.Unmarshal(msg.Payload, &payloadInterface)
	Expect(err).Should(BeNil())
	registerValidatorPayload, ok := payloadInterface.(*warpMessages.RegisterSubnetValidator)
	Expect(ok).Should(BeTrue())

	Expect(ver).Should(Equal(uint16(warpMessages.CodecVersion)))
	Expect(registerValidatorPayload.NodeID).Should(Equal(nodeID))
	Expect(registerValidatorPayload.Weight).Should(Equal(weight))
	Expect(registerValidatorPayload.SubnetID).Should(Equal(subnetID))
	Expect(registerValidatorPayload.BlsPubKey[:]).Should(Equal(blsPublicKey[:]))
}

func ValidateSetSubnetValidatorWeightMessage(
	signedWarpMessage *avalancheWarp.Message,
	validationID ids.ID,
	weight uint64,
	nonce uint64,
) {
	msg, err := warpPayload.ParseAddressedCall(signedWarpMessage.UnsignedMessage.Payload)
	Expect(err).Should(BeNil())
	// Check that the addressed call payload is a registered Warp message type
	var payloadInterface warpMessages.Payload
	ver, err := warpMessages.Codec.Unmarshal(msg.Payload, &payloadInterface)
	Expect(err).Should(BeNil())
	registerValidatorPayload, ok := payloadInterface.(*warpMessages.SetSubnetValidatorWeight)
	Expect(ok).Should(BeTrue())

	Expect(ver).Should(Equal(uint16(warpMessages.CodecVersion)))
	Expect(registerValidatorPayload.ValidationID).Should(Equal(validationID))
	Expect(registerValidatorPayload.Weight).Should(Equal(weight))
	Expect(registerValidatorPayload.Nonce).Should(Equal(nonce))
}

func WaitMinStakeDuration(
	subnet interfaces.SubnetTestInfo,
	fundedKey *ecdsa.PrivateKey,
) {
	// Make sure minimum stake duration has passed
	time.Sleep(time.Duration(DefaultMinStakeDurationSeconds) * time.Second)

	// Send a loopback transaction to self to force a block production
	// before delisting the validator.
	SendNativeTransfer(
		context.Background(),
		subnet,
		fundedKey,
		common.Address{},
		big.NewInt(10),
	)
}

func CalculateSubnetConversionValidationId(convertSubnetTxID ids.ID, validatorIdx uint32) ids.ID {
	preImage := make([]byte, 36)
	copy(preImage[0:32], convertSubnetTxID[:])
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
		"ConvertSubnetTxID":            reflect.TypeOf([32]byte{}),
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

	convertSubnetTxID := v.FieldByName("ConvertSubnetTxID").Interface().([32]byte)
	validatorManagerBlockchainID := v.FieldByName("ValidatorManagerBlockchainID").Interface().([32]byte)
	validatorManagerAddress := v.FieldByName("ValidatorManagerAddress").Interface().(common.Address)
	initialValidators := v.FieldByName("InitialValidators")

	packedLen := 92 + initialValidatorPackedLen*initialValidators.Len()
	b := make([]byte, packedLen)
	copy(b[0:32], convertSubnetTxID[:])
	copy(b[32:64], validatorManagerBlockchainID[:])
	// These are evm addresses and have lengths of 20 so hardcoding here
	binary.BigEndian.PutUint32(b[64:68], uint32(20))
	copy(b[68:88], validatorManagerAddress.Bytes())
	binary.BigEndian.PutUint32(b[88:92], uint32(initialValidators.Len()))
	// Pack each InitialValidator struct
	for i := 0; i < initialValidators.Len(); i++ {
		iv := initialValidators.Index(i).Interface()
		ivPacked, err := PackInitialValidator(iv)
		if err != nil {
			return nil, fmt.Errorf("failed to pack InitialValidator: %w", err)
		}
		if len(ivPacked) != initialValidatorPackedLen {
			return nil, fmt.Errorf(
				"expected packed InitialValidator to be %d bytes, got %d",
				initialValidatorPackedLen,
				len(ivPacked),
			)
		}
		copy(b[92+i*initialValidatorPackedLen:92+(i+1)*initialValidatorPackedLen], ivPacked)
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
		"NodeID":       reflect.TypeOf([32]byte{}),
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
	nodeID := v.FieldByName("NodeID").Interface().([32]byte)
	weight := v.FieldByName("Weight").Interface().(uint64)
	blsPublicKey := v.FieldByName("BlsPublicKey").Interface().([]byte)

	b := make([]byte, initialValidatorPackedLen)
	copy(b[0:32], nodeID[:])
	binary.BigEndian.PutUint64(b[32:40], weight)
	copy(b[40:88], blsPublicKey)
	return b, nil
}
