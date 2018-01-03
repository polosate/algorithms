package _6_heap

type heapArray []int64

func heapSort(unsorted heapArray) {
	size := len(unsorted)
	for i := size/2 - 1; i >= 0; i-- {
		unsorted.trickleDown(i, size)
	}
	for i := size - 1; i > 0; i-- {
		biggestNode := unsorted.remove(i)
		unsorted.insertAt(i, biggestNode)
	}
}

func (ha heapArray) trickleDown(index int, curSize int) {
	var largerChild, leftChild, rightChild int
	top := ha[index]
	for index < (curSize)/2 {
		leftChild = 2*index + 1
		rightChild = leftChild + 1
		if rightChild < curSize && ha[leftChild] < ha[rightChild] {
			largerChild = rightChild
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
	ha.trickleDown(0, curSize)
	return biggest
}

func (ha heapArray) insertAt(i int, value int64) {
	ha[i] = value
}
