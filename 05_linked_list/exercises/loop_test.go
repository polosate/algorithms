package exercises

import (
	"testing"
)

func TestLoop(t *testing.T) {
	list := NewSingleList()
	list.InsertFirst(1)
	list.InsertFirst(2)
	list.InsertFirst(3)
	list.InsertFirst(4)
	list.InsertFirst(5)
	list.InsertFirst(6)
	list.InsertFirst(7)
	list.InsertFirst(8)
	list.InsertFirst(9)
	list.InsertFirst(10)

	iter1 := list.GetIterator()
	iter2 := list.GetIterator()

	for iter1.GetCurrent().next != nil {
		iter1.NextLink()
	}
	for i := 0; i < 5; i++ {
		iter2.NextLink()
	}
	iter1.GetCurrent().next = iter2.GetCurrent()

	loopStart := findLoopStart(list)
	if loopStart.value != 5 {
		t.Error("Want", 5, "have", loopStart.value)
	}
	//// Trololo -)
	//list.DisplayList()
}
