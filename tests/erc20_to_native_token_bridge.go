package tests

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	erc20tokensource "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/NativeTokenBridge/ERC20TokenSource"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/Mocks/ExampleERC20"
	nativetokendestination "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/NativeTokenBridge/NativeTokenDestination"
	"github.com/ava-labs/teleporter/tests/utils"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func ERC20ToNativeTokenBridge() {
	const (
		tokenReserve  = uint64(1e15)
		valueToSend1  = tokenReserve / 4
		valueToSend2  = tokenReserve
		valueToReturn = valueToSend1 / 4

		bridgeDeployerKeyStr               = "aad7440febfc8f9d73a58c3cb1f1754779a566978f9ebffcd4f4698e9b043985"
		ExampleERC20ByteCodeFile           = "./contracts/out/ExampleERC20.sol/ExampleERC20.json"
		ERC20TokenSourceByteCodeFile       = "./contracts/out/ERC20TokenSource.sol/ERC20TokenSource.json"
		NativeTokenDestinationByteCodeFile = "./contracts/out/NativeTokenDestination.sol/NativeTokenDestination.json"
	)
	var (
		ctx                       = context.Background()
		nativeTokenBridgeDeployer = common.HexToAddress("0x1337cfd2dCff6270615B90938aCB1efE79801704")
		tokenReceiverAddress      = common.HexToAddress("0x0123456789012345678901234567890123456789")
	)

	subnetA := utils.GetSubnetATestInfo()
	subnetB := utils.GetSubnetBTestInfo()
	teleporterContractAddress := utils.GetTeleporterContractAddress()

	// Info we need to calculate for the test
	nativeTokenBridgeDeployerPK, err := crypto.HexToECDSA(bridgeDeployerKeyStr)
	Expect(err).Should(BeNil())
	nativeTokenBridgeContractAddress, err := deploymentUtils.DeriveEVMContractAddress(nativeTokenBridgeDeployer, 0)
	Expect(err).Should(BeNil())
	log.Info("Native Token Bridge Contract Address: " + nativeTokenBridgeContractAddress.Hex())
	var	exampleERC20ContractAddress common.Address

	{ // Deploy the contracts
		erc20TokenSourceBytecode, err := deploymentUtils.ExtractByteCode(ERC20TokenSourceByteCodeFile)
		Expect(err).Should(BeNil())
		chainATransactor, err := bind.NewKeyedTransactorWithChainID(nativeTokenBridgeDeployerPK, subnetA.ChainID)
		Expect(err).Should(BeNil())
		erc20TokenSourceAbi, err := erc20tokensource.ERC20TokenSourceMetaData.GetAbi()
		Expect(err).Should(BeNil())
		_, txA, _, err := bind.DeployContract(chainATransactor, *erc20TokenSourceAbi, erc20TokenSourceBytecode, subnetA.WSClient, teleporterContractAddress, subnetB.BlockchainID, nativeTokenBridgeContractAddress)
		Expect(err).Should(BeNil())

		// Both contracts in this test will be deployed to 0xAcB633F5B00099c7ec187eB00156c5cd9D854b5B,
		// though they do not necessarily have to be deployed at the same address, each contract needs
		// to know the address of the other.
		// The nativeTokenDestination contract must be added to "adminAddresses" of "contractNativeMinterConfig"
		// in the genesis file for the subnet. This will allow it to call the native minter precompile.
		nativeTokenDestinationBytecode, err := deploymentUtils.ExtractByteCode(NativeTokenDestinationByteCodeFile)
		Expect(err).Should(BeNil())
		chainBTransactor, err := bind.NewKeyedTransactorWithChainID(nativeTokenBridgeDeployerPK, subnetB.ChainID)
		Expect(err).Should(BeNil())
		nativeTokenDestinationAbi, err := nativetokendestination.NativeTokenDestinationMetaData.GetAbi()
		Expect(err).Should(BeNil())
		_, txB, _, err := bind.DeployContract(chainBTransactor, *nativeTokenDestinationAbi, nativeTokenDestinationBytecode, subnetB.WSClient, teleporterContractAddress, subnetA.BlockchainID, nativeTokenBridgeContractAddress, new(big.Int).SetUint64(tokenReserve))
		Expect(err).Should(BeNil())

		// Deploy an example ERC20 contract to be used as the source token
		exampleERC20Bytecode, err := deploymentUtils.ExtractByteCode(ExampleERC20ByteCodeFile)
		Expect(err).Should(BeNil())
		exampleERC20Abi, err := exampleerc20.ExampleERC20MetaData.GetAbi()
		Expect(err).Should(BeNil())
		exampleERC20ContractAddress, txExampleERC20, _, err := bind.DeployContract(chainATransactor, *exampleERC20Abi, exampleERC20Bytecode, subnetA.WSClient, "Ignored", "Params")
		Expect(err).Should(BeNil())

		// Wait for transaction, then check code was deployed
		utils.WaitForTransaction(ctx, txA.Hash(), subnetA.WSClient)
		bridgeCodeA, err := subnetA.WSClient.CodeAt(ctx, nativeTokenBridgeContractAddress, nil)
		Expect(err).Should(BeNil())
		Expect(len(bridgeCodeA)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode

		// Wait for transaction, then check code was deployed
		utils.WaitForTransaction(ctx, txB.Hash(), subnetB.WSClient)
		bridgeCodeB, err := subnetB.WSClient.CodeAt(ctx, nativeTokenBridgeContractAddress, nil)
		Expect(err).Should(BeNil())
		Expect(len(bridgeCodeB)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode

		// Wait for transaction, then check code was deployed
		utils.WaitForTransaction(ctx, txExampleERC20.Hash(), subnetA.WSClient)
		exampleERC20Code, err := subnetB.WSClient.CodeAt(ctx, exampleERC20ContractAddress, nil)
		Expect(err).Should(BeNil())
		Expect(len(exampleERC20Code)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode

		log.Info("Finished deploying contracts")
	}

	// Create abi objects to call the contract with
	nativeTokenDestination, err := nativetokendestination.NewNativeTokenDestination(nativeTokenBridgeContractAddress, subnetB.WSClient)
	Expect(err).Should(BeNil())
	erc20TokenSource, err := erc20tokensource.NewERC20TokenSource(nativeTokenBridgeContractAddress, subnetA.WSClient)
	Expect(err).Should(BeNil())
	exampleERC20, err := exampleerc20.NewExampleERC20(exampleERC20ContractAddress, subnetA.WSClient)
	Expect(err).Should(BeNil())

	// Helper function
	sendTokensToSource := func(valueToSend uint64, fromKey *ecdsa.PrivateKey, toAddress common.Address) *types.Receipt {
		transactor, err := bind.NewKeyedTransactorWithChainID(fromKey, subnetB.ChainID)
		Expect(err).Should(BeNil())
		transactor.Value = new(big.Int).SetUint64(valueToSend)

		tx, err := nativeTokenDestination.TransferToSource(transactor, toAddress, common.Address{}, big.NewInt(0), []common.Address{})
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

		tx, err := erc20TokenSource.TransferToDestination(transactor, toAddress, exampleERC20ContractAddress, new(big.Int).SetUint64(valueToSend), big.NewInt(0), []common.Address{})
		Expect(err).Should(BeNil())
		log.Info("Sent TransferToDestination transaction on source chain", "destinationChainID", subnetB.BlockchainID, "txHash", tx.Hash().Hex())

		receipt := utils.WaitForTransaction(ctx, tx.Hash(), subnetA.WSClient)
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

		utils.RelayMessage(ctx, receipt.BlockHash, receipt.BlockNumber, subnetA, subnetB)
		return receipt
	}

	{ // Transfer some tokens A -> B
		// Check starting balance is 0
		bal, err := exampleERC20.BalanceOf(&bind.CallOpts{ Accepted: true }, tokenReceiverAddress)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(BeZero())

		sendTokensToDestination(valueToSend1, nativeTokenBridgeDeployerPK, tokenReceiverAddress)

		// Check intermediate balance, no tokens should be minted because we haven't collateralized
		bal, err = subnetB.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(BeZero())
	}

	{ // Fail to Transfer tokens B -> A because bridge is not collateralized
		// Check starting balance is 0
		bal, err := exampleERC20.BalanceOf(&bind.CallOpts{ Accepted: true }, tokenReceiverAddress)
		Expect(err).Should(BeNil())
		Expect(bal.Cmp(common.Big0)).Should(BeZero())

		transactor, err := bind.NewKeyedTransactorWithChainID(nativeTokenBridgeDeployerPK, subnetB.ChainID)
		Expect(err).Should(BeNil())
		transactor.Value = new(big.Int).SetUint64(valueToSend1)

		// This transfer should revert because the bridge isn't collateralized
		_, err = nativeTokenDestination.TransferToSource(transactor, tokenReceiverAddress, common.Address{}, big.NewInt(0), []common.Address{})
		Expect(err).ShouldNot(BeNil())

		// Check we should fail to send because we're not collateralized
		bal, err = exampleERC20.BalanceOf(&bind.CallOpts{ Accepted: true }, tokenReceiverAddress)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(BeZero())
	}

	{ // Transfer more tokens A -> B to collateralize the bridge
		// Check starting balance is 0
		bal, err := subnetB.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(BeZero())

		sendTokensToDestination(valueToSend2, nativeTokenBridgeDeployerPK, tokenReceiverAddress)

		// We should have minted the excess coins after checking the collateral
		bal, err = subnetB.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(Equal(valueToSend1 + valueToSend2 - tokenReserve))
	}

	{ // Transfer tokens B -> A
		sendTokensToSource(valueToReturn, nativeTokenBridgeDeployerPK, tokenReceiverAddress)

		// Check we should fail to send because we're not collateralized
		bal, err := exampleERC20.BalanceOf(&bind.CallOpts{ Accepted: true }, tokenReceiverAddress)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(Equal(valueToReturn))
	}
}
