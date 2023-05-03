// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package newGrandpa

// TODO prob later thing to impl

// WarpSyncFragment A proof of an authority set change.
type WarpSyncFragment struct {
	// The last block that the given authority set finalized. This block should contain a digest
	// signaling an authority set change from which we can fetch the next authority set.
	//pub header: Block::Header,

	// A justification for the header above which proves its finality. In order to validate it the
	// verifier must be aware of the authorities and set id for which the justification refers to.
	//pub justification: GrandpaJustification<Block>,
}

// WarpSyncProof An accumulated proof of multiple authority set changes.
type WarpSyncProof struct {
	//proofs: Vec<WarpSyncFragment<Block>>,
	//is_finished: bool,
}

// Generates a warp sync proof starting at the given block. It will generate authority set
// change proofs for all changes that happened from `begin` until the current authority set
// (capped by MAX_WARP_SYNC_PROOF_SIZE).
func (s *WarpSyncProof) generate() {}

// Verifies the warp sync proof starting at the given set id and with the given authorities.
// Verification stops when either the proof is exhausted or finality for the target header can
// be proven. If the proof is valid the new set id and authorities is returned.
func (s *WarpSyncProof) verify() {}

// NetworkProvider Implements network API for warp sync.
type NetworkProvider struct {
	//backend: Arc<Backend>,
	//authority_set: SharedAuthoritySet<Block::Hash, NumberFor<Block>>,
	//hard_forks: HashMap<(Block::Hash, NumberFor<Block>), (SetId, AuthorityList)>,
}

// Create a new instance for a given backend and authority set.
func (s *NetworkProvider) new() {}

// This is from network i believe
type WarpSyncProvider struct{}

func (s *WarpSyncProvider) generate() {}

func (s *WarpSyncProvider) verify() {}
