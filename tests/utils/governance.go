package utils

import (
	"context"
	"crypto/ecdsa"

	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	validatorsetsig "github.com/ava-labs/teleporter/abi-bindings/go/governance/ValidatorSetSig"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func DeployValidatorSetSig(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	contractSubnet interfaces.SubnetTestInfo,
	validatorSubnet interfaces.SubnetTestInfo,
) (common.Address, *validatorsetsig.ValidatorSetSig) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, contractSubnet.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, validatorSetSig, err := validatorsetsig.DeployValidatorSetSig(
		opts,
		contractSubnet.RPCClient,
		validatorSubnet.BlockchainID,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, contractSubnet, tx.Hash())

	return address, validatorSetSig
}

// Returns Receipt for the transaction unlike TeleporterRegistry version since this is a non-teleporter case
// and we don't want to add the ValidatorSetSig ABI to the subnetInfo
func ExecuteValidatorSetSigCallAndVerify(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
	validatorSetSigAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	expectSuccess bool,
) *types.Receipt {
	signedWarpMsg := GetSignedMessage(ctx, source, destination, unsignedMessage.ID())
	log.Info("Got signed warp message", "messageID", signedWarpMsg.ID())

	signedPredicateTx := CreateExecuteCallPredicateTransaction(
		ctx,
		signedWarpMsg,
		validatorSetSigAddress,
		senderKey,
		destination,
	)

	// Wait for tx to be accepted and verify events emitted
	if expectSuccess {
		return SendTransactionAndWaitForSuccess(ctx, destination, signedPredicateTx)
	}
	return SendTransactionAndWaitForFailure(ctx, destination, signedPredicateTx)
}

func InitOffChainMessageChainConfigValidatorSetSig(
	networkID uint32,
	subnet interfaces.SubnetTestInfo,
	validatorSetSigAddress common.Address,
	validatorSetSigMessages []validatorsetsig.ValidatorSetSigMessage,
) ([]avalancheWarp.UnsignedMessage, string) {
	unsignedMessages := []avalancheWarp.UnsignedMessage{}
	for _, message := range validatorSetSigMessages {
		unsignedMessage := CreateOffChainValidatorSetSigMessage(networkID, subnet, message)
		unsignedMessages = append(unsignedMessages, *unsignedMessage)
		log.Info("Adding validatorSetSig off-chain message to Warp chain config",
			"messageID", unsignedMessage.ID(),
			"blockchainID", subnet.BlockchainID.String())
	}
	return unsignedMessages, GetChainConfigWithOffChainMessages(unsignedMessages)
}

// Creates an off-chain Warp message pointing to a function, contract and payload to be executed
// if the validator set signs this message
func CreateOffChainValidatorSetSigMessage(
	networkID uint32,
	subnet interfaces.SubnetTestInfo,
	message validatorsetsig.ValidatorSetSigMessage,
) *avalancheWarp.UnsignedMessage {
	sourceAddress := []byte{}
	payloadBytes, err := message.Pack()
	Expect(err).Should(BeNil())

	addressedPayload, err := payload.NewAddressedCall(sourceAddress, payloadBytes)
	Expect(err).Should(BeNil())

	unsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		subnet.BlockchainID,
		addressedPayload.Bytes(),
	)
	Expect(err).Should(BeNil())

	return unsignedMessage
}
