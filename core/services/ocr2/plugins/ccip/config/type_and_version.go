package config

import (
	"strings"

	"github.com/Masterminds/semver/v3"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	type_and_version "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/type_and_version_interface_wrapper"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

type ContractType string

var (
	EVM2EVMOnRamp  ContractType = "EVM2EVMOnRamp"
	EVM2EVMOffRamp ContractType = "EVM2EVMOffRamp"
	CommitStore    ContractType = "CommitStore"
	PriceRegistry  ContractType = "PriceRegistry"
	ContractTypes               = mapset.NewSet[ContractType](
		EVM2EVMOffRamp,
		EVM2EVMOnRamp,
		CommitStore,
		PriceRegistry,
	)
)

func VerifyTypeAndVersion(addr common.Address, client bind.ContractBackend, expectedType ContractType) (semver.Version, error) {
	contractType, version, err := TypeAndVersion(addr, client)
	if err != nil {
		return semver.Version{}, errors.Errorf("failed getting type and version %v", err)
	}
	if contractType != expectedType {
		return semver.Version{}, errors.Errorf("Wrong contract type %s", contractType)
	}
	return version, nil
}

func TypeAndVersion(addr common.Address, client bind.ContractBackend) (ContractType, semver.Version, error) {
	var log logger.Logger

	tv, err := type_and_version.NewTypeAndVersionInterface(addr, client)
	if err != nil {
		return "", semver.Version{}, err
	}

	log.Infof("New Type and version interface initiated.... %s", tv.Address().String())

	tvStr, err := tv.TypeAndVersion(nil)
	if err != nil {
		return "", semver.Version{}, errors.Errorf("Error while retrieving type and version %v %s", err, addr.String())
	}

	contractType, versionStr, err := ParseTypeAndVersion(tvStr)
	if err != nil {
		return "", semver.Version{}, err
	}
	v, err := semver.NewVersion(versionStr)
	if err != nil {
		return "", semver.Version{}, errors.Wrapf(err, "failed parsing version %s", versionStr)
	}

	if !ContractTypes.Contains(ContractType(contractType)) {
		return "", semver.Version{}, errors.Errorf("unrecognized contract type %v", contractType)
	}
	return ContractType(contractType), *v, nil
}

func ParseTypeAndVersion(tvStr string) (string, string, error) {
	typeAndVersionValues := strings.Split(tvStr, " ")

	if len(typeAndVersionValues) < 2 {
		return "", "", errors.Errorf("invalid type and version %s", tvStr)
	}
	return typeAndVersionValues[0], typeAndVersionValues[1], nil
}
