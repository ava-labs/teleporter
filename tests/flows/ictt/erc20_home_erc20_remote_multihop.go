package ictt

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20tokenhome "github.com/ava-labs/teleporter/abi-bindings/go/ictt/TokenHome/ERC20TokenHome"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a ERC20 token home on the primary network
 * Deploys ERC20 token remote to L1 A and L1 B
 * Transfers C-Chain example ERC20 tokens to L1 A
 * Transfer tokens from L1 A to L1 B through multi-hop
 * Transfer back tokens from L1 B to L1 A through multi-hop
 */
func ERC20TokenHomeERC20TokenRemoteMultiHop(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	L1AInfo, L1BInfo := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on L1 A as the token to be transferred
	exampleERC20Address, exampleERC20 := utils.DeployExampleERC20Decimals(
		ctx,
		fundedKey,
		cChainInfo,
		erc20TokenHomeDecimals,
	)

	homeTokenDecimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Create an ERC20TokenHome for transferring the ERC20 token
	erc20TokenHomeAddress, erc20TokenHome := utils.DeployERC20TokenHome(
		ctx,
		teleporter,
		fundedKey,
		cChainInfo,
		fundedAddress,
		exampleERC20Address,
		homeTokenDecimals,
	)

	// Token representation on L1s A and B will have same name, symbol, and decimals
	tokenName, err := exampleERC20.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := exampleERC20.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20TokenRemote tp L1 A
	erc20TokenRemoteAddressA, erc20TokenRemoteA := utils.DeployERC20TokenRemote(
		ctx,
		teleporter,
		fundedKey,
		L1AInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		homeTokenDecimals,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Deploy an ERC20TokenRemote for L1 B
	erc20TokenRemoteAddressB, erc20TokenRemoteB := utils.DeployERC20TokenRemote(
		ctx,
		teleporter,
		fundedKey,
		L1BInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		homeTokenDecimals,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Register both ERC20TokenRemote instances on the ERC20TokenHome
	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		teleporter,
		cChainInfo,
		erc20TokenHomeAddress,
		L1AInfo,
		erc20TokenRemoteAddressA,
		fundedKey,
	)
	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		teleporter,
		cChainInfo,
		erc20TokenHomeAddress,
		L1BInfo,
		erc20TokenRemoteAddressB,
		fundedKey,
	)

	// Generate new recipient to receive transferred tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to L1 A
	input := erc20tokenhome.SendTokensInput{
		DestinationBlockchainID:            L1AInfo.BlockchainID,
		DestinationTokenTransferrerAddress: erc20TokenRemoteAddressA,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             exampleERC20Address,
		PrimaryFee:                         big.NewInt(1e18),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultERC20RequiredGas,
	}
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(13))

	receipt, transferredAmount := utils.SendERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		input,
		amount,
		fundedKey,
	)

	// Relay the message to L1 A and check for ERC20TokenRemote withdrawal
	receipt = teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		cChainInfo,
		L1AInfo,
		true,
		fundedKey,
	)

	utils.CheckERC20TokenRemoteWithdrawal(
		ctx,
		erc20TokenRemoteA,
		receipt,
		recipientAddress,
		transferredAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20TokenRemoteA.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(transferredAmount))

	// Multi-hop transfer to L1 B
	transferredAmount = big.NewInt(0).Div(transferredAmount, big.NewInt(2))
	secondaryFeeAmount := big.NewInt(0).Div(transferredAmount, big.NewInt(4))
	utils.SendERC20TokenMultiHopAndVerify(
		ctx,
		teleporter,
		fundedKey,
		recipientKey,
		recipientAddress,
		L1AInfo,
		erc20TokenRemoteA,
		erc20TokenRemoteAddressA,
		L1BInfo,
		erc20TokenRemoteB,
		erc20TokenRemoteAddressB,
		cChainInfo,
		transferredAmount,
		secondaryFeeAmount,
	)

	// Multi-hop transfer back to L1 A
	transferredAmount = big.NewInt(0).Sub(transferredAmount, secondaryFeeAmount)
	secondaryFeeAmount = big.NewInt(0).Div(transferredAmount, big.NewInt(4))
	utils.SendERC20TokenMultiHopAndVerify(
		ctx,
		teleporter,
		fundedKey,
		recipientKey,
		recipientAddress,
		L1BInfo,
		erc20TokenRemoteB,
		erc20TokenRemoteAddressB,
		L1AInfo,
		erc20TokenRemoteA,
		erc20TokenRemoteAddressA,
		cChainInfo,
		transferredAmount,
		secondaryFeeAmount,
	)
}
