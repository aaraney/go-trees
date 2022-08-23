package graph

import "testing"

func TestBFS(t *testing.T) {
	g := New[int]()
	g.Insert(1, 2, 4)
	g.Insert(2, 3, 4)
	g.Insert(3, 4, 5)
	g.Insert(5, 6)

	paths := BFS(g, 1)

	for _, level := range paths {
		t.Logf("\n%s\n", level)
	}

}
