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
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, subnetCInfo := network.GetTwoSubnets()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Get the expected teleporter message ID for Subnet C
	//
	expectedAtoCMessageID, err := teleporter.TeleporterMessenger(subnetAInfo).GetNextMessageID(
		&bind.CallOpts{},
		subnetCInfo.BlockchainID,
	)
	Expect(err).Should(BeNil())

	//
	// Submit a message to be sent from SubnetA to SubnetB
	//
	ctx := context.Background()
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: subnetBInfo.BlockchainID, // Message intended for SubnetB
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
		"destinationBlockchainID", subnetBInfo.BlockchainID,
	)

	receipt, _ := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		teleporter.TeleporterMessenger(subnetAInfo),
		subnetAInfo,
		subnetBInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	teleporter.RelayTeleporterMessage(ctx, receipt, subnetAInfo, subnetCInfo, false, fundedKey, nil, network.GetSignatureAggregator())

	//
	// Check that the message was not received on the Subnet C
	//
	delivered, err := teleporter.TeleporterMessenger(subnetCInfo).MessageReceived(
		&bind.CallOpts{}, expectedAtoCMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())
}
