package _6_heap

import (
	"errors"
	"fmt"
)

type heap struct {
	curSize   int
	size      int
	heapArray []*node
}

func NewHeap(maxSize int) heap {
	return heap{
		curSize:   0,
		size:      maxSize,
		heapArray: make([]*node, maxSize),
	}
}

func (h *heap) isEmpty() bool {
	return h.curSize == 0
}

func (h *heap) isFull() bool {
	return h.curSize == h.size
}

func (h *heap) insert(key int64) (bool, error) {
	if h.curSize == h.size {
		return false, errors.New("heap is full")
	}
	newNode := newNode(key)
	h.heapArray[h.curSize] = newNode

	h.trickleUp(h.curSize)
	h.curSize++
	return true, nil
}

func (h *heap) remove() (*node, error) {
	if h.isEmpty() {
		return nil, errors.New("heap is empty")
	}
	root := h.heapArray[0]
	h.curSize--
	h.heapArray[0] = h.heapArray[h.curSize]
	h.trickleDown(0)
	return root, nil
}

func (h *heap) trickleUp(index int) {
	parent := (index - 1) / 2
	bottom := h.heapArray[index]
	for index > 0 && bottom.getKey() > h.heapArray[parent].getKey() {
		h.heapArray[index] = h.heapArray[parent]
		index = parent
		parent = (parent - 1) / 2
	}
	h.heapArray[index] = bottom
}

func (h *heap) trickleDown(index int) {
	var lagrerChild, leftChild, rightChild int
	top := h.heapArray[index]
	for index < h.curSize/2 {
		leftChild = 2*index + 1
		rightChild = leftChild + 1
		if rightChild < h.curSize && h.heapArray[leftChild].getKey() < h.heapArray[rightChild].getKey() {
			lagrerChild = rightChild
		} else {
			lagrerChild = leftChild
		}
		if top.getKey() > h.heapArray[lagrerChild].getKey() {
			break
		}
		h.heapArray[index] = h.heapArray[lagrerChild]
		index = lagrerChild
	}
	h.heapArray[index] = top
}

func (h *heap) toss(key int64) bool {
	if h.curSize == h.size {
		return false
	}
	newNode := newNode(key)
	h.heapArray[h.curSize] = newNode
	h.curSize++
	return true
}

func (h *heap) restoreHeap() {
	for i := h.curSize/2 - 1; i >= 0; i-- {
		h.trickleDown(i)
	}
}

func (h *heap) displayHeap() {
	for i := 0; i < h.curSize; i++ {
		fmt.Print(h.heapArray[i].getKey(), " ")
	}
	fmt.Println()
}
