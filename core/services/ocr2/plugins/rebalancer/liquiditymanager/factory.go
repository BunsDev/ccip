package liquiditymanager

import (
	"fmt"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

// Factory initializes a new liquidity manager instance.
//
//go:generate mockery --quiet --name Factory --output ../rebalancermocks --filename lm_factory_mock.go --case=underscore
type Factory interface {
	NewRebalancer(networkID models.NetworkSelector, address models.Address) (Rebalancer, error)
}

type evmDep struct {
	lp        logpoller.LogPoller
	ethClient client.Client
}

type BaseRebalancerFactory struct {
	evmDeps map[models.NetworkSelector]evmDep
	tokens  map[models.NetworkSelector]models.Address
}

type Opt func(f *BaseRebalancerFactory)

func NewBaseRebalancerFactory(tokens map[models.NetworkSelector]models.Address, opts ...Opt) *BaseRebalancerFactory {
	f := &BaseRebalancerFactory{
		evmDeps: make(map[models.NetworkSelector]evmDep),
		tokens:  tokens,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func WithEvmDep(networkID models.NetworkSelector, lp logpoller.LogPoller, ethClient client.Client) Opt {
	return func(f *BaseRebalancerFactory) {
		f.evmDeps[networkID] = evmDep{
			lp:        lp,
			ethClient: ethClient,
		}
	}
}

func (b *BaseRebalancerFactory) NewRebalancer(networkID models.NetworkSelector, address models.Address) (Rebalancer, error) {
	switch typ := networkID.Type(); typ {
	case models.NetworkTypeEvm:
		evmDeps, exists := b.evmDeps[networkID]
		if !exists {
			return nil, fmt.Errorf("evm dependencies not found")
		}

		tokenAddr, exists := b.tokens[networkID]
		if !exists {
			return nil, fmt.Errorf("token address for network %d not found", networkID)
		}

		return NewEvmRebalancer(address, networkID, evmDeps.ethClient, evmDeps.lp, tokenAddr)
	default:
		return nil, fmt.Errorf("liquidity manager of type %v is not supported", typ)
	}
}
