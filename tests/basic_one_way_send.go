package tests

import (
	"context"
	"math/big"

	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

// Ginkgo describe node that acts as a container for the teleporter e2e tests. This test suite
// will run through the following steps in order:
// 1. Send a transaction to the Teleporter contract on Subnet A
// 2. Aggregate signatures and send the Warp message to Subnet B
// 3. Verify receipt of the message on Subnet B
func BasicOneWaySend() {
	var (
		teleporterMessageID *big.Int
	)

	subnetAInfo := utils.GetSubnetATestInfo()
	subnetBInfo := utils.GetSubnetBTestInfo()
	teleporterContractAddress := utils.GetTeleporterContractAddress()
	fundedAddress, fundedKey := utils.GetFundedAccountInfo()

	//
	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	//
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
	signedTx := utils.CreateSendCrossChainMessageTransaction(ctx, subnetAInfo, sendCrossChainMessageInput, fundedKey, teleporterContractAddress)

	log.Info("Sending Teleporter transaction on source chain", "destinationChainID", subnetBInfo.BlockchainID, "txHash", signedTx.Hash())
	receipt := utils.SendAndWaitForTransaction(ctx, signedTx, subnetAInfo.WSClient)

	teleporterMessengerA, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetAInfo.WSClient)
	Expect(err).Should(BeNil())
	event, err := utils.GetEventFromLogs(receipt.Logs, teleporterMessengerA.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationChainID[:]).Should(Equal(subnetBInfo.BlockchainID[:]))

	teleporterMessageID = event.Message.MessageID

	//
	// Relay the message to the destination
	//

	// This is a nonsense message, so we don't check for successful execution of the teleporter message.
	utils.RelayMessage(ctx, receipt.BlockHash, receipt.BlockNumber, subnetAInfo, subnetBInfo)

	//
	// Check Teleporter message received on the destination
	//

	teleporterMessengerB, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetBInfo.WSClient)
	Expect(err).Should(BeNil())
	success, err := teleporterMessengerB.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, teleporterMessageID)
	Expect(err).Should(BeNil())
	Expect(success).Should(BeTrue())
}
