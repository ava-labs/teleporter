package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func ResubmitAlteredMessageGinkgo() {
	ResubmitAlteredMessage(&network.LocalNetwork{})
}

func ResubmitAlteredMessage(network network.Network) {
	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress, subnetAInfo.ChainRPCClient,
	)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress, subnetBInfo.ChainRPCClient,
	)
	Expect(err).Should(BeNil())

	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	ctx := context.Background()

	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationChainID: subnetBInfo.BlockchainID,
		DestinationAddress: fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			ContractAddress: fundedAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	receipt, messageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, subnetAInfo, subnetBInfo, sendCrossChainMessageInput, fundedAddress, fundedKey, subnetATeleporterMessenger)

	// Relay the message to the destination
	receipt = network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, false, true)

	log.Info("Checking the message was received on the destination")
	delivered, err := subnetBTeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Get the Teleporter message from receive event
	event, err := utils.GetEventFromLogs(receipt.Logs, subnetBTeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.MessageID).Should(Equal(messageID))
	teleporterMessage := event.Message

	// Alter the message
	alteredMessage := make([]byte, len(teleporterMessage.Message))
	copy(alteredMessage, teleporterMessage.Message)
	alteredMessage[0] = ^alteredMessage[0]
	Expect(alteredMessage[:]).ShouldNot(Equal(teleporterMessage.Message[:]))
	teleporterMessage.Message = alteredMessage

	// Resubmit the altered message
	log.Info("Submitting the altered Teleporter message on the destination")
	opts := utils.CreateTransactorOpts(ctx, subnetAInfo, fundedAddress, fundedKey)
	_, err = subnetATeleporterMessenger.RetrySendCrossChainMessage(opts, subnetBInfo.BlockchainID, teleporterMessage)
	Expect(err).ShouldNot(BeNil())
}
