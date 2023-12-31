// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	transaction_log "gitlab.id.vin/gami/ps2-gami-common/adapters/transaction_log"
)

// Adapter is an autogenerated mock type for the Adapter type
type Adapter struct {
	mock.Mock
}

// GetTransactionLog provides a mock function with given fields: ctx, request
func (_m *Adapter) GetTransactionLog(ctx context.Context, request *transaction_log.GetTransactionRequest) (*transaction_log.GetTransactionResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *transaction_log.GetTransactionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *transaction_log.GetTransactionRequest) *transaction_log.GetTransactionResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transaction_log.GetTransactionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *transaction_log.GetTransactionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAdapter interface {
	mock.TestingT
	Cleanup(func())
}

// NewAdapter creates a new instance of Adapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAdapter(t mockConstructorTestingTNewAdapter) *Adapter {
	mock := &Adapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
