package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	localUtils "github.com/ava-labs/teleporter/tests/utils/local-network-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

func SendSpecificReceiptsGinkgo() {
	SendSpecificReceipts(&network.LocalNetwork{})
}

func SendSpecificReceipts(network network.Network) {
	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress, subnetAInfo.ChainRPCClient,
	)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress, subnetBInfo.ChainRPCClient,
	)
	Expect(err).Should(BeNil())

	// Use mock token as the fee token
	mockTokenAddress, mockToken := localUtils.DeployExampleERC20(
		context.Background(), fundedAddress, fundedKey, subnetAInfo,
	)
	localUtils.ExampleERC20Approve(
		ctx,
		mockToken,
		teleporterContractAddress,
		big.NewInt(0).Mul(big.NewInt(1e18),
			big.NewInt(10)),
		subnetAInfo,
		fundedAddress,
		fundedKey,
	)

	relayerFeePerMessage := big.NewInt(5)
	totalAccumulatedRelayerFee := big.NewInt(10)

	destinationKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	destinationAddress := crypto.PubkeyToAddress(destinationKey.PublicKey)

	// Send two messages from Subnet A to Subnet B
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationChainID: subnetBInfo.BlockchainID,
		DestinationAddress: destinationAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			ContractAddress: mockTokenAddress,
			Amount:          relayerFeePerMessage,
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	// Send first message from Subnet A to Subnet B with fee amount 5
	sendCrossChainMsgReceipt, messageID1 := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, subnetAInfo, subnetBInfo, sendCrossChainMessageInput, fundedAddress, fundedKey, subnetATeleporterMessenger)

	// Relay message from SubnetA to SubnetB
	network.RelayMessage(ctx, sendCrossChainMsgReceipt, subnetAInfo, subnetBInfo, false, true)
	// Check messge delivered
	delivered, err := subnetBTeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, messageID1)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Send second message from Subnet A to Subnet B with fee amount 5
	sendCrossChainMsgReceipt, messageID2 := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, subnetAInfo, subnetBInfo, sendCrossChainMessageInput, fundedAddress, fundedKey, subnetATeleporterMessenger)

	// Relay message from SubnetA to SubnetB
	network.RelayMessage(ctx, sendCrossChainMsgReceipt, subnetAInfo, subnetBInfo, false, true)
	// Check delivered
	delivered, err = subnetBTeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, messageID2)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Relayer send specific receipts to get reward of relaying two messages
	receipt, messageID := utils.SendSpecifiedReceiptsAndWaitForAcceptance(
		ctx,
		subnetAInfo.BlockchainID,
		subnetBInfo,
		[]*big.Int{messageID1, messageID2},
		teleportermessenger.TeleporterFeeInfo{
			ContractAddress: mockTokenAddress,
			Amount:          big.NewInt(0),
		},
		[]common.Address{},
		fundedAddress,
		fundedKey,
		subnetBTeleporterMessenger,
	)

	// Relay message from Subnet B to Subnet A
	network.RelayMessage(ctx, receipt, subnetBInfo, subnetAInfo, false, true)
	// Check delivered
	delivered, err = subnetATeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetBInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check the reward amount. The reward amount should be 10
	amount, err := subnetATeleporterMessenger.CheckRelayerRewardAmount(&bind.CallOpts{}, fundedAddress, mockTokenAddress)
	Expect(err).Should(BeNil())
	Expect(amount).Should(Equal(totalAccumulatedRelayerFee))

	// Send message from Subnet B to Subnet A to trigger the "regular" method of delivering receipts.
	// The next message from B->A will contain the same receipts that were manually sent in the above steps,
	// but they should not be processed again on Subnet A.
	sendCrossChainMessageInput = teleportermessenger.TeleporterMessageInput{
		DestinationChainID: subnetAInfo.BlockchainID,
		DestinationAddress: destinationAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			ContractAddress: mockTokenAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	// This message will also have the same receipt as the previous message
	receipt, messageID = utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, subnetBInfo, subnetAInfo, sendCrossChainMessageInput, fundedAddress, fundedKey, subnetBTeleporterMessenger)

	// Relay message from Subnet B to Subnet A
	network.RelayMessage(ctx, receipt, subnetBInfo, subnetAInfo, false, true)
	// Check delivered
	delivered, err = subnetATeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetBInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check the reward amount remains the same
	amount, err = subnetATeleporterMessenger.CheckRelayerRewardAmount(&bind.CallOpts{}, fundedAddress, mockTokenAddress)
	Expect(err).Should(BeNil())
	Expect(amount).Should(Equal(totalAccumulatedRelayerFee))
}
