// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dojimanetwork/gossamer/dot/rpc/modules (interfaces: StorageAPI,BlockAPI,Telemetry)

// Package modules is a generated GoMock package.
package modules

import (
	json "encoding/json"
	reflect "reflect"

	state "github.com/dojimanetwork/gossamer/dot/state"
	types "github.com/dojimanetwork/gossamer/dot/types"
	common "github.com/dojimanetwork/gossamer/lib/common"
	runtime "github.com/dojimanetwork/gossamer/lib/runtime"
	trie "github.com/dojimanetwork/gossamer/lib/trie"
	gomock "github.com/golang/mock/gomock"
)

// MockStorageAPI is a mock of StorageAPI interface.
type MockStorageAPI struct {
	ctrl     *gomock.Controller
	recorder *MockStorageAPIMockRecorder
}

// MockStorageAPIMockRecorder is the mock recorder for MockStorageAPI.
type MockStorageAPIMockRecorder struct {
	mock *MockStorageAPI
}

// NewMockStorageAPI creates a new mock instance.
func NewMockStorageAPI(ctrl *gomock.Controller) *MockStorageAPI {
	mock := &MockStorageAPI{ctrl: ctrl}
	mock.recorder = &MockStorageAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageAPI) EXPECT() *MockStorageAPIMockRecorder {
	return m.recorder
}

// Entries mocks base method.
func (m *MockStorageAPI) Entries(arg0 *common.Hash) (map[string][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entries", arg0)
	ret0, _ := ret[0].(map[string][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Entries indicates an expected call of Entries.
func (mr *MockStorageAPIMockRecorder) Entries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entries", reflect.TypeOf((*MockStorageAPI)(nil).Entries), arg0)
}

// GetKeysWithPrefix mocks base method.
func (m *MockStorageAPI) GetKeysWithPrefix(arg0 *common.Hash, arg1 []byte) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeysWithPrefix", arg0, arg1)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeysWithPrefix indicates an expected call of GetKeysWithPrefix.
func (mr *MockStorageAPIMockRecorder) GetKeysWithPrefix(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeysWithPrefix", reflect.TypeOf((*MockStorageAPI)(nil).GetKeysWithPrefix), arg0, arg1)
}

// GetStateRootFromBlock mocks base method.
func (m *MockStorageAPI) GetStateRootFromBlock(arg0 *common.Hash) (*common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateRootFromBlock", arg0)
	ret0, _ := ret[0].(*common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateRootFromBlock indicates an expected call of GetStateRootFromBlock.
func (mr *MockStorageAPIMockRecorder) GetStateRootFromBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateRootFromBlock", reflect.TypeOf((*MockStorageAPI)(nil).GetStateRootFromBlock), arg0)
}

// GetStorage mocks base method.
func (m *MockStorageAPI) GetStorage(arg0 *common.Hash, arg1 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorage", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorage indicates an expected call of GetStorage.
func (mr *MockStorageAPIMockRecorder) GetStorage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorage", reflect.TypeOf((*MockStorageAPI)(nil).GetStorage), arg0, arg1)
}

// GetStorageByBlockHash mocks base method.
func (m *MockStorageAPI) GetStorageByBlockHash(arg0 *common.Hash, arg1 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageByBlockHash", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageByBlockHash indicates an expected call of GetStorageByBlockHash.
func (mr *MockStorageAPIMockRecorder) GetStorageByBlockHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageByBlockHash", reflect.TypeOf((*MockStorageAPI)(nil).GetStorageByBlockHash), arg0, arg1)
}

// GetStorageChild mocks base method.
func (m *MockStorageAPI) GetStorageChild(arg0 *common.Hash, arg1 []byte) (*trie.Trie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageChild", arg0, arg1)
	ret0, _ := ret[0].(*trie.Trie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageChild indicates an expected call of GetStorageChild.
func (mr *MockStorageAPIMockRecorder) GetStorageChild(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageChild", reflect.TypeOf((*MockStorageAPI)(nil).GetStorageChild), arg0, arg1)
}

// GetStorageFromChild mocks base method.
func (m *MockStorageAPI) GetStorageFromChild(arg0 *common.Hash, arg1, arg2 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageFromChild", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageFromChild indicates an expected call of GetStorageFromChild.
func (mr *MockStorageAPIMockRecorder) GetStorageFromChild(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageFromChild", reflect.TypeOf((*MockStorageAPI)(nil).GetStorageFromChild), arg0, arg1, arg2)
}

// RegisterStorageObserver mocks base method.
func (m *MockStorageAPI) RegisterStorageObserver(arg0 state.Observer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterStorageObserver", arg0)
}

// RegisterStorageObserver indicates an expected call of RegisterStorageObserver.
func (mr *MockStorageAPIMockRecorder) RegisterStorageObserver(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterStorageObserver", reflect.TypeOf((*MockStorageAPI)(nil).RegisterStorageObserver), arg0)
}

// UnregisterStorageObserver mocks base method.
func (m *MockStorageAPI) UnregisterStorageObserver(arg0 state.Observer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnregisterStorageObserver", arg0)
}

// UnregisterStorageObserver indicates an expected call of UnregisterStorageObserver.
func (mr *MockStorageAPIMockRecorder) UnregisterStorageObserver(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterStorageObserver", reflect.TypeOf((*MockStorageAPI)(nil).UnregisterStorageObserver), arg0)
}

// MockBlockAPI is a mock of BlockAPI interface.
type MockBlockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockBlockAPIMockRecorder
}

// MockBlockAPIMockRecorder is the mock recorder for MockBlockAPI.
type MockBlockAPIMockRecorder struct {
	mock *MockBlockAPI
}

// NewMockBlockAPI creates a new mock instance.
func NewMockBlockAPI(ctrl *gomock.Controller) *MockBlockAPI {
	mock := &MockBlockAPI{ctrl: ctrl}
	mock.recorder = &MockBlockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockAPI) EXPECT() *MockBlockAPIMockRecorder {
	return m.recorder
}

// BestBlockHash mocks base method.
func (m *MockBlockAPI) BestBlockHash() common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BestBlockHash")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// BestBlockHash indicates an expected call of BestBlockHash.
func (mr *MockBlockAPIMockRecorder) BestBlockHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BestBlockHash", reflect.TypeOf((*MockBlockAPI)(nil).BestBlockHash))
}

// FreeFinalisedNotifierChannel mocks base method.
func (m *MockBlockAPI) FreeFinalisedNotifierChannel(arg0 chan *types.FinalisationInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FreeFinalisedNotifierChannel", arg0)
}

// FreeFinalisedNotifierChannel indicates an expected call of FreeFinalisedNotifierChannel.
func (mr *MockBlockAPIMockRecorder) FreeFinalisedNotifierChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreeFinalisedNotifierChannel", reflect.TypeOf((*MockBlockAPI)(nil).FreeFinalisedNotifierChannel), arg0)
}

// FreeImportedBlockNotifierChannel mocks base method.
func (m *MockBlockAPI) FreeImportedBlockNotifierChannel(arg0 chan *types.Block) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FreeImportedBlockNotifierChannel", arg0)
}

// FreeImportedBlockNotifierChannel indicates an expected call of FreeImportedBlockNotifierChannel.
func (mr *MockBlockAPIMockRecorder) FreeImportedBlockNotifierChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreeImportedBlockNotifierChannel", reflect.TypeOf((*MockBlockAPI)(nil).FreeImportedBlockNotifierChannel), arg0)
}

// GetBlockByHash mocks base method.
func (m *MockBlockAPI) GetBlockByHash(arg0 common.Hash) (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByHash", arg0)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHash indicates an expected call of GetBlockByHash.
func (mr *MockBlockAPIMockRecorder) GetBlockByHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHash", reflect.TypeOf((*MockBlockAPI)(nil).GetBlockByHash), arg0)
}

// GetFinalisedHash mocks base method.
func (m *MockBlockAPI) GetFinalisedHash(arg0, arg1 uint64) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalisedHash", arg0, arg1)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFinalisedHash indicates an expected call of GetFinalisedHash.
func (mr *MockBlockAPIMockRecorder) GetFinalisedHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalisedHash", reflect.TypeOf((*MockBlockAPI)(nil).GetFinalisedHash), arg0, arg1)
}

// GetFinalisedNotifierChannel mocks base method.
func (m *MockBlockAPI) GetFinalisedNotifierChannel() chan *types.FinalisationInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalisedNotifierChannel")
	ret0, _ := ret[0].(chan *types.FinalisationInfo)
	return ret0
}

// GetFinalisedNotifierChannel indicates an expected call of GetFinalisedNotifierChannel.
func (mr *MockBlockAPIMockRecorder) GetFinalisedNotifierChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalisedNotifierChannel", reflect.TypeOf((*MockBlockAPI)(nil).GetFinalisedNotifierChannel))
}

// GetHashByNumber mocks base method.
func (m *MockBlockAPI) GetHashByNumber(arg0 uint) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHashByNumber", arg0)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHashByNumber indicates an expected call of GetHashByNumber.
func (mr *MockBlockAPIMockRecorder) GetHashByNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHashByNumber", reflect.TypeOf((*MockBlockAPI)(nil).GetHashByNumber), arg0)
}

// GetHeader mocks base method.
func (m *MockBlockAPI) GetHeader(arg0 common.Hash) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", arg0)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeader indicates an expected call of GetHeader.
func (mr *MockBlockAPIMockRecorder) GetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockBlockAPI)(nil).GetHeader), arg0)
}

// GetHighestFinalisedHash mocks base method.
func (m *MockBlockAPI) GetHighestFinalisedHash() (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHighestFinalisedHash")
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHighestFinalisedHash indicates an expected call of GetHighestFinalisedHash.
func (mr *MockBlockAPIMockRecorder) GetHighestFinalisedHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHighestFinalisedHash", reflect.TypeOf((*MockBlockAPI)(nil).GetHighestFinalisedHash))
}

// GetImportedBlockNotifierChannel mocks base method.
func (m *MockBlockAPI) GetImportedBlockNotifierChannel() chan *types.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImportedBlockNotifierChannel")
	ret0, _ := ret[0].(chan *types.Block)
	return ret0
}

// GetImportedBlockNotifierChannel indicates an expected call of GetImportedBlockNotifierChannel.
func (mr *MockBlockAPIMockRecorder) GetImportedBlockNotifierChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImportedBlockNotifierChannel", reflect.TypeOf((*MockBlockAPI)(nil).GetImportedBlockNotifierChannel))
}

// GetJustification mocks base method.
func (m *MockBlockAPI) GetJustification(arg0 common.Hash) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJustification", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJustification indicates an expected call of GetJustification.
func (mr *MockBlockAPIMockRecorder) GetJustification(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJustification", reflect.TypeOf((*MockBlockAPI)(nil).GetJustification), arg0)
}

// GetRuntime mocks base method.
func (m *MockBlockAPI) GetRuntime(arg0 common.Hash) (runtime.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRuntime", arg0)
	ret0, _ := ret[0].(runtime.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRuntime indicates an expected call of GetRuntime.
func (mr *MockBlockAPIMockRecorder) GetRuntime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuntime", reflect.TypeOf((*MockBlockAPI)(nil).GetRuntime), arg0)
}

// HasJustification mocks base method.
func (m *MockBlockAPI) HasJustification(arg0 common.Hash) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasJustification", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasJustification indicates an expected call of HasJustification.
func (mr *MockBlockAPIMockRecorder) HasJustification(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasJustification", reflect.TypeOf((*MockBlockAPI)(nil).HasJustification), arg0)
}

// RangeInMemory mocks base method.
func (m *MockBlockAPI) RangeInMemory(arg0, arg1 common.Hash) ([]common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RangeInMemory", arg0, arg1)
	ret0, _ := ret[0].([]common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RangeInMemory indicates an expected call of RangeInMemory.
func (mr *MockBlockAPIMockRecorder) RangeInMemory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RangeInMemory", reflect.TypeOf((*MockBlockAPI)(nil).RangeInMemory), arg0, arg1)
}

// RegisterRuntimeUpdatedChannel mocks base method.
func (m *MockBlockAPI) RegisterRuntimeUpdatedChannel(arg0 chan<- runtime.Version) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterRuntimeUpdatedChannel", arg0)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterRuntimeUpdatedChannel indicates an expected call of RegisterRuntimeUpdatedChannel.
func (mr *MockBlockAPIMockRecorder) RegisterRuntimeUpdatedChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRuntimeUpdatedChannel", reflect.TypeOf((*MockBlockAPI)(nil).RegisterRuntimeUpdatedChannel), arg0)
}

// UnregisterRuntimeUpdatedChannel mocks base method.
func (m *MockBlockAPI) UnregisterRuntimeUpdatedChannel(arg0 uint32) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterRuntimeUpdatedChannel", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UnregisterRuntimeUpdatedChannel indicates an expected call of UnregisterRuntimeUpdatedChannel.
func (mr *MockBlockAPIMockRecorder) UnregisterRuntimeUpdatedChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterRuntimeUpdatedChannel", reflect.TypeOf((*MockBlockAPI)(nil).UnregisterRuntimeUpdatedChannel), arg0)
}

// MockTelemetry is a mock of Telemetry interface.
type MockTelemetry struct {
	ctrl     *gomock.Controller
	recorder *MockTelemetryMockRecorder
}

// MockTelemetryMockRecorder is the mock recorder for MockTelemetry.
type MockTelemetryMockRecorder struct {
	mock *MockTelemetry
}

// NewMockTelemetry creates a new mock instance.
func NewMockTelemetry(ctrl *gomock.Controller) *MockTelemetry {
	mock := &MockTelemetry{ctrl: ctrl}
	mock.recorder = &MockTelemetryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelemetry) EXPECT() *MockTelemetryMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockTelemetry) SendMessage(arg0 json.Marshaler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendMessage", arg0)
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockTelemetryMockRecorder) SendMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockTelemetry)(nil).SendMessage), arg0)
}
