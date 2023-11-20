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

func UnallowedRelayerGinkgo() {
	UnallowedRelayer(&network.LocalNetwork{})
}

func UnallowedRelayer(network network.Network) {
	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	// The Teleporter message includes an allowed relayer list that does NOT include the relayer
	//
	ctx := context.Background()

	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationChainID: subnetBInfo.BlockchainID,
		DestinationAddress: fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			ContractAddress: fundedAddress,
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
		"destinationChainID", subnetBInfo.BlockchainID,
	)
	receipt, teleporterMessageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, subnetAInfo, subnetBInfo, sendCrossChainMessageInput, fundedKey, subnetAInfo.TeleporterMessenger,
	)

	//
	// Relay the message to the destination
	//

	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, false, false)

	//
	// Check Teleporter message was not received on the destination
	//
	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, subnetAInfo.BlockchainID, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())
}
