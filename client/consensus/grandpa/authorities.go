// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package grandpa

import (
	"errors"
	"github.com/ChainSafe/gossamer/dot/types"
	"github.com/ChainSafe/gossamer/lib/common"
)

// AuthorityList A list of Grandpa authorities with associated weights.
type AuthorityList []types.Authority

// DelayKind Kinds of delays for pending changes.
// This is an enum
type DelayKind struct {
	// TODO impl
}

// PendingChange A pending change to the authority set.
//
// This will be applied when the announcing block is at some depth within
// the finalized or unfinalized chain.
type PendingChange struct {
	nextAuthorities AuthorityList
	delay           uint
	canonHeight     uint
	canonHash       common.Hash
	delayKind       DelayKind
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

// InvalidAuthorityList authority sets must be non-empty and all weights must be greater than 0
func (authSet *AuthoritySet) InvalidAuthorityList(authorities AuthorityList) bool {
	if len(authorities) == 0 {
		return false
	}

	for _, authority := range authorities {
		if authority.Weight == 0 {
			return false
		}
	}
	return true
}

func (authSet *AuthoritySet) AddPendingChange(pending PendingChange, isDescendentOf bool) error {
	if authSet.InvalidAuthorityList(pending.nextAuthorities) {
		return errors.New("invalid authority set, either empty or with an authority weight set to 0")
	}

	return nil
}

// Get the earliest limit-block number, if any. If there are pending changes across
// different forks, this method will return the earliest effective number (across the
// different branches) that is higher or equal to the given min number.
//
// Only standard changes are taken into account for the current
// limit, since any existing forced change should preclude the voter from voting.
func (authSet *AuthoritySet) CurrentLimit(min uint) {}

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
