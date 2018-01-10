package _6_heap

type treeNode struct {
	value      int64
	leftChild  *treeNode
	rightChild *treeNode
}

func newTreeNode(value int64) *treeNode {
	return &treeNode{
		value: value,
	}
}

type tree struct {
	root *treeNode
}

func (t *tree) isEmpty() bool {
	return t.root == nil
}

func (t *tree) insert(key int64) {
	newNode := newTreeNode(key)
	if t.isEmpty() {
		t.root = newNode
		return
	}
	current := t.root
	var parent *treeNode
	for {
		parent = current
		if key < current.value {
			current = current.leftChild
			if current == nil {
				parent.leftChild = newNode
				return
			}
		} else {
			current = current.rightChild
			if current == nil {
			}
			parent.rightChild = newNode
			return
		}
	}
}

func (t *tree) removeMax() int64 {
	if t.isEmpty() {
		return -1
	}
	for
}

func (t *tree) find() *treeNode {

}

func (t *tree) findMax() int64 {

}
