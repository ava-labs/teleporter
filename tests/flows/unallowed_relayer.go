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

func UnallowedRelayer(network interfaces.Network) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := utils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	// The Teleporter message includes an allowed relayer list that does NOT include the relayer
	//
	ctx := context.Background()

	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: subnetBInfo.BlockchainID,
		DestinationAddress:      fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: fundedAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit: big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{
			common.HexToAddress("0x0123456789012345678901234567890123456789"),
		},
		Message: []byte{1, 2, 3, 4},
	}

	log.Info(
		"Sending Teleporter transaction on source chain",
		"destinationBlockchainID", subnetBInfo.BlockchainID,
	)
	receipt, teleporterMessageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, subnetAInfo, subnetBInfo, sendCrossChainMessageInput, fundedKey,
	)

	//
	// Relay the message to the destination
	//
	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, false)

	//
	// Check Teleporter message was not received on the destination
	//
	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())
}
