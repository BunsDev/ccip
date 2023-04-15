package ccip

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ocrConfigHelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"golang.org/x/crypto/curve25519"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"

	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/maybe_revert_message_receiver"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/mock_afn_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/weth9"
	ccipplugin "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
)

// CCIPContractsDeployer provides the implementations for deploying CCIP ETH contracts
type CCIPContractsDeployer struct {
	evmClient   blockchain.EVMClient
	EthDeployer *contracts.EthereumContractDeployer
}

// NewCCIPContractsDeployer returns an instance of a contract deployer for CCIP
func NewCCIPContractsDeployer(bcClient blockchain.EVMClient) (*CCIPContractsDeployer, error) {
	return &CCIPContractsDeployer{
		evmClient:   bcClient,
		EthDeployer: contracts.NewEthereumContractDeployer(bcClient),
	}, nil
}

func (e *CCIPContractsDeployer) DeployLinkTokenContract() (*LinkToken, error) {
	address, _, instance, err := e.evmClient.DeployContract("Link Token", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return link_token_interface.DeployLinkToken(auth, backend)
	})

	if err != nil {
		return nil, err
	}
	return &LinkToken{
		client:     e.evmClient,
		instance:   instance.(*link_token_interface.LinkToken),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) NewLinkTokenContract(addr common.Address) (*LinkToken, error) {
	token, err := link_token_interface.NewLinkToken(addr, e.evmClient.Backend())

	if err != nil {
		return nil, err
	}
	log.Info().
		Str("Contract Address", addr.Hex()).
		Str("Contract Name", "Link Token").
		Str("From", e.evmClient.GetDefaultWallet().Address()).
		Str("Network Name", e.evmClient.GetNetworkConfig().Name).
		Msg("New contract")
	return &LinkToken{
		client:     e.evmClient,
		instance:   token,
		EthAddress: addr,
	}, err
}

func (e *CCIPContractsDeployer) NewLockReleaseTokenPoolContract(addr common.Address) (
	*LockReleaseTokenPool,
	error,
) {
	pool, err := lock_release_token_pool.NewLockReleaseTokenPool(addr, e.evmClient.Backend())

	if err != nil {
		return nil, err
	}
	log.Info().
		Str("Contract Address", addr.Hex()).
		Str("Contract Name", "Native Token Pool").
		Str("From", e.evmClient.GetDefaultWallet().Address()).
		Str("Network Name", e.evmClient.GetNetworkConfig().Name).
		Msg("New contract")
	return &LockReleaseTokenPool{
		client:     e.evmClient,
		instance:   pool,
		EthAddress: addr,
	}, err
}

func (e *CCIPContractsDeployer) DeployLockReleaseTokenPoolContract(linkAddr string) (
	*LockReleaseTokenPool,
	error,
) {
	log.Debug().Str("token", linkAddr).Msg("Deploying native token pool")
	token := common.HexToAddress(linkAddr)
	address, _, instance, err := e.evmClient.DeployContract("Native Token Pool", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return lock_release_token_pool.DeployLockReleaseTokenPool(
			auth,
			backend,
			token,
			lock_release_token_pool.RateLimiterConfig{
				Capacity:  new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e9)),
				Rate:      new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e5)),
				IsEnabled: true,
			})
	})

	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPool{
		client:     e.evmClient,
		instance:   instance.(*lock_release_token_pool.LockReleaseTokenPool),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployAFNContract() (*AFN, error) {
	address, _, instance, err := e.evmClient.DeployContract("Mock AFN Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return mock_afn_contract.DeployMockAFNContract(auth, backend)
	})
	if err != nil {
		return nil, err
	}
	return &AFN{
		client:     e.evmClient,
		instance:   instance.(*mock_afn_contract.MockAFNContract),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) NewAFNContract(addr common.Address) (*AFN, error) {
	afn, err := mock_afn_contract.NewMockAFNContract(addr, e.evmClient.Backend())
	if err != nil {
		return nil, err
	}
	log.Info().
		Str("Contract Address", addr.Hex()).
		Str("Contract Name", "Mock AFN Contract").
		Str("From", e.evmClient.GetDefaultWallet().Address()).
		Str("Network Name", e.evmClient.GetNetworkConfig().Name).
		Msg("New contract")

	return &AFN{
		client:     e.evmClient,
		instance:   afn,
		EthAddress: addr,
	}, err
}

func (e *CCIPContractsDeployer) NewCommitStore(addr common.Address) (
	*CommitStore,
	error,
) {
	ins, err := commit_store.NewCommitStore(addr, e.evmClient.Backend())
	log.Info().
		Str("Contract Address", addr.Hex()).
		Str("Contract Name", "CommitStore").
		Str("From", e.evmClient.GetDefaultWallet().Address()).
		Str("Network Name", e.evmClient.GetNetworkConfig().Name).
		Msg("New contract")
	return &CommitStore{
		client:     e.evmClient,
		Instance:   ins,
		EthAddress: addr,
	}, err
}

func (e *CCIPContractsDeployer) DeployCommitStore(sourceChainId, destChainId uint64, afn, onRamp, priceRegistry common.Address) (*CommitStore, error) {
	address, _, instance, err := e.evmClient.DeployContract("CommitStore Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return commit_store.DeployCommitStore(
			auth,
			backend,
			commit_store.CommitStoreStaticConfig{
				ChainId:       destChainId,
				SourceChainId: sourceChainId,
				OnRamp:        onRamp,
			},
			commit_store.CommitStoreDynamicConfig{
				PriceRegistry: priceRegistry,
				Afn:           afn,
			},
		)
	})
	if err != nil {
		return nil, err
	}
	return &CommitStore{
		client:     e.evmClient,
		Instance:   instance.(*commit_store.CommitStore),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeploySimpleMessageReceiver() (
	*MessageReceiver,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("SimpleMessageReceiver Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return simple_message_receiver.DeploySimpleMessageReceiver(auth, backend)
	})
	if err != nil {
		return nil, err
	}
	return &MessageReceiver{
		client:     e.evmClient,
		instance:   instance.(*simple_message_receiver.SimpleMessageReceiver),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployReceiverDapp(router common.Address) (
	*ReceiverDapp,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("ReceiverDapp", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return maybe_revert_message_receiver.DeployMaybeRevertMessageReceiver(auth, backend, false)
	})
	if err != nil {
		return nil, err
	}
	return &ReceiverDapp{
		client:     e.evmClient,
		instance:   instance.(*maybe_revert_message_receiver.MaybeRevertMessageReceiver),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) NewReceiverDapp(addr common.Address) (
	*ReceiverDapp,
	error,
) {
	ins, err := maybe_revert_message_receiver.NewMaybeRevertMessageReceiver(addr, e.evmClient.Backend())
	log.Info().
		Str("Contract Address", addr.Hex()).
		Str("Contract Name", "ReceiverDapp").
		Str("From", e.evmClient.GetDefaultWallet().Address()).
		Str("Network Name", e.evmClient.GetNetworkConfig().Name).
		Msg("New contract")
	return &ReceiverDapp{
		client:     e.evmClient,
		instance:   ins,
		EthAddress: addr,
	}, err
}

func (e *CCIPContractsDeployer) DeployRouter(wrappedNative common.Address) (
	*Router,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("Router", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return router.DeployRouter(auth, backend, wrappedNative)
	})
	if err != nil {
		return nil, err
	}
	return &Router{
		client:     e.evmClient,
		Instance:   instance.(*router.Router),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) NewRouter(addr common.Address) (
	*Router,
	error,
) {
	r, err := router.NewRouter(addr, e.evmClient.Backend())
	log.Info().
		Str("Contract Address", addr.Hex()).
		Str("Contract Name", "Router").
		Str("From", e.evmClient.GetDefaultWallet().Address()).
		Str("Network Name", e.evmClient.GetNetworkConfig().Name).
		Msg("New contract")
	if err != nil {
		return nil, err
	}
	return &Router{
		client:     e.evmClient,
		Instance:   r,
		EthAddress: addr,
	}, err
}

func (e *CCIPContractsDeployer) NewPriceRegistry(addr common.Address) (
	*PriceRegistry,
	error,
) {
	ins, err := price_registry.NewPriceRegistry(addr, e.evmClient.Backend())
	log.Info().
		Str("Contract Address", addr.Hex()).
		Str("Contract Name", "PriceRegistry").
		Str("From", e.evmClient.GetDefaultWallet().Address()).
		Str("Network Name", e.evmClient.GetNetworkConfig().Name).
		Msg("New contract")
	return &PriceRegistry{
		client:     e.evmClient,
		instance:   ins,
		EthAddress: addr,
	}, err
}

func (e *CCIPContractsDeployer) DeployPriceRegistry(
	feeUpdates price_registry.InternalPriceUpdates,
) (
	*PriceRegistry,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("PriceRegistry", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		feeTokens := []common.Address{}
		if len(feeUpdates.TokenPriceUpdates) > 0 {
			feeTokens = append(feeTokens, feeUpdates.TokenPriceUpdates[0].SourceToken)
		}
		return price_registry.DeployPriceRegistry(auth, backend, feeUpdates, nil, feeTokens, 60*60*24*14)
	})
	if err != nil {
		return nil, err
	}
	return &PriceRegistry{
		client:     e.evmClient,
		instance:   instance.(*price_registry.PriceRegistry),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) NewOnRamp(addr common.Address) (
	*OnRamp,
	error,
) {
	ins, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(addr, e.evmClient.Backend())
	log.Info().
		Str("Contract Address", addr.Hex()).
		Str("Contract Name", "OnRamp").
		Str("From", e.evmClient.GetDefaultWallet().Address()).
		Str("Network Name", e.evmClient.GetNetworkConfig().Name).
		Msg("New contract")
	return &OnRamp{
		client:     e.evmClient,
		Instance:   ins,
		EthAddress: addr,
	}, err
}

func (e *CCIPContractsDeployer) DeployOnRamp(
	sourceChainId, destChainId uint64,
	allowList []common.Address,
	tokensAndPools []evm_2_evm_onramp.EVM2EVMOnRampTokenAndPool,
	afn, router, priceRegistry common.Address,
	opts RateLimiterConfig,
	feeConfig []evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs,
	linkTokenAddress common.Address,
) (
	*OnRamp,
	error,
) {
	address, _, instance, err := e.evmClient.DeployContract("OnRamp", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return evm_2_evm_onramp.DeployEVM2EVMOnRamp(
			auth,
			backend,

			evm_2_evm_onramp.EVM2EVMOnRampStaticConfig{
				LinkToken:         linkTokenAddress,
				ChainId:           sourceChainId, // source chain id
				DestChainId:       destChainId,   // destinationChainId
				DefaultTxGasLimit: 200_000,
			},
			evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig{
				Router:          router,
				PriceRegistry:   priceRegistry,
				MaxDataSize:     1e5,
				MaxTokensLength: 5,
				MaxGasLimit:     4_000_000,
				Afn:             afn,
			},
			tokensAndPools,
			allowList,
			evm_2_evm_onramp.RateLimiterConfig{
				Capacity: opts.Capacity,
				Rate:     opts.Rate,
			},
			feeConfig,
			[]evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight{},
		)
	})
	if err != nil {
		return nil, err
	}
	return &OnRamp{
		client:     e.evmClient,
		Instance:   instance.(*evm_2_evm_onramp.EVM2EVMOnRamp),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) NewOffRamp(addr common.Address) (
	*OffRamp,
	error,
) {
	ins, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(addr, e.evmClient.Backend())
	log.Info().
		Str("Contract Address", addr.Hex()).
		Str("Contract Name", "OffRamp").
		Str("From", e.evmClient.GetDefaultWallet().Address()).
		Str("Network Name", e.evmClient.GetNetworkConfig().Name).
		Msg("New contract")
	return &OffRamp{
		client:     e.evmClient,
		Instance:   ins,
		EthAddress: addr,
	}, err
}

func (e *CCIPContractsDeployer) DeployOffRamp(sourceChainId, destChainId uint64, commitStore, onRamp, afn, destRouter common.Address, sourceToken, pools []common.Address, opts RateLimiterConfig) (*OffRamp, error) {
	address, _, instance, err := e.evmClient.DeployContract("OffRamp Contract", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return evm_2_evm_offramp.DeployEVM2EVMOffRamp(
			auth,
			backend,
			evm_2_evm_offramp.EVM2EVMOffRampStaticConfig{
				CommitStore:   commitStore,
				ChainId:       destChainId,
				SourceChainId: sourceChainId,
				OnRamp:        onRamp,
			},
			evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig{
				PermissionLessExecutionThresholdSeconds: 24 * 7 * 60 * 60,
				Router:                                  destRouter,
				Afn:                                     afn,
				MaxDataSize:                             1e5,
				MaxTokensLength:                         15,
			},
			sourceToken,
			pools,
			evm_2_evm_offramp.RateLimiterConfig{
				Rate:      opts.Rate,
				Capacity:  opts.Capacity,
				IsEnabled: true,
			},
		)
	})
	if err != nil {
		return nil, err
	}
	return &OffRamp{
		client:     e.evmClient,
		Instance:   instance.(*evm_2_evm_offramp.EVM2EVMOffRamp),
		EthAddress: *address,
	}, err
}

func (e *CCIPContractsDeployer) DeployWrappedNative() (*common.Address, error) {
	address, _, _, err := e.evmClient.DeployContract("WrappedNative", func(
		auth *bind.TransactOpts,
		backend bind.ContractBackend,
	) (common.Address, *types.Transaction, interface{}, error) {
		return weth9.DeployWETH9(auth, backend)
	})
	if err != nil {
		return nil, err
	}
	return address, err
}

func DefaultOffChainAggregatorV2Config(numberNodes int) contracts.OffChainAggregatorV2Config {
	if numberNodes <= 4 {
		log.Err(fmt.Errorf("insufficient number of nodes (%d) supplied for OCR, need at least 5", numberNodes)).
			Int("Number Chainlink Nodes", numberNodes).
			Msg("You likely need more chainlink nodes to properly configure OCR, try 5 or more.")
	}
	s := make([]int, 0)
	for i := 0; i < numberNodes; i++ {
		s = append(s, 1)
	}
	faultyNodes := 0
	if numberNodes > 1 {
		faultyNodes = numberNodes/3 - 1
	}
	if faultyNodes == 0 {
		faultyNodes = 1
	}
	return contracts.OffChainAggregatorV2Config{
		DeltaProgress:                           70 * time.Second,
		DeltaResend:                             5 * time.Second,
		DeltaRound:                              30 * time.Second,
		DeltaGrace:                              2 * time.Second,
		DeltaStage:                              40 * time.Second,
		RMax:                                    3,
		S:                                       s,
		F:                                       faultyNodes,
		Oracles:                                 []ocrConfigHelper2.OracleIdentityExtra{},
		MaxDurationQuery:                        5 * time.Second,
		MaxDurationObservation:                  32 * time.Second,
		MaxDurationReport:                       20 * time.Second,
		MaxDurationShouldAcceptFinalizedReport:  10 * time.Second,
		MaxDurationShouldTransmitAcceptedReport: 10 * time.Second,
		OnchainConfig:                           []byte{},
	}
}

func stripKeyPrefix(key string) string {
	chunks := strings.Split(key, "_")
	if len(chunks) == 3 {
		return chunks[2]
	}
	return key
}

func NewOffChainAggregatorV2Config[T ccipplugin.OffchainConfig](
	nodes []*client.CLNodesWithKeys,
	offchainCfg T,
) (
	signers []common.Address,
	transmitters []common.Address,
	f_ uint8,
	onchainConfig_ []byte,
	offchainConfigVersion uint64,
	offchainConfig []byte,
	err error,
) {
	oracleIdentities := make([]ocrConfigHelper2.OracleIdentityExtra, 0)
	ocrConfig := DefaultOffChainAggregatorV2Config(len(nodes))
	var onChainKeys []ocrtypes2.OnchainPublicKey
	for i, nodeWithKeys := range nodes {
		ocr2Key := nodeWithKeys.KeysBundle.OCR2Key.Data
		offChainPubKeyTemp, err := hex.DecodeString(stripKeyPrefix(ocr2Key.Attributes.OffChainPublicKey))
		if err != nil {
			return nil, nil, 0, nil, 0, nil, err
		}
		formattedOnChainPubKey := stripKeyPrefix(ocr2Key.Attributes.OnChainPublicKey)
		cfgPubKeyTemp, err := hex.DecodeString(stripKeyPrefix(ocr2Key.Attributes.ConfigPublicKey))
		if err != nil {
			return nil, nil, 0, nil, 0, nil, err
		}
		cfgPubKeyBytes := [ed25519.PublicKeySize]byte{}
		copy(cfgPubKeyBytes[:], cfgPubKeyTemp)
		offChainPubKey := [curve25519.PointSize]byte{}
		copy(offChainPubKey[:], offChainPubKeyTemp)
		ethAddress := nodeWithKeys.KeysBundle.EthAddress
		p2pKeys := nodeWithKeys.KeysBundle.P2PKeys
		peerID := p2pKeys.Data[0].Attributes.PeerID
		oracleIdentities = append(oracleIdentities, ocrConfigHelper2.OracleIdentityExtra{
			OracleIdentity: ocrConfigHelper2.OracleIdentity{
				OffchainPublicKey: offChainPubKey,
				OnchainPublicKey:  common.HexToAddress(formattedOnChainPubKey).Bytes(),
				PeerID:            peerID,
				TransmitAccount:   ocrtypes2.Account(ethAddress),
			},
			ConfigEncryptionPublicKey: cfgPubKeyBytes,
		})
		onChainKeys = append(onChainKeys, oracleIdentities[i].OnchainPublicKey)
		transmitters = append(transmitters, common.HexToAddress(ethAddress))
	}
	signers, err = ocrcommon.OnchainPublicKeyToAddress(onChainKeys)
	if err != nil {
		return nil, nil, 0, nil, 0, nil, err
	}
	ocrConfig.Oracles = oracleIdentities
	ocrConfig.ReportingPluginConfig, err = ccipplugin.EncodeOffchainConfig[T](offchainCfg)
	if err != nil {
		return nil, nil, 0, nil, 0, nil, err
	}
	_, _, f_, onchainConfig_, offchainConfigVersion, offchainConfig, err = ocrConfigHelper2.ContractSetConfigArgsForTests(
		ocrConfig.DeltaProgress,
		ocrConfig.DeltaResend,
		ocrConfig.DeltaRound,
		ocrConfig.DeltaGrace,
		ocrConfig.DeltaStage,
		ocrConfig.RMax,
		ocrConfig.S,
		ocrConfig.Oracles,
		ocrConfig.ReportingPluginConfig,
		ocrConfig.MaxDurationQuery,
		ocrConfig.MaxDurationObservation,
		ocrConfig.MaxDurationReport,
		ocrConfig.MaxDurationShouldAcceptFinalizedReport,
		ocrConfig.MaxDurationShouldTransmitAcceptedReport,
		ocrConfig.F,
		ocrConfig.OnchainConfig,
	)
	return
}
