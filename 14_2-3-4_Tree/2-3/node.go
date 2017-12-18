package __3_4

import "fmt"

const order = 3

type node struct {
	childArray []*node
	itemArray  []*dataItem
	numItems   int
	parent     *node
}

func newNode() *node {
	return &node{
		childArray: make([]*node, order),
		itemArray:  make([]*dataItem, order-1),
		numItems:   0,
		parent:     nil,
	}
}

func (n *node) findItem(key int64) int {
	for i := 0; i < order; i++ {
		if n.itemArray[i] == nil {
			break
		}
		if n.itemArray[i].key == key {
			return i
		}
	}
	return -1
}

func (n *node) insertItem(newItem *dataItem) int {
	n.numItems++
	newKey := newItem.key

	if newKey < n.itemArray[0].key {
		n.itemArray[1] = n.itemArray[0]
		n.itemArray[0] = newItem
		return 0
	} else {
		n.itemArray[1] = newItem
		return 1
	}
}

func (n *node) removeItem() *dataItem {
	temp := n.itemArray[n.numItems-1]
	n.itemArray[n.numItems-1] = nil
	n.numItems--
	return temp
}

func (n *node) connectChild(childNum int, child *node) {
	n.childArray[childNum] = child
	if child != nil {
		child.parent = n
	}

}

func (n *node) disconnectChild(childNum int) *node {
	tempChild := n.childArray[childNum]
	n.childArray[childNum] = nil
	return tempChild

}

func (n *node) getChild(childNum int) *node {
	return n.childArray[childNum]
}

func (n *node) getItem(itemNum int) *dataItem {
	return n.itemArray[itemNum]
}

func (n *node) getParent() *node {
	return n.parent
}

func (n *node) isLeaf() bool {
	return n.childArray[0] == nil
}

func (n *node) isFull() bool {
	return n.numItems == order-1
}

func (n *node) getNumItems() int {
	return n.numItems
}

func (n *node) displayNode() {
	for i := 0; i < n.numItems; i++ {
		n.itemArray[i].displayDataItem()
	}
	fmt.Print("/  ")
}
