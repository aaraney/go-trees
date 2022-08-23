package graph

import (
	"fmt"

	"github.com/aaraney/go-trees/linked_list"
)

type UGraph[T comparable] map[T]*linked_list.List[T]

func New[T comparable]() UGraph[T] {
	return make(UGraph[T])
}

func (l UGraph[T]) Insert(vert T, edges ...T) {
	l.insert(vert, edges...)

	// update edge connections
	l.update(vert, edges...)
}

func (l UGraph[T]) Get(vert T) (*linked_list.List[T], error) {
	v, ok := l[vert]
	if !ok {
		return nil, fmt.Errorf("vertex, %v, not in graph", vert)
	}

	return v, nil
}

func (l UGraph[T]) Delete(vert T) {
	v, ok := l[vert]

	// vert not in map, so also not in any adjacency list
	if !ok {
		return
	}

	iter := v.Iter()
	for edge := iter.Next(); edge != nil; edge = iter.Next() {
		// remove instances of `vert` from other adjacency lists
		l[edge.Value].Delete(vert)
	}

	// remove key and value from map
	delete(l, vert)
}

func (l UGraph[T]) update(vert T, edges ...T) {
	for _, edge := range edges {
		l.insert(edge, vert)
	}
}

func (l UGraph[T]) insert(vert T, edges ...T) {
	v, ok := l[vert]

	if !ok {
		l[vert] = linked_list.NewList(edges...)
		return
	}

	v.Prepend(edges...)
}

func (l UGraph[T]) String() string {
	var s []byte

	for key, val := range l {
		buf := []byte(fmt.Sprintf("vert: %v\n", key))

		iter := val.Iter()
		for item := iter.Next(); item != nil; item = iter.Next() {
			buf = append(buf, []byte(fmt.Sprintf("%v ", item.Value))...)
		}
		buf = append(buf, []byte("\n")...)

		s = append(s, buf...)
	}

	return fmt.Sprintf("%s", s)
}
