// Code generated by MockGen. DO NOT EDIT.
// Source: main.go

// Package mock_main is a generated GoMock package.
package mock_main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICommander is a mock of ICommander interface
type MockICommander struct {
	ctrl     *gomock.Controller
	recorder *MockICommanderMockRecorder
}

// MockICommanderMockRecorder is the mock recorder for MockICommander
type MockICommanderMockRecorder struct {
	mock *MockICommander
}

// NewMockICommander creates a new mock instance
func NewMockICommander(ctrl *gomock.Controller) *MockICommander {
	mock := &MockICommander{ctrl: ctrl}
	mock.recorder = &MockICommanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICommander) EXPECT() *MockICommanderMockRecorder {
	return m.recorder
}

// GetGitRemoteInfo mocks base method
func (m *MockICommander) GetGitRemoteInfo() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGitRemoteInfo")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGitRemoteInfo indicates an expected call of GetGitRemoteInfo
func (mr *MockICommanderMockRecorder) GetGitRemoteInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGitRemoteInfo", reflect.TypeOf((*MockICommander)(nil).GetGitRemoteInfo))
}

// Printf mocks base method
func (m *MockICommander) Printf(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Printf", msg)
}

// Printf indicates an expected call of Printf
func (mr *MockICommanderMockRecorder) Printf(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Printf", reflect.TypeOf((*MockICommander)(nil).Printf), msg)
}

// PrintErr mocks base method
func (m *MockICommander) PrintErr(msg error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PrintErr", msg)
}

// PrintErr indicates an expected call of PrintErr
func (mr *MockICommanderMockRecorder) PrintErr(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintErr", reflect.TypeOf((*MockICommander)(nil).PrintErr), msg)
}