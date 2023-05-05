// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package grandpa

// TODO investigate how to impl this, this is just filler
type SharedAuthoritySet struct{}

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
	inner        ClientForGrandpa
	selectChain  SelectChain
	authoritySet SharedAuthoritySet
	// TODO figure out best way to do this in go
	//sendVoterCommands TracingUnboundedSender
	// TODO investigate this
	//authoritySetHardForks HashMap<K, V, S = RandomState>
	// TODO more communication stuff
	//justification_sender GrandpaJustificationSender
	telemetry Telemetry
	// TODO do we need backend? Do we ned PhantomData
	//phantom PhantomData<Backend>
}
