package teleporter

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

func AddFeeAmount(network interfaces.Network, teleporterInfo utils.TeleporterTestInfo) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := utils.GetTwoSubnets(network)
	teleporterContractAddress := teleporterInfo[subnetAInfo.BlockchainID].TeleporterMessengerAddress
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	// Use mock token as the fee token
	mockTokenAddress, mockToken := utils.DeployExampleERC20(
		context.Background(),
		fundedKey,
		subnetAInfo,
	)
	utils.ERC20Approve(
		ctx,
		mockToken,
		teleporterContractAddress,
		big.NewInt(1e18),
		subnetAInfo,
		fundedKey,
	)

	initFeeAmount := big.NewInt(1)

	// Send a transaction to Subnet A that sends a message to Subnet B.
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: subnetBInfo.BlockchainID,
		DestinationAddress:      fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: mockTokenAddress,
			Amount:          initFeeAmount,
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	sendCrossChainMsgReceipt, messageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		teleporterInfo[subnetAInfo.BlockchainID].TeleporterMessenger,
		subnetAInfo,
		subnetBInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	// Add a fee amount to the message.
	additionalFeeAmount := big.NewInt(2)
	utils.SendAddFeeAmountAndWaitForAcceptance(
		ctx,
		subnetAInfo,
		subnetBInfo,
		messageID,
		additionalFeeAmount,
		mockTokenAddress,
		fundedKey,
		teleporterInfo[subnetAInfo.BlockchainID].TeleporterMessenger,
	)

	// Relay message from Subnet A to Subnet B
	deliveryReceipt := teleporterInfo.RelayTeleporterMessage(ctx, sendCrossChainMsgReceipt, subnetAInfo, subnetBInfo, true, fundedKey)
	receiveEvent, err := utils.GetEventFromLogs(
		deliveryReceipt.Logs,
		teleporterInfo[subnetBInfo.BlockchainID].TeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())

	// Check Teleporter message received on the destination (Subnet B)
	delivered, err := teleporterInfo[subnetBInfo.BlockchainID].TeleporterMessenger.MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check the initial relayer reward amount on Subnet A.
	initialRewardAmount, err := teleporterInfo[subnetAInfo.BlockchainID].TeleporterMessenger.CheckRelayerRewardAmount(
		&bind.CallOpts{},
		receiveEvent.RewardRedeemer,
		mockTokenAddress)
	Expect(err).Should(BeNil())

	// Send a message from Subnet B back to Subnet A that includes the specific receipt for the message.
	sendSpecificReceiptsReceipt, sendSpecificReceiptsMessageID := utils.SendSpecifiedReceiptsAndWaitForAcceptance(
		ctx,
		teleporterInfo[subnetBInfo.BlockchainID].TeleporterMessenger,
		subnetBInfo,
		subnetAInfo.BlockchainID,
		[][32]byte{receiveEvent.MessageID},
		teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: mockTokenAddress,
			Amount:          big.NewInt(0),
		},
		[]common.Address{},
		fundedKey)

	// Relay message containing the specific receipt from Subnet B to Subnet A
	teleporterInfo.RelayTeleporterMessage(ctx, sendSpecificReceiptsReceipt, subnetBInfo, subnetAInfo, true, fundedKey)

	// Check message delivered
	delivered, err = teleporterInfo[subnetAInfo.BlockchainID].TeleporterMessenger.MessageReceived(&bind.CallOpts{}, sendSpecificReceiptsMessageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check the updated relayer reward amount
	expectedIncrease := new(big.Int).Add(initFeeAmount, additionalFeeAmount)
	newRewardAmount, err := teleporterInfo[subnetAInfo.BlockchainID].TeleporterMessenger.CheckRelayerRewardAmount(
		&bind.CallOpts{},
		receiveEvent.RewardRedeemer,
		mockTokenAddress)
	Expect(err).Should(BeNil())
	Expect(newRewardAmount).Should(Equal(new(big.Int).Add(initialRewardAmount, expectedIncrease)))

	// If the funded address is the one able to redeem the rewards, do so and check the reward amount is reset.
	if fundedAddress == receiveEvent.RewardRedeemer {
		utils.RedeemRelayerRewardsAndConfirm(
			ctx,
			teleporterInfo[subnetAInfo.BlockchainID].TeleporterMessenger,
			subnetAInfo,
			mockToken,
			mockTokenAddress,
			fundedKey,
			newRewardAmount,
		)
	}
}
