package huffman_code

import (
	"fmt"
	"strings"
)

type Node struct {
	value      string
	leftChild  *Node
	rightChild *Node
}

func NewNode(value string) *Node {
	return &Node{value: value}
}

func (n *Node) GetValue() string {
	return n.value
}

func (n *Node) DisplayNode() {
	fmt.Print(n.value, "; ")
}

type Tree struct {
	root *Node
}

func NewTree() Tree {
	return Tree{root: nil}
}

type huffmanTree struct {
	text         string
	frequencyMap map[int]string
	codeMap      map[byte]string
}

func NewHuffmanTree() huffmanTree {
	return huffmanTree{}
}
