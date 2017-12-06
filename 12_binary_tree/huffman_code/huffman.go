package huffman_code

import (
	"bytes"
)

type hCode struct {
	hTree Tree
}

func New() hCode {
	return hCode{}
}

func (h *hCode) Encode(text string) string {
	return h.encode([]byte(text))
}

func (h *hCode) encode(text []byte) string {
	fMap := makeFrequencyMap(text)
	h.buildTree(fMap)
	codeMap := buildCodeMap(h.hTree)

	encoded := ""
	for _, v := range text {
		encoded += codeMap[v]
	}
	return encoded
}

func (h *hCode) decode(encoded string) string {
	decoded := []byte{}
	current := h.hTree.root
	for i := range encoded {
		if encoded[i] == byte('0') {
			current = current.leftChild
		} else {
			current = current.rightChild
		}
		if current.rightChild == nil && current.leftChild == nil {
			decoded = append(decoded, current.char)
			current = h.hTree.root
			continue
		}
	}
	return string(decoded)
}

func makeFrequencyMap(text []byte) map[byte]int {
	frequencyMap := map[byte]int{}
	for i := range text {
		currentChar := text[i]
		if frequencyMap[currentChar] > 0 {
			continue
		} else {
			frequencyMap[currentChar] = bytes.Count(text, []byte{currentChar})
		}
	}
	return frequencyMap
}

func (h *hCode) buildTree(frequencyMap map[byte]int) {
	list := NewList()
	for k, v := range frequencyMap {
		list.Insert(NewNode(k, v))
	}

	for list.first.next != nil {
		node1 := list.Remove()
		node2 := list.Remove()
		newNode := NewNode(0x0, node1.count+node2.count)
		if node1.count < node2.count {
			newNode.leftChild = node1
			newNode.rightChild = node2
		} else {
			newNode.leftChild = node2
			newNode.rightChild = node1
		}
		list.Insert(newNode)
	}
	tree := NewTree()
	tree.root = list.Remove()
	h.hTree = tree
}

func recWalk(node *Node, code string, codeMap map[byte]string) {
	if node.rightChild == nil && node.leftChild == nil {
		codeMap[node.char] = code
	}

	if node.leftChild != nil {
		recWalk(node.leftChild, code+"0", codeMap)
	}
	if node.rightChild != nil {
		recWalk(node.rightChild, code+"1", codeMap)
	}
}

func buildCodeMap(tree Tree) map[byte]string {
	codeMap := map[byte]string{}
	recWalk(tree.root, "", codeMap)
	return codeMap
}
