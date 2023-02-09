// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	notification "gitlab.id.vin/gami/gami-common/adapters/notification"
)

// Adapter is an autogenerated mock type for the Adapter type
type Adapter struct {
	mock.Mock
}

// SendNotification provides a mock function with given fields: ctx, request
func (_m *Adapter) SendNotification(ctx context.Context, request *notification.Request) (*notification.SendNotificationResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *notification.SendNotificationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *notification.Request) *notification.SendNotificationResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notification.SendNotificationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *notification.Request) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendNotificationV2 provides a mock function with given fields: ctx, request
func (_m *Adapter) SendNotificationV2(ctx context.Context, request *notification.Request) (*notification.SendNotificationV2Response, error) {
	ret := _m.Called(ctx, request)

	var r0 *notification.SendNotificationV2Response
	if rf, ok := ret.Get(0).(func(context.Context, *notification.Request) *notification.SendNotificationV2Response); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notification.SendNotificationV2Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *notification.Request) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendNtfVinShop provides a mock function with given fields: ctx, request
func (_m *Adapter) SendNtfVinShop(ctx context.Context, request *notification.NtfVinShopRequest) (*notification.SendNtfVinShopResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *notification.SendNtfVinShopResponse
	if rf, ok := ret.Get(0).(func(context.Context, *notification.NtfVinShopRequest) *notification.SendNtfVinShopResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notification.SendNtfVinShopResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *notification.NtfVinShopRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendNtfVinShopV2 provides a mock function with given fields: ctx, request
func (_m *Adapter) SendNtfVinShopV2(ctx context.Context, request *notification.NtfVinShopRequest) (*notification.NtfVinShopV2Response, error) {
	ret := _m.Called(ctx, request)

	var r0 *notification.NtfVinShopV2Response
	if rf, ok := ret.Get(0).(func(context.Context, *notification.NtfVinShopRequest) *notification.NtfVinShopV2Response); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*notification.NtfVinShopV2Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *notification.NtfVinShopRequest) error); ok {
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
