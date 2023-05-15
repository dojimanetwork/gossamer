// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package grandpa

import (
	"github.com/ChainSafe/gossamer/dot/types"
	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/ChainSafe/gossamer/lib/crypto/ed25519"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func staticIsDescendentOf(value bool) IsDescendentOf {
	return func(common.Hash, common.Hash) (bool, error) { return value, nil }
}

func TestCurrentLimitFiltersMin(t *testing.T) {
	var currentAuthorities AuthorityList
	kp, err := ed25519.GenerateKeypair()
	require.NoError(t, err)
	currentAuthorities = append(currentAuthorities, types.Authority{
		Key:    kp.Public(),
		Weight: 1,
	})

	finalizedKind := Finalized{}
	delayKind := NewDelayKind(finalizedKind)

	pendingChange1 := PendingChange{
		nextAuthorities: currentAuthorities,
		delay:           0,
		canonHeight:     1,
		canonHash:       common.BytesToHash([]byte{1}),
		delayKind:       delayKind,
	}

	pendingChange2 := PendingChange{
		nextAuthorities: currentAuthorities,
		delay:           0,
		canonHeight:     2,
		canonHash:       common.BytesToHash([]byte{2}),
		delayKind:       delayKind,
	}

	ctrl := gomock.NewController(t)
	mockForkTree1 := NewMockForkTree(ctrl)
	mockForkTree1.EXPECT().Import(common.BytesToHash([]byte{1}), uint(1), pendingChange1, gomock.Any())

	mockForkTree2 := NewMockForkTree(ctrl)
	mockForkTree2.EXPECT().Import(common.BytesToHash([]byte{2}), uint(2), pendingChange2, gomock.Any())

	authorities1 := AuthoritySet{
		currentAuthorities:     currentAuthorities,
		setId:                  0,
		pendingStandardChanges: mockForkTree1,
		pendingForcedChanges:   []PendingChange{},
		authoritySetChanges:    AuthoritySetChanges{},
	}

	authorities2 := AuthoritySet{
		currentAuthorities:     currentAuthorities,
		setId:                  0,
		pendingStandardChanges: mockForkTree2,
		pendingForcedChanges:   []PendingChange{},
		authoritySetChanges:    AuthoritySetChanges{},
	}

	err = authorities1.AddPendingChange(pendingChange1, staticIsDescendentOf(false))
	require.NoError(t, err)

	err = authorities2.AddPendingChange(pendingChange2, staticIsDescendentOf(false))
	require.NoError(t, err)

	//require.Equal(t, authorities.current)

}

func TestAuthoritySet_InvalidAuthorityList(t *testing.T) {
	type args struct {
		authorities  AuthorityList
		authoritySet AuthoritySet
	}
	tests := []struct {
		name string
		args args
		exp  bool
	}{
		// TODO: Add test cases.
		{
			name: "nil authorities",
			args: args{
				authorities:  nil,
				authoritySet: AuthoritySet{},
			},
		},
		{
			name: "empty authorities",
			args: args{
				authorities:  AuthorityList{},
				authoritySet: AuthoritySet{},
			},
		},
		{
			name: "invalid authorities weight",
			args: args{
				authorities: AuthorityList{
					types.Authority{
						Weight: 0,
					},
				},
				authoritySet: AuthoritySet{},
			},
		},
		{
			name: "valid authority list",
			args: args{
				authorities: AuthorityList{
					types.Authority{
						Weight: 1,
					},
				},
				authoritySet: AuthoritySet{},
			},
			exp: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.args.authoritySet.InvalidAuthorityList(tt.args.authorities); got != tt.exp {
				t.Errorf("InvalidAuthorityList() = %v, want %v", got, tt.exp)
			}
		})
	}
}
