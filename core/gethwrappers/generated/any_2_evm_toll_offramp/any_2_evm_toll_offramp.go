// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package any_2_evm_toll_offramp

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

type AggregateRateLimiterInterfaceRateLimiterConfig struct {
	Rate     *big.Int
	Capacity *big.Int
}

type AggregateRateLimiterInterfaceTokenBucket struct {
	Rate        *big.Int
	Capacity    *big.Int
	Tokens      *big.Int
	LastUpdated *big.Int
}

type BaseOffRampInterfaceOffRampConfig struct {
	OnRampAddress                           common.Address
	PermissionLessExecutionThresholdSeconds uint32
	ExecutionDelaySeconds                   uint64
	MaxDataSize                             uint64
	MaxTokensLength                         uint64
}

type CCIPAny2EVMMessageFromSender struct {
	SourceChainId        *big.Int
	Sender               []byte
	Receiver             common.Address
	Data                 []byte
	DestPools            []common.Address
	DestTokensAndAmounts []CCIPEVMTokenAndAmount
	GasLimit             *big.Int
}

type CCIPEVM2EVMTollMessage struct {
	SourceChainId     *big.Int
	SequenceNumber    uint64
	Sender            common.Address
	Receiver          common.Address
	Data              []byte
	TokensAndAmounts  []CCIPEVMTokenAndAmount
	FeeTokenAndAmount CCIPEVMTokenAndAmount
	GasLimit          *big.Int
}

type CCIPEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

type CCIPExecutionReport struct {
	SequenceNumbers          []uint64
	TokenPerFeeCoinAddresses []common.Address
	TokenPerFeeCoin          []*big.Int
	EncodedMessages          [][]byte
	InnerProofs              [][32]byte
	InnerProofFlagBits       *big.Int
	OuterProofs              [][32]byte
	OuterProofFlagBits       *big.Int
}

var EVM2EVMTollOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeTaken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobVerifier\",\"outputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"merkleGasShare\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount\",\"name\":\"feeTokenAndAmount\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVM2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"overheadGasToll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"}],\"name\":\"setBlobVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620064fe380380620064fe83398101604081905262000034916200063c565b6000805460ff191681556001908a908a908a908a908a908a908a908a908a90829082908690869089903390819081620000b45760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ee57620000ee8162000385565b5050506001600160a01b0381166200011957604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015b5760405162d8548360e71b815260040160405180910390fd5b81516200017090600490602085019062000436565b5060005b82518110156200023b57600082828151811062000195576200019562000796565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001df57620001df62000796565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b0292909116919091179055506200023381620007ac565b905062000174565b5050600580546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600855600993909355600a55600b91909155875116620002b9576040516342bcdf7f60e11b815260040160405180910390fd5b5050506080958652505060a0929092528051600e8054602084015160408501516001600160401b03908116600160c01b026001600160c01b0363ffffffff909316600160a01b026001600160c01b03199094166001600160a01b03968716179390931791909116919091179091556060830151600f80549490960151821668010000000000000000026001600160801b031990941691161791909117909255600d8054919092166001600160a01b031991909116179055151560c05250620007d4975050505050505050565b336001600160a01b03821603620003df5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ab565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200048e579160200282015b828111156200048e57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000457565b506200049c929150620004a0565b5090565b5b808211156200049c5760008155600101620004a1565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715620004f257620004f2620004b7565b60405290565b6001600160a01b03811681146200050e57600080fd5b50565b80516200051e81620004f8565b919050565b80516001600160401b03811681146200051e57600080fd5b600082601f8301126200054d57600080fd5b815160206001600160401b03808311156200056c576200056c620004b7565b8260051b604051601f19603f83011681018181108482111715620005945762000594620004b7565b604052938452858101830193838101925087851115620005b357600080fd5b83870191505b84821015620005df578151620005cf81620004f8565b83529183019190830190620005b9565b979650505050505050565b600060408284031215620005fd57600080fd5b604080519081016001600160401b0381118282101715620006225762000622620004b7565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000898b036101c08112156200065d57600080fd5b8a51995060208b0151985060a0603f19820112156200067b57600080fd5b5062000686620004cd565b60408b01516200069681620004f8565b815260608b015163ffffffff81168114620006b057600080fd5b6020820152620006c360808c0162000523565b6040820152620006d660a08c0162000523565b6060820152620006e960c08c0162000523565b60808201529650620006fe60e08b0162000511565b95506200070f6101008b0162000511565b6101208b01519095506001600160401b03808211156200072e57600080fd5b6200073c8d838e016200053b565b95506101408c01519150808211156200075457600080fd5b50620007638c828d016200053b565b935050620007768b6101608c01620005ea565b9150620007876101a08b0162000511565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b600060018201620007cd57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c051615ce56200081960003960006116a001526000818161030b0152611f0a0152600081816102e801528181611ee5015261387d0152615ce56000f3fe608060405234801561001057600080fd5b50600436106102c85760003560e01c80638456cb591161017b578063b1dc65a4116100d8578063c3f909d41161008c578063e3d0e71211610071578063e3d0e712146107a7578063eb511dd4146107ba578063f2fde38b146107cd57600080fd5b8063c3f909d41461067c578063c9029f6a1461079457600080fd5b8063bbe4f6db116100bd578063bbe4f6db1461062a578063be9b03f114610656578063c0d786551461066957600080fd5b8063b1dc65a414610604578063b4069b311461061757600080fd5b80639438ff631161012f578063a8b640c111610114578063a8b640c1146105b3578063afcb95d7146105d3578063b0f479a1146105f357600080fd5b80639438ff6314610592578063a639d1c0146105a057600080fd5b80638da5cb5b116101605780638da5cb5b1461052e57806390c2339b14610544578063918725431461057f57600080fd5b80638456cb591461051e57806389c065681461052657600080fd5b80634741062e11610229578063681fba16116101dd57806379ba5097116101c257806379ba5097146104d157806381411834146104d957806381ff7048146104ee57600080fd5b8063681fba16146104a9578063744b92e2146104be57600080fd5b80635b4dc8121161020e5780635b4dc8121461046c5780635c975abb1461048d5780636133dc241461049857600080fd5b80634741062e1461043b578063599f64311461045b57600080fd5b8063181f5a771161028057806339aa92641161026557806339aa92641461040d5780633f4ba83a146104205780634352fa9f1461042857600080fd5b8063181f5a771461039f5780632222dd42146103e857600080fd5b8063108ee5fc116102b1578063108ee5fc1461033a578063142a98fc1461034d578063147809b31461038757600080fd5b806307a22a07146102cd578063087ae6df146102e2575b600080fd5b6102e06102db366004614650565b6107e0565b005b604080517f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006020820152015b60405180910390f35b6102e0610348366004614751565b610854565b61037a61035b36600461478f565b67ffffffffffffffff1660009081526010602052604090205460ff1690565b60405161033191906147db565b61038f61090b565b6040519015158152602001610331565b6103db6040518060400160405280601881526020017f45564d3245564d546f6c6c4f666652616d7020312e302e30000000000000000081525081565b6040516103319190614874565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610331565b6102e061041b366004614751565b610998565b6102e06109cf565b6102e06104363660046148e2565b6109e1565b61044e610449366004614946565b610c36565b60405161033191906149be565b6005546001600160a01b03166103f5565b61047f61047a3660046149d1565b610cfe565b604051908152602001610331565b60005460ff1661038f565b600d546001600160a01b03166103f5565b6104b1610daf565b6040516103319190614ac9565b6102e06104cc366004614b16565b610e74565b6102e06111b8565b6104e16112a0565b6040516103319190614b88565b6013546011546040805163ffffffff80851682526401000000009094049093166020840152820152606001610331565b6102e0611302565b6104b1611312565b60005461010090046001600160a01b03166103f5565b61054c611372565b60405161033191908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6102e061058d366004614b9b565b611413565b6102e06102c8366004614bea565b6102e06105ae366004614751565b611546565b61047f6105c1366004614c25565b60176020526000908152604090205481565b604080516001815260006020820181905291810191909152606001610331565b600c546001600160a01b03166103f5565b6102e0610612366004614c8a565b61157d565b6103f5610625366004614751565b611b50565b6103f5610638366004614751565b6001600160a01b039081166000908152600360205260409020541690565b6102e0610664366004614e6c565b611c3e565b6102e0610677366004614751565b612367565b6107386040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a081018252600e546001600160a01b038116825263ffffffff74010000000000000000000000000000000000000000820416602083015267ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811692820192909252600f548083166060830152680100000000000000009004909116608082015290565b6040805182516001600160a01b0316815260208084015163ffffffff16908201528282015167ffffffffffffffff908116928201929092526060808401518316908201526080928301519091169181019190915260a001610331565b6102e06107a2366004614fb5565b6123c6565b6102e06107b5366004615067565b612565565b6102e06107c8366004614b16565b612e5a565b6102e06107db366004614751565b613032565b333014610819576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a08101515115610848576108318160a00151613043565b61084881608001518260a0015183604001516131dd565b6108518161327b565b50565b61085c613353565b6001600160a01b03811661089c576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa15801561096e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109929190615134565b15905090565b6109a0613353565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6109d7613353565b6109df6133b2565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015610a1e57506005546001600160a01b03163314155b15610a55576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610a91576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610aeb576006600060078381548110610ab657610ab6615151565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610ae4816151af565b9050610a97565b5060005b82811015610c1b576000858281518110610b0b57610b0b615151565b6020026020010151905060006001600160a01b0316816001600160a01b031603610b61576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610b7357610b73615151565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610bd857610bd8615151565b6020026020010151604051610c029291906001600160a01b03929092168252602082015260400190565b60405180910390a150610c14816151af565b9050610aef565b508351610c2f90600790602087019061430e565b5050505050565b80516060908067ffffffffffffffff811115610c5457610c546143b0565b604051908082528060200260200182016040528015610c7d578160200160208202803683370190505b50915060005b81811015610cf75760066000858381518110610ca157610ca1615151565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610cdc57610cdc615151565b6020908102919091010152610cf0816151af565b9050610c83565b5050919050565b6000808260800151518360a001515160206014610d1b91906151e7565b610d2591906151ff565b610d309060866151e7565b610d3a91906151e7565b90506000610d496010836151ff565b9050610a28611bbc8560a00151516001610d6391906151e7565b610d6f90618aac6151ff565b6156b8610d7c89866151e7565b610d8691906151e7565b610d9091906151e7565b610d9a91906151e7565b610da491906151e7565b925050505b92915050565b60045460609067ffffffffffffffff811115610dcd57610dcd6143b0565b604051908082528060200260200182016040528015610df6578160200160208202803683370190505b50905060005b600454811015610e7057610e3660048281548110610e1c57610e1c615151565b6000918252602090912001546001600160a01b0316611b50565b828281518110610e4857610e48615151565b6001600160a01b0390921660209283029190910190910152610e69816151af565b9050610dfc565b5090565b610e7c613353565b6004546000819003610eba576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610f48576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614610f97576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610fa660018561523c565b81548110610fb657610fb6615151565b9060005260206000200160009054906101000a90046001600160a01b03169050600482602001516bffffffffffffffffffffffff1681548110610ffb57610ffb615151565b6000918252602090912001546001600160a01b0316600461101d60018661523c565b8154811061102d5761102d615151565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483602001516bffffffffffffffffffffffff168154811061108157611081615151565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000291909216179055600480548061110b5761110b615253565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9084018101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03878116808452600383526040808520949094558351908152908716918101919091527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146112175760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b606060168054806020026020016040519081016040528092919081815260200182805480156112f857602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116112da575b5050505050905090565b61130a613353565b6109df61344e565b606060048054806020026020016040519081016040528092919081815260200182805480156112f8576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116112da575050505050905090565b61139d6040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b546060820181905242906000906113d7908361523c565b60208401518451919250611403916113ef90846151ff565b85604001516113fe91906151e7565b6134d6565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561145057506005546001600160a01b03163314155b15611487576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff116114db576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114e560086134ec565b602081015160098190558151600855600a5461150191906134d6565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b61154e613353565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916115d391849163ffffffff851691908e908e908190840183828082843760009201919091525061359992505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260115480825260125460ff8082166020850152610100909104169282019290925290831461168e5760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d617463680000000000000000000000604482015260640161120e565b61169c8b8b8b8b8b8b613612565b60007f0000000000000000000000000000000000000000000000000000000000000000156116f9576002826020015183604001516116da9190615282565b6116e491906152d6565b6116ef906001615282565b60ff16905061170f565b6020820151611709906001615282565b60ff1690505b88811461175e5760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604482015260640161120e565b8887146117ad5760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604482015260640161120e565b3360009081526014602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156117f0576117f06147ac565b6002811115611801576118016147ac565b905250905060028160200151600281111561181e5761181e6147ac565b14801561185857506016816000015160ff168154811061184057611840615151565b6000918252602090912001546001600160a01b031633145b6118a45760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604482015260640161120e565b5050505050600088886040516118bb9291906152f8565b6040519081900381206118d2918c90602001615308565b6040516020818303038152906040528051906020012090506118f261437c565b604080518082019091526000808252602082015260005b88811015611b2e57600060018588846020811061192857611928615151565b61193591901a601b615282565b8d8d8681811061194757611947615151565b905060200201358c8c8781811061196057611960615151565b905060200201356040516000815260200160405260405161199d949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156119bf573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526014602090815290849020838501909452835460ff80821685529296509294508401916101009004166002811115611a1457611a146147ac565b6002811115611a2557611a256147ac565b9052509250600183602001516002811115611a4257611a426147ac565b14611a8f5760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e0000604482015260640161120e565b8251849060ff16601f8110611aa657611aa6615151565b602002015115611af85760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e6174757265000000000000000000000000604482015260640161120e565b600184846000015160ff16601f8110611b1357611b13615151565b9115156020909202015250611b27816151af565b9050611909565b5050505063ffffffff8110611b4557611b45615324565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611ba4576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611c13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c379190615353565b9392505050565b60005460ff1615611c915760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161120e565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ce4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d089190615134565b15611d3e576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c546001600160a01b0316611d80576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060820151516000819003611dc1576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115611ddc57611ddc6143b0565b604051908082528060200260200182016040528015611e6357816020015b611e5060408051610100810182526000808252602080830182905282840182905260608084018390526080840181905260a084015283518085019094528184528301529060c08201908152602001600081525090565b815260200190600190039081611dfa5790505b50905060008267ffffffffffffffff811115611e8157611e816143b0565b604051908082528060200260200182016040528015611eaa578160200160208202803683370190505b5090506000611f5b7f31a97e998befaf21ee85c8e3e1879003d67d46abb0b6b552882133e4c782986d600e54604080516020808201949094527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201526001600160a01b039092166080808401919091528151808403909101815260a09092019052805191012090565b905060005b84811015611ff557600087606001518281518110611f8057611f80615151565b6020026020010151806020019051810190611f9b919061546a565b9050611fa78184613699565b848381518110611fb957611fb9615151565b60200260200101818152505080858381518110611fd857611fd8615151565b60200260200101819052505080611fee906151af565b9050611f60565b506000806120168489608001518a60a001518b60c001518c60e00151613781565b600e54919350915060009074010000000000000000000000000000000000000000900463ffffffff16612049844261523c565b11905060005b8781101561235b57600087828151811061206b5761206b615151565b60200260200101519050600061209e826020015167ffffffffffffffff1660009081526010602052604090205460ff1690565b905060028160038111156120b4576120b46147ac565b036120fd5760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161120e565b8a15806121075750835b8061212357506003816003811115612121576121216147ac565b145b612159576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61216282613879565b600080826003811115612177576121776147ac565b14801561218257508b155b156121db5761219e8a51876121979190615557565b8e856139ab565b905060006121b38460c0015160000151613b98565b90506121c0818330613bfa565b818460c001516020018181516121d6919061523c565b905250505b60008260038111156121ef576121ef6147ac565b1461222b5760208084015167ffffffffffffffff16600090815260178252604090205460c0850151909101805161222790839061523c565b9052505b60208381015167ffffffffffffffff166000908152601090915260408120805460ff1916600117905561226561226085613c48565b613f11565b60208086015167ffffffffffffffff1660009081526010909152604090208054919250829160ff191660018360038111156122a2576122a26147ac565b021790555060008360038111156122bb576122bb6147ac565b1480156122d9575060038160038111156122d7576122d76147ac565b145b156123005760208085015167ffffffffffffffff1660009081526017909152604090208290555b836020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf8260405161233e91906147db565b60405180910390a25050505080612354906151af565b905061204f565b50505050505050505050565b61236f613353565b600c805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b6123ce613353565b80516001600160a01b031661240f576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600e805460208085018051604080880180516001600160a01b039098167fffffffffffffffff00000000000000000000000000000000000000000000000090961686177401000000000000000000000000000000000000000063ffffffff948516021777ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff998a16021790965560608089018051600f80546080808e018051948e167fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169290921768010000000000000000948e1694909402939093179091558451988952955190941695870195909552955187169085015291518516938301939093529151909216908201527fd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa09060a00161153b565b855185518560ff16601f8311156125d8576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e65727300000000000000000000000000000000604482015260640161120e565b80600003612642576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f7369746976650000000000000000000000000000604482015260640161120e565b8183146126d0576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e00000000000000000000000000000000000000000000000000000000606482015260840161120e565b6126db8160036151ff565b8311612743576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604482015260640161120e565b61274b613353565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b6015541561290e576015546000906127a39060019061523c565b90506000601582815481106127ba576127ba615151565b6000918252602082200154601680546001600160a01b03909216935090849081106127e7576127e7615151565b60009182526020808320909101546001600160a01b0385811684526014909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009081169091559290911680845292208054909116905560158054919250908061285a5761285a615253565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908101805473ffffffffffffffffffffffffffffffffffffffff1916905501905560168054806128b8576128b8615253565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908101805473ffffffffffffffffffffffffffffffffffffffff1916905501905550612789915050565b60005b815151811015612cd3576000601460008460000151848151811061293757612937615151565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115612974576129746147ac565b146129db576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e65722061646472657373000000000000000000604482015260640161120e565b6040805180820190915260ff82168152600160208201528251805160149160009185908110612a0c57612a0c615151565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612a8257612a826147ac565b021790555060009150612a929050565b6014600084602001518481518110612aac57612aac615151565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115612ae957612ae96147ac565b14612b50576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161120e565b6040805180820190915260ff821681526020810160028152506014600084602001518481518110612b8357612b83615151565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612bf957612bf96147ac565b021790555050825180516015925083908110612c1757612c17615151565b6020908102919091018101518254600181018455600093845292829020909201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039093169290921790915582015180516016919083908110612c7b57612c7b615151565b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909216919091179055612ccc816151af565b9050612911565b5060408101516012805460ff191660ff909216919091179055601380547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612d4792869290821691161761556b565b92506101000a81548163ffffffff021916908363ffffffff160217905550612da64630601360009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151614045565b6011819055825180516012805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560135460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612e45988b98919763ffffffff909216969095919491939192615593565b60405180910390a15050505050505050505050565b612e62613353565b6001600160a01b0382161580612e7f57506001600160a01b038116155b15612eb6576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612f45576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201805473ffffffffffffffffffffffffffffffffffffffff1916821790558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b61303a613353565b610851816140d2565b6000805b82518110156131425760006006600085848151811061306857613068615151565b6020026020010151600001516001600160a01b03166001600160a01b03168152602001908152602001600020549050806000036130fb578382815181106130b1576130b1615151565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260240161120e565b83828151811061310d5761310d615151565b6020026020010151602001518161312491906151ff565b61312e90846151e7565b9250508061313b906151af565b9050613047565b5080156131d95761315360086134ec565b600a5481111561318f576040517f3bfa6f3800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600860020160008282546131a4919061523c565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378906020016108ff565b5050565b8151835114613218576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b83518110156132755761326584828151811061323957613239615151565b602002602001015184838151811061325357613253615151565b60200260200101516020015184613bfa565b61326e816151af565b905061321b565b50505050565b60408101516001600160a01b03163b6132915750565b600c546040517fda52b4c40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063da52b4c4906132da908490600401615676565b6020604051808303816000875af11580156132f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061331d9190615134565b610851576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005461010090046001600160a01b031633146109df5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161120e565b60005460ff166134045760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161120e565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60005460ff16156134a15760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161120e565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586134313390565b60008183106134e55781611c37565b5090919050565b60018101546002820154429114806135075750808260030154145b15613510575050565b816001015482600201541115613552576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482613564919061523c565b6001840154845491925061358b9161357c90846151ff565b85600201546113fe91906151e7565b600284015550600390910155565b306001600160a01b031663be9b03f1828060200190518101906135bc91906158c3565b60006040518363ffffffff1660e01b81526004016135db929190615a8b565b600060405180830381600087803b1580156135f557600080fd5b505af1158015613609573d6000803e3d6000fd5b50505050505050565b600061361f8260206151ff565b61362a8560206151ff565b613636886101446151e7565b61364091906151e7565b61364a91906151e7565b6136559060006151e7565b9050368114613609576040517f8e1192e10000000000000000000000000000000000000000000000000000000081526004810182905236602482015260440161120e565b60008060001b828460200151856040015186606001518760800151805190602001208860a001516040516020016136d09190615b82565b604051602081830303815290604052805190602001208960e001518a60c00151604051602001613763999897969594939291909889526020808a019890985267ffffffffffffffff9690961660408901526001600160a01b039485166060890152928416608088015260a087019190915260c086015260e085015281511661010084015201516101208201526101400190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600d546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce906137de908c908c908c908c908c90600401615b95565b6020604051808303816000875af11580156137fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138219190615be7565b90506000811161385d576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a613869908461523c565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146138d95780516040517fd44bc9eb000000000000000000000000000000000000000000000000000000008152600481019190915260240161120e565b600f5460a0820151516801000000000000000090910467ffffffffffffffff1610156139435760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161120e565b600f5460808201515167ffffffffffffffff909116101561085157600f546080820151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9092166004830152602482015260440161120e565b60008060006139c18460c0015160000151613b98565b6001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156139fe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a229190615353565b905060005b856020015151811015613a9b57816001600160a01b031686602001518281518110613a5457613a54615151565b60200260200101516001600160a01b031603613a8b5785604001518181518110613a8057613a80615151565b602002602001015192505b613a94816151af565b9050613a27565b5081613ade576040517fce480bcc0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161120e565b6000670de0b6b3a7640000833a8760e00151613afa8b8a610cfe565b613b0491906151e7565b613b0e91906151ff565b613b1891906151ff565b613b229190615557565b90508460c0015160200151811115613b8e5760208086015160c0870151909101516040517f394a2c2700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909216600483015260248201839052604482015260640161120e565b9695505050505050565b6001600160a01b038181166000908152600360205260409020541680613bf5576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161120e565b919050565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a2906044016135db565b613c916040518060e00160405280600081526020016060815260200160006001600160a01b03168152602001606081526020016060815260200160608152602001600081525090565b6000613ca58360a001518460c0015161418e565b805190915060008167ffffffffffffffff811115613cc557613cc56143b0565b604051908082528060200260200182016040528015613d0a57816020015b6040805180820190915260008082526020820152815260200190600190039081613ce35790505b50905060008267ffffffffffffffff811115613d2857613d286143b0565b604051908082528060200260200182016040528015613d51578160200160208202803683370190505b50905060005b83811015613e87576000613d87868381518110613d7657613d76615151565b602002602001015160000151613b98565b905080838381518110613d9c57613d9c615151565b60200260200101906001600160a01b031690816001600160a01b0316815250506040518060400160405280826001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613e05573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e299190615353565b6001600160a01b03168152602001878481518110613e4957613e49615151565b602002602001015160200151815250848381518110613e6a57613e6a615151565b60200260200101819052505080613e80906151af565b9050613d57565b506040518060e00160405280876000015181526020018760400151604051602001613ec191906001600160a01b0391909116815260200190565b604051602081830303815290604052815260200187606001516001600160a01b03168152602001876080015181526020018281526020018381526020018760e00151815250945050505050919050565b6040517f07a22a0700000000000000000000000000000000000000000000000000000000815260009030906307a22a0790613f50908590600401615676565b600060405180830381600087803b158015613f6a57600080fd5b505af1925050508015613f7b575060015b61403d573d808015613fa9576040519150601f19603f3d011682016040523d82523d6000602084013e613fae565b606091505b50613fb881615c00565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da800000000000000000000000000000000000000000000000000000000036140095750600392915050565b806040517fcf19edfd00000000000000000000000000000000000000000000000000000000815260040161120e9190614874565b506002919050565b6000808a8a8a8a8a8a8a8a8a60405160200161406999989796959493929190615c50565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b0382160361412a5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161120e565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b606060005b83518110156142225782600001516001600160a01b03168482815181106141bc576141bc615151565b6020026020010151600001516001600160a01b0316036142125782602001518482815181106141ed576141ed615151565b602002602001015160200181815161420591906151e7565b905250839150610da99050565b61421b816151af565b9050614193565b5060008351600161423391906151e7565b67ffffffffffffffff81111561424b5761424b6143b0565b60405190808252806020026020018201604052801561429057816020015b60408051808201909152600080825260208201528152602001906001900390816142695790505b50905060005b84518110156142e7578481815181106142b1576142b1615151565b60200260200101518282815181106142cb576142cb615151565b6020026020010181905250806142e0906151af565b9050614296565b5082818551815181106142fc576142fc615151565b60209081029190910101529392505050565b828054828255906000526020600020908101928215614370579160200282015b82811115614370578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390911617825560209092019160019091019061432e565b50610e7092915061439b565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610e70576000815560010161439c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715614402576144026143b0565b60405290565b60405160e0810167ffffffffffffffff81118282101715614402576144026143b0565b604051610100810167ffffffffffffffff81118282101715614402576144026143b0565b604051601f8201601f1916810167ffffffffffffffff81118282101715614478576144786143b0565b604052919050565b600067ffffffffffffffff82111561449a5761449a6143b0565b50601f01601f191660200190565b600082601f8301126144b957600080fd5b81356144cc6144c782614480565b61444f565b8181528460208386010111156144e157600080fd5b816020850160208301376000918101602001919091529392505050565b6001600160a01b038116811461085157600080fd5b8035613bf5816144fe565b600067ffffffffffffffff821115614538576145386143b0565b5060051b60200190565b600082601f83011261455357600080fd5b813560206145636144c78361451e565b82815260059290921b8401810191818101908684111561458257600080fd5b8286015b848110156145a6578035614599816144fe565b8352918301918301614586565b509695505050505050565b6000604082840312156145c357600080fd5b6145cb6143df565b905081356145d8816144fe565b808252506020820135602082015292915050565b600082601f8301126145fd57600080fd5b8135602061460d6144c78361451e565b82815260069290921b8401810191818101908684111561462c57600080fd5b8286015b848110156145a65761464288826145b1565b835291830191604001614630565b60006020828403121561466257600080fd5b813567ffffffffffffffff8082111561467a57600080fd5b9083019060e0828603121561468e57600080fd5b614696614408565b823581526020830135828111156146ac57600080fd5b6146b8878286016144a8565b6020830152506146ca60408401614513565b60408201526060830135828111156146e157600080fd5b6146ed878286016144a8565b60608301525060808301358281111561470557600080fd5b61471187828601614542565b60808301525060a08301358281111561472957600080fd5b614735878286016145ec565b60a08301525060c083013560c082015280935050505092915050565b60006020828403121561476357600080fd5b8135611c37816144fe565b67ffffffffffffffff8116811461085157600080fd5b8035613bf58161476e565b6000602082840312156147a157600080fd5b8135611c378161476e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310614816577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b8381101561483757818101518382015260200161481f565b838111156132755750506000910152565b6000815180845261486081602086016020860161481c565b601f01601f19169290920160200192915050565b602081526000611c376020830184614848565b600082601f83011261489857600080fd5b813560206148a86144c78361451e565b82815260059290921b840181019181810190868411156148c757600080fd5b8286015b848110156145a657803583529183019183016148cb565b600080604083850312156148f557600080fd5b823567ffffffffffffffff8082111561490d57600080fd5b61491986838701614542565b9350602085013591508082111561492f57600080fd5b5061493c85828601614887565b9150509250929050565b60006020828403121561495857600080fd5b813567ffffffffffffffff81111561496f57600080fd5b61497b84828501614542565b949350505050565b600081518084526020808501945080840160005b838110156149b357815187529582019590820190600101614997565b509495945050505050565b602081526000611c376020830184614983565b600080604083850312156149e457600080fd5b82359150602083013567ffffffffffffffff80821115614a0357600080fd5b908401906101208287031215614a1857600080fd5b614a2061442b565b82358152614a3060208401614784565b6020820152614a4160408401614513565b6040820152614a5260608401614513565b6060820152608083013582811115614a6957600080fd5b614a75888286016144a8565b60808301525060a083013582811115614a8d57600080fd5b614a99888286016145ec565b60a083015250614aac8760c085016145b1565b60c082015261010083013560e08201528093505050509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614b0a5783516001600160a01b031683529284019291840191600101614ae5565b50909695505050505050565b60008060408385031215614b2957600080fd5b8235614b34816144fe565b91506020830135614b44816144fe565b809150509250929050565b600081518084526020808501945080840160005b838110156149b35781516001600160a01b031687529582019590820190600101614b63565b602081526000611c376020830184614b4f565b600060408284031215614bad57600080fd5b6040516040810181811067ffffffffffffffff82111715614bd057614bd06143b0565b604052823581526020928301359281019290925250919050565b600060208284031215614bfc57600080fd5b813567ffffffffffffffff811115614c1357600080fd5b820160e08185031215611c3757600080fd5b600060208284031215614c3757600080fd5b5035919050565b60008083601f840112614c5057600080fd5b50813567ffffffffffffffff811115614c6857600080fd5b6020830191508360208260051b8501011115614c8357600080fd5b9250929050565b60008060008060008060008060e0898b031215614ca657600080fd5b606089018a811115614cb757600080fd5b8998503567ffffffffffffffff80821115614cd157600080fd5b818b0191508b601f830112614ce557600080fd5b813581811115614cf457600080fd5b8c6020828501011115614d0657600080fd5b6020830199508098505060808b0135915080821115614d2457600080fd5b614d308c838d01614c3e565b909750955060a08b0135915080821115614d4957600080fd5b50614d568b828c01614c3e565b999c989b50969995989497949560c00135949350505050565b600082601f830112614d8057600080fd5b81356020614d906144c78361451e565b82815260059290921b84018101918181019086841115614daf57600080fd5b8286015b848110156145a6578035614dc68161476e565b8352918301918301614db3565b600082601f830112614de457600080fd5b81356020614df46144c78361451e565b82815260059290921b84018101918181019086841115614e1357600080fd5b8286015b848110156145a657803567ffffffffffffffff811115614e375760008081fd5b614e458986838b01016144a8565b845250918301918301614e17565b801515811461085157600080fd5b8035613bf581614e53565b60008060408385031215614e7f57600080fd5b823567ffffffffffffffff80821115614e9757600080fd5b908401906101008287031215614eac57600080fd5b614eb461442b565b823582811115614ec357600080fd5b614ecf88828601614d6f565b825250602083013582811115614ee457600080fd5b614ef088828601614542565b602083015250604083013582811115614f0857600080fd5b614f1488828601614887565b604083015250606083013582811115614f2c57600080fd5b614f3888828601614dd3565b606083015250608083013582811115614f5057600080fd5b614f5c88828601614887565b60808301525060a083013560a082015260c083013582811115614f7e57600080fd5b614f8a88828601614887565b60c08301525060e083013560e0820152809450505050614fac60208401614e61565b90509250929050565b600060a08284031215614fc757600080fd5b60405160a0810181811067ffffffffffffffff82111715614fea57614fea6143b0565b6040528235614ff8816144fe565b8152602083013563ffffffff8116811461501157600080fd5b602082015260408301356150248161476e565b604082015260608301356150378161476e565b6060820152608083013561504a8161476e565b60808201529392505050565b803560ff81168114613bf557600080fd5b60008060008060008060c0878903121561508057600080fd5b863567ffffffffffffffff8082111561509857600080fd5b6150a48a838b01614542565b975060208901359150808211156150ba57600080fd5b6150c68a838b01614542565b96506150d460408a01615056565b955060608901359150808211156150ea57600080fd5b6150f68a838b016144a8565b945061510460808a01614784565b935060a089013591508082111561511a57600080fd5b5061512789828a016144a8565b9150509295509295509295565b60006020828403121561514657600080fd5b8151611c3781614e53565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036151e0576151e0615180565b5060010190565b600082198211156151fa576151fa615180565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561523757615237615180565b500290565b60008282101561524e5761524e615180565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff0382111561529f5761529f615180565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff8316806152e9576152e96152a7565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60006020828403121561536557600080fd5b8151611c37816144fe565b8051613bf58161476e565b8051613bf5816144fe565b600082601f83011261539757600080fd5b81516153a56144c782614480565b8181528460208386010111156153ba57600080fd5b61497b82602083016020870161481c565b6000604082840312156153dd57600080fd5b6153e56143df565b905081516153f2816144fe565b808252506020820151602082015292915050565b600082601f83011261541757600080fd5b815160206154276144c78361451e565b82815260069290921b8401810191818101908684111561544657600080fd5b8286015b848110156145a65761545c88826153cb565b83529183019160400161544a565b60006020828403121561547c57600080fd5b815167ffffffffffffffff8082111561549457600080fd5b9083019061012082860312156154a957600080fd5b6154b161442b565b825181526154c160208401615370565b60208201526154d26040840161537b565b60408201526154e36060840161537b565b60608201526080830151828111156154fa57600080fd5b61550687828601615386565b60808301525060a08301518281111561551e57600080fd5b61552a87828601615406565b60a08301525061553d8660c085016153cb565b60c0820152610100929092015160e0830152509392505050565b600082615566576155666152a7565b500490565b600063ffffffff80831681851680830382111561558a5761558a615180565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526155c38184018a614b4f565b905082810360808401526155d78189614b4f565b905060ff871660a084015282810360c08401526155f48187614848565b905067ffffffffffffffff851660e08401528281036101008401526156198185614848565b9c9b505050505050505050505050565b600081518084526020808501945080840160005b838110156149b35761566387835180516001600160a01b03168252602090810151910152565b604096909601959082019060010161563d565b60208152815160208201526000602083015160e0604084015261569d610100840182614848565b90506001600160a01b0360408501511660608401526060840151601f19808584030160808601526156ce8383614848565b925060808601519150808584030160a08601526156eb8383614b4f565b925060a08601519150808584030160c0860152506157098282615629565b91505060c084015160e08401528091505092915050565b600082601f83011261573157600080fd5b815160206157416144c78361451e565b82815260059290921b8401810191818101908684111561576057600080fd5b8286015b848110156145a65780516157778161476e565b8352918301918301615764565b600082601f83011261579557600080fd5b815160206157a56144c78361451e565b82815260059290921b840181019181810190868411156157c457600080fd5b8286015b848110156145a65780516157db816144fe565b83529183019183016157c8565b600082601f8301126157f957600080fd5b815160206158096144c78361451e565b82815260059290921b8401810191818101908684111561582857600080fd5b8286015b848110156145a6578051835291830191830161582c565b600082601f83011261585457600080fd5b815160206158646144c78361451e565b82815260059290921b8401810191818101908684111561588357600080fd5b8286015b848110156145a657805167ffffffffffffffff8111156158a75760008081fd5b6158b58986838b0101615386565b845250918301918301615887565b6000602082840312156158d557600080fd5b815167ffffffffffffffff808211156158ed57600080fd5b90830190610100828603121561590257600080fd5b61590a61442b565b82518281111561591957600080fd5b61592587828601615720565b82525060208301518281111561593a57600080fd5b61594687828601615784565b60208301525060408301518281111561595e57600080fd5b61596a878286016157e8565b60408301525060608301518281111561598257600080fd5b61598e87828601615843565b6060830152506080830151828111156159a657600080fd5b6159b2878286016157e8565b60808301525060a083015160a082015260c0830151828111156159d457600080fd5b6159e0878286016157e8565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b838110156149b357815167ffffffffffffffff1687529582019590820190600101615a10565b600081518084526020808501808196508360051b8101915082860160005b85811015615a7e578284038952615a6c848351614848565b98850198935090840190600101615a54565b5091979650505050505050565b6040815260008351610100806040850152615aaa6101408501836159fc565b915060208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080868503016060870152615ae68483614b4f565b93506040880151915080868503016080870152615b038483614983565b935060608801519150808685030160a0870152615b208483615a36565b935060808801519150808685030160c0870152615b3d8483614983565b935060a088015160e087015260c0880151915080868503018387015250615b648382614983565b60e088015161012087015286151560208701529350611c3792505050565b602081526000611c376020830184615629565b60a081526000615ba860a0830188614983565b8281036020840152615bba8188614983565b90508560408401528281036060840152615bd48186614983565b9150508260808301529695505050505050565b600060208284031215615bf957600080fd5b5051919050565b6000815160208301517fffffffff0000000000000000000000000000000000000000000000000000000080821693506004831015615c485780818460040360031b1b83161693505b505050919050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b166040850152816060850152615c8a8285018b614b4f565b91508382036080850152615c9e828a614b4f565b915060ff881660a085015283820360c0850152615cbb8288614848565b90861660e08501528381036101008501529050615619818561484856fea164736f6c634300080f000a",
}

var EVM2EVMTollOffRampABI = EVM2EVMTollOffRampMetaData.ABI

var EVM2EVMTollOffRampBin = EVM2EVMTollOffRampMetaData.Bin

func DeployEVM2EVMTollOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, blobVerifier common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address) (common.Address, *types.Transaction, *EVM2EVMTollOffRamp, error) {
	parsed, err := EVM2EVMTollOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMTollOffRampBin), backend, sourceChainId, chainId, offRampConfig, blobVerifier, afn, sourceTokens, pools, rateLimiterConfig, tokenLimitsAdmin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMTollOffRamp{EVM2EVMTollOffRampCaller: EVM2EVMTollOffRampCaller{contract: contract}, EVM2EVMTollOffRampTransactor: EVM2EVMTollOffRampTransactor{contract: contract}, EVM2EVMTollOffRampFilterer: EVM2EVMTollOffRampFilterer{contract: contract}}, nil
}

type EVM2EVMTollOffRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMTollOffRampCaller
	EVM2EVMTollOffRampTransactor
	EVM2EVMTollOffRampFilterer
}

type EVM2EVMTollOffRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampSession struct {
	Contract     *EVM2EVMTollOffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOffRampCallerSession struct {
	Contract *EVM2EVMTollOffRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMTollOffRampTransactorSession struct {
	Contract     *EVM2EVMTollOffRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOffRampRaw struct {
	Contract *EVM2EVMTollOffRamp
}

type EVM2EVMTollOffRampCallerRaw struct {
	Contract *EVM2EVMTollOffRampCaller
}

type EVM2EVMTollOffRampTransactorRaw struct {
	Contract *EVM2EVMTollOffRampTransactor
}

func NewEVM2EVMTollOffRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMTollOffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMTollOffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMTollOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRamp{address: address, abi: abi, EVM2EVMTollOffRampCaller: EVM2EVMTollOffRampCaller{contract: contract}, EVM2EVMTollOffRampTransactor: EVM2EVMTollOffRampTransactor{contract: contract}, EVM2EVMTollOffRampFilterer: EVM2EVMTollOffRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMTollOffRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMTollOffRampCaller, error) {
	contract, err := bindEVM2EVMTollOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampCaller{contract: contract}, nil
}

func NewEVM2EVMTollOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMTollOffRampTransactor, error) {
	contract, err := bindEVM2EVMTollOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampTransactor{contract: contract}, nil
}

func NewEVM2EVMTollOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMTollOffRampFilterer, error) {
	contract, err := bindEVM2EVMTollOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampFilterer{contract: contract}, nil
}

func bindEVM2EVMTollOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMTollOffRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOffRamp.Contract.EVM2EVMTollOffRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.EVM2EVMTollOffRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.EVM2EVMTollOffRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(AggregateRateLimiterInterfaceTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(AggregateRateLimiterInterfaceTokenBucket)).(*AggregateRateLimiterInterfaceTokenBucket)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMTollOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMTollOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMTollOffRamp.Contract.CcipReceive(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMTollOffRamp.Contract.CcipReceive(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) FeeTaken(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "feeTaken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) FeeTaken(arg0 *big.Int) (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.FeeTaken(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) FeeTaken(arg0 *big.Int) (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.FeeTaken(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetAFN(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetAFN(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetBlobVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getBlobVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetBlobVerifier() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetBlobVerifier(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetBlobVerifier() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetBlobVerifier(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getChainIDs")

	outstruct := new(GetChainIDs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMTollOffRamp.Contract.GetChainIDs(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMTollOffRamp.Contract.GetChainIDs(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOffRampInterfaceOffRampConfig)).(*BaseOffRampInterfaceOffRampConfig)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMTollOffRamp.Contract.GetConfig(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMTollOffRamp.Contract.GetConfig(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetDestinationToken(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetDestinationToken(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetDestinationTokens(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetDestinationTokens(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMTollOffRamp.Contract.GetExecutionState(&_EVM2EVMTollOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMTollOffRamp.Contract.GetExecutionState(&_EVM2EVMTollOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPool(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPool(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPoolTokens(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPoolTokens(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPricesForTokens(&_EVM2EVMTollOffRamp.CallOpts, tokens)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPricesForTokens(&_EVM2EVMTollOffRamp.CallOpts, tokens)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetRouter(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetRouter(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.IsAFNHealthy(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.IsAFNHealthy(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMTollOffRamp.Contract.LatestConfigDetails(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMTollOffRamp.Contract.LatestConfigDetails(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMTollOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMTollOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) OverheadGasToll(opts *bind.CallOpts, merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "overheadGasToll", merkleGasShare, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) OverheadGasToll(merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.OverheadGasToll(&_EVM2EVMTollOffRamp.CallOpts, merkleGasShare, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) OverheadGasToll(merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.OverheadGasToll(&_EVM2EVMTollOffRamp.CallOpts, merkleGasShare, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.Owner(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.Owner(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Paused() (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.Paused(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.Paused(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.Transmitters(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.Transmitters(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOffRamp.Contract.TypeAndVersion(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOffRamp.Contract.TypeAndVersion(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.AcceptOwnership(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.AcceptOwnership(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.AddPool(&_EVM2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.AddPool(&_EVM2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) Execute(opts *bind.TransactOpts, report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "execute", report, manualExecution)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Execute(&_EVM2EVMTollOffRamp.TransactOpts, report, manualExecution)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Execute(&_EVM2EVMTollOffRamp.TransactOpts, report, manualExecution)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "executeSingleMessage", message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRamp.TransactOpts, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRamp.TransactOpts, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Pause(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Pause(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.RemovePool(&_EVM2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.RemovePool(&_EVM2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetAFN(&_EVM2EVMTollOffRamp.TransactOpts, afn)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetAFN(&_EVM2EVMTollOffRamp.TransactOpts, afn)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetBlobVerifier(opts *bind.TransactOpts, blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setBlobVerifier", blobVerifier)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetBlobVerifier(&_EVM2EVMTollOffRamp.TransactOpts, blobVerifier)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetBlobVerifier(&_EVM2EVMTollOffRamp.TransactOpts, blobVerifier)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetConfig(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetConfig0(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setConfig0", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig0(&_EVM2EVMTollOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig0(&_EVM2EVMTollOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetPrices(&_EVM2EVMTollOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetPrices(&_EVM2EVMTollOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetRouter(&_EVM2EVMTollOffRamp.TransactOpts, router)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetRouter(&_EVM2EVMTollOffRamp.TransactOpts, router)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMTollOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMTollOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.TransferOwnership(&_EVM2EVMTollOffRamp.TransactOpts, to)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.TransferOwnership(&_EVM2EVMTollOffRamp.TransactOpts, to)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Transmit(&_EVM2EVMTollOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Transmit(&_EVM2EVMTollOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Unpause(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Unpause(&_EVM2EVMTollOffRamp.TransactOpts)
}

type EVM2EVMTollOffRampAFNSetIterator struct {
	Event *EVM2EVMTollOffRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampAFNSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampAFNSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampAFNSetIterator{contract: _EVM2EVMTollOffRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampAFNSet)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseAFNSet(log types.Log) (*EVM2EVMTollOffRampAFNSet, error) {
	event := new(EVM2EVMTollOffRampAFNSet)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampConfigChangedIterator struct {
	Event *EVM2EVMTollOffRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampConfigChangedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampConfigChanged)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMTollOffRampConfigChanged, error) {
	event := new(EVM2EVMTollOffRampConfigChanged)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampConfigSetIterator struct {
	Event *EVM2EVMTollOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampConfigSetIterator{contract: _EVM2EVMTollOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampConfigSet)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseConfigSet(log types.Log) (*EVM2EVMTollOffRampConfigSet, error) {
	event := new(EVM2EVMTollOffRampConfigSet)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampExecutionStateChangedIterator struct {
	Event *EVM2EVMTollOffRampExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampExecutionStateChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampExecutionStateChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampExecutionStateChanged struct {
	SequenceNumber uint64
	State          uint8
	Raw            types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMTollOffRampExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampExecutionStateChangedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampExecutionStateChanged)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseExecutionStateChanged(log types.Log) (*EVM2EVMTollOffRampExecutionStateChanged, error) {
	event := new(EVM2EVMTollOffRampExecutionStateChanged)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampOffRampConfigSetIterator struct {
	Event *EVM2EVMTollOffRampOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampOffRampConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampOffRampConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampOffRampConfigSet struct {
	Config BaseOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampOffRampConfigSetIterator{contract: _EVM2EVMTollOffRamp.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampOffRampConfigSet)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseOffRampConfigSet(log types.Log) (*EVM2EVMTollOffRampOffRampConfigSet, error) {
	event := new(EVM2EVMTollOffRampOffRampConfigSet)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampOffRampRouterSetIterator struct {
	Event *EVM2EVMTollOffRampOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampOffRampRouterSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampOffRampRouterSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampOffRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMTollOffRampOffRampRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampOffRampRouterSetIterator{contract: _EVM2EVMTollOffRamp.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampOffRampRouterSet)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseOffRampRouterSet(log types.Log) (*EVM2EVMTollOffRampOffRampRouterSet, error) {
	event := new(EVM2EVMTollOffRampOffRampRouterSet)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMTollOffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampOwnershipTransferRequestedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampOwnershipTransferRequested)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOffRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMTollOffRampOwnershipTransferRequested)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampOwnershipTransferredIterator struct {
	Event *EVM2EVMTollOffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampOwnershipTransferredIterator{contract: _EVM2EVMTollOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampOwnershipTransferred)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOffRampOwnershipTransferred, error) {
	event := new(EVM2EVMTollOffRampOwnershipTransferred)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampPausedIterator struct {
	Event *EVM2EVMTollOffRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampPausedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampPaused)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParsePaused(log types.Log) (*EVM2EVMTollOffRampPaused, error) {
	event := new(EVM2EVMTollOffRampPaused)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampPoolAddedIterator struct {
	Event *EVM2EVMTollOffRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampPoolAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampPoolAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampPoolAddedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampPoolAdded)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMTollOffRampPoolAdded, error) {
	event := new(EVM2EVMTollOffRampPoolAdded)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampPoolRemovedIterator struct {
	Event *EVM2EVMTollOffRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampPoolRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampPoolRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampPoolRemovedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampPoolRemoved)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMTollOffRampPoolRemoved, error) {
	event := new(EVM2EVMTollOffRampPoolRemoved)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampTokenPriceChangedIterator struct {
	Event *EVM2EVMTollOffRampTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampTokenPriceChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampTokenPriceChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampTokenPriceChangedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampTokenPriceChanged)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMTollOffRampTokenPriceChanged, error) {
	event := new(EVM2EVMTollOffRampTokenPriceChanged)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMTollOffRampTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampTokensRemovedFromBucket)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampTokensRemovedFromBucket)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampTokensRemovedFromBucketIterator{contract: _EVM2EVMTollOffRamp.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampTokensRemovedFromBucket)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMTollOffRampTokensRemovedFromBucket, error) {
	event := new(EVM2EVMTollOffRampTokensRemovedFromBucket)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampTransmittedIterator struct {
	Event *EVM2EVMTollOffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampTransmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampTransmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTransmittedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampTransmittedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTransmitted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampTransmitted)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseTransmitted(log types.Log) (*EVM2EVMTollOffRampTransmitted, error) {
	event := new(EVM2EVMTollOffRampTransmitted)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampUnpausedIterator struct {
	Event *EVM2EVMTollOffRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMTollOffRampUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMTollOffRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampUnpausedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampUnpaused)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMTollOffRampUnpaused, error) {
	event := new(EVM2EVMTollOffRampUnpaused)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetChainIDs struct {
	SourceChainId *big.Int
	ChainId       *big.Int
}
type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}
type LatestConfigDigestAndEpoch struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMTollOffRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMTollOffRamp.ParseAFNSet(log)
	case _EVM2EVMTollOffRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMTollOffRamp.ParseConfigChanged(log)
	case _EVM2EVMTollOffRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMTollOffRamp.ParseConfigSet(log)
	case _EVM2EVMTollOffRamp.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMTollOffRamp.ParseExecutionStateChanged(log)
	case _EVM2EVMTollOffRamp.abi.Events["OffRampConfigSet"].ID:
		return _EVM2EVMTollOffRamp.ParseOffRampConfigSet(log)
	case _EVM2EVMTollOffRamp.abi.Events["OffRampRouterSet"].ID:
		return _EVM2EVMTollOffRamp.ParseOffRampRouterSet(log)
	case _EVM2EVMTollOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMTollOffRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMTollOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMTollOffRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMTollOffRamp.abi.Events["Paused"].ID:
		return _EVM2EVMTollOffRamp.ParsePaused(log)
	case _EVM2EVMTollOffRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMTollOffRamp.ParsePoolAdded(log)
	case _EVM2EVMTollOffRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMTollOffRamp.ParsePoolRemoved(log)
	case _EVM2EVMTollOffRamp.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMTollOffRamp.ParseTokenPriceChanged(log)
	case _EVM2EVMTollOffRamp.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMTollOffRamp.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMTollOffRamp.abi.Events["Transmitted"].ID:
		return _EVM2EVMTollOffRamp.ParseTransmitted(log)
	case _EVM2EVMTollOffRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMTollOffRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMTollOffRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMTollOffRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMTollOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMTollOffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf")
}

func (EVM2EVMTollOffRampOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa0")
}

func (EVM2EVMTollOffRampOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
}

func (EVM2EVMTollOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMTollOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMTollOffRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMTollOffRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMTollOffRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMTollOffRampTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMTollOffRampTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMTollOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (EVM2EVMTollOffRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRamp) Address() common.Address {
	return _EVM2EVMTollOffRamp.address
}

type EVM2EVMTollOffRampInterface interface {
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error)

	CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error

	FeeTaken(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetBlobVerifier(opts *bind.CallOpts) (common.Address, error)

	GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

		error)

	GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	OverheadGasToll(opts *bind.CallOpts, merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetBlobVerifier(opts *bind.TransactOpts, blobVerifier common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error)

	SetConfig0(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMTollOffRampAFNSet, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMTollOffRampConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMTollOffRampConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMTollOffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMTollOffRampExecutionStateChanged, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*EVM2EVMTollOffRampOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMTollOffRampOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*EVM2EVMTollOffRampOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOffRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMTollOffRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMTollOffRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMTollOffRampPoolRemoved, error)

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMTollOffRampTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMTollOffRampTokensRemovedFromBucket, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMTollOffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMTollOffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
