// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleportermessenger

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestEventString(t *testing.T) {
	var (
		tests = []struct {
			event Event
			str   string
		}{
			{Unknown, unknownStr},
			{SendCrossChainMessage, sendCrossChainMessageStr},
			{ReceiveCrossChainMessage, receiveCrossChainMessageStr},
			{AddFeeAmount, addFeeAmountStr},
			{MessageExecutionFailed, messageExecutionFailedStr},
			{MessageExecuted, messageExecutedStr},
			{RelayerRewardsRedeemed, relayerRewardsRedeemedStr},
		}
	)

	for _, test := range tests {
		t.Run(test.event.String(), func(t *testing.T) {
			require.Equal(t, test.event.String(), test.str)
		})
	}
}

func TestToEvent(t *testing.T) {
	var (
		tests = []struct {
			str     string
			event   Event
			isError bool
		}{
			{unknownStr, Unknown, true},
			{sendCrossChainMessageStr, SendCrossChainMessage, false},
			{receiveCrossChainMessageStr, ReceiveCrossChainMessage, false},
			{addFeeAmountStr, AddFeeAmount, false},
			{messageExecutionFailedStr, MessageExecutionFailed, false},
			{messageExecutedStr, MessageExecuted, false},
			{relayerRewardsRedeemedStr, RelayerRewardsRedeemed, false},
		}
	)

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			event, err := ToEvent(test.str)
			if test.isError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, test.event, event)
		})
	}
}

func TestFilterTeleporterEvents(t *testing.T) {
	mockBlockchainID := [32]byte{1, 2, 3, 4}
	messageID := big.NewInt(1)
	message := createTestTeleporterMessage(messageID.Int64())
	feeInfo := TeleporterFeeInfo{
		FeeTokenAddress: common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		Amount:          big.NewInt(1),
	}
	deliverer := common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567")

	teleporterABI, err := TeleporterMessengerMetaData.GetAbi()
	require.NoError(t, err)

	var (
		tests = []struct {
			event    Event
			args     []interface{}
			expected interface{}
		}{
			{
				event: SendCrossChainMessage,
				args: []interface{}{
					mockBlockchainID,
					messageID,
					message,
					feeInfo,
				},
				expected: &TeleporterMessengerSendCrossChainMessage{
					DestinationBlockchainID: mockBlockchainID,
					MessageID:               messageID,
					Message:                 message,
					FeeInfo:                 feeInfo,
				},
			},
			{
				event: ReceiveCrossChainMessage,
				args: []interface{}{
					mockBlockchainID,
					messageID,
					deliverer,
					deliverer,
					message,
				},
				expected: &TeleporterMessengerReceiveCrossChainMessage{
					OriginBlockchainID: mockBlockchainID,
					MessageID:          messageID,
					Deliverer:          deliverer,
					RewardRedeemer:     deliverer,
					Message:            message,
				},
			},
			{
				event: MessageExecuted,
				args: []interface{}{
					mockBlockchainID,
					messageID,
				},
				expected: &TeleporterMessengerMessageExecuted{
					OriginBlockchainID: mockBlockchainID,
					MessageID:          messageID,
				},
			},
		}
	)

	for _, test := range tests {
		t.Run(test.event.String(), func(t *testing.T) {
			topics, data, err := teleporterABI.PackEvent(test.event.String(), test.args...)
			require.NoError(t, err)

			out, err := FilterTeleporterEvents(topics, data, test.event.String())
			require.NoError(t, err)

			require.Equal(t, test.expected, out)
		})
	}
}
