package unordered

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	array := newArray(10)

	// Insert elements
	array.Insert(10)
	array.Insert(55)
	array.Insert(7)
	array.Insert(18)
	array.Insert(89)
	array.Insert(33)
	array.Insert(66)
	array.Insert(11)
	array.Insert(84)
	array.Insert(99)

	expected := []int64{10, 55, 7, 18, 89, 33, 66, 11, 84, 99}
	res := array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected array", expected, "have array", res)
	}
	if array.GetLen() != 10 {
		t.Error("expected len", 10, "have array", array.GetLen())
	}

	// Insert to full array
	inserted := array.Insert(22)
	if inserted {
		t.Error("expected inserted", false, "have array", inserted)
	}

	// Find element
	index := array.Find(55)
	if index != 1 {
		t.Error("expected index", 1, "have index", index)
	}

	// Delete nonexistent element
	deleted := array.Delete(1)
	if deleted {
		t.Error("expected deleted", false, "have deleted", deleted)
	}

	// Delete middle element
	deleted = array.Delete(89)
	if !deleted {
		t.Error("expected deleted", true, "have deleted", deleted)
	}
	expected = []int64{10, 55, 7, 18, 33, 66, 11, 84, 99}
	res = array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected array", expected, "have array", res)
	}
	if array.GetLen() != 9 {
		t.Error("expected len", 9, "have array", array.GetLen())
	}

	// Delete first element
	deleted = array.Delete(10)
	if !deleted {
		t.Error("expected deleted", true, "have deleted", deleted)
	}
	expected = []int64{55, 7, 18, 33, 66, 11, 84, 99}
	res = array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected array", expected, "have array", res)
	}
	if array.GetLen() != 8 {
		t.Error("expected len", 8, "have array", array.GetLen())
	}

	// Delete last element
	deleted = array.Delete(99)
	if !deleted {
		t.Error("expected deleted", true, "have deleted", deleted)
	}
	expected = []int64{55, 7, 18, 33, 66, 11, 84}
	res = array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected array", expected, "have array", res)
	}
	if array.GetLen() != 7 {
		t.Error("expected len", 7, "have array", array.GetLen())
	}

	// Insert element again
	inserted = array.Insert(111)
	if !inserted {
		t.Error("expected inserted", true, "have array", inserted)
	}

	expected = []int64{55, 7, 18, 33, 66, 11, 84, 111}
	res = array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected array", expected, "have array", res)
	}
	if array.GetLen() != 8 {
		t.Error("expected len", 8, "have array", array.GetLen())
	}
}
