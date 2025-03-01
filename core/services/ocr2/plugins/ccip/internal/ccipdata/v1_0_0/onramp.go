package v1_0_0

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciptypes"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/cache"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/hashlib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/logpollerutil"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

const (
	CCIPSendRequestedEventName = "CCIPSendRequested"
	ConfigSetEventName         = "ConfigSet"
)

var _ ccipdata.OnRampReader = &OnRamp{}

type OnRamp struct {
	address                          common.Address
	onRamp                           *evm_2_evm_onramp_1_0_0.EVM2EVMOnRamp
	lp                               logpoller.LogPoller
	lggr                             logger.Logger
	client                           client.Client
	leafHasher                       ccipdata.LeafHasherInterface[[32]byte]
	sendRequestedEventSig            common.Hash
	sendRequestedSeqNumberWord       int
	filters                          []logpoller.Filter
	cachedSourcePriceRegistryAddress cache.AutoSync[cciptypes.Address]
	// Static config can be cached, because it's never expected to change.
	// The only way to change that is through the contract's constructor (redeployment)
	cachedStaticConfig cache.OnceCtxFunction[evm_2_evm_onramp_1_0_0.EVM2EVMOnRampStaticConfig]
}

func NewOnRamp(lggr logger.Logger, sourceSelector, destSelector uint64, onRampAddress common.Address, sourceLP logpoller.LogPoller, source client.Client) (*OnRamp, error) {
	onRamp, err := evm_2_evm_onramp_1_0_0.NewEVM2EVMOnRamp(onRampAddress, source)
	if err != nil {
		return nil, err
	}
	onRampABI := abihelpers.MustParseABI(evm_2_evm_onramp_1_0_0.EVM2EVMOnRampABI)
	eventSig := abihelpers.MustGetEventID(CCIPSendRequestedEventName, onRampABI)
	configSetEventSig := abihelpers.MustGetEventID(ConfigSetEventName, onRampABI)
	filters := []logpoller.Filter{
		{
			Name:      logpoller.FilterName(ccipdata.COMMIT_CCIP_SENDS, onRampAddress),
			EventSigs: []common.Hash{eventSig},
			Addresses: []common.Address{onRampAddress},
		},
		{
			Name:      logpoller.FilterName(ccipdata.CONFIG_CHANGED, onRampAddress),
			EventSigs: []common.Hash{configSetEventSig},
			Addresses: []common.Address{onRampAddress},
		},
	}
	cachedStaticConfig := cache.OnceCtxFunction[evm_2_evm_onramp_1_0_0.EVM2EVMOnRampStaticConfig](func(ctx context.Context) (evm_2_evm_onramp_1_0_0.EVM2EVMOnRampStaticConfig, error) {
		return onRamp.GetStaticConfig(&bind.CallOpts{Context: ctx})
	})
	return &OnRamp{
		lggr:       lggr,
		address:    onRampAddress,
		onRamp:     onRamp,
		client:     source,
		filters:    filters,
		lp:         sourceLP,
		leafHasher: NewLeafHasher(sourceSelector, destSelector, onRampAddress, hashlib.NewKeccakCtx(), onRamp),
		// offset || sourceChainID || seqNum || ...
		sendRequestedSeqNumberWord: 2,
		sendRequestedEventSig:      eventSig,
		cachedSourcePriceRegistryAddress: cache.NewLogpollerEventsBased[cciptypes.Address](
			sourceLP,
			[]common.Hash{configSetEventSig},
			onRampAddress,
		),
		cachedStaticConfig: cachedStaticConfig,
	}, nil
}

func (o *OnRamp) Address() (cciptypes.Address, error) {
	return cciptypes.Address(o.onRamp.Address().String()), nil
}

func (o *OnRamp) GetDynamicConfig() (cciptypes.OnRampDynamicConfig, error) {
	if o.onRamp == nil {
		return cciptypes.OnRampDynamicConfig{}, fmt.Errorf("onramp not initialized")
	}
	legacyDynamicConfig, err := o.onRamp.GetDynamicConfig(nil)
	if err != nil {
		return cciptypes.OnRampDynamicConfig{}, err
	}
	return cciptypes.OnRampDynamicConfig{
		Router:                            cciptypes.Address(legacyDynamicConfig.Router.String()),
		MaxNumberOfTokensPerMsg:           legacyDynamicConfig.MaxTokensLength,
		DestGasOverhead:                   0,
		DestGasPerPayloadByte:             0,
		DestDataAvailabilityOverheadGas:   0,
		DestGasPerDataAvailabilityByte:    0,
		DestDataAvailabilityMultiplierBps: 0,
		PriceRegistry:                     cciptypes.Address(legacyDynamicConfig.PriceRegistry.String()),
		MaxDataBytes:                      legacyDynamicConfig.MaxDataSize,
		MaxPerMsgGasLimit:                 uint32(legacyDynamicConfig.MaxGasLimit),
	}, nil
}

func (o *OnRamp) SourcePriceRegistryAddress(ctx context.Context) (cciptypes.Address, error) {
	return o.cachedSourcePriceRegistryAddress.Get(ctx, func(ctx context.Context) (cciptypes.Address, error) {
		c, err := o.GetDynamicConfig()
		if err != nil {
			return "", err
		}
		return c.PriceRegistry, nil
	})
}

func (o *OnRamp) GetSendRequestsBetweenSeqNums(ctx context.Context, seqNumMin, seqNumMax uint64, finalized bool) ([]cciptypes.EVM2EVMMessageWithTxMeta, error) {
	logs, err := o.lp.LogsDataWordRange(
		o.sendRequestedEventSig,
		o.address,
		o.sendRequestedSeqNumberWord,
		logpoller.EvmWord(seqNumMin),
		logpoller.EvmWord(seqNumMax),
		ccipdata.LogsConfirmations(finalized),
		pg.WithParentCtx(ctx))
	if err != nil {
		return nil, err
	}

	parsedLogs, err := ccipdata.ParseLogs[cciptypes.EVM2EVMMessage](logs, o.lggr, o.logToMessage)
	if err != nil {
		return nil, err
	}

	res := make([]cciptypes.EVM2EVMMessageWithTxMeta, 0, len(parsedLogs))
	for _, log := range parsedLogs {
		res = append(res, cciptypes.EVM2EVMMessageWithTxMeta{
			TxMeta:         log.TxMeta,
			EVM2EVMMessage: log.Data,
		})
	}
	return res, nil
}

func (o *OnRamp) RouterAddress() (cciptypes.Address, error) {
	config, err := o.onRamp.GetDynamicConfig(nil)
	if err != nil {
		return "", err
	}
	return cciptypes.Address(config.Router.String()), nil
}

func (o *OnRamp) IsSourceChainHealthy(context.Context) (bool, error) {
	if err := o.lp.Healthy(); err != nil {
		return false, nil
	}
	return true, nil
}

func (o *OnRamp) IsSourceCursed(ctx context.Context) (bool, error) {
	staticConfig, err := o.cachedStaticConfig(ctx)
	if err != nil {
		return false, err
	}

	arm, err := arm_contract.NewARMContract(staticConfig.ArmProxy, o.client)
	if err != nil {
		return false, fmt.Errorf("intializing Arm contract through the ArmProxy: %w", err)
	}

	cursed, err := arm.IsCursed(&bind.CallOpts{Context: ctx})
	if err != nil {
		return false, fmt.Errorf("checking if source Arm is cursed: %w", err)
	}
	return cursed, nil
}

func (o *OnRamp) GetLastUSDCMessagePriorToLogIndexInTx(ctx context.Context, logIndex int64, txHash common.Hash) ([]byte, error) {
	return nil, errors.New("USDC not supported in < 1.2.0")
}

func (o *OnRamp) Close(qopts ...pg.QOpt) error {
	return logpollerutil.UnregisterLpFilters(o.lp, o.filters, qopts...)
}

func (o *OnRamp) RegisterFilters(qopts ...pg.QOpt) error {
	return logpollerutil.RegisterLpFilters(o.lp, o.filters, qopts...)
}

func (o *OnRamp) logToMessage(log types.Log) (*cciptypes.EVM2EVMMessage, error) {
	msg, err := o.onRamp.ParseCCIPSendRequested(log)
	if err != nil {
		return nil, err
	}
	h, err := o.leafHasher.HashLeaf(log)
	if err != nil {
		return nil, err
	}
	tokensAndAmounts := make([]cciptypes.TokenAmount, len(msg.Message.TokenAmounts))
	for i, tokenAndAmount := range msg.Message.TokenAmounts {
		tokensAndAmounts[i] = cciptypes.TokenAmount{
			Token:  cciptypes.Address(tokenAndAmount.Token.String()),
			Amount: tokenAndAmount.Amount,
		}
	}
	return &cciptypes.EVM2EVMMessage{
		SequenceNumber:      msg.Message.SequenceNumber,
		GasLimit:            msg.Message.GasLimit,
		Nonce:               msg.Message.Nonce,
		MessageID:           msg.Message.MessageId,
		SourceChainSelector: msg.Message.SourceChainSelector,
		Sender:              cciptypes.Address(msg.Message.Sender.String()),
		Receiver:            cciptypes.Address(msg.Message.Receiver.String()),
		Strict:              msg.Message.Strict,
		FeeToken:            cciptypes.Address(msg.Message.FeeToken.String()),
		FeeTokenAmount:      msg.Message.FeeTokenAmount,
		Data:                msg.Message.Data,
		TokenAmounts:        tokensAndAmounts,
		SourceTokenData:     make([][]byte, len(msg.Message.TokenAmounts)), // Always empty in 1.0
		Hash:                h,
	}, nil
}
