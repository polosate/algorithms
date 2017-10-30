package single

import (
	"testing"
)

func Test(t *testing.T) {
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
