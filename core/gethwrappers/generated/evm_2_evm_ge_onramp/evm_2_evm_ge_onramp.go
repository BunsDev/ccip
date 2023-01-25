// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_ge_onramp

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

type CommonEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

type GEConsumerEVM2AnyGEMessage struct {
	Receiver         []byte
	Data             []byte
	TokensAndAmounts []CommonEVMTokenAndAmount
	FeeToken         common.Address
	ExtraArgs        []byte
}

type GEEVM2EVMGEMessage struct {
	SourceChainId    uint64
	SequenceNumber   uint64
	FeeTokenAmount   *big.Int
	Sender           common.Address
	Nonce            uint64
	GasLimit         *big.Int
	Strict           bool
	Receiver         common.Address
	Data             []byte
	TokensAndAmounts []CommonEVMTokenAndAmount
	FeeToken         common.Address
	MessageId        [32]byte
}

type IAggregateRateLimiterRateLimiterConfig struct {
	Rate     *big.Int
	Capacity *big.Int
}

type IAggregateRateLimiterTokenBucket struct {
	Rate        *big.Int
	Capacity    *big.Int
	Tokens      *big.Int
	LastUpdated *big.Int
}

type IBaseOnRampOnRampConfig struct {
	CommitFeeJuels  uint64
	MaxDataSize     uint32
	MaxTokensLength uint16
	MaxGasLimit     uint64
}

type IEVM2EVMGEOnRampDynamicFeeConfig struct {
	LinkToken       common.Address
	FeeAmount       *big.Int
	DestGasOverhead *big.Int
	Multiplier      *big.Int
	FeeManager      common.Address
	DestChainId     uint64
}

var EVM2EVMGEOnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"commitFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structIBaseOnRamp.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"},{\"internalType\":\"contractIGERouter\",\"name\":\"router\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destGasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMGEOnRamp.DynamicFeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"expected\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"got\",\"type\":\"bytes4\"}],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chain\",\"type\":\"uint64\"}],\"name\":\"TokenOrChainNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"AllowListSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structGE.EVM2EVMGEMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeAdmin\",\"type\":\"address\"}],\"name\":\"FeeAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destGasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMGEOnRamp.DynamicFeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"FeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"commitFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIBaseOnRamp.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OnRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structGEConsumer.EVM2AnyGEMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structGEConsumer.EVM2AnyGEMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOnRampConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"commitFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structIBaseOnRamp.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"setAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAllowlistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAdmin\",\"type\":\"address\"}],\"name\":\"setFeeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destGasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMGEOnRamp.DynamicFeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"setFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"commitFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structIBaseOnRamp.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOnRampConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200460e3803806200460e833981016040819052620000349162000aa0565b6000805460ff191681558b908b908b908b908b908b908b908b908b908b9083908390889088903390819081620000b15760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000eb57620000eb8162000644565b5050506001600160a01b0381166200011657604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580511562000168576002805460ff60a01b1916600160a01b179055805162000166906003906020840190620006f5565b505b60005b8151811015620001d55760016004600084848151811062000190576200019062000bdd565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055620001cd8162000bf3565b90506200016b565b505080600560006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060800160405280836000015181526020018360200151815260200183602001518152602001428152506008600082015181600001556020820151816001015560408201518160020155606082015181600301559050505050896001600160401b03166080816001600160401b031681525050886001600160401b031660a0816001600160401b03168152505083600d60008201518160000160006101000a8154816001600160401b0302191690836001600160401b0316021790555060208201518160000160086101000a81548163ffffffff021916908363ffffffff160217905550604082015181600001600c6101000a81548161ffff021916908361ffff160217905550606082015181600001600e6101000a8154816001600160401b0302191690836001600160401b0316021790555090505080600e60006101000a8154816001600160a01b0302191690836001600160a01b031602179055506000600c60006101000a8154816001600160401b0302191690836001600160401b031602179055508651885114620003a95760405162d8548360e71b815260040160405180910390fd5b8751620003be9060109060208b0190620006f5565b5060005b885181101562000485576040518060400160405280898381518110620003ec57620003ec62000bdd565b60200260200101516001600160a01b0316815260200160011515815250600f60008b848151811062000422576200042262000bdd565b6020908102919091018101516001600160a01b0390811683528282019390935260409091016000208351815494909201511515600160a01b026001600160a81b031990941691909216179190911790556200047d8162000bf3565b9050620003c2565b505050505050505050505080601260008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060a08201518160040160146101000a8154816001600160401b0302191690836001600160401b031602179055509050507fba22a5847647789e6efe1840c86bc66129ac89e03d7b95e0eebdf7fa43763fdd8b8b306040516020016200059c94939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b60408051808303601f19018152828252805160209182012060c090815284516001600160a01b03908116855285830151928501929092528483015192840192909252606080850151908401526080808501519091169083015260a0808401516001600160401b0316908301527f92c1d3bc951c7322787dfee144ba08e1dbc36535f211ae5fbdbbc3c27c3b26be910160405180910390a1505050505050505050505062000c1b565b336001600160a01b038216036200069e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a8565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200074d579160200282015b828111156200074d57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000716565b506200075b9291506200075f565b5090565b5b808211156200075b576000815560010162000760565b80516001600160401b03811681146200078e57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620007d457620007d462000793565b604052919050565b60006001600160401b03821115620007f857620007f862000793565b5060051b60200190565b6001600160a01b03811681146200081857600080fd5b50565b80516200078e8162000802565b600082601f8301126200083a57600080fd5b81516020620008536200084d83620007dc565b620007a9565b82815260059290921b840181019181810190868411156200087357600080fd5b8286015b848110156200089b5780516200088d8162000802565b835291830191830162000877565b509695505050505050565b600082601f830112620008b857600080fd5b81516020620008cb6200084d83620007dc565b82815260059290921b84018101918181019086841115620008eb57600080fd5b8286015b848110156200089b578051620009058162000802565b8352918301918301620008ef565b6000608082840312156200092657600080fd5b604051608081016001600160401b03811182821017156200094b576200094b62000793565b6040529050806200095c8362000776565b8152602083015163ffffffff811681146200097657600080fd5b6020820152604083015161ffff811681146200099157600080fd5b6040820152620009a46060840162000776565b60608201525092915050565b600060408284031215620009c357600080fd5b604080519081016001600160401b0381118282101715620009e857620009e862000793565b604052825181526020928301519281019290925250919050565b600060c0828403121562000a1557600080fd5b60405160c081016001600160401b038111828210171562000a3a5762000a3a62000793565b8060405250809150825162000a4f8162000802565b80825250602083015160208201526040830151604082015260608301516060820152608083015162000a818162000802565b608082015262000a9460a0840162000776565b60a08201525092915050565b60008060008060008060008060008060006102808c8e03121562000ac357600080fd5b62000ace8c62000776565b9a5062000ade60208d0162000776565b60408d0151909a506001600160401b0381111562000afb57600080fd5b62000b098e828f0162000828565b60608e0151909a5090506001600160401b0381111562000b2857600080fd5b62000b368e828f01620008a6565b60808e015190995090506001600160401b0381111562000b5557600080fd5b62000b638e828f0162000828565b97505062000b7460a08d016200081b565b955062000b858d60c08e0162000913565b945062000b978d6101408e01620009b0565b935062000ba86101808d016200081b565b925062000bb96101a08d016200081b565b915062000bcb8d6101c08e0162000a02565b90509295989b509295989b9093969950565b634e487b7160e01b600052603260045260246000fd5b60006001820162000c1457634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c0516139bc62000c526000396000611813015260006103b5015260008181610303015261162b01526139bc6000f3fe608060405234801561001057600080fd5b506004361061025c5760003560e01c8063744b92e211610145578063b0f479a1116100bd578063d7644ba21161008c578063eb511dd411610071578063eb511dd414610651578063f2fde38b14610664578063f78faa321461067757600080fd5b8063d7644ba21461062b578063dcd210f11461063e57600080fd5b8063b0f479a1146105ea578063c0d78655146105fb578063c5eff3d01461060e578063d3c7c2c71461062357600080fd5b80638da5cb5b1161011457806391872543116100f957806391872543146105b1578063a11855be146105c4578063a7d3e02f146105d757600080fd5b80638da5cb5b1461056057806390c2339b1461057657600080fd5b8063744b92e21461050a57806379ba50971461051d5780638456cb5914610525578063856c82471461052d57600080fd5b80634352fa9f116101d8578063552b818b116101a75780635c975abb1161018c5780635c975abb146104d95780635d86f141146104e45780636eb2d031146104f757600080fd5b8063552b818b146104b5578063599f6431146104c857600080fd5b80634352fa9f146103805780634741062e146103935780634894b536146103b3578063499bb0e2146103d957600080fd5b80633408e4701161022f57806339aa92641161021457806339aa92641461035d5780633f4ba83a146103705780634120fccd1461037857600080fd5b80633408e4701461030157806338724a951461033c57600080fd5b8063108ee5fc14610261578063147809b314610276578063181f5a77146102935780632222dd42146102dc575b600080fd5b61027461026f366004612c1d565b610689565b005b61027e610740565b60405190151581526020015b60405180910390f35b6102cf6040518060400160405280601581526020017f45564d3245564d47454f6e52616d7020312e302e30000000000000000000000081525081565b60405161028a9190612c96565b6002546001600160a01b03165b6040516001600160a01b03909116815260200161028a565b7f00000000000000000000000000000000000000000000000000000000000000005b60405167ffffffffffffffff909116815260200161028a565b61034f61034a366004612cc1565b6107cd565b60405190815260200161028a565b61027461036b366004612c1d565b61097e565b6102746109b5565b6103236109c7565b61027461038e366004612e17565b6109e7565b6103a66103a1366004612ed2565b610c3c565b60405161028a9190612f07565b7f0000000000000000000000000000000000000000000000000000000000000000610323565b61046b6040805160808101825260008082526020820181905291810182905260608101919091525060408051608081018252600d5467ffffffffffffffff808216835263ffffffff68010000000000000000830416602084015261ffff6c01000000000000000000000000830416938301939093526e0100000000000000000000000000009004909116606082015290565b60405161028a9190815167ffffffffffffffff908116825260208084015163ffffffff169083015260408084015161ffff1690830152606092830151169181019190915260800190565b6102746104c3366004612f4b565b610d04565b6005546001600160a01b03166102e9565b60005460ff1661027e565b6102e96104f2366004612c1d565b610e8e565b610274610505366004612c1d565b610e9f565b610274610518366004612fc0565b610f09565b610274611059565b61027461113c565b61032361053b366004612c1d565b6001600160a01b031660009081526011602052604090205467ffffffffffffffff1690565b60005461010090046001600160a01b03166102e9565b61057e61114c565b60405161028a91908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6102746105bf366004612ff9565b6111ed565b6102746105d236600461302b565b611319565b61034f6105e536600461303d565b611360565b600e546001600160a01b03166102e9565b610274610609366004612c1d565b611885565b6106166118e8565b60405161028a9190613098565b61061661194a565b6102746106393660046130e7565b611954565b61027461064c366004613104565b6119c4565b61027461065f366004612fc0565b611a0b565b610274610672366004612c1d565b611bac565b600254600160a01b900460ff1661027e565b610691611bc0565b6001600160a01b0381166106d1576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa1580156107a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c79190613116565b15905090565b6000806107e56107e06080850185613133565b611c1f565b516016549091506000906001600160a01b031663268e5d4861080d6080870160608801612c1d565b60165460405160e084901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b039092166004830152600160a01b900467ffffffffffffffff166024820152604401602060405180830381865afa158015610882573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108a6919061319f565b6fffffffffffffffffffffffffffffffff1690508060000361092e576108d26080850160608601612c1d565b6016546040517f102e3c280000000000000000000000000000000000000000000000000000000081526001600160a01b039092166004830152600160a01b900467ffffffffffffffff1660248201526044015b60405180910390fd5b601554601454670de0b6b3a76400009190839061094b9086613200565b6109559190613218565b61095f9190613218565b6109699190613255565b6013546109769190613200565b949350505050565b610986611bc0565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6109bd611bc0565b6109c5611d85565b565b600c546000906109e29067ffffffffffffffff166001613290565b905090565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015610a2457506005546001600160a01b03163314155b15610a5b576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610a97576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610af1576006600060078381548110610abc57610abc6132bc565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610aea816132eb565b9050610a9d565b5060005b82811015610c21576000858281518110610b1157610b116132bc565b6020026020010151905060006001600160a01b0316816001600160a01b031603610b67576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610b7957610b796132bc565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610bde57610bde6132bc565b6020026020010151604051610c089291906001600160a01b03929092168252602082015260400190565b60405180910390a150610c1a816132eb565b9050610af5565b508351610c35906007906020870190612b21565b5050505050565b80516060908067ffffffffffffffff811115610c5a57610c5a612cf6565b604051908082528060200260200182016040528015610c83578160200160208202803683370190505b50915060005b81811015610cfd5760066000858381518110610ca757610ca76132bc565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610ce257610ce26132bc565b6020908102919091010152610cf6816132eb565b9050610c89565b5050919050565b610d0c611bc0565b60006003805480602002602001604051908101604052809291908181526020018280548015610d6457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d46575b5050505050905060005b8151811015610dd157600060046000848481518110610d8f57610d8f6132bc565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055610dca816132eb565b9050610d6e565b50610dde60038484612b93565b5060005b82811015610e4f57600160046000868685818110610e0257610e026132bc565b9050602002016020810190610e179190612c1d565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055610e48816132eb565b9050610de2565b507ff8adc5fee247b62a85f63294cb46e4be61da815e4756bc57672a83b24faf0dda8383604051610e81929190613323565b60405180910390a1505050565b6000610e9982611e21565b92915050565b610ea7611bc0565b6017805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f34efeb1f04731080ec2147b3b9c7e38f9b884e035020914e40269450f53b4815906020015b60405180910390a150565b610f11611bc0565b6001600160a01b038281166000908152600f6020908152604091829020825180840190935254928316808352600160a01b90930460ff1615159082015290610f90576040517f73913ebd0000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401610925565b816001600160a01b031681600001516001600160a01b031614610fdf576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038381166000818152600f602090815260409182902080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690558151928352928516928201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610e81565b6001546001600160a01b031633146110b35760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610925565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b611144611bc0565b6109c5611ec4565b6111776040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b546060820181905242906000906111b19083613366565b602084015184519192506111dd916111c99084613218565b85604001516111d89190613200565b611f4c565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561122a57506005546001600160a01b03163314155b15611261576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff116112b5576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112bf6008611f62565b602081015160098190558151600855600a546112db9190611f4c565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d9101610efe565b611321611bc0565b80601261132e8282613393565b9050507f92c1d3bc951c7322787dfee144ba08e1dbc36535f211ae5fbdbbc3c27c3b26be81604051610efe9190613479565b6000805460ff16156113b45760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610925565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611407573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061142b9190613116565b15611461576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e546001600160a01b031633146114a5576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6012546001600160a01b03166114c16080860160608701612c1d565b6001600160a01b03160361156d5760006114e46104f26080870160608801612c1d565b90506001600160a01b038116611542576115046080860160608701612c1d565b6040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610925565b61156781856115576080890160608a01612c1d565b6001600160a01b0316919061200f565b5061158e565b60165461158e906001600160a01b0316846115576080880160608901612c1d565b60006115a06107e06080870187613133565b905061161b6115b26020870187613133565b83519091506115c460408901896134ee565b808060200260200160405190810160405280939291908181526020016000905b828210156116105761160160408302860136819003810190613556565b815260200190600101906115e4565b505050505086612094565b60006040518061018001604052807f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168152602001600c600081819054906101000a900467ffffffffffffffff1661167b90613590565b825467ffffffffffffffff9182166101009390930a8381029083021990911617909255825260208083018990526001600160a01b038816604080850182905260009182526011909252908120805460609094019390926116db9116613590565b825467ffffffffffffffff9182166101009390930a8381029202191617909155815283516020808301919091528401511515604082015260600161171f8880613133565b81019061172c9190612c1d565b6001600160a01b0316815260200187806020019061174a9190613133565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200161179160408901896134ee565b808060200260200160405190810160405280939291908181526020016000905b828210156117dd576117ce60408302860136819003810190613556565b815260200190600101906117b1565b50505091835250506020016117f86080890160608a01612c1d565b6001600160a01b0316815260006020909101529050611837817f00000000000000000000000000000000000000000000000000000000000000006123c8565b6101608201526040517faffc45517195d6499808c643bd4a7b0ffeedf95bea5852840d7bfcf63f59e8219061186d908390613606565b60405180910390a161016001519150505b9392505050565b61188d611bc0565b600e805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527fc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d1590602001610efe565b6060600380548060200260200160405190810160405280929190818152602001828054801561194057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611922575b5050505050905090565b60606109e26124d2565b61195c611bc0565b60028054821515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9091161790556040517fccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df03290610efe90831515815260200190565b6119cc611bc0565b80600d6119d98282613747565b9050507fe8e69f40b790527d400ff1d06e78519a73e7725dc6e5c04f263cc7758143c4ba81604051610efe9190613883565b611a13611bc0565b6001600160a01b0382161580611a3057506001600160a01b038116155b15611a67576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166000908152600f6020526040902054600160a01b900460ff1615611ac1576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182526001600160a01b03838116808352600160208085018281528885166000818152600f84528881209751885493511515600160a01b027fffffffffffffffffffffff000000000000000000000000000000000000000000909416971696909617919091179095556010805492830181559093527f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae67201805473ffffffffffffffffffffffffffffffffffffffff1916841790558351928352908201527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c9101610734565b611bb4611bc0565b611bbd8161266d565b50565b60005461010090046001600160a01b031633146109c55760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610925565b60408051808201909152600080825260208201526000829003611c5957506040805180820190915262030d40815260006020820152610e99565b7f97a657c900000000000000000000000000000000000000000000000000000000611c886004600085876138f5565b611c919161391f565b7fffffffff000000000000000000000000000000000000000000000000000000001614611d4e577f97a657c900000000000000000000000000000000000000000000000000000000611ce76004600085876138f5565b611cf09161391f565b6040517f55a0e02c0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000928316600482015291166024820152604401610925565b6040805180820190915280611d676024600486886138f5565b810190611d749190613967565b815260006020909101529392505050565b60005460ff16611dd75760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610925565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6001600160a01b038181166000908152600f602090815260408083208151808301909252549384168152600160a01b90930460ff1615801591840191909152909190611e875750506001600160a01b039081166000908152600f60205260409020541690565b6040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401610925565b60005460ff1615611f175760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610925565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258611e043390565b6000818310611f5b578161187e565b5090919050565b6001810154600282015442911480611f7d5750808260030154145b15611f86575050565b816001015482600201541115611fc8576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482611fda9190613366565b6001840154845491925061200191611ff29084613218565b85600201546111d89190613200565b600284015550600390910155565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261208f908490612729565b505050565b600e546001600160a01b03166120d6576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038116612116576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5468010000000000000000900463ffffffff1684111561218257600d546040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090910463ffffffff16600482015260248101859052604401610925565b600d546e010000000000000000000000000000900467ffffffffffffffff168311156121da576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151600d546c01000000000000000000000000900461ffff1681111561222c576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254600160a01b900460ff16801561225e57506001600160a01b03821660009081526004602052604090205460ff16155b156122a0576040517fd0d259760000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610925565b6122a98361280e565b60005b818110156123c05760008482815181106122c8576122c86132bc565b6020026020010151905060008160000151905060006122e682610e8e565b90506001600160a01b038116612333576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610925565b60208301516040517f503c285800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0382169063503c285890602401600060405180830381600087803b15801561239457600080fd5b505af11580156123a8573d6000803e3d6000fd5b50505050505050806123b9906132eb565b90506122ac565b505050505050565b60008060001b828460200151856080015186606001518760e00151886101000151805190602001208961012001516040516020016124069190613980565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d604001516040516020016124b49c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b60606000805b60105481101561254c57600f6000601083815481106124f9576124f96132bc565b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff600160a01b909104161561253c5781612538816132eb565b9250505b612545816132eb565b90506124d8565b5060008167ffffffffffffffff81111561256857612568612cf6565b604051908082528060200260200182016040528015612591578160200160208202803683370190505b5090506000805b60105481101561266457600f6000601083815481106125b9576125b96132bc565b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff600160a01b90910416156126545760108181548110612601576126016132bc565b6000918252602090912001546001600160a01b03168383612621816132eb565b945081518110612633576126336132bc565b60200260200101906001600160a01b031690816001600160a01b0316815250505b61265d816132eb565b9050612598565b50909392505050565b336001600160a01b038216036126c55760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610925565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600061277e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316612a169092919063ffffffff16565b80519091501561208f578080602001905181019061279c9190613116565b61208f5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610925565b6000805b825181101561290d57600060066000858481518110612833576128336132bc565b6020026020010151600001516001600160a01b03166001600160a01b03168152602001908152602001600020549050806000036128c65783828151811061287c5761287c6132bc565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610925565b8382815181106128d8576128d86132bc565b602002602001015160200151816128ef9190613218565b6128f99084613200565b92505080612906906132eb565b9050612812565b508015612a125761291e6008611f62565b600954811115612968576009546040517f688ccf77000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610925565b600a548111156129c857600854600a54600091906129869084613366565b6129909190613255565b9050806040517fe31e0f3200000000000000000000000000000000000000000000000000000000815260040161092591815260200190565b80600860020160008282546129dd9190613366565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba37890602001610734565b5050565b6060610976848460008585843b612a6f5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610925565b600080866001600160a01b03168587604051612a8b9190613993565b60006040518083038185875af1925050503d8060008114612ac8576040519150601f19603f3d011682016040523d82523d6000602084013e612acd565b606091505b5091509150612add828286612ae8565b979650505050505050565b60608315612af757508161187e565b825115612b075782518084602001fd5b8160405162461bcd60e51b81526004016109259190612c96565b828054828255906000526020600020908101928215612b83579160200282015b82811115612b83578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190612b41565b50612b8f929150612bf3565b5090565b828054828255906000526020600020908101928215612b83579160200282015b82811115612b8357815473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03843516178255602090920191600190910190612bb3565b5b80821115612b8f5760008155600101612bf4565b6001600160a01b0381168114611bbd57600080fd5b600060208284031215612c2f57600080fd5b813561187e81612c08565b60005b83811015612c55578181015183820152602001612c3d565b83811115612c64576000848401525b50505050565b60008151808452612c82816020860160208601612c3a565b601f01601f19169290920160200192915050565b60208152600061187e6020830184612c6a565b600060a08284031215612cbb57600080fd5b50919050565b600060208284031215612cd357600080fd5b813567ffffffffffffffff811115612cea57600080fd5b61097684828501612ca9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715612d4857612d48612cf6565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715612d7757612d77612cf6565b604052919050565b600067ffffffffffffffff821115612d9957612d99612cf6565b5060051b60200190565b600082601f830112612db457600080fd5b81356020612dc9612dc483612d7f565b612d4e565b82815260059290921b84018101918181019086841115612de857600080fd5b8286015b84811015612e0c578035612dff81612c08565b8352918301918301612dec565b509695505050505050565b60008060408385031215612e2a57600080fd5b823567ffffffffffffffff80821115612e4257600080fd5b612e4e86838701612da3565b9350602091508185013581811115612e6557600080fd5b85019050601f81018613612e7857600080fd5b8035612e86612dc482612d7f565b81815260059190911b82018301908381019088831115612ea557600080fd5b928401925b82841015612ec357833582529284019290840190612eaa565b80955050505050509250929050565b600060208284031215612ee457600080fd5b813567ffffffffffffffff811115612efb57600080fd5b61097684828501612da3565b6020808252825182820181905260009190848201906040850190845b81811015612f3f57835183529284019291840191600101612f23565b50909695505050505050565b60008060208385031215612f5e57600080fd5b823567ffffffffffffffff80821115612f7657600080fd5b818501915085601f830112612f8a57600080fd5b813581811115612f9957600080fd5b8660208260051b8501011115612fae57600080fd5b60209290920196919550909350505050565b60008060408385031215612fd357600080fd5b8235612fde81612c08565b91506020830135612fee81612c08565b809150509250929050565b60006040828403121561300b57600080fd5b613013612d25565b82358152602083013560208201528091505092915050565b600060c08284031215612cbb57600080fd5b60008060006060848603121561305257600080fd5b833567ffffffffffffffff81111561306957600080fd5b61307586828701612ca9565b93505060208401359150604084013561308d81612c08565b809150509250925092565b6020808252825182820181905260009190848201906040850190845b81811015612f3f5783516001600160a01b0316835292840192918401916001016130b4565b8015158114611bbd57600080fd5b6000602082840312156130f957600080fd5b813561187e816130d9565b600060808284031215612cbb57600080fd5b60006020828403121561312857600080fd5b815161187e816130d9565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261316857600080fd5b83018035915067ffffffffffffffff82111561318357600080fd5b60200191503681900382131561319857600080fd5b9250929050565b6000602082840312156131b157600080fd5b81516fffffffffffffffffffffffffffffffff8116811461187e57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115613213576132136131d1565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615613250576132506131d1565b500290565b60008261328b577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600067ffffffffffffffff8083168185168083038211156132b3576132b36131d1565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361331c5761331c6131d1565b5060010190565b60208082528181018390526000908460408401835b86811015612e0c57823561334b81612c08565b6001600160a01b031682529183019190830190600101613338565b600082821015613378576133786131d1565b500390565b67ffffffffffffffff81168114611bbd57600080fd5b813561339e81612c08565b815473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038216178255506020820135600182015560408201356002820155606082013560038201556004810160808301356133f581612c08565b815473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0382161782555060a083013561342a8161337d565b81547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a09190911b7bffffffffffffffff0000000000000000000000000000000000000000161790555050565b60c08101823561348881612c08565b6001600160a01b038082168452602085013560208501526040850135604085015260608501356060850152608085013591506134c382612c08565b16608083015260a08301356134d78161337d565b67ffffffffffffffff811660a08401525092915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261352357600080fd5b83018035915067ffffffffffffffff82111561353e57600080fd5b6020019150600681901b360382131561319857600080fd5b60006040828403121561356857600080fd5b613570612d25565b823561357b81612c08565b81526020928301359281019290925250919050565b600067ffffffffffffffff8083168181036135ad576135ad6131d1565b6001019392505050565b600081518084526020808501945080840160005b838110156135fb57815180516001600160a01b0316885283015183880152604090960195908201906001016135cb565b509495945050505050565b6020815261362160208201835167ffffffffffffffff169052565b6000602083015161363e604084018267ffffffffffffffff169052565b5060408301516060830152606083015161366360808401826001600160a01b03169052565b50608083015167ffffffffffffffff811660a08401525060a083015160c083015260c083015161369760e084018215159052565b5060e08301516101006136b4818501836001600160a01b03169052565b8085015191505061018061012081818601526136d46101a0860184612c6a565b9250808601519050610140601f1986850301818701526136f484836135b7565b935080870151915050610160613714818701836001600160a01b03169052565b959095015193019290925250919050565b63ffffffff81168114611bbd57600080fd5b61ffff81168114611bbd57600080fd5b81356137528161337d565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008216178355602084013561379681613725565b6bffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffffffffffff00000000000000000000000084161717845560408501356137e181613737565b6dffff0000000000000000000000008160601b169050837fffffffffffffffffffffffffffffffffffff0000000000000000000000000000841617935080848317178555606086013592506138358361337d565b75ffffffffffffffff00000000000000000000000000008360701b167fffffffffffffffffffff0000000000000000ffffffffffffffffffffffffffff851683178217178555505050505050565b6080810182356138928161337d565b67ffffffffffffffff90811683526020840135906138af82613725565b63ffffffff82166020850152604085013591506138cb82613737565b61ffff82166040850152606085013591506138e58261337d565b8082166060850152505092915050565b6000808585111561390557600080fd5b8386111561391257600080fd5b5050820193919092039150565b7fffffffff00000000000000000000000000000000000000000000000000000000813581811691600485101561395f5780818660040360031b1b83161692505b505092915050565b60006020828403121561397957600080fd5b5035919050565b60208152600061187e60208301846135b7565b600082516139a5818460208701612c3a565b919091019291505056fea164736f6c634300080f000a",
}

var EVM2EVMGEOnRampABI = EVM2EVMGEOnRampMetaData.ABI

var EVM2EVMGEOnRampBin = EVM2EVMGEOnRampMetaData.Bin

func DeployEVM2EVMGEOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, chainId uint64, destinationChainId uint64, tokens []common.Address, pools []common.Address, allowlist []common.Address, afn common.Address, config IBaseOnRampOnRampConfig, rateLimiterConfig IAggregateRateLimiterRateLimiterConfig, tokenLimitsAdmin common.Address, router common.Address, feeConfig IEVM2EVMGEOnRampDynamicFeeConfig) (common.Address, *types.Transaction, *EVM2EVMGEOnRamp, error) {
	parsed, err := EVM2EVMGEOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMGEOnRampBin), backend, chainId, destinationChainId, tokens, pools, allowlist, afn, config, rateLimiterConfig, tokenLimitsAdmin, router, feeConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMGEOnRamp{EVM2EVMGEOnRampCaller: EVM2EVMGEOnRampCaller{contract: contract}, EVM2EVMGEOnRampTransactor: EVM2EVMGEOnRampTransactor{contract: contract}, EVM2EVMGEOnRampFilterer: EVM2EVMGEOnRampFilterer{contract: contract}}, nil
}

type EVM2EVMGEOnRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMGEOnRampCaller
	EVM2EVMGEOnRampTransactor
	EVM2EVMGEOnRampFilterer
}

type EVM2EVMGEOnRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMGEOnRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMGEOnRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMGEOnRampSession struct {
	Contract     *EVM2EVMGEOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMGEOnRampCallerSession struct {
	Contract *EVM2EVMGEOnRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMGEOnRampTransactorSession struct {
	Contract     *EVM2EVMGEOnRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMGEOnRampRaw struct {
	Contract *EVM2EVMGEOnRamp
}

type EVM2EVMGEOnRampCallerRaw struct {
	Contract *EVM2EVMGEOnRampCaller
}

type EVM2EVMGEOnRampTransactorRaw struct {
	Contract *EVM2EVMGEOnRampTransactor
}

func NewEVM2EVMGEOnRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMGEOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMGEOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMGEOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRamp{address: address, abi: abi, EVM2EVMGEOnRampCaller: EVM2EVMGEOnRampCaller{contract: contract}, EVM2EVMGEOnRampTransactor: EVM2EVMGEOnRampTransactor{contract: contract}, EVM2EVMGEOnRampFilterer: EVM2EVMGEOnRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMGEOnRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMGEOnRampCaller, error) {
	contract, err := bindEVM2EVMGEOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampCaller{contract: contract}, nil
}

func NewEVM2EVMGEOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMGEOnRampTransactor, error) {
	contract, err := bindEVM2EVMGEOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampTransactor{contract: contract}, nil
}

func NewEVM2EVMGEOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMGEOnRampFilterer, error) {
	contract, err := bindEVM2EVMGEOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampFilterer{contract: contract}, nil
}

func bindEVM2EVMGEOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMGEOnRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMGEOnRamp.Contract.EVM2EVMGEOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.EVM2EVMGEOnRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.EVM2EVMGEOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMGEOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(IAggregateRateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(IAggregateRateLimiterTokenBucket)).(*IAggregateRateLimiterTokenBucket)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMGEOnRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMGEOnRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetAFN() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetAFN(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetAFN(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetAllowlist(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getAllowlist")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetAllowlist() ([]common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetAllowlist(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetAllowlist() ([]common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetAllowlist(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getAllowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetAllowlistEnabled() (bool, error) {
	return _EVM2EVMGEOnRamp.Contract.GetAllowlistEnabled(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetAllowlistEnabled() (bool, error) {
	return _EVM2EVMGEOnRamp.Contract.GetAllowlistEnabled(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetChainId() (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetChainId(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetChainId() (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetChainId(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetDestinationChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getDestinationChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetDestinationChainId() (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetDestinationChainId(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetDestinationChainId() (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetDestinationChainId(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getExpectedNextSequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetFee(opts *bind.CallOpts, message GEConsumerEVM2AnyGEMessage) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getFee", message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetFee(message GEConsumerEVM2AnyGEMessage) (*big.Int, error) {
	return _EVM2EVMGEOnRamp.Contract.GetFee(&_EVM2EVMGEOnRamp.CallOpts, message)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetFee(message GEConsumerEVM2AnyGEMessage) (*big.Int, error) {
	return _EVM2EVMGEOnRamp.Contract.GetFee(&_EVM2EVMGEOnRamp.CallOpts, message)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetOnRampConfig(opts *bind.CallOpts) (IBaseOnRampOnRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getOnRampConfig")

	if err != nil {
		return *new(IBaseOnRampOnRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IBaseOnRampOnRampConfig)).(*IBaseOnRampOnRampConfig)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetOnRampConfig() (IBaseOnRampOnRampConfig, error) {
	return _EVM2EVMGEOnRamp.Contract.GetOnRampConfig(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetOnRampConfig() (IBaseOnRampOnRampConfig, error) {
	return _EVM2EVMGEOnRamp.Contract.GetOnRampConfig(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getPoolBySourceToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetPoolBySourceToken(&_EVM2EVMGEOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetPoolBySourceToken(&_EVM2EVMGEOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMGEOnRamp.Contract.GetPricesForTokens(&_EVM2EVMGEOnRamp.CallOpts, tokens)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMGEOnRamp.Contract.GetPricesForTokens(&_EVM2EVMGEOnRamp.CallOpts, tokens)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetRouter() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetRouter(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetRouter(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getSenderNonce", sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetSenderNonce(&_EVM2EVMGEOnRamp.CallOpts, sender)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetSenderNonce(&_EVM2EVMGEOnRamp.CallOpts, sender)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetSupportedTokens(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetSupportedTokens(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMGEOnRamp.Contract.IsAFNHealthy(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMGEOnRamp.Contract.IsAFNHealthy(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) Owner() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.Owner(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.Owner(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) Paused() (bool, error) {
	return _EVM2EVMGEOnRamp.Contract.Paused(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMGEOnRamp.Contract.Paused(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMGEOnRamp.Contract.TypeAndVersion(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMGEOnRamp.Contract.TypeAndVersion(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.AcceptOwnership(&_EVM2EVMGEOnRamp.TransactOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.AcceptOwnership(&_EVM2EVMGEOnRamp.TransactOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.AddPool(&_EVM2EVMGEOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.AddPool(&_EVM2EVMGEOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) ForwardFromRouter(opts *bind.TransactOpts, message GEConsumerEVM2AnyGEMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "forwardFromRouter", message, feeTokenAmount, originalSender)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) ForwardFromRouter(message GEConsumerEVM2AnyGEMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.ForwardFromRouter(&_EVM2EVMGEOnRamp.TransactOpts, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) ForwardFromRouter(message GEConsumerEVM2AnyGEMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.ForwardFromRouter(&_EVM2EVMGEOnRamp.TransactOpts, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.Pause(&_EVM2EVMGEOnRamp.TransactOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.Pause(&_EVM2EVMGEOnRamp.TransactOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.RemovePool(&_EVM2EVMGEOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.RemovePool(&_EVM2EVMGEOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetAFN(&_EVM2EVMGEOnRamp.TransactOpts, afn)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetAFN(&_EVM2EVMGEOnRamp.TransactOpts, afn)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setAllowlist", allowlist)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetAllowlist(&_EVM2EVMGEOnRamp.TransactOpts, allowlist)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetAllowlist(&_EVM2EVMGEOnRamp.TransactOpts, allowlist)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setAllowlistEnabled", enabled)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetAllowlistEnabled(&_EVM2EVMGEOnRamp.TransactOpts, enabled)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetAllowlistEnabled(&_EVM2EVMGEOnRamp.TransactOpts, enabled)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetFeeAdmin(opts *bind.TransactOpts, feeAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setFeeAdmin", feeAdmin)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetFeeAdmin(feeAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetFeeAdmin(&_EVM2EVMGEOnRamp.TransactOpts, feeAdmin)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetFeeAdmin(feeAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetFeeAdmin(&_EVM2EVMGEOnRamp.TransactOpts, feeAdmin)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetFeeConfig(opts *bind.TransactOpts, feeConfig IEVM2EVMGEOnRampDynamicFeeConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setFeeConfig", feeConfig)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetFeeConfig(feeConfig IEVM2EVMGEOnRampDynamicFeeConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetFeeConfig(&_EVM2EVMGEOnRamp.TransactOpts, feeConfig)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetFeeConfig(feeConfig IEVM2EVMGEOnRampDynamicFeeConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetFeeConfig(&_EVM2EVMGEOnRamp.TransactOpts, feeConfig)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetOnRampConfig(opts *bind.TransactOpts, config IBaseOnRampOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setOnRampConfig", config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetOnRampConfig(config IBaseOnRampOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetOnRampConfig(&_EVM2EVMGEOnRamp.TransactOpts, config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetOnRampConfig(config IBaseOnRampOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetOnRampConfig(&_EVM2EVMGEOnRamp.TransactOpts, config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetPrices(&_EVM2EVMGEOnRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetPrices(&_EVM2EVMGEOnRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetRateLimiterConfig(&_EVM2EVMGEOnRamp.TransactOpts, config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetRateLimiterConfig(&_EVM2EVMGEOnRamp.TransactOpts, config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetRouter(&_EVM2EVMGEOnRamp.TransactOpts, router)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetRouter(&_EVM2EVMGEOnRamp.TransactOpts, router)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMGEOnRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMGEOnRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.TransferOwnership(&_EVM2EVMGEOnRamp.TransactOpts, to)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.TransferOwnership(&_EVM2EVMGEOnRamp.TransactOpts, to)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.Unpause(&_EVM2EVMGEOnRamp.TransactOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.Unpause(&_EVM2EVMGEOnRamp.TransactOpts)
}

type EVM2EVMGEOnRampAFNSetIterator struct {
	Event *EVM2EVMGEOnRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampAFNSet)
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
		it.Event = new(EVM2EVMGEOnRampAFNSet)
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

func (it *EVM2EVMGEOnRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampAFNSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampAFNSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseAFNSet(log types.Log) (*EVM2EVMGEOnRampAFNSet, error) {
	event := new(EVM2EVMGEOnRampAFNSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampAllowListEnabledSetIterator struct {
	Event *EVM2EVMGEOnRampAllowListEnabledSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampAllowListEnabledSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampAllowListEnabledSet)
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
		it.Event = new(EVM2EVMGEOnRampAllowListEnabledSet)
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

func (it *EVM2EVMGEOnRampAllowListEnabledSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampAllowListEnabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampAllowListEnabledSet struct {
	Enabled bool
	Raw     types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterAllowListEnabledSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampAllowListEnabledSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "AllowListEnabledSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampAllowListEnabledSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "AllowListEnabledSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchAllowListEnabledSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampAllowListEnabledSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "AllowListEnabledSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampAllowListEnabledSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "AllowListEnabledSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseAllowListEnabledSet(log types.Log) (*EVM2EVMGEOnRampAllowListEnabledSet, error) {
	event := new(EVM2EVMGEOnRampAllowListEnabledSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "AllowListEnabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampAllowListSetIterator struct {
	Event *EVM2EVMGEOnRampAllowListSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampAllowListSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampAllowListSet)
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
		it.Event = new(EVM2EVMGEOnRampAllowListSet)
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

func (it *EVM2EVMGEOnRampAllowListSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampAllowListSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampAllowListSet struct {
	Allowlist []common.Address
	Raw       types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterAllowListSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampAllowListSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "AllowListSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampAllowListSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "AllowListSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchAllowListSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampAllowListSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "AllowListSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampAllowListSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "AllowListSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseAllowListSet(log types.Log) (*EVM2EVMGEOnRampAllowListSet, error) {
	event := new(EVM2EVMGEOnRampAllowListSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "AllowListSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampCCIPSendRequestedIterator struct {
	Event *EVM2EVMGEOnRampCCIPSendRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampCCIPSendRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampCCIPSendRequested)
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
		it.Event = new(EVM2EVMGEOnRampCCIPSendRequested)
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

func (it *EVM2EVMGEOnRampCCIPSendRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampCCIPSendRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampCCIPSendRequested struct {
	Message GEEVM2EVMGEMessage
	Raw     types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMGEOnRampCCIPSendRequestedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampCCIPSendRequestedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "CCIPSendRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampCCIPSendRequested) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampCCIPSendRequested)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseCCIPSendRequested(log types.Log) (*EVM2EVMGEOnRampCCIPSendRequested, error) {
	event := new(EVM2EVMGEOnRampCCIPSendRequested)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampConfigChangedIterator struct {
	Event *EVM2EVMGEOnRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampConfigChanged)
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
		it.Event = new(EVM2EVMGEOnRampConfigChanged)
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

func (it *EVM2EVMGEOnRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMGEOnRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampConfigChangedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampConfigChanged)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMGEOnRampConfigChanged, error) {
	event := new(EVM2EVMGEOnRampConfigChanged)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampFeeAdminSetIterator struct {
	Event *EVM2EVMGEOnRampFeeAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampFeeAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampFeeAdminSet)
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
		it.Event = new(EVM2EVMGEOnRampFeeAdminSet)
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

func (it *EVM2EVMGEOnRampFeeAdminSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampFeeAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampFeeAdminSet struct {
	FeeAdmin common.Address
	Raw      types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterFeeAdminSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampFeeAdminSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "FeeAdminSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampFeeAdminSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "FeeAdminSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchFeeAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampFeeAdminSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "FeeAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampFeeAdminSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "FeeAdminSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseFeeAdminSet(log types.Log) (*EVM2EVMGEOnRampFeeAdminSet, error) {
	event := new(EVM2EVMGEOnRampFeeAdminSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "FeeAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampFeeConfigSetIterator struct {
	Event *EVM2EVMGEOnRampFeeConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampFeeConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampFeeConfigSet)
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
		it.Event = new(EVM2EVMGEOnRampFeeConfigSet)
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

func (it *EVM2EVMGEOnRampFeeConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampFeeConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampFeeConfigSet struct {
	FeeConfig IEVM2EVMGEOnRampDynamicFeeConfig
	Raw       types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampFeeConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "FeeConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampFeeConfigSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "FeeConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampFeeConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "FeeConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampFeeConfigSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "FeeConfigSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseFeeConfigSet(log types.Log) (*EVM2EVMGEOnRampFeeConfigSet, error) {
	event := new(EVM2EVMGEOnRampFeeConfigSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "FeeConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampOnRampConfigSetIterator struct {
	Event *EVM2EVMGEOnRampOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampOnRampConfigSet)
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
		it.Event = new(EVM2EVMGEOnRampOnRampConfigSet)
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

func (it *EVM2EVMGEOnRampOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampOnRampConfigSet struct {
	Config IBaseOnRampOnRampConfig
	Raw    types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterOnRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampOnRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampOnRampConfigSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "OnRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampOnRampConfigSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseOnRampConfigSet(log types.Log) (*EVM2EVMGEOnRampOnRampConfigSet, error) {
	event := new(EVM2EVMGEOnRampOnRampConfigSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMGEOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMGEOnRampOwnershipTransferRequested)
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

func (it *EVM2EVMGEOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMGEOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampOwnershipTransferRequestedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampOwnershipTransferRequested)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMGEOnRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMGEOnRampOwnershipTransferRequested)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampOwnershipTransferredIterator struct {
	Event *EVM2EVMGEOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMGEOnRampOwnershipTransferred)
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

func (it *EVM2EVMGEOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMGEOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampOwnershipTransferredIterator{contract: _EVM2EVMGEOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampOwnershipTransferred)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMGEOnRampOwnershipTransferred, error) {
	event := new(EVM2EVMGEOnRampOwnershipTransferred)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampPausedIterator struct {
	Event *EVM2EVMGEOnRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampPaused)
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
		it.Event = new(EVM2EVMGEOnRampPaused)
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

func (it *EVM2EVMGEOnRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMGEOnRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampPausedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampPaused)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParsePaused(log types.Log) (*EVM2EVMGEOnRampPaused, error) {
	event := new(EVM2EVMGEOnRampPaused)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampPoolAddedIterator struct {
	Event *EVM2EVMGEOnRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampPoolAdded)
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
		it.Event = new(EVM2EVMGEOnRampPoolAdded)
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

func (it *EVM2EVMGEOnRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMGEOnRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampPoolAddedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampPoolAdded)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMGEOnRampPoolAdded, error) {
	event := new(EVM2EVMGEOnRampPoolAdded)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampPoolRemovedIterator struct {
	Event *EVM2EVMGEOnRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampPoolRemoved)
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
		it.Event = new(EVM2EVMGEOnRampPoolRemoved)
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

func (it *EVM2EVMGEOnRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMGEOnRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampPoolRemovedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampPoolRemoved)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMGEOnRampPoolRemoved, error) {
	event := new(EVM2EVMGEOnRampPoolRemoved)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampRouterSetIterator struct {
	Event *EVM2EVMGEOnRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampRouterSet)
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
		it.Event = new(EVM2EVMGEOnRampRouterSet)
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

func (it *EVM2EVMGEOnRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterRouterSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampRouterSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampRouterSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "RouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampRouterSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampRouterSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseRouterSet(log types.Log) (*EVM2EVMGEOnRampRouterSet, error) {
	event := new(EVM2EVMGEOnRampRouterSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampTokenPriceChangedIterator struct {
	Event *EVM2EVMGEOnRampTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampTokenPriceChanged)
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
		it.Event = new(EVM2EVMGEOnRampTokenPriceChanged)
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

func (it *EVM2EVMGEOnRampTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMGEOnRampTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampTokenPriceChangedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampTokenPriceChanged)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMGEOnRampTokenPriceChanged, error) {
	event := new(EVM2EVMGEOnRampTokenPriceChanged)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMGEOnRampTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMGEOnRampTokensRemovedFromBucket)
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

func (it *EVM2EVMGEOnRampTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMGEOnRampTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampTokensRemovedFromBucketIterator{contract: _EVM2EVMGEOnRamp.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampTokensRemovedFromBucket)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMGEOnRampTokensRemovedFromBucket, error) {
	event := new(EVM2EVMGEOnRampTokensRemovedFromBucket)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampUnpausedIterator struct {
	Event *EVM2EVMGEOnRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampUnpaused)
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
		it.Event = new(EVM2EVMGEOnRampUnpaused)
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

func (it *EVM2EVMGEOnRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMGEOnRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampUnpausedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampUnpaused)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMGEOnRampUnpaused, error) {
	event := new(EVM2EVMGEOnRampUnpaused)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMGEOnRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMGEOnRamp.ParseAFNSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["AllowListEnabledSet"].ID:
		return _EVM2EVMGEOnRamp.ParseAllowListEnabledSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["AllowListSet"].ID:
		return _EVM2EVMGEOnRamp.ParseAllowListSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["CCIPSendRequested"].ID:
		return _EVM2EVMGEOnRamp.ParseCCIPSendRequested(log)
	case _EVM2EVMGEOnRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMGEOnRamp.ParseConfigChanged(log)
	case _EVM2EVMGEOnRamp.abi.Events["FeeAdminSet"].ID:
		return _EVM2EVMGEOnRamp.ParseFeeAdminSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["FeeConfigSet"].ID:
		return _EVM2EVMGEOnRamp.ParseFeeConfigSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["OnRampConfigSet"].ID:
		return _EVM2EVMGEOnRamp.ParseOnRampConfigSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMGEOnRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMGEOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMGEOnRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMGEOnRamp.abi.Events["Paused"].ID:
		return _EVM2EVMGEOnRamp.ParsePaused(log)
	case _EVM2EVMGEOnRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMGEOnRamp.ParsePoolAdded(log)
	case _EVM2EVMGEOnRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMGEOnRamp.ParsePoolRemoved(log)
	case _EVM2EVMGEOnRamp.abi.Events["RouterSet"].ID:
		return _EVM2EVMGEOnRamp.ParseRouterSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMGEOnRamp.ParseTokenPriceChanged(log)
	case _EVM2EVMGEOnRamp.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMGEOnRamp.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMGEOnRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMGEOnRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMGEOnRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMGEOnRampAllowListEnabledSet) Topic() common.Hash {
	return common.HexToHash("0xccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df032")
}

func (EVM2EVMGEOnRampAllowListSet) Topic() common.Hash {
	return common.HexToHash("0xf8adc5fee247b62a85f63294cb46e4be61da815e4756bc57672a83b24faf0dda")
}

func (EVM2EVMGEOnRampCCIPSendRequested) Topic() common.Hash {
	return common.HexToHash("0xaffc45517195d6499808c643bd4a7b0ffeedf95bea5852840d7bfcf63f59e821")
}

func (EVM2EVMGEOnRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMGEOnRampFeeAdminSet) Topic() common.Hash {
	return common.HexToHash("0x34efeb1f04731080ec2147b3b9c7e38f9b884e035020914e40269450f53b4815")
}

func (EVM2EVMGEOnRampFeeConfigSet) Topic() common.Hash {
	return common.HexToHash("0x92c1d3bc951c7322787dfee144ba08e1dbc36535f211ae5fbdbbc3c27c3b26be")
}

func (EVM2EVMGEOnRampOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xe8e69f40b790527d400ff1d06e78519a73e7725dc6e5c04f263cc7758143c4ba")
}

func (EVM2EVMGEOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMGEOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMGEOnRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMGEOnRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMGEOnRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMGEOnRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0xc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15")
}

func (EVM2EVMGEOnRampTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMGEOnRampTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMGEOnRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRamp) Address() common.Address {
	return _EVM2EVMGEOnRamp.address
}

type EVM2EVMGEOnRampInterface interface {
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetAllowlist(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	GetChainId(opts *bind.CallOpts) (uint64, error)

	GetDestinationChainId(opts *bind.CallOpts) (uint64, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetFee(opts *bind.CallOpts, message GEConsumerEVM2AnyGEMessage) (*big.Int, error)

	GetOnRampConfig(opts *bind.CallOpts) (IBaseOnRampOnRampConfig, error)

	GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error)

	GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, message GEConsumerEVM2AnyGEMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error)

	SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error)

	SetFeeAdmin(opts *bind.TransactOpts, feeAdmin common.Address) (*types.Transaction, error)

	SetFeeConfig(opts *bind.TransactOpts, feeConfig IEVM2EVMGEOnRampDynamicFeeConfig) (*types.Transaction, error)

	SetOnRampConfig(opts *bind.TransactOpts, config IBaseOnRampOnRampConfig) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMGEOnRampAFNSet, error)

	FilterAllowListEnabledSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampAllowListEnabledSetIterator, error)

	WatchAllowListEnabledSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampAllowListEnabledSet) (event.Subscription, error)

	ParseAllowListEnabledSet(log types.Log) (*EVM2EVMGEOnRampAllowListEnabledSet, error)

	FilterAllowListSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampAllowListSetIterator, error)

	WatchAllowListSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampAllowListSet) (event.Subscription, error)

	ParseAllowListSet(log types.Log) (*EVM2EVMGEOnRampAllowListSet, error)

	FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMGEOnRampCCIPSendRequestedIterator, error)

	WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampCCIPSendRequested) (event.Subscription, error)

	ParseCCIPSendRequested(log types.Log) (*EVM2EVMGEOnRampCCIPSendRequested, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMGEOnRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMGEOnRampConfigChanged, error)

	FilterFeeAdminSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampFeeAdminSetIterator, error)

	WatchFeeAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampFeeAdminSet) (event.Subscription, error)

	ParseFeeAdminSet(log types.Log) (*EVM2EVMGEOnRampFeeAdminSet, error)

	FilterFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampFeeConfigSetIterator, error)

	WatchFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampFeeConfigSet) (event.Subscription, error)

	ParseFeeConfigSet(log types.Log) (*EVM2EVMGEOnRampFeeConfigSet, error)

	FilterOnRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampOnRampConfigSetIterator, error)

	WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampOnRampConfigSet) (event.Subscription, error)

	ParseOnRampConfigSet(log types.Log) (*EVM2EVMGEOnRampOnRampConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMGEOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMGEOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMGEOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMGEOnRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMGEOnRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMGEOnRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMGEOnRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMGEOnRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMGEOnRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMGEOnRampPoolRemoved, error)

	FilterRouterSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampRouterSetIterator, error)

	WatchRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampRouterSet) (event.Subscription, error)

	ParseRouterSet(log types.Log) (*EVM2EVMGEOnRampRouterSet, error)

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMGEOnRampTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMGEOnRampTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMGEOnRampTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMGEOnRampTokensRemovedFromBucket, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMGEOnRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMGEOnRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
