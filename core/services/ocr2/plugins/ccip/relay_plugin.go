package ccip

import (
	"encoding/json"
	"math/big"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	type_and_version "github.com/smartcontractkit/chainlink/core/gethwrappers/generated/type_and_version_interface_wrapper"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	ccipconfig "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
)

type ContractType string

var (
	EVM2EVMTollOnRamp          ContractType = "EVM2EVMTollOnRamp"
	EVM2EVMTollOffRamp         ContractType = "EVM2EVMTollOffRamp"
	EVM2EVMSubscriptionOnRamp  ContractType = "EVM2EVMSubscriptionOnRamp"
	EVM2EVMSubscriptionOffRamp ContractType = "EVM2EVMSubscriptionOffRamp"
	BlobVerifier               ContractType = "BlobVerifier"
	ContractTypes                           = map[ContractType]struct{}{
		EVM2EVMTollOnRamp:          {},
		EVM2EVMTollOffRamp:         {},
		EVM2EVMSubscriptionOnRamp:  {},
		EVM2EVMSubscriptionOffRamp: {},
		BlobVerifier:               {},
	}
)

func TypeAndVersion(addr common.Address, client bind.ContractBackend) (ContractType, semver.Version, error) {
	tv, err := type_and_version.NewTypeAndVersionInterface(addr, client)
	if err != nil {
		return "", semver.Version{}, errors.Wrap(err, "failed creating a type and version")
	}
	tvStr, err := tv.TypeAndVersion(nil)
	if err != nil {
		return "", semver.Version{}, errors.Wrap(err, "failed to call type and version")
	}
	typeAndVersionValues := strings.Split(tvStr, " ")
	contractType, version := typeAndVersionValues[0], typeAndVersionValues[1]
	v, err := semver.NewVersion(version)
	if err != nil {
		return "", semver.Version{}, err
	}
	if _, ok := ContractTypes[ContractType(contractType)]; !ok {
		return "", semver.Version{}, errors.Errorf("unrecognized contract type %v", contractType)
	}
	return ContractType(contractType), *v, nil
}

func NewRelayServices(lggr logger.Logger, spec *job.OCR2OracleSpec, chainSet evm.ChainSet, new bool, argsNoPlugin libocr2.OracleArgs) ([]job.ServiceCtx, error) {
	var pluginConfig ccipconfig.RelayPluginConfig
	err := json.Unmarshal(spec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return nil, err
	}
	err = pluginConfig.ValidateRelayPluginConfig()
	if err != nil {
		return nil, err
	}
	lggr.Infof("CCIP relay plugin initialized with offchainConfig: %+v", pluginConfig)

	sourceChainId, destChainId := big.NewInt(0).SetUint64(pluginConfig.SourceChainID), big.NewInt(0).SetUint64(pluginConfig.DestChainID)

	sourceChain, err := chainSet.Get(sourceChainId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open source chain")
	}
	destChain, err := chainSet.Get(destChainId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open destination chain")
	}

	if !common.IsHexAddress(spec.ContractID) {
		return nil, errors.Wrap(err, "spec.ContractID is not a valid hex address")
	}
	blobVerifier, err := blob_verifier.NewBlobVerifier(common.HexToAddress(spec.ContractID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading the blobVerifier")
	}
	onRampSeqParsers := make(map[common.Address]func(log logpoller.Log) (uint64, error))
	onRampToReqEventSig := make(map[common.Address]common.Hash)
	var onRamps []common.Address
	var onRampToHasher = make(map[common.Address]LeafHasher[[32]byte])
	hashingCtx := hasher.NewKeccakCtx()

	for _, onRampID := range pluginConfig.OnRampIDs {
		addr := common.HexToAddress(onRampID)
		onRamps = append(onRamps, addr)
		contractType, _, err2 := TypeAndVersion(addr, sourceChain.Client())
		if err2 != nil {
			return nil, errors.Errorf("failed getting type and version %v", err2)
		}

		switch contractType {
		case EVM2EVMTollOnRamp:
			onRamp, err3 := evm_2_evm_toll_onramp.NewEVM2EVMTollOnRamp(addr, sourceChain.Client())
			if err3 != nil {
				return nil, errors.Wrap(err3, "failed creating a new onramp")
			}
			onRampSeqParsers[common.HexToAddress(onRampID)] = func(log logpoller.Log) (uint64, error) {
				req, err4 := onRamp.ParseCCIPSendRequested(types.Log{Data: log.Data, Topics: log.GetTopics()})
				if err4 != nil {
					lggr.Warnf("failed to parse log: %+v", log)
					return 0, err4
				}
				return req.Message.SequenceNumber, nil
			}
			// Subscribe to all relevant relay logs.
			_, err = sourceChain.LogPoller().RegisterFilter(logpoller.Filter{EventSigs: []common.Hash{CCIPTollSendRequested}, Addresses: []common.Address{onRamp.Address()}})
			if err != nil {
				return nil, err
			}
			onRampToReqEventSig[onRamp.Address()] = CCIPTollSendRequested
			onRampToHasher[onRamp.Address()] = NewTollLeafHasher(sourceChainId, destChainId, onRamp.Address(), hashingCtx)
		case EVM2EVMSubscriptionOnRamp:
			onRamp, err3 := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(addr, sourceChain.Client())
			if err3 != nil {
				return nil, errors.Wrap(err3, "failed creating a new onramp")
			}
			onRampSeqParsers[common.HexToAddress(onRampID)] = func(log logpoller.Log) (uint64, error) {
				req, err4 := onRamp.ParseCCIPSendRequested(types.Log{Data: log.Data, Topics: log.GetTopics()})
				if err4 != nil {
					lggr.Warnf("failed to parse log: %+v", log)
					return 0, err4
				}
				return req.Message.SequenceNumber, nil
			}
			// Subscribe to all relevant relay logs.
			_, err = sourceChain.LogPoller().RegisterFilter(logpoller.Filter{EventSigs: []common.Hash{CCIPSubSendRequested}, Addresses: []common.Address{onRamp.Address()}})
			if err != nil {
				return nil, err
			}
			onRampToReqEventSig[onRamp.Address()] = CCIPSubSendRequested
			onRampToHasher[onRamp.Address()] = NewSubscriptionLeafHasher(sourceChainId, destChainId, onRamp.Address(), hashingCtx)
		default:
			return nil, errors.Errorf("unrecognized onramp %v", onRampID)
		}
	}
	argsNoPlugin.ReportingPluginFactory = NewRelayReportingPluginFactory(lggr, sourceChain.LogPoller(), blobVerifier, onRampSeqParsers, onRampToReqEventSig, onRamps, onRampToHasher)
	oracle, err := libocr2.NewOracle(argsNoPlugin)
	if err != nil {
		return nil, err
	}
	// If this is a brand-new job, then we make use of the start blocks. If not then we're rebooting and log poller will pick up where we left off.
	if new {
		return []job.ServiceCtx{&BackfilledOracle{
			srcStartBlock: pluginConfig.SourceStartBlock,
			dstStartBlock: pluginConfig.DestStartBlock,
			src:           sourceChain.LogPoller(),
			dst:           destChain.LogPoller(),
			oracle:        job.NewServiceAdapter(oracle),
			lggr:          lggr,
		}}, nil
	}
	return []job.ServiceCtx{job.NewServiceAdapter(oracle)}, nil
}
