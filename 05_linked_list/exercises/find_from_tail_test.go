package exercises

import (
	"testing"
)

func TestFind(t *testing.T) {
	list := NewSingleList()
	value, err := FindFromTail(list, 7)
	if err == nil {
		t.Error("Want error", "have", nil)
	}
	if err.Error() != "list is empty" {
		t.Error("Want error", "list is empty", "have", err.Error())
	}
	if value != -1 {
		t.Error("Want value", -1, "have", value)
	}
	list.InsertFirst(1)
	list.InsertFirst(2)
	list.InsertFirst(3)
	list.InsertFirst(5)
	list.InsertFirst(8)
	list.InsertFirst(4)

	// first from tail
	value, err = FindFromTail(list, 0)
	if err != nil {
		t.Error("Want error", nil, "have", err.Error())
	}
	if value != 1 {
		t.Error("Want value", 1, "have", value)
	}

	value, err = FindFromTail(list, 1)
	if err != nil {
		t.Error("Want error", nil, "have", err.Error())
	}
	if value != 2 {
		t.Error("Want value", 2, "have", value)
	}

	// last from tail
	value, err = FindFromTail(list, 5)
	if err != nil {
		t.Error("Want error", nil, "have", err.Error())
	}
	if value != 4 {
		t.Error("Want value", 4, "have", value)
	}

	value, err = FindFromTail(list, 4)
	if err != nil {
		t.Error("Want error", nil, "have", err.Error())
	}
	if value != 8 {
		t.Error("Want value", 8, "have", value)
	}

	// from middle
	value, err = FindFromTail(list, 3)
	if err != nil {
		t.Error("Want error", nil, "have", err.Error())
	}
	if value != 5 {
		t.Error("Want value", 5, "have", value)
	}

	// k more than len(list)
	value, err = FindFromTail(list, 7)
	if err == nil {
		t.Error("Want error", "have", nil)
	}
	if err.Error() != "list has length less than k" {
		t.Error("Want error", "list has length less than k", "have", err.Error())
	}
	if value != -1 {
		t.Error("Want value", -1, "have", value)
	}
}
