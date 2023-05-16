// Copyright 2023 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

// Package ForkTree
// Utility library for managing tree-like ordered data with logic for pruning
// the tree while finalizing nodes.
package ForkTree

import (
	"fmt"
	"github.com/ChainSafe/gossamer/lib/common"
)

// ForkTree A tree data structure that stores several nodes across multiple branches.
//
// Top-level branches are called roots. The tree has functionality for
// finalizing nodes, which means that node is traversed, and all competing
// branches are pruned. It also guarantees that nodes in the tree are finalized
// in order. Each node is uniquely identified by its hash but can be ordered by
// its number. In order to build the tree an external function must be provided
// when interacting with the tree to establish a node's ancestry.
type ForkTree[V any] struct {
	roots               []node[V]
	bestFinalizedNumber *uint
}

// IsDescendentOf is a type to represent the function signature of a IsDescendentOf function
type IsDescendentOf func(h1 common.Hash, h2 common.Hash) (bool, error)

type predicate func(V any) bool

func (ft *ForkTree[V]) Import(hash common.Hash, number uint, data V, isDescendentOf IsDescendentOf) (bool, error) {
	// TODO impl
	if ft.bestFinalizedNumber != nil {
		if number <= *ft.bestFinalizedNumber {
			return false, fmt.Errorf("tried to import or finalize node that is an ancestor of a previously finalized node")
		}
	}

	// seems this always returns true???
	p := func(V any) bool {
		return true
	}
	ft.findNodeWhere(hash, number, isDescendentOf, p)
	return false, nil
}

// Same as [`find_node_where`](ForkTree::find_node_where), but returns mutable reference.
func (ft *ForkTree[V]) findNodeWhere(hash common.Hash, number uint, isDescendentOf IsDescendentOf, predicate predicate) {
	_, _ = ft.findNodeIndexWhere(hash, number, isDescendentOf, predicate)
}

// Same as [`find_node_where`](ForkTree::find_node_where), but returns indices.
//
// The returned indices represent the full path to reach the matching node starting
// from one of the roots, i.e. the earliest index in the traverse path goes first,
// and the final index in the traverse path goes last.
//
// If a node is found that matches the predicate the returned path should always
// contain at least one index, otherwise `None` is returned.
//
// WARNING: some users of this method (i.e. consensus epoch changes tree) currently silently
// rely on a **post-order DFS** traversal. If we are using instead a top-down traversal method
// then the `is_descendent_of` closure, when used after a warp-sync, will end up querying the
// backend for a block (the one corresponding to the root) that is not present and thus will
// return a wrong result.
func (ft *ForkTree[V]) findNodeIndexWhere(hash common.Hash, number uint, isDescendentOf IsDescendentOf, predicate predicate) (*[]uintptr, error) {
	var stack stack[V]
	rootIdx := 0
	//found := false
	//isDescendent := false

	for rootIdx < len(ft.roots) {
		if number <= ft.roots[rootIdx].Number {
			rootIdx += 1
			continue
		}

		// The second element in the stack tuple tracks what is the **next** children
		// index to search into. If we find an ancestor then we stop searching into
		// alternative branches and we focus on the current path up to the root.
		newStackElem := stackElem[V]{
			node: ft.roots[rootIdx],
			val:  0,
		}
		stack.push(newStackElem)
	}
	return nil, nil
}

type node[V any] struct {
	Hash     common.Hash
	Number   uint
	Data     V
	Children []node[V]
}
