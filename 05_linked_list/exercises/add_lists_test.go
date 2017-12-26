package exercises

import "testing"

func TestAddReversedLists(t *testing.T) {
	l1 := NewSingleList()
	l1.InsertFirst(3)
	l1.InsertFirst(2)
	l1.InsertFirst(1)

	l2 := NewSingleList()
	l2.InsertFirst(9)
	l2.InsertFirst(1)

	res := addLists(l1, l2)

	expected := []int64{2, 1, 4}
	i := 0
	cur := res.first
	for cur != nil {
		if cur.value != expected[i] {
			t.Error("Want", expected[i], "have", cur.value)
		}
		i++
		cur = cur.next
	}
}

func TestAddDirectLists(t *testing.T) {
	l1 := NewSingleList()
	l1.InsertFirst(3)
	l1.InsertFirst(2)
	l1.InsertFirst(1)

	l2 := NewSingleList()
	l2.InsertFirst(9)
	l2.InsertFirst(1)

	res := addDirectLists(l1, l2)

	expected := []int64{1, 4, 2}
	i := 0
	cur := res.first
	for cur != nil {
		if cur.value != expected[i] {
			t.Error("Want", expected[i], "have", cur.value)
		}
		i++
		cur = cur.next
	}
}
