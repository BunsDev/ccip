package chaos_test

import (
	"math/big"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-env/chaos"
	"github.com/smartcontractkit/chainlink-env/environment"
	a "github.com/smartcontractkit/chainlink-env/pkg/alias"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
)

func TestChaosCCIP(t *testing.T) {
	var (
		tearDown         func()
		numOfCommitNodes = 5
		numOfRequests    = 3
		testEnvironment  *environment.Environment
		lane             *actions.CCIPLane
		testSetup        actions.CCIPTestEnv
	)
	t.Cleanup(func() {
		tearDown()
	})

	lane, _, tearDown = actions.CCIPDefaultTestSetUp(t, "chaos-ccip",
		map[string]interface{}{
			"replicas": "12",
			"toml":     actions.DefaultCCIPCLNodeEnv(t),
			"env": map[string]interface{}{
				"CL_DEV": "true",
			},
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
		}, []*big.Int{big.NewInt(1e8)}, numOfCommitNodes, false, false)
	require.NoError(t, lane.IsLaneDeployed())
	testEnvironment = lane.TestEnv.K8Env
	testSetup = *lane.TestEnv

	inputs := []struct {
		chaosFunc            chaos.ManifestFunc
		chaosProps           *chaos.Props
		waitForChaosRecovery bool
	}{
		{
			chaosFunc: chaos.NewFailPods,
			chaosProps: &chaos.Props{
				LabelsSelector: &map[string]*string{actions.ChaosGroupCommitFaultyPlus: a.Str("1")},
				DurationStr:    "1m",
			},
			waitForChaosRecovery: true,
		},
		{
			chaosFunc: chaos.NewFailPods,
			chaosProps: &chaos.Props{
				LabelsSelector: &map[string]*string{actions.ChaosGroupExecutionFaultyPlus: a.Str("1")},
				DurationStr:    "1m",
			},
			waitForChaosRecovery: true,
		},
		{
			chaosFunc: chaos.NewFailPods,
			chaosProps: &chaos.Props{
				LabelsSelector: &map[string]*string{actions.ChaosGroupCommitFaulty: a.Str("1")},
				DurationStr:    "90s",
			},
			waitForChaosRecovery: false,
		},
		{
			chaosFunc: chaos.NewFailPods,
			chaosProps: &chaos.Props{
				LabelsSelector: &map[string]*string{actions.ChaosGroupExecutionFaulty: a.Str("1")},
				DurationStr:    "90s",
			},
			waitForChaosRecovery: false,
		},
	}
	for _, in := range inputs {
		t.Run("", func(t *testing.T) {
			testSetup.ChaosLabel(t)

			// apply chaos
			chaosId, err := testEnvironment.Chaos.Run(in.chaosFunc(testEnvironment.Cfg.Namespace, in.chaosProps))
			require.NoError(t, err)
			// Send the ccip-request while the chaos is at play
			lane.SendGERequests(numOfRequests)
			if in.waitForChaosRecovery {
				// wait for chaos to be recovered before further validation
				testEnvironment.Chaos.WaitForAllRecovered(chaosId)
			} else {
				log.Info().Msg("proceeding without waiting for chaos recovery")
			}
			lane.ValidateGERequests()
		})
	}
}
