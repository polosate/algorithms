package exercises

import "fmt"

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
}

func NewList() list {
	return list{}
}

func (l *list) IsEmpty() bool {
	return l.first == nil
}

func (l *list) InsertFirst(tree Tree) {
	newLink := NewLink(tree)
	newLink.next = l.first
	l.first = newLink
}

func (l *list) RemoveFirst() Tree {
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
		fmt.Println("+++++++++++++++++++")
		current := newList.first
		for current.next != nil {
			fmt.Println("======")
			fmt.Println(current.tree.root)
			l1 := newList.RemoveFirst()
			l2 := newList.RemoveFirst()
			current = newList.first
			resTree := newTree(NewNode("+"))
			resTree.root.leftChild = l1.root
			resTree.root.rightChild = l2.root
			newList.InsertFirst(*resTree)

			//current = current.next
			//fmt.Println(current.tree.root)
			fmt.Println("======")
		}
		fmt.Println("+++++++++++++++++++")
	}
	return newList.RemoveFirst()
}

func (t *Tree) DisplayTree() {

}
