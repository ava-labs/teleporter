package flows

import (
	"context"
	"math/big"

	runner_sdk "github.com/ava-labs/avalanche-network-runner/client"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/Mocks/ExampleERC20"
	validatorsetsig "github.com/ava-labs/teleporter/abi-bindings/go/Utilities/ValidatorSetSig"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func ValidatorSetSig(network interfaces.LocalNetwork) {
	// Deploy ValidatorSetSig contract to contract subnet to be validated from validator subnet
	// Deploy a mock ERC20 contract to contract subnet
	// Construct ValidatorSetSig message with mock ERC20 as the target contract
	// Create a OffChainWarp Message using the ValidatorSetSig message to be signed by the validator subnet
	// Confirm the event is emitted

	cChainInfo := network.GetPrimaryNetworkInfo()
	contractSubnet, validatorSubnet := utils.GetTwoSubnets(network)
	_, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy a ValidatorSetSigContract to contract subnet
	validatorSetSigContractAddress, validatorSetSig := utils.DeployValidatorSetSig(
		ctx,
		fundedKey,
		contractSubnet,
		validatorSubnet,
	)

	// Deploy a mock ERC20 contract to contract subnet
	exampleERC20ContractAddress, exampleErc20 := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		contractSubnet,
	)
	erc20ABI, err := exampleerc20.ExampleERC20MetaData.GetAbi()
	Expect(err).Should(BeNil())

	// Confirm that the validatorContract has a balance of 0 to begin with
	startingBalance, err := exampleErc20.BalanceOf(&bind.CallOpts{}, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(startingBalance.Cmp(big.NewInt(0))).Should(BeZero())

	// Construct a ValidatorSetSig message with mock ERC20 as the target contract
	// and mint 100 tokens as the TxPayload
	callData, err := erc20ABI.Pack("mint", big.NewInt(100))
	Expect(err).Should(BeNil())

	vssMessage1 := validatorsetsig.ValidatorSetSigMessage{
		ValidatorSetSigAddress: validatorSetSigContractAddress,
		TargetContractAddress:  exampleERC20ContractAddress,
		TargetBlockchainID:     contractSubnet.BlockchainID,
		Nonce:                  big.NewInt(1),
		Payload:                callData,
	}

	// Construct a second ValidatorSetSig message with mock ERC20 as the target contract
	// and mint 50 tokens as the TxPayload
	callData2, err := erc20ABI.Pack("mint", big.NewInt(50))
	Expect(err).Should(BeNil())
	vssMessage2 := validatorsetsig.ValidatorSetSigMessage{
		ValidatorSetSigAddress: validatorSetSigContractAddress,
		TargetContractAddress:  exampleERC20ContractAddress,
		TargetBlockchainID:     contractSubnet.BlockchainID,
		Nonce:                  big.NewInt(2),
		Payload:                callData2,
	}

	networkID := network.GetNetworkID()

	noMessageWarpEnabledChainConfig := utils.GetWarpEnabledChainConfig(nil)
	// Create chain config file with off chain validatorsetsig message
	offchainMessages, warpEnabledChainConfigWithMsg := utils.InitOffChainMessageChainConfigValidatorSetSig(
		networkID,
		validatorSubnet,
		validatorSetSigContractAddress,
		[]validatorsetsig.ValidatorSetSigMessage{vssMessage1, vssMessage2})

	// Create chain config with off chain messages
	chainConfigs := make(map[string]string)
	utils.SetChainConfig(chainConfigs, cChainInfo, noMessageWarpEnabledChainConfig)
	utils.SetChainConfig(chainConfigs, validatorSubnet, warpEnabledChainConfigWithMsg)
	utils.SetChainConfig(chainConfigs, contractSubnet, noMessageWarpEnabledChainConfig)

	// Restart nodes with new chain config
	nodeNames := network.GetAllNodeNames()
	network.RestartNodes(ctx, nodeNames, runner_sdk.WithChainConfigs(chainConfigs))

	// Execute the ValidatorSetSig executeCall and wait for acceptance
	receipt := utils.ExecuteValidatorSetSigCallAndVerify(
		ctx,
		network,
		validatorSubnet,
		contractSubnet,
		validatorSetSigContractAddress,
		fundedKey,
		&offchainMessages[0],
		true,
	)

	// Confirm that the Delivered event is emitted and that validatorSetSig contract has a balance of 100 of ERC20
	deliveredEvent, err := utils.GetEventFromLogs(receipt.Logs, validatorSetSig.ParseDelivered)
	Expect(err).Should(BeNil())
	Expect(deliveredEvent.TargetContractAddress).Should(Equal(exampleERC20ContractAddress))
	Expect(deliveredEvent.Nonce).Should(Equal(big.NewInt(1)))

	endingBalance, err := exampleErc20.BalanceOf(nil, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(endingBalance).Should(Equal(big.NewInt(100)))

	// Resend the same message again and it should fail due to nonce being consumed

	_ = utils.ExecuteValidatorSetSigCallAndVerify(
		ctx,
		network,
		validatorSubnet,
		contractSubnet,
		validatorSetSigContractAddress,
		fundedKey,
		&offchainMessages[0],
		false,
	)

	// Confirm that the validatorSetSig contract still has a balance of 100 of ERC20
	endingBalance, err = exampleErc20.BalanceOf(nil, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(endingBalance).Should(Equal(big.NewInt(100)))

	// Finally send another valid transaction with the incremented nonce
	receipt2 := utils.ExecuteValidatorSetSigCallAndVerify(
		ctx,
		network,
		validatorSubnet,
		contractSubnet,
		validatorSetSigContractAddress,
		fundedKey,
		&offchainMessages[1],
		true,
	)

	// Confirm that the Delivered event is emitted and that validatorSetSig contract has a balance of 150 of ERC20
	deliveredEvent2, err := utils.GetEventFromLogs(receipt2.Logs, validatorSetSig.ParseDelivered)
	Expect(err).Should(BeNil())
	Expect(deliveredEvent2.TargetContractAddress).Should(Equal(exampleERC20ContractAddress))
	Expect(deliveredEvent2.Nonce).Should(Equal(big.NewInt(2)))

	endingBalance, err = exampleErc20.BalanceOf(nil, validatorSetSigContractAddress)
	Expect(err).Should(BeNil())
	Expect(endingBalance).Should(Equal(big.NewInt(150)))
}
