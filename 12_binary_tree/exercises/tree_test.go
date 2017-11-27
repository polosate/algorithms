package exercises

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tr := NewTree([]string{"A", "B", "C"})
	fmt.Println(tr.root.value)
	fmt.Println(tr.root.leftChild.root.value)
	fmt.Println(tr.root.rightChild.root.value)
	fmt.Println(tr.root.leftChild.root.leftChild.root.value)
	fmt.Println(tr.root.leftChild.root.rightChild.root.value)

}
