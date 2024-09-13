package staking

import (
	"context"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	poavalidatormanager "github.com/ava-labs/teleporter/abi-bindings/go/staking/PoAValidatorManager"
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
	weight := uint64(1)

	{
		// Try to call with invalid owner
		opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
		Expect(err).Should(BeNil())

		nodeID := ids.GenerateTestID()
		blsPublicKey := [bls.PublicKeyLen]byte{}
		_, err = validatorManager.InitializeValidatorRegistration(
			opts,
			poavalidatormanager.ValidatorRegistrationInput{
				NodeID:             nodeID,
				RegistrationExpiry: uint64(time.Now().Add(24 * time.Hour).Unix()),
				BlsPublicKey:       blsPublicKey[:],
			},
			weight,
		)
		Expect(err).ShouldNot(BeNil())

		// Initiate validator registration
		validationID = utils.InitializeAndCompletePoAValidatorRegistration(
			network,
			signatureAggregator,
			ownerKey,
			fundedKey,
			subnetAInfo,
			pChainInfo,
			validatorManager,
			validatorManagerAddress,
			weight,
		)
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

		utils.InitializeAndCompleteEndPoAValidation(
			network,
			signatureAggregator,
			ownerKey,
			fundedKey,
			subnetAInfo,
			pChainInfo,
			validatorManager,
			validatorManagerAddress,
			validationID,
			weight,
			1,
		)
	}
}
