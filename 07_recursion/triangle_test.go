package _7_recursion

import "testing"

func TestTriangle(t *testing.T) {
	sum := triangle(10000000)
	if sum != 15 {
		t.Error("Want sum", 15, "have", sum)
	}
}

func TestRecursionTriangle(t *testing.T) {
	sum := recursionTriangle(10000000)
	if sum != 15 {
		t.Error("Want sum", 15, "have", sum)
	}
}
