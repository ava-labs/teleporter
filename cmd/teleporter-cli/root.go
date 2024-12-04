// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"os"

	"github.com/ava-labs/avalanchego/utils/logging"
	teleportermessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/TeleporterMessenger"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/spf13/cobra"
)

var (
	logger        logging.Logger
	teleporterABI *abi.ABI
)

var rootCmd = &cobra.Command{
	Use:   "teleporter-cli",
	Short: "A CLI that integrates with the Teleporter protocol",
	Long: `A CLI that integrates with the Teleporter protocol, and allows you
to debug Teleporter on chain activity. The CLI can help decode
TeleporterMessenger and ICM events, as well as parsing Teleporter messages.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	logLevelArg := rootCmd.PersistentFlags().StringP("log", "l", "", "Log level i.e. debug, info...")
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		return rootPreRunE(logLevelArg)
	}
}

func rootPreRunE(logLevelArg *string) error {
	if *logLevelArg == "" {
		*logLevelArg = logging.Info.LowerString()
	}

	logLevel, err := logging.ToLevel(*logLevelArg)
	if err != nil {
		return err
	}
	logger = logging.NewLogger(
		"teleporter-cli",
		logging.NewWrappedCore(
			logLevel,
			os.Stdout,
			logging.Plain.ConsoleEncoder(),
		),
	)
	abi, err := teleportermessenger.TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return err
	}
	teleporterABI = abi
	return nil
}

func callPersistentPreRunE(cmd *cobra.Command, args []string) error {
	if parent := cmd.Parent(); parent != nil {
		if parent.PersistentPreRunE != nil {
			return parent.PersistentPreRunE(parent, args)
		}
	}
	return nil
}

func main() {
	Execute()
}
