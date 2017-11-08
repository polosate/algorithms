package _7_recursion

import "testing"

func TestBinarySearchOdd(t *testing.T) {
	arr := NewArray([]int{1, 2, 3, 4, 5})

	index := arr.Find(1)
	if index != 0 {
		t.Error("Wand index", 0, "have", index)
	}
	index = arr.Find(3)
	if index != 2 {
		t.Error("Wand index", 2, "have", index)
	}
	index = arr.Find(5)
	if index != 4 {
		t.Error("Wand index", 4, "have", index)
	}
	index = arr.Find(6)
	if index != 5 {
		t.Error("Wand index", 5, "have", index)
	}
}

func TestBinarySearchEven(t *testing.T) {
	array := NewArray([]int{1, 2, 3, 4})

	index := array.Find(1)
	if index != 0 {
		t.Error("Wand index", 0, "have", index)
	}
	index = array.Find(3)
	if index != 2 {
		t.Error("Wand index", 2, "have", index)
	}
	index = array.Find(4)
	if index != 3 {
		t.Error("Wand index", 3, "have", index)
	}
	index = array.Find(6)
	if index != 4 {
		t.Error("Wand index", 4, "have", index)
	}
}
