package utils

import (
	"context"
	"crypto/ecdsa"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	proxyadmin "github.com/ava-labs/teleporter/abi-bindings/go/ProxyAdmin"
	transparentupgradeableproxy "github.com/ava-labs/teleporter/abi-bindings/go/TransparentUpgradeableProxy"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

func DeployTransparentUpgradeableProxy[T any](
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	senderKey *ecdsa.PrivateKey,
	implAddress common.Address,
	newInstance func(address common.Address, backend bind.ContractBackend) (*T, error),
) (common.Address, *proxyadmin.ProxyAdmin, *T) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should((BeNil()))

	senderAddress := crypto.PubkeyToAddress(senderKey.PublicKey)
	proxyAddress, tx, proxy, err := transparentupgradeableproxy.DeployTransparentUpgradeableProxy(
		opts,
		subnet.RPCClient,
		implAddress,
		senderAddress,
		[]byte{},
	)
	Expect(err).Should(BeNil())
	receipt := WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	proxyAdminEvent, err := GetEventFromLogs(receipt.Logs, proxy.ParseAdminChanged)
	Expect(err).Should(BeNil())

	proxyAdmin, err := proxyadmin.NewProxyAdmin(proxyAdminEvent.NewAdmin, subnet.RPCClient)
	Expect(err).Should(BeNil())

	contract, err := newInstance(proxyAddress, subnet.RPCClient)
	Expect(err).Should(BeNil())

	return proxyAddress, proxyAdmin, contract
}
