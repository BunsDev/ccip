// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"

	event "github.com/ethereum/go-ethereum/event"

	forwarder "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/forwarder"

	generated "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// ForwarderInterface is an autogenerated mock type for the ForwarderInterface type
type ForwarderInterface struct {
	mock.Mock
}

// AcceptOwnership provides a mock function with given fields: opts
func (_m *ForwarderInterface) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(opts)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) (*types.Transaction, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Address provides a mock function with given fields:
func (_m *ForwarderInterface) Address() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// EIP712DOMAINTYPE provides a mock function with given fields: opts
func (_m *ForwarderInterface) EIP712DOMAINTYPE(opts *bind.CallOpts) (string, error) {
	ret := _m.Called(opts)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (string, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) string); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Execute provides a mock function with given fields: opts, req, domainSeparator, requestTypeHash, suffixData, sig
func (_m *ForwarderInterface) Execute(opts *bind.TransactOpts, req forwarder.IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, req, domainSeparator, requestTypeHash, suffixData, sig)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, forwarder.IForwarderForwardRequest, [32]byte, [32]byte, []byte, []byte) (*types.Transaction, error)); ok {
		return rf(opts, req, domainSeparator, requestTypeHash, suffixData, sig)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, forwarder.IForwarderForwardRequest, [32]byte, [32]byte, []byte, []byte) *types.Transaction); ok {
		r0 = rf(opts, req, domainSeparator, requestTypeHash, suffixData, sig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, forwarder.IForwarderForwardRequest, [32]byte, [32]byte, []byte, []byte) error); ok {
		r1 = rf(opts, req, domainSeparator, requestTypeHash, suffixData, sig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterDomainRegistered provides a mock function with given fields: opts, domainSeparator
func (_m *ForwarderInterface) FilterDomainRegistered(opts *bind.FilterOpts, domainSeparator [][32]byte) (*forwarder.ForwarderDomainRegisteredIterator, error) {
	ret := _m.Called(opts, domainSeparator)

	var r0 *forwarder.ForwarderDomainRegisteredIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, [][32]byte) (*forwarder.ForwarderDomainRegisteredIterator, error)); ok {
		return rf(opts, domainSeparator)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, [][32]byte) *forwarder.ForwarderDomainRegisteredIterator); ok {
		r0 = rf(opts, domainSeparator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*forwarder.ForwarderDomainRegisteredIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, [][32]byte) error); ok {
		r1 = rf(opts, domainSeparator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterForwardSucceeded provides a mock function with given fields: opts, from, target, domainSeparator
func (_m *ForwarderInterface) FilterForwardSucceeded(opts *bind.FilterOpts, from []common.Address, target []common.Address, domainSeparator [][32]byte) (*forwarder.ForwarderForwardSucceededIterator, error) {
	ret := _m.Called(opts, from, target, domainSeparator)

	var r0 *forwarder.ForwarderForwardSucceededIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address, [][32]byte) (*forwarder.ForwarderForwardSucceededIterator, error)); ok {
		return rf(opts, from, target, domainSeparator)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address, [][32]byte) *forwarder.ForwarderForwardSucceededIterator); ok {
		r0 = rf(opts, from, target, domainSeparator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*forwarder.ForwarderForwardSucceededIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address, [][32]byte) error); ok {
		r1 = rf(opts, from, target, domainSeparator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterOwnershipTransferRequested provides a mock function with given fields: opts, from, to
func (_m *ForwarderInterface) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*forwarder.ForwarderOwnershipTransferRequestedIterator, error) {
	ret := _m.Called(opts, from, to)

	var r0 *forwarder.ForwarderOwnershipTransferRequestedIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) (*forwarder.ForwarderOwnershipTransferRequestedIterator, error)); ok {
		return rf(opts, from, to)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) *forwarder.ForwarderOwnershipTransferRequestedIterator); ok {
		r0 = rf(opts, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*forwarder.ForwarderOwnershipTransferRequestedIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterOwnershipTransferred provides a mock function with given fields: opts, from, to
func (_m *ForwarderInterface) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*forwarder.ForwarderOwnershipTransferredIterator, error) {
	ret := _m.Called(opts, from, to)

	var r0 *forwarder.ForwarderOwnershipTransferredIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) (*forwarder.ForwarderOwnershipTransferredIterator, error)); ok {
		return rf(opts, from, to)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) *forwarder.ForwarderOwnershipTransferredIterator); ok {
		r0 = rf(opts, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*forwarder.ForwarderOwnershipTransferredIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterRequestTypeRegistered provides a mock function with given fields: opts, typeHash
func (_m *ForwarderInterface) FilterRequestTypeRegistered(opts *bind.FilterOpts, typeHash [][32]byte) (*forwarder.ForwarderRequestTypeRegisteredIterator, error) {
	ret := _m.Called(opts, typeHash)

	var r0 *forwarder.ForwarderRequestTypeRegisteredIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, [][32]byte) (*forwarder.ForwarderRequestTypeRegisteredIterator, error)); ok {
		return rf(opts, typeHash)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, [][32]byte) *forwarder.ForwarderRequestTypeRegisteredIterator); ok {
		r0 = rf(opts, typeHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*forwarder.ForwarderRequestTypeRegisteredIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, [][32]byte) error); ok {
		r1 = rf(opts, typeHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GENERICPARAMS provides a mock function with given fields: opts
func (_m *ForwarderInterface) GENERICPARAMS(opts *bind.CallOpts) (string, error) {
	ret := _m.Called(opts)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (string, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) string); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDomainSeparator provides a mock function with given fields: opts, name, version
func (_m *ForwarderInterface) GetDomainSeparator(opts *bind.CallOpts, name string, version string) ([]byte, error) {
	ret := _m.Called(opts, name, version)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, string, string) ([]byte, error)); ok {
		return rf(opts, name, version)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, string, string) []byte); ok {
		r0 = rf(opts, name, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, string, string) error); ok {
		r1 = rf(opts, name, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEncoded provides a mock function with given fields: opts, req, requestTypeHash, suffixData
func (_m *ForwarderInterface) GetEncoded(opts *bind.CallOpts, req forwarder.IForwarderForwardRequest, requestTypeHash [32]byte, suffixData []byte) ([]byte, error) {
	ret := _m.Called(opts, req, requestTypeHash, suffixData)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, forwarder.IForwarderForwardRequest, [32]byte, []byte) ([]byte, error)); ok {
		return rf(opts, req, requestTypeHash, suffixData)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, forwarder.IForwarderForwardRequest, [32]byte, []byte) []byte); ok {
		r0 = rf(opts, req, requestTypeHash, suffixData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, forwarder.IForwarderForwardRequest, [32]byte, []byte) error); ok {
		r1 = rf(opts, req, requestTypeHash, suffixData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Owner provides a mock function with given fields: opts
func (_m *ForwarderInterface) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (common.Address, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseDomainRegistered provides a mock function with given fields: log
func (_m *ForwarderInterface) ParseDomainRegistered(log types.Log) (*forwarder.ForwarderDomainRegistered, error) {
	ret := _m.Called(log)

	var r0 *forwarder.ForwarderDomainRegistered
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*forwarder.ForwarderDomainRegistered, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *forwarder.ForwarderDomainRegistered); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*forwarder.ForwarderDomainRegistered)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseForwardSucceeded provides a mock function with given fields: log
func (_m *ForwarderInterface) ParseForwardSucceeded(log types.Log) (*forwarder.ForwarderForwardSucceeded, error) {
	ret := _m.Called(log)

	var r0 *forwarder.ForwarderForwardSucceeded
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*forwarder.ForwarderForwardSucceeded, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *forwarder.ForwarderForwardSucceeded); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*forwarder.ForwarderForwardSucceeded)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseLog provides a mock function with given fields: log
func (_m *ForwarderInterface) ParseLog(log types.Log) (generated.AbigenLog, error) {
	ret := _m.Called(log)

	var r0 generated.AbigenLog
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (generated.AbigenLog, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) generated.AbigenLog); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.AbigenLog)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseOwnershipTransferRequested provides a mock function with given fields: log
func (_m *ForwarderInterface) ParseOwnershipTransferRequested(log types.Log) (*forwarder.ForwarderOwnershipTransferRequested, error) {
	ret := _m.Called(log)

	var r0 *forwarder.ForwarderOwnershipTransferRequested
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*forwarder.ForwarderOwnershipTransferRequested, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *forwarder.ForwarderOwnershipTransferRequested); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*forwarder.ForwarderOwnershipTransferRequested)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseOwnershipTransferred provides a mock function with given fields: log
func (_m *ForwarderInterface) ParseOwnershipTransferred(log types.Log) (*forwarder.ForwarderOwnershipTransferred, error) {
	ret := _m.Called(log)

	var r0 *forwarder.ForwarderOwnershipTransferred
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*forwarder.ForwarderOwnershipTransferred, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *forwarder.ForwarderOwnershipTransferred); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*forwarder.ForwarderOwnershipTransferred)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseRequestTypeRegistered provides a mock function with given fields: log
func (_m *ForwarderInterface) ParseRequestTypeRegistered(log types.Log) (*forwarder.ForwarderRequestTypeRegistered, error) {
	ret := _m.Called(log)

	var r0 *forwarder.ForwarderRequestTypeRegistered
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*forwarder.ForwarderRequestTypeRegistered, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *forwarder.ForwarderRequestTypeRegistered); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*forwarder.ForwarderRequestTypeRegistered)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Receive provides a mock function with given fields: opts
func (_m *ForwarderInterface) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(opts)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) (*types.Transaction, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterDomainSeparator provides a mock function with given fields: opts, name, version
func (_m *ForwarderInterface) RegisterDomainSeparator(opts *bind.TransactOpts, name string, version string) (*types.Transaction, error) {
	ret := _m.Called(opts, name, version)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, string, string) (*types.Transaction, error)); ok {
		return rf(opts, name, version)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, string, string) *types.Transaction); ok {
		r0 = rf(opts, name, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, string, string) error); ok {
		r1 = rf(opts, name, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterRequestType provides a mock function with given fields: opts, typeName, typeSuffix
func (_m *ForwarderInterface) RegisterRequestType(opts *bind.TransactOpts, typeName string, typeSuffix string) (*types.Transaction, error) {
	ret := _m.Called(opts, typeName, typeSuffix)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, string, string) (*types.Transaction, error)); ok {
		return rf(opts, typeName, typeSuffix)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, string, string) *types.Transaction); ok {
		r0 = rf(opts, typeName, typeSuffix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, string, string) error); ok {
		r1 = rf(opts, typeName, typeSuffix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SDomains provides a mock function with given fields: opts, arg0
func (_m *ForwarderInterface) SDomains(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	ret := _m.Called(opts, arg0)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, [32]byte) (bool, error)); ok {
		return rf(opts, arg0)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, [32]byte) bool); ok {
		r0 = rf(opts, arg0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, [32]byte) error); ok {
		r1 = rf(opts, arg0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// STypeHashes provides a mock function with given fields: opts, arg0
func (_m *ForwarderInterface) STypeHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	ret := _m.Called(opts, arg0)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, [32]byte) (bool, error)); ok {
		return rf(opts, arg0)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, [32]byte) bool); ok {
		r0 = rf(opts, arg0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, [32]byte) error); ok {
		r1 = rf(opts, arg0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SupportsInterface provides a mock function with given fields: opts, interfaceId
func (_m *ForwarderInterface) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	ret := _m.Called(opts, interfaceId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, [4]byte) (bool, error)); ok {
		return rf(opts, interfaceId)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, [4]byte) bool); ok {
		r0 = rf(opts, interfaceId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, [4]byte) error); ok {
		r1 = rf(opts, interfaceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferOwnership provides a mock function with given fields: opts, to
func (_m *ForwarderInterface) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, to)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) (*types.Transaction, error)); ok {
		return rf(opts, to)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Verify provides a mock function with given fields: opts, req, domainSeparator, requestTypeHash, suffixData, sig
func (_m *ForwarderInterface) Verify(opts *bind.CallOpts, req forwarder.IForwarderForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) error {
	ret := _m.Called(opts, req, domainSeparator, requestTypeHash, suffixData, sig)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, forwarder.IForwarderForwardRequest, [32]byte, [32]byte, []byte, []byte) error); ok {
		r0 = rf(opts, req, domainSeparator, requestTypeHash, suffixData, sig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WatchDomainRegistered provides a mock function with given fields: opts, sink, domainSeparator
func (_m *ForwarderInterface) WatchDomainRegistered(opts *bind.WatchOpts, sink chan<- *forwarder.ForwarderDomainRegistered, domainSeparator [][32]byte) (event.Subscription, error) {
	ret := _m.Called(opts, sink, domainSeparator)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderDomainRegistered, [][32]byte) (event.Subscription, error)); ok {
		return rf(opts, sink, domainSeparator)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderDomainRegistered, [][32]byte) event.Subscription); ok {
		r0 = rf(opts, sink, domainSeparator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderDomainRegistered, [][32]byte) error); ok {
		r1 = rf(opts, sink, domainSeparator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchForwardSucceeded provides a mock function with given fields: opts, sink, from, target, domainSeparator
func (_m *ForwarderInterface) WatchForwardSucceeded(opts *bind.WatchOpts, sink chan<- *forwarder.ForwarderForwardSucceeded, from []common.Address, target []common.Address, domainSeparator [][32]byte) (event.Subscription, error) {
	ret := _m.Called(opts, sink, from, target, domainSeparator)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderForwardSucceeded, []common.Address, []common.Address, [][32]byte) (event.Subscription, error)); ok {
		return rf(opts, sink, from, target, domainSeparator)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderForwardSucceeded, []common.Address, []common.Address, [][32]byte) event.Subscription); ok {
		r0 = rf(opts, sink, from, target, domainSeparator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderForwardSucceeded, []common.Address, []common.Address, [][32]byte) error); ok {
		r1 = rf(opts, sink, from, target, domainSeparator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchOwnershipTransferRequested provides a mock function with given fields: opts, sink, from, to
func (_m *ForwarderInterface) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *forwarder.ForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, from, to)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderOwnershipTransferRequested, []common.Address, []common.Address) (event.Subscription, error)); ok {
		return rf(opts, sink, from, to)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderOwnershipTransferRequested, []common.Address, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderOwnershipTransferRequested, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, sink, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchOwnershipTransferred provides a mock function with given fields: opts, sink, from, to
func (_m *ForwarderInterface) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *forwarder.ForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, from, to)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderOwnershipTransferred, []common.Address, []common.Address) (event.Subscription, error)); ok {
		return rf(opts, sink, from, to)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderOwnershipTransferred, []common.Address, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderOwnershipTransferred, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, sink, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchRequestTypeRegistered provides a mock function with given fields: opts, sink, typeHash
func (_m *ForwarderInterface) WatchRequestTypeRegistered(opts *bind.WatchOpts, sink chan<- *forwarder.ForwarderRequestTypeRegistered, typeHash [][32]byte) (event.Subscription, error) {
	ret := _m.Called(opts, sink, typeHash)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderRequestTypeRegistered, [][32]byte) (event.Subscription, error)); ok {
		return rf(opts, sink, typeHash)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderRequestTypeRegistered, [][32]byte) event.Subscription); ok {
		r0 = rf(opts, sink, typeHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *forwarder.ForwarderRequestTypeRegistered, [][32]byte) error); ok {
		r1 = rf(opts, sink, typeHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewForwarderInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewForwarderInterface creates a new instance of ForwarderInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewForwarderInterface(t mockConstructorTestingTNewForwarderInterface) *ForwarderInterface {
	mock := &ForwarderInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
