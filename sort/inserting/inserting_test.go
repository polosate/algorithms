package inserting

import (
	"reflect"
	"testing"
)

func TestUnarySearch(t *testing.T) {
	array := newArray(10)

	// Insert elements
	// 22 77 44 33 66 55 11 111 1 34
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

	array.InsertingSort()

	expected = []int64{11, 15, 22, 33, 34, 44, 55, 66, 77, 111}
	res = array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected array", expected, "have array", res)
	}
}

//func Test_BinarySearch(t *testing.T) {
//	array := newArray(10)
//
//	// Insert elements
//	// 22 77 44 33 66 55 11 111 1 34
//	array.Insert(22)
//	array.Insert(77)
//	array.Insert(44)
//	array.Insert(33)
//	array.Insert(66)
//	array.Insert(55)
//	array.Insert(11)
//	array.Insert(111)
//	array.Insert(15)
//	array.Insert(34)
//
//	expected := []int64{22, 77, 44, 33, 66, 55, 11, 111, 15, 34}
//	res := array.GetData()
//	if !reflect.DeepEqual(res, expected) {
//		t.Error("expected array", expected, "have array", res)
//	}
//
//	array.InsertingSortWithBinarySearch()
//
//	expected = []int64{11, 15, 22, 33, 34, 44, 55, 66, 77, 111}
//	res = array.GetData()
//	if !reflect.DeepEqual(res, expected) {
//		t.Error("expected array", expected, "have array", res)
//	}
//}

func TestMedianEven(t *testing.T) {
	array := newArray(4)

	array.Insert(1)
	array.Insert(2)
	array.Insert(3)
	array.Insert(4)

	m := array.Median()
	if !reflect.DeepEqual(m, float64(2.5)) {
		t.Error("expected array", float64(2.5), "have array", m)
	}
}

func TestMedianOdd(t *testing.T) {
	array := newArray(5)

	array.Insert(1)
	array.Insert(2)
	array.Insert(3)
	array.Insert(4)
	array.Insert(5)

	m := array.Median()
	if !reflect.DeepEqual(m, float64(3)) {
		t.Error("expected array", float64(3), "have array", m)
	}
}

func TestNoDups(t *testing.T) {
	array := newArray(10)

	array.Insert(1)
	array.Insert(2)
	array.Insert(6)
	array.Insert(1)
	array.Insert(1)
	array.Insert(4)
	array.Insert(3)
	array.Insert(5)
	array.Insert(5)
	array.Insert(6)

	array.NoDups()
	res := array.GetData()
	expected := []int64{1, 2, 3, 4, 5, 6}
	res = array.GetData()
	if !reflect.DeepEqual(res, expected) {
		t.Error("expected array", expected, "have array", res)
	}
}
