package __3_4

import "fmt"

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
	for {
		if curNode.isFull() {
			t.split(curNode)
			curNode = curNode.getParent()
			curNode = getNextChild(curNode, key)
		} else if curNode.isLeaf() {
			break
		} else {
			curNode = getNextChild(curNode, key)
		}
	}
	curNode.insertItem(newItem)
}

func (t *tree) split(curNode *node) {
	var (
		child2, child3, parent *node
		itemB, itemC           *dataItem
	)

	itemC = curNode.removeItem()
	itemB = curNode.removeItem()
	child3 = curNode.disconnectChild(3)
	child2 = curNode.disconnectChild(2)
	rightNode := newNode()

	if curNode == t.root {
		t.root = newNode()
		parent = t.root
		t.root.connectChild(0, curNode)
	} else {
		parent = curNode.getParent()
	}
	itemIndex := parent.insertItem(itemB)
	n := parent.numItems
	for i := n - 1; i > itemIndex; i-- {
		tempChild := parent.disconnectChild(i)
		parent.connectChild(i+1, tempChild)
	}
	parent.connectChild(itemIndex+1, rightNode)

	rightNode.insertItem(itemC)
	rightNode.connectChild(0, child2)
	rightNode.connectChild(1, child3)
}

func (t *tree) min() int64 {
	curNode := t.root
	for !curNode.isLeaf() {
		curNode = curNode.getChild(0)
	}
	return curNode.getItem(0).key
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
			return curNode.childArray[i]
		}
	}
	return curNode.childArray[i]
}
