package _0_quick_sort

import (
	"reflect"
	"testing"
)

func TestNewQuickSort02(t *testing.T) {
	data := []int64{5, 2, 15, 20, 9, 10, 3, 6, 40, 30, 1, 11, 7}
	array := NewQuickSortArray2(data)
	array.QuickSort()
	expected := []int64{1, 2, 3, 5, 6, 7, 9, 10, 11, 15, 20, 30, 40}
	if !reflect.DeepEqual(array.array, expected) {
		t.Error("want", expected, "have", array.array)
	}
}

func TestManualSort(t *testing.T) {
	seed := []int64{1, 2, 3}
	combi := doCombinations(seed, []int64{})

	for _, v := range combi {
		array := NewQuickSortArray2(v)
		array.manualSort(0, 2)
		expected := []int64{1, 2, 3}
		if !reflect.DeepEqual(expected, array.array) {
			t.Error("want", expected, "have", array.array)
		}
	}

	seed = []int64{1, 2}
	combi = doCombinations(seed, []int64{})

	for _, v := range combi {
		array := NewQuickSortArray2(v)
		array.manualSort(0, 1)
		expected := []int64{1, 2}
		if !reflect.DeepEqual(expected, array.array) {
			t.Error("want", expected, "have", array.array)
		}
	}
}

func doCombinations(array []int64, prefix []int64) [][]int64 {
	if len(array) == 0 {
		res := [][]int64{}
		res = append(res, prefix)
		return res
	}
	var res [][]int64
	for i := 0; i < len(array); i++ {
		r := doCombinations(array[1:], append(prefix, array[0]))
		res = append(res, r...)
		shift(array)
	}
	return res
}

func shift(array []int64) {
	temp := array[len(array)-1]
	for i := len(array) - 2; i >= 0; i-- {
		array[i+1] = array[i]
	}
	array[0] = temp
}
