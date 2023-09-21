// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/avalanche-network-runner/rpcpb"
	"github.com/ava-labs/avalanchego/ids"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/awm-relayer/messages/teleporter"
	relayerEvm "github.com/ava-labs/awm-relayer/vms/evm"
	"github.com/ava-labs/coreth/rpc"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ava-labs/subnet-evm/params"
	"github.com/ava-labs/subnet-evm/plugin/evm"
	"github.com/ava-labs/subnet-evm/tests/utils/runner"
	predicateutils "github.com/ava-labs/subnet-evm/utils/predicate"
	warpBackend "github.com/ava-labs/subnet-evm/warp"
	warpPayload "github.com/ava-labs/subnet-evm/warp/payload"

	"github.com/ava-labs/subnet-evm/x/warp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	fundedKeyStr = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
)

var (
	// Vars used in the suite setup
	warpGenesisFile            = fmt.Sprintf("%s/tests/warp-genesis.json", os.Getenv("REPO_ROOT"))
	teleporterByteCodeFile     = fmt.Sprintf("%s/contracts/out/TeleporterMessenger.sol/TeleporterMessenger.json", os.Getenv("REPO_ROOT"))
	anrConfig                  = runner.NewDefaultANRConfig()
	manager                    = runner.NewNetworkManager(anrConfig)
	warpChainConfigPath        string
	subnetIDs                  []ids.ID
	subnetA, subnetB           ids.ID
	chainARPCURI, chainBRPCURI string

	// Vars used in the tests. Any variables that are set in the process1 function of SynchronizedBeforeSuite should
	// be set again in the allProcesses function of SynchronizedAfterSuite so that the tests can access them in their
	// respective parallel process
	// TODONOW: provide each suite with its own funded address
	fundedAddress             = common.HexToAddress("0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC")
	teleporterContractAddress common.Address
	teleporterMessage         = teleporter.TeleporterMessage{
		MessageID:               big.NewInt(1),
		SenderAddress:           fundedAddress,
		DestinationAddress:      fundedAddress,
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Receipts:                []teleporter.TeleporterMessageReceipt{},
		Message:                 []byte{1, 2, 3, 4},
	}
	blockchainIDA, blockchainIDB     ids.ID
	chainANodeURIs, chainBNodeURIs   []string
	fundedKey                        *ecdsa.PrivateKey
	chainAWSClient, chainBWSClient   ethclient.Client
	chainARPCClient, chainBRPCClient ethclient.Client
	chainAIDInt, chainBIDInt         *big.Int
	newHeadsA, newHeadsB             chan *types.Header
)

type SubnetStruct struct {
	SubnetID      ids.ID   `json:"subnetIDs"`
	ValidatorURIs []string `json:"validatorURIs"`
	BlockchainID  ids.ID   `json:"blockchainID"`
}

type SetupStruct struct {
	Subnets []SubnetStruct `json:"subnets"`
}

func TestE2E(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Teleporter e2e test")
}

// BeforeSuite starts the default network and adds 10 new nodes as validators with BLS keys
// registered on the P-Chain.
// Adds two disjoint sets of 5 of the new validator nodes to validate two new subnets with a
// a single Subnet-EVM blockchain.
var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {
	log.Info("Running ginkgo before suite")
	ctx := context.Background()
	var err error

	// Name 10 new validators (which should have BLS key registered)
	subnetANodeNames := []string{}
	subnetBNodeNames := []string{}
	for i := 1; i <= 10; i++ {
		n := fmt.Sprintf("node%d-bls", i)
		if i <= 5 {
			subnetANodeNames = append(subnetANodeNames, n)
		} else {
			subnetBNodeNames = append(subnetBNodeNames, n)
		}
	}
	f, err := os.CreateTemp(os.TempDir(), "config.json")
	Expect(err).Should(BeNil())
	_, err = f.Write([]byte(`{"warp-api-enabled": true}`))
	Expect(err).Should(BeNil())
	warpChainConfigPath = f.Name()

	// Make sure that the warp genesis file exists
	_, err = os.Stat(warpGenesisFile)
	Expect(err).Should(BeNil())

	// Construct the network using the avalanche-network-runner
	_, err = manager.StartDefaultNetwork(ctx)
	Expect(err).Should(BeNil())
	err = manager.SetupNetwork(
		ctx,
		anrConfig.AvalancheGoExecPath,
		[]*rpcpb.BlockchainSpec{
			{
				VmName:      evm.IDStr,
				Genesis:     warpGenesisFile,
				ChainConfig: warpChainConfigPath,
				SubnetSpec: &rpcpb.SubnetSpec{
					SubnetConfig: "",
					Participants: subnetANodeNames,
				},
			},
			{
				VmName:      evm.IDStr,
				Genesis:     warpGenesisFile,
				ChainConfig: warpChainConfigPath,
				SubnetSpec: &rpcpb.SubnetSpec{
					SubnetConfig: "",
					Participants: subnetBNodeNames,
				},
			},
		},
	)
	Expect(err).Should(BeNil())

	// Issue transactions to activate the proposerVM fork on the receiving chain
	fundedKey, err = crypto.HexToECDSA(fundedKeyStr)
	Expect(err).Should(BeNil())
	setUpProposerVm(ctx, fundedKey, manager, 0)
	setUpProposerVm(ctx, fundedKey, manager, 1)

	// Set up subnet URIs
	subnetIDs = manager.GetSubnets()
	Expect(len(subnetIDs)).Should(Equal(2))

	subnetA = subnetIDs[0]
	subnetADetails, ok := manager.GetSubnet(subnetA)
	Expect(ok).Should(BeTrue())
	Expect(len(subnetADetails.ValidatorURIs)).Should(Equal(5))
	blockchainIDA = subnetADetails.BlockchainID
	chainANodeURIs = append(chainANodeURIs, subnetADetails.ValidatorURIs...)

	subnetB = subnetIDs[1]
	subnetBDetails, ok := manager.GetSubnet(subnetB)
	Expect(ok).Should(BeTrue())
	Expect(len(subnetBDetails.ValidatorURIs)).Should(Equal(5))
	blockchainIDB = subnetBDetails.BlockchainID
	chainBNodeURIs = append(chainBNodeURIs, subnetBDetails.ValidatorURIs...)

	log.Info(
		"Created URIs for subnets",
		"ChainAURIs", chainANodeURIs,
		"ChainBURIs", chainBNodeURIs,
		"blockchainIDA", blockchainIDA,
		"blockchainIDB", blockchainIDB,
	)

	chainAWSURI := httpToWebsocketURI(chainANodeURIs[0], blockchainIDA.String())
	chainARPCURI = httpToRPCURI(chainANodeURIs[0], blockchainIDA.String())
	log.Info("Creating ethclient for blockchainA", "wsURI", chainAWSURI, "rpcURL, chainARPCURI")
	chainAWSClient, err = ethclient.Dial(chainAWSURI)
	Expect(err).Should(BeNil())
	chainARPCClient, err = ethclient.Dial(chainARPCURI)
	Expect(err).Should(BeNil())

	chainAIDInt, err = chainARPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	chainBWSURI := httpToWebsocketURI(chainBNodeURIs[0], blockchainIDB.String())
	chainBRPCURI = httpToRPCURI(chainBNodeURIs[0], blockchainIDB.String())
	log.Info("Creating ethclient for blockchainB", "wsURI", chainBWSURI)
	chainBWSClient, err = ethclient.Dial(chainBWSURI)
	Expect(err).Should(BeNil())
	chainBRPCClient, err = ethclient.Dial(chainBRPCURI)
	Expect(err).Should(BeNil())

	chainBIDInt, err = chainBRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	newHeadsA = make(chan *types.Header, 10)
	newHeadsB = make(chan *types.Header, 10)

	log.Info("Finished setting up e2e test subnet variables")

	log.Info("Deploying Teleporter contract to subnets")
	// Generate the Teleporter deployment values
	var (
		teleporterDeployerAddress     common.Address
		teleporterDeployerTransaction []byte
	)

	teleporterDeployerTransaction, teleporterDeployerAddress, teleporterContractAddress, err = ConstructKeylessTransaction(teleporterByteCodeFile, false)
	Expect(err).Should(BeNil())

	nonceA, err := chainARPCClient.NonceAt(ctx, fundedAddress, nil)
	Expect(err).Should(BeNil())

	nonceB, err := chainBRPCClient.NonceAt(ctx, fundedAddress, nil)
	Expect(err).Should(BeNil())

	gasTipCapA, err := chainARPCClient.SuggestGasTipCap(context.Background())
	Expect(err).Should(BeNil())
	gasTipCapB, err := chainBRPCClient.SuggestGasTipCap(context.Background())
	Expect(err).Should(BeNil())

	baseFeeA, err := chainARPCClient.EstimateBaseFee(context.Background())
	Expect(err).Should(BeNil())
	gasFeeCapA := baseFeeA.Mul(baseFeeA, big.NewInt(relayerEvm.BaseFeeFactor))
	gasFeeCapA.Add(gasFeeCapA, big.NewInt(relayerEvm.MaxPriorityFeePerGas))

	baseFeeB, err := chainBRPCClient.EstimateBaseFee(context.Background())
	Expect(err).Should(BeNil())
	gasFeeCapB := baseFeeB.Mul(baseFeeB, big.NewInt(relayerEvm.BaseFeeFactor))
	gasFeeCapB.Add(gasFeeCapB, big.NewInt(relayerEvm.MaxPriorityFeePerGas))

	// Fund the deployer address
	{
		value := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)) // 10eth
		txA := types.NewTx(&types.DynamicFeeTx{
			ChainID:   chainAIDInt,
			Nonce:     nonceA,
			To:        &teleporterDeployerAddress,
			Gas:       defaultTeleporterMessageGas,
			GasFeeCap: gasFeeCapA,
			GasTipCap: gasTipCapA,
			Value:     value,
		})
		txSignerA := types.LatestSignerForChainID(chainAIDInt)
		triggerTxA, err := types.SignTx(txA, txSignerA, fundedKey)
		Expect(err).Should(BeNil())
		err = chainARPCClient.SendTransaction(ctx, triggerTxA)
		Expect(err).Should(BeNil())
		time.Sleep(5 * time.Second)
		receipt, err := chainARPCClient.TransactionReceipt(ctx, triggerTxA.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	}
	{
		value := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)) // 10eth
		txB := types.NewTx(&types.DynamicFeeTx{
			ChainID:   chainBIDInt,
			Nonce:     nonceB,
			To:        &teleporterDeployerAddress,
			Gas:       defaultTeleporterMessageGas,
			GasFeeCap: gasFeeCapB,
			GasTipCap: gasTipCapB,
			Value:     value,
		})
		txSignerB := types.LatestSignerForChainID(chainBIDInt)
		triggerTxB, err := types.SignTx(txB, txSignerB, fundedKey)
		Expect(err).Should(BeNil())
		err = chainBRPCClient.SendTransaction(ctx, triggerTxB)
		Expect(err).Should(BeNil())
		time.Sleep(5 * time.Second)
		receipt, err := chainBRPCClient.TransactionReceipt(ctx, triggerTxB.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	}
	// Deploy Teleporter on the two subnets
	{
		rpcClient, err := rpc.DialContext(ctx, chainARPCURI)
		Expect(err).Should(BeNil())
		err = rpcClient.CallContext(ctx, nil, "eth_sendRawTransaction", hexutil.Encode(teleporterDeployerTransaction))
		Expect(err).Should(BeNil())
		time.Sleep(5 * time.Second)
		teleporterCode, err := chainARPCClient.CodeAt(ctx, teleporterContractAddress, nil)
		Expect(err).Should(BeNil())
		Expect(len(teleporterCode)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
	}
	{
		rpcClient, err := rpc.DialContext(ctx, chainBRPCURI)
		Expect(err).Should(BeNil())
		err = rpcClient.CallContext(ctx, nil, "eth_sendRawTransaction", hexutil.Encode(teleporterDeployerTransaction))
		Expect(err).Should(BeNil())
		time.Sleep(5 * time.Second)
		teleporterCode, err := chainBRPCClient.CodeAt(ctx, teleporterContractAddress, nil)
		Expect(err).Should(BeNil())
		Expect(len(teleporterCode)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
	}
	log.Info("Finished deploying Teleporter contracts")

	log.Info("Set up ginkgo before suite")

	setupStruct := SetupStruct{
		Subnets: []SubnetStruct{
			{
				SubnetID:      subnetA,
				ValidatorURIs: chainANodeURIs,
				BlockchainID:  blockchainIDA,
			},
			{
				SubnetID:      subnetB,
				ValidatorURIs: chainBNodeURIs,
				BlockchainID:  blockchainIDB,
			},
		},
	}
	setupBytes, err := json.Marshal(setupStruct)
	Expect(err).Should(BeNil())

	return setupBytes
},
	// Re-set all global variables that are referenced in the tests and set in the above function
	func(setupBytes []byte) {
		var setupStruct SetupStruct
		err := json.Unmarshal(setupBytes, &setupStruct)

		_, _, teleporterContractAddress, err = ConstructKeylessTransaction(teleporterByteCodeFile, false)
		Expect(err).Should(BeNil())

		// Set up subnet URIs
		subnetA = setupStruct.Subnets[0].SubnetID
		Expect(len(setupStruct.Subnets[0].ValidatorURIs)).Should(Equal(5))
		blockchainIDA = setupStruct.Subnets[0].BlockchainID
		chainANodeURIs = append(chainANodeURIs, setupStruct.Subnets[0].ValidatorURIs...)

		subnetB = setupStruct.Subnets[1].SubnetID
		Expect(len(setupStruct.Subnets[1].ValidatorURIs)).Should(Equal(5))
		blockchainIDB = setupStruct.Subnets[1].BlockchainID
		chainBNodeURIs = append(chainBNodeURIs, setupStruct.Subnets[1].ValidatorURIs...)

		chainAWSURI := httpToWebsocketURI(chainANodeURIs[0], blockchainIDA.String())
		chainARPCURI = httpToRPCURI(chainANodeURIs[0], blockchainIDA.String())
		log.Info("Creating ethclient for blockchainA", "wsURI", chainAWSURI, "rpcURL, chainARPCURI")
		chainAWSClient, err = ethclient.Dial(chainAWSURI)
		Expect(err).Should(BeNil())
		chainARPCClient, err = ethclient.Dial(chainARPCURI)
		Expect(err).Should(BeNil())

		chainAIDInt, err = chainARPCClient.ChainID(context.Background())
		Expect(err).Should(BeNil())

		chainBWSURI := httpToWebsocketURI(chainBNodeURIs[0], blockchainIDB.String())
		chainBRPCURI = httpToRPCURI(chainBNodeURIs[0], blockchainIDB.String())
		log.Info("Creating ethclient for blockchainB", "wsURI", chainBWSURI)
		chainBWSClient, err = ethclient.Dial(chainBWSURI)
		Expect(err).Should(BeNil())
		chainBRPCClient, err = ethclient.Dial(chainBRPCURI)
		Expect(err).Should(BeNil())

		chainBIDInt, err = chainBRPCClient.ChainID(context.Background())
		Expect(err).Should(BeNil())

		newHeadsA = make(chan *types.Header, 10)
		newHeadsB = make(chan *types.Header, 10)

		// Create and fund the address to be used by the suite
		// TODONOW:
		// How do we do this? We cant generate the key on the fly, since we need to fund it using the genesis key,
		// which we can't do without the nonce.
		// Instead, we could cap the number of parallel runners (this is reasonable) and fund that number of addresses
		// in the genesis. We just need a way to uniquely assign addresses to each runner. We have to do that here, since
		// the test suites don't necessarily map one-to-one with parallel runners. This is an issue if the number of tests
		// exceeds the number of parallel runners
		fundedKey, err = crypto.HexToECDSA(fundedKeyStr)
		Expect(err).Should(BeNil())

	})

var _ = ginkgo.SynchronizedAfterSuite(
	func() {},
	func() {
		log.Info("Running ginkgo after suite")
		Expect(manager).ShouldNot(BeNil())
		Expect(manager.TeardownNetwork()).Should(BeNil())
		Expect(os.Remove(warpChainConfigPath)).Should(BeNil())
	})

// Ginkgo describe node that acts as a container for the teleporter e2e tests. This test suite
// will run through the following steps in order:
// 1. Send a transaction to the Teleporter contract on Subnet A
// 2. Aggregate signatures and send the Warp message to Subnet B
// 3. Verify receipt of the message on Subnet B
var _ = ginkgo.Describe("[Teleporter one way send]", ginkgo.Ordered, func() {
	var (
		receivedWarpMessage *avalancheWarp.Message
		payload             []byte
		teleporterMessageID *big.Int
	)

	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	ginkgo.It("Send Message from A to B", ginkgo.Label("Teleporter", "SendTeleporter"), func() {
		ctx := context.Background()

		nonceA, err := chainARPCClient.NonceAt(ctx, fundedAddress, nil)
		Expect(err).Should(BeNil())

		nonceB, err := chainBRPCClient.NonceAt(ctx, fundedAddress, nil)
		Expect(err).Should(BeNil())

		log.Info("Packing teleporter message", "nonceA", nonceA, "nonceB", nonceB)
		payload, err = teleporter.PackSendCrossChainMessageEvent(common.Hash(blockchainIDB), teleporterMessage)
		Expect(err).Should(BeNil())

		data, err := teleporter.EVMTeleporterContractABI.Pack(
			"sendCrossChainMessage",
			TeleporterMessageInput{
				DestinationChainID: blockchainIDB,
				DestinationAddress: fundedAddress,
				FeeInfo: FeeInfo{
					ContractAddress: fundedAddress,
					Amount:          big.NewInt(0),
				},
				RequiredGasLimit:        big.NewInt(1),
				AllowedRelayerAddresses: []common.Address{},
				Message:                 []byte{1, 2, 3, 4},
			},
		)
		Expect(err).Should(BeNil())

		// Send a transaction to the Teleporter contract
		tx := newTestTeleporterMessage(chainAIDInt, teleporterContractAddress, nonceA, data)

		txSigner := types.LatestSignerForChainID(chainAIDInt)
		signedTx, err := types.SignTx(tx, txSigner, fundedKey)
		Expect(err).Should(BeNil())

		subA, err := chainAWSClient.SubscribeNewHead(ctx, newHeadsA)
		Expect(err).Should(BeNil())
		defer subA.Unsubscribe()

		log.Info("Sending Teleporter transaction on source chain", "destinationChainID", blockchainIDB, "txHash", signedTx.Hash())
		err = chainARPCClient.SendTransaction(ctx, signedTx)
		Expect(err).Should(BeNil())

		// Sleep to ensure the new block is published to the subscriber
		time.Sleep(5 * time.Second)
		receipt, err := chainARPCClient.TransactionReceipt(ctx, signedTx.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	})

	ginkgo.It("Relay message to destination", ginkgo.Label("Teleporter", "RelayMessage"), func() {
		ctx := context.Background()

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
		for i, uri := range chainANodeURIs {
			chainAWSURI := httpToWebsocketURI(uri, blockchainIDA.String())
			log.Info("Creating ethclient for blockchainA", "wsURI", chainAWSURI)
			client, err := ethclient.Dial(chainAWSURI)
			Expect(err).Should(BeNil())

			// Loop until each node has advanced to >= the height of the block that emitted the warp log
			for {
				block, err := client.BlockByNumber(ctx, nil)
				Expect(err).Should(BeNil())
				if block.NumberU64() >= newHeadA.Number.Uint64() {
					log.Info("client accepted the block containing SendWarpMessage", "client", i, "height", block.NumberU64())
					break
				}
			}
		}

		// Get the aggregate signature for the Warp message
		log.Info("Fetching aggregate signature from the source chain validators")
		warpClient, err := warpBackend.NewWarpClient(chainANodeURIs[0], blockchainIDA.String())
		Expect(err).Should(BeNil())
		signedWarpMessageBytes, err := warpClient.GetAggregateSignature(ctx, unsignedWarpMessageID, params.WarpQuorumDenominator)
		Expect(err).Should(BeNil())

		// Construct the transaction to send the Warp message to the destination chain
		log.Info("Constructing transaction for the destination chain")
		signedMessage, err := avalancheWarp.ParseMessage(signedWarpMessageBytes)
		Expect(err).Should(BeNil())

		numSigners, err := signedMessage.Signature.NumSigners()
		Expect(err).Should(BeNil())

		gasLimit, err := teleporter.CalculateReceiveMessageGasLimit(numSigners, teleporterMessage.RequiredGasLimit)
		Expect(err).Should(BeNil())

		callData, err := teleporter.EVMTeleporterContractABI.Pack("receiveCrossChainMessage", fundedAddress)
		Expect(err).Should(BeNil())

		baseFee, err := chainBRPCClient.EstimateBaseFee(ctx)
		Expect(err).Should(BeNil())

		gasTipCap, err := chainBRPCClient.SuggestGasTipCap(ctx)
		Expect(err).Should(BeNil())

		nonce, err := chainBRPCClient.NonceAt(ctx, fundedAddress, nil)
		Expect(err).Should(BeNil())

		gasFeeCap := baseFee.Mul(baseFee, big.NewInt(2))
		gasFeeCap.Add(gasFeeCap, big.NewInt(2500000000))
		destinationTx := predicateutils.NewPredicateTx(
			chainBIDInt,
			nonce,
			&teleporterContractAddress,
			gasLimit,
			gasFeeCap,
			gasTipCap,
			big.NewInt(0),
			callData,
			types.AccessList{},
			warp.ContractAddress,
			signedMessage.Bytes(),
		)

		// Sign and send the transaction on the destination chain
		signer := types.LatestSignerForChainID(chainBIDInt)
		signedTxB, err := types.SignTx(destinationTx, signer, fundedKey)
		Expect(err).Should(BeNil())

		log.Info("Subscribing to new heads on destination chain")
		subB, err := chainBWSClient.SubscribeNewHead(ctx, newHeadsB)
		Expect(err).Should(BeNil())
		defer subB.Unsubscribe()

		log.Info("Sending transaction to destination chain")
		err = chainBRPCClient.SendTransaction(context.Background(), signedTxB)
		Expect(err).Should(BeNil())

		// Sleep to ensure the new block is published to the subscriber
		time.Sleep(5 * time.Second)
		receipt, err := chainBRPCClient.TransactionReceipt(ctx, signedTxB.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	})

	ginkgo.It("Receive message on Subnet B", ginkgo.Label("Teleporter", "ReceiveTeleporter"), func() {
		ctx := context.Background()

		// Get the latest block from Subnet B
		log.Info("Waiting for new block confirmation")
		newHeadB := <-newHeadsB
		log.Info("Received new head", "height", newHeadB.Number.Uint64())
		blockHashB := newHeadB.Hash()
		block, err := chainBRPCClient.BlockByHash(ctx, blockHashB)
		Expect(err).Should(BeNil())
		log.Info(
			"Got block",
			"blockHash", blockHashB,
			"blockNumber", block.NumberU64(),
			"transactions", block.Transactions(),
			"numTransactions", len(block.Transactions()),
			"block", block,
		)
		accessLists := block.Transactions()[0].AccessList()
		Expect(len(accessLists)).Should(Equal(1))
		Expect(accessLists[0].Address).Should(Equal(warp.Module.Address))

		// Check the transaction storage key has warp message we're expecting
		storageKeyHashes := accessLists[0].StorageKeys
		packedPredicate := predicateutils.HashSliceToBytes(storageKeyHashes)
		predicateBytes, err := predicateutils.UnpackPredicate(packedPredicate)
		Expect(err).Should(BeNil())
		receivedWarpMessage, err = avalancheWarp.ParseMessage(predicateBytes)
		Expect(err).Should(BeNil())

		// Check that the transaction has successful receipt status
		txHash := block.Transactions()[0].Hash()
		receipt, err := chainBRPCClient.TransactionReceipt(ctx, txHash)
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

		log.Info("Finished sending warp message, closing down output channel")

	})

	ginkgo.It("Validate Received Warp Message Values", ginkgo.Label("Teleporter", "VerifyWarp"), func() {
		Expect(receivedWarpMessage.SourceChainID).Should(Equal(blockchainIDA))
		addressedPayload, err := warpPayload.ParseAddressedPayload(receivedWarpMessage.Payload)
		Expect(err).Should(BeNil())

		receivedDestinationID, err := ids.ToID(addressedPayload.DestinationChainID.Bytes())
		Expect(err).Should(BeNil())
		Expect(receivedDestinationID).Should(Equal(blockchainIDB))
		Expect(addressedPayload.DestinationAddress).Should(Equal(teleporterContractAddress))
		Expect(addressedPayload.Payload).Should(Equal(payload))

		// Check that the teleporter message is correct
		receivedTeleporterMessage, err := teleporter.UnpackTeleporterMessage(addressedPayload.Payload)
		Expect(err).Should(BeNil())
		Expect(*receivedTeleporterMessage).Should(Equal(teleporterMessage))

		teleporterMessageID = receivedTeleporterMessage.MessageID
	})

	ginkgo.It("Check Teleporter Message Received", ginkgo.Label("Teleporter", "TeleporterMessageReceived"), func() {
		data, err := teleporter.PackMessageReceivedMessage(teleporter.MessageReceivedInput{
			OriginChainID: blockchainIDA,
			MessageID:     teleporterMessageID,
		})
		Expect(err).Should(BeNil())
		callMessage := interfaces.CallMsg{
			To:   &teleporterContractAddress,
			Data: data,
		}
		result, err := chainBRPCClient.CallContract(context.Background(), callMessage, nil)
		Expect(err).Should(BeNil())

		// check the contract call result
		delivered, err := teleporter.UnpackMessageReceivedResult(result)
		Expect(err).Should(BeNil())
		Expect(delivered).Should(BeTrue())
	})

})

// Replica test
// TODONOW: remove this
var _ = ginkgo.Describe("[Teleporter one way send replica]", ginkgo.Ordered, func() {
	var (
		receivedWarpMessage *avalancheWarp.Message
		payload             []byte
		teleporterMessageID *big.Int
	)

	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	ginkgo.It("Send Message from A to B", ginkgo.Label("Teleporter", "SendTeleporter"), func() {
		ctx := context.Background()

		nonceA, err := chainARPCClient.NonceAt(ctx, fundedAddress, nil)
		Expect(err).Should(BeNil())

		nonceB, err := chainBRPCClient.NonceAt(ctx, fundedAddress, nil)
		Expect(err).Should(BeNil())

		log.Info("Packing teleporter message", "nonceA", nonceA, "nonceB", nonceB)
		payload, err = teleporter.PackSendCrossChainMessageEvent(common.Hash(blockchainIDB), teleporterMessage)
		Expect(err).Should(BeNil())

		data, err := teleporter.EVMTeleporterContractABI.Pack(
			"sendCrossChainMessage",
			TeleporterMessageInput{
				DestinationChainID: blockchainIDB,
				DestinationAddress: fundedAddress,
				FeeInfo: FeeInfo{
					ContractAddress: fundedAddress,
					Amount:          big.NewInt(0),
				},
				RequiredGasLimit:        big.NewInt(1),
				AllowedRelayerAddresses: []common.Address{},
				Message:                 []byte{1, 2, 3, 4},
			},
		)
		Expect(err).Should(BeNil())

		// Send a transaction to the Teleporter contract
		tx := newTestTeleporterMessage(chainAIDInt, teleporterContractAddress, nonceA, data)

		txSigner := types.LatestSignerForChainID(chainAIDInt)
		signedTx, err := types.SignTx(tx, txSigner, fundedKey)
		Expect(err).Should(BeNil())

		subA, err := chainAWSClient.SubscribeNewHead(ctx, newHeadsA)
		Expect(err).Should(BeNil())
		defer subA.Unsubscribe()

		log.Info("Sending Teleporter transaction on source chain", "destinationChainID", blockchainIDB, "txHash", signedTx.Hash())
		err = chainARPCClient.SendTransaction(ctx, signedTx)
		Expect(err).Should(BeNil())

		// Sleep to ensure the new block is published to the subscriber
		time.Sleep(5 * time.Second)
		receipt, err := chainARPCClient.TransactionReceipt(ctx, signedTx.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	})

	ginkgo.It("Relay message to destination", ginkgo.Label("Teleporter", "RelayMessage"), func() {
		ctx := context.Background()

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
		for i, uri := range chainANodeURIs {
			chainAWSURI := httpToWebsocketURI(uri, blockchainIDA.String())
			log.Info("Creating ethclient for blockchainA", "wsURI", chainAWSURI)
			client, err := ethclient.Dial(chainAWSURI)
			Expect(err).Should(BeNil())

			// Loop until each node has advanced to >= the height of the block that emitted the warp log
			for {
				block, err := client.BlockByNumber(ctx, nil)
				Expect(err).Should(BeNil())
				if block.NumberU64() >= newHeadA.Number.Uint64() {
					log.Info("client accepted the block containing SendWarpMessage", "client", i, "height", block.NumberU64())
					break
				}
			}
		}

		// Get the aggregate signature for the Warp message
		log.Info("Fetching aggregate signature from the source chain validators")
		warpClient, err := warpBackend.NewWarpClient(chainANodeURIs[0], blockchainIDA.String())
		Expect(err).Should(BeNil())
		signedWarpMessageBytes, err := warpClient.GetAggregateSignature(ctx, unsignedWarpMessageID, params.WarpQuorumDenominator)
		Expect(err).Should(BeNil())

		// Construct the transaction to send the Warp message to the destination chain
		log.Info("Constructing transaction for the destination chain")
		signedMessage, err := avalancheWarp.ParseMessage(signedWarpMessageBytes)
		Expect(err).Should(BeNil())

		numSigners, err := signedMessage.Signature.NumSigners()
		Expect(err).Should(BeNil())

		gasLimit, err := teleporter.CalculateReceiveMessageGasLimit(numSigners, teleporterMessage.RequiredGasLimit)
		Expect(err).Should(BeNil())

		callData, err := teleporter.EVMTeleporterContractABI.Pack("receiveCrossChainMessage", fundedAddress)
		Expect(err).Should(BeNil())

		baseFee, err := chainBRPCClient.EstimateBaseFee(ctx)
		Expect(err).Should(BeNil())

		gasTipCap, err := chainBRPCClient.SuggestGasTipCap(ctx)
		Expect(err).Should(BeNil())

		nonce, err := chainBRPCClient.NonceAt(ctx, fundedAddress, nil)
		Expect(err).Should(BeNil())

		gasFeeCap := baseFee.Mul(baseFee, big.NewInt(2))
		gasFeeCap.Add(gasFeeCap, big.NewInt(2500000000))
		destinationTx := predicateutils.NewPredicateTx(
			chainBIDInt,
			nonce,
			&teleporterContractAddress,
			gasLimit,
			gasFeeCap,
			gasTipCap,
			big.NewInt(0),
			callData,
			types.AccessList{},
			warp.ContractAddress,
			signedMessage.Bytes(),
		)

		// Sign and send the transaction on the destination chain
		signer := types.LatestSignerForChainID(chainBIDInt)
		signedTxB, err := types.SignTx(destinationTx, signer, fundedKey)
		Expect(err).Should(BeNil())

		log.Info("Subscribing to new heads on destination chain")
		subB, err := chainBWSClient.SubscribeNewHead(ctx, newHeadsB)
		Expect(err).Should(BeNil())
		defer subB.Unsubscribe()

		log.Info("Sending transaction to destination chain")
		err = chainBRPCClient.SendTransaction(context.Background(), signedTxB)
		Expect(err).Should(BeNil())

		// Sleep to ensure the new block is published to the subscriber
		time.Sleep(5 * time.Second)
		receipt, err := chainBRPCClient.TransactionReceipt(ctx, signedTxB.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	})

	ginkgo.It("Receive message on Subnet B", ginkgo.Label("Teleporter", "ReceiveTeleporter"), func() {
		ctx := context.Background()

		// Get the latest block from Subnet B
		log.Info("Waiting for new block confirmation")
		newHeadB := <-newHeadsB
		log.Info("Received new head", "height", newHeadB.Number.Uint64())
		blockHashB := newHeadB.Hash()
		block, err := chainBRPCClient.BlockByHash(ctx, blockHashB)
		Expect(err).Should(BeNil())
		log.Info(
			"Got block",
			"blockHash", blockHashB,
			"blockNumber", block.NumberU64(),
			"transactions", block.Transactions(),
			"numTransactions", len(block.Transactions()),
			"block", block,
		)
		accessLists := block.Transactions()[0].AccessList()
		Expect(len(accessLists)).Should(Equal(1))
		Expect(accessLists[0].Address).Should(Equal(warp.Module.Address))

		// Check the transaction storage key has warp message we're expecting
		storageKeyHashes := accessLists[0].StorageKeys
		packedPredicate := predicateutils.HashSliceToBytes(storageKeyHashes)
		predicateBytes, err := predicateutils.UnpackPredicate(packedPredicate)
		Expect(err).Should(BeNil())
		receivedWarpMessage, err = avalancheWarp.ParseMessage(predicateBytes)
		Expect(err).Should(BeNil())

		// Check that the transaction has successful receipt status
		txHash := block.Transactions()[0].Hash()
		receipt, err := chainBRPCClient.TransactionReceipt(ctx, txHash)
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

		log.Info("Finished sending warp message, closing down output channel")

	})

	ginkgo.It("Validate Received Warp Message Values", ginkgo.Label("Teleporter", "VerifyWarp"), func() {
		Expect(receivedWarpMessage.SourceChainID).Should(Equal(blockchainIDA))
		addressedPayload, err := warpPayload.ParseAddressedPayload(receivedWarpMessage.Payload)
		Expect(err).Should(BeNil())

		receivedDestinationID, err := ids.ToID(addressedPayload.DestinationChainID.Bytes())
		Expect(err).Should(BeNil())
		Expect(receivedDestinationID).Should(Equal(blockchainIDB))
		Expect(addressedPayload.DestinationAddress).Should(Equal(teleporterContractAddress))
		Expect(addressedPayload.Payload).Should(Equal(payload))

		// Check that the teleporter message is correct
		receivedTeleporterMessage, err := teleporter.UnpackTeleporterMessage(addressedPayload.Payload)
		Expect(err).Should(BeNil())
		Expect(*receivedTeleporterMessage).Should(Equal(teleporterMessage))

		teleporterMessageID = receivedTeleporterMessage.MessageID
	})

	ginkgo.It("Check Teleporter Message Received", ginkgo.Label("Teleporter", "TeleporterMessageReceived"), func() {
		data, err := teleporter.PackMessageReceivedMessage(teleporter.MessageReceivedInput{
			OriginChainID: blockchainIDA,
			MessageID:     teleporterMessageID,
		})
		Expect(err).Should(BeNil())
		callMessage := interfaces.CallMsg{
			To:   &teleporterContractAddress,
			Data: data,
		}
		result, err := chainBRPCClient.CallContract(context.Background(), callMessage, nil)
		Expect(err).Should(BeNil())

		// check the contract call result
		delivered, err := teleporter.UnpackMessageReceivedResult(result)
		Expect(err).Should(BeNil())
		Expect(delivered).Should(BeTrue())
	})

})
