package dione

import (
	"fmt"
	"math/big"
	"net/http"
	"net/url"

	"github.com/smartcontractkit/chainlink/integration-tests/client"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	"github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/job"
)

const (
	PollPeriod = "1s"
)

type Environment string

const (
	StagingAlpha Environment = "staging-alpha"
	StagingBeta  Environment = "staging-beta"
	Production   Environment = "prod"
)

type JobType string

const (
	Relay     JobType = "relay"
	Execution JobType = "exec"
	Boostrap  JobType = "bootstrap"
)

type Chain string

const (
	Rinkeby        Chain = "Rinkeby"
	Goerli         Chain = "Goerli"
	OptimismGoerli Chain = "420"
	Sepolia        Chain = "Sepolia"
	AvaxFuji       Chain = "Avax Fuji"
)

type ChainConfig struct {
	ChainID  uint64
	RpcUrl   string
	EIP1559  bool
	GasPrice uint64
}

type NodesConfig struct {
	Bootstrap NodeConfig
	Nodes     []NodeConfig
}

type NodeConfig struct {
	EthKeys map[string]string
	PeerID  string
	OCRKeys client.OCR2Keys
}

type DON struct {
	Nodes     []*client.Chainlink
	bootstrap *client.Chainlink
	OfflineDON
}

func NewDON(env Environment, lggr logger.Logger) DON {
	creds, err := ReadCredentials(env)
	common.PanicErr(err)
	nodes, bootstrap, err := creds.DialNodes()
	common.PanicErr(err)

	return DON{
		Nodes:      nodes,
		bootstrap:  bootstrap,
		OfflineDON: NewOfflineDON(env, lggr),
	}
}

func (don *DON) PopulateOCR2Keys() {
	for i, node := range don.Nodes {
		keys, _, err := node.ReadOCR2Keys()
		common.PanicErr(err)
		don.Config.Nodes[i].OCRKeys = *keys
	}
}

func createKey(c *client.Chainlink, chain string) (*http.Response, error) {
	createUrl := url.URL{
		Path: "/v2/keys/evm",
	}
	query := createUrl.Query()
	query.Set("evmChainID", chain)

	createUrl.RawQuery = query.Encode()
	resp, err := c.APIClient.R().Post(createUrl.String())
	if err != nil {
		return nil, err
	}

	return resp.RawResponse, nil
}

func deleteKnownETHKey(node *client.Chainlink, key string) (*http.Response, error) {
	deleteUrl := url.URL{
		Path: "/v2/keys/evm/" + key,
	}
	query := deleteUrl.Query()
	query.Set("hard", "true")
	deleteUrl.RawQuery = query.Encode()

	resp, err := node.APIClient.R().
		Delete(deleteUrl.String())
	if err != nil {
		return nil, err
	}
	return resp.RawResponse, err
}

func (don *DON) DeleteKnownKey(chainID string) {
	for i, node := range don.Nodes {
		// Only remove a key if it exists
		if key, ok := don.Config.Nodes[i].EthKeys[chainID]; ok {
			resp, err := deleteKnownETHKey(node, key)
			if err != nil {
				don.lggr.Infof("Failed to delete key: %s", resp.Status)
			}
		}
	}
}

func (don *DON) CreateNewEthKeysForChain(chainID *big.Int) {
	for i, node := range don.Nodes {
		_, err := createKey(node, chainID.String())
		common.PanicErr(err)
		don.lggr.Infof("Node [%2d] Created new eth key", i)
	}
}

func (don *DON) PopulatePeerId() {
	for i, node := range don.Nodes {
		p2pkeys, err := node.MustReadP2PKeys()
		common.PanicErr(err)

		don.Config.Nodes[i].PeerID = p2pkeys.Data[0].Attributes.PeerID
	}

	p2pkeys, err := don.bootstrap.MustReadP2PKeys()
	common.PanicErr(err)
	don.Config.Bootstrap.PeerID = p2pkeys.Data[0].Attributes.PeerID
}

func (don *DON) PopulateEthKeys() {
	for i, node := range don.Nodes {
		keys, err := node.MustReadETHKeys()
		if err != nil {
			don.lggr.Infof("Failed getting keys for node %d", i)
		}

		don.Config.Nodes[i].EthKeys = make(map[string]string)
		don.lggr.Infof("Read %d keys for node %2d", len(keys.Data), i)
		for _, key := range keys.Data {
			don.Config.Nodes[i].EthKeys[key.Attributes.ChainID] = key.Attributes.Address
		}
	}
}

func (don *DON) ClearJobSpecs(jobType JobType, source Chain, destination Chain) {
	jobToDelete := fmt.Sprintf("ccip-%s-%s-%s", jobType, source, destination)

	for i, node := range don.Nodes {
		jobs, _, err := node.ReadJobs()
		common.PanicErr(err)

		for _, maps := range jobs.Data {
			jb := maps["attributes"].(map[string]interface{})
			jobName := jb["name"].(string)
			id := maps["id"].(string)

			don.lggr.Infof("Node [%2d]: Job %s: %s", i, id, jobName)

			if jobToDelete == jobName {
				don.lggr.Infof("Node [%2d]:Deleting job %s: %s", i, id, jobName)
				s, err := node.DeleteJob(id)
				common.PanicErr(err)
				don.lggr.Infof(s.Status)
			}
		}
	}
}

func (don *DON) ListJobSpecs() {
	for i, node := range don.Nodes {
		jobs, _, err := node.ReadJobs()
		common.PanicErr(err)

		for _, maps := range jobs.Data {
			jb := maps["attributes"].(map[string]interface{})
			jobName := jb["name"].(string)
			id := maps["id"].(string)

			don.lggr.Infof("Node [%2d]: Job %3s: %-28s %+v", i, id, jobName, jb)
		}
	}
}

func (don *DON) AddRawJobSpec(node *client.Chainlink, spec string) {
	jb, tx, err := node.CreateJobRaw(spec)
	common.PanicErr(err)

	don.lggr.Infof("Created job %3s. Status code %s", jb.Data.ID, tx.Status)
}

func (don *DON) LoadCurrentNodeParams() {
	don.PopulateOCR2Keys()
	don.PopulateEthKeys()
	don.PopulatePeerId()
	don.PrintConfig()
}

func (don *DON) ClearAllJobs(chainA Chain, chainB Chain) {
	don.ClearJobSpecs(Relay, chainA, chainB)
	don.ClearJobSpecs(Execution, chainA, chainB)
	don.ClearJobSpecs(Relay, chainB, chainA)
	don.ClearJobSpecs(Execution, chainB, chainA)
}

func (don *DON) AddTwoWaySpecs(chainA rhea.EvmChainConfig, chainB rhea.EvmChainConfig) {
	relaySpecAB := generateRelayJobSpecs(&chainA, &chainB)
	don.AddJobSpecs(relaySpecAB)
	executionSpecAB := generateExecutionJobSpecs(&chainA, &chainB)
	don.AddJobSpecs(executionSpecAB)
	relaySpecBA := generateRelayJobSpecs(&chainB, &chainA)
	don.AddJobSpecs(relaySpecBA)
	executionSpecBA := generateExecutionJobSpecs(&chainB, &chainA)
	don.AddJobSpecs(executionSpecBA)
}

func (don *DON) AddJobSpecs(spec job.Job) {
	chainID := spec.OCR2OracleSpec.RelayConfig["chainID"].(string)

	for i, node := range don.Nodes {
		evmKeyBundle := GetOCRkeysForChainType(don.Config.Nodes[i].OCRKeys, "evm")
		transmitterIDs := don.Config.Nodes[i].EthKeys

		spec.OCR2OracleSpec.OCRKeyBundleID.SetValid(evmKeyBundle.ID)
		spec.OCR2OracleSpec.TransmitterID.SetValid(transmitterIDs[chainID])

		var specString string
		if spec.OCR2OracleSpec.PluginType == job.CCIPRelay {
			specString = RelaySpecToString(spec)
		} else {
			specString = ExecSpecToString(spec)
		}
		don.lggr.Infof(specString)
		don.AddRawJobSpec(node, specString)
	}
}
