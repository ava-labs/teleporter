// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleportermessenger

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func createTestTeleporterMessage(messageID int64) TeleporterMessage {
	m := TeleporterMessage{
		MessageID:               big.NewInt(messageID),
		SenderAddress:           common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		DestinationBlockchainID: [32]byte{1, 2, 3, 4},
		DestinationAddress:      common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		RequiredGasLimit:        big.NewInt(2),
		AllowedRelayerAddresses: []common.Address{
			common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		},
		Receipts: []TeleporterMessageReceipt{
			{
				ReceivedMessageID:    big.NewInt(1),
				RelayerRewardAddress: common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
			},
		},
		Message: []byte{1, 2, 3, 4},
	}
	return m
}

func TestPackUnpackTeleporterMessage(t *testing.T) {
	var (
		messageID int64 = 4
	)
	message := createTestTeleporterMessage(messageID)

	b, err := PackTeleporterMessage(message)
	if err != nil {
		t.Errorf("failed to pack teleporter message: %v", err)
		t.FailNow()
	}

	unpacked, err := UnpackTeleporterMessage(b)
	if err != nil {
		t.Errorf("failed to unpack teleporter message: %v", err)
		t.FailNow()
	}

	for i := 0; i < len(message.AllowedRelayerAddresses); i++ {
		require.Equal(t, unpacked.AllowedRelayerAddresses[i], message.AllowedRelayerAddresses[i])
	}

	for i := 0; i < len(message.Receipts); i++ {
		require.Equal(t, message.Receipts[i].ReceivedMessageID, unpacked.Receipts[i].ReceivedMessageID)
		require.Equal(t, message.Receipts[i].RelayerRewardAddress, unpacked.Receipts[i].RelayerRewardAddress)
	}

	require.True(t, bytes.Equal(message.Message, unpacked.Message))
}

func TestUnpackEvent(t *testing.T) {
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
			out      interface{}
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
				out: new(TeleporterMessengerSendCrossChainMessage),
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
				out: new(TeleporterMessengerReceiveCrossChainMessage),
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
				out: new(TeleporterMessengerMessageExecuted),
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

			err = UnpackEvent(test.out, test.event.String(), topics, data)
			require.NoError(t, err)

			require.Equal(t, test.expected, test.out)
		})
	}
}
