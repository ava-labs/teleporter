// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func CalculateMessageID(
	teleporterMessengerAddress common.Address,
	sourceBlockchainID ids.ID,
	destinationBlockchainID ids.ID,
	nonce *big.Int,
) (ids.ID, error) {
	uint256Ty, _ := abi.NewType("uint256", "uint256", nil)
	bytes32Ty, _ := abi.NewType("bytes32", "bytes32", nil)
	addressTy, _ := abi.NewType("address", "address", nil)

	arguments := abi.Arguments{
		{Type: addressTy},
		{Type: bytes32Ty},
		{Type: bytes32Ty},
		{Type: uint256Ty},
	}

	bytes, err := arguments.Pack(
		teleporterMessengerAddress,
		sourceBlockchainID,
		destinationBlockchainID,
		nonce,
	)
	if err != nil {
		return ids.ID{}, err
	}

	return ids.ID(crypto.Keccak256Hash(bytes)), nil
}
