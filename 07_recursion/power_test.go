package _7_recursion

import "testing"

func TestPower(t *testing.T) {
	a := Power(2, 4)
	if a != 16 {
		t.Error("Want", 16, "have", a)
	}

	b := Power(2, 5)
	if b != 32 {
		t.Error("Want", 32, "have", b)
	}
}
