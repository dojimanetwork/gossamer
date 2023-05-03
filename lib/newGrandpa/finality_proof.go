// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package newGrandpa

// FinalityProofProvider Finality proof provider for serving network requests.
type FinalityProofProvider struct {
	//backend: Arc<BE>,
	//shared_authority_set: Option<SharedAuthoritySet<Block::Hash, NumberFor<Block>>>,
}

// Create new finality proof provider using:
//
// - backend for accessing blockchain data;
// - authority_provider for calling and proving runtime methods.
// - shared_authority_set for accessing authority set data
func (s *FinalityProofProvider) new() {}

// Create new finality proof provider for the service using:
//
// - backend for accessing blockchain data;
// - storage_provider, which is generally a client.
// - shared_authority_set for accessing authority set data
func (s *FinalityProofProvider) new_for_service() {}

// Prove finality for the given block number by returning a Justification for the last block of
// the authority set in bytes.
func (s *FinalityProofProvider) prove_finality() {}

// Prove finality for the given block number by returning a Justification for the last block of
// the authority set.
//
// If `collect_unknown_headers` is true, the finality proof will include all headers from the
// requested block until the block the justification refers to.
func (s *FinalityProofProvider) prove_finality_proof() {}

// FinalityProof Finality for block B is proved by providing:
// 1) the justification for the descendant block F;
// 2) headers sub-chain (B; F] if B != F;
type FinalityProof struct {
	// The hash of block F for which justification is provided.
	//pub block: Header::Hash,

	// Justification of the block F.
	//pub justification: Vec<u8>,

	// The set of headers in the range (B; F] that we believe are unknown to the caller. Ordered.
	//pub unknown_headers: Vec<Header>,
}

type FinalityProofError error

// Prove finality for the given block number by returning a justification for the last block of
// the authority set of which the given block is part of, or a justification for the latest
// finalized block if the given block is part of the current authority set.
//
// If `collect_unknown_headers` is true, the finality proof will include all headers from the
// requested block until the block the justification refers to.
func prove_finality() {

}
