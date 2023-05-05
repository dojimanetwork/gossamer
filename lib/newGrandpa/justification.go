// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package newGrandpa

import (
	"github.com/ChainSafe/gossamer/dot/types"
	"github.com/ChainSafe/gossamer/lib/common"
)

// TODO maybe this is a good starting point?

// GrandpaJustificationSP A GRANDPA justification for block finality, it includes a commit message and
// an ancestry proof including all headers routing all precommit target blocks
// to the commit target block. Due to the current voting strategy the precommit
// targets should be the same as the commit target, since honest voters don't
// vote past authority set change blocks.
//
// This is meant to be stored in the db and passed around the network to other
// nodes, and are used by syncing nodes to prove authority set handoffs.
//
// from sp_consensus_grandpa::GrandpaJustification<Block::Header>
type GrandpaJustificationSP struct {
	round uint64
	//commit grandpa::CompactCommit
	vote_ancestries []types.Header
}

// GrandpaJustification A GRANDPA justification for block finality, it includes a commit message and
// an ancestry proof including all headers routing all precommit target blocks
// to the commit target block. Due to the current voting strategy the precommit
// targets should be the same as the commit target, since honest voters don't
// vote past authority set change blocks.
//
// This is meant to be stored in the db and passed around the network to other
// nodes, and are used by syncing nodes to prove authority set handoffs.
type GrandpaJustification struct {
	// TODO might be able to just use sp_consensus_grandpa::GrandpaJustification<Block::Header> type
	// TODO I think if we dont need _block then we can remove this

	// The GRANDPA justification for block finality.
	//pub justification: sp_consensus_grandpa::GrandpaJustification<Block::Header>,
	justification GrandpaJustificationSP

	// TODO Since this is phantom data, maybe this isnt needed?
	//_block: PhantomData<Block>,
}

// Create a GRANDPA justification from the given commit. This method
// assumes the commit is valid and well-formed.
// Params:
// - client: &Arc<C> where C is type blockchain backend
// - round u64
// - commit: Commit<Block::Header>
func (s *GrandpaJustification) from_commit() {
	// votes_ancestries_hashes
	// vote_ancestries

	// we pick the precommit for the lowest block as the base that
	// should serve as the root block for populating ancestry (i.e.
	// collect all headers from all precommit blocks to the base)

	// iterate throught precommits

	// Ok(sp_consensus_grandpa::GrandpaJustification { round, commit, votes_ancestries }.into())
}

// Decode a GRANDPA justification and validate the commit and the votes'
// ancestry proofs finalize the given block.
// params:
// - encoded: &[u8],
// - finalized_target: (Block::Hash, NumberFor<Block>),
// - set_id: u64,
// - voters: &VoterSet<AuthorityId>,
func (s *GrandpaJustification) decode_and_verify_finalizes(_ []byte, _ common.Hash, _ uint, _ uint64) {
	
}

// Validate the commit and the votes' ancestry proofs.
func (s *GrandpaJustification) verify() {}

// Validate the commit and the votes' ancestry proofs.
func (s *GrandpaJustification) verify_with_voter_set() {}

// The target block number and hash that this justifications proves finality for.
func (s *GrandpaJustification) target() {}

// AncestryChain A utility trait implementing `finality_grandpa::Chain` using a given set of headers.
// This is useful when validating commits, using the given set of headers to
// verify a valid ancestry route to the target commit block.
type AncestryChain struct {
	//ancestry: HashMap<Block::Hash, Block::Header>,
}

func (s *AncestryChain) ancestry() {}
