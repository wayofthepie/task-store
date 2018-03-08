// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// Listener is an autogenerated mock type for the Listener type
type Listener struct {
	mock.Mock
}

// ListenForTaskStatus provides a mock function with given fields:
func (_m *Listener) ListenForTaskStatus() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
