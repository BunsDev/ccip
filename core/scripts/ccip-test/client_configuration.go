package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_subscription_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/governance_dapp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ping_pong_demo"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/receiver_dapp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/subscription_sender_dapp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/dione"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/core/utils"
)

func (client *CCIPClient) wip(t *testing.T, sourceClient *rhea.EvmDeploymentConfig, destClient *rhea.EvmDeploymentConfig) {

}

func (client *CCIPClient) startPingPong(t *testing.T) {
	tx, err := client.Source.PingPongDapp.StartPingPong(client.Source.Owner)
	require.NoError(t, err)
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) setPingPongPaused(t *testing.T, paused bool) {
	tx, err := client.Source.PingPongDapp.SetPaused(client.Source.Owner, paused)
	require.NoError(t, err)
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) fundPingPong(t *testing.T) {
	fundingAmount := big.NewInt(1e18)
	client.Dest.ApproveLinkFrom(t, client.Dest.Owner, client.Dest.OffRampRouter.Address(), fundingAmount)
	tx, err := client.Dest.OffRampRouter.FundSubscription(client.Dest.Owner, client.Dest.PingPongDapp.Address(), fundingAmount)
	require.NoError(t, err)
	shared.WaitForMined(t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	client.Dest.logger.Infof(fmt.Sprintf("Ping pong funded %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), tx.Hash())))
}

type Client struct {
	Owner            *bind.TransactOpts
	Users            []*bind.TransactOpts
	Client           *ethclient.Client
	ChainId          *big.Int
	LinkToken        *link_token_interface.LinkToken
	LinkTokenAddress common.Address
	BridgeTokens     []common.Address
	TokenPools       []*native_token_pool.NativeTokenPool
	TokenPrices      []*big.Int
	GovernanceDapp   *governance_dapp.GovernanceDapp
	PingPongDapp     *ping_pong_demo.PingPongDemo
	Afn              *afn_contract.AFNContract
	logger           logger.Logger
	t                *testing.T
}

type SourceClient struct {
	Client
	OnRamp       *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRamp
	OnRampRouter *evm_2_any_subscription_onramp_router.EVM2AnySubscriptionOnRampRouter
	SenderDapp   *subscription_sender_dapp.SubscriptionSenderDapp
}

func NewSourceClient(t *testing.T, config rhea.EvmDeploymentConfig) SourceClient {
	client := rhea.GetClient(t, config.ChainConfig.EthUrl)
	LinkToken, err := link_token_interface.NewLinkToken(config.ChainConfig.LinkToken, client)
	require.NoError(t, err)
	var tokenPools []*native_token_pool.NativeTokenPool
	for _, poolAddress := range config.ChainConfig.TokenPools {
		tokenPool, err2 := native_token_pool.NewNativeTokenPool(poolAddress, client)
		require.NoError(t, err2)
		tokenPools = append(tokenPools, tokenPool)
	}

	afn, err := afn_contract.NewAFNContract(config.ChainConfig.Afn, client)
	require.NoError(t, err)
	onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(config.LaneConfig.OnRamp, client)
	require.NoError(t, err)
	senderDapp, err := subscription_sender_dapp.NewSubscriptionSenderDapp(config.LaneConfig.TokenSender, client)
	require.NoError(t, err)
	onRampRouter, err := evm_2_any_subscription_onramp_router.NewEVM2AnySubscriptionOnRampRouter(config.ChainConfig.OnRampRouter, client)
	require.NoError(t, err)
	governanceDapp, err := governance_dapp.NewGovernanceDapp(config.LaneConfig.GovernanceDapp, client)
	require.NoError(t, err)
	pingPongDapp, err := ping_pong_demo.NewPingPongDemo(config.LaneConfig.PingPongDapp, client)
	require.NoError(t, err)

	return SourceClient{
		Client: Client{
			Client:           client,
			ChainId:          config.ChainConfig.ChainId,
			LinkTokenAddress: config.ChainConfig.LinkToken,
			LinkToken:        LinkToken,
			Afn:              afn,
			BridgeTokens:     config.ChainConfig.BridgeTokens,
			TokenPools:       tokenPools,
			TokenPrices:      config.ChainConfig.TokenPrices,
			GovernanceDapp:   governanceDapp,
			PingPongDapp:     pingPongDapp,
			logger:           logger.TestLogger(t).Named(helpers.ChainName(config.ChainConfig.ChainId.Int64())),
			t:                t,
		},
		OnRamp:       onRamp,
		OnRampRouter: onRampRouter,
		SenderDapp:   senderDapp,
	}
}

type DestClient struct {
	Client
	BlobVerifier    *blob_verifier.BlobVerifier
	MessageReceiver *simple_message_receiver.SimpleMessageReceiver
	ReceiverDapp    *receiver_dapp.ReceiverDapp
	OffRamp         *any_2_evm_subscription_offramp.EVM2EVMSubscriptionOffRamp
	OffRampRouter   *any_2_evm_subscription_offramp_router.Any2EVMSubscriptionOffRampRouter
}

func NewDestinationClient(t *testing.T, config rhea.EvmDeploymentConfig) DestClient {
	client := rhea.GetClient(t, config.ChainConfig.EthUrl)
	LinkToken, err := link_token_interface.NewLinkToken(config.ChainConfig.LinkToken, client)
	require.NoError(t, err)

	var tokenPools []*native_token_pool.NativeTokenPool
	for _, poolAddress := range config.ChainConfig.TokenPools {
		tokenPool, err2 := native_token_pool.NewNativeTokenPool(poolAddress, client)
		require.NoError(t, err2)
		tokenPools = append(tokenPools, tokenPool)
	}

	afn, err := afn_contract.NewAFNContract(config.ChainConfig.Afn, client)
	require.NoError(t, err)
	blobVerifier, err := blob_verifier.NewBlobVerifier(config.LaneConfig.BlobVerifier, client)
	require.NoError(t, err)
	offRamp, err := any_2_evm_subscription_offramp.NewEVM2EVMSubscriptionOffRamp(config.LaneConfig.OffRamp, client)
	require.NoError(t, err)
	messageReceiver, err := simple_message_receiver.NewSimpleMessageReceiver(config.LaneConfig.MessageReceiver, client)
	require.NoError(t, err)
	receiverDapp, err := receiver_dapp.NewReceiverDapp(config.LaneConfig.ReceiverDapp, client)
	require.NoError(t, err)
	offRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(config.ChainConfig.OffRampRouter, client)
	require.NoError(t, err)
	governanceDapp, err := governance_dapp.NewGovernanceDapp(config.LaneConfig.GovernanceDapp, client)
	require.NoError(t, err)
	pingPongDapp, err := ping_pong_demo.NewPingPongDemo(config.LaneConfig.PingPongDapp, client)
	require.NoError(t, err)

	return DestClient{
		Client: Client{
			Client:           client,
			ChainId:          config.ChainConfig.ChainId,
			LinkTokenAddress: config.ChainConfig.LinkToken,
			LinkToken:        LinkToken,
			BridgeTokens:     config.ChainConfig.BridgeTokens,
			TokenPools:       tokenPools,
			TokenPrices:      config.ChainConfig.TokenPrices,
			GovernanceDapp:   governanceDapp,
			PingPongDapp:     pingPongDapp,
			Afn:              afn,
			logger:           logger.TestLogger(t).Named(helpers.ChainName(config.ChainConfig.ChainId.Int64())),
			t:                t,
		},
		BlobVerifier:    blobVerifier,
		OffRampRouter:   offRampRouter,
		MessageReceiver: messageReceiver,
		ReceiverDapp:    receiverDapp,
		OffRamp:         offRamp,
	}
}

// CCIPClient contains a source chain and destination chain client and implements many methods
// that are useful for testing CCIP functionality on chain.
type CCIPClient struct {
	Source SourceClient
	Dest   DestClient
}

// NewCcipClient returns a new CCIPClient with initialised source and destination clients.
func NewCcipClient(t *testing.T, sourceConfig rhea.EvmDeploymentConfig, destConfig rhea.EvmDeploymentConfig, ownerKey string, seedKey string) CCIPClient {
	source := NewSourceClient(t, sourceConfig)
	source.SetOwnerAndUsers(t, ownerKey, seedKey, sourceConfig.ChainConfig.GasSettings)
	dest := NewDestinationClient(t, destConfig)
	dest.SetOwnerAndUsers(t, ownerKey, seedKey, destConfig.ChainConfig.GasSettings)

	return CCIPClient{
		Source: source,
		Dest:   dest,
	}
}

func GetSetupChain(t *testing.T, ownerPrivateKey string, chain rhea.EvmDeploymentConfig) *rhea.EvmDeploymentConfig {
	chain.SetupChain(t, ownerPrivateKey)
	return &chain
}

// SetOwnerAndUsers sets the owner and 10 users on a given client. It also set the proper
// gas parameters on these users.
func (client *Client) SetOwnerAndUsers(t *testing.T, ownerPrivateKey string, seedKey string, gasSettings rhea.EVMGasSettings) {
	client.Owner = rhea.GetOwner(t, ownerPrivateKey, client.ChainId, gasSettings)

	var users []*bind.TransactOpts
	seedKeyWithoutFirstChar := seedKey[1:]
	fmt.Println("--- Addresses of the seed key")
	for i := 0; i <= 9; i++ {
		_, err := hex.DecodeString(strconv.Itoa(i) + seedKeyWithoutFirstChar)
		require.NoError(t, err)
		key, err := crypto.HexToECDSA(strconv.Itoa(i) + seedKeyWithoutFirstChar)
		require.NoError(t, err)
		user, err := bind.NewKeyedTransactorWithChainID(key, client.ChainId)
		require.NoError(t, err)
		rhea.SetGasFees(user, gasSettings)
		users = append(users, user)
		fmt.Println(user.From.Hex())
	}
	fmt.Println("---")

	client.Users = users
}

func (client *Client) TypeAndVersion(addr common.Address) (ccip.ContractType, semver.Version, error) {
	return ccip.TypeAndVersion(addr, client.Client)
}

func (client *Client) ApproveLinkFrom(t *testing.T, user *bind.TransactOpts, approvedFor common.Address, amount *big.Int) {
	client.logger.Warnf("Approving %d link for %s", amount.Int64(), approvedFor.Hex())
	tx, err := client.LinkToken.Approve(user, approvedFor, amount)
	require.NoError(t, err)

	shared.WaitForMined(client.t, client.logger, client.Client, tx.Hash(), true)
	client.logger.Warnf("Link approved %s", helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash()))
}

func (client *Client) ApproveLink(t *testing.T, approvedFor common.Address, amount *big.Int) {
	client.ApproveLinkFrom(t, client.Owner, approvedFor, amount)
}

func (client *CCIPClient) ChangeGovernanceParameters(t *testing.T) {
	feeConfig := governance_dapp.GovernanceDappFeeConfig{
		FeeAmount:           big.NewInt(10),
		SubscriptionManager: client.Source.Owner.From,
		ChangedAtBlock:      big.NewInt(0),
	}
	DestBlockNum := GetCurrentBlockNumber(client.Dest.Client.Client)
	sourceBlockNum := GetCurrentBlockNumber(client.Source.Client.Client)

	tx, err := client.Source.GovernanceDapp.VoteForNewFeeConfig(client.Source.Owner, feeConfig)
	require.NoError(t, err)
	sendRequest := WaitForCrossChainSendRequest(client.Source, sourceBlockNum, tx.Hash())
	client.WaitForRelay(t, DestBlockNum)
	client.WaitForExecution(t, DestBlockNum, sendRequest.Message.SequenceNumber)
}

func (client *CCIPClient) SendMessage(t *testing.T) {
	DestBlockNum := GetCurrentBlockNumber(client.Dest.Client.Client)

	// ABI encoded message
	bts, err := hex.DecodeString("00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000005626c616e6b000000000000000000000000000000000000000000000000000000")
	require.NoError(t, err)

	msg := evm_2_any_subscription_onramp_router.CCIPEVM2AnySubscriptionMessage{
		Receiver: testhelpers.MustEncodeAddress(t, client.Dest.MessageReceiver.Address()),
		Data:     bts,
		Tokens:   []common.Address{client.Source.LinkTokenAddress},
		Amounts:  []*big.Int{big.NewInt(1)},
		GasLimit: big.NewInt(3e5),
	}

	tx, err := client.Source.OnRampRouter.CcipSend(client.Source.Owner, client.Dest.ChainId, msg)
	require.NoError(t, err)
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	client.WaitForRelay(t, DestBlockNum)
}

func (client *CCIPClient) DonExecutionHappyPath(t *testing.T) {
	client.Source.logger.Infof("Starting cross chain tx with DON execution")

	tokenAmount := big.NewInt(500)
	client.Source.ApproveLink(t, client.Source.OnRampRouter.Address(), tokenAmount)

	DestBlockNum := GetCurrentBlockNumber(client.Dest.Client.Client)
	crossChainRequest := client.SendToOnrampWithExecution(t, client.Source, client.Source.Owner, client.Dest.ReceiverDapp.Address(), tokenAmount)
	client.Source.logger.Infof("Don executed tx submitted with sequence number: %d", crossChainRequest.Message.SequenceNumber)

	client.WaitForRelay(t, DestBlockNum)
	client.WaitForExecution(t, DestBlockNum, crossChainRequest.Message.SequenceNumber)
}

func (client *CCIPClient) WaitForRelay(t *testing.T, DestBlockNum uint64) {
	client.Dest.logger.Infof("Waiting for relay")

	relayEvent := make(chan *blob_verifier.BlobVerifierReportAccepted)
	sub, err := client.Dest.BlobVerifier.WatchReportAccepted(
		&bind.WatchOpts{
			Context: context.Background(),
			Start:   &DestBlockNum,
		},
		relayEvent,
	)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	select {
	case event := <-relayEvent:
		client.Dest.logger.Infof("Relay in tx %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), event.Raw.TxHash))
		return
	case err = <-sub.Err():
		panic(err)
	}
}

func (client *CCIPClient) WaitForExecution(t *testing.T, DestBlockNum uint64, sequenceNumber uint64) {
	client.Dest.logger.Infof("Waiting for execution...")

	events := make(chan *any_2_evm_subscription_offramp.EVM2EVMSubscriptionOffRampExecutionStateChanged)
	sub, err := client.Dest.OffRamp.WatchExecutionStateChanged(
		&bind.WatchOpts{
			Context: context.Background(),
			Start:   &DestBlockNum,
		},
		events,
		[]uint64{sequenceNumber})
	require.NoError(t, err)
	defer sub.Unsubscribe()

	select {
	case event := <-events:
		client.Dest.logger.Infof("Execution in tx %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), event.Raw.TxHash))
		return
	case err = <-sub.Err():
		panic(err)
	}
}

func (client *CCIPClient) ExecuteManually(seqNr uint64) error {
	// Find the seq num
	// Find the corresponding relay report
	end := uint64(11436244)
	reportIterator, err := client.Dest.BlobVerifier.FilterReportAccepted(&bind.FilterOpts{
		Start: end - 10000,
		End:   &end,
	})
	if err != nil {
		return err
	}
	var onRampIdx int
	var report *blob_verifier.CCIPRelayReport
	for reportIterator.Next() {
		for i, onRamp := range reportIterator.Event.Report.OnRamps {
			if onRamp == client.Source.OnRamp.Address() {
				if reportIterator.Event.Report.Intervals[i].Min <= seqNr && reportIterator.Event.Report.Intervals[i].Max >= seqNr {
					onRampIdx = i
					report = &reportIterator.Event.Report
					fmt.Println("Found root")
					break
				}
			}
		}
	}
	reportIterator.Close()
	if report == nil {
		return errors.New("unable to find seq num")
	}
	ctx := hasher.NewKeccakCtx()
	leafHasher := ccip.NewSubscriptionLeafHasher(client.Source.ChainId, client.Dest.ChainId, client.Source.OnRamp.Address(), ctx)
	// Get all seqNrs in that range.
	end = uint64(7651526)
	sendRequestedIterator, err := client.Source.OnRamp.FilterCCIPSendRequested(&bind.FilterOpts{
		Start: end - 10000,
		End:   &end,
	})
	if err != nil {
		return err
	}
	var leaves [][32]byte
	var curr, prove int
	var originalMsg []byte
	for sendRequestedIterator.Next() {
		// Assume in order?
		if sendRequestedIterator.Event.Message.SequenceNumber <= report.Intervals[onRampIdx].Max && sendRequestedIterator.Event.Message.SequenceNumber >= report.Intervals[onRampIdx].Min {
			fmt.Println("Found seq num", sendRequestedIterator.Event.Message.SequenceNumber, report.Intervals[onRampIdx])
			hash, err2 := leafHasher.HashLeaf(sendRequestedIterator.Event.Raw)
			if err2 != nil {
				return err2
			}
			leaves = append(leaves, hash)
			if sendRequestedIterator.Event.Message.SequenceNumber == seqNr {
				fmt.Printf("Found proving %d %+v\n", curr, sendRequestedIterator.Event.Message)
				originalMsg = sendRequestedIterator.Event.Raw.Data
				prove = curr
			}
			curr++
		}
	}
	sendRequestedIterator.Close()
	if originalMsg == nil {
		return errors.New("unable to find")
	}
	tree, err := merklemulti.NewTree(ctx, leaves)
	if err != nil {
		return err
	}
	innerProof := tree.Prove([]int{prove})
	if tree.Root() != report.MerkleRoots[onRampIdx] {
		return errors.New("inner root doesn't match")
	}
	outerTree, err := merklemulti.NewTree(ctx, report.MerkleRoots)
	if err != nil {
		return err
	}
	if outerTree.Root() != report.RootOfRoots {
		return errors.New("outer root doesn't match")
	}
	outerProof := outerTree.Prove([]int{onRampIdx})
	executionReport := any_2_evm_subscription_offramp.CCIPExecutionReport{
		SequenceNumbers:          []uint64{seqNr},
		TokenPerFeeCoinAddresses: []common.Address{client.Dest.LinkTokenAddress},
		TokenPerFeeCoin:          []*big.Int{big.NewInt(1)},
		EncodedMessages:          [][]byte{originalMsg},
		InnerProofs:              innerProof.Hashes,
		InnerProofFlagBits:       ccip.ProofFlagsToBits(innerProof.SourceFlags),
		OuterProofs:              outerProof.Hashes,
		OuterProofFlagBits:       ccip.ProofFlagsToBits(outerProof.SourceFlags),
	}
	tx, err := client.Dest.OffRamp.Execute(client.Dest.Owner, executionReport, true)
	if err != nil {
		fmt.Printf("%+v err %v\n", executionReport, err)
		return err
	}
	fmt.Println(client.Dest.Owner.From, tx.Hash(), err)
	return nil
}

//func (client CCIPClient) ExternalExecutionHappyPath(t *testing.T) {
//	ctx := context.Background()
//	offrampBlockNumber := GetCurrentBlockNumber(client.Dest.Client.Client)
//	onrampBlockNumber := GetCurrentBlockNumber(client.Source.Client.Client)
//
//	amount, _ := new(big.Int).SetString("10", 10)
//	client.Source.ApproveLink(t, client.Source.OnRamRouter.Address(), amount)
//
//	onrampRequest := client.SendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
//	sequenceNumber := onrampRequest.Message.SequenceNumber
//
//	// Gets the report that our transaction is included in
//	client.Dest.logger.Info("Getting report")
//	report, err := client.GetReportForSequenceNumber(ctx, sequenceNumber, offrampBlockNumber)
//	require.NoError(t, err)
//
//	// Get all requests included in the given report
//	client.Dest.logger.Info("Getting recent cross chain requests")
//	requests := client.GetCrossChainSendRequestsForRange(ctx, t, report, onrampBlockNumber)
//
//	// Generate the proof
//	client.Dest.logger.Info("Generating proof")
//	proof := client.ValidateMerkleRoot(t, onrampRequest, requests, report)
//
//	// Execute the transaction on the offramp
//	client.Dest.logger.Info("Executing offramp TX")
//	tx, err := client.ExecuteOffRampTransaction(t, proof, onrampRequest.Raw.Data)
//	require.NoError(t, err)
//
//	WaitForMined(t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
//	client.Dest.logger.Infof("Cross chain tx sent %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), tx.Hash()))
//}

func (client *CCIPClient) CrossChainSendPausedOnrampShouldFail(t *testing.T) {
	client.PauseOnramp()
	amount := big.NewInt(100)
	client.Source.ApproveLink(t, client.Source.SenderDapp.Address(), amount)
	client.Source.Owner.GasLimit = 1e6
	tx, err := client.Source.SenderDapp.SendMessage(client.Source.Owner,
		subscription_sender_dapp.CCIPEVM2AnySubscriptionMessage{
			Receiver: testhelpers.MustEncodeAddress(t, client.Dest.Owner.From),
			Tokens:   []common.Address{client.Source.LinkTokenAddress},
			Amounts:  []*big.Int{amount},
			GasLimit: big.NewInt(100_000),
		})
	require.NoError(t, err)
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), false)
}

//func (client CCIPClient) CrossChainSendPausedOfframpShouldFail(t *testing.T) {
//	client.PauseBlobVerifier()
//	ctx := context.Background()
//	offrampBlockNumber := GetCurrentBlockNumber(client.Dest.Client.Client)
//
//	amount, _ := new(big.Int).SetString("10", 10)
//	client.Source.ApproveLink(t, client.Source.SenderDapp.Address(), amount)
//	onrampRequest := client.SendToDappWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
//
//	client.Dest.logger.Info("Waiting for report...")
//	_, err := client.GetReportForSequenceNumber(ctx, onrampRequest.Message.SequenceNumber, offrampBlockNumber)
//	if err.Error() == "No report found within the given time" {
//		client.Dest.logger.Info("Success, no oracle report sent to paused offramp.")
//	} else {
//		panic("report found despite paused offramp")
//	}
//}

func (client *CCIPClient) NotEnoughFundsInBucketShouldFail(t *testing.T) {
	amount := big.NewInt(2e18) // 2 LINK, bucket size is 1 LINK
	client.Source.ApproveLink(t, client.Source.SenderDapp.Address(), amount)
	client.Source.Owner.GasLimit = 1e6
	tx, err := client.Source.SenderDapp.SendMessage(client.Source.Owner,
		subscription_sender_dapp.CCIPEVM2AnySubscriptionMessage{
			Receiver: testhelpers.MustEncodeAddress(t, client.Dest.Owner.From),
			Tokens:   []common.Address{client.Source.LinkTokenAddress},
			Amounts:  []*big.Int{amount},
			GasLimit: big.NewInt(100_000),
		})
	require.NoError(t, err)
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), false)
}

//func (client CCIPClient) ExternalExecutionSubmitOfframpTwiceShouldFail(t *testing.T) {
//	ctx := context.Background()
//	offrampBlockNumber := GetCurrentBlockNumber(client.Dest.Client.Client)
//	onrampBlockNumber := GetCurrentBlockNumber(client.Source.Client.Client)
//
//	amount, _ := new(big.Int).SetString("10", 10)
//	client.Source.ApproveLink(t, client.Source.SenderDapp.Address(), amount)
//
//	onrampRequest := client.SendToDappWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
//
//	// Gets the report that our transaction is included in
//	client.Dest.logger.Info("Getting report")
//	report, err := client.GetReportForSequenceNumber(ctx, onrampRequest.Message.SequenceNumber, offrampBlockNumber)
//	require.NoError(t, err)
//
//	// Get all requests included in the given report
//	client.Dest.logger.Info("Getting recent cross chain requests")
//	requests := client.GetCrossChainSendRequestsForRange(ctx, t, report, onrampBlockNumber)
//
//	// Generate the proof
//	client.Dest.logger.Info("Generating proof")
//	proof := client.ValidateMerkleRoot(t, onrampRequest, requests, report)
//
//	// Execute the transaction on the offramp
//	client.Dest.logger.Info("Executing first offramp TX - should succeed")
//	tx, err := client.ExecuteOffRampTransaction(t, proof, onrampRequest.Raw.Data)
//	require.NoError(t, err)
//	WaitForMined(t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
//
//	// Execute the transaction on the offramp
//	client.Dest.logger.Info("Executing second offramp TX - should fail")
//	client.Dest.Owner.GasLimit = 1e6
//	tx, err = client.ExecuteOffRampTransaction(t, proof, onrampRequest.Raw.Data)
//	require.NoError(t, err)
//	WaitForMined(t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), false)
//}

func (client *CCIPClient) SendDappTx(t *testing.T) {
	amount := big.NewInt(500)
	destBlockNumber := GetCurrentBlockNumber(client.Dest.Client.Client)

	client.Source.ApproveLink(t, client.Source.SenderDapp.Address(), amount)
	crossChainRequest := client.SendToDappWithExecution(t, client.Source, client.Source.Owner, client.Dest.Owner.From, amount)
	client.WaitForRelay(t, destBlockNumber)
	client.WaitForExecution(t, destBlockNumber, crossChainRequest.Message.SequenceNumber)
}

// ScalingAndBatching should scale so that we see batching on the nodes
func (client *CCIPClient) ScalingAndBatching(t *testing.T) {
	amount := big.NewInt(10)
	toAddress := common.HexToAddress("0x57359120D900fab8cE74edC2c9959b21660d3887")
	DestBlockNum := GetCurrentBlockNumber(client.Dest.Client.Client)
	var seqNum uint64

	var wg sync.WaitGroup
	for _, user := range client.Source.Users {
		wg.Add(1)
		go func(user *bind.TransactOpts) {
			defer wg.Done()
			client.Source.ApproveLinkFrom(t, user, client.Source.SenderDapp.Address(), amount)
			crossChainRequest := client.SendToDappWithExecution(t, client.Source, user, toAddress, amount)
			client.Source.logger.Info("Don executed tx submitted with sequence number: ", crossChainRequest.Message.SequenceNumber)
			seqNum = crossChainRequest.Message.SequenceNumber
		}(user)
	}
	wg.Wait()
	client.WaitForRelay(t, DestBlockNum)
	client.WaitForExecution(t, DestBlockNum, seqNum)
	client.Source.logger.Info("Sent 10 txs to onramp.")
}

//func (client CCIPClient) ExecuteOffRampTransaction(t *testing.T, proof merklemulti.Proof[[32]byte], encodedMessage []byte) (*types.Transaction, error) {
//	decodedMsg, err := ccip.DecodeCCIPMessage(encodedMessage)
//	require.NoError(t, err)
//	_, err = ccip.MakeTollCCIPMsgArgs().PackValues([]interface{}{*decodedMsg})
//	require.NoError(t, err)
//
//	client.Dest.logger.Infof("Cross chain message %+v", decodedMsg)
//
//	report := any_2_evm_toll_offramp.CCIPExecutionReport{
//		Messages:       []any_2_evm_toll_offramp.CCIPAny2EVMTollMessage{*decodedMsg},
//		Proofs:         proof.Hashes,
//		ProofFlagsBits: ccip.ProofFlagsToBits(proof.SourceFlags),
//	}
//
//	tx, err := client.Dest.BlobVerifier.ExecuteTransaction(client.Dest.Owner, report, false)
//	if err != nil {
//		reason, err2 := evmclient.ExtractRevertReasonFromRPCError(err)
//		require.NoError(t, err2)
//		client.Dest.logger.Errorf("Extracting revert reason \"%s\" err \"%s\"", reason, err)
//	}
//
//	return tx, errors.Wrap(err, "Executing offramp tx")
//}

//func (client CCIPClient) GetCrossChainSendRequestsForRange(
//	ctx context.Context,
//	t *testing.T,
//	report blob_verifier.CCIPRelayReport,
//	onrampBlockNumber uint64) []*evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested {
//	// Get the other transactions in the proof, we look 1000 blocks back for transaction
//	// should be fine? Needs fine-tuning after improved batching strategies are developed
//	// in milestone 4
//	reqsIterator, err := client.Source.OnRamp.FilterCCIPSendRequested(&bind.FilterOpts{
//		Context: ctx,
//		Start:   onrampBlockNumber - 1000,
//	})
//	require.NoError(t, err)
//
//	var requests []*evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested
//	var minFound = report.MaxSequenceNumber
//
//	for reqsIterator.Next() {
//		num := reqsIterator.Event.Message.SequenceNumber
//		if num < minFound {
//			minFound = num
//		}
//		if num >= report.MinSequenceNumber && num <= report.MaxSequenceNumber {
//			requests = append(requests, reqsIterator.Event)
//		}
//	}
//
//	// TODO: Even if this check passes, we may not have fetched all necessary requests if
//	// minFound == report.MinSequenceNumber
//	if minFound > report.MinSequenceNumber {
//		t.Log("Not all cross chain requests found in the last 1000 blocks")
//		t.FailNow()
//	}
//
//	return requests
//}

//// GetReportForSequenceNumber return the offramp.CCIPRelayReport for a given ccip requests sequence number.
//func (client CCIPClient) GetReportForSequenceNumber(ctx context.Context, sequenceNumber uint64, minBlockNumber uint64) (blob_verifier.CCIPRelayReport, error) {
//	client.Dest.logger.Infof("Looking for sequenceNumber %d", sequenceNumber)
//	report, err := client.Dest.OffRamp.GetLastReport(&bind.CallOpts{Context: ctx, Pending: false})
//	if err != nil {
//		return blob_verifier.CCIPRelayReport{}, err
//	}
//
//	client.Dest.logger.Infof("Last report found for range %d-%d", report.MinSequenceNumber, report.MaxSequenceNumber)
//	// our tx is in the latest report
//	if sequenceNumber >= report.MinSequenceNumber && sequenceNumber <= report.MaxSequenceNumber {
//		return report, nil
//	}
//	// report isn't out yet, it will be in a future report
//	if sequenceNumber > report.MaxSequenceNumber {
//		maxIterations := CrossChainTimout / RetryTiming
//		for i := 0; i < int(maxIterations); i++ {
//			report, err = client.Dest.BlobVerifier.GetLastReport(&bind.CallOpts{Context: ctx, Pending: false})
//			if err != nil {
//				return blob_verifier.CCIPRelayReport{}, err
//			}
//			client.Dest.logger.Infof("Last report found for range %d-%d", report.MinSequenceNumber, report.MaxSequenceNumber)
//			if sequenceNumber >= report.MinSequenceNumber && sequenceNumber <= report.MaxSequenceNumber {
//				return report, nil
//			}
//			time.Sleep(RetryTiming)
//		}
//		return blob_verifier.CCIPRelayReport{}, errors.New("No report found within the given timeout")
//	}
//
//	// it is in a past report, start looking at the earliest block number possible, the one
//	// before we started the entire transaction on the onramp.
//	reports, err := client.Dest.BlobVerifier.FilterReportAccepted(&bind.FilterOpts{
//		Start:   minBlockNumber,
//		End:     nil,
//		Context: ctx,
//	})
//	if err != nil {
//		return blob_verifier.CCIPRelayReport{}, err
//	}
//
//	for reports.Next() {
//		report = reports.Event.Report
//		if sequenceNumber >= report.MinSequenceNumber && sequenceNumber <= report.MaxSequenceNumber {
//			return report, nil
//		}
//	}
//
//	// Somehow the transaction was not included in any report within blocks produced after
//	// the transaction was initialized but the sequence number is lower than we are currently at
//	return blob_verifier.CCIPRelayReport{}, errors.New("No report found for given sequence number")
//}

func (client *CCIPClient) SetBlobVerifierConfig(t *testing.T) {
	config := blob_verifier.BlobVerifierInterfaceBlobVerifierConfig{
		OnRamps:          []common.Address{client.Source.OnRamp.Address()},
		MinSeqNrByOnRamp: []uint64{3},
	}
	tx, err := client.Dest.BlobVerifier.SetConfig(client.Dest.Owner, config)
	require.NoError(t, err)
	shared.WaitForMined(t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func GetCurrentBlockNumber(chain *ethclient.Client) uint64 {
	blockNumber, err := chain.BlockNumber(context.Background())
	helpers.PanicErr(err)
	return blockNumber
}

func (client *CCIPClient) ValidateMerkleRoot(
	t *testing.T,
	request *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested,
	reportRequests []*evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested,
	report blob_verifier.CCIPRelayReport,
) merklemulti.Proof[[32]byte] {
	mctx := hasher.NewKeccakCtx()
	var leafHashes [][32]byte
	for _, req := range reportRequests {
		leafHashes = append(leafHashes, mctx.Hash(req.Raw.Data))
	}

	tree, err := merklemulti.NewTree(mctx, leafHashes)
	require.NoError(t, err)
	rootIndex := -1
	for i, root := range report.MerkleRoots {
		if tree.Root() == root {
			rootIndex = i
		}

	}
	if rootIndex < 0 {
		t.Log("Merkle root does not match any root in the report")
		t.FailNow()
	}

	exists, err := client.Dest.BlobVerifier.GetMerkleRoot(nil, tree.Root())
	require.NoError(t, err)
	if exists.Uint64() < 1 {
		panic("Path is not present in the offramp")
	}
	index := request.Message.SequenceNumber - report.Intervals[rootIndex].Min
	client.Dest.logger.Info("index is ", index)
	return tree.Prove([]int{int(index)})
}

func (client *CCIPClient) TryGetTokensFromPausedPool() {
	client.PauseOnrampPool()

	paused, err := client.Source.TokenPools[0].Paused(nil)
	helpers.PanicErr(err)
	if !paused {
		helpers.PanicErr(errors.New("Should be paused"))
	}

	client.Source.Owner.GasLimit = 2e6
	tx, err := client.Source.TokenPools[0].LockOrBurn(client.Source.Owner, big.NewInt(1000))
	helpers.PanicErr(err)
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), false)
}

// SendToDappWithExecution executes a cross chain transactions using the sender dapp interface.
func (client *CCIPClient) SendToDappWithExecution(t *testing.T, source SourceClient, from *bind.TransactOpts, toAddress common.Address, amount *big.Int) *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested {
	SourceBlockNumber := GetCurrentBlockNumber(source.Client.Client)

	tx, err := source.SenderDapp.SendMessage(from, subscription_sender_dapp.CCIPEVM2AnySubscriptionMessage{
		Receiver: testhelpers.MustEncodeAddress(t, toAddress),
		Tokens:   []common.Address{source.LinkTokenAddress},
		Amounts:  []*big.Int{amount},
		GasLimit: big.NewInt(100_000),
	})
	helpers.PanicErr(err)
	source.logger.Infof("Send tokens tx %s", helpers.ExplorerLink(source.ChainId.Int64(), tx.Hash()))

	return WaitForCrossChainSendRequest(source, SourceBlockNumber, tx.Hash())
}

// SendToOnrampWithExecution executes a cross chain transactions using the onramp interface.
func (client *CCIPClient) SendToOnrampWithExecution(t *testing.T, source SourceClient, from *bind.TransactOpts, toAddress common.Address, amount *big.Int) *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested {
	SourceBlockNumber := GetCurrentBlockNumber(source.Client.Client)

	senderAndReceiver, err := utils.ABIEncode(`[{"type":"address"}, {"type":"address"}]`, source.Owner.From, source.Owner.From)
	helpers.PanicErr(err)

	payload := evm_2_any_subscription_onramp_router.CCIPEVM2AnySubscriptionMessage{
		Tokens:   []common.Address{},
		Amounts:  []*big.Int{},
		Receiver: testhelpers.MustEncodeAddress(t, toAddress),
		Data:     senderAndReceiver,
		GasLimit: big.NewInt(3e5),
	}
	source.logger.Infof("Send tx with payload %+v", payload)

	tx, err := source.OnRampRouter.CcipSend(from, client.Dest.ChainId, payload)
	if err != nil {
		t.Log(err.Error())
		printRevertReason(err, evm_2_any_subscription_onramp_router.EVM2AnySubscriptionOnRampRouterABI)
	}
	helpers.PanicErr(err)
	source.logger.Infof("Send tokens tx %s", helpers.ExplorerLink(source.ChainId.Int64(), tx.Hash()))
	return WaitForCrossChainSendRequest(source, SourceBlockNumber, tx.Hash())
}

func printRevertReason(errorData interface{}, abiString string) {
	dataError := errorData.(rpc.DataError)
	data, err := hex.DecodeString(dataError.ErrorData().(string)[2:])
	helpers.PanicErr(err)
	jsonABI, err := abi.JSON(strings.NewReader(abiString))
	helpers.PanicErr(err)
	for k, abiError := range jsonABI.Errors {
		if bytes.Equal(data[:4], abiError.ID.Bytes()[:4]) {
			// Found a matching error
			v, err := abiError.Unpack(data)
			helpers.PanicErr(err)
			fmt.Printf("Error \"%v\" args \"%v\"\n", k, v)
			return
		}
	}
}

// WaitForCrossChainSendRequest checks on chain for a successful onramp send event with the given tx hash.
// If not immediately found it will keep retrying in intervals of the globally specified RetryTiming.
func WaitForCrossChainSendRequest(source SourceClient, fromBlockNum uint64, txhash common.Hash) *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested {
	filter := bind.FilterOpts{Start: fromBlockNum}
	source.logger.Infof("Waiting for cross chain send... ")

	for {
		iterator, err := source.OnRamp.FilterCCIPSendRequested(&filter)
		helpers.PanicErr(err)
		for iterator.Next() {
			if iterator.Event.Raw.TxHash.Hex() == txhash.Hex() {
				source.logger.Infof("Cross chain send event found in tx: %s ", helpers.ExplorerLink(source.ChainId.Int64(), txhash))
				return iterator.Event
			}
		}
		time.Sleep(shared.RetryTiming)
	}
}

func (client *CCIPClient) PauseOfframpPool() {
	paused, err := client.Dest.TokenPools[0].Paused(nil)
	helpers.PanicErr(err)
	if paused {
		return
	}
	client.Dest.logger.Info("pausing offramp pool...")
	tx, err := client.Dest.TokenPools[0].Pause(client.Dest.Owner)
	helpers.PanicErr(err)
	client.Dest.logger.Info("Offramp pool paused, tx hash: %s", tx.Hash())
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) PauseOnrampPool() {
	paused, err := client.Source.TokenPools[0].Paused(nil)
	helpers.PanicErr(err)
	if paused {
		return
	}
	client.Source.logger.Info("pausing onramp pool...")
	tx, err := client.Source.TokenPools[0].Pause(client.Source.Owner)
	helpers.PanicErr(err)
	client.Source.logger.Info("Onramp pool paused, tx hash: %s", tx.Hash())
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) UnpauseOfframpPool() {
	paused, err := client.Dest.TokenPools[0].Paused(nil)
	helpers.PanicErr(err)
	if !paused {
		return
	}
	client.Dest.logger.Info("unpausing offramp pool...")
	tx, err := client.Dest.TokenPools[0].Unpause(client.Dest.Owner)
	helpers.PanicErr(err)
	client.Dest.logger.Info("Offramp pool unpaused, tx hash: %s", tx.Hash())
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) UnpauseOnrampPool() {
	paused, err := client.Source.TokenPools[0].Paused(nil)
	helpers.PanicErr(err)
	if !paused {
		return
	}
	client.Source.logger.Info("unpausing onramp pool...")
	tx, err := client.Source.TokenPools[0].Unpause(client.Source.Owner)
	helpers.PanicErr(err)
	client.Source.logger.Info("Onramp pool unpaused, tx hash: %s", tx.Hash())
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) PauseOnramp() {
	paused, err := client.Source.OnRamp.Paused(nil)
	helpers.PanicErr(err)
	if paused {
		return
	}
	client.Source.logger.Info("pausing onramp...")
	tx, err := client.Source.OnRamp.Pause(client.Source.Owner)
	helpers.PanicErr(err)
	client.Source.logger.Info("Onramp paused, tx hash: %s", tx.Hash())
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) PauseBlobVerifier() {
	paused, err := client.Dest.BlobVerifier.Paused(nil)
	helpers.PanicErr(err)
	if paused {
		return
	}
	client.Dest.logger.Info("pausing offramp...")
	tx, err := client.Dest.BlobVerifier.Pause(client.Dest.Owner)
	helpers.PanicErr(err)
	client.Dest.logger.Info("Offramp paused, tx hash: %s", tx.Hash())
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) UnpauseOnramp() {
	paused, err := client.Source.OnRamp.Paused(nil)
	helpers.PanicErr(err)
	if !paused {
		return
	}
	client.Source.logger.Info("unpausing onramp...")
	tx, err := client.Source.OnRamp.Unpause(client.Source.Owner)
	helpers.PanicErr(err)
	client.Source.logger.Info("Onramp unpaused, tx hash: %s", tx.Hash())
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) UnpauseBlobVerifier() {
	paused, err := client.Dest.BlobVerifier.Paused(nil)
	helpers.PanicErr(err)
	if !paused {
		return
	}
	client.Dest.logger.Info("unpausing offramp...")
	tx, err := client.Dest.BlobVerifier.Unpause(client.Dest.Owner)
	helpers.PanicErr(err)
	client.Dest.logger.Info("Offramp unpaused, tx hash: %s", tx.Hash())
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) UnpauseAll() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		client.UnpauseOnramp()
	}()
	go func() {
		defer wg.Done()
		client.UnpauseBlobVerifier()
	}()
	go func() {
		defer wg.Done()
		client.UnpauseOnrampPool()
	}()
	go func() {
		defer wg.Done()
		client.UnpauseOfframpPool()
	}()
	wg.Wait()
}

func (client *CCIPClient) SetOCRConfig(env dione.Environment) {
	verifierOCRConfig, err := client.Dest.BlobVerifier.LatestConfigDetails(&bind.CallOpts{})
	helpers.PanicErr(err)
	if verifierOCRConfig.BlockNumber != 0 {
		client.Dest.logger.Infof("BlobVerifier OCR config already found: %+v", verifierOCRConfig.ConfigDigest)
		client.Dest.logger.Infof("The new config will overwrite the current one.")
	}

	rampOCRConfig, err := client.Dest.OffRamp.LatestConfigDetails(&bind.CallOpts{})
	helpers.PanicErr(err)
	if rampOCRConfig.BlockNumber != 0 {
		client.Dest.logger.Infof("OffRamp OCR config already found: %+v", rampOCRConfig.ConfigDigest)
		client.Dest.logger.Infof("The new config will overwrite the current one.")
	}

	ccipConfig, err := ccip.OffchainConfig{
		SourceIncomingConfirmations: 10,
		DestIncomingConfirmations:   10,
	}.Encode()
	helpers.PanicErr(err)

	don := dione.NewOfflineDON(env, client.Dest.logger)
	faults := len(don.Config.Nodes) / 3

	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper2.ContractSetConfigArgsForTests(
		70*time.Second, // deltaProgress
		5*time.Second,  // deltaResend
		30*time.Second, // deltaRound
		2*time.Second,  // deltaGrace
		40*time.Second, // deltaStage
		3,
		[]int{1, 1, 2, 3}, // Transmission schedule: 1 oracle in first deltaStage, 2 in the second and so on.
		don.GenerateOracleIdentities(client.Dest.ChainId.String()),
		ccipConfig,
		5*time.Second,
		32*time.Second,
		20*time.Second,
		10*time.Second,
		10*time.Second,
		faults,
		nil,
	)
	helpers.PanicErr(err)

	signerAddresses, err := ocrcommon.OnchainPublicKeyToAddress(signers)
	helpers.PanicErr(err)
	transmitterAddresses, err := ocrcommon.AccountToAddress(transmitters)
	helpers.PanicErr(err)

	tx, err := client.Dest.BlobVerifier.SetConfig0(
		client.Dest.Owner,
		signerAddresses,
		transmitterAddresses,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	helpers.PanicErr(err)
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	client.Dest.logger.Infof("Config set on blob verifier %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), tx.Hash()))

	tx, err = client.Dest.OffRamp.SetConfig0(
		client.Dest.Owner,
		signerAddresses,
		transmitterAddresses,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	helpers.PanicErr(err)
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	client.Dest.logger.Infof("Config set on offramp %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), tx.Hash()))
}

func (client *CCIPClient) AcceptOwnership(t *testing.T) {
	tx, err := client.Dest.BlobVerifier.AcceptOwnership(client.Dest.Owner)
	require.NoError(t, err)
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)

	tx, err = client.Dest.OffRamp.AcceptOwnership(client.Dest.Owner)
	require.NoError(t, err)
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) PrepareSetSenders(t *testing.T) {
	sender := []common.Address{client.Source.SenderDapp.Address()}
	tx, err := client.Dest.OffRampRouter.PrepareSetSubscriptionSenders(client.Dest.Owner, client.Dest.ReceiverDapp.Address(), sender)
	require.NoError(t, err)
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) SetSubscriptionSenders(t *testing.T) {
	sender := []common.Address{client.Source.SenderDapp.Address()}
	tx, err := client.Dest.OffRampRouter.SetSubscriptionSenders(client.Dest.Owner, client.Dest.ReceiverDapp.Address(), sender)
	require.NoError(t, err)
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

type tokenPoolRegistry interface {
	Address() common.Address
	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)
	GetPool(opts *bind.CallOpts, token common.Address) (common.Address, error)
	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)
	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)
}

type aggregateRateLimiter interface {
	Address() common.Address
	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)
	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)
}

func syncPools(client *Client, registry tokenPoolRegistry, bridgeTokens []common.Address, txOpts *bind.TransactOpts) []*types.Transaction {
	registeredTokens, err := registry.GetPoolTokens(&bind.CallOpts{})
	require.NoError(client.t, err)

	pendingTxs := make([]*types.Transaction, 0)
	// remove registered tokenPools not present in config
	for _, token := range registeredTokens {
		if !slices.Contains(bridgeTokens, token) {
			pool, err := registry.GetPool(&bind.CallOpts{}, token)
			require.NoError(client.t, err)
			tx, err := registry.RemovePool(txOpts, token, pool)
			require.NoError(client.t, err)
			client.logger.Infof("removePool(token=%s, pool=%s) from registry=%s: tx=%s", token, pool, registry.Address(), tx.Hash())
			pendingTxs = append(pendingTxs, tx) // queue txs for wait
			txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1)) // increment nonce
		}
	}
	// add tokenPools present in config and not yet registered
	for i, token := range bridgeTokens {
		// remove tokenPools not present in config
		if !slices.Contains(registeredTokens, token) {
			pool := client.TokenPools[i].Address()
			tx, err := registry.AddPool(txOpts, token, pool)
			require.NoError(client.t, err)
			client.logger.Infof("addPool(token=%s, pool=%s) from registry=%s: tx=%s", token, pool, registry.Address(), tx.Hash())
			pendingTxs = append(pendingTxs, tx) // queue txs for wait
			txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1)) // increment nonce
		}
	}
	return pendingTxs
}

func syncPrices(client *Client, limiter aggregateRateLimiter, txOpts *bind.TransactOpts) *types.Transaction {
	// sync tokenPrices if needed
	if len(client.TokenPrices) == 0 {
		return nil
	}
	if len(client.TokenPrices) != len(client.BridgeTokens) {
		client.t.Fatal("if config.TokenPrices isn't empty, it must correspond to BridgeTokens")
	}
	limiterTokenPrices, err := limiter.GetPricesForTokens(&bind.CallOpts{}, client.BridgeTokens)
	require.NoError(client.t, err)
	for i := range client.BridgeTokens {
		// on first difference, setPrices then return
		if client.TokenPrices[i].Cmp(limiterTokenPrices[i]) != 0 {
			tx, err := limiter.SetPrices(txOpts, client.BridgeTokens, client.TokenPrices)
			require.NoError(client.t, err)
			client.logger.Infof("setPrices(tokens=%s, prices=%s) for limiter=%s: tx=%s", client.BridgeTokens, client.TokenPrices, limiter.Address(), tx.Hash())
			txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1)) // increment nonce
			return tx
		}
	}
	return nil
}

func waitPendingTxs(client *Client, pendingTxs *[]*types.Transaction) {
	// wait for all queued txs
	for _, tx := range *pendingTxs {
		shared.WaitForMined(client.t, client.logger, client.Client, tx.Hash(), true)
	}
	*pendingTxs = (*pendingTxs)[:0] // clear pending txs
}

func (client *CCIPClient) SyncTokenPools(t *testing.T) {
	// use local txOpts, so we can cache/increment nonce manually before waiting on all txs
	sourceTxOpts := *client.Source.Owner
	sourceTxOpts.GasLimit = 120_000 // hardcode gasLimit (enough for each tx here), to avoid race from mis-estimating
	sourcePendingNonce, err := client.Source.Client.Client.PendingNonceAt(context.Background(), client.Source.Owner.From)
	require.NoError(t, err)
	sourceTxOpts.Nonce = big.NewInt(int64(sourcePendingNonce))

	// onRamp maps source tokens to source pools
	sourcePendingTxs := syncPools(&client.Source.Client, client.Source.OnRamp, client.Source.BridgeTokens, &sourceTxOpts)

	// same as above, for offRamp
	destTxOpts := *client.Dest.Owner
	destTxOpts.GasLimit = 120_000 // hardcode gasLimit (enough for each tx here), to avoid race from mis-estimating
	destPendingNonce, err := client.Dest.Client.Client.PendingNonceAt(context.Background(), client.Dest.Owner.From)
	require.NoError(t, err)
	destTxOpts.Nonce = big.NewInt(int64(destPendingNonce))

	// offRamp maps *source* tokens to *dest* pools
	destPendingTxs := syncPools(&client.Dest.Client, client.Dest.OffRamp, client.Source.BridgeTokens, &destTxOpts)

	waitPendingTxs(&client.Source.Client, &sourcePendingTxs)
	waitPendingTxs(&client.Dest.Client, &destPendingTxs)

	if tx := syncPrices(&client.Source.Client, client.Source.OnRamp, &sourceTxOpts); tx != nil {
		sourcePendingTxs = append(sourcePendingTxs, tx)
	}
	if tx := syncPrices(&client.Dest.Client, client.Dest.OffRamp, &destTxOpts); tx != nil {
		destPendingTxs = append(destPendingTxs, tx)
	}

	waitPendingTxs(&client.Source.Client, &sourcePendingTxs)
	waitPendingTxs(&client.Dest.Client, &destPendingTxs)
}
