// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleportermessenger

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/core/types"
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
		out = new(ReadableTeleporterMessengerSendCrossChainMessage)
	case ReceiveCrossChainMessage:
		out = new(ReadableTeleporterMessengerReceiveCrossChainMessage)
	case AddFeeAmount:
		out = new(ReadableTeleporterMessengerAddFeeAmount)
	case MessageExecutionFailed:
		out = new(ReadableTeleporterMessengerMessageExecutionFailed)
	case MessageExecuted:
		out = new(ReadableTeleporterMessengerMessageExecuted)
	case RelayerRewardsRedeemed:
		out = new(TeleporterMessengerRelayerRewardsRedeemed)
	case ReceiptReceived:
		out = new(ReadableTeleporterMessengerReceiptReceived)
	default:
		return nil, fmt.Errorf("unknown event %s", e.String())
	}
	if err := UnpackEvent(out, e.String(), topics, data); err != nil {
		return nil, err
	}
	return out, nil
}

type ReadableTeleporterMessengerSendCrossChainMessage struct {
	MessageID               common.Hash
	DestinationBlockchainID ids.ID
	Message                 ReadableTeleporterMessage
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log
}

type ReadableTeleporterMessengerReceiveCrossChainMessage struct {
	MessageID          common.Hash
	SourceBlockchainID ids.ID
	Deliverer          common.Address
	RewardRedeemer     common.Address
	Message            ReadableTeleporterMessage
	Raw                types.Log
}

type ReadableTeleporterMessengerAddFeeAmount struct {
	MessageID      common.Hash
	UpdatedFeeInfo TeleporterFeeInfo
	Raw            types.Log
}

type ReadableTeleporterMessengerMessageExecutionFailed struct {
	MessageID          common.Hash
	SourceBlockchainID ids.ID
	Message            ReadableTeleporterMessage
	Raw                types.Log
}

type ReadableTeleporterMessengerMessageExecuted struct {
	MessageID          common.Hash
	SourceBlockchainID ids.ID
	Raw                types.Log
}

type ReadableTeleporterMessengerReceiptReceived struct {
	MessageID               common.Hash
	DestinationBlockchainID ids.ID
	RelayerRewardAddress    common.Address
	FeeInfo                 TeleporterFeeInfo
	Raw                     types.Log
}

// TeleporterMessage is an auto generated low-level Go binding around an user-defined struct.
type ReadableTeleporterMessage struct {
	MessageNonce            *big.Int
	OriginSenderAddress     common.Address
	DestinationBlockchainID ids.ID
	DestinationAddress      common.Address
	RequiredGasLimit        *big.Int
	AllowedRelayerAddresses []common.Address
	Receipts                []TeleporterMessageReceipt
	Message                 []byte
}
