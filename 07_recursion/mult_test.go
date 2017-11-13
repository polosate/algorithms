package _7_recursion

import "testing"

func TestMultiply(t *testing.T) {
	a := multiply(5, 2)
	if a != 10 {
		t.Error("want", 10, "have", a)
	}
	a = multiply(5, 1)
	if a != 5 {
		t.Error("want", 5, "have", a)
	}
	a = multiply(5, 0)
	if a != 0 {
		t.Error("want", 0, "have", a)
	}
}
