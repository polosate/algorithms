package sort_benchmarks

import (
	"math/rand"
	"testing"

	bubble "algorithms/02_compare_sort/bubble"
	insert "algorithms/02_compare_sort/inserting"
	sel "algorithms/02_compare_sort/selecting"
	recursion "algorithms/07_recursion"
	shell "algorithms/08_shell_sort"
)

var (
	mergeA  recursion.MArray
	bubbleA bubble.BArray
	insertA insert.IArray
	selectA sel.SArray
	shellA  shell.ShellArray
)

func TestMain(m *testing.M) {
	res := []int64{}
	for i := 0; i < 500000	; i++ {
		res = append(res, int64(rand.Int()))
	}
	mergeA = recursion.NewMArray(res)
	bubbleA = bubble.NewBArray(res)
	insertA = insert.NewIArray(res)
	selectA = sel.NewSArray(res)
	shellA = shell.NewShellArray(res)
	m.Run()
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
