// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dojimanetwork/gossamer/dot/rpc/subscription (interfaces: TransactionStateAPI)

// Package subscription is a generated GoMock package.
package subscription

import (
	reflect "reflect"

	types "github.com/dojimanetwork/gossamer/dot/types"
	transaction "github.com/dojimanetwork/gossamer/lib/transaction"
	gomock "github.com/golang/mock/gomock"
)

// MockTransactionStateAPI is a mock of TransactionStateAPI interface.
type MockTransactionStateAPI struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionStateAPIMockRecorder
}

// MockTransactionStateAPIMockRecorder is the mock recorder for MockTransactionStateAPI.
type MockTransactionStateAPIMockRecorder struct {
	mock *MockTransactionStateAPI
}

// NewMockTransactionStateAPI creates a new mock instance.
func NewMockTransactionStateAPI(ctrl *gomock.Controller) *MockTransactionStateAPI {
	mock := &MockTransactionStateAPI{ctrl: ctrl}
	mock.recorder = &MockTransactionStateAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionStateAPI) EXPECT() *MockTransactionStateAPIMockRecorder {
	return m.recorder
}

// FreeStatusNotifierChannel mocks base method.
func (m *MockTransactionStateAPI) FreeStatusNotifierChannel(arg0 chan transaction.Status) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FreeStatusNotifierChannel", arg0)
}

// FreeStatusNotifierChannel indicates an expected call of FreeStatusNotifierChannel.
func (mr *MockTransactionStateAPIMockRecorder) FreeStatusNotifierChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreeStatusNotifierChannel", reflect.TypeOf((*MockTransactionStateAPI)(nil).FreeStatusNotifierChannel), arg0)
}

// GetStatusNotifierChannel mocks base method.
func (m *MockTransactionStateAPI) GetStatusNotifierChannel(arg0 types.Extrinsic) chan transaction.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatusNotifierChannel", arg0)
	ret0, _ := ret[0].(chan transaction.Status)
	return ret0
}

// GetStatusNotifierChannel indicates an expected call of GetStatusNotifierChannel.
func (mr *MockTransactionStateAPIMockRecorder) GetStatusNotifierChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatusNotifierChannel", reflect.TypeOf((*MockTransactionStateAPI)(nil).GetStatusNotifierChannel), arg0)
}
