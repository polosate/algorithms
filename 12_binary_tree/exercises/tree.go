package exercises

import "math"

//type nodeLink struct {
//	node Node
//	next *nodeLink
//}
//
//func NewNodeLink(node Node) *nodeLink {
//	return &nodeLink{
//		node: node,
//	}
//}
//
//type nodeList struct {
//	first *nodeLink
//}
//
//func NewNodeList() nodeList {
//	return nodeList{}
//}
//
//func (l *nodeList) IsEmpty() bool {
//	return l.first == nil
//}
//
//func (l *nodeList) InsertFirst(node Node) {
//	newLink := NewNodeLink(node)
//	newLink.next = l.first
//	l.first = newLink
//}
//
//func (l *nodeList) RemoveFirst() Node {
//	temp := l.first
//	l.first = temp.next
//	return temp.node
//}

type link struct {
	tree Tree
	next *link
}

func NewLink(node Tree) *link {
	return &link{tree: node}
}

type list struct {
	first *link
	last  *link
}

func NewList() list {
	return list{}
}

func (l *list) IsEmpty() bool {
	return l.first == nil
}

func (l *list) InsertFirst(tree Tree) {
	newLink := NewLink(tree)
	if l.IsEmpty() {
		l.last = newLink
	}
	newLink.next = l.first
	l.first = newLink
}

func (l *list) InsertLast(tree Tree) {
	newLink := NewLink(tree)
	if l.IsEmpty() {
		l.first = newLink
		l.last = newLink
		return
	}
	l.last.next = newLink
	l.last = newLink
}

func (l *list) RemoveFirst() Tree {
	if l.first == l.last {
		l.last = nil
	}
	temp := l.first
	l.first = temp.next
	return temp.tree
}

type Node struct {
	value      string
	leftChild  *Node
	rightChild *Node
}

func NewNode(value string) *Node {
	return &Node{value: value}
}

type Tree struct {
	root *Node
}

func newTree(root *Node) *Tree {
	return &Tree{root: root}
}

func NewTree(array []string) Tree {
	newList := NewList()
	for _, v := range array {
		newNode := NewNode(v)
		newTree := newTree(newNode)
		newList.InsertFirst(*newTree)
	}
	for newList.first.next != nil {
		leaf1 := newList.RemoveFirst()
		leaf2 := newList.RemoveFirst()
		resultTree := newTree(NewNode("+"))
		resultTree.root.leftChild = leaf1.root
		resultTree.root.rightChild = leaf2.root
		newList.InsertFirst(*resultTree)
	}
	return newList.RemoveFirst()
}

func NewBalancedTree(array []string) Tree {
	newList := NewList()
	for _, v := range array {
		newNode := NewNode(v)
		newTree := newTree(newNode)
		newList.InsertFirst(*newTree)
	}

	for newList.first.next != nil {
		l1 := newList.RemoveFirst()
		l2 := newList.RemoveFirst()
		resTree := newTree(NewNode("+"))
		resTree.root.leftChild = l1.root
		resTree.root.rightChild = l2.root
		newList.InsertLast(*resTree)
	}
	return newList.RemoveFirst()
}

func (t *Tree) isBalanced() bool {
	if checkHeight(t.root) == -1 {
		return false
	}
	return true
}

func checkHeight(n *Node) int {
	if n == nil {
		return 0
	}
	leftChild := checkHeight(n.leftChild)
	if leftChild == -1 {
		return -1
	}
	rightChild := checkHeight(n.rightChild)
	if rightChild == -1 {
		return -1
	}
	currentDepth := rightChild - leftChild
	if math.Abs(float64(currentDepth)) > 1 {
		return -1
	}
	return int(math.Max(float64(leftChild), float64(rightChild))) + 1
}

// 2*i + 1  right child
// 2*i + 2  left child
// 0 1 2 3 4 5
// A B C D E F
//     A
//   B   C
// D E  F -
func NewFullTree(array []string) Tree {
	root := NewNode(array[0])
	tree := newTree(root)
	recFullTree(0, root, array)
	return *tree

}

func recFullTree(ind int, node *Node, array []string) {
	if 2*ind+1 > len(array)-1 {
		return
	} else {
		if 2*ind+2 > len(array)-1 {
			leftChild := NewNode(array[2*ind+1])
			node.leftChild = leftChild
		} else {
			leftChild := NewNode(array[2*ind+1])
			node.leftChild = leftChild
			recFullTree(2*ind+1, leftChild, array)
			rightChild := NewNode(array[2*ind+2])
			node.rightChild = rightChild
			recFullTree(2*ind+2, rightChild, array)
		}

	}
}

func (t *Tree) DisplayTree() {

}
