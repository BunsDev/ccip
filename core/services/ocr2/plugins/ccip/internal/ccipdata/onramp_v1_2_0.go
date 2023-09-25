package ccipdata

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/hashlib"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

var (
	// Backwards compat for integration tests
	CCIPSendRequestEventSigV1_2_0 common.Hash
)

const (
	CCIPSendRequestSeqNumIndexV1_2_0 = 4
)

func init() {
	onRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_onramp.EVM2EVMOnRampABI))
	if err != nil {
		panic(err)
	}
	CCIPSendRequestEventSigV1_2_0 = abihelpers.GetIDOrPanic("CCIPSendRequested", onRampABI)
}

type LeafHasherV1_2_0 struct {
	metaDataHash [32]byte
	ctx          hashlib.Ctx[[32]byte]
	onRamp       *evm_2_evm_onramp.EVM2EVMOnRamp
}

func NewLeafHasherV1_2_0(sourceChainSelector uint64, destChainSelector uint64, onRampId common.Address, ctx hashlib.Ctx[[32]byte], onRamp *evm_2_evm_onramp.EVM2EVMOnRamp) *LeafHasherV1_2_0 {
	return &LeafHasherV1_2_0{
		metaDataHash: getMetaDataHash(ctx, ctx.Hash([]byte("EVM2EVMMessageHashV2")), sourceChainSelector, onRampId, destChainSelector),
		ctx:          ctx,
		onRamp:       onRamp,
	}
}

func (t *LeafHasherV1_2_0) HashLeaf(log types.Log) ([32]byte, error) {
	msg, err := t.onRamp.ParseCCIPSendRequested(log)
	if err != nil {
		return [32]byte{}, err
	}

	encodedTokens, err := abihelpers.TokenAmountsArgs.PackValues([]interface{}{msg.Message.TokenAmounts})
	if err != nil {
		return [32]byte{}, err
	}

	bytesArray, err := abi.NewType("bytes[]", "bytes[]", nil)
	if err != nil {
		return [32]byte{}, err
	}

	encodedSourceTokenData, err := abi.Arguments{abi.Argument{Type: bytesArray}}.PackValues([]interface{}{msg.Message.SourceTokenData})
	if err != nil {
		return [32]byte{}, err
	}

	packedValues, err := utils.ABIEncode(
		`[
{"name": "leafDomainSeparator","type":"bytes1"},
{"name": "metadataHash", "type":"bytes32"},
{"name": "sequenceNumber", "type":"uint64"},
{"name": "nonce", "type":"uint64"},
{"name": "sender", "type":"address"},
{"name": "receiver", "type":"address"},
{"name": "dataHash", "type":"bytes32"},
{"name": "tokenAmountsHash", "type":"bytes32"},
{"name": "sourceTokenDataHash", "type":"bytes32"},
{"name": "gasLimit", "type":"uint256"},
{"name": "strict", "type":"bool"},
{"name": "feeToken","type": "address"},
{"name": "feeTokenAmount","type": "uint256"}
]`,
		leafDomainSeparator,
		t.metaDataHash,
		msg.Message.SequenceNumber,
		msg.Message.Nonce,
		msg.Message.Sender,
		msg.Message.Receiver,
		t.ctx.Hash(msg.Message.Data),
		t.ctx.Hash(encodedTokens),
		t.ctx.Hash(encodedSourceTokenData),
		msg.Message.GasLimit,
		msg.Message.Strict,
		msg.Message.FeeToken,
		msg.Message.FeeTokenAmount,
	)
	if err != nil {
		return [32]byte{}, err
	}
	return t.ctx.Hash(packedValues), nil
}

var _ OnRampReader = &OnRampV1_2_0{}

// Significant change in 1.2:
// - CCIPSendRequested event signature has changed
type OnRampV1_2_0 struct {
	onRamp                     *evm_2_evm_onramp.EVM2EVMOnRamp
	address                    common.Address
	lggr                       logger.Logger
	lp                         logpoller.LogPoller
	leafHasher                 LeafHasherInterface[[32]byte]
	client                     client.Client
	finalityTags               bool
	filterName                 string
	sendRequestedEventSig      common.Hash
	sendRequestedSeqNumberWord int
}

func (o *OnRampV1_2_0) logToMessage(log types.Log) (*EVM2EVMMessage, error) {
	msg, err := o.onRamp.ParseCCIPSendRequested(log)
	if err != nil {
		return nil, err
	}
	h, err := o.leafHasher.HashLeaf(log)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMessage{
		SequenceNumber: msg.Message.SequenceNumber,
		GasLimit:       msg.Message.GasLimit,
		Nonce:          msg.Message.Nonce,
		Hash:           h,
		Log:            log,
	}, nil
}

func (o *OnRampV1_2_0) GetSendRequestsGteSeqNum(ctx context.Context, seqNum uint64, confs int) ([]Event[EVM2EVMMessage], error) {
	if !o.finalityTags {
		logs, err2 := o.lp.LogsDataWordGreaterThan(
			o.sendRequestedEventSig,
			o.address,
			o.sendRequestedSeqNumberWord,
			abihelpers.EvmWord(seqNum),
			confs,
			pg.WithParentCtx(ctx),
		)
		if err2 != nil {
			return nil, fmt.Errorf("logs data word greater than: %w", err2)
		}
		return parseLogs[EVM2EVMMessage](logs, o.lggr, o.logToMessage)
	}
	latestFinalizedHash, err := latestFinalizedBlockHash(ctx, o.client)
	if err != nil {
		return nil, err
	}
	logs, err := o.lp.LogsUntilBlockHashDataWordGreaterThan(
		o.sendRequestedEventSig,
		o.address,
		o.sendRequestedSeqNumberWord,
		abihelpers.EvmWord(seqNum),
		latestFinalizedHash,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("logs until block hash data word greater than: %w", err)
	}
	return parseLogs[EVM2EVMMessage](logs, o.lggr, o.logToMessage)
}

func (o *OnRampV1_2_0) GetSendRequestsBetweenSeqNums(ctx context.Context, seqNumMin, seqNumMax uint64, confs int) ([]Event[EVM2EVMMessage], error) {
	logs, err := o.lp.LogsDataWordRange(
		o.sendRequestedEventSig,
		o.address,
		o.sendRequestedSeqNumberWord,
		logpoller.EvmWord(seqNumMin),
		logpoller.EvmWord(seqNumMax),
		confs,
		pg.WithParentCtx(ctx))
	if err != nil {
		return nil, err
	}
	return parseLogs[EVM2EVMMessage](logs, o.lggr, o.logToMessage)
}

func (o *OnRampV1_2_0) Router() common.Address {
	config, _ := o.onRamp.GetDynamicConfig(nil)
	return config.Router
}

func (o *OnRampV1_2_0) ToOffRampMessage(message EVM2EVMMessage) (*evm_2_evm_offramp.InternalEVM2EVMMessage, error) {
	m, err := o.onRamp.ParseCCIPSendRequested(message.Log)
	if err != nil {
		return nil, err
	}
	tokensAndAmounts := make([]evm_2_evm_offramp.ClientEVMTokenAmount, len(m.Message.TokenAmounts))
	for i, tokenAndAmount := range m.Message.TokenAmounts {
		tokensAndAmounts[i] = evm_2_evm_offramp.ClientEVMTokenAmount{
			Token:  tokenAndAmount.Token,
			Amount: tokenAndAmount.Amount,
		}
	}
	return &evm_2_evm_offramp.InternalEVM2EVMMessage{
		SourceChainSelector: m.Message.SourceChainSelector,
		Sender:              m.Message.Sender,
		Receiver:            m.Message.Receiver,
		SequenceNumber:      m.Message.SequenceNumber,
		GasLimit:            m.Message.GasLimit,
		Strict:              m.Message.Strict,
		Nonce:               m.Message.Nonce,
		FeeToken:            m.Message.FeeToken,
		FeeTokenAmount:      m.Message.FeeTokenAmount,
		Data:                m.Message.Data,
		TokenAmounts:        tokensAndAmounts,
		SourceTokenData:     m.Message.SourceTokenData, // BREAKING CHANGE IN 1.2
		MessageId:           m.Message.MessageId,
	}, nil
}

func (o *OnRampV1_2_0) Close() error {
	return o.lp.UnregisterFilter(o.filterName)
}

func NewOnRampV1_2_0(
	lggr logger.Logger,
	sourceSelector,
	destSelector uint64,
	onRampAddress common.Address,
	sourceLP logpoller.LogPoller,
	source client.Client,
	finalityTags bool,
) (*OnRampV1_2_0, error) {
	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddress, source)
	if err != nil {
		panic(err) // ABI failure ok to panic
	}
	onRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_onramp.EVM2EVMOnRampABI))
	if err != nil {
		return nil, err
	}
	// Subscribe to the relevant logs
	// Note we can keep the same prefix across 1.0/1.1 and 1.2 because the onramp addresses will be different
	name := logpoller.FilterName(COMMIT_CCIP_SENDS, onRampAddress)
	err = sourceLP.RegisterFilter(logpoller.Filter{
		Name:      name,
		EventSigs: []common.Hash{abihelpers.GetIDOrPanic("CCIPSendRequested", onRampABI)},
		Addresses: []common.Address{onRampAddress},
	})
	return &OnRampV1_2_0{
		finalityTags: finalityTags,
		lggr:         lggr,
		client:       source,
		lp:           sourceLP,
		leafHasher:   NewLeafHasherV1_2_0(sourceSelector, destSelector, onRampAddress, hashlib.NewKeccakCtx(), onRamp),
		onRamp:       onRamp,
		filterName:   name,
	}, nil
}
