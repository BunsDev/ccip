package main

import (
	"os"
	"testing"
)

// FullFeatureCCIP can be run as a test (prefix with Test) with the following config
// Env vars:
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
// SEED_KEY   private key used for multi-user tests. Not needed when using the "deploy" command.
// COMMAND    what function to run e.g. "deploy", "setConfig", or "externalExecution".
//
// Use "-v" as a Go tool argument for streaming log output.
func TestFullFeatureCCIP(t *testing.T) {
	ownerKey := os.Getenv("OWNER_KEY")
	command := os.Getenv("COMMAND")
	if ownerKey == "" {
		if command == "" {
			t.Log("No command given, skipping ccip-test-script. This is intended behaviour for automated testing.")
			t.SkipNow()
		}
		t.Log("Must set owner key")
		t.FailNow()
	}

	switch command {
	case "":
		t.Log("No command given, exit successfully")
		t.SkipNow()
	case "deploy":
		deployCCIPContracts(t, ownerKey,
			&Rinkeby,
			[]*EvmChainConfig{&Kovan, &BSCTestnet, &PolygonMumbai})
	case "printJobs":
		printContractConfig(GetSetupChain(t, ownerKey, Rinkeby),
			[]*EvmChainConfig{
				GetSetupChain(t, ownerKey, Kovan),
				GetSetupChain(t, ownerKey, BSCTestnet),
				GetSetupChain(t, ownerKey, PolygonMumbai)})
	default:
		runCommand(t, ownerKey, command)
	}
}

func runCommand(t *testing.T, ownerKey string, command string) {
	// The seed key is used to generate 10 keys from a single key by changing the
	// first character of the given seed with the digits 0-9
	seedKey := os.Getenv("SEED_KEY")
	if seedKey == "" {
		t.Error("must set seed key")
	}

	// Configures a client to run tests with using the network defaults and given keys.
	// After updating any contracts be sure to update the network defaults to reflect
	// those changes.
	client := NewCcipClient(t,
		// Source chain
		Rinkeby,
		// Dest chain
		Kovan,
		ownerKey,
		seedKey,
	)

	client.Source.Client.AssureHealth(t)
	client.Dest.Client.AssureHealth(t)
	client.UnpauseAll()

	switch command {
	case "setConfig":
		// Set the config to the message executor and the offramp
		client.SetConfig()
	case "externalExecution":
		// Cross chain request with the client manually proving and executing the transaction
		client.ExternalExecutionHappyPath(t)
	case "noRepeat":
		// Executing the same request twice should fail
		client.ExternalExecutionSubmitOfframpTwiceShouldFail(t)
	case "don":
		// Cross chain request with DON execution
		client.DonExecutionHappyPath(t)
	case "batching":
		// Submit 10 txs. This should result in the txs being batched together
		client.ScalingAndBatching(t)
	case "exceedBucket":
		// Should not be able to send funds greater than the amount in the bucket
		client.NotEnoughFundsInBucketShouldFail(t)
	case "tryPausedPool":
		// Should fail because the pool is paused
		client.TryGetTokensFromPausedPool()
	case "tryPausedOfframp":
		// Should not be included in a report because the offramp is paused
		client.CrossChainSendPausedOfframpShouldFail(t)
	case "tryPausedOnramp":
		// Should not succeed because the onramp is paused
		client.CrossChainSendPausedOnrampShouldFail(t)
	default:
		t.Errorf("Unknown command \"%s\"", command)
	}
}
