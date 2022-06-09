// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudfoundry/bosh-cli/v6/deployment/disk (interfaces: Disk,Manager)

// Package mocks is a generated GoMock package.
package mocks

import (
	disk "github.com/cloudfoundry/bosh-cli/v6/deployment/disk"
	manifest "github.com/cloudfoundry/bosh-cli/v6/deployment/manifest"
	ui "github.com/cloudfoundry/bosh-cli/v6/ui"
	property "github.com/cloudfoundry/bosh-utils/property"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDisk is a mock of Disk interface
type MockDisk struct {
	ctrl     *gomock.Controller
	recorder *MockDiskMockRecorder
}

// MockDiskMockRecorder is the mock recorder for MockDisk
type MockDiskMockRecorder struct {
	mock *MockDisk
}

// NewMockDisk creates a new mock instance
func NewMockDisk(ctrl *gomock.Controller) *MockDisk {
	mock := &MockDisk{ctrl: ctrl}
	mock.recorder = &MockDiskMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDisk) EXPECT() *MockDiskMockRecorder {
	return m.recorder
}

// CID mocks base method
func (m *MockDisk) CID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CID")
	ret0, _ := ret[0].(string)
	return ret0
}

// CID indicates an expected call of CID
func (mr *MockDiskMockRecorder) CID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CID", reflect.TypeOf((*MockDisk)(nil).CID))
}

// Delete mocks base method
func (m *MockDisk) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockDiskMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDisk)(nil).Delete))
}

// NeedsMigration mocks base method
func (m *MockDisk) NeedsMigration(arg0 int, arg1 property.Map) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeedsMigration", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0, nil
}

// NeedsMigration indicates an expected call of NeedsMigration
func (mr *MockDiskMockRecorder) NeedsMigration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeedsMigration", reflect.TypeOf((*MockDisk)(nil).NeedsMigration), arg0, arg1)
}

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockManager) Create(arg0 manifest.DiskPool, arg1 string) (disk.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(disk.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockManagerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockManager)(nil).Create), arg0, arg1)
}

// DeleteUnused mocks base method
func (m *MockManager) DeleteUnused(arg0 ui.Stage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUnused", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUnused indicates an expected call of DeleteUnused
func (mr *MockManagerMockRecorder) DeleteUnused(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUnused", reflect.TypeOf((*MockManager)(nil).DeleteUnused), arg0)
}

// FindCurrent mocks base method
func (m *MockManager) FindCurrent() ([]disk.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCurrent")
	ret0, _ := ret[0].([]disk.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCurrent indicates an expected call of FindCurrent
func (mr *MockManagerMockRecorder) FindCurrent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCurrent", reflect.TypeOf((*MockManager)(nil).FindCurrent))
}

// FindUnused mocks base method
func (m *MockManager) FindUnused() ([]disk.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUnused")
	ret0, _ := ret[0].([]disk.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUnused indicates an expected call of FindUnused
func (mr *MockManagerMockRecorder) FindUnused() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUnused", reflect.TypeOf((*MockManager)(nil).FindUnused))
}
