package double

import "testing"

func TestInsertFirst(t *testing.T) {
	ll := NewDoubleLinkedList()
	ll.InsertFirst(1)
	ll.InsertFirst(2)
	ll.DisplayForward()
	ll.DisplayBackward()
}

func TestInsertLast(t *testing.T) {
	ll := NewDoubleLinkedList()
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.DisplayForward()
	ll.DisplayBackward()
}

func TestInsertAfter(t *testing.T) {
	ll := NewDoubleLinkedList()
	ll.InsertLast(3.3)
	ll.InsertLast(1.1)
	ll.InsertLast(4.4)
	ll.InsertAfter(3.3, 2.2)
	ll.DisplayForward()
}

func TestDeleteFirst(t *testing.T) {
	ll := NewDoubleLinkedList()
	el := ll.DeleteFirst()
	if el != nil {
		t.Error("Expected element", nil, "have", el)
	}
	ll.InsertFirst(1)
	el = ll.DeleteFirst()
	if el.dData != 1 {
		t.Error("Expected element", 1, "have", el.dData)
	}
	ll.InsertFirst(2)
	ll.InsertFirst(3)
	ll.InsertFirst(4)
	el = ll.DeleteFirst()
	if el.dData != 4 {
		t.Error("Expected element", 4, "have", el.dData)
	}
	ll.DisplayForward()
	ll.DisplayBackward()
}

func TestDeleteLast(t *testing.T) {
	ll := NewDoubleLinkedList()
	el := ll.DeleteLast()
	if el != nil {
		t.Error("Expected element", nil, "have", el)
	}
	ll.InsertFirst(1)
	el = ll.DeleteLast()
	if el.dData != 1 {
		t.Error("Expected element", 1, "have", el.dData)
	}
	ll.InsertFirst(2)
	ll.InsertFirst(3)
	ll.InsertFirst(4)
	el = ll.DeleteLast()
	if el.dData != 2 {
		t.Error("Expected element", 2, "have", el.dData)
	}
	ll.DisplayForward()
	ll.DisplayBackward()
}

func TestDeleteAfter(t *testing.T) {
	ll := NewDoubleLinkedList()
	el := ll.DeleteAfter(1)
	if el != nil {
		t.Error("Expected element", nil, "have", el)
	}
	ll.InsertFirst(1)
	el = ll.DeleteAfter(1)
	if el != nil {
		t.Error("Expected element", nil, "have", el)
	}
	ll.InsertFirst(2)
	ll.InsertFirst(3)
	ll.InsertFirst(4)
	el = ll.DeleteAfter(3)
	if el.dData != 2 {
		t.Error("Expected element", 2, "have", el.dData)
	}
	el = ll.DeleteAfter(10)
	if el != nil {
		t.Error("Expected element", nil, "have", el)
	}
	ll.DisplayForward()
	ll.DisplayBackward()
}
