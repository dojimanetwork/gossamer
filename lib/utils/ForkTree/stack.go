// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package ForkTree

import "golang.org/x/exp/slices"

type stackElem[V any] struct {
	node node[V]
	val  uintptr
}

type stack[V any] []stackElem[V]

func (s stack[V]) push(elem stackElem[V]) {
	s = append(s, elem)
}

func (s stack[V]) pop() stack[V] {
	return slices.Delete(s, len(s)-1, len(s))
}
