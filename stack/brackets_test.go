package stack

import (
	"testing"
)

func TestBracketsExpectOk(t *testing.T) {
	err := BracketsChecker("s{sds{s} m[sss(s)ss]sd}")
	if err != nil {
		t.Error("expected err", nil, "have", err.Error())
	}
}

func TestBracketsNotEnoughBrackets(t *testing.T) {
	err := BracketsChecker("s{sds{s} m[sss(s)ss]sd")
	if err.Error() != "not enough closing brackets" {
		t.Error("expected err", "not enough closing brackets", "have", err.Error())
	}
}

func TestBracketsIncorrectBracket(t *testing.T) {
	err := BracketsChecker("s{sds{s} m[sss(s}ss]sd}")
	if err.Error() != "expected ), but received } in 16" {
		t.Error("expected err", "have expected ), but received } in 16", "have", err.Error())
	}
}
