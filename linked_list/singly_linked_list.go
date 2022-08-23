package linked_list

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func (n *Node[T]) prepend(value T) *Node[T] {
	return &Node[T]{Value: value, Next: n}
}

// cant be in-place because of how go represents pointers
func (n *Node[T]) delete(value T) (*Node[T], error) {
	// deleting empty
	if n == nil {
		return nil, errors.New("empty list")
	}

	// delete first
	if n.Value == value {
		// empty list
		if n.Next == nil {
			return nil, nil
		}

		*n = *n.Next
		return n, nil
	}

	cursor := n
	// peak ahead to find value
	for {
		next, err := cursor.peek()
		// could not delete
		// next is end
		if err != nil {
			return n, err
		}

		// did not find value; advance cursor
		if next.Value != value {
			cursor = next
			continue
		}

		// found case
		nextNext, err := next.peek()

		// next next is the end of list
		//
		// value = 2
		//   [ 1 ] -> [ 2 ] -> [ \ ]
		// (cursor)  (next)  (next next)
		if err != nil {
			// cursor is new end of list
			cursor.Next = nil
			return n, nil
		}

		cursor.Next = nextNext
		// drop pointer
		next.Next = nil
		return n, nil
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
