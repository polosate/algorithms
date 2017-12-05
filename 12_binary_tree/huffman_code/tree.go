package huffman_code

type Node struct {
	char       string
	count      int
	leftChild  *Node
	rightChild *Node
}

func NewNode(char string, count int) *Node {
	return &Node{char: char, count: count}
}

//func (n *Node) GetValue() string {
//	return n.value
//}
//
//func (n *Node) DisplayNode() {
//	fmt.Print(n.value, "; ")
//}

type Tree struct {
	root *Node
}

func NewTree() Tree {
	return Tree{root: nil}
}
