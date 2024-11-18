package network

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"os"
	"slices"
	"sort"
	"time"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/config"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	pwallet "github.com/ava-labs/avalanchego/wallet/chain/p/wallet"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	subnetEvmTestUtils "github.com/ava-labs/subnet-evm/tests/utils"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/teleporter/TeleporterMessenger"
	testmessenger "github.com/ava-labs/teleporter/abi-bindings/go/teleporter/tests/TestMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

// Implements Network, pointing to the network setup in local_network_setup.go
type LocalNetwork struct {
	primaryNetworkInfo *interfaces.L1TestInfo
	L1Infos            map[ids.ID]*interfaces.L1TestInfo

	extraNodes []*tmpnet.Node // to add as more L1 validators in the tests

	globalFundedKey *ecdsa.PrivateKey
	pChainWallet    pwallet.Wallet

	// Internal vars only used to set up the local network
	tmpnet *tmpnet.Network
}

const (
	fundedKeyStr = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
	timeout      = 120 * time.Second
)

type L1Spec struct {
	Name       string
	EVMChainID uint64
	NodeCount  int

	// Optional fields
	TeleporterContractAddress  common.Address
	TeleporterDeployedBytecode string
	TeleporterDeployerAddress  common.Address
}

func NewLocalNetwork(
	ctx context.Context,
	name string,
	warpGenesisTemplateFile string,
	l1Specs []L1Spec,
	extraNodeCount int, // for use by tests, eg to add new L1 validators
) *LocalNetwork {
	var err error

	// Create extra nodes to be used to add more validators later
	extraNodes := subnetEvmTestUtils.NewTmpnetNodes(extraNodeCount)

	var allNodes []*tmpnet.Node
	allNodes = append(allNodes, extraNodes...) // to be appended w/ L1 validators

	fundedKey, err := hex.DecodeString(fundedKeyStr)
	Expect(err).Should(BeNil())
	globalFundedKey, err := secp256k1.ToPrivateKey(fundedKey)
	Expect(err).Should(BeNil())

	globalFundedECDSAKey := globalFundedKey.ToECDSA()
	Expect(err).Should(BeNil())

	var l1s []*tmpnet.Subnet
	for _, subnetSpec := range l1Specs {
		nodes := subnetEvmTestUtils.NewTmpnetNodes(subnetSpec.NodeCount)
		allNodes = append(allNodes, nodes...)

		l1 := subnetEvmTestUtils.NewTmpnetSubnet(
			subnetSpec.Name,
			utils.InstantiateGenesisTemplate(
				warpGenesisTemplateFile,
				subnetSpec.EVMChainID,
				subnetSpec.TeleporterContractAddress,
				subnetSpec.TeleporterDeployedBytecode,
				subnetSpec.TeleporterDeployerAddress,
			),
			utils.WarpEnabledChainConfig,
			nodes...,
		)
		l1.OwningKey = globalFundedKey
		l1s = append(l1s, l1)
	}

	network := subnetEvmTestUtils.NewTmpnetNetwork(
		name,
		allNodes,
		utils.WarpEnabledChainConfig,
		l1s...,
	)
	Expect(network).ShouldNot(BeNil())

	// Activate Etna
	upgrades := upgrade.Default
	upgrades.EtnaTime = time.Now().Add(-1 * time.Minute)
	upgradeJSON, err := json.Marshal(upgrades)
	Expect(err).Should(BeNil())

	upgradeBase64 := base64.StdEncoding.EncodeToString(upgradeJSON)
	network.DefaultFlags.SetDefaults(tmpnet.FlagsMap{
		config.UpgradeFileContentKey: upgradeBase64,
	})

	avalancheGoBuildPath, ok := os.LookupEnv("AVALANCHEGO_BUILD_PATH")
	Expect(ok).Should(Equal(true))

	ctx, cancelBootstrap := context.WithCancel(ctx)
	defer cancelBootstrap()
	err = tmpnet.BootstrapNewNetwork(
		ctx,
		os.Stdout,
		network,
		"",
		avalancheGoBuildPath+"/avalanchego",
		avalancheGoBuildPath+"/plugins",
	)
	Expect(err).Should(BeNil())

	// Issue transactions to activate the proposerVM fork on the chains
	for _, l1 := range network.Subnets {
		utils.SetupProposerVM(ctx, globalFundedECDSAKey, network, l1.SubnetID)
	}

	localNetwork := &LocalNetwork{
		primaryNetworkInfo: &interfaces.L1TestInfo{},
		L1Infos:            make(map[ids.ID]*interfaces.L1TestInfo),
		extraNodes:         extraNodes,
		globalFundedKey:    globalFundedECDSAKey,
		tmpnet:             network,
	}
	for _, l1 := range network.Subnets {
		localNetwork.setL1Values(l1)
	}
	localNetwork.setPrimaryNetworkValues()

	// Create the P-Chain wallet to issue transactions
	kc := secp256k1fx.NewKeychain(globalFundedKey)
	localNetwork.GetL1Infos()
	var l1IDs []ids.ID
	for _, l1 := range localNetwork.GetL1Infos() {
		l1IDs = append(l1IDs, l1.L1ID)
	}
	wallet, err := primary.MakeWallet(ctx, &primary.WalletConfig{
		URI:          localNetwork.GetPrimaryNetworkInfo().NodeURIs[0],
		AVAXKeychain: kc,
		EthKeychain:  kc,
		SubnetIDs:    l1IDs,
	})
	Expect(err).Should(BeNil())
	localNetwork.pChainWallet = wallet.P()

	return localNetwork
}

// Should be called after setL1Values for all L1s
func (n *LocalNetwork) setPrimaryNetworkValues() {
	// Get the C-Chain node URIs.
	// All L1 nodes validate the C-Chain, so we can include them all here
	var nodeURIs []string
	for _, l1Info := range n.L1Infos {
		nodeURIs = append(nodeURIs, l1Info.NodeURIs...)
	}
	for _, extraNode := range n.extraNodes {
		uri, err := n.tmpnet.GetURIForNodeID(extraNode.NodeID)
		Expect(err).Should(BeNil())
		nodeURIs = append(nodeURIs, uri)
	}

	cChainBlockchainID, err := info.NewClient(nodeURIs[0]).GetBlockchainID(context.Background(), "C")
	Expect(err).Should(BeNil())
	Expect(cChainBlockchainID).ShouldNot(Equal(ids.Empty))

	chainWSURI := utils.HttpToWebsocketURI(nodeURIs[0], cChainBlockchainID.String())
	chainRPCURI := utils.HttpToRPCURI(nodeURIs[0], cChainBlockchainID.String())
	if n.primaryNetworkInfo != nil && n.primaryNetworkInfo.WSClient != nil {
		n.primaryNetworkInfo.WSClient.Close()
	}
	chainWSClient, err := ethclient.Dial(chainWSURI)
	Expect(err).Should(BeNil())
	if n.primaryNetworkInfo != nil && n.primaryNetworkInfo.RPCClient != nil {
		n.primaryNetworkInfo.RPCClient.Close()
	}
	chainRPCClient, err := ethclient.Dial(chainRPCURI)
	Expect(err).Should(BeNil())
	chainIDInt, err := chainRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	n.primaryNetworkInfo.L1ID = constants.PrimaryNetworkID
	n.primaryNetworkInfo.BlockchainID = cChainBlockchainID
	n.primaryNetworkInfo.NodeURIs = nodeURIs
	n.primaryNetworkInfo.WSClient = chainWSClient
	n.primaryNetworkInfo.RPCClient = chainRPCClient
	n.primaryNetworkInfo.EVMChainID = chainIDInt
}

func (n *LocalNetwork) setL1Values(l1 *tmpnet.Subnet) {
	blockchainID := l1.Chains[0].ChainID

	var chainNodeURIs []string
	for _, validatorID := range l1.ValidatorIDs {
		uri, err := n.tmpnet.GetURIForNodeID(validatorID)
		Expect(err).Should(BeNil(), "failed to get URI for node ID %s", validatorID)
		Expect(uri).ShouldNot(HaveLen(0))
		chainNodeURIs = append(chainNodeURIs, uri)
	}

	chainWSURI := utils.HttpToWebsocketURI(chainNodeURIs[0], blockchainID.String())
	chainRPCURI := utils.HttpToRPCURI(chainNodeURIs[0], blockchainID.String())

	L1ID := l1.SubnetID

	if n.L1Infos[L1ID] != nil && n.L1Infos[L1ID].WSClient != nil {
		n.L1Infos[L1ID].WSClient.Close()
	}
	chainWSClient, err := ethclient.Dial(chainWSURI)
	Expect(err).Should(BeNil())
	if n.L1Infos[L1ID] != nil && n.L1Infos[L1ID].RPCClient != nil {
		n.L1Infos[L1ID].RPCClient.Close()
	}
	chainRPCClient, err := ethclient.Dial(chainRPCURI)
	Expect(err).Should(BeNil())
	chainIDInt, err := chainRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	// Set the new values in the L1Infos map
	if n.L1Infos[L1ID] == nil {
		n.L1Infos[L1ID] = &interfaces.L1TestInfo{}
	}
	n.L1Infos[L1ID].L1Name = l1.Name
	n.L1Infos[L1ID].L1ID = L1ID
	n.L1Infos[L1ID].BlockchainID = blockchainID
	n.L1Infos[L1ID].NodeURIs = chainNodeURIs
	n.L1Infos[L1ID].WSClient = chainWSClient
	n.L1Infos[L1ID].RPCClient = chainRPCClient
	n.L1Infos[L1ID].EVMChainID = chainIDInt
}

// Returns all L1 info sorted in lexicographic order of L1Name.
func (n *LocalNetwork) GetL1Infos() []interfaces.L1TestInfo {
	L1Infos := make([]interfaces.L1TestInfo, 0, len(n.L1Infos))
	for _, l1Info := range n.L1Infos {
		L1Infos = append(L1Infos, *l1Info)
	}
	sort.Slice(L1Infos, func(i, j int) bool {
		return L1Infos[i].L1Name < L1Infos[j].L1Name
	})
	return L1Infos
}

func (n *LocalNetwork) GetPrimaryNetworkInfo() interfaces.L1TestInfo {
	return *n.primaryNetworkInfo
}

// Returns L1 info for all L1s, including the primary network
func (n *LocalNetwork) GetAllL1Infos() []interfaces.L1TestInfo {
	l1s := n.GetL1Infos()
	return append(l1s, n.GetPrimaryNetworkInfo())
}

func (n *LocalNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	fundedAddress := crypto.PubkeyToAddress(n.globalFundedKey.PublicKey)
	return fundedAddress, n.globalFundedKey
}

func (n *LocalNetwork) setAllL1Values() {
	L1IDs := n.GetL1Infos()
	Expect(len(L1IDs)).Should(Equal(2))

	for _, l1Info := range n.L1Infos {
		l1 := n.tmpnet.GetSubnet(l1Info.L1Name)
		Expect(l1).ShouldNot(BeNil())
		n.setL1Values(l1)
	}

	n.setPrimaryNetworkValues()
}

func (n *LocalNetwork) TearDownNetwork() {
	log.Info("Tearing down network")
	Expect(n).ShouldNot(BeNil())
	Expect(n.tmpnet).ShouldNot(BeNil())
	Expect(n.tmpnet.Stop(context.Background())).Should(BeNil())
}

func (n *LocalNetwork) AddL1Validators(ctx context.Context, l1ID ids.ID, count uint) {
	Expect(count > 0).Should(BeTrue(), "can't add 0 validators")
	Expect(uint(len(n.extraNodes)) >= count).Should(
		BeTrue(),
		"not enough extra nodes to use",
	)

	l1 := n.tmpnet.Subnets[slices.IndexFunc(
		n.tmpnet.Subnets,
		func(s *tmpnet.Subnet) bool { return s.SubnetID == l1ID },
	)]

	// consume some of the extraNodes
	var newValidatorNodes []*tmpnet.Node
	newValidatorNodes = append(newValidatorNodes, n.extraNodes[0:count]...)
	n.extraNodes = n.extraNodes[count:]

	apiURI, err := n.tmpnet.GetURIForNodeID(l1.ValidatorIDs[0])
	Expect(err).Should(BeNil())

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	err = l1.AddValidators(
		ctx,
		os.Stdout,
		apiURI,
		newValidatorNodes...,
	)
	Expect(err).Should(BeNil())

	for _, node := range newValidatorNodes {
		l1.ValidatorIDs = append(l1.ValidatorIDs, node.NodeID)
		node.Flags[config.TrackSubnetsKey] = l1ID.String()
	}

	tmpnet.WaitForActiveValidators(ctx, os.Stdout, platformvm.NewClient(n.tmpnet.Nodes[0].URI), l1)

	nodeIdsToRestart := make([]ids.NodeID, len(newValidatorNodes))
	for i, node := range newValidatorNodes {
		nodeIdsToRestart[i] = node.NodeID
	}
	n.RestartNodes(ctx, nodeIdsToRestart)

	n.setAllL1Values()
}

// Restarts the nodes with the given nodeIDs. If nodeIDs is empty, restarts all nodes.
func (n *LocalNetwork) RestartNodes(ctx context.Context, nodeIDs []ids.NodeID) {
	log.Info("Restarting nodes", "nodeIDs", nodeIDs)
	var nodes []*tmpnet.Node
	if len(nodeIDs) == 0 {
		nodes = n.tmpnet.Nodes
	} else {
		for _, nodeID := range nodeIDs {
			for _, node := range n.tmpnet.Nodes {
				if node.NodeID == nodeID {
					nodes = append(nodes, node)
				}
			}
		}
	}

	for _, node := range nodes {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		err := node.SaveAPIPort()
		Expect(err).Should(BeNil())

		err = node.Stop(ctx)
		Expect(err).Should(BeNil())

		err = n.tmpnet.StartNode(ctx, os.Stdout, node)
		Expect(err).Should(BeNil())
	}

	log.Info("Waiting for all nodes to report healthy")
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, node := range nodes {
		err := tmpnet.WaitForHealthy(ctx, node)
		Expect(err).Should(BeNil())
	}

	n.setAllL1Values()
}

func (n *LocalNetwork) SetChainConfigs(chainConfigs map[string]string) {
	for chainIDStr, chainConfig := range chainConfigs {
		if chainIDStr == utils.CChainPathSpecifier {
			var cfg tmpnet.FlagsMap
			err := json.Unmarshal([]byte(chainConfig), &cfg)
			if err != nil {
				log.Error(
					"failed to unmarshal chain config",
					"error", err,
					"chainConfig", chainConfig,
				)
			}
			n.tmpnet.ChainConfigs[utils.CChainPathSpecifier] = cfg
			continue
		}

		for _, l1 := range n.tmpnet.Subnets {
			for _, chain := range l1.Chains {
				if chain.ChainID.String() == chainIDStr {
					chain.Config = chainConfig
				}
			}
		}
	}
	err := n.tmpnet.Write()
	if err != nil {
		log.Error("failed to write network", "error", err)
	}
	for _, l1 := range n.tmpnet.Subnets {
		err := l1.Write(n.tmpnet.GetSubnetDir(), n.tmpnet.GetChainConfigDir())
		if err != nil {
			log.Error("failed to write L1s", "error", err)
		}
	}
}

func (n *LocalNetwork) GetNetworkID() uint32 {
	return n.tmpnet.Genesis.NetworkID
}

func (n *LocalNetwork) Dir() string {
	return n.tmpnet.Dir
}

func (n *LocalNetwork) GetPChainWallet() pwallet.Wallet {
	return n.pChainWallet
}

func (n *LocalNetwork) ClearReceiptQueue(
	ctx context.Context,
	teleporter utils.TeleporterTestInfo,
	fundedKey *ecdsa.PrivateKey,
	source interfaces.L1TestInfo,
	destination interfaces.L1TestInfo,
) {
	sourceTeleporterMessenger := teleporter.TeleporterMessenger(source)
	outstandReceiptCount := utils.GetOutstandingReceiptCount(
		teleporter.TeleporterMessenger(source),
		destination.BlockchainID,
	)
	for outstandReceiptCount.Cmp(big.NewInt(0)) != 0 {
		log.Info("Emptying receipt queue", "remainingReceipts", outstandReceiptCount.String())
		// Send message from L1 B to L1 A to trigger the "regular" method of delivering receipts.
		// The next message from B->A will contain the same receipts that were manually sent in the above steps,
		// but they should not be processed again on L1 A.
		sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
			DestinationBlockchainID: destination.BlockchainID,
			DestinationAddress:      common.HexToAddress("0x1111111111111111111111111111111111111111"),
			RequiredGasLimit:        big.NewInt(1),
			FeeInfo: teleportermessenger.TeleporterFeeInfo{
				FeeTokenAddress: common.Address{},
				Amount:          big.NewInt(0),
			},
			AllowedRelayerAddresses: []common.Address{},
			Message:                 []byte{1, 2, 3, 4},
		}

		// This message will also have the same receipts as the previous message
		receipt, _ := utils.SendCrossChainMessageAndWaitForAcceptance(
			ctx, sourceTeleporterMessenger, source, destination, sendCrossChainMessageInput, fundedKey)

		// Relay message
		teleporter.RelayTeleporterMessage(ctx, receipt, source, destination, true, fundedKey)

		outstandReceiptCount = utils.GetOutstandingReceiptCount(sourceTeleporterMessenger, destination.BlockchainID)
	}
	log.Info("Receipt queue emptied")
}

// Returns Receipt for the transaction unlike TeleporterRegistry version since this is a non-teleporter case
// and we don't want to add the ValidatorSetSig ABI to the l1Info
func (n *LocalNetwork) ExecuteValidatorSetSigCallAndVerify(
	ctx context.Context,
	source interfaces.L1TestInfo,
	destination interfaces.L1TestInfo,
	validatorSetSigAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	unsignedMessage *avalancheWarp.UnsignedMessage,
	expectSuccess bool,
) *types.Receipt {
	signedWarpMsg := utils.GetSignedMessage(ctx, source, destination, unsignedMessage.ID())
	log.Info("Got signed warp message", "messageID", signedWarpMsg.ID())

	signedPredicateTx := utils.CreateExecuteCallPredicateTransaction(
		ctx,
		signedWarpMsg,
		validatorSetSigAddress,
		senderKey,
		destination,
	)

	// Wait for tx to be accepted and verify events emitted
	if expectSuccess {
		return utils.SendTransactionAndWaitForSuccess(ctx, destination, signedPredicateTx)
	}
	return utils.SendTransactionAndWaitForFailure(ctx, destination, signedPredicateTx)
}

func (n *LocalNetwork) AddProtocolVersionAndWaitForAcceptance(
	ctx context.Context,
	teleporter utils.TeleporterTestInfo,
	l1 interfaces.L1TestInfo,
	newTeleporterAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	unsignedMessage *avalancheWarp.UnsignedMessage,
) {
	signedWarpMsg := utils.GetSignedMessage(ctx, l1, l1, unsignedMessage.ID())
	log.Info("Got signed warp message", "messageID", signedWarpMsg.ID())

	// Construct tx to add protocol version and send to destination chain
	signedTx := utils.CreateAddProtocolVersionTransaction(
		ctx,
		signedWarpMsg,
		teleporter.TeleporterRegistryAddress(l1),
		senderKey,
		l1,
	)

	curLatestVersion := teleporter.GetLatestTeleporterVersion(l1)
	expectedLatestVersion := big.NewInt(curLatestVersion.Int64() + 1)

	// Wait for tx to be accepted, and verify events emitted
	receipt := utils.SendTransactionAndWaitForSuccess(ctx, l1, signedTx)
	teleporterRegistry := teleporter.TeleporterRegistry(l1)
	addProtocolVersionEvent, err := utils.GetEventFromLogs(receipt.Logs, teleporterRegistry.ParseAddProtocolVersion)
	Expect(err).Should(BeNil())
	Expect(addProtocolVersionEvent.Version.Cmp(expectedLatestVersion)).Should(Equal(0))
	Expect(addProtocolVersionEvent.ProtocolAddress).Should(Equal(newTeleporterAddress))

	versionUpdatedEvent, err := utils.GetEventFromLogs(receipt.Logs, teleporterRegistry.ParseLatestVersionUpdated)
	Expect(err).Should(BeNil())
	Expect(versionUpdatedEvent.OldVersion.Cmp(curLatestVersion)).Should(Equal(0))
	Expect(versionUpdatedEvent.NewVersion.Cmp(expectedLatestVersion)).Should(Equal(0))
}

func (n *LocalNetwork) GetTwoL1s() (
	interfaces.L1TestInfo,
	interfaces.L1TestInfo,
) {
	l1s := n.GetL1Infos()
	Expect(len(l1s)).Should(BeNumerically(">=", 2))
	return l1s[0], l1s[1]
}

func (n *LocalNetwork) SendExampleCrossChainMessageAndVerify(
	ctx context.Context,
	teleporter utils.TeleporterTestInfo,
	source interfaces.L1TestInfo,
	sourceExampleMessenger *testmessenger.TestMessenger,
	destination interfaces.L1TestInfo,
	destExampleMessengerAddress common.Address,
	destExampleMessenger *testmessenger.TestMessenger,
	senderKey *ecdsa.PrivateKey,
	message string,
	expectSuccess bool,
) {
	// Call the example messenger contract on L1 A
	optsA, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := sourceExampleMessenger.SendMessage(
		optsA,
		destination.BlockchainID,
		destExampleMessengerAddress,
		common.BigToAddress(common.Big0),
		big.NewInt(0),
		testmessenger.SendMessageRequiredGas,
		message,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt := utils.WaitForTransactionSuccess(ctx, source, tx.Hash())

	sourceTeleporterMessenger := teleporter.TeleporterMessenger(source)
	destTeleporterMessenger := teleporter.TeleporterMessenger(destination)

	event, err := utils.GetEventFromLogs(receipt.Logs, sourceTeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(destination.BlockchainID[:]))

	teleporterMessageID := event.MessageID

	//
	// Relay the message to the destination
	//
	receipt = teleporter.RelayTeleporterMessage(ctx, receipt, source, destination, true, senderKey)

	//
	// Check Teleporter message received on the destination
	//
	delivered, err := destTeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	if expectSuccess {
		// Check that message execution was successful
		messageExecutedEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			destTeleporterMessenger.ParseMessageExecuted,
		)
		Expect(err).Should(BeNil())
		Expect(messageExecutedEvent.MessageID[:]).Should(Equal(teleporterMessageID[:]))
	} else {
		// Check that message execution failed
		messageExecutionFailedEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			destTeleporterMessenger.ParseMessageExecutionFailed,
		)
		Expect(err).Should(BeNil())
		Expect(messageExecutionFailedEvent.MessageID[:]).Should(Equal(teleporterMessageID[:]))
	}

	//
	// Verify we received the expected string
	//
	_, currMessage, err := destExampleMessenger.GetCurrentMessage(&bind.CallOpts{}, source.BlockchainID)
	Expect(err).Should(BeNil())
	if expectSuccess {
		Expect(currMessage).Should(Equal(message))
	} else {
		Expect(currMessage).ShouldNot(Equal(message))
	}
}
