// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	liquiditygraph "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditygraph"

	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

// Rebalancer is an autogenerated mock type for the Rebalancer type
type Rebalancer struct {
	mock.Mock
}

// ComputeTransfersToBalance provides a mock function with given fields: g, inflightTransfers, medianLiquidityPerChain
func (_m *Rebalancer) ComputeTransfersToBalance(g liquiditygraph.LiquidityGraph, inflightTransfers []models.PendingTransfer, medianLiquidityPerChain []models.NetworkLiquidity) ([]models.Transfer, error) {
	ret := _m.Called(g, inflightTransfers, medianLiquidityPerChain)

	if len(ret) == 0 {
		panic("no return value specified for ComputeTransfersToBalance")
	}

	var r0 []models.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(liquiditygraph.LiquidityGraph, []models.PendingTransfer, []models.NetworkLiquidity) ([]models.Transfer, error)); ok {
		return rf(g, inflightTransfers, medianLiquidityPerChain)
	}
	if rf, ok := ret.Get(0).(func(liquiditygraph.LiquidityGraph, []models.PendingTransfer, []models.NetworkLiquidity) []models.Transfer); ok {
		r0 = rf(g, inflightTransfers, medianLiquidityPerChain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transfer)
		}
	}

	if rf, ok := ret.Get(1).(func(liquiditygraph.LiquidityGraph, []models.PendingTransfer, []models.NetworkLiquidity) error); ok {
		r1 = rf(g, inflightTransfers, medianLiquidityPerChain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRebalancer creates a new instance of Rebalancer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRebalancer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Rebalancer {
	mock := &Rebalancer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
