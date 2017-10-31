package single

import "testing"

func TestSortedList(t *testing.T) {
	list := NewSortedList()
	list.Insert(5, 5.5)
	list.Insert(2, 2.2)
	list.Insert(3, 3.3)
	list.DisplayList()

	el, _ := list.Remove()
	if el.dData != 2.2 {
		t.Error("Expected link.dData", 2.2, "have", el)
	}
	list.DisplayList()
}
