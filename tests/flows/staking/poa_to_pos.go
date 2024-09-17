package staking

import (
	"context"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	nativetokenstakingmanager "github.com/ava-labs/teleporter/abi-bindings/go/staking/NativeTokenStakingManager"
	poavalidatormanager "github.com/ava-labs/teleporter/abi-bindings/go/staking/PoAValidatorManager"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
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
func PoAMigrationToPoS(network interfaces.LocalNetwork) {
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

	// Deploy PoAValidatorManager contract
	implAddress, _ := utils.DeployPoAValidatorManager(
		ctx,
		fundedKey,
		subnetAInfo,
	)

	// Deploy TransparentUpgradeableProxy contract pointing to PoAValidatorManager
	proxyAddress, proxyAdmin, poaValidatorManager := utils.DeployTransparentUpgradeableProxy(
		ctx,
		subnetAInfo,
		fundedKey,
		implAddress,
		poavalidatormanager.NewPoAValidatorManager,
	)
	opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey,
		subnetAInfo.EVMChainID,
	)
	Expect(err).Should(BeNil())

	tx, err := poaValidatorManager.Initialize(
		opts,
		poavalidatormanager.ValidatorManagerSettings{
			PChainBlockchainID:     pChainInfo.BlockchainID,
			SubnetID:               subnetAInfo.SubnetID,
			ChurnPeriodSeconds:     uint64(0),
			MaximumChurnPercentage: uint8(20),
		},
		ownerAddress,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(context.Background(), subnetAInfo, tx.Hash())

	_ = utils.InitializePoAValidatorSet(
		ctx,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		poaValidatorManager,
		proxyAddress,
		network,
		signatureAggregator,
		20,
	)

	// Register a validator
	poaWeight := uint64(1)
	poaValidationID := utils.InitializeAndCompletePoAValidatorRegistration(
		network,
		signatureAggregator,
		ownerKey,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		poaValidatorManager,
		proxyAddress,
		poaWeight,
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
	newImplAddress, _ := utils.DeployNativeTokenStakingManager(
		ctx,
		fundedKey,
		subnetAInfo,
	)

	// Upgrade the TransparentUpgradeableProxy contract to use the new logic contract
	tx, err = proxyAdmin.UpgradeAndCall(opts, proxyAddress, newImplAddress, []byte{})
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx.Hash())

	// Change the proxy contract type to NativeTokenStakingManager and initialize it
	posValidatorManager, err := nativetokenstakingmanager.NewNativeTokenStakingManager(
		proxyAddress,
		subnetAInfo.RPCClient,
	)
	Expect(err).Should(BeNil())

	tx, err = posValidatorManager.Initialize(
		opts,
		nativetokenstakingmanager.PoSValidatorManagerSettings{
			BaseSettings: nativetokenstakingmanager.ValidatorManagerSettings{
				PChainBlockchainID:     pChainInfo.BlockchainID,
				SubnetID:               subnetAInfo.SubnetID,
				ChurnPeriodSeconds:     utils.DefaultChurnPeriodSeconds,
				MaximumChurnPercentage: utils.DefaultMaxChurnPercentage,
			},
			MinimumStakeAmount:       big.NewInt(0).SetUint64(utils.DefaultMinStakeAmount),
			MaximumStakeAmount:       big.NewInt(0).SetUint64(utils.DefaultMaxStakeAmount),
			MinimumStakeDuration:     utils.DefaultMinStakeDurationSeconds,
			MinimumDelegationFeeBips: utils.DefaultMinDelegateFeeBips,
			MaximumStakeMultiplier:   utils.DefaultMaxStakeMultiplier,
			RewardCalculator:         common.Address{},
		},
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(context.Background(), subnetAInfo, tx.Hash())

	// Check that previous validator is still active
	validationID, err := posValidatorManager.ActiveValidators(&bind.CallOpts{}, poaNodeID)
	Expect(err).Should(BeNil())
	Expect(validationID[:]).Should(Equal(poaValidationID[:]))

	// Register a PoS validator
	stakeAmount := big.NewInt(1e18)
	posWeight, err := posValidatorManager.ValueToWeight(
		&bind.CallOpts{},
		stakeAmount,
	)
	Expect(err).Should(BeNil())

	posValidationID := utils.InitializeAndCompleteNativeValidatorRegistration(
		network,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		posValidatorManager,
		proxyAddress,
		stakeAmount,
	)

	// Attempt to delist previous PoA validator with wrong owner and check that it fails
	_, err = posValidatorManager.InitializeEndValidation(
		opts,
		poaValidationID,
		false,
		0,
	)
	Expect(err).ShouldNot(BeNil())

	// Delist the previous PoA validator properly
	utils.InitializeAndCompleteEndNativeValidation(
		network,
		signatureAggregator,
		ownerKey,
		subnetAInfo,
		pChainInfo,
		posValidatorManager,
		proxyAddress,
		poaValidationID,
		poaWeight,
		1,
	)

	// Delist the PoS validator
	utils.InitializeAndCompleteEndNativeValidation(
		network,
		signatureAggregator,
		fundedKey,
		subnetAInfo,
		pChainInfo,
		posValidatorManager,
		proxyAddress,
		posValidationID,
		posWeight,
		1,
	)
}
