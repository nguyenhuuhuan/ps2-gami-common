// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	user "gitlab.id.vin/gami/ps2-gami-common/adapters/user"
)

// AuthAdapter is an autogenerated mock type for the AuthAdapter type
type AuthAdapter struct {
	mock.Mock
}

// GetCacheToken provides a mock function with given fields: key
func (_m *AuthAdapter) GetCacheToken(key string) (string, bool) {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetListUserProfile provides a mock function with given fields: request
func (_m *AuthAdapter) GetListUserProfile(request *user.ListUserRequest) (*user.ListUserResponse, error) {
	ret := _m.Called(request)

	var r0 *user.ListUserResponse
	if rf, ok := ret.Get(0).(func(*user.ListUserRequest) *user.ListUserResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.ListUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.ListUserRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetToken provides a mock function with given fields: request
func (_m *AuthAdapter) GetToken(request *user.TokenRequest) (*user.GetTokenResponse, error) {
	ret := _m.Called(request)

	var r0 *user.GetTokenResponse
	if rf, ok := ret.Get(0).(func(*user.TokenRequest) *user.GetTokenResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.GetTokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.TokenRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserProfile provides a mock function with given fields: request
func (_m *AuthAdapter) GetUserProfile(request *user.ProfileRequest) (*user.ProfileResponse, error) {
	ret := _m.Called(request)

	var r0 *user.ProfileResponse
	if rf, ok := ret.Get(0).(func(*user.ProfileRequest) *user.ProfileResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.ProfileResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.ProfileRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetCacheToken provides a mock function with given fields: key, token
func (_m *AuthAdapter) SetCacheToken(key string, token string) bool {
	ret := _m.Called(key, token)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(key, token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewAuthAdapter interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthAdapter creates a new instance of AuthAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthAdapter(t mockConstructorTestingTNewAuthAdapter) *AuthAdapter {
	mock := &AuthAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
