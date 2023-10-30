package tests

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	nativetokendestination "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/NativeTokenBridge/NativeTokenDestination"
	nativetokensource "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/NativeTokenBridge/NativeTokenSource"
	"github.com/ava-labs/teleporter/tests/utils"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

func NativeTokenBridge() {
	const (
		tokenReserve  = uint64(1e15)
		valueToSend1  = tokenReserve / 4
		valueToSend2  = tokenReserve
		valueToReturn = valueToSend1 / 4

		deployerKeyStr                     = "aad7440febfc8f9d73a58c3cb1f1754779a566978f9ebffcd4f4698e9b043985"
		NativeTokenSourceByteCodeFile      = "./contracts/out/NativeTokenSource.sol/NativeTokenSource.json"
		NativeTokenDestinationByteCodeFile = "./contracts/out/NativeTokenDestination.sol/NativeTokenDestination.json"
	)
	var (
		ctx                  = context.Background()
		deployerAddress      = common.HexToAddress("0x1337cfd2dCff6270615B90938aCB1efE79801704")
		tokenReceiverAddress = common.HexToAddress("0x0123456789012345678901234567890123456789")
	)

	subnetA := utils.GetSubnetATestInfo()
	subnetB := utils.GetSubnetBTestInfo()
	teleporterContractAddress := utils.GetTeleporterContractAddress()

	// Info we need to calculate for the test
	deployerPK, err := crypto.HexToECDSA(deployerKeyStr)
	Expect(err).Should(BeNil())
	nativeTokenBridgeContractAddress, err := deploymentUtils.DeriveEVMContractAddress(deployerAddress, 0)
	Expect(err).Should(BeNil())
	log.Info("Native Token Bridge Contract Address: " + nativeTokenBridgeContractAddress.Hex())

	{ // Deploy the contracts
		nativeTokenSourceBytecode, err := deploymentUtils.ExtractByteCode(NativeTokenSourceByteCodeFile)
		Expect(err).Should(BeNil())
		chainATransactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetA.ChainID)
		Expect(err).Should(BeNil())
		nativeTokenSourceAbi, err := nativetokensource.NativeTokenSourceMetaData.GetAbi()
		Expect(err).Should(BeNil())
		_, txA, _, err := bind.DeployContract(chainATransactor, *nativeTokenSourceAbi, nativeTokenSourceBytecode, subnetA.WSClient, teleporterContractAddress, subnetB.BlockchainID, nativeTokenBridgeContractAddress)
		Expect(err).Should(BeNil())

		// Both contracts in this test will be deployed to 0xAcB633F5B00099c7ec187eB00156c5cd9D854b5B,
		// though they do not necessarily have to be deployed at the same address, each contract needs
		// to know the address of the other.
		// The nativeTokenDestination contract must be added to "adminAddresses" of "contractNativeMinterConfig"
		// in the genesis file for the subnet. This will allow it to call the native minter precompile.
		nativeTokenDestinationBytecode, err := deploymentUtils.ExtractByteCode(NativeTokenDestinationByteCodeFile)
		Expect(err).Should(BeNil())
		chainBTransactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetB.ChainID)
		Expect(err).Should(BeNil())
		nativeTokenDestinationAbi, err := nativetokendestination.NativeTokenDestinationMetaData.GetAbi()
		Expect(err).Should(BeNil())
		_, txB, _, err := bind.DeployContract(chainBTransactor, *nativeTokenDestinationAbi, nativeTokenDestinationBytecode, subnetB.WSClient, teleporterContractAddress, subnetA.BlockchainID, nativeTokenBridgeContractAddress, new(big.Int).SetUint64(tokenReserve))
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

		log.Info("Finished deploying Bridge contracts")
	}

	// Create abi objects to call the contract with
	nativeTokenDestination, err := nativetokendestination.NewNativeTokenDestination(nativeTokenBridgeContractAddress, subnetB.WSClient)
	Expect(err).Should(BeNil())
	nativeTokenSource, err := nativetokensource.NewNativeTokenSource(nativeTokenBridgeContractAddress, subnetA.WSClient)
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
		transactor.Value = new(big.Int).SetUint64(valueToSend)

		tx, err := nativeTokenSource.TransferToDestination(transactor, toAddress, common.Address{}, big.NewInt(0), []common.Address{})
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
		bal, err := subnetB.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(Equal(common.Big0.Uint64()))

		sendTokensToDestination(valueToSend1, deployerPK, tokenReceiverAddress)

		// Check intermediate balance, no tokens should be minted because we haven't collateralized
		bal, err = subnetB.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(BeZero())
	}

	{ // Fail to Transfer tokens B -> A because bridge is not collateralized
		// Check starting balance is 0
		bal, err := subnetA.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(BeZero())

		transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetB.ChainID)
		Expect(err).Should(BeNil())
		transactor.Value = new(big.Int).SetUint64(valueToSend1)

		// This transfer should revert because the bridge isn't collateralized
		_, err = nativeTokenDestination.TransferToSource(transactor, tokenReceiverAddress, common.Address{}, big.NewInt(0), []common.Address{})
		Expect(err).ShouldNot(BeNil())

		// Check we should fail to send because we're not collateralized
		bal, err = subnetA.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
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
		Expect(bal.Uint64()).Should(Equal(valueToSend1 + valueToSend2 - tokenReserve))
	}

	{ // Transfer tokens B -> A
		sendTokensToSource(valueToReturn, deployerPK, tokenReceiverAddress)

		// Check we should fail to send because we're not collateralized
		bal, err := subnetA.WSClient.BalanceAt(ctx, tokenReceiverAddress, nil)
		Expect(err).Should(BeNil())
		Expect(bal.Uint64()).Should(Equal(valueToReturn))
	}
}
