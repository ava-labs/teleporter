package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	. "github.com/onsi/gomega"
)

func ReceiveMessageTwiceGinkgo() {
	ReceiveMessageTwice(&network.LocalNetwork{})
}

func ReceiveMessageTwice(network network.Network) {
	var (
		teleporterMessageID *big.Int
	)

	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetAInfo.ChainRPCClient)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetBInfo.ChainRPCClient)
	Expect(err).Should(BeNil())

	//
	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	//
	ctx := context.Background()

	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationChainID: subnetBInfo.BlockchainID,
		DestinationAddress: fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			ContractAddress: fundedAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}
	signedTx := utils.CreateSendCrossChainMessageTransaction(ctx, subnetAInfo, sendCrossChainMessageInput, fundedAddress, fundedKey, teleporterContractAddress)

	log.Info("Sending Teleporter transaction on source chain", "destinationChainID", subnetBInfo.BlockchainID, "txHash", signedTx.Hash())
	receipt := utils.SendTransactionAndWaitForAcceptance(ctx, subnetAInfo.ChainWSClient, subnetAInfo.ChainRPCClient, signedTx, true)

	event, err := utils.GetEventFromLogs(receipt.Logs, subnetATeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationChainID[:]).Should(Equal(subnetBInfo.BlockchainID[:]))

	teleporterMessageID = event.Message.MessageID

	//
	// Relay the message to the destination
	//

	receipt = network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, false, true)
	teleporterTx, _, err := subnetBInfo.ChainRPCClient.TransactionByHash(ctx, receipt.TxHash)

	//
	// Check Teleporter message received on the destination
	//
	log.Info("Checking the message was received on the destination")
	delivered, err := subnetBTeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, teleporterMessageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Attempt to send the same message again
	//
	log.Info("Submitting the same Teleporter message again on the destination")
	gasFeeCap, gasTipCap, nonce := utils.CalculateTxParams(ctx, subnetBInfo.ChainRPCClient, fundedAddress)
	secondTeleporterTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:    subnetBInfo.ChainIDInt,
		Nonce:      nonce,
		To:         &teleporterContractAddress,
		Gas:        teleporterTx.Gas(),
		GasFeeCap:  gasFeeCap,
		GasTipCap:  gasTipCap,
		Value:      big.NewInt(0),
		Data:       teleporterTx.Data(),
		AccessList: teleporterTx.AccessList(),
	})

	signedTx = utils.SignTransaction(secondTeleporterTx, fundedKey, subnetBInfo.ChainIDInt)
	utils.SendTransactionAndWaitForAcceptance(ctx, subnetBInfo.ChainWSClient, subnetBInfo.ChainRPCClient, signedTx, false)
}
