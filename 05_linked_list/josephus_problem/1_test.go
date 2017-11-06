package josephus_problem

import (
	"testing"
)

func Test(t *testing.T) {
	alive := solve(7, 3)
	if alive != 2 {
		t.Error("Expected alive", 2, "have", alive)
	}
}
