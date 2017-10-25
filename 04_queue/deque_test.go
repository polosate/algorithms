package queue

import (
	"reflect"
	"testing"
)

func TestDeque(t *testing.T) {
	d := newDeque(5)
	empty := d.isEmpty()
	if !empty {
		t.Error("Expected IsEmpty", true, "have", false)
	}
	size := d.size()
	if size != 0 {
		t.Error("Expected Size", 0, "have", size)
	}

	d.insertLeft(1)
	size = d.size()
	if size != 1 {
		t.Error("Expected Size", 1, "have", size)
	}
	d.insertLeft(2)
	size = d.size()
	if size != 2 {
		t.Error("Expected Size", 2, "have", size)
	}
	d.insertRight(3)
	size = d.size()
	if size != 3 {
		t.Error("Expected Size", 3, "have", size)
	}
	d.insertRight(4)
	size = d.size()
	if size != 4 {
		t.Error("Expected Size", 4, "have", size)
	}
	d.insertLeft(5)
	size = d.size()
	if size != 5 {
		t.Error("Expected Size", 5, "have", size)
	}

	expect := []interface{}{3, 4, nil, 5, 2, 1}
	if !reflect.DeepEqual(expect, d.data) {
		t.Error("Expected 01_array", expect, "have", d.data)
	}

	e, _ := d.removeLeft()
	if !reflect.DeepEqual(e, 5) {
		t.Error("Expected element", 5, "have", e)
	}
	size = d.size()
	if size != 4 {
		t.Error("Expected Size", 4, "have", size)
	}

	e, _ = d.removeRight()
	if !reflect.DeepEqual(e, 4) {
		t.Error("Expected element", 4, "have", e)
	}
	size = d.size()
	if size != 3 {
		t.Error("Expected Size", 3, "have", size)
	}
}
