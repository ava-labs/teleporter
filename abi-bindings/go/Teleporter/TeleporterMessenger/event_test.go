// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleportermessenger

import (
	"testing"

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
