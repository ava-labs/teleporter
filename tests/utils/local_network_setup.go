package utils

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
	examplecrosschainmessenger "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/ExampleMessenger/ExampleCrossChainMessenger"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/Mocks/ExampleERC20"
	teleporterregistry "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/upgrades/TeleporterRegistry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

const (
	fundedKeyStr = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
)

var (
	teleporterContractAddress                              common.Address
	subnetA, subnetB                                       ids.ID
	blockchainIDA, blockchainIDB                           ids.ID
	chainANodeURIs, chainBNodeURIs                         []string
	fundedAddress                                          = common.HexToAddress("0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC")
	fundedKey                                              *ecdsa.PrivateKey
	chainAWSClient, chainBWSClient                         ethclient.Client
	chainARPCClient, chainBRPCClient                       ethclient.Client
	chainAIDInt, chainBIDInt                               *big.Int
	teleporterRegistryAddressA, teleporterRegistryAddressB common.Address

	// Internal vars only used to set up the local network
	anrClient           runner_sdk.Client
	manager             *runner.NetworkManager
	warpChainConfigPath string
)

//
// Global test state getters. Should be called within a test spec, after SetupNetwork has been called
//

type SubnetTestInfo struct {
	SubnetID                  ids.ID
	BlockchainID              ids.ID
	ChainNodeURIs             []string
	ChainWSClient             ethclient.Client
	ChainRPCClient            ethclient.Client
	ChainIDInt                *big.Int
	TeleporterRegistryAddress common.Address
}

func GetSubnetsInfo() []SubnetTestInfo {
	return []SubnetTestInfo{
		GetSubnetATestInfo(),
		GetSubnetBTestInfo(),
	}
}

func GetSubnetATestInfo() SubnetTestInfo {
	return SubnetTestInfo{
		SubnetID:                  subnetA,
		BlockchainID:              blockchainIDA,
		ChainNodeURIs:             chainANodeURIs,
		ChainWSClient:             chainAWSClient,
		ChainRPCClient:            chainARPCClient,
		ChainIDInt:                big.NewInt(0).Set(chainAIDInt),
		TeleporterRegistryAddress: teleporterRegistryAddressA,
	}
}
func GetSubnetBTestInfo() SubnetTestInfo {
	return SubnetTestInfo{
		SubnetID:                  subnetB,
		BlockchainID:              blockchainIDB,
		ChainNodeURIs:             chainBNodeURIs,
		ChainWSClient:             chainBWSClient,
		ChainRPCClient:            chainBRPCClient,
		ChainIDInt:                big.NewInt(0).Set(chainBIDInt),
		TeleporterRegistryAddress: teleporterRegistryAddressB,
	}
}
func GetTeleporterContractAddress() common.Address {
	return teleporterContractAddress
}
func GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	key, err := crypto.ToECDSA(crypto.FromECDSA(fundedKey))
	Expect(err).Should(BeNil())
	return fundedAddress, key
}

// SetupNetwork starts the default network and adds 10 new nodes as validators with BLS keys
// registered on the P-Chain.
// Adds two disjoint sets of 5 of the new validator nodes to validate two new subnets with a
// a single Subnet-EVM blockchain.
func SetupNetwork(warpGenesisFile string) {
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

	anrConfig := runner.NewDefaultANRConfig()
	log.Info("dbg", "anrConfig", anrConfig)
	manager = runner.NewNetworkManager(anrConfig)

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
					SubnetConfig: `{
						"proposerMinBlockDelay":0
					}`,
					Participants: subnetANodeNames,
				},
			},
			{
				VmName:      evm.IDStr,
				Genesis:     warpGenesisFile,
				ChainConfig: warpChainConfigPath,
				SubnetSpec: &rpcpb.SubnetSpec{
					SubnetConfig: `{
						"proposerMinBlockDelay":0
					}`,
					Participants: subnetBNodeNames,
				},
			},
		},
	)
	Expect(err).Should(BeNil())

	// Issue transactions to activate the proposerVM fork on the chains
	fundedKey, err = crypto.HexToECDSA(fundedKeyStr)
	Expect(err).Should(BeNil())
	SetupProposerVM(ctx, fundedKey, manager, 0)
	SetupProposerVM(ctx, fundedKey, manager, 1)

	// Create the ANR client
	logLevel, err := logging.ToLevel("info")
	Expect(err).Should(BeNil())

	logFactory := logging.NewFactory(logging.Config{
		DisplayLevel: logLevel,
		LogLevel:     logLevel,
	})
	zapLog, err := logFactory.Make("main")
	Expect(err).Should(BeNil())

	anrClient, err = runner_sdk.New(runner_sdk.Config{
		Endpoint:    "0.0.0.0:12352",
		DialTimeout: 10 * time.Second,
	}, zapLog)

	SetSubnetValues()

	log.Info("Finished setting up e2e test subnet variables")
}

func SetSubnetValues() {
	subnetIDs := manager.GetSubnets()
	Expect(len(subnetIDs)).Should(Equal(2))

	subnetA = subnetIDs[0]
	SetSubnetAValues(subnetA)

	subnetB = subnetIDs[1]
	SetSubnetBValues(subnetB)
}

func SetSubnetBValues(subnetID ids.ID) {
	var err error
	subnetBDetails, ok := manager.GetSubnet(subnetID)
	Expect(ok).Should(BeTrue())
	// Expect(len(subnetBDetails.ValidatorURIs)).Should(Equal(5))
	blockchainIDB = subnetBDetails.BlockchainID

	// remove
	log.Info("debug", "subnetBURIs before", subnetBDetails.ValidatorURIs)
	subnetBDetails.ValidatorURIs = nil
	status, err := anrClient.Status(context.Background())
	Expect(err).Should(BeNil())
	nodeInfos := status.GetClusterInfo().GetNodeInfos()

	var nodeNames []string
	for i := 6; i <= 10; i++ {
		n := fmt.Sprintf("node%d-bls", i)
		nodeNames = append(nodeNames, n)
	}

	anrClient.Status(context.Background())
	for _, nodeName := range nodeNames {
		subnetBDetails.ValidatorURIs = append(subnetBDetails.ValidatorURIs, nodeInfos[nodeName].Uri)
	}
	log.Info("debug", "subnetBURIs after", subnetBDetails.ValidatorURIs)

	// remove

	chainBNodeURIs = nil
	chainBNodeURIs = append(chainBNodeURIs, subnetBDetails.ValidatorURIs...)

	chainBWSURI := HttpToWebsocketURI(chainBNodeURIs[0], blockchainIDB.String())
	chainBRPCURI := HttpToRPCURI(chainBNodeURIs[0], blockchainIDB.String())

	if chainBWSClient != nil {
		chainBWSClient.Close()
	}
	chainBWSClient, err = ethclient.Dial(chainBWSURI)
	Expect(err).Should(BeNil())
	if chainBRPCClient != nil {
		chainBRPCClient.Close()
	}
	chainBRPCClient, err = ethclient.Dial(chainBRPCURI)
	Expect(err).Should(BeNil())
	chainBIDInt, err = chainBRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())
}

func SetSubnetAValues(subnetID ids.ID) {
	var err error
	subnetADetails, ok := manager.GetSubnet(subnetID)
	Expect(ok).Should(BeTrue())
	// Expect(len(subnetADetails.ValidatorURIs)).Should(Equal(5))
	blockchainIDA = subnetADetails.BlockchainID

	// remove
	log.Info("debug", "subnetAURIs before", subnetADetails.ValidatorURIs)
	subnetADetails.ValidatorURIs = nil
	status, err := anrClient.Status(context.Background())
	Expect(err).Should(BeNil())
	nodeInfos := status.GetClusterInfo().GetNodeInfos()

	var nodeNames []string
	for i := 1; i <= 5; i++ {
		n := fmt.Sprintf("node%d-bls", i)
		nodeNames = append(nodeNames, n)
	}

	anrClient.Status(context.Background())
	for _, nodeName := range nodeNames {
		subnetADetails.ValidatorURIs = append(subnetADetails.ValidatorURIs, nodeInfos[nodeName].Uri)
	}
	log.Info("debug", "subnetAURIs after", subnetADetails.ValidatorURIs)

	// remove

	chainANodeURIs = nil
	chainANodeURIs = append(chainANodeURIs, subnetADetails.ValidatorURIs...)

	chainAWSURI := HttpToWebsocketURI(chainANodeURIs[0], blockchainIDA.String())
	chainARPCURI := HttpToRPCURI(chainANodeURIs[0], blockchainIDA.String())

	if chainAWSClient != nil {
		chainAWSClient.Close()
	}
	chainAWSClient, err = ethclient.Dial(chainAWSURI)
	Expect(err).Should(BeNil())
	if chainARPCClient != nil {
		chainARPCClient.Close()
	}
	chainARPCClient, err = ethclient.Dial(chainARPCURI)
	Expect(err).Should(BeNil())
	chainAIDInt, err = chainARPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())
}

// DeployTeleporterContracts deploys the Teleporter contract to the two subnets. The caller is responsible for generating the
// deployment transaction information
func DeployTeleporterContracts(transactionBytes []byte, deployerAddress common.Address, contractAddress common.Address) {
	log.Info("Deploying Teleporter contract to subnets")

	subnetsInfo := GetSubnetsInfo()

	// Set the package level teleporterContractAddress
	teleporterContractAddress = contractAddress

	ctx := context.Background()

	for _, subnetInfo := range subnetsInfo {
		// Fund the deployer address
		{
			fundAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)) // 10eth
			fundDeployerTx := CreateNativeTransferTransaction(ctx, subnetInfo, fundedAddress, fundedKey, deployerAddress, fundAmount)
			SendTransactionAndWaitForAcceptance(ctx, subnetInfo.ChainWSClient, subnetInfo.ChainRPCClient, fundDeployerTx, true)
		}
		log.Info("Finished funding Teleporter deployer", "blockchainID", subnetInfo.BlockchainID.Hex())

		// Deploy Teleporter contract
		{
			rpcClient, err := rpc.DialContext(ctx, HttpToRPCURI(subnetInfo.ChainNodeURIs[0], subnetInfo.BlockchainID.String()))
			Expect(err).Should(BeNil())
			defer rpcClient.Close()

			newHeads := make(chan *types.Header, 10)
			subA, err := subnetInfo.ChainWSClient.SubscribeNewHead(ctx, newHeads)
			Expect(err).Should(BeNil())
			defer subA.Unsubscribe()

			err = rpcClient.CallContext(ctx, nil, "eth_sendRawTransaction", hexutil.Encode(transactionBytes))
			Expect(err).Should(BeNil())

			<-newHeads
			teleporterCode, err := subnetInfo.ChainRPCClient.CodeAt(ctx, teleporterContractAddress, nil)
			Expect(err).Should(BeNil())
			Expect(len(teleporterCode)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
		}
		log.Info("Finished deploying Teleporter contract", "blockchainID", subnetInfo.BlockchainID.Hex())
	}
	log.Info("Deployed Teleporter contracts to all subnets")
}

func DeployTeleporterRegistryContracts(teleporterAddress common.Address) {
	log.Info("Deploying TeleporterRegistry contract to subnets")
	ctx := context.Background()

	entries := []teleporterregistry.ProtocolRegistryEntry{
		{
			Version:         big.NewInt(1),
			ProtocolAddress: teleporterAddress,
		},
	}

	subnetAInfo := GetSubnetATestInfo()
	subnetBInfo := GetSubnetBTestInfo()

	var (
		err error
		tx  *types.Transaction
	)
	optsA := CreateTransactorOpts(ctx, subnetAInfo, fundedAddress, fundedKey)
	teleporterRegistryAddressA, tx, _, err = teleporterregistry.DeployTeleporterRegistry(optsA, subnetAInfo.ChainRPCClient, entries)
	Expect(err).Should(BeNil())
	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, subnetAInfo.ChainRPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	optsB := CreateTransactorOpts(ctx, subnetBInfo, fundedAddress, fundedKey)
	teleporterRegistryAddressB, tx, _, err = teleporterregistry.DeployTeleporterRegistry(optsB, subnetBInfo.ChainRPCClient, entries)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err = bind.WaitMined(ctx, subnetBInfo.ChainRPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	log.Info("Deployed TeleporterRegistry contracts to all subnets")
}

func DeployMockToken(ctx context.Context, fundedAddress common.Address, fundedKey *ecdsa.PrivateKey, source SubnetTestInfo) (common.Address, *exampleerc20.ExampleERC20) {
	opts := CreateTransactorOpts(ctx, source, fundedAddress, fundedKey)

	// Deploy Mock ERC20 contract
	address, txn, mockToken, err := exampleerc20.DeployExampleERC20(opts, source.ChainRPCClient)
	Expect(err).Should(BeNil())
	log.Info("Deployed Mock ERC20 contract", "address", address.Hex(), "txHash", txn.Hash().Hex())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, source.ChainRPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	return address, mockToken
}

func DeployExampleCrossChainMessenger(ctx context.Context, fundedAddress common.Address, fundedKey *ecdsa.PrivateKey, source SubnetTestInfo) (common.Address, *examplecrosschainmessenger.ExampleCrossChainMessenger) {
	optsA := CreateTransactorOpts(ctx, source, fundedAddress, fundedKey)
	address, tx, exampleMessenger, err := examplecrosschainmessenger.DeployExampleCrossChainMessenger(optsA, source.ChainRPCClient, source.TeleporterRegistryAddress)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, source.ChainRPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	return address, exampleMessenger
}

func ExampleERC20Approve(ctx context.Context, mockToken *exampleerc20.ExampleERC20, spender common.Address, amount *big.Int, source SubnetTestInfo) {
	opts := CreateTransactorOpts(ctx, source, fundedAddress, fundedKey)
	txn, err := mockToken.Approve(opts, spender, amount)
	Expect(err).Should(BeNil())
	log.Info("Approved Mock ERC20", "spender", teleporterContractAddress.Hex(), "txHash", txn.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, source.ChainRPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
}

func TearDownNetwork() {
	log.Info("Tearing down network")
	Expect(manager).ShouldNot(BeNil())
	Expect(manager.TeardownNetwork()).Should(BeNil())
	Expect(os.Remove(warpChainConfigPath)).Should(BeNil())
}

func RemoveSubnetValidators(ctx context.Context, subnetID ids.ID, nodeNames []string) {
	_, err := anrClient.RemoveSubnetValidator(ctx, []*rpcpb.RemoveSubnetValidatorSpec{
		{
			SubnetId:  subnetID.String(),
			NodeNames: nodeNames,
		},
	})
	Expect(err).Should(BeNil())
	SetSubnetValues()
}

func AddSubnetValidators(ctx context.Context, subnetID ids.ID, nodeNames []string) {
	log.Info("debug",
		"subnetID", subnetID.String(),
		"nodeNames", nodeNames,
	)
	_, err := anrClient.AddSubnetValidators(ctx, []*rpcpb.SubnetValidatorsSpec{
		{
			SubnetId:  subnetID.String(),
			NodeNames: nodeNames,
		},
	})
	Expect(err).Should(BeNil())
	SetSubnetValues()
}

func RestartNodes(ctx context.Context, nodeNames []string) {
	for _, nodeName := range nodeNames {
		log.Info("debug", "restarting node", "nodeName", nodeName)
		_, err := anrClient.RestartNode(ctx, nodeName)
		Expect(err).Should(BeNil())
	}
	time.Sleep(30 * time.Second)
	SetSubnetValues()
}
