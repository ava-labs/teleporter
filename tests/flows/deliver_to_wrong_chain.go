package flows

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func DeliverToWrongChain(network interfaces.Network) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, subnetCInfo := utils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Get the expected teleporter message ID for Subnet C
	//
	expectedAtoCMessageID, err :=
		subnetAInfo.TeleporterMessenger.GetNextMessageID(&bind.CallOpts{}, subnetCInfo.BlockchainID)
	Expect(err).Should(BeNil())

	//
	// Submit a message to be sent from SubnetA to SubnetB
	//
	ctx := context.Background()
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: subnetBInfo.BlockchainID, // Message intended for SubnetB
		DestinationAddress:      common.HexToAddress("0x1111111111111111111111111111111111111111"),
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: fundedAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	log.Info(
		"Sending Teleporter transaction on source chain",
		"destinationBlockchainID", subnetBInfo.BlockchainID,
	)

	receipt, _ := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		subnetAInfo,
		subnetBInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	if network.SupportsIndependentRelaying() {
		//
		// Try to relay the message to subnet C, should fail
		//
		network.RelayMessage(ctx, receipt, subnetAInfo, subnetCInfo, false)
	} else {
		//
		// Wait for external relayer to properly deliver the message to subnet B
		//
		deliveryReceipt := network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)
		deliveryTx, isPending, err := subnetBInfo.RPCClient.TransactionByHash(ctx, deliveryReceipt.TxHash)
		Expect(err).Should(BeNil())
		Expect(isPending).Should(BeFalse())

		//
		// Take the successful delivery transaction, and use it to create a transaction that attempts to deliver
		// the same message to subnet C.
		//
		wrongChainDeliveryTx := createWrongChainDeliveryTransaction(ctx, deliveryTx, fundedKey, fundedAddress, subnetCInfo)
		utils.SendTransactionAndWaitForFailure(ctx, subnetCInfo, wrongChainDeliveryTx)
	}

	//
	// Check that the message was not received on the Subnet C
	//
	delivered, err := subnetCInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, expectedAtoCMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())
}

func createWrongChainDeliveryTransaction(
	ctx context.Context,
	deliveryTx *types.Transaction,
	fundedKey *ecdsa.PrivateKey,
	fundedAddress common.Address,
	destination interfaces.SubnetTestInfo,
) *types.Transaction {
	gasFeeCap, gasTipCap, nonce := utils.CalculateTxParams(ctx, destination, fundedAddress)
	unsignedTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:    destination.EVMChainID,
		Nonce:      nonce,
		To:         deliveryTx.To(),
		Gas:        deliveryTx.Gas(),
		GasFeeCap:  gasFeeCap,
		GasTipCap:  gasTipCap,
		Value:      big.NewInt(0),
		Data:       deliveryTx.Data(),
		AccessList: deliveryTx.AccessList(),
	})
	return utils.SignTransaction(unsignedTx, fundedKey, destination.EVMChainID)
}
