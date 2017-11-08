package matrix

import (
	"testing"
	"fmt"
)

func TestMatrixCreate(t *testing.T) {
	m := NewMatrix(5, 5)
	m.DisplayMatrix()
	fmt.Println()
	m.SetValue(2, 2, 5)
	m.DisplayMatrix()

}
