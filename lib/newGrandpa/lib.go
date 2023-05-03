// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package newGrandpa

// Integration of the GRANDPA finality gadget into substrate.
//
// This crate is unstable and the API and usage may change.
//
// This crate provides a long-running future that produces finality notifications.
//
// # Usage
//
// First, create a block-import wrapper with the `block_import` function. The
// GRANDPA worker needs to be linked together with this block import object, so
// a `LinkHalf` is returned as well. All blocks imported (from network or
// consensus or otherwise) must pass through this wrapper, otherwise consensus
// is likely to break in unexpected ways.
//
// Next, use the `LinkHalf` and a local configuration to `run_grandpa_voter`.
// This requires a `Network` implementation. The returned future should be
// driven to completion and will finalize blocks in the background.
//
// # Changing authority sets
//
// The rough idea behind changing authority sets in GRANDPA is that at some point,
// we obtain agreement for some maximum block height that the current set can
// finalize, and once a block with that height is finalized the next set will
// pick up finalization from there.
//
// Technically speaking, this would be implemented as a voting rule which says,
// "if there is a signal for a change in N blocks in block B, only vote on
// chains with length NUM(B) + N if they contain B". This conditional-inclusion
// logic is complex to compute because it requires looking arbitrarily far
// back in the chain.
//
// Instead, we keep track of a list of all signals we've seen so far (across
// all forks), sorted ascending by the block number they would be applied at.
// We never vote on chains with number higher than the earliest handoff block
// number (this is num(signal) + N). When finalizing a block, we either apply
// or prune any signaled changes based on whether the signaling block is
// included in the newly-finalized chain.

// CommunicationIn A global communication input stream for commits and catch up messages. Not
// exposed publicly, used internally to simplify types in the communication
// layer.
type CommunicationIn struct{}

// CommunicationInH Global communication input stream for commits and catch up messages, with
// the hash type not being derived from the block, useful for forcing the hash
// to some type (e.g. `H256`) when the compiler can't do the inference.
type CommunicationInH struct {
	// Do i need this?
}

// Global communication sink for commits with the hash type not being derived
// from the block, useful for forcing the hash to some type (e.g. `H256`) when
// the compiler can't do the inference.
type CommunicationOutH struct {
	// Do i need this?
}

// SharedVoterState Shared voter state for querying.
type SharedVoterState struct {
	// inner: Arc<RwLock<Option<Box<dyn voter::VoterState<AuthorityId> + Sync + Send>>>>,
}

// Create a new empty `SharedVoterState` instance.
func (s *SharedVoterState) empty() {}

func (s *SharedVoterState) reset() {}

// Get the inner `VoterState` instance.
func (s *SharedVoterState) voter_state() {}

// Config Configuration for the GRANDPA service
type Config struct {
	// The expected duration for a message to be gossiped across the network.
	//pub gossip_duration: Duration,

	// Justification generation period (in blocks). GRANDPA will try to generate justifications
	// at least every justification_period blocks. There are some other events which might cause
	// justification generation.
	//pub justification_period: u32,

	// Whether the GRANDPA observer protocol is live on the network and thereby
	// a full-node not running as a validator is running the GRANDPA observer
	// protocol (we will only issue catch-up requests to authorities when the
	// observer protocol is enabled).
	//pub observer_enabled: bool,

	// The role of the local node (i.e. authority, full-node or light).
	//pub local_role: sc_network::config::Role,

	// Some local identifier of the voter.
	//pub name: Option<String>,

	// The keystore that manages the keys of this node.
	//pub keystore: Option<KeystorePtr>,

	// TelemetryHandle instance.
	//pub telemetry: Option<TelemetryHandle>,

	// Chain specific GRANDPA protocol name. See [`crate::protocol_standard_name`].
	//pub protocol_name: ProtocolName,
}

func (s *Config) name() {}

type BlockStatus struct {
	// trait
}

func (s *BlockStatus) block_number() {}

// ClientForGrandpa A trait that includes all the client functionalities grandpa requires.
// Ideally this would be a trait alias, we're not there yet.
// tracking issue <https://github.com/rust-lang/rust/issues/41517>
type ClientForGrandpa struct {
	// trait
}

// BlockSyncRequester Something that one can ask to do a block sync request.
type BlockSyncRequester struct {
	// trait
}

// Notifies the sync service to try and sync the given block from the given
// peers.
//
// If the given vector of peers is empty then the underlying implementation
// should make a best effort to fetch the block from any peers it is
// connected to (NOTE: this assumption will change in the future #3629).
func (s *BlockSyncRequester) set_sync_fork_request() {
	// uses the NetworkBridge
}

// NewAuthoritySet A new authority set along with the canonical block it changed at.
type NewAuthoritySet struct {
	//pub(crate) canon_number: N,
	//pub(crate) canon_hash: H,
	//pub(crate) set_id: SetId,
	//pub(crate) authorities: AuthorityList,
}

// VoterCommand Commands issued to the voter.
type VoterCommand struct {
	//enum
}

// CommandOrError Signals either an early exit of a voter or an error.
type CommandOrError struct {
	//enum
}

// LinkHalf Link between the block importer and the background voter.
type LinkHalf struct {
	//client: Arc<C>,
	//select_chain: SC,
	//persistent_data: PersistentData<Block>,
	//voter_commands_rx: TracingUnboundedReceiver<VoterCommand<Block::Hash, NumberFor<Block>>>,
	//justification_sender: GrandpaJustificationSender<Block>,
	//justification_stream: GrandpaJustificationStream<Block>,
	//telemetry: Option<TelemetryHandle>,
}

// Get the shared authority set.
func (s *LinkHalf) shared_authority_set() {}

// Get the receiving end of justification notifications.
func (s *LinkHalf) justification_stream() {}

// GenesisAuthoritySetProvider Provider for the Grandpa authority set configured on the genesis block.
type GenesisAuthoritySetProvider struct {
	//trait
}

// Get the authority set at the genesis block.
func (s *LinkHalf) get() {}

// Make block importer and link half necessary to tie the background voter
// to it.
func block_import() {}

// A descriptor for an authority set hard fork. These are authority set changes
// that are not signalled by the runtime and instead are defined off-chain
// (hence the hard fork).
type AuthoritySetHardFork struct {
	// The new authority set id.
	//pub set_id: SetId,

	// The block hash and number at which the hard fork should be applied.
	//pub block: (Block::Hash, NumberFor<Block>),

	// The authorities in the new set.
	//pub authorities: AuthorityList,

	// The latest block number that was finalized before this authority set
	// hard fork. When defined, the authority set change will be forced, i.e.
	// the node won't wait for the block above to be finalized before enacting
	// the change, and the given finalized number will be used as a base for
	// voting.
	//pub last_finalized: Option<NumberFor<Block>>,
}

// Make block importer and link half necessary to tie the background voter to
// it. A vector of authority set hard forks can be passed, any authority set
// change signaled at the given block (either already signalled or in a further
// block when importing it) will be replaced by a standard change with the
// given static authorities.
func block_import_with_authority_set_hard_forks() {}

func global_communication() {}

type GrandpaParams struct {
	/// Configuration for the GRANDPA service.
	//pub config: Config,

	/// A link to the block import worker.
	//pub link: LinkHalf<Block, C, SC>,

	/// The Network instance.
	///
	/// It is assumed that this network will feed us Grandpa notifications. When using the
	/// `sc_network` crate, it is assumed that the Grandpa notifications protocol has been passed
	/// to the configuration of the networking. See [`grandpa_peers_set_config`].
	//pub network: N,

	/// Event stream for syncing-related events.
	//pub sync: S,

	/// A voting rule used to potentially restrict target votes.
	//pub voting_rule: VR,

	/// The prometheus metrics registry.
	//pub prometheus_registry: Option<prometheus_endpoint::Registry>,

	/// The voter state is exposed at an RPC endpoint.
	//pub shared_voter_state: SharedVoterState,

	/// TelemetryHandle instance.
	//pub telemetry: Option<TelemetryHandle>,
}

// Returns the configuration value to put in
// [`sc_network::config::NetworkConfiguration::extra_sets`].
// For standard protocol name see [`crate::protocol_standard_name`].
func grandpa_peers_set_config() {}

// Run a GRANDPA voter as a task. Provide configuration and a link to a
// block import worker that has already been instantiated with `block_import`.
func run_grandpa_voter() {}

type Metrics3 struct{}

// VoterWork Future that powers the voter.
type VoterWork struct{}

func (s *VoterWork) new() {}

// Rebuilds the `self.voter` field using the current authority set
// state. This method should be called when we know that the authority set
// has changed (e.g. as signalled by a voter command).
func (s *VoterWork) rebuild_voter() {}

func (s *VoterWork) handle_voter_command() {}

func (s *VoterWork) poll() {}

// Checks if this node has any available keys in the keystore for any authority id in the given
// voter set.  Returns the authority id for which keys are available, or `None` if no keys are
// available.
func local_authority_id() {}

// Reverts protocol aux data to at most the last finalized block.
// In particular, standard and forced authority set changes announced after the
// revert point are removed.
func revert() {
	// do i need this
}
