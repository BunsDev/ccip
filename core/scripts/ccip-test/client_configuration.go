package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/lock_unlock_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_sender"
	"github.com/smartcontractkit/chainlink/core/services/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
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

func NewSourceClient(config EvmChainConfig) SourceClient {
	client := GetClient(config.EthUrl)
	LinkToken, err := link_token_interface.NewLinkToken(config.LinkToken, client)
	PanicErr(err)
	lockUnlockPool, err := lock_unlock_pool.NewLockUnlockPool(config.LockUnlockPool, client)
	PanicErr(err)
	afn, err := afn_contract.NewAFNContract(config.Afn, client)
	PanicErr(err)
	singleTokenOnramp, err := single_token_onramp.NewSingleTokenOnRamp(config.SingleTokenOnramp, client)
	PanicErr(err)
	simpleTokenSender, err := single_token_sender.NewEOASingleTokenSender(config.SingleTokenSender, client)
	PanicErr(err)

	return SourceClient{
		Client: Client{
			Client:           client,
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

func NewDestinationClient(config EvmChainConfig) DestClient {
	client := GetClient(config.EthUrl)
	LinkToken, err := link_token_interface.NewLinkToken(config.LinkToken, client)
	PanicErr(err)
	lockUnlockPool, err := lock_unlock_pool.NewLockUnlockPool(config.LockUnlockPool, client)
	PanicErr(err)
	afn, err := afn_contract.NewAFNContract(config.Afn, client)
	PanicErr(err)
	singleTokenOfframp, err := single_token_offramp.NewSingleTokenOffRamp(config.SingleTokenOfframp, client)
	PanicErr(err)
	messageExecutor, err := message_executor.NewMessageExecutor(config.MessageExecutor, client)
	PanicErr(err)
	simpleMessageReceiver, err := simple_message_receiver.NewSimpleMessageReceiver(config.SimpleMessageReceiver, client)
	PanicErr(err)
	singleTokenReceiver, err := single_token_receiver.NewEOASingleTokenReceiver(config.SingleTokenReceiver, client)
	PanicErr(err)

	return DestClient{
		Client: Client{
			Client:           client,
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

type CCIPClient struct {
	Source SourceClient
	Dest   DestClient
}

func NewCcipClient(sourceConfig EvmChainConfig, destConfig EvmChainConfig, ownerKey string, seedKey string) CCIPClient {
	source := NewSourceClient(sourceConfig)
	source.SetOwnerAndUsers(ownerKey, seedKey)
	dest := NewDestinationClient(destConfig)
	dest.SetOwnerAndUsers(ownerKey, seedKey)

	return CCIPClient{
		Source: source,
		Dest:   dest,
	}
}

func GetOwner(ownerPrivateKey string, chainId *big.Int) *bind.TransactOpts {
	ownerKey, err := crypto.HexToECDSA(ownerPrivateKey)
	PanicErr(err)
	user, err := bind.NewKeyedTransactorWithChainID(ownerKey, chainId)
	PanicErr(err)
	fmt.Println("--- Owner address ")
	fmt.Println(user.From.Hex())
	// to not use geth-only tip fee method
	// https://github.com/ethereum/go-ethereum/pull/23484
	user.GasTipCap = big.NewInt(2e9)
	return user
}

func GetClient(ethUrl string) *ethclient.Client {
	client, err := ethclient.Dial(ethUrl)
	PanicErr(err)
	return client
}

func (client *Client) SetOwnerAndUsers(ownerPrivateKey string, seedKey string) {
	client.Owner = GetOwner(ownerPrivateKey, client.ChainId)

	var users []*bind.TransactOpts
	seedKeyWithoutFirstChar := seedKey[1:]
	fmt.Println("--- Addresses of the seed key")
	for i := 0; i <= 9; i++ {
		key, err := crypto.HexToECDSA(strconv.Itoa(i) + seedKeyWithoutFirstChar)
		PanicErr(err)
		user, err := bind.NewKeyedTransactorWithChainID(key, client.ChainId)
		PanicErr(err)
		// to not use geth-only tip fee method
		// https://github.com/ethereum/go-ethereum/pull/23484
		user.GasTipCap = big.NewInt(2e9)
		users = append(users, user)
		fmt.Println(user.From.Hex())
	}
	fmt.Println("---")

	client.Users = users
}

func (client *Client) AssureHealth() {
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

func (client *Client) ApproveLinkFrom(user *bind.TransactOpts, approvedFor common.Address, amount *big.Int) {
	ctx := context.Background()
	tx, err := client.LinkToken.Approve(user, approvedFor, amount)
	PanicErr(err)

	WaitForMined(ctx, client.Client, tx.Hash(), true)
	fmt.Println("approve tx hash", tx.Hash().Hex())
}

func (client *Client) ApproveLink(approvedFor common.Address, amount *big.Int) {
	client.ApproveLinkFrom(client.Owner, approvedFor, amount)
}

func (client CCIPClient) SendMessage() {
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
	PanicErr(err)
	WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) DonExecutionHappyPath() {
	amount := big.NewInt(100)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)
	DestBlockNum := GetCurrentBlockNumber(client.Dest.Client.Client)
	crossChainRequest := client.SendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, client.Dest.MessageExecutor.Address())
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
	PanicErr(err)
	defer sub.Unsubscribe()

	select {
	case event := <-events:
		fmt.Printf("found Destination execution in transaction: %s\n", event.Raw.TxHash.Hex())
		return
	case err := <-sub.Err():
		panic(err)
	}
}

func (client CCIPClient) ExternalExecutionHappyPath() {
	ctx := context.Background()
	offrampBlockNumber := GetCurrentBlockNumber(client.Dest.Client.Client)
	onrampBlockNumber := GetCurrentBlockNumber(client.Source.Client.Client)

	amount, _ := new(big.Int).SetString("10", 10)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)

	onrampRequest := client.SendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
	sequenceNumber := onrampRequest.Message.SequenceNumber.Int64()

	// Gets the report that our transaction is included in
	fmt.Println("Getting report")
	report, err := client.GetReportForSequenceNumber(ctx, sequenceNumber, offrampBlockNumber)
	PanicErr(err)

	// Get all requests included in the given report
	fmt.Println("Getting recent cross chain requests")
	requests := client.GetCrossChainSendRequestsForRange(ctx, report, onrampBlockNumber)

	// Generate the proof
	fmt.Println("Generating proof")
	proof := client.ValidateMerkleRoot(onrampRequest, requests, report)

	// Execute the transaction on the offramp
	fmt.Println("Executing offramp TX")
	tx, err := client.ExecuteOfframpTransaction(proof, onrampRequest.Raw.Data)
	PanicErr(err)
	WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) CrossChainSendPausedOnrampShouldFail() {
	client.PauseOnramp()
	amount := big.NewInt(100)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)
	client.Source.Owner.GasLimit = 1e6
	tx, err := client.Source.SingleTokenSender.SendTokens(client.Source.Owner, client.Dest.Owner.From, amount, client.Dest.MessageExecutor.Address())
	PanicErr(err)
	WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), false)
}

func (client CCIPClient) CrossChainSendPausedOfframpShouldFail() {
	client.PauseOfframp()
	ctx := context.Background()
	offrampBlockNumber := GetCurrentBlockNumber(client.Dest.Client.Client)

	amount, _ := new(big.Int).SetString("10", 10)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)

	onrampRequest := client.SendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
	sequenceNumber := onrampRequest.Message.SequenceNumber.Int64()

	fmt.Println("Waiting for report...")
	_, err := client.GetReportForSequenceNumber(ctx, sequenceNumber, offrampBlockNumber)
	if err.Error() == "No report found within the given time" {
		fmt.Println("Success, no oracle report sent to paused offramp.")
	} else {
		panic(fmt.Errorf("report found despite paused offramp"))
	}
}

func (client CCIPClient) NotEnoughFundsInBucketShouldFail() {
	amount := big.NewInt(2e18) // 2 LINK, bucket size is 1 LINK
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)
	client.Source.Owner.GasLimit = 1e6
	tx, err := client.Source.SingleTokenSender.SendTokens(client.Source.Owner, client.Dest.Owner.From, amount, client.Dest.MessageExecutor.Address())
	PanicErr(err)
	WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), false)
}

func (client CCIPClient) ExternalExecutionSubmitOfframpTwiceShouldFail() {
	ctx := context.Background()
	offrampBlockNumber := GetCurrentBlockNumber(client.Dest.Client.Client)
	onrampBlockNumber := GetCurrentBlockNumber(client.Source.Client.Client)

	amount, _ := new(big.Int).SetString("10", 10)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)

	onrampRequest := client.SendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
	sequenceNumber := onrampRequest.Message.SequenceNumber.Int64()

	// Gets the report that our transaction is included in
	fmt.Println("Getting report")
	report, err := client.GetReportForSequenceNumber(ctx, sequenceNumber, offrampBlockNumber)
	PanicErr(err)

	// Get all requests included in the given report
	fmt.Println("Getting recent cross chain requests")
	requests := client.GetCrossChainSendRequestsForRange(ctx, report, onrampBlockNumber)

	// Generate the proof
	fmt.Println("Generating proof")
	proof := client.ValidateMerkleRoot(onrampRequest, requests, report)

	// Execute the transaction on the offramp
	fmt.Println("Executing first offramp TX - should succeed")
	tx, err := client.ExecuteOfframpTransaction(proof, onrampRequest.Raw.Data)
	PanicErr(err)
	WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), true)

	// Execute the transaction on the offramp
	fmt.Println("Executing second offramp TX - should fail")
	client.Dest.Owner.GasLimit = 1e6
	tx, err = client.ExecuteOfframpTransaction(proof, onrampRequest.Raw.Data)
	PanicErr(err)
	WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), false)
}

// ScalingAndBatching should scale so that we see batching on the nodes
func (client CCIPClient) ScalingAndBatching() {
	amount := big.NewInt(10)
	toAddress := common.HexToAddress("0x57359120D900fab8cE74edC2c9959b21660d3887")

	var wg sync.WaitGroup
	for _, user := range client.Source.Users {
		wg.Add(1)
		go func(user *bind.TransactOpts) {
			defer wg.Done()
			client.Source.ApproveLinkFrom(user, client.Source.SingleTokenSender.Address(), amount)
			crossChainRequest := client.SendToOnrampWithExecution(client.Source, user, toAddress, amount, client.Dest.MessageExecutor.Address())
			fmt.Println("Don executed tx submitted with sequence number: ", crossChainRequest.Message.SequenceNumber.Int64())
		}(user)
	}
	wg.Wait()
	fmt.Println("Sent 10 txs to onramp.")
}

func (client CCIPClient) ExecuteOfframpTransaction(proof ccip.MerkleProof, encodedMessage []byte) (*types.Transaction, error) {
	decodedMsg, err := abihelpers.DecodeCCIPMessage(encodedMessage)
	PanicErr(err)
	_, err = abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{*decodedMsg})
	PanicErr(err)

	tx, err := client.Dest.SingleTokenOfframp.ExecuteTransaction(client.Dest.Owner, proof.PathForExecute(), *decodedMsg, proof.Index())
	return tx, errors.Wrap(err, "executing offramp tx")
}

func (client CCIPClient) GetCrossChainSendRequestsForRange(
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
	PanicErr(err)

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
		PanicErr(errors.New("Not all cross chain requests found in the last 1000 blocks"))
	}

	return requests
}

func (client CCIPClient) GetReportForSequenceNumber(ctx context.Context, sequenceNumber int64, minBlockNumber uint64) (single_token_offramp.CCIPRelayReport, error) {
	report, err := client.Dest.SingleTokenOfframp.GetLastReport(&bind.CallOpts{Context: ctx, Pending: false})
	if err != nil {
		return single_token_offramp.CCIPRelayReport{}, err
	}

	// our tx is in the latest report
	if sequenceNumber >= report.MinSequenceNumber.Int64() && sequenceNumber <= report.MaxSequenceNumber.Int64() {
		return report, nil
	}
	// report isn't out yet, it will be in a future report
	if sequenceNumber > report.MaxSequenceNumber.Int64() {
		maxIterations := CrossChainTimout / RetryTiming
		for i := 0; i < int(maxIterations); i++ {
			report, err = client.Dest.SingleTokenOfframp.GetLastReport(&bind.CallOpts{Context: ctx, Pending: false})
			if err != nil {
				return single_token_offramp.CCIPRelayReport{}, err
			}
			if sequenceNumber >= report.MinSequenceNumber.Int64() && sequenceNumber <= report.MaxSequenceNumber.Int64() {
				return report, nil
			}
			time.Sleep(RetryTiming)
		}
		return single_token_offramp.CCIPRelayReport{}, errors.New("No report found within the given timeout")
	}

	// it is in a past report, start looking at the earliest block number possible, the one
	// before we started the entire transaction on the onramp.
	reports, err := client.Dest.SingleTokenOfframp.FilterReportAccepted(&bind.FilterOpts{
		Start:   minBlockNumber,
		End:     nil,
		Context: ctx,
	})
	if err != nil {
		return single_token_offramp.CCIPRelayReport{}, err
	}

	for reports.Next() {
		report = reports.Event.Report
		if sequenceNumber >= report.MinSequenceNumber.Int64() && sequenceNumber <= report.MaxSequenceNumber.Int64() {
			return report, nil
		}
	}

	// Somehow the transaction was not included in any report within blocks produced after
	// the transaction was initialized but the sequence number is lower than we are currently at
	return single_token_offramp.CCIPRelayReport{}, errors.New("No report found for given sequence number")
}

func GetCurrentBlockNumber(chain *ethclient.Client) uint64 {
	blockNumber, err := chain.BlockNumber(context.Background())
	PanicErr(err)
	return blockNumber
}

func (client CCIPClient) ValidateMerkleRoot(
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
	if !bytes.Equal(root[:], report.MerkleRoot[:]) {
		PanicErr(errors.New("Merkle root does not match the report"))
	}

	genRoot := ccip.GenerateMerkleRoot(leaves[int(index.Int64())], proof)
	if !reflect.DeepEqual(root[:], genRoot[:]) {
		PanicErr(errors.New("Root does not verify"))
	}

	exists, err := client.Dest.SingleTokenOfframp.GetMerkleRoot(nil, root)
	PanicErr(err)
	if exists.Uint64() < 1 {
		PanicErr(errors.New("Proof is not present in the offramp"))
	}
	return proof
}

func (client CCIPClient) TryGetTokensFromPausedPool() {
	client.PauseOnrampPool()

	paused, err := client.Source.LockUnlockPool.Paused(nil)
	PanicErr(err)
	if !paused {
		PanicErr(errors.New("Should be paused"))
	}

	client.Source.Owner.GasLimit = 2e6
	tx, err := client.Source.LockUnlockPool.LockOrBurn(client.Source.Owner, client.Source.Owner.From, big.NewInt(1000))
	PanicErr(err)
	WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), false)
}

func (client CCIPClient) SendToOnrampWithExecution(Source SourceClient, from *bind.TransactOpts, toAddress common.Address, amount *big.Int, executor common.Address) *single_token_onramp.SingleTokenOnRampCrossChainSendRequested {
	ctx := context.Background()
	SourceBlockNumber := GetCurrentBlockNumber(Source.Client.Client)

	tx, err := Source.SingleTokenSender.SendTokens(from, toAddress, amount, executor)
	PanicErr(err)
	fmt.Println("send tokens hash ", tx.Hash())
	WaitForMined(ctx, Source.Client.Client, tx.Hash(), true)

	return WaitForCrossChainSendRequest(Source, SourceBlockNumber, tx.Hash())
}

func WaitForCrossChainSendRequest(Source SourceClient, fromBlockNum uint64, txhash common.Hash) *single_token_onramp.SingleTokenOnRampCrossChainSendRequested {
	filter := bind.FilterOpts{Start: fromBlockNum}
	for {
		iterator, err := Source.SingleTokenOnramp.FilterCrossChainSendRequested(&filter)
		PanicErr(err)
		for iterator.Next() {
			if iterator.Event.Raw.TxHash.Hex() == txhash.Hex() {
				fmt.Println("cross chain send event found in tx: ", txhash.Hex())
				return iterator.Event
			}
		}
		time.Sleep(RetryTiming)
	}
}

func (client CCIPClient) PauseOfframpPool() {
	paused, err := client.Dest.LockUnlockPool.Paused(nil)
	PanicErr(err)
	if paused {
		return
	}
	fmt.Println("pausing offramp pool...")
	tx, err := client.Dest.LockUnlockPool.Pause(client.Dest.Owner)
	PanicErr(err)
	fmt.Println("Offramp pool paused, tx hash:", tx.Hash())
	WaitForMined(context.Background(), client.Dest.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) PauseOnrampPool() {
	paused, err := client.Source.LockUnlockPool.Paused(nil)
	PanicErr(err)
	if paused {
		return
	}
	fmt.Println("pausing onramp pool...")
	tx, err := client.Source.LockUnlockPool.Pause(client.Source.Owner)
	PanicErr(err)
	fmt.Println("Onramp pool paused, tx hash:", tx.Hash())
	WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) UnpauseOfframpPool() {
	paused, err := client.Dest.LockUnlockPool.Paused(nil)
	PanicErr(err)
	if !paused {
		return
	}
	fmt.Println("unpausing offramp pool...")
	tx, err := client.Dest.LockUnlockPool.Unpause(client.Dest.Owner)
	PanicErr(err)
	fmt.Println("Offramp pool unpaused, tx hash:", tx.Hash())
	WaitForMined(context.Background(), client.Dest.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) UnpauseOnrampPool() {
	paused, err := client.Source.LockUnlockPool.Paused(nil)
	PanicErr(err)
	if !paused {
		return
	}
	fmt.Println("unpausing onramp pool...")
	tx, err := client.Source.LockUnlockPool.Unpause(client.Source.Owner)
	PanicErr(err)
	fmt.Println("Onramp pool unpaused, tx hash:", tx.Hash())
	WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) PauseOnramp() {
	paused, err := client.Source.SingleTokenOnramp.Paused(nil)
	PanicErr(err)
	if paused {
		return
	}
	fmt.Println("pausing onramp...")
	tx, err := client.Source.SingleTokenOnramp.Pause(client.Source.Owner)
	PanicErr(err)
	fmt.Println("Onramp paused, tx hash:", tx.Hash())
	WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) PauseOfframp() {
	paused, err := client.Dest.SingleTokenOfframp.Paused(nil)
	PanicErr(err)
	if paused {
		return
	}
	fmt.Println("pausing offramp...")
	tx, err := client.Dest.SingleTokenOfframp.Pause(client.Dest.Owner)
	PanicErr(err)
	fmt.Println("Offramp paused, tx hash:", tx.Hash())
	WaitForMined(context.Background(), client.Dest.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) UnpauseOnramp() {
	paused, err := client.Source.SingleTokenOnramp.Paused(nil)
	PanicErr(err)
	if !paused {
		return
	}
	fmt.Println("unpausing onramp...")
	tx, err := client.Source.SingleTokenOnramp.Unpause(client.Source.Owner)
	PanicErr(err)
	fmt.Println("Onramp unpaused, tx hash:", tx.Hash())
	WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) UnpauseOfframp() {
	paused, err := client.Dest.SingleTokenOfframp.Paused(nil)
	PanicErr(err)
	if !paused {
		return
	}
	fmt.Println("unpausing offramp...")
	tx, err := client.Dest.SingleTokenOfframp.Unpause(client.Dest.Owner)
	PanicErr(err)
	fmt.Println("Offramp unpaused, tx hash:", tx.Hash())
	WaitForMined(context.Background(), client.Dest.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) UnpauseAll() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		client.UnpauseOnramp()
	}()
	go func() {
		defer wg.Done()
		client.UnpauseOfframp()
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

func (client CCIPClient) SetConfig() {
	ccipConfig, err := ccip.OffchainConfig{
		SourceIncomingConfirmations: 0,
		DestIncomingConfirmations:   0,
	}.Encode()
	PanicErr(err)
	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper2.ContractSetConfigArgsForTests(
		60*time.Second, // deltaProgress
		1*time.Second,  // deltaResend
		20*time.Second, // deltaRound
		2*time.Second,  // deltaGrace
		30*time.Second, // deltaStage
		3,
		[]int{1, 2, 3, 4}, // Transmission schedule: 1 oracle in first deltaStage, 2 in the second and so on.
		Oracles,
		ccipConfig,
		1*time.Second,
		10*time.Second,
		20*time.Second,
		10*time.Second,
		10*time.Second,
		1, // faults
		nil,
	)
	PanicErr(err)

	ctx := context.Background()

	signerAddresses, err := ocrcommon.OnchainPublicKeyToAddress(signers)
	PanicErr(err)
	transmitterAddresses, err := ocrcommon.AccountToAddress(transmitters)
	PanicErr(err)

	tx, err := client.Dest.SingleTokenOfframp.SetConfig(
		client.Dest.Owner,
		signerAddresses,
		transmitterAddresses,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	PanicErr(err)
	WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), true)
	fmt.Println("Config set on offramp. Tx hash:", tx.Hash().Hex())

	tx, err = client.Dest.MessageExecutor.SetConfig(
		client.Dest.Owner,
		signerAddresses,
		transmitterAddresses,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	PanicErr(err)
	WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), true)
	fmt.Println("Config set on message executor. Tx hash:", tx.Hash().Hex())
}
