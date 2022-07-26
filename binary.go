package main

import (
	"fmt"
)

type tree struct {
	val   int
	left  *tree
	right *tree
}

func (t *tree) insert(val int) {
	// left
	if val < t.val {
		if t.left == nil {
			t.left = &tree{val: val}
			return
		}
		t.left.insert(val)
	} else {

		if t.right == nil {
			t.right = &tree{val: val}
			return
		}
		t.right.insert(val)
	}

}

func (t *tree) batchInsert(vals ...int) {
	for _, val := range vals {
		t.insert(val)
	}
}

func (t *tree) isLeaf() bool {
	if t != nil {
		return t.left == nil && t.right == nil
	}
	return false
}

func levelOrder(t *tree, level int) []int {
	if t == nil {
		return []int{}
	}
	if level == 1 {
		return []int{t.val}
	}
	return append(levelOrder(t.left, level-1), levelOrder(t.right, level-1)...)
}

func allLevels(t *tree) [][]int {
	var outer [][]int

	for i := 1; ; i++ {
		inner := levelOrder(t, i)
		if len(inner) == 0 {
			break
		}
		outer = append(outer, inner)
	}

	return outer

}

func helper(t *tree, vs *[]int) {
	if t != nil {
		helper(t.left, vs)
		*vs = append(*vs, t.val)
		helper(t.right, vs)
	}
}

func inOrder(t *tree) []int {
	var l []int
	helper(t, &l)
	return l
}

func preHelper(t *tree, vs *[]int) {
	if t != nil {
		*vs = append(*vs, t.val)
		preHelper(t.left, vs)
		preHelper(t.right, vs)
	}
}

func preOrder(t *tree) []int {
	var l []int
	preHelper(t, &l)
	return l
}

func postHelper(t *tree, vs *[]int) {
	if t != nil {
		postHelper(t.left, vs)
		postHelper(t.right, vs)
		*vs = append(*vs, t.val)
	}
}

func postOrder(t *tree) []int {
	var l []int
	postHelper(t, &l)
	return l
}

func reverse(t *tree) {
	if t != nil {
		reverse(t.left)

		if t.left != nil || t.right != nil {
			t.left, t.right = t.right, t.left
			// go left b.c. swap. right is now left
			reverse(t.left)
		}
	}
}

func main() {
	//   5
	//  4 6
	// 3

	// root := &tree{val: 5}
	// root.insert(4)
	// root.insert(3)
	// root.insert(6)
	// fmt.Printf("%v\n", root.left.left.val)

	preOrderVal := []int{25, 15, 10, 4, 12, 22, 18, 24, 50, 35, 31, 44, 70, 66, 90}
	root := &tree{val: preOrderVal[0]}
	root.batchInsert(preOrderVal[1:]...)

	fmt.Printf("level-order: %v\n", levelOrder(root, 2))
	fmt.Printf("all-levels: %v\n", allLevels(root))
	fmt.Printf("in-order: %v\n", inOrder(root))
	fmt.Printf("pre-order: %v\n", preOrder(root))
	fmt.Printf("post-order: %v\n", postOrder(root))
	reverse(root)
	fmt.Printf("reversed: %v\n", allLevels(root))
}
