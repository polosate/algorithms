package huffman_code

import (
	"fmt"
	"sort"
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

type huffmanCode struct {
	text         string
	frequencyMap map[string]int
	codeMap      map[string]byte
}

func NewHuffmanCode(text string) huffmanCode {
	return huffmanCode{
		text:         text,
		frequencyMap: map[string]int{},
	}
}

func (h *huffmanCode) makeFrequencyMap() {
	byteArr := []byte(h.text)
	for i := range byteArr {
		if h.frequencyMap[string(byteArr[i])] > 0 {
			continue
		} else {
			h.frequencyMap[string(byteArr[i])] = strings.Count(h.text, string(byteArr[i]))
		}
	}
	var valueArray []int
	for _, v := range h.frequencyMap {
		valueArray = append(valueArray, v)
	}
	sort.Ints(valueArray)
	fmt.Println(valueArray)
}
