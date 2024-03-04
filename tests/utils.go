// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"strings"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/awm-relayer/config"
	offchainregistry "github.com/ava-labs/awm-relayer/messages/off-chain-registry"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	teleporterTestUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

var (
	// Write the test database to /tmp since the data is not needed after the test
	storageLocation = fmt.Sprintf("%s/.awm-relayer-storage", os.TempDir())
)

func BuildAndRunRelayerExecutable(ctx context.Context, relayerConfigPath string) context.CancelFunc {
	// Build the awm-relayer binary
	cmd := exec.Command("./scripts/build.sh")
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	Expect(err).Should(BeNil())

	cmdOutput := make(chan string)

	// Run awm relayer binary with config path
	var relayerContext context.Context
	relayerContext, relayerCancel := context.WithCancel(ctx)
	relayerCmd := exec.CommandContext(relayerContext, "./build/awm-relayer", "--config-file", relayerConfigPath)

	// Set up a pipe to capture the command's output
	cmdStdOutReader, err := relayerCmd.StdoutPipe()
	Expect(err).Should(BeNil())
	cmdStdErrReader, err := relayerCmd.StderrPipe()
	Expect(err).Should(BeNil())

	// Start the command
	log.Info("Starting the relayer executable")
	err = relayerCmd.Start()
	Expect(err).Should(BeNil())

	// Start goroutines to read and output the command's stdout and stderr
	go func() {
		scanner := bufio.NewScanner(cmdStdOutReader)
		for scanner.Scan() {
			log.Info(scanner.Text())
		}
		cmdOutput <- "Command execution finished"
	}()
	go func() {
		scanner := bufio.NewScanner(cmdStdErrReader)
		for scanner.Scan() {
			log.Error(scanner.Text())
		}
		cmdOutput <- "Command execution finished"
	}()
	return func() {
		relayerCancel()
		relayerCmd.Wait()
	}
}

func ReadHexTextFile(filename string) string {
	fileData, err := os.ReadFile(filename)
	Expect(err).Should(BeNil())
	return strings.TrimRight(string(fileData), "\n")
}

// Constructs a relayer config with all subnets as sources and destinations
func CreateDefaultRelayerConfig(
	subnetsInfo []interfaces.SubnetTestInfo,
	teleporterContractAddress common.Address,
	fundedAddress common.Address,
	relayerKey *ecdsa.PrivateKey,
) config.Config {
	log.Info(
		"Setting up relayer config",
	)
	// Construct the config values for each subnet
	sources := make([]*config.SourceSubnet, len(subnetsInfo))
	destinations := make([]*config.DestinationSubnet, len(subnetsInfo))
	for i, subnetInfo := range subnetsInfo {
		host, port, err := teleporterTestUtils.GetURIHostAndPort(subnetInfo.NodeURIs[0])
		Expect(err).Should(BeNil())

		sources[i] = &config.SourceSubnet{
			SubnetID:          subnetInfo.SubnetID.String(),
			BlockchainID:      subnetInfo.BlockchainID.String(),
			VM:                config.EVM.String(),
			EncryptConnection: false,
			APINodeHost:       host,
			APINodePort:       port,
			MessageContracts: map[string]config.MessageProtocolConfig{
				teleporterContractAddress.Hex(): {
					MessageFormat: config.TELEPORTER.String(),
					Settings: map[string]interface{}{
						"reward-address": fundedAddress.Hex(),
					},
				},
				offchainregistry.OffChainRegistrySourceAddress.Hex(): {
					MessageFormat: config.OFF_CHAIN_REGISTRY.String(),
					Settings: map[string]interface{}{
						"teleporter-registry-address": subnetInfo.TeleporterRegistryAddress.Hex(),
					},
				},
			},
		}

		destinations[i] = &config.DestinationSubnet{
			SubnetID:          subnetInfo.SubnetID.String(),
			BlockchainID:      subnetInfo.BlockchainID.String(),
			VM:                config.EVM.String(),
			EncryptConnection: false,
			APINodeHost:       host,
			APINodePort:       port,
			AccountPrivateKey: hex.EncodeToString(relayerKey.D.Bytes()),
		}

		log.Info(
			"Creating relayer config for subnet",
			"subnetID", subnetInfo.SubnetID.String(),
			"blockchainID", subnetInfo.BlockchainID.String(),
			"host", host,
			"port", port,
		)
	}

	return config.Config{
		LogLevel: logging.Info.LowerString(),
		// TODO: There's currently a bug in ANR v1.7.4-rc.0 that specifies the network ID as 0. We should change this back to constants.DefaultNetworkID once fixed.
		NetworkID:           0,
		PChainAPIURL:        subnetsInfo[0].NodeURIs[0],
		EncryptConnection:   false,
		StorageLocation:     RelayerStorageLocation(),
		ProcessMissedBlocks: false,
		SourceSubnets:       sources,
		DestinationSubnets:  destinations,
	}
}

func RelayerStorageLocation() string {
	return storageLocation
}

func ClearRelayerStorage() error {
	return os.RemoveAll(storageLocation)
}

func FundRelayers(
	ctx context.Context,
	subnetsInfo []interfaces.SubnetTestInfo,
	fundedKey *ecdsa.PrivateKey,
	relayerKey *ecdsa.PrivateKey,
) {
	relayerAddress := crypto.PubkeyToAddress(relayerKey.PublicKey)
	fundAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)) // 10eth

	for _, subnetInfo := range subnetsInfo {
		fundRelayerTx := utils.CreateNativeTransferTransaction(
			ctx, subnetInfo, fundedKey, relayerAddress, fundAmount,
		)
		utils.SendTransactionAndWaitForSuccess(ctx, subnetInfo, fundRelayerTx)
	}
}

// SanitizeHexString removes the "0x" prefix from a hex string if it exists.
// Otherwise, returns the original string.
func SanitizeHexString(hex string) string {
	return strings.TrimPrefix(hex, "0x")
}
