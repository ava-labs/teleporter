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
	"github.com/ava-labs/avalanchego/genesis"
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
	primaryNetworkInfo *interfaces.SubnetTestInfo
	subnetsInfo        map[ids.ID]*interfaces.SubnetTestInfo

	extraNodes []*tmpnet.Node // to add as more subnet validators in the tests

	globalFundedKey *ecdsa.PrivateKey
	pChainWallet    pwallet.Wallet

	// Internal vars only used to set up the local network
	tmpnet *tmpnet.Network
}

const (
	fundedKeyStr = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
	timeout      = 120 * time.Second
)

type SubnetSpec struct {
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
	subnetSpecs []SubnetSpec,
	extraNodeCount int, // for use by tests, eg to add new subnet validators
) *LocalNetwork {
	var err error

	// Create extra nodes to be used to add more validators later
	extraNodes := subnetEvmTestUtils.NewTmpnetNodes(extraNodeCount)

	var allNodes []*tmpnet.Node
	allNodes = append(allNodes, extraNodes...) // to be appended w/ subnet validators

	fundedKey, err := hex.DecodeString(fundedKeyStr)
	Expect(err).Should(BeNil())
	globalFundedKey, err := secp256k1.ToPrivateKey(fundedKey)
	Expect(err).Should(BeNil())

	globalFundedECDSAKey := globalFundedKey.ToECDSA()
	Expect(err).Should(BeNil())

	var subnets []*tmpnet.Subnet
	var bootstrapNodes []*tmpnet.Node
	for _, subnetSpec := range subnetSpecs {
		// Create a single bootstrap node. This will be removed from the subnet validator set after it is converted, but will remain a primary network validator
		boostrapNodes := subnetEvmTestUtils.NewTmpnetNodes(1) // One bootstrap node per subnet
		allNodes = append(allNodes, boostrapNodes...)
		bootstrapNodes = append(bootstrapNodes, boostrapNodes...)

		// Create validators to specify as the initial vdr set in the subnet conversion
		initialVdrNodes := subnetEvmTestUtils.NewTmpnetNodes(subnetSpec.NodeCount)
		allNodes = append(allNodes, initialVdrNodes...)
		extraNodes = append(extraNodes, initialVdrNodes...)

		subnet := subnetEvmTestUtils.NewTmpnetSubnet(
			subnetSpec.Name,
			utils.InstantiateGenesisTemplate(
				warpGenesisTemplateFile,
				subnetSpec.EVMChainID,
				subnetSpec.TeleporterContractAddress,
				subnetSpec.TeleporterDeployedBytecode,
				subnetSpec.TeleporterDeployerAddress,
			),
			utils.WarpEnabledChainConfig,
			boostrapNodes...,
		)
		subnet.OwningKey = globalFundedKey
		subnets = append(subnets, subnet)
	}

	network := subnetEvmTestUtils.NewTmpnetNetwork(
		name,
		allNodes,
		utils.WarpEnabledChainConfig,
		subnets...,
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

	// Specify only a subset of the nodes to be bootstrapped
	keysToFund := []*secp256k1.PrivateKey{
		genesis.VMRQKey,
		genesis.EWOQKey,
		tmpnet.HardhatKey,
	}
	keysToFund = append(keysToFund, network.PreFundedKeys...)
	genesis, err := tmpnet.NewTestGenesis(88888, bootstrapNodes, keysToFund)
	Expect(err).Should(BeNil())
	network.Genesis = genesis

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
	for _, subnet := range network.Subnets {
		utils.SetupProposerVM(ctx, globalFundedECDSAKey, network, subnet.SubnetID)
	}

	localNetwork := &LocalNetwork{
		primaryNetworkInfo: &interfaces.SubnetTestInfo{},
		subnetsInfo:        make(map[ids.ID]*interfaces.SubnetTestInfo),
		extraNodes:         extraNodes,
		globalFundedKey:    globalFundedECDSAKey,
		tmpnet:             network,
	}
	for _, subnet := range network.Subnets {
		localNetwork.setSubnetValues(subnet)
	}
	localNetwork.setPrimaryNetworkValues()

	// Create the P-Chain wallet to issue transactions
	kc := secp256k1fx.NewKeychain(globalFundedKey)
	localNetwork.GetSubnetsInfo()
	var subnetIDs []ids.ID
	for _, subnet := range localNetwork.GetSubnetsInfo() {
		subnetIDs = append(subnetIDs, subnet.SubnetID)
	}
	wallet, err := primary.MakeWallet(ctx, &primary.WalletConfig{
		URI:          localNetwork.GetPrimaryNetworkInfo().NodeURIs[0],
		AVAXKeychain: kc,
		EthKeychain:  kc,
		SubnetIDs:    subnetIDs,
	})
	Expect(err).Should(BeNil())
	localNetwork.pChainWallet = wallet.P()

	return localNetwork
}

func (n *LocalNetwork) GetExtraNodes(count uint) []*tmpnet.Node {
	Expect(count > 0).Should(BeTrue(), "can't add 0 validators")
	Expect(uint(len(n.extraNodes)) >= count).Should(
		BeTrue(),
		"not enough extra nodes to use",
	)
	nodes := n.extraNodes[0:count]
	n.extraNodes = n.extraNodes[count:]
	return nodes
}

// Should be called after setSubnetValues for all subnets
func (n *LocalNetwork) setPrimaryNetworkValues() {
	// Get the C-Chain node URIs.
	// All subnet nodes validate the C-Chain, so we can include them all here
	var nodeURIs []string
	for _, subnetInfo := range n.subnetsInfo {
		nodeURIs = append(nodeURIs, subnetInfo.NodeURIs...)
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

	n.primaryNetworkInfo.SubnetID = constants.PrimaryNetworkID
	n.primaryNetworkInfo.BlockchainID = cChainBlockchainID
	n.primaryNetworkInfo.NodeURIs = nodeURIs
	n.primaryNetworkInfo.WSClient = chainWSClient
	n.primaryNetworkInfo.RPCClient = chainRPCClient
	n.primaryNetworkInfo.EVMChainID = chainIDInt
}

func (n *LocalNetwork) setSubnetValues(subnet *tmpnet.Subnet) {
	blockchainID := subnet.Chains[0].ChainID

	var chainNodeURIs []string
	for _, validatorID := range subnet.ValidatorIDs {
		uri, err := n.tmpnet.GetURIForNodeID(validatorID)
		Expect(err).Should(BeNil(), "failed to get URI for node ID %s", validatorID)
		Expect(uri).ShouldNot(HaveLen(0))
		chainNodeURIs = append(chainNodeURIs, uri)
	}

	chainWSURI := utils.HttpToWebsocketURI(chainNodeURIs[0], blockchainID.String())
	chainRPCURI := utils.HttpToRPCURI(chainNodeURIs[0], blockchainID.String())

	subnetID := subnet.SubnetID

	if n.subnetsInfo[subnetID] != nil && n.subnetsInfo[subnetID].WSClient != nil {
		n.subnetsInfo[subnetID].WSClient.Close()
	}
	chainWSClient, err := ethclient.Dial(chainWSURI)
	Expect(err).Should(BeNil())
	if n.subnetsInfo[subnetID] != nil && n.subnetsInfo[subnetID].RPCClient != nil {
		n.subnetsInfo[subnetID].RPCClient.Close()
	}
	chainRPCClient, err := ethclient.Dial(chainRPCURI)
	Expect(err).Should(BeNil())
	chainIDInt, err := chainRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	// Set the new values in the subnetsInfo map
	if n.subnetsInfo[subnetID] == nil {
		n.subnetsInfo[subnetID] = &interfaces.SubnetTestInfo{}
	}
	n.subnetsInfo[subnetID].SubnetName = subnet.Name
	n.subnetsInfo[subnetID].SubnetID = subnetID
	n.subnetsInfo[subnetID].BlockchainID = blockchainID
	n.subnetsInfo[subnetID].NodeURIs = chainNodeURIs
	n.subnetsInfo[subnetID].WSClient = chainWSClient
	n.subnetsInfo[subnetID].RPCClient = chainRPCClient
	n.subnetsInfo[subnetID].EVMChainID = chainIDInt
}

// Returns all subnet info sorted in lexicographic order of SubnetName.
func (n *LocalNetwork) GetSubnetsInfo() []interfaces.SubnetTestInfo {
	subnetsInfo := make([]interfaces.SubnetTestInfo, 0, len(n.subnetsInfo))
	for _, subnetInfo := range n.subnetsInfo {
		subnetsInfo = append(subnetsInfo, *subnetInfo)
	}
	sort.Slice(subnetsInfo, func(i, j int) bool {
		return subnetsInfo[i].SubnetName < subnetsInfo[j].SubnetName
	})
	return subnetsInfo
}

func (n *LocalNetwork) GetPrimaryNetworkInfo() interfaces.SubnetTestInfo {
	return *n.primaryNetworkInfo
}

// Returns subnet info for all subnets, including the primary network
func (n *LocalNetwork) GetAllSubnetsInfo() []interfaces.SubnetTestInfo {
	subnets := n.GetSubnetsInfo()
	return append(subnets, n.GetPrimaryNetworkInfo())
}

func (n *LocalNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	fundedAddress := crypto.PubkeyToAddress(n.globalFundedKey.PublicKey)
	return fundedAddress, n.globalFundedKey
}

func (n *LocalNetwork) setAllSubnetValues() {
	subnetIDs := n.GetSubnetsInfo()
	Expect(len(subnetIDs)).Should(Equal(2))

	for _, subnetInfo := range n.subnetsInfo {
		subnet := n.tmpnet.GetSubnet(subnetInfo.SubnetName)
		Expect(subnet).ShouldNot(BeNil())
		n.setSubnetValues(subnet)
	}

	n.setPrimaryNetworkValues()
}

func (n *LocalNetwork) TearDownNetwork() {
	log.Info("Tearing down network")
	Expect(n).ShouldNot(BeNil())
	Expect(n.tmpnet).ShouldNot(BeNil())
	Expect(n.tmpnet.Stop(context.Background())).Should(BeNil())
}

func (n *LocalNetwork) AddSubnetValidators(ctx context.Context, subnetID ids.ID, count uint) {
	subnet := n.tmpnet.Subnets[slices.IndexFunc(
		n.tmpnet.Subnets,
		func(s *tmpnet.Subnet) bool { return s.SubnetID == subnetID },
	)]

	// consume some of the extraNodes
	newValidatorNodes := n.GetExtraNodes(count)

	apiURI, err := n.tmpnet.GetURIForNodeID(subnet.ValidatorIDs[0])
	Expect(err).Should(BeNil())

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	err = subnet.AddValidators(
		ctx,
		os.Stdout,
		apiURI,
		newValidatorNodes...,
	)
	Expect(err).Should(BeNil())

	for _, node := range newValidatorNodes {
		subnet.ValidatorIDs = append(subnet.ValidatorIDs, node.NodeID)
		node.Flags[config.TrackSubnetsKey] = subnetID.String()
	}

	tmpnet.WaitForActiveValidators(ctx, os.Stdout, platformvm.NewClient(n.tmpnet.Nodes[0].URI), subnet)

	nodeIdsToRestart := make([]ids.NodeID, len(newValidatorNodes))
	for i, node := range newValidatorNodes {
		nodeIdsToRestart[i] = node.NodeID
	}
	n.RestartNodes(ctx, nodeIdsToRestart)

	n.setAllSubnetValues()
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

	n.setAllSubnetValues()
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

		for _, subnet := range n.tmpnet.Subnets {
			for _, chain := range subnet.Chains {
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
	for _, subnet := range n.tmpnet.Subnets {
		err := subnet.Write(n.tmpnet.GetSubnetDir(), n.tmpnet.GetChainConfigDir())
		if err != nil {
			log.Error("failed to write subnets", "error", err)
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
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
) {
	sourceTeleporterMessenger := teleporter.TeleporterMessenger(source)
	outstandReceiptCount := utils.GetOutstandingReceiptCount(
		teleporter.TeleporterMessenger(source),
		destination.BlockchainID,
	)
	for outstandReceiptCount.Cmp(big.NewInt(0)) != 0 {
		log.Info("Emptying receipt queue", "remainingReceipts", outstandReceiptCount.String())
		// Send message from Subnet B to Subnet A to trigger the "regular" method of delivering receipts.
		// The next message from B->A will contain the same receipts that were manually sent in the above steps,
		// but they should not be processed again on Subnet A.
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
// and we don't want to add the ValidatorSetSig ABI to the subnetInfo
func (n *LocalNetwork) ExecuteValidatorSetSigCallAndVerify(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
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
	subnet interfaces.SubnetTestInfo,
	newTeleporterAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	unsignedMessage *avalancheWarp.UnsignedMessage,
) {
	signedWarpMsg := utils.GetSignedMessage(ctx, subnet, subnet, unsignedMessage.ID())
	log.Info("Got signed warp message", "messageID", signedWarpMsg.ID())

	// Construct tx to add protocol version and send to destination chain
	signedTx := utils.CreateAddProtocolVersionTransaction(
		ctx,
		signedWarpMsg,
		teleporter.TeleporterRegistryAddress(subnet),
		senderKey,
		subnet,
	)

	curLatestVersion := teleporter.GetLatestTeleporterVersion(subnet)
	expectedLatestVersion := big.NewInt(curLatestVersion.Int64() + 1)

	// Wait for tx to be accepted, and verify events emitted
	receipt := utils.SendTransactionAndWaitForSuccess(ctx, subnet, signedTx)
	teleporterRegistry := teleporter.TeleporterRegistry(subnet)
	addProtocolVersionEvent, err := utils.GetEventFromLogs(receipt.Logs, teleporterRegistry.ParseAddProtocolVersion)
	Expect(err).Should(BeNil())
	Expect(addProtocolVersionEvent.Version.Cmp(expectedLatestVersion)).Should(Equal(0))
	Expect(addProtocolVersionEvent.ProtocolAddress).Should(Equal(newTeleporterAddress))

	versionUpdatedEvent, err := utils.GetEventFromLogs(receipt.Logs, teleporterRegistry.ParseLatestVersionUpdated)
	Expect(err).Should(BeNil())
	Expect(versionUpdatedEvent.OldVersion.Cmp(curLatestVersion)).Should(Equal(0))
	Expect(versionUpdatedEvent.NewVersion.Cmp(expectedLatestVersion)).Should(Equal(0))
}

func (n *LocalNetwork) GetTwoSubnets() (
	interfaces.SubnetTestInfo,
	interfaces.SubnetTestInfo,
) {
	subnets := n.GetSubnetsInfo()
	Expect(len(subnets)).Should(BeNumerically(">=", 2))
	return subnets[0], subnets[1]
}

func (n *LocalNetwork) SendExampleCrossChainMessageAndVerify(
	ctx context.Context,
	teleporter utils.TeleporterTestInfo,
	source interfaces.SubnetTestInfo,
	sourceExampleMessenger *testmessenger.TestMessenger,
	destination interfaces.SubnetTestInfo,
	destExampleMessengerAddress common.Address,
	destExampleMessenger *testmessenger.TestMessenger,
	senderKey *ecdsa.PrivateKey,
	message string,
	expectSuccess bool,
) {
	// Call the example messenger contract on Subnet A
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
