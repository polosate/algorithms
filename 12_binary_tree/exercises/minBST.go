package exercises

import (
	"fmt"
)

type node struct {
	value      int
	leftChild  *node
	rightChild *node
}

type tree struct {
	root *node
}

func minBST(array []int) tree {
	return tree{
		root: makeTreeRec(array, 0, len(array)-1),
	}
}

func makeTreeRec(array []int, start, end int) *node {
	if end < start {
		return nil
	}
	mid := (start + end) / 2
	root := node{value: array[mid]}
	root.leftChild = makeTreeRec(array, start, mid-1)
	root.rightChild = makeTreeRec(array, mid+1, end)
	return &root
}

func (t *tree) display() {
	recDisplay(t.root)
}

func recDisplay(n *node) {
	if n == nil {
		return
	}
	fmt.Print(n.value)
	recDisplay(n.leftChild)
	recDisplay(n.rightChild)
}
