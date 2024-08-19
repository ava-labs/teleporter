package flows

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	warpMessages "github.com/ava-labs/avalanchego/vms/platformvm/warp/messages"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

// Registers a erc20 token staking validator on a subnet. The steps are as follows:
// - Deploy the ERCTokenStakingManager
// - Initiate validator registration
// - Deliver the Warp message to the P-Chain (not implemented)
// - Aggregate P-Chain signatures on the response Warp message
// - Deliver the Warp message to the subnet
// - Verify that the validator is registered in the staking contract
// Delists the validator from the subnet. The steps are as follows:
// - Initiate validator delisting
// - Deliver the Warp message to the P-Chain (not implemented)
// - Aggregate P-Chain signatures on the response Warp message
// - Deliver the Warp message to the subnet
// - Verify that the validator is delisted from the staking contract
func ERC20TokenStakingManager(network interfaces.LocalNetwork) {
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
	stakingManagerAddress, stakingManager, _, erc20 := utils.DeployAndInitializeERC20TokenStakingManager(
		context.Background(),
		fundedKey,
		subnetAInfo,
		pChainInfo,
	)

	//
	// Register a validator
	//
	var validationID ids.ID // To be used in the delisting step
	stakeAmount := uint64(1e18)
	weight := uint64(1e6) // stakeAmount / 1e12
	{
		// Iniatiate validator registration
		nodeID := ids.GenerateTestID()
		blsPublicKey := [bls.PublicKeyLen]byte{}
		var receipt *types.Receipt
		receipt, validationID = utils.CallERC20InitializeValidatorRegistration(
			fundedKey,
			subnetAInfo,
			stakeAmount,
			erc20,
			stakingManagerAddress,
			nodeID,
			blsPublicKey,
			stakingManager,
		)

		// Gather subnet-evm Warp signatures for the RegisterSubnetValidatorMessage & relay to the P-Chain
		// (Sending to the P-Chain will be skipped for now)
		signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetAInfo, pChainInfo)

		// Validate the Warp message, (this will be done on the P-Chain in the future)
		msg, err := warpPayload.ParseAddressedCall(signedWarpMessage.UnsignedMessage.Payload)
		Expect(err).Should(BeNil())
		// Check that the addressed call payload is a registered Warp message type
		var payloadInterface warpMessages.Payload
		ver, err := warpMessages.Codec.Unmarshal(msg.Payload, &payloadInterface)
		Expect(err).Should(BeNil())
		registerValidatorPayload, ok := payloadInterface.(*warpMessages.RegisterSubnetValidator)
		Expect(ok).Should(BeTrue())

		Expect(ver).Should(Equal(uint16(warpMessages.CodecVersion)))
		Expect(registerValidatorPayload.NodeID).Should(Equal(nodeID))
		Expect(registerValidatorPayload.Weight).Should(Equal(weight))
		Expect(registerValidatorPayload.SubnetID).Should(Equal(subnetAInfo.SubnetID))
		Expect(registerValidatorPayload.BlsPubKey[:]).Should(Equal(blsPublicKey[:]))

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
		receipt = utils.CallERC20CompleteValidatorRegistration(
			fundedKey,
			subnetAInfo,
			stakingManagerAddress,
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
		receipt := utils.CallERC20InitializeEndValidation(
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
		Expect(validatorRemovalEvent.StakeAmount.Uint64()).Should(Equal(weight))

		// Gather subnet-evm Warp signatures for the SetSubnetValidatorWeightMessage & relay to the P-Chain
		// (Sending to the P-Chain will be skipped for now)
		signedWarpMessage := network.ConstructSignedWarpMessage(context.Background(), receipt, subnetAInfo, pChainInfo)
		Expect(err).Should(BeNil())
		// Validate the Warp message, (this will be done on the P-Chain in the future)
		msg, err := warpPayload.ParseAddressedCall(signedWarpMessage.UnsignedMessage.Payload)
		Expect(err).Should(BeNil())
		// Check that the addressed call payload is a registered Warp message type
		var payloadInterface warpMessages.Payload
		ver, err := warpMessages.Codec.Unmarshal(msg.Payload, &payloadInterface)
		Expect(err).Should(BeNil())
		registerValidatorPayload, ok := payloadInterface.(*warpMessages.SetSubnetValidatorWeight)
		Expect(ok).Should(BeTrue())

		Expect(ver).Should(Equal(uint16(warpMessages.CodecVersion)))
		Expect(registerValidatorPayload.ValidationID).Should(Equal(validationID))
		Expect(registerValidatorPayload.Weight).Should(Equal(uint64(0)))
		Expect(registerValidatorPayload.Nonce).Should(Equal(uint64(0)))

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
		receipt = utils.CallERC20CompleteEndValidation(
			fundedKey,
			subnetAInfo,
			stakingManagerAddress,
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
