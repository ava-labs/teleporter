package testnet

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	"github.com/ava-labs/subnet-evm/rpc"
	warpBackend "github.com/ava-labs/subnet-evm/warp"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/teleporter/TeleporterMessenger"
	teleporterregistry "github.com/ava-labs/teleporter/abi-bindings/go/teleporter/registry/TeleporterRegistry"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	. "github.com/onsi/gomega"
)

const (
	subnetAPrefix                   = "subnet_a"
	subnetBPrefix                   = "subnet_b"
	subnetCPrefix                   = "subnet_c"
	cChainPrefix                    = "c_chain"
	teleporterContractAddress       = "teleporter_contract_address"
	teleporterRegistryAddressSuffix = "_teleporter_registry_address"
	subnetIDSuffix                  = "_subnet_id"
	blockchainIDSuffix              = "_blockchain_id"
	rpcURLSuffix                    = "_rpc_url"
	wsURLSuffix                     = "_ws_url"
	userAddress                     = "user_address"
	userPrivateKey                  = "user_private_key"

	receiveCrossChainMessageEventName      = "ReceiveCrossChainMessage"
	receiveCrossChainMessageLookBackBlocks = 500

	privateKeyHexLength = 64
)

var errInvalidPrivateKeyString = errors.New("invalid private key string")

var _ interfaces.Network = &testNetwork{}

type testNetwork struct {
	teleporterContractAddress common.Address
	primaryNetwork            *interfaces.SubnetTestInfo
	subnetsInfo               map[ids.ID]*interfaces.SubnetTestInfo
	fundedAddress             common.Address
	fundedKey                 *ecdsa.PrivateKey
}

func initializeSubnetInfo(
	subnetPrefix string,
	teleporterContractAddress common.Address,
) (interfaces.SubnetTestInfo, error) {
	subnetIDStr := os.Getenv(subnetPrefix + subnetIDSuffix)
	subnetID, err := ids.FromString(subnetIDStr)
	if err != nil {
		log.Info("Error decoding subnet ID", "subnetPrefix", subnetPrefix)
		return interfaces.SubnetTestInfo{}, err
	}

	blockchainIDStr := os.Getenv(subnetPrefix + blockchainIDSuffix)
	blockchainID, err := ids.FromString(blockchainIDStr)
	if err != nil {
		log.Info("Error decoding blockchain ID", "subnetPrefix", subnetPrefix)
		return interfaces.SubnetTestInfo{}, err
	}

	rpcURLStr := os.Getenv(subnetPrefix + rpcURLSuffix)

	// Create the client using a cookiejar to try to use the same node for each
	// request when using public RPC endpoints. Having requests routed to different
	// nodes behind a load balancer may cause issues with nodes serving slightly stale
	// data from before they see recently accepted transactions. Responses generally
	// have cookies identifying which node served the request, and those cookies can
	// be added to the cookiejar to be included on future requests to attempt to have
	// the same node serve it.
	// See here: https://docs.avax.network/tooling/rpc-providers#sticky-sessions
	jar, err := cookiejar.New(nil)
	if err != nil {
		return interfaces.SubnetTestInfo{}, err
	}

	httpClient := &http.Client{
		Jar: jar,
	}

	rpcClient, err := rpc.DialOptions(
		context.Background(),
		rpcURLStr,
		rpc.WithHTTPClient(httpClient),
	)
	if err != nil {
		return interfaces.SubnetTestInfo{}, err
	}

	ethClient := ethclient.NewClient(rpcClient)

	evmChainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		return interfaces.SubnetTestInfo{}, err
	}

	teleporterRegistryAddressStr := os.Getenv(subnetPrefix + teleporterRegistryAddressSuffix)
	if !common.IsHexAddress(teleporterRegistryAddressStr) {
		return interfaces.SubnetTestInfo{}, errors.New("invalid teleporter regirstry address")
	}
	teleporterRegistryAddress := common.HexToAddress(teleporterRegistryAddressStr)

	teleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress, ethClient,
	)
	if err != nil {
		return interfaces.SubnetTestInfo{}, err
	}

	teleporterRegistry, err := teleporterregistry.NewTeleporterRegistry(teleporterRegistryAddress, ethClient)
	if err != nil {
		return interfaces.SubnetTestInfo{}, err
	}

	return interfaces.SubnetTestInfo{
		SubnetID:                  subnetID,
		BlockchainID:              blockchainID,
		NodeURIs:                  []string{}, // no specific node URIs for a testnet subnet, only RPC endpoints.
		RPCClient:                 ethClient,
		WSClient:                  nil,
		EVMChainID:                evmChainID,
		TeleporterRegistryAddress: teleporterRegistryAddress,
		TeleporterRegistry:        teleporterRegistry,
		TeleporterMessenger:       teleporterMessenger,
	}, nil
}

func NewTestNetwork() (*testNetwork, error) {
	teleporterContractAddressStr := os.Getenv(teleporterContractAddress)
	teleporterContractAddress := common.HexToAddress(teleporterContractAddressStr)
	log.Info("Set teleporter contract address", "teleporterContractAddress", teleporterContractAddressStr)

	subnetAInfo, err := initializeSubnetInfo(subnetAPrefix, teleporterContractAddress)
	if err != nil {
		return nil, err
	}

	subnetBInfo, err := initializeSubnetInfo(subnetBPrefix, teleporterContractAddress)
	if err != nil {
		return nil, err
	}

	cChainInfo, err := initializeSubnetInfo(cChainPrefix, teleporterContractAddress)
	if err != nil {
		return nil, err
	}
	log.Info("Set testnet subnet info", subnetAPrefix, subnetAInfo, subnetBPrefix, subnetBInfo)

	fundedAddressStr := os.Getenv(userAddress)
	fundedKeyStr := os.Getenv(userPrivateKey)
	if len(fundedKeyStr) >= 2 && fundedKeyStr[0:2] == "0x" {
		fundedKeyStr = fundedKeyStr[2:]
	}
	if len(fundedKeyStr) != privateKeyHexLength {
		return nil, errInvalidPrivateKeyString
	}
	fundedKey, err := crypto.HexToECDSA(fundedKeyStr)
	if err != nil {
		return nil, err
	}
	log.Info("Set user funded address", "address", fundedAddressStr)

	subnetsInfo := make(map[ids.ID]*interfaces.SubnetTestInfo)
	subnetsInfo[subnetAInfo.SubnetID] = &subnetAInfo
	subnetsInfo[subnetBInfo.SubnetID] = &subnetBInfo
	return &testNetwork{
		teleporterContractAddress: teleporterContractAddress,
		primaryNetwork:            &cChainInfo,
		subnetsInfo:               subnetsInfo,
		fundedAddress:             common.HexToAddress(fundedAddressStr),
		fundedKey:                 fundedKey,
	}, nil
}

func (n *testNetwork) GetPrimaryNetworkInfo() interfaces.SubnetTestInfo {
	return *n.primaryNetwork
}

func (n *testNetwork) GetSubnetsInfo() []interfaces.SubnetTestInfo {
	subnetsInfo := make([]interfaces.SubnetTestInfo, 0, len(n.subnetsInfo))
	for _, subnetInfo := range n.subnetsInfo {
		subnetsInfo = append(subnetsInfo, *subnetInfo)
	}
	return subnetsInfo
}

// Returns subnet info for all subnets, including the primary network
func (n *testNetwork) GetAllSubnetsInfo() []interfaces.SubnetTestInfo {
	subnets := n.GetSubnetsInfo()
	return append(subnets, n.GetPrimaryNetworkInfo())
}

func (n *testNetwork) GetTeleporterContractAddress() common.Address {
	return n.teleporterContractAddress
}

func (n *testNetwork) SetTeleporterContractAddress(newTeleporterAddress common.Address) {
	n.teleporterContractAddress = newTeleporterAddress
	subnets := n.GetAllSubnetsInfo()
	for _, subnetInfo := range subnets {
		teleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
			n.teleporterContractAddress, subnetInfo.RPCClient,
		)
		Expect(err).Should(BeNil())
		if subnetInfo.SubnetID == constants.PrimaryNetworkID {
			n.primaryNetwork.TeleporterMessenger = teleporterMessenger
		} else {
			n.subnetsInfo[subnetInfo.SubnetID].TeleporterMessenger = teleporterMessenger
		}
	}
}

func (n *testNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return n.fundedAddress, n.fundedKey
}

func (n *testNetwork) IsExternalNetwork() bool {
	return true
}

func (n *testNetwork) SupportsIndependentRelaying() bool {
	// The test application cannot relay its own messages on testnets
	// because it can't query validators directly for their BLS signatures.
	return false
}

// For testnet messages, rely on a separately deployed relayer to relay the message.
// The implementation checks for the deliver of the given message on the destination
// within a time window of {relayWaitTime} seconds, and returns the receipt of the
// transaction that delivered the message.
func (n *testNetwork) RelayMessage(
	ctx context.Context,
	sourceReceipt *types.Receipt,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
	expectSuccess bool,
) *types.Receipt {
	// Get the Teleporter message ID from the receipt
	sendEvent, err := utils.GetEventFromLogs(
		sourceReceipt.Logs, source.TeleporterMessenger.ParseSendCrossChainMessage,
	)
	Expect(err).Should(BeNil())

	teleporterMessageID := sendEvent.MessageID

	receipt, err := n.getMessageDeliveryTransactionReceipt(ctx, source.BlockchainID, destination, teleporterMessageID)
	Expect(err).Should(BeNil())
	Expect(receipt).ShouldNot(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	return receipt
}

func (n *testNetwork) getMessageDeliveryTransactionReceipt(
	ctx context.Context,
	sourceBlockchainID ids.ID,
	destination interfaces.SubnetTestInfo,
	teleporterMessageID ids.ID,
) (*types.Receipt, error) {
	// Wait until the message is delivered.
	queryTicker := time.NewTicker(time.Millisecond * 500)
	defer queryTicker.Stop()

	// Wait a maximum of 20 seconds.
	cctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	for {
		delivered, err := destination.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, teleporterMessageID)
		if err != nil {
			return nil, err
		}
		if delivered {
			break
		}

		// Wait for the next round.
		select {
		case <-cctx.Done():
			return nil, cctx.Err()
		case <-queryTicker.C:
		}
	}

	// Get the latest block height
	currentBlockHeight, err := destination.RPCClient.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	var startBlock uint64
	if currentBlockHeight > receiveCrossChainMessageLookBackBlocks {
		startBlock = currentBlockHeight - receiveCrossChainMessageLookBackBlocks
	} else {
		startBlock = 0
	}

	// Get the log event of the delivery. The log must be in the last {receiveCrossChainMessageLookBackBlocks} blocks.
	logIter, err := destination.TeleporterMessenger.FilterReceiveCrossChainMessage(
		&bind.FilterOpts{
			Start:   startBlock,
			Context: cctx,
		},
		[][32]byte{teleporterMessageID},
		[][32]byte{sourceBlockchainID},
		nil,
	)
	if err != nil {
		return nil, err
	}

	// There should be exactly one matching event.
	if !logIter.Next() {
		return nil, errors.New("failed to find ReceiveCrossChainMessage log for relayed message")
	}
	receiveEvent := logIter.Event
	if logIter.Next() {
		return nil, errors.New("found multiple ReceiveCrossChainMessage logs for relayed message")
	}

	// The transaction should already be mined, but WaitMined will also wait for the eth_blockNumber
	// endpoint to reflect the block that the transaction has been included in.
	return utils.WaitMined(ctx, destination.RPCClient, receiveEvent.Raw.TxHash)
}

func (n *testNetwork) GetSignedMessage(
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
