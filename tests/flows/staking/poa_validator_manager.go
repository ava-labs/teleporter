package staking

import (
	"context"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/*
 * Register a PoA validator manager on a L1. The steps are as follows:
 * - Generate random address to be the owner address
 * - Fund native assets to the owner address
 * - Deploy the PoAValidatorManager contract
 * - Attempt to initiate with non owner and check that it fails
 * - Initiate validator registration
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the L1
 * - Verify that the validator is registered in the validator manager contract
 *
 * Delists the validator from the L1. The steps are as follows:
 * - Attempt to initiate with non owner and check that it fails
 * - Initiate validator delisting
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the L1
 * - Verify that the validator is delisted from the validator manager contract
 */
func PoAValidatorManager(network interfaces.LocalNetwork) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := utils.GetTwoSubnets(network)
	_, fundedKey := network.GetFundedAccountInfo()
	pChainInfo := utils.GetPChainInfo(cChainInfo)

	signatureAggregator := utils.NewSignatureAggregator(
		cChainInfo.NodeURIs[0],
		[]ids.ID{
			subnetAInfo.SubnetID,
			ids.Empty, // Primary network subnet ID
		},
	)

	// Generate random address to be the owner address
	ownerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	ownerAddress := crypto.PubkeyToAddress(ownerKey.PublicKey)

	// Transfer native assets to the owner account
	ctx := context.Background()
	fundAmount := big.NewInt(1e18) // 1avax
	utils.SendNativeTransfer(
		ctx,
		subnetAInfo,
		fundedKey,
		ownerAddress,
		fundAmount,
	)

	validatorManagerAddress, validatorManager := utils.DeployAndInitializePoAValidatorManager(
		ctx,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		ownerAddress,
	)

	var validationID ids.ID // To be used in the delisting step
	nodeID := ids.GenerateTestID()
	blsPublicKey := [bls.PublicKeyLen]byte{}
	weight := uint64(1)

	{
		// Try to call with invalid owner
		opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
		Expect(err).Should(BeNil())

		_, err = validatorManager.InitializeValidatorRegistration(
			opts,
			weight,
			nodeID,
			uint64(time.Now().Add(24*time.Hour).Unix()),
			blsPublicKey[:],
		)
		Expect(err).ShouldNot(BeNil())

		// Initiate validator registration
		var receipt *types.Receipt
		receipt, validationID = utils.InitializePoAValidatorRegistration(
			ownerKey,
			subnetAInfo,
			weight,
			nodeID,
			blsPublicKey,
			validatorManager,
		)

		// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
		// (Sending to the P-Chain will be skipped for now)
		signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetAInfo, pChainInfo)

		// Validate the Warp message, (this will be done on the P-Chain in the future)
		utils.ValidateRegisterSubnetValidatorMessage(
			signedWarpMessage,
			nodeID,
			weight,
			subnetAInfo.SubnetID,
			blsPublicKey,
		)

		// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
		registrationSignedMessage := utils.ConstructSubnetValidatorRegistrationMessage(
			validationID,
			true,
			subnetAInfo,
			pChainInfo,
			network,
			signatureAggregator,
		)

		// Deliver the Warp message to the subnet
		receipt = utils.CompletePoAValidatorRegistration(
			fundedKey,
			subnetAInfo,
			validatorManagerAddress,
			registrationSignedMessage,
		)
		// Check that the validator is registered in the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			validatorManager.ParseValidationPeriodRegistered,
		)
		Expect(err).Should(BeNil())
		Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
	}

	//
	// Delist the validator
	//
	{
		// Try with invalid validator owner
		opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
		Expect(err).Should(BeNil())
		_, err = validatorManager.InitializeEndValidation(
			opts,
			validationID,
		)
		Expect(err).ShouldNot(BeNil())

		receipt := utils.InitializeEndPoAValidation(
			ownerKey,
			subnetAInfo,
			validatorManager,
			validationID,
		)
		validatorRemovalEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			validatorManager.ParseValidatorRemovalInitialized,
		)
		Expect(err).Should(BeNil())
		Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(validatorRemovalEvent.Weight.Uint64()).Should(Equal(weight))

		// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
		// (Sending to the P-Chain will be skipped for now)
		signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetAInfo, pChainInfo)
		Expect(err).Should(BeNil())

		// Validate the Warp message, (this will be done on the P-Chain in the future)
		utils.ValidateSetSubnetValidatorWeightMessage(signedWarpMessage, validationID, 0, 0)

		// Construct a SubnetValidatorRegistrationMessage Warp message from the P-Chain
		registrationSignedMessage := utils.ConstructSubnetValidatorRegistrationMessage(
			validationID,
			false,
			subnetAInfo,
			pChainInfo,
			network,
			signatureAggregator,
		)

		// Deliver the Warp message to the subnet
		receipt = utils.CompleteEndPoAValidation(
			ownerKey,
			subnetAInfo,
			validatorManagerAddress,
			registrationSignedMessage,
		)

		// Check that the validator is has been delisted from the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			validatorManager.ParseValidationPeriodEnded,
		)
		Expect(err).Should(BeNil())
		Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
	}
}
