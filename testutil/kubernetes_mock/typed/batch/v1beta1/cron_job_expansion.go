// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// CronJobExpansion is an autogenerated mock type for the CronJobExpansion type
type CronJobExpansion struct {
	mock.Mock
}

// NewCronJobExpansion creates a new instance of CronJobExpansion. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCronJobExpansion(t testing.TB) *CronJobExpansion {
	mock := &CronJobExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
