package _6_heap

func heapSort(unsorted []int64) {
	size := len(unsorted)
	heap := NewHeap(size)
	for i := 0; i < size; i++ {
		heap.insertAt(i, newNode(unsorted[i]))
		heap.incrementSize()
	}

	for i := size/2 - 1; i >= 0; i-- {
		heap.trickleDown(i)
	}

	for i := size - 1; i >= 0; i-- {
		biggestNode, _ := heap.remove()
		heap.incrementSize()
		heap.insertAt(i, biggestNode)
	}
}
