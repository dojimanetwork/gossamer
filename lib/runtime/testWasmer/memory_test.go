package testWasmer

import (
	"github.com/stretchr/testify/require"
	"github.com/wasmerio/wasmer-go/wasmer"
	"testing"
	"unsafe"
)

func createInstance(t *testing.T) (*wasmer.Instance, error) {
	// We are using the text representation of the module here.
	// Taken from: https://github.com/wasmerio/wasmer-go/blob/23d786b6b81ad93e2b974d2f4510bea374f0f37c/examples/example_memory_test.go#L28 //nolint:lll
	wasmBytes := []byte(`
		(module
		  (type $mem_size_t (func (result i32)))
		  (type $get_at_t (func (param i32) (result i32)))
		  (type $set_at_t (func (param i32) (param i32)))
		  (memory $mem 1)
		  (func $get_at (type $get_at_t) (param $idx i32) (result i32)
		    (i32.load (local.get $idx)))
		  (func $set_at (type $set_at_t) (param $idx i32) (param $val i32)
		    (i32.store (local.get $idx) (local.get $val)))
		  (func $mem_size (type $mem_size_t) (result i32)
		    (memory.size))
		  (export "get_at" (func $get_at))
		  (export "set_at" (func $set_at))
		  (export "mem_size" (func $mem_size))
		  (export "memory" (memory $mem)))
	`)

	wasmerConfig := wasmer.NewConfig()
	wasmerConfig.UseSinglepassCompiler()

	engine := wasmer.NewEngineWithConfig(wasmerConfig)
	store := wasmer.NewStore(engine)

	// Compile module
	module, err := wasmer.NewModule(store, wasmBytes)
	require.NoError(t, err)

	importObject := wasmer.NewImportObject()

	// Instantiate the Wasm module.
	return wasmer.NewInstance(module, importObject)
}

func TestMemory_Length(t *testing.T) {
	const pageLength uint32 = 65536
	instance, err := createInstance(t)
	require.NoError(t, err)

	wasmerMemory, err := instance.Exports.GetMemory("memory")
	require.NoError(t, err)

	memory := Memory{
		memory: wasmerMemory,
	}

	memLength := memory.Length()
	require.Equal(t, pageLength, memLength)
}

func TestMemory_Grow(t *testing.T) {
	const pageLength uint32 = 65536
	instance, err := createInstance(t)
	require.NoError(t, err)

	wasmerMemory, err := instance.Exports.GetMemory("memory")
	require.NoError(t, err)

	memory := Memory{
		memory: wasmerMemory,
	}

	memLength := memory.Length()
	require.Equal(t, pageLength, memLength)

	err = memory.Grow(1)
	require.NoError(t, err)

	memLength = memory.Length()
	require.Equal(t, pageLength*2, memLength)
}

func TestMemory_Data(t *testing.T) {
	const pageLength uint32 = 65536
	instance, err := createInstance(t)
	require.NoError(t, err)

	// Grab exported utility functions from the module
	getAt, err := instance.Exports.GetFunction("get_at")
	require.NoError(t, err)

	setAt, err := instance.Exports.GetFunction("set_at")
	require.NoError(t, err)

	wasmerMemory, err := instance.Exports.GetMemory("memory")
	require.NoError(t, err)

	memory := Memory{
		memory: wasmerMemory,
	}

	memAddr := 0x0
	const val int32 = 0xFEFEFFE
	_, err = setAt(memAddr, val)
	require.NoError(t, err)

	// Compare bytes at address 0x0
	expectedFirstBytes := []byte{254, 239, 239, 15}
	memData := memory.Data()
	require.Equal(t, expectedFirstBytes, memData[:4])

	result, err := getAt(memAddr)
	require.NoError(t, err)
	require.Equal(t, val, result)

	// Write value at end of page
	pageSize := 0x1_0000
	memAddr = (pageSize) - int(unsafe.Sizeof(val))
	const val2 int32 = 0xFEA09
	_, err = setAt(memAddr, val2)
	require.NoError(t, err)

	result, err = getAt(memAddr)
	require.NoError(t, err)
	require.Equal(t, val2, result)

}
