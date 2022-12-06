package ccip

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/core/utils/mathutil"
)

const (
	MessageStateUntouched = iota
	MessageStateInProgress
	MessageStateSuccess
	MessageStateFailure
)

const (
	BatchGasLimit            = 5_000_000                 // TODO: think if a good value for this
	GasLimitPerTx            = BatchGasLimit - 1_000_000 // Leave a buffer for overhead.
	MaxPayloadLength         = 1000
	MaxTokensPerMessage      = 5
	MaxExecutionReportLength = 150_000 // TODO
	MaxGasPrice              = 200e9   // 200 gwei. TODO: probably want this to be some dynamic value, a multiplier of the current gas price.
	TokenPriceBufferPercent  = 10      // Amount that the leader adds as a token price buffer in Query.
)

var (
	_                    types.ReportingPluginFactory = &ExecutionReportingPluginFactory{}
	_                    types.ReportingPlugin        = &ExecutionReportingPlugin{}
	ErrCommitStoreIsDown                              = errors.New("commitStore is down")
)

func EncodeTollExecutionReport(seqNums []uint64,
	tokensPerFeeCoin map[common.Address]*big.Int,
	msgs [][]byte,
	innerProofs [][32]byte,
	innerProofSourceFlags []bool,
	outerProofs [][32]byte,
	outerProofSourceFlags []bool,
) (types.Report, error) {
	var tokensPerFeeCoinAddresses []common.Address
	var tokensPerFeeCoinValues []*big.Int
	for addr := range tokensPerFeeCoin {
		tokensPerFeeCoinAddresses = append(tokensPerFeeCoinAddresses, addr)
	}
	// Sort the addresses for determinism.
	sort.Slice(tokensPerFeeCoinAddresses, func(i, j int) bool {
		return bytes.Compare(tokensPerFeeCoinAddresses[i].Bytes(), tokensPerFeeCoinAddresses[j].Bytes()) < 0
	})
	for _, addr := range tokensPerFeeCoinAddresses {
		tokensPerFeeCoinValues = append(tokensPerFeeCoinValues, tokensPerFeeCoin[addr])
	}
	report, err := makeTollExecutionReportArgs().PackValues([]interface{}{&evm_2_evm_toll_offramp.TollExecutionReport{
		SequenceNumbers:          seqNums,
		EncodedMessages:          msgs,
		TokenPerFeeCoinAddresses: tokensPerFeeCoinAddresses,
		TokenPerFeeCoin:          tokensPerFeeCoinValues,
		InnerProofs:              innerProofs,
		InnerProofFlagBits:       ProofFlagsToBits(innerProofSourceFlags),
		OuterProofs:              outerProofs,
		OuterProofFlagBits:       ProofFlagsToBits(outerProofSourceFlags),
	}})
	if err != nil {
		return nil, err
	}
	return report, nil
}

func EncodeGEExecutionReport(seqNums []uint64,
	tokensPerFeeCoin map[common.Address]*big.Int,
	msgs [][]byte,
	innerProofs [][32]byte,
	innerProofSourceFlags []bool,
	outerProofs [][32]byte,
	outerProofSourceFlags []bool,
	feeUpdates []evm_2_evm_ge_offramp.GEFeeUpdate,
) (types.Report, error) {
	var tokensPerFeeCoinAddresses []common.Address
	var tokensPerFeeCoinValues []*big.Int
	for addr := range tokensPerFeeCoin {
		tokensPerFeeCoinAddresses = append(tokensPerFeeCoinAddresses, addr)
	}
	// Sort the addresses for determinism.
	sort.Slice(tokensPerFeeCoinAddresses, func(i, j int) bool {
		return bytes.Compare(tokensPerFeeCoinAddresses[i].Bytes(), tokensPerFeeCoinAddresses[j].Bytes()) < 0
	})
	for _, addr := range tokensPerFeeCoinAddresses {
		tokensPerFeeCoinValues = append(tokensPerFeeCoinValues, tokensPerFeeCoin[addr])
	}
	report, err := makeExecutionReportArgs().PackValues([]interface{}{&evm_2_evm_ge_offramp.GEExecutionReport{
		SequenceNumbers:          seqNums,
		FeeUpdates:               feeUpdates,
		EncodedMessages:          msgs,
		TokenPerFeeCoinAddresses: tokensPerFeeCoinAddresses,
		TokenPerFeeCoin:          tokensPerFeeCoinValues,
		InnerProofs:              innerProofs,
		InnerProofFlagBits:       ProofFlagsToBits(innerProofSourceFlags),
		OuterProofs:              outerProofs,
		OuterProofFlagBits:       ProofFlagsToBits(outerProofSourceFlags),
	}})
	if err != nil {
		return nil, err
	}
	return report, nil
}

func MessagesFromTollExecutionReport(report types.Report) ([]uint64, [][]byte, error) {
	tollReport, err := DecodeTollExecutionReport(report)
	if err != nil {
		return nil, nil, err
	}
	return tollReport.SequenceNumbers, tollReport.EncodedMessages, nil
}

func MessagesFromGEExecutionReport(report types.Report) ([]uint64, [][]byte, error) {
	geReport, err := DecodeGEExecutionReport(report)
	if err != nil {
		return nil, nil, err
	}
	return geReport.SequenceNumbers, geReport.EncodedMessages, nil
}

func DecodeGEExecutionReport(report types.Report) (*evm_2_evm_ge_offramp.GEExecutionReport, error) {
	unpacked, err := makeExecutionReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, errors.New("assumptionViolation: expected at least one element")
	}

	// Must be anonymous struct here
	erStruct, ok := unpacked[0].(struct {
		SequenceNumbers          []uint64         `json:"sequenceNumbers"`
		TokenPerFeeCoinAddresses []common.Address `json:"tokenPerFeeCoinAddresses"`
		TokenPerFeeCoin          []*big.Int       `json:"tokenPerFeeCoin"`
		FeeUpdates               []struct {
			ChainId        *big.Int `json:"chainId"`
			LinkPerUnitGas *big.Int `json:"linkPerUnitGas"`
		} `json:"feeUpdates"`
		EncodedMessages    [][]byte    `json:"encodedMessages"`
		InnerProofs        [][32]uint8 `json:"innerProofs"`
		InnerProofFlagBits *big.Int    `json:"innerProofFlagBits"`
		OuterProofs        [][32]uint8 `json:"outerProofs"`
		OuterProofFlagBits *big.Int    `json:"outerProofFlagBits"`
	})
	if !ok {
		return nil, fmt.Errorf("got %T", unpacked[0])
	}
	var er evm_2_evm_ge_offramp.GEExecutionReport
	er.EncodedMessages = append(er.EncodedMessages, erStruct.EncodedMessages...)
	er.InnerProofs = append(er.InnerProofs, erStruct.InnerProofs...)
	er.OuterProofs = append(er.OuterProofs, erStruct.OuterProofs...)

	er.FeeUpdates = []evm_2_evm_ge_offramp.GEFeeUpdate{}

	for _, feeUpdate := range erStruct.FeeUpdates {
		er.FeeUpdates = append(er.FeeUpdates, evm_2_evm_ge_offramp.GEFeeUpdate{
			ChainId:        feeUpdate.ChainId,
			LinkPerUnitGas: feeUpdate.LinkPerUnitGas,
		})
	}

	er.SequenceNumbers = erStruct.SequenceNumbers
	// Unpack will populate with big.Int{false, <allocated empty nat>} for 0 values,
	// which is different from the expected big.NewInt(0). Rebuild to the expected value for this case.
	er.InnerProofFlagBits = big.NewInt(erStruct.InnerProofFlagBits.Int64())
	er.OuterProofFlagBits = big.NewInt(erStruct.OuterProofFlagBits.Int64())
	er.TokenPerFeeCoinAddresses = erStruct.TokenPerFeeCoinAddresses
	er.TokenPerFeeCoin = erStruct.TokenPerFeeCoin
	return &er, nil
}

func DecodeTollExecutionReport(report types.Report) (*evm_2_evm_toll_offramp.TollExecutionReport, error) {
	unpacked, err := makeTollExecutionReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, errors.New("assumptionViolation: expected at least one element")
	}

	// Must be anonymous struct here
	erStruct, ok := unpacked[0].(struct {
		SequenceNumbers          []uint64         `json:"sequenceNumbers"`
		TokenPerFeeCoinAddresses []common.Address `json:"tokenPerFeeCoinAddresses"`
		TokenPerFeeCoin          []*big.Int       `json:"tokenPerFeeCoin"`
		EncodedMessages          [][]byte         `json:"encodedMessages"`
		InnerProofs              [][32]uint8      `json:"innerProofs"`
		InnerProofFlagBits       *big.Int         `json:"innerProofFlagBits"`
		OuterProofs              [][32]uint8      `json:"outerProofs"`
		OuterProofFlagBits       *big.Int         `json:"outerProofFlagBits"`
	})
	if !ok {
		return nil, fmt.Errorf("got %T", unpacked[0])
	}
	var er evm_2_evm_toll_offramp.TollExecutionReport
	er.EncodedMessages = append(er.EncodedMessages, erStruct.EncodedMessages...)
	er.InnerProofs = append(er.InnerProofs, erStruct.InnerProofs...)
	er.OuterProofs = append(er.OuterProofs, erStruct.OuterProofs...)
	er.SequenceNumbers = erStruct.SequenceNumbers
	// Unpack will populate with big.Int{false, <allocated empty nat>} for 0 values,
	// which is different from the expected big.NewInt(0). Rebuild to the expected value for this case.
	er.InnerProofFlagBits = big.NewInt(erStruct.InnerProofFlagBits.Int64())
	er.OuterProofFlagBits = big.NewInt(erStruct.OuterProofFlagBits.Int64())
	er.TokenPerFeeCoinAddresses = erStruct.TokenPerFeeCoinAddresses
	er.TokenPerFeeCoin = erStruct.TokenPerFeeCoin
	return &er, nil
}

func aggregateTokenValue(tokenLimitPrices map[common.Address]*big.Int, srcToDst map[common.Address]common.Address, tokens []common.Address, amounts []*big.Int) (*big.Int, error) {
	sum := big.NewInt(0)
	for i := 0; i < len(tokens); i++ {
		price, ok := tokenLimitPrices[srcToDst[tokens[i]]]
		if !ok {
			return nil, errors.Errorf("do not have price for src token %x", tokens[i])
		}
		sum.Add(sum, new(big.Int).Mul(price, amounts[i]))
	}
	return sum, nil
}

// EventSignatures contain pluginType specific signatures and indexes.
// Indexes are zero indexed
type EventSignatures struct {
	SendRequested                            common.Hash
	SendRequestedSequenceNumberIndex         int
	ExecutionStateChanged                    common.Hash
	ExecutionStateChangedSequenceNumberIndex int
}

type ExecutionReportingPluginFactory struct {
	lggr                                 logger.Logger
	source, dest                         logpoller.LogPoller
	onRamp, offRampAddr                  common.Address
	offRamp                              OffRamp
	commitStore                          *commit_store.CommitStore
	builder                              BatchBuilder
	onRampSeqParser                      func(log logpoller.Log) (uint64, error)
	eventSignatures                      EventSignatures
	priceGetter                          PriceGetter
	onRampToHasher                       map[common.Address]LeafHasher[[32]byte]
	rootSnoozeTime                       time.Duration
	inflightCacheExpiry                  time.Duration
	sourceGasEstimator, destGasEstimator gas.Estimator
	sourceChainID                        *big.Int
}

func NewExecutionReportingPluginFactory(
	lggr logger.Logger,
	onRamp common.Address,
	commitStore *commit_store.CommitStore,
	source, dest logpoller.LogPoller,
	offRampAddr common.Address,
	offRamp OffRamp,
	builder BatchBuilder,
	onRampSeqParser func(log logpoller.Log) (uint64, error),
	eventSignatures EventSignatures,
	priceGetter PriceGetter,
	onRampToHasher map[common.Address]LeafHasher[[32]byte],
	rootSnoozeTime time.Duration,
	inflightCacheExpiry time.Duration,
	sourceGasEstimator, destGasEstimator gas.Estimator,
	sourceChainID *big.Int,
) types.ReportingPluginFactory {
	return &ExecutionReportingPluginFactory{lggr: lggr, onRamp: onRamp, commitStore: commitStore, offRamp: offRamp, source: source, dest: dest, offRampAddr: offRampAddr, builder: builder,
		onRampSeqParser: onRampSeqParser, eventSignatures: eventSignatures, priceGetter: priceGetter, onRampToHasher: onRampToHasher, rootSnoozeTime: rootSnoozeTime, inflightCacheExpiry: inflightCacheExpiry,
		sourceGasEstimator: sourceGasEstimator, destGasEstimator: destGasEstimator, sourceChainID: sourceChainID}
}

func (rf *ExecutionReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	offchainConfig, err := Decode(config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	return &ExecutionReportingPlugin{
			lggr:           rf.lggr.Named("ExecutionReportingPlugin"),
			F:              config.F,
			offRampAddr:    rf.offRampAddr,
			offRamp:        rf.offRamp,
			onRamp:         rf.onRamp,
			commitStore:    rf.commitStore,
			source:         rf.source,
			dest:           rf.dest,
			offchainConfig: offchainConfig,
			builder: NewExecutionBatchBuilder(
				BatchGasLimit,
				rf.rootSnoozeTime,
				rf.commitStore,
				rf.onRamp,
				rf.offRampAddr,
				rf.source,
				rf.dest,
				rf.builder,
				offchainConfig,
				rf.offRamp,
				rf.eventSignatures,
				rf.lggr),
			onRampSeqParser:     rf.onRampSeqParser,
			eventSignatures:     rf.eventSignatures,
			priceGetter:         rf.priceGetter,
			onRampToHasher:      rf.onRampToHasher,
			inflightCacheExpiry: rf.inflightCacheExpiry,
			sourceGasEstimator:  rf.sourceGasEstimator,
			destGasEstimator:    rf.destGasEstimator,
			sourceChainID:       rf.sourceChainID,
		}, types.ReportingPluginInfo{
			Name:          "CCIPExecution",
			UniqueReports: true,
			Limits: types.ReportingPluginLimits{
				MaxQueryLength:       MaxQueryLength,
				MaxObservationLength: MaxObservationLength,
				MaxReportLength:      MaxExecutionReportLength,
			},
		}, nil
}

type ExecutionReportingPlugin struct {
	lggr         logger.Logger
	F            int
	offRampAddr  common.Address
	onRamp       common.Address
	offRamp      OffRamp
	commitStore  *commit_store.CommitStore
	source, dest logpoller.LogPoller
	// We need to synchronize access to the inflight structure
	// as reporting plugin methods may be called from separate goroutines,
	// e.g. reporting vs transmission protocol.
	inFlightMu                           sync.RWMutex
	inFlight                             []InflightExecutionReport
	inflightCacheExpiry                  time.Duration
	offchainConfig                       OffchainConfig
	builder                              *ExecutionBatchBuilder
	onRampSeqParser                      func(log logpoller.Log) (uint64, error)
	eventSignatures                      EventSignatures
	priceGetter                          PriceGetter
	onRampToHasher                       map[common.Address]LeafHasher[[32]byte]
	sourceGasEstimator, destGasEstimator gas.Estimator
	sourceChainID                        *big.Int
}

type InflightExecutionReport struct {
	createdAt   time.Time
	seqNrs      []uint64
	encMessages [][]byte
}

type Query struct {
	TokenPrices  map[common.Address]*big.Int `json:"tokenPrices"` // TODO: We should simplify this to just link for toll as well.
	DestGasPrice *big.Int                    `json:"destGasPrice"`
}

// expect percentMultiplier to be [0, 100]
func (r *ExecutionReportingPlugin) tokenPrices(percentMultiplier *big.Int) (map[common.Address]*big.Int, error) {
	tokensPerFeeCoin := make(map[common.Address]*big.Int)
	executionFeeTokens, err := r.offRamp.GetSupportedTokensForExecutionFee()
	if err != nil {
		return nil, err
	}
	prices, err := r.priceGetter.TokensPerFeeCoin(context.Background(), executionFeeTokens)
	if err != nil {
		return nil, err
	}
	for token, price := range prices {
		buffer := big.NewInt(0).Div(price, percentMultiplier)
		tokensPerFeeCoin[token] = big.NewInt(0).Add(price, buffer)
	}
	return tokensPerFeeCoin, nil
}

func (r *ExecutionReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	// The leader queries an overestimated set of token prices, which are used by all the followers
	// to compute message executability, ensuring that the set of executable messages is deterministic.
	tokensPerFeeCoin, err := r.tokenPrices(big.NewInt(TokenPriceBufferPercent))
	if err != nil {
		return nil, err
	}
	// In the context of GE CCIP, latency is much less important than cost, so generally we'd prefer to wait vs bump at all.
	// Options:
	// - Disable bumping entirely. Execute messages up to max loss given current price and simply wait until it is included. Means that we can potentially
	// block execution for all jobs until that tx goes though.
	// - Have very very relaxed bumping settings and a relatively low cap (say 1k gwei). The worst case is that we decide to execute a max loss tx,
	// prices remain high for a long time to invoke bumping and increase our loss up to the bumped cap. Benefit is we will unblock the jobs ourselves.
	// Should be possible to ensure the max bumped loss is incurred with some extremely low probability (something much worse than a luna type meltdown of 4hr 8k gwei spike).
	// TODO: Switch between 1559 and non-1559 here based on chain (or wrap estimator at a higher level).
	destGasPrice, _, err := r.destGasEstimator.GetLegacyGas(ctx, nil, BatchGasLimit, assets.NewWei(big.NewInt(int64(MaxGasPrice))))
	if err != nil {
		return nil, err
	}
	return json.Marshal(Query{TokenPrices: tokensPerFeeCoin, DestGasPrice: destGasPrice.ToInt()})
}

func (r *ExecutionReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	// Query contains the tokenPricesPerFeeCoin
	lggr := r.lggr.Named("ExecutionObservation")
	if isCommitStoreDownNow(lggr, r.commitStore) {
		return nil, ErrCommitStoreIsDown
	}
	// Expire any inflight reports.
	r.expireInflight(lggr)

	var q Query
	if err := json.Unmarshal(query, &q); err != nil {
		return nil, err
	}
	// Read and make a copy for the builder.
	r.inFlightMu.RLock()
	inFlight := make([]InflightExecutionReport, len(r.inFlight))
	copy(inFlight[:], r.inFlight[:])
	r.inFlightMu.RUnlock()

	batchBuilderStart := time.Now()
	// IMPORTANT: We build executable set based on the leaders token prices, ensuring consistency across followers.
	executableSequenceNumbers, err := r.builder.getExecutableSeqNrs(q.DestGasPrice, q.TokenPrices, inFlight)
	lggr.Infof("Batch building took %d ms", time.Since(batchBuilderStart).Milliseconds())
	if err != nil {
		return nil, err
	}
	// Toll requires seqnrs
	_, isToll := r.offRamp.(*tollOffRamp)
	if isToll && len(executableSequenceNumbers) == 0 {
		return nil, errors.New("No observations")
	}
	lggr.Infof("executable seq nums %v %x", executableSequenceNumbers, r.eventSignatures.SendRequested)
	followerTokensPerFeeCoin, err := r.tokenPrices(big.NewInt(TokenPriceBufferPercent))
	if err != nil {
		return nil, err
	}
	// Observe a source chain price for GE pricing.
	// TODO: 1559 support
	sourceGasPrice, _, err := r.sourceGasEstimator.GetLegacyGas(ctx, nil, BatchGasLimit, assets.NewWei(big.NewInt(int64(MaxGasPrice))))
	if err != nil {
		return nil, err
	}
	return ExecutionObservation{
		SeqNrs:           executableSequenceNumbers, // Note can be empty
		TokensPerFeeCoin: followerTokensPerFeeCoin,
		SourceGasPrice:   sourceGasPrice.ToInt(),
	}.Marshal()
}

func contiguousReqs(lggr logger.Logger, min, max uint64, seqNrs []uint64) bool {
	for i, j := min, 0; i < max && j < len(seqNrs); i, j = i+1, j+1 {
		if seqNrs[j] != i {
			lggr.Errorw("unexpected gap in seq nums", "seq", i)
			return false
		}
	}
	return true
}

func leafsFromIntervals(lggr logger.Logger, onRampToEventSig map[common.Address]EventSignatures, seqParsers map[common.Address]func(logpoller.Log) (uint64, error), intervalByOnRamp map[common.Address]commit_store.InternalInterval, srcLogPoller logpoller.LogPoller, onRampToHasher map[common.Address]LeafHasher[[32]byte], confs int) (map[common.Address][][32]byte, error) {
	leafsByOnRamp := make(map[common.Address][][32]byte)
	for onRamp, interval := range intervalByOnRamp {
		// Logs are guaranteed to be in order of seq num, since these are finalized logs only
		// and the contract's seq num is auto-incrementing.
		logs, err := srcLogPoller.LogsDataWordRange(
			onRampToEventSig[onRamp].SendRequested,
			onRamp,
			onRampToEventSig[onRamp].SendRequestedSequenceNumberIndex,
			logpoller.EvmWord(interval.Min),
			logpoller.EvmWord(interval.Max),
			confs)
		if err != nil {
			return nil, err
		}
		var seqNrs []uint64
		for _, log := range logs {
			seqNr, err2 := seqParsers[onRamp](log)
			if err2 != nil {
				return nil, err2
			}
			seqNrs = append(seqNrs, seqNr)
		}
		if !contiguousReqs(lggr, interval.Min, interval.Max, seqNrs) {
			return nil, errors.Errorf("do not have full range [%v, %v] have %v", interval.Min, interval.Max, seqNrs)
		}
		var leafs [][32]byte
		for _, log := range logs {
			hash, err2 := onRampToHasher[onRamp].HashLeaf(LogPollerLogToEthLog(log))
			if err2 != nil {
				return nil, err2
			}
			leafs = append(leafs, hash)
		}
		leafsByOnRamp[onRamp] = leafs
	}
	return leafsByOnRamp, nil
}

func (r *ExecutionReportingPlugin) generateFeeUpdate(sourceGasPrice *big.Int, juelsPerFeeCoin *big.Int) []evm_2_evm_ge_offramp.GEFeeUpdate {
	// TODO: Check gas fee updated logs
	linkPerUnitGas := big.NewInt(0).Div(big.NewInt(0).Mul(sourceGasPrice, juelsPerFeeCoin), big.NewInt(1e18))
	return []evm_2_evm_ge_offramp.GEFeeUpdate{
		{
			ChainId: r.sourceChainID,
			// (juels/eth) * (wei / gas) / (1 eth / 1e18 wei) = juels/gas
			// TODO: Think more about this offchain/onchain computation split
			LinkPerUnitGas: linkPerUnitGas,
		},
	}
}

type MessageExecution struct {
	encMsgs               [][]byte
	innerProofs           [][32]byte
	innerProofSourceFlags []bool
	outerProofs           [][32]byte
	outerProofSourceFlags []bool
}

// Assumes non-empty report. Messages to execute can span more than one report, but are assumed to be in order of increasing
// sequence number.
func (r *ExecutionReportingPlugin) buildReport(lggr logger.Logger, finalSeqNums []uint64, tokensPerFeeCoin map[common.Address]*big.Int, sourceGasPrice *big.Int) ([]byte, error) {
	execTokens, err := r.offRamp.GetSupportedTokensForExecutionFee()
	if err != nil {
		return nil, err
	}
	// TODO: Hack assume link is the first token
	_, isGE := r.offRamp.(*geOffRamp)
	linkToken := execTokens[0]
	var me *MessageExecution
	if len(finalSeqNums) > 0 {
		nextMin, err := r.commitStore.GetExpectedNextSequenceNumber(nil, r.onRamp)
		if err != nil {
			return nil, err
		}
		if mathutil.Max(finalSeqNums[0], finalSeqNums[1:]...) >= nextMin {
			return nil, errors.Errorf("Cannot execute uncommitted seq num. nextMin %v, seqNums %v", nextMin, finalSeqNums)
		}
		rep, err := r.builder.commitReport(finalSeqNums[0])
		if err != nil {
			return nil, err
		}
		lggr.Infow("Building execution report", "finalSeqNums", finalSeqNums, "rootOfRoots", hexutil.Encode(rep.RootOfRoots[:]), "report", rep)
		// Otherwise we have messages to include as well.
		var interval commit_store.InternalInterval
		var onRampIdx int
		var outerTreeLeafs [][32]byte
		for i, onRamp := range rep.OnRamps {
			if onRamp == r.onRamp {
				interval = rep.Intervals[i]
				onRampIdx = i
			}
			outerTreeLeafs = append(outerTreeLeafs, rep.MerkleRoots[i])
		}
		if interval.Max == 0 {
			return nil, errors.New("interval not found for ramp " + r.onRamp.Hex())
		}
		msgsInRoot, err := r.source.LogsDataWordRange(r.eventSignatures.SendRequested, r.onRamp, r.eventSignatures.SendRequestedSequenceNumberIndex, EvmWord(interval.Min), EvmWord(interval.Max), int(r.offchainConfig.SourceIncomingConfirmations))
		if err != nil {
			return nil, err
		}
		if len(msgsInRoot) != int(interval.Max-interval.Min+1) {
			return nil, errors.Errorf("unexpected missing msgs in committed root %x have %d want %d", rep.MerkleRoots[onRampIdx], len(msgsInRoot), int(interval.Max-interval.Min+1))
		}
		leafsByOnRamp, err := leafsFromIntervals(
			lggr,
			map[common.Address]EventSignatures{r.onRamp: r.eventSignatures},
			map[common.Address]func(log logpoller.Log) (uint64, error){r.onRamp: r.onRampSeqParser},
			map[common.Address]commit_store.InternalInterval{r.onRamp: interval},
			r.source,
			r.onRampToHasher, int(r.offchainConfig.SourceIncomingConfirmations))
		if err != nil {
			return nil, err
		}
		ctx := hasher.NewKeccakCtx()
		outerTree, err := merklemulti.NewTree[[32]byte](ctx, outerTreeLeafs)
		if err != nil {
			return nil, err
		}
		outerProof := outerTree.Prove([]int{onRampIdx})
		innerTree, err := merklemulti.NewTree[[32]byte](ctx, leafsByOnRamp[r.onRamp])
		if err != nil {
			return nil, err
		}

		var innerIdxs []int
		var encMsgs [][]byte
		var hashes [][32]byte
		for _, seqNum := range finalSeqNums {
			if seqNum < interval.Min || seqNum > interval.Max {
				// We only return messages from a single root (the root of the first message).
				continue
			}
			innerIdx := int(seqNum - interval.Min)
			innerIdxs = append(innerIdxs, innerIdx)
			encMsgs = append(encMsgs, msgsInRoot[innerIdx].Data)
			hash, err2 := r.onRampToHasher[r.onRamp].HashLeaf(LogPollerLogToEthLog(msgsInRoot[innerIdx]))
			if err2 != nil {
				return nil, err2
			}
			hashes = append(hashes, hash)
		}
		innerProof := innerTree.Prove(innerIdxs)
		// Double check this verifies before sending.
		res, err := r.commitStore.Verify(nil, hashes, innerProof.Hashes, ProofFlagsToBits(innerProof.SourceFlags), outerProof.Hashes, ProofFlagsToBits(outerProof.SourceFlags))
		if err != nil {
			lggr.Errorw("Unable to call verify", "seqNums", finalSeqNums, "indices", innerIdxs, "root", rep.RootOfRoots[:], "seqRange", rep.Intervals[onRampIdx], "onRampReport", rep.OnRamps[onRampIdx].Hex(), "onRampHave", r.onRamp.Hex(), "err", err)
			return nil, err
		}
		// No timestamp, means failed to verify root.
		if res.Cmp(big.NewInt(0)) == 0 {
			ir := innerTree.Root()
			or := outerTree.Root()
			lggr.Errorf("Root does not verify for messages: %v (indices %v) our inner root %x our outer root %x contract outer root %x",
				finalSeqNums, innerIdxs, ir[:], or[:], rep.RootOfRoots[:])
			return nil, errors.New("root does not verify")
		}
		me = &MessageExecution{
			encMsgs:               encMsgs,
			innerProofs:           innerProof.Hashes,
			innerProofSourceFlags: innerProof.SourceFlags,
			outerProofs:           outerProof.Hashes,
			outerProofSourceFlags: outerProof.SourceFlags,
		}
	}
	if isGE {
		gasFeeUpdates := r.generateFeeUpdate(sourceGasPrice, tokensPerFeeCoin[linkToken])
		if len(gasFeeUpdates) == 0 && len(finalSeqNums) == 0 {
			return nil, errors.New("No report needed")
		}
		if len(finalSeqNums) != 0 {
			return EncodeGEExecutionReport(finalSeqNums,
				tokensPerFeeCoin,
				me.encMsgs,
				me.innerProofs,
				me.innerProofSourceFlags,
				me.outerProofs,
				me.outerProofSourceFlags,
				gasFeeUpdates,
			)
		}
		lggr.Infow("Building execution report fee update only", "feeUpdates", gasFeeUpdates)
		return EncodeGEExecutionReport(finalSeqNums, tokensPerFeeCoin, nil, nil, nil, nil, nil, gasFeeUpdates)
	}
	return EncodeTollExecutionReport(finalSeqNums, tokensPerFeeCoin, me.encMsgs,
		me.innerProofs, me.innerProofSourceFlags, me.outerProofs, me.outerProofSourceFlags)
}

func median(vals []*big.Int) *big.Int {
	valsCopy := make([]*big.Int, len(vals))
	copy(valsCopy[:], vals[:])
	sort.Slice(valsCopy, func(i, j int) bool {
		return valsCopy[i].Cmp(valsCopy[j]) == -1
	})
	return valsCopy[len(valsCopy)/2]
}

func (r *ExecutionReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	lggr := r.lggr.Named("Report")
	if isCommitStoreDownNow(lggr, r.commitStore) {
		return false, nil, ErrCommitStoreIsDown
	}
	actualMaybeObservations := getNonEmptyObservations[ExecutionObservation](lggr, observations)
	var actualObservations []ExecutionObservation
	tokens, err := r.offRamp.GetSupportedTokensForExecutionFee()
	if err != nil {
		return false, nil, err
	}
	priceObservations := make(map[common.Address][]*big.Int)
	var sourceGasObservations []*big.Int
	for _, obs := range actualMaybeObservations {
		hasAllPrices := true
		for _, token := range tokens {
			if _, ok := obs.TokensPerFeeCoin[token]; !ok {
				hasAllPrices = false
				break
			}
		}
		if !hasAllPrices {
			continue
		}
		if obs.SourceGasPrice == nil {
			lggr.Errorw("Expect source gas price to be present")
			continue
		}
		// If it has all the prices, add each price to observations
		for token, price := range obs.TokensPerFeeCoin {
			priceObservations[token] = append(priceObservations[token], price)
		}
		// Add source gas price
		sourceGasObservations = append(sourceGasObservations, obs.SourceGasPrice)
		actualObservations = append(actualObservations, obs)
	}
	// Need at least F+1 observations
	if len(actualObservations) <= r.F {
		lggr.Tracew("Non-empty observations <= F, need at least F+1 to continue")
		return false, nil, nil
	}
	// If we have sufficient observations, only build a report if
	// the leaders prices is >= the median of the followers prices.
	// Note we accept that this can result in the leader stalling progress,
	// by setting an extremely high set of prices, but a malicious leader always had that ability
	// and eventually we'll elect a new one.
	var q Query
	if err2 := json.Unmarshal(query, &q); err2 != nil {
		return false, nil, err2
	}
	medianTokensPerFeeCoin := make(map[common.Address]*big.Int)
	for _, token := range tokens {
		medianTokensPerFeeCoin[token] = median(priceObservations[token])
		if medianTokensPerFeeCoin[token].Cmp(q.TokenPrices[token]) == 1 {
			// Leader specified a price which is too low, reject this report.
			// TODO: Error or not here?
			lggr.Warnw("Leader price is too low, skipping report", "leader", q.TokenPrices[token], "followers", medianTokensPerFeeCoin[token])
			return false, nil, nil
		}
	}
	// If we make it here, the leader has specified a valid set of prices.
	tally := make(map[uint64]int)
	for _, obs := range actualObservations {
		for _, seqNr := range obs.SeqNrs {
			tally[seqNr]++
		}
	}
	var finalSequenceNumbers []uint64
	for seqNr, count := range tally {
		// Note spec deviation - I think it's ok to rely on the batch builder for capping the number of messages vs capping in two places/ways?
		if count > r.F {
			finalSequenceNumbers = append(finalSequenceNumbers, seqNr)
		}
	}
	// buildReport expects sorted sequence numbers (tally map is non-deterministic).
	sort.Slice(finalSequenceNumbers, func(i, j int) bool {
		return finalSequenceNumbers[i] < finalSequenceNumbers[j]
	})
	// Important we actually execute based on the medianTokensPrices, which we ensure
	// is <= than prices used to determine executability.
	report, err := r.buildReport(lggr, finalSequenceNumbers, medianTokensPerFeeCoin, median(sourceGasObservations))
	if err != nil {
		return false, nil, err
	}
	lggr.Infow("Built report",
		"onRamp", r.onRamp,
		"finalSeqNums", finalSequenceNumbers,
		"executionPrices", medianTokensPerFeeCoin)
	return true, report, nil
}

func (r *ExecutionReportingPlugin) expireInflight(lggr logger.Logger) {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	// Reap old inflight txs and check if any messages in the report are inflight.
	var stillInFlight []InflightExecutionReport
	for _, report := range r.inFlight {
		if time.Since(report.createdAt) > r.inflightCacheExpiry {
			// Happy path: inflight report was successfully transmitted onchain, we remove it from inflight and onchain state reflects inflight.
			// Sad path: inflight report reverts onchain, we remove it from inflight, onchain state does not reflect the chains so we retry.
			lggr.Infow("Inflight report expired", "seqNums", report.seqNrs)
		} else {
			stillInFlight = append(stillInFlight, report)
		}
	}
	r.inFlight = stillInFlight
}

func (r *ExecutionReportingPlugin) addToInflight(lggr logger.Logger, seqNrs []uint64, encMsgs [][]byte) error {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	for _, report := range r.inFlight {
		// TODO: Think about if this fails in reorgs
		if (len(report.seqNrs) > 0 && len(seqNrs) > 0) && (report.seqNrs[0] == seqNrs[0]) {
			return errors.Errorf("report is already in flight")
		}
	}
	// Otherwise not already in flight, add it.
	lggr.Infow("Added report to inflight",
		"seqNums", seqNrs)
	r.inFlight = append(r.inFlight, InflightExecutionReport{
		createdAt:   time.Now(),
		seqNrs:      seqNrs,
		encMessages: encMsgs,
	})
	return nil
}

func (r *ExecutionReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("ShouldAcceptFinalizedReport")
	seqNrs, encMsgs, err := r.parseReport(report)
	if err != nil {
		lggr.Errorw("unable to decode report", "err", err)
		return false, nil
	}
	if len(seqNrs) > 0 {
		lggr.Infof("Seq nums %v", seqNrs)
		// If the first message is executed already, this execution report is stale, and we do not accept it.
		stale, err2 := r.isStaleReport(seqNrs[0])
		if err2 != nil {
			return !stale, err2
		}
		if stale {
			return false, nil
		}
	}
	// Else just assume in flight
	if err = r.addToInflight(lggr, seqNrs, encMsgs); err != nil {
		return false, err
	}
	return true, nil
}

func (r *ExecutionReportingPlugin) parseReport(report []byte) ([]uint64, [][]byte, error) {
	// TODO: Cleaner to separate these plugins?
	_, isGE := r.offRamp.(*geOffRamp)
	if isGE {
		return MessagesFromGEExecutionReport(report)
	}
	return MessagesFromTollExecutionReport(report)
}

func (r *ExecutionReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	seqNrs, _, err := r.parseReport(report)
	if err != nil {
		return false, nil
	}
	// If report is not stale we transmit.
	// When the executeTransmitter enqueues the tx for tx manager,
	// we mark it as execution_sent, removing it from the set of inflight messages.
	if len(seqNrs) > 0 {
		stale, err := r.isStaleReport(seqNrs[0])
		return !stale, err
	}
	// TODO: how to check for staleness on a purely fee update report?
	return false, nil
}

func (r *ExecutionReportingPlugin) isStaleReport(min uint64) (bool, error) {
	// If the first message is executed already, this execution report is stale.
	msgState, err := r.offRamp.GetExecutionState(nil, min)
	if err != nil {
		// TODO: do we need to check for not present error?
		return true, err
	}
	if msgState == MessageStateFailure || msgState == MessageStateSuccess {
		return true, nil
	}
	return false, nil
}

func (r *ExecutionReportingPlugin) Close() error {
	return nil
}
