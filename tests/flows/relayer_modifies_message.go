package flows

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	"github.com/ava-labs/subnet-evm/x/warp"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	gasUtils "github.com/ava-labs/teleporter/utils/gas-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

type constructSignedMessageFunc func(
	ctx context.Context,
	sourceReceipt *types.Receipt,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
) []byte

// Disallow this test from being run on anything but a local network, since it requires special behavior by the relayer
func RelayerModifiesMessage(network interfaces.Network, constructSignedMessageFunc constructSignedMessageFunc) {
	subnets := network.GetSubnetsInfo()
	Expect(len(subnets)).Should(BeNumerically(">=", 2))
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	ctx := context.Background()

	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationChainID: subnetBInfo.BlockchainID,
		DestinationAddress: fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: fundedAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	receipt, messageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, subnetAInfo, subnetBInfo, sendCrossChainMessageInput, fundedKey)

	// Relay the message to the destination
	// Relayer modifies the message in flight
	relayAlteredMessage(ctx, receipt, subnetAInfo, subnetBInfo, fundedKey, network.GetTeleporterContractAddress(), constructSignedMessageFunc)

	// Check Teleporter message was not received on the destination
	delivered, err :=
		subnetBInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())
}

func relayAlteredMessage(
	ctx context.Context,
	sourceReceipt *types.Receipt,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
	fundedKey *ecdsa.PrivateKey,
	teleporterContractAddress common.Address,
	constructSignedMessageFunc constructSignedMessageFunc,
) {
	// Fetch the Teleporter message from the logs
	sendEvent, err :=
		utils.GetEventFromLogs(sourceReceipt.Logs, source.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())

	signedWarpMessageBytes := constructSignedMessageFunc(ctx, sourceReceipt, source, destination)

	// Construct the transaction to send the Warp message to the destination chain
	signedTx := createAlteredReceiveCrossChainMessageTransaction(
		ctx,
		signedWarpMessageBytes,
		sendEvent.Message.RequiredGasLimit,
		teleporterContractAddress,
		fundedKey,
		destination,
	)

	log.Info("Sending transaction to destination chain")
	utils.SendTransactionAndWaitForAcceptance(ctx, destination, signedTx, false)
}

func createAlteredReceiveCrossChainMessageTransaction(
	ctx context.Context,
	warpMessageBytes []byte,
	requiredGasLimit *big.Int,
	teleporterContractAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
) *types.Transaction {
	fundedAddress := crypto.PubkeyToAddress(fundedKey.PublicKey)
	// Construct the transaction to send the Warp message to the destination chain
	log.Info("Constructing transaction for the destination chain")
	signedMessage, err := avalancheWarp.ParseMessage(warpMessageBytes)
	Expect(err).Should(BeNil())

	numSigners, err := signedMessage.Signature.NumSigners()
	Expect(err).Should(BeNil())

	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(numSigners, requiredGasLimit)
	Expect(err).Should(BeNil())

	callData, err := teleportermessenger.PackReceiveCrossChainMessage(0, fundedAddress)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := utils.CalculateTxParams(ctx, subnetInfo, fundedAddress)

	alterTeleporterMessage(signedMessage)

	destinationTx := predicateutils.NewPredicateTx(
		subnetInfo.ChainIDInt,
		nonce,
		&teleporterContractAddress,
		gasLimit,
		gasFeeCap,
		gasTipCap,
		big.NewInt(0),
		callData,
		types.AccessList{},
		warp.ContractAddress,
		signedMessage.Bytes(),
	)

	return utils.SignTransaction(destinationTx, fundedKey, subnetInfo.ChainIDInt)
}

func alterTeleporterMessage(signedMessage *avalancheWarp.Message) {
	warpMsgPayload, err := warpPayload.ParseAddressedCall(signedMessage.UnsignedMessage.Payload)
	Expect(err).Should(BeNil())

	teleporterMessage, err := teleportermessenger.UnpackTeleporterMessage(warpMsgPayload.Payload)
	Expect(err).Should(BeNil())
	// Alter the message
	teleporterMessage.Message[0] = ^teleporterMessage.Message[0]

	// Pack the teleporter message
	teleporterMessageBytes, err := teleportermessenger.PackTeleporterMessage(*teleporterMessage)
	Expect(err).Should(BeNil())

	payload, err := warpPayload.NewAddressedCall(warpMsgPayload.SourceAddress, teleporterMessageBytes)
	Expect(err).Should(BeNil())

	signedMessage.UnsignedMessage.Payload = payload.Bytes()

	signedMessage.Initialize()
}
