// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/cloudfoundry/bosh-cli/v6/test_support (interfaces: Spy)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Spy interface
type MockSpy struct {
	ctrl     *gomock.Controller
	recorder *_MockSpyRecorder
}

// Recorder for MockSpy (not exported)
type _MockSpyRecorder struct {
	mock *MockSpy
}

func NewMockSpy(ctrl *gomock.Controller) *MockSpy {
	mock := &MockSpy{ctrl: ctrl}
	mock.recorder = &_MockSpyRecorder{mock}
	return mock
}

func (_m *MockSpy) EXPECT() *_MockSpyRecorder {
	return _m.recorder
}

func (_m *MockSpy) Record() {
	_m.ctrl.Call(_m, "Record")
}

func (_mr *_MockSpyRecorder) Record() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Record")
}
