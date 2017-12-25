package _6_heap

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	unsorted := []int64{81, 6, 23, 38, 95, 71, 72, 39, 34, 53}
	heapSort(unsorted)
}
