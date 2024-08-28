package staking

import (
	"context"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

/*
 * Registers a native token staking validator on a subnet. The steps are as follows:
 * - Deploy the NativeTokenStakingManager
 * - Initiate validator registration
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the subnet
 * - Verify that the validator is registered in the staking contract
 *
 * Delists the validator from the subnet. The steps are as follows:
 * - Initiate validator delisting
 * - Deliver the Warp message to the P-Chain (not implemented)
 * - Aggregate P-Chain signatures on the response Warp message
 * - Deliver the Warp message to the subnet
 * - Verify that the validator is delisted from the staking contract
 */
func NativeTokenStakingManager(network interfaces.LocalNetwork) {
	// Get the subnets info
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

	// Deploy the staking manager contract
	stakingManagerContractAddress, stakingManager := utils.DeployAndInitializeNativeTokenStakingManager(
		context.Background(),
		fundedKey,
		subnetAInfo,
		pChainInfo,
	)

	//
	// Register a validator
	//
	var validationID ids.ID // To be used in the delisting step
	stakeAmount := big.NewInt(1e18)
	weight, err := stakingManager.ValueToWeight(
		&bind.CallOpts{},
		stakeAmount,
	)
	Expect(err).Should(BeNil())
	{
		// Iniatiate validator registration
		nodeID := ids.GenerateTestID()
		blsPublicKey := [bls.PublicKeyLen]byte{}
		var receipt *types.Receipt
		receipt, validationID = utils.InitializeNativeValidatorRegistration(
			fundedKey,
			subnetAInfo,
			stakeAmount,
			nodeID,
			blsPublicKey,
			stakingManager,
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
		receipt = utils.CompleteNativeValidatorRegistration(
			fundedKey,
			subnetAInfo,
			stakingManagerContractAddress,
			registrationSignedMessage,
		)
		// Check that the validator is registered in the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			stakingManager.ParseValidationPeriodRegistered,
		)
		Expect(err).Should(BeNil())
		Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
	}

	//
	// Delist the validator
	//
	{
		receipt := utils.InitializeEndNativeValidation(
			fundedKey,
			subnetAInfo,
			stakingManager,
			validationID,
		)
		validatorRemovalEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			stakingManager.ParseValidatorRemovalInitialized,
		)
		Expect(err).Should(BeNil())
		Expect(validatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(validatorRemovalEvent.Weight.Uint64()).Should(Equal(weight))

		// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
		// (Sending to the P-Chain will be skipped for now)
		signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetAInfo, pChainInfo)
		Expect(err).Should(BeNil())

		// Validate the Warp message, (this will be done on the P-Chain in the future)
		utils.ValidateSetSubnetValidatorWeightMessage(signedWarpMessage, validationID, 0, 1)

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
		receipt = utils.CompleteEndNativeValidation(
			fundedKey,
			subnetAInfo,
			stakingManagerContractAddress,
			registrationSignedMessage,
		)

		// Check that the validator is has been delisted from the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			stakingManager.ParseValidationPeriodEnded,
		)
		Expect(err).Should(BeNil())
		Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
	}
}
