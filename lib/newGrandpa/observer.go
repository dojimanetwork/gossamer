// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package newGrandpa

type ObserverChain struct {
	//client: &'a Arc<Client>,
	//_phantom: PhantomData<Block>,PhantomData<Backend>,red_authority_set: Option<SharedAuthoritySet<Block::Hash, NumberFor<Block>>>,
}

func (s *ObserverChain) ancestry() {}

func grandpa_observer() {}

// Run a GRANDPA observer as a task, the observer will finalize blocks only by
// listening for and validating GRANDPA commits instead of following the full
// protocol. Provide configuration and a link to a block import worker that has
// already been instantiated with `block_import`.
// NOTE: this is currently not part of the crate's public API since we don't consider
// it stable enough to use on a live network.
func run_grandpa_observer() {
	// TODO says not live, should build?
}

// ObserverWork Future that powers the observer.
type ObserverWork struct {
	//observer:
	//Pin<Box<dyn Future<Output = Result<(), CommandOrError<B::Hash, NumberFor<B>>>> + Send>>,
	//client: Arc<Client>,
	//network: NetworkBridge<B, N, S>,
	//persistent_data: PersistentData<B>,
	//keystore: Option<KeystorePtr>,
	//voter_commands_rx: TracingUnboundedReceiver<VoterCommand<B::Hash, NumberFor<B>>>,
	//justification_sender: Option<GrandpaJustificationSender<B>>,
	//telemetry: Option<TelemetryHandle>,
	//_phantom: PhantomData<BE>,antom: PhantomData<Block>,PhantomData<Backend>,red_authority_set: Option<SharedAuthoritySet<Block::Hash, NumberFor<Block>>>,
}

func (s *ObserverWork) new() {}

// Rebuilds the `self.observer` field using the current authority set
// state. This method should be called when we know that the authority set
// has changed (e.g. as signalled by a voter command).
func (s *ObserverWork) rebuild_observer() {}

func (s *ObserverWork) handle_voter_command() {}

func (s *ObserverWork) poll() {}
