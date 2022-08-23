package graph

import (
	"fmt"

	"github.com/aaraney/go-trees/set"
)

type graphTuple[T comparable] struct {
	self T
	to   T
}

func (t graphTuple[T]) String() string {
	return fmt.Sprintf("%v --> %v", t.to, t.self)
}

func BFS[T comparable](g UGraph[T], root T) [][]graphTuple[T] {
	horizon := 0
	var levels [][]graphTuple[T]
	rootHorizon := []graphTuple[T]{{self: root}}

	seen := set.New[T]()
	seen.Add(root)
	levels = append(levels, rootHorizon)

	for {
		var level []graphTuple[T]

		for l := 0; l < len(levels[horizon]); l++ {
			e := levels[horizon][l]
			edges, _ := g.Get(e.self)

			iter := edges.Iter()
			for item := iter.Next(); item != nil; item = iter.Next() {

				if seen.In(item.Value) {
					continue
				}

				seen.Add(item.Value)
				level = append(level, graphTuple[T]{self: item.Value, to: e.self})
			}

		}
		horizon++
		levels = append(levels, level)
		if len(level) == 0 {
			return levels
		}
	}
}
