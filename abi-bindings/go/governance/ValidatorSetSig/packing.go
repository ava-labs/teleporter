// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorsetsig

import (
	"fmt"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/pkg/errors"
)

var validatorSetSigMessageType abi.Type

func init() {
	// Create an ABI binding for ValidatorSetSigMessage, defined in ValidatorSetSig.sol
	// abigen does not support ABI bindings for standalone structs, only methods and events,
	// so we must manually keep this up-to-date with the struct defined in the contract.
	var err error
	validatorSetSigMessageType, err = abi.NewType("tuple", "struct Overloader.F", []abi.ArgumentMarshaling{
		{Name: "targetBlockchainID", Type: "bytes32"},
		{Name: "validatorSetSigAddress", Type: "address"},
		{Name: "targetContractAddress", Type: "address"},
		{Name: "nonce", Type: "uint256"},
		{Name: "value", Type: "uint256"},
		{Name: "payload", Type: "bytes"},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create ValidatorSetSigMessage ABI type: %v", err))
	}
}

func (m *ValidatorSetSigMessage) Pack() ([]byte, error) {
	args := abi.Arguments{
		{
			Name: "validatorSetSigMessage",
			Type: validatorSetSigMessageType,
		},
	}
	return args.Pack(m)
}

func (m *ValidatorSetSigMessage) Unpack(messageBytes []byte) error {
	args := abi.Arguments{
		{
			Name: "validatorSetSigMessage",
			Type: validatorSetSigMessageType,
		},
	}
	unpacked, err := args.Unpack(messageBytes)
	if err != nil {
		return fmt.Errorf("failed to unpack to ValidatorSetSigMessage with err: %v", err)
	}
	return args.Copy(&m, unpacked)
}

// PackExecuteCall packs the input to form a call to the executeCall function
func PackExecuteCall(messageIndex uint32) ([]byte, error) {
	abi, err := ValidatorSetSigMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}
	return abi.Pack("executeCall", messageIndex)
}
