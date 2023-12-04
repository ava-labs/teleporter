package flows

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func SendSpecificReceipts(network interfaces.Network) {
	subnetAInfo, subnetBInfo, _ := utils.GetThreeSubnets(network)
	teleporterContractAddress := network.GetTeleporterContractAddress()
	_, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	// Clear the receipt queue from Subnet B -> Subnet A to have a clean slate for the test flow.
	// This is only done if the test non-external networks because external networks may have
	// an arbitrarily high number of receipts to be cleared from a given queue from unrelated messages.
	if !network.IsExternalNetwork() {
		clearReceiptQueue(ctx, network, fundedKey, subnetBInfo, subnetAInfo)
	}

	// Use mock token as the fee token
	mockTokenAddress, mockToken := utils.DeployExampleERC20(
		ctx, fundedKey, subnetAInfo,
	)
	utils.ERC20Approve(
		ctx,
		mockToken,
		teleporterContractAddress,
		big.NewInt(0).Mul(big.NewInt(1e18),
			big.NewInt(10)),
		subnetAInfo,
		fundedKey,
	)

	// Send two messages from Subnet A to Subnet B
	relayerFeePerMessage := big.NewInt(5)
	destinationAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: subnetBInfo.BlockchainID,
		DestinationAddress:      destinationAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: mockTokenAddress,
			Amount:          relayerFeePerMessage,
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	// Send first message from Subnet A to Subnet B with fee amount 5
	sendCrossChainMsgReceipt, messageID1 := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, subnetAInfo, subnetBInfo, sendCrossChainMessageInput, fundedKey)

	// Relay the message from SubnetA to SubnetB
	deliveryReceipt1 := network.RelayMessage(ctx, sendCrossChainMsgReceipt, subnetAInfo, subnetBInfo, true)
	receiveEvent1, err := utils.GetEventFromLogs(
		deliveryReceipt1.Logs,
		subnetBInfo.TeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())

	// Check that the first message was delivered
	delivered, err :=
		subnetBInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, messageID1)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Send second message from Subnet A to Subnet B with fee amount 5
	sendCrossChainMsgReceipt, messageID2 := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, subnetAInfo, subnetBInfo, sendCrossChainMessageInput, fundedKey)

	// Relay the message from SubnetA to SubnetB
	deliveryReceipt2 := network.RelayMessage(ctx, sendCrossChainMsgReceipt, subnetAInfo, subnetBInfo, true)
	receiveEvent2, err := utils.GetEventFromLogs(
		deliveryReceipt2.Logs,
		subnetBInfo.TeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())

	// Check that the second message was delivered
	delivered, err =
		subnetBInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, messageID2)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Call send specific receipts to get reward of relaying two messages
	receipt, messageID := utils.SendSpecifiedReceiptsAndWaitForAcceptance(
		ctx,
		subnetAInfo.BlockchainID,
		subnetBInfo,
		[]*big.Int{messageID1, messageID2},
		teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: mockTokenAddress,
			Amount:          big.NewInt(0),
		},
		[]common.Address{},
		fundedKey,
	)

	// Relay message from Subnet B to Subnet A
	network.RelayMessage(ctx, receipt, subnetBInfo, subnetAInfo, true)

	// Check that the message back to Subnet A was delivered
	delivered, err = subnetAInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetBInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check the reward amounts.
	checkExpectedRewardAmounts(subnetAInfo, receiveEvent1, receiveEvent2, mockTokenAddress, relayerFeePerMessage)

	// If the network is internal to the test application, send a message from Subnet B to Subnet A to trigger
	// the "regular" method of delivering receipts. The next message from B->A will contain the same receipts
	// that were manually sent in the above steps, but they should not be processed again on Subnet A.
	// These checks are not performed for external networks because unrelated messages may have already changed
	// the state of the receipt queues.
	if !network.IsExternalNetwork() {
		sendCrossChainMessageInput = teleportermessenger.TeleporterMessageInput{
			DestinationBlockchainID: subnetAInfo.BlockchainID,
			DestinationAddress:      destinationAddress,
			FeeInfo: teleportermessenger.TeleporterFeeInfo{
				FeeTokenAddress: mockTokenAddress,
				Amount:          big.NewInt(0),
			},
			RequiredGasLimit:        big.NewInt(1),
			AllowedRelayerAddresses: []common.Address{},
			Message:                 []byte{1, 2, 3, 4},
		}

		// This message will also have the same receipts as the previous message
		receipt, messageID = utils.SendCrossChainMessageAndWaitForAcceptance(
			ctx, subnetBInfo, subnetAInfo, sendCrossChainMessageInput, fundedKey)

		// Relay message from Subnet B to Subnet A
		receipt = network.RelayMessage(ctx, receipt, subnetBInfo, subnetAInfo, true)
		// Check delivered
		delivered, err = subnetAInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetBInfo.BlockchainID, messageID)
		Expect(err).Should(BeNil())
		Expect(delivered).Should(BeTrue())
		// Get the Teleporter message from receive event and confirm that the receipts are delivered again
		receiveEvent, err :=
			utils.GetEventFromLogs(receipt.Logs, subnetAInfo.TeleporterMessenger.ParseReceiveCrossChainMessage)
		Expect(err).Should(BeNil())
		log.Info("Receipt included", "count", len(receiveEvent.Message.Receipts), "receipts", receiveEvent.Message.Receipts)
		Expect(receiptIncluded(messageID1, receiveEvent.Message.Receipts)).Should(BeTrue())
		Expect(receiptIncluded(messageID2, receiveEvent.Message.Receipts)).Should(BeTrue())

		// Check the reward amount remains the same
		checkExpectedRewardAmounts(subnetAInfo, receiveEvent1, receiveEvent2, mockTokenAddress, relayerFeePerMessage)
	}
}

func clearReceiptQueue(
	ctx context.Context,
	network interfaces.Network,
	fundedKey *ecdsa.PrivateKey,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
) {
	outstandReceiptCount := getOutstandingReceiptCount(source, destination.BlockchainID)
	for outstandReceiptCount.Cmp(big.NewInt(0)) != 0 {
		log.Info("Emptying receipt queue", "remainingReceipts", outstandReceiptCount.String())
		// Send message from Subnet B to Subnet A to trigger the "regular" method of delivering receipts.
		// The next message from B->A will contain the same receipts that were manually sent in the above steps,
		// but they should not be processed again on Subnet A.
		sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
			DestinationBlockchainID: destination.BlockchainID,
			DestinationAddress:      common.HexToAddress("0x1111111111111111111111111111111111111111"),
			RequiredGasLimit:        big.NewInt(1),
			FeeInfo: teleportermessenger.TeleporterFeeInfo{
				FeeTokenAddress: common.Address{},
				Amount:          big.NewInt(0),
			},
			AllowedRelayerAddresses: []common.Address{},
			Message:                 []byte{1, 2, 3, 4},
		}

		// This message will also have the same receipts as the previous message
		receipt, _ := utils.SendCrossChainMessageAndWaitForAcceptance(
			ctx, source, destination, sendCrossChainMessageInput, fundedKey)

		// Relay message
		network.RelayMessage(ctx, receipt, source, destination, true)

		outstandReceiptCount = getOutstandingReceiptCount(source, destination.BlockchainID)
	}
	log.Info("Receipt queue emptied")
}

func getOutstandingReceiptCount(source interfaces.SubnetTestInfo, destinationBlockchainID ids.ID) *big.Int {
	size, err := source.TeleporterMessenger.GetReceiptQueueSize(&bind.CallOpts{}, destinationBlockchainID)
	Expect(err).Should(BeNil())
	return size
}

// Checks the given message ID is included in the list of receipts.
func receiptIncluded(
	expectedMessageID *big.Int,
	receipts []teleportermessenger.TeleporterMessageReceipt,
) bool {
	for _, receipt := range receipts {
		if receipt.ReceivedMessageID.Cmp(expectedMessageID) == 0 {
			return true
		}
	}
	return false
}

func checkExpectedRewardAmounts(
	sourceSubnet interfaces.SubnetTestInfo,
	receiveEvent1 *teleportermessenger.TeleporterMessengerReceiveCrossChainMessage,
	receiveEvent2 *teleportermessenger.TeleporterMessengerReceiveCrossChainMessage,
	tokenAddress common.Address,
	feePerMessage *big.Int,
) {
	// Check the reward amounts.
	if receiveEvent1.RewardRedeemer == receiveEvent2.RewardRedeemer {
		amount, err := sourceSubnet.TeleporterMessenger.CheckRelayerRewardAmount(
			&bind.CallOpts{},
			receiveEvent1.RewardRedeemer,
			tokenAddress)
		Expect(err).Should(BeNil())
		Expect(amount).Should(Equal(new(big.Int).Mul(feePerMessage, big.NewInt(2))))
	} else {
		amount1, err := sourceSubnet.TeleporterMessenger.CheckRelayerRewardAmount(
			&bind.CallOpts{},
			receiveEvent1.RewardRedeemer,
			tokenAddress)
		Expect(err).Should(BeNil())
		Expect(amount1).Should(Equal(feePerMessage))
		amount2, err := sourceSubnet.TeleporterMessenger.CheckRelayerRewardAmount(
			&bind.CallOpts{},
			receiveEvent2.RewardRedeemer,
			tokenAddress)
		Expect(err).Should(BeNil())
		Expect(amount2).Should(Equal(feePerMessage))
	}
}
