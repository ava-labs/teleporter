// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/awm-relayer/messages/teleporter"
	"github.com/ava-labs/awm-relayer/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/params"
	"github.com/ethereum/go-ethereum/common"
	"github.com/onsi/gomega"
)

var (
	defaultTeleporterMessageGas       uint64 = 200_000
	defaultTeleporterMessageGasFeeCap        = big.NewInt(225 * params.GWei)
	defaultTeleporterMessageGasTipCap        = big.NewInt(params.GWei)
	defaultTeleporterMessageValue            = common.Big0
)

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
	gomega.Expect(len(uri)).Should(gomega.BeNumerically(">", 7))
	if uri[:7] == "http://" {
		uri = uri[7:]
	} else if uri[:8] == "https://" {
		uri = uri[8:]
	} else {
		return "", 0, fmt.Errorf("invalid uri: %s", uri)
	}

	// Split the uri into host and port
	hostAndPort := strings.Split(uri, ":")
	gomega.Expect(len(hostAndPort)).Should(gomega.Equal(2))

	// Parse the port
	port, err := strconv.ParseUint(hostAndPort[1], 10, 32)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse port: %w", err)
	}

	return hostAndPort[0], uint32(port), nil
}

func newTestTeleporterMessage(chainIDInt *big.Int, teleporterAddress common.Address, nonce uint64, data []byte) *types.Transaction {
	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainIDInt,
		Nonce:     nonce,
		To:        &teleporterAddress,
		Gas:       defaultTeleporterMessageGas,
		GasFeeCap: defaultTeleporterMessageGasFeeCap,
		GasTipCap: defaultTeleporterMessageGasTipCap,
		Value:     defaultTeleporterMessageValue,
		Data:      data,
	})
}

func readHexTextFile(filename string) []byte {
	fileData, err := os.ReadFile(filename)
	gomega.Expect(err).Should(gomega.BeNil())
	hexString := utils.SanitizeHashString(string(fileData))
	data, err := hex.DecodeString(hexString)
	gomega.Expect(err).Should(gomega.BeNil())
	return data
}

// TODONOW: remove these once awm-relayer e2e tests merged
// unpack Teleporter message bytes according to EVM ABI encoding rules
func unpackTeleporterMessage(messageBytes []byte) (*teleporter.TeleporterMessage, error) {
	args := abi.Arguments{
		{
			Name: "teleporterMessage",
			Type: teleporter.TeleporterMessageABI,
		},
	}
	unpacked, err := args.Unpack(messageBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack to teleporter message with err: %v", err)
	}
	type teleporterMessageArg struct {
		TeleporterMessage teleporter.TeleporterMessage `json:"teleporterMessage"`
	}
	var teleporterMessage teleporterMessageArg
	err = args.Copy(&teleporterMessage, unpacked)
	if err != nil {
		return nil, err
	}
	return &teleporterMessage.TeleporterMessage, nil
}

// TODONOW: this should be named packSendCrossChainMessageEvent in awm-relayer
func packTeleporterMessage(destinationChainID common.Hash, message teleporter.TeleporterMessage) ([]byte, error) {
	_, hashes, err := teleporter.EVMTeleporterContractABI.PackEvent("SendCrossChainMessage", destinationChainID, message.MessageID, message)
	return hashes, err
}
