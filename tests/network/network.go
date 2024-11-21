package network

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	goLog "log"
	"os"
	"sort"
	"time"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/config"
	"github.com/ava-labs/avalanchego/genesis"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	warpMessage "github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	pwallet "github.com/ava-labs/avalanchego/wallet/chain/p/wallet"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary"
	"github.com/ava-labs/awm-relayer/signature-aggregator/aggregator"
	"github.com/ava-labs/subnet-evm/ethclient"
	subnetEvmTestUtils "github.com/ava-labs/subnet-evm/tests/utils"
	proxyadmin "github.com/ava-labs/teleporter/abi-bindings/go/ProxyAdmin"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

// Implements Network, pointing to the network setup in local_network_setup.go
type LocalNetwork struct {
	tmpnet.Network

	extraNodes               []*tmpnet.Node // to add as more subnet validators in the tests
	primaryNetworkValidators []*tmpnet.Node
	globalFundedKey          *secp256k1.PrivateKey
	validatorManagers        map[ids.ID]common.Address
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
	numPrimaryNetworkValidators int,
	extraNodeCount int, // for use by tests, eg to add new subnet validators
) *LocalNetwork {
	// There must be at least one primary network validator per subnet
	Expect(numPrimaryNetworkValidators).Should(BeNumerically(">=", len(subnetSpecs)))

	// Create extra nodes to be used to add more validators later
	extraNodes := subnetEvmTestUtils.NewTmpnetNodes(extraNodeCount)

	fundedKey, err := hex.DecodeString(fundedKeyStr)
	Expect(err).Should(BeNil())
	globalFundedKey, err := secp256k1.ToPrivateKey(fundedKey)
	Expect(err).Should(BeNil())

	globalFundedECDSAKey := globalFundedKey.ToECDSA()
	Expect(err).Should(BeNil())

	var subnets []*tmpnet.Subnet
	bootstrapNodes := subnetEvmTestUtils.NewTmpnetNodes(numPrimaryNetworkValidators)
	for i, subnetSpec := range subnetSpecs {
		// Create a single bootstrap node. This will be removed from the subnet validator set after it is converted,
		// but will remain a primary network validator
		initialSubnetBootstrapper := bootstrapNodes[i] // One bootstrap node per subnet

		// Create validators to specify as the initial vdr set in the subnet conversion, and store them as extra nodes
		initialVdrNodes := subnetEvmTestUtils.NewTmpnetNodes(subnetSpec.NodeCount)
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
			initialSubnetBootstrapper,
		)
		subnet.OwningKey = globalFundedKey
		subnets = append(subnets, subnet)
	}
	network := subnetEvmTestUtils.NewTmpnetNetwork(
		name,
		bootstrapNodes,
		utils.WarpEnabledChainConfig,
		subnets...,
	)
	Expect(network).ShouldNot(BeNil())

	// Activate Etna
	// upgrades := upgrade.Default
	// upgrades.EtnaTime = time.Now().Add(-1 * time.Minute)
	// upgradeJSON, err := json.Marshal(upgrades)
	// Expect(err).Should(BeNil())

	// upgradeBase64 := base64.StdEncoding.EncodeToString(upgradeJSON)
	// network.DefaultFlags.SetDefaults(tmpnet.FlagsMap{
	// 	config.UpgradeFileContentKey: upgradeBase64,
	// })

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
	goLog.Println("Network bootstrapped")

	// Issue transactions to activate the proposerVM fork on the chains
	for _, subnet := range network.Subnets {
		utils.SetupProposerVM(ctx, globalFundedECDSAKey, network, subnet.SubnetID)
	}

	// All nodes are specified as bootstrap validators
	var primaryNetworkValidators []*tmpnet.Node
	fmt.Println("Primary Network Validators")
	for _, node := range network.Nodes {
		fmt.Println(node.NodeID, node.URI)
		primaryNetworkValidators = append(primaryNetworkValidators, node)
	}

	localNetwork := &LocalNetwork{
		Network:                  *network,
		extraNodes:               extraNodes,
		globalFundedKey:          globalFundedKey,
		primaryNetworkValidators: primaryNetworkValidators,
		validatorManagers:        make(map[ids.ID]common.Address),
	}

	return localNetwork
}

func (n *LocalNetwork) ConvertSubnet(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	managerType utils.ValidatorManagerConcreteType,
	weights []uint64,
	senderKey *ecdsa.PrivateKey,
	proxy bool,
) ([]utils.Node, []ids.ID, *proxyadmin.ProxyAdmin) {
	goLog.Println("Converting subnet", subnet.SubnetID)
	cChainInfo := n.GetPrimaryNetworkInfo()
	pClient := platformvm.NewClient(cChainInfo.NodeURIs[0])
	currentValidators, err := pClient.GetCurrentValidators(ctx, subnet.SubnetID, nil)
	Expect(err).Should(BeNil())

	vdrManagerAddress, proxyAdmin := utils.DeployAndInitializeValidatorManager(
		ctx,
		senderKey,
		subnet,
		managerType,
		proxy,
	)
	n.validatorManagers[subnet.SubnetID] = vdrManagerAddress

	tmpnetNodes := n.GetExtraNodes(len(weights))
	sort.Slice(tmpnetNodes, func(i, j int) bool {
		return string(tmpnetNodes[i].NodeID.Bytes()) < string(tmpnetNodes[j].NodeID.Bytes())
	})

	var nodes []utils.Node
	// Construct the convert subnet info
	destAddr, err := address.ParseToID(utils.DefaultPChainAddress)
	Expect(err).Should(BeNil())
	vdrs := make([]*txs.ConvertSubnetToL1Validator, len(tmpnetNodes))
	for i, node := range tmpnetNodes {
		signer, err := node.GetProofOfPossession()
		Expect(err).Should(BeNil())
		nodes = append(nodes, utils.Node{
			NodeID:  node.NodeID,
			NodePoP: signer,
			Weight:  weights[i],
		})
		vdrs[i] = &txs.ConvertSubnetToL1Validator{
			NodeID:  node.NodeID.Bytes(),
			Weight:  weights[i],
			Balance: units.Avax * 100,
			Signer:  *signer,
			RemainingBalanceOwner: warpMessage.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
			DeactivationOwner: warpMessage.PChainOwner{
				Threshold: 1,
				Addresses: []ids.ShortID{destAddr},
			},
		}
	}
	pChainWallet := n.GetPChainWallet()
	_, err = pChainWallet.IssueConvertSubnetToL1Tx(
		subnet.SubnetID,
		subnet.BlockchainID,
		vdrManagerAddress[:],
		vdrs,
	)
	Expect(err).Should(BeNil())

	subnet = n.AddSubnetValidators(tmpnetNodes, subnet)

	utils.PChainProposerVMWorkaround(pChainWallet)
	utils.AdvanceProposerVM(ctx, subnet, senderKey, 5)

	aggregator := n.GetSignatureAggregator()
	validationIDs := utils.InitializeValidatorSet(
		ctx,
		senderKey,
		subnet,
		utils.GetPChainInfo(cChainInfo),
		vdrManagerAddress,
		n.GetNetworkID(),
		aggregator,
		nodes,
	)
	aggregator.Shutdown()

	// Remove the bootstrap nodes as subnet validators
	for _, vdr := range currentValidators {
		_, err := pChainWallet.IssueRemoveSubnetValidatorTx(vdr.NodeID, subnet.SubnetID)
		Expect(err).Should(BeNil())
		for _, node := range n.Network.Nodes {
			if node.NodeID == vdr.NodeID {
				fmt.Println("Restarting bootstrap node", node.NodeID)
				n.Network.RestartNode(ctx, os.Stdout, node)
			}
		}
	}
	utils.PChainProposerVMWorkaround(pChainWallet)
	utils.AdvanceProposerVM(ctx, subnet, senderKey, 5)

	return nodes, validationIDs, proxyAdmin
}

func (n *LocalNetwork) AddSubnetValidators(
	nodes []*tmpnet.Node,
	subnet interfaces.SubnetTestInfo,
) interfaces.SubnetTestInfo {
	// Modify the each node's config to track the subnet
	for _, node := range nodes {
		goLog.Printf("Adding node %s @ %s to subnet %s", node.NodeID, node.URI, subnet.SubnetID)
		existingTrackedSubnets, err := node.Flags.GetStringVal(config.TrackSubnetsKey)
		Expect(err).Should(BeNil())
		if existingTrackedSubnets == subnet.SubnetID.String() {
			goLog.Printf("Node %s @ %s already tracking subnet %s\n", node.NodeID, node.URI, subnet.SubnetID)
			continue
		}
		node.Flags[config.TrackSubnetsKey] = subnet.SubnetID.String()

		// Add the node to the network
		n.Network.Nodes = append(n.Network.Nodes, node)
	}
	err := n.Network.StartNodes(context.Background(), os.Stdout, nodes...)
	Expect(err).Should(BeNil())

	// Update the tmpnet Subnet struct
	for _, tmpnetSubnet := range n.Network.Subnets {
		if tmpnetSubnet.SubnetID == subnet.SubnetID {
			for _, tmpnetNode := range nodes {
				tmpnetSubnet.ValidatorIDs = append(tmpnetSubnet.ValidatorIDs, tmpnetNode.NodeID)
			}
		}
	}

	// Refresh the subnet info after restarting the nodes
	return n.GetSubnetInfo(subnet.SubnetID)
}

func (n *LocalNetwork) GetValidatorManager(subnetID ids.ID) common.Address {
	return n.validatorManagers[subnetID]
}

func (n *LocalNetwork) GetSignatureAggregator() *aggregator.SignatureAggregator {
	var subnetIDs []ids.ID
	for _, subnet := range n.GetSubnetsInfo() {
		subnetIDs = append(subnetIDs, subnet.SubnetID)
	}
	return utils.NewSignatureAggregator(
		n.GetPrimaryNetworkInfo().NodeURIs[0],
		subnetIDs,
	)
}

func (n *LocalNetwork) GetExtraNodes(count int) []*tmpnet.Node {
	Expect(len(n.extraNodes) >= count).Should(
		BeTrue(),
		"not enough extra nodes to use",
	)
	nodes := n.extraNodes[0:count]
	n.extraNodes = n.extraNodes[count:]
	return nodes
}

func (n *LocalNetwork) GetPrimaryNetworkValidators() []*tmpnet.Node {
	return n.primaryNetworkValidators
}

func (n *LocalNetwork) GetPrimaryNetworkInfo() interfaces.SubnetTestInfo {
	var nodeURIs []string
	for _, node := range n.primaryNetworkValidators {
		nodeURIs = append(nodeURIs, node.URI)
	}
	infoClient := info.NewClient(nodeURIs[0])
	cChainBlockchainID, err := infoClient.GetBlockchainID(context.Background(), "C")
	Expect(err).Should(BeNil())

	wsClient, err := ethclient.Dial(utils.HttpToWebsocketURI(nodeURIs[0], cChainBlockchainID.String()))
	Expect(err).Should(BeNil())

	rpcClient, err := ethclient.Dial(utils.HttpToRPCURI(nodeURIs[0], cChainBlockchainID.String()))
	Expect(err).Should(BeNil())

	evmChainID, err := rpcClient.ChainID(context.Background())
	Expect(err).Should(BeNil())
	return interfaces.SubnetTestInfo{
		SubnetID:     ids.Empty,
		BlockchainID: cChainBlockchainID,
		NodeURIs:     nodeURIs,
		WSClient:     wsClient,
		RPCClient:    rpcClient,
		EVMChainID:   evmChainID,
	}
}

func (n *LocalNetwork) GetSubnetInfo(subnetID ids.ID) interfaces.SubnetTestInfo {
	for _, subnet := range n.Network.Subnets {
		if subnet.SubnetID == subnetID {
			var nodeURIs []string
			for _, nodeID := range subnet.ValidatorIDs {
				uri, err := n.Network.GetURIForNodeID(nodeID)
				Expect(err).Should(BeNil())

				nodeURIs = append(nodeURIs, uri)
			}
			blockchainID := subnet.Chains[0].ChainID
			wsClient, err := ethclient.Dial(utils.HttpToWebsocketURI(nodeURIs[0], blockchainID.String()))
			Expect(err).Should(BeNil())

			rpcClient, err := ethclient.Dial(utils.HttpToRPCURI(nodeURIs[0], blockchainID.String()))
			Expect(err).Should(BeNil())
			evmChainID, err := rpcClient.ChainID(context.Background())
			Expect(err).Should(BeNil())
			return interfaces.SubnetTestInfo{
				SubnetID:     subnetID,
				BlockchainID: blockchainID,
				NodeURIs:     nodeURIs,
				WSClient:     wsClient,
				RPCClient:    rpcClient,
				EVMChainID:   evmChainID,
			}
		}
	}
	return interfaces.SubnetTestInfo{}
}

// Returns all subnet info sorted in lexicographic order of SubnetName.
func (n *LocalNetwork) GetSubnetsInfo() []interfaces.SubnetTestInfo {
	subnets := make([]interfaces.SubnetTestInfo, len(n.Network.Subnets))
	for i, subnet := range n.Network.Subnets {
		var nodeURIs []string
		for _, nodeID := range subnet.ValidatorIDs {
			uri, err := n.Network.GetURIForNodeID(nodeID)
			Expect(err).Should(BeNil())

			nodeURIs = append(nodeURIs, uri)
		}
		blockchainID := subnet.Chains[0].ChainID
		wsClient, err := ethclient.Dial(utils.HttpToWebsocketURI(nodeURIs[0], blockchainID.String()))
		Expect(err).Should(BeNil())

		rpcClient, err := ethclient.Dial(utils.HttpToRPCURI(nodeURIs[0], blockchainID.String()))
		Expect(err).Should(BeNil())
		evmChainID, err := rpcClient.ChainID(context.Background())
		Expect(err).Should(BeNil())
		subnets[i] = interfaces.SubnetTestInfo{
			SubnetID:     subnet.SubnetID,
			BlockchainID: blockchainID,
			NodeURIs:     nodeURIs,
			WSClient:     wsClient,
			RPCClient:    rpcClient,
			EVMChainID:   evmChainID,
		}
	}
	return subnets
}

// Returns subnet info for all subnets, including the primary network
func (n *LocalNetwork) GetAllSubnetsInfo() []interfaces.SubnetTestInfo {
	subnets := n.GetSubnetsInfo()
	return append(subnets, n.GetPrimaryNetworkInfo())
}

func (n *LocalNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	ecdsaKey := n.globalFundedKey.ToECDSA()
	fundedAddress := crypto.PubkeyToAddress(ecdsaKey.PublicKey)
	return fundedAddress, ecdsaKey
}

func (n *LocalNetwork) TearDownNetwork() {
	log.Info("Tearing down network")
	Expect(n).ShouldNot(BeNil())
	Expect(n.Network).ShouldNot(BeNil())
	Expect(n.Network.Stop(context.Background())).Should(BeNil())
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
			n.Network.ChainConfigs[utils.CChainPathSpecifier] = cfg
			continue
		}

		for _, subnet := range n.Network.Subnets {
			for _, chain := range subnet.Chains {
				if chain.ChainID.String() == chainIDStr {
					chain.Config = chainConfig
				}
			}
		}
	}
	err := n.Network.Write()
	if err != nil {
		log.Error("failed to write network", "error", err)
	}
	for _, subnet := range n.Network.Subnets {
		err := subnet.Write(n.Network.GetSubnetDir(), n.Network.GetChainConfigDir())
		if err != nil {
			log.Error("failed to write subnets", "error", err)
		}
	}

	// Restart the network to apply the new chain configs
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60*len(n.Network.Nodes))*time.Second)
	defer cancel()
	err = n.Network.Restart(ctx, os.Stdout)
	Expect(err).Should(BeNil())
}

func (n *LocalNetwork) GetNetworkID() uint32 {
	return n.Network.Genesis.NetworkID
}

func (n *LocalNetwork) Dir() string {
	return n.Network.Dir
}

func (n *LocalNetwork) GetPChainWallet() pwallet.Wallet {
	// Create the P-Chain wallet to issue transactions
	kc := secp256k1fx.NewKeychain(n.globalFundedKey)
	var subnetIDs []ids.ID
	for _, subnet := range n.GetSubnetsInfo() {
		subnetIDs = append(subnetIDs, subnet.SubnetID)
	}
	wallet, err := primary.MakeWallet(context.Background(), &primary.WalletConfig{
		URI:          n.GetPrimaryNetworkInfo().NodeURIs[0],
		AVAXKeychain: kc,
		EthKeychain:  kc,
		SubnetIDs:    subnetIDs,
	})
	Expect(err).Should(BeNil())
	return wallet.P()
}

func (n *LocalNetwork) GetTwoSubnets() (
	interfaces.SubnetTestInfo,
	interfaces.SubnetTestInfo,
) {
	subnets := n.GetSubnetsInfo()
	Expect(len(subnets)).Should(BeNumerically(">=", 2))
	return subnets[0], subnets[1]
}
