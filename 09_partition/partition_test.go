package _9_partition

import (
	"testing"
)

func TestPartition(t *testing.T) {
	data := []int64{5, 2, 15, 20, 9, 10, 3, 6, 40, 30}
	var pivot int64 = 7
	a := NewPartArray(data)
	pivotInd := a.Partition(0, len(data)-1, pivot)
	if pivotInd != 4 {
		t.Error("Want partition index", 4, "have", pivotInd)
	}
	for i, v := range a.array {
		if i < pivotInd && v > pivot {
			t.Error("want element less then", pivot, "got", v)
		}
		if i >= pivotInd && v < pivot {
			t.Error("want element more then", pivot, "got", v)
		}
	}
}
