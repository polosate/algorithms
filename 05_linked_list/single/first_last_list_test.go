package single

import "testing"

func TestFirstLastList(t *testing.T) {
	ll := NewFirstLastList()
	el := ll.DeleteFirst()
	if el != nil {
		t.Error("Expected element", nil, "have", el)
	}

	ll.InsertFirst(1, 1.1)
	ll.InsertLast(3, 3.3)
	ll.InsertFirst(2, 2.2)
	ll.InsertLast(4, 4.4)
	ll.DisplayList()

	el = ll.DeleteFirst()
	if el.iData != 2 {
		t.Error("Expected element iData", 2, "have", el.iData)
	}
	el = ll.DeleteFirst()
	if el.iData != 1 {
		t.Error("Expected element iData", 1, "have", el.iData)
	}
	ll.DisplayList()

}
