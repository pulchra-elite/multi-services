// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// RoleExpansion is an autogenerated mock type for the RoleExpansion type
type RoleExpansion struct {
	mock.Mock
}

// NewRoleExpansion creates a new instance of RoleExpansion. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoleExpansion(t testing.TB) *RoleExpansion {
	mock := &RoleExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
