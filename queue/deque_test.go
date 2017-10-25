package queue

import (
	"reflect"
	"testing"
)

func TestDeque(t *testing.T) {
	d := newDeque(5)
	empty := d.isEmpty()
	if !empty {
		t.Error("Expected isEmpty", true, "have", false)
	}

	d.insertLeft(1)
	d.insertLeft(2)
	d.insertRight(3)
	d.insertRight(4)
	d.insertLeft(5)

	expect := []interface{}{3, 4, nil, 5, 2, 1}
	if !reflect.DeepEqual(expect, d.data) {
		t.Error("Expected array", expect, "have", d.data)
	}

	e, _ := d.removeLeft()
	if !reflect.DeepEqual(e, 5) {
		t.Error("Expected element", 5, "have", e)
	}

	e, _ = d.removeRight()
	if !reflect.DeepEqual(e, 4) {
		t.Error("Expected element", 4, "have", e)
	}
}
