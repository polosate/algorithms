package parse_expression

import (
	"fmt"
	"testing"
)

func TestInToPost(t *testing.T) {
	expr := NewInToPost("A+B*(C–D)")
	expr.DoTrans()

	outStr := expr.outStr
	expectedOut := "ABCD-*+"

	fmt.Printf("Input %s\n", expr.inStr)
	if outStr != expectedOut {
		t.Error("Expected output", expectedOut, "have", outStr)
	}
}
