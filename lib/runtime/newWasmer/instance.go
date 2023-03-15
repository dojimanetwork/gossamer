// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package newWasmer

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ChainSafe/gossamer/internal/log"
	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/ChainSafe/gossamer/lib/crypto"
	"github.com/ChainSafe/gossamer/lib/keystore"
	"github.com/ChainSafe/gossamer/lib/runtime"
	"github.com/ChainSafe/gossamer/lib/runtime/offchain"
	"github.com/ChainSafe/gossamer/lib/trie"
	"github.com/klauspost/compress/zstd"
	"github.com/wasmerio/wasmer-go/wasmer"
	"os"
	"path/filepath"
	"sync"
)

// Name represents the name of the interpreter
const Name = "wasmer"

var (
	logger = log.NewFromGlobal(
		log.AddContext("pkg", "runtime"),
		log.AddContext("module", "go-wasmer"),
	)
)

var (
	ErrCodeEmpty      = errors.New("code is empty")
	ErrWASMDecompress = errors.New("wasm decompression failed")
)

// Context is the context for the wasm interpreter's imported functions
type Context struct {
	Storage         Storage
	Allocator       *runtime.FreeingBumpHeapAllocator
	Keystore        *keystore.GlobalKeystore
	Validator       bool
	NodeStorage     runtime.NodeStorage
	Network         BasicNetwork
	Transaction     TransactionState
	SigVerifier     *crypto.SignatureVerifier
	OffchainHTTPSet *offchain.HTTPSet
	Version         runtime.Version
	Memory          Memory
}

// Instance represents a runtime go-wasmer instance
type Instance struct {
	vm       wasmer.Instance
	ctx      *Context
	isClosed bool
	codeHash common.Hash
	mutex    sync.Mutex
}

// NewRuntimeFromGenesis creates a runtime instance from the genesis data
func NewRuntimeFromGenesis(cfg Config) (instance *Instance, err error) {
	if cfg.Storage == nil {
		return nil, errors.New("storage is nil")
	}

	code := cfg.Storage.LoadCode()
	if len(code) == 0 {
		return nil, fmt.Errorf("cannot find :code in state")
	}

	return NewInstance(code, cfg)
}

// NewInstanceFromTrie returns a new runtime instance with the code provided in the given trie
func NewInstanceFromTrie(t *trie.Trie, cfg Config) (*Instance, error) {
	code := t.Get(common.CodeKey)
	if len(code) == 0 {
		return nil, fmt.Errorf("cannot find :code in trie")
	}

	return NewInstance(code, cfg)
}

// NewInstanceFromFile instantiates a runtime from a .wasm file
func NewInstanceFromFile(fp string, cfg Config) (*Instance, error) {
	// Reads the WebAssembly module as bytes.
	bytes, err := os.ReadFile(filepath.Clean(fp))
	if err != nil {
		return nil, err
	}

	return NewInstance(bytes, cfg)
}

// NewInstance instantiates a runtime from raw wasm bytecode
// TODO should cfg be a pointer?
func NewInstance(code []byte, cfg Config) (*Instance, error) {
	return newInstance(code, cfg)
}

// TODO refactor
func newInstance(code []byte, cfg Config) (*Instance, error) {
	logger.Patch(log.SetLevel(cfg.LogLvl), log.SetCallerFunc(true))
	if len(code) == 0 {
		return nil, ErrCodeEmpty
	}

	code, err := decompressWasm(code)
	if err != nil {
		// Note the sentinel error is wrapped here since the ztsd Go library
		// does not return any exported sentinel errors.
		return nil, fmt.Errorf("%w: %s", ErrWASMDecompress, err)
	}

	// TODO add new get imports function
	//imports, err := importsNodeRuntime()
	var imports *wasmer.ImportObject
	if err != nil {
		return nil, fmt.Errorf("creating node runtime imports: %w", err)
	}

	// Create engine and store with default values
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compile the module
	module, err := wasmer.NewModule(store, code)
	if err != nil {
		return nil, err
	}

	// Get memory descriptor from module, if it imports memory
	moduleImports := module.Imports()
	var memImport *wasmer.ImportType
	for _, im := range moduleImports {
		if im.Name() == "memory" {
			memImport = im
			break
		}
	}

	var memoryType *wasmer.MemoryType
	if memImport != nil {
		memoryType = memImport.Type().IntoMemoryType()
	}

	// Check if module exports memory
	hasExportedMemory := false
	moduleExports := module.Exports()
	for _, export := range moduleExports {
		if export.Name() == "memory" {
			hasExportedMemory = true
			break
		}
	}

	var memory *wasmer.Memory
	// create memory to import, if it's expecting imported memory
	if !hasExportedMemory {
		if memoryType == nil {
			// values from newer kusama/polkadot runtimes
			lim, err := wasmer.NewLimits(23, 4294967295) //nolint
			if err != nil {
				return nil, err
			}
			memoryType = wasmer.NewMemoryType(lim)
		}

		memory = wasmer.NewMemory(store, memoryType)
	}

	runtimeCtx := &Context{
		Storage:         cfg.Storage,
		Keystore:        cfg.Keystore,
		Validator:       cfg.Role == common.AuthorityRole,
		NodeStorage:     cfg.NodeStorage,
		Network:         cfg.Network,
		Transaction:     cfg.Transaction,
		SigVerifier:     crypto.NewSignatureVerifier(logger),
		OffchainHTTPSet: offchain.NewHTTPSet(),
	}

	wasmInstance, err := wasmer.NewInstance(module, imports)
	if err != nil {
		return nil, err
	}

	logger.Info("instantiated runtime!!!")

	if hasExportedMemory {
		memory, err = wasmInstance.Exports.GetMemory("memory")
		if err != nil {
			return nil, err
		}
	}

	runtimeCtx.Memory = Memory{memory}

	// set heap base for allocator, start allocating at heap base
	heapBase, err := wasmInstance.Exports.Get("__heap_base")
	if err != nil {
		return nil, err
	}

	hb, err := heapBase.IntoGlobal().Get()
	if err != nil {
		return nil, err
	}

	runtimeCtx.Allocator = runtime.NewAllocator(runtimeCtx.Memory, uint32(hb.(int32)))
	instance := &Instance{
		vm:       *wasmInstance,
		ctx:      runtimeCtx,
		codeHash: cfg.CodeHash,
	}

	// TODO this should work when we bring in exports
	//if cfg.testVersion != nil {
	//	instance.ctx.Version = *cfg.testVersion
	//} else {
	//	instance.ctx.Version, err = instance.version()
	//	if err != nil {
	//		instance.close()
	//		return nil, fmt.Errorf("getting instance version: %w", err)
	//	}
	//}
	return instance, nil
}

// decompressWasm decompresses a Wasm blob that may or may not be compressed with zstd
// ref: https://github.com/paritytech/substrate/blob/master/primitives/maybe-compressed-blob/src/lib.rs
func decompressWasm(code []byte) ([]byte, error) {
	compressionFlag := []byte{82, 188, 83, 118, 70, 219, 142, 5}
	if !bytes.HasPrefix(code, compressionFlag) {
		return code, nil
	}

	decoder, err := zstd.NewReader(nil)
	if err != nil {
		return nil, fmt.Errorf("creating zstd reader: %s", err)
	}

	return decoder.DecodeAll(code[len(compressionFlag):], nil)
}
