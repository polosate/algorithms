package parse_expression

import (
	"testing"
)

func TestInToPostCase01(t *testing.T) {
	expr := NewInToPost("A+B*(C-D)")
	expr.DoTrans()

	outStr := expr.outStr
	expectedOut := "ABCD-*+"

	if outStr != expectedOut {
		t.Error("Expected output", expectedOut, "have", outStr)
	}
}

func TestInToPostCase02(t *testing.T) {
	expr := NewInToPost("A*(B*C)")
	expr.DoTrans()

	outStr := expr.outStr
	expectedOut := "ABC**"

	if outStr != expectedOut {
		t.Error("Expected output", expectedOut, "have", outStr)
	}
}

func TestInToPostCase03(t *testing.T) {
	expr := NewInToPost("A+B/C-D*(E+F)")
	expr.DoTrans()

	outStr := expr.outStr
	expectedOut := "ABC/+DEF+*-"

	if outStr != expectedOut {
		t.Error("Expected output", expectedOut, "have", outStr)
	}
}
