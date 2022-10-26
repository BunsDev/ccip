package deployments

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/secrets"
)

// Chains

var Prod_Goerli = rhea.EVMChainConfig{
	ChainId: big.NewInt(5),
	EthUrl:  secrets.GoerliEthURL,
	GasSettings: rhea.EVMGasSettings{
		EIP1559:   true,
		GasTipCap: rhea.DefaultGasTipFee,
	},
	LinkToken:     gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
	BridgeTokens:  []gethcommon.Address{gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"), gethcommon.HexToAddress("0x5680dC17bD191EE04d048719b57983335c5E6153")},
	TokenPools:    []gethcommon.Address{gethcommon.HexToAddress("0x4c10d67E4B8e18a67A7606DEFDCe42CCc281D39B"), gethcommon.HexToAddress("0x1fce171011B16F3b0D16198e3F59FD72c091f43B")},
	TokenPrices:   []*big.Int{big.NewInt(1), big.NewInt(1)},
	OnRampRouter:  gethcommon.HexToAddress("0xA189971a2c5AcA0DFC5Ee7a2C44a2Ae27b3CF389"),
	OffRampRouter: gethcommon.HexToAddress("0xb78d314d32EB4B01C459EDE0774cc3b6AF244Dd7"),
	Afn:           gethcommon.HexToAddress("0x8a710bBd77661D168D5A6725bD2E514ba1bFf59d"),
}

var Prod_OptimismGoerli = rhea.EVMChainConfig{
	ChainId: big.NewInt(420),
	EthUrl:  secrets.OptimismGoerliURL,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	LinkToken:     gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
	BridgeTokens:  []gethcommon.Address{gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"), gethcommon.HexToAddress("0xfe628556155F681dd897e3FD029e5ED699a9248E")},
	TokenPools:    []gethcommon.Address{gethcommon.HexToAddress("0xE4aB69C077896252FAFBD49EFD26B5D171A32410"), gethcommon.HexToAddress("0xc5CCb84C3d8eAD52C081dDB24e7Add615c0c9Daf")},
	TokenPrices:   []*big.Int{big.NewInt(1), big.NewInt(1)},
	OnRampRouter:  gethcommon.HexToAddress("0xE591bf0A0CF924A0674d7792db046B23CEbF5f34"),
	OffRampRouter: gethcommon.HexToAddress("0x2b7aB40413DA5077E168546eA376920591Aee8E7"),
	Afn:           gethcommon.HexToAddress("0x4c10d67E4B8e18a67A7606DEFDCe42CCc281D39B"),
}

var Prod_AvaxFuji = rhea.EVMChainConfig{
	ChainId: big.NewInt(43113),
	EthUrl:  secrets.AvaxFujiURL,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	LinkToken:     gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
	BridgeTokens:  []gethcommon.Address{gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"), gethcommon.HexToAddress("0x3C3de1Dd82eA10B664C693C9a3c19645Ab9635EB")},
	TokenPools:    []gethcommon.Address{gethcommon.HexToAddress("0xb6f1Fe2CDE891eFd5Efd2A563C4C2F2549163718"), gethcommon.HexToAddress("0x43A2A4C2ECB74FF45Eca704a14111d8f2B1c0fA0")},
	TokenPrices:   []*big.Int{big.NewInt(1), big.NewInt(1)},
	OnRampRouter:  gethcommon.HexToAddress("0xc0A2c03115d1B48BAA59f676c108EfE5Ba3ee062"),
	OffRampRouter: gethcommon.HexToAddress("0x7d5297c5506ee2A7Ef121Da9bE02b6a6AD30b392"),
	Afn:           gethcommon.HexToAddress("0xb2958D1Bd07448865E555FeeFf32b58D254ffB4C"),
}

// Lanes

var Prod_GoerliToOptimism = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Goerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0x6A14cFB4Ee9B0A3950CBf731c5634FFEc32A324f"),
		OffRamp:         gethcommon.HexToAddress("0xaEbC37b4FE7F325eF2065e09d50f97E11ec1E5E0"),
		BlobVerifier:    gethcommon.HexToAddress("0x701Fe16916dd21EFE2f535CA59611D818B017877"),
		TokenSender:     gethcommon.HexToAddress("0xc3e8bB61e1db9adE45F76237d75AAfaCca2066AF"),
		MessageReceiver: gethcommon.HexToAddress("0xe0D4860bD0429B87f508f0aE8d1789cC0adbbfcA"),
		ReceiverDapp:    gethcommon.HexToAddress("0x84B7B012c95f8A152B44Ab3e952f2dEE424fA8e1"),
		GovernanceDapp:  gethcommon.HexToAddress(""),
		PingPongDapp:    gethcommon.HexToAddress("0x201D1843707764CA2F236bd69E37CCbefF0827D4"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployBlobVerifier:   false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           7802070,
	},
}

var Prod_OptimismToGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0x692ED98151834BA8800462fd5e17737eA29f4f11"),
		OffRamp:         gethcommon.HexToAddress("0xDda968682C04f82F4b812CB98fa6273b6403D388"),
		BlobVerifier:    gethcommon.HexToAddress("0x4A1d9c5a7f9f9de7D5d8eC0f96f7213b0AB953d9"),
		TokenSender:     gethcommon.HexToAddress("0x51298c07eF8849f89552C2B3184741a759d4B37C"),
		MessageReceiver: gethcommon.HexToAddress("0x2321F13659889c2f1e7a62A7700744E36F9C60E5"),
		ReceiverDapp:    gethcommon.HexToAddress("0xA189971a2c5AcA0DFC5Ee7a2C44a2Ae27b3CF389"),
		GovernanceDapp:  gethcommon.HexToAddress(""),
		PingPongDapp:    gethcommon.HexToAddress("0xb6E24bd5376f808a8f4cEf945c96ec5582791255"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployBlobVerifier:   false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           2084429,
	},
}

var Prod_GoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Goerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0x6eA3dE96a33617c3620b7c33c22656f860DDC255"),
		OffRamp:         gethcommon.HexToAddress("0xfb402f1ed3f05B9552eB27E036FA7a70Bd8D9AB5"),
		BlobVerifier:    gethcommon.HexToAddress("0x56eDC4D8367932F0e36B966CbBd95dF48E9DB40F"),
		TokenSender:     gethcommon.HexToAddress("0xC5662F413AffaE59d214FC84BE92B469a92c077C"),
		MessageReceiver: gethcommon.HexToAddress("0x670bAeAa765CA179B82aDAA21947Ff02f819EbC0"),
		ReceiverDapp:    gethcommon.HexToAddress("0x6D984b7515604C27413BEFF5E92b3a1146E84B18"),
		GovernanceDapp:  gethcommon.HexToAddress(""),
		PingPongDapp:    gethcommon.HexToAddress("0x43A2A4C2ECB74FF45Eca704a14111d8f2B1c0fA0"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployBlobVerifier:   false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           7802070,
	},
}

var Prod_AvaxFujiToGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0x17d1399f0558A31b40C6e8997fb356C84CEb7A8C"),
		OffRamp:         gethcommon.HexToAddress("0x01D20791D6713C01666250dE51F81d64d05aEff8"),
		BlobVerifier:    gethcommon.HexToAddress("0x177e068bc512AD99eC73dB6FEB7c731d9fea0CB3"),
		TokenSender:     gethcommon.HexToAddress("0xD6B8378092f590a39C360e8196101290551a66EA"),
		MessageReceiver: gethcommon.HexToAddress("0x4d57C6d8037C65fa66D6231844785a428310a735"),
		ReceiverDapp:    gethcommon.HexToAddress("0x8AB103843ED9D28D2C5DAf5FdB9c3e1CE2B6c876"),
		GovernanceDapp:  gethcommon.HexToAddress(""),
		PingPongDapp:    gethcommon.HexToAddress("0xACD8713E31B2CD1cf936673C4ccb8B5f16156129"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployBlobVerifier:   false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           14800337,
	},
}
