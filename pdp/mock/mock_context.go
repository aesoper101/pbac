// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aesoper101/pbac/pdp/types (interfaces: EvalContextor)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEvalContextor is a mock of EvalContextor interface.
type MockEvalContextor struct {
	ctrl     *gomock.Controller
	recorder *MockEvalContextorMockRecorder
}

// MockEvalContextorMockRecorder is the mock recorder for MockEvalContextor.
type MockEvalContextorMockRecorder struct {
	mock *MockEvalContextor
}

// NewMockEvalContextor creates a new mock instance.
func NewMockEvalContextor(ctrl *gomock.Controller) *MockEvalContextor {
	mock := &MockEvalContextor{ctrl: ctrl}
	mock.recorder = &MockEvalContextorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvalContextor) EXPECT() *MockEvalContextorMockRecorder {
	return m.recorder
}

// GetAttribute mocks base method.
func (m *MockEvalContextor) GetAttribute(arg0 string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttribute", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetAttribute indicates an expected call of GetAttribute.
func (mr *MockEvalContextorMockRecorder) GetAttribute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttribute", reflect.TypeOf((*MockEvalContextor)(nil).GetAttribute), arg0)
}