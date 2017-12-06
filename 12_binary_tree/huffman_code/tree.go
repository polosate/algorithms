package huffman_code

type Node struct {
	char       byte
	count      int
	leftChild  *Node
	rightChild *Node
}

func NewNode(char byte, count int) *Node {
	return &Node{char: char, count: count}
}

type Tree struct {
	root *Node
}

func NewTree() Tree {
	return Tree{root: nil}
}
