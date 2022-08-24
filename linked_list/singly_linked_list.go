package linked_list

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func (n Node[T]) String() string {
	return fmt.Sprintf("Value: %#v Next: %#v", n.Value, n.Next)
}

func (n *Node[T]) prepend(value T) *Node[T] {
	return &Node[T]{Value: value, Next: n}
}

// cant be in-place because of how go represents pointers
func (n *Node[T]) delete(value T) (*Node[T], error) {
	var previous *Node[T]
	current := n

	for {
		if current == nil {
			return n, fmt.Errorf("value, %v, not in list", value)
		}

		// found value
		if current.Value == value {
			// case when `current` is the fist item in list
			if previous == nil {
				root := current.Next
				current.Next = nil
				return root, nil
			}

			previous.Next = current.Next
			// drop pointer from deleted item to it's next element
			current.Next = nil
			return n, nil
		}

		// next
		previous = current
		current = current.Next
	}
}

func (n *Node[T]) get(value T) (*Node[T], error) {
	node := n

	for {
		if node.Value == value {
			return node, nil
		}
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return nil, fmt.Errorf("value, %v, not in list", value)
}

func (n *Node[T]) peek() (*Node[T], error) {
	if n.Next == nil {
		return nil, errors.New("end of list")
	}
	return n.Next, nil
}
