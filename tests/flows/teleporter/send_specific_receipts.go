package teleporter

import (
	"bytes"
	"context"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	teleporterutils "github.com/ava-labs/teleporter/utils/teleporter-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func SendSpecificReceipts(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	L1AInfo := network.GetPrimaryNetworkInfo()
	L1BInfo, _ := network.GetTwoL1s()
	l1ATeleporterMessenger := teleporter.TeleporterMessenger(L1AInfo)
	l1BTeleporterMessenger := teleporter.TeleporterMessenger(L1BInfo)
	teleporterContractAddress := teleporter.TeleporterMessengerAddress(L1AInfo)
	_, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	// Clear the receipt queue from L1 B -> L1 A to have a clean slate for the test flow.
	network.ClearReceiptQueue(ctx, teleporter, fundedKey, L1BInfo, L1AInfo)

	// Use mock token as the fee token
	mockTokenAddress, mockToken := utils.DeployExampleERC20(
		ctx, fundedKey, L1AInfo,
	)
	utils.ERC20Approve(
		ctx,
		mockToken,
		teleporterContractAddress,
		big.NewInt(0).Mul(big.NewInt(1e18),
			big.NewInt(10)),
		L1AInfo,
		fundedKey,
	)

	// Send two messages from L1 A to L1 B
	relayerFeePerMessage := big.NewInt(5)
	destinationAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: L1BInfo.BlockchainID,
		DestinationAddress:      destinationAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: mockTokenAddress,
			Amount:          relayerFeePerMessage,
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	// Send first message from L1 A to L1 B with fee amount 5
	sendCrossChainMsgReceipt, messageID1 := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, l1ATeleporterMessenger, L1AInfo, L1BInfo, sendCrossChainMessageInput, fundedKey)

	// Relay the message from L1A to L1B
	deliveryReceipt1 := teleporter.RelayTeleporterMessage(
		ctx,
		sendCrossChainMsgReceipt,
		L1AInfo,
		L1BInfo,
		true,
		fundedKey,
	)
	receiveEvent1, err := utils.GetEventFromLogs(
		deliveryReceipt1.Logs,
		l1BTeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(receiveEvent1.MessageID[:]).Should(Equal(messageID1[:]))

	// Check that the first message was delivered
	delivered, err := l1BTeleporterMessenger.MessageReceived(&bind.CallOpts{}, messageID1)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Send second message from L1 A to L1 B with fee amount 5
	sendCrossChainMsgReceipt, messageID2 := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, l1ATeleporterMessenger, L1AInfo, L1BInfo, sendCrossChainMessageInput, fundedKey)

	// Relay the message from L1A to L1B
	deliveryReceipt2 := teleporter.RelayTeleporterMessage(
		ctx,
		sendCrossChainMsgReceipt,
		L1AInfo,
		L1BInfo,
		true,
		fundedKey,
	)
	receiveEvent2, err := utils.GetEventFromLogs(
		deliveryReceipt2.Logs,
		l1BTeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(receiveEvent2.MessageID[:]).Should(Equal(messageID2[:]))

	// Check that the second message was delivered
	delivered, err = l1BTeleporterMessenger.MessageReceived(&bind.CallOpts{}, messageID2)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Call send specific receipts to get reward of relaying two messages
	receipt, messageID := utils.SendSpecifiedReceiptsAndWaitForAcceptance(
		ctx,
		l1BTeleporterMessenger,
		L1BInfo,
		L1AInfo.BlockchainID,
		[][32]byte{messageID1, messageID2},
		teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: mockTokenAddress,
			Amount:          big.NewInt(0),
		},
		[]common.Address{},
		fundedKey,
	)

	// Relay message from L1 B to L1 A
	receipt = teleporter.RelayTeleporterMessage(ctx, receipt, L1BInfo, L1AInfo, true, fundedKey)

	// Check that the message back to L1 A was delivered
	delivered, err = l1ATeleporterMessenger.MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check that the expected receipts were received and emitted ReceiptReceived
	Expect(utils.CheckReceiptReceived(receipt, messageID1, l1ATeleporterMessenger)).Should(BeTrue())
	Expect(utils.CheckReceiptReceived(receipt, messageID2, l1ATeleporterMessenger)).Should(BeTrue())

	// Check the reward amounts.
	// Even on external networks, the relayer should only have the expected fee amount
	// for this asset because the asset contract was newly deployed by this test.
	checkExpectedRewardAmounts(
		teleporter,
		L1AInfo,
		receiveEvent1,
		receiveEvent2,
		mockTokenAddress,
		relayerFeePerMessage,
	)

	// Send a message from L1 B to L1 A to trigger
	// the "regular" method of delivering receipts. The next message from B->A will contain the same receipts
	// that were manually sent in the above steps, but they should not be processed again on L1 A.
	sendCrossChainMessageInput = teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: L1AInfo.BlockchainID,
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
		ctx, l1BTeleporterMessenger, L1BInfo, L1AInfo, sendCrossChainMessageInput, fundedKey)

	// Relay message from L1 B to L1 A
	receipt = teleporter.RelayTeleporterMessage(ctx, receipt, L1BInfo, L1AInfo, true, fundedKey)
	// Check delivered
	delivered, err = l1ATeleporterMessenger.MessageReceived(
		&bind.CallOpts{},
		messageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check that the expected receipts were included in the message but did not emit ReceiptReceived
	// because they were previously received
	Expect(utils.CheckReceiptReceived(receipt, messageID1, l1ATeleporterMessenger)).Should(BeFalse())
	Expect(utils.CheckReceiptReceived(receipt, messageID2, l1ATeleporterMessenger)).Should(BeFalse())

	receiveEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		l1ATeleporterMessenger.ParseReceiveCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	log.Info("Receipt included", "count", len(receiveEvent.Message.Receipts), "receipts", receiveEvent.Message.Receipts)
	Expect(receiptIncluded(
		teleporterContractAddress,
		messageID1,
		L1AInfo,
		L1BInfo,
		receiveEvent.Message.Receipts)).Should(BeTrue())
	Expect(receiptIncluded(
		teleporterContractAddress,
		messageID2,
		L1AInfo,
		L1BInfo,
		receiveEvent.Message.Receipts)).Should(BeTrue())

	// Check the reward amount remains the same
	checkExpectedRewardAmounts(
		teleporter,
		L1AInfo,
		receiveEvent1,
		receiveEvent2,
		mockTokenAddress,
		relayerFeePerMessage,
	)
}

// Checks the given message ID is included in the list of receipts.
func receiptIncluded(
	teleporterMessengerAddress common.Address,
	expectedMessageID ids.ID,
	sourceL1 interfaces.L1TestInfo,
	destinationL1 interfaces.L1TestInfo,
	receipts []teleportermessenger.TeleporterMessageReceipt,
) bool {
	for _, receipt := range receipts {
		messageID, err := teleporterutils.CalculateMessageID(
			teleporterMessengerAddress,
			sourceL1.BlockchainID,
			destinationL1.BlockchainID,
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
	teleporter utils.TeleporterTestInfo,
	sourceL1 interfaces.L1TestInfo,
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
		amount, err := teleporter[sourceL1.BlockchainID].TeleporterMessenger.CheckRelayerRewardAmount(
			&bind.CallOpts{},
			receiveEvent1.RewardRedeemer,
			tokenAddress)
		Expect(err).Should(BeNil())
		Expect(amount).Should(Equal(new(big.Int).Mul(feePerMessage, big.NewInt(2))))
	} else {
		amount1, err := teleporter[sourceL1.BlockchainID].TeleporterMessenger.CheckRelayerRewardAmount(
			&bind.CallOpts{},
			receiveEvent1.RewardRedeemer,
			tokenAddress)
		Expect(err).Should(BeNil())
		Expect(amount1).Should(Equal(feePerMessage))
		amount2, err := teleporter[sourceL1.BlockchainID].TeleporterMessenger.CheckRelayerRewardAmount(
			&bind.CallOpts{},
			receiveEvent2.RewardRedeemer,
			tokenAddress)
		Expect(err).Should(BeNil())
		Expect(amount2).Should(Equal(feePerMessage))
	}
}
