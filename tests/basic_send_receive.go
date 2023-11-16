package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

func BasicSendReceiveGinkgo() {
	BasicSendReceive(&network.LocalNetwork{})
}

// Tests basic one-way send from Subnet A to Subnet B and vice versa
func BasicSendReceive(network network.Network) {
	var (
		teleporterMessageID *big.Int
	)

	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

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

	receipt, teleporterMessageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		subnetAInfo,
		subnetBInfo,
		sendCrossChainMessageInput,
		fundedKey,
		subnetAInfo.TeleporterMessenger,
	)

	//
	// Relay the message to the destination
	//

	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)

	//
	// Check Teleporter message received on the destination
	//
	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, subnetAInfo.BlockchainID, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Send a transaction to Subnet B to issue a Warp Message from the Teleporter contract to Subnet A
	//
	sendCrossChainMessageInput.DestinationChainID = subnetAInfo.BlockchainID
	receipt, teleporterMessageID = utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		subnetBInfo,
		subnetAInfo,
		sendCrossChainMessageInput,
		fundedKey,
		subnetBInfo.TeleporterMessenger,
	)

	//
	// Relay the message to the destination
	//
	network.RelayMessage(ctx, receipt, subnetBInfo, subnetAInfo, true)

	//
	// Check Teleporter message received on the destination
	//
	delivered, err = subnetAInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, subnetBInfo.BlockchainID, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())
}
