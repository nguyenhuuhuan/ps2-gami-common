// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gorm "github.com/jinzhu/gorm"
	database "gitlab.id.vin/gami/gami-common/adapters/database"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// DBAdapter is an autogenerated mock type for the DBAdapter type
type DBAdapter struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *DBAdapter) Begin() database.DBAdapter {
	ret := _m.Called()

	var r0 database.DBAdapter
	if rf, ok := ret.Get(0).(func() database.DBAdapter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.DBAdapter)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *DBAdapter) Close() {
	_m.Called()
}

// Commit provides a mock function with given fields:
func (_m *DBAdapter) Commit() {
	_m.Called()
}

// DB provides a mock function with given fields:
func (_m *DBAdapter) DB() *sql.DB {
	ret := _m.Called()

	var r0 *sql.DB
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	return r0
}

// Gormer provides a mock function with given fields:
func (_m *DBAdapter) Gormer() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Open provides a mock function with given fields: _a0
func (_m *DBAdapter) Open(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RollbackUselessCommitted provides a mock function with given fields:
func (_m *DBAdapter) RollbackUselessCommitted() {
	_m.Called()
}

type mockConstructorTestingTNewDBAdapter interface {
	mock.TestingT
	Cleanup(func())
}

// NewDBAdapter creates a new instance of DBAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDBAdapter(t mockConstructorTestingTNewDBAdapter) *DBAdapter {
	mock := &DBAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
