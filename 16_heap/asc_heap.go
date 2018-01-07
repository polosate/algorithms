package _6_heap

import (
	"errors"
	"fmt"
)

type ascHeap struct {
	curSize   int
	size      int
	heapArray []*node
}

func NewAscHeap(maxSize int) ascHeap {
	return ascHeap{
		curSize:   0,
		size:      maxSize,
		heapArray: make([]*node, maxSize),
	}
}

func (h *ascHeap) isEmpty() bool {
	return h.curSize == 0
}

func (h *ascHeap) insert(key int64) (bool, error) {
	if h.curSize == h.size {
		return false, errors.New("heap is full")
	}
	newNode := newNode(key)
	h.heapArray[h.curSize] = newNode

	h.trickleUp(h.curSize)
	h.curSize++
	return true, nil
}

func (h *ascHeap) remove() (*node, error) {
	if h.isEmpty() {
		return nil, errors.New("heap is empty")
	}
	root := h.heapArray[0]
	h.curSize--
	h.heapArray[0] = h.heapArray[h.curSize]
	h.trickleDown(0)
	return root, nil
}

func (h *ascHeap) trickleUp(index int) {
	parent := (index - 1) / 2
	bottom := h.heapArray[index]
	for index > 0 && bottom.getKey() < h.heapArray[parent].getKey() {
		h.heapArray[index] = h.heapArray[parent]
		index = parent
		parent = (parent - 1) / 2
	}
	h.heapArray[index] = bottom
}

func (h *ascHeap) trickleDown(index int) {
	var smallerChild, leftChild, rightChild int
	top := h.heapArray[index]
	for index < h.curSize/2 {
		leftChild = 2*index + 1
		rightChild = leftChild + 1
		if rightChild < h.curSize && h.heapArray[rightChild].getKey() < h.heapArray[leftChild].getKey() {
			smallerChild = rightChild
		} else {
			smallerChild = leftChild
		}
		if top.getKey() < h.heapArray[smallerChild].getKey() {
			break
		}
		h.heapArray[index] = h.heapArray[smallerChild]
		index = smallerChild
	}
	h.heapArray[index] = top
}

func (h *ascHeap) displayHeap() {
	for i := 0; i < h.curSize; i++ {
		fmt.Print(h.heapArray[i].getKey(), " ")
	}
	fmt.Println()
}
