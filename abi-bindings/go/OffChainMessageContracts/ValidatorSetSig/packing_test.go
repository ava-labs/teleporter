// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorsetsig

import (
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestPackUnpackValidatorSetSigMessage(t *testing.T) {
	msg := ValidatorSetSigMessage{
		ValidatorSetSigAddress: common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		TargetContractAddress:  common.HexToAddress("0x0123456789abcdef0123456789abcdef01234568"),
		TargetBlockChainID:     ids.ID{1, 2, 3, 4},
		Nonce:                  big.NewInt(1),
		TxPayload:              []byte{1, 2, 3, 4},
	}
	b, err := PackValidatorSetSigWarpPayload(msg)
	require.NoError(t, err)

	unpackedMsg, err := UnpackValidatorSetSigWarpPayload(b)
	require.NoError(t, err)

	require.Equal(t, msg, unpackedMsg)

}
