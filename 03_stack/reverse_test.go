package stack

import (
	"testing"
)

func TestReverse(t *testing.T) {
	r, err := Reverse("golang")
	if err != nil {
		t.Error("expected err", nil, "have", err.Error())
	}

	if r != "gnalog" {
		t.Error("expected", "gnalog", "have", r)
	}
}
