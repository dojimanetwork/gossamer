// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package grandpa

import (
	"github.com/ChainSafe/gossamer/dot/types"
	"github.com/ChainSafe/gossamer/lib/crypto/ed25519"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCurrentLimitFiltersMin(t *testing.T) {
	var currentAuthorities AuthorityList
	pubKey, err := ed25519.NewPublicKey([]byte{1})
	require.NoError(t, err)
	currentAuthorities[0] = types.Authority{
		Key:    pubKey,
		Weight: 1,
	}

	_ = AuthoritySet{
		currentAuthorities:     currentAuthorities,
		setId:                  0,
		pendingStandardChanges: ForkTree{},
		pendingForcedChanges:   []PendingChange{},
		authoritySetChanges:    AuthoritySetChanges{},
	}
}
