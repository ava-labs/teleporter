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
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/plugin/evm"
	"github.com/ava-labs/subnet-evm/rpc"
	"github.com/ava-labs/subnet-evm/tests/utils/runner"
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

const (
	fundedKeyStr     = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
	fundedAddressStr = "0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC"
)

// Implements Network, pointing to the network setup in local_network_setup.go
type localNetwork struct {
	teleporterContractAddress common.Address
	subnetAInfo, subnetBInfo  interfaces.SubnetTestInfo
	fundedAddress             common.Address
	fundedKey                 *ecdsa.PrivateKey

	anrClient           runner_sdk.Client
	manager             *runner.NetworkManager
	warpChainConfigPath string
}

func newLocalNetwork(warpGenesisFile string) *localNetwork {
	anrConfig := runner.NewDefaultANRConfig()
	manager := runner.NewNetworkManager(anrConfig)
	Expect(manager).ShouldNot(BeNil())

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
	warpChainConfigPath := f.Name()

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

	// Issue transactions to activate the proposerVM fork on the chains
	fundedKey, err := crypto.HexToECDSA(fundedKeyStr)
	Expect(err).Should(BeNil())
	setupProposerVM(ctx, fundedKey, manager, 0)
	setupProposerVM(ctx, fundedKey, manager, 1)

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

	// Set up subnet URIs
	subnetIDs := manager.GetSubnets()
	Expect(len(subnetIDs)).Should(Equal(2))

	subnetAID := subnetIDs[0]
	subnetADetails, ok := manager.GetSubnet(subnetAID)
	Expect(ok).Should(BeTrue())
	Expect(len(subnetADetails.ValidatorURIs)).Should(Equal(5))
	blockchainIDA := subnetADetails.BlockchainID
	chainANodeURIs := append([]string{}, subnetADetails.ValidatorURIs...)

	subnetBID := subnetIDs[1]
	subnetBDetails, ok := manager.GetSubnet(subnetBID)
	Expect(ok).Should(BeTrue())
	Expect(len(subnetBDetails.ValidatorURIs)).Should(Equal(5))
	blockchainIDB := subnetBDetails.BlockchainID
	chainBNodeURIs := append([]string{}, subnetBDetails.ValidatorURIs...)

	log.Info(
		"Created URIs for subnets",
		"chainAURIs", chainANodeURIs,
		"chainBURIs", chainBNodeURIs,
		"blockchainIDA", blockchainIDA,
		"blockchainIDB", blockchainIDB,
	)

	chainAWSURI := utils.HttpToWebsocketURI(chainANodeURIs[0], blockchainIDA.String())
	chainARPCURI := utils.HttpToRPCURI(chainANodeURIs[0], blockchainIDA.String())

	log.Info("Creating ethclients for blockchainA", "wsURI", chainAWSURI, "rpcURI", chainARPCURI)
	chainAWSClient, err := ethclient.Dial(chainAWSURI)
	Expect(err).Should(BeNil())
	chainARPCClient, err := ethclient.Dial(chainARPCURI)
	Expect(err).Should(BeNil())
	chainAIDInt, err := chainARPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	chainBWSURI := utils.HttpToWebsocketURI(chainBNodeURIs[0], blockchainIDB.String())
	chainBRPCURI := utils.HttpToRPCURI(chainBNodeURIs[0], blockchainIDB.String())
	log.Info("Creating ethclients for blockchainB", "wsURI", chainBWSURI, "rpcURI", chainBRPCURI)

	chainBWSClient, err := ethclient.Dial(chainBWSURI)
	Expect(err).Should(BeNil())
	chainBRPCClient, err := ethclient.Dial(chainBRPCURI)
	Expect(err).Should(BeNil())
	chainBIDInt, err := chainBRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	log.Info("Finished setting up e2e test subnet variables")
	return &localNetwork{
		teleporterContractAddress: common.Address{}, // Set by deployTeleporterContracts
		subnetAInfo: interfaces.SubnetTestInfo{
			SubnetID:                  subnetAID,
			BlockchainID:              blockchainIDA,
			NodeURIs:                  chainANodeURIs,
			NodeNames:                 subnetANodeNames,
			WSClient:                  chainAWSClient,
			RPCClient:                 chainARPCClient,
			EVMChainID:                chainAIDInt,
			TeleporterRegistryAddress: common.Address{}, // Set by deployTeleporterRegistryContracts
		},
		subnetBInfo: interfaces.SubnetTestInfo{
			SubnetID:                  subnetAID,
			BlockchainID:              blockchainIDB,
			NodeURIs:                  chainBNodeURIs,
			NodeNames:                 subnetBNodeNames,
			WSClient:                  chainBWSClient,
			RPCClient:                 chainBRPCClient,
			EVMChainID:                chainBIDInt,
			TeleporterRegistryAddress: common.Address{}, // Set by deployTeleporterRegistryContracts,
		},
		fundedAddress:       common.HexToAddress(fundedAddressStr),
		fundedKey:           fundedKey,
		anrClient:           anrClient,
		manager:             manager,
		warpChainConfigPath: warpChainConfigPath,
	}
}

func (n *localNetwork) deployTeleporterContracts(
	transactionBytes []byte,
	deployerAddress common.Address,
	contractAddress common.Address) {
	log.Info("Deploying Teleporter contract to subnets")

	// Set the package level teleporterContractAddress
	n.teleporterContractAddress = contractAddress

	ctx := context.Background()

	for _, subnetInfo := range []interfaces.SubnetTestInfo{n.subnetAInfo, n.subnetBInfo} {
		// Fund the deployer address
		{
			fundAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)) // 10eth
			fundDeployerTx := utils.CreateNativeTransferTransaction(
				ctx,
				subnetInfo,
				n.fundedAddress,
				n.fundedKey,
				deployerAddress,
				fundAmount)
			utils.SendTransactionAndWaitForAcceptance(ctx, subnetInfo.RPCClient, fundDeployerTx, true)
		}
		log.Info("Finished funding Teleporter deployer", "blockchainID", subnetInfo.BlockchainID.Hex())

		// Deploy Teleporter contract
		{
			rpcClient, err := rpc.DialContext(ctx, utils.HttpToRPCURI(subnetInfo.NodeURIs[0], subnetInfo.BlockchainID.String()))
			Expect(err).Should(BeNil())
			defer rpcClient.Close()

			newHeads := make(chan *types.Header, 10)
			subA, err := subnetInfo.WSClient.SubscribeNewHead(ctx, newHeads)
			Expect(err).Should(BeNil())
			defer subA.Unsubscribe()

			err = rpcClient.CallContext(ctx, nil, "eth_sendRawTransaction", hexutil.Encode(transactionBytes))
			Expect(err).Should(BeNil())

			<-newHeads
			teleporterCode, err := subnetInfo.RPCClient.CodeAt(ctx, n.teleporterContractAddress, nil)
			Expect(err).Should(BeNil())
			Expect(len(teleporterCode)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
		}
		log.Info("Finished deploying Teleporter contract", "blockchainID", subnetInfo.BlockchainID.Hex())
	}
	log.Info("Deployed Teleporter contracts to all subnets")
}

func (n *localNetwork) deployTeleporterRegistryContracts(teleporterAddress common.Address) {
	log.Info("Deploying TeleporterRegistry contract to subnets")
	ctx := context.Background()

	entries := []teleporterregistry.ProtocolRegistryEntry{
		{
			Version:         big.NewInt(1),
			ProtocolAddress: teleporterAddress,
		},
	}

	var (
		err error
		tx  *types.Transaction
	)
	optsA := utils.CreateTransactorOpts(ctx, n.subnetAInfo, n.fundedAddress, n.fundedKey)
	teleporterRegistryAddressA, tx, _, err := teleporterregistry.DeployTeleporterRegistry(
		optsA,
		n.subnetAInfo.RPCClient,
		entries)
	Expect(err).Should(BeNil())
	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, n.subnetAInfo.RPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	optsB := utils.CreateTransactorOpts(ctx, n.subnetBInfo, n.fundedAddress, n.fundedKey)
	teleporterRegistryAddressB, tx, _, err := teleporterregistry.DeployTeleporterRegistry(
		optsB,
		n.subnetBInfo.RPCClient,
		entries)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err = bind.WaitMined(ctx, n.subnetBInfo.RPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	log.Info("Deployed TeleporterRegistry contracts to all subnets")
	n.subnetAInfo.TeleporterRegistryAddress = teleporterRegistryAddressA
	n.subnetBInfo.TeleporterRegistryAddress = teleporterRegistryAddressB
}

func (n *localNetwork) GetSubnetInfo() (interfaces.SubnetTestInfo, interfaces.SubnetTestInfo) {
	return n.subnetAInfo, n.subnetBInfo
}

func (n *localNetwork) GetTeleporterContractAddress() common.Address {
	return n.teleporterContractAddress
}

func (n *localNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return n.fundedAddress, n.fundedKey
}

func (n *localNetwork) RelayMessage(ctx context.Context,
	sourceReceipt *types.Receipt,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
	alterMessage bool,
	expectSuccess bool) *types.Receipt {
	return relayMessage(
		ctx,
		sourceReceipt,
		source,
		destination,
		n.teleporterContractAddress,
		n.fundedAddress,
		n.fundedKey,
		alterMessage,
		expectSuccess)
}

func (n *localNetwork) tearDownNetwork() {
	Expect(n.manager).ShouldNot(BeNil())
	Expect(n.manager.TeardownNetwork()).Should(BeNil())
	Expect(os.Remove(n.warpChainConfigPath)).Should(BeNil())
}

func (n *localNetwork) getCurrentSubnetValues(subnetInfo interfaces.SubnetTestInfo) interfaces.SubnetTestInfo {
	subnetDetails, ok := n.manager.GetSubnet(subnetInfo.SubnetID)
	Expect(ok).Should(BeTrue())
	blockchainID := subnetDetails.BlockchainID

	// Reset the validator URIs, as they may have changed
	subnetDetails.ValidatorURIs = nil
	status, err := n.anrClient.Status(context.Background())
	Expect(err).Should(BeNil())
	nodeInfos := status.GetClusterInfo().GetNodeInfos()

	for _, nodeName := range subnetInfo.NodeNames {
		subnetDetails.ValidatorURIs = append(subnetDetails.ValidatorURIs, nodeInfos[nodeName].Uri)
	}
	var nodeURIs []string
	nodeURIs = append(nodeURIs, subnetDetails.ValidatorURIs...)

	wsURI := utils.HttpToWebsocketURI(nodeURIs[0], blockchainID.String())
	rpcURI := utils.HttpToRPCURI(nodeURIs[0], blockchainID.String())

	if subnetInfo.WSClient != nil {
		subnetInfo.WSClient.Close()
	}
	wsClient, err := ethclient.Dial(wsURI)
	Expect(err).Should(BeNil())
	if subnetInfo.RPCClient != nil {
		subnetInfo.RPCClient.Close()
	}
	rpcClient, err := ethclient.Dial(rpcURI)
	Expect(err).Should(BeNil())
	evmChainID, err := rpcClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	return interfaces.SubnetTestInfo{
		SubnetID:                  subnetInfo.SubnetID,
		BlockchainID:              subnetInfo.BlockchainID,
		NodeURIs:                  nodeURIs,
		NodeNames:                 subnetInfo.NodeNames,
		WSClient:                  wsClient,
		RPCClient:                 rpcClient,
		EVMChainID:                evmChainID,
		TeleporterRegistryAddress: subnetInfo.TeleporterRegistryAddress,
	}
}

func (n *localNetwork) addSubnetAValidators(ctx context.Context, nodeNames []string) {
	_, err := n.anrClient.AddSubnetValidators(ctx, []*rpcpb.SubnetValidatorsSpec{
		{
			SubnetId:  n.subnetAInfo.SubnetID.String(),
			NodeNames: nodeNames,
		},
	})
	Expect(err).Should(BeNil())

	// Add the new node names
	n.subnetAInfo.NodeNames = append(n.subnetAInfo.NodeNames, nodeNames...)

	n.subnetAInfo = n.getCurrentSubnetValues(n.subnetAInfo)
}
