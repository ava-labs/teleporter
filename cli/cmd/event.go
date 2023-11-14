/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	topicArgs []string
	data      []byte
)

// eventCmd represents the event command
var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(0),

	Run: func(cmd *cobra.Command, args []string) {
		var topics []common.Hash
		for _, topic := range topicArgs {
			topics = append(topics, common.HexToHash(topic))
		}

		err := FilterTeleporterEvents(topics, data)
		cobra.CheckErr(err)
	},
}

func FilterTeleporterEvents(topics []common.Hash, data []byte) error {
	event, err := teleporterABI.EventByID(topics[0])
	if err != nil {
		return err
	}
	var out interface{}
	switch event.Name {
	case "SendCrossChainMessage":
		out = new(teleportermessenger.TeleporterMessengerSendCrossChainMessage)
	case "ReceiveCrossChainMessage":
		out = new(teleportermessenger.TeleporterMessengerReceiveCrossChainMessage)
	case "AddFeeAmount":
		out = new(teleportermessenger.TeleporterMessengerAddFeeAmount)
	case "MessageExecutedFailed":
		out = new(teleportermessenger.TeleporterMessengerMessageExecutionFailed)
	case "MessageExecuted":
		out = new(teleportermessenger.TeleporterMessengerMessageExecuted)
	case "RelayerRewardsRedeemed":
		out = new(teleportermessenger.TeleporterMessengerRelayerRewardsRedeemed)
	default:
		logger.Error("Unknown event", zap.String("event", event.Name))
		return fmt.Errorf("unknown event %s", event.Name)
	}
	if err = UnpackEvent(out, event.Name, topics, data); err != nil {
		return err
	}
	logger.Info(fmt.Sprintf("Event %s unpacked", event.Name), zap.Any("event", out))
	return nil
}

func UnpackEvent(out interface{}, event string, topics []common.Hash, data []byte) error {
	if len(data) > 0 {
		if err := teleporterABI.UnpackIntoInterface(out, event, data); err != nil {
			return err
		}
	}

	var indexed abi.Arguments
	for _, arg := range teleporterABI.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, topics[1:])
}

func init() {
	rootCmd.AddCommand(eventCmd)
	eventCmd.PersistentFlags().StringSliceVar(&topicArgs, "topics", []string{}, "The topics of the event")
	eventCmd.Flags().BytesHexVar(&data, "data", []byte{}, "The data of the event")

	err := eventCmd.MarkPersistentFlagRequired("topics")
	cobra.CheckErr(err)
}
