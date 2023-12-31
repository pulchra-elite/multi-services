// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	v1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
)

// PodDisruptionBudgetsGetter is an autogenerated mock type for the PodDisruptionBudgetsGetter type
type PodDisruptionBudgetsGetter struct {
	mock.Mock
}

// PodDisruptionBudgets provides a mock function with given fields: namespace
func (_m *PodDisruptionBudgetsGetter) PodDisruptionBudgets(namespace string) v1beta1.PodDisruptionBudgetInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.PodDisruptionBudgetInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.PodDisruptionBudgetInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.PodDisruptionBudgetInterface)
		}
	}

	return r0
}

// NewPodDisruptionBudgetsGetter creates a new instance of PodDisruptionBudgetsGetter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPodDisruptionBudgetsGetter(t testing.TB) *PodDisruptionBudgetsGetter {
	mock := &PodDisruptionBudgetsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
