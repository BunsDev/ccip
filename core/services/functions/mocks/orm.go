// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	functions "github.com/smartcontractkit/chainlink/v2/core/services/functions"
	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/v2/core/services/pg"

	time "time"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// CreateRequest provides a mock function with given fields: requestID, receivedAt, requestTxHash, qopts
func (_m *ORM) CreateRequest(requestID functions.RequestID, receivedAt time.Time, requestTxHash *common.Hash, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, requestID, receivedAt, requestTxHash)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(functions.RequestID, time.Time, *common.Hash, ...pg.QOpt) error); ok {
		r0 = rf(requestID, receivedAt, requestTxHash, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: requestID, qopts
func (_m *ORM) FindById(requestID functions.RequestID, qopts ...pg.QOpt) (*functions.Request, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, requestID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *functions.Request
	var r1 error
	if rf, ok := ret.Get(0).(func(functions.RequestID, ...pg.QOpt) (*functions.Request, error)); ok {
		return rf(requestID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(functions.RequestID, ...pg.QOpt) *functions.Request); ok {
		r0 = rf(requestID, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*functions.Request)
		}
	}

	if rf, ok := ret.Get(1).(func(functions.RequestID, ...pg.QOpt) error); ok {
		r1 = rf(requestID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOldestEntriesByState provides a mock function with given fields: state, limit, qopts
func (_m *ORM) FindOldestEntriesByState(state functions.RequestState, limit uint32, qopts ...pg.QOpt) ([]functions.Request, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, state, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []functions.Request
	var r1 error
	if rf, ok := ret.Get(0).(func(functions.RequestState, uint32, ...pg.QOpt) ([]functions.Request, error)); ok {
		return rf(state, limit, qopts...)
	}
	if rf, ok := ret.Get(0).(func(functions.RequestState, uint32, ...pg.QOpt) []functions.Request); ok {
		r0 = rf(state, limit, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]functions.Request)
		}
	}

	if rf, ok := ret.Get(1).(func(functions.RequestState, uint32, ...pg.QOpt) error); ok {
		r1 = rf(state, limit, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetConfirmed provides a mock function with given fields: requestID, qopts
func (_m *ORM) SetConfirmed(requestID functions.RequestID, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, requestID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(functions.RequestID, ...pg.QOpt) error); ok {
		r0 = rf(requestID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetError provides a mock function with given fields: requestID, runID, errorType, computationError, readyAt, readyForProcessing, qopts
func (_m *ORM) SetError(requestID functions.RequestID, runID int64, errorType functions.ErrType, computationError []byte, readyAt time.Time, readyForProcessing bool, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, requestID, runID, errorType, computationError, readyAt, readyForProcessing)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(functions.RequestID, int64, functions.ErrType, []byte, time.Time, bool, ...pg.QOpt) error); ok {
		r0 = rf(requestID, runID, errorType, computationError, readyAt, readyForProcessing, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetFinalized provides a mock function with given fields: requestID, reportedResult, reportedError, qopts
func (_m *ORM) SetFinalized(requestID functions.RequestID, reportedResult []byte, reportedError []byte, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, requestID, reportedResult, reportedError)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(functions.RequestID, []byte, []byte, ...pg.QOpt) error); ok {
		r0 = rf(requestID, reportedResult, reportedError, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetResult provides a mock function with given fields: requestID, runID, computationResult, readyAt, qopts
func (_m *ORM) SetResult(requestID functions.RequestID, runID int64, computationResult []byte, readyAt time.Time, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, requestID, runID, computationResult, readyAt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(functions.RequestID, int64, []byte, time.Time, ...pg.QOpt) error); ok {
		r0 = rf(requestID, runID, computationResult, readyAt, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TimeoutExpiredResults provides a mock function with given fields: cutoff, limit, qopts
func (_m *ORM) TimeoutExpiredResults(cutoff time.Time, limit uint32, qopts ...pg.QOpt) ([]functions.RequestID, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, cutoff, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []functions.RequestID
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, uint32, ...pg.QOpt) ([]functions.RequestID, error)); ok {
		return rf(cutoff, limit, qopts...)
	}
	if rf, ok := ret.Get(0).(func(time.Time, uint32, ...pg.QOpt) []functions.RequestID); ok {
		r0 = rf(cutoff, limit, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]functions.RequestID)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time, uint32, ...pg.QOpt) error); ok {
		r1 = rf(cutoff, limit, qopts...)
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
