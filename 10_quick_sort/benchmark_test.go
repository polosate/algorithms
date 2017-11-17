package _0_quick_sort

import (
	"testing"
)

func BenchmarkQuickSort1(b *testing.B) {
	var res []int64
	for i := 50000; i > 0; i-- {
		res = append(res, int64(i))
	}
	quick1 := NewQuickSortArray(res)
	quick1.QuickSort()
}

func BenchmarkQuickSort2(b *testing.B) {
	var res []int64
	for i := 50000; i > 0; i-- {
		res = append(res, int64(i))
	}
	quick2 := NewQuickSortArray2(res)
	quick2.QuickSort()
}
