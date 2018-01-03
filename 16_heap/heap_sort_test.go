package _6_heap

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	unsorted := []int64{9, 1, 2, 4, 10, 7, 8, 5, 3, 6}
	heapSort(unsorted)

}
