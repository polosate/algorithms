package _6_heap

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	unsorted := []int64{9, 1, 2, 4, 10, 7, 8, 5, 3, 6}
	heapSort(unsorted)
	expected := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(unsorted, expected) {
		t.Error("want", expected, "have", unsorted)
	}

	unsorted = []int64{17, 10, 99, 77, 55}
	heapSort(unsorted)
	expected = []int64{10, 17, 55, 77, 99}
	if !reflect.DeepEqual(unsorted, expected) {
		t.Error("want", expected, "have", unsorted)
	}
}
