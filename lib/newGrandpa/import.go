// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package newGrandpa

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
	//inner: Arc<Client>,
	//select_chain: SC,
	//authority_set: SharedAuthoritySet<Block::Hash, NumberFor<Block>>,
	//send_voter_commands: TracingUnboundedSender<VoterCommand<Block::Hash, NumberFor<Block>>>,
	//authority_set_hard_forks: HashMap<Block::Hash, PendingChange<Block::Hash, NumberFor<Block>>>,
	//justification_sender: GrandpaJustificationSender<Block>,
	//telemetry: Option<TelemetryHandle>,
	//_phantom: PhantomData<Backend>,red_authority_set: Option<SharedAuthoritySet<Block::Hash, NumberFor<Block>>>,
}

func (s *GrandpaBlockImport) on_start() {}

func (s *GrandpaBlockImport) import_justification() {}

// TODO this is actually an enum
type AppliedChanges struct{}

func (s *AppliedChanges) needs_justification() {}

type PendingSetChanges struct {
	//	just_in_case: Option<(
	//	AuthoritySet<Block::Hash, NumberFor<Block>>,
	//	SharedDataLockedUpgradable<AuthoritySet<Block::Hash, NumberFor<Block>>>,
	//)>,
	//	applied_changes: AppliedChanges<Block::Hash, NumberFor<Block>>,
	//	do_pause: bool,
}

func (s *PendingSetChanges) revert() {}

func (s *PendingSetChanges) defuse() {}

func (s *PendingSetChanges) drop() {}

// Checks the given header for a consensus digest signalling a **standard** scheduled change and
// extracts it.
func find_scheduled_change() {}

// Checks the given header for a consensus digest signalling a **forced** scheduled change and
// extracts it.
func find_forced_change() {}

func (s *GrandpaBlockImport) check_new_change() {}

func (s *GrandpaBlockImport) make_authorities_changes() {}

func (s *GrandpaBlockImport) current_set_id() {}

func (s *GrandpaBlockImport) import_state() {}

func (s *GrandpaBlockImport) import_block() {}

func (s *GrandpaBlockImport) check_block() {}

func (s *GrandpaBlockImport) new() {}

// Import a block justification and finalize the block.
//
// If `enacts_change` is set to true, then finalizing this block *must*
// enact an authority set change, the function will panic otherwise.
func (s *GrandpaBlockImport) import_justification2() {}
