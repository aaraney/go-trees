package linked_list

type List[T comparable] struct {
	root *Node[T]
}

func NewList[T comparable](values ...T) *List[T] {
	var node *Node[T]

	for i := len(values) - 1; i > -1; i-- {
		node = &Node[T]{Value: values[i], Next: node}
	}

	return &List[T]{root: node}
}

func (n *List[T]) Root() *Node[T] {
	return n.root
}

func (n *List[T]) Prepend(values ...T) {
	for i := len(values) - 1; i > -1; i-- {
		n.root = n.root.prepend(values[i])
	}
}

func (n *List[T]) Delete(value T) error {
	newRoot, err := n.root.delete(value)
	n.root = newRoot
	return err
}

func (n *List[T]) Get(value T) (*Node[T], error) {
	return n.root.get(value)
}

func (n *List[T]) Peek() (*Node[T], error) {
	return n.root.peek()
}

func (n *List[T]) Iter() Iterator[*Node[T]] {
	return &iter[T]{state: n}
}

type Iterator[T any] interface {
	Next() T
}

type iter[T comparable] struct {
	state *List[T]
}

func (i *iter[T]) Next() *Node[T] {
	if i == nil || i.state == nil {
		return nil
	}

	// to yield
	root := i.state.Root()

	if root == nil {
		return nil
	}

	next, _ := i.state.Peek()
	i.state = &List[T]{root: next}

	return root
}
