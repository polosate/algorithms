package __3_4

import (
	"fmt"
)

type tree struct {
	root *node
}

func newTree() tree {
	return tree{
		root: newNode(),
	}
}

func (t *tree) find(key int64) int {
	curNode := t.root
	var childNumber int

	for {
		if childNumber = curNode.findItem(key); childNumber != -1 {
			return childNumber
		} else if curNode.isLeaf() {
			return -1
		} else {
			curNode = getNextChild(curNode, key)
		}
	}

}

func (t *tree) insert(key int64) {
	newItem := newDataItem(key)
	curNode := t.root

	for !curNode.isLeaf() {
		curNode = getNextChild(curNode, key)
	}

	if curNode.getNumItems() == 1 {
		curNode.insertItem(newItem)
	} else {
		t.split(curNode, newItem)
	}
}

func (t *tree) split(curNode *node, item *dataItem) {
	rightNode := newNode()

}

func (t *tree) displayTree() {
	t.recDisplayTree(t.root, 0, 0)
	fmt.Println()
}

func (t *tree) recDisplayTree(curNode *node, level int, childNum int) {
	fmt.Print("Level=", level, " ChildNumber=", childNum, "  ")
	curNode.displayNode()

	numItems := curNode.getNumItems()
	for i := 0; i < numItems+1; i++ {
		nextNode := curNode.getChild(i)
		if nextNode != nil {
			t.recDisplayTree(nextNode, level+1, i)
		} else {
			return
		}
	}
}

func getNextChild(curNode *node, key int64) *node {
	i := 0
	for ; i < curNode.numItems; i++ {
		if curNode.itemArray[i].key > key {
			return curNode.getChild(i)
		}
	}
	return curNode.childArray[i]
}
