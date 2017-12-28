package exercises

import (
	"testing"
)

func TestStackWithMin(t *testing.T) {
	s := newStackWithMin(10)
	s.push(5)
	s.push(6)
	s.push(3)
	s.push(7)

	min, _ := s.min()
	if min != 3 {
		t.Error()
	}
	s.pop()
	s.pop()
	min, _ = s.min()
	if min != 5 {
		t.Error()
	}
	s.pop()
	s.pop()
	_, err := s.min()
	if err == nil {
		t.Error()
	}
}
