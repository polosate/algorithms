package _2_binary_tree

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
	current, parent := t.find(key)
	if current == nil {
		return nil
	}

	// Element is root
	if current == t.root {
		if current.rightChild == nil && current.leftChild == nil {
			t.root = nil
			return current
		}
		if current.rightChild == nil {
			t.root = current.leftChild
			return current
		}
		if current.leftChild == nil {
			t.root = current.rightChild
			return current
		}
	}

	isRightChild := false
	if parent.rightChild != nil && parent.rightChild.GetValue() == key {
		isRightChild = true
	}

	// Element is leaf
	if current.rightChild == nil && current.leftChild == nil {
		if isRightChild {
			parent.rightChild = nil
			return current
		} else {
			parent.leftChild = nil
			return current
		}
	}

	// Element has only left child
	if current.rightChild == nil {
		if isRightChild {
			parent.rightChild = current.leftChild
			return current
		} else {
			parent.leftChild = current.leftChild
			return current
		}
	}

	// Element has only right child
	if current.leftChild == nil {
		if isRightChild {
			parent.rightChild = current.rightChild
			return current
		} else {
			parent.leftChild = current.rightChild
			return current
		}
	}

	return nil
}

func (t *Tree) getSuccessor(node *Node) *Node {
	

	successor := node.rightChild
	for successor.leftChild != nil {
		successor = successor.leftChild
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
