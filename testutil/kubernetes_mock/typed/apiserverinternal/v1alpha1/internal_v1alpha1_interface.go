// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	testing "testing"

	v1alpha1 "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1"
)

// InternalV1alpha1Interface is an autogenerated mock type for the InternalV1alpha1Interface type
type InternalV1alpha1Interface struct {
	mock.Mock
}

// RESTClient provides a mock function with given fields:
func (_m *InternalV1alpha1Interface) RESTClient() rest.Interface {
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

// StorageVersions provides a mock function with given fields:
func (_m *InternalV1alpha1Interface) StorageVersions() v1alpha1.StorageVersionInterface {
	ret := _m.Called()

	var r0 v1alpha1.StorageVersionInterface
	if rf, ok := ret.Get(0).(func() v1alpha1.StorageVersionInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.StorageVersionInterface)
		}
	}

	return r0
}

// NewInternalV1alpha1Interface creates a new instance of InternalV1alpha1Interface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewInternalV1alpha1Interface(t testing.TB) *InternalV1alpha1Interface {
	mock := &InternalV1alpha1Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
