package ccip

import (
	"encoding/binary"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/logger"
)

const (
	MaxObservationLength = 500 // TODO: Think about what to set this too
	MaxQueryLength       = 500 // TODO: Think about what to set this too, roughly maxTokens*maxPriceInBytesPerToken
)

func EvmWord(i uint64) common.Hash {
	var b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return common.BigToHash(big.NewInt(0).SetBytes(b))
}

type RelayObservation struct {
	IntervalsByOnRamp map[common.Address]blob_verifier.CCIPInterval `json:"intervalsByOnRamp"`
}

func (o RelayObservation) Marshal() ([]byte, error) {
	return json.Marshal(&o)
}

type ExecutionObservation struct {
	SeqNrs           []uint64                    `json:"seqNrs"`
	TokensPerFeeCoin map[common.Address]*big.Int `json:"tokensPerFeeCoin"`
}

func (o ExecutionObservation) Marshal() ([]byte, error) {
	return json.Marshal(&o)
}

// getNonEmptyObservations checks the given observations for formatting and value errors.
// It returns all valid observations, potentially being an empty list. It will log
// malformed observations but never error.
func getNonEmptyObservations[O RelayObservation | ExecutionObservation](l logger.Logger, observations []types.AttributedObservation) []O {
	var nonEmptyObservations []O
	for _, ao := range observations {
		if len(ao.Observation) == 0 {
			// Empty observation
			l.Tracew("Discarded empty observation %+v", ao)
			continue
		}
		var ob O
		err := json.Unmarshal(ao.Observation, &ob)
		if err != nil {
			l.Errorw("Received unmarshallable observation", "err", err, "observation", string(ao.Observation))
			continue
		}
		nonEmptyObservations = append(nonEmptyObservations, ob)
	}
	return nonEmptyObservations
}
