package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func InsufficientGasGinkgo() {
	insufficientGas(&network.LocalNetwork{})
}

func insufficientGas(network network.Network) {
	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	testAddress, testKey := network.GetTestAccountInfo()
	ctx := context.Background()

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetAInfo.ChainRPCClient)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetBInfo.ChainRPCClient)
	Expect(err).Should(BeNil())

	// Remove the balance from the funded address to testAddress
	balance, err := subnetAInfo.ChainRPCClient.BalanceAt(ctx, fundedAddress, nil)
	balance.Sub(balance, big.NewInt(1e14))
	Expect(err).Should(BeNil())

	testAddressBalance, err := subnetAInfo.ChainRPCClient.BalanceAt(ctx, testAddress, nil)
	Expect(err).Should(BeNil())
	moveBalanceTx := utils.CreateNativeTransferTransaction(ctx, subnetAInfo, fundedAddress, fundedKey, testAddress, balance)
	utils.SendTransactionAndWaitForAcceptance(ctx, subnetAInfo.ChainWSClient, subnetAInfo.ChainRPCClient, moveBalanceTx)
	testAddressBalanceAfter, err := subnetAInfo.ChainRPCClient.BalanceAt(ctx, testAddress, nil)
	Expect(err).Should(BeNil())
	log.Info("check test address balance")
	Expect(testAddressBalanceAfter).Should(Equal(testAddressBalance.Add(testAddressBalance, balance)))

	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
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

	opts := utils.CreateTransactorOpts(ctx, subnetAInfo, fundedAddress, fundedKey)

	log.Info("Sending cross chain message")

	// Send a transaction to the Teleporter contract
	_, err = subnetATeleporterMessenger.SendCrossChainMessage(opts, sendCrossChainMessageInput)
	Expect(err).ShouldNot(BeNil())

	// transfer balance back to funded address
	moveBalanceTx2 := utils.CreateNativeTransferTransaction(ctx, subnetAInfo, testAddress, testKey, fundedAddress, balance)
	utils.SendTransactionAndWaitForAcceptance(ctx, subnetAInfo.ChainWSClient, subnetAInfo.ChainRPCClient, moveBalanceTx2)

	// Retry message execution
	receipt, messageID := utils.SendCrossChainMessageAndWaitForAcceptance(ctx, subnetAInfo, subnetBInfo, sendCrossChainMessageInput, fundedAddress, fundedKey, subnetATeleporterMessenger)
	// relay message from SubnetA to SubnetB
	network.RelayMessage(ctx, receipt.BlockHash, receipt.BlockNumber, subnetAInfo, subnetBInfo)

	// Check Teleporter message received on the destination
	delivered, err := subnetBTeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())
}
