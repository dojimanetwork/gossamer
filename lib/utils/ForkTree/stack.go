// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package ForkTree

// TODO should I use this instead https://github.com/zeroflucs-given/generics

type stackElem[V any] struct {
	node node[V]
	val  uintptr
}

type stack[V any] []stackElem[V]

func (s *stack[V]) push(elem stackElem[V]) {
	*s = append(*s, elem)
}

func (s *stack[V]) pop() stackElem[V] {
	n := len(*s) - 1
	newS := *s
	val := newS[n]
	newS = newS[:n]
	*s = newS[:n]
	return val
}

func (s *stack[V]) len() int {
	return len(*s)
}

func (s *stack[V]) peek() stackElem[V] {
	newS := *s
	val := newS[len(*s)-1]
	return val
}
