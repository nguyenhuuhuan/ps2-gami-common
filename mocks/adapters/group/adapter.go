// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	group "gitlab.id.vin/gami/ps2-gami-common/adapters/group"
)

// Adapter is an autogenerated mock type for the Adapter type
type Adapter struct {
	mock.Mock
}

// IsInGroups provides a mock function with given fields: ctx, req
func (_m *Adapter) IsInGroups(ctx context.Context, req *group.IsInGroupsRequest) (*group.IsInGroupsResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *group.IsInGroupsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *group.IsInGroupsRequest) *group.IsInGroupsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*group.IsInGroupsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *group.IsInGroupsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsUserInGroups provides a mock function with given fields: ctx, req
func (_m *Adapter) IsUserInGroups(ctx context.Context, req *group.IsUserInGroupsRequest) (*group.IsUserInGroupsResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *group.IsUserInGroupsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *group.IsUserInGroupsRequest) *group.IsUserInGroupsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*group.IsUserInGroupsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *group.IsUserInGroupsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValuesInGroups provides a mock function with given fields: ctx, req
func (_m *Adapter) ValuesInGroups(ctx context.Context, req *group.ValuesInGroupsRequest) (*group.ValuesInGroupsResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *group.ValuesInGroupsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *group.ValuesInGroupsRequest) *group.ValuesInGroupsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*group.ValuesInGroupsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *group.ValuesInGroupsRequest) error); ok {
		r1 = rf(ctx, req)
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
