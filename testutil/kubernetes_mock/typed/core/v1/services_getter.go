// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// ServicesGetter is an autogenerated mock type for the ServicesGetter type
type ServicesGetter struct {
	mock.Mock
}

// Services provides a mock function with given fields: namespace
func (_m *ServicesGetter) Services(namespace string) v1.ServiceInterface {
	ret := _m.Called(namespace)

	var r0 v1.ServiceInterface
	if rf, ok := ret.Get(0).(func(string) v1.ServiceInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.ServiceInterface)
		}
	}

	return r0
}

// NewServicesGetter creates a new instance of ServicesGetter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewServicesGetter(t testing.TB) *ServicesGetter {
	mock := &ServicesGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
