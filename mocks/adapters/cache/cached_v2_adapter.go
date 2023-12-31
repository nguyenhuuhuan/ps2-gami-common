// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cache "gitlab.id.vin/gami/ps2-gami-common/adapters/cache"

	mock "github.com/stretchr/testify/mock"

	redis "github.com/go-redis/redis/v8"

	time "time"
)

// CachedV2Adapter is an autogenerated mock type for the CachedV2Adapter type
type CachedV2Adapter struct {
	mock.Mock
}

// AppendIntArray provides a mock function with given fields: ctx, key, value
func (_m *CachedV2Adapter) AppendIntArray(ctx context.Context, key string, value []int64) (bool, error) {
	ret := _m.Called(ctx, key, value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, []int64) bool); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []int64) error); ok {
		r1 = rf(ctx, key, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BitCount provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) BitCount(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BitCountAll provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) BitCountAll(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BitField provides a mock function with given fields: ctx, key, args
func (_m *CachedV2Adapter) BitField(ctx context.Context, key string, args []cache.BitFieldModel) ([]int64, error) {
	ret := _m.Called(ctx, key, args)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(context.Context, string, []cache.BitFieldModel) []int64); ok {
		r0 = rf(ctx, key, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []cache.BitFieldModel) error); ok {
		r1 = rf(ctx, key, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BitOpAnd provides a mock function with given fields: ctx, destKey, keys
func (_m *CachedV2Adapter) BitOpAnd(ctx context.Context, destKey string, keys ...string) (int64, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, destKey)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) int64); ok {
		r0 = rf(ctx, destKey, keys...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...string) error); ok {
		r1 = rf(ctx, destKey, keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BitOpNot provides a mock function with given fields: ctx, destKey, key
func (_m *CachedV2Adapter) BitOpNot(ctx context.Context, destKey string, key string) (int64, error) {
	ret := _m.Called(ctx, destKey, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(ctx, destKey, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, destKey, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BitOpOr provides a mock function with given fields: ctx, destKey, keys
func (_m *CachedV2Adapter) BitOpOr(ctx context.Context, destKey string, keys ...string) (int64, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, destKey)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) int64); ok {
		r0 = rf(ctx, destKey, keys...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...string) error); ok {
		r1 = rf(ctx, destKey, keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BitOpXor provides a mock function with given fields: ctx, destKey, keys
func (_m *CachedV2Adapter) BitOpXor(ctx context.Context, destKey string, keys ...string) (int64, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, destKey)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) int64); ok {
		r0 = rf(ctx, destKey, keys...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...string) error); ok {
		r1 = rf(ctx, destKey, keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BitPos provides a mock function with given fields: ctx, key, bit, pos
func (_m *CachedV2Adapter) BitPos(ctx context.Context, key string, bit int64, pos ...int64) (int64, error) {
	_va := make([]interface{}, len(pos))
	for _i := range pos {
		_va[_i] = pos[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key, bit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, ...int64) int64); ok {
		r0 = rf(ctx, key, bit, pos...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64, ...int64) error); ok {
		r1 = rf(ctx, key, bit, pos...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Del provides a mock function with given fields: ctx, keys
func (_m *CachedV2Adapter) Del(ctx context.Context, keys ...string) error {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...string) error); ok {
		r0 = rf(ctx, keys...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) Exists(ctx context.Context, key string) bool {
	ret := _m.Called(ctx, key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ExistsV2 provides a mock function with given fields: ctx, keys
func (_m *CachedV2Adapter) ExistsV2(ctx context.Context, keys ...string) (int64, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, ...string) int64); ok {
		r0 = rf(ctx, keys...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...string) error); ok {
		r1 = rf(ctx, keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Expire provides a mock function with given fields: ctx, key, expiration
func (_m *CachedV2Adapter) Expire(ctx context.Context, key string, expiration time.Duration) error {
	ret := _m.Called(ctx, key, expiration)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration) error); ok {
		r0 = rf(ctx, key, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Flush provides a mock function with given fields: ctx
func (_m *CachedV2Adapter) Flush(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, key, v
func (_m *CachedV2Adapter) Get(ctx context.Context, key string, v interface{}) error {
	ret := _m.Called(ctx, key, v)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, key, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBit provides a mock function with given fields: ctx, key, offset
func (_m *CachedV2Adapter) GetBit(ctx context.Context, key string, offset int64) (int64, error) {
	ret := _m.Called(ctx, key, offset)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) int64); ok {
		r0 = rf(ctx, key, offset)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, key, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBytes provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) GetBytes(ctx context.Context, key string) ([]byte, error) {
	ret := _m.Called(ctx, key)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInt64 provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) GetInt64(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetString provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) GetString(ctx context.Context, key string) (string, error) {
	ret := _m.Called(ctx, key)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HDel provides a mock function with given fields: ctx, key, member
func (_m *CachedV2Adapter) HDel(ctx context.Context, key string, member string) error {
	ret := _m.Called(ctx, key, member)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, key, member)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HExists provides a mock function with given fields: ctx, key, field
func (_m *CachedV2Adapter) HExists(ctx context.Context, key string, field string) bool {
	ret := _m.Called(ctx, key, field)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, key, field)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HExistsV2 provides a mock function with given fields: ctx, key, field
func (_m *CachedV2Adapter) HExistsV2(ctx context.Context, key string, field string) (bool, error) {
	ret := _m.Called(ctx, key, field)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, key, field)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, key, field)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HGet provides a mock function with given fields: ctx, key, member, v
func (_m *CachedV2Adapter) HGet(ctx context.Context, key string, member string, v interface{}) error {
	ret := _m.Called(ctx, key, member, v)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, interface{}) error); ok {
		r0 = rf(ctx, key, member, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HGetAll provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	ret := _m.Called(ctx, key)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string]string); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HGetInt64 provides a mock function with given fields: ctx, key, member
func (_m *CachedV2Adapter) HGetInt64(ctx context.Context, key string, member string) (int64, error) {
	ret := _m.Called(ctx, key, member)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(ctx, key, member)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, key, member)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HIncr provides a mock function with given fields: ctx, key, member, quantity
func (_m *CachedV2Adapter) HIncr(ctx context.Context, key string, member string, quantity int64) (int64, error) {
	ret := _m.Called(ctx, key, member, quantity)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) int64); ok {
		r0 = rf(ctx, key, member, quantity)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) error); ok {
		r1 = rf(ctx, key, member, quantity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HIncrBy provides a mock function with given fields: ctx, key, field, incr
func (_m *CachedV2Adapter) HIncrBy(ctx context.Context, key string, field string, incr int64) (int64, error) {
	ret := _m.Called(ctx, key, field, incr)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) int64); ok {
		r0 = rf(ctx, key, field, incr)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) error); ok {
		r1 = rf(ctx, key, field, incr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HLen provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) HLen(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HMGet provides a mock function with given fields: ctx, key, field
func (_m *CachedV2Adapter) HMGet(ctx context.Context, key string, field []string) ([]interface{}, error) {
	ret := _m.Called(ctx, key, field)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) []interface{}); ok {
		r0 = rf(ctx, key, field)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, key, field)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HMSet provides a mock function with given fields: ctx, key, fields
func (_m *CachedV2Adapter) HMSet(ctx context.Context, key string, fields interface{}) error {
	ret := _m.Called(ctx, key, fields)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, key, fields)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HSet provides a mock function with given fields: ctx, key, field, value
func (_m *CachedV2Adapter) HSet(ctx context.Context, key string, field string, value interface{}) error {
	ret := _m.Called(ctx, key, field, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, interface{}) error); ok {
		r0 = rf(ctx, key, field, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HSetInt64 provides a mock function with given fields: ctx, key, member, value
func (_m *CachedV2Adapter) HSetInt64(ctx context.Context, key string, member string, value int64) error {
	ret := _m.Called(ctx, key, member, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) error); ok {
		r0 = rf(ctx, key, member, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Incr provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) Incr(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrBy provides a mock function with given fields: ctx, key, quantity
func (_m *CachedV2Adapter) IncrBy(ctx context.Context, key string, quantity int64) (int64, error) {
	ret := _m.Called(ctx, key, quantity)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) int64); ok {
		r0 = rf(ctx, key, quantity)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, key, quantity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Keys provides a mock function with given fields: ctx, pattern
func (_m *CachedV2Adapter) Keys(ctx context.Context, pattern string) ([]string, error) {
	ret := _m.Called(ctx, pattern)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, pattern)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, pattern)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LLength provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) LLength(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LPush provides a mock function with given fields: ctx, key, val
func (_m *CachedV2Adapter) LPush(ctx context.Context, key string, val ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, val...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) error); ok {
		r0 = rf(ctx, key, val...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LRange provides a mock function with given fields: ctx, key, start, stop
func (_m *CachedV2Adapter) LRange(ctx context.Context, key string, start int64, stop int64) ([]string, error) {
	ret := _m.Called(ctx, key, start, stop)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64) []string); ok {
		r0 = rf(ctx, key, start, stop)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64, int64) error); ok {
		r1 = rf(ctx, key, start, stop)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LRem provides a mock function with given fields: ctx, key, val
func (_m *CachedV2Adapter) LRem(ctx context.Context, key string, val interface{}) (int64, error) {
	ret := _m.Called(ctx, key, val)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) int64); ok {
		r0 = rf(ctx, key, val)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, key, val)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Lock provides a mock function with given fields: ctx, key, expiration, acquire, interval
func (_m *CachedV2Adapter) Lock(ctx context.Context, key string, expiration time.Duration, acquire time.Duration, interval time.Duration) error {
	ret := _m.Called(ctx, key, expiration, acquire, interval)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration, time.Duration, time.Duration) error); ok {
		r0 = rf(ctx, key, expiration, acquire, interval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MGet provides a mock function with given fields: ctx, keys
func (_m *CachedV2Adapter) MGet(ctx context.Context, keys []string) ([]interface{}, error) {
	ret := _m.Called(ctx, keys)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []interface{}); ok {
		r0 = rf(ctx, keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MSet provides a mock function with given fields: ctx, m
func (_m *CachedV2Adapter) MSet(ctx context.Context, m map[string]interface{}) error {
	ret := _m.Called(ctx, m)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) error); ok {
		r0 = rf(ctx, m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MSetNX provides a mock function with given fields: ctx, m
func (_m *CachedV2Adapter) MSetNX(ctx context.Context, m map[string]interface{}) (bool, error) {
	ret := _m.Called(ctx, m)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) bool); ok {
		r0 = rf(ctx, m)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PFAdd provides a mock function with given fields: ctx, key, list
func (_m *CachedV2Adapter) PFAdd(ctx context.Context, key string, list []string) error {
	ret := _m.Called(ctx, key, list)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) error); ok {
		r0 = rf(ctx, key, list)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PFCount provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) PFCount(ctx context.Context, key string) int64 {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// RAppendIntArray provides a mock function with given fields: ctx, key, value
func (_m *CachedV2Adapter) RAppendIntArray(ctx context.Context, key string, value []int64) (bool, error) {
	ret := _m.Called(ctx, key, value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, []int64) bool); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []int64) error); ok {
		r1 = rf(ctx, key, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RBitCount provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) RBitCount(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RGetBit provides a mock function with given fields: ctx, key, offset
func (_m *CachedV2Adapter) RGetBit(ctx context.Context, key string, offset int64) (int64, error) {
	ret := _m.Called(ctx, key, offset)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) int64); ok {
		r0 = rf(ctx, key, offset)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, key, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RGetBits provides a mock function with given fields: ctx, keys, offset
func (_m *CachedV2Adapter) RGetBits(ctx context.Context, keys []string, offset int64) ([]int, error) {
	ret := _m.Called(ctx, keys, offset)

	var r0 []int
	if rf, ok := ret.Get(0).(func(context.Context, []string, int64) []int); ok {
		r0 = rf(ctx, keys, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, int64) error); ok {
		r1 = rf(ctx, keys, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RGetIntArray provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) RGetIntArray(ctx context.Context, key string) ([]int64, error) {
	ret := _m.Called(ctx, key)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(context.Context, string) []int64); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ROptimize provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) ROptimize(ctx context.Context, key string) (bool, error) {
	ret := _m.Called(ctx, key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RPop provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) RPop(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RPush provides a mock function with given fields: ctx, key, val
func (_m *CachedV2Adapter) RPush(ctx context.Context, key string, val ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, val...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) error); ok {
		r0 = rf(ctx, key, val...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RSetBit provides a mock function with given fields: ctx, key, offset, value
func (_m *CachedV2Adapter) RSetBit(ctx context.Context, key string, offset int64, value int) (int64, error) {
	ret := _m.Called(ctx, key, offset, value)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int) int64); ok {
		r0 = rf(ctx, key, offset, value)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64, int) error); ok {
		r1 = rf(ctx, key, offset, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RSetIntArray provides a mock function with given fields: ctx, key, value
func (_m *CachedV2Adapter) RSetIntArray(ctx context.Context, key string, value []int64) (bool, error) {
	ret := _m.Called(ctx, key, value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, []int64) bool); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []int64) error); ok {
		r1 = rf(ctx, key, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SAdd provides a mock function with given fields: ctx, key, list
func (_m *CachedV2Adapter) SAdd(ctx context.Context, key string, list []string) error {
	ret := _m.Called(ctx, key, list)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) error); ok {
		r0 = rf(ctx, key, list)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SCard provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) SCard(ctx context.Context, key string) int64 {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// SMMem provides a mock function with given fields: ctx, key, list
func (_m *CachedV2Adapter) SMMem(ctx context.Context, key string, list []string) []bool {
	ret := _m.Called(ctx, key, list)

	var r0 []bool
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) []bool); ok {
		r0 = rf(ctx, key, list)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bool)
		}
	}

	return r0
}

// SMem provides a mock function with given fields: ctx, key, value
func (_m *CachedV2Adapter) SMem(ctx context.Context, key string, value string) bool {
	ret := _m.Called(ctx, key, value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SRem provides a mock function with given fields: ctx, key, list
func (_m *CachedV2Adapter) SRem(ctx context.Context, key string, list []string) error {
	ret := _m.Called(ctx, key, list)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) error); ok {
		r0 = rf(ctx, key, list)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Set provides a mock function with given fields: ctx, key, v, expiration
func (_m *CachedV2Adapter) Set(ctx context.Context, key string, v interface{}, expiration time.Duration) error {
	ret := _m.Called(ctx, key, v, expiration)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, v, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetBit provides a mock function with given fields: ctx, key, offset, value
func (_m *CachedV2Adapter) SetBit(ctx context.Context, key string, offset int64, value int) error {
	ret := _m.Called(ctx, key, offset, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int) error); ok {
		r0 = rf(ctx, key, offset, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetNX provides a mock function with given fields: ctx, key, value, expiration
func (_m *CachedV2Adapter) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	ret := _m.Called(ctx, key, value, expiration)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) bool); ok {
		r0 = rf(ctx, key, value, expiration)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r1 = rf(ctx, key, value, expiration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetString provides a mock function with given fields: ctx, key, v, expiration
func (_m *CachedV2Adapter) SetString(ctx context.Context, key string, v interface{}, expiration time.Duration) error {
	ret := _m.Called(ctx, key, v, expiration)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, v, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StrLen provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) StrLen(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ZAdd provides a mock function with given fields: ctx, key, v
func (_m *CachedV2Adapter) ZAdd(ctx context.Context, key string, v *redis.Z) error {
	ret := _m.Called(ctx, key, v)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *redis.Z) error); ok {
		r0 = rf(ctx, key, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ZAdds provides a mock function with given fields: ctx, key, v
func (_m *CachedV2Adapter) ZAdds(ctx context.Context, key string, v ...*redis.Z) error {
	_va := make([]interface{}, len(v))
	for _i := range v {
		_va[_i] = v[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...*redis.Z) error); ok {
		r0 = rf(ctx, key, v...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ZCard provides a mock function with given fields: ctx, key
func (_m *CachedV2Adapter) ZCard(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ZCount provides a mock function with given fields: ctx, key, min, max
func (_m *CachedV2Adapter) ZCount(ctx context.Context, key string, min int64, max int64) (int64, error) {
	ret := _m.Called(ctx, key, min, max)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64) int64); ok {
		r0 = rf(ctx, key, min, max)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64, int64) error); ok {
		r1 = rf(ctx, key, min, max)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ZIncr provides a mock function with given fields: ctx, key, member
func (_m *CachedV2Adapter) ZIncr(ctx context.Context, key string, member *redis.Z) error {
	ret := _m.Called(ctx, key, member)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *redis.Z) error); ok {
		r0 = rf(ctx, key, member)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ZMScore provides a mock function with given fields: ctx, key, members
func (_m *CachedV2Adapter) ZMScore(ctx context.Context, key string, members ...string) ([]float64, error) {
	_va := make([]interface{}, len(members))
	for _i := range members {
		_va[_i] = members[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []float64
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) []float64); ok {
		r0 = rf(ctx, key, members...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...string) error); ok {
		r1 = rf(ctx, key, members...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ZRank provides a mock function with given fields: ctx, key, member
func (_m *CachedV2Adapter) ZRank(ctx context.Context, key string, member string) (int64, error) {
	ret := _m.Called(ctx, key, member)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(ctx, key, member)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, key, member)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ZRem provides a mock function with given fields: ctx, key, member
func (_m *CachedV2Adapter) ZRem(ctx context.Context, key string, member interface{}) error {
	ret := _m.Called(ctx, key, member)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, key, member)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ZRevRange provides a mock function with given fields: ctx, key, min, max
func (_m *CachedV2Adapter) ZRevRange(ctx context.Context, key string, min int64, max int64) ([]redis.Z, error) {
	ret := _m.Called(ctx, key, min, max)

	var r0 []redis.Z
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64) []redis.Z); ok {
		r0 = rf(ctx, key, min, max)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]redis.Z)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64, int64) error); ok {
		r1 = rf(ctx, key, min, max)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ZRevRangeByScore provides a mock function with given fields: ctx, key, min, max, offset, count
func (_m *CachedV2Adapter) ZRevRangeByScore(ctx context.Context, key string, min string, max string, offset int64, count int64) ([]redis.Z, error) {
	ret := _m.Called(ctx, key, min, max, offset, count)

	var r0 []redis.Z
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, int64, int64) []redis.Z); ok {
		r0 = rf(ctx, key, min, max, offset, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]redis.Z)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, int64, int64) error); ok {
		r1 = rf(ctx, key, min, max, offset, count)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ZRevRank provides a mock function with given fields: ctx, key, member
func (_m *CachedV2Adapter) ZRevRank(ctx context.Context, key string, member string) (int64, error) {
	ret := _m.Called(ctx, key, member)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(ctx, key, member)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, key, member)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ZScan provides a mock function with given fields: ctx, key, match
func (_m *CachedV2Adapter) ZScan(ctx context.Context, key string, match string) ([]string, error) {
	ret := _m.Called(ctx, key, match)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []string); ok {
		r0 = rf(ctx, key, match)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, key, match)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ZScore provides a mock function with given fields: ctx, key, member
func (_m *CachedV2Adapter) ZScore(ctx context.Context, key string, member string) (float64, error) {
	ret := _m.Called(ctx, key, member)

	var r0 float64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) float64); ok {
		r0 = rf(ctx, key, member)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, key, member)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCachedV2Adapter interface {
	mock.TestingT
	Cleanup(func())
}

// NewCachedV2Adapter creates a new instance of CachedV2Adapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCachedV2Adapter(t mockConstructorTestingTNewCachedV2Adapter) *CachedV2Adapter {
	mock := &CachedV2Adapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
