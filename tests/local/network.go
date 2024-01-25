package local

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"time"

	runner_sdk "github.com/ava-labs/avalanche-network-runner/client"
	"github.com/ava-labs/avalanche-network-runner/rpcpb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	subnetEvmInterfaces "github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ava-labs/subnet-evm/plugin/evm"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	"github.com/ava-labs/subnet-evm/rpc"
	"github.com/ava-labs/subnet-evm/tests/utils/runner"
	warpBackend "github.com/ava-labs/subnet-evm/warp"

	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	teleporterregistry "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/upgrades/TeleporterRegistry"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

var _ interfaces.LocalNetwork = &LocalNetwork{}

// Implements Network, pointing to the network setup in local_network_setup.go
type LocalNetwork struct {
	teleporterContractAddress common.Address
	primaryNetworkInfo        *interfaces.SubnetTestInfo
	subnetAID, subnetBID      ids.ID
	subnetsInfo               map[ids.ID]*interfaces.SubnetTestInfo
	subnetNodeNames           map[ids.ID][]string

	globalFundedKey *ecdsa.PrivateKey

	// Internal vars only used to set up the local network
	anrClient           runner_sdk.Client
	manager             *runner.NetworkManager
	warpChainConfigPath string
}

const (
	fundedKeyStr           = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
	warpEnabledChainConfig = `{
		"warp-api-enabled": true, 
		"eth-apis":["eth","eth-filter","net","admin","web3",
					"internal-eth","internal-blockchain","internal-transaction",
					"internal-debug","internal-account","internal-personal",
					"debug","debug-tracer","debug-file-tracer","debug-handler"]
	}`
)

func NewLocalNetwork(warpGenesisFile string) *LocalNetwork {
	ctx := context.Background()
	var err error

	// Name 10 new validators (which should have BLS key registered)
	var subnetANodeNames []string
	var subnetBNodeNames []string
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
	_, err = f.Write([]byte(warpEnabledChainConfig))
	Expect(err).Should(BeNil())
	warpChainConfigPath := f.Name()

	// Make sure that the warp genesis file exists
	_, err = os.Stat(warpGenesisFile)
	Expect(err).Should(BeNil())

	anrConfig := runner.NewDefaultANRConfig()
	anrConfig.GlobalCChainConfig = warpEnabledChainConfig
	manager := runner.NewNetworkManager(anrConfig)

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

	// Issue transactions to activate the proposerVM fork on the chains
	globalFundedKey, err := crypto.HexToECDSA(fundedKeyStr)
	Expect(err).Should(BeNil())
	setupProposerVM(ctx, globalFundedKey, manager, 0)
	setupProposerVM(ctx, globalFundedKey, manager, 1)

	// Create the ANR client
	logLevel, err := logging.ToLevel("info")
	Expect(err).Should(BeNil())

	logFactory := logging.NewFactory(logging.Config{
		DisplayLevel: logLevel,
		LogLevel:     logLevel,
	})
	zapLog, err := logFactory.Make("main")
	Expect(err).Should(BeNil())

	anrClient, err := runner_sdk.New(runner_sdk.Config{
		Endpoint:    "0.0.0.0:12352",
		DialTimeout: 10 * time.Second,
	}, zapLog)
	Expect(err).Should(BeNil())

	// On initial startup, we need to first set the subnet node names
	// before calling setSubnetValues for the two subnets
	subnetIDs := manager.GetSubnets()
	Expect(len(subnetIDs)).Should(Equal(2))
	subnetAID := subnetIDs[0]
	subnetBID := subnetIDs[1]

	res := &LocalNetwork{
		subnetAID:          subnetAID,
		subnetBID:          subnetBID,
		primaryNetworkInfo: &interfaces.SubnetTestInfo{},
		subnetsInfo:        make(map[ids.ID]*interfaces.SubnetTestInfo),
		subnetNodeNames: map[ids.ID][]string{
			subnetAID: subnetANodeNames,
			subnetBID: subnetBNodeNames,
		},
		globalFundedKey:     globalFundedKey,
		anrClient:           anrClient,
		manager:             manager,
		warpChainConfigPath: warpChainConfigPath,
	}
	res.setSubnetValues(subnetAID)
	res.setSubnetValues(subnetBID)
	res.setPrimaryNetworkValues()
	return res
}

// Should be called after setSubnetValues for all subnets
func (n *LocalNetwork) setPrimaryNetworkValues() {
	// Get the C-Chain node URIs.
	// All subnet nodes validate the C-Chain, so we can include them all here
	var nodeURIs []string
	nodeURIs = append(nodeURIs, n.subnetsInfo[n.subnetAID].NodeURIs...)
	nodeURIs = append(nodeURIs, n.subnetsInfo[n.subnetBID].NodeURIs...)
	pChainClient := platformvm.NewClient(nodeURIs[0])
	blockChains, err := pChainClient.GetBlockchains(context.Background())
	Expect(err).Should(BeNil())

	var cChainBlockchainID ids.ID
	for _, chain := range blockChains {
		if chain.Name == "C-Chain" {
			cChainBlockchainID = chain.ID
		}
	}
	Expect(cChainBlockchainID).ShouldNot(Equal(ids.Empty))

	chainWSURI := utils.HttpToWebsocketURI(nodeURIs[0], utils.CChainPathSpecifier)
	chainRPCURI := utils.HttpToRPCURI(nodeURIs[0], utils.CChainPathSpecifier)
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

	// TeleporterMessenger is set in DeployTeleporterContracts
	// TeleporterRegistryAddress is set in DeployTeleporterRegistryContracts
}

func (n *LocalNetwork) setSubnetValues(subnetID ids.ID) {
	subnetDetails, ok := n.manager.GetSubnet(subnetID)
	Expect(ok).Should(BeTrue())
	blockchainID := subnetDetails.BlockchainID

	// Reset the validator URIs, as they may have changed
	subnetDetails.ValidatorURIs = nil
	status, err := n.anrClient.Status(context.Background())
	Expect(err).Should(BeNil())
	nodeInfos := status.GetClusterInfo().GetNodeInfos()

	for _, nodeName := range n.subnetNodeNames[subnetID] {
		log.Info("Adding validator URI", "nodeName", nodeName, "uri", nodeInfos[nodeName].Uri)
		subnetDetails.ValidatorURIs = append(subnetDetails.ValidatorURIs, nodeInfos[nodeName].Uri)
	}
	var chainNodeURIs []string
	chainNodeURIs = append(chainNodeURIs, subnetDetails.ValidatorURIs...)

	chainWSURI := utils.HttpToWebsocketURI(chainNodeURIs[0], blockchainID.String())
	chainRPCURI := utils.HttpToRPCURI(chainNodeURIs[0], blockchainID.String())

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
	n.subnetsInfo[subnetID].SubnetID = subnetID
	n.subnetsInfo[subnetID].BlockchainID = blockchainID
	n.subnetsInfo[subnetID].NodeURIs = chainNodeURIs
	n.subnetsInfo[subnetID].WSClient = chainWSClient
	n.subnetsInfo[subnetID].RPCClient = chainRPCClient
	n.subnetsInfo[subnetID].EVMChainID = chainIDInt

	// TeleporterMessenger is set in DeployTeleporterContracts
	// TeleporterRegistryAddress is set in DeployTeleporterRegistryContracts
}

// DeployTeleporterContracts deploys the Teleporter contract to all subnets.
// The caller is responsible for generating the deployment transaction information
func (n *LocalNetwork) DeployTeleporterContracts(
	transactionBytes []byte,
	deployerAddress common.Address,
	contractAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	updateNetworkTeleporter bool,
) {
	log.Info("Deploying Teleporter contract to subnets", "contractAddress", contractAddress.String())

	ctx := context.Background()

	subnets := n.GetAllSubnetsInfo()
	for _, subnetInfo := range subnets {
		// Fund the deployer address
		{
			fundAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(11)) // 11 AVAX
			fundDeployerTx := utils.CreateNativeTransferTransaction(
				ctx, subnetInfo, fundedKey, deployerAddress, fundAmount,
			)
			utils.SendTransactionAndWaitForSuccess(ctx, subnetInfo, fundDeployerTx)
		}
		log.Info("Finished funding Teleporter deployer", "blockchainID", subnetInfo.BlockchainID.Hex())

		// Deploy Teleporter contract
		{
			rpcClient, err := rpc.DialContext(
				ctx,
				utils.HttpToRPCURI(subnetInfo.NodeURIs[0], subnetInfo.BlockchainID.String()),
			)
			Expect(err).Should(BeNil())
			defer rpcClient.Close()

			newHeads := make(chan *types.Header, 10)
			sub, err := subnetInfo.WSClient.SubscribeNewHead(ctx, newHeads)
			Expect(err).Should(BeNil())
			defer sub.Unsubscribe()

			err = rpcClient.CallContext(ctx, nil, "eth_sendRawTransaction", hexutil.Encode(transactionBytes))
			Expect(err).Should(BeNil())

			<-newHeads
			teleporterCode, err := subnetInfo.RPCClient.CodeAt(ctx, contractAddress, nil)
			Expect(err).Should(BeNil())
			Expect(len(teleporterCode)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
		}
		log.Info("Finished deploying Teleporter contract", "blockchainID", subnetInfo.BlockchainID.Hex())
	}

	if updateNetworkTeleporter {
		n.SetTeleporterContractAddress(contractAddress)
	}
	log.Info("Deployed Teleporter contracts to all subnets")
}

func (n *LocalNetwork) DeployTeleporterRegistryContracts(
	teleporterAddress common.Address,
	deployerKey *ecdsa.PrivateKey,
) {
	log.Info("Deploying TeleporterRegistry contract to subnets")
	ctx := context.Background()

	entries := []teleporterregistry.ProtocolRegistryEntry{
		{
			Version:         big.NewInt(1),
			ProtocolAddress: teleporterAddress,
		},
	}

	subnets := n.GetAllSubnetsInfo()
	for _, subnetInfo := range subnets {
		opts, err := bind.NewKeyedTransactorWithChainID(deployerKey, subnetInfo.EVMChainID)
		Expect(err).Should(BeNil())
		teleporterRegistryAddress, tx, teleporterRegistry, err := teleporterregistry.DeployTeleporterRegistry(
			opts, subnetInfo.RPCClient, entries,
		)
		Expect(err).Should(BeNil())
		// Wait for the transaction to be mined
		utils.WaitForTransactionSuccess(ctx, subnetInfo, tx)

		if subnetInfo.SubnetID == constants.PrimaryNetworkID {
			n.primaryNetworkInfo.TeleporterRegistryAddress = teleporterRegistryAddress
			n.primaryNetworkInfo.TeleporterRegistry = teleporterRegistry
		} else {
			n.subnetsInfo[subnetInfo.SubnetID].TeleporterRegistryAddress = teleporterRegistryAddress
			n.subnetsInfo[subnetInfo.SubnetID].TeleporterRegistry = teleporterRegistry
		}

		log.Info("Deployed TeleporterRegistry contract to subnet", subnetInfo.SubnetID.Hex(),
			"address", teleporterRegistryAddress.Hex())
	}

	log.Info("Deployed TeleporterRegistry contracts to all subnets")
}

func (n *LocalNetwork) GetSubnetsInfo() []interfaces.SubnetTestInfo {
	return []interfaces.SubnetTestInfo{
		*n.subnetsInfo[n.subnetAID],
		*n.subnetsInfo[n.subnetBID],
	}
}

func (n *LocalNetwork) GetPrimaryNetworkInfo() interfaces.SubnetTestInfo {
	return *n.primaryNetworkInfo
}

// Returns subnet info for all subnets, including the primary network
func (n *LocalNetwork) GetAllSubnetsInfo() []interfaces.SubnetTestInfo {
	subnets := n.GetSubnetsInfo()
	return append(subnets, n.GetPrimaryNetworkInfo())
}

func (n *LocalNetwork) GetTeleporterContractAddress() common.Address {
	return n.teleporterContractAddress
}

func (n *LocalNetwork) SetTeleporterContractAddress(newTeleporterAddress common.Address) {
	n.teleporterContractAddress = newTeleporterAddress
	subnets := n.GetAllSubnetsInfo()
	for _, subnetInfo := range subnets {
		teleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
			n.teleporterContractAddress, subnetInfo.RPCClient,
		)
		Expect(err).Should(BeNil())
		if subnetInfo.SubnetID == constants.PrimaryNetworkID {
			n.primaryNetworkInfo.TeleporterMessenger = teleporterMessenger
		} else {
			n.subnetsInfo[subnetInfo.SubnetID].TeleporterMessenger = teleporterMessenger
		}
	}
}

func (n *LocalNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	fundedAddress := crypto.PubkeyToAddress(n.globalFundedKey.PublicKey)
	return fundedAddress, n.globalFundedKey
}

func (n *LocalNetwork) IsExternalNetwork() bool {
	return false
}

func (n *LocalNetwork) SupportsIndependentRelaying() bool {
	// Messages can be relayed by the test application for local
	// networks with connections to each node.
	return true
}

func (n *LocalNetwork) RelayMessage(ctx context.Context,
	sourceReceipt *types.Receipt,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
	expectSuccess bool,
) *types.Receipt {
	// Fetch the Teleporter message from the logs
	sendEvent, err := utils.GetEventFromLogs(sourceReceipt.Logs, source.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())

	signedWarpMessage := n.ConstructSignedWarpMessage(ctx, sourceReceipt, source, destination)

	// Construct the transaction to send the Warp message to the destination chain
	signedTx := utils.CreateReceiveCrossChainMessageTransaction(
		ctx,
		signedWarpMessage,
		sendEvent.Message.RequiredGasLimit,
		n.teleporterContractAddress,
		n.globalFundedKey,
		destination,
	)

	log.Info("Sending transaction to destination chain")
	if !expectSuccess {
		return utils.SendTransactionAndWaitForFailure(ctx, destination, signedTx)
	}

	receipt := utils.SendTransactionAndWaitForSuccess(ctx, destination, signedTx)

	// Check the transaction logs for the ReceiveCrossChainMessage event emitted by the Teleporter contract
	receiveEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		destination.TeleporterMessenger.ParseReceiveCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	Expect(receiveEvent.SourceBlockchainID[:]).Should(Equal(source.BlockchainID[:]))
	return receipt
}

func (n *LocalNetwork) setAllSubnetValues() {
	subnetIDs := n.manager.GetSubnets()
	Expect(len(subnetIDs)).Should(Equal(2))

	n.subnetAID = subnetIDs[0]
	n.setSubnetValues(n.subnetAID)

	n.subnetBID = subnetIDs[1]
	n.setSubnetValues(n.subnetBID)

	n.setPrimaryNetworkValues()
}

func (n *LocalNetwork) TearDownNetwork() {
	log.Info("Tearing down network")
	Expect(n.manager).ShouldNot(BeNil())
	Expect(n.manager.TeardownNetwork()).Should(BeNil())
	Expect(os.Remove(n.warpChainConfigPath)).Should(BeNil())
}

func (n *LocalNetwork) AddSubnetValidators(ctx context.Context, subnetID ids.ID, nodeNames []string) {
	_, err := n.anrClient.AddSubnetValidators(ctx, []*rpcpb.SubnetValidatorsSpec{
		{
			SubnetId:  subnetID.String(),
			NodeNames: nodeNames,
		},
	})
	Expect(err).Should(BeNil())

	// Add the new node names
	n.subnetNodeNames[subnetID] = append(n.subnetNodeNames[subnetID], nodeNames...)

	n.setAllSubnetValues()
}

// GetAllNodeNames returns a slice that copies all node names in the network
func (n *LocalNetwork) GetAllNodeNames() []string {
	// The network starts off with 5 nodes that validate the primary network.
	// These nodes were not added by this network setup, and are not in n.subnetNodeNames.
	// So we query the ANR client to get the full list of node names.
	status, err := n.anrClient.Status(context.Background())
	Expect(err).Should(BeNil())
	return status.GetClusterInfo().GetNodeNames()
}

func (n *LocalNetwork) RestartNodes(ctx context.Context, nodeNames []string, opts ...runner_sdk.OpOption) {
	log.Info("Network restarting nodes", "nodeNames", nodeNames)
	opts = append(opts,
		runner_sdk.WithExecPath(n.manager.ANRConfig.AvalancheGoExecPath),
		runner_sdk.WithPluginDir(n.manager.ANRConfig.PluginDir))
	for _, nodeName := range nodeNames {
		_, err := n.anrClient.RestartNode(ctx, nodeName, opts...)
		Expect(err).Should(BeNil())
	}

	log.Info("Waiting for all VMs to report healthy")
	for {
		v, err := n.anrClient.Health(ctx)
		log.Info("Pinged CLI Health", "result", v, "err", err)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	n.setAllSubnetValues()
}

func (n *LocalNetwork) ConstructSignedWarpMessage(
	ctx context.Context,
	sourceReceipt *types.Receipt,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
) *avalancheWarp.Message {
	log.Info("Fetching relevant warp logs from the newly produced block")
	logs, err := source.RPCClient.FilterLogs(ctx, subnetEvmInterfaces.FilterQuery{
		BlockHash: &sourceReceipt.BlockHash,
		Addresses: []common.Address{warp.Module.Address},
	})
	Expect(err).Should(BeNil())
	Expect(len(logs)).Should(Equal(1))

	// Check for relevant warp log from subscription and ensure that it matches
	// the log extracted from the last block.
	txLog := logs[0]
	log.Info("Parsing logData as unsigned warp message")
	unsignedMsg, err := warp.UnpackSendWarpEventDataToMessage(txLog.Data)
	Expect(err).Should(BeNil())

	// Set local variables for the duration of the test
	unsignedWarpMessageID := unsignedMsg.ID()
	log.Info(
		"Parsed unsignedWarpMsg",
		"unsignedWarpMessageID", unsignedWarpMessageID,
		"unsignedWarpMessage", unsignedMsg,
	)

	// Loop over each client on source chain to ensure they all have time to accept the block.
	// Note: if we did not confirm this here, the next stage could be racy since it assumes every node
	// has accepted the block.
	waitForAllValidatorsToAcceptBlock(ctx, source.NodeURIs, source.BlockchainID, sourceReceipt.BlockNumber.Uint64())

	// Get the aggregate signature for the Warp message
	log.Info("Fetching aggregate signature from the source chain validators")
	return n.GetSignedMessage(ctx, source, destination, unsignedWarpMessageID)
}

func (n *LocalNetwork) GetSignedMessage(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
	unsignedWarpMessageID ids.ID,
) *avalancheWarp.Message {
	Expect(len(source.NodeURIs)).Should(BeNumerically(">", 0))
	warpClient, err := warpBackend.NewClient(source.NodeURIs[0], source.BlockchainID.String())
	Expect(err).Should(BeNil())

	signingSubnetID := source.SubnetID
	if source.SubnetID == constants.PrimaryNetworkID {
		signingSubnetID = destination.SubnetID
	}

	// Get the aggregate signature for the Warp message
	signedWarpMessageBytes, err := warpClient.GetMessageAggregateSignature(
		ctx,
		unsignedWarpMessageID,
		warp.WarpDefaultQuorumNumerator,
		signingSubnetID.String(),
	)
	Expect(err).Should(BeNil())

	signedWarpMsg, err := avalancheWarp.ParseMessage(signedWarpMessageBytes)
	Expect(err).Should(BeNil())

	return signedWarpMsg
}

func (n *LocalNetwork) GetNetworkID() uint32 {
	status, err := n.anrClient.Status(context.Background())
	Expect(err).Should(BeNil())
	return status.GetClusterInfo().GetNetworkId()
}
