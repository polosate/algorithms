package exercises

import "testing"

func TestPartition(t *testing.T) {
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
}
