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
	. "github.com/onsi/gomega"
)

func AddFeeAmountGinkgo() {
	AddFeeAmount(&network.LocalNetwork{})
}

func AddFeeAmount(network network.Network) {
	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	// Use mock token as the fee token
	mockTokenAddress, mockToken := localUtils.DeployExampleERC20(
		context.Background(),
		fundedKey,
		subnetAInfo,
	)
	localUtils.ExampleERC20Approve(
		ctx,
		mockToken,
		teleporterContractAddress,
		big.NewInt(0).Mul(big.NewInt(1e18),
			big.NewInt(10)),
		subnetAInfo,
		fundedKey,
	)

	initFeeAmount := big.NewInt(1)
	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationChainID: subnetBInfo.BlockchainID,
		DestinationAddress: fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: mockTokenAddress,
			Amount:          initFeeAmount,
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	sendCrossChainMsgReceipt, messageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, subnetAInfo, subnetBInfo, sendCrossChainMessageInput, fundedKey, subnetAInfo.TeleporterMessenger)

	// Add fee amount
	additionalFeeAmount := big.NewInt(2)
	utils.SendAddFeeAmountAndWaitForAcceptance(
		ctx,
		subnetAInfo,
		subnetBInfo,
		messageID,
		additionalFeeAmount,
		mockTokenAddress,
		fundedKey,
		subnetAInfo.TeleporterMessenger,
	)

	// Relay message from SubnetA to SubnetB
	network.RelayMessage(ctx, sendCrossChainMsgReceipt, subnetAInfo, subnetBInfo, true)

	// Check Teleporter message received on the destination
	delivered, err :=
		subnetBInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Send message from SubnetB to SubnetA. This will include the receipt for the previous message from A->B
	sendCrossChainMessageInput.DestinationChainID = subnetAInfo.BlockchainID
	sendCrossChainMessageInput.FeeInfo.Amount = big.NewInt(0)

	sendCrossChainMsgReceipt, messageID = utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, subnetBInfo, subnetAInfo, sendCrossChainMessageInput, fundedKey, subnetBInfo.TeleporterMessenger)

	// Relay message from SubnetB to SubnetA
	network.RelayMessage(ctx, sendCrossChainMsgReceipt, subnetBInfo, subnetAInfo, true)

	// Check message delivered
	delivered, err = subnetAInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetBInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check the relayer reward amount
	amount, err :=
		subnetAInfo.TeleporterMessenger.CheckRelayerRewardAmount(&bind.CallOpts{}, fundedAddress, mockTokenAddress)
	Expect(err).Should(BeNil())
	Expect(amount).Should(Equal(additionalFeeAmount.Add(additionalFeeAmount, initFeeAmount)))

	utils.RedeemRelayerRewardsAndConfirm(
		ctx, subnetAInfo, mockToken, mockTokenAddress, fundedKey, amount,
	)
}
