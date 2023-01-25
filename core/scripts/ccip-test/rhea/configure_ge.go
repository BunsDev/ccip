package rhea

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/fee_manager"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ge_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

func setOffRampOnTokenPools(t *testing.T, destClient *EvmDeploymentConfig) {
	for _, tokenConfig := range destClient.ChainConfig.SupportedTokens {
		pool, err := lock_release_token_pool.NewLockReleaseTokenPool(tokenConfig.Pool, destClient.Client)
		require.NoError(t, err)

		// Configure offramp address on pool
		tx, err := pool.SetOffRamp(destClient.Owner, destClient.LaneConfig.OffRamp, true)
		require.NoError(t, err)
		shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
		destClient.Logger.Infof("Offramp pool configured with offramp address: %s", helpers.ExplorerLink(int64(destClient.ChainConfig.ChainId), tx.Hash()))
	}
}

func setFeeManagerPrices(t *testing.T, client *EvmDeploymentConfig, destChainId uint64) {
	feeManager, err := fee_manager.NewFeeManager(client.ChainConfig.FeeManager, client.Client)
	require.NoError(t, err)

	tx, err := feeManager.UpdateFees(client.Owner, []fee_manager.GEFeeUpdate{
		{
			SourceFeeToken:              client.ChainConfig.LinkToken,
			DestChainId:                 destChainId,
			FeeTokenBaseUnitsPerUnitGas: big.NewInt(1e18),
		},
	})
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
}

func setOnRampOnRouter(t *testing.T, sourceClient *EvmDeploymentConfig, destChainId uint64) {
	sourceClient.Logger.Infof("Setting the onRamp on the Router")
	router, err := ge_router.NewGERouter(sourceClient.ChainConfig.Router, sourceClient.Client)
	require.NoError(t, err)
	sourceClient.Logger.Infof("Registering new onRamp")
	tx, err := router.SetOnRamp(sourceClient.Owner, destChainId, sourceClient.LaneConfig.OnRamp)
	require.NoError(t, err)
	shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
}

func setOnRampOnTokenPools(t *testing.T, sourceClient *EvmDeploymentConfig) {
	for _, tokenConfig := range sourceClient.ChainConfig.SupportedTokens {
		pool, err := lock_release_token_pool.NewLockReleaseTokenPool(tokenConfig.Pool, sourceClient.Client)
		require.NoError(t, err)

		// Configure offramp address on pool
		tx, err := pool.SetOnRamp(sourceClient.Owner, sourceClient.LaneConfig.OnRamp, true)
		require.NoError(t, err)
		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("Onramp pool configured with offramp address: %s", helpers.ExplorerLink(int64(sourceClient.ChainConfig.ChainId), tx.Hash()))
	}
}

func setOnRampOnCommitStore(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	commitStore, err := commit_store.NewCommitStore(destClient.LaneConfig.CommitStore, destClient.Client)
	require.NoError(t, err)

	config, err := commitStore.GetCommitStoreConfig(&bind.CallOpts{})
	require.NoError(t, err)

	config.OnRamps = append(config.OnRamps, sourceClient.LaneConfig.OnRamp)
	config.MinSeqNrByOnRamp = append(config.MinSeqNrByOnRamp, 1)

	tx, err := commitStore.SetCommitStoreConfig(destClient.Owner, config)
	require.NoError(t, err)
	destClient.Logger.Infof(fmt.Sprintf("Adding new onRamp to commitStore in tx %s", helpers.ExplorerLink(int64(destClient.ChainConfig.ChainId), tx.Hash())))
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
}

func setRouterOnOffRamp(t *testing.T, destClient *EvmDeploymentConfig) {
	offRamp, err := evm_2_evm_ge_offramp.NewEVM2EVMGEOffRamp(destClient.LaneConfig.OffRamp, destClient.Client)
	require.NoError(t, err)
	tx, err := offRamp.SetRouter(destClient.Owner, destClient.ChainConfig.Router)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.Logger.Infof(fmt.Sprintf("Router set on offRamp in tx %s", helpers.ExplorerLink(int64(destClient.ChainConfig.ChainId), tx.Hash())))
}

func setOffRampOnRouter(t *testing.T, client *EvmDeploymentConfig) {
	client.Logger.Infof("Setting the offRamp on the Router")
	router, err := ge_router.NewGERouter(client.ChainConfig.Router, client.Client)
	require.NoError(t, err)

	isOffRamp, err := router.IsOffRamp(&bind.CallOpts{}, client.LaneConfig.OffRamp)
	require.NoError(t, err)
	if isOffRamp {
		client.Logger.Infof("OffRamp already configured on router. Skipping")
		return
	}

	tx, err := router.AddOffRamp(client.Owner, client.LaneConfig.OffRamp)
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
}

func setFeeManagerUpdater(t *testing.T, client *EvmDeploymentConfig) {
	feeManager, err := fee_manager.NewFeeManager(client.ChainConfig.FeeManager, client.Client)
	require.NoError(t, err)

	tx, err := feeManager.SetFeeUpdater(client.Owner, client.LaneConfig.OffRamp)
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
}

/*
func revokeOffRampOnOffRampRouter(t *testing.T, destClient *EvmDeploymentConfig, offRamp common.Address) {
	destClient.Logger.Infof("Revoking the offRamp on the offRampRouter")
	offRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(destClient.ChainConfig.OffRampRouter, destClient.Client)
	require.NoError(t, err)

	tx, err := offRampRouter.RemoveOffRamp(destClient.Owner, offRamp)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
}
*/

/*
func revokeOffRampOnTokenPools(t *testing.T, destClient *EvmDeploymentConfig, offRamp common.Address) {
	// TODO
}
*/

func fillPoolWithTokens(t *testing.T, client *EvmDeploymentConfig, pool *lock_release_token_pool.LockReleaseTokenPool) {
	destLinkToken, err := link_token_interface.NewLinkToken(client.ChainConfig.LinkToken, client.Client)
	require.NoError(t, err)

	// fill offramp token pool with 0.5 LINK
	amount := big.NewInt(5e17)
	tx, err := destLinkToken.Approve(client.Owner, pool.Address(), amount)
	require.NoError(t, err)
	client.Logger.Infof("Approving token to the token pool: %s", helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	tx, err = pool.AddLiquidity(client.Owner, amount)
	require.NoError(t, err)
	client.Logger.Infof("Adding liquidity to the token pool: %s", helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	client.Logger.Infof("Pool filled with tokens: %s", helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
}
