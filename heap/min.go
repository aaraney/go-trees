package heap

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type MinHeap[T constraints.Ordered] []T

func (h MinHeap[T]) BuildHeap() {
	for start := h.lastParent(); start >= 0; start-- {
		h.HeapifyDown(start)
	}
}

// return min, parent tuple
func (h MinHeap[T]) heapify(parentIdx int) (int, int) {
	left := h.LeftChild(parentIdx)
	right := h.RightChild(parentIdx)

	min := parentIdx

	if left < len(h) && h[left] < h[min] {
		min = left
	}

	if right < len(h) && h[right] < h[min] {
		min = right
	}

	// parent is min
	if min == parentIdx {
		return parentIdx, parentIdx
	}

	h[min], h[parentIdx] = h[parentIdx], h[min]
	return min, parentIdx
}

func (h MinHeap[T]) HeapifyUp(parentIdx int) {
	min, parent := h.heapify(parentIdx)
	grand_parent := h.Parent(parent)
	if min == parent || grand_parent < 0 {
		return
	}
	h.HeapifyUp(grand_parent)
}

func (h MinHeap[T]) HeapifyDown(parentIdx int) {
	min, parent := h.heapify(parentIdx)
	if min == parent {
		return
	}
	h.HeapifyDown(min)

}

// idx: 0 1 2 3 4 5 6
// val: 1 2 3 4 5 6 7
//
//     1
//  2    3
// 4 5  6 7
//
// parent: i // 2
// left : i*2 + 1 (b.c. zero counting)
// right: i*2 + 2
func (h MinHeap[T]) Parent(i int) int {
	// floor divide by two
	return i >> 1
}

func (h MinHeap[T]) LeftChild(i int) int {
	return 2*i + 1
}

func (h MinHeap[T]) RightChild(i int) int {
	return 2*i + 2
}

func (h MinHeap[T]) lastParent() int {
	return len(h)>>1 - 1
}

func (h *MinHeap[T]) Pop() (T, error) {

	if len(*h) == 0 {
		return *new(T), errors.New("empty heap")
	}

	min := (*h)[0]
	*h = (*h)[1:]
	h.HeapifyDown(0)
	return min, nil
}

func (h *MinHeap[T]) Insert(element T) {
	*h = append(*h, element)
	h.HeapifyUp(h.lastParent())
}
