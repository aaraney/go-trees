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

func (h MinHeap[T]) HeapifyUp(idx int) {
	parent := h.Parent(idx)
	if parent < 0 {
		return
	}

	if idx < len(h) && h[idx] < h[parent] {
		h[idx], h[parent] = h[parent], h[idx]
		h.HeapifyUp(parent)
	}
}

func (h MinHeap[T]) HeapifyDown(parentIdx int) {
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
		return
	}

	h[min], h[parentIdx] = h[parentIdx], h[min]
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
	h.HeapifyUp(len(*h) - 1)
}
