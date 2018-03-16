package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/execd/task-store/pkg/model"

import uuid "github.com/satori/go.uuid"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// AddTaskToExecutingSet provides a mock function with given fields: id
func (_m *Store) AddTaskToExecutingSet(id *uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExecutingSetSize provides a mock function with given fields:
func (_m *Store) ExecutingSetSize() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTask provides a mock function with given fields: id
func (_m *Store) GetTask(id *uuid.UUID) (*model.Spec, error) {
	ret := _m.Called(id)

	var r0 *model.Spec
	if rf, ok := ret.Get(0).(func(*uuid.UUID) *model.Spec); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Spec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsTaskExecuting provides a mock function with given fields: id
func (_m *Store) IsTaskExecuting(id *uuid.UUID) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*uuid.UUID) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListenForTaskCreatedEvents provides a mock function with given fields:
func (_m *Store) ListenForTaskCreatedEvents() <-chan uuid.UUID {
	ret := _m.Called()

	var r0 <-chan uuid.UUID
	if rf, ok := ret.Get(0).(func() <-chan uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan uuid.UUID)
		}
	}

	return r0
}

// PopTask provides a mock function with given fields:
func (_m *Store) PopTask() (*uuid.UUID, error) {
	ret := _m.Called()

	var r0 *uuid.UUID
	if rf, ok := ret.Get(0).(func() *uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*uuid.UUID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublishTaskCreatedEvent provides a mock function with given fields: id
func (_m *Store) PublishTaskCreatedEvent(id *uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PushTask provides a mock function with given fields: id
func (_m *Store) PushTask(id *uuid.UUID) (int64, error) {
	ret := _m.Called(id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*uuid.UUID) int64); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveTaskFromExecutingSet provides a mock function with given fields: id
func (_m *Store) RemoveTaskFromExecutingSet(id *uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreTask provides a mock function with given fields: _a0
func (_m *Store) StoreTask(_a0 model.Spec) (*uuid.UUID, error) {
	ret := _m.Called(_a0)

	var r0 *uuid.UUID
	if rf, ok := ret.Get(0).(func(model.Spec) *uuid.UUID); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*uuid.UUID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Spec) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TaskQueueSize provides a mock function with given fields:
func (_m *Store) TaskQueueSize() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTaskInfo provides a mock function with given fields: info
func (_m *Store) UpdateTaskInfo(info *model.Info) error {
	ret := _m.Called(info)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Info) error); ok {
		r0 = rf(info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
