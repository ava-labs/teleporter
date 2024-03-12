// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	erc20destination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Destination"
	erc20source "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Source"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/Mocks/ExampleERC20"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/gomega"
)

func DeployERC20Source(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	tokenSourceAddress common.Address,
) (common.Address, *erc20source.ERC20Source) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, erc20Source, err := erc20source.DeployERC20Source(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		tokenSourceAddress,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, erc20Source
}

func DeployERC20Destination(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	sourceSubnet interfaces.SubnetTestInfo,
	tokenSourceAddress common.Address,
	tokenName string,
	tokenSymbol string,
	tokenDecimals uint8,

) (common.Address, *erc20destination.ERC20Destination) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, erc20Destination, err := erc20destination.DeployERC20Destination(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		sourceSubnet.BlockchainID,
		tokenSourceAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, erc20Destination
}

func SendERC20Source(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20Source *erc20source.ERC20Source,
	erc20SourceAddress common.Address,
	sourceToken *exampleerc20.ExampleERC20,
	input erc20source.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, [32]byte, *big.Int) {
	// Approve the ERC20Source to spend the tokens
	teleporterUtils.ERC20Approve(
		ctx,
		sourceToken,
		erc20SourceAddress,
		amount,
		subnet,
		senderKey,
	)

	// Send the tokens and verify expected events
	optsA, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20Source.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(BeNil())
	bridgedAmount := big.NewInt(0).Sub(amount, input.PrimaryFee)

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Source.ParseSendTokens)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	Expect(event.Amount).Should(Equal(bridgedAmount))

	return receipt, event.TeleporterMessageID, event.Amount
}

func SendERC20Destination(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20Destination *erc20destination.ERC20Destination,
	erc20DestinationAddress common.Address,
	input erc20destination.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, [32]byte, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20Destination.Approve(
		opts,
		erc20DestinationAddress,
		amount,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	// Bridge the tokens back to subnet A
	tx, err = erc20Destination.Send(
		opts,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	bridgedAmount := big.NewInt(0).Sub(amount, input.PrimaryFee)
	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Destination.ParseSendTokens)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	Expect(event.Amount).Should(Equal(bridgedAmount))

	return receipt, event.TeleporterMessageID, event.Amount
}

func CheckERC20SourceWithdrawal(
	ctx context.Context,
	erc20SourceAddress common.Address,
	sourceToken *exampleerc20.ExampleERC20,
	recipientAddress common.Address,
	amount *big.Int,
	receipt *types.Receipt,
) {
	sourceTransferEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, sourceToken.ParseTransfer)
	Expect(err).Should(BeNil())
	Expect(sourceTransferEvent.From).Should(Equal(erc20SourceAddress))
	Expect(sourceTransferEvent.To).Should(Equal(recipientAddress))
	Expect(sourceTransferEvent.Value).Should(Equal(amount))
}

func CheckERC20DestinationWithdrawal(
	ctx context.Context,
	erc20Destination *erc20destination.ERC20Destination,
	recipientAddress common.Address,
	amount *big.Int,
	receipt *types.Receipt,
) {
	transferEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Destination.ParseTransfer)
	Expect(err).Should(BeNil())
	Expect(transferEvent.From).Should(Equal(common.Address{}))
	Expect(transferEvent.To).Should(Equal(recipientAddress))
	Expect(transferEvent.Value).Should(Equal(amount))
}
