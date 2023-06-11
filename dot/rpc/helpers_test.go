// Copyright 2022 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package rpc

import (
	"testing"

	"github.com/dojimanetwork/gossamer/dot/types"
	"github.com/dojimanetwork/gossamer/lib/common"
	"github.com/dojimanetwork/gossamer/lib/genesis"
	"github.com/dojimanetwork/gossamer/lib/runtime/wasmer"
	"github.com/dojimanetwork/gossamer/lib/trie"
	"github.com/dojimanetwork/gossamer/lib/utils"
	"github.com/stretchr/testify/require"
)

func newWestendDevGenesisWithTrieAndHeader(t *testing.T) (
	gen genesis.Genesis, genesisTrie trie.Trie, genesisHeader types.Header) {
	t.Helper()

	genesisPath := utils.GetWestendDevRawGenesisPath(t)
	genesisPtr, err := genesis.NewGenesisFromJSONRaw(genesisPath)
	require.NoError(t, err)
	gen = *genesisPtr

	genesisTrie, err = wasmer.NewTrieFromGenesis(gen)
	require.NoError(t, err)

	parentHash := common.NewHash([]byte{0})
	stateRoot := genesisTrie.MustHash()
	extrinsicRoot := trie.EmptyHash
	const number = 0
	digest := types.NewDigest()
	genesisHeader = *types.NewHeader(parentHash,
		stateRoot, extrinsicRoot, number, digest)

	return gen, genesisTrie, genesisHeader
}
