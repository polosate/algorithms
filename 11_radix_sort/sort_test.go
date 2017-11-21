package _1_radix_sort

import (
	"reflect"
	"testing"
)

func TestRadixSort(t *testing.T) {
	data := []int64{435, 123, 2, 45, 67, 911, 234}
	array := newRadixArray(data)
	array.Sort()
	expected := []int64{2, 45, 67, 123, 234, 435, 911}
	if !reflect.DeepEqual(array.array, expected) {
		t.Error("Want array", expected, "have", array.array)
	}
}
