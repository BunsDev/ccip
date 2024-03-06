package ccipcommon

import (
	"errors"
	"math/rand"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/mocks"
)

func TestGetMessageIDsAsHexString(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		hashes := make([]cciptypes.Hash, 10)
		for i := range hashes {
			hashes[i] = cciptypes.Hash(common.HexToHash(strconv.Itoa(rand.Intn(100000))))
		}

		msgs := make([]cciptypes.EVM2EVMMessage, len(hashes))
		for i := range msgs {
			msgs[i] = cciptypes.EVM2EVMMessage{MessageID: hashes[i]}
		}

		messageIDs := GetMessageIDsAsHexString(msgs)
		for i := range messageIDs {
			assert.Equal(t, hashes[i].String(), messageIDs[i])
		}
	})

	t.Run("empty", func(t *testing.T) {
		messageIDs := GetMessageIDsAsHexString(nil)
		assert.Empty(t, messageIDs)
	})
}

func TestFlattenUniqueSlice(t *testing.T) {
	testCases := []struct {
		name           string
		inputSlices    [][]int
		expectedOutput []int
	}{
		{name: "empty", inputSlices: nil, expectedOutput: []int{}},
		{name: "empty 2", inputSlices: [][]int{}, expectedOutput: []int{}},
		{name: "single", inputSlices: [][]int{{1, 2, 3, 3, 3, 4}}, expectedOutput: []int{1, 2, 3, 4}},
		{name: "simple", inputSlices: [][]int{{1, 2, 3}, {2, 3, 4}}, expectedOutput: []int{1, 2, 3, 4}},
		{
			name:           "more complex case",
			inputSlices:    [][]int{{1, 3}, {2, 4, 3}, {5, 2, -1, 7, 10}},
			expectedOutput: []int{1, 3, 2, 4, 5, -1, 7, 10},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := FlattenUniqueSlice(tc.inputSlices...)
			assert.Equal(t, tc.expectedOutput, res)
		})
	}
}

func TestVerifyNotDown(t *testing.T) {
	ctx := tests.Context(t)

	testCases := []struct {
		name            string
		commitStoreDown bool
		commitStoreErr  error
		onRampCursed    bool
		onRampErr       error
		expectedErr     bool
	}{
		{
			name:        "Neither down nor cursed",
			expectedErr: false,
		},
		{
			name:            "CommitStore is down",
			commitStoreDown: true,
			expectedErr:     true,
		},
		{
			name:         "OnRamp is cursed",
			onRampCursed: true,
			expectedErr:  true,
		},
		{
			name:           "CommitStore error",
			commitStoreErr: errors.New("commit store error"),
			expectedErr:    true,
		},
		{
			name:        "OnRamp error",
			onRampErr:   errors.New("onramp error"),
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCommitStore := mocks.NewCommitStoreReader(t)
			mockOnRamp := mocks.NewOnRampReader(t)

			mockCommitStore.On("IsDown", ctx).Return(tc.commitStoreDown, tc.commitStoreErr)
			mockOnRamp.On("IsSourceCursed", ctx).Return(tc.onRampCursed, tc.onRampErr)

			err := VerifyNotDown(ctx, logger.TestLogger(t), mockCommitStore, mockOnRamp)

			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
