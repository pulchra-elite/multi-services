// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	v1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
)

// StorageClassesGetter is an autogenerated mock type for the StorageClassesGetter type
type StorageClassesGetter struct {
	mock.Mock
}

// StorageClasses provides a mock function with given fields:
func (_m *StorageClassesGetter) StorageClasses() v1beta1.StorageClassInterface {
	ret := _m.Called()

	var r0 v1beta1.StorageClassInterface
	if rf, ok := ret.Get(0).(func() v1beta1.StorageClassInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.StorageClassInterface)
		}
	}

	return r0
}

// NewStorageClassesGetter creates a new instance of StorageClassesGetter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorageClassesGetter(t testing.TB) *StorageClassesGetter {
	mock := &StorageClassesGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
