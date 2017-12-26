package exercises

import "testing"

func TestPartition(t *testing.T) {
	list := NewSingleList()
	list.InsertFirst(2)
	list.InsertFirst(1)
	list.InsertFirst(3)
	list.InsertFirst(33)
	list.InsertFirst(11)
	list.InsertFirst(5)
	list.InsertFirst(9)
	list.InsertFirst(66)
	list.InsertFirst(12)
	list.InsertFirst(15)
	list.InsertFirst(6)
	l := partition(list, 10)

	expected := []int64{2, 1, 3, 5, 9, 6, 33, 11, 66, 12, 15}
	iter := l.GetIterator()
	i := 0
	for iter.GetCurrent() != nil {
		if iter.GetCurrent().value != expected[i] {
			t.FailNow()
		}
		i++
		iter.NextLink()
	}
}
