package _0_quick_sort

import (
	"reflect"
	"testing"
)

func TestNewQuickSort(t *testing.T) {
	data := []int64{5, 2, 15, 20, 9, 10, 3, 6, 40, 30, 1, 11, 7}
	array := NewQuickSortArray(data)
	array.QuickSort()

	expected := []int64{1, 2, 3, 5, 6, 7, 9, 10, 11, 15, 20, 30, 40}
	if !reflect.DeepEqual(array.array, expected) {
		t.Error("want", expected, "have", array.array)
	}
}

func TestNewQuickSortReverse(t *testing.T) {
	data := []int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	array := NewQuickSortArray(data)
	array.QuickSort()

	expected := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(array.array, expected) {
		t.Error("want", expected, "have", array.array)
	}
}
