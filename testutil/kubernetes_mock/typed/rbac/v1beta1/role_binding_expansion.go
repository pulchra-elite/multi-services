// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// RoleBindingExpansion is an autogenerated mock type for the RoleBindingExpansion type
type RoleBindingExpansion struct {
	mock.Mock
}

// NewRoleBindingExpansion creates a new instance of RoleBindingExpansion. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoleBindingExpansion(t testing.TB) *RoleBindingExpansion {
	mock := &RoleBindingExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
