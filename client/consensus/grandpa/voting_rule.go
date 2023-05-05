// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package grandpa

// TODO probably move to interfaces.go
type VotingRule interface {
	restrictVote()
}
