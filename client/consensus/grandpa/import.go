// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package grandpa

import (
	"github.com/ChainSafe/gossamer/client/consensus"
	"github.com/ChainSafe/gossamer/dot/types"
	"github.com/ChainSafe/gossamer/lib/common"
)

// GrandpaBlockImport A block-import handler for GRANDPA.
//
// This scans each imported block for signals of changing authority set.
// If the block being imported enacts an authority set change then:
// - If the current authority set is still live: we import the block
// - Otherwise, the block must include a valid justification.
//
// When using GRANDPA, the block import worker should be using this block import
// object.
type GrandpaBlockImport struct {
	inner        ClientForGrandpa
	selectChain  SelectChain
	authoritySet SharedAuthoritySet
	// TODO figure out best way to do this in go
	//sendVoterCommands TracingUnboundedSender
	// TODO investigate this
	// It looks like grandpa changes are defined in state, why? Feel like best here
	authoritySetHardForks map[common.Hash]consensus.PendingChange
	// TODO more communication stuff
	//justification_sender GrandpaJustificationSender
	telemetry Telemetry
	// TODO do we need backend? Do we ned PhantomData
	//phantom PhantomData<Backend>
}

// I think I can account for diff trait impls just by defining appropriate types
// Maybe can also use nested interfaces?

// JustificationImport impl, async trait

func (gbi *GrandpaBlockImport) onStart() ([]types.GrandpaVote, error) {
	// TODO implement
	return []types.GrandpaVote{}, nil
}

// TODO impl justification
func (gbi *GrandpaBlockImport) importJustification(hash common.Hash, number uint) {
	// TODO implement
	// this justification was requested by the sync service, therefore we
	// are not sure if it should enact a change or not. it could have been a
	// request made as part of initial sync but that means the justification
	// wasn't part of the block and was requested asynchronously, probably
	// makes sense to log in that case.
	justification := []byte{} // dummy data
	gbi.importJustification2(hash, number, justification, false, false)
}

// GrandpaBlockImport impl

// check for a new authority set change.
func (gbi *GrandpaBlockImport) checkNewChange(header *types.Header, hash common.Hash) {
	// TODO implement
}

func (gbi *GrandpaBlockImport) makeAuthoritiesChanges() {
	// TODO implement
}

// Read current set id form a given state.
func (gbi *GrandpaBlockImport) currentSetId(hash common.Hash) (setId uint64, err error) {
	// TODO implement
	return 0, nil
}

// Import whole new state and reset authority set.
func (gbi *GrandpaBlockImport) importState() {
	// TODO implement
}

// BlockImport impl, async trait

func (gbi *GrandpaBlockImport) importBlock() {
	// TODO implement
}

func (gbi *GrandpaBlockImport) checkBlock() {
	// TODO implement
}

// GrandpaBlockImport impl
func (gbi *GrandpaBlockImport) new() {
	// TODO implement
}

// Import a block justification and finalize the block.
//
// If `enacts_change` is set to true, then finalizing this block *must*
// enact an authority set change, the function will panic otherwise.
func (gbi *GrandpaBlockImport) importJustification2(hash common.Hash, number uint, justification []byte, enactsChange bool, initialSync bool) {
	// TODO implement
}
