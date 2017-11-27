package exercises

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
	leftChild  *Tree
	rightChild *Tree
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
		resultTree.root.leftChild = &leaf1
		resultTree.root.rightChild = &leaf2
		newList.InsertFirst(*resultTree)
	}
	return newList.RemoveFirst()
}

func (t *Tree) DisplayTree() {

}
