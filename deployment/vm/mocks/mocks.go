// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudfoundry/bosh-cli/v6/deployment/vm (interfaces: ManagerFactory)

// Package mocks is a generated GoMock package.
package mocks

import (
	agentclient "github.com/cloudfoundry/bosh-agent/agentclient"
	cloud "github.com/cloudfoundry/bosh-cli/v6/cloud"
	vm "github.com/cloudfoundry/bosh-cli/v6/deployment/vm"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockManagerFactory is a mock of ManagerFactory interface
type MockManagerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockManagerFactoryMockRecorder
}

// MockManagerFactoryMockRecorder is the mock recorder for MockManagerFactory
type MockManagerFactoryMockRecorder struct {
	mock *MockManagerFactory
}

// NewMockManagerFactory creates a new mock instance
func NewMockManagerFactory(ctrl *gomock.Controller) *MockManagerFactory {
	mock := &MockManagerFactory{ctrl: ctrl}
	mock.recorder = &MockManagerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManagerFactory) EXPECT() *MockManagerFactoryMockRecorder {
	return m.recorder
}

// NewManager mocks base method
func (m *MockManagerFactory) NewManager(arg0 cloud.Cloud, arg1 agentclient.AgentClient) vm.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewManager", arg0, arg1)
	ret0, _ := ret[0].(vm.Manager)
	return ret0
}

// NewManager indicates an expected call of NewManager
func (mr *MockManagerFactoryMockRecorder) NewManager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewManager", reflect.TypeOf((*MockManagerFactory)(nil).NewManager), arg0, arg1)
}
