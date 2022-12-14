// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// ICache is an autogenerated mock type for the ICache type
type ICache struct {
	mock.Mock
}

// Del provides a mock function with given fields: ctx, key
func (_m *ICache) Del(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, key
func (_m *ICache) Get(ctx context.Context, key string) (string, error) {
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

// Set provides a mock function with given fields: ctx, key, value, exp
func (_m *ICache) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	ret := _m.Called(ctx, key, value, exp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, value, exp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
