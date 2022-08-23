package linked_list

import "testing"

func TestGet(t *testing.T) {
	root := Node[int]{Value: 7}
	val, err := root.get(7)

	if err != nil || val.Value != 7 {
		t.Error("val should equal 7")
	}

}

func TestPrepend(t *testing.T) {
	root := Node[int]{Value: 7}
	next := root.prepend(8)

	if next.Value != 8 {
		t.Error("val should equal 8")
	}

	if next.Next.Value != 7 {
		t.Error("val should equal 7")
	}
}

func TestDelete(t *testing.T) {
	root := &Node[int]{Value: 1}
	root, err := root.delete(1)
	if err != nil {
		t.Errorf("should not have errored, got %q", err)
	}
	if root != nil {
		t.Error("root should be nil")
	}

	// 1 -> 2 -> /
	root = &Node[int]{Value: 2}
	root = root.prepend(1)

	root, err = root.delete(1)
	if err != nil {
		t.Errorf("should not have errored, got %q", err)
	}
	if root.Value != 2 {
		t.Error("root value should be 2")
	}
	if root.Next != nil {
		t.Error("root.Next should be nil. i.e. end of list")
	}

	// 1 -> 2 -> /
	root = &Node[int]{Value: 2}
	root = root.prepend(1)

	root, err = root.delete(2)
	if err != nil {
		t.Errorf("should not have errored, got %q", err)
	}
	if root.Value != 1 {
		t.Error("root value should be 1")
	}
	if root.Next != nil {
		t.Error("root.Next should be nil. i.e. end of list")
	}

	root = &Node[int]{Value: 3}
	root = root.prepend(2)
	root = root.prepend(1)

	root, err = root.delete(2)
	if err != nil {
		t.Errorf("should not have errored, got %q", err)
	}

	if root.Value != 1 {
		t.Error("root value should be 1")
	}

	if root.Next == nil {
		t.Error("root.Next should not be nil. i.e. end of list")
	}

	if root.Next.Value != 3 {
		t.Error("root.Next.value should equal 3.")
	}

}
