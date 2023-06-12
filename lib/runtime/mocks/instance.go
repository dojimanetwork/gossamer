// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	common "github.com/dojimanetwork/gossamer/lib/common"
	keystore "github.com/dojimanetwork/gossamer/lib/keystore"

	mock "github.com/stretchr/testify/mock"

	runtime "github.com/dojimanetwork/gossamer/lib/runtime"

	transaction "github.com/dojimanetwork/gossamer/lib/transaction"

	types "github.com/dojimanetwork/gossamer/dot/types"
)

// Instance is an autogenerated mock type for the Instance type
type Instance struct {
	mock.Mock
}

// ApplyExtrinsic provides a mock function with given fields: data
func (_m *Instance) ApplyExtrinsic(data types.Extrinsic) ([]byte, error) {
	ret := _m.Called(data)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(types.Extrinsic) []byte); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Extrinsic) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BabeConfiguration provides a mock function with given fields:
func (_m *Instance) BabeConfiguration() (*types.BabeConfiguration, error) {
	ret := _m.Called()

	var r0 *types.BabeConfiguration
	if rf, ok := ret.Get(0).(func() *types.BabeConfiguration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BabeConfiguration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckInherents provides a mock function with given fields:
func (_m *Instance) CheckInherents() {
	_m.Called()
}

// CheckRuntimeVersion provides a mock function with given fields: _a0
func (_m *Instance) CheckRuntimeVersion(_a0 []byte) (runtime.Version, error) {
	ret := _m.Called(_a0)

	var r0 runtime.Version
	if rf, ok := ret.Get(0).(func([]byte) runtime.Version); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeSessionKeys provides a mock function with given fields: enc
func (_m *Instance) DecodeSessionKeys(enc []byte) ([]byte, error) {
	ret := _m.Called(enc)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(enc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(enc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exec provides a mock function with given fields: function, data
func (_m *Instance) Exec(function string, data []byte) ([]byte, error) {
	ret := _m.Called(function, data)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, []byte) []byte); ok {
		r0 = rf(function, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte) error); ok {
		r1 = rf(function, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteBlock provides a mock function with given fields: block
func (_m *Instance) ExecuteBlock(block *types.Block) ([]byte, error) {
	ret := _m.Called(block)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*types.Block) []byte); ok {
		r0 = rf(block)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Block) error); ok {
		r1 = rf(block)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FinalizeBlock provides a mock function with given fields:
func (_m *Instance) FinalizeBlock() (*types.Header, error) {
	ret := _m.Called()

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func() *types.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateSessionKeys provides a mock function with given fields:
func (_m *Instance) GenerateSessionKeys() {
	_m.Called()
}

// GetCodeHash provides a mock function with given fields:
func (_m *Instance) GetCodeHash() common.Hash {
	ret := _m.Called()

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func() common.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// GrandpaAuthorities provides a mock function with given fields:
func (_m *Instance) GrandpaAuthorities() ([]types.Authority, error) {
	ret := _m.Called()

	var r0 []types.Authority
	if rf, ok := ret.Get(0).(func() []types.Authority); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Authority)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InherentExtrinsics provides a mock function with given fields: data
func (_m *Instance) InherentExtrinsics(data []byte) ([]byte, error) {
	ret := _m.Called(data)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitializeBlock provides a mock function with given fields: header
func (_m *Instance) InitializeBlock(header *types.Header) error {
	ret := _m.Called(header)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Header) error); ok {
		r0 = rf(header)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Keystore provides a mock function with given fields:
func (_m *Instance) Keystore() *keystore.GlobalKeystore {
	ret := _m.Called()

	var r0 *keystore.GlobalKeystore
	if rf, ok := ret.Get(0).(func() *keystore.GlobalKeystore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keystore.GlobalKeystore)
		}
	}

	return r0
}

// Metadata provides a mock function with given fields:
func (_m *Instance) Metadata() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NetworkService provides a mock function with given fields:
func (_m *Instance) NetworkService() runtime.BasicNetwork {
	ret := _m.Called()

	var r0 runtime.BasicNetwork
	if rf, ok := ret.Get(0).(func() runtime.BasicNetwork); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.BasicNetwork)
		}
	}

	return r0
}

// NodeStorage provides a mock function with given fields:
func (_m *Instance) NodeStorage() runtime.NodeStorage {
	ret := _m.Called()

	var r0 runtime.NodeStorage
	if rf, ok := ret.Get(0).(func() runtime.NodeStorage); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(runtime.NodeStorage)
	}

	return r0
}

// OffchainWorker provides a mock function with given fields:
func (_m *Instance) OffchainWorker() {
	_m.Called()
}

// PaymentQueryInfo provides a mock function with given fields: ext
func (_m *Instance) PaymentQueryInfo(ext []byte) (*types.TransactionPaymentQueryInfo, error) {
	ret := _m.Called(ext)

	var r0 *types.TransactionPaymentQueryInfo
	if rf, ok := ret.Get(0).(func([]byte) *types.TransactionPaymentQueryInfo); ok {
		r0 = rf(ext)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TransactionPaymentQueryInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(ext)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RandomSeed provides a mock function with given fields:
func (_m *Instance) RandomSeed() {
	_m.Called()
}

// SetContextStorage provides a mock function with given fields: s
func (_m *Instance) SetContextStorage(s runtime.Storage) {
	_m.Called(s)
}

// Stop provides a mock function with given fields:
func (_m *Instance) Stop() {
	_m.Called()
}

// UpdateRuntimeCode provides a mock function with given fields: _a0
func (_m *Instance) UpdateRuntimeCode(_a0 []byte) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateTransaction provides a mock function with given fields: e
func (_m *Instance) ValidateTransaction(e types.Extrinsic) (*transaction.Validity, error) {
	ret := _m.Called(e)

	var r0 *transaction.Validity
	if rf, ok := ret.Get(0).(func(types.Extrinsic) *transaction.Validity); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transaction.Validity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Extrinsic) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validator provides a mock function with given fields:
func (_m *Instance) Validator() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Version provides a mock function with given fields:
func (_m *Instance) Version() (runtime.Version, error) {
	ret := _m.Called()

	var r0 runtime.Version
	if rf, ok := ret.Get(0).(func() runtime.Version); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
