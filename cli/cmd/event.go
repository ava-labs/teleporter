// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

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

var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Parses a Teleporter log's topics and data",
	Long: `Given the topics and data of a Teleporter log, parses the log into
	the corresponding Teleporter event. Topics are represented by a hash, and data
	is the hex encoding of the bytes. For example:

	teleporter-cli event --topics topic1, topic2 --data 0x1234`,
	Args: cobra.NoArgs,

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
