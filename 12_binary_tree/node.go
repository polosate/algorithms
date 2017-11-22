package _2_binary_tree

import "fmt"

type NodeValue struct {
	iData int64
}

type Node struct {
	value      *NodeValue
	leftChild  *Node
	rightChild *Node
}

func NewNodeValue(iData int64) *NodeValue {
	return &NodeValue{iData: iData}
}

func NewNode(value int64) *Node {
	nodeValue := NewNodeValue(value)
	return &Node{value: nodeValue}
}

func (n *Node) GetValue() int64 {
	return n.value.iData
}

func (n *Node) DisplayNode() {
	fmt.Print(n.value.iData, "; ")
}
