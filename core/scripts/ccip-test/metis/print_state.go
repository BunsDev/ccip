package metis

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_subscription_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ping_pong_demo"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/receiver_dapp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/subscription_sender_dapp"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/dione"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

func PrintCCIPState(source *rhea.EvmChainConfig, destination *rhea.EvmChainConfig) {
	printPoolBalances(source)
	printPoolBalances(destination)

	printDappSanityCheck(source)
	printDappSanityCheck(destination)

	printRampSanityCheck(source, destination.OnRamp)
	printRampSanityCheck(destination, source.OnRamp)

	printSourceSubscriptionBalances(source)
	printDestinationSubscriptionBalances(destination)

	printPaused(source)
	printPaused(destination)

	printRateLimitingStatus(source)
	printRateLimitingStatus(destination)

	printTxStatuses(source, destination)
}

type CCIPTXStatus struct {
	message     *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested
	relayReport *blob_verifier.BlobVerifierReportAccepted
	execStatus  *any_2_evm_subscription_offramp.EVM2EVMSubscriptionOffRampExecutionStateChanged
}

type ExecutionStatus uint8

const (
	Untouched  ExecutionStatus = 0
	InProgress ExecutionStatus = 1
	Success    ExecutionStatus = 2
	Failed     ExecutionStatus = 3
)

func (e ExecutionStatus) String() string {
	switch e {
	case Untouched:
		return "Untouched"
	case InProgress:
		return "InProgress"
	case Success:
		return "Success"
	case Failed:
		return "Failed"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

func printTxStatuses(source *rhea.EvmChainConfig, destination *rhea.EvmChainConfig) {
	onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(source.OnRamp, source.Client)
	helpers.PanicErr(err)

	block, err := source.Client.BlockNumber(context.Background())
	helpers.PanicErr(err)

	sendRequested, err := onRamp.FilterCCIPSendRequested(&bind.FilterOpts{
		Start: block - 9990,
	})
	helpers.PanicErr(err)

	txs := make(map[uint64]*CCIPTXStatus)
	maxSeqNum := uint64(0)
	minSeqNum := uint64(1)
	var seqNums []uint64

	for sendRequested.Next() {
		txs[sendRequested.Event.Message.SequenceNumber] = &CCIPTXStatus{
			message: sendRequested.Event,
		}
		if sendRequested.Event.Message.SequenceNumber > maxSeqNum {
			maxSeqNum = sendRequested.Event.Message.SequenceNumber
		}
		if minSeqNum == 1 {
			minSeqNum = sendRequested.Event.Message.SequenceNumber
		}
		seqNums = append(seqNums, sendRequested.Event.Message.SequenceNumber)
	}

	blobVerifier, err := blob_verifier.NewBlobVerifier(destination.BlobVerifier, destination.Client)
	helpers.PanicErr(err)

	block, err = destination.Client.BlockNumber(context.Background())
	helpers.PanicErr(err)

	reports, err := blobVerifier.FilterReportAccepted(&bind.FilterOpts{
		Start: block - 9990,
	})
	helpers.PanicErr(err)

	for reports.Next() {
		for i, interval := range reports.Event.Report.Intervals {
			if reports.Event.Report.OnRamps[i] != source.OnRamp {
				continue
			}
			for j := interval.Min; j <= interval.Max; j++ {
				if _, ok := txs[j]; ok {
					txs[j].relayReport = reports.Event
				}
			}
		}
	}

	offRamp, err := any_2_evm_subscription_offramp.NewEVM2EVMSubscriptionOffRamp(destination.OffRamp, destination.Client)
	helpers.PanicErr(err)

	stateChanges, err := offRamp.FilterExecutionStateChanged(&bind.FilterOpts{
		Start: block - 9990,
	}, seqNums)
	helpers.PanicErr(err)

	for stateChanges.Next() {
		if _, ok := txs[stateChanges.Event.SequenceNumber]; !ok {
			txs[stateChanges.Event.SequenceNumber] = &CCIPTXStatus{}
			if stateChanges.Event.SequenceNumber > maxSeqNum {
				maxSeqNum = stateChanges.Event.SequenceNumber
			}
		}
		txs[stateChanges.Event.SequenceNumber].execStatus = stateChanges.Event
	}

	var sb strings.Builder
	sb.WriteString("\n")
	tableHeaders := []string{"SequenceNumber", "Relayed in block", "Execution status", "Executed in block", "Nonce"}
	headerLengths := []int{18, 18, 20, 18, 18}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	if minSeqNum > 1 {
		sb.WriteString(fmt.Sprintf("| %18d | %18d | %41s | %18s | \n", 1, minSeqNum-1, "Probably > 10k blocks in the past", ""))
	}

	for i := minSeqNum; i <= maxSeqNum; i++ {
		tx := txs[i]
		relayedAt := "-"
		if tx == nil {
			sb.WriteString(fmt.Sprintf("| %18d | %18s | %41s | %18s | \n", i, "TX MISSING", "", ""))
			continue
		}
		if tx.relayReport != nil {
			relayedAt = strconv.Itoa(int(tx.relayReport.Raw.BlockNumber))
		}

		if tx.message == nil {
			sb.WriteString(fmt.Sprintf("| %18s | %18s | %20v | %18d | %18s | \n", "MISSING", relayedAt, ExecutionStatus(tx.execStatus.State), tx.execStatus.Raw.BlockNumber, "-"))
		} else if tx.execStatus != nil {
			sb.WriteString(fmt.Sprintf("| %18d | %18s | %20v | %18d | %18d | %s \n",
				tx.message.Message.SequenceNumber,
				relayedAt,
				ExecutionStatus(tx.execStatus.State),
				tx.execStatus.Raw.BlockNumber,
				tx.message.Message.Nonce,
				helpers.ExplorerLink(destination.ChainId.Int64(), tx.execStatus.Raw.TxHash)))
		} else {
			sb.WriteString(fmt.Sprintf("| %18d | %18s | %20v | %18s | %18d | %s \n",
				tx.message.Message.SequenceNumber,
				relayedAt,
				"-",
				"-",
				tx.message.Message.Nonce,
				""))
		}
	}
	sb.WriteString(generateSeparator(headerLengths))

	destination.Logger.Info(sb.String())
}

func printDappSanityCheck(source *rhea.EvmChainConfig) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Dapp sanity checks for %s\n", helpers.ChainName(source.ChainId.Int64())))

	tableHeaders := []string{"Dapp", "Router Set"}
	headerLengths := []int{30, 15}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	senderDapp, err := subscription_sender_dapp.NewSubscriptionSenderDapp(source.TokenSender, source.Client)
	helpers.PanicErr(err)
	router, err := senderDapp.IOnRampRouter(&bind.CallOpts{})
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "Sender dapp", router == source.OnRampRouter))

	receiverDap, err := receiver_dapp.NewReceiverDapp(source.ReceiverDapp, source.Client)
	helpers.PanicErr(err)
	router, err = receiverDap.SRouter(&bind.CallOpts{})
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "Receiver dapp", router == source.OffRampRouter))

	if source.PingPongDapp != common.HexToAddress("") {
		pingDapp, err := ping_pong_demo.NewPingPongDemo(source.PingPongDapp, source.Client)
		helpers.PanicErr(err)
		receiver, sender, err := pingDapp.GetRouters(&bind.CallOpts{})
		helpers.PanicErr(err)
		sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "Ping dapp receiver", receiver == source.OffRampRouter))
		sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "Ping dapp sender", sender == source.OnRampRouter))
	}

	sb.WriteString(generateSeparator(headerLengths))

	source.Logger.Info(sb.String())
}

func printRampSanityCheck(chain *rhea.EvmChainConfig, sourceOnRamp common.Address) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Ramp checks for %s\n", helpers.ChainName(chain.ChainId.Int64())))

	tableHeaders := []string{"Contract", "Config correct"}
	headerLengths := []int{30, 15}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	afn, err := afn_contract.NewAFNContract(chain.Afn, chain.Client)
	helpers.PanicErr(err)
	badSignal, err := afn.BadSignalReceived(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "AFN healthy", !badSignal))

	onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(chain.OnRamp, chain.Client)
	helpers.PanicErr(err)
	router, err := onRamp.GetRouter(&bind.CallOpts{})
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "OnRamp Router set", router == chain.OnRampRouter))

	offRamp, err := any_2_evm_subscription_offramp.NewEVM2EVMSubscriptionOffRamp(chain.OffRamp, chain.Client)
	helpers.PanicErr(err)
	router, err = offRamp.GetRouter(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "OffRamp Router set", router == chain.OffRampRouter))

	configDetails, err := offRamp.LatestConfigDetails(&bind.CallOpts{})
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "OffRamp OCR2 configured", configDetails.ConfigCount != 0))

	blobVerifier, err := blob_verifier.NewBlobVerifier(chain.BlobVerifier, chain.Client)
	helpers.PanicErr(err)

	config, err := blobVerifier.GetConfig(&bind.CallOpts{})
	helpers.PanicErr(err)

	rampSet := false
	for _, ramp := range config.OnRamps {
		if ramp == sourceOnRamp {
			rampSet = true
		}
	}

	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "BlobVerifier Onramp set", rampSet))

	blobConfigDetails, err := blobVerifier.LatestConfigDetails(&bind.CallOpts{})
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "BlobVerifier OCR2 configured", blobConfigDetails.ConfigCount != 0))

	offRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(chain.OffRampRouter, chain.Client)
	helpers.PanicErr(err)

	isRamp, err := offRampRouter.IsOffRamp(&bind.CallOpts{}, chain.OffRamp)
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "OffRampRouter has offRamp Set", isRamp))

	sb.WriteString(generateSeparator(headerLengths))

	chain.Logger.Info(sb.String())
}

func printSourceSubscriptionBalances(source *rhea.EvmChainConfig) {
	onRampRouter, err := evm_2_any_subscription_onramp_router.NewEVM2AnySubscriptionOnRampRouter(source.OnRampRouter, source.Client)
	helpers.PanicErr(err)

	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Source subscription balances for %s\n", helpers.ChainName(source.ChainId.Int64())))

	tableHeaders := []string{"Sender", "Address", "Balance", "#Txs funded"}
	headerLengths := []int{20, 42, 22, 22}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	fee, err := onRampRouter.GetFee(&bind.CallOpts{})
	helpers.PanicErr(err)

	possibleTxs := "∞"
	balance, err := onRampRouter.GetBalance(&bind.CallOpts{}, source.Owner.From)
	helpers.PanicErr(err)
	if fee.Int64() != 0 {
		possibleTxs = big.NewInt(0).Div(balance, fee).String()
	}
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %22s |\n", "owner", source.Owner.From.Hex(), balance, possibleTxs))

	balance, err = onRampRouter.GetBalance(&bind.CallOpts{}, source.TokenSender)
	helpers.PanicErr(err)
	if fee.Int64() != 0 {
		possibleTxs = big.NewInt(0).Div(balance, fee).String()
	}
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %22s |\n", "sender dapp", source.TokenSender.Hex(), balance, possibleTxs))

	balance, err = onRampRouter.GetBalance(&bind.CallOpts{}, source.GovernanceDapp)
	helpers.PanicErr(err)
	if fee.Int64() != 0 {
		possibleTxs = big.NewInt(0).Div(balance, fee).String()
	}
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %22s |\n", "governance dapp", source.GovernanceDapp.Hex(), balance, possibleTxs))

	balance, err = onRampRouter.GetBalance(&bind.CallOpts{}, source.PingPongDapp)
	helpers.PanicErr(err)
	if fee.Int64() != 0 {
		possibleTxs = big.NewInt(0).Div(balance, fee).String()
	}
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %22s |\n", "ping pong dapp", source.PingPongDapp.Hex(), balance, possibleTxs))

	sb.WriteString(generateSeparator(headerLengths))

	sb.WriteString(fmt.Sprintf("| %-20s | %92d |\n", "relay fee", fee))

	sb.WriteString(generateSeparator(headerLengths))

	source.Logger.Info(sb.String())
}

func printDestinationSubscriptionBalances(destination *rhea.EvmChainConfig) {
	offRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(destination.OffRampRouter, destination.Client)
	helpers.PanicErr(err)

	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Destination subscription balances for %s\n", helpers.ChainName(destination.ChainId.Int64())))

	tableHeaders := []string{"Receiver", "Address", "Balance", "Allowed senders"}
	headerLengths := []int{20, 42, 22, 44}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	subscription, err := offRampRouter.GetSubscription(&bind.CallOpts{}, destination.ReceiverDapp)
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %42v |\n", "receiver dapp", destination.ReceiverDapp.Hex(), subscription.Balance, subscription.Senders))

	subscription, err = offRampRouter.GetSubscription(&bind.CallOpts{}, destination.MessageReceiver)
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %42v |\n", "message receiver", destination.MessageReceiver.Hex(), subscription.Balance, subscription.Senders))

	subscription, err = offRampRouter.GetSubscription(&bind.CallOpts{}, destination.GovernanceDapp)
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %42v |\n", "governance dapp", destination.GovernanceDapp.Hex(), subscription.Balance, subscription.Senders))

	subscription, err = offRampRouter.GetSubscription(&bind.CallOpts{}, destination.PingPongDapp)
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %42v |\n", "ping pong dapp", destination.PingPongDapp.Hex(), subscription.Balance, subscription.Senders))

	sb.WriteString(generateSeparator(headerLengths))

	destination.Logger.Info(sb.String())
}

func printRateLimitingStatus(chain *rhea.EvmChainConfig) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Rate limits for %s\n", helpers.ChainName(chain.ChainId.Int64())))

	tableHeaders := []string{"Contract", "Tokens"}
	headerLengths := []int{25, 42}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(chain.OnRamp, chain.Client)
	helpers.PanicErr(err)
	bucketState, err := onRamp.CalculateCurrentTokenBucketState(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42d |\n", "onramp", bucketState.Tokens))

	offRamp, err := any_2_evm_subscription_offramp.NewEVM2EVMSubscriptionOffRamp(chain.OffRamp, chain.Client)
	helpers.PanicErr(err)
	offRampBucketState, err := offRamp.CalculateCurrentTokenBucketState(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42d |\n", "offramp", offRampBucketState.Tokens))

	sb.WriteString(generateSeparator(headerLengths))
	chain.Logger.Info(sb.String())
}

func printPaused(chain *rhea.EvmChainConfig) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Paused addresses for %s\n", helpers.ChainName(chain.ChainId.Int64())))

	tableHeaders := []string{"Contract", "Address", "Running"}
	headerLengths := []int{25, 42, 20}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	for _, pool := range chain.TokenPools {
		tokenPool, err := native_token_pool.NewNativeTokenPool(pool, chain.Client)
		helpers.PanicErr(err)
		paused, err := tokenPool.Paused(&bind.CallOpts{})
		helpers.PanicErr(err)

		sb.WriteString(fmt.Sprintf("| %-25s | %42s | %20t |\n", "token pool", pool.Hex(), !paused))
	}

	onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(chain.OnRamp, chain.Client)
	helpers.PanicErr(err)
	paused, err := onRamp.Paused(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42s | %20t |\n", "onramp", onRamp.Address(), !paused))

	offRamp, err := any_2_evm_subscription_offramp.NewEVM2EVMSubscriptionOffRamp(chain.OffRamp, chain.Client)
	helpers.PanicErr(err)
	paused, err = offRamp.Paused(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42s | %20t |\n", "offramp", offRamp.Address(), !paused))

	blobVerifier, err := blob_verifier.NewBlobVerifier(chain.BlobVerifier, chain.Client)
	helpers.PanicErr(err)
	paused, err = blobVerifier.Paused(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42s | %20t |\n", "blobverifier", blobVerifier.Address(), !paused))

	sb.WriteString(generateSeparator(headerLengths))
	chain.Logger.Info(sb.String())
}

func PrintNodeBalances(chain *rhea.EvmChainConfig, addresses []common.Address) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Paused addresses for %s\n", helpers.ChainName(chain.ChainId.Int64())))

	tableHeaders := []string{"Sender", "Balance"}
	headerLengths := []int{42, 18}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	for _, sender := range addresses {
		balanceAt, err := chain.Client.BalanceAt(context.Background(), sender, nil)
		helpers.PanicErr(err)

		sb.WriteString(fmt.Sprintf("| %42s |   %-16s |\n", sender.Hex(), new(big.Float).Quo(new(big.Float).SetInt(balanceAt), big.NewFloat(1e18)).String()))
	}

	sb.WriteString(generateSeparator(headerLengths))
	chain.Logger.Info(sb.String())
}

func printPoolBalances(chain *rhea.EvmChainConfig) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Pool balances for %s\n", helpers.ChainName(chain.ChainId.Int64())))

	tableHeaders := []string{"Token", "Pool", "Balance", "Onramp", "OffRamp"}
	headerLengths := []int{42, 42, 20, 10, 10}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))
	for _, pool := range chain.TokenPools {
		tokenPool, err := native_token_pool.NewNativeTokenPool(pool, chain.Client)
		helpers.PanicErr(err)

		tokenAddress, err := tokenPool.GetToken(&bind.CallOpts{})
		helpers.PanicErr(err)

		token, err := link_token_interface.NewLinkToken(tokenAddress, chain.Client)
		helpers.PanicErr(err)

		balance, err := token.BalanceOf(&bind.CallOpts{}, pool)
		helpers.PanicErr(err)

		isAllowedOnRamp, err := tokenPool.IsOnRamp(&bind.CallOpts{}, chain.OnRamp)
		helpers.PanicErr(err)

		isAllowedOffRamp, err := tokenPool.IsOffRamp(&bind.CallOpts{}, chain.OffRamp)
		helpers.PanicErr(err)

		sb.WriteString(fmt.Sprintf("| %s | %s | %20d | %10t | %10t |\n", tokenAddress.Hex(), pool.Hex(), balance, isAllowedOnRamp, isAllowedOffRamp))
	}

	sb.WriteString(generateSeparator(headerLengths))

	chain.Logger.Info(sb.String())
}

func generateHeader(headers []string, headerLengths []int) string {
	var sb strings.Builder

	sb.WriteString(generateSeparator(headerLengths))
	sb.WriteString("|")
	for i, header := range headers {
		sb.WriteString(fmt.Sprintf(" %-"+strconv.Itoa(headerLengths[i])+"s |", header))
	}
	sb.WriteString("\n")
	sb.WriteString(generateSeparator(headerLengths))

	return sb.String()
}

func generateSeparator(headerLengths []int) string {
	length := 1

	for _, headerLength := range headerLengths {
		length += headerLength + 3
	}
	return strings.Repeat("─", length) + "\n"
}

// PrintJobSpecs prints the job spec for each node and CCIP spec type, as well as a bootstrap spec.
func PrintJobSpecs(env dione.Environment, onramp, blobVerifier, offRamp common.Address, sourceChainID, destChainID *big.Int, destLinkToken common.Address, sourceReplayFrom, destReplayFrom uint64) {
	jobs := fmt.Sprintf(bootstrapTemplate+"\n", helpers.ChainName(destChainID.Int64()), blobVerifier, destChainID)
	don := dione.NewOfflineDON(env, nil)

	for i, oracle := range don.Config.Nodes {
		jobs += fmt.Sprintf("// [Node %d]\n", i)
		jobs += fmt.Sprintf(relayTemplate+"\n",
			helpers.ChainName(sourceChainID.Int64())+"-"+helpers.ChainName(destChainID.Int64()),
			blobVerifier,
			dione.GetOCRkeysForChainType(oracle.OCRKeys, "evm").ID,
			oracle.EthKeys[destChainID.String()],
			sourceChainID,
			destChainID,
			onramp,
			dione.PollPeriod,
			sourceReplayFrom,
			destReplayFrom,
			destChainID,
		)
		jobs += fmt.Sprintf(executionTemplate+"\n",
			helpers.ChainName(sourceChainID.Int64())+"-"+helpers.ChainName(destChainID.Int64()),
			offRamp,
			dione.GetOCRkeysForChainType(oracle.OCRKeys, "evm").ID,
			oracle.EthKeys[destChainID.String()],
			sourceChainID,
			destChainID,
			onramp,
			blobVerifier,
			sourceReplayFrom,
			destReplayFrom,
			destLinkToken,
			destChainID,
		)
	}
	fmt.Println(jobs)
}

const bootstrapTemplate = `
// Bootstrap Node
# BootstrapSpec
type                               = "bootstrap"
name                               = "bootstrap-%s"
relay                              = "evm"
schemaVersion                      = 1
contractID                         = "%s"
contractConfigConfirmations        = 1
contractConfigTrackerPollInterval  = "60s"
[relayConfig]
chainID                            = %s
`

const relayTemplate = `
# CCIPRelaySpec
type               = "offchainreporting2"
name               = "ccip-relay-%s"
pluginType         = "ccip-relay"
relay              = "evm"
schemaVersion      = 1
contractID         = "%s"
ocrKeyBundleID     = "%s"
transmitterID      = "%s"

[pluginConfig]
sourceChainID      = %d
destChainID        = %d
onRampIDs          = ["%s"]
pollPeriod         = "%s"
SourceStartBlock   = %d
DestStartBlock     = %d

[relayConfig]
chainID            = %d
`

const executionTemplate = `
# CCIPExecutionSpec
type              = "offchainreporting2"
name              = "ccip-exec-%s"
pluginType        = "ccip-execution"
relay             = "evm"
schemaVersion     = 1
contractID        = "%s"
ocrKeyBundleID    = "%s"
transmitterID     = "%s"

[pluginConfig]
sourceChainID     = %d
destChainID       = %d
onRampID          = "%s"
blobVerifierID    = "%s"
SourceStartBlock  = %d
DestStartBlock    = %d
tokensPerFeeCoinPipeline = """merge [type=merge left="{}" right="{\\\"%s\\\":\\\"1000000000000000000\\\"}"];"""

[relayConfig]
chainID           = %d
`
