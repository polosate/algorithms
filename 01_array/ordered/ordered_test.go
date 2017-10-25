package ordered

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	array := newArray(10)

	// Insert elements
	array.Insert(22)
	array.Insert(77)
	array.Insert(44)
	array.Insert(33)
	array.Insert(66)
	array.Insert(55)
	array.Insert(11)
	array.Insert(111)
	array.Insert(15)
	array.Insert(34)

	expected := []int64{11, 15, 22, 33, 34, 44, 55, 66, 77, 111}
	res := array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected 01_array", expected, "have 01_array", res)
	}
	if array.GetLen() != 10 {
		t.Error("expected len", 10, "have 01_array", array.GetLen())
	}

	// Insert to full 01_array
	inserted := array.Insert(22)
	if inserted {
		t.Error("expected inserted", false, "have 01_array", inserted)
	}

	// BinarySearch element
	index := array.BinarySearch(55)
	if index != 6 {
		t.Error("expected index", 6, "have index", index)
	}

	// Delete nonexistent element
	deleted := array.Delete(1)
	if deleted {
		t.Error("expected deleted", false, "have deleted", deleted)
	}

	// Delete middle element
	deleted = array.Delete(44)
	if !deleted {
		t.Error("expected deleted", true, "have deleted", deleted)
	}
	expected = []int64{11, 15, 22, 33, 34, 55, 66, 77, 111}
	res = array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected 01_array", expected, "have 01_array", res)
	}
	if array.GetLen() != 9 {
		t.Error("expected len", 9, "have 01_array", array.GetLen())
	}

	// Delete first element
	deleted = array.Delete(11)
	if !deleted {
		t.Error("expected deleted", true, "have deleted", deleted)
	}
	expected = []int64{15, 22, 33, 34, 55, 66, 77, 111}
	res = array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected 01_array", expected, "have 01_array", res)
	}
	if array.GetLen() != 8 {
		t.Error("expected len", 8, "have 01_array", array.GetLen())
	}

	// Delete last element
	deleted = array.Delete(111)
	if !deleted {
		t.Error("expected deleted", true, "have deleted", deleted)
	}
	expected = []int64{15, 22, 33, 34, 55, 66, 77}
	res = array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected 01_array", expected, "have 01_array", res)
	}
	if array.GetLen() != 7 {
		t.Error("expected len", 7, "have 01_array", array.GetLen())
	}

	// Insert element again
	inserted = array.Insert(56)
	if !inserted {
		t.Error("expected inserted", true, "have 01_array", inserted)
	}

	expected = []int64{15, 22, 33, 34, 55, 56, 66, 77}
	res = array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected 01_array", expected, "have 01_array", res)
	}
	if array.GetLen() != 8 {
		t.Error("expected len", 8, "have 01_array", array.GetLen())
	}

	// Reverse
	expected = []int64{77, 66, 56, 55, 34, 33, 22, 15}
	array.Reverse()
	res = array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected 01_array", expected, "have 01_array", res)
	}
	deleted = array.Delete(55)
	expected = []int64{15, 22, 33, 34, 56, 66, 77}
	array.Reverse()
	res = array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected 01_array", expected, "have 01_array", res)
	}

}
