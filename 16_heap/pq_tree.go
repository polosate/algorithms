package _6_heap

type pqTree struct {
	t tree
}

func newPQTree(size int) IPriorityQueue {
	return &pqHeap{
		h: NewHeap(size),
	}
}

func (pq *pqTree) isEmpty() bool {
	return pq.t.isEmpty()
}

func (pq *pqTree) isFull() bool {
	return false
}

func (pq *pqTree) insert(key int64) {
	pq.t.insert(key)
}

func (pq *pqTree) remove() int64 {
	return pq.t.removeMax()
}

func (pq *pqTree) peek() int64 {
	if pq.isEmpty() {
		return -1
	}
	return pq.t.findMax()
}
