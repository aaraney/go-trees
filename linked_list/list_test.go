package linked_list

import "testing"

func TestListPrepend(t *testing.T) {
	root := NewList(1, 2)
	// root.Prepend(1)

	node := root.Root()

	if node.Value != 1 {
		t.Error("val should equal 1")
	}

	if node.Next.Value != 2 {
		t.Error("val should equal 2")
	}
}

func TestListIter(t *testing.T) {
	root := NewList(2)
	root.Prepend(1)
}

func TestNewList(t *testing.T) {
	root := NewList[int]()
	root.Prepend(1)

	if root.root.Value != 1 {
		t.Logf("%#v", root)
		t.Error("root should equal 1")
	}
}

func TestIter(t *testing.T) {
	items := []int{0, 1, 2, 3}
	root := NewList(items...)
	iter := root.Iter()

	i := 0
	for node := iter.Next(); node != nil; node = iter.Next() {
		if node.Value != items[i] {
			t.Errorf("node, %d, should equal %d", node.Value, items[i])
		}
		i++
	}

}
