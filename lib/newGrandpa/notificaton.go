// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package newGrandpa

// TODO These will all probably be channels, just structs for now

// GrandpaJustificationSender The sending half of the Grandpa justification channel(s).
//
// Used to send notifications about justifications generated
// at the end of a Grandpa round.
type GrandpaJustificationSender struct{}

// GrandpaJustificationStream The receiving half of the Grandpa justification channel.
//
// Used to receive notifications about justifications generated
// at the end of a Grandpa round.
// The `GrandpaJustificationStream` entity stores the `SharedJustificationSenders`
// so it can be used to add more subscriptions.
type GrandpaJustificationStream struct{}

// GrandpaJustificationsTracingKey Provides tracing key for GRANDPA justifications stream.
type GrandpaJustificationsTracingKey struct{}
