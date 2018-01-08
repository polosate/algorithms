package _6_heap

type priorityQ struct {
	h heap
}

func newPriorityQ(size int) priorityQ {
	return priorityQ{
		h: NewHeap(size),
	}
}

func (pq *priorityQ) isEmpty() bool {
	return pq.h.isEmpty()
}

func (pq *priorityQ) isFull() bool {
	return pq.h.isFull()
}

func (pq *priorityQ) insert(key int64) {
	if pq.isFull() {
		return
	}
	pq.h.insert(key)
}

func (pq *priorityQ) remove() int64 {
	if pq.isEmpty() {
		return -1
	}
	n, err := pq.h.remove()
	if err != nil {
		return -1
	}
	return n.getKey()
}

func (pq *priorityQ) peek() int64 {
	if pq.isEmpty() {
		return -1
	}
	return pq.h.heapArray[0].getKey()
}
