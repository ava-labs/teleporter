package teleporter

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/teleporter/TeleporterMessenger"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

// Tests basic one-way send from Subnet A to Subnet B and vice versa
func BasicSendReceive(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := network.GetTwoSubnets()
	teleporterContractAddress := teleporter.TeleporterMessengerAddress(subnetAInfo)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	ctx := context.Background()

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	// Clear the receipt queue from Subnet B -> Subnet A to have a clean slate for the test flow.
	// This is only done if the test non-external networks because external networks may have
	// an arbitrarily high number of receipts to be cleared from a given queue from unrelated messages.
	teleporter.ClearReceiptQueue(ctx, fundedKey, subnetBInfo, subnetAInfo, aggregator)

	feeAmount := big.NewInt(1)
	feeTokenAddress, feeToken := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		subnetAInfo,
	)
	utils.ERC20Approve(
		ctx,
		feeToken,
		teleporterContractAddress,
		big.NewInt(0).Mul(big.NewInt(1e18),
			big.NewInt(10)),
		subnetAInfo,
		fundedKey,
	)

	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: subnetBInfo.BlockchainID,
		DestinationAddress:      fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: feeTokenAddress,
			Amount:          feeAmount,
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	receipt, teleporterMessageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		teleporter.TeleporterMessenger(subnetAInfo),
		subnetAInfo,
		subnetBInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)
	expectedReceiptID := teleporterMessageID

	// Relay the message to the destination
	deliveryReceipt := teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		subnetAInfo,
		subnetBInfo,
		true,
		fundedKey,
		nil,
		aggregator,
	)
	receiveEvent, err := utils.GetEventFromLogs(
		deliveryReceipt.Logs,
		teleporter.TeleporterMessenger(subnetBInfo).ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())

	// Check Teleporter message received on the destination
	delivered, err := teleporter.TeleporterMessenger(subnetBInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Send a transaction to Subnet B to issue a Warp Message from the Teleporter contract to Subnet A
	sendCrossChainMessageInput.DestinationBlockchainID = subnetAInfo.BlockchainID
	sendCrossChainMessageInput.FeeInfo.Amount = big.NewInt(0)
	receipt, teleporterMessageID = utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		teleporter.TeleporterMessenger(subnetBInfo),
		subnetBInfo,
		subnetAInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	// Relay the message to the destination
	deliveryReceipt = teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		subnetBInfo,
		subnetAInfo,
		true,
		fundedKey,
		nil,
		aggregator,
	)

	Expect(utils.CheckReceiptReceived(
		deliveryReceipt,
		expectedReceiptID,
		teleporter.TeleporterMessenger(subnetAInfo))).Should(BeTrue())

	// Check Teleporter message received on the destination
	delivered, err = teleporter.TeleporterMessenger(subnetAInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// If the reward address of the message from A->B is the funded address, which is able to send
	// transactions on subnet A, then redeem the rewards.
	if receiveEvent.RewardRedeemer == fundedAddress {
		utils.RedeemRelayerRewardsAndConfirm(
			ctx, teleporter.TeleporterMessenger(subnetAInfo), subnetAInfo, feeToken, feeTokenAddress, fundedKey, feeAmount,
		)
	}
}
