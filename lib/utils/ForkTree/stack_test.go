// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package ForkTree

import (
	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStackPush(t *testing.T) {
	elem0 := stackElem[int]{
		node: node[int]{
			Hash:   common.Hash{},
			Number: 0,
			Data:   0,
		},
		val: 0,
	}

	elem1 := stackElem[int]{
		node: node[int]{
			Hash:   common.Hash{},
			Number: 1,
			Data:   1,
		},
		val: 1,
	}

	s := stack[int]{elem0}
	require.True(t, len(s) == 1)

	s.push(elem1)
	require.True(t, len(s) == 2)
	require.Equal(t, s[len(s)-1], elem1)
}

func TestStackPop(t *testing.T) {
	elem0 := stackElem[int]{
		node: node[int]{
			Hash:   common.Hash{},
			Number: 0,
			Data:   0,
		},
		val: 0,
	}

	elem1 := stackElem[int]{
		node: node[int]{
			Hash:   common.Hash{},
			Number: 1,
			Data:   1,
		},
		val: 1,
	}

	s := stack[int]{elem0, elem1}
	require.True(t, s.len() == 2)

	val := s.pop()
	require.True(t, s.len() == 1)
	require.Equal(t, elem1, val)
	require.Equal(t, s[s.len()-1], elem0)
}

func TestStackPeek(t *testing.T) {
	elem0 := stackElem[int]{
		node: node[int]{
			Hash:   common.Hash{},
			Number: 0,
			Data:   0,
		},
		val: 0,
	}

	elem1 := stackElem[int]{
		node: node[int]{
			Hash:   common.Hash{},
			Number: 1,
			Data:   1,
		},
		val: 1,
	}

	s := stack[int]{elem0, elem1}
	require.True(t, s.len() == 2)

	val := s.peek()
	require.True(t, s.len() == 2)
	require.Equal(t, elem1, val)
	require.Equal(t, s[s.len()-1], elem1)
}
