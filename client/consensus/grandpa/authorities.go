// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package grandpa

import (
	"github.com/ChainSafe/gossamer/dot/types"
)

// AuthorityList A list of Grandpa authorities with associated weights.
type AuthorityList []types.Authority

// PendingChange A pending change to the authority set.
//
// This will be applied when the announcing block is at some depth within
// the finalized or unfinalized chain.
type PendingChange struct {
	// TODO impl
}

// AuthoritySetChanges Tracks historical authority set changes. We store the block numbers for the last block
// of each authority set, once they have been finalized. These blocks are guaranteed to
// have a justification unless they were triggered by a forced change.
type AuthoritySetChanges []struct {
	setId       uint64
	blockNumber uint
}

// AuthoritySet A set of authorities.
type AuthoritySet struct {
	// The current active authorities.
	currentAuthorities AuthorityList
	// The current set id.
	setId uint64
	// Tree of pending standard changes across forks. Standard changes are
	// enacted on finality and must be enacted (i.e. finalized) in-order across
	// a given branch
	pendingStandardChanges ForkTree
	// Pending forced changes across different forks (at most one per fork).
	// Forced changes are enacted on block depth (not finality), for this
	// reason only one forced change should exist per fork. When trying to
	// apply forced changes we keep track of any pending standard changes that
	// they may depend on, this is done by making sure that any pending change
	// that is an ancestor of the forced changed and its effective block number
	// is lower than the last finalized block (as signaled in the forced
	// change) must be applied beforehand.
	pendingForcedChanges []PendingChange
	// Track at which blocks the set id changed. This is useful when we need to prove finality for
	// a given block since we can figure out what set the block belongs to and when the set
	// started/ended.
	authoritySetChanges AuthoritySetChanges
}

// SharedAuthoritySet A shared authority set.
// TODO thought: I wonder if I can just hold the data and a mutex for this case?
type SharedAuthoritySet struct {
	// TODO make shared
	authoritySet AuthoritySet
}

// Inner Returns access to the [`AuthoritySet`].
func (sas *SharedAuthoritySet) Inner() AuthoritySet {
	return sas.authoritySet
}

// InnerLocked
// Returns access to the [`AuthoritySet`] and locks it.
//
// For more information see [`SharedDataLocked`].
func (sas *SharedAuthoritySet) InnerLocked() AuthoritySet {
	return sas.authoritySet
}
