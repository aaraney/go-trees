package graph

import "testing"

func TestDFS(t *testing.T) {
	g := New[int]()
	g.Insert(1, 2, 4)
	g.Insert(2, 3, 4)
	g.Insert(3, 4, 5)
	g.Insert(5, 6)

	paths := DFS(g, 1)

	t.Logf("\n%v\n", paths)

}
