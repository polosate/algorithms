package _6_heap

import "fmt"

type heapArray []int64

func heapSort(unsorted heapArray) {
	size := len(unsorted)
	for i := size/2 - 1; i >= 0; i-- {
		unsorted.trickleDown(i, size)
	}
	for i := size - 1; i >= 0; i-- {
		biggestNode := unsorted.remove(i)
		unsorted.insertAt(i, biggestNode)
	}
	fmt.Println(unsorted)
}

func (ha heapArray) trickleDown(index int, curSize int) {
	var largerChild, leftChild, rigrhChild int
	top := ha[index]
	for index < curSize/2 {
		leftChild = 2*index + 1
		rigrhChild = leftChild + 1
		if rigrhChild < curSize && ha[leftChild] < ha[rigrhChild] {
			largerChild = rigrhChild
		} else {
			largerChild = leftChild
		}
		if top > ha[largerChild] {
			break
		}
		ha[index] = ha[largerChild]
		index = largerChild
	}
	ha[index] = top
}

func (ha heapArray) remove(curSize int) int64 {
	biggest := ha[0]
	ha[0] = ha[curSize]
	ha.trickleDown(0, curSize-1)
	return biggest
}

func (ha heapArray) insertAt(i int, value int64) {
	ha[i] = value
}
