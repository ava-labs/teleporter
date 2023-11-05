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

func AddFeeAmountGinkgo() {
	addFeeAmount(&network.LocalNetwork{})
}

func addFeeAmount(network network.Network) {
	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetAInfo.ChainRPCClient)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetBInfo.ChainRPCClient)
	Expect(err).Should(BeNil())

	// Use mock token as the fee token
	mockTokenAddress, mockToken := utils.DeployMockToken(context.Background(), fundedAddress, fundedKey, subnetAInfo)
	utils.ExampleERC20Approve(ctx, mockToken, teleporterContractAddress, big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)), subnetAInfo)

	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationChainID: subnetBInfo.BlockchainID,
		DestinationAddress: fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			ContractAddress: mockTokenAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	sendCrossChainMsgReceipt, messageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, subnetAInfo, subnetBInfo, sendCrossChainMessageInput, fundedAddress, fundedKey, subnetATeleporterMessenger)

	// add fee amount
	utils.SendAddFeeAmountAndWaitForAcceptance(
		ctx, subnetAInfo, subnetBInfo, messageID, big.NewInt(1), mockTokenAddress, fundedAddress, fundedKey, subnetATeleporterMessenger)

	// relay message from SubnetA to SubnetB
	network.RelayMessage(ctx, sendCrossChainMsgReceipt.BlockHash, sendCrossChainMsgReceipt.BlockNumber, subnetAInfo, subnetBInfo)

	// Check Teleporter message received on the destination
	delivered, err := subnetBTeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// send message from SubnetB to SubnetA
	sendCrossChainMessageInput.DestinationChainID = subnetAInfo.BlockchainID

	sendCrossChainMsgReceipt, messageID = utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, subnetBInfo, subnetAInfo, sendCrossChainMessageInput, fundedAddress, fundedKey, subnetBTeleporterMessenger)

	// relay message from SubnetB to SubnetA
	network.RelayMessage(ctx, sendCrossChainMsgReceipt.BlockHash, sendCrossChainMsgReceipt.BlockNumber, subnetBInfo, subnetAInfo)

	// check delivered
	delivered, err = subnetATeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetBInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// check the fee amount
	amount, err := subnetATeleporterMessenger.CheckRelayerRewardAmount(&bind.CallOpts{}, fundedAddress, mockTokenAddress)
	Expect(err).Should(BeNil())
	Expect(amount).Should(Equal(big.NewInt(1)))
}
