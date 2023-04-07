package dione

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

func TestGetTokenPricesUSDPipeline(t *testing.T) {
	srcWeth := rhea.EVMBridgedToken{
		Token: common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
		Price: new(big.Int).Mul(big.NewInt(1500), big.NewInt(1e18)),
	}
	dstLink := rhea.EVMBridgedToken{
		Token: common.HexToAddress("0x514910771af9ca656af840dff83e8264ecf986ca"),
		Price: new(big.Int).Mul(big.NewInt(10), big.NewInt(1e18)),
	}
	dstWeth := rhea.EVMBridgedToken{
		Token: common.HexToAddress("0x4200000000000000000000000000000000000006"),
		Price: new(big.Int).Mul(big.NewInt(1500), big.NewInt(1e18)),
	}
	var tt = []struct {
		pipelineTokens []rhea.EVMBridgedToken
		expected       string
	}{
		{
			[]rhea.EVMBridgedToken{dstLink, srcWeth},
			fmt.Sprintf(`merge [type=merge left="{}" right="{\\\"%s\\\":\\\"10000000000000000000\\\",\\\"%s\\\":\\\"1500000000000000000000\\\"}"];`,
				dstLink.Token.Hex(), srcWeth.Token.Hex()),
		},
		{
			[]rhea.EVMBridgedToken{dstLink, dstWeth, srcWeth},
			fmt.Sprintf(`merge [type=merge left="{}" right="{\\\"%s\\\":\\\"10000000000000000000\\\",\\\"%s\\\":\\\"1500000000000000000000\\\",\\\"%s\\\":\\\"1500000000000000000000\\\"}"];`,
				dstLink.Token.Hex(), dstWeth.Token.Hex(), srcWeth.Token.Hex()),
		},
	}

	for _, tc := range tt {
		tc := tc
		a := GetTokenPricesUSDPipeline(tc.pipelineTokens)
		assert.Equal(t, tc.expected, a)
	}
}

func TestGetTokenRealPricesUSDPipeline(t *testing.T) {
	dstLink := rhea.EVMBridgedToken{
		Token:          common.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789"),
		ChainId:        11155111,
		TokenPriceType: rhea.PriceFeeds,
		PriceFeed: rhea.PriceFeed{
			Aggregator: common.HexToAddress("0x5A2734CC0341ea6564dF3D00171cc99C63B1A7d3"),
			Multiplier: big.NewInt(1e10),
		},
	}
	dstWeth := rhea.EVMBridgedToken{
		Token:          common.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789"),
		ChainId:        11155111,
		TokenPriceType: rhea.TokenPrices,
		Price:          new(big.Int).Mul(big.NewInt(1500), big.NewInt(1e18)),
	}
	dstCustom := rhea.EVMBridgedToken{
		Token:   common.HexToAddress("0x779877A7B0D9E8603169DdbDS836e478b4624789"),
		ChainId: 11155111,
		Price:   new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e18)),
	}
	srcWrappedNative := rhea.EVMBridgedToken{
		Token:          common.HexToAddress("0xd00ae08403B9bbb9124bB305C09058E32C39A48c"),
		ChainId:        43113,
		TokenPriceType: rhea.PriceFeeds,
		PriceFeed: rhea.PriceFeed{
			Aggregator: common.HexToAddress("0x6C2441920404835155f33d88faf0545B895871b1"),
			Multiplier: big.NewInt(1e10),
		},
	}

	var tt = []struct {
		pipelineTokens []rhea.EVMBridgedToken
		expected       string
	}{
		{
			[]rhea.EVMBridgedToken{dstLink, srcWrappedNative, dstWeth, dstCustom},
			fmt.Sprintf(`
encode_call_token1_usd  [type="ethabiencode" abi="latestRoundData()"]

call_token1_usd [type="ethcall"
evmChainId=11155111
contract="%s"
data="$(encode_call_token1_usd)"]

decode_result_token1_usd [type="ethabidecode"
abi="uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound"
data="$(call_token1_usd)"]

multiply_token1_usd [type="multiply" input="$(decode_result_token1_usd.answer)" times=10000000000]

encode_call_token1_usd -> call_token1_usd -> decode_result_token1_usd -> multiply_token1_usd

encode_call_token2_usd  [type="ethabiencode" abi="latestRoundData()"]

call_token2_usd [type="ethcall"
evmChainId=43113
contract="%s"
data="$(encode_call_token2_usd)"]

decode_result_token2_usd [type="ethabidecode"
abi="uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound"
data="$(call_token2_usd)"]

multiply_token2_usd [type="multiply" input="$(decode_result_token2_usd.answer)" times=10000000000]

encode_call_token2_usd -> call_token2_usd -> decode_result_token2_usd -> multiply_token2_usd
merge [type=merge left="{}" right="{\\\"%s\\\":$(multiply_token1_usd),\\\"%s\\\":$(multiply_token2_usd),\\\"%s\\\":\\\"1500000000000000000000\\\",\\\"%s\\\":\\\"1000000000000000000000\\\"}"];`,
				dstLink.PriceFeed.Aggregator.Hex(),
				srcWrappedNative.PriceFeed.Aggregator.Hex(),
				dstLink.Token.Hex(), srcWrappedNative.Token.Hex(), dstWeth.Token.Hex(), dstCustom.Token.Hex()),
		},
	}

	for _, tc := range tt {
		tc := tc
		a := GetTokenPricesUSDPipeline(tc.pipelineTokens)
		assert.Equal(t, tc.expected, a)
	}
}
