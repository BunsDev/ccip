core/scripts/ccip-test/deploy-ccip-contracts/main.go                                                000644  000765  000024  00000022471 14165357743 023555  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/lock_unlock_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_sender"
	ccip_shared "github.com/smartcontractkit/chainlink/core/scripts/ccip-test/ccip-shared"
)

func main() {
	// This key is used to deploy all contracts on both source and Dest chains
	k := os.Getenv("OWNER_KEY")
	if k == "" {
		panic("must set owner key")
	}

	// This method will deploy all source and Destination chain contracts using the
	// owner key. Only run this of the currently deployed contracts are outdated or
	// when initializing a new chain.
	deploySourceAndDestContracts(
		ccip_shared.Kovan.SetOwner(k).SetupClient(),
		ccip_shared.Rinkeby.SetOwner(k).SetupClient(),
	)
}

func deploySourceAndDestContracts(source *ccip_shared.EvmChainConfig, dest *ccip_shared.EvmChainConfig) {
	// 2 gwei to not use eip1559
	source.Owner.GasTipCap = big.NewInt(2000000000)
	dest.Owner.GasTipCap = big.NewInt(2000000000)

	// After running this code please update the configuration to reflect the newly
	// deployed contract addresses.
	onramp := deployOnramp(source, dest.ChainId, dest.LinkToken)
	fmt.Println("Onramp fully deployed:", onramp.Address().Hex())
	offramp, executor, singleTokenReceiver := deployOfframp(dest, source.ChainId)
	fmt.Println("Offramp fully deployed:", offramp.Address().Hex())

	// Deploy onramp EOA token sender
	eoaTokenSenderAddress, tx, _, err := single_token_sender.DeployEOASingleTokenSender(source.Owner, source.Client, onramp.Address(), singleTokenReceiver)
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), source.Client, tx.Hash(), true)
	fmt.Println("Onramp EOA token sender deployed:", eoaTokenSenderAddress.Hex())

	PrintJobSpecs(onramp.Address(), offramp.Address(), executor.Address())
}

func deployOnramp(source *ccip_shared.EvmChainConfig, offrampChainId *big.Int, offrampLinkTokenAddress common.Address) *single_token_onramp.SingleTokenOnRamp {
	sourcePool := deployLockUnlockPool(source, true)
	afn := deployAFN(source, true)

	// deploy onramp
	onRampAddress, tx, _, err := single_token_onramp.DeploySingleTokenOnRamp(
		source.Owner,                    // user
		source.Client,                   // client
		source.ChainId,                  // source chain id
		source.LinkToken,                // source token
		sourcePool.Address(),            // source pool
		offrampChainId,                  // dest chain id
		offrampLinkTokenAddress,         // remoteToken
		[]common.Address{},              // allow list
		false,                           // enableAllowList
		big.NewInt(1),                   // token bucket rate
		big.NewInt(1000000000000000000), // token bucket capacity, 1 LINK
		afn.Address(),                   // AFN
		// 86400 seconds = one day
		big.NewInt(86400), //maxTimeWithoutAFNSignal
	)
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), source.Client, tx.Hash(), true)

	onRamp, err := single_token_onramp.NewSingleTokenOnRamp(onRampAddress, source.Client)
	ccip_shared.PanicErr(err)
	fmt.Println("Onramp deployed on:", onRampAddress.String())

	// Configure onramp address on pool
	tx, err = sourcePool.SetOnRamp(source.Owner, onRampAddress, true)
	ccip_shared.PanicErr(err)

	fmt.Println("Onramp pool configured with onramp on:", tx.Hash().Hex())
	return onRamp
}

func deployOfframp(dest *ccip_shared.EvmChainConfig, onrampChainId *big.Int) (*single_token_offramp.SingleTokenOffRamp, *message_executor.MessageExecutor, common.Address) {
	pool := deployLockUnlockPool(dest, true)
	fillPoolWithLink(dest, pool)
	afn := deployAFN(dest, true)

	// deploy offramp on Rinkeby
	offrampAddress, tx, _, err := single_token_offramp.DeploySingleTokenOffRamp(
		dest.Owner,                      // user
		dest.Client,                     // client
		onrampChainId,                   // source chain id
		dest.ChainId,                    // dest chain id
		dest.LinkToken,                  // link token address
		pool.Address(),                  // dest pool address
		big.NewInt(1),                   // token bucket rate
		big.NewInt(1000000000000000000), // token bucket capacity
		afn.Address(),                   // AFN address
		// 86400 seconds = one day
		big.NewInt(86400), // max timeout without AFN signal
		big.NewInt(0),     // execution delay in seconds
	)
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), dest.Client, tx.Hash(), true)
	fmt.Println("Offramp deployed on:", offrampAddress.Hex())

	offramp, err := single_token_offramp.NewSingleTokenOffRamp(offrampAddress, dest.Client)
	ccip_shared.PanicErr(err)

	// Configure offramp address on pool
	tx, err = pool.SetOffRamp(dest.Owner, offramp.Address(), true)
	ccip_shared.PanicErr(err)
	fmt.Println("Offramp pool configured with offramp address, tx hash:", tx.Hash().Hex())

	// Deploy offramp contract token receiver
	messageReceiverAddress, tx, _, err := simple_message_receiver.DeploySimpleMessageReceiver(dest.Owner, dest.Client)
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), dest.Client, tx.Hash(), true)
	fmt.Println("Offramp contract message receiver deployed on:", messageReceiverAddress.Hex())

	// Deploy offramp EOA token receiver
	tokenReceiverAddress, tx, _, err := single_token_receiver.DeployEOASingleTokenReceiver(dest.Owner, dest.Client, offramp.Address())
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), dest.Client, tx.Hash(), true)
	fmt.Println("Offramp EOA token receiver deployed on:", tokenReceiverAddress.Hex())
	// Deploy the message executor ocr2 contract
	executorAddress, tx, _, err := message_executor.DeployMessageExecutor(dest.Owner, dest.Client, offramp.Address())
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), dest.Client, tx.Hash(), true)
	fmt.Println("Message executor ocr2 contract deployed on:", executorAddress.Hex())

	executor, err := message_executor.NewMessageExecutor(executorAddress, dest.Client)
	ccip_shared.PanicErr(err)

	return offramp, executor, tokenReceiverAddress
}

func deployLockUnlockPool(client *ccip_shared.EvmChainConfig, deployNew bool) *lock_unlock_pool.LockUnlockPool {
	if deployNew {
		address, tx, _, err := lock_unlock_pool.DeployLockUnlockPool(client.Owner, client.Client, client.LinkToken)
		ccip_shared.PanicErr(err)
		ccip_shared.WaitForMined(context.Background(), client.Client, tx.Hash(), true)
		fmt.Println("Lock/unlock pool deployed on:", address.Hex())
		pool, err := lock_unlock_pool.NewLockUnlockPool(address, client.Client)
		ccip_shared.PanicErr(err)
		return pool
	}
	if client.LockUnlockPool.Hex() == "0x0000000000000000000000000000000000000000" {
		ccip_shared.PanicErr(errors.New("deploy new lock unlock pool set to false but no lock unlock pool given in config"))
	}
	sourcePool, err := lock_unlock_pool.NewLockUnlockPool(client.LockUnlockPool, client.Client)
	ccip_shared.PanicErr(err)
	fmt.Println("Lock unlock pool loaded from:", sourcePool.Address().Hex())
	return sourcePool
}

func deployAFN(client *ccip_shared.EvmChainConfig, deployNew bool) *afn_contract.AFNContract {
	if deployNew {
		address, tx, _, err := afn_contract.DeployAFNContract(
			client.Owner,
			client.Client,
			[]common.Address{client.Owner.From},
			[]*big.Int{big.NewInt(1)},
			big.NewInt(1),
			big.NewInt(1),
		)
		ccip_shared.PanicErr(err)
		ccip_shared.WaitForMined(context.Background(), client.Client, tx.Hash(), true)
		fmt.Println("AFN deployed on:", address.Hex())
		afn, err := afn_contract.NewAFNContract(address, client.Client)
		ccip_shared.PanicErr(err)
		return afn
	}
	if client.Afn.Hex() == "0x0000000000000000000000000000000000000000" {
		ccip_shared.PanicErr(errors.New("deploy new afn set to false but no afn given in config"))
	}
	afn, err := afn_contract.NewAFNContract(client.Afn, client.Client)
	ccip_shared.PanicErr(err)
	fmt.Println("AFN loaded from:", afn.Address().Hex())
	return afn
}

func fillPoolWithLink(client *ccip_shared.EvmChainConfig, pool *lock_unlock_pool.LockUnlockPool) {
	destLinkToken, err := link_token_interface.NewLinkToken(client.LinkToken, client.Client)
	ccip_shared.PanicErr(err)

	// fill offramp token pool with 5 LINK
	amount, _ := new(big.Int).SetString("5000000000000000000", 10)
	tx, err := destLinkToken.Approve(client.Owner, pool.Address(), amount)
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), client.Client, tx.Hash(), true)

	tx, err = pool.LockOrBurn(client.Owner, client.Owner.From, amount)
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), client.Client, tx.Hash(), true)
	fmt.Println("Dest pool filled with funds, tx hash:", tx.Hash().Hex())
}
                                                                                                                                                                                                       core/scripts/ccip-test/ccip-shared/client_configuration.go                                          000644  000765  000024  00000014010 14165357743 025000  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip_shared

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/lock_unlock_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_sender"
)

type Client struct {
	Owner            *bind.TransactOpts
	Users            []*bind.TransactOpts
	Client           *ethclient.Client
	ChainId          *big.Int
	LinkToken        *link_token_interface.LinkToken
	LinkTokenAddress common.Address
	LockUnlockPool   *lock_unlock_pool.LockUnlockPool
	Afn              *afn_contract.AFNContract
}

type SourceClient struct {
	Client
	SingleTokenOnramp *single_token_onramp.SingleTokenOnRamp
	SingleTokenSender *single_token_sender.EOASingleTokenSender
}

func NewSourceClient(config *EvmChainConfig) SourceClient {
	LinkToken, err := link_token_interface.NewLinkToken(config.LinkToken, config.Client)
	PanicErr(err)
	lockUnlockPool, err := lock_unlock_pool.NewLockUnlockPool(config.LockUnlockPool, config.Client)
	PanicErr(err)
	afn, err := afn_contract.NewAFNContract(config.Afn, config.Client)
	PanicErr(err)
	singleTokenOnramp, err := single_token_onramp.NewSingleTokenOnRamp(config.SingleTokenOnramp, config.Client)
	PanicErr(err)
	simpleTokenSender, err := single_token_sender.NewEOASingleTokenSender(config.SingleTokenSender, config.Client)
	PanicErr(err)

	return SourceClient{
		Client: Client{
			Client:           config.Client,
			Users:            config.Users,
			ChainId:          config.ChainId,
			LinkTokenAddress: config.LinkToken,
			LinkToken:        LinkToken,
			Afn:              afn,
			LockUnlockPool:   lockUnlockPool,
		},
		SingleTokenOnramp: singleTokenOnramp,
		SingleTokenSender: simpleTokenSender,
	}
}

type DestClient struct {
	Client
	SingleTokenOfframp    *single_token_offramp.SingleTokenOffRamp
	SimpleMessageReceiver *simple_message_receiver.SimpleMessageReceiver
	SingleTokenReceiver   *single_token_receiver.EOASingleTokenReceiver
	MessageExecutor       *message_executor.MessageExecutor
}

func NewDestinationClient(config *EvmChainConfig) DestClient {
	LinkToken, err := link_token_interface.NewLinkToken(config.LinkToken, config.Client)
	PanicErr(err)
	lockUnlockPool, err := lock_unlock_pool.NewLockUnlockPool(config.LockUnlockPool, config.Client)
	PanicErr(err)
	afn, err := afn_contract.NewAFNContract(config.Afn, config.Client)
	PanicErr(err)
	singleTokenOfframp, err := single_token_offramp.NewSingleTokenOffRamp(config.SingleTokenOfframp, config.Client)
	PanicErr(err)
	messageExecutor, err := message_executor.NewMessageExecutor(config.MessageExecutor, config.Client)
	PanicErr(err)
	simpleMessageReceiver, err := simple_message_receiver.NewSimpleMessageReceiver(config.SimpleMessageReceiver, config.Client)
	PanicErr(err)
	singleTokenReceiver, err := single_token_receiver.NewEOASingleTokenReceiver(config.SingleTokenReceiver, config.Client)
	PanicErr(err)

	return DestClient{
		Client: Client{
			Owner:            config.Owner,
			Users:            config.Users,
			Client:           config.Client,
			ChainId:          config.ChainId,
			LinkTokenAddress: config.LinkToken,
			LinkToken:        LinkToken,
			LockUnlockPool:   lockUnlockPool,
			Afn:              afn,
		},
		SingleTokenOfframp:    singleTokenOfframp,
		SimpleMessageReceiver: simpleMessageReceiver,
		SingleTokenReceiver:   singleTokenReceiver,
		MessageExecutor:       messageExecutor,
	}
}

type CcipClient struct {
	Source SourceClient
	Dest   DestClient
}

func NewCcipClient(sourceConfig *EvmChainConfig, destConfig *EvmChainConfig) CcipClient {
	// to not use geth-only tip fee method
	// https://github.com/ethereum/go-ethereum/pull/23484
	var twoGwei = big.NewInt(2e9)
	sourceConfig.Owner.GasTipCap = twoGwei
	destConfig.Owner.GasTipCap = twoGwei
	for _, user := range sourceConfig.Users {
		user.GasTipCap = twoGwei
	}
	for _, user := range destConfig.Users {
		user.GasTipCap = twoGwei
	}

	source := NewSourceClient(sourceConfig.SetupClient())
	dest := NewDestinationClient(destConfig.SetupClient())

	return CcipClient{
		Source: source,
		Dest:   dest,
	}
}

func (client Client) AssureHealth() {
	standardAfnTimeout := int64(86400)
	status, err := client.Afn.GetLastHeartbeat(&bind.CallOpts{
		Pending: false,
		Context: nil,
	})
	PanicErr(err)
	timeNow := time.Now().Unix()

	if timeNow > status.Timestamp.Int64()+standardAfnTimeout {
		tx, err := client.Afn.VoteGood(client.Owner, big.NewInt(status.Round.Int64()+1))
		PanicErr(err)
		WaitForMined(context.Background(), client.Client, tx.Hash(), true)
		fmt.Printf("[HEALTH] client with chainId %d set healthy for %d hours\n", client.ChainId.Int64(), standardAfnTimeout/60/60)
	} else {
		fmt.Printf("[HEALTH] client with chainId %d is already healthy for %d more hours\n", client.ChainId.Int64(), (standardAfnTimeout-(timeNow-status.Timestamp.Int64()))/60/60)
	}
}

func (client Client) ApproveLinkFrom(user *bind.TransactOpts, approvedFor common.Address, amount *big.Int) {
	ctx := context.Background()
	tx, err := client.LinkToken.Approve(user, approvedFor, amount)
	PanicErr(err)

	WaitForMined(ctx, client.Client, tx.Hash(), true)
	fmt.Println("approve tx hash", tx.Hash().Hex())
}

func (client Client) ApproveLink(approvedFor common.Address, amount *big.Int) {
	client.ApproveLinkFrom(client.Owner, approvedFor, amount)
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        core/scripts/ccip-test/ccip-shared/chain_configurations.go                                          000644  000765  000024  00000013613 14165357743 024777  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip_shared

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"
)

type EvmChainConfig struct {
	Owner                 *bind.TransactOpts
	Users                 []*bind.TransactOpts
	Client                *ethclient.Client
	ChainId               *big.Int
	LinkToken             common.Address
	SingleTokenOnramp     common.Address
	SingleTokenOfframp    common.Address
	LockUnlockPool        common.Address
	SimpleMessageReceiver common.Address
	SingleTokenSender     common.Address
	SingleTokenReceiver   common.Address
	MessageExecutor       common.Address
	Afn                   common.Address
	EthUrl                string
}

var Kovan = EvmChainConfig{
	LinkToken:          common.HexToAddress("0xa36085F69e2889c224210F603D836748e7dC0088"),
	SingleTokenOfframp: common.Address{},
	SingleTokenOnramp:  common.HexToAddress("0x7590f49f088a7B74596712A1DB4C39D9FAB00347"),
	SingleTokenSender:  common.HexToAddress("0xCfBc79f4042Be99a7292Ff33466cEE6652F40485"),
	LockUnlockPool:     common.HexToAddress("0x0C710E14226A43301028cEf2D0D492fDE7E3024A"),
	Afn:                common.HexToAddress("0xB86654A84CF21a913f39d9Da126C27f56Df07166"),
	EthUrl:             "wss://parity-kovan.eth.devnet.tools/ws",
	ChainId:            big.NewInt(42),
}

var Rinkeby = EvmChainConfig{
	LinkToken:             common.HexToAddress("0x01be23585060835e02b77ef475b0cc51aa1e0709"),
	SingleTokenOfframp:    common.HexToAddress("0xC3376eD981978E0C107a52f0488f1d131B9ECCAc"),
	SingleTokenOnramp:     common.Address{},
	LockUnlockPool:        common.HexToAddress("0x33079B10A1417EF666040BF5aAF5623FCc90FAFe"),
	SimpleMessageReceiver: common.HexToAddress("0x0389eF5B01822F673cFb87cdf3D8f97E0FaDBf77"),
	SingleTokenReceiver:   common.HexToAddress("0xF47C5C5cEeE3F77954Fa2eA58690e44fD6658B9F"),
	MessageExecutor:       common.HexToAddress("0x21d5C93B2A22Bdc315F1760E92960Cda23D93f3E"),
	Afn:                   common.HexToAddress("0x1E275452a2bD9154EC0F46aE21881E47Aed03E3e"),
	EthUrl:                "wss://geth-rinkeby.eth.devnet.tools/ws",
	ChainId:               big.NewInt(4),
}

const BootstrapPeerID = "12D3KooWFHTQLnS1dzmRoqit8zyLx7ost7sm8pSjFQSByfjsoqyT"

var Oracles = []confighelper2.OracleIdentityExtra{
	{
		// Node 0
		OracleIdentity: confighelper2.OracleIdentity{
			OnchainPublicKey:  common.HexToAddress("0x69B8fADd511A2BE6d90A5dA5F617EB48cE3FA132").Bytes(),
			TransmitAccount:   ocrtypes2.Account("0x1b9aC605d2b2E2E9Db4cac561181Ec10A938390c"),
			OffchainPublicKey: hexutil.MustDecode("0x17992ca120fe8a3075e6c8b3e8c93f06fc3fc5dc5f989d54ec14def8cf080d06"),
			PeerID:            "12D3KooWPRpNDEzJKJevcwhdjKvTWEBV4o9RFJ8FmzPf9ErsPtBM",
		},
		ConfigEncryptionPublicKey: stringTo32Bytes("0x69a21497b875787e4810d2d825aefca5f9ee6dc3e97f51b93b33de67300c402f"),
	},
	{
		// Node 1
		OracleIdentity: confighelper2.OracleIdentity{
			OnchainPublicKey:  common.HexToAddress("0x51A4282729AFE2A7967ab24ff707AffCe1dcc678").Bytes(),
			TransmitAccount:   ocrtypes2.Account("0x2FF79Fff751a157054629eECF2B32aE671d72Bf8"),
			OffchainPublicKey: hexutil.MustDecode("0xd7f949bb2ff6242f2d5158b2f54eb0b629904dddfaa9d699736e7265eb87bb2f"),
			PeerID:            "12D3KooWAPnKdfa3wPobf3FdErZu1VAKKMmuoEHwvmcjnSQhYSvD",
		},
		ConfigEncryptionPublicKey: stringTo32Bytes("0x4320cf6a9be0ffdd4e44787551bfda49950288c31d6854ba5f243e9ea23e5278"),
	},
	{
		// Node 2
		OracleIdentity: confighelper2.OracleIdentity{
			OnchainPublicKey:  common.HexToAddress("0x9d51eeF5292d2fFE9bEa7c263CF1fe18e9f35148").Bytes(),
			TransmitAccount:   ocrtypes2.Account("0xC81C5cccfcA5B95526609575235D55077A25F105"),
			OffchainPublicKey: hexutil.MustDecode("0x65b165e268405827411a79384bae8648f7725d701bc4d8373fdd55838802e4f6"),
			PeerID:            "12D3KooWJnTuDhN1GCSbxWjNW51P7z1QRbC3VJbtY6wuL1VUXkQu",
		},
		ConfigEncryptionPublicKey: stringTo32Bytes("0xc4717af64f5e4235c07e893159c522b12dc0809982f09f519d873d6194129a43"),
	},
	{
		// Node 3
		OracleIdentity: confighelper2.OracleIdentity{
			OnchainPublicKey:  common.HexToAddress("0xaaeB8784265a6ee8181729dDD0Aea99c60814482").Bytes(),
			TransmitAccount:   ocrtypes2.Account("0x1FD884B9088d2013B6c2EC2F9640F551578e2f1C"),
			OffchainPublicKey: hexutil.MustDecode("0x61c4c6a6e9a2ac020e87e2e7e8c88e32373a503a1ae7d1a651b1ac08bb7c31f5"),
			PeerID:            "12D3KooWSDzVm7Kv3xSHB17aUQ5UvBJay2cxC7XfTjGsqGn7MDK7",
		},
		ConfigEncryptionPublicKey: stringTo32Bytes("0xe0b8876b62cb1c5c827be6f6dc271ce7702f06a8bf7d5c289486b2a6c8a21e19"),
	},
}

func stringTo32Bytes(s string) [32]byte {
	var b [32]byte
	copy(b[:], hexutil.MustDecode(s))
	return b
}

func (config *EvmChainConfig) SetupClient() *EvmChainConfig {
	client, err := ethclient.Dial(config.EthUrl)
	PanicErr(err)
	config.Client = client
	return config
}

func (config *EvmChainConfig) SetOwnerAndUsers(ownerPrivateKey string, seedKey string) *EvmChainConfig {
	config.SetOwner(ownerPrivateKey)

	var users []*bind.TransactOpts
	seedKeyWithoutFirstChar := seedKey[1:]
	fmt.Println("--- Addresses of the seed key")
	for i := 0; i <= 9; i++ {
		key, err := crypto.HexToECDSA(strconv.Itoa(i) + seedKeyWithoutFirstChar)
		PanicErr(err)
		user, err := bind.NewKeyedTransactorWithChainID(key, config.ChainId)
		PanicErr(err)
		users = append(users, user)
		fmt.Println(user.From.Hex())
	}
	fmt.Println("---")

	config.Users = users

	return config
}

func (config *EvmChainConfig) SetOwner(ownerPrivateKey string) *EvmChainConfig {
	ownerKey, err := crypto.HexToECDSA(ownerPrivateKey)
	PanicErr(err)
	user, err := bind.NewKeyedTransactorWithChainID(ownerKey, config.ChainId)
	PanicErr(err)
	fmt.Println("--- Owner address ")
	fmt.Println(user.From.Hex())
	config.Owner = user
	return config
}
                                                                                                                     core/scripts/ccip-test/ccip-shared/helpers.go                                                       000644  000765  000024  00000001562 14165357743 022245  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip_shared

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const RetryTiming = 5 * time.Second

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func WaitForMined(ctx context.Context, client *ethclient.Client, hash common.Hash, shouldSucceed bool) {
	for {
		fmt.Println("[MINING] waiting for tx to be mined...")
		receipt, _ := client.TransactionReceipt(ctx, hash)

		if receipt != nil {
			if shouldSucceed && receipt.Status == 0 {
				fmt.Println("[MINING] ERROR tx reverted!", hash.Hex())
				panic(receipt)
			} else if !shouldSucceed && receipt.Status != 0 {
				fmt.Println("[MINING] ERROR expected tx to revert!", hash.Hex())
				panic(receipt)
			}
			fmt.Println("[MINING] tx mined", hash.Hex(), "successful", shouldSucceed)
			break
		}

		time.Sleep(RetryTiming)
	}
}
                                                                                                                                              core/scripts/ccip-test/ccip-test-functions/main.go                                                  000644  000765  000024  00000053117 14165357743 023251  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services/ccip/abihelpers"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	ccipshared "github.com/smartcontractkit/chainlink/core/scripts/ccip-test/ccip-shared"
	"github.com/smartcontractkit/chainlink/core/services/ccip"
)

type ccipClient ccipshared.CcipClient

func main() {
	// This key is used to deploy all contracts on both source and Dest chains
	k := os.Getenv("OWNER_KEY")
	if k == "" {
		panic("must set owner key")
	}

	// The seed key is used to generate 10 keys from a single key by changing the
	// first character of the given seed with the digits 0-9
	seedKey := os.Getenv("SEED_KEY")
	if seedKey == "" {
		panic("must set seed key")
	}

	// Configures a client to run tests with using the network defaults and given keys.
	// After updating any contracts be sure to update the network defaults to reflect
	// those changes.
	client := ccipClient(ccipshared.NewCcipClient(
		// Source chain
		ccipshared.Kovan.SetOwnerAndUsers(k, seedKey),
		// Dest chain
		ccipshared.Rinkeby.SetOwnerAndUsers(k, seedKey)))

	client.Source.Client.AssureHealth()
	client.Dest.Client.AssureHealth()
	client.unpauseAll()

	// Set the config to the message executor and the offramp
	//client.setConfig()

	// Cross chain request with the client manually proving and executing the transaction
	//client.externalExecutionHappyPath()

	// Executing the same request twice should fail
	//client.externalExecutionSubmitOfframpTwiceShouldFail()

	// Cross chain request with DON execution
	//client.donExecutionHappyPath()

	// Submit 10 txs. This should result in the txs being batched together
	//client.scalingAndBatching()

	// Should not be able to send funds greater than the amount in the bucket
	//client.notEnoughFundsInBucketShouldFail()

	//client.tryGetTokensFromPausedPool()

	client.crossChainSendPausedOfframpShouldFail()

	//client.crossChainSendPausedOnrampShouldFail()
}

func (client ccipClient) sendMessage() {
	// ABI encoded message
	bytes, _ := hex.DecodeString("00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000005626c616e6b000000000000000000000000000000000000000000000000000000")

	msg := single_token_onramp.CCIPMessagePayload{
		Receiver: client.Dest.SimpleMessageReceiver.Address(),
		Data:     bytes,
		Tokens:   []common.Address{client.Source.LinkToken.Address()},
		Amounts:  []*big.Int{big.NewInt(1)},
		Options:  []byte{},
		Executor: client.Dest.MessageExecutor.Address(),
	}

	client.Source.ApproveLink(client.Source.LockUnlockPool.Address(), big.NewInt(1))
	tx, err := client.Source.SingleTokenOnramp.RequestCrossChainSend(client.Source.Owner, msg)
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client ccipClient) donExecutionHappyPath() {
	amount := big.NewInt(100)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)
	DestBlockNum := getCurrentBlockNumber(client.Dest.Client.Client)
	crossChainRequest := client.sendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, client.Dest.MessageExecutor.Address())
	fmt.Println("Don executed tx submitted with sequence number: ", crossChainRequest.Message.SequenceNumber.Int64())
	fmt.Println("Waiting for Destination funds transfer...")

	events := make(chan *single_token_offramp.SingleTokenOffRampCrossChainMessageExecuted)
	sub, err := client.Dest.SingleTokenOfframp.WatchCrossChainMessageExecuted(
		&bind.WatchOpts{
			Context: context.Background(),
			Start:   &DestBlockNum,
		},
		events,
		[]*big.Int{crossChainRequest.Message.SequenceNumber})
	ccipshared.PanicErr(err)
	defer sub.Unsubscribe()

	select {
	case event := <-events:
		fmt.Printf("found Destination execution in transaction: %s\n", event.Raw.TxHash.Hex())
		return
	case err := <-sub.Err():
		panic(err)
	}
}

func (client ccipClient) externalExecutionHappyPath() {
	ctx := context.Background()
	offrampBlockNumber := getCurrentBlockNumber(client.Dest.Client.Client)
	onrampBlockNumber := getCurrentBlockNumber(client.Source.Client.Client)

	amount, _ := new(big.Int).SetString("10", 10)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)

	onrampRequest := client.sendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
	sequenceNumber := onrampRequest.Message.SequenceNumber.Int64()

	// Gets the report that our transaction is included in
	fmt.Println("Getting report")
	report := client.getReportForSequenceNumber(ctx, sequenceNumber, offrampBlockNumber)

	// Get all requests included in the given report
	fmt.Println("Getting recent cross chain requests")
	requests := client.getCrossChainSendRequestsForRange(ctx, report, onrampBlockNumber)

	// Generate the proof
	fmt.Println("Generating proof")
	proof := client.validateMerkleRoot(onrampRequest, requests, report)

	// Execute the transaction on the offramp
	client.Dest.Owner.GasLimit = 2e9
	fmt.Println("Executing offramp TX")
	tx, err := client.executeOfframpTransaction(proof, onrampRequest.Raw.Data)
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), true)
}

func (client ccipClient) crossChainSendPausedOnrampShouldFail() {
	client.pauseOnramp()
	amount := big.NewInt(100)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)
	client.Source.Owner.GasLimit = 1e6
	tx, err := client.Source.SingleTokenSender.SendTokens(client.Source.Owner, client.Dest.Owner.From, amount, client.Dest.MessageExecutor.Address())
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), false)
}

func (client ccipClient) crossChainSendPausedOfframpShouldFail() {
	client.pauseOfframp()
	ctx := context.Background()
	offrampBlockNumber := getCurrentBlockNumber(client.Dest.Client.Client)

	amount, _ := new(big.Int).SetString("10", 10)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)

	onrampRequest := client.sendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
	sequenceNumber := onrampRequest.Message.SequenceNumber.Int64()

	fmt.Println("Waiting for report...")
	result := make(chan single_token_offramp.CCIPRelayReport)
	go func() {
		result <- client.getReportForSequenceNumber(ctx, sequenceNumber, offrampBlockNumber)
	}()

	select {
	case r := <-result:
		panic(fmt.Errorf("report found despite paused offramp: %+v", r))
	case <-time.After(time.Minute):
		fmt.Println("Success, no oracle report sent to paused offramp.")
	}
}

func (client ccipClient) notEnoughFundsInBucketShouldFail() {
	amount := big.NewInt(2e18) // 2 LINK, bucket size is 1 LINK
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)
	client.Source.Owner.GasLimit = 1e6
	tx, err := client.Source.SingleTokenSender.SendTokens(client.Source.Owner, client.Dest.Owner.From, amount, client.Dest.MessageExecutor.Address())
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), false)
}

func (client ccipClient) externalExecutionSubmitOfframpTwiceShouldFail() {
	ctx := context.Background()
	offrampBlockNumber := getCurrentBlockNumber(client.Dest.Client.Client)
	onrampBlockNumber := getCurrentBlockNumber(client.Source.Client.Client)

	amount, _ := new(big.Int).SetString("10", 10)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)

	onrampRequest := client.sendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
	sequenceNumber := onrampRequest.Message.SequenceNumber.Int64()

	// Gets the report that our transaction is included in
	fmt.Println("Getting report")
	report := client.getReportForSequenceNumber(ctx, sequenceNumber, offrampBlockNumber)

	// Get all requests included in the given report
	fmt.Println("Getting recent cross chain requests")
	requests := client.getCrossChainSendRequestsForRange(ctx, report, onrampBlockNumber)

	// Generate the proof
	fmt.Println("Generating proof")
	proof := client.validateMerkleRoot(onrampRequest, requests, report)

	// Execute the transaction on the offramp
	fmt.Println("Executing first offramp TX - should succeed")
	tx, err := client.executeOfframpTransaction(proof, onrampRequest.Raw.Data)
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), true)

	// Execute the transaction on the offramp
	fmt.Println("Executing second offramp TX - should fail")
	client.Dest.Owner.GasLimit = 1e6
	tx, err = client.executeOfframpTransaction(proof, onrampRequest.Raw.Data)
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), false)
}

// Scale so that we see batching on the nodes
func (client ccipClient) scalingAndBatching() {
	amount := big.NewInt(10)
	toAddress := common.HexToAddress("0x57359120D900fab8cE74edC2c9959b21660d3887")

	var wg sync.WaitGroup
	for _, user := range client.Source.Users {
		wg.Add(1)
		go func(user *bind.TransactOpts) {
			defer wg.Done()
			client.Source.ApproveLinkFrom(user, client.Source.SingleTokenSender.Address(), amount)
			crossChainRequest := client.sendToOnrampWithExecution(client.Source, user, toAddress, amount, client.Dest.MessageExecutor.Address())
			fmt.Println("Don executed tx submitted with sequence number: ", crossChainRequest.Message.SequenceNumber.Int64())
		}(user)
	}
	wg.Wait()
	fmt.Println("Sent 10 txs to onramp.")
}

func (client ccipClient) executeOfframpTransaction(proof ccip.MerkleProof, encodedMessage []byte) (*types.Transaction, error) {
	decodedMsg, err := abihelpers.DecodeCCIPMessage(encodedMessage)
	ccipshared.PanicErr(err)
	_, err = abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{*decodedMsg})
	ccipshared.PanicErr(err)

	tx, err := client.Dest.SingleTokenOfframp.ExecuteTransaction(client.Dest.Owner, proof.PathForExecute(), *decodedMsg, proof.Index())
	return tx, errors.Wrap(err, "executing offramp tx")
}

func (client ccipClient) getCrossChainSendRequestsForRange(
	ctx context.Context,
	report single_token_offramp.CCIPRelayReport,
	onrampBlockNumber uint64) []*single_token_onramp.SingleTokenOnRampCrossChainSendRequested {
	// Get the other transactions in the proof, we look 1000 blocks back for transaction
	// should be fine? Needs fine-tuning after improved batching strategies are developed
	// in milestone 4
	reqsIterator, err := client.Source.SingleTokenOnramp.FilterCrossChainSendRequested(&bind.FilterOpts{
		Context: ctx,
		Start:   onrampBlockNumber - 1000,
	})
	ccipshared.PanicErr(err)

	var requests []*single_token_onramp.SingleTokenOnRampCrossChainSendRequested

	var minFound = report.MaxSequenceNumber.Int64()
	for reqsIterator.Next() {
		num := reqsIterator.Event.Message.SequenceNumber.Int64()
		if num < minFound {
			minFound = num
		}
		if num >= report.MinSequenceNumber.Int64() && num <= report.MaxSequenceNumber.Int64() {
			requests = append(requests, reqsIterator.Event)
		}
	}

	// TODO: Even if this check passes, we may not have fetched all necessary requests if
	// minFound == report.MinSequenceNumber
	if minFound > report.MinSequenceNumber.Int64() {
		ccipshared.PanicErr(errors.New("Not all cross chain requests found in the last 1000 blocks"))
	}

	return requests
}

func (client ccipClient) getReportForSequenceNumber(ctx context.Context, sequenceNumber int64, minBlockNumber uint64) single_token_offramp.CCIPRelayReport {
	report, err := client.Dest.SingleTokenOfframp.GetLastReport(&bind.CallOpts{
		Pending: false,
	})
	ccipshared.PanicErr(err)

	// our tx is in the latest report
	if sequenceNumber >= report.MinSequenceNumber.Int64() && sequenceNumber <= report.MaxSequenceNumber.Int64() {
		return report
	}
	// report isn't out yet, it will be in a future report
	if sequenceNumber > report.MaxSequenceNumber.Int64() {
		for {
			report, err = client.Dest.SingleTokenOfframp.GetLastReport(&bind.CallOpts{
				Pending: false,
				Context: ctx,
			})
			ccipshared.PanicErr(err)
			if sequenceNumber >= report.MinSequenceNumber.Int64() && sequenceNumber <= report.MaxSequenceNumber.Int64() {
				return report
			}
			time.Sleep(ccipshared.RetryTiming)
		}
	}

	// it is in a past report, start looking at the earliest block number possible, the one
	// before we started the entire transaction on the onramp.
	reports, err := client.Dest.SingleTokenOfframp.FilterReportAccepted(&bind.FilterOpts{
		Start:   minBlockNumber,
		End:     nil,
		Context: ctx,
	})
	ccipshared.PanicErr(err)

	for reports.Next() {
		report = reports.Event.Report
		if sequenceNumber >= report.MinSequenceNumber.Int64() && sequenceNumber <= report.MaxSequenceNumber.Int64() {
			return report
		}
	}

	// Somehow the transaction was not included in any report within blocks produced after
	// the transaction was initialized but the sequence number is lower than we are currently at
	ccipshared.PanicErr(errors.New("No report found"))
	return single_token_offramp.CCIPRelayReport{}
}

func getCurrentBlockNumber(chain *ethclient.Client) uint64 {
	blockNumber, err := chain.BlockNumber(context.Background())
	ccipshared.PanicErr(err)
	return blockNumber
}

func (client ccipClient) validateMerkleRoot(
	request *single_token_onramp.SingleTokenOnRampCrossChainSendRequested,
	reportRequests []*single_token_onramp.SingleTokenOnRampCrossChainSendRequested,
	report single_token_offramp.CCIPRelayReport,
) ccip.MerkleProof {
	var leaves [][]byte
	for _, req := range reportRequests {
		leaves = append(leaves, req.Raw.Data)
	}

	index := big.NewInt(0).Sub(request.Message.SequenceNumber, report.MinSequenceNumber)
	fmt.Println("index is", index)
	root, proof := ccip.GenerateMerkleProof(32, leaves, int(index.Int64()))
	if !reflect.DeepEqual(root[:], report.MerkleRoot[:]) {
		ccipshared.PanicErr(errors.New("Merkle root does not match the report"))
	}

	genRoot := ccip.GenerateMerkleRoot(leaves[int(index.Int64())], proof)
	if !reflect.DeepEqual(root[:], genRoot[:]) {
		ccipshared.PanicErr(errors.New("Root does not verify"))
	}

	exists, err := client.Dest.SingleTokenOfframp.GetMerkleRoot(nil, root)
	ccipshared.PanicErr(err)
	if exists.Uint64() < 1 {
		ccipshared.PanicErr(errors.New("Proof is not present in the offramp"))
	}
	return proof
}

func (client ccipClient) tryGetTokensFromPausedPool() {
	client.pauseOnrampPool()

	paused, err := client.Source.LockUnlockPool.Paused(nil)
	ccipshared.PanicErr(err)
	if !paused {
		ccipshared.PanicErr(errors.New("Should be paused"))
	}

	client.Source.Owner.GasLimit = 2e6
	tx, err := client.Source.LockUnlockPool.LockOrBurn(client.Source.Owner, client.Source.Owner.From, big.NewInt(1000))
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), false)
}

func (client ccipClient) sendToOnrampWithExecution(Source ccipshared.SourceClient, from *bind.TransactOpts, toAddress common.Address, amount *big.Int, executor common.Address) *single_token_onramp.SingleTokenOnRampCrossChainSendRequested {
	ctx := context.Background()
	SourceBlockNumber := getCurrentBlockNumber(Source.Client.Client)

	tx, err := Source.SingleTokenSender.SendTokens(from, toAddress, amount, executor)
	ccipshared.PanicErr(err)
	fmt.Println("send tokens hash ", tx.Hash())
	ccipshared.WaitForMined(ctx, Source.Client.Client, tx.Hash(), true)

	return waitForCrossChainSendRequest(Source, SourceBlockNumber, tx.Hash())
}

func waitForCrossChainSendRequest(Source ccipshared.SourceClient, fromBlockNum uint64, txhash common.Hash) *single_token_onramp.SingleTokenOnRampCrossChainSendRequested {
	filter := bind.FilterOpts{Start: fromBlockNum}
	for {
		iterator, err := Source.SingleTokenOnramp.FilterCrossChainSendRequested(&filter)
		ccipshared.PanicErr(err)
		for iterator.Next() {
			if iterator.Event.Raw.TxHash.Hex() == txhash.Hex() {
				fmt.Println("cross chain send event found in tx: ", txhash.Hex())
				return iterator.Event
			}
		}
		time.Sleep(ccipshared.RetryTiming)
	}
}

func (client ccipClient) pauseOfframpPool() {
	paused, err := client.Dest.LockUnlockPool.Paused(nil)
	ccipshared.PanicErr(err)
	if paused {
		return
	}
	fmt.Println("pausing offramp pool...")
	tx, err := client.Dest.LockUnlockPool.Pause(client.Dest.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Offramp pool paused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Dest.Client.Client, tx.Hash(), true)
}

func (client ccipClient) pauseOnrampPool() {
	paused, err := client.Source.LockUnlockPool.Paused(nil)
	ccipshared.PanicErr(err)
	if paused {
		return
	}
	fmt.Println("pausing onramp pool...")
	tx, err := client.Source.LockUnlockPool.Pause(client.Source.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Onramp pool paused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client ccipClient) unpauseOfframpPool() {
	paused, err := client.Dest.LockUnlockPool.Paused(nil)
	ccipshared.PanicErr(err)
	if !paused {
		return
	}
	fmt.Println("unpausing offramp pool...")
	tx, err := client.Dest.LockUnlockPool.Unpause(client.Dest.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Offramp pool unpaused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Dest.Client.Client, tx.Hash(), true)
}

func (client ccipClient) unpauseOnrampPool() {
	paused, err := client.Source.LockUnlockPool.Paused(nil)
	ccipshared.PanicErr(err)
	if !paused {
		return
	}
	fmt.Println("unpausing onramp pool...")
	tx, err := client.Source.LockUnlockPool.Unpause(client.Source.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Onramp pool unpaused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client ccipClient) pauseOnramp() {
	paused, err := client.Source.SingleTokenOnramp.Paused(nil)
	ccipshared.PanicErr(err)
	if paused {
		return
	}
	fmt.Println("pausing onramp...")
	tx, err := client.Source.SingleTokenOnramp.Pause(client.Source.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Onramp paused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client ccipClient) pauseOfframp() {
	paused, err := client.Dest.SingleTokenOfframp.Paused(nil)
	ccipshared.PanicErr(err)
	if paused {
		return
	}
	fmt.Println("pausing offramp...")
	tx, err := client.Dest.SingleTokenOfframp.Pause(client.Dest.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Offramp paused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Dest.Client.Client, tx.Hash(), true)
}

func (client ccipClient) unpauseOnramp() {
	paused, err := client.Source.SingleTokenOnramp.Paused(nil)
	ccipshared.PanicErr(err)
	if !paused {
		return
	}
	fmt.Println("unpausing onramp...")
	tx, err := client.Source.SingleTokenOnramp.Unpause(client.Source.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Onramp unpaused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client ccipClient) unpauseOfframp() {
	paused, err := client.Dest.SingleTokenOfframp.Paused(nil)
	ccipshared.PanicErr(err)
	if !paused {
		return
	}
	fmt.Println("unpausing offramp...")
	tx, err := client.Dest.SingleTokenOfframp.Unpause(client.Dest.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Offramp unpaused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Dest.Client.Client, tx.Hash(), true)
}

func (client ccipClient) unpauseAll() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		client.unpauseOnramp()
	}()
	go func() {
		defer wg.Done()
		client.unpauseOfframp()
	}()
	go func() {
		defer wg.Done()
		client.unpauseOnrampPool()
	}()
	go func() {
		defer wg.Done()
		client.unpauseOfframpPool()
	}()
	wg.Wait()
}

func (client ccipClient) setConfig() {
	ccipConfig, err := ccip.OffchainConfig{
		SourceIncomingConfirmations: 0,
		DestIncomingConfirmations:   0,
	}.Encode()
	ccipshared.PanicErr(err)
	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper2.ContractSetConfigArgs(
		60*time.Second, // deltaProgress
		1*time.Second,  // deltaResend
		20*time.Second, // deltaRound
		2*time.Second,  // deltaGrace
		30*time.Second, // deltaStage
		3,
		[]int{1, 2, 3, 4}, // Transmission schedule: 1 oracle in first deltaStage, 2 in the second and so on.
		ccipshared.Oracles,
		ccipConfig,
		1*time.Second,
		10*time.Second,
		20*time.Second,
		10*time.Second,
		10*time.Second,
		1, // faults
		nil,
	)
	ccipshared.PanicErr(err)

	ctx := context.Background()

	tx, err := client.Dest.SingleTokenOfframp.SetConfig(
		client.Dest.Owner,
		signers,
		transmitters,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), true)
	fmt.Println("Config set on offramp. Tx hash:", tx.Hash().Hex())

	tx, err = client.Dest.MessageExecutor.SetConfig(
		client.Dest.Owner,
		signers,
		transmitters,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), true)
	fmt.Println("Config set on message executor. Tx hash:", tx.Hash().Hex())
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                 core/services/ccip/mocks/orm.go                                                                     000644  000765  000024  00000010174 14165346401 017475  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"
	ccip "github.com/smartcontractkit/chainlink/core/services/ccip"

	mock "github.com/stretchr/testify/mock"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// RelayReport provides a mock function with given fields: seqNum
func (_m *ORM) RelayReport(seqNum *big.Int) (ccip.RelayReport, error) {
	ret := _m.Called(seqNum)

	var r0 ccip.RelayReport
	if rf, ok := ret.Get(0).(func(*big.Int) ccip.RelayReport); ok {
		r0 = rf(seqNum)
	} else {
		r0 = ret.Get(0).(ccip.RelayReport)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(seqNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Requests provides a mock function with given fields: sourceChainId, destChainId, minSeqNum, maxSeqNum, status, executor, options
func (_m *ORM) Requests(sourceChainId *big.Int, destChainId *big.Int, minSeqNum *big.Int, maxSeqNum *big.Int, status ccip.RequestStatus, executor *common.Address, options []byte) ([]*ccip.Request, error) {
	ret := _m.Called(sourceChainId, destChainId, minSeqNum, maxSeqNum, status, executor, options)

	var r0 []*ccip.Request
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, *big.Int, *big.Int, ccip.RequestStatus, *common.Address, []byte) []*ccip.Request); ok {
		r0 = rf(sourceChainId, destChainId, minSeqNum, maxSeqNum, status, executor, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ccip.Request)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int, *big.Int, *big.Int, *big.Int, ccip.RequestStatus, *common.Address, []byte) error); ok {
		r1 = rf(sourceChainId, destChainId, minSeqNum, maxSeqNum, status, executor, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetExpiredRequests provides a mock function with given fields: sourceChainId, destChainId, expiryTimeoutSeconds, fromStatus, toStatus
func (_m *ORM) ResetExpiredRequests(sourceChainId *big.Int, destChainId *big.Int, expiryTimeoutSeconds int, fromStatus ccip.RequestStatus, toStatus ccip.RequestStatus) error {
	ret := _m.Called(sourceChainId, destChainId, expiryTimeoutSeconds, fromStatus, toStatus)

	var r0 error
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, int, ccip.RequestStatus, ccip.RequestStatus) error); ok {
		r0 = rf(sourceChainId, destChainId, expiryTimeoutSeconds, fromStatus, toStatus)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveRelayReport provides a mock function with given fields: report
func (_m *ORM) SaveRelayReport(report ccip.RelayReport) error {
	ret := _m.Called(report)

	var r0 error
	if rf, ok := ret.Get(0).(func(ccip.RelayReport) error); ok {
		r0 = rf(report)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveRequest provides a mock function with given fields: request
func (_m *ORM) SaveRequest(request *ccip.Request) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ccip.Request) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRequestSetStatus provides a mock function with given fields: sourceChainId, destChainId, seqNums, status
func (_m *ORM) UpdateRequestSetStatus(sourceChainId *big.Int, destChainId *big.Int, seqNums []*big.Int, status ccip.RequestStatus) error {
	ret := _m.Called(sourceChainId, destChainId, seqNums, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, []*big.Int, ccip.RequestStatus) error); ok {
		r0 = rf(sourceChainId, destChainId, seqNums, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRequestStatus provides a mock function with given fields: sourceChainId, destChainId, minSeqNum, maxSeqNum, status
func (_m *ORM) UpdateRequestStatus(sourceChainId *big.Int, destChainId *big.Int, minSeqNum *big.Int, maxSeqNum *big.Int, status ccip.RequestStatus) error {
	ret := _m.Called(sourceChainId, destChainId, minSeqNum, maxSeqNum, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, *big.Int, *big.Int, ccip.RequestStatus) error); ok {
		r0 = rf(sourceChainId, destChainId, minSeqNum, maxSeqNum, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
                                                                                                                                                                                                                                                                                                                                                                                                    core/services/ccip/mocks/lastreporter/off_ramp_last_reporter.go                                     000644  000765  000024  00000001752 14165346401 026166  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	mock "github.com/stretchr/testify/mock"

	single_token_offramp "github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
)

// OffRampLastReporter is an autogenerated mock type for the OffRampLastReporter type
type OffRampLastReporter struct {
	mock.Mock
}

// GetLastReport provides a mock function with given fields: opts
func (_m *OffRampLastReporter) GetLastReport(opts *bind.CallOpts) (single_token_offramp.CCIPRelayReport, error) {
	ret := _m.Called(opts)

	var r0 single_token_offramp.CCIPRelayReport
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) single_token_offramp.CCIPRelayReport); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(single_token_offramp.CCIPRelayReport)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
                      core/services/ccip/models.go                                                                        000644  000765  000024  00000005433 14165346401 017051  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink/core/utils"

	"github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"
)

type Request struct {
	SeqNum        utils.Big
	SourceChainID string // Note this will be some super set which includes evm_chain_id
	DestChainID   string // Note this will be some super set which includes evm_chain_id
	Sender        common.Address
	Receiver      common.Address
	Data          []byte
	Tokens        pq.StringArray
	Amounts       pq.StringArray
	Executor      common.Address
	Options       []byte
	Raw           []byte // Full ABI-encoded event for merkle tree
	Status        RequestStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (Request) TableName() string {
	return "ccip_requests"
}

func MakeOptions() abi.Arguments {
	mustType := func(ts string) abi.Type {
		ty, _ := abi.NewType(ts, "", nil)
		return ty
	}
	return []abi.Argument{
		{
			Type: mustType("bool"),
			Name: "oracleExecute",
		},
	}
}

func mustStringToBigInt(s string) *big.Int {
	i, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		panic(fmt.Sprintf("invalid big.Int string %v", s))
	}
	return i
}

func (r Request) ToMessage() Message {
	var tokens []common.Address
	for _, t := range r.Tokens {
		tokens = append(tokens, common.HexToAddress(t))
	}
	var amounts []*big.Int
	for _, a := range r.Amounts {
		amounts = append(amounts, mustStringToBigInt(a))
	}
	return Message{
		SequenceNumber:     r.SeqNum.ToInt(),
		SourceChainId:      mustStringToBigInt(r.SourceChainID),
		DestinationChainId: mustStringToBigInt(r.DestChainID),
		Sender:             r.Sender,
		Payload: struct {
			Receiver common.Address   `json:"receiver"`
			Data     []uint8          `json:"data"`
			Tokens   []common.Address `json:"tokens"`
			Amounts  []*big.Int       `json:"amounts"`
			Executor common.Address   `json:"executor"`
			Options  []uint8          `json:"options"`
		}{
			Receiver: r.Receiver,
			Data:     r.Data,
			Tokens:   tokens,
			Amounts:  amounts,
			Executor: r.Executor,
			Options:  r.Options,
		},
	}
}

type RequestStatus string

const (
	RequestStatusUnstarted    RequestStatus = "unstarted"
	RequestStatusRelayPending RequestStatus = "relay_pending"
	// We only mark relay confirmed after we've seen the report accepted log with sufficient
	// number of confirmations
	RequestStatusRelayConfirmed   RequestStatus = "relay_confirmed"
	RequestStatusExecutionPending RequestStatus = "execution_pending"
	// We only mark execution confirmed after we've seen the Message executed log with sufficient
	// number of confirmations
	RequestStatusExecutionConfirmed RequestStatus = "execution_confirmed"
)

type RelayReport struct {
	Root      []byte
	MinSeqNum utils.Big
	MaxSeqNum utils.Big
	CreatedAt time.Time
}
                                                                                                                                                                                                                                     core/services/ccip/config.go                                                                        000644  000765  000024  00000010521 14165346401 017025  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"math/big"
	"time"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/chains/evm"
	evmconfig "github.com/smartcontractkit/chainlink/core/chains/evm/config"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ethkey"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/p2pkey"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/core/store/models"

	ocrcommontypes "github.com/smartcontractkit/libocr/commontypes"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"
)

// Fallback to config if explicit spec parameters are not set
func computeLocalConfig(config evmconfig.OCR2Config, dev bool, bt time.Duration, confs uint16, poll time.Duration) ocrtypes.LocalConfig {
	var blockchainTimeout time.Duration
	if bt != 0 {
		blockchainTimeout = bt
	} else {
		blockchainTimeout = config.OCR2BlockchainTimeout()
	}

	var contractConfirmations uint16
	if confs != 0 {
		contractConfirmations = confs
	} else {
		contractConfirmations = config.OCR2ContractConfirmations()
	}

	var contractConfigTrackerPollInterval time.Duration
	if poll != 0 {
		contractConfigTrackerPollInterval = poll
	} else {
		contractConfigTrackerPollInterval = config.OCR2ContractPollInterval()
	}

	lc := ocrtypes.LocalConfig{
		BlockchainTimeout:                  blockchainTimeout,
		ContractConfigConfirmations:        contractConfirmations,
		ContractConfigTrackerPollInterval:  contractConfigTrackerPollInterval,
		ContractTransmitterTransmitTimeout: config.OCR2ContractTransmitterTransmitTimeout(),
		DatabaseTimeout:                    config.OCR2DatabaseTimeout(),
	}
	if dev {
		// Skips config validation so we can use any config parameters we want.
		// For example to lower contractConfigTrackerPollInterval to speed up tests.
		lc.DevelopmentMode = ocrtypes.EnableDangerousDevelopmentMode
	}
	return lc
}

func parseBootstrapPeers(peers []string) (bootstrapPeers []ocrcommontypes.BootstrapperLocator, err error) {
	for _, bs := range peers {
		var bsl ocrcommontypes.BootstrapperLocator
		err = bsl.UnmarshalText([]byte(bs))
		if err != nil {
			return nil, err
		}
		bootstrapPeers = append(bootstrapPeers, bsl)
	}
	return
}

func getValidatedBootstrapPeers(specPeers []string, chain evm.Chain) ([]ocrcommontypes.BootstrapperLocator, error) {
	bootstrapPeers, err := parseBootstrapPeers(specPeers)
	if err != nil {
		return nil, err
	}
	if len(bootstrapPeers) == 0 {
		bootstrapPeers = chain.Config().P2PV2Bootstrappers()
		if err != nil {
			return nil, err
		}
	}
	return bootstrapPeers, nil
}

func validatePeerWrapper(specID *p2pkey.PeerID, chain evm.Chain, pw *ocrcommon.SingletonPeerWrapper) error {
	var peerID p2pkey.PeerID
	if specID != nil {
		peerID = *specID
	} else {
		peerID = chain.Config().P2PPeerID()
	}
	if !pw.IsStarted() {
		return errors.New("peerWrapper is not started. OCR2 jobs require a started and running peer. Did you forget to specify P2P_LISTEN_PORT?")
	} else if pw.PeerID != peerID {
		return errors.Errorf("given peer with ID '%s' does not match OCR2 configured peer with ID: %s", pw.PeerID.String(), peerID.String())
	}
	return nil
}

func getValidatedKeyBundle(specBundleID *models.Sha256Hash, chain evm.Chain, ks keystore.OCR2) (kb ocr2key.KeyBundle, err error) {
	var kbs string
	if specBundleID != nil {
		kbs = specBundleID.String()
	} else if kbs, err = chain.Config().OCR2KeyBundleID(); err != nil {
		return kb, err
	}
	key, err := ks.Get(kbs)
	if err != nil {
		return kb, err
	}
	return key, nil
}

func getTransmitterAddress(specAddress *ethkey.EIP55Address, chain evm.Chain) (ta ethkey.EIP55Address, err error) {
	if specAddress != nil {
		ta = *specAddress
	} else if ta, err = chain.Config().OCR2TransmitterAddress(); err != nil {
		return ta, err
	}
	return ta, nil
}

// Multi-chain tests using the sim have to be remapped to the default
// sim chainID because its a hardcoded constant in the geth code base and so
// and CHAINID op codes will ALWAYS be 1337.
func maybeRemapChainID(chainID *big.Int) *big.Int {
	testChainIDs := []*big.Int{big.NewInt(1000), big.NewInt(2000)}
	for _, testChainID := range testChainIDs {
		if chainID.Cmp(testChainID) == 0 {
			return big.NewInt(1337)
		}
	}
	return chainID
}
                                                                                                                                                                               core/services/ccip/contract_tracker.go                                                              000644  000765  000024  00000036413 14165346401 021120  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // Note lifted from https://github.com/smartcontractkit/chainlink/pull/4809/files
// TODO: Pull into common ocr library
package ccip

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink/core/chains"
	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/eth"
	httypes "github.com/smartcontractkit/chainlink/core/services/headtracker/types"
	"github.com/smartcontractkit/chainlink/core/services/log"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/core/services/postgres"
	"github.com/smartcontractkit/chainlink/core/utils"
	"github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

var (
	_ ocrtypes.ContractConfigTracker = &CCIPContractTracker{}
	_ httypes.HeadTrackable          = &CCIPContractTracker{}

	OCRContractConfigSet = getEventTopic("ConfigSet")
)

type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}

type ConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []gethCommon.Address
	Transmitters              []gethCommon.Address
	F                         uint8
	OnchainConfig             []byte
	EncodedConfigVersion      uint64
	Encoded                   []byte
	Raw                       gethTypes.Log
}

type OffchainConfig struct {
	SourceIncomingConfirmations uint32
	DestIncomingConfirmations   uint32
}

type OCR2 interface {
	Address() gethCommon.Address
	ParseLog(log gethTypes.Log) (generated.AbigenLog, error)
	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails, error)
	ParseConfigSet(log gethTypes.Log) (ConfigSet, error)
}

type offrampTracker struct {
	*single_token_offramp.SingleTokenOffRamp
}

func (ot offrampTracker) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails, error) {
	dets, err := ot.SingleTokenOffRamp.LatestConfigDetails(opts)
	return LatestConfigDetails(dets), err
}

func (ot offrampTracker) ParseConfigSet(log gethTypes.Log) (ConfigSet, error) {
	c, err := ot.SingleTokenOffRamp.ParseConfigSet(log)
	return ConfigSet(*c), err
}

type executorTracker struct {
	*message_executor.MessageExecutor
}

func (ot executorTracker) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails, error) {
	dets, err := ot.MessageExecutor.LatestConfigDetails(opts)
	return LatestConfigDetails(dets), err
}

func (ot executorTracker) ParseConfigSet(log gethTypes.Log) (ConfigSet, error) {
	c, err := ot.MessageExecutor.ParseConfigSet(log)
	return ConfigSet(*c), err
}

// CCIPContractTracker complies with ContractConfigTracker interface and
// handles log events related to the contract more generally
type CCIPContractTracker struct {
	utils.StartStopOnce

	ethClient       eth.Client
	contract        OCR2
	logBroadcaster  log.Broadcaster
	jobID           int32
	logger          logger.Logger
	gdb             *gorm.DB
	blockTranslator ocrcommon.BlockTranslator
	chain           evm.Chain

	// HeadBroadcaster
	headBroadcaster  httypes.HeadBroadcaster
	unsubscribeHeads func()

	// Start/Stop lifecycle
	ctx             context.Context
	ctxCancel       context.CancelFunc
	wg              sync.WaitGroup
	unsubscribeLogs func()

	// ContractConfig
	configsMB utils.Mailbox
	chConfigs chan ocrtypes.ContractConfig

	// LatestBlockHeight
	latestBlockHeight   int64
	latestBlockHeightMu sync.RWMutex
}

// NewCCIPContractTracker makes a new CCIPContractTracker
func NewCCIPContractTracker(
	contract OCR2,
	ethClient eth.Client,
	logBroadcaster log.Broadcaster,
	jobID int32,
	logger logger.Logger,
	gdb *gorm.DB,
	chain evm.Chain,
	headBroadcaster httypes.HeadBroadcaster,
) (o *CCIPContractTracker) {
	ctx, cancel := context.WithCancel(context.Background())
	return &CCIPContractTracker{
		utils.StartStopOnce{},
		ethClient,
		contract,
		logBroadcaster,
		jobID,
		logger,
		gdb,
		ocrcommon.NewBlockTranslator(chain.Config(), ethClient, logger),
		chain,
		headBroadcaster,
		nil,
		ctx,
		cancel,
		sync.WaitGroup{},
		nil,
		// Should only ever be 0 or 1 config in the mailbox, put a sanity bound of 100.
		*utils.NewMailbox(100),
		make(chan ocrtypes.ContractConfig),
		-1,
		sync.RWMutex{},
	}
}

// Start must be called before logs can be delivered
// It ought to be called before starting OCR
func (t *CCIPContractTracker) Start() error {
	return t.StartOnce("CCIPContractTracker", func() (err error) {
		t.logger.Infow("CCIPContractTracker: registering for config set logs")
		t.unsubscribeLogs = t.logBroadcaster.Register(t, log.ListenerOpts{
			Contract: t.contract.Address(),
			ParseLog: t.contract.ParseLog,
			LogsWithTopics: map[gethCommon.Hash][][]log.Topic{
				single_token_offramp.SingleTokenOffRampConfigSet{}.Topic(): nil,
			},
			NumConfirmations: 1,
		})

		var latestHead *eth.Head
		latestHead, t.unsubscribeHeads = t.headBroadcaster.Subscribe(t)
		if latestHead != nil {
			t.setLatestBlockHeight(*latestHead)
		}

		t.wg.Add(1)
		go t.processLogs()
		return nil
	})
}

// Close should be called after teardown of the OCR job relying on this tracker
func (t *CCIPContractTracker) Close() error {
	return t.StopOnce("CCIPContractTracker", func() error {
		t.ctxCancel()
		t.wg.Wait()
		t.unsubscribeHeads()
		t.unsubscribeLogs()
		close(t.chConfigs)
		return nil
	})
}

// Connect conforms to HeadTrackable
func (t *CCIPContractTracker) Connect(*eth.Head) error { return nil }

// OnNewLongestChain conformed to HeadTrackable and updates latestBlockHeight
func (t *CCIPContractTracker) OnNewLongestChain(_ context.Context, h eth.Head) {
	t.setLatestBlockHeight(h)
}

func (t *CCIPContractTracker) setLatestBlockHeight(h eth.Head) {
	var num int64
	if h.L1BlockNumber.Valid {
		num = h.L1BlockNumber.Int64
	} else {
		num = h.Number
	}
	t.latestBlockHeightMu.Lock()
	defer t.latestBlockHeightMu.Unlock()
	if num > t.latestBlockHeight {
		t.latestBlockHeight = num
	}
}

func (t *CCIPContractTracker) getLatestBlockHeight() int64 {
	t.latestBlockHeightMu.RLock()
	defer t.latestBlockHeightMu.RUnlock()
	return t.latestBlockHeight
}

func (t *CCIPContractTracker) processLogs() {
	defer t.wg.Done()
	for {
		select {
		case <-t.configsMB.Notify():
			// NOTE: libocr could take an arbitrary amount of time to process a
			// new config. To avoid blocking the log broadcaster, we use this
			// background thread to deliver them and a mailbox as the buffer.
			for {
				x, exists := t.configsMB.Retrieve()
				if !exists {
					break
				}
				cc, ok := x.(ocrtypes.ContractConfig)
				if !ok {
					panic(fmt.Sprintf("expected ocrtypes.ContractConfig but got %T", x))
				}
				select {
				case t.chConfigs <- cc:
				case <-t.ctx.Done():
					return
				}
			}
		case <-t.ctx.Done():
			return
		}
	}
}

// HandleLog complies with LogListener interface
// It is not thread safe
func (t *CCIPContractTracker) HandleLog(lb log.Broadcast) {
	t.logger.Infow("CCIPContractTracker: config set log received", "log", lb.String())
	was, err := t.logBroadcaster.WasAlreadyConsumed(t.gdb, lb)
	if err != nil {
		t.logger.Errorw("OCRContract: could not determine if log was already consumed", "error", err)
		return
	} else if was {
		return
	}

	raw := lb.RawLog()
	if raw.Address != t.contract.Address() {
		t.logger.Errorf("log address of 0x%x does not match configured contract address of 0x%x", raw.Address, t.contract.Address())
		t.logger.ErrorIfCalling(func() error { return t.logBroadcaster.MarkConsumed(t.gdb, lb) })
		return
	}
	topics := raw.Topics
	if len(topics) == 0 {
		t.logger.ErrorIfCalling(func() error { return t.logBroadcaster.MarkConsumed(t.gdb, lb) })
		return
	}

	var consumed bool
	switch topics[0] {
	case single_token_offramp.SingleTokenOffRampConfigSet{}.Topic():
		configSet, err := t.contract.ParseConfigSet(raw)
		if err != nil {
			t.logger.Errorw("could not parse config set", "err", err)
			t.logger.ErrorIfCalling(func() error { return t.logBroadcaster.MarkConsumed(t.gdb, lb) })
			return
		}
		configSet.Raw = raw
		cc := ContractConfigFromConfigSetEvent(configSet)

		wasOverCapacity := t.configsMB.Deliver(cc)
		if wasOverCapacity {
			t.logger.Error("config mailbox is over capacity - dropped the oldest unprocessed item")
		}
	default:
		logger.Debugw("CCIPContractTracker: got unrecognised log topic", "topic", topics[0])
	}
	if !consumed {
		ctx, cancel := postgres.DefaultQueryCtx()
		defer cancel()
		t.logger.ErrorIfCalling(func() error { return t.logBroadcaster.MarkConsumed(t.gdb.WithContext(ctx), lb) })
	}
}

// IsLaterThan returns true if the first log was emitted "after" the second log
// from the blockchain's point of view
func IsLaterThan(incoming gethTypes.Log, existing gethTypes.Log) bool {
	return incoming.BlockNumber > existing.BlockNumber ||
		(incoming.BlockNumber == existing.BlockNumber && incoming.TxIndex > existing.TxIndex) ||
		(incoming.BlockNumber == existing.BlockNumber && incoming.TxIndex == existing.TxIndex && incoming.Index > existing.Index)
}

// IsV2Job complies with LogListener interface
func (t *CCIPContractTracker) IsV2Job() bool {
	return true
}

// JobID complies with LogListener interface
func (t *CCIPContractTracker) JobID() int32 {
	return t.jobID
}

// Notify returns a channel that can wake up the contract tracker to let it
// know when a new config is available
func (t *CCIPContractTracker) Notify() <-chan struct{} {
	return nil
}

// LatestConfigDetails queries the eth node
func (t *CCIPContractTracker) LatestConfigDetails(ctx context.Context) (changedInBlock uint64, configDigest ocrtypes.ConfigDigest, err error) {
	var cancel context.CancelFunc
	ctx, cancel = utils.CombinedContext(t.ctx, ctx)
	defer cancel()

	opts := bind.CallOpts{Context: ctx, Pending: false}
	result, err := t.contract.LatestConfigDetails(&opts)
	if err != nil {
		return 0, configDigest, errors.Wrap(err, "error getting LatestConfigDetails")
	}

	t.logger.Infow("CCIPContractTracker: latest config details", "digest", hexutil.Encode(result.ConfigDigest[:]))
	configDigest, err = ocrtypes.BytesToConfigDigest(result.ConfigDigest[:])
	if err != nil {
		return 0, configDigest, errors.Wrap(err, fmt.Sprintf("error getting LatestConfigDetails %v", t.contract.Address()))
	}
	return uint64(result.BlockNumber), configDigest, err
}

// Return the latest configuration
func (t *CCIPContractTracker) LatestConfig(ctx context.Context, changedInBlock uint64) (ocrtypes.ContractConfig, error) {
	fromBlock, toBlock := t.blockTranslator.NumberToQueryRange(ctx, changedInBlock)
	q := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []gethCommon.Address{t.contract.Address()},
		Topics: [][]gethCommon.Hash{
			{OCRContractConfigSet},
		},
	}
	var cancel context.CancelFunc
	ctx, cancel = utils.CombinedContext(t.ctx, ctx)
	defer cancel()

	logs, err := t.ethClient.FilterLogs(ctx, q)
	if err != nil {
		return ocrtypes.ContractConfig{}, err
	}
	if len(logs) == 0 {
		return ocrtypes.ContractConfig{}, errors.Errorf("ConfigFromLogs: OCRContract with address 0x%x has no logs", t.contract.Address())
	}

	latest, err := t.contract.ParseConfigSet(logs[len(logs)-1])
	if err != nil {
		return ocrtypes.ContractConfig{}, errors.Wrap(err, "ConfigFromLogs failed to ParseConfigSet")
	}
	latest.Raw = logs[len(logs)-1]
	if latest.Raw.Address != t.contract.Address() {
		return ocrtypes.ContractConfig{}, errors.Errorf("log address of 0x%x does not match configured contract address of 0x%x", latest.Raw.Address, t.contract.Address())
	}

	cc := ContractConfigFromConfigSetEvent(latest)
	t.logger.Infow("CCIPContractTracker: latest config", "digest", hexutil.Encode(cc.ConfigDigest[:]))
	return cc, err
}

func ContractConfigFromConfigSetEvent(changed ConfigSet) ocrtypes.ContractConfig {
	var transmitAccounts []ocrtypes.Account
	for _, addr := range changed.Transmitters {
		transmitAccounts = append(transmitAccounts, ocrtypes.Account(addr.Hex()))
	}
	var signers []ocrtypes.OnchainPublicKey
	for _, addr := range changed.Signers {
		addr := addr
		signers = append(signers, addr[:])
	}
	return ocrtypes.ContractConfig{
		ConfigDigest:          changed.ConfigDigest,
		ConfigCount:           changed.ConfigCount,
		Signers:               signers,
		Transmitters:          transmitAccounts,
		F:                     changed.F,
		OnchainConfig:         changed.OnchainConfig,
		OffchainConfigVersion: changed.EncodedConfigVersion,
		OffchainConfig:        changed.Encoded,
	}
}

// LatestBlockHeight queries the eth node for the most recent header
func (t *CCIPContractTracker) LatestBlockHeight(ctx context.Context) (blockheight uint64, err error) {
	// We skip confirmation checking anyway on Optimism so there's no need to
	// care about the block height; we have no way of getting the L1 block
	// height anyway
	if t.chain.Config().ChainType() == chains.Optimism {
		return 0, nil
	}
	latestBlockHeight := t.getLatestBlockHeight()
	if latestBlockHeight >= 0 {
		return uint64(latestBlockHeight), nil
	}

	t.logger.Debugw("CCIPContractTracker: still waiting for first head, falling back to on-chain lookup")

	var cancel context.CancelFunc
	ctx, cancel = utils.CombinedContext(t.ctx, ctx)
	defer cancel()

	h, err := t.ethClient.HeadByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	if h == nil {
		return 0, errors.New("got nil head")
	}

	if h.L1BlockNumber.Valid {
		return uint64(h.L1BlockNumber.Int64), nil
	}

	return uint64(h.Number), nil
}

func (t *CCIPContractTracker) GetOffchainConfig() (OffchainConfig, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	changedInBlock, _, err := t.LatestConfigDetails(ctx)
	if err != nil {
		return OffchainConfig{}, errors.Wrap(err, "could not get block number for latest config change")
	}
	config, err := t.LatestConfig(ctx, changedInBlock)
	if err != nil {
		return OffchainConfig{}, errors.Wrap(err, "could not get latest config")
	}
	publicConfig, err := confighelper.PublicConfigFromContractConfig(false, config)
	if err != nil {
		return OffchainConfig{}, errors.Wrap(err, "could not parse latest config")
	}
	ccipConfig, err := Decode(publicConfig.ReportingPluginConfig)
	if err != nil {
		return OffchainConfig{}, errors.Wrap(err, "could not decode latest config")
	}
	return ccipConfig, nil
}

func Decode(encodedConfig []byte) (OffchainConfig, error) {
	var result OffchainConfig
	err := json.Unmarshal(encodedConfig, &result)
	return result, err
}

func (occ OffchainConfig) Encode() ([]byte, error) {
	return json.Marshal(occ)
}

func getEventTopic(name string) gethCommon.Hash {
	abi, err := abi.JSON(strings.NewReader(single_token_offramp.SingleTokenOffRampABI))
	if err != nil {
		panic("could not parse singletoken ABI: " + err.Error())
	}
	event, exists := abi.Events[name]
	if !exists {
		panic(fmt.Sprintf("abi.Events was missing %s", name))
	}
	return event.ID
}
                                                                                                                                                                                                                                                     core/services/ccip/orm_test.go                                                                      000644  000765  000024  00000012112 14165346401 017412  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip_test

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/core/services/ccip"
	"github.com/smartcontractkit/chainlink/core/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestORM(t *testing.T) {
	// Use a real db so we can do timestamp testing.
	_, db, _ := heavyweight.FullTestDB(t, "orm_test", true, false)
	orm := ccip.NewORM(db)
	source := big.NewInt(1)
	dest := big.NewInt(2)

	// Check we can read/write requests.
	req := ccip.Request{
		SourceChainID: source.String(),
		DestChainID:   dest.String(),
		SeqNum:        *utils.NewBigI(10),
		Sender:        common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
		Receiver:      common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
		Data:          []byte("hello"),
		Tokens:        pq.StringArray{},
		Amounts:       pq.StringArray{},
		Options:       []byte{},
	}
	err := orm.SaveRequest(&req)
	require.NoError(t, err)
	reqRead, err := orm.Requests(source, dest, req.SeqNum.ToInt(), req.SeqNum.ToInt(), "", nil, nil)
	require.NoError(t, err)
	require.Equal(t, 1, len(reqRead))
	assert.True(t, reqRead[0].UpdatedAt != time.Time{})
	assert.True(t, reqRead[0].CreatedAt != time.Time{})
	assert.Equal(t, req.Data, reqRead[0].Data)

	// Check we can update the request status.
	err = orm.UpdateRequestStatus(source, dest, req.SeqNum.ToInt(), req.SeqNum.ToInt(), ccip.RequestStatusRelayPending)
	require.NoError(t, err)
	// Updating an non-existent reqID should error
	err = orm.UpdateRequestStatus(source, dest, big.NewInt(1337), big.NewInt(1337), ccip.RequestStatusUnstarted)
	require.Error(t, err)
	reqReadAfterUpdate, err := orm.Requests(source, dest, req.SeqNum.ToInt(), req.SeqNum.ToInt(), "", nil, nil)
	require.NoError(t, err)
	require.Equal(t, 1, len(reqReadAfterUpdate))
	assert.Equal(t, ccip.RequestStatusRelayPending, reqReadAfterUpdate[0].Status)
	assert.True(t, reqReadAfterUpdate[0].UpdatedAt.After(reqRead[0].UpdatedAt), fmt.Sprintf("before %v after %v", reqRead[0].UpdatedAt, reqReadAfterUpdate[0].UpdatedAt))
	assert.Equal(t, reqReadAfterUpdate[0].CreatedAt, reqRead[0].CreatedAt)

	// Check we can read/write relay reports.
	var aroot = [32]byte{0x01}
	require.NoError(t, orm.SaveRelayReport(ccip.RelayReport{
		Root:      aroot[:],
		MinSeqNum: *utils.NewBig(big.NewInt(1)),
		MaxSeqNum: *utils.NewBig(big.NewInt(2)),
	}))
	r, err := orm.RelayReport(big.NewInt(1))
	require.NoError(t, err)
	assert.Equal(t, byte(0x01), r.Root[0])
	require.NoError(t, err)

	// Check we can filter by status and executor with multiple requests present.
	executor := common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB5")
	reqForOracleExecution := ccip.Request{
		SourceChainID: source.String(),
		DestChainID:   dest.String(),
		SeqNum:        *utils.NewBigI(11),
		Sender:        common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
		Receiver:      common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
		Data:          []byte("hello"),
		Tokens:        pq.StringArray{},
		Amounts:       pq.StringArray{},
		Executor:      executor,
		Options:       []byte{},
	}
	require.NoError(t, orm.SaveRequest(&reqForOracleExecution))
	require.NoError(t, orm.UpdateRequestStatus(source, dest, big.NewInt(11), big.NewInt(11), ccip.RequestStatusRelayConfirmed))
	reqsForOracle, err := orm.Requests(source, dest, nil, nil, ccip.RequestStatusRelayConfirmed, nil, nil)
	require.NoError(t, err)
	require.Len(t, reqsForOracle, 1)
	reqsForOracle, err = orm.Requests(source, dest, nil, nil, ccip.RequestStatusRelayConfirmed, &executor, nil)
	require.NoError(t, err)
	require.Len(t, reqsForOracle, 1)

	// Check we can update the status with specific seq nums, as opposed to a range.
	reqsBefore, err := orm.Requests(source, dest, big.NewInt(10), big.NewInt(11), ccip.RequestStatusRelayConfirmed, nil, nil)
	require.NoError(t, err)
	require.NoError(t, orm.UpdateRequestSetStatus(source, dest, []*big.Int{big.NewInt(10), big.NewInt(11)}, ccip.RequestStatusExecutionConfirmed))
	reqs, err := orm.Requests(source, dest, nil, nil, ccip.RequestStatusExecutionConfirmed, nil, nil)
	require.NoError(t, err)
	require.Len(t, reqs, 2)
	assert.True(t, reqs[0].UpdatedAt.After(reqsBefore[0].UpdatedAt), fmt.Sprintf("before %v after %v", reqRead[0].UpdatedAt, reqReadAfterUpdate[0].UpdatedAt))

	// Check that we can reset the status of expired requests.
	res, err := db.Exec(`UPDATE ccip_requests SET updated_at = $1`, time.Now().Add(-2*time.Second))
	require.NoError(t, err)
	n, err := res.RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(2), n)
	// Now they should be recognized as being 1s old, so we can reset them with a timeout of 1s.
	require.NoError(t, orm.ResetExpiredRequests(source, dest, 1, ccip.RequestStatusExecutionConfirmed, ccip.RequestStatusRelayConfirmed))
	// Should all be relay confirmed now.
	reqs, err = orm.Requests(source, dest, nil, nil, ccip.RequestStatusRelayConfirmed, nil, nil)
	require.NoError(t, err)
	require.Len(t, reqs, 2)
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                      core/services/ccip/validate_test.go                                                                 000644  000765  000024  00000004515 14165346401 020416  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"testing"

	"github.com/smartcontractkit/chainlink/core/internal/testutils/configtest"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateCCIPSpec(t *testing.T) {
	var tt = []struct {
		name         string
		toml         string
		setGlobalCfg func(t *testing.T, c *configtest.TestGeneralConfig)
		assertion    func(t *testing.T, os job.Job, err error)
	}{
		{
			name: "decodes valid ccip spec toml",
			toml: `
type               = "ccip-relay"
schemaVersion      = 1
contractAddress    = "0x613a38AC1659769640aaE063C651F48E0250454C"
p2pPeerID          = "12D3KooWHfYFQ8hGttAYbMCevQVESEQhzJAqFZokMVtom8bNxwGq"
p2pBootstrapPeers  = [
"/dns4/chain.link/tcp/1234/p2p/16Uiu2HAm58SP7UL8zsnpeuwHfytLocaqgnyaYKP8wu7qRdrixLju",
]
keyBundleID        = "73e8966a78ca09bb912e9565cfb79fbe8a6048fab1f0cf49b18047c3895e0447"
monitoringEndpoint = "chain.link:4321"
transmitterAddress = "0xF67D0290337bca0847005C7ffD1BC75BA9AAE6e4"
observationTimeout = "10s"
observationSource = """
ds1          [type=bridge name=voter_turnout];
ds1_parse    [type=jsonparse path="one,two"];
ds1_multiply [type=multiply times=1.23];
ds1 -> ds1_parse -> ds1_multiply -> answer1;
answer1      [type=median Index=0];
"""
`,
			assertion: func(t *testing.T, os job.Job, err error) {
				require.NoError(t, err)
				assert.Equal(t, 1, int(os.SchemaVersion))
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s, err := ValidatedCCIPSpec(tc.toml)
			tc.assertion(t, s, err)
		})
	}
}

func TestValidateCCIPBootstrapSpec(t *testing.T) {
	var tt = []struct {
		name         string
		toml         string
		setGlobalCfg func(t *testing.T, c *configtest.TestGeneralConfig)
		assertion    func(t *testing.T, os job.Job, err error)
	}{
		{
			name: "decodes valid ccip bootstrap spec toml",
			toml: `
type               = "ccip-bootstrap"
schemaVersion      = 1
contractAddress    = "0x613a38AC1659769640aaE063C651F48E0250454C"
monitoringEndpoint = "chain.link:4321"
`,
			assertion: func(t *testing.T, os job.Job, err error) {
				require.NoError(t, err)
				assert.Equal(t, 1, int(os.SchemaVersion))
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s, err := ValidatedCCIPBootstrapSpec(tc.toml)
			tc.assertion(t, s, err)
		})
	}
}
                                                                                                                                                                                   core/services/ccip/orm.go                                                                           000644  000765  000024  00000014475 14165346401 016371  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/smartcontractkit/chainlink/core/services/postgres"

	"github.com/smartcontractkit/sqlx"
)

// ORM We intend to use the same table for all xchain requests.
// TODO: I think we may need to pass in string based chainIDs
// in the future when we support non-evm chains, for now keep the interface EVM
// The triplet (seqNum, source, dest) defined the Message.
//go:generate mockery --name ORM --output ./mocks/ --case=underscore
type ORM interface {
	// Note always returns them sorted by seqNum
	Requests(sourceChainId, destChainId, minSeqNum, maxSeqNum *big.Int, status RequestStatus, executor *common.Address, options []byte) ([]*Request, error)
	SaveRequest(request *Request) error
	UpdateRequestStatus(sourceChainId, destChainId, minSeqNum, maxSeqNum *big.Int, status RequestStatus) error
	UpdateRequestSetStatus(sourceChainId, destChainId *big.Int, seqNums []*big.Int, status RequestStatus) error
	ResetExpiredRequests(sourceChainId, destChainId *big.Int, expiryTimeoutSeconds int, fromStatus RequestStatus, toStatus RequestStatus) error
	RelayReport(seqNum *big.Int) (RelayReport, error)
	SaveRelayReport(report RelayReport) error
}

type orm struct {
	db *sqlx.DB
}

var _ORM = (*orm)(nil)

func NewORM(db *sqlx.DB) ORM {
	return &orm{db}
}

// Note that executor can be left unset in the request, meaning anyone can execute.
// A nil executor as an argument here however means "don't filter on executor" and so it will return requests with both unset and set executors.
func (o *orm) Requests(sourceChainId, destChainId *big.Int, minSeqNum, maxSeqNum *big.Int, status RequestStatus, executor *common.Address, options []byte) (reqs []*Request, err error) {
	q := `SELECT * FROM ccip_requests WHERE true`
	if sourceChainId != nil {
		q += fmt.Sprintf(" AND source_chain_id = '%s'", sourceChainId.String())
	}
	if destChainId != nil {
		q += fmt.Sprintf(" AND dest_chain_id = '%s'", destChainId.String())
	}
	if minSeqNum != nil {
		q += fmt.Sprintf(" AND seq_num >= CAST(%s AS NUMERIC(78,0))", minSeqNum.String())
	}
	if maxSeqNum != nil {
		q += fmt.Sprintf(" AND seq_num <= CAST(%s AS NUMERIC(78,0))", maxSeqNum.String())
	}
	if status != "" {
		q += fmt.Sprintf(" AND status = '%s'", status)
	}
	if executor != nil {
		q += fmt.Sprintf(` AND executor = '\x%v'`, executor.String()[2:])
	}
	if options != nil {
		q += fmt.Sprintf(` AND options = '\x%v'`, hexutil.Encode(options)[2:])
	}
	q += ` ORDER BY seq_num ASC`
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	err = o.db.SelectContext(ctx, &reqs, q)
	return
}

func (o *orm) UpdateRequestStatus(sourceChainId, destChainId, minSeqNum, maxSeqNum *big.Int, status RequestStatus) error {
	// We return seqNum here to error if it doesn't exist
	q := `UPDATE ccip_requests SET status = $1, updated_at = now()
		WHERE seq_num >= CAST($2 AS NUMERIC(78,0))
		  AND seq_num <= CAST($3 AS NUMERIC(78,0))
		  AND source_chain_id = $4 
		  AND dest_chain_id = $5 
		RETURNING seq_num`
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	res, err := o.db.ExecContext(ctx, q, status, minSeqNum.String(), maxSeqNum.String(), sourceChainId.String(), destChainId.String())
	if err != nil {
		return err
	}
	seqRange := big.NewInt(0).Sub(maxSeqNum, minSeqNum)
	expectedUpdates := seqRange.Add(seqRange, big.NewInt(1)).Int64()
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n != expectedUpdates {
		return fmt.Errorf("did not update expected num rows, got %v want %v", n, expectedUpdates)
	}
	return nil
}

func (o *orm) UpdateRequestSetStatus(sourceChainId, destChainId *big.Int, seqNums []*big.Int, status RequestStatus) error {
	if len(seqNums) == 0 {
		return nil
	}
	seqNumsSet := fmt.Sprintf(`(CAST('%s' AS NUMERIC(78,0))`, seqNums[0].String())
	for _, n := range seqNums[1:] {
		seqNumsSet += fmt.Sprintf(`,CAST('%s' AS NUMERIC(78,0))`, n.String())
	}
	seqNumsSet += `)`
	q := fmt.Sprintf(`UPDATE ccip_requests SET status = $1, updated_at = now()
		WHERE seq_num IN %s 
		  AND source_chain_id = $2 
		  AND dest_chain_id = $3 
		RETURNING seq_num`, seqNumsSet)
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	res, err := o.db.ExecContext(ctx, q, status, sourceChainId.String(), destChainId.String())
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if int(n) != len(seqNums) {
		return fmt.Errorf("did not update expected num rows, got %v want %v", n, len(seqNums))
	}
	return err
}

func (o *orm) ResetExpiredRequests(sourceChainId, destChainId *big.Int, expiryTimeoutSeconds int, fromStatus RequestStatus, toStatus RequestStatus) error {
	q := fmt.Sprintf(`UPDATE ccip_requests SET status = $1, updated_at = now()
		WHERE now() > (updated_at + interval '%d seconds') 
			AND source_chain_id = $2
			AND dest_chain_id = $3
			AND status = $4`, expiryTimeoutSeconds)
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	_, err := o.db.ExecContext(ctx, q, toStatus, sourceChainId.String(), destChainId.String(), fromStatus)
	return err
}

// Note requests will only be added in an unstarted status
func (o *orm) SaveRequest(request *Request) error {
	q := `INSERT INTO ccip_requests 
    (seq_num, source_chain_id, dest_chain_id, sender, receiver, data, tokens, amounts, executor, options, raw, status, created_at, updated_at) 
    VALUES (:seq_num, :source_chain_id, :dest_chain_id, :sender, :receiver, :data, :tokens, :amounts, :executor, :options, :raw, 'unstarted', now(), now())
   	ON CONFLICT DO NOTHING `
	stmt, err := o.db.PrepareNamed(q)
	if err != nil {
		return err
	}
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	_, err = stmt.ExecContext(ctx, request)
	return err
}

func (o *orm) RelayReport(seqNum *big.Int) (report RelayReport, err error) {
	q := `SELECT * FROM ccip_relay_reports WHERE min_seq_num <= $1 and max_seq_num >= $1`
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	err = o.db.GetContext(ctx, &report, q, seqNum.String())
	return
}

func (o *orm) SaveRelayReport(report RelayReport) error {
	q := `INSERT INTO ccip_relay_reports (root, min_seq_num, max_seq_num, created_at) VALUES ($1, $2, $3, now()) ON CONFLICT DO NOTHING`
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	_, err := o.db.ExecContext(ctx, q, report.Root[:], report.MinSeqNum.String(), report.MaxSeqNum.String())
	return err
}
                                                                                                                                                                                                   core/services/ccip/database.go                                                                      000644  000765  000024  00000023776 14165346401 017344  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // Note lifted from https://github.com/smartcontractkit/chainlink/pull/4809/files
// TODO: Pull into common ocr library
package ccip

import (
	"context"
	"database/sql"
	"encoding/binary"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/logger"
	ocrcommon "github.com/smartcontractkit/libocr/commontypes"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"
)

type db struct {
	*sql.DB
	//oracleSpecID int32
	contractAddr common.Address
}

var (
	_ ocrtypes.Database = &db{}
	//_ OCRContractTrackerDB = &db{}
)

// NewDB returns a new DB scoped to this oracleSpecID
func NewDB(sqldb *sql.DB, contractAddr common.Address) *db {
	return &db{sqldb, contractAddr}
}

func (d *db) ReadState(ctx context.Context, cd ocrtypes.ConfigDigest) (ps *ocrtypes.PersistentState, err error) {
	q := d.QueryRowContext(ctx, `
SELECT epoch, highest_sent_epoch, highest_received_epoch
FROM ccip_persistent_states
WHERE contract_address = $1 AND config_digest = $2
LIMIT 1`, d.contractAddr, cd)

	ps = new(ocrtypes.PersistentState)

	var tmp []int64
	err = q.Scan(&ps.Epoch, &ps.HighestSentEpoch, pq.Array(&tmp))

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "ReadState failed")
	}

	for _, v := range tmp {
		ps.HighestReceivedEpoch = append(ps.HighestReceivedEpoch, uint32(v))
	}

	return ps, nil
}

func (d *db) WriteState(ctx context.Context, cd ocrtypes.ConfigDigest, state ocrtypes.PersistentState) error {
	var highestReceivedEpoch []int64
	for _, v := range state.HighestReceivedEpoch {
		highestReceivedEpoch = append(highestReceivedEpoch, int64(v))
	}
	_, err := d.ExecContext(ctx, `
INSERT INTO ccip_persistent_states (
	contract_address,
	config_digest,
	epoch,
	highest_sent_epoch,
	highest_received_epoch,
	created_at,
	updated_at
)
VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
ON CONFLICT (contract_address, config_digest)
DO UPDATE SET (
		epoch,
		highest_sent_epoch,
		highest_received_epoch,
		updated_at
	) = (
	 EXCLUDED.epoch,
	 EXCLUDED.highest_sent_epoch,
	 EXCLUDED.highest_received_epoch,
	 NOW()
	)`, d.contractAddr, cd, state.Epoch, state.HighestSentEpoch, pq.Array(&highestReceivedEpoch))

	return errors.Wrap(err, "WriteState failed")
}

func (d *db) ReadConfig(ctx context.Context) (c *ocrtypes.ContractConfig, err error) {
	q := d.QueryRowContext(ctx, `
SELECT
	config_digest,
	config_count,
	signers,
	transmitters,
	f,
	onchain_config,
	offchain_config_version,
	offchain_config
FROM ccip_contract_configs
WHERE contract_address = $1
LIMIT 1`, d.contractAddr)

	c = new(ocrtypes.ContractConfig)

	digest := []byte{}
	signers := [][]byte{}
	transmitters := [][]byte{}

	err = q.Scan(
		&digest,
		&c.ConfigCount,
		(*pq.ByteaArray)(&signers),
		(*pq.ByteaArray)(&transmitters),
		&c.F,
		&c.OnchainConfig,
		&c.OffchainConfigVersion,
		&c.OffchainConfig,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "ReadConfig failed")
	}

	copy(c.ConfigDigest[:], digest)

	c.Signers = []ocrtypes.OnchainPublicKey{}
	for _, s := range signers {
		signer := ocrtypes.OnchainPublicKey{}
		copy(signer, s[:])
		c.Signers = append(c.Signers, signer)
	}

	c.Transmitters = []ocrtypes.Account{}
	for _, t := range transmitters {
		transmitter := ocrtypes.Account(t)
		c.Transmitters = append(c.Transmitters, transmitter)
	}

	return
}

func (d *db) WriteConfig(ctx context.Context, c ocrtypes.ContractConfig) error {
	var signers [][]byte
	for _, s := range c.Signers {
		signers = append(signers, []byte(s))
	}
	_, err := d.ExecContext(ctx, `
INSERT INTO ccip_contract_configs (
	contract_address,
	config_digest,
	config_count,
	signers,
	transmitters,
	f,
	onchain_config,
	offchain_config_version,
	offchain_config,
	created_at,
	updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())
ON CONFLICT (contract_address) DO UPDATE SET
	config_digest = EXCLUDED.config_digest,
	config_count = EXCLUDED.config_count,
	signers = EXCLUDED.signers,
	transmitters = EXCLUDED.transmitters,
	f = EXCLUDED.f,
	onchain_config = EXCLUDED.onchain_config,
	offchain_config_version = EXCLUDED.offchain_config_version,
	offchain_config = EXCLUDED.offchain_config,
	updated_at = NOW()
`,
		d.contractAddr,
		c.ConfigDigest,
		c.ConfigCount,
		pq.ByteaArray(signers),
		c.Transmitters,
		c.F,
		c.OnchainConfig,
		c.OffchainConfigVersion,
		c.OffchainConfig,
	)

	return errors.Wrap(err, "WriteConfig failed")
}

func (d *db) StorePendingTransmission(ctx context.Context, t ocrtypes.ReportTimestamp, tx ocrtypes.PendingTransmission) error {
	var signatures [][]byte
	for _, s := range tx.AttributedSignatures {
		signatures = append(signatures, s.Signature)
		buffer := make([]byte, binary.MaxVarintLen64)
		binary.PutVarint(buffer, int64(s.Signer))
		signatures = append(signatures, buffer)
	}

	digest := make([]byte, 32)
	copy(digest, t.ConfigDigest[:])

	extraHash := make([]byte, 32)
	copy(extraHash[:], tx.ExtraHash[:])

	_, err := d.ExecContext(ctx, `
INSERT INTO ccip_pending_transmissions (
	contract_address,
	config_digest,
	epoch,
	round,
	time,
	extra_hash,
	report,
	attributed_signatures,
	created_at,
	updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
ON CONFLICT (contract_address, config_digest, epoch, round) DO UPDATE SET
	contract_address = EXCLUDED.contract_address,
	config_digest = EXCLUDED.config_digest,
	epoch = EXCLUDED.epoch,
	round = EXCLUDED.round,
	time = EXCLUDED.time,
	extra_hash = EXCLUDED.extra_hash,
	report = EXCLUDED.report,
	attributed_signatures = EXCLUDED.attributed_signatures,
	updated_at = NOW()
`,
		d.contractAddr,
		digest,
		t.Epoch,
		t.Round,
		tx.Time,
		extraHash,
		tx.Report,
		pq.ByteaArray(signatures),
	)

	return errors.Wrap(err, "StorePendingTransmission failed")
}

func (d *db) PendingTransmissionsWithConfigDigest(ctx context.Context, cd ocrtypes.ConfigDigest) (map[ocrtypes.ReportTimestamp]ocrtypes.PendingTransmission, error) {
	rows, err := d.QueryContext(ctx, `
SELECT
	config_digest,
	epoch,
	round,
	time,
	extra_hash,
	report,
	attributed_signatures
FROM ccip_pending_transmissions
WHERE contract_address = $1 AND config_digest = $2
`, d.contractAddr, cd)
	if err != nil {
		return nil, errors.Wrap(err, "PendingTransmissionsWithConfigDigest failed to query rows")
	}
	defer logger.ErrorIfCalling(rows.Close)
	if rows.Err() != nil {
		return nil, errors.Wrap(err, "PendingTransmissionsWithConfigDigest row iteration error")
	}

	m := make(map[ocrtypes.ReportTimestamp]ocrtypes.PendingTransmission)

	for rows.Next() {
		k := ocrtypes.ReportTimestamp{}
		p := ocrtypes.PendingTransmission{}

		signatures := [][]byte{}
		digest := []byte{}
		extraHash := []byte{}
		report := []byte{}

		if err := rows.Scan(&digest, &k.Epoch, &k.Round, &p.Time, &extraHash, &report, (*pq.ByteaArray)(&signatures)); err != nil {
			return nil, errors.Wrap(err, "PendingTransmissionsWithConfigDigest failed to scan row")
		}

		copy(k.ConfigDigest[:], digest)
		copy(p.ExtraHash[:], extraHash)
		p.Report = make([]byte, len(report))
		copy(p.Report[:], report)

		for index := 0; index < len(signatures); index += 2 {
			signature := signatures[index]
			signer, _ := binary.Varint(signatures[index+1])
			sig := ocrtypes.AttributedOnchainSignature{
				Signature: signature,
				Signer:    ocrcommon.OracleID(signer),
			}
			p.AttributedSignatures = append(p.AttributedSignatures, sig)
		}
		m[k] = p
	}

	return m, nil
}

func (d *db) DeletePendingTransmission(ctx context.Context, t ocrtypes.ReportTimestamp) (err error) {
	_, err = d.ExecContext(ctx, `
DELETE FROM ccip_pending_transmissions
WHERE contract_address = $1 AND  config_digest = $2 AND epoch = $3 AND round = $4
`, d.contractAddr, t.ConfigDigest, t.Epoch, t.Round)

	err = errors.Wrap(err, "DeletePendingTransmission failed")

	return
}

func (d *db) DeletePendingTransmissionsOlderThan(ctx context.Context, t time.Time) (err error) {
	_, err = d.ExecContext(ctx, `
DELETE FROM ccip_pending_transmissions
WHERE contract_address = $1 AND time < $2
`, d.contractAddr, t)

	err = errors.Wrap(err, "DeletePendingTransmissionsOlderThan failed")

	return
}

//func (d *db) SaveLatestRoundRequested(tx *sql.Tx, rr ocr2aggregator.OCR2AggregatorRoundRequested) error {
//	rawLog, err := json.Marshal(rr.Raw)
//	if err != nil {
//		return errors.Wrap(err, "could not marshal log as JSON")
//	}
//	_, err = tx.Exec(`
//INSERT INTO offchainreporting2_latest_round_requested (offchainreporting2_oracle_spec_id, requester, config_digest, epoch, round, raw)
//VALUES ($1,$2,$3,$4,$5,$6) ON CONFLICT (offchainreporting2_oracle_spec_id) DO UPDATE SET
//	requester = EXCLUDED.requester,
//	config_digest = EXCLUDED.config_digest,
//	epoch = EXCLUDED.epoch,
//	round = EXCLUDED.round,
//	raw = EXCLUDED.raw
//`, d.oracleSpecID, rr.Requester, rr.ConfigDigest[:], rr.Epoch, rr.Round, rawLog)
//
//	return errors.Wrap(err, "could not save latest round requested")
//}
//
//func (d *db) LoadLatestRoundRequested() (ocr2aggregator.OCR2AggregatorRoundRequested, error) {
//	rr := ocr2aggregator.OCR2AggregatorRoundRequested{}
//	rows, err := d.Query(`
//SELECT requester, config_digest, epoch, round, raw
//FROM offchainreporting2_latest_round_requested
//WHERE offchainreporting2_oracle_spec_id = $1
//LIMIT 1
//`, d.oracleSpecID)
//	if err != nil {
//		return rr, errors.Wrap(err, "LoadLatestRoundRequested failed to query rows")
//	}
//	if rows.Err() != nil {
//		return rr, errors.Wrap(err, "LoadLatestRoundRequested row iteration error")
//	}
//
//	for rows.Next() {
//		var configDigest []byte
//		var rawLog []byte
//
//		err = rows.Scan(&rr.Requester, &configDigest, &rr.Epoch, &rr.Round, &rawLog)
//		if err != nil {
//			return rr, errors.Wrap(err, "LoadLatestRoundRequested failed to scan row")
//		}
//
//		rr.ConfigDigest, err = ocrtypes.BytesToConfigDigest(configDigest)
//		if err != nil {
//			return rr, errors.Wrap(err, "LoadLatestRoundRequested failed to decode config digest")
//		}
//
//		err = json.Unmarshal(rawLog, &rr.Raw)
//		if err != nil {
//			return rr, errors.Wrap(err, "LoadLatestRoundRequested failed to unmarshal raw log")
//		}
//	}
//
//	return rr, nil
//}
  core/services/ccip/execution_reporting_plugin_test.go                                               000644  000765  000024  00000033027 14165357743 024312  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip_test

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/ccip/abihelpers"

	"github.com/lib/pq"
	"github.com/smartcontractkit/chainlink/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/services/ccip"
	lastreportermocks "github.com/smartcontractkit/chainlink/core/services/ccip/mocks/lastreporter"
	"github.com/smartcontractkit/chainlink/core/utils"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/stretchr/testify/mock"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/lock_unlock_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor_helper"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp_helper"
	"github.com/smartcontractkit/libocr/gethwrappers/link_token_interface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecutionReportEncoding(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	destUser, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	destChain := backends.NewSimulatedBackend(core.GenesisAlloc{
		destUser.From: {Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1000000000000000000))}},
		ethconfig.Defaults.Miner.GasCeil)
	// Deploy link token
	destLinkTokenAddress, _, destLinkToken, err := link_token_interface.DeployLinkToken(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()
	_, err = link_token_interface.NewLinkToken(destLinkTokenAddress, destChain)
	require.NoError(t, err)

	// Deploy destination pool
	destPoolAddress, _, _, err := lock_unlock_pool.DeployLockUnlockPool(destUser, destChain, destLinkTokenAddress)
	require.NoError(t, err)
	destChain.Commit()
	destPool, err := lock_unlock_pool.NewLockUnlockPool(destPoolAddress, destChain)
	require.NoError(t, err)

	// Fund dest pool
	_, err = destLinkToken.Approve(destUser, destPoolAddress, big.NewInt(1000000))
	require.NoError(t, err)
	destChain.Commit()
	_, err = destPool.LockOrBurn(destUser, destUser.From, big.NewInt(1000000))
	require.NoError(t, err)
	destChain.Commit()

	afnAddress, _, _, err := afn_contract.DeployAFNContract(
		destUser,
		destChain,
		[]common.Address{destUser.From},
		[]*big.Int{big.NewInt(1)},
		big.NewInt(1),
		big.NewInt(1),
	)
	require.NoError(t, err)
	destChain.Commit()

	offRampAddress, _, _, err := single_token_offramp_helper.DeploySingleTokenOffRampHelper(
		destUser,
		destChain,
		big.NewInt(1337),
		big.NewInt(1338),
		destLinkTokenAddress,
		destPoolAddress,
		big.NewInt(1),    // token bucket rate
		big.NewInt(1000), // token bucket capacity
		afnAddress,       // AFN address
		// 86400 seconds = one day
		big.NewInt(86400), // max timeout without AFN signal
		big.NewInt(0),     // execution delay in seconds
	)
	require.NoError(t, err)
	offRamp, err := single_token_offramp_helper.NewSingleTokenOffRampHelper(offRampAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()
	_, err = destPool.SetOffRamp(destUser, offRampAddress, true)
	require.NoError(t, err)
	receiverAddress, _, _, err := simple_message_receiver.DeploySimpleMessageReceiver(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()

	m := ccip.Request{
		SeqNum:        *utils.NewBigI(10),
		SourceChainID: "1337",
		DestChainID:   "1338",
		Sender:        destUser.From,
		Receiver:      receiverAddress,
		Data:          []byte("hello"),
		Tokens:        []string{destLinkTokenAddress.String()},
		Amounts:       []string{"100"},
		Options:       []byte{},
	}
	msgBytes, err := abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{m.ToMessage()})
	require.NoError(t, err)
	r, proof := ccip.GenerateMerkleProof(2, [][]byte{msgBytes}, 0)
	var root [32]byte
	copy(root[:], r[:])
	rootLocal := ccip.GenerateMerkleRoot(msgBytes, proof)
	require.True(t, bytes.Equal(rootLocal[:], r[:]))

	out, err := ccip.EncodeRelayReport(&single_token_offramp.CCIPRelayReport{
		MerkleRoot:        root,
		MinSequenceNumber: big.NewInt(10),
		MaxSequenceNumber: big.NewInt(10),
	})
	require.NoError(t, err)
	_, err = ccip.DecodeRelayReport(out)
	require.NoError(t, err)

	// RelayReport that Message
	tx, err := offRamp.Report(destUser, out)
	require.NoError(t, err)
	destChain.Commit()

	// Now execute that Message via the executor
	t.Log(offRampAddress)
	executorAddress, _, _, err := message_executor_helper.DeployMessageExecutorHelper(
		destUser,
		destChain,
		offRampAddress)
	require.NoError(t, err)
	executor, err := message_executor_helper.NewMessageExecutorHelper(executorAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	executorReport, err := ccip.EncodeExecutionReport([]ccip.ExecutableMessage{
		{
			Proof:   proof.PathForExecute(),
			Message: m.ToMessage(),
			Index:   proof.Index(),
		},
	})
	require.NoError(t, err)
	ems, err := ccip.DecodeExecutionReport(executorReport)
	require.NoError(t, err)
	t.Log(ems)

	generatedRoot, err := offRamp.GenerateMerkleRoot(nil, proof.PathForExecute(), ccip.HashLeaf(msgBytes), proof.Index())
	require.NoError(t, err)
	require.Equal(t, root, generatedRoot)
	tx, err = executor.Report(destUser, executorReport)
	require.NoError(t, err)
	destChain.Commit()
	res, err := destChain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	assert.Equal(t, uint64(1), res.Status)
}

func TestExecutionPlugin(t *testing.T) {
	// Avoid using txdb: it has bugs and currently has savepoints disabled (to be able to use with gorm)
	// and so any ctx cancellation poisons the tx.
	_, db, _ := heavyweight.FullTestDB(t, "executor_plugin", true, false)
	orm := ccip.NewORM(db)
	lr := new(lastreportermocks.OffRampLastReporter)
	executor := common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB5")
	rf := ccip.NewExecutionReportingPluginFactory(logger.TestLogger(t), orm, big.NewInt(1), big.NewInt(2), executor, lr)
	rp, _, err := rf.NewReportingPlugin(types.ReportingPluginConfig{F: 1})
	require.NoError(t, err)
	sid, did := big.NewInt(1), big.NewInt(2)
	// Observe with nothing in the db should error with no observations
	obs, err := rp.Observation(context.Background(), types.ReportTimestamp{}, types.Query{})
	require.Error(t, err)
	require.Len(t, obs, 0)

	// Observe with a non-relay-confirmed request should still return no requests
	req := ccip.Request{
		SeqNum:        *utils.NewBigI(2),
		SourceChainID: sid.String(),
		DestChainID:   did.String(),
		Sender:        common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
		Data:          []byte("hello"),
		Tokens:        pq.StringArray{},
		Amounts:       pq.StringArray{},
		Executor:      executor,
		Options:       []byte{},
	}
	b, err := abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{req.ToMessage()})
	require.NoError(t, err)
	req.Raw = b
	require.NoError(t, orm.SaveRequest(&req))
	obs, err = rp.Observation(context.Background(), types.ReportTimestamp{}, types.Query{})
	require.Error(t, err)
	require.Equal(t, "no requests for oracle execution", err.Error())
	require.Len(t, obs, 0)

	// We should see an error if the latest report doesn't have a higher seq num
	lr.On("GetLastReport", mock.Anything).Return(getLastReportMock(1)).Once()
	require.NoError(t, orm.UpdateRequestSetStatus(sid, did, []*big.Int{big.NewInt(2)}, ccip.RequestStatusRelayConfirmed))
	obs, err = rp.Observation(context.Background(), types.ReportTimestamp{}, types.Query{})
	require.Error(t, err)
	// Should succeed if we do have a higher seq num
	lr.On("GetLastReport", mock.Anything).Return(getLastReportMock(2)).Once()
	obs, err = rp.Observation(context.Background(), types.ReportTimestamp{}, types.Query{})
	require.NoError(t, err)
	var o ccip.Observation
	require.NoError(t, json.Unmarshal(obs, &o))
	require.Equal(t, "2", o.MinSeqNum.String())
	require.Equal(t, "2", o.MaxSeqNum.String())

	// If all the nodes report the same, this should succeed
	// First add the relay report
	root, _ := ccip.GenerateMerkleProof(32, [][]byte{b}, 0)
	require.NoError(t, orm.SaveRelayReport(ccip.RelayReport{Root: root[:], MinSeqNum: *utils.NewBigI(2), MaxSeqNum: *utils.NewBigI(2)}))
	lr.On("GetLastReport", mock.Anything).Return(getLastReportMock(2)).Once()
	finalizeReport, rep, err := rp.Report(context.Background(), types.ReportTimestamp{}, types.Query{}, []types.AttributedObservation{
		{Observation: obs}, {Observation: obs}, {Observation: obs}, {Observation: obs},
	})
	require.NoError(t, err)
	require.True(t, finalizeReport)
	executableMessages, err := ccip.DecodeExecutionReport(rep)
	require.NoError(t, err)
	// Should see our one message there
	require.Len(t, executableMessages, 1)
	require.Equal(t, "2", executableMessages[0].Message.SequenceNumber.String())

	// If we have < F observations, we should not get a report
	finalizeReport, rep, err = rp.Report(context.Background(), types.ReportTimestamp{}, types.Query{}, []types.AttributedObservation{
		{Observation: nil}, {Observation: nil}, {Observation: nil}, {Observation: obs},
	})
	require.False(t, finalizeReport)
	// With F=1, that means a single value cannot corrupt our report
	var fakeObs = ccip.Observation{
		MinSeqNum: *utils.NewBigI(10000),
		MaxSeqNum: *utils.NewBigI(10000),
	}
	b, err = json.Marshal(fakeObs)
	require.NoError(t, err)
	lr.On("GetLastReport", mock.Anything).Return(getLastReportMock(2)).Once()
	finalizeReport, rep, err = rp.Report(context.Background(), types.ReportTimestamp{}, types.Query{}, []types.AttributedObservation{
		{Observation: obs}, {Observation: obs}, {Observation: obs}, {Observation: b},
	})
	require.NoError(t, err)
	// Still our message 2 despite the fakeObs
	executableMessages, err = ccip.DecodeExecutionReport(rep)
	require.NoError(t, err)
	require.Len(t, executableMessages, 1)
	require.Equal(t, "2", executableMessages[0].Message.SequenceNumber.String())

	// Should not accept or transmit if the report is stale
	orm.UpdateRequestSetStatus(sid, did, []*big.Int{big.NewInt(2)}, ccip.RequestStatusExecutionConfirmed)
	accept, err := rp.ShouldAcceptFinalizedReport(context.Background(), types.ReportTimestamp{}, rep)
	require.NoError(t, err)
	require.False(t, accept)
	accept, err = rp.ShouldTransmitAcceptedReport(context.Background(), types.ReportTimestamp{}, rep)
	require.NoError(t, err)
	require.False(t, accept)

	// Ensure observing and reporting works with batches.
	// Let's save a batch of seqnums {3,4,5}
	var leaves [][]byte
	for i := 3; i < 6; i++ {
		req := ccip.Request{
			SeqNum:        *utils.NewBigI(int64(i)),
			SourceChainID: sid.String(),
			DestChainID:   did.String(),
			Sender:        common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
			Data:          []byte("hello"),
			Tokens:        pq.StringArray{},
			Amounts:       pq.StringArray{},
			Executor:      executor,
			Options:       []byte{},
		}
		b, err := abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{req.ToMessage()})
		require.NoError(t, err)
		req.Raw = b
		require.NoError(t, orm.SaveRequest(&req))
		leaves = append(leaves, b)
	}
	require.NoError(t, orm.UpdateRequestStatus(sid, did, big.NewInt(3), big.NewInt(5), ccip.RequestStatusRelayConfirmed))
	lr.On("GetLastReport", mock.Anything).Return(getLastReportMock(5)).Once()
	obs, err = rp.Observation(context.Background(), types.ReportTimestamp{}, types.Query{})
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(obs, &o))
	require.Equal(t, "3", o.MinSeqNum.String())
	require.Equal(t, "5", o.MaxSeqNum.String())

	// Let's put 2 in one report and 1 in a different report then assert the execution report makes sense
	root1, _ := ccip.GenerateMerkleProof(32, [][]byte{leaves[0]}, 0)
	require.NoError(t, orm.SaveRelayReport(ccip.RelayReport{Root: root1[:], MinSeqNum: *utils.NewBigI(3), MaxSeqNum: *utils.NewBigI(3)}))
	root2, _ := ccip.GenerateMerkleProof(32, [][]byte{leaves[1], leaves[2]}, 0)
	require.NoError(t, orm.SaveRelayReport(ccip.RelayReport{Root: root2[:], MinSeqNum: *utils.NewBigI(4), MaxSeqNum: *utils.NewBigI(5)}))
	lr.On("GetLastReport", mock.Anything).Return(getLastReportMock(5)).Once()
	finalizeReport, rep, err = rp.Report(context.Background(), types.ReportTimestamp{}, types.Query{}, []types.AttributedObservation{
		{Observation: obs}, {Observation: obs}, {Observation: obs}, {Observation: obs},
	})
	require.NoError(t, err)
	msgs, err := ccip.DecodeExecutionReport(rep)
	require.NoError(t, err)
	require.Len(t, msgs, 3)
	rootLeaf1 := ccip.GenerateMerkleRoot(leaves[0], ccip.NewMerkleProof(int(msgs[0].Index.Int64()), msgs[0].Proof))
	rootLeaf2 := ccip.GenerateMerkleRoot(leaves[1], ccip.NewMerkleProof(int(msgs[1].Index.Int64()), msgs[1].Proof))
	rootLeaf3 := ccip.GenerateMerkleRoot(leaves[1], ccip.NewMerkleProof(int(msgs[1].Index.Int64()), msgs[1].Proof))
	require.True(t, bytes.Equal(rootLeaf1[:], root1[:]))
	require.True(t, bytes.Equal(rootLeaf2[:], root2[:]))
	require.True(t, bytes.Equal(rootLeaf3[:], root2[:]))
}

func getLastReportMock(maxSequenceNumber int64) (single_token_offramp.CCIPRelayReport, error) {
	maxSequenceNumberBig := big.NewInt(maxSequenceNumber)
	return single_token_offramp.CCIPRelayReport{
		MaxSequenceNumber: maxSequenceNumberBig,
	}, nil
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         core/services/ccip/delegate_relayer.go                                                              000644  000765  000024  00000014405 14165346401 021062  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"context"
	"strings"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/services/bulletprooftxmanager"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/postgres"
	ocr "github.com/smartcontractkit/libocr/offchainreporting2"
	"github.com/smartcontractkit/libocr/offchainreporting2/chains/evmutil"
	"gorm.io/gorm"
)

var _ job.Delegate = (*RelayDelegate)(nil)

type RelayDelegate struct {
	db          *gorm.DB
	jobORM      job.ORM
	orm         ORM
	chainSet    evm.ChainSet
	keyStore    keystore.OCR2
	peerWrapper *ocrcommon.SingletonPeerWrapper
}

// TODO: Register this delegate behind a FF
func NewRelayDelegate(
	db *gorm.DB,
	jobORM job.ORM,
	chainSet evm.ChainSet,
	keyStore keystore.OCR2,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
) *RelayDelegate {
	return &RelayDelegate{
		db:          db,
		jobORM:      jobORM,
		orm:         NewORM(postgres.UnwrapGormDB(db)),
		chainSet:    chainSet,
		keyStore:    keyStore,
		peerWrapper: peerWrapper,
	}
}

func (d RelayDelegate) JobType() job.Type {
	return job.CCIPRelay
}

func (d RelayDelegate) getOracleArgs(l logger.Logger, jb job.Job, offRamp *single_token_offramp.SingleTokenOffRamp, chain evm.Chain, contractTracker *CCIPContractTracker, offchainConfigDigester evmutil.EVMOffchainConfigDigester) (*ocr.OracleArgs, error) {
	ta, err := getTransmitterAddress(jb.CCIPRelaySpec.TransmitterAddress, chain)
	if err != nil {
		return nil, err
	}
	offRampABI, err := abi.JSON(strings.NewReader(single_token_offramp.SingleTokenOffRampABI))
	if err != nil {
		return nil, errors.Wrap(err, "could not get contract ABI JSON")
	}
	contractTransmitter := NewOfframpTransmitter(
		offRamp,
		offRampABI,
		NewRelayTransmitter(chain.TxManager(),
			d.db,
			jb.CCIPRelaySpec.SourceEVMChainID.ToInt(),
			jb.CCIPRelaySpec.DestEVMChainID.ToInt(), ta.Address(),
			chain.Config().EvmGasLimitDefault(),
			bulletprooftxmanager.NewQueueingTxStrategy(jb.ExternalJobID,
				chain.Config().OCR2DefaultTransactionQueueDepth(), false),
			chain.Client()),
	)
	ocrLogger := logger.NewOCRWrapper(l, true, func(msg string) {
		d.jobORM.RecordError(context.Background(), jb.ID, msg)
	})
	key, err := getValidatedKeyBundle(jb.CCIPRelaySpec.EncryptedOCRKeyBundleID, chain, d.keyStore)
	if err != nil {
		return nil, err
	}
	if err = validatePeerWrapper(jb.CCIPRelaySpec.P2PPeerID, chain, d.peerWrapper); err != nil {
		return nil, err
	}
	bootstrapPeers, err := getValidatedBootstrapPeers(jb.CCIPRelaySpec.P2PBootstrapPeers, chain)
	if err != nil {
		return nil, err
	}

	gormdb, errdb := d.db.DB()
	if errdb != nil {
		return nil, errors.Wrap(errdb, "unable to open sql db")
	}
	ocrdb := NewDB(gormdb, jb.CCIPRelaySpec.OffRampAddress.Address())
	return &ocr.OracleArgs{
		BinaryNetworkEndpointFactory: d.peerWrapper.Peer2,
		V2Bootstrappers:              bootstrapPeers,
		ContractTransmitter:          contractTransmitter,
		ContractConfigTracker:        contractTracker,
		Database:                     ocrdb,
		LocalConfig: computeLocalConfig(chain.Config(), chain.Config().Dev(),
			jb.CCIPRelaySpec.BlockchainTimeout.Duration(),
			jb.CCIPRelaySpec.ContractConfigConfirmations, jb.CCIPRelaySpec.ContractConfigTrackerPollInterval.Duration()),
		Logger:                 ocrLogger,
		MonitoringEndpoint:     nil, // TODO
		OffchainConfigDigester: offchainConfigDigester,
		OffchainKeyring:        &key.OffchainKeyring,
		OnchainKeyring:         &key.OnchainKeyring,
		ReportingPluginFactory: NewRelayReportingPluginFactory(l, d.orm, offRamp),
	}, nil
}

func (d RelayDelegate) ServicesForSpec(jb job.Job) ([]job.Service, error) {
	if jb.CCIPRelaySpec == nil {
		return nil, errors.New("no ccip job specified")
	}
	l := logger.Default.With(
		"jobID", jb.ID,
		"externalJobID", jb.ExternalJobID,
		"offRampContract", jb.CCIPRelaySpec.OffRampAddress,
		"onRampContract", jb.CCIPRelaySpec.OnRampAddress,
	)

	destChain, err := d.chainSet.Get(jb.CCIPRelaySpec.DestEVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open chain")
	}
	sourceChain, err := d.chainSet.Get(jb.CCIPRelaySpec.SourceEVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open chain")
	}
	offRamp, err := single_token_offramp.NewSingleTokenOffRamp(jb.CCIPRelaySpec.OffRampAddress.Address(), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not instantiate NewOffchainAggregator")
	}
	contractTracker := NewCCIPContractTracker(
		offrampTracker{offRamp},
		destChain.Client(),
		destChain.LogBroadcaster(),
		jb.ID,
		logger.Default,
		d.db,
		destChain,
		destChain.HeadBroadcaster(),
	)
	offchainConfigDigester := evmutil.EVMOffchainConfigDigester{
		ChainID:         maybeRemapChainID(destChain.Config().ChainID()).Uint64(),
		ContractAddress: jb.CCIPRelaySpec.OffRampAddress.Address(),
	}
	oracleArgs, err := d.getOracleArgs(l, jb, offRamp, destChain, contractTracker, offchainConfigDigester)
	if err != nil {
		return nil, err
	}
	oracle, err := ocr.NewOracle(*oracleArgs)
	if err != nil {
		return nil, err
	}
	singleTokenOnRamp, err := single_token_onramp.NewSingleTokenOnRamp(jb.CCIPRelaySpec.OnRampAddress.Address(), sourceChain.Client())
	if err != nil {
		return nil, err
	}

	encodedCCIPConfig, err := contractTracker.GetOffchainConfig()
	if err != nil {
		return nil, errors.Wrap(err, "could not get the latest encoded config")
	}
	// TODO: Its conceivable we may want pull out this log listener into its own job spec so to avoid repeating
	// All the log subscriptions
	logListener := NewLogListener(l,
		sourceChain.LogBroadcaster(),
		destChain.LogBroadcaster(),
		singleTokenOnRamp,
		offRamp,
		encodedCCIPConfig,
		d.db,
		jb.ID)
	return []job.Service{contractTracker, oracle, logListener}, nil
}

func (d RelayDelegate) AfterJobCreated(spec job.Job) {
}

func (d RelayDelegate) BeforeJobDeleted(spec job.Job) {
}
                                                                                                                                                                                                                                                           core/services/ccip/abihelpers/abi_helpers.go                                                        000644  000765  000024  00000004723 14165357743 022175  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package abihelpers

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
)

func DecodeCCIPMessage(b []byte) (*single_token_offramp.CCIPMessage, error) {
	unpacked, err := MakeCCIPMsgArgs().Unpack(b)
	if err != nil {
		return nil, err
	}
	// Note must use unnamed type here
	receivedCp, ok := unpacked[0].(struct {
		SequenceNumber     *big.Int       `json:"sequenceNumber"`
		SourceChainId      *big.Int       `json:"sourceChainId"`
		DestinationChainId *big.Int       `json:"destinationChainId"`
		Sender             common.Address `json:"sender"`
		Payload            struct {
			Receiver common.Address   `json:"receiver"`
			Data     []uint8          `json:"data"`
			Tokens   []common.Address `json:"tokens"`
			Amounts  []*big.Int       `json:"amounts"`
			Executor common.Address   `json:"executor"`
			Options  []uint8          `json:"options"`
		} `json:"payload"`
	})
	if !ok {
		return nil, fmt.Errorf("invalid format have %T want %T", unpacked[0], receivedCp)
	}
	return &single_token_offramp.CCIPMessage{
		SequenceNumber:     receivedCp.SequenceNumber,
		SourceChainId:      receivedCp.SourceChainId,
		DestinationChainId: receivedCp.DestinationChainId,
		Sender:             receivedCp.Sender,
		Payload: single_token_offramp.CCIPMessagePayload{
			Receiver: receivedCp.Payload.Receiver,
			Data:     receivedCp.Payload.Data,
			Tokens:   receivedCp.Payload.Tokens,
			Amounts:  receivedCp.Payload.Amounts,
			Executor: receivedCp.Payload.Executor,
			Options:  receivedCp.Payload.Options,
		},
	}, nil
}

func MakeCCIPMsgArgs() abi.Arguments {
	var tuples = []abi.ArgumentMarshaling{
		{
			Name: "sequenceNumber",
			Type: "uint256",
		},
		{
			Name: "sourceChainId",
			Type: "uint256",
		},
		{
			Name: "destinationChainId",
			Type: "uint256",
		},
		{
			Name: "sender",
			Type: "address",
		},
		{
			Name: "payload",
			Type: "tuple",
			Components: []abi.ArgumentMarshaling{
				{
					Name: "receiver",
					Type: "address",
				},
				{
					Name: "data",
					Type: "bytes",
				},
				{
					Name: "tokens",
					Type: "address[]",
				},
				{
					Name: "amounts",
					Type: "uint256[]",
				},
				{
					Name: "executor",
					Type: "address",
				},
				{
					Name: "options",
					Type: "bytes",
				},
			},
		},
	}
	ty, _ := abi.NewType("tuple", "", tuples)
	return abi.Arguments{
		{
			Type: ty,
		},
	}
}
                                             core/services/ccip/merkle.go                                                                        000644  000765  000024  00000006022 14165346401 017040  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"math/big"

	"golang.org/x/crypto/sha3"
)

type MerkleProof struct {
	index int
	path  [][32]byte
}

func NewMerkleProof(index int, path [][32]byte) MerkleProof {
	return MerkleProof{
		index: index,
		path:  path,
	}
}

func (mp MerkleProof) PathForExecute() [][32]byte {
	return mp.path
}

func (mp MerkleProof) Index() *big.Int {
	return big.NewInt(int64(mp.index))
}

func GenerateMerkleProof(treeHeight int, leaves [][]byte, index int) ([32]byte, MerkleProof) {
	zhs := computeZeroHashes(treeHeight)
	var level [][32]byte
	for _, leaf := range leaves {
		level = append(level, HashLeaf(leaf))
	}
	levelIndex := index
	var path [][32]byte
	// Go level by level up the tree starting from the bottom.
	// Record the path of sibling nodes for the Index node required
	// to get the top.
	for height := 0; height < treeHeight-1; height++ {
		// If we have an odd number of level elements
		if len(level)%2 == 1 {
			level = append(level, zhs[height])
		}
		pathIndex := levelIndex + 1
		if levelIndex%2 == 1 {
			pathIndex = levelIndex - 1
		}
		path = append(path, level[pathIndex])
		// Floor division here
		// E.g. [0, 1, 2, 3]
		// maps to [0, 1] on the next level.
		// So 0,1 -> 0 and 2,3 -> 1
		levelIndex /= 2
		// Compute the next level by hashing each pair of nodes.
		// (we know there is an even number of them)
		var newLevel [][32]byte
		for i := 0; i < len(level)-1; i += 2 {
			newLevel = append(newLevel, hashInternal(level[i], level[i+1]))
		}
		level = newLevel
	}
	if len(level) != 1 {
		panic("invalid")
	}
	return level[0], MerkleProof{path: path, index: index}
}

func GenerateMerkleRoot(leaf []byte, proof MerkleProof) [32]byte {
	// Make a deep copy of the path
	var path [][32]byte
	for _, p := range proof.path {
		var pc [32]byte
		copy(pc[:], p[:])
		path = append(path, pc)
	}
	index := proof.index
	h := HashLeaf(leaf)
	var l, r [32]byte
	for {
		if len(path) == 0 {
			break
		}
		if index%2 == 0 {
			// if Index is even then our Index is on the left
			// and the Proof element is on the right
			l = h
			r = path[0]
		} else {
			// if Index is odd then our Index is on the right
			// and the Proof element is on the left
			l = path[0]
			r = h
		}
		path = path[1:] // done with that Proof element
		h = hashInternal(l, r)
		index >>= 1
	}
	return h
}

func hashInternal(l, r [32]byte) [32]byte {
	hash := sha3.NewLegacyKeccak256()
	// Ignore errors
	hash.Write([]byte{0x01})
	hash.Write(l[:])
	hash.Write(r[:])
	var res [32]byte
	copy(res[:], hash.Sum(nil))
	return res
}

func HashLeaf(b []byte) [32]byte {
	hash := sha3.NewLegacyKeccak256()
	// Ignore errors
	hash.Write([]byte{0x00})
	hash.Write(b)
	var r [32]byte
	copy(r[:], hash.Sum(nil))
	return r
}

func computeZeroHashes(height int) [][32]byte {
	// Pre-compute all-zero trees for each depth
	// i.e. [0x00, H(0x00), H(H(0x00)||H(0x00)), ...]
	var zeroHashes = make([][32]byte, height)
	for i := 0; i < height-1; i++ {
		if i == 0 {
			var zh [32]byte
			zeroHashes[i] = zh
		}
		zeroHashes[i+1] = hashInternal(zeroHashes[i], zeroHashes[i])
	}
	return zeroHashes
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              core/services/ccip/relay_reporting_plugin.go                                                        000644  000765  000024  00000021633 14165346401 022351  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sort"

	"github.com/smartcontractkit/chainlink/core/logger"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink/core/utils"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"

	"github.com/smartcontractkit/libocr/offchainreporting2/types"
)

const (
	RELAY_MAX_INFLIGHT_TIME_SECONDS = 180
)

var _ types.ReportingPluginFactory = &RelayReportingPluginFactory{}
var _ types.ReportingPlugin = &RelayReportingPlugin{}

type Observation struct {
	MinSeqNum utils.Big `json:"minSeqNum"`
	MaxSeqNum utils.Big `json:"maxSeqNum"`
}

func makeRelayReportArgs() abi.Arguments {
	mustType := func(ts string) abi.Type {
		ty, _ := abi.NewType(ts, "", nil)
		return ty
	}
	return []abi.Argument{
		{
			Name: "merkleRoot",
			Type: mustType("bytes32"),
		},
		{
			Name: "minSequenceNumber",
			Type: mustType("uint256"),
		},
		{
			Name: "maxSequenceNumber",
			Type: mustType("uint256"),
		},
	}
}

func EncodeRelayReport(relayReport *single_token_offramp.CCIPRelayReport) (types.Report, error) {
	report, err := makeRelayReportArgs().PackValues([]interface{}{relayReport.MerkleRoot, relayReport.MinSequenceNumber, relayReport.MaxSequenceNumber})
	if err != nil {
		return nil, err
	}
	return report, nil
}

func DecodeRelayReport(report types.Report) (*single_token_offramp.CCIPRelayReport, error) {
	unpacked, err := makeRelayReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) != 3 {
		return nil, errors.New("invalid num fields in report")
	}
	root, ok := unpacked[0].([32]byte)
	if !ok {
		return nil, errors.New("invalid root")
	}
	min, ok := unpacked[1].(*big.Int)
	if !ok {
		return nil, errors.New("invalid min")
	}
	max, ok := unpacked[1].(*big.Int)
	if !ok {
		return nil, errors.New("invalid max")
	}
	return &single_token_offramp.CCIPRelayReport{
		MerkleRoot:        root,
		MinSequenceNumber: min,
		MaxSequenceNumber: max,
	}, nil
}

type RelayReportingPluginFactory struct {
	l       logger.Logger
	orm     ORM
	jobID   int32
	offRamp *single_token_offramp.SingleTokenOffRamp
}

func NewRelayReportingPluginFactory(l logger.Logger, orm ORM, offRamp *single_token_offramp.SingleTokenOffRamp) types.ReportingPluginFactory {
	return &RelayReportingPluginFactory{l: l, orm: orm, offRamp: offRamp}
}

func (rf *RelayReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	destChainId, err := rf.offRamp.CHAINID(nil)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	sourcChainId, err := rf.offRamp.SOURCECHAINID(nil)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	return RelayReportingPlugin{rf.l, config.F, rf.orm, sourcChainId, destChainId, rf.offRamp}, types.ReportingPluginInfo{
		Name:              "CCIPRelay",
		UniqueReports:     true,
		MaxQueryLen:       0,      // We do not use the query phase.
		MaxObservationLen: 100000, // TODO
		MaxReportLen:      100000, // TODO
	}, nil
}

type RelayReportingPlugin struct {
	l             logger.Logger
	F             int
	orm           ORM
	sourceChainId *big.Int
	destChainId   *big.Int
	offRamp       *single_token_offramp.SingleTokenOffRamp
}

func (r RelayReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	return types.Query{}, nil
}

func (r RelayReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	lastReport, err := r.offRamp.GetLastReport(nil)
	if err != nil {
		return nil, err
	}
	unstartedReqs, err := r.orm.Requests(r.sourceChainId, r.destChainId, big.NewInt(0).Add(lastReport.MaxSequenceNumber, big.NewInt(1)), nil, RequestStatusUnstarted, nil, nil)
	if err != nil {
		return nil, err
	}
	// No request to process
	// Return an empty observation
	// which should not result in a report generated.
	if len(unstartedReqs) == 0 {
		return nil, fmt.Errorf("no requests with seq num greater than %v", lastReport.MaxSequenceNumber)
	}

	b, err := json.Marshal(&Observation{
		MinSeqNum: unstartedReqs[0].SeqNum,
		MaxSeqNum: unstartedReqs[len(unstartedReqs)-1].SeqNum,
	})
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r RelayReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	// Need at least F+1 valid observations
	var nonEmptyObservations []Observation
	for _, ao := range observations {
		var ob Observation
		err := json.Unmarshal(ao.Observation, &ob)
		if err != nil {
			r.l.Errorw("received unmarshallable observation", "err", err)
			continue
		}
		nonEmptyObservations = append(nonEmptyObservations, ob)
	}
	if len(nonEmptyObservations) <= r.F {
		return false, nil, nil
	}
	// We have at least F+1 valid observations
	// Extract the min and max
	sort.Slice(nonEmptyObservations, func(i, j int) bool {
		return nonEmptyObservations[i].MinSeqNum.ToInt().Cmp(nonEmptyObservations[j].MinSeqNum.ToInt()) < 0
	})
	min := nonEmptyObservations[r.F].MinSeqNum.ToInt()
	sort.Slice(nonEmptyObservations, func(i, j int) bool {
		return nonEmptyObservations[i].MaxSeqNum.ToInt().Cmp(nonEmptyObservations[j].MaxSeqNum.ToInt()) < 0
	})
	max := nonEmptyObservations[r.F].MaxSeqNum.ToInt()
	if max.Cmp(min) < 0 {
		return false, nil, errors.New("max seq num smaller than min")
	}
	reqs, err := r.orm.Requests(r.sourceChainId, r.destChainId, min, nil, RequestStatusUnstarted, nil, nil)
	if err != nil {
		return false, nil, err
	}
	// Cannot construct a report for which we haven't seen all the messages.
	if len(reqs) == 0 {
		return false, nil, fmt.Errorf("do not have all the messages in report, have zero messages, report has min %v max %v", min, max)
	}
	if reqs[len(reqs)-1].SeqNum.ToInt().Cmp(max) < 0 {
		return false, nil, fmt.Errorf("do not have all the messages in report, our max %v reports max %v", reqs[len(reqs)-1].SeqNum, max)
	}
	if r.isStale(min) {
		return false, nil, nil
	}

	report, err := r.buildReport(min, max)
	if err != nil {
		return false, nil, err
	}
	return true, report, nil
}

func (r RelayReportingPlugin) isStale(minSeqNum *big.Int) bool {
	lastReport, err := r.offRamp.GetLastReport(nil)
	if err != nil {
		// Assume its a transient issue getting the last report
		// Will try again on the next round
		return true
	}
	// If the last report onchain has a lower bound
	// strictly greater than this minSeqNum, then this minSeqNum
	// is stale.
	return lastReport.MinSequenceNumber.Cmp(minSeqNum) > 0
}

func (r RelayReportingPlugin) buildReport(min *big.Int, max *big.Int) ([]byte, error) {
	reqs, err := r.orm.Requests(r.sourceChainId, r.destChainId, min, max, "", nil, nil)
	if err != nil {
		return nil, err
	}
	// Take all these request and produce a merkle root of them
	var leaves [][]byte
	for _, req := range reqs {
		leaves = append(leaves, req.Raw)
	}

	// Note Index doesn't matter, we just want the root
	root, _ := GenerateMerkleProof(32, leaves, 0)
	report, err := EncodeRelayReport(&single_token_offramp.CCIPRelayReport{
		MerkleRoot:        root,
		MinSequenceNumber: min,
		MaxSequenceNumber: max,
	})
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (r RelayReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeRelayReport(report)
	if err != nil {
		return false, nil
	}
	// Note its ok to leave the unstarted requests behind, since the
	// Observe is always based on the last reports onchain min seq num.
	if r.isStale(parsedReport.MinSequenceNumber) {
		return false, nil
	}
	// Any timed out requests should be set back to RequestStatusExecutionPending so their execution can be retried in a subsequent report.
	if err = r.orm.ResetExpiredRequests(r.sourceChainId, r.destChainId, RELAY_MAX_INFLIGHT_TIME_SECONDS, RequestStatusRelayPending, RequestStatusUnstarted); err != nil {
		// Ok to continue here, we'll try to reset them again on the next round.
		r.l.Errorw("unable to reset expired requests", "err", err)
	}
	// Marking new requests as pending/in-flight
	err = r.orm.UpdateRequestStatus(r.sourceChainId, r.destChainId, parsedReport.MinSequenceNumber, parsedReport.MaxSequenceNumber, RequestStatusRelayPending)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (r RelayReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeRelayReport(report)
	if err != nil {
		return false, nil
	}
	// If report is not stale we transmit.
	// When the relayTransmitter enqueues the tx for bptxm,
	// we mark it as fulfilled, effectively removing it from the set of inflight messages.
	return !r.isStale(parsedReport.MinSequenceNumber), nil
}

func (r RelayReportingPlugin) Start() error {
	return nil
}

func (r RelayReportingPlugin) Close() error {
	return nil
}
                                                                                                     core/services/ccip/log_listener.go                                                                  000644  000765  000024  00000027340 14165346401 020255  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"fmt"
	"math/big"
	"reflect"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/log"
	"github.com/smartcontractkit/chainlink/core/services/postgres"
	"github.com/smartcontractkit/chainlink/core/utils"
	"github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	"gorm.io/gorm"
)

var (
	_ log.Listener = &LogListener{}
	_ job.Service  = &LogListener{}
)

type LogListener struct {
	utils.StartStopOnce

	logger                     logger.Logger
	sourceChainLogBroadcaster  log.Broadcaster
	destChainLogBroadcaster    log.Broadcaster
	singleTokenOnRamp          *single_token_onramp.SingleTokenOnRamp
	singleTokenOffRamp         *single_token_offramp.SingleTokenOffRamp
	sourceChainId, destChainId *big.Int
	// this can get overwritten by on-chain changes but doesn't need mutexes
	// because this is a single goroutine service.
	offchainConfig OffchainConfig
	jobID          int32

	unsubscribeLogsOnRamp  func()
	unsubscribeLogsOffRamp func()

	db  *gorm.DB
	orm ORM

	wgShutdown sync.WaitGroup
	mbLogs     *utils.Mailbox
	chStop     chan struct{}
}

func NewLogListener(
	l logger.Logger,
	sourceChainLogBroadcaster log.Broadcaster,
	destChainLogBroadcaster log.Broadcaster,
	singleTokenOnRamp *single_token_onramp.SingleTokenOnRamp,
	singleTokenOffRamp *single_token_offramp.SingleTokenOffRamp,
	offchainConfig OffchainConfig,
	db *gorm.DB,
	jobID int32,
) *LogListener {
	return &LogListener{
		logger:                    l,
		sourceChainLogBroadcaster: sourceChainLogBroadcaster,
		destChainLogBroadcaster:   destChainLogBroadcaster,
		jobID:                     jobID,
		db:                        db,
		orm:                       NewORM(postgres.UnwrapGormDB(db)),
		singleTokenOnRamp:         singleTokenOnRamp,
		singleTokenOffRamp:        singleTokenOffRamp,
		offchainConfig:            offchainConfig,
		mbLogs:                    utils.NewMailbox(10000),
		chStop:                    make(chan struct{}),
	}
}

// Start complies with job.Service
func (l *LogListener) Start() error {
	return l.StartOnce("CCIP_LogListener", func() error {
		sourceChainId, err := l.singleTokenOnRamp.CHAINID(nil)
		if err != nil {
			return err
		}
		destChainId, err := l.singleTokenOffRamp.CHAINID(nil)
		if err != nil {
			return err
		}
		l.sourceChainId = sourceChainId
		l.destChainId = destChainId
		l.subscribeSourceChainLogBroadcaster()
		l.subscribeDestChainLogBroadcaster()
		l.wgShutdown.Add(1)
		l.logger.Infow("CCIP_LogListener: Starting", "onRamp", l.singleTokenOnRamp.Address(), "offRamp", l.singleTokenOffRamp.Address())
		go l.run()

		return nil
	})
}

func (l *LogListener) subscribeSourceChainLogBroadcaster() {
	l.unsubscribeLogsOnRamp = l.sourceChainLogBroadcaster.Register(l, log.ListenerOpts{
		Contract: l.singleTokenOnRamp.Address(),
		LogsWithTopics: map[common.Hash][][]log.Topic{
			// Both relayer and executor save to db
			single_token_onramp.SingleTokenOnRampCrossChainSendRequested{}.Topic(): {},
		},
		ParseLog:         l.singleTokenOnRamp.ParseLog,
		NumConfirmations: uint64(l.offchainConfig.SourceIncomingConfirmations),
	})
}

func (l *LogListener) subscribeDestChainLogBroadcaster() {
	l.unsubscribeLogsOffRamp = l.destChainLogBroadcaster.Register(l, log.ListenerOpts{
		Contract: l.singleTokenOffRamp.Address(),
		LogsWithTopics: map[common.Hash][][]log.Topic{
			// Both relayer and executor mark as report_confirmed state
			single_token_offramp.SingleTokenOffRampReportAccepted{}.Topic(): {},
			// Both relayer and executor mark as execution_confirmed state
			single_token_offramp.SingleTokenOffRampCrossChainMessageExecuted{}.Topic(): {},
			// The offramp listens to config changed
			single_token_offramp.SingleTokenOffRampConfigSet{}.Topic(): {},
		},
		ParseLog:         l.singleTokenOffRamp.ParseLog,
		NumConfirmations: uint64(l.offchainConfig.DestIncomingConfirmations),
	})
}

// Close complies with job.Service
func (l *LogListener) Close() error {
	return l.StopOnce("CCIP_LogListener", func() error {
		close(l.chStop)
		l.wgShutdown.Wait()
		return nil
	})
}

func (l *LogListener) HandleLog(lb log.Broadcast) {
	wasOverCapacity := l.mbLogs.Deliver(lb)
	if wasOverCapacity {
		l.logger.Error("CCIP_LogListener: log mailbox is over capacity - dropped the oldest log")
	}
}

func (l *LogListener) run() {
	for {
		select {
		case <-l.chStop:
			l.unsubscribeLogsOffRamp()
			l.unsubscribeLogsOnRamp()
			l.wgShutdown.Done()
			return
		case <-l.mbLogs.Notify():
			l.handleReceivedLogs()
		}
	}
}

func (l *LogListener) handleReceivedLogs() {
	for {
		i, exists := l.mbLogs.Retrieve()
		if !exists {
			return
		}
		lb, ok := i.(log.Broadcast)
		if !ok {
			panic(errors.Errorf("CCIP_LogListener: invariant violation, expected log.Broadcast but got %T", lb))
		}

		logObj := lb.DecodedLog()
		if logObj == nil || reflect.ValueOf(logObj).IsNil() {
			l.logger.Error("CCIP_LogListener: HandleLog: ignoring nil value")
			return
		}

		// TODO: think about a way to do a single switch
		var logBroadcaster log.Broadcaster
		switch logObj.(type) {
		case *single_token_onramp.SingleTokenOnRampCrossChainSendRequested:
			logBroadcaster = l.sourceChainLogBroadcaster
		case *single_token_offramp.SingleTokenOffRampCrossChainMessageExecuted, *single_token_offramp.SingleTokenOffRampReportAccepted, *single_token_offramp.SingleTokenOffRampConfigSet:
			logBroadcaster = l.destChainLogBroadcaster
		default:
			l.logger.Warnf("CCIP_LogListener: unexpected log type %T", logObj)
		}

		ctx, cancel := postgres.DefaultQueryCtx()
		wasConsumed, err := logBroadcaster.WasAlreadyConsumed(l.db.WithContext(ctx), lb)
		cancel()
		if err != nil {
			l.logger.Errorw("CCIP_LogListener: could not determine if log was already consumed", "error", err)
			return
		} else if wasConsumed {
			return
		}

		switch log := logObj.(type) {
		case *single_token_onramp.SingleTokenOnRampCrossChainSendRequested:
			l.handleCrossChainSendRequested(log, lb)
		case *single_token_offramp.SingleTokenOffRampCrossChainMessageExecuted:
			l.handleCrossChainMessageExecuted(log, lb)
		case *single_token_offramp.SingleTokenOffRampReportAccepted:
			l.handleCrossChainReportRelayed(log, lb)
		case *single_token_offramp.SingleTokenOffRampConfigSet:
			if err := l.updateIncomingConfirmationsConfig(lb.RawLog()); err != nil {
				l.logger.Errorw("could not parse config set", "err", err)
			}
		default:
			l.logger.Warnf("CCIP_LogListener: unexpected log type %T", logObj)
		}
	}
}

func (l *LogListener) updateIncomingConfirmationsConfig(log types.Log) error {
	offrampConfigSet, err := l.singleTokenOffRamp.ParseConfigSet(log)
	if err != nil {
		return err
	}
	contractConfig := ContractConfigFromConfigSetEvent(ConfigSet(*offrampConfigSet))
	publicConfig, err := confighelper.PublicConfigFromContractConfig(false, contractConfig)
	if err != nil {
		return err
	}
	ccipConfig, err := Decode(publicConfig.ReportingPluginConfig)
	if err != nil {
		return err
	}
	if l.offchainConfig.SourceIncomingConfirmations != ccipConfig.SourceIncomingConfirmations {
		l.offchainConfig.SourceIncomingConfirmations = ccipConfig.SourceIncomingConfirmations
		l.unsubscribeLogsOnRamp()
		l.subscribeSourceChainLogBroadcaster()
	}

	if l.offchainConfig.DestIncomingConfirmations != ccipConfig.DestIncomingConfirmations {
		l.offchainConfig.DestIncomingConfirmations = ccipConfig.DestIncomingConfirmations
		l.unsubscribeLogsOffRamp()
		l.subscribeDestChainLogBroadcaster()
	}
	return nil
}

func (l *LogListener) handleCrossChainMessageExecuted(executed *single_token_offramp.SingleTokenOffRampCrossChainMessageExecuted, lb log.Broadcast) {
	l.logger.Infow("CCIP_LogListener: cross chain request executed",
		"seqNum", fmt.Sprintf("%0x", executed.SequenceNumber),
		"jobID", lb.JobID(),
	)
	err := l.orm.UpdateRequestStatus(l.sourceChainId, l.destChainId, executed.SequenceNumber, executed.SequenceNumber, RequestStatusExecutionConfirmed)
	if err != nil {
		// We can replay the logs if needed
		l.logger.Errorw("failed to save CCIP request", "error", err)
		return
	}
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	if err := l.destChainLogBroadcaster.MarkConsumed(l.db.WithContext(ctx), lb); err != nil {
		l.logger.Errorw("CCIP_LogListener: failed mark consumed", "err", err)
	}
}

func (l *LogListener) handleCrossChainReportRelayed(relayed *single_token_offramp.SingleTokenOffRampReportAccepted, lb log.Broadcast) {
	l.logger.Infow("CCIP_LogListener: cross chain report relayed",
		"minSeqNum", fmt.Sprintf("%0x", relayed.Report.MinSequenceNumber),
		"maxSeqNum", fmt.Sprintf("%0x", relayed.Report.MaxSequenceNumber),
		"jobID", lb.JobID(),
	)

	// TODO: should be in the same tx
	err := l.orm.UpdateRequestStatus(l.sourceChainId, l.destChainId, relayed.Report.MinSequenceNumber, relayed.Report.MaxSequenceNumber, RequestStatusRelayConfirmed)
	if err != nil {
		// We can replay the logs if needed
		l.logger.Errorw("failed to save CCIP request", "error", err)
		return
	}
	err = l.orm.SaveRelayReport(RelayReport{
		Root:      relayed.Report.MerkleRoot[:],
		MinSeqNum: *utils.NewBig(relayed.Report.MinSequenceNumber),
		MaxSeqNum: *utils.NewBig(relayed.Report.MaxSequenceNumber),
	})
	if err != nil {
		// We can replay the logs if needed
		l.logger.Errorw("failed to save CCIP report", "error", err)
		return
	}
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	if err := l.destChainLogBroadcaster.MarkConsumed(l.db.WithContext(ctx), lb); err != nil {
		l.logger.Errorw("CCIP_LogListener: failed mark consumed", "err", err)
	}
}

// We assume a bounded Message size which is enforced on-chain,
// TODO: add Message bounds to onramp and include assertion offchain as well.
func (l *LogListener) handleCrossChainSendRequested(request *single_token_onramp.SingleTokenOnRampCrossChainSendRequested, lb log.Broadcast) {
	l.logger.Infow("CCIP_LogListener: cross chain send request received",
		"requestId", fmt.Sprintf("%0x", request.Message.SequenceNumber),
		"sender", request.Message.Sender,
		"receiver", request.Message.Payload.Receiver,
		"sourceChainId", request.Message.SourceChainId,
		"destChainId", request.Message.DestinationChainId,
		"tokens", request.Message.Payload.Tokens,
		"amounts", request.Message.Payload.Amounts,
		"options", request.Message.Payload.Options,
		"jobID", lb.JobID(),
	)

	var tokens []string
	for _, token := range request.Message.Payload.Tokens {
		tokens = append(tokens, token.String())
	}
	var amounts []string
	for _, amount := range request.Message.Payload.Amounts {
		amounts = append(amounts, amount.String())
	}
	err := l.orm.SaveRequest(&Request{
		SeqNum:        *utils.NewBig(request.Message.SequenceNumber),
		SourceChainID: request.Message.SourceChainId.String(),
		DestChainID:   request.Message.DestinationChainId.String(),
		Sender:        request.Message.Sender,
		Receiver:      request.Message.Payload.Receiver,
		Data:          request.Message.Payload.Data,
		Tokens:        tokens,
		Amounts:       amounts,
		Executor:      request.Message.Payload.Executor,
		Options:       request.Message.Payload.Options,
		Raw:           request.Raw.Data,
		Status:        RequestStatusUnstarted,
	})
	if err != nil {
		// We can replay the logs if needed
		l.logger.Errorw("failed to save CCIP request", "error", err)
		return
	}

	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	if err := l.sourceChainLogBroadcaster.MarkConsumed(l.db.WithContext(ctx), lb); err != nil {
		l.logger.Errorw("CCIP_LogListener: failed mark consumed", "err", err)
	}
}

// JobID complies with log.Listener
func (l *LogListener) JobID() int32 {
	return l.jobID
}
                                                                                                                                                                                                                                                                                                core/services/ccip/execution_reporting_plugin.go                                                    000644  000765  000024  00000032753 14165346401 023245  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/utils"

	"github.com/smartcontractkit/chainlink/core/logger"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
)

const (
	EXECUTION_MAX_INFLIGHT_TIME_SECONDS = 180
)

var _ types.ReportingPluginFactory = &ExecutionReportingPluginFactory{}
var _ types.ReportingPlugin = &ExecutionReportingPlugin{}

type Message struct {
	SequenceNumber     *big.Int       `json:"sequenceNumber"`
	SourceChainId      *big.Int       `json:"sourceChainId"`
	DestinationChainId *big.Int       `json:"destinationChainId"`
	Sender             common.Address `json:"sender"`
	Payload            struct {
		Receiver common.Address   `json:"receiver"`
		Data     []uint8          `json:"data"`
		Tokens   []common.Address `json:"tokens"`
		Amounts  []*big.Int       `json:"amounts"`
		Executor common.Address   `json:"executor"`
		Options  []uint8          `json:"options"`
	} `json:"payload"`
}

type ExecutableMessage struct {
	Proof   [][32]byte `json:"proof"`
	Message Message    `json:"message"`
	Index   *big.Int   `json:"index"`
}

type ExecutableMessages []ExecutableMessage

func (ems ExecutableMessages) SeqNums() (nums []*big.Int) {
	for i := range ems {
		nums = append(nums, ems[i].Message.SequenceNumber)
	}
	return
}

// ExecutionObservation Note there can be gaps in this range of sequence numbers,
// indicative of some messages being non-DON executed.
type ExecutionObservation struct {
	MinSeqNum utils.Big `json:"minSeqNum"`
	MaxSeqNum utils.Big `json:"maxSeqNum"`
}

func makeExecutionReportArgs() abi.Arguments {
	mustType := func(ts string, components []abi.ArgumentMarshaling) abi.Type {
		ty, _ := abi.NewType(ts, "", components)
		return ty
	}
	return []abi.Argument{
		{
			Name: "executableMessages",
			Type: mustType("tuple[]", []abi.ArgumentMarshaling{
				{
					Name: "Proof",
					Type: "bytes32[]",
				},
				{
					Name: "Message",
					Type: "tuple",
					Components: []abi.ArgumentMarshaling{
						{
							Name: "sequenceNumber",
							Type: "uint256",
						},
						{
							Name: "sourceChainId",
							Type: "uint256",
						},
						{
							Name: "destinationChainId",
							Type: "uint256",
						},
						{
							Name: "sender",
							Type: "address",
						},
						{
							Name: "payload",
							Type: "tuple",
							Components: []abi.ArgumentMarshaling{
								{
									Name: "receiver",
									Type: "address",
								},
								{
									Name: "data",
									Type: "bytes",
								},
								{
									Name: "tokens",
									Type: "address[]",
								},
								{
									Name: "amounts",
									Type: "uint256[]",
								},
								{
									Name: "executor",
									Type: "address",
								},
								{
									Name: "options",
									Type: "bytes",
								},
							},
						},
					},
				},
				{
					Name: "Index",
					Type: "uint256",
				},
			}),
		},
	}
}

func EncodeExecutionReport(ems []ExecutableMessage) (types.Report, error) {
	report, err := makeExecutionReportArgs().PackValues([]interface{}{ems})
	if err != nil {
		return nil, err
	}
	return report, nil
}

func DecodeExecutionReport(report types.Report) ([]ExecutableMessage, error) {
	unpacked, err := makeExecutionReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, nil
	}

	// Must be anonymous struct here
	msgs, ok := unpacked[0].([]struct {
		Proof   [][32]uint8 `json:"Proof"`
		Message struct {
			SequenceNumber     *big.Int       `json:"sequenceNumber"`
			SourceChainId      *big.Int       `json:"sourceChainId"`
			DestinationChainId *big.Int       `json:"destinationChainId"`
			Sender             common.Address `json:"sender"`
			Payload            struct {
				Receiver common.Address   `json:"receiver"`
				Data     []uint8          `json:"data"`
				Tokens   []common.Address `json:"tokens"`
				Amounts  []*big.Int       `json:"amounts"`
				Executor common.Address   `json:"executor"`
				Options  []uint8          `json:"options"`
			} `json:"payload"`
		} `json:"Message"`
		Index *big.Int `json:"Index"`
	})
	if !ok {
		return nil, fmt.Errorf("got %T", unpacked[0])
	}
	var ems []ExecutableMessage
	for _, emi := range msgs {
		ems = append(ems, ExecutableMessage{
			Proof:   emi.Proof,
			Message: emi.Message,
			Index:   emi.Index,
		})
	}
	return ems, nil
}

//go:generate mockery --name OffRampLastReporter --output ./mocks/lastreporter --case=underscore
type OffRampLastReporter interface {
	GetLastReport(opts *bind.CallOpts) (single_token_offramp.CCIPRelayReport, error)
}

type ExecutionReportingPluginFactory struct {
	l            logger.Logger
	orm          ORM
	source, dest *big.Int
	lastReporter OffRampLastReporter
	executor     common.Address
}

func NewExecutionReportingPluginFactory(l logger.Logger, orm ORM, source, dest *big.Int, executor common.Address, lastReporter OffRampLastReporter) types.ReportingPluginFactory {
	return &ExecutionReportingPluginFactory{l: l, orm: orm, source: source, dest: dest, executor: executor, lastReporter: lastReporter}
}

func (rf *ExecutionReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	return ExecutionReportingPlugin{rf.l, config.F, rf.orm, rf.source, rf.dest, rf.executor, rf.lastReporter}, types.ReportingPluginInfo{
		Name:              "CCIPExecution",
		UniqueReports:     true,
		MaxQueryLen:       0,      // We do not use the query phase.
		MaxObservationLen: 100000, // TODO
		MaxReportLen:      100000, // TODO
	}, nil
}

type ExecutionReportingPlugin struct {
	l             logger.Logger
	F             int
	orm           ORM
	sourceChainId *big.Int
	destChainId   *big.Int
	executor      common.Address
	// We also use the offramp for defensive checks
	lastReporter OffRampLastReporter
}

func (r ExecutionReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	return types.Query{}, nil
}

func (r ExecutionReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	// We want to execute any messages which satisfy the following:
	// 1. Have the executor field set to the DONs message executor contract
	// 2. There exists a confirmed relay report containing its sequence number, i.e. it's status is RequestStatusRelayConfirmed
	reqs, err := r.orm.Requests(r.sourceChainId, r.destChainId, nil, nil, RequestStatusRelayConfirmed, &r.executor, nil)
	if err != nil {
		return nil, err
	}
	// No request to process
	// Return an empty observation
	// which should not result in a report generated.
	if len(reqs) == 0 {
		return nil, fmt.Errorf("no requests for oracle execution")
	}
	// Double check the latest sequence number onchain is >= our max relayed seq num
	lr, err := r.lastReporter.GetLastReport(nil)
	if err != nil {
		return nil, err
	}
	if reqs[len(reqs)-1].SeqNum.ToInt().Cmp(lr.MaxSequenceNumber) > 0 {
		return nil, fmt.Errorf("invariant violated, mismatch between relay_confirmed requests and last report")
	}
	b, err := json.Marshal(&ExecutionObservation{
		MinSeqNum: reqs[0].SeqNum,
		MaxSeqNum: reqs[len(reqs)-1].SeqNum,
	})
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r ExecutionReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	var nonEmptyObservations []ExecutionObservation
	for _, ao := range observations {
		var ob ExecutionObservation
		err := json.Unmarshal(ao.Observation, &ob)
		if err != nil {
			r.l.Errorw("unmarshallable observation", "ao", ao.Observation, "err", err)
			continue
		}
		nonEmptyObservations = append(nonEmptyObservations, ob)
	}
	// Need at least F+1 observations
	if len(nonEmptyObservations) <= r.F {
		return false, nil, nil
	}
	// We have at least F+1 valid observations
	// Extract the min and max
	sort.Slice(nonEmptyObservations, func(i, j int) bool {
		return nonEmptyObservations[i].MinSeqNum.ToInt().Cmp(nonEmptyObservations[j].MinSeqNum.ToInt()) < 0
	})
	min := nonEmptyObservations[r.F].MinSeqNum.ToInt()
	sort.Slice(nonEmptyObservations, func(i, j int) bool {
		return nonEmptyObservations[i].MaxSeqNum.ToInt().Cmp(nonEmptyObservations[j].MaxSeqNum.ToInt()) < 0
	})
	max := nonEmptyObservations[r.F].MaxSeqNum.ToInt()
	if max.Cmp(min) < 0 {
		return false, nil, errors.New("max seq num smaller than min")
	}
	reqs, err := r.orm.Requests(r.sourceChainId, r.destChainId, min, max, RequestStatusRelayConfirmed, &r.executor, nil)
	if err != nil {
		return false, nil, err
	}
	// Cannot construct a report for which we haven't seen all the messages.
	if len(reqs) == 0 {
		return false, nil, fmt.Errorf("do not have all the messages in report, have zero messages, report has min %v max %v", min, max)
	}
	lr, err := r.lastReporter.GetLastReport(nil)
	if err != nil {
		return false, nil, err
	}
	if reqs[len(reqs)-1].SeqNum.ToInt().Cmp(lr.MaxSequenceNumber) > 0 {
		return false, nil, fmt.Errorf("invariant violated, mismatch between relay_confirmed requests (max %v) and last report (max %v)", reqs[len(reqs)-1].SeqNum, lr.MaxSequenceNumber)
	}
	report, err := r.buildReport(reqs)
	if err != nil {
		return false, nil, err
	}
	return true, report, nil
}

// For each message in the given range of sequence numbers (with potential holes):
// 1. Lookup the report associated with that sequence number
// 2. Generate a merkle proof that the message was in that report
// 3. Encode those proofs and messages into a report for the executor contract
// TODO: We may want to combine these queries for performance, hold off
// until we decide whether we move forward with batch proving.
func (r ExecutionReportingPlugin) buildReport(reqs []*Request) ([]byte, error) {
	var executable []ExecutableMessage
	for _, req := range reqs {
		// Look up all the messages that are in the same report
		// as this one (even externally executed ones), generate a Proof and double-check the root checks out.
		rep, err2 := r.orm.RelayReport(req.SeqNum.ToInt())
		if err2 != nil {
			r.l.Errorw("could not find relay report for request", "err", err2, "seq num", req.SeqNum.String())
			continue
		}
		allReqsInReport, err3 := r.orm.Requests(r.sourceChainId, r.destChainId, rep.MinSeqNum.ToInt(), rep.MaxSeqNum.ToInt(), "", nil, nil)
		if err3 != nil {
			continue
		}
		var leaves [][]byte
		for _, reqInReport := range allReqsInReport {
			leaves = append(leaves, reqInReport.Raw)
		}
		index := big.NewInt(0).Sub(req.SeqNum.ToInt(), rep.MinSeqNum.ToInt())
		root, proof := GenerateMerkleProof(32, leaves, int(index.Int64()))
		if !bytes.Equal(root[:], rep.Root[:]) {
			continue
		}
		executable = append(executable, ExecutableMessage{
			Proof:   proof.PathForExecute(),
			Message: req.ToMessage(),
			Index:   proof.Index(),
		})
	}

	report, err := EncodeExecutionReport(executable)
	if err != nil {
		return nil, err
	}

	return report, nil
}

func (r ExecutionReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	ems, err := DecodeExecutionReport(report)
	if err != nil {
		return false, nil
	}
	// If the report is stale, we do not accept it.
	stale, err := r.isStale(ems[0].Message.SequenceNumber)
	if err != nil {
		return !stale, err
	}
	if stale {
		return false, err
	}
	// Any timed out requests should be set back to RequestStatusExecutionPending so their execution can be retried in a subsequent report.
	if err = r.orm.ResetExpiredRequests(r.sourceChainId, r.destChainId, EXECUTION_MAX_INFLIGHT_TIME_SECONDS, RequestStatusExecutionPending, RequestStatusRelayConfirmed); err != nil {
		// Ok to continue here, we'll try to reset them again on the next round.
		r.l.Errorw("unable to reset expired requests", "err", err)
	}
	if err := r.orm.UpdateRequestSetStatus(r.sourceChainId, r.destChainId, ExecutableMessages(ems).SeqNums(), RequestStatusExecutionPending); err != nil {
		return false, err
	}
	return true, nil
}

func (r ExecutionReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeExecutionReport(report)
	if err != nil {
		return false, nil
	}
	// If report is not stale we transmit.
	// When the executeTransmitter enqueues the tx for bptxm,
	// we mark it as execution_sent, removing it from the set of inflight messages.
	stale, err := r.isStale(parsedReport[0].Message.SequenceNumber)
	return !stale, err
}

func (r ExecutionReportingPlugin) isStale(min *big.Int) (bool, error) {
	// If the first message is executed already, this execution report is stale.
	req, err := r.orm.Requests(r.sourceChainId, r.destChainId, min, min, "", nil, nil)
	if err != nil {
		// if we can't find the request, assume transient db issue
		// and wait until the next OCR2 round (don't submit)
		return true, err
	}
	if len(req) != 1 {
		// If we don't have the request at all, this likely means we never had the request to begin with
		// (say our eth subscription is down) and we want to let other oracles continue the protocol.
		return false, errors.New("could not find first message in execution report")
	}
	return req[0].Status == RequestStatusExecutionConfirmed, nil
}

func (r ExecutionReportingPlugin) Start() error {
	return nil
}

func (r ExecutionReportingPlugin) Close() error {
	return nil
}
                     core/services/ccip/delegate_bootstrap.go                                                            000644  000765  000024  00000007130 14165346401 021431  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // TODO: I think we might be able to make the bootstrap job type just generic for all genocr jobs?
package ccip

import (
	"context"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/core/services/postgres"
	ocrcommontypes "github.com/smartcontractkit/libocr/commontypes"
	ocr "github.com/smartcontractkit/libocr/offchainreporting2"
	"github.com/smartcontractkit/libocr/offchainreporting2/chains/evmutil"
	"gorm.io/gorm"
)

type DelegateBootstrap struct {
	bootstrappers []ocrcommontypes.BootstrapperLocator
	db            *gorm.DB
	jobORM        job.ORM
	orm           ORM
	chainSet      evm.ChainSet
	peerWrapper   *ocrcommon.SingletonPeerWrapper
}

// TODO: Register this delegate behind a FF
func NewDelegateBootstrap(
	db *gorm.DB,
	jobORM job.ORM,
	chainSet evm.ChainSet,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
) *DelegateBootstrap {
	return &DelegateBootstrap{
		db:          db,
		jobORM:      jobORM,
		orm:         NewORM(postgres.UnwrapGormDB(db)),
		chainSet:    chainSet,
		peerWrapper: peerWrapper,
	}
}

func (d DelegateBootstrap) JobType() job.Type {
	return job.CCIPBootstrap
}

func (d DelegateBootstrap) ServicesForSpec(jb job.Job) ([]job.Service, error) {
	if jb.CCIPBootstrapSpec == nil {
		return nil, errors.New("no bootstrap job specified")
	}
	l := logger.Default.With(
		"jobID", jb.ID,
		"externalJobID", jb.ExternalJobID,
		"coordinatorAddress", jb.CCIPBootstrapSpec.ContractAddress,
	)

	c, err := d.chainSet.Get(jb.CCIPBootstrapSpec.EVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open chain")
	}
	// Bootstrap could either be an offramp or an executor, should work in both cases
	offRamp, err := single_token_offramp.NewSingleTokenOffRamp(jb.CCIPBootstrapSpec.ContractAddress.Address(), c.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not instantiate NewOffchainAggregator")
	}

	gormdb, errdb := d.db.DB()
	if errdb != nil {
		return nil, errors.Wrap(errdb, "unable to open sql db")
	}
	ocrdb := NewDB(gormdb, jb.CCIPBootstrapSpec.ContractAddress.Address())
	contractTracker := NewCCIPContractTracker(
		offrampTracker{offRamp},
		c.Client(),
		c.LogBroadcaster(),
		jb.ID,
		logger.Default,
		d.db,
		c,
		c.HeadBroadcaster(),
	)
	ocrLogger := logger.NewOCRWrapper(l, true, func(msg string) {
		d.jobORM.RecordError(context.Background(), jb.ID, msg)
	})
	offchainConfigDigester := evmutil.EVMOffchainConfigDigester{
		ChainID:         maybeRemapChainID(c.Config().ChainID()).Uint64(),
		ContractAddress: jb.CCIPBootstrapSpec.ContractAddress.Address(),
	}
	bootstrapNode, err := ocr.NewBootstrapper(ocr.BootstrapperArgs{
		BootstrapperFactory:   d.peerWrapper.Peer2,
		ContractConfigTracker: contractTracker,
		Database:              ocrdb,
		LocalConfig: computeLocalConfig(c.Config(), c.Config().Dev(),
			jb.CCIPBootstrapSpec.BlockchainTimeout.Duration(),
			jb.CCIPBootstrapSpec.ContractConfigConfirmations, jb.CCIPBootstrapSpec.ContractConfigTrackerPollInterval.Duration()),
		Logger:                 ocrLogger,
		MonitoringEndpoint:     nil, // TODO
		OffchainConfigDigester: offchainConfigDigester,
	})
	if err != nil {
		return nil, err
	}
	return []job.Service{contractTracker, bootstrapNode}, nil
}

func (d DelegateBootstrap) AfterJobCreated(spec job.Job) {
}

func (d DelegateBootstrap) BeforeJobDeleted(spec job.Job) {
}
                                                                                                                                                                                                                                                                                                                                                                                                                                        core/services/ccip/validate.go                                                                      000644  000765  000024  00000002337 14165346401 017357  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services/job"
)

func ValidatedCCIPSpec(tomlString string) (job.Job, error) {
	var jb = job.Job{}
	tree, err := toml.Load(tomlString)
	if err != nil {
		return jb, err
	}
	err = tree.Unmarshal(&jb)
	if err != nil {
		return jb, err
	}
	switch jb.Type {
	case job.CCIPRelay:
		var spec job.CCIPRelaySpec
		err = tree.Unmarshal(&spec)
		if err != nil {
			return jb, err
		}
		jb.CCIPRelaySpec = &spec
	case job.CCIPExecution:
		var spec job.CCIPExecutionSpec
		err = tree.Unmarshal(&spec)
		if err != nil {
			return jb, err
		}
		jb.CCIPExecutionSpec = &spec
	default:
		return jb, errors.Errorf("unsupported type %s", jb.Type)
	}

	return jb, nil
}

func ValidatedCCIPBootstrapSpec(tomlString string) (job.Job, error) {
	var jb = job.Job{}
	tree, err := toml.Load(tomlString)
	if err != nil {
		return jb, err
	}
	err = tree.Unmarshal(&jb)
	if err != nil {
		return jb, err
	}
	var spec job.CCIPBootstrapSpec
	err = tree.Unmarshal(&spec)
	if err != nil {
		return jb, err
	}
	jb.CCIPBootstrapSpec = &spec

	if jb.Type != job.CCIPBootstrap {
		return jb, errors.Errorf("unsupported type %s", jb.Type)
	}
	return jb, nil
}
                                                                                                                                                                                                                                                                                                 core/services/ccip/log_listener_test.go                                                             000644  000765  000024  00000027634 14165357743 021335  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"bytes"
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/services/ccip/abihelpers"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/lib/pq"
	"github.com/onsi/gomega"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/lock_unlock_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/eth"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/log"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/core/testdata/testspecs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type lc struct {
}

func (l lc) BlockBackfillDepth() uint64 {
	return 1
}

func (l lc) BlockBackfillSkip() bool {
	return false
}

func (l lc) EvmFinalityDepth() uint32 {
	return 50
}

func (l lc) EvmLogBackfillBatchSize() uint32 {
	return 1
}

func TestLogListener_SavesRequests(t *testing.T) {
	// Deploy contract
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	user, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		user.From: {Balance: big.NewInt(1000000000000000000)}},
		ethconfig.Defaults.Miner.GasCeil)
	linkTokenAddress, _, linkToken, err := link_token_interface.DeployLinkToken(user, backend)
	require.NoError(t, err)
	poolAddress, _, pool, err := lock_unlock_pool.DeployLockUnlockPool(user, backend, linkTokenAddress)
	require.NoError(t, err)
	afn := deployAfn(t, user, backend)

	onRampAddress, _, _, err := single_token_onramp.DeploySingleTokenOnRamp(
		user,               // user
		backend,            // client
		big.NewInt(2),      // source chain id
		linkTokenAddress,   // source token
		poolAddress,        // source pool
		big.NewInt(1),      // dest chain id
		linkTokenAddress,   // remoteToken
		[]common.Address{}, // allow list
		false,              // enableAllowList
		big.NewInt(1),      // token bucket rate
		big.NewInt(1000),   // token bucket capacity
		afn,                // AFN
		// 86400 seconds = one day
		big.NewInt(86400), //maxTimeWithoutAFNSignal
	)
	require.NoError(t, err)
	onRamp, err := single_token_onramp.NewSingleTokenOnRamp(onRampAddress, backend)
	require.NoError(t, err)
	_, err = pool.SetOnRamp(user, onRampAddress, true)
	require.NoError(t, err)
	_, err = linkToken.Approve(user, poolAddress, big.NewInt(100))
	require.NoError(t, err)
	offRampAddress, _, _, err := single_token_offramp.DeploySingleTokenOffRamp(
		user,             // user
		backend,          // client
		big.NewInt(1),    // source chain id
		big.NewInt(2),    // dest chain id
		linkTokenAddress, // link token address
		poolAddress,      // dest pool address
		big.NewInt(1),    // token bucket rate
		big.NewInt(1000), // token bucket capacity
		afn,              // AFN address
		// 86400 seconds = one day
		big.NewInt(86400), // max timeout without AFN signal
		big.NewInt(0),     // execution delay in seconds
	)
	require.NoError(t, err)
	offRamp, err := single_token_offramp.NewSingleTokenOffRamp(offRampAddress, backend)
	require.NoError(t, err)
	backend.Commit()

	// Start the log broadcaster/log listener
	// and add a CCIP job.
	db := pgtest.NewGormDB(t)
	ethClient := eth.NewClientFromSim(backend, big.NewInt(1337))
	lggr := logger.TestLogger(t)
	lorm := log.NewORM(db, *big.NewInt(1337))
	r, err := lorm.FindConsumedLogs(0, 100)
	require.NoError(t, err)
	t.Log(r)
	lb := log.NewBroadcaster(lorm, ethClient, lc{}, lggr, nil)
	require.NoError(t, lb.Start())
	jobORM := job.NewORM(db, nil, pipeline.NewORM(db), nil, lggr)
	ccipSpec, err := ValidatedCCIPSpec(testspecs.GenerateCCIPSpec(testspecs.CCIPSpecParams{}).Toml())
	require.NoError(t, err)
	jb, err := jobORM.CreateJob(context.Background(), &ccipSpec, ccipSpec.Pipeline)
	require.NoError(t, err)
	ccipConfig := OffchainConfig{
		SourceIncomingConfirmations: 0,
		DestIncomingConfirmations:   0,
	}
	logListener := NewLogListener(logger.Default, lb, lb, onRamp, offRamp, ccipConfig, db, jb.ID)
	t.Log("Ramp address", onRampAddress, onRamp.Address())
	require.NoError(t, logListener.Start())

	// Update the ccip config on chain and assert that the log listener uses the new config values
	newCcipConfig := OffchainConfig{
		SourceIncomingConfirmations: 1,
		DestIncomingConfirmations:   5,
	}
	updateOffchainConfig(t, newCcipConfig, offRamp, user)
	backend.Commit()

	// Send blocks until that request is saved.
	head, err := backend.HeaderByNumber(context.Background(), nil)
	require.NoError(t, err)
	startHead := head.Number.Int64()
	var reqs []*Request
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		lb.OnNewLongestChain(context.Background(), eth.Head{Hash: head.Hash(), Number: startHead})
		startHead++
		reqs, err = logListener.orm.Requests(big.NewInt(2), big.NewInt(1), big.NewInt(0), nil, RequestStatusUnstarted, nil, nil)
		require.NoError(t, err)
		t.Logf("log %+v\n", reqs)
		return logListener.offchainConfig.DestIncomingConfirmations == newCcipConfig.DestIncomingConfirmations &&
			logListener.offchainConfig.SourceIncomingConfirmations == newCcipConfig.SourceIncomingConfirmations
	}, 3*time.Second, 100*time.Millisecond).Should(gomega.BeTrue())

	//Send a request.
	executor := common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4")
	msg := single_token_onramp.CCIPMessagePayload{
		Receiver: linkTokenAddress,
		Data:     []byte("hello xchain world"),
		Tokens:   []common.Address{linkTokenAddress},
		Amounts:  []*big.Int{big.NewInt(100)},
		Executor: executor,
		Options:  nil,
	}
	_, err = onRamp.RequestCrossChainSend(user, msg)
	require.NoError(t, err)
	backend.Commit()

	// Send blocks until that request is saved.
	head, err = backend.HeaderByNumber(context.Background(), nil)
	require.NoError(t, err)
	startHead = head.Number.Int64()
	reqs = []*Request{}
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		lb.OnNewLongestChain(context.Background(), eth.Head{Hash: head.Hash(), Number: startHead})
		startHead++
		reqs, err = logListener.orm.Requests(big.NewInt(2), big.NewInt(1), big.NewInt(0), nil, RequestStatusUnstarted, nil, nil)
		require.NoError(t, err)
		t.Logf("log %+v\n", reqs)
		return len(reqs) == 1
	}, 3*time.Second, 100*time.Millisecond).Should(gomega.BeTrue())

	// Assert the xchain request was saved correctly.
	assert.Equal(t, "100", reqs[0].Amounts[0])
	assert.Equal(t, msg.Data, reqs[0].Data)
	assert.Equal(t, pq.StringArray{linkTokenAddress.String()}, reqs[0].Tokens)
	assert.Equal(t, msg.Receiver, reqs[0].Receiver)
	assert.Equal(t, msg.Executor.String(), reqs[0].Executor.String())
	assert.Equal(t, []byte{}, reqs[0].Options)
	// We expect the raw request bytes to be the abi.encoded CCIP Message
	b, err := abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{single_token_onramp.CCIPMessage{
		SequenceNumber:     big.NewInt(1),
		SourceChainId:      big.NewInt(2),
		DestinationChainId: big.NewInt(1),
		Sender:             user.From,
		Payload:            msg,
	}})
	require.NoError(t, err)
	require.True(t, bytes.Equal(reqs[0].Raw, b))
	// Round trip should be the same bytes
	cmsg, err := abihelpers.DecodeCCIPMessage(b)
	require.NoError(t, err)
	b2, err := abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{cmsg})
	require.NoError(t, err)
	require.True(t, bytes.Equal(b2, b))

	require.NoError(t, lb.Close())
	require.NoError(t, logListener.Close())
	require.NoError(t, jobORM.DeleteJob(context.Background(), jb.ID))
}

func updateOffchainConfig(t *testing.T, reportingPluginConfig OffchainConfig, offRamp *single_token_offramp.SingleTokenOffRamp, user *bind.TransactOpts) {
	encoded, err := reportingPluginConfig.Encode()
	require.NoError(t, err)

	var oracles = []confighelper2.OracleIdentityExtra{
		{
			// Node 1
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  common.HexToAddress("0xf4e7b2426718b11d8df7008d688d48c8926768d3").Bytes(),
				TransmitAccount:   ocrtypes2.Account("0x016D97857a21A501a0C10b526011516000cE4586"),
				OffchainPublicKey: hexutil.MustDecode("0x510bdd47650e70f3006b24261944d5c3685bc1b8194e5e209beea02916189952"),
				PeerID:            "12D3KooWENNxGhdSx7wXWRXcrZ2uKrY8FEagUCntS6Jw55gXqrTX",
			},
			ConfigEncryptionPublicKey: stringTo32Bytes("0xb2b25ce373a833e3fa7f23538a6ace837673e4ef890db7f7e02830e8d5b6d009"),
		},
		{
			// Node 2
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  common.HexToAddress("0x33a96c0976DD8c10Cc3e9709Ed25f2CF7d7d970E").Bytes(),
				TransmitAccount:   ocrtypes2.Account("0xcca943C692b27b47a43cB532b2354591BD8a7E9b"),
				OffchainPublicKey: hexutil.MustDecode("0x705cec8e7df7ca42fb8465a60e68ff4e02afd90e17dfef2b01e1166c8dd0cb96"),
				PeerID:            "12D3KooWJtEHwtgkC96umAg2C3Gc8oWpqqT81z6RQXEhkFZK1P21",
			},
			ConfigEncryptionPublicKey: stringTo32Bytes("0x0661dc7f751df3c97b1303a78d310d09d7cf32c24df5404136c6275a0385d172"),
		},
		{
			// Node 3
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  common.HexToAddress("0x19dec24A8748c117b102Bb29418F36c45E8C94f1").Bytes(),
				TransmitAccount:   ocrtypes2.Account("0x2fD8930F52bD73Eb01C78b375E8449D6c107170c"),
				OffchainPublicKey: hexutil.MustDecode("0xccc929da9f3185f018c357a14d427cb9c982e981e3d4e20c391cbfb13d9fbb81"),
				PeerID:            "12D3KooWEC7dxiVkSRTCbFV72R4MSn2EZhDtnH7sH5mtYZifzqCW",
			},
			ConfigEncryptionPublicKey: stringTo32Bytes("0x3c21f181098f39d854cc77a4189b3a56b37bee7fec2386abe04e1e36b9177d15"),
		},
		{
			// Node 4
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  common.HexToAddress("0x257ca0ff00204861bbeb626d70a733ece8dc71fa").Bytes(),
				TransmitAccount:   ocrtypes2.Account("0x338820995b4772fAafCEd3bF56824D4b7a6996De"),
				OffchainPublicKey: hexutil.MustDecode("0x2b6fe2d95b217e93da7192bc495828bd5a7c8fc5e7deee919a21c19bc4b951c7"),
				PeerID:            "12D3KooWAyafDntpPKSnGeT4ybu7onfDtUAe54LNzaJGKGnfBx6c",
			},
			ConfigEncryptionPublicKey: stringTo32Bytes("0xd14d160383b80e13dff1130fcdaed3afd54eabbb1f1c1136d3ea6b77e802744b"),
		},
	}
	// Change the offramp config
	signers, transmitters, threshold, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper2.ContractSetConfigArgs(
		2*time.Second,        // deltaProgress
		1*time.Second,        // deltaResend
		1*time.Second,        // deltaRound
		500*time.Millisecond, // deltaGrace
		2*time.Second,        // deltaStage
		3,
		[]int{1, 1, 1, 1},
		oracles,
		encoded,
		50*time.Millisecond,
		50*time.Millisecond,
		50*time.Millisecond,
		50*time.Millisecond,
		50*time.Millisecond,
		1, // faults
		nil,
	)
	_, err = offRamp.SetConfig(user, signers, transmitters, threshold, onchainConfig, offchainConfigVersion, offchainConfig)
	require.NoError(t, err)
}

func stringTo32Bytes(s string) [32]byte {
	var b [32]byte
	copy(b[:], hexutil.MustDecode(s))
	return b
}

func deployAfn(t *testing.T, user *bind.TransactOpts, chain *backends.SimulatedBackend) common.Address {
	afnSourceAddress, _, _, err := afn_contract.DeployAFNContract(
		user,
		chain,
		[]common.Address{user.From},
		[]*big.Int{big.NewInt(1)},
		big.NewInt(1),
		big.NewInt(1),
	)
	require.NoError(t, err)
	chain.Commit()
	return afnSourceAddress
}
                                                                                                    core/services/ccip/contract_transmitter.go                                                          000644  000765  000024  00000022107 14165357743 022047  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/bulletprooftxmanager"
	"github.com/smartcontractkit/chainlink/core/services/eth"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/chains/evmutil"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"gorm.io/gorm"
)

var (
	_ ocrtypes.ContractTransmitter = &OfframpTransmitter{}
	_ ocrtypes.ContractTransmitter = &ExecutionTransmitter{}
)

type ExecutionTransmitter struct {
	contractABI abi.ABI
	transmitter Transmitter
	contract    *message_executor.MessageExecutor
}

func NewExecutionTransmitter(
	contract *message_executor.MessageExecutor,
	contractABI abi.ABI,
	transmitter Transmitter,
) *ExecutionTransmitter {
	return &ExecutionTransmitter{
		contractABI: contractABI,
		transmitter: transmitter,
		contract:    contract,
	}
}

func (oc *ExecutionTransmitter) Transmit(ctx context.Context, reportCtx ocrtypes.ReportContext, report ocrtypes.Report, signatures []ocrtypes.AttributedOnchainSignature) error {
	rs, ss, vs := splitSigs(signatures)
	rawReportCtx := evmutil.RawReportContext(reportCtx)
	logger.Infow("executor transmitting report", "report", hex.EncodeToString(report), "rawReportCtx", rawReportCtx, "contractAddress", oc.contract.Address())

	payload, err := oc.contractABI.Pack("transmit", rawReportCtx, []byte(report), rs, ss, vs)
	if err != nil {
		return errors.Wrap(err, "abi.Pack failed")
	}

	return errors.Wrap(oc.transmitter.CreateEthTransaction(ctx, oc.contract.Address(), payload, report), "failed to send Eth transaction")
}

func (oc *ExecutionTransmitter) LatestConfigDigestAndEpoch(ctx context.Context) (configDigest ocrtypes.ConfigDigest, epoch uint32, err error) {
	//! TODO: not efficient!
	it, err := oc.contract.FilterTransmited(&bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: ctx,
	})
	if err != nil {
		return ocrtypes.ConfigDigest{}, 0, err
	}
	defer it.Close()
	for it.Next() {
		fmt.Println("LatestConfigDigestAndEpoch:", it.Event)
		configDigest = it.Event.ConfigDigest
		epoch = it.Event.Epoch
	}

	if it.Error() != nil {
		return ocrtypes.ConfigDigest{}, 0, it.Error()
	}
	return configDigest, epoch, nil
}

func (oc *ExecutionTransmitter) FromAccount() ocrtypes.Account {
	return ocrtypes.Account(oc.transmitter.FromAddress().String())
}

type OfframpTransmitter struct {
	contractABI abi.ABI
	transmitter Transmitter
	contract    *single_token_offramp.SingleTokenOffRamp
}

func NewOfframpTransmitter(
	contract *single_token_offramp.SingleTokenOffRamp,
	contractABI abi.ABI,
	transmitter Transmitter,
) *OfframpTransmitter {
	return &OfframpTransmitter{
		contractABI: contractABI,
		transmitter: transmitter,
		contract:    contract,
	}
}

func splitSigs(signatures []ocrtypes.AttributedOnchainSignature) (rs [][32]byte, ss [][32]byte, vs [32]byte) {
	for i, as := range signatures {
		r, s, v, err := evmutil.SplitSignature(as.Signature)
		if err != nil {
			panic("eventTransmit(ev): error in SplitSignature")
		}
		rs = append(rs, r)
		ss = append(ss, s)
		vs[i] = v
	}
	return
}

func (oc *OfframpTransmitter) Transmit(ctx context.Context, reportCtx ocrtypes.ReportContext, report ocrtypes.Report, signatures []ocrtypes.AttributedOnchainSignature) error {
	rs, ss, vs := splitSigs(signatures)
	rawReportCtx := evmutil.RawReportContext(reportCtx)
	logger.Debugw("Transmitting report", "report", hex.EncodeToString(report), "rawReportCtx", rawReportCtx, "contractAddress", oc.contract.Address())

	payload, err := oc.contractABI.Pack("transmit", rawReportCtx, []byte(report), rs, ss, vs)
	if err != nil {
		return errors.Wrap(err, "abi.Pack failed")
	}

	return errors.Wrap(oc.transmitter.CreateEthTransaction(ctx, oc.contract.Address(), payload, report), "failed to send Eth transaction")
}

func (oc *OfframpTransmitter) LatestConfigDigestAndEpoch(ctx context.Context) (configDigest ocrtypes.ConfigDigest, epoch uint32, err error) {
	//! TODO: not efficient!
	it, err := oc.contract.FilterTransmited(&bind.FilterOpts{
		Start:   0,
		End:     nil,
		Context: ctx,
	})
	if err != nil {
		return ocrtypes.ConfigDigest{}, 0, err
	}
	defer it.Close()
	for it.Next() {
		fmt.Println("LatestConfigDigestAndEpoch:", it.Event)
		configDigest = it.Event.ConfigDigest
		epoch = it.Event.Epoch
	}

	if it.Error() != nil {
		return ocrtypes.ConfigDigest{}, 0, it.Error()
	}
	return configDigest, epoch, nil
}

func (oc *OfframpTransmitter) FromAccount() ocrtypes.Account {
	return ocrtypes.Account(oc.transmitter.FromAddress().String())
}

type relayTransmitter struct {
	txm                        TxManager
	db                         *gorm.DB
	fromAddress                gethCommon.Address
	gasLimit                   uint64
	strategy                   bulletprooftxmanager.TxStrategy
	ec                         eth.Client
	sourceChainID, destChainID *big.Int
}

type TxManager interface {
	CreateEthTransaction(db *gorm.DB, newTx bulletprooftxmanager.NewTx) (etx bulletprooftxmanager.EthTx, err error)
}

type Transmitter interface {
	CreateEthTransaction(ctx context.Context, toAddress gethCommon.Address, payload []byte, report []byte) error
	FromAddress() gethCommon.Address
}

// NewTransmitter creates a new eth relayTransmitter
func NewRelayTransmitter(txm TxManager, db *gorm.DB, sourceChainID, destChainID *big.Int, fromAddress gethCommon.Address, gasLimit uint64, strategy bulletprooftxmanager.TxStrategy, ec eth.Client) Transmitter {
	return &relayTransmitter{
		txm:           txm,
		db:            db,
		fromAddress:   fromAddress,
		gasLimit:      gasLimit,
		strategy:      strategy,
		ec:            ec,
		sourceChainID: sourceChainID,
		destChainID:   destChainID,
	}
}

func (t *relayTransmitter) CreateEthTransaction(ctx context.Context, toAddress gethCommon.Address, payload []byte, report []byte) error {
	twoGwei := big.NewInt(2_000_000_000)
	a := toAddress
	gasEstimate, err := t.ec.EstimateGas(ctx, ethereum.CallMsg{
		From:      t.fromAddress,
		To:        &a,
		Gas:       0,
		GasFeeCap: twoGwei,
		GasTipCap: twoGwei,
		Data:      payload,
	})
	if err != nil {
		return errors.Wrap(err, "failed to estimating gas cost for ccip execute transaction")
	}
	if gasEstimate > t.gasLimit {
		return errors.Wrap(err, fmt.Sprintf("gas estimate of %d exceeds gas limit set by node %d", gasEstimate, t.gasLimit))
	}
	// TODO: As soon as gorm is removed, these can two db ops need to be in the same transaction
	_, err = t.txm.CreateEthTransaction(t.db, bulletprooftxmanager.NewTx{
		FromAddress:    t.fromAddress,
		ToAddress:      toAddress,
		EncodedPayload: payload,
		GasLimit:       gasEstimate,
		Meta:           nil,
		Strategy:       t.strategy,
	})
	return errors.Wrap(err, "creating ETH ccip relay transaction")
}

func (t *relayTransmitter) FromAddress() gethCommon.Address {
	return t.fromAddress
}

type executeTransmitter struct {
	txm                        TxManager
	db                         *gorm.DB
	fromAddress                gethCommon.Address
	gasLimit                   uint64
	strategy                   bulletprooftxmanager.TxStrategy
	ec                         eth.Client
	sourceChainID, destChainID *big.Int
}

// NewTransmitter creates a new eth relayTransmitter
func NewExecuteTransmitter(txm TxManager, db *gorm.DB, sourceChainID, destChainID *big.Int, fromAddress gethCommon.Address, gasLimit uint64, strategy bulletprooftxmanager.TxStrategy, ec eth.Client) Transmitter {
	return &executeTransmitter{
		txm:           txm,
		db:            db,
		fromAddress:   fromAddress,
		gasLimit:      gasLimit,
		strategy:      strategy,
		ec:            ec,
		sourceChainID: sourceChainID,
		destChainID:   destChainID,
	}
}

func (t *executeTransmitter) CreateEthTransaction(ctx context.Context, toAddress gethCommon.Address, payload []byte, report []byte) error {
	twoGwei := big.NewInt(2_000_000_000)
	a := toAddress
	gasEstimate, err := t.ec.EstimateGas(ctx, ethereum.CallMsg{
		From:      t.fromAddress,
		To:        &a,
		Gas:       0,
		GasFeeCap: twoGwei,
		GasTipCap: twoGwei,
		Data:      payload,
	})
	if err != nil {
		return errors.Wrap(err, "failed to estimating gas cost for ccip execute transaction")
	}
	if gasEstimate > t.gasLimit {
		return errors.Wrap(err, fmt.Sprintf("gas estimate of %d exceeds gas limit set by node %d", gasEstimate, t.gasLimit))
	}
	// TODO: As soon as gorm is removed, these can two db ops need to be in the same transaction
	_, err = t.txm.CreateEthTransaction(t.db, bulletprooftxmanager.NewTx{
		FromAddress:    t.fromAddress,
		ToAddress:      toAddress,
		EncodedPayload: payload,
		GasLimit:       gasEstimate,
		Meta:           nil,
		Strategy:       t.strategy,
	})
	return errors.Wrap(err, "creating ETH ccip execute transaction")
}

func (t *executeTransmitter) FromAddress() gethCommon.Address {
	return t.fromAddress
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                         core/services/ccip/contract_tracker_test.go                                                         000644  000765  000024  00000001172 14165346401 022151  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOffchainConfig_Encode_Decode(t *testing.T) {
	tests := map[string]struct {
		want  OffchainConfig
		error bool
	}{
		"Success": {
			want: OffchainConfig{
				SourceIncomingConfirmations: 3,
				DestIncomingConfirmations:   6,
			},
		},
		"Missing value as 0": {
			want: OffchainConfig{
				SourceIncomingConfirmations: 99999999,
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			encode, err := tc.want.Encode()
			got, err := Decode(encode)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}
                                                                                                                                                                                                                                                                                                                                                                                                      core/services/ccip/integration_test.go                                                              000644  000765  000024  00000075500 14165357743 021165  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip_test

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/onsi/gomega"
	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/services/ccip/abihelpers"
	"github.com/smartcontractkit/libocr/commontypes"
	ocrnetworking "github.com/smartcontractkit/libocr/networking"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"

	evmtypes "github.com/smartcontractkit/chainlink/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/core/gracefulpanic"
	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/lock_unlock_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_sender"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/configtest"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/evmtest"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/bulletprooftxmanager"
	"github.com/smartcontractkit/chainlink/core/services/ccip"
	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/services/eth"
	"github.com/smartcontractkit/chainlink/core/services/headtracker"
	httypes "github.com/smartcontractkit/chainlink/core/services/headtracker/types"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/core/services/log"
	"github.com/smartcontractkit/chainlink/core/services/postgres"
	"github.com/smartcontractkit/chainlink/core/utils"
)

func setupChain(t *testing.T) (*backends.SimulatedBackend, *bind.TransactOpts) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	user, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	chain := backends.NewSimulatedBackend(core.GenesisAlloc{
		user.From: {Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1000000000000000000))}},
		ethconfig.Defaults.Miner.GasCeil)
	return chain, user
}

type CCIPContracts struct {
	sourceUser, destUser           *bind.TransactOpts
	sourceChain, destChain         *backends.SimulatedBackend
	sourcePool, destPool           *lock_unlock_pool.LockUnlockPool
	onRamp                         *single_token_onramp.SingleTokenOnRamp
	sourceLinkToken, destLinkToken *link_token_interface.LinkToken
	offRamp                        *single_token_offramp.SingleTokenOffRamp
	messageReceiver                *simple_message_receiver.SimpleMessageReceiver
	eoaTokenSender                 *single_token_sender.EOASingleTokenSender
	eoaTokenReceiver               *single_token_receiver.EOASingleTokenReceiver
	executor                       *message_executor.MessageExecutor
}

func setupCCIPContracts(t *testing.T) CCIPContracts {
	sourceChain, sourceUser := setupChain(t)
	destChain, destUser := setupChain(t)

	// Deploy link token and pool on source chain
	sourceLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(sourceUser, sourceChain)
	require.NoError(t, err)
	sourceChain.Commit()
	sourceLinkToken, err := link_token_interface.NewLinkToken(sourceLinkTokenAddress, sourceChain)
	require.NoError(t, err)
	sourcePoolAddress, _, _, err := lock_unlock_pool.DeployLockUnlockPool(sourceUser, sourceChain, sourceLinkTokenAddress)
	require.NoError(t, err)
	sourceChain.Commit()
	sourcePool, err := lock_unlock_pool.NewLockUnlockPool(sourcePoolAddress, sourceChain)
	require.NoError(t, err)

	// Deploy link token and pool on destination chain
	destLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()
	destLinkToken, err := link_token_interface.NewLinkToken(destLinkTokenAddress, destChain)
	require.NoError(t, err)
	destPoolAddress, _, _, err := lock_unlock_pool.DeployLockUnlockPool(destUser, destChain, destLinkTokenAddress)
	require.NoError(t, err)
	destChain.Commit()
	destPool, err := lock_unlock_pool.NewLockUnlockPool(destPoolAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	// Float the offramp pool with 1M juels
	// Dest user is the owner of the dest pool, so he can store
	o, err := destPool.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, destUser.From.String(), o.String())
	b, err := destLinkToken.BalanceOf(nil, destUser.From)
	require.NoError(t, err)
	t.Log("balance", b)
	_, err = destLinkToken.Approve(destUser, destPoolAddress, big.NewInt(1000000))
	require.NoError(t, err)
	destChain.Commit()
	_, err = destPool.LockOrBurn(destUser, destUser.From, big.NewInt(1000000))
	require.NoError(t, err)

	afnSourceAddress, _, _, err := afn_contract.DeployAFNContract(
		sourceUser,
		sourceChain,
		[]common.Address{sourceUser.From},
		[]*big.Int{big.NewInt(1)},
		big.NewInt(1),
		big.NewInt(1),
	)
	require.NoError(t, err)
	sourceChain.Commit()

	// Deploy onramp source chain
	onRampAddress, _, _, err := single_token_onramp.DeploySingleTokenOnRamp(
		sourceUser,             // users
		sourceChain,            // backend
		sourceChainID,          // source chain id
		sourceLinkTokenAddress, // token
		sourcePoolAddress,      // pool
		destChainID,            // remoteChainId
		destLinkTokenAddress,   // remoteToken
		[]common.Address{},     // allow list
		false,                  // enableAllowList
		big.NewInt(1),          // token bucket rate
		big.NewInt(1000),       // token bucket capacity,
		afnSourceAddress,       // AFN
		big.NewInt(86400),      //maxTimeWithoutAFNSignal 86400 seconds = one day
	)
	require.NoError(t, err)
	// We do this so onRamp.Address() works
	onRamp, err := single_token_onramp.NewSingleTokenOnRamp(onRampAddress, sourceChain)
	require.NoError(t, err)
	_, err = sourcePool.SetOnRamp(sourceUser, onRampAddress, true)
	require.NoError(t, err)

	afnDestAddress, _, _, err := afn_contract.DeployAFNContract(
		destUser,
		destChain,
		[]common.Address{destUser.From},
		[]*big.Int{big.NewInt(1)},
		big.NewInt(1),
		big.NewInt(1),
	)
	require.NoError(t, err)
	destChain.Commit()

	// Deploy offramp dest chain
	offRampAddress, _, _, err := single_token_offramp.DeploySingleTokenOffRamp(
		destUser,
		destChain,
		sourceChainID,
		destChainID,
		destLinkTokenAddress,
		destPoolAddress,
		big.NewInt(1),     // token bucket rate
		big.NewInt(1000),  // token bucket capacity,
		afnDestAddress,    // AFN
		big.NewInt(86400), //maxTimeWithoutAFNSignal 86400 seconds = one day
		big.NewInt(0),     // execution delay in seconds
	)
	require.NoError(t, err)
	offRamp, err := single_token_offramp.NewSingleTokenOffRamp(offRampAddress, destChain)
	require.NoError(t, err)
	// Set the pool to be the offramp
	_, err = destPool.SetOffRamp(destUser, offRampAddress, true)
	require.NoError(t, err)

	// Deploy offramp contract token receiver
	messageReceiverAddress, _, _, err := simple_message_receiver.DeploySimpleMessageReceiver(destUser, destChain)
	require.NoError(t, err)
	messageReceiver, err := simple_message_receiver.NewSimpleMessageReceiver(messageReceiverAddress, destChain)
	require.NoError(t, err)
	// Deploy offramp EOA token receiver
	eoaTokenReceiverAddress, _, _, err := single_token_receiver.DeployEOASingleTokenReceiver(destUser, destChain, offRampAddress)
	require.NoError(t, err)
	eoaTokenReceiver, err := single_token_receiver.NewEOASingleTokenReceiver(eoaTokenReceiverAddress, destChain)
	require.NoError(t, err)
	// Deploy onramp EOA token sender
	eoaTokenSenderAddress, _, _, err := single_token_sender.DeployEOASingleTokenSender(sourceUser, sourceChain, onRampAddress, eoaTokenReceiverAddress)
	require.NoError(t, err)
	eoaTokenSender, err := single_token_sender.NewEOASingleTokenSender(eoaTokenSenderAddress, sourceChain)
	require.NoError(t, err)

	// Deploy the message executor ocr2 contract
	executorAddress, _, _, err := message_executor.DeployMessageExecutor(destUser, destChain, offRampAddress)
	require.NoError(t, err)
	executor, err := message_executor.NewMessageExecutor(executorAddress, destChain)
	require.NoError(t, err)

	sourceChain.Commit()
	destChain.Commit()

	return CCIPContracts{
		sourceUser:       sourceUser,
		destUser:         destUser,
		sourceChain:      sourceChain,
		destChain:        destChain,
		sourcePool:       sourcePool,
		destPool:         destPool,
		onRamp:           onRamp,
		sourceLinkToken:  sourceLinkToken,
		destLinkToken:    destLinkToken,
		offRamp:          offRamp,
		messageReceiver:  messageReceiver,
		eoaTokenReceiver: eoaTokenReceiver,
		eoaTokenSender:   eoaTokenSender,
		executor:         executor,
	}
}

var (
	sourceChainID = big.NewInt(1000)
	destChainID   = big.NewInt(2000)
)

type EthKeyStoreSim struct {
	keystore.Eth
}

func (ks EthKeyStoreSim) SignTx(address common.Address, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	if chainID.String() == "1000" || chainID.String() == "2000" {
		// A terrible hack, just for the multichain test
		// Needs to actually use the sim
		return ks.Eth.SignTx(address, tx, big.NewInt(1337))
	}
	return ks.Eth.SignTx(address, tx, chainID)
}

var _ keystore.Eth = EthKeyStoreSim{}

func setupNodeCCIP(t *testing.T, owner *bind.TransactOpts, port uint16, dbName string, sourceChain *backends.SimulatedBackend, destChain *backends.SimulatedBackend) (chainlink.Application, string, common.Address, ocr2key.KeyBundle, *configtest.TestGeneralConfig, func()) {
	// Do not want to load fixtures as they contain a dummy chainID.
	config, _, db := heavyweight.FullTestDB(t, fmt.Sprintf("%s%d", dbName, port), true, false)
	config.Overrides.FeatureOffchainReporting = null.BoolFrom(false)
	config.Overrides.FeatureOffchainReporting2 = null.BoolFrom(true)
	config.Overrides.GlobalGasEstimatorMode = null.NewString("FixedPrice", true)
	config.Overrides.DefaultChainID = nil
	config.Overrides.P2PListenPort = port
	config.Overrides.P2PNetworkingStack = ocrnetworking.NetworkingStackV2
	// Disables ocr spec validation so we can have fast polling for the test.
	config.Overrides.Dev = null.BoolFrom(true)

	eventBroadcaster := postgres.NewEventBroadcaster(config.DatabaseURL(), 0, 0)
	var lggr = logger.TestLogger(t)
	shutdown := gracefulpanic.NewSignal()
	sqlxDB := postgres.UnwrapGormDB(db)

	// We fake different chainIDs using the wrapped sim cltest.SimulatedBackend
	chainORM := evm.NewORM(sqlxDB)
	_, err := chainORM.CreateChain(*utils.NewBig(sourceChainID), evmtypes.ChainCfg{})
	require.NoError(t, err)
	_, err = chainORM.CreateChain(*utils.NewBig(destChainID), evmtypes.ChainCfg{})
	require.NoError(t, err)
	sourceClient := cltest.NewSimulatedBackendClient(t, sourceChain, sourceChainID)
	destClient := cltest.NewSimulatedBackendClient(t, destChain, destChainID)

	keyStore := keystore.New(db, utils.FastScryptParams, lggr)
	simEthKeyStore := EthKeyStoreSim{Eth: keyStore.Eth()}

	// Create our chainset manually so we can have custom eth clients
	// (the wrapped sims faking different chainIDs)
	chainSet, err := evm.LoadChainSet(evm.ChainSetOpts{
		ORM:              chainORM,
		Config:           config,
		Logger:           lggr,
		GormDB:           db,
		SQLxDB:           sqlxDB,
		KeyStore:         simEthKeyStore,
		EventBroadcaster: eventBroadcaster,
		GenEthClient: func(c evmtypes.Chain) eth.Client {
			if c.ID.String() == sourceChainID.String() {
				return sourceClient
			} else if c.ID.String() == destChainID.String() {
				return destClient
			}
			t.Fatalf("invalid chain ID %v", c.ID.String())
			return nil
		},
		GenHeadTracker: func(c evmtypes.Chain, ht httypes.HeadBroadcaster) httypes.Tracker {
			if c.ID.String() == sourceChainID.String() {
				return headtracker.NewHeadTracker(lggr, sourceClient, evmtest.NewChainScopedConfig(t, config), headtracker.NewORM(db, *sourceChainID), ht)
			} else if c.ID.String() == destChainID.String() {
				return headtracker.NewHeadTracker(lggr, destClient, evmtest.NewChainScopedConfig(t, config), headtracker.NewORM(db, *destChainID), ht)
			}
			t.Fatalf("invalid chain ID %v", c.ID.String())
			return nil
		},
		GenLogBroadcaster: func(c evmtypes.Chain) log.Broadcaster {
			if c.ID.String() == sourceChainID.String() {
				t.Log("Generating log broadcaster source")
				return log.NewBroadcaster(log.NewORM(db, *sourceChainID), sourceClient,
					evmtest.NewChainScopedConfig(t, config), lggr, nil)
			} else if c.ID.String() == destChainID.String() {
				return log.NewBroadcaster(log.NewORM(db, *destChainID), destClient,
					evmtest.NewChainScopedConfig(t, config), lggr, nil)
			}
			t.Fatalf("invalid chain ID %v", c.ID.String())
			return nil
		},
		GenTxManager: func(c evmtypes.Chain) bulletprooftxmanager.TxManager {
			if c.ID.String() == sourceChainID.String() {
				return bulletprooftxmanager.NewBulletproofTxManager(db, sourceClient, evmtest.NewChainScopedConfig(t, config), simEthKeyStore, eventBroadcaster, lggr)
			} else if c.ID.String() == destChainID.String() {
				return bulletprooftxmanager.NewBulletproofTxManager(db, destClient, evmtest.NewChainScopedConfig(t, config), simEthKeyStore, eventBroadcaster, lggr)
			}
			t.Fatalf("invalid chain ID %v", c.ID.String())
			return nil
		},
	})
	if err != nil {
		logger.Fatal(err)
	}
	app, err := chainlink.NewApplication(chainlink.ApplicationOpts{
		Config:                   config,
		EventBroadcaster:         eventBroadcaster,
		ShutdownSignal:           shutdown,
		GormDB:                   db,
		SqlxDB:                   sqlxDB,
		KeyStore:                 keyStore,
		ChainSet:                 chainSet,
		Logger:                   lggr,
		ExternalInitiatorManager: nil,
	})
	require.NoError(t, err)
	require.NoError(t, app.GetKeyStore().Unlock("password"))
	_, err = app.GetKeyStore().P2P().Create()
	require.NoError(t, err)

	p2pIDs, err := app.GetKeyStore().P2P().GetAll()
	require.NoError(t, err)
	require.Len(t, p2pIDs, 1)
	peerID := p2pIDs[0].PeerID()

	config.Overrides.P2PPeerID = peerID
	config.Overrides.P2PListenPort = port
	p2paddresses := []string{
		fmt.Sprintf("127.0.0.1:%d", port),
	}
	config.Overrides.P2PV2ListenAddresses = p2paddresses
	config.Overrides.P2PV2AnnounceAddresses = p2paddresses

	_, err = app.GetKeyStore().Eth().Create(destChainID)
	require.NoError(t, err)
	sendingKeys, err := app.GetKeyStore().Eth().SendingKeys()
	require.NoError(t, err)
	require.Len(t, sendingKeys, 1)
	transmitter := sendingKeys[0].Address.Address()
	s, err := app.GetKeyStore().Eth().GetState(sendingKeys[0].ID())
	require.NoError(t, err)
	logger.Debug(fmt.Sprintf("Transmitter address %s chainID %s", transmitter, s.EVMChainID.String()))

	// Fund the relayTransmitter address with some ETH
	n, err := destChain.NonceAt(context.Background(), owner.From, nil)
	require.NoError(t, err)

	tx := types.NewTransaction(n, transmitter, big.NewInt(1000000000000000000), 21000, big.NewInt(1000000000), nil)
	signedTx, err := owner.Signer(owner.From, tx)
	require.NoError(t, err)
	err = destChain.SendTransaction(context.Background(), signedTx)
	require.NoError(t, err)
	destChain.Commit()

	kb, err := app.GetKeyStore().OCR2().Create()
	require.NoError(t, err)
	return app, peerID.Raw(), transmitter, kb, config, func() {
		app.Stop()
	}
}

func TestIntegration_CCIP(t *testing.T) {
	ccipContracts := setupCCIPContracts(t)
	// Oracles need ETH on the destination chain
	bootstrapNodePort := uint16(19599)
	appBootstrap, bootstrapPeerID, _, _, _, _ := setupNodeCCIP(t, ccipContracts.destUser, bootstrapNodePort, "bootstrap_ccip", ccipContracts.sourceChain, ccipContracts.destChain)
	var (
		oracles      []confighelper2.OracleIdentityExtra
		transmitters []common.Address
		kbs          []ocr2key.KeyBundle
		apps         []chainlink.Application
	)
	// Set up the minimum 4 oracles all funded with destination ETH
	for i := uint16(0); i < 4; i++ {
		app, peerID, transmitter, kb, cfg, _ := setupNodeCCIP(t, ccipContracts.destUser, bootstrapNodePort+1+i, fmt.Sprintf("oracle_ccip%d", i), ccipContracts.sourceChain, ccipContracts.destChain)
		// Supply the bootstrap IP and port as a V2 peer address
		cfg.Overrides.P2PV2Bootstrappers = []commontypes.BootstrapperLocator{
			{PeerID: bootstrapPeerID, Addrs: []string{
				fmt.Sprintf("127.0.0.1:%d", bootstrapNodePort),
			}},
		}
		kbs = append(kbs, kb)
		apps = append(apps, app)
		transmitters = append(transmitters, transmitter)
		oracles = append(oracles, confighelper2.OracleIdentityExtra{
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  kb.OnchainKeyring.SigningAddress().Bytes(),
				TransmitAccount:   ocrtypes2.Account(transmitter.String()),
				OffchainPublicKey: kb.OffchainKeyring.OffchainPublicKey(),
				PeerID:            peerID,
			},
			ConfigEncryptionPublicKey: kb.OffchainKeyring.ConfigEncryptionPublicKey(),
		})
	}

	reportingPluginConfig, err := ccip.OffchainConfig{
		SourceIncomingConfirmations: 0,
		DestIncomingConfirmations:   1,
	}.Encode()
	require.NoError(t, err)

	setupOnchainConfig(t, ccipContracts, oracles, reportingPluginConfig)

	err = appBootstrap.Start()
	require.NoError(t, err)
	defer appBootstrap.Stop()

	// Add the bootstrap job
	chainSet := appBootstrap.GetChainSet()
	require.NotNil(t, chainSet)
	ocrJob, err := ccip.ValidatedCCIPBootstrapSpec(fmt.Sprintf(`
type               = "ccip-bootstrap"
schemaVersion      = 1
evmChainID         = "%s"
name               = "boot"
contractAddress    = "%s"
isBootstrapPeer    = true
contractConfigConfirmations = 1
contractConfigTrackerPollInterval = "1s"
`, destChainID, ccipContracts.offRamp.Address()))
	require.NoError(t, err)
	_, err = appBootstrap.AddJobV2(context.Background(), ocrJob, null.NewString("boot", true))
	require.NoError(t, err)

	// For each oracle add a relayer and job
	for i := 0; i < 4; i++ {
		err = apps[i].Start()
		require.NoError(t, err)
		defer apps[i].Stop()
		// Wait for peer wrapper to start
		time.Sleep(1 * time.Second)
		ccipJob, err := ccip.ValidatedCCIPSpec(fmt.Sprintf(`
type               = "ccip-relay"
schemaVersion      = 1
name               = "ccip-job-%d"
onRampAddress = "%s"
offRampAddress = "%s"
sourceEvmChainID   = "%s"
destEvmChainID     = "%s"
keyBundleID        = "%s"
transmitterAddress = "%s"
contractConfigConfirmations = 1
contractConfigTrackerPollInterval = "1s"
`, i, ccipContracts.onRamp.Address(), ccipContracts.offRamp.Address(), sourceChainID, destChainID, kbs[i].ID(), transmitters[i]))
		require.NoError(t, err)
		_, err = apps[i].AddJobV2(context.Background(), ccipJob, null.NewString("ccip", true))
		require.NoError(t, err)
		// Add executor job
		ccipExecutionJob, err := ccip.ValidatedCCIPSpec(fmt.Sprintf(`
type               = "ccip-execution"
schemaVersion      = 1
name               = "ccip-executor-job-%d"
onRampAddress = "%s"
offRampAddress = "%s"
executorAddress = "%s"
sourceEvmChainID   = "%s"
destEvmChainID     = "%s"
keyBundleID        = "%s"
transmitterAddress = "%s"
contractConfigConfirmations = 1
contractConfigTrackerPollInterval = "1s"
`, i, ccipContracts.onRamp.Address(), ccipContracts.offRamp.Address(), ccipContracts.executor.Address(), sourceChainID, destChainID, kbs[i].ID(), transmitters[i]))
		require.NoError(t, err)
		_, err = apps[i].AddJobV2(context.Background(), ccipExecutionJob, null.NewString("ccip-executor", true))
		require.NoError(t, err)
	}
	// Send a request.
	// Jobs are booting but that is ok, the log broadcaster
	// will backfill this request log.
	ccipContracts.sourceUser.GasLimit = 500000
	_, err = ccipContracts.sourceLinkToken.Approve(ccipContracts.sourceUser, ccipContracts.sourcePool.Address(), big.NewInt(100))
	ccipContracts.sourceChain.Commit()
	msg := single_token_onramp.CCIPMessagePayload{
		Receiver: ccipContracts.messageReceiver.Address(),
		Data:     []byte("hello xchain world"),
		Tokens:   []common.Address{ccipContracts.sourceLinkToken.Address()},
		Amounts:  []*big.Int{big.NewInt(100)},
		Options:  nil,
	}
	tx, err := ccipContracts.onRamp.RequestCrossChainSend(ccipContracts.sourceUser, msg)
	require.NoError(t, err)
	ccipContracts.sourceChain.Commit()
	rec, err := ccipContracts.sourceChain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), rec.Status)

	reportingPluginConfig, err = ccip.OffchainConfig{
		SourceIncomingConfirmations: 1,
		DestIncomingConfirmations:   0,
	}.Encode()
	require.NoError(t, err)

	setupOnchainConfig(t, ccipContracts, oracles, reportingPluginConfig)

	// Request should appear on all nodes eventually
	for i := 0; i < 4; i++ {
		var reqs []*ccip.Request
		ccipReqORM := ccip.NewORM(postgres.UnwrapGormDB(apps[i].GetDB()))
		gomega.NewGomegaWithT(t).Eventually(func() bool {
			ccipContracts.sourceChain.Commit()
			reqs, err = ccipReqORM.Requests(sourceChainID, destChainID, big.NewInt(0), nil, ccip.RequestStatusUnstarted, nil, nil)
			return len(reqs) == 1
		}, 5*time.Second, 1*time.Second).Should(gomega.BeTrue())
	}

	// Once all nodes have the request, the reporting plugin should run to generate and submit a report onchain.
	// So we should eventually see a successful offramp submission.
	// Note that since we only send blocks here, it's likely that all the nodes will enter the transmission
	// phase before someone has submitted, so 1 report will succeed and 3 will revert.
	var report single_token_offramp.CCIPRelayReport
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		report, err = ccipContracts.offRamp.GetLastReport(nil)
		require.NoError(t, err)
		ccipContracts.destChain.Commit()
		return report.MinSequenceNumber.String() == "1" && report.MaxSequenceNumber.String() == "1"
	}, 10*time.Second, 1*time.Second).Should(gomega.BeTrue())

	// We should see the request in a fulfilled state on all nodes
	// after the offramp submission. There should be no
	// remaining valid requests.
	for i := 0; i < 4; i++ {
		gomega.NewGomegaWithT(t).Eventually(func() bool {
			ccipReqORM := ccip.NewORM(postgres.UnwrapGormDB(apps[i].GetDB()))
			ccipContracts.destChain.Commit()
			reqs, err := ccipReqORM.Requests(sourceChainID, destChainID, report.MinSequenceNumber, report.MaxSequenceNumber, ccip.RequestStatusRelayConfirmed, nil, nil)
			require.NoError(t, err)
			valid, err := ccipReqORM.Requests(sourceChainID, destChainID, report.MinSequenceNumber, nil, ccip.RequestStatusUnstarted, nil, nil)
			require.NoError(t, err)
			return len(reqs) == 1 && len(valid) == 0
		}, 10*time.Second, 1*time.Second).Should(gomega.BeTrue())
	}

	// Now the merkle root is across.
	// Let's try to execute a request as an external party.
	// The raw log in the merkle root should be the abi-encoded version of the CCIPMessage
	ccipReqORM := ccip.NewORM(postgres.UnwrapGormDB(apps[0].GetDB()))
	reqs, err := ccipReqORM.Requests(sourceChainID, destChainID, report.MinSequenceNumber, report.MaxSequenceNumber, "", nil, nil)
	require.NoError(t, err)
	root, proof := ccip.GenerateMerkleProof(32, [][]byte{reqs[0].Raw}, 0)
	// Root should match the report root
	require.True(t, bytes.Equal(root[:], report.MerkleRoot[:]))

	// Proof should verify.
	genRoot := ccip.GenerateMerkleRoot(reqs[0].Raw, proof)
	require.True(t, bytes.Equal(root[:], genRoot[:]))
	exists, err := ccipContracts.offRamp.GetMerkleRoot(nil, report.MerkleRoot)
	require.NoError(t, err)
	require.True(t, exists.Int64() > 0)

	h := utils.MustFixedKeccak256(append([]byte{0x00}, reqs[0].Raw...))
	onchainRoot, err := ccipContracts.offRamp.GenerateMerkleRoot(nil, proof.PathForExecute(), h, proof.Index())
	require.NoError(t, err)
	require.Equal(t, genRoot, onchainRoot)

	// Execute the Message
	decodedMsg, err := abihelpers.DecodeCCIPMessage(reqs[0].Raw)
	require.NoError(t, err)
	abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{*decodedMsg})
	tx, err = ccipContracts.offRamp.ExecuteTransaction(ccipContracts.destUser, proof.PathForExecute(), *decodedMsg, proof.Index())
	require.NoError(t, err)
	ccipContracts.destChain.Commit()

	// We should now have the Message in the offchain receiver
	receivedMsg, err := ccipContracts.messageReceiver.SMessage(nil)
	require.NoError(t, err)
	assert.Equal(t, "hello xchain world", string(receivedMsg.Payload.Data))

	// Now let's send an EOA to EOA request
	// We can just use the sourceUser and destUser
	startBalanceSource, err := ccipContracts.sourceLinkToken.BalanceOf(nil, ccipContracts.sourceUser.From)
	require.NoError(t, err)
	startBalanceDest, err := ccipContracts.destLinkToken.BalanceOf(nil, ccipContracts.destUser.From)
	require.NoError(t, err)
	t.Log(startBalanceSource, startBalanceDest)

	ccipContracts.sourceUser.GasLimit = 500000
	// Approve the sender contract to take the tokens
	_, err = ccipContracts.sourceLinkToken.Approve(ccipContracts.sourceUser, ccipContracts.eoaTokenSender.Address(), big.NewInt(100))
	ccipContracts.sourceChain.Commit()
	// Send the tokens. Should invoke the onramp.
	// Only the destUser can execute.
	tx, err = ccipContracts.eoaTokenSender.SendTokens(ccipContracts.sourceUser, ccipContracts.destUser.From, big.NewInt(100), ccipContracts.destUser.From)
	require.NoError(t, err)
	ccipContracts.sourceChain.Commit()

	// DON should eventually send another report
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		report, err = ccipContracts.offRamp.GetLastReport(nil)
		require.NoError(t, err)
		ccipContracts.destChain.Commit()
		return report.MinSequenceNumber.String() == "2" && report.MaxSequenceNumber.String() == "2"
	}, 10*time.Second, 1*time.Second).Should(gomega.BeTrue())

	eoaReq, err := ccipReqORM.Requests(sourceChainID, destChainID, report.MinSequenceNumber, report.MaxSequenceNumber, "", nil, nil)
	require.NoError(t, err)
	root, proof = ccip.GenerateMerkleProof(32, [][]byte{eoaReq[0].Raw}, 0)
	// Root should match the report root
	require.True(t, bytes.Equal(root[:], report.MerkleRoot[:]))

	// Execute the Message
	decodedMsg, err = abihelpers.DecodeCCIPMessage(eoaReq[0].Raw)
	require.NoError(t, err)
	abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{*decodedMsg})
	tx, err = ccipContracts.offRamp.ExecuteTransaction(ccipContracts.destUser, proof.PathForExecute(), *decodedMsg, proof.Index())
	require.NoError(t, err)
	ccipContracts.destChain.Commit()

	// The destination user's balance should increase
	endBalanceSource, err := ccipContracts.sourceLinkToken.BalanceOf(nil, ccipContracts.sourceUser.From)
	require.NoError(t, err)
	endBalanceDest, err := ccipContracts.destLinkToken.BalanceOf(nil, ccipContracts.destUser.From)
	require.NoError(t, err)
	t.Log("Start balances", startBalanceSource, startBalanceDest)
	t.Log("End balances", endBalanceSource, endBalanceDest)
	assert.Equal(t, "100", big.NewInt(0).Sub(startBalanceSource, endBalanceSource).String())
	assert.Equal(t, "100", big.NewInt(0).Sub(endBalanceDest, startBalanceDest).String())

	// Now let's send a request flagged for oracle execution
	_, err = ccipContracts.sourceLinkToken.Approve(ccipContracts.sourceUser, ccipContracts.sourcePool.Address(), big.NewInt(100))
	require.NoError(t, err)
	ccipContracts.sourceChain.Commit()
	require.NoError(t, err)
	msg = single_token_onramp.CCIPMessagePayload{
		Receiver: ccipContracts.messageReceiver.Address(),
		Data:     []byte("hey DON, execute for me"),
		Tokens:   []common.Address{ccipContracts.sourceLinkToken.Address()},
		Amounts:  []*big.Int{big.NewInt(100)},
		Executor: ccipContracts.executor.Address(),
		Options:  []byte{},
	}
	_, err = ccipContracts.onRamp.RequestCrossChainSend(ccipContracts.sourceUser, msg)
	require.NoError(t, err)
	ccipContracts.sourceChain.Commit()

	// Should first be relayed, seq number 3
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		report, err = ccipContracts.offRamp.GetLastReport(nil)
		require.NoError(t, err)
		ccipContracts.destChain.Commit()
		return report.MinSequenceNumber.String() == "3" && report.MaxSequenceNumber.String() == "3"
	}, 10*time.Second, 1*time.Second).Should(gomega.BeTrue())

	// Should see the 3rd message be executed
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		it, err := ccipContracts.offRamp.FilterCrossChainMessageExecuted(nil, nil)
		require.NoError(t, err)
		ecount := 0
		for it.Next() {
			t.Log("executed", it.Event.SequenceNumber)
			ecount++
		}
		ccipContracts.destChain.Commit()
		return ecount == 3
	}, 20*time.Second, 1*time.Second).Should(gomega.BeTrue())
	// In total, we should see 3 relay reports containing seq 1,2,3
	// and 3 execution_confirmed messages
	reqs, err = ccipReqORM.Requests(sourceChainID, destChainID, big.NewInt(1), big.NewInt(3), ccip.RequestStatusExecutionConfirmed, nil, nil)
	require.NoError(t, err)
	require.Len(t, reqs, 3)
	_, err = ccipReqORM.RelayReport(big.NewInt(1))
	require.NoError(t, err)
	_, err = ccipReqORM.RelayReport(big.NewInt(2))
	require.NoError(t, err)
	_, err = ccipReqORM.RelayReport(big.NewInt(3))
	require.NoError(t, err)
}

func setupOnchainConfig(t *testing.T, ccipContracts CCIPContracts, oracles []confighelper2.OracleIdentityExtra, reportingPluginConfig []byte) {
	// Note We do NOT set the payees, payment is done in the OCR2Base implementation
	// Set the offramp config.
	signers, transmitters, threshold, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper2.ContractSetConfigArgs(
		2*time.Second,        // deltaProgress
		1*time.Second,        // deltaResend
		1*time.Second,        // deltaRound
		500*time.Millisecond, // deltaGrace
		2*time.Second,        // deltaStage
		3,
		[]int{1, 1, 1, 1},
		oracles,
		reportingPluginConfig,
		50*time.Millisecond,
		50*time.Millisecond,
		50*time.Millisecond,
		50*time.Millisecond,
		50*time.Millisecond,
		1, // faults
		nil,
	)

	require.NoError(t, err)
	logger.Debugw("Setting Config on Oracle Contract",
		"signers", signers,
		"transmitters", transmitters,
		"threshold", threshold,
		"onchainConfig", onchainConfig,
		"encodedConfigVersion", offchainConfigVersion,
	)

	// Set the DON on the offramp
	_, err = ccipContracts.offRamp.SetConfig(
		ccipContracts.destUser,
		signers,
		transmitters,
		threshold,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	require.NoError(t, err)
	ccipContracts.destChain.Commit()

	// Same DON on the message executor
	_, err = ccipContracts.executor.SetConfig(
		ccipContracts.destUser,
		signers,
		transmitters,
		threshold,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	require.NoError(t, err)
	ccipContracts.destChain.Commit()
}
                                                                                                                                                                                                core/services/ccip/relay_reporting_plugin_test.go                                                   000644  000765  000024  00000007715 14165346401 023415  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/lock_unlock_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp_helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRelayReportEncoding(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	destUser, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	destChain := backends.NewSimulatedBackend(core.GenesisAlloc{
		destUser.From: {Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1000000000000000000))}},
		ethconfig.Defaults.Miner.GasCeil)
	destLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()
	_, err = link_token_interface.NewLinkToken(destLinkTokenAddress, destChain)
	require.NoError(t, err)
	destPoolAddress, _, _, err := lock_unlock_pool.DeployLockUnlockPool(destUser, destChain, destLinkTokenAddress)
	require.NoError(t, err)
	destChain.Commit()
	_, err = lock_unlock_pool.NewLockUnlockPool(destPoolAddress, destChain)
	require.NoError(t, err)
	destAfn := deployAfn(t, destUser, destChain)

	offRampAddress, _, _, err := single_token_offramp_helper.DeploySingleTokenOffRampHelper(
		destUser,             // user
		destChain,            // client
		big.NewInt(1337),     // source chain id
		big.NewInt(1338),     // dest chain id
		destLinkTokenAddress, // link token address
		destPoolAddress,      // dest pool address
		big.NewInt(1),        // token bucket rate
		big.NewInt(1000),     // token bucket capacity
		destAfn,              // AFN address
		// 86400 seconds = one day
		big.NewInt(86400), // max timeout without AFN signal
		big.NewInt(0),     // execution delay in seconds
	)
	require.NoError(t, err)
	offRamp, err := single_token_offramp_helper.NewSingleTokenOffRampHelper(offRampAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	r, proof := GenerateMerkleProof(2, [][]byte{{0xaa}}, 0)
	rootLocal := GenerateMerkleRoot([]byte{0xaa}, proof)
	require.True(t, bytes.Equal(rootLocal[:], r[:]))
	t.Log(proof.PathForExecute(), proof.path)

	out, err := EncodeRelayReport(&single_token_offramp.CCIPRelayReport{
		MerkleRoot:        r,
		MinSequenceNumber: big.NewInt(1),
		MaxSequenceNumber: big.NewInt(10),
	})
	require.NoError(t, err)
	_, err = DecodeRelayReport(out)
	require.NoError(t, err)

	tx, err := offRamp.Report(destUser, out)
	require.NoError(t, err)
	destChain.Commit()
	res, err := destChain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	assert.Equal(t, uint64(1), res.Status)

	rep, err := offRamp.GetLastReport(nil)
	require.NoError(t, err)
	// Verify it locally
	require.True(t, bytes.Equal(rep.MerkleRoot[:], rootLocal[:]), fmt.Sprintf("Got %v want %v", hexutil.Encode(rootLocal[:]), hexutil.Encode(rep.MerkleRoot[:])))
	exists, err := offRamp.GetMerkleRoot(nil, rep.MerkleRoot)
	require.NoError(t, err)
	require.True(t, exists.Int64() > 0)

	// Verify it onchain
	lh := HashLeaf([]byte{0xaa})
	// Should merely be doing H(lhash, 32 zero bytes) and obtaining the same hash
	root, err := offRamp.GenerateMerkleRoot(nil, proof.PathForExecute(), lh, proof.Index())
	require.NoError(t, err)

	t.Log("verifies", root, "path", proof.PathForExecute(), "Index", proof.Index(), "root", rep.MerkleRoot, "rootlocal", hashInternal(lh, proof.PathForExecute()[0]))
	require.Equal(t, rootLocal, root)
}
                                                   core/services/ccip/merkle_test.go                                                                   000644  000765  000024  00000003030 14165346401 020073  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func bytes32Equal(l, r [32]byte) bool {
	return bytes.Equal(l[:], r[:])
}

func TestMerkleProver(t *testing.T) {
	zhs := computeZeroHashes(2)
	require.Equal(t, 2, len(zhs))
	var zh [32]byte
	assert.True(t, bytes32Equal(zh, zhs[0]))
	assert.True(t, bytes32Equal(hashInternal(zh, zh), zhs[1]))
	zhs = computeZeroHashes(32)
	require.Equal(t, 32, len(zhs))

	leaves := make([][]byte, 2)
	leaves[0] = []byte{0xaa}
	leaves[1] = []byte{0xbb}

	// With a tree height of 2 and 2 elements, the root should simply be
	// h(h(leaf0) || h(leaf1))
	root, proof := GenerateMerkleProof(2, leaves, 0)
	assert.True(t, bytes32Equal(root, hashInternal(HashLeaf(leaves[0]), HashLeaf(leaves[1]))))
	assert.True(t, bytes32Equal(root, GenerateMerkleRoot(leaves[0], proof)))

	// With a tree height of 3 and 2 elements, we expect
	// h((h(leaf0) || h(leaf1)) || h(0 || 0))
	root, proof = GenerateMerkleProof(3, leaves, 0)
	assert.True(t, bytes32Equal(root,
		hashInternal(hashInternal(HashLeaf(leaves[0]), HashLeaf(leaves[1])), hashInternal(zh, zh))))
	assert.True(t, bytes32Equal(root, GenerateMerkleRoot(leaves[0], proof)))

	// One element tree height 2
	root, proof = GenerateMerkleProof(2, leaves[:1], 0)
	assert.True(t, bytes32Equal(root,
		hashInternal(HashLeaf(leaves[0]), zh)))
	assert.True(t, bytes32Equal(zh, proof.path[0]))
	assert.True(t, bytes32Equal(zh, proof.path[0]))
	assert.True(t, bytes32Equal(root, GenerateMerkleRoot(leaves[0], proof)))
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        core/services/ccip/delegate_executor.go                                                             000644  000765  000024  00000015633 14165346401 021261  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         package ccip

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/services/bulletprooftxmanager"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/postgres"
	ocr "github.com/smartcontractkit/libocr/offchainreporting2"
	"github.com/smartcontractkit/libocr/offchainreporting2/chains/evmutil"
	"gorm.io/gorm"
)

var _ job.Delegate = (*ExecutionDelegate)(nil)

type ExecutionDelegate struct {
	db          *gorm.DB
	jobORM      job.ORM
	orm         ORM
	chainSet    evm.ChainSet
	keyStore    keystore.OCR2
	peerWrapper *ocrcommon.SingletonPeerWrapper
}

// TODO: Register this delegate behind a FF
func NewExecutionDelegate(
	db *gorm.DB,
	jobORM job.ORM,
	chainSet evm.ChainSet,
	keyStore keystore.OCR2,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
) *ExecutionDelegate {
	return &ExecutionDelegate{
		db:          db,
		jobORM:      jobORM,
		orm:         NewORM(postgres.UnwrapGormDB(db)),
		chainSet:    chainSet,
		keyStore:    keyStore,
		peerWrapper: peerWrapper,
	}
}

func (d ExecutionDelegate) JobType() job.Type {
	return job.CCIPExecution
}

func (d ExecutionDelegate) getOracleArgs(l logger.Logger, jb job.Job, executor *message_executor.MessageExecutor, chain evm.Chain, contractTracker *CCIPContractTracker, offchainConfigDigester evmutil.EVMOffchainConfigDigester, offRamp *single_token_offramp.SingleTokenOffRamp) (*ocr.OracleArgs, error) {
	ta, err := getTransmitterAddress(jb.CCIPExecutionSpec.TransmitterAddress, chain)
	if err != nil {
		return nil, err
	}
	executorABI, err := abi.JSON(strings.NewReader(message_executor.MessageExecutorABI))
	if err != nil {
		return nil, errors.Wrap(err, "could not get contract ABI JSON")
	}
	contractTransmitter := NewExecutionTransmitter(
		executor,
		executorABI,
		NewExecuteTransmitter(chain.TxManager(),
			d.db,
			jb.CCIPExecutionSpec.SourceEVMChainID.ToInt(),
			jb.CCIPExecutionSpec.DestEVMChainID.ToInt(), ta.Address(),
			chain.Config().EvmGasLimitDefault(),
			bulletprooftxmanager.NewQueueingTxStrategy(jb.ExternalJobID,
				chain.Config().OCR2DefaultTransactionQueueDepth(), false),
			chain.Client()),
	)
	ocrLogger := logger.NewOCRWrapper(l, true, func(msg string) {
		d.jobORM.RecordError(context.Background(), jb.ID, msg)
	})
	key, err := getValidatedKeyBundle(jb.CCIPExecutionSpec.EncryptedOCRKeyBundleID, chain, d.keyStore)
	if err != nil {
		return nil, err
	}
	if err = validatePeerWrapper(jb.CCIPExecutionSpec.P2PPeerID, chain, d.peerWrapper); err != nil {
		return nil, err
	}
	bootstrapPeers, err := getValidatedBootstrapPeers(jb.CCIPExecutionSpec.P2PBootstrapPeers, chain)
	if err != nil {
		return nil, err
	}

	gormdb, errdb := d.db.DB()
	if errdb != nil {
		return nil, errors.Wrap(errdb, "unable to open sql db")
	}
	ocrdb := NewDB(gormdb, jb.CCIPExecutionSpec.ExecutorAddress.Address())
	return &ocr.OracleArgs{
		BinaryNetworkEndpointFactory: d.peerWrapper.Peer2,
		V2Bootstrappers:              bootstrapPeers,
		ContractTransmitter:          contractTransmitter,
		ContractConfigTracker:        contractTracker,
		Database:                     ocrdb,
		LocalConfig: computeLocalConfig(chain.Config(), chain.Config().Dev(),
			jb.CCIPExecutionSpec.BlockchainTimeout.Duration(),
			jb.CCIPExecutionSpec.ContractConfigConfirmations, jb.CCIPExecutionSpec.ContractConfigTrackerPollInterval.Duration()),
		Logger:                 ocrLogger,
		MonitoringEndpoint:     nil, // TODO
		OffchainConfigDigester: offchainConfigDigester,
		OffchainKeyring:        &key.OffchainKeyring,
		OnchainKeyring:         &key.OnchainKeyring,
		ReportingPluginFactory: NewExecutionReportingPluginFactory(l, d.orm, jb.CCIPExecutionSpec.SourceEVMChainID.ToInt(), jb.CCIPExecutionSpec.DestEVMChainID.ToInt(), jb.CCIPExecutionSpec.ExecutorAddress.Address(), offRamp),
	}, nil
}

func (d ExecutionDelegate) ServicesForSpec(jb job.Job) ([]job.Service, error) {
	if jb.CCIPExecutionSpec == nil {
		return nil, errors.New("no ccip job specified")
	}
	l := logger.Default.With(
		"jobID", jb.ID,
		"externalJobID", jb.ExternalJobID,
		"offRampAddress", jb.CCIPExecutionSpec.OffRampAddress,
		"onRampAddress", jb.CCIPExecutionSpec.OnRampAddress,
		"executorAddress", jb.CCIPExecutionSpec.OnRampAddress,
	)

	destChain, err := d.chainSet.Get(jb.CCIPExecutionSpec.DestEVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open chain")
	}
	sourceChain, err := d.chainSet.Get(jb.CCIPExecutionSpec.SourceEVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open chain")
	}
	contract, err := message_executor.NewMessageExecutor(jb.CCIPExecutionSpec.ExecutorAddress.Address(), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not instantiate NewOffchainAggregator")
	}
	singleTokenOffRamp, err := single_token_offramp.NewSingleTokenOffRamp(jb.CCIPExecutionSpec.OffRampAddress.Address(), destChain.Client())
	if err != nil {
		return nil, err
	}
	offchainConfigDigester := evmutil.EVMOffchainConfigDigester{
		ChainID:         maybeRemapChainID(destChain.Config().ChainID()).Uint64(),
		ContractAddress: jb.CCIPExecutionSpec.ExecutorAddress.Address(),
	}
	contractTracker := NewCCIPContractTracker(
		executorTracker{contract},
		destChain.Client(),
		destChain.LogBroadcaster(),
		jb.ID,
		logger.Default,
		d.db,
		destChain,
		destChain.HeadBroadcaster(),
	)
	oracleArgs, err := d.getOracleArgs(l, jb, contract, destChain, contractTracker, offchainConfigDigester, singleTokenOffRamp)
	if err != nil {
		return nil, err
	}
	oracle, err := ocr.NewOracle(*oracleArgs)
	if err != nil {
		return nil, err
	}

	singleTokenOnRamp, err := single_token_onramp.NewSingleTokenOnRamp(jb.CCIPExecutionSpec.OnRampAddress.Address(), sourceChain.Client())
	if err != nil {
		return nil, err
	}

	encodedCCIPConfig, err := contractTracker.GetOffchainConfig()
	if err != nil {
		return nil, errors.Wrap(err, "could not get the latest encoded config")
	}
	// TODO: Its conceivable we may want pull out this log listener into its own job spec so to avoid repeating
	// all the log subscriptions.
	logListener := NewLogListener(l,
		sourceChain.LogBroadcaster(),
		destChain.LogBroadcaster(),
		singleTokenOnRamp,
		singleTokenOffRamp,
		encodedCCIPConfig,
		d.db,
		jb.ID)
	return []job.Service{contractTracker, oracle, logListener}, nil
}

func (d ExecutionDelegate) AfterJobCreated(spec job.Job) {
}

func (d ExecutionDelegate) BeforeJobDeleted(spec job.Job) {
}
                                                                                                     contracts/test/v0.8/ccip/pools/LockUnlockPool.test.ts                                               000644  000765  000024  00000027002 14165346401 023512  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         import hre from 'hardhat'
import { Artifact } from 'hardhat/types'
import { expect } from 'chai'
import { MockERC20, LockUnlockPool } from '../../../../typechain'
import { BigNumber, Signer } from 'ethers'
import { getUsers } from '../../../test-helpers/setup'
import { evmRevert } from '../../../test-helpers/matchers'

const { deployContract } = hre.waffle

const DEPOSIT = BigNumber.from(123)

describe('LockUnlockPool', function () {
  before(async function () {
    const users = await getUsers()
    this.roles = users.roles
  })

  beforeEach(async function () {
    const MockERC20Artifact: Artifact = await hre.artifacts.readArtifact(
      'MockERC20',
    )
    this.token = <MockERC20>(
      await deployContract(this.roles.defaultAccount, MockERC20Artifact, [
        'Test Token',
        'TEST',
        this.roles.defaultAccount.address,
        BigNumber.from(1000000),
      ])
    )

    // For testing onlyRamp
    this.onRamp = this.roles.oracleNode1
    this.offRamp = this.roles.oracleNode2
    this.token
      .connect(this.roles.defaultAccount)
      .mint(this.onRamp.address, BigNumber.from(1000000))
    this.token
      .connect(this.roles.defaultAccount)
      .mint(this.offRamp.address, BigNumber.from(1000000))
    const LockUnlockPoolArtifact: Artifact = await hre.artifacts.readArtifact(
      'LockUnlockPool',
    )
    this.pool = <LockUnlockPool>(
      await deployContract(this.roles.defaultAccount, LockUnlockPoolArtifact, [
        this.token.address,
      ])
    )
  })

  describe('#constructor', () => {
    it('is initialized correctly', async function () {
      expect(await this.pool.getToken()).to.equal(this.token.address)
      expect(await this.pool.owner()).to.equal(
        this.roles.defaultAccount.address,
      )
    })
  })

  describe('#pause', () => {
    it('owner can pause pool', async function () {
      const account = this.roles.defaultAccount
      await expect(this.pool.connect(account).pause())
        .to.emit(this.pool, 'Paused')
        .withArgs(account.address)
    })

    it('unknown account cannot pause pool', async function () {
      await expect(this.pool.connect(this.onRamp).pause()).to.be.revertedWith(
        'Only callable by owner',
      )
      await expect(this.pool.connect(this.offRamp).pause()).to.be.revertedWith(
        'Only callable by owner',
      )
      await expect(
        this.pool.connect(this.roles.stranger).pause(),
      ).to.be.revertedWith('Only callable by owner')
    })
  })

  describe('#unpause', () => {
    beforeEach(async function () {
      await this.pool.connect(this.roles.defaultAccount).pause()
    })

    it('owner can unpause pool', async function () {
      const account = this.roles.defaultAccount
      await expect(this.pool.connect(account).unpause())
        .to.emit(this.pool, 'Unpaused')
        .withArgs(account.address)
    })

    it('unknown account cannot unpause pool', async function () {
      await expect(this.pool.connect(this.onRamp).unpause()).to.be.revertedWith(
        'Only callable by owner',
      )
      await expect(
        this.pool.connect(this.offRamp).unpause(),
      ).to.be.revertedWith('Only callable by owner')
      await expect(
        this.pool.connect(this.roles.stranger).unpause(),
      ).to.be.revertedWith('Only callable by owner')
    })
  })

  describe('#lockOrBurn', () => {
    let account: Signer
    let sender: string
    let depositor: string

    describe('called by the owner', () => {
      beforeEach(async function () {
        account = this.roles.defaultAccount
        sender = await account.getAddress()
        depositor = await account.getAddress()

        await this.token.connect(account).approve(this.pool.address, DEPOSIT)
        const allowance = await this.token.allowance(
          depositor,
          this.pool.address,
        )
        await expect(allowance).to.equal(DEPOSIT)
      })

      it('can lock tokens', async function () {
        await expect(this.pool.connect(account).lockOrBurn(depositor, DEPOSIT))
          .to.emit(this.pool, 'Locked')
          .withArgs(sender, depositor, DEPOSIT)

        const poolBalance = await this.token.balanceOf(this.pool.address)
        await expect(poolBalance).to.equal(DEPOSIT)
      })

      it("can't store when paused", async function () {
        await this.pool.connect(this.roles.defaultAccount).pause()

        await evmRevert(
          this.pool.connect(account).lockOrBurn(depositor, DEPOSIT),
          'Pausable: paused',
        )
      })
    })

    describe('called by the onRamp', () => {
      beforeEach(async function () {
        account = this.onRamp
        sender = await account.getAddress()
        depositor = await account.getAddress()

        await this.token.connect(account).approve(this.pool.address, DEPOSIT)
        const allowance = await this.token.allowance(
          depositor,
          this.pool.address,
        )
        await expect(allowance).to.equal(DEPOSIT)
      })

      describe('when the onRamp is not set yet', () => {
        it('Fails with a permissions error', async function () {
          await evmRevert(
            this.pool.connect(account).lockOrBurn(depositor, DEPOSIT),
            'PermissionsError()',
          )
        })
      })

      describe('Once the onRamp is set', async function () {
        beforeEach(async function () {
          await this.pool
            .connect(this.roles.defaultAccount)
            .setOnRamp(this.onRamp.address, true)
          expect(await this.pool.isOnRamp(this.onRamp.address)).to.equal(true)
        })

        it('tokens can be locked', async function () {
          await expect(
            this.pool.connect(account).lockOrBurn(depositor, DEPOSIT),
          )
            .to.emit(this.pool, 'Locked')
            .withArgs(sender, depositor, DEPOSIT)

          const poolBalance = await this.token.balanceOf(this.pool.address)
          await expect(poolBalance).to.equal(DEPOSIT)
        })

        it("can't store when paused", async function () {
          await this.pool.connect(this.roles.defaultAccount).pause()

          await evmRevert(
            this.pool.connect(account).lockOrBurn(depositor, DEPOSIT),
            'Pausable: paused',
          )
        })
      })
    })

    it('fails when called by an unknown account', async function () {
      const account = this.roles.stranger
      const depositor = account.address
      await this.token.connect(account).approve(this.pool.address, DEPOSIT)
      await evmRevert(
        this.pool.connect(account).lockOrBurn(depositor, DEPOSIT),
        `PermissionsError()`,
      )
    })
  })

  describe('#releaseOrMint', () => {
    let account: Signer
    let sender: string
    let depositor: string
    let recipient: string

    describe('called by the offRamp', () => {
      beforeEach(async function () {
        account = this.offRamp
        sender = await account.getAddress()
        depositor = await account.getAddress()
        recipient = this.roles.stranger.address

        // Store using the onRamp first
        await this.pool
          .connect(this.roles.defaultAccount)
          .setOnRamp(this.onRamp.address, true)
        expect(await this.pool.isOnRamp(this.onRamp.address)).to.equal(true)

        await this.token
          .connect(this.onRamp)
          .approve(this.pool.address, DEPOSIT)
        await this.pool
          .connect(this.onRamp)
          .lockOrBurn(this.onRamp.address, DEPOSIT)
        // // Check pool balance
        const poolBalance = await this.token.balanceOf(this.pool.address)
        await expect(poolBalance).to.equal(DEPOSIT)
      })

      describe('when the offRamp is not set yet', () => {
        it('Fails with a permissions error', async function () {
          await evmRevert(
            this.pool.connect(account).releaseOrMint(depositor, DEPOSIT),
            'PermissionsError()',
          )
        })
      })

      describe('once the offRamp is set', () => {
        beforeEach(async function () {
          await this.pool
            .connect(this.roles.defaultAccount)
            .setOffRamp(this.offRamp.address, true)
          expect(await this.pool.isOffRamp(this.offRamp.address)).to.equal(true)
        })

        it('can release tokens', async function () {
          await expect(
            this.pool.connect(account).releaseOrMint(recipient, DEPOSIT),
          )
            .to.emit(this.pool, 'Released')
            .withArgs(sender, recipient, DEPOSIT)
          const recipientBalance = await this.token.balanceOf(recipient)
          await expect(recipientBalance).to.equal(DEPOSIT)
        })

        it("can't release tokens if paused", async function () {
          await this.pool.connect(this.roles.defaultAccount).pause()

          await evmRevert(
            this.pool.connect(account).releaseOrMint(recipient, DEPOSIT),
            'Pausable: paused',
          )
        })
      })
    })

    describe('called by the owner', () => {
      beforeEach(async function () {
        account = this.roles.defaultAccount
        sender = await account.getAddress()
        depositor = await account.getAddress()
        recipient = this.roles.stranger.address

        await this.token.connect(account).approve(this.pool.address, DEPOSIT)
        await this.pool.connect(account).lockOrBurn(depositor, DEPOSIT)
        const poolBalance = await this.token.balanceOf(this.pool.address)
        await expect(poolBalance).to.equal(DEPOSIT)
      })

      it('can release tokens', async function () {
        await expect(
          this.pool.connect(account).releaseOrMint(recipient, DEPOSIT),
        )
          .to.emit(this.pool, 'Released')
          .withArgs(sender, recipient, DEPOSIT)
        const recipientBalance = await this.token.balanceOf(recipient)
        await expect(recipientBalance).to.equal(DEPOSIT)
      })

      it("can't release tokens if paused", async function () {
        await this.pool.connect(this.roles.defaultAccount).pause()

        await evmRevert(
          this.pool.connect(account).releaseOrMint(recipient, DEPOSIT),
          'Pausable: paused',
        )
      })
    })

    it('fails when called by an unknown account', async function () {
      const account = this.roles.stranger
      const depositor = account.address
      await evmRevert(
        this.pool.connect(account).releaseOrMint(depositor, DEPOSIT),
        `PermissionsError()`,
      )
    })
  })

  describe('#setOnRamp', () => {
    let expectedRamp: string
    beforeEach(async function () {
      expectedRamp = this.roles.oracleNode3.address
    })

    it('sets the on ramp when called by the owner', async function () {
      await this.pool
        .connect(this.roles.defaultAccount)
        .setOnRamp(expectedRamp, true)
      expect(await this.pool.isOnRamp(expectedRamp)).to.equal(true)
    })

    it('reverts when called by any other account', async function () {
      await evmRevert(
        this.pool.connect(this.roles.stranger).setOnRamp(expectedRamp, true),
        'Only callable by owner',
      )
    })
  })

  describe('#setOffRamp', () => {
    let expectedRamp: string
    beforeEach(async function () {
      expectedRamp = this.roles.oracleNode3.address
    })

    it('sets the on ramp when called by the owner', async function () {
      await this.pool
        .connect(this.roles.defaultAccount)
        .setOffRamp(expectedRamp, true)
      expect(await this.pool.isOffRamp(expectedRamp)).to.equal(true)
    })

    it('reverts when called by any other account', async function () {
      await evmRevert(
        this.pool.connect(this.roles.stranger).setOffRamp(expectedRamp, true),
        'Only callable by owner',
      )
    })
  })
})
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              contracts/test/v0.8/ccip/utils/TokenLimits.test.ts                                                  000644  000765  000024  00000016645 14165346401 023075  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         import hre from 'hardhat'
import { expect } from 'chai'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { TokenLimitsHelper } from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { BigNumber } from '@ethersproject/bignumber'
import { increaseTime5Minutes } from '../../../test-helpers/helpers'
import { evmRevert } from '../../../test-helpers/matchers'

const { deployContract } = hre.waffle
let roles: Roles
let TokenLimitsHelperArtifact: Artifact
let helper: TokenLimitsHelper

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('TokenLimits library', () => {
  beforeEach(async () => {
    TokenLimitsHelperArtifact = await hre.artifacts.readArtifact(
      'TokenLimitsHelper',
    )

    helper = <TokenLimitsHelper>(
      await deployContract(roles.defaultAccount, TokenLimitsHelperArtifact)
    )
  })

  describe('#constructTokenBucket', () => {
    let rate: BigNumber
    let capacity: BigNumber

    beforeEach(async () => {
      rate = BigNumber.from('10')
      capacity = BigNumber.from('100')
    })

    it('returns an empty bucket', async () => {
      await helper
        .connect(roles.defaultAccount)
        .constructTokenBucket(rate, capacity, false)
      const bucket = await helper.s_bucket()

      expect(bucket.rate).to.equal(rate)
      expect(bucket.capacity).to.equal(capacity)
      expect(bucket.tokens).to.equal(BigNumber.from('0'))
    })

    it('returns a full bucket', async () => {
      await helper
        .connect(roles.defaultAccount)
        .constructTokenBucket(rate, capacity, true)
      const bucket = await helper.s_bucket()

      expect(bucket.rate).to.equal(rate)
      expect(bucket.capacity).to.equal(capacity)
      expect(bucket.tokens).to.equal(capacity)
    })
  })

  describe('#update', () => {
    let rate: BigNumber
    let capacity: BigNumber

    beforeEach(async () => {
      rate = BigNumber.from('1')
      capacity = BigNumber.from('1000000')
      await helper
        .connect(roles.defaultAccount)
        .constructTokenBucket(rate, capacity, false)
    })

    it('increases the tokens in the bucket over time at expected rate', async () => {
      await increaseTime5Minutes(hre.ethers.provider)
      await helper.connect(roles.defaultAccount).update()
      const bucket = await helper.s_bucket()
      const threehundred = BigNumber.from(300)
      expect(bucket.tokens)
        .to.be.at.least(threehundred)
        .but.be.below(threehundred.add(5))
    })

    it('does not update the timestamp if the bucket is already full', async () => {
      await helper
        .connect(roles.defaultAccount)
        .constructTokenBucket(rate, capacity, true)
      const initialBucket = await helper.s_bucket()
      await helper.connect(roles.defaultAccount).update()
      const updatedBucket = await helper.s_bucket()
      expect(initialBucket.lastUpdated).to.equal(updatedBucket.lastUpdated)
    })

    describe('altering the capacity', () => {
      let newCapacity: BigNumber

      beforeEach(async () => {
        await helper
          .connect(roles.defaultAccount)
          .constructTokenBucket(rate, capacity, true)
      })

      it('reverts if the capacity is reduced to below the current tokens amount', async () => {
        newCapacity = capacity.div(2)
        await helper.connect(roles.defaultAccount).alterCapacity(newCapacity)
        await evmRevert(
          helper.connect(roles.defaultAccount).update(),
          'BucketOverfilled()',
        )
      })

      it('works correctly when capacity increased', async () => {
        newCapacity = capacity.mul(2)
        const initialBucket = await helper.s_bucket()
        await helper.connect(roles.defaultAccount).alterCapacity(newCapacity)
        await helper.connect(roles.defaultAccount).update()
        const updatedBucket = await helper.s_bucket()
        expect(updatedBucket.lastUpdated).to.be.gt(initialBucket.lastUpdated)
        expect(updatedBucket.tokens).to.be.gt(capacity)
      })
    })
  })

  describe('#remove', () => {
    let rate: BigNumber
    let capacity: BigNumber

    beforeEach(async () => {
      rate = BigNumber.from('1')
      capacity = BigNumber.from('1000000')
      await helper
        .connect(roles.defaultAccount)
        .constructTokenBucket(rate, capacity, true)
    })

    it('removes from bucket', async () => {
      const removeAmount = BigNumber.from('100')
      const tx = await helper.connect(roles.defaultAccount).remove(removeAmount)
      await expect(tx).to.emit(helper, 'RemovalSuccess').withArgs(true)
      const bucket = await helper.s_bucket()
      expect(bucket.tokens).to.equal(bucket.capacity.sub(removeAmount))
    })

    it('does not remove if token amount greater than capacity', async () => {
      const removeAmount = capacity.add(1)
      const tx = await helper.connect(roles.defaultAccount).remove(removeAmount)
      await expect(tx).to.emit(helper, 'RemovalSuccess').withArgs(false)
      const bucket = await helper.s_bucket()
      expect(bucket.tokens).to.equal(bucket.capacity)
    })

    it('does not remove if token amount greater than tokens in bucket', async () => {
      const removeAmount = capacity.sub(1)
      await helper.connect(roles.defaultAccount).remove(removeAmount)
      const tx = await helper.connect(roles.defaultAccount).remove(removeAmount)
      await expect(tx).to.emit(helper, 'RemovalSuccess').withArgs(false)
    })

    describe('altering the capacity', () => {
      let newCapacity: BigNumber

      beforeEach(async () => {
        await helper
          .connect(roles.defaultAccount)
          .constructTokenBucket(rate, capacity, true)
      })

      describe('when the capacity is reduced', () => {
        beforeEach(async () => {
          newCapacity = capacity.div(2)
          await helper.connect(roles.defaultAccount).alterCapacity(newCapacity)
        })

        it('fails when the tokens amount is greater than the capacity', async () => {
          const removeAmount = BigNumber.from('100')
          await evmRevert(
            helper.connect(roles.defaultAccount).remove(removeAmount),
            'BucketOverfilled()',
          )
        })
      })

      describe('when the capacity is increased', () => {
        let initialBucket: any

        beforeEach(async () => {
          initialBucket = await helper.s_bucket()
          newCapacity = capacity.mul(2)
          await helper.connect(roles.defaultAccount).alterCapacity(newCapacity)
        })

        it('removes from the bucket', async () => {
          const removeAmount = BigNumber.from('100')
          const tx = await helper
            .connect(roles.defaultAccount)
            .remove(removeAmount)
          await expect(tx).to.emit(helper, 'RemovalSuccess').withArgs(true)
          const bucket = await helper.s_bucket()
          // Depending on provider time, the bucket might update itself
          expect(bucket.tokens)
            .to.be.at.least(initialBucket.tokens.sub(removeAmount))
            .but.be.below(initialBucket.tokens.sub(removeAmount).add(5))
        })

        it('does not remove if the token amount greater than tokens in the bucket', async () => {
          const removeAmount = newCapacity.sub(1)
          await helper.connect(roles.defaultAccount).remove(removeAmount)
          const tx = await helper
            .connect(roles.defaultAccount)
            .remove(removeAmount)
          await expect(tx).to.emit(helper, 'RemovalSuccess').withArgs(false)
        })
      })
    })
  })
})
                                                                                           contracts/test/v0.8/ccip/utils/HealthChecker.test.ts                                                000644  000765  000024  00000013222 14165346401 023311  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         import hre from 'hardhat'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { HealthCheckerHelper, MockAFN } from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { BigNumber } from '@ethersproject/bignumber'
import { expect } from 'chai'
import { evmRevert } from '../../../test-helpers/matchers'
import { constants } from 'ethers'

const { deployContract } = hre.waffle
let roles: Roles
let HealthCheckerArtifact: Artifact
let MockAFNArtifact: Artifact
let healthChecker: HealthCheckerHelper
let afn: MockAFN

let maxTimeBetweenAFNSignals: BigNumber

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('HealthChecker', () => {
  beforeEach(async () => {
    MockAFNArtifact = await hre.artifacts.readArtifact('MockAFN')
    HealthCheckerArtifact = await hre.artifacts.readArtifact(
      'HealthCheckerHelper',
    )
    maxTimeBetweenAFNSignals = BigNumber.from(60).mul(60) // 1 hour

    afn = <MockAFN>await deployContract(roles.defaultAccount, MockAFNArtifact)
    healthChecker = <HealthCheckerHelper>(
      await deployContract(roles.defaultAccount, HealthCheckerArtifact, [
        afn.address,
        maxTimeBetweenAFNSignals,
      ])
    )
  })

  describe('#constructor', () => {
    it('sets the correct storage vars', async () => {
      expect(await healthChecker.getAFN()).to.equal(afn.address)
      expect(await healthChecker.getMaxSecondsWithoutAFNHeartbeat()).to.equal(
        maxTimeBetweenAFNSignals,
      )
    })

    it('fails if zero values are used', async () => {
      // Zero address afn
      await evmRevert(
        deployContract(roles.defaultAccount, HealthCheckerArtifact, [
          constants.AddressZero,
          maxTimeBetweenAFNSignals,
        ]),
        `BadHealthConfig()`,
      )
      // Zero time
      await evmRevert(
        deployContract(roles.defaultAccount, HealthCheckerArtifact, [
          afn.address,
          0,
        ]),
        `BadHealthConfig()`,
      )
    })
  })

  describe('#pause', () => {
    it('owner can pause healthChecker', async () => {
      const account = roles.defaultAccount
      await expect(healthChecker.connect(account).pause())
        .to.emit(healthChecker, 'Paused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot pause pool', async function () {
      const account = roles.stranger
      await expect(healthChecker.connect(account).pause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#unpause', () => {
    beforeEach(async () => {
      await healthChecker.connect(roles.defaultAccount).pause()
    })

    it('owner can unpause healthChecker', async () => {
      const account = roles.defaultAccount
      await expect(healthChecker.connect(account).unpause())
        .to.emit(healthChecker, 'Unpaused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot unpause pool', async function () {
      const account = roles.stranger
      await expect(healthChecker.connect(account).unpause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#setAFN', () => {
    let newAFN: MockAFN

    beforeEach(async () => {
      newAFN = <MockAFN>(
        await deployContract(roles.defaultAccount, MockAFNArtifact)
      )
    })

    it('only callable by owner', async () => {
      await expect(
        healthChecker.connect(roles.stranger).setAFN(newAFN.address),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('fails with zero value', async () => {
      await evmRevert(
        healthChecker
          .connect(roles.defaultAccount)
          .setAFN(constants.AddressZero),
        `BadHealthConfig()`,
      )
    })

    it('sets the new AFN', async () => {
      const tx = await healthChecker
        .connect(roles.defaultAccount)
        .setAFN(newAFN.address)
      expect(await healthChecker.getAFN()).to.equal(newAFN.address)
      await expect(tx)
        .to.emit(healthChecker, 'AFNSet')
        .withArgs(afn.address, newAFN.address)
    })
  })

  describe('#setMaxTimeWithoutAFNSignal', () => {
    let newTime: BigNumber

    beforeEach(async () => {
      newTime = maxTimeBetweenAFNSignals.mul(2)
    })

    it('only callable by owner', async () => {
      await expect(
        healthChecker
          .connect(roles.stranger)
          .setMaxSecondsWithoutAFNHeartbeat(newTime),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('fails with zero value', async () => {
      await evmRevert(
        healthChecker
          .connect(roles.defaultAccount)
          .setMaxSecondsWithoutAFNHeartbeat(0),
        `BadHealthConfig()`,
      )
    })

    it('sets the new max time without afn signal', async () => {
      const tx = await healthChecker
        .connect(roles.defaultAccount)
        .setMaxSecondsWithoutAFNHeartbeat(newTime)
      expect(await healthChecker.getMaxSecondsWithoutAFNHeartbeat()).to.equal(
        newTime,
      )
      await expect(tx)
        .to.emit(healthChecker, 'AFNMaxHeartbeatTimeSet')
        .withArgs(maxTimeBetweenAFNSignals, newTime)
    })
  })

  describe('#whenHealthy', () => {
    // Uses HealthCheckerHelper.whenHealthyFunction() to simulate modifier

    it('fails if the afn has emitted a bad signal', async () => {
      await afn.voteBad()
      await evmRevert(healthChecker.whenHealthyFunction(), 'BadAFNSignal()')
    })

    it('fails if the heartbeat is stale', async () => {
      await afn.setTimestamp(1)
      await evmRevert(healthChecker.whenHealthyFunction(), 'StaleAFNHeartbeat')
    })

    it('it does nothing if all is well', async () => {
      await healthChecker.whenHealthyFunction()
    })
  })
})
                                                                                                                                                                                                                                                                                                                                                                              contracts/test/v0.8/ccip/utils/AFN.test.ts                                                          000644  000765  000024  00000027446 14165346401 021240  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         import hre from 'hardhat'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { AFN } from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { BigNumber } from '@ethersproject/bignumber'
import { Signer } from '@ethersproject/abstract-signer'
import { expect } from 'chai'
import { evmRevert } from '../../../test-helpers/matchers'
import { ContractTransaction } from 'ethers'

const { deployContract } = hre.waffle
let roles: Roles
let AFNArtifact: Artifact
let afn: AFN
let partyAccounts: Array<Signer>
let parties: Array<string>
let weights: Array<BigNumber>
let goodQuorum: BigNumber
let badQuorum: BigNumber

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('AFN', () => {
  beforeEach(async () => {
    partyAccounts = [
      roles.oracleNode1,
      roles.oracleNode2,
      roles.oracleNode3,
      roles.oracleNode4,
    ]
    parties = await Promise.all(partyAccounts.map((acc) => acc.getAddress()))
    weights = [1, 2, 3, 4].map((num) => BigNumber.from(num))
    badQuorum = BigNumber.from(3)
    goodQuorum = BigNumber.from(4)

    AFNArtifact = await hre.artifacts.readArtifact('AFN')

    afn = <AFN>(
      await deployContract(roles.defaultAccount, AFNArtifact, [
        parties,
        weights,
        goodQuorum,
        badQuorum,
      ])
    )
  })

  describe('#constructor', () => {
    it('deploys correctly', async () => {
      const initialBadSignal = await afn.hasBadSignal()
      const initialLastHeartbeat = await afn.getLastHeartbeat()
      const initialQuorums = await afn.getQuorums()
      const initialParties = await afn.getParties()
      const initialRound = await afn.getRound()
      const initialCommitteeVersion = await afn.getCommitteeVersion()
      expect(initialBadSignal).to.be.false
      expect(initialLastHeartbeat.timestamp).to.equal(0)
      expect(initialQuorums.good).to.equal(goodQuorum)
      expect(initialQuorums.bad).to.equal(badQuorum)
      expect(initialParties).to.deep.equal(parties)
      expect(initialRound).to.equal(1)
      expect(initialCommitteeVersion).to.equal(1)

      for (let i = 0; i < parties.length; i++) {
        const party = parties[i]
        const initialWeight = await afn.getWeight(party)
        expect(initialWeight).to.equal(weights[i])
      }
    })
  })

  describe('#voteGood', () => {
    describe('failure', () => {
      it('fails when the round is wrong', async () => {
        await evmRevert(
          afn.connect(partyAccounts[1]).voteGood(2),
          'IncorrectRound(1, 2)',
        )
      })
      it('fails if the signal is bad', async () => {
        await afn.connect(partyAccounts[3]).voteBad()
        await evmRevert(
          afn.connect(partyAccounts[1]).voteGood(1),
          'MustRecoverFromBadSignal',
        )
      })
      it('fails if the voter is not a registered party', async () => {
        await evmRevert(
          afn.connect(roles.defaultAccount).voteGood(1),
          `InvalidVoter("${await roles.defaultAccount.getAddress()}")`,
        )
      })
      it('fails if the voter already voted in this round', async () => {
        await afn.connect(partyAccounts[1]).voteGood(1)
        await evmRevert(
          afn.connect(partyAccounts[1]).voteGood(1),
          `AlreadyVoted()`,
        )
      })
    })

    describe('success', () => {
      let tx: ContractTransaction
      let index: number
      describe('single vote without reaching quorum', () => {
        beforeEach(async () => {
          index = 1
          tx = await afn.connect(partyAccounts[index]).voteGood(1)
        })
        it('sets the last good vote of the voter to this round', async () => {
          const lastGoodVote = await afn.getLastGoodVote(parties[index])
          expect(lastGoodVote).to.equal(1)
        })
        it('adds the votes to this round', async () => {
          const votes = await afn.getGoodVotes(1)
          expect(votes).to.equal(weights[index])
        })
        it('emits a good vote event', async () => {
          await expect(tx).to.emit(afn, 'GoodVote').withArgs(parties[index], 1)
        })
      })

      describe('reaching good quorum', () => {
        beforeEach(async () => {
          index = 3
          tx = await afn.connect(partyAccounts[index]).voteGood(1)
        })

        it('sets the last heartbeat', async () => {
          const heartbeat = await afn.getLastHeartbeat()
          expect(heartbeat.round).to.equal(1)
          expect(heartbeat.committeeVersion).to.equal(1)
          expect(heartbeat.timestamp).to.not.equal(0)
        })
        it('increments the round', async () => {
          const roundSet = await afn.getRound()
          expect(roundSet).to.equal(2)
        })
        it('emits a heatbeat event', async () => {
          await expect(tx).to.emit(afn, 'AFNHeartbeat')
        })
      })
    })
  })

  describe('#voteBad', () => {
    describe('failure', () => {
      it('fails if the signal is already bad', async () => {
        await afn.connect(partyAccounts[3]).voteBad()
        await evmRevert(
          afn.connect(partyAccounts[2]).voteBad(),
          'MustRecoverFromBadSignal()',
        )
      })
      it('fails if the voter is not a registered party', async () => {
        await evmRevert(
          afn.connect(roles.defaultAccount).voteBad(),
          `InvalidVoter("${await roles.defaultAccount.getAddress()}")`,
        )
      })
      it('fails is the voter has already voted bad', async () => {
        await afn.connect(partyAccounts[0]).voteBad()
        await evmRevert(
          afn.connect(partyAccounts[0]).voteBad(),
          'AlreadyVoted()',
        )
      })
    })

    describe('success', () => {
      it('increments votes, adds party to voters and sets s_hasVotedBad for sender', async () => {
        const index = 1
        await afn.connect(partyAccounts[index]).voteBad()
        const votersAndVotes = await afn.getBadVotersAndVotes()
        const hasVotedBad = await afn.hasVotedBad(parties[index])
        expect(votersAndVotes.voters).to.deep.equal([parties[index]])
        expect(votersAndVotes.votes).to.equal(weights[index])
        expect(hasVotedBad).to.be.true
      })
      describe('reaching bad quorum', () => {
        let tx: ContractTransaction
        beforeEach(async () => {
          tx = await afn.connect(partyAccounts[3]).voteBad()
        })

        it('sets the bad signal', async () => {
          expect(await afn.hasBadSignal()).to.be.true
        })
        it('emits an event', async () => {
          await expect(tx).to.emit(afn, 'AFNBadSignal')
        })
      })
    })
  })

  describe('#recover', () => {
    describe('failure', () => {
      it('only allows the owner to call', async () => {
        await evmRevert(
          afn.connect(roles.stranger).recover(),
          'Only callable by owner',
        )
      })
      it('fails if there is no bad signal', async () => {
        await evmRevert(
          afn.connect(roles.defaultAccount).recover(),
          'RecoveryNotNecessary()',
        )
      })
    })

    describe('success', () => {
      let tx: ContractTransaction
      beforeEach(async () => {
        await afn.connect(partyAccounts[3]).voteBad()
        tx = await afn.connect(roles.defaultAccount).recover()
      })

      it('resets s_badVoters, s_hasVotedBad and s_badVotes', async () => {
        const votersAndVotes = await afn.getBadVotersAndVotes()
        expect(votersAndVotes.voters.length).to.equal(0)
        expect(votersAndVotes.votes).to.equal(0)
      })
      it('turns off the bad signal', async () => {
        const hasBadSignal = await afn
          .connect(roles.defaultAccount)
          .hasBadSignal()
        expect(hasBadSignal).to.be.false
      })
      it('emits a Recovered event', async () => {
        await expect(tx).to.emit(afn, 'RecoveredFromBadSignal')
      })
    })
  })

  describe('#setConfig', () => {
    let newParties: Array<string>
    let newWeights: Array<BigNumber>
    let newGoodQuorum: BigNumber
    let newBadQuorum: BigNumber

    describe('failure', () => {
      beforeEach(async () => {
        newParties = [
          await roles.consumer.getAddress(),
          await roles.stranger.getAddress(),
        ]
        newWeights = [BigNumber.from(8), BigNumber.from(9)]
        newGoodQuorum = BigNumber.from(10)
        newBadQuorum = BigNumber.from(8)
      })

      it('only allows the owner to set config', async () => {
        await evmRevert(
          afn.connect(partyAccounts[0]).setConfig([], [], 1, 1),
          'Only callable by owner',
        )
      })

      it('fails if the parties length is 0', async () => {
        await evmRevert(
          afn.connect(roles.defaultAccount).setConfig([], newWeights, 1, 1),
          'InvalidConfig()',
        )
      })
      it('fails if the weights length is 0', async () => {
        await evmRevert(
          afn.connect(roles.defaultAccount).setConfig(newParties, [], 1, 1),
          'InvalidConfig()',
        )
      })
      it('fails if the goodQuorum is 0', async () => {
        await evmRevert(
          afn
            .connect(roles.defaultAccount)
            .setConfig(newParties, newWeights, 0, 1),
          'InvalidConfig()',
        )
      })
      it('fails if the badQuorum is 0', async () => {
        await evmRevert(
          afn
            .connect(roles.defaultAccount)
            .setConfig(newParties, newWeights, 1, 0),
          'InvalidConfig()',
        )
      })
      it('fails if a weight is 0', async () => {
        await evmRevert(
          afn.connect(roles.defaultAccount).setConfig(newParties, [0, 0], 1, 1),
          'InvalidWeight()',
        )
      })
    })
    describe('success', () => {
      let tx: ContractTransaction
      let initialRound: BigNumber
      let initialCommitteeVersion: BigNumber

      beforeEach(async () => {
        initialRound = await afn.getRound()
        initialCommitteeVersion = await afn.getCommitteeVersion()

        newParties = [
          await roles.consumer.getAddress(),
          await roles.stranger.getAddress(),
        ]
        newWeights = [BigNumber.from(8), BigNumber.from(9)]
        newGoodQuorum = BigNumber.from(10)
        newBadQuorum = BigNumber.from(8)
        tx = await afn
          .connect(roles.defaultAccount)
          .setConfig(newParties, newWeights, newGoodQuorum, newBadQuorum)
      })

      it('removes the old configs', async () => {
        for (let i = 0; i < parties.length; i++) {
          const party = parties[i]
          const setWeight = await afn
            .connect(roles.defaultAccount)
            .getWeight(party)
          expect(setWeight).to.equal(0)
        }
        const quorums = await afn.getQuorums()
        expect(quorums.good).to.not.equal(goodQuorum)
        expect(quorums.bad).to.not.equal(badQuorum)
        const setRound = await afn.getRound()
        const setCommitteeVersion = await afn.getCommitteeVersion()
        expect(setRound).to.not.equal(initialRound)
        expect(setCommitteeVersion).to.not.equal(initialCommitteeVersion)
      })

      it('sets the new configs', async () => {
        for (let i = 0; i < newParties.length; i++) {
          const party = newParties[i]
          const setWeight = await afn
            .connect(roles.defaultAccount)
            .getWeight(party)
          expect(setWeight).to.equal(newWeights[i])
        }
        const quorums = await afn.getQuorums()
        expect(quorums.good).to.not.equal(goodQuorum)
        expect(quorums.bad).to.not.equal(badQuorum)
        const setRound = await afn.getRound()
        const setCommitteeVersion = await afn.getCommitteeVersion()
        expect(setRound).to.equal(initialRound.add(1))
        expect(setCommitteeVersion).to.equal(initialCommitteeVersion.add(1))
      })

      it('emits an event', async () => {
        await expect(tx)
          .to.emit(afn, 'ConfigSet')
          .withArgs(newParties, newWeights, newGoodQuorum, newBadQuorum)
      })
    })
  })
})
                                                                                                                                                                                                                          contracts/test/v0.8/ccip/ramps/SingleTokenOnRamp.test.ts                                            000644  000765  000024  00000035660 14165346401 024152  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         import hre from 'hardhat'
import { publicAbi } from '../../../test-helpers/helpers'
import { expect } from 'chai'
import { BigNumber, ContractReceipt, ContractTransaction, Signer } from 'ethers'
import { Roles, getUsers } from '../../../test-helpers/setup'
import {
  MockERC20,
  LockUnlockPool,
  SingleTokenOnRamp,
  MockAFN,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { evmRevert } from '../../../test-helpers/matchers'
import {
  CCIPMessagePayload,
  requestEventArgsEqual,
} from '../../../test-helpers/ccip'

const { deployContract } = hre.waffle

let roles: Roles

let afn: MockAFN
let ramp: SingleTokenOnRamp
let token: MockERC20
let pool: LockUnlockPool
let destinationTokenAddress: string

let MockAFNArtifact: Artifact
let TokenArtifact: Artifact
let PoolArtifact: Artifact
let RampArtifact: Artifact

const sourceChainId: number = 123
const destinationChainId: number = 234
let bucketRate: BigNumber
let bucketCapactiy: BigNumber
let maxTimeWithoutAFNSignal: BigNumber

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('SingleTokenOnRamp', () => {
  beforeEach(async () => {
    destinationTokenAddress = await roles.stranger.getAddress()

    bucketRate = BigNumber.from('10000000000000000')
    bucketCapactiy = BigNumber.from('100000000000000000')

    MockAFNArtifact = await hre.artifacts.readArtifact('MockAFN')
    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')
    PoolArtifact = await hre.artifacts.readArtifact('LockUnlockPool')
    RampArtifact = await hre.artifacts.readArtifact('SingleTokenOnRamp')

    afn = <MockAFN>await deployContract(roles.defaultAccount, MockAFNArtifact)
    maxTimeWithoutAFNSignal = BigNumber.from(60).mul(60) // 1 hour
    token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'LINK Token',
        'LINK',
        await roles.defaultAccount.getAddress(),
        BigNumber.from('1000000000000000000'),
      ])
    )
    pool = <LockUnlockPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [token.address])
    )
    ramp = <SingleTokenOnRamp>(
      await deployContract(roles.defaultAccount, RampArtifact, [
        sourceChainId,
        token.address,
        pool.address,
        destinationChainId,
        destinationTokenAddress,
        [roles.defaultAccount.getAddress()],
        true,
        bucketRate,
        bucketCapactiy,
        afn.address,
        maxTimeWithoutAFNSignal,
      ])
    )
    await pool.connect(roles.defaultAccount).setOnRamp(ramp.address, true)
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(ramp, [
      // SingleTokenRamp
      'requestCrossChainSend',
      'TOKEN',
      'DESTINATION_TOKEN',
      'POOL',
      'DESTINATION_CHAIN_ID',
      'CHAIN_ID',
      'setAllowlistEnabled',
      'getAllowlistEnabled',
      'setAllowlist',
      'getAllowlist',
      'configureTokenBucket',
      'getTokenBucket',
      // HealthChecker
      'setAFN',
      'getAFN',
      'setMaxSecondsWithoutAFNHeartbeat',
      'getMaxSecondsWithoutAFNHeartbeat',
      // TypeAndVersionInterface
      'typeAndVersion',
      // Ownership
      'owner',
      'transferOwnership',
      'acceptOwnership',
      // Pausable
      'paused',
      'pause',
      'unpause',
    ])
  })

  describe('#constructor', () => {
    it('should deploy correctly', async () => {
      const owner = await roles.defaultAccount.getAddress()
      await expect(await ramp.TOKEN()).to.equal(token.address)
      await expect(await ramp.POOL()).to.not.equal(
        hre.ethers.constants.AddressZero,
      )
      await expect(await ramp.DESTINATION_CHAIN_ID()).to.equal(
        destinationChainId,
      )
      await expect(await ramp.owner()).to.equal(owner)
      await expect(await ramp.getAllowlistEnabled()).to.be.true
      await expect(await ramp.getAllowlist()).to.deep.equal([
        await roles.defaultAccount.getAddress(),
      ])
      await expect(await ramp.getAFN()).to.equal(afn.address)
      await expect(await ramp.getMaxSecondsWithoutAFNHeartbeat()).to.equal(
        maxTimeWithoutAFNSignal,
      )
      const tokenBucket = await ramp.getTokenBucket()
      await expect(tokenBucket.rate).to.equal(bucketRate)
      await expect(tokenBucket.capacity).to.equal(bucketCapactiy)

      await expect(await pool.owner()).to.equal(owner)
      await expect(await pool.isOnRamp(ramp.address)).to.equal(true)
      await expect(await pool.getToken()).to.equal(token.address)
    })

    it('should fail if the pool token is different from the ramp token', async () => {
      const differentToken = <MockERC20>(
        await deployContract(roles.defaultAccount, TokenArtifact, [
          'LINK Token',
          'LINK',
          await roles.defaultAccount.getAddress(),
          BigNumber.from('1000000000000000000'),
        ])
      )
      await evmRevert(
        deployContract(roles.defaultAccount, RampArtifact, [
          sourceChainId,
          differentToken.address,
          pool.address,
          destinationChainId,
          destinationTokenAddress,
          [roles.defaultAccount.getAddress()],
          true,
          bucketRate,
          bucketCapactiy,
          afn.address,
          maxTimeWithoutAFNSignal,
        ]),
        `TokenMismatch()`,
      )
    })
  })

  describe('#requestCrossChainSend', () => {
    let receiver: string
    let messagedata: string
    let options: string
    let amount: BigNumber
    let message: CCIPMessagePayload

    beforeEach(async () => {
      receiver = await roles.stranger.getAddress()
      messagedata = hre.ethers.constants.HashZero
      options = hre.ethers.constants.HashZero
      amount = BigNumber.from('1000000000000000')
      message = {
        receiver: receiver,
        data: messagedata,
        tokens: [token.address],
        amounts: [amount],
        executor: hre.ethers.constants.AddressZero,
        options: options,
      }
    })

    describe('when contract not paused', () => {
      it('fails if there are not enough or too many tokens', async () => {
        message.tokens = []
        await evmRevert(
          ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
          `UnsupportedNumberOfTokens()`,
        )
        message.tokens = [token.address, token.address]
        await evmRevert(
          ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
          `UnsupportedNumberOfTokens()`,
        )
      })
      it('fails if there are not enough or too many amounts', async () => {
        message.amounts = []
        await evmRevert(
          ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
          `UnsupportedNumberOfTokens()`,
        )
        message.amounts = [amount, amount]
        await evmRevert(
          ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
          `UnsupportedNumberOfTokens()`,
        )
      })
      it('fails if token is not configured token', async () => {
        message.tokens = [receiver]
        await evmRevert(
          ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
          `UnsupportedToken("${token.address}", "${receiver}")`,
        )
      })

      it('fails if sent by a non-allowlisted address', async () => {
        await evmRevert(
          ramp.connect(roles.stranger).requestCrossChainSend(message),
          `SenderNotAllowed("${await roles.stranger.getAddress()}")`,
        )
      })

      it('fails if the send amount is greater than the bucket allows', async () => {
        message.amounts = [bucketCapactiy.add(1)]
        await evmRevert(
          ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
          `ExceedsTokenLimit(${bucketCapactiy}, ${message.amounts[0]})`,
        )
      })

      describe('sending a message', () => {
        let tx: ContractTransaction
        let owner: Signer
        let initialPoolBalance: BigNumber

        beforeEach(async () => {
          owner = roles.defaultAccount
          initialPoolBalance = await token.balanceOf(pool.address)
          await token.approve(pool.address, amount)
          tx = await ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(message)
        })

        it('emits a Locked event in the pool', async () => {
          await expect(tx)
            .to.emit(pool, 'Locked')
            .withArgs(ramp.address, await owner.getAddress(), amount)
        })

        it('transfers the tokens to the pool', async () => {
          const expectedBalance = initialPoolBalance.add(amount)
          await expect(await token.balanceOf(pool.address)).to.equal(
            expectedBalance,
          )
        })

        it('emits a message send request', async () => {
          const receipt: ContractReceipt = await tx.wait()
          const eventArgs = receipt.events?.[3]?.args?.[0]
          requestEventArgsEqual(eventArgs, {
            sequenceNumber: eventArgs?.sequenceNumber,
            sourceChainId: BigNumber.from(sourceChainId),
            destinationChainId: BigNumber.from(destinationChainId),
            sender: await owner.getAddress(),
            receiver: receiver,
            data: messagedata,
            tokens: [destinationTokenAddress],
            amounts: [amount],
            options: options,
          })
        })
      })
    })

    it('fails when the ramp is paused', async () => {
      await ramp.pause()
      await evmRevert(
        ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
        'Pausable: paused',
      )
    })

    it('fails whenn the AFN signal is bad', async () => {
      await afn.voteBad()
      await evmRevert(
        ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
        'BadAFNSignal()',
      )
    })

    it('fails when the AFN signal is stale', async () => {
      await afn.setTimestamp(BigNumber.from(1))
      await evmRevert(
        ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
        'StaleAFNHeartbeat()',
      )
    })
  })

  describe('#pause', () => {
    it('owner can pause ramp', async () => {
      const account = roles.defaultAccount
      await expect(ramp.connect(account).pause())
        .to.emit(ramp, 'Paused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot pause pool', async function () {
      const account = roles.stranger
      await expect(ramp.connect(account).pause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#unpause', () => {
    beforeEach(async () => {
      await ramp.connect(roles.defaultAccount).pause()
    })

    it('owner can unpause ramp', async () => {
      const account = roles.defaultAccount
      await expect(ramp.connect(account).unpause())
        .to.emit(ramp, 'Unpaused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot unpause pool', async function () {
      const account = roles.stranger
      await expect(ramp.connect(account).unpause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#setAllowlistEnabled', () => {
    it('only allows owner to set', async () => {
      await evmRevert(
        ramp.connect(roles.stranger).setAllowlistEnabled(false),
        'Only callable by owner',
      )
    })

    it('sets the allowlistEnabled flag correctly', async () => {
      let tx = await ramp
        .connect(roles.defaultAccount)
        .setAllowlistEnabled(false)
      await expect(await ramp.getAllowlistEnabled()).to.be.false
      await expect(tx).to.emit(ramp, 'AllowlistEnabledSet').withArgs(false)

      tx = await ramp.connect(roles.defaultAccount).setAllowlistEnabled(true)
      await expect(await ramp.getAllowlistEnabled()).to.be.true
      await expect(tx).to.emit(ramp, 'AllowlistEnabledSet').withArgs(true)
    })
  })

  describe('#setAllowlist', () => {
    let newAllowList: Array<string>

    beforeEach(async () => {
      newAllowList = [
        await roles.oracleNode1.getAddress(),
        await roles.oracleNode2.getAddress(),
      ]
    })

    it('only allows owner to set', async () => {
      await evmRevert(
        ramp.connect(roles.stranger).setAllowlist(newAllowList),
        'Only callable by owner',
      )
    })

    it('sets the correct allowlist', async () => {
      await ramp.connect(roles.defaultAccount).setAllowlist(newAllowList)
      await expect(await ramp.getAllowlist()).to.deep.equal(newAllowList)
    })
  })

  describe('#setAFN', () => {
    let newAFN: MockAFN

    beforeEach(async () => {
      newAFN = <MockAFN>(
        await deployContract(roles.defaultAccount, MockAFNArtifact)
      )
    })

    it('only callable by owner', async () => {
      await expect(
        ramp.connect(roles.stranger).setAFN(newAFN.address),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('sets the new AFN', async () => {
      const tx = await ramp.connect(roles.defaultAccount).setAFN(newAFN.address)
      expect(await ramp.getAFN()).to.equal(newAFN.address)
      await expect(tx)
        .to.emit(ramp, 'AFNSet')
        .withArgs(afn.address, newAFN.address)
    })
  })

  describe('#setMaxSecondsWithoutAFNHeartbeat', () => {
    let newTime: BigNumber

    beforeEach(async () => {
      newTime = maxTimeWithoutAFNSignal.mul(2)
    })

    it('only callable by owner', async () => {
      await expect(
        ramp.connect(roles.stranger).setMaxSecondsWithoutAFNHeartbeat(newTime),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('sets the new max time without afn signal', async () => {
      const tx = await ramp
        .connect(roles.defaultAccount)
        .setMaxSecondsWithoutAFNHeartbeat(newTime)
      expect(await ramp.getMaxSecondsWithoutAFNHeartbeat()).to.equal(newTime)
      await expect(tx)
        .to.emit(ramp, 'AFNMaxHeartbeatTimeSet')
        .withArgs(maxTimeWithoutAFNSignal, newTime)
    })
  })

  describe('#configureTokenBucket', () => {
    let newRate: BigNumber
    let newCapacity: BigNumber

    beforeEach(async () => {
      newRate = BigNumber.from(5)
      newCapacity = bucketCapactiy.add(10)
    })

    it('only callable by owner', async () => {
      await expect(
        ramp
          .connect(roles.stranger)
          .configureTokenBucket(newRate, newCapacity, true),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('sets the new max time without afn signal', async () => {
      const tx = await ramp
        .connect(roles.defaultAccount)
        .configureTokenBucket(newRate, newCapacity, true)
      const tokenBucketParams = await ramp.getTokenBucket()
      expect(tokenBucketParams.rate).to.equal(newRate)
      expect(tokenBucketParams.capacity).to.equal(newCapacity)
      await expect(tx)
        .to.emit(ramp, 'NewTokenBucketConstructed')
        .withArgs(newRate, newCapacity, true)
    })
  })

  describe('#typeAndVersion', () => {
    it('should return the correct type and version', async () => {
      const response = await ramp.typeAndVersion()
      await expect(response).to.equal('SingleTokenOnRamp 1.1.0')
    })
  })
})
                                                                                contracts/test/v0.8/ccip/ramps/SingleTokenOffRamp.test.ts                                           000644  000765  000024  00000070416 14165346401 024306  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         import hre from 'hardhat'
import {
  numToBytes32,
  publicAbi,
  stringToBytes,
} from '../../../test-helpers/helpers'
import { expect } from 'chai'
import {
  BigNumber,
  Contract,
  ContractFactory,
  ContractTransaction,
} from 'ethers'
import { Roles, getUsers } from '../../../test-helpers/setup'
import {
  SimpleMessageReceiver,
  MockERC20,
  LockUnlockPool,
  MockAFN,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { evmRevert } from '../../../test-helpers/matchers'
import {
  CCIPMessage,
  CCIPMessagePayload,
  encodeReport,
  generateMerkleTreeFromHashes,
  hashMessage,
  MerkleTree,
  messageDeepEqual,
  RelayReport,
} from '../../../test-helpers/ccip'

const { deployContract } = hre.waffle

function constructReport(
  message: CCIPMessage,
  minSequenceNumber: BigNumber,
  maxSequenceNumber: BigNumber,
): RelayReport {
  const rootHash = hashMessage(message)
  let report: RelayReport = {
    merkleRoot: rootHash,
    minSequenceNumber: minSequenceNumber,
    maxSequenceNumber: maxSequenceNumber,
  }
  return report
}

let roles: Roles

// This has to be ethers.Contract because of an issue with
// `address.call(abi.encodeWithSelector(...))` using typechain artifacts.
let ramp: Contract
let afn: MockAFN
let token: MockERC20
let receiver: SimpleMessageReceiver
let pool: LockUnlockPool

let MockAFNArtifact: Artifact
let TokenArtifact: Artifact
let PoolArtifact: Artifact
let rampFactory: ContractFactory

const sourceChainId: number = 123
const destinationChainId: number = 234
const initialExecutionDelay: number = 0
let bucketRate: BigNumber
let bucketCapactiy: BigNumber
let maxTimeBetweenAFNSignals: BigNumber

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('SingleTokenOffRamp', () => {
  beforeEach(async () => {
    MockAFNArtifact = await hre.artifacts.readArtifact('MockAFN')
    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')
    PoolArtifact = await hre.artifacts.readArtifact('LockUnlockPool')
    rampFactory = await hre.ethers.getContractFactory(
      'SingleTokenOffRampHelper',
    )
    const SimpleMessageReceiverArtifact: Artifact =
      await hre.artifacts.readArtifact('SimpleMessageReceiver')
    bucketRate = BigNumber.from('10000000000000000')
    bucketCapactiy = BigNumber.from('100000000000000000')
    const mintAmount = BigNumber.from('100000000000000000000')
    maxTimeBetweenAFNSignals = BigNumber.from(60).mul(60) // 1 hour
    token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'LINK Token',
        'LINK',
        await roles.defaultAccount.getAddress(),
        mintAmount,
      ])
    )
    pool = <LockUnlockPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [token.address])
    )
    await token
      .connect(roles.defaultAccount)
      .transfer(pool.address, mintAmount.div(2))
    afn = <MockAFN>await deployContract(roles.defaultAccount, MockAFNArtifact)
    ramp = await rampFactory
      .connect(roles.defaultAccount)
      .deploy(
        sourceChainId,
        destinationChainId,
        token.address,
        pool.address,
        bucketRate,
        bucketCapactiy,
        afn.address,
        maxTimeBetweenAFNSignals,
        initialExecutionDelay,
      )
    await pool.connect(roles.defaultAccount).setOffRamp(ramp.address, true)
    receiver = <SimpleMessageReceiver>(
      await deployContract(roles.defaultAccount, SimpleMessageReceiverArtifact)
    )
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(ramp, [
      // SingleTokenRamp
      'TOKEN',
      'POOL',
      'SOURCE_CHAIN_ID',
      'CHAIN_ID',
      'executeTransaction',
      'generateMerkleRoot',
      'getMerkleRoot',
      'getExecuted',
      'getLastReport',
      'getExecutionDelaySeconds',
      'setExecutionDelaySeconds',
      'configureTokenBucket',
      'getTokenBucket',
      // HealthChecker
      'setAFN',
      'getAFN',
      'setMaxSecondsWithoutAFNHeartbeat',
      'getMaxSecondsWithoutAFNHeartbeat',
      //SingleTokenOffRampHelper
      'report',
      // OCR2Base
      'setConfig',
      'latestConfigDetails',
      'transmitters',
      'transmit',
      // TypeAndVersionInterface
      'typeAndVersion',
      // Ownership
      'owner',
      'transferOwnership',
      'acceptOwnership',
      // Pausable
      'paused',
      'pause',
      'unpause',
    ])
  })

  describe('#constructor', () => {
    it('should deploy correctly', async () => {
      const owner = await roles.defaultAccount.getAddress()
      await expect(await ramp.TOKEN()).to.equal(token.address)
      await expect(await ramp.POOL()).to.not.equal(
        hre.ethers.constants.AddressZero,
      )
      await expect(await ramp.SOURCE_CHAIN_ID()).to.equal(sourceChainId)
      await expect(await ramp.owner()).to.equal(owner)
      await expect(await ramp.getExecutionDelaySeconds()).to.equal(0)

      await expect(await pool.owner()).to.equal(owner)
      await expect(await pool.isOffRamp(ramp.address)).to.equal(true)
      await expect(await pool.getToken()).to.equal(token.address)
    })

    it('should fail if the pool token is different from the ramp token', async () => {
      const differentToken = <MockERC20>(
        await deployContract(roles.defaultAccount, TokenArtifact, [
          'LINK Token',
          'LINK',
          await roles.defaultAccount.getAddress(),
          BigNumber.from('100000000000000000000'),
        ])
      )
      await evmRevert(
        rampFactory
          .connect(roles.defaultAccount)
          .deploy(
            sourceChainId,
            destinationChainId,
            differentToken.address,
            pool.address,
            bucketRate,
            bucketCapactiy,
            afn.address,
            maxTimeBetweenAFNSignals,
            initialExecutionDelay,
          ),
        `TokenMismatch()`,
      )
    })
  })

  describe('#generateMerkleProof', () => {
    let messages: Array<string>
    let merkle: any

    it('generates', async () => {
      messages = [
        hre.ethers.utils.defaultAbiCoder.encode(['uint256'], [1]),
        hre.ethers.utils.defaultAbiCoder.encode(['uint256'], [2]),
        hre.ethers.utils.defaultAbiCoder.encode(['uint256'], [3]),
        hre.ethers.utils.defaultAbiCoder.encode(['uint256'], [4]),
      ]
      merkle = generateMerkleTreeFromHashes(messages)
      for (let i = 0; i < merkle.leaves.length; i++) {
        const leaf = merkle.leaves[i]
        const proof = leaf.recursiveProof([])
        const hash = leaf.hash
        expect(await ramp.generateMerkleRoot(proof, hash, i)).to.equal(
          merkle.root.hash,
        )
      }
    })
  })

  describe('#report', () => {
    describe('failure', () => {
      let report: RelayReport
      beforeEach(async () => {
        report = {
          merkleRoot: numToBytes32(1),
          minSequenceNumber: BigNumber.from(2),
          maxSequenceNumber: BigNumber.from(3),
        }
      })

      it('reverts when paused', async () => {
        await ramp.connect(roles.defaultAccount).pause()
        await evmRevert(
          ramp.connect(roles.defaultAccount).report(stringToBytes('')),
          'Pausable: paused',
        )
      })

      it('fails whenn the AFN signal is bad', async () => {
        await afn.voteBad()
        await evmRevert(
          ramp.connect(roles.defaultAccount).report(stringToBytes('')),
          'BadAFNSignal()',
        )
      })

      it('fails when the AFN signal is stale', async () => {
        await afn.setTimestamp(BigNumber.from(1))
        await evmRevert(
          ramp.connect(roles.defaultAccount).report(stringToBytes('')),
          'StaleAFNHeartbeat()',
        )
      })

      it('reverts when the minSequenceNumber is greater than the maxSequenceNumber', async () => {
        report.maxSequenceNumber = BigNumber.from(1)
        await evmRevert(
          ramp.connect(roles.defaultAccount).report(encodeReport(report)),
          'RelayReportError()',
        )
      })

      it('reverts when the minSequenceNumber is not 1 greater than the previous report maxSequenceNumber', async () => {
        await ramp.connect(roles.defaultAccount).report(encodeReport(report))
        report = {
          merkleRoot: numToBytes32(2),
          minSequenceNumber: BigNumber.from(3),
          maxSequenceNumber: BigNumber.from(4),
        }
        await evmRevert(
          ramp.connect(roles.defaultAccount).report(encodeReport(report)),
          `SequenceError(3, 3)`,
        )
      })
    })

    describe('success', () => {
      let report: RelayReport
      let root: string
      let response: ContractTransaction
      beforeEach(async () => {
        root = numToBytes32(1)
        report = {
          merkleRoot: root,
          minSequenceNumber: BigNumber.from(1),
          maxSequenceNumber: BigNumber.from(2),
        }
        response = await ramp
          .connect(roles.defaultAccount)
          .report(encodeReport(report))
      })
      it('stores the root', async () => {
        const stored = await ramp.getMerkleRoot(root)
        expect(stored).to.not.equal(0)
      })
      it('stores the report in s_lastReport', async () => {
        const response = await ramp.getLastReport()
        expect(response.merkleRoot).to.equal(root)
        expect(response.minSequenceNumber).to.equal(report.minSequenceNumber)
        expect(response.maxSequenceNumber).to.equal(report.maxSequenceNumber)
      })
      it('emits a ReportAccepted event', async () => {
        expect(response)
          .to.emit(ramp, 'ReportAccepted')
          .withArgs([root, report.minSequenceNumber, report.maxSequenceNumber])
      })
    })
  })

  describe('#executeTransaction', () => {
    let sequenceNumber: BigNumber
    let sourceId: BigNumber
    let destinationId: BigNumber
    let sender: string
    let messagedata: string
    let amount: BigNumber
    let options: string
    let message: CCIPMessage
    let payload: CCIPMessagePayload
    let report: RelayReport
    beforeEach(async () => {
      sequenceNumber = BigNumber.from(1)
      sourceId = BigNumber.from(sourceChainId)
      destinationId = BigNumber.from(destinationChainId)
      sender = await roles.oracleNode.getAddress()
      messagedata = stringToBytes('Message')
      amount = BigNumber.from('10000000000')
      options = stringToBytes('options')
      payload = {
        receiver: receiver.address,
        data: messagedata,
        tokens: [token.address],
        amounts: [amount],
        executor: hre.ethers.constants.AddressZero,
        options: options,
      }
      message = {
        sequenceNumber: sequenceNumber,
        sourceChainId: sourceId,
        destinationChainId: destinationId,
        sender: sender,
        payload: payload,
      }
    })

    describe('failure', () => {
      describe('verifyMerkleProof failures', () => {
        let hashes: string[]
        let root: MerkleTree
        let leaves: MerkleTree[]

        beforeEach(async () => {
          const hash1 = hashMessage(message)
          const sequenceNumber2 = BigNumber.from(2)
          const payload2 = {
            receiver: receiver.address,
            data: messagedata,
            tokens: [token.address],
            amounts: [BigNumber.from('9999999')],
            executor: hre.ethers.constants.AddressZero,
            options: options,
          }
          const message2 = {
            sequenceNumber: sequenceNumber2,
            sourceChainId: sourceId,
            destinationChainId: destinationId,
            sender: sender,
            payload: payload2,
          }
          const hash2 = hashMessage(message2)
          hashes = [hash1, hash2]
          const merkle = generateMerkleTreeFromHashes(hashes)
          root = merkle.root
          leaves = merkle.leaves
          report = {
            merkleRoot: root.hash!,
            minSequenceNumber: sequenceNumber,
            maxSequenceNumber: sequenceNumber2,
          }
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
        })

        it('fails when the payload is wrong', async () => {
          const proof = leaves[0].recursiveProof([])
          message.payload.options = stringToBytes('loremipsum')
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(proof, message, 0),
            `MerkleProofError(["${proof[0]}"], [${message.sequenceNumber}, ${message.sourceChainId}, ${message.destinationChainId}, "${message.sender}", ["${message.payload.receiver}", "${message.payload.data}", ["${message.payload.tokens[0]}"], [${message.payload.amounts[0]}], "${message.payload.executor}", "${message.payload.options}"]], 0)`,
          )
        })

        it('fails when the proof is wrong', async () => {
          const proof = [numToBytes32(1)]
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(proof, message, 0),
            `MerkleProofError(["${proof[0]}"], [${message.sequenceNumber}, ${message.sourceChainId}, ${message.destinationChainId}, "${message.sender}", ["${message.payload.receiver}", "${message.payload.data}", ["${message.payload.tokens[0]}"], [${message.payload.amounts[0]}], "${message.payload.executor}", "${message.payload.options}"]], 0)`,
          )
        })

        it('fails when the index is wrong', async () => {
          const proof = leaves[0].recursiveProof([])
          const wrongIndex = 1

          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(proof, message, wrongIndex),
            `MerkleProofError(["${proof[0]}"], [${message.sequenceNumber}, ${message.sourceChainId}, ${message.destinationChainId}, "${message.sender}", ["${message.payload.receiver}", "${message.payload.data}", ["${message.payload.tokens[0]}"], [${message.payload.amounts[0]}], "${message.payload.executor}", "${message.payload.options}"]], ${wrongIndex})`,
          )
        })

        it('fails when the execution delay has not yet passed', async () => {
          const proof = leaves[0].recursiveProof([])
          await ramp
            .connect(roles.defaultAccount)
            .setExecutionDelaySeconds(60 * 60)
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(proof, message, 0),
            `ExecutionDelayError()`,
          )
        })
      })
      describe('validation fails', () => {
        it('fails if the receiver is the ramp', async () => {
          message.payload.receiver = ramp.address
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `InvalidReceiver("${message.payload.receiver}")`,
          )
        })
        it('fails if the receiver is the pool', async () => {
          message.payload.receiver = pool.address
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `InvalidReceiver("${message.payload.receiver}")`,
          )
        })
        it('fails if the receiver is the token', async () => {
          message.payload.receiver = token.address
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `InvalidReceiver("${message.payload.receiver}")`,
          )
        })
        it('fails if the receiver is not a contract', async () => {
          message.payload.receiver = await roles.oracleNode1.getAddress()
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `InvalidReceiver("${message.payload.receiver}")`,
          )
        })
        it('fails when the message executor is invalid', async () => {
          // Set the executor to a specific address, then executing with a different
          // one should revert.
          message.payload.executor = await roles.oracleNode1.getAddress()
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `InvalidExecutor(${message.sequenceNumber})`,
          )
        })
        it('fails when the message is already executed', async () => {
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await ramp
            .connect(roles.defaultAccount)
            .executeTransaction([], message, 0)
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `AlreadyExecuted(${message.sequenceNumber})`,
          )
        })
        it('should fail if sent from an unsupported source chain', async () => {
          message.sourceChainId = BigNumber.from(999)
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `InvalidSourceChain(${message.sourceChainId})`,
          )
        })
        it('should fail if the number of tokens sent is not 1', async () => {
          message.payload.tokens.push(await roles.oracleNode.getAddress())
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `UnsupportedNumberOfTokens()`,
          )
        })
        it('should fail if the number of amounts of tokens to send is not 1', async () => {
          message.payload.amounts.push(BigNumber.from(50000))
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `UnsupportedNumberOfTokens()`,
          )
        })
        it('should fail if sent using an unsupported token', async () => {
          message.payload.tokens[0] = await roles.oracleNode2.getAddress()
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `UnsupportedToken("${message.payload.tokens[0]}")`,
          )
        })
        it('should fail if sending more tokens than the tokenBucket allows', async () => {
          message.payload.amounts[0] = bucketCapactiy.add(1)
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `ExceedsTokenLimit(${bucketCapactiy}, ${message.payload.amounts[0]})`,
          )
        })
        it('should fail if the receiver does not support CrossChainMessageReceiverInterface', async () => {
          const nonReceiver = <MockERC20>(
            await deployContract(roles.defaultAccount, TokenArtifact, [
              'FAKE Token',
              'FAKE',
              await roles.defaultAccount.getAddress(),
              100,
            ])
          )
          message.payload.receiver = nonReceiver.address
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `ExecutionError(${message.sequenceNumber}, "0x")`,
          )
        })
        it('should fail if the contract is paused', async () => {
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await ramp.connect(roles.defaultAccount).pause()
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `Pausable: paused`,
          )
        })
        it('fails whenn the AFN signal is bad', async () => {
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await afn.voteBad()
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `BadAFNSignal()`,
          )
        })

        it('fails when the AFN signal is stale', async () => {
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await afn.setTimestamp(BigNumber.from(1))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `StaleAFNHeartbeat()`,
          )
        })
      })
    })

    describe('success', () => {
      let tx: ContractTransaction
      beforeEach(async () => {
        await ramp
          .connect(roles.defaultAccount)
          .report(
            encodeReport(
              constructReport(message, sequenceNumber, sequenceNumber),
            ),
          )
        tx = await ramp
          .connect(roles.defaultAccount)
          .executeTransaction([], message, 0)
      })
      it('should set s_executed to true', async () => {
        expect(await ramp.getExecuted(message.sequenceNumber)).to.be.true
      })
      it('should deliver the message to the receiver', async () => {
        messageDeepEqual(await receiver.s_message(), message)
      })
      it('should send the funds to the receiver contract', async () => {
        expect(await token.balanceOf(receiver.address)).to.equal(
          message.payload.amounts[0],
        )
      })
      it('should emit a CrossChainMessageExecuted event', async () => {
        expect(tx)
          .to.emit(ramp, 'CrossChainMessageExecuted')
          .withArgs(message.sequenceNumber)
      })
      it('should execute a message specifying an executor', async () => {
        message.payload.executor = await roles.oracleNode1.getAddress()
        message.sequenceNumber = message.sequenceNumber.add(1)
        await ramp
          .connect(roles.defaultAccount)
          .report(
            encodeReport(
              constructReport(
                message,
                sequenceNumber.add(1),
                sequenceNumber.add(1),
              ),
            ),
          )
        // Should not revert
        await expect(
          ramp.connect(roles.oracleNode1).executeTransaction([], message, 0),
        )
          .to.emit(ramp, 'CrossChainMessageExecuted')
          .withArgs(message.sequenceNumber)
      })
    })
  })

  describe('#pause', () => {
    it('owner can pause ramp', async () => {
      const account = roles.defaultAccount
      await expect(ramp.connect(account).pause())
        .to.emit(ramp, 'Paused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot pause pool', async function () {
      const account = roles.stranger
      await expect(ramp.connect(account).pause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#unpause', () => {
    beforeEach(async () => {
      await ramp.connect(roles.defaultAccount).pause()
    })

    it('owner can unpause ramp', async () => {
      const account = roles.defaultAccount
      await expect(ramp.connect(account).unpause())
        .to.emit(ramp, 'Unpaused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot unpause pool', async function () {
      const account = roles.stranger
      await expect(ramp.connect(account).unpause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#configureTokenBucket', () => {
    let newRate: BigNumber
    let newCapacity: BigNumber

    beforeEach(async () => {
      newRate = BigNumber.from(5)
      newCapacity = bucketCapactiy.add(10)
    })

    it('only callable by owner', async () => {
      await expect(
        ramp
          .connect(roles.stranger)
          .configureTokenBucket(newRate, newCapacity, true),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('sets the new max time without afn signal', async () => {
      const tx = await ramp
        .connect(roles.defaultAccount)
        .configureTokenBucket(newRate, newCapacity, true)
      const tokenBucketParams = await ramp.getTokenBucket()
      expect(tokenBucketParams.rate).to.equal(newRate)
      expect(tokenBucketParams.capacity).to.equal(newCapacity)
      await expect(tx)
        .to.emit(ramp, 'NewTokenBucketConstructed')
        .withArgs(newRate, newCapacity, true)
    })
  })

  describe('#setExecutionDelaySeconds', () => {
    it('can only be called by the owner', async () => {
      await evmRevert(
        ramp.connect(roles.stranger).setExecutionDelaySeconds(60),
        'Only callable by owner',
      )
    })

    it('sets the execution delay', async () => {
      const delaySeconds = 60
      const tx = await ramp
        .connect(roles.defaultAccount)
        .setExecutionDelaySeconds(delaySeconds)
      await expect(tx)
        .to.emit(ramp, 'ExecutionDelaySecondsSet')
        .withArgs(delaySeconds)
      const actualDelaySeconds = await ramp.getExecutionDelaySeconds()
      expect(actualDelaySeconds).to.equal(delaySeconds)
    })
  })

  describe('#setAFN', () => {
    let newAFN: MockAFN

    beforeEach(async () => {
      newAFN = <MockAFN>(
        await deployContract(roles.defaultAccount, MockAFNArtifact)
      )
    })

    it('only callable by owner', async () => {
      await expect(
        ramp.connect(roles.stranger).setAFN(newAFN.address),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('sets the new AFN', async () => {
      const tx = await ramp.connect(roles.defaultAccount).setAFN(newAFN.address)
      expect(await ramp.getAFN()).to.equal(newAFN.address)
      await expect(tx)
        .to.emit(ramp, 'AFNSet')
        .withArgs(afn.address, newAFN.address)
    })
  })

  describe('#setMaxSecondsWithoutAFNHeartbeat', () => {
    let newTime: BigNumber

    beforeEach(async () => {
      newTime = maxTimeBetweenAFNSignals.mul(2)
    })

    it('only callable by owner', async () => {
      await expect(
        ramp.connect(roles.stranger).setMaxSecondsWithoutAFNHeartbeat(newTime),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('sets the new max time without afn signal', async () => {
      const tx = await ramp
        .connect(roles.defaultAccount)
        .setMaxSecondsWithoutAFNHeartbeat(newTime)
      expect(await ramp.getMaxSecondsWithoutAFNHeartbeat()).to.equal(newTime)
      await expect(tx)
        .to.emit(ramp, 'AFNMaxHeartbeatTimeSet')
        .withArgs(maxTimeBetweenAFNSignals, newTime)
    })
  })

  describe('#typeAndVersion', () => {
    it('should return the correct type and version', async () => {
      const response = await ramp.typeAndVersion()
      await expect(response).to.equal('SingleTokenOffRamp 1.1.0')
    })
  })
})
                                                                                                                                                                                                                                                  contracts/test/v0.8/ccip/ramps/MessageExecutor.test.ts                                              000644  000765  000024  00000007217 14165346401 023713  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         import hre, { ethers } from 'hardhat'
import { expect } from 'chai'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { MockOffRamp, MessageExecutorHelper } from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { CCIPMessage, messageDeepEqual } from '../../../test-helpers/ccip'
import { BigNumber } from '@ethersproject/bignumber'
import { numToBytes32 } from '../../../test-helpers/helpers'

interface ExecutableMessage {
  proof: string[]
  message: CCIPMessage
  index: BigNumber
}

function encodeExecutableMessages(messages: ExecutableMessage[]): string {
  return ethers.utils.defaultAbiCoder.encode(
    [
      'tuple(bytes32[] proof, tuple(uint256 sequenceNumber, uint256 sourceChainId, uint256 destinationChainId, address sender, tuple(address receiver, bytes data, address[] tokens, uint256[] amounts, address executor, bytes options) payload) message, uint256 index)[] report',
    ],
    [messages],
  )
}

const { deployContract } = hre.waffle

let roles: Roles

let RampArtifact: Artifact
let ExecutorArtifact: Artifact

let ramp: MockOffRamp
let executor: MessageExecutorHelper

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('MessageExecutor', () => {
  beforeEach(async () => {
    RampArtifact = await hre.artifacts.readArtifact('MockOffRamp')
    ExecutorArtifact = await hre.artifacts.readArtifact('MessageExecutorHelper')

    ramp = <MockOffRamp>(
      await deployContract(roles.defaultAccount, RampArtifact, [])
    )
    executor = <MessageExecutorHelper>(
      await deployContract(roles.defaultAccount, ExecutorArtifact, [
        ramp.address,
      ])
    )
  })

  it('deploys correctly', async () => {
    expect(await executor.s_offRamp()).to.equal(ramp.address)
  })

  it('executes 2 messages in the same tx', async () => {
    const message1: CCIPMessage = {
      sequenceNumber: BigNumber.from(1),
      sourceChainId: BigNumber.from(1),
      destinationChainId: BigNumber.from(2),
      sender: await roles.oracleNode1.getAddress(),
      payload: {
        receiver: await roles.oracleNode2.getAddress(),
        data: numToBytes32(3),
        tokens: [],
        amounts: [],
        executor: hre.ethers.constants.AddressZero,
        options: numToBytes32(4),
      },
    }
    const message2: CCIPMessage = {
      sequenceNumber: BigNumber.from(2),
      sourceChainId: BigNumber.from(1),
      destinationChainId: BigNumber.from(2),
      sender: await roles.oracleNode3.getAddress(),
      payload: {
        receiver: await roles.oracleNode4.getAddress(),
        data: numToBytes32(7),
        tokens: [],
        amounts: [],
        executor: hre.ethers.constants.AddressZero,
        options: numToBytes32(8),
      },
    }
    const proof1 = [numToBytes32(9)]
    const proof2 = [numToBytes32(10)]
    const index1 = BigNumber.from(0)
    const index2 = BigNumber.from(1)

    const em1: ExecutableMessage = {
      proof: proof1,
      message: message1,
      index: index1,
    }
    const em2: ExecutableMessage = {
      proof: proof2,
      message: message2,
      index: index2,
    }
    const tx = await executor
      .connect(roles.defaultAccount)
      .report(encodeExecutableMessages([em1, em2]))
    const receipt = await tx.wait()
    const event1 = ramp.interface.parseLog(receipt.logs[0])
    const event2 = ramp.interface.parseLog(receipt.logs[1])

    expect(event1.args.proof).to.deep.equal(proof1)
    expect(event1.args.index).to.equal(index1)
    messageDeepEqual(event1.args.message, message1)

    expect(event2.args.proof).to.deep.equal(proof2)
    expect(event2.args.index).to.equal(index2)
    messageDeepEqual(event2.args.message, message2)
  })
})
                                                                                                                                                                                                                                                                                                                                                                                 contracts/test/v0.8/ccip/ramps/SingleTokenContractEndToEnd.test.ts                                  000644  000765  000024  00000016033 14165346401 026105  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         import hre, { ethers } from 'hardhat'
import { stringToBytes } from '../../../test-helpers/helpers'
import { expect } from 'chai'
import { BigNumber, Contract, ContractReceipt } from 'ethers'
import { Roles, getUsers } from '../../../test-helpers/setup'
import {
  MockERC20,
  LockUnlockPool,
  SimpleMessageReceiver,
  SingleTokenOnRamp,
  MockAFN,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import {
  CCIPMessage,
  CCIPMessagePayload,
  encodeReport,
  hashMessage,
  messageDeepEqual,
  RelayReport,
} from '../../../test-helpers/ccip'

const { deployContract } = hre.waffle

let roles: Roles

let chain1AFN: MockAFN
let chain1OnRamp: SingleTokenOnRamp
let chain1Token: MockERC20
let chain1Pool: LockUnlockPool
const chain1ID: number = 1

// This has to be ethers.Contract because of an issue with
// `address.call(abi.encodeWithSelector(...))` using typechain artifacts.
let chain2OffRamp: Contract
let chain2AFN: MockAFN
let chain2Token: MockERC20
let chain2Receiver: SimpleMessageReceiver
let chain2Pool: LockUnlockPool
const chain2ID: number = 2

const sendAmount = BigNumber.from('1000000000000000000')
const maxTimeBetweenAFNSignals = sendAmount
const executionDelay = 0

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('Single Token Contract End to End', () => {
  beforeEach(async () => {
    const adminAddress = await roles.defaultAccount.getAddress()

    const MockAFNArtifact: Artifact = await hre.artifacts.readArtifact(
      'MockAFN',
    )
    const TokenArtifact: Artifact = await hre.artifacts.readArtifact(
      'MockERC20',
    )
    const PoolArtifact: Artifact = await hre.artifacts.readArtifact(
      'LockUnlockPool',
    )
    const offRampFactory = await ethers.getContractFactory(
      'SingleTokenOffRampHelper',
    )
    const OnRampArtifact: Artifact = await hre.artifacts.readArtifact(
      'SingleTokenOnRamp',
    )
    const SimpleMessageReceiverArtifact: Artifact =
      await hre.artifacts.readArtifact('SimpleMessageReceiver')

    // Deploy chain2 contracts
    chain2Token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'Chain 2 LINK Token',
        'LINK',
        adminAddress,
        BigNumber.from('100000000000000000000'),
      ])
    )
    chain2Pool = <LockUnlockPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [
        chain2Token.address,
      ])
    )
    chain2AFN = <MockAFN>(
      await deployContract(roles.defaultAccount, MockAFNArtifact)
    )
    chain2OffRamp = await offRampFactory.connect(roles.defaultAccount).deploy(
      chain1ID,
      chain2ID,
      chain2Token.address,
      chain2Pool.address,
      sendAmount, // bucketRate
      sendAmount, // bucketCapacity
      chain2AFN.address,
      maxTimeBetweenAFNSignals,
      executionDelay,
    )
    await chain2Pool
      .connect(roles.defaultAccount)
      .setOffRamp(chain2OffRamp.address, true)
    await chain2Token
      .connect(roles.defaultAccount)
      .approve(chain2Pool.address, sendAmount)
    await chain2Pool
      .connect(roles.defaultAccount)
      .lockOrBurn(adminAddress, sendAmount)
    chain2Receiver = <SimpleMessageReceiver>(
      await deployContract(roles.defaultAccount, SimpleMessageReceiverArtifact)
    )

    // Deploy chain1 contracts
    chain1Token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'Chain 1 LINK Token',
        'LINK',
        adminAddress,
        BigNumber.from('100000000000000000000'),
      ])
    )
    chain1Pool = <LockUnlockPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [
        chain1Token.address,
      ])
    )
    chain1AFN = <MockAFN>(
      await deployContract(roles.defaultAccount, MockAFNArtifact)
    )
    chain1OnRamp = <SingleTokenOnRamp>(
      await deployContract(roles.defaultAccount, OnRampArtifact, [
        chain1ID,
        chain1Token.address,
        chain1Pool.address,
        chain2ID,
        chain2Token.address,
        [roles.defaultAccount.getAddress()],
        true,
        sendAmount, // bucketRate
        sendAmount, // bucketCapacity
        chain1AFN.address,
        maxTimeBetweenAFNSignals,
      ])
    )
    await chain1Pool
      .connect(roles.defaultAccount)
      .setOnRamp(chain1OnRamp.address, true)
  })

  it('should send a message and tokens from chain1 to chain2', async () => {
    const messagedata = stringToBytes('Message')
    const options = hre.ethers.constants.HashZero
    const payload: CCIPMessagePayload = {
      receiver: chain2Receiver.address,
      data: messagedata,
      tokens: [chain1Token.address],
      amounts: [sendAmount],
      executor: hre.ethers.constants.AddressZero,
      options: options,
    }

    const initialChain1PoolBalance = await chain1Token.balanceOf(
      chain1Pool.address,
    )
    const initialChain2ReceiverBalance = await chain2Token.balanceOf(
      chain2Receiver.address,
    )
    // approve tokens and send message
    const chain1PoolAddress = await chain1OnRamp.POOL()
    await chain1Token.approve(chain1PoolAddress, sendAmount)
    let tx = await chain1OnRamp.requestCrossChainSend(payload)

    // Check tokens are locked
    await expect(await chain1Token.balanceOf(chain1Pool.address)).to.equal(
      initialChain1PoolBalance.add(sendAmount),
    )

    // DON picks up event and reads
    let receipt: ContractReceipt = await tx.wait()
    let eventArgs = receipt.events?.[3]?.args?.[0]
    const sequenceNumber = eventArgs?.sequenceNumber
    const donPayload: CCIPMessagePayload = {
      receiver: eventArgs?.payload.receiver,
      data: eventArgs?.payload.data,
      tokens: eventArgs?.payload.tokens,
      amounts: eventArgs?.payload.amounts,
      executor: eventArgs?.payload.executor,
      options: eventArgs?.payload.options,
    }
    const donMessage: CCIPMessage = {
      sequenceNumber: sequenceNumber,
      sourceChainId: BigNumber.from(chain1ID),
      destinationChainId: BigNumber.from(chain2ID),
      sender: eventArgs?.sender,
      payload: donPayload,
    }

    // DON encodes, reports and executes the message
    let report: RelayReport = {
      merkleRoot: hashMessage(donMessage),
      minSequenceNumber: sequenceNumber,
      maxSequenceNumber: sequenceNumber,
    }
    await chain2OffRamp
      .connect(roles.defaultAccount)
      .report(encodeReport(report))
    tx = await chain2OffRamp
      .connect(roles.defaultAccount)
      .executeTransaction([], donMessage, 0)
    receipt = await tx.wait()

    // Check that events are emitted and receiver receives the message
    await expect(tx)
      .to.emit(chain2OffRamp, 'CrossChainMessageExecuted')
      .withArgs(donMessage.sequenceNumber)

    await expect(tx).to.emit(chain2Receiver, 'MessageReceived')
    const receivedPayload = await chain2Receiver.s_message()
    messageDeepEqual(receivedPayload, donMessage)

    // Check balance of contract
    const afterChain2ReceiverBalance = await chain2Token.balanceOf(
      chain2Receiver.address,
    )
    expect(afterChain2ReceiverBalance).to.equal(
      initialChain2ReceiverBalance.add(sendAmount),
    )
  })
})
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     contracts/test/v0.8/ccip/scripts/deployments.ts                                                     000644  000765  000024  00000021070 14165357743 022546  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         import { Contract, ContractFactory } from '@ethersproject/contracts'
import { NonceManager } from '@ethersproject/experimental'
import { HardhatRuntimeEnvironment } from 'hardhat/types'

interface PoolInformation {
  chainId: number
  link: string
  pool: Contract
  wallet: NonceManager
}

interface Lane {
  source: {
    eoaSender: Contract
    onRamp: Contract
    poolInfo: PoolInformation
  }
  destination: {
    eoaReceiver: Contract
    offRamp: Contract
    poolInfo: PoolInformation
  }
}

export async function deployPools(
  _: any,
  hre: HardhatRuntimeEnvironment,
  wallets: any,
  envVars: any,
) {
  const poolFactory: ContractFactory = await hre.ethers.getContractFactory(
    'LockUnlockPool',
  )
  console.log('Deploying to Kovan...')
  const kovanPool = await poolFactory
    .connect(wallets.kovan)
    .deploy(envVars.KOVAN_LINK)
  const kovanPoolVerifyCmd = `yarn hardhat verify --network kovan ${kovanPool.address} ${envVars.KOVAN_LINK}`
  console.log('Deploying to Rinkeby...')
  const rinkebyPool = await poolFactory
    .connect(wallets.rinkeby)
    .deploy(envVars.RINKEBY_LINK)
  const rinkebyPoolVerifyCmd = `yarn hardhat verify --network rinkeby ${rinkebyPool.address} ${envVars.RINKEBY_LINK}`
  console.log('Deploying to Georli...')
  const goerliPool = await poolFactory
    .connect(wallets.goerli)
    .deploy(envVars.GOERLI_LINK)
  const goerliPoolVerifyCmd = `yarn hardhat verify --network goerli ${goerliPool.address} ${envVars.GOERLI_LINK}`

  const output = `
==== DEPLOYED POOLS ====
Kovan:    ${kovanPool.address}
Rinkeby:  ${rinkebyPool.address}
Goerli:   ${goerliPool.address}
==== -------------- ====
To verify the pools, run the following commands:

${kovanPoolVerifyCmd}

${rinkebyPoolVerifyCmd}

${goerliPoolVerifyCmd}`

  console.log(output)
}

/**
 * Create a lane between two networks
 * @param hre
 * @param wallets
 * @param envVars
 */
async function createLane(
  hre: HardhatRuntimeEnvironment,
  network1: PoolInformation,
  network2: PoolInformation,
): Promise<Lane> {
  const offRampFactory: ContractFactory = await hre.ethers.getContractFactory(
    'SingleTokenOffRamp',
  )
  const eoaReceiverFactory: ContractFactory =
    await hre.ethers.getContractFactory('EOASingleTokenReceiver')
  const onRampFactory: ContractFactory = await hre.ethers.getContractFactory(
    'SingleTokenOnRamp',
  )
  const eoaSenderFactory: ContractFactory = await hre.ethers.getContractFactory(
    'EOASingleTokenSender',
  )

  // Set up network1 -> network2 lane end to end
  // network2 side receiving contracts first
  console.log('deploy offramp')
  const oneToTwoOffRamp = await offRampFactory
    .connect(network2.wallet)
    .deploy(
      network1.chainId!,
      network2.chainId!,
      network2.link,
      network2.pool.address,
    )
  const setOffRampGas = await network2.pool
    .connect(network2.wallet)
    .estimateGas.setOnRamp(oneToTwoOffRamp.address, true)
  console.log('set off ramp with gas', setOffRampGas.toString())
  await (
    await network2.pool
      .connect(network2.wallet)
      .setOffRamp(oneToTwoOffRamp.address, true, {
        gasLimit: setOffRampGas.add(10000),
      })
  ).wait()
  console.log('deploy receiver')
  const oneToTwoEOAReceiver = await eoaReceiverFactory
    .connect(network2.wallet)
    .deploy(oneToTwoOffRamp.address)
  // network 1 side sending contracts next
  console.log('deploy onramp')
  const oneToTwoOnRamp = await onRampFactory
    .connect(network1.wallet)
    .deploy(
      network1.chainId,
      network1.link,
      network1.pool.address,
      network2.chainId,
      network2.link,
    )
  const setOnRampGas = await network1.pool
    .connect(network1.wallet)
    .estimateGas.setOnRamp(oneToTwoOnRamp.address, true)
  console.log('set on ramp with gas', setOnRampGas.toString())
  await (
    await network1.pool
      .connect(network1.wallet)
      .setOnRamp(oneToTwoOnRamp.address, true, {
        gasLimit: setOnRampGas.add(10000),
      })
  ).wait()
  console.log('deploy sender')
  const oneToTwoEOASender = await eoaSenderFactory
    .connect(network1.wallet)
    .deploy(oneToTwoOnRamp.address, oneToTwoEOAReceiver.address)

  const lane: Lane = {
    source: {
      eoaSender: oneToTwoEOASender,
      onRamp: oneToTwoOnRamp,
      poolInfo: network1,
    },
    destination: {
      eoaReceiver: oneToTwoEOAReceiver,
      offRamp: oneToTwoOffRamp,
      poolInfo: network2,
    },
  }
  return lane
}

/**
 * Deploy ramps to create lanes between networks
 * @param args
 * @param hre
 * @param wallets
 * @param envVars
 */
export async function deployRamps(
  args: any,
  hre: HardhatRuntimeEnvironment,
  wallets: any,
  envVars: any,
) {
  const poolFactory: ContractFactory = await hre.ethers.getContractFactory(
    'LockUnlockPool',
  )

  const kovan: PoolInformation = {
    chainId: envVars.KOVAN_ID,
    link: envVars.KOVAN_LINK,
    pool: await poolFactory.connect(wallets.kovan).attach(args.kp),
    wallet: wallets.kovan,
  }
  const rinkeby: PoolInformation = {
    chainId: envVars.RINKEBY_ID,
    link: envVars.RINKEBY_LINK,
    pool: await poolFactory.connect(wallets.rinkeby).attach(args.rp),
    wallet: wallets.rinkeby,
  }
  const goerli: PoolInformation = {
    chainId: envVars.GOERLI_ID,
    link: envVars.GOERLI_LINK,
    pool: await poolFactory.connect(wallets.goerli).attach(args.gp),
    wallet: wallets.goerli,
  }

  // kovan->rinkeby
  console.log('deploying kovan to rinkeby lane')
  const kovanToRinkeby: Lane = await createLane(hre, kovan, rinkeby)
  // rinkeby->kovan
  console.log('deploying rinkeby to kovan lane')
  const rinkebyToKovan: Lane = await createLane(hre, rinkeby, kovan)

  // kovan->goerli
  console.log('deploying kovan to goerli lane')
  const kovanToGoerli: Lane = await createLane(hre, kovan, goerli)
  // goerli->kovan
  console.log('deploying goerli to kovan lane')
  const goerliToKovan: Lane = await createLane(hre, goerli, kovan)

  // rinkeby->goerli
  console.log('deploying rinkeby to goerli lane')
  const rinkebyToGoerli: Lane = await createLane(hre, rinkeby, goerli)
  // goerli->rinkeby
  console.log('deploying goerli to rinkeby lane')
  const goerliToRinkeby: Lane = await createLane(hre, goerli, rinkeby)

  console.log(
    `
==== DEPLOYED LANES ====

--- KOVAN <-> RINKEBY ---
- KOVAN -
EOASender:      ${kovanToRinkeby.source.eoaSender.address}
OnRamp:         ${kovanToRinkeby.source.onRamp.address}
OffRamp:        ${rinkebyToKovan.destination.offRamp.address}
EOAReceiver:    ${rinkebyToKovan.destination.eoaReceiver.address}

- Rinkeby -
EOASender:      ${rinkebyToKovan.source.eoaSender.address}
OnRamp:         ${rinkebyToKovan.source.onRamp.address}
OffRamp:        ${kovanToRinkeby.destination.offRamp.address}
EOAReceiver:    ${kovanToRinkeby.destination.eoaReceiver.address}

--- KOVAN <-> GOERLI ---
- KOVAN -
EOASender:      ${kovanToGoerli.source.eoaSender.address}
OnRamp:         ${kovanToGoerli.source.onRamp.address}
OffRamp:        ${goerliToKovan.destination.offRamp.address}
EOAReceiver:    ${goerliToKovan.destination.eoaReceiver.address}

- Goerli -
EOASender:      ${goerliToKovan.source.eoaSender.address}
OnRamp:         ${goerliToKovan.source.onRamp.address}
OffRamp:        ${kovanToGoerli.destination.offRamp.address}
EOAReceiver:    ${kovanToGoerli.destination.eoaReceiver.address}

--- RINKEBY <-> GOERLI ---
- Rinkeby -
EOASender:      ${rinkebyToGoerli.source.eoaSender.address}
OnRamp:         ${rinkebyToGoerli.source.onRamp.address}
OffRamp:        ${goerliToRinkeby.destination.offRamp.address}
EOAReceiver:    ${goerliToRinkeby.destination.eoaReceiver.address}

- Goerli -
EOASender:      ${goerliToRinkeby.source.eoaSender.address}
OnRamp:         ${goerliToRinkeby.source.onRamp.address}
OffRamp:        ${rinkebyToGoerli.destination.offRamp.address}
EOAReceiver:    ${rinkebyToGoerli.destination.eoaReceiver.address}
`,
  )
}

export async function transferOwnership(
  args: any,
  hre: HardhatRuntimeEnvironment,
  wallets: any,
) {
  const contractFactory = await hre.ethers.getContractFactory(
    'src/v0.8/ConfirmedOwner.sol:ConfirmedOwner',
  )
  let wallet
  if (args.chain == 'kovan') {
    wallet = wallets.kovan
  } else if (args.chain == 'rinkeby') {
    wallet = wallets.rinkeby
  } else if (args.chain == 'goerli') {
    wallet = wallets.goerli
  } else {
    throw new Error("Chain config doesn't exist")
  }

  const contract = await contractFactory.connect(wallet).attach(args.contract)
  const gasLimit = await contract
    .connect(wallet)
    .estimateGas.transferOwnership(args.to)
  console.log('Transfer ownership with gas', gasLimit.toString())
  await (
    await contract
      .connect(wallet)
      .transferOwnership(args.to, { gasLimit: gasLimit.add(10000) })
  ).wait()
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                        contracts/test/v0.8/ccip/applications/EOASingleTokenSender.test.ts                                  000644  000765  000024  00000010601 14165346401 026053  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         import hre, { ethers } from 'hardhat'
import { expect } from 'chai'
import { Roles, getUsers } from '../../../test-helpers/setup'
import {
  MockOnRamp,
  EOASingleTokenSender,
  MockERC20,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { BigNumber } from '@ethersproject/bignumber'
import { evmRevert } from '../../../test-helpers/matchers'

const { deployContract } = hre.waffle

let roles: Roles

let SenderArtifact: Artifact
let RampArtifact: Artifact
let TokenArtifact: Artifact

let sourceChainId: BigNumber
let token: MockERC20
let destinationToken: string
let destinationChainId: BigNumber

let ramp: MockOnRamp
let senderContract: EOASingleTokenSender
let destinationContract: string

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
  destinationContract = await users.contracts.contract8.getAddress()
})

describe('EOASingleTokenSender', () => {
  beforeEach(async () => {
    sourceChainId = BigNumber.from(1)
    destinationToken = await roles.oracleNode2.getAddress()
    destinationChainId = BigNumber.from(2)

    SenderArtifact = await hre.artifacts.readArtifact('EOASingleTokenSender')
    RampArtifact = await hre.artifacts.readArtifact('MockOnRamp')
    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')

    token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'LINK Token',
        'LINK',
        await roles.defaultAccount.getAddress(),
        BigNumber.from('10000000000000000000'),
      ])
    )

    ramp = <MockOnRamp>(
      await deployContract(roles.defaultAccount, RampArtifact, [
        sourceChainId,
        token.address,
        destinationToken,
        await roles.oracleNode3.getAddress(),
        destinationChainId,
      ])
    )

    senderContract = <EOASingleTokenSender>(
      await deployContract(roles.defaultAccount, SenderArtifact, [
        ramp.address,
        destinationContract,
      ])
    )
  })

  describe('#constructor', () => {
    it('should set the onRamp', async () => {
      const onRamp = await senderContract.ON_RAMP()
      expect(onRamp).to.equal(ramp.address)
    })

    it('#should set the destination contract', async () => {
      const destContract = await senderContract.DESTINATION_CONTRACT()
      expect(destContract).to.equal(destinationContract)
    })
  })

  describe('#sendMessage', () => {
    let senderAddress: string
    let destinationAddress: string
    let data: string
    let amount: BigNumber
    let options: string

    beforeEach(async () => {
      senderAddress = await roles.defaultAccount.getAddress()
      destinationAddress = senderAddress
      data = ethers.utils.defaultAbiCoder.encode(
        ['address', 'address'],
        [senderAddress, destinationAddress],
      )
      amount = BigNumber.from('1000000000000000000')
      options = '0x'
    })

    it('should send a request to the onRamp', async () => {
      const expectedResponse = [
        destinationContract,
        data,
        [token.address],
        [amount],
        options,
      ]

      await token.approve(senderContract.address, amount)
      await senderContract.sendTokens(
        destinationAddress,
        amount,
        ethers.constants.AddressZero,
      )
      const response = await ramp.getMessagePayload()
      for (let i = 0; i < response.length; i++) {
        const actual = response[i].toString()
        const expected = expectedResponse[i].toString()
        expect(actual).to.deep.equal(expected)
      }
    })

    it('should fail if the destination address is zero address', async () => {
      await evmRevert(
        senderContract.sendTokens(
          ethers.constants.AddressZero,
          amount,
          ethers.constants.AddressZero,
        ),
        `InvalidDestinationAddress("${ethers.constants.AddressZero}")`,
      )
    })
  })

  describe('#rampDetails', () => {
    it('returns the correct destination chain ID', async () => {
      const response = await senderContract.rampDetails()
      expect(response.destinationChainId).to.equal(destinationChainId)
      expect(response.token).to.equal(token.address)
      expect(response.destinationChainToken).to.equal(destinationToken)
    })
  })

  describe('#typeAndVersion', () => {
    it('should return the correct type and version', async () => {
      expect(await senderContract.typeAndVersion()).to.equal(
        'EOASingleTokenSender 1.0.0',
      )
    })
  })
})
                                                                                                                               contracts/test/v0.8/ccip/applications/EOASingleTokenReceiver.test.ts                                000644  000765  000024  00000007643 14165346401 026413  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         import hre, { ethers } from 'hardhat'
import { expect } from 'chai'
import { Roles, getUsers } from '../../../test-helpers/setup'
import {
  MockERC20,
  MockOffRamp,
  EOASingleTokenReceiver,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { evmRevert } from '../../../test-helpers/matchers'
import { CCIPMessage } from '../../../test-helpers/ccip'
import { BigNumber } from '@ethersproject/bignumber'

const { deployContract } = hre.waffle

let roles: Roles

let ReceiverArtifact: Artifact
let RampArtifact: Artifact
let TokenArtifact: Artifact

let ramp: MockOffRamp
let receiverContract: EOASingleTokenReceiver
let token: MockERC20

let balance: BigNumber

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('EOASingleTokenReceiver', () => {
  beforeEach(async () => {
    balance = BigNumber.from('12000000000000000000')
    ReceiverArtifact = await hre.artifacts.readArtifact(
      'EOASingleTokenReceiver',
    )
    RampArtifact = await hre.artifacts.readArtifact('MockOffRamp')
    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')

    ramp = <MockOffRamp>(
      await deployContract(roles.defaultAccount, RampArtifact, [])
    )

    receiverContract = <EOASingleTokenReceiver>(
      await deployContract(roles.defaultAccount, ReceiverArtifact, [
        ramp.address,
      ])
    )
    token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'LINK Token',
        'LINK',
        receiverContract.address,
        balance,
      ])
    )
    await ramp.setToken(token.address)
  })

  describe('#constructor', () => {
    it('sets the off ramp', async () => {
      const response = await receiverContract.OFF_RAMP()
      expect(response).to.equal(ramp.address)
    })
  })

  describe('#receiveMessage', () => {
    let accountAddr: string

    it('fails if the sender is not the off ramp', async () => {
      const message: CCIPMessage = {
        sequenceNumber: BigNumber.from(1),
        sourceChainId: BigNumber.from(1),
        destinationChainId: BigNumber.from(2),
        sender: ethers.constants.AddressZero,
        payload: {
          receiver: ethers.constants.AddressZero,
          data: ethers.constants.HashZero,
          tokens: [],
          amounts: [],
          executor: ethers.constants.AddressZero,
          options: ethers.constants.HashZero,
        },
      }
      accountAddr = await roles.defaultAccount.getAddress()
      await evmRevert(
        receiverContract.connect(roles.defaultAccount).receiveMessage(message),
        `InvalidDeliverer("${accountAddr}")`,
      )
    })
    describe('success', () => {
      let data: string
      let sequenceNumber: BigNumber
      let amount: BigNumber

      beforeEach(async () => {
        accountAddr = await roles.defaultAccount.getAddress()
        data = ethers.utils.defaultAbiCoder.encode(
          ['address', 'address'],
          [accountAddr, accountAddr],
        )
        sequenceNumber = BigNumber.from(1)
        amount = balance
        const message: CCIPMessage = {
          sequenceNumber,
          sourceChainId: BigNumber.from(5),
          destinationChainId: BigNumber.from(2),
          sender: receiverContract.address,
          payload: {
            receiver: receiverContract.address,
            data,
            tokens: [token.address],
            amounts: [amount],
            executor: ethers.constants.AddressZero,
            options: ethers.constants.HashZero,
          },
        }
        await ramp.deliverMessageTo(receiverContract.address, message)
      })

      it('forwards the tokens', async () => {
        expect(await token.balanceOf(accountAddr)).to.equal(amount)
      })
    })
  })

  describe('#typeAndVersion', () => {
    it('returns the type and version', async () => {
      const response = await receiverContract.typeAndVersion()
      expect(response).to.equal('EOASingleTokenReceiver 1.1.0')
    })
  })
})
                                                                                             contracts/test/v0.8/ccip/applications/EOASingleTokenEndToEnd.test.ts                                000644  000765  000024  00000015571 14165346401 026306  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         import hre, { ethers } from 'hardhat'
import { BigNumber, Contract } from 'ethers'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { expect } from 'chai'
import {
  MockERC20,
  LockUnlockPool,
  EOASingleTokenReceiver,
  SingleTokenOnRamp,
  EOASingleTokenSender,
  MockAFN,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import {
  CCIPMessage,
  encodeReport,
  hashMessage,
  RelayReport,
} from '../../../test-helpers/ccip'

const { deployContract } = hre.waffle

let roles: Roles

let chain1OnApp: EOASingleTokenSender
let chain1AFN: MockAFN
let chain1OnRamp: SingleTokenOnRamp
let chain1Token: MockERC20
let chain1Pool: LockUnlockPool
const chain1ID: number = 1

// This has to be ethers.Contract because of an issue with
// `address.call(abi.encodeWithSelector(...))` and try-catch using typechain artifacts.
let chain2OffRamp: Contract
let chain2AFN: MockAFN
let chain2OffApp: EOASingleTokenReceiver
let chain2Token: MockERC20
let chain2Pool: LockUnlockPool
const chain2ID: number = 2

const sendAmount = BigNumber.from('1000000000000000000')
const maxTimeBetweenAFNSignals = sendAmount
const executionDelay = 0

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('Single Token EOA End to End', () => {
  beforeEach(async () => {
    const adminAddress = await roles.defaultAccount.getAddress()

    const MockAFNArtifact: Artifact = await hre.artifacts.readArtifact(
      'MockAFN',
    )
    const TokenArtifact: Artifact = await hre.artifacts.readArtifact(
      'MockERC20',
    )
    const PoolArtifact: Artifact = await hre.artifacts.readArtifact(
      'LockUnlockPool',
    )
    const offRampFactory = await ethers.getContractFactory(
      'SingleTokenOffRampHelper',
    )
    const OnRampSenderArtifact: Artifact = await hre.artifacts.readArtifact(
      'EOASingleTokenSender',
    )
    const OnRampArtifact: Artifact = await hre.artifacts.readArtifact(
      'SingleTokenOnRamp',
    )
    const OffRampReceiverArtifact: Artifact = await hre.artifacts.readArtifact(
      'EOASingleTokenReceiver',
    )

    // Deploy chain2 contracts
    chain2Token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'Chain 2 LINK Token',
        'LINK',
        adminAddress,
        BigNumber.from('100000000000000000000'),
      ])
    )
    chain2Pool = <LockUnlockPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [
        chain2Token.address,
      ])
    )
    chain2AFN = <MockAFN>(
      await deployContract(roles.defaultAccount, MockAFNArtifact)
    )
    chain2OffRamp = await offRampFactory.connect(roles.defaultAccount).deploy(
      chain1ID,
      chain2ID,
      chain2Token.address,
      chain2Pool.address,
      sendAmount, //bucketRate
      sendAmount, //bucketCapacity
      chain2AFN.address,
      maxTimeBetweenAFNSignals,
      executionDelay,
    )
    await chain2Pool
      .connect(roles.defaultAccount)
      .setOffRamp(chain2OffRamp.address, true)
    await chain2Token
      .connect(roles.defaultAccount)
      .approve(chain2Pool.address, sendAmount)
    await chain2Pool
      .connect(roles.defaultAccount)
      .lockOrBurn(adminAddress, sendAmount)
    chain2OffApp = <EOASingleTokenReceiver>(
      await deployContract(roles.defaultAccount, OffRampReceiverArtifact, [
        chain2OffRamp.address,
      ])
    )

    // Deploy chain1 contracts
    chain1Token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'Chain 1 LINK Token',
        'LINK',
        adminAddress,
        BigNumber.from('100000000000000000000'),
      ])
    )
    chain1Pool = <LockUnlockPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [
        chain1Token.address,
      ])
    )
    chain1AFN = <MockAFN>(
      await deployContract(roles.defaultAccount, MockAFNArtifact)
    )
    chain1OnRamp = <SingleTokenOnRamp>(
      await deployContract(roles.defaultAccount, OnRampArtifact, [
        chain1ID,
        chain1Token.address,
        chain1Pool.address,
        chain2ID,
        chain2Token.address,
        [],
        true,
        sendAmount, // bucketRate
        sendAmount, // bucketCapacity
        chain1AFN.address,
        maxTimeBetweenAFNSignals,
      ])
    )
    await chain1Pool
      .connect(roles.defaultAccount)
      .setOnRamp(chain1OnRamp.address, true)
    chain1OnApp = <EOASingleTokenSender>(
      await deployContract(roles.defaultAccount, OnRampSenderArtifact, [
        chain1OnRamp.address,
        chain2OffApp.address,
      ])
    )
    await chain1OnRamp.setAllowlist([chain1OnApp.address])
    await chain1Token.transfer(await roles.stranger.getAddress(), sendAmount)
  })

  it('should send tokens from chain1 to chain2 EOAs', async () => {
    // Initial balances
    const chain1StrangerInitialBalance = await chain1Token.balanceOf(
      await roles.stranger.getAddress(),
    )
    const chain2StrangerInitialBalance = await chain2Token.balanceOf(
      await roles.stranger.getAddress(),
    )

    // approve tokens and send message
    await chain1Token
      .connect(roles.stranger)
      .approve(chain1OnApp.address, sendAmount)
    let tx = await chain1OnApp
      .connect(roles.stranger)
      .sendTokens(
        await roles.stranger.getAddress(),
        sendAmount,
        ethers.constants.AddressZero,
      )

    // Parse log
    let receipt = await tx.wait()
    const log = receipt.logs[6]
    const decodedLog = chain1OnRamp.interface.parseLog(log)
    const logArgs = decodedLog.args[0]

    // Send messge to chain2
    const message: CCIPMessage = {
      sequenceNumber: logArgs.sequenceNumber,
      sourceChainId: BigNumber.from(chain1ID),
      destinationChainId: BigNumber.from(chain2ID),
      sender: logArgs.sender,
      payload: {
        receiver: logArgs.payload.receiver,
        data: logArgs.payload.data,
        tokens: logArgs.payload.tokens,
        amounts: logArgs.payload.amounts,
        executor: logArgs.payload.executor,
        options: logArgs.payload.options,
      },
    }
    // DON encodes, reports and executes the message
    let report: RelayReport = {
      merkleRoot: hashMessage(message),
      minSequenceNumber: logArgs.sequenceNumber,
      maxSequenceNumber: logArgs.sequenceNumber,
    }
    await chain2OffRamp
      .connect(roles.defaultAccount)
      .report(encodeReport(report))
    tx = await chain2OffRamp
      .connect(roles.defaultAccount)
      .executeTransaction([], message, 0)
    receipt = await tx.wait()

    const chain1StrangerBalanceAfter = await chain1Token.balanceOf(
      await roles.stranger.getAddress(),
    )
    const chain2StrangerBalanceAfter = await chain2Token.balanceOf(
      await roles.stranger.getAddress(),
    )

    expect(
      chain1StrangerInitialBalance.sub(chain1StrangerBalanceAfter),
    ).to.equal(sendAmount)
    expect(
      chain2StrangerBalanceAfter.sub(chain2StrangerInitialBalance),
    ).to.equal(sendAmount)
  })
})
                                                                                                                                       contracts/test/test-helpers/ccip.ts                                                                 000644  000765  000024  00000014413 14165346401 020430  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         import { BigNumber, BigNumberish, BytesLike } from 'ethers'
import { ethers } from 'hardhat'
import { expect } from 'chai'
export interface RelayReport {
  merkleRoot: string
  minSequenceNumber: BigNumber
  maxSequenceNumber: BigNumber
}

export interface CCIPMessage {
  sequenceNumber: BigNumber
  sourceChainId: BigNumber
  destinationChainId: BigNumber
  sender: string
  payload: CCIPMessagePayload
}

export interface CCIPMessagePayload {
  receiver: string
  data: BytesLike
  tokens: string[]
  amounts: BigNumberish[]
  executor: string
  options: BytesLike
}

export class MerkleTree {
  public parent?: MerkleTree

  /**
   * Left subtree
   */
  public left?: MerkleTree

  /**
   * Right subtree
   */
  public right?: MerkleTree

  /**
   * Hash that is either provide or populated
   */
  public hash?: string

  constructor(hash?: string) {
    this.hash = hash
  }

  public getSiblingHash(hash?: string): string {
    if (hash == this.left?.hash) {
      return this.right?.hash!
    } else if (hash == this.right?.hash) {
      return this.left?.hash!
    } else {
      throw new Error('Hash not found')
    }
  }

  public recursiveProof(proof: string[]): string[] {
    if (this.parent != undefined) {
      proof.push(this.parent?.getSiblingHash(this.hash)!)
      this.parent.recursiveProof(proof)
    }
    return proof
  }

  /**
   * Computes the hash based on the children. If no right child
   * exists, reuse the left child's value
   */
  public computeHash(): string {
    const leftHash = this.left!.hash
    const rightHash = this.right ? this.right.hash : leftHash
    // Add the internal node domain separator.
    return ethers.utils.solidityKeccak256(
      ['bytes', 'bytes32', 'bytes32'],
      ['0x01', leftHash, rightHash],
    )
  }
}

export function generateMerkleTreeFromHashes(hashes: string[]): any {
  // Convert the initial hashes into leaf nodes. We will use these
  // leaf nodes to construuct the Merkle tree from the bottom up by
  // successively combining pairing nodes at each level to construct
  // the parent
  let nodes = hashes.map((p) => new MerkleTree(p))
  let leaves: MerkleTree[] = []

  // Loop until we reach a single node, which will be our Merkle root
  while (nodes.length > 1) {
    const parents = []

    // Successively pair up nodes at each level
    for (let i = 0; i < nodes.length; i += 2) {
      // Create the parent node, which we will add a left, try add
      // a right, then calculate the hash for the node
      const parent = new MerkleTree()
      parents.push(parent)

      // Assign the left, which will always be there
      parent.left = nodes[i]
      nodes[i].parent = parent

      // Assign the right, which won't always be there. However,
      // in JavaScript, an array overflow simply returns undefined
      // which in this context, is the same as a null pointer.
      parent.right = nodes[i + 1]
      nodes[i + 1].parent = parent

      // Finally compute the hash, which will be based on the
      // number of children.
      parent.hash = parent.computeHash()

      // Add to the leaves if we're still on the bottom level
      if (leaves.length < hashes.length) {
        leaves.push(nodes[i], nodes[i + 1])
      }
    }

    // Once all pairs have been made, the parents now become the
    // children and we start all over again
    nodes = parents
  }

  // Return the single node as our root
  return {
    root: nodes[0],
    leaves: leaves,
  }
}

export function encodeReport(report: RelayReport) {
  return ethers.utils.defaultAbiCoder.encode(
    [
      'tuple(bytes32 merkleRoot, uint256 minSequenceNumber, uint256 maxSequenceNumber) report',
    ],
    [report],
  )
}

export function hashMessage(message: CCIPMessage) {
  const bytesMessage = ethers.utils.defaultAbiCoder.encode(
    [
      'tuple(uint256 sequenceNumber, uint256 sourceChainId, uint256 destinationChainId, address sender, tuple(address receiver, bytes data, address[] tokens, uint256[] amounts, address executor, bytes options) payload) message',
    ],
    [message],
  )
  // Add the leaf domain separator 0x00.
  return ethers.utils.solidityKeccak256(
    ['bytes', 'bytes'],
    ['0x00', bytesMessage],
  )
}

export function messageDeepEqual(
  actualMessage: any,
  expectedMessage: CCIPMessage,
) {
  expect(actualMessage?.sequenceNumber).to.equal(expectedMessage.sequenceNumber)
  expect(actualMessage?.sourceChainId).to.equal(expectedMessage.sourceChainId)
  expect(actualMessage?.destinationChainId).to.equal(
    expectedMessage.destinationChainId,
  )
  expect(actualMessage?.sender).to.equal(expectedMessage.sender)
  const actualMessagePayload = actualMessage?.payload
  expect(actualMessagePayload?.receiver).to.equal(
    expectedMessage.payload.receiver,
  )
  expect(actualMessagePayload?.data).to.equal(expectedMessage.payload.data)
  expect(actualMessagePayload.tokens).to.deep.equal(
    expectedMessage.payload.tokens,
  )
  const expectedAmounts = actualMessagePayload.amounts
  expect(actualMessagePayload.amounts.length).to.equal(expectedAmounts.length)
  for (let i = 0; i < expectedAmounts.length; i++) {
    const expectedAmount = expectedAmounts[i].toString()
    expect(actualMessagePayload.amounts[i].toString()).to.equal(expectedAmount)
  }
  expect(actualMessagePayload?.options).to.equal(
    expectedMessage.payload.options,
  )
}

export function requestEventArgsEqual(
  actualRequestArgs: any,
  expectedRequestArgs: any,
) {
  expect(actualRequestArgs?.sequenceNumber).to.equal(
    expectedRequestArgs.sequenceNumber,
  )

  expect(actualRequestArgs?.chainId).to.equal(expectedRequestArgs.chainId)
  expect(actualRequestArgs?.sender).to.equal(expectedRequestArgs.sender)
  expect(actualRequestArgs?.payload.receiver).to.equal(
    expectedRequestArgs.receiver,
  )
  expect(actualRequestArgs?.payload.data).to.equal(expectedRequestArgs.data)
  expect(actualRequestArgs.payload.tokens).to.deep.equal(
    expectedRequestArgs.tokens,
  )
  expect(actualRequestArgs.payload.amounts.length).to.equal(
    expectedRequestArgs.amounts.length,
  )
  for (let i = 0; i < expectedRequestArgs.amounts.length; i++) {
    const expectedAmount = expectedRequestArgs?.amounts[i].toString()
    expect(actualRequestArgs?.payload.amounts[i].toString()).to.equal(
      expectedAmount,
    )
  }
  expect(actualRequestArgs?.payload.options).to.equal(
    expectedRequestArgs?.options,
  )
}
                                                                                                                                                                                                                                                     contracts/src/v0.8/ccip/pools/LockUnlockPool.sol                                                    000644  000765  000024  00000004314 14165346401 022514  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT

pragma solidity ^0.8.6;

import "../../vendor/SafeERC20.sol";
import "../../vendor/Pausable.sol";
import "../interfaces/OnRampInterface.sol";
import "../interfaces/OffRampInterface.sol";
import "../interfaces/PoolInterface.sol";
import "../access/OwnerIsCreator.sol";

contract LockUnlockPool is PoolInterface, OwnerIsCreator, Pausable {
  using SafeERC20 for IERC20;

  IERC20 private immutable s_token;
  mapping(OnRampInterface => bool) private s_onRamps;
  mapping(OffRampInterface => bool) private s_offRamps;

  error PermissionsError();

  constructor(IERC20 token) {
    s_token = token;
  }

  function lockOrBurn(address depositor, uint256 amount) external override whenNotPaused onlyOwnerOrOnRamp {
    getToken().safeTransferFrom(depositor, address(this), amount);
    emit Locked(msg.sender, depositor, amount);
  }

  function releaseOrMint(address recipient, uint256 amount) external override whenNotPaused onlyOwnerOrOffRamp {
    getToken().safeTransfer(recipient, amount);
    emit Released(msg.sender, recipient, amount);
  }

  function pause() external override onlyOwner {
    _pause();
  }

  function unpause() external override onlyOwner {
    _unpause();
  }

  function setOnRamp(OnRampInterface onRamp, bool permission) public onlyOwner {
    s_onRamps[onRamp] = permission;
  }

  function setOffRamp(OffRampInterface offRamp, bool permission) public onlyOwner {
    s_offRamps[offRamp] = permission;
  }

  function isOnRamp(OnRampInterface onRamp) public view returns (bool) {
    return s_onRamps[onRamp];
  }

  function isOffRamp(OffRampInterface offRamp) public view returns (bool) {
    return s_offRamps[offRamp];
  }

  function getToken() public view override returns (IERC20 token) {
    return s_token;
  }

  function _validateOwnerOrOnRamp() internal view {
    if (msg.sender != owner() && !isOnRamp(OnRampInterface(msg.sender))) revert PermissionsError();
  }

  function _validateOwnerOrOffRamp() internal view {
    if (msg.sender != owner() && !isOffRamp(OffRampInterface(msg.sender))) revert PermissionsError();
  }

  modifier onlyOwnerOrOnRamp() {
    _validateOwnerOrOnRamp();
    _;
  }

  modifier onlyOwnerOrOffRamp() {
    _validateOwnerOrOffRamp();
    _;
  }
}
                                                                                                                                                                                                                                                                                                                    contracts/src/v0.8/ccip/test/mocks/MockOffRamp.sol                                                  000644  000765  000024  00000001553 14165346401 022723  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../../interfaces/OffRampInterface.sol";

contract MockOffRamp is OffRampInterface {
  event MessageExecuted(bytes32[] proof, CCIP.Message message, uint256 index);

  IERC20 public s_token;

  function deliverMessageTo(CrossChainMessageReceiverInterface recipient, CCIP.Message calldata message) external {
    recipient.receiveMessage(message);
  }

  function SOURCE_CHAIN_ID() external view returns (uint256) {}

  function CHAIN_ID() external view returns (uint256) {}

  function executeTransaction(
    bytes32[] memory proof,
    CCIP.Message memory message,
    uint256 index
  ) external override {
    emit MessageExecuted(proof, message, index);
  }

  function setToken(IERC20 token) external {
    s_token = token;
  }

  function TOKEN() external view returns (IERC20) {
    return s_token;
  }
}
                                                                                                                                                     contracts/src/v0.8/ccip/test/mocks/MockOnRamp.sol                                                   000644  000765  000024  00000002072 14165346401 022562  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../interfaces/OnRampInterface.sol";

contract MockOnRamp is OnRampInterface {
  uint256 public immutable CHAIN_ID;
  IERC20 public immutable TOKEN;
  IERC20 public immutable DESTINATION_TOKEN;
  PoolInterface public immutable POOL;
  uint256 public immutable DESTINATION_CHAIN_ID;

  CCIP.MessagePayload public mp;

  constructor(
    uint256 chainId,
    IERC20 token,
    IERC20 destinationToken,
    PoolInterface pool,
    uint256 destinationChainId
  ) {
    CHAIN_ID = chainId;
    TOKEN = token;
    DESTINATION_TOKEN = destinationToken;
    POOL = pool;
    DESTINATION_CHAIN_ID = destinationChainId;
  }

  function requestCrossChainSend(CCIP.MessagePayload calldata payload) external override returns (uint256) {
    mp = payload;
    return 0;
  }

  function getMessagePayload()
    external
    view
    returns (
      address,
      bytes memory,
      IERC20[] memory,
      uint256[] memory,
      bytes memory
    )
  {
    return (mp.receiver, mp.data, mp.tokens, mp.amounts, mp.options);
  }
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      contracts/src/v0.8/ccip/test/mocks/MockERC20.sol                                                    000644  000765  000024  00000001416 14165346401 022142  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT

pragma solidity ^0.8.6;

import "../../../vendor/ERC20.sol";

contract MockERC20 is ERC20 {
  constructor(
    string memory name,
    string memory symbol,
    address initialAccount,
    uint256 initialBalance
  ) payable ERC20(name, symbol) {
    _mint(initialAccount, initialBalance);
  }

  function mint(address account, uint256 amount) public {
    _mint(account, amount);
  }

  function burn(address account, uint256 amount) public {
    _burn(account, amount);
  }

  function transferInternal(
    address from,
    address to,
    uint256 value
  ) public {
    _transfer(from, to, value);
  }

  function approveInternal(
    address owner,
    address spender,
    uint256 value
  ) public {
    _approve(owner, spender, value);
  }
}
                                                                                                                                                                                                                                                  contracts/src/v0.8/ccip/test/mocks/MockAFN.sol                                                      000644  000765  000024  00000002026 14165346401 021771  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../../interfaces/AFNInterface.sol";

contract MockAFN is AFNInterface {
  Heartbeat public s_lastHeartbeat;
  bool public s_badSignal;

  constructor() {
    s_lastHeartbeat = Heartbeat({round: 1, timestamp: block.timestamp, committeeVersion: 1});
  }

  function setTimestamp(uint64 newTimestamp) external {
    s_lastHeartbeat.timestamp = newTimestamp;
  }

  function hasBadSignal() external view override returns (bool) {
    return s_badSignal;
  }

  function getLastHeartbeat() external view override returns (Heartbeat memory) {
    return s_lastHeartbeat;
  }

  function voteGood(
    uint256 /*round*/
  ) external override {
    s_badSignal = false;
  }

  function voteBad() external override {
    s_badSignal = true;
  }

  function recover() external override {
    s_badSignal = false;
  }

  function setConfig(
    address[] memory parties,
    uint256[] memory weights,
    uint256 goodQuorum,
    uint256 badQuorum
  ) external override {
    // nothing
  }
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          contracts/src/v0.8/ccip/test/helpers/TokenLimitsHelper.sol                                          000644  000765  000024  00000001335 14165346401 024505  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../../utils/TokenLimits.sol";

contract TokenLimitsHelper {
  using TokenLimits for TokenLimits.TokenBucket;

  TokenLimits.TokenBucket public s_bucket;

  event RemovalSuccess(bool success);

  function constructTokenBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) public {
    s_bucket = TokenLimits.constructTokenBucket(rate, capacity, full);
  }

  function alterCapacity(uint256 newCapacity) public {
    s_bucket.capacity = newCapacity;
  }

  function update() public {
    s_bucket.update();
  }

  function remove(uint256 tokens) public returns (bool removed) {
    removed = s_bucket.remove(tokens);
    emit RemovalSuccess(removed);
  }
}
                                                                                                                                                                                                                                                                                                   contracts/src/v0.8/ccip/test/helpers/SingleTokenOffRampHelper.sol                                   000644  000765  000024  00000001443 14165346401 025740  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../../ramps/SingleTokenOffRamp.sol";

contract SingleTokenOffRampHelper is SingleTokenOffRamp {
  constructor(
    uint256 sourceChainId,
    uint256 destinationChainId,
    IERC20 token,
    PoolInterface pool,
    uint256 tokenBucketRate,
    uint256 tokenBucketCapacity,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    uint256 executionDelaySeconds
  )
    SingleTokenOffRamp(
      sourceChainId,
      destinationChainId,
      token,
      pool,
      tokenBucketRate,
      tokenBucketCapacity,
      afn,
      maxTimeWithoutAFNSignal,
      executionDelaySeconds
    )
  {}

  /**
   * @dev Expose _report for tests
   */
  function report(bytes memory merkle) external {
    _report(bytes32(0), 0, merkle);
  }
}
                                                                                                                                                                                                                             contracts/src/v0.8/ccip/test/helpers/SimpleMessageReceiver.sol                                      000644  000765  000024  00000000724 14165346401 025327  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../../interfaces/CrossChainMessageReceiverInterface.sol";
import "../../interfaces/OffRampInterface.sol";

contract SimpleMessageReceiver is CrossChainMessageReceiverInterface {
  CCIP.Message public s_message;

  event MessageReceived(CCIP.Message message);

  function receiveMessage(CCIP.Message calldata message) external override {
    s_message = message;
    emit MessageReceived(message);
  }
}
                                            contracts/src/v0.8/ccip/test/helpers/MessageExecutorHelper.sol                                      000644  000765  000024  00000000521 14165346401 025342  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../../ramps/MessageExecutor.sol";

contract MessageExecutorHelper is MessageExecutor {
  constructor(OffRampInterface offRamp) MessageExecutor(offRamp) {}

  function report(bytes memory executableMessages) external {
    _report(bytes32(0), 0, executableMessages);
  }
}
                                                                                                                                                                               contracts/src/v0.8/ccip/test/helpers/HealthCheckerHelper.sol                                        000644  000765  000024  00000000477 14165346401 024743  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../../health/HealthChecker.sol";

contract HealthCheckerHelper is HealthChecker {
  constructor(AFNInterface afn, uint256 maxTimeWithoutAFNSignal) HealthChecker(afn, maxTimeWithoutAFNSignal) {}

  function whenHealthyFunction() external whenHealthy {}
}
                                                                                                                                                                                                 contracts/src/v0.8/ccip/access/OwnerIsCreator.sol                                                   000644  000765  000024  00000000450 14165346401 022626  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../ConfirmedOwner.sol";

/**
 * @title The OwnerIsCreator contract
 * @notice A contract with helpers for basic contract ownership.
 */
contract OwnerIsCreator is ConfirmedOwner {
  constructor() ConfirmedOwner(msg.sender) {}
}
                                                                                                                                                                                                                        contracts/src/v0.8/ccip/health/HealthChecker.sol                                                    000644  000765  000024  00000005613 14165346401 022424  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../../vendor/Pausable.sol";
import "../interfaces/AFNInterface.sol";
import "../access/OwnerIsCreator.sol";

contract HealthChecker is Pausable, OwnerIsCreator {
  // AFN contract to check health of the system
  AFNInterface private s_afn;
  // The maximum time since the last AFN heartbeat before it is considered unhealthy
  uint256 private s_maxSecondsWithoutAFNHeartbeat;

  error BadAFNSignal();
  error StaleAFNHeartbeat();
  error BadHealthConfig();

  event AFNSet(AFNInterface oldAFN, AFNInterface newAFN);
  event AFNMaxHeartbeatTimeSet(uint256 oldTime, uint256 newTime);

  /**
   * @param afn The AFN contract to check health
   * @param maxSecondsWithoutAFNHeartbeat maximum seconds allowed between heartbeats to consider
   * the network "healthy".
   */
  constructor(AFNInterface afn, uint256 maxSecondsWithoutAFNHeartbeat) {
    if (address(afn) == address(0) || maxSecondsWithoutAFNHeartbeat == 0) revert BadHealthConfig();
    s_afn = afn;
    s_maxSecondsWithoutAFNHeartbeat = maxSecondsWithoutAFNHeartbeat;
  }

  /**
   * @notice Pause the contract
   * @dev only callable by the owner
   */
  function pause() external onlyOwner {
    _pause();
  }

  /**
   * @notice Unpause the contract
   * @dev only callable by the owner
   */
  function unpause() external onlyOwner {
    _unpause();
  }

  /**
   * @notice Change the afn contract to track
   * @dev only callable by the owner
   * @param afn new AFN contract
   */
  function setAFN(AFNInterface afn) external onlyOwner {
    if (address(afn) == address(0)) revert BadHealthConfig();
    AFNInterface old = s_afn;
    s_afn = afn;
    emit AFNSet(old, afn);
  }

  /**
   * @notice Get the current AFN contract
   * @return Current AFN
   */
  function getAFN() external view returns (AFNInterface) {
    return s_afn;
  }

  /**
   * @notice Change the mixumum time allowed without a heartbeat
   * @dev only callable by the owner
   * @param newTime the new max time
   */
  function setMaxSecondsWithoutAFNHeartbeat(uint256 newTime) external onlyOwner {
    if (newTime == 0) revert BadHealthConfig();
    uint256 oldTime = s_maxSecondsWithoutAFNHeartbeat;
    s_maxSecondsWithoutAFNHeartbeat = newTime;
    emit AFNMaxHeartbeatTimeSet(oldTime, newTime);
  }

  /**
   * @notice Get the current max time without heartbeat
   * @return current max time
   */
  function getMaxSecondsWithoutAFNHeartbeat() external view returns (uint256) {
    return s_maxSecondsWithoutAFNHeartbeat;
  }

  /**
   * @notice Ensure that the AFN has not emitted a bad signal, and that the latest heartbeat is not stale.
   */
  modifier whenHealthy() {
    if (s_afn.hasBadSignal()) revert BadAFNSignal();
    AFNInterface.Heartbeat memory lastHeartbeat = s_afn.getLastHeartbeat();
    if ((block.timestamp - uint256(lastHeartbeat.timestamp)) > s_maxSecondsWithoutAFNHeartbeat)
      revert StaleAFNHeartbeat();
    _;
  }
}
                                                                                                                     contracts/src/v0.8/ccip/health/AFN.sol                                                              000644  000765  000024  00000014177 14165346401 020343  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../interfaces/AFNInterface.sol";
import "../access/OwnerIsCreator.sol";

contract AFN is AFNInterface, OwnerIsCreator {
  // AFN party addresses and weights
  mapping(address => uint256) private s_parties;
  // List of AFN party addresses
  address[] private s_partyList;
  // Quorum of good votes to reach
  uint256 private s_goodQuorum;
  // Quorum of bad votes to reach
  uint256 private s_badQuorum;
  // The current round ID
  uint256 private s_round;
  // Version of the set of parties
  uint256 private s_committeeVersion;

  // Last heartbeat
  Heartbeat private s_lastHeartbeat;
  // The last round that a party voted good
  mapping(address => uint256) private s_lastGoodVote;
  // round => total good votes
  mapping(uint256 => uint256) private s_goodVotes;

  // Has a party voted bad
  mapping(address => bool) private s_hasVotedBad;
  // Parties that have voted bad
  address[] private s_badVoters;
  // Total bad votes
  uint256 private s_badVotes;
  // Whether or not there is a bad signal
  bool private s_badSignal;

  constructor(
    address[] memory parties,
    uint256[] memory weights,
    uint256 goodQuorum,
    uint256 badQuorum
  ) {
    _setConfig(parties, weights, goodQuorum, badQuorum, 1, 1);
  }

  ////////  VOTING  ////////

  /**
   * @notice Submit a good vote
   * @dev msg.sender must be a registered party
   * @param round the current round
   */
  function voteGood(uint256 round) external override {
    uint256 currentRound = s_round;
    if (round != currentRound) revert IncorrectRound(currentRound, round);
    if (s_badSignal) revert MustRecoverFromBadSignal();
    address sender = msg.sender;
    if (s_parties[sender] == 0) revert InvalidVoter(sender);
    if (s_lastGoodVote[sender] == currentRound) revert AlreadyVoted();

    s_lastGoodVote[sender] = currentRound;
    s_goodVotes[currentRound] += s_parties[sender];
    emit GoodVote(sender, currentRound);

    if (s_goodVotes[currentRound] >= s_goodQuorum) {
      Heartbeat memory heartbeat = Heartbeat({
        round: currentRound,
        timestamp: uint64(block.timestamp),
        committeeVersion: s_committeeVersion
      });
      s_lastHeartbeat = heartbeat;
      s_round++;
      emit AFNHeartbeat(heartbeat);
    }
  }

  /**
   * @notice Submit a bad vote
   * @dev msg.sender must be a registered party
   */
  function voteBad() external override {
    if (s_badSignal) revert MustRecoverFromBadSignal();
    address sender = msg.sender;
    uint256 senderWeight = s_parties[sender];
    if (senderWeight == 0) revert InvalidVoter(sender);
    if (s_hasVotedBad[sender]) revert AlreadyVoted();

    s_hasVotedBad[sender] = true;
    s_badVoters.push(sender);
    s_badVotes += senderWeight;

    if (s_badVotes >= s_badQuorum) {
      s_badSignal = true;
      emit AFNBadSignal(block.timestamp);
    }
  }

  ////////  OnlyOwner ////////

  /**
   * @notice Recover from a bad signal
   * @dev only callable by the owner
   */
  function recover() external override onlyOwner {
    if (!s_badSignal) revert RecoveryNotNecessary();
    address[] memory badVoters = s_badVoters;
    for (uint256 i = 0; i < badVoters.length; i++) {
      s_hasVotedBad[badVoters[i]] = false;
    }
    s_badVotes = 0;
    delete s_badVoters;
    s_badSignal = false;
    emit RecoveredFromBadSignal();
  }

  /**
   * @notice Set config storage vars
   * @dev only callable by the owner
   * @param parties parties allowed to vote
   * @param weights weights of each party's vote
   * @param goodQuorum threshold to emit a heartbeat
   * @param badQuorum threashold to emit a bad signal
   */
  function setConfig(
    address[] memory parties,
    uint256[] memory weights,
    uint256 goodQuorum,
    uint256 badQuorum
  ) external override onlyOwner {
    _setConfig(parties, weights, goodQuorum, badQuorum, s_round + 1, s_committeeVersion + 1);
  }

  ////////  Views ////////

  function hasBadSignal() external view override returns (bool) {
    return s_badSignal;
  }

  function getLastHeartbeat() external view override returns (Heartbeat memory) {
    return s_lastHeartbeat;
  }

  function getQuorums() external view returns (uint256 good, uint256 bad) {
    return (s_goodQuorum, s_badQuorum);
  }

  function getParties() external view returns (address[] memory) {
    return s_partyList;
  }

  function getWeight(address party) external view returns (uint256) {
    return s_parties[party];
  }

  function getRound() external view returns (uint256) {
    return s_round;
  }

  function getCommitteeVersion() external view returns (uint256) {
    return s_committeeVersion;
  }

  function getLastGoodVote(address party) external view returns (uint256) {
    return s_lastGoodVote[party];
  }

  function getGoodVotes(uint256 round) external view returns (uint256) {
    return s_goodVotes[round];
  }

  function getBadVotersAndVotes() external view returns (address[] memory voters, uint256 votes) {
    return (s_badVoters, s_badVotes);
  }

  function hasVotedBad(address party) external view returns (bool) {
    return s_hasVotedBad[party];
  }

  ////////  Private ////////

  /**
   * @notice Set detailed config storage vars
   */
  function _setConfig(
    address[] memory parties,
    uint256[] memory weights,
    uint256 goodQuorum,
    uint256 badQuorum,
    uint256 round,
    uint256 committeeVersion
  ) private {
    if (
      parties.length != weights.length ||
      parties.length == 0 ||
      goodQuorum == 0 ||
      badQuorum == 0 ||
      round == 0 ||
      committeeVersion == 0
    ) {
      revert InvalidConfig();
    }
    // Unset existing parties
    address[] memory existingParties = s_partyList;
    for (uint256 i = 0; i < existingParties.length; i++) {
      s_parties[existingParties[i]] = 0;
    }

    // Update round, committee and quorum details
    s_goodQuorum = goodQuorum;
    s_badQuorum = badQuorum;
    s_round = round;
    s_committeeVersion = committeeVersion;

    // Set new parties
    s_partyList = parties;
    for (uint256 i = 0; i < parties.length; i++) {
      if (weights[i] == 0) revert InvalidWeight();
      s_parties[parties[i]] = weights[i];
    }
    emit ConfigSet(parties, weights, goodQuorum, badQuorum);
  }
}
                                                                                                                                                                                                                                                                                                                                                                                                 contracts/src/v0.8/ccip/ocr/OCR2Base.sol                                                            000644  000765  000024  00000032772 14165346401 020556  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../access/OwnerIsCreator.sol";
import "../../interfaces/TypeAndVersionInterface.sol";

/**
  * @notice Onchain verification of reports from the offchain reporting protocol

  * @dev For details on its operation, see the offchain reporting protocol design
  * doc, which refers to this contract as simply the "contract".

  * @dev This contract is meant to aid rapid development of new applications based on OCR2.
  * However, for actual production contracts, it is expected that most of the logic of this contract
  * will be folded directly into the application contract. Inheritance prevents us from doing lots
  * of juicy storage layout optimizations, leading to a substantial increase in gas cost.

  * @dev THIS CONTRACT HAS NOT GONE THROUGH ANY SECURITY REVIEW. DO NOT USE IN PROD
*/
abstract contract OCR2Base is OwnerIsCreator, TypeAndVersionInterface {
  bool internal immutable UNIQUE_REPORTS;

  constructor(bool uniqueReports) {
    UNIQUE_REPORTS = uniqueReports;
  }

  uint256 private constant maxUint32 = (1 << 32) - 1;

  // Maximum number of oracles the offchain reporting protocol is designed for
  uint256 internal constant maxNumOracles = 31;

  // Storing these fields used on the hot path in a ConfigInfo variable reduces the
  // retrieval of all of them to a single SLOAD. If any further fields are
  // added, make sure that storage of the struct still takes at most 32 bytes.
  struct ConfigInfo {
    bytes32 latestConfigDigest;
    uint8 f; // TODO: could be optimized by squeezing into one slot
    uint8 n;
  }
  ConfigInfo internal s_configInfo;

  // incremented each time a new config is posted. This count is incorporated
  // into the config digest, to prevent replay attacks.
  uint32 internal s_configCount;
  uint32 internal s_latestConfigBlockNumber; // makes it easier for offchain systems
  // to extract config from logs.

  // Used for s_oracles[a].role, where a is an address, to track the purpose
  // of the address, or to indicate that the address is unset.
  enum Role {
    // No oracle role has been set for address a
    Unset,
    // Signing address for the s_oracles[a].index'th oracle. I.e., report
    // signatures from this oracle should ecrecover back to address a.
    Signer,
    // Transmission address for the s_oracles[a].index'th oracle. I.e., if a
    // report is received by OCR2Aggregator.transmit in which msg.sender is
    // a, it is attributed to the s_oracles[a].index'th oracle.
    Transmitter
  }

  struct Oracle {
    uint8 index; // Index of oracle in s_signers/s_transmitters
    Role role; // Role of the address which mapped to this struct
  }

  mapping(address => Oracle) /* signer OR transmitter address */
    internal s_oracles;

  // s_signers contains the signing address of each oracle
  address[] internal s_signers;

  // s_transmitters contains the transmission address of each oracle,
  // i.e. the address the oracle actually sends transactions to the contract from
  address[] internal s_transmitters;

  /*
   * Config logic
   */

  /**
   * @notice triggers a new run of the offchain reporting protocol
   * @param previousConfigBlockNumber block in which the previous config was set, to simplify historic analysis
   * @param configCount ordinal number of this config setting among all config settings over the life of this contract
   * @param signers ith element is address ith oracle uses to sign a report
   * @param transmitters ith element is address ith oracle uses to transmit a report via the transmit method
   * @param f maximum number of faulty/dishonest oracles the protocol can tolerate while still working correctly
   * @param encodedConfigVersion version of the serialization format used for "encoded" parameter
   * @param encoded serialized data used by oracles to configure their offchain operation
   */
  event ConfigSet(
    uint32 previousConfigBlockNumber,
    bytes32 configDigest,
    uint64 configCount,
    address[] signers,
    address[] transmitters,
    uint8 f,
    bytes onchainConfig,
    uint64 encodedConfigVersion,
    bytes encoded
  );

  // Reverts transaction if config args are invalid
  modifier checkConfigValid(
    uint256 _numSigners,
    uint256 _numTransmitters,
    uint256 _f
  ) {
    require(_numSigners <= maxNumOracles, "too many signers");
    require(_f > 0, "f must be positive");
    require(_numSigners == _numTransmitters, "oracle addresses out of registration");
    require(_numSigners > 3 * _f, "faulty-oracle f too high");
    _;
  }

  struct SetConfigArgs {
    address[] signers;
    address[] transmitters;
    uint8 f;
    bytes onchainConfig;
    uint64 offchainConfigVersion;
    bytes offchainConfig;
  }

  /**
   * @notice sets offchain reporting protocol configuration incl. participating oracles
   * @param _signers addresses with which oracles sign the reports
   * @param _transmitters addresses oracles use to transmit the reports
   * @param _f number of faulty oracles the system can tolerate
   * @param _onchainConfig encoded on-chain contract configuration
   * @param _offchainConfigVersion version number for offchainEncoding schema
   * @param _offchainConfig encoded off-chain oracle configuration
   */
  function setConfig(
    address[] memory _signers,
    address[] memory _transmitters,
    uint8 _f,
    bytes memory _onchainConfig,
    uint64 _offchainConfigVersion,
    bytes memory _offchainConfig
  ) external checkConfigValid(_signers.length, _transmitters.length, _f) onlyOwner {
    SetConfigArgs memory args = SetConfigArgs({
      signers: _signers,
      transmitters: _transmitters,
      f: _f,
      onchainConfig: _onchainConfig,
      offchainConfigVersion: _offchainConfigVersion,
      offchainConfig: _offchainConfig
    });

    _beforeSetConfig(args.f, args.onchainConfig);

    while (s_signers.length != 0) {
      // remove any old signer/transmitter addresses
      uint256 lastIdx = s_signers.length - 1;
      address signer = s_signers[lastIdx];
      address transmitter = s_transmitters[lastIdx];
      delete s_oracles[signer];
      delete s_oracles[transmitter];
      s_signers.pop();
      s_transmitters.pop();
    }

    for (uint256 i = 0; i < args.signers.length; i++) {
      // add new signer/transmitter addresses
      require(s_oracles[args.signers[i]].role == Role.Unset, "repeated signer address");
      s_oracles[args.signers[i]] = Oracle(uint8(i), Role.Signer);
      require(s_oracles[args.transmitters[i]].role == Role.Unset, "repeated transmitter address");
      s_oracles[args.transmitters[i]] = Oracle(uint8(i), Role.Transmitter);
      s_signers.push(args.signers[i]);
      s_transmitters.push(args.transmitters[i]);
    }
    s_configInfo.f = args.f;
    uint32 previousConfigBlockNumber = s_latestConfigBlockNumber;
    s_latestConfigBlockNumber = uint32(block.number);
    s_configCount += 1;
    {
      s_configInfo.latestConfigDigest = configDigestFromConfigData(
        block.chainid,
        address(this),
        s_configCount,
        args.signers,
        args.transmitters,
        args.f,
        args.onchainConfig,
        args.offchainConfigVersion,
        args.offchainConfig
      );
    }
    s_configInfo.n = uint8(args.signers.length);

    emit ConfigSet(
      previousConfigBlockNumber,
      s_configInfo.latestConfigDigest,
      s_configCount,
      args.signers,
      args.transmitters,
      args.f,
      args.onchainConfig,
      args.offchainConfigVersion,
      args.offchainConfig
    );

    _afterSetConfig(args.f, args.onchainConfig);
  }

  function configDigestFromConfigData(
    uint256 _chainId,
    address _contractAddress,
    uint64 _configCount,
    address[] memory _signers,
    address[] memory _transmitters,
    uint8 _f,
    bytes memory _onchainConfig,
    uint64 _encodedConfigVersion,
    bytes memory _encodedConfig
  ) internal pure returns (bytes32) {
    uint256 h = uint256(
      keccak256(
        abi.encode(
          _chainId,
          _contractAddress,
          _configCount,
          _signers,
          _transmitters,
          _f,
          _onchainConfig,
          _encodedConfigVersion,
          _encodedConfig
        )
      )
    );
    uint256 prefixMask = type(uint256).max << (256 - 16); // 0xFFFF00..00
    uint256 prefix = 0x0001 << (256 - 16); // 0x000100..00
    return bytes32((prefix & prefixMask) | (h & ~prefixMask));
  }

  /**
   * @notice information about current offchain reporting protocol configuration

   * @return configCount ordinal number of current config, out of all configs applied to this contract so far
   * @return blockNumber block at which this config was set
   * @return configDigest domain-separation tag for current config (see configDigestFromConfigData)
   */
  function latestConfigDetails()
    external
    view
    returns (
      uint32 configCount,
      uint32 blockNumber,
      bytes32 configDigest
    )
  {
    return (s_configCount, s_latestConfigBlockNumber, s_configInfo.latestConfigDigest);
  }

  /**
   * @return list of addresses permitted to transmit reports to this contract

   * @dev The list will match the order used to specify the transmitter during setConfig
   */
  function transmitters() external view returns (address[] memory) {
    return s_transmitters;
  }

  function _beforeSetConfig(uint8 _f, bytes memory _onchainConfig) internal virtual;

  function _afterSetConfig(uint8 _f, bytes memory _onchainConfig) internal virtual;

  function _report(
    bytes32 configDigest,
    uint40 epochAndRound,
    bytes memory report
  ) internal virtual;

  function _payTransmitter(uint32 initialGas, address transmitter) internal virtual;

  // The constant-length components of the msg.data sent to transmit.
  // See the "If we wanted to call sam" example on for example reasoning
  // https://solidity.readthedocs.io/en/v0.7.2/abi-spec.html
  uint16 private constant TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT =
    4 + // function selector
      32 *
      3 + // 3 words containing reportContext
      32 + // word containing start location of abiencoded report value
      32 + // word containing location start of abiencoded rs value
      32 + // word containing start location of abiencoded ss value
      32 + // rawVs value
      32 + // word containing length of report
      32 + // word containing length rs
      32 + // word containing length of ss
      0; // placeholder

  function requireExpectedMsgDataLength(
    bytes calldata report,
    bytes32[] calldata rs,
    bytes32[] calldata ss
  ) private pure {
    // calldata will never be big enough to make this overflow
    uint256 expected = uint256(TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT) +
      report.length + // one byte pure entry in _report
      rs.length *
      32 + // 32 bytes per entry in _rs
      ss.length *
      32 + // 32 bytes per entry in _ss
      0; // placeholder
    require(msg.data.length == expected, "calldata length mismatch");
  }

  event Transmited(bytes32 configDigest, uint32 epoch);

  /**
   * @notice transmit is called to post a new report to the contract
   * @param report serialized report, which the signatures are signing.
   * @param rs ith element is the R components of the ith signature on report. Must have at most maxNumOracles entries
   * @param ss ith element is the S components of the ith signature on report. Must have at most maxNumOracles entries
   * @param rawVs ith element is the the V component of the ith signature
   */
  function transmit(
    // NOTE: If these parameters are changed, expectedMsgDataLength and/or
    // TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT need to be changed accordingly
    bytes32[3] calldata reportContext,
    bytes calldata report,
    bytes32[] calldata rs,
    bytes32[] calldata ss,
    bytes32 rawVs // signatures
  ) external {
    uint256 initialGas = gasleft(); // This line must come first

    {
      // reportContext consists of:
      // reportContext[0]: ConfigDigest
      // reportContext[1]: 27 byte padding, 4-byte epoch and 1-byte round
      // reportContext[2]: ExtraHash
      bytes32 configDigest = reportContext[0];
      uint40 epochAndRound = uint40(uint256(reportContext[1]));

      _report(configDigest, epochAndRound, report);

      emit Transmited(configDigest, uint32(epochAndRound >> 8));

      ConfigInfo memory configInfo = s_configInfo;
      require(configInfo.latestConfigDigest == configDigest, "configDigest mismatch");

      requireExpectedMsgDataLength(report, rs, ss);

      uint256 expectedNumSignatures;
      if (UNIQUE_REPORTS) {
        expectedNumSignatures = (configInfo.n + configInfo.f) / 2 + 1;
      } else {
        expectedNumSignatures = configInfo.f + 1;
      }

      require(rs.length == expectedNumSignatures, "wrong number of signatures");
      require(rs.length == ss.length, "signatures out of registration");

      Oracle memory transmitter = s_oracles[msg.sender];
      require( // Check that sender is authorized to report
        transmitter.role == Role.Transmitter && msg.sender == s_transmitters[transmitter.index],
        "unauthorized transmitter"
      );
    }

    {
      // Verify signatures attached to report
      bytes32 h = keccak256(abi.encodePacked(keccak256(report), reportContext));
      bool[maxNumOracles] memory signed;

      Oracle memory o;
      for (uint256 i = 0; i < rs.length; i++) {
        address signer = ecrecover(h, uint8(rawVs[i]) + 27, rs[i], ss[i]);
        o = s_oracles[signer];
        require(o.role == Role.Signer, "address not authorized to sign");
        require(!signed[o.index], "non-unique signature");
        signed[o.index] = true;
      }
    }

    assert(initialGas < maxUint32);
    _payTransmitter(uint32(initialGas), msg.sender);
  }
}
      contracts/src/v0.8/ccip/utils/TokenLimits.sol                                                       000644  000765  000024  00000005235 14165346401 022067  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

/**
 * @notice This library enables token rate limiting using a `TokenBucket`.
 * The bucket holds the number of tokens that can be transferred at any
 * given time. It has:
 *  - capacity: maximum number of tokens possible
 *  - rate: rate at which the bucket refills per second
 *  - tokens: current number of tokens in the bucket
 *  - lastUpdated: timestamp of the last refill
 */
library TokenLimits {
  // Token Bucket used for rate limiting
  struct TokenBucket {
    uint256 rate;
    uint256 capacity;
    uint256 tokens;
    uint256 lastUpdated;
  }

  error TimeError();
  error BucketOverfilled();

  /**
   * @notice Create a fresh token bucket
   * @param rate Refill rate
   * @param capacity Maximum capacity of the bucket
   * @return tokenBucket
   */
  function constructTokenBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) internal view returns (TokenBucket memory) {
    uint256 tokens = full ? capacity : 0;
    return TokenBucket({rate: rate, capacity: capacity, tokens: tokens, lastUpdated: block.timestamp});
  }

  /**
   * @notice Remove tokens from the buck if possible.
   * @dev This acts upon a storage variable in the calling contract.
   * @param bucket token bucket (MUST BE STORAGE)
   * @param tokens number of tokens
   * @return tokens removed (true if removed, false otherwise)
   */
  function remove(TokenBucket storage bucket, uint256 tokens) internal returns (bool) {
    // Refill the bucket if possible
    update(bucket);
    // Remove tokens if available in bucket
    if (bucket.tokens < tokens) return false;
    bucket.tokens -= tokens;
    return true;
  }

  /**
   * @notice Update the tokens in the bucket
   * @dev Uses the `rate` and block timestamp to refill the bucket.
   * @dev This acts upon a storage variable in the calling contract.
   * @param bucket token bucket (MUST BE STORAGE)
   */
  function update(TokenBucket storage bucket) internal {
    // Revert if the tokens in the bucket exceed its capacity
    if (bucket.tokens > bucket.capacity) revert BucketOverfilled();
    // Return if there's nothing to update
    if (bucket.tokens == bucket.capacity) return;
    uint256 timeNow = block.timestamp;
    if (timeNow < bucket.lastUpdated) revert TimeError();
    uint256 difference = timeNow - bucket.lastUpdated;
    bucket.tokens = min(bucket.capacity, bucket.tokens + difference * bucket.rate);
    bucket.lastUpdated = timeNow;
  }

  /**
   * @notice Return the smallest of two integers
   * @param a first int
   * @param b second int
   * @return smallest
   */
  function min(uint256 a, uint256 b) private pure returns (uint256) {
    return a < b ? a : b;
  }
}
                                                                                                                                                                                                                                                                                                                                                                   contracts/src/v0.8/ccip/utils/CCIP.sol                                                              000644  000765  000024  00000001257 14165346401 020343  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../vendor/IERC20.sol";

contract CCIP {
  /// @notice High level message
  struct Message {
    uint256 sequenceNumber;
    uint256 sourceChainId;
    uint256 destinationChainId;
    address sender;
    MessagePayload payload;
  }

  /// @notice Payload within the message
  struct MessagePayload {
    address receiver;
    bytes data;
    IERC20[] tokens;
    uint256[] amounts;
    address executor;
    bytes options;
  }

  /// @notice Report that is relayed by the observing DON at the relay phase
  struct RelayReport {
    bytes32 merkleRoot;
    uint256 minSequenceNumber;
    uint256 maxSequenceNumber;
  }
}
                                                                                                                                                                                                                                                                                                                                                 contracts/src/v0.8/ccip/ramps/SingleTokenOffRamp.sol                                                000644  000765  000024  00000021005 14165346401 023275  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../interfaces/OffRampInterface.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../ocr/OCR2Base.sol";
import "../utils/CCIP.sol";
import "../utils/TokenLimits.sol";
import "../health/HealthChecker.sol";
import "../../vendor/Address.sol";

contract SingleTokenOffRamp is OffRampInterface, TypeAndVersionInterface, HealthChecker, OCR2Base {
  using Address for address;
  using TokenLimits for TokenLimits.TokenBucket;

  // Chain ID of the source chain
  uint256 public immutable SOURCE_CHAIN_ID;
  // Chain ID of this chain
  uint256 public immutable CHAIN_ID;
  // Token pool contract
  PoolInterface public immutable POOL;
  // Token contract
  IERC20 public immutable TOKEN;
  // Offchain leaf domain separator
  bytes1 private constant LEAF_DOMAIN_SEPARATOR = 0x00;
  // Internal domain separator used in proofs
  bytes1 private constant INTERNAL_DOMAIN_SEPARATOR = 0x01;
  // merkleRoot => timestamp when received
  mapping(bytes32 => uint256) private s_merkleRoots;
  // sequenceNumber => executed
  mapping(uint256 => bool) private s_executed;
  // execution delay in seconds
  uint256 private s_executionDelaySeconds;
  // Last relay report
  CCIP.RelayReport private s_lastReport;
  // Token bucket for token rate limiting
  TokenLimits.TokenBucket private s_tokenBucket;

  constructor(
    uint256 sourceChainId,
    uint256 chainId,
    IERC20 token,
    PoolInterface pool,
    uint256 tokenBucketRate,
    uint256 tokenBucketCapacity,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    uint256 executionDelaySeconds
  ) OCR2Base(true) HealthChecker(afn, maxTimeWithoutAFNSignal) {
    if (pool.getToken() != token) revert TokenMismatch();
    SOURCE_CHAIN_ID = sourceChainId;
    CHAIN_ID = chainId;
    TOKEN = token;
    POOL = pool;
    s_tokenBucket = TokenLimits.constructTokenBucket(tokenBucketRate, tokenBucketCapacity, true);
    s_executionDelaySeconds = executionDelaySeconds;
  }

  /**
   * @notice Extending OCR2Base._report
   * @dev assumes the report is a bytes encoded bytes32 merkle root
   * @dev will be called by Chainlink nodes on transmit()
   */
  function _report(
    bytes32, /*configDigest*/
    uint40, /*epochAndRound*/
    bytes memory report
  ) internal override whenNotPaused whenHealthy {
    CCIP.RelayReport memory newRelayReport = abi.decode(report, (CCIP.RelayReport));
    // check that the sequence numbers make sense
    if (newRelayReport.minSequenceNumber > newRelayReport.maxSequenceNumber) revert RelayReportError();
    CCIP.RelayReport memory lastRelayReport = s_lastReport;
    // if this is not the first relay report, make sure the sequence numbers
    // are greater than the previous report.
    if (lastRelayReport.merkleRoot != bytes32(0)) {
      if (newRelayReport.minSequenceNumber != lastRelayReport.maxSequenceNumber + 1) {
        revert SequenceError(lastRelayReport.maxSequenceNumber, newRelayReport.minSequenceNumber);
      }
    }

    s_merkleRoots[newRelayReport.merkleRoot] = block.timestamp;
    s_lastReport = newRelayReport;
    emit ReportAccepted(newRelayReport);
  }

  /**
   * @notice Execute a specific payload
   * @param proof Merkle proof in the order bottom to top of the tree
   * @param message Message that is to be sent
   * @param index Index of the leaf
   * @dev Can be called by anyone
   */
  function executeTransaction(
    bytes32[] memory proof,
    CCIP.Message memory message,
    uint256 index
  ) external override whenNotPaused whenHealthy {
    // Verify merkle proof
    // The leaf offchain is keccak256(LEAF_DOMAIN_SEPARATOR || CrossChainSendRequested event data),
    // where the CrossChainSendRequested event data is abi.encode(CCIP.Message).
    bytes32 leaf = keccak256(abi.encodePacked(LEAF_DOMAIN_SEPARATOR, abi.encode(message)));

    // Get root from proof
    bytes32 root = generateMerkleRoot(proof, leaf, index);

    // Check that root has been relayed
    uint256 reportTimestamp = s_merkleRoots[root];
    if (reportTimestamp == 0) revert MerkleProofError(proof, message, index);

    // Execution delay
    if (reportTimestamp + s_executionDelaySeconds >= block.timestamp) revert ExecutionDelayError();

    // Disallow double-execution.
    if (s_executed[message.sequenceNumber]) revert AlreadyExecuted(message.sequenceNumber);

    // The transaction can only be executed by the designated executor, if one exists.
    if (message.payload.executor != address(0) && message.payload.executor != msg.sender)
      revert InvalidExecutor(message.sequenceNumber);

    // Validity checks for the message.
    _isWellFormed(message);

    // Avoid shooting ourselves in the foot by disallowing calls to some
    // privileged OffRamp function as OffRamp.
    // In the wild: https://rekt.news/polynetwork-rekt/
    _validateReceiver(message);

    // Mark as executed before external calls
    s_executed[message.sequenceNumber] = true;

    // Remove the tokens from the rate limiting bucket
    uint256 numberOfTokens = message.payload.amounts[0];
    if (!s_tokenBucket.remove(numberOfTokens)) revert ExceedsTokenLimit(s_tokenBucket.tokens, numberOfTokens);

    // Release tokens to receiver
    POOL.releaseOrMint(message.payload.receiver, message.payload.amounts[0]);

    // Try send the message, emit fulfillment error if fails
    try CrossChainMessageReceiverInterface(message.payload.receiver).receiveMessage(message) {
      emit CrossChainMessageExecuted(message.sequenceNumber);
    } catch (bytes memory reason) {
      revert ExecutionError(message.sequenceNumber, reason);
    }
  }

  /**
   * @notice Generate a Merkle Root from Proof.
   * @param proof Merkle proof in the order bottom to top of the tree
   * @param leaf bytes32 leaf hash
   * @param index Index of the leaf
   * @return bytes32 root generated by proof
   */
  function generateMerkleRoot(
    bytes32[] memory proof,
    bytes32 leaf,
    uint256 index
  ) public pure returns (bytes32) {
    bytes32 hash = leaf;

    for (uint256 i = 0; i < proof.length; i++) {
      bytes32 proofElement = proof[i];

      if (index % 2 == 0) {
        hash = keccak256(abi.encodePacked(INTERNAL_DOMAIN_SEPARATOR, hash, proofElement));
      } else {
        hash = keccak256(abi.encodePacked(INTERNAL_DOMAIN_SEPARATOR, proofElement, hash));
      }
      index = index / 2;
    }
    return hash;
  }

  /**
   * @notice Message receiver checks
   */
  function _validateReceiver(CCIP.Message memory message) private view {
    if (
      address(message.payload.receiver) == address(this) ||
      address(message.payload.receiver) == address(POOL) ||
      address(message.payload.receiver) == address(TOKEN) ||
      !address(message.payload.receiver).isContract()
    ) revert InvalidReceiver(message.payload.receiver);
  }

  function _isWellFormed(CCIP.Message memory message) private view {
    if (message.sourceChainId != SOURCE_CHAIN_ID) revert InvalidSourceChain(message.sourceChainId);
    if (message.payload.tokens.length != 1 || message.payload.amounts.length != 1) revert UnsupportedNumberOfTokens();
    if (message.payload.tokens[0] != TOKEN) revert UnsupportedToken(message.payload.tokens[0]);
  }

  function _beforeSetConfig(uint8 _threshold, bytes memory _onchainConfig) internal override {
    // TODO
  }

  function _afterSetConfig(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {
    // TODO
  }

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {
    // TODO
  }

  function configureTokenBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) external onlyOwner {
    s_tokenBucket = TokenLimits.constructTokenBucket(rate, capacity, full);
    emit NewTokenBucketConstructed(rate, capacity, full);
  }

  function getTokenBucket() external view returns (TokenLimits.TokenBucket memory) {
    return s_tokenBucket;
  }

  function setExecutionDelaySeconds(uint256 executionDelaySeconds) external onlyOwner {
    s_executionDelaySeconds = executionDelaySeconds;
    emit ExecutionDelaySecondsSet(executionDelaySeconds);
  }

  function getExecutionDelaySeconds() external view returns (uint256) {
    return s_executionDelaySeconds;
  }

  function getMerkleRoot(bytes32 root) external view returns (uint256) {
    return s_merkleRoots[root];
  }

  function getExecuted(uint256 sequenceNumber) external view returns (bool) {
    return s_executed[sequenceNumber];
  }

  function getLastReport() external view returns (CCIP.RelayReport memory) {
    return s_lastReport;
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "SingleTokenOffRamp 1.1.0";
  }
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           contracts/src/v0.8/ccip/ramps/SingleTokenOnRamp.sol                                                 000644  000765  000024  00000012505 14165346401 023144  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../interfaces/OnRampInterface.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../utils/CCIP.sol";
import "../utils/TokenLimits.sol";
import "../health/HealthChecker.sol";

/**
 * @notice An implementation of an On Ramp, which enables just a single token to be
 * used in the protocol.
 */
contract SingleTokenOnRamp is OnRampInterface, TypeAndVersionInterface, HealthChecker {
  using TokenLimits for TokenLimits.TokenBucket;

  // Chain ID of the destination chain. This is sent in the request to the DON.
  uint256 public immutable DESTINATION_CHAIN_ID;
  // Address of the token on the destination chain. This is sent in the request to the DON.
  IERC20 public immutable DESTINATION_TOKEN;

  // Chain ID of the source chain (where this contract is deployed)
  uint256 public immutable CHAIN_ID;
  // Token pool responsible for managing the TOKEN.
  PoolInterface public immutable POOL;
  // Token that this ramp enables to be sent using the protocol.
  IERC20 public immutable TOKEN;

  // Whether the allowlist is enabled
  bool private s_allowlistEnabled;
  // Addresses that are allowed to send messages
  mapping(address => bool) private s_allowed;
  // List of allowed addresses
  address[] private s_allowList;
  // Simple incremental nonce.
  uint256 private s_sequenceNumber;
  // Token bucket for token rate limiting
  TokenLimits.TokenBucket private s_tokenBucket;

  constructor(
    uint256 sourceChainId,
    IERC20 sourceToken,
    PoolInterface sourcePool,
    uint256 destinationChainId,
    IERC20 destinationToken,
    address[] memory allowlist,
    bool enableAllowlist,
    uint256 tokenBucketRate,
    uint256 tokenBucketCapacity,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal
  ) HealthChecker(afn, maxTimeWithoutAFNSignal) {
    if (sourcePool.getToken() != sourceToken) revert TokenMismatch();
    CHAIN_ID = sourceChainId;
    TOKEN = sourceToken;
    POOL = sourcePool;
    DESTINATION_CHAIN_ID = destinationChainId;
    DESTINATION_TOKEN = destinationToken;
    s_sequenceNumber = 1;
    s_allowlistEnabled = enableAllowlist;
    s_allowList = allowlist;
    for (uint256 i = 0; i < allowlist.length; i++) {
      s_allowed[allowlist[i]] = true;
    }
    s_tokenBucket = TokenLimits.constructTokenBucket(tokenBucketRate, tokenBucketCapacity, true);
  }

  /**
   * @notice Send a message to the remote chain
   * @dev tokens must be of length 1 and be the token allowed by this contract
   * @dev amounts must also be of length 1, be greater than zero, and approve() must have already
   * been called on the token using the POOL address as the spender.
   * @dev if the contract is paused, this function will revert.
   * @param payload Message struct to send
   */
  function requestCrossChainSend(CCIP.MessagePayload memory payload)
    external
    override
    whenNotPaused
    whenHealthy
    returns (uint256)
  {
    address sender = msg.sender;
    if (s_allowlistEnabled && !s_allowed[sender]) revert SenderNotAllowed(sender);
    // Check that inputs are correct
    if (payload.tokens.length != 1 || payload.amounts.length != 1) revert UnsupportedNumberOfTokens();
    if (payload.tokens[0] != TOKEN) revert UnsupportedToken(TOKEN, payload.tokens[0]);
    // This step will be a mapping filled with a loop in future when more than one token is suported.
    IERC20[] memory mappedRemoteTokens = new IERC20[](1);
    mappedRemoteTokens[0] = DESTINATION_TOKEN;
    payload.tokens = mappedRemoteTokens;
    // Check that sending these tokens falls within the bucket limits
    if (!s_tokenBucket.remove(payload.amounts[0])) revert ExceedsTokenLimit(s_tokenBucket.tokens, payload.amounts[0]);
    // Store in pool
    POOL.lockOrBurn(sender, payload.amounts[0]);
    // Emit message request
    CCIP.Message memory message = CCIP.Message({
      sequenceNumber: s_sequenceNumber,
      sourceChainId: CHAIN_ID,
      destinationChainId: DESTINATION_CHAIN_ID,
      sender: sender,
      payload: payload
    });
    emit CrossChainSendRequested(message);
    s_sequenceNumber++;
    return message.sequenceNumber;
  }

  function setAllowlistEnabled(bool enabled) external onlyOwner {
    s_allowlistEnabled = enabled;
    emit AllowlistEnabledSet(enabled);
  }

  function getAllowlistEnabled() external view returns (bool) {
    return s_allowlistEnabled;
  }

  function setAllowlist(address[] calldata allowlist) external onlyOwner {
    // Remove existing allowlist
    address[] memory existingList = s_allowList;
    for (uint256 i = 0; i < existingList.length; i++) {
      s_allowed[existingList[i]] = false;
    }

    // Set the new allowlist
    s_allowList = allowlist;
    for (uint256 i = 0; i < allowlist.length; i++) {
      s_allowed[allowlist[i]] = true;
    }
    emit AllowlistSet(allowlist);
  }

  function getAllowlist() external view returns (address[] memory) {
    return s_allowList;
  }

  function configureTokenBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) external onlyOwner {
    s_tokenBucket = TokenLimits.constructTokenBucket(rate, capacity, full);
    emit NewTokenBucketConstructed(rate, capacity, full);
  }

  function getTokenBucket() external view returns (TokenLimits.TokenBucket memory) {
    return s_tokenBucket;
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "SingleTokenOnRamp 1.1.0";
  }
}
                                                                                                                                                                                           contracts/src/v0.8/ccip/ramps/MessageExecutor.sol                                                   000644  000765  000024  00000003202 14165346401 022702  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../interfaces/OffRampInterface.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../ocr/OCR2Base.sol";
import "../utils/CCIP.sol";

/**
 * @notice MessageExecutor enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract MessageExecutor is TypeAndVersionInterface, OCR2Base {
  struct ExecutableMessage {
    bytes32[] proof;
    CCIP.Message message;
    uint256 index;
  }

  OffRampInterface public immutable s_offRamp;

  constructor(OffRampInterface offRamp) OCR2Base(true) {
    s_offRamp = offRamp;
  }

  /**
   * @notice Entry point for execution, called by the OCR network
   * @dev Expects an encoded array of ExectableMessage tuples.
   */
  function _report(
    bytes32, /*configDigest*/
    uint40, /*epochAndRound*/
    bytes memory report
  ) internal override {
    ExecutableMessage[] memory executableMessages = abi.decode(report, (ExecutableMessage[]));
    for (uint256 i = 0; i < executableMessages.length; i++) {
      ExecutableMessage memory em = executableMessages[i];
      s_offRamp.executeTransaction(em.proof, em.message, em.index);
    }
  }

  function _beforeSetConfig(uint8 _threshold, bytes memory _onchainConfig) internal override {
    // TODO
  }

  function _afterSetConfig(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {
    // TODO
  }

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {
    // TODO
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "MessageExecutor 1.0.0";
  }
}
                                                                                                                                                                                                                                                                                                                                                                                              contracts/src/v0.8/ccip/applications/EOASingleTokenSender.sol                                       000644  000765  000024  00000005626 14165346401 025067  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../ramps/SingleTokenOnRamp.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../utils/CCIP.sol";
import "../../vendor/SafeERC20.sol";

/**
 * @notice This contract enables EOAs to send a single asset across to the chain
 * represented by the On Ramp. Consider this an "Application Layer" contract that utilise the
 * underlying protocol.
 */
contract EOASingleTokenSender is TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  // On ramp contract responsible for interacting with the DON.
  SingleTokenOnRamp public immutable ON_RAMP;
  // Corresponding contract on the destination chain responsible for receiving the message
  // and enabling the EOA on the destination chain to access the tokens that are sent.
  // For this scenario, it would be the address of a deployed EOASingleTokenReceiver.
  address public immutable DESTINATION_CONTRACT;

  error InvalidDestinationAddress(address invalidAddress);

  constructor(SingleTokenOnRamp onRamp, address destinationContract) {
    ON_RAMP = onRamp;
    DESTINATION_CONTRACT = destinationContract;
  }

  /**
   * @notice Send tokens to the destination chain.
   * @dev msg.sender must first call TOKEN.approve for this contract to spend the tokens.
   */
  function sendTokens(
    address destinationAddress,
    uint256 amount,
    address executor
  ) external returns (uint256 sequenceNumber) {
    if (destinationAddress == address(0)) revert InvalidDestinationAddress(destinationAddress);
    bytes memory options;
    // Set tokens using the ramp token
    IERC20[] memory tokens = new IERC20[](1);
    tokens[0] = ON_RAMP.TOKEN();
    // Set the amounts using the amount parameter
    uint256[] memory amounts = new uint256[](1);
    amounts[0] = amount;
    address originalSender = msg.sender;
    // Init the MessagePayload struct
    // `payload.data` format:
    //  - EOA sender address
    //  - EOA destination address
    CCIP.MessagePayload memory payload = CCIP.MessagePayload({
      receiver: DESTINATION_CONTRACT,
      data: abi.encode(originalSender, destinationAddress),
      tokens: tokens,
      amounts: amounts,
      executor: executor,
      options: options
    });
    tokens[0].safeTransferFrom(originalSender, address(this), amount);
    tokens[0].approve(address(ON_RAMP.POOL()), amount);
    sequenceNumber = ON_RAMP.requestCrossChainSend(payload);
  }

  /**
   * @notice Get the details of the ramp. This includes the destination chain details
   */
  function rampDetails()
    external
    view
    returns (
      IERC20 token,
      uint256 destinationChainId,
      IERC20 destinationChainToken
    )
  {
    token = ON_RAMP.TOKEN();
    destinationChainId = ON_RAMP.DESTINATION_CHAIN_ID();
    destinationChainToken = ON_RAMP.DESTINATION_TOKEN();
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "EOASingleTokenSender 1.0.0";
  }
}
                                                                                                          contracts/src/v0.8/ccip/applications/EOASingleTokenReceiver.sol                                     000644  000765  000024  00000002620 14165346401 025402  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../interfaces/CrossChainMessageReceiverInterface.sol";
import "../../vendor/SafeERC20.sol";
import "../ramps/SingleTokenOffRamp.sol";
import "../../interfaces/TypeAndVersionInterface.sol";

/**
 * @notice Appliation contract for receiving messages from the OffRamp on behalf of an EOA
 */
contract EOASingleTokenReceiver is CrossChainMessageReceiverInterface, TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  SingleTokenOffRamp public immutable OFF_RAMP;

  error InvalidDeliverer(address deliverer);

  constructor(SingleTokenOffRamp offRamp) {
    OFF_RAMP = offRamp;
  }

  /**
   * @notice Called by the OffRamp, this function receives a message and forwards
   * the tokens sent with it to the designated EOA
   * @param message CCIP Message
   */
  function receiveMessage(CCIP.Message calldata message) external override {
    if (msg.sender != address(OFF_RAMP)) revert InvalidDeliverer(msg.sender);
    (
      ,
      /* address originalSender */
      address destinationAddress
    ) = abi.decode(message.payload.data, (address, address));
    if (destinationAddress != address(0) && message.payload.amounts[0] != 0) {
      OFF_RAMP.TOKEN().transfer(destinationAddress, message.payload.amounts[0]);
    }
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "EOASingleTokenReceiver 1.1.0";
  }
}
                                                                                                                contracts/src/v0.8/ccip/interfaces/CrossChainMessageReceiverInterface.sol                           000644  000765  000024  00000000655 14165346401 027460  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../utils/CCIP.sol";

/**
 * @notice Application contracts that intend to receive messages from
 * the OffRamp should implement this interface.
 */
interface CrossChainMessageReceiverInterface {
  /**
   * @notice Called by the OffRamp to deliver a message
   * @param message CCIP Message
   */
  function receiveMessage(CCIP.Message calldata message) external;
}
                                                                                   contracts/src/v0.8/ccip/interfaces/OnRampInterface.sol                                              000644  000765  000024  00000001623 14165346401 023622  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../utils/CCIP.sol";
import "../interfaces/PoolInterface.sol";
import "../interfaces/AFNInterface.sol";

interface OnRampInterface {
  error TokenMismatch();
  error UnsupportedNumberOfTokens();
  error UnsupportedToken(IERC20 expected, IERC20 given);
  error ExceedsTokenLimit(uint256 currentLimit, uint256 requested);
  error SenderNotAllowed(address sender);

  event CrossChainSendRequested(CCIP.Message message);
  event AllowlistEnabledSet(bool enabled);
  event AllowlistSet(address[] allowlist);
  event NewTokenBucketConstructed(uint256 rate, uint256 capacity, bool full);

  /**
   * @notice Request a message to be sent to the destination chain
   * @param payload The message payload
   * @return The sequence number of the message
   */
  function requestCrossChainSend(CCIP.MessagePayload calldata payload) external returns (uint256);
}
                                                                                                             contracts/src/v0.8/ccip/interfaces/AFNInterface.sol                                                 000644  000765  000024  00000002120 14165346401 023023  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

interface AFNInterface {
  struct Heartbeat {
    uint256 round;
    uint256 timestamp;
    uint256 committeeVersion;
  }

  event GoodVote(address voter, uint256 round);
  event BadVote(address voter, uint256 round);
  event AFNHeartbeat(Heartbeat heartbeat);
  event AFNBadSignal(uint256 timestamp);
  event RecoveredFromBadSignal();
  event ConfigSet(address[] parties, uint256[] weights, uint256 goodQuorum, uint256 badQuorum);

  error IncorrectRound(uint256 expected, uint256 received);
  error InvalidVoter(address voter);
  error AlreadyVoted();
  error InvalidConfig();
  error InvalidWeight();
  error MustRecoverFromBadSignal();
  error RecoveryNotNecessary();

  function hasBadSignal() external returns (bool);

  function getLastHeartbeat() external returns (Heartbeat memory);

  function voteGood(uint256 round) external;

  function voteBad() external;

  function recover() external;

  function setConfig(
    address[] memory parties,
    uint256[] memory weights,
    uint256 goodQuorum,
    uint256 badQuorum
  ) external;
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                contracts/src/v0.8/ccip/interfaces/PoolInterface.sol                                                000644  000765  000024  00000001521 14165346401 023334  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "../../vendor/IERC20.sol";

// Shared public interface for multiple pool types.
// Each pool type handles a different child token model (lock/unlock, mint/burn.)
interface PoolInterface {
  event Locked(address indexed sender, address indexed depositor, uint256 amount);
  event Burnt(address indexed sender, address indexed depositor, uint256 amount);
  event Released(address indexed sender, address indexed recipient, uint256 amount);
  event Minted(address indexed sender, address indexed recipient, uint256 amount);

  function lockOrBurn(address depositor, uint256 amount) external;

  function releaseOrMint(address recipient, uint256 amount) external;

  function getToken() external view returns (IERC20 pool);

  function pause() external;

  function unpause() external;
}
                                                                                                                                                                               contracts/src/v0.8/ccip/interfaces/OffRampInterface.sol                                             000644  000765  000024  00000002662 14165346401 023764  0                                                                                                    ustar 00kostis                          staff                           000000  000000                                                                                                                                                                         // SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../interfaces/PoolInterface.sol";
import "../interfaces/CrossChainMessageReceiverInterface.sol";
import "../utils/CCIP.sol";

interface OffRampInterface {
  error RelayReportError();
  error SequenceError(uint256 lastMaxSequenceNumber, uint256 newMinSequenceNumber);
  error MerkleProofError(bytes32[] proof, CCIP.Message message, uint256 index);
  error TokenMismatch();
  error UnsupportedNumberOfTokens();
  error UnsupportedToken(IERC20 token);
  error ExceedsTokenLimit(uint256 currentLimit, uint256 requested);
  error AlreadyExecuted(uint256 sequenceNumber);
  error InvalidExecutor(uint256 sequenceNumber);
  error ExecutionError(uint256 sequenceNumber, bytes reason);
  error ExecutionDelayError();
  error InvalidReceiver(address receiver);
  error InvalidSourceChain(uint256 sourceChainId);

  event ReportAccepted(CCIP.RelayReport report);
  event CrossChainMessageExecuted(uint256 indexed sequenceNumber);
  event ExecutionDelaySecondsSet(uint256 delay);
  event NewTokenBucketConstructed(uint256 rate, uint256 capacity, bool full);

  /**
   * @notice Execute the delivery of a message by using its merkle proof
   * @param proof Merkle proof
   * @param message Original message object
   * @param index Index of the message in the original tree
   */
  function executeTransaction(
    bytes32[] memory proof,
    CCIP.Message memory message,
    uint256 index
  ) external;
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              