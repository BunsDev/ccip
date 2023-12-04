package test_utils

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

// NewSimulation returns a client and a simulated backend.
func NewSimulation(t *testing.T) (*bind.TransactOpts, *client.SimulatedBackendClient) {
	user := testutils.MustNewSimTransactor(t)
	simulatedBackend := backends.NewSimulatedBackend(map[common.Address]core.GenesisAccount{
		user.From: {
			Balance: big.NewInt(0).Mul(big.NewInt(3), big.NewInt(1e18)),
		},
	}, 10e6)
	simulatedBackendClient := client.NewSimulatedBackendClient(t, simulatedBackend, testutils.SimulatedChainID)
	return user, simulatedBackendClient
}

// AssertNonRevert Verify that a transaction was not reverted.
func AssertNonRevert(t *testing.T, tx *types.Transaction, bc *client.SimulatedBackendClient, user *bind.TransactOpts) {
	require.NotNil(t, tx, "Transaction should not be nil")
	receipt, err := bc.TransactionReceipt(user.Context, tx.Hash())
	require.NoError(t, err)
	require.NotEqual(t, uint64(0), receipt.Status, "Transaction should not have reverted")
}

func AssertFilterRegistration(t *testing.T, lp *lpmocks.LogPoller, buildCloser func(lp *lpmocks.LogPoller, addr common.Address) ccipdata.Closer, numFilter int) {
	// Expected filter properties for a closer:
	// - Should be the same filter set registered that is unregistered
	// - Should be registered to the address specified
	// - Number of events specific to this component should be registered
	addr := common.HexToAddress("0x1234")
	var filters []logpoller.Filter

	lp.On("RegisterFilter", mock.Anything).Run(func(args mock.Arguments) {
		f := args.Get(0).(logpoller.Filter)
		require.Equal(t, len(f.Addresses), 1)
		require.Equal(t, f.Addresses[0], addr)
		filters = append(filters, f)
	}).Return(nil).Times(numFilter)

	c := buildCloser(lp, addr)
	for _, filter := range filters {
		lp.On("UnregisterFilter", filter.Name).Return(nil)
	}

	require.NoError(t, c.Close())
	lp.AssertExpectations(t)
}

func CommitAndGetBlockTs(ec *client.SimulatedBackendClient) uint64 {
	h := ec.Commit()
	b, _ := ec.BlockByHash(context.Background(), h)
	return b.Time()
}

func NewSim(t *testing.T) (*bind.TransactOpts, *client.SimulatedBackendClient) {
	user := testutils.MustNewSimTransactor(t)
	sim := backends.NewSimulatedBackend(map[common.Address]core.GenesisAccount{
		user.From: {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)),
		},
	}, 10e6)
	ec := client.NewSimulatedBackendClient(t, sim, testutils.SimulatedChainID)
	return user, ec
}
