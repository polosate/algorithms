package _6_heap

type pqHeap struct {
	h heap
}

func newPQHeap(size int) IPriorityQueue {
	return &pqHeap{
		h: NewHeap(size),
	}
}

func (pq *pqHeap) isEmpty() bool {
	return pq.h.isEmpty()
}

func (pq *pqHeap) isFull() bool {
	return pq.h.isFull()
}

func (pq *pqHeap) insert(key int64) {
	if pq.isFull() {
		return
	}
	pq.h.insert(key)
}

func (pq *pqHeap) remove() int64 {
	if pq.isEmpty() {
		return -1
	}
	n, err := pq.h.remove()
	if err != nil {
		return -1
	}
	return n.getKey()
}

func (pq *pqHeap) peek() int64 {
	if pq.isEmpty() {
		return -1
	}
	return pq.h.heapArray[0].getKey()
}
