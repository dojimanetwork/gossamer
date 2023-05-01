// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package newGrandpa

import "errors"

//! Utilities for dealing with authorities, authority sets, and handoffs.

// ErrAuthoritySet Error type returned on operations on the `AuthoritySet`.
// TODO fill in with VDT and/or proper errors
var ErrAuthoritySet = errors.New("authority set error fill in error")

// TODO dont think we need this function
//fn from(err: fork_tree::Error<E>) -> Error<N, E> {}

// SharedAuthoritySet A shared authority set.
type SharedAuthoritySet struct {
	//  SharedData<AuthoritySet<H, N>>,
}

// TODO this maybe ill need to implement SharedDataLocked or something related
// Impl 1 from substrate

// Returns access to the [`AuthoritySet`].
// returns a MappedMutexGuard<AuthoritySet<H, N>>
func (s *SharedAuthoritySet) inner() {
	//TODO implement
}

// Returns access to the [`AuthoritySet`] and locks it.
//
// For more information see [`SharedDataLocked`].
// returns a SharedDataLocked<AuthoritySet<H, N>>
func (s *SharedAuthoritySet) inner_locked() {
	//TODO implement
}

// Impl 2 from substrate

// Get the earliest limit-block number that's higher or equal to the given
// min number, if any.
func (s *SharedAuthoritySet) current_limit() {
	//TODO implement
}

// Get the current set ID. This is incremented every time the set changes.
func (s *SharedAuthoritySet) set_id() {
	//TODO implement
}

// Get the current authorities and their weights (for the current set ID).
func (s *SharedAuthoritySet) current_authorities() {
	//TODO implement
}

// Clone the inner `AuthoritySet`.
func (s *SharedAuthoritySet) clone_inner() {
	//TODO implement
}

// Clone the inner `AuthoritySetChanges`.
func (s *SharedAuthoritySet) authority_set_changes() {
	//TODO implement
}

// Status of the set after changes were applied.
type Status struct {
	// Whether internal changes were made.
	Changed bool

	// `Some` when underlying authority set has changed, containing the
	// block where that set changed.
	//pub(crate) new_set_block: Option<(H, N)>
}

// AuthoritySet A set of authorities.
type AuthoritySet struct {
	/// The current active authorities.
	//pub(crate) current_authorities: AuthorityList,

	/// The current set id.
	//pub(crate) set_id: u64,

	/// Tree of pending standard changes across forks. Standard changes are
	/// enacted on finality and must be enacted (i.e. finalized) in-order across
	/// a given branch
	//pub(crate) pending_standard_changes: ForkTree<H, N, PendingChange<H, N>>,

	/// Pending forced changes across different forks (at most one per fork).
	/// Forced changes are enacted on block depth (not finality), for this
	/// reason only one forced change should exist per fork. When trying to
	/// apply forced changes we keep track of any pending standard changes that
	/// they may depend on, this is done by making sure that any pending change
	/// that is an ancestor of the forced changed and its effective block number
	/// is lower than the last finalized block (as signaled in the forced
	/// change) must be applied beforehand.
	//pending_forced_changes: Vec<PendingChange<H, N>>,

	/// Track at which blocks the set id changed. This is useful when we need to prove finality for
	/// a given block since we can figure out what set the block belongs to and when the set
	/// started/ended.
	//pub(crate) authority_set_changes: AuthoritySetChanges<N>,
}

// Kinds of delays for pending changes.
// TODO this is implemented as an enum, investigate if we need this or not
//pub enum DelayKind<N> {
//	/// Depth in finalized chain.
//	Finalized,
//	/// Depth in best chain. The median last finalized block is calculated at the time the
//	/// change was signaled.
//	Best { median_last_finalized: N },
//}

/// A pending change to the authority set.
///
/// This will be applied when the announcing block is at some depth within
/// the finalized or unfinalized chain.

type PendingChange struct {
	/// The new authorities and weights to apply.
	//pub(crate) next_authorities: AuthorityList,

	/// How deep in the chain the announcing block must be
	/// before the change is applied.
	//pub(crate) delay: N,

	/// The announcing block's height.
	//pub(crate) canon_height: N,

	/// The announcing block's hash.
	//pub(crate) canon_hash: H,

	/// The delay kind.
	//pub(crate) delay_kind: DelayKind<N>,
}

// AuthoritySetChanges Tracks historical authority set changes. We store the block numbers for the last block
// of each authority set, once they have been finalized. These blocks are guaranteed to
// have a justification unless they were triggered by a forced change.
// pub struct AuthoritySetChanges<N>(Vec<(u64, N)>);
type AuthoritySetChanges struct{}

/// The response when querying for a set id for a specific block. Either we get a set id
/// together with a block number for the last block in the set, or that the requested block is in
/// the latest set, or that we don't know what set id the given block belongs to.
// TODO this is implemented as an enum, investigate if we need this or not
//pub enum AuthoritySetChangeId<N> {
// The requested block is in the latest set.
//Latest,

// Tuple containing the set id and the last block number of that set.
//Set(SetId, N),

// We don't know which set id the request block belongs to (this can only happen due to
/// missing data).
//Unknown,
//}
