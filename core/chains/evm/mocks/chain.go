// Code generated by mockery v2.10.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	client "github.com/smartcontractkit/chainlink/core/chains/evm/client"
	config "github.com/smartcontractkit/chainlink/core/chains/evm/config"

	context "context"

	log "github.com/smartcontractkit/chainlink/core/chains/evm/log"

	logger "github.com/smartcontractkit/chainlink/core/logger"

	logpoller "github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"

	mock "github.com/stretchr/testify/mock"

	monitor "github.com/smartcontractkit/chainlink/core/chains/evm/monitor"

	txmgr "github.com/smartcontractkit/chainlink/core/chains/evm/txmgr"

	types "github.com/smartcontractkit/chainlink/core/chains/evm/headtracker/types"
)

// Chain is an autogenerated mock type for the Chain type
type Chain struct {
	mock.Mock
}

// BalanceMonitor provides a mock function with given fields:
func (_m *Chain) BalanceMonitor() monitor.BalanceMonitor {
	ret := _m.Called()

	var r0 monitor.BalanceMonitor
	if rf, ok := ret.Get(0).(func() monitor.BalanceMonitor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(monitor.BalanceMonitor)
		}
	}

	return r0
}

// Client provides a mock function with given fields:
func (_m *Chain) Client() client.Client {
	ret := _m.Called()

	var r0 client.Client
	if rf, ok := ret.Get(0).(func() client.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.Client)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Chain) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Config provides a mock function with given fields:
func (_m *Chain) Config() config.ChainScopedConfig {
	ret := _m.Called()

	var r0 config.ChainScopedConfig
	if rf, ok := ret.Get(0).(func() config.ChainScopedConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.ChainScopedConfig)
		}
	}

	return r0
}

// HeadBroadcaster provides a mock function with given fields:
func (_m *Chain) HeadBroadcaster() types.HeadBroadcaster {
	ret := _m.Called()

	var r0 types.HeadBroadcaster
	if rf, ok := ret.Get(0).(func() types.HeadBroadcaster); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.HeadBroadcaster)
		}
	}

	return r0
}

// HeadTracker provides a mock function with given fields:
func (_m *Chain) HeadTracker() types.HeadTracker {
	ret := _m.Called()

	var r0 types.HeadTracker
	if rf, ok := ret.Get(0).(func() types.HeadTracker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.HeadTracker)
		}
	}

	return r0
}

// Healthy provides a mock function with given fields:
func (_m *Chain) Healthy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Chain) ID() *big.Int {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// LogBroadcaster provides a mock function with given fields:
func (_m *Chain) LogBroadcaster() log.Broadcaster {
	ret := _m.Called()

	var r0 log.Broadcaster
	if rf, ok := ret.Get(0).(func() log.Broadcaster); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Broadcaster)
		}
	}

	return r0
}

// LogPoller provides a mock function with given fields:
func (_m *Chain) LogPoller() logpoller.LogPoller {
	ret := _m.Called()

	var r0 logpoller.LogPoller
	if rf, ok := ret.Get(0).(func() logpoller.LogPoller); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logpoller.LogPoller)
		}
	}

	return r0
}

// Logger provides a mock function with given fields:
func (_m *Chain) Logger() logger.Logger {
	ret := _m.Called()

	var r0 logger.Logger
	if rf, ok := ret.Get(0).(func() logger.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.Logger)
		}
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *Chain) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *Chain) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxManager provides a mock function with given fields:
func (_m *Chain) TxManager() txmgr.TxManager {
	ret := _m.Called()

	var r0 txmgr.TxManager
	if rf, ok := ret.Get(0).(func() txmgr.TxManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(txmgr.TxManager)
		}
	}

	return r0
}
