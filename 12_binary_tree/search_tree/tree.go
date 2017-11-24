package search_tree

type Tree struct {
	root *Node
}

func NewTree() Tree {
	return Tree{root: nil}
}

func (t *Tree) Find(key int64) *Node {
	current, _ := t.find(key)
	return current
}

func (t *Tree) Insert(value int64) {
	newNode := NewNode(value)
	if t.root == nil {
		t.root = newNode
		return
	}
	current := t.root
	var parent *Node
	for {
		parent = current
		if value < current.GetValue() {
			current = current.leftChild
			if current == nil {
				parent.leftChild = newNode
				return
			}
		} else {
			current = current.rightChild
			if current == nil {
				parent.rightChild = newNode
				return
			}
		}
	}
}

func (t *Tree) TraversInOrder() {
	t.recTraversInOrder(t.root)

}

func (t *Tree) Remove(key int64) *Node {
	delNode, parent := t.find(key)
	if delNode == nil {
		return nil
	}

	// delNode is root
	if delNode == t.root {
		if delNode.rightChild == nil && delNode.leftChild == nil {
			t.root = nil
			return delNode
		}
		if delNode.rightChild == nil {
			t.root = delNode.leftChild
			return delNode
		}
		if delNode.leftChild == nil {
			t.root = delNode.rightChild
			return delNode
		}
		successor := t.getSuccessor(delNode)
		t.root = successor
		t.root.leftChild = delNode.leftChild
		return delNode
	}

	isRightChild := false
	if parent.rightChild != nil && parent.rightChild.GetValue() == key {
		isRightChild = true
	}

	// delNode is leaf
	if delNode.rightChild == nil && delNode.leftChild == nil {
		if isRightChild {
			parent.rightChild = nil
			return delNode
		} else {
			parent.leftChild = nil
			return delNode
		}
	}

	// delNode has only left child
	if delNode.rightChild == nil {
		if isRightChild {
			parent.rightChild = delNode.leftChild
			return delNode
		} else {
			parent.leftChild = delNode.leftChild
			return delNode
		}
	}

	// delNode has only right child
	if delNode.leftChild == nil {
		if isRightChild {
			parent.rightChild = delNode.rightChild
			return delNode
		} else {
			parent.leftChild = delNode.rightChild
			return delNode
		}
	}

	// delNode has both children
	successor := t.getSuccessor(delNode)
	if isRightChild {
		parent.rightChild = successor
	} else {
		parent.leftChild = successor
	}
	successor.leftChild = delNode.leftChild
	return delNode
}

func (t *Tree) getSuccessor(delNode *Node) *Node {
	successor := delNode
	successorParent := delNode
	current := successor.rightChild

	for current != nil {
		successorParent = successor
		successor = current
		current = current.leftChild
	}

	if successor != delNode.rightChild {
		successorParent.leftChild = successor.rightChild
		successor.rightChild = delNode.rightChild
	}
	return successor
}

func (t *Tree) find(key int64) (node *Node, parent *Node) {
	current := t.root
	for current != nil && current.GetValue() != key {
		parent = current
		if key < current.GetValue() {
			current = current.leftChild
		} else {
			current = current.rightChild
		}
	}
	return current, parent
}

func (t *Tree) recTraversInOrder(currentRoot *Node) {
	if currentRoot != nil {
		t.recTraversInOrder(currentRoot.leftChild)
		currentRoot.DisplayNode()
		t.recTraversInOrder(currentRoot.rightChild)
	}
}
