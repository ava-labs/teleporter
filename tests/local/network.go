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
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/plugin/evm"
	"github.com/ava-labs/subnet-evm/rpc"
	"github.com/ava-labs/subnet-evm/tests/utils/runner"
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

var _ interfaces.Network = &localNetwork{}

// Implements Network, pointing to the network setup in local_network_setup.go
type localNetwork struct {
	teleporterContractAddress       common.Address
	subnetAID, subnetBID, subnetCID ids.ID
	subnetsInfo                     map[ids.ID]*interfaces.SubnetTestInfo
	subnetNodeNames                 map[ids.ID][]string

	globalFundedKey *ecdsa.PrivateKey

	// Internal vars only used to set up the local network
	anrClient           runner_sdk.Client
	manager             *runner.NetworkManager
	warpChainConfigPath string
}

const (
	fundedKeyStr = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
)

func newLocalNetwork(warpGenesisFile string) *localNetwork {
	ctx := context.Background()
	var err error

	// Name 10 new validators (which should have BLS key registered)
	var subnetANodeNames []string
	var subnetBNodeNames []string
	var subnetCNodeNames []string
	for i := 1; i <= 15; i++ {
		n := fmt.Sprintf("node%d-bls", i)
		if i <= 5 {
			subnetANodeNames = append(subnetANodeNames, n)
		} else if i <= 10 {
			subnetBNodeNames = append(subnetBNodeNames, n)
		} else {
			subnetCNodeNames = append(subnetCNodeNames, n)
		}
	}

	f, err := os.CreateTemp(os.TempDir(), "config.json")
	Expect(err).Should(BeNil())
	_, err = f.Write([]byte(`{"warp-api-enabled": true}`))
	Expect(err).Should(BeNil())
	warpChainConfigPath := f.Name()

	// Make sure that the warp genesis file exists
	_, err = os.Stat(warpGenesisFile)
	Expect(err).Should(BeNil())

	anrConfig := runner.NewDefaultANRConfig()
	log.Info("dbg", "anrConfig", anrConfig)
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
			{
				VmName:      evm.IDStr,
				Genesis:     warpGenesisFile,
				ChainConfig: warpChainConfigPath,
				SubnetSpec: &rpcpb.SubnetSpec{
					SubnetConfig: "",
					Participants: subnetCNodeNames,
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
	setupProposerVM(ctx, globalFundedKey, manager, 2)

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
	Expect(len(subnetIDs)).Should(Equal(3))
	subnetAID := subnetIDs[0]
	subnetBID := subnetIDs[1]
	subnetCID := subnetIDs[2]

	res := &localNetwork{
		subnetAID:   subnetAID,
		subnetBID:   subnetBID,
		subnetCID:   subnetCID,
		subnetsInfo: make(map[ids.ID]*interfaces.SubnetTestInfo),
		subnetNodeNames: map[ids.ID][]string{
			subnetAID: subnetANodeNames,
			subnetBID: subnetBNodeNames,
			subnetCID: subnetCNodeNames,
		},
		globalFundedKey:     globalFundedKey,
		anrClient:           anrClient,
		manager:             manager,
		warpChainConfigPath: warpChainConfigPath,
	}
	res.setSubnetValues(subnetAID)
	res.setSubnetValues(subnetBID)
	res.setSubnetValues(subnetCID)
	return res
}

func (n *localNetwork) setSubnetValues(subnetID ids.ID) {
	subnetDetails, ok := n.manager.GetSubnet(subnetID)
	Expect(ok).Should(BeTrue())
	blockchainID := subnetDetails.BlockchainID

	// Reset the validator URIs, as they may have changed
	subnetDetails.ValidatorURIs = nil
	status, err := n.anrClient.Status(context.Background())
	Expect(err).Should(BeNil())
	nodeInfos := status.GetClusterInfo().GetNodeInfos()

	for _, nodeName := range n.subnetNodeNames[subnetID] {
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

// deployTeleporterContracts deploys the Teleporter contract to all subnets.
// The caller is responsible for generating the deployment transaction information
func (n *localNetwork) deployTeleporterContracts(
	transactionBytes []byte,
	deployerAddress common.Address,
	contractAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
) {
	log.Info("Deploying Teleporter contract to subnets")

	subnetsInfoList := n.GetSubnetsInfo()

	// Set the package level teleporterContractAddress
	n.teleporterContractAddress = contractAddress

	ctx := context.Background()

	for _, subnetInfo := range subnetsInfoList {
		// Fund the deployer address
		{
			fundAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)) // 10eth
			fundDeployerTx := utils.CreateNativeTransferTransaction(
				ctx, subnetInfo, fundedKey, deployerAddress, fundAmount,
			)
			utils.SendTransactionAndWaitForAcceptance(ctx, subnetInfo, fundDeployerTx, true)
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
			teleporterCode, err := subnetInfo.RPCClient.CodeAt(ctx, n.teleporterContractAddress, nil)
			Expect(err).Should(BeNil())
			Expect(len(teleporterCode)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
		}
		teleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
			n.teleporterContractAddress, subnetInfo.RPCClient,
		)
		Expect(err).Should(BeNil())
		n.subnetsInfo[subnetInfo.SubnetID].TeleporterMessenger = teleporterMessenger
		log.Info("Finished deploying Teleporter contract", "blockchainID", subnetInfo.BlockchainID.Hex())
	}
	log.Info("Deployed Teleporter contracts to all subnets")
}

func (n *localNetwork) deployTeleporterRegistryContracts(
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

	for _, subnetInfo := range n.GetSubnetsInfo() {
		opts, err := bind.NewKeyedTransactorWithChainID(deployerKey, subnetInfo.EVMChainID)
		Expect(err).Should(BeNil())
		teleporterRegistryAddress, tx, _, err := teleporterregistry.DeployTeleporterRegistry(
			opts, subnetInfo.RPCClient, entries,
		)
		Expect(err).Should(BeNil())

		n.subnetsInfo[subnetInfo.SubnetID].TeleporterRegistryAddress = teleporterRegistryAddress
		// Wait for the transaction to be mined
		receipt, err := bind.WaitMined(ctx, subnetInfo.RPCClient, tx)
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
		log.Info("Deployed TeleporterRegistry contract to subnet", subnetInfo.SubnetID.Hex(),
			"Deploy address", teleporterRegistryAddress.Hex())
	}

	log.Info("Deployed TeleporterRegistry contracts to all subnets")
}

func (n *localNetwork) GetSubnetsInfo() []interfaces.SubnetTestInfo {
	return []interfaces.SubnetTestInfo{
		*n.subnetsInfo[n.subnetAID],
		*n.subnetsInfo[n.subnetBID],
		*n.subnetsInfo[n.subnetCID],
	}
}

func (n *localNetwork) GetTeleporterContractAddress() common.Address {
	return n.teleporterContractAddress
}

func (n *localNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	fundedAddress := crypto.PubkeyToAddress(n.globalFundedKey.PublicKey)
	return fundedAddress, n.globalFundedKey
}

func (n *localNetwork) SupportsIndependentRelaying() bool {
	// Messages can be relayed by the test application for local
	// networks with connections to each node.
	return true
}

func (n *localNetwork) RelayMessage(ctx context.Context,
	sourceReceipt *types.Receipt,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
	expectSuccess bool) *types.Receipt {
	return relayMessage(
		ctx,
		n.teleporterContractAddress,
		n.globalFundedKey,
		sourceReceipt,
		source,
		destination,
		expectSuccess)
}

func (n *localNetwork) setAllSubnetValues() {
	subnetIDs := n.manager.GetSubnets()
	Expect(len(subnetIDs)).Should(Equal(3))

	n.subnetAID = subnetIDs[0]
	n.setSubnetValues(n.subnetAID)

	n.subnetBID = subnetIDs[1]
	n.setSubnetValues(n.subnetBID)

	n.subnetCID = subnetIDs[2]
	n.setSubnetValues(n.subnetCID)
}

func (n *localNetwork) tearDownNetwork() {
	log.Info("Tearing down network")
	Expect(n.manager).ShouldNot(BeNil())
	Expect(n.manager.TeardownNetwork()).Should(BeNil())
	Expect(os.Remove(n.warpChainConfigPath)).Should(BeNil())
}

func (n *localNetwork) addSubnetValidators(ctx context.Context, subnetID ids.ID, nodeNames []string) {
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
