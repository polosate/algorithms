package bubble

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

	expected := []int64{22, 77, 44, 33, 66, 55, 11, 111, 15, 34}
	res := array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected array", expected, "have array", res)
	}

	array.BubbleSort()

	expected = []int64{11, 15, 22, 33, 34, 44, 55, 66, 77, 111}
	res = array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected array", expected, "have array", res)
	}
}
