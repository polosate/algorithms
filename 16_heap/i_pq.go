package _6_heap

type IPriorityQueue interface {
	isEmpty() bool
	isFull() bool
	insert(key int64)
	remove() int64
	peek() int64
}
