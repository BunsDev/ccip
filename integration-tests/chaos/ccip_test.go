package chaos_test

import (
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink-env/chaos"
	a "github.com/smartcontractkit/chainlink-env/pkg/alias"
	"github.com/smartcontractkit/chainlink-testing-framework/utils"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/testsetups"
)

/* @network-chaos and @pod-chaos are split intentionally into 2 parallel groups
we can't use chaos.NewNetworkPartition and chaos.NewFailPods in parallel
because of jsii runtime bug, see Makefile and please use those targets to run tests
In .github/workflows/ccip-chaos-tests.yml we use these tags to run these tests separately
*/

func TestChaosCCIP(t *testing.T) {
	inputs := []struct {
		testName             string
		chaosFunc            chaos.ManifestFunc
		chaosProps           *chaos.Props
		waitForChaosRecovery bool
	}{
		{
			testName:  "CCIP works after rpc is down for NetworkA @network-chaos",
			chaosFunc: chaos.NewNetworkPartition,
			chaosProps: &chaos.Props{
				FromLabels: &map[string]*string{actions.ChaosGroupNetworkACCIPGeth: a.Str("1")},
				// chainlink-0 is default label set for all cll nodes
				ToLabels:    &map[string]*string{"app": a.Str("chainlink-0")},
				DurationStr: "1m",
			},
			waitForChaosRecovery: true,
		},
		{
			testName:  "CCIP works after rpc is down for NetworkB @network-chaos",
			chaosFunc: chaos.NewNetworkPartition,
			chaosProps: &chaos.Props{
				FromLabels:  &map[string]*string{actions.ChaosGroupNetworkBCCIPGeth: a.Str("1")},
				ToLabels:    &map[string]*string{"app": a.Str("chainlink-0")},
				DurationStr: "1m",
			},
			waitForChaosRecovery: true,
		},
		{
			testName:  "CCIP works after 2 rpc's are down for all cll nodes @network-chaos",
			chaosFunc: chaos.NewNetworkPartition,
			chaosProps: &chaos.Props{
				FromLabels:  &map[string]*string{"geth": a.Str(actions.ChaosGroupCCIPGeth)},
				ToLabels:    &map[string]*string{"app": a.Str("chainlink-0")},
				DurationStr: "1m",
			},
			waitForChaosRecovery: true,
		},
		{
			testName:  "CCIP Commit works after majority of CL nodes are recovered from pod failure @pod-chaos",
			chaosFunc: chaos.NewFailPods,
			chaosProps: &chaos.Props{
				LabelsSelector: &map[string]*string{actions.ChaosGroupCommitFaultyPlus: a.Str("1")},
				DurationStr:    "1m",
			},
			waitForChaosRecovery: true,
		},
		{
			testName:  "CCIP Execution works after majority of CL nodes are recovered from pod failure @pod-chaos",
			chaosFunc: chaos.NewFailPods,
			chaosProps: &chaos.Props{
				LabelsSelector: &map[string]*string{actions.ChaosGroupExecutionFaultyPlus: a.Str("1")},
				DurationStr:    "1m",
			},
			waitForChaosRecovery: true,
		},
		{
			testName:  "CCIP Commit works while minority of CL nodes are in failed state for pod failure @pod-chaos",
			chaosFunc: chaos.NewFailPods,
			chaosProps: &chaos.Props{
				LabelsSelector: &map[string]*string{actions.ChaosGroupCommitFaulty: a.Str("1")},
				DurationStr:    "90s",
			},
			waitForChaosRecovery: false,
		},
		{
			testName:  "CCIP Execution works while minority of CL nodes are in failed state for pod failure @pod-chaos",
			chaosFunc: chaos.NewFailPods,
			chaosProps: &chaos.Props{
				LabelsSelector: &map[string]*string{actions.ChaosGroupExecutionFaulty: a.Str("1")},
				DurationStr:    "90s",
			},
			waitForChaosRecovery: false,
		},
	}

	for _, in := range inputs {
		in := in
		t.Run(in.testName, func(t *testing.T) {
			t.Parallel()
			l := utils.GetTestLogger(t)
			testCfg := testsetups.NewCCIPTestConfig(t, l, testsetups.Chaos)
			var (
				numOfCommitNodes = 5
				numOfRequests    = 3
			)

			setUpArgs := testsetups.CCIPDefaultTestSetUp(t, l, "chaos-ccip", map[string]interface{}{
				"replicas": "12",
				"db": map[string]interface{}{
					"stateful": true,
					"capacity": "10Gi",
					"resources": map[string]interface{}{
						"requests": map[string]interface{}{
							"cpu":    "250m",
							"memory": "256Mi",
						},
						"limits": map[string]interface{}{
							"cpu":    "250m",
							"memory": "256Mi",
						},
					},
				},
			}, []*big.Int{big.NewInt(1e8)}, numOfCommitNodes, false, false, testCfg)

			if len(setUpArgs.Lanes) == 0 {
				return
			}

			lane := setUpArgs.Lanes[0].ForwardLane

			tearDown := setUpArgs.TearDown
			testEnvironment := lane.TestEnv.K8Env
			testSetup := lane.TestEnv

			testSetup.ChaosLabel(t, lane.SourceChain.GetNetworkName(), lane.DestChain.GetNetworkName())

			// apply chaos
			chaosId, err := testEnvironment.Chaos.Run(in.chaosFunc(testEnvironment.Cfg.Namespace, in.chaosProps))
			require.NoError(t, err)
			t.Cleanup(func() {
				if chaosId != "" {
					testEnvironment.Chaos.Stop(chaosId)
				}
				tearDown()
			})
			lane.RecordStateBeforeTransfer()
			// Send the ccip-request while the chaos is at play
			_, err = lane.SendRequests(numOfRequests)
			require.NoError(t, err)
			if in.waitForChaosRecovery {
				// wait for chaos to be recovered before further validation
				testEnvironment.Chaos.WaitForAllRecovered(chaosId)
			} else {
				l.Info().Msg("proceeding without waiting for chaos recovery")
			}
			lane.ValidateRequests()
		})
	}
}
