// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleportermessenger

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// Event is a Teleporter log event
type Event uint8

const (
	Unknown Event = iota
	SendCrossChainMessage
	ReceiveCrossChainMessage
	AddFeeAmount
	MessageExecutionFailed
	MessageExecuted
	RelayerRewardsRedeemed
	ReceiptReceived

	sendCrossChainMessageStr    = "SendCrossChainMessage"
	receiveCrossChainMessageStr = "ReceiveCrossChainMessage"
	addFeeAmountStr             = "AddFeeAmount"
	messageExecutionFailedStr   = "MessageExecutionFailed"
	messageExecutedStr          = "MessageExecuted"
	relayerRewardsRedeemedStr   = "RelayerRewardsRedeemed"
	receiptReceivedStr          = "ReceiptReceived"
	unknownStr                  = "Unknown"
)

// String returns the string representation of an Event
func (e Event) String() string {
	switch e {
	case SendCrossChainMessage:
		return sendCrossChainMessageStr
	case ReceiveCrossChainMessage:
		return receiveCrossChainMessageStr
	case AddFeeAmount:
		return addFeeAmountStr
	case MessageExecutionFailed:
		return messageExecutionFailedStr
	case MessageExecuted:
		return messageExecutedStr
	case RelayerRewardsRedeemed:
		return relayerRewardsRedeemedStr
	case ReceiptReceived:
		return receiptReceivedStr
	default:
		return unknownStr
	}
}

// ToEvent converts a string to an Event
func ToEvent(e string) (Event, error) {
	switch strings.ToLower(e) {
	case strings.ToLower(sendCrossChainMessageStr):
		return SendCrossChainMessage, nil
	case strings.ToLower(receiveCrossChainMessageStr):
		return ReceiveCrossChainMessage, nil
	case strings.ToLower(addFeeAmountStr):
		return AddFeeAmount, nil
	case strings.ToLower(messageExecutionFailedStr):
		return MessageExecutionFailed, nil
	case strings.ToLower(messageExecutedStr):
		return MessageExecuted, nil
	case strings.ToLower(relayerRewardsRedeemedStr):
		return RelayerRewardsRedeemed, nil
	case strings.ToLower(receiptReceivedStr):
		return ReceiptReceived, nil
	default:
		return Unknown, fmt.Errorf("unknown event %s", e)
	}
}

// FilterTeleporterEvents parses the topics and data of a Teleporter log into the corresponding Teleporter event
func FilterTeleporterEvents(topics []common.Hash, data []byte, event string) (interface{}, error) {
	e, err := ToEvent(event)
	if err != nil {
		return nil, err
	}
	var out interface{}
	switch e {
	case SendCrossChainMessage:
		out = new(TeleporterMessengerSendCrossChainMessage)
	case ReceiveCrossChainMessage:
		out = new(TeleporterMessengerReceiveCrossChainMessage)
	case AddFeeAmount:
		out = new(TeleporterMessengerAddFeeAmount)
	case MessageExecutionFailed:
		out = new(TeleporterMessengerMessageExecutionFailed)
	case MessageExecuted:
		out = new(TeleporterMessengerMessageExecuted)
	case RelayerRewardsRedeemed:
		out = new(TeleporterMessengerRelayerRewardsRedeemed)
	case ReceiptReceived:
		out = new(TeleporterMessengerReceiptReceived)
	default:
		return nil, fmt.Errorf("unknown event %s", e.String())
	}
	if err := UnpackEvent(out, e.String(), topics, data); err != nil {
		return nil, err
	}
	return out, nil
}
