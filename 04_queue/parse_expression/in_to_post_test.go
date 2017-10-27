package parse_expression

import (
	"fmt"
	"testing"
)

func TestInToPost_Case01(t *testing.T) {
	expr := NewInToPost("A+B*(C-D)")
	expr.DoTrans()

	outStr := expr.outStr
	expectedOut := "ABCD-*+"

	fmt.Printf("Input %s\n", expr.inStr)
	if outStr != expectedOut {
		t.Error("Expected output", expectedOut, "have", outStr)
	}
}

func TestInToPost_Case02(t *testing.T) {
	expr := NewInToPost("A*(B*C)")
	expr.DoTrans()

	outStr := expr.outStr
	expectedOut := "ABC**"

	fmt.Printf("Input %s\n", expr.inStr)
	if outStr != expectedOut {
		t.Error("Expected output", expectedOut, "have", outStr)
	}
}

func TestInToPost_Case03(t *testing.T) {
	expr := NewInToPost("A+B/C-D*(E+F)")
	expr.DoTrans()

	outStr := expr.outStr
	expectedOut := ""

	fmt.Printf("Input %s\n", expr.inStr)
	if outStr != expectedOut {
		t.Error("Expected output", expectedOut, "have", outStr)
	}
}
