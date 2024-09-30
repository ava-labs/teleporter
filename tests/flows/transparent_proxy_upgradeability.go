package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20tokenhome "github.com/ava-labs/teleporter/abi-bindings/go/ictt/TokenHome/ERC20TokenHome"
	erc20tokenhomeupgradeable "github.com/ava-labs/teleporter/abi-bindings/go/ictt/TokenHome/ERC20TokenHomeUpgradeable"
	erc20tokenremote "github.com/ava-labs/teleporter/abi-bindings/go/ictt/TokenRemote/ERC20TokenRemote"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy an upgradeable ERC20TokenHome on the primary network
 * Deploys a transparent upgradeable proxy that uses the ERC20TokenHome logic contract
 * Deploys a proxy admin contract to manage the upgradeable proxy
 * Deploy an ERC20TokenRemote to Subnet A
 * Transfers example erc20 tokens from the primary network to Subnet A
 * Deploy a new ERC20TokenHome logic contract on the primary network
 * Upgrade the transparent upgradeable proxy to use the new logic contract
 * Transfer tokens from Subnet A back to the primary network
 * Check that the transfer was successful, and expected balances are correct
 */

func TransparentUpgradeableProxy(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on the primary network as the token to be transferred
	exampleERC20Address, exampleERC20 := utils.DeployExampleERC20Decimals(
		ctx,
		fundedKey,
		cChainInfo,
		erc20TokenHomeDecimals,
	)

	tokenName, err := exampleERC20.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := exampleERC20.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey,
		cChainInfo.EVMChainID,
	)
	Expect(err).Should(BeNil())
	implAddress, tx, _, err := erc20tokenhomeupgradeable.DeployERC20TokenHomeUpgradeable(
		opts,
		cChainInfo.RPCClient,
		uint8(1),
	)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())

	// Deploy a TransparentUpgradeableProxy contract on primary network for the ERC20TokenHome logic contract
	erc20TokenHomeAddress, proxyAdmin, erc20TokenHome := utils.DeployTransparentUpgradeableProxy(
		ctx,
		cChainInfo,
		fundedKey,
		implAddress,
		erc20tokenhome.NewERC20TokenHome,
	)

	tx, err = erc20TokenHome.Initialize(
		opts,
		cChainInfo.TeleporterRegistryAddress,
		fundedAddress,
		big.NewInt(1),
		exampleERC20Address,
		tokenDecimals,
	)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())

	// Deploy the ERC20TokenRemote contract on Subnet A
	erc20TokenRemoteAddress, erc20TokenRemote := utils.DeployERC20TokenRemote(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		tokenDecimals,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		network,
		cChainInfo,
		erc20TokenHomeAddress,
		subnetAInfo,
		erc20TokenRemoteAddress,
	)

	// Send a transfer from primary network to Subnet A
	// Generate new recipient to receive transferred tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	input := erc20tokenhome.SendTokensInput{
		DestinationBlockchainID:            subnetAInfo.BlockchainID,
		DestinationTokenTransferrerAddress: erc20TokenRemoteAddress,
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

	// Check that the transfer was successful, and expected balances are correct
	receipt = network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	utils.CheckERC20TokenRemoteWithdrawal(
		ctx,
		erc20TokenRemote,
		receipt,
		recipientAddress,
		transferredAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20TokenRemote.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(transferredAmount))

	// Deploy a new ERC20TokenHome logic contract on primary network
	newLogic, tx, _, err := erc20tokenhomeupgradeable.DeployERC20TokenHomeUpgradeable(
		opts,
		cChainInfo.RPCClient,
		uint8(1),
	)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())

	// Upgrade the TransparentUpgradeableProxy contract to use the new logic contract
	tx, err = proxyAdmin.UpgradeAndCall(opts, erc20TokenHomeAddress, newLogic, []byte{})
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())

	// Send a transfer from Subnet A back to primary network
	teleporterUtils.SendNativeTransfer(
		ctx,
		subnetAInfo,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	inputB := erc20tokenremote.SendTokensInput{
		DestinationBlockchainID:            cChainInfo.BlockchainID,
		DestinationTokenTransferrerAddress: erc20TokenHomeAddress,
		Recipient:                          recipientAddress,
		PrimaryFeeTokenAddress:             erc20TokenRemoteAddress,
		PrimaryFee:                         big.NewInt(1e10),
		SecondaryFee:                       big.NewInt(0),
		RequiredGasLimit:                   utils.DefaultERC20RequiredGas,
	}

	receipt, transferredAmount = utils.SendERC20TokenRemote(
		ctx,
		subnetAInfo,
		erc20TokenRemote,
		erc20TokenRemoteAddress,
		inputB,
		teleporterUtils.BigIntSub(transferredAmount, inputB.PrimaryFee),
		recipientKey,
	)

	receipt = network.RelayMessage(
		ctx,
		receipt,
		subnetAInfo,
		cChainInfo,
		true,
	)

	// Check that the transfer was successful, and expected balances are correct
	utils.CheckERC20TokenHomeWithdrawal(
		ctx,
		erc20TokenHomeAddress,
		exampleERC20,
		receipt,
		recipientAddress,
		transferredAmount,
	)

	// Check that the recipient received the tokens
	balance, err = exampleERC20.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(transferredAmount))
}
