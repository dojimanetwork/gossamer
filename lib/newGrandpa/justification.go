// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package newGrandpa

// TODO maybe this is a good starting point?

// GrandpaJustification A GRANDPA justification for block finality, it includes a commit message and
// an ancestry proof including all headers routing all precommit target blocks
// to the commit target block. Due to the current voting strategy the precommit
// targets should be the same as the commit target, since honest voters don't
// vote past authority set change blocks.
//
// This is meant to be stored in the db and passed around the network to other
// nodes, and are used by syncing nodes to prove authority set handoffs.
type GrandpaJustification struct {
	// The GRANDPA justification for block finality.
	//pub justification: sp_consensus_grandpa::GrandpaJustification<Block::Header>,
	//_block: PhantomData<Block>,
}

// Create a GRANDPA justification from the given commit. This method
// assumes the commit is valid and well-formed.
func (s *GrandpaJustification) from_commit() {}

// Decode a GRANDPA justification and validate the commit and the votes'
// ancestry proofs finalize the given block.
func (s *GrandpaJustification) decode_and_verify_finalizes() {}

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
