// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package grandpa

import (
	"github.com/ChainSafe/gossamer/dot/types"
	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/ChainSafe/gossamer/lib/crypto/ed25519"
	"github.com/ChainSafe/gossamer/lib/keystore"
	finalityGrandpa "github.com/ChainSafe/gossamer/pkg/finality-grandpa"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

// AuthorityId Identity of a Grandpa authority.
type AuthorityId = ed25519.PublicKeyBytes

// SharedVoterState Shared voter state for querying.
type SharedVoterState struct {
	// TODO add RwLock
	inner finalityGrandpa.VoterState[AuthorityId]
}

// ClientForGrandpa A trait that includes all the client functionalities grandpa requires.
// TODO investigate how to deal with client/backend
type ClientForGrandpa interface {
	HeaderBackend
}

// HeaderBackend Blockchain database header backend. Does not perform any validation.
// This is not located in grandpa, but is in primatives/blockchain/src/backend.rs
// TODO investigate how to deal with backend
type HeaderBackend interface {
	header()
	info() Info
	status()
	number()
	hash()
	block_hash_from_id()
	block_number_from_id()
	expect_header()
	expect_block_number_from_id()
	expect_block_hash_from_id()
}

// Gap is used to represent missing blocks after warp sync
type Gap struct {
	start uint
	end   uint
}

// Info Blockchain info, this is also in backend
type Info struct {
	// Best block hash.
	bestHash common.Hash
	// Best block number.
	bestNumber uint
	// Genesis block hash.
	genesisHash common.Hash
	// The head of the finalized chain.
	finalizedHash common.Hash
	// Last finalized block number.
	finalizedNumber uint
	// Last finalized state.
	// Below is actually an option containing a tuple TODO think about
	finalizedStateHash   common.Hash
	finalizedStateNumber uint
	// Number of concurrent leave forks.
	numberLeaves uintptr // TODO is this go equiv of usize?
	// Missing blocks after warp sync. (start, end).
	blockGap *Gap
}

// The SelectChain trait defines the strategy upon which the head is chosen
// if multiple forks are present for an opaque definition of "best" in the
// specific chain build.
//
// The Strategy can be customized for the two use cases of authoring new blocks
// upon the best chain or which fork to finalize. Unless implemented differently
// by default finalization methods fall back to use authoring, so as a minimum
// `_authoring`-functions must be implemented.
//
// Any particular user must make explicit, however, whether they intend to finalize
// or author through the using the right function call, as these might differ in
// some implementations.
//
// Non-deterministically finalizing chains may only use the `_authoring` functions.
type SelectChain interface{}

// TODO just here for now, believe need to impl
type S struct{}

// TODO investigate best way to do this in go
type TracingUnboundedReceiver struct{}

// LinkHalf Link between the block importer and the background voter.
type LinkHalf struct {
	client      ClientForGrandpa
	selectChain SelectChain
	// TODO part of aux schema, investigate if needed
	//persistentData      PersistentData

	// TODO all 3 of below are related to communication, investigate
	//voterCommandsRx     TracingUnboundedReceiver
	//justificationSender GrandpaJustificationSender
	//justificationStream GrandpaJustificationStream
	telemetry Telemetry
}

type Config struct {
	// The expected duration for a message to be gossiped across the network.
	gossipDuration time.Duration
	// Justification generation period (in blocks). GRANDPA will try to generate justifications
	// at least every justification_period blocks. There are some other events which might cause
	// justification generation.
	justificationPeriod uint32
	// Whether the GRANDPA observer protocol is live on the network and thereby
	// a full-node not running as a validator is running the GRANDPA observer
	// protocol (we will only issue catch-up requests to authorities when the
	// observer protocol is enabled).
	observerEnabled bool
	// The role of the local node (i.e. authority, full-node or light).
	localRole common.Roles
	// Some local identifier of the voter.
	name string
	// The keystore that manages the keys of this node.
	keystore keystore.Keystore
	// TelemetryHandle instance.
	telemetry Telemetry
	// Chain specific GRANDPA protocol name. See [`crate::protocol_standard_name`].
	protocolName string // this is an enum, but i think ok as string
}

type GrandpaParams struct {
	// Configuration for the GRANDPA service.
	config Config
	// A link to the block import worker.
	link LinkHalf
	// The Network instance.
	//
	// It is assumed that this network will feed us Grandpa notifications. When using the
	// `sc_network` crate, it is assumed that the Grandpa notifications protocol has been passed
	// to the configuration of the networking. See [`grandpa_peers_set_config`].
	network Network // TODO make sure this does what is needed
	// Event stream for syncing-related events.
	sync S
	// A voting rule used to potentially restrict target votes.
	votingRule VotingRule
	// The prometheus metrics registry.
	prometheumRegistry prometheus.Registry
	// The voter state is exposed at an RPC endpoint.
	sharedVoterState SharedVoterState
	// TelemetryHandle instance.
	telemetry Telemetry
}

// VoterWork Future that powers the voter.
type VoterWork struct{}

func (*VoterWork) New(grandpaParams GrandpaParams) *VoterWork {
	return &VoterWork{}
}

// GenesisAuthoritySetProvider Provider for the Grandpa authority set configured on the genesis block.
type GenesisAuthoritySetProvider interface {
	// Get the authority set at the genesis block.
	get()
}

// RunGrandpaVoter Run a GRANDPA voter as a task. Provide configuration and a link to a
// block import worker that has already been instantiated with `block_import`.
func RunGrandpaVoter(grandpaParams GrandpaParams) {
	// TODO check is this is still valid
	// NOTE: we have recently removed `run_grandpa_observer` from the public
	// API, I felt it is easier to just ignore this field rather than removing
	// it from the config temporarily. This should be removed after #5013 is
	// fixed and we re-add the observer to the public API.
	grandpaParams.config.observerEnabled = false

	// TODO investigate NetworkBridge
	//network := NetworkBridge{}

	// Telemetry stuff

	var voterWork *VoterWork
	voterWork = voterWork.New(grandpaParams)
}

// A descriptor for an authority set hard fork. These are authority set changes
// that are not signalled by the runtime and instead are defined off-chain
// (hence the hard fork).
type AuthoritySetHardFork struct {
	// The new authority set id.
	setId uint64
	// The block hash and number at which the hard fork should be applied.
	blockHash   common.Hash
	blockNumber uint //verify this
	// The authorities in the new set.
	authorities []types.Authority // I believe this should work
	// The latest block number that was finalized before this authority set
	// hard fork. When defined, the authority set change will be forced, i.e.
	// the node won't wait for the block above to be finalized before enacting
	// the change, and the given finalized number will be used as a base for
	// voting.
	// option for block number
	lastFinalized *uint
}

// BlockImport Make block importer and link half necessary to tie the background voter
// to it.
func BlockImport(client ClientForGrandpa,
	genesisAuthoritiesProvider GenesisAuthoritySetProvider,
	selectChain SelectChain, telemetry Telemetry) (GrandpaBlockImport, LinkHalf, error) {

	return BlockImportWIthAuthoritySetHardForks(client, genesisAuthoritiesProvider, selectChain, []AuthoritySetHardFork{}, telemetry)
}

// BlockImportWIthAuthoritySetHardForks Make block importer and link half necessary to tie the background voter to
// it. A vector of authority set hard forks can be passed, any authority set
// change signaled at the given block (either already signalled or in a further
// block when importing it) will be replaced by a standard change with the
// given static authorities.
func BlockImportWIthAuthoritySetHardForks(client ClientForGrandpa,
	genesisAuthoritiesProvider GenesisAuthoritySetProvider,
	selectChain SelectChain, authorities []AuthoritySetHardFork, telemetry Telemetry) (GrandpaBlockImport, LinkHalf, error) {

	chainInfo := client.info()
	_ = chainInfo.genesisHash

	//get Persistant data using aux schema

	// set up tracing unbounded sender and reciever

	// set up notification stream for justification sender

	// create pending change objects with 0 delay for each authority set hard fork.

	grandpaBlockImport := GrandpaBlockImport{
		inner:       client,
		selectChain: selectChain,
		// persistent_data.authority_set
		//authoritySet: SharedAuthoritySet{},
		telemetry: telemetry,
	}

	linkHalf := LinkHalf{
		client:      client,
		selectChain: selectChain,
		telemetry:   telemetry,
	}

	return grandpaBlockImport, linkHalf, nil
}
