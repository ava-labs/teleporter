// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/awm-relayer/messages/teleporter"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/params"
	"github.com/ava-labs/subnet-evm/tests/utils"
	"github.com/ava-labs/subnet-evm/tests/utils/runner"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

var (
	defaultTeleporterTransactionGas       uint64 = 200_000
	defaultTeleporterTransactionGasFeeCap        = big.NewInt(225 * params.GWei)
	defaultTeleporterTransactionGasTipCap        = big.NewInt(params.GWei)
	defaultTeleporterTransactionValue            = common.Big0
)

type SendCrossChainMessageEvent struct {
	DestinationChainID ids.ID
	MessageID          *big.Int
	Message            teleporter.TeleporterMessage
}

// Teleporter contract sendCrossChainMessage input type
type TeleporterMessageInput struct {
	DestinationChainID      ids.ID
	DestinationAddress      common.Address
	FeeInfo                 FeeInfo
	RequiredGasLimit        *big.Int
	Message                 []byte
	AllowedRelayerAddresses []common.Address
}

type FeeInfo struct {
	ContractAddress common.Address
	Amount          *big.Int
}

func httpToWebsocketURI(uri string, blockchainID string) string {
	return fmt.Sprintf("ws://%s/ext/bc/%s/ws", strings.TrimPrefix(uri, "http://"), blockchainID)
}

func httpToRPCURI(uri string, blockchainID string) string {
	return fmt.Sprintf("http://%s/ext/bc/%s/rpc", strings.TrimPrefix(uri, "http://"), blockchainID)
}

func getURIHostAndPort(uri string) (string, uint32, error) {
	// At a minimum uri should have http:// of 7 characters
	Expect(len(uri)).Should(BeNumerically(">", 7))
	if uri[:7] == "http://" {
		uri = uri[7:]
	} else if uri[:8] == "https://" {
		uri = uri[8:]
	} else {
		return "", 0, fmt.Errorf("invalid uri: %s", uri)
	}

	// Split the uri into host and port
	hostAndPort := strings.Split(uri, ":")
	Expect(len(hostAndPort)).Should(Equal(2))

	// Parse the port
	port, err := strconv.ParseUint(hostAndPort[1], 10, 32)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse port: %w", err)
	}

	return hostAndPort[0], uint32(port), nil
}

func newTestTeleporterTransaction(chainIDInt *big.Int, teleporterAddress common.Address, nonce uint64, data []byte) *types.Transaction {
	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainIDInt,
		Nonce:     nonce,
		To:        &teleporterAddress,
		Gas:       defaultTeleporterTransactionGas,
		GasFeeCap: defaultTeleporterTransactionGasFeeCap,
		GasTipCap: defaultTeleporterTransactionGasTipCap,
		Value:     defaultTeleporterTransactionValue,
		Data:      data,
	})
}

func setUpProposerVm(ctx context.Context, fundedKey *ecdsa.PrivateKey, manager *runner.NetworkManager, index int) {
	subnet := manager.GetSubnets()[index]
	subnetDetails, ok := manager.GetSubnet(subnet)
	Expect(ok).Should(BeTrue())

	chainID := subnetDetails.BlockchainID
	uri := httpToWebsocketURI(subnetDetails.ValidatorURIs[0], chainID.String())

	client, err := ethclient.Dial(uri)
	Expect(err).Should(BeNil())
	chainIDInt, err := client.ChainID(ctx)
	Expect(err).Should(BeNil())

	err = utils.IssueTxsToActivateProposerVMFork(ctx, chainIDInt, fundedKey, client)
	Expect(err).Should(BeNil())
}
