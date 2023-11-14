/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile       string
	logger        logging.Logger
	teleporterABI *abi.ABI
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "teleporter-cli",
	Short: "A CLI that integrates with the Teleporter protocol",
	Long: `A CLI that integrates with the Teleporter protocol, and allows you
	to debug Teleporter on chain activity. The CLI can help decode Teleporter data
	as well as query for on chain state`,
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	logLevelArg := rootCmd.PersistentFlags().StringP("log", "l", "", "Log level")
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
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
				logging.JSON.ConsoleEncoder(),
			),
		)
		abi, err := teleportermessenger.TeleporterMessengerMetaData.GetAbi()
		if err != nil {
			return err
		}
		teleporterABI = abi
		return nil
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
