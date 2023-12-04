package price_registry_test

import (
	"context"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	price_registry2 "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/test_utils"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestPriceRegistryFilters(t *testing.T) {
	cl := mocks.NewClient(t)

	test_utils.AssertFilterRegistration(t, new(lpmocks.LogPoller), func(lp *lpmocks.LogPoller, addr common.Address) ccipdata.Closer {
		c, err := price_registry2.NewPriceRegistryV1_0_0(logger.TestLogger(t), addr, lp, cl)
		require.NoError(t, err)
		return c
	}, 3)

	test_utils.AssertFilterRegistration(t, new(lpmocks.LogPoller), func(lp *lpmocks.LogPoller, addr common.Address) ccipdata.Closer {
		c, err := price_registry2.NewPriceRegistryV1_2_0(logger.TestLogger(t), addr, lp, cl)
		require.NoError(t, err)
		return c
	}, 3)
}

type priceRegReaderTH struct {
	lp      logpoller.LogPollerTest
	ec      client.Client
	lggr    logger.Logger
	user    *bind.TransactOpts
	readers map[string]price_registry2.PriceRegistryReader

	// Expected state
	blockTs              []uint64
	expectedFeeTokens    []common.Address
	expectedGasUpdates   map[uint64][]price_registry2.GasPrice
	expectedTokenUpdates map[uint64][]price_registry2.TokenPrice
	dest                 uint64
}

// setupPriceRegistryReaderTH instantiates all versions of the price registry reader
// with a snapshot of data so reader tests can do multi-version assertions.
func setupPriceRegistryReaderTH(t *testing.T) priceRegReaderTH {
	user, ec := test_utils.NewSim(t)
	lggr := logger.TestLogger(t)
	// TODO: We should be able to use an in memory log poller ORM here to speed up the tests.
	lp := logpoller.NewLogPoller(logpoller.NewORM(testutils.SimulatedChainID, pgtest.NewSqlxDB(t), lggr, pgtest.NewQConfig(true)), ec, lggr, 100*time.Millisecond, false, 2, 3, 2, 1000)

	feeTokens := []common.Address{utils.RandomAddress(), utils.RandomAddress()}
	dest := uint64(10)
	gasPriceUpdatesBlock1 := []price_registry2.GasPrice{
		{
			DestChainSelector: dest,
			Value:             big.NewInt(11),
		},
	}
	gasPriceUpdatesBlock2 := []price_registry2.GasPrice{
		{
			DestChainSelector: dest,           // Reset same gas price
			Value:             big.NewInt(12), // Intentionally different from block1
		},
	}
	token1 := utils.RandomAddress()
	token2 := utils.RandomAddress()
	tokenPriceUpdatesBlock1 := []price_registry2.TokenPrice{
		{
			Token: token1,
			Value: big.NewInt(12),
		},
	}
	tokenPriceUpdatesBlock2 := []price_registry2.TokenPrice{
		{
			Token: token1,
			Value: big.NewInt(13), // Intentionally change token1 value
		},
		{
			Token: token2,
			Value: big.NewInt(12), // Intentionally set a same value different token
		},
	}
	addr, _, _, err := price_registry_1_0_0.DeployPriceRegistry(user, ec, nil, feeTokens, 1000)
	require.NoError(t, err)
	addr2, _, _, err := price_registry.DeployPriceRegistry(user, ec, nil, feeTokens, 1000)
	require.NoError(t, err)
	test_utils.CommitAndGetBlockTs(ec) // Deploy these
	pr10r, err := price_registry2.NewPriceRegistryReader(lggr, addr, lp, ec)
	require.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(pr10r).String(), reflect.TypeOf(&price_registry2.PriceRegistryV1_0_0{}).String())
	pr12r, err := price_registry2.NewPriceRegistryReader(lggr, addr2, lp, ec)
	require.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(pr12r).String(), reflect.TypeOf(&price_registry2.PriceRegistryV1_2_0{}).String())
	// Apply block1.
	price_registry2.ApplyPriceRegistryUpdateV1_0_0(t, user, addr, ec, gasPriceUpdatesBlock1, tokenPriceUpdatesBlock1)
	price_registry2.ApplyPriceRegistryUpdateV1_2_0(t, user, addr2, ec, gasPriceUpdatesBlock1, tokenPriceUpdatesBlock1)
	b1 := test_utils.CommitAndGetBlockTs(ec)
	// Apply block2
	price_registry2.ApplyPriceRegistryUpdateV1_0_0(t, user, addr, ec, gasPriceUpdatesBlock2, tokenPriceUpdatesBlock2)
	price_registry2.ApplyPriceRegistryUpdateV1_2_0(t, user, addr2, ec, gasPriceUpdatesBlock2, tokenPriceUpdatesBlock2)
	b2 := test_utils.CommitAndGetBlockTs(ec)

	// Capture all lp data.
	lp.PollAndSaveLogs(context.Background(), 1)

	return priceRegReaderTH{
		lp:   lp,
		ec:   ec,
		lggr: lggr,
		user: user,
		readers: map[string]price_registry2.PriceRegistryReader{
			ccipdata.V1_0_0: pr10r, ccipdata.V1_2_0: pr12r,
		},
		expectedFeeTokens: feeTokens,
		expectedGasUpdates: map[uint64][]price_registry2.GasPrice{
			b1: gasPriceUpdatesBlock1,
			b2: gasPriceUpdatesBlock2,
		},
		expectedTokenUpdates: map[uint64][]price_registry2.TokenPrice{
			b1: tokenPriceUpdatesBlock1,
			b2: tokenPriceUpdatesBlock2,
		},
		blockTs: []uint64{b1, b2},
		dest:    dest,
	}
}

func testPriceRegistryReader(t *testing.T, th priceRegReaderTH, pr price_registry2.PriceRegistryReader) {
	// Assert have expected fee tokens.
	gotFeeTokens, err := pr.GetFeeTokens(context.Background())
	require.NoError(t, err)
	assert.Equal(t, th.expectedFeeTokens, gotFeeTokens)

	// Note unsupported chain selector simply returns an empty set not an error
	gasUpdates, err := pr.GetGasPriceUpdatesCreatedAfter(context.Background(), 1e6, time.Unix(0, 0), 0)
	require.NoError(t, err)
	assert.Len(t, gasUpdates, 0)

	for i, ts := range th.blockTs {
		// Should see all updates >= ts.
		var expectedGas []price_registry2.GasPrice
		var expectedToken []price_registry2.TokenPrice
		for j := i; j < len(th.blockTs); j++ {
			expectedGas = append(expectedGas, th.expectedGasUpdates[th.blockTs[j]]...)
			expectedToken = append(expectedToken, th.expectedTokenUpdates[th.blockTs[j]]...)
		}
		gasUpdates, err = pr.GetGasPriceUpdatesCreatedAfter(context.Background(), th.dest, time.Unix(int64(ts-1), 0), 0)
		require.NoError(t, err)
		assert.Len(t, gasUpdates, len(expectedGas))

		tokenUpdates, err2 := pr.GetTokenPriceUpdatesCreatedAfter(context.Background(), time.Unix(int64(ts-1), 0), 0)
		require.NoError(t, err2)
		assert.Len(t, tokenUpdates, len(expectedToken))
	}

	// Empty token set should return empty set no error.
	gotEmpty, err := pr.GetTokenPrices(context.Background(), []common.Address{})
	require.NoError(t, err)
	assert.Len(t, gotEmpty, 0)

	// We expect latest token prices to apply
	allTokenUpdates, err := pr.GetTokenPriceUpdatesCreatedAfter(context.Background(), time.Unix(0, 0), 0)
	require.NoError(t, err)
	// Build latest map
	latest := make(map[common.Address]*big.Int)
	// Comes back in ascending order (oldest first)
	var allTokens []common.Address
	for i := len(allTokenUpdates) - 1; i >= 0; i-- {
		_, have := latest[allTokenUpdates[i].Data.Token]
		if have {
			continue
		}
		latest[allTokenUpdates[i].Data.Token] = allTokenUpdates[i].Data.Value
		allTokens = append(allTokens, allTokenUpdates[i].Data.Token)
	}
	tokenPrices, err := pr.GetTokenPrices(context.Background(), allTokens)
	require.NoError(t, err)
	require.Len(t, tokenPrices, len(allTokens))
	for _, p := range tokenPrices {
		assert.Equal(t, p.Value, latest[p.Token])
	}

	// We expect 2 fee token events (added/removed). Exact event sigs may differ.
	assert.Len(t, pr.FeeTokenEvents(), 2)

}

func TestPriceRegistryReader(t *testing.T) {
	th := setupPriceRegistryReaderTH(t)
	// Assert all readers produce the same expected results.
	for version, pr := range th.readers {
		pr := pr
		t.Run("PriceRegistryReader"+version, func(t *testing.T) {
			testPriceRegistryReader(t, th, pr)
		})
	}
}
