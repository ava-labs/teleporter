package tests

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	erc20tokensource "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/NativeTokenBridge/ERC20TokenSource"
	nativetokendestination "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/NativeTokenBridge/NativeTokenDestination"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/Mocks/ExampleERC20"
	"github.com/ava-labs/teleporter/tests/utils"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func ERC20ToNativeTokenBridge() {
	const (
		initialReserveImbalance = uint64(1e15)
		valueToSend1            = initialReserveImbalance / 4
		valueToSend2            = initialReserveImbalance
		valueToReturn           = valueToSend1 / 4

		// TODO use a unique deployer for each test file
		deployerKeyStr                     = "ca7269c1fe2a5b86884a119aa516b8d5b641670b83aac0ebf9f2d71affcc12e4"
		ExampleERC20ByteCodeFile           = "./contracts/out/ExampleERC20.sol/ExampleERC20.json"
		ERC20TokenSourceByteCodeFile       = "./contracts/out/ERC20TokenSource.sol/ERC20TokenSource.json"
		NativeTokenDestinationByteCodeFile = "./contracts/out/NativeTokenDestination.sol/NativeTokenDestination.json"
	)
	var (
		ctx                  = context.Background()
		deployerAddress      = common.HexToAddress("0x539447ab8Be7e927bE8E005663C81ff2AE951337")
		tokenReceiverAddress = common.HexToAddress("0x4444444444444444444444444444444444444444")
	)

	subnetA := utils.GetSubnetATestInfo()
	subnetB := utils.GetSubnetBTestInfo()
	teleporterContractAddress := utils.GetTeleporterContractAddress()

	// Info we need to calculate for the test
	deployerPK, err := crypto.HexToECDSA(deployerKeyStr)
	Expect(err).Should(BeNil())
	bridgeContractAddress, err := deploymentUtils.DeriveEVMContractAddress(deployerAddress, 0)
	Expect(err).Should(BeNil())
	log.Info("Native Token Bridge Contract Address: " + bridgeContractAddress.Hex())
	exampleERC20ContractAddress, err := deploymentUtils.DeriveEVMContractAddress(deployerAddress, 1)
	Expect(err).Should(BeNil())
	log.Info("Example ERC20 Contract Address: " + exampleERC20ContractAddress.Hex())

	// Deploy the contracts
	// Both contracts in this test will be deployed to 0xAcB633F5B00099c7ec187eB00156c5cd9D854b5B,
	// though they do not necessarily have to be deployed at the same address, each contract needs
	// to know the address of the other.
	// The nativeTokenDestination contract must be added to "adminAddresses" of "contractNativeMinterConfig"
	// in the genesis file for the subnet. This will allow it to call the native minter precompile.
	erc20TokenSourceAbi, err := erc20tokensource.ERC20TokenSourceMetaData.GetAbi()
	Expect(err).Should(BeNil())
	DeployContract(ctx, ERC20TokenSourceByteCodeFile, deployerPK, subnetA, erc20TokenSourceAbi, teleporterContractAddress, subnetB.BlockchainID, bridgeContractAddress, exampleERC20ContractAddress)

	nativeTokenDestinationAbi, err := nativetokendestination.NativeTokenDestinationMetaData.GetAbi()
	Expect(err).Should(BeNil())
	DeployContract(ctx, NativeTokenDestinationByteCodeFile, deployerPK, subnetB, nativeTokenDestinationAbi, teleporterContractAddress, subnetA.BlockchainID, bridgeContractAddress, new(big.Int).SetUint64(initialReserveImbalance))

	exampleERC20Abi, err := exampleerc20.ExampleERC20MetaData.GetAbi()
	Expect(err).Should(BeNil())
	DeployContract(ctx, ExampleERC20ByteCodeFile, deployerPK, subnetA, exampleERC20Abi)

	log.Info("Finished deploying contracts")

	// Create abi objects to call the contract with
	nativeTokenDestination, err := nativetokendestination.NewNativeTokenDestination(bridgeContractAddress, subnetB.WSClient)
	Expect(err).Should(BeNil())
	erc20TokenSource, err := erc20tokensource.NewERC20TokenSource(bridgeContractAddress, subnetA.WSClient)
	Expect(err).Should(BeNil())
	exampleERC20, err := exampleerc20.NewExampleERC20(exampleERC20ContractAddress, subnetA.WSClient)
	Expect(err).Should(BeNil())

	// Helper function
	sendTokensToSource := func(valueToSend uint64, fromKey *ecdsa.PrivateKey, toAddress common.Address) *types.Receipt {
		transactor, err := bind.NewKeyedTransactorWithChainID(fromKey, subnetB.ChainID)
		Expect(err).Should(BeNil())
		transactor.Value = new(big.Int).SetUint64(valueToSend)

		tx, err := nativeTokenDestination.TransferToSource(transactor, toAddress, nativetokendestination.TeleporterFeeInfo{}, []common.Address{})
		Expect(err).Should(BeNil())
		log.Info("Sent TransferToSource transaction on destination chain", "sourceChainID", subnetA.BlockchainID, "txHash", tx.Hash().Hex())

		receipt := utils.WaitForTransaction(ctx, tx.Hash(), subnetB.WSClient)
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

		utils.RelayMessage(ctx, receipt.BlockHash, receipt.BlockNumber, subnetB, subnetA)
		return receipt
	}

	// Helper function
	sendTokensToDestination := func(valueToSend uint64, fromKey *ecdsa.PrivateKey, toAddress common.Address) *types.Receipt {
		transactor, err := bind.NewKeyedTransactorWithChainID(fromKey, subnetA.ChainID)
		Expect(err).Should(BeNil())

		tx, err := erc20TokenSource.TransferToDestination(transactor, toAddress, new(big.Int).SetUint64(valueToSend), big.NewInt(0), []common.Address{})
		Expect(err).Should(BeNil())
		log.Info("Sent TransferToDestination transaction on source chain", "destinationChainID", subnetB.BlockchainID, "txHash", tx.Hash().Hex())

		receipt := utils.WaitForTransaction(ctx, tx.Hash(), subnetA.WSClient)
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

		utils.RelayMessage(ctx, receipt.BlockHash, receipt.BlockNumber, subnetA, subnetB)
		return receipt
	}

	{ // Give erc20TokenSource allowance to spend all of the deployer's ERC20 Tokens
		bal, err := exampleERC20.BalanceOf(nil, deployerAddress)
		Expect(err).Should(BeNil())

		transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetA.ChainID)
		Expect(err).Should(BeNil())
		tx, err := exampleERC20.Approve(transactor, bridgeContractAddress, bal)
		Expect(err).Should(BeNil())

		receipt := utils.WaitForTransaction(ctx, tx.Hash(), subnetA.WSClient)
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	}

	{ // Transfer some tokens A -> B
		// Check starting balance is 0
		bal, err := exampleERC20.BalanceOf(nil, tokenReceiverAddress)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(BeZero())

		sendTokensToDestination(valueToSend1, deployerPK, tokenReceiverAddress)

		// Check intermediate balance, no tokens should be minted because we haven't collateralized
		bal, err = subnetB.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(BeZero())
	}

	{ // Fail to Transfer tokens B -> A because bridge is not collateralized
		// Check starting balance is 0
		bal, err := exampleERC20.BalanceOf(&bind.CallOpts{Accepted: true}, tokenReceiverAddress)
		Expect(err).Should(BeNil())
		Expect(bal.Cmp(common.Big0)).Should(BeZero())

		transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetB.ChainID)
		Expect(err).Should(BeNil())
		transactor.Value = new(big.Int).SetUint64(valueToSend1)

		// This transfer should revert because the bridge isn't collateralized
		_, err = nativeTokenDestination.TransferToSource(transactor, tokenReceiverAddress, nativetokendestination.TeleporterFeeInfo{}, []common.Address{})
		Expect(err).ShouldNot(BeNil())

		// Check we should fail to send because we're not collateralized
		bal, err = exampleERC20.BalanceOf(&bind.CallOpts{Accepted: true}, tokenReceiverAddress)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(BeZero())
	}

	{ // Transfer more tokens A -> B to collateralize the bridge
		// Check starting balance is 0
		bal, err := subnetB.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(BeZero())

		sendTokensToDestination(valueToSend2, deployerPK, tokenReceiverAddress)

		// We should have minted the excess coins after checking the collateral
		bal, err = subnetB.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(Equal(valueToSend1 + valueToSend2 - initialReserveImbalance))
	}

	{ // Transfer tokens B -> A
		sendTokensToSource(valueToReturn, deployerPK, tokenReceiverAddress)

		// Check we should fail to send because we're not collateralized
		bal, err := exampleERC20.BalanceOf(nil, tokenReceiverAddress)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(Equal(valueToReturn))
	}
}

func DeployContract(ctx context.Context, byteCodeFileName string, deployerPK *ecdsa.PrivateKey, subnetInfo utils.SubnetTestInfo, abi *abi.ABI, constructorArgs ...interface{}) {
	// Deploy an example ERC20 contract to be used as the source token
	byteCode, err := deploymentUtils.ExtractByteCode(byteCodeFileName)
	Expect(err).Should(BeNil())
	Expect(len(byteCode) > 0).Should(BeTrue())
	transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetInfo.ChainID)
	Expect(err).Should(BeNil())
	contractAddress, tx, _, err := bind.DeployContract(transactor, *abi, byteCode, subnetInfo.WSClient, constructorArgs...)
	Expect(err).Should(BeNil())

	// Wait for transaction, then check code was deployed
	utils.WaitForTransaction(ctx, tx.Hash(), subnetInfo.WSClient)
	code, err := subnetInfo.WSClient.CodeAt(ctx, contractAddress, nil)
	Expect(err).Should(BeNil())
	Expect(len(code)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
}
