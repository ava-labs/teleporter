package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ava-labs/avalanche-network-runner/rpcpb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/plugin/evm"
	"github.com/ava-labs/subnet-evm/rpc"
	"github.com/ava-labs/subnet-evm/tests/utils/runner"
	teleporterregistry "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/upgrades/TeleporterRegistry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

const (
	fundedKeyStr                   = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
	TeleporterRegistryByteCodeFile = "./contracts/out/TeleporterRegistry.sol/TeleporterRegistry.json"
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
	anrConfig           = runner.NewDefaultANRConfig()
	manager             = runner.NewNetworkManager(anrConfig)
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
	fundedKey, err = crypto.HexToECDSA(fundedKeyStr)
	Expect(err).Should(BeNil())
	SetupProposerVM(ctx, fundedKey, manager, 0)
	SetupProposerVM(ctx, fundedKey, manager, 1)

	// Set up subnet URIs
	subnetIDs := manager.GetSubnets()
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
		"chainAURIs", chainANodeURIs,
		"chainBURIs", chainBNodeURIs,
		"blockchainIDA", blockchainIDA,
		"blockchainIDB", blockchainIDB,
	)

	chainAWSURI := HttpToWebsocketURI(chainANodeURIs[0], blockchainIDA.String())
	chainARPCURI := HttpToRPCURI(chainANodeURIs[0], blockchainIDA.String())

	log.Info("Creating ethclients for blockchainA", "wsURI", chainAWSURI, "rpcURI", chainARPCURI)
	chainAWSClient, err = ethclient.Dial(chainAWSURI)
	Expect(err).Should(BeNil())
	chainARPCClient, err = ethclient.Dial(chainARPCURI)
	Expect(err).Should(BeNil())
	chainAIDInt, err = chainARPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	chainBWSURI := HttpToWebsocketURI(chainBNodeURIs[0], blockchainIDB.String())
	chainBRPCURI := HttpToRPCURI(chainBNodeURIs[0], blockchainIDB.String())
	log.Info("Creating ethclients for blockchainB", "wsURI", chainBWSURI, "rpcURI", chainBRPCURI)

	chainBWSClient, err = ethclient.Dial(chainBWSURI)
	Expect(err).Should(BeNil())
	chainBRPCClient, err = ethclient.Dial(chainBRPCURI)
	Expect(err).Should(BeNil())
	chainBIDInt, err = chainBRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	log.Info("Finished setting up e2e test subnet variables")
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

	teleporterRegistryABI, err := teleporterregistry.TeleporterRegistryMetaData.GetAbi()
	Expect(err).Should(BeNil())

	teleporterRegistryAddressA = DeployContract(ctx, TeleporterRegistryByteCodeFile, fundedKey, GetSubnetATestInfo(), teleporterRegistryABI, entries)
	teleporterRegistryAddressB = DeployContract(ctx, TeleporterRegistryByteCodeFile, fundedKey, GetSubnetBTestInfo(), teleporterRegistryABI, entries)

	log.Info("Deployed TeleporterRegistry contracts to all subnets")
}

func TearDownNetwork() {
	log.Info("Tearing down network")
	Expect(manager).ShouldNot(BeNil())
	Expect(manager.TeardownNetwork()).Should(BeNil())
	Expect(os.Remove(warpChainConfigPath)).Should(BeNil())
}
