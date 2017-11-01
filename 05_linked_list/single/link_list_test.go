package single

import (
	"fmt"
	"testing"
)

func TestInsertDeleteFirst(t *testing.T) {
	ll := NewLinkList()
	if !ll.IsEmpty() {
		t.Error("Expected IsEmpty", true, "have", false)
	}
	ll.InsertFirst(1, 1.1)
	ll.InsertFirst(2, 2.2)
	ll.DisplayList()

	l := ll.DeleteFirst()
	if l.iData != 2 {
		t.Error("Expected link.iData", 2, "have", l.iData)
	}
	if l.dData != 2.2 {
		t.Error("Expected link.dData", 2.2, "have", l.iData)
	}
	if ll.IsEmpty() {
		t.Error("Expected IsEmpty", false, "have", true)
	}
}

func TestFind(t *testing.T) {
	ll := NewLinkList()
	el := ll.Find(3)
	if el != nil {
		t.Error("Expected next", nil, "have", el.next)
	}

	ll.InsertFirst(1, 1.1)
	ll.InsertFirst(2, 2.2)
	ll.InsertFirst(3, 3.3)
	ll.InsertFirst(4, 4.4)
	ll.InsertFirst(5, 5.5)

	el = ll.Find(3)
	if el.dData != 3.3 {
		t.Error("Expected next", 3.3, "have", el.dData)
	}

	el = ll.Find(6)
	if el != nil {
		t.Error("Expected next", nil, "have", el)
	}

	el = ll.Find(1)
	if el.dData != 1.1 {
		t.Error("Expected next", 1.1, "have", el.dData)
	}

	el = ll.Find(5)
	if el.dData != 5.5 {
		t.Error("Expected next", 5.5, "have", el.dData)
	}
}

func TestDelete(t *testing.T) {
	ll := NewLinkList()
	el := ll.Delete(3)
	if el != nil {
		t.Error("Expected next", nil, "have", el.next)
	}

	ll.InsertFirst(1, 1.1)
	ll.InsertFirst(2, 2.2)
	ll.InsertFirst(3, 3.3)
	ll.InsertFirst(4, 4.4)
	ll.InsertFirst(5, 5.5)

	el = ll.Delete(3)
	if el.dData != 3.3 {
		t.Error("Expected next", 3.3, "have", el.dData)
	}
	ll.DisplayList()

	el = ll.Delete(6)
	if el != nil {
		t.Error("Expected next", nil, "have", el)
	}
	ll.DisplayList()

	el = ll.Delete(1)
	if el.dData != 1.1 {
		t.Error("Expected next", 1.1, "have", el.dData)
	}
	ll.DisplayList()

	el = ll.Delete(5)
	if el.dData != 5.5 {
		t.Error("Expected next", 5.5, "have", el.dData)
	}
	ll.DisplayList()
}

func TestReverseEmptyList(t *testing.T) {
	ll := NewLinkList()
	if !ll.IsEmpty() {
		t.Error("Expected IsEmpty", true, "have", false)
	}
	ll.Reverse()
	if !ll.IsEmpty() {
		t.Error("Expected IsEmpty", true, "have", false)
	}

}

func TestReverseListWithOneItem(t *testing.T) {
	ll := NewLinkList()
	ll.InsertFirst(1, 1.1)
	ll.DisplayList()
	ll.Reverse()
	fmt.Println("Reversed:")
	ll.DisplayList()
	el := ll.DeleteFirst()
	if el.iData != 1 {
		t.Error("Expected iData", 1, "have", el.iData)
	}
	if !ll.IsEmpty() {
		t.Error("Expected IsEmpty", true, "have", false)
	}
}

func TestReverseListWithTwoItems(t *testing.T) {
	ll := NewLinkList()
	ll.InsertFirst(1, 1.1)
	ll.InsertFirst(2, 2.2)
	ll.DisplayList()
	ll.Reverse()
	fmt.Println("Reversed:")
	ll.DisplayList()
	el := ll.DeleteFirst()
	if el.iData != 1 {
		t.Error("Expected iData", 1, "have", el.iData)
	}
}

func TestReverse(t *testing.T) {
	ll := NewLinkList()

	ll.InsertFirst(1, 1.1)
	ll.InsertFirst(2, 2.2)
	ll.InsertFirst(3, 3.3)
	ll.InsertFirst(4, 4.4)
	ll.InsertFirst(5, 5.5)

	el := ll.DeleteFirst()
	if el.iData != 5 {
		t.Error("Expected iData", 5, "have", el.iData)
	}

	ll.DisplayList()

	ll.Reverse()
	fmt.Println("Reversed:")
	ll.DisplayList()

	el = ll.DeleteFirst()
	if el.iData != 1 {
		t.Error("Expected iData", 1, "have", el.iData)
	}
	el = ll.DeleteFirst()
	if el.iData != 2 {
		t.Error("Expected iData", 2, "have", el.iData)
	}
	el = ll.DeleteFirst()
	if el.iData != 3 {
		t.Error("Expected iData", 3, "have", el.iData)
	}

}
