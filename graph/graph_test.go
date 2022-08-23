package graph

import "testing"

func TestUGraph(t *testing.T) {
	g := New[int]()
	g.Insert(1, 2, 4)
	g.Insert(2, 3, 4)
	g.Insert(3, 4, 5)
	g.Insert(5, 6)

	t.Logf("\n%s\n", g)
}
