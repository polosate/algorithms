package _0_quick_sort

import (
	"testing"
)

var (
	quick1 QuickSortArray
	quick2 QuickSortArray2
)

func TestMain(m *testing.M) {
	var res []int64
	for i := 50000; i > 0; i-- {
		res = append(res, int64(i))
	}
	quick1 = NewQuickSortArray(res)
	quick2 = NewQuickSortArray2(res)
	m.Run()
}

func BenchmarkQuickSort1(b *testing.B) {
	quick1.QuickSort()
}

func BenchmarkQuickSort2(b *testing.B) {
	quick2.QuickSort()
}
