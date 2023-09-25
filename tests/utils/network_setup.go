package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ava-labs/avalanche-network-runner/rpcpb"
	"github.com/ava-labs/avalanchego/ids"
	relayerEvm "github.com/ava-labs/awm-relayer/vms/evm"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/plugin/evm"
	"github.com/ava-labs/subnet-evm/rpc"
	"github.com/ava-labs/subnet-evm/tests/utils/runner"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

const (
	FundedKeyStr = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"
)

var (
	TeleporterContractAddress        common.Address
	SubnetIDs                        []ids.ID
	SubnetA, SubnetB                 ids.ID
	BlockchainIDA, BlockchainIDB     ids.ID
	ChainANodeURIs, ChainBNodeURIs   []string
	FundedAddress                    = common.HexToAddress("0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC")
	FundedKey                        *ecdsa.PrivateKey
	ChainAWSClient, ChainBWSClient   ethclient.Client
	ChainARPCClient, ChainBRPCClient ethclient.Client
	ChainARPCURI, ChainBRPCURI       string
	ChainAIDInt, ChainBIDInt         *big.Int
	NewHeadsA                        chan *types.Header

	// Internal vars only used to set up the local network
	anrConfig           = runner.NewDefaultANRConfig()
	manager             = runner.NewNetworkManager(anrConfig)
	warpChainConfigPath string
)

// setupNetwork starts the default network and adds 10 new nodes as validators with BLS keys
// registered on the P-Chain.
// Adds two disjoint sets of 5 of the new validator nodes to validate two new subnets with a
// a single Subnet-EVM blockchain.
func SetupNetwork(warpGenesisFile string) {
	ctx := context.Background()
	var err error

	// Name 10 new validators (which should have BLS key registered)
	SubnetANodeNames := []string{}
	SubnetBNodeNames := []string{}
	for i := 1; i <= 10; i++ {
		n := fmt.Sprintf("node%d-bls", i)
		if i <= 5 {
			SubnetANodeNames = append(SubnetANodeNames, n)
		} else {
			SubnetBNodeNames = append(SubnetBNodeNames, n)
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
					Participants: SubnetANodeNames,
				},
			},
			{
				VmName:      evm.IDStr,
				Genesis:     warpGenesisFile,
				ChainConfig: warpChainConfigPath,
				SubnetSpec: &rpcpb.SubnetSpec{
					SubnetConfig: "",
					Participants: SubnetBNodeNames,
				},
			},
		},
	)
	Expect(err).Should(BeNil())

	// Issue transactions to activate the proposerVM fork on the chains
	FundedKey, err = crypto.HexToECDSA(FundedKeyStr)
	Expect(err).Should(BeNil())
	SetUpProposerVm(ctx, FundedKey, manager, 0)
	SetUpProposerVm(ctx, FundedKey, manager, 1)

	// Set up subnet URIs
	SubnetIDs = manager.GetSubnets()
	Expect(len(SubnetIDs)).Should(Equal(2))

	SubnetA = SubnetIDs[0]
	SubnetADetails, ok := manager.GetSubnet(SubnetA)
	Expect(ok).Should(BeTrue())
	Expect(len(SubnetADetails.ValidatorURIs)).Should(Equal(5))
	BlockchainIDA = SubnetADetails.BlockchainID
	ChainANodeURIs = append(ChainANodeURIs, SubnetADetails.ValidatorURIs...)

	SubnetB = SubnetIDs[1]
	SubnetBDetails, ok := manager.GetSubnet(SubnetB)
	Expect(ok).Should(BeTrue())
	Expect(len(SubnetBDetails.ValidatorURIs)).Should(Equal(5))
	BlockchainIDB = SubnetBDetails.BlockchainID
	ChainBNodeURIs = append(ChainBNodeURIs, SubnetBDetails.ValidatorURIs...)

	log.Info(
		"Created URIs for subnets",
		"ChainAURIs", ChainANodeURIs,
		"ChainBURIs", ChainBNodeURIs,
		"BlockchainIDA", BlockchainIDA,
		"BlockchainIDB", BlockchainIDB,
	)

	chainAWSURI := HttpToWebsocketURI(ChainANodeURIs[0], BlockchainIDA.String())
	ChainARPCURI = HttpToRPCURI(ChainANodeURIs[0], BlockchainIDA.String())
	log.Info("Creating ethclient for blockchainA", "wsURI", chainAWSURI, "rpcURL, ChainARPCURI")
	ChainAWSClient, err = ethclient.Dial(chainAWSURI)
	Expect(err).Should(BeNil())
	ChainARPCClient, err = ethclient.Dial(ChainARPCURI)
	Expect(err).Should(BeNil())

	ChainAIDInt, err = ChainARPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	chainBWSURI := HttpToWebsocketURI(ChainBNodeURIs[0], BlockchainIDB.String())
	ChainBRPCURI = HttpToRPCURI(ChainBNodeURIs[0], BlockchainIDB.String())
	log.Info("Creating ethclient for blockchainB", "wsURI", chainBWSURI)
	ChainBWSClient, err = ethclient.Dial(chainBWSURI)
	Expect(err).Should(BeNil())
	ChainBRPCClient, err = ethclient.Dial(ChainBRPCURI)
	Expect(err).Should(BeNil())

	ChainBIDInt, err = ChainBRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	NewHeadsA = make(chan *types.Header, 10)

	log.Info("Finished setting up e2e test subnet variables")

	log.Info("Deploying Teleporter contract to subnets")

	log.Info("Set up ginkgo before suite")
}

// DeployTeleporterContract deploys the Teleporter contract to the two subnets. The caller is responsible for generating the
// deployment transaction information
func DeployTeleporterContract(transactionBytes []byte, deployerAddress common.Address, contractAddress common.Address) {
	// Set the package level TeleporterContractAddress
	TeleporterContractAddress = contractAddress

	ctx := context.Background()
	nonceA, err := ChainARPCClient.NonceAt(ctx, FundedAddress, nil)
	Expect(err).Should(BeNil())

	nonceB, err := ChainBRPCClient.NonceAt(ctx, FundedAddress, nil)
	Expect(err).Should(BeNil())

	gasTipCapA, err := ChainARPCClient.SuggestGasTipCap(context.Background())
	Expect(err).Should(BeNil())
	gasTipCapB, err := ChainBRPCClient.SuggestGasTipCap(context.Background())
	Expect(err).Should(BeNil())

	baseFeeA, err := ChainARPCClient.EstimateBaseFee(context.Background())
	Expect(err).Should(BeNil())
	gasFeeCapA := baseFeeA.Mul(baseFeeA, big.NewInt(relayerEvm.BaseFeeFactor))
	gasFeeCapA.Add(gasFeeCapA, big.NewInt(relayerEvm.MaxPriorityFeePerGas))

	baseFeeB, err := ChainBRPCClient.EstimateBaseFee(context.Background())
	Expect(err).Should(BeNil())
	gasFeeCapB := baseFeeB.Mul(baseFeeB, big.NewInt(relayerEvm.BaseFeeFactor))
	gasFeeCapB.Add(gasFeeCapB, big.NewInt(relayerEvm.MaxPriorityFeePerGas))

	// Fund the deployer address
	{
		value := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)) // 10eth
		txA := types.NewTx(&types.DynamicFeeTx{
			ChainID:   ChainAIDInt,
			Nonce:     nonceA,
			To:        &deployerAddress,
			Gas:       DefaultTeleporterTransactionGas,
			GasFeeCap: gasFeeCapA,
			GasTipCap: gasTipCapA,
			Value:     value,
		})
		txSignerA := types.LatestSignerForChainID(ChainAIDInt)
		triggerTxA, err := types.SignTx(txA, txSignerA, FundedKey)
		Expect(err).Should(BeNil())
		err = ChainARPCClient.SendTransaction(ctx, triggerTxA)
		Expect(err).Should(BeNil())
		time.Sleep(5 * time.Second)
		receipt, err := ChainARPCClient.TransactionReceipt(ctx, triggerTxA.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	}
	{
		value := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)) // 10eth
		txB := types.NewTx(&types.DynamicFeeTx{
			ChainID:   ChainBIDInt,
			Nonce:     nonceB,
			To:        &deployerAddress,
			Gas:       DefaultTeleporterTransactionGas,
			GasFeeCap: gasFeeCapB,
			GasTipCap: gasTipCapB,
			Value:     value,
		})
		txSignerB := types.LatestSignerForChainID(ChainBIDInt)
		triggerTxB, err := types.SignTx(txB, txSignerB, FundedKey)
		Expect(err).Should(BeNil())
		err = ChainBRPCClient.SendTransaction(ctx, triggerTxB)
		Expect(err).Should(BeNil())
		time.Sleep(5 * time.Second)
		receipt, err := ChainBRPCClient.TransactionReceipt(ctx, triggerTxB.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	}
	// Deploy Teleporter on the two subnets
	{
		rpcClient, err := rpc.DialContext(ctx, ChainARPCURI)
		Expect(err).Should(BeNil())
		err = rpcClient.CallContext(ctx, nil, "eth_sendRawTransaction", hexutil.Encode(transactionBytes))
		Expect(err).Should(BeNil())
		time.Sleep(5 * time.Second)
		teleporterCode, err := ChainARPCClient.CodeAt(ctx, TeleporterContractAddress, nil)
		Expect(err).Should(BeNil())
		Expect(len(teleporterCode)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
	}
	{
		rpcClient, err := rpc.DialContext(ctx, ChainBRPCURI)
		Expect(err).Should(BeNil())
		err = rpcClient.CallContext(ctx, nil, "eth_sendRawTransaction", hexutil.Encode(transactionBytes))
		Expect(err).Should(BeNil())
		time.Sleep(5 * time.Second)
		teleporterCode, err := ChainBRPCClient.CodeAt(ctx, TeleporterContractAddress, nil)
		Expect(err).Should(BeNil())
		Expect(len(teleporterCode)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
	}
	log.Info("Finished deploying Teleporter contracts")
}

func TearDownNetwork() {
	log.Info("Running ginkgo after suite")
	Expect(manager).ShouldNot(BeNil())
	Expect(manager.TeardownNetwork()).Should(BeNil())
	Expect(os.Remove(warpChainConfigPath)).Should(BeNil())
}
