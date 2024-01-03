package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func ResubmitAlteredMessage(network interfaces.Network) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := utils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	ctx := context.Background()

	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: subnetBInfo.BlockchainID,
		DestinationAddress:      fundedAddress,
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
	receipt = network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)

	log.Info("Checking the message was received on the destination")
	delivered, err :=
		subnetBInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Get the Teleporter message from receive event
	event, err := utils.GetEventFromLogs(receipt.Logs, subnetBInfo.TeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.MessageID[:]).Should(Equal(messageID[:]))
	teleporterMessage := event.Message

	// Alter the message
	alteredMessage := make([]byte, len(teleporterMessage.Message))
	copy(alteredMessage, teleporterMessage.Message)
	alteredMessage[0] = ^alteredMessage[0]
	Expect(alteredMessage[:]).ShouldNot(Equal(teleporterMessage.Message[:]))
	teleporterMessage.Message = alteredMessage

	// Resubmit the altered message
	log.Info("Submitting the altered Teleporter message on the source chain")
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err :=
		subnetAInfo.TeleporterMessenger.RetrySendCrossChainMessage(opts, teleporterMessage)
	Expect(err).ShouldNot(BeNil())

	// We expect the tx to be nil because the Warp message failed verification, which happens in the predicate
	// In that case, the block is never built, and the transaction is never mined
	Expect(tx).Should(BeNil())
}
