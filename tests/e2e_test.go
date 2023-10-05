// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"os"
	"testing"

	deploymentUtils "github.com/ava-labs/teleporter/contract-deployment/utils"
	testUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

const (
	teleporterByteCodeFile = "./contracts/out/TeleporterMessenger.sol/TeleporterMessenger.json"
	warpGenesisFile        = "./tests/utils/warp-genesis.json"
)

func TestE2E(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	format.UseStringerRepresentation = true
	ginkgo.RunSpecs(t, "Teleporter e2e test")
}

// Define the Teleporter before and after suite functions.
var _ = ginkgo.BeforeSuite(func() {
	testUtils.SetupNetwork(warpGenesisFile)
	// Generate the Teleporter deployment values
	teleporterDeployerTransaction, teleporterDeployerAddress, teleporterContractAddress, err := deploymentUtils.ConstructKeylessTransaction(teleporterByteCodeFile, false)
	Expect(err).Should(BeNil())

	testUtils.DeployTeleporterContracts(teleporterDeployerTransaction, teleporterDeployerAddress, teleporterContractAddress)
	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(testUtils.TearDownNetwork)

var _ = ginkgo.Describe("[Teleporter integration tests]", func() {
	ginkgo.It("Send a message from Subnet A to Subnet B", BasicOneWaySend)
})

var _ = ginkgo.Describe("[NativeTransfer two-way send]", ginkgo.Ordered, func() {
	const (
		valueToSend = int64(1000_000_000_000_000_000)
		valueToReturn = valueToSend/4
	)
	var (
		ctx = context.Background()
		err error
	)

	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	ginkgo.It("Deploy Contracts on chains A and B", ginkgo.Label("NativeTransfer", "DeployContracs"), func() {
		// Deploy Native Token bridge on the two subnets
		nativeTokenBridgeDeployerPK, err = crypto.HexToECDSA(bridgeDeployerKeyStr)
		Expect(err).Should(BeNil())
		nativeTokenBridgeContractAddress, err = deploymentUtils.DeriveEVMContractAddress(nativeTokenBridgeDeployer, 0)
		Expect(err).Should(BeNil())
		fmt.Println("Native Token Bridge Contract Address: ", nativeTokenBridgeContractAddress.Hex())


		nativeTokenSourceBytecode, err := deploymentUtils.ExtractByteCode(NativeTokenSourceByteCodeFile)
		Expect(err).Should(BeNil())
		chainATransactor, err := bind.NewKeyedTransactorWithChainID(nativeTokenBridgeDeployerPK, chainAIDInt)
		Expect(err).Should(BeNil())
		nativeTokenSourceAbi, err := abi.NativeTokenSourceMetaData.GetAbi()
		Expect(err).Should(BeNil())
		_, _, _, err = bind.DeployContract(chainATransactor, *nativeTokenSourceAbi, nativeTokenSourceBytecode, chainARPCClient, teleporterContractAddress, blockchainIDB, nativeTokenBridgeContractAddress)
		Expect(err).Should(BeNil())

		nativeTokenDestinationBytecode, err := deploymentUtils.ExtractByteCode(NativeTokenDestinationByteCodeFile)
		Expect(err).Should(BeNil())
		chainBTransactor, err := bind.NewKeyedTransactorWithChainID(nativeTokenBridgeDeployerPK, chainBIDInt)
		Expect(err).Should(BeNil())
		nativeTokenDestinationAbi, err := abi.NativeTokenDestinationMetaData.GetAbi()
		Expect(err).Should(BeNil())
		_, _, _, err = bind.DeployContract(chainBTransactor, *nativeTokenDestinationAbi, nativeTokenDestinationBytecode, chainBRPCClient, teleporterContractAddress, blockchainIDA, nativeTokenBridgeContractAddress, common.Big0)
		Expect(err).Should(BeNil())


		time.Sleep(5 * time.Second)
		bridgeCodeA, err := chainARPCClient.CodeAt(ctx, nativeTokenBridgeContractAddress, nil)
		Expect(err).Should(BeNil())
		Expect(len(bridgeCodeA)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode

		bridgeCodeB, err := chainBRPCClient.CodeAt(ctx, nativeTokenBridgeContractAddress, nil)
		Expect(err).Should(BeNil())
		Expect(len(bridgeCodeB)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode

		log.Info("Finished deploying Bridge contracts")
	})

	ginkgo.It("Transfer tokens from A to B", ginkgo.Label("NativeTransfer", "Lock and Mint"), func() {
		nativeTokenSource, err := abi.NewNativeTokenSource(nativeTokenBridgeContractAddress, chainARPCClient)
		Expect(err).Should(BeNil())
		transactor, err := bind.NewKeyedTransactorWithChainID(nativeTokenBridgeDeployerPK, chainAIDInt)
		Expect(err).Should(BeNil())
		transactor.Value = big.NewInt(valueToSend)

		subA, err := chainAWSClient.SubscribeNewHead(ctx, newHeadsA)
		Expect(err).Should(BeNil())
		defer subA.Unsubscribe()

		tx, err := nativeTokenSource.TransferToDestination(transactor, tokenReceiverAddress, tokenReceiverAddress, big.NewInt(0))
		Expect(err).Should(BeNil())
		log.Info("Sent TransferToDestination transaction on source chain", "destinationChainID", blockchainIDB, "txHash", tx.Hash().Hex())

		// Sleep to ensure the new block is published to the subscriber
		time.Sleep(5 * time.Second)
		receipt, err := chainARPCClient.TransactionReceipt(ctx, tx.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	})

	ginkgo.It("Relay message to destination", ginkgo.Label("NativeTransfer", "RelayMessage A -> B"), func() {
		// Get the latest block from Subnet A, and retrieve the warp message from the logs
		log.Info("Waiting for new block confirmation")
		newHeadA := <-newHeadsA
		blockHashA := newHeadA.Hash()

		log.Info("Fetching relevant warp logs from the newly produced block")
		logs, err := chainARPCClient.FilterLogs(ctx, interfaces.FilterQuery{
			BlockHash: &blockHashA,
			Addresses: []common.Address{warp.Module.Address},
		})
		Expect(err).Should(BeNil())
		Expect(len(logs)).Should(Equal(1))

		// Check for relevant warp log from subscription and ensure that it matches
		// the log extracted from the last block.
		txLog := logs[0]
		log.Info("Parsing logData as unsigned warp message")
		unsignedMsg, err := avalancheWarp.ParseUnsignedMessage(txLog.Data)
		Expect(err).Should(BeNil())

		// Set local variables for the duration of the test
		unsignedWarpMessageID := unsignedMsg.ID()
		unsignedWarpMsg := unsignedMsg
		log.Info("Parsed unsignedWarpMsg", "unsignedWarpMessageID", unsignedWarpMessageID, "unsignedWarpMessage", unsignedWarpMsg)

		// Loop over each client on chain A to ensure they all have time to accept the block.
		// Note: if we did not confirm this here, the next stage could be racy since it assumes every node
		// has accepted the block.
		waitForAllValidatorsToAcceptBlock(ctx, chainANodeURIs, blockchainIDA, newHeadA.Number.Uint64())

		// Get the aggregate signature for the Warp message
		log.Info("Fetching aggregate signature from the source chain validators")
		warpClient, err := warpBackend.NewClient(chainANodeURIs[0], blockchainIDA.String())
		Expect(err).Should(BeNil())
		signedWarpMessageBytes, err := warpClient.GetAggregateSignature(ctx, unsignedWarpMessageID, params.WarpQuorumDenominator)
		Expect(err).Should(BeNil())

		// Check starting balance is 0
		bal, err := chainBRPCClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Int64()).Should(Equal(common.Big0.Int64()))

		signedTxB := constructAndSendTransaction(
			ctx,
			signedWarpMessageBytes,
			big.NewInt(1),
			teleporterContractAddress,
			fundedAddress,
			fundedKey,
			chainBRPCClient,
			chainBIDInt,
		)

		// Sleep to ensure the new block is published to the subscriber
		time.Sleep(5 * time.Second)
		receipt, err := chainBRPCClient.TransactionReceipt(ctx, signedTxB.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

		sendCrossChainMessageLog := receipt.Logs[2]
		var event SendCrossChainMessageEvent
		err = teleporter.EVMTeleporterContractABI.UnpackIntoInterface(&event, "SendCrossChainMessage", sendCrossChainMessageLog.Data)
		Expect(err).Should(BeNil())

		bal, err = chainBRPCClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Int64()).Should(Equal(valueToSend))
	})

	ginkgo.It("Transfer tokens from B to A", ginkgo.Label("NativeTransfer", "Burn and Unlock"), func() {
		nativeTokenDestination, err := abi.NewNativeTokenDestination(nativeTokenBridgeContractAddress, chainBRPCClient)
		Expect(err).Should(BeNil())
		transactor, err := bind.NewKeyedTransactorWithChainID(nativeTokenBridgeDeployerPK, chainBIDInt)
		Expect(err).Should(BeNil())
		transactor.Value = big.NewInt(valueToReturn)

		subB, err := chainBWSClient.SubscribeNewHead(ctx, newHeadsB)
		Expect(err).Should(BeNil())
		defer subB.Unsubscribe()

		tx, err := nativeTokenDestination.TransferToSource(transactor, tokenReceiverAddress, tokenReceiverAddress, big.NewInt(0))
		Expect(err).Should(BeNil())
		log.Info("Sent TransferToSource transaction on destination chain", "sourceChainID", blockchainIDA, "txHash", tx.Hash().Hex())

		// Sleep to ensure the new block is published to the subscriber
		time.Sleep(5 * time.Second)
		receipt, err := chainBRPCClient.TransactionReceipt(ctx, tx.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	})

	ginkgo.It("Relay message to source", ginkgo.Label("NativeTransfer", "RelayMessage B -> A"), func() {
		// Get the latest block from Subnet A, and retrieve the warp message from the logs
		log.Info("Waiting for new block confirmation")
		newHeadB := <-newHeadsB
		blockHashB := newHeadB.Hash()

		log.Info("Fetching relevant warp logs from the newly produced block")
		logs, err := chainBRPCClient.FilterLogs(ctx, interfaces.FilterQuery{
			BlockHash: &blockHashB,
			Addresses: []common.Address{warp.Module.Address},
		})
		Expect(err).Should(BeNil())
		Expect(len(logs)).Should(Equal(1))

		// Check for relevant warp log from subscription and ensure that it matches
		// the log extracted from the last block.
		txLog := logs[0]
		log.Info("Parsing logData as unsigned warp message")
		unsignedMsg, err := avalancheWarp.ParseUnsignedMessage(txLog.Data)
		Expect(err).Should(BeNil())

		// Set local variables for the duration of the test
		unsignedWarpMessageID := unsignedMsg.ID()
		unsignedWarpMsg := unsignedMsg
		log.Info("Parsed unsignedWarpMsg", "unsignedWarpMessageID", unsignedWarpMessageID, "unsignedWarpMessage", unsignedWarpMsg)

		// Loop over each client on chain A to ensure they all have time to accept the block.
		// Note: if we did not confirm this here, the next stage could be racy since it assumes every node
		// has accepted the block.
		waitForAllValidatorsToAcceptBlock(ctx, chainBNodeURIs, blockchainIDB, newHeadB.Number.Uint64())

		// Get the aggregate signature for the Warp message
		log.Info("Fetching aggregate signature from the source chain validators")
		warpClient, err := warpBackend.NewClient(chainBNodeURIs[0], blockchainIDB.String())
		Expect(err).Should(BeNil())
		signedWarpMessageBytes, err := warpClient.GetAggregateSignature(ctx, unsignedWarpMessageID, params.WarpQuorumDenominator)
		Expect(err).Should(BeNil())

		// Check starting balance is 0
		bal, err := chainARPCClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Int64()).Should(Equal(common.Big0.Int64()))

		signedTxA := constructAndSendTransaction(
			ctx,
			signedWarpMessageBytes,
			big.NewInt(1),
			teleporterContractAddress,
			fundedAddress,
			fundedKey,
			chainARPCClient,
			chainAIDInt,
		)

		// Sleep to ensure the new block is published to the subscriber
		time.Sleep(5 * time.Second)
		receipt, err := chainARPCClient.TransactionReceipt(ctx, signedTxA.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

		sendCrossChainMessageLog := receipt.Logs[2]
		var event SendCrossChainMessageEvent
		err = teleporter.EVMTeleporterContractABI.UnpackIntoInterface(&event, "SendCrossChainMessage", sendCrossChainMessageLog.Data)
		Expect(err).Should(BeNil())

		bal, err = chainARPCClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Int64()).Should(Equal(valueToReturn))
	})

})