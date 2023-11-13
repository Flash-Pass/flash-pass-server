// Code generated by MockGen. DO NOT EDIT.
// Source: snowflake.go

// Package SnowflakeMocks is a generated GoMock package.
package SnowflakeMocks

import (
	reflect "reflect"

	snowflake "github.com/bwmarrin/snowflake"
	gomock "github.com/golang/mock/gomock"
)

// MockIHandle is a mock of IHandle interface.
type MockIHandle struct {
	ctrl     *gomock.Controller
	recorder *MockIHandleMockRecorder
}

// MockIHandleMockRecorder is the mock recorder for MockIHandle.
type MockIHandleMockRecorder struct {
	mock *MockIHandle
}

// NewMockIHandle creates a new mock instance.
func NewMockIHandle(ctrl *gomock.Controller) *MockIHandle {
	mock := &MockIHandle{ctrl: ctrl}
	mock.recorder = &MockIHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHandle) EXPECT() *MockIHandleMockRecorder {
	return m.recorder
}

// GetId mocks base method.
func (m *MockIHandle) GetId() snowflake.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetId")
	ret0, _ := ret[0].(snowflake.ID)
	return ret0
}

// GetId indicates an expected call of GetId.
func (mr *MockIHandleMockRecorder) GetId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetId", reflect.TypeOf((*MockIHandle)(nil).GetId))
}

// GetUInt64Id mocks base method.
func (m *MockIHandle) GetUInt64Id() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUInt64Id")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetUInt64Id indicates an expected call of GetUInt64Id.
func (mr *MockIHandleMockRecorder) GetUInt64Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUInt64Id", reflect.TypeOf((*MockIHandle)(nil).GetUInt64Id))
}