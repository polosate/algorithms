package matrix

import (
	"testing"
)

func TestMatrixCreate(t *testing.T) {
	m := NewMatrix(5, 5)
	m.SetValue(0, 0, 5.5)
	v, _ := m.GetValue(0, 0)
	if v != 5.5 {
		t.Error("Expected value", 5.5, "have", v)
	}
	m.SetValue(1, 1, 5.5)
	v, _ = m.GetValue(1, 1)
	if v != 5.5 {
		t.Error("Expected value", 5.5, "have", v)
	}
	m.SetValue(2, 2, 5.5)
	v, _ = m.GetValue(2, 2)
	if v != 5.5 {
		t.Error("Expected value", 5.5, "have", v)
	}
	m.SetValue(4, 4, 5.5)
	v, _ = m.GetValue(4, 4)
	if v != 5.5 {
		t.Error("Expected value", 5.5, "have", v)
	}
	m.SetValue(5, 5, 5.5)
	_, err := m.GetValue(5, 5)
	if err.Error() != "incorrect coordinates" {
		t.Error("Expected error", "incorrect coordinates", "have", err.Error())
	}
	m.DisplayMatrix()
}
