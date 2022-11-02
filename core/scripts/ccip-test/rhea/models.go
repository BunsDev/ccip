package rhea

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/test-go/testify/require"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/secrets"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

// DefaultGasTipFee is the default gas tip fee of 1 gwei.
var DefaultGasTipFee = big.NewInt(1e9)

// EVMGasSettings specifies the gas configuration for an EVM chain.
type EVMGasSettings struct {
	EIP1559   bool
	GasPrice  *big.Int
	GasTipCap *big.Int
}

type DeploySettings struct {
	DeployAFN            bool
	DeployTokenPools     bool
	DeployRamp           bool
	DeployRouter         bool
	DeployBlobVerifier   bool
	DeployGovernanceDapp bool
	DeployPingPongDapp   bool
	DeployedAt           uint64
}

type EVMChainConfig struct {
	ChainId     *big.Int
	GasSettings EVMGasSettings
	LinkToken   gethcommon.Address

	SupportedTokens map[gethcommon.Address]EVMBridgedToken
	OnRampRouter    gethcommon.Address
	OffRampRouter   gethcommon.Address
	Afn             gethcommon.Address
}

type EVMBridgedToken struct {
	Pool  gethcommon.Address
	Price *big.Int
}

type EVMLaneConfig struct {
	OnRamp       gethcommon.Address
	OffRamp      gethcommon.Address
	BlobVerifier gethcommon.Address

	TokenSender     gethcommon.Address
	MessageReceiver gethcommon.Address
	ReceiverDapp    gethcommon.Address
	GovernanceDapp  gethcommon.Address
	PingPongDapp    gethcommon.Address
}

type EvmDeploymentConfig struct {
	Owner          *bind.TransactOpts
	Client         *ethclient.Client
	Logger         logger.Logger
	DeploySettings DeploySettings

	ChainConfig EVMChainConfig
	LaneConfig  EVMLaneConfig
}

func (chain *EvmDeploymentConfig) SetupChain(t *testing.T, ownerPrivateKey string) {
	chain.Owner = GetOwner(t, ownerPrivateKey, chain.ChainConfig.ChainId, chain.ChainConfig.GasSettings)
	chain.Client = GetClient(t, secrets.GetRPC(chain.ChainConfig.ChainId))
	chain.Logger = logger.TestLogger(t).Named(helpers.ChainName(chain.ChainConfig.ChainId.Int64()))
	chain.Logger.Info("Completed chain setup")
}

func (chain *EvmDeploymentConfig) SetupReadOnlyChain(lggr logger.Logger) error {
	client, err := ethclient.Dial(secrets.GetRPC(chain.ChainConfig.ChainId))
	if err != nil {
		return err
	}
	chain.Logger = lggr
	chain.Client = client

	return nil
}

// GetOwner sets the owner user credentials and ensures a GasTipCap is set for the resulting user.
func GetOwner(t *testing.T, ownerPrivateKey string, chainId *big.Int, gasSettings EVMGasSettings) *bind.TransactOpts {
	ownerKey, err := crypto.HexToECDSA(ownerPrivateKey)
	require.NoError(t, err)
	user, err := bind.NewKeyedTransactorWithChainID(ownerKey, chainId)
	require.NoError(t, err)
	fmt.Println("--- Owner address ")
	fmt.Println(user.From.Hex())
	SetGasFees(user, gasSettings)

	return user
}

// GetClient dials a given EVM client url and returns the resulting client.
func GetClient(t *testing.T, ethUrl string) *ethclient.Client {
	client, err := ethclient.Dial(ethUrl)
	require.NoError(t, err)
	return client
}

// SetGasFees configures the chain client with the given EVMGasSettings. This method is needed for EIP txs
// to function because of the geth-only tip fee method.
func SetGasFees(owner *bind.TransactOpts, config EVMGasSettings) {
	if config.EIP1559 {
		// to not use geth-only tip fee method when EIP1559 is enabled
		// https://github.com/ethereum/go-ethereum/pull/23484
		owner.GasTipCap = config.GasTipCap
	} else {
		owner.GasPrice = config.GasPrice
	}
}

func PrintContractConfig(source *EvmDeploymentConfig, destination *EvmDeploymentConfig) {
	source.Logger.Infof(`
Source chain config

LinkToken:      common.HexToAddress("%s"),
BridgeTokens:   %+v,
TokenPools:     %s,
OnRamp:         common.HexToAddress("%s"),
OnRampRouter:   common.HexToAddress("%s"),
TokenSender:    common.HexToAddress("%s"),
Afn:            common.HexToAddress("%s"),
GovernanceDapp: common.HexToAddress("%s"),
PingPongDapp:   common.HexToAddress("%s"),
	
`,
		source.ChainConfig.LinkToken,
		source.ChainConfig.SupportedTokens,
		source.LaneConfig.OnRamp,
		source.ChainConfig.OnRampRouter,
		source.LaneConfig.TokenSender,
		source.ChainConfig.Afn,
		source.LaneConfig.GovernanceDapp,
		source.LaneConfig.PingPongDapp)

	destination.Logger.Infof(`
Destination chain config

LinkToken:       common.HexToAddress("%s"),
BridgeTokens:    %+v,
OffRamp:         common.HexToAddress("%s"),
OffRampRouter:   common.HexToAddress("%s"),
BlobVerifier:    common.HexToAddress("%s"),	
MessageReceiver: common.HexToAddress("%s"),
ReceiverDapp:    common.HexToAddress("%s"),
Afn:             common.HexToAddress("%s"),
GovernanceDapp:  common.HexToAddress("%s"),
PingPongDapp:    common.HexToAddress("%s"),
`,
		destination.ChainConfig.LinkToken,
		destination.ChainConfig.SupportedTokens,
		destination.LaneConfig.OffRamp,
		destination.ChainConfig.OffRampRouter,
		destination.LaneConfig.BlobVerifier,
		destination.LaneConfig.MessageReceiver,
		destination.LaneConfig.ReceiverDapp,
		destination.ChainConfig.Afn,
		destination.LaneConfig.GovernanceDapp,
		destination.LaneConfig.PingPongDapp)
}
