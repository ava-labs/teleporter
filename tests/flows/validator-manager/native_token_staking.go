package staking

import (
	"context"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
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
	ctx := context.Background()

	// Deploy the staking manager contract
	stakingManagerContractAddress, stakingManager := utils.DeployAndInitializeNativeTokenStakingManager(
		ctx,
		fundedKey,
		subnetAInfo,
		pChainInfo,
	)

	utils.AddNativeMinterAdmin(ctx, subnetAInfo, fundedKey, stakingManagerContractAddress)

	_ = utils.InitializeNativeTokenValidatorSet(
		ctx,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		stakingManager,
		stakingManagerContractAddress,
		network,
		signatureAggregator,
		utils.DefaultMinStakeAmount*10,
	)

	//
	// Register a validator
	//
	stakeAmount := new(big.Int).SetUint64(utils.DefaultMinStakeAmount)
	weight, err := stakingManager.ValueToWeight(
		&bind.CallOpts{},
		stakeAmount,
	)
	Expect(err).Should(BeNil())
	// Iniatiate validator registration
	validationID := utils.InitializeAndCompleteNativeValidatorRegistration(
		ctx,
		network,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		stakingManager,
		stakingManagerContractAddress,
		stakeAmount,
	)

	//
	// Delist the validator
	//
	utils.InitializeAndCompleteEndNativeValidation(
		ctx,
		network,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		stakingManager,
		stakingManagerContractAddress,
		validationID,
		weight,
		1,
	)
}
