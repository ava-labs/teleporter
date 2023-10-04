package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/teleporter/tests/utils"

	"github.com/ava-labs/awm-relayer/messages/teleporter"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

// Ginkgo describe node that acts as a container for the teleporter e2e tests. This test suite
// will run through the following steps in order:
// 1. Send a transaction to the Teleporter contract on Subnet A
// 2. Aggregate signatures and send the Warp message to Subnet B
// 3. Verify receipt of the message on Subnet B
func BasicOneWaySend() {
	var (
		teleporterMessageID *big.Int
	)

	subnetAInfo := utils.GetSubnetATestInfo()
	subnetBInfo := utils.GetSubnetBTestInfo()
	teleporterContractAddress := utils.GetTeleporterContractAddress()
	fundedAddress, fundedKey := utils.GetFundedAccountInfo()

	//
	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	//
	ctx := context.Background()

	sendCrossChainMessageInput := utils.SendCrossChainMessageInput{
		DestinationChainID: subnetBInfo.BlockchainID,
		DestinationAddress: fundedAddress,
		FeeInfo: utils.FeeInfo{
			ContractAddress: fundedAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}
	signedTx := utils.CreateSendCrossChainMessageTransaction(ctx, subnetAInfo, sendCrossChainMessageInput, fundedAddress, fundedKey, teleporterContractAddress)

	log.Info("Sending Teleporter transaction on source chain", "destinationChainID", subnetBInfo.BlockchainID, "txHash", signedTx.Hash())
	receipt := utils.SendTransactionAndWaitForAcceptance(ctx, subnetAInfo.ChainWSClient, signedTx)

	sendCrossChainMessageLog := receipt.Logs[0]
	var event utils.SendCrossChainMessageEvent
	err := teleporter.EVMTeleporterContractABI.UnpackIntoInterface(&event, "SendCrossChainMessage", sendCrossChainMessageLog.Data)
	Expect(err).Should(BeNil())

	teleporterMessageID = event.Message.MessageID

	//
	// Relay the message to the destination
	//

	utils.RelayMessage(ctx, receipt.BlockHash, receipt.BlockNumber, subnetAInfo, subnetBInfo)

	//
	// Check Teleporter message received on the destination
	//
	data, err := teleporter.PackMessageReceived(teleporter.MessageReceivedInput{
		OriginChainID: subnetAInfo.BlockchainID,
		MessageID:     teleporterMessageID,
	})
	Expect(err).Should(BeNil())
	callMessage := interfaces.CallMsg{
		To:   &teleporterContractAddress,
		Data: data,
	}
	result, err := subnetBInfo.ChainWSClient.CallContract(context.Background(), callMessage, nil)
	Expect(err).Should(BeNil())

	// check the contract call result
	delivered, err := teleporter.UnpackMessageReceivedResult(result)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())
}
