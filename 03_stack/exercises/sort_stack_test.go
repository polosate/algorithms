package exercises

import (
	"testing"
)

func TestSortStack(t *testing.T) {
	input := newStack1(5)
	input.push(3)
	input.push(5)
	input.push(2)
	input.push(4)
	input.push(1)
	sorted := sortStack(input)

	expected := []int64{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		data := sorted.pop()
		if data != expected[i] {
			t.Error("want", expected[i], "have", data)
		}
	}
}
