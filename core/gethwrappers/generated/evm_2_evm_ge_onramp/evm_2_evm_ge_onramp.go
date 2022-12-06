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

type BaseOnRampInterfaceOnRampConfig struct {
	CommitFeeJuels  uint64
	MaxDataSize     uint64
	MaxTokensLength uint64
	MaxGasLimit     uint64
}

type CommonEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

type EVM2EVMGEOnRampInterfaceDynamicFeeConfig struct {
	FeeToken        common.Address
	FeeAmount       *big.Int
	DestGasOverhead *big.Int
	Multiplier      *big.Int
	GasFeeCache     common.Address
	DestChainId     *big.Int
}

type GEConsumerEVM2AnyGEMessage struct {
	Receiver         []byte
	Data             []byte
	TokensAndAmounts []CommonEVMTokenAndAmount
	FeeToken         common.Address
	ExtraArgs        []byte
}

type GEEVM2EVMGEMessage struct {
	SourceChainId    *big.Int
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

var EVM2EVMGEOnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"commitFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"},{\"internalType\":\"contractGERouterInterface\",\"name\":\"router\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destGasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"internalType\":\"structEVM2EVMGEOnRampInterface.DynamicFeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"expected\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"got\",\"type\":\"bytes4\"}],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"got\",\"type\":\"address\"}],\"name\":\"MismatchedFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"AllowListSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structGE.EVM2EVMGEMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeAdmin\",\"type\":\"address\"}],\"name\":\"FeeAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destGasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMGEOnRampInterface.DynamicFeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"FeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"commitFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OnRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structGEConsumer.EVM2AnyGEMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"commitFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structGEConsumer.EVM2AnyGEMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_destinationChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"setAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAllowlistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"commitFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAdmin\",\"type\":\"address\"}],\"name\":\"setFeeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destGasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"internalType\":\"structEVM2EVMGEOnRampInterface.DynamicFeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"setFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200400e3803806200400e833981016040819052620000349162000993565b6000805460ff191681558b908b908b908b908b908b908b908b908b908b9083908390889088903390819081620000b15760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000eb57620000eb8162000552565b5050506001600160a01b0381166200011657604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580511562000168576002805460ff60a01b1916600160a01b17905580516200016690600390602084019062000603565b505b60005b8151811015620001d55760016004600084848151811062000190576200019062000abe565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055620001cd8162000ad4565b90506200016b565b5050600580546001600160a01b039283166001600160a01b03199182161790915560408051608080820183528551808352602096870180518885018190529051848601819052426060958601819052600893909355600991909155600a55600b558f905260a08e90528851600d8054968b0151938b0151928b01516001600160401b03908116600160c01b026001600160c01b03948216600160801b02949094166001600160801b0395821668010000000000000000026001600160801b0319909916919093161796909617929092169190911717909255600e805491841691909216179055600c80546001600160401b03191690558651885114620002ed5760405162d8548360e71b815260040160405180910390fd5b8751620003029060109060208b019062000603565b5060005b8851811015620003c957604051806040016040528089838151811062000330576200033062000abe565b60200260200101516001600160a01b0316815260200160011515815250600f60008b848151811062000366576200036662000abe565b6020908102919091018101516001600160a01b0390811683528282019390935260409091016000208351815494909201511515600160a01b026001600160a81b03199094169190921617919091179055620003c18162000ad4565b905062000306565b505050505050505050505080601260008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060a082015181600501559050507fba22a5847647789e6efe1840c86bc66129ac89e03d7b95e0eebdf7fa43763fdd8b8b30604051602001620004b39493929190938452602084019290925260408301526001600160a01b0316606082015260800190565b60408051808303601f19018152828252805160209182012060c090815284516001600160a01b03908116855285830151928501929092528483015192840192909252606080850151908401526080808501519091169083015260a080840151908301527f690597a8e557c5794b70ef51c386875de46b2ac98f234cc214b803785c1e7d84910160405180910390a1505050505050505050505062000afc565b336001600160a01b03821603620005ac5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a8565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200065b579160200282015b828111156200065b57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000624565b50620006699291506200066d565b5090565b5b808211156200066957600081556001016200066e565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620006c557620006c562000684565b604052919050565b60006001600160401b03821115620006e957620006e962000684565b5060051b60200190565b6001600160a01b03811681146200070957600080fd5b50565b600082601f8301126200071e57600080fd5b81516020620007376200073183620006cd565b6200069a565b82815260059290921b840181019181810190868411156200075757600080fd5b8286015b848110156200077f5780516200077181620006f3565b83529183019183016200075b565b509695505050505050565b600082601f8301126200079c57600080fd5b81516020620007af6200073183620006cd565b82815260059290921b84018101918181019086841115620007cf57600080fd5b8286015b848110156200077f578051620007e981620006f3565b8352918301918301620007d3565b80516200080481620006f3565b919050565b80516001600160401b03811681146200080457600080fd5b6000608082840312156200083457600080fd5b604051608081016001600160401b038111828210171562000859576200085962000684565b6040529050806200086a8362000809565b81526200087a6020840162000809565b60208201526200088d6040840162000809565b6040820152620008a06060840162000809565b60608201525092915050565b600060408284031215620008bf57600080fd5b604080519081016001600160401b0381118282101715620008e457620008e462000684565b604052825181526020928301519281019290925250919050565b600060c082840312156200091157600080fd5b60405160c081016001600160401b038111828210171562000936576200093662000684565b806040525080915082516200094b81620006f3565b8082525060208301516020820152604083015160408201526060830151606082015260808301516200097d81620006f3565b608082015260a092830151920191909152919050565b60008060008060008060008060008060006102808c8e031215620009b657600080fd5b8b5160208d015160408e0151919c509a506001600160401b03811115620009dc57600080fd5b620009ea8e828f016200070c565b60608e0151909a5090506001600160401b0381111562000a0957600080fd5b62000a178e828f016200078a565b60808e015190995090506001600160401b0381111562000a3657600080fd5b62000a448e828f016200070c565b97505062000a5560a08d01620007f7565b955062000a668d60c08e0162000821565b945062000a788d6101408e01620008ac565b935062000a896101808d01620007f7565b925062000a9a6101a08d01620007f7565b915062000aac8d6101c08e01620008fe565b90509295989b509295989b9093969950565b634e487b7160e01b600052603260045260246000fd5b60006001820162000af557634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c0516134db62000b3360003960006118dd015260006104a60152600081816102eb015261170201526134db6000f3fe608060405234801561001057600080fd5b50600436106102415760003560e01c8063744b92e211610145578063b0f479a1116100bd578063c5eff3d01161008c578063eb511dd411610071578063eb511dd41461060b578063f2fde38b1461061e578063f78faa321461063157600080fd5b8063c5eff3d0146105f0578063d7644ba2146105f857600080fd5b8063b0f479a1146104db578063bcddbbd3146104ec578063c0d78655146104ff578063c3f909d41461051257600080fd5b80638da5cb5b1161011457806391872543116100f9578063918725431461048e578063a7217195146104a1578063a7d3e02f146104c857600080fd5b80638da5cb5b1461043d57806390c2339b1461045357600080fd5b8063744b92e21461040557806379ba5097146104185780638456cb591461042057806389c065681461042857600080fd5b80634120fccd116101d8578063599f6431116101a75780635d86f1411161018c5780635d86f141146103cc5780636eb2d031146103df57806371239ff4146103f257600080fd5b8063599f6431146103b05780635c975abb146103c157600080fd5b80634120fccd146103495780634352fa9f1461036a5780634741062e1461037d578063552b818b1461039d57600080fd5b806328094b591161021457806328094b59146102e657806338724a951461031b57806339aa92641461032e5780633f4ba83a1461034157600080fd5b8063108ee5fc14610246578063147809b31461025b578063181f5a77146102785780632222dd42146102c1575b600080fd5b610259610254366004612848565b610643565b005b6102636106fa565b60405190151581526020015b60405180910390f35b6102b46040518060400160405280601581526020017f45564d3245564d47454f6e52616d7020312e302e30000000000000000000000081525081565b60405161026f91906128b2565b6002546001600160a01b03165b6040516001600160a01b03909116815260200161026f565b61030d7f000000000000000000000000000000000000000000000000000000000000000081565b60405190815260200161026f565b61030d6103293660046128dd565b610787565b61025961033c366004612848565b610913565b61025961094a565b61035161095c565b60405167ffffffffffffffff909116815260200161026f565b610259610378366004612a33565b61097c565b61039061038b366004612aee565b610bd1565b60405161026f9190612b23565b6102596103ab366004612b67565b610c99565b6005546001600160a01b03166102ce565b60005460ff16610263565b6102ce6103da366004612848565b610e23565b6102596103ed366004612848565b610ec6565b610259610400366004612bdc565b610f30565b610259610413366004612bee565b610f77565b6102596110c7565b6102596111aa565b6104306111ba565b60405161026f9190612c27565b60005461010090046001600160a01b03166102ce565b61045b611355565b60405161026f91908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b61025961049c366004612c68565b6113f6565b61030d7f000000000000000000000000000000000000000000000000000000000000000081565b61030d6104d6366004612c9a565b611522565b600e546001600160a01b03166102ce565b6102596104fa366004612cf5565b61194d565b61025961050d366004612848565b611994565b6105ac6040805160808101825260008082526020820181905291810182905260608101919091525060408051608081018252600d5467ffffffffffffffff808216835268010000000000000000820481166020840152700100000000000000000000000000000000820481169383019390935278010000000000000000000000000000000000000000000000009004909116606082015290565b60405161026f9190815167ffffffffffffffff9081168252602080840151821690830152604080840151821690830152606092830151169181019190915260800190565b6104306119f7565b610259610606366004612d15565b611a59565b610259610619366004612bee565b611ac9565b61025961062c366004612848565b611c6a565b600254600160a01b900460ff16610263565b61064b611c7e565b6001600160a01b03811661068b576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa15801561075d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107819190612d32565b15905090565b60006107996080830160608401612848565b6012546001600160a01b03908116911614610814576012546001600160a01b03166107ca6080840160608501612848565b6040517fc8201c920000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152911660248201526044015b60405180910390fd5b600061082b6108266080850185612d4f565b611cdd565b516016546017546040517ffcee45f40000000000000000000000000000000000000000000000000000000081529293506000926001600160a01b039092169163fcee45f4916108809160040190815260200190565b602060405180830381865afa15801561089d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c19190612dbb565b601554601454919250670de0b6b3a76400009183906108e09086612e03565b6108ea9190612e1b565b6108f49190612e1b565b6108fe9190612e58565b60135461090b9190612e03565b949350505050565b61091b611c7e565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b610952611c7e565b61095a611e45565b565b600c546000906109779067ffffffffffffffff166001612e93565b905090565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b0316141580156109b957506005546001600160a01b03163314155b156109f0576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610a2c576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610a86576006600060078381548110610a5157610a51612ebf565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610a7f81612eee565b9050610a32565b5060005b82811015610bb6576000858281518110610aa657610aa6612ebf565b6020026020010151905060006001600160a01b0316816001600160a01b031603610afc576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610b0e57610b0e612ebf565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610b7357610b73612ebf565b6020026020010151604051610b9d9291906001600160a01b03929092168252602082015260400190565b60405180910390a150610baf81612eee565b9050610a8a565b508351610bca90600790602087019061274c565b5050505050565b80516060908067ffffffffffffffff811115610bef57610bef612912565b604051908082528060200260200182016040528015610c18578160200160208202803683370190505b50915060005b81811015610c925760066000858381518110610c3c57610c3c612ebf565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610c7757610c77612ebf565b6020908102919091010152610c8b81612eee565b9050610c1e565b5050919050565b610ca1611c7e565b60006003805480602002602001604051908101604052809291908181526020018280548015610cf957602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610cdb575b5050505050905060005b8151811015610d6657600060046000848481518110610d2457610d24612ebf565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055610d5f81612eee565b9050610d03565b50610d73600384846127be565b5060005b82811015610de457600160046000868685818110610d9757610d97612ebf565b9050602002016020810190610dac9190612848565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055610ddd81612eee565b9050610d77565b507ff8adc5fee247b62a85f63294cb46e4be61da815e4756bc57672a83b24faf0dda8383604051610e16929190612f26565b60405180910390a1505050565b6001600160a01b038181166000908152600f602090815260408083208151808301909252549384168152600160a01b90930460ff1615801591840191909152909190610e895750506001600160a01b039081166000908152600f60205260409020541690565b6040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b038416600482015260240161080b565b610ece611c7e565b6018805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f34efeb1f04731080ec2147b3b9c7e38f9b884e035020914e40269450f53b4815906020015b60405180910390a150565b610f38611c7e565b806012610f458282612f69565b9050507f690597a8e557c5794b70ef51c386875de46b2ac98f234cc214b803785c1e7d8481604051610f259190613000565b610f7f611c7e565b6001600160a01b038281166000908152600f6020908152604091829020825180840190935254928316808352600160a01b90930460ff1615159082015290610ffe576040517f73913ebd0000000000000000000000000000000000000000000000000000000081526001600160a01b038416600482015260240161080b565b816001600160a01b031681600001516001600160a01b03161461104d576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038381166000818152600f602090815260409182902080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690558151928352928516928201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610e16565b6001546001600160a01b031633146111215760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161080b565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6111b2611c7e565b61095a611ee1565b60606000805b60105481101561123457600f6000601083815481106111e1576111e1612ebf565b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff600160a01b9091041615611224578161122081612eee565b9250505b61122d81612eee565b90506111c0565b5060008167ffffffffffffffff81111561125057611250612912565b604051908082528060200260200182016040528015611279578160200160208202803683370190505b5090506000805b60105481101561134c57600f6000601083815481106112a1576112a1612ebf565b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff600160a01b909104161561133c57601081815481106112e9576112e9612ebf565b6000918252602090912001546001600160a01b0316838361130981612eee565b94508151811061131b5761131b612ebf565b60200260200101906001600160a01b031690816001600160a01b0316815250505b61134581612eee565b9050611280565b50909392505050565b6113806040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b546060820181905242906000906113ba9083613061565b602084015184519192506113e6916113d29084612e1b565b85604001516113e19190612e03565b611f69565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561143357506005546001600160a01b03163314155b1561146a576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff116114be576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114c86008611f81565b602081015160098190558151600855600a546114e49190611f69565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d9101610f25565b6000805460ff16156115765760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161080b565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ed9190612d32565b15611623576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e546001600160a01b03163314611667576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006116796108266080870187612d4f565b90506116f461168b6020870187612d4f565b835190915061169d6040890189613078565b808060200260200160405190810160405280939291908181526020016000905b828210156116e9576116da604083028601368190038101906130e0565b815260200190600101906116bd565b50505050508661202e565b6040805161018081019091527f00000000000000000000000000000000000000000000000000000000000000008152600c805460009291602083019184906117459067ffffffffffffffff1661311a565b825467ffffffffffffffff9182166101009390930a8381029083021990911617909255825260208083018990526001600160a01b038816604080850182905260009182526011909252908120805460609094019390926117a5911661311a565b825467ffffffffffffffff9182166101009390930a838102920219161790915581528351602080830191909152840151151560408201526060016117e98880612d4f565b8101906117f69190612848565b6001600160a01b031681526020018780602001906118149190612d4f565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200161185b6040890189613078565b808060200260200160405190810160405280939291908181526020016000905b828210156118a757611898604083028601368190038101906130e0565b8152602001906001019061187b565b50505091835250506020016118c26080890160608a01612848565b6001600160a01b0316815260006020909101529050611901817f000000000000000000000000000000000000000000000000000000000000000061237e565b6101608201526040517fd94df26131d6ea08cb76088e9db81d712d59856f9a5872cac5459a9ad822e65590611937908390613190565b60405180910390a1610160015195945050505050565b611955611c7e565b80600d61196282826132b5565b9050507f0447ae479bc793c12cd12089f932a4c0b4ac50f1da17f1379c3d420af34407b881604051610f2591906133c6565b61199c611c7e565b600e805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527fc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d1590602001610f25565b60606003805480602002602001604051908101604052809291908181526020018280548015611a4f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611a31575b5050505050905090565b611a61611c7e565b60028054821515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9091161790556040517fccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df03290610f2590831515815260200190565b611ad1611c7e565b6001600160a01b0382161580611aee57506001600160a01b038116155b15611b25576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166000908152600f6020526040902054600160a01b900460ff1615611b7f576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182526001600160a01b03838116808352600160208085018281528885166000818152600f84528881209751885493511515600160a01b027fffffffffffffffffffffff000000000000000000000000000000000000000000909416971696909617919091179095556010805492830181559093527f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae67201805473ffffffffffffffffffffffffffffffffffffffff1916841790558351928352908201527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c91016106ee565b611c72611c7e565b611c7b81612488565b50565b60005461010090046001600160a01b0316331461095a5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161080b565b60408051808201909152600080825260208201526000829003611d1757506040805180820190915262030d40815260006020820152611e3f565b7f97a657c900000000000000000000000000000000000000000000000000000000611d46600460008587613430565b611d4f9161345a565b7fffffffff000000000000000000000000000000000000000000000000000000001614611e0c577f97a657c900000000000000000000000000000000000000000000000000000000611da5600460008587613430565b611dae9161345a565b6040517f55a0e02c0000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000092831660048201529116602482015260440161080b565b6040805180820190915280611e25602460048688613430565b810190611e3291906134a2565b8152600060209091015290505b92915050565b60005460ff16611e975760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161080b565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60005460ff1615611f345760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161080b565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258611ec43390565b6000818310611f785781611f7a565b825b9392505050565b6001810154600282015442911480611f9c5750808260030154145b15611fa5575050565b816001015482600201541115611fe7576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482611ff99190613061565b60018401548454919250612020916120119084612e1b565b85600201546113e19190612e03565b600284015550600390910155565b600e546001600160a01b0316612070576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381166120b0576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5468010000000000000000900467ffffffffffffffff1684111561212457600d546040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090910467ffffffffffffffff1660048201526024810185905260440161080b565b600d547801000000000000000000000000000000000000000000000000900467ffffffffffffffff16831115612186576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151600d54700100000000000000000000000000000000900467ffffffffffffffff168111156121e2576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254600160a01b900460ff16801561221457506001600160a01b03821660009081526004602052604090205460ff16155b15612256576040517fd0d259760000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161080b565b61225f83612544565b60005b8181101561237657600084828151811061227e5761227e612ebf565b60200260200101519050600081600001519050600061229c82610e23565b90506001600160a01b0381166122e9576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161080b565b60208301516040517f503c285800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0382169063503c285890602401600060405180830381600087803b15801561234a57600080fd5b505af115801561235e573d6000803e3d6000fd5b505050505050508061236f90612eee565b9050612262565b505050505050565b60008060001b828460200151856080015186606001518760e00151886101000151805190602001208961012001516040516020016123bc91906134bb565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d6040015160405160200161246a9c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b336001600160a01b038216036124e05760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161080b565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000805b82518110156126435760006006600085848151811061256957612569612ebf565b6020026020010151600001516001600160a01b03166001600160a01b03168152602001908152602001600020549050806000036125fc578382815181106125b2576125b2612ebf565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260240161080b565b83828151811061260e5761260e612ebf565b602002602001015160200151816126259190612e1b565b61262f9084612e03565b9250508061263c90612eee565b9050612548565b508015612748576126546008611f81565b60095481111561269e576009546040517f688ccf7700000000000000000000000000000000000000000000000000000000815260048101919091526024810182905260440161080b565b600a548111156126fe57600854600a54600091906126bc9084613061565b6126c69190612e58565b9050806040517fe31e0f3200000000000000000000000000000000000000000000000000000000815260040161080b91815260200190565b80600860020160008282546127139190613061565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378906020016106ee565b5050565b8280548282559060005260206000209081019282156127ae579160200282015b828111156127ae578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390911617825560209092019160019091019061276c565b506127ba92915061281e565b5090565b8280548282559060005260206000209081019282156127ae579160200282015b828111156127ae57815473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038435161782556020909201916001909101906127de565b5b808211156127ba576000815560010161281f565b6001600160a01b0381168114611c7b57600080fd5b60006020828403121561285a57600080fd5b8135611f7a81612833565b6000815180845260005b8181101561288b5760208185018101518683018201520161286f565b8181111561289d576000602083870101525b50601f01601f19169290920160200192915050565b602081526000611f7a6020830184612865565b600060a082840312156128d757600080fd5b50919050565b6000602082840312156128ef57600080fd5b813567ffffffffffffffff81111561290657600080fd5b61090b848285016128c5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561296457612964612912565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561299357612993612912565b604052919050565b600067ffffffffffffffff8211156129b5576129b5612912565b5060051b60200190565b600082601f8301126129d057600080fd5b813560206129e56129e08361299b565b61296a565b82815260059290921b84018101918181019086841115612a0457600080fd5b8286015b84811015612a28578035612a1b81612833565b8352918301918301612a08565b509695505050505050565b60008060408385031215612a4657600080fd5b823567ffffffffffffffff80821115612a5e57600080fd5b612a6a868387016129bf565b9350602091508185013581811115612a8157600080fd5b85019050601f81018613612a9457600080fd5b8035612aa26129e08261299b565b81815260059190911b82018301908381019088831115612ac157600080fd5b928401925b82841015612adf57833582529284019290840190612ac6565b80955050505050509250929050565b600060208284031215612b0057600080fd5b813567ffffffffffffffff811115612b1757600080fd5b61090b848285016129bf565b6020808252825182820181905260009190848201906040850190845b81811015612b5b57835183529284019291840191600101612b3f565b50909695505050505050565b60008060208385031215612b7a57600080fd5b823567ffffffffffffffff80821115612b9257600080fd5b818501915085601f830112612ba657600080fd5b813581811115612bb557600080fd5b8660208260051b8501011115612bca57600080fd5b60209290920196919550909350505050565b600060c082840312156128d757600080fd5b60008060408385031215612c0157600080fd5b8235612c0c81612833565b91506020830135612c1c81612833565b809150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015612b5b5783516001600160a01b031683529284019291840191600101612c43565b600060408284031215612c7a57600080fd5b612c82612941565b82358152602083013560208201528091505092915050565b600080600060608486031215612caf57600080fd5b833567ffffffffffffffff811115612cc657600080fd5b612cd2868287016128c5565b935050602084013591506040840135612cea81612833565b809150509250925092565b6000608082840312156128d757600080fd5b8015158114611c7b57600080fd5b600060208284031215612d2757600080fd5b8135611f7a81612d07565b600060208284031215612d4457600080fd5b8151611f7a81612d07565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112612d8457600080fd5b83018035915067ffffffffffffffff821115612d9f57600080fd5b602001915036819003821315612db457600080fd5b9250929050565b600060208284031215612dcd57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115612e1657612e16612dd4565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615612e5357612e53612dd4565b500290565b600082612e8e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600067ffffffffffffffff808316818516808303821115612eb657612eb6612dd4565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612f1f57612f1f612dd4565b5060010190565b60208082528181018390526000908460408401835b86811015612a28578235612f4e81612833565b6001600160a01b031682529183019190830190600101612f3b565b8135612f7481612833565b815473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038216178255506020820135600182015560408201356002820155606082013560038201556080820135612fc781612833565b60048201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383161790555060a082013560058201555050565b60c08101823561300f81612833565b6001600160a01b0380821684526020850135602085015260408501356040850152606085013560608501526080850135915061304a82612833565b16608083015260a092830135929091019190915290565b60008282101561307357613073612dd4565b500390565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126130ad57600080fd5b83018035915067ffffffffffffffff8211156130c857600080fd5b6020019150600681901b3603821315612db457600080fd5b6000604082840312156130f257600080fd5b6130fa612941565b823561310581612833565b81526020928301359281019290925250919050565b600067ffffffffffffffff80831681810361313757613137612dd4565b6001019392505050565b600081518084526020808501945080840160005b8381101561318557815180516001600160a01b031688528301518388015260409096019590820190600101613155565b509495945050505050565b6020815281516020820152600060208301516131b8604084018267ffffffffffffffff169052565b506040830151606083015260608301516131dd60808401826001600160a01b03169052565b50608083015167ffffffffffffffff811660a08401525060a083015160c083015260c083015161321160e084018215159052565b5060e083015161010061322e818501836001600160a01b03169052565b80850151915050610180610120818186015261324e6101a0860184612865565b9250808601519050610140601f19868503018187015261326e8483613141565b93508087015191505061016061328e818701836001600160a01b03169052565b959095015193019290925250919050565b67ffffffffffffffff81168114611c7b57600080fd5b81356132c08161329f565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000821617835560208401356133048161329f565b6fffffffffffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffff0000000000000000000000000000000084161717845560408501356133538161329f565b77ffffffffffffffff000000000000000000000000000000008160801b1690507fffffffffffffffff0000000000000000000000000000000000000000000000008185828616178417178655606087013593506133af8461329f565b808460c01b16858417831717865550505050505050565b6080810182356133d58161329f565b67ffffffffffffffff90811683526020840135906133f28261329f565b90811660208401526040840135906134098261329f565b90811660408401526060840135906134208261329f565b8082166060850152505092915050565b6000808585111561344057600080fd5b8386111561344d57600080fd5b5050820193919092039150565b7fffffffff00000000000000000000000000000000000000000000000000000000813581811691600485101561349a5780818660040360031b1b83161692505b505092915050565b6000602082840312156134b457600080fd5b5035919050565b602081526000611f7a602083018461314156fea164736f6c634300080f000a",
}

var EVM2EVMGEOnRampABI = EVM2EVMGEOnRampMetaData.ABI

var EVM2EVMGEOnRampBin = EVM2EVMGEOnRampMetaData.Bin

func DeployEVM2EVMGEOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, chainId *big.Int, destinationChainId *big.Int, tokens []common.Address, pools []common.Address, allowlist []common.Address, afn common.Address, config BaseOnRampInterfaceOnRampConfig, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address, router common.Address, feeConfig EVM2EVMGEOnRampInterfaceDynamicFeeConfig) (common.Address, *types.Transaction, *EVM2EVMGEOnRamp, error) {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(AggregateRateLimiterInterfaceTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(AggregateRateLimiterInterfaceTokenBucket)).(*AggregateRateLimiterInterfaceTokenBucket)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMGEOnRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetConfig(opts *bind.CallOpts) (BaseOnRampInterfaceOnRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOnRampInterfaceOnRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOnRampInterfaceOnRampConfig)).(*BaseOnRampInterfaceOnRampConfig)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetConfig() (BaseOnRampInterfaceOnRampConfig, error) {
	return _EVM2EVMGEOnRamp.Contract.GetConfig(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetConfig() (BaseOnRampInterfaceOnRampConfig, error) {
	return _EVM2EVMGEOnRamp.Contract.GetConfig(&_EVM2EVMGEOnRamp.CallOpts)
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetPoolTokens(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetPoolTokens(&_EVM2EVMGEOnRamp.CallOpts)
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) IChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "i_chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) IChainId() (*big.Int, error) {
	return _EVM2EVMGEOnRamp.Contract.IChainId(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) IChainId() (*big.Int, error) {
	return _EVM2EVMGEOnRamp.Contract.IChainId(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) IDestinationChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "i_destinationChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) IDestinationChainId() (*big.Int, error) {
	return _EVM2EVMGEOnRamp.Contract.IDestinationChainId(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) IDestinationChainId() (*big.Int, error) {
	return _EVM2EVMGEOnRamp.Contract.IDestinationChainId(&_EVM2EVMGEOnRamp.CallOpts)
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetConfig(opts *bind.TransactOpts, config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetConfig(config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetConfig(&_EVM2EVMGEOnRamp.TransactOpts, config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetConfig(config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetConfig(&_EVM2EVMGEOnRamp.TransactOpts, config)
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetFeeConfig(opts *bind.TransactOpts, feeConfig EVM2EVMGEOnRampInterfaceDynamicFeeConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setFeeConfig", feeConfig)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetFeeConfig(feeConfig EVM2EVMGEOnRampInterfaceDynamicFeeConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetFeeConfig(&_EVM2EVMGEOnRamp.TransactOpts, feeConfig)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetFeeConfig(feeConfig EVM2EVMGEOnRampInterfaceDynamicFeeConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetFeeConfig(&_EVM2EVMGEOnRamp.TransactOpts, feeConfig)
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetRateLimiterConfig(&_EVM2EVMGEOnRamp.TransactOpts, config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
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
	FeeConfig EVM2EVMGEOnRampInterfaceDynamicFeeConfig
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
	Config BaseOnRampInterfaceOnRampConfig
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
	return common.HexToHash("0xd94df26131d6ea08cb76088e9db81d712d59856f9a5872cac5459a9ad822e655")
}

func (EVM2EVMGEOnRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMGEOnRampFeeAdminSet) Topic() common.Hash {
	return common.HexToHash("0x34efeb1f04731080ec2147b3b9c7e38f9b884e035020914e40269450f53b4815")
}

func (EVM2EVMGEOnRampFeeConfigSet) Topic() common.Hash {
	return common.HexToHash("0x690597a8e557c5794b70ef51c386875de46b2ac98f234cc214b803785c1e7d84")
}

func (EVM2EVMGEOnRampOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x0447ae479bc793c12cd12089f932a4c0b4ac50f1da17f1379c3d420af34407b8")
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
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetAllowlist(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	GetConfig(opts *bind.CallOpts) (BaseOnRampInterfaceOnRampConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetFee(opts *bind.CallOpts, message GEConsumerEVM2AnyGEMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	IChainId(opts *bind.CallOpts) (*big.Int, error)

	IDestinationChainId(opts *bind.CallOpts) (*big.Int, error)

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

	SetConfig(opts *bind.TransactOpts, config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error)

	SetFeeAdmin(opts *bind.TransactOpts, feeAdmin common.Address) (*types.Transaction, error)

	SetFeeConfig(opts *bind.TransactOpts, feeConfig EVM2EVMGEOnRampInterfaceDynamicFeeConfig) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error)

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
