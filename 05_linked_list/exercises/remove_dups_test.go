package exercises

import (
	"testing"
)

func TestRemoveDups(t *testing.T) {
	list := NewSingleList()
	list.InsertFirst(1)
	list.InsertFirst(2)
	list.InsertFirst(3)
	list.InsertFirst(3)
	list.InsertFirst(1)
	list.InsertFirst(5)
	list.InsertFirst(1)
	list.InsertFirst(6)
	list.InsertFirst(2)
	list.InsertFirst(2)

	result := []int64{2, 6, 1, 5, 3}

	RemoveDups(list)

	current := list.first
	for _, v := range result {
		if current.value != v {
			t.Error("Want", v, "got", current.value)
		}
		current = current.next
	}
}
