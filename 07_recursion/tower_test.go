package _7_recursion

import (
	"reflect"
	"testing"
)

func TestTower(t *testing.T) {
	newTower := NewTower(4)
	from := newTower.from
	newTower.Move()

	if !reflect.DeepEqual(from, newTower.to) {
		t.Error("Want tower", from, "have", newTower.to)
	}
}
