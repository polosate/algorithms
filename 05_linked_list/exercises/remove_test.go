package exercises

import (
	"testing"
)

func TestRemove(t *testing.T) {
	list := NewSingleList()
	list.InsertFirst(1)
	list.InsertFirst(2)
	list.InsertFirst(3)
	list.InsertFirst(5)
	list.InsertFirst(8)
	list.InsertFirst(4)

	iter1 := list.GetIterator()

	for i := 0; i < 1; i++ {
		iter1.NextLink()
	}

	Remove(iter1.GetCurrent())

	iter2 := list.GetIterator()
	for i := 0; i < 1; i++ {
		iter2.NextLink()
	}
	if iter2.GetCurrent().value != 5 {
		t.Error("Want", 5, "have", iter2.GetCurrent().value)
	}
}
