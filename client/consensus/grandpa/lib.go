// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package grandpa

import (
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
type ClientForGrandpa interface{}

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

// BlockImport Make block importer and link half necessary to tie the background voter
// to it.
func BlockImport(client ClientForGrandpa,
	genesisAuthoritiesProvider GenesisAuthoritySetProvider,
	selectCHain SelectChain, _ Telemetry) (GrandpaBlockImport, LinkHalf, error) {

	return GrandpaBlockImport{}, LinkHalf{}, nil
}
