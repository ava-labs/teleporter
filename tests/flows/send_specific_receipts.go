package flows

import (
	"bytes"
	"context"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	teleporterutils "github.com/ava-labs/teleporter/utils/teleporter-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func SendSpecificReceipts(network interfaces.Network) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := utils.GetTwoSubnets(network)
	teleporterContractAddress := network.GetTeleporterContractAddress()
	_, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	// Clear the receipt queue from Subnet B -> Subnet A to have a clean slate for the test flow.
	// This is only done if the test non-external networks because external networks may have
	// an arbitrarily high number of receipts to be cleared from a given queue from unrelated messages.
	if !network.IsExternalNetwork() {
		utils.ClearReceiptQueue(ctx, network, fundedKey, subnetBInfo, subnetAInfo)
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
	Expect(receiveEvent1.MessageID[:]).Should(Equal(messageID1[:]))

	// Check that the first message was delivered
	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, messageID1)
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
	Expect(receiveEvent2.MessageID[:]).Should(Equal(messageID2[:]))

	// Check that the second message was delivered
	delivered, err = subnetBInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, messageID2)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Call send specific receipts to get reward of relaying two messages
	receipt, messageID := utils.SendSpecifiedReceiptsAndWaitForAcceptance(
		ctx,
		subnetAInfo.BlockchainID,
		subnetBInfo,
		[][32]byte{messageID1, messageID2},
		teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: mockTokenAddress,
			Amount:          big.NewInt(0),
		},
		[]common.Address{},
		fundedKey,
	)

	// Relay message from Subnet B to Subnet A
	receipt = network.RelayMessage(ctx, receipt, subnetBInfo, subnetAInfo, true)

	// Check that the message back to Subnet A was delivered
	delivered, err = subnetAInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check that the expected receipts were received and emitted ReceiptReceived
	Expect(utils.CheckReceiptReceived(receipt, messageID1, subnetAInfo.TeleporterMessenger)).Should(BeTrue())
	Expect(utils.CheckReceiptReceived(receipt, messageID2, subnetAInfo.TeleporterMessenger)).Should(BeTrue())

	// Check the reward amounts.
	// Even on external networks, the relayer should only have the expected fee amount
	// for this asset because the asset contract was newly deployed by this test.
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
		delivered, err = subnetAInfo.TeleporterMessenger.MessageReceived(
			&bind.CallOpts{},
			messageID)
		Expect(err).Should(BeNil())
		Expect(delivered).Should(BeTrue())

		// Check that the expected receipts were included in the message but did not emit ReceiptReceived
		// because they were previously received
		Expect(utils.CheckReceiptReceived(receipt, messageID1, subnetAInfo.TeleporterMessenger)).Should(BeFalse())
		Expect(utils.CheckReceiptReceived(receipt, messageID2, subnetAInfo.TeleporterMessenger)).Should(BeFalse())

		receiveEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			subnetAInfo.TeleporterMessenger.ParseReceiveCrossChainMessage,
		)
		Expect(err).Should(BeNil())
		log.Info("Receipt included", "count", len(receiveEvent.Message.Receipts), "receipts", receiveEvent.Message.Receipts)
		Expect(receiptIncluded(
			teleporterContractAddress,
			messageID1,
			subnetAInfo,
			subnetBInfo,
			receiveEvent.Message.Receipts)).Should(BeTrue())
		Expect(receiptIncluded(
			teleporterContractAddress,
			messageID2,
			subnetAInfo,
			subnetBInfo,
			receiveEvent.Message.Receipts)).Should(BeTrue())

		// Check the reward amount remains the same
		checkExpectedRewardAmounts(subnetAInfo, receiveEvent1, receiveEvent2, mockTokenAddress, relayerFeePerMessage)
	}
}

// Checks the given message ID is included in the list of receipts.
func receiptIncluded(
	teleporterMessengerAddress common.Address,
	expectedMessageID ids.ID,
	sourceSubnet interfaces.SubnetTestInfo,
	destinationSubnet interfaces.SubnetTestInfo,
	receipts []teleportermessenger.TeleporterMessageReceipt,
) bool {
	for _, receipt := range receipts {
		messageID, err := teleporterutils.CalculateMessageID(
			teleporterMessengerAddress,
			sourceSubnet.BlockchainID,
			destinationSubnet.BlockchainID,
			receipt.ReceivedMessageNonce,
		)
		Expect(err).Should(BeNil())
		if bytes.Equal(messageID[:], expectedMessageID[:]) {
			return true
		}
	}
	return false
}

// Checks that the reward redeemers specified by the two provided message receipts
// are able to redeem the correct amount of the rewards of the given token. It is
// assumed that the {tokenAddress} was used as the fee asset for each of the messages,
// and that each message individually had a fee of {feePerMessage}.
func checkExpectedRewardAmounts(
	sourceSubnet interfaces.SubnetTestInfo,
	receiveEvent1 *teleportermessenger.TeleporterMessengerReceiveCrossChainMessage,
	receiveEvent2 *teleportermessenger.TeleporterMessengerReceiveCrossChainMessage,
	tokenAddress common.Address,
	feePerMessage *big.Int,
) {
	// Check the reward amounts.
	// If the same address is the reward redeemer for both messages,
	// it should be able to redeem {feePerMessage}*2. Otherwise,
	// each distinct reward redeemer should be able to redeem {feePerMessage}.
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
