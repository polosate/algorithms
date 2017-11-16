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

var (
	mergeA  recursion.MArray
	bubbleA bubble.BArray
	insertA insert.IArray
	selectA sel.SArray
	shellA  shell.ShellArray
	quickA  quick.QuickSortArray
)

func TestMain(m *testing.M) {
	var res []int64
	for i := 0; i < 500000; i++ {
		res = append(res, int64(rand.Int()))
	}
	mergeA = recursion.NewMArray(res)
	bubbleA = bubble.NewBArray(res)
	insertA = insert.NewIArray(res)
	selectA = sel.NewSArray(res)
	shellA = shell.NewShellArray(res)
	quickA = quick.NewQuickSortArray(res)
	m.Run()
}

func BenchmarkQuickSort(b *testing.B) {
	quickA.QuickSort()
}

func BenchmarkShellSort(b *testing.B) {
	shellA.ShellSort()
}

func BenchmarkMergeSort(b *testing.B) {
	mergeA.MergeSort()
}

func BenchmarkInsertingSort(b *testing.B) {
	insertA.InsertingSort()
}

func BenchmarkSelectionSort(b *testing.B) {
	selectA.SelectingSort()
}

func BenchmarkBubbleSort(b *testing.B) {
	bubbleA.BubbleSort()
}
