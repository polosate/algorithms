package exercises

type link struct {
	value *Node
	next  *link
}

type list struct {
	first *link
}

type stack struct {
	list
}

type Node struct {
	value      string
	leftChild  *Node
	rightChild *Node
}
type Tree struct {
	root *Node
}

func NewTree() Tree {
	return Tree{}
}

func (t *Tree) build(array []string) {

}
