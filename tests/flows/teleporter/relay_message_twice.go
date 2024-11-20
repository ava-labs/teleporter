package teleporter

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/teleporter/TeleporterMessenger"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	. "github.com/onsi/gomega"
)

func RelayMessageTwice(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	L1AInfo := network.GetPrimaryNetworkInfo()
	L1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Send a transaction to L1 A to issue a ICM Message from the Teleporter contract to L1 B
	//
	ctx := context.Background()

	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: L1BInfo.BlockchainID,
		DestinationAddress:      fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: fundedAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	log.Info(
		"Sending Teleporter transaction on source chain",
		"destinationBlockchainID", L1BInfo.BlockchainID,
	)
	receipt, teleporterMessageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, teleporter.TeleporterMessenger(L1AInfo), L1AInfo, L1BInfo, sendCrossChainMessageInput, fundedKey,
	)

	//
	// Relay the message to the destination
	//
	teleporter.RelayTeleporterMessage(ctx, receipt, L1AInfo, L1BInfo, true, fundedKey)

	//
	// Check Teleporter message received on the destination
	//
	log.Info("Checking the message was received on the destination")
	delivered, err := teleporter.TeleporterMessenger(L1BInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Attempt to send the same message again, should fail
	//
	log.Info("Relaying the same Teleporter message again on the destination")
	teleporter.RelayTeleporterMessage(ctx, receipt, L1AInfo, L1BInfo, false, fundedKey)
}
