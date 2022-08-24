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

func TestDeleteListItem(t *testing.T) {
	root := NewList(1, 2)
	one, _ := root.Get(1)

	err := root.Delete(1)
	if err != nil {
		t.Errorf("should not have received err: %q", err)
	}

	if root.Root().Value != 2 {
		t.Error("root should have value of 2")
	}

	if one.Next != nil {
		t.Error("deleted item should have nil Next pointer", err)
		t.Errorf("%s", one)
	}

	root = NewList(1, 2, 3)
	root.Delete(2)
	root.Delete(1)

	if root.Root().Value != 3 {
		t.Error("root should have value of 3")
	}

	iter := root.Iter()
	target_n := 1
	n := 0

	// count number of items in list
	for item := iter.Next(); item != nil; item = iter.Next() {
		n++
	}

	if target_n != n {
		t.Errorf("expected: %d got: %d. %d != %d. too many items in list.", target_n, n, target_n, n)
	}
}
