// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dojimanetwork/gossamer/dot/core (interfaces: CodeSubstitutedState)

// Package modules is a generated GoMock package.
package modules

import (
	reflect "reflect"

	common "github.com/dojimanetwork/gossamer/lib/common"
	gomock "github.com/golang/mock/gomock"
)

// MockCodeSubstitutedState is a mock of CodeSubstitutedState interface.
type MockCodeSubstitutedState struct {
	ctrl     *gomock.Controller
	recorder *MockCodeSubstitutedStateMockRecorder
}

// MockCodeSubstitutedStateMockRecorder is the mock recorder for MockCodeSubstitutedState.
type MockCodeSubstitutedStateMockRecorder struct {
	mock *MockCodeSubstitutedState
}

// NewMockCodeSubstitutedState creates a new mock instance.
func NewMockCodeSubstitutedState(ctrl *gomock.Controller) *MockCodeSubstitutedState {
	mock := &MockCodeSubstitutedState{ctrl: ctrl}
	mock.recorder = &MockCodeSubstitutedStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCodeSubstitutedState) EXPECT() *MockCodeSubstitutedStateMockRecorder {
	return m.recorder
}

// StoreCodeSubstitutedBlockHash mocks base method.
func (m *MockCodeSubstitutedState) StoreCodeSubstitutedBlockHash(arg0 common.Hash) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreCodeSubstitutedBlockHash", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreCodeSubstitutedBlockHash indicates an expected call of StoreCodeSubstitutedBlockHash.
func (mr *MockCodeSubstitutedStateMockRecorder) StoreCodeSubstitutedBlockHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreCodeSubstitutedBlockHash", reflect.TypeOf((*MockCodeSubstitutedState)(nil).StoreCodeSubstitutedBlockHash), arg0)
}
