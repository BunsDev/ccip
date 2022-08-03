// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_subscription_onramp

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
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"
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

type BaseOnRampInterfaceOnRampConfig struct {
	RelayingFeeJuels uint64
	MaxDataSize      uint64
	MaxTokensLength  uint64
}

type CCIPEVM2AnySubscriptionMessage struct {
	Receiver common.Address
	Data     []byte
	Tokens   []common.Address
	Amounts  []*big.Int
	GasLimit *big.Int
}

type CCIPEVM2EVMSubscriptionEvent struct {
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         common.Address
	Receiver       common.Address
	Nonce          uint64
	Data           []byte
	Tokens         []common.Address
	Amounts        []*big.Int
	GasLimit       *big.Int
}

var EVM2EVMSubscriptionOnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"contractAggregatorV2V3Interface[]\",\"name\":\"feeds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTimeWithoutAFNSignal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"internalType\":\"contractAny2EVMSubscriptionOnRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTokenAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeedDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriceFeedConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleAFNHeartbeat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenConfigMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenFeedMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"AFNMaxHeartbeatTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"AllowListSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCCIP.EVM2EVMSubscriptionEvent\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OnRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DESTINATION_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"addFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVM2AnySubscriptionMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeed\",\"outputs\":[{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNow\",\"type\":\"uint256\"}],\"name\":\"isHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV2V3Interface\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"removeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"setAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAllowlistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"relayingFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOnRampInterface.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"setMaxSecondsWithoutAFNHeartbeat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAccumulatedFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162004118380380620041188339810160408190526200003491620007fc565b6000805460ff191681558a908a908a908a908a908a908a908a908a908a9085908990889082908b90899089903390819081620000b75760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000f157620000f181620004b5565b5050506001600160a01b038216158062000109575080155b156200012857604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03939093169290921790915560035580518251146200016e5760405162d8548360e71b815260040160405180910390fd5b81516200018390600590602085019062000566565b5060005b825181101562000265576000828281518110620001a857620001a86200090e565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060046000868581518110620001f257620001f26200090e565b6020908102919091018101516001600160a01b03908116835282820193909352604091820160009081208551958301516001600160601b0316600160a01b0295851695909517909455939091168252600690925220805460ff191660011790556200025d8162000924565b905062000187565b50505080518251146200028b5760405163ee9d106b60e01b815260040160405180910390fd5b8151620002a090600890602085019062000566565b5060005b82518110156200036b576000828281518110620002c557620002c56200090e565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b0316815250600760008685815181106200030f576200030f6200090e565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b029290911691909117905550620003638162000924565b9050620002a4565b505081511590506200039b576009805460ff1916600117905580516200039990600a90602084019062000566565b505b60005b815181101562000408576001600b6000848481518110620003c357620003c36200090e565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055620004008162000924565b90506200039e565b505060809990995260a09790975250508451600d805460208801516040909801516001600160401b03908116600160801b02600160801b600160c01b031999821668010000000000000000026001600160801b031990931691909416171796909616179094555050600e80546001600160a01b039094166001600160a01b0319909416939093179092555050600c80546001600160401b0319169055506200094c98505050505050505050565b336001600160a01b038216036200050f5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ae565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620005be579160200282015b82811115620005be57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000587565b50620005cc929150620005d0565b5090565b5b80821115620005cc5760008155600101620005d1565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620006285762000628620005e7565b604052919050565b60006001600160401b038211156200064c576200064c620005e7565b5060051b60200190565b6001600160a01b03811681146200066c57600080fd5b50565b600082601f8301126200068157600080fd5b815160206200069a620006948362000630565b620005fd565b82815260059290921b84018101918181019086841115620006ba57600080fd5b8286015b84811015620006e2578051620006d48162000656565b8352918301918301620006be565b509695505050505050565b600082601f830112620006ff57600080fd5b8151602062000712620006948362000630565b82815260059290921b840181019181810190868411156200073257600080fd5b8286015b84811015620006e25780516200074c8162000656565b835291830191830162000736565b8051620007678162000656565b919050565b80516001600160401b03811681146200076757600080fd5b6000606082840312156200079757600080fd5b604051606081016001600160401b0381118282101715620007bc57620007bc620005e7565b604052905080620007cd836200076c565b8152620007dd602084016200076c565b6020820152620007f0604084016200076c565b60408201525092915050565b6000806000806000806000806000806101808b8d0312156200081d57600080fd5b8a5160208c015160408d0151919b5099506001600160401b03808211156200084457600080fd5b620008528e838f016200066f565b995060608d01519150808211156200086957600080fd5b620008778e838f016200066f565b985060808d01519150808211156200088e57600080fd5b6200089c8e838f01620006ed565b975060a08d0151915080821115620008b357600080fd5b50620008c28d828e01620006ed565b955050620008d360c08c016200075a565b935060e08b01519250620008ec8c6101008d0162000784565b9150620008fd6101608c016200075a565b90509295989b9194979a5092959850565b634e487b7160e01b600052603260045260246000fd5b6000600182016200094557634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05161379f6200097960003960006103860152600081816104d70152611b60015261379f6000f3fe608060405234801561001057600080fd5b50600436106102775760003560e01c806379ba509711610160578063b4069b31116100d8578063c5eff3d01161008c578063eb511dd411610071578063eb511dd414610699578063f2fde38b146106ac578063f78faa32146106bf57600080fd5b8063c5eff3d014610671578063d7644ba21461068657600080fd5b8063bbe4f6db116100bd578063bbe4f6db14610583578063c0d78655146105bc578063c3f909d4146105cf57600080fd5b8063b4069b311461055d578063b6608c3b1461057057600080fd5b806389c065681161012f578063ae990dce11610114578063ae990dce14610524578063b034909c14610537578063b0f479a11461053f57600080fd5b806389c06568146104f95780638da5cb5b1461050157600080fd5b806379ba5097146104ba57806381be8fa4146104c25780638456cb59146104ca57806385e1f4d0146104d257600080fd5b8063552b818b116101f35780635b16ebb7116101c2578063671dc337116101a7578063671dc3371461047f578063681fba1614610492578063744b92e2146104a757600080fd5b80635b16ebb71461043b5780635c975abb1461047457600080fd5b8063552b818b146103df578063567c814b146103f25780635853c6271461041557806359e96b5b1461042857600080fd5b80632222dd421161024a5780632ea023691161022f5780632ea02369146103815780633f4ba83a146103b65780634120fccd146103be57600080fd5b80632222dd42146103505780632b898c251461036e57600080fd5b806304c2a34a1461027c578063108ee5fc146102b957806316b8e731146102ce578063181f5a7714610307575b600080fd5b61028f61028a366004612cfe565b6106ca565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6102cc6102c7366004612cfe565b6106fb565b005b61028f6102dc366004612cfe565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b6103436040518060400160405280601f81526020017f45564d3245564d537562736372697074696f6e4f6e52616d7020312e302e300081525081565b6040516102b09190612d95565b60025473ffffffffffffffffffffffffffffffffffffffff1661028f565b6102cc61037c366004612da8565b6107d7565b6103a87f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016102b0565b6102cc610ba7565b6103c6610bb9565b60405167ffffffffffffffff90911681526020016102b0565b6102cc6103ed366004612de1565b610bd9565b610405610400366004612e56565b610dc6565b60405190151581526020016102b0565b6102cc610423366004612da8565b610f0c565b6102cc610436366004612e7f565b61111b565b610405610449366004612cfe565b73ffffffffffffffffffffffffffffffffffffffff1660009081526006602052604090205460ff1690565b60005460ff16610405565b6102cc61048d366004612ec0565b611199565b61049a6111eb565b6040516102b09190612f29565b6102cc6104b5366004612da8565b6112ca565b6102cc6116bf565b61049a6117e6565b6102cc611855565b6103a87f000000000000000000000000000000000000000000000000000000000000000081565b61049a611865565b600054610100900473ffffffffffffffffffffffffffffffffffffffff1661028f565b6103c6610532366004613164565b6118d2565b6003546103a8565b600e5473ffffffffffffffffffffffffffffffffffffffff1661028f565b61028f61056b366004612cfe565b611d15565b6102cc61057e366004612e56565b611e1a565b61028f610591366004612cfe565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b6102cc6105ca366004612cfe565b611e9a565b61063e60408051606081018252600080825260208201819052918101919091525060408051606081018252600d5467ffffffffffffffff808216835268010000000000000000820481166020840152700100000000000000000000000000000000909104169181019190915290565b60408051825167ffffffffffffffff908116825260208085015182169083015292820151909216908201526060016102b0565b610679611f15565b6040516102b09190613247565b6102cc6106943660046132af565b611f82565b6102cc6106a7366004612da8565b611fe9565b6102cc6106ba366004612cfe565b612229565b60095460ff16610405565b73ffffffffffffffffffffffffffffffffffffffff8082166000908152600460205260408120549091165b92915050565b61070361223d565b73ffffffffffffffffffffffffffffffffffffffff8116610750576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b6107df61223d565b600854600081900361081d576040517f2e70248b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906108b8576040517f3917193900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610921576040517f9403a50500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060086109306001856132fb565b8154811061094057610940613312565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600882602001516bffffffffffffffffffffffff168154811061099257610992613312565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1660086109c16001866132fb565b815481106109d1576109d1613312565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600883602001516bffffffffffffffffffffffff1681548110610a3f57610a3f613312565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526007909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556008805480610ae157610ae1613341565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600783526040808520949094558351908152908716918101919091527fa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d91015b60405180910390a15050505050565b610baf61223d565b610bb76122c3565b565b600c54600090610bd49067ffffffffffffffff166001613370565b905090565b610be161223d565b6000600a805480602002602001604051908101604052809291908181526020018280548015610c4657602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610c1b575b5050505050905060005b8151811015610cde576000600b6000848481518110610c7157610c71613312565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055610cd78161339c565b9050610c50565b50610ceb600a8484612c48565b5060005b82811015610d87576001600b6000868685818110610d0f57610d0f613312565b9050602002016020810190610d249190612cfe565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055610d808161339c565b9050610cef565b507ff8adc5fee247b62a85f63294cb46e4be61da815e4756bc57672a83b24faf0dda8383604051610db99291906133d4565b60405180910390a1505050565b600254604080517fcf72b39b000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163cf72b39b9160048083019260209291908290030181865afa158015610e36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5a9190613424565b1580156106f55750600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663343157b46040518163ffffffff1660e01b8152600401606060405180830381865afa158015610ed2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef69190613441565b60200151610f0490846132fb565b111592915050565b610f1461223d565b73ffffffffffffffffffffffffffffffffffffffff82161580610f4b575073ffffffffffffffffffffffffffffffffffffffff8116155b15610f82576040517fee9d106b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260076020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152901561101e576040517f965ffb7b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600880546bffffffffffffffffffffffff908116602080870191825288861660008181526007835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee390920180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821790558351908152908101919091527f037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb6013059101610db9565b61112361223d565b61114473ffffffffffffffffffffffffffffffffffffffff841683836123a4565b6040805173ffffffffffffffffffffffffffffffffffffffff8086168252841660208201529081018290527f5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa890606001610db9565b6111a161223d565b80600d6111ae82826134b3565b9050507fcc6ce9e57c1de2adf58a81e94b96b43d77ea6973e3f08e6ea4fe83d62ae60e9e816040516111e091906135a1565b60405180910390a150565b60055460609067ffffffffffffffff81111561120957611209612f3c565b604051908082528060200260200182016040528015611232578160200160208202803683370190505b50905060005b6005548110156112c65761127f6005828154811061125857611258613312565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16611d15565b82828151811061129157611291613312565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101526112bf8161339c565b9050611238565b5090565b6112d261223d565b6005546000819003611310576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906113ab576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614611414576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060056114236001856132fb565b8154811061143357611433613312565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600582602001516bffffffffffffffffffffffff168154811061148557611485613312565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1660056114b46001866132fb565b815481106114c4576114c4613312565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600583602001516bffffffffffffffffffffffff168154811061153257611532613312565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558481015184841683526004909152604090912080546bffffffffffffffffffffffff90921674010000000000000000000000000000000000000000029190921617905560058054806115d4576115d4613341565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff878116808452600483526040808520859055918816808552600684529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558151908152918201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610b98565b60015473ffffffffffffffffffffffffffffffffffffffff163314611745576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405173ffffffffffffffffffffffffffffffffffffffff919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060600880548060200260200160405190810160405280929190818152602001828054801561184b57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611820575b5050505050905090565b61185d61223d565b610bb7612436565b6060600580548060200260200160405190810160405280929190818152602001828054801561184b5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611820575050505050905090565b6000805460ff1615611940576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161173c565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cf72b39b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119d19190613424565b15611a07576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254604080517f343157b4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163343157b49160048083019260609291908290030181865afa158015611a77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a9b9190613441565b9050600354816020015142611ab091906132fb565b1115611ae8576040517fa8c8866900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e5473ffffffffffffffffffffffffffffffffffffffff163314611b39576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611b5284602001515185604001518660600151866124f6565b6040805161012081019091527f00000000000000000000000000000000000000000000000000000000000000008152600c80546000929160208301918490611ba39067ffffffffffffffff166135f4565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001866000015173ffffffffffffffffffffffffffffffffffffffff168152602001600f6000886000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081819054906101000a900467ffffffffffffffff16611c76906135f4565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff168152602001866020015181526020018660400151815260200186606001518152602001866080015181525090507f73dfb9df8214728e699dbaaf6ba97aa125afaaba83a5d0de7903062e7c5b313981604051611d01919061364b565b60405180910390a160200151949350505050565b73ffffffffffffffffffffffffffffffffffffffff80821660009081526004602052604081205490911680611d76576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8084166000908152600460208181526040928390205483517f21df0da700000000000000000000000000000000000000000000000000000000815293519416936321df0da79380840193908290030181865afa158015611def573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e139190613759565b9392505050565b611e2261223d565b80600003611e5c576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380549082905560408051828152602081018490527f72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c25191016107cb565b611ea261223d565b600e80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15906020016111e0565b6060600a80548060200260200160405190810160405280929190818152602001828054801561184b5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611820575050505050905090565b611f8a61223d565b600980547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168215159081179091556040519081527fccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df032906020016111e0565b611ff161223d565b73ffffffffffffffffffffffffffffffffffffffff82161580612028575073ffffffffffffffffffffffffffffffffffffffff8116155b1561205f576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116600090815260046020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156120fb576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff828116808352600580546bffffffffffffffffffffffff908116602080870191825288861660008181526004835260408082208a51955190961674010000000000000000000000000000000000000000029490981693909317909355835460018082019095557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055848252600681529085902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169093179092558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c9101610db9565b61223161223d565b61223a81612857565b50565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610bb7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161173c565b60005460ff1661232f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161173c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052612431908490612952565b505050565b60005460ff16156124a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161173c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861237a3390565b600e5473ffffffffffffffffffffffffffffffffffffffff16612545576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116612592576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5468010000000000000000900467ffffffffffffffff1684111561260657600d546040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090910467ffffffffffffffff1660048201526024810185905260440161173c565b8251600d54700100000000000000000000000000000000900467ffffffffffffffff16811180612637575082518114155b1561266e576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60095460ff1680156126a6575073ffffffffffffffffffffffffffffffffffffffff82166000908152600b602052604090205460ff16155b156126f5576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161173c565b60005b8181101561284f57600085828151811061271457612714613312565b60200260200101519050600061274f8273ffffffffffffffffffffffffffffffffffffffff9081166000908152600460205260409020541690565b905073ffffffffffffffffffffffffffffffffffffffff81166127b6576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161173c565b8073ffffffffffffffffffffffffffffffffffffffff1663503c28588785815181106127e4576127e4613312565b60200260200101516040518263ffffffff1660e01b815260040161280a91815260200190565b600060405180830381600087803b15801561282457600080fd5b505af1158015612838573d6000803e3d6000fd5b505050505050806128489061339c565b90506126f8565b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036128d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161173c565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006129b4826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16612a5e9092919063ffffffff16565b80519091501561243157808060200190518101906129d29190613424565b612431576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161173c565b6060612a6d8484600085612a75565b949350505050565b606082471015612b07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161173c565b843b612b6f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161173c565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051612b989190613776565b60006040518083038185875af1925050503d8060008114612bd5576040519150601f19603f3d011682016040523d82523d6000602084013e612bda565b606091505b5091509150612bea828286612bf5565b979650505050505050565b60608315612c04575081611e13565b825115612c145782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161173c9190612d95565b828054828255906000526020600020908101928215612cc0579160200282015b82811115612cc05781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff843516178255602090920191600190910190612c68565b506112c69291505b808211156112c65760008155600101612cc8565b73ffffffffffffffffffffffffffffffffffffffff8116811461223a57600080fd5b600060208284031215612d1057600080fd5b8135611e1381612cdc565b60005b83811015612d36578181015183820152602001612d1e565b83811115612d45576000848401525b50505050565b60008151808452612d63816020860160208601612d1b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611e136020830184612d4b565b60008060408385031215612dbb57600080fd5b8235612dc681612cdc565b91506020830135612dd681612cdc565b809150509250929050565b60008060208385031215612df457600080fd5b823567ffffffffffffffff80821115612e0c57600080fd5b818501915085601f830112612e2057600080fd5b813581811115612e2f57600080fd5b8660208260051b8501011115612e4457600080fd5b60209290920196919550909350505050565b600060208284031215612e6857600080fd5b5035919050565b8035612e7a81612cdc565b919050565b600080600060608486031215612e9457600080fd5b8335612e9f81612cdc565b92506020840135612eaf81612cdc565b929592945050506040919091013590565b600060608284031215612ed257600080fd5b50919050565b600081518084526020808501945080840160005b83811015612f1e57815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101612eec565b509495945050505050565b602081526000611e136020830184612ed8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715612f8e57612f8e612f3c565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612fdb57612fdb612f3c565b604052919050565b600082601f830112612ff457600080fd5b813567ffffffffffffffff81111561300e5761300e612f3c565b61303f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601612f94565b81815284602083860101111561305457600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff82111561308b5761308b612f3c565b5060051b60200190565b600082601f8301126130a657600080fd5b813560206130bb6130b683613071565b612f94565b82815260059290921b840181019181810190868411156130da57600080fd5b8286015b848110156130fe5780356130f181612cdc565b83529183019183016130de565b509695505050505050565b600082601f83011261311a57600080fd5b8135602061312a6130b683613071565b82815260059290921b8401810191818101908684111561314957600080fd5b8286015b848110156130fe578035835291830191830161314d565b6000806040838503121561317757600080fd5b823567ffffffffffffffff8082111561318f57600080fd5b9084019060a082870312156131a357600080fd5b6131ab612f6b565b6131b483612e6f565b81526020830135828111156131c857600080fd5b6131d488828601612fe3565b6020830152506040830135828111156131ec57600080fd5b6131f888828601613095565b60408301525060608301358281111561321057600080fd5b61321c88828601613109565b6060830152506080830135608082015280945050505061323e60208401612e6f565b90509250929050565b6020808252825182820181905260009190848201906040850190845b8181101561329557835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101613263565b50909695505050505050565b801515811461223a57600080fd5b6000602082840312156132c157600080fd5b8135611e13816132a1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561330d5761330d6132cc565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600067ffffffffffffffff808316818516808303821115613393576133936132cc565b01949350505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036133cd576133cd6132cc565b5060010190565b60208082528181018390526000908460408401835b868110156130fe5782356133fc81612cdc565b73ffffffffffffffffffffffffffffffffffffffff16825291830191908301906001016133e9565b60006020828403121561343657600080fd5b8151611e13816132a1565b60006060828403121561345357600080fd5b6040516060810181811067ffffffffffffffff8211171561347657613476612f3c565b80604052508251815260208301516020820152604083015160408201528091505092915050565b67ffffffffffffffff8116811461223a57600080fd5b81356134be8161349d565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000821617835560208401356135028161349d565b6fffffffffffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffff0000000000000000000000000000000084161717845560408501356135518161349d565b77ffffffffffffffff000000000000000000000000000000008160801b16847fffffffffffffffff0000000000000000000000000000000000000000000000008516178317178555505050505050565b6060810182356135b08161349d565b67ffffffffffffffff90811683526020840135906135cd8261349d565b90811660208401526040840135906135e48261349d565b8082166040850152505092915050565b600067ffffffffffffffff808316818103613611576136116132cc565b6001019392505050565b600081518084526020808501945080840160005b83811015612f1e5781518752958201959082019060010161362f565b602081528151602082015260006020830151613673604084018267ffffffffffffffff169052565b50604083015173ffffffffffffffffffffffffffffffffffffffff8116606084015250606083015173ffffffffffffffffffffffffffffffffffffffff8116608084015250608083015167ffffffffffffffff811660a08401525060a08301516101208060c08501526136ea610140850183612d4b565b915060c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808685030160e08701526137268483612ed8565b935060e08701519150610100818786030181880152613745858461361b565b970151959092019490945250929392505050565b60006020828403121561376b57600080fd5b8151611e1381612cdc565b60008251613788818460208701612d1b565b919091019291505056fea164736f6c634300080f000a",
}

var EVM2EVMSubscriptionOnRampABI = EVM2EVMSubscriptionOnRampMetaData.ABI

var EVM2EVMSubscriptionOnRampBin = EVM2EVMSubscriptionOnRampMetaData.Bin

func DeployEVM2EVMSubscriptionOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, chainId *big.Int, destinationChainId *big.Int, tokens []common.Address, pools []common.Address, feeds []common.Address, allowlist []common.Address, afn common.Address, maxTimeWithoutAFNSignal *big.Int, config BaseOnRampInterfaceOnRampConfig, router common.Address) (common.Address, *types.Transaction, *EVM2EVMSubscriptionOnRamp, error) {
	parsed, err := EVM2EVMSubscriptionOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMSubscriptionOnRampBin), backend, chainId, destinationChainId, tokens, pools, feeds, allowlist, afn, maxTimeWithoutAFNSignal, config, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMSubscriptionOnRamp{EVM2EVMSubscriptionOnRampCaller: EVM2EVMSubscriptionOnRampCaller{contract: contract}, EVM2EVMSubscriptionOnRampTransactor: EVM2EVMSubscriptionOnRampTransactor{contract: contract}, EVM2EVMSubscriptionOnRampFilterer: EVM2EVMSubscriptionOnRampFilterer{contract: contract}}, nil
}

type EVM2EVMSubscriptionOnRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMSubscriptionOnRampCaller
	EVM2EVMSubscriptionOnRampTransactor
	EVM2EVMSubscriptionOnRampFilterer
}

type EVM2EVMSubscriptionOnRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMSubscriptionOnRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMSubscriptionOnRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMSubscriptionOnRampSession struct {
	Contract     *EVM2EVMSubscriptionOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMSubscriptionOnRampCallerSession struct {
	Contract *EVM2EVMSubscriptionOnRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMSubscriptionOnRampTransactorSession struct {
	Contract     *EVM2EVMSubscriptionOnRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMSubscriptionOnRampRaw struct {
	Contract *EVM2EVMSubscriptionOnRamp
}

type EVM2EVMSubscriptionOnRampCallerRaw struct {
	Contract *EVM2EVMSubscriptionOnRampCaller
}

type EVM2EVMSubscriptionOnRampTransactorRaw struct {
	Contract *EVM2EVMSubscriptionOnRampTransactor
}

func NewEVM2EVMSubscriptionOnRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMSubscriptionOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMSubscriptionOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMSubscriptionOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRamp{address: address, abi: abi, EVM2EVMSubscriptionOnRampCaller: EVM2EVMSubscriptionOnRampCaller{contract: contract}, EVM2EVMSubscriptionOnRampTransactor: EVM2EVMSubscriptionOnRampTransactor{contract: contract}, EVM2EVMSubscriptionOnRampFilterer: EVM2EVMSubscriptionOnRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMSubscriptionOnRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMSubscriptionOnRampCaller, error) {
	contract, err := bindEVM2EVMSubscriptionOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampCaller{contract: contract}, nil
}

func NewEVM2EVMSubscriptionOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMSubscriptionOnRampTransactor, error) {
	contract, err := bindEVM2EVMSubscriptionOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampTransactor{contract: contract}, nil
}

func NewEVM2EVMSubscriptionOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMSubscriptionOnRampFilterer, error) {
	contract, err := bindEVM2EVMSubscriptionOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampFilterer{contract: contract}, nil
}

func bindEVM2EVMSubscriptionOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMSubscriptionOnRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMSubscriptionOnRamp.Contract.EVM2EVMSubscriptionOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.EVM2EVMSubscriptionOnRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.EVM2EVMSubscriptionOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMSubscriptionOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) CHAINID() (*big.Int, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.CHAINID(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) CHAINID() (*big.Int, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.CHAINID(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "DESTINATION_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.DESTINATIONCHAINID(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) DESTINATIONCHAINID() (*big.Int, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.DESTINATIONCHAINID(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetAFN() (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetAFN(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetAFN(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetAllowlist(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getAllowlist")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetAllowlist() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetAllowlist(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetAllowlist() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetAllowlist(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getAllowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetAllowlistEnabled() (bool, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetAllowlistEnabled(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetAllowlistEnabled() (bool, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetAllowlistEnabled(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetConfig(opts *bind.CallOpts) (BaseOnRampInterfaceOnRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOnRampInterfaceOnRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOnRampInterfaceOnRampConfig)).(*BaseOnRampInterfaceOnRampConfig)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetConfig() (BaseOnRampInterfaceOnRampConfig, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetConfig(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetConfig() (BaseOnRampInterfaceOnRampConfig, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetConfig(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetDestinationToken(&_EVM2EVMSubscriptionOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetDestinationToken(&_EVM2EVMSubscriptionOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetDestinationTokens(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetDestinationTokens(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getExpectedNextSequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getFeed", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetFeed(token common.Address) (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetFeed(&_EVM2EVMSubscriptionOnRamp.CallOpts, token)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetFeed(token common.Address) (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetFeed(&_EVM2EVMSubscriptionOnRamp.CallOpts, token)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getFeedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetFeedTokens() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetFeedTokens(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetFeedTokens() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetFeedTokens(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getMaxSecondsWithoutAFNHeartbeat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetMaxSecondsWithoutAFNHeartbeat() (*big.Int, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetPool(&_EVM2EVMSubscriptionOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetPool(&_EVM2EVMSubscriptionOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetPoolTokens(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetPoolTokens(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetRouter() (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetRouter(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetRouter(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) GetTokenPool(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "getTokenPool", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) GetTokenPool(token common.Address) (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetTokenPool(&_EVM2EVMSubscriptionOnRamp.CallOpts, token)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) GetTokenPool(token common.Address) (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.GetTokenPool(&_EVM2EVMSubscriptionOnRamp.CallOpts, token)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "isHealthy", timeNow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.IsHealthy(&_EVM2EVMSubscriptionOnRamp.CallOpts, timeNow)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) IsHealthy(timeNow *big.Int) (bool, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.IsHealthy(&_EVM2EVMSubscriptionOnRamp.CallOpts, timeNow)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) IsPool(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "isPool", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) IsPool(addr common.Address) (bool, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.IsPool(&_EVM2EVMSubscriptionOnRamp.CallOpts, addr)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) IsPool(addr common.Address) (bool, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.IsPool(&_EVM2EVMSubscriptionOnRamp.CallOpts, addr)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) Owner() (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.Owner(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.Owner(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) Paused() (bool, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.Paused(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.Paused(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMSubscriptionOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.TypeAndVersion(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.TypeAndVersion(&_EVM2EVMSubscriptionOnRamp.CallOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.AcceptOwnership(&_EVM2EVMSubscriptionOnRamp.TransactOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.AcceptOwnership(&_EVM2EVMSubscriptionOnRamp.TransactOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "addFeed", token, feed)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.AddFeed(&_EVM2EVMSubscriptionOnRamp.TransactOpts, token, feed)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) AddFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.AddFeed(&_EVM2EVMSubscriptionOnRamp.TransactOpts, token, feed)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.AddPool(&_EVM2EVMSubscriptionOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.AddPool(&_EVM2EVMSubscriptionOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) ForwardFromRouter(opts *bind.TransactOpts, message CCIPEVM2AnySubscriptionMessage, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "forwardFromRouter", message, originalSender)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) ForwardFromRouter(message CCIPEVM2AnySubscriptionMessage, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.ForwardFromRouter(&_EVM2EVMSubscriptionOnRamp.TransactOpts, message, originalSender)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) ForwardFromRouter(message CCIPEVM2AnySubscriptionMessage, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.ForwardFromRouter(&_EVM2EVMSubscriptionOnRamp.TransactOpts, message, originalSender)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.Pause(&_EVM2EVMSubscriptionOnRamp.TransactOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.Pause(&_EVM2EVMSubscriptionOnRamp.TransactOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "removeFeed", token, feed)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.RemoveFeed(&_EVM2EVMSubscriptionOnRamp.TransactOpts, token, feed)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) RemoveFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.RemoveFeed(&_EVM2EVMSubscriptionOnRamp.TransactOpts, token, feed)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.RemovePool(&_EVM2EVMSubscriptionOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.RemovePool(&_EVM2EVMSubscriptionOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.SetAFN(&_EVM2EVMSubscriptionOnRamp.TransactOpts, afn)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.SetAFN(&_EVM2EVMSubscriptionOnRamp.TransactOpts, afn)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "setAllowlist", allowlist)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.SetAllowlist(&_EVM2EVMSubscriptionOnRamp.TransactOpts, allowlist)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.SetAllowlist(&_EVM2EVMSubscriptionOnRamp.TransactOpts, allowlist)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "setAllowlistEnabled", enabled)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.SetAllowlistEnabled(&_EVM2EVMSubscriptionOnRamp.TransactOpts, enabled)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.SetAllowlistEnabled(&_EVM2EVMSubscriptionOnRamp.TransactOpts, enabled)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) SetConfig(opts *bind.TransactOpts, config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) SetConfig(config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.SetConfig(&_EVM2EVMSubscriptionOnRamp.TransactOpts, config)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) SetConfig(config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.SetConfig(&_EVM2EVMSubscriptionOnRamp.TransactOpts, config)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "setMaxSecondsWithoutAFNHeartbeat", newTime)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMSubscriptionOnRamp.TransactOpts, newTime)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) SetMaxSecondsWithoutAFNHeartbeat(newTime *big.Int) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.SetMaxSecondsWithoutAFNHeartbeat(&_EVM2EVMSubscriptionOnRamp.TransactOpts, newTime)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.SetRouter(&_EVM2EVMSubscriptionOnRamp.TransactOpts, router)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.SetRouter(&_EVM2EVMSubscriptionOnRamp.TransactOpts, router)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.TransferOwnership(&_EVM2EVMSubscriptionOnRamp.TransactOpts, to)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.TransferOwnership(&_EVM2EVMSubscriptionOnRamp.TransactOpts, to)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.Unpause(&_EVM2EVMSubscriptionOnRamp.TransactOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.Unpause(&_EVM2EVMSubscriptionOnRamp.TransactOpts)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactor) WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.contract.Transact(opts, "withdrawAccumulatedFees", feeToken, recipient, amount)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.WithdrawAccumulatedFees(&_EVM2EVMSubscriptionOnRamp.TransactOpts, feeToken, recipient, amount)
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampTransactorSession) WithdrawAccumulatedFees(feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EVM2EVMSubscriptionOnRamp.Contract.WithdrawAccumulatedFees(&_EVM2EVMSubscriptionOnRamp.TransactOpts, feeToken, recipient, amount)
}

type EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSetIterator struct {
	Event *EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSet)
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
		it.Event = new(EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSet)
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

func (it *EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSet struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSetIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSetIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "AFNMaxHeartbeatTimeSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "AFNMaxHeartbeatTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSet)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseAFNMaxHeartbeatTimeSet(log types.Log) (*EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSet, error) {
	event := new(EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSet)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "AFNMaxHeartbeatTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampAFNSetIterator struct {
	Event *EVM2EVMSubscriptionOnRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampAFNSet)
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
		it.Event = new(EVM2EVMSubscriptionOnRampAFNSet)
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

func (it *EVM2EVMSubscriptionOnRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampAFNSetIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampAFNSet)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseAFNSet(log types.Log) (*EVM2EVMSubscriptionOnRampAFNSet, error) {
	event := new(EVM2EVMSubscriptionOnRampAFNSet)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampAllowListEnabledSetIterator struct {
	Event *EVM2EVMSubscriptionOnRampAllowListEnabledSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampAllowListEnabledSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampAllowListEnabledSet)
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
		it.Event = new(EVM2EVMSubscriptionOnRampAllowListEnabledSet)
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

func (it *EVM2EVMSubscriptionOnRampAllowListEnabledSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampAllowListEnabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampAllowListEnabledSet struct {
	Enabled bool
	Raw     types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterAllowListEnabledSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampAllowListEnabledSetIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "AllowListEnabledSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampAllowListEnabledSetIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "AllowListEnabledSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchAllowListEnabledSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampAllowListEnabledSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "AllowListEnabledSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampAllowListEnabledSet)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "AllowListEnabledSet", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseAllowListEnabledSet(log types.Log) (*EVM2EVMSubscriptionOnRampAllowListEnabledSet, error) {
	event := new(EVM2EVMSubscriptionOnRampAllowListEnabledSet)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "AllowListEnabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampAllowListSetIterator struct {
	Event *EVM2EVMSubscriptionOnRampAllowListSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampAllowListSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampAllowListSet)
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
		it.Event = new(EVM2EVMSubscriptionOnRampAllowListSet)
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

func (it *EVM2EVMSubscriptionOnRampAllowListSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampAllowListSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampAllowListSet struct {
	Allowlist []common.Address
	Raw       types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterAllowListSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampAllowListSetIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "AllowListSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampAllowListSetIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "AllowListSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchAllowListSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampAllowListSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "AllowListSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampAllowListSet)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "AllowListSet", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseAllowListSet(log types.Log) (*EVM2EVMSubscriptionOnRampAllowListSet, error) {
	event := new(EVM2EVMSubscriptionOnRampAllowListSet)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "AllowListSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampCCIPSendRequestedIterator struct {
	Event *EVM2EVMSubscriptionOnRampCCIPSendRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampCCIPSendRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampCCIPSendRequested)
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
		it.Event = new(EVM2EVMSubscriptionOnRampCCIPSendRequested)
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

func (it *EVM2EVMSubscriptionOnRampCCIPSendRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampCCIPSendRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampCCIPSendRequested struct {
	Message CCIPEVM2EVMSubscriptionEvent
	Raw     types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampCCIPSendRequestedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampCCIPSendRequestedIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "CCIPSendRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampCCIPSendRequested) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampCCIPSendRequested)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseCCIPSendRequested(log types.Log) (*EVM2EVMSubscriptionOnRampCCIPSendRequested, error) {
	event := new(EVM2EVMSubscriptionOnRampCCIPSendRequested)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampFeeChargedIterator struct {
	Event *EVM2EVMSubscriptionOnRampFeeCharged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampFeeChargedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampFeeCharged)
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
		it.Event = new(EVM2EVMSubscriptionOnRampFeeCharged)
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

func (it *EVM2EVMSubscriptionOnRampFeeChargedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampFeeCharged struct {
	From common.Address
	To   common.Address
	Fee  *big.Int
	Raw  types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterFeeCharged(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampFeeChargedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampFeeChargedIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampFeeCharged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "FeeCharged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampFeeCharged)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseFeeCharged(log types.Log) (*EVM2EVMSubscriptionOnRampFeeCharged, error) {
	event := new(EVM2EVMSubscriptionOnRampFeeCharged)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampFeedAddedIterator struct {
	Event *EVM2EVMSubscriptionOnRampFeedAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampFeedAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampFeedAdded)
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
		it.Event = new(EVM2EVMSubscriptionOnRampFeedAdded)
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

func (it *EVM2EVMSubscriptionOnRampFeedAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampFeedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampFeedAdded struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterFeedAdded(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampFeedAddedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampFeedAddedIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "FeedAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampFeedAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampFeedAdded)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "FeedAdded", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseFeedAdded(log types.Log) (*EVM2EVMSubscriptionOnRampFeedAdded, error) {
	event := new(EVM2EVMSubscriptionOnRampFeedAdded)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "FeedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampFeedRemovedIterator struct {
	Event *EVM2EVMSubscriptionOnRampFeedRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampFeedRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampFeedRemoved)
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
		it.Event = new(EVM2EVMSubscriptionOnRampFeedRemoved)
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

func (it *EVM2EVMSubscriptionOnRampFeedRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampFeedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampFeedRemoved struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterFeedRemoved(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampFeedRemovedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampFeedRemovedIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "FeedRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampFeedRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampFeedRemoved)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseFeedRemoved(log types.Log) (*EVM2EVMSubscriptionOnRampFeedRemoved, error) {
	event := new(EVM2EVMSubscriptionOnRampFeedRemoved)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampFeesWithdrawnIterator struct {
	Event *EVM2EVMSubscriptionOnRampFeesWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampFeesWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampFeesWithdrawn)
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
		it.Event = new(EVM2EVMSubscriptionOnRampFeesWithdrawn)
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

func (it *EVM2EVMSubscriptionOnRampFeesWithdrawnIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampFeesWithdrawn struct {
	FeeToken  common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampFeesWithdrawnIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampFeesWithdrawnIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampFeesWithdrawn)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseFeesWithdrawn(log types.Log) (*EVM2EVMSubscriptionOnRampFeesWithdrawn, error) {
	event := new(EVM2EVMSubscriptionOnRampFeesWithdrawn)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampOnRampConfigSetIterator struct {
	Event *EVM2EVMSubscriptionOnRampOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampOnRampConfigSet)
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
		it.Event = new(EVM2EVMSubscriptionOnRampOnRampConfigSet)
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

func (it *EVM2EVMSubscriptionOnRampOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampOnRampConfigSet struct {
	Config BaseOnRampInterfaceOnRampConfig
	Raw    types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterOnRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampOnRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampOnRampConfigSetIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "OnRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampOnRampConfigSet)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseOnRampConfigSet(log types.Log) (*EVM2EVMSubscriptionOnRampOnRampConfigSet, error) {
	event := new(EVM2EVMSubscriptionOnRampOnRampConfigSet)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMSubscriptionOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMSubscriptionOnRampOwnershipTransferRequested)
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

func (it *EVM2EVMSubscriptionOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMSubscriptionOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampOwnershipTransferRequestedIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampOwnershipTransferRequested)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMSubscriptionOnRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMSubscriptionOnRampOwnershipTransferRequested)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampOwnershipTransferredIterator struct {
	Event *EVM2EVMSubscriptionOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMSubscriptionOnRampOwnershipTransferred)
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

func (it *EVM2EVMSubscriptionOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMSubscriptionOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampOwnershipTransferredIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampOwnershipTransferred)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMSubscriptionOnRampOwnershipTransferred, error) {
	event := new(EVM2EVMSubscriptionOnRampOwnershipTransferred)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampPausedIterator struct {
	Event *EVM2EVMSubscriptionOnRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampPaused)
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
		it.Event = new(EVM2EVMSubscriptionOnRampPaused)
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

func (it *EVM2EVMSubscriptionOnRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampPausedIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampPaused)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParsePaused(log types.Log) (*EVM2EVMSubscriptionOnRampPaused, error) {
	event := new(EVM2EVMSubscriptionOnRampPaused)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampPoolAddedIterator struct {
	Event *EVM2EVMSubscriptionOnRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampPoolAdded)
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
		it.Event = new(EVM2EVMSubscriptionOnRampPoolAdded)
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

func (it *EVM2EVMSubscriptionOnRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampPoolAddedIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampPoolAdded)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMSubscriptionOnRampPoolAdded, error) {
	event := new(EVM2EVMSubscriptionOnRampPoolAdded)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampPoolRemovedIterator struct {
	Event *EVM2EVMSubscriptionOnRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampPoolRemoved)
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
		it.Event = new(EVM2EVMSubscriptionOnRampPoolRemoved)
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

func (it *EVM2EVMSubscriptionOnRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampPoolRemovedIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampPoolRemoved)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMSubscriptionOnRampPoolRemoved, error) {
	event := new(EVM2EVMSubscriptionOnRampPoolRemoved)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampRouterSetIterator struct {
	Event *EVM2EVMSubscriptionOnRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampRouterSet)
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
		it.Event = new(EVM2EVMSubscriptionOnRampRouterSet)
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

func (it *EVM2EVMSubscriptionOnRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterRouterSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampRouterSetIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampRouterSetIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "RouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampRouterSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampRouterSet)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseRouterSet(log types.Log) (*EVM2EVMSubscriptionOnRampRouterSet, error) {
	event := new(EVM2EVMSubscriptionOnRampRouterSet)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMSubscriptionOnRampUnpausedIterator struct {
	Event *EVM2EVMSubscriptionOnRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMSubscriptionOnRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMSubscriptionOnRampUnpaused)
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
		it.Event = new(EVM2EVMSubscriptionOnRampUnpaused)
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

func (it *EVM2EVMSubscriptionOnRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMSubscriptionOnRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMSubscriptionOnRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMSubscriptionOnRampUnpausedIterator{contract: _EVM2EVMSubscriptionOnRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMSubscriptionOnRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMSubscriptionOnRampUnpaused)
				if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMSubscriptionOnRampUnpaused, error) {
	event := new(EVM2EVMSubscriptionOnRampUnpaused)
	if err := _EVM2EVMSubscriptionOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMSubscriptionOnRamp.abi.Events["AFNMaxHeartbeatTimeSet"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseAFNMaxHeartbeatTimeSet(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseAFNSet(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["AllowListEnabledSet"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseAllowListEnabledSet(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["AllowListSet"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseAllowListSet(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["CCIPSendRequested"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseCCIPSendRequested(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["FeeCharged"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseFeeCharged(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["FeedAdded"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseFeedAdded(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["FeedRemoved"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseFeedRemoved(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["FeesWithdrawn"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseFeesWithdrawn(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["OnRampConfigSet"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseOnRampConfigSet(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["Paused"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParsePaused(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParsePoolAdded(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParsePoolRemoved(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["RouterSet"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseRouterSet(log)
	case _EVM2EVMSubscriptionOnRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMSubscriptionOnRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSet) Topic() common.Hash {
	return common.HexToHash("0x72e72b3dfd44fb4d803f52b2d895c7347b912da657d0f77599a0afc26956c251")
}

func (EVM2EVMSubscriptionOnRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMSubscriptionOnRampAllowListEnabledSet) Topic() common.Hash {
	return common.HexToHash("0xccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df032")
}

func (EVM2EVMSubscriptionOnRampAllowListSet) Topic() common.Hash {
	return common.HexToHash("0xf8adc5fee247b62a85f63294cb46e4be61da815e4756bc57672a83b24faf0dda")
}

func (EVM2EVMSubscriptionOnRampCCIPSendRequested) Topic() common.Hash {
	return common.HexToHash("0x73dfb9df8214728e699dbaaf6ba97aa125afaaba83a5d0de7903062e7c5b3139")
}

func (EVM2EVMSubscriptionOnRampFeeCharged) Topic() common.Hash {
	return common.HexToHash("0x945458c62aa39df7a4d87d6c4dbaaab7de5d870c9a1fe40e2b7571d84f158a8d")
}

func (EVM2EVMSubscriptionOnRampFeedAdded) Topic() common.Hash {
	return common.HexToHash("0x037e7fb95c491187e3e2fbb914fac34809e73da6bfe5119bb916b263fb601305")
}

func (EVM2EVMSubscriptionOnRampFeedRemoved) Topic() common.Hash {
	return common.HexToHash("0xa551ef23eb9f5fcdfd41e19414c3eed81c9412d63fa26c01f3902c6431e1950d")
}

func (EVM2EVMSubscriptionOnRampFeesWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x5e110f8bc8a20b65dcc87f224bdf1cc039346e267118bae2739847f07321ffa8")
}

func (EVM2EVMSubscriptionOnRampOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xcc6ce9e57c1de2adf58a81e94b96b43d77ea6973e3f08e6ea4fe83d62ae60e9e")
}

func (EVM2EVMSubscriptionOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMSubscriptionOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMSubscriptionOnRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMSubscriptionOnRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMSubscriptionOnRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMSubscriptionOnRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0xc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15")
}

func (EVM2EVMSubscriptionOnRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMSubscriptionOnRamp *EVM2EVMSubscriptionOnRamp) Address() common.Address {
	return _EVM2EVMSubscriptionOnRamp.address
}

type EVM2EVMSubscriptionOnRampInterface interface {
	CHAINID(opts *bind.CallOpts) (*big.Int, error)

	DESTINATIONCHAINID(opts *bind.CallOpts) (*big.Int, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetAllowlist(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	GetConfig(opts *bind.CallOpts) (BaseOnRampInterfaceOnRampConfig, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetFeed(opts *bind.CallOpts, token common.Address) (common.Address, error)

	GetFeedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetMaxSecondsWithoutAFNHeartbeat(opts *bind.CallOpts) (*big.Int, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetTokenPool(opts *bind.CallOpts, token common.Address) (common.Address, error)

	IsHealthy(opts *bind.CallOpts, timeNow *big.Int) (bool, error)

	IsPool(opts *bind.CallOpts, addr common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, message CCIPEVM2AnySubscriptionMessage, originalSender common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error)

	SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, config BaseOnRampInterfaceOnRampConfig) (*types.Transaction, error)

	SetMaxSecondsWithoutAFNHeartbeat(opts *bind.TransactOpts, newTime *big.Int) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawAccumulatedFees(opts *bind.TransactOpts, feeToken common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	FilterAFNMaxHeartbeatTimeSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSetIterator, error)

	WatchAFNMaxHeartbeatTimeSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSet) (event.Subscription, error)

	ParseAFNMaxHeartbeatTimeSet(log types.Log) (*EVM2EVMSubscriptionOnRampAFNMaxHeartbeatTimeSet, error)

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMSubscriptionOnRampAFNSet, error)

	FilterAllowListEnabledSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampAllowListEnabledSetIterator, error)

	WatchAllowListEnabledSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampAllowListEnabledSet) (event.Subscription, error)

	ParseAllowListEnabledSet(log types.Log) (*EVM2EVMSubscriptionOnRampAllowListEnabledSet, error)

	FilterAllowListSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampAllowListSetIterator, error)

	WatchAllowListSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampAllowListSet) (event.Subscription, error)

	ParseAllowListSet(log types.Log) (*EVM2EVMSubscriptionOnRampAllowListSet, error)

	FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampCCIPSendRequestedIterator, error)

	WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampCCIPSendRequested) (event.Subscription, error)

	ParseCCIPSendRequested(log types.Log) (*EVM2EVMSubscriptionOnRampCCIPSendRequested, error)

	FilterFeeCharged(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampFeeChargedIterator, error)

	WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampFeeCharged) (event.Subscription, error)

	ParseFeeCharged(log types.Log) (*EVM2EVMSubscriptionOnRampFeeCharged, error)

	FilterFeedAdded(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampFeedAddedIterator, error)

	WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampFeedAdded) (event.Subscription, error)

	ParseFeedAdded(log types.Log) (*EVM2EVMSubscriptionOnRampFeedAdded, error)

	FilterFeedRemoved(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampFeedRemovedIterator, error)

	WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampFeedRemoved) (event.Subscription, error)

	ParseFeedRemoved(log types.Log) (*EVM2EVMSubscriptionOnRampFeedRemoved, error)

	FilterFeesWithdrawn(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampFeesWithdrawnIterator, error)

	WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampFeesWithdrawn) (event.Subscription, error)

	ParseFeesWithdrawn(log types.Log) (*EVM2EVMSubscriptionOnRampFeesWithdrawn, error)

	FilterOnRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampOnRampConfigSetIterator, error)

	WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampOnRampConfigSet) (event.Subscription, error)

	ParseOnRampConfigSet(log types.Log) (*EVM2EVMSubscriptionOnRampOnRampConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMSubscriptionOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMSubscriptionOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMSubscriptionOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMSubscriptionOnRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMSubscriptionOnRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMSubscriptionOnRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMSubscriptionOnRampPoolRemoved, error)

	FilterRouterSet(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampRouterSetIterator, error)

	WatchRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampRouterSet) (event.Subscription, error)

	ParseRouterSet(log types.Log) (*EVM2EVMSubscriptionOnRampRouterSet, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMSubscriptionOnRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMSubscriptionOnRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMSubscriptionOnRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
