// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// VolumeAttachmentExpansion is an autogenerated mock type for the VolumeAttachmentExpansion type
type VolumeAttachmentExpansion struct {
	mock.Mock
}

// NewVolumeAttachmentExpansion creates a new instance of VolumeAttachmentExpansion. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewVolumeAttachmentExpansion(t testing.TB) *VolumeAttachmentExpansion {
	mock := &VolumeAttachmentExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
