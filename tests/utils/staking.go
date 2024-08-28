package utils

import (
	"context"
	"crypto/ecdsa"
	"math/big"
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
	erc20tokenstakingmanager "github.com/ava-labs/teleporter/abi-bindings/go/staking/ERC20TokenStakingManager"
	nativetokenstakingmanager "github.com/ava-labs/teleporter/abi-bindings/go/staking/NativeTokenStakingManager"
	poavalidatormanager "github.com/ava-labs/teleporter/abi-bindings/go/staking/PoAValidatorManager"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/onsi/gomega"
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
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.Initialize(
		opts,
		nativetokenstakingmanager.PoSValidatorManagerSettings{
			BaseSettings: nativetokenstakingmanager.ValidatorManagerSettings{
				PChainBlockchainID: pChainInfo.BlockchainID,
				SubnetID:           subnet.SubnetID,
				MaximumHourlyChurn: 0,
			},
			MinimumStakeAmount:   big.NewInt(0).SetUint64(1e6),
			MaximumStakeAmount:   big.NewInt(0).SetUint64(10e6),
			MinimumStakeDuration: uint64(24 * time.Hour),
			RewardCalculator:     common.Address{},
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
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.Initialize(
		opts,
		erc20tokenstakingmanager.PoSValidatorManagerSettings{
			BaseSettings: erc20tokenstakingmanager.ValidatorManagerSettings{
				PChainBlockchainID: pChainInfo.BlockchainID,
				SubnetID:           subnet.SubnetID,
				MaximumHourlyChurn: 0,
			},
			MinimumStakeAmount:   big.NewInt(0).SetUint64(1e6),
			MaximumStakeAmount:   big.NewInt(0).SetUint64(10e6),
			MinimumStakeDuration: uint64(24 * time.Hour),
			RewardCalculator:     common.Address{},
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
			PChainBlockchainID: pChainInfo.BlockchainID,
			SubnetID:           subnet.SubnetID,
			MaximumHourlyChurn: 0,
		},
		ownerAddress,
	)
	Expect(err).Should(BeNil())
	WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())

	return validatorManagerAddress, validatorManager
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
		nodeID,
		uint64(time.Now().Add(24*time.Hour).Unix()),
		blsPublicKey[:],
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
		stakeAmount,
		nodeID,
		uint64(time.Now().Add(24*time.Hour).Unix()),
		blsPublicKey[:],
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
		weight,
		nodeID,
		uint64(time.Now().Add(24*time.Hour).Unix()),
		blsPublicKey[:],
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

func InitializeEndERC20Validation(
	sendingKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	stakingManager *erc20tokenstakingmanager.ERC20TokenStakingManager,
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
	callData, err := abi.Pack("completeEndValidation", uint32(0), false)
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
	callData, err := abi.Pack("completeEndValidation", uint32(0), false)
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
	callData, err := abi.Pack("completeEndValidation", uint32(0), false)
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
	delegator common.Address,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeDelegatorRegistration", uint32(0), delegator)
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
	validationID ids.ID,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(sendingKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := stakingManager.InitializeEndDelegation(
		opts,
		validationID,
	)
	Expect(err).Should(BeNil())
	return WaitForTransactionSuccess(context.Background(), subnet, tx.Hash())
}

func CompleteEndERC20Delegation(
	sendingKey *ecdsa.PrivateKey,
	delegator common.Address,
	subnet interfaces.SubnetTestInfo,
	stakingManagerContractAddress common.Address,
	signedMessage *avalancheWarp.Message,
) *types.Receipt {
	abi, err := erc20tokenstakingmanager.ERC20TokenStakingManagerMetaData.GetAbi()
	Expect(err).Should(BeNil())
	callData, err := abi.Pack("completeEndDelegation", uint32(0), delegator)
	Expect(err).Should(BeNil())
	return CallWarpReceiver(
		callData,
		sendingKey,
		subnet,
		stakingManagerContractAddress,
		signedMessage,
	)
}

func InitializeNativeDelegatorRegistration() *types.Receipt {
	return nil
}

func CompleteNativeDelegatorRegistration() *types.Receipt {
	return nil
}

func InitializeEndNativeDelegation() *types.Receipt {
	return nil
}

func CompleteEndNativeDelegation() *types.Receipt {
	return nil
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
