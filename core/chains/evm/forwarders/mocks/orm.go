// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	forwarders "github.com/smartcontractkit/chainlink/v2/core/chains/evm/forwarders"
	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/v2/core/services/pg"

	utils "github.com/smartcontractkit/chainlink/v2/core/utils"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// CreateForwarder provides a mock function with given fields: addr, evmChainId
func (_m *ORM) CreateForwarder(addr common.Address, evmChainId utils.Big) (forwarders.Forwarder, error) {
	ret := _m.Called(addr, evmChainId)

	var r0 forwarders.Forwarder
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, utils.Big) (forwarders.Forwarder, error)); ok {
		return rf(addr, evmChainId)
	}
	if rf, ok := ret.Get(0).(func(common.Address, utils.Big) forwarders.Forwarder); ok {
		r0 = rf(addr, evmChainId)
	} else {
		r0 = ret.Get(0).(forwarders.Forwarder)
	}

	if rf, ok := ret.Get(1).(func(common.Address, utils.Big) error); ok {
		r1 = rf(addr, evmChainId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteForwarder provides a mock function with given fields: id, cleanup
func (_m *ORM) DeleteForwarder(id int64, cleanup func(pg.Queryer, int64, common.Address) error) error {
	ret := _m.Called(id, cleanup)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, func(pg.Queryer, int64, common.Address) error) error); ok {
		r0 = rf(id, cleanup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindForwarders provides a mock function with given fields: offset, limit
func (_m *ORM) FindForwarders(offset int, limit int) ([]forwarders.Forwarder, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []forwarders.Forwarder
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]forwarders.Forwarder, int, error)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []forwarders.Forwarder); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]forwarders.Forwarder)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindForwardersByChain provides a mock function with given fields: evmChainId
func (_m *ORM) FindForwardersByChain(evmChainId utils.Big) ([]forwarders.Forwarder, error) {
	ret := _m.Called(evmChainId)

	var r0 []forwarders.Forwarder
	var r1 error
	if rf, ok := ret.Get(0).(func(utils.Big) ([]forwarders.Forwarder, error)); ok {
		return rf(evmChainId)
	}
	if rf, ok := ret.Get(0).(func(utils.Big) []forwarders.Forwarder); ok {
		r0 = rf(evmChainId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]forwarders.Forwarder)
		}
	}

	if rf, ok := ret.Get(1).(func(utils.Big) error); ok {
		r1 = rf(evmChainId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindForwardersInListByChain provides a mock function with given fields: evmChainId, addrs
func (_m *ORM) FindForwardersInListByChain(evmChainId utils.Big, addrs []common.Address) ([]forwarders.Forwarder, error) {
	ret := _m.Called(evmChainId, addrs)

	var r0 []forwarders.Forwarder
	var r1 error
	if rf, ok := ret.Get(0).(func(utils.Big, []common.Address) ([]forwarders.Forwarder, error)); ok {
		return rf(evmChainId, addrs)
	}
	if rf, ok := ret.Get(0).(func(utils.Big, []common.Address) []forwarders.Forwarder); ok {
		r0 = rf(evmChainId, addrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]forwarders.Forwarder)
		}
	}

	if rf, ok := ret.Get(1).(func(utils.Big, []common.Address) error); ok {
		r1 = rf(evmChainId, addrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewORM interface {
	mock.TestingT
	Cleanup(func())
}

// NewORM creates a new instance of ORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewORM(t mockConstructorTestingTNewORM) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
