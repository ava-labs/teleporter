package staking

import (
	"context"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

/*
 * Registers a delegator with a native token staking validator on a subnet. The steps are as follows:
 * - Deploy the NativeTokenStakingManager
 * - Register a validator
 * - Register a delegator
 * - Deleist the delegator
 *
 * TODO: as delegation gets built out, this test will need to be updated to cover:
 * - Delegator rewards
 * - Implicit delegation end at validation end
 */
func NativeDelegation(network interfaces.LocalNetwork) {
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
	stakingManagerAddress, stakingManager := utils.DeployAndInitializeNativeTokenStakingManager(
		context.Background(),
		fundedKey,
		subnetAInfo,
		pChainInfo,
	)

	//
	// Register a validator
	//
	var validationID ids.ID // To be used in the delisting step
	validatorStake := new(big.Int).SetUint64(utils.DefaultMinStakeAmount)
	validatorWeight, err := stakingManager.ValueToWeight(
		&bind.CallOpts{},
		validatorStake,
	)
	Expect(err).Should(BeNil())
	nodeID := ids.GenerateTestID()
	blsPublicKey := [bls.PublicKeyLen]byte{}
	stakeAmount := new(big.Int).SetUint64(utils.DefaultMinStakeAmount)
	validationID = utils.InitializeAndCompleteNativeValidatorRegistration(
		network,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		stakingManager,
		stakingManagerAddress,
		nodeID,
		blsPublicKey,
		stakeAmount,
	)
	//
	// Register a delegator
	//
	var delegationID ids.ID
	{
		delegatorStake := big.NewInt(1e17)
		delegatorWeight, err := stakingManager.ValueToWeight(
			&bind.CallOpts{},
			delegatorStake,
		)
		Expect(err).Should(BeNil())
		newValidatorWeight := validatorWeight + delegatorWeight

		nonce := uint64(1)

		receipt := utils.InitializeNativeDelegatorRegistration(
			fundedKey,
			subnetAInfo,
			validationID,
			delegatorStake,
			stakingManagerAddress,
			stakingManager,
		)
		initRegistrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			stakingManager.ParseDelegatorAdded,
		)
		Expect(err).Should(BeNil())
		delegationID = initRegistrationEvent.DelegationID

		// Gather subnet-evm Warp signatures for the SubnetValidatorWeightUpdateMessage & relay to the P-Chain
		// (Sending to the P-Chain will be skipped for now)
		signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetAInfo, pChainInfo)

		// Validate the Warp message, (this will be done on the P-Chain in the future)
		utils.ValidateSetSubnetValidatorWeightMessage(
			signedWarpMessage,
			validationID,
			newValidatorWeight,
			nonce,
		)

		// Construct a SubnetValidatorWeightUpdateMessage Warp message from the P-Chain
		registrationSignedMessage := utils.ConstructSubnetValidatorWeightUpdateMessage(
			validationID,
			nonce,
			newValidatorWeight,
			subnetAInfo,
			pChainInfo,
			network,
			signatureAggregator,
		)

		// Deliver the Warp message to the subnet
		receipt = utils.CompleteNativeDelegatorRegistration(
			fundedKey,
			delegationID,
			subnetAInfo,
			stakingManagerAddress,
			registrationSignedMessage,
		)
		// Check that the validator is registered in the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			stakingManager.ParseDelegatorRegistered,
		)
		Expect(err).Should(BeNil())
		Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(registrationEvent.DelegationID[:]).Should(Equal(delegationID[:]))
	}
	//
	// Delist the delegator
	//
	{
		nonce := uint64(2)
		receipt := utils.InitializeEndNativeDelegation(
			fundedKey,
			subnetAInfo,
			stakingManager,
			delegationID,
		)
		delegatorRemovalEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			stakingManager.ParseDelegatorRemovalInitialized,
		)
		Expect(err).Should(BeNil())
		Expect(delegatorRemovalEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(delegatorRemovalEvent.DelegationID[:]).Should(Equal(delegationID[:]))

		// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
		// (Sending to the P-Chain will be skipped for now)
		signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetAInfo, pChainInfo)
		Expect(err).Should(BeNil())

		// Validate the Warp message, (this will be done on the P-Chain in the future)
		utils.ValidateSetSubnetValidatorWeightMessage(signedWarpMessage, validationID, validatorWeight, 2)

		// Construct a SubnetValidatorWeightUpdateMessage Warp message from the P-Chain
		signedMessage := utils.ConstructSubnetValidatorWeightUpdateMessage(
			validationID,
			nonce,
			validatorWeight,
			subnetAInfo,
			pChainInfo,
			network,
			signatureAggregator,
		)

		// Deliver the Warp message to the subnet
		receipt = utils.CompleteEndNativeDelegation(
			fundedKey,
			delegationID,
			subnetAInfo,
			stakingManagerAddress,
			signedMessage,
		)

		// Check that the delegator has been delisted from the staking contract
		registrationEvent, err := utils.GetEventFromLogs(
			receipt.Logs,
			stakingManager.ParseDelegationEnded,
		)
		Expect(err).Should(BeNil())
		Expect(registrationEvent.ValidationID[:]).Should(Equal(validationID[:]))
		Expect(registrationEvent.DelegationID[:]).Should(Equal(delegationID[:]))
	}
}
