// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package newGrandpa

type HistoricalVotes struct {
	// set equal to finality_grandpa::HistoricalVotes
}

// CompletedRound Data about a completed round. The set of votes that is stored must be
// minimal, i.e. at most one equivocation is stored per voter.
type CompletedRound struct {
	// The round number.
	//pub number: RoundNumber,

	// The round state (prevote ghost, estimate, finalized, etc.)
	//pub state: RoundState<Block::Hash, NumberFor<Block>>,

	// The target block base used for voting in the round.
	//pub base: (Block::Hash, NumberFor<Block>),

	// All the votes observed in the round.
	//pub votes: Vec<SignedMessage<Block::Header>>,
}

// CompletedRounds Data about last completed rounds within a single voter set. Stores
// NUM_LAST_COMPLETED_ROUNDS and always contains data about at least one round
// (genesis).
type CompletedRounds struct {
	//rounds: Vec<CompletedRound<Block>>,
	//set_id: SetId,
	//voters: Vec<AuthorityId>,
}

//	NUM_LAST_COMPLETED_ROUNDS
//	NOTE: the current strategy for persisting completed rounds is very naive
//
// (update everything) and we also rely on cloning to do atomic updates,
// therefore this value should be kept small for now.
// TODO prob delete this comment
const NUM_LAST_COMPLETED_ROUNDS = 2

func (s *CompletedRounds) encode() {}

func (s *CompletedRounds) decode() {}

// Create a new completed rounds tracker with NUM_LAST_COMPLETED_ROUNDS capacity.
func (s *CompletedRounds) new() {}

// Get the set-id and voter set of the completed rounds.
func (s *CompletedRounds) set_info() {}

// Iterate over all completed rounds.
func (s *CompletedRounds) iter() {}

// Returns the last (latest) completed round.
func (s *CompletedRounds) last() {}

// Push a new completed round, oldest round is evicted if number of rounds
// is higher than `NUM_LAST_COMPLETED_ROUNDS`.
func (s *CompletedRounds) push() {}

// A map with voter status information for currently live rounds,
// which votes have we cast and what are they.
//TODO figure out how to impl this
//pub type CurrentRounds<Block> = BTreeMap<RoundNumber, HasVoted<<Block as BlockT>::Header>>;

// VoterSetState The state of the current voter set, whether it is currently active or not
// and information related to the previously completed rounds. Current round
// voting status is used when restarting the voter, i.e. it will re-use the
// previous votes for a given round if appropriate (same round and same local
// key).
// TODO figure out if need this as enum
type VoterSetState struct{}

//pub enum VoterSetState<Block: BlockT> {
//	// The voter is live, i.e. participating in rounds.
//	Live {
//	// The previously completed rounds.
//	completed_rounds: CompletedRounds<Block>,
//	// Voter status for the currently live rounds.
//	current_rounds: CurrentRounds<Block>,
//	},
//	// The voter is paused, i.e. not casting or importing any votes.
//	Paused {
//	// The previously completed rounds.
//	completed_rounds: CompletedRounds<Block>,
//	},
//}

// Create a new live VoterSetState with round 0 as a completed round using
// the given genesis state and the given authorities. Round 1 is added as a
// current round (with state `HasVoted::No`).
func (s *VoterSetState) live() {}

// Returns the last completed rounds.
func (s *VoterSetState) completed_rounds() {}

// Returns the last completed round.
func (s *VoterSetState) last_completed_round() {}

// Returns the voter set state validating that it includes the given round
// in current rounds and that the voter isn't paused.
func (s *VoterSetState) with_current_round() {}

// Whether we've voted already during a prior run of the program.
// TODO figure out
type HasVoted struct{}

//pub enum HasVoted<Header: HeaderT> {
//	// Has not voted already in this round.
//	No,
//	// Has voted in this round.
//	Yes(AuthorityId, Vote<Header>),
//}

// The votes cast by this voter already during a prior run of the program.
// TODO figure out
//pub enum Vote<Header: HeaderT> {
//	// Has cast a proposal.
//	Propose(PrimaryPropose<Header>),
//	// Has cast a prevote.
//	Prevote(Option<PrimaryPropose<Header>>, Prevote<Header>),
//	// Has cast a precommit (implies prevote.)
//	Precommit(Option<PrimaryPropose<Header>>, Prevote<Header>, Precommit<Header>),
//}

// Returns the proposal we should vote with (if any.)
func (s *HasVoted) propose() {}

// Returns the prevote we should vote with (if any.)
func (s *HasVoted) prevote() {}

// Returns the precommit we should vote with (if any.)
func (s *HasVoted) precommit() {}

// Returns true if the voter can still propose, false otherwise.
func (s *HasVoted) can_propose() {}

// Returns true if the voter can still prevote, false otherwise.
func (s *HasVoted) can_prevote() {}

// Returns true if the voter can still precommit, false otherwise.
func (s *HasVoted) can_precommit() {}

// SharedVoterSetState A voter set state meant to be shared safely across multiple owners.
type SharedVoterSetState struct {
	//// The inner shared `VoterSetState`.
	//inner: Arc<RwLock<VoterSetState<Block>>>,
	//
	//// A tracker for the rounds that we are actively participating on (i.e. voting)
	//// and the authority id under which we are doing it.
	//voting: Arc<RwLock<HashMap<RoundNumber, AuthorityId>>>,
}

// Create a new shared voter set tracker with the given state.
func (s *HasVoted) new() {}

// Read the inner voter set state.
func (s *HasVoted) read() {}

// Get the authority id that we are using to vote on the given round, if any.
func (s *HasVoted) voting_on() {}

// Note that we started voting on the give round with the given authority id.
func (s *HasVoted) started_voting_on() {}

// Note that we have finished voting on the given round. If we were voting on
// the given round, the authority id that we were using to do it will be
// cleared.
func (s *HasVoted) finished_voting_on() {}

// Return vote status information for the current round.
func (s *HasVoted) has_voted() {}

// NOTE: not exposed outside of this module intentionally.
// TODO might not need this
func (s *HasVoted) with() {}

// Metrics Prometheus metrics for GRANDPA.
type Metrics struct {
	//finality_grandpa_round: Gauge<U64>,
	//finality_grandpa_prevotes: Counter<U64>,
	//finality_grandpa_precommits: Counter<U64>,
}

func (s *Metrics) register() {}

// Environment The environment we run GRANDPA in.
type Environment struct{}

// Updates the voter set state using the given closure. The write lock is
// held during evaluation of the closure and the environment's voter set
// state is set to its result if successful.
func (s *Environment) update_voter_set_state() {}

// Report the given equivocation to the GRANDPA runtime module. This method
// generates a session membership proof of the offender and then submits an
// extrinsic to report the equivocation. In particular, the session membership
// proof must be generated at the block at which the given set was active which
// isn't necessarily the best block if there are pending authority set changes.
func (s *Environment) report_equivocation() {}

// TODO finality_grandpa::Chain impl for environment
type finality_grandpaChain struct{}

// impl<BE, Block, C, N, S, SC, VR> finality_grandpa::Chain<Block::Hash, NumberFor<Block>>
// for Environment<BE, Block, C, N, S, SC, VR
func (finality_grandpaChain) ancestry() {}

// TODO voter::Environment impl for environment, also from finality grandpa
type voterEnvironment struct{}

func (voterEnvironment) best_chain_containing() {}

func (voterEnvironment) round_data() {}

func (voterEnvironment) proposed() {}

func (voterEnvironment) prevoted() {}

func (voterEnvironment) precommitted() {}

func (voterEnvironment) completed() {}

func (voterEnvironment) concluded() {}

func (voterEnvironment) finalize_block() {}

func (voterEnvironment) round_commit_timer() {}

func (voterEnvironment) prevote_equivocation() {}

func (voterEnvironment) precommit_equivocation() {}

// TODO figure out
//pub(crate) enum JustificationOrCommit<Block: BlockT> {
//	Justification(GrandpaJustification<Block>),
//	Commit((RoundNumber, Commit<Block::Header>)),
//}

func best_chain_containing() {}

// Finalize the given block and apply any authority set changes. If an
// authority set change is enacted then a justification is created (if not
// given) and stored with the block when finalizing it.
// This method assumes that the block being finalized has already been imported.
func finalize_block() {}
