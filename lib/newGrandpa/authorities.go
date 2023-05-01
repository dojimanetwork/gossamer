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
func (s *SharedAuthoritySet) inner() {}

// Returns access to the [`AuthoritySet`] and locks it.
//
// For more information see [`SharedDataLocked`].
// returns a SharedDataLocked<AuthoritySet<H, N>>
func (s *SharedAuthoritySet) inner_locked() {}

// Impl 2 from substrate

// Get the earliest limit-block number that's higher or equal to the given
// min number, if any.
func (s *SharedAuthoritySet) current_limit() {}

// Get the current set ID. This is incremented every time the set changes.
func (s *SharedAuthoritySet) set_id() {}

// Get the current authorities and their weights (for the current set ID).
func (s *SharedAuthoritySet) current_authorities() {}

// Clone the inner `AuthoritySet`.
func (s *SharedAuthoritySet) clone_inner() {}

// Clone the inner `AuthoritySetChanges`.
func (s *SharedAuthoritySet) authority_set_changes() {}

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

// Impl 1

// authority sets must be non-empty and all weights must be greater than 0
func (s *AuthoritySet) invalid_authority_list() {}

// Get a genesis set with given authorities.
func (s *AuthoritySet) genesis() {}

// Create a new authority set.
func (s *AuthoritySet) new() {}

// Get the current set id and a reference to the current authority set.
func (s *AuthoritySet) current() {}

// Revert to a specified block given its `hash` and `number`.
// This removes all the authority set changes that were announced after
// the revert point.
// Revert point is identified by `number` and `hash`.
func (s *AuthoritySet) revert() {}

// Impl 2

// Returns the block hash and height at which the next pending change in
// the given chain (i.e. it includes `best_hash`) was signalled, `None` if
// there are no pending changes for the given chain.
//
// This is useful since we know that when a change is signalled the
// underlying runtime authority set management module (e.g. session module)
// has updated its internal state (e.g. a new session started).
func (s *AuthoritySet) next_change() {}

func (s *AuthoritySet) add_standard_change() {}

func (s *AuthoritySet) add_forced_change() {}

// Note an upcoming pending transition. Multiple pending standard changes
// on the same branch can be added as long as they don't overlap. Forced
// changes are restricted to one per fork. This method assumes that changes
// on the same branch will be added in-order. The given function
// `is_descendent_of` should return `true` if the second hash (target) is a
// descendent of the first hash (base).
func (s *AuthoritySet) add_pending_change() {}

// Inspect pending changes. Standard pending changes are iterated first,
// and the changes in the tree are traversed in pre-order, afterwards all
// forced changes are iterated.
func (s *AuthoritySet) pending_changes() {}

// Get the earliest limit-block number, if any. If there are pending changes across
// different forks, this method will return the earliest effective number (across the
// different branches) that is higher or equal to the given min number.
//
// Only standard changes are taken into account for the current
// limit, since any existing forced change should preclude the voter from voting.
func (s *AuthoritySet) current_limit() {}

// Apply or prune any pending transitions based on a best-block trigger.
//
// Returns `Ok((median, new_set))` when a forced change has occurred. The
// median represents the median last finalized block at the time the change
// was signaled, and it should be used as the canon block when starting the
// new grandpa voter. Only alters the internal state in this case.
//
// These transitions are always forced and do not lead to justifications
// which light clients can follow.
//
// Forced changes can only be applied after all pending standard changes
// that it depends on have been applied. If any pending standard change
// exists that is an ancestor of a given forced changed and which effective
// block number is lower than the last finalized block (as defined by the
// forced change), then the forced change cannot be applied. An error will
// be returned in that case which will prevent block import.
func (s *AuthoritySet) apply_forced_changes() {}

// Apply or prune any pending transitions based on a finality trigger. This
// method ensures that if there are multiple changes in the same branch,
// finalizing this block won't finalize past multiple transitions (i.e.
// transitions must be finalized in-order). The given function
// `is_descendent_of` should return `true` if the second hash (target) is a
// descendent of the first hash (base).
//
// When the set has changed, the return value will be `Ok(Some((H, N)))`
// which is the canonical block where the set last changed (i.e. the given
// hash and number).
func (s *AuthoritySet) apply_standard_changes() {}

// Check whether the given finalized block number enacts any standard
// authority set change (without triggering it), ensuring that if there are
// multiple changes in the same branch, finalizing this block won't
// finalize past multiple transitions (i.e. transitions must be finalized
// in-order). Returns `Some(true)` if the block being finalized enacts a
// change that can be immediately applied, `Some(false)` if the block being
// finalized enacts a change but it cannot be applied yet since there are
// other dependent changes, and `None` if no change is enacted. The given
// function `is_descendent_of` should return `true` if the second hash
// (target) is a descendent of the first hash (base).
func (s *AuthoritySet) enacts_standard_change() {}

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

func (s *PendingChange) decode() {}

// Returns the effective number this change will be applied at.
func (s *PendingChange) effective_number() {}

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

func (s *AuthoritySetChanges) empty() {}

func (s *AuthoritySetChanges) append() {}

func (s *AuthoritySetChanges) get_set_id() {}

func (s *AuthoritySetChanges) insert() {}

// Returns an iterator over all historical authority set changes starting at the given block
// number (excluded). The iterator yields a tuple representing the set id and the block number
// of the last block in that set.
func (s *AuthoritySetChanges) iter_from() {}
