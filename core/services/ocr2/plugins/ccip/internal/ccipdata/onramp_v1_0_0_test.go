package ccipdata

import (
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/hashlib"
)

func TestHasher(t *testing.T) {
	sourceChainSelector, destChainSelector := uint64(1), uint64(4)
	onRampAddress := common.HexToAddress("0x5550000000000000000000000000000000000001")
	onRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_onramp_1_0_0.EVM2EVMOnRampABI))
	require.NoError(t, err)

	ramp, err := evm_2_evm_onramp_1_0_0.NewEVM2EVMOnRamp(onRampAddress, nil)
	require.NoError(t, err)
	hashingCtx := hashlib.NewKeccakCtx()
	hasher := NewLeafHasherV1_0_0(sourceChainSelector, destChainSelector, onRampAddress, hashingCtx, ramp)

	message := evm_2_evm_onramp_1_0_0.InternalEVM2EVMMessage{
		SourceChainSelector: sourceChainSelector,
		Sender:              common.HexToAddress("0x1110000000000000000000000000000000000001"),
		Receiver:            common.HexToAddress("0x2220000000000000000000000000000000000001"),
		SequenceNumber:      1337,
		GasLimit:            big.NewInt(100),
		Strict:              false,
		Nonce:               1337,
		FeeToken:            common.Address{},
		FeeTokenAmount:      big.NewInt(1),
		Data:                []byte{},
		TokenAmounts:        []evm_2_evm_onramp_1_0_0.ClientEVMTokenAmount{{Token: common.HexToAddress("0x4440000000000000000000000000000000000001"), Amount: big.NewInt(12345678900)}},
		MessageId:           [32]byte{},
	}

	data, err := onRampABI.Events["CCIPSendRequested"].Inputs.Pack(message)
	require.NoError(t, err)
	hash, err := hasher.HashLeaf(types.Log{Topics: []common.Hash{abihelpers.GetIDOrPanic("CCIPSendRequested", onRampABI)}, Data: data})
	require.NoError(t, err)

	// NOTE: Must match spec
	require.Equal(t, "46ad031bfb052db2e4a2514fed8dc480b98e5ce4acb55d5640d91407e0d8a3e9", hex.EncodeToString(hash[:]))

	message = evm_2_evm_onramp_1_0_0.InternalEVM2EVMMessage{
		SourceChainSelector: sourceChainSelector,
		Sender:              common.HexToAddress("0x1110000000000000000000000000000000000001"),
		Receiver:            common.HexToAddress("0x2220000000000000000000000000000000000001"),
		SequenceNumber:      1337,
		GasLimit:            big.NewInt(100),
		Strict:              false,
		Nonce:               1337,
		FeeToken:            common.Address{},
		FeeTokenAmount:      big.NewInt(1e12),
		Data:                []byte("foo bar baz"),
		TokenAmounts: []evm_2_evm_onramp_1_0_0.ClientEVMTokenAmount{
			{Token: common.HexToAddress("0x4440000000000000000000000000000000000001"), Amount: big.NewInt(12345678900)},
			{Token: common.HexToAddress("0x6660000000000000000000000000000000000001"), Amount: big.NewInt(4204242)},
		},
		MessageId: [32]byte{},
	}

	data, err = onRampABI.Events["CCIPSendRequested"].Inputs.Pack(message)
	require.NoError(t, err)
	hash, err = hasher.HashLeaf(types.Log{Topics: []common.Hash{abihelpers.GetIDOrPanic("CCIPSendRequested", onRampABI)}, Data: data})
	require.NoError(t, err)

	// NOTE: Must match spec
	require.Equal(t, "4362a13a42e52ff5ce4324e7184dc7aa41704c3146bc842d35d95b94b32a78b6", hex.EncodeToString(hash[:]))
}

func TestMetaDataHash(t *testing.T) {
	sourceChainSelector, destChainSelector := uint64(1), uint64(4)
	onRampAddress := common.HexToAddress("0x5550000000000000000000000000000000000001")
	ctx := hashlib.NewKeccakCtx()
	hash := getMetaDataHash(ctx, ctx.Hash([]byte("EVM2EVMSubscriptionMessagePlus")), sourceChainSelector, onRampAddress, destChainSelector)
	require.Equal(t, "e8b93c9d01a7a72ec6c7235e238701cf1511b267a31fdb78dd342649ee58c08d", hex.EncodeToString(hash[:]))
}
