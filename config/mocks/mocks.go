// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudfoundry/bosh-cli/v6/config (interfaces: LegacyDeploymentStateMigrator)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLegacyDeploymentStateMigrator is a mock of LegacyDeploymentStateMigrator interface
type MockLegacyDeploymentStateMigrator struct {
	ctrl     *gomock.Controller
	recorder *MockLegacyDeploymentStateMigratorMockRecorder
}

// MockLegacyDeploymentStateMigratorMockRecorder is the mock recorder for MockLegacyDeploymentStateMigrator
type MockLegacyDeploymentStateMigratorMockRecorder struct {
	mock *MockLegacyDeploymentStateMigrator
}

// NewMockLegacyDeploymentStateMigrator creates a new mock instance
func NewMockLegacyDeploymentStateMigrator(ctrl *gomock.Controller) *MockLegacyDeploymentStateMigrator {
	mock := &MockLegacyDeploymentStateMigrator{ctrl: ctrl}
	mock.recorder = &MockLegacyDeploymentStateMigratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLegacyDeploymentStateMigrator) EXPECT() *MockLegacyDeploymentStateMigratorMockRecorder {
	return m.recorder
}

// MigrateIfExists mocks base method
func (m *MockLegacyDeploymentStateMigrator) MigrateIfExists(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrateIfExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MigrateIfExists indicates an expected call of MigrateIfExists
func (mr *MockLegacyDeploymentStateMigratorMockRecorder) MigrateIfExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateIfExists", reflect.TypeOf((*MockLegacyDeploymentStateMigrator)(nil).MigrateIfExists), arg0)
}
