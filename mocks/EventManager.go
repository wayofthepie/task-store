package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/execd/task-store/pkg/model"

// EventManager is an autogenerated mock type for the EventManager type
type EventManager struct {
	mock.Mock
}

// ListenForProgress provides a mock function with given fields: quit
func (_m *EventManager) ListenForProgress(quit <-chan int) (<-chan model.Info, <-chan error) {
	ret := _m.Called(quit)

	var r0 <-chan model.Info
	if rf, ok := ret.Get(0).(func(<-chan int) <-chan model.Info); ok {
		r0 = rf(quit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan model.Info)
		}
	}

	var r1 <-chan error
	if rf, ok := ret.Get(1).(func(<-chan int) <-chan error); ok {
		r1 = rf(quit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan error)
		}
	}

	return r0, r1
}

// PublishWork provides a mock function with given fields: _a0
func (_m *EventManager) PublishWork(_a0 *model.Spec) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Spec) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}