// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// JobExpansion is an autogenerated mock type for the JobExpansion type
type JobExpansion struct {
	mock.Mock
}

// NewJobExpansion creates a new instance of JobExpansion. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewJobExpansion(t testing.TB) *JobExpansion {
	mock := &JobExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
