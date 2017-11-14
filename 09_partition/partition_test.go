package _9_partition

import (
	"testing"
)

func TestPartition(t *testing.T) {
	data := []int64{5, 2, 15, 20, 9, 10, 3, 6, 4, 0}
	a := NewPartArray(data)
	pivotInd := a.Partition(0, len(data)-1, 6)
	if pivotInd != 5 {
		t.Error("Want partition index", 5, "have", pivotInd)
	}
	for i, v := range a.array {
		if i <= 5 && v > 6 {
			t.Fail()
		}
		if i > 5 && v <= 6 {
			t.Fail()
		}
	}
}
