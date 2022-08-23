package graph

import (
	"github.com/aaraney/go-trees/set"
)

func DFS[T comparable](g UGraph[T], root T) []T {
	nodes := []T{root}
	seen := set.New[T]()
	seen.Add(root)

	DFS_helper(g, &root, &nodes, seen)
	return nodes
}

func DFS_helper[T comparable](g UGraph[T], root *T, nodes *[]T, seen set.Set[T]) {
	edges, _ := g.Get(*root)

	iter := edges.Iter()
	for item := iter.Next(); item != nil; item = iter.Next() {

		if seen.In(item.Value) {
			continue
		}

		seen.Add(item.Value)
		*nodes = append(*nodes, item.Value)
		DFS_helper(g, &item.Value, nodes, seen)
	}

}
