// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	testing "testing"

	v1beta1 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
)

// EventsV1beta1Interface is an autogenerated mock type for the EventsV1beta1Interface type
type EventsV1beta1Interface struct {
	mock.Mock
}

// Events provides a mock function with given fields: namespace
func (_m *EventsV1beta1Interface) Events(namespace string) v1beta1.EventInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.EventInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.EventInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.EventInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *EventsV1beta1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}

// NewEventsV1beta1Interface creates a new instance of EventsV1beta1Interface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventsV1beta1Interface(t testing.TB) *EventsV1beta1Interface {
	mock := &EventsV1beta1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
