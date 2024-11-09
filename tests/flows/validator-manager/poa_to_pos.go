package staking

import (
	"context"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	nativetokenstakingmanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/NativeTokenStakingManager"
	poavalidatormanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/PoAValidatorManager"
	iposvalidatormanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/interfaces/IPoSValidatorManager"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/gomega"
)

/*
 * Register a PoA validator manager on a L1 with a proxy. The steps are as follows:
 * - Generate random address to be the owner address
 * - Fund native assets to the owner address
 * - Deploy the PoAValidatorManager contract
 * - Deploy a TransparentUpgradeableProxy contract that points to the PoAValidatorManager
 * - Call initialize on the PoAValidatorManager through the proxy
 * - Initialize and complete PoA validator registration
 *
 * Migrates the proxy to a PoS validator manager. The steps are as follows:
 * - Deploy the PoSValidatorManager contract
 * - Upgrade the TransparentUpgradeableProxy to point to the PoSValidatorManager
 * - Call initialize on the PoSValidatorManager through the proxy
 * - Check that previous validator is still active
 * - Initialize and complete PoS validator registration
 * - Attempt to delist previous PoA validator with wrong owner and check that it fails
 * - Delist the previous PoA validator properly
 * - Delist the PoS validator
 */
func PoAMigrationToPoS(network *localnetwork.LocalNetwork) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := network.GetTwoSubnets()
	_, fundedKey := network.GetFundedAccountInfo()
	pChainInfo := utils.GetPChainInfo(cChainInfo)

	signatureAggregator := utils.NewSignatureAggregator(
		cChainInfo.NodeURIs[0],
		[]ids.ID{
			subnetAInfo.SubnetID,
		},
	)

	// Generate random address to be the owner address
	ownerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	ownerAddress := crypto.PubkeyToAddress(ownerKey.PublicKey)

	// Transfer native assets to the owner account
	ctx := context.Background()
	fundAmount := big.NewInt(1e18) // 10avax
	fundAmount.Mul(fundAmount, big.NewInt(10))
	utils.SendNativeTransfer(
		ctx,
		subnetAInfo,
		fundedKey,
		ownerAddress,
		fundAmount,
	)

	// Deploy PoAValidatorManager contract with a proxy
	nodes, initialValidationIDs, proxyAdmin := network.ConvertSubnet(
		ctx,
		subnetAInfo,
		utils.PoAValidatorManager,
		2,
		ownerKey,
		true,
	)
	proxyAddress := network.GetValidatorManager(subnetAInfo.SubnetID)
	poaValidatorManager, err := poavalidatormanager.NewPoAValidatorManager(proxyAddress, subnetAInfo.RPCClient)
	Expect(err).Should(BeNil())

	//
	// Delist one initial validator
	//
	utils.InitializeAndCompleteEndInitialPoAValidation(
		ctx,
		signatureAggregator,
		ownerKey,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		poaValidatorManager,
		proxyAddress,
		initialValidationIDs[0],
		0,
		nodes[0].Weight,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	// Try to call with invalid owner
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())

	_, err = poaValidatorManager.InitializeValidatorRegistration(
		opts,
		poavalidatormanager.ValidatorRegistrationInput{
			NodeID:             nodes[0].NodeID[:],
			RegistrationExpiry: uint64(time.Now().Add(24 * time.Hour).Unix()),
			BlsPublicKey:       nodes[0].NodePoP.PublicKey[:],
		},
		nodes[0].Weight,
	)
	Expect(err).ShouldNot(BeNil())

	//
	// Re-register the validator as a SoV validator
	//
	expiry := uint64(time.Now().Add(24 * time.Hour).Unix())
	poaValidationID := utils.InitializeAndCompletePoAValidatorRegistration(
		ctx,
		signatureAggregator,
		ownerKey,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		poaValidatorManager,
		proxyAddress,
		expiry,
		nodes[0],
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
	poaValidator, err := poaValidatorManager.GetValidator(&bind.CallOpts{}, poaValidationID)
	Expect(err).Should(BeNil())
	poaNodeID := poaValidator.NodeID

	/*
	 ******************
	 * Migrate PoAValidatorManager to PoSValidatorManager
	 ******************
	 */

	// Deploy PoSValidatorManager contract
	newImplAddress, _ := utils.DeployValidatorManager(
		ctx,
		fundedKey,
		subnetAInfo,
		utils.NativeTokenStakingManager,
	)

	// Upgrade the TransparentUpgradeableProxy contract to use the new logic contract
	opts, err = bind.NewKeyedTransactorWithChainID(ownerKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := proxyAdmin.UpgradeAndCall(opts, proxyAddress, newImplAddress, []byte{})
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx.Hash())

	// Change the proxy contract type to NativeTokenStakingManager and initialize it
	nativeStakingManager, err := nativetokenstakingmanager.NewNativeTokenStakingManager(
		proxyAddress,
		subnetAInfo.RPCClient,
	)
	Expect(err).Should(BeNil())

	utils.AddNativeMinterAdmin(ctx, subnetAInfo, fundedKey, proxyAddress)

	rewardCalculatorAddress, _ := utils.DeployExampleRewardCalculator(
		ctx,
		fundedKey,
		subnetAInfo,
		uint64(10),
	)

	tx, err = nativeStakingManager.Initialize(
		opts,
		nativetokenstakingmanager.PoSValidatorManagerSettings{
			BaseSettings: nativetokenstakingmanager.ValidatorManagerSettings{
				SubnetID:               subnetAInfo.SubnetID,
				ChurnPeriodSeconds:     utils.DefaultChurnPeriodSeconds,
				MaximumChurnPercentage: utils.DefaultMaxChurnPercentage,
			},
			MinimumStakeAmount:       big.NewInt(0).SetUint64(utils.DefaultMinStakeAmount),
			MaximumStakeAmount:       big.NewInt(0).SetUint64(utils.DefaultMaxStakeAmount),
			MinimumStakeDuration:     utils.DefaultMinStakeDurationSeconds,
			MinimumDelegationFeeBips: utils.DefaultMinDelegateFeeBips,
			MaximumStakeMultiplier:   utils.DefaultMaxStakeMultiplier,
			WeightToValueFactor:      big.NewInt(0).SetUint64(utils.DefaultWeightToValueFactor),
			RewardCalculator:         rewardCalculatorAddress,
		},
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(context.Background(), subnetAInfo, tx.Hash())

	// Check that previous validator is still registered
	validationID, err := nativeStakingManager.RegisteredValidators(&bind.CallOpts{}, poaNodeID)
	Expect(err).Should(BeNil())
	Expect(validationID[:]).Should(Equal(poaValidationID[:]))

	//
	// Remove the PoA validator and re-register as a PoS validator
	//
	posStakingManager, err := iposvalidatormanager.NewIPoSValidatorManager(proxyAddress, subnetAInfo.RPCClient)
	Expect(err).Should(BeNil())
	utils.InitializeAndCompleteEndPoSValidation(
		ctx,
		signatureAggregator,
		ownerKey,
		subnetAInfo,
		pChainInfo,
		posStakingManager,
		proxyAddress,
		poaValidationID,
		expiry,
		nodes[0],
		1,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	expiry2 := uint64(time.Now().Add(24 * time.Hour).Unix())
	posValidationID := utils.InitializeAndCompleteNativeValidatorRegistration(
		ctx,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		nativeStakingManager,
		proxyAddress,
		expiry2,
		nodes[0],
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)

	// Delist the PoS validator
	utils.InitializeAndCompleteEndPoSValidation(
		ctx,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		posStakingManager,
		proxyAddress,
		posValidationID,
		expiry2,
		nodes[0],
		1,
		network.GetPChainWallet(),
		network.GetNetworkID(),
	)
}
