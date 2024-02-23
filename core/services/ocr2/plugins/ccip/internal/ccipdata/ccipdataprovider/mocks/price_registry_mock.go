// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cciptypes "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciptypes"

	mock "github.com/stretchr/testify/mock"
)

// PriceRegistry is an autogenerated mock type for the PriceRegistry type
type PriceRegistry struct {
	mock.Mock
}

// NewPriceRegistryReader provides a mock function with given fields: ctx, addr
func (_m *PriceRegistry) NewPriceRegistryReader(ctx context.Context, addr cciptypes.Address) (cciptypes.PriceRegistryReader, error) {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for NewPriceRegistryReader")
	}

	var r0 cciptypes.PriceRegistryReader
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cciptypes.Address) (cciptypes.PriceRegistryReader, error)); ok {
		return rf(ctx, addr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cciptypes.Address) cciptypes.PriceRegistryReader); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cciptypes.PriceRegistryReader)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cciptypes.Address) error); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPriceRegistry creates a new instance of PriceRegistry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPriceRegistry(t interface {
	mock.TestingT
	Cleanup(func())
}) *PriceRegistry {
	mock := &PriceRegistry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
