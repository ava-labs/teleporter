// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cmd

import (
	"context"

	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/x/warp"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	warpPrecompileAddress = "0x0200000000000000000000000000000000000005"
)

var (
	rpcEndpoint     string
	contractAddress common.Address
	client          ethclient.Client
)

var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Parses relevant Teleporter logs from a transaction",
	Long: `Given a transaction this command looks through the transaction's receipt
	for Teleporter and Warp log events. When corresponding log events are found,
	the command parses to log event fields to a more human readable format. For example:

	teleporter-cli transaction 0x1234 --rpc $RPC_URL --contract-address 0x1234`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		receipt, err := client.TransactionReceipt(context.Background(),
			common.HexToHash(args[0]))
		cobra.CheckErr(err)
		for _, log := range receipt.Logs {
			if log.Address == contractAddress {
				logger.Debug("Processing Teleporter log", zap.Any("log", log))

				event, err := teleporterABI.EventByID(log.Topics[0])
				cobra.CheckErr(err)

				out, err := teleportermessenger.FilterTeleporterEvents(log.Topics, log.Data, event.Name)
				cobra.CheckErr(err)
				logger.Info("Parsed Teleporter event", zap.Any("event", out))
			}

			if log.Address == common.HexToAddress(warpPrecompileAddress) {
				logger.Debug("Processing Warp log", zap.Any("log", log))

				unsignedMsg, err := warp.UnpackSendWarpEventDataToMessage(log.Data)
				cobra.CheckErr(err)

				warpPayload, err := warpPayload.ParseAddressedCall(unsignedMsg.Payload)
				cobra.CheckErr(err)

				teleporterMessage, err := teleportermessenger.UnpackTeleporterMessage(warpPayload.Payload)
				cobra.CheckErr(err)
				logger.Info("Parsed Teleporter message", zap.Any("message", teleporterMessage))
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(transactionCmd)
	transactionCmd.PersistentFlags().StringVar(&rpcEndpoint, "rpc", "", "RPC endpoint to connect to the node")
	address := transactionCmd.PersistentFlags().StringP("contract-address", "c", "", "Teleporter contract address")
	err := transactionCmd.MarkPersistentFlagRequired("rpc")
	cobra.CheckErr(err)
	err = transactionCmd.MarkPersistentFlagRequired("contract-address")
	cobra.CheckErr(err)
	transactionCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if rootCmd.PersistentPreRunE != nil {
			if err := rootCmd.PersistentPreRunE(cmd, args); err != nil {
				return err
			}
		}
		contractAddress = common.HexToAddress(*address)
		c, err := ethclient.Dial(rpcEndpoint)
		if err != nil {
			return err
		}

		client = c
		return err
	}
}
