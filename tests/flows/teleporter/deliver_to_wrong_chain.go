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

func DeliverToWrongChain(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	L1AInfo := network.GetPrimaryNetworkInfo()
	L1BInfo, L1CInfo := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Get the expected teleporter message ID for L1 C
	//
	expectedAtoCMessageID, err := teleporter.TeleporterMessenger(L1AInfo).GetNextMessageID(
		&bind.CallOpts{},
		L1CInfo.BlockchainID,
	)
	Expect(err).Should(BeNil())

	//
	// Submit a message to be sent from L1A to L1B
	//
	ctx := context.Background()
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: L1BInfo.BlockchainID, // Message intended for L1B
		DestinationAddress:      common.HexToAddress("0x1111111111111111111111111111111111111111"),
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

	receipt, _ := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		teleporter.TeleporterMessenger(L1AInfo),
		L1AInfo,
		L1BInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	teleporter.RelayTeleporterMessage(ctx, receipt, L1AInfo, L1CInfo, false, fundedKey)

	//
	// Check that the message was not received on the L1 C
	//
	delivered, err := teleporter.TeleporterMessenger(L1CInfo).MessageReceived(
		&bind.CallOpts{}, expectedAtoCMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())
}
