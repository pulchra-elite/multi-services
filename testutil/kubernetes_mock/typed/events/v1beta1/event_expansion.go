// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	v1beta1 "k8s.io/api/events/v1beta1"
)

// EventExpansion is an autogenerated mock type for the EventExpansion type
type EventExpansion struct {
	mock.Mock
}

// CreateWithEventNamespace provides a mock function with given fields: event
func (_m *EventExpansion) CreateWithEventNamespace(event *v1beta1.Event) (*v1beta1.Event, error) {
	ret := _m.Called(event)

	var r0 *v1beta1.Event
	if rf, ok := ret.Get(0).(func(*v1beta1.Event) *v1beta1.Event); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.Event) error); ok {
		r1 = rf(event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PatchWithEventNamespace provides a mock function with given fields: event, data
func (_m *EventExpansion) PatchWithEventNamespace(event *v1beta1.Event, data []byte) (*v1beta1.Event, error) {
	ret := _m.Called(event, data)

	var r0 *v1beta1.Event
	if rf, ok := ret.Get(0).(func(*v1beta1.Event, []byte) *v1beta1.Event); ok {
		r0 = rf(event, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.Event, []byte) error); ok {
		r1 = rf(event, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWithEventNamespace provides a mock function with given fields: event
func (_m *EventExpansion) UpdateWithEventNamespace(event *v1beta1.Event) (*v1beta1.Event, error) {
	ret := _m.Called(event)

	var r0 *v1beta1.Event
	if rf, ok := ret.Get(0).(func(*v1beta1.Event) *v1beta1.Event); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.Event) error); ok {
		r1 = rf(event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEventExpansion creates a new instance of EventExpansion. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventExpansion(t testing.TB) *EventExpansion {
	mock := &EventExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
