package exercises

import (
	"testing"
)

func TestMinBST(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	tree := minBST(a)
	tree.display()
}
