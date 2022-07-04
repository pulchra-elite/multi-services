// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	testing "testing"

	v2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
)

// AutoscalingV2beta2Interface is an autogenerated mock type for the AutoscalingV2beta2Interface type
type AutoscalingV2beta2Interface struct {
	mock.Mock
}

// HorizontalPodAutoscalers provides a mock function with given fields: namespace
func (_m *AutoscalingV2beta2Interface) HorizontalPodAutoscalers(namespace string) v2beta2.HorizontalPodAutoscalerInterface {
	ret := _m.Called(namespace)

	var r0 v2beta2.HorizontalPodAutoscalerInterface
	if rf, ok := ret.Get(0).(func(string) v2beta2.HorizontalPodAutoscalerInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v2beta2.HorizontalPodAutoscalerInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *AutoscalingV2beta2Interface) RESTClient() rest.Interface {
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

// NewAutoscalingV2beta2Interface creates a new instance of AutoscalingV2beta2Interface. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewAutoscalingV2beta2Interface(t testing.TB) *AutoscalingV2beta2Interface {
	mock := &AutoscalingV2beta2Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
