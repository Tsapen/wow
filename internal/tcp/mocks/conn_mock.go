// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	net "net"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// ConnMock is an autogenerated mock type for the ConnMock type
type ConnMock struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ConnMock) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LocalAddr provides a mock function with given fields:
func (_m *ConnMock) LocalAddr() net.Addr {
	ret := _m.Called()

	var r0 net.Addr
	if rf, ok := ret.Get(0).(func() net.Addr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Addr)
		}
	}

	return r0
}

// Read provides a mock function with given fields: b
func (_m *ConnMock) Read(b []byte) (int, error) {
	ret := _m.Called(b)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(b)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoteAddr provides a mock function with given fields:
func (_m *ConnMock) RemoteAddr() net.Addr {
	ret := _m.Called()

	var r0 net.Addr
	if rf, ok := ret.Get(0).(func() net.Addr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Addr)
		}
	}

	return r0
}

// SetDeadline provides a mock function with given fields: t
func (_m *ConnMock) SetDeadline(t time.Time) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetReadDeadline provides a mock function with given fields: t
func (_m *ConnMock) SetReadDeadline(t time.Time) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWriteDeadline provides a mock function with given fields: t
func (_m *ConnMock) SetWriteDeadline(t time.Time) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: b
func (_m *ConnMock) Write(b []byte) (int, error) {
	ret := _m.Called(b)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(b)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewConnMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewConnMock creates a new instance of ConnMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConnMock(t mockConstructorTestingTNewConnMock) *ConnMock {
	mock := &ConnMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
