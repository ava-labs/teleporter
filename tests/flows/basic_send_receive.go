package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

// Tests basic one-way send from Subnet A to Subnet B and vice versa
func BasicSendReceive(network interfaces.Network) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := utils.GetTwoSubnets(network)
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	//
	ctx := context.Background()

	// Clear the receipt queue from Subnet B -> Subnet A to have a clean slate for the test flow.
	// This is only done if the test non-external networks because external networks may have
	// an arbitrarily high number of receipts to be cleared from a given queue from unrelated messages.
	if !network.IsExternalNetwork() {
		utils.ClearReceiptQueue(ctx, network, fundedKey, subnetBInfo, subnetAInfo)
	}

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
		subnetAInfo,
		subnetBInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)
	expectedReceiptID := teleporterMessageID

	//
	// Relay the message to the destination
	//
	deliveryReceipt := network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)
	receiveEvent, err := utils.GetEventFromLogs(
		deliveryReceipt.Logs,
		subnetBInfo.TeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())

	//
	// Check Teleporter message received on the destination
	//
	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Send a transaction to Subnet B to issue a Warp Message from the Teleporter contract to Subnet A
	//
	sendCrossChainMessageInput.DestinationBlockchainID = subnetAInfo.BlockchainID
	sendCrossChainMessageInput.FeeInfo.Amount = big.NewInt(0)
	receipt, teleporterMessageID = utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		subnetBInfo,
		subnetAInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	//
	// Relay the message to the destination
	//
	deliveryReceipt = network.RelayMessage(ctx, receipt, subnetBInfo, subnetAInfo, true)

	// Check that the receipt was received for expected Teleporter message ID
	// This check is not performed for external networks because unrelated messages may have already changed
	// the state of the receipt queues.
	if !network.IsExternalNetwork() {
		Expect(utils.CheckReceiptReceived(
			deliveryReceipt,
			expectedReceiptID,
			subnetAInfo.TeleporterMessenger)).Should(BeTrue())
	}

	//
	// Check Teleporter message received on the destination
	//
	delivered, err = subnetAInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// If the reward address of the message from A->B is the funded address, which is able to send
	// transactions on subnet A, then redeem the rewards.
	if receiveEvent.RewardRedeemer == fundedAddress {
		utils.RedeemRelayerRewardsAndConfirm(
			ctx, subnetAInfo, feeToken, feeTokenAddress, fundedKey, feeAmount,
		)
	}
}
