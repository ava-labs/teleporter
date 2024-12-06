package staking

import (
	"context"
	"log"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	foo "github.com/ava-labs/teleporter/abi-bindings/go/foo-upgradeable/Foo"
	foov2 "github.com/ava-labs/teleporter/abi-bindings/go/foo-upgradeable/FooV2"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

func UpgradeFoo(network *localnetwork.LocalNetwork) {
	subnetAInfo, _ := network.GetTwoSubnets()
	_, fundedKey := network.GetFundedAccountInfo()

	// Generate random address to be the owner address
	ownerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	ownerAddress := crypto.PubkeyToAddress(ownerKey.PublicKey)

	opts, err := bind.NewKeyedTransactorWithChainID(ownerKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())

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

	// Deploy Foo with a proxy
	fooImplAddress, tx, _, err := foo.DeployFoo(opts, subnetAInfo.RPCClient)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx.Hash())

	fooAddress, proxyAdmin := utils.DeployTransparentUpgradeableProxy(
		ctx,
		subnetAInfo,
		ownerKey,
		fooImplAddress,
	)

	// Set a value that's stored in Bar's state
	foo, err := foo.NewFoo(fooAddress, subnetAInfo.RPCClient)
	Expect(err).Should(BeNil())
	foo.SetBar(opts, big.NewInt(42))
	utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx.Hash())

	// Upgrade Foo to FooV2
	fooV2ImplAddress, tx, _, err := foov2.DeployFooV2(opts, subnetAInfo.RPCClient)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx.Hash())

	// Upgrade the TransparentUpgradeableProxy to point to the FooV2 implementation
	tx, err = proxyAdmin.UpgradeAndCall(opts, fooAddress, fooV2ImplAddress, []byte{})
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx.Hash())

	// Re-initialize the FooV2 contract to migrate Bar's state
	fooV2, err := foov2.NewFooV2(fooAddress, subnetAInfo.RPCClient)
	Expect(err).Should(BeNil())
	tx, err = fooV2.Initialize(opts)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx.Hash())

	// Check that the value is still stored in Bar's state
	value, err := fooV2.GetBar(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	log.Println(value.Uint64())
	Expect(value.Cmp(big.NewInt(42))).Should(Equal(0))
}
