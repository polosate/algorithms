package sort_benchmarks

import (
	"math/rand"
	"testing"

	"algorithms/02_compare_sort/bubble"
	insert "algorithms/02_compare_sort/inserting"
	sel "algorithms/02_compare_sort/selecting"
	recursion "algorithms/07_recursion"
	shell "algorithms/08_shell_sort"
	quick "algorithms/10_quick_sort"
)

const (
	smallSize = 50000
	bigSize   = 5000000
)

var (
	mergeA  recursion.MArray
	bubbleA bubble.BArray
	insertA insert.IArray
	selectA sel.SArray
	shellA  shell.ShellArray
	quickA  quick.QuickSortArray
)

func BenchmarkQuickSort(b *testing.B) {
	var res []int64
	for i := 0; i < bigSize; i++ {
		res = append(res, int64(rand.Int()))
	}
	quickA = quick.NewQuickSortArray(res)
	quickA.QuickSort()
}

func BenchmarkShellSort(b *testing.B) {
	var res []int64
	for i := 0; i < bigSize; i++ {
		res = append(res, int64(rand.Int()))
	}
	shellA = shell.NewShellArray(res)
	shellA.ShellSort()
}

func BenchmarkMergeSort(b *testing.B) {
	var res []int64
	for i := 0; i < bigSize; i++ {
		res = append(res, int64(rand.Int()))
	}
	mergeA = recursion.NewMArray(res)
	mergeA.MergeSort()
}

func BenchmarkInsertingSort(b *testing.B) {
	var res []int64
	for i := 0; i < smallSize; i++ {
		res = append(res, int64(rand.Int()))
	}
	insertA = insert.NewIArray(res)
	insertA.InsertingSort()
}

func BenchmarkSelectionSort(b *testing.B) {
	var res []int64
	for i := 0; i < smallSize; i++ {
		res = append(res, int64(rand.Int()))
	}
	selectA = sel.NewSArray(res)
	selectA.SelectingSort()
}

func BenchmarkBubbleSort(b *testing.B) {
	var res []int64
	for i := 0; i < smallSize; i++ {
		res = append(res, int64(rand.Int()))
	}
	bubbleA = bubble.NewBArray(res)
	bubbleA.BubbleSort()
}
