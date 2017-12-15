package _4_2_3_4_Tree

import "fmt"

type node struct {
	childArray []*node
	itemArray  []*dataItem
	numItems   int
	parent     *node
}

func NewNode() *node {
	return &node{
		childArray: make([]*node, 4),
		itemArray:  make([]*dataItem, 3),
		numItems:   0,
		parent:     nil,
	}
}

func (n *node) findItem(key int64) *dataItem {
}

func (n *node) inertItem(key int64) {

}

func (n *node) removeItem(key int64) *dataItem {
}

func (n *node) displayNode() {
	for _, v := range n.itemArray {
		v.displayDataItem()
	}
	fmt.Print("/")
}
