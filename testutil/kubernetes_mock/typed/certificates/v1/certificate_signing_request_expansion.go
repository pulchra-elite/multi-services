// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// CertificateSigningRequestExpansion is an autogenerated mock type for the CertificateSigningRequestExpansion type
type CertificateSigningRequestExpansion struct {
	mock.Mock
}

// NewCertificateSigningRequestExpansion creates a new instance of CertificateSigningRequestExpansion. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCertificateSigningRequestExpansion(t testing.TB) *CertificateSigningRequestExpansion {
	mock := &CertificateSigningRequestExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
